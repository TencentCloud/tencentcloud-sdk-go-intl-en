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

package v20200210

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2020-02-10"

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


func NewAbortAgentCruiseDialingCampaignRequest() (request *AbortAgentCruiseDialingCampaignRequest) {
    request = &AbortAgentCruiseDialingCampaignRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "AbortAgentCruiseDialingCampaign")
    
    
    return
}

func NewAbortAgentCruiseDialingCampaignResponse() (response *AbortAgentCruiseDialingCampaignResponse) {
    response = &AbortAgentCruiseDialingCampaignResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AbortAgentCruiseDialingCampaign
// If you want to stop running agent's individual auto task, then call AbortAgentCruiseDialingCampaign to terminate.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) AbortAgentCruiseDialingCampaign(request *AbortAgentCruiseDialingCampaignRequest) (response *AbortAgentCruiseDialingCampaignResponse, err error) {
    return c.AbortAgentCruiseDialingCampaignWithContext(context.Background(), request)
}

// AbortAgentCruiseDialingCampaign
// If you want to stop running agent's individual auto task, then call AbortAgentCruiseDialingCampaign to terminate.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) AbortAgentCruiseDialingCampaignWithContext(ctx context.Context, request *AbortAgentCruiseDialingCampaignRequest) (response *AbortAgentCruiseDialingCampaignResponse, err error) {
    if request == nil {
        request = NewAbortAgentCruiseDialingCampaignRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AbortAgentCruiseDialingCampaign require credential")
    }

    request.SetContext(ctx)
    
    response = NewAbortAgentCruiseDialingCampaignResponse()
    err = c.Send(request, response)
    return
}

func NewAbortPredictiveDialingCampaignRequest() (request *AbortPredictiveDialingCampaignRequest) {
    request = &AbortPredictiveDialingCampaignRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "AbortPredictiveDialingCampaign")
    
    
    return
}

func NewAbortPredictiveDialingCampaignResponse() (response *AbortPredictiveDialingCampaignResponse) {
    response = &AbortPredictiveDialingCampaignResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AbortPredictiveDialingCampaign
// This API is used to pause predictive dialing campaign
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) AbortPredictiveDialingCampaign(request *AbortPredictiveDialingCampaignRequest) (response *AbortPredictiveDialingCampaignResponse, err error) {
    return c.AbortPredictiveDialingCampaignWithContext(context.Background(), request)
}

// AbortPredictiveDialingCampaign
// This API is used to pause predictive dialing campaign
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) AbortPredictiveDialingCampaignWithContext(ctx context.Context, request *AbortPredictiveDialingCampaignRequest) (response *AbortPredictiveDialingCampaignResponse, err error) {
    if request == nil {
        request = NewAbortPredictiveDialingCampaignRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AbortPredictiveDialingCampaign require credential")
    }

    request.SetContext(ctx)
    
    response = NewAbortPredictiveDialingCampaignResponse()
    err = c.Send(request, response)
    return
}

func NewBindNumberCallOutSkillGroupRequest() (request *BindNumberCallOutSkillGroupRequest) {
    request = &BindNumberCallOutSkillGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "BindNumberCallOutSkillGroup")
    
    
    return
}

