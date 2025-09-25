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

package v20210125

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2021-01-25"

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


func NewAddUsersToWorkGroupRequest() (request *AddUsersToWorkGroupRequest) {
    request = &AddUsersToWorkGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "AddUsersToWorkGroup")
    
    
    return
}

func NewAddUsersToWorkGroupResponse() (response *AddUsersToWorkGroupResponse) {
    response = &AddUsersToWorkGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddUsersToWorkGroup
// This API is used to add users to working groups.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDGROUPID = "InvalidParameter.InvalidGroupId"
//  INVALIDPARAMETER_INVALIDUSERNAME = "InvalidParameter.InvalidUserName"
//  UNAUTHORIZEDOPERATION_ADDUSERSTOWORKGROUP = "UnauthorizedOperation.AddUsersToWorkgroup"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
func (c *Client) AddUsersToWorkGroup(request *AddUsersToWorkGroupRequest) (response *AddUsersToWorkGroupResponse, err error) {
    return c.AddUsersToWorkGroupWithContext(context.Background(), request)
}

// AddUsersToWorkGroup
// This API is used to add users to working groups.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDGROUPID = "InvalidParameter.InvalidGroupId"
//  INVALIDPARAMETER_INVALIDUSERNAME = "InvalidParameter.InvalidUserName"
//  UNAUTHORIZEDOPERATION_ADDUSERSTOWORKGROUP = "UnauthorizedOperation.AddUsersToWorkgroup"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
func (c *Client) AddUsersToWorkGroupWithContext(ctx context.Context, request *AddUsersToWorkGroupRequest) (response *AddUsersToWorkGroupResponse, err error) {
    if request == nil {
        request = NewAddUsersToWorkGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "AddUsersToWorkGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddUsersToWorkGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddUsersToWorkGroupResponse()
    err = c.Send(request, response)
    return
}

func NewAlterDMSDatabaseRequest() (request *AlterDMSDatabaseRequest) {
    request = &AlterDMSDatabaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "AlterDMSDatabase")
    
    
    return
}

func NewAlterDMSDatabaseResponse() (response *AlterDMSDatabaseResponse) {
    response = &AlterDMSDatabaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AlterDMSDatabase
// This API is used to update databases in the DMS metadata module.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AlterDMSDatabase(request *AlterDMSDatabaseRequest) (response *AlterDMSDatabaseResponse, err error) {
    return c.AlterDMSDatabaseWithContext(context.Background(), request)
}

// AlterDMSDatabase
// This API is used to update databases in the DMS metadata module.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AlterDMSDatabaseWithContext(ctx context.Context, request *AlterDMSDatabaseRequest) (response *AlterDMSDatabaseResponse, err error) {
    if request == nil {
        request = NewAlterDMSDatabaseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "AlterDMSDatabase")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AlterDMSDatabase require credential")
    }

    request.SetContext(ctx)
    
    response = NewAlterDMSDatabaseResponse()
    err = c.Send(request, response)
    return
}

func NewAttachUserPolicyRequest() (request *AttachUserPolicyRequest) {
    request = &AttachUserPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "AttachUserPolicy")
    
    
    return
}

