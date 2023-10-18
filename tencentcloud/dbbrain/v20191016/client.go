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

package v20191016

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2019-10-16"

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


func NewAddUserContactRequest() (request *AddUserContactRequest) {
    request = &AddUserContactRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "AddUserContact")
    
    
    return
}

func NewAddUserContactResponse() (response *AddUserContactResponse) {
    response = &AddUserContactResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddUserContact
// This API is used to add the contact name and email.. The return value is the successfully added contact ID. Select Guangzhou for Region.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddUserContact(request *AddUserContactRequest) (response *AddUserContactResponse, err error) {
    return c.AddUserContactWithContext(context.Background(), request)
}

// AddUserContact
// This API is used to add the contact name and email.. The return value is the successfully added contact ID. Select Guangzhou for Region.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddUserContactWithContext(ctx context.Context, request *AddUserContactRequest) (response *AddUserContactResponse, err error) {
    if request == nil {
        request = NewAddUserContactRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddUserContact require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddUserContactResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDBDiagReportTaskRequest() (request *CreateDBDiagReportTaskRequest) {
    request = &CreateDBDiagReportTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "CreateDBDiagReportTask")
    
    
    return
}

func NewCreateDBDiagReportTaskResponse() (response *CreateDBDiagReportTaskResponse) {
    response = &CreateDBDiagReportTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDBDiagReportTask
// This API is used to create a health report and send it via email as configured.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) CreateDBDiagReportTask(request *CreateDBDiagReportTaskRequest) (response *CreateDBDiagReportTaskResponse, err error) {
    return c.CreateDBDiagReportTaskWithContext(context.Background(), request)
}

// CreateDBDiagReportTask
// This API is used to create a health report and send it via email as configured.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) CreateDBDiagReportTaskWithContext(ctx context.Context, request *CreateDBDiagReportTaskRequest) (response *CreateDBDiagReportTaskResponse, err error) {
    if request == nil {
        request = NewCreateDBDiagReportTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDBDiagReportTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDBDiagReportTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDBDiagReportUrlRequest() (request *CreateDBDiagReportUrlRequest) {
    request = &CreateDBDiagReportUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "CreateDBDiagReportUrl")
    
    
    return
}

func NewCreateDBDiagReportUrlResponse() (response *CreateDBDiagReportUrlResponse) {
    response = &CreateDBDiagReportUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDBDiagReportUrl
// This API is used to create a URL for a health report.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateDBDiagReportUrl(request *CreateDBDiagReportUrlRequest) (response *CreateDBDiagReportUrlResponse, err error) {
    return c.CreateDBDiagReportUrlWithContext(context.Background(), request)
}

// CreateDBDiagReportUrl
// This API is used to create a URL for a health report.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateDBDiagReportUrlWithContext(ctx context.Context, request *CreateDBDiagReportUrlRequest) (response *CreateDBDiagReportUrlResponse, err error) {
    if request == nil {
        request = NewCreateDBDiagReportUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDBDiagReportUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDBDiagReportUrlResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMailProfileRequest() (request *CreateMailProfileRequest) {
    request = &CreateMailProfileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "CreateMailProfile")
    
    
    return
}

func NewCreateMailProfileResponse() (response *CreateMailProfileResponse) {
    response = &CreateMailProfileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateMailProfile
// This API is used to create the email configuration. The input parameter `ProfileType` represents the type of the email configuration. Valid values: `dbScan_mail_configuration` (email configuration of database inspection report) and `scheduler_mail_configuration` (email sending configuration of regularly generated health report). Select Guangzhou for Region, regardless of the region where the instance belongs.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateMailProfile(request *CreateMailProfileRequest) (response *CreateMailProfileResponse, err error) {
    return c.CreateMailProfileWithContext(context.Background(), request)
}

// CreateMailProfile
// This API is used to create the email configuration. The input parameter `ProfileType` represents the type of the email configuration. Valid values: `dbScan_mail_configuration` (email configuration of database inspection report) and `scheduler_mail_configuration` (email sending configuration of regularly generated health report). Select Guangzhou for Region, regardless of the region where the instance belongs.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateMailProfileWithContext(ctx context.Context, request *CreateMailProfileRequest) (response *CreateMailProfileResponse, err error) {
    if request == nil {
        request = NewCreateMailProfileRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMailProfile require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMailProfileResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSchedulerMailProfileRequest() (request *CreateSchedulerMailProfileRequest) {
    request = &CreateSchedulerMailProfileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "CreateSchedulerMailProfile")
    
    
    return
}

