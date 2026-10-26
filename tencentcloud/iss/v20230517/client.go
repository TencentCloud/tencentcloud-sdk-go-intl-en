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

package v20230517

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2023-05-17"

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


func NewAddOrganizationRequest() (request *AddOrganizationRequest) {
    request = &AddOrganizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "AddOrganization")
    
    
    return
}

func NewAddOrganizationResponse() (response *AddOrganizationResponse) {
    response = &AddOrganizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddOrganization
// Used to add an organization.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDORGNAME = "InvalidParameterValue.InvalidOrgName"
//  INVALIDPARAMETERVALUE_ORGNAMEREPEAT = "InvalidParameterValue.OrgNameRepeat"
//  OPERATIONDENIED_EXCEEDEDMAXIMUMDEPTH = "OperationDenied.ExceededMaximumDepth"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddOrganization(request *AddOrganizationRequest) (response *AddOrganizationResponse, err error) {
    return c.AddOrganizationWithContext(context.Background(), request)
}

// AddOrganization
// Used to add an organization.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDORGNAME = "InvalidParameterValue.InvalidOrgName"
//  INVALIDPARAMETERVALUE_ORGNAMEREPEAT = "InvalidParameterValue.OrgNameRepeat"
//  OPERATIONDENIED_EXCEEDEDMAXIMUMDEPTH = "OperationDenied.ExceededMaximumDepth"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddOrganizationWithContext(ctx context.Context, request *AddOrganizationRequest) (response *AddOrganizationResponse, err error) {
    if request == nil {
        request = NewAddOrganizationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "AddOrganization")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddOrganization require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddOrganizationResponse()
    err = c.Send(request, response)
    return
}

func NewAddRecordBackupPlanRequest() (request *AddRecordBackupPlanRequest) {
    request = &AddRecordBackupPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "AddRecordBackupPlan")
    
    
    return
}

func NewAddRecordBackupPlanResponse() (response *AddRecordBackupPlanResponse) {
    response = &AddRecordBackupPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddRecordBackupPlan
// Used to add a recording cloud migration plan (currently applicable only to devices/video channels connected via the GB28181 protocol and gateway).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDNAME = "InvalidParameterValue.InvalidName"
//  INVALIDPARAMETERVALUE_INVALIDORGANIZATIONID = "InvalidParameterValue.InvalidOrganizationId"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  INVALIDPARAMETERVALUE_PLANNAMEREPEAT = "InvalidParameterValue.PlanNameRepeat"
//  RESOURCEINUSE_PLANNAMEREPEAT = "ResourceInUse.PlanNameRepeat"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddRecordBackupPlan(request *AddRecordBackupPlanRequest) (response *AddRecordBackupPlanResponse, err error) {
    return c.AddRecordBackupPlanWithContext(context.Background(), request)
}

// AddRecordBackupPlan
// Used to add a recording cloud migration plan (currently applicable only to devices/video channels connected via the GB28181 protocol and gateway).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDNAME = "InvalidParameterValue.InvalidName"
//  INVALIDPARAMETERVALUE_INVALIDORGANIZATIONID = "InvalidParameterValue.InvalidOrganizationId"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  INVALIDPARAMETERVALUE_PLANNAMEREPEAT = "InvalidParameterValue.PlanNameRepeat"
//  RESOURCEINUSE_PLANNAMEREPEAT = "ResourceInUse.PlanNameRepeat"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddRecordBackupPlanWithContext(ctx context.Context, request *AddRecordBackupPlanRequest) (response *AddRecordBackupPlanResponse, err error) {
    if request == nil {
        request = NewAddRecordBackupPlanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "AddRecordBackupPlan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddRecordBackupPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddRecordBackupPlanResponse()
    err = c.Send(request, response)
    return
}

func NewAddRecordBackupTemplateRequest() (request *AddRecordBackupTemplateRequest) {
    request = &AddRecordBackupTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "AddRecordBackupTemplate")
    
    
    return
}

func NewAddRecordBackupTemplateResponse() (response *AddRecordBackupTemplateResponse) {
    response = &AddRecordBackupTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddRecordBackupTemplate
// This API is used to create a video recording cloud template.
//
// > The feature essentially pulls local device recording data to the cloud (i.e., recordings stored in the IPC camera memory card or NVR hard disk). When operating, first set the recording time period (i.e., the local device recording you want to upload to the cloud), then set the Cloud Upload Time Period and upload speed. The platform will pull the previous day's Time Period Data from the device at the specified speed during the Cloud Upload Time Period.
//
// 
//
// > Set the Cloud Upload Time Period to at least meet (Cloud Upload Time Period = previous day's recording time period / cloud upload speed multiplication). It is advisable to set about 10% more time for the Cloud Upload Time Period to avoid incomplete data retrieval due to network fluctuations.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BAKTIMENOTENOUGH = "InvalidParameterValue.BakTimeNotEnough"
//  INVALIDPARAMETERVALUE_DATAOUTOFRANGE = "InvalidParameterValue.DataOutOfRange"
//  INVALIDPARAMETERVALUE_INVALIDNAME = "InvalidParameterValue.InvalidName"
//  INVALIDPARAMETERVALUE_INVALIDTIMEFORMAT = "InvalidParameterValue.InvalidTimeFormat"
//  INVALIDPARAMETERVALUE_STARTOVERENDTIME = "InvalidParameterValue.StartOverEndTime"
//  INVALIDPARAMETERVALUE_TIMELESSTHANTENMINUTES = "InvalidParameterValue.TimeLessThanTenMinutes"
//  INVALIDPARAMETERVALUE_UNSUPPORTSCALE = "InvalidParameterValue.UnSupportScale"
//  RESOURCEINUSE_TEMPLATENAMEREPEAT = "ResourceInUse.TemplateNameRepeat"
func (c *Client) AddRecordBackupTemplate(request *AddRecordBackupTemplateRequest) (response *AddRecordBackupTemplateResponse, err error) {
    return c.AddRecordBackupTemplateWithContext(context.Background(), request)
}

// AddRecordBackupTemplate
// This API is used to create a video recording cloud template.
//
// > The feature essentially pulls local device recording data to the cloud (i.e., recordings stored in the IPC camera memory card or NVR hard disk). When operating, first set the recording time period (i.e., the local device recording you want to upload to the cloud), then set the Cloud Upload Time Period and upload speed. The platform will pull the previous day's Time Period Data from the device at the specified speed during the Cloud Upload Time Period.
//
// 
//
// > Set the Cloud Upload Time Period to at least meet (Cloud Upload Time Period = previous day's recording time period / cloud upload speed multiplication). It is advisable to set about 10% more time for the Cloud Upload Time Period to avoid incomplete data retrieval due to network fluctuations.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BAKTIMENOTENOUGH = "InvalidParameterValue.BakTimeNotEnough"
//  INVALIDPARAMETERVALUE_DATAOUTOFRANGE = "InvalidParameterValue.DataOutOfRange"
//  INVALIDPARAMETERVALUE_INVALIDNAME = "InvalidParameterValue.InvalidName"
//  INVALIDPARAMETERVALUE_INVALIDTIMEFORMAT = "InvalidParameterValue.InvalidTimeFormat"
//  INVALIDPARAMETERVALUE_STARTOVERENDTIME = "InvalidParameterValue.StartOverEndTime"
//  INVALIDPARAMETERVALUE_TIMELESSTHANTENMINUTES = "InvalidParameterValue.TimeLessThanTenMinutes"
//  INVALIDPARAMETERVALUE_UNSUPPORTSCALE = "InvalidParameterValue.UnSupportScale"
//  RESOURCEINUSE_TEMPLATENAMEREPEAT = "ResourceInUse.TemplateNameRepeat"
func (c *Client) AddRecordBackupTemplateWithContext(ctx context.Context, request *AddRecordBackupTemplateRequest) (response *AddRecordBackupTemplateResponse, err error) {
    if request == nil {
        request = NewAddRecordBackupTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "AddRecordBackupTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddRecordBackupTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddRecordBackupTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewAddRecordPlanRequest() (request *AddRecordPlanRequest) {
    request = &AddRecordPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "AddRecordPlan")
    
    
    return
}

func NewAddRecordPlanResponse() (response *AddRecordPlanResponse) {
    response = &AddRecordPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddRecordPlan
// This API is used to create a real-time cloud migration plan.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETER_INVALIDCHANNELS = "InvalidParameter.InvalidChannels"
//  INVALIDPARAMETER_INVALIDORGANIZATIONPARAM = "InvalidParameter.InvalidOrganizationParam"
//  INVALIDPARAMETERVALUE_EMPTYCHANNELID = "InvalidParameterValue.EmptyChannelId"
//  INVALIDPARAMETERVALUE_EMPTYDEVICEID = "InvalidParameterValue.EmptyDeviceId"
//  INVALIDPARAMETERVALUE_EMPTYNAME = "InvalidParameterValue.EmptyName"
//  INVALIDPARAMETERVALUE_EMPTYORGANIZATIONID = "InvalidParameterValue.EmptyOrganizationId"
//  INVALIDPARAMETERVALUE_EMPTYTEMPLATEID = "InvalidParameterValue.EmptyTemplateId"
//  INVALIDPARAMETERVALUE_ILLEGALCHANNELID = "InvalidParameterValue.IllegalChannelId"
//  INVALIDPARAMETERVALUE_ILLEGALDESCRIBE = "InvalidParameterValue.IllegalDescribe"
//  INVALIDPARAMETERVALUE_ILLEGALDEVICEID = "InvalidParameterValue.IllegalDeviceId"
//  INVALIDPARAMETERVALUE_ILLEGALNAME = "InvalidParameterValue.IllegalName"
//  INVALIDPARAMETERVALUE_ILLEGALORGANIZATIONID = "InvalidParameterValue.IllegalOrganizationId"
//  INVALIDPARAMETERVALUE_ILLEGALSTREAMTYPE = "InvalidParameterValue.IllegalStreamType"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELIDORDEVICEID = "InvalidParameterValue.InvalidChannelIdOrDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDREPAIRMODE = "InvalidParameterValue.InvalidRepairMode"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  INVALIDPARAMETERVALUE_LIMITREPAIRTRANSITION = "InvalidParameterValue.LimitRepairTransition"
//  INVALIDPARAMETERVALUE_ORGANIZATIONCOUNTEXCEEDSRANGE = "InvalidParameterValue.OrganizationCountExceedsRange"
//  INVALIDPARAMETERVALUE_PLANCHANNELSEXCEEDSRANGE = "InvalidParameterValue.PlanChannelsExceedsRange"
//  INVALIDPARAMETERVALUE_TOOLONGDESCRIBE = "InvalidParameterValue.TooLongDescribe"
//  INVALIDPARAMETERVALUE_TOOLONGNAME = "InvalidParameterValue.TooLongName"
//  INVALIDPARAMETERVALUE_TOOLONGSTREAMTYPE = "InvalidParameterValue.TooLongStreamType"
//  RESOURCEINUSE_PLANNAMEREPEAT = "ResourceInUse.PlanNameRepeat"
//  RESOURCENOTFOUND_ORGANIZATIONIDNOTEXIST = "ResourceNotFound.OrganizationIdNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) AddRecordPlan(request *AddRecordPlanRequest) (response *AddRecordPlanResponse, err error) {
    return c.AddRecordPlanWithContext(context.Background(), request)
}

// AddRecordPlan
// This API is used to create a real-time cloud migration plan.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETER_INVALIDCHANNELS = "InvalidParameter.InvalidChannels"
//  INVALIDPARAMETER_INVALIDORGANIZATIONPARAM = "InvalidParameter.InvalidOrganizationParam"
//  INVALIDPARAMETERVALUE_EMPTYCHANNELID = "InvalidParameterValue.EmptyChannelId"
//  INVALIDPARAMETERVALUE_EMPTYDEVICEID = "InvalidParameterValue.EmptyDeviceId"
//  INVALIDPARAMETERVALUE_EMPTYNAME = "InvalidParameterValue.EmptyName"
//  INVALIDPARAMETERVALUE_EMPTYORGANIZATIONID = "InvalidParameterValue.EmptyOrganizationId"
//  INVALIDPARAMETERVALUE_EMPTYTEMPLATEID = "InvalidParameterValue.EmptyTemplateId"
//  INVALIDPARAMETERVALUE_ILLEGALCHANNELID = "InvalidParameterValue.IllegalChannelId"
//  INVALIDPARAMETERVALUE_ILLEGALDESCRIBE = "InvalidParameterValue.IllegalDescribe"
//  INVALIDPARAMETERVALUE_ILLEGALDEVICEID = "InvalidParameterValue.IllegalDeviceId"
//  INVALIDPARAMETERVALUE_ILLEGALNAME = "InvalidParameterValue.IllegalName"
//  INVALIDPARAMETERVALUE_ILLEGALORGANIZATIONID = "InvalidParameterValue.IllegalOrganizationId"
//  INVALIDPARAMETERVALUE_ILLEGALSTREAMTYPE = "InvalidParameterValue.IllegalStreamType"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELIDORDEVICEID = "InvalidParameterValue.InvalidChannelIdOrDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDREPAIRMODE = "InvalidParameterValue.InvalidRepairMode"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  INVALIDPARAMETERVALUE_LIMITREPAIRTRANSITION = "InvalidParameterValue.LimitRepairTransition"
//  INVALIDPARAMETERVALUE_ORGANIZATIONCOUNTEXCEEDSRANGE = "InvalidParameterValue.OrganizationCountExceedsRange"
//  INVALIDPARAMETERVALUE_PLANCHANNELSEXCEEDSRANGE = "InvalidParameterValue.PlanChannelsExceedsRange"
//  INVALIDPARAMETERVALUE_TOOLONGDESCRIBE = "InvalidParameterValue.TooLongDescribe"
//  INVALIDPARAMETERVALUE_TOOLONGNAME = "InvalidParameterValue.TooLongName"
//  INVALIDPARAMETERVALUE_TOOLONGSTREAMTYPE = "InvalidParameterValue.TooLongStreamType"
//  RESOURCEINUSE_PLANNAMEREPEAT = "ResourceInUse.PlanNameRepeat"
//  RESOURCENOTFOUND_ORGANIZATIONIDNOTEXIST = "ResourceNotFound.OrganizationIdNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) AddRecordPlanWithContext(ctx context.Context, request *AddRecordPlanRequest) (response *AddRecordPlanResponse, err error) {
    if request == nil {
        request = NewAddRecordPlanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "AddRecordPlan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddRecordPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddRecordPlanResponse()
    err = c.Send(request, response)
    return
}

func NewAddRecordRetrieveTaskRequest() (request *AddRecordRetrieveTaskRequest) {
    request = &AddRecordRetrieveTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "AddRecordRetrieveTask")
    
    
    return
}

func NewAddRecordRetrieveTaskResponse() (response *AddRecordRetrieveTaskResponse) {
    response = &AddRecordRetrieveTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddRecordRetrieveTask
// This API is used to create a retrieval task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETER_INVALIDCHANNELS = "InvalidParameter.InvalidChannels"
//  INVALIDPARAMETER_INVALIDLIFERULEPARAM = "InvalidParameter.InvalidLifeRuleParam"
//  INVALIDPARAMETERVALUE_EMPTYCHANNELID = "InvalidParameterValue.EmptyChannelId"
//  INVALIDPARAMETERVALUE_EMPTYDEVICEID = "InvalidParameterValue.EmptyDeviceId"
//  INVALIDPARAMETERVALUE_EMPTYNAME = "InvalidParameterValue.EmptyName"
//  INVALIDPARAMETERVALUE_ILLEGALDESCRIBE = "InvalidParameterValue.IllegalDescribe"
//  INVALIDPARAMETERVALUE_ILLEGALNAME = "InvalidParameterValue.IllegalName"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELIDORDEVICEID = "InvalidParameterValue.InvalidChannelIdOrDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDEXPIRATION = "InvalidParameterValue.InvalidExpiration"
//  INVALIDPARAMETERVALUE_INVALIDRETRIEVALMODE = "InvalidParameterValue.InvalidRetrievalMode"
//  INVALIDPARAMETERVALUE_INVALIDSTARTTIMEORENDTIME = "InvalidParameterValue.InvalidStartTimeOrEndTime"
//  INVALIDPARAMETERVALUE_NONSAMEDAY = "InvalidParameterValue.NonSameDay"
//  INVALIDPARAMETERVALUE_RETRIEVETASKCHANNELSEXCEEDSRANGE = "InvalidParameterValue.RetrieveTaskChannelsExceedsRange"
//  INVALIDPARAMETERVALUE_STARTTIMEGREATERTHANOREQUALENDTIME = "InvalidParameterValue.StartTimeGreaterThanOrEqualEndTime"
//  INVALIDPARAMETERVALUE_TOOLONGDESCRIBE = "InvalidParameterValue.TooLongDescribe"
//  INVALIDPARAMETERVALUE_TOOLONGNAME = "InvalidParameterValue.TooLongName"
//  MISSINGPARAMETER_MISSINGRETRIEVETASKPARAM = "MissingParameter.MissingRetrieveTaskParam"
//  RESOURCEINUSE_RETRIEVETASKNAMEREPEAT = "ResourceInUse.RetrieveTaskNameRepeat"
func (c *Client) AddRecordRetrieveTask(request *AddRecordRetrieveTaskRequest) (response *AddRecordRetrieveTaskResponse, err error) {
    return c.AddRecordRetrieveTaskWithContext(context.Background(), request)
}

// AddRecordRetrieveTask
// This API is used to create a retrieval task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETER_INVALIDCHANNELS = "InvalidParameter.InvalidChannels"
//  INVALIDPARAMETER_INVALIDLIFERULEPARAM = "InvalidParameter.InvalidLifeRuleParam"
//  INVALIDPARAMETERVALUE_EMPTYCHANNELID = "InvalidParameterValue.EmptyChannelId"
//  INVALIDPARAMETERVALUE_EMPTYDEVICEID = "InvalidParameterValue.EmptyDeviceId"
//  INVALIDPARAMETERVALUE_EMPTYNAME = "InvalidParameterValue.EmptyName"
//  INVALIDPARAMETERVALUE_ILLEGALDESCRIBE = "InvalidParameterValue.IllegalDescribe"
//  INVALIDPARAMETERVALUE_ILLEGALNAME = "InvalidParameterValue.IllegalName"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELIDORDEVICEID = "InvalidParameterValue.InvalidChannelIdOrDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDEXPIRATION = "InvalidParameterValue.InvalidExpiration"
//  INVALIDPARAMETERVALUE_INVALIDRETRIEVALMODE = "InvalidParameterValue.InvalidRetrievalMode"
//  INVALIDPARAMETERVALUE_INVALIDSTARTTIMEORENDTIME = "InvalidParameterValue.InvalidStartTimeOrEndTime"
//  INVALIDPARAMETERVALUE_NONSAMEDAY = "InvalidParameterValue.NonSameDay"
//  INVALIDPARAMETERVALUE_RETRIEVETASKCHANNELSEXCEEDSRANGE = "InvalidParameterValue.RetrieveTaskChannelsExceedsRange"
//  INVALIDPARAMETERVALUE_STARTTIMEGREATERTHANOREQUALENDTIME = "InvalidParameterValue.StartTimeGreaterThanOrEqualEndTime"
//  INVALIDPARAMETERVALUE_TOOLONGDESCRIBE = "InvalidParameterValue.TooLongDescribe"
//  INVALIDPARAMETERVALUE_TOOLONGNAME = "InvalidParameterValue.TooLongName"
//  MISSINGPARAMETER_MISSINGRETRIEVETASKPARAM = "MissingParameter.MissingRetrieveTaskParam"
//  RESOURCEINUSE_RETRIEVETASKNAMEREPEAT = "ResourceInUse.RetrieveTaskNameRepeat"
func (c *Client) AddRecordRetrieveTaskWithContext(ctx context.Context, request *AddRecordRetrieveTaskRequest) (response *AddRecordRetrieveTaskResponse, err error) {
    if request == nil {
        request = NewAddRecordRetrieveTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "AddRecordRetrieveTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddRecordRetrieveTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddRecordRetrieveTaskResponse()
    err = c.Send(request, response)
    return
}

func NewAddRecordTemplateRequest() (request *AddRecordTemplateRequest) {
    request = &AddRecordTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "AddRecordTemplate")
    
    
    return
}

func NewAddRecordTemplateResponse() (response *AddRecordTemplateResponse) {
    response = &AddRecordTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddRecordTemplate
// Used to add a real-time cloud template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETER_INVALIDTIMESECTION = "InvalidParameter.InvalidTimeSection"
//  INVALIDPARAMETERVALUE_EMPTYNAME = "InvalidParameterValue.EmptyName"
//  INVALIDPARAMETERVALUE_ILLEGALNAME = "InvalidParameterValue.IllegalName"
//  INVALIDPARAMETERVALUE_INVALIDTIMESECTIONVALUE = "InvalidParameterValue.InvalidTimeSectionValue"
//  INVALIDPARAMETERVALUE_STARTTIMEGREATERTHANOREQUALENDTIME = "InvalidParameterValue.StartTimeGreaterThanOrEqualEndTime"
//  INVALIDPARAMETERVALUE_TIMELESSTHANTENMINUTES = "InvalidParameterValue.TimeLessThanTenMinutes"
//  INVALIDPARAMETERVALUE_TOOLONGNAME = "InvalidParameterValue.TooLongName"
//  MISSINGPARAMETER_EMPTYTIMESECTION = "MissingParameter.EmptyTimeSection"
//  RESOURCEINUSE_TEMPLATENAMEREPEAT = "ResourceInUse.TemplateNameRepeat"
func (c *Client) AddRecordTemplate(request *AddRecordTemplateRequest) (response *AddRecordTemplateResponse, err error) {
    return c.AddRecordTemplateWithContext(context.Background(), request)
}

// AddRecordTemplate
// Used to add a real-time cloud template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETER_INVALIDTIMESECTION = "InvalidParameter.InvalidTimeSection"
//  INVALIDPARAMETERVALUE_EMPTYNAME = "InvalidParameterValue.EmptyName"
//  INVALIDPARAMETERVALUE_ILLEGALNAME = "InvalidParameterValue.IllegalName"
//  INVALIDPARAMETERVALUE_INVALIDTIMESECTIONVALUE = "InvalidParameterValue.InvalidTimeSectionValue"
//  INVALIDPARAMETERVALUE_STARTTIMEGREATERTHANOREQUALENDTIME = "InvalidParameterValue.StartTimeGreaterThanOrEqualEndTime"
//  INVALIDPARAMETERVALUE_TIMELESSTHANTENMINUTES = "InvalidParameterValue.TimeLessThanTenMinutes"
//  INVALIDPARAMETERVALUE_TOOLONGNAME = "InvalidParameterValue.TooLongName"
//  MISSINGPARAMETER_EMPTYTIMESECTION = "MissingParameter.EmptyTimeSection"
//  RESOURCEINUSE_TEMPLATENAMEREPEAT = "ResourceInUse.TemplateNameRepeat"
func (c *Client) AddRecordTemplateWithContext(ctx context.Context, request *AddRecordTemplateRequest) (response *AddRecordTemplateResponse, err error) {
    if request == nil {
        request = NewAddRecordTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "AddRecordTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddRecordTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddRecordTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewAddStreamAuthRequest() (request *AddStreamAuthRequest) {
    request = &AddStreamAuthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "AddStreamAuth")
    
    
    return
}

func NewAddStreamAuthResponse() (response *AddStreamAuthResponse) {
    response = &AddStreamAuthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddStreamAuth
// Used to set push/pull stream authentication configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPULLSTATE = "InvalidParameterValue.InvalidPullState"
//  INVALIDPARAMETERVALUE_INVALIDPUSHSTATE = "InvalidParameterValue.InvalidPushState"
//  INVALIDPARAMETERVALUE_INVALIDSECRET = "InvalidParameterValue.InvalidSecret"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddStreamAuth(request *AddStreamAuthRequest) (response *AddStreamAuthResponse, err error) {
    return c.AddStreamAuthWithContext(context.Background(), request)
}

// AddStreamAuth
// Used to set push/pull stream authentication configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPULLSTATE = "InvalidParameterValue.InvalidPullState"
//  INVALIDPARAMETERVALUE_INVALIDPUSHSTATE = "InvalidParameterValue.InvalidPushState"
//  INVALIDPARAMETERVALUE_INVALIDSECRET = "InvalidParameterValue.InvalidSecret"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddStreamAuthWithContext(ctx context.Context, request *AddStreamAuthRequest) (response *AddStreamAuthResponse, err error) {
    if request == nil {
        request = NewAddStreamAuthRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "AddStreamAuth")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddStreamAuth require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddStreamAuthResponse()
    err = c.Send(request, response)
    return
}

func NewAddUserDeviceRequest() (request *AddUserDeviceRequest) {
    request = &AddUserDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "AddUserDevice")
    
    
    return
}

