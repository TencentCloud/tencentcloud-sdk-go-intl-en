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
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CancelSparkSessionBatchSQL(request *CancelSparkSessionBatchSQLRequest) (response *CancelSparkSessionBatchSQLResponse, err error) {
    return c.CancelSparkSessionBatchSQLWithContext(context.Background(), request)
}

// CancelSparkSessionBatchSQL
// This API is used to cancel a Spark SQL batch task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CancelSparkSessionBatchSQLWithContext(ctx context.Context, request *CancelSparkSessionBatchSQLRequest) (response *CancelSparkSessionBatchSQLResponse, err error) {
    if request == nil {
        request = NewCancelSparkSessionBatchSQLRequest()
    }
    
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
//  INVALIDPARAMETER_INVALIDTASKID = "InvalidParameter.InvalidTaskId"
//  INVALIDPARAMETER_TASKALREADYFINISHED = "InvalidParameter.TaskAlreadyFinished"
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
//  INVALIDPARAMETER_INVALIDTASKID = "InvalidParameter.InvalidTaskId"
//  INVALIDPARAMETER_TASKALREADYFINISHED = "InvalidParameter.TaskAlreadyFinished"
func (c *Client) CancelTaskWithContext(ctx context.Context, request *CancelTaskRequest) (response *CancelTaskResponse, err error) {
    if request == nil {
        request = NewCancelTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelTaskResponse()
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
//  INVALIDPARAMETER_DUPLICATEDATAENGINENAME = "InvalidParameter.DuplicateDataEngineName"
//  INVALIDPARAMETER_INVALIDDATAENGINEMODE = "InvalidParameter.InvalidDataEngineMode"
//  INVALIDPARAMETER_INVALIDDATAENGINENAME = "InvalidParameter.InvalidDataEngineName"
//  INVALIDPARAMETER_INVALIDDATAENGINESPECS = "InvalidParameter.InvalidDataEngineSpecs"
//  INVALIDPARAMETER_INVALIDENGINETYPE = "InvalidParameter.InvalidEngineType"
//  INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPayMode"
//  INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_VPCCIDRFORMATERROR = "InvalidParameter.VpcCidrFormatError"
//  RESOURCENOTFOUND = "ResourceNotFound"
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
//  INVALIDPARAMETER_DUPLICATEDATAENGINENAME = "InvalidParameter.DuplicateDataEngineName"
//  INVALIDPARAMETER_INVALIDDATAENGINEMODE = "InvalidParameter.InvalidDataEngineMode"
//  INVALIDPARAMETER_INVALIDDATAENGINENAME = "InvalidParameter.InvalidDataEngineName"
//  INVALIDPARAMETER_INVALIDDATAENGINESPECS = "InvalidParameter.InvalidDataEngineSpecs"
//  INVALIDPARAMETER_INVALIDENGINETYPE = "InvalidParameter.InvalidEngineType"
//  INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPayMode"
//  INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"
//  INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"
//  INVALIDPARAMETER_VPCCIDRFORMATERROR = "InvalidParameter.VpcCidrFormatError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOPAYMENTAUTHORITY = "UnauthorizedOperation.NoPaymentAuthority"
func (c *Client) CreateDataEngineWithContext(ctx context.Context, request *CreateDataEngineRequest) (response *CreateDataEngineResponse, err error) {
    if request == nil {
        request = NewCreateDataEngineRequest()
    }
    
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
// This API is used to create a Spark application.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDROLEARN = "InvalidParameter.InvalidRoleArn"
func (c *Client) CreateSparkApp(request *CreateSparkAppRequest) (response *CreateSparkAppResponse, err error) {
    return c.CreateSparkAppWithContext(context.Background(), request)
}

// CreateSparkApp
// This API is used to create a Spark application.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDROLEARN = "InvalidParameter.InvalidRoleArn"
func (c *Client) CreateSparkAppWithContext(ctx context.Context, request *CreateSparkAppRequest) (response *CreateSparkAppResponse, err error) {
    if request == nil {
        request = NewCreateSparkAppRequest()
    }
    
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
// This API is used to create a Spark task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDSPARKAPPPARAM = "InvalidParameter.InvalidSparkAppParam"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_USECOMPUTINGENGINE = "UnauthorizedOperation.UseComputingEngine"
func (c *Client) CreateSparkAppTask(request *CreateSparkAppTaskRequest) (response *CreateSparkAppTaskResponse, err error) {
    return c.CreateSparkAppTaskWithContext(context.Background(), request)
}

// CreateSparkAppTask
// This API is used to create a Spark task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDSPARKAPPPARAM = "InvalidParameter.InvalidSparkAppParam"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_USECOMPUTINGENGINE = "UnauthorizedOperation.UseComputingEngine"
func (c *Client) CreateSparkAppTaskWithContext(ctx context.Context, request *CreateSparkAppTaskRequest) (response *CreateSparkAppTaskResponse, err error) {
    if request == nil {
        request = NewCreateSparkAppTaskRequest()
    }
    
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
// This API is used to submit a Spark SQL batch task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDCODE_SESSIONINSUFFICIENTRESOURCES = "ResourceNotFound.ResourceNotFoundCode_SessionInsufficientResources"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  RESOURCENOTFOUND_SESSIONSTATEDEAD = "ResourceNotFound.SessionStateDead"
func (c *Client) CreateSparkSessionBatchSQL(request *CreateSparkSessionBatchSQLRequest) (response *CreateSparkSessionBatchSQLResponse, err error) {
    return c.CreateSparkSessionBatchSQLWithContext(context.Background(), request)
}

// CreateSparkSessionBatchSQL
// This API is used to submit a Spark SQL batch task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDCODE_SESSIONINSUFFICIENTRESOURCES = "ResourceNotFound.ResourceNotFoundCode_SessionInsufficientResources"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  RESOURCENOTFOUND_SESSIONSTATEDEAD = "ResourceNotFound.SessionStateDead"
func (c *Client) CreateSparkSessionBatchSQLWithContext(ctx context.Context, request *CreateSparkSessionBatchSQLRequest) (response *CreateSparkSessionBatchSQLResponse, err error) {
    if request == nil {
        request = NewCreateSparkSessionBatchSQLRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSparkSessionBatchSQL require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSparkSessionBatchSQLResponse()
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
// This API is used to create a SQL query task. (We recommend you use the `CreateTasks` API instead.)
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAENGINENAME = "InvalidParameter.InvalidDataEngineName"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  INVALIDPARAMETER_INVALIDSQLNUM = "InvalidParameter.InvalidSQLNum"
//  INVALIDPARAMETER_INVALIDSTORELOCATION = "InvalidParameter.InvalidStoreLocation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESULTOUTPUTPATHNOTFOUND = "ResourceNotFound.ResultOutputPathNotFound"
//  RESOURCEUNAVAILABLE_BALANCEINSUFFICIENT = "ResourceUnavailable.BalanceInsufficient"
//  UNAUTHORIZEDOPERATION_USECOMPUTINGENGINE = "UnauthorizedOperation.UseComputingEngine"
func (c *Client) CreateTask(request *CreateTaskRequest) (response *CreateTaskResponse, err error) {
    return c.CreateTaskWithContext(context.Background(), request)
}

// CreateTask
// This API is used to create a SQL query task. (We recommend you use the `CreateTasks` API instead.)
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAENGINENAME = "InvalidParameter.InvalidDataEngineName"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  INVALIDPARAMETER_INVALIDSQLNUM = "InvalidParameter.InvalidSQLNum"
//  INVALIDPARAMETER_INVALIDSTORELOCATION = "InvalidParameter.InvalidStoreLocation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESULTOUTPUTPATHNOTFOUND = "ResourceNotFound.ResultOutputPathNotFound"
//  RESOURCEUNAVAILABLE_BALANCEINSUFFICIENT = "ResourceUnavailable.BalanceInsufficient"
//  UNAUTHORIZEDOPERATION_USECOMPUTINGENGINE = "UnauthorizedOperation.UseComputingEngine"
func (c *Client) CreateTaskWithContext(ctx context.Context, request *CreateTaskRequest) (response *CreateTaskResponse, err error) {
    if request == nil {
        request = NewCreateTaskRequest()
    }
    
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
// This API is used to create tasks in batches.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAENGINENAME = "InvalidParameter.InvalidDataEngineName"
//  INVALIDPARAMETER_INVALIDFAILURETOLERANCE = "InvalidParameter.InvalidFailureTolerance"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  INVALIDPARAMETER_INVALIDSQLNUM = "InvalidParameter.InvalidSQLNum"
//  INVALIDPARAMETER_INVALIDSTORELOCATION = "InvalidParameter.InvalidStoreLocation"
//  INVALIDPARAMETER_INVALIDTASKTYPE = "InvalidParameter.InvalidTaskType"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESULTOUTPUTPATHNOTFOUND = "ResourceNotFound.ResultOutputPathNotFound"
//  RESOURCEUNAVAILABLE_BALANCEINSUFFICIENT = "ResourceUnavailable.BalanceInsufficient"
//  UNAUTHORIZEDOPERATION_USECOMPUTINGENGINE = "UnauthorizedOperation.UseComputingEngine"
func (c *Client) CreateTasks(request *CreateTasksRequest) (response *CreateTasksResponse, err error) {
    return c.CreateTasksWithContext(context.Background(), request)
}

// CreateTasks
// This API is used to create tasks in batches.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAENGINENAME = "InvalidParameter.InvalidDataEngineName"
//  INVALIDPARAMETER_INVALIDFAILURETOLERANCE = "InvalidParameter.InvalidFailureTolerance"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  INVALIDPARAMETER_INVALIDSQLNUM = "InvalidParameter.InvalidSQLNum"
//  INVALIDPARAMETER_INVALIDSTORELOCATION = "InvalidParameter.InvalidStoreLocation"
//  INVALIDPARAMETER_INVALIDTASKTYPE = "InvalidParameter.InvalidTaskType"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESULTOUTPUTPATHNOTFOUND = "ResourceNotFound.ResultOutputPathNotFound"
//  RESOURCEUNAVAILABLE_BALANCEINSUFFICIENT = "ResourceUnavailable.BalanceInsufficient"
//  UNAUTHORIZEDOPERATION_USECOMPUTINGENGINE = "UnauthorizedOperation.UseComputingEngine"
func (c *Client) CreateTasksWithContext(ctx context.Context, request *CreateTasksRequest) (response *CreateTasksResponse, err error) {
    if request == nil {
        request = NewCreateTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTasksResponse()
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
// This API is used to delete a Spark application.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAENGINENAME = "InvalidParameter.InvalidDataEngineName"
//  INVALIDPARAMETER_INVALIDFAILURETOLERANCE = "InvalidParameter.InvalidFailureTolerance"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  INVALIDPARAMETER_INVALIDSQLNUM = "InvalidParameter.InvalidSQLNum"
//  INVALIDPARAMETER_INVALIDSTORELOCATION = "InvalidParameter.InvalidStoreLocation"
//  INVALIDPARAMETER_INVALIDTASKTYPE = "InvalidParameter.InvalidTaskType"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESULTOUTPUTPATHNOTFOUND = "ResourceNotFound.ResultOutputPathNotFound"
//  RESOURCEUNAVAILABLE_BALANCEINSUFFICIENT = "ResourceUnavailable.BalanceInsufficient"
//  UNAUTHORIZEDOPERATION_USECOMPUTINGENGINE = "UnauthorizedOperation.UseComputingEngine"
func (c *Client) DeleteSparkApp(request *DeleteSparkAppRequest) (response *DeleteSparkAppResponse, err error) {
    return c.DeleteSparkAppWithContext(context.Background(), request)
}

// DeleteSparkApp
// This API is used to delete a Spark application.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAENGINENAME = "InvalidParameter.InvalidDataEngineName"
//  INVALIDPARAMETER_INVALIDFAILURETOLERANCE = "InvalidParameter.InvalidFailureTolerance"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  INVALIDPARAMETER_INVALIDSQLNUM = "InvalidParameter.InvalidSQLNum"
//  INVALIDPARAMETER_INVALIDSTORELOCATION = "InvalidParameter.InvalidStoreLocation"
//  INVALIDPARAMETER_INVALIDTASKTYPE = "InvalidParameter.InvalidTaskType"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESULTOUTPUTPATHNOTFOUND = "ResourceNotFound.ResultOutputPathNotFound"
//  RESOURCEUNAVAILABLE_BALANCEINSUFFICIENT = "ResourceUnavailable.BalanceInsufficient"
//  UNAUTHORIZEDOPERATION_USECOMPUTINGENGINE = "UnauthorizedOperation.UseComputingEngine"
func (c *Client) DeleteSparkAppWithContext(ctx context.Context, request *DeleteSparkAppRequest) (response *DeleteSparkAppResponse, err error) {
    if request == nil {
        request = NewDeleteSparkAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSparkApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSparkAppResponse()
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
// This API is used to get the list of disabled table attributes.
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
// This API is used to get the list of disabled table attributes.
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeForbiddenTablePro require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeForbiddenTableProResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLakeFsInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLakeFsInfoResponse()
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
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeResultDownload(request *DescribeResultDownloadRequest) (response *DescribeResultDownloadResponse, err error) {
    return c.DescribeResultDownloadWithContext(context.Background(), request)
}

// DescribeResultDownload
// This API is used to get a query result download task.
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
func (c *Client) DescribeResultDownloadWithContext(ctx context.Context, request *DescribeResultDownloadRequest) (response *DescribeResultDownloadResponse, err error) {
    if request == nil {
        request = NewDescribeResultDownloadRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResultDownload require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResultDownloadResponse()
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
// This API is used to query a specific Spark application.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSPARKAPPPARAM = "InvalidParameter.InvalidSparkAppParam"
func (c *Client) DescribeSparkAppJob(request *DescribeSparkAppJobRequest) (response *DescribeSparkAppJobResponse, err error) {
    return c.DescribeSparkAppJobWithContext(context.Background(), request)
}

// DescribeSparkAppJob
// This API is used to query a specific Spark application.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSPARKAPPPARAM = "InvalidParameter.InvalidSparkAppParam"
func (c *Client) DescribeSparkAppJobWithContext(ctx context.Context, request *DescribeSparkAppJobRequest) (response *DescribeSparkAppJobResponse, err error) {
    if request == nil {
        request = NewDescribeSparkAppJobRequest()
    }
    
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
// This API is used to get the list of Spark applications.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSparkAppJobs(request *DescribeSparkAppJobsRequest) (response *DescribeSparkAppJobsResponse, err error) {
    return c.DescribeSparkAppJobsWithContext(context.Background(), request)
}

// DescribeSparkAppJobs
// This API is used to get the list of Spark applications.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSparkAppJobsWithContext(ctx context.Context, request *DescribeSparkAppJobsRequest) (response *DescribeSparkAppJobsResponse, err error) {
    if request == nil {
        request = NewDescribeSparkAppJobsRequest()
    }
    
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
// This API is used to query the list of running task instances of a Spark application.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeSparkAppTasks(request *DescribeSparkAppTasksRequest) (response *DescribeSparkAppTasksResponse, err error) {
    return c.DescribeSparkAppTasksWithContext(context.Background(), request)
}

// DescribeSparkAppTasks
// This API is used to query the list of running task instances of a Spark application.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeSparkAppTasksWithContext(ctx context.Context, request *DescribeSparkAppTasksRequest) (response *DescribeSparkAppTasksResponse, err error) {
    if request == nil {
        request = NewDescribeSparkAppTasksRequest()
    }
    
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
// This API is used to obtain the logs of a Spark SQL batch task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSparkSessionBatchSqlLog(request *DescribeSparkSessionBatchSqlLogRequest) (response *DescribeSparkSessionBatchSqlLogResponse, err error) {
    return c.DescribeSparkSessionBatchSqlLogWithContext(context.Background(), request)
}

// DescribeSparkSessionBatchSqlLog
// This API is used to obtain the logs of a Spark SQL batch task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSparkSessionBatchSqlLogWithContext(ctx context.Context, request *DescribeSparkSessionBatchSqlLogRequest) (response *DescribeSparkSessionBatchSqlLogResponse, err error) {
    if request == nil {
        request = NewDescribeSparkSessionBatchSqlLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSparkSessionBatchSqlLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSparkSessionBatchSqlLogResponse()
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
//  INVALIDPARAMETER_INVALIDTASKID = "InvalidParameter.InvalidTaskId"
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
//  INVALIDPARAMETER_INVALIDTASKID = "InvalidParameter.InvalidTaskId"
func (c *Client) DescribeTaskResultWithContext(ctx context.Context, request *DescribeTaskResultRequest) (response *DescribeTaskResultResponse, err error) {
    if request == nil {
        request = NewDescribeTaskResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskResultResponse()
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
//  INVALIDPARAMETER = "InvalidParameter"
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
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeTasksWithContext(ctx context.Context, request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    if request == nil {
        request = NewDescribeTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTasksResponse()
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
//  INVALIDPARAMETER = "InvalidParameter"
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
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GenerateCreateMangedTableSqlWithContext(ctx context.Context, request *GenerateCreateMangedTableSqlRequest) (response *GenerateCreateMangedTableSqlResponse, err error) {
    if request == nil {
        request = NewGenerateCreateMangedTableSqlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GenerateCreateMangedTableSql require credential")
    }

    request.SetContext(ctx)
    
    response = NewGenerateCreateMangedTableSqlResponse()
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
// This API is used to update a Spark application.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDDATAENGINENAME = "InvalidParameter.InvalidDataEngineName"
func (c *Client) ModifySparkApp(request *ModifySparkAppRequest) (response *ModifySparkAppResponse, err error) {
    return c.ModifySparkAppWithContext(context.Background(), request)
}

// ModifySparkApp
// This API is used to update a Spark application.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDDATAENGINENAME = "InvalidParameter.InvalidDataEngineName"
func (c *Client) ModifySparkAppWithContext(ctx context.Context, request *ModifySparkAppRequest) (response *ModifySparkAppResponse, err error) {
    if request == nil {
        request = NewModifySparkAppRequest()
    }
    
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
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifySparkAppBatch(request *ModifySparkAppBatchRequest) (response *ModifySparkAppBatchResponse, err error) {
    return c.ModifySparkAppBatchWithContext(context.Background(), request)
}

// ModifySparkAppBatch
// This API is used to modify Spark job parameters in batches.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifySparkAppBatchWithContext(ctx context.Context, request *ModifySparkAppBatchRequest) (response *ModifySparkAppBatchResponse, err error) {
    if request == nil {
        request = NewModifySparkAppBatchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySparkAppBatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySparkAppBatchResponse()
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
// This API is used to suspend or resume a data engine.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  UNAUTHORIZEDOPERATION_OPERATECOMPUTINGENGINE = "UnauthorizedOperation.OperateComputingEngine"
func (c *Client) SuspendResumeDataEngine(request *SuspendResumeDataEngineRequest) (response *SuspendResumeDataEngineResponse, err error) {
    return c.SuspendResumeDataEngineWithContext(context.Background(), request)
}

// SuspendResumeDataEngine
// This API is used to suspend or resume a data engine.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  UNAUTHORIZEDOPERATION_OPERATECOMPUTINGENGINE = "UnauthorizedOperation.OperateComputingEngine"
func (c *Client) SuspendResumeDataEngineWithContext(ctx context.Context, request *SuspendResumeDataEngineRequest) (response *SuspendResumeDataEngineResponse, err error) {
    if request == nil {
        request = NewSuspendResumeDataEngineRequest()
    }
    
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
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
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
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  UNAUTHORIZEDOPERATION_OPERATECOMPUTINGENGINE = "UnauthorizedOperation.OperateComputingEngine"
func (c *Client) SwitchDataEngineWithContext(ctx context.Context, request *SwitchDataEngineRequest) (response *SwitchDataEngineResponse, err error) {
    if request == nil {
        request = NewSwitchDataEngineRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SwitchDataEngine require credential")
    }

    request.SetContext(ctx)
    
    response = NewSwitchDataEngineResponse()
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  UNAUTHORIZEDOPERATION_OPERATECOMPUTINGENGINE = "UnauthorizedOperation.OperateComputingEngine"
func (c *Client) UpdateRowFilter(request *UpdateRowFilterRequest) (response *UpdateRowFilterResponse, err error) {
    return c.UpdateRowFilterWithContext(context.Background(), request)
}

// UpdateRowFilter
// This API is used to update row filters. Please note that it updates filters only but not catalogs, databases, or tables.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"
//  UNAUTHORIZEDOPERATION_OPERATECOMPUTINGENGINE = "UnauthorizedOperation.OperateComputingEngine"
func (c *Client) UpdateRowFilterWithContext(ctx context.Context, request *UpdateRowFilterRequest) (response *UpdateRowFilterResponse, err error) {
    if request == nil {
        request = NewUpdateRowFilterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateRowFilter require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateRowFilterResponse()
    err = c.Send(request, response)
    return
}