func NewCreateSchedulerMailProfileResponse() (response *CreateSchedulerMailProfileResponse) {
    response = &CreateSchedulerMailProfileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSchedulerMailProfile
// This API is used to create the regular generation time of the health reports and the regular email sending configuration. Pass in the regular generation time of the health reports as a parameter (Monday to Sunday) to set the regular generation time of the health reports, and save the corresponding regular email sending configuration.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSchedulerMailProfile(request *CreateSchedulerMailProfileRequest) (response *CreateSchedulerMailProfileResponse, err error) {
    return c.CreateSchedulerMailProfileWithContext(context.Background(), request)
}

// CreateSchedulerMailProfile
// This API is used to create the regular generation time of the health reports and the regular email sending configuration. Pass in the regular generation time of the health reports as a parameter (Monday to Sunday) to set the regular generation time of the health reports, and save the corresponding regular email sending configuration.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSchedulerMailProfileWithContext(ctx context.Context, request *CreateSchedulerMailProfileRequest) (response *CreateSchedulerMailProfileResponse, err error) {
    if request == nil {
        request = NewCreateSchedulerMailProfileRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSchedulerMailProfile require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSchedulerMailProfileResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllUserContactRequest() (request *DescribeAllUserContactRequest) {
    request = &DescribeAllUserContactRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeAllUserContact")
    
    
    return
}

func NewDescribeAllUserContactResponse() (response *DescribeAllUserContactResponse) {
    response = &DescribeAllUserContactResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAllUserContact
// This API is used to obtain the information of the contact in the email.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeAllUserContact(request *DescribeAllUserContactRequest) (response *DescribeAllUserContactResponse, err error) {
    return c.DescribeAllUserContactWithContext(context.Background(), request)
}

// DescribeAllUserContact
// This API is used to obtain the information of the contact in the email.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeAllUserContactWithContext(ctx context.Context, request *DescribeAllUserContactRequest) (response *DescribeAllUserContactResponse, err error) {
    if request == nil {
        request = NewDescribeAllUserContactRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAllUserContact require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAllUserContactResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllUserGroupRequest() (request *DescribeAllUserGroupRequest) {
    request = &DescribeAllUserGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeAllUserGroup")
    
    
    return
}

func NewDescribeAllUserGroupResponse() (response *DescribeAllUserGroupResponse) {
    response = &DescribeAllUserGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAllUserGroup
// This API is used to obtain the information of the contact group in the email.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeAllUserGroup(request *DescribeAllUserGroupRequest) (response *DescribeAllUserGroupResponse, err error) {
    return c.DescribeAllUserGroupWithContext(context.Background(), request)
}

// DescribeAllUserGroup
// This API is used to obtain the information of the contact group in the email.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeAllUserGroupWithContext(ctx context.Context, request *DescribeAllUserGroupRequest) (response *DescribeAllUserGroupResponse, err error) {
    if request == nil {
        request = NewDescribeAllUserGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAllUserGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAllUserGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBDiagEventRequest() (request *DescribeDBDiagEventRequest) {
    request = &DescribeDBDiagEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeDBDiagEvent")
    
    
    return
}

func NewDescribeDBDiagEventResponse() (response *DescribeDBDiagEventResponse) {
    response = &DescribeDBDiagEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBDiagEvent
// This API is used to get the details of an instance exception diagnosis event.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDBDiagEvent(request *DescribeDBDiagEventRequest) (response *DescribeDBDiagEventResponse, err error) {
    return c.DescribeDBDiagEventWithContext(context.Background(), request)
}

// DescribeDBDiagEvent
// This API is used to get the details of an instance exception diagnosis event.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDBDiagEventWithContext(ctx context.Context, request *DescribeDBDiagEventRequest) (response *DescribeDBDiagEventResponse, err error) {
    if request == nil {
        request = NewDescribeDBDiagEventRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBDiagEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBDiagEventResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBDiagHistoryRequest() (request *DescribeDBDiagHistoryRequest) {
    request = &DescribeDBDiagHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeDBDiagHistory")
    
    
    return
}

func NewDescribeDBDiagHistoryResponse() (response *DescribeDBDiagHistoryResponse) {
    response = &DescribeDBDiagHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBDiagHistory
// This API is used to get the list of instance diagnosis events.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDBDiagHistory(request *DescribeDBDiagHistoryRequest) (response *DescribeDBDiagHistoryResponse, err error) {
    return c.DescribeDBDiagHistoryWithContext(context.Background(), request)
}

// DescribeDBDiagHistory
// This API is used to get the list of instance diagnosis events.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDBDiagHistoryWithContext(ctx context.Context, request *DescribeDBDiagHistoryRequest) (response *DescribeDBDiagHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeDBDiagHistoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBDiagHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBDiagHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBDiagReportTasksRequest() (request *DescribeDBDiagReportTasksRequest) {
    request = &DescribeDBDiagReportTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeDBDiagReportTasks")
    
    
    return
}

func NewDescribeDBDiagReportTasksResponse() (response *DescribeDBDiagReportTasksResponse) {
    response = &DescribeDBDiagReportTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBDiagReportTasks
// This API is used to query the list of health report generation tasks.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeDBDiagReportTasks(request *DescribeDBDiagReportTasksRequest) (response *DescribeDBDiagReportTasksResponse, err error) {
    return c.DescribeDBDiagReportTasksWithContext(context.Background(), request)
}

// DescribeDBDiagReportTasks
// This API is used to query the list of health report generation tasks.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeDBDiagReportTasksWithContext(ctx context.Context, request *DescribeDBDiagReportTasksRequest) (response *DescribeDBDiagReportTasksResponse, err error) {
    if request == nil {
        request = NewDescribeDBDiagReportTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBDiagReportTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBDiagReportTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBSpaceStatusRequest() (request *DescribeDBSpaceStatusRequest) {
    request = &DescribeDBSpaceStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeDBSpaceStatus")
    
    
    return
}

func NewDescribeDBSpaceStatusResponse() (response *DescribeDBSpaceStatusResponse) {
    response = &DescribeDBSpaceStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBSpaceStatus
// This API is used to query the overview of instance space usage during a specified time period, including disk usage growth (MB), available disk space (MB), total disk space (MB), and estimated number of available days.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDBSpaceStatus(request *DescribeDBSpaceStatusRequest) (response *DescribeDBSpaceStatusResponse, err error) {
    return c.DescribeDBSpaceStatusWithContext(context.Background(), request)
}

// DescribeDBSpaceStatus
// This API is used to query the overview of instance space usage during a specified time period, including disk usage growth (MB), available disk space (MB), total disk space (MB), and estimated number of available days.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDBSpaceStatusWithContext(ctx context.Context, request *DescribeDBSpaceStatusRequest) (response *DescribeDBSpaceStatusResponse, err error) {
    if request == nil {
        request = NewDescribeDBSpaceStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBSpaceStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBSpaceStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDiagDBInstancesRequest() (request *DescribeDiagDBInstancesRequest) {
    request = &DescribeDiagDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeDiagDBInstances")
    
    
    return
}

func NewDescribeDiagDBInstancesResponse() (response *DescribeDiagDBInstancesResponse) {
    response = &DescribeDiagDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDiagDBInstances
// This API is used to obtain the instance information list. Select Guangzhou for Region.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDiagDBInstances(request *DescribeDiagDBInstancesRequest) (response *DescribeDiagDBInstancesResponse, err error) {
    return c.DescribeDiagDBInstancesWithContext(context.Background(), request)
}

// DescribeDiagDBInstances
// This API is used to obtain the instance information list. Select Guangzhou for Region.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDiagDBInstancesWithContext(ctx context.Context, request *DescribeDiagDBInstancesRequest) (response *DescribeDiagDBInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeDiagDBInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDiagDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDiagDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHealthScoreRequest() (request *DescribeHealthScoreRequest) {
    request = &DescribeHealthScoreRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeHealthScore")
    
    
    return
}

func NewDescribeHealthScoreResponse() (response *DescribeHealthScoreResponse) {
    response = &DescribeHealthScoreResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHealthScore
// This API is used to obtain the health score and deduction for exceptions in the specified time period (30 minutes) based on the instance ID.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeHealthScore(request *DescribeHealthScoreRequest) (response *DescribeHealthScoreResponse, err error) {
    return c.DescribeHealthScoreWithContext(context.Background(), request)
}

// DescribeHealthScore
// This API is used to obtain the health score and deduction for exceptions in the specified time period (30 minutes) based on the instance ID.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeHealthScoreWithContext(ctx context.Context, request *DescribeHealthScoreRequest) (response *DescribeHealthScoreResponse, err error) {
    if request == nil {
        request = NewDescribeHealthScoreRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHealthScore require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHealthScoreResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMailProfileRequest() (request *DescribeMailProfileRequest) {
    request = &DescribeMailProfileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeMailProfile")
    
    
    return
}

func NewDescribeMailProfileResponse() (response *DescribeMailProfileResponse) {
    response = &DescribeMailProfileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMailProfile
// This API is used to obtain the email sending configurations, including the email configuration for database inspection and the email sending configuration for regularly generated health reports. Select Guangzhou for Region.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeMailProfile(request *DescribeMailProfileRequest) (response *DescribeMailProfileResponse, err error) {
    return c.DescribeMailProfileWithContext(context.Background(), request)
}

// DescribeMailProfile
// This API is used to obtain the email sending configurations, including the email configuration for database inspection and the email sending configuration for regularly generated health reports. Select Guangzhou for Region.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeMailProfileWithContext(ctx context.Context, request *DescribeMailProfileRequest) (response *DescribeMailProfileResponse, err error) {
    if request == nil {
        request = NewDescribeMailProfileRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMailProfile require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMailProfileResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlowLogTimeSeriesStatsRequest() (request *DescribeSlowLogTimeSeriesStatsRequest) {
    request = &DescribeSlowLogTimeSeriesStatsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeSlowLogTimeSeriesStats")
    
    
    return
}

func NewDescribeSlowLogTimeSeriesStatsResponse() (response *DescribeSlowLogTimeSeriesStatsResponse) {
    response = &DescribeSlowLogTimeSeriesStatsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSlowLogTimeSeriesStats
// This API is used to get the slow log statistics histogram.
//
// error code that may be returned:
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeSlowLogTimeSeriesStats(request *DescribeSlowLogTimeSeriesStatsRequest) (response *DescribeSlowLogTimeSeriesStatsResponse, err error) {
    return c.DescribeSlowLogTimeSeriesStatsWithContext(context.Background(), request)
}

// DescribeSlowLogTimeSeriesStats
// This API is used to get the slow log statistics histogram.
//
// error code that may be returned:
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeSlowLogTimeSeriesStatsWithContext(ctx context.Context, request *DescribeSlowLogTimeSeriesStatsRequest) (response *DescribeSlowLogTimeSeriesStatsResponse, err error) {
    if request == nil {
        request = NewDescribeSlowLogTimeSeriesStatsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSlowLogTimeSeriesStats require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSlowLogTimeSeriesStatsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlowLogTopSqlsRequest() (request *DescribeSlowLogTopSqlsRequest) {
    request = &DescribeSlowLogTopSqlsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeSlowLogTopSqls")
    
    
    return
}

func NewDescribeSlowLogTopSqlsResponse() (response *DescribeSlowLogTopSqlsResponse) {
    response = &DescribeSlowLogTopSqlsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSlowLogTopSqls
// This API is used to get and sort the top slow SQL statements in a specified time period by the aggregation mode of SQL template plus schema.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeSlowLogTopSqls(request *DescribeSlowLogTopSqlsRequest) (response *DescribeSlowLogTopSqlsResponse, err error) {
    return c.DescribeSlowLogTopSqlsWithContext(context.Background(), request)
}

// DescribeSlowLogTopSqls
// This API is used to get and sort the top slow SQL statements in a specified time period by the aggregation mode of SQL template plus schema.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeSlowLogTopSqlsWithContext(ctx context.Context, request *DescribeSlowLogTopSqlsRequest) (response *DescribeSlowLogTopSqlsResponse, err error) {
    if request == nil {
        request = NewDescribeSlowLogTopSqlsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSlowLogTopSqls require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSlowLogTopSqlsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlowLogUserHostStatsRequest() (request *DescribeSlowLogUserHostStatsRequest) {
    request = &DescribeSlowLogUserHostStatsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeSlowLogUserHostStats")
    
    
    return
}

func NewDescribeSlowLogUserHostStatsResponse() (response *DescribeSlowLogUserHostStatsResponse) {
    response = &DescribeSlowLogUserHostStatsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSlowLogUserHostStats
// This API is used to obtain the statistical distribution chart of slow log source addresses.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeSlowLogUserHostStats(request *DescribeSlowLogUserHostStatsRequest) (response *DescribeSlowLogUserHostStatsResponse, err error) {
    return c.DescribeSlowLogUserHostStatsWithContext(context.Background(), request)
}

// DescribeSlowLogUserHostStats
// This API is used to obtain the statistical distribution chart of slow log source addresses.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeSlowLogUserHostStatsWithContext(ctx context.Context, request *DescribeSlowLogUserHostStatsRequest) (response *DescribeSlowLogUserHostStatsResponse, err error) {
    if request == nil {
        request = NewDescribeSlowLogUserHostStatsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSlowLogUserHostStats require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSlowLogUserHostStatsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopSpaceSchemaTimeSeriesRequest() (request *DescribeTopSpaceSchemaTimeSeriesRequest) {
    request = &DescribeTopSpaceSchemaTimeSeriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeTopSpaceSchemaTimeSeries")
    
    
    return
}

func NewDescribeTopSpaceSchemaTimeSeriesResponse() (response *DescribeTopSpaceSchemaTimeSeriesResponse) {
    response = &DescribeTopSpaceSchemaTimeSeriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTopSpaceSchemaTimeSeries
// This API is used to query the daily space data of top databases consuming the most instance space. The data is daily collected by DBbrain during a specified time period. The return results are sorted by size by default.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeTopSpaceSchemaTimeSeries(request *DescribeTopSpaceSchemaTimeSeriesRequest) (response *DescribeTopSpaceSchemaTimeSeriesResponse, err error) {
    return c.DescribeTopSpaceSchemaTimeSeriesWithContext(context.Background(), request)
}

// DescribeTopSpaceSchemaTimeSeries
// This API is used to query the daily space data of top databases consuming the most instance space. The data is daily collected by DBbrain during a specified time period. The return results are sorted by size by default.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeTopSpaceSchemaTimeSeriesWithContext(ctx context.Context, request *DescribeTopSpaceSchemaTimeSeriesRequest) (response *DescribeTopSpaceSchemaTimeSeriesResponse, err error) {
    if request == nil {
        request = NewDescribeTopSpaceSchemaTimeSeriesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopSpaceSchemaTimeSeries require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopSpaceSchemaTimeSeriesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopSpaceSchemasRequest() (request *DescribeTopSpaceSchemasRequest) {
    request = &DescribeTopSpaceSchemasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeTopSpaceSchemas")
    
    
    return
}

func NewDescribeTopSpaceSchemasResponse() (response *DescribeTopSpaceSchemasResponse) {
    response = &DescribeTopSpaceSchemasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTopSpaceSchemas
// This API is used to query real-time space statistics of top databases. The return results are sorted by size by default.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeTopSpaceSchemas(request *DescribeTopSpaceSchemasRequest) (response *DescribeTopSpaceSchemasResponse, err error) {
    return c.DescribeTopSpaceSchemasWithContext(context.Background(), request)
}

// DescribeTopSpaceSchemas
// This API is used to query real-time space statistics of top databases. The return results are sorted by size by default.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeTopSpaceSchemasWithContext(ctx context.Context, request *DescribeTopSpaceSchemasRequest) (response *DescribeTopSpaceSchemasResponse, err error) {
    if request == nil {
        request = NewDescribeTopSpaceSchemasRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopSpaceSchemas require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopSpaceSchemasResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopSpaceTableTimeSeriesRequest() (request *DescribeTopSpaceTableTimeSeriesRequest) {
    request = &DescribeTopSpaceTableTimeSeriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeTopSpaceTableTimeSeries")
    
    
    return
}

func NewDescribeTopSpaceTableTimeSeriesResponse() (response *DescribeTopSpaceTableTimeSeriesResponse) {
    response = &DescribeTopSpaceTableTimeSeriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTopSpaceTableTimeSeries
// This API is used to query the daily space data of top tables consuming the most instance space. The data is daily collected by DBbrain during a specified time period. The return results are sorted by size by default.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTopSpaceTableTimeSeries(request *DescribeTopSpaceTableTimeSeriesRequest) (response *DescribeTopSpaceTableTimeSeriesResponse, err error) {
    return c.DescribeTopSpaceTableTimeSeriesWithContext(context.Background(), request)
}

// DescribeTopSpaceTableTimeSeries
// This API is used to query the daily space data of top tables consuming the most instance space. The data is daily collected by DBbrain during a specified time period. The return results are sorted by size by default.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTopSpaceTableTimeSeriesWithContext(ctx context.Context, request *DescribeTopSpaceTableTimeSeriesRequest) (response *DescribeTopSpaceTableTimeSeriesResponse, err error) {
    if request == nil {
        request = NewDescribeTopSpaceTableTimeSeriesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopSpaceTableTimeSeries require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopSpaceTableTimeSeriesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopSpaceTablesRequest() (request *DescribeTopSpaceTablesRequest) {
    request = &DescribeTopSpaceTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeTopSpaceTables")
    
    
    return
}

func NewDescribeTopSpaceTablesResponse() (response *DescribeTopSpaceTablesResponse) {
    response = &DescribeTopSpaceTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTopSpaceTables
// This API is used to query real-time space statistics of top tables of an instance. The return results are sorted by size by default.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTopSpaceTables(request *DescribeTopSpaceTablesRequest) (response *DescribeTopSpaceTablesResponse, err error) {
    return c.DescribeTopSpaceTablesWithContext(context.Background(), request)
}

// DescribeTopSpaceTables
// This API is used to query real-time space statistics of top tables of an instance. The return results are sorted by size by default.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTopSpaceTablesWithContext(ctx context.Context, request *DescribeTopSpaceTablesRequest) (response *DescribeTopSpaceTablesResponse, err error) {
    if request == nil {
        request = NewDescribeTopSpaceTablesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopSpaceTables require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopSpaceTablesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserSqlAdviceRequest() (request *DescribeUserSqlAdviceRequest) {
    request = &DescribeUserSqlAdviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeUserSqlAdvice")
    
    
    return
}

func NewDescribeUserSqlAdviceResponse() (response *DescribeUserSqlAdviceResponse) {
    response = &DescribeUserSqlAdviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserSqlAdvice
// This API is used to obtain SQL statement optimization suggestions.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeUserSqlAdvice(request *DescribeUserSqlAdviceRequest) (response *DescribeUserSqlAdviceResponse, err error) {
    return c.DescribeUserSqlAdviceWithContext(context.Background(), request)
}

// DescribeUserSqlAdvice
// This API is used to obtain SQL statement optimization suggestions.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeUserSqlAdviceWithContext(ctx context.Context, request *DescribeUserSqlAdviceRequest) (response *DescribeUserSqlAdviceResponse, err error) {
    if request == nil {
        request = NewDescribeUserSqlAdviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserSqlAdvice require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserSqlAdviceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDiagDBInstanceConfRequest() (request *ModifyDiagDBInstanceConfRequest) {
    request = &ModifyDiagDBInstanceConfRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "ModifyDiagDBInstanceConf")
    
    
    return
}

func NewModifyDiagDBInstanceConfResponse() (response *ModifyDiagDBInstanceConfResponse) {
    response = &ModifyDiagDBInstanceConfResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDiagDBInstanceConf
// This API is used to enable/disable instance inspection.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyDiagDBInstanceConf(request *ModifyDiagDBInstanceConfRequest) (response *ModifyDiagDBInstanceConfResponse, err error) {
    return c.ModifyDiagDBInstanceConfWithContext(context.Background(), request)
}

// ModifyDiagDBInstanceConf
// This API is used to enable/disable instance inspection.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyDiagDBInstanceConfWithContext(ctx context.Context, request *ModifyDiagDBInstanceConfRequest) (response *ModifyDiagDBInstanceConfResponse, err error) {
    if request == nil {
        request = NewModifyDiagDBInstanceConfRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDiagDBInstanceConf require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDiagDBInstanceConfResponse()
    err = c.Send(request, response)
    return
}