func NewAddUserDeviceResponse() (response *AddUserDeviceResponse) {
    response = &AddUserDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddUserDevice
// This API is used to add a single device. After the device is added, configure it based on the returned results, then wait for device registration or stream pushing.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDCLUSTERID = "InvalidParameterValue.InvalidClusterId"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEPASSWORDLENGTH = "InvalidParameterValue.InvalidDevicePasswordLength"
//  INVALIDPARAMETERVALUE_INVALIDGATEWAYPROTOCOLTYPE = "InvalidParameterValue.InvalidGatewayProtocolType"
//  INVALIDPARAMETERVALUE_INVALIDIPV4 = "InvalidParameterValue.InvalidIpv4"
//  INVALIDPARAMETERVALUE_INVALIDNAME = "InvalidParameterValue.InvalidName"
//  INVALIDPARAMETERVALUE_INVALIDRTMPAPPNAME = "InvalidParameterValue.InvalidRTMPAppName"
//  INVALIDPARAMETERVALUE_INVALIDRTMPSTREAMNAME = "InvalidParameterValue.InvalidRTMPStreamName"
//  INVALIDPARAMETERVALUE_INVALIDUSERNAME = "InvalidParameterValue.InvalidUsername"
//  INVALIDPARAMETERVALUE_RTMPPUSHSTREAMPARAMREPEAT = "InvalidParameterValue.RTMPPushStreamParamRepeat"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  REGIONERROR_RESOURCEUNREACHABLE = "RegionError.ResourceUnreachable"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOTFOUNDCLUSTER = "ResourceNotFound.NotFoundCluster"
//  UNSUPPORTEDOPERATION_PUSHDOMAINNOTEXIST = "UnsupportedOperation.PushDomainNotExist"
func (c *Client) AddUserDevice(request *AddUserDeviceRequest) (response *AddUserDeviceResponse, err error) {
    return c.AddUserDeviceWithContext(context.Background(), request)
}

// AddUserDevice
// This API is used to add a single device. After the device is added, configure it based on the returned results, then wait for device registration or stream pushing.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDCLUSTERID = "InvalidParameterValue.InvalidClusterId"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEPASSWORDLENGTH = "InvalidParameterValue.InvalidDevicePasswordLength"
//  INVALIDPARAMETERVALUE_INVALIDGATEWAYPROTOCOLTYPE = "InvalidParameterValue.InvalidGatewayProtocolType"
//  INVALIDPARAMETERVALUE_INVALIDIPV4 = "InvalidParameterValue.InvalidIpv4"
//  INVALIDPARAMETERVALUE_INVALIDNAME = "InvalidParameterValue.InvalidName"
//  INVALIDPARAMETERVALUE_INVALIDRTMPAPPNAME = "InvalidParameterValue.InvalidRTMPAppName"
//  INVALIDPARAMETERVALUE_INVALIDRTMPSTREAMNAME = "InvalidParameterValue.InvalidRTMPStreamName"
//  INVALIDPARAMETERVALUE_INVALIDUSERNAME = "InvalidParameterValue.InvalidUsername"
//  INVALIDPARAMETERVALUE_RTMPPUSHSTREAMPARAMREPEAT = "InvalidParameterValue.RTMPPushStreamParamRepeat"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  REGIONERROR_RESOURCEUNREACHABLE = "RegionError.ResourceUnreachable"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOTFOUNDCLUSTER = "ResourceNotFound.NotFoundCluster"
//  UNSUPPORTEDOPERATION_PUSHDOMAINNOTEXIST = "UnsupportedOperation.PushDomainNotExist"
func (c *Client) AddUserDeviceWithContext(ctx context.Context, request *AddUserDeviceRequest) (response *AddUserDeviceResponse, err error) {
    if request == nil {
        request = NewAddUserDeviceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "AddUserDevice")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddUserDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddUserDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewBatchAddVideoSnapshotConfigChannelRequest() (request *BatchAddVideoSnapshotConfigChannelRequest) {
    request = &BatchAddVideoSnapshotConfigChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "BatchAddVideoSnapshotConfigChannel")
    
    
    return
}

func NewBatchAddVideoSnapshotConfigChannelResponse() (response *BatchAddVideoSnapshotConfigChannelResponse) {
    response = &BatchAddVideoSnapshotConfigChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchAddVideoSnapshotConfigChannel
// Add channels for screenshot configuration in batches.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_CONTAININVALIDCHANNELID = "InvalidParameterValue.ContainInvalidChannelId"
//  INVALIDPARAMETERVALUE_CONTAININVALIDSNAPSHOTTIME = "InvalidParameterValue.ContainInvalidSnapshotTime"
//  INVALIDPARAMETERVALUE_INVALIDCONFIGID = "InvalidParameterValue.InvalidConfigId"
//  LIMITEXCEEDED_BATCHADDCHANNELNUMEXCEEDED = "LimitExceeded.BatchAddChannelNumExceeded"
//  LIMITEXCEEDED_CHANNELNUMEXCEEDED = "LimitExceeded.ChannelNumExceeded"
//  OPERATIONDENIED_CONFIGCHANGING = "OperationDenied.ConfigChanging"
//  OPERATIONDENIED_INPUTDUPLICATECHANNELID = "OperationDenied.InputDuplicateChannelId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ACCESSUNAUTHORIZED = "ResourceNotFound.AccessUnauthorized"
//  RESOURCENOTFOUND_SNAPSHOTCONFIGNOTEXIST = "ResourceNotFound.SnapshotConfigNotExist"
func (c *Client) BatchAddVideoSnapshotConfigChannel(request *BatchAddVideoSnapshotConfigChannelRequest) (response *BatchAddVideoSnapshotConfigChannelResponse, err error) {
    return c.BatchAddVideoSnapshotConfigChannelWithContext(context.Background(), request)
}

// BatchAddVideoSnapshotConfigChannel
// Add channels for screenshot configuration in batches.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_CONTAININVALIDCHANNELID = "InvalidParameterValue.ContainInvalidChannelId"
//  INVALIDPARAMETERVALUE_CONTAININVALIDSNAPSHOTTIME = "InvalidParameterValue.ContainInvalidSnapshotTime"
//  INVALIDPARAMETERVALUE_INVALIDCONFIGID = "InvalidParameterValue.InvalidConfigId"
//  LIMITEXCEEDED_BATCHADDCHANNELNUMEXCEEDED = "LimitExceeded.BatchAddChannelNumExceeded"
//  LIMITEXCEEDED_CHANNELNUMEXCEEDED = "LimitExceeded.ChannelNumExceeded"
//  OPERATIONDENIED_CONFIGCHANGING = "OperationDenied.ConfigChanging"
//  OPERATIONDENIED_INPUTDUPLICATECHANNELID = "OperationDenied.InputDuplicateChannelId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ACCESSUNAUTHORIZED = "ResourceNotFound.AccessUnauthorized"
//  RESOURCENOTFOUND_SNAPSHOTCONFIGNOTEXIST = "ResourceNotFound.SnapshotConfigNotExist"
func (c *Client) BatchAddVideoSnapshotConfigChannelWithContext(ctx context.Context, request *BatchAddVideoSnapshotConfigChannelRequest) (response *BatchAddVideoSnapshotConfigChannelResponse, err error) {
    if request == nil {
        request = NewBatchAddVideoSnapshotConfigChannelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "BatchAddVideoSnapshotConfigChannel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchAddVideoSnapshotConfigChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchAddVideoSnapshotConfigChannelResponse()
    err = c.Send(request, response)
    return
}

func NewBatchAssignVLPlanChannelsRequest() (request *BatchAssignVLPlanChannelsRequest) {
    request = &BatchAssignVLPlanChannelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "BatchAssignVLPlanChannels")
    
    
    return
}

func NewBatchAssignVLPlanChannelsResponse() (response *BatchAssignVLPlanChannelsResponse) {
    response = &BatchAssignVLPlanChannelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchAssignVLPlanChannels
// Batch associate visual understanding plan channels
//
// error code that may be returned:
//  AUTHFAILURE_CONTAINUNAUTHORIZEDCHANNELID = "AuthFailure.ContainUnauthorizedChannelId"
//  AUTHFAILURE_ORGANIZATIONUNAUTHORIZED = "AuthFailure.OrganizationUnauthorized"
//  FAILEDOPERATION_CHANNELALREADYASSIGNED = "FailedOperation.ChannelAlreadyAssigned"
//  FAILEDOPERATION_CHANNELNOTASSIGNED = "FailedOperation.ChannelNotAssigned"
//  INVALIDPARAMETERVALUE_CONTAININVALIDCHANNELID = "InvalidParameterValue.ContainInvalidChannelId"
//  INVALIDPARAMETERVALUE_CONTAINNONEXISTENTCHANNELID = "InvalidParameterValue.ContainNonexistentChannelId"
//  INVALIDPARAMETERVALUE_EMPTYCHANNELIDS = "InvalidParameterValue.EmptyChannelIds"
//  INVALIDPARAMETERVALUE_IMAGEVLPLANIDNOTEXISTED = "InvalidParameterValue.ImageVLPlanIDNotExisted"
//  INVALIDPARAMETERVALUE_INVALIDCMD = "InvalidParameterValue.InvalidCmd"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEVLPLANID = "InvalidParameterValue.InvalidImageVLPlanID"
//  INVALIDPARAMETERVALUE_ORGANIZATIONIDNOTEXISTED = "InvalidParameterValue.OrganizationIDNotExisted"
//  INVALIDPARAMETERVALUE_ORGANIZATIONIDSNOTALLOWEDFORUNASSIGN = "InvalidParameterValue.OrganizationIdsNotAllowedForUnAssign"
//  INVALIDPARAMETERVALUE_TOOMANYCHANNELIDS = "InvalidParameterValue.TooManyChannelIds"
func (c *Client) BatchAssignVLPlanChannels(request *BatchAssignVLPlanChannelsRequest) (response *BatchAssignVLPlanChannelsResponse, err error) {
    return c.BatchAssignVLPlanChannelsWithContext(context.Background(), request)
}

// BatchAssignVLPlanChannels
// Batch associate visual understanding plan channels
//
// error code that may be returned:
//  AUTHFAILURE_CONTAINUNAUTHORIZEDCHANNELID = "AuthFailure.ContainUnauthorizedChannelId"
//  AUTHFAILURE_ORGANIZATIONUNAUTHORIZED = "AuthFailure.OrganizationUnauthorized"
//  FAILEDOPERATION_CHANNELALREADYASSIGNED = "FailedOperation.ChannelAlreadyAssigned"
//  FAILEDOPERATION_CHANNELNOTASSIGNED = "FailedOperation.ChannelNotAssigned"
//  INVALIDPARAMETERVALUE_CONTAININVALIDCHANNELID = "InvalidParameterValue.ContainInvalidChannelId"
//  INVALIDPARAMETERVALUE_CONTAINNONEXISTENTCHANNELID = "InvalidParameterValue.ContainNonexistentChannelId"
//  INVALIDPARAMETERVALUE_EMPTYCHANNELIDS = "InvalidParameterValue.EmptyChannelIds"
//  INVALIDPARAMETERVALUE_IMAGEVLPLANIDNOTEXISTED = "InvalidParameterValue.ImageVLPlanIDNotExisted"
//  INVALIDPARAMETERVALUE_INVALIDCMD = "InvalidParameterValue.InvalidCmd"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEVLPLANID = "InvalidParameterValue.InvalidImageVLPlanID"
//  INVALIDPARAMETERVALUE_ORGANIZATIONIDNOTEXISTED = "InvalidParameterValue.OrganizationIDNotExisted"
//  INVALIDPARAMETERVALUE_ORGANIZATIONIDSNOTALLOWEDFORUNASSIGN = "InvalidParameterValue.OrganizationIdsNotAllowedForUnAssign"
//  INVALIDPARAMETERVALUE_TOOMANYCHANNELIDS = "InvalidParameterValue.TooManyChannelIds"
func (c *Client) BatchAssignVLPlanChannelsWithContext(ctx context.Context, request *BatchAssignVLPlanChannelsRequest) (response *BatchAssignVLPlanChannelsResponse, err error) {
    if request == nil {
        request = NewBatchAssignVLPlanChannelsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "BatchAssignVLPlanChannels")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchAssignVLPlanChannels require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchAssignVLPlanChannelsResponse()
    err = c.Send(request, response)
    return
}

func NewBatchDeleteVideoDownloadTaskRequest() (request *BatchDeleteVideoDownloadTaskRequest) {
    request = &BatchDeleteVideoDownloadTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "BatchDeleteVideoDownloadTask")
    
    
    return
}

func NewBatchDeleteVideoDownloadTaskResponse() (response *BatchDeleteVideoDownloadTaskResponse) {
    response = &BatchDeleteVideoDownloadTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchDeleteVideoDownloadTask
// This API is used to delete failed local recording download tasks in batches.
//
// error code that may be returned:
//  AUTHFAILURE_CONTAINUNAUTHORIZEDCHANNELID = "AuthFailure.ContainUnauthorizedChannelId"
//  AUTHFAILURE_ORGANIZATIONUNAUTHORIZED = "AuthFailure.OrganizationUnauthorized"
//  FAILEDOPERATION_CHANNELALREADYASSIGNED = "FailedOperation.ChannelAlreadyAssigned"
//  FAILEDOPERATION_CHANNELNOTASSIGNED = "FailedOperation.ChannelNotAssigned"
//  INVALIDPARAMETERVALUE_CONTAININVALIDCHANNELID = "InvalidParameterValue.ContainInvalidChannelId"
//  INVALIDPARAMETERVALUE_CONTAINNONEXISTENTCHANNELID = "InvalidParameterValue.ContainNonexistentChannelId"
//  INVALIDPARAMETERVALUE_EMPTYCHANNELIDS = "InvalidParameterValue.EmptyChannelIds"
//  INVALIDPARAMETERVALUE_IMAGEVLPLANIDNOTEXISTED = "InvalidParameterValue.ImageVLPlanIDNotExisted"
//  INVALIDPARAMETERVALUE_INVALIDCMD = "InvalidParameterValue.InvalidCmd"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEVLPLANID = "InvalidParameterValue.InvalidImageVLPlanID"
//  INVALIDPARAMETERVALUE_ORGANIZATIONIDNOTEXISTED = "InvalidParameterValue.OrganizationIDNotExisted"
//  INVALIDPARAMETERVALUE_ORGANIZATIONIDSNOTALLOWEDFORUNASSIGN = "InvalidParameterValue.OrganizationIdsNotAllowedForUnAssign"
//  INVALIDPARAMETERVALUE_TOOMANYCHANNELIDS = "InvalidParameterValue.TooManyChannelIds"
func (c *Client) BatchDeleteVideoDownloadTask(request *BatchDeleteVideoDownloadTaskRequest) (response *BatchDeleteVideoDownloadTaskResponse, err error) {
    return c.BatchDeleteVideoDownloadTaskWithContext(context.Background(), request)
}

// BatchDeleteVideoDownloadTask
// This API is used to delete failed local recording download tasks in batches.
//
// error code that may be returned:
//  AUTHFAILURE_CONTAINUNAUTHORIZEDCHANNELID = "AuthFailure.ContainUnauthorizedChannelId"
//  AUTHFAILURE_ORGANIZATIONUNAUTHORIZED = "AuthFailure.OrganizationUnauthorized"
//  FAILEDOPERATION_CHANNELALREADYASSIGNED = "FailedOperation.ChannelAlreadyAssigned"
//  FAILEDOPERATION_CHANNELNOTASSIGNED = "FailedOperation.ChannelNotAssigned"
//  INVALIDPARAMETERVALUE_CONTAININVALIDCHANNELID = "InvalidParameterValue.ContainInvalidChannelId"
//  INVALIDPARAMETERVALUE_CONTAINNONEXISTENTCHANNELID = "InvalidParameterValue.ContainNonexistentChannelId"
//  INVALIDPARAMETERVALUE_EMPTYCHANNELIDS = "InvalidParameterValue.EmptyChannelIds"
//  INVALIDPARAMETERVALUE_IMAGEVLPLANIDNOTEXISTED = "InvalidParameterValue.ImageVLPlanIDNotExisted"
//  INVALIDPARAMETERVALUE_INVALIDCMD = "InvalidParameterValue.InvalidCmd"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEVLPLANID = "InvalidParameterValue.InvalidImageVLPlanID"
//  INVALIDPARAMETERVALUE_ORGANIZATIONIDNOTEXISTED = "InvalidParameterValue.OrganizationIDNotExisted"
//  INVALIDPARAMETERVALUE_ORGANIZATIONIDSNOTALLOWEDFORUNASSIGN = "InvalidParameterValue.OrganizationIdsNotAllowedForUnAssign"
//  INVALIDPARAMETERVALUE_TOOMANYCHANNELIDS = "InvalidParameterValue.TooManyChannelIds"
func (c *Client) BatchDeleteVideoDownloadTaskWithContext(ctx context.Context, request *BatchDeleteVideoDownloadTaskRequest) (response *BatchDeleteVideoDownloadTaskResponse, err error) {
    if request == nil {
        request = NewBatchDeleteVideoDownloadTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "BatchDeleteVideoDownloadTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchDeleteVideoDownloadTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchDeleteVideoDownloadTaskResponse()
    err = c.Send(request, response)
    return
}

func NewBatchGetVLPlanSourceFileURLRequest() (request *BatchGetVLPlanSourceFileURLRequest) {
    request = &BatchGetVLPlanSourceFileURLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "BatchGetVLPlanSourceFileURL")
    
    
    return
}

func NewBatchGetVLPlanSourceFileURLResponse() (response *BatchGetVLPlanSourceFileURLResponse) {
    response = &BatchGetVLPlanSourceFileURLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchGetVLPlanSourceFileURL
// This API is used to batch obtain source material URLs for visual understanding plans.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDEXPIRE = "InvalidParameterValue.InvalidExpire"
//  INVALIDPARAMETERVALUE_INVALIDVLPLANRESULTIDS = "InvalidParameterValue.InvalidVLPlanResultIds"
//  INVALIDPARAMETERVALUE_VLPLANRESULTIDSEXCEEDLIMIT = "InvalidParameterValue.VLPlanResultIdsExceedLimit"
func (c *Client) BatchGetVLPlanSourceFileURL(request *BatchGetVLPlanSourceFileURLRequest) (response *BatchGetVLPlanSourceFileURLResponse, err error) {
    return c.BatchGetVLPlanSourceFileURLWithContext(context.Background(), request)
}

// BatchGetVLPlanSourceFileURL
// This API is used to batch obtain source material URLs for visual understanding plans.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDEXPIRE = "InvalidParameterValue.InvalidExpire"
//  INVALIDPARAMETERVALUE_INVALIDVLPLANRESULTIDS = "InvalidParameterValue.InvalidVLPlanResultIds"
//  INVALIDPARAMETERVALUE_VLPLANRESULTIDSEXCEEDLIMIT = "InvalidParameterValue.VLPlanResultIdsExceedLimit"
func (c *Client) BatchGetVLPlanSourceFileURLWithContext(ctx context.Context, request *BatchGetVLPlanSourceFileURLRequest) (response *BatchGetVLPlanSourceFileURLResponse, err error) {
    if request == nil {
        request = NewBatchGetVLPlanSourceFileURLRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "BatchGetVLPlanSourceFileURL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchGetVLPlanSourceFileURL require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchGetVLPlanSourceFileURLResponse()
    err = c.Send(request, response)
    return
}

func NewBatchGetVideoSnapshotImageUrlRequest() (request *BatchGetVideoSnapshotImageUrlRequest) {
    request = &BatchGetVideoSnapshotImageUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "BatchGetVideoSnapshotImageUrl")
    
    
    return
}

func NewBatchGetVideoSnapshotImageUrlResponse() (response *BatchGetVideoSnapshotImageUrlResponse) {
    response = &BatchGetVideoSnapshotImageUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchGetVideoSnapshotImageUrl
// Bulk gets screenshot image URLs
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_CONTAININVALIDIMAGEID = "InvalidParameterValue.ContainInvalidImageId"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDIMAGEURLEXPIRETIME = "InvalidParameterValue.UnsupportedImageUrlExpireTime"
//  LIMITEXCEEDED_IMAGEIDSLICEEXCEEDLIMIT = "LimitExceeded.ImageIdSliceExceedLimit"
//  UNSUPPORTEDOPERATION_IMAGEDOMAINNOTEXIST = "UnsupportedOperation.ImageDomainNotExist"
func (c *Client) BatchGetVideoSnapshotImageUrl(request *BatchGetVideoSnapshotImageUrlRequest) (response *BatchGetVideoSnapshotImageUrlResponse, err error) {
    return c.BatchGetVideoSnapshotImageUrlWithContext(context.Background(), request)
}

// BatchGetVideoSnapshotImageUrl
// Bulk gets screenshot image URLs
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_CONTAININVALIDIMAGEID = "InvalidParameterValue.ContainInvalidImageId"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDIMAGEURLEXPIRETIME = "InvalidParameterValue.UnsupportedImageUrlExpireTime"
//  LIMITEXCEEDED_IMAGEIDSLICEEXCEEDLIMIT = "LimitExceeded.ImageIdSliceExceedLimit"
//  UNSUPPORTEDOPERATION_IMAGEDOMAINNOTEXIST = "UnsupportedOperation.ImageDomainNotExist"
func (c *Client) BatchGetVideoSnapshotImageUrlWithContext(ctx context.Context, request *BatchGetVideoSnapshotImageUrlRequest) (response *BatchGetVideoSnapshotImageUrlResponse, err error) {
    if request == nil {
        request = NewBatchGetVideoSnapshotImageUrlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "BatchGetVideoSnapshotImageUrl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchGetVideoSnapshotImageUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchGetVideoSnapshotImageUrlResponse()
    err = c.Send(request, response)
    return
}

func NewBatchOperateDeviceRequest() (request *BatchOperateDeviceRequest) {
    request = &BatchOperateDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "BatchOperateDevice")
    
    
    return
}

func NewBatchOperateDeviceResponse() (response *BatchOperateDeviceResponse) {
    response = &BatchOperateDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchOperateDevice
// Used for batch operations (enable, disable, delete) on devices.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_UNSUPPORTOPERATECMD = "InvalidParameterValue.UnSupportOperateCMD"
//  REGIONERROR_RESOURCEUNREACHABLE = "RegionError.ResourceUnreachable"
func (c *Client) BatchOperateDevice(request *BatchOperateDeviceRequest) (response *BatchOperateDeviceResponse, err error) {
    return c.BatchOperateDeviceWithContext(context.Background(), request)
}

// BatchOperateDevice
// Used for batch operations (enable, disable, delete) on devices.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_UNSUPPORTOPERATECMD = "InvalidParameterValue.UnSupportOperateCMD"
//  REGIONERROR_RESOURCEUNREACHABLE = "RegionError.ResourceUnreachable"
func (c *Client) BatchOperateDeviceWithContext(ctx context.Context, request *BatchOperateDeviceRequest) (response *BatchOperateDeviceResponse, err error) {
    if request == nil {
        request = NewBatchOperateDeviceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "BatchOperateDevice")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchOperateDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchOperateDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewCallISAPIRequest() (request *CallISAPIRequest) {
    request = &CallISAPIRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "CallISAPI")
    
    
    return
}

func NewCallISAPIResponse() (response *CallISAPIResponse) {
    response = &CallISAPIResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CallISAPI
// This API can be based on the Hikvision ISUP 5.0 protocol to pass through ISAPI request data. Before calling the API, ensure the device has successfully registered to this platform using the ISUP protocol.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_UNSUPPORTOPERATECMD = "InvalidParameterValue.UnSupportOperateCMD"
//  REGIONERROR_RESOURCEUNREACHABLE = "RegionError.ResourceUnreachable"
func (c *Client) CallISAPI(request *CallISAPIRequest) (response *CallISAPIResponse, err error) {
    return c.CallISAPIWithContext(context.Background(), request)
}

// CallISAPI
// This API can be based on the Hikvision ISUP 5.0 protocol to pass through ISAPI request data. Before calling the API, ensure the device has successfully registered to this platform using the ISUP protocol.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_UNSUPPORTOPERATECMD = "InvalidParameterValue.UnSupportOperateCMD"
//  REGIONERROR_RESOURCEUNREACHABLE = "RegionError.ResourceUnreachable"
func (c *Client) CallISAPIWithContext(ctx context.Context, request *CallISAPIRequest) (response *CallISAPIResponse, err error) {
    if request == nil {
        request = NewCallISAPIRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "CallISAPI")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CallISAPI require credential")
    }

    request.SetContext(ctx)
    
    response = NewCallISAPIResponse()
    err = c.Send(request, response)
    return
}

func NewControlDevicePTZRequest() (request *ControlDevicePTZRequest) {
    request = &ControlDevicePTZRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ControlDevicePTZ")
    
    
    return
}

func NewControlDevicePTZResponse() (response *ControlDevicePTZResponse) {
    response = &ControlDevicePTZResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ControlDevicePTZ
// Used for device channel panning control, including turn, zoom scale, zoom, and aperture.
//
// error code that may be returned:
//  FAILEDOPERATION_DEVICERESPONSETIMEOUT = "FailedOperation.DeviceResponseTimeOut"
//  FAILEDOPERATION_DEVICERESULTTIMEOUT = "FailedOperation.DeviceResultTimeOut"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDPTZCMD = "InvalidParameterValue.UnSupportedPTZCMD"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDSCALEPARAM = "InvalidParameterValue.UnSupportedScaleParam"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CHANNELOFFLINE = "ResourceUnavailable.ChannelOffline"
//  RESOURCEUNAVAILABLE_DEVDISABLE = "ResourceUnavailable.DevDisable"
//  RESOURCEUNAVAILABLE_DEVNOREGISTER = "ResourceUnavailable.DevNoRegister"
//  RESOURCEUNAVAILABLE_DEVOFFLINE = "ResourceUnavailable.DevOffline"
func (c *Client) ControlDevicePTZ(request *ControlDevicePTZRequest) (response *ControlDevicePTZResponse, err error) {
    return c.ControlDevicePTZWithContext(context.Background(), request)
}

// ControlDevicePTZ
// Used for device channel panning control, including turn, zoom scale, zoom, and aperture.
//
// error code that may be returned:
//  FAILEDOPERATION_DEVICERESPONSETIMEOUT = "FailedOperation.DeviceResponseTimeOut"
//  FAILEDOPERATION_DEVICERESULTTIMEOUT = "FailedOperation.DeviceResultTimeOut"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDPTZCMD = "InvalidParameterValue.UnSupportedPTZCMD"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDSCALEPARAM = "InvalidParameterValue.UnSupportedScaleParam"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CHANNELOFFLINE = "ResourceUnavailable.ChannelOffline"
//  RESOURCEUNAVAILABLE_DEVDISABLE = "ResourceUnavailable.DevDisable"
//  RESOURCEUNAVAILABLE_DEVNOREGISTER = "ResourceUnavailable.DevNoRegister"
//  RESOURCEUNAVAILABLE_DEVOFFLINE = "ResourceUnavailable.DevOffline"
func (c *Client) ControlDevicePTZWithContext(ctx context.Context, request *ControlDevicePTZRequest) (response *ControlDevicePTZResponse, err error) {
    if request == nil {
        request = NewControlDevicePTZRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "ControlDevicePTZ")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ControlDevicePTZ require credential")
    }

    request.SetContext(ctx)
    
    response = NewControlDevicePTZResponse()
    err = c.Send(request, response)
    return
}

func NewControlDevicePresetRequest() (request *ControlDevicePresetRequest) {
    request = &ControlDevicePresetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ControlDevicePreset")
    
    
    return
}

func NewControlDevicePresetResponse() (response *ControlDevicePresetResponse) {
    response = &ControlDevicePresetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ControlDevicePreset
// Used to operate device preset positions, including settings, deletion, and calls.
//
// error code that may be returned:
//  FAILEDOPERATION_DEVICERESPONSETIMEOUT = "FailedOperation.DeviceResponseTimeOut"
//  FAILEDOPERATION_DEVICERESULTTIMEOUT = "FailedOperation.DeviceResultTimeOut"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDPRESETCMD = "InvalidParameterValue.UnSupportedPresetCMD"
//  RESOURCEUNAVAILABLE_CHANNELOFFLINE = "ResourceUnavailable.ChannelOffline"
//  RESOURCEUNAVAILABLE_DEVDISABLE = "ResourceUnavailable.DevDisable"
//  RESOURCEUNAVAILABLE_DEVNOREGISTER = "ResourceUnavailable.DevNoRegister"
//  RESOURCEUNAVAILABLE_DEVOFFLINE = "ResourceUnavailable.DevOffline"
func (c *Client) ControlDevicePreset(request *ControlDevicePresetRequest) (response *ControlDevicePresetResponse, err error) {
    return c.ControlDevicePresetWithContext(context.Background(), request)
}

// ControlDevicePreset
// Used to operate device preset positions, including settings, deletion, and calls.
//
// error code that may be returned:
//  FAILEDOPERATION_DEVICERESPONSETIMEOUT = "FailedOperation.DeviceResponseTimeOut"
//  FAILEDOPERATION_DEVICERESULTTIMEOUT = "FailedOperation.DeviceResultTimeOut"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDPRESETCMD = "InvalidParameterValue.UnSupportedPresetCMD"
//  RESOURCEUNAVAILABLE_CHANNELOFFLINE = "ResourceUnavailable.ChannelOffline"
//  RESOURCEUNAVAILABLE_DEVDISABLE = "ResourceUnavailable.DevDisable"
//  RESOURCEUNAVAILABLE_DEVNOREGISTER = "ResourceUnavailable.DevNoRegister"
//  RESOURCEUNAVAILABLE_DEVOFFLINE = "ResourceUnavailable.DevOffline"
func (c *Client) ControlDevicePresetWithContext(ctx context.Context, request *ControlDevicePresetRequest) (response *ControlDevicePresetResponse, err error) {
    if request == nil {
        request = NewControlDevicePresetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "ControlDevicePreset")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ControlDevicePreset require credential")
    }

    request.SetContext(ctx)
    
    response = NewControlDevicePresetResponse()
    err = c.Send(request, response)
    return
}

func NewControlDeviceSnapshotRequest() (request *ControlDeviceSnapshotRequest) {
    request = &ControlDeviceSnapshotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ControlDeviceSnapshot")
    
    
    return
}

func NewControlDeviceSnapshotResponse() (response *ControlDeviceSnapshotResponse) {
    response = &ControlDeviceSnapshotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ControlDeviceSnapshot
// Control device capture (single). Currently only support national standard devices.
//
// error code that may be returned:
//  FAILEDOPERATION_DEVICERESPONSETIMEOUT = "FailedOperation.DeviceResponseTimeOut"
//  FAILEDOPERATION_DEVICERESULTTIMEOUT = "FailedOperation.DeviceResultTimeOut"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDPRESETCMD = "InvalidParameterValue.UnSupportedPresetCMD"
//  RESOURCEUNAVAILABLE_CHANNELOFFLINE = "ResourceUnavailable.ChannelOffline"
//  RESOURCEUNAVAILABLE_DEVDISABLE = "ResourceUnavailable.DevDisable"
//  RESOURCEUNAVAILABLE_DEVNOREGISTER = "ResourceUnavailable.DevNoRegister"
//  RESOURCEUNAVAILABLE_DEVOFFLINE = "ResourceUnavailable.DevOffline"
func (c *Client) ControlDeviceSnapshot(request *ControlDeviceSnapshotRequest) (response *ControlDeviceSnapshotResponse, err error) {
    return c.ControlDeviceSnapshotWithContext(context.Background(), request)
}

// ControlDeviceSnapshot
// Control device capture (single). Currently only support national standard devices.
//
// error code that may be returned:
//  FAILEDOPERATION_DEVICERESPONSETIMEOUT = "FailedOperation.DeviceResponseTimeOut"
//  FAILEDOPERATION_DEVICERESULTTIMEOUT = "FailedOperation.DeviceResultTimeOut"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDPRESETCMD = "InvalidParameterValue.UnSupportedPresetCMD"
//  RESOURCEUNAVAILABLE_CHANNELOFFLINE = "ResourceUnavailable.ChannelOffline"
//  RESOURCEUNAVAILABLE_DEVDISABLE = "ResourceUnavailable.DevDisable"
//  RESOURCEUNAVAILABLE_DEVNOREGISTER = "ResourceUnavailable.DevNoRegister"
//  RESOURCEUNAVAILABLE_DEVOFFLINE = "ResourceUnavailable.DevOffline"
func (c *Client) ControlDeviceSnapshotWithContext(ctx context.Context, request *ControlDeviceSnapshotRequest) (response *ControlDeviceSnapshotResponse, err error) {
    if request == nil {
        request = NewControlDeviceSnapshotRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "ControlDeviceSnapshot")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ControlDeviceSnapshot require credential")
    }

    request.SetContext(ctx)
    
    response = NewControlDeviceSnapshotResponse()
    err = c.Send(request, response)
    return
}

func NewControlDeviceStreamRequest() (request *ControlDeviceStreamRequest) {
    request = &ControlDeviceStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ControlDeviceStream")
    
    
    return
}

func NewControlDeviceStreamResponse() (response *ControlDeviceStreamResponse) {
    response = &ControlDeviceStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ControlDeviceStream
// This API is used to obtain the real-time streaming address of a device.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DEVICEMISMATCHCHANNEL = "InvalidParameterValue.DeviceMismatchChannel"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDSTREAMTYPE = "InvalidParameterValue.InvalidStreamType"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CHANNELOFFLINE = "ResourceUnavailable.ChannelOffline"
//  RESOURCEUNAVAILABLE_DEVDISABLE = "ResourceUnavailable.DevDisable"
//  RESOURCEUNAVAILABLE_DEVNOREGISTER = "ResourceUnavailable.DevNoRegister"
//  RESOURCEUNAVAILABLE_DEVOFFLINE = "ResourceUnavailable.DevOffline"
//  UNSUPPORTEDOPERATION_STREAMTYPEORRESOLUTION = "UnsupportedOperation.StreamTypeOrResolution"
func (c *Client) ControlDeviceStream(request *ControlDeviceStreamRequest) (response *ControlDeviceStreamResponse, err error) {
    return c.ControlDeviceStreamWithContext(context.Background(), request)
}

// ControlDeviceStream
// This API is used to obtain the real-time streaming address of a device.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DEVICEMISMATCHCHANNEL = "InvalidParameterValue.DeviceMismatchChannel"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDSTREAMTYPE = "InvalidParameterValue.InvalidStreamType"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CHANNELOFFLINE = "ResourceUnavailable.ChannelOffline"
//  RESOURCEUNAVAILABLE_DEVDISABLE = "ResourceUnavailable.DevDisable"
//  RESOURCEUNAVAILABLE_DEVNOREGISTER = "ResourceUnavailable.DevNoRegister"
//  RESOURCEUNAVAILABLE_DEVOFFLINE = "ResourceUnavailable.DevOffline"
//  UNSUPPORTEDOPERATION_STREAMTYPEORRESOLUTION = "UnsupportedOperation.StreamTypeOrResolution"
func (c *Client) ControlDeviceStreamWithContext(ctx context.Context, request *ControlDeviceStreamRequest) (response *ControlDeviceStreamResponse, err error) {
    if request == nil {
        request = NewControlDeviceStreamRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "ControlDeviceStream")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ControlDeviceStream require credential")
    }

    request.SetContext(ctx)
    
    response = NewControlDeviceStreamResponse()
    err = c.Send(request, response)
    return
}

func NewControlRecordTimelineRequest() (request *ControlRecordTimelineRequest) {
    request = &ControlRecordTimelineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ControlRecordTimeline")
    
    
    return
}

func NewControlRecordTimelineResponse() (response *ControlRecordTimelineResponse) {
    response = &ControlRecordTimelineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ControlRecordTimeline
// This API is used to query the local recording timeline info of a device, which is stored locally on the NVR/IPC.
//
// error code that may be returned:
//  FAILEDOPERATION_DEVICERESPONSETIMEOUT = "FailedOperation.DeviceResponseTimeOut"
//  FAILEDOPERATION_DEVICERESULTTIMEOUT = "FailedOperation.DeviceResultTimeOut"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ENDTIMEZERO = "InvalidParameterValue.EndTimeZero"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_STARTOVERENDTIME = "InvalidParameterValue.StartOverEndTime"
//  INVALIDPARAMETERVALUE_STARTOVERNOWTIME = "InvalidParameterValue.StartOverNowTime"
//  INVALIDPARAMETERVALUE_STARTTIMEGREATERTHANOREQUALENDTIME = "InvalidParameterValue.StartTimeGreaterThanOrEqualEndTime"
//  INVALIDPARAMETERVALUE_STARTTIMEZERO = "InvalidParameterValue.StartTimeZero"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_DEVDISABLE = "ResourceUnavailable.DevDisable"
//  RESOURCEUNAVAILABLE_DEVNOREGISTER = "ResourceUnavailable.DevNoRegister"
//  RESOURCEUNAVAILABLE_DEVOFFLINE = "ResourceUnavailable.DevOffline"
func (c *Client) ControlRecordTimeline(request *ControlRecordTimelineRequest) (response *ControlRecordTimelineResponse, err error) {
    return c.ControlRecordTimelineWithContext(context.Background(), request)
}

// ControlRecordTimeline
// This API is used to query the local recording timeline info of a device, which is stored locally on the NVR/IPC.
//
// error code that may be returned:
//  FAILEDOPERATION_DEVICERESPONSETIMEOUT = "FailedOperation.DeviceResponseTimeOut"
//  FAILEDOPERATION_DEVICERESULTTIMEOUT = "FailedOperation.DeviceResultTimeOut"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ENDTIMEZERO = "InvalidParameterValue.EndTimeZero"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_STARTOVERENDTIME = "InvalidParameterValue.StartOverEndTime"
//  INVALIDPARAMETERVALUE_STARTOVERNOWTIME = "InvalidParameterValue.StartOverNowTime"
//  INVALIDPARAMETERVALUE_STARTTIMEGREATERTHANOREQUALENDTIME = "InvalidParameterValue.StartTimeGreaterThanOrEqualEndTime"
//  INVALIDPARAMETERVALUE_STARTTIMEZERO = "InvalidParameterValue.StartTimeZero"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_DEVDISABLE = "ResourceUnavailable.DevDisable"
//  RESOURCEUNAVAILABLE_DEVNOREGISTER = "ResourceUnavailable.DevNoRegister"
//  RESOURCEUNAVAILABLE_DEVOFFLINE = "ResourceUnavailable.DevOffline"
func (c *Client) ControlRecordTimelineWithContext(ctx context.Context, request *ControlRecordTimelineRequest) (response *ControlRecordTimelineResponse, err error) {
    if request == nil {
        request = NewControlRecordTimelineRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "ControlRecordTimeline")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ControlRecordTimeline require credential")
    }

    request.SetContext(ctx)
    
    response = NewControlRecordTimelineResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVLPlanRequest() (request *CreateVLPlanRequest) {
    request = &CreateVLPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "CreateVLPlan")
    
    
    return
}

func NewCreateVLPlanResponse() (response *CreateVLPlanResponse) {
    response = &CreateVLPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateVLPlan
// Creates a vision understanding plan
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_IMAGEVLPLANNAMECONTAINILLEGALCHARACTER = "InvalidParameterValue.ImageVLPlanNameContainIllegalCharacter"
//  INVALIDPARAMETERVALUE_INVALIDCALLBACKTOKEN = "InvalidParameterValue.InvalidCallbackToken"
//  INVALIDPARAMETERVALUE_INVALIDCALLBACKURL = "InvalidParameterValue.InvalidCallbackUrl"
//  INVALIDPARAMETERVALUE_INVALIDFILEEXPIRES = "InvalidParameterValue.InvalidFileExpires"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEVLPLANDESCRIBE = "InvalidParameterValue.InvalidImageVLPlanDescribe"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEVLPLANNAME = "InvalidParameterValue.InvalidImageVLPlanName"
//  INVALIDPARAMETERVALUE_TEMPLATEIDNOTEXISTED = "InvalidParameterValue.TemplateIDNotExisted"
func (c *Client) CreateVLPlan(request *CreateVLPlanRequest) (response *CreateVLPlanResponse, err error) {
    return c.CreateVLPlanWithContext(context.Background(), request)
}

// CreateVLPlan
// Creates a vision understanding plan
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_IMAGEVLPLANNAMECONTAINILLEGALCHARACTER = "InvalidParameterValue.ImageVLPlanNameContainIllegalCharacter"
//  INVALIDPARAMETERVALUE_INVALIDCALLBACKTOKEN = "InvalidParameterValue.InvalidCallbackToken"
//  INVALIDPARAMETERVALUE_INVALIDCALLBACKURL = "InvalidParameterValue.InvalidCallbackUrl"
//  INVALIDPARAMETERVALUE_INVALIDFILEEXPIRES = "InvalidParameterValue.InvalidFileExpires"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEVLPLANDESCRIBE = "InvalidParameterValue.InvalidImageVLPlanDescribe"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEVLPLANNAME = "InvalidParameterValue.InvalidImageVLPlanName"
//  INVALIDPARAMETERVALUE_TEMPLATEIDNOTEXISTED = "InvalidParameterValue.TemplateIDNotExisted"
func (c *Client) CreateVLPlanWithContext(ctx context.Context, request *CreateVLPlanRequest) (response *CreateVLPlanResponse, err error) {
    if request == nil {
        request = NewCreateVLPlanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "CreateVLPlan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVLPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVLPlanResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVLTemplateRequest() (request *CreateVLTemplateRequest) {
    request = &CreateVLTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "CreateVLTemplate")
    
    
    return
}

func NewCreateVLTemplateResponse() (response *CreateVLTemplateResponse) {
    response = &CreateVLTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateVLTemplate
// Create a visual understanding template.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDAIMODEL = "InvalidParameterValue.InvalidAIModel"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIBE = "InvalidParameterValue.InvalidDescribe"
//  INVALIDPARAMETERVALUE_INVALIDIMAGECONFIGINTERVALFORLOWFREQ = "InvalidParameterValue.InvalidImageConfigIntervalForLowFreq"
//  INVALIDPARAMETERVALUE_INVALIDIMAGECONFIGINTERVALFORREALTIME = "InvalidParameterValue.InvalidImageConfigIntervalForRealtime"
//  INVALIDPARAMETERVALUE_INVALIDIMAGECONFIGTYPE = "InvalidParameterValue.InvalidImageConfigType"
//  INVALIDPARAMETERVALUE_INVALIDMEDIATYPE = "InvalidParameterValue.InvalidMediaType"
//  INVALIDPARAMETERVALUE_INVALIDNAME = "InvalidParameterValue.InvalidName"
//  INVALIDPARAMETERVALUE_INVALIDPROMPT = "InvalidParameterValue.InvalidPrompt"
//  INVALIDPARAMETERVALUE_INVALIDTIMES = "InvalidParameterValue.InvalidTimes"
//  INVALIDPARAMETERVALUE_MISSINGIMAGECONFIG = "InvalidParameterValue.MissingImageConfig"
//  INVALIDPARAMETERVALUE_MISSINGTIMES = "InvalidParameterValue.MissingTimes"
func (c *Client) CreateVLTemplate(request *CreateVLTemplateRequest) (response *CreateVLTemplateResponse, err error) {
    return c.CreateVLTemplateWithContext(context.Background(), request)
}

// CreateVLTemplate
// Create a visual understanding template.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDAIMODEL = "InvalidParameterValue.InvalidAIModel"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIBE = "InvalidParameterValue.InvalidDescribe"
//  INVALIDPARAMETERVALUE_INVALIDIMAGECONFIGINTERVALFORLOWFREQ = "InvalidParameterValue.InvalidImageConfigIntervalForLowFreq"
//  INVALIDPARAMETERVALUE_INVALIDIMAGECONFIGINTERVALFORREALTIME = "InvalidParameterValue.InvalidImageConfigIntervalForRealtime"
//  INVALIDPARAMETERVALUE_INVALIDIMAGECONFIGTYPE = "InvalidParameterValue.InvalidImageConfigType"
//  INVALIDPARAMETERVALUE_INVALIDMEDIATYPE = "InvalidParameterValue.InvalidMediaType"
//  INVALIDPARAMETERVALUE_INVALIDNAME = "InvalidParameterValue.InvalidName"
//  INVALIDPARAMETERVALUE_INVALIDPROMPT = "InvalidParameterValue.InvalidPrompt"
//  INVALIDPARAMETERVALUE_INVALIDTIMES = "InvalidParameterValue.InvalidTimes"
//  INVALIDPARAMETERVALUE_MISSINGIMAGECONFIG = "InvalidParameterValue.MissingImageConfig"
//  INVALIDPARAMETERVALUE_MISSINGTIMES = "InvalidParameterValue.MissingTimes"
func (c *Client) CreateVLTemplateWithContext(ctx context.Context, request *CreateVLTemplateRequest) (response *CreateVLTemplateResponse, err error) {
    if request == nil {
        request = NewCreateVLTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "CreateVLTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVLTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVLTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVideoDownloadTaskRequest() (request *CreateVideoDownloadTaskRequest) {
    request = &CreateVideoDownloadTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "CreateVideoDownloadTask")
    
    
    return
}

func NewCreateVideoDownloadTaskResponse() (response *CreateVideoDownloadTaskResponse) {
    response = &CreateVideoDownloadTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateVideoDownloadTask
// This API is used to create a local recording download task.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDAIMODEL = "InvalidParameterValue.InvalidAIModel"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIBE = "InvalidParameterValue.InvalidDescribe"
//  INVALIDPARAMETERVALUE_INVALIDIMAGECONFIGINTERVALFORLOWFREQ = "InvalidParameterValue.InvalidImageConfigIntervalForLowFreq"
//  INVALIDPARAMETERVALUE_INVALIDIMAGECONFIGINTERVALFORREALTIME = "InvalidParameterValue.InvalidImageConfigIntervalForRealtime"
//  INVALIDPARAMETERVALUE_INVALIDIMAGECONFIGTYPE = "InvalidParameterValue.InvalidImageConfigType"
//  INVALIDPARAMETERVALUE_INVALIDMEDIATYPE = "InvalidParameterValue.InvalidMediaType"
//  INVALIDPARAMETERVALUE_INVALIDNAME = "InvalidParameterValue.InvalidName"
//  INVALIDPARAMETERVALUE_INVALIDPROMPT = "InvalidParameterValue.InvalidPrompt"
//  INVALIDPARAMETERVALUE_INVALIDTIMES = "InvalidParameterValue.InvalidTimes"
//  INVALIDPARAMETERVALUE_MISSINGIMAGECONFIG = "InvalidParameterValue.MissingImageConfig"
//  INVALIDPARAMETERVALUE_MISSINGTIMES = "InvalidParameterValue.MissingTimes"
func (c *Client) CreateVideoDownloadTask(request *CreateVideoDownloadTaskRequest) (response *CreateVideoDownloadTaskResponse, err error) {
    return c.CreateVideoDownloadTaskWithContext(context.Background(), request)
}

// CreateVideoDownloadTask
// This API is used to create a local recording download task.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDAIMODEL = "InvalidParameterValue.InvalidAIModel"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIBE = "InvalidParameterValue.InvalidDescribe"
//  INVALIDPARAMETERVALUE_INVALIDIMAGECONFIGINTERVALFORLOWFREQ = "InvalidParameterValue.InvalidImageConfigIntervalForLowFreq"
//  INVALIDPARAMETERVALUE_INVALIDIMAGECONFIGINTERVALFORREALTIME = "InvalidParameterValue.InvalidImageConfigIntervalForRealtime"
//  INVALIDPARAMETERVALUE_INVALIDIMAGECONFIGTYPE = "InvalidParameterValue.InvalidImageConfigType"
//  INVALIDPARAMETERVALUE_INVALIDMEDIATYPE = "InvalidParameterValue.InvalidMediaType"
//  INVALIDPARAMETERVALUE_INVALIDNAME = "InvalidParameterValue.InvalidName"
//  INVALIDPARAMETERVALUE_INVALIDPROMPT = "InvalidParameterValue.InvalidPrompt"
//  INVALIDPARAMETERVALUE_INVALIDTIMES = "InvalidParameterValue.InvalidTimes"
//  INVALIDPARAMETERVALUE_MISSINGIMAGECONFIG = "InvalidParameterValue.MissingImageConfig"
//  INVALIDPARAMETERVALUE_MISSINGTIMES = "InvalidParameterValue.MissingTimes"
func (c *Client) CreateVideoDownloadTaskWithContext(ctx context.Context, request *CreateVideoDownloadTaskRequest) (response *CreateVideoDownloadTaskResponse, err error) {
    if request == nil {
        request = NewCreateVideoDownloadTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "CreateVideoDownloadTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVideoDownloadTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVideoDownloadTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVideoSnapshotConfigRequest() (request *CreateVideoSnapshotConfigRequest) {
    request = &CreateVideoSnapshotConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "CreateVideoSnapshotConfig")
    
    
    return
}

func NewCreateVideoSnapshotConfigResponse() (response *CreateVideoSnapshotConfigResponse) {
    response = &CreateVideoSnapshotConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateVideoSnapshotConfig
// This API is used to create a screenshot configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_CONTAININVALIDTOKEN = "InvalidParameterValue.ContainInvalidToken"
//  INVALIDPARAMETERVALUE_INVALIDCALLBACKURL = "InvalidParameterValue.InvalidCallbackUrl"
//  INVALIDPARAMETERVALUE_INVALIDNAME = "InvalidParameterValue.InvalidName"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDFILEEXPIRETIME = "InvalidParameterValue.UnsupportedFileExpireTime"
//  LIMITEXCEEDED_SNAPSHOTCONFIGNUMEXCEEDED = "LimitExceeded.SnapshotConfigNumExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateVideoSnapshotConfig(request *CreateVideoSnapshotConfigRequest) (response *CreateVideoSnapshotConfigResponse, err error) {
    return c.CreateVideoSnapshotConfigWithContext(context.Background(), request)
}

// CreateVideoSnapshotConfig
// This API is used to create a screenshot configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_CONTAININVALIDTOKEN = "InvalidParameterValue.ContainInvalidToken"
//  INVALIDPARAMETERVALUE_INVALIDCALLBACKURL = "InvalidParameterValue.InvalidCallbackUrl"
//  INVALIDPARAMETERVALUE_INVALIDNAME = "InvalidParameterValue.InvalidName"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDFILEEXPIRETIME = "InvalidParameterValue.UnsupportedFileExpireTime"
//  LIMITEXCEEDED_SNAPSHOTCONFIGNUMEXCEEDED = "LimitExceeded.SnapshotConfigNumExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateVideoSnapshotConfigWithContext(ctx context.Context, request *CreateVideoSnapshotConfigRequest) (response *CreateVideoSnapshotConfigResponse, err error) {
    if request == nil {
        request = NewCreateVideoSnapshotConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "CreateVideoSnapshotConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVideoSnapshotConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVideoSnapshotConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVideoSnapshotTemplateRequest() (request *CreateVideoSnapshotTemplateRequest) {
    request = &CreateVideoSnapshotTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "CreateVideoSnapshotTemplate")
    
    
    return
}

func NewCreateVideoSnapshotTemplateResponse() (response *CreateVideoSnapshotTemplateResponse) {
    response = &CreateVideoSnapshotTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateVideoSnapshotTemplate
// Create a screenshot template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_CONTAININVALIDSNAPSHOTTIME = "InvalidParameterValue.ContainInvalidSnapshotTime"
//  INVALIDPARAMETERVALUE_INVALIDNAME = "InvalidParameterValue.InvalidName"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDINTERVAL = "InvalidParameterValue.UnsupportedInterval"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDSNAPSHOTTOTAL = "InvalidParameterValue.UnsupportedSnapshotTotal"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDSNAPSHOTTYPE = "InvalidParameterValue.UnsupportedSnapshotType"
//  LIMITEXCEEDED_SNAPSHOTTEMPLATENUMEXCEEDED = "LimitExceeded.SnapshotTemplateNumExceeded"
//  LIMITEXCEEDED_TIMELINESLICEEXCEEDLIMIT = "LimitExceeded.TimelineSliceExceedLimit"
//  UNSUPPORTEDOPERATION_IMAGEDOMAINNOTEXIST = "UnsupportedOperation.ImageDomainNotExist"
func (c *Client) CreateVideoSnapshotTemplate(request *CreateVideoSnapshotTemplateRequest) (response *CreateVideoSnapshotTemplateResponse, err error) {
    return c.CreateVideoSnapshotTemplateWithContext(context.Background(), request)
}

// CreateVideoSnapshotTemplate
// Create a screenshot template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_CONTAININVALIDSNAPSHOTTIME = "InvalidParameterValue.ContainInvalidSnapshotTime"
//  INVALIDPARAMETERVALUE_INVALIDNAME = "InvalidParameterValue.InvalidName"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDINTERVAL = "InvalidParameterValue.UnsupportedInterval"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDSNAPSHOTTOTAL = "InvalidParameterValue.UnsupportedSnapshotTotal"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDSNAPSHOTTYPE = "InvalidParameterValue.UnsupportedSnapshotType"
//  LIMITEXCEEDED_SNAPSHOTTEMPLATENUMEXCEEDED = "LimitExceeded.SnapshotTemplateNumExceeded"
//  LIMITEXCEEDED_TIMELINESLICEEXCEEDLIMIT = "LimitExceeded.TimelineSliceExceedLimit"
//  UNSUPPORTEDOPERATION_IMAGEDOMAINNOTEXIST = "UnsupportedOperation.ImageDomainNotExist"
func (c *Client) CreateVideoSnapshotTemplateWithContext(ctx context.Context, request *CreateVideoSnapshotTemplateRequest) (response *CreateVideoSnapshotTemplateResponse, err error) {
    if request == nil {
        request = NewCreateVideoSnapshotTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "CreateVideoSnapshotTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVideoSnapshotTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVideoSnapshotTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDomainRequest() (request *DeleteDomainRequest) {
    request = &DeleteDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DeleteDomain")
    
    
    return
}

func NewDeleteDomainResponse() (response *DeleteDomainResponse) {
    response = &DeleteDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDomain
// This API is used to delete domain names.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDomain(request *DeleteDomainRequest) (response *DeleteDomainResponse, err error) {
    return c.DeleteDomainWithContext(context.Background(), request)
}

// DeleteDomain
// This API is used to delete domain names.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDomainWithContext(ctx context.Context, request *DeleteDomainRequest) (response *DeleteDomainResponse, err error) {
    if request == nil {
        request = NewDeleteDomainRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "DeleteDomain")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteOrganizationRequest() (request *DeleteOrganizationRequest) {
    request = &DeleteOrganizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DeleteOrganization")
    
    
    return
}

func NewDeleteOrganizationResponse() (response *DeleteOrganizationResponse) {
    response = &DeleteOrganizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteOrganization
// This API is used to delete an organization.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_ORGLINKDEV = "UnsupportedOperation.OrgLinkDev"
//  UNSUPPORTEDOPERATION_ORGLINKORG = "UnsupportedOperation.OrgLinkOrg"
func (c *Client) DeleteOrganization(request *DeleteOrganizationRequest) (response *DeleteOrganizationResponse, err error) {
    return c.DeleteOrganizationWithContext(context.Background(), request)
}

// DeleteOrganization
// This API is used to delete an organization.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_ORGLINKDEV = "UnsupportedOperation.OrgLinkDev"
//  UNSUPPORTEDOPERATION_ORGLINKORG = "UnsupportedOperation.OrgLinkOrg"
func (c *Client) DeleteOrganizationWithContext(ctx context.Context, request *DeleteOrganizationRequest) (response *DeleteOrganizationResponse, err error) {
    if request == nil {
        request = NewDeleteOrganizationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "DeleteOrganization")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteOrganization require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteOrganizationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRecordBackupPlanRequest() (request *DeleteRecordBackupPlanRequest) {
    request = &DeleteRecordBackupPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DeleteRecordBackupPlan")
    
    
    return
}

func NewDeleteRecordBackupPlanResponse() (response *DeleteRecordBackupPlanResponse) {
    response = &DeleteRecordBackupPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRecordBackupPlan
// Used to delete a video recording cloud template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDPLANID = "InvalidParameterValue.InvalidPlanId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRecordBackupPlan(request *DeleteRecordBackupPlanRequest) (response *DeleteRecordBackupPlanResponse, err error) {
    return c.DeleteRecordBackupPlanWithContext(context.Background(), request)
}

// DeleteRecordBackupPlan
// Used to delete a video recording cloud template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDPLANID = "InvalidParameterValue.InvalidPlanId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRecordBackupPlanWithContext(ctx context.Context, request *DeleteRecordBackupPlanRequest) (response *DeleteRecordBackupPlanResponse, err error) {
    if request == nil {
        request = NewDeleteRecordBackupPlanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "DeleteRecordBackupPlan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRecordBackupPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRecordBackupPlanResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRecordBackupTemplateRequest() (request *DeleteRecordBackupTemplateRequest) {
    request = &DeleteRecordBackupTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DeleteRecordBackupTemplate")
    
    
    return
}

func NewDeleteRecordBackupTemplateResponse() (response *DeleteRecordBackupTemplateResponse) {
    response = &DeleteRecordBackupTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRecordBackupTemplate
// This API is used to delete a video recording cloud template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  RESOURCEINUSE_PLANLINKTEMPLATE = "ResourceInUse.PlanLinkTemplate"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRecordBackupTemplate(request *DeleteRecordBackupTemplateRequest) (response *DeleteRecordBackupTemplateResponse, err error) {
    return c.DeleteRecordBackupTemplateWithContext(context.Background(), request)
}

// DeleteRecordBackupTemplate
// This API is used to delete a video recording cloud template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  RESOURCEINUSE_PLANLINKTEMPLATE = "ResourceInUse.PlanLinkTemplate"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRecordBackupTemplateWithContext(ctx context.Context, request *DeleteRecordBackupTemplateRequest) (response *DeleteRecordBackupTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteRecordBackupTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "DeleteRecordBackupTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRecordBackupTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRecordBackupTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRecordPlanRequest() (request *DeleteRecordPlanRequest) {
    request = &DeleteRecordPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DeleteRecordPlan")
    
    
    return
}

func NewDeleteRecordPlanResponse() (response *DeleteRecordPlanResponse) {
    response = &DeleteRecordPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRecordPlan
// This API is used to delete a real-time cloud migration plan.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_INVALIDPLANID = "InvalidParameterValue.InvalidPlanId"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
func (c *Client) DeleteRecordPlan(request *DeleteRecordPlanRequest) (response *DeleteRecordPlanResponse, err error) {
    return c.DeleteRecordPlanWithContext(context.Background(), request)
}

// DeleteRecordPlan
// This API is used to delete a real-time cloud migration plan.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_INVALIDPLANID = "InvalidParameterValue.InvalidPlanId"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
func (c *Client) DeleteRecordPlanWithContext(ctx context.Context, request *DeleteRecordPlanRequest) (response *DeleteRecordPlanResponse, err error) {
    if request == nil {
        request = NewDeleteRecordPlanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "DeleteRecordPlan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRecordPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRecordPlanResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRecordRetrieveTaskRequest() (request *DeleteRecordRetrieveTaskRequest) {
    request = &DeleteRecordRetrieveTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DeleteRecordRetrieveTask")
    
    
    return
}

func NewDeleteRecordRetrieveTaskResponse() (response *DeleteRecordRetrieveTaskResponse) {
    response = &DeleteRecordRetrieveTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRecordRetrieveTask
// This API is used to delete a retrieval task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_INVALIDRETRIEVETASKID = "InvalidParameterValue.InvalidRetrieveTaskId"
//  RESOURCEINUSE_RETRIEVETASKEXECUTING = "ResourceInUse.RetrieveTaskExecuting"
//  RESOURCENOTFOUND_RETRIEVETASKNOTEXIST = "ResourceNotFound.RetrieveTaskNotExist"
func (c *Client) DeleteRecordRetrieveTask(request *DeleteRecordRetrieveTaskRequest) (response *DeleteRecordRetrieveTaskResponse, err error) {
    return c.DeleteRecordRetrieveTaskWithContext(context.Background(), request)
}

// DeleteRecordRetrieveTask
// This API is used to delete a retrieval task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_INVALIDRETRIEVETASKID = "InvalidParameterValue.InvalidRetrieveTaskId"
//  RESOURCEINUSE_RETRIEVETASKEXECUTING = "ResourceInUse.RetrieveTaskExecuting"
//  RESOURCENOTFOUND_RETRIEVETASKNOTEXIST = "ResourceNotFound.RetrieveTaskNotExist"
func (c *Client) DeleteRecordRetrieveTaskWithContext(ctx context.Context, request *DeleteRecordRetrieveTaskRequest) (response *DeleteRecordRetrieveTaskResponse, err error) {
    if request == nil {
        request = NewDeleteRecordRetrieveTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "DeleteRecordRetrieveTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRecordRetrieveTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRecordRetrieveTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRecordTemplateRequest() (request *DeleteRecordTemplateRequest) {
    request = &DeleteRecordTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DeleteRecordTemplate")
    
    
    return
}

func NewDeleteRecordTemplateResponse() (response *DeleteRecordTemplateResponse) {
    response = &DeleteRecordTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRecordTemplate
// Used to delete a real-time cloud template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  RESOURCEINUSE_PLANLINKTEMPLATE = "ResourceInUse.PlanLinkTemplate"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteRecordTemplate(request *DeleteRecordTemplateRequest) (response *DeleteRecordTemplateResponse, err error) {
    return c.DeleteRecordTemplateWithContext(context.Background(), request)
}

// DeleteRecordTemplate
// Used to delete a real-time cloud template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  RESOURCEINUSE_PLANLINKTEMPLATE = "ResourceInUse.PlanLinkTemplate"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteRecordTemplateWithContext(ctx context.Context, request *DeleteRecordTemplateRequest) (response *DeleteRecordTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteRecordTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "DeleteRecordTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRecordTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRecordTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTaskRequest() (request *DeleteTaskRequest) {
    request = &DeleteTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DeleteTask")
    
    
    return
}

func NewDeleteTaskResponse() (response *DeleteTaskResponse) {
    response = &DeleteTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTask
// Used to delete completed tasks.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  RESOURCEINUSE_PLANLINKTEMPLATE = "ResourceInUse.PlanLinkTemplate"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteTask(request *DeleteTaskRequest) (response *DeleteTaskResponse, err error) {
    return c.DeleteTaskWithContext(context.Background(), request)
}

// DeleteTask
// Used to delete completed tasks.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  RESOURCEINUSE_PLANLINKTEMPLATE = "ResourceInUse.PlanLinkTemplate"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DeleteTaskWithContext(ctx context.Context, request *DeleteTaskRequest) (response *DeleteTaskResponse, err error) {
    if request == nil {
        request = NewDeleteTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "DeleteTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUserDeviceRequest() (request *DeleteUserDeviceRequest) {
    request = &DeleteUserDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DeleteUserDevice")
    
    
    return
}

func NewDeleteUserDeviceResponse() (response *DeleteUserDeviceResponse) {
    response = &DeleteUserDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteUserDevice
// This API is used to delete added devices.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  REGIONERROR_RESOURCEUNREACHABLE = "RegionError.ResourceUnreachable"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteUserDevice(request *DeleteUserDeviceRequest) (response *DeleteUserDeviceResponse, err error) {
    return c.DeleteUserDeviceWithContext(context.Background(), request)
}

// DeleteUserDevice
// This API is used to delete added devices.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  REGIONERROR_RESOURCEUNREACHABLE = "RegionError.ResourceUnreachable"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteUserDeviceWithContext(ctx context.Context, request *DeleteUserDeviceRequest) (response *DeleteUserDeviceResponse, err error) {
    if request == nil {
        request = NewDeleteUserDeviceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "DeleteUserDevice")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUserDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUserDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteVLPlanRequest() (request *DeleteVLPlanRequest) {
    request = &DeleteVLPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DeleteVLPlan")
    
    
    return
}

func NewDeleteVLPlanResponse() (response *DeleteVLPlanResponse) {
    response = &DeleteVLPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteVLPlan
// Delete a vision understanding plan
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_IMAGEVLPLANIDNOTEXISTED = "InvalidParameterValue.ImageVLPlanIDNotExisted"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEVLPLANID = "InvalidParameterValue.InvalidImageVLPlanID"
func (c *Client) DeleteVLPlan(request *DeleteVLPlanRequest) (response *DeleteVLPlanResponse, err error) {
    return c.DeleteVLPlanWithContext(context.Background(), request)
}

// DeleteVLPlan
// Delete a vision understanding plan
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_IMAGEVLPLANIDNOTEXISTED = "InvalidParameterValue.ImageVLPlanIDNotExisted"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEVLPLANID = "InvalidParameterValue.InvalidImageVLPlanID"
func (c *Client) DeleteVLPlanWithContext(ctx context.Context, request *DeleteVLPlanRequest) (response *DeleteVLPlanResponse, err error) {
    if request == nil {
        request = NewDeleteVLPlanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "DeleteVLPlan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteVLPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteVLPlanResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteVLTemplateRequest() (request *DeleteVLTemplateRequest) {
    request = &DeleteVLTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DeleteVLTemplate")
    
    
    return
}

func NewDeleteVLTemplateResponse() (response *DeleteVLTemplateResponse) {
    response = &DeleteVLTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteVLTemplate
// Delete a visual understanding template.
//
// error code that may be returned:
//  FAILEDOPERATION_TEMPLATEINUSE = "FailedOperation.TemplateInUse"
//  INVALIDPARAMETERVALUE_TEMPLATEIDNOTEXISTED = "InvalidParameterValue.TemplateIDNotExisted"
func (c *Client) DeleteVLTemplate(request *DeleteVLTemplateRequest) (response *DeleteVLTemplateResponse, err error) {
    return c.DeleteVLTemplateWithContext(context.Background(), request)
}

// DeleteVLTemplate
// Delete a visual understanding template.
//
// error code that may be returned:
//  FAILEDOPERATION_TEMPLATEINUSE = "FailedOperation.TemplateInUse"
//  INVALIDPARAMETERVALUE_TEMPLATEIDNOTEXISTED = "InvalidParameterValue.TemplateIDNotExisted"
func (c *Client) DeleteVLTemplateWithContext(ctx context.Context, request *DeleteVLTemplateRequest) (response *DeleteVLTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteVLTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "DeleteVLTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteVLTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteVLTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteVideoSnapshotConfigRequest() (request *DeleteVideoSnapshotConfigRequest) {
    request = &DeleteVideoSnapshotConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DeleteVideoSnapshotConfig")
    
    
    return
}

func NewDeleteVideoSnapshotConfigResponse() (response *DeleteVideoSnapshotConfigResponse) {
    response = &DeleteVideoSnapshotConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteVideoSnapshotConfig
// Deletes screenshot configuration
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDCONFIGID = "InvalidParameterValue.InvalidConfigId"
//  OPERATIONDENIED_CONFIGCHANGING = "OperationDenied.ConfigChanging"
//  RESOURCENOTFOUND_ACCESSUNAUTHORIZED = "ResourceNotFound.AccessUnauthorized"
//  RESOURCENOTFOUND_SNAPSHOTCONFIGNOTEXIST = "ResourceNotFound.SnapshotConfigNotExist"
func (c *Client) DeleteVideoSnapshotConfig(request *DeleteVideoSnapshotConfigRequest) (response *DeleteVideoSnapshotConfigResponse, err error) {
    return c.DeleteVideoSnapshotConfigWithContext(context.Background(), request)
}

// DeleteVideoSnapshotConfig
// Deletes screenshot configuration
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDCONFIGID = "InvalidParameterValue.InvalidConfigId"
//  OPERATIONDENIED_CONFIGCHANGING = "OperationDenied.ConfigChanging"
//  RESOURCENOTFOUND_ACCESSUNAUTHORIZED = "ResourceNotFound.AccessUnauthorized"
//  RESOURCENOTFOUND_SNAPSHOTCONFIGNOTEXIST = "ResourceNotFound.SnapshotConfigNotExist"
func (c *Client) DeleteVideoSnapshotConfigWithContext(ctx context.Context, request *DeleteVideoSnapshotConfigRequest) (response *DeleteVideoSnapshotConfigResponse, err error) {
    if request == nil {
        request = NewDeleteVideoSnapshotConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "DeleteVideoSnapshotConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteVideoSnapshotConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteVideoSnapshotConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteVideoSnapshotTemplateRequest() (request *DeleteVideoSnapshotTemplateRequest) {
    request = &DeleteVideoSnapshotTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DeleteVideoSnapshotTemplate")
    
    
    return
}

func NewDeleteVideoSnapshotTemplateResponse() (response *DeleteVideoSnapshotTemplateResponse) {
    response = &DeleteVideoSnapshotTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteVideoSnapshotTemplate
// This API is used to delete a screenshot template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  RESOURCEINUSE_TEMPLATEISUSED = "ResourceInUse.TemplateIsUsed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteVideoSnapshotTemplate(request *DeleteVideoSnapshotTemplateRequest) (response *DeleteVideoSnapshotTemplateResponse, err error) {
    return c.DeleteVideoSnapshotTemplateWithContext(context.Background(), request)
}

// DeleteVideoSnapshotTemplate
// This API is used to delete a screenshot template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  RESOURCEINUSE_TEMPLATEISUSED = "ResourceInUse.TemplateIsUsed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteVideoSnapshotTemplateWithContext(ctx context.Context, request *DeleteVideoSnapshotTemplateRequest) (response *DeleteVideoSnapshotTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteVideoSnapshotTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "DeleteVideoSnapshotTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteVideoSnapshotTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteVideoSnapshotTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCNAMERequest() (request *DescribeCNAMERequest) {
    request = &DescribeCNAMERequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeCNAME")
    
    
    return
}

func NewDescribeCNAMEResponse() (response *DescribeCNAMEResponse) {
    response = &DescribeCNAMEResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCNAME
// This API is used to get the CNAME value based on the service node.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDDOMAINTYPE = "InvalidParameterValue.InvalidDomainType"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCNAME(request *DescribeCNAMERequest) (response *DescribeCNAMEResponse, err error) {
    return c.DescribeCNAMEWithContext(context.Background(), request)
}

// DescribeCNAME
// This API is used to get the CNAME value based on the service node.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDDOMAINTYPE = "InvalidParameterValue.InvalidDomainType"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCNAMEWithContext(ctx context.Context, request *DescribeCNAMERequest) (response *DescribeCNAMEResponse, err error) {
    if request == nil {
        request = NewDescribeCNAMERequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "DescribeCNAME")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCNAME require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCNAMEResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceChannelRequest() (request *DescribeDeviceChannelRequest) {
    request = &DescribeDeviceChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeDeviceChannel")
    
    
    return
}

func NewDescribeDeviceChannelResponse() (response *DescribeDeviceChannelResponse) {
    response = &DescribeDeviceChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeviceChannel
// This API is used to query the channels of a device.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDeviceChannel(request *DescribeDeviceChannelRequest) (response *DescribeDeviceChannelResponse, err error) {
    return c.DescribeDeviceChannelWithContext(context.Background(), request)
}

// DescribeDeviceChannel
// This API is used to query the channels of a device.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDeviceChannelWithContext(ctx context.Context, request *DescribeDeviceChannelRequest) (response *DescribeDeviceChannelResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceChannelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "DescribeDeviceChannel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceChannelResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDevicePresetRequest() (request *DescribeDevicePresetRequest) {
    request = &DescribeDevicePresetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeDevicePreset")
    
    
    return
}

func NewDescribeDevicePresetResponse() (response *DescribeDevicePresetResponse) {
    response = &DescribeDevicePresetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDevicePreset
// This API is used to query preset position information of device channels.
//
// error code that may be returned:
//  FAILEDOPERATION_DEVICERESPONSETIMEOUT = "FailedOperation.DeviceResponseTimeOut"
//  FAILEDOPERATION_DEVICERESULTTIMEOUT = "FailedOperation.DeviceResultTimeOut"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CHANNELOFFLINE = "ResourceUnavailable.ChannelOffline"
//  RESOURCEUNAVAILABLE_DEVDISABLE = "ResourceUnavailable.DevDisable"
//  RESOURCEUNAVAILABLE_DEVNOREGISTER = "ResourceUnavailable.DevNoRegister"
//  RESOURCEUNAVAILABLE_DEVOFFLINE = "ResourceUnavailable.DevOffline"
func (c *Client) DescribeDevicePreset(request *DescribeDevicePresetRequest) (response *DescribeDevicePresetResponse, err error) {
    return c.DescribeDevicePresetWithContext(context.Background(), request)
}

// DescribeDevicePreset
// This API is used to query preset position information of device channels.
//
// error code that may be returned:
//  FAILEDOPERATION_DEVICERESPONSETIMEOUT = "FailedOperation.DeviceResponseTimeOut"
//  FAILEDOPERATION_DEVICERESULTTIMEOUT = "FailedOperation.DeviceResultTimeOut"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CHANNELOFFLINE = "ResourceUnavailable.ChannelOffline"
//  RESOURCEUNAVAILABLE_DEVDISABLE = "ResourceUnavailable.DevDisable"
//  RESOURCEUNAVAILABLE_DEVNOREGISTER = "ResourceUnavailable.DevNoRegister"
//  RESOURCEUNAVAILABLE_DEVOFFLINE = "ResourceUnavailable.DevOffline"
func (c *Client) DescribeDevicePresetWithContext(ctx context.Context, request *DescribeDevicePresetRequest) (response *DescribeDevicePresetResponse, err error) {
    if request == nil {
        request = NewDescribeDevicePresetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "DescribeDevicePreset")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDevicePreset require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDevicePresetResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainRequest() (request *DescribeDomainRequest) {
    request = &DescribeDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeDomain")
    
    
    return
}

func NewDescribeDomainResponse() (response *DescribeDomainResponse) {
    response = &DescribeDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDomain
// This API is used to query the list of added domains.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDomain(request *DescribeDomainRequest) (response *DescribeDomainResponse, err error) {
    return c.DescribeDomainWithContext(context.Background(), request)
}

// DescribeDomain
// This API is used to query the list of added domains.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDomainWithContext(ctx context.Context, request *DescribeDomainRequest) (response *DescribeDomainResponse, err error) {
    if request == nil {
        request = NewDescribeDomainRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "DescribeDomain")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainRegionRequest() (request *DescribeDomainRegionRequest) {
    request = &DescribeDomainRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeDomainRegion")
    
    
    return
}

func NewDescribeDomainRegionResponse() (response *DescribeDomainRegionResponse) {
    response = &DescribeDomainRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDomainRegion
// This API is used to query the service nodes that can be bound when a user adds a domain name. The result is all service nodes supported by the platform. Each service node can only be bound to one domain name.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDomainRegion(request *DescribeDomainRegionRequest) (response *DescribeDomainRegionResponse, err error) {
    return c.DescribeDomainRegionWithContext(context.Background(), request)
}

// DescribeDomainRegion
// This API is used to query the service nodes that can be bound when a user adds a domain name. The result is all service nodes supported by the platform. Each service node can only be bound to one domain name.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDomainRegionWithContext(ctx context.Context, request *DescribeDomainRegionRequest) (response *DescribeDomainRegionResponse, err error) {
    if request == nil {
        request = NewDescribeDomainRegionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "DescribeDomainRegion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomainRegion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainRegionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGBDeviceAddrRequest() (request *DescribeGBDeviceAddrRequest) {
    request = &DescribeGBDeviceAddrRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeGBDeviceAddr")
    
    
    return
}

func NewDescribeGBDeviceAddrResponse() (response *DescribeGBDeviceAddrResponse) {
    response = &DescribeGBDeviceAddrResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGBDeviceAddr
// This API is used to obtain the public network address of a national standard device.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CONTAININVALIDDEVICEID = "InvalidParameterValue.ContainInvalidDeviceId"
//  LIMITEXCEEDED_GBDEVICENUMEXCEEDED = "LimitExceeded.GBDeviceNumExceeded"
func (c *Client) DescribeGBDeviceAddr(request *DescribeGBDeviceAddrRequest) (response *DescribeGBDeviceAddrResponse, err error) {
    return c.DescribeGBDeviceAddrWithContext(context.Background(), request)
}

// DescribeGBDeviceAddr
// This API is used to obtain the public network address of a national standard device.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CONTAININVALIDDEVICEID = "InvalidParameterValue.ContainInvalidDeviceId"
//  LIMITEXCEEDED_GBDEVICENUMEXCEEDED = "LimitExceeded.GBDeviceNumExceeded"
func (c *Client) DescribeGBDeviceAddrWithContext(ctx context.Context, request *DescribeGBDeviceAddrRequest) (response *DescribeGBDeviceAddrResponse, err error) {
    if request == nil {
        request = NewDescribeGBDeviceAddrRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "DescribeGBDeviceAddr")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGBDeviceAddr require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGBDeviceAddrResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganizationRequest() (request *DescribeOrganizationRequest) {
    request = &DescribeOrganizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeOrganization")
    
    
    return
}

func NewDescribeOrganizationResponse() (response *DescribeOrganizationResponse) {
    response = &DescribeOrganizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOrganization
// Used for querying organizations.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeOrganization(request *DescribeOrganizationRequest) (response *DescribeOrganizationResponse, err error) {
    return c.DescribeOrganizationWithContext(context.Background(), request)
}

// DescribeOrganization
// Used for querying organizations.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeOrganizationWithContext(ctx context.Context, request *DescribeOrganizationRequest) (response *DescribeOrganizationResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "DescribeOrganization")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganization require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordBackupPlanRequest() (request *DescribeRecordBackupPlanRequest) {
    request = &DescribeRecordBackupPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeRecordBackupPlan")
    
    
    return
}

func NewDescribeRecordBackupPlanResponse() (response *DescribeRecordBackupPlanResponse) {
    response = &DescribeRecordBackupPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRecordBackupPlan
// This API is used to query recording cloud migration plan details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDPLANID = "InvalidParameterValue.InvalidPlanId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRecordBackupPlan(request *DescribeRecordBackupPlanRequest) (response *DescribeRecordBackupPlanResponse, err error) {
    return c.DescribeRecordBackupPlanWithContext(context.Background(), request)
}

// DescribeRecordBackupPlan
// This API is used to query recording cloud migration plan details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDPLANID = "InvalidParameterValue.InvalidPlanId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRecordBackupPlanWithContext(ctx context.Context, request *DescribeRecordBackupPlanRequest) (response *DescribeRecordBackupPlanResponse, err error) {
    if request == nil {
        request = NewDescribeRecordBackupPlanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "DescribeRecordBackupPlan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordBackupPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordBackupPlanResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordBackupTemplateRequest() (request *DescribeRecordBackupTemplateRequest) {
    request = &DescribeRecordBackupTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeRecordBackupTemplate")
    
    
    return
}

func NewDescribeRecordBackupTemplateResponse() (response *DescribeRecordBackupTemplateResponse) {
    response = &DescribeRecordBackupTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRecordBackupTemplate
// This API is used to query video recording cloud template details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRecordBackupTemplate(request *DescribeRecordBackupTemplateRequest) (response *DescribeRecordBackupTemplateResponse, err error) {
    return c.DescribeRecordBackupTemplateWithContext(context.Background(), request)
}

// DescribeRecordBackupTemplate
// This API is used to query video recording cloud template details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRecordBackupTemplateWithContext(ctx context.Context, request *DescribeRecordBackupTemplateRequest) (response *DescribeRecordBackupTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeRecordBackupTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "DescribeRecordBackupTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordBackupTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordBackupTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordFileRequest() (request *DescribeRecordFileRequest) {
    request = &DescribeRecordFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeRecordFile")
    
    
    return
}

func NewDescribeRecordFileResponse() (response *DescribeRecordFileResponse) {
    response = &DescribeRecordFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRecordFile
// This API is used to query the device cloud recording timeline information, which is the cloud storage video recording after setting a recording schedule for video upload to cloud.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ILLEGALCHANNELID = "InvalidParameterValue.IllegalChannelId"
//  INVALIDPARAMETERVALUE_ILLEGALDEVICEID = "InvalidParameterValue.IllegalDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDSTARTTIMEORENDTIME = "InvalidParameterValue.InvalidStartTimeOrEndTime"
//  INVALIDPARAMETERVALUE_NONSAMEDAY = "InvalidParameterValue.NonSameDay"
//  INVALIDPARAMETERVALUE_STARTTIMEGREATERTHANOREQUALENDTIME = "InvalidParameterValue.StartTimeGreaterThanOrEqualEndTime"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
func (c *Client) DescribeRecordFile(request *DescribeRecordFileRequest) (response *DescribeRecordFileResponse, err error) {
    return c.DescribeRecordFileWithContext(context.Background(), request)
}

// DescribeRecordFile
// This API is used to query the device cloud recording timeline information, which is the cloud storage video recording after setting a recording schedule for video upload to cloud.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ILLEGALCHANNELID = "InvalidParameterValue.IllegalChannelId"
//  INVALIDPARAMETERVALUE_ILLEGALDEVICEID = "InvalidParameterValue.IllegalDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDSTARTTIMEORENDTIME = "InvalidParameterValue.InvalidStartTimeOrEndTime"
//  INVALIDPARAMETERVALUE_NONSAMEDAY = "InvalidParameterValue.NonSameDay"
//  INVALIDPARAMETERVALUE_STARTTIMEGREATERTHANOREQUALENDTIME = "InvalidParameterValue.StartTimeGreaterThanOrEqualEndTime"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
func (c *Client) DescribeRecordFileWithContext(ctx context.Context, request *DescribeRecordFileRequest) (response *DescribeRecordFileResponse, err error) {
    if request == nil {
        request = NewDescribeRecordFileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "DescribeRecordFile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordFileResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordPlanRequest() (request *DescribeRecordPlanRequest) {
    request = &DescribeRecordPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeRecordPlan")
    
    
    return
}

func NewDescribeRecordPlanResponse() (response *DescribeRecordPlanResponse) {
    response = &DescribeRecordPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRecordPlan
// This API is used to query real-time cloud migration plan details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_INVALIDPLANID = "InvalidParameterValue.InvalidPlanId"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
func (c *Client) DescribeRecordPlan(request *DescribeRecordPlanRequest) (response *DescribeRecordPlanResponse, err error) {
    return c.DescribeRecordPlanWithContext(context.Background(), request)
}

// DescribeRecordPlan
// This API is used to query real-time cloud migration plan details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_INVALIDPLANID = "InvalidParameterValue.InvalidPlanId"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
func (c *Client) DescribeRecordPlanWithContext(ctx context.Context, request *DescribeRecordPlanRequest) (response *DescribeRecordPlanResponse, err error) {
    if request == nil {
        request = NewDescribeRecordPlanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "DescribeRecordPlan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordPlanResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordPlaybackUrlRequest() (request *DescribeRecordPlaybackUrlRequest) {
    request = &DescribeRecordPlaybackUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeRecordPlaybackUrl")
    
    
    return
}

func NewDescribeRecordPlaybackUrlResponse() (response *DescribeRecordPlaybackUrlResponse) {
    response = &DescribeRecordPlaybackUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRecordPlaybackUrl
// This API is used to obtain the cloud recording playback url address.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ILLEGALCHANNELID = "InvalidParameterValue.IllegalChannelId"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_INVALIDSTARTTIMEORENDTIME = "InvalidParameterValue.InvalidStartTimeOrEndTime"
//  INVALIDPARAMETERVALUE_NONSAMEDAY = "InvalidParameterValue.NonSameDay"
//  INVALIDPARAMETERVALUE_STARTTIMEGREATERTHANOREQUALENDTIME = "InvalidParameterValue.StartTimeGreaterThanOrEqualEndTime"
//  RESOURCENOTFOUND_CHANNELNOTEXIST = "ResourceNotFound.ChannelNotExist"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNSUPPORTEDOPERATION_UNOPENEDINTRANETSERVICESINREGION = "UnsupportedOperation.UnopenedIntranetServicesInRegion"
func (c *Client) DescribeRecordPlaybackUrl(request *DescribeRecordPlaybackUrlRequest) (response *DescribeRecordPlaybackUrlResponse, err error) {
    return c.DescribeRecordPlaybackUrlWithContext(context.Background(), request)
}

// DescribeRecordPlaybackUrl
// This API is used to obtain the cloud recording playback url address.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ILLEGALCHANNELID = "InvalidParameterValue.IllegalChannelId"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_INVALIDSTARTTIMEORENDTIME = "InvalidParameterValue.InvalidStartTimeOrEndTime"
//  INVALIDPARAMETERVALUE_NONSAMEDAY = "InvalidParameterValue.NonSameDay"
//  INVALIDPARAMETERVALUE_STARTTIMEGREATERTHANOREQUALENDTIME = "InvalidParameterValue.StartTimeGreaterThanOrEqualEndTime"
//  RESOURCENOTFOUND_CHANNELNOTEXIST = "ResourceNotFound.ChannelNotExist"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNSUPPORTEDOPERATION_UNOPENEDINTRANETSERVICESINREGION = "UnsupportedOperation.UnopenedIntranetServicesInRegion"
func (c *Client) DescribeRecordPlaybackUrlWithContext(ctx context.Context, request *DescribeRecordPlaybackUrlRequest) (response *DescribeRecordPlaybackUrlResponse, err error) {
    if request == nil {
        request = NewDescribeRecordPlaybackUrlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "DescribeRecordPlaybackUrl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordPlaybackUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordPlaybackUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordRetrieveTaskRequest() (request *DescribeRecordRetrieveTaskRequest) {
    request = &DescribeRecordRetrieveTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeRecordRetrieveTask")
    
    
    return
}

func NewDescribeRecordRetrieveTaskResponse() (response *DescribeRecordRetrieveTaskResponse) {
    response = &DescribeRecordRetrieveTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRecordRetrieveTask
// Used to query the details of a cloud recording retrieval task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_INVALIDRETRIEVETASKID = "InvalidParameterValue.InvalidRetrieveTaskId"
//  RESOURCENOTFOUND_RETRIEVETASKNOTEXIST = "ResourceNotFound.RetrieveTaskNotExist"
func (c *Client) DescribeRecordRetrieveTask(request *DescribeRecordRetrieveTaskRequest) (response *DescribeRecordRetrieveTaskResponse, err error) {
    return c.DescribeRecordRetrieveTaskWithContext(context.Background(), request)
}

// DescribeRecordRetrieveTask
// Used to query the details of a cloud recording retrieval task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_INVALIDRETRIEVETASKID = "InvalidParameterValue.InvalidRetrieveTaskId"
//  RESOURCENOTFOUND_RETRIEVETASKNOTEXIST = "ResourceNotFound.RetrieveTaskNotExist"
func (c *Client) DescribeRecordRetrieveTaskWithContext(ctx context.Context, request *DescribeRecordRetrieveTaskRequest) (response *DescribeRecordRetrieveTaskResponse, err error) {
    if request == nil {
        request = NewDescribeRecordRetrieveTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "DescribeRecordRetrieveTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordRetrieveTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordRetrieveTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordTemplateRequest() (request *DescribeRecordTemplateRequest) {
    request = &DescribeRecordTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeRecordTemplate")
    
    
    return
}

func NewDescribeRecordTemplateResponse() (response *DescribeRecordTemplateResponse) {
    response = &DescribeRecordTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRecordTemplate
// This API is used to query real-time cloud migration template details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeRecordTemplate(request *DescribeRecordTemplateRequest) (response *DescribeRecordTemplateResponse, err error) {
    return c.DescribeRecordTemplateWithContext(context.Background(), request)
}

// DescribeRecordTemplate
// This API is used to query real-time cloud migration template details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) DescribeRecordTemplateWithContext(ctx context.Context, request *DescribeRecordTemplateRequest) (response *DescribeRecordTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeRecordTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "DescribeRecordTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamAuthRequest() (request *DescribeStreamAuthRequest) {
    request = &DescribeStreamAuthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeStreamAuth")
    
    
    return
}

func NewDescribeStreamAuthResponse() (response *DescribeStreamAuthResponse) {
    response = &DescribeStreamAuthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStreamAuth
// Used to query push/pull stream authentication configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeStreamAuth(request *DescribeStreamAuthRequest) (response *DescribeStreamAuthResponse, err error) {
    return c.DescribeStreamAuthWithContext(context.Background(), request)
}

// DescribeStreamAuth
// Used to query push/pull stream authentication configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeStreamAuthWithContext(ctx context.Context, request *DescribeStreamAuthRequest) (response *DescribeStreamAuthResponse, err error) {
    if request == nil {
        request = NewDescribeStreamAuthRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "DescribeStreamAuth")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamAuth require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamAuthResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskRequest() (request *DescribeTaskRequest) {
    request = &DescribeTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeTask")
    
    
    return
}

func NewDescribeTaskResponse() (response *DescribeTaskResponse) {
    response = &DescribeTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTask
// This API is used to query task details.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEERROR = "FailedOperation.DatabaseError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETER_INVALIDPARAMETERFORMAT = "InvalidParameter.InvalidParameterFormat"
//  INVALIDPARAMETER_REQUIREDHEADERPARAMETEREMPTY = "InvalidParameter.RequiredHeaderParameterEmpty"
//  INVALIDPARAMETER_TASKIDNOTEXIST = "InvalidParameter.TaskIdNotExist"
//  INVALIDPARAMETERVALUE_TASKTYPENOTSUPPORTED = "InvalidParameterValue.TaskTypeNotSupported"
func (c *Client) DescribeTask(request *DescribeTaskRequest) (response *DescribeTaskResponse, err error) {
    return c.DescribeTaskWithContext(context.Background(), request)
}

// DescribeTask
// This API is used to query task details.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEERROR = "FailedOperation.DatabaseError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETER_INVALIDPARAMETERFORMAT = "InvalidParameter.InvalidParameterFormat"
//  INVALIDPARAMETER_REQUIREDHEADERPARAMETEREMPTY = "InvalidParameter.RequiredHeaderParameterEmpty"
//  INVALIDPARAMETER_TASKIDNOTEXIST = "InvalidParameter.TaskIdNotExist"
//  INVALIDPARAMETERVALUE_TASKTYPENOTSUPPORTED = "InvalidParameterValue.TaskTypeNotSupported"
func (c *Client) DescribeTaskWithContext(ctx context.Context, request *DescribeTaskRequest) (response *DescribeTaskResponse, err error) {
    if request == nil {
        request = NewDescribeTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "DescribeTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserDeviceRequest() (request *DescribeUserDeviceRequest) {
    request = &DescribeUserDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeUserDevice")
    
    
    return
}

func NewDescribeUserDeviceResponse() (response *DescribeUserDeviceResponse) {
    response = &DescribeUserDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserDevice
// This API is used to query detailed information of a device.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeUserDevice(request *DescribeUserDeviceRequest) (response *DescribeUserDeviceResponse, err error) {
    return c.DescribeUserDeviceWithContext(context.Background(), request)
}

// DescribeUserDevice
// This API is used to query detailed information of a device.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeUserDeviceWithContext(ctx context.Context, request *DescribeUserDeviceRequest) (response *DescribeUserDeviceResponse, err error) {
    if request == nil {
        request = NewDescribeUserDeviceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "DescribeUserDevice")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserDeviceListRequest() (request *DescribeUserDeviceListRequest) {
    request = &DescribeUserDeviceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeUserDeviceList")
    
    
    return
}

func NewDescribeUserDeviceListResponse() (response *DescribeUserDeviceListResponse) {
    response = &DescribeUserDeviceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserDeviceList
// This API is used to query device details in batches.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_DEICEIDQUANTITYEXCEEDEDLIMIT = "LimitExceeded.DeiceIdQuantityExceededLimit"
func (c *Client) DescribeUserDeviceList(request *DescribeUserDeviceListRequest) (response *DescribeUserDeviceListResponse, err error) {
    return c.DescribeUserDeviceListWithContext(context.Background(), request)
}

// DescribeUserDeviceList
// This API is used to query device details in batches.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_DEICEIDQUANTITYEXCEEDEDLIMIT = "LimitExceeded.DeiceIdQuantityExceededLimit"
func (c *Client) DescribeUserDeviceListWithContext(ctx context.Context, request *DescribeUserDeviceListRequest) (response *DescribeUserDeviceListResponse, err error) {
    if request == nil {
        request = NewDescribeUserDeviceListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "DescribeUserDeviceList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserDeviceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserDeviceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVLPlanRequest() (request *DescribeVLPlanRequest) {
    request = &DescribeVLPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeVLPlan")
    
    
    return
}

func NewDescribeVLPlanResponse() (response *DescribeVLPlanResponse) {
    response = &DescribeVLPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVLPlan
// Query the plan details of visual understanding.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_IMAGEVLPLANIDNOTEXISTED = "InvalidParameterValue.ImageVLPlanIDNotExisted"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEVLPLANID = "InvalidParameterValue.InvalidImageVLPlanID"
func (c *Client) DescribeVLPlan(request *DescribeVLPlanRequest) (response *DescribeVLPlanResponse, err error) {
    return c.DescribeVLPlanWithContext(context.Background(), request)
}

// DescribeVLPlan
// Query the plan details of visual understanding.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_IMAGEVLPLANIDNOTEXISTED = "InvalidParameterValue.ImageVLPlanIDNotExisted"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEVLPLANID = "InvalidParameterValue.InvalidImageVLPlanID"
func (c *Client) DescribeVLPlanWithContext(ctx context.Context, request *DescribeVLPlanRequest) (response *DescribeVLPlanResponse, err error) {
    if request == nil {
        request = NewDescribeVLPlanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "DescribeVLPlan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVLPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVLPlanResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVLPlanResultRequest() (request *DescribeVLPlanResultRequest) {
    request = &DescribeVLPlanResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeVLPlanResult")
    
    
    return
}

func NewDescribeVLPlanResultResponse() (response *DescribeVLPlanResultResponse) {
    response = &DescribeVLPlanResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVLPlanResult
// Query the result details of a vision understanding plan
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDVLPLANRESULTID = "InvalidParameterValue.InvalidVLPlanResultID"
//  INVALIDPARAMETERVALUE_VLPLANRESULTIDNOTEXISTED = "InvalidParameterValue.VLPlanResultIDNotExisted"
func (c *Client) DescribeVLPlanResult(request *DescribeVLPlanResultRequest) (response *DescribeVLPlanResultResponse, err error) {
    return c.DescribeVLPlanResultWithContext(context.Background(), request)
}

// DescribeVLPlanResult
// Query the result details of a vision understanding plan
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDVLPLANRESULTID = "InvalidParameterValue.InvalidVLPlanResultID"
//  INVALIDPARAMETERVALUE_VLPLANRESULTIDNOTEXISTED = "InvalidParameterValue.VLPlanResultIDNotExisted"
func (c *Client) DescribeVLPlanResultWithContext(ctx context.Context, request *DescribeVLPlanResultRequest) (response *DescribeVLPlanResultResponse, err error) {
    if request == nil {
        request = NewDescribeVLPlanResultRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "DescribeVLPlanResult")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVLPlanResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVLPlanResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVLTemplateRequest() (request *DescribeVLTemplateRequest) {
    request = &DescribeVLTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeVLTemplate")
    
    
    return
}

func NewDescribeVLTemplateResponse() (response *DescribeVLTemplateResponse) {
    response = &DescribeVLTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVLTemplate
// Query the details of a visual understanding template.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDVLPLANRESULTID = "InvalidParameterValue.InvalidVLPlanResultID"
//  INVALIDPARAMETERVALUE_VLPLANRESULTIDNOTEXISTED = "InvalidParameterValue.VLPlanResultIDNotExisted"
func (c *Client) DescribeVLTemplate(request *DescribeVLTemplateRequest) (response *DescribeVLTemplateResponse, err error) {
    return c.DescribeVLTemplateWithContext(context.Background(), request)
}

// DescribeVLTemplate
// Query the details of a visual understanding template.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDVLPLANRESULTID = "InvalidParameterValue.InvalidVLPlanResultID"
//  INVALIDPARAMETERVALUE_VLPLANRESULTIDNOTEXISTED = "InvalidParameterValue.VLPlanResultIDNotExisted"
func (c *Client) DescribeVLTemplateWithContext(ctx context.Context, request *DescribeVLTemplateRequest) (response *DescribeVLTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeVLTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "DescribeVLTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVLTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVLTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVideoBitRateRequest() (request *DescribeVideoBitRateRequest) {
    request = &DescribeVideoBitRateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeVideoBitRate")
    
    
    return
}

func NewDescribeVideoBitRateResponse() (response *DescribeVideoBitRateResponse) {
    response = &DescribeVideoBitRateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVideoBitRate
// This API is used to obtain the bitrate info of a video channel.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CONTAININVALIDCHANNELID = "InvalidParameterValue.ContainInvalidChannelId"
//  LIMITEXCEEDED_CHANNELNUMEXCEEDED = "LimitExceeded.ChannelNumExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVideoBitRate(request *DescribeVideoBitRateRequest) (response *DescribeVideoBitRateResponse, err error) {
    return c.DescribeVideoBitRateWithContext(context.Background(), request)
}

// DescribeVideoBitRate
// This API is used to obtain the bitrate info of a video channel.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CONTAININVALIDCHANNELID = "InvalidParameterValue.ContainInvalidChannelId"
//  LIMITEXCEEDED_CHANNELNUMEXCEEDED = "LimitExceeded.ChannelNumExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVideoBitRateWithContext(ctx context.Context, request *DescribeVideoBitRateRequest) (response *DescribeVideoBitRateResponse, err error) {
    if request == nil {
        request = NewDescribeVideoBitRateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "DescribeVideoBitRate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVideoBitRate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVideoBitRateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVideoDownloadUrlRequest() (request *DescribeVideoDownloadUrlRequest) {
    request = &DescribeVideoDownloadUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeVideoDownloadUrl")
    
    
    return
}

func NewDescribeVideoDownloadUrlResponse() (response *DescribeVideoDownloadUrlResponse) {
    response = &DescribeVideoDownloadUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVideoDownloadUrl
// This API is used to obtain the download url of cloud recordings.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_DOWNLOADURLERROR = "InvalidParameter.DownloadUrlError"
//  INVALIDPARAMETER_DOWNLOADURLHASEXPIRED = "InvalidParameter.DownloadUrlHasExpired"
//  INVALIDPARAMETERVALUE_INVALIDBEGINTIME = "InvalidParameterValue.InvalidBeginTime"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_INVALIDENDTIME = "InvalidParameterValue.InvalidEndTime"
//  INVALIDPARAMETERVALUE_INVALIDFILETYPE = "InvalidParameterValue.InvalidFileType"
//  INVALIDPARAMETERVALUE_NONSAMEDAY = "InvalidParameterValue.NonSameDay"
//  INVALIDPARAMETERVALUE_OUTOFTIMERANGE = "InvalidParameterValue.OutOfTimeRange"
//  OPERATIONDENIED_BANDWIDTHLIMITZERO = "OperationDenied.BandwidthLimitZero"
//  OPERATIONDENIED_CONCURRENTDOWNLOADSOVERLIMIT = "OperationDenied.ConcurrentDownloadsOverLimit"
//  OPERATIONDENIED_CONNECTSLIMITZERO = "OperationDenied.ConnectsLimitZero"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_VIDEONOTFOUND = "ResourceNotFound.VideoNotFound"
//  RESOURCEUNAVAILABLE_VIDEOARCHIVED = "ResourceUnavailable.VideoArchived"
//  UNSUPPORTEDOPERATION_UNOPENEDINTRANETSERVICESINREGION = "UnsupportedOperation.UnopenedIntranetServicesInRegion"
func (c *Client) DescribeVideoDownloadUrl(request *DescribeVideoDownloadUrlRequest) (response *DescribeVideoDownloadUrlResponse, err error) {
    return c.DescribeVideoDownloadUrlWithContext(context.Background(), request)
}

// DescribeVideoDownloadUrl
// This API is used to obtain the download url of cloud recordings.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_DOWNLOADURLERROR = "InvalidParameter.DownloadUrlError"
//  INVALIDPARAMETER_DOWNLOADURLHASEXPIRED = "InvalidParameter.DownloadUrlHasExpired"
//  INVALIDPARAMETERVALUE_INVALIDBEGINTIME = "InvalidParameterValue.InvalidBeginTime"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_INVALIDENDTIME = "InvalidParameterValue.InvalidEndTime"
//  INVALIDPARAMETERVALUE_INVALIDFILETYPE = "InvalidParameterValue.InvalidFileType"
//  INVALIDPARAMETERVALUE_NONSAMEDAY = "InvalidParameterValue.NonSameDay"
//  INVALIDPARAMETERVALUE_OUTOFTIMERANGE = "InvalidParameterValue.OutOfTimeRange"
//  OPERATIONDENIED_BANDWIDTHLIMITZERO = "OperationDenied.BandwidthLimitZero"
//  OPERATIONDENIED_CONCURRENTDOWNLOADSOVERLIMIT = "OperationDenied.ConcurrentDownloadsOverLimit"
//  OPERATIONDENIED_CONNECTSLIMITZERO = "OperationDenied.ConnectsLimitZero"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_VIDEONOTFOUND = "ResourceNotFound.VideoNotFound"
//  RESOURCEUNAVAILABLE_VIDEOARCHIVED = "ResourceUnavailable.VideoArchived"
//  UNSUPPORTEDOPERATION_UNOPENEDINTRANETSERVICESINREGION = "UnsupportedOperation.UnopenedIntranetServicesInRegion"
func (c *Client) DescribeVideoDownloadUrlWithContext(ctx context.Context, request *DescribeVideoDownloadUrlRequest) (response *DescribeVideoDownloadUrlResponse, err error) {
    if request == nil {
        request = NewDescribeVideoDownloadUrlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "DescribeVideoDownloadUrl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVideoDownloadUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVideoDownloadUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVideoSnapshotConfigRequest() (request *DescribeVideoSnapshotConfigRequest) {
    request = &DescribeVideoSnapshotConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeVideoSnapshotConfig")
    
    
    return
}

func NewDescribeVideoSnapshotConfigResponse() (response *DescribeVideoSnapshotConfigResponse) {
    response = &DescribeVideoSnapshotConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVideoSnapshotConfig
// Queries screenshot configuration details
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDCONFIGID = "InvalidParameterValue.InvalidConfigId"
//  RESOURCENOTFOUND_SNAPSHOTCONFIGNOTEXIST = "ResourceNotFound.SnapshotConfigNotExist"
func (c *Client) DescribeVideoSnapshotConfig(request *DescribeVideoSnapshotConfigRequest) (response *DescribeVideoSnapshotConfigResponse, err error) {
    return c.DescribeVideoSnapshotConfigWithContext(context.Background(), request)
}

// DescribeVideoSnapshotConfig
// Queries screenshot configuration details
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDCONFIGID = "InvalidParameterValue.InvalidConfigId"
//  RESOURCENOTFOUND_SNAPSHOTCONFIGNOTEXIST = "ResourceNotFound.SnapshotConfigNotExist"
func (c *Client) DescribeVideoSnapshotConfigWithContext(ctx context.Context, request *DescribeVideoSnapshotConfigRequest) (response *DescribeVideoSnapshotConfigResponse, err error) {
    if request == nil {
        request = NewDescribeVideoSnapshotConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "DescribeVideoSnapshotConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVideoSnapshotConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVideoSnapshotConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVideoSnapshotTemplateRequest() (request *DescribeVideoSnapshotTemplateRequest) {
    request = &DescribeVideoSnapshotTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "DescribeVideoSnapshotTemplate")
    
    
    return
}

func NewDescribeVideoSnapshotTemplateResponse() (response *DescribeVideoSnapshotTemplateResponse) {
    response = &DescribeVideoSnapshotTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVideoSnapshotTemplate
// Get screenshot template details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVideoSnapshotTemplate(request *DescribeVideoSnapshotTemplateRequest) (response *DescribeVideoSnapshotTemplateResponse, err error) {
    return c.DescribeVideoSnapshotTemplateWithContext(context.Background(), request)
}

// DescribeVideoSnapshotTemplate
// Get screenshot template details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVideoSnapshotTemplateWithContext(ctx context.Context, request *DescribeVideoSnapshotTemplateRequest) (response *DescribeVideoSnapshotTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeVideoSnapshotTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "DescribeVideoSnapshotTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVideoSnapshotTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVideoSnapshotTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewListDeviceSnapshotsRequest() (request *ListDeviceSnapshotsRequest) {
    request = &ListDeviceSnapshotsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ListDeviceSnapshots")
    
    
    return
}

func NewListDeviceSnapshotsResponse() (response *ListDeviceSnapshotsResponse) {
    response = &ListDeviceSnapshotsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListDeviceSnapshots
// Get device capture result list
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_PAGESIZEEXCEEDLIMIT = "InvalidParameterValue.PageSizeExceedLimit"
//  INVALIDPARAMETERVALUE_STARTOVERENDTIME = "InvalidParameterValue.StartOverEndTime"
//  INVALIDPARAMETERVALUE_STARTOVERNOWTIME = "InvalidParameterValue.StartOverNowTime"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ListDeviceSnapshots(request *ListDeviceSnapshotsRequest) (response *ListDeviceSnapshotsResponse, err error) {
    return c.ListDeviceSnapshotsWithContext(context.Background(), request)
}

// ListDeviceSnapshots
// Get device capture result list
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_PAGESIZEEXCEEDLIMIT = "InvalidParameterValue.PageSizeExceedLimit"
//  INVALIDPARAMETERVALUE_STARTOVERENDTIME = "InvalidParameterValue.StartOverEndTime"
//  INVALIDPARAMETERVALUE_STARTOVERNOWTIME = "InvalidParameterValue.StartOverNowTime"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ListDeviceSnapshotsWithContext(ctx context.Context, request *ListDeviceSnapshotsRequest) (response *ListDeviceSnapshotsResponse, err error) {
    if request == nil {
        request = NewListDeviceSnapshotsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "ListDeviceSnapshots")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListDeviceSnapshots require credential")
    }

    request.SetContext(ctx)
    
    response = NewListDeviceSnapshotsResponse()
    err = c.Send(request, response)
    return
}

func NewListDevicesRequest() (request *ListDevicesRequest) {
    request = &ListDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ListDevices")
    
    
    return
}

func NewListDevicesResponse() (response *ListDevicesResponse) {
    response = &ListDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListDevices
// This API is used to obtain the device list under the corresponding organization.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDACCESSPROTOCOL = "InvalidParameterValue.InvalidAccessProtocol"
//  INVALIDPARAMETERVALUE_INVALIDCLUSTERID = "InvalidParameterValue.InvalidClusterId"
//  INVALIDPARAMETERVALUE_INVALIDDEVICESTATUS = "InvalidParameterValue.InvalidDeviceStatus"
//  INVALIDPARAMETERVALUE_INVALIDDEVICETYPE = "InvalidParameterValue.InvalidDeviceType"
//  INVALIDPARAMETERVALUE_INVALIDKEYWORD = "InvalidParameterValue.InvalidKeyword"
//  INVALIDPARAMETERVALUE_INVALIDORGANIZATIONID = "InvalidParameterValue.InvalidOrganizationId"
//  INVALIDPARAMETERVALUE_INVALIDPAGENUMBER = "InvalidParameterValue.InvalidPageNumber"
//  INVALIDPARAMETERVALUE_INVALIDPAGESIZE = "InvalidParameterValue.InvalidPageSize"
func (c *Client) ListDevices(request *ListDevicesRequest) (response *ListDevicesResponse, err error) {
    return c.ListDevicesWithContext(context.Background(), request)
}

// ListDevices
// This API is used to obtain the device list under the corresponding organization.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDACCESSPROTOCOL = "InvalidParameterValue.InvalidAccessProtocol"
//  INVALIDPARAMETERVALUE_INVALIDCLUSTERID = "InvalidParameterValue.InvalidClusterId"
//  INVALIDPARAMETERVALUE_INVALIDDEVICESTATUS = "InvalidParameterValue.InvalidDeviceStatus"
//  INVALIDPARAMETERVALUE_INVALIDDEVICETYPE = "InvalidParameterValue.InvalidDeviceType"
//  INVALIDPARAMETERVALUE_INVALIDKEYWORD = "InvalidParameterValue.InvalidKeyword"
//  INVALIDPARAMETERVALUE_INVALIDORGANIZATIONID = "InvalidParameterValue.InvalidOrganizationId"
//  INVALIDPARAMETERVALUE_INVALIDPAGENUMBER = "InvalidParameterValue.InvalidPageNumber"
//  INVALIDPARAMETERVALUE_INVALIDPAGESIZE = "InvalidParameterValue.InvalidPageSize"
func (c *Client) ListDevicesWithContext(ctx context.Context, request *ListDevicesRequest) (response *ListDevicesResponse, err error) {
    if request == nil {
        request = NewListDevicesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "ListDevices")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListDevices require credential")
    }

    request.SetContext(ctx)
    
    response = NewListDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewListOrganizationChannelsRequest() (request *ListOrganizationChannelsRequest) {
    request = &ListOrganizationChannelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ListOrganizationChannels")
    
    
    return
}

func NewListOrganizationChannelsResponse() (response *ListOrganizationChannelsResponse) {
    response = &ListOrganizationChannelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListOrganizationChannels
// This API is used to query the channel list under the organizational unit directory.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_EMPTYNAME = "InvalidParameterValue.EmptyName"
//  INVALIDPARAMETERVALUE_EMPTYORGANIZATIONID = "InvalidParameterValue.EmptyOrganizationId"
//  INVALIDPARAMETERVALUE_ILLEGALNAME = "InvalidParameterValue.IllegalName"
//  INVALIDPARAMETERVALUE_ILLEGALORGANIZATIONID = "InvalidParameterValue.IllegalOrganizationId"
//  INVALIDPARAMETERVALUE_INVALIDPAGENUMBER = "InvalidParameterValue.InvalidPageNumber"
//  INVALIDPARAMETERVALUE_INVALIDPAGESIZE = "InvalidParameterValue.InvalidPageSize"
//  INVALIDPARAMETERVALUE_TOOLONGNAME = "InvalidParameterValue.TooLongName"
//  RESOURCENOTFOUND_ORGANIZATIONIDNOTEXIST = "ResourceNotFound.OrganizationIdNotExist"
func (c *Client) ListOrganizationChannels(request *ListOrganizationChannelsRequest) (response *ListOrganizationChannelsResponse, err error) {
    return c.ListOrganizationChannelsWithContext(context.Background(), request)
}

// ListOrganizationChannels
// This API is used to query the channel list under the organizational unit directory.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_EMPTYNAME = "InvalidParameterValue.EmptyName"
//  INVALIDPARAMETERVALUE_EMPTYORGANIZATIONID = "InvalidParameterValue.EmptyOrganizationId"
//  INVALIDPARAMETERVALUE_ILLEGALNAME = "InvalidParameterValue.IllegalName"
//  INVALIDPARAMETERVALUE_ILLEGALORGANIZATIONID = "InvalidParameterValue.IllegalOrganizationId"
//  INVALIDPARAMETERVALUE_INVALIDPAGENUMBER = "InvalidParameterValue.InvalidPageNumber"
//  INVALIDPARAMETERVALUE_INVALIDPAGESIZE = "InvalidParameterValue.InvalidPageSize"
//  INVALIDPARAMETERVALUE_TOOLONGNAME = "InvalidParameterValue.TooLongName"
//  RESOURCENOTFOUND_ORGANIZATIONIDNOTEXIST = "ResourceNotFound.OrganizationIdNotExist"
func (c *Client) ListOrganizationChannelsWithContext(ctx context.Context, request *ListOrganizationChannelsRequest) (response *ListOrganizationChannelsResponse, err error) {
    if request == nil {
        request = NewListOrganizationChannelsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "ListOrganizationChannels")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListOrganizationChannels require credential")
    }

    request.SetContext(ctx)
    
    response = NewListOrganizationChannelsResponse()
    err = c.Send(request, response)
    return
}

func NewListRecordBackupPlanDevicesRequest() (request *ListRecordBackupPlanDevicesRequest) {
    request = &ListRecordBackupPlanDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ListRecordBackupPlanDevices")
    
    
    return
}

func NewListRecordBackupPlanDevicesResponse() (response *ListRecordBackupPlanDevicesResponse) {
    response = &ListRecordBackupPlanDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListRecordBackupPlanDevices
// This API is used to query the device channel list under a recording cloud migration plan.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELNAME = "InvalidParameterValue.InvalidChannelName"
//  INVALIDPARAMETERVALUE_INVALIDDEVICENAME = "InvalidParameterValue.InvalidDeviceName"
//  INVALIDPARAMETERVALUE_INVALIDORGNAME = "InvalidParameterValue.InvalidOrgName"
//  INVALIDPARAMETERVALUE_INVALIDPLANID = "InvalidParameterValue.InvalidPlanId"
func (c *Client) ListRecordBackupPlanDevices(request *ListRecordBackupPlanDevicesRequest) (response *ListRecordBackupPlanDevicesResponse, err error) {
    return c.ListRecordBackupPlanDevicesWithContext(context.Background(), request)
}

// ListRecordBackupPlanDevices
// This API is used to query the device channel list under a recording cloud migration plan.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELNAME = "InvalidParameterValue.InvalidChannelName"
//  INVALIDPARAMETERVALUE_INVALIDDEVICENAME = "InvalidParameterValue.InvalidDeviceName"
//  INVALIDPARAMETERVALUE_INVALIDORGNAME = "InvalidParameterValue.InvalidOrgName"
//  INVALIDPARAMETERVALUE_INVALIDPLANID = "InvalidParameterValue.InvalidPlanId"
func (c *Client) ListRecordBackupPlanDevicesWithContext(ctx context.Context, request *ListRecordBackupPlanDevicesRequest) (response *ListRecordBackupPlanDevicesResponse, err error) {
    if request == nil {
        request = NewListRecordBackupPlanDevicesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "ListRecordBackupPlanDevices")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListRecordBackupPlanDevices require credential")
    }

    request.SetContext(ctx)
    
    response = NewListRecordBackupPlanDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewListRecordBackupPlansRequest() (request *ListRecordBackupPlansRequest) {
    request = &ListRecordBackupPlansRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ListRecordBackupPlans")
    
    
    return
}

func NewListRecordBackupPlansResponse() (response *ListRecordBackupPlansResponse) {
    response = &ListRecordBackupPlansResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListRecordBackupPlans
// This API is used to query the recording cloud migration plan list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListRecordBackupPlans(request *ListRecordBackupPlansRequest) (response *ListRecordBackupPlansResponse, err error) {
    return c.ListRecordBackupPlansWithContext(context.Background(), request)
}

// ListRecordBackupPlans
// This API is used to query the recording cloud migration plan list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListRecordBackupPlansWithContext(ctx context.Context, request *ListRecordBackupPlansRequest) (response *ListRecordBackupPlansResponse, err error) {
    if request == nil {
        request = NewListRecordBackupPlansRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "ListRecordBackupPlans")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListRecordBackupPlans require credential")
    }

    request.SetContext(ctx)
    
    response = NewListRecordBackupPlansResponse()
    err = c.Send(request, response)
    return
}

func NewListRecordBackupTemplatesRequest() (request *ListRecordBackupTemplatesRequest) {
    request = &ListRecordBackupTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ListRecordBackupTemplates")
    
    
    return
}

func NewListRecordBackupTemplatesResponse() (response *ListRecordBackupTemplatesResponse) {
    response = &ListRecordBackupTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListRecordBackupTemplates
// This API is used to query the video recording cloud template list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListRecordBackupTemplates(request *ListRecordBackupTemplatesRequest) (response *ListRecordBackupTemplatesResponse, err error) {
    return c.ListRecordBackupTemplatesWithContext(context.Background(), request)
}

// ListRecordBackupTemplates
// This API is used to query the video recording cloud template list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListRecordBackupTemplatesWithContext(ctx context.Context, request *ListRecordBackupTemplatesRequest) (response *ListRecordBackupTemplatesResponse, err error) {
    if request == nil {
        request = NewListRecordBackupTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "ListRecordBackupTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListRecordBackupTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewListRecordBackupTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewListRecordPlanDevicesRequest() (request *ListRecordPlanDevicesRequest) {
    request = &ListRecordPlanDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ListRecordPlanDevices")
    
    
    return
}

func NewListRecordPlanDevicesResponse() (response *ListRecordPlanDevicesResponse) {
    response = &ListRecordPlanDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListRecordPlanDevices
// This API is used to query the device channel list under a real-time cloud migration plan.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_EMPTYNAME = "InvalidParameterValue.EmptyName"
//  INVALIDPARAMETERVALUE_ILLEGALNAME = "InvalidParameterValue.IllegalName"
//  INVALIDPARAMETERVALUE_INVALIDPAGENUMBER = "InvalidParameterValue.InvalidPageNumber"
//  INVALIDPARAMETERVALUE_INVALIDPAGESIZE = "InvalidParameterValue.InvalidPageSize"
//  INVALIDPARAMETERVALUE_INVALIDPLANID = "InvalidParameterValue.InvalidPlanId"
//  INVALIDPARAMETERVALUE_TOOLONGNAME = "InvalidParameterValue.TooLongName"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
func (c *Client) ListRecordPlanDevices(request *ListRecordPlanDevicesRequest) (response *ListRecordPlanDevicesResponse, err error) {
    return c.ListRecordPlanDevicesWithContext(context.Background(), request)
}

// ListRecordPlanDevices
// This API is used to query the device channel list under a real-time cloud migration plan.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETERVALUE_EMPTYNAME = "InvalidParameterValue.EmptyName"
//  INVALIDPARAMETERVALUE_ILLEGALNAME = "InvalidParameterValue.IllegalName"
//  INVALIDPARAMETERVALUE_INVALIDPAGENUMBER = "InvalidParameterValue.InvalidPageNumber"
//  INVALIDPARAMETERVALUE_INVALIDPAGESIZE = "InvalidParameterValue.InvalidPageSize"
//  INVALIDPARAMETERVALUE_INVALIDPLANID = "InvalidParameterValue.InvalidPlanId"
//  INVALIDPARAMETERVALUE_TOOLONGNAME = "InvalidParameterValue.TooLongName"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
func (c *Client) ListRecordPlanDevicesWithContext(ctx context.Context, request *ListRecordPlanDevicesRequest) (response *ListRecordPlanDevicesResponse, err error) {
    if request == nil {
        request = NewListRecordPlanDevicesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "ListRecordPlanDevices")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListRecordPlanDevices require credential")
    }

    request.SetContext(ctx)
    
    response = NewListRecordPlanDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewListRecordPlansRequest() (request *ListRecordPlansRequest) {
    request = &ListRecordPlansRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ListRecordPlans")
    
    
    return
}

func NewListRecordPlansResponse() (response *ListRecordPlansResponse) {
    response = &ListRecordPlansResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListRecordPlans
// This API is used to query the real-time cloud migration plan list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListRecordPlans(request *ListRecordPlansRequest) (response *ListRecordPlansResponse, err error) {
    return c.ListRecordPlansWithContext(context.Background(), request)
}

// ListRecordPlans
// This API is used to query the real-time cloud migration plan list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListRecordPlansWithContext(ctx context.Context, request *ListRecordPlansRequest) (response *ListRecordPlansResponse, err error) {
    if request == nil {
        request = NewListRecordPlansRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "ListRecordPlans")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListRecordPlans require credential")
    }

    request.SetContext(ctx)
    
    response = NewListRecordPlansResponse()
    err = c.Send(request, response)
    return
}

func NewListRecordRetrieveTasksRequest() (request *ListRecordRetrieveTasksRequest) {
    request = &ListRecordRetrieveTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ListRecordRetrieveTasks")
    
    
    return
}

func NewListRecordRetrieveTasksResponse() (response *ListRecordRetrieveTasksResponse) {
    response = &ListRecordRetrieveTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListRecordRetrieveTasks
// This API is used to query the retrieval task list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListRecordRetrieveTasks(request *ListRecordRetrieveTasksRequest) (response *ListRecordRetrieveTasksResponse, err error) {
    return c.ListRecordRetrieveTasksWithContext(context.Background(), request)
}

// ListRecordRetrieveTasks
// This API is used to query the retrieval task list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListRecordRetrieveTasksWithContext(ctx context.Context, request *ListRecordRetrieveTasksRequest) (response *ListRecordRetrieveTasksResponse, err error) {
    if request == nil {
        request = NewListRecordRetrieveTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "ListRecordRetrieveTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListRecordRetrieveTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewListRecordRetrieveTasksResponse()
    err = c.Send(request, response)
    return
}

func NewListRecordTemplatesRequest() (request *ListRecordTemplatesRequest) {
    request = &ListRecordTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ListRecordTemplates")
    
    
    return
}

func NewListRecordTemplatesResponse() (response *ListRecordTemplatesResponse) {
    response = &ListRecordTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListRecordTemplates
// This API is used to query the real-time cloud migration template list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListRecordTemplates(request *ListRecordTemplatesRequest) (response *ListRecordTemplatesResponse, err error) {
    return c.ListRecordTemplatesWithContext(context.Background(), request)
}

// ListRecordTemplates
// This API is used to query the real-time cloud migration template list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListRecordTemplatesWithContext(ctx context.Context, request *ListRecordTemplatesRequest) (response *ListRecordTemplatesResponse, err error) {
    if request == nil {
        request = NewListRecordTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "ListRecordTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListRecordTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewListRecordTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewListSubTasksRequest() (request *ListSubTasksRequest) {
    request = &ListSubTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ListSubTasks")
    
    
    return
}

func NewListSubTasksResponse() (response *ListSubTasksResponse) {
    response = &ListSubTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListSubTasks
// This API is used to query the subtask list of a query task.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEERROR = "FailedOperation.DatabaseError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETER_INVALIDPARAMETERFORMAT = "InvalidParameter.InvalidParameterFormat"
//  INVALIDPARAMETER_REQUIREDHEADERPARAMETEREMPTY = "InvalidParameter.RequiredHeaderParameterEmpty"
//  INVALIDPARAMETER_TASKIDNOTEXIST = "InvalidParameter.TaskIdNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPAGEPARAMETER = "InvalidParameterValue.InvalidPageParameter"
//  INVALIDPARAMETERVALUE_INVALIDSTATUS = "InvalidParameterValue.InvalidStatus"
//  INVALIDPARAMETERVALUE_TASKTYPENOTSUPPORTED = "InvalidParameterValue.TaskTypeNotSupported"
func (c *Client) ListSubTasks(request *ListSubTasksRequest) (response *ListSubTasksResponse, err error) {
    return c.ListSubTasksWithContext(context.Background(), request)
}

// ListSubTasks
// This API is used to query the subtask list of a query task.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEERROR = "FailedOperation.DatabaseError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETER_INVALIDPARAMETERFORMAT = "InvalidParameter.InvalidParameterFormat"
//  INVALIDPARAMETER_REQUIREDHEADERPARAMETEREMPTY = "InvalidParameter.RequiredHeaderParameterEmpty"
//  INVALIDPARAMETER_TASKIDNOTEXIST = "InvalidParameter.TaskIdNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPAGEPARAMETER = "InvalidParameterValue.InvalidPageParameter"
//  INVALIDPARAMETERVALUE_INVALIDSTATUS = "InvalidParameterValue.InvalidStatus"
//  INVALIDPARAMETERVALUE_TASKTYPENOTSUPPORTED = "InvalidParameterValue.TaskTypeNotSupported"
func (c *Client) ListSubTasksWithContext(ctx context.Context, request *ListSubTasksRequest) (response *ListSubTasksResponse, err error) {
    if request == nil {
        request = NewListSubTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "ListSubTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListSubTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewListSubTasksResponse()
    err = c.Send(request, response)
    return
}

func NewListTasksRequest() (request *ListTasksRequest) {
    request = &ListTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ListTasks")
    
    
    return
}

func NewListTasksResponse() (response *ListTasksResponse) {
    response = &ListTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListTasks
// This API is used to query the batch task and simple task list.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEERROR = "FailedOperation.DatabaseError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETER_INVALIDACTION = "InvalidParameter.InvalidAction"
//  INVALIDPARAMETER_INVALIDPARAMETERFORMAT = "InvalidParameter.InvalidParameterFormat"
//  INVALIDPARAMETER_REQUIREDHEADERPARAMETEREMPTY = "InvalidParameter.RequiredHeaderParameterEmpty"
//  INVALIDPARAMETERVALUE_INVALIDPAGEPARAMETER = "InvalidParameterValue.InvalidPageParameter"
//  INVALIDPARAMETERVALUE_INVALIDSTATUS = "InvalidParameterValue.InvalidStatus"
//  INVALIDPARAMETERVALUE_TASKTYPENOTSUPPORTED = "InvalidParameterValue.TaskTypeNotSupported"
func (c *Client) ListTasks(request *ListTasksRequest) (response *ListTasksResponse, err error) {
    return c.ListTasksWithContext(context.Background(), request)
}

// ListTasks
// This API is used to query the batch task and simple task list.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEERROR = "FailedOperation.DatabaseError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETER_INVALIDACTION = "InvalidParameter.InvalidAction"
//  INVALIDPARAMETER_INVALIDPARAMETERFORMAT = "InvalidParameter.InvalidParameterFormat"
//  INVALIDPARAMETER_REQUIREDHEADERPARAMETEREMPTY = "InvalidParameter.RequiredHeaderParameterEmpty"
//  INVALIDPARAMETERVALUE_INVALIDPAGEPARAMETER = "InvalidParameterValue.InvalidPageParameter"
//  INVALIDPARAMETERVALUE_INVALIDSTATUS = "InvalidParameterValue.InvalidStatus"
//  INVALIDPARAMETERVALUE_TASKTYPENOTSUPPORTED = "InvalidParameterValue.TaskTypeNotSupported"
func (c *Client) ListTasksWithContext(ctx context.Context, request *ListTasksRequest) (response *ListTasksResponse, err error) {
    if request == nil {
        request = NewListTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "ListTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewListTasksResponse()
    err = c.Send(request, response)
    return
}

func NewListVLPlansRequest() (request *ListVLPlansRequest) {
    request = &ListVLPlansRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ListVLPlans")
    
    
    return
}

func NewListVLPlansResponse() (response *ListVLPlansResponse) {
    response = &ListVLPlansResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListVLPlans
// Get the visual understanding plan list.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDPAGENUMBER = "InvalidParameterValue.InvalidPageNumber"
//  INVALIDPARAMETERVALUE_INVALIDPAGESIZE = "InvalidParameterValue.InvalidPageSize"
func (c *Client) ListVLPlans(request *ListVLPlansRequest) (response *ListVLPlansResponse, err error) {
    return c.ListVLPlansWithContext(context.Background(), request)
}

// ListVLPlans
// Get the visual understanding plan list.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDPAGENUMBER = "InvalidParameterValue.InvalidPageNumber"
//  INVALIDPARAMETERVALUE_INVALIDPAGESIZE = "InvalidParameterValue.InvalidPageSize"
func (c *Client) ListVLPlansWithContext(ctx context.Context, request *ListVLPlansRequest) (response *ListVLPlansResponse, err error) {
    if request == nil {
        request = NewListVLPlansRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "ListVLPlans")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListVLPlans require credential")
    }

    request.SetContext(ctx)
    
    response = NewListVLPlansResponse()
    err = c.Send(request, response)
    return
}

func NewListVLTemplatesRequest() (request *ListVLTemplatesRequest) {
    request = &ListVLTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ListVLTemplates")
    
    
    return
}

func NewListVLTemplatesResponse() (response *ListVLTemplatesResponse) {
    response = &ListVLTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListVLTemplates
// Get a list of visual understanding templates.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDCAPTURETYPE = "InvalidParameterValue.InvalidCaptureType"
//  INVALIDPARAMETERVALUE_INVALIDMEDIATYPE = "InvalidParameterValue.InvalidMediaType"
//  INVALIDPARAMETERVALUE_INVALIDPAGENUMBER = "InvalidParameterValue.InvalidPageNumber"
//  INVALIDPARAMETERVALUE_INVALIDPAGESIZE = "InvalidParameterValue.InvalidPageSize"
func (c *Client) ListVLTemplates(request *ListVLTemplatesRequest) (response *ListVLTemplatesResponse, err error) {
    return c.ListVLTemplatesWithContext(context.Background(), request)
}

// ListVLTemplates
// Get a list of visual understanding templates.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDCAPTURETYPE = "InvalidParameterValue.InvalidCaptureType"
//  INVALIDPARAMETERVALUE_INVALIDMEDIATYPE = "InvalidParameterValue.InvalidMediaType"
//  INVALIDPARAMETERVALUE_INVALIDPAGENUMBER = "InvalidParameterValue.InvalidPageNumber"
//  INVALIDPARAMETERVALUE_INVALIDPAGESIZE = "InvalidParameterValue.InvalidPageSize"
func (c *Client) ListVLTemplatesWithContext(ctx context.Context, request *ListVLTemplatesRequest) (response *ListVLTemplatesResponse, err error) {
    if request == nil {
        request = NewListVLTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "ListVLTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListVLTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewListVLTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewListVideoDownloadTaskRequest() (request *ListVideoDownloadTaskRequest) {
    request = &ListVideoDownloadTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ListVideoDownloadTask")
    
    
    return
}

func NewListVideoDownloadTaskResponse() (response *ListVideoDownloadTaskResponse) {
    response = &ListVideoDownloadTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListVideoDownloadTask
// This API is used to query the recording download task list in this store.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDCAPTURETYPE = "InvalidParameterValue.InvalidCaptureType"
//  INVALIDPARAMETERVALUE_INVALIDMEDIATYPE = "InvalidParameterValue.InvalidMediaType"
//  INVALIDPARAMETERVALUE_INVALIDPAGENUMBER = "InvalidParameterValue.InvalidPageNumber"
//  INVALIDPARAMETERVALUE_INVALIDPAGESIZE = "InvalidParameterValue.InvalidPageSize"
func (c *Client) ListVideoDownloadTask(request *ListVideoDownloadTaskRequest) (response *ListVideoDownloadTaskResponse, err error) {
    return c.ListVideoDownloadTaskWithContext(context.Background(), request)
}

// ListVideoDownloadTask
// This API is used to query the recording download task list in this store.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDCAPTURETYPE = "InvalidParameterValue.InvalidCaptureType"
//  INVALIDPARAMETERVALUE_INVALIDMEDIATYPE = "InvalidParameterValue.InvalidMediaType"
//  INVALIDPARAMETERVALUE_INVALIDPAGENUMBER = "InvalidParameterValue.InvalidPageNumber"
//  INVALIDPARAMETERVALUE_INVALIDPAGESIZE = "InvalidParameterValue.InvalidPageSize"
func (c *Client) ListVideoDownloadTaskWithContext(ctx context.Context, request *ListVideoDownloadTaskRequest) (response *ListVideoDownloadTaskResponse, err error) {
    if request == nil {
        request = NewListVideoDownloadTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "ListVideoDownloadTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListVideoDownloadTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewListVideoDownloadTaskResponse()
    err = c.Send(request, response)
    return
}

func NewListVideoSnapshotTemplateRequest() (request *ListVideoSnapshotTemplateRequest) {
    request = &ListVideoSnapshotTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ListVideoSnapshotTemplate")
    
    
    return
}

func NewListVideoSnapshotTemplateResponse() (response *ListVideoSnapshotTemplateResponse) {
    response = &ListVideoSnapshotTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListVideoSnapshotTemplate
// This API is used to search the screenshot template details list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_PAGESIZEEXCEEDLIMIT = "InvalidParameterValue.PageSizeExceedLimit"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDSNAPSHOTTYPE = "InvalidParameterValue.UnsupportedSnapshotType"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDSORTRULE = "InvalidParameterValue.UnsupportedSortRule"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ListVideoSnapshotTemplate(request *ListVideoSnapshotTemplateRequest) (response *ListVideoSnapshotTemplateResponse, err error) {
    return c.ListVideoSnapshotTemplateWithContext(context.Background(), request)
}

// ListVideoSnapshotTemplate
// This API is used to search the screenshot template details list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_PAGESIZEEXCEEDLIMIT = "InvalidParameterValue.PageSizeExceedLimit"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDSNAPSHOTTYPE = "InvalidParameterValue.UnsupportedSnapshotType"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDSORTRULE = "InvalidParameterValue.UnsupportedSortRule"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ListVideoSnapshotTemplateWithContext(ctx context.Context, request *ListVideoSnapshotTemplateRequest) (response *ListVideoSnapshotTemplateResponse, err error) {
    if request == nil {
        request = NewListVideoSnapshotTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "ListVideoSnapshotTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListVideoSnapshotTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewListVideoSnapshotTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVLPlanRequest() (request *ModifyVLPlanRequest) {
    request = &ModifyVLPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ModifyVLPlan")
    
    
    return
}

func NewModifyVLPlanResponse() (response *ModifyVLPlanResponse) {
    response = &ModifyVLPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyVLPlan
// Modify a vision understanding plan
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_IMAGEVLPLANIDNOTEXISTED = "InvalidParameterValue.ImageVLPlanIDNotExisted"
//  INVALIDPARAMETERVALUE_IMAGEVLPLANNAMECONTAINILLEGALCHARACTER = "InvalidParameterValue.ImageVLPlanNameContainIllegalCharacter"
//  INVALIDPARAMETERVALUE_INVALIDCALLBACKTOKEN = "InvalidParameterValue.InvalidCallbackToken"
//  INVALIDPARAMETERVALUE_INVALIDCALLBACKURL = "InvalidParameterValue.InvalidCallbackUrl"
//  INVALIDPARAMETERVALUE_INVALIDFILEEXPIRES = "InvalidParameterValue.InvalidFileExpires"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEVLPLANDESCRIBE = "InvalidParameterValue.InvalidImageVLPlanDescribe"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEVLPLANID = "InvalidParameterValue.InvalidImageVLPlanID"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEVLPLANNAME = "InvalidParameterValue.InvalidImageVLPlanName"
//  INVALIDPARAMETERVALUE_TEMPLATEIDNOTEXISTED = "InvalidParameterValue.TemplateIDNotExisted"
func (c *Client) ModifyVLPlan(request *ModifyVLPlanRequest) (response *ModifyVLPlanResponse, err error) {
    return c.ModifyVLPlanWithContext(context.Background(), request)
}

// ModifyVLPlan
// Modify a vision understanding plan
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_IMAGEVLPLANIDNOTEXISTED = "InvalidParameterValue.ImageVLPlanIDNotExisted"
//  INVALIDPARAMETERVALUE_IMAGEVLPLANNAMECONTAINILLEGALCHARACTER = "InvalidParameterValue.ImageVLPlanNameContainIllegalCharacter"
//  INVALIDPARAMETERVALUE_INVALIDCALLBACKTOKEN = "InvalidParameterValue.InvalidCallbackToken"
//  INVALIDPARAMETERVALUE_INVALIDCALLBACKURL = "InvalidParameterValue.InvalidCallbackUrl"
//  INVALIDPARAMETERVALUE_INVALIDFILEEXPIRES = "InvalidParameterValue.InvalidFileExpires"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEVLPLANDESCRIBE = "InvalidParameterValue.InvalidImageVLPlanDescribe"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEVLPLANID = "InvalidParameterValue.InvalidImageVLPlanID"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEVLPLANNAME = "InvalidParameterValue.InvalidImageVLPlanName"
//  INVALIDPARAMETERVALUE_TEMPLATEIDNOTEXISTED = "InvalidParameterValue.TemplateIDNotExisted"
func (c *Client) ModifyVLPlanWithContext(ctx context.Context, request *ModifyVLPlanRequest) (response *ModifyVLPlanResponse, err error) {
    if request == nil {
        request = NewModifyVLPlanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "ModifyVLPlan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVLPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVLPlanResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVideoSnapshotConfigRequest() (request *ModifyVideoSnapshotConfigRequest) {
    request = &ModifyVideoSnapshotConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "ModifyVideoSnapshotConfig")
    
    
    return
}

func NewModifyVideoSnapshotConfigResponse() (response *ModifyVideoSnapshotConfigResponse) {
    response = &ModifyVideoSnapshotConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyVideoSnapshotConfig
// Update screenshot configuration
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_CONTAININVALIDTOKEN = "InvalidParameterValue.ContainInvalidToken"
//  INVALIDPARAMETERVALUE_INVALIDCALLBACKURL = "InvalidParameterValue.InvalidCallbackUrl"
//  INVALIDPARAMETERVALUE_INVALIDCONFIGID = "InvalidParameterValue.InvalidConfigId"
//  INVALIDPARAMETERVALUE_INVALIDNAME = "InvalidParameterValue.InvalidName"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDFILEEXPIRETIME = "InvalidParameterValue.UnsupportedFileExpireTime"
//  LIMITEXCEEDED_SNAPSHOTCONFIGNUMEXCEEDED = "LimitExceeded.SnapshotConfigNumExceeded"
//  OPERATIONDENIED_CONFIGCHANGING = "OperationDenied.ConfigChanging"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SNAPSHOTCONFIGNOTEXIST = "ResourceNotFound.SnapshotConfigNotExist"
func (c *Client) ModifyVideoSnapshotConfig(request *ModifyVideoSnapshotConfigRequest) (response *ModifyVideoSnapshotConfigResponse, err error) {
    return c.ModifyVideoSnapshotConfigWithContext(context.Background(), request)
}

// ModifyVideoSnapshotConfig
// Update screenshot configuration
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_CONTAININVALIDTOKEN = "InvalidParameterValue.ContainInvalidToken"
//  INVALIDPARAMETERVALUE_INVALIDCALLBACKURL = "InvalidParameterValue.InvalidCallbackUrl"
//  INVALIDPARAMETERVALUE_INVALIDCONFIGID = "InvalidParameterValue.InvalidConfigId"
//  INVALIDPARAMETERVALUE_INVALIDNAME = "InvalidParameterValue.InvalidName"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDFILEEXPIRETIME = "InvalidParameterValue.UnsupportedFileExpireTime"
//  LIMITEXCEEDED_SNAPSHOTCONFIGNUMEXCEEDED = "LimitExceeded.SnapshotConfigNumExceeded"
//  OPERATIONDENIED_CONFIGCHANGING = "OperationDenied.ConfigChanging"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SNAPSHOTCONFIGNOTEXIST = "ResourceNotFound.SnapshotConfigNotExist"
func (c *Client) ModifyVideoSnapshotConfigWithContext(ctx context.Context, request *ModifyVideoSnapshotConfigRequest) (response *ModifyVideoSnapshotConfigResponse, err error) {
    if request == nil {
        request = NewModifyVideoSnapshotConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "ModifyVideoSnapshotConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVideoSnapshotConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVideoSnapshotConfigResponse()
    err = c.Send(request, response)
    return
}

func NewPlayRecordRequest() (request *PlayRecordRequest) {
    request = &PlayRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "PlayRecord")
    
    
    return
}

func NewPlayRecordResponse() (response *PlayRecordResponse) {
    response = &PlayRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PlayRecord
// This API is used to obtain the local device recording URL address.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ENDTIMEZERO = "InvalidParameterValue.EndTimeZero"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_STARTOVERENDTIME = "InvalidParameterValue.StartOverEndTime"
//  INVALIDPARAMETERVALUE_STARTOVERNOWTIME = "InvalidParameterValue.StartOverNowTime"
//  INVALIDPARAMETERVALUE_STARTTIMEGREATERTHANOREQUALENDTIME = "InvalidParameterValue.StartTimeGreaterThanOrEqualEndTime"
//  INVALIDPARAMETERVALUE_STARTTIMEZERO = "InvalidParameterValue.StartTimeZero"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_STREAMTYPEORRESOLUTION = "UnsupportedOperation.StreamTypeOrResolution"
func (c *Client) PlayRecord(request *PlayRecordRequest) (response *PlayRecordResponse, err error) {
    return c.PlayRecordWithContext(context.Background(), request)
}

// PlayRecord
// This API is used to obtain the local device recording URL address.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ENDTIMEZERO = "InvalidParameterValue.EndTimeZero"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_STARTOVERENDTIME = "InvalidParameterValue.StartOverEndTime"
//  INVALIDPARAMETERVALUE_STARTOVERNOWTIME = "InvalidParameterValue.StartOverNowTime"
//  INVALIDPARAMETERVALUE_STARTTIMEGREATERTHANOREQUALENDTIME = "InvalidParameterValue.StartTimeGreaterThanOrEqualEndTime"
//  INVALIDPARAMETERVALUE_STARTTIMEZERO = "InvalidParameterValue.StartTimeZero"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_STREAMTYPEORRESOLUTION = "UnsupportedOperation.StreamTypeOrResolution"
func (c *Client) PlayRecordWithContext(ctx context.Context, request *PlayRecordRequest) (response *PlayRecordResponse, err error) {
    if request == nil {
        request = NewPlayRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "PlayRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("PlayRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewPlayRecordResponse()
    err = c.Send(request, response)
    return
}

func NewQueryForbidPlayChannelListRequest() (request *QueryForbidPlayChannelListRequest) {
    request = &QueryForbidPlayChannelListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "QueryForbidPlayChannelList")
    
    
    return
}

func NewQueryForbidPlayChannelListResponse() (response *QueryForbidPlayChannelListResponse) {
    response = &QueryForbidPlayChannelListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryForbidPlayChannelList
// Query the disabled channel list
//
// error code that may be returned:
//  INTERNALERROR_SERVERINTERNALERROR = "InternalError.ServerInternalError"
//  INVALIDPARAMETERVALUE_INVALIDPAGENUMBER = "InvalidParameterValue.InvalidPageNumber"
//  INVALIDPARAMETERVALUE_INVALIDPAGESIZE = "InvalidParameterValue.InvalidPageSize"
//  INVALIDPARAMETERVALUE_INVALIDUSERPARAM = "InvalidParameterValue.InvalidUserParam"
//  MISSINGPARAMETER_MISSINGUSERINFO = "MissingParameter.MissingUserInfo"
//  OPERATIONDENIED_NOWHITELIST = "OperationDenied.NoWhitelist"
func (c *Client) QueryForbidPlayChannelList(request *QueryForbidPlayChannelListRequest) (response *QueryForbidPlayChannelListResponse, err error) {
    return c.QueryForbidPlayChannelListWithContext(context.Background(), request)
}

// QueryForbidPlayChannelList
// Query the disabled channel list
//
// error code that may be returned:
//  INTERNALERROR_SERVERINTERNALERROR = "InternalError.ServerInternalError"
//  INVALIDPARAMETERVALUE_INVALIDPAGENUMBER = "InvalidParameterValue.InvalidPageNumber"
//  INVALIDPARAMETERVALUE_INVALIDPAGESIZE = "InvalidParameterValue.InvalidPageSize"
//  INVALIDPARAMETERVALUE_INVALIDUSERPARAM = "InvalidParameterValue.InvalidUserParam"
//  MISSINGPARAMETER_MISSINGUSERINFO = "MissingParameter.MissingUserInfo"
//  OPERATIONDENIED_NOWHITELIST = "OperationDenied.NoWhitelist"
func (c *Client) QueryForbidPlayChannelListWithContext(ctx context.Context, request *QueryForbidPlayChannelListRequest) (response *QueryForbidPlayChannelListResponse, err error) {
    if request == nil {
        request = NewQueryForbidPlayChannelListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "QueryForbidPlayChannelList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryForbidPlayChannelList require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryForbidPlayChannelListResponse()
    err = c.Send(request, response)
    return
}

func NewRefreshDeviceChannelRequest() (request *RefreshDeviceChannelRequest) {
    request = &RefreshDeviceChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "RefreshDeviceChannel")
    
    
    return
}

func NewRefreshDeviceChannelResponse() (response *RefreshDeviceChannelResponse) {
    response = &RefreshDeviceChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RefreshDeviceChannel
// This API is used to synchronize channels of national standard devices. After the API is called, a request is triggered to get the channel list from devices. Newly-added channels are stored, and channels deleted on devices need to be deleted manually and are not automatically deleted in the backend.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DEVICERESPONSETIMEOUT = "FailedOperation.DeviceResponseTimeOut"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_DEVDISABLE = "ResourceUnavailable.DevDisable"
//  RESOURCEUNAVAILABLE_DEVNOREGISTER = "ResourceUnavailable.DevNoRegister"
//  RESOURCEUNAVAILABLE_DEVOFFLINE = "ResourceUnavailable.DevOffline"
func (c *Client) RefreshDeviceChannel(request *RefreshDeviceChannelRequest) (response *RefreshDeviceChannelResponse, err error) {
    return c.RefreshDeviceChannelWithContext(context.Background(), request)
}

// RefreshDeviceChannel
// This API is used to synchronize channels of national standard devices. After the API is called, a request is triggered to get the channel list from devices. Newly-added channels are stored, and channels deleted on devices need to be deleted manually and are not automatically deleted in the backend.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DEVICERESPONSETIMEOUT = "FailedOperation.DeviceResponseTimeOut"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_DEVDISABLE = "ResourceUnavailable.DevDisable"
//  RESOURCEUNAVAILABLE_DEVNOREGISTER = "ResourceUnavailable.DevNoRegister"
//  RESOURCEUNAVAILABLE_DEVOFFLINE = "ResourceUnavailable.DevOffline"
func (c *Client) RefreshDeviceChannelWithContext(ctx context.Context, request *RefreshDeviceChannelRequest) (response *RefreshDeviceChannelResponse, err error) {
    if request == nil {
        request = NewRefreshDeviceChannelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "RefreshDeviceChannel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RefreshDeviceChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewRefreshDeviceChannelResponse()
    err = c.Send(request, response)
    return
}

func NewSetForbidPlayChannelsRequest() (request *SetForbidPlayChannelsRequest) {
    request = &SetForbidPlayChannelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "SetForbidPlayChannels")
    
    
    return
}

func NewSetForbidPlayChannelsResponse() (response *SetForbidPlayChannelsResponse) {
    response = &SetForbidPlayChannelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetForbidPlayChannels
// Forbid the live preview of video channels for root accounts and sub-accounts.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTHAVESUBUSER = "FailedOperation.NotHaveSubUser"
//  INTERNALERROR_SERVERINTERNALERROR = "InternalError.ServerInternalError"
//  INVALIDPARAMETER_INVALIDCHANNELS = "InvalidParameter.InvalidChannels"
//  INVALIDPARAMETERVALUE_CHANNELSEXCEEDSRANGE = "InvalidParameterValue.ChannelsExceedsRange"
//  INVALIDPARAMETERVALUE_ILLEGALCHANNELID = "InvalidParameterValue.IllegalChannelId"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_INVALIDENALBEVALUE = "InvalidParameterValue.InvalidEnalbeValue"
//  INVALIDPARAMETERVALUE_INVALIDUSERPARAM = "InvalidParameterValue.InvalidUserParam"
//  MISSINGPARAMETER_MISSINGMUSTPARAMETER = "MissingParameter.MissingMustParameter"
//  MISSINGPARAMETER_MISSINGUSERINFO = "MissingParameter.MissingUserInfo"
//  OPERATIONDENIED_NOWHITELIST = "OperationDenied.NoWhitelist"
//  OPERATIONDENIED_PLAYAUTHNOTENABLED = "OperationDenied.PlayAuthNotEnabled"
func (c *Client) SetForbidPlayChannels(request *SetForbidPlayChannelsRequest) (response *SetForbidPlayChannelsResponse, err error) {
    return c.SetForbidPlayChannelsWithContext(context.Background(), request)
}

// SetForbidPlayChannels
// Forbid the live preview of video channels for root accounts and sub-accounts.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTHAVESUBUSER = "FailedOperation.NotHaveSubUser"
//  INTERNALERROR_SERVERINTERNALERROR = "InternalError.ServerInternalError"
//  INVALIDPARAMETER_INVALIDCHANNELS = "InvalidParameter.InvalidChannels"
//  INVALIDPARAMETERVALUE_CHANNELSEXCEEDSRANGE = "InvalidParameterValue.ChannelsExceedsRange"
//  INVALIDPARAMETERVALUE_ILLEGALCHANNELID = "InvalidParameterValue.IllegalChannelId"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_INVALIDENALBEVALUE = "InvalidParameterValue.InvalidEnalbeValue"
//  INVALIDPARAMETERVALUE_INVALIDUSERPARAM = "InvalidParameterValue.InvalidUserParam"
//  MISSINGPARAMETER_MISSINGMUSTPARAMETER = "MissingParameter.MissingMustParameter"
//  MISSINGPARAMETER_MISSINGUSERINFO = "MissingParameter.MissingUserInfo"
//  OPERATIONDENIED_NOWHITELIST = "OperationDenied.NoWhitelist"
//  OPERATIONDENIED_PLAYAUTHNOTENABLED = "OperationDenied.PlayAuthNotEnabled"
func (c *Client) SetForbidPlayChannelsWithContext(ctx context.Context, request *SetForbidPlayChannelsRequest) (response *SetForbidPlayChannelsResponse, err error) {
    if request == nil {
        request = NewSetForbidPlayChannelsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "SetForbidPlayChannels")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetForbidPlayChannels require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetForbidPlayChannelsResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDeviceOrganizationRequest() (request *UpdateDeviceOrganizationRequest) {
    request = &UpdateDeviceOrganizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "UpdateDeviceOrganization")
    
    
    return
}

func NewUpdateDeviceOrganizationResponse() (response *UpdateDeviceOrganizationResponse) {
    response = &UpdateDeviceOrganizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateDeviceOrganization
// Used to change the organization of devices in batches.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateDeviceOrganization(request *UpdateDeviceOrganizationRequest) (response *UpdateDeviceOrganizationResponse, err error) {
    return c.UpdateDeviceOrganizationWithContext(context.Background(), request)
}

// UpdateDeviceOrganization
// Used to change the organization of devices in batches.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateDeviceOrganizationWithContext(ctx context.Context, request *UpdateDeviceOrganizationRequest) (response *UpdateDeviceOrganizationResponse, err error) {
    if request == nil {
        request = NewUpdateDeviceOrganizationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "UpdateDeviceOrganization")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateDeviceOrganization require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateDeviceOrganizationResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDeviceStatusRequest() (request *UpdateDeviceStatusRequest) {
    request = &UpdateDeviceStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "UpdateDeviceStatus")
    
    
    return
}

func NewUpdateDeviceStatusResponse() (response *UpdateDeviceStatusResponse) {
    response = &UpdateDeviceStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateDeviceStatus
// Used to enable or disable a device. After it is disabled, device registration is denied.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDDEVICESTATUS = "InvalidParameterValue.InvalidDeviceStatus"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateDeviceStatus(request *UpdateDeviceStatusRequest) (response *UpdateDeviceStatusResponse, err error) {
    return c.UpdateDeviceStatusWithContext(context.Background(), request)
}

// UpdateDeviceStatus
// Used to enable or disable a device. After it is disabled, device registration is denied.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDDEVICESTATUS = "InvalidParameterValue.InvalidDeviceStatus"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateDeviceStatusWithContext(ctx context.Context, request *UpdateDeviceStatusRequest) (response *UpdateDeviceStatusResponse, err error) {
    if request == nil {
        request = NewUpdateDeviceStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "UpdateDeviceStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateDeviceStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateDeviceStatusResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateOrganizationRequest() (request *UpdateOrganizationRequest) {
    request = &UpdateOrganizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "UpdateOrganization")
    
    
    return
}

func NewUpdateOrganizationResponse() (response *UpdateOrganizationResponse) {
    response = &UpdateOrganizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateOrganization
// Used to modify an organization.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDORGNAME = "InvalidParameterValue.InvalidOrgName"
//  INVALIDPARAMETERVALUE_ORGNAMEREPEAT = "InvalidParameterValue.OrgNameRepeat"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateOrganization(request *UpdateOrganizationRequest) (response *UpdateOrganizationResponse, err error) {
    return c.UpdateOrganizationWithContext(context.Background(), request)
}

// UpdateOrganization
// Used to modify an organization.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDORGNAME = "InvalidParameterValue.InvalidOrgName"
//  INVALIDPARAMETERVALUE_ORGNAMEREPEAT = "InvalidParameterValue.OrgNameRepeat"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateOrganizationWithContext(ctx context.Context, request *UpdateOrganizationRequest) (response *UpdateOrganizationResponse, err error) {
    if request == nil {
        request = NewUpdateOrganizationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "UpdateOrganization")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateOrganization require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateOrganizationResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateRecordBackupPlanRequest() (request *UpdateRecordBackupPlanRequest) {
    request = &UpdateRecordBackupPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "UpdateRecordBackupPlan")
    
    
    return
}

func NewUpdateRecordBackupPlanResponse() (response *UpdateRecordBackupPlanResponse) {
    response = &UpdateRecordBackupPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateRecordBackupPlan
// Used to modify the recording cloud migration plan.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDORGANIZATIONID = "InvalidParameterValue.InvalidOrganizationId"
//  INVALIDPARAMETERVALUE_INVALIDPLANID = "InvalidParameterValue.InvalidPlanId"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  RESOURCEINUSE_PLANNAMEREPEAT = "ResourceInUse.PlanNameRepeat"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateRecordBackupPlan(request *UpdateRecordBackupPlanRequest) (response *UpdateRecordBackupPlanResponse, err error) {
    return c.UpdateRecordBackupPlanWithContext(context.Background(), request)
}

// UpdateRecordBackupPlan
// Used to modify the recording cloud migration plan.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDORGANIZATIONID = "InvalidParameterValue.InvalidOrganizationId"
//  INVALIDPARAMETERVALUE_INVALIDPLANID = "InvalidParameterValue.InvalidPlanId"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  RESOURCEINUSE_PLANNAMEREPEAT = "ResourceInUse.PlanNameRepeat"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateRecordBackupPlanWithContext(ctx context.Context, request *UpdateRecordBackupPlanRequest) (response *UpdateRecordBackupPlanResponse, err error) {
    if request == nil {
        request = NewUpdateRecordBackupPlanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "UpdateRecordBackupPlan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateRecordBackupPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateRecordBackupPlanResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateRecordBackupTemplateRequest() (request *UpdateRecordBackupTemplateRequest) {
    request = &UpdateRecordBackupTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "UpdateRecordBackupTemplate")
    
    
    return
}

func NewUpdateRecordBackupTemplateResponse() (response *UpdateRecordBackupTemplateResponse) {
    response = &UpdateRecordBackupTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateRecordBackupTemplate
// Used to modify a video recording cloud template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_BAKTIMENOTENOUGH = "InvalidParameterValue.BakTimeNotEnough"
//  INVALIDPARAMETERVALUE_DATAOUTOFRANGE = "InvalidParameterValue.DataOutOfRange"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  INVALIDPARAMETERVALUE_INVALIDTIMEFORMAT = "InvalidParameterValue.InvalidTimeFormat"
//  INVALIDPARAMETERVALUE_STARTTIMEGREATERTHANOREQUALENDTIME = "InvalidParameterValue.StartTimeGreaterThanOrEqualEndTime"
//  INVALIDPARAMETERVALUE_TIMELESSTHANTENMINUTES = "InvalidParameterValue.TimeLessThanTenMinutes"
//  INVALIDPARAMETERVALUE_UNSUPPORTSCALE = "InvalidParameterValue.UnSupportScale"
//  RESOURCEINUSE_TEMPLATENAMEREPEAT = "ResourceInUse.TemplateNameRepeat"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateRecordBackupTemplate(request *UpdateRecordBackupTemplateRequest) (response *UpdateRecordBackupTemplateResponse, err error) {
    return c.UpdateRecordBackupTemplateWithContext(context.Background(), request)
}

// UpdateRecordBackupTemplate
// Used to modify a video recording cloud template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_BAKTIMENOTENOUGH = "InvalidParameterValue.BakTimeNotEnough"
//  INVALIDPARAMETERVALUE_DATAOUTOFRANGE = "InvalidParameterValue.DataOutOfRange"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  INVALIDPARAMETERVALUE_INVALIDTIMEFORMAT = "InvalidParameterValue.InvalidTimeFormat"
//  INVALIDPARAMETERVALUE_STARTTIMEGREATERTHANOREQUALENDTIME = "InvalidParameterValue.StartTimeGreaterThanOrEqualEndTime"
//  INVALIDPARAMETERVALUE_TIMELESSTHANTENMINUTES = "InvalidParameterValue.TimeLessThanTenMinutes"
//  INVALIDPARAMETERVALUE_UNSUPPORTSCALE = "InvalidParameterValue.UnSupportScale"
//  RESOURCEINUSE_TEMPLATENAMEREPEAT = "ResourceInUse.TemplateNameRepeat"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateRecordBackupTemplateWithContext(ctx context.Context, request *UpdateRecordBackupTemplateRequest) (response *UpdateRecordBackupTemplateResponse, err error) {
    if request == nil {
        request = NewUpdateRecordBackupTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "UpdateRecordBackupTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateRecordBackupTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateRecordBackupTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateRecordPlanRequest() (request *UpdateRecordPlanRequest) {
    request = &UpdateRecordPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "UpdateRecordPlan")
    
    
    return
}

func NewUpdateRecordPlanResponse() (response *UpdateRecordPlanResponse) {
    response = &UpdateRecordPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateRecordPlan
// This API is used to modify a real-time cloud migration plan.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETER_INVALIDCHANNELS = "InvalidParameter.InvalidChannels"
//  INVALIDPARAMETER_INVALIDLIFERULEPARAM = "InvalidParameter.InvalidLifeRuleParam"
//  INVALIDPARAMETER_INVALIDORGANIZATIONPARAM = "InvalidParameter.InvalidOrganizationParam"
//  INVALIDPARAMETERVALUE_EMPTYCHANNELID = "InvalidParameterValue.EmptyChannelId"
//  INVALIDPARAMETERVALUE_EMPTYDEVICEID = "InvalidParameterValue.EmptyDeviceId"
//  INVALIDPARAMETERVALUE_EMPTYNAME = "InvalidParameterValue.EmptyName"
//  INVALIDPARAMETERVALUE_EMPTYORGANIZATIONID = "InvalidParameterValue.EmptyOrganizationId"
//  INVALIDPARAMETERVALUE_ILLEGALCHANNELID = "InvalidParameterValue.IllegalChannelId"
//  INVALIDPARAMETERVALUE_ILLEGALDESCRIBE = "InvalidParameterValue.IllegalDescribe"
//  INVALIDPARAMETERVALUE_ILLEGALDEVICEID = "InvalidParameterValue.IllegalDeviceId"
//  INVALIDPARAMETERVALUE_ILLEGALNAME = "InvalidParameterValue.IllegalName"
//  INVALIDPARAMETERVALUE_ILLEGALORGANIZATIONID = "InvalidParameterValue.IllegalOrganizationId"
//  INVALIDPARAMETERVALUE_ILLEGALSTREAMTYPE = "InvalidParameterValue.IllegalStreamType"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELIDORDEVICEID = "InvalidParameterValue.InvalidChannelIdOrDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDEXPIRATIONRANGE = "InvalidParameterValue.InvalidExpirationRange"
//  INVALIDPARAMETERVALUE_INVALIDPLANID = "InvalidParameterValue.InvalidPlanId"
//  INVALIDPARAMETERVALUE_INVALIDREPAIRMODE = "InvalidParameterValue.InvalidRepairMode"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  INVALIDPARAMETERVALUE_INVALIDTRANSITIONRANGE = "InvalidParameterValue.InvalidTransitionRange"
//  INVALIDPARAMETERVALUE_LIMITREPAIRTRANSITION = "InvalidParameterValue.LimitRepairTransition"
//  INVALIDPARAMETERVALUE_PLANCHANNELSEXCEEDSRANGE = "InvalidParameterValue.PlanChannelsExceedsRange"
//  INVALIDPARAMETERVALUE_TOOLONGDESCRIBE = "InvalidParameterValue.TooLongDescribe"
//  INVALIDPARAMETERVALUE_TOOLONGNAME = "InvalidParameterValue.TooLongName"
//  INVALIDPARAMETERVALUE_TOOLONGSTREAMTYPE = "InvalidParameterValue.TooLongStreamType"
//  MISSINGPARAMETER_MISSINGLIFERULEPARAM = "MissingParameter.MissingLifeRuleParam"
//  MISSINGPARAMETER_MISSINGUPDATEDINFO = "MissingParameter.MissingUpdatedInfo"
//  RESOURCEINUSE_CHANNELREPEATADD = "ResourceInUse.ChannelRepeatAdd"
//  RESOURCEINUSE_PLANDELETING = "ResourceInUse.PlanDeleting"
//  RESOURCEINUSE_PLANNAMEREPEAT = "ResourceInUse.PlanNameRepeat"
//  RESOURCENOTFOUND_ORGANIZATIONIDNOTEXIST = "ResourceNotFound.OrganizationIdNotExist"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) UpdateRecordPlan(request *UpdateRecordPlanRequest) (response *UpdateRecordPlanResponse, err error) {
    return c.UpdateRecordPlanWithContext(context.Background(), request)
}

// UpdateRecordPlan
// This API is used to modify a real-time cloud migration plan.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETER_INVALIDCHANNELS = "InvalidParameter.InvalidChannels"
//  INVALIDPARAMETER_INVALIDLIFERULEPARAM = "InvalidParameter.InvalidLifeRuleParam"
//  INVALIDPARAMETER_INVALIDORGANIZATIONPARAM = "InvalidParameter.InvalidOrganizationParam"
//  INVALIDPARAMETERVALUE_EMPTYCHANNELID = "InvalidParameterValue.EmptyChannelId"
//  INVALIDPARAMETERVALUE_EMPTYDEVICEID = "InvalidParameterValue.EmptyDeviceId"
//  INVALIDPARAMETERVALUE_EMPTYNAME = "InvalidParameterValue.EmptyName"
//  INVALIDPARAMETERVALUE_EMPTYORGANIZATIONID = "InvalidParameterValue.EmptyOrganizationId"
//  INVALIDPARAMETERVALUE_ILLEGALCHANNELID = "InvalidParameterValue.IllegalChannelId"
//  INVALIDPARAMETERVALUE_ILLEGALDESCRIBE = "InvalidParameterValue.IllegalDescribe"
//  INVALIDPARAMETERVALUE_ILLEGALDEVICEID = "InvalidParameterValue.IllegalDeviceId"
//  INVALIDPARAMETERVALUE_ILLEGALNAME = "InvalidParameterValue.IllegalName"
//  INVALIDPARAMETERVALUE_ILLEGALORGANIZATIONID = "InvalidParameterValue.IllegalOrganizationId"
//  INVALIDPARAMETERVALUE_ILLEGALSTREAMTYPE = "InvalidParameterValue.IllegalStreamType"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELID = "InvalidParameterValue.InvalidChannelId"
//  INVALIDPARAMETERVALUE_INVALIDCHANNELIDORDEVICEID = "InvalidParameterValue.InvalidChannelIdOrDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDEXPIRATIONRANGE = "InvalidParameterValue.InvalidExpirationRange"
//  INVALIDPARAMETERVALUE_INVALIDPLANID = "InvalidParameterValue.InvalidPlanId"
//  INVALIDPARAMETERVALUE_INVALIDREPAIRMODE = "InvalidParameterValue.InvalidRepairMode"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  INVALIDPARAMETERVALUE_INVALIDTRANSITIONRANGE = "InvalidParameterValue.InvalidTransitionRange"
//  INVALIDPARAMETERVALUE_LIMITREPAIRTRANSITION = "InvalidParameterValue.LimitRepairTransition"
//  INVALIDPARAMETERVALUE_PLANCHANNELSEXCEEDSRANGE = "InvalidParameterValue.PlanChannelsExceedsRange"
//  INVALIDPARAMETERVALUE_TOOLONGDESCRIBE = "InvalidParameterValue.TooLongDescribe"
//  INVALIDPARAMETERVALUE_TOOLONGNAME = "InvalidParameterValue.TooLongName"
//  INVALIDPARAMETERVALUE_TOOLONGSTREAMTYPE = "InvalidParameterValue.TooLongStreamType"
//  MISSINGPARAMETER_MISSINGLIFERULEPARAM = "MissingParameter.MissingLifeRuleParam"
//  MISSINGPARAMETER_MISSINGUPDATEDINFO = "MissingParameter.MissingUpdatedInfo"
//  RESOURCEINUSE_CHANNELREPEATADD = "ResourceInUse.ChannelRepeatAdd"
//  RESOURCEINUSE_PLANDELETING = "ResourceInUse.PlanDeleting"
//  RESOURCEINUSE_PLANNAMEREPEAT = "ResourceInUse.PlanNameRepeat"
//  RESOURCENOTFOUND_ORGANIZATIONIDNOTEXIST = "ResourceNotFound.OrganizationIdNotExist"
//  RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) UpdateRecordPlanWithContext(ctx context.Context, request *UpdateRecordPlanRequest) (response *UpdateRecordPlanResponse, err error) {
    if request == nil {
        request = NewUpdateRecordPlanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "UpdateRecordPlan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateRecordPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateRecordPlanResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateRecordTemplateRequest() (request *UpdateRecordTemplateRequest) {
    request = &UpdateRecordTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "UpdateRecordTemplate")
    
    
    return
}

func NewUpdateRecordTemplateResponse() (response *UpdateRecordTemplateResponse) {
    response = &UpdateRecordTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateRecordTemplate
// Used to modify a real-time cloud template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETER_INVALIDTIMESECTION = "InvalidParameter.InvalidTimeSection"
//  INVALIDPARAMETERVALUE_EMPTYNAME = "InvalidParameterValue.EmptyName"
//  INVALIDPARAMETERVALUE_ILLEGALNAME = "InvalidParameterValue.IllegalName"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  INVALIDPARAMETERVALUE_INVALIDTIMESECTIONVALUE = "InvalidParameterValue.InvalidTimeSectionValue"
//  INVALIDPARAMETERVALUE_STARTTIMEGREATERTHANOREQUALENDTIME = "InvalidParameterValue.StartTimeGreaterThanOrEqualEndTime"
//  INVALIDPARAMETERVALUE_TIMELESSTHANTENMINUTES = "InvalidParameterValue.TimeLessThanTenMinutes"
//  INVALIDPARAMETERVALUE_TOOLONGNAME = "InvalidParameterValue.TooLongName"
//  MISSINGPARAMETER_EMPTYTIMESECTION = "MissingParameter.EmptyTimeSection"
//  MISSINGPARAMETER_MISSINGUPDATEDINFO = "MissingParameter.MissingUpdatedInfo"
//  RESOURCEINUSE_TEMPLATENAMEREPEAT = "ResourceInUse.TemplateNameRepeat"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) UpdateRecordTemplate(request *UpdateRecordTemplateRequest) (response *UpdateRecordTemplateResponse, err error) {
    return c.UpdateRecordTemplateWithContext(context.Background(), request)
}

// UpdateRecordTemplate
// Used to modify a real-time cloud template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDBODYFORMAT = "InvalidParameter.InvalidBodyFormat"
//  INVALIDPARAMETER_INVALIDTIMESECTION = "InvalidParameter.InvalidTimeSection"
//  INVALIDPARAMETERVALUE_EMPTYNAME = "InvalidParameterValue.EmptyName"
//  INVALIDPARAMETERVALUE_ILLEGALNAME = "InvalidParameterValue.IllegalName"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEID = "InvalidParameterValue.InvalidTemplateId"
//  INVALIDPARAMETERVALUE_INVALIDTIMESECTIONVALUE = "InvalidParameterValue.InvalidTimeSectionValue"
//  INVALIDPARAMETERVALUE_STARTTIMEGREATERTHANOREQUALENDTIME = "InvalidParameterValue.StartTimeGreaterThanOrEqualEndTime"
//  INVALIDPARAMETERVALUE_TIMELESSTHANTENMINUTES = "InvalidParameterValue.TimeLessThanTenMinutes"
//  INVALIDPARAMETERVALUE_TOOLONGNAME = "InvalidParameterValue.TooLongName"
//  MISSINGPARAMETER_EMPTYTIMESECTION = "MissingParameter.EmptyTimeSection"
//  MISSINGPARAMETER_MISSINGUPDATEDINFO = "MissingParameter.MissingUpdatedInfo"
//  RESOURCEINUSE_TEMPLATENAMEREPEAT = "ResourceInUse.TemplateNameRepeat"
//  RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"
func (c *Client) UpdateRecordTemplateWithContext(ctx context.Context, request *UpdateRecordTemplateRequest) (response *UpdateRecordTemplateResponse, err error) {
    if request == nil {
        request = NewUpdateRecordTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "UpdateRecordTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateRecordTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateRecordTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateUserDeviceRequest() (request *UpdateUserDeviceRequest) {
    request = &UpdateUserDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iss", APIVersion, "UpdateUserDevice")
    
    
    return
}

func NewUpdateUserDeviceResponse() (response *UpdateUserDeviceResponse) {
    response = &UpdateUserDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateUserDevice
// This API is used to modify device configuration information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEPASSWORDLENGTH = "InvalidParameterValue.InvalidDevicePasswordLength"
//  INVALIDPARAMETERVALUE_INVALIDGATEWAYPROTOCOLTYPE = "InvalidParameterValue.InvalidGatewayProtocolType"
//  INVALIDPARAMETERVALUE_INVALIDIPV4 = "InvalidParameterValue.InvalidIpv4"
//  INVALIDPARAMETERVALUE_INVALIDNAME = "InvalidParameterValue.InvalidName"
//  INVALIDPARAMETERVALUE_INVALIDUSERNAME = "InvalidParameterValue.InvalidUsername"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDSTREAMPROTOCOL = "InvalidParameterValue.UnsupportedStreamProtocol"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateUserDevice(request *UpdateUserDeviceRequest) (response *UpdateUserDeviceResponse, err error) {
    return c.UpdateUserDeviceWithContext(context.Background(), request)
}

// UpdateUserDevice
// This API is used to modify device configuration information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEID = "InvalidParameterValue.InvalidDeviceId"
//  INVALIDPARAMETERVALUE_INVALIDDEVICEPASSWORDLENGTH = "InvalidParameterValue.InvalidDevicePasswordLength"
//  INVALIDPARAMETERVALUE_INVALIDGATEWAYPROTOCOLTYPE = "InvalidParameterValue.InvalidGatewayProtocolType"
//  INVALIDPARAMETERVALUE_INVALIDIPV4 = "InvalidParameterValue.InvalidIpv4"
//  INVALIDPARAMETERVALUE_INVALIDNAME = "InvalidParameterValue.InvalidName"
//  INVALIDPARAMETERVALUE_INVALIDUSERNAME = "InvalidParameterValue.InvalidUsername"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDACCESSTYPE = "InvalidParameterValue.UnSupportedAccessType"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDSTREAMPROTOCOL = "InvalidParameterValue.UnsupportedStreamProtocol"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateUserDeviceWithContext(ctx context.Context, request *UpdateUserDeviceRequest) (response *UpdateUserDeviceResponse, err error) {
    if request == nil {
        request = NewUpdateUserDeviceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iss", APIVersion, "UpdateUserDevice")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateUserDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateUserDeviceResponse()
    err = c.Send(request, response)
    return
}
