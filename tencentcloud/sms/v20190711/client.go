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

package v20190711

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2019-07-11"

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


func NewAddSmsSignRequest() (request *AddSmsSignRequest) {
    request = &AddSmsSignRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sms", APIVersion, "AddSmsSign")
    
    
    return
}

func NewAddSmsSignResponse() (response *AddSmsSignResponse) {
    response = &AddSmsSignResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddSmsSign
// This API is used to add an SMS signature. Please read the [Tencent Cloud SMS Signature Review Standards](https://intl.cloud.tencent.com/document/product/382/39022?from_cn_redirect=1) before starting an application.
//
// > Note: individual users cannot use this API to apply for SMS signatures. For more information, please see [Identity Verification Overview](https://intl.cloud.tencent.com/document/product/378/3629?from_cn_redirect=1). If your account identity is individual, please log in to the console to apply for SMS signatures. For detailed directions, please see [Creating SMS Signatures](https://intl.cloud.tencent.com/document/product/382/36136?from_cn_redirect=1#Sign).
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_APPIDANDBIZID = "InvalidParameter.AppidAndBizId"
//  INVALIDPARAMETERVALUE_IMAGEINVALID = "InvalidParameterValue.ImageInvalid"
//  INVALIDPARAMETERVALUE_INVALIDDOCUMENTTYPE = "InvalidParameterValue.InvalidDocumentType"
//  INVALIDPARAMETERVALUE_INVALIDINTERNATIONAL = "InvalidParameterValue.InvalidInternational"
//  INVALIDPARAMETERVALUE_INVALIDUSEDMETHOD = "InvalidParameterValue.InvalidUsedMethod"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppidNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) AddSmsSign(request *AddSmsSignRequest) (response *AddSmsSignResponse, err error) {
    return c.AddSmsSignWithContext(context.Background(), request)
}

// AddSmsSign
// This API is used to add an SMS signature. Please read the [Tencent Cloud SMS Signature Review Standards](https://intl.cloud.tencent.com/document/product/382/39022?from_cn_redirect=1) before starting an application.
//
// > Note: individual users cannot use this API to apply for SMS signatures. For more information, please see [Identity Verification Overview](https://intl.cloud.tencent.com/document/product/378/3629?from_cn_redirect=1). If your account identity is individual, please log in to the console to apply for SMS signatures. For detailed directions, please see [Creating SMS Signatures](https://intl.cloud.tencent.com/document/product/382/36136?from_cn_redirect=1#Sign).
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_APPIDANDBIZID = "InvalidParameter.AppidAndBizId"
//  INVALIDPARAMETERVALUE_IMAGEINVALID = "InvalidParameterValue.ImageInvalid"
//  INVALIDPARAMETERVALUE_INVALIDDOCUMENTTYPE = "InvalidParameterValue.InvalidDocumentType"
//  INVALIDPARAMETERVALUE_INVALIDINTERNATIONAL = "InvalidParameterValue.InvalidInternational"
//  INVALIDPARAMETERVALUE_INVALIDUSEDMETHOD = "InvalidParameterValue.InvalidUsedMethod"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppidNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) AddSmsSignWithContext(ctx context.Context, request *AddSmsSignRequest) (response *AddSmsSignResponse, err error) {
    if request == nil {
        request = NewAddSmsSignRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddSmsSign require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddSmsSignResponse()
    err = c.Send(request, response)
    return
}

func NewAddSmsTemplateRequest() (request *AddSmsTemplateRequest) {
    request = &AddSmsTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sms", APIVersion, "AddSmsTemplate")
    
    
    return
}

func NewAddSmsTemplateResponse() (response *AddSmsTemplateResponse) {
    response = &AddSmsTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddSmsTemplate
// This API is used to add an SMS template. Please read the [Tencent Cloud SMS Body Template Review Standards](https://intl.cloud.tencent.com/document/product/382/39023?from_cn_redirect=1) before starting an application.
//
// > Note: individual users cannot use this API to apply for SMS body templates. For more information, please see [Identity Verification Overview](https://intl.cloud.tencent.com/document/product/378/3629?from_cn_redirect=1). If your account identity is individual, please log in to the console to apply for SMS body templates. For detailed directions, please see [Creating SMS Body Templates](https://intl.cloud.tencent.com/document/product/382/36136?from_cn_redirect=1#Template).
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_MISSINGSIGNATURE = "FailedOperation.MissingSignature"
//  FAILEDOPERATION_MISSINGTEMPLATETOMODIFY = "FailedOperation.MissingTemplateToModify"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  FAILEDOPERATION_TEMPLATEALREADYPASSEDCHECK = "FailedOperation.TemplateAlreadyPassedCheck"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMETERS = "InvalidParameter.InvalidParameters"
//  INVALIDPARAMETERVALUE_INVALIDINTERNATIONAL = "InvalidParameterValue.InvalidInternational"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppidIsDisabled"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) AddSmsTemplate(request *AddSmsTemplateRequest) (response *AddSmsTemplateResponse, err error) {
    return c.AddSmsTemplateWithContext(context.Background(), request)
}

