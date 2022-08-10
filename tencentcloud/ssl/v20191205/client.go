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

package v20191205

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
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
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  INTERNALERROR = "InternalError"
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
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  INTERNALERROR = "InternalError"
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
// This API is used to submit a certificate order.
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
func (c *Client) CommitCertificateInformation(request *CommitCertificateInformationRequest) (response *CommitCertificateInformationResponse, err error) {
    return c.CommitCertificateInformationWithContext(context.Background(), request)
}

// CommitCertificateInformation
// This API is used to submit a certificate order.
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
func (c *Client) DeleteCertificate(request *DeleteCertificateRequest) (response *DeleteCertificateResponse, err error) {
    return c.DeleteCertificateWithContext(context.Background(), request)
}

// DeleteCertificate
// This API is used to delete a certificate.
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
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeCertificateDetail(request *DescribeCertificateDetailRequest) (response *DescribeCertificateDetailResponse, err error) {
    return c.DescribeCertificateDetailWithContext(context.Background(), request)
}

// DescribeCertificateDetail
// This API is used to get certificate details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
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
func (c *Client) DescribeCertificates(request *DescribeCertificatesRequest) (response *DescribeCertificatesResponse, err error) {
    return c.DescribeCertificatesWithContext(context.Background(), request)
}

// DescribeCertificates
// This API is used to get the certificate list.
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
func (c *Client) ReplaceCertificate(request *ReplaceCertificateRequest) (response *ReplaceCertificateResponse, err error) {
    return c.ReplaceCertificateWithContext(context.Background(), request)
}

// ReplaceCertificate
// This API is used to reissue a certificate. Note that if you have applied for a free certificate, only an RSA-2048 certificate will be reissued, and the certificate can be reissued only once.
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
// This API is used to submit certificate information.
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
// This API is used to submit certificate information.
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
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEEXISTS = "FailedOperation.CertificateExists"
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
func (c *Client) UploadCertificate(request *UploadCertificateRequest) (response *UploadCertificateResponse, err error) {
    return c.UploadCertificateWithContext(context.Background(), request)
}

// UploadCertificate
// This API is used to upload a certificate.
//
// error code that may be returned:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEEXISTS = "FailedOperation.CertificateExists"
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