func NewBindNumberCallOutSkillGroupResponse() (response *BindNumberCallOutSkillGroupResponse) {
    response = &BindNumberCallOutSkillGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindNumberCallOutSkillGroup
// This API is used to assign outbound skill group(s) to your number
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
func (c *Client) BindNumberCallOutSkillGroup(request *BindNumberCallOutSkillGroupRequest) (response *BindNumberCallOutSkillGroupResponse, err error) {
    return c.BindNumberCallOutSkillGroupWithContext(context.Background(), request)
}

// BindNumberCallOutSkillGroup
// This API is used to assign outbound skill group(s) to your number
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
func (c *Client) BindNumberCallOutSkillGroupWithContext(ctx context.Context, request *BindNumberCallOutSkillGroupRequest) (response *BindNumberCallOutSkillGroupResponse, err error) {
    if request == nil {
        request = NewBindNumberCallOutSkillGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindNumberCallOutSkillGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindNumberCallOutSkillGroupResponse()
    err = c.Send(request, response)
    return
}

func NewBindStaffSkillGroupListRequest() (request *BindStaffSkillGroupListRequest) {
    request = &BindStaffSkillGroupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "BindStaffSkillGroupList")
    
    
    return
}

func NewBindStaffSkillGroupListResponse() (response *BindStaffSkillGroupListResponse) {
    response = &BindStaffSkillGroupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindStaffSkillGroupList
// This API is used to assign an agent to skill group
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) BindStaffSkillGroupList(request *BindStaffSkillGroupListRequest) (response *BindStaffSkillGroupListResponse, err error) {
    return c.BindStaffSkillGroupListWithContext(context.Background(), request)
}

// BindStaffSkillGroupList
// This API is used to assign an agent to skill group
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) BindStaffSkillGroupListWithContext(ctx context.Context, request *BindStaffSkillGroupListRequest) (response *BindStaffSkillGroupListResponse, err error) {
    if request == nil {
        request = NewBindStaffSkillGroupListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindStaffSkillGroupList require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindStaffSkillGroupListResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAdminURLRequest() (request *CreateAdminURLRequest) {
    request = &CreateAdminURLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "CreateAdminURL")
    
    
    return
}

func NewCreateAdminURLResponse() (response *CreateAdminURLResponse) {
    response = &CreateAdminURLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAdminURL
// This API is used to create a management access link.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) CreateAdminURL(request *CreateAdminURLRequest) (response *CreateAdminURLResponse, err error) {
    return c.CreateAdminURLWithContext(context.Background(), request)
}

// CreateAdminURL
// This API is used to create a management access link.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) CreateAdminURLWithContext(ctx context.Context, request *CreateAdminURLRequest) (response *CreateAdminURLResponse, err error) {
    if request == nil {
        request = NewCreateAdminURLRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAdminURL require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAdminURLResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAgentCruiseDialingCampaignRequest() (request *CreateAgentCruiseDialingCampaignRequest) {
    request = &CreateAgentCruiseDialingCampaignRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "CreateAgentCruiseDialingCampaign")
    
    
    return
}

func NewCreateAgentCruiseDialingCampaignResponse() (response *CreateAgentCruiseDialingCampaignResponse) {
    response = &CreateAgentCruiseDialingCampaignResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAgentCruiseDialingCampaign
// This document shows how to call API to create an individual auto dialing campaign for agent
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_OUTOFCOUNTLIMIT = "LimitExceeded.OutOfCountLimit"
func (c *Client) CreateAgentCruiseDialingCampaign(request *CreateAgentCruiseDialingCampaignRequest) (response *CreateAgentCruiseDialingCampaignResponse, err error) {
    return c.CreateAgentCruiseDialingCampaignWithContext(context.Background(), request)
}

// CreateAgentCruiseDialingCampaign
// This document shows how to call API to create an individual auto dialing campaign for agent
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_OUTOFCOUNTLIMIT = "LimitExceeded.OutOfCountLimit"
func (c *Client) CreateAgentCruiseDialingCampaignWithContext(ctx context.Context, request *CreateAgentCruiseDialingCampaignRequest) (response *CreateAgentCruiseDialingCampaignResponse, err error) {
    if request == nil {
        request = NewCreateAgentCruiseDialingCampaignRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAgentCruiseDialingCampaign require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAgentCruiseDialingCampaignResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAutoCalloutTaskRequest() (request *CreateAutoCalloutTaskRequest) {
    request = &CreateAutoCalloutTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "CreateAutoCalloutTask")
    
    
    return
}

func NewCreateAutoCalloutTaskResponse() (response *CreateAutoCalloutTaskResponse) {
    response = &CreateAutoCalloutTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAutoCalloutTask
// This API is used to create the automatic outbound call task.
//
// error code that may be returned:
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateAutoCalloutTask(request *CreateAutoCalloutTaskRequest) (response *CreateAutoCalloutTaskResponse, err error) {
    return c.CreateAutoCalloutTaskWithContext(context.Background(), request)
}

// CreateAutoCalloutTask
// This API is used to create the automatic outbound call task.
//
// error code that may be returned:
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateAutoCalloutTaskWithContext(ctx context.Context, request *CreateAutoCalloutTaskRequest) (response *CreateAutoCalloutTaskResponse, err error) {
    if request == nil {
        request = NewCreateAutoCalloutTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAutoCalloutTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAutoCalloutTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCCCSkillGroupRequest() (request *CreateCCCSkillGroupRequest) {
    request = &CreateCCCSkillGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "CreateCCCSkillGroup")
    
    
    return
}

func NewCreateCCCSkillGroupResponse() (response *CreateCCCSkillGroupResponse) {
    response = &CreateCCCSkillGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCCCSkillGroup
// This API is used to create a new skill group.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_SKILLGROUPEXIST = "InvalidParameterValue.SkillGroupExist"
func (c *Client) CreateCCCSkillGroup(request *CreateCCCSkillGroupRequest) (response *CreateCCCSkillGroupResponse, err error) {
    return c.CreateCCCSkillGroupWithContext(context.Background(), request)
}

// CreateCCCSkillGroup
// This API is used to create a new skill group.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_SKILLGROUPEXIST = "InvalidParameterValue.SkillGroupExist"
func (c *Client) CreateCCCSkillGroupWithContext(ctx context.Context, request *CreateCCCSkillGroupRequest) (response *CreateCCCSkillGroupResponse, err error) {
    if request == nil {
        request = NewCreateCCCSkillGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCCCSkillGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCCCSkillGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCallOutSessionRequest() (request *CreateCallOutSessionRequest) {
    request = &CreateCallOutSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "CreateCallOutSession")
    
    
    return
}

func NewCreateCallOutSessionResponse() (response *CreateCallOutSessionResponse) {
    response = &CreateCallOutSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCallOutSession
// This API is used to create outbound sessions. Currently, only dual call is supported. That is, firstly, please use the platform number to call the agent's cell phone. After the agent answers, then please make outbound calls to the user. Due to ISP frequency restrictions, the agent's phone number must first be added to the allowlist to avoid frequency control which may lead to the failure of the outbound call.
//
// error code that may be returned:
//  FAILEDOPERATION_CALLOUTFAILED = "FailedOperation.CallOutFailed"
//  FAILEDOPERATION_CALLEEISLIMITED = "FailedOperation.CalleeIsLimited"
//  FAILEDOPERATION_CALLEROVERFREQUENCY = "FailedOperation.CallerOverFrequency"
//  FAILEDOPERATION_NOCALLOUTNUMBER = "FailedOperation.NoCallOutNumber"
//  FAILEDOPERATION_SEATSTATUSBUSY = "FailedOperation.SeatStatusBusy"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXIST = "InvalidParameterValue.AccountNotExist"
//  OPERATIONDENIED_NOTINWHITELIST = "OperationDenied.NotInWhiteList"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCallOutSession(request *CreateCallOutSessionRequest) (response *CreateCallOutSessionResponse, err error) {
    return c.CreateCallOutSessionWithContext(context.Background(), request)
}

// CreateCallOutSession
// This API is used to create outbound sessions. Currently, only dual call is supported. That is, firstly, please use the platform number to call the agent's cell phone. After the agent answers, then please make outbound calls to the user. Due to ISP frequency restrictions, the agent's phone number must first be added to the allowlist to avoid frequency control which may lead to the failure of the outbound call.
//
// error code that may be returned:
//  FAILEDOPERATION_CALLOUTFAILED = "FailedOperation.CallOutFailed"
//  FAILEDOPERATION_CALLEEISLIMITED = "FailedOperation.CalleeIsLimited"
//  FAILEDOPERATION_CALLEROVERFREQUENCY = "FailedOperation.CallerOverFrequency"
//  FAILEDOPERATION_NOCALLOUTNUMBER = "FailedOperation.NoCallOutNumber"
//  FAILEDOPERATION_SEATSTATUSBUSY = "FailedOperation.SeatStatusBusy"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXIST = "InvalidParameterValue.AccountNotExist"
//  OPERATIONDENIED_NOTINWHITELIST = "OperationDenied.NotInWhiteList"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCallOutSessionWithContext(ctx context.Context, request *CreateCallOutSessionRequest) (response *CreateCallOutSessionResponse, err error) {
    if request == nil {
        request = NewCreateCallOutSessionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCallOutSession require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCallOutSessionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateExtensionRequest() (request *CreateExtensionRequest) {
    request = &CreateExtensionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "CreateExtension")
    
    
    return
}

func NewCreateExtensionResponse() (response *CreateExtensionResponse) {
    response = &CreateExtensionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateExtension
// This API is used to create the telephone account.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) CreateExtension(request *CreateExtensionRequest) (response *CreateExtensionResponse, err error) {
    return c.CreateExtensionWithContext(context.Background(), request)
}

// CreateExtension
// This API is used to create the telephone account.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) CreateExtensionWithContext(ctx context.Context, request *CreateExtensionRequest) (response *CreateExtensionResponse, err error) {
    if request == nil {
        request = NewCreateExtensionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateExtension require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateExtensionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIVRSessionRequest() (request *CreateIVRSessionRequest) {
    request = &CreateIVRSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "CreateIVRSession")
    
    
    return
}

func NewCreateIVRSessionResponse() (response *CreateIVRSessionResponse) {
    response = &CreateIVRSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateIVRSession
// Create a session associated with IVR. This feature is supported only in the Advanced Version. Currently, it supports inbound and automatic outbound IVR types. Upon receiving the request, TCCC will first attempt to call the callee, then enter the IVR flow.
//
// error code that may be returned:
//  FAILEDOPERATION_CALLOUTFAILED = "FailedOperation.CallOutFailed"
//  FAILEDOPERATION_CALLEEISBLACKUSER = "FailedOperation.CalleeIsBlackUser"
//  FAILEDOPERATION_CALLEROVERFREQUENCY = "FailedOperation.CallerOverFrequency"
//  FAILEDOPERATION_CALLOUTRULEBLINDAREA = "FailedOperation.CalloutRuleBlindArea"
//  FAILEDOPERATION_CALLOUTRULEMAXCALLCOUNTCALLEEINTERVALTIME = "FailedOperation.CalloutRuleMaxCallCountCalleeIntervalTime"
//  FAILEDOPERATION_CALLOUTRULEMAXCALLCOUNTCALLEEPERDAYAPPID = "FailedOperation.CalloutRuleMaxCallCountCalleePerDayAppID"
//  FAILEDOPERATION_CALLOUTRULENOTWORKTIME = "FailedOperation.CalloutRuleNotWorkTime"
//  FAILEDOPERATION_NOCALLOUTNUMBER = "FailedOperation.NoCallOutNumber"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  LIMITEXCEEDED_BASEPACKAGEEXPIRED = "LimitExceeded.BasePackageExpired"
func (c *Client) CreateIVRSession(request *CreateIVRSessionRequest) (response *CreateIVRSessionResponse, err error) {
    return c.CreateIVRSessionWithContext(context.Background(), request)
}

// CreateIVRSession
// Create a session associated with IVR. This feature is supported only in the Advanced Version. Currently, it supports inbound and automatic outbound IVR types. Upon receiving the request, TCCC will first attempt to call the callee, then enter the IVR flow.
//
// error code that may be returned:
//  FAILEDOPERATION_CALLOUTFAILED = "FailedOperation.CallOutFailed"
//  FAILEDOPERATION_CALLEEISBLACKUSER = "FailedOperation.CalleeIsBlackUser"
//  FAILEDOPERATION_CALLEROVERFREQUENCY = "FailedOperation.CallerOverFrequency"
//  FAILEDOPERATION_CALLOUTRULEBLINDAREA = "FailedOperation.CalloutRuleBlindArea"
//  FAILEDOPERATION_CALLOUTRULEMAXCALLCOUNTCALLEEINTERVALTIME = "FailedOperation.CalloutRuleMaxCallCountCalleeIntervalTime"
//  FAILEDOPERATION_CALLOUTRULEMAXCALLCOUNTCALLEEPERDAYAPPID = "FailedOperation.CalloutRuleMaxCallCountCalleePerDayAppID"
//  FAILEDOPERATION_CALLOUTRULENOTWORKTIME = "FailedOperation.CalloutRuleNotWorkTime"
//  FAILEDOPERATION_NOCALLOUTNUMBER = "FailedOperation.NoCallOutNumber"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  LIMITEXCEEDED_BASEPACKAGEEXPIRED = "LimitExceeded.BasePackageExpired"
func (c *Client) CreateIVRSessionWithContext(ctx context.Context, request *CreateIVRSessionRequest) (response *CreateIVRSessionResponse, err error) {
    if request == nil {
        request = NewCreateIVRSessionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIVRSession require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateIVRSessionResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePredictiveDialingCampaignRequest() (request *CreatePredictiveDialingCampaignRequest) {
    request = &CreatePredictiveDialingCampaignRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "CreatePredictiveDialingCampaign")
    
    
    return
}

func NewCreatePredictiveDialingCampaignResponse() (response *CreatePredictiveDialingCampaignResponse) {
    response = &CreatePredictiveDialingCampaignResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePredictiveDialingCampaign
// This API is used to create the predictive outbound call task.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_OUTOFCOUNTLIMIT = "LimitExceeded.OutOfCountLimit"
func (c *Client) CreatePredictiveDialingCampaign(request *CreatePredictiveDialingCampaignRequest) (response *CreatePredictiveDialingCampaignResponse, err error) {
    return c.CreatePredictiveDialingCampaignWithContext(context.Background(), request)
}

// CreatePredictiveDialingCampaign
// This API is used to create the predictive outbound call task.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_OUTOFCOUNTLIMIT = "LimitExceeded.OutOfCountLimit"
func (c *Client) CreatePredictiveDialingCampaignWithContext(ctx context.Context, request *CreatePredictiveDialingCampaignRequest) (response *CreatePredictiveDialingCampaignResponse, err error) {
    if request == nil {
        request = NewCreatePredictiveDialingCampaignRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePredictiveDialingCampaign require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePredictiveDialingCampaignResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSDKLoginTokenRequest() (request *CreateSDKLoginTokenRequest) {
    request = &CreateSDKLoginTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "CreateSDKLoginToken")
    
    
    return
}

func NewCreateSDKLoginTokenResponse() (response *CreateSDKLoginTokenResponse) {
    response = &CreateSDKLoginTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSDKLoginToken
// This API is used to create the SDK log-in token.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXIST = "InvalidParameterValue.AccountNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  OPERATIONDENIED_UINDISABLED = "OperationDenied.UinDisabled"
func (c *Client) CreateSDKLoginToken(request *CreateSDKLoginTokenRequest) (response *CreateSDKLoginTokenResponse, err error) {
    return c.CreateSDKLoginTokenWithContext(context.Background(), request)
}

// CreateSDKLoginToken
// This API is used to create the SDK log-in token.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXIST = "InvalidParameterValue.AccountNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  OPERATIONDENIED_UINDISABLED = "OperationDenied.UinDisabled"
func (c *Client) CreateSDKLoginTokenWithContext(ctx context.Context, request *CreateSDKLoginTokenRequest) (response *CreateSDKLoginTokenResponse, err error) {
    if request == nil {
        request = NewCreateSDKLoginTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSDKLoginToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSDKLoginTokenResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStaffRequest() (request *CreateStaffRequest) {
    request = &CreateStaffRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "CreateStaff")
    
    
    return
}

func NewCreateStaffResponse() (response *CreateStaffResponse) {
    response = &CreateStaffResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateStaff
// This API is used to create the customer service account.
//
// error code that may be returned:
//  FAILEDOPERATION_DUPLICATEDACCOUNT = "FailedOperation.DuplicatedAccount"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateStaff(request *CreateStaffRequest) (response *CreateStaffResponse, err error) {
    return c.CreateStaffWithContext(context.Background(), request)
}

// CreateStaff
// This API is used to create the customer service account.
//
// error code that may be returned:
//  FAILEDOPERATION_DUPLICATEDACCOUNT = "FailedOperation.DuplicatedAccount"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateStaffWithContext(ctx context.Context, request *CreateStaffRequest) (response *CreateStaffResponse, err error) {
    if request == nil {
        request = NewCreateStaffRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateStaff require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateStaffResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteExtensionRequest() (request *DeleteExtensionRequest) {
    request = &DeleteExtensionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DeleteExtension")
    
    
    return
}

func NewDeleteExtensionResponse() (response *DeleteExtensionResponse) {
    response = &DeleteExtensionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteExtension
// This API is used to delete telephone accounts.
//
// error code that may be returned:
//  FAILEDOPERATION_DUPLICATEDACCOUNT = "FailedOperation.DuplicatedAccount"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DeleteExtension(request *DeleteExtensionRequest) (response *DeleteExtensionResponse, err error) {
    return c.DeleteExtensionWithContext(context.Background(), request)
}

// DeleteExtension
// This API is used to delete telephone accounts.
//
// error code that may be returned:
//  FAILEDOPERATION_DUPLICATEDACCOUNT = "FailedOperation.DuplicatedAccount"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DeleteExtensionWithContext(ctx context.Context, request *DeleteExtensionRequest) (response *DeleteExtensionResponse, err error) {
    if request == nil {
        request = NewDeleteExtensionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteExtension require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteExtensionResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePredictiveDialingCampaignRequest() (request *DeletePredictiveDialingCampaignRequest) {
    request = &DeletePredictiveDialingCampaignRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DeletePredictiveDialingCampaign")
    
    
    return
}

func NewDeletePredictiveDialingCampaignResponse() (response *DeletePredictiveDialingCampaignResponse) {
    response = &DeletePredictiveDialingCampaignResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePredictiveDialingCampaign
// This API is used to delete the predictive outbound call task.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeletePredictiveDialingCampaign(request *DeletePredictiveDialingCampaignRequest) (response *DeletePredictiveDialingCampaignResponse, err error) {
    return c.DeletePredictiveDialingCampaignWithContext(context.Background(), request)
}

// DeletePredictiveDialingCampaign
// This API is used to delete the predictive outbound call task.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeletePredictiveDialingCampaignWithContext(ctx context.Context, request *DeletePredictiveDialingCampaignRequest) (response *DeletePredictiveDialingCampaignResponse, err error) {
    if request == nil {
        request = NewDeletePredictiveDialingCampaignRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePredictiveDialingCampaign require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePredictiveDialingCampaignResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteStaffRequest() (request *DeleteStaffRequest) {
    request = &DeleteStaffRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DeleteStaff")
    
    
    return
}

func NewDeleteStaffResponse() (response *DeleteStaffResponse) {
    response = &DeleteStaffResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteStaff
// This API is used to delete the agent information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DeleteStaff(request *DeleteStaffRequest) (response *DeleteStaffResponse, err error) {
    return c.DeleteStaffWithContext(context.Background(), request)
}

// DeleteStaff
// This API is used to delete the agent information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DeleteStaffWithContext(ctx context.Context, request *DeleteStaffRequest) (response *DeleteStaffResponse, err error) {
    if request == nil {
        request = NewDeleteStaffRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteStaff require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteStaffResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentCruiseDialingCampaignRequest() (request *DescribeAgentCruiseDialingCampaignRequest) {
    request = &DescribeAgentCruiseDialingCampaignRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeAgentCruiseDialingCampaign")
    
    
    return
}

func NewDescribeAgentCruiseDialingCampaignResponse() (response *DescribeAgentCruiseDialingCampaignResponse) {
    response = &DescribeAgentCruiseDialingCampaignResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAgentCruiseDialingCampaign
// Query Agent Cruise-style Outbound Call Task
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeAgentCruiseDialingCampaign(request *DescribeAgentCruiseDialingCampaignRequest) (response *DescribeAgentCruiseDialingCampaignResponse, err error) {
    return c.DescribeAgentCruiseDialingCampaignWithContext(context.Background(), request)
}

// DescribeAgentCruiseDialingCampaign
// Query Agent Cruise-style Outbound Call Task
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeAgentCruiseDialingCampaignWithContext(ctx context.Context, request *DescribeAgentCruiseDialingCampaignRequest) (response *DescribeAgentCruiseDialingCampaignResponse, err error) {
    if request == nil {
        request = NewDescribeAgentCruiseDialingCampaignRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentCruiseDialingCampaign require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentCruiseDialingCampaignResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutoCalloutTaskRequest() (request *DescribeAutoCalloutTaskRequest) {
    request = &DescribeAutoCalloutTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeAutoCalloutTask")
    
    
    return
}

func NewDescribeAutoCalloutTaskResponse() (response *DescribeAutoCalloutTaskResponse) {
    response = &DescribeAutoCalloutTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAutoCalloutTask
// This API is used to query automatic outbound call task details.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeAutoCalloutTask(request *DescribeAutoCalloutTaskRequest) (response *DescribeAutoCalloutTaskResponse, err error) {
    return c.DescribeAutoCalloutTaskWithContext(context.Background(), request)
}

// DescribeAutoCalloutTask
// This API is used to query automatic outbound call task details.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeAutoCalloutTaskWithContext(ctx context.Context, request *DescribeAutoCalloutTaskRequest) (response *DescribeAutoCalloutTaskResponse, err error) {
    if request == nil {
        request = NewDescribeAutoCalloutTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAutoCalloutTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAutoCalloutTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutoCalloutTasksRequest() (request *DescribeAutoCalloutTasksRequest) {
    request = &DescribeAutoCalloutTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeAutoCalloutTasks")
    
    
    return
}

func NewDescribeAutoCalloutTasksResponse() (response *DescribeAutoCalloutTasksResponse) {
    response = &DescribeAutoCalloutTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAutoCalloutTasks
// Batch Query Automatic Outbound Call Tasks
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeAutoCalloutTasks(request *DescribeAutoCalloutTasksRequest) (response *DescribeAutoCalloutTasksResponse, err error) {
    return c.DescribeAutoCalloutTasksWithContext(context.Background(), request)
}

// DescribeAutoCalloutTasks
// Batch Query Automatic Outbound Call Tasks
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeAutoCalloutTasksWithContext(ctx context.Context, request *DescribeAutoCalloutTasksRequest) (response *DescribeAutoCalloutTasksResponse, err error) {
    if request == nil {
        request = NewDescribeAutoCalloutTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAutoCalloutTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAutoCalloutTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCCCBuyInfoListRequest() (request *DescribeCCCBuyInfoListRequest) {
    request = &DescribeCCCBuyInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeCCCBuyInfoList")
    
    
    return
}

func NewDescribeCCCBuyInfoListResponse() (response *DescribeCCCBuyInfoListResponse) {
    response = &DescribeCCCBuyInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCCCBuyInfoList
// This API is used to access the user purchasing information list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXIST = "InvalidParameterValue.AccountNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCCCBuyInfoList(request *DescribeCCCBuyInfoListRequest) (response *DescribeCCCBuyInfoListResponse, err error) {
    return c.DescribeCCCBuyInfoListWithContext(context.Background(), request)
}

// DescribeCCCBuyInfoList
// This API is used to access the user purchasing information list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXIST = "InvalidParameterValue.AccountNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCCCBuyInfoListWithContext(ctx context.Context, request *DescribeCCCBuyInfoListRequest) (response *DescribeCCCBuyInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeCCCBuyInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCCCBuyInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCCCBuyInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCallInMetricsRequest() (request *DescribeCallInMetricsRequest) {
    request = &DescribeCallInMetricsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeCallInMetrics")
    
    
    return
}

func NewDescribeCallInMetricsResponse() (response *DescribeCallInMetricsResponse) {
    response = &DescribeCallInMetricsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCallInMetrics
// This API is used to access the inbound real-time data statistical metrics.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeCallInMetrics(request *DescribeCallInMetricsRequest) (response *DescribeCallInMetricsResponse, err error) {
    return c.DescribeCallInMetricsWithContext(context.Background(), request)
}

// DescribeCallInMetrics
// This API is used to access the inbound real-time data statistical metrics.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeCallInMetricsWithContext(ctx context.Context, request *DescribeCallInMetricsRequest) (response *DescribeCallInMetricsResponse, err error) {
    if request == nil {
        request = NewDescribeCallInMetricsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCallInMetrics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCallInMetricsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExtensionRequest() (request *DescribeExtensionRequest) {
    request = &DescribeExtensionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeExtension")
    
    
    return
}

func NewDescribeExtensionResponse() (response *DescribeExtensionResponse) {
    response = &DescribeExtensionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeExtension
// This API is used to access the telephone information.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
func (c *Client) DescribeExtension(request *DescribeExtensionRequest) (response *DescribeExtensionResponse, err error) {
    return c.DescribeExtensionWithContext(context.Background(), request)
}

// DescribeExtension
// This API is used to access the telephone information.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
func (c *Client) DescribeExtensionWithContext(ctx context.Context, request *DescribeExtensionRequest) (response *DescribeExtensionResponse, err error) {
    if request == nil {
        request = NewDescribeExtensionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExtension require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExtensionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExtensionsRequest() (request *DescribeExtensionsRequest) {
    request = &DescribeExtensionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeExtensions")
    
    
    return
}

func NewDescribeExtensionsResponse() (response *DescribeExtensionsResponse) {
    response = &DescribeExtensionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeExtensions
// This API is used to query the telephone list information.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeExtensions(request *DescribeExtensionsRequest) (response *DescribeExtensionsResponse, err error) {
    return c.DescribeExtensionsWithContext(context.Background(), request)
}

// DescribeExtensions
// This API is used to query the telephone list information.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeExtensionsWithContext(ctx context.Context, request *DescribeExtensionsRequest) (response *DescribeExtensionsResponse, err error) {
    if request == nil {
        request = NewDescribeExtensionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExtensions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExtensionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIvrAudioListRequest() (request *DescribeIvrAudioListRequest) {
    request = &DescribeIvrAudioListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeIvrAudioList")
    
    
    return
}

func NewDescribeIvrAudioListResponse() (response *DescribeIvrAudioListResponse) {
    response = &DescribeIvrAudioListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIvrAudioList
// Query IVR Audio File List Information
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeIvrAudioList(request *DescribeIvrAudioListRequest) (response *DescribeIvrAudioListResponse, err error) {
    return c.DescribeIvrAudioListWithContext(context.Background(), request)
}

// DescribeIvrAudioList
// Query IVR Audio File List Information
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeIvrAudioListWithContext(ctx context.Context, request *DescribeIvrAudioListRequest) (response *DescribeIvrAudioListResponse, err error) {
    if request == nil {
        request = NewDescribeIvrAudioListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIvrAudioList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIvrAudioListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNumbersRequest() (request *DescribeNumbersRequest) {
    request = &DescribeNumbersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeNumbers")
    
    
    return
}

func NewDescribeNumbersResponse() (response *DescribeNumbersResponse) {
    response = &DescribeNumbersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNumbers
// This API is used to query the number list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeNumbers(request *DescribeNumbersRequest) (response *DescribeNumbersResponse, err error) {
    return c.DescribeNumbersWithContext(context.Background(), request)
}

// DescribeNumbers
// This API is used to query the number list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeNumbersWithContext(ctx context.Context, request *DescribeNumbersRequest) (response *DescribeNumbersResponse, err error) {
    if request == nil {
        request = NewDescribeNumbersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNumbers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNumbersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePSTNActiveSessionListRequest() (request *DescribePSTNActiveSessionListRequest) {
    request = &DescribePSTNActiveSessionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribePSTNActiveSessionList")
    
    
    return
}

func NewDescribePSTNActiveSessionListResponse() (response *DescribePSTNActiveSessionListResponse) {
    response = &DescribePSTNActiveSessionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePSTNActiveSessionList
// This API is used to access the current calling session list.
//
// error code that may be returned:
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_OUTOFCOUNTLIMIT = "LimitExceeded.OutOfCountLimit"
func (c *Client) DescribePSTNActiveSessionList(request *DescribePSTNActiveSessionListRequest) (response *DescribePSTNActiveSessionListResponse, err error) {
    return c.DescribePSTNActiveSessionListWithContext(context.Background(), request)
}

// DescribePSTNActiveSessionList
// This API is used to access the current calling session list.
//
// error code that may be returned:
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_OUTOFCOUNTLIMIT = "LimitExceeded.OutOfCountLimit"
func (c *Client) DescribePSTNActiveSessionListWithContext(ctx context.Context, request *DescribePSTNActiveSessionListRequest) (response *DescribePSTNActiveSessionListResponse, err error) {
    if request == nil {
        request = NewDescribePSTNActiveSessionListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePSTNActiveSessionList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePSTNActiveSessionListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePredictiveDialingCampaignRequest() (request *DescribePredictiveDialingCampaignRequest) {
    request = &DescribePredictiveDialingCampaignRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribePredictiveDialingCampaign")
    
    
    return
}

func NewDescribePredictiveDialingCampaignResponse() (response *DescribePredictiveDialingCampaignResponse) {
    response = &DescribePredictiveDialingCampaignResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePredictiveDialingCampaign
// This API is used to query the predictive outbound call task.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribePredictiveDialingCampaign(request *DescribePredictiveDialingCampaignRequest) (response *DescribePredictiveDialingCampaignResponse, err error) {
    return c.DescribePredictiveDialingCampaignWithContext(context.Background(), request)
}

// DescribePredictiveDialingCampaign
// This API is used to query the predictive outbound call task.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribePredictiveDialingCampaignWithContext(ctx context.Context, request *DescribePredictiveDialingCampaignRequest) (response *DescribePredictiveDialingCampaignResponse, err error) {
    if request == nil {
        request = NewDescribePredictiveDialingCampaignRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePredictiveDialingCampaign require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePredictiveDialingCampaignResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePredictiveDialingCampaignsRequest() (request *DescribePredictiveDialingCampaignsRequest) {
    request = &DescribePredictiveDialingCampaignsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribePredictiveDialingCampaigns")
    
    
    return
}

func NewDescribePredictiveDialingCampaignsResponse() (response *DescribePredictiveDialingCampaignsResponse) {
    response = &DescribePredictiveDialingCampaignsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePredictiveDialingCampaigns
// This API is used to query the predictive outbound call task list.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribePredictiveDialingCampaigns(request *DescribePredictiveDialingCampaignsRequest) (response *DescribePredictiveDialingCampaignsResponse, err error) {
    return c.DescribePredictiveDialingCampaignsWithContext(context.Background(), request)
}

// DescribePredictiveDialingCampaigns
// This API is used to query the predictive outbound call task list.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribePredictiveDialingCampaignsWithContext(ctx context.Context, request *DescribePredictiveDialingCampaignsRequest) (response *DescribePredictiveDialingCampaignsResponse, err error) {
    if request == nil {
        request = NewDescribePredictiveDialingCampaignsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePredictiveDialingCampaigns require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePredictiveDialingCampaignsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePredictiveDialingSessionsRequest() (request *DescribePredictiveDialingSessionsRequest) {
    request = &DescribePredictiveDialingSessionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribePredictiveDialingSessions")
    
    
    return
}

func NewDescribePredictiveDialingSessionsResponse() (response *DescribePredictiveDialingSessionsResponse) {
    response = &DescribePredictiveDialingSessionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePredictiveDialingSessions
// This API is used to query the predictive outbound call list.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribePredictiveDialingSessions(request *DescribePredictiveDialingSessionsRequest) (response *DescribePredictiveDialingSessionsResponse, err error) {
    return c.DescribePredictiveDialingSessionsWithContext(context.Background(), request)
}

// DescribePredictiveDialingSessions
// This API is used to query the predictive outbound call list.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribePredictiveDialingSessionsWithContext(ctx context.Context, request *DescribePredictiveDialingSessionsRequest) (response *DescribePredictiveDialingSessionsResponse, err error) {
    if request == nil {
        request = NewDescribePredictiveDialingSessionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePredictiveDialingSessions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePredictiveDialingSessionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProtectedTelCdrRequest() (request *DescribeProtectedTelCdrRequest) {
    request = &DescribeProtectedTelCdrRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeProtectedTelCdr")
    
    
    return
}

func NewDescribeProtectedTelCdrResponse() (response *DescribeProtectedTelCdrResponse) {
    response = &DescribeProtectedTelCdrResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProtectedTelCdr
// This API is used to access protected phone service records and recordings for both inbound and outbound calls.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeProtectedTelCdr(request *DescribeProtectedTelCdrRequest) (response *DescribeProtectedTelCdrResponse, err error) {
    return c.DescribeProtectedTelCdrWithContext(context.Background(), request)
}

// DescribeProtectedTelCdr
// This API is used to access protected phone service records and recordings for both inbound and outbound calls.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeProtectedTelCdrWithContext(ctx context.Context, request *DescribeProtectedTelCdrRequest) (response *DescribeProtectedTelCdrResponse, err error) {
    if request == nil {
        request = NewDescribeProtectedTelCdrRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProtectedTelCdr require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProtectedTelCdrResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSkillGroupInfoListRequest() (request *DescribeSkillGroupInfoListRequest) {
    request = &DescribeSkillGroupInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeSkillGroupInfoList")
    
    
    return
}

func NewDescribeSkillGroupInfoListResponse() (response *DescribeSkillGroupInfoListResponse) {
    response = &DescribeSkillGroupInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSkillGroupInfoList
// This API is used to access the skill group information list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeSkillGroupInfoList(request *DescribeSkillGroupInfoListRequest) (response *DescribeSkillGroupInfoListResponse, err error) {
    return c.DescribeSkillGroupInfoListWithContext(context.Background(), request)
}

// DescribeSkillGroupInfoList
// This API is used to access the skill group information list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeSkillGroupInfoListWithContext(ctx context.Context, request *DescribeSkillGroupInfoListRequest) (response *DescribeSkillGroupInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeSkillGroupInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSkillGroupInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSkillGroupInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStaffInfoListRequest() (request *DescribeStaffInfoListRequest) {
    request = &DescribeStaffInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeStaffInfoList")
    
    
    return
}

func NewDescribeStaffInfoListResponse() (response *DescribeStaffInfoListResponse) {
    response = &DescribeStaffInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStaffInfoList
// This API is used to access the agent information list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeStaffInfoList(request *DescribeStaffInfoListRequest) (response *DescribeStaffInfoListResponse, err error) {
    return c.DescribeStaffInfoListWithContext(context.Background(), request)
}

// DescribeStaffInfoList
// This API is used to access the agent information list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeStaffInfoListWithContext(ctx context.Context, request *DescribeStaffInfoListRequest) (response *DescribeStaffInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeStaffInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStaffInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStaffInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStaffStatusMetricsRequest() (request *DescribeStaffStatusMetricsRequest) {
    request = &DescribeStaffStatusMetricsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeStaffStatusMetrics")
    
    
    return
}

func NewDescribeStaffStatusMetricsResponse() (response *DescribeStaffStatusMetricsResponse) {
    response = &DescribeStaffStatusMetricsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStaffStatusMetrics
// This API is used to access the real-time status statistics metrics of the agent.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeStaffStatusMetrics(request *DescribeStaffStatusMetricsRequest) (response *DescribeStaffStatusMetricsResponse, err error) {
    return c.DescribeStaffStatusMetricsWithContext(context.Background(), request)
}

// DescribeStaffStatusMetrics
// This API is used to access the real-time status statistics metrics of the agent.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeStaffStatusMetricsWithContext(ctx context.Context, request *DescribeStaffStatusMetricsRequest) (response *DescribeStaffStatusMetricsResponse, err error) {
    if request == nil {
        request = NewDescribeStaffStatusMetricsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStaffStatusMetrics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStaffStatusMetricsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTelCallInfoRequest() (request *DescribeTelCallInfoRequest) {
    request = &DescribeTelCallInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeTelCallInfo")
    
    
    return
}

func NewDescribeTelCallInfoResponse() (response *DescribeTelCallInfoResponse) {
    response = &DescribeTelCallInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTelCallInfo
// This API is used to access call detail records by application
//
// error code that may be returned:
//  FAILEDOPERATION_DUPLICATEDACCOUNT = "FailedOperation.DuplicatedAccount"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeTelCallInfo(request *DescribeTelCallInfoRequest) (response *DescribeTelCallInfoResponse, err error) {
    return c.DescribeTelCallInfoWithContext(context.Background(), request)
}

// DescribeTelCallInfo
// This API is used to access call detail records by application
//
// error code that may be returned:
//  FAILEDOPERATION_DUPLICATEDACCOUNT = "FailedOperation.DuplicatedAccount"
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeTelCallInfoWithContext(ctx context.Context, request *DescribeTelCallInfoRequest) (response *DescribeTelCallInfoResponse, err error) {
    if request == nil {
        request = NewDescribeTelCallInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTelCallInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTelCallInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTelCdrRequest() (request *DescribeTelCdrRequest) {
    request = &DescribeTelCdrRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeTelCdr")
    
    
    return
}

func NewDescribeTelCdrResponse() (response *DescribeTelCdrResponse) {
    response = &DescribeTelCdrResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTelCdr
// This API is used to access phone service records and recordings.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeTelCdr(request *DescribeTelCdrRequest) (response *DescribeTelCdrResponse, err error) {
    return c.DescribeTelCdrWithContext(context.Background(), request)
}

// DescribeTelCdr
// This API is used to access phone service records and recordings.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeTelCdrWithContext(ctx context.Context, request *DescribeTelCdrRequest) (response *DescribeTelCdrResponse, err error) {
    if request == nil {
        request = NewDescribeTelCdrRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTelCdr require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTelCdrResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTelSessionRequest() (request *DescribeTelSessionRequest) {
    request = &DescribeTelSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeTelSession")
    
    
    return
}

func NewDescribeTelSessionResponse() (response *DescribeTelSessionResponse) {
    response = &DescribeTelSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTelSession
// This API is used to access the PSTN session information.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeTelSession(request *DescribeTelSessionRequest) (response *DescribeTelSessionResponse, err error) {
    return c.DescribeTelSessionWithContext(context.Background(), request)
}

// DescribeTelSession
// This API is used to access the PSTN session information.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeTelSessionWithContext(ctx context.Context, request *DescribeTelSessionRequest) (response *DescribeTelSessionResponse, err error) {
    if request == nil {
        request = NewDescribeTelSessionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTelSession require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTelSessionResponse()
    err = c.Send(request, response)
    return
}

func NewDisableCCCPhoneNumberRequest() (request *DisableCCCPhoneNumberRequest) {
    request = &DisableCCCPhoneNumberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "DisableCCCPhoneNumber")
    
    
    return
}

func NewDisableCCCPhoneNumberResponse() (response *DisableCCCPhoneNumberResponse) {
    response = &DisableCCCPhoneNumberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisableCCCPhoneNumber
// This API is used to disable numbers.
//
// error code that may be returned:
//  FAILEDOPERATION_CURSTATENOTALLOWMODIFY = "FailedOperation.CurStateNotAllowModify"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_DUPLICATEADDRESS = "InvalidParameter.DuplicateAddress"
//  INVALIDPARAMETER_DUPLICATEPHONENUMBER = "InvalidParameter.DuplicatePhoneNumber"
//  INVALIDPARAMETER_DUPLICATESIPACCOUNT = "InvalidParameter.DuplicateSipAccount"
//  INVALIDPARAMETER_ILLEGALADDRESS = "InvalidParameter.IllegalAddress"
//  INVALIDPARAMETER_ILLEGALPHONENUMBER = "InvalidParameter.IllegalPhoneNumber"
//  INVALIDPARAMETER_INVALIDADDRESS = "InvalidParameter.InvalidAddress"
//  INVALIDPARAMETER_INVALIDIP = "InvalidParameter.InvalidIP"
//  INVALIDPARAMETER_INVALIDPHONENUMBER = "InvalidParameter.InvalidPhoneNumber"
//  INVALIDPARAMETER_INVALIDPORT = "InvalidParameter.InvalidPort"
//  INVALIDPARAMETER_SIPACCOUNTPASSWORDFORMAT = "InvalidParameter.SipAccountPasswordFormat"
//  INVALIDPARAMETER_SIPACCOUNTUSERFORMAT = "InvalidParameter.SipAccountUserFormat"
//  INVALIDPARAMETER_SIPTRUNKINUSED = "InvalidParameter.SipTrunkInUsed"
//  INVALIDPARAMETER_SIPTRUNKNOTFOUND = "InvalidParameter.SipTrunkNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DisableCCCPhoneNumber(request *DisableCCCPhoneNumberRequest) (response *DisableCCCPhoneNumberResponse, err error) {
    return c.DisableCCCPhoneNumberWithContext(context.Background(), request)
}

// DisableCCCPhoneNumber
// This API is used to disable numbers.
//
// error code that may be returned:
//  FAILEDOPERATION_CURSTATENOTALLOWMODIFY = "FailedOperation.CurStateNotAllowModify"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_DUPLICATEADDRESS = "InvalidParameter.DuplicateAddress"
//  INVALIDPARAMETER_DUPLICATEPHONENUMBER = "InvalidParameter.DuplicatePhoneNumber"
//  INVALIDPARAMETER_DUPLICATESIPACCOUNT = "InvalidParameter.DuplicateSipAccount"
//  INVALIDPARAMETER_ILLEGALADDRESS = "InvalidParameter.IllegalAddress"
//  INVALIDPARAMETER_ILLEGALPHONENUMBER = "InvalidParameter.IllegalPhoneNumber"
//  INVALIDPARAMETER_INVALIDADDRESS = "InvalidParameter.InvalidAddress"
//  INVALIDPARAMETER_INVALIDIP = "InvalidParameter.InvalidIP"
//  INVALIDPARAMETER_INVALIDPHONENUMBER = "InvalidParameter.InvalidPhoneNumber"
//  INVALIDPARAMETER_INVALIDPORT = "InvalidParameter.InvalidPort"
//  INVALIDPARAMETER_SIPACCOUNTPASSWORDFORMAT = "InvalidParameter.SipAccountPasswordFormat"
//  INVALIDPARAMETER_SIPACCOUNTUSERFORMAT = "InvalidParameter.SipAccountUserFormat"
//  INVALIDPARAMETER_SIPTRUNKINUSED = "InvalidParameter.SipTrunkInUsed"
//  INVALIDPARAMETER_SIPTRUNKNOTFOUND = "InvalidParameter.SipTrunkNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DisableCCCPhoneNumberWithContext(ctx context.Context, request *DisableCCCPhoneNumberRequest) (response *DisableCCCPhoneNumberResponse, err error) {
    if request == nil {
        request = NewDisableCCCPhoneNumberRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableCCCPhoneNumber require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableCCCPhoneNumberResponse()
    err = c.Send(request, response)
    return
}

func NewHangUpCallRequest() (request *HangUpCallRequest) {
    request = &HangUpCallRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "HangUpCall")
    
    
    return
}

func NewHangUpCallResponse() (response *HangUpCallResponse) {
    response = &HangUpCallResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// HangUpCall
// This API is used to hang up the phone.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
func (c *Client) HangUpCall(request *HangUpCallRequest) (response *HangUpCallResponse, err error) {
    return c.HangUpCallWithContext(context.Background(), request)
}

// HangUpCall
// This API is used to hang up the phone.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
func (c *Client) HangUpCallWithContext(ctx context.Context, request *HangUpCallRequest) (response *HangUpCallResponse, err error) {
    if request == nil {
        request = NewHangUpCallRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("HangUpCall require credential")
    }

    request.SetContext(ctx)
    
    response = NewHangUpCallResponse()
    err = c.Send(request, response)
    return
}

func NewModifyExtensionRequest() (request *ModifyExtensionRequest) {
    request = &ModifyExtensionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "ModifyExtension")
    
    
    return
}

func NewModifyExtensionResponse() (response *ModifyExtensionResponse) {
    response = &ModifyExtensionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyExtension
// This API is used to modify telephone accounts (binding skill group, binding agent account).
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) ModifyExtension(request *ModifyExtensionRequest) (response *ModifyExtensionResponse, err error) {
    return c.ModifyExtensionWithContext(context.Background(), request)
}

// ModifyExtension
// This API is used to modify telephone accounts (binding skill group, binding agent account).
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) ModifyExtensionWithContext(ctx context.Context, request *ModifyExtensionRequest) (response *ModifyExtensionResponse, err error) {
    if request == nil {
        request = NewModifyExtensionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyExtension require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyExtensionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyStaffRequest() (request *ModifyStaffRequest) {
    request = &ModifyStaffRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "ModifyStaff")
    
    
    return
}

func NewModifyStaffResponse() (response *ModifyStaffResponse) {
    response = &ModifyStaffResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyStaff
// This API is used to modify  customer service / agent account.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXIST = "InvalidParameterValue.AccountNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_PHONENUMISBOUNDOTHERACCOUNT = "InvalidParameterValue.PhoneNumIsBoundOtherAccount"
//  INVALIDPARAMETERVALUE_SKILLGROUPERROR = "InvalidParameterValue.SkillGroupError"
//  OPERATIONDENIED_NOTINWHITELIST = "OperationDenied.NotInWhiteList"
func (c *Client) ModifyStaff(request *ModifyStaffRequest) (response *ModifyStaffResponse, err error) {
    return c.ModifyStaffWithContext(context.Background(), request)
}

// ModifyStaff
// This API is used to modify  customer service / agent account.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXIST = "InvalidParameterValue.AccountNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_PHONENUMISBOUNDOTHERACCOUNT = "InvalidParameterValue.PhoneNumIsBoundOtherAccount"
//  INVALIDPARAMETERVALUE_SKILLGROUPERROR = "InvalidParameterValue.SkillGroupError"
//  OPERATIONDENIED_NOTINWHITELIST = "OperationDenied.NotInWhiteList"
func (c *Client) ModifyStaffWithContext(ctx context.Context, request *ModifyStaffRequest) (response *ModifyStaffResponse, err error) {
    if request == nil {
        request = NewModifyStaffRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyStaff require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyStaffResponse()
    err = c.Send(request, response)
    return
}

func NewModifyStaffPasswordRequest() (request *ModifyStaffPasswordRequest) {
    request = &ModifyStaffPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "ModifyStaffPassword")
    
    
    return
}

func NewModifyStaffPasswordResponse() (response *ModifyStaffPasswordResponse) {
    response = &ModifyStaffPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyStaffPassword
// Modify Agent's Password
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXIST = "InvalidParameterValue.AccountNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) ModifyStaffPassword(request *ModifyStaffPasswordRequest) (response *ModifyStaffPasswordResponse, err error) {
    return c.ModifyStaffPasswordWithContext(context.Background(), request)
}

// ModifyStaffPassword
// Modify Agent's Password
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXIST = "InvalidParameterValue.AccountNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) ModifyStaffPasswordWithContext(ctx context.Context, request *ModifyStaffPasswordRequest) (response *ModifyStaffPasswordResponse, err error) {
    if request == nil {
        request = NewModifyStaffPasswordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyStaffPassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyStaffPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewPausePredictiveDialingCampaignRequest() (request *PausePredictiveDialingCampaignRequest) {
    request = &PausePredictiveDialingCampaignRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "PausePredictiveDialingCampaign")
    
    
    return
}

func NewPausePredictiveDialingCampaignResponse() (response *PausePredictiveDialingCampaignResponse) {
    response = &PausePredictiveDialingCampaignResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PausePredictiveDialingCampaign
// This API is used to pause the predictive outbound call task.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) PausePredictiveDialingCampaign(request *PausePredictiveDialingCampaignRequest) (response *PausePredictiveDialingCampaignResponse, err error) {
    return c.PausePredictiveDialingCampaignWithContext(context.Background(), request)
}

// PausePredictiveDialingCampaign
// This API is used to pause the predictive outbound call task.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) PausePredictiveDialingCampaignWithContext(ctx context.Context, request *PausePredictiveDialingCampaignRequest) (response *PausePredictiveDialingCampaignResponse, err error) {
    if request == nil {
        request = NewPausePredictiveDialingCampaignRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PausePredictiveDialingCampaign require credential")
    }

    request.SetContext(ctx)
    
    response = NewPausePredictiveDialingCampaignResponse()
    err = c.Send(request, response)
    return
}

func NewResetExtensionPasswordRequest() (request *ResetExtensionPasswordRequest) {
    request = &ResetExtensionPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "ResetExtensionPassword")
    
    
    return
}

func NewResetExtensionPasswordResponse() (response *ResetExtensionPasswordResponse) {
    response = &ResetExtensionPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetExtensionPassword
// This API is used to reset the telephone register password.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ResetExtensionPassword(request *ResetExtensionPasswordRequest) (response *ResetExtensionPasswordResponse, err error) {
    return c.ResetExtensionPasswordWithContext(context.Background(), request)
}

// ResetExtensionPassword
// This API is used to reset the telephone register password.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ResetExtensionPasswordWithContext(ctx context.Context, request *ResetExtensionPasswordRequest) (response *ResetExtensionPasswordResponse, err error) {
    if request == nil {
        request = NewResetExtensionPasswordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetExtensionPassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetExtensionPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewResumePredictiveDialingCampaignRequest() (request *ResumePredictiveDialingCampaignRequest) {
    request = &ResumePredictiveDialingCampaignRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "ResumePredictiveDialingCampaign")
    
    
    return
}

func NewResumePredictiveDialingCampaignResponse() (response *ResumePredictiveDialingCampaignResponse) {
    response = &ResumePredictiveDialingCampaignResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResumePredictiveDialingCampaign
// This API is used to resume the predictive outbound call task.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ResumePredictiveDialingCampaign(request *ResumePredictiveDialingCampaignRequest) (response *ResumePredictiveDialingCampaignResponse, err error) {
    return c.ResumePredictiveDialingCampaignWithContext(context.Background(), request)
}

// ResumePredictiveDialingCampaign
// This API is used to resume the predictive outbound call task.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ResumePredictiveDialingCampaignWithContext(ctx context.Context, request *ResumePredictiveDialingCampaignRequest) (response *ResumePredictiveDialingCampaignResponse, err error) {
    if request == nil {
        request = NewResumePredictiveDialingCampaignRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResumePredictiveDialingCampaign require credential")
    }

    request.SetContext(ctx)
    
    response = NewResumePredictiveDialingCampaignResponse()
    err = c.Send(request, response)
    return
}

func NewStopAutoCalloutTaskRequest() (request *StopAutoCalloutTaskRequest) {
    request = &StopAutoCalloutTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "StopAutoCalloutTask")
    
    
    return
}

func NewStopAutoCalloutTaskResponse() (response *StopAutoCalloutTaskResponse) {
    response = &StopAutoCalloutTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopAutoCalloutTask
// This API is used to stop the automatic outbound call task.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StopAutoCalloutTask(request *StopAutoCalloutTaskRequest) (response *StopAutoCalloutTaskResponse, err error) {
    return c.StopAutoCalloutTaskWithContext(context.Background(), request)
}

// StopAutoCalloutTask
// This API is used to stop the automatic outbound call task.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StopAutoCalloutTaskWithContext(ctx context.Context, request *StopAutoCalloutTaskRequest) (response *StopAutoCalloutTaskResponse, err error) {
    if request == nil {
        request = NewStopAutoCalloutTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopAutoCalloutTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopAutoCalloutTaskResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindNumberCallOutSkillGroupRequest() (request *UnbindNumberCallOutSkillGroupRequest) {
    request = &UnbindNumberCallOutSkillGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "UnbindNumberCallOutSkillGroup")
    
    
    return
}

func NewUnbindNumberCallOutSkillGroupResponse() (response *UnbindNumberCallOutSkillGroupResponse) {
    response = &UnbindNumberCallOutSkillGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnbindNumberCallOutSkillGroup
// This API is used to unbind the number from the outbound call skill group.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UnbindNumberCallOutSkillGroup(request *UnbindNumberCallOutSkillGroupRequest) (response *UnbindNumberCallOutSkillGroupResponse, err error) {
    return c.UnbindNumberCallOutSkillGroupWithContext(context.Background(), request)
}

// UnbindNumberCallOutSkillGroup
// This API is used to unbind the number from the outbound call skill group.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UnbindNumberCallOutSkillGroupWithContext(ctx context.Context, request *UnbindNumberCallOutSkillGroupRequest) (response *UnbindNumberCallOutSkillGroupResponse, err error) {
    if request == nil {
        request = NewUnbindNumberCallOutSkillGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindNumberCallOutSkillGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindNumberCallOutSkillGroupResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindStaffSkillGroupListRequest() (request *UnbindStaffSkillGroupListRequest) {
    request = &UnbindStaffSkillGroupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "UnbindStaffSkillGroupList")
    
    
    return
}

func NewUnbindStaffSkillGroupListResponse() (response *UnbindStaffSkillGroupListResponse) {
    response = &UnbindStaffSkillGroupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnbindStaffSkillGroupList
// This API is used to unbind the agent's skill group.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) UnbindStaffSkillGroupList(request *UnbindStaffSkillGroupListRequest) (response *UnbindStaffSkillGroupListResponse, err error) {
    return c.UnbindStaffSkillGroupListWithContext(context.Background(), request)
}

// UnbindStaffSkillGroupList
// This API is used to unbind the agent's skill group.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) UnbindStaffSkillGroupListWithContext(ctx context.Context, request *UnbindStaffSkillGroupListRequest) (response *UnbindStaffSkillGroupListResponse, err error) {
    if request == nil {
        request = NewUnbindStaffSkillGroupListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindStaffSkillGroupList require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindStaffSkillGroupListResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCCCSkillGroupRequest() (request *UpdateCCCSkillGroupRequest) {
    request = &UpdateCCCSkillGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "UpdateCCCSkillGroup")
    
    
    return
}

func NewUpdateCCCSkillGroupResponse() (response *UpdateCCCSkillGroupResponse) {
    response = &UpdateCCCSkillGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateCCCSkillGroup
// This API is used to update the skill group.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_SKILLGROUPEXIST = "InvalidParameterValue.SkillGroupExist"
func (c *Client) UpdateCCCSkillGroup(request *UpdateCCCSkillGroupRequest) (response *UpdateCCCSkillGroupResponse, err error) {
    return c.UpdateCCCSkillGroupWithContext(context.Background(), request)
}

// UpdateCCCSkillGroup
// This API is used to update the skill group.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_SKILLGROUPEXIST = "InvalidParameterValue.SkillGroupExist"
func (c *Client) UpdateCCCSkillGroupWithContext(ctx context.Context, request *UpdateCCCSkillGroupRequest) (response *UpdateCCCSkillGroupResponse, err error) {
    if request == nil {
        request = NewUpdateCCCSkillGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCCCSkillGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCCCSkillGroupResponse()
    err = c.Send(request, response)
    return
}

func NewUpdatePredictiveDialingCampaignRequest() (request *UpdatePredictiveDialingCampaignRequest) {
    request = &UpdatePredictiveDialingCampaignRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "UpdatePredictiveDialingCampaign")
    
    
    return
}

func NewUpdatePredictiveDialingCampaignResponse() (response *UpdatePredictiveDialingCampaignResponse) {
    response = &UpdatePredictiveDialingCampaignResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdatePredictiveDialingCampaign
// This API is used to update the predictive outbound call task before it starts.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdatePredictiveDialingCampaign(request *UpdatePredictiveDialingCampaignRequest) (response *UpdatePredictiveDialingCampaignResponse, err error) {
    return c.UpdatePredictiveDialingCampaignWithContext(context.Background(), request)
}

// UpdatePredictiveDialingCampaign
// This API is used to update the predictive outbound call task before it starts.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdatePredictiveDialingCampaignWithContext(ctx context.Context, request *UpdatePredictiveDialingCampaignRequest) (response *UpdatePredictiveDialingCampaignResponse, err error) {
    if request == nil {
        request = NewUpdatePredictiveDialingCampaignRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdatePredictiveDialingCampaign require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdatePredictiveDialingCampaignResponse()
    err = c.Send(request, response)
    return
}

func NewUploadIvrAudioRequest() (request *UploadIvrAudioRequest) {
    request = &UploadIvrAudioRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ccc", APIVersion, "UploadIvrAudio")
    
    
    return
}

func NewUploadIvrAudioResponse() (response *UploadIvrAudioResponse) {
    response = &UploadIvrAudioResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UploadIvrAudio
// Upload audio files used in IVR, with a daily upload limit of 50 files. (It is recommended to use temporary links stored in Tencent Cloud Cos for the audio file URL in the parameters)
//
// error code that may be returned:
//  FAILEDOPERATION_UPLOADFILEOVERFLOW = "FailedOperation.UploadFileOverflow"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UploadIvrAudio(request *UploadIvrAudioRequest) (response *UploadIvrAudioResponse, err error) {
    return c.UploadIvrAudioWithContext(context.Background(), request)
}

// UploadIvrAudio
// Upload audio files used in IVR, with a daily upload limit of 50 files. (It is recommended to use temporary links stored in Tencent Cloud Cos for the audio file URL in the parameters)
//
// error code that may be returned:
//  FAILEDOPERATION_UPLOADFILEOVERFLOW = "FailedOperation.UploadFileOverflow"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UploadIvrAudioWithContext(ctx context.Context, request *UploadIvrAudioRequest) (response *UploadIvrAudioResponse, err error) {
    if request == nil {
        request = NewUploadIvrAudioRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadIvrAudio require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadIvrAudioResponse()
    err = c.Send(request, response)
    return
}
