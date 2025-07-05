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

package v20191205

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2019-12-05"

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


func NewApplyCertificateRequest() (request *ApplyCertificateRequest) {
    request = &ApplyCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "ApplyCertificate")
    
    
    return
}

func NewApplyCertificateResponse() (response *ApplyCertificateResponse) {
    response = &ApplyCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ApplyCertificate
// This API is used to apply for a free certificate.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_MAINDOMAINCERTIFICATECOUNTLIMIT = "FailedOperation.MainDomainCertificateCountLimit"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  FAILEDOPERATION_PACKAGECOUNTLIMIT = "FailedOperation.PackageCountLimit"
//  FAILEDOPERATION_PACKAGEEXPIRED = "FailedOperation.PackageExpired"
//  FAILEDOPERATION_PACKAGENOTFOUND = "FailedOperation.PackageNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PACKAGEIDSINVALID = "InvalidParameter.PackageIdsInvalid"
func (c *Client) ApplyCertificate(request *ApplyCertificateRequest) (response *ApplyCertificateResponse, err error) {
    return c.ApplyCertificateWithContext(context.Background(), request)
}

// ApplyCertificate
// This API is used to apply for a free certificate.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_MAINDOMAINCERTIFICATECOUNTLIMIT = "FailedOperation.MainDomainCertificateCountLimit"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  FAILEDOPERATION_PACKAGECOUNTLIMIT = "FailedOperation.PackageCountLimit"
//  FAILEDOPERATION_PACKAGEEXPIRED = "FailedOperation.PackageExpired"
//  FAILEDOPERATION_PACKAGENOTFOUND = "FailedOperation.PackageNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PACKAGEIDSINVALID = "InvalidParameter.PackageIdsInvalid"
func (c *Client) ApplyCertificateWithContext(ctx context.Context, request *ApplyCertificateRequest) (response *ApplyCertificateResponse, err error) {
    if request == nil {
        request = NewApplyCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewBatchDeleteCSRRequest() (request *BatchDeleteCSRRequest) {
    request = &BatchDeleteCSRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "BatchDeleteCSR")
    
    
    return
}

func NewBatchDeleteCSRResponse() (response *BatchDeleteCSRResponse) {
    response = &BatchDeleteCSRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchDeleteCSR
// This API is used to batch delete CSRs.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCSRID = "InvalidParameter.InvalidCSRId"
//  INVALIDPARAMETER_WITHDETAILREASON = "InvalidParameter.WithDetailReason"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) BatchDeleteCSR(request *BatchDeleteCSRRequest) (response *BatchDeleteCSRResponse, err error) {
    return c.BatchDeleteCSRWithContext(context.Background(), request)
}

// BatchDeleteCSR
// This API is used to batch delete CSRs.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCSRID = "InvalidParameter.InvalidCSRId"
//  INVALIDPARAMETER_WITHDETAILREASON = "InvalidParameter.WithDetailReason"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) BatchDeleteCSRWithContext(ctx context.Context, request *BatchDeleteCSRRequest) (response *BatchDeleteCSRResponse, err error) {
    if request == nil {
        request = NewBatchDeleteCSRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchDeleteCSR require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchDeleteCSRResponse()
    err = c.Send(request, response)
    return
}

func NewCancelAuditCertificateRequest() (request *CancelAuditCertificateRequest) {
    request = &CancelAuditCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "CancelAuditCertificate")
    
    
    return
}

func NewCancelAuditCertificateResponse() (response *CancelAuditCertificateResponse) {
    response = &CancelAuditCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CancelAuditCertificate
// This API is used to cancel certificate review.
//
// error code that may be returned:
//  FAILEDOPERATION_CERTIFICATENOTFOUNDORCANTCANCEL = "FailedOperation.CertificateNotFoundOrCantCancel"
func (c *Client) CancelAuditCertificate(request *CancelAuditCertificateRequest) (response *CancelAuditCertificateResponse, err error) {
    return c.CancelAuditCertificateWithContext(context.Background(), request)
}

// CancelAuditCertificate
// This API is used to cancel certificate review.
//
// error code that may be returned:
//  FAILEDOPERATION_CERTIFICATENOTFOUNDORCANTCANCEL = "FailedOperation.CertificateNotFoundOrCantCancel"
func (c *Client) CancelAuditCertificateWithContext(ctx context.Context, request *CancelAuditCertificateRequest) (response *CancelAuditCertificateResponse, err error) {
    if request == nil {
        request = NewCancelAuditCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelAuditCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelAuditCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewCancelCertificateOrderRequest() (request *CancelCertificateOrderRequest) {
    request = &CancelCertificateOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "CancelCertificateOrder")
    
    
    return
}

func NewCancelCertificateOrderResponse() (response *CancelCertificateOrderResponse) {
    response = &CancelCertificateOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CancelCertificateOrder
// This API is used to cancel a certificate order.
//
// error code that may be returned:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
func (c *Client) CancelCertificateOrder(request *CancelCertificateOrderRequest) (response *CancelCertificateOrderResponse, err error) {
    return c.CancelCertificateOrderWithContext(context.Background(), request)
}

// CancelCertificateOrder
// This API is used to cancel a certificate order.
//
// error code that may be returned:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
func (c *Client) CancelCertificateOrderWithContext(ctx context.Context, request *CancelCertificateOrderRequest) (response *CancelCertificateOrderResponse, err error) {
    if request == nil {
        request = NewCancelCertificateOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelCertificateOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelCertificateOrderResponse()
    err = c.Send(request, response)
    return
}

func NewCommitCertificateInformationRequest() (request *CommitCertificateInformationRequest) {
    request = &CommitCertificateInformationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "CommitCertificateInformation")
    
    
    return
}

func NewCommitCertificateInformationResponse() (response *CommitCertificateInformationResponse) {
    response = &CommitCertificateInformationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CommitCertificateInformation
// Submit payment certificate orders; This API does not maintain new features, and you can use the new API to submit orders. [CertificateOrderSubmit](https://intl.cloud.tencent.com/document/product/400/116032?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDRESPONSEERROR = "InternalError.BackendResponseError"
func (c *Client) CommitCertificateInformation(request *CommitCertificateInformationRequest) (response *CommitCertificateInformationResponse, err error) {
    return c.CommitCertificateInformationWithContext(context.Background(), request)
}

// CommitCertificateInformation
// Submit payment certificate orders; This API does not maintain new features, and you can use the new API to submit orders. [CertificateOrderSubmit](https://intl.cloud.tencent.com/document/product/400/116032?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDRESPONSEERROR = "InternalError.BackendResponseError"
func (c *Client) CommitCertificateInformationWithContext(ctx context.Context, request *CommitCertificateInformationRequest) (response *CommitCertificateInformationResponse, err error) {
    if request == nil {
        request = NewCommitCertificateInformationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CommitCertificateInformation require credential")
    }

    request.SetContext(ctx)
    
    response = NewCommitCertificateInformationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCSRRequest() (request *CreateCSRRequest) {
    request = &CreateCSRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "CreateCSR")
    
    
    return
}

func NewCreateCSRResponse() (response *CreateCSRResponse) {
    response = &CreateCSRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCSR
// This API is used to create a CSR.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CAMAUTHORIZEDFAIL = "FailedOperation.CAMAuthorizedFail"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
func (c *Client) CreateCSR(request *CreateCSRRequest) (response *CreateCSRResponse, err error) {
    return c.CreateCSRWithContext(context.Background(), request)
}

// CreateCSR
// This API is used to create a CSR.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CAMAUTHORIZEDFAIL = "FailedOperation.CAMAuthorizedFail"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
func (c *Client) CreateCSRWithContext(ctx context.Context, request *CreateCSRRequest) (response *CreateCSRResponse, err error) {
    if request == nil {
        request = NewCreateCSRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCSR require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCSRResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCertificateRequest() (request *CreateCertificateRequest) {
    request = &CreateCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "CreateCertificate")
    
    
    return
}

func NewCreateCertificateResponse() (response *CreateCertificateResponse) {
    response = &CreateCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCertificate
// This API is used to purchase a certificate.
//
// error code that may be returned:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDRESPONSEERROR = "InternalError.BackendResponseError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateCertificate(request *CreateCertificateRequest) (response *CreateCertificateResponse, err error) {
    return c.CreateCertificateWithContext(context.Background(), request)
}

// CreateCertificate
// This API is used to purchase a certificate.
//
// error code that may be returned:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDRESPONSEERROR = "InternalError.BackendResponseError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateCertificateWithContext(ctx context.Context, request *CreateCertificateRequest) (response *CreateCertificateResponse, err error) {
    if request == nil {
        request = NewCreateCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCertificateBindResourceSyncTaskRequest() (request *CreateCertificateBindResourceSyncTaskRequest) {
    request = &CreateCertificateBindResourceSyncTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "CreateCertificateBindResourceSyncTask")
    
    
    return
}

func NewCreateCertificateBindResourceSyncTaskResponse() (response *CreateCertificateBindResourceSyncTaskResponse) {
    response = &CreateCertificateBindResourceSyncTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCertificateBindResourceSyncTask
// This API is used to create an async task for querying the cloud resources associated with a certificate. If such a task already exists under the certificate ID, the ID of this task is returned as the result. The following types of cloud resources are supported: CLB, CDN, WAF, LIVE, VOD, DDOS, TKE, APIGATEWAY, TCB, and TEO (EDGEONE). You can query the result of this task using the `DescribeCertificateBindResourceTaskResult` API.
//
// error code that may be returned:
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
func (c *Client) CreateCertificateBindResourceSyncTask(request *CreateCertificateBindResourceSyncTaskRequest) (response *CreateCertificateBindResourceSyncTaskResponse, err error) {
    return c.CreateCertificateBindResourceSyncTaskWithContext(context.Background(), request)
}

// CreateCertificateBindResourceSyncTask
// This API is used to create an async task for querying the cloud resources associated with a certificate. If such a task already exists under the certificate ID, the ID of this task is returned as the result. The following types of cloud resources are supported: CLB, CDN, WAF, LIVE, VOD, DDOS, TKE, APIGATEWAY, TCB, and TEO (EDGEONE). You can query the result of this task using the `DescribeCertificateBindResourceTaskResult` API.
//
// error code that may be returned:
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
func (c *Client) CreateCertificateBindResourceSyncTaskWithContext(ctx context.Context, request *CreateCertificateBindResourceSyncTaskRequest) (response *CreateCertificateBindResourceSyncTaskResponse, err error) {
    if request == nil {
        request = NewCreateCertificateBindResourceSyncTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCertificateBindResourceSyncTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCertificateBindResourceSyncTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCertificateRequest() (request *DeleteCertificateRequest) {
    request = &DeleteCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DeleteCertificate")
    
    
    return
}

func NewDeleteCertificateResponse() (response *DeleteCertificateResponse) {
    response = &DeleteCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCertificate
// This API is used to delete a certificate.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_BOUNDRESOURCES = "FailedOperation.BoundResources"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_DELETEFAILEDTIMENOTUP = "FailedOperation.DeleteFailedTimeNotUp"
//  FAILEDOPERATION_DELETERESOURCEFAILED = "FailedOperation.DeleteResourceFailed"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_RENEWNOTISSUED = "FailedOperation.RenewNotIssued"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DeleteCertificate(request *DeleteCertificateRequest) (response *DeleteCertificateResponse, err error) {
    return c.DeleteCertificateWithContext(context.Background(), request)
}

// DeleteCertificate
// This API is used to delete a certificate.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_BOUNDRESOURCES = "FailedOperation.BoundResources"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_DELETEFAILEDTIMENOTUP = "FailedOperation.DeleteFailedTimeNotUp"
//  FAILEDOPERATION_DELETERESOURCEFAILED = "FailedOperation.DeleteResourceFailed"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_RENEWNOTISSUED = "FailedOperation.RenewNotIssued"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DeleteCertificateWithContext(ctx context.Context, request *DeleteCertificateRequest) (response *DeleteCertificateResponse, err error) {
    if request == nil {
        request = NewDeleteCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCSRRequest() (request *DescribeCSRRequest) {
    request = &DescribeCSRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeCSR")
    
    
    return
}

func NewDescribeCSRResponse() (response *DescribeCSRResponse) {
    response = &DescribeCSRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCSR
// This API is used to query the details of a CSR.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_INVALIDCSRID = "InvalidParameter.InvalidCSRId"
//  INVALIDPARAMETER_WITHDETAILREASON = "InvalidParameter.WithDetailReason"
func (c *Client) DescribeCSR(request *DescribeCSRRequest) (response *DescribeCSRResponse, err error) {
    return c.DescribeCSRWithContext(context.Background(), request)
}

// DescribeCSR
// This API is used to query the details of a CSR.
//
// error code that may be returned:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_INVALIDCSRID = "InvalidParameter.InvalidCSRId"
//  INVALIDPARAMETER_WITHDETAILREASON = "InvalidParameter.WithDetailReason"
func (c *Client) DescribeCSRWithContext(ctx context.Context, request *DescribeCSRRequest) (response *DescribeCSRResponse, err error) {
    if request == nil {
        request = NewDescribeCSRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCSR require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCSRResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCSRSetRequest() (request *DescribeCSRSetRequest) {
    request = &DescribeCSRSetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeCSRSet")
    
    
    return
}

func NewDescribeCSRSetResponse() (response *DescribeCSRSetResponse) {
    response = &DescribeCSRSetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCSRSet
// This API is used to query the CSR list.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMAUTHORIZEDFAIL = "FailedOperation.CAMAuthorizedFail"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
func (c *Client) DescribeCSRSet(request *DescribeCSRSetRequest) (response *DescribeCSRSetResponse, err error) {
    return c.DescribeCSRSetWithContext(context.Background(), request)
}

// DescribeCSRSet
// This API is used to query the CSR list.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMAUTHORIZEDFAIL = "FailedOperation.CAMAuthorizedFail"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
func (c *Client) DescribeCSRSetWithContext(ctx context.Context, request *DescribeCSRSetRequest) (response *DescribeCSRSetResponse, err error) {
    if request == nil {
        request = NewDescribeCSRSetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCSRSet require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCSRSetResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCertificateRequest() (request *DescribeCertificateRequest) {
    request = &DescribeCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeCertificate")
    
    
    return
}

func NewDescribeCertificateResponse() (response *DescribeCertificateResponse) {
    response = &DescribeCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCertificate
// This API is used to get certificate information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeCertificate(request *DescribeCertificateRequest) (response *DescribeCertificateResponse, err error) {
    return c.DescribeCertificateWithContext(context.Background(), request)
}

// DescribeCertificate
// This API is used to get certificate information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeCertificateWithContext(ctx context.Context, request *DescribeCertificateRequest) (response *DescribeCertificateResponse, err error) {
    if request == nil {
        request = NewDescribeCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCertificateBindResourceTaskDetailRequest() (request *DescribeCertificateBindResourceTaskDetailRequest) {
    request = &DescribeCertificateBindResourceTaskDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeCertificateBindResourceTaskDetail")
    
    
    return
}

func NewDescribeCertificateBindResourceTaskDetailResponse() (response *DescribeCertificateBindResourceTaskDetailResponse) {
    response = &DescribeCertificateBindResourceTaskDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCertificateBindResourceTaskDetail
// This API is used to query the task result of CreateCertificateBindResourceSyncTask, returning the asynchronous task result of the certificate associated with cloud resources, supporting the following cloud resources: clb, cdn, waf, live, vod, ddos, tke, apigateway, tcb, teo (edgeOne), cos.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATESYNCTASKIDINVALID = "FailedOperation.CertificateSyncTaskIdInvalid"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
func (c *Client) DescribeCertificateBindResourceTaskDetail(request *DescribeCertificateBindResourceTaskDetailRequest) (response *DescribeCertificateBindResourceTaskDetailResponse, err error) {
    return c.DescribeCertificateBindResourceTaskDetailWithContext(context.Background(), request)
}

// DescribeCertificateBindResourceTaskDetail
// This API is used to query the task result of CreateCertificateBindResourceSyncTask, returning the asynchronous task result of the certificate associated with cloud resources, supporting the following cloud resources: clb, cdn, waf, live, vod, ddos, tke, apigateway, tcb, teo (edgeOne), cos.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATESYNCTASKIDINVALID = "FailedOperation.CertificateSyncTaskIdInvalid"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
func (c *Client) DescribeCertificateBindResourceTaskDetailWithContext(ctx context.Context, request *DescribeCertificateBindResourceTaskDetailRequest) (response *DescribeCertificateBindResourceTaskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCertificateBindResourceTaskDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCertificateBindResourceTaskDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCertificateBindResourceTaskDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCertificateBindResourceTaskResultRequest() (request *DescribeCertificateBindResourceTaskResultRequest) {
    request = &DescribeCertificateBindResourceTaskResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeCertificateBindResourceTaskResult")
    
    
    return
}

func NewDescribeCertificateBindResourceTaskResultResponse() (response *DescribeCertificateBindResourceTaskResultResponse) {
    response = &DescribeCertificateBindResourceTaskResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCertificateBindResourceTaskResult
// This API is used to query the result of an async task created with `CreateCertificateBindResourceSyncTask` to query cloud resources associated with a certificate. The following types of cloud resources are supported: CLB, CDN, WAF, LIVE, VOD, DDOS, TKE, APIGATEWAY, TCB, and TEO (EDGEONE).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATESYNCTASKIDINVALID = "FailedOperation.CertificateSyncTaskIdInvalid"
func (c *Client) DescribeCertificateBindResourceTaskResult(request *DescribeCertificateBindResourceTaskResultRequest) (response *DescribeCertificateBindResourceTaskResultResponse, err error) {
    return c.DescribeCertificateBindResourceTaskResultWithContext(context.Background(), request)
}

// DescribeCertificateBindResourceTaskResult
// This API is used to query the result of an async task created with `CreateCertificateBindResourceSyncTask` to query cloud resources associated with a certificate. The following types of cloud resources are supported: CLB, CDN, WAF, LIVE, VOD, DDOS, TKE, APIGATEWAY, TCB, and TEO (EDGEONE).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATESYNCTASKIDINVALID = "FailedOperation.CertificateSyncTaskIdInvalid"
func (c *Client) DescribeCertificateBindResourceTaskResultWithContext(ctx context.Context, request *DescribeCertificateBindResourceTaskResultRequest) (response *DescribeCertificateBindResourceTaskResultResponse, err error) {
    if request == nil {
        request = NewDescribeCertificateBindResourceTaskResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCertificateBindResourceTaskResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCertificateBindResourceTaskResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCertificateDetailRequest() (request *DescribeCertificateDetailRequest) {
    request = &DescribeCertificateDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeCertificateDetail")
    
    
    return
}

func NewDescribeCertificateDetailResponse() (response *DescribeCertificateDetailResponse) {
    response = &DescribeCertificateDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCertificateDetail
// This API is used to get certificate details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDRESPONSEEMPTY = "InternalError.BackendResponseEmpty"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
//  RESOURCENOTFOUND_MANAGER = "ResourceNotFound.Manager"
func (c *Client) DescribeCertificateDetail(request *DescribeCertificateDetailRequest) (response *DescribeCertificateDetailResponse, err error) {
    return c.DescribeCertificateDetailWithContext(context.Background(), request)
}

// DescribeCertificateDetail
// This API is used to get certificate details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDRESPONSEEMPTY = "InternalError.BackendResponseEmpty"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
//  RESOURCENOTFOUND_MANAGER = "ResourceNotFound.Manager"
func (c *Client) DescribeCertificateDetailWithContext(ctx context.Context, request *DescribeCertificateDetailRequest) (response *DescribeCertificateDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCertificateDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCertificateDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCertificateDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCertificateOperateLogsRequest() (request *DescribeCertificateOperateLogsRequest) {
    request = &DescribeCertificateOperateLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeCertificateOperateLogs")
    
    
    return
}

func NewDescribeCertificateOperateLogsResponse() (response *DescribeCertificateOperateLogsResponse) {
    response = &DescribeCertificateOperateLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCertificateOperateLogs
// This API is used to get certificate operation logs in the current account.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeCertificateOperateLogs(request *DescribeCertificateOperateLogsRequest) (response *DescribeCertificateOperateLogsResponse, err error) {
    return c.DescribeCertificateOperateLogsWithContext(context.Background(), request)
}

// DescribeCertificateOperateLogs
// This API is used to get certificate operation logs in the current account.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeCertificateOperateLogsWithContext(ctx context.Context, request *DescribeCertificateOperateLogsRequest) (response *DescribeCertificateOperateLogsResponse, err error) {
    if request == nil {
        request = NewDescribeCertificateOperateLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCertificateOperateLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCertificateOperateLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCertificatesRequest() (request *DescribeCertificatesRequest) {
    request = &DescribeCertificatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeCertificates")
    
    
    return
}

func NewDescribeCertificatesResponse() (response *DescribeCertificatesResponse) {
    response = &DescribeCertificatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCertificates
// This API is used to get the certificate list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CAMAUTHORIZEDFAIL = "FailedOperation.CAMAuthorizedFail"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDRESPONSEEMPTY = "InternalError.BackendResponseEmpty"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeCertificates(request *DescribeCertificatesRequest) (response *DescribeCertificatesResponse, err error) {
    return c.DescribeCertificatesWithContext(context.Background(), request)
}

// DescribeCertificates
// This API is used to get the certificate list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CAMAUTHORIZEDFAIL = "FailedOperation.CAMAuthorizedFail"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDRESPONSEEMPTY = "InternalError.BackendResponseEmpty"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeCertificatesWithContext(ctx context.Context, request *DescribeCertificatesRequest) (response *DescribeCertificatesResponse, err error) {
    if request == nil {
        request = NewDescribeCertificatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCertificates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCertificatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostTeoInstanceListRequest() (request *DescribeHostTeoInstanceListRequest) {
    request = &DescribeHostTeoInstanceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeHostTeoInstanceList")
    
    
    return
}

func NewDescribeHostTeoInstanceListResponse() (response *DescribeHostTeoInstanceListResponse) {
    response = &DescribeHostTeoInstanceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHostTeoInstanceList
// This API is used to query the list of EdgeOne instances to which a certificate can be deployed.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostTeoInstanceList(request *DescribeHostTeoInstanceListRequest) (response *DescribeHostTeoInstanceListResponse, err error) {
    return c.DescribeHostTeoInstanceListWithContext(context.Background(), request)
}

// DescribeHostTeoInstanceList
// This API is used to query the list of EdgeOne instances to which a certificate can be deployed.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostTeoInstanceListWithContext(ctx context.Context, request *DescribeHostTeoInstanceListRequest) (response *DescribeHostTeoInstanceListResponse, err error) {
    if request == nil {
        request = NewDescribeHostTeoInstanceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHostTeoInstanceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostTeoInstanceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostUpdateRecordRequest() (request *DescribeHostUpdateRecordRequest) {
    request = &DescribeHostUpdateRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeHostUpdateRecord")
    
    
    return
}

func NewDescribeHostUpdateRecordResponse() (response *DescribeHostUpdateRecordResponse) {
    response = &DescribeHostUpdateRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHostUpdateRecord
// Query certificate cloud resource update record list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostUpdateRecord(request *DescribeHostUpdateRecordRequest) (response *DescribeHostUpdateRecordResponse, err error) {
    return c.DescribeHostUpdateRecordWithContext(context.Background(), request)
}

// DescribeHostUpdateRecord
// Query certificate cloud resource update record list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostUpdateRecordWithContext(ctx context.Context, request *DescribeHostUpdateRecordRequest) (response *DescribeHostUpdateRecordResponse, err error) {
    if request == nil {
        request = NewDescribeHostUpdateRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHostUpdateRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostUpdateRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostUpdateRecordDetailRequest() (request *DescribeHostUpdateRecordDetailRequest) {
    request = &DescribeHostUpdateRecordDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeHostUpdateRecordDetail")
    
    
    return
}

func NewDescribeHostUpdateRecordDetailResponse() (response *DescribeHostUpdateRecordDetailResponse) {
    response = &DescribeHostUpdateRecordDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHostUpdateRecordDetail
// This API is used to query the update record details of certificate cloud resources.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEDEPLOYNOTEXIST = "FailedOperation.CertificateDeployNotExist"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostUpdateRecordDetail(request *DescribeHostUpdateRecordDetailRequest) (response *DescribeHostUpdateRecordDetailResponse, err error) {
    return c.DescribeHostUpdateRecordDetailWithContext(context.Background(), request)
}

// DescribeHostUpdateRecordDetail
// This API is used to query the update record details of certificate cloud resources.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEDEPLOYNOTEXIST = "FailedOperation.CertificateDeployNotExist"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostUpdateRecordDetailWithContext(ctx context.Context, request *DescribeHostUpdateRecordDetailRequest) (response *DescribeHostUpdateRecordDetailResponse, err error) {
    if request == nil {
        request = NewDescribeHostUpdateRecordDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHostUpdateRecordDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostUpdateRecordDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostUploadUpdateRecordRequest() (request *DescribeHostUploadUpdateRecordRequest) {
    request = &DescribeHostUploadUpdateRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeHostUploadUpdateRecord")
    
    
    return
}

func NewDescribeHostUploadUpdateRecordResponse() (response *DescribeHostUploadUpdateRecordResponse) {
    response = &DescribeHostUploadUpdateRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHostUploadUpdateRecord
// This API is used to query the record list of cloud resource updates for certificates (certificate ID unchanged).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostUploadUpdateRecord(request *DescribeHostUploadUpdateRecordRequest) (response *DescribeHostUploadUpdateRecordResponse, err error) {
    return c.DescribeHostUploadUpdateRecordWithContext(context.Background(), request)
}

// DescribeHostUploadUpdateRecord
// This API is used to query the record list of cloud resource updates for certificates (certificate ID unchanged).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeHostUploadUpdateRecordWithContext(ctx context.Context, request *DescribeHostUploadUpdateRecordRequest) (response *DescribeHostUploadUpdateRecordResponse, err error) {
    if request == nil {
        request = NewDescribeHostUploadUpdateRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHostUploadUpdateRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostUploadUpdateRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostUploadUpdateRecordDetailRequest() (request *DescribeHostUploadUpdateRecordDetailRequest) {
    request = &DescribeHostUploadUpdateRecordDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeHostUploadUpdateRecordDetail")
    
    
    return
}

func NewDescribeHostUploadUpdateRecordDetailResponse() (response *DescribeHostUploadUpdateRecordDetailResponse) {
    response = &DescribeHostUploadUpdateRecordDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHostUploadUpdateRecordDetail
// This API is used to query the deployment record details of certificate update records (certificate ID unchanged).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeHostUploadUpdateRecordDetail(request *DescribeHostUploadUpdateRecordDetailRequest) (response *DescribeHostUploadUpdateRecordDetailResponse, err error) {
    return c.DescribeHostUploadUpdateRecordDetailWithContext(context.Background(), request)
}

// DescribeHostUploadUpdateRecordDetail
// This API is used to query the deployment record details of certificate update records (certificate ID unchanged).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeHostUploadUpdateRecordDetailWithContext(ctx context.Context, request *DescribeHostUploadUpdateRecordDetailRequest) (response *DescribeHostUploadUpdateRecordDetailResponse, err error) {
    if request == nil {
        request = NewDescribeHostUploadUpdateRecordDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHostUploadUpdateRecordDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostUploadUpdateRecordDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadCertificateRequest() (request *DownloadCertificateRequest) {
    request = &DownloadCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DownloadCertificate")
    
    
    return
}

func NewDownloadCertificateResponse() (response *DownloadCertificateResponse) {
    response = &DownloadCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DownloadCertificate
// This API is used to download a certificate.
//
// error code that may be returned:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  INTERNALERROR = "InternalError"
func (c *Client) DownloadCertificate(request *DownloadCertificateRequest) (response *DownloadCertificateResponse, err error) {
    return c.DownloadCertificateWithContext(context.Background(), request)
}

// DownloadCertificate
// This API is used to download a certificate.
//
// error code that may be returned:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  INTERNALERROR = "InternalError"
func (c *Client) DownloadCertificateWithContext(ctx context.Context, request *DownloadCertificateRequest) (response *DownloadCertificateResponse, err error) {
    if request == nil {
        request = NewDownloadCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DownloadCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDownloadCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCSRRequest() (request *ModifyCSRRequest) {
    request = &ModifyCSRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "ModifyCSR")
    
    
    return
}

func NewModifyCSRResponse() (response *ModifyCSRResponse) {
    response = &ModifyCSRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCSR
// This API is used to modify the information of a CSR.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  INVALIDPARAMETER_INVALIDCSRID = "InvalidParameter.InvalidCSRId"
//  INVALIDPARAMETER_WITHDETAILREASON = "InvalidParameter.WithDetailReason"
func (c *Client) ModifyCSR(request *ModifyCSRRequest) (response *ModifyCSRResponse, err error) {
    return c.ModifyCSRWithContext(context.Background(), request)
}

// ModifyCSR
// This API is used to modify the information of a CSR.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  INVALIDPARAMETER_INVALIDCSRID = "InvalidParameter.InvalidCSRId"
//  INVALIDPARAMETER_WITHDETAILREASON = "InvalidParameter.WithDetailReason"
func (c *Client) ModifyCSRWithContext(ctx context.Context, request *ModifyCSRRequest) (response *ModifyCSRResponse, err error) {
    if request == nil {
        request = NewModifyCSRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCSR require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCSRResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCertificateAliasRequest() (request *ModifyCertificateAliasRequest) {
    request = &ModifyCertificateAliasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "ModifyCertificateAlias")
    
    
    return
}

func NewModifyCertificateAliasResponse() (response *ModifyCertificateAliasResponse) {
    response = &ModifyCertificateAliasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCertificateAlias
// This API is used to modify a certificate alias by passing in the certificate ID and new alias.
//
// error code that may be returned:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyCertificateAlias(request *ModifyCertificateAliasRequest) (response *ModifyCertificateAliasResponse, err error) {
    return c.ModifyCertificateAliasWithContext(context.Background(), request)
}

// ModifyCertificateAlias
// This API is used to modify a certificate alias by passing in the certificate ID and new alias.
//
// error code that may be returned:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyCertificateAliasWithContext(ctx context.Context, request *ModifyCertificateAliasRequest) (response *ModifyCertificateAliasResponse, err error) {
    if request == nil {
        request = NewModifyCertificateAliasRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCertificateAlias require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCertificateAliasResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCertificateProjectRequest() (request *ModifyCertificateProjectRequest) {
    request = &ModifyCertificateProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "ModifyCertificateProject")
    
    
    return
}

func NewModifyCertificateProjectResponse() (response *ModifyCertificateProjectResponse) {
    response = &ModifyCertificateProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCertificateProject
// This API is used to modify the projects of multiple certificates.
//
// error code that may be returned:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CAMAUTHORIZEDFAIL = "FailedOperation.CAMAuthorizedFail"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyCertificateProject(request *ModifyCertificateProjectRequest) (response *ModifyCertificateProjectResponse, err error) {
    return c.ModifyCertificateProjectWithContext(context.Background(), request)
}

// ModifyCertificateProject
// This API is used to modify the projects of multiple certificates.
//
// error code that may be returned:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CAMAUTHORIZEDFAIL = "FailedOperation.CAMAuthorizedFail"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyCertificateProjectWithContext(ctx context.Context, request *ModifyCertificateProjectRequest) (response *ModifyCertificateProjectResponse, err error) {
    if request == nil {
        request = NewModifyCertificateProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCertificateProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCertificateProjectResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCertificateResubmitRequest() (request *ModifyCertificateResubmitRequest) {
    request = &ModifyCertificateResubmitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "ModifyCertificateResubmit")
    
    
    return
}

func NewModifyCertificateResubmitResponse() (response *ModifyCertificateResubmitResponse) {
    response = &ModifyCertificateResubmitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCertificateResubmit
// This API is used to re-submit a review application for a paid certificate whose review failed or was canceled.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CERTIFICATESTATUSNOTALLOWRESUBMIT = "FailedOperation.CertificateStatusNotAllowResubmit"
//  INVALIDPARAMETER_CERTIFICATESTATUSNOTALLOWRESUBMIT = "InvalidParameter.CertificateStatusNotAllowResubmit"
//  INVALIDPARAMETER_WITHDETAILREASON = "InvalidParameter.WithDetailReason"
func (c *Client) ModifyCertificateResubmit(request *ModifyCertificateResubmitRequest) (response *ModifyCertificateResubmitResponse, err error) {
    return c.ModifyCertificateResubmitWithContext(context.Background(), request)
}

// ModifyCertificateResubmit
// This API is used to re-submit a review application for a paid certificate whose review failed or was canceled.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CERTIFICATESTATUSNOTALLOWRESUBMIT = "FailedOperation.CertificateStatusNotAllowResubmit"
//  INVALIDPARAMETER_CERTIFICATESTATUSNOTALLOWRESUBMIT = "InvalidParameter.CertificateStatusNotAllowResubmit"
//  INVALIDPARAMETER_WITHDETAILREASON = "InvalidParameter.WithDetailReason"
func (c *Client) ModifyCertificateResubmitWithContext(ctx context.Context, request *ModifyCertificateResubmitRequest) (response *ModifyCertificateResubmitResponse, err error) {
    if request == nil {
        request = NewModifyCertificateResubmitRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCertificateResubmit require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCertificateResubmitResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCertificatesExpiringNotificationSwitchRequest() (request *ModifyCertificatesExpiringNotificationSwitchRequest) {
    request = &ModifyCertificatesExpiringNotificationSwitchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "ModifyCertificatesExpiringNotificationSwitch")
    
    
    return
}

func NewModifyCertificatesExpiringNotificationSwitchResponse() (response *ModifyCertificatesExpiringNotificationSwitchResponse) {
    response = &ModifyCertificatesExpiringNotificationSwitchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCertificatesExpiringNotificationSwitch
// Modify to ignore certificate expiration notifications. Enable or disable certificate expiration notifications.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  INVALIDPARAMETER_CERTIFICATESNUMBEREXCEEDED = "InvalidParameter.CertificatesNumberExceeded"
func (c *Client) ModifyCertificatesExpiringNotificationSwitch(request *ModifyCertificatesExpiringNotificationSwitchRequest) (response *ModifyCertificatesExpiringNotificationSwitchResponse, err error) {
    return c.ModifyCertificatesExpiringNotificationSwitchWithContext(context.Background(), request)
}

// ModifyCertificatesExpiringNotificationSwitch
// Modify to ignore certificate expiration notifications. Enable or disable certificate expiration notifications.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  INVALIDPARAMETER_CERTIFICATESNUMBEREXCEEDED = "InvalidParameter.CertificatesNumberExceeded"
func (c *Client) ModifyCertificatesExpiringNotificationSwitchWithContext(ctx context.Context, request *ModifyCertificatesExpiringNotificationSwitchRequest) (response *ModifyCertificatesExpiringNotificationSwitchResponse, err error) {
    if request == nil {
        request = NewModifyCertificatesExpiringNotificationSwitchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCertificatesExpiringNotificationSwitch require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCertificatesExpiringNotificationSwitchResponse()
    err = c.Send(request, response)
    return
}

func NewReplaceCertificateRequest() (request *ReplaceCertificateRequest) {
    request = &ReplaceCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "ReplaceCertificate")
    
    
    return
}

func NewReplaceCertificateResponse() (response *ReplaceCertificateResponse) {
    response = &ReplaceCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReplaceCertificate
// This API is used to reissue a certificate. Note that if you have applied for a free certificate, only an RSA-2048 certificate will be reissued, and the certificate can be reissued only once.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESOURCE = "FailedOperation.InvalidCertificateSource"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_RENEWALGORITHMINVALID = "InvalidParameter.RenewAlgorithmInvalid"
func (c *Client) ReplaceCertificate(request *ReplaceCertificateRequest) (response *ReplaceCertificateResponse, err error) {
    return c.ReplaceCertificateWithContext(context.Background(), request)
}

// ReplaceCertificate
// This API is used to reissue a certificate. Note that if you have applied for a free certificate, only an RSA-2048 certificate will be reissued, and the certificate can be reissued only once.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESOURCE = "FailedOperation.InvalidCertificateSource"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_RENEWALGORITHMINVALID = "InvalidParameter.RenewAlgorithmInvalid"
func (c *Client) ReplaceCertificateWithContext(ctx context.Context, request *ReplaceCertificateRequest) (response *ReplaceCertificateResponse, err error) {
    if request == nil {
        request = NewReplaceCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReplaceCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewReplaceCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitCertificateInformationRequest() (request *SubmitCertificateInformationRequest) {
    request = &SubmitCertificateInformationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "SubmitCertificateInformation")
    
    
    return
}

func NewSubmitCertificateInformationResponse() (response *SubmitCertificateInformationResponse) {
    response = &SubmitCertificateInformationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitCertificateInformation
// Submit documentation for paid certificates; This API does not maintain new features, and you can use the new API to submit documentation. [CertificateInfoSubmit](https://intl.cloud.tencent.com/document/product/400/116033?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
func (c *Client) SubmitCertificateInformation(request *SubmitCertificateInformationRequest) (response *SubmitCertificateInformationResponse, err error) {
    return c.SubmitCertificateInformationWithContext(context.Background(), request)
}

// SubmitCertificateInformation
// Submit documentation for paid certificates; This API does not maintain new features, and you can use the new API to submit documentation. [CertificateInfoSubmit](https://intl.cloud.tencent.com/document/product/400/116033?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
func (c *Client) SubmitCertificateInformationWithContext(ctx context.Context, request *SubmitCertificateInformationRequest) (response *SubmitCertificateInformationResponse, err error) {
    if request == nil {
        request = NewSubmitCertificateInformationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitCertificateInformation require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitCertificateInformationResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCertificateInstanceRequest() (request *UpdateCertificateInstanceRequest) {
    request = &UpdateCertificateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "UpdateCertificateInstance")
    
    
    return
}

func NewUpdateCertificateInstanceResponse() (response *UpdateCertificateInstanceResponse) {
    response = &UpdateCertificateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateCertificateInstance
// This API is used to one-click update old certificate resources. This API is asynchronous. After calling it, if DeployRecordId is 0, it means the task is in progress. Repeat the request to this API. When the returned DeployRecordId is greater than 0, it means the task creation is successful. If it is not created successfully, an exception will be thrown.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEDEPLOYHASPENDINGRECORD = "FailedOperation.CertificateDeployHasPendingRecord"
//  FAILEDOPERATION_CERTIFICATEDEPLOYINSTANCEEMPTY = "FailedOperation.CertificateDeployInstanceEmpty"
//  FAILEDOPERATION_CERTIFICATEEXISTS = "FailedOperation.CertificateExists"
//  FAILEDOPERATION_CERTIFICATEHOSTDEPLOYCANNOTALLOW = "FailedOperation.CertificateHostDeployCanNotAllow"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINSTANCEHUGELIMIT = "FailedOperation.CertificateHostResourceInstanceHugeLimit"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCETYPEINVALID = "FailedOperation.CertificateHostResourceTypeInvalid"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMATCHERROR = "FailedOperation.CertificateMatchError"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTDEPLOYINSTANCE = "FailedOperation.CertificateNotDeployInstance"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CERTIFICATEPARSEERROR = "FailedOperation.CertificateParseError"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_WITHDETAILREASON = "InvalidParameter.WithDetailReason"
//  INVALIDPARAMETERVALUE_INTERMEDIATECERTNOTSAME = "InvalidParameterValue.IntermediateCertNotSame"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) UpdateCertificateInstance(request *UpdateCertificateInstanceRequest) (response *UpdateCertificateInstanceResponse, err error) {
    return c.UpdateCertificateInstanceWithContext(context.Background(), request)
}

// UpdateCertificateInstance
// This API is used to one-click update old certificate resources. This API is asynchronous. After calling it, if DeployRecordId is 0, it means the task is in progress. Repeat the request to this API. When the returned DeployRecordId is greater than 0, it means the task creation is successful. If it is not created successfully, an exception will be thrown.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEDEPLOYHASPENDINGRECORD = "FailedOperation.CertificateDeployHasPendingRecord"
//  FAILEDOPERATION_CERTIFICATEDEPLOYINSTANCEEMPTY = "FailedOperation.CertificateDeployInstanceEmpty"
//  FAILEDOPERATION_CERTIFICATEEXISTS = "FailedOperation.CertificateExists"
//  FAILEDOPERATION_CERTIFICATEHOSTDEPLOYCANNOTALLOW = "FailedOperation.CertificateHostDeployCanNotAllow"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINSTANCEHUGELIMIT = "FailedOperation.CertificateHostResourceInstanceHugeLimit"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCETYPEINVALID = "FailedOperation.CertificateHostResourceTypeInvalid"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMATCHERROR = "FailedOperation.CertificateMatchError"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTDEPLOYINSTANCE = "FailedOperation.CertificateNotDeployInstance"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CERTIFICATEPARSEERROR = "FailedOperation.CertificateParseError"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_WITHDETAILREASON = "InvalidParameter.WithDetailReason"
//  INVALIDPARAMETERVALUE_INTERMEDIATECERTNOTSAME = "InvalidParameterValue.IntermediateCertNotSame"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) UpdateCertificateInstanceWithContext(ctx context.Context, request *UpdateCertificateInstanceRequest) (response *UpdateCertificateInstanceResponse, err error) {
    if request == nil {
        request = NewUpdateCertificateInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCertificateInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCertificateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCertificateRecordRetryRequest() (request *UpdateCertificateRecordRetryRequest) {
    request = &UpdateCertificateRecordRetryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "UpdateCertificateRecordRetry")
    
    
    return
}

func NewUpdateCertificateRecordRetryResponse() (response *UpdateCertificateRecordRetryResponse) {
    response = &UpdateCertificateRecordRetryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateCertificateRecordRetry
// Cloud resource update deployment retry record
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEDEPLOYHASPENDINGRECORD = "FailedOperation.CertificateDeployHasPendingRecord"
//  FAILEDOPERATION_CERTIFICATEDEPLOYRETRYSTATUSINVALID = "FailedOperation.CertificateDeployRetryStatusInvalid"
//  FAILEDOPERATION_CERTIFICATENOTDEPLOYINSTANCE = "FailedOperation.CertificateNotDeployInstance"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) UpdateCertificateRecordRetry(request *UpdateCertificateRecordRetryRequest) (response *UpdateCertificateRecordRetryResponse, err error) {
    return c.UpdateCertificateRecordRetryWithContext(context.Background(), request)
}

// UpdateCertificateRecordRetry
// Cloud resource update deployment retry record
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEDEPLOYHASPENDINGRECORD = "FailedOperation.CertificateDeployHasPendingRecord"
//  FAILEDOPERATION_CERTIFICATEDEPLOYRETRYSTATUSINVALID = "FailedOperation.CertificateDeployRetryStatusInvalid"
//  FAILEDOPERATION_CERTIFICATENOTDEPLOYINSTANCE = "FailedOperation.CertificateNotDeployInstance"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) UpdateCertificateRecordRetryWithContext(ctx context.Context, request *UpdateCertificateRecordRetryRequest) (response *UpdateCertificateRecordRetryResponse, err error) {
    if request == nil {
        request = NewUpdateCertificateRecordRetryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCertificateRecordRetry require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCertificateRecordRetryResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCertificateRecordRollbackRequest() (request *UpdateCertificateRecordRollbackRequest) {
    request = &UpdateCertificateRecordRollbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "UpdateCertificateRecordRollback")
    
    
    return
}

func NewUpdateCertificateRecordRollbackResponse() (response *UpdateCertificateRecordRollbackResponse) {
    response = &UpdateCertificateRecordRollbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateCertificateRecordRollback
// Cloud resource update one-click rollback
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEDEPLOYDETAILROLLBACKSTATUSINVALID = "FailedOperation.CertificateDeployDetailRollbackStatusInvalid"
//  FAILEDOPERATION_CERTIFICATEDEPLOYHASPENDINGRECORD = "FailedOperation.CertificateDeployHasPendingRecord"
//  FAILEDOPERATION_CERTIFICATEDEPLOYRETRYSTATUSINVALID = "FailedOperation.CertificateDeployRetryStatusInvalid"
//  FAILEDOPERATION_CERTIFICATEDEPLOYROLLBACKSTATUSINVALID = "FailedOperation.CertificateDeployRollbackStatusInvalid"
//  FAILEDOPERATION_CERTIFICATENOTDEPLOYINSTANCE = "FailedOperation.CertificateNotDeployInstance"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) UpdateCertificateRecordRollback(request *UpdateCertificateRecordRollbackRequest) (response *UpdateCertificateRecordRollbackResponse, err error) {
    return c.UpdateCertificateRecordRollbackWithContext(context.Background(), request)
}

// UpdateCertificateRecordRollback
// Cloud resource update one-click rollback
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEDEPLOYDETAILROLLBACKSTATUSINVALID = "FailedOperation.CertificateDeployDetailRollbackStatusInvalid"
//  FAILEDOPERATION_CERTIFICATEDEPLOYHASPENDINGRECORD = "FailedOperation.CertificateDeployHasPendingRecord"
//  FAILEDOPERATION_CERTIFICATEDEPLOYRETRYSTATUSINVALID = "FailedOperation.CertificateDeployRetryStatusInvalid"
//  FAILEDOPERATION_CERTIFICATEDEPLOYROLLBACKSTATUSINVALID = "FailedOperation.CertificateDeployRollbackStatusInvalid"
//  FAILEDOPERATION_CERTIFICATENOTDEPLOYINSTANCE = "FailedOperation.CertificateNotDeployInstance"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) UpdateCertificateRecordRollbackWithContext(ctx context.Context, request *UpdateCertificateRecordRollbackRequest) (response *UpdateCertificateRecordRollbackResponse, err error) {
    if request == nil {
        request = NewUpdateCertificateRecordRollbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCertificateRecordRollback require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCertificateRecordRollbackResponse()
    err = c.Send(request, response)
    return
}

func NewUploadCertificateRequest() (request *UploadCertificateRequest) {
    request = &UploadCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "UploadCertificate")
    
    
    return
}

func NewUploadCertificateResponse() (response *UploadCertificateResponse) {
    response = &UploadCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UploadCertificate
// This API is used to upload a certificate.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CAMAUTHORIZEDFAIL = "FailedOperation.CAMAuthorizedFail"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATECAERROR = "FailedOperation.CertificateCaError"
//  FAILEDOPERATION_CERTIFICATEENCRYPTINVALID = "FailedOperation.CertificateEncryptInvalid"
//  FAILEDOPERATION_CERTIFICATEEXISTS = "FailedOperation.CertificateExists"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMATCHERROR = "FailedOperation.CertificateMatchError"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CERTIFICATEPARSEERROR = "FailedOperation.CertificateParseError"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_MISSMATCHPRIVATEKEY = "InvalidParameter.MissMatchPrivateKey"
//  INVALIDPARAMETER_WITHDETAILREASON = "InvalidParameter.WithDetailReason"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) UploadCertificate(request *UploadCertificateRequest) (response *UploadCertificateResponse, err error) {
    return c.UploadCertificateWithContext(context.Background(), request)
}

// UploadCertificate
// This API is used to upload a certificate.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CAMAUTHORIZEDFAIL = "FailedOperation.CAMAuthorizedFail"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATECAERROR = "FailedOperation.CertificateCaError"
//  FAILEDOPERATION_CERTIFICATEENCRYPTINVALID = "FailedOperation.CertificateEncryptInvalid"
//  FAILEDOPERATION_CERTIFICATEEXISTS = "FailedOperation.CertificateExists"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMATCHERROR = "FailedOperation.CertificateMatchError"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CERTIFICATEPARSEERROR = "FailedOperation.CertificateParseError"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_MISSMATCHPRIVATEKEY = "InvalidParameter.MissMatchPrivateKey"
//  INVALIDPARAMETER_WITHDETAILREASON = "InvalidParameter.WithDetailReason"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) UploadCertificateWithContext(ctx context.Context, request *UploadCertificateRequest) (response *UploadCertificateResponse, err error) {
    if request == nil {
        request = NewUploadCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewUploadConfirmLetterRequest() (request *UploadConfirmLetterRequest) {
    request = &UploadConfirmLetterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "UploadConfirmLetter")
    
    
    return
}

func NewUploadConfirmLetterResponse() (response *UploadConfirmLetterResponse) {
    response = &UploadConfirmLetterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UploadConfirmLetter
// This API is used to upload the confirmation letter for a certificate.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CONFIRMLETTERTOOLARGE = "FailedOperation.ConfirmLetterTooLarge"
//  FAILEDOPERATION_CONFIRMLETTERTOOSMALL = "FailedOperation.ConfirmLetterTooSmall"
//  FAILEDOPERATION_INVALIDCERTIFICATESOURCE = "FailedOperation.InvalidCertificateSource"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDCONFIRMLETTERFORMAT = "FailedOperation.InvalidConfirmLetterFormat"
//  FAILEDOPERATION_INVALIDCONFIRMLETTERFORMATWOSIGN = "FailedOperation.InvalidConfirmLetterFormatWosign"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  INTERNALERROR = "InternalError"
func (c *Client) UploadConfirmLetter(request *UploadConfirmLetterRequest) (response *UploadConfirmLetterResponse, err error) {
    return c.UploadConfirmLetterWithContext(context.Background(), request)
}

// UploadConfirmLetter
// This API is used to upload the confirmation letter for a certificate.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CONFIRMLETTERTOOLARGE = "FailedOperation.ConfirmLetterTooLarge"
//  FAILEDOPERATION_CONFIRMLETTERTOOSMALL = "FailedOperation.ConfirmLetterTooSmall"
//  FAILEDOPERATION_INVALIDCERTIFICATESOURCE = "FailedOperation.InvalidCertificateSource"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDCONFIRMLETTERFORMAT = "FailedOperation.InvalidConfirmLetterFormat"
//  FAILEDOPERATION_INVALIDCONFIRMLETTERFORMATWOSIGN = "FailedOperation.InvalidConfirmLetterFormatWosign"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  INTERNALERROR = "InternalError"
func (c *Client) UploadConfirmLetterWithContext(ctx context.Context, request *UploadConfirmLetterRequest) (response *UploadConfirmLetterResponse, err error) {
    if request == nil {
        request = NewUploadConfirmLetterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadConfirmLetter require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadConfirmLetterResponse()
    err = c.Send(request, response)
    return
}

func NewUploadUpdateCertificateInstanceRequest() (request *UploadUpdateCertificateInstanceRequest) {
    request = &UploadUpdateCertificateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "UploadUpdateCertificateInstance")
    
    
    return
}

func NewUploadUpdateCertificateInstanceResponse() (response *UploadUpdateCertificateInstanceResponse) {
    response = &UploadUpdateCertificateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UploadUpdateCertificateInstance
// This API is used to update certificate content (certificate ID unchanged) and update associated Tencent Cloud resources. This is an asynchronous API. After calling, a DeployRecordId of 0 indicates that the task is in progress. Repeatedly request this API, and when DeployRecordId is greater than 0, it means the task has been successfully created. If the task is not successfully created, an exception will be thrown.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEDEPLOYHASPENDINGRECORD = "FailedOperation.CertificateDeployHasPendingRecord"
//  FAILEDOPERATION_CERTIFICATEDEPLOYINSTANCEEMPTY = "FailedOperation.CertificateDeployInstanceEmpty"
//  FAILEDOPERATION_CERTIFICATEEXISTS = "FailedOperation.CertificateExists"
//  FAILEDOPERATION_CERTIFICATEHOSTDEPLOYCANNOTALLOW = "FailedOperation.CertificateHostDeployCanNotAllow"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINSTANCEHUGELIMIT = "FailedOperation.CertificateHostResourceInstanceHugeLimit"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCETYPEINVALID = "FailedOperation.CertificateHostResourceTypeInvalid"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMATCHERROR = "FailedOperation.CertificateMatchError"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTDEPLOYINSTANCE = "FailedOperation.CertificateNotDeployInstance"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CERTIFICATEPARSEERROR = "FailedOperation.CertificateParseError"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_WITHDETAILREASON = "InvalidParameter.WithDetailReason"
//  INVALIDPARAMETERVALUE_INTERMEDIATECERTNOTSAME = "InvalidParameterValue.IntermediateCertNotSame"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) UploadUpdateCertificateInstance(request *UploadUpdateCertificateInstanceRequest) (response *UploadUpdateCertificateInstanceResponse, err error) {
    return c.UploadUpdateCertificateInstanceWithContext(context.Background(), request)
}

// UploadUpdateCertificateInstance
// This API is used to update certificate content (certificate ID unchanged) and update associated Tencent Cloud resources. This is an asynchronous API. After calling, a DeployRecordId of 0 indicates that the task is in progress. Repeatedly request this API, and when DeployRecordId is greater than 0, it means the task has been successfully created. If the task is not successfully created, an exception will be thrown.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEDEPLOYHASPENDINGRECORD = "FailedOperation.CertificateDeployHasPendingRecord"
//  FAILEDOPERATION_CERTIFICATEDEPLOYINSTANCEEMPTY = "FailedOperation.CertificateDeployInstanceEmpty"
//  FAILEDOPERATION_CERTIFICATEEXISTS = "FailedOperation.CertificateExists"
//  FAILEDOPERATION_CERTIFICATEHOSTDEPLOYCANNOTALLOW = "FailedOperation.CertificateHostDeployCanNotAllow"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINSTANCEHUGELIMIT = "FailedOperation.CertificateHostResourceInstanceHugeLimit"
//  FAILEDOPERATION_CERTIFICATEHOSTRESOURCETYPEINVALID = "FailedOperation.CertificateHostResourceTypeInvalid"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMATCHERROR = "FailedOperation.CertificateMatchError"
//  FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"
//  FAILEDOPERATION_CERTIFICATENOTDEPLOYINSTANCE = "FailedOperation.CertificateNotDeployInstance"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CERTIFICATEPARSEERROR = "FailedOperation.CertificateParseError"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_WITHDETAILREASON = "InvalidParameter.WithDetailReason"
//  INVALIDPARAMETERVALUE_INTERMEDIATECERTNOTSAME = "InvalidParameterValue.IntermediateCertNotSame"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) UploadUpdateCertificateInstanceWithContext(ctx context.Context, request *UploadUpdateCertificateInstanceRequest) (response *UploadUpdateCertificateInstanceResponse, err error) {
    if request == nil {
        request = NewUploadUpdateCertificateInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadUpdateCertificateInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadUpdateCertificateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUploadUpdateCertificateRecordRetryRequest() (request *UploadUpdateCertificateRecordRetryRequest) {
    request = &UploadUpdateCertificateRecordRetryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "UploadUpdateCertificateRecordRetry")
    
    
    return
}

func NewUploadUpdateCertificateRecordRetryResponse() (response *UploadUpdateCertificateRecordRetryResponse) {
    response = &UploadUpdateCertificateRecordRetryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UploadUpdateCertificateRecordRetry
// Cloud Resource Update (Certificate ID Unchanged) Deployment Retry Record.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEDEPLOYHASPENDINGRECORD = "FailedOperation.CertificateDeployHasPendingRecord"
//  FAILEDOPERATION_CERTIFICATEDEPLOYRETRYSTATUSINVALID = "FailedOperation.CertificateDeployRetryStatusInvalid"
//  FAILEDOPERATION_CERTIFICATENOTDEPLOYINSTANCE = "FailedOperation.CertificateNotDeployInstance"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) UploadUpdateCertificateRecordRetry(request *UploadUpdateCertificateRecordRetryRequest) (response *UploadUpdateCertificateRecordRetryResponse, err error) {
    return c.UploadUpdateCertificateRecordRetryWithContext(context.Background(), request)
}

// UploadUpdateCertificateRecordRetry
// Cloud Resource Update (Certificate ID Unchanged) Deployment Retry Record.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEDEPLOYHASPENDINGRECORD = "FailedOperation.CertificateDeployHasPendingRecord"
//  FAILEDOPERATION_CERTIFICATEDEPLOYRETRYSTATUSINVALID = "FailedOperation.CertificateDeployRetryStatusInvalid"
//  FAILEDOPERATION_CERTIFICATENOTDEPLOYINSTANCE = "FailedOperation.CertificateNotDeployInstance"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) UploadUpdateCertificateRecordRetryWithContext(ctx context.Context, request *UploadUpdateCertificateRecordRetryRequest) (response *UploadUpdateCertificateRecordRetryResponse, err error) {
    if request == nil {
        request = NewUploadUpdateCertificateRecordRetryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadUpdateCertificateRecordRetry require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadUpdateCertificateRecordRetryResponse()
    err = c.Send(request, response)
    return
}

func NewUploadUpdateCertificateRecordRollbackRequest() (request *UploadUpdateCertificateRecordRollbackRequest) {
    request = &UploadUpdateCertificateRecordRollbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "UploadUpdateCertificateRecordRollback")
    
    
    return
}

func NewUploadUpdateCertificateRecordRollbackResponse() (response *UploadUpdateCertificateRecordRollbackResponse) {
    response = &UploadUpdateCertificateRecordRollbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UploadUpdateCertificateRecordRollback
// This API is used to roll back the full task when cloud resource update succeeds with unchanged certificate ID.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEDEPLOYDETAILROLLBACKSTATUSINVALID = "FailedOperation.CertificateDeployDetailRollbackStatusInvalid"
//  FAILEDOPERATION_CERTIFICATEDEPLOYHASPENDINGRECORD = "FailedOperation.CertificateDeployHasPendingRecord"
//  FAILEDOPERATION_CERTIFICATEDEPLOYRETRYSTATUSINVALID = "FailedOperation.CertificateDeployRetryStatusInvalid"
//  FAILEDOPERATION_CERTIFICATEDEPLOYROLLBACKSTATUSINVALID = "FailedOperation.CertificateDeployRollbackStatusInvalid"
//  FAILEDOPERATION_CERTIFICATENOTDEPLOYINSTANCE = "FailedOperation.CertificateNotDeployInstance"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) UploadUpdateCertificateRecordRollback(request *UploadUpdateCertificateRecordRollbackRequest) (response *UploadUpdateCertificateRecordRollbackResponse, err error) {
    return c.UploadUpdateCertificateRecordRollbackWithContext(context.Background(), request)
}

// UploadUpdateCertificateRecordRollback
// This API is used to roll back the full task when cloud resource update succeeds with unchanged certificate ID.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEDEPLOYDETAILROLLBACKSTATUSINVALID = "FailedOperation.CertificateDeployDetailRollbackStatusInvalid"
//  FAILEDOPERATION_CERTIFICATEDEPLOYHASPENDINGRECORD = "FailedOperation.CertificateDeployHasPendingRecord"
//  FAILEDOPERATION_CERTIFICATEDEPLOYRETRYSTATUSINVALID = "FailedOperation.CertificateDeployRetryStatusInvalid"
//  FAILEDOPERATION_CERTIFICATEDEPLOYROLLBACKSTATUSINVALID = "FailedOperation.CertificateDeployRollbackStatusInvalid"
//  FAILEDOPERATION_CERTIFICATENOTDEPLOYINSTANCE = "FailedOperation.CertificateNotDeployInstance"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) UploadUpdateCertificateRecordRollbackWithContext(ctx context.Context, request *UploadUpdateCertificateRecordRollbackRequest) (response *UploadUpdateCertificateRecordRollbackResponse, err error) {
    if request == nil {
        request = NewUploadUpdateCertificateRecordRollbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadUpdateCertificateRecordRollback require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadUpdateCertificateRecordRollbackResponse()
    err = c.Send(request, response)
    return
}
