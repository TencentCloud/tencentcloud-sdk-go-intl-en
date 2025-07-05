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

package v20170312

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2017-03-12"

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


func NewAttachInstancesRequest() (request *AttachInstancesRequest) {
    request = &AttachInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "AttachInstances")
    
    
    return
}

func NewAttachInstancesResponse() (response *AttachInstancesResponse) {
    response = &AttachInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AttachInstances
// This API is used to add existing instances to the compute environment.
//
// Requirements: <br/>
//
// 1. The instance is not in the batch compute system.<br/>
//
// 2. The instance is in “Running” status.<br/>
//
// 3. Spot instances are not supported.<b/>
//
// 
//
// For instances added to the compute environment, their UserData will be reset, and the operating systems will be reinstalled.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLCVM = "InternalError.CallCvm"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  INVALIDPARAMETERVALUE_INSTANCEIDDUPLICATED = "InvalidParameterValue.InstanceIdDuplicated"
//  RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INSTANCESNOTALLOWTOATTACH = "UnsupportedOperation.InstancesNotAllowToAttach"
func (c *Client) AttachInstances(request *AttachInstancesRequest) (response *AttachInstancesResponse, err error) {
    return c.AttachInstancesWithContext(context.Background(), request)
}

// AttachInstances
// This API is used to add existing instances to the compute environment.
//
// Requirements: <br/>
//
// 1. The instance is not in the batch compute system.<br/>
//
// 2. The instance is in “Running” status.<br/>
//
// 3. Spot instances are not supported.<b/>
//
// 
//
// For instances added to the compute environment, their UserData will be reset, and the operating systems will be reinstalled.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLCVM = "InternalError.CallCvm"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  INVALIDPARAMETERVALUE_INSTANCEIDDUPLICATED = "InvalidParameterValue.InstanceIdDuplicated"
//  RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INSTANCESNOTALLOWTOATTACH = "UnsupportedOperation.InstancesNotAllowToAttach"
func (c *Client) AttachInstancesWithContext(ctx context.Context, request *AttachInstancesRequest) (response *AttachInstancesResponse, err error) {
    if request == nil {
        request = NewAttachInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AttachInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewAttachInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateComputeEnvRequest() (request *CreateComputeEnvRequest) {
    request = &CreateComputeEnvRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "CreateComputeEnv")
    
    
    return
}

func NewCreateComputeEnvResponse() (response *CreateComputeEnvResponse) {
    response = &CreateComputeEnvResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateComputeEnv
// This API is used to create a compute environment.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLCAM = "InternalError.CallCam"
//  INTERNALERROR_CALLCVM = "InternalError.CallCvm"
//  INTERNALERROR_CALLTAGAPI = "InternalError.CallTagAPI"
//  INVALIDPARAMETER_CVMPARAMETERS = "InvalidParameter.CvmParameters"
//  INVALIDPARAMETER_ENVDESCRIPTIONTOOLONG = "InvalidParameter.EnvDescriptionTooLong"
//  INVALIDPARAMETER_ENVNAMETOOLONG = "InvalidParameter.EnvNameTooLong"
//  INVALIDPARAMETER_IMAGEIDMALFORMED = "InvalidParameter.ImageIdMalformed"
//  INVALIDPARAMETER_INVALIDPARAMETERCOMBINATION = "InvalidParameter.InvalidParameterCombination"
//  INVALIDPARAMETER_NOTIFICATIONEVENTNAMEDUPLICATE = "InvalidParameter.NotificationEventNameDuplicate"
//  INVALIDPARAMETER_NOTIFICATIONTOPICNAME = "InvalidParameter.NotificationTopicName"
//  INVALIDPARAMETER_NOTIFICATIONTOPICNAMETOOLONG = "InvalidParameter.NotificationTopicNameTooLong"
//  INVALIDPARAMETERVALUE_COMPUTEENV = "InvalidParameterValue.ComputeEnv"
//  INVALIDPARAMETERVALUE_INSTANCETYPE = "InvalidParameterValue.InstanceType"
//  INVALIDPARAMETERVALUE_INVALIDDATATYPEANY = "InvalidParameterValue.InvalidDataTypeAny"
//  INVALIDPARAMETERVALUE_INVALIDZONEMISMATCHREGION = "InvalidParameterValue.InvalidZoneMismatchRegion"
//  INVALIDPARAMETERVALUE_LOCALPATH = "InvalidParameterValue.LocalPath"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_REMOTESTORAGEPATH = "InvalidParameterValue.RemoteStoragePath"
//  INVALIDPARAMETERVALUE_REMOTESTORAGESCHEMETYPE = "InvalidParameterValue.RemoteStorageSchemeType"
//  INVALIDPARAMETERVALUE_UNAVAILABLEZONE = "InvalidParameterValue.UnavailableZone"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  LIMITEXCEEDED_COMPUTEENVQUOTA = "LimitExceeded.ComputeEnvQuota"
//  LIMITEXCEEDED_CPUQUOTA = "LimitExceeded.CpuQuota"
//  UNAUTHORIZEDOPERATION_USERNOTALLOWEDTOUSEBATCH = "UnauthorizedOperation.UserNotAllowedToUseBatch"
func (c *Client) CreateComputeEnv(request *CreateComputeEnvRequest) (response *CreateComputeEnvResponse, err error) {
    return c.CreateComputeEnvWithContext(context.Background(), request)
}

// CreateComputeEnv
// This API is used to create a compute environment.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLCAM = "InternalError.CallCam"
//  INTERNALERROR_CALLCVM = "InternalError.CallCvm"
//  INTERNALERROR_CALLTAGAPI = "InternalError.CallTagAPI"
//  INVALIDPARAMETER_CVMPARAMETERS = "InvalidParameter.CvmParameters"
//  INVALIDPARAMETER_ENVDESCRIPTIONTOOLONG = "InvalidParameter.EnvDescriptionTooLong"
//  INVALIDPARAMETER_ENVNAMETOOLONG = "InvalidParameter.EnvNameTooLong"
//  INVALIDPARAMETER_IMAGEIDMALFORMED = "InvalidParameter.ImageIdMalformed"
//  INVALIDPARAMETER_INVALIDPARAMETERCOMBINATION = "InvalidParameter.InvalidParameterCombination"
//  INVALIDPARAMETER_NOTIFICATIONEVENTNAMEDUPLICATE = "InvalidParameter.NotificationEventNameDuplicate"
//  INVALIDPARAMETER_NOTIFICATIONTOPICNAME = "InvalidParameter.NotificationTopicName"
//  INVALIDPARAMETER_NOTIFICATIONTOPICNAMETOOLONG = "InvalidParameter.NotificationTopicNameTooLong"
//  INVALIDPARAMETERVALUE_COMPUTEENV = "InvalidParameterValue.ComputeEnv"
//  INVALIDPARAMETERVALUE_INSTANCETYPE = "InvalidParameterValue.InstanceType"
//  INVALIDPARAMETERVALUE_INVALIDDATATYPEANY = "InvalidParameterValue.InvalidDataTypeAny"
//  INVALIDPARAMETERVALUE_INVALIDZONEMISMATCHREGION = "InvalidParameterValue.InvalidZoneMismatchRegion"
//  INVALIDPARAMETERVALUE_LOCALPATH = "InvalidParameterValue.LocalPath"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_REMOTESTORAGEPATH = "InvalidParameterValue.RemoteStoragePath"
//  INVALIDPARAMETERVALUE_REMOTESTORAGESCHEMETYPE = "InvalidParameterValue.RemoteStorageSchemeType"
//  INVALIDPARAMETERVALUE_UNAVAILABLEZONE = "InvalidParameterValue.UnavailableZone"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  LIMITEXCEEDED_COMPUTEENVQUOTA = "LimitExceeded.ComputeEnvQuota"
//  LIMITEXCEEDED_CPUQUOTA = "LimitExceeded.CpuQuota"
//  UNAUTHORIZEDOPERATION_USERNOTALLOWEDTOUSEBATCH = "UnauthorizedOperation.UserNotAllowedToUseBatch"
func (c *Client) CreateComputeEnvWithContext(ctx context.Context, request *CreateComputeEnvRequest) (response *CreateComputeEnvResponse, err error) {
    if request == nil {
        request = NewCreateComputeEnvRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateComputeEnv require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateComputeEnvResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTaskTemplateRequest() (request *CreateTaskTemplateRequest) {
    request = &CreateTaskTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "CreateTaskTemplate")
    
    
    return
}

func NewCreateTaskTemplateResponse() (response *CreateTaskTemplateResponse) {
    response = &CreateTaskTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTaskTemplate
// This API is used to create a task template.
//
// error code that may be returned:
//  ALLOWEDONEATTRIBUTEINENVIDANDCOMPUTEENV = "AllowedOneAttributeInEnvIdAndComputeEnv"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLTAGAPI = "InternalError.CallTagAPI"
//  INVALIDPARAMETER_CVMPARAMETERS = "InvalidParameter.CvmParameters"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  INVALIDPARAMETER_IMAGEIDMALFORMED = "InvalidParameter.ImageIdMalformed"
//  INVALIDPARAMETER_TASKNAME = "InvalidParameter.TaskName"
//  INVALIDPARAMETER_TASKNAMETOOLONG = "InvalidParameter.TaskNameTooLong"
//  INVALIDPARAMETER_TASKTEMPLATENAME = "InvalidParameter.TaskTemplateName"
//  INVALIDPARAMETER_TASKTEMPLATENAMETOOLONG = "InvalidParameter.TaskTemplateNameTooLong"
//  INVALIDPARAMETERVALUE_COMPUTEENV = "InvalidParameterValue.ComputeEnv"
//  INVALIDPARAMETERVALUE_DEPENDENCENOTFOUNDTASKNAME = "InvalidParameterValue.DependenceNotFoundTaskName"
//  INVALIDPARAMETERVALUE_DEPENDENCEUNFEASIBLE = "InvalidParameterValue.DependenceUnfeasible"
//  INVALIDPARAMETERVALUE_INVALIDDATATYPEANY = "InvalidParameterValue.InvalidDataTypeAny"
//  INVALIDPARAMETERVALUE_LOCALPATH = "InvalidParameterValue.LocalPath"
//  INVALIDPARAMETERVALUE_MAXRETRYCOUNT = "InvalidParameterValue.MaxRetryCount"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_REMOTESTORAGEPATH = "InvalidParameterValue.RemoteStoragePath"
//  INVALIDPARAMETERVALUE_REMOTESTORAGESCHEMETYPE = "InvalidParameterValue.RemoteStorageSchemeType"
//  LIMITEXCEEDED_TASKTEMPLATEQUOTA = "LimitExceeded.TaskTemplateQuota"
func (c *Client) CreateTaskTemplate(request *CreateTaskTemplateRequest) (response *CreateTaskTemplateResponse, err error) {
    return c.CreateTaskTemplateWithContext(context.Background(), request)
}

// CreateTaskTemplate
// This API is used to create a task template.
//
// error code that may be returned:
//  ALLOWEDONEATTRIBUTEINENVIDANDCOMPUTEENV = "AllowedOneAttributeInEnvIdAndComputeEnv"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLTAGAPI = "InternalError.CallTagAPI"
//  INVALIDPARAMETER_CVMPARAMETERS = "InvalidParameter.CvmParameters"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  INVALIDPARAMETER_IMAGEIDMALFORMED = "InvalidParameter.ImageIdMalformed"
//  INVALIDPARAMETER_TASKNAME = "InvalidParameter.TaskName"
//  INVALIDPARAMETER_TASKNAMETOOLONG = "InvalidParameter.TaskNameTooLong"
//  INVALIDPARAMETER_TASKTEMPLATENAME = "InvalidParameter.TaskTemplateName"
//  INVALIDPARAMETER_TASKTEMPLATENAMETOOLONG = "InvalidParameter.TaskTemplateNameTooLong"
//  INVALIDPARAMETERVALUE_COMPUTEENV = "InvalidParameterValue.ComputeEnv"
//  INVALIDPARAMETERVALUE_DEPENDENCENOTFOUNDTASKNAME = "InvalidParameterValue.DependenceNotFoundTaskName"
//  INVALIDPARAMETERVALUE_DEPENDENCEUNFEASIBLE = "InvalidParameterValue.DependenceUnfeasible"
//  INVALIDPARAMETERVALUE_INVALIDDATATYPEANY = "InvalidParameterValue.InvalidDataTypeAny"
//  INVALIDPARAMETERVALUE_LOCALPATH = "InvalidParameterValue.LocalPath"
//  INVALIDPARAMETERVALUE_MAXRETRYCOUNT = "InvalidParameterValue.MaxRetryCount"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_REMOTESTORAGEPATH = "InvalidParameterValue.RemoteStoragePath"
//  INVALIDPARAMETERVALUE_REMOTESTORAGESCHEMETYPE = "InvalidParameterValue.RemoteStorageSchemeType"
//  LIMITEXCEEDED_TASKTEMPLATEQUOTA = "LimitExceeded.TaskTemplateQuota"
func (c *Client) CreateTaskTemplateWithContext(ctx context.Context, request *CreateTaskTemplateRequest) (response *CreateTaskTemplateResponse, err error) {
    if request == nil {
        request = NewCreateTaskTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTaskTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTaskTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteComputeEnvRequest() (request *DeleteComputeEnvRequest) {
    request = &DeleteComputeEnvRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DeleteComputeEnv")
    
    
    return
}

func NewDeleteComputeEnvResponse() (response *DeleteComputeEnvResponse) {
    response = &DeleteComputeEnvResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteComputeEnv
// This API is used to delete a compute environment.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ACCEPTOTHERREQUEST = "UnsupportedOperation.AcceptOtherRequest"
//  UNSUPPORTEDOPERATION_COMPUTEENVACCEPTOTHERREQUEST = "UnsupportedOperation.ComputeEnvAcceptOtherRequest"
//  UNSUPPORTEDOPERATION_COMPUTEENVOPERATION = "UnsupportedOperation.ComputeEnvOperation"
//  UNSUPPORTEDOPERATION_COMPUTENODEFORBIDTERMINATE = "UnsupportedOperation.ComputeNodeForbidTerminate"
func (c *Client) DeleteComputeEnv(request *DeleteComputeEnvRequest) (response *DeleteComputeEnvResponse, err error) {
    return c.DeleteComputeEnvWithContext(context.Background(), request)
}

// DeleteComputeEnv
// This API is used to delete a compute environment.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ACCEPTOTHERREQUEST = "UnsupportedOperation.AcceptOtherRequest"
//  UNSUPPORTEDOPERATION_COMPUTEENVACCEPTOTHERREQUEST = "UnsupportedOperation.ComputeEnvAcceptOtherRequest"
//  UNSUPPORTEDOPERATION_COMPUTEENVOPERATION = "UnsupportedOperation.ComputeEnvOperation"
//  UNSUPPORTEDOPERATION_COMPUTENODEFORBIDTERMINATE = "UnsupportedOperation.ComputeNodeForbidTerminate"
func (c *Client) DeleteComputeEnvWithContext(ctx context.Context, request *DeleteComputeEnvRequest) (response *DeleteComputeEnvResponse, err error) {
    if request == nil {
        request = NewDeleteComputeEnvRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteComputeEnv require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteComputeEnvResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteJobRequest() (request *DeleteJobRequest) {
    request = &DeleteJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DeleteJob")
    
    
    return
}

func NewDeleteJobResponse() (response *DeleteJobResponse) {
    response = &DeleteJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteJob
// This API is used to delete a job.
//
// When a job is deleted, all related information is deleted and the job cannot be queried.
//
// To delete a job, the job and all its task instances must be in SUCCEED or FAILED status.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JOBIDMALFORMED = "InvalidParameter.JobIdMalformed"
//  RESOURCEINUSE_JOB = "ResourceInUse.Job"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteJob(request *DeleteJobRequest) (response *DeleteJobResponse, err error) {
    return c.DeleteJobWithContext(context.Background(), request)
}

// DeleteJob
// This API is used to delete a job.
//
// When a job is deleted, all related information is deleted and the job cannot be queried.
//
// To delete a job, the job and all its task instances must be in SUCCEED or FAILED status.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JOBIDMALFORMED = "InvalidParameter.JobIdMalformed"
//  RESOURCEINUSE_JOB = "ResourceInUse.Job"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteJobWithContext(ctx context.Context, request *DeleteJobRequest) (response *DeleteJobResponse, err error) {
    if request == nil {
        request = NewDeleteJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteJobResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTaskTemplatesRequest() (request *DeleteTaskTemplatesRequest) {
    request = &DeleteTaskTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DeleteTaskTemplates")
    
    
    return
}

func NewDeleteTaskTemplatesResponse() (response *DeleteTaskTemplatesResponse) {
    response = &DeleteTaskTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTaskTemplates
// This API is used to delete task template information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_TASKTEMPLATEIDMALFORMED = "InvalidParameter.TaskTemplateIdMalformed"
//  RESOURCENOTFOUND_TASKTEMPLATE = "ResourceNotFound.TaskTemplate"
func (c *Client) DeleteTaskTemplates(request *DeleteTaskTemplatesRequest) (response *DeleteTaskTemplatesResponse, err error) {
    return c.DeleteTaskTemplatesWithContext(context.Background(), request)
}

// DeleteTaskTemplates
// This API is used to delete task template information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_TASKTEMPLATEIDMALFORMED = "InvalidParameter.TaskTemplateIdMalformed"
//  RESOURCENOTFOUND_TASKTEMPLATE = "ResourceNotFound.TaskTemplate"
func (c *Client) DeleteTaskTemplatesWithContext(ctx context.Context, request *DeleteTaskTemplatesRequest) (response *DeleteTaskTemplatesResponse, err error) {
    if request == nil {
        request = NewDeleteTaskTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTaskTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTaskTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAvailableCvmInstanceTypesRequest() (request *DescribeAvailableCvmInstanceTypesRequest) {
    request = &DescribeAvailableCvmInstanceTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DescribeAvailableCvmInstanceTypes")
    
    
    return
}

func NewDescribeAvailableCvmInstanceTypesResponse() (response *DescribeAvailableCvmInstanceTypesResponse) {
    response = &DescribeAvailableCvmInstanceTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAvailableCvmInstanceTypes
// This API is used to view the information of available CVM model configurations.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLCVM = "InternalError.CallCvm"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDZONEMISMATCHREGION = "InvalidParameterValue.InvalidZoneMismatchRegion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
func (c *Client) DescribeAvailableCvmInstanceTypes(request *DescribeAvailableCvmInstanceTypesRequest) (response *DescribeAvailableCvmInstanceTypesResponse, err error) {
    return c.DescribeAvailableCvmInstanceTypesWithContext(context.Background(), request)
}

// DescribeAvailableCvmInstanceTypes
// This API is used to view the information of available CVM model configurations.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLCVM = "InternalError.CallCvm"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDZONEMISMATCHREGION = "InvalidParameterValue.InvalidZoneMismatchRegion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
func (c *Client) DescribeAvailableCvmInstanceTypesWithContext(ctx context.Context, request *DescribeAvailableCvmInstanceTypesRequest) (response *DescribeAvailableCvmInstanceTypesResponse, err error) {
    if request == nil {
        request = NewDescribeAvailableCvmInstanceTypesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAvailableCvmInstanceTypes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAvailableCvmInstanceTypesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComputeEnvRequest() (request *DescribeComputeEnvRequest) {
    request = &DescribeComputeEnvRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DescribeComputeEnv")
    
    
    return
}

func NewDescribeComputeEnvResponse() (response *DescribeComputeEnvResponse) {
    response = &DescribeComputeEnvResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeComputeEnv
// This API is used to query compute environment details.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"
func (c *Client) DescribeComputeEnv(request *DescribeComputeEnvRequest) (response *DescribeComputeEnvResponse, err error) {
    return c.DescribeComputeEnvWithContext(context.Background(), request)
}

// DescribeComputeEnv
// This API is used to query compute environment details.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"
func (c *Client) DescribeComputeEnvWithContext(ctx context.Context, request *DescribeComputeEnvRequest) (response *DescribeComputeEnvResponse, err error) {
    if request == nil {
        request = NewDescribeComputeEnvRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeComputeEnv require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeComputeEnvResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComputeEnvActivitiesRequest() (request *DescribeComputeEnvActivitiesRequest) {
    request = &DescribeComputeEnvActivitiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DescribeComputeEnvActivities")
    
    
    return
}

func NewDescribeComputeEnvActivitiesResponse() (response *DescribeComputeEnvActivitiesResponse) {
    response = &DescribeComputeEnvActivitiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeComputeEnvActivities
// This API is used to query the information of activities in the compute environment.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_COMPUTENODEIDMALFORMED = "InvalidParameter.ComputeNodeIdMalformed"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeComputeEnvActivities(request *DescribeComputeEnvActivitiesRequest) (response *DescribeComputeEnvActivitiesResponse, err error) {
    return c.DescribeComputeEnvActivitiesWithContext(context.Background(), request)
}

// DescribeComputeEnvActivities
// This API is used to query the information of activities in the compute environment.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_COMPUTENODEIDMALFORMED = "InvalidParameter.ComputeNodeIdMalformed"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeComputeEnvActivitiesWithContext(ctx context.Context, request *DescribeComputeEnvActivitiesRequest) (response *DescribeComputeEnvActivitiesResponse, err error) {
    if request == nil {
        request = NewDescribeComputeEnvActivitiesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeComputeEnvActivities require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeComputeEnvActivitiesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComputeEnvCreateInfoRequest() (request *DescribeComputeEnvCreateInfoRequest) {
    request = &DescribeComputeEnvCreateInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DescribeComputeEnvCreateInfo")
    
    
    return
}

func NewDescribeComputeEnvCreateInfoResponse() (response *DescribeComputeEnvCreateInfoResponse) {
    response = &DescribeComputeEnvCreateInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeComputeEnvCreateInfo
// This API is used to query the compute environment creation information.
//
// error code that may be returned:
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"
func (c *Client) DescribeComputeEnvCreateInfo(request *DescribeComputeEnvCreateInfoRequest) (response *DescribeComputeEnvCreateInfoResponse, err error) {
    return c.DescribeComputeEnvCreateInfoWithContext(context.Background(), request)
}

// DescribeComputeEnvCreateInfo
// This API is used to query the compute environment creation information.
//
// error code that may be returned:
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"
func (c *Client) DescribeComputeEnvCreateInfoWithContext(ctx context.Context, request *DescribeComputeEnvCreateInfoRequest) (response *DescribeComputeEnvCreateInfoResponse, err error) {
    if request == nil {
        request = NewDescribeComputeEnvCreateInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeComputeEnvCreateInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeComputeEnvCreateInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComputeEnvCreateInfosRequest() (request *DescribeComputeEnvCreateInfosRequest) {
    request = &DescribeComputeEnvCreateInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DescribeComputeEnvCreateInfos")
    
    
    return
}

func NewDescribeComputeEnvCreateInfosResponse() (response *DescribeComputeEnvCreateInfosResponse) {
    response = &DescribeComputeEnvCreateInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeComputeEnvCreateInfos
// This API is used to view the list of information of compute environment creation, including name, description, type, environment parameters, notifications, and number of desired nodes.
//
// error code that may be returned:
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  INVALIDPARAMETER_INVALIDPARAMETERCOMBINATION = "InvalidParameter.InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeComputeEnvCreateInfos(request *DescribeComputeEnvCreateInfosRequest) (response *DescribeComputeEnvCreateInfosResponse, err error) {
    return c.DescribeComputeEnvCreateInfosWithContext(context.Background(), request)
}

// DescribeComputeEnvCreateInfos
// This API is used to view the list of information of compute environment creation, including name, description, type, environment parameters, notifications, and number of desired nodes.
//
// error code that may be returned:
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  INVALIDPARAMETER_INVALIDPARAMETERCOMBINATION = "InvalidParameter.InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeComputeEnvCreateInfosWithContext(ctx context.Context, request *DescribeComputeEnvCreateInfosRequest) (response *DescribeComputeEnvCreateInfosResponse, err error) {
    if request == nil {
        request = NewDescribeComputeEnvCreateInfosRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeComputeEnvCreateInfos require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeComputeEnvCreateInfosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComputeEnvsRequest() (request *DescribeComputeEnvsRequest) {
    request = &DescribeComputeEnvsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DescribeComputeEnvs")
    
    
    return
}

func NewDescribeComputeEnvsResponse() (response *DescribeComputeEnvsResponse) {
    response = &DescribeComputeEnvsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeComputeEnvs
// This API is used to get the list of compute environments.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  INVALIDPARAMETER_INVALIDPARAMETERCOMBINATION = "InvalidParameter.InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_RESOURCETYPE = "InvalidParameterValue.ResourceType"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeComputeEnvs(request *DescribeComputeEnvsRequest) (response *DescribeComputeEnvsResponse, err error) {
    return c.DescribeComputeEnvsWithContext(context.Background(), request)
}

// DescribeComputeEnvs
// This API is used to get the list of compute environments.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  INVALIDPARAMETER_INVALIDPARAMETERCOMBINATION = "InvalidParameter.InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_RESOURCETYPE = "InvalidParameterValue.ResourceType"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeComputeEnvsWithContext(ctx context.Context, request *DescribeComputeEnvsRequest) (response *DescribeComputeEnvsResponse, err error) {
    if request == nil {
        request = NewDescribeComputeEnvsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeComputeEnvs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeComputeEnvsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCvmZoneInstanceConfigInfosRequest() (request *DescribeCvmZoneInstanceConfigInfosRequest) {
    request = &DescribeCvmZoneInstanceConfigInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DescribeCvmZoneInstanceConfigInfos")
    
    
    return
}

func NewDescribeCvmZoneInstanceConfigInfosResponse() (response *DescribeCvmZoneInstanceConfigInfosResponse) {
    response = &DescribeCvmZoneInstanceConfigInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCvmZoneInstanceConfigInfos
// This API is used to get the model configuration information of the availability zone of BatchCompute.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLCVM = "InternalError.CallCvm"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDZONEMISMATCHREGION = "InvalidParameterValue.InvalidZoneMismatchRegion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDBATCHINSTANCECHARGETYPE = "InvalidParameterValue.UnsupportedBatchInstanceChargeType"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
func (c *Client) DescribeCvmZoneInstanceConfigInfos(request *DescribeCvmZoneInstanceConfigInfosRequest) (response *DescribeCvmZoneInstanceConfigInfosResponse, err error) {
    return c.DescribeCvmZoneInstanceConfigInfosWithContext(context.Background(), request)
}

// DescribeCvmZoneInstanceConfigInfos
// This API is used to get the model configuration information of the availability zone of BatchCompute.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLCVM = "InternalError.CallCvm"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDZONEMISMATCHREGION = "InvalidParameterValue.InvalidZoneMismatchRegion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDBATCHINSTANCECHARGETYPE = "InvalidParameterValue.UnsupportedBatchInstanceChargeType"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
func (c *Client) DescribeCvmZoneInstanceConfigInfosWithContext(ctx context.Context, request *DescribeCvmZoneInstanceConfigInfosRequest) (response *DescribeCvmZoneInstanceConfigInfosResponse, err error) {
    if request == nil {
        request = NewDescribeCvmZoneInstanceConfigInfosRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCvmZoneInstanceConfigInfos require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCvmZoneInstanceConfigInfosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceCategoriesRequest() (request *DescribeInstanceCategoriesRequest) {
    request = &DescribeInstanceCategoriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DescribeInstanceCategories")
    
    
    return
}

func NewDescribeInstanceCategoriesResponse() (response *DescribeInstanceCategoriesResponse) {
    response = &DescribeInstanceCategoriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceCategories
// Currently, CVM instance families are classified into different category, and each category contains several instance families. This API is used to query the instance category information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstanceCategories(request *DescribeInstanceCategoriesRequest) (response *DescribeInstanceCategoriesResponse, err error) {
    return c.DescribeInstanceCategoriesWithContext(context.Background(), request)
}

// DescribeInstanceCategories
// Currently, CVM instance families are classified into different category, and each category contains several instance families. This API is used to query the instance category information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstanceCategoriesWithContext(ctx context.Context, request *DescribeInstanceCategoriesRequest) (response *DescribeInstanceCategoriesResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceCategoriesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceCategories require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceCategoriesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeJobRequest() (request *DescribeJobRequest) {
    request = &DescribeJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DescribeJob")
    
    
    return
}

func NewDescribeJobResponse() (response *DescribeJobResponse) {
    response = &DescribeJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeJob
// This API is used to query the details of a job, including internal task (`Task`) and dependency (`Dependence`) information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JOBIDMALFORMED = "InvalidParameter.JobIdMalformed"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
func (c *Client) DescribeJob(request *DescribeJobRequest) (response *DescribeJobResponse, err error) {
    return c.DescribeJobWithContext(context.Background(), request)
}

// DescribeJob
// This API is used to query the details of a job, including internal task (`Task`) and dependency (`Dependence`) information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JOBIDMALFORMED = "InvalidParameter.JobIdMalformed"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
func (c *Client) DescribeJobWithContext(ctx context.Context, request *DescribeJobRequest) (response *DescribeJobResponse, err error) {
    if request == nil {
        request = NewDescribeJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeJobSubmitInfoRequest() (request *DescribeJobSubmitInfoRequest) {
    request = &DescribeJobSubmitInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DescribeJobSubmitInfo")
    
    
    return
}

func NewDescribeJobSubmitInfoResponse() (response *DescribeJobSubmitInfoResponse) {
    response = &DescribeJobSubmitInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeJobSubmitInfo
// This API is used to query the submission information of the specified job, with the return including the job submission information used as input parameters in the JobId and SubmitJob APIs.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JOBIDMALFORMED = "InvalidParameter.JobIdMalformed"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
func (c *Client) DescribeJobSubmitInfo(request *DescribeJobSubmitInfoRequest) (response *DescribeJobSubmitInfoResponse, err error) {
    return c.DescribeJobSubmitInfoWithContext(context.Background(), request)
}

// DescribeJobSubmitInfo
// This API is used to query the submission information of the specified job, with the return including the job submission information used as input parameters in the JobId and SubmitJob APIs.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JOBIDMALFORMED = "InvalidParameter.JobIdMalformed"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
func (c *Client) DescribeJobSubmitInfoWithContext(ctx context.Context, request *DescribeJobSubmitInfoRequest) (response *DescribeJobSubmitInfoResponse, err error) {
    if request == nil {
        request = NewDescribeJobSubmitInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeJobSubmitInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeJobSubmitInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeJobsRequest() (request *DescribeJobsRequest) {
    request = &DescribeJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DescribeJobs")
    
    
    return
}

func NewDescribeJobsResponse() (response *DescribeJobsResponse) {
    response = &DescribeJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeJobs
// This API is used to query the overview information of several jobs.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER_INVALIDPARAMETERCOMBINATION = "InvalidParameter.InvalidParameterCombination"
//  INVALIDPARAMETER_JOBIDMALFORMED = "InvalidParameter.JobIdMalformed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeJobs(request *DescribeJobsRequest) (response *DescribeJobsResponse, err error) {
    return c.DescribeJobsWithContext(context.Background(), request)
}

// DescribeJobs
// This API is used to query the overview information of several jobs.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER_INVALIDPARAMETERCOMBINATION = "InvalidParameter.InvalidParameterCombination"
//  INVALIDPARAMETER_JOBIDMALFORMED = "InvalidParameter.JobIdMalformed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeJobsWithContext(ctx context.Context, request *DescribeJobsRequest) (response *DescribeJobsResponse, err error) {
    if request == nil {
        request = NewDescribeJobsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskRequest() (request *DescribeTaskRequest) {
    request = &DescribeTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DescribeTask")
    
    
    return
}

func NewDescribeTaskResponse() (response *DescribeTaskResponse) {
    response = &DescribeTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTask
// This API is used to query the details of a specified task, including information of the task instances inside the task.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JOBIDMALFORMED = "InvalidParameter.JobIdMalformed"
//  RESOURCENOTFOUND_TASK = "ResourceNotFound.Task"
func (c *Client) DescribeTask(request *DescribeTaskRequest) (response *DescribeTaskResponse, err error) {
    return c.DescribeTaskWithContext(context.Background(), request)
}

// DescribeTask
// This API is used to query the details of a specified task, including information of the task instances inside the task.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JOBIDMALFORMED = "InvalidParameter.JobIdMalformed"
//  RESOURCENOTFOUND_TASK = "ResourceNotFound.Task"
func (c *Client) DescribeTaskWithContext(ctx context.Context, request *DescribeTaskRequest) (response *DescribeTaskResponse, err error) {
    if request == nil {
        request = NewDescribeTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskLogsRequest() (request *DescribeTaskLogsRequest) {
    request = &DescribeTaskLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DescribeTaskLogs")
    
    
    return
}

func NewDescribeTaskLogsResponse() (response *DescribeTaskLogsResponse) {
    response = &DescribeTaskLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTaskLogs
// This API is used to get the standard outputs and standard error logs of multiple task instances.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JOBIDMALFORMED = "InvalidParameter.JobIdMalformed"
//  INVALIDPARAMETER_TASKNAME = "InvalidParameter.TaskName"
//  INVALIDPARAMETER_TASKNAMETOOLONG = "InvalidParameter.TaskNameTooLong"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
//  RESOURCENOTFOUND_TASK = "ResourceNotFound.Task"
//  RESOURCENOTFOUND_TASKINSTANCE = "ResourceNotFound.TaskInstance"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTaskLogs(request *DescribeTaskLogsRequest) (response *DescribeTaskLogsResponse, err error) {
    return c.DescribeTaskLogsWithContext(context.Background(), request)
}

// DescribeTaskLogs
// This API is used to get the standard outputs and standard error logs of multiple task instances.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JOBIDMALFORMED = "InvalidParameter.JobIdMalformed"
//  INVALIDPARAMETER_TASKNAME = "InvalidParameter.TaskName"
//  INVALIDPARAMETER_TASKNAMETOOLONG = "InvalidParameter.TaskNameTooLong"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
//  RESOURCENOTFOUND_TASK = "ResourceNotFound.Task"
//  RESOURCENOTFOUND_TASKINSTANCE = "ResourceNotFound.TaskInstance"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTaskLogsWithContext(ctx context.Context, request *DescribeTaskLogsRequest) (response *DescribeTaskLogsResponse, err error) {
    if request == nil {
        request = NewDescribeTaskLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskTemplatesRequest() (request *DescribeTaskTemplatesRequest) {
    request = &DescribeTaskTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DescribeTaskTemplates")
    
    
    return
}

func NewDescribeTaskTemplatesResponse() (response *DescribeTaskTemplatesResponse) {
    response = &DescribeTaskTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTaskTemplates
// This API is used to query the information of task templates.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETERCOMBINATION = "InvalidParameter.InvalidParameterCombination"
//  INVALIDPARAMETER_TASKTEMPLATEIDMALFORMED = "InvalidParameter.TaskTemplateIdMalformed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTaskTemplates(request *DescribeTaskTemplatesRequest) (response *DescribeTaskTemplatesResponse, err error) {
    return c.DescribeTaskTemplatesWithContext(context.Background(), request)
}

// DescribeTaskTemplates
// This API is used to query the information of task templates.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETERCOMBINATION = "InvalidParameter.InvalidParameterCombination"
//  INVALIDPARAMETER_TASKTEMPLATEIDMALFORMED = "InvalidParameter.TaskTemplateIdMalformed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTaskTemplatesWithContext(ctx context.Context, request *DescribeTaskTemplatesRequest) (response *DescribeTaskTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeTaskTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDetachInstancesRequest() (request *DetachInstancesRequest) {
    request = &DetachInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DetachInstances")
    
    
    return
}

func NewDetachInstancesResponse() (response *DetachInstancesResponse) {
    response = &DetachInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DetachInstances
// This API is used to remove instances that from compute environment. 
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLCVM = "InternalError.CallCvm"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  INVALIDPARAMETERVALUE_INSTANCEIDDUPLICATED = "InvalidParameterValue.InstanceIdDuplicated"
//  RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DetachInstances(request *DetachInstancesRequest) (response *DetachInstancesResponse, err error) {
    return c.DetachInstancesWithContext(context.Background(), request)
}

// DetachInstances
// This API is used to remove instances that from compute environment. 
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLCVM = "InternalError.CallCvm"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  INVALIDPARAMETERVALUE_INSTANCEIDDUPLICATED = "InvalidParameterValue.InstanceIdDuplicated"
//  RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DetachInstancesWithContext(ctx context.Context, request *DetachInstancesRequest) (response *DetachInstancesResponse, err error) {
    if request == nil {
        request = NewDetachInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetachInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetachInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyComputeEnvRequest() (request *ModifyComputeEnvRequest) {
    request = &ModifyComputeEnvRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "ModifyComputeEnv")
    
    
    return
}

func NewModifyComputeEnvResponse() (response *ModifyComputeEnvResponse) {
    response = &ModifyComputeEnvResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyComputeEnv
// This API is used to modify the attributes of a compute environment.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLCPMAPI = "InternalError.CallCpmAPI"
//  INVALIDPARAMETER_ENVDESCRIPTIONTOOLONG = "InvalidParameter.EnvDescriptionTooLong"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  INVALIDPARAMETER_ENVNAMETOOLONG = "InvalidParameter.EnvNameTooLong"
//  INVALIDPARAMETERVALUE_INSTANCETYPE = "InvalidParameterValue.InstanceType"
//  INVALIDPARAMETERVALUE_INSTANCETYPEDUPLICATE = "InvalidParameterValue.InstanceTypeDuplicate"
//  INVALIDPARAMETERVALUE_INSTANCETYPESEMPTY = "InvalidParameterValue.InstanceTypesEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  LIMITEXCEEDED_CPUQUOTA = "LimitExceeded.CpuQuota"
//  RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"
//  UNSUPPORTEDOPERATION_ACCEPTOTHERREQUEST = "UnsupportedOperation.AcceptOtherRequest"
//  UNSUPPORTEDOPERATION_COMPUTEENVACCEPTOTHERREQUEST = "UnsupportedOperation.ComputeEnvAcceptOtherRequest"
//  UNSUPPORTEDOPERATION_COMPUTENODEFORBIDTERMINATE = "UnsupportedOperation.ComputeNodeForbidTerminate"
//  UNSUPPORTEDOPERATION_NOTENOUGHCOMPUTENODESTOTERMINATE = "UnsupportedOperation.NotEnoughComputeNodesToTerminate"
func (c *Client) ModifyComputeEnv(request *ModifyComputeEnvRequest) (response *ModifyComputeEnvResponse, err error) {
    return c.ModifyComputeEnvWithContext(context.Background(), request)
}

// ModifyComputeEnv
// This API is used to modify the attributes of a compute environment.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLCPMAPI = "InternalError.CallCpmAPI"
//  INVALIDPARAMETER_ENVDESCRIPTIONTOOLONG = "InvalidParameter.EnvDescriptionTooLong"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  INVALIDPARAMETER_ENVNAMETOOLONG = "InvalidParameter.EnvNameTooLong"
//  INVALIDPARAMETERVALUE_INSTANCETYPE = "InvalidParameterValue.InstanceType"
//  INVALIDPARAMETERVALUE_INSTANCETYPEDUPLICATE = "InvalidParameterValue.InstanceTypeDuplicate"
//  INVALIDPARAMETERVALUE_INSTANCETYPESEMPTY = "InvalidParameterValue.InstanceTypesEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  LIMITEXCEEDED_CPUQUOTA = "LimitExceeded.CpuQuota"
//  RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"
//  UNSUPPORTEDOPERATION_ACCEPTOTHERREQUEST = "UnsupportedOperation.AcceptOtherRequest"
//  UNSUPPORTEDOPERATION_COMPUTEENVACCEPTOTHERREQUEST = "UnsupportedOperation.ComputeEnvAcceptOtherRequest"
//  UNSUPPORTEDOPERATION_COMPUTENODEFORBIDTERMINATE = "UnsupportedOperation.ComputeNodeForbidTerminate"
//  UNSUPPORTEDOPERATION_NOTENOUGHCOMPUTENODESTOTERMINATE = "UnsupportedOperation.NotEnoughComputeNodesToTerminate"
func (c *Client) ModifyComputeEnvWithContext(ctx context.Context, request *ModifyComputeEnvRequest) (response *ModifyComputeEnvResponse, err error) {
    if request == nil {
        request = NewModifyComputeEnvRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyComputeEnv require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyComputeEnvResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTaskTemplateRequest() (request *ModifyTaskTemplateRequest) {
    request = &ModifyTaskTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "ModifyTaskTemplate")
    
    
    return
}

func NewModifyTaskTemplateResponse() (response *ModifyTaskTemplateResponse) {
    response = &ModifyTaskTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTaskTemplate
// This API is used to modify a task template.
//
// error code that may be returned:
//  ALLOWEDONEATTRIBUTEINENVIDANDCOMPUTEENV = "AllowedOneAttributeInEnvIdAndComputeEnv"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CVMPARAMETERS = "InvalidParameter.CvmParameters"
//  INVALIDPARAMETER_TASKNAME = "InvalidParameter.TaskName"
//  INVALIDPARAMETER_TASKNAMETOOLONG = "InvalidParameter.TaskNameTooLong"
//  INVALIDPARAMETER_TASKTEMPLATEDESCRIPTIONTOOLONG = "InvalidParameter.TaskTemplateDescriptionTooLong"
//  INVALIDPARAMETER_TASKTEMPLATEIDMALFORMED = "InvalidParameter.TaskTemplateIdMalformed"
//  INVALIDPARAMETER_TASKTEMPLATENAME = "InvalidParameter.TaskTemplateName"
//  INVALIDPARAMETER_TASKTEMPLATENAMETOOLONG = "InvalidParameter.TaskTemplateNameTooLong"
//  INVALIDPARAMETERATLEASTONEATTRIBUTE = "InvalidParameterAtLeastOneAttribute"
//  INVALIDPARAMETERVALUE_COMPUTEENV = "InvalidParameterValue.ComputeEnv"
//  INVALIDPARAMETERVALUE_DEPENDENCENOTFOUNDTASKNAME = "InvalidParameterValue.DependenceNotFoundTaskName"
//  INVALIDPARAMETERVALUE_DEPENDENCEUNFEASIBLE = "InvalidParameterValue.DependenceUnfeasible"
//  INVALIDPARAMETERVALUE_LOCALPATH = "InvalidParameterValue.LocalPath"
//  INVALIDPARAMETERVALUE_MAXRETRYCOUNT = "InvalidParameterValue.MaxRetryCount"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_REMOTESTORAGEPATH = "InvalidParameterValue.RemoteStoragePath"
//  INVALIDPARAMETERVALUE_REMOTESTORAGESCHEMETYPE = "InvalidParameterValue.RemoteStorageSchemeType"
//  RESOURCENOTFOUND_TASKTEMPLATE = "ResourceNotFound.TaskTemplate"
func (c *Client) ModifyTaskTemplate(request *ModifyTaskTemplateRequest) (response *ModifyTaskTemplateResponse, err error) {
    return c.ModifyTaskTemplateWithContext(context.Background(), request)
}

// ModifyTaskTemplate
// This API is used to modify a task template.
//
// error code that may be returned:
//  ALLOWEDONEATTRIBUTEINENVIDANDCOMPUTEENV = "AllowedOneAttributeInEnvIdAndComputeEnv"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CVMPARAMETERS = "InvalidParameter.CvmParameters"
//  INVALIDPARAMETER_TASKNAME = "InvalidParameter.TaskName"
//  INVALIDPARAMETER_TASKNAMETOOLONG = "InvalidParameter.TaskNameTooLong"
//  INVALIDPARAMETER_TASKTEMPLATEDESCRIPTIONTOOLONG = "InvalidParameter.TaskTemplateDescriptionTooLong"
//  INVALIDPARAMETER_TASKTEMPLATEIDMALFORMED = "InvalidParameter.TaskTemplateIdMalformed"
//  INVALIDPARAMETER_TASKTEMPLATENAME = "InvalidParameter.TaskTemplateName"
//  INVALIDPARAMETER_TASKTEMPLATENAMETOOLONG = "InvalidParameter.TaskTemplateNameTooLong"
//  INVALIDPARAMETERATLEASTONEATTRIBUTE = "InvalidParameterAtLeastOneAttribute"
//  INVALIDPARAMETERVALUE_COMPUTEENV = "InvalidParameterValue.ComputeEnv"
//  INVALIDPARAMETERVALUE_DEPENDENCENOTFOUNDTASKNAME = "InvalidParameterValue.DependenceNotFoundTaskName"
//  INVALIDPARAMETERVALUE_DEPENDENCEUNFEASIBLE = "InvalidParameterValue.DependenceUnfeasible"
//  INVALIDPARAMETERVALUE_LOCALPATH = "InvalidParameterValue.LocalPath"
//  INVALIDPARAMETERVALUE_MAXRETRYCOUNT = "InvalidParameterValue.MaxRetryCount"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_REMOTESTORAGEPATH = "InvalidParameterValue.RemoteStoragePath"
//  INVALIDPARAMETERVALUE_REMOTESTORAGESCHEMETYPE = "InvalidParameterValue.RemoteStorageSchemeType"
//  RESOURCENOTFOUND_TASKTEMPLATE = "ResourceNotFound.TaskTemplate"
func (c *Client) ModifyTaskTemplateWithContext(ctx context.Context, request *ModifyTaskTemplateRequest) (response *ModifyTaskTemplateResponse, err error) {
    if request == nil {
        request = NewModifyTaskTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTaskTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTaskTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewRetryJobsRequest() (request *RetryJobsRequest) {
    request = &RetryJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "RetryJobs")
    
    
    return
}

func NewRetryJobsResponse() (response *RetryJobsResponse) {
    response = &RetryJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RetryJobs
// This API is used to retry the failed task instances in a job.
//
// Job retry is supported only if a job is in the "FAILED" state. After the retry operation succeeds, the job will retry the failing task instances in each task in turn according to the task dependencies specified in the "DAG". The history information of the task instances will be reset, the instances will participate in subsequent scheduling and execution as if they are run for the first time.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JOBIDMALFORMED = "InvalidParameter.JobIdMalformed"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RetryJobs(request *RetryJobsRequest) (response *RetryJobsResponse, err error) {
    return c.RetryJobsWithContext(context.Background(), request)
}

// RetryJobs
// This API is used to retry the failed task instances in a job.
//
// Job retry is supported only if a job is in the "FAILED" state. After the retry operation succeeds, the job will retry the failing task instances in each task in turn according to the task dependencies specified in the "DAG". The history information of the task instances will be reset, the instances will participate in subsequent scheduling and execution as if they are run for the first time.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JOBIDMALFORMED = "InvalidParameter.JobIdMalformed"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RetryJobsWithContext(ctx context.Context, request *RetryJobsRequest) (response *RetryJobsResponse, err error) {
    if request == nil {
        request = NewRetryJobsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RetryJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewRetryJobsResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateComputeNodeRequest() (request *TerminateComputeNodeRequest) {
    request = &TerminateComputeNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "TerminateComputeNode")
    
    
    return
}

func NewTerminateComputeNodeResponse() (response *TerminateComputeNodeResponse) {
    response = &TerminateComputeNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TerminateComputeNode
// This API is used to terminate a compute node.
//
// Termination is allowed for nodes in the CREATED, CREATION_FAILED, RUNNING or ABNORMAL state.
//
// error code that may be returned:
//  INVALIDPARAMETER_COMPUTENODEIDMALFORMED = "InvalidParameter.ComputeNodeIdMalformed"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"
//  RESOURCENOTFOUND_COMPUTENODE = "ResourceNotFound.ComputeNode"
//  UNSUPPORTEDOPERATION_ACCEPTOTHERREQUEST = "UnsupportedOperation.AcceptOtherRequest"
//  UNSUPPORTEDOPERATION_COMPUTEENVACCEPTOTHERREQUEST = "UnsupportedOperation.ComputeEnvAcceptOtherRequest"
//  UNSUPPORTEDOPERATION_COMPUTEENVOPERATION = "UnsupportedOperation.ComputeEnvOperation"
//  UNSUPPORTEDOPERATION_COMPUTENODEFORBIDTERMINATE = "UnsupportedOperation.ComputeNodeForbidTerminate"
//  UNSUPPORTEDOPERATION_COMPUTENODEISTERMINATING = "UnsupportedOperation.ComputeNodeIsTerminating"
func (c *Client) TerminateComputeNode(request *TerminateComputeNodeRequest) (response *TerminateComputeNodeResponse, err error) {
    return c.TerminateComputeNodeWithContext(context.Background(), request)
}

// TerminateComputeNode
// This API is used to terminate a compute node.
//
// Termination is allowed for nodes in the CREATED, CREATION_FAILED, RUNNING or ABNORMAL state.
//
// error code that may be returned:
//  INVALIDPARAMETER_COMPUTENODEIDMALFORMED = "InvalidParameter.ComputeNodeIdMalformed"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"
//  RESOURCENOTFOUND_COMPUTENODE = "ResourceNotFound.ComputeNode"
//  UNSUPPORTEDOPERATION_ACCEPTOTHERREQUEST = "UnsupportedOperation.AcceptOtherRequest"
//  UNSUPPORTEDOPERATION_COMPUTEENVACCEPTOTHERREQUEST = "UnsupportedOperation.ComputeEnvAcceptOtherRequest"
//  UNSUPPORTEDOPERATION_COMPUTEENVOPERATION = "UnsupportedOperation.ComputeEnvOperation"
//  UNSUPPORTEDOPERATION_COMPUTENODEFORBIDTERMINATE = "UnsupportedOperation.ComputeNodeForbidTerminate"
//  UNSUPPORTEDOPERATION_COMPUTENODEISTERMINATING = "UnsupportedOperation.ComputeNodeIsTerminating"
func (c *Client) TerminateComputeNodeWithContext(ctx context.Context, request *TerminateComputeNodeRequest) (response *TerminateComputeNodeResponse, err error) {
    if request == nil {
        request = NewTerminateComputeNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateComputeNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateComputeNodeResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateComputeNodesRequest() (request *TerminateComputeNodesRequest) {
    request = &TerminateComputeNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "TerminateComputeNodes")
    
    
    return
}

func NewTerminateComputeNodesResponse() (response *TerminateComputeNodesResponse) {
    response = &TerminateComputeNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TerminateComputeNodes
// This API is used to terminate compute nodes in batches. It is not allowed to repeatedly terminate the same node.
//
// error code that may be returned:
//  INVALIDPARAMETER_COMPUTENODEIDMALFORMED = "InvalidParameter.ComputeNodeIdMalformed"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"
//  RESOURCENOTFOUND_COMPUTENODE = "ResourceNotFound.ComputeNode"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ACCEPTOTHERREQUEST = "UnsupportedOperation.AcceptOtherRequest"
//  UNSUPPORTEDOPERATION_COMPUTEENVACCEPTOTHERREQUEST = "UnsupportedOperation.ComputeEnvAcceptOtherRequest"
//  UNSUPPORTEDOPERATION_COMPUTENODEFORBIDTERMINATE = "UnsupportedOperation.ComputeNodeForbidTerminate"
func (c *Client) TerminateComputeNodes(request *TerminateComputeNodesRequest) (response *TerminateComputeNodesResponse, err error) {
    return c.TerminateComputeNodesWithContext(context.Background(), request)
}

// TerminateComputeNodes
// This API is used to terminate compute nodes in batches. It is not allowed to repeatedly terminate the same node.
//
// error code that may be returned:
//  INVALIDPARAMETER_COMPUTENODEIDMALFORMED = "InvalidParameter.ComputeNodeIdMalformed"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"
//  RESOURCENOTFOUND_COMPUTENODE = "ResourceNotFound.ComputeNode"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ACCEPTOTHERREQUEST = "UnsupportedOperation.AcceptOtherRequest"
//  UNSUPPORTEDOPERATION_COMPUTEENVACCEPTOTHERREQUEST = "UnsupportedOperation.ComputeEnvAcceptOtherRequest"
//  UNSUPPORTEDOPERATION_COMPUTENODEFORBIDTERMINATE = "UnsupportedOperation.ComputeNodeForbidTerminate"
func (c *Client) TerminateComputeNodesWithContext(ctx context.Context, request *TerminateComputeNodesRequest) (response *TerminateComputeNodesResponse, err error) {
    if request == nil {
        request = NewTerminateComputeNodesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateComputeNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateComputeNodesResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateJobRequest() (request *TerminateJobRequest) {
    request = &TerminateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "TerminateJob")
    
    
    return
}

func NewTerminateJobResponse() (response *TerminateJobResponse) {
    response = &TerminateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TerminateJob
// This API is used to terminate a job.
//
// Termination is prohibited if a job is in the `SUBMITTED` state and does not take effect if it is in the `SUCCEED` state.
//
// Job termination is an asynchronous process, and the time it takes to complete the entire process is directly proportional to the total number of tasks. The effect of terminating a job is equivalent to performing the TerminateTaskInstance operation on all the task instances contained in the job. For more information on the specific effect and usage, see TerminateTaskInstance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JOBIDMALFORMED = "InvalidParameter.JobIdMalformed"
//  RESOURCEINUSE_JOB = "ResourceInUse.Job"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_TERMINATEOPERATIONWITHENVID = "UnsupportedOperation.TerminateOperationWithEnvId"
func (c *Client) TerminateJob(request *TerminateJobRequest) (response *TerminateJobResponse, err error) {
    return c.TerminateJobWithContext(context.Background(), request)
}

// TerminateJob
// This API is used to terminate a job.
//
// Termination is prohibited if a job is in the `SUBMITTED` state and does not take effect if it is in the `SUCCEED` state.
//
// Job termination is an asynchronous process, and the time it takes to complete the entire process is directly proportional to the total number of tasks. The effect of terminating a job is equivalent to performing the TerminateTaskInstance operation on all the task instances contained in the job. For more information on the specific effect and usage, see TerminateTaskInstance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JOBIDMALFORMED = "InvalidParameter.JobIdMalformed"
//  RESOURCEINUSE_JOB = "ResourceInUse.Job"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_TERMINATEOPERATIONWITHENVID = "UnsupportedOperation.TerminateOperationWithEnvId"
func (c *Client) TerminateJobWithContext(ctx context.Context, request *TerminateJobRequest) (response *TerminateJobResponse, err error) {
    if request == nil {
        request = NewTerminateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateJobResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateTaskInstanceRequest() (request *TerminateTaskInstanceRequest) {
    request = &TerminateTaskInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "TerminateTaskInstance")
    
    
    return
}

func NewTerminateTaskInstanceResponse() (response *TerminateTaskInstanceResponse) {
    response = &TerminateTaskInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TerminateTaskInstance
// This API is used to terminate a task instance.
//
// `SUCCEED` and `FAILED` task instances: No action
//
// `SUBMITTED`, `PENDING`, and `RUNNABLE` task instances: Change status to `FAILED`.
//
// `STARTING`, `RUNNING` and `FAILED_INTERRUPTED` task instances: If `EnvId` is not specified, the CVM instance is terminated, and then the task status goes to `FAILED`. If `EnvId` is specified, the task instance changes to `FAILED`, then the related CVM instance is restarted. 
//
// `FAILED_INTERRUPTED` task instances: The related resources and quotas will be released only after the termination actually succeeds.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JOBIDMALFORMED = "InvalidParameter.JobIdMalformed"
//  RESOURCENOTFOUND_TASKINSTANCE = "ResourceNotFound.TaskInstance"
//  UNSUPPORTEDOPERATION_TERMINATEOPERATIONWITHENVID = "UnsupportedOperation.TerminateOperationWithEnvId"
func (c *Client) TerminateTaskInstance(request *TerminateTaskInstanceRequest) (response *TerminateTaskInstanceResponse, err error) {
    return c.TerminateTaskInstanceWithContext(context.Background(), request)
}

// TerminateTaskInstance
// This API is used to terminate a task instance.
//
// `SUCCEED` and `FAILED` task instances: No action
//
// `SUBMITTED`, `PENDING`, and `RUNNABLE` task instances: Change status to `FAILED`.
//
// `STARTING`, `RUNNING` and `FAILED_INTERRUPTED` task instances: If `EnvId` is not specified, the CVM instance is terminated, and then the task status goes to `FAILED`. If `EnvId` is specified, the task instance changes to `FAILED`, then the related CVM instance is restarted. 
//
// `FAILED_INTERRUPTED` task instances: The related resources and quotas will be released only after the termination actually succeeds.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JOBIDMALFORMED = "InvalidParameter.JobIdMalformed"
//  RESOURCENOTFOUND_TASKINSTANCE = "ResourceNotFound.TaskInstance"
//  UNSUPPORTEDOPERATION_TERMINATEOPERATIONWITHENVID = "UnsupportedOperation.TerminateOperationWithEnvId"
func (c *Client) TerminateTaskInstanceWithContext(ctx context.Context, request *TerminateTaskInstanceRequest) (response *TerminateTaskInstanceResponse, err error) {
    if request == nil {
        request = NewTerminateTaskInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateTaskInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateTaskInstanceResponse()
    err = c.Send(request, response)
    return
}