// AddSmsTemplate
// This API is used to add an SMS template. Please read the [Tencent Cloud SMS Body Template Review Standards](https://intl.cloud.tencent.com/document/product/382/39023?from_cn_redirect=1) before starting an application.
//
// > Note: individual users cannot use this API to apply for SMS body templates. For more information, please see [Identity Verification Overview](https://intl.cloud.tencent.com/document/product/378/3629?from_cn_redirect=1). If your account identity is individual, please log in to the console to apply for SMS body templates. For detailed directions, please see [Creating SMS Body Templates](https://intl.cloud.tencent.com/document/product/382/36136?from_cn_redirect=1#Template).
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_MISSINGSIGNATURE = "FailedOperation.MissingSignature"
//  FAILEDOPERATION_MISSINGTEMPLATETOMODIFY = "FailedOperation.MissingTemplateToModify"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  FAILEDOPERATION_TEMPLATEALREADYPASSEDCHECK = "FailedOperation.TemplateAlreadyPassedCheck"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMETERS = "InvalidParameter.InvalidParameters"
//  INVALIDPARAMETERVALUE_INVALIDINTERNATIONAL = "InvalidParameterValue.InvalidInternational"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppidIsDisabled"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) AddSmsTemplateWithContext(ctx context.Context, request *AddSmsTemplateRequest) (response *AddSmsTemplateResponse, err error) {
    if request == nil {
        request = NewAddSmsTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddSmsTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddSmsTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCallbackStatusStatisticsRequest() (request *CallbackStatusStatisticsRequest) {
    request = &CallbackStatusStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sms", APIVersion, "CallbackStatusStatistics")
    
    
    return
}

func NewCallbackStatusStatisticsResponse() (response *CallbackStatusStatisticsResponse) {
    response = &CallbackStatusStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CallbackStatusStatistics
// This API is used to collect statistics on user receipts.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PHONENUMBERONBLACKLIST = "FailedOperation.PhoneNumberOnBlacklist"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppidNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppidIsDisabled"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppidVerifyFail"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) CallbackStatusStatistics(request *CallbackStatusStatisticsRequest) (response *CallbackStatusStatisticsResponse, err error) {
    return c.CallbackStatusStatisticsWithContext(context.Background(), request)
}

// CallbackStatusStatistics
// This API is used to collect statistics on user receipts.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PHONENUMBERONBLACKLIST = "FailedOperation.PhoneNumberOnBlacklist"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppidNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppidIsDisabled"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppidVerifyFail"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) CallbackStatusStatisticsWithContext(ctx context.Context, request *CallbackStatusStatisticsRequest) (response *CallbackStatusStatisticsResponse, err error) {
    if request == nil {
        request = NewCallbackStatusStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CallbackStatusStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewCallbackStatusStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSmsSignRequest() (request *DeleteSmsSignRequest) {
    request = &DeleteSmsSignRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sms", APIVersion, "DeleteSmsSign")
    
    
    return
}

func NewDeleteSmsSignResponse() (response *DeleteSmsSignResponse) {
    response = &DeleteSmsSignResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSmsSign
// > Note: individual users cannot use this API to delete SMS signatures. For more information, please see [Identity Verification Overview](https://intl.cloud.tencent.com/document/product/378/3629?from_cn_redirect=1). You can log in to the console to delete SMS signatures. For detailed directions, please see the notes on deleting SMS signatures in [SMS Signature Operations](https://intl.cloud.tencent.com/document/product/382/36136?from_cn_redirect=1#Sign).
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_MISSINGSIGNATURETOMODIFY = "FailedOperation.MissingSignatureToModify"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) DeleteSmsSign(request *DeleteSmsSignRequest) (response *DeleteSmsSignResponse, err error) {
    return c.DeleteSmsSignWithContext(context.Background(), request)
}

// DeleteSmsSign
// > Note: individual users cannot use this API to delete SMS signatures. For more information, please see [Identity Verification Overview](https://intl.cloud.tencent.com/document/product/378/3629?from_cn_redirect=1). You can log in to the console to delete SMS signatures. For detailed directions, please see the notes on deleting SMS signatures in [SMS Signature Operations](https://intl.cloud.tencent.com/document/product/382/36136?from_cn_redirect=1#Sign).
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_MISSINGSIGNATURETOMODIFY = "FailedOperation.MissingSignatureToModify"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) DeleteSmsSignWithContext(ctx context.Context, request *DeleteSmsSignRequest) (response *DeleteSmsSignResponse, err error) {
    if request == nil {
        request = NewDeleteSmsSignRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSmsSign require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSmsSignResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSmsTemplateRequest() (request *DeleteSmsTemplateRequest) {
    request = &DeleteSmsTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sms", APIVersion, "DeleteSmsTemplate")
    
    
    return
}

func NewDeleteSmsTemplateResponse() (response *DeleteSmsTemplateResponse) {
    response = &DeleteSmsTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSmsTemplate
// > Note: individual users cannot use this API to delete SMS body templates. For more information, please see [Identity Verification Overview](https://intl.cloud.tencent.com/document/product/378/3629?from_cn_redirect=1). You can log in to the console to delete SMS body templates. For detailed directions, please see the notes on deleting SMS body templates in [SMS Body Template Operations](https://intl.cloud.tencent.com/document/product/382/36136?from_cn_redirect=1#Template).
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_MISSINGTEMPLATETOMODIFY = "FailedOperation.MissingTemplateToModify"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppidIsDisabled"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) DeleteSmsTemplate(request *DeleteSmsTemplateRequest) (response *DeleteSmsTemplateResponse, err error) {
    return c.DeleteSmsTemplateWithContext(context.Background(), request)
}

// DeleteSmsTemplate
// > Note: individual users cannot use this API to delete SMS body templates. For more information, please see [Identity Verification Overview](https://intl.cloud.tencent.com/document/product/378/3629?from_cn_redirect=1). You can log in to the console to delete SMS body templates. For detailed directions, please see the notes on deleting SMS body templates in [SMS Body Template Operations](https://intl.cloud.tencent.com/document/product/382/36136?from_cn_redirect=1#Template).
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_MISSINGTEMPLATETOMODIFY = "FailedOperation.MissingTemplateToModify"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppidIsDisabled"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) DeleteSmsTemplateWithContext(ctx context.Context, request *DeleteSmsTemplateRequest) (response *DeleteSmsTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteSmsTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSmsTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSmsTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSmsSignListRequest() (request *DescribeSmsSignListRequest) {
    request = &DescribeSmsSignListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sms", APIVersion, "DescribeSmsSignList")
    
    
    return
}

func NewDescribeSmsSignListResponse() (response *DescribeSmsSignListResponse) {
    response = &DescribeSmsSignListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSmsSignList
// > Note: individual users cannot use this API to query SMS signatures. For more information, please see [Identity Verification Overview](https://intl.cloud.tencent.com/document/product/378/3629?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_MISSINGSIGNATURETOMODIFY = "FailedOperation.MissingSignatureToModify"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppidIsDisabled"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) DescribeSmsSignList(request *DescribeSmsSignListRequest) (response *DescribeSmsSignListResponse, err error) {
    return c.DescribeSmsSignListWithContext(context.Background(), request)
}

// DescribeSmsSignList
// > Note: individual users cannot use this API to query SMS signatures. For more information, please see [Identity Verification Overview](https://intl.cloud.tencent.com/document/product/378/3629?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_MISSINGSIGNATURETOMODIFY = "FailedOperation.MissingSignatureToModify"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppidIsDisabled"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) DescribeSmsSignListWithContext(ctx context.Context, request *DescribeSmsSignListRequest) (response *DescribeSmsSignListResponse, err error) {
    if request == nil {
        request = NewDescribeSmsSignListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSmsSignList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSmsSignListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSmsTemplateListRequest() (request *DescribeSmsTemplateListRequest) {
    request = &DescribeSmsTemplateListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sms", APIVersion, "DescribeSmsTemplateList")
    
    
    return
}

func NewDescribeSmsTemplateListResponse() (response *DescribeSmsTemplateListResponse) {
    response = &DescribeSmsTemplateListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSmsTemplateList
// > Note: individual users cannot use this API to query SMS body templates. For more information, please see [Identity Verification Overview](https://intl.cloud.tencent.com/document/product/378/3629?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_MISSINGTEMPLATETOMODIFY = "FailedOperation.MissingTemplateToModify"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppidIsDisabled"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) DescribeSmsTemplateList(request *DescribeSmsTemplateListRequest) (response *DescribeSmsTemplateListResponse, err error) {
    return c.DescribeSmsTemplateListWithContext(context.Background(), request)
}

// DescribeSmsTemplateList
// > Note: individual users cannot use this API to query SMS body templates. For more information, please see [Identity Verification Overview](https://intl.cloud.tencent.com/document/product/378/3629?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_MISSINGTEMPLATETOMODIFY = "FailedOperation.MissingTemplateToModify"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppidIsDisabled"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) DescribeSmsTemplateListWithContext(ctx context.Context, request *DescribeSmsTemplateListRequest) (response *DescribeSmsTemplateListResponse, err error) {
    if request == nil {
        request = NewDescribeSmsTemplateListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSmsTemplateList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSmsTemplateListResponse()
    err = c.Send(request, response)
    return
}

func NewModifySmsSignRequest() (request *ModifySmsSignRequest) {
    request = &ModifySmsSignRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sms", APIVersion, "ModifySmsSign")
    
    
    return
}

func NewModifySmsSignResponse() (response *ModifySmsSignResponse) {
    response = &ModifySmsSignResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySmsSign
// This API is used to modify an SMS signature. Please read the [Tencent Cloud SMS Signature Review Standards](https://intl.cloud.tencent.com/document/product/382/39022?from_cn_redirect=1) before making a modification.
//
// >-  Note: individual users cannot use this API to modify SMS signatures. For more information, please see [Identity Verification Overview](https://intl.cloud.tencent.com/document/product/378/3629?from_cn_redirect=1). If your account identity is individual, you can log in to the console to modify SMS signatures.
//
// >- Modifications can be made only if the signature status is **pending review** or **rejected**. **Approved** signatures cannot be modified.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_MISSINGSIGNATURETOMODIFY = "FailedOperation.MissingSignatureToModify"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  FAILEDOPERATION_TEMPLATEALREADYPASSEDCHECK = "FailedOperation.TemplateAlreadyPassedCheck"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPIDANDBIZID = "InvalidParameter.AppidAndBizId"
//  INVALIDPARAMETERVALUE_IMAGEINVALID = "InvalidParameterValue.ImageInvalid"
//  INVALIDPARAMETERVALUE_INVALIDDOCUMENTTYPE = "InvalidParameterValue.InvalidDocumentType"
//  INVALIDPARAMETERVALUE_INVALIDINTERNATIONAL = "InvalidParameterValue.InvalidInternational"
//  INVALIDPARAMETERVALUE_INVALIDUSEDMETHOD = "InvalidParameterValue.InvalidUsedMethod"
//  INVALIDPARAMETERVALUE_MISSINGSIGNATURELIST = "InvalidParameterValue.MissingSignatureList"
//  INVALIDPARAMETERVALUE_SIGNALREADYPASSEDCHECK = "InvalidParameterValue.SignAlreadyPassedCheck"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) ModifySmsSign(request *ModifySmsSignRequest) (response *ModifySmsSignResponse, err error) {
    return c.ModifySmsSignWithContext(context.Background(), request)
}

// ModifySmsSign
// This API is used to modify an SMS signature. Please read the [Tencent Cloud SMS Signature Review Standards](https://intl.cloud.tencent.com/document/product/382/39022?from_cn_redirect=1) before making a modification.
//
// >-  Note: individual users cannot use this API to modify SMS signatures. For more information, please see [Identity Verification Overview](https://intl.cloud.tencent.com/document/product/378/3629?from_cn_redirect=1). If your account identity is individual, you can log in to the console to modify SMS signatures.
//
// >- Modifications can be made only if the signature status is **pending review** or **rejected**. **Approved** signatures cannot be modified.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_MISSINGSIGNATURETOMODIFY = "FailedOperation.MissingSignatureToModify"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  FAILEDOPERATION_TEMPLATEALREADYPASSEDCHECK = "FailedOperation.TemplateAlreadyPassedCheck"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPIDANDBIZID = "InvalidParameter.AppidAndBizId"
//  INVALIDPARAMETERVALUE_IMAGEINVALID = "InvalidParameterValue.ImageInvalid"
//  INVALIDPARAMETERVALUE_INVALIDDOCUMENTTYPE = "InvalidParameterValue.InvalidDocumentType"
//  INVALIDPARAMETERVALUE_INVALIDINTERNATIONAL = "InvalidParameterValue.InvalidInternational"
//  INVALIDPARAMETERVALUE_INVALIDUSEDMETHOD = "InvalidParameterValue.InvalidUsedMethod"
//  INVALIDPARAMETERVALUE_MISSINGSIGNATURELIST = "InvalidParameterValue.MissingSignatureList"
//  INVALIDPARAMETERVALUE_SIGNALREADYPASSEDCHECK = "InvalidParameterValue.SignAlreadyPassedCheck"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) ModifySmsSignWithContext(ctx context.Context, request *ModifySmsSignRequest) (response *ModifySmsSignResponse, err error) {
    if request == nil {
        request = NewModifySmsSignRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySmsSign require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySmsSignResponse()
    err = c.Send(request, response)
    return
}

func NewModifySmsTemplateRequest() (request *ModifySmsTemplateRequest) {
    request = &ModifySmsTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sms", APIVersion, "ModifySmsTemplate")
    
    
    return
}

func NewModifySmsTemplateResponse() (response *ModifySmsTemplateResponse) {
    response = &ModifySmsTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySmsTemplate
// This API is used to modify an SMS body template. Please read the [Tencent Cloud SMS Body Template Review Standards](https://intl.cloud.tencent.com/document/product/382/39023?from_cn_redirect=1) before making a modification.
//
// >-  Note: individual users cannot use this API to modify SMS body templates. For more information, please see [Identity Verification Overview](https://intl.cloud.tencent.com/document/product/378/3629?from_cn_redirect=1). If your account identity is individual, you can log in to the console to modify SMS body templates.
//
// >- Modifications can be made only if the body template status is **pending review** or **rejected**. **Approved** body templates cannot be modified.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_MISSINGTEMPLATETOMODIFY = "FailedOperation.MissingTemplateToModify"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  FAILEDOPERATION_TEMPLATEALREADYPASSEDCHECK = "FailedOperation.TemplateAlreadyPassedCheck"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMETERS = "InvalidParameter.InvalidParameters"
//  INVALIDPARAMETERVALUE_INVALIDINTERNATIONAL = "InvalidParameterValue.InvalidInternational"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppidIsDisabled"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) ModifySmsTemplate(request *ModifySmsTemplateRequest) (response *ModifySmsTemplateResponse, err error) {
    return c.ModifySmsTemplateWithContext(context.Background(), request)
}

// ModifySmsTemplate
// This API is used to modify an SMS body template. Please read the [Tencent Cloud SMS Body Template Review Standards](https://intl.cloud.tencent.com/document/product/382/39023?from_cn_redirect=1) before making a modification.
//
// >-  Note: individual users cannot use this API to modify SMS body templates. For more information, please see [Identity Verification Overview](https://intl.cloud.tencent.com/document/product/378/3629?from_cn_redirect=1). If your account identity is individual, you can log in to the console to modify SMS body templates.
//
// >- Modifications can be made only if the body template status is **pending review** or **rejected**. **Approved** body templates cannot be modified.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_MISSINGTEMPLATETOMODIFY = "FailedOperation.MissingTemplateToModify"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  FAILEDOPERATION_TEMPLATEALREADYPASSEDCHECK = "FailedOperation.TemplateAlreadyPassedCheck"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMETERS = "InvalidParameter.InvalidParameters"
//  INVALIDPARAMETERVALUE_INVALIDINTERNATIONAL = "InvalidParameterValue.InvalidInternational"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppidIsDisabled"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) ModifySmsTemplateWithContext(ctx context.Context, request *ModifySmsTemplateRequest) (response *ModifySmsTemplateResponse, err error) {
    if request == nil {
        request = NewModifySmsTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySmsTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySmsTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewPullSmsReplyStatusRequest() (request *PullSmsReplyStatusRequest) {
    request = &PullSmsReplyStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sms", APIVersion, "PullSmsReplyStatus")
    
    
    return
}

func NewPullSmsReplyStatusResponse() (response *PullSmsReplyStatusResponse) {
    response = &PullSmsReplyStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PullSmsReplyStatus
// This API is used to pull SMS reply status.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PHONENUMBERONBLACKLIST = "FailedOperation.PhoneNumberOnBlacklist"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppidNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppidIsDisabled"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppidVerifyFail"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) PullSmsReplyStatus(request *PullSmsReplyStatusRequest) (response *PullSmsReplyStatusResponse, err error) {
    return c.PullSmsReplyStatusWithContext(context.Background(), request)
}

// PullSmsReplyStatus
// This API is used to pull SMS reply status.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PHONENUMBERONBLACKLIST = "FailedOperation.PhoneNumberOnBlacklist"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppidNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppidIsDisabled"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppidVerifyFail"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) PullSmsReplyStatusWithContext(ctx context.Context, request *PullSmsReplyStatusRequest) (response *PullSmsReplyStatusResponse, err error) {
    if request == nil {
        request = NewPullSmsReplyStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PullSmsReplyStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewPullSmsReplyStatusResponse()
    err = c.Send(request, response)
    return
}

func NewPullSmsReplyStatusByPhoneNumberRequest() (request *PullSmsReplyStatusByPhoneNumberRequest) {
    request = &PullSmsReplyStatusByPhoneNumberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sms", APIVersion, "PullSmsReplyStatusByPhoneNumber")
    
    
    return
}

func NewPullSmsReplyStatusByPhoneNumberResponse() (response *PullSmsReplyStatusByPhoneNumberResponse) {
    response = &PullSmsReplyStatusByPhoneNumberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PullSmsReplyStatusByPhoneNumber
// This API is used to pull SMS reply status for one single number.
//
// error code that may be returned:
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PHONENUMBERINBLACKLIST = "FailedOperation.PhoneNumberInBlacklist"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_INVALIDSTARTTIME = "InvalidParameterValue.InvalidStartTime"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppidNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppidIsDisabled"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppidVerifyFail"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) PullSmsReplyStatusByPhoneNumber(request *PullSmsReplyStatusByPhoneNumberRequest) (response *PullSmsReplyStatusByPhoneNumberResponse, err error) {
    return c.PullSmsReplyStatusByPhoneNumberWithContext(context.Background(), request)
}

// PullSmsReplyStatusByPhoneNumber
// This API is used to pull SMS reply status for one single number.
//
// error code that may be returned:
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PHONENUMBERINBLACKLIST = "FailedOperation.PhoneNumberInBlacklist"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_INVALIDSTARTTIME = "InvalidParameterValue.InvalidStartTime"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppidNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppidIsDisabled"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppidVerifyFail"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) PullSmsReplyStatusByPhoneNumberWithContext(ctx context.Context, request *PullSmsReplyStatusByPhoneNumberRequest) (response *PullSmsReplyStatusByPhoneNumberResponse, err error) {
    if request == nil {
        request = NewPullSmsReplyStatusByPhoneNumberRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PullSmsReplyStatusByPhoneNumber require credential")
    }

    request.SetContext(ctx)
    
    response = NewPullSmsReplyStatusByPhoneNumberResponse()
    err = c.Send(request, response)
    return
}

func NewPullSmsSendStatusRequest() (request *PullSmsSendStatusRequest) {
    request = &PullSmsSendStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sms", APIVersion, "PullSmsSendStatus")
    
    
    return
}

func NewPullSmsSendStatusResponse() (response *PullSmsSendStatusResponse) {
    response = &PullSmsSendStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PullSmsSendStatus
// This API is used to pull SMS delivery status.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PHONENUMBERONBLACKLIST = "FailedOperation.PhoneNumberOnBlacklist"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppidNotExist"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppidIsDisabled"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppidVerifyFail"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) PullSmsSendStatus(request *PullSmsSendStatusRequest) (response *PullSmsSendStatusResponse, err error) {
    return c.PullSmsSendStatusWithContext(context.Background(), request)
}

// PullSmsSendStatus
// This API is used to pull SMS delivery status.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PHONENUMBERONBLACKLIST = "FailedOperation.PhoneNumberOnBlacklist"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppidNotExist"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppidIsDisabled"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppidVerifyFail"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) PullSmsSendStatusWithContext(ctx context.Context, request *PullSmsSendStatusRequest) (response *PullSmsSendStatusResponse, err error) {
    if request == nil {
        request = NewPullSmsSendStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PullSmsSendStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewPullSmsSendStatusResponse()
    err = c.Send(request, response)
    return
}

func NewPullSmsSendStatusByPhoneNumberRequest() (request *PullSmsSendStatusByPhoneNumberRequest) {
    request = &PullSmsSendStatusByPhoneNumberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sms", APIVersion, "PullSmsSendStatusByPhoneNumber")
    
    
    return
}

func NewPullSmsSendStatusByPhoneNumberResponse() (response *PullSmsSendStatusByPhoneNumberResponse) {
    response = &PullSmsSendStatusByPhoneNumberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PullSmsSendStatusByPhoneNumber
// This API is used to pull SMS delivery status for one single number.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PHONENUMBERONBLACKLIST = "FailedOperation.PhoneNumberOnBlacklist"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_INVALIDSTARTTIME = "InvalidParameterValue.InvalidStartTime"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppidNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppidIsDisabled"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppidVerifyFail"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) PullSmsSendStatusByPhoneNumber(request *PullSmsSendStatusByPhoneNumberRequest) (response *PullSmsSendStatusByPhoneNumberResponse, err error) {
    return c.PullSmsSendStatusByPhoneNumberWithContext(context.Background(), request)
}

// PullSmsSendStatusByPhoneNumber
// This API is used to pull SMS delivery status for one single number.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PHONENUMBERONBLACKLIST = "FailedOperation.PhoneNumberOnBlacklist"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_INVALIDSTARTTIME = "InvalidParameterValue.InvalidStartTime"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppidNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppidIsDisabled"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppidVerifyFail"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) PullSmsSendStatusByPhoneNumberWithContext(ctx context.Context, request *PullSmsSendStatusByPhoneNumberRequest) (response *PullSmsSendStatusByPhoneNumberResponse, err error) {
    if request == nil {
        request = NewPullSmsSendStatusByPhoneNumberRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PullSmsSendStatusByPhoneNumber require credential")
    }

    request.SetContext(ctx)
    
    response = NewPullSmsSendStatusByPhoneNumberResponse()
    err = c.Send(request, response)
    return
}

func NewSendSmsRequest() (request *SendSmsRequest) {
    request = &SendSmsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sms", APIVersion, "SendSms")
    
    
    return
}

func NewSendSmsResponse() (response *SendSmsResponse) {
    response = &SendSmsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SendSms
// This API is used to send SMS verification codes, notification, or marketing messages to users.
//
// 
//
// error code that may be returned:
//  FAILEDOPERATION_CONTAINSENSITIVEWORD = "FailedOperation.ContainSensitiveWord"
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_INSUFFICIENTBALANCEINSMSPACKAGE = "FailedOperation.InsufficientBalanceInSmsPackage"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_MARKETINGSENDTIMECONSTRAINT = "FailedOperation.MarketingSendTimeConstraint"
//  FAILEDOPERATION_PHONENUMBERINBLACKLIST = "FailedOperation.PhoneNumberInBlacklist"
//  FAILEDOPERATION_SIGNATUREINCORRECTORUNAPPROVED = "FailedOperation.SignatureIncorrectOrUnapproved"
//  FAILEDOPERATION_TEMPLATEINCORRECTORUNAPPROVED = "FailedOperation.TemplateIncorrectOrUnapproved"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CONTENTLENGTHLIMIT = "InvalidParameterValue.ContentLengthLimit"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_PROHIBITEDUSEURLINTEMPLATEPARAMETER = "InvalidParameterValue.ProhibitedUseUrlInTemplateParameter"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppidNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEPARAMETERFORMATERROR = "InvalidParameterValue.TemplateParameterFormatError"
//  INVALIDPARAMETERVALUE_TEMPLATEPARAMETERLENGTHLIMIT = "InvalidParameterValue.TemplateParameterLengthLimit"
//  LIMITEXCEEDED_APPDAILYLIMIT = "LimitExceeded.AppDailyLimit"
//  LIMITEXCEEDED_DAILYLIMIT = "LimitExceeded.DailyLimit"
//  LIMITEXCEEDED_DELIVERYFREQUENCYLIMIT = "LimitExceeded.DeliveryFrequencyLimit"
//  LIMITEXCEEDED_PHONENUMBERCOUNTLIMIT = "LimitExceeded.PhoneNumberCountLimit"
//  LIMITEXCEEDED_PHONENUMBERDAILYLIMIT = "LimitExceeded.PhoneNumberDailyLimit"
//  LIMITEXCEEDED_PHONENUMBERONEHOURLIMIT = "LimitExceeded.PhoneNumberOneHourLimit"
//  LIMITEXCEEDED_PHONENUMBERSAMECONTENTDAILYLIMIT = "LimitExceeded.PhoneNumberSameContentDailyLimit"
//  LIMITEXCEEDED_PHONENUMBERTHIRTYSECONDLIMIT = "LimitExceeded.PhoneNumberThirtySecondLimit"
//  MISSINGPARAMETER_EMPTYPHONENUMBERSET = "MissingParameter.EmptyPhoneNumberSet"
//  UNAUTHORIZEDOPERATION_INDIVIDUALUSERMARKETINGSMSPERMISSIONDENY = "UnauthorizedOperation.IndividualUserMarketingSmsPermissionDeny"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppidIsDisabled"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppidVerifyFail"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
//  UNSUPPORTEDOPERATION_CONTAINDOMESTICANDINTERNATIONALPHONENUMBER = "UnsupportedOperation.ContainDomesticAndInternationalPhoneNumber"
//  UNSUPPORTEDOPERATION_UNSUPORTEDREGION = "UnsupportedOperation.UnsuportedRegion"
func (c *Client) SendSms(request *SendSmsRequest) (response *SendSmsResponse, err error) {
    return c.SendSmsWithContext(context.Background(), request)
}

// SendSms
// This API is used to send SMS verification codes, notification, or marketing messages to users.
//
// 
//
// error code that may be returned:
//  FAILEDOPERATION_CONTAINSENSITIVEWORD = "FailedOperation.ContainSensitiveWord"
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_INSUFFICIENTBALANCEINSMSPACKAGE = "FailedOperation.InsufficientBalanceInSmsPackage"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_MARKETINGSENDTIMECONSTRAINT = "FailedOperation.MarketingSendTimeConstraint"
//  FAILEDOPERATION_PHONENUMBERINBLACKLIST = "FailedOperation.PhoneNumberInBlacklist"
//  FAILEDOPERATION_SIGNATUREINCORRECTORUNAPPROVED = "FailedOperation.SignatureIncorrectOrUnapproved"
//  FAILEDOPERATION_TEMPLATEINCORRECTORUNAPPROVED = "FailedOperation.TemplateIncorrectOrUnapproved"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CONTENTLENGTHLIMIT = "InvalidParameterValue.ContentLengthLimit"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_PROHIBITEDUSEURLINTEMPLATEPARAMETER = "InvalidParameterValue.ProhibitedUseUrlInTemplateParameter"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppidNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEPARAMETERFORMATERROR = "InvalidParameterValue.TemplateParameterFormatError"
//  INVALIDPARAMETERVALUE_TEMPLATEPARAMETERLENGTHLIMIT = "InvalidParameterValue.TemplateParameterLengthLimit"
//  LIMITEXCEEDED_APPDAILYLIMIT = "LimitExceeded.AppDailyLimit"
//  LIMITEXCEEDED_DAILYLIMIT = "LimitExceeded.DailyLimit"
//  LIMITEXCEEDED_DELIVERYFREQUENCYLIMIT = "LimitExceeded.DeliveryFrequencyLimit"
//  LIMITEXCEEDED_PHONENUMBERCOUNTLIMIT = "LimitExceeded.PhoneNumberCountLimit"
//  LIMITEXCEEDED_PHONENUMBERDAILYLIMIT = "LimitExceeded.PhoneNumberDailyLimit"
//  LIMITEXCEEDED_PHONENUMBERONEHOURLIMIT = "LimitExceeded.PhoneNumberOneHourLimit"
//  LIMITEXCEEDED_PHONENUMBERSAMECONTENTDAILYLIMIT = "LimitExceeded.PhoneNumberSameContentDailyLimit"
//  LIMITEXCEEDED_PHONENUMBERTHIRTYSECONDLIMIT = "LimitExceeded.PhoneNumberThirtySecondLimit"
//  MISSINGPARAMETER_EMPTYPHONENUMBERSET = "MissingParameter.EmptyPhoneNumberSet"
//  UNAUTHORIZEDOPERATION_INDIVIDUALUSERMARKETINGSMSPERMISSIONDENY = "UnauthorizedOperation.IndividualUserMarketingSmsPermissionDeny"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppidIsDisabled"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppidVerifyFail"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
//  UNSUPPORTEDOPERATION_CONTAINDOMESTICANDINTERNATIONALPHONENUMBER = "UnsupportedOperation.ContainDomesticAndInternationalPhoneNumber"
//  UNSUPPORTEDOPERATION_UNSUPORTEDREGION = "UnsupportedOperation.UnsuportedRegion"
func (c *Client) SendSmsWithContext(ctx context.Context, request *SendSmsRequest) (response *SendSmsResponse, err error) {
    if request == nil {
        request = NewSendSmsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendSms require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendSmsResponse()
    err = c.Send(request, response)
    return
}

func NewSendStatusStatisticsRequest() (request *SendStatusStatisticsRequest) {
    request = &SendStatusStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sms", APIVersion, "SendStatusStatistics")
    
    
    return
}

func NewSendStatusStatisticsResponse() (response *SendStatusStatisticsResponse) {
    response = &SendStatusStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SendStatusStatistics
// This API is used to collect statistics on SMS sent by users.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PHONENUMBERINBLACKLIST = "FailedOperation.PhoneNumberInBlacklist"
//  FAILEDOPERATION_TEMPLATEINCORRECTORUNAPPROVED = "FailedOperation.TemplateIncorrectOrUnapproved"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppidNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppidIsDisabled"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppidVerifyFail"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) SendStatusStatistics(request *SendStatusStatisticsRequest) (response *SendStatusStatisticsResponse, err error) {
    return c.SendStatusStatisticsWithContext(context.Background(), request)
}

// SendStatusStatistics
// This API is used to collect statistics on SMS sent by users.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PHONENUMBERINBLACKLIST = "FailedOperation.PhoneNumberInBlacklist"
//  FAILEDOPERATION_TEMPLATEINCORRECTORUNAPPROVED = "FailedOperation.TemplateIncorrectOrUnapproved"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppidNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppidIsDisabled"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppidVerifyFail"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) SendStatusStatisticsWithContext(ctx context.Context, request *SendStatusStatisticsRequest) (response *SendStatusStatisticsResponse, err error) {
    if request == nil {
        request = NewSendStatusStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendStatusStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendStatusStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewSmsPackagesStatisticsRequest() (request *SmsPackagesStatisticsRequest) {
    request = &SmsPackagesStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sms", APIVersion, "SmsPackagesStatistics")
    
    
    return
}

func NewSmsPackagesStatisticsResponse() (response *SmsPackagesStatisticsResponse) {
    response = &SmsPackagesStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SmsPackagesStatistics
// This API is used to collect statistics on user's packages.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PHONENUMBERINBLACKLIST = "FailedOperation.PhoneNumberInBlacklist"
//  FAILEDOPERATION_SIGNATUREINCORRECTORUNAPPROVED = "FailedOperation.SignatureIncorrectOrUnapproved"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppidNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppidIsDisabled"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppidVerifyFail"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) SmsPackagesStatistics(request *SmsPackagesStatisticsRequest) (response *SmsPackagesStatisticsResponse, err error) {
    return c.SmsPackagesStatisticsWithContext(context.Background(), request)
}

// SmsPackagesStatistics
// This API is used to collect statistics on user's packages.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PHONENUMBERINBLACKLIST = "FailedOperation.PhoneNumberInBlacklist"
//  FAILEDOPERATION_SIGNATUREINCORRECTORUNAPPROVED = "FailedOperation.SignatureIncorrectOrUnapproved"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppidNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppidIsDisabled"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppidVerifyFail"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) SmsPackagesStatisticsWithContext(ctx context.Context, request *SmsPackagesStatisticsRequest) (response *SmsPackagesStatisticsResponse, err error) {
    if request == nil {
        request = NewSmsPackagesStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SmsPackagesStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewSmsPackagesStatisticsResponse()
    err = c.Send(request, response)
    return
}