func NewAttachUserPolicyResponse() (response *AttachUserPolicyResponse) {
    response = &AttachUserPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AttachUserPolicy
// This API is used to bind the authentication policy to the user.
//
// error code that may be returned:
//  FAILEDOPERATION_GETPOLICYFAILED = "FailedOperation.GetPolicyFailed"
//  FAILEDOPERATION_GETUSERINFOFAILED = "FailedOperation.GetUserInfoFailed"
//  FAILEDOPERATION_GETWORKGROUPINFOFAILED = "FailedOperation.GetWorkGroupInfoFailed"
//  FAILEDOPERATION_GRANTPOLICYFAILED = "FailedOperation.GrantPolicyFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_INVALIDACCESSPOLICY = "InvalidParameter.InvalidAccessPolicy"
//  INVALIDPARAMETER_INVALIDUSERNAME = "InvalidParameter.InvalidUserName"
//  INVALIDPARAMETER_INVALIDUSERTYPE = "InvalidParameter.InvalidUserType"
//  RESOURCESSOLDOUT_UNAUTHORIZEDGRANTPOLICY = "ResourcesSoldOut.UnauthorizedGrantPolicy"
//  UNAUTHORIZEDOPERATION_GRANTPOLICY = "UnauthorizedOperation.GrantPolicy"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
//  UNSUPPORTEDOPERATION_MODIFYOWNERUNSUPPORTED = "UnsupportedOperation.ModifyOwnerUnsupported"
func (c *Client) AttachUserPolicy(request *AttachUserPolicyRequest) (response *AttachUserPolicyResponse, err error) {
    return c.AttachUserPolicyWithContext(context.Background(), request)
}

// AttachUserPolicy
// This API is used to bind the authentication policy to the user.
//
// error code that may be returned:
//  FAILEDOPERATION_GETPOLICYFAILED = "FailedOperation.GetPolicyFailed"
//  FAILEDOPERATION_GETUSERINFOFAILED = "FailedOperation.GetUserInfoFailed"
//  FAILEDOPERATION_GETWORKGROUPINFOFAILED = "FailedOperation.GetWorkGroupInfoFailed"
//  FAILEDOPERATION_GRANTPOLICYFAILED = "FailedOperation.GrantPolicyFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_INVALIDACCESSPOLICY = "InvalidParameter.InvalidAccessPolicy"
//  INVALIDPARAMETER_INVALIDUSERNAME = "InvalidParameter.InvalidUserName"
//  INVALIDPARAMETER_INVALIDUSERTYPE = "InvalidParameter.InvalidUserType"
//  RESOURCESSOLDOUT_UNAUTHORIZEDGRANTPOLICY = "ResourcesSoldOut.UnauthorizedGrantPolicy"
//  UNAUTHORIZEDOPERATION_GRANTPOLICY = "UnauthorizedOperation.GrantPolicy"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
//  UNSUPPORTEDOPERATION_MODIFYOWNERUNSUPPORTED = "UnsupportedOperation.ModifyOwnerUnsupported"
func (c *Client) AttachUserPolicyWithContext(ctx context.Context, request *AttachUserPolicyRequest) (response *AttachUserPolicyResponse, err error) {
    if request == nil {
        request = NewAttachUserPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "AttachUserPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AttachUserPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewAttachUserPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewAttachWorkGroupPolicyRequest() (request *AttachWorkGroupPolicyRequest) {
    request = &AttachWorkGroupPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "AttachWorkGroupPolicy")
    
    
    return
}

func NewAttachWorkGroupPolicyResponse() (response *AttachWorkGroupPolicyResponse) {
    response = &AttachWorkGroupPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AttachWorkGroupPolicy
// This API is used to bind an authentication policy to a working group.
//
// error code that may be returned:
//  FAILEDOPERATION_GETPOLICYFAILED = "FailedOperation.GetPolicyFailed"
//  FAILEDOPERATION_GETUSERINFOFAILED = "FailedOperation.GetUserInfoFailed"
//  FAILEDOPERATION_GETWORKGROUPINFOFAILED = "FailedOperation.GetWorkGroupInfoFailed"
//  FAILEDOPERATION_GRANTPOLICYFAILED = "FailedOperation.GrantPolicyFailed"
//  FAILEDOPERATION_REVOKEPOLICYFAILED = "FailedOperation.RevokePolicyFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDACCESSPOLICY = "InvalidParameter.InvalidAccessPolicy"
//  INVALIDPARAMETER_INVALIDGROUPID = "InvalidParameter.InvalidGroupId"
//  UNAUTHORIZEDOPERATION_GRANTPOLICY = "UnauthorizedOperation.GrantPolicy"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
func (c *Client) AttachWorkGroupPolicy(request *AttachWorkGroupPolicyRequest) (response *AttachWorkGroupPolicyResponse, err error) {
    return c.AttachWorkGroupPolicyWithContext(context.Background(), request)
}

// AttachWorkGroupPolicy
// This API is used to bind an authentication policy to a working group.
//
// error code that may be returned:
//  FAILEDOPERATION_GETPOLICYFAILED = "FailedOperation.GetPolicyFailed"
//  FAILEDOPERATION_GETUSERINFOFAILED = "FailedOperation.GetUserInfoFailed"
//  FAILEDOPERATION_GETWORKGROUPINFOFAILED = "FailedOperation.GetWorkGroupInfoFailed"
//  FAILEDOPERATION_GRANTPOLICYFAILED = "FailedOperation.GrantPolicyFailed"
//  FAILEDOPERATION_REVOKEPOLICYFAILED = "FailedOperation.RevokePolicyFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDACCESSPOLICY = "InvalidParameter.InvalidAccessPolicy"
//  INVALIDPARAMETER_INVALIDGROUPID = "InvalidParameter.InvalidGroupId"
//  UNAUTHORIZEDOPERATION_GRANTPOLICY = "UnauthorizedOperation.GrantPolicy"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
func (c *Client) AttachWorkGroupPolicyWithContext(ctx context.Context, request *AttachWorkGroupPolicyRequest) (response *AttachWorkGroupPolicyResponse, err error) {
    if request == nil {
        request = NewAttachWorkGroupPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "AttachWorkGroupPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AttachWorkGroupPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewAttachWorkGroupPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewBindWorkGroupsToUserRequest() (request *BindWorkGroupsToUserRequest) {
    request = &BindWorkGroupsToUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "BindWorkGroupsToUser")
    
    
    return
}

func NewBindWorkGroupsToUserResponse() (response *BindWorkGroupsToUserResponse) {
    response = &BindWorkGroupsToUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindWorkGroupsToUser
// This API is used to bind working groups to users.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDGROUPID = "InvalidParameter.InvalidGroupId"
//  INVALIDPARAMETER_INVALIDUSERNAME = "InvalidParameter.InvalidUserName"
//  UNAUTHORIZEDOPERATION_BINDWORKGROUPSTOUSER = "UnauthorizedOperation.BindWorkgroupsToUser"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
//  UNSUPPORTEDOPERATION_MODIFYOWNERUNSUPPORTED = "UnsupportedOperation.ModifyOwnerUnsupported"
func (c *Client) BindWorkGroupsToUser(request *BindWorkGroupsToUserRequest) (response *BindWorkGroupsToUserResponse, err error) {
    return c.BindWorkGroupsToUserWithContext(context.Background(), request)
}

// BindWorkGroupsToUser
// This API is used to bind working groups to users.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDGROUPID = "InvalidParameter.InvalidGroupId"
//  INVALIDPARAMETER_INVALIDUSERNAME = "InvalidParameter.InvalidUserName"
//  UNAUTHORIZEDOPERATION_BINDWORKGROUPSTOUSER = "UnauthorizedOperation.BindWorkgroupsToUser"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
//  UNSUPPORTEDOPERATION_MODIFYOWNERUNSUPPORTED = "UnsupportedOperation.ModifyOwnerUnsupported"
func (c *Client) BindWorkGroupsToUserWithContext(ctx context.Context, request *BindWorkGroupsToUserRequest) (response *BindWorkGroupsToUserResponse, err error) {
    if request == nil {
        request = NewBindWorkGroupsToUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "BindWorkGroupsToUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindWorkGroupsToUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindWorkGroupsToUserResponse()
    err = c.Send(request, response)
    return
}

func NewCancelSparkSessionBatchSQLRequest() (request *CancelSparkSessionBatchSQLRequest) {
    request = &CancelSparkSessionBatchSQLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CancelSparkSessionBatchSQL")
    
    
    return
}

func NewCancelSparkSessionBatchSQLResponse() (response *CancelSparkSessionBatchSQLResponse) {
    response = &CancelSparkSessionBatchSQLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CancelSparkSessionBatchSQL
// This API is used to cancel a Spark SQL batch task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BATCHSQLTASKNOTFOUND = "ResourceNotFound.BatchSQLTaskNotFound"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  RESOURCENOTFOUND_SESSIONSTATEDEAD = "ResourceNotFound.SessionStateDead"
func (c *Client) CancelSparkSessionBatchSQL(request *CancelSparkSessionBatchSQLRequest) (response *CancelSparkSessionBatchSQLResponse, err error) {
    return c.CancelSparkSessionBatchSQLWithContext(context.Background(), request)
}

// CancelSparkSessionBatchSQL
// This API is used to cancel a Spark SQL batch task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BATCHSQLTASKNOTFOUND = "ResourceNotFound.BatchSQLTaskNotFound"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  RESOURCENOTFOUND_SESSIONSTATEDEAD = "ResourceNotFound.SessionStateDead"
func (c *Client) CancelSparkSessionBatchSQLWithContext(ctx context.Context, request *CancelSparkSessionBatchSQLRequest) (response *CancelSparkSessionBatchSQLResponse, err error) {
    if request == nil {
        request = NewCancelSparkSessionBatchSQLRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "CancelSparkSessionBatchSQL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelSparkSessionBatchSQL require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelSparkSessionBatchSQLResponse()
    err = c.Send(request, response)
    return
}

func NewCancelTaskRequest() (request *CancelTaskRequest) {
    request = &CancelTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CancelTask")
    
    
    return
}

func NewCancelTaskResponse() (response *CancelTaskResponse) {
    response = &CancelTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CancelTask
// This API is used to cancel a task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ANOTHERREQUESTPROCESSING = "FailedOperation.AnotherRequestProcessing"
//  FAILEDOPERATION_HTTPCLIENTDOREQUESTFAILED = "FailedOperation.HttpClientDoRequestFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_INVALIDTASKID = "InvalidParameter.InvalidTaskId"
//  INVALIDPARAMETER_PARAMETERBASE64DECODEFAILED = "InvalidParameter.ParameterBase64DecodeFailed"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  INVALIDPARAMETER_TASKALREADYFINISHED = "InvalidParameter.TaskAlreadyFinished"
//  RESOURCENOTFOUND_TASKALREADYFAILED = "ResourceNotFound.TaskAlreadyFailed"
//  RESOURCENOTFOUND_TASKALREADYFINISHED = "ResourceNotFound.TaskAlreadyFinished"
func (c *Client) CancelTask(request *CancelTaskRequest) (response *CancelTaskResponse, err error) {
    return c.CancelTaskWithContext(context.Background(), request)
}

// CancelTask
// This API is used to cancel a task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ANOTHERREQUESTPROCESSING = "FailedOperation.AnotherRequestProcessing"
//  FAILEDOPERATION_HTTPCLIENTDOREQUESTFAILED = "FailedOperation.HttpClientDoRequestFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_INVALIDTASKID = "InvalidParameter.InvalidTaskId"
//  INVALIDPARAMETER_PARAMETERBASE64DECODEFAILED = "InvalidParameter.ParameterBase64DecodeFailed"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  INVALIDPARAMETER_TASKALREADYFINISHED = "InvalidParameter.TaskAlreadyFinished"
//  RESOURCENOTFOUND_TASKALREADYFAILED = "ResourceNotFound.TaskAlreadyFailed"
//  RESOURCENOTFOUND_TASKALREADYFINISHED = "ResourceNotFound.TaskAlreadyFinished"
func (c *Client) CancelTaskWithContext(ctx context.Context, request *CancelTaskRequest) (response *CancelTaskResponse, err error) {
    if request == nil {
        request = NewCancelTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "CancelTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCheckDataEngineConfigPairsValidityRequest() (request *CheckDataEngineConfigPairsValidityRequest) {
    request = &CheckDataEngineConfigPairsValidityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CheckDataEngineConfigPairsValidity")
    
    
    return
}

func NewCheckDataEngineConfigPairsValidityResponse() (response *CheckDataEngineConfigPairsValidityResponse) {
    response = &CheckDataEngineConfigPairsValidityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckDataEngineConfigPairsValidity
// This API is used to check the validity of the engine's user-defined parameters.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_IMAGEENGINETYPENOTMATCH = "InvalidParameter.ImageEngineTypeNotMatch"
//  INVALIDPARAMETER_IMAGEISPUBLICNOTMATCH = "InvalidParameter.ImageIsPublicNotMatch"
//  INVALIDPARAMETER_IMAGEPARAMETERSUBMITMETHODNOTMATCH = "InvalidParameter.ImageParameterSubmitMethodNotMatch"
//  INVALIDPARAMETER_IMAGEPARAMETERTYPENOTMATCH = "InvalidParameter.ImageParameterTypeNotMatch"
//  INVALIDPARAMETER_IMAGESTATENOTMATCH = "InvalidParameter.ImageStateNotMatch"
//  INVALIDPARAMETER_IMAGEUSERRECORDSTYPENOTMATCH = "InvalidParameter.ImageUserRecordsTypeNotMatch"
//  INVALIDPARAMETER_INVALIDWHITELISTKEY = "InvalidParameter.InvalidWhiteListKey"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_IMAGEVERSIONNOTFOUND = "ResourceNotFound.ImageVersionNotFound"
func (c *Client) CheckDataEngineConfigPairsValidity(request *CheckDataEngineConfigPairsValidityRequest) (response *CheckDataEngineConfigPairsValidityResponse, err error) {
    return c.CheckDataEngineConfigPairsValidityWithContext(context.Background(), request)
}

// CheckDataEngineConfigPairsValidity
// This API is used to check the validity of the engine's user-defined parameters.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_IMAGEENGINETYPENOTMATCH = "InvalidParameter.ImageEngineTypeNotMatch"
//  INVALIDPARAMETER_IMAGEISPUBLICNOTMATCH = "InvalidParameter.ImageIsPublicNotMatch"
//  INVALIDPARAMETER_IMAGEPARAMETERSUBMITMETHODNOTMATCH = "InvalidParameter.ImageParameterSubmitMethodNotMatch"
//  INVALIDPARAMETER_IMAGEPARAMETERTYPENOTMATCH = "InvalidParameter.ImageParameterTypeNotMatch"
//  INVALIDPARAMETER_IMAGESTATENOTMATCH = "InvalidParameter.ImageStateNotMatch"
//  INVALIDPARAMETER_IMAGEUSERRECORDSTYPENOTMATCH = "InvalidParameter.ImageUserRecordsTypeNotMatch"
//  INVALIDPARAMETER_INVALIDWHITELISTKEY = "InvalidParameter.InvalidWhiteListKey"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_IMAGEVERSIONNOTFOUND = "ResourceNotFound.ImageVersionNotFound"
func (c *Client) CheckDataEngineConfigPairsValidityWithContext(ctx context.Context, request *CheckDataEngineConfigPairsValidityRequest) (response *CheckDataEngineConfigPairsValidityResponse, err error) {
    if request == nil {
        request = NewCheckDataEngineConfigPairsValidityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "CheckDataEngineConfigPairsValidity")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckDataEngineConfigPairsValidity require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckDataEngineConfigPairsValidityResponse()
    err = c.Send(request, response)
    return
}

func NewCheckDataEngineImageCanBeRollbackRequest() (request *CheckDataEngineImageCanBeRollbackRequest) {
    request = &CheckDataEngineImageCanBeRollbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CheckDataEngineImageCanBeRollback")
    
    
    return
}

func NewCheckDataEngineImageCanBeRollbackResponse() (response *CheckDataEngineImageCanBeRollbackResponse) {
    response = &CheckDataEngineImageCanBeRollbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckDataEngineImageCanBeRollback
// This API is used to check whether the cluster can be rolled back.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAENGINEIMAGEOPERATENOTMATCH = "InvalidParameter.DataEngineImageOperateNotMatch"
//  INVALIDPARAMETER_IMAGEENGINETYPENOTMATCH = "InvalidParameter.ImageEngineTypeNotMatch"
//  INVALIDPARAMETER_IMAGEISPUBLICNOTMATCH = "InvalidParameter.ImageIsPublicNotMatch"
//  INVALIDPARAMETER_IMAGESTATENOTMATCH = "InvalidParameter.ImageStateNotMatch"
//  INVALIDPARAMETER_IMAGEUSERRECORDSTYPENOTMATCH = "InvalidParameter.ImageUserRecordsTypeNotMatch"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTUNIQUE = "ResourceNotFound.DataEngineNotUnique"
//  RESOURCENOTFOUND_IMAGEVERSIONNOTACTIVITY = "ResourceNotFound.ImageVersionNotActivity"
//  RESOURCENOTFOUND_IMAGEVERSIONNOTUNIQUE = "ResourceNotFound.ImageVersionNotUnique"
func (c *Client) CheckDataEngineImageCanBeRollback(request *CheckDataEngineImageCanBeRollbackRequest) (response *CheckDataEngineImageCanBeRollbackResponse, err error) {
    return c.CheckDataEngineImageCanBeRollbackWithContext(context.Background(), request)
}

// CheckDataEngineImageCanBeRollback
// This API is used to check whether the cluster can be rolled back.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAENGINEIMAGEOPERATENOTMATCH = "InvalidParameter.DataEngineImageOperateNotMatch"
//  INVALIDPARAMETER_IMAGEENGINETYPENOTMATCH = "InvalidParameter.ImageEngineTypeNotMatch"
//  INVALIDPARAMETER_IMAGEISPUBLICNOTMATCH = "InvalidParameter.ImageIsPublicNotMatch"
//  INVALIDPARAMETER_IMAGESTATENOTMATCH = "InvalidParameter.ImageStateNotMatch"
//  INVALIDPARAMETER_IMAGEUSERRECORDSTYPENOTMATCH = "InvalidParameter.ImageUserRecordsTypeNotMatch"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTUNIQUE = "ResourceNotFound.DataEngineNotUnique"
//  RESOURCENOTFOUND_IMAGEVERSIONNOTACTIVITY = "ResourceNotFound.ImageVersionNotActivity"
//  RESOURCENOTFOUND_IMAGEVERSIONNOTUNIQUE = "ResourceNotFound.ImageVersionNotUnique"
func (c *Client) CheckDataEngineImageCanBeRollbackWithContext(ctx context.Context, request *CheckDataEngineImageCanBeRollbackRequest) (response *CheckDataEngineImageCanBeRollbackResponse, err error) {
    if request == nil {
        request = NewCheckDataEngineImageCanBeRollbackRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "CheckDataEngineImageCanBeRollback")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckDataEngineImageCanBeRollback require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckDataEngineImageCanBeRollbackResponse()
    err = c.Send(request, response)
    return
}

func NewCheckDataEngineImageCanBeUpgradeRequest() (request *CheckDataEngineImageCanBeUpgradeRequest) {
    request = &CheckDataEngineImageCanBeUpgradeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CheckDataEngineImageCanBeUpgrade")
    
    
    return
}

func NewCheckDataEngineImageCanBeUpgradeResponse() (response *CheckDataEngineImageCanBeUpgradeResponse) {
    response = &CheckDataEngineImageCanBeUpgradeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckDataEngineImageCanBeUpgrade
// This API is used to check whether the cluster image can be upgraded.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_IMAGEENGINETYPENOTMATCH = "InvalidParameter.ImageEngineTypeNotMatch"
//  INVALIDPARAMETER_IMAGEISPUBLICNOTMATCH = "InvalidParameter.ImageIsPublicNotMatch"
//  INVALIDPARAMETER_IMAGESTATENOTMATCH = "InvalidParameter.ImageStateNotMatch"
//  INVALIDPARAMETER_IMAGEUSERRECORDSTYPENOTMATCH = "InvalidParameter.ImageUserRecordsTypeNotMatch"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTUNIQUE = "ResourceNotFound.DataEngineNotUnique"
//  RESOURCENOTFOUND_IMAGEVERSIONNOTFOUND = "ResourceNotFound.ImageVersionNotFound"
func (c *Client) CheckDataEngineImageCanBeUpgrade(request *CheckDataEngineImageCanBeUpgradeRequest) (response *CheckDataEngineImageCanBeUpgradeResponse, err error) {
    return c.CheckDataEngineImageCanBeUpgradeWithContext(context.Background(), request)
}

// CheckDataEngineImageCanBeUpgrade
// This API is used to check whether the cluster image can be upgraded.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_IMAGEENGINETYPENOTMATCH = "InvalidParameter.ImageEngineTypeNotMatch"
//  INVALIDPARAMETER_IMAGEISPUBLICNOTMATCH = "InvalidParameter.ImageIsPublicNotMatch"
//  INVALIDPARAMETER_IMAGESTATENOTMATCH = "InvalidParameter.ImageStateNotMatch"
//  INVALIDPARAMETER_IMAGEUSERRECORDSTYPENOTMATCH = "InvalidParameter.ImageUserRecordsTypeNotMatch"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTUNIQUE = "ResourceNotFound.DataEngineNotUnique"
//  RESOURCENOTFOUND_IMAGEVERSIONNOTFOUND = "ResourceNotFound.ImageVersionNotFound"
func (c *Client) CheckDataEngineImageCanBeUpgradeWithContext(ctx context.Context, request *CheckDataEngineImageCanBeUpgradeRequest) (response *CheckDataEngineImageCanBeUpgradeResponse, err error) {
    if request == nil {
        request = NewCheckDataEngineImageCanBeUpgradeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "CheckDataEngineImageCanBeUpgrade")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckDataEngineImageCanBeUpgrade require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckDataEngineImageCanBeUpgradeResponse()
    err = c.Send(request, response)
    return
}

func NewCheckGrantedPermissionRequest() (request *CheckGrantedPermissionRequest) {
    request = &CheckGrantedPermissionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CheckGrantedPermission")
    
    
    return
}

func NewCheckGrantedPermissionResponse() (response *CheckGrantedPermissionResponse) {
    response = &CheckGrantedPermissionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckGrantedPermission
// This API is used to check the permission status.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CheckGrantedPermission(request *CheckGrantedPermissionRequest) (response *CheckGrantedPermissionResponse, err error) {
    return c.CheckGrantedPermissionWithContext(context.Background(), request)
}

// CheckGrantedPermission
// This API is used to check the permission status.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CheckGrantedPermissionWithContext(ctx context.Context, request *CheckGrantedPermissionRequest) (response *CheckGrantedPermissionResponse, err error) {
    if request == nil {
        request = NewCheckGrantedPermissionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "CheckGrantedPermission")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckGrantedPermission require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckGrantedPermissionResponse()
    err = c.Send(request, response)
    return
}

func NewCopyDLCTableRequest() (request *CopyDLCTableRequest) {
    request = &CopyDLCTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CopyDLCTable")
    
    
    return
}

func NewCopyDLCTableResponse() (response *CopyDLCTableResponse) {
    response = &CopyDLCTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CopyDLCTable
// This API is used to copy a table.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CopyDLCTable(request *CopyDLCTableRequest) (response *CopyDLCTableResponse, err error) {
    return c.CopyDLCTableWithContext(context.Background(), request)
}

// CopyDLCTable
// This API is used to copy a table.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CopyDLCTableWithContext(ctx context.Context, request *CopyDLCTableRequest) (response *CopyDLCTableResponse, err error) {
    if request == nil {
        request = NewCopyDLCTableRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "CopyDLCTable")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CopyDLCTable require credential")
    }

    request.SetContext(ctx)
    
    response = NewCopyDLCTableResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCHDFSBindingProductRequest() (request *CreateCHDFSBindingProductRequest) {
    request = &CreateCHDFSBindingProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CreateCHDFSBindingProduct")
    
    
    return
}

func NewCreateCHDFSBindingProductResponse() (response *CreateCHDFSBindingProductResponse) {
    response = &CreateCHDFSBindingProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCHDFSBindingProduct
// This API is used to create metadata acceleration buckets and the binding relationship between products.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateCHDFSBindingProduct(request *CreateCHDFSBindingProductRequest) (response *CreateCHDFSBindingProductResponse, err error) {
    return c.CreateCHDFSBindingProductWithContext(context.Background(), request)
}

// CreateCHDFSBindingProduct
// This API is used to create metadata acceleration buckets and the binding relationship between products.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateCHDFSBindingProductWithContext(ctx context.Context, request *CreateCHDFSBindingProductRequest) (response *CreateCHDFSBindingProductResponse, err error) {
    if request == nil {
        request = NewCreateCHDFSBindingProductRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "CreateCHDFSBindingProduct")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCHDFSBindingProduct require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCHDFSBindingProductResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDLCTableRequest() (request *CreateDLCTableRequest) {
    request = &CreateDLCTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CreateDLCTable")
    
    
    return
}

func NewCreateDLCTableResponse() (response *CreateDLCTableResponse) {
    response = &CreateDLCTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDLCTable
// This API is used to create a table.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateDLCTable(request *CreateDLCTableRequest) (response *CreateDLCTableResponse, err error) {
    return c.CreateDLCTableWithContext(context.Background(), request)
}

// CreateDLCTable
// This API is used to create a table.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateDLCTableWithContext(ctx context.Context, request *CreateDLCTableRequest) (response *CreateDLCTableResponse, err error) {
    if request == nil {
        request = NewCreateDLCTableRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "CreateDLCTable")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDLCTable require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDLCTableResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDMSDatabaseRequest() (request *CreateDMSDatabaseRequest) {
    request = &CreateDMSDatabaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CreateDMSDatabase")
    
    
    return
}

func NewCreateDMSDatabaseResponse() (response *CreateDMSDatabaseResponse) {
    response = &CreateDMSDatabaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDMSDatabase
// This API is used to create databases in the DMS metadata module.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateDMSDatabase(request *CreateDMSDatabaseRequest) (response *CreateDMSDatabaseResponse, err error) {
    return c.CreateDMSDatabaseWithContext(context.Background(), request)
}

// CreateDMSDatabase
// This API is used to create databases in the DMS metadata module.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateDMSDatabaseWithContext(ctx context.Context, request *CreateDMSDatabaseRequest) (response *CreateDMSDatabaseResponse, err error) {
    if request == nil {
        request = NewCreateDMSDatabaseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "CreateDMSDatabase")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDMSDatabase require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDMSDatabaseResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDataEngineRequest() (request *CreateDataEngineRequest) {
    request = &CreateDataEngineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CreateDataEngine")
    
    
    return
}

func NewCreateDataEngineResponse() (response *CreateDataEngineResponse) {
    response = &CreateDataEngineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDataEngine
// This API is used to create a data engine.
//
// error code that may be returned:
//  FAILEDOPERATION_ABNORMALORDERSTATUS = "FailedOperation.AbnormalOrderStatus"
//  FAILEDOPERATION_ANOTHERCREATEPROCESSRUNNING = "FailedOperation.AnotherCreateProcessRunning"
//  FAILEDOPERATION_ANOTHERPROCESSRUNNING = "FailedOperation.AnotherProcessRunning"
//  FAILEDOPERATION_ANOTHERREQUESTPROCESSING = "FailedOperation.AnotherRequestProcessing"
//  FAILEDOPERATION_BALANCENOTENOUGH = "FailedOperation.BalanceNotEnough"
//  FAILEDOPERATION_BILLINGSYSTEMERROR = "FailedOperation.BillingSystemError"
//  FAILEDOPERATION_BINDTOOMANYTAGS = "FailedOperation.BindTooManyTags"
//  FAILEDOPERATION_CREATEDATAENGINEFAILED = "FailedOperation.CreateDataEngineFailed"
//  FAILEDOPERATION_DELIVERGOODSFAILED = "FailedOperation.DeliverGoodsFailed"
//  FAILEDOPERATION_DUPLICATETAGKEY = "FailedOperation.DuplicateTagKey"
//  FAILEDOPERATION_FEEDEDUCTIONFAILED = "FailedOperation.FeeDeductionFailed"
//  FAILEDOPERATION_GETPRODUCTINFORMATIONFAILED = "FailedOperation.GetProductInformationFailed"
//  FAILEDOPERATION_ILLEGALRESOURCE = "FailedOperation.IllegalResource"
//  FAILEDOPERATION_ILLEGALTAGKEY = "FailedOperation.IllegalTagKey"
//  FAILEDOPERATION_ILLEGALTAGVALUE = "FailedOperation.IllegalTagValue"
//  FAILEDOPERATION_INQUIREPRICEFAILED = "FailedOperation.InquirePriceFailed"
//  FAILEDOPERATION_MODIFYINSTANCEFAILED = "FailedOperation.ModifyInstanceFailed"
//  FAILEDOPERATION_NOPERMISSION = "FailedOperation.NoPermission"
//  FAILEDOPERATION_NOREALNAMEAUTHENTICATION = "FailedOperation.NoRealNameAuthentication"
//  FAILEDOPERATION_NUMBEREXCEEDLIMIT = "FailedOperation.NumberExceedLimit"
//  FAILEDOPERATION_PARAMETERVALIDATIONFAILED = "FailedOperation.ParameterValidationFailed"
//  FAILEDOPERATION_REFUNDDEPOSITFAILED = "FailedOperation.RefundDepositFailed"
//  FAILEDOPERATION_TAGALREADYATTACHED = "FailedOperation.TagAlreadyAttached"
//  FAILEDOPERATION_TAGKEYTOOLONG = "FailedOperation.TagKeyTooLong"
//  FAILEDOPERATION_TAGNOTEXIST = "FailedOperation.TagNotExist"
//  FAILEDOPERATION_TAGVALUETOOLONG = "FailedOperation.TagValueTooLong"
//  FAILEDOPERATION_TOOMANYRESOURCES = "FailedOperation.TooManyResources"
//  FAILEDOPERATION_TOOMANYTAGS = "FailedOperation.TooManyTags"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_DATAENGINEMODENOTMATCH = "InvalidParameter.DataEngineModeNotMatch"
//  INVALIDPARAMETER_DATAENGINESIZENOTMATCH = "InvalidParameter.DataEngineSizeNotMatch"
//  INVALIDPARAMETER_DUPLICATEDATAENGINENAME = "InvalidParameter.DuplicateDataEngineName"
//  INVALIDPARAMETER_IMAGEPARAMETERNOTFOUND = "InvalidParameter.ImageParameterNotFound"
//  INVALIDPARAMETER_INVALIDDATAENGINECIDRFORMAT = "InvalidParameter.InvalidDataEngineCidrFormat"
//  INVALIDPARAMETER_INVALIDDATAENGINEMODE = "InvalidParameter.InvalidDataEngineMode"
//  INVALIDPARAMETER_INVALIDDATAENGINENAME = "InvalidParameter.InvalidDataEngineName"
//  INVALIDPARAMETER_INVALIDDATAENGINESPECS = "InvalidParameter.InvalidDataEngineSpecs"
//  INVALIDPARAMETER_INVALIDDATAENGINETIMESPAN = "InvalidParameter.InvalidDataEngineTimeSpan"
//  INVALIDPARAMETER_INVALIDDATAENGINETIMEUNIT = "InvalidParameter.InvalidDataEngineTimeUnit"
//  INVALIDPARAMETER_INVALIDENGINETYPE = "InvalidParameter.InvalidEngineType"
//  INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPayMode"
//  INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_VPCCIDRFORMATERROR = "InvalidParameter.VpcCidrFormatError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTUNIQUE = "ResourceNotFound.DataEngineNotUnique"
//  UNAUTHORIZEDOPERATION_NOPAYMENTAUTHORITY = "UnauthorizedOperation.NoPaymentAuthority"
func (c *Client) CreateDataEngine(request *CreateDataEngineRequest) (response *CreateDataEngineResponse, err error) {
    return c.CreateDataEngineWithContext(context.Background(), request)
}

// CreateDataEngine
// This API is used to create a data engine.
//
// error code that may be returned:
//  FAILEDOPERATION_ABNORMALORDERSTATUS = "FailedOperation.AbnormalOrderStatus"
//  FAILEDOPERATION_ANOTHERCREATEPROCESSRUNNING = "FailedOperation.AnotherCreateProcessRunning"
//  FAILEDOPERATION_ANOTHERPROCESSRUNNING = "FailedOperation.AnotherProcessRunning"
//  FAILEDOPERATION_ANOTHERREQUESTPROCESSING = "FailedOperation.AnotherRequestProcessing"
//  FAILEDOPERATION_BALANCENOTENOUGH = "FailedOperation.BalanceNotEnough"
//  FAILEDOPERATION_BILLINGSYSTEMERROR = "FailedOperation.BillingSystemError"
//  FAILEDOPERATION_BINDTOOMANYTAGS = "FailedOperation.BindTooManyTags"
//  FAILEDOPERATION_CREATEDATAENGINEFAILED = "FailedOperation.CreateDataEngineFailed"
//  FAILEDOPERATION_DELIVERGOODSFAILED = "FailedOperation.DeliverGoodsFailed"
//  FAILEDOPERATION_DUPLICATETAGKEY = "FailedOperation.DuplicateTagKey"
//  FAILEDOPERATION_FEEDEDUCTIONFAILED = "FailedOperation.FeeDeductionFailed"
//  FAILEDOPERATION_GETPRODUCTINFORMATIONFAILED = "FailedOperation.GetProductInformationFailed"
//  FAILEDOPERATION_ILLEGALRESOURCE = "FailedOperation.IllegalResource"
//  FAILEDOPERATION_ILLEGALTAGKEY = "FailedOperation.IllegalTagKey"
//  FAILEDOPERATION_ILLEGALTAGVALUE = "FailedOperation.IllegalTagValue"
//  FAILEDOPERATION_INQUIREPRICEFAILED = "FailedOperation.InquirePriceFailed"
//  FAILEDOPERATION_MODIFYINSTANCEFAILED = "FailedOperation.ModifyInstanceFailed"
//  FAILEDOPERATION_NOPERMISSION = "FailedOperation.NoPermission"
//  FAILEDOPERATION_NOREALNAMEAUTHENTICATION = "FailedOperation.NoRealNameAuthentication"
//  FAILEDOPERATION_NUMBEREXCEEDLIMIT = "FailedOperation.NumberExceedLimit"
//  FAILEDOPERATION_PARAMETERVALIDATIONFAILED = "FailedOperation.ParameterValidationFailed"
//  FAILEDOPERATION_REFUNDDEPOSITFAILED = "FailedOperation.RefundDepositFailed"
//  FAILEDOPERATION_TAGALREADYATTACHED = "FailedOperation.TagAlreadyAttached"
//  FAILEDOPERATION_TAGKEYTOOLONG = "FailedOperation.TagKeyTooLong"
//  FAILEDOPERATION_TAGNOTEXIST = "FailedOperation.TagNotExist"
//  FAILEDOPERATION_TAGVALUETOOLONG = "FailedOperation.TagValueTooLong"
//  FAILEDOPERATION_TOOMANYRESOURCES = "FailedOperation.TooManyResources"
//  FAILEDOPERATION_TOOMANYTAGS = "FailedOperation.TooManyTags"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_DATAENGINEMODENOTMATCH = "InvalidParameter.DataEngineModeNotMatch"
//  INVALIDPARAMETER_DATAENGINESIZENOTMATCH = "InvalidParameter.DataEngineSizeNotMatch"
//  INVALIDPARAMETER_DUPLICATEDATAENGINENAME = "InvalidParameter.DuplicateDataEngineName"
//  INVALIDPARAMETER_IMAGEPARAMETERNOTFOUND = "InvalidParameter.ImageParameterNotFound"
//  INVALIDPARAMETER_INVALIDDATAENGINECIDRFORMAT = "InvalidParameter.InvalidDataEngineCidrFormat"
//  INVALIDPARAMETER_INVALIDDATAENGINEMODE = "InvalidParameter.InvalidDataEngineMode"
//  INVALIDPARAMETER_INVALIDDATAENGINENAME = "InvalidParameter.InvalidDataEngineName"
//  INVALIDPARAMETER_INVALIDDATAENGINESPECS = "InvalidParameter.InvalidDataEngineSpecs"
//  INVALIDPARAMETER_INVALIDDATAENGINETIMESPAN = "InvalidParameter.InvalidDataEngineTimeSpan"
//  INVALIDPARAMETER_INVALIDDATAENGINETIMEUNIT = "InvalidParameter.InvalidDataEngineTimeUnit"
//  INVALIDPARAMETER_INVALIDENGINETYPE = "InvalidParameter.InvalidEngineType"
//  INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPayMode"
//  INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_VPCCIDRFORMATERROR = "InvalidParameter.VpcCidrFormatError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTUNIQUE = "ResourceNotFound.DataEngineNotUnique"
//  UNAUTHORIZEDOPERATION_NOPAYMENTAUTHORITY = "UnauthorizedOperation.NoPaymentAuthority"
func (c *Client) CreateDataEngineWithContext(ctx context.Context, request *CreateDataEngineRequest) (response *CreateDataEngineResponse, err error) {
    if request == nil {
        request = NewCreateDataEngineRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "CreateDataEngine")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDataEngine require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDataEngineResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInternalTableRequest() (request *CreateInternalTableRequest) {
    request = &CreateInternalTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CreateInternalTable")
    
    
    return
}

func NewCreateInternalTableResponse() (response *CreateInternalTableResponse) {
    response = &CreateInternalTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateInternalTable
// This API is used to create a managed internal table. It has been disused.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateInternalTable(request *CreateInternalTableRequest) (response *CreateInternalTableResponse, err error) {
    return c.CreateInternalTableWithContext(context.Background(), request)
}

// CreateInternalTable
// This API is used to create a managed internal table. It has been disused.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateInternalTableWithContext(ctx context.Context, request *CreateInternalTableRequest) (response *CreateInternalTableResponse, err error) {
    if request == nil {
        request = NewCreateInternalTableRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "CreateInternalTable")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInternalTable require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInternalTableResponse()
    err = c.Send(request, response)
    return
}

func NewCreateResultDownloadRequest() (request *CreateResultDownloadRequest) {
    request = &CreateResultDownloadRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CreateResultDownload")
    
    
    return
}

func NewCreateResultDownloadResponse() (response *CreateResultDownloadResponse) {
    response = &CreateResultDownloadResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateResultDownload
// This API is used to create a query result download task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOPERMISSION = "FailedOperation.NoPermission"
func (c *Client) CreateResultDownload(request *CreateResultDownloadRequest) (response *CreateResultDownloadResponse, err error) {
    return c.CreateResultDownloadWithContext(context.Background(), request)
}

// CreateResultDownload
// This API is used to create a query result download task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOPERMISSION = "FailedOperation.NoPermission"
func (c *Client) CreateResultDownloadWithContext(ctx context.Context, request *CreateResultDownloadRequest) (response *CreateResultDownloadResponse, err error) {
    if request == nil {
        request = NewCreateResultDownloadRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "CreateResultDownload")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateResultDownload require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateResultDownloadResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSparkAppRequest() (request *CreateSparkAppRequest) {
    request = &CreateSparkAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CreateSparkApp")
    
    
    return
}

func NewCreateSparkAppResponse() (response *CreateSparkAppResponse) {
    response = &CreateSparkAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSparkApp
// This API is used to create a Spark job.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_INVALIDAPPFILEFORMAT = "InvalidParameter.InvalidAppFileFormat"
//  INVALIDPARAMETER_INVALIDDRIVERSIZE = "InvalidParameter.InvalidDriverSize"
//  INVALIDPARAMETER_INVALIDEXECUTORSIZE = "InvalidParameter.InvalidExecutorSize"
//  INVALIDPARAMETER_INVALIDFILEPATHFORMAT = "InvalidParameter.InvalidFilePathFormat"
//  INVALIDPARAMETER_INVALIDROLEARN = "InvalidParameter.InvalidRoleArn"
//  INVALIDPARAMETER_SPARKJOBNOTUNIQUE = "InvalidParameter.SparkJobNotUnique"
//  INVALIDPARAMETER_SPARKJOBONLYSUPPORTSPARKBATCHENGINE = "InvalidParameter.SparkJobOnlySupportSparkBatchEngine"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
//  RESOURCENOTFOUND_SESSIONINSUFFICIENTRESOURCES = "ResourceNotFound.SessionInsufficientResources"
func (c *Client) CreateSparkApp(request *CreateSparkAppRequest) (response *CreateSparkAppResponse, err error) {
    return c.CreateSparkAppWithContext(context.Background(), request)
}

// CreateSparkApp
// This API is used to create a Spark job.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_INVALIDAPPFILEFORMAT = "InvalidParameter.InvalidAppFileFormat"
//  INVALIDPARAMETER_INVALIDDRIVERSIZE = "InvalidParameter.InvalidDriverSize"
//  INVALIDPARAMETER_INVALIDEXECUTORSIZE = "InvalidParameter.InvalidExecutorSize"
//  INVALIDPARAMETER_INVALIDFILEPATHFORMAT = "InvalidParameter.InvalidFilePathFormat"
//  INVALIDPARAMETER_INVALIDROLEARN = "InvalidParameter.InvalidRoleArn"
//  INVALIDPARAMETER_SPARKJOBNOTUNIQUE = "InvalidParameter.SparkJobNotUnique"
//  INVALIDPARAMETER_SPARKJOBONLYSUPPORTSPARKBATCHENGINE = "InvalidParameter.SparkJobOnlySupportSparkBatchEngine"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
//  RESOURCENOTFOUND_SESSIONINSUFFICIENTRESOURCES = "ResourceNotFound.SessionInsufficientResources"
func (c *Client) CreateSparkAppWithContext(ctx context.Context, request *CreateSparkAppRequest) (response *CreateSparkAppResponse, err error) {
    if request == nil {
        request = NewCreateSparkAppRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "CreateSparkApp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSparkApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSparkAppResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSparkAppTaskRequest() (request *CreateSparkAppTaskRequest) {
    request = &CreateSparkAppTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CreateSparkAppTask")
    
    
    return
}

func NewCreateSparkAppTaskResponse() (response *CreateSparkAppTaskResponse) {
    response = &CreateSparkAppTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSparkAppTask
// This API is used to start a Spark job.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOPERMISSIONTOUSETHEDATAENGINE = "FailedOperation.NoPermissionToUseTheDataEngine"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_IMAGEENGINETYPENOTMATCH = "InvalidParameter.ImageEngineTypeNotMatch"
//  INVALIDPARAMETER_IMAGEISPUBLICNOTMATCH = "InvalidParameter.ImageIsPublicNotMatch"
//  INVALIDPARAMETER_IMAGEPARAMETERSUBMITMETHODNOTMATCH = "InvalidParameter.ImageParameterSubmitMethodNotMatch"
//  INVALIDPARAMETER_IMAGEPARAMETERTYPENOTMATCH = "InvalidParameter.ImageParameterTypeNotMatch"
//  INVALIDPARAMETER_IMAGESESSIONPARAMETERSFORMATNOTJSON = "InvalidParameter.ImageSessionParametersFormatNotJson"
//  INVALIDPARAMETER_IMAGESTATENOTMATCH = "InvalidParameter.ImageStateNotMatch"
//  INVALIDPARAMETER_IMAGEUSERRECORDSTYPENOTMATCH = "InvalidParameter.ImageUserRecordsTypeNotMatch"
//  INVALIDPARAMETER_INVALIDROLEARN = "InvalidParameter.InvalidRoleArn"
//  INVALIDPARAMETER_INVALIDSPARKAPPPARAM = "InvalidParameter.InvalidSparkAppParam"
//  INVALIDPARAMETER_INVALIDSPARKCONFIGFORMAT = "InvalidParameter.InvalidSparkConfigFormat"
//  INVALIDPARAMETER_INVALIDTCRSPARKIMAGEFORMAT = "InvalidParameter.InvalidTcrSparkImageFormat"
//  INVALIDPARAMETER_INVALIDWHITELISTKEY = "InvalidParameter.InvalidWhiteListKey"
//  INVALIDPARAMETER_NUMBEROFSQLEXCEEDSTHELIMIT = "InvalidParameter.NumberOfSQLExceedsTheLimit"
//  INVALIDPARAMETER_PARAMETERBASE64DECODEFAILED = "InvalidParameter.ParameterBase64DecodeFailed"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  INVALIDPARAMETER_SQLBASE64DECODEFAIL = "InvalidParameter.SQLBase64DecodeFail"
//  INVALIDPARAMETER_SQLPARAMETERPREPROCESSINGFAILED = "InvalidParameter.SQLParameterPreprocessingFailed"
//  INVALIDPARAMETER_SPARKJOBNOTFOUND = "InvalidParameter.SparkJobNotFound"
//  INVALIDPARAMETER_SPARKJOBNOTUNIQUE = "InvalidParameter.SparkJobNotUnique"
//  INVALIDPARAMETER_SPARKJOBROLEARNNOTFOUND = "InvalidParameter.SparkJobRoleArnNotFound"
//  RESOURCEINSUFFICIENT_SPARKJOBINSUFFICIENTRESOURCES = "ResourceInsufficient.SparkJobInsufficientResources"
//  RESOURCENOTFOUND_DATAENGINECONFIGINSTANCENOTFOUND = "ResourceNotFound.DataEngineConfigInstanceNotFound"
//  RESOURCENOTFOUND_DATAENGINECONFIGINSTANCENOTUNIQUE = "ResourceNotFound.DataEngineConfigInstanceNotUnique"
//  RESOURCENOTFOUND_DATAENGINENOTACTIVITY = "ResourceNotFound.DataEngineNotActivity"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTUNIQUE = "ResourceNotFound.DataEngineNotUnique"
//  RESOURCENOTFOUND_IMAGEVERSIONNOTFOUND = "ResourceNotFound.ImageVersionNotFound"
//  RESOURCENOTFOUND_IMAGEVERSIONNOTUNIQUE = "ResourceNotFound.ImageVersionNotUnique"
//  RESOURCENOTFOUND_RESOURCEUSAGEOUTOFLIMIT = "ResourceNotFound.ResourceUsageOutOfLimit"
//  RESOURCENOTFOUND_SHUFFLEDIRNOTFOUND = "ResourceNotFound.ShuffleDirNotFound"
//  RESOURCENOTFOUND_WAREHOUSEDIRNOTFOUND = "ResourceNotFound.WarehouseDirNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_USECOMPUTINGENGINE = "UnauthorizedOperation.UseComputingEngine"
func (c *Client) CreateSparkAppTask(request *CreateSparkAppTaskRequest) (response *CreateSparkAppTaskResponse, err error) {
    return c.CreateSparkAppTaskWithContext(context.Background(), request)
}

// CreateSparkAppTask
// This API is used to start a Spark job.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOPERMISSIONTOUSETHEDATAENGINE = "FailedOperation.NoPermissionToUseTheDataEngine"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_IMAGEENGINETYPENOTMATCH = "InvalidParameter.ImageEngineTypeNotMatch"
//  INVALIDPARAMETER_IMAGEISPUBLICNOTMATCH = "InvalidParameter.ImageIsPublicNotMatch"
//  INVALIDPARAMETER_IMAGEPARAMETERSUBMITMETHODNOTMATCH = "InvalidParameter.ImageParameterSubmitMethodNotMatch"
//  INVALIDPARAMETER_IMAGEPARAMETERTYPENOTMATCH = "InvalidParameter.ImageParameterTypeNotMatch"
//  INVALIDPARAMETER_IMAGESESSIONPARAMETERSFORMATNOTJSON = "InvalidParameter.ImageSessionParametersFormatNotJson"
//  INVALIDPARAMETER_IMAGESTATENOTMATCH = "InvalidParameter.ImageStateNotMatch"
//  INVALIDPARAMETER_IMAGEUSERRECORDSTYPENOTMATCH = "InvalidParameter.ImageUserRecordsTypeNotMatch"
//  INVALIDPARAMETER_INVALIDROLEARN = "InvalidParameter.InvalidRoleArn"
//  INVALIDPARAMETER_INVALIDSPARKAPPPARAM = "InvalidParameter.InvalidSparkAppParam"
//  INVALIDPARAMETER_INVALIDSPARKCONFIGFORMAT = "InvalidParameter.InvalidSparkConfigFormat"
//  INVALIDPARAMETER_INVALIDTCRSPARKIMAGEFORMAT = "InvalidParameter.InvalidTcrSparkImageFormat"
//  INVALIDPARAMETER_INVALIDWHITELISTKEY = "InvalidParameter.InvalidWhiteListKey"
//  INVALIDPARAMETER_NUMBEROFSQLEXCEEDSTHELIMIT = "InvalidParameter.NumberOfSQLExceedsTheLimit"
//  INVALIDPARAMETER_PARAMETERBASE64DECODEFAILED = "InvalidParameter.ParameterBase64DecodeFailed"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  INVALIDPARAMETER_SQLBASE64DECODEFAIL = "InvalidParameter.SQLBase64DecodeFail"
//  INVALIDPARAMETER_SQLPARAMETERPREPROCESSINGFAILED = "InvalidParameter.SQLParameterPreprocessingFailed"
//  INVALIDPARAMETER_SPARKJOBNOTFOUND = "InvalidParameter.SparkJobNotFound"
//  INVALIDPARAMETER_SPARKJOBNOTUNIQUE = "InvalidParameter.SparkJobNotUnique"
//  INVALIDPARAMETER_SPARKJOBROLEARNNOTFOUND = "InvalidParameter.SparkJobRoleArnNotFound"
//  RESOURCEINSUFFICIENT_SPARKJOBINSUFFICIENTRESOURCES = "ResourceInsufficient.SparkJobInsufficientResources"
//  RESOURCENOTFOUND_DATAENGINECONFIGINSTANCENOTFOUND = "ResourceNotFound.DataEngineConfigInstanceNotFound"
//  RESOURCENOTFOUND_DATAENGINECONFIGINSTANCENOTUNIQUE = "ResourceNotFound.DataEngineConfigInstanceNotUnique"
//  RESOURCENOTFOUND_DATAENGINENOTACTIVITY = "ResourceNotFound.DataEngineNotActivity"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTUNIQUE = "ResourceNotFound.DataEngineNotUnique"
//  RESOURCENOTFOUND_IMAGEVERSIONNOTFOUND = "ResourceNotFound.ImageVersionNotFound"
//  RESOURCENOTFOUND_IMAGEVERSIONNOTUNIQUE = "ResourceNotFound.ImageVersionNotUnique"
//  RESOURCENOTFOUND_RESOURCEUSAGEOUTOFLIMIT = "ResourceNotFound.ResourceUsageOutOfLimit"
//  RESOURCENOTFOUND_SHUFFLEDIRNOTFOUND = "ResourceNotFound.ShuffleDirNotFound"
//  RESOURCENOTFOUND_WAREHOUSEDIRNOTFOUND = "ResourceNotFound.WarehouseDirNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_USECOMPUTINGENGINE = "UnauthorizedOperation.UseComputingEngine"
func (c *Client) CreateSparkAppTaskWithContext(ctx context.Context, request *CreateSparkAppTaskRequest) (response *CreateSparkAppTaskResponse, err error) {
    if request == nil {
        request = NewCreateSparkAppTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "CreateSparkAppTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSparkAppTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSparkAppTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSparkSessionBatchSQLRequest() (request *CreateSparkSessionBatchSQLRequest) {
    request = &CreateSparkSessionBatchSQLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CreateSparkSessionBatchSQL")
    
    
    return
}

func NewCreateSparkSessionBatchSQLResponse() (response *CreateSparkSessionBatchSQLResponse) {
    response = &CreateSparkSessionBatchSQLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSparkSessionBatchSQL
// This API is used to submit a Spark SQL batch task to the job engine.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOPERMISSIONTOUSETHEDATAENGINE = "FailedOperation.NoPermissionToUseTheDataEngine"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BATCHSQLCUSTOMKEYNOTUNIQUE = "InvalidParameter.BatchSQLCustomKeyNotUnique"
//  INVALIDPARAMETER_DATAENGINEONLYSUPPORTSPARKBATCH = "InvalidParameter.DataEngineOnlySupportSparkBatch"
//  INVALIDPARAMETER_IMAGEENGINETYPENOTMATCH = "InvalidParameter.ImageEngineTypeNotMatch"
//  INVALIDPARAMETER_IMAGEISPUBLICNOTMATCH = "InvalidParameter.ImageIsPublicNotMatch"
//  INVALIDPARAMETER_IMAGEPARAMETERSUBMITMETHODNOTMATCH = "InvalidParameter.ImageParameterSubmitMethodNotMatch"
//  INVALIDPARAMETER_IMAGEPARAMETERTYPENOTMATCH = "InvalidParameter.ImageParameterTypeNotMatch"
//  INVALIDPARAMETER_IMAGESTATENOTMATCH = "InvalidParameter.ImageStateNotMatch"
//  INVALIDPARAMETER_INVALIDDRIVERSIZE = "InvalidParameter.InvalidDriverSize"
//  INVALIDPARAMETER_INVALIDDYNAMICALLOCATIONMAXEXECUTORS = "InvalidParameter.InvalidDynamicAllocationMaxExecutors"
//  INVALIDPARAMETER_INVALIDEXECUTORSIZE = "InvalidParameter.InvalidExecutorSize"
//  INVALIDPARAMETER_INVALIDFILECOMPRESSIONFORMAT = "InvalidParameter.InvalidFileCompressionFormat"
//  INVALIDPARAMETER_INVALIDFILEPATHFORMAT = "InvalidParameter.InvalidFilePathFormat"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  INVALIDPARAMETER_INVALIDSESSIONKINDTYPE = "InvalidParameter.InvalidSessionKindType"
//  INVALIDPARAMETER_INVALIDSTATEMENTKINDTYPE = "InvalidParameter.InvalidStatementKindType"
//  INVALIDPARAMETER_INVALIDWHITELISTKEY = "InvalidParameter.InvalidWhiteListKey"
//  INVALIDPARAMETER_NUMBEROFSQLEXCEEDSTHELIMIT = "InvalidParameter.NumberOfSQLExceedsTheLimit"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  INVALIDPARAMETER_SQLBASE64DECODEFAIL = "InvalidParameter.SQLBase64DecodeFail"
//  INVALIDPARAMETER_SQLPARAMETERPREPROCESSINGFAILED = "InvalidParameter.SQLParameterPreprocessingFailed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTRUNNING = "ResourceNotFound.DataEngineNotRunning"
//  RESOURCENOTFOUND_DATAENGINENOTUNIQUE = "ResourceNotFound.DataEngineNotUnique"
//  RESOURCENOTFOUND_IMAGESESSIONCONFIGNOTFOUND = "ResourceNotFound.ImageSessionConfigNotFound"
//  RESOURCENOTFOUND_IMAGESESSIONCONFIGNOTUNIQUE = "ResourceNotFound.ImageSessionConfigNotUnique"
//  RESOURCENOTFOUND_IMAGEVERSIONNOTFOUND = "ResourceNotFound.ImageVersionNotFound"
//  RESOURCENOTFOUND_RESULTSAVEPATHNOTFOUND = "ResourceNotFound.ResultSavePathNotFound"
//  RESOURCENOTFOUND_ROLEARNRESOURCENOTFOUND = "ResourceNotFound.RoleArnResourceNotFound"
//  RESOURCENOTFOUND_SESSIONINSUFFICIENTRESOURCES = "ResourceNotFound.SessionInsufficientResources"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  RESOURCENOTFOUND_SESSIONSTATEDEAD = "ResourceNotFound.SessionStateDead"
//  RESOURCENOTFOUND_SHUFFLEDIRNOTFOUND = "ResourceNotFound.ShuffleDirNotFound"
//  RESOURCENOTFOUND_WAREHOUSEDIRNOTFOUND = "ResourceNotFound.WarehouseDirNotFound"
func (c *Client) CreateSparkSessionBatchSQL(request *CreateSparkSessionBatchSQLRequest) (response *CreateSparkSessionBatchSQLResponse, err error) {
    return c.CreateSparkSessionBatchSQLWithContext(context.Background(), request)
}

// CreateSparkSessionBatchSQL
// This API is used to submit a Spark SQL batch task to the job engine.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOPERMISSIONTOUSETHEDATAENGINE = "FailedOperation.NoPermissionToUseTheDataEngine"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BATCHSQLCUSTOMKEYNOTUNIQUE = "InvalidParameter.BatchSQLCustomKeyNotUnique"
//  INVALIDPARAMETER_DATAENGINEONLYSUPPORTSPARKBATCH = "InvalidParameter.DataEngineOnlySupportSparkBatch"
//  INVALIDPARAMETER_IMAGEENGINETYPENOTMATCH = "InvalidParameter.ImageEngineTypeNotMatch"
//  INVALIDPARAMETER_IMAGEISPUBLICNOTMATCH = "InvalidParameter.ImageIsPublicNotMatch"
//  INVALIDPARAMETER_IMAGEPARAMETERSUBMITMETHODNOTMATCH = "InvalidParameter.ImageParameterSubmitMethodNotMatch"
//  INVALIDPARAMETER_IMAGEPARAMETERTYPENOTMATCH = "InvalidParameter.ImageParameterTypeNotMatch"
//  INVALIDPARAMETER_IMAGESTATENOTMATCH = "InvalidParameter.ImageStateNotMatch"
//  INVALIDPARAMETER_INVALIDDRIVERSIZE = "InvalidParameter.InvalidDriverSize"
//  INVALIDPARAMETER_INVALIDDYNAMICALLOCATIONMAXEXECUTORS = "InvalidParameter.InvalidDynamicAllocationMaxExecutors"
//  INVALIDPARAMETER_INVALIDEXECUTORSIZE = "InvalidParameter.InvalidExecutorSize"
//  INVALIDPARAMETER_INVALIDFILECOMPRESSIONFORMAT = "InvalidParameter.InvalidFileCompressionFormat"
//  INVALIDPARAMETER_INVALIDFILEPATHFORMAT = "InvalidParameter.InvalidFilePathFormat"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  INVALIDPARAMETER_INVALIDSESSIONKINDTYPE = "InvalidParameter.InvalidSessionKindType"
//  INVALIDPARAMETER_INVALIDSTATEMENTKINDTYPE = "InvalidParameter.InvalidStatementKindType"
//  INVALIDPARAMETER_INVALIDWHITELISTKEY = "InvalidParameter.InvalidWhiteListKey"
//  INVALIDPARAMETER_NUMBEROFSQLEXCEEDSTHELIMIT = "InvalidParameter.NumberOfSQLExceedsTheLimit"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  INVALIDPARAMETER_SQLBASE64DECODEFAIL = "InvalidParameter.SQLBase64DecodeFail"
//  INVALIDPARAMETER_SQLPARAMETERPREPROCESSINGFAILED = "InvalidParameter.SQLParameterPreprocessingFailed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTRUNNING = "ResourceNotFound.DataEngineNotRunning"
//  RESOURCENOTFOUND_DATAENGINENOTUNIQUE = "ResourceNotFound.DataEngineNotUnique"
//  RESOURCENOTFOUND_IMAGESESSIONCONFIGNOTFOUND = "ResourceNotFound.ImageSessionConfigNotFound"
//  RESOURCENOTFOUND_IMAGESESSIONCONFIGNOTUNIQUE = "ResourceNotFound.ImageSessionConfigNotUnique"
//  RESOURCENOTFOUND_IMAGEVERSIONNOTFOUND = "ResourceNotFound.ImageVersionNotFound"
//  RESOURCENOTFOUND_RESULTSAVEPATHNOTFOUND = "ResourceNotFound.ResultSavePathNotFound"
//  RESOURCENOTFOUND_ROLEARNRESOURCENOTFOUND = "ResourceNotFound.RoleArnResourceNotFound"
//  RESOURCENOTFOUND_SESSIONINSUFFICIENTRESOURCES = "ResourceNotFound.SessionInsufficientResources"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  RESOURCENOTFOUND_SESSIONSTATEDEAD = "ResourceNotFound.SessionStateDead"
//  RESOURCENOTFOUND_SHUFFLEDIRNOTFOUND = "ResourceNotFound.ShuffleDirNotFound"
//  RESOURCENOTFOUND_WAREHOUSEDIRNOTFOUND = "ResourceNotFound.WarehouseDirNotFound"
func (c *Client) CreateSparkSessionBatchSQLWithContext(ctx context.Context, request *CreateSparkSessionBatchSQLRequest) (response *CreateSparkSessionBatchSQLResponse, err error) {
    if request == nil {
        request = NewCreateSparkSessionBatchSQLRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "CreateSparkSessionBatchSQL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSparkSessionBatchSQL require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSparkSessionBatchSQLResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStoreLocationRequest() (request *CreateStoreLocationRequest) {
    request = &CreateStoreLocationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CreateStoreLocation")
    
    
    return
}

func NewCreateStoreLocationResponse() (response *CreateStoreLocationResponse) {
    response = &CreateStoreLocationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateStoreLocation
// This API is used to add or overwrite the storage location of results.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSTORELOCATION = "InvalidParameter.InvalidStoreLocation"
func (c *Client) CreateStoreLocation(request *CreateStoreLocationRequest) (response *CreateStoreLocationResponse, err error) {
    return c.CreateStoreLocationWithContext(context.Background(), request)
}

// CreateStoreLocation
// This API is used to add or overwrite the storage location of results.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSTORELOCATION = "InvalidParameter.InvalidStoreLocation"
func (c *Client) CreateStoreLocationWithContext(ctx context.Context, request *CreateStoreLocationRequest) (response *CreateStoreLocationResponse, err error) {
    if request == nil {
        request = NewCreateStoreLocationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "CreateStoreLocation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateStoreLocation require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateStoreLocationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTaskRequest() (request *CreateTaskRequest) {
    request = &CreateTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CreateTask")
    
    
    return
}

func NewCreateTaskResponse() (response *CreateTaskResponse) {
    response = &CreateTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTask
// This API is used to create and execute a SQL task. (`CreateTasks` is recommended.)
//
// error code that may be returned:
//  FAILEDOPERATION_NOPERMISSIONTOUSETHEDATAENGINE = "FailedOperation.NoPermissionToUseTheDataEngine"
//  FAILEDOPERATION_SQLTASKPARSEFAILED = "FailedOperation.SQLTaskParseFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAENGINEONLYSUPPORTSQL = "InvalidParameter.DataEngineOnlySupportSQL"
//  INVALIDPARAMETER_IMAGEENGINETYPENOTMATCH = "InvalidParameter.ImageEngineTypeNotMatch"
//  INVALIDPARAMETER_IMAGEISPUBLICNOTMATCH = "InvalidParameter.ImageIsPublicNotMatch"
//  INVALIDPARAMETER_IMAGEPARAMETERNOTFOUND = "InvalidParameter.ImageParameterNotFound"
//  INVALIDPARAMETER_IMAGEPARAMETERSUBMITMETHODNOTMATCH = "InvalidParameter.ImageParameterSubmitMethodNotMatch"
//  INVALIDPARAMETER_IMAGEPARAMETERTYPENOTMATCH = "InvalidParameter.ImageParameterTypeNotMatch"
//  INVALIDPARAMETER_INVALIDCONFIGKEYNOTFOUND = "InvalidParameter.InvalidConfigKeyNotFound"
//  INVALIDPARAMETER_INVALIDCONFIGVALUELENGTHOUTLIMIT = "InvalidParameter.InvalidConfigValueLengthOutLimit"
//  INVALIDPARAMETER_INVALIDCONFIGVALUEREGEXPNOTMATCH = "InvalidParameter.InvalidConfigValueRegexpNotMatch"
//  INVALIDPARAMETER_INVALIDDATAENGINENAME = "InvalidParameter.InvalidDataEngineName"
//  INVALIDPARAMETER_INVALIDFAILURETOLERANCE = "InvalidParameter.InvalidFailureTolerance"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  INVALIDPARAMETER_INVALIDSQLCONFIGSQL = "InvalidParameter.InvalidSQLConfigSQL"
//  INVALIDPARAMETER_INVALIDSQLNUM = "InvalidParameter.InvalidSQLNum"
//  INVALIDPARAMETER_INVALIDSTORELOCATION = "InvalidParameter.InvalidStoreLocation"
//  INVALIDPARAMETER_INVALIDTASKTYPE = "InvalidParameter.InvalidTaskType"
//  INVALIDPARAMETER_INVALIDWHITELISTKEY = "InvalidParameter.InvalidWhiteListKey"
//  INVALIDPARAMETER_PARAMETERBASE64DECODEFAILED = "InvalidParameter.ParameterBase64DecodeFailed"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  INVALIDPARAMETER_SQLBASE64DECODEFAIL = "InvalidParameter.SQLBase64DecodeFail"
//  INVALIDPARAMETER_SQLPARAMETERPREPROCESSINGFAILED = "InvalidParameter.SQLParameterPreprocessingFailed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTRUNNING = "ResourceNotFound.DataEngineNotRunning"
//  RESOURCENOTFOUND_DATAENGINENOTUNIQUE = "ResourceNotFound.DataEngineNotUnique"
//  RESOURCENOTFOUND_DEFAULTDATAENGINENOTFOUND = "ResourceNotFound.DefaultDataEngineNotFound"
//  RESOURCENOTFOUND_RESULTOUTPUTPATHNOTFOUND = "ResourceNotFound.ResultOutputPathNotFound"
//  RESOURCEUNAVAILABLE_BALANCEINSUFFICIENT = "ResourceUnavailable.BalanceInsufficient"
//  UNAUTHORIZEDOPERATION_USECOMPUTINGENGINE = "UnauthorizedOperation.UseComputingEngine"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDFILETYPE = "UnsupportedOperation.UnsupportedFileType"
func (c *Client) CreateTask(request *CreateTaskRequest) (response *CreateTaskResponse, err error) {
    return c.CreateTaskWithContext(context.Background(), request)
}

// CreateTask
// This API is used to create and execute a SQL task. (`CreateTasks` is recommended.)
//
// error code that may be returned:
//  FAILEDOPERATION_NOPERMISSIONTOUSETHEDATAENGINE = "FailedOperation.NoPermissionToUseTheDataEngine"
//  FAILEDOPERATION_SQLTASKPARSEFAILED = "FailedOperation.SQLTaskParseFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAENGINEONLYSUPPORTSQL = "InvalidParameter.DataEngineOnlySupportSQL"
//  INVALIDPARAMETER_IMAGEENGINETYPENOTMATCH = "InvalidParameter.ImageEngineTypeNotMatch"
//  INVALIDPARAMETER_IMAGEISPUBLICNOTMATCH = "InvalidParameter.ImageIsPublicNotMatch"
//  INVALIDPARAMETER_IMAGEPARAMETERNOTFOUND = "InvalidParameter.ImageParameterNotFound"
//  INVALIDPARAMETER_IMAGEPARAMETERSUBMITMETHODNOTMATCH = "InvalidParameter.ImageParameterSubmitMethodNotMatch"
//  INVALIDPARAMETER_IMAGEPARAMETERTYPENOTMATCH = "InvalidParameter.ImageParameterTypeNotMatch"
//  INVALIDPARAMETER_INVALIDCONFIGKEYNOTFOUND = "InvalidParameter.InvalidConfigKeyNotFound"
//  INVALIDPARAMETER_INVALIDCONFIGVALUELENGTHOUTLIMIT = "InvalidParameter.InvalidConfigValueLengthOutLimit"
//  INVALIDPARAMETER_INVALIDCONFIGVALUEREGEXPNOTMATCH = "InvalidParameter.InvalidConfigValueRegexpNotMatch"
//  INVALIDPARAMETER_INVALIDDATAENGINENAME = "InvalidParameter.InvalidDataEngineName"
//  INVALIDPARAMETER_INVALIDFAILURETOLERANCE = "InvalidParameter.InvalidFailureTolerance"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  INVALIDPARAMETER_INVALIDSQLCONFIGSQL = "InvalidParameter.InvalidSQLConfigSQL"
//  INVALIDPARAMETER_INVALIDSQLNUM = "InvalidParameter.InvalidSQLNum"
//  INVALIDPARAMETER_INVALIDSTORELOCATION = "InvalidParameter.InvalidStoreLocation"
//  INVALIDPARAMETER_INVALIDTASKTYPE = "InvalidParameter.InvalidTaskType"
//  INVALIDPARAMETER_INVALIDWHITELISTKEY = "InvalidParameter.InvalidWhiteListKey"
//  INVALIDPARAMETER_PARAMETERBASE64DECODEFAILED = "InvalidParameter.ParameterBase64DecodeFailed"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  INVALIDPARAMETER_SQLBASE64DECODEFAIL = "InvalidParameter.SQLBase64DecodeFail"
//  INVALIDPARAMETER_SQLPARAMETERPREPROCESSINGFAILED = "InvalidParameter.SQLParameterPreprocessingFailed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTRUNNING = "ResourceNotFound.DataEngineNotRunning"
//  RESOURCENOTFOUND_DATAENGINENOTUNIQUE = "ResourceNotFound.DataEngineNotUnique"
//  RESOURCENOTFOUND_DEFAULTDATAENGINENOTFOUND = "ResourceNotFound.DefaultDataEngineNotFound"
//  RESOURCENOTFOUND_RESULTOUTPUTPATHNOTFOUND = "ResourceNotFound.ResultOutputPathNotFound"
//  RESOURCEUNAVAILABLE_BALANCEINSUFFICIENT = "ResourceUnavailable.BalanceInsufficient"
//  UNAUTHORIZEDOPERATION_USECOMPUTINGENGINE = "UnauthorizedOperation.UseComputingEngine"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDFILETYPE = "UnsupportedOperation.UnsupportedFileType"
func (c *Client) CreateTaskWithContext(ctx context.Context, request *CreateTaskRequest) (response *CreateTaskResponse, err error) {
    if request == nil {
        request = NewCreateTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "CreateTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTasksRequest() (request *CreateTasksRequest) {
    request = &CreateTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CreateTasks")
    
    
    return
}

func NewCreateTasksResponse() (response *CreateTasksResponse) {
    response = &CreateTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTasks
// This API is used to create and execute SQL tasks in batches.
//
// error code that may be returned:
//  FAILEDOPERATION_NOPERMISSIONTOUSETHEDATAENGINE = "FailedOperation.NoPermissionToUseTheDataEngine"
//  FAILEDOPERATION_SQLTASKPARSEFAILED = "FailedOperation.SQLTaskParseFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAENGINEONLYSUPPORTSQL = "InvalidParameter.DataEngineOnlySupportSQL"
//  INVALIDPARAMETER_IMAGEENGINETYPENOTMATCH = "InvalidParameter.ImageEngineTypeNotMatch"
//  INVALIDPARAMETER_IMAGEISPUBLICNOTMATCH = "InvalidParameter.ImageIsPublicNotMatch"
//  INVALIDPARAMETER_IMAGEPARAMETERNOTFOUND = "InvalidParameter.ImageParameterNotFound"
//  INVALIDPARAMETER_IMAGEPARAMETERSUBMITMETHODNOTMATCH = "InvalidParameter.ImageParameterSubmitMethodNotMatch"
//  INVALIDPARAMETER_IMAGEPARAMETERTYPENOTMATCH = "InvalidParameter.ImageParameterTypeNotMatch"
//  INVALIDPARAMETER_INVALIDCONFIGKEYNOTFOUND = "InvalidParameter.InvalidConfigKeyNotFound"
//  INVALIDPARAMETER_INVALIDCONFIGVALUELENGTHOUTLIMIT = "InvalidParameter.InvalidConfigValueLengthOutLimit"
//  INVALIDPARAMETER_INVALIDCONFIGVALUEREGEXPNOTMATCH = "InvalidParameter.InvalidConfigValueRegexpNotMatch"
//  INVALIDPARAMETER_INVALIDDATAENGINENAME = "InvalidParameter.InvalidDataEngineName"
//  INVALIDPARAMETER_INVALIDFAILURETOLERANCE = "InvalidParameter.InvalidFailureTolerance"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  INVALIDPARAMETER_INVALIDSQLNUM = "InvalidParameter.InvalidSQLNum"
//  INVALIDPARAMETER_INVALIDSTORELOCATION = "InvalidParameter.InvalidStoreLocation"
//  INVALIDPARAMETER_INVALIDTASKTYPE = "InvalidParameter.InvalidTaskType"
//  INVALIDPARAMETER_INVALIDWHITELISTKEY = "InvalidParameter.InvalidWhiteListKey"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  INVALIDPARAMETER_SQLBASE64DECODEFAIL = "InvalidParameter.SQLBase64DecodeFail"
//  INVALIDPARAMETER_SQLPARAMETERPREPROCESSINGFAILED = "InvalidParameter.SQLParameterPreprocessingFailed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTRUNNING = "ResourceNotFound.DataEngineNotRunning"
//  RESOURCENOTFOUND_DATAENGINENOTUNIQUE = "ResourceNotFound.DataEngineNotUnique"
//  RESOURCENOTFOUND_DEFAULTDATAENGINENOTFOUND = "ResourceNotFound.DefaultDataEngineNotFound"
//  RESOURCENOTFOUND_RESULTOUTPUTPATHNOTFOUND = "ResourceNotFound.ResultOutputPathNotFound"
//  RESOURCEUNAVAILABLE_BALANCEINSUFFICIENT = "ResourceUnavailable.BalanceInsufficient"
//  UNAUTHORIZEDOPERATION_USECOMPUTINGENGINE = "UnauthorizedOperation.UseComputingEngine"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDFILETYPE = "UnsupportedOperation.UnsupportedFileType"
func (c *Client) CreateTasks(request *CreateTasksRequest) (response *CreateTasksResponse, err error) {
    return c.CreateTasksWithContext(context.Background(), request)
}

// CreateTasks
// This API is used to create and execute SQL tasks in batches.
//
// error code that may be returned:
//  FAILEDOPERATION_NOPERMISSIONTOUSETHEDATAENGINE = "FailedOperation.NoPermissionToUseTheDataEngine"
//  FAILEDOPERATION_SQLTASKPARSEFAILED = "FailedOperation.SQLTaskParseFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAENGINEONLYSUPPORTSQL = "InvalidParameter.DataEngineOnlySupportSQL"
//  INVALIDPARAMETER_IMAGEENGINETYPENOTMATCH = "InvalidParameter.ImageEngineTypeNotMatch"
//  INVALIDPARAMETER_IMAGEISPUBLICNOTMATCH = "InvalidParameter.ImageIsPublicNotMatch"
//  INVALIDPARAMETER_IMAGEPARAMETERNOTFOUND = "InvalidParameter.ImageParameterNotFound"
//  INVALIDPARAMETER_IMAGEPARAMETERSUBMITMETHODNOTMATCH = "InvalidParameter.ImageParameterSubmitMethodNotMatch"
//  INVALIDPARAMETER_IMAGEPARAMETERTYPENOTMATCH = "InvalidParameter.ImageParameterTypeNotMatch"
//  INVALIDPARAMETER_INVALIDCONFIGKEYNOTFOUND = "InvalidParameter.InvalidConfigKeyNotFound"
//  INVALIDPARAMETER_INVALIDCONFIGVALUELENGTHOUTLIMIT = "InvalidParameter.InvalidConfigValueLengthOutLimit"
//  INVALIDPARAMETER_INVALIDCONFIGVALUEREGEXPNOTMATCH = "InvalidParameter.InvalidConfigValueRegexpNotMatch"
//  INVALIDPARAMETER_INVALIDDATAENGINENAME = "InvalidParameter.InvalidDataEngineName"
//  INVALIDPARAMETER_INVALIDFAILURETOLERANCE = "InvalidParameter.InvalidFailureTolerance"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  INVALIDPARAMETER_INVALIDSQLNUM = "InvalidParameter.InvalidSQLNum"
//  INVALIDPARAMETER_INVALIDSTORELOCATION = "InvalidParameter.InvalidStoreLocation"
//  INVALIDPARAMETER_INVALIDTASKTYPE = "InvalidParameter.InvalidTaskType"
//  INVALIDPARAMETER_INVALIDWHITELISTKEY = "InvalidParameter.InvalidWhiteListKey"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  INVALIDPARAMETER_SQLBASE64DECODEFAIL = "InvalidParameter.SQLBase64DecodeFail"
//  INVALIDPARAMETER_SQLPARAMETERPREPROCESSINGFAILED = "InvalidParameter.SQLParameterPreprocessingFailed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTRUNNING = "ResourceNotFound.DataEngineNotRunning"
//  RESOURCENOTFOUND_DATAENGINENOTUNIQUE = "ResourceNotFound.DataEngineNotUnique"
//  RESOURCENOTFOUND_DEFAULTDATAENGINENOTFOUND = "ResourceNotFound.DefaultDataEngineNotFound"
//  RESOURCENOTFOUND_RESULTOUTPUTPATHNOTFOUND = "ResourceNotFound.ResultOutputPathNotFound"
//  RESOURCEUNAVAILABLE_BALANCEINSUFFICIENT = "ResourceUnavailable.BalanceInsufficient"
//  UNAUTHORIZEDOPERATION_USECOMPUTINGENGINE = "UnauthorizedOperation.UseComputingEngine"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDFILETYPE = "UnsupportedOperation.UnsupportedFileType"
func (c *Client) CreateTasksWithContext(ctx context.Context, request *CreateTasksRequest) (response *CreateTasksResponse, err error) {
    if request == nil {
        request = NewCreateTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "CreateTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTasksResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserRequest() (request *CreateUserRequest) {
    request = &CreateUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CreateUser")
    
    
    return
}

func NewCreateUserResponse() (response *CreateUserResponse) {
    response = &CreateUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateUser
// This API is used to create users.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GRANTPOLICYFAILED = "FailedOperation.GrantPolicyFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_DUPLICATEUSERNAME = "InvalidParameter.DuplicateUserName"
//  INVALIDPARAMETER_INVALIDACCESSPOLICY = "InvalidParameter.InvalidAccessPolicy"
//  INVALIDPARAMETER_INVALIDDESCRIPTION = "InvalidParameter.InvalidDescription"
//  INVALIDPARAMETER_INVALIDUSERALIAS = "InvalidParameter.InvalidUserAlias"
//  INVALIDPARAMETER_INVALIDUSERNAME = "InvalidParameter.InvalidUserName"
//  INVALIDPARAMETER_INVALIDUSERTYPE = "InvalidParameter.InvalidUserType"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  RESOURCESSOLDOUT_UNAUTHORIZEDGRANTPOLICY = "ResourcesSoldOut.UnauthorizedGrantPolicy"
//  RESOURCESSOLDOUT_UNAUTHORIZEDOPERATION = "ResourcesSoldOut.UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_GRANTPOLICY = "UnauthorizedOperation.GrantPolicy"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
func (c *Client) CreateUser(request *CreateUserRequest) (response *CreateUserResponse, err error) {
    return c.CreateUserWithContext(context.Background(), request)
}

// CreateUser
// This API is used to create users.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GRANTPOLICYFAILED = "FailedOperation.GrantPolicyFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_DUPLICATEUSERNAME = "InvalidParameter.DuplicateUserName"
//  INVALIDPARAMETER_INVALIDACCESSPOLICY = "InvalidParameter.InvalidAccessPolicy"
//  INVALIDPARAMETER_INVALIDDESCRIPTION = "InvalidParameter.InvalidDescription"
//  INVALIDPARAMETER_INVALIDUSERALIAS = "InvalidParameter.InvalidUserAlias"
//  INVALIDPARAMETER_INVALIDUSERNAME = "InvalidParameter.InvalidUserName"
//  INVALIDPARAMETER_INVALIDUSERTYPE = "InvalidParameter.InvalidUserType"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  RESOURCESSOLDOUT_UNAUTHORIZEDGRANTPOLICY = "ResourcesSoldOut.UnauthorizedGrantPolicy"
//  RESOURCESSOLDOUT_UNAUTHORIZEDOPERATION = "ResourcesSoldOut.UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_GRANTPOLICY = "UnauthorizedOperation.GrantPolicy"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
func (c *Client) CreateUserWithContext(ctx context.Context, request *CreateUserRequest) (response *CreateUserResponse, err error) {
    if request == nil {
        request = NewCreateUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "CreateUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWorkGroupRequest() (request *CreateWorkGroupRequest) {
    request = &CreateWorkGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "CreateWorkGroup")
    
    
    return
}

func NewCreateWorkGroupResponse() (response *CreateWorkGroupResponse) {
    response = &CreateWorkGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateWorkGroup
// This API is used to create working groups.
//
// error code that may be returned:
//  FAILEDOPERATION_GRANTPOLICYFAILED = "FailedOperation.GrantPolicyFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_DUPLICATEGROUPNAME = "InvalidParameter.DuplicateGroupName"
//  INVALIDPARAMETER_INVALIDACCESSPOLICY = "InvalidParameter.InvalidAccessPolicy"
//  INVALIDPARAMETER_INVALIDDESCRIPTION = "InvalidParameter.InvalidDescription"
//  INVALIDPARAMETER_INVALIDWORKGROUPNAME = "InvalidParameter.InvalidWorkGroupName"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  RESOURCESSOLDOUT_UNAUTHORIZEDGRANTPOLICY = "ResourcesSoldOut.UnauthorizedGrantPolicy"
//  RESOURCESSOLDOUT_UNAUTHORIZEDOPERATION = "ResourcesSoldOut.UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CREATEWORKGROUP = "UnauthorizedOperation.CreateWorkgroup"
//  UNAUTHORIZEDOPERATION_GRANTPOLICY = "UnauthorizedOperation.GrantPolicy"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
func (c *Client) CreateWorkGroup(request *CreateWorkGroupRequest) (response *CreateWorkGroupResponse, err error) {
    return c.CreateWorkGroupWithContext(context.Background(), request)
}

// CreateWorkGroup
// This API is used to create working groups.
//
// error code that may be returned:
//  FAILEDOPERATION_GRANTPOLICYFAILED = "FailedOperation.GrantPolicyFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_DUPLICATEGROUPNAME = "InvalidParameter.DuplicateGroupName"
//  INVALIDPARAMETER_INVALIDACCESSPOLICY = "InvalidParameter.InvalidAccessPolicy"
//  INVALIDPARAMETER_INVALIDDESCRIPTION = "InvalidParameter.InvalidDescription"
//  INVALIDPARAMETER_INVALIDWORKGROUPNAME = "InvalidParameter.InvalidWorkGroupName"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  RESOURCESSOLDOUT_UNAUTHORIZEDGRANTPOLICY = "ResourcesSoldOut.UnauthorizedGrantPolicy"
//  RESOURCESSOLDOUT_UNAUTHORIZEDOPERATION = "ResourcesSoldOut.UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CREATEWORKGROUP = "UnauthorizedOperation.CreateWorkgroup"
//  UNAUTHORIZEDOPERATION_GRANTPOLICY = "UnauthorizedOperation.GrantPolicy"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
func (c *Client) CreateWorkGroupWithContext(ctx context.Context, request *CreateWorkGroupRequest) (response *CreateWorkGroupResponse, err error) {
    if request == nil {
        request = NewCreateWorkGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "CreateWorkGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWorkGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWorkGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCHDFSBindingProductRequest() (request *DeleteCHDFSBindingProductRequest) {
    request = &DeleteCHDFSBindingProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DeleteCHDFSBindingProduct")
    
    
    return
}

func NewDeleteCHDFSBindingProductResponse() (response *DeleteCHDFSBindingProductResponse) {
    response = &DeleteCHDFSBindingProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCHDFSBindingProduct
// This API is used to delete the binding relationship between metadata acceleration buckets and products.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) DeleteCHDFSBindingProduct(request *DeleteCHDFSBindingProductRequest) (response *DeleteCHDFSBindingProductResponse, err error) {
    return c.DeleteCHDFSBindingProductWithContext(context.Background(), request)
}

// DeleteCHDFSBindingProduct
// This API is used to delete the binding relationship between metadata acceleration buckets and products.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) DeleteCHDFSBindingProductWithContext(ctx context.Context, request *DeleteCHDFSBindingProductRequest) (response *DeleteCHDFSBindingProductResponse, err error) {
    if request == nil {
        request = NewDeleteCHDFSBindingProductRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DeleteCHDFSBindingProduct")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCHDFSBindingProduct require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCHDFSBindingProductResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDataEngineRequest() (request *DeleteDataEngineRequest) {
    request = &DeleteDataEngineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DeleteDataEngine")
    
    
    return
}

func NewDeleteDataEngineResponse() (response *DeleteDataEngineResponse) {
    response = &DeleteDataEngineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDataEngine
// This API is used to delete the data engine.
//
// error code that may be returned:
//  FAILEDOPERATION_ABNORMALORDERSTATUS = "FailedOperation.AbnormalOrderStatus"
//  FAILEDOPERATION_ANOTHERPROCESSRUNNING = "FailedOperation.AnotherProcessRunning"
//  FAILEDOPERATION_ANOTHERREQUESTPROCESSING = "FailedOperation.AnotherRequestProcessing"
//  FAILEDOPERATION_BALANCENOTENOUGH = "FailedOperation.BalanceNotEnough"
//  FAILEDOPERATION_BILLINGSYSTEMERROR = "FailedOperation.BillingSystemError"
//  FAILEDOPERATION_DELETEDATAENGINEFAILED = "FailedOperation.DeleteDataEngineFailed"
//  FAILEDOPERATION_DELIVERGOODSFAILED = "FailedOperation.DeliverGoodsFailed"
//  FAILEDOPERATION_FEEDEDUCTIONFAILED = "FailedOperation.FeeDeductionFailed"
//  FAILEDOPERATION_GETPRODUCTINFORMATIONFAILED = "FailedOperation.GetProductInformationFailed"
//  FAILEDOPERATION_INQUIREPRICEFAILED = "FailedOperation.InquirePriceFailed"
//  FAILEDOPERATION_MODIFYINSTANCEFAILED = "FailedOperation.ModifyInstanceFailed"
//  FAILEDOPERATION_NOREALNAMEAUTHENTICATION = "FailedOperation.NoRealNameAuthentication"
//  FAILEDOPERATION_NUMBEREXCEEDLIMIT = "FailedOperation.NumberExceedLimit"
//  FAILEDOPERATION_PARAMETERVALIDATIONFAILED = "FailedOperation.ParameterValidationFailed"
//  FAILEDOPERATION_REFUNDDEPOSITFAILED = "FailedOperation.RefundDepositFailed"
//  FAILEDOPERATION_TOOMANYRESOURCES = "FailedOperation.TooManyResources"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INSTANCEINPROCESS = "InvalidParameter.InstanceInProcess"
//  INVALIDPARAMETER_INVALIDDATAENGINENAME = "InvalidParameter.InvalidDataEngineName"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTUNIQUE = "ResourceNotFound.DataEngineNotUnique"
//  UNAUTHORIZEDOPERATION_DELETECOMPUTINGENGINE = "UnauthorizedOperation.DeleteComputingEngine"
//  UNAUTHORIZEDOPERATION_NOPAYMENTAUTHORITY = "UnauthorizedOperation.NoPaymentAuthority"
func (c *Client) DeleteDataEngine(request *DeleteDataEngineRequest) (response *DeleteDataEngineResponse, err error) {
    return c.DeleteDataEngineWithContext(context.Background(), request)
}

// DeleteDataEngine
// This API is used to delete the data engine.
//
// error code that may be returned:
//  FAILEDOPERATION_ABNORMALORDERSTATUS = "FailedOperation.AbnormalOrderStatus"
//  FAILEDOPERATION_ANOTHERPROCESSRUNNING = "FailedOperation.AnotherProcessRunning"
//  FAILEDOPERATION_ANOTHERREQUESTPROCESSING = "FailedOperation.AnotherRequestProcessing"
//  FAILEDOPERATION_BALANCENOTENOUGH = "FailedOperation.BalanceNotEnough"
//  FAILEDOPERATION_BILLINGSYSTEMERROR = "FailedOperation.BillingSystemError"
//  FAILEDOPERATION_DELETEDATAENGINEFAILED = "FailedOperation.DeleteDataEngineFailed"
//  FAILEDOPERATION_DELIVERGOODSFAILED = "FailedOperation.DeliverGoodsFailed"
//  FAILEDOPERATION_FEEDEDUCTIONFAILED = "FailedOperation.FeeDeductionFailed"
//  FAILEDOPERATION_GETPRODUCTINFORMATIONFAILED = "FailedOperation.GetProductInformationFailed"
//  FAILEDOPERATION_INQUIREPRICEFAILED = "FailedOperation.InquirePriceFailed"
//  FAILEDOPERATION_MODIFYINSTANCEFAILED = "FailedOperation.ModifyInstanceFailed"
//  FAILEDOPERATION_NOREALNAMEAUTHENTICATION = "FailedOperation.NoRealNameAuthentication"
//  FAILEDOPERATION_NUMBEREXCEEDLIMIT = "FailedOperation.NumberExceedLimit"
//  FAILEDOPERATION_PARAMETERVALIDATIONFAILED = "FailedOperation.ParameterValidationFailed"
//  FAILEDOPERATION_REFUNDDEPOSITFAILED = "FailedOperation.RefundDepositFailed"
//  FAILEDOPERATION_TOOMANYRESOURCES = "FailedOperation.TooManyResources"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INSTANCEINPROCESS = "InvalidParameter.InstanceInProcess"
//  INVALIDPARAMETER_INVALIDDATAENGINENAME = "InvalidParameter.InvalidDataEngineName"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTUNIQUE = "ResourceNotFound.DataEngineNotUnique"
//  UNAUTHORIZEDOPERATION_DELETECOMPUTINGENGINE = "UnauthorizedOperation.DeleteComputingEngine"
//  UNAUTHORIZEDOPERATION_NOPAYMENTAUTHORITY = "UnauthorizedOperation.NoPaymentAuthority"
func (c *Client) DeleteDataEngineWithContext(ctx context.Context, request *DeleteDataEngineRequest) (response *DeleteDataEngineResponse, err error) {
    if request == nil {
        request = NewDeleteDataEngineRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DeleteDataEngine")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDataEngine require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDataEngineResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSparkAppRequest() (request *DeleteSparkAppRequest) {
    request = &DeleteSparkAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DeleteSparkApp")
    
    
    return
}

func NewDeleteSparkAppResponse() (response *DeleteSparkAppResponse) {
    response = &DeleteSparkAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSparkApp
// This API is used to delete a Spark job.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  INVALIDPARAMETER_SPARKJOBNOTFOUND = "InvalidParameter.SparkJobNotFound"
func (c *Client) DeleteSparkApp(request *DeleteSparkAppRequest) (response *DeleteSparkAppResponse, err error) {
    return c.DeleteSparkAppWithContext(context.Background(), request)
}

// DeleteSparkApp
// This API is used to delete a Spark job.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  INVALIDPARAMETER_SPARKJOBNOTFOUND = "InvalidParameter.SparkJobNotFound"
func (c *Client) DeleteSparkAppWithContext(ctx context.Context, request *DeleteSparkAppRequest) (response *DeleteSparkAppResponse, err error) {
    if request == nil {
        request = NewDeleteSparkAppRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DeleteSparkApp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSparkApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSparkAppResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteThirdPartyAccessUserRequest() (request *DeleteThirdPartyAccessUserRequest) {
    request = &DeleteThirdPartyAccessUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DeleteThirdPartyAccessUser")
    
    
    return
}

func NewDeleteThirdPartyAccessUserResponse() (response *DeleteThirdPartyAccessUserResponse) {
    response = &DeleteThirdPartyAccessUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteThirdPartyAccessUser
// This API is used to remove visits through the third-party platform.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteThirdPartyAccessUser(request *DeleteThirdPartyAccessUserRequest) (response *DeleteThirdPartyAccessUserResponse, err error) {
    return c.DeleteThirdPartyAccessUserWithContext(context.Background(), request)
}

// DeleteThirdPartyAccessUser
// This API is used to remove visits through the third-party platform.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteThirdPartyAccessUserWithContext(ctx context.Context, request *DeleteThirdPartyAccessUserRequest) (response *DeleteThirdPartyAccessUserResponse, err error) {
    if request == nil {
        request = NewDeleteThirdPartyAccessUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DeleteThirdPartyAccessUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteThirdPartyAccessUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteThirdPartyAccessUserResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUserRequest() (request *DeleteUserRequest) {
    request = &DeleteUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DeleteUser")
    
    
    return
}

func NewDeleteUserResponse() (response *DeleteUserResponse) {
    response = &DeleteUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteUser
// This API is used to delete users.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETUSERINFOFAILED = "FailedOperation.GetUserInfoFailed"
//  FAILEDOPERATION_REVOKEPOLICYFAILED = "FailedOperation.RevokePolicyFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDUSERNAME = "InvalidParameter.InvalidUserName"
//  RESOURCESSOLDOUT_UNAUTHORIZEDREVOKEPOLICY = "ResourcesSoldOut.UnauthorizedRevokePolicy"
//  UNAUTHORIZEDOPERATION_DELETEUSER = "UnauthorizedOperation.DeleteUser"
//  UNAUTHORIZEDOPERATION_REVOKEPOLICY = "UnauthorizedOperation.RevokePolicy"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
//  UNSUPPORTEDOPERATION_MODIFYOWNERUNSUPPORTED = "UnsupportedOperation.ModifyOwnerUnsupported"
func (c *Client) DeleteUser(request *DeleteUserRequest) (response *DeleteUserResponse, err error) {
    return c.DeleteUserWithContext(context.Background(), request)
}

// DeleteUser
// This API is used to delete users.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETUSERINFOFAILED = "FailedOperation.GetUserInfoFailed"
//  FAILEDOPERATION_REVOKEPOLICYFAILED = "FailedOperation.RevokePolicyFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDUSERNAME = "InvalidParameter.InvalidUserName"
//  RESOURCESSOLDOUT_UNAUTHORIZEDREVOKEPOLICY = "ResourcesSoldOut.UnauthorizedRevokePolicy"
//  UNAUTHORIZEDOPERATION_DELETEUSER = "UnauthorizedOperation.DeleteUser"
//  UNAUTHORIZEDOPERATION_REVOKEPOLICY = "UnauthorizedOperation.RevokePolicy"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
//  UNSUPPORTEDOPERATION_MODIFYOWNERUNSUPPORTED = "UnsupportedOperation.ModifyOwnerUnsupported"
func (c *Client) DeleteUserWithContext(ctx context.Context, request *DeleteUserRequest) (response *DeleteUserResponse, err error) {
    if request == nil {
        request = NewDeleteUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DeleteUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUserResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUsersFromWorkGroupRequest() (request *DeleteUsersFromWorkGroupRequest) {
    request = &DeleteUsersFromWorkGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DeleteUsersFromWorkGroup")
    
    
    return
}

func NewDeleteUsersFromWorkGroupResponse() (response *DeleteUsersFromWorkGroupResponse) {
    response = &DeleteUsersFromWorkGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteUsersFromWorkGroup
// This API is used to delete users from working groups.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDGROUPID = "InvalidParameter.InvalidGroupId"
//  UNAUTHORIZEDOPERATION_DELETEUSERSFROMWORKGROUP = "UnauthorizedOperation.DeleteUsersFromWorkgroup"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
func (c *Client) DeleteUsersFromWorkGroup(request *DeleteUsersFromWorkGroupRequest) (response *DeleteUsersFromWorkGroupResponse, err error) {
    return c.DeleteUsersFromWorkGroupWithContext(context.Background(), request)
}

// DeleteUsersFromWorkGroup
// This API is used to delete users from working groups.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDGROUPID = "InvalidParameter.InvalidGroupId"
//  UNAUTHORIZEDOPERATION_DELETEUSERSFROMWORKGROUP = "UnauthorizedOperation.DeleteUsersFromWorkgroup"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
func (c *Client) DeleteUsersFromWorkGroupWithContext(ctx context.Context, request *DeleteUsersFromWorkGroupRequest) (response *DeleteUsersFromWorkGroupResponse, err error) {
    if request == nil {
        request = NewDeleteUsersFromWorkGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DeleteUsersFromWorkGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUsersFromWorkGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUsersFromWorkGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWorkGroupRequest() (request *DeleteWorkGroupRequest) {
    request = &DeleteWorkGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DeleteWorkGroup")
    
    
    return
}

func NewDeleteWorkGroupResponse() (response *DeleteWorkGroupResponse) {
    response = &DeleteWorkGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteWorkGroup
// This API is used to delete working groups.
//
// error code that may be returned:
//  FAILEDOPERATION_REVOKEPOLICYFAILED = "FailedOperation.RevokePolicyFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  UNAUTHORIZEDOPERATION_DELETEWORKGROUP = "UnauthorizedOperation.DeleteWorkgroup"
//  UNAUTHORIZEDOPERATION_REVOKEPOLICY = "UnauthorizedOperation.RevokePolicy"
func (c *Client) DeleteWorkGroup(request *DeleteWorkGroupRequest) (response *DeleteWorkGroupResponse, err error) {
    return c.DeleteWorkGroupWithContext(context.Background(), request)
}

// DeleteWorkGroup
// This API is used to delete working groups.
//
// error code that may be returned:
//  FAILEDOPERATION_REVOKEPOLICYFAILED = "FailedOperation.RevokePolicyFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  UNAUTHORIZEDOPERATION_DELETEWORKGROUP = "UnauthorizedOperation.DeleteWorkgroup"
//  UNAUTHORIZEDOPERATION_REVOKEPOLICY = "UnauthorizedOperation.RevokePolicy"
func (c *Client) DeleteWorkGroupWithContext(ctx context.Context, request *DeleteWorkGroupRequest) (response *DeleteWorkGroupResponse, err error) {
    if request == nil {
        request = NewDeleteWorkGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DeleteWorkGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWorkGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteWorkGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAdvancedStoreLocationRequest() (request *DescribeAdvancedStoreLocationRequest) {
    request = &DescribeAdvancedStoreLocationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeAdvancedStoreLocation")
    
    
    return
}

func NewDescribeAdvancedStoreLocationResponse() (response *DescribeAdvancedStoreLocationResponse) {
    response = &DescribeAdvancedStoreLocationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAdvancedStoreLocation
// This API is used to query the advanced settings of the SQL query interface.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAdvancedStoreLocation(request *DescribeAdvancedStoreLocationRequest) (response *DescribeAdvancedStoreLocationResponse, err error) {
    return c.DescribeAdvancedStoreLocationWithContext(context.Background(), request)
}

// DescribeAdvancedStoreLocation
// This API is used to query the advanced settings of the SQL query interface.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAdvancedStoreLocationWithContext(ctx context.Context, request *DescribeAdvancedStoreLocationRequest) (response *DescribeAdvancedStoreLocationResponse, err error) {
    if request == nil {
        request = NewDescribeAdvancedStoreLocationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeAdvancedStoreLocation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAdvancedStoreLocation require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAdvancedStoreLocationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDLCCatalogAccessRequest() (request *DescribeDLCCatalogAccessRequest) {
    request = &DescribeDLCCatalogAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeDLCCatalogAccess")
    
    
    return
}

func NewDescribeDLCCatalogAccessResponse() (response *DescribeDLCCatalogAccessResponse) {
    response = &DescribeDLCCatalogAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDLCCatalogAccess
// This API is used to query the DLC Catalog authorization list.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeDLCCatalogAccess(request *DescribeDLCCatalogAccessRequest) (response *DescribeDLCCatalogAccessResponse, err error) {
    return c.DescribeDLCCatalogAccessWithContext(context.Background(), request)
}

// DescribeDLCCatalogAccess
// This API is used to query the DLC Catalog authorization list.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeDLCCatalogAccessWithContext(ctx context.Context, request *DescribeDLCCatalogAccessRequest) (response *DescribeDLCCatalogAccessResponse, err error) {
    if request == nil {
        request = NewDescribeDLCCatalogAccessRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeDLCCatalogAccess")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDLCCatalogAccess require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDLCCatalogAccessResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDLCTableRequest() (request *DescribeDLCTableRequest) {
    request = &DescribeDLCTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeDLCTable")
    
    
    return
}

func NewDescribeDLCTableResponse() (response *DescribeDLCTableResponse) {
    response = &DescribeDLCTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDLCTable
// This API is used to obtain the table.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDLCTable(request *DescribeDLCTableRequest) (response *DescribeDLCTableResponse, err error) {
    return c.DescribeDLCTableWithContext(context.Background(), request)
}

// DescribeDLCTable
// This API is used to obtain the table.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDLCTableWithContext(ctx context.Context, request *DescribeDLCTableRequest) (response *DescribeDLCTableResponse, err error) {
    if request == nil {
        request = NewDescribeDLCTableRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeDLCTable")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDLCTable require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDLCTableResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDLCTableListRequest() (request *DescribeDLCTableListRequest) {
    request = &DescribeDLCTableListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeDLCTableList")
    
    
    return
}

func NewDescribeDLCTableListResponse() (response *DescribeDLCTableListResponse) {
    response = &DescribeDLCTableListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDLCTableList
// This API is used to obtain the list of tables.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDLCTableList(request *DescribeDLCTableListRequest) (response *DescribeDLCTableListResponse, err error) {
    return c.DescribeDLCTableListWithContext(context.Background(), request)
}

// DescribeDLCTableList
// This API is used to obtain the list of tables.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDLCTableListWithContext(ctx context.Context, request *DescribeDLCTableListRequest) (response *DescribeDLCTableListResponse, err error) {
    if request == nil {
        request = NewDescribeDLCTableListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeDLCTableList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDLCTableList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDLCTableListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDMSDatabaseRequest() (request *DescribeDMSDatabaseRequest) {
    request = &DescribeDMSDatabaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeDMSDatabase")
    
    
    return
}

func NewDescribeDMSDatabaseResponse() (response *DescribeDMSDatabaseResponse) {
    response = &DescribeDMSDatabaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDMSDatabase
// This API is used to obtain databases in the DMS metadata module.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDMSDatabase(request *DescribeDMSDatabaseRequest) (response *DescribeDMSDatabaseResponse, err error) {
    return c.DescribeDMSDatabaseWithContext(context.Background(), request)
}

// DescribeDMSDatabase
// This API is used to obtain databases in the DMS metadata module.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDMSDatabaseWithContext(ctx context.Context, request *DescribeDMSDatabaseRequest) (response *DescribeDMSDatabaseResponse, err error) {
    if request == nil {
        request = NewDescribeDMSDatabaseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeDMSDatabase")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDMSDatabase require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDMSDatabaseResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDMSDatabaseListRequest() (request *DescribeDMSDatabaseListRequest) {
    request = &DescribeDMSDatabaseListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeDMSDatabaseList")
    
    
    return
}

func NewDescribeDMSDatabaseListResponse() (response *DescribeDMSDatabaseListResponse) {
    response = &DescribeDMSDatabaseListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDMSDatabaseList
// This API is used to obtain the list of databases.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDMSDatabaseList(request *DescribeDMSDatabaseListRequest) (response *DescribeDMSDatabaseListResponse, err error) {
    return c.DescribeDMSDatabaseListWithContext(context.Background(), request)
}

// DescribeDMSDatabaseList
// This API is used to obtain the list of databases.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDMSDatabaseListWithContext(ctx context.Context, request *DescribeDMSDatabaseListRequest) (response *DescribeDMSDatabaseListResponse, err error) {
    if request == nil {
        request = NewDescribeDMSDatabaseListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeDMSDatabaseList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDMSDatabaseList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDMSDatabaseListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataEngineRequest() (request *DescribeDataEngineRequest) {
    request = &DescribeDataEngineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeDataEngine")
    
    
    return
}

func NewDescribeDataEngineResponse() (response *DescribeDataEngineResponse) {
    response = &DescribeDataEngineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDataEngine
// This API is used to obtain detailed data engine information based on names.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
func (c *Client) DescribeDataEngine(request *DescribeDataEngineRequest) (response *DescribeDataEngineResponse, err error) {
    return c.DescribeDataEngineWithContext(context.Background(), request)
}

// DescribeDataEngine
// This API is used to obtain detailed data engine information based on names.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
func (c *Client) DescribeDataEngineWithContext(ctx context.Context, request *DescribeDataEngineRequest) (response *DescribeDataEngineResponse, err error) {
    if request == nil {
        request = NewDescribeDataEngineRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeDataEngine")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataEngine require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataEngineResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataEngineImageVersionsRequest() (request *DescribeDataEngineImageVersionsRequest) {
    request = &DescribeDataEngineImageVersionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeDataEngineImageVersions")
    
    
    return
}

func NewDescribeDataEngineImageVersionsResponse() (response *DescribeDataEngineImageVersionsResponse) {
    response = &DescribeDataEngineImageVersionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDataEngineImageVersions
// This API is used to obtain the major version image list of exclusive clusters.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_IMAGEENGINETYPENOTMATCH = "InvalidParameter.ImageEngineTypeNotMatch"
//  INVALIDPARAMETER_IMAGEISPUBLICNOTMATCH = "InvalidParameter.ImageIsPublicNotMatch"
//  INVALIDPARAMETER_IMAGEUSERRECORDSTYPENOTMATCH = "InvalidParameter.ImageUserRecordsTypeNotMatch"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_IMAGEVERSIONNOTACTIVITY = "ResourceNotFound.ImageVersionNotActivity"
func (c *Client) DescribeDataEngineImageVersions(request *DescribeDataEngineImageVersionsRequest) (response *DescribeDataEngineImageVersionsResponse, err error) {
    return c.DescribeDataEngineImageVersionsWithContext(context.Background(), request)
}

// DescribeDataEngineImageVersions
// This API is used to obtain the major version image list of exclusive clusters.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_IMAGEENGINETYPENOTMATCH = "InvalidParameter.ImageEngineTypeNotMatch"
//  INVALIDPARAMETER_IMAGEISPUBLICNOTMATCH = "InvalidParameter.ImageIsPublicNotMatch"
//  INVALIDPARAMETER_IMAGEUSERRECORDSTYPENOTMATCH = "InvalidParameter.ImageUserRecordsTypeNotMatch"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_IMAGEVERSIONNOTACTIVITY = "ResourceNotFound.ImageVersionNotActivity"
func (c *Client) DescribeDataEngineImageVersionsWithContext(ctx context.Context, request *DescribeDataEngineImageVersionsRequest) (response *DescribeDataEngineImageVersionsResponse, err error) {
    if request == nil {
        request = NewDescribeDataEngineImageVersionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeDataEngineImageVersions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataEngineImageVersions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataEngineImageVersionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataEnginePythonSparkImagesRequest() (request *DescribeDataEnginePythonSparkImagesRequest) {
    request = &DescribeDataEnginePythonSparkImagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeDataEnginePythonSparkImages")
    
    
    return
}

func NewDescribeDataEnginePythonSparkImagesResponse() (response *DescribeDataEnginePythonSparkImagesResponse) {
    response = &DescribeDataEnginePythonSparkImagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDataEnginePythonSparkImages
// This API is used to obtain the PYSPARK image list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_IMAGEUSERRECORDSTYPENOTMATCH = "InvalidParameter.ImageUserRecordsTypeNotMatch"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDataEnginePythonSparkImages(request *DescribeDataEnginePythonSparkImagesRequest) (response *DescribeDataEnginePythonSparkImagesResponse, err error) {
    return c.DescribeDataEnginePythonSparkImagesWithContext(context.Background(), request)
}

// DescribeDataEnginePythonSparkImages
// This API is used to obtain the PYSPARK image list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_IMAGEUSERRECORDSTYPENOTMATCH = "InvalidParameter.ImageUserRecordsTypeNotMatch"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDataEnginePythonSparkImagesWithContext(ctx context.Context, request *DescribeDataEnginePythonSparkImagesRequest) (response *DescribeDataEnginePythonSparkImagesResponse, err error) {
    if request == nil {
        request = NewDescribeDataEnginePythonSparkImagesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeDataEnginePythonSparkImages")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataEnginePythonSparkImages require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataEnginePythonSparkImagesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataEnginesScaleDetailRequest() (request *DescribeDataEnginesScaleDetailRequest) {
    request = &DescribeDataEnginesScaleDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeDataEnginesScaleDetail")
    
    
    return
}

func NewDescribeDataEnginesScaleDetailResponse() (response *DescribeDataEnginesScaleDetailResponse) {
    response = &DescribeDataEnginesScaleDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDataEnginesScaleDetail
// This API is used to query engine specification details.
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_MONITORCOMPUTINGENGINE = "UnauthorizedOperation.MonitorComputingEngine"
func (c *Client) DescribeDataEnginesScaleDetail(request *DescribeDataEnginesScaleDetailRequest) (response *DescribeDataEnginesScaleDetailResponse, err error) {
    return c.DescribeDataEnginesScaleDetailWithContext(context.Background(), request)
}

// DescribeDataEnginesScaleDetail
// This API is used to query engine specification details.
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_MONITORCOMPUTINGENGINE = "UnauthorizedOperation.MonitorComputingEngine"
func (c *Client) DescribeDataEnginesScaleDetailWithContext(ctx context.Context, request *DescribeDataEnginesScaleDetailRequest) (response *DescribeDataEnginesScaleDetailResponse, err error) {
    if request == nil {
        request = NewDescribeDataEnginesScaleDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeDataEnginesScaleDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataEnginesScaleDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataEnginesScaleDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEngineUsageInfoRequest() (request *DescribeEngineUsageInfoRequest) {
    request = &DescribeEngineUsageInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeEngineUsageInfo")
    
    
    return
}

func NewDescribeEngineUsageInfoResponse() (response *DescribeEngineUsageInfoResponse) {
    response = &DescribeEngineUsageInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEngineUsageInfo
// This API is used to query the resource usage of a data engine based on its ID.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeEngineUsageInfo(request *DescribeEngineUsageInfoRequest) (response *DescribeEngineUsageInfoResponse, err error) {
    return c.DescribeEngineUsageInfoWithContext(context.Background(), request)
}

// DescribeEngineUsageInfo
// This API is used to query the resource usage of a data engine based on its ID.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeEngineUsageInfoWithContext(ctx context.Context, request *DescribeEngineUsageInfoRequest) (response *DescribeEngineUsageInfoResponse, err error) {
    if request == nil {
        request = NewDescribeEngineUsageInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeEngineUsageInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEngineUsageInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEngineUsageInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeForbiddenTableProRequest() (request *DescribeForbiddenTableProRequest) {
    request = &DescribeForbiddenTableProRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeForbiddenTablePro")
    
    
    return
}

func NewDescribeForbiddenTableProResponse() (response *DescribeForbiddenTableProResponse) {
    response = &DescribeForbiddenTableProResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeForbiddenTablePro
// This API is used to get the list of disabled table attributes (new).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeForbiddenTablePro(request *DescribeForbiddenTableProRequest) (response *DescribeForbiddenTableProResponse, err error) {
    return c.DescribeForbiddenTableProWithContext(context.Background(), request)
}

// DescribeForbiddenTablePro
// This API is used to get the list of disabled table attributes (new).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeForbiddenTableProWithContext(ctx context.Context, request *DescribeForbiddenTableProRequest) (response *DescribeForbiddenTableProResponse, err error) {
    if request == nil {
        request = NewDescribeForbiddenTableProRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeForbiddenTablePro")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeForbiddenTablePro require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeForbiddenTableProResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeJobRequest() (request *DescribeJobRequest) {
    request = &DescribeJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeJob")
    
    
    return
}

func NewDescribeJobResponse() (response *DescribeJobResponse) {
    response = &DescribeJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeJob
// This API is used to obtain the job information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_JOBNOTFOUND = "InvalidParameter.JobNotFound"
func (c *Client) DescribeJob(request *DescribeJobRequest) (response *DescribeJobResponse, err error) {
    return c.DescribeJobWithContext(context.Background(), request)
}

// DescribeJob
// This API is used to obtain the job information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_JOBNOTFOUND = "InvalidParameter.JobNotFound"
func (c *Client) DescribeJobWithContext(ctx context.Context, request *DescribeJobRequest) (response *DescribeJobResponse, err error) {
    if request == nil {
        request = NewDescribeJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeJobsRequest() (request *DescribeJobsRequest) {
    request = &DescribeJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeJobs")
    
    
    return
}

func NewDescribeJobsResponse() (response *DescribeJobsResponse) {
    response = &DescribeJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeJobs
// This API is used to obtain the list of job information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_JOBNOTFOUND = "InvalidParameter.JobNotFound"
func (c *Client) DescribeJobs(request *DescribeJobsRequest) (response *DescribeJobsResponse, err error) {
    return c.DescribeJobsWithContext(context.Background(), request)
}

// DescribeJobs
// This API is used to obtain the list of job information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_JOBNOTFOUND = "InvalidParameter.JobNotFound"
func (c *Client) DescribeJobsWithContext(ctx context.Context, request *DescribeJobsRequest) (response *DescribeJobsResponse, err error) {
    if request == nil {
        request = NewDescribeJobsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeJobs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLakeFsDirSummaryRequest() (request *DescribeLakeFsDirSummaryRequest) {
    request = &DescribeLakeFsDirSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeLakeFsDirSummary")
    
    
    return
}

func NewDescribeLakeFsDirSummaryResponse() (response *DescribeLakeFsDirSummaryResponse) {
    response = &DescribeLakeFsDirSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLakeFsDirSummary
// This API is used to query the summary of a specified directory in a managed storage.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLakeFsDirSummary(request *DescribeLakeFsDirSummaryRequest) (response *DescribeLakeFsDirSummaryResponse, err error) {
    return c.DescribeLakeFsDirSummaryWithContext(context.Background(), request)
}

// DescribeLakeFsDirSummary
// This API is used to query the summary of a specified directory in a managed storage.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLakeFsDirSummaryWithContext(ctx context.Context, request *DescribeLakeFsDirSummaryRequest) (response *DescribeLakeFsDirSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeLakeFsDirSummaryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeLakeFsDirSummary")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLakeFsDirSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLakeFsDirSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLakeFsInfoRequest() (request *DescribeLakeFsInfoRequest) {
    request = &DescribeLakeFsInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeLakeFsInfo")
    
    
    return
}

func NewDescribeLakeFsInfoResponse() (response *DescribeLakeFsInfoResponse) {
    response = &DescribeLakeFsInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLakeFsInfo
// This API is used to query managed storage information.
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
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLakeFsInfo(request *DescribeLakeFsInfoRequest) (response *DescribeLakeFsInfoResponse, err error) {
    return c.DescribeLakeFsInfoWithContext(context.Background(), request)
}

// DescribeLakeFsInfo
// This API is used to query managed storage information.
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
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLakeFsInfoWithContext(ctx context.Context, request *DescribeLakeFsInfoRequest) (response *DescribeLakeFsInfoResponse, err error) {
    if request == nil {
        request = NewDescribeLakeFsInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeLakeFsInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLakeFsInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLakeFsInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOtherCHDFSBindingListRequest() (request *DescribeOtherCHDFSBindingListRequest) {
    request = &DescribeOtherCHDFSBindingListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeOtherCHDFSBindingList")
    
    
    return
}

func NewDescribeOtherCHDFSBindingListResponse() (response *DescribeOtherCHDFSBindingListResponse) {
    response = &DescribeOtherCHDFSBindingListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOtherCHDFSBindingList
// This API is used to query the list of metadata acceleration buckets bound to other products.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeOtherCHDFSBindingList(request *DescribeOtherCHDFSBindingListRequest) (response *DescribeOtherCHDFSBindingListResponse, err error) {
    return c.DescribeOtherCHDFSBindingListWithContext(context.Background(), request)
}

// DescribeOtherCHDFSBindingList
// This API is used to query the list of metadata acceleration buckets bound to other products.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeOtherCHDFSBindingListWithContext(ctx context.Context, request *DescribeOtherCHDFSBindingListRequest) (response *DescribeOtherCHDFSBindingListResponse, err error) {
    if request == nil {
        request = NewDescribeOtherCHDFSBindingListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeOtherCHDFSBindingList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOtherCHDFSBindingList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOtherCHDFSBindingListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQueryRequest() (request *DescribeQueryRequest) {
    request = &DescribeQueryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeQuery")
    
    
    return
}

func NewDescribeQueryResponse() (response *DescribeQueryResponse) {
    response = &DescribeQueryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeQuery
// This API is used to obtain the query results.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
func (c *Client) DescribeQuery(request *DescribeQueryRequest) (response *DescribeQueryResponse, err error) {
    return c.DescribeQueryWithContext(context.Background(), request)
}

// DescribeQuery
// This API is used to obtain the query results.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
func (c *Client) DescribeQueryWithContext(ctx context.Context, request *DescribeQueryRequest) (response *DescribeQueryResponse, err error) {
    if request == nil {
        request = NewDescribeQueryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeQuery")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeQuery require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeQueryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResultDownloadRequest() (request *DescribeResultDownloadRequest) {
    request = &DescribeResultDownloadRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeResultDownload")
    
    
    return
}

func NewDescribeResultDownloadResponse() (response *DescribeResultDownloadResponse) {
    response = &DescribeResultDownloadResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeResultDownload
// This API is used to get a query result download task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeResultDownload(request *DescribeResultDownloadRequest) (response *DescribeResultDownloadResponse, err error) {
    return c.DescribeResultDownloadWithContext(context.Background(), request)
}

// DescribeResultDownload
// This API is used to get a query result download task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeResultDownloadWithContext(ctx context.Context, request *DescribeResultDownloadRequest) (response *DescribeResultDownloadResponse, err error) {
    if request == nil {
        request = NewDescribeResultDownloadRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeResultDownload")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResultDownload require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResultDownloadResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSessionImageVersionRequest() (request *DescribeSessionImageVersionRequest) {
    request = &DescribeSessionImageVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeSessionImageVersion")
    
    
    return
}

func NewDescribeSessionImageVersionResponse() (response *DescribeSessionImageVersionResponse) {
    response = &DescribeSessionImageVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSessionImageVersion
// This API is used to retrieve all built-in images of all minor versions under a specified major version.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_INVALIDCONFIGVALUEREGEXPNOTMATCH = "InvalidParameter.InvalidConfigValueRegexpNotMatch"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
func (c *Client) DescribeSessionImageVersion(request *DescribeSessionImageVersionRequest) (response *DescribeSessionImageVersionResponse, err error) {
    return c.DescribeSessionImageVersionWithContext(context.Background(), request)
}

// DescribeSessionImageVersion
// This API is used to retrieve all built-in images of all minor versions under a specified major version.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_INVALIDCONFIGVALUEREGEXPNOTMATCH = "InvalidParameter.InvalidConfigValueRegexpNotMatch"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
func (c *Client) DescribeSessionImageVersionWithContext(ctx context.Context, request *DescribeSessionImageVersionRequest) (response *DescribeSessionImageVersionResponse, err error) {
    if request == nil {
        request = NewDescribeSessionImageVersionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeSessionImageVersion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSessionImageVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSessionImageVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSparkAppJobRequest() (request *DescribeSparkAppJobRequest) {
    request = &DescribeSparkAppJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeSparkAppJob")
    
    
    return
}

func NewDescribeSparkAppJobResponse() (response *DescribeSparkAppJobResponse) {
    response = &DescribeSparkAppJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSparkAppJob
// u200cThis API is used to query the information of a Spark job.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSPARKAPPPARAM = "InvalidParameter.InvalidSparkAppParam"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  INVALIDPARAMETER_SPARKJOBNOTFOUND = "InvalidParameter.SparkJobNotFound"
//  INVALIDPARAMETER_SPARKJOBNOTUNIQUE = "InvalidParameter.SparkJobNotUnique"
func (c *Client) DescribeSparkAppJob(request *DescribeSparkAppJobRequest) (response *DescribeSparkAppJobResponse, err error) {
    return c.DescribeSparkAppJobWithContext(context.Background(), request)
}

// DescribeSparkAppJob
// u200cThis API is used to query the information of a Spark job.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSPARKAPPPARAM = "InvalidParameter.InvalidSparkAppParam"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  INVALIDPARAMETER_SPARKJOBNOTFOUND = "InvalidParameter.SparkJobNotFound"
//  INVALIDPARAMETER_SPARKJOBNOTUNIQUE = "InvalidParameter.SparkJobNotUnique"
func (c *Client) DescribeSparkAppJobWithContext(ctx context.Context, request *DescribeSparkAppJobRequest) (response *DescribeSparkAppJobResponse, err error) {
    if request == nil {
        request = NewDescribeSparkAppJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeSparkAppJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSparkAppJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSparkAppJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSparkAppJobsRequest() (request *DescribeSparkAppJobsRequest) {
    request = &DescribeSparkAppJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeSparkAppJobs")
    
    
    return
}

func NewDescribeSparkAppJobsResponse() (response *DescribeSparkAppJobsResponse) {
    response = &DescribeSparkAppJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSparkAppJobs
// This API is used to query the list of Spark jobs.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  INVALIDPARAMETER_INVALIDTIMEFORMAT = "InvalidParameter.InvalidTimeFormat"
//  INVALIDPARAMETER_SPARKJOBFILTERSKEYTYPENOTMATH = "InvalidParameter.SparkJobFiltersKeyTypeNotMath"
//  INVALIDPARAMETER_SPARKJOBSORTBYTYPENOTMATCH = "InvalidParameter.SparkJobSortByTypeNotMatch"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BATCHSQLTASKNOTFOUND = "ResourceNotFound.BatchSQLTaskNotFound"
func (c *Client) DescribeSparkAppJobs(request *DescribeSparkAppJobsRequest) (response *DescribeSparkAppJobsResponse, err error) {
    return c.DescribeSparkAppJobsWithContext(context.Background(), request)
}

// DescribeSparkAppJobs
// This API is used to query the list of Spark jobs.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  INVALIDPARAMETER_INVALIDTIMEFORMAT = "InvalidParameter.InvalidTimeFormat"
//  INVALIDPARAMETER_SPARKJOBFILTERSKEYTYPENOTMATH = "InvalidParameter.SparkJobFiltersKeyTypeNotMath"
//  INVALIDPARAMETER_SPARKJOBSORTBYTYPENOTMATCH = "InvalidParameter.SparkJobSortByTypeNotMatch"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BATCHSQLTASKNOTFOUND = "ResourceNotFound.BatchSQLTaskNotFound"
func (c *Client) DescribeSparkAppJobsWithContext(ctx context.Context, request *DescribeSparkAppJobsRequest) (response *DescribeSparkAppJobsResponse, err error) {
    if request == nil {
        request = NewDescribeSparkAppJobsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeSparkAppJobs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSparkAppJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSparkAppJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSparkAppTasksRequest() (request *DescribeSparkAppTasksRequest) {
    request = &DescribeSparkAppTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeSparkAppTasks")
    
    
    return
}

func NewDescribeSparkAppTasksResponse() (response *DescribeSparkAppTasksResponse) {
    response = &DescribeSparkAppTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSparkAppTasks
// This API is used to query the list of running task instances of a Spark job.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_FILTERSVALUESNUMBEROUTOFLIMIT = "InvalidParameter.FiltersValuesNumberOutOfLimit"
//  INVALIDPARAMETER_INVALIDTIMEFORMAT = "InvalidParameter.InvalidTimeFormat"
//  INVALIDPARAMETER_INVALIDTIMEPARAMETER = "InvalidParameter.InvalidTimeParameter"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  INVALIDPARAMETER_SPARKJOBFILTERSKEYTYPENOTMATH = "InvalidParameter.SparkJobFiltersKeyTypeNotMath"
//  INVALIDPARAMETER_SPARKJOBNOTFOUND = "InvalidParameter.SparkJobNotFound"
//  INVALIDPARAMETER_SPARKJOBNOTUNIQUE = "InvalidParameter.SparkJobNotUnique"
//  RESOURCEUNAVAILABLE_WHITELISTFUNCTION = "ResourceUnavailable.WhiteListFunction"
func (c *Client) DescribeSparkAppTasks(request *DescribeSparkAppTasksRequest) (response *DescribeSparkAppTasksResponse, err error) {
    return c.DescribeSparkAppTasksWithContext(context.Background(), request)
}

// DescribeSparkAppTasks
// This API is used to query the list of running task instances of a Spark job.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_FILTERSVALUESNUMBEROUTOFLIMIT = "InvalidParameter.FiltersValuesNumberOutOfLimit"
//  INVALIDPARAMETER_INVALIDTIMEFORMAT = "InvalidParameter.InvalidTimeFormat"
//  INVALIDPARAMETER_INVALIDTIMEPARAMETER = "InvalidParameter.InvalidTimeParameter"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  INVALIDPARAMETER_SPARKJOBFILTERSKEYTYPENOTMATH = "InvalidParameter.SparkJobFiltersKeyTypeNotMath"
//  INVALIDPARAMETER_SPARKJOBNOTFOUND = "InvalidParameter.SparkJobNotFound"
//  INVALIDPARAMETER_SPARKJOBNOTUNIQUE = "InvalidParameter.SparkJobNotUnique"
//  RESOURCEUNAVAILABLE_WHITELISTFUNCTION = "ResourceUnavailable.WhiteListFunction"
func (c *Client) DescribeSparkAppTasksWithContext(ctx context.Context, request *DescribeSparkAppTasksRequest) (response *DescribeSparkAppTasksResponse, err error) {
    if request == nil {
        request = NewDescribeSparkAppTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeSparkAppTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSparkAppTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSparkAppTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSparkSessionBatchSqlLogRequest() (request *DescribeSparkSessionBatchSqlLogRequest) {
    request = &DescribeSparkSessionBatchSqlLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeSparkSessionBatchSqlLog")
    
    
    return
}

func NewDescribeSparkSessionBatchSqlLogResponse() (response *DescribeSparkSessionBatchSqlLogResponse) {
    response = &DescribeSparkSessionBatchSqlLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSparkSessionBatchSqlLog
// This API is used to query Spark SQL batch task logs.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BATCHSQLTASKNOTFOUND = "ResourceNotFound.BatchSQLTaskNotFound"
//  RESOURCENOTFOUND_BATCHSQLTASKNOTUNIQUE = "ResourceNotFound.BatchSQLTaskNotUnique"
func (c *Client) DescribeSparkSessionBatchSqlLog(request *DescribeSparkSessionBatchSqlLogRequest) (response *DescribeSparkSessionBatchSqlLogResponse, err error) {
    return c.DescribeSparkSessionBatchSqlLogWithContext(context.Background(), request)
}

// DescribeSparkSessionBatchSqlLog
// This API is used to query Spark SQL batch task logs.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BATCHSQLTASKNOTFOUND = "ResourceNotFound.BatchSQLTaskNotFound"
//  RESOURCENOTFOUND_BATCHSQLTASKNOTUNIQUE = "ResourceNotFound.BatchSQLTaskNotUnique"
func (c *Client) DescribeSparkSessionBatchSqlLogWithContext(ctx context.Context, request *DescribeSparkSessionBatchSqlLogRequest) (response *DescribeSparkSessionBatchSqlLogResponse, err error) {
    if request == nil {
        request = NewDescribeSparkSessionBatchSqlLogRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeSparkSessionBatchSqlLog")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSparkSessionBatchSqlLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSparkSessionBatchSqlLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStoreLocationRequest() (request *DescribeStoreLocationRequest) {
    request = &DescribeStoreLocationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeStoreLocation")
    
    
    return
}

func NewDescribeStoreLocationResponse() (response *DescribeStoreLocationResponse) {
    response = &DescribeStoreLocationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStoreLocation
// This API is used to query the storage location of calculation results.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeStoreLocation(request *DescribeStoreLocationRequest) (response *DescribeStoreLocationResponse, err error) {
    return c.DescribeStoreLocationWithContext(context.Background(), request)
}

// DescribeStoreLocation
// This API is used to query the storage location of calculation results.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeStoreLocationWithContext(ctx context.Context, request *DescribeStoreLocationRequest) (response *DescribeStoreLocationResponse, err error) {
    if request == nil {
        request = NewDescribeStoreLocationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeStoreLocation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStoreLocation require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStoreLocationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubUserAccessPolicyRequest() (request *DescribeSubUserAccessPolicyRequest) {
    request = &DescribeSubUserAccessPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeSubUserAccessPolicy")
    
    
    return
}

func NewDescribeSubUserAccessPolicyResponse() (response *DescribeSubUserAccessPolicyResponse) {
    response = &DescribeSubUserAccessPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSubUserAccessPolicy
// This API is used to query the sub-user's visiting policy for users accessing through the third-party platform.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSubUserAccessPolicy(request *DescribeSubUserAccessPolicyRequest) (response *DescribeSubUserAccessPolicyResponse, err error) {
    return c.DescribeSubUserAccessPolicyWithContext(context.Background(), request)
}

// DescribeSubUserAccessPolicy
// This API is used to query the sub-user's visiting policy for users accessing through the third-party platform.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSubUserAccessPolicyWithContext(ctx context.Context, request *DescribeSubUserAccessPolicyRequest) (response *DescribeSubUserAccessPolicyResponse, err error) {
    if request == nil {
        request = NewDescribeSubUserAccessPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeSubUserAccessPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSubUserAccessPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSubUserAccessPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTablesNameRequest() (request *DescribeTablesNameRequest) {
    request = &DescribeTablesNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeTablesName")
    
    
    return
}

func NewDescribeTablesNameResponse() (response *DescribeTablesNameResponse) {
    response = &DescribeTablesNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTablesName
// This API is used to query the data table name list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_DATASOURCENOTFOUND = "ResourceNotFound.DatasourceNotFound"
func (c *Client) DescribeTablesName(request *DescribeTablesNameRequest) (response *DescribeTablesNameResponse, err error) {
    return c.DescribeTablesNameWithContext(context.Background(), request)
}

// DescribeTablesName
// This API is used to query the data table name list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_DATASOURCENOTFOUND = "ResourceNotFound.DatasourceNotFound"
func (c *Client) DescribeTablesNameWithContext(ctx context.Context, request *DescribeTablesNameRequest) (response *DescribeTablesNameResponse, err error) {
    if request == nil {
        request = NewDescribeTablesNameRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeTablesName")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTablesName require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTablesNameResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskResultRequest() (request *DescribeTaskResultRequest) {
    request = &DescribeTaskResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeTaskResult")
    
    
    return
}

func NewDescribeTaskResultResponse() (response *DescribeTaskResultResponse) {
    response = &DescribeTaskResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTaskResult
// This API is used to query the result of a task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_HTTPCLIENTDOREQUESTFAILED = "FailedOperation.HttpClientDoRequestFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDMAXRESULTS = "InvalidParameter.InvalidMaxResults"
//  INVALIDPARAMETER_INVALIDSQLTASKMAXRESULTS = "InvalidParameter.InvalidSQLTaskMaxResults"
//  INVALIDPARAMETER_INVALIDTASKID = "InvalidParameter.InvalidTaskId"
//  INVALIDPARAMETER_MAXRESULTONLYSUPPORTHUNDRED = "InvalidParameter.MaxResultOnlySupportHundred"
func (c *Client) DescribeTaskResult(request *DescribeTaskResultRequest) (response *DescribeTaskResultResponse, err error) {
    return c.DescribeTaskResultWithContext(context.Background(), request)
}

// DescribeTaskResult
// This API is used to query the result of a task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_HTTPCLIENTDOREQUESTFAILED = "FailedOperation.HttpClientDoRequestFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDMAXRESULTS = "InvalidParameter.InvalidMaxResults"
//  INVALIDPARAMETER_INVALIDSQLTASKMAXRESULTS = "InvalidParameter.InvalidSQLTaskMaxResults"
//  INVALIDPARAMETER_INVALIDTASKID = "InvalidParameter.InvalidTaskId"
//  INVALIDPARAMETER_MAXRESULTONLYSUPPORTHUNDRED = "InvalidParameter.MaxResultOnlySupportHundred"
func (c *Client) DescribeTaskResultWithContext(ctx context.Context, request *DescribeTaskResultRequest) (response *DescribeTaskResultResponse, err error) {
    if request == nil {
        request = NewDescribeTaskResultRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeTaskResult")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskStatisticsRequest() (request *DescribeTaskStatisticsRequest) {
    request = &DescribeTaskStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeTaskStatistics")
    
    
    return
}

func NewDescribeTaskStatisticsResponse() (response *DescribeTaskStatisticsResponse) {
    response = &DescribeTaskStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTaskStatistics
// This API is used to describe the information on task statistics.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
func (c *Client) DescribeTaskStatistics(request *DescribeTaskStatisticsRequest) (response *DescribeTaskStatisticsResponse, err error) {
    return c.DescribeTaskStatisticsWithContext(context.Background(), request)
}

// DescribeTaskStatistics
// This API is used to describe the information on task statistics.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
func (c *Client) DescribeTaskStatisticsWithContext(ctx context.Context, request *DescribeTaskStatisticsRequest) (response *DescribeTaskStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeTaskStatisticsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeTaskStatistics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTasksRequest() (request *DescribeTasksRequest) {
    request = &DescribeTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeTasks")
    
    
    return
}

func NewDescribeTasksResponse() (response *DescribeTasksResponse) {
    response = &DescribeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTasks
// This API is used to query the list of tasks.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FILTERSVALUESNUMBEROUTOFLIMIT = "InvalidParameter.FiltersValuesNumberOutOfLimit"
//  INVALIDPARAMETER_INVALIDFILTERLENGTH = "InvalidParameter.InvalidFilterLength"
//  INVALIDPARAMETER_INVALIDTIMEFORMAT = "InvalidParameter.InvalidTimeFormat"
//  INVALIDPARAMETER_PARAMETERBASE64DECODEFAILED = "InvalidParameter.ParameterBase64DecodeFailed"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  INVALIDPARAMETER_SQLTASKFILTERSKEYTYPENOTMATH = "InvalidParameter.SQLTaskFiltersKeyTypeNotMath"
//  INVALIDPARAMETER_SQLTASKNOTFOUND = "InvalidParameter.SQLTaskNotFound"
//  INVALIDPARAMETER_SQLTASKSORTBYTYPENOTMATCH = "InvalidParameter.SQLTaskSortByTypeNotMatch"
//  INVALIDPARAMETER_SPARKJOBNOTFOUND = "InvalidParameter.SparkJobNotFound"
//  INVALIDPARAMETER_TASKSTATETYPENOTMATH = "InvalidParameter.TaskStateTypeNotMath"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeTasks(request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    return c.DescribeTasksWithContext(context.Background(), request)
}

// DescribeTasks
// This API is used to query the list of tasks.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FILTERSVALUESNUMBEROUTOFLIMIT = "InvalidParameter.FiltersValuesNumberOutOfLimit"
//  INVALIDPARAMETER_INVALIDFILTERLENGTH = "InvalidParameter.InvalidFilterLength"
//  INVALIDPARAMETER_INVALIDTIMEFORMAT = "InvalidParameter.InvalidTimeFormat"
//  INVALIDPARAMETER_PARAMETERBASE64DECODEFAILED = "InvalidParameter.ParameterBase64DecodeFailed"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  INVALIDPARAMETER_SQLTASKFILTERSKEYTYPENOTMATH = "InvalidParameter.SQLTaskFiltersKeyTypeNotMath"
//  INVALIDPARAMETER_SQLTASKNOTFOUND = "InvalidParameter.SQLTaskNotFound"
//  INVALIDPARAMETER_SQLTASKSORTBYTYPENOTMATCH = "InvalidParameter.SQLTaskSortByTypeNotMatch"
//  INVALIDPARAMETER_SPARKJOBNOTFOUND = "InvalidParameter.SparkJobNotFound"
//  INVALIDPARAMETER_TASKSTATETYPENOTMATH = "InvalidParameter.TaskStateTypeNotMath"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeTasksWithContext(ctx context.Context, request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    if request == nil {
        request = NewDescribeTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeThirdPartyAccessUserRequest() (request *DescribeThirdPartyAccessUserRequest) {
    request = &DescribeThirdPartyAccessUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeThirdPartyAccessUser")
    
    
    return
}

func NewDescribeThirdPartyAccessUserResponse() (response *DescribeThirdPartyAccessUserResponse) {
    response = &DescribeThirdPartyAccessUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeThirdPartyAccessUser
// This API is used to query the information of users visiting through the third-party platform.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeThirdPartyAccessUser(request *DescribeThirdPartyAccessUserRequest) (response *DescribeThirdPartyAccessUserResponse, err error) {
    return c.DescribeThirdPartyAccessUserWithContext(context.Background(), request)
}

// DescribeThirdPartyAccessUser
// This API is used to query the information of users visiting through the third-party platform.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeThirdPartyAccessUserWithContext(ctx context.Context, request *DescribeThirdPartyAccessUserRequest) (response *DescribeThirdPartyAccessUserResponse, err error) {
    if request == nil {
        request = NewDescribeThirdPartyAccessUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeThirdPartyAccessUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeThirdPartyAccessUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeThirdPartyAccessUserResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUpdatableDataEnginesRequest() (request *DescribeUpdatableDataEnginesRequest) {
    request = &DescribeUpdatableDataEnginesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeUpdatableDataEngines")
    
    
    return
}

func NewDescribeUpdatableDataEnginesResponse() (response *DescribeUpdatableDataEnginesResponse) {
    response = &DescribeUpdatableDataEnginesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUpdatableDataEngines
// This API is used to query the list of engines that are able to upgrade their configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAENGINEDESCRIPTION = "InvalidParameter.InvalidDataEngineDescription"
//  INVALIDPARAMETER_INVALIDENGINEEXECTYPE = "InvalidParameter.InvalidEngineExecType"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeUpdatableDataEngines(request *DescribeUpdatableDataEnginesRequest) (response *DescribeUpdatableDataEnginesResponse, err error) {
    return c.DescribeUpdatableDataEnginesWithContext(context.Background(), request)
}

// DescribeUpdatableDataEngines
// This API is used to query the list of engines that are able to upgrade their configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAENGINEDESCRIPTION = "InvalidParameter.InvalidDataEngineDescription"
//  INVALIDPARAMETER_INVALIDENGINEEXECTYPE = "InvalidParameter.InvalidEngineExecType"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeUpdatableDataEnginesWithContext(ctx context.Context, request *DescribeUpdatableDataEnginesRequest) (response *DescribeUpdatableDataEnginesResponse, err error) {
    if request == nil {
        request = NewDescribeUpdatableDataEnginesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeUpdatableDataEngines")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUpdatableDataEngines require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUpdatableDataEnginesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserDataEngineConfigRequest() (request *DescribeUserDataEngineConfigRequest) {
    request = &DescribeUserDataEngineConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeUserDataEngineConfig")
    
    
    return
}

func NewDescribeUserDataEngineConfigResponse() (response *DescribeUserDataEngineConfigResponse) {
    response = &DescribeUserDataEngineConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserDataEngineConfig
// This API is used to query user-defined engine parameters.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeUserDataEngineConfig(request *DescribeUserDataEngineConfigRequest) (response *DescribeUserDataEngineConfigResponse, err error) {
    return c.DescribeUserDataEngineConfigWithContext(context.Background(), request)
}

// DescribeUserDataEngineConfig
// This API is used to query user-defined engine parameters.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeUserDataEngineConfigWithContext(ctx context.Context, request *DescribeUserDataEngineConfigRequest) (response *DescribeUserDataEngineConfigResponse, err error) {
    if request == nil {
        request = NewDescribeUserDataEngineConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeUserDataEngineConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserDataEngineConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserDataEngineConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserInfoRequest() (request *DescribeUserInfoRequest) {
    request = &DescribeUserInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeUserInfo")
    
    
    return
}

func NewDescribeUserInfoResponse() (response *DescribeUserInfoResponse) {
    response = &DescribeUserInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserInfo
// This API is used to get detailed user information.
//
// error code that may be returned:
//  FAILEDOPERATION_GETPOLICYFAILED = "FailedOperation.GetPolicyFailed"
//  FAILEDOPERATION_GETUSERINFOFAILED = "FailedOperation.GetUserInfoFailed"
//  FAILEDOPERATION_GETWORKGROUPINFOFAILED = "FailedOperation.GetWorkGroupInfoFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDACCESSPOLICY = "InvalidParameter.InvalidAccessPolicy"
//  INVALIDPARAMETER_INVALIDFILTERKEY = "InvalidParameter.InvalidFilterKey"
//  INVALIDPARAMETER_INVALIDINFOTYPE = "InvalidParameter.InvalidInfoType"
//  INVALIDPARAMETER_INVALIDSORTBYTYPE = "InvalidParameter.InvalidSortByType"
//  INVALIDPARAMETER_INVALIDUSERNAME = "InvalidParameter.InvalidUserName"
//  INVALIDPARAMETER_INVALIDUSERTYPE = "InvalidParameter.InvalidUserType"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
func (c *Client) DescribeUserInfo(request *DescribeUserInfoRequest) (response *DescribeUserInfoResponse, err error) {
    return c.DescribeUserInfoWithContext(context.Background(), request)
}

// DescribeUserInfo
// This API is used to get detailed user information.
//
// error code that may be returned:
//  FAILEDOPERATION_GETPOLICYFAILED = "FailedOperation.GetPolicyFailed"
//  FAILEDOPERATION_GETUSERINFOFAILED = "FailedOperation.GetUserInfoFailed"
//  FAILEDOPERATION_GETWORKGROUPINFOFAILED = "FailedOperation.GetWorkGroupInfoFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDACCESSPOLICY = "InvalidParameter.InvalidAccessPolicy"
//  INVALIDPARAMETER_INVALIDFILTERKEY = "InvalidParameter.InvalidFilterKey"
//  INVALIDPARAMETER_INVALIDINFOTYPE = "InvalidParameter.InvalidInfoType"
//  INVALIDPARAMETER_INVALIDSORTBYTYPE = "InvalidParameter.InvalidSortByType"
//  INVALIDPARAMETER_INVALIDUSERNAME = "InvalidParameter.InvalidUserName"
//  INVALIDPARAMETER_INVALIDUSERTYPE = "InvalidParameter.InvalidUserType"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
func (c *Client) DescribeUserInfoWithContext(ctx context.Context, request *DescribeUserInfoRequest) (response *DescribeUserInfoResponse, err error) {
    if request == nil {
        request = NewDescribeUserInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeUserInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserRolesRequest() (request *DescribeUserRolesRequest) {
    request = &DescribeUserRolesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeUserRoles")
    
    
    return
}

func NewDescribeUserRolesResponse() (response *DescribeUserRolesResponse) {
    response = &DescribeUserRolesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserRoles
// This API is used to enumerate user roles.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeUserRoles(request *DescribeUserRolesRequest) (response *DescribeUserRolesResponse, err error) {
    return c.DescribeUserRolesWithContext(context.Background(), request)
}

// DescribeUserRoles
// This API is used to enumerate user roles.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeUserRolesWithContext(ctx context.Context, request *DescribeUserRolesRequest) (response *DescribeUserRolesResponse, err error) {
    if request == nil {
        request = NewDescribeUserRolesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeUserRoles")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserRoles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserRolesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserTypeRequest() (request *DescribeUserTypeRequest) {
    request = &DescribeUserTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeUserType")
    
    
    return
}

func NewDescribeUserTypeResponse() (response *DescribeUserTypeResponse) {
    response = &DescribeUserTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserType
// This API is used to get the types of users.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeUserType(request *DescribeUserTypeRequest) (response *DescribeUserTypeResponse, err error) {
    return c.DescribeUserTypeWithContext(context.Background(), request)
}

// DescribeUserType
// This API is used to get the types of users.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeUserTypeWithContext(ctx context.Context, request *DescribeUserTypeRequest) (response *DescribeUserTypeResponse, err error) {
    if request == nil {
        request = NewDescribeUserTypeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeUserType")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserType require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserTypeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsersRequest() (request *DescribeUsersRequest) {
    request = &DescribeUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeUsers")
    
    
    return
}

func NewDescribeUsersResponse() (response *DescribeUsersResponse) {
    response = &DescribeUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUsers
// This API is used to obtain the user list.
//
// error code that may be returned:
//  FAILEDOPERATION_GETPOLICYFAILED = "FailedOperation.GetPolicyFailed"
//  FAILEDOPERATION_GETUSERINFOFAILED = "FailedOperation.GetUserInfoFailed"
//  FAILEDOPERATION_GETWORKGROUPINFOFAILED = "FailedOperation.GetWorkGroupInfoFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_INVALIDACCESSPOLICY = "InvalidParameter.InvalidAccessPolicy"
//  INVALIDPARAMETER_INVALIDOFFSET = "InvalidParameter.InvalidOffset"
//  INVALIDPARAMETER_INVALIDSORTBYTYPE = "InvalidParameter.InvalidSortByType"
func (c *Client) DescribeUsers(request *DescribeUsersRequest) (response *DescribeUsersResponse, err error) {
    return c.DescribeUsersWithContext(context.Background(), request)
}

// DescribeUsers
// This API is used to obtain the user list.
//
// error code that may be returned:
//  FAILEDOPERATION_GETPOLICYFAILED = "FailedOperation.GetPolicyFailed"
//  FAILEDOPERATION_GETUSERINFOFAILED = "FailedOperation.GetUserInfoFailed"
//  FAILEDOPERATION_GETWORKGROUPINFOFAILED = "FailedOperation.GetWorkGroupInfoFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_INVALIDACCESSPOLICY = "InvalidParameter.InvalidAccessPolicy"
//  INVALIDPARAMETER_INVALIDOFFSET = "InvalidParameter.InvalidOffset"
//  INVALIDPARAMETER_INVALIDSORTBYTYPE = "InvalidParameter.InvalidSortByType"
func (c *Client) DescribeUsersWithContext(ctx context.Context, request *DescribeUsersRequest) (response *DescribeUsersResponse, err error) {
    if request == nil {
        request = NewDescribeUsersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeUsers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUsersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWorkGroupInfoRequest() (request *DescribeWorkGroupInfoRequest) {
    request = &DescribeWorkGroupInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeWorkGroupInfo")
    
    
    return
}

func NewDescribeWorkGroupInfoResponse() (response *DescribeWorkGroupInfoResponse) {
    response = &DescribeWorkGroupInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWorkGroupInfo
// This API is used to get detailed information about working groups.
//
// error code that may be returned:
//  FAILEDOPERATION_GETPOLICYFAILED = "FailedOperation.GetPolicyFailed"
//  FAILEDOPERATION_GETUSERINFOFAILED = "FailedOperation.GetUserInfoFailed"
//  FAILEDOPERATION_GETWORKGROUPINFOFAILED = "FailedOperation.GetWorkGroupInfoFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFILTERKEY = "InvalidParameter.InvalidFilterKey"
//  INVALIDPARAMETER_INVALIDGROUPID = "InvalidParameter.InvalidGroupId"
//  INVALIDPARAMETER_INVALIDINFOTYPE = "InvalidParameter.InvalidInfoType"
//  INVALIDPARAMETER_INVALIDSORTBYTYPE = "InvalidParameter.InvalidSortByType"
func (c *Client) DescribeWorkGroupInfo(request *DescribeWorkGroupInfoRequest) (response *DescribeWorkGroupInfoResponse, err error) {
    return c.DescribeWorkGroupInfoWithContext(context.Background(), request)
}

// DescribeWorkGroupInfo
// This API is used to get detailed information about working groups.
//
// error code that may be returned:
//  FAILEDOPERATION_GETPOLICYFAILED = "FailedOperation.GetPolicyFailed"
//  FAILEDOPERATION_GETUSERINFOFAILED = "FailedOperation.GetUserInfoFailed"
//  FAILEDOPERATION_GETWORKGROUPINFOFAILED = "FailedOperation.GetWorkGroupInfoFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFILTERKEY = "InvalidParameter.InvalidFilterKey"
//  INVALIDPARAMETER_INVALIDGROUPID = "InvalidParameter.InvalidGroupId"
//  INVALIDPARAMETER_INVALIDINFOTYPE = "InvalidParameter.InvalidInfoType"
//  INVALIDPARAMETER_INVALIDSORTBYTYPE = "InvalidParameter.InvalidSortByType"
func (c *Client) DescribeWorkGroupInfoWithContext(ctx context.Context, request *DescribeWorkGroupInfoRequest) (response *DescribeWorkGroupInfoResponse, err error) {
    if request == nil {
        request = NewDescribeWorkGroupInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeWorkGroupInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWorkGroupInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWorkGroupInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWorkGroupsRequest() (request *DescribeWorkGroupsRequest) {
    request = &DescribeWorkGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DescribeWorkGroups")
    
    
    return
}

func NewDescribeWorkGroupsResponse() (response *DescribeWorkGroupsResponse) {
    response = &DescribeWorkGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWorkGroups
// This API is used to get a list of working groups.
//
// error code that may be returned:
//  FAILEDOPERATION_GETPOLICYFAILED = "FailedOperation.GetPolicyFailed"
//  FAILEDOPERATION_GETUSERINFOFAILED = "FailedOperation.GetUserInfoFailed"
//  FAILEDOPERATION_GETWORKGROUPINFOFAILED = "FailedOperation.GetWorkGroupInfoFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_INVALIDACCESSPOLICY = "InvalidParameter.InvalidAccessPolicy"
//  INVALIDPARAMETER_INVALIDFILTERKEY = "InvalidParameter.InvalidFilterKey"
//  INVALIDPARAMETER_INVALIDOFFSET = "InvalidParameter.InvalidOffset"
//  INVALIDPARAMETER_INVALIDSORTBYTYPE = "InvalidParameter.InvalidSortByType"
//  INVALIDPARAMETER_INVALIDSORTING = "InvalidParameter.InvalidSorting"
func (c *Client) DescribeWorkGroups(request *DescribeWorkGroupsRequest) (response *DescribeWorkGroupsResponse, err error) {
    return c.DescribeWorkGroupsWithContext(context.Background(), request)
}

// DescribeWorkGroups
// This API is used to get a list of working groups.
//
// error code that may be returned:
//  FAILEDOPERATION_GETPOLICYFAILED = "FailedOperation.GetPolicyFailed"
//  FAILEDOPERATION_GETUSERINFOFAILED = "FailedOperation.GetUserInfoFailed"
//  FAILEDOPERATION_GETWORKGROUPINFOFAILED = "FailedOperation.GetWorkGroupInfoFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_INVALIDACCESSPOLICY = "InvalidParameter.InvalidAccessPolicy"
//  INVALIDPARAMETER_INVALIDFILTERKEY = "InvalidParameter.InvalidFilterKey"
//  INVALIDPARAMETER_INVALIDOFFSET = "InvalidParameter.InvalidOffset"
//  INVALIDPARAMETER_INVALIDSORTBYTYPE = "InvalidParameter.InvalidSortByType"
//  INVALIDPARAMETER_INVALIDSORTING = "InvalidParameter.InvalidSorting"
func (c *Client) DescribeWorkGroupsWithContext(ctx context.Context, request *DescribeWorkGroupsRequest) (response *DescribeWorkGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeWorkGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DescribeWorkGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWorkGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWorkGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDetachUserPolicyRequest() (request *DetachUserPolicyRequest) {
    request = &DetachUserPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DetachUserPolicy")
    
    
    return
}

func NewDetachUserPolicyResponse() (response *DetachUserPolicyResponse) {
    response = &DetachUserPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DetachUserPolicy
// This API is used to unbind the authentication policy from the user.
//
// error code that may be returned:
//  FAILEDOPERATION_GETPOLICYFAILED = "FailedOperation.GetPolicyFailed"
//  FAILEDOPERATION_GETUSERINFOFAILED = "FailedOperation.GetUserInfoFailed"
//  FAILEDOPERATION_GETWORKGROUPINFOFAILED = "FailedOperation.GetWorkGroupInfoFailed"
//  FAILEDOPERATION_GRANTPOLICYFAILED = "FailedOperation.GrantPolicyFailed"
//  FAILEDOPERATION_REVOKEPOLICYFAILED = "FailedOperation.RevokePolicyFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_INVALIDACCESSPOLICY = "InvalidParameter.InvalidAccessPolicy"
//  INVALIDPARAMETER_INVALIDUSERNAME = "InvalidParameter.InvalidUserName"
//  UNAUTHORIZEDOPERATION_GRANTPOLICY = "UnauthorizedOperation.GrantPolicy"
//  UNAUTHORIZEDOPERATION_PROHIBITEDOPERATIONADMIN = "UnauthorizedOperation.ProhibitedOperationAdmin"
//  UNAUTHORIZEDOPERATION_REVOKEPOLICY = "UnauthorizedOperation.RevokePolicy"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
//  UNSUPPORTEDOPERATION_MODIFYOWNERUNSUPPORTED = "UnsupportedOperation.ModifyOwnerUnsupported"
func (c *Client) DetachUserPolicy(request *DetachUserPolicyRequest) (response *DetachUserPolicyResponse, err error) {
    return c.DetachUserPolicyWithContext(context.Background(), request)
}

// DetachUserPolicy
// This API is used to unbind the authentication policy from the user.
//
// error code that may be returned:
//  FAILEDOPERATION_GETPOLICYFAILED = "FailedOperation.GetPolicyFailed"
//  FAILEDOPERATION_GETUSERINFOFAILED = "FailedOperation.GetUserInfoFailed"
//  FAILEDOPERATION_GETWORKGROUPINFOFAILED = "FailedOperation.GetWorkGroupInfoFailed"
//  FAILEDOPERATION_GRANTPOLICYFAILED = "FailedOperation.GrantPolicyFailed"
//  FAILEDOPERATION_REVOKEPOLICYFAILED = "FailedOperation.RevokePolicyFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_INVALIDACCESSPOLICY = "InvalidParameter.InvalidAccessPolicy"
//  INVALIDPARAMETER_INVALIDUSERNAME = "InvalidParameter.InvalidUserName"
//  UNAUTHORIZEDOPERATION_GRANTPOLICY = "UnauthorizedOperation.GrantPolicy"
//  UNAUTHORIZEDOPERATION_PROHIBITEDOPERATIONADMIN = "UnauthorizedOperation.ProhibitedOperationAdmin"
//  UNAUTHORIZEDOPERATION_REVOKEPOLICY = "UnauthorizedOperation.RevokePolicy"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
//  UNSUPPORTEDOPERATION_MODIFYOWNERUNSUPPORTED = "UnsupportedOperation.ModifyOwnerUnsupported"
func (c *Client) DetachUserPolicyWithContext(ctx context.Context, request *DetachUserPolicyRequest) (response *DetachUserPolicyResponse, err error) {
    if request == nil {
        request = NewDetachUserPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DetachUserPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetachUserPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetachUserPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDetachWorkGroupPolicyRequest() (request *DetachWorkGroupPolicyRequest) {
    request = &DetachWorkGroupPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DetachWorkGroupPolicy")
    
    
    return
}

func NewDetachWorkGroupPolicyResponse() (response *DetachWorkGroupPolicyResponse) {
    response = &DetachWorkGroupPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DetachWorkGroupPolicy
// This API is used to unbind the authentication policy from the working group.
//
// error code that may be returned:
//  FAILEDOPERATION_GETPOLICYFAILED = "FailedOperation.GetPolicyFailed"
//  FAILEDOPERATION_GETUSERINFOFAILED = "FailedOperation.GetUserInfoFailed"
//  FAILEDOPERATION_GETWORKGROUPINFOFAILED = "FailedOperation.GetWorkGroupInfoFailed"
//  FAILEDOPERATION_GRANTPOLICYFAILED = "FailedOperation.GrantPolicyFailed"
//  FAILEDOPERATION_REVOKEPOLICYFAILED = "FailedOperation.RevokePolicyFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDACCESSPOLICY = "InvalidParameter.InvalidAccessPolicy"
//  UNAUTHORIZEDOPERATION_GRANTPOLICY = "UnauthorizedOperation.GrantPolicy"
//  UNAUTHORIZEDOPERATION_REVOKEPOLICY = "UnauthorizedOperation.RevokePolicy"
func (c *Client) DetachWorkGroupPolicy(request *DetachWorkGroupPolicyRequest) (response *DetachWorkGroupPolicyResponse, err error) {
    return c.DetachWorkGroupPolicyWithContext(context.Background(), request)
}

// DetachWorkGroupPolicy
// This API is used to unbind the authentication policy from the working group.
//
// error code that may be returned:
//  FAILEDOPERATION_GETPOLICYFAILED = "FailedOperation.GetPolicyFailed"
//  FAILEDOPERATION_GETUSERINFOFAILED = "FailedOperation.GetUserInfoFailed"
//  FAILEDOPERATION_GETWORKGROUPINFOFAILED = "FailedOperation.GetWorkGroupInfoFailed"
//  FAILEDOPERATION_GRANTPOLICYFAILED = "FailedOperation.GrantPolicyFailed"
//  FAILEDOPERATION_REVOKEPOLICYFAILED = "FailedOperation.RevokePolicyFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDACCESSPOLICY = "InvalidParameter.InvalidAccessPolicy"
//  UNAUTHORIZEDOPERATION_GRANTPOLICY = "UnauthorizedOperation.GrantPolicy"
//  UNAUTHORIZEDOPERATION_REVOKEPOLICY = "UnauthorizedOperation.RevokePolicy"
func (c *Client) DetachWorkGroupPolicyWithContext(ctx context.Context, request *DetachWorkGroupPolicyRequest) (response *DetachWorkGroupPolicyResponse, err error) {
    if request == nil {
        request = NewDetachWorkGroupPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DetachWorkGroupPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetachWorkGroupPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetachWorkGroupPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDropDLCTableRequest() (request *DropDLCTableRequest) {
    request = &DropDLCTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DropDLCTable")
    
    
    return
}

func NewDropDLCTableResponse() (response *DropDLCTableResponse) {
    response = &DropDLCTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DropDLCTable
// This API is used to delete the table.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DropDLCTable(request *DropDLCTableRequest) (response *DropDLCTableResponse, err error) {
    return c.DropDLCTableWithContext(context.Background(), request)
}

// DropDLCTable
// This API is used to delete the table.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DropDLCTableWithContext(ctx context.Context, request *DropDLCTableRequest) (response *DropDLCTableResponse, err error) {
    if request == nil {
        request = NewDropDLCTableRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DropDLCTable")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DropDLCTable require credential")
    }

    request.SetContext(ctx)
    
    response = NewDropDLCTableResponse()
    err = c.Send(request, response)
    return
}

func NewDropDMSDatabaseRequest() (request *DropDMSDatabaseRequest) {
    request = &DropDMSDatabaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DropDMSDatabase")
    
    
    return
}

func NewDropDMSDatabaseResponse() (response *DropDMSDatabaseResponse) {
    response = &DropDMSDatabaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DropDMSDatabase
// This API is used to delete databases in the DMS metadata module.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DropDMSDatabase(request *DropDMSDatabaseRequest) (response *DropDMSDatabaseResponse, err error) {
    return c.DropDMSDatabaseWithContext(context.Background(), request)
}

// DropDMSDatabase
// This API is used to delete databases in the DMS metadata module.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DropDMSDatabaseWithContext(ctx context.Context, request *DropDMSDatabaseRequest) (response *DropDMSDatabaseResponse, err error) {
    if request == nil {
        request = NewDropDMSDatabaseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DropDMSDatabase")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DropDMSDatabase require credential")
    }

    request.SetContext(ctx)
    
    response = NewDropDMSDatabaseResponse()
    err = c.Send(request, response)
    return
}

func NewDropDMSTableRequest() (request *DropDMSTableRequest) {
    request = &DropDMSTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "DropDMSTable")
    
    
    return
}

func NewDropDMSTableResponse() (response *DropDMSTableResponse) {
    response = &DropDMSTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DropDMSTable
// This API is used to delete tables in the DMS metadata module.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DropDMSTable(request *DropDMSTableRequest) (response *DropDMSTableResponse, err error) {
    return c.DropDMSTableWithContext(context.Background(), request)
}

// DropDMSTable
// This API is used to delete tables in the DMS metadata module.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DropDMSTableWithContext(ctx context.Context, request *DropDMSTableRequest) (response *DropDMSTableResponse, err error) {
    if request == nil {
        request = NewDropDMSTableRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "DropDMSTable")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DropDMSTable require credential")
    }

    request.SetContext(ctx)
    
    response = NewDropDMSTableResponse()
    err = c.Send(request, response)
    return
}

func NewGenerateCreateMangedTableSqlRequest() (request *GenerateCreateMangedTableSqlRequest) {
    request = &GenerateCreateMangedTableSqlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "GenerateCreateMangedTableSql")
    
    
    return
}

func NewGenerateCreateMangedTableSqlResponse() (response *GenerateCreateMangedTableSqlResponse) {
    response = &GenerateCreateMangedTableSqlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GenerateCreateMangedTableSql
// This API is used to generate SQL statements for creating a managed table.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCOLUMNNAMELENGTH = "InvalidParameter.InvalidColumnNameLength"
//  INVALIDPARAMETER_INVALIDCOLUMNNUMBER = "InvalidParameter.InvalidColumnNumber"
//  INVALIDPARAMETER_INVALIDDECIMALTYPE = "InvalidParameter.InvalidDecimalType"
//  INVALIDPARAMETER_INVALIDTABLENAMELENGTH = "InvalidParameter.InvalidTableNameLength"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GenerateCreateMangedTableSql(request *GenerateCreateMangedTableSqlRequest) (response *GenerateCreateMangedTableSqlResponse, err error) {
    return c.GenerateCreateMangedTableSqlWithContext(context.Background(), request)
}

// GenerateCreateMangedTableSql
// This API is used to generate SQL statements for creating a managed table.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCOLUMNNAMELENGTH = "InvalidParameter.InvalidColumnNameLength"
//  INVALIDPARAMETER_INVALIDCOLUMNNUMBER = "InvalidParameter.InvalidColumnNumber"
//  INVALIDPARAMETER_INVALIDDECIMALTYPE = "InvalidParameter.InvalidDecimalType"
//  INVALIDPARAMETER_INVALIDTABLENAMELENGTH = "InvalidParameter.InvalidTableNameLength"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GenerateCreateMangedTableSqlWithContext(ctx context.Context, request *GenerateCreateMangedTableSqlRequest) (response *GenerateCreateMangedTableSqlResponse, err error) {
    if request == nil {
        request = NewGenerateCreateMangedTableSqlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "GenerateCreateMangedTableSql")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GenerateCreateMangedTableSql require credential")
    }

    request.SetContext(ctx)
    
    response = NewGenerateCreateMangedTableSqlResponse()
    err = c.Send(request, response)
    return
}

func NewGetOptimizerPolicyRequest() (request *GetOptimizerPolicyRequest) {
    request = &GetOptimizerPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "GetOptimizerPolicy")
    
    
    return
}

func NewGetOptimizerPolicyResponse() (response *GetOptimizerPolicyResponse) {
    response = &GetOptimizerPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetOptimizerPolicy
// GetOptimizerPolicy
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetOptimizerPolicy(request *GetOptimizerPolicyRequest) (response *GetOptimizerPolicyResponse, err error) {
    return c.GetOptimizerPolicyWithContext(context.Background(), request)
}

// GetOptimizerPolicy
// GetOptimizerPolicy
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetOptimizerPolicyWithContext(ctx context.Context, request *GetOptimizerPolicyRequest) (response *GetOptimizerPolicyResponse, err error) {
    if request == nil {
        request = NewGetOptimizerPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "GetOptimizerPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetOptimizerPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetOptimizerPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewGrantDLCCatalogAccessRequest() (request *GrantDLCCatalogAccessRequest) {
    request = &GrantDLCCatalogAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "GrantDLCCatalogAccess")
    
    
    return
}

func NewGrantDLCCatalogAccessResponse() (response *GrantDLCCatalogAccessResponse) {
    response = &GrantDLCCatalogAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GrantDLCCatalogAccess
// This API is used to grant permissions for visiting DLC Catalog.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GrantDLCCatalogAccess(request *GrantDLCCatalogAccessRequest) (response *GrantDLCCatalogAccessResponse, err error) {
    return c.GrantDLCCatalogAccessWithContext(context.Background(), request)
}

// GrantDLCCatalogAccess
// This API is used to grant permissions for visiting DLC Catalog.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GrantDLCCatalogAccessWithContext(ctx context.Context, request *GrantDLCCatalogAccessRequest) (response *GrantDLCCatalogAccessResponse, err error) {
    if request == nil {
        request = NewGrantDLCCatalogAccessRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "GrantDLCCatalogAccess")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GrantDLCCatalogAccess require credential")
    }

    request.SetContext(ctx)
    
    response = NewGrantDLCCatalogAccessResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAdvancedStoreLocationRequest() (request *ModifyAdvancedStoreLocationRequest) {
    request = &ModifyAdvancedStoreLocationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "ModifyAdvancedStoreLocation")
    
    
    return
}

func NewModifyAdvancedStoreLocationResponse() (response *ModifyAdvancedStoreLocationResponse) {
    response = &ModifyAdvancedStoreLocationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAdvancedStoreLocation
// This API is used to modify the advanced settings of the SQL query interface.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSTORELOCATION = "InvalidParameter.InvalidStoreLocation"
func (c *Client) ModifyAdvancedStoreLocation(request *ModifyAdvancedStoreLocationRequest) (response *ModifyAdvancedStoreLocationResponse, err error) {
    return c.ModifyAdvancedStoreLocationWithContext(context.Background(), request)
}

// ModifyAdvancedStoreLocation
// This API is used to modify the advanced settings of the SQL query interface.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSTORELOCATION = "InvalidParameter.InvalidStoreLocation"
func (c *Client) ModifyAdvancedStoreLocationWithContext(ctx context.Context, request *ModifyAdvancedStoreLocationRequest) (response *ModifyAdvancedStoreLocationResponse, err error) {
    if request == nil {
        request = NewModifyAdvancedStoreLocationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "ModifyAdvancedStoreLocation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAdvancedStoreLocation require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAdvancedStoreLocationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDataEngineDescriptionRequest() (request *ModifyDataEngineDescriptionRequest) {
    request = &ModifyDataEngineDescriptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "ModifyDataEngineDescription")
    
    
    return
}

func NewModifyDataEngineDescriptionResponse() (response *ModifyDataEngineDescriptionResponse) {
    response = &ModifyDataEngineDescriptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDataEngineDescription
// This API is used to modify the engine's description.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INVALIDDATAENGINEDESCRIPTION = "InvalidParameter.InvalidDataEngineDescription"
//  INVALIDPARAMETER_INVALIDDATAENGINENAME = "InvalidParameter.InvalidDataEngineName"
//  UNAUTHORIZEDOPERATION_MODIFYCOMPUTINGENGINE = "UnauthorizedOperation.ModifyComputingEngine"
func (c *Client) ModifyDataEngineDescription(request *ModifyDataEngineDescriptionRequest) (response *ModifyDataEngineDescriptionResponse, err error) {
    return c.ModifyDataEngineDescriptionWithContext(context.Background(), request)
}

// ModifyDataEngineDescription
// This API is used to modify the engine's description.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INVALIDDATAENGINEDESCRIPTION = "InvalidParameter.InvalidDataEngineDescription"
//  INVALIDPARAMETER_INVALIDDATAENGINENAME = "InvalidParameter.InvalidDataEngineName"
//  UNAUTHORIZEDOPERATION_MODIFYCOMPUTINGENGINE = "UnauthorizedOperation.ModifyComputingEngine"
func (c *Client) ModifyDataEngineDescriptionWithContext(ctx context.Context, request *ModifyDataEngineDescriptionRequest) (response *ModifyDataEngineDescriptionResponse, err error) {
    if request == nil {
        request = NewModifyDataEngineDescriptionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "ModifyDataEngineDescription")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDataEngineDescription require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDataEngineDescriptionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyGovernEventRuleRequest() (request *ModifyGovernEventRuleRequest) {
    request = &ModifyGovernEventRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "ModifyGovernEventRule")
    
    
    return
}

func NewModifyGovernEventRuleResponse() (response *ModifyGovernEventRuleResponse) {
    response = &ModifyGovernEventRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyGovernEventRule
// This API is used to change data governance event thresholds.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyGovernEventRule(request *ModifyGovernEventRuleRequest) (response *ModifyGovernEventRuleResponse, err error) {
    return c.ModifyGovernEventRuleWithContext(context.Background(), request)
}

// ModifyGovernEventRule
// This API is used to change data governance event thresholds.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyGovernEventRuleWithContext(ctx context.Context, request *ModifyGovernEventRuleRequest) (response *ModifyGovernEventRuleResponse, err error) {
    if request == nil {
        request = NewModifyGovernEventRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "ModifyGovernEventRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyGovernEventRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyGovernEventRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifySparkAppRequest() (request *ModifySparkAppRequest) {
    request = &ModifySparkAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "ModifySparkApp")
    
    
    return
}

func NewModifySparkAppResponse() (response *ModifySparkAppResponse) {
    response = &ModifySparkAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySparkApp
// This API is used to update a Spark job.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_INVALIDAPPFILEFORMAT = "InvalidParameter.InvalidAppFileFormat"
//  INVALIDPARAMETER_INVALIDDATAENGINENAME = "InvalidParameter.InvalidDataEngineName"
//  INVALIDPARAMETER_INVALIDDRIVERSIZE = "InvalidParameter.InvalidDriverSize"
//  INVALIDPARAMETER_INVALIDEXECUTORSIZE = "InvalidParameter.InvalidExecutorSize"
//  INVALIDPARAMETER_INVALIDFILECOMPRESSIONFORMAT = "InvalidParameter.InvalidFileCompressionFormat"
//  INVALIDPARAMETER_INVALIDFILEPATHFORMAT = "InvalidParameter.InvalidFilePathFormat"
//  INVALIDPARAMETER_SQLBASE64DECODEFAIL = "InvalidParameter.SQLBase64DecodeFail"
//  INVALIDPARAMETER_SPARKJOBNOTFOUND = "InvalidParameter.SparkJobNotFound"
//  INVALIDPARAMETER_SPARKJOBONLYSUPPORTSPARKBATCHENGINE = "InvalidParameter.SparkJobOnlySupportSparkBatchEngine"
//  RESOURCEINSUFFICIENT_SPARKJOBINSUFFICIENTRESOURCES = "ResourceInsufficient.SparkJobInsufficientResources"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
func (c *Client) ModifySparkApp(request *ModifySparkAppRequest) (response *ModifySparkAppResponse, err error) {
    return c.ModifySparkAppWithContext(context.Background(), request)
}

// ModifySparkApp
// This API is used to update a Spark job.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_INVALIDAPPFILEFORMAT = "InvalidParameter.InvalidAppFileFormat"
//  INVALIDPARAMETER_INVALIDDATAENGINENAME = "InvalidParameter.InvalidDataEngineName"
//  INVALIDPARAMETER_INVALIDDRIVERSIZE = "InvalidParameter.InvalidDriverSize"
//  INVALIDPARAMETER_INVALIDEXECUTORSIZE = "InvalidParameter.InvalidExecutorSize"
//  INVALIDPARAMETER_INVALIDFILECOMPRESSIONFORMAT = "InvalidParameter.InvalidFileCompressionFormat"
//  INVALIDPARAMETER_INVALIDFILEPATHFORMAT = "InvalidParameter.InvalidFilePathFormat"
//  INVALIDPARAMETER_SQLBASE64DECODEFAIL = "InvalidParameter.SQLBase64DecodeFail"
//  INVALIDPARAMETER_SPARKJOBNOTFOUND = "InvalidParameter.SparkJobNotFound"
//  INVALIDPARAMETER_SPARKJOBONLYSUPPORTSPARKBATCHENGINE = "InvalidParameter.SparkJobOnlySupportSparkBatchEngine"
//  RESOURCEINSUFFICIENT_SPARKJOBINSUFFICIENTRESOURCES = "ResourceInsufficient.SparkJobInsufficientResources"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
func (c *Client) ModifySparkAppWithContext(ctx context.Context, request *ModifySparkAppRequest) (response *ModifySparkAppResponse, err error) {
    if request == nil {
        request = NewModifySparkAppRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "ModifySparkApp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySparkApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySparkAppResponse()
    err = c.Send(request, response)
    return
}

func NewModifySparkAppBatchRequest() (request *ModifySparkAppBatchRequest) {
    request = &ModifySparkAppBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "ModifySparkAppBatch")
    
    
    return
}

func NewModifySparkAppBatchResponse() (response *ModifySparkAppBatchResponse) {
    response = &ModifySparkAppBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySparkAppBatch
// This API is used to modify Spark job parameters in batches.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  INVALIDPARAMETER_SPARKJOBISINHERITTYPENOTMATCH = "InvalidParameter.SparkJobIsInheritTypeNotMatch"
//  RESOURCEINSUFFICIENT_SPARKJOBINSUFFICIENTRESOURCES = "ResourceInsufficient.SparkJobInsufficientResources"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTUNIQUE = "ResourceNotFound.DataEngineNotUnique"
func (c *Client) ModifySparkAppBatch(request *ModifySparkAppBatchRequest) (response *ModifySparkAppBatchResponse, err error) {
    return c.ModifySparkAppBatchWithContext(context.Background(), request)
}

// ModifySparkAppBatch
// This API is used to modify Spark job parameters in batches.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  INVALIDPARAMETER_SPARKJOBISINHERITTYPENOTMATCH = "InvalidParameter.SparkJobIsInheritTypeNotMatch"
//  RESOURCEINSUFFICIENT_SPARKJOBINSUFFICIENTRESOURCES = "ResourceInsufficient.SparkJobInsufficientResources"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTUNIQUE = "ResourceNotFound.DataEngineNotUnique"
func (c *Client) ModifySparkAppBatchWithContext(ctx context.Context, request *ModifySparkAppBatchRequest) (response *ModifySparkAppBatchResponse, err error) {
    if request == nil {
        request = NewModifySparkAppBatchRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "ModifySparkAppBatch")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySparkAppBatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySparkAppBatchResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserRequest() (request *ModifyUserRequest) {
    request = &ModifyUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "ModifyUser")
    
    
    return
}

func NewModifyUserResponse() (response *ModifyUserResponse) {
    response = &ModifyUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUser
// This API is used to modify user information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCESSOLDOUT_UNAUTHORIZEDOPERATION = "ResourcesSoldOut.UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_MODIFYUSERINFO = "UnauthorizedOperation.ModifyUserInfo"
//  UNSUPPORTEDOPERATION_MODIFYOWNERUNSUPPORTED = "UnsupportedOperation.ModifyOwnerUnsupported"
func (c *Client) ModifyUser(request *ModifyUserRequest) (response *ModifyUserResponse, err error) {
    return c.ModifyUserWithContext(context.Background(), request)
}

// ModifyUser
// This API is used to modify user information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCESSOLDOUT_UNAUTHORIZEDOPERATION = "ResourcesSoldOut.UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_MODIFYUSERINFO = "UnauthorizedOperation.ModifyUserInfo"
//  UNSUPPORTEDOPERATION_MODIFYOWNERUNSUPPORTED = "UnsupportedOperation.ModifyOwnerUnsupported"
func (c *Client) ModifyUserWithContext(ctx context.Context, request *ModifyUserRequest) (response *ModifyUserResponse, err error) {
    if request == nil {
        request = NewModifyUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "ModifyUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserTypeRequest() (request *ModifyUserTypeRequest) {
    request = &ModifyUserTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "ModifyUserType")
    
    
    return
}

func NewModifyUserTypeResponse() (response *ModifyUserTypeResponse) {
    response = &ModifyUserTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUserType
// This API is used to modify the types of users. Only users who are also administrators can call this API to conduct the operation.
//
// error code that may be returned:
//  FAILEDOPERATION_GETUSERINFOFAILED = "FailedOperation.GetUserInfoFailed"
//  FAILEDOPERATION_GRANTPOLICYFAILED = "FailedOperation.GrantPolicyFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDACCESSPOLICY = "InvalidParameter.InvalidAccessPolicy"
//  INVALIDPARAMETER_INVALIDUSERNAME = "InvalidParameter.InvalidUserName"
//  INVALIDPARAMETER_INVALIDUSERTYPE = "InvalidParameter.InvalidUserType"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_GRANTPOLICY = "UnauthorizedOperation.GrantPolicy"
//  UNAUTHORIZEDOPERATION_MODIFYUSERTYPE = "UnauthorizedOperation.ModifyUserType"
//  UNAUTHORIZEDOPERATION_REVOKEPOLICY = "UnauthorizedOperation.RevokePolicy"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
//  UNSUPPORTEDOPERATION_MODIFYOWNERUNSUPPORTED = "UnsupportedOperation.ModifyOwnerUnsupported"
func (c *Client) ModifyUserType(request *ModifyUserTypeRequest) (response *ModifyUserTypeResponse, err error) {
    return c.ModifyUserTypeWithContext(context.Background(), request)
}

// ModifyUserType
// This API is used to modify the types of users. Only users who are also administrators can call this API to conduct the operation.
//
// error code that may be returned:
//  FAILEDOPERATION_GETUSERINFOFAILED = "FailedOperation.GetUserInfoFailed"
//  FAILEDOPERATION_GRANTPOLICYFAILED = "FailedOperation.GrantPolicyFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDACCESSPOLICY = "InvalidParameter.InvalidAccessPolicy"
//  INVALIDPARAMETER_INVALIDUSERNAME = "InvalidParameter.InvalidUserName"
//  INVALIDPARAMETER_INVALIDUSERTYPE = "InvalidParameter.InvalidUserType"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_GRANTPOLICY = "UnauthorizedOperation.GrantPolicy"
//  UNAUTHORIZEDOPERATION_MODIFYUSERTYPE = "UnauthorizedOperation.ModifyUserType"
//  UNAUTHORIZEDOPERATION_REVOKEPOLICY = "UnauthorizedOperation.RevokePolicy"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
//  UNSUPPORTEDOPERATION_MODIFYOWNERUNSUPPORTED = "UnsupportedOperation.ModifyOwnerUnsupported"
func (c *Client) ModifyUserTypeWithContext(ctx context.Context, request *ModifyUserTypeRequest) (response *ModifyUserTypeResponse, err error) {
    if request == nil {
        request = NewModifyUserTypeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "ModifyUserType")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUserType require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserTypeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWorkGroupRequest() (request *ModifyWorkGroupRequest) {
    request = &ModifyWorkGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "ModifyWorkGroup")
    
    
    return
}

func NewModifyWorkGroupResponse() (response *ModifyWorkGroupResponse) {
    response = &ModifyWorkGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyWorkGroup
// This API is used to modify working group information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDGROUPID = "InvalidParameter.InvalidGroupId"
//  UNAUTHORIZEDOPERATION_MODIFYWORKGROUPINFO = "UnauthorizedOperation.ModifyWorkgroupInfo"
func (c *Client) ModifyWorkGroup(request *ModifyWorkGroupRequest) (response *ModifyWorkGroupResponse, err error) {
    return c.ModifyWorkGroupWithContext(context.Background(), request)
}

// ModifyWorkGroup
// This API is used to modify working group information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDGROUPID = "InvalidParameter.InvalidGroupId"
//  UNAUTHORIZEDOPERATION_MODIFYWORKGROUPINFO = "UnauthorizedOperation.ModifyWorkgroupInfo"
func (c *Client) ModifyWorkGroupWithContext(ctx context.Context, request *ModifyWorkGroupRequest) (response *ModifyWorkGroupResponse, err error) {
    if request == nil {
        request = NewModifyWorkGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "ModifyWorkGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWorkGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWorkGroupResponse()
    err = c.Send(request, response)
    return
}

func NewQueryResultRequest() (request *QueryResultRequest) {
    request = &QueryResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "QueryResult")
    
    
    return
}

func NewQueryResultResponse() (response *QueryResultResponse) {
    response = &QueryResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryResult
// This API is used to query the result of obtaining tasks.
//
// error code that may be returned:
//  FAILEDOPERATION_HTTPCLIENTDOREQUESTFAILED = "FailedOperation.HttpClientDoRequestFailed"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
func (c *Client) QueryResult(request *QueryResultRequest) (response *QueryResultResponse, err error) {
    return c.QueryResultWithContext(context.Background(), request)
}

// QueryResult
// This API is used to query the result of obtaining tasks.
//
// error code that may be returned:
//  FAILEDOPERATION_HTTPCLIENTDOREQUESTFAILED = "FailedOperation.HttpClientDoRequestFailed"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
func (c *Client) QueryResultWithContext(ctx context.Context, request *QueryResultRequest) (response *QueryResultResponse, err error) {
    if request == nil {
        request = NewQueryResultRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "QueryResult")
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryResultResponse()
    err = c.Send(request, response)
    return
}

func NewQueryTaskCostDetailRequest() (request *QueryTaskCostDetailRequest) {
    request = &QueryTaskCostDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "QueryTaskCostDetail")
    
    
    return
}

func NewQueryTaskCostDetailResponse() (response *QueryTaskCostDetailResponse) {
    response = &QueryTaskCostDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryTaskCostDetail
// This API is used to query task consumption details.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FILTERSVALUESNUMBEROUTOFLIMIT = "InvalidParameter.FiltersValuesNumberOutOfLimit"
//  INVALIDPARAMETER_INVALIDFILTERLENGTH = "InvalidParameter.InvalidFilterLength"
//  INVALIDPARAMETER_INVALIDTIMEFORMAT = "InvalidParameter.InvalidTimeFormat"
//  INVALIDPARAMETER_PARAMETERBASE64DECODEFAILED = "InvalidParameter.ParameterBase64DecodeFailed"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  INVALIDPARAMETER_SQLTASKFILTERSKEYTYPENOTMATH = "InvalidParameter.SQLTaskFiltersKeyTypeNotMath"
//  INVALIDPARAMETER_SQLTASKNOTFOUND = "InvalidParameter.SQLTaskNotFound"
//  INVALIDPARAMETER_SQLTASKSORTBYTYPENOTMATCH = "InvalidParameter.SQLTaskSortByTypeNotMatch"
//  INVALIDPARAMETER_SPARKJOBNOTFOUND = "InvalidParameter.SparkJobNotFound"
//  INVALIDPARAMETER_TASKSTATETYPENOTMATH = "InvalidParameter.TaskStateTypeNotMath"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) QueryTaskCostDetail(request *QueryTaskCostDetailRequest) (response *QueryTaskCostDetailResponse, err error) {
    return c.QueryTaskCostDetailWithContext(context.Background(), request)
}

// QueryTaskCostDetail
// This API is used to query task consumption details.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FILTERSVALUESNUMBEROUTOFLIMIT = "InvalidParameter.FiltersValuesNumberOutOfLimit"
//  INVALIDPARAMETER_INVALIDFILTERLENGTH = "InvalidParameter.InvalidFilterLength"
//  INVALIDPARAMETER_INVALIDTIMEFORMAT = "InvalidParameter.InvalidTimeFormat"
//  INVALIDPARAMETER_PARAMETERBASE64DECODEFAILED = "InvalidParameter.ParameterBase64DecodeFailed"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  INVALIDPARAMETER_SQLTASKFILTERSKEYTYPENOTMATH = "InvalidParameter.SQLTaskFiltersKeyTypeNotMath"
//  INVALIDPARAMETER_SQLTASKNOTFOUND = "InvalidParameter.SQLTaskNotFound"
//  INVALIDPARAMETER_SQLTASKSORTBYTYPENOTMATCH = "InvalidParameter.SQLTaskSortByTypeNotMatch"
//  INVALIDPARAMETER_SPARKJOBNOTFOUND = "InvalidParameter.SparkJobNotFound"
//  INVALIDPARAMETER_TASKSTATETYPENOTMATH = "InvalidParameter.TaskStateTypeNotMath"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) QueryTaskCostDetailWithContext(ctx context.Context, request *QueryTaskCostDetailRequest) (response *QueryTaskCostDetailResponse, err error) {
    if request == nil {
        request = NewQueryTaskCostDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "QueryTaskCostDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryTaskCostDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryTaskCostDetailResponse()
    err = c.Send(request, response)
    return
}

func NewRegisterThirdPartyAccessUserRequest() (request *RegisterThirdPartyAccessUserRequest) {
    request = &RegisterThirdPartyAccessUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "RegisterThirdPartyAccessUser")
    
    
    return
}

func NewRegisterThirdPartyAccessUserResponse() (response *RegisterThirdPartyAccessUserResponse) {
    response = &RegisterThirdPartyAccessUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RegisterThirdPartyAccessUser
// This API is used to enable visits to the third-party platform.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) RegisterThirdPartyAccessUser(request *RegisterThirdPartyAccessUserRequest) (response *RegisterThirdPartyAccessUserResponse, err error) {
    return c.RegisterThirdPartyAccessUserWithContext(context.Background(), request)
}

// RegisterThirdPartyAccessUser
// This API is used to enable visits to the third-party platform.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) RegisterThirdPartyAccessUserWithContext(ctx context.Context, request *RegisterThirdPartyAccessUserRequest) (response *RegisterThirdPartyAccessUserResponse, err error) {
    if request == nil {
        request = NewRegisterThirdPartyAccessUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "RegisterThirdPartyAccessUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RegisterThirdPartyAccessUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewRegisterThirdPartyAccessUserResponse()
    err = c.Send(request, response)
    return
}

func NewRenewDataEngineRequest() (request *RenewDataEngineRequest) {
    request = &RenewDataEngineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "RenewDataEngine")
    
    
    return
}

func NewRenewDataEngineResponse() (response *RenewDataEngineResponse) {
    response = &RenewDataEngineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RenewDataEngine
// This API is used to renew the data engine.
//
// error code that may be returned:
//  INVALIDPARAMETER_INSTANCEINPROCESS = "InvalidParameter.InstanceInProcess"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
//  UNAUTHORIZEDOPERATION_RENEWCOMPUTINGENGINE = "UnauthorizedOperation.RenewComputingEngine"
func (c *Client) RenewDataEngine(request *RenewDataEngineRequest) (response *RenewDataEngineResponse, err error) {
    return c.RenewDataEngineWithContext(context.Background(), request)
}

// RenewDataEngine
// This API is used to renew the data engine.
//
// error code that may be returned:
//  INVALIDPARAMETER_INSTANCEINPROCESS = "InvalidParameter.InstanceInProcess"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
//  UNAUTHORIZEDOPERATION_RENEWCOMPUTINGENGINE = "UnauthorizedOperation.RenewComputingEngine"
func (c *Client) RenewDataEngineWithContext(ctx context.Context, request *RenewDataEngineRequest) (response *RenewDataEngineResponse, err error) {
    if request == nil {
        request = NewRenewDataEngineRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "RenewDataEngine")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RenewDataEngine require credential")
    }

    request.SetContext(ctx)
    
    response = NewRenewDataEngineResponse()
    err = c.Send(request, response)
    return
}

func NewRestartDataEngineRequest() (request *RestartDataEngineRequest) {
    request = &RestartDataEngineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "RestartDataEngine")
    
    
    return
}

func NewRestartDataEngineResponse() (response *RestartDataEngineResponse) {
    response = &RestartDataEngineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RestartDataEngine
// This API is used to restart engines.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOPERMISSIONTOUSETHEDATAENGINE = "FailedOperation.NoPermissionToUseTheDataEngine"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  RESOURCENOTFOUND_DATAENGINENOTACTIVITY = "ResourceNotFound.DataEngineNotActivity"
//  RESOURCENOTFOUND_DATAENGINENOTRUNNING = "ResourceNotFound.DataEngineNotRunning"
//  RESOURCENOTFOUND_DEFAULTDATAENGINENOTFOUND = "ResourceNotFound.DefaultDataEngineNotFound"
func (c *Client) RestartDataEngine(request *RestartDataEngineRequest) (response *RestartDataEngineResponse, err error) {
    return c.RestartDataEngineWithContext(context.Background(), request)
}

// RestartDataEngine
// This API is used to restart engines.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOPERMISSIONTOUSETHEDATAENGINE = "FailedOperation.NoPermissionToUseTheDataEngine"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  RESOURCENOTFOUND_DATAENGINENOTACTIVITY = "ResourceNotFound.DataEngineNotActivity"
//  RESOURCENOTFOUND_DATAENGINENOTRUNNING = "ResourceNotFound.DataEngineNotRunning"
//  RESOURCENOTFOUND_DEFAULTDATAENGINENOTFOUND = "ResourceNotFound.DefaultDataEngineNotFound"
func (c *Client) RestartDataEngineWithContext(ctx context.Context, request *RestartDataEngineRequest) (response *RestartDataEngineResponse, err error) {
    if request == nil {
        request = NewRestartDataEngineRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "RestartDataEngine")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartDataEngine require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestartDataEngineResponse()
    err = c.Send(request, response)
    return
}

func NewRevokeDLCCatalogAccessRequest() (request *RevokeDLCCatalogAccessRequest) {
    request = &RevokeDLCCatalogAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "RevokeDLCCatalogAccess")
    
    
    return
}

func NewRevokeDLCCatalogAccessResponse() (response *RevokeDLCCatalogAccessResponse) {
    response = &RevokeDLCCatalogAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RevokeDLCCatalogAccess
// This API is used to revoke permissions for visiting DLC Catalog.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOPERMISSIONTOUSETHEDATAENGINE = "FailedOperation.NoPermissionToUseTheDataEngine"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  RESOURCENOTFOUND_DATAENGINENOTACTIVITY = "ResourceNotFound.DataEngineNotActivity"
//  RESOURCENOTFOUND_DATAENGINENOTRUNNING = "ResourceNotFound.DataEngineNotRunning"
//  RESOURCENOTFOUND_DEFAULTDATAENGINENOTFOUND = "ResourceNotFound.DefaultDataEngineNotFound"
func (c *Client) RevokeDLCCatalogAccess(request *RevokeDLCCatalogAccessRequest) (response *RevokeDLCCatalogAccessResponse, err error) {
    return c.RevokeDLCCatalogAccessWithContext(context.Background(), request)
}

// RevokeDLCCatalogAccess
// This API is used to revoke permissions for visiting DLC Catalog.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOPERMISSIONTOUSETHEDATAENGINE = "FailedOperation.NoPermissionToUseTheDataEngine"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  RESOURCENOTFOUND_DATAENGINENOTACTIVITY = "ResourceNotFound.DataEngineNotActivity"
//  RESOURCENOTFOUND_DATAENGINENOTRUNNING = "ResourceNotFound.DataEngineNotRunning"
//  RESOURCENOTFOUND_DEFAULTDATAENGINENOTFOUND = "ResourceNotFound.DefaultDataEngineNotFound"
func (c *Client) RevokeDLCCatalogAccessWithContext(ctx context.Context, request *RevokeDLCCatalogAccessRequest) (response *RevokeDLCCatalogAccessResponse, err error) {
    if request == nil {
        request = NewRevokeDLCCatalogAccessRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "RevokeDLCCatalogAccess")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RevokeDLCCatalogAccess require credential")
    }

    request.SetContext(ctx)
    
    response = NewRevokeDLCCatalogAccessResponse()
    err = c.Send(request, response)
    return
}

func NewRollbackDataEngineImageRequest() (request *RollbackDataEngineImageRequest) {
    request = &RollbackDataEngineImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "RollbackDataEngineImage")
    
    
    return
}

func NewRollbackDataEngineImageResponse() (response *RollbackDataEngineImageResponse) {
    response = &RollbackDataEngineImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RollbackDataEngineImage
// This API is used to roll back the versions of the engine image.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOPERMISSIONTOUSETHEDATAENGINE = "FailedOperation.NoPermissionToUseTheDataEngine"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_DATAENGINEEXECTYPENOTMATCH = "InvalidParameter.DataEngineExecTypeNotMatch"
//  INVALIDPARAMETER_DATAENGINEIMAGEOPERATENOTMATCH = "InvalidParameter.DataEngineImageOperateNotMatch"
//  INVALIDPARAMETER_DATAENGINEPAYMODETYPENOTMATCH = "InvalidParameter.DataEnginePayModeTypeNotMatch"
//  INVALIDPARAMETER_DATAENGINETYPENOTMATCH = "InvalidParameter.DataEngineTypeNotMatch"
//  INVALIDPARAMETER_IMAGECLUSTERPARAMETERSFORMATNOTJSON = "InvalidParameter.ImageClusterParametersFormatNotJson"
//  INVALIDPARAMETER_IMAGEENGINETYPENOTMATCH = "InvalidParameter.ImageEngineTypeNotMatch"
//  INVALIDPARAMETER_IMAGEISPUBLICNOTMATCH = "InvalidParameter.ImageIsPublicNotMatch"
//  INVALIDPARAMETER_IMAGEPARAMETERNOTFOUND = "InvalidParameter.ImageParameterNotFound"
//  INVALIDPARAMETER_IMAGESESSIONPARAMETERSFORMATNOTJSON = "InvalidParameter.ImageSessionParametersFormatNotJson"
//  INVALIDPARAMETER_IMAGESTATENOTMATCH = "InvalidParameter.ImageStateNotMatch"
//  INVALIDPARAMETER_IMAGEUSERRECORDSTYPENOTMATCH = "InvalidParameter.ImageUserRecordsTypeNotMatch"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  RESOURCENOTFOUND_DATAENGINECONFIGINSTANCENOTFOUND = "ResourceNotFound.DataEngineConfigInstanceNotFound"
//  RESOURCENOTFOUND_DATAENGINECONFIGINSTANCENOTUNIQUE = "ResourceNotFound.DataEngineConfigInstanceNotUnique"
//  RESOURCENOTFOUND_DATAENGINENOTACTIVITY = "ResourceNotFound.DataEngineNotActivity"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTMULTIVERSION = "ResourceNotFound.DataEngineNotMultiVersion"
//  RESOURCENOTFOUND_DATAENGINENOTUNIQUE = "ResourceNotFound.DataEngineNotUnique"
//  RESOURCENOTFOUND_IMAGESESSIONCONFIGNOTFOUND = "ResourceNotFound.ImageSessionConfigNotFound"
//  RESOURCENOTFOUND_IMAGESESSIONCONFIGNOTUNIQUE = "ResourceNotFound.ImageSessionConfigNotUnique"
//  RESOURCENOTFOUND_IMAGEVERSIONNOTACTIVITY = "ResourceNotFound.ImageVersionNotActivity"
//  RESOURCENOTFOUND_IMAGEVERSIONNOTUNIQUE = "ResourceNotFound.ImageVersionNotUnique"
func (c *Client) RollbackDataEngineImage(request *RollbackDataEngineImageRequest) (response *RollbackDataEngineImageResponse, err error) {
    return c.RollbackDataEngineImageWithContext(context.Background(), request)
}

// RollbackDataEngineImage
// This API is used to roll back the versions of the engine image.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOPERMISSIONTOUSETHEDATAENGINE = "FailedOperation.NoPermissionToUseTheDataEngine"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_DATAENGINEEXECTYPENOTMATCH = "InvalidParameter.DataEngineExecTypeNotMatch"
//  INVALIDPARAMETER_DATAENGINEIMAGEOPERATENOTMATCH = "InvalidParameter.DataEngineImageOperateNotMatch"
//  INVALIDPARAMETER_DATAENGINEPAYMODETYPENOTMATCH = "InvalidParameter.DataEnginePayModeTypeNotMatch"
//  INVALIDPARAMETER_DATAENGINETYPENOTMATCH = "InvalidParameter.DataEngineTypeNotMatch"
//  INVALIDPARAMETER_IMAGECLUSTERPARAMETERSFORMATNOTJSON = "InvalidParameter.ImageClusterParametersFormatNotJson"
//  INVALIDPARAMETER_IMAGEENGINETYPENOTMATCH = "InvalidParameter.ImageEngineTypeNotMatch"
//  INVALIDPARAMETER_IMAGEISPUBLICNOTMATCH = "InvalidParameter.ImageIsPublicNotMatch"
//  INVALIDPARAMETER_IMAGEPARAMETERNOTFOUND = "InvalidParameter.ImageParameterNotFound"
//  INVALIDPARAMETER_IMAGESESSIONPARAMETERSFORMATNOTJSON = "InvalidParameter.ImageSessionParametersFormatNotJson"
//  INVALIDPARAMETER_IMAGESTATENOTMATCH = "InvalidParameter.ImageStateNotMatch"
//  INVALIDPARAMETER_IMAGEUSERRECORDSTYPENOTMATCH = "InvalidParameter.ImageUserRecordsTypeNotMatch"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  RESOURCENOTFOUND_DATAENGINECONFIGINSTANCENOTFOUND = "ResourceNotFound.DataEngineConfigInstanceNotFound"
//  RESOURCENOTFOUND_DATAENGINECONFIGINSTANCENOTUNIQUE = "ResourceNotFound.DataEngineConfigInstanceNotUnique"
//  RESOURCENOTFOUND_DATAENGINENOTACTIVITY = "ResourceNotFound.DataEngineNotActivity"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTMULTIVERSION = "ResourceNotFound.DataEngineNotMultiVersion"
//  RESOURCENOTFOUND_DATAENGINENOTUNIQUE = "ResourceNotFound.DataEngineNotUnique"
//  RESOURCENOTFOUND_IMAGESESSIONCONFIGNOTFOUND = "ResourceNotFound.ImageSessionConfigNotFound"
//  RESOURCENOTFOUND_IMAGESESSIONCONFIGNOTUNIQUE = "ResourceNotFound.ImageSessionConfigNotUnique"
//  RESOURCENOTFOUND_IMAGEVERSIONNOTACTIVITY = "ResourceNotFound.ImageVersionNotActivity"
//  RESOURCENOTFOUND_IMAGEVERSIONNOTUNIQUE = "ResourceNotFound.ImageVersionNotUnique"
func (c *Client) RollbackDataEngineImageWithContext(ctx context.Context, request *RollbackDataEngineImageRequest) (response *RollbackDataEngineImageResponse, err error) {
    if request == nil {
        request = NewRollbackDataEngineImageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "RollbackDataEngineImage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RollbackDataEngineImage require credential")
    }

    request.SetContext(ctx)
    
    response = NewRollbackDataEngineImageResponse()
    err = c.Send(request, response)
    return
}

func NewSuspendResumeDataEngineRequest() (request *SuspendResumeDataEngineRequest) {
    request = &SuspendResumeDataEngineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "SuspendResumeDataEngine")
    
    
    return
}

func NewSuspendResumeDataEngineResponse() (response *SuspendResumeDataEngineResponse) {
    response = &SuspendResumeDataEngineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SuspendResumeDataEngine
// This API is used to suspend or start a data engine.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAENGINECLUSTERTYPENOTMATCH = "InvalidParameter.DataEngineClusterTypeNotMatch"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
//  UNAUTHORIZEDOPERATION_OPERATECOMPUTINGENGINE = "UnauthorizedOperation.OperateComputingEngine"
func (c *Client) SuspendResumeDataEngine(request *SuspendResumeDataEngineRequest) (response *SuspendResumeDataEngineResponse, err error) {
    return c.SuspendResumeDataEngineWithContext(context.Background(), request)
}

// SuspendResumeDataEngine
// This API is used to suspend or start a data engine.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAENGINECLUSTERTYPENOTMATCH = "InvalidParameter.DataEngineClusterTypeNotMatch"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
//  UNAUTHORIZEDOPERATION_OPERATECOMPUTINGENGINE = "UnauthorizedOperation.OperateComputingEngine"
func (c *Client) SuspendResumeDataEngineWithContext(ctx context.Context, request *SuspendResumeDataEngineRequest) (response *SuspendResumeDataEngineResponse, err error) {
    if request == nil {
        request = NewSuspendResumeDataEngineRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "SuspendResumeDataEngine")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SuspendResumeDataEngine require credential")
    }

    request.SetContext(ctx)
    
    response = NewSuspendResumeDataEngineResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchDataEngineRequest() (request *SwitchDataEngineRequest) {
    request = &SwitchDataEngineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "SwitchDataEngine")
    
    
    return
}

func NewSwitchDataEngineResponse() (response *SwitchDataEngineResponse) {
    response = &SwitchDataEngineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SwitchDataEngine
// This API is used to switch between the primary and standby clusters.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAENGINECLUSTERTYPENOTMATCH = "InvalidParameter.DataEngineClusterTypeNotMatch"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
//  UNAUTHORIZEDOPERATION_OPERATECOMPUTINGENGINE = "UnauthorizedOperation.OperateComputingEngine"
func (c *Client) SwitchDataEngine(request *SwitchDataEngineRequest) (response *SwitchDataEngineResponse, err error) {
    return c.SwitchDataEngineWithContext(context.Background(), request)
}

// SwitchDataEngine
// This API is used to switch between the primary and standby clusters.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAENGINECLUSTERTYPENOTMATCH = "InvalidParameter.DataEngineClusterTypeNotMatch"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
//  UNAUTHORIZEDOPERATION_OPERATECOMPUTINGENGINE = "UnauthorizedOperation.OperateComputingEngine"
func (c *Client) SwitchDataEngineWithContext(ctx context.Context, request *SwitchDataEngineRequest) (response *SwitchDataEngineResponse, err error) {
    if request == nil {
        request = NewSwitchDataEngineRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "SwitchDataEngine")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SwitchDataEngine require credential")
    }

    request.SetContext(ctx)
    
    response = NewSwitchDataEngineResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchDataEngineImageRequest() (request *SwitchDataEngineImageRequest) {
    request = &SwitchDataEngineImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "SwitchDataEngineImage")
    
    
    return
}

func NewSwitchDataEngineImageResponse() (response *SwitchDataEngineImageResponse) {
    response = &SwitchDataEngineImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SwitchDataEngineImage
// This API is used to switch the versions of the engine image.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOPERMISSIONTOUSETHEDATAENGINE = "FailedOperation.NoPermissionToUseTheDataEngine"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_DATAENGINEEXECTYPENOTMATCH = "InvalidParameter.DataEngineExecTypeNotMatch"
//  INVALIDPARAMETER_DATAENGINEPAYMODETYPENOTMATCH = "InvalidParameter.DataEnginePayModeTypeNotMatch"
//  INVALIDPARAMETER_IMAGECLUSTERPARAMETERSFORMATNOTJSON = "InvalidParameter.ImageClusterParametersFormatNotJson"
//  INVALIDPARAMETER_IMAGEENGINETYPENOTMATCH = "InvalidParameter.ImageEngineTypeNotMatch"
//  INVALIDPARAMETER_IMAGEISPUBLICNOTMATCH = "InvalidParameter.ImageIsPublicNotMatch"
//  INVALIDPARAMETER_IMAGEPARAMETERNOTFOUND = "InvalidParameter.ImageParameterNotFound"
//  INVALIDPARAMETER_IMAGESESSIONPARAMETERSFORMATNOTJSON = "InvalidParameter.ImageSessionParametersFormatNotJson"
//  INVALIDPARAMETER_IMAGESTATENOTMATCH = "InvalidParameter.ImageStateNotMatch"
//  INVALIDPARAMETER_IMAGEUSERRECORDSTYPENOTMATCH = "InvalidParameter.ImageUserRecordsTypeNotMatch"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  RESOURCENOTFOUND_DATAENGINECONFIGINSTANCENOTFOUND = "ResourceNotFound.DataEngineConfigInstanceNotFound"
//  RESOURCENOTFOUND_DATAENGINECONFIGINSTANCENOTUNIQUE = "ResourceNotFound.DataEngineConfigInstanceNotUnique"
//  RESOURCENOTFOUND_DATAENGINENOTACTIVITY = "ResourceNotFound.DataEngineNotActivity"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTMULTIVERSION = "ResourceNotFound.DataEngineNotMultiVersion"
//  RESOURCENOTFOUND_IMAGESESSIONCONFIGNOTFOUND = "ResourceNotFound.ImageSessionConfigNotFound"
//  RESOURCENOTFOUND_IMAGESESSIONCONFIGNOTUNIQUE = "ResourceNotFound.ImageSessionConfigNotUnique"
//  RESOURCENOTFOUND_IMAGEVERSIONNOTACTIVITY = "ResourceNotFound.ImageVersionNotActivity"
//  RESOURCENOTFOUND_IMAGEVERSIONNOTFOUND = "ResourceNotFound.ImageVersionNotFound"
func (c *Client) SwitchDataEngineImage(request *SwitchDataEngineImageRequest) (response *SwitchDataEngineImageResponse, err error) {
    return c.SwitchDataEngineImageWithContext(context.Background(), request)
}

// SwitchDataEngineImage
// This API is used to switch the versions of the engine image.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOPERMISSIONTOUSETHEDATAENGINE = "FailedOperation.NoPermissionToUseTheDataEngine"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_DATAENGINEEXECTYPENOTMATCH = "InvalidParameter.DataEngineExecTypeNotMatch"
//  INVALIDPARAMETER_DATAENGINEPAYMODETYPENOTMATCH = "InvalidParameter.DataEnginePayModeTypeNotMatch"
//  INVALIDPARAMETER_IMAGECLUSTERPARAMETERSFORMATNOTJSON = "InvalidParameter.ImageClusterParametersFormatNotJson"
//  INVALIDPARAMETER_IMAGEENGINETYPENOTMATCH = "InvalidParameter.ImageEngineTypeNotMatch"
//  INVALIDPARAMETER_IMAGEISPUBLICNOTMATCH = "InvalidParameter.ImageIsPublicNotMatch"
//  INVALIDPARAMETER_IMAGEPARAMETERNOTFOUND = "InvalidParameter.ImageParameterNotFound"
//  INVALIDPARAMETER_IMAGESESSIONPARAMETERSFORMATNOTJSON = "InvalidParameter.ImageSessionParametersFormatNotJson"
//  INVALIDPARAMETER_IMAGESTATENOTMATCH = "InvalidParameter.ImageStateNotMatch"
//  INVALIDPARAMETER_IMAGEUSERRECORDSTYPENOTMATCH = "InvalidParameter.ImageUserRecordsTypeNotMatch"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  RESOURCENOTFOUND_DATAENGINECONFIGINSTANCENOTFOUND = "ResourceNotFound.DataEngineConfigInstanceNotFound"
//  RESOURCENOTFOUND_DATAENGINECONFIGINSTANCENOTUNIQUE = "ResourceNotFound.DataEngineConfigInstanceNotUnique"
//  RESOURCENOTFOUND_DATAENGINENOTACTIVITY = "ResourceNotFound.DataEngineNotActivity"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTMULTIVERSION = "ResourceNotFound.DataEngineNotMultiVersion"
//  RESOURCENOTFOUND_IMAGESESSIONCONFIGNOTFOUND = "ResourceNotFound.ImageSessionConfigNotFound"
//  RESOURCENOTFOUND_IMAGESESSIONCONFIGNOTUNIQUE = "ResourceNotFound.ImageSessionConfigNotUnique"
//  RESOURCENOTFOUND_IMAGEVERSIONNOTACTIVITY = "ResourceNotFound.ImageVersionNotActivity"
//  RESOURCENOTFOUND_IMAGEVERSIONNOTFOUND = "ResourceNotFound.ImageVersionNotFound"
func (c *Client) SwitchDataEngineImageWithContext(ctx context.Context, request *SwitchDataEngineImageRequest) (response *SwitchDataEngineImageResponse, err error) {
    if request == nil {
        request = NewSwitchDataEngineImageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "SwitchDataEngineImage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SwitchDataEngineImage require credential")
    }

    request.SetContext(ctx)
    
    response = NewSwitchDataEngineImageResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindWorkGroupsFromUserRequest() (request *UnbindWorkGroupsFromUserRequest) {
    request = &UnbindWorkGroupsFromUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "UnbindWorkGroupsFromUser")
    
    
    return
}

func NewUnbindWorkGroupsFromUserResponse() (response *UnbindWorkGroupsFromUserResponse) {
    response = &UnbindWorkGroupsFromUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnbindWorkGroupsFromUser
// This API is used to unbind a user group from a user.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_DUPLICATEUSERNAME = "InvalidParameter.DuplicateUserName"
//  INVALIDPARAMETER_INVALIDGROUPID = "InvalidParameter.InvalidGroupId"
//  UNAUTHORIZEDOPERATION_UNBINDWORKGROUPSFROMUSER = "UnauthorizedOperation.UnbindWorkgroupsFromUser"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
func (c *Client) UnbindWorkGroupsFromUser(request *UnbindWorkGroupsFromUserRequest) (response *UnbindWorkGroupsFromUserResponse, err error) {
    return c.UnbindWorkGroupsFromUserWithContext(context.Background(), request)
}

// UnbindWorkGroupsFromUser
// This API is used to unbind a user group from a user.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_DUPLICATEUSERNAME = "InvalidParameter.DuplicateUserName"
//  INVALIDPARAMETER_INVALIDGROUPID = "InvalidParameter.InvalidGroupId"
//  UNAUTHORIZEDOPERATION_UNBINDWORKGROUPSFROMUSER = "UnauthorizedOperation.UnbindWorkgroupsFromUser"
//  UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"
func (c *Client) UnbindWorkGroupsFromUserWithContext(ctx context.Context, request *UnbindWorkGroupsFromUserRequest) (response *UnbindWorkGroupsFromUserResponse, err error) {
    if request == nil {
        request = NewUnbindWorkGroupsFromUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "UnbindWorkGroupsFromUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindWorkGroupsFromUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindWorkGroupsFromUserResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDataEngineRequest() (request *UpdateDataEngineRequest) {
    request = &UpdateDataEngineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "UpdateDataEngine")
    
    
    return
}

func NewUpdateDataEngineResponse() (response *UpdateDataEngineResponse) {
    response = &UpdateDataEngineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateDataEngine
// This API is used to upgrade data engine configuration.
//
// error code that may be returned:
//  FAILEDOPERATION_ABNORMALORDERSTATUS = "FailedOperation.AbnormalOrderStatus"
//  FAILEDOPERATION_ANOTHERPROCESSRUNNING = "FailedOperation.AnotherProcessRunning"
//  FAILEDOPERATION_ANOTHERREQUESTPROCESSING = "FailedOperation.AnotherRequestProcessing"
//  FAILEDOPERATION_BALANCENOTENOUGH = "FailedOperation.BalanceNotEnough"
//  FAILEDOPERATION_BILLINGSYSTEMERROR = "FailedOperation.BillingSystemError"
//  FAILEDOPERATION_DELIVERGOODSFAILED = "FailedOperation.DeliverGoodsFailed"
//  FAILEDOPERATION_FEEDEDUCTIONFAILED = "FailedOperation.FeeDeductionFailed"
//  FAILEDOPERATION_GETPRODUCTINFORMATIONFAILED = "FailedOperation.GetProductInformationFailed"
//  FAILEDOPERATION_MODIFYINSTANCEFAILED = "FailedOperation.ModifyInstanceFailed"
//  FAILEDOPERATION_NUMBEREXCEEDLIMIT = "FailedOperation.NumberExceedLimit"
//  FAILEDOPERATION_PARAMETERVALIDATIONFAILED = "FailedOperation.ParameterValidationFailed"
//  FAILEDOPERATION_REFUNDDEPOSITFAILED = "FailedOperation.RefundDepositFailed"
//  FAILEDOPERATION_TOOMANYRESOURCES = "FailedOperation.TooManyResources"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCEINPROCESS = "InvalidParameter.InstanceInProcess"
//  INVALIDPARAMETER_INVALIDDATAENGINEDESCRIPTION = "InvalidParameter.InvalidDataEngineDescription"
//  INVALIDPARAMETER_INVALIDDATAENGINEMODE = "InvalidParameter.InvalidDataEngineMode"
//  INVALIDPARAMETER_INVALIDDATAENGINENAME = "InvalidParameter.InvalidDataEngineName"
//  INVALIDPARAMETER_INVALIDDATAENGINESPECS = "InvalidParameter.InvalidDataEngineSpecs"
//  INVALIDPARAMETER_INVALIDDEFAULTDATAENGINE = "InvalidParameter.InvalidDefaultDataEngine"
//  INVALIDPARAMETER_INVALIDDESCRIPTION = "InvalidParameter.InvalidDescription"
//  INVALIDPARAMETER_INVALIDMINCLUSTERS = "InvalidParameter.InvalidMinClusters"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCEINUSE_UNFINISHEDSQLS = "ResourceInUse.UnfinishedSQLs"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_MODIFYCOMPUTINGENGINE = "UnauthorizedOperation.ModifyComputingEngine"
//  UNAUTHORIZEDOPERATION_NOPAYMENTAUTHORITY = "UnauthorizedOperation.NoPaymentAuthority"
func (c *Client) UpdateDataEngine(request *UpdateDataEngineRequest) (response *UpdateDataEngineResponse, err error) {
    return c.UpdateDataEngineWithContext(context.Background(), request)
}

// UpdateDataEngine
// This API is used to upgrade data engine configuration.
//
// error code that may be returned:
//  FAILEDOPERATION_ABNORMALORDERSTATUS = "FailedOperation.AbnormalOrderStatus"
//  FAILEDOPERATION_ANOTHERPROCESSRUNNING = "FailedOperation.AnotherProcessRunning"
//  FAILEDOPERATION_ANOTHERREQUESTPROCESSING = "FailedOperation.AnotherRequestProcessing"
//  FAILEDOPERATION_BALANCENOTENOUGH = "FailedOperation.BalanceNotEnough"
//  FAILEDOPERATION_BILLINGSYSTEMERROR = "FailedOperation.BillingSystemError"
//  FAILEDOPERATION_DELIVERGOODSFAILED = "FailedOperation.DeliverGoodsFailed"
//  FAILEDOPERATION_FEEDEDUCTIONFAILED = "FailedOperation.FeeDeductionFailed"
//  FAILEDOPERATION_GETPRODUCTINFORMATIONFAILED = "FailedOperation.GetProductInformationFailed"
//  FAILEDOPERATION_MODIFYINSTANCEFAILED = "FailedOperation.ModifyInstanceFailed"
//  FAILEDOPERATION_NUMBEREXCEEDLIMIT = "FailedOperation.NumberExceedLimit"
//  FAILEDOPERATION_PARAMETERVALIDATIONFAILED = "FailedOperation.ParameterValidationFailed"
//  FAILEDOPERATION_REFUNDDEPOSITFAILED = "FailedOperation.RefundDepositFailed"
//  FAILEDOPERATION_TOOMANYRESOURCES = "FailedOperation.TooManyResources"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCEINPROCESS = "InvalidParameter.InstanceInProcess"
//  INVALIDPARAMETER_INVALIDDATAENGINEDESCRIPTION = "InvalidParameter.InvalidDataEngineDescription"
//  INVALIDPARAMETER_INVALIDDATAENGINEMODE = "InvalidParameter.InvalidDataEngineMode"
//  INVALIDPARAMETER_INVALIDDATAENGINENAME = "InvalidParameter.InvalidDataEngineName"
//  INVALIDPARAMETER_INVALIDDATAENGINESPECS = "InvalidParameter.InvalidDataEngineSpecs"
//  INVALIDPARAMETER_INVALIDDEFAULTDATAENGINE = "InvalidParameter.InvalidDefaultDataEngine"
//  INVALIDPARAMETER_INVALIDDESCRIPTION = "InvalidParameter.InvalidDescription"
//  INVALIDPARAMETER_INVALIDMINCLUSTERS = "InvalidParameter.InvalidMinClusters"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCEINUSE_UNFINISHEDSQLS = "ResourceInUse.UnfinishedSQLs"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_MODIFYCOMPUTINGENGINE = "UnauthorizedOperation.ModifyComputingEngine"
//  UNAUTHORIZEDOPERATION_NOPAYMENTAUTHORITY = "UnauthorizedOperation.NoPaymentAuthority"
func (c *Client) UpdateDataEngineWithContext(ctx context.Context, request *UpdateDataEngineRequest) (response *UpdateDataEngineResponse, err error) {
    if request == nil {
        request = NewUpdateDataEngineRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "UpdateDataEngine")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateDataEngine require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateDataEngineResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDataEngineConfigRequest() (request *UpdateDataEngineConfigRequest) {
    request = &UpdateDataEngineConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "UpdateDataEngineConfig")
    
    
    return
}

func NewUpdateDataEngineConfigResponse() (response *UpdateDataEngineConfigResponse) {
    response = &UpdateDataEngineConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateDataEngineConfig
// This API is used to trigger the modification of the engine configuration by the user through a certain operation.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOPERMISSIONTOUSETHEDATAENGINE = "FailedOperation.NoPermissionToUseTheDataEngine"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_DATAENGINEEXECTYPENOTMATCH = "InvalidParameter.DataEngineExecTypeNotMatch"
//  INVALIDPARAMETER_DATAENGINEPAYMODETYPENOTMATCH = "InvalidParameter.DataEnginePayModeTypeNotMatch"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  RESOURCENOTFOUND_DATAENGINECONFIGINSTANCENOTFOUND = "ResourceNotFound.DataEngineConfigInstanceNotFound"
//  RESOURCENOTFOUND_DATAENGINECONFIGINSTANCENOTUNIQUE = "ResourceNotFound.DataEngineConfigInstanceNotUnique"
//  RESOURCENOTFOUND_DATAENGINENOTACTIVITY = "ResourceNotFound.DataEngineNotActivity"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTMULTIVERSION = "ResourceNotFound.DataEngineNotMultiVersion"
func (c *Client) UpdateDataEngineConfig(request *UpdateDataEngineConfigRequest) (response *UpdateDataEngineConfigResponse, err error) {
    return c.UpdateDataEngineConfigWithContext(context.Background(), request)
}

// UpdateDataEngineConfig
// This API is used to trigger the modification of the engine configuration by the user through a certain operation.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOPERMISSIONTOUSETHEDATAENGINE = "FailedOperation.NoPermissionToUseTheDataEngine"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_DATAENGINEEXECTYPENOTMATCH = "InvalidParameter.DataEngineExecTypeNotMatch"
//  INVALIDPARAMETER_DATAENGINEPAYMODETYPENOTMATCH = "InvalidParameter.DataEnginePayModeTypeNotMatch"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  RESOURCENOTFOUND_DATAENGINECONFIGINSTANCENOTFOUND = "ResourceNotFound.DataEngineConfigInstanceNotFound"
//  RESOURCENOTFOUND_DATAENGINECONFIGINSTANCENOTUNIQUE = "ResourceNotFound.DataEngineConfigInstanceNotUnique"
//  RESOURCENOTFOUND_DATAENGINENOTACTIVITY = "ResourceNotFound.DataEngineNotActivity"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTMULTIVERSION = "ResourceNotFound.DataEngineNotMultiVersion"
func (c *Client) UpdateDataEngineConfigWithContext(ctx context.Context, request *UpdateDataEngineConfigRequest) (response *UpdateDataEngineConfigResponse, err error) {
    if request == nil {
        request = NewUpdateDataEngineConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "UpdateDataEngineConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateDataEngineConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateDataEngineConfigResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateRowFilterRequest() (request *UpdateRowFilterRequest) {
    request = &UpdateRowFilterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "UpdateRowFilter")
    
    
    return
}

func NewUpdateRowFilterResponse() (response *UpdateRowFilterResponse) {
    response = &UpdateRowFilterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateRowFilter
// This API is used to update row filters. Please note that it updates filters only but not catalogs, databases, or tables.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOPERMISSIONTOUSETHEDATAENGINE = "FailedOperation.NoPermissionToUseTheDataEngine"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_DATAENGINEEXECTYPENOTMATCH = "InvalidParameter.DataEngineExecTypeNotMatch"
//  INVALIDPARAMETER_DATAENGINEPAYMODETYPENOTMATCH = "InvalidParameter.DataEnginePayModeTypeNotMatch"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  RESOURCENOTFOUND_DATAENGINECONFIGINSTANCENOTFOUND = "ResourceNotFound.DataEngineConfigInstanceNotFound"
//  RESOURCENOTFOUND_DATAENGINECONFIGINSTANCENOTUNIQUE = "ResourceNotFound.DataEngineConfigInstanceNotUnique"
//  RESOURCENOTFOUND_DATAENGINENOTACTIVITY = "ResourceNotFound.DataEngineNotActivity"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTMULTIVERSION = "ResourceNotFound.DataEngineNotMultiVersion"
func (c *Client) UpdateRowFilter(request *UpdateRowFilterRequest) (response *UpdateRowFilterResponse, err error) {
    return c.UpdateRowFilterWithContext(context.Background(), request)
}

// UpdateRowFilter
// This API is used to update row filters. Please note that it updates filters only but not catalogs, databases, or tables.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOPERMISSIONTOUSETHEDATAENGINE = "FailedOperation.NoPermissionToUseTheDataEngine"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_DATAENGINEEXECTYPENOTMATCH = "InvalidParameter.DataEngineExecTypeNotMatch"
//  INVALIDPARAMETER_DATAENGINEPAYMODETYPENOTMATCH = "InvalidParameter.DataEnginePayModeTypeNotMatch"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  RESOURCENOTFOUND_DATAENGINECONFIGINSTANCENOTFOUND = "ResourceNotFound.DataEngineConfigInstanceNotFound"
//  RESOURCENOTFOUND_DATAENGINECONFIGINSTANCENOTUNIQUE = "ResourceNotFound.DataEngineConfigInstanceNotUnique"
//  RESOURCENOTFOUND_DATAENGINENOTACTIVITY = "ResourceNotFound.DataEngineNotActivity"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTMULTIVERSION = "ResourceNotFound.DataEngineNotMultiVersion"
func (c *Client) UpdateRowFilterWithContext(ctx context.Context, request *UpdateRowFilterRequest) (response *UpdateRowFilterResponse, err error) {
    if request == nil {
        request = NewUpdateRowFilterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "UpdateRowFilter")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateRowFilter require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateRowFilterResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateUserDataEngineConfigRequest() (request *UpdateUserDataEngineConfigRequest) {
    request = &UpdateUserDataEngineConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "UpdateUserDataEngineConfig")
    
    
    return
}

func NewUpdateUserDataEngineConfigResponse() (response *UpdateUserDataEngineConfigResponse) {
    response = &UpdateUserDataEngineConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateUserDataEngineConfig
// This API is used to modify the custom configuration of the user's engine.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_DATAENGINECONFIGPAIRSDUPLICATE = "InvalidParameter.DataEngineConfigPairsDuplicate"
//  INVALIDPARAMETER_IMAGEENGINETYPENOTMATCH = "InvalidParameter.ImageEngineTypeNotMatch"
//  INVALIDPARAMETER_IMAGEISPUBLICNOTMATCH = "InvalidParameter.ImageIsPublicNotMatch"
//  INVALIDPARAMETER_IMAGEPARAMETERSUBMITMETHODNOTMATCH = "InvalidParameter.ImageParameterSubmitMethodNotMatch"
//  INVALIDPARAMETER_IMAGEPARAMETERTYPENOTMATCH = "InvalidParameter.ImageParameterTypeNotMatch"
//  INVALIDPARAMETER_INVALIDDATAENGINECONFIGPAIRS = "InvalidParameter.InvalidDataEngineConfigPairs"
//  INVALIDPARAMETER_INVALIDWHITELISTKEY = "InvalidParameter.InvalidWhiteListKey"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  RESOURCENOTFOUND_DATAENGINECONFIGINSTANCENOTFOUND = "ResourceNotFound.DataEngineConfigInstanceNotFound"
//  RESOURCENOTFOUND_DATAENGINECONFIGINSTANCENOTUNIQUE = "ResourceNotFound.DataEngineConfigInstanceNotUnique"
//  RESOURCENOTFOUND_DATAENGINENOTACTIVITY = "ResourceNotFound.DataEngineNotActivity"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTMULTIVERSION = "ResourceNotFound.DataEngineNotMultiVersion"
//  RESOURCENOTFOUND_DATAENGINENOTUNIQUE = "ResourceNotFound.DataEngineNotUnique"
//  RESOURCENOTFOUND_DATAENGINETYPEONLYSUPPORTBATCH = "ResourceNotFound.DataEngineTypeOnlySupportBatch"
//  RESOURCENOTFOUND_DEFAULTDATAENGINENOTFOUND = "ResourceNotFound.DefaultDataEngineNotFound"
//  RESOURCENOTFOUND_IMAGESESSIONCONFIGNOTFOUND = "ResourceNotFound.ImageSessionConfigNotFound"
//  RESOURCENOTFOUND_IMAGESESSIONCONFIGNOTUNIQUE = "ResourceNotFound.ImageSessionConfigNotUnique"
func (c *Client) UpdateUserDataEngineConfig(request *UpdateUserDataEngineConfigRequest) (response *UpdateUserDataEngineConfigResponse, err error) {
    return c.UpdateUserDataEngineConfigWithContext(context.Background(), request)
}

// UpdateUserDataEngineConfig
// This API is used to modify the custom configuration of the user's engine.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_DATAENGINECONFIGPAIRSDUPLICATE = "InvalidParameter.DataEngineConfigPairsDuplicate"
//  INVALIDPARAMETER_IMAGEENGINETYPENOTMATCH = "InvalidParameter.ImageEngineTypeNotMatch"
//  INVALIDPARAMETER_IMAGEISPUBLICNOTMATCH = "InvalidParameter.ImageIsPublicNotMatch"
//  INVALIDPARAMETER_IMAGEPARAMETERSUBMITMETHODNOTMATCH = "InvalidParameter.ImageParameterSubmitMethodNotMatch"
//  INVALIDPARAMETER_IMAGEPARAMETERTYPENOTMATCH = "InvalidParameter.ImageParameterTypeNotMatch"
//  INVALIDPARAMETER_INVALIDDATAENGINECONFIGPAIRS = "InvalidParameter.InvalidDataEngineConfigPairs"
//  INVALIDPARAMETER_INVALIDWHITELISTKEY = "InvalidParameter.InvalidWhiteListKey"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  RESOURCENOTFOUND_DATAENGINECONFIGINSTANCENOTFOUND = "ResourceNotFound.DataEngineConfigInstanceNotFound"
//  RESOURCENOTFOUND_DATAENGINECONFIGINSTANCENOTUNIQUE = "ResourceNotFound.DataEngineConfigInstanceNotUnique"
//  RESOURCENOTFOUND_DATAENGINENOTACTIVITY = "ResourceNotFound.DataEngineNotActivity"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTMULTIVERSION = "ResourceNotFound.DataEngineNotMultiVersion"
//  RESOURCENOTFOUND_DATAENGINENOTUNIQUE = "ResourceNotFound.DataEngineNotUnique"
//  RESOURCENOTFOUND_DATAENGINETYPEONLYSUPPORTBATCH = "ResourceNotFound.DataEngineTypeOnlySupportBatch"
//  RESOURCENOTFOUND_DEFAULTDATAENGINENOTFOUND = "ResourceNotFound.DefaultDataEngineNotFound"
//  RESOURCENOTFOUND_IMAGESESSIONCONFIGNOTFOUND = "ResourceNotFound.ImageSessionConfigNotFound"
//  RESOURCENOTFOUND_IMAGESESSIONCONFIGNOTUNIQUE = "ResourceNotFound.ImageSessionConfigNotUnique"
func (c *Client) UpdateUserDataEngineConfigWithContext(ctx context.Context, request *UpdateUserDataEngineConfigRequest) (response *UpdateUserDataEngineConfigResponse, err error) {
    if request == nil {
        request = NewUpdateUserDataEngineConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "UpdateUserDataEngineConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateUserDataEngineConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateUserDataEngineConfigResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeDataEngineImageRequest() (request *UpgradeDataEngineImageRequest) {
    request = &UpgradeDataEngineImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dlc", APIVersion, "UpgradeDataEngineImage")
    
    
    return
}

func NewUpgradeDataEngineImageResponse() (response *UpgradeDataEngineImageResponse) {
    response = &UpgradeDataEngineImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpgradeDataEngineImage
// This API is used to upgrade the engine image.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOPERMISSIONTOUSETHEDATAENGINE = "FailedOperation.NoPermissionToUseTheDataEngine"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_DATAENGINEEXECTYPENOTMATCH = "InvalidParameter.DataEngineExecTypeNotMatch"
//  INVALIDPARAMETER_DATAENGINEPAYMODETYPENOTMATCH = "InvalidParameter.DataEnginePayModeTypeNotMatch"
//  INVALIDPARAMETER_IMAGECLUSTERPARAMETERSFORMATNOTJSON = "InvalidParameter.ImageClusterParametersFormatNotJson"
//  INVALIDPARAMETER_IMAGEENGINETYPENOTMATCH = "InvalidParameter.ImageEngineTypeNotMatch"
//  INVALIDPARAMETER_IMAGEISPUBLICNOTMATCH = "InvalidParameter.ImageIsPublicNotMatch"
//  INVALIDPARAMETER_IMAGEPARAMETERNOTFOUND = "InvalidParameter.ImageParameterNotFound"
//  INVALIDPARAMETER_IMAGESESSIONPARAMETERSFORMATNOTJSON = "InvalidParameter.ImageSessionParametersFormatNotJson"
//  INVALIDPARAMETER_IMAGESTATENOTMATCH = "InvalidParameter.ImageStateNotMatch"
//  INVALIDPARAMETER_IMAGEUSERRECORDSTYPENOTMATCH = "InvalidParameter.ImageUserRecordsTypeNotMatch"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  RESOURCENOTFOUND_DATAENGINECONFIGINSTANCENOTFOUND = "ResourceNotFound.DataEngineConfigInstanceNotFound"
//  RESOURCENOTFOUND_DATAENGINECONFIGINSTANCENOTUNIQUE = "ResourceNotFound.DataEngineConfigInstanceNotUnique"
//  RESOURCENOTFOUND_DATAENGINENOTACTIVITY = "ResourceNotFound.DataEngineNotActivity"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTMULTIVERSION = "ResourceNotFound.DataEngineNotMultiVersion"
//  RESOURCENOTFOUND_DATAENGINENOTUNIQUE = "ResourceNotFound.DataEngineNotUnique"
//  RESOURCENOTFOUND_IMAGESESSIONCONFIGNOTFOUND = "ResourceNotFound.ImageSessionConfigNotFound"
//  RESOURCENOTFOUND_IMAGESESSIONCONFIGNOTUNIQUE = "ResourceNotFound.ImageSessionConfigNotUnique"
//  RESOURCENOTFOUND_IMAGEVERSIONNOTFOUND = "ResourceNotFound.ImageVersionNotFound"
func (c *Client) UpgradeDataEngineImage(request *UpgradeDataEngineImageRequest) (response *UpgradeDataEngineImageResponse, err error) {
    return c.UpgradeDataEngineImageWithContext(context.Background(), request)
}

// UpgradeDataEngineImage
// This API is used to upgrade the engine image.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOPERMISSIONTOUSETHEDATAENGINE = "FailedOperation.NoPermissionToUseTheDataEngine"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"
//  INVALIDPARAMETER_DATAENGINEEXECTYPENOTMATCH = "InvalidParameter.DataEngineExecTypeNotMatch"
//  INVALIDPARAMETER_DATAENGINEPAYMODETYPENOTMATCH = "InvalidParameter.DataEnginePayModeTypeNotMatch"
//  INVALIDPARAMETER_IMAGECLUSTERPARAMETERSFORMATNOTJSON = "InvalidParameter.ImageClusterParametersFormatNotJson"
//  INVALIDPARAMETER_IMAGEENGINETYPENOTMATCH = "InvalidParameter.ImageEngineTypeNotMatch"
//  INVALIDPARAMETER_IMAGEISPUBLICNOTMATCH = "InvalidParameter.ImageIsPublicNotMatch"
//  INVALIDPARAMETER_IMAGEPARAMETERNOTFOUND = "InvalidParameter.ImageParameterNotFound"
//  INVALIDPARAMETER_IMAGESESSIONPARAMETERSFORMATNOTJSON = "InvalidParameter.ImageSessionParametersFormatNotJson"
//  INVALIDPARAMETER_IMAGESTATENOTMATCH = "InvalidParameter.ImageStateNotMatch"
//  INVALIDPARAMETER_IMAGEUSERRECORDSTYPENOTMATCH = "InvalidParameter.ImageUserRecordsTypeNotMatch"
//  INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"
//  RESOURCENOTFOUND_DATAENGINECONFIGINSTANCENOTFOUND = "ResourceNotFound.DataEngineConfigInstanceNotFound"
//  RESOURCENOTFOUND_DATAENGINECONFIGINSTANCENOTUNIQUE = "ResourceNotFound.DataEngineConfigInstanceNotUnique"
//  RESOURCENOTFOUND_DATAENGINENOTACTIVITY = "ResourceNotFound.DataEngineNotActivity"
//  RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"
//  RESOURCENOTFOUND_DATAENGINENOTMULTIVERSION = "ResourceNotFound.DataEngineNotMultiVersion"
//  RESOURCENOTFOUND_DATAENGINENOTUNIQUE = "ResourceNotFound.DataEngineNotUnique"
//  RESOURCENOTFOUND_IMAGESESSIONCONFIGNOTFOUND = "ResourceNotFound.ImageSessionConfigNotFound"
//  RESOURCENOTFOUND_IMAGESESSIONCONFIGNOTUNIQUE = "ResourceNotFound.ImageSessionConfigNotUnique"
//  RESOURCENOTFOUND_IMAGEVERSIONNOTFOUND = "ResourceNotFound.ImageVersionNotFound"
func (c *Client) UpgradeDataEngineImageWithContext(ctx context.Context, request *UpgradeDataEngineImageRequest) (response *UpgradeDataEngineImageResponse, err error) {
    if request == nil {
        request = NewUpgradeDataEngineImageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dlc", APIVersion, "UpgradeDataEngineImage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeDataEngineImage require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeDataEngineImageResponse()
    err = c.Send(request, response)
    return
}
