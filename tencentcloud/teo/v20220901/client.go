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

package v20220901

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2022-09-01"

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


func NewApplyFreeCertificateRequest() (request *ApplyFreeCertificateRequest) {
    request = &ApplyFreeCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ApplyFreeCertificate")
    
    
    return
}

func NewApplyFreeCertificateResponse() (response *ApplyFreeCertificateResponse) {
    response = &ApplyFreeCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ApplyFreeCertificate
// When applying for a free certificate, if you need to apply through DNS delegation verification or file verification, you can call this API to initiate the certificate application and get the corresponding verification content based on the application method. The order in which to call the API is as follows:
//
// Step 1: Call ApplyFreeCertificate to specify the verification method for the free certificate application and obtain the verification content.
//
// Step 2: Configure the content for the corresponding domain as required for verification.
//
// Step 3: Call CheckFreeCertificateVerification for verification. The free certificate application is completed after verification passes.
//
// Step 4: Call ModifyHostsCertificate to configure the domain name certificate with an EdgeOne free certificate.
//
// 
//
// For the introduction to application methods, see the referenced document: Free Certificate Application Description (https://www.tencentcloud.com/document/product/1552/90437?from_cn_redirect=1) 
//
// Description:
//
// -Only the CNAME access mode can call this API to specify the free certificate application method. The NS/DNSPod hosting access mode uses automatic validation to apply for a free certificate, with no need to call this API.
//
// -If you need to switch the verification method for a free certificate, you can call this API again to update it by changing the VerificationMethod field.
//
// -A domain name can only apply for one free certificate. After calling this API, the backend will trigger the free certificate application task. You need to complete the domain name verification info configuration within 2 days, and then complete certificate authentication.
func (c *Client) ApplyFreeCertificate(request *ApplyFreeCertificateRequest) (response *ApplyFreeCertificateResponse, err error) {
    return c.ApplyFreeCertificateWithContext(context.Background(), request)
}

// ApplyFreeCertificate
// When applying for a free certificate, if you need to apply through DNS delegation verification or file verification, you can call this API to initiate the certificate application and get the corresponding verification content based on the application method. The order in which to call the API is as follows:
//
// Step 1: Call ApplyFreeCertificate to specify the verification method for the free certificate application and obtain the verification content.
//
// Step 2: Configure the content for the corresponding domain as required for verification.
//
// Step 3: Call CheckFreeCertificateVerification for verification. The free certificate application is completed after verification passes.
//
// Step 4: Call ModifyHostsCertificate to configure the domain name certificate with an EdgeOne free certificate.
//
// 
//
// For the introduction to application methods, see the referenced document: Free Certificate Application Description (https://www.tencentcloud.com/document/product/1552/90437?from_cn_redirect=1) 
//
// Description:
//
// -Only the CNAME access mode can call this API to specify the free certificate application method. The NS/DNSPod hosting access mode uses automatic validation to apply for a free certificate, with no need to call this API.
//
// -If you need to switch the verification method for a free certificate, you can call this API again to update it by changing the VerificationMethod field.
//
// -A domain name can only apply for one free certificate. After calling this API, the backend will trigger the free certificate application task. You need to complete the domain name verification info configuration within 2 days, and then complete certificate authentication.
func (c *Client) ApplyFreeCertificateWithContext(ctx context.Context, request *ApplyFreeCertificateRequest) (response *ApplyFreeCertificateResponse, err error) {
    if request == nil {
        request = NewApplyFreeCertificateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ApplyFreeCertificate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyFreeCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyFreeCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewBindSecurityTemplateToEntityRequest() (request *BindSecurityTemplateToEntityRequest) {
    request = &BindSecurityTemplateToEntityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "BindSecurityTemplateToEntity")
    
    
    return
}

func NewBindSecurityTemplateToEntityResponse() (response *BindSecurityTemplateToEntityResponse) {
    response = &BindSecurityTemplateToEntityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindSecurityTemplateToEntity
// Operate a security policy template. This API is used to bind or rebind a domain name to a designated policy template, or unbind it from a designated policy template.
//
// error code that may be returned:
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BindSecurityTemplateToEntity(request *BindSecurityTemplateToEntityRequest) (response *BindSecurityTemplateToEntityResponse, err error) {
    return c.BindSecurityTemplateToEntityWithContext(context.Background(), request)
}

// BindSecurityTemplateToEntity
// Operate a security policy template. This API is used to bind or rebind a domain name to a designated policy template, or unbind it from a designated policy template.
//
// error code that may be returned:
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BindSecurityTemplateToEntityWithContext(ctx context.Context, request *BindSecurityTemplateToEntityRequest) (response *BindSecurityTemplateToEntityResponse, err error) {
    if request == nil {
        request = NewBindSecurityTemplateToEntityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "BindSecurityTemplateToEntity")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindSecurityTemplateToEntity require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindSecurityTemplateToEntityResponse()
    err = c.Send(request, response)
    return
}

func NewBindSharedCNAMERequest() (request *BindSharedCNAMERequest) {
    request = &BindSharedCNAMERequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "BindSharedCNAME")
    
    
    return
}

func NewBindSharedCNAMEResponse() (response *BindSharedCNAMEResponse) {
    response = &BindSharedCNAMEResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindSharedCNAME
// Used to bind or unbind a shared CNAME for an acceleration domain name. This feature is in beta test.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) BindSharedCNAME(request *BindSharedCNAMERequest) (response *BindSharedCNAMEResponse, err error) {
    return c.BindSharedCNAMEWithContext(context.Background(), request)
}

// BindSharedCNAME
// Used to bind or unbind a shared CNAME for an acceleration domain name. This feature is in beta test.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) BindSharedCNAMEWithContext(ctx context.Context, request *BindSharedCNAMERequest) (response *BindSharedCNAMEResponse, err error) {
    if request == nil {
        request = NewBindSharedCNAMERequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "BindSharedCNAME")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindSharedCNAME require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindSharedCNAMEResponse()
    err = c.Send(request, response)
    return
}

func NewBindZoneToPlanRequest() (request *BindZoneToPlanRequest) {
    request = &BindZoneToPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "BindZoneToPlan")
    
    
    return
}

func NewBindZoneToPlanResponse() (response *BindZoneToPlanResponse) {
    response = &BindZoneToPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindZoneToPlan
// Bind the selected site not bound to a package to an existing package.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PLANNOTFOUND = "InvalidParameter.PlanNotFound"
//  INVALIDPARAMETER_ZONEHASBEENBOUND = "InvalidParameter.ZoneHasBeenBound"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) BindZoneToPlan(request *BindZoneToPlanRequest) (response *BindZoneToPlanResponse, err error) {
    return c.BindZoneToPlanWithContext(context.Background(), request)
}

// BindZoneToPlan
// Bind the selected site not bound to a package to an existing package.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PLANNOTFOUND = "InvalidParameter.PlanNotFound"
//  INVALIDPARAMETER_ZONEHASBEENBOUND = "InvalidParameter.ZoneHasBeenBound"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) BindZoneToPlanWithContext(ctx context.Context, request *BindZoneToPlanRequest) (response *BindZoneToPlanResponse, err error) {
    if request == nil {
        request = NewBindZoneToPlanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "BindZoneToPlan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindZoneToPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindZoneToPlanResponse()
    err = c.Send(request, response)
    return
}

func NewCheckCnameStatusRequest() (request *CheckCnameStatusRequest) {
    request = &CheckCnameStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CheckCnameStatus")
    
    
    return
}

func NewCheckCnameStatusResponse() (response *CheckCnameStatusResponse) {
    response = &CheckCnameStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckCnameStatus
// When the site access type is CNAME, all domain names under the site must complete CNAME record configuration as specified by the CNAME domain assigned by EdgeOne.
//
// This API is used to obtain the CNAME domain name assigned by EdgeOne for the integrated domain name, and to validate the CNAME configuration status of the integrated domain name.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) CheckCnameStatus(request *CheckCnameStatusRequest) (response *CheckCnameStatusResponse, err error) {
    return c.CheckCnameStatusWithContext(context.Background(), request)
}

// CheckCnameStatus
// When the site access type is CNAME, all domain names under the site must complete CNAME record configuration as specified by the CNAME domain assigned by EdgeOne.
//
// This API is used to obtain the CNAME domain name assigned by EdgeOne for the integrated domain name, and to validate the CNAME configuration status of the integrated domain name.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) CheckCnameStatusWithContext(ctx context.Context, request *CheckCnameStatusRequest) (response *CheckCnameStatusResponse, err error) {
    if request == nil {
        request = NewCheckCnameStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CheckCnameStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckCnameStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckCnameStatusResponse()
    err = c.Send(request, response)
    return
}

func NewCheckFreeCertificateVerificationRequest() (request *CheckFreeCertificateVerificationRequest) {
    request = &CheckFreeCertificateVerificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CheckFreeCertificateVerification")
    
    
    return
}

func NewCheckFreeCertificateVerificationResponse() (response *CheckFreeCertificateVerificationResponse) {
    response = &CheckFreeCertificateVerificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckFreeCertificateVerification
// This API is used to verify a free certificate and obtain the application result. If verification is passed, you can query the free certificate information for the corresponding domain through this API. If the application failed, this API will return the corresponding failure description.
//
// After triggering the free certificate application API (https://www.tencentcloud.com/document/product/1552/90437?from_cn_redirect=1), you can check the free certificate application result through this API. Once the free certificate application is successful, you need to configure it through the domain certificate configuration API (https://www.tencentcloud.com/document/product/1552/80764?from_cn_redirect=1) to deploy the free certificate to the acceleration domain.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) CheckFreeCertificateVerification(request *CheckFreeCertificateVerificationRequest) (response *CheckFreeCertificateVerificationResponse, err error) {
    return c.CheckFreeCertificateVerificationWithContext(context.Background(), request)
}

// CheckFreeCertificateVerification
// This API is used to verify a free certificate and obtain the application result. If verification is passed, you can query the free certificate information for the corresponding domain through this API. If the application failed, this API will return the corresponding failure description.
//
// After triggering the free certificate application API (https://www.tencentcloud.com/document/product/1552/90437?from_cn_redirect=1), you can check the free certificate application result through this API. Once the free certificate application is successful, you need to configure it through the domain certificate configuration API (https://www.tencentcloud.com/document/product/1552/80764?from_cn_redirect=1) to deploy the free certificate to the acceleration domain.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) CheckFreeCertificateVerificationWithContext(ctx context.Context, request *CheckFreeCertificateVerificationRequest) (response *CheckFreeCertificateVerificationResponse, err error) {
    if request == nil {
        request = NewCheckFreeCertificateVerificationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CheckFreeCertificateVerification")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckFreeCertificateVerification require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckFreeCertificateVerificationResponse()
    err = c.Send(request, response)
    return
}

func NewConfirmMultiPathGatewayOriginACLRequest() (request *ConfirmMultiPathGatewayOriginACLRequest) {
    request = &ConfirmMultiPathGatewayOriginACLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ConfirmMultiPathGatewayOriginACL")
    
    
    return
}

func NewConfirmMultiPathGatewayOriginACLResponse() (response *ConfirmMultiPathGatewayOriginACLResponse) {
    response = &ConfirmMultiPathGatewayOriginACLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ConfirmMultiPathGatewayOriginACL
// This API is used to confirm that the latest origin-pull IP ranges have been updated to the origin server firewall when the origin-pull IP ranges of the multi-channel security acceleration gateway change.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ConfirmMultiPathGatewayOriginACL(request *ConfirmMultiPathGatewayOriginACLRequest) (response *ConfirmMultiPathGatewayOriginACLResponse, err error) {
    return c.ConfirmMultiPathGatewayOriginACLWithContext(context.Background(), request)
}

// ConfirmMultiPathGatewayOriginACL
// This API is used to confirm that the latest origin-pull IP ranges have been updated to the origin server firewall when the origin-pull IP ranges of the multi-channel security acceleration gateway change.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ConfirmMultiPathGatewayOriginACLWithContext(ctx context.Context, request *ConfirmMultiPathGatewayOriginACLRequest) (response *ConfirmMultiPathGatewayOriginACLResponse, err error) {
    if request == nil {
        request = NewConfirmMultiPathGatewayOriginACLRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ConfirmMultiPathGatewayOriginACL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ConfirmMultiPathGatewayOriginACL require credential")
    }

    request.SetContext(ctx)
    
    response = NewConfirmMultiPathGatewayOriginACLResponse()
    err = c.Send(request, response)
    return
}

func NewConfirmOriginACLUpdateRequest() (request *ConfirmOriginACLUpdateRequest) {
    request = &ConfirmOriginACLUpdateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ConfirmOriginACLUpdate")
    
    
    return
}

func NewConfirmOriginACLUpdateResponse() (response *ConfirmOriginACLUpdateResponse) {
    response = &ConfirmOriginACLUpdateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ConfirmOriginACLUpdate
// This API is used to confirm that the latest origin-pull IP ranges have been updated to the origin server firewall when the origin-pull IP ranges change. After confirming the update to the latest origin-pull IP ranges, related change notifications will stop being pushed.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_LATESTVERSIONNOW = "OperationDenied.LatestVersionNow"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ConfirmOriginACLUpdate(request *ConfirmOriginACLUpdateRequest) (response *ConfirmOriginACLUpdateResponse, err error) {
    return c.ConfirmOriginACLUpdateWithContext(context.Background(), request)
}

// ConfirmOriginACLUpdate
// This API is used to confirm that the latest origin-pull IP ranges have been updated to the origin server firewall when the origin-pull IP ranges change. After confirming the update to the latest origin-pull IP ranges, related change notifications will stop being pushed.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_LATESTVERSIONNOW = "OperationDenied.LatestVersionNow"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ConfirmOriginACLUpdateWithContext(ctx context.Context, request *ConfirmOriginACLUpdateRequest) (response *ConfirmOriginACLUpdateResponse, err error) {
    if request == nil {
        request = NewConfirmOriginACLUpdateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ConfirmOriginACLUpdate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ConfirmOriginACLUpdate require credential")
    }

    request.SetContext(ctx)
    
    response = NewConfirmOriginACLUpdateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAccelerationDomainRequest() (request *CreateAccelerationDomainRequest) {
    request = &CreateAccelerationDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateAccelerationDomain")
    
    
    return
}

func NewCreateAccelerationDomainResponse() (response *CreateAccelerationDomainResponse) {
    response = &CreateAccelerationDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAccelerationDomain
// After creating a site, you can create an acceleration domain name through this API. 
//
// 
//
// In CNAME mode, if site ownership verification is ongoing, this API will return domain ownership verification info. You can verify domain ownership separately. For details, see [Site/Domain Name Ownership Verification](https://www.tencentcloud.com/document/product/1552/70789?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INVALIDPARAMETER_ADVANCEDORIGINROUTINGNOTSUPPORTNEWLYADDEDHOST = "InvalidParameter.AdvancedOriginRoutingNotSupportNewlyAddedHost"
//  INVALIDPARAMETER_CONFLICTHOSTORIGIN = "InvalidParameter.ConflictHostOrigin"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_INVALIDACCELERATETYPE = "InvalidParameter.InvalidAccelerateType"
//  INVALIDPARAMETER_INVALIDAWSPRIVATEACCESS = "InvalidParameter.InvalidAwsPrivateAccess"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDCLIENTIPORIGIN = "InvalidParameter.InvalidClientIpOrigin"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_INVALIDPRIVATEACCESSPARAMS = "InvalidParameter.InvalidPrivateAccessParams"
//  INVALIDPARAMETER_INVALIDPRIVATEACCESSSWITCH = "InvalidParameter.InvalidPrivateAccessSwitch"
//  INVALIDPARAMETER_INVALIDQUICBILLING = "InvalidParameter.InvalidQuicBilling"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGCLIENTIP = "InvalidParameter.InvalidStandardDebugClientIp"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_SPACENOTBINDORIGIN = "InvalidParameter.SpaceNotBindOrigin"
//  INVALIDPARAMETER_ZONEISGRAYPUBLISHING = "InvalidParameter.ZoneIsGrayPublishing"
//  INVALIDPARAMETERVALUE_CONFLICTRECORD = "InvalidParameterValue.ConflictRecord"
//  INVALIDPARAMETERVALUE_CONFLICTWITHDNSSEC = "InvalidParameterValue.ConflictWithDNSSEC"
//  INVALIDPARAMETERVALUE_CONFLICTWITHDOMAIN = "InvalidParameterValue.ConflictWithDomain"
//  INVALIDPARAMETERVALUE_CONFLICTWITHNSRECORD = "InvalidParameterValue.ConflictWithNSRecord"
//  INVALIDPARAMETERVALUE_CONTENTSAMEASNAME = "InvalidParameterValue.ContentSameAsName"
//  INVALIDPARAMETERVALUE_DOMAINNOTMATCHZONE = "InvalidParameterValue.DomainNotMatchZone"
//  INVALIDPARAMETERVALUE_INVALIDDNSCONTENT = "InvalidParameterValue.InvalidDNSContent"
//  INVALIDPARAMETERVALUE_INVALIDDNSNAME = "InvalidParameterValue.InvalidDNSName"
//  INVALIDPARAMETERVALUE_INVALIDDOMAINNAME = "InvalidParameterValue.InvalidDomainName"
//  INVALIDPARAMETERVALUE_INVALIDPROXYORIGIN = "InvalidParameterValue.InvalidProxyOrigin"
//  INVALIDPARAMETERVALUE_INVALIDSITEFAILOVERUNSUPPORTED = "InvalidParameterValue.InvalidSiteFailoverUnsupported"
//  INVALIDPARAMETERVALUE_ORIGINGROUPNOTEXISTS = "InvalidParameterValue.OriginGroupNotExists"
//  INVALIDPARAMETERVALUE_SITEFAILOVERNOTSUPPORTHOSTORIGINTYPEVOD = "InvalidParameterValue.SiteFailoverNotSupportHostOriginTypeVod"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDDISABLE = "OperationDenied.AccelerateMainlandDisable"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_DOMAINNOICP = "OperationDenied.DomainNoICP"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_INVALIDADVANCEDDEFENSEZONEAREA = "OperationDenied.InvalidAdvancedDefenseZoneArea"
//  OPERATIONDENIED_LOADBALANCESTATUSNOTINONLINE = "OperationDenied.LoadBalanceStatusNotInOnline"
//  OPERATIONDENIED_RECORDISFORBIDDEN = "OperationDenied.RecordIsForbidden"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  OPERATIONDENIED_VERSIONCONTROLISGRAYING = "OperationDenied.VersionControlIsGraying"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_DNSRECORD = "ResourceInUse.DnsRecord"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_DOMAINALREADYEXISTS = "ResourceUnavailable.DomainAlreadyExists"
//  RESOURCESSOLDOUT_L7LACKOFRESOURCES = "ResourcesSoldOut.L7LackOfResources"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) CreateAccelerationDomain(request *CreateAccelerationDomainRequest) (response *CreateAccelerationDomainResponse, err error) {
    return c.CreateAccelerationDomainWithContext(context.Background(), request)
}

// CreateAccelerationDomain
// After creating a site, you can create an acceleration domain name through this API. 
//
// 
//
// In CNAME mode, if site ownership verification is ongoing, this API will return domain ownership verification info. You can verify domain ownership separately. For details, see [Site/Domain Name Ownership Verification](https://www.tencentcloud.com/document/product/1552/70789?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INVALIDPARAMETER_ADVANCEDORIGINROUTINGNOTSUPPORTNEWLYADDEDHOST = "InvalidParameter.AdvancedOriginRoutingNotSupportNewlyAddedHost"
//  INVALIDPARAMETER_CONFLICTHOSTORIGIN = "InvalidParameter.ConflictHostOrigin"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_INVALIDACCELERATETYPE = "InvalidParameter.InvalidAccelerateType"
//  INVALIDPARAMETER_INVALIDAWSPRIVATEACCESS = "InvalidParameter.InvalidAwsPrivateAccess"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDCLIENTIPORIGIN = "InvalidParameter.InvalidClientIpOrigin"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_INVALIDPRIVATEACCESSPARAMS = "InvalidParameter.InvalidPrivateAccessParams"
//  INVALIDPARAMETER_INVALIDPRIVATEACCESSSWITCH = "InvalidParameter.InvalidPrivateAccessSwitch"
//  INVALIDPARAMETER_INVALIDQUICBILLING = "InvalidParameter.InvalidQuicBilling"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGCLIENTIP = "InvalidParameter.InvalidStandardDebugClientIp"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_SPACENOTBINDORIGIN = "InvalidParameter.SpaceNotBindOrigin"
//  INVALIDPARAMETER_ZONEISGRAYPUBLISHING = "InvalidParameter.ZoneIsGrayPublishing"
//  INVALIDPARAMETERVALUE_CONFLICTRECORD = "InvalidParameterValue.ConflictRecord"
//  INVALIDPARAMETERVALUE_CONFLICTWITHDNSSEC = "InvalidParameterValue.ConflictWithDNSSEC"
//  INVALIDPARAMETERVALUE_CONFLICTWITHDOMAIN = "InvalidParameterValue.ConflictWithDomain"
//  INVALIDPARAMETERVALUE_CONFLICTWITHNSRECORD = "InvalidParameterValue.ConflictWithNSRecord"
//  INVALIDPARAMETERVALUE_CONTENTSAMEASNAME = "InvalidParameterValue.ContentSameAsName"
//  INVALIDPARAMETERVALUE_DOMAINNOTMATCHZONE = "InvalidParameterValue.DomainNotMatchZone"
//  INVALIDPARAMETERVALUE_INVALIDDNSCONTENT = "InvalidParameterValue.InvalidDNSContent"
//  INVALIDPARAMETERVALUE_INVALIDDNSNAME = "InvalidParameterValue.InvalidDNSName"
//  INVALIDPARAMETERVALUE_INVALIDDOMAINNAME = "InvalidParameterValue.InvalidDomainName"
//  INVALIDPARAMETERVALUE_INVALIDPROXYORIGIN = "InvalidParameterValue.InvalidProxyOrigin"
//  INVALIDPARAMETERVALUE_INVALIDSITEFAILOVERUNSUPPORTED = "InvalidParameterValue.InvalidSiteFailoverUnsupported"
//  INVALIDPARAMETERVALUE_ORIGINGROUPNOTEXISTS = "InvalidParameterValue.OriginGroupNotExists"
//  INVALIDPARAMETERVALUE_SITEFAILOVERNOTSUPPORTHOSTORIGINTYPEVOD = "InvalidParameterValue.SiteFailoverNotSupportHostOriginTypeVod"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDDISABLE = "OperationDenied.AccelerateMainlandDisable"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_DOMAINNOICP = "OperationDenied.DomainNoICP"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_INVALIDADVANCEDDEFENSEZONEAREA = "OperationDenied.InvalidAdvancedDefenseZoneArea"
//  OPERATIONDENIED_LOADBALANCESTATUSNOTINONLINE = "OperationDenied.LoadBalanceStatusNotInOnline"
//  OPERATIONDENIED_RECORDISFORBIDDEN = "OperationDenied.RecordIsForbidden"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  OPERATIONDENIED_VERSIONCONTROLISGRAYING = "OperationDenied.VersionControlIsGraying"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_DNSRECORD = "ResourceInUse.DnsRecord"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_DOMAINALREADYEXISTS = "ResourceUnavailable.DomainAlreadyExists"
//  RESOURCESSOLDOUT_L7LACKOFRESOURCES = "ResourcesSoldOut.L7LackOfResources"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) CreateAccelerationDomainWithContext(ctx context.Context, request *CreateAccelerationDomainRequest) (response *CreateAccelerationDomainResponse, err error) {
    if request == nil {
        request = NewCreateAccelerationDomainRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreateAccelerationDomain")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAccelerationDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAccelerationDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAliasDomainRequest() (request *CreateAliasDomainRequest) {
    request = &CreateAliasDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateAliasDomain")
    
    
    return
}

func NewCreateAliasDomainResponse() (response *CreateAliasDomainResponse) {
    response = &CreateAliasDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAliasDomain
// This API is used to create an alias domain name.
//
// This feature is only supported in the enterprise plan and is currently in beta test. To use it, please [contact us](https://www.tencentcloud.com/contact-us).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INVALIDPARAMETER_ALIASDOMAINNOTSUPPORTSMCERT = "InvalidParameter.AliasDomainNotSupportSMCert"
//  INVALIDPARAMETER_CERTNOTMATCHDOMAIN = "InvalidParameter.CertNotMatchDomain"
//  INVALIDPARAMETER_INVALIDHTTPSCERTINFO = "InvalidParameter.InvalidHttpsCertInfo"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDALIASDOMAINNAME = "InvalidParameterValue.InvalidAliasDomainName"
//  INVALIDPARAMETERVALUE_INVALIDALIASNAMESUFFIX = "InvalidParameterValue.InvalidAliasNameSuffix"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_DOMAINISBLOCKED = "OperationDenied.DomainIsBlocked"
//  OPERATIONDENIED_DOMAINNOICP = "OperationDenied.DomainNoICP"
//  OPERATIONDENIED_VERSIONCONTROLLOCKED = "OperationDenied.VersionControlLocked"
//  RESOURCEINUSE_ALIASNAME = "ResourceInUse.AliasName"
//  RESOURCEINUSE_ALREADYEXISTSASANACCELERATIONDOMAIN = "ResourceInUse.AlreadyExistsAsAnAccelerationDomain"
//  RESOURCEINUSE_DUPLICATENAME = "ResourceInUse.DuplicateName"
//  RESOURCEINUSE_ZONE = "ResourceInUse.Zone"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_DOMAINALREADYEXISTS = "ResourceUnavailable.DomainAlreadyExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNSUPPORTEDOPERATION_TARGETNAMEORIGINTYPECOS = "UnsupportedOperation.TargetNameOriginTypeCos"
func (c *Client) CreateAliasDomain(request *CreateAliasDomainRequest) (response *CreateAliasDomainResponse, err error) {
    return c.CreateAliasDomainWithContext(context.Background(), request)
}

// CreateAliasDomain
// This API is used to create an alias domain name.
//
// This feature is only supported in the enterprise plan and is currently in beta test. To use it, please [contact us](https://www.tencentcloud.com/contact-us).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INVALIDPARAMETER_ALIASDOMAINNOTSUPPORTSMCERT = "InvalidParameter.AliasDomainNotSupportSMCert"
//  INVALIDPARAMETER_CERTNOTMATCHDOMAIN = "InvalidParameter.CertNotMatchDomain"
//  INVALIDPARAMETER_INVALIDHTTPSCERTINFO = "InvalidParameter.InvalidHttpsCertInfo"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDALIASDOMAINNAME = "InvalidParameterValue.InvalidAliasDomainName"
//  INVALIDPARAMETERVALUE_INVALIDALIASNAMESUFFIX = "InvalidParameterValue.InvalidAliasNameSuffix"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_DOMAINISBLOCKED = "OperationDenied.DomainIsBlocked"
//  OPERATIONDENIED_DOMAINNOICP = "OperationDenied.DomainNoICP"
//  OPERATIONDENIED_VERSIONCONTROLLOCKED = "OperationDenied.VersionControlLocked"
//  RESOURCEINUSE_ALIASNAME = "ResourceInUse.AliasName"
//  RESOURCEINUSE_ALREADYEXISTSASANACCELERATIONDOMAIN = "ResourceInUse.AlreadyExistsAsAnAccelerationDomain"
//  RESOURCEINUSE_DUPLICATENAME = "ResourceInUse.DuplicateName"
//  RESOURCEINUSE_ZONE = "ResourceInUse.Zone"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_DOMAINALREADYEXISTS = "ResourceUnavailable.DomainAlreadyExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNSUPPORTEDOPERATION_TARGETNAMEORIGINTYPECOS = "UnsupportedOperation.TargetNameOriginTypeCos"
func (c *Client) CreateAliasDomainWithContext(ctx context.Context, request *CreateAliasDomainRequest) (response *CreateAliasDomainResponse, err error) {
    if request == nil {
        request = NewCreateAliasDomainRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreateAliasDomain")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAliasDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAliasDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApplicationProxyRequest() (request *CreateApplicationProxyRequest) {
    request = &CreateApplicationProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateApplicationProxy")
    
    
    return
}

func NewCreateApplicationProxyResponse() (response *CreateApplicationProxyResponse) {
    response = &CreateApplicationProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApplicationProxy
// This API is on an earlier version. If you want to call it, switch to the latest version to create layer 4 proxy instances (https://www.tencentcloud.com/document/product/1552/103417?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETER_PROXYNAMEDUPLICATING = "InvalidParameter.ProxyNameDuplicating"
//  INVALIDPARAMETER_PROXYNAMENOTMATCHED = "InvalidParameter.ProxyNameNotMatched"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_USERQUOTALIMITED = "LimitExceeded.UserQuotaLimited"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_PLATTYPEIPACCELERATEMAINLANDNOTSUPPORT = "OperationDenied.PlatTypeIPAccelerateMainlandNotSupport"
//  OPERATIONDENIED_ZONENOTACTIVE = "OperationDenied.ZoneNotActive"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateApplicationProxy(request *CreateApplicationProxyRequest) (response *CreateApplicationProxyResponse, err error) {
    return c.CreateApplicationProxyWithContext(context.Background(), request)
}

// CreateApplicationProxy
// This API is on an earlier version. If you want to call it, switch to the latest version to create layer 4 proxy instances (https://www.tencentcloud.com/document/product/1552/103417?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETER_PROXYNAMEDUPLICATING = "InvalidParameter.ProxyNameDuplicating"
//  INVALIDPARAMETER_PROXYNAMENOTMATCHED = "InvalidParameter.ProxyNameNotMatched"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_USERQUOTALIMITED = "LimitExceeded.UserQuotaLimited"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_PLATTYPEIPACCELERATEMAINLANDNOTSUPPORT = "OperationDenied.PlatTypeIPAccelerateMainlandNotSupport"
//  OPERATIONDENIED_ZONENOTACTIVE = "OperationDenied.ZoneNotActive"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateApplicationProxyWithContext(ctx context.Context, request *CreateApplicationProxyRequest) (response *CreateApplicationProxyResponse, err error) {
    if request == nil {
        request = NewCreateApplicationProxyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreateApplicationProxy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApplicationProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApplicationProxyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApplicationProxyRuleRequest() (request *CreateApplicationProxyRuleRequest) {
    request = &CreateApplicationProxyRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateApplicationProxyRule")
    
    
    return
}

func NewCreateApplicationProxyRuleResponse() (response *CreateApplicationProxyRuleResponse) {
    response = &CreateApplicationProxyRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApplicationProxyRule
// This API is on an earlier version. If you want to call it, switch to the latest version. For details, see [Create Layer 4 Proxy Forwarding Rule].
//
// ](https://www.tencentcloud.com/document/product/1552/103416?from_cn_redirect=1) .
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_INVALIDRULEPROTO = "InvalidParameter.InvalidRuleProto"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETER_RULEORIGINFORMATERROR = "InvalidParameter.RuleOriginFormatError"
//  INVALIDPARAMETER_RULEORIGINMULTIDOMAIN = "InvalidParameter.RuleOriginMultiDomain"
//  INVALIDPARAMETER_RULEORIGINPORTINTEGER = "InvalidParameter.RuleOriginPortInteger"
//  INVALIDPARAMETER_RULEORIGINVALUEERROR = "InvalidParameter.RuleOriginValueError"
//  INVALIDPARAMETER_RULEPORTDUPLICATING = "InvalidParameter.RulePortDuplicating"
//  INVALIDPARAMETER_RULEPORTGROUP = "InvalidParameter.RulePortGroup"
//  INVALIDPARAMETER_RULEPORTINTEGER = "InvalidParameter.RulePortInteger"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateApplicationProxyRule(request *CreateApplicationProxyRuleRequest) (response *CreateApplicationProxyRuleResponse, err error) {
    return c.CreateApplicationProxyRuleWithContext(context.Background(), request)
}

// CreateApplicationProxyRule
// This API is on an earlier version. If you want to call it, switch to the latest version. For details, see [Create Layer 4 Proxy Forwarding Rule].
//
// ](https://www.tencentcloud.com/document/product/1552/103416?from_cn_redirect=1) .
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_INVALIDRULEPROTO = "InvalidParameter.InvalidRuleProto"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETER_RULEORIGINFORMATERROR = "InvalidParameter.RuleOriginFormatError"
//  INVALIDPARAMETER_RULEORIGINMULTIDOMAIN = "InvalidParameter.RuleOriginMultiDomain"
//  INVALIDPARAMETER_RULEORIGINPORTINTEGER = "InvalidParameter.RuleOriginPortInteger"
//  INVALIDPARAMETER_RULEORIGINVALUEERROR = "InvalidParameter.RuleOriginValueError"
//  INVALIDPARAMETER_RULEPORTDUPLICATING = "InvalidParameter.RulePortDuplicating"
//  INVALIDPARAMETER_RULEPORTGROUP = "InvalidParameter.RulePortGroup"
//  INVALIDPARAMETER_RULEPORTINTEGER = "InvalidParameter.RulePortInteger"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateApplicationProxyRuleWithContext(ctx context.Context, request *CreateApplicationProxyRuleRequest) (response *CreateApplicationProxyRuleResponse, err error) {
    if request == nil {
        request = NewCreateApplicationProxyRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreateApplicationProxyRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApplicationProxyRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApplicationProxyRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCLSIndexRequest() (request *CreateCLSIndexRequest) {
    request = &CreateCLSIndexRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateCLSIndex")
    
    
    return
}

func NewCreateCLSIndexResponse() (response *CreateCLSIndexResponse) {
    response = &CreateCLSIndexResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCLSIndex
// This API is used to create key-value indexes for delivered log fields in the corresponding Tencent Cloud CLS log topic for a specified real-time log delivery task (task-id). If indexes have been created in CLS, this API will append indexes through merging.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATECLSLOGSETFAILED = "FailedOperation.CreateClsLogSetFailed"
func (c *Client) CreateCLSIndex(request *CreateCLSIndexRequest) (response *CreateCLSIndexResponse, err error) {
    return c.CreateCLSIndexWithContext(context.Background(), request)
}

// CreateCLSIndex
// This API is used to create key-value indexes for delivered log fields in the corresponding Tencent Cloud CLS log topic for a specified real-time log delivery task (task-id). If indexes have been created in CLS, this API will append indexes through merging.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATECLSLOGSETFAILED = "FailedOperation.CreateClsLogSetFailed"
func (c *Client) CreateCLSIndexWithContext(ctx context.Context, request *CreateCLSIndexRequest) (response *CreateCLSIndexResponse, err error) {
    if request == nil {
        request = NewCreateCLSIndexRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreateCLSIndex")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCLSIndex require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCLSIndexResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConfigGroupVersionRequest() (request *CreateConfigGroupVersionRequest) {
    request = &CreateConfigGroupVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateConfigGroupVersion")
    
    
    return
}

func NewCreateConfigGroupVersionResponse() (response *CreateConfigGroupVersionResponse) {
    response = &CreateConfigGroupVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateConfigGroupVersion
// In version management mode, this API is used to create a new version of a specified configuration group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CONFIGCONDITIONSYNTAXERROR = "FailedOperation.ConfigConditionSyntaxError"
//  FAILEDOPERATION_CONFIGCONDITIONUNKNOWNTARGET = "FailedOperation.ConfigConditionUnknownTarget"
//  FAILEDOPERATION_CONFIGCONDITIONVALUEEMPTYERROR = "FailedOperation.ConfigConditionValueEmptyError"
//  FAILEDOPERATION_CONFIGDUPLICATEKEYERROR = "FailedOperation.ConfigDuplicateKeyError"
//  FAILEDOPERATION_CONFIGFIELDTYPEERROR = "FailedOperation.ConfigFieldTypeError"
//  FAILEDOPERATION_CONFIGFORMATERROR = "FailedOperation.ConfigFormatError"
//  FAILEDOPERATION_CONFIGJSONFORMATERROR = "FailedOperation.ConfigJSONFormatError"
//  FAILEDOPERATION_CONFIGMALFORMEDCONTENT = "FailedOperation.ConfigMalformedContent"
//  FAILEDOPERATION_CONFIGPARAMVALIDATEERRORS = "FailedOperation.ConfigParamValidateErrors"
//  FAILEDOPERATION_CONFIGTIMEPARSINGERROR = "FailedOperation.ConfigTimeParsingError"
//  FAILEDOPERATION_CONFIGUNKNOWNFIELD = "FailedOperation.ConfigUnknownField"
//  FAILEDOPERATION_CONFIGUNSUPPORTEDACTION = "FailedOperation.ConfigUnsupportedAction"
//  FAILEDOPERATION_CONFIGUNSUPPORTEDFORMATVERSION = "FailedOperation.ConfigUnsupportedFormatVersion"
//  FAILEDOPERATION_MISSINGCONFIGCHUNK = "FailedOperation.MissingConfigChunk"
//  FAILEDOPERATION_UNKNOWNCONFIGGROUPTYPE = "FailedOperation.UnknownConfigGroupType"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DUPLICATERULE = "InvalidParameter.DuplicateRule"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"
//  INVALIDPARAMETER_INVALIDCACHEKEYQUERYSTRINGACTION = "InvalidParameter.InvalidCacheKeyQueryStringAction"
//  INVALIDPARAMETER_INVALIDCONDITIONS = "InvalidParameter.InvalidConditions"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINE = "InvalidParameter.InvalidDynamicRoutine"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGEFOLLOWORIGIN = "InvalidParameter.InvalidMaxAgeFollowOrigin"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUG = "InvalidParameter.InvalidStandardDebug"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_STATUSCODECACHEINVALIDSTATUSCODE = "InvalidParameter.StatusCodeCacheInvalidStatusCode"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RULELIMITEXCEEDED = "LimitExceeded.RuleLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDIPV6CONFLICT = "OperationDenied.AccelerateMainlandIpv6Conflict"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_NOTINVERSIONCONTROLWHITELIST = "OperationDenied.NotInVersionControlWhiteList"
//  OPERATIONDENIED_OWNERSHIPVERIFICATIONNOTPASSED = "OperationDenied.OwnershipVerificationNotPassed"
//  OPERATIONDENIED_VERSIONCONTROLLOCKED = "OperationDenied.VersionControlLocked"
//  OPERATIONDENIED_WORKMODENOTINVERSIONCONTROL = "OperationDenied.WorkModeNotInVersionControl"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_FUNCTIONNOTFOUND = "ResourceUnavailable.FunctionNotFound"
func (c *Client) CreateConfigGroupVersion(request *CreateConfigGroupVersionRequest) (response *CreateConfigGroupVersionResponse, err error) {
    return c.CreateConfigGroupVersionWithContext(context.Background(), request)
}

// CreateConfigGroupVersion
// In version management mode, this API is used to create a new version of a specified configuration group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CONFIGCONDITIONSYNTAXERROR = "FailedOperation.ConfigConditionSyntaxError"
//  FAILEDOPERATION_CONFIGCONDITIONUNKNOWNTARGET = "FailedOperation.ConfigConditionUnknownTarget"
//  FAILEDOPERATION_CONFIGCONDITIONVALUEEMPTYERROR = "FailedOperation.ConfigConditionValueEmptyError"
//  FAILEDOPERATION_CONFIGDUPLICATEKEYERROR = "FailedOperation.ConfigDuplicateKeyError"
//  FAILEDOPERATION_CONFIGFIELDTYPEERROR = "FailedOperation.ConfigFieldTypeError"
//  FAILEDOPERATION_CONFIGFORMATERROR = "FailedOperation.ConfigFormatError"
//  FAILEDOPERATION_CONFIGJSONFORMATERROR = "FailedOperation.ConfigJSONFormatError"
//  FAILEDOPERATION_CONFIGMALFORMEDCONTENT = "FailedOperation.ConfigMalformedContent"
//  FAILEDOPERATION_CONFIGPARAMVALIDATEERRORS = "FailedOperation.ConfigParamValidateErrors"
//  FAILEDOPERATION_CONFIGTIMEPARSINGERROR = "FailedOperation.ConfigTimeParsingError"
//  FAILEDOPERATION_CONFIGUNKNOWNFIELD = "FailedOperation.ConfigUnknownField"
//  FAILEDOPERATION_CONFIGUNSUPPORTEDACTION = "FailedOperation.ConfigUnsupportedAction"
//  FAILEDOPERATION_CONFIGUNSUPPORTEDFORMATVERSION = "FailedOperation.ConfigUnsupportedFormatVersion"
//  FAILEDOPERATION_MISSINGCONFIGCHUNK = "FailedOperation.MissingConfigChunk"
//  FAILEDOPERATION_UNKNOWNCONFIGGROUPTYPE = "FailedOperation.UnknownConfigGroupType"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DUPLICATERULE = "InvalidParameter.DuplicateRule"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"
//  INVALIDPARAMETER_INVALIDCACHEKEYQUERYSTRINGACTION = "InvalidParameter.InvalidCacheKeyQueryStringAction"
//  INVALIDPARAMETER_INVALIDCONDITIONS = "InvalidParameter.InvalidConditions"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINE = "InvalidParameter.InvalidDynamicRoutine"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGEFOLLOWORIGIN = "InvalidParameter.InvalidMaxAgeFollowOrigin"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUG = "InvalidParameter.InvalidStandardDebug"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_STATUSCODECACHEINVALIDSTATUSCODE = "InvalidParameter.StatusCodeCacheInvalidStatusCode"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RULELIMITEXCEEDED = "LimitExceeded.RuleLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDIPV6CONFLICT = "OperationDenied.AccelerateMainlandIpv6Conflict"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_NOTINVERSIONCONTROLWHITELIST = "OperationDenied.NotInVersionControlWhiteList"
//  OPERATIONDENIED_OWNERSHIPVERIFICATIONNOTPASSED = "OperationDenied.OwnershipVerificationNotPassed"
//  OPERATIONDENIED_VERSIONCONTROLLOCKED = "OperationDenied.VersionControlLocked"
//  OPERATIONDENIED_WORKMODENOTINVERSIONCONTROL = "OperationDenied.WorkModeNotInVersionControl"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_FUNCTIONNOTFOUND = "ResourceUnavailable.FunctionNotFound"
func (c *Client) CreateConfigGroupVersionWithContext(ctx context.Context, request *CreateConfigGroupVersionRequest) (response *CreateConfigGroupVersionResponse, err error) {
    if request == nil {
        request = NewCreateConfigGroupVersionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreateConfigGroupVersion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConfigGroupVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConfigGroupVersionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateContentIdentifierRequest() (request *CreateContentIdentifierRequest) {
    request = &CreateContentIdentifierRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateContentIdentifier")
    
    
    return
}

func NewCreateContentIdentifierResponse() (response *CreateContentIdentifierResponse) {
    response = &CreateContentIdentifierResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateContentIdentifier
// This API is used to create a content identifier. You can set the description, tag, and other information. At the same time, you need to bind an enterprise package for billing data statistics. A content identifier can only be bound to one billing package, while a billing package can bind multiple content identifiers. This functionality is only available to allowlisted users.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CONFIGCONDITIONSYNTAXERROR = "FailedOperation.ConfigConditionSyntaxError"
//  FAILEDOPERATION_CONFIGCONDITIONUNKNOWNTARGET = "FailedOperation.ConfigConditionUnknownTarget"
//  FAILEDOPERATION_CONFIGCONDITIONVALUEEMPTYERROR = "FailedOperation.ConfigConditionValueEmptyError"
//  FAILEDOPERATION_CONFIGDUPLICATEKEYERROR = "FailedOperation.ConfigDuplicateKeyError"
//  FAILEDOPERATION_CONFIGFIELDTYPEERROR = "FailedOperation.ConfigFieldTypeError"
//  FAILEDOPERATION_CONFIGFORMATERROR = "FailedOperation.ConfigFormatError"
//  FAILEDOPERATION_CONFIGJSONFORMATERROR = "FailedOperation.ConfigJSONFormatError"
//  FAILEDOPERATION_CONFIGMALFORMEDCONTENT = "FailedOperation.ConfigMalformedContent"
//  FAILEDOPERATION_CONFIGPARAMVALIDATEERRORS = "FailedOperation.ConfigParamValidateErrors"
//  FAILEDOPERATION_CONFIGTIMEPARSINGERROR = "FailedOperation.ConfigTimeParsingError"
//  FAILEDOPERATION_CONFIGUNKNOWNFIELD = "FailedOperation.ConfigUnknownField"
//  FAILEDOPERATION_CONFIGUNSUPPORTEDACTION = "FailedOperation.ConfigUnsupportedAction"
//  FAILEDOPERATION_CONFIGUNSUPPORTEDFORMATVERSION = "FailedOperation.ConfigUnsupportedFormatVersion"
//  FAILEDOPERATION_MISSINGCONFIGCHUNK = "FailedOperation.MissingConfigChunk"
//  FAILEDOPERATION_UNKNOWNCONFIGGROUPTYPE = "FailedOperation.UnknownConfigGroupType"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DUPLICATERULE = "InvalidParameter.DuplicateRule"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"
//  INVALIDPARAMETER_INVALIDCACHEKEYQUERYSTRINGACTION = "InvalidParameter.InvalidCacheKeyQueryStringAction"
//  INVALIDPARAMETER_INVALIDCONDITIONS = "InvalidParameter.InvalidConditions"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINE = "InvalidParameter.InvalidDynamicRoutine"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGEFOLLOWORIGIN = "InvalidParameter.InvalidMaxAgeFollowOrigin"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUG = "InvalidParameter.InvalidStandardDebug"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_STATUSCODECACHEINVALIDSTATUSCODE = "InvalidParameter.StatusCodeCacheInvalidStatusCode"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RULELIMITEXCEEDED = "LimitExceeded.RuleLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDIPV6CONFLICT = "OperationDenied.AccelerateMainlandIpv6Conflict"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_NOTINVERSIONCONTROLWHITELIST = "OperationDenied.NotInVersionControlWhiteList"
//  OPERATIONDENIED_OWNERSHIPVERIFICATIONNOTPASSED = "OperationDenied.OwnershipVerificationNotPassed"
//  OPERATIONDENIED_VERSIONCONTROLLOCKED = "OperationDenied.VersionControlLocked"
//  OPERATIONDENIED_WORKMODENOTINVERSIONCONTROL = "OperationDenied.WorkModeNotInVersionControl"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_FUNCTIONNOTFOUND = "ResourceUnavailable.FunctionNotFound"
func (c *Client) CreateContentIdentifier(request *CreateContentIdentifierRequest) (response *CreateContentIdentifierResponse, err error) {
    return c.CreateContentIdentifierWithContext(context.Background(), request)
}

// CreateContentIdentifier
// This API is used to create a content identifier. You can set the description, tag, and other information. At the same time, you need to bind an enterprise package for billing data statistics. A content identifier can only be bound to one billing package, while a billing package can bind multiple content identifiers. This functionality is only available to allowlisted users.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CONFIGCONDITIONSYNTAXERROR = "FailedOperation.ConfigConditionSyntaxError"
//  FAILEDOPERATION_CONFIGCONDITIONUNKNOWNTARGET = "FailedOperation.ConfigConditionUnknownTarget"
//  FAILEDOPERATION_CONFIGCONDITIONVALUEEMPTYERROR = "FailedOperation.ConfigConditionValueEmptyError"
//  FAILEDOPERATION_CONFIGDUPLICATEKEYERROR = "FailedOperation.ConfigDuplicateKeyError"
//  FAILEDOPERATION_CONFIGFIELDTYPEERROR = "FailedOperation.ConfigFieldTypeError"
//  FAILEDOPERATION_CONFIGFORMATERROR = "FailedOperation.ConfigFormatError"
//  FAILEDOPERATION_CONFIGJSONFORMATERROR = "FailedOperation.ConfigJSONFormatError"
//  FAILEDOPERATION_CONFIGMALFORMEDCONTENT = "FailedOperation.ConfigMalformedContent"
//  FAILEDOPERATION_CONFIGPARAMVALIDATEERRORS = "FailedOperation.ConfigParamValidateErrors"
//  FAILEDOPERATION_CONFIGTIMEPARSINGERROR = "FailedOperation.ConfigTimeParsingError"
//  FAILEDOPERATION_CONFIGUNKNOWNFIELD = "FailedOperation.ConfigUnknownField"
//  FAILEDOPERATION_CONFIGUNSUPPORTEDACTION = "FailedOperation.ConfigUnsupportedAction"
//  FAILEDOPERATION_CONFIGUNSUPPORTEDFORMATVERSION = "FailedOperation.ConfigUnsupportedFormatVersion"
//  FAILEDOPERATION_MISSINGCONFIGCHUNK = "FailedOperation.MissingConfigChunk"
//  FAILEDOPERATION_UNKNOWNCONFIGGROUPTYPE = "FailedOperation.UnknownConfigGroupType"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DUPLICATERULE = "InvalidParameter.DuplicateRule"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"
//  INVALIDPARAMETER_INVALIDCACHEKEYQUERYSTRINGACTION = "InvalidParameter.InvalidCacheKeyQueryStringAction"
//  INVALIDPARAMETER_INVALIDCONDITIONS = "InvalidParameter.InvalidConditions"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINE = "InvalidParameter.InvalidDynamicRoutine"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGEFOLLOWORIGIN = "InvalidParameter.InvalidMaxAgeFollowOrigin"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUG = "InvalidParameter.InvalidStandardDebug"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_STATUSCODECACHEINVALIDSTATUSCODE = "InvalidParameter.StatusCodeCacheInvalidStatusCode"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RULELIMITEXCEEDED = "LimitExceeded.RuleLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDIPV6CONFLICT = "OperationDenied.AccelerateMainlandIpv6Conflict"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_NOTINVERSIONCONTROLWHITELIST = "OperationDenied.NotInVersionControlWhiteList"
//  OPERATIONDENIED_OWNERSHIPVERIFICATIONNOTPASSED = "OperationDenied.OwnershipVerificationNotPassed"
//  OPERATIONDENIED_VERSIONCONTROLLOCKED = "OperationDenied.VersionControlLocked"
//  OPERATIONDENIED_WORKMODENOTINVERSIONCONTROL = "OperationDenied.WorkModeNotInVersionControl"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_FUNCTIONNOTFOUND = "ResourceUnavailable.FunctionNotFound"
func (c *Client) CreateContentIdentifierWithContext(ctx context.Context, request *CreateContentIdentifierRequest) (response *CreateContentIdentifierResponse, err error) {
    if request == nil {
        request = NewCreateContentIdentifierRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreateContentIdentifier")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateContentIdentifier require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateContentIdentifierResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCustomGroupRequest() (request *CreateCustomGroupRequest) {
    request = &CreateCustomGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateCustomGroup")
    
    
    return
}

func NewCreateCustomGroupResponse() (response *CreateCustomGroupResponse) {
    response = &CreateCustomGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCustomGroup
// This API is used to create a custom group. A custom group is used to centrally manage a set of reusable matching entries. For example, you can save multiple JA3/JA4 TLS fingerprints as a TLS fingerprint group and directly reference the group in security rules, avoiding the need to maintain the same fingerprint list in multiple rules.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CONFIGCONDITIONSYNTAXERROR = "FailedOperation.ConfigConditionSyntaxError"
//  FAILEDOPERATION_CONFIGCONDITIONUNKNOWNTARGET = "FailedOperation.ConfigConditionUnknownTarget"
//  FAILEDOPERATION_CONFIGCONDITIONVALUEEMPTYERROR = "FailedOperation.ConfigConditionValueEmptyError"
//  FAILEDOPERATION_CONFIGDUPLICATEKEYERROR = "FailedOperation.ConfigDuplicateKeyError"
//  FAILEDOPERATION_CONFIGFIELDTYPEERROR = "FailedOperation.ConfigFieldTypeError"
//  FAILEDOPERATION_CONFIGFORMATERROR = "FailedOperation.ConfigFormatError"
//  FAILEDOPERATION_CONFIGJSONFORMATERROR = "FailedOperation.ConfigJSONFormatError"
//  FAILEDOPERATION_CONFIGMALFORMEDCONTENT = "FailedOperation.ConfigMalformedContent"
//  FAILEDOPERATION_CONFIGPARAMVALIDATEERRORS = "FailedOperation.ConfigParamValidateErrors"
//  FAILEDOPERATION_CONFIGTIMEPARSINGERROR = "FailedOperation.ConfigTimeParsingError"
//  FAILEDOPERATION_CONFIGUNKNOWNFIELD = "FailedOperation.ConfigUnknownField"
//  FAILEDOPERATION_CONFIGUNSUPPORTEDACTION = "FailedOperation.ConfigUnsupportedAction"
//  FAILEDOPERATION_CONFIGUNSUPPORTEDFORMATVERSION = "FailedOperation.ConfigUnsupportedFormatVersion"
//  FAILEDOPERATION_MISSINGCONFIGCHUNK = "FailedOperation.MissingConfigChunk"
//  FAILEDOPERATION_UNKNOWNCONFIGGROUPTYPE = "FailedOperation.UnknownConfigGroupType"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DUPLICATERULE = "InvalidParameter.DuplicateRule"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"
//  INVALIDPARAMETER_INVALIDCACHEKEYQUERYSTRINGACTION = "InvalidParameter.InvalidCacheKeyQueryStringAction"
//  INVALIDPARAMETER_INVALIDCONDITIONS = "InvalidParameter.InvalidConditions"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINE = "InvalidParameter.InvalidDynamicRoutine"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGEFOLLOWORIGIN = "InvalidParameter.InvalidMaxAgeFollowOrigin"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUG = "InvalidParameter.InvalidStandardDebug"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_STATUSCODECACHEINVALIDSTATUSCODE = "InvalidParameter.StatusCodeCacheInvalidStatusCode"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RULELIMITEXCEEDED = "LimitExceeded.RuleLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDIPV6CONFLICT = "OperationDenied.AccelerateMainlandIpv6Conflict"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_NOTINVERSIONCONTROLWHITELIST = "OperationDenied.NotInVersionControlWhiteList"
//  OPERATIONDENIED_OWNERSHIPVERIFICATIONNOTPASSED = "OperationDenied.OwnershipVerificationNotPassed"
//  OPERATIONDENIED_VERSIONCONTROLLOCKED = "OperationDenied.VersionControlLocked"
//  OPERATIONDENIED_WORKMODENOTINVERSIONCONTROL = "OperationDenied.WorkModeNotInVersionControl"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_FUNCTIONNOTFOUND = "ResourceUnavailable.FunctionNotFound"
func (c *Client) CreateCustomGroup(request *CreateCustomGroupRequest) (response *CreateCustomGroupResponse, err error) {
    return c.CreateCustomGroupWithContext(context.Background(), request)
}

// CreateCustomGroup
// This API is used to create a custom group. A custom group is used to centrally manage a set of reusable matching entries. For example, you can save multiple JA3/JA4 TLS fingerprints as a TLS fingerprint group and directly reference the group in security rules, avoiding the need to maintain the same fingerprint list in multiple rules.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CONFIGCONDITIONSYNTAXERROR = "FailedOperation.ConfigConditionSyntaxError"
//  FAILEDOPERATION_CONFIGCONDITIONUNKNOWNTARGET = "FailedOperation.ConfigConditionUnknownTarget"
//  FAILEDOPERATION_CONFIGCONDITIONVALUEEMPTYERROR = "FailedOperation.ConfigConditionValueEmptyError"
//  FAILEDOPERATION_CONFIGDUPLICATEKEYERROR = "FailedOperation.ConfigDuplicateKeyError"
//  FAILEDOPERATION_CONFIGFIELDTYPEERROR = "FailedOperation.ConfigFieldTypeError"
//  FAILEDOPERATION_CONFIGFORMATERROR = "FailedOperation.ConfigFormatError"
//  FAILEDOPERATION_CONFIGJSONFORMATERROR = "FailedOperation.ConfigJSONFormatError"
//  FAILEDOPERATION_CONFIGMALFORMEDCONTENT = "FailedOperation.ConfigMalformedContent"
//  FAILEDOPERATION_CONFIGPARAMVALIDATEERRORS = "FailedOperation.ConfigParamValidateErrors"
//  FAILEDOPERATION_CONFIGTIMEPARSINGERROR = "FailedOperation.ConfigTimeParsingError"
//  FAILEDOPERATION_CONFIGUNKNOWNFIELD = "FailedOperation.ConfigUnknownField"
//  FAILEDOPERATION_CONFIGUNSUPPORTEDACTION = "FailedOperation.ConfigUnsupportedAction"
//  FAILEDOPERATION_CONFIGUNSUPPORTEDFORMATVERSION = "FailedOperation.ConfigUnsupportedFormatVersion"
//  FAILEDOPERATION_MISSINGCONFIGCHUNK = "FailedOperation.MissingConfigChunk"
//  FAILEDOPERATION_UNKNOWNCONFIGGROUPTYPE = "FailedOperation.UnknownConfigGroupType"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DUPLICATERULE = "InvalidParameter.DuplicateRule"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"
//  INVALIDPARAMETER_INVALIDCACHEKEYQUERYSTRINGACTION = "InvalidParameter.InvalidCacheKeyQueryStringAction"
//  INVALIDPARAMETER_INVALIDCONDITIONS = "InvalidParameter.InvalidConditions"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINE = "InvalidParameter.InvalidDynamicRoutine"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGEFOLLOWORIGIN = "InvalidParameter.InvalidMaxAgeFollowOrigin"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUG = "InvalidParameter.InvalidStandardDebug"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_STATUSCODECACHEINVALIDSTATUSCODE = "InvalidParameter.StatusCodeCacheInvalidStatusCode"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RULELIMITEXCEEDED = "LimitExceeded.RuleLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDIPV6CONFLICT = "OperationDenied.AccelerateMainlandIpv6Conflict"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_NOTINVERSIONCONTROLWHITELIST = "OperationDenied.NotInVersionControlWhiteList"
//  OPERATIONDENIED_OWNERSHIPVERIFICATIONNOTPASSED = "OperationDenied.OwnershipVerificationNotPassed"
//  OPERATIONDENIED_VERSIONCONTROLLOCKED = "OperationDenied.VersionControlLocked"
//  OPERATIONDENIED_WORKMODENOTINVERSIONCONTROL = "OperationDenied.WorkModeNotInVersionControl"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_FUNCTIONNOTFOUND = "ResourceUnavailable.FunctionNotFound"
func (c *Client) CreateCustomGroupWithContext(ctx context.Context, request *CreateCustomGroupRequest) (response *CreateCustomGroupResponse, err error) {
    if request == nil {
        request = NewCreateCustomGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreateCustomGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCustomGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCustomGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCustomizeErrorPageRequest() (request *CreateCustomizeErrorPageRequest) {
    request = &CreateCustomizeErrorPageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateCustomizeErrorPage")
    
    
    return
}

func NewCreateCustomizeErrorPageResponse() (response *CreateCustomizeErrorPageResponse) {
    response = &CreateCustomizeErrorPageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCustomizeErrorPage
// This API is used to create a custom response page.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_CONTENTTYPENOTMATCH = "InvalidParameterValue.ContentTypeNotMatch"
//  INVALIDPARAMETERVALUE_PAGENAMEALREADYEXIST = "InvalidParameterValue.PageNameAlreadyExist"
func (c *Client) CreateCustomizeErrorPage(request *CreateCustomizeErrorPageRequest) (response *CreateCustomizeErrorPageResponse, err error) {
    return c.CreateCustomizeErrorPageWithContext(context.Background(), request)
}

// CreateCustomizeErrorPage
// This API is used to create a custom response page.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_CONTENTTYPENOTMATCH = "InvalidParameterValue.ContentTypeNotMatch"
//  INVALIDPARAMETERVALUE_PAGENAMEALREADYEXIST = "InvalidParameterValue.PageNameAlreadyExist"
func (c *Client) CreateCustomizeErrorPageWithContext(ctx context.Context, request *CreateCustomizeErrorPageRequest) (response *CreateCustomizeErrorPageResponse, err error) {
    if request == nil {
        request = NewCreateCustomizeErrorPageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreateCustomizeErrorPage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCustomizeErrorPage require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCustomizeErrorPageResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDnsRecordRequest() (request *CreateDnsRecordRequest) {
    request = &CreateDnsRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateDnsRecord")
    
    
    return
}

func NewCreateDnsRecordResponse() (response *CreateDnsRecordResponse) {
    response = &CreateDnsRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDnsRecord
// After a site is created and accessed in NS mode, you can create DNS records through this API.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_CONTENTTYPENOTMATCH = "InvalidParameterValue.ContentTypeNotMatch"
//  INVALIDPARAMETERVALUE_PAGENAMEALREADYEXIST = "InvalidParameterValue.PageNameAlreadyExist"
func (c *Client) CreateDnsRecord(request *CreateDnsRecordRequest) (response *CreateDnsRecordResponse, err error) {
    return c.CreateDnsRecordWithContext(context.Background(), request)
}

// CreateDnsRecord
// After a site is created and accessed in NS mode, you can create DNS records through this API.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_CONTENTTYPENOTMATCH = "InvalidParameterValue.ContentTypeNotMatch"
//  INVALIDPARAMETERVALUE_PAGENAMEALREADYEXIST = "InvalidParameterValue.PageNameAlreadyExist"
func (c *Client) CreateDnsRecordWithContext(ctx context.Context, request *CreateDnsRecordRequest) (response *CreateDnsRecordResponse, err error) {
    if request == nil {
        request = NewCreateDnsRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreateDnsRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDnsRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDnsRecordResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEdgeKVNamespaceRequest() (request *CreateEdgeKVNamespaceRequest) {
    request = &CreateEdgeKVNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateEdgeKVNamespace")
    
    
    return
}

func NewCreateEdgeKVNamespaceResponse() (response *CreateEdgeKVNamespaceResponse) {
    response = &CreateEdgeKVNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateEdgeKVNamespace
// This API is used to create a KV namespace at a specified site.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_BADNAMESPACENAME = "InvalidParameter.BadNamespaceName"
//  INVALIDPARAMETER_DUPLICATEBINDINGNAME = "InvalidParameter.DuplicateBindingName"
//  INVALIDPARAMETER_REMARKTOOLONG = "InvalidParameter.RemarkTooLong"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateEdgeKVNamespace(request *CreateEdgeKVNamespaceRequest) (response *CreateEdgeKVNamespaceResponse, err error) {
    return c.CreateEdgeKVNamespaceWithContext(context.Background(), request)
}

// CreateEdgeKVNamespace
// This API is used to create a KV namespace at a specified site.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_BADNAMESPACENAME = "InvalidParameter.BadNamespaceName"
//  INVALIDPARAMETER_DUPLICATEBINDINGNAME = "InvalidParameter.DuplicateBindingName"
//  INVALIDPARAMETER_REMARKTOOLONG = "InvalidParameter.RemarkTooLong"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateEdgeKVNamespaceWithContext(ctx context.Context, request *CreateEdgeKVNamespaceRequest) (response *CreateEdgeKVNamespaceResponse, err error) {
    if request == nil {
        request = NewCreateEdgeKVNamespaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreateEdgeKVNamespace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEdgeKVNamespace require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEdgeKVNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFunctionRequest() (request *CreateFunctionRequest) {
    request = &CreateFunctionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateFunction")
    
    
    return
}

func NewCreateFunctionResponse() (response *CreateFunctionResponse) {
    response = &CreateFunctionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateFunction
// This API is used to create and deploy an edge function to EdgeOne edge nodes.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FUNCTIONDEPLOYING = "FailedOperation.FunctionDeploying"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_BADCONTENT = "InvalidParameter.BadContent"
//  INVALIDPARAMETER_BADFUNCTIONNAME = "InvalidParameter.BadFunctionName"
//  INVALIDPARAMETER_CONTENTEXCEEDSLIMIT = "InvalidParameter.ContentExceedsLimit"
//  INVALIDPARAMETER_FUNCTIONNAMECONFLICT = "InvalidParameter.FunctionNameConflict"
//  INVALIDPARAMETER_LENGTHEXCEEDSLIMIT = "InvalidParameter.LengthExceedsLimit"
//  LIMITEXCEEDED_FUNCTIONLIMITEXCEEDED = "LimitExceeded.FunctionLimitExceeded"
//  OPERATIONDENIED_VERSIONCONTROLLOCKED = "OperationDenied.VersionControlLocked"
//  RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateFunction(request *CreateFunctionRequest) (response *CreateFunctionResponse, err error) {
    return c.CreateFunctionWithContext(context.Background(), request)
}

// CreateFunction
// This API is used to create and deploy an edge function to EdgeOne edge nodes.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FUNCTIONDEPLOYING = "FailedOperation.FunctionDeploying"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_BADCONTENT = "InvalidParameter.BadContent"
//  INVALIDPARAMETER_BADFUNCTIONNAME = "InvalidParameter.BadFunctionName"
//  INVALIDPARAMETER_CONTENTEXCEEDSLIMIT = "InvalidParameter.ContentExceedsLimit"
//  INVALIDPARAMETER_FUNCTIONNAMECONFLICT = "InvalidParameter.FunctionNameConflict"
//  INVALIDPARAMETER_LENGTHEXCEEDSLIMIT = "InvalidParameter.LengthExceedsLimit"
//  LIMITEXCEEDED_FUNCTIONLIMITEXCEEDED = "LimitExceeded.FunctionLimitExceeded"
//  OPERATIONDENIED_VERSIONCONTROLLOCKED = "OperationDenied.VersionControlLocked"
//  RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateFunctionWithContext(ctx context.Context, request *CreateFunctionRequest) (response *CreateFunctionResponse, err error) {
    if request == nil {
        request = NewCreateFunctionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreateFunction")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFunction require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFunctionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFunctionReplicaRequest() (request *CreateFunctionReplicaRequest) {
    request = &CreateFunctionReplicaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateFunctionReplica")
    
    
    return
}

func NewCreateFunctionReplicaResponse() (response *CreateFunctionReplicaResponse) {
    response = &CreateFunctionReplicaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateFunctionReplica
// This API is used to create copies of a specified edge function. After creating copies, when client requests match the configured trigger rules or default domain name, you can access a specific function copy by adding EO-Function-Replica-Name:[copy name] to the request header. Each function supports creating two copies by default.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FUNCTIONDEPLOYING = "FailedOperation.FunctionDeploying"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_BADCONTENT = "InvalidParameter.BadContent"
//  INVALIDPARAMETER_BADFUNCTIONNAME = "InvalidParameter.BadFunctionName"
//  INVALIDPARAMETER_CONTENTEXCEEDSLIMIT = "InvalidParameter.ContentExceedsLimit"
//  INVALIDPARAMETER_FUNCTIONNAMECONFLICT = "InvalidParameter.FunctionNameConflict"
//  INVALIDPARAMETER_LENGTHEXCEEDSLIMIT = "InvalidParameter.LengthExceedsLimit"
//  LIMITEXCEEDED_FUNCTIONLIMITEXCEEDED = "LimitExceeded.FunctionLimitExceeded"
//  OPERATIONDENIED_VERSIONCONTROLLOCKED = "OperationDenied.VersionControlLocked"
//  RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateFunctionReplica(request *CreateFunctionReplicaRequest) (response *CreateFunctionReplicaResponse, err error) {
    return c.CreateFunctionReplicaWithContext(context.Background(), request)
}

// CreateFunctionReplica
// This API is used to create copies of a specified edge function. After creating copies, when client requests match the configured trigger rules or default domain name, you can access a specific function copy by adding EO-Function-Replica-Name:[copy name] to the request header. Each function supports creating two copies by default.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FUNCTIONDEPLOYING = "FailedOperation.FunctionDeploying"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_BADCONTENT = "InvalidParameter.BadContent"
//  INVALIDPARAMETER_BADFUNCTIONNAME = "InvalidParameter.BadFunctionName"
//  INVALIDPARAMETER_CONTENTEXCEEDSLIMIT = "InvalidParameter.ContentExceedsLimit"
//  INVALIDPARAMETER_FUNCTIONNAMECONFLICT = "InvalidParameter.FunctionNameConflict"
//  INVALIDPARAMETER_LENGTHEXCEEDSLIMIT = "InvalidParameter.LengthExceedsLimit"
//  LIMITEXCEEDED_FUNCTIONLIMITEXCEEDED = "LimitExceeded.FunctionLimitExceeded"
//  OPERATIONDENIED_VERSIONCONTROLLOCKED = "OperationDenied.VersionControlLocked"
//  RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateFunctionReplicaWithContext(ctx context.Context, request *CreateFunctionReplicaRequest) (response *CreateFunctionReplicaResponse, err error) {
    if request == nil {
        request = NewCreateFunctionReplicaRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreateFunctionReplica")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFunctionReplica require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFunctionReplicaResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFunctionRuleRequest() (request *CreateFunctionRuleRequest) {
    request = &CreateFunctionRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateFunctionRule")
    
    
    return
}

func NewCreateFunctionRuleResponse() (response *CreateFunctionRuleResponse) {
    response = &CreateFunctionRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateFunctionRule
// Creates a trigger rule for an edge function. It supports customized filter conditions to determine whether to execute a function. When a function needs to be executed, multiple ways are provided to select the target function, including directly specifying, selecting based on the client's associated region, and weight-based selection.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RULEOPERATIONCONFLICT = "FailedOperation.RuleOperationConflict"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_DUPLICATERULE = "InvalidParameter.DuplicateRule"
//  INVALIDPARAMETER_INVALIDCONDITIONS = "InvalidParameter.InvalidConditions"
//  LIMITEXCEEDED_RULELIMITEXCEEDED = "LimitExceeded.RuleLimitExceeded"
//  OPERATIONDENIED_VERSIONCONTROLLOCKED = "OperationDenied.VersionControlLocked"
//  RESOURCEUNAVAILABLE_FUNCTIONNOTFOUND = "ResourceUnavailable.FunctionNotFound"
//  RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"
func (c *Client) CreateFunctionRule(request *CreateFunctionRuleRequest) (response *CreateFunctionRuleResponse, err error) {
    return c.CreateFunctionRuleWithContext(context.Background(), request)
}

// CreateFunctionRule
// Creates a trigger rule for an edge function. It supports customized filter conditions to determine whether to execute a function. When a function needs to be executed, multiple ways are provided to select the target function, including directly specifying, selecting based on the client's associated region, and weight-based selection.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RULEOPERATIONCONFLICT = "FailedOperation.RuleOperationConflict"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_DUPLICATERULE = "InvalidParameter.DuplicateRule"
//  INVALIDPARAMETER_INVALIDCONDITIONS = "InvalidParameter.InvalidConditions"
//  LIMITEXCEEDED_RULELIMITEXCEEDED = "LimitExceeded.RuleLimitExceeded"
//  OPERATIONDENIED_VERSIONCONTROLLOCKED = "OperationDenied.VersionControlLocked"
//  RESOURCEUNAVAILABLE_FUNCTIONNOTFOUND = "ResourceUnavailable.FunctionNotFound"
//  RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"
func (c *Client) CreateFunctionRuleWithContext(ctx context.Context, request *CreateFunctionRuleRequest) (response *CreateFunctionRuleResponse, err error) {
    if request == nil {
        request = NewCreateFunctionRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreateFunctionRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFunctionRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFunctionRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInferenceAPITokenRequest() (request *CreateInferenceAPITokenRequest) {
    request = &CreateInferenceAPITokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateInferenceAPIToken")
    
    
    return
}

func NewCreateInferenceAPITokenResponse() (response *CreateInferenceAPITokenResponse) {
    response = &CreateInferenceAPITokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateInferenceAPIToken
// Create an inference API Token for authentication when accessing the inference service. The Token content is returned only once when creating, and each site can create up to 100 tokens.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ZONENOTFOUND = "InvalidParameterValue.ZoneNotFound"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateInferenceAPIToken(request *CreateInferenceAPITokenRequest) (response *CreateInferenceAPITokenResponse, err error) {
    return c.CreateInferenceAPITokenWithContext(context.Background(), request)
}

// CreateInferenceAPIToken
// Create an inference API Token for authentication when accessing the inference service. The Token content is returned only once when creating, and each site can create up to 100 tokens.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ZONENOTFOUND = "InvalidParameterValue.ZoneNotFound"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateInferenceAPITokenWithContext(ctx context.Context, request *CreateInferenceAPITokenRequest) (response *CreateInferenceAPITokenResponse, err error) {
    if request == nil {
        request = NewCreateInferenceAPITokenRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreateInferenceAPIToken")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInferenceAPIToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInferenceAPITokenResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInferenceDomainRequest() (request *CreateInferenceDomainRequest) {
    request = &CreateInferenceDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateInferenceDomain")
    
    
    return
}

func NewCreateInferenceDomainResponse() (response *CreateInferenceDomainResponse) {
    response = &CreateInferenceDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateInferenceDomain
// Create an inference service domain name for accessing the inference service through a custom domain name. An inference service supports creating a maximum of 5 custom domain names.
//
// After successful creation, complete the following steps to make the domain name function normally for external access:
//
// 1. To validate the CNAME configuration status of a domain, see the CheckCnameStatus api.
//
// 2. To verify ownership, refer to the VerifyOwnership API (https://www.tencentcloud.com/document/api/1552/98879?from_cn_redirect=1);
//
// 3. To configure domain name certificates, refer to the ModifyHostsCertificate api.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_ZONENOTFOUND = "InvalidParameterValue.ZoneNotFound"
//  RESOURCENOTFOUND_INFERENCEDOMAIN = "ResourceNotFound.InferenceDomain"
//  RESOURCENOTFOUND_INFERENCESERVICE = "ResourceNotFound.InferenceService"
func (c *Client) CreateInferenceDomain(request *CreateInferenceDomainRequest) (response *CreateInferenceDomainResponse, err error) {
    return c.CreateInferenceDomainWithContext(context.Background(), request)
}

// CreateInferenceDomain
// Create an inference service domain name for accessing the inference service through a custom domain name. An inference service supports creating a maximum of 5 custom domain names.
//
// After successful creation, complete the following steps to make the domain name function normally for external access:
//
// 1. To validate the CNAME configuration status of a domain, see the CheckCnameStatus api.
//
// 2. To verify ownership, refer to the VerifyOwnership API (https://www.tencentcloud.com/document/api/1552/98879?from_cn_redirect=1);
//
// 3. To configure domain name certificates, refer to the ModifyHostsCertificate api.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_ZONENOTFOUND = "InvalidParameterValue.ZoneNotFound"
//  RESOURCENOTFOUND_INFERENCEDOMAIN = "ResourceNotFound.InferenceDomain"
//  RESOURCENOTFOUND_INFERENCESERVICE = "ResourceNotFound.InferenceService"
func (c *Client) CreateInferenceDomainWithContext(ctx context.Context, request *CreateInferenceDomainRequest) (response *CreateInferenceDomainResponse, err error) {
    if request == nil {
        request = NewCreateInferenceDomainRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreateInferenceDomain")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInferenceDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInferenceDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInferenceServiceRequest() (request *CreateInferenceServiceRequest) {
    request = &CreateInferenceServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateInferenceService")
    
    
    return
}

func NewCreateInferenceServiceResponse() (response *CreateInferenceServiceResponse) {
    response = &CreateInferenceServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateInferenceService
// This API is used to create an inference service. You can set the service name, listening port, container image configuration, and resource configuration. An access address is provided after successful creation.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INFERENCESERVICENAME = "InvalidParameterValue.InferenceServiceName"
//  INVALIDPARAMETERVALUE_INFERENCESERVICENAMEDUPLICATE = "InvalidParameterValue.InferenceServiceNameDuplicate"
//  INVALIDPARAMETERVALUE_ZONENOTFOUND = "InvalidParameterValue.ZoneNotFound"
//  LIMITEXCEEDED_INFERENCESERVICE = "LimitExceeded.InferenceService"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateInferenceService(request *CreateInferenceServiceRequest) (response *CreateInferenceServiceResponse, err error) {
    return c.CreateInferenceServiceWithContext(context.Background(), request)
}

// CreateInferenceService
// This API is used to create an inference service. You can set the service name, listening port, container image configuration, and resource configuration. An access address is provided after successful creation.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INFERENCESERVICENAME = "InvalidParameterValue.InferenceServiceName"
//  INVALIDPARAMETERVALUE_INFERENCESERVICENAMEDUPLICATE = "InvalidParameterValue.InferenceServiceNameDuplicate"
//  INVALIDPARAMETERVALUE_ZONENOTFOUND = "InvalidParameterValue.ZoneNotFound"
//  LIMITEXCEEDED_INFERENCESERVICE = "LimitExceeded.InferenceService"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateInferenceServiceWithContext(ctx context.Context, request *CreateInferenceServiceRequest) (response *CreateInferenceServiceResponse, err error) {
    if request == nil {
        request = NewCreateInferenceServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreateInferenceService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInferenceService require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInferenceServiceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateJustInTimeTranscodeTemplateRequest() (request *CreateJustInTimeTranscodeTemplateRequest) {
    request = &CreateJustInTimeTranscodeTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateJustInTimeTranscodeTemplate")
    
    
    return
}

func NewCreateJustInTimeTranscodeTemplateResponse() (response *CreateJustInTimeTranscodeTemplateResponse) {
    response = &CreateJustInTimeTranscodeTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateJustInTimeTranscodeTemplate
// JIT transcoding already provides preset transcoding templates to meet most needs. For personalized transcoding requirements, you can create custom transcoding templates through this API, with up to 100 custom templates.
//
// To ensure consistency of JIT transcoding effects and avoid video output exceptions caused by template changes during EO cache or M3U8 sharding, the template cannot be modified after creation.
//
// To learn about JIT transcoding capacity, see EdgeOne Video Instant Processing Function Introduction (https://www.tencentcloud.com/document/product/1552/111927?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION_PRODUCTDISCONTINUED = "FailedOperation.ProductDiscontinued"
//  FAILEDOPERATION_PRODUCTNOTACTIVATED = "FailedOperation.ProductNotActivated"
//  FAILEDOPERATION_TEMPLATEOVERLIMIT = "FailedOperation.TemplateOverLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ZONENOTFOUND = "InvalidParameterValue.ZoneNotFound"
//  INVALIDPARAMETERVALUE_ZONEPAUSED = "InvalidParameterValue.ZonePaused"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateJustInTimeTranscodeTemplate(request *CreateJustInTimeTranscodeTemplateRequest) (response *CreateJustInTimeTranscodeTemplateResponse, err error) {
    return c.CreateJustInTimeTranscodeTemplateWithContext(context.Background(), request)
}

// CreateJustInTimeTranscodeTemplate
// JIT transcoding already provides preset transcoding templates to meet most needs. For personalized transcoding requirements, you can create custom transcoding templates through this API, with up to 100 custom templates.
//
// To ensure consistency of JIT transcoding effects and avoid video output exceptions caused by template changes during EO cache or M3U8 sharding, the template cannot be modified after creation.
//
// To learn about JIT transcoding capacity, see EdgeOne Video Instant Processing Function Introduction (https://www.tencentcloud.com/document/product/1552/111927?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION_PRODUCTDISCONTINUED = "FailedOperation.ProductDiscontinued"
//  FAILEDOPERATION_PRODUCTNOTACTIVATED = "FailedOperation.ProductNotActivated"
//  FAILEDOPERATION_TEMPLATEOVERLIMIT = "FailedOperation.TemplateOverLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ZONENOTFOUND = "InvalidParameterValue.ZoneNotFound"
//  INVALIDPARAMETERVALUE_ZONEPAUSED = "InvalidParameterValue.ZonePaused"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateJustInTimeTranscodeTemplateWithContext(ctx context.Context, request *CreateJustInTimeTranscodeTemplateRequest) (response *CreateJustInTimeTranscodeTemplateResponse, err error) {
    if request == nil {
        request = NewCreateJustInTimeTranscodeTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreateJustInTimeTranscodeTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateJustInTimeTranscodeTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateJustInTimeTranscodeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateL4ProxyRequest() (request *CreateL4ProxyRequest) {
    request = &CreateL4ProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateL4Proxy")
    
    
    return
}

func NewCreateL4ProxyResponse() (response *CreateL4ProxyResponse) {
    response = &CreateL4ProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateL4Proxy
// This API is used to create a layer 4 proxy instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_PROXYNAMEDUPLICATING = "InvalidParameter.ProxyNameDuplicating"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDDISABLE = "OperationDenied.AccelerateMainlandDisable"
//  OPERATIONDENIED_ACCELERATEMAINLANDIPV6CONFLICT = "OperationDenied.AccelerateMainlandIpv6Conflict"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_IPV6ADVANCEDCONFLICT = "OperationDenied.Ipv6AdvancedConflict"
//  OPERATIONDENIED_IPV6STATICIPCONFLICT = "OperationDenied.Ipv6StaticIpConflict"
//  OPERATIONDENIED_L4LACKOFRESOURCES = "OperationDenied.L4LackOfResources"
//  OPERATIONDENIED_MSGIPV6ADVANCEDCONFLICT = "OperationDenied.MsgIpv6AdvancedConflict"
//  OPERATIONDENIED_STATICIPAREACONFLICT = "OperationDenied.StaticIpAreaConflict"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateL4Proxy(request *CreateL4ProxyRequest) (response *CreateL4ProxyResponse, err error) {
    return c.CreateL4ProxyWithContext(context.Background(), request)
}

// CreateL4Proxy
// This API is used to create a layer 4 proxy instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_PROXYNAMEDUPLICATING = "InvalidParameter.ProxyNameDuplicating"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDDISABLE = "OperationDenied.AccelerateMainlandDisable"
//  OPERATIONDENIED_ACCELERATEMAINLANDIPV6CONFLICT = "OperationDenied.AccelerateMainlandIpv6Conflict"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_IPV6ADVANCEDCONFLICT = "OperationDenied.Ipv6AdvancedConflict"
//  OPERATIONDENIED_IPV6STATICIPCONFLICT = "OperationDenied.Ipv6StaticIpConflict"
//  OPERATIONDENIED_L4LACKOFRESOURCES = "OperationDenied.L4LackOfResources"
//  OPERATIONDENIED_MSGIPV6ADVANCEDCONFLICT = "OperationDenied.MsgIpv6AdvancedConflict"
//  OPERATIONDENIED_STATICIPAREACONFLICT = "OperationDenied.StaticIpAreaConflict"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateL4ProxyWithContext(ctx context.Context, request *CreateL4ProxyRequest) (response *CreateL4ProxyResponse, err error) {
    if request == nil {
        request = NewCreateL4ProxyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreateL4Proxy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateL4Proxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateL4ProxyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateL4ProxyRulesRequest() (request *CreateL4ProxyRulesRequest) {
    request = &CreateL4ProxyRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateL4ProxyRules")
    
    
    return
}

func NewCreateL4ProxyRulesResponse() (response *CreateL4ProxyRulesResponse) {
    response = &CreateL4ProxyRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateL4ProxyRules
// This API is used to create Layer 4 proxy instance rules and supports individual and batch creation.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_INVALIDORIGINVALUE = "InvalidParameter.InvalidOriginValue"
//  INVALIDPARAMETER_RULEPORTDUPLICATING = "InvalidParameter.RulePortDuplicating"
//  LIMITEXCEEDED_PROXYRULESLIMITEXCEEDED = "LimitExceeded.ProxyRulesLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_L4PORTLACKOFRESOURCES = "OperationDenied.L4PortLackOfResources"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateL4ProxyRules(request *CreateL4ProxyRulesRequest) (response *CreateL4ProxyRulesResponse, err error) {
    return c.CreateL4ProxyRulesWithContext(context.Background(), request)
}

// CreateL4ProxyRules
// This API is used to create Layer 4 proxy instance rules and supports individual and batch creation.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_INVALIDORIGINVALUE = "InvalidParameter.InvalidOriginValue"
//  INVALIDPARAMETER_RULEPORTDUPLICATING = "InvalidParameter.RulePortDuplicating"
//  LIMITEXCEEDED_PROXYRULESLIMITEXCEEDED = "LimitExceeded.ProxyRulesLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_L4PORTLACKOFRESOURCES = "OperationDenied.L4PortLackOfResources"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateL4ProxyRulesWithContext(ctx context.Context, request *CreateL4ProxyRulesRequest) (response *CreateL4ProxyRulesResponse, err error) {
    if request == nil {
        request = NewCreateL4ProxyRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreateL4ProxyRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateL4ProxyRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateL4ProxyRulesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateL7AccRulesRequest() (request *CreateL7AccRulesRequest) {
    request = &CreateL7AccRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateL7AccRules")
    
    
    return
}

func NewCreateL7AccRulesResponse() (response *CreateL7AccRulesResponse) {
    response = &CreateL7AccRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateL7AccRules
// This API is used to create rules in the [rule engine](https://www.tencentcloud.com/document/product/1552/70901?from_cn_redirect=1), with support for batch creation.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACCESSREDIRECTREGEXERROR = "InvalidParameter.AccessRedirectRegexError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_ADVANCEDORIGINROUTINGNOTSUPPORTACCELERATEMAINLAND = "InvalidParameter.AdvancedOriginRoutingNotSupportAccelerateMainland"
//  INVALIDPARAMETER_ADVANCEDORIGINROUTINGNOTSUPPORTCURRENTPLANORIGINWHITELIST = "InvalidParameter.AdvancedOriginRoutingNotSupportCurrentPlanOriginWhitelist"
//  INVALIDPARAMETER_ADVANCEDORIGINROUTINGNOTSUPPORTMULTIPLYLAYER = "InvalidParameter.AdvancedOriginRoutingNotSupportMultiplyLayer"
//  INVALIDPARAMETER_ADVANCEDORIGINROUTINGNOTSUPPORTOCDIRECTORIGIN = "InvalidParameter.AdvancedOriginRoutingNotSupportOCDirectOrigin"
//  INVALIDPARAMETER_ADVANCEDORIGINROUTINGNOTSUPPORTRACINGPLATFORM = "InvalidParameter.AdvancedOriginRoutingNotSupportRacingPlatform"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_ERRACTIONUNSUPPORTTARGET = "InvalidParameter.ErrActionUnsupportTarget"
//  INVALIDPARAMETER_ERRINVALIDACTION = "InvalidParameter.ErrInvalidAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONORIGINPRIVATEADDRESS = "InvalidParameter.ErrInvalidActionOriginPrivateAddress"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMACTION = "InvalidParameter.ErrInvalidActionParamAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMBADVALUETYPE = "InvalidParameter.ErrInvalidActionParamBadValueType"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMDUPLICATENAME = "InvalidParameter.ErrInvalidActionParamDuplicateName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMNAME = "InvalidParameter.ErrInvalidActionParamName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_ERRINVALIDACTIONTYPE = "InvalidParameter.ErrInvalidActionType"
//  INVALIDPARAMETER_ERRINVALIDCONDITION = "InvalidParameter.ErrInvalidCondition"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONCANNOTONLYCONTAINHOSTWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidConditionCannotOnlyContainHostWhenModifyOriginActionConfigured"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONIGNORECASE = "InvalidParameter.ErrInvalidConditionIgnoreCase"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMEBADNAME = "InvalidParameter.ErrInvalidConditionNameBadName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMETARGETNOTSUPPORTNAME = "InvalidParameter.ErrInvalidConditionNameTargetNotSupportName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADREGULAR = "InvalidParameter.ErrInvalidConditionValueBadRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADURL = "InvalidParameter.ErrInvalidConditionValueBadUrl"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUECONTAINFILENAMEEXTENSION = "InvalidParameter.ErrInvalidConditionValueBadValueContainFileNameExtension"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOLONGVALUE = "InvalidParameter.ErrInvalidConditionValueTooLongValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYREGULAR = "InvalidParameter.ErrInvalidConditionValueTooManyRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYVALUES = "InvalidParameter.ErrInvalidConditionValueTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYWILDCARD = "InvalidParameter.ErrInvalidConditionValueTooManyWildcard"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEZEROLENGTH = "InvalidParameter.ErrInvalidConditionValueZeroLength"
//  INVALIDPARAMETER_ERRNILCONDITION = "InvalidParameter.ErrNilCondition"
//  INVALIDPARAMETER_GRPCREQUIREHTTP2 = "InvalidParameter.GrpcRequireHttp2"
//  INVALIDPARAMETER_HOSTNOTFOUND = "InvalidParameter.HostNotFound"
//  INVALIDPARAMETER_INVALIDADVANCEDORIGINROUTINGUNSUPPORTED = "InvalidParameter.InvalidAdvancedOriginRoutingUnsupported"
//  INVALIDPARAMETER_INVALIDAUTHENTICATION = "InvalidParameter.InvalidAuthentication"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPEEXPIRETIME = "InvalidParameter.InvalidAuthenticationTypeExpireTime"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESIGNPARAM = "InvalidParameter.InvalidAuthenticationTypeSignParam"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDBACKUPSERVERNAME = "InvalidParameter.InvalidBackupServerName"
//  INVALIDPARAMETER_INVALIDCACHECONFIGCACHE = "InvalidParameter.InvalidCacheConfigCache"
//  INVALIDPARAMETER_INVALIDCACHEKEY = "InvalidParameter.InvalidCacheKey"
//  INVALIDPARAMETER_INVALIDCACHEKEYQUERYSTRINGVALUE = "InvalidParameter.InvalidCacheKeyQueryStringValue"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINE = "InvalidParameter.InvalidDynamicRoutine"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINEBILLING = "InvalidParameter.InvalidDynamicRoutineBilling"
//  INVALIDPARAMETER_INVALIDERRORPAGEREDIRECTURL = "InvalidParameter.InvalidErrorPageRedirectUrl"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGETIME = "InvalidParameter.InvalidMaxAgeTime"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPOSTSIZEVALUE = "InvalidParameter.InvalidPostSizeValue"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAMEXFF = "InvalidParameter.InvalidRequestHeaderNameXff"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERVALUE = "InvalidParameter.InvalidRequestHeaderValue"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERNAME = "InvalidParameter.InvalidResponseHeaderName"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERVALUE = "InvalidParameter.InvalidResponseHeaderValue"
//  INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSEXTENSION = "InvalidParameter.InvalidRuleEngineTargetsExtension"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSURL = "InvalidParameter.InvalidRuleEngineTargetsUrl"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_INVALIDSHIELDUNSUPPORTED = "InvalidParameter.InvalidShieldUnsupported"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGEXPIRETIMELIMIT = "InvalidParameter.InvalidStandardDebugExpireTimeLimit"
//  INVALIDPARAMETER_INVALIDUPSTREAMREQUESTQUERYSTRINGVALUE = "InvalidParameter.InvalidUpstreamRequestQueryStringValue"
//  INVALIDPARAMETER_INVALIDURLREDIRECT = "InvalidParameter.InvalidUrlRedirect"
//  INVALIDPARAMETER_INVALIDURLREDIRECTHOST = "InvalidParameter.InvalidUrlRedirectHost"
//  INVALIDPARAMETER_INVALIDURLREDIRECTURL = "InvalidParameter.InvalidUrlRedirectUrl"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_KEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.KeyRulesInvalidQueryStringValue"
//  INVALIDPARAMETER_LOADBALANCEINSTANCEIDISREQUIRED = "InvalidParameter.LoadBalanceInstanceIdIsRequired"
//  INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"
//  INVALIDPARAMETER_ORIGINORIGINGROUPIDISREQUIRED = "InvalidParameter.OriginOriginGroupIdIsRequired"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_SHIELDNOTSUPPORTHOSTCACHEPREFRESH = "InvalidParameter.ShieldNotSupportHostCachePrefresh"
//  INVALIDPARAMETER_SHIELDNOTSUPPORTHOSTORIGINWHITELIST = "InvalidParameter.ShieldNotSupportHostOriginWhitelist"
//  INVALIDPARAMETER_SHIELDNOTSUPPORTHOSTSMARTROUTING = "InvalidParameter.ShieldNotSupportHostSmartRouting"
//  INVALIDPARAMETER_SHIELDNOTSUPPORTZONECACHEPREFRESH = "InvalidParameter.ShieldNotSupportZoneCachePrefresh"
//  INVALIDPARAMETER_SHIELDNOTSUPPORTZONESMARTROUTING = "InvalidParameter.ShieldNotSupportZoneSmartRouting"
//  INVALIDPARAMETER_SHIELDSPACENOTFOUND = "InvalidParameter.ShieldSpaceNotFound"
//  INVALIDPARAMETER_SHIELDSPACEREQUIREDERROR = "InvalidParameter.ShieldSpaceRequiredError"
//  INVALIDPARAMETER_SHIELDSPACESTATUSMUSTONLINE = "InvalidParameter.ShieldSpaceStatusMustOnline"
//  INVALIDPARAMETER_TASKSYSTEMERROR = "InvalidParameter.TaskSystemError"
//  INVALIDPARAMETERVALUE_FORMATMISMATCH = "InvalidParameterValue.FormatMismatch"
//  INVALIDPARAMETERVALUE_GENERALMISMATCH = "InvalidParameterValue.GeneralMismatch"
//  INVALIDPARAMETERVALUE_INCLUDEINVALIDVALUE = "InvalidParameterValue.IncludeInvalidValue"
//  INVALIDPARAMETERVALUE_INVALIDSITEFAILOVERUNSUPPORTED = "InvalidParameterValue.InvalidSiteFailoverUnsupported"
//  INVALIDPARAMETERVALUE_MISSINGNECESSARYPARAM = "InvalidParameterValue.MissingNecessaryParam"
//  INVALIDPARAMETERVALUE_NOTINENUMERATION = "InvalidParameterValue.NotInEnumeration"
//  INVALIDPARAMETERVALUE_NOTWITHINRANGE = "InvalidParameterValue.NotWithinRange"
//  INVALIDPARAMETERVALUE_REGEXMISMATCH = "InvalidParameterValue.RegExMismatch"
//  INVALIDPARAMETERVALUE_SITEFAILOVERNOTSUPPORTHOSTORIGINPULLRATELIMIT = "InvalidParameterValue.SiteFailoverNotSupportHostOriginPullRateLimit"
//  INVALIDPARAMETERVALUE_SITEFAILOVERNOTSUPPORTHOSTORIGINTYPEVOD = "InvalidParameterValue.SiteFailoverNotSupportHostOriginTypeVod"
//  INVALIDPARAMETERVALUE_SITEFAILOVERNOTSUPPORTHOSTUPSTREAMVERIFY = "InvalidParameterValue.SiteFailoverNotSupportHostUpstreamVerify"
//  INVALIDPARAMETERVALUE_UNRECOGNIZABLEVALUE = "InvalidParameterValue.UnrecognizableValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDMULTIPLYLAYERCONFLICT = "OperationDenied.AccelerateMainlandMultiplyLayerConflict"
//  OPERATIONDENIED_ADVANCEDORIGINROUTINGCONFLICTWITHDCIOVERSEALAN = "OperationDenied.AdvancedOriginRoutingConflictWithDciOverseaLan"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_INVALIDADVANCEDDEFENSESECURITYTYPE = "OperationDenied.InvalidAdvancedDefenseSecurityType"
//  OPERATIONDENIED_INVALIDADVANCEDORIGINROUTINGBILLING = "OperationDenied.InvalidAdvancedOriginRoutingBilling"
//  OPERATIONDENIED_NOTINADVANCEDORIGINROUTINGWHITELIST = "OperationDenied.NotInAdvancedOriginRoutingWhiteList"
//  OPERATIONDENIED_NOTINSHIELDSPACEWHITELIST = "OperationDenied.NotInShieldSpaceWhiteList"
//  OPERATIONDENIED_NOTINSITEFAILOVERWHITELIST = "OperationDenied.NotInSiteFailoverWhiteList"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateL7AccRules(request *CreateL7AccRulesRequest) (response *CreateL7AccRulesResponse, err error) {
    return c.CreateL7AccRulesWithContext(context.Background(), request)
}

// CreateL7AccRules
// This API is used to create rules in the [rule engine](https://www.tencentcloud.com/document/product/1552/70901?from_cn_redirect=1), with support for batch creation.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACCESSREDIRECTREGEXERROR = "InvalidParameter.AccessRedirectRegexError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_ADVANCEDORIGINROUTINGNOTSUPPORTACCELERATEMAINLAND = "InvalidParameter.AdvancedOriginRoutingNotSupportAccelerateMainland"
//  INVALIDPARAMETER_ADVANCEDORIGINROUTINGNOTSUPPORTCURRENTPLANORIGINWHITELIST = "InvalidParameter.AdvancedOriginRoutingNotSupportCurrentPlanOriginWhitelist"
//  INVALIDPARAMETER_ADVANCEDORIGINROUTINGNOTSUPPORTMULTIPLYLAYER = "InvalidParameter.AdvancedOriginRoutingNotSupportMultiplyLayer"
//  INVALIDPARAMETER_ADVANCEDORIGINROUTINGNOTSUPPORTOCDIRECTORIGIN = "InvalidParameter.AdvancedOriginRoutingNotSupportOCDirectOrigin"
//  INVALIDPARAMETER_ADVANCEDORIGINROUTINGNOTSUPPORTRACINGPLATFORM = "InvalidParameter.AdvancedOriginRoutingNotSupportRacingPlatform"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_ERRACTIONUNSUPPORTTARGET = "InvalidParameter.ErrActionUnsupportTarget"
//  INVALIDPARAMETER_ERRINVALIDACTION = "InvalidParameter.ErrInvalidAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONORIGINPRIVATEADDRESS = "InvalidParameter.ErrInvalidActionOriginPrivateAddress"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMACTION = "InvalidParameter.ErrInvalidActionParamAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMBADVALUETYPE = "InvalidParameter.ErrInvalidActionParamBadValueType"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMDUPLICATENAME = "InvalidParameter.ErrInvalidActionParamDuplicateName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMNAME = "InvalidParameter.ErrInvalidActionParamName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_ERRINVALIDACTIONTYPE = "InvalidParameter.ErrInvalidActionType"
//  INVALIDPARAMETER_ERRINVALIDCONDITION = "InvalidParameter.ErrInvalidCondition"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONCANNOTONLYCONTAINHOSTWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidConditionCannotOnlyContainHostWhenModifyOriginActionConfigured"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONIGNORECASE = "InvalidParameter.ErrInvalidConditionIgnoreCase"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMEBADNAME = "InvalidParameter.ErrInvalidConditionNameBadName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMETARGETNOTSUPPORTNAME = "InvalidParameter.ErrInvalidConditionNameTargetNotSupportName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADREGULAR = "InvalidParameter.ErrInvalidConditionValueBadRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADURL = "InvalidParameter.ErrInvalidConditionValueBadUrl"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUECONTAINFILENAMEEXTENSION = "InvalidParameter.ErrInvalidConditionValueBadValueContainFileNameExtension"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOLONGVALUE = "InvalidParameter.ErrInvalidConditionValueTooLongValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYREGULAR = "InvalidParameter.ErrInvalidConditionValueTooManyRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYVALUES = "InvalidParameter.ErrInvalidConditionValueTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYWILDCARD = "InvalidParameter.ErrInvalidConditionValueTooManyWildcard"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEZEROLENGTH = "InvalidParameter.ErrInvalidConditionValueZeroLength"
//  INVALIDPARAMETER_ERRNILCONDITION = "InvalidParameter.ErrNilCondition"
//  INVALIDPARAMETER_GRPCREQUIREHTTP2 = "InvalidParameter.GrpcRequireHttp2"
//  INVALIDPARAMETER_HOSTNOTFOUND = "InvalidParameter.HostNotFound"
//  INVALIDPARAMETER_INVALIDADVANCEDORIGINROUTINGUNSUPPORTED = "InvalidParameter.InvalidAdvancedOriginRoutingUnsupported"
//  INVALIDPARAMETER_INVALIDAUTHENTICATION = "InvalidParameter.InvalidAuthentication"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPEEXPIRETIME = "InvalidParameter.InvalidAuthenticationTypeExpireTime"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESIGNPARAM = "InvalidParameter.InvalidAuthenticationTypeSignParam"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDBACKUPSERVERNAME = "InvalidParameter.InvalidBackupServerName"
//  INVALIDPARAMETER_INVALIDCACHECONFIGCACHE = "InvalidParameter.InvalidCacheConfigCache"
//  INVALIDPARAMETER_INVALIDCACHEKEY = "InvalidParameter.InvalidCacheKey"
//  INVALIDPARAMETER_INVALIDCACHEKEYQUERYSTRINGVALUE = "InvalidParameter.InvalidCacheKeyQueryStringValue"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINE = "InvalidParameter.InvalidDynamicRoutine"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINEBILLING = "InvalidParameter.InvalidDynamicRoutineBilling"
//  INVALIDPARAMETER_INVALIDERRORPAGEREDIRECTURL = "InvalidParameter.InvalidErrorPageRedirectUrl"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGETIME = "InvalidParameter.InvalidMaxAgeTime"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPOSTSIZEVALUE = "InvalidParameter.InvalidPostSizeValue"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAMEXFF = "InvalidParameter.InvalidRequestHeaderNameXff"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERVALUE = "InvalidParameter.InvalidRequestHeaderValue"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERNAME = "InvalidParameter.InvalidResponseHeaderName"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERVALUE = "InvalidParameter.InvalidResponseHeaderValue"
//  INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSEXTENSION = "InvalidParameter.InvalidRuleEngineTargetsExtension"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSURL = "InvalidParameter.InvalidRuleEngineTargetsUrl"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_INVALIDSHIELDUNSUPPORTED = "InvalidParameter.InvalidShieldUnsupported"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGEXPIRETIMELIMIT = "InvalidParameter.InvalidStandardDebugExpireTimeLimit"
//  INVALIDPARAMETER_INVALIDUPSTREAMREQUESTQUERYSTRINGVALUE = "InvalidParameter.InvalidUpstreamRequestQueryStringValue"
//  INVALIDPARAMETER_INVALIDURLREDIRECT = "InvalidParameter.InvalidUrlRedirect"
//  INVALIDPARAMETER_INVALIDURLREDIRECTHOST = "InvalidParameter.InvalidUrlRedirectHost"
//  INVALIDPARAMETER_INVALIDURLREDIRECTURL = "InvalidParameter.InvalidUrlRedirectUrl"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_KEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.KeyRulesInvalidQueryStringValue"
//  INVALIDPARAMETER_LOADBALANCEINSTANCEIDISREQUIRED = "InvalidParameter.LoadBalanceInstanceIdIsRequired"
//  INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"
//  INVALIDPARAMETER_ORIGINORIGINGROUPIDISREQUIRED = "InvalidParameter.OriginOriginGroupIdIsRequired"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_SHIELDNOTSUPPORTHOSTCACHEPREFRESH = "InvalidParameter.ShieldNotSupportHostCachePrefresh"
//  INVALIDPARAMETER_SHIELDNOTSUPPORTHOSTORIGINWHITELIST = "InvalidParameter.ShieldNotSupportHostOriginWhitelist"
//  INVALIDPARAMETER_SHIELDNOTSUPPORTHOSTSMARTROUTING = "InvalidParameter.ShieldNotSupportHostSmartRouting"
//  INVALIDPARAMETER_SHIELDNOTSUPPORTZONECACHEPREFRESH = "InvalidParameter.ShieldNotSupportZoneCachePrefresh"
//  INVALIDPARAMETER_SHIELDNOTSUPPORTZONESMARTROUTING = "InvalidParameter.ShieldNotSupportZoneSmartRouting"
//  INVALIDPARAMETER_SHIELDSPACENOTFOUND = "InvalidParameter.ShieldSpaceNotFound"
//  INVALIDPARAMETER_SHIELDSPACEREQUIREDERROR = "InvalidParameter.ShieldSpaceRequiredError"
//  INVALIDPARAMETER_SHIELDSPACESTATUSMUSTONLINE = "InvalidParameter.ShieldSpaceStatusMustOnline"
//  INVALIDPARAMETER_TASKSYSTEMERROR = "InvalidParameter.TaskSystemError"
//  INVALIDPARAMETERVALUE_FORMATMISMATCH = "InvalidParameterValue.FormatMismatch"
//  INVALIDPARAMETERVALUE_GENERALMISMATCH = "InvalidParameterValue.GeneralMismatch"
//  INVALIDPARAMETERVALUE_INCLUDEINVALIDVALUE = "InvalidParameterValue.IncludeInvalidValue"
//  INVALIDPARAMETERVALUE_INVALIDSITEFAILOVERUNSUPPORTED = "InvalidParameterValue.InvalidSiteFailoverUnsupported"
//  INVALIDPARAMETERVALUE_MISSINGNECESSARYPARAM = "InvalidParameterValue.MissingNecessaryParam"
//  INVALIDPARAMETERVALUE_NOTINENUMERATION = "InvalidParameterValue.NotInEnumeration"
//  INVALIDPARAMETERVALUE_NOTWITHINRANGE = "InvalidParameterValue.NotWithinRange"
//  INVALIDPARAMETERVALUE_REGEXMISMATCH = "InvalidParameterValue.RegExMismatch"
//  INVALIDPARAMETERVALUE_SITEFAILOVERNOTSUPPORTHOSTORIGINPULLRATELIMIT = "InvalidParameterValue.SiteFailoverNotSupportHostOriginPullRateLimit"
//  INVALIDPARAMETERVALUE_SITEFAILOVERNOTSUPPORTHOSTORIGINTYPEVOD = "InvalidParameterValue.SiteFailoverNotSupportHostOriginTypeVod"
//  INVALIDPARAMETERVALUE_SITEFAILOVERNOTSUPPORTHOSTUPSTREAMVERIFY = "InvalidParameterValue.SiteFailoverNotSupportHostUpstreamVerify"
//  INVALIDPARAMETERVALUE_UNRECOGNIZABLEVALUE = "InvalidParameterValue.UnrecognizableValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDMULTIPLYLAYERCONFLICT = "OperationDenied.AccelerateMainlandMultiplyLayerConflict"
//  OPERATIONDENIED_ADVANCEDORIGINROUTINGCONFLICTWITHDCIOVERSEALAN = "OperationDenied.AdvancedOriginRoutingConflictWithDciOverseaLan"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_INVALIDADVANCEDDEFENSESECURITYTYPE = "OperationDenied.InvalidAdvancedDefenseSecurityType"
//  OPERATIONDENIED_INVALIDADVANCEDORIGINROUTINGBILLING = "OperationDenied.InvalidAdvancedOriginRoutingBilling"
//  OPERATIONDENIED_NOTINADVANCEDORIGINROUTINGWHITELIST = "OperationDenied.NotInAdvancedOriginRoutingWhiteList"
//  OPERATIONDENIED_NOTINSHIELDSPACEWHITELIST = "OperationDenied.NotInShieldSpaceWhiteList"
//  OPERATIONDENIED_NOTINSITEFAILOVERWHITELIST = "OperationDenied.NotInSiteFailoverWhiteList"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateL7AccRulesWithContext(ctx context.Context, request *CreateL7AccRulesRequest) (response *CreateL7AccRulesResponse, err error) {
    if request == nil {
        request = NewCreateL7AccRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreateL7AccRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateL7AccRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateL7AccRulesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLoadBalancerRequest() (request *CreateLoadBalancerRequest) {
    request = &CreateLoadBalancerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateLoadBalancer")
    
    
    return
}

func NewCreateLoadBalancerResponse() (response *CreateLoadBalancerResponse) {
    response = &CreateLoadBalancerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLoadBalancer
// This API is used to create a LoadBalancer. For details, see [Quickly Create Load Balancers](https://www.tencentcloud.com/document/product/1552/104223?from_cn_redirect=1). The load balancing feature is in beta test. If you need to use it, [contact us](https://www.tencentcloud.com/contact-us).
//
// error code that may be returned:
//  INVALIDPARAMETER_LOADBALANCERBINDORIGINGROUPINVALID = "InvalidParameter.LoadBalancerBindOriginGroupInvalid"
//  INVALIDPARAMETER_LOADBALANCERNAMEREPEATED = "InvalidParameter.LoadBalancerNameRepeated"
//  INVALIDPARAMETER_ORIGINGROUPTYPECANNOTMATCHLBTYPE = "InvalidParameter.OriginGroupTypeCanNotMatchLBType"
//  INVALIDPARAMETER_SOMEORIGINGROUPNOTEXIST = "InvalidParameter.SomeOriginGroupNotExist"
//  LIMITEXCEEDED_LOADBALANCINGCOUNTLIMITEXCEEDED = "LimitExceeded.LoadBalancingCountLimitExceeded"
func (c *Client) CreateLoadBalancer(request *CreateLoadBalancerRequest) (response *CreateLoadBalancerResponse, err error) {
    return c.CreateLoadBalancerWithContext(context.Background(), request)
}

// CreateLoadBalancer
// This API is used to create a LoadBalancer. For details, see [Quickly Create Load Balancers](https://www.tencentcloud.com/document/product/1552/104223?from_cn_redirect=1). The load balancing feature is in beta test. If you need to use it, [contact us](https://www.tencentcloud.com/contact-us).
//
// error code that may be returned:
//  INVALIDPARAMETER_LOADBALANCERBINDORIGINGROUPINVALID = "InvalidParameter.LoadBalancerBindOriginGroupInvalid"
//  INVALIDPARAMETER_LOADBALANCERNAMEREPEATED = "InvalidParameter.LoadBalancerNameRepeated"
//  INVALIDPARAMETER_ORIGINGROUPTYPECANNOTMATCHLBTYPE = "InvalidParameter.OriginGroupTypeCanNotMatchLBType"
//  INVALIDPARAMETER_SOMEORIGINGROUPNOTEXIST = "InvalidParameter.SomeOriginGroupNotExist"
//  LIMITEXCEEDED_LOADBALANCINGCOUNTLIMITEXCEEDED = "LimitExceeded.LoadBalancingCountLimitExceeded"
func (c *Client) CreateLoadBalancerWithContext(ctx context.Context, request *CreateLoadBalancerRequest) (response *CreateLoadBalancerResponse, err error) {
    if request == nil {
        request = NewCreateLoadBalancerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreateLoadBalancer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLoadBalancer require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLoadBalancerResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLogAnalysisDownloadTaskRequest() (request *CreateLogAnalysisDownloadTaskRequest) {
    request = &CreateLogAnalysisDownloadTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateLogAnalysisDownloadTask")
    
    
    return
}

func NewCreateLogAnalysisDownloadTaskResponse() (response *CreateLogAnalysisDownloadTaskResponse) {
    response = &CreateLogAnalysisDownloadTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLogAnalysisDownloadTask
// This API is used to create a log analysis download task. After creation, you can query the download task through the DescribeLogAnalysisDownloadTasks API.
//
// Note:
//
// 1. Supports up to 50 million log entries per download.
//
// 2. Log files will be retained for 3 days.
//
// 3. If multiple tasks exist at the same time, they will be processed sequentially based on task creation time.
//
// error code that may be returned:
//  INVALIDPARAMETER_LOADBALANCERBINDORIGINGROUPINVALID = "InvalidParameter.LoadBalancerBindOriginGroupInvalid"
//  INVALIDPARAMETER_LOADBALANCERNAMEREPEATED = "InvalidParameter.LoadBalancerNameRepeated"
//  INVALIDPARAMETER_ORIGINGROUPTYPECANNOTMATCHLBTYPE = "InvalidParameter.OriginGroupTypeCanNotMatchLBType"
//  INVALIDPARAMETER_SOMEORIGINGROUPNOTEXIST = "InvalidParameter.SomeOriginGroupNotExist"
//  LIMITEXCEEDED_LOADBALANCINGCOUNTLIMITEXCEEDED = "LimitExceeded.LoadBalancingCountLimitExceeded"
func (c *Client) CreateLogAnalysisDownloadTask(request *CreateLogAnalysisDownloadTaskRequest) (response *CreateLogAnalysisDownloadTaskResponse, err error) {
    return c.CreateLogAnalysisDownloadTaskWithContext(context.Background(), request)
}

// CreateLogAnalysisDownloadTask
// This API is used to create a log analysis download task. After creation, you can query the download task through the DescribeLogAnalysisDownloadTasks API.
//
// Note:
//
// 1. Supports up to 50 million log entries per download.
//
// 2. Log files will be retained for 3 days.
//
// 3. If multiple tasks exist at the same time, they will be processed sequentially based on task creation time.
//
// error code that may be returned:
//  INVALIDPARAMETER_LOADBALANCERBINDORIGINGROUPINVALID = "InvalidParameter.LoadBalancerBindOriginGroupInvalid"
//  INVALIDPARAMETER_LOADBALANCERNAMEREPEATED = "InvalidParameter.LoadBalancerNameRepeated"
//  INVALIDPARAMETER_ORIGINGROUPTYPECANNOTMATCHLBTYPE = "InvalidParameter.OriginGroupTypeCanNotMatchLBType"
//  INVALIDPARAMETER_SOMEORIGINGROUPNOTEXIST = "InvalidParameter.SomeOriginGroupNotExist"
//  LIMITEXCEEDED_LOADBALANCINGCOUNTLIMITEXCEEDED = "LimitExceeded.LoadBalancingCountLimitExceeded"
func (c *Client) CreateLogAnalysisDownloadTaskWithContext(ctx context.Context, request *CreateLogAnalysisDownloadTaskRequest) (response *CreateLogAnalysisDownloadTaskResponse, err error) {
    if request == nil {
        request = NewCreateLogAnalysisDownloadTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreateLogAnalysisDownloadTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLogAnalysisDownloadTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLogAnalysisDownloadTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMultiPathGatewayRequest() (request *CreateMultiPathGatewayRequest) {
    request = &CreateMultiPathGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateMultiPathGateway")
    
    
    return
}

func NewCreateMultiPathGatewayResponse() (response *CreateMultiPathGatewayResponse) {
    response = &CreateMultiPathGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateMultiPathGateway
// This API is used to create a multi-channel security acceleration gateway, including Cloud Gateway and Private Gateway. Use the DescribeMultiPathGateway interface to query the status. If the status is online, the creation succeeded.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateMultiPathGateway(request *CreateMultiPathGatewayRequest) (response *CreateMultiPathGatewayResponse, err error) {
    return c.CreateMultiPathGatewayWithContext(context.Background(), request)
}

// CreateMultiPathGateway
// This API is used to create a multi-channel security acceleration gateway, including Cloud Gateway and Private Gateway. Use the DescribeMultiPathGateway interface to query the status. If the status is online, the creation succeeded.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateMultiPathGatewayWithContext(ctx context.Context, request *CreateMultiPathGatewayRequest) (response *CreateMultiPathGatewayResponse, err error) {
    if request == nil {
        request = NewCreateMultiPathGatewayRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreateMultiPathGateway")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMultiPathGateway require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMultiPathGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMultiPathGatewayLineRequest() (request *CreateMultiPathGatewayLineRequest) {
    request = &CreateMultiPathGatewayLineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateMultiPathGatewayLine")
    
    
    return
}

func NewCreateMultiPathGatewayLineResponse() (response *CreateMultiPathGatewayLineResponse) {
    response = &CreateMultiPathGatewayLineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateMultiPathGatewayLine
// This API is used to create lines for integrating with the multi-channel security acceleration gateway, including EdgeOne Layer 4 proxy lines and custom lines.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateMultiPathGatewayLine(request *CreateMultiPathGatewayLineRequest) (response *CreateMultiPathGatewayLineResponse, err error) {
    return c.CreateMultiPathGatewayLineWithContext(context.Background(), request)
}

// CreateMultiPathGatewayLine
// This API is used to create lines for integrating with the multi-channel security acceleration gateway, including EdgeOne Layer 4 proxy lines and custom lines.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateMultiPathGatewayLineWithContext(ctx context.Context, request *CreateMultiPathGatewayLineRequest) (response *CreateMultiPathGatewayLineResponse, err error) {
    if request == nil {
        request = NewCreateMultiPathGatewayLineRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreateMultiPathGatewayLine")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMultiPathGatewayLine require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMultiPathGatewayLineResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMultiPathGatewaySecretKeyRequest() (request *CreateMultiPathGatewaySecretKeyRequest) {
    request = &CreateMultiPathGatewaySecretKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateMultiPathGatewaySecretKey")
    
    
    return
}

func NewCreateMultiPathGatewaySecretKeyResponse() (response *CreateMultiPathGatewaySecretKeyResponse) {
    response = &CreateMultiPathGatewaySecretKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateMultiPathGatewaySecretKey
// This API is used to create a key for integrating with a multi-channel security acceleration gateway. Customers integrate with the multi-channel security acceleration gateway based on key signatures. Each site contains only one key, which can be used to integrate all gateways under the site. The key can be queried through the DescribeMultiPathGatewaySecretKey API.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateMultiPathGatewaySecretKey(request *CreateMultiPathGatewaySecretKeyRequest) (response *CreateMultiPathGatewaySecretKeyResponse, err error) {
    return c.CreateMultiPathGatewaySecretKeyWithContext(context.Background(), request)
}

// CreateMultiPathGatewaySecretKey
// This API is used to create a key for integrating with a multi-channel security acceleration gateway. Customers integrate with the multi-channel security acceleration gateway based on key signatures. Each site contains only one key, which can be used to integrate all gateways under the site. The key can be queried through the DescribeMultiPathGatewaySecretKey API.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateMultiPathGatewaySecretKeyWithContext(ctx context.Context, request *CreateMultiPathGatewaySecretKeyRequest) (response *CreateMultiPathGatewaySecretKeyResponse, err error) {
    if request == nil {
        request = NewCreateMultiPathGatewaySecretKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreateMultiPathGatewaySecretKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMultiPathGatewaySecretKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMultiPathGatewaySecretKeyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOriginGroupRequest() (request *CreateOriginGroupRequest) {
    request = &CreateOriginGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateOriginGroup")
    
    
    return
}

func NewCreateOriginGroupResponse() (response *CreateOriginGroupResponse) {
    response = &CreateOriginGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateOriginGroup
// Create an origin server group to manage origin servers by group. The origin server group configured here can be referenced in features such as **adding an acceleration domain name** and **Layer 4 proxy**.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INVALIDPARAMETER_HOSTHEADERINVALID = "InvalidParameter.HostHeaderInvalid"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDORIGINGROUPTYPE = "InvalidParameter.InvalidOriginGroupType"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_ORIGINNAMEEXISTS = "InvalidParameter.OriginNameExists"
//  INVALIDPARAMETER_ORIGINRECORDFORMATERROR = "InvalidParameter.OriginRecordFormatError"
//  INVALIDPARAMETER_ORIGINRECORDWEIGHTVALUE = "InvalidParameter.OriginRecordWeightValue"
//  INVALIDPARAMETER_ORIGINTHIRDPARTYPARAMFORMATERROR = "InvalidParameter.OriginThirdPartyParamFormatError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_LOADBALANCINGZONEISNOTACTIVE = "OperationDenied.LoadBalancingZoneIsNotActive"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) CreateOriginGroup(request *CreateOriginGroupRequest) (response *CreateOriginGroupResponse, err error) {
    return c.CreateOriginGroupWithContext(context.Background(), request)
}

// CreateOriginGroup
// Create an origin server group to manage origin servers by group. The origin server group configured here can be referenced in features such as **adding an acceleration domain name** and **Layer 4 proxy**.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INVALIDPARAMETER_HOSTHEADERINVALID = "InvalidParameter.HostHeaderInvalid"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDORIGINGROUPTYPE = "InvalidParameter.InvalidOriginGroupType"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_ORIGINNAMEEXISTS = "InvalidParameter.OriginNameExists"
//  INVALIDPARAMETER_ORIGINRECORDFORMATERROR = "InvalidParameter.OriginRecordFormatError"
//  INVALIDPARAMETER_ORIGINRECORDWEIGHTVALUE = "InvalidParameter.OriginRecordWeightValue"
//  INVALIDPARAMETER_ORIGINTHIRDPARTYPARAMFORMATERROR = "InvalidParameter.OriginThirdPartyParamFormatError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_LOADBALANCINGZONEISNOTACTIVE = "OperationDenied.LoadBalancingZoneIsNotActive"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) CreateOriginGroupWithContext(ctx context.Context, request *CreateOriginGroupRequest) (response *CreateOriginGroupResponse, err error) {
    if request == nil {
        request = NewCreateOriginGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreateOriginGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOriginGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOriginGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePlanRequest() (request *CreatePlanRequest) {
    request = &CreatePlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreatePlan")
    
    
    return
}

func NewCreatePlanResponse() (response *CreatePlanResponse) {
    response = &CreatePlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePlan
// If you need to use the Edgeone product, you need to create a billing plan through this interface.
//
// > After creating a package, you need to complete creating a site and binding the plan through [CreateZone](https://www.tencentcloud.com/document/product/1552/80719?from_cn_redirect=1) before Edgeone can provide services properly.
//
// error code that may be returned:
//  FAILEDOPERATION_INSUFFICIENTACCOUNTBALANCE = "FailedOperation.InsufficientAccountBalance"
//  INVALIDPARAMETER_INVALIDAUTOUSEVOUCHER = "InvalidParameter.InvalidAutoUseVoucher"
//  INVALIDPARAMETER_INVALIDPERIOD = "InvalidParameter.InvalidPeriod"
//  INVALIDPARAMETER_INVALIDPLANTYPE = "InvalidParameter.InvalidPlanType"
//  INVALIDPARAMETER_INVALIDRENEWFLAG = "InvalidParameter.InvalidRenewFlag"
//  OPERATIONDENIED_PLEASECONTACTBUSINESSPERSONNEL = "OperationDenied.PleaseContactBusinessPersonnel"
func (c *Client) CreatePlan(request *CreatePlanRequest) (response *CreatePlanResponse, err error) {
    return c.CreatePlanWithContext(context.Background(), request)
}

// CreatePlan
// If you need to use the Edgeone product, you need to create a billing plan through this interface.
//
// > After creating a package, you need to complete creating a site and binding the plan through [CreateZone](https://www.tencentcloud.com/document/product/1552/80719?from_cn_redirect=1) before Edgeone can provide services properly.
//
// error code that may be returned:
//  FAILEDOPERATION_INSUFFICIENTACCOUNTBALANCE = "FailedOperation.InsufficientAccountBalance"
//  INVALIDPARAMETER_INVALIDAUTOUSEVOUCHER = "InvalidParameter.InvalidAutoUseVoucher"
//  INVALIDPARAMETER_INVALIDPERIOD = "InvalidParameter.InvalidPeriod"
//  INVALIDPARAMETER_INVALIDPLANTYPE = "InvalidParameter.InvalidPlanType"
//  INVALIDPARAMETER_INVALIDRENEWFLAG = "InvalidParameter.InvalidRenewFlag"
//  OPERATIONDENIED_PLEASECONTACTBUSINESSPERSONNEL = "OperationDenied.PleaseContactBusinessPersonnel"
func (c *Client) CreatePlanWithContext(ctx context.Context, request *CreatePlanRequest) (response *CreatePlanResponse, err error) {
    if request == nil {
        request = NewCreatePlanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreatePlan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePlanResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePlanForZoneRequest() (request *CreatePlanForZoneRequest) {
    request = &CreatePlanForZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreatePlanForZone")
    
    
    return
}

func NewCreatePlanForZoneResponse() (response *CreatePlanForZoneResponse) {
    response = &CreatePlanForZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePlanForZone
// Purchase a package for the selected site not bound to a package.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSUFFICIENTACCOUNTBALANCE = "FailedOperation.InsufficientAccountBalance"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_ZONEHASBEENBOUND = "InvalidParameter.ZoneHasBeenBound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DOMAINNOICP = "OperationDenied.DomainNoICP"
func (c *Client) CreatePlanForZone(request *CreatePlanForZoneRequest) (response *CreatePlanForZoneResponse, err error) {
    return c.CreatePlanForZoneWithContext(context.Background(), request)
}

// CreatePlanForZone
// Purchase a package for the selected site not bound to a package.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSUFFICIENTACCOUNTBALANCE = "FailedOperation.InsufficientAccountBalance"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_ZONEHASBEENBOUND = "InvalidParameter.ZoneHasBeenBound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DOMAINNOICP = "OperationDenied.DomainNoICP"
func (c *Client) CreatePlanForZoneWithContext(ctx context.Context, request *CreatePlanForZoneRequest) (response *CreatePlanForZoneResponse, err error) {
    if request == nil {
        request = NewCreatePlanForZoneRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreatePlanForZone")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePlanForZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePlanForZoneResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrefetchTaskRequest() (request *CreatePrefetchTaskRequest) {
    request = &CreatePrefetchTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreatePrefetchTask")
    
    
    return
}

func NewCreatePrefetchTaskResponse() (response *CreatePrefetchTaskResponse) {
    response = &CreatePrefetchTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePrefetchTask
// This API is used to create a preheating task.
//
// error code that may be returned:
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DOMAINCONFIG = "InternalError.DomainConfig"
//  INTERNALERROR_FAILEDTOGENERATEURL = "InternalError.FailedToGenerateUrl"
//  INTERNALERROR_QUOTASYSTEM = "InternalError.QuotaSystem"
//  INVALIDPARAMETER_DOMAINNOTFOUND = "InvalidParameter.DomainNotFound"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETER_TARGET = "InvalidParameter.Target"
//  INVALIDPARAMETER_TASKNOTGENERATED = "InvalidParameter.TaskNotGenerated"
//  INVALIDPARAMETER_UPLOADURL = "InvalidParameter.UploadUrl"
//  LIMITEXCEEDED_BATCHQUOTA = "LimitExceeded.BatchQuota"
//  LIMITEXCEEDED_DAILYQUOTA = "LimitExceeded.DailyQuota"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) CreatePrefetchTask(request *CreatePrefetchTaskRequest) (response *CreatePrefetchTaskResponse, err error) {
    return c.CreatePrefetchTaskWithContext(context.Background(), request)
}

// CreatePrefetchTask
// This API is used to create a preheating task.
//
// error code that may be returned:
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DOMAINCONFIG = "InternalError.DomainConfig"
//  INTERNALERROR_FAILEDTOGENERATEURL = "InternalError.FailedToGenerateUrl"
//  INTERNALERROR_QUOTASYSTEM = "InternalError.QuotaSystem"
//  INVALIDPARAMETER_DOMAINNOTFOUND = "InvalidParameter.DomainNotFound"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETER_TARGET = "InvalidParameter.Target"
//  INVALIDPARAMETER_TASKNOTGENERATED = "InvalidParameter.TaskNotGenerated"
//  INVALIDPARAMETER_UPLOADURL = "InvalidParameter.UploadUrl"
//  LIMITEXCEEDED_BATCHQUOTA = "LimitExceeded.BatchQuota"
//  LIMITEXCEEDED_DAILYQUOTA = "LimitExceeded.DailyQuota"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) CreatePrefetchTaskWithContext(ctx context.Context, request *CreatePrefetchTaskRequest) (response *CreatePrefetchTaskResponse, err error) {
    if request == nil {
        request = NewCreatePrefetchTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreatePrefetchTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrefetchTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrefetchTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePurgeTaskRequest() (request *CreatePurgeTaskRequest) {
    request = &CreatePurgeTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreatePurgeTask")
    
    
    return
}

func NewCreatePurgeTaskResponse() (response *CreatePurgeTaskResponse) {
    response = &CreatePurgeTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePurgeTask
// When origin server resources update but the node cache TTL has not expired, users still access old resources. At this point, use this API to update node resources. The following two methods can trigger updates: <li>Direct deletion: Delete the node cache without any validation. User requests trigger origin-pull.</li><li>Mark expiration: Mark node resources as expired. User requests trigger back-to-origin validation, which sends HTTP conditional requests with If-None-Match and If-Modified-Since headers. If the origin server responds with 200, the node fetches new resources from the origin and updates the cache. If the origin server responds with 304, the cache will not be updated.</li>
//
// 
//
// For details on purging cache tasks, see [Purge Cache](https://www.tencentcloud.com/document/product/1552/70759?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DOMAINCONFIG = "InternalError.DomainConfig"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_QUOTASYSTEM = "InternalError.QuotaSystem"
//  INVALIDPARAMETER_DOMAINNOTFOUND = "InvalidParameter.DomainNotFound"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETER_TARGET = "InvalidParameter.Target"
//  INVALIDPARAMETER_TASKNOTGENERATED = "InvalidParameter.TaskNotGenerated"
//  INVALIDPARAMETER_UPLOADURL = "InvalidParameter.UploadUrl"
//  LIMITEXCEEDED_BATCHQUOTA = "LimitExceeded.BatchQuota"
//  LIMITEXCEEDED_DAILYQUOTA = "LimitExceeded.DailyQuota"
//  LIMITEXCEEDED_PACKNOTALLOW = "LimitExceeded.PackNotAllow"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) CreatePurgeTask(request *CreatePurgeTaskRequest) (response *CreatePurgeTaskResponse, err error) {
    return c.CreatePurgeTaskWithContext(context.Background(), request)
}

// CreatePurgeTask
// When origin server resources update but the node cache TTL has not expired, users still access old resources. At this point, use this API to update node resources. The following two methods can trigger updates: <li>Direct deletion: Delete the node cache without any validation. User requests trigger origin-pull.</li><li>Mark expiration: Mark node resources as expired. User requests trigger back-to-origin validation, which sends HTTP conditional requests with If-None-Match and If-Modified-Since headers. If the origin server responds with 200, the node fetches new resources from the origin and updates the cache. If the origin server responds with 304, the cache will not be updated.</li>
//
// 
//
// For details on purging cache tasks, see [Purge Cache](https://www.tencentcloud.com/document/product/1552/70759?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DOMAINCONFIG = "InternalError.DomainConfig"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_QUOTASYSTEM = "InternalError.QuotaSystem"
//  INVALIDPARAMETER_DOMAINNOTFOUND = "InvalidParameter.DomainNotFound"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETER_TARGET = "InvalidParameter.Target"
//  INVALIDPARAMETER_TASKNOTGENERATED = "InvalidParameter.TaskNotGenerated"
//  INVALIDPARAMETER_UPLOADURL = "InvalidParameter.UploadUrl"
//  LIMITEXCEEDED_BATCHQUOTA = "LimitExceeded.BatchQuota"
//  LIMITEXCEEDED_DAILYQUOTA = "LimitExceeded.DailyQuota"
//  LIMITEXCEEDED_PACKNOTALLOW = "LimitExceeded.PackNotAllow"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) CreatePurgeTaskWithContext(ctx context.Context, request *CreatePurgeTaskRequest) (response *CreatePurgeTaskResponse, err error) {
    if request == nil {
        request = NewCreatePurgeTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreatePurgeTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePurgeTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePurgeTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRealtimeLogDeliveryTaskRequest() (request *CreateRealtimeLogDeliveryTaskRequest) {
    request = &CreateRealtimeLogDeliveryTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateRealtimeLogDeliveryTask")
    
    
    return
}

func NewCreateRealtimeLogDeliveryTaskResponse() (response *CreateRealtimeLogDeliveryTaskResponse) {
    response = &CreateRealtimeLogDeliveryTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRealtimeLogDeliveryTask
// This API is used to create a real-time log delivery task.
//
// 
//
// **Restrictions**
//
// 
//
// An entity (a Layer 7 domain name, a Layer 4 proxy instance, or an Edge Function instance) under the same combination of `LogType` and `Area` can only have one task created for each `TaskType`. The mutual exclusion rules for different `LogType` values are as follows:
//
// 
//
// | LogType | Conflicting TaskType |
//
// |---|---|
//
// | Layer 4 proxy logs, Edge Function execution logs | Choose one of `custom_endpoint`, `s3`, and `aws_s3` |
//
// 
//
// When `TaskType` is `log_analysis`:
//
// 
//
// <ul>
//
// <li>Only `LogType` = `l7-access-logs` or `web-attack` is supported.</li>
//
// <li>Under the same combination of `ZoneId` and `Area`, only one task can be created for each `LogType`.</li>
//
// </ul>
//
// 
//
// **Recommendation**: Before creating a task, call the [DescribeRealtimeLogDeliveryTasks](https://www.tencentcloud.com/document/product/1145/59503) API to query existing tasks and avoid duplicate additions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATECLSLOGSETFAILED = "FailedOperation.CreateClsLogSetFailed"
//  FAILEDOPERATION_CREATECLSLOGTOPICTASKFAILED = "FailedOperation.CreateClsLogTopicTaskFailed"
//  FAILEDOPERATION_CREATELOGTOPICTASKAUTHFAILURE = "FailedOperation.CreateLogTopicTaskAuthFailure"
//  FAILEDOPERATION_REALTIMELOGAUTHFAILURE = "FailedOperation.RealtimeLogAuthFailure"
//  FAILEDOPERATION_REALTIMELOGNOTFOUND = "FailedOperation.RealtimeLogNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDLOGFORMATFIELDDELIMITER = "InvalidParameter.InvalidLogFormatFieldDelimiter"
//  INVALIDPARAMETER_INVALIDLOGFORMATFORMATTYPE = "InvalidParameter.InvalidLogFormatFormatType"
//  INVALIDPARAMETER_INVALIDLOGFORMATRECORDDELIMITER = "InvalidParameter.InvalidLogFormatRecordDelimiter"
//  INVALIDPARAMETER_REALTIMELOGENTITYALREADYCREATED = "InvalidParameter.RealtimeLogEntityAlreadyCreated"
//  INVALIDPARAMETER_REALTIMELOGINVALIDDELIVERYAREA = "InvalidParameter.RealtimeLogInvalidDeliveryArea"
//  INVALIDPARAMETER_REALTIMELOGINVALIDLOGTYPE = "InvalidParameter.RealtimeLogInvalidLogType"
//  INVALIDPARAMETER_REALTIMELOGINVALIDTASKTYPE = "InvalidParameter.RealtimeLogInvalidTaskType"
//  INVALIDPARAMETER_REALTIMELOGNUMSEXCEEDLIMIT = "InvalidParameter.RealtimeLogNumsExceedLimit"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CUSTOMLOGFIELDREGEXLIMITEXCEEDED = "LimitExceeded.CustomLogFieldRegexLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) CreateRealtimeLogDeliveryTask(request *CreateRealtimeLogDeliveryTaskRequest) (response *CreateRealtimeLogDeliveryTaskResponse, err error) {
    return c.CreateRealtimeLogDeliveryTaskWithContext(context.Background(), request)
}

// CreateRealtimeLogDeliveryTask
// This API is used to create a real-time log delivery task.
//
// 
//
// **Restrictions**
//
// 
//
// An entity (a Layer 7 domain name, a Layer 4 proxy instance, or an Edge Function instance) under the same combination of `LogType` and `Area` can only have one task created for each `TaskType`. The mutual exclusion rules for different `LogType` values are as follows:
//
// 
//
// | LogType | Conflicting TaskType |
//
// |---|---|
//
// | Layer 4 proxy logs, Edge Function execution logs | Choose one of `custom_endpoint`, `s3`, and `aws_s3` |
//
// 
//
// When `TaskType` is `log_analysis`:
//
// 
//
// <ul>
//
// <li>Only `LogType` = `l7-access-logs` or `web-attack` is supported.</li>
//
// <li>Under the same combination of `ZoneId` and `Area`, only one task can be created for each `LogType`.</li>
//
// </ul>
//
// 
//
// **Recommendation**: Before creating a task, call the [DescribeRealtimeLogDeliveryTasks](https://www.tencentcloud.com/document/product/1145/59503) API to query existing tasks and avoid duplicate additions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATECLSLOGSETFAILED = "FailedOperation.CreateClsLogSetFailed"
//  FAILEDOPERATION_CREATECLSLOGTOPICTASKFAILED = "FailedOperation.CreateClsLogTopicTaskFailed"
//  FAILEDOPERATION_CREATELOGTOPICTASKAUTHFAILURE = "FailedOperation.CreateLogTopicTaskAuthFailure"
//  FAILEDOPERATION_REALTIMELOGAUTHFAILURE = "FailedOperation.RealtimeLogAuthFailure"
//  FAILEDOPERATION_REALTIMELOGNOTFOUND = "FailedOperation.RealtimeLogNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDLOGFORMATFIELDDELIMITER = "InvalidParameter.InvalidLogFormatFieldDelimiter"
//  INVALIDPARAMETER_INVALIDLOGFORMATFORMATTYPE = "InvalidParameter.InvalidLogFormatFormatType"
//  INVALIDPARAMETER_INVALIDLOGFORMATRECORDDELIMITER = "InvalidParameter.InvalidLogFormatRecordDelimiter"
//  INVALIDPARAMETER_REALTIMELOGENTITYALREADYCREATED = "InvalidParameter.RealtimeLogEntityAlreadyCreated"
//  INVALIDPARAMETER_REALTIMELOGINVALIDDELIVERYAREA = "InvalidParameter.RealtimeLogInvalidDeliveryArea"
//  INVALIDPARAMETER_REALTIMELOGINVALIDLOGTYPE = "InvalidParameter.RealtimeLogInvalidLogType"
//  INVALIDPARAMETER_REALTIMELOGINVALIDTASKTYPE = "InvalidParameter.RealtimeLogInvalidTaskType"
//  INVALIDPARAMETER_REALTIMELOGNUMSEXCEEDLIMIT = "InvalidParameter.RealtimeLogNumsExceedLimit"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CUSTOMLOGFIELDREGEXLIMITEXCEEDED = "LimitExceeded.CustomLogFieldRegexLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) CreateRealtimeLogDeliveryTaskWithContext(ctx context.Context, request *CreateRealtimeLogDeliveryTaskRequest) (response *CreateRealtimeLogDeliveryTaskResponse, err error) {
    if request == nil {
        request = NewCreateRealtimeLogDeliveryTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreateRealtimeLogDeliveryTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRealtimeLogDeliveryTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRealtimeLogDeliveryTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRuleRequest() (request *CreateRuleRequest) {
    request = &CreateRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateRule")
    
    
    return
}

func NewCreateRuleResponse() (response *CreateRuleResponse) {
    response = &CreateRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRule
// This API is on an earlier version for creating rule engine APIs. EdgeOne comprehensively upgraded the rule engine APIs on January 21, 2025. For details on the new API for creating layer-7 acceleration rules, refer to [CreateL7AccRules](https://www.tencentcloud.com/document/product/1552/115822?from_cn_redirect=1).
//
// <p style="color: red;">Note: Starting from January 21, 2025, the old version APIs will stop update and iteration. New features will only be provided by the latest version interface. The existing capacity supported by the old version APIs will not be affected. To avoid data field conflicts when using the old version APIs, migrate to the new version rule engine APIs as soon as possible.</p>
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACCESSREDIRECTREGEXERROR = "InvalidParameter.AccessRedirectRegexError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_ERRACTIONUNSUPPORTTARGET = "InvalidParameter.ErrActionUnsupportTarget"
//  INVALIDPARAMETER_ERRINVALIDACTION = "InvalidParameter.ErrInvalidAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONORIGINPRIVATEADDRESS = "InvalidParameter.ErrInvalidActionOriginPrivateAddress"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMACTION = "InvalidParameter.ErrInvalidActionParamAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMBADVALUETYPE = "InvalidParameter.ErrInvalidActionParamBadValueType"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMDUPLICATENAME = "InvalidParameter.ErrInvalidActionParamDuplicateName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMNAME = "InvalidParameter.ErrInvalidActionParamName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_ERRINVALIDACTIONTYPE = "InvalidParameter.ErrInvalidActionType"
//  INVALIDPARAMETER_ERRINVALIDCONDITION = "InvalidParameter.ErrInvalidCondition"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONCANNOTONLYCONTAINHOSTWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidConditionCannotOnlyContainHostWhenModifyOriginActionConfigured"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONIGNORECASE = "InvalidParameter.ErrInvalidConditionIgnoreCase"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMEBADNAME = "InvalidParameter.ErrInvalidConditionNameBadName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMETARGETNOTSUPPORTNAME = "InvalidParameter.ErrInvalidConditionNameTargetNotSupportName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADREGULAR = "InvalidParameter.ErrInvalidConditionValueBadRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADURL = "InvalidParameter.ErrInvalidConditionValueBadUrl"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUECONTAINFILENAMEEXTENSION = "InvalidParameter.ErrInvalidConditionValueBadValueContainFileNameExtension"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOLONGVALUE = "InvalidParameter.ErrInvalidConditionValueTooLongValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYREGULAR = "InvalidParameter.ErrInvalidConditionValueTooManyRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYVALUES = "InvalidParameter.ErrInvalidConditionValueTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYWILDCARD = "InvalidParameter.ErrInvalidConditionValueTooManyWildcard"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEZEROLENGTH = "InvalidParameter.ErrInvalidConditionValueZeroLength"
//  INVALIDPARAMETER_ERRNILCONDITION = "InvalidParameter.ErrNilCondition"
//  INVALIDPARAMETER_GRPCREQUIREHTTP2 = "InvalidParameter.GrpcRequireHttp2"
//  INVALIDPARAMETER_HOSTNOTFOUND = "InvalidParameter.HostNotFound"
//  INVALIDPARAMETER_INVALIDAUTHENTICATION = "InvalidParameter.InvalidAuthentication"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPEEXPIRETIME = "InvalidParameter.InvalidAuthenticationTypeExpireTime"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESIGNPARAM = "InvalidParameter.InvalidAuthenticationTypeSignParam"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDBACKUPSERVERNAME = "InvalidParameter.InvalidBackupServerName"
//  INVALIDPARAMETER_INVALIDCACHECONFIGCACHE = "InvalidParameter.InvalidCacheConfigCache"
//  INVALIDPARAMETER_INVALIDCACHEKEY = "InvalidParameter.InvalidCacheKey"
//  INVALIDPARAMETER_INVALIDCACHEKEYQUERYSTRINGVALUE = "InvalidParameter.InvalidCacheKeyQueryStringValue"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINE = "InvalidParameter.InvalidDynamicRoutine"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINEBILLING = "InvalidParameter.InvalidDynamicRoutineBilling"
//  INVALIDPARAMETER_INVALIDERRORPAGEREDIRECTURL = "InvalidParameter.InvalidErrorPageRedirectUrl"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGETIME = "InvalidParameter.InvalidMaxAgeTime"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPOSTSIZEVALUE = "InvalidParameter.InvalidPostSizeValue"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAMEXFF = "InvalidParameter.InvalidRequestHeaderNameXff"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERVALUE = "InvalidParameter.InvalidRequestHeaderValue"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERNAME = "InvalidParameter.InvalidResponseHeaderName"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERVALUE = "InvalidParameter.InvalidResponseHeaderValue"
//  INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSEXTENSION = "InvalidParameter.InvalidRuleEngineTargetsExtension"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSURL = "InvalidParameter.InvalidRuleEngineTargetsUrl"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGEXPIRETIMELIMIT = "InvalidParameter.InvalidStandardDebugExpireTimeLimit"
//  INVALIDPARAMETER_INVALIDUPSTREAMREQUESTQUERYSTRINGVALUE = "InvalidParameter.InvalidUpstreamRequestQueryStringValue"
//  INVALIDPARAMETER_INVALIDURLREDIRECT = "InvalidParameter.InvalidUrlRedirect"
//  INVALIDPARAMETER_INVALIDURLREDIRECTHOST = "InvalidParameter.InvalidUrlRedirectHost"
//  INVALIDPARAMETER_INVALIDURLREDIRECTURL = "InvalidParameter.InvalidUrlRedirectUrl"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_KEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.KeyRulesInvalidQueryStringValue"
//  INVALIDPARAMETER_LOADBALANCEINSTANCEIDISREQUIRED = "InvalidParameter.LoadBalanceInstanceIdIsRequired"
//  INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"
//  INVALIDPARAMETER_ORIGINORIGINGROUPIDISREQUIRED = "InvalidParameter.OriginOriginGroupIdIsRequired"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_TASKSYSTEMERROR = "InvalidParameter.TaskSystemError"
//  INVALIDPARAMETERVALUE_FORMATMISMATCH = "InvalidParameterValue.FormatMismatch"
//  INVALIDPARAMETERVALUE_GENERALMISMATCH = "InvalidParameterValue.GeneralMismatch"
//  INVALIDPARAMETERVALUE_INCLUDEINVALIDVALUE = "InvalidParameterValue.IncludeInvalidValue"
//  INVALIDPARAMETERVALUE_MISSINGNECESSARYPARAM = "InvalidParameterValue.MissingNecessaryParam"
//  INVALIDPARAMETERVALUE_NOTINENUMERATION = "InvalidParameterValue.NotInEnumeration"
//  INVALIDPARAMETERVALUE_NOTWITHINRANGE = "InvalidParameterValue.NotWithinRange"
//  INVALIDPARAMETERVALUE_REGEXMISMATCH = "InvalidParameterValue.RegExMismatch"
//  INVALIDPARAMETERVALUE_UNRECOGNIZABLEVALUE = "InvalidParameterValue.UnrecognizableValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDMULTIPLYLAYERCONFLICT = "OperationDenied.AccelerateMainlandMultiplyLayerConflict"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_INVALIDADVANCEDDEFENSESECURITYTYPE = "OperationDenied.InvalidAdvancedDefenseSecurityType"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateRule(request *CreateRuleRequest) (response *CreateRuleResponse, err error) {
    return c.CreateRuleWithContext(context.Background(), request)
}

// CreateRule
// This API is on an earlier version for creating rule engine APIs. EdgeOne comprehensively upgraded the rule engine APIs on January 21, 2025. For details on the new API for creating layer-7 acceleration rules, refer to [CreateL7AccRules](https://www.tencentcloud.com/document/product/1552/115822?from_cn_redirect=1).
//
// <p style="color: red;">Note: Starting from January 21, 2025, the old version APIs will stop update and iteration. New features will only be provided by the latest version interface. The existing capacity supported by the old version APIs will not be affected. To avoid data field conflicts when using the old version APIs, migrate to the new version rule engine APIs as soon as possible.</p>
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACCESSREDIRECTREGEXERROR = "InvalidParameter.AccessRedirectRegexError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_ERRACTIONUNSUPPORTTARGET = "InvalidParameter.ErrActionUnsupportTarget"
//  INVALIDPARAMETER_ERRINVALIDACTION = "InvalidParameter.ErrInvalidAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONORIGINPRIVATEADDRESS = "InvalidParameter.ErrInvalidActionOriginPrivateAddress"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMACTION = "InvalidParameter.ErrInvalidActionParamAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMBADVALUETYPE = "InvalidParameter.ErrInvalidActionParamBadValueType"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMDUPLICATENAME = "InvalidParameter.ErrInvalidActionParamDuplicateName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMNAME = "InvalidParameter.ErrInvalidActionParamName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_ERRINVALIDACTIONTYPE = "InvalidParameter.ErrInvalidActionType"
//  INVALIDPARAMETER_ERRINVALIDCONDITION = "InvalidParameter.ErrInvalidCondition"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONCANNOTONLYCONTAINHOSTWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidConditionCannotOnlyContainHostWhenModifyOriginActionConfigured"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONIGNORECASE = "InvalidParameter.ErrInvalidConditionIgnoreCase"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMEBADNAME = "InvalidParameter.ErrInvalidConditionNameBadName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMETARGETNOTSUPPORTNAME = "InvalidParameter.ErrInvalidConditionNameTargetNotSupportName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADREGULAR = "InvalidParameter.ErrInvalidConditionValueBadRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADURL = "InvalidParameter.ErrInvalidConditionValueBadUrl"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUECONTAINFILENAMEEXTENSION = "InvalidParameter.ErrInvalidConditionValueBadValueContainFileNameExtension"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOLONGVALUE = "InvalidParameter.ErrInvalidConditionValueTooLongValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYREGULAR = "InvalidParameter.ErrInvalidConditionValueTooManyRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYVALUES = "InvalidParameter.ErrInvalidConditionValueTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYWILDCARD = "InvalidParameter.ErrInvalidConditionValueTooManyWildcard"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEZEROLENGTH = "InvalidParameter.ErrInvalidConditionValueZeroLength"
//  INVALIDPARAMETER_ERRNILCONDITION = "InvalidParameter.ErrNilCondition"
//  INVALIDPARAMETER_GRPCREQUIREHTTP2 = "InvalidParameter.GrpcRequireHttp2"
//  INVALIDPARAMETER_HOSTNOTFOUND = "InvalidParameter.HostNotFound"
//  INVALIDPARAMETER_INVALIDAUTHENTICATION = "InvalidParameter.InvalidAuthentication"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPEEXPIRETIME = "InvalidParameter.InvalidAuthenticationTypeExpireTime"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESIGNPARAM = "InvalidParameter.InvalidAuthenticationTypeSignParam"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDBACKUPSERVERNAME = "InvalidParameter.InvalidBackupServerName"
//  INVALIDPARAMETER_INVALIDCACHECONFIGCACHE = "InvalidParameter.InvalidCacheConfigCache"
//  INVALIDPARAMETER_INVALIDCACHEKEY = "InvalidParameter.InvalidCacheKey"
//  INVALIDPARAMETER_INVALIDCACHEKEYQUERYSTRINGVALUE = "InvalidParameter.InvalidCacheKeyQueryStringValue"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINE = "InvalidParameter.InvalidDynamicRoutine"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINEBILLING = "InvalidParameter.InvalidDynamicRoutineBilling"
//  INVALIDPARAMETER_INVALIDERRORPAGEREDIRECTURL = "InvalidParameter.InvalidErrorPageRedirectUrl"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGETIME = "InvalidParameter.InvalidMaxAgeTime"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPOSTSIZEVALUE = "InvalidParameter.InvalidPostSizeValue"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAMEXFF = "InvalidParameter.InvalidRequestHeaderNameXff"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERVALUE = "InvalidParameter.InvalidRequestHeaderValue"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERNAME = "InvalidParameter.InvalidResponseHeaderName"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERVALUE = "InvalidParameter.InvalidResponseHeaderValue"
//  INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSEXTENSION = "InvalidParameter.InvalidRuleEngineTargetsExtension"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSURL = "InvalidParameter.InvalidRuleEngineTargetsUrl"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGEXPIRETIMELIMIT = "InvalidParameter.InvalidStandardDebugExpireTimeLimit"
//  INVALIDPARAMETER_INVALIDUPSTREAMREQUESTQUERYSTRINGVALUE = "InvalidParameter.InvalidUpstreamRequestQueryStringValue"
//  INVALIDPARAMETER_INVALIDURLREDIRECT = "InvalidParameter.InvalidUrlRedirect"
//  INVALIDPARAMETER_INVALIDURLREDIRECTHOST = "InvalidParameter.InvalidUrlRedirectHost"
//  INVALIDPARAMETER_INVALIDURLREDIRECTURL = "InvalidParameter.InvalidUrlRedirectUrl"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_KEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.KeyRulesInvalidQueryStringValue"
//  INVALIDPARAMETER_LOADBALANCEINSTANCEIDISREQUIRED = "InvalidParameter.LoadBalanceInstanceIdIsRequired"
//  INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"
//  INVALIDPARAMETER_ORIGINORIGINGROUPIDISREQUIRED = "InvalidParameter.OriginOriginGroupIdIsRequired"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_TASKSYSTEMERROR = "InvalidParameter.TaskSystemError"
//  INVALIDPARAMETERVALUE_FORMATMISMATCH = "InvalidParameterValue.FormatMismatch"
//  INVALIDPARAMETERVALUE_GENERALMISMATCH = "InvalidParameterValue.GeneralMismatch"
//  INVALIDPARAMETERVALUE_INCLUDEINVALIDVALUE = "InvalidParameterValue.IncludeInvalidValue"
//  INVALIDPARAMETERVALUE_MISSINGNECESSARYPARAM = "InvalidParameterValue.MissingNecessaryParam"
//  INVALIDPARAMETERVALUE_NOTINENUMERATION = "InvalidParameterValue.NotInEnumeration"
//  INVALIDPARAMETERVALUE_NOTWITHINRANGE = "InvalidParameterValue.NotWithinRange"
//  INVALIDPARAMETERVALUE_REGEXMISMATCH = "InvalidParameterValue.RegExMismatch"
//  INVALIDPARAMETERVALUE_UNRECOGNIZABLEVALUE = "InvalidParameterValue.UnrecognizableValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDMULTIPLYLAYERCONFLICT = "OperationDenied.AccelerateMainlandMultiplyLayerConflict"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_INVALIDADVANCEDDEFENSESECURITYTYPE = "OperationDenied.InvalidAdvancedDefenseSecurityType"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateRuleWithContext(ctx context.Context, request *CreateRuleRequest) (response *CreateRuleResponse, err error) {
    if request == nil {
        request = NewCreateRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreateRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSecurityAPIResourceRequest() (request *CreateSecurityAPIResourceRequest) {
    request = &CreateSecurityAPIResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateSecurityAPIResource")
    
    
    return
}

func NewCreateSecurityAPIResourceResponse() (response *CreateSecurityAPIResourceResponse) {
    response = &CreateSecurityAPIResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSecurityAPIResource
// Used to create API resources.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACCESSREDIRECTREGEXERROR = "InvalidParameter.AccessRedirectRegexError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_ERRACTIONUNSUPPORTTARGET = "InvalidParameter.ErrActionUnsupportTarget"
//  INVALIDPARAMETER_ERRINVALIDACTION = "InvalidParameter.ErrInvalidAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONORIGINPRIVATEADDRESS = "InvalidParameter.ErrInvalidActionOriginPrivateAddress"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMACTION = "InvalidParameter.ErrInvalidActionParamAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMBADVALUETYPE = "InvalidParameter.ErrInvalidActionParamBadValueType"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMDUPLICATENAME = "InvalidParameter.ErrInvalidActionParamDuplicateName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMNAME = "InvalidParameter.ErrInvalidActionParamName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_ERRINVALIDACTIONTYPE = "InvalidParameter.ErrInvalidActionType"
//  INVALIDPARAMETER_ERRINVALIDCONDITION = "InvalidParameter.ErrInvalidCondition"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONCANNOTONLYCONTAINHOSTWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidConditionCannotOnlyContainHostWhenModifyOriginActionConfigured"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONIGNORECASE = "InvalidParameter.ErrInvalidConditionIgnoreCase"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMEBADNAME = "InvalidParameter.ErrInvalidConditionNameBadName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMETARGETNOTSUPPORTNAME = "InvalidParameter.ErrInvalidConditionNameTargetNotSupportName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADREGULAR = "InvalidParameter.ErrInvalidConditionValueBadRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADURL = "InvalidParameter.ErrInvalidConditionValueBadUrl"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUECONTAINFILENAMEEXTENSION = "InvalidParameter.ErrInvalidConditionValueBadValueContainFileNameExtension"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOLONGVALUE = "InvalidParameter.ErrInvalidConditionValueTooLongValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYREGULAR = "InvalidParameter.ErrInvalidConditionValueTooManyRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYVALUES = "InvalidParameter.ErrInvalidConditionValueTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYWILDCARD = "InvalidParameter.ErrInvalidConditionValueTooManyWildcard"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEZEROLENGTH = "InvalidParameter.ErrInvalidConditionValueZeroLength"
//  INVALIDPARAMETER_ERRNILCONDITION = "InvalidParameter.ErrNilCondition"
//  INVALIDPARAMETER_GRPCREQUIREHTTP2 = "InvalidParameter.GrpcRequireHttp2"
//  INVALIDPARAMETER_HOSTNOTFOUND = "InvalidParameter.HostNotFound"
//  INVALIDPARAMETER_INVALIDAUTHENTICATION = "InvalidParameter.InvalidAuthentication"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPEEXPIRETIME = "InvalidParameter.InvalidAuthenticationTypeExpireTime"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESIGNPARAM = "InvalidParameter.InvalidAuthenticationTypeSignParam"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDBACKUPSERVERNAME = "InvalidParameter.InvalidBackupServerName"
//  INVALIDPARAMETER_INVALIDCACHECONFIGCACHE = "InvalidParameter.InvalidCacheConfigCache"
//  INVALIDPARAMETER_INVALIDCACHEKEY = "InvalidParameter.InvalidCacheKey"
//  INVALIDPARAMETER_INVALIDCACHEKEYQUERYSTRINGVALUE = "InvalidParameter.InvalidCacheKeyQueryStringValue"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINE = "InvalidParameter.InvalidDynamicRoutine"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINEBILLING = "InvalidParameter.InvalidDynamicRoutineBilling"
//  INVALIDPARAMETER_INVALIDERRORPAGEREDIRECTURL = "InvalidParameter.InvalidErrorPageRedirectUrl"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGETIME = "InvalidParameter.InvalidMaxAgeTime"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPOSTSIZEVALUE = "InvalidParameter.InvalidPostSizeValue"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAMEXFF = "InvalidParameter.InvalidRequestHeaderNameXff"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERVALUE = "InvalidParameter.InvalidRequestHeaderValue"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERNAME = "InvalidParameter.InvalidResponseHeaderName"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERVALUE = "InvalidParameter.InvalidResponseHeaderValue"
//  INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSEXTENSION = "InvalidParameter.InvalidRuleEngineTargetsExtension"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSURL = "InvalidParameter.InvalidRuleEngineTargetsUrl"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGEXPIRETIMELIMIT = "InvalidParameter.InvalidStandardDebugExpireTimeLimit"
//  INVALIDPARAMETER_INVALIDUPSTREAMREQUESTQUERYSTRINGVALUE = "InvalidParameter.InvalidUpstreamRequestQueryStringValue"
//  INVALIDPARAMETER_INVALIDURLREDIRECT = "InvalidParameter.InvalidUrlRedirect"
//  INVALIDPARAMETER_INVALIDURLREDIRECTHOST = "InvalidParameter.InvalidUrlRedirectHost"
//  INVALIDPARAMETER_INVALIDURLREDIRECTURL = "InvalidParameter.InvalidUrlRedirectUrl"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_KEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.KeyRulesInvalidQueryStringValue"
//  INVALIDPARAMETER_LOADBALANCEINSTANCEIDISREQUIRED = "InvalidParameter.LoadBalanceInstanceIdIsRequired"
//  INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"
//  INVALIDPARAMETER_ORIGINORIGINGROUPIDISREQUIRED = "InvalidParameter.OriginOriginGroupIdIsRequired"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_TASKSYSTEMERROR = "InvalidParameter.TaskSystemError"
//  INVALIDPARAMETERVALUE_FORMATMISMATCH = "InvalidParameterValue.FormatMismatch"
//  INVALIDPARAMETERVALUE_GENERALMISMATCH = "InvalidParameterValue.GeneralMismatch"
//  INVALIDPARAMETERVALUE_INCLUDEINVALIDVALUE = "InvalidParameterValue.IncludeInvalidValue"
//  INVALIDPARAMETERVALUE_MISSINGNECESSARYPARAM = "InvalidParameterValue.MissingNecessaryParam"
//  INVALIDPARAMETERVALUE_NOTINENUMERATION = "InvalidParameterValue.NotInEnumeration"
//  INVALIDPARAMETERVALUE_NOTWITHINRANGE = "InvalidParameterValue.NotWithinRange"
//  INVALIDPARAMETERVALUE_REGEXMISMATCH = "InvalidParameterValue.RegExMismatch"
//  INVALIDPARAMETERVALUE_UNRECOGNIZABLEVALUE = "InvalidParameterValue.UnrecognizableValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDMULTIPLYLAYERCONFLICT = "OperationDenied.AccelerateMainlandMultiplyLayerConflict"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_INVALIDADVANCEDDEFENSESECURITYTYPE = "OperationDenied.InvalidAdvancedDefenseSecurityType"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateSecurityAPIResource(request *CreateSecurityAPIResourceRequest) (response *CreateSecurityAPIResourceResponse, err error) {
    return c.CreateSecurityAPIResourceWithContext(context.Background(), request)
}

// CreateSecurityAPIResource
// Used to create API resources.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACCESSREDIRECTREGEXERROR = "InvalidParameter.AccessRedirectRegexError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_ERRACTIONUNSUPPORTTARGET = "InvalidParameter.ErrActionUnsupportTarget"
//  INVALIDPARAMETER_ERRINVALIDACTION = "InvalidParameter.ErrInvalidAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONORIGINPRIVATEADDRESS = "InvalidParameter.ErrInvalidActionOriginPrivateAddress"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMACTION = "InvalidParameter.ErrInvalidActionParamAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMBADVALUETYPE = "InvalidParameter.ErrInvalidActionParamBadValueType"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMDUPLICATENAME = "InvalidParameter.ErrInvalidActionParamDuplicateName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMNAME = "InvalidParameter.ErrInvalidActionParamName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_ERRINVALIDACTIONTYPE = "InvalidParameter.ErrInvalidActionType"
//  INVALIDPARAMETER_ERRINVALIDCONDITION = "InvalidParameter.ErrInvalidCondition"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONCANNOTONLYCONTAINHOSTWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidConditionCannotOnlyContainHostWhenModifyOriginActionConfigured"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONIGNORECASE = "InvalidParameter.ErrInvalidConditionIgnoreCase"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMEBADNAME = "InvalidParameter.ErrInvalidConditionNameBadName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMETARGETNOTSUPPORTNAME = "InvalidParameter.ErrInvalidConditionNameTargetNotSupportName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADREGULAR = "InvalidParameter.ErrInvalidConditionValueBadRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADURL = "InvalidParameter.ErrInvalidConditionValueBadUrl"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUECONTAINFILENAMEEXTENSION = "InvalidParameter.ErrInvalidConditionValueBadValueContainFileNameExtension"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOLONGVALUE = "InvalidParameter.ErrInvalidConditionValueTooLongValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYREGULAR = "InvalidParameter.ErrInvalidConditionValueTooManyRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYVALUES = "InvalidParameter.ErrInvalidConditionValueTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYWILDCARD = "InvalidParameter.ErrInvalidConditionValueTooManyWildcard"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEZEROLENGTH = "InvalidParameter.ErrInvalidConditionValueZeroLength"
//  INVALIDPARAMETER_ERRNILCONDITION = "InvalidParameter.ErrNilCondition"
//  INVALIDPARAMETER_GRPCREQUIREHTTP2 = "InvalidParameter.GrpcRequireHttp2"
//  INVALIDPARAMETER_HOSTNOTFOUND = "InvalidParameter.HostNotFound"
//  INVALIDPARAMETER_INVALIDAUTHENTICATION = "InvalidParameter.InvalidAuthentication"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPEEXPIRETIME = "InvalidParameter.InvalidAuthenticationTypeExpireTime"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESIGNPARAM = "InvalidParameter.InvalidAuthenticationTypeSignParam"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDBACKUPSERVERNAME = "InvalidParameter.InvalidBackupServerName"
//  INVALIDPARAMETER_INVALIDCACHECONFIGCACHE = "InvalidParameter.InvalidCacheConfigCache"
//  INVALIDPARAMETER_INVALIDCACHEKEY = "InvalidParameter.InvalidCacheKey"
//  INVALIDPARAMETER_INVALIDCACHEKEYQUERYSTRINGVALUE = "InvalidParameter.InvalidCacheKeyQueryStringValue"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINE = "InvalidParameter.InvalidDynamicRoutine"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINEBILLING = "InvalidParameter.InvalidDynamicRoutineBilling"
//  INVALIDPARAMETER_INVALIDERRORPAGEREDIRECTURL = "InvalidParameter.InvalidErrorPageRedirectUrl"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGETIME = "InvalidParameter.InvalidMaxAgeTime"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPOSTSIZEVALUE = "InvalidParameter.InvalidPostSizeValue"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAMEXFF = "InvalidParameter.InvalidRequestHeaderNameXff"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERVALUE = "InvalidParameter.InvalidRequestHeaderValue"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERNAME = "InvalidParameter.InvalidResponseHeaderName"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERVALUE = "InvalidParameter.InvalidResponseHeaderValue"
//  INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSEXTENSION = "InvalidParameter.InvalidRuleEngineTargetsExtension"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSURL = "InvalidParameter.InvalidRuleEngineTargetsUrl"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGEXPIRETIMELIMIT = "InvalidParameter.InvalidStandardDebugExpireTimeLimit"
//  INVALIDPARAMETER_INVALIDUPSTREAMREQUESTQUERYSTRINGVALUE = "InvalidParameter.InvalidUpstreamRequestQueryStringValue"
//  INVALIDPARAMETER_INVALIDURLREDIRECT = "InvalidParameter.InvalidUrlRedirect"
//  INVALIDPARAMETER_INVALIDURLREDIRECTHOST = "InvalidParameter.InvalidUrlRedirectHost"
//  INVALIDPARAMETER_INVALIDURLREDIRECTURL = "InvalidParameter.InvalidUrlRedirectUrl"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_KEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.KeyRulesInvalidQueryStringValue"
//  INVALIDPARAMETER_LOADBALANCEINSTANCEIDISREQUIRED = "InvalidParameter.LoadBalanceInstanceIdIsRequired"
//  INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"
//  INVALIDPARAMETER_ORIGINORIGINGROUPIDISREQUIRED = "InvalidParameter.OriginOriginGroupIdIsRequired"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_TASKSYSTEMERROR = "InvalidParameter.TaskSystemError"
//  INVALIDPARAMETERVALUE_FORMATMISMATCH = "InvalidParameterValue.FormatMismatch"
//  INVALIDPARAMETERVALUE_GENERALMISMATCH = "InvalidParameterValue.GeneralMismatch"
//  INVALIDPARAMETERVALUE_INCLUDEINVALIDVALUE = "InvalidParameterValue.IncludeInvalidValue"
//  INVALIDPARAMETERVALUE_MISSINGNECESSARYPARAM = "InvalidParameterValue.MissingNecessaryParam"
//  INVALIDPARAMETERVALUE_NOTINENUMERATION = "InvalidParameterValue.NotInEnumeration"
//  INVALIDPARAMETERVALUE_NOTWITHINRANGE = "InvalidParameterValue.NotWithinRange"
//  INVALIDPARAMETERVALUE_REGEXMISMATCH = "InvalidParameterValue.RegExMismatch"
//  INVALIDPARAMETERVALUE_UNRECOGNIZABLEVALUE = "InvalidParameterValue.UnrecognizableValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDMULTIPLYLAYERCONFLICT = "OperationDenied.AccelerateMainlandMultiplyLayerConflict"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_INVALIDADVANCEDDEFENSESECURITYTYPE = "OperationDenied.InvalidAdvancedDefenseSecurityType"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateSecurityAPIResourceWithContext(ctx context.Context, request *CreateSecurityAPIResourceRequest) (response *CreateSecurityAPIResourceResponse, err error) {
    if request == nil {
        request = NewCreateSecurityAPIResourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreateSecurityAPIResource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSecurityAPIResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSecurityAPIResourceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSecurityAPIServiceRequest() (request *CreateSecurityAPIServiceRequest) {
    request = &CreateSecurityAPIServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateSecurityAPIService")
    
    
    return
}

func NewCreateSecurityAPIServiceResponse() (response *CreateSecurityAPIServiceResponse) {
    response = &CreateSecurityAPIServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSecurityAPIService
// Used to create an API service.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACCESSREDIRECTREGEXERROR = "InvalidParameter.AccessRedirectRegexError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_ERRACTIONUNSUPPORTTARGET = "InvalidParameter.ErrActionUnsupportTarget"
//  INVALIDPARAMETER_ERRINVALIDACTION = "InvalidParameter.ErrInvalidAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONORIGINPRIVATEADDRESS = "InvalidParameter.ErrInvalidActionOriginPrivateAddress"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMACTION = "InvalidParameter.ErrInvalidActionParamAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMBADVALUETYPE = "InvalidParameter.ErrInvalidActionParamBadValueType"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMDUPLICATENAME = "InvalidParameter.ErrInvalidActionParamDuplicateName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMNAME = "InvalidParameter.ErrInvalidActionParamName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_ERRINVALIDACTIONTYPE = "InvalidParameter.ErrInvalidActionType"
//  INVALIDPARAMETER_ERRINVALIDCONDITION = "InvalidParameter.ErrInvalidCondition"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONCANNOTONLYCONTAINHOSTWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidConditionCannotOnlyContainHostWhenModifyOriginActionConfigured"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONIGNORECASE = "InvalidParameter.ErrInvalidConditionIgnoreCase"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMEBADNAME = "InvalidParameter.ErrInvalidConditionNameBadName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMETARGETNOTSUPPORTNAME = "InvalidParameter.ErrInvalidConditionNameTargetNotSupportName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADREGULAR = "InvalidParameter.ErrInvalidConditionValueBadRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADURL = "InvalidParameter.ErrInvalidConditionValueBadUrl"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUECONTAINFILENAMEEXTENSION = "InvalidParameter.ErrInvalidConditionValueBadValueContainFileNameExtension"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOLONGVALUE = "InvalidParameter.ErrInvalidConditionValueTooLongValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYREGULAR = "InvalidParameter.ErrInvalidConditionValueTooManyRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYVALUES = "InvalidParameter.ErrInvalidConditionValueTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYWILDCARD = "InvalidParameter.ErrInvalidConditionValueTooManyWildcard"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEZEROLENGTH = "InvalidParameter.ErrInvalidConditionValueZeroLength"
//  INVALIDPARAMETER_ERRNILCONDITION = "InvalidParameter.ErrNilCondition"
//  INVALIDPARAMETER_GRPCREQUIREHTTP2 = "InvalidParameter.GrpcRequireHttp2"
//  INVALIDPARAMETER_HOSTNOTFOUND = "InvalidParameter.HostNotFound"
//  INVALIDPARAMETER_INVALIDAUTHENTICATION = "InvalidParameter.InvalidAuthentication"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPEEXPIRETIME = "InvalidParameter.InvalidAuthenticationTypeExpireTime"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESIGNPARAM = "InvalidParameter.InvalidAuthenticationTypeSignParam"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDBACKUPSERVERNAME = "InvalidParameter.InvalidBackupServerName"
//  INVALIDPARAMETER_INVALIDCACHECONFIGCACHE = "InvalidParameter.InvalidCacheConfigCache"
//  INVALIDPARAMETER_INVALIDCACHEKEY = "InvalidParameter.InvalidCacheKey"
//  INVALIDPARAMETER_INVALIDCACHEKEYQUERYSTRINGVALUE = "InvalidParameter.InvalidCacheKeyQueryStringValue"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINE = "InvalidParameter.InvalidDynamicRoutine"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINEBILLING = "InvalidParameter.InvalidDynamicRoutineBilling"
//  INVALIDPARAMETER_INVALIDERRORPAGEREDIRECTURL = "InvalidParameter.InvalidErrorPageRedirectUrl"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGETIME = "InvalidParameter.InvalidMaxAgeTime"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPOSTSIZEVALUE = "InvalidParameter.InvalidPostSizeValue"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAMEXFF = "InvalidParameter.InvalidRequestHeaderNameXff"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERVALUE = "InvalidParameter.InvalidRequestHeaderValue"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERNAME = "InvalidParameter.InvalidResponseHeaderName"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERVALUE = "InvalidParameter.InvalidResponseHeaderValue"
//  INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSEXTENSION = "InvalidParameter.InvalidRuleEngineTargetsExtension"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSURL = "InvalidParameter.InvalidRuleEngineTargetsUrl"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGEXPIRETIMELIMIT = "InvalidParameter.InvalidStandardDebugExpireTimeLimit"
//  INVALIDPARAMETER_INVALIDUPSTREAMREQUESTQUERYSTRINGVALUE = "InvalidParameter.InvalidUpstreamRequestQueryStringValue"
//  INVALIDPARAMETER_INVALIDURLREDIRECT = "InvalidParameter.InvalidUrlRedirect"
//  INVALIDPARAMETER_INVALIDURLREDIRECTHOST = "InvalidParameter.InvalidUrlRedirectHost"
//  INVALIDPARAMETER_INVALIDURLREDIRECTURL = "InvalidParameter.InvalidUrlRedirectUrl"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_KEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.KeyRulesInvalidQueryStringValue"
//  INVALIDPARAMETER_LOADBALANCEINSTANCEIDISREQUIRED = "InvalidParameter.LoadBalanceInstanceIdIsRequired"
//  INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"
//  INVALIDPARAMETER_ORIGINORIGINGROUPIDISREQUIRED = "InvalidParameter.OriginOriginGroupIdIsRequired"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_TASKSYSTEMERROR = "InvalidParameter.TaskSystemError"
//  INVALIDPARAMETERVALUE_FORMATMISMATCH = "InvalidParameterValue.FormatMismatch"
//  INVALIDPARAMETERVALUE_GENERALMISMATCH = "InvalidParameterValue.GeneralMismatch"
//  INVALIDPARAMETERVALUE_INCLUDEINVALIDVALUE = "InvalidParameterValue.IncludeInvalidValue"
//  INVALIDPARAMETERVALUE_MISSINGNECESSARYPARAM = "InvalidParameterValue.MissingNecessaryParam"
//  INVALIDPARAMETERVALUE_NOTINENUMERATION = "InvalidParameterValue.NotInEnumeration"
//  INVALIDPARAMETERVALUE_NOTWITHINRANGE = "InvalidParameterValue.NotWithinRange"
//  INVALIDPARAMETERVALUE_REGEXMISMATCH = "InvalidParameterValue.RegExMismatch"
//  INVALIDPARAMETERVALUE_UNRECOGNIZABLEVALUE = "InvalidParameterValue.UnrecognizableValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDMULTIPLYLAYERCONFLICT = "OperationDenied.AccelerateMainlandMultiplyLayerConflict"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_INVALIDADVANCEDDEFENSESECURITYTYPE = "OperationDenied.InvalidAdvancedDefenseSecurityType"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateSecurityAPIService(request *CreateSecurityAPIServiceRequest) (response *CreateSecurityAPIServiceResponse, err error) {
    return c.CreateSecurityAPIServiceWithContext(context.Background(), request)
}

// CreateSecurityAPIService
// Used to create an API service.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACCESSREDIRECTREGEXERROR = "InvalidParameter.AccessRedirectRegexError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_ERRACTIONUNSUPPORTTARGET = "InvalidParameter.ErrActionUnsupportTarget"
//  INVALIDPARAMETER_ERRINVALIDACTION = "InvalidParameter.ErrInvalidAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONORIGINPRIVATEADDRESS = "InvalidParameter.ErrInvalidActionOriginPrivateAddress"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMACTION = "InvalidParameter.ErrInvalidActionParamAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMBADVALUETYPE = "InvalidParameter.ErrInvalidActionParamBadValueType"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMDUPLICATENAME = "InvalidParameter.ErrInvalidActionParamDuplicateName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMNAME = "InvalidParameter.ErrInvalidActionParamName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_ERRINVALIDACTIONTYPE = "InvalidParameter.ErrInvalidActionType"
//  INVALIDPARAMETER_ERRINVALIDCONDITION = "InvalidParameter.ErrInvalidCondition"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONCANNOTONLYCONTAINHOSTWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidConditionCannotOnlyContainHostWhenModifyOriginActionConfigured"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONIGNORECASE = "InvalidParameter.ErrInvalidConditionIgnoreCase"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMEBADNAME = "InvalidParameter.ErrInvalidConditionNameBadName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMETARGETNOTSUPPORTNAME = "InvalidParameter.ErrInvalidConditionNameTargetNotSupportName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADREGULAR = "InvalidParameter.ErrInvalidConditionValueBadRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADURL = "InvalidParameter.ErrInvalidConditionValueBadUrl"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUECONTAINFILENAMEEXTENSION = "InvalidParameter.ErrInvalidConditionValueBadValueContainFileNameExtension"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOLONGVALUE = "InvalidParameter.ErrInvalidConditionValueTooLongValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYREGULAR = "InvalidParameter.ErrInvalidConditionValueTooManyRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYVALUES = "InvalidParameter.ErrInvalidConditionValueTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYWILDCARD = "InvalidParameter.ErrInvalidConditionValueTooManyWildcard"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEZEROLENGTH = "InvalidParameter.ErrInvalidConditionValueZeroLength"
//  INVALIDPARAMETER_ERRNILCONDITION = "InvalidParameter.ErrNilCondition"
//  INVALIDPARAMETER_GRPCREQUIREHTTP2 = "InvalidParameter.GrpcRequireHttp2"
//  INVALIDPARAMETER_HOSTNOTFOUND = "InvalidParameter.HostNotFound"
//  INVALIDPARAMETER_INVALIDAUTHENTICATION = "InvalidParameter.InvalidAuthentication"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPEEXPIRETIME = "InvalidParameter.InvalidAuthenticationTypeExpireTime"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESIGNPARAM = "InvalidParameter.InvalidAuthenticationTypeSignParam"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDBACKUPSERVERNAME = "InvalidParameter.InvalidBackupServerName"
//  INVALIDPARAMETER_INVALIDCACHECONFIGCACHE = "InvalidParameter.InvalidCacheConfigCache"
//  INVALIDPARAMETER_INVALIDCACHEKEY = "InvalidParameter.InvalidCacheKey"
//  INVALIDPARAMETER_INVALIDCACHEKEYQUERYSTRINGVALUE = "InvalidParameter.InvalidCacheKeyQueryStringValue"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINE = "InvalidParameter.InvalidDynamicRoutine"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINEBILLING = "InvalidParameter.InvalidDynamicRoutineBilling"
//  INVALIDPARAMETER_INVALIDERRORPAGEREDIRECTURL = "InvalidParameter.InvalidErrorPageRedirectUrl"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGETIME = "InvalidParameter.InvalidMaxAgeTime"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPOSTSIZEVALUE = "InvalidParameter.InvalidPostSizeValue"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAMEXFF = "InvalidParameter.InvalidRequestHeaderNameXff"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERVALUE = "InvalidParameter.InvalidRequestHeaderValue"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERNAME = "InvalidParameter.InvalidResponseHeaderName"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERVALUE = "InvalidParameter.InvalidResponseHeaderValue"
//  INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSEXTENSION = "InvalidParameter.InvalidRuleEngineTargetsExtension"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSURL = "InvalidParameter.InvalidRuleEngineTargetsUrl"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGEXPIRETIMELIMIT = "InvalidParameter.InvalidStandardDebugExpireTimeLimit"
//  INVALIDPARAMETER_INVALIDUPSTREAMREQUESTQUERYSTRINGVALUE = "InvalidParameter.InvalidUpstreamRequestQueryStringValue"
//  INVALIDPARAMETER_INVALIDURLREDIRECT = "InvalidParameter.InvalidUrlRedirect"
//  INVALIDPARAMETER_INVALIDURLREDIRECTHOST = "InvalidParameter.InvalidUrlRedirectHost"
//  INVALIDPARAMETER_INVALIDURLREDIRECTURL = "InvalidParameter.InvalidUrlRedirectUrl"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_KEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.KeyRulesInvalidQueryStringValue"
//  INVALIDPARAMETER_LOADBALANCEINSTANCEIDISREQUIRED = "InvalidParameter.LoadBalanceInstanceIdIsRequired"
//  INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"
//  INVALIDPARAMETER_ORIGINORIGINGROUPIDISREQUIRED = "InvalidParameter.OriginOriginGroupIdIsRequired"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_TASKSYSTEMERROR = "InvalidParameter.TaskSystemError"
//  INVALIDPARAMETERVALUE_FORMATMISMATCH = "InvalidParameterValue.FormatMismatch"
//  INVALIDPARAMETERVALUE_GENERALMISMATCH = "InvalidParameterValue.GeneralMismatch"
//  INVALIDPARAMETERVALUE_INCLUDEINVALIDVALUE = "InvalidParameterValue.IncludeInvalidValue"
//  INVALIDPARAMETERVALUE_MISSINGNECESSARYPARAM = "InvalidParameterValue.MissingNecessaryParam"
//  INVALIDPARAMETERVALUE_NOTINENUMERATION = "InvalidParameterValue.NotInEnumeration"
//  INVALIDPARAMETERVALUE_NOTWITHINRANGE = "InvalidParameterValue.NotWithinRange"
//  INVALIDPARAMETERVALUE_REGEXMISMATCH = "InvalidParameterValue.RegExMismatch"
//  INVALIDPARAMETERVALUE_UNRECOGNIZABLEVALUE = "InvalidParameterValue.UnrecognizableValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDMULTIPLYLAYERCONFLICT = "OperationDenied.AccelerateMainlandMultiplyLayerConflict"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_INVALIDADVANCEDDEFENSESECURITYTYPE = "OperationDenied.InvalidAdvancedDefenseSecurityType"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateSecurityAPIServiceWithContext(ctx context.Context, request *CreateSecurityAPIServiceRequest) (response *CreateSecurityAPIServiceResponse, err error) {
    if request == nil {
        request = NewCreateSecurityAPIServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreateSecurityAPIService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSecurityAPIService require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSecurityAPIServiceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSecurityClientAttesterRequest() (request *CreateSecurityClientAttesterRequest) {
    request = &CreateSecurityClientAttesterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateSecurityClientAttester")
    
    
    return
}

func NewCreateSecurityClientAttesterResponse() (response *CreateSecurityClientAttesterResponse) {
    response = &CreateSecurityClientAttesterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSecurityClientAttester
// Create client authentication options.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACCESSREDIRECTREGEXERROR = "InvalidParameter.AccessRedirectRegexError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_ERRACTIONUNSUPPORTTARGET = "InvalidParameter.ErrActionUnsupportTarget"
//  INVALIDPARAMETER_ERRINVALIDACTION = "InvalidParameter.ErrInvalidAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONORIGINPRIVATEADDRESS = "InvalidParameter.ErrInvalidActionOriginPrivateAddress"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMACTION = "InvalidParameter.ErrInvalidActionParamAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMBADVALUETYPE = "InvalidParameter.ErrInvalidActionParamBadValueType"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMDUPLICATENAME = "InvalidParameter.ErrInvalidActionParamDuplicateName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMNAME = "InvalidParameter.ErrInvalidActionParamName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_ERRINVALIDACTIONTYPE = "InvalidParameter.ErrInvalidActionType"
//  INVALIDPARAMETER_ERRINVALIDCONDITION = "InvalidParameter.ErrInvalidCondition"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONCANNOTONLYCONTAINHOSTWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidConditionCannotOnlyContainHostWhenModifyOriginActionConfigured"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONIGNORECASE = "InvalidParameter.ErrInvalidConditionIgnoreCase"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMEBADNAME = "InvalidParameter.ErrInvalidConditionNameBadName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMETARGETNOTSUPPORTNAME = "InvalidParameter.ErrInvalidConditionNameTargetNotSupportName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADREGULAR = "InvalidParameter.ErrInvalidConditionValueBadRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADURL = "InvalidParameter.ErrInvalidConditionValueBadUrl"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUECONTAINFILENAMEEXTENSION = "InvalidParameter.ErrInvalidConditionValueBadValueContainFileNameExtension"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOLONGVALUE = "InvalidParameter.ErrInvalidConditionValueTooLongValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYREGULAR = "InvalidParameter.ErrInvalidConditionValueTooManyRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYVALUES = "InvalidParameter.ErrInvalidConditionValueTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYWILDCARD = "InvalidParameter.ErrInvalidConditionValueTooManyWildcard"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEZEROLENGTH = "InvalidParameter.ErrInvalidConditionValueZeroLength"
//  INVALIDPARAMETER_ERRNILCONDITION = "InvalidParameter.ErrNilCondition"
//  INVALIDPARAMETER_GRPCREQUIREHTTP2 = "InvalidParameter.GrpcRequireHttp2"
//  INVALIDPARAMETER_HOSTNOTFOUND = "InvalidParameter.HostNotFound"
//  INVALIDPARAMETER_INVALIDAUTHENTICATION = "InvalidParameter.InvalidAuthentication"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPEEXPIRETIME = "InvalidParameter.InvalidAuthenticationTypeExpireTime"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESIGNPARAM = "InvalidParameter.InvalidAuthenticationTypeSignParam"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDBACKUPSERVERNAME = "InvalidParameter.InvalidBackupServerName"
//  INVALIDPARAMETER_INVALIDCACHECONFIGCACHE = "InvalidParameter.InvalidCacheConfigCache"
//  INVALIDPARAMETER_INVALIDCACHEKEY = "InvalidParameter.InvalidCacheKey"
//  INVALIDPARAMETER_INVALIDCACHEKEYQUERYSTRINGVALUE = "InvalidParameter.InvalidCacheKeyQueryStringValue"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINE = "InvalidParameter.InvalidDynamicRoutine"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINEBILLING = "InvalidParameter.InvalidDynamicRoutineBilling"
//  INVALIDPARAMETER_INVALIDERRORPAGEREDIRECTURL = "InvalidParameter.InvalidErrorPageRedirectUrl"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGETIME = "InvalidParameter.InvalidMaxAgeTime"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPOSTSIZEVALUE = "InvalidParameter.InvalidPostSizeValue"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAMEXFF = "InvalidParameter.InvalidRequestHeaderNameXff"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERVALUE = "InvalidParameter.InvalidRequestHeaderValue"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERNAME = "InvalidParameter.InvalidResponseHeaderName"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERVALUE = "InvalidParameter.InvalidResponseHeaderValue"
//  INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSEXTENSION = "InvalidParameter.InvalidRuleEngineTargetsExtension"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSURL = "InvalidParameter.InvalidRuleEngineTargetsUrl"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGEXPIRETIMELIMIT = "InvalidParameter.InvalidStandardDebugExpireTimeLimit"
//  INVALIDPARAMETER_INVALIDUPSTREAMREQUESTQUERYSTRINGVALUE = "InvalidParameter.InvalidUpstreamRequestQueryStringValue"
//  INVALIDPARAMETER_INVALIDURLREDIRECT = "InvalidParameter.InvalidUrlRedirect"
//  INVALIDPARAMETER_INVALIDURLREDIRECTHOST = "InvalidParameter.InvalidUrlRedirectHost"
//  INVALIDPARAMETER_INVALIDURLREDIRECTURL = "InvalidParameter.InvalidUrlRedirectUrl"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_KEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.KeyRulesInvalidQueryStringValue"
//  INVALIDPARAMETER_LOADBALANCEINSTANCEIDISREQUIRED = "InvalidParameter.LoadBalanceInstanceIdIsRequired"
//  INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"
//  INVALIDPARAMETER_ORIGINORIGINGROUPIDISREQUIRED = "InvalidParameter.OriginOriginGroupIdIsRequired"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_TASKSYSTEMERROR = "InvalidParameter.TaskSystemError"
//  INVALIDPARAMETERVALUE_FORMATMISMATCH = "InvalidParameterValue.FormatMismatch"
//  INVALIDPARAMETERVALUE_GENERALMISMATCH = "InvalidParameterValue.GeneralMismatch"
//  INVALIDPARAMETERVALUE_INCLUDEINVALIDVALUE = "InvalidParameterValue.IncludeInvalidValue"
//  INVALIDPARAMETERVALUE_MISSINGNECESSARYPARAM = "InvalidParameterValue.MissingNecessaryParam"
//  INVALIDPARAMETERVALUE_NOTINENUMERATION = "InvalidParameterValue.NotInEnumeration"
//  INVALIDPARAMETERVALUE_NOTWITHINRANGE = "InvalidParameterValue.NotWithinRange"
//  INVALIDPARAMETERVALUE_REGEXMISMATCH = "InvalidParameterValue.RegExMismatch"
//  INVALIDPARAMETERVALUE_UNRECOGNIZABLEVALUE = "InvalidParameterValue.UnrecognizableValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDMULTIPLYLAYERCONFLICT = "OperationDenied.AccelerateMainlandMultiplyLayerConflict"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_INVALIDADVANCEDDEFENSESECURITYTYPE = "OperationDenied.InvalidAdvancedDefenseSecurityType"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateSecurityClientAttester(request *CreateSecurityClientAttesterRequest) (response *CreateSecurityClientAttesterResponse, err error) {
    return c.CreateSecurityClientAttesterWithContext(context.Background(), request)
}

// CreateSecurityClientAttester
// Create client authentication options.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACCESSREDIRECTREGEXERROR = "InvalidParameter.AccessRedirectRegexError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_ERRACTIONUNSUPPORTTARGET = "InvalidParameter.ErrActionUnsupportTarget"
//  INVALIDPARAMETER_ERRINVALIDACTION = "InvalidParameter.ErrInvalidAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONORIGINPRIVATEADDRESS = "InvalidParameter.ErrInvalidActionOriginPrivateAddress"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMACTION = "InvalidParameter.ErrInvalidActionParamAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMBADVALUETYPE = "InvalidParameter.ErrInvalidActionParamBadValueType"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMDUPLICATENAME = "InvalidParameter.ErrInvalidActionParamDuplicateName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMNAME = "InvalidParameter.ErrInvalidActionParamName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_ERRINVALIDACTIONTYPE = "InvalidParameter.ErrInvalidActionType"
//  INVALIDPARAMETER_ERRINVALIDCONDITION = "InvalidParameter.ErrInvalidCondition"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONCANNOTONLYCONTAINHOSTWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidConditionCannotOnlyContainHostWhenModifyOriginActionConfigured"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONIGNORECASE = "InvalidParameter.ErrInvalidConditionIgnoreCase"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMEBADNAME = "InvalidParameter.ErrInvalidConditionNameBadName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMETARGETNOTSUPPORTNAME = "InvalidParameter.ErrInvalidConditionNameTargetNotSupportName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADREGULAR = "InvalidParameter.ErrInvalidConditionValueBadRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADURL = "InvalidParameter.ErrInvalidConditionValueBadUrl"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUECONTAINFILENAMEEXTENSION = "InvalidParameter.ErrInvalidConditionValueBadValueContainFileNameExtension"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOLONGVALUE = "InvalidParameter.ErrInvalidConditionValueTooLongValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYREGULAR = "InvalidParameter.ErrInvalidConditionValueTooManyRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYVALUES = "InvalidParameter.ErrInvalidConditionValueTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYWILDCARD = "InvalidParameter.ErrInvalidConditionValueTooManyWildcard"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEZEROLENGTH = "InvalidParameter.ErrInvalidConditionValueZeroLength"
//  INVALIDPARAMETER_ERRNILCONDITION = "InvalidParameter.ErrNilCondition"
//  INVALIDPARAMETER_GRPCREQUIREHTTP2 = "InvalidParameter.GrpcRequireHttp2"
//  INVALIDPARAMETER_HOSTNOTFOUND = "InvalidParameter.HostNotFound"
//  INVALIDPARAMETER_INVALIDAUTHENTICATION = "InvalidParameter.InvalidAuthentication"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPEEXPIRETIME = "InvalidParameter.InvalidAuthenticationTypeExpireTime"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESIGNPARAM = "InvalidParameter.InvalidAuthenticationTypeSignParam"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDBACKUPSERVERNAME = "InvalidParameter.InvalidBackupServerName"
//  INVALIDPARAMETER_INVALIDCACHECONFIGCACHE = "InvalidParameter.InvalidCacheConfigCache"
//  INVALIDPARAMETER_INVALIDCACHEKEY = "InvalidParameter.InvalidCacheKey"
//  INVALIDPARAMETER_INVALIDCACHEKEYQUERYSTRINGVALUE = "InvalidParameter.InvalidCacheKeyQueryStringValue"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINE = "InvalidParameter.InvalidDynamicRoutine"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINEBILLING = "InvalidParameter.InvalidDynamicRoutineBilling"
//  INVALIDPARAMETER_INVALIDERRORPAGEREDIRECTURL = "InvalidParameter.InvalidErrorPageRedirectUrl"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGETIME = "InvalidParameter.InvalidMaxAgeTime"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPOSTSIZEVALUE = "InvalidParameter.InvalidPostSizeValue"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAMEXFF = "InvalidParameter.InvalidRequestHeaderNameXff"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERVALUE = "InvalidParameter.InvalidRequestHeaderValue"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERNAME = "InvalidParameter.InvalidResponseHeaderName"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERVALUE = "InvalidParameter.InvalidResponseHeaderValue"
//  INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSEXTENSION = "InvalidParameter.InvalidRuleEngineTargetsExtension"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSURL = "InvalidParameter.InvalidRuleEngineTargetsUrl"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGEXPIRETIMELIMIT = "InvalidParameter.InvalidStandardDebugExpireTimeLimit"
//  INVALIDPARAMETER_INVALIDUPSTREAMREQUESTQUERYSTRINGVALUE = "InvalidParameter.InvalidUpstreamRequestQueryStringValue"
//  INVALIDPARAMETER_INVALIDURLREDIRECT = "InvalidParameter.InvalidUrlRedirect"
//  INVALIDPARAMETER_INVALIDURLREDIRECTHOST = "InvalidParameter.InvalidUrlRedirectHost"
//  INVALIDPARAMETER_INVALIDURLREDIRECTURL = "InvalidParameter.InvalidUrlRedirectUrl"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_KEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.KeyRulesInvalidQueryStringValue"
//  INVALIDPARAMETER_LOADBALANCEINSTANCEIDISREQUIRED = "InvalidParameter.LoadBalanceInstanceIdIsRequired"
//  INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"
//  INVALIDPARAMETER_ORIGINORIGINGROUPIDISREQUIRED = "InvalidParameter.OriginOriginGroupIdIsRequired"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_TASKSYSTEMERROR = "InvalidParameter.TaskSystemError"
//  INVALIDPARAMETERVALUE_FORMATMISMATCH = "InvalidParameterValue.FormatMismatch"
//  INVALIDPARAMETERVALUE_GENERALMISMATCH = "InvalidParameterValue.GeneralMismatch"
//  INVALIDPARAMETERVALUE_INCLUDEINVALIDVALUE = "InvalidParameterValue.IncludeInvalidValue"
//  INVALIDPARAMETERVALUE_MISSINGNECESSARYPARAM = "InvalidParameterValue.MissingNecessaryParam"
//  INVALIDPARAMETERVALUE_NOTINENUMERATION = "InvalidParameterValue.NotInEnumeration"
//  INVALIDPARAMETERVALUE_NOTWITHINRANGE = "InvalidParameterValue.NotWithinRange"
//  INVALIDPARAMETERVALUE_REGEXMISMATCH = "InvalidParameterValue.RegExMismatch"
//  INVALIDPARAMETERVALUE_UNRECOGNIZABLEVALUE = "InvalidParameterValue.UnrecognizableValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDMULTIPLYLAYERCONFLICT = "OperationDenied.AccelerateMainlandMultiplyLayerConflict"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_INVALIDADVANCEDDEFENSESECURITYTYPE = "OperationDenied.InvalidAdvancedDefenseSecurityType"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateSecurityClientAttesterWithContext(ctx context.Context, request *CreateSecurityClientAttesterRequest) (response *CreateSecurityClientAttesterResponse, err error) {
    if request == nil {
        request = NewCreateSecurityClientAttesterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreateSecurityClientAttester")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSecurityClientAttester require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSecurityClientAttesterResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSecurityIPGroupRequest() (request *CreateSecurityIPGroupRequest) {
    request = &CreateSecurityIPGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateSecurityIPGroup")
    
    
    return
}

func NewCreateSecurityIPGroupResponse() (response *CreateSecurityIPGroupResponse) {
    response = &CreateSecurityIPGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSecurityIPGroup
// Create a security IP group
//
// error code that may be returned:
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateSecurityIPGroup(request *CreateSecurityIPGroupRequest) (response *CreateSecurityIPGroupResponse, err error) {
    return c.CreateSecurityIPGroupWithContext(context.Background(), request)
}

// CreateSecurityIPGroup
// Create a security IP group
//
// error code that may be returned:
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateSecurityIPGroupWithContext(ctx context.Context, request *CreateSecurityIPGroupRequest) (response *CreateSecurityIPGroupResponse, err error) {
    if request == nil {
        request = NewCreateSecurityIPGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreateSecurityIPGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSecurityIPGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSecurityIPGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSecurityJSInjectionRuleRequest() (request *CreateSecurityJSInjectionRuleRequest) {
    request = &CreateSecurityJSInjectionRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateSecurityJSInjectionRule")
    
    
    return
}

func NewCreateSecurityJSInjectionRuleResponse() (response *CreateSecurityJSInjectionRuleResponse) {
    response = &CreateSecurityJSInjectionRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSecurityJSInjectionRule
// This API is used to create JavaScript injection rules.
//
// error code that may be returned:
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateSecurityJSInjectionRule(request *CreateSecurityJSInjectionRuleRequest) (response *CreateSecurityJSInjectionRuleResponse, err error) {
    return c.CreateSecurityJSInjectionRuleWithContext(context.Background(), request)
}

// CreateSecurityJSInjectionRule
// This API is used to create JavaScript injection rules.
//
// error code that may be returned:
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateSecurityJSInjectionRuleWithContext(ctx context.Context, request *CreateSecurityJSInjectionRuleRequest) (response *CreateSecurityJSInjectionRuleResponse, err error) {
    if request == nil {
        request = NewCreateSecurityJSInjectionRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreateSecurityJSInjectionRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSecurityJSInjectionRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSecurityJSInjectionRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSharedCNAMERequest() (request *CreateSharedCNAMERequest) {
    request = &CreateSharedCNAMERequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateSharedCNAME")
    
    
    return
}

func NewCreateSharedCNAMEResponse() (response *CreateSharedCNAMEResponse) {
    response = &CreateSharedCNAMEResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSharedCNAME
// Used to create a shared CNAME. This feature is in beta test.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_NOTALLOWEDWILDCARDSHAREDCNAME = "InvalidParameterValue.NotAllowedWildcardSharedCNAME"
//  INVALIDPARAMETERVALUE_SHAREDCNAMEPREFIXNOTMATCH = "InvalidParameterValue.SharedCNAMEPrefixNotMatch"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  RESOURCEUNAVAILABLE_SHAREDCNAMEALREADYEXISTS = "ResourceUnavailable.SharedCNAMEAlreadyExists"
func (c *Client) CreateSharedCNAME(request *CreateSharedCNAMERequest) (response *CreateSharedCNAMEResponse, err error) {
    return c.CreateSharedCNAMEWithContext(context.Background(), request)
}

// CreateSharedCNAME
// Used to create a shared CNAME. This feature is in beta test.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_NOTALLOWEDWILDCARDSHAREDCNAME = "InvalidParameterValue.NotAllowedWildcardSharedCNAME"
//  INVALIDPARAMETERVALUE_SHAREDCNAMEPREFIXNOTMATCH = "InvalidParameterValue.SharedCNAMEPrefixNotMatch"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  RESOURCEUNAVAILABLE_SHAREDCNAMEALREADYEXISTS = "ResourceUnavailable.SharedCNAMEAlreadyExists"
func (c *Client) CreateSharedCNAMEWithContext(ctx context.Context, request *CreateSharedCNAMERequest) (response *CreateSharedCNAMEResponse, err error) {
    if request == nil {
        request = NewCreateSharedCNAMERequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreateSharedCNAME")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSharedCNAME require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSharedCNAMEResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWebSecurityTemplateRequest() (request *CreateWebSecurityTemplateRequest) {
    request = &CreateWebSecurityTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateWebSecurityTemplate")
    
    
    return
}

func NewCreateWebSecurityTemplateResponse() (response *CreateWebSecurityTemplateResponse) {
    response = &CreateWebSecurityTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateWebSecurityTemplate
// Create a security policy configuration template
//
// error code that may be returned:
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  LIMITEXCEEDED_SECURITY = "LimitExceeded.Security"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateWebSecurityTemplate(request *CreateWebSecurityTemplateRequest) (response *CreateWebSecurityTemplateResponse, err error) {
    return c.CreateWebSecurityTemplateWithContext(context.Background(), request)
}

// CreateWebSecurityTemplate
// Create a security policy configuration template
//
// error code that may be returned:
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  LIMITEXCEEDED_SECURITY = "LimitExceeded.Security"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateWebSecurityTemplateWithContext(ctx context.Context, request *CreateWebSecurityTemplateRequest) (response *CreateWebSecurityTemplateResponse, err error) {
    if request == nil {
        request = NewCreateWebSecurityTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreateWebSecurityTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWebSecurityTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWebSecurityTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateZoneRequest() (request *CreateZoneRequest) {
    request = &CreateZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateZone")
    
    
    return
}

func NewCreateZoneResponse() (response *CreateZoneResponse) {
    response = &CreateZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateZone
// EdgeOne offers three access methods: CNAME, NS, and domainless access. You need to finalize site creation through this API first. For CNAME and NS connected site scenarios, see [Quickly access EdgeOne from scratch](https://www.tencentcloud.com/document/product/1552/87601?from_cn_redirect=1). For domainless access scenarios, see [Quickly enable layer 4 proxy service](https://www.tencentcloud.com/document/product/1552/96051?from_cn_redirect=1).
//
// 
//
// > It is recommended to call this API to create a site when a package already exists under the account. Please input `PlanId` in the parameter to directly bind the site to the package. If `PlanId` is not input, the created site will be inactive and unable to provide normal services. You need to complete package binding through [BindZoneToPlan](https://www.tencentcloud.com/document/product/1552/83042?from_cn_redirect=1) before the site can provide services normally. If you currently have no bindable package, go to the console to purchase a package to finalize site creation.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION_FAILEDTOCALLDNSPOD = "FailedOperation.FailedToCallDNSPod"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_LENGTHEXCEEDSLIMIT = "InvalidParameter.LengthExceedsLimit"
//  INVALIDPARAMETER_PLANNOTFOUND = "InvalidParameter.PlanNotFound"
//  INVALIDPARAMETER_ZONEHASBEENBOUND = "InvalidParameter.ZoneHasBeenBound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACCESSBLACKLIST = "InvalidParameterValue.AccessBlacklist"
//  INVALIDPARAMETERVALUE_INVALIDTAGVALUE = "InvalidParameterValue.InvalidTagValue"
//  INVALIDPARAMETERVALUE_TOPLEVELDOMAINNOTSUPPORT = "InvalidParameterValue.TopLevelDomainNotSupport"
//  INVALIDPARAMETERVALUE_ZONENAMEINVALID = "InvalidParameterValue.ZoneNameInvalid"
//  INVALIDPARAMETERVALUE_ZONENAMENOTSUPPORTPUNYCODE = "InvalidParameterValue.ZoneNameNotSupportPunyCode"
//  INVALIDPARAMETERVALUE_ZONENAMENOTSUPPORTSUBDOMAIN = "InvalidParameterValue.ZoneNameNotSupportSubDomain"
//  INVALIDPARAMETERVALUE_ZONESAMEASNAME = "InvalidParameterValue.ZoneSameAsName"
//  LIMITEXCEEDED_ZONEBINDPLAN = "LimitExceeded.ZoneBindPlan"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_COMPLIANCEFORBIDDEN = "OperationDenied.ComplianceForbidden"
//  OPERATIONDENIED_DNSPODUNAUTHORIZEDROLEOPERATION = "OperationDenied.DNSPodUnauthorizedRoleOperation"
//  OPERATIONDENIED_DOMAINISBLOCKED = "OperationDenied.DomainIsBlocked"
//  OPERATIONDENIED_RECORDISFORBIDDEN = "OperationDenied.RecordIsForbidden"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_ALIASDOMAIN = "ResourceInUse.AliasDomain"
//  RESOURCEINUSE_CNAME = "ResourceInUse.Cname"
//  RESOURCEINUSE_DNS = "ResourceInUse.Dns"
//  RESOURCEINUSE_GENERICHOST = "ResourceInUse.GenericHost"
//  RESOURCEINUSE_HOST = "ResourceInUse.Host"
//  RESOURCEINUSE_NS = "ResourceInUse.NS"
//  RESOURCEINUSE_OTHERS = "ResourceInUse.Others"
//  RESOURCEINUSE_OTHERSALIASDOMAIN = "ResourceInUse.OthersAliasDomain"
//  RESOURCEINUSE_OTHERSCNAME = "ResourceInUse.OthersCname"
//  RESOURCEINUSE_OTHERSHOST = "ResourceInUse.OthersHost"
//  RESOURCEINUSE_OTHERSNS = "ResourceInUse.OthersNS"
//  RESOURCEINUSE_SELFANDOTHERSCNAME = "ResourceInUse.SelfAndOthersCname"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DNSPODDOMAINNOTINACCOUNT = "ResourceNotFound.DNSPodDomainNotInAccount"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) CreateZone(request *CreateZoneRequest) (response *CreateZoneResponse, err error) {
    return c.CreateZoneWithContext(context.Background(), request)
}

// CreateZone
// EdgeOne offers three access methods: CNAME, NS, and domainless access. You need to finalize site creation through this API first. For CNAME and NS connected site scenarios, see [Quickly access EdgeOne from scratch](https://www.tencentcloud.com/document/product/1552/87601?from_cn_redirect=1). For domainless access scenarios, see [Quickly enable layer 4 proxy service](https://www.tencentcloud.com/document/product/1552/96051?from_cn_redirect=1).
//
// 
//
// > It is recommended to call this API to create a site when a package already exists under the account. Please input `PlanId` in the parameter to directly bind the site to the package. If `PlanId` is not input, the created site will be inactive and unable to provide normal services. You need to complete package binding through [BindZoneToPlan](https://www.tencentcloud.com/document/product/1552/83042?from_cn_redirect=1) before the site can provide services normally. If you currently have no bindable package, go to the console to purchase a package to finalize site creation.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION_FAILEDTOCALLDNSPOD = "FailedOperation.FailedToCallDNSPod"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_LENGTHEXCEEDSLIMIT = "InvalidParameter.LengthExceedsLimit"
//  INVALIDPARAMETER_PLANNOTFOUND = "InvalidParameter.PlanNotFound"
//  INVALIDPARAMETER_ZONEHASBEENBOUND = "InvalidParameter.ZoneHasBeenBound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACCESSBLACKLIST = "InvalidParameterValue.AccessBlacklist"
//  INVALIDPARAMETERVALUE_INVALIDTAGVALUE = "InvalidParameterValue.InvalidTagValue"
//  INVALIDPARAMETERVALUE_TOPLEVELDOMAINNOTSUPPORT = "InvalidParameterValue.TopLevelDomainNotSupport"
//  INVALIDPARAMETERVALUE_ZONENAMEINVALID = "InvalidParameterValue.ZoneNameInvalid"
//  INVALIDPARAMETERVALUE_ZONENAMENOTSUPPORTPUNYCODE = "InvalidParameterValue.ZoneNameNotSupportPunyCode"
//  INVALIDPARAMETERVALUE_ZONENAMENOTSUPPORTSUBDOMAIN = "InvalidParameterValue.ZoneNameNotSupportSubDomain"
//  INVALIDPARAMETERVALUE_ZONESAMEASNAME = "InvalidParameterValue.ZoneSameAsName"
//  LIMITEXCEEDED_ZONEBINDPLAN = "LimitExceeded.ZoneBindPlan"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_COMPLIANCEFORBIDDEN = "OperationDenied.ComplianceForbidden"
//  OPERATIONDENIED_DNSPODUNAUTHORIZEDROLEOPERATION = "OperationDenied.DNSPodUnauthorizedRoleOperation"
//  OPERATIONDENIED_DOMAINISBLOCKED = "OperationDenied.DomainIsBlocked"
//  OPERATIONDENIED_RECORDISFORBIDDEN = "OperationDenied.RecordIsForbidden"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_ALIASDOMAIN = "ResourceInUse.AliasDomain"
//  RESOURCEINUSE_CNAME = "ResourceInUse.Cname"
//  RESOURCEINUSE_DNS = "ResourceInUse.Dns"
//  RESOURCEINUSE_GENERICHOST = "ResourceInUse.GenericHost"
//  RESOURCEINUSE_HOST = "ResourceInUse.Host"
//  RESOURCEINUSE_NS = "ResourceInUse.NS"
//  RESOURCEINUSE_OTHERS = "ResourceInUse.Others"
//  RESOURCEINUSE_OTHERSALIASDOMAIN = "ResourceInUse.OthersAliasDomain"
//  RESOURCEINUSE_OTHERSCNAME = "ResourceInUse.OthersCname"
//  RESOURCEINUSE_OTHERSHOST = "ResourceInUse.OthersHost"
//  RESOURCEINUSE_OTHERSNS = "ResourceInUse.OthersNS"
//  RESOURCEINUSE_SELFANDOTHERSCNAME = "ResourceInUse.SelfAndOthersCname"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DNSPODDOMAINNOTINACCOUNT = "ResourceNotFound.DNSPodDomainNotInAccount"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) CreateZoneWithContext(ctx context.Context, request *CreateZoneRequest) (response *CreateZoneResponse, err error) {
    if request == nil {
        request = NewCreateZoneRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "CreateZone")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateZoneResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAccelerationDomainsRequest() (request *DeleteAccelerationDomainsRequest) {
    request = &DeleteAccelerationDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteAccelerationDomains")
    
    
    return
}

func NewDeleteAccelerationDomainsResponse() (response *DeleteAccelerationDomainsResponse) {
    response = &DeleteAccelerationDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAccelerationDomains
// Delete acceleration domain names in batches.
//
// error code that may be returned:
//  INVALIDPARAMETER_ZONEISGRAYPUBLISHING = "InvalidParameter.ZoneIsGrayPublishing"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  INVALIDPARAMETERVALUE_INVALIDDOMAINSTATUS = "InvalidParameterValue.InvalidDomainStatus"
//  INVALIDPARAMETERVALUE_RECORDNEEDDEFAULTLOCATION = "InvalidParameterValue.RecordNeedDefaultLocation"
//  INVALIDPARAMETERVALUE_UNSUPPORTDELETEWEIGHTRECORD = "InvalidParameterValue.UnSupportDeleteWeightRecord"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  OPERATIONDENIED_UNSUPPORTDELETELOCATIONRECORD = "OperationDenied.UnSupportDeleteLocationRecord"
//  OPERATIONDENIED_UNSUPPORTDELETEWEIGHTRECORD = "OperationDenied.UnSupportDeleteWeightRecord"
//  OPERATIONDENIED_VERSIONCONTROLISGRAYING = "OperationDenied.VersionControlIsGraying"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE_DOMAINNOTFOUND = "ResourceUnavailable.DomainNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteAccelerationDomains(request *DeleteAccelerationDomainsRequest) (response *DeleteAccelerationDomainsResponse, err error) {
    return c.DeleteAccelerationDomainsWithContext(context.Background(), request)
}

// DeleteAccelerationDomains
// Delete acceleration domain names in batches.
//
// error code that may be returned:
//  INVALIDPARAMETER_ZONEISGRAYPUBLISHING = "InvalidParameter.ZoneIsGrayPublishing"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  INVALIDPARAMETERVALUE_INVALIDDOMAINSTATUS = "InvalidParameterValue.InvalidDomainStatus"
//  INVALIDPARAMETERVALUE_RECORDNEEDDEFAULTLOCATION = "InvalidParameterValue.RecordNeedDefaultLocation"
//  INVALIDPARAMETERVALUE_UNSUPPORTDELETEWEIGHTRECORD = "InvalidParameterValue.UnSupportDeleteWeightRecord"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  OPERATIONDENIED_UNSUPPORTDELETELOCATIONRECORD = "OperationDenied.UnSupportDeleteLocationRecord"
//  OPERATIONDENIED_UNSUPPORTDELETEWEIGHTRECORD = "OperationDenied.UnSupportDeleteWeightRecord"
//  OPERATIONDENIED_VERSIONCONTROLISGRAYING = "OperationDenied.VersionControlIsGraying"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE_DOMAINNOTFOUND = "ResourceUnavailable.DomainNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteAccelerationDomainsWithContext(ctx context.Context, request *DeleteAccelerationDomainsRequest) (response *DeleteAccelerationDomainsResponse, err error) {
    if request == nil {
        request = NewDeleteAccelerationDomainsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DeleteAccelerationDomains")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAccelerationDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAccelerationDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAliasDomainRequest() (request *DeleteAliasDomainRequest) {
    request = &DeleteAliasDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteAliasDomain")
    
    
    return
}

func NewDeleteAliasDomainResponse() (response *DeleteAliasDomainResponse) {
    response = &DeleteAliasDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAliasDomain
// Delete an alias domain name.
//
// This feature is only supported in the enterprise plan and is currently in beta test. To use it, please [contact us](https://www.tencentcloud.com/contact-us).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteAliasDomain(request *DeleteAliasDomainRequest) (response *DeleteAliasDomainResponse, err error) {
    return c.DeleteAliasDomainWithContext(context.Background(), request)
}

// DeleteAliasDomain
// Delete an alias domain name.
//
// This feature is only supported in the enterprise plan and is currently in beta test. To use it, please [contact us](https://www.tencentcloud.com/contact-us).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteAliasDomainWithContext(ctx context.Context, request *DeleteAliasDomainRequest) (response *DeleteAliasDomainResponse, err error) {
    if request == nil {
        request = NewDeleteAliasDomainRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DeleteAliasDomain")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAliasDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAliasDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApplicationProxyRequest() (request *DeleteApplicationProxyRequest) {
    request = &DeleteApplicationProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteApplicationProxy")
    
    
    return
}

func NewDeleteApplicationProxyResponse() (response *DeleteApplicationProxyResponse) {
    response = &DeleteApplicationProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteApplicationProxy
// This API is on an earlier version. If you want to call it, switch to the latest version. For details, see Deleting a Layer 4 Proxy Instance.
//
// ](https://www.tencentcloud.com/document/product/1552/103415?from_cn_redirect=1) .
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteApplicationProxy(request *DeleteApplicationProxyRequest) (response *DeleteApplicationProxyResponse, err error) {
    return c.DeleteApplicationProxyWithContext(context.Background(), request)
}

// DeleteApplicationProxy
// This API is on an earlier version. If you want to call it, switch to the latest version. For details, see Deleting a Layer 4 Proxy Instance.
//
// ](https://www.tencentcloud.com/document/product/1552/103415?from_cn_redirect=1) .
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteApplicationProxyWithContext(ctx context.Context, request *DeleteApplicationProxyRequest) (response *DeleteApplicationProxyResponse, err error) {
    if request == nil {
        request = NewDeleteApplicationProxyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DeleteApplicationProxy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApplicationProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteApplicationProxyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApplicationProxyRuleRequest() (request *DeleteApplicationProxyRuleRequest) {
    request = &DeleteApplicationProxyRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteApplicationProxyRule")
    
    
    return
}

func NewDeleteApplicationProxyRuleResponse() (response *DeleteApplicationProxyRuleResponse) {
    response = &DeleteApplicationProxyRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteApplicationProxyRule
// This API is on an earlier version. If you want to call it, switch to the latest version. For details, see Deleting L4 Proxy Rules (https://www.tencentcloud.com/document/product/1552/103414?from_cn_redirect=1).
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteApplicationProxyRule(request *DeleteApplicationProxyRuleRequest) (response *DeleteApplicationProxyRuleResponse, err error) {
    return c.DeleteApplicationProxyRuleWithContext(context.Background(), request)
}

// DeleteApplicationProxyRule
// This API is on an earlier version. If you want to call it, switch to the latest version. For details, see Deleting L4 Proxy Rules (https://www.tencentcloud.com/document/product/1552/103414?from_cn_redirect=1).
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteApplicationProxyRuleWithContext(ctx context.Context, request *DeleteApplicationProxyRuleRequest) (response *DeleteApplicationProxyRuleResponse, err error) {
    if request == nil {
        request = NewDeleteApplicationProxyRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DeleteApplicationProxyRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApplicationProxyRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteApplicationProxyRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteContentIdentifierRequest() (request *DeleteContentIdentifierRequest) {
    request = &DeleteContentIdentifierRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteContentIdentifier")
    
    
    return
}

func NewDeleteContentIdentifierResponse() (response *DeleteContentIdentifierResponse) {
    response = &DeleteContentIdentifierResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteContentIdentifier
// Delete the specified content identifier. The feature is open only in the allowlist.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteContentIdentifier(request *DeleteContentIdentifierRequest) (response *DeleteContentIdentifierResponse, err error) {
    return c.DeleteContentIdentifierWithContext(context.Background(), request)
}

// DeleteContentIdentifier
// Delete the specified content identifier. The feature is open only in the allowlist.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteContentIdentifierWithContext(ctx context.Context, request *DeleteContentIdentifierRequest) (response *DeleteContentIdentifierResponse, err error) {
    if request == nil {
        request = NewDeleteContentIdentifierRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DeleteContentIdentifier")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteContentIdentifier require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteContentIdentifierResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCustomErrorPageRequest() (request *DeleteCustomErrorPageRequest) {
    request = &DeleteCustomErrorPageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteCustomErrorPage")
    
    
    return
}

func NewDeleteCustomErrorPageResponse() (response *DeleteCustomErrorPageResponse) {
    response = &DeleteCustomErrorPageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCustomErrorPage
// This API is used to delete a custom response page.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteCustomErrorPage(request *DeleteCustomErrorPageRequest) (response *DeleteCustomErrorPageResponse, err error) {
    return c.DeleteCustomErrorPageWithContext(context.Background(), request)
}

// DeleteCustomErrorPage
// This API is used to delete a custom response page.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteCustomErrorPageWithContext(ctx context.Context, request *DeleteCustomErrorPageRequest) (response *DeleteCustomErrorPageResponse, err error) {
    if request == nil {
        request = NewDeleteCustomErrorPageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DeleteCustomErrorPage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCustomErrorPage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCustomErrorPageResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCustomGroupRequest() (request *DeleteCustomGroupRequest) {
    request = &DeleteCustomGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteCustomGroup")
    
    
    return
}

func NewDeleteCustomGroupResponse() (response *DeleteCustomGroupResponse) {
    response = &DeleteCustomGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCustomGroup
// This API is used to delete a specified custom group. If the group is referenced by a rule, deletion is not allowed. Remove the reference first.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteCustomGroup(request *DeleteCustomGroupRequest) (response *DeleteCustomGroupResponse, err error) {
    return c.DeleteCustomGroupWithContext(context.Background(), request)
}

// DeleteCustomGroup
// This API is used to delete a specified custom group. If the group is referenced by a rule, deletion is not allowed. Remove the reference first.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteCustomGroupWithContext(ctx context.Context, request *DeleteCustomGroupRequest) (response *DeleteCustomGroupResponse, err error) {
    if request == nil {
        request = NewDeleteCustomGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DeleteCustomGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCustomGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCustomGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDnsRecordsRequest() (request *DeleteDnsRecordsRequest) {
    request = &DeleteDnsRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteDnsRecords")
    
    
    return
}

func NewDeleteDnsRecordsResponse() (response *DeleteDnsRecordsResponse) {
    response = &DeleteDnsRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDnsRecords
// You can use this API to delete DNS records in batches.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDnsRecords(request *DeleteDnsRecordsRequest) (response *DeleteDnsRecordsResponse, err error) {
    return c.DeleteDnsRecordsWithContext(context.Background(), request)
}

// DeleteDnsRecords
// You can use this API to delete DNS records in batches.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDnsRecordsWithContext(ctx context.Context, request *DeleteDnsRecordsRequest) (response *DeleteDnsRecordsResponse, err error) {
    if request == nil {
        request = NewDeleteDnsRecordsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DeleteDnsRecords")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDnsRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDnsRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEdgeKVNamespaceRequest() (request *DeleteEdgeKVNamespaceRequest) {
    request = &DeleteEdgeKVNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteEdgeKVNamespace")
    
    
    return
}

func NewDeleteEdgeKVNamespaceResponse() (response *DeleteEdgeKVNamespaceResponse) {
    response = &DeleteEdgeKVNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteEdgeKVNamespace
// This API is used to delete a specified KV namespace. After deletion, all key-value pair data in the namespace will be cleared and cannot be recovered. If the namespace is referenced by an edge function, unbind the relationship before deletion.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_NAMESPACEINUSE = "InvalidParameter.NamespaceInUse"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEUNAVAILABLE_NAMESPACENOTFOUND = "ResourceUnavailable.NamespaceNotFound"
func (c *Client) DeleteEdgeKVNamespace(request *DeleteEdgeKVNamespaceRequest) (response *DeleteEdgeKVNamespaceResponse, err error) {
    return c.DeleteEdgeKVNamespaceWithContext(context.Background(), request)
}

// DeleteEdgeKVNamespace
// This API is used to delete a specified KV namespace. After deletion, all key-value pair data in the namespace will be cleared and cannot be recovered. If the namespace is referenced by an edge function, unbind the relationship before deletion.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_NAMESPACEINUSE = "InvalidParameter.NamespaceInUse"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEUNAVAILABLE_NAMESPACENOTFOUND = "ResourceUnavailable.NamespaceNotFound"
func (c *Client) DeleteEdgeKVNamespaceWithContext(ctx context.Context, request *DeleteEdgeKVNamespaceRequest) (response *DeleteEdgeKVNamespaceResponse, err error) {
    if request == nil {
        request = NewDeleteEdgeKVNamespaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DeleteEdgeKVNamespace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteEdgeKVNamespace require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteEdgeKVNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFunctionRequest() (request *DeleteFunctionRequest) {
    request = &DeleteFunctionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteFunction")
    
    
    return
}

func NewDeleteFunctionResponse() (response *DeleteFunctionResponse) {
    response = &DeleteFunctionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteFunction
// This API is used to delete an edge function. Once deleted, the function cannot be recovered, and associated trigger rules, environment variables, and replicas are also deleted.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RULEOPERATIONCONFLICT = "FailedOperation.RuleOperationConflict"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  OPERATIONDENIED_VERSIONCONTROLLOCKED = "OperationDenied.VersionControlLocked"
//  RESOURCEUNAVAILABLE_FUNCTIONNOTFOUND = "ResourceUnavailable.FunctionNotFound"
//  RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"
func (c *Client) DeleteFunction(request *DeleteFunctionRequest) (response *DeleteFunctionResponse, err error) {
    return c.DeleteFunctionWithContext(context.Background(), request)
}

// DeleteFunction
// This API is used to delete an edge function. Once deleted, the function cannot be recovered, and associated trigger rules, environment variables, and replicas are also deleted.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RULEOPERATIONCONFLICT = "FailedOperation.RuleOperationConflict"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  OPERATIONDENIED_VERSIONCONTROLLOCKED = "OperationDenied.VersionControlLocked"
//  RESOURCEUNAVAILABLE_FUNCTIONNOTFOUND = "ResourceUnavailable.FunctionNotFound"
//  RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"
func (c *Client) DeleteFunctionWithContext(ctx context.Context, request *DeleteFunctionRequest) (response *DeleteFunctionResponse, err error) {
    if request == nil {
        request = NewDeleteFunctionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DeleteFunction")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteFunction require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteFunctionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFunctionReplicaRequest() (request *DeleteFunctionReplicaRequest) {
    request = &DeleteFunctionReplicaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteFunctionReplica")
    
    
    return
}

func NewDeleteFunctionReplicaResponse() (response *DeleteFunctionReplicaResponse) {
    response = &DeleteFunctionReplicaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteFunctionReplica
// This API is used to delete a specified edge function replica.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RULEOPERATIONCONFLICT = "FailedOperation.RuleOperationConflict"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  OPERATIONDENIED_VERSIONCONTROLLOCKED = "OperationDenied.VersionControlLocked"
//  RESOURCEUNAVAILABLE_FUNCTIONNOTFOUND = "ResourceUnavailable.FunctionNotFound"
//  RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"
func (c *Client) DeleteFunctionReplica(request *DeleteFunctionReplicaRequest) (response *DeleteFunctionReplicaResponse, err error) {
    return c.DeleteFunctionReplicaWithContext(context.Background(), request)
}

// DeleteFunctionReplica
// This API is used to delete a specified edge function replica.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RULEOPERATIONCONFLICT = "FailedOperation.RuleOperationConflict"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  OPERATIONDENIED_VERSIONCONTROLLOCKED = "OperationDenied.VersionControlLocked"
//  RESOURCEUNAVAILABLE_FUNCTIONNOTFOUND = "ResourceUnavailable.FunctionNotFound"
//  RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"
func (c *Client) DeleteFunctionReplicaWithContext(ctx context.Context, request *DeleteFunctionReplicaRequest) (response *DeleteFunctionReplicaResponse, err error) {
    if request == nil {
        request = NewDeleteFunctionReplicaRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DeleteFunctionReplica")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteFunctionReplica require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteFunctionReplicaResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFunctionRulesRequest() (request *DeleteFunctionRulesRequest) {
    request = &DeleteFunctionRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteFunctionRules")
    
    
    return
}

func NewDeleteFunctionRulesResponse() (response *DeleteFunctionRulesResponse) {
    response = &DeleteFunctionRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteFunctionRules
// This API is used to delete a trigger rule for an edge function.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RULEOPERATIONCONFLICT = "FailedOperation.RuleOperationConflict"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  RESOURCEUNAVAILABLE_RULENOTFOUND = "ResourceUnavailable.RuleNotFound"
//  RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"
func (c *Client) DeleteFunctionRules(request *DeleteFunctionRulesRequest) (response *DeleteFunctionRulesResponse, err error) {
    return c.DeleteFunctionRulesWithContext(context.Background(), request)
}

// DeleteFunctionRules
// This API is used to delete a trigger rule for an edge function.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RULEOPERATIONCONFLICT = "FailedOperation.RuleOperationConflict"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  RESOURCEUNAVAILABLE_RULENOTFOUND = "ResourceUnavailable.RuleNotFound"
//  RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"
func (c *Client) DeleteFunctionRulesWithContext(ctx context.Context, request *DeleteFunctionRulesRequest) (response *DeleteFunctionRulesResponse, err error) {
    if request == nil {
        request = NewDeleteFunctionRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DeleteFunctionRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteFunctionRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteFunctionRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteInferenceAPITokenRequest() (request *DeleteInferenceAPITokenRequest) {
    request = &DeleteInferenceAPITokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteInferenceAPIToken")
    
    
    return
}

func NewDeleteInferenceAPITokenResponse() (response *DeleteInferenceAPITokenResponse) {
    response = &DeleteInferenceAPITokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteInferenceAPIToken
// Delete an inference API Token. Once deleted, the Token expires immediately, and requests using it to access the inference service will not pass authentication.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_ZONENOTFOUND = "InvalidParameterValue.ZoneNotFound"
func (c *Client) DeleteInferenceAPIToken(request *DeleteInferenceAPITokenRequest) (response *DeleteInferenceAPITokenResponse, err error) {
    return c.DeleteInferenceAPITokenWithContext(context.Background(), request)
}

// DeleteInferenceAPIToken
// Delete an inference API Token. Once deleted, the Token expires immediately, and requests using it to access the inference service will not pass authentication.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_ZONENOTFOUND = "InvalidParameterValue.ZoneNotFound"
func (c *Client) DeleteInferenceAPITokenWithContext(ctx context.Context, request *DeleteInferenceAPITokenRequest) (response *DeleteInferenceAPITokenResponse, err error) {
    if request == nil {
        request = NewDeleteInferenceAPITokenRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DeleteInferenceAPIToken")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteInferenceAPIToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteInferenceAPITokenResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteJustInTimeTranscodeTemplatesRequest() (request *DeleteJustInTimeTranscodeTemplatesRequest) {
    request = &DeleteJustInTimeTranscodeTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteJustInTimeTranscodeTemplates")
    
    
    return
}

func NewDeleteJustInTimeTranscodeTemplatesResponse() (response *DeleteJustInTimeTranscodeTemplatesResponse) {
    response = &DeleteJustInTimeTranscodeTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteJustInTimeTranscodeTemplates
// Delete a just in time transcoding template based on the unique template identifier under the site ID.
//
// error code that may be returned:
//  FAILEDOPERATION_PRODUCTDISCONTINUED = "FailedOperation.ProductDiscontinued"
//  FAILEDOPERATION_PRODUCTNOTACTIVATED = "FailedOperation.ProductNotActivated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TEMPLATENOTCUSTOM = "InvalidParameterValue.TemplateNotCustom"
//  INVALIDPARAMETERVALUE_TEMPLATENOTFOUND = "InvalidParameterValue.TemplateNotFound"
//  INVALIDPARAMETERVALUE_ZONENOTFOUND = "InvalidParameterValue.ZoneNotFound"
//  INVALIDPARAMETERVALUE_ZONEPAUSED = "InvalidParameterValue.ZonePaused"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteJustInTimeTranscodeTemplates(request *DeleteJustInTimeTranscodeTemplatesRequest) (response *DeleteJustInTimeTranscodeTemplatesResponse, err error) {
    return c.DeleteJustInTimeTranscodeTemplatesWithContext(context.Background(), request)
}

// DeleteJustInTimeTranscodeTemplates
// Delete a just in time transcoding template based on the unique template identifier under the site ID.
//
// error code that may be returned:
//  FAILEDOPERATION_PRODUCTDISCONTINUED = "FailedOperation.ProductDiscontinued"
//  FAILEDOPERATION_PRODUCTNOTACTIVATED = "FailedOperation.ProductNotActivated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TEMPLATENOTCUSTOM = "InvalidParameterValue.TemplateNotCustom"
//  INVALIDPARAMETERVALUE_TEMPLATENOTFOUND = "InvalidParameterValue.TemplateNotFound"
//  INVALIDPARAMETERVALUE_ZONENOTFOUND = "InvalidParameterValue.ZoneNotFound"
//  INVALIDPARAMETERVALUE_ZONEPAUSED = "InvalidParameterValue.ZonePaused"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteJustInTimeTranscodeTemplatesWithContext(ctx context.Context, request *DeleteJustInTimeTranscodeTemplatesRequest) (response *DeleteJustInTimeTranscodeTemplatesResponse, err error) {
    if request == nil {
        request = NewDeleteJustInTimeTranscodeTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DeleteJustInTimeTranscodeTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteJustInTimeTranscodeTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteJustInTimeTranscodeTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteL4ProxyRequest() (request *DeleteL4ProxyRequest) {
    request = &DeleteL4ProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteL4Proxy")
    
    
    return
}

func NewDeleteL4ProxyResponse() (response *DeleteL4ProxyResponse) {
    response = &DeleteL4ProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteL4Proxy
// This API is used to delete a layer 4 proxy instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteL4Proxy(request *DeleteL4ProxyRequest) (response *DeleteL4ProxyResponse, err error) {
    return c.DeleteL4ProxyWithContext(context.Background(), request)
}

// DeleteL4Proxy
// This API is used to delete a layer 4 proxy instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteL4ProxyWithContext(ctx context.Context, request *DeleteL4ProxyRequest) (response *DeleteL4ProxyResponse, err error) {
    if request == nil {
        request = NewDeleteL4ProxyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DeleteL4Proxy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteL4Proxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteL4ProxyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteL4ProxyRulesRequest() (request *DeleteL4ProxyRulesRequest) {
    request = &DeleteL4ProxyRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteL4ProxyRules")
    
    
    return
}

func NewDeleteL4ProxyRulesResponse() (response *DeleteL4ProxyRulesResponse) {
    response = &DeleteL4ProxyRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteL4ProxyRules
// This API is used to delete l4 proxy rules and supports both individual and batch operation.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteL4ProxyRules(request *DeleteL4ProxyRulesRequest) (response *DeleteL4ProxyRulesResponse, err error) {
    return c.DeleteL4ProxyRulesWithContext(context.Background(), request)
}

// DeleteL4ProxyRules
// This API is used to delete l4 proxy rules and supports both individual and batch operation.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteL4ProxyRulesWithContext(ctx context.Context, request *DeleteL4ProxyRulesRequest) (response *DeleteL4ProxyRulesResponse, err error) {
    if request == nil {
        request = NewDeleteL4ProxyRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DeleteL4ProxyRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteL4ProxyRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteL4ProxyRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteL7AccRulesRequest() (request *DeleteL7AccRulesRequest) {
    request = &DeleteL7AccRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteL7AccRules")
    
    
    return
}

func NewDeleteL7AccRulesResponse() (response *DeleteL7AccRulesResponse) {
    response = &DeleteL7AccRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteL7AccRules
// This API is used to delete rules from the rule engine (https://www.tencentcloud.com/document/product/1552/70901?from_cn_redirect=1). Batch deletion supported.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_INVALIDHTTPS = "InvalidParameter.InvalidHttps"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDRULEENGINENOTFOUND = "InvalidParameter.InvalidRuleEngineNotFound"
//  INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteL7AccRules(request *DeleteL7AccRulesRequest) (response *DeleteL7AccRulesResponse, err error) {
    return c.DeleteL7AccRulesWithContext(context.Background(), request)
}

// DeleteL7AccRules
// This API is used to delete rules from the rule engine (https://www.tencentcloud.com/document/product/1552/70901?from_cn_redirect=1). Batch deletion supported.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_INVALIDHTTPS = "InvalidParameter.InvalidHttps"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDRULEENGINENOTFOUND = "InvalidParameter.InvalidRuleEngineNotFound"
//  INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteL7AccRulesWithContext(ctx context.Context, request *DeleteL7AccRulesRequest) (response *DeleteL7AccRulesResponse, err error) {
    if request == nil {
        request = NewDeleteL7AccRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DeleteL7AccRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteL7AccRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteL7AccRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLoadBalancerRequest() (request *DeleteLoadBalancerRequest) {
    request = &DeleteLoadBalancerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteLoadBalancer")
    
    
    return
}

func NewDeleteLoadBalancerResponse() (response *DeleteLoadBalancerResponse) {
    response = &DeleteLoadBalancerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLoadBalancer
// This API is used to delete a load balancing instance. If the load balancing instance is referenced by other services, such as a Layer 4 proxy, it cannot be deleted. Remove the reference relationship first. The load balancer feature is in beta test. If you need to use it, contact us (https://www.tencentcloud.com/online?from_cn_redirect=1-service).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_LOADBALANCERUSEDINL4PROXY = "InvalidParameter.LoadBalancerUsedInL4Proxy"
//  INVALIDPARAMETER_LOADBALANCERUSEDINL7DOMAIN = "InvalidParameter.LoadBalancerUsedInL7Domain"
//  INVALIDPARAMETER_LOADBALANCERUSEDINRULEENGINE = "InvalidParameter.LoadBalancerUsedInRuleEngine"
func (c *Client) DeleteLoadBalancer(request *DeleteLoadBalancerRequest) (response *DeleteLoadBalancerResponse, err error) {
    return c.DeleteLoadBalancerWithContext(context.Background(), request)
}

// DeleteLoadBalancer
// This API is used to delete a load balancing instance. If the load balancing instance is referenced by other services, such as a Layer 4 proxy, it cannot be deleted. Remove the reference relationship first. The load balancer feature is in beta test. If you need to use it, contact us (https://www.tencentcloud.com/online?from_cn_redirect=1-service).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_LOADBALANCERUSEDINL4PROXY = "InvalidParameter.LoadBalancerUsedInL4Proxy"
//  INVALIDPARAMETER_LOADBALANCERUSEDINL7DOMAIN = "InvalidParameter.LoadBalancerUsedInL7Domain"
//  INVALIDPARAMETER_LOADBALANCERUSEDINRULEENGINE = "InvalidParameter.LoadBalancerUsedInRuleEngine"
func (c *Client) DeleteLoadBalancerWithContext(ctx context.Context, request *DeleteLoadBalancerRequest) (response *DeleteLoadBalancerResponse, err error) {
    if request == nil {
        request = NewDeleteLoadBalancerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DeleteLoadBalancer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLoadBalancer require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLoadBalancerResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMultiPathGatewayRequest() (request *DeleteMultiPathGatewayRequest) {
    request = &DeleteMultiPathGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteMultiPathGateway")
    
    
    return
}

func NewDeleteMultiPathGatewayResponse() (response *DeleteMultiPathGatewayResponse) {
    response = &DeleteMultiPathGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteMultiPathGateway
// This API is used to delete multi-channel security acceleration gateways, including private gateways and Cloud Gateways.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DeleteMultiPathGateway(request *DeleteMultiPathGatewayRequest) (response *DeleteMultiPathGatewayResponse, err error) {
    return c.DeleteMultiPathGatewayWithContext(context.Background(), request)
}

// DeleteMultiPathGateway
// This API is used to delete multi-channel security acceleration gateways, including private gateways and Cloud Gateways.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DeleteMultiPathGatewayWithContext(ctx context.Context, request *DeleteMultiPathGatewayRequest) (response *DeleteMultiPathGatewayResponse, err error) {
    if request == nil {
        request = NewDeleteMultiPathGatewayRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DeleteMultiPathGateway")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMultiPathGateway require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMultiPathGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMultiPathGatewayLineRequest() (request *DeleteMultiPathGatewayLineRequest) {
    request = &DeleteMultiPathGatewayLineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteMultiPathGatewayLine")
    
    
    return
}

func NewDeleteMultiPathGatewayLineResponse() (response *DeleteMultiPathGatewayLineResponse) {
    response = &DeleteMultiPathGatewayLineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteMultiPathGatewayLine
// This API is used to delete lines integrated with multi-channel security acceleration gateways. Only custom lines support deletion.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DeleteMultiPathGatewayLine(request *DeleteMultiPathGatewayLineRequest) (response *DeleteMultiPathGatewayLineResponse, err error) {
    return c.DeleteMultiPathGatewayLineWithContext(context.Background(), request)
}

// DeleteMultiPathGatewayLine
// This API is used to delete lines integrated with multi-channel security acceleration gateways. Only custom lines support deletion.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DeleteMultiPathGatewayLineWithContext(ctx context.Context, request *DeleteMultiPathGatewayLineRequest) (response *DeleteMultiPathGatewayLineResponse, err error) {
    if request == nil {
        request = NewDeleteMultiPathGatewayLineRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DeleteMultiPathGatewayLine")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMultiPathGatewayLine require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMultiPathGatewayLineResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteOriginGroupRequest() (request *DeleteOriginGroupRequest) {
    request = &DeleteOriginGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteOriginGroup")
    
    
    return
}

func NewDeleteOriginGroupResponse() (response *DeleteOriginGroupResponse) {
    response = &DeleteOriginGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteOriginGroup
// Delete an origin group. If the origin group is still referenced by a service (for example, Layer 4 proxy, DNS, load balancing, or a rule), deletion is not allowed.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ORIGINGROUPACCELERATIONDOMAINUSED = "OperationDenied.OriginGroupAccelerationDomainUsed"
//  OPERATIONDENIED_ORIGINGROUPL4USED = "OperationDenied.OriginGroupL4Used"
//  OPERATIONDENIED_ORIGINGROUPLBUSED = "OperationDenied.OriginGroupLBUsed"
//  OPERATIONDENIED_ORIGINGROUPRULEENGINEUSED = "OperationDenied.OriginGroupRuleEngineUsed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteOriginGroup(request *DeleteOriginGroupRequest) (response *DeleteOriginGroupResponse, err error) {
    return c.DeleteOriginGroupWithContext(context.Background(), request)
}

// DeleteOriginGroup
// Delete an origin group. If the origin group is still referenced by a service (for example, Layer 4 proxy, DNS, load balancing, or a rule), deletion is not allowed.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ORIGINGROUPACCELERATIONDOMAINUSED = "OperationDenied.OriginGroupAccelerationDomainUsed"
//  OPERATIONDENIED_ORIGINGROUPL4USED = "OperationDenied.OriginGroupL4Used"
//  OPERATIONDENIED_ORIGINGROUPLBUSED = "OperationDenied.OriginGroupLBUsed"
//  OPERATIONDENIED_ORIGINGROUPRULEENGINEUSED = "OperationDenied.OriginGroupRuleEngineUsed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteOriginGroupWithContext(ctx context.Context, request *DeleteOriginGroupRequest) (response *DeleteOriginGroupResponse, err error) {
    if request == nil {
        request = NewDeleteOriginGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DeleteOriginGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteOriginGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteOriginGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRealtimeLogDeliveryTaskRequest() (request *DeleteRealtimeLogDeliveryTaskRequest) {
    request = &DeleteRealtimeLogDeliveryTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteRealtimeLogDeliveryTask")
    
    
    return
}

func NewDeleteRealtimeLogDeliveryTaskResponse() (response *DeleteRealtimeLogDeliveryTaskResponse) {
    response = &DeleteRealtimeLogDeliveryTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRealtimeLogDeliveryTask
// This API is used to delete a real-time log delivery task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteRealtimeLogDeliveryTask(request *DeleteRealtimeLogDeliveryTaskRequest) (response *DeleteRealtimeLogDeliveryTaskResponse, err error) {
    return c.DeleteRealtimeLogDeliveryTaskWithContext(context.Background(), request)
}

// DeleteRealtimeLogDeliveryTask
// This API is used to delete a real-time log delivery task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteRealtimeLogDeliveryTaskWithContext(ctx context.Context, request *DeleteRealtimeLogDeliveryTaskRequest) (response *DeleteRealtimeLogDeliveryTaskResponse, err error) {
    if request == nil {
        request = NewDeleteRealtimeLogDeliveryTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DeleteRealtimeLogDeliveryTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRealtimeLogDeliveryTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRealtimeLogDeliveryTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRulesRequest() (request *DeleteRulesRequest) {
    request = &DeleteRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteRules")
    
    
    return
}

func NewDeleteRulesResponse() (response *DeleteRulesResponse) {
    response = &DeleteRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRules
// This API is on an earlier version for deleting rule engine APIs. EdgeOne completed a comprehensive upgrade of rule engine APIs on January 21, 2025. For details on the new API for deleting L7 acceleration rules, see [DeleteL7AccRules](https://www.tencentcloud.com/document/product/1552/115821?from_cn_redirect=1).
//
// <p style="color: red;">Note: Starting from January 21, 2025, the old version API will stop update and iteration. New features will only be provided in the new version API, and the existing capabilities supported by the old version API will not be affected. To avoid data field conflicts when using the old version API, migrate to the new version rule engine API as soon as possible.</p>
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_INVALIDHTTPS = "InvalidParameter.InvalidHttps"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDRULEENGINENOTFOUND = "InvalidParameter.InvalidRuleEngineNotFound"
//  INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteRules(request *DeleteRulesRequest) (response *DeleteRulesResponse, err error) {
    return c.DeleteRulesWithContext(context.Background(), request)
}

// DeleteRules
// This API is on an earlier version for deleting rule engine APIs. EdgeOne completed a comprehensive upgrade of rule engine APIs on January 21, 2025. For details on the new API for deleting L7 acceleration rules, see [DeleteL7AccRules](https://www.tencentcloud.com/document/product/1552/115821?from_cn_redirect=1).
//
// <p style="color: red;">Note: Starting from January 21, 2025, the old version API will stop update and iteration. New features will only be provided in the new version API, and the existing capabilities supported by the old version API will not be affected. To avoid data field conflicts when using the old version API, migrate to the new version rule engine API as soon as possible.</p>
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_INVALIDHTTPS = "InvalidParameter.InvalidHttps"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDRULEENGINENOTFOUND = "InvalidParameter.InvalidRuleEngineNotFound"
//  INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteRulesWithContext(ctx context.Context, request *DeleteRulesRequest) (response *DeleteRulesResponse, err error) {
    if request == nil {
        request = NewDeleteRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DeleteRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSecurityAPIResourceRequest() (request *DeleteSecurityAPIResourceRequest) {
    request = &DeleteSecurityAPIResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteSecurityAPIResource")
    
    
    return
}

func NewDeleteSecurityAPIResourceResponse() (response *DeleteSecurityAPIResourceResponse) {
    response = &DeleteSecurityAPIResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSecurityAPIResource
// Used to delete API resources.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_INVALIDHTTPS = "InvalidParameter.InvalidHttps"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDRULEENGINENOTFOUND = "InvalidParameter.InvalidRuleEngineNotFound"
//  INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteSecurityAPIResource(request *DeleteSecurityAPIResourceRequest) (response *DeleteSecurityAPIResourceResponse, err error) {
    return c.DeleteSecurityAPIResourceWithContext(context.Background(), request)
}

// DeleteSecurityAPIResource
// Used to delete API resources.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_INVALIDHTTPS = "InvalidParameter.InvalidHttps"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDRULEENGINENOTFOUND = "InvalidParameter.InvalidRuleEngineNotFound"
//  INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteSecurityAPIResourceWithContext(ctx context.Context, request *DeleteSecurityAPIResourceRequest) (response *DeleteSecurityAPIResourceResponse, err error) {
    if request == nil {
        request = NewDeleteSecurityAPIResourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DeleteSecurityAPIResource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSecurityAPIResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSecurityAPIResourceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSecurityAPIServiceRequest() (request *DeleteSecurityAPIServiceRequest) {
    request = &DeleteSecurityAPIServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteSecurityAPIService")
    
    
    return
}

func NewDeleteSecurityAPIServiceResponse() (response *DeleteSecurityAPIServiceResponse) {
    response = &DeleteSecurityAPIServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSecurityAPIService
// Used to delete an API service.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_INVALIDHTTPS = "InvalidParameter.InvalidHttps"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDRULEENGINENOTFOUND = "InvalidParameter.InvalidRuleEngineNotFound"
//  INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteSecurityAPIService(request *DeleteSecurityAPIServiceRequest) (response *DeleteSecurityAPIServiceResponse, err error) {
    return c.DeleteSecurityAPIServiceWithContext(context.Background(), request)
}

// DeleteSecurityAPIService
// Used to delete an API service.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_INVALIDHTTPS = "InvalidParameter.InvalidHttps"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDRULEENGINENOTFOUND = "InvalidParameter.InvalidRuleEngineNotFound"
//  INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteSecurityAPIServiceWithContext(ctx context.Context, request *DeleteSecurityAPIServiceRequest) (response *DeleteSecurityAPIServiceResponse, err error) {
    if request == nil {
        request = NewDeleteSecurityAPIServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DeleteSecurityAPIService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSecurityAPIService require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSecurityAPIServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSecurityClientAttesterRequest() (request *DeleteSecurityClientAttesterRequest) {
    request = &DeleteSecurityClientAttesterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteSecurityClientAttester")
    
    
    return
}

func NewDeleteSecurityClientAttesterResponse() (response *DeleteSecurityClientAttesterResponse) {
    response = &DeleteSecurityClientAttesterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSecurityClientAttester
// Delete client authentication options.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_INVALIDHTTPS = "InvalidParameter.InvalidHttps"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDRULEENGINENOTFOUND = "InvalidParameter.InvalidRuleEngineNotFound"
//  INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteSecurityClientAttester(request *DeleteSecurityClientAttesterRequest) (response *DeleteSecurityClientAttesterResponse, err error) {
    return c.DeleteSecurityClientAttesterWithContext(context.Background(), request)
}

// DeleteSecurityClientAttester
// Delete client authentication options.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_INVALIDHTTPS = "InvalidParameter.InvalidHttps"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDRULEENGINENOTFOUND = "InvalidParameter.InvalidRuleEngineNotFound"
//  INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteSecurityClientAttesterWithContext(ctx context.Context, request *DeleteSecurityClientAttesterRequest) (response *DeleteSecurityClientAttesterResponse, err error) {
    if request == nil {
        request = NewDeleteSecurityClientAttesterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DeleteSecurityClientAttester")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSecurityClientAttester require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSecurityClientAttesterResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSecurityIPGroupRequest() (request *DeleteSecurityIPGroupRequest) {
    request = &DeleteSecurityIPGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteSecurityIPGroup")
    
    
    return
}

func NewDeleteSecurityIPGroupResponse() (response *DeleteSecurityIPGroupResponse) {
    response = &DeleteSecurityIPGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSecurityIPGroup
// Delete a specified IP group. If a rule references the IP group, deletion is not allowed.
//
// error code that may be returned:
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteSecurityIPGroup(request *DeleteSecurityIPGroupRequest) (response *DeleteSecurityIPGroupResponse, err error) {
    return c.DeleteSecurityIPGroupWithContext(context.Background(), request)
}

// DeleteSecurityIPGroup
// Delete a specified IP group. If a rule references the IP group, deletion is not allowed.
//
// error code that may be returned:
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteSecurityIPGroupWithContext(ctx context.Context, request *DeleteSecurityIPGroupRequest) (response *DeleteSecurityIPGroupResponse, err error) {
    if request == nil {
        request = NewDeleteSecurityIPGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DeleteSecurityIPGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSecurityIPGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSecurityIPGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSecurityJSInjectionRuleRequest() (request *DeleteSecurityJSInjectionRuleRequest) {
    request = &DeleteSecurityJSInjectionRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteSecurityJSInjectionRule")
    
    
    return
}

func NewDeleteSecurityJSInjectionRuleResponse() (response *DeleteSecurityJSInjectionRuleResponse) {
    response = &DeleteSecurityJSInjectionRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSecurityJSInjectionRule
// This API is used to delete JavaScript injection rules.
//
// error code that may be returned:
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteSecurityJSInjectionRule(request *DeleteSecurityJSInjectionRuleRequest) (response *DeleteSecurityJSInjectionRuleResponse, err error) {
    return c.DeleteSecurityJSInjectionRuleWithContext(context.Background(), request)
}

// DeleteSecurityJSInjectionRule
// This API is used to delete JavaScript injection rules.
//
// error code that may be returned:
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteSecurityJSInjectionRuleWithContext(ctx context.Context, request *DeleteSecurityJSInjectionRuleRequest) (response *DeleteSecurityJSInjectionRuleResponse, err error) {
    if request == nil {
        request = NewDeleteSecurityJSInjectionRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DeleteSecurityJSInjectionRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSecurityJSInjectionRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSecurityJSInjectionRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSharedCNAMERequest() (request *DeleteSharedCNAMERequest) {
    request = &DeleteSharedCNAMERequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteSharedCNAME")
    
    
    return
}

func NewDeleteSharedCNAMEResponse() (response *DeleteSharedCNAMEResponse) {
    response = &DeleteSharedCNAMEResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSharedCNAME
// Used to delete a shared CNAME. This feature is in beta test.
//
// error code that may be returned:
//  RESOURCEINUSE_SHAREDCNAME = "ResourceInUse.SharedCNAME"
func (c *Client) DeleteSharedCNAME(request *DeleteSharedCNAMERequest) (response *DeleteSharedCNAMEResponse, err error) {
    return c.DeleteSharedCNAMEWithContext(context.Background(), request)
}

// DeleteSharedCNAME
// Used to delete a shared CNAME. This feature is in beta test.
//
// error code that may be returned:
//  RESOURCEINUSE_SHAREDCNAME = "ResourceInUse.SharedCNAME"
func (c *Client) DeleteSharedCNAMEWithContext(ctx context.Context, request *DeleteSharedCNAMERequest) (response *DeleteSharedCNAMEResponse, err error) {
    if request == nil {
        request = NewDeleteSharedCNAMERequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DeleteSharedCNAME")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSharedCNAME require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSharedCNAMEResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWebSecurityTemplateRequest() (request *DeleteWebSecurityTemplateRequest) {
    request = &DeleteWebSecurityTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteWebSecurityTemplate")
    
    
    return
}

func NewDeleteWebSecurityTemplateResponse() (response *DeleteWebSecurityTemplateResponse) {
    response = &DeleteWebSecurityTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteWebSecurityTemplate
// Delete a security policy configuration template
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DeleteWebSecurityTemplate(request *DeleteWebSecurityTemplateRequest) (response *DeleteWebSecurityTemplateResponse, err error) {
    return c.DeleteWebSecurityTemplateWithContext(context.Background(), request)
}

// DeleteWebSecurityTemplate
// Delete a security policy configuration template
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DeleteWebSecurityTemplateWithContext(ctx context.Context, request *DeleteWebSecurityTemplateRequest) (response *DeleteWebSecurityTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteWebSecurityTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DeleteWebSecurityTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWebSecurityTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteWebSecurityTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteZoneRequest() (request *DeleteZoneRequest) {
    request = &DeleteZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteZone")
    
    
    return
}

func NewDeleteZoneResponse() (response *DeleteZoneResponse) {
    response = &DeleteZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteZone
// Delete a site.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_COMPLIANCEFORBIDDEN = "OperationDenied.ComplianceForbidden"
//  OPERATIONDENIED_DELETEZONEPRECHECKFAILED = "OperationDenied.DeleteZonePreCheckFailed"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_ZONEISBINDINGSHAREDCNAME = "OperationDenied.ZoneIsBindingSharedCNAME"
//  OPERATIONDENIED_ZONEISREFERENCECUSTOMERRORPAGE = "OperationDenied.ZoneIsReferenceCustomErrorPage"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteZone(request *DeleteZoneRequest) (response *DeleteZoneResponse, err error) {
    return c.DeleteZoneWithContext(context.Background(), request)
}

// DeleteZone
// Delete a site.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_COMPLIANCEFORBIDDEN = "OperationDenied.ComplianceForbidden"
//  OPERATIONDENIED_DELETEZONEPRECHECKFAILED = "OperationDenied.DeleteZonePreCheckFailed"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_ZONEISBINDINGSHAREDCNAME = "OperationDenied.ZoneIsBindingSharedCNAME"
//  OPERATIONDENIED_ZONEISREFERENCECUSTOMERRORPAGE = "OperationDenied.ZoneIsReferenceCustomErrorPage"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteZoneWithContext(ctx context.Context, request *DeleteZoneRequest) (response *DeleteZoneResponse, err error) {
    if request == nil {
        request = NewDeleteZoneRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DeleteZone")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteZoneResponse()
    err = c.Send(request, response)
    return
}

func NewDeployConfigGroupVersionRequest() (request *DeployConfigGroupVersionRequest) {
    request = &DeployConfigGroupVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeployConfigGroupVersion")
    
    
    return
}

func NewDeployConfigGroupVersionResponse() (response *DeployConfigGroupVersionResponse) {
    response = &DeployConfigGroupVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeployConfigGroupVersion
// In version management mode, this API is used for version release. You can release a version to the testing environment or production environment by EnvId. The version management feature is in beta test and currently only available to the allowlist.
//
// error code that may be returned:
//  FAILEDOPERATION_CONFIGCONDITIONSYNTAXERROR = "FailedOperation.ConfigConditionSyntaxError"
//  FAILEDOPERATION_CONFIGCONDITIONUNKNOWNTARGET = "FailedOperation.ConfigConditionUnknownTarget"
//  FAILEDOPERATION_CONFIGCONDITIONVALUEEMPTYERROR = "FailedOperation.ConfigConditionValueEmptyError"
//  FAILEDOPERATION_CONFIGFIELDTYPEERROR = "FailedOperation.ConfigFieldTypeError"
//  FAILEDOPERATION_CONFIGFORMATERROR = "FailedOperation.ConfigFormatError"
//  FAILEDOPERATION_CONFIGMALFORMEDCONTENT = "FailedOperation.ConfigMalformedContent"
//  FAILEDOPERATION_CONFIGPARAMVALIDATEERRORS = "FailedOperation.ConfigParamValidateErrors"
//  FAILEDOPERATION_CONFIGUNKNOWNFIELD = "FailedOperation.ConfigUnknownField"
//  FAILEDOPERATION_CONFIGUNSUPPORTEDFORMATVERSION = "FailedOperation.ConfigUnsupportedFormatVersion"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_DOMAINSTATUSUNSTABLE = "OperationDenied.DomainStatusUnstable"
//  OPERATIONDENIED_ENVNOTREADY = "OperationDenied.EnvNotReady"
//  OPERATIONDENIED_NOTINVERSIONCONTROLWHITELIST = "OperationDenied.NotInVersionControlWhiteList"
//  OPERATIONDENIED_VERSIONCONTROLLOCKED = "OperationDenied.VersionControlLocked"
//  OPERATIONDENIED_WORKMODENOTINVERSIONCONTROL = "OperationDenied.WorkModeNotInVersionControl"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_VERSIONNOTFOUND = "ResourceNotFound.VersionNotFound"
func (c *Client) DeployConfigGroupVersion(request *DeployConfigGroupVersionRequest) (response *DeployConfigGroupVersionResponse, err error) {
    return c.DeployConfigGroupVersionWithContext(context.Background(), request)
}

// DeployConfigGroupVersion
// In version management mode, this API is used for version release. You can release a version to the testing environment or production environment by EnvId. The version management feature is in beta test and currently only available to the allowlist.
//
// error code that may be returned:
//  FAILEDOPERATION_CONFIGCONDITIONSYNTAXERROR = "FailedOperation.ConfigConditionSyntaxError"
//  FAILEDOPERATION_CONFIGCONDITIONUNKNOWNTARGET = "FailedOperation.ConfigConditionUnknownTarget"
//  FAILEDOPERATION_CONFIGCONDITIONVALUEEMPTYERROR = "FailedOperation.ConfigConditionValueEmptyError"
//  FAILEDOPERATION_CONFIGFIELDTYPEERROR = "FailedOperation.ConfigFieldTypeError"
//  FAILEDOPERATION_CONFIGFORMATERROR = "FailedOperation.ConfigFormatError"
//  FAILEDOPERATION_CONFIGMALFORMEDCONTENT = "FailedOperation.ConfigMalformedContent"
//  FAILEDOPERATION_CONFIGPARAMVALIDATEERRORS = "FailedOperation.ConfigParamValidateErrors"
//  FAILEDOPERATION_CONFIGUNKNOWNFIELD = "FailedOperation.ConfigUnknownField"
//  FAILEDOPERATION_CONFIGUNSUPPORTEDFORMATVERSION = "FailedOperation.ConfigUnsupportedFormatVersion"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_DOMAINSTATUSUNSTABLE = "OperationDenied.DomainStatusUnstable"
//  OPERATIONDENIED_ENVNOTREADY = "OperationDenied.EnvNotReady"
//  OPERATIONDENIED_NOTINVERSIONCONTROLWHITELIST = "OperationDenied.NotInVersionControlWhiteList"
//  OPERATIONDENIED_VERSIONCONTROLLOCKED = "OperationDenied.VersionControlLocked"
//  OPERATIONDENIED_WORKMODENOTINVERSIONCONTROL = "OperationDenied.WorkModeNotInVersionControl"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_VERSIONNOTFOUND = "ResourceNotFound.VersionNotFound"
func (c *Client) DeployConfigGroupVersionWithContext(ctx context.Context, request *DeployConfigGroupVersionRequest) (response *DeployConfigGroupVersionResponse, err error) {
    if request == nil {
        request = NewDeployConfigGroupVersionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DeployConfigGroupVersion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeployConfigGroupVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeployConfigGroupVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccelerationDomainsRequest() (request *DescribeAccelerationDomainsRequest) {
    request = &DescribeAccelerationDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeAccelerationDomains")
    
    
    return
}

func NewDescribeAccelerationDomainsResponse() (response *DescribeAccelerationDomainsResponse) {
    response = &DescribeAccelerationDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccelerationDomains
// This API is used to view domain information under a site, including acceleration domain names, origin servers, and domain name status. You can view information of all domains under a site, or specify filter conditions to query specific domain information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_DOMAINONTRAFFICSCHEDULING = "InvalidParameter.DomainOnTrafficScheduling"
//  INVALIDPARAMETER_TOOMANYFILTERVALUES = "InvalidParameter.TooManyFilterValues"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeAccelerationDomains(request *DescribeAccelerationDomainsRequest) (response *DescribeAccelerationDomainsResponse, err error) {
    return c.DescribeAccelerationDomainsWithContext(context.Background(), request)
}

// DescribeAccelerationDomains
// This API is used to view domain information under a site, including acceleration domain names, origin servers, and domain name status. You can view information of all domains under a site, or specify filter conditions to query specific domain information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_DOMAINONTRAFFICSCHEDULING = "InvalidParameter.DomainOnTrafficScheduling"
//  INVALIDPARAMETER_TOOMANYFILTERVALUES = "InvalidParameter.TooManyFilterValues"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeAccelerationDomainsWithContext(ctx context.Context, request *DescribeAccelerationDomainsRequest) (response *DescribeAccelerationDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeAccelerationDomainsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeAccelerationDomains")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccelerationDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccelerationDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAliasDomainsRequest() (request *DescribeAliasDomainsRequest) {
    request = &DescribeAliasDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeAliasDomains")
    
    
    return
}

func NewDescribeAliasDomainsResponse() (response *DescribeAliasDomainsResponse) {
    response = &DescribeAliasDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAliasDomains
// Query the list of alias domain names.
//
// This feature is only supported in the enterprise plan and is currently in beta test. To use it, please [contact us](https://www.tencentcloud.com/contact-us).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeAliasDomains(request *DescribeAliasDomainsRequest) (response *DescribeAliasDomainsResponse, err error) {
    return c.DescribeAliasDomainsWithContext(context.Background(), request)
}

// DescribeAliasDomains
// Query the list of alias domain names.
//
// This feature is only supported in the enterprise plan and is currently in beta test. To use it, please [contact us](https://www.tencentcloud.com/contact-us).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeAliasDomainsWithContext(ctx context.Context, request *DescribeAliasDomainsRequest) (response *DescribeAliasDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeAliasDomainsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeAliasDomains")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAliasDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAliasDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationProxiesRequest() (request *DescribeApplicationProxiesRequest) {
    request = &DescribeApplicationProxiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeApplicationProxies")
    
    
    return
}

func NewDescribeApplicationProxiesResponse() (response *DescribeApplicationProxiesResponse) {
    response = &DescribeApplicationProxiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApplicationProxies
// This API is on an earlier version. If you want to call it, switch to the latest version. The latest version splits into two APIs: one for querying the Layer 4 proxy instance list and the other for querying Layer 4 forwarding rules. For details, see [Query Layer 4 Proxy Instance List](https://www.tencentcloud.com/document/product/1552/103413?from_cn_redirect=1) and [Query Layer 4 Proxy Forwarding Rule List](https://www.tencentcloud.com/document/product/1552/103412?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeApplicationProxies(request *DescribeApplicationProxiesRequest) (response *DescribeApplicationProxiesResponse, err error) {
    return c.DescribeApplicationProxiesWithContext(context.Background(), request)
}

// DescribeApplicationProxies
// This API is on an earlier version. If you want to call it, switch to the latest version. The latest version splits into two APIs: one for querying the Layer 4 proxy instance list and the other for querying Layer 4 forwarding rules. For details, see [Query Layer 4 Proxy Instance List](https://www.tencentcloud.com/document/product/1552/103413?from_cn_redirect=1) and [Query Layer 4 Proxy Forwarding Rule List](https://www.tencentcloud.com/document/product/1552/103412?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeApplicationProxiesWithContext(ctx context.Context, request *DescribeApplicationProxiesRequest) (response *DescribeApplicationProxiesResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationProxiesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeApplicationProxies")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationProxies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationProxiesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAvailableCustomActionsForRuleEngineRequest() (request *DescribeAvailableCustomActionsForRuleEngineRequest) {
    request = &DescribeAvailableCustomActionsForRuleEngineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeAvailableCustomActionsForRuleEngine")
    
    
    return
}

func NewDescribeAvailableCustomActionsForRuleEngineResponse() (response *DescribeAvailableCustomActionsForRuleEngineResponse) {
    response = &DescribeAvailableCustomActionsForRuleEngineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAvailableCustomActionsForRuleEngine
// Queries the list of available custom configurations for site rule engine operations.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_ZONENOTFOUND = "InvalidParameterValue.ZoneNotFound"
func (c *Client) DescribeAvailableCustomActionsForRuleEngine(request *DescribeAvailableCustomActionsForRuleEngineRequest) (response *DescribeAvailableCustomActionsForRuleEngineResponse, err error) {
    return c.DescribeAvailableCustomActionsForRuleEngineWithContext(context.Background(), request)
}

// DescribeAvailableCustomActionsForRuleEngine
// Queries the list of available custom configurations for site rule engine operations.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_ZONENOTFOUND = "InvalidParameterValue.ZoneNotFound"
func (c *Client) DescribeAvailableCustomActionsForRuleEngineWithContext(ctx context.Context, request *DescribeAvailableCustomActionsForRuleEngineRequest) (response *DescribeAvailableCustomActionsForRuleEngineResponse, err error) {
    if request == nil {
        request = NewDescribeAvailableCustomActionsForRuleEngineRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeAvailableCustomActionsForRuleEngine")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAvailableCustomActionsForRuleEngine require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAvailableCustomActionsForRuleEngineResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAvailableOriginACLFamilyRequest() (request *DescribeAvailableOriginACLFamilyRequest) {
    request = &DescribeAvailableOriginACLFamilyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeAvailableOriginACLFamily")
    
    
    return
}

func NewDescribeAvailableOriginACLFamilyResponse() (response *DescribeAvailableOriginACLFamilyResponse) {
    response = &DescribeAvailableOriginACLFamilyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAvailableOriginACLFamily
// This API is used to query detailed information about the origin protection IP range control domain, including the version number and specific IP range information.
//
// The main difference between standard control domain and streamlined control domain is the number of IP ranges provided. The latter has fewer IP ranges, but has some usage restrictions. For specific restrictions, consult the product. The format description is as follows:
//
// Standard control domain:
//
// <li>gaz: standard global control domain;</li>
//
// <li>mlc: Standard China control domain;</li>
//
// <li>emc: standard overseas (worldwide excluding China) control domain;</li>
//
// Streamline the control domain.
//
// <li>plat-gaz: simplified global control domain;</li>
//
// <li>plat-mlc: Simplified China control domain;</li>
//
// <li>plat-emc: Simplified overseas (global excluding China) control domain;</li>
//
// <li>plat-specific-gaz: custom edition global availability zone control domain;</li>
//
// <li>plat-specific-mlc: custom edition for controlling the Chinese mainland availability zone control domain;</li>
//
// <li>plat-specific-emc: custom edition control domain for availability zones worldwide (excluding Chinese mainland);</li>
//
// Abbreviation description:
//
// <li>gaz:Global AZ Availability Zone;</li>
//
// <li>mlc:mainlandChina;</li>
//
// <li>emc:Exclude mainlandChina.</li>
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_ZONENOTFOUND = "InvalidParameterValue.ZoneNotFound"
func (c *Client) DescribeAvailableOriginACLFamily(request *DescribeAvailableOriginACLFamilyRequest) (response *DescribeAvailableOriginACLFamilyResponse, err error) {
    return c.DescribeAvailableOriginACLFamilyWithContext(context.Background(), request)
}

// DescribeAvailableOriginACLFamily
// This API is used to query detailed information about the origin protection IP range control domain, including the version number and specific IP range information.
//
// The main difference between standard control domain and streamlined control domain is the number of IP ranges provided. The latter has fewer IP ranges, but has some usage restrictions. For specific restrictions, consult the product. The format description is as follows:
//
// Standard control domain:
//
// <li>gaz: standard global control domain;</li>
//
// <li>mlc: Standard China control domain;</li>
//
// <li>emc: standard overseas (worldwide excluding China) control domain;</li>
//
// Streamline the control domain.
//
// <li>plat-gaz: simplified global control domain;</li>
//
// <li>plat-mlc: Simplified China control domain;</li>
//
// <li>plat-emc: Simplified overseas (global excluding China) control domain;</li>
//
// <li>plat-specific-gaz: custom edition global availability zone control domain;</li>
//
// <li>plat-specific-mlc: custom edition for controlling the Chinese mainland availability zone control domain;</li>
//
// <li>plat-specific-emc: custom edition control domain for availability zones worldwide (excluding Chinese mainland);</li>
//
// Abbreviation description:
//
// <li>gaz:Global AZ Availability Zone;</li>
//
// <li>mlc:mainlandChina;</li>
//
// <li>emc:Exclude mainlandChina.</li>
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_ZONENOTFOUND = "InvalidParameterValue.ZoneNotFound"
func (c *Client) DescribeAvailableOriginACLFamilyWithContext(ctx context.Context, request *DescribeAvailableOriginACLFamilyRequest) (response *DescribeAvailableOriginACLFamilyResponse, err error) {
    if request == nil {
        request = NewDescribeAvailableOriginACLFamilyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeAvailableOriginACLFamily")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAvailableOriginACLFamily require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAvailableOriginACLFamilyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAvailablePlansRequest() (request *DescribeAvailablePlansRequest) {
    request = &DescribeAvailablePlansRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeAvailablePlans")
    
    
    return
}

func NewDescribeAvailablePlansResponse() (response *DescribeAvailablePlansResponse) {
    response = &DescribeAvailablePlansResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAvailablePlans
// Query the list of available package information for the current account.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
func (c *Client) DescribeAvailablePlans(request *DescribeAvailablePlansRequest) (response *DescribeAvailablePlansResponse, err error) {
    return c.DescribeAvailablePlansWithContext(context.Background(), request)
}

// DescribeAvailablePlans
// Query the list of available package information for the current account.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
func (c *Client) DescribeAvailablePlansWithContext(ctx context.Context, request *DescribeAvailablePlansRequest) (response *DescribeAvailablePlansResponse, err error) {
    if request == nil {
        request = NewDescribeAvailablePlansRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeAvailablePlans")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAvailablePlans require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAvailablePlansResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillingDataRequest() (request *DescribeBillingDataRequest) {
    request = &DescribeBillingDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeBillingData")
    
    
    return
}

func NewDescribeBillingDataResponse() (response *DescribeBillingDataResponse) {
    response = &DescribeBillingDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillingData
// Query billing data via this API.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_GROUPBYLIMITEXCEEDED = "InvalidParameter.GroupByLimitExceeded"
//  INVALIDPARAMETER_INVALIDINTERVAL = "InvalidParameter.InvalidInterval"
//  INVALIDPARAMETER_INVALIDMETRIC = "InvalidParameter.InvalidMetric"
//  INVALIDPARAMETER_ZONEHASNOTBEENBOUNDTOPLAN = "InvalidParameter.ZoneHasNotBeenBoundToPlan"
func (c *Client) DescribeBillingData(request *DescribeBillingDataRequest) (response *DescribeBillingDataResponse, err error) {
    return c.DescribeBillingDataWithContext(context.Background(), request)
}

// DescribeBillingData
// Query billing data via this API.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_GROUPBYLIMITEXCEEDED = "InvalidParameter.GroupByLimitExceeded"
//  INVALIDPARAMETER_INVALIDINTERVAL = "InvalidParameter.InvalidInterval"
//  INVALIDPARAMETER_INVALIDMETRIC = "InvalidParameter.InvalidMetric"
//  INVALIDPARAMETER_ZONEHASNOTBEENBOUNDTOPLAN = "InvalidParameter.ZoneHasNotBeenBoundToPlan"
func (c *Client) DescribeBillingDataWithContext(ctx context.Context, request *DescribeBillingDataRequest) (response *DescribeBillingDataResponse, err error) {
    if request == nil {
        request = NewDescribeBillingDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeBillingData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillingData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillingDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigGroupVersionDetailRequest() (request *DescribeConfigGroupVersionDetailRequest) {
    request = &DescribeConfigGroupVersionDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeConfigGroupVersionDetail")
    
    
    return
}

func NewDescribeConfigGroupVersionDetailResponse() (response *DescribeConfigGroupVersionDetailResponse) {
    response = &DescribeConfigGroupVersionDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConfigGroupVersionDetail
// In version management mode, this API is used to obtain detailed version information, including version ID, description, status, creation time, associated configuration group information, and the content of the version configuration file. The version management feature is in beta test and currently only available to the allowlist.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeConfigGroupVersionDetail(request *DescribeConfigGroupVersionDetailRequest) (response *DescribeConfigGroupVersionDetailResponse, err error) {
    return c.DescribeConfigGroupVersionDetailWithContext(context.Background(), request)
}

// DescribeConfigGroupVersionDetail
// In version management mode, this API is used to obtain detailed version information, including version ID, description, status, creation time, associated configuration group information, and the content of the version configuration file. The version management feature is in beta test and currently only available to the allowlist.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeConfigGroupVersionDetailWithContext(ctx context.Context, request *DescribeConfigGroupVersionDetailRequest) (response *DescribeConfigGroupVersionDetailResponse, err error) {
    if request == nil {
        request = NewDescribeConfigGroupVersionDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeConfigGroupVersionDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConfigGroupVersionDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConfigGroupVersionDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigGroupVersionsRequest() (request *DescribeConfigGroupVersionsRequest) {
    request = &DescribeConfigGroupVersionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeConfigGroupVersions")
    
    
    return
}

func NewDescribeConfigGroupVersionsResponse() (response *DescribeConfigGroupVersionsResponse) {
    response = &DescribeConfigGroupVersionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConfigGroupVersions
// In version management mode, this API is used to query the version list of a specified configuration group. The version management feature is in beta test and currently only available to the allowlist.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeConfigGroupVersions(request *DescribeConfigGroupVersionsRequest) (response *DescribeConfigGroupVersionsResponse, err error) {
    return c.DescribeConfigGroupVersionsWithContext(context.Background(), request)
}

// DescribeConfigGroupVersions
// In version management mode, this API is used to query the version list of a specified configuration group. The version management feature is in beta test and currently only available to the allowlist.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeConfigGroupVersionsWithContext(ctx context.Context, request *DescribeConfigGroupVersionsRequest) (response *DescribeConfigGroupVersionsResponse, err error) {
    if request == nil {
        request = NewDescribeConfigGroupVersionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeConfigGroupVersions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConfigGroupVersions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConfigGroupVersionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContentIdentifiersRequest() (request *DescribeContentIdentifiersRequest) {
    request = &DescribeContentIdentifiersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeContentIdentifiers")
    
    
    return
}

func NewDescribeContentIdentifiersResponse() (response *DescribeContentIdentifiersResponse) {
    response = &DescribeContentIdentifiersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeContentIdentifiers
// This API is used to query content identifiers in batches. They can be filtered by ID, description, status, or tags. Deleted content identifiers queried by status are only retained for three months. This feature is only open in the allowlist.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeContentIdentifiers(request *DescribeContentIdentifiersRequest) (response *DescribeContentIdentifiersResponse, err error) {
    return c.DescribeContentIdentifiersWithContext(context.Background(), request)
}

// DescribeContentIdentifiers
// This API is used to query content identifiers in batches. They can be filtered by ID, description, status, or tags. Deleted content identifiers queried by status are only retained for three months. This feature is only open in the allowlist.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeContentIdentifiersWithContext(ctx context.Context, request *DescribeContentIdentifiersRequest) (response *DescribeContentIdentifiersResponse, err error) {
    if request == nil {
        request = NewDescribeContentIdentifiersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeContentIdentifiers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeContentIdentifiers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeContentIdentifiersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContentQuotaRequest() (request *DescribeContentQuotaRequest) {
    request = &DescribeContentQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeContentQuota")
    
    
    return
}

func NewDescribeContentQuotaResponse() (response *DescribeContentQuotaResponse) {
    response = &DescribeContentQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeContentQuota
// Query the content management API quota
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeContentQuota(request *DescribeContentQuotaRequest) (response *DescribeContentQuotaResponse, err error) {
    return c.DescribeContentQuotaWithContext(context.Background(), request)
}

// DescribeContentQuota
// Query the content management API quota
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeContentQuotaWithContext(ctx context.Context, request *DescribeContentQuotaRequest) (response *DescribeContentQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeContentQuotaRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeContentQuota")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeContentQuota require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeContentQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomErrorPagesRequest() (request *DescribeCustomErrorPagesRequest) {
    request = &DescribeCustomErrorPagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeCustomErrorPages")
    
    
    return
}

func NewDescribeCustomErrorPagesResponse() (response *DescribeCustomErrorPagesResponse) {
    response = &DescribeCustomErrorPagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCustomErrorPages
// This API is used to query the custom error page list.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
func (c *Client) DescribeCustomErrorPages(request *DescribeCustomErrorPagesRequest) (response *DescribeCustomErrorPagesResponse, err error) {
    return c.DescribeCustomErrorPagesWithContext(context.Background(), request)
}

// DescribeCustomErrorPages
// This API is used to query the custom error page list.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
func (c *Client) DescribeCustomErrorPagesWithContext(ctx context.Context, request *DescribeCustomErrorPagesRequest) (response *DescribeCustomErrorPagesResponse, err error) {
    if request == nil {
        request = NewDescribeCustomErrorPagesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeCustomErrorPages")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCustomErrorPages require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCustomErrorPagesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomGroupRequest() (request *DescribeCustomGroupRequest) {
    request = &DescribeCustomGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeCustomGroup")
    
    
    return
}

func NewDescribeCustomGroupResponse() (response *DescribeCustomGroupResponse) {
    response = &DescribeCustomGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCustomGroup
// This API is used to query the detailed information of a custom group, including all entries in the group (returned with pagination via Limit/Offset). The DescribeCustomGroups list API returns up to the first 500 entries for each group. To obtain all entries, use this API. The GroupId can be obtained from the response of CreateCustomGroup or by querying DescribeCustomGroups.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
func (c *Client) DescribeCustomGroup(request *DescribeCustomGroupRequest) (response *DescribeCustomGroupResponse, err error) {
    return c.DescribeCustomGroupWithContext(context.Background(), request)
}

// DescribeCustomGroup
// This API is used to query the detailed information of a custom group, including all entries in the group (returned with pagination via Limit/Offset). The DescribeCustomGroups list API returns up to the first 500 entries for each group. To obtain all entries, use this API. The GroupId can be obtained from the response of CreateCustomGroup or by querying DescribeCustomGroups.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
func (c *Client) DescribeCustomGroupWithContext(ctx context.Context, request *DescribeCustomGroupRequest) (response *DescribeCustomGroupResponse, err error) {
    if request == nil {
        request = NewDescribeCustomGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeCustomGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCustomGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCustomGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomGroupReferencesRequest() (request *DescribeCustomGroupReferencesRequest) {
    request = &DescribeCustomGroupReferencesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeCustomGroupReferences")
    
    
    return
}

func NewDescribeCustomGroupReferencesResponse() (response *DescribeCustomGroupReferencesResponse) {
    response = &DescribeCustomGroupReferencesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCustomGroupReferences
// This API is used to obtain the security policy configuration that uses a custom group.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
func (c *Client) DescribeCustomGroupReferences(request *DescribeCustomGroupReferencesRequest) (response *DescribeCustomGroupReferencesResponse, err error) {
    return c.DescribeCustomGroupReferencesWithContext(context.Background(), request)
}

// DescribeCustomGroupReferences
// This API is used to obtain the security policy configuration that uses a custom group.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
func (c *Client) DescribeCustomGroupReferencesWithContext(ctx context.Context, request *DescribeCustomGroupReferencesRequest) (response *DescribeCustomGroupReferencesResponse, err error) {
    if request == nil {
        request = NewDescribeCustomGroupReferencesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeCustomGroupReferences")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCustomGroupReferences require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCustomGroupReferencesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomGroupsRequest() (request *DescribeCustomGroupsRequest) {
    request = &DescribeCustomGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeCustomGroups")
    
    
    return
}

func NewDescribeCustomGroupsResponse() (response *DescribeCustomGroupsResponse) {
    response = &DescribeCustomGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCustomGroups
// This API is used to query the list of custom groups under a specified site. This API returns only the first 500 entries of each custom group. To query all entries in a group, use the DescribeSecurityCustomGroup API.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
func (c *Client) DescribeCustomGroups(request *DescribeCustomGroupsRequest) (response *DescribeCustomGroupsResponse, err error) {
    return c.DescribeCustomGroupsWithContext(context.Background(), request)
}

// DescribeCustomGroups
// This API is used to query the list of custom groups under a specified site. This API returns only the first 500 entries of each custom group. To query all entries in a group, use the DescribeSecurityCustomGroup API.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
func (c *Client) DescribeCustomGroupsWithContext(ctx context.Context, request *DescribeCustomGroupsRequest) (response *DescribeCustomGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeCustomGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeCustomGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCustomGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCustomGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSAttackDataRequest() (request *DescribeDDoSAttackDataRequest) {
    request = &DescribeDDoSAttackDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeDDoSAttackData")
    
    
    return
}

func NewDescribeDDoSAttackDataResponse() (response *DescribeDDoSAttackDataResponse) {
    response = &DescribeDDoSAttackDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDDoSAttackData
// This API is used to query DDoS attack time-series data.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeDDoSAttackData(request *DescribeDDoSAttackDataRequest) (response *DescribeDDoSAttackDataResponse, err error) {
    return c.DescribeDDoSAttackDataWithContext(context.Background(), request)
}

// DescribeDDoSAttackData
// This API is used to query DDoS attack time-series data.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeDDoSAttackDataWithContext(ctx context.Context, request *DescribeDDoSAttackDataRequest) (response *DescribeDDoSAttackDataResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSAttackDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeDDoSAttackData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDDoSAttackData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDDoSAttackDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSAttackEventRequest() (request *DescribeDDoSAttackEventRequest) {
    request = &DescribeDDoSAttackEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeDDoSAttackEvent")
    
    
    return
}

func NewDescribeDDoSAttackEventResponse() (response *DescribeDDoSAttackEventResponse) {
    response = &DescribeDDoSAttackEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDDoSAttackEvent
// This API is used to query the list of DDoS attack events.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeDDoSAttackEvent(request *DescribeDDoSAttackEventRequest) (response *DescribeDDoSAttackEventResponse, err error) {
    return c.DescribeDDoSAttackEventWithContext(context.Background(), request)
}

// DescribeDDoSAttackEvent
// This API is used to query the list of DDoS attack events.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeDDoSAttackEventWithContext(ctx context.Context, request *DescribeDDoSAttackEventRequest) (response *DescribeDDoSAttackEventResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSAttackEventRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeDDoSAttackEvent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDDoSAttackEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDDoSAttackEventResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSAttackTopDataRequest() (request *DescribeDDoSAttackTopDataRequest) {
    request = &DescribeDDoSAttackTopDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeDDoSAttackTopData")
    
    
    return
}

func NewDescribeDDoSAttackTopDataResponse() (response *DescribeDDoSAttackTopDataResponse) {
    response = &DescribeDDoSAttackTopDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDDoSAttackTopData
// This API is used to query Top-ranked data of DDoS attacks.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeDDoSAttackTopData(request *DescribeDDoSAttackTopDataRequest) (response *DescribeDDoSAttackTopDataResponse, err error) {
    return c.DescribeDDoSAttackTopDataWithContext(context.Background(), request)
}

// DescribeDDoSAttackTopData
// This API is used to query Top-ranked data of DDoS attacks.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeDDoSAttackTopDataWithContext(ctx context.Context, request *DescribeDDoSAttackTopDataRequest) (response *DescribeDDoSAttackTopDataResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSAttackTopDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeDDoSAttackTopData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDDoSAttackTopData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDDoSAttackTopDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSProtectionRequest() (request *DescribeDDoSProtectionRequest) {
    request = &DescribeDDoSProtectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeDDoSProtection")
    
    
    return
}

func NewDescribeDDoSProtectionResponse() (response *DescribeDDoSProtectionResponse) {
    response = &DescribeDDoSProtectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDDoSProtection
// This API is used to query exclusive DDoS protection info of a site.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeDDoSProtection(request *DescribeDDoSProtectionRequest) (response *DescribeDDoSProtectionResponse, err error) {
    return c.DescribeDDoSProtectionWithContext(context.Background(), request)
}

// DescribeDDoSProtection
// This API is used to query exclusive DDoS protection info of a site.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeDDoSProtectionWithContext(ctx context.Context, request *DescribeDDoSProtectionRequest) (response *DescribeDDoSProtectionResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSProtectionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeDDoSProtection")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDDoSProtection require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDDoSProtectionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDefaultCertificatesRequest() (request *DescribeDefaultCertificatesRequest) {
    request = &DescribeDefaultCertificatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeDefaultCertificates")
    
    
    return
}

func NewDescribeDefaultCertificatesResponse() (response *DescribeDefaultCertificatesResponse) {
    response = &DescribeDefaultCertificatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDefaultCertificates
// Query the default certificate list
//
// <p style="color: red;">Note: This earlier version interface is no longer updated. New features will only be provided in the latest version interface. Existing capacities supported by the earlier version interface will not be affected. To query the cert configured for a domain name, please use DescribeAccelerationDomains.</p>
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeDefaultCertificates(request *DescribeDefaultCertificatesRequest) (response *DescribeDefaultCertificatesResponse, err error) {
    return c.DescribeDefaultCertificatesWithContext(context.Background(), request)
}

// DescribeDefaultCertificates
// Query the default certificate list
//
// <p style="color: red;">Note: This earlier version interface is no longer updated. New features will only be provided in the latest version interface. Existing capacities supported by the earlier version interface will not be affected. To query the cert configured for a domain name, please use DescribeAccelerationDomains.</p>
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeDefaultCertificatesWithContext(ctx context.Context, request *DescribeDefaultCertificatesRequest) (response *DescribeDefaultCertificatesResponse, err error) {
    if request == nil {
        request = NewDescribeDefaultCertificatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeDefaultCertificates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDefaultCertificates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDefaultCertificatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeployHistoryRequest() (request *DescribeDeployHistoryRequest) {
    request = &DescribeDeployHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeDeployHistory")
    
    
    return
}

func NewDescribeDeployHistoryResponse() (response *DescribeDeployHistoryResponse) {
    response = &DescribeDeployHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeployHistory
// In version management mode, this API is used to query the version release history of production/test environments. The version management feature is in beta test and currently only available to the allowlist.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDeployHistory(request *DescribeDeployHistoryRequest) (response *DescribeDeployHistoryResponse, err error) {
    return c.DescribeDeployHistoryWithContext(context.Background(), request)
}

// DescribeDeployHistory
// In version management mode, this API is used to query the version release history of production/test environments. The version management feature is in beta test and currently only available to the allowlist.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDeployHistoryWithContext(ctx context.Context, request *DescribeDeployHistoryRequest) (response *DescribeDeployHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeDeployHistoryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeDeployHistory")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeployHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeployHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDnsRecordsRequest() (request *DescribeDnsRecordsRequest) {
    request = &DescribeDnsRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeDnsRecords")
    
    
    return
}

func NewDescribeDnsRecordsResponse() (response *DescribeDnsRecordsResponse) {
    response = &DescribeDnsRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDnsRecords
// This API is used to view DNS records under a site, including DNS record names, record types, and record content. It supports specifying filter conditions to query specific DNS records.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDnsRecords(request *DescribeDnsRecordsRequest) (response *DescribeDnsRecordsResponse, err error) {
    return c.DescribeDnsRecordsWithContext(context.Background(), request)
}

// DescribeDnsRecords
// This API is used to view DNS records under a site, including DNS record names, record types, and record content. It supports specifying filter conditions to query specific DNS records.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDnsRecordsWithContext(ctx context.Context, request *DescribeDnsRecordsRequest) (response *DescribeDnsRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeDnsRecordsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeDnsRecords")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDnsRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDnsRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeKVNamespacesRequest() (request *DescribeEdgeKVNamespacesRequest) {
    request = &DescribeEdgeKVNamespacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeEdgeKVNamespaces")
    
    
    return
}

func NewDescribeEdgeKVNamespacesResponse() (response *DescribeEdgeKVNamespacesResponse) {
    response = &DescribeEdgeKVNamespacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEdgeKVNamespaces
// Queries the KV namespace list under a specified site. Supports pagination, sorting, and conditional filtering. Returns basic namespace info, storage capacity utilization, and reference relationships. If no data is found, an empty array is returned.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDFILTERNAME = "InvalidParameter.InvalidFilterName"
//  INVALIDPARAMETER_INVALIDSORTBY = "InvalidParameter.InvalidSortBy"
//  INVALIDPARAMETER_INVALIDSORTORDER = "InvalidParameter.InvalidSortOrder"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_NAMESPACELIMITEXCEEDED = "LimitExceeded.NamespaceLimitExceeded"
func (c *Client) DescribeEdgeKVNamespaces(request *DescribeEdgeKVNamespacesRequest) (response *DescribeEdgeKVNamespacesResponse, err error) {
    return c.DescribeEdgeKVNamespacesWithContext(context.Background(), request)
}

// DescribeEdgeKVNamespaces
// Queries the KV namespace list under a specified site. Supports pagination, sorting, and conditional filtering. Returns basic namespace info, storage capacity utilization, and reference relationships. If no data is found, an empty array is returned.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDFILTERNAME = "InvalidParameter.InvalidFilterName"
//  INVALIDPARAMETER_INVALIDSORTBY = "InvalidParameter.InvalidSortBy"
//  INVALIDPARAMETER_INVALIDSORTORDER = "InvalidParameter.InvalidSortOrder"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_NAMESPACELIMITEXCEEDED = "LimitExceeded.NamespaceLimitExceeded"
func (c *Client) DescribeEdgeKVNamespacesWithContext(ctx context.Context, request *DescribeEdgeKVNamespacesRequest) (response *DescribeEdgeKVNamespacesResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeKVNamespacesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeEdgeKVNamespaces")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeKVNamespaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeKVNamespacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnvironmentsRequest() (request *DescribeEnvironmentsRequest) {
    request = &DescribeEnvironmentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeEnvironments")
    
    
    return
}

func NewDescribeEnvironmentsResponse() (response *DescribeEnvironmentsResponse) {
    response = &DescribeEnvironmentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEnvironments
// In version management mode, this API is used to query environment information, including the environment ID, type, and current effective version. The version management feature is in beta test and currently only available to the allowlist.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDFILTERNAME = "InvalidParameter.InvalidFilterName"
//  INVALIDPARAMETER_INVALIDSORTBY = "InvalidParameter.InvalidSortBy"
//  INVALIDPARAMETER_INVALIDSORTORDER = "InvalidParameter.InvalidSortOrder"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_NAMESPACELIMITEXCEEDED = "LimitExceeded.NamespaceLimitExceeded"
func (c *Client) DescribeEnvironments(request *DescribeEnvironmentsRequest) (response *DescribeEnvironmentsResponse, err error) {
    return c.DescribeEnvironmentsWithContext(context.Background(), request)
}

// DescribeEnvironments
// In version management mode, this API is used to query environment information, including the environment ID, type, and current effective version. The version management feature is in beta test and currently only available to the allowlist.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDFILTERNAME = "InvalidParameter.InvalidFilterName"
//  INVALIDPARAMETER_INVALIDSORTBY = "InvalidParameter.InvalidSortBy"
//  INVALIDPARAMETER_INVALIDSORTORDER = "InvalidParameter.InvalidSortOrder"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_NAMESPACELIMITEXCEEDED = "LimitExceeded.NamespaceLimitExceeded"
func (c *Client) DescribeEnvironmentsWithContext(ctx context.Context, request *DescribeEnvironmentsRequest) (response *DescribeEnvironmentsResponse, err error) {
    if request == nil {
        request = NewDescribeEnvironmentsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeEnvironments")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEnvironments require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEnvironmentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFunctionComponentBindingsRequest() (request *DescribeFunctionComponentBindingsRequest) {
    request = &DescribeFunctionComponentBindingsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeFunctionComponentBindings")
    
    
    return
}

func NewDescribeFunctionComponentBindingsResponse() (response *DescribeFunctionComponentBindingsResponse) {
    response = &DescribeFunctionComponentBindingsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFunctionComponentBindings
// This API is used to query the component binding list of a specified edge function. It supports pagination and conditional filtering, and returns detailed information such as the bound component type, variable name, and configuration parameters. The currently supported binding component type is KV namespace (kv_namespace).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDFILTERNAME = "InvalidParameter.InvalidFilterName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeFunctionComponentBindings(request *DescribeFunctionComponentBindingsRequest) (response *DescribeFunctionComponentBindingsResponse, err error) {
    return c.DescribeFunctionComponentBindingsWithContext(context.Background(), request)
}

// DescribeFunctionComponentBindings
// This API is used to query the component binding list of a specified edge function. It supports pagination and conditional filtering, and returns detailed information such as the bound component type, variable name, and configuration parameters. The currently supported binding component type is KV namespace (kv_namespace).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDFILTERNAME = "InvalidParameter.InvalidFilterName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeFunctionComponentBindingsWithContext(ctx context.Context, request *DescribeFunctionComponentBindingsRequest) (response *DescribeFunctionComponentBindingsResponse, err error) {
    if request == nil {
        request = NewDescribeFunctionComponentBindingsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeFunctionComponentBindings")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFunctionComponentBindings require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFunctionComponentBindingsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFunctionReplicasRequest() (request *DescribeFunctionReplicasRequest) {
    request = &DescribeFunctionReplicasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeFunctionReplicas")
    
    
    return
}

func NewDescribeFunctionReplicasResponse() (response *DescribeFunctionReplicasResponse) {
    response = &DescribeFunctionReplicasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFunctionReplicas
// This API is used to query the replica list of edge functions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDFILTERNAME = "InvalidParameter.InvalidFilterName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeFunctionReplicas(request *DescribeFunctionReplicasRequest) (response *DescribeFunctionReplicasResponse, err error) {
    return c.DescribeFunctionReplicasWithContext(context.Background(), request)
}

// DescribeFunctionReplicas
// This API is used to query the replica list of edge functions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDFILTERNAME = "InvalidParameter.InvalidFilterName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeFunctionReplicasWithContext(ctx context.Context, request *DescribeFunctionReplicasRequest) (response *DescribeFunctionReplicasResponse, err error) {
    if request == nil {
        request = NewDescribeFunctionReplicasRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeFunctionReplicas")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFunctionReplicas require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFunctionReplicasResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFunctionRulesRequest() (request *DescribeFunctionRulesRequest) {
    request = &DescribeFunctionRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeFunctionRules")
    
    
    return
}

func NewDescribeFunctionRulesResponse() (response *DescribeFunctionRulesResponse) {
    response = &DescribeFunctionRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFunctionRules
// This API is used to query the list of trigger rules for an edge function. It supports filtering by rule ID, function ID, rule description, and so on.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INVALIDFILTERNAME = "InvalidParameter.InvalidFilterName"
//  INVALIDPARAMETER_TOOMANYFILTERVALUES = "InvalidParameter.TooManyFilterValues"
//  INVALIDPARAMETER_TOOMANYFILTERS = "InvalidParameter.TooManyFilters"
//  RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"
func (c *Client) DescribeFunctionRules(request *DescribeFunctionRulesRequest) (response *DescribeFunctionRulesResponse, err error) {
    return c.DescribeFunctionRulesWithContext(context.Background(), request)
}

// DescribeFunctionRules
// This API is used to query the list of trigger rules for an edge function. It supports filtering by rule ID, function ID, rule description, and so on.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INVALIDFILTERNAME = "InvalidParameter.InvalidFilterName"
//  INVALIDPARAMETER_TOOMANYFILTERVALUES = "InvalidParameter.TooManyFilterValues"
//  INVALIDPARAMETER_TOOMANYFILTERS = "InvalidParameter.TooManyFilters"
//  RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"
func (c *Client) DescribeFunctionRulesWithContext(ctx context.Context, request *DescribeFunctionRulesRequest) (response *DescribeFunctionRulesResponse, err error) {
    if request == nil {
        request = NewDescribeFunctionRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeFunctionRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFunctionRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFunctionRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFunctionRuntimeEnvironmentRequest() (request *DescribeFunctionRuntimeEnvironmentRequest) {
    request = &DescribeFunctionRuntimeEnvironmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeFunctionRuntimeEnvironment")
    
    
    return
}

func NewDescribeFunctionRuntimeEnvironmentResponse() (response *DescribeFunctionRuntimeEnvironmentResponse) {
    response = &DescribeFunctionRuntimeEnvironmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFunctionRuntimeEnvironment
// This API is used to query the runtime environment of an edge function, including environment variables.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INVALIDFILTERNAME = "InvalidParameter.InvalidFilterName"
//  RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"
func (c *Client) DescribeFunctionRuntimeEnvironment(request *DescribeFunctionRuntimeEnvironmentRequest) (response *DescribeFunctionRuntimeEnvironmentResponse, err error) {
    return c.DescribeFunctionRuntimeEnvironmentWithContext(context.Background(), request)
}

// DescribeFunctionRuntimeEnvironment
// This API is used to query the runtime environment of an edge function, including environment variables.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INVALIDFILTERNAME = "InvalidParameter.InvalidFilterName"
//  RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"
func (c *Client) DescribeFunctionRuntimeEnvironmentWithContext(ctx context.Context, request *DescribeFunctionRuntimeEnvironmentRequest) (response *DescribeFunctionRuntimeEnvironmentResponse, err error) {
    if request == nil {
        request = NewDescribeFunctionRuntimeEnvironmentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeFunctionRuntimeEnvironment")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFunctionRuntimeEnvironment require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFunctionRuntimeEnvironmentResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFunctionsRequest() (request *DescribeFunctionsRequest) {
    request = &DescribeFunctionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeFunctions")
    
    
    return
}

func NewDescribeFunctionsResponse() (response *DescribeFunctionsResponse) {
    response = &DescribeFunctionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFunctions
// This API is used to query the list of edge functions. It supports filtering by function ID, name, description, and so on.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INVALIDFILTERNAME = "InvalidParameter.InvalidFilterName"
//  INVALIDPARAMETER_LENGTHEXCEEDSLIMIT = "InvalidParameter.LengthExceedsLimit"
//  INVALIDPARAMETER_TOOMANYFILTERVALUES = "InvalidParameter.TooManyFilterValues"
//  INVALIDPARAMETER_TOOMANYFILTERS = "InvalidParameter.TooManyFilters"
//  RESOURCEUNAVAILABLE_FUNCTIONNOTFOUND = "ResourceUnavailable.FunctionNotFound"
//  RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"
func (c *Client) DescribeFunctions(request *DescribeFunctionsRequest) (response *DescribeFunctionsResponse, err error) {
    return c.DescribeFunctionsWithContext(context.Background(), request)
}

// DescribeFunctions
// This API is used to query the list of edge functions. It supports filtering by function ID, name, description, and so on.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INVALIDFILTERNAME = "InvalidParameter.InvalidFilterName"
//  INVALIDPARAMETER_LENGTHEXCEEDSLIMIT = "InvalidParameter.LengthExceedsLimit"
//  INVALIDPARAMETER_TOOMANYFILTERVALUES = "InvalidParameter.TooManyFilterValues"
//  INVALIDPARAMETER_TOOMANYFILTERS = "InvalidParameter.TooManyFilters"
//  RESOURCEUNAVAILABLE_FUNCTIONNOTFOUND = "ResourceUnavailable.FunctionNotFound"
//  RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"
func (c *Client) DescribeFunctionsWithContext(ctx context.Context, request *DescribeFunctionsRequest) (response *DescribeFunctionsResponse, err error) {
    if request == nil {
        request = NewDescribeFunctionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeFunctions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFunctions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFunctionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostsSettingRequest() (request *DescribeHostsSettingRequest) {
    request = &DescribeHostsSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeHostsSetting")
    
    
    return
}

func NewDescribeHostsSettingResponse() (response *DescribeHostsSettingResponse) {
    response = &DescribeHostsSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHostsSetting
// This API is on an earlier version. EdgeOne has completed a comprehensive upgrade of rule engine APIs. You can use [DescribeL7AccSetting](https://www.tencentcloud.com/document/product/1552/115819?from_cn_redirect=1) and [DescribeL7AccRules](https://www.tencentcloud.com/document/product/1552/115820?from_cn_redirect=1) to obtain detailed configuration of a domain.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeHostsSetting(request *DescribeHostsSettingRequest) (response *DescribeHostsSettingResponse, err error) {
    return c.DescribeHostsSettingWithContext(context.Background(), request)
}

// DescribeHostsSetting
// This API is on an earlier version. EdgeOne has completed a comprehensive upgrade of rule engine APIs. You can use [DescribeL7AccSetting](https://www.tencentcloud.com/document/product/1552/115819?from_cn_redirect=1) and [DescribeL7AccRules](https://www.tencentcloud.com/document/product/1552/115820?from_cn_redirect=1) to obtain detailed configuration of a domain.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeHostsSettingWithContext(ctx context.Context, request *DescribeHostsSettingRequest) (response *DescribeHostsSettingResponse, err error) {
    if request == nil {
        request = NewDescribeHostsSettingRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeHostsSetting")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHostsSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostsSettingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIPGroupReferencesRequest() (request *DescribeIPGroupReferencesRequest) {
    request = &DescribeIPGroupReferencesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeIPGroupReferences")
    
    
    return
}

func NewDescribeIPGroupReferencesResponse() (response *DescribeIPGroupReferencesResponse) {
    response = &DescribeIPGroupReferencesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIPGroupReferences
// This API is used to obtain policy settings that use IP grouping.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeIPGroupReferences(request *DescribeIPGroupReferencesRequest) (response *DescribeIPGroupReferencesResponse, err error) {
    return c.DescribeIPGroupReferencesWithContext(context.Background(), request)
}

// DescribeIPGroupReferences
// This API is used to obtain policy settings that use IP grouping.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeIPGroupReferencesWithContext(ctx context.Context, request *DescribeIPGroupReferencesRequest) (response *DescribeIPGroupReferencesResponse, err error) {
    if request == nil {
        request = NewDescribeIPGroupReferencesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeIPGroupReferences")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIPGroupReferences require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIPGroupReferencesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIPRegionRequest() (request *DescribeIPRegionRequest) {
    request = &DescribeIPRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeIPRegion")
    
    
    return
}

func NewDescribeIPRegionResponse() (response *DescribeIPRegionResponse) {
    response = &DescribeIPRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIPRegion
// This API is used to query if the IP is an EdgeOne IP.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
func (c *Client) DescribeIPRegion(request *DescribeIPRegionRequest) (response *DescribeIPRegionResponse, err error) {
    return c.DescribeIPRegionWithContext(context.Background(), request)
}

// DescribeIPRegion
// This API is used to query if the IP is an EdgeOne IP.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
func (c *Client) DescribeIPRegionWithContext(ctx context.Context, request *DescribeIPRegionRequest) (response *DescribeIPRegionResponse, err error) {
    if request == nil {
        request = NewDescribeIPRegionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeIPRegion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIPRegion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIPRegionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIdentificationsRequest() (request *DescribeIdentificationsRequest) {
    request = &DescribeIdentificationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeIdentifications")
    
    
    return
}

func NewDescribeIdentificationsResponse() (response *DescribeIdentificationsResponse) {
    response = &DescribeIdentificationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIdentifications
// Queries verification information of a site.
//
// error code that may be returned:
//  OPERATIONDENIED_DOMAININSHARECNAMEGROUP = "OperationDenied.DomainInShareCnameGroup"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeIdentifications(request *DescribeIdentificationsRequest) (response *DescribeIdentificationsResponse, err error) {
    return c.DescribeIdentificationsWithContext(context.Background(), request)
}

// DescribeIdentifications
// Queries verification information of a site.
//
// error code that may be returned:
//  OPERATIONDENIED_DOMAININSHARECNAMEGROUP = "OperationDenied.DomainInShareCnameGroup"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeIdentificationsWithContext(ctx context.Context, request *DescribeIdentificationsRequest) (response *DescribeIdentificationsResponse, err error) {
    if request == nil {
        request = NewDescribeIdentificationsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeIdentifications")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIdentifications require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIdentificationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInferenceAPITokensRequest() (request *DescribeInferenceAPITokensRequest) {
    request = &DescribeInferenceAPITokensRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeInferenceAPITokens")
    
    
    return
}

func NewDescribeInferenceAPITokensResponse() (response *DescribeInferenceAPITokensResponse) {
    response = &DescribeInferenceAPITokensResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInferenceAPITokens
// Queries the list of reasoning API tokens and returns the Token ID, name, content, and creation time. Paging query is supported.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_ZONENOTFOUND = "InvalidParameterValue.ZoneNotFound"
func (c *Client) DescribeInferenceAPITokens(request *DescribeInferenceAPITokensRequest) (response *DescribeInferenceAPITokensResponse, err error) {
    return c.DescribeInferenceAPITokensWithContext(context.Background(), request)
}

// DescribeInferenceAPITokens
// Queries the list of reasoning API tokens and returns the Token ID, name, content, and creation time. Paging query is supported.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_ZONENOTFOUND = "InvalidParameterValue.ZoneNotFound"
func (c *Client) DescribeInferenceAPITokensWithContext(ctx context.Context, request *DescribeInferenceAPITokensRequest) (response *DescribeInferenceAPITokensResponse, err error) {
    if request == nil {
        request = NewDescribeInferenceAPITokensRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeInferenceAPITokens")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInferenceAPITokens require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInferenceAPITokensResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInferenceDomainsRequest() (request *DescribeInferenceDomainsRequest) {
    request = &DescribeInferenceDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeInferenceDomains")
    
    
    return
}

func NewDescribeInferenceDomainsResponse() (response *DescribeInferenceDomainsResponse) {
    response = &DescribeInferenceDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInferenceDomains
// This API is used to query the inference service domain name list and return the Cname address, ownership verification, and certificate information of the domain.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_ZONENOTFOUND = "InvalidParameterValue.ZoneNotFound"
//  RESOURCENOTFOUND_INFERENCESERVICE = "ResourceNotFound.InferenceService"
func (c *Client) DescribeInferenceDomains(request *DescribeInferenceDomainsRequest) (response *DescribeInferenceDomainsResponse, err error) {
    return c.DescribeInferenceDomainsWithContext(context.Background(), request)
}

// DescribeInferenceDomains
// This API is used to query the inference service domain name list and return the Cname address, ownership verification, and certificate information of the domain.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_ZONENOTFOUND = "InvalidParameterValue.ZoneNotFound"
//  RESOURCENOTFOUND_INFERENCESERVICE = "ResourceNotFound.InferenceService"
func (c *Client) DescribeInferenceDomainsWithContext(ctx context.Context, request *DescribeInferenceDomainsRequest) (response *DescribeInferenceDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeInferenceDomainsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeInferenceDomains")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInferenceDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInferenceDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInferenceHardwareSpecificationsRequest() (request *DescribeInferenceHardwareSpecificationsRequest) {
    request = &DescribeInferenceHardwareSpecificationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeInferenceHardwareSpecifications")
    
    
    return
}

func NewDescribeInferenceHardwareSpecificationsResponse() (response *DescribeInferenceHardwareSpecificationsResponse) {
    response = &DescribeInferenceHardwareSpecificationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInferenceHardwareSpecifications
// This API is used to query the inference hardware specification list and return the CPU, memory, GPU, and video memory configurations of each specification. You can select the required specification from the list when creating a service.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_ZONENOTFOUND = "InvalidParameterValue.ZoneNotFound"
func (c *Client) DescribeInferenceHardwareSpecifications(request *DescribeInferenceHardwareSpecificationsRequest) (response *DescribeInferenceHardwareSpecificationsResponse, err error) {
    return c.DescribeInferenceHardwareSpecificationsWithContext(context.Background(), request)
}

// DescribeInferenceHardwareSpecifications
// This API is used to query the inference hardware specification list and return the CPU, memory, GPU, and video memory configurations of each specification. You can select the required specification from the list when creating a service.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_ZONENOTFOUND = "InvalidParameterValue.ZoneNotFound"
func (c *Client) DescribeInferenceHardwareSpecificationsWithContext(ctx context.Context, request *DescribeInferenceHardwareSpecificationsRequest) (response *DescribeInferenceHardwareSpecificationsResponse, err error) {
    if request == nil {
        request = NewDescribeInferenceHardwareSpecificationsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeInferenceHardwareSpecifications")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInferenceHardwareSpecifications require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInferenceHardwareSpecificationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInferenceServiceDeploymentLogsRequest() (request *DescribeInferenceServiceDeploymentLogsRequest) {
    request = &DescribeInferenceServiceDeploymentLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeInferenceServiceDeploymentLogs")
    
    
    return
}

func NewDescribeInferenceServiceDeploymentLogsResponse() (response *DescribeInferenceServiceDeploymentLogsResponse) {
    response = &DescribeInferenceServiceDeploymentLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInferenceServiceDeploymentLogs
// This API is used to query the logs of a specified deployment of an inference service, return log content and generation time, and support time range search, pagination, and sorting.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ZONENOTFOUND = "InvalidParameterValue.ZoneNotFound"
func (c *Client) DescribeInferenceServiceDeploymentLogs(request *DescribeInferenceServiceDeploymentLogsRequest) (response *DescribeInferenceServiceDeploymentLogsResponse, err error) {
    return c.DescribeInferenceServiceDeploymentLogsWithContext(context.Background(), request)
}

// DescribeInferenceServiceDeploymentLogs
// This API is used to query the logs of a specified deployment of an inference service, return log content and generation time, and support time range search, pagination, and sorting.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ZONENOTFOUND = "InvalidParameterValue.ZoneNotFound"
func (c *Client) DescribeInferenceServiceDeploymentLogsWithContext(ctx context.Context, request *DescribeInferenceServiceDeploymentLogsRequest) (response *DescribeInferenceServiceDeploymentLogsResponse, err error) {
    if request == nil {
        request = NewDescribeInferenceServiceDeploymentLogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeInferenceServiceDeploymentLogs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInferenceServiceDeploymentLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInferenceServiceDeploymentLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInferenceServiceDeploymentRecordsRequest() (request *DescribeInferenceServiceDeploymentRecordsRequest) {
    request = &DescribeInferenceServiceDeploymentRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeInferenceServiceDeploymentRecords")
    
    
    return
}

func NewDescribeInferenceServiceDeploymentRecordsResponse() (response *DescribeInferenceServiceDeploymentRecordsResponse) {
    response = &DescribeInferenceServiceDeploymentRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInferenceServiceDeploymentRecords
// This API is used to query the deployment history list of an inference service. It returns the operation type, status, duration, configuration snapshot, and whether it is the currently effective configuration for each deployment. Pagination and sorting are supported.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ZONENOTFOUND = "InvalidParameterValue.ZoneNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeInferenceServiceDeploymentRecords(request *DescribeInferenceServiceDeploymentRecordsRequest) (response *DescribeInferenceServiceDeploymentRecordsResponse, err error) {
    return c.DescribeInferenceServiceDeploymentRecordsWithContext(context.Background(), request)
}

// DescribeInferenceServiceDeploymentRecords
// This API is used to query the deployment history list of an inference service. It returns the operation type, status, duration, configuration snapshot, and whether it is the currently effective configuration for each deployment. Pagination and sorting are supported.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ZONENOTFOUND = "InvalidParameterValue.ZoneNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeInferenceServiceDeploymentRecordsWithContext(ctx context.Context, request *DescribeInferenceServiceDeploymentRecordsRequest) (response *DescribeInferenceServiceDeploymentRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeInferenceServiceDeploymentRecordsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeInferenceServiceDeploymentRecords")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInferenceServiceDeploymentRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInferenceServiceDeploymentRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInferenceServiceMonitorDataRequest() (request *DescribeInferenceServiceMonitorDataRequest) {
    request = &DescribeInferenceServiceMonitorDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeInferenceServiceMonitorData")
    
    
    return
}

func NewDescribeInferenceServiceMonitorDataResponse() (response *DescribeInferenceServiceMonitorDataResponse) {
    response = &DescribeInferenceServiceMonitorDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInferenceServiceMonitorData
// This API is used to query the monitoring data of inference services, supporting indicators such as CPU, memory, GPU, video memory utilization, and instance count. You can specify the time range and aggregation granularity, and query data from the most recent 30 days.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ZONENOTFOUND = "InvalidParameterValue.ZoneNotFound"
func (c *Client) DescribeInferenceServiceMonitorData(request *DescribeInferenceServiceMonitorDataRequest) (response *DescribeInferenceServiceMonitorDataResponse, err error) {
    return c.DescribeInferenceServiceMonitorDataWithContext(context.Background(), request)
}

// DescribeInferenceServiceMonitorData
// This API is used to query the monitoring data of inference services, supporting indicators such as CPU, memory, GPU, video memory utilization, and instance count. You can specify the time range and aggregation granularity, and query data from the most recent 30 days.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ZONENOTFOUND = "InvalidParameterValue.ZoneNotFound"
func (c *Client) DescribeInferenceServiceMonitorDataWithContext(ctx context.Context, request *DescribeInferenceServiceMonitorDataRequest) (response *DescribeInferenceServiceMonitorDataResponse, err error) {
    if request == nil {
        request = NewDescribeInferenceServiceMonitorDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeInferenceServiceMonitorData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInferenceServiceMonitorData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInferenceServiceMonitorDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInferenceServicesRequest() (request *DescribeInferenceServicesRequest) {
    request = &DescribeInferenceServicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeInferenceServices")
    
    
    return
}

func NewDescribeInferenceServicesResponse() (response *DescribeInferenceServicesResponse) {
    response = &DescribeInferenceServicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInferenceServices
// This API is used to query the inference service list, support filtering by service name, service ID, and status, and return the service configuration, running state, number of instances, and inference access address.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_ZONENOTFOUND = "InvalidParameterValue.ZoneNotFound"
func (c *Client) DescribeInferenceServices(request *DescribeInferenceServicesRequest) (response *DescribeInferenceServicesResponse, err error) {
    return c.DescribeInferenceServicesWithContext(context.Background(), request)
}

// DescribeInferenceServices
// This API is used to query the inference service list, support filtering by service name, service ID, and status, and return the service configuration, running state, number of instances, and inference access address.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_ZONENOTFOUND = "InvalidParameterValue.ZoneNotFound"
func (c *Client) DescribeInferenceServicesWithContext(ctx context.Context, request *DescribeInferenceServicesRequest) (response *DescribeInferenceServicesResponse, err error) {
    if request == nil {
        request = NewDescribeInferenceServicesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeInferenceServices")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInferenceServices require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInferenceServicesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeJustInTimeTranscodeTemplatesRequest() (request *DescribeJustInTimeTranscodeTemplatesRequest) {
    request = &DescribeJustInTimeTranscodeTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeJustInTimeTranscodeTemplates")
    
    
    return
}

func NewDescribeJustInTimeTranscodeTemplatesResponse() (response *DescribeJustInTimeTranscodeTemplatesResponse) {
    response = &DescribeJustInTimeTranscodeTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeJustInTimeTranscodeTemplates
// This API is used to retrieve the transcoding template detail list based on the just in time transcoding template name, template type, or unique ID. The returned results include all eligible custom templates and preset templates.
//
// error code that may be returned:
//  FAILEDOPERATION_PRODUCTDISCONTINUED = "FailedOperation.ProductDiscontinued"
//  FAILEDOPERATION_PRODUCTNOTACTIVATED = "FailedOperation.ProductNotActivated"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ZONENOTFOUND = "InvalidParameterValue.ZoneNotFound"
//  INVALIDPARAMETERVALUE_ZONEPAUSED = "InvalidParameterValue.ZonePaused"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeJustInTimeTranscodeTemplates(request *DescribeJustInTimeTranscodeTemplatesRequest) (response *DescribeJustInTimeTranscodeTemplatesResponse, err error) {
    return c.DescribeJustInTimeTranscodeTemplatesWithContext(context.Background(), request)
}

// DescribeJustInTimeTranscodeTemplates
// This API is used to retrieve the transcoding template detail list based on the just in time transcoding template name, template type, or unique ID. The returned results include all eligible custom templates and preset templates.
//
// error code that may be returned:
//  FAILEDOPERATION_PRODUCTDISCONTINUED = "FailedOperation.ProductDiscontinued"
//  FAILEDOPERATION_PRODUCTNOTACTIVATED = "FailedOperation.ProductNotActivated"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ZONENOTFOUND = "InvalidParameterValue.ZoneNotFound"
//  INVALIDPARAMETERVALUE_ZONEPAUSED = "InvalidParameterValue.ZonePaused"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeJustInTimeTranscodeTemplatesWithContext(ctx context.Context, request *DescribeJustInTimeTranscodeTemplatesRequest) (response *DescribeJustInTimeTranscodeTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeJustInTimeTranscodeTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeJustInTimeTranscodeTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeJustInTimeTranscodeTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeJustInTimeTranscodeTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeL4ProxyRequest() (request *DescribeL4ProxyRequest) {
    request = &DescribeL4ProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeL4Proxy")
    
    
    return
}

func NewDescribeL4ProxyResponse() (response *DescribeL4ProxyResponse) {
    response = &DescribeL4ProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeL4Proxy
// This API is used to query the layer 4 proxy instance list.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeL4Proxy(request *DescribeL4ProxyRequest) (response *DescribeL4ProxyResponse, err error) {
    return c.DescribeL4ProxyWithContext(context.Background(), request)
}

// DescribeL4Proxy
// This API is used to query the layer 4 proxy instance list.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeL4ProxyWithContext(ctx context.Context, request *DescribeL4ProxyRequest) (response *DescribeL4ProxyResponse, err error) {
    if request == nil {
        request = NewDescribeL4ProxyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeL4Proxy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeL4Proxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeL4ProxyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeL4ProxyRulesRequest() (request *DescribeL4ProxyRulesRequest) {
    request = &DescribeL4ProxyRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeL4ProxyRules")
    
    
    return
}

func NewDescribeL4ProxyRulesResponse() (response *DescribeL4ProxyRulesResponse) {
    response = &DescribeL4ProxyRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeL4ProxyRules
// This API is used to query the forwarding rule list under a Layer 4 proxy instance.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeL4ProxyRules(request *DescribeL4ProxyRulesRequest) (response *DescribeL4ProxyRulesResponse, err error) {
    return c.DescribeL4ProxyRulesWithContext(context.Background(), request)
}

// DescribeL4ProxyRules
// This API is used to query the forwarding rule list under a Layer 4 proxy instance.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeL4ProxyRulesWithContext(ctx context.Context, request *DescribeL4ProxyRulesRequest) (response *DescribeL4ProxyRulesResponse, err error) {
    if request == nil {
        request = NewDescribeL4ProxyRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeL4ProxyRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeL4ProxyRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeL4ProxyRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeL7AccRulesRequest() (request *DescribeL7AccRulesRequest) {
    request = &DescribeL7AccRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeL7AccRules")
    
    
    return
}

func NewDescribeL7AccRulesResponse() (response *DescribeL7AccRulesResponse) {
    response = &DescribeL7AccRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeL7AccRules
// This API is used to query the rule list of the [rule engine](https://www.tencentcloud.com/document/product/1552/70901?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeL7AccRules(request *DescribeL7AccRulesRequest) (response *DescribeL7AccRulesResponse, err error) {
    return c.DescribeL7AccRulesWithContext(context.Background(), request)
}

// DescribeL7AccRules
// This API is used to query the rule list of the [rule engine](https://www.tencentcloud.com/document/product/1552/70901?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeL7AccRulesWithContext(ctx context.Context, request *DescribeL7AccRulesRequest) (response *DescribeL7AccRulesResponse, err error) {
    if request == nil {
        request = NewDescribeL7AccRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeL7AccRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeL7AccRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeL7AccRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeL7AccSettingRequest() (request *DescribeL7AccSettingRequest) {
    request = &DescribeL7AccSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeL7AccSetting")
    
    
    return
}

func NewDescribeL7AccSettingResponse() (response *DescribeL7AccSettingResponse) {
    response = &DescribeL7AccSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeL7AccSetting
// This API is used to query the global configuration of site acceleration.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_SETTINGINVALIDPARAM = "InvalidParameter.SettingInvalidParam"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeL7AccSetting(request *DescribeL7AccSettingRequest) (response *DescribeL7AccSettingResponse, err error) {
    return c.DescribeL7AccSettingWithContext(context.Background(), request)
}

// DescribeL7AccSetting
// This API is used to query the global configuration of site acceleration.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_SETTINGINVALIDPARAM = "InvalidParameter.SettingInvalidParam"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeL7AccSettingWithContext(ctx context.Context, request *DescribeL7AccSettingRequest) (response *DescribeL7AccSettingResponse, err error) {
    if request == nil {
        request = NewDescribeL7AccSettingRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeL7AccSetting")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeL7AccSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeL7AccSettingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLoadBalancerListRequest() (request *DescribeLoadBalancerListRequest) {
    request = &DescribeLoadBalancerListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeLoadBalancerList")
    
    
    return
}

func NewDescribeLoadBalancerListResponse() (response *DescribeLoadBalancerListResponse) {
    response = &DescribeLoadBalancerListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLoadBalancerList
// This API is used to query the load balancing instance list. The load balancer feature is in beta test. If you need to use it, [contact us](https://www.tencentcloud.com/contact-us).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeLoadBalancerList(request *DescribeLoadBalancerListRequest) (response *DescribeLoadBalancerListResponse, err error) {
    return c.DescribeLoadBalancerListWithContext(context.Background(), request)
}

// DescribeLoadBalancerList
// This API is used to query the load balancing instance list. The load balancer feature is in beta test. If you need to use it, [contact us](https://www.tencentcloud.com/contact-us).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeLoadBalancerListWithContext(ctx context.Context, request *DescribeLoadBalancerListRequest) (response *DescribeLoadBalancerListResponse, err error) {
    if request == nil {
        request = NewDescribeLoadBalancerListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeLoadBalancerList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLoadBalancerList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLoadBalancerListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogAnalysisDetailRequest() (request *DescribeLogAnalysisDetailRequest) {
    request = &DescribeLogAnalysisDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeLogAnalysisDetail")
    
    
    return
}

func NewDescribeLogAnalysisDetailResponse() (response *DescribeLogAnalysisDetailResponse) {
    response = &DescribeLogAnalysisDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLogAnalysisDetail
// This API is used to query log analysis details. The data comes from the log data pushed by real-time log push tasks under the site with the destination "log-analysis".
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeLogAnalysisDetail(request *DescribeLogAnalysisDetailRequest) (response *DescribeLogAnalysisDetailResponse, err error) {
    return c.DescribeLogAnalysisDetailWithContext(context.Background(), request)
}

// DescribeLogAnalysisDetail
// This API is used to query log analysis details. The data comes from the log data pushed by real-time log push tasks under the site with the destination "log-analysis".
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeLogAnalysisDetailWithContext(ctx context.Context, request *DescribeLogAnalysisDetailRequest) (response *DescribeLogAnalysisDetailResponse, err error) {
    if request == nil {
        request = NewDescribeLogAnalysisDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeLogAnalysisDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogAnalysisDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogAnalysisDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogAnalysisDownloadTasksRequest() (request *DescribeLogAnalysisDownloadTasksRequest) {
    request = &DescribeLogAnalysisDownloadTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeLogAnalysisDownloadTasks")
    
    
    return
}

func NewDescribeLogAnalysisDownloadTasksResponse() (response *DescribeLogAnalysisDownloadTasksResponse) {
    response = &DescribeLogAnalysisDownloadTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLogAnalysisDownloadTasks
// This API is used to query the log download task list for log analysis. Note: Only download task records from the last three days are retained.
//
// error code that may be returned:
//  FAILEDOPERATION_REALTIMELOGLOGANALYSISNOTSUPPORT = "FailedOperation.RealtimeLogLogAnalysisNotSupport"
//  INVALIDPARAMETER_INVALIDLOGANALYSISCONDITION = "InvalidParameter.InvalidLogAnalysisCondition"
func (c *Client) DescribeLogAnalysisDownloadTasks(request *DescribeLogAnalysisDownloadTasksRequest) (response *DescribeLogAnalysisDownloadTasksResponse, err error) {
    return c.DescribeLogAnalysisDownloadTasksWithContext(context.Background(), request)
}

// DescribeLogAnalysisDownloadTasks
// This API is used to query the log download task list for log analysis. Note: Only download task records from the last three days are retained.
//
// error code that may be returned:
//  FAILEDOPERATION_REALTIMELOGLOGANALYSISNOTSUPPORT = "FailedOperation.RealtimeLogLogAnalysisNotSupport"
//  INVALIDPARAMETER_INVALIDLOGANALYSISCONDITION = "InvalidParameter.InvalidLogAnalysisCondition"
func (c *Client) DescribeLogAnalysisDownloadTasksWithContext(ctx context.Context, request *DescribeLogAnalysisDownloadTasksRequest) (response *DescribeLogAnalysisDownloadTasksResponse, err error) {
    if request == nil {
        request = NewDescribeLogAnalysisDownloadTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeLogAnalysisDownloadTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogAnalysisDownloadTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogAnalysisDownloadTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogAnalysisStatisticsRequest() (request *DescribeLogAnalysisStatisticsRequest) {
    request = &DescribeLogAnalysisStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeLogAnalysisStatistics")
    
    
    return
}

func NewDescribeLogAnalysisStatisticsResponse() (response *DescribeLogAnalysisStatisticsResponse) {
    response = &DescribeLogAnalysisStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLogAnalysisStatistics
// This API is used to query log analysis statistics results. Users append an SQL aggregation statement after the query statement ([retrieval condition] | [SQL statement]). The backend executes the aggregation and returns columnar statistics results (column definition + data row).
//
// error code that may be returned:
//  FAILEDOPERATION_REALTIMELOGLOGANALYSISNOTSUPPORT = "FailedOperation.RealtimeLogLogAnalysisNotSupport"
//  INVALIDPARAMETER_INVALIDLOGANALYSISCONDITION = "InvalidParameter.InvalidLogAnalysisCondition"
func (c *Client) DescribeLogAnalysisStatistics(request *DescribeLogAnalysisStatisticsRequest) (response *DescribeLogAnalysisStatisticsResponse, err error) {
    return c.DescribeLogAnalysisStatisticsWithContext(context.Background(), request)
}

// DescribeLogAnalysisStatistics
// This API is used to query log analysis statistics results. Users append an SQL aggregation statement after the query statement ([retrieval condition] | [SQL statement]). The backend executes the aggregation and returns columnar statistics results (column definition + data row).
//
// error code that may be returned:
//  FAILEDOPERATION_REALTIMELOGLOGANALYSISNOTSUPPORT = "FailedOperation.RealtimeLogLogAnalysisNotSupport"
//  INVALIDPARAMETER_INVALIDLOGANALYSISCONDITION = "InvalidParameter.InvalidLogAnalysisCondition"
func (c *Client) DescribeLogAnalysisStatisticsWithContext(ctx context.Context, request *DescribeLogAnalysisStatisticsRequest) (response *DescribeLogAnalysisStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeLogAnalysisStatisticsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeLogAnalysisStatistics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogAnalysisStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogAnalysisStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMultiPathGatewayRequest() (request *DescribeMultiPathGatewayRequest) {
    request = &DescribeMultiPathGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeMultiPathGateway")
    
    
    return
}

func NewDescribeMultiPathGatewayResponse() (response *DescribeMultiPathGatewayResponse) {
    response = &DescribeMultiPathGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMultiPathGateway
// This API is used to query details of a multi-channel security acceleration gateway, such as name, Gateway ID, IP, port, and type.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeMultiPathGateway(request *DescribeMultiPathGatewayRequest) (response *DescribeMultiPathGatewayResponse, err error) {
    return c.DescribeMultiPathGatewayWithContext(context.Background(), request)
}

// DescribeMultiPathGateway
// This API is used to query details of a multi-channel security acceleration gateway, such as name, Gateway ID, IP, port, and type.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeMultiPathGatewayWithContext(ctx context.Context, request *DescribeMultiPathGatewayRequest) (response *DescribeMultiPathGatewayResponse, err error) {
    if request == nil {
        request = NewDescribeMultiPathGatewayRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeMultiPathGateway")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMultiPathGateway require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMultiPathGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMultiPathGatewayLineRequest() (request *DescribeMultiPathGatewayLineRequest) {
    request = &DescribeMultiPathGatewayLineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeMultiPathGatewayLine")
    
    
    return
}

func NewDescribeMultiPathGatewayLineResponse() (response *DescribeMultiPathGatewayLineResponse) {
    response = &DescribeMultiPathGatewayLineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMultiPathGatewayLine
// This API is used to query lines connected to the multi-channel security acceleration gateway, including direct connections, EdgeOne Layer 4 proxy lines, and custom lines.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeMultiPathGatewayLine(request *DescribeMultiPathGatewayLineRequest) (response *DescribeMultiPathGatewayLineResponse, err error) {
    return c.DescribeMultiPathGatewayLineWithContext(context.Background(), request)
}

// DescribeMultiPathGatewayLine
// This API is used to query lines connected to the multi-channel security acceleration gateway, including direct connections, EdgeOne Layer 4 proxy lines, and custom lines.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeMultiPathGatewayLineWithContext(ctx context.Context, request *DescribeMultiPathGatewayLineRequest) (response *DescribeMultiPathGatewayLineResponse, err error) {
    if request == nil {
        request = NewDescribeMultiPathGatewayLineRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeMultiPathGatewayLine")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMultiPathGatewayLine require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMultiPathGatewayLineResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMultiPathGatewayOriginACLRequest() (request *DescribeMultiPathGatewayOriginACLRequest) {
    request = &DescribeMultiPathGatewayOriginACLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeMultiPathGatewayOriginACL")
    
    
    return
}

func NewDescribeMultiPathGatewayOriginACLResponse() (response *DescribeMultiPathGatewayOriginACLResponse) {
    response = &DescribeMultiPathGatewayOriginACLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMultiPathGatewayOriginACL
// This API is used to query the binding relationship between a multi-channel security acceleration gateway instance and origin IP ranges, as well as origin IP range details. If the MultiPathGatewayNextOriginACL field has a return value, synchronize the latest origin IP ranges to the origin server firewall settings.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeMultiPathGatewayOriginACL(request *DescribeMultiPathGatewayOriginACLRequest) (response *DescribeMultiPathGatewayOriginACLResponse, err error) {
    return c.DescribeMultiPathGatewayOriginACLWithContext(context.Background(), request)
}

// DescribeMultiPathGatewayOriginACL
// This API is used to query the binding relationship between a multi-channel security acceleration gateway instance and origin IP ranges, as well as origin IP range details. If the MultiPathGatewayNextOriginACL field has a return value, synchronize the latest origin IP ranges to the origin server firewall settings.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeMultiPathGatewayOriginACLWithContext(ctx context.Context, request *DescribeMultiPathGatewayOriginACLRequest) (response *DescribeMultiPathGatewayOriginACLResponse, err error) {
    if request == nil {
        request = NewDescribeMultiPathGatewayOriginACLRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeMultiPathGatewayOriginACL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMultiPathGatewayOriginACL require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMultiPathGatewayOriginACLResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMultiPathGatewayRegionsRequest() (request *DescribeMultiPathGatewayRegionsRequest) {
    request = &DescribeMultiPathGatewayRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeMultiPathGatewayRegions")
    
    
    return
}

func NewDescribeMultiPathGatewayRegionsResponse() (response *DescribeMultiPathGatewayRegionsResponse) {
    response = &DescribeMultiPathGatewayRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMultiPathGatewayRegions
// This API is used to query the available region list of user-created multi-channel security acceleration gateways.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeMultiPathGatewayRegions(request *DescribeMultiPathGatewayRegionsRequest) (response *DescribeMultiPathGatewayRegionsResponse, err error) {
    return c.DescribeMultiPathGatewayRegionsWithContext(context.Background(), request)
}

// DescribeMultiPathGatewayRegions
// This API is used to query the available region list of user-created multi-channel security acceleration gateways.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeMultiPathGatewayRegionsWithContext(ctx context.Context, request *DescribeMultiPathGatewayRegionsRequest) (response *DescribeMultiPathGatewayRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeMultiPathGatewayRegionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeMultiPathGatewayRegions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMultiPathGatewayRegions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMultiPathGatewayRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMultiPathGatewaySecretKeyRequest() (request *DescribeMultiPathGatewaySecretKeyRequest) {
    request = &DescribeMultiPathGatewaySecretKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeMultiPathGatewaySecretKey")
    
    
    return
}

func NewDescribeMultiPathGatewaySecretKeyResponse() (response *DescribeMultiPathGatewaySecretKeyResponse) {
    response = &DescribeMultiPathGatewaySecretKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMultiPathGatewaySecretKey
// This API is used to query the key for integrating a multi-channel security acceleration gateway. Customers sign with the integration key to access the multi-channel security acceleration gateway.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeMultiPathGatewaySecretKey(request *DescribeMultiPathGatewaySecretKeyRequest) (response *DescribeMultiPathGatewaySecretKeyResponse, err error) {
    return c.DescribeMultiPathGatewaySecretKeyWithContext(context.Background(), request)
}

// DescribeMultiPathGatewaySecretKey
// This API is used to query the key for integrating a multi-channel security acceleration gateway. Customers sign with the integration key to access the multi-channel security acceleration gateway.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeMultiPathGatewaySecretKeyWithContext(ctx context.Context, request *DescribeMultiPathGatewaySecretKeyRequest) (response *DescribeMultiPathGatewaySecretKeyResponse, err error) {
    if request == nil {
        request = NewDescribeMultiPathGatewaySecretKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeMultiPathGatewaySecretKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMultiPathGatewaySecretKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMultiPathGatewaySecretKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMultiPathGatewaysRequest() (request *DescribeMultiPathGatewaysRequest) {
    request = &DescribeMultiPathGatewaysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeMultiPathGateways")
    
    
    return
}

func NewDescribeMultiPathGatewaysResponse() (response *DescribeMultiPathGatewaysResponse) {
    response = &DescribeMultiPathGatewaysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMultiPathGateways
// This API is used to query the list of user-created multi-channel security acceleration gateways. Pagination is supported.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeMultiPathGateways(request *DescribeMultiPathGatewaysRequest) (response *DescribeMultiPathGatewaysResponse, err error) {
    return c.DescribeMultiPathGatewaysWithContext(context.Background(), request)
}

// DescribeMultiPathGateways
// This API is used to query the list of user-created multi-channel security acceleration gateways. Pagination is supported.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeMultiPathGatewaysWithContext(ctx context.Context, request *DescribeMultiPathGatewaysRequest) (response *DescribeMultiPathGatewaysResponse, err error) {
    if request == nil {
        request = NewDescribeMultiPathGatewaysRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeMultiPathGateways")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMultiPathGateways require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMultiPathGatewaysResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOriginACLRequest() (request *DescribeOriginACLRequest) {
    request = &DescribeOriginACLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeOriginACL")
    
    
    return
}

func NewDescribeOriginACLResponse() (response *DescribeOriginACLResponse) {
    response = &DescribeOriginACLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOriginACL
// This API is used to query the binding relationship between L7 acceleration domains/L4 proxy instances and origin IP ranges under a site, as well as origin IP range details. If you want to periodically get the latest version of origin IP ranges through an automation script, you can poll this API at a low frequency (once every three days is recommended). If the NextOriginACL field has a return value, synchronize the latest origin IP ranges to the origin server firewall configuration.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeOriginACL(request *DescribeOriginACLRequest) (response *DescribeOriginACLResponse, err error) {
    return c.DescribeOriginACLWithContext(context.Background(), request)
}

// DescribeOriginACL
// This API is used to query the binding relationship between L7 acceleration domains/L4 proxy instances and origin IP ranges under a site, as well as origin IP range details. If you want to periodically get the latest version of origin IP ranges through an automation script, you can poll this API at a low frequency (once every three days is recommended). If the NextOriginACL field has a return value, synchronize the latest origin IP ranges to the origin server firewall configuration.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeOriginACLWithContext(ctx context.Context, request *DescribeOriginACLRequest) (response *DescribeOriginACLResponse, err error) {
    if request == nil {
        request = NewDescribeOriginACLRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeOriginACL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOriginACL require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOriginACLResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOriginGroupRequest() (request *DescribeOriginGroupRequest) {
    request = &DescribeOriginGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeOriginGroup")
    
    
    return
}

func NewDescribeOriginGroupResponse() (response *DescribeOriginGroupResponse) {
    response = &DescribeOriginGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOriginGroup
// This API is used to query the origin server group list.
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeOriginGroup(request *DescribeOriginGroupRequest) (response *DescribeOriginGroupResponse, err error) {
    return c.DescribeOriginGroupWithContext(context.Background(), request)
}

// DescribeOriginGroup
// This API is used to query the origin server group list.
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeOriginGroupWithContext(ctx context.Context, request *DescribeOriginGroupRequest) (response *DescribeOriginGroupResponse, err error) {
    if request == nil {
        request = NewDescribeOriginGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeOriginGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOriginGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOriginGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOriginGroupHealthStatusRequest() (request *DescribeOriginGroupHealthStatusRequest) {
    request = &DescribeOriginGroupHealthStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeOriginGroupHealthStatus")
    
    
    return
}

func NewDescribeOriginGroupHealthStatusResponse() (response *DescribeOriginGroupHealthStatusResponse) {
    response = &DescribeOriginGroupHealthStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOriginGroupHealthStatus
// This API is used to query the health status of origin server groups under a load balancing instance. The load balancer feature is in beta test. If you need to use it, [contact us](https://www.tencentcloud.com/contact-us).
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeOriginGroupHealthStatus(request *DescribeOriginGroupHealthStatusRequest) (response *DescribeOriginGroupHealthStatusResponse, err error) {
    return c.DescribeOriginGroupHealthStatusWithContext(context.Background(), request)
}

// DescribeOriginGroupHealthStatus
// This API is used to query the health status of origin server groups under a load balancing instance. The load balancer feature is in beta test. If you need to use it, [contact us](https://www.tencentcloud.com/contact-us).
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeOriginGroupHealthStatusWithContext(ctx context.Context, request *DescribeOriginGroupHealthStatusRequest) (response *DescribeOriginGroupHealthStatusResponse, err error) {
    if request == nil {
        request = NewDescribeOriginGroupHealthStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeOriginGroupHealthStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOriginGroupHealthStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOriginGroupHealthStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOriginProtectionRequest() (request *DescribeOriginProtectionRequest) {
    request = &DescribeOriginProtectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeOriginProtection")
    
    
    return
}

func NewDescribeOriginProtectionResponse() (response *DescribeOriginProtectionResponse) {
    response = &DescribeOriginProtectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOriginProtection
// This API is on an earlier version for origin protection queries. EdgeOne comprehensively upgraded the origin protection APIs on June 27, 2025. For details on the new version of the origin protection query API, refer to [DescribeOriginACL](https://www.tencentcloud.com/document/product/1552/120408?from_cn_redirect=1).
//
// 
//
// <p style="color: red;">Note: Starting from June 27, 2025, earlier version APIs will stop update and iteration, and new features will only be provided in the new version of APIs. To avoid data field conflicts when using earlier version APIs, migrate to the new version of origin protection APIs as soon as possible.</p>
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeOriginProtection(request *DescribeOriginProtectionRequest) (response *DescribeOriginProtectionResponse, err error) {
    return c.DescribeOriginProtectionWithContext(context.Background(), request)
}

// DescribeOriginProtection
// This API is on an earlier version for origin protection queries. EdgeOne comprehensively upgraded the origin protection APIs on June 27, 2025. For details on the new version of the origin protection query API, refer to [DescribeOriginACL](https://www.tencentcloud.com/document/product/1552/120408?from_cn_redirect=1).
//
// 
//
// <p style="color: red;">Note: Starting from June 27, 2025, earlier version APIs will stop update and iteration, and new features will only be provided in the new version of APIs. To avoid data field conflicts when using earlier version APIs, migrate to the new version of origin protection APIs as soon as possible.</p>
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeOriginProtectionWithContext(ctx context.Context, request *DescribeOriginProtectionRequest) (response *DescribeOriginProtectionResponse, err error) {
    if request == nil {
        request = NewDescribeOriginProtectionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeOriginProtection")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOriginProtection require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOriginProtectionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOverviewL7DataRequest() (request *DescribeOverviewL7DataRequest) {
    request = &DescribeOverviewL7DataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeOverviewL7Data")
    
    
    return
}

func NewDescribeOverviewL7DataResponse() (response *DescribeOverviewL7DataResponse) {
    response = &DescribeOverviewL7DataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOverviewL7Data
// This API is used to query layer-7 monitoring time series traffic Data. This API is to be deprecated. Please use the <a href="https://www.tencentcloud.com/document/product/1552/80648?from_cn_redirect=1">DescribeTimingL7AnalysisData</a> API.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeOverviewL7Data(request *DescribeOverviewL7DataRequest) (response *DescribeOverviewL7DataResponse, err error) {
    return c.DescribeOverviewL7DataWithContext(context.Background(), request)
}

// DescribeOverviewL7Data
// This API is used to query layer-7 monitoring time series traffic Data. This API is to be deprecated. Please use the <a href="https://www.tencentcloud.com/document/product/1552/80648?from_cn_redirect=1">DescribeTimingL7AnalysisData</a> API.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeOverviewL7DataWithContext(ctx context.Context, request *DescribeOverviewL7DataRequest) (response *DescribeOverviewL7DataResponse, err error) {
    if request == nil {
        request = NewDescribeOverviewL7DataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeOverviewL7Data")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOverviewL7Data require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOverviewL7DataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePlansRequest() (request *DescribePlansRequest) {
    request = &DescribePlansRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribePlans")
    
    
    return
}

func NewDescribePlansResponse() (response *DescribePlansResponse) {
    response = &DescribePlansResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePlans
// Query the package information list with pagination.
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribePlans(request *DescribePlansRequest) (response *DescribePlansResponse, err error) {
    return c.DescribePlansWithContext(context.Background(), request)
}

// DescribePlans
// Query the package information list with pagination.
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribePlansWithContext(ctx context.Context, request *DescribePlansRequest) (response *DescribePlansResponse, err error) {
    if request == nil {
        request = NewDescribePlansRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribePlans")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePlans require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePlansResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrefetchOriginLimitRequest() (request *DescribePrefetchOriginLimitRequest) {
    request = &DescribePrefetchOriginLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribePrefetchOriginLimit")
    
    
    return
}

func NewDescribePrefetchOriginLimitResponse() (response *DescribePrefetchOriginLimitResponse) {
    response = &DescribePrefetchOriginLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrefetchOriginLimit
// This API is used to query origin-pull rate limits. This feature is in beta test for allowlisted users.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINNOTFOUND = "InvalidParameter.DomainNotFound"
//  INVALIDPARAMETER_PREFETCHORIGINLIMITAREAINVALID = "InvalidParameter.PrefetchOriginLimitAreaInvalid"
//  OPERATIONDENIED_NOTINPREFETCHORIGINLIMITWHITELIST = "OperationDenied.NotInPrefetchOriginLimitWhiteList"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribePrefetchOriginLimit(request *DescribePrefetchOriginLimitRequest) (response *DescribePrefetchOriginLimitResponse, err error) {
    return c.DescribePrefetchOriginLimitWithContext(context.Background(), request)
}

// DescribePrefetchOriginLimit
// This API is used to query origin-pull rate limits. This feature is in beta test for allowlisted users.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINNOTFOUND = "InvalidParameter.DomainNotFound"
//  INVALIDPARAMETER_PREFETCHORIGINLIMITAREAINVALID = "InvalidParameter.PrefetchOriginLimitAreaInvalid"
//  OPERATIONDENIED_NOTINPREFETCHORIGINLIMITWHITELIST = "OperationDenied.NotInPrefetchOriginLimitWhiteList"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribePrefetchOriginLimitWithContext(ctx context.Context, request *DescribePrefetchOriginLimitRequest) (response *DescribePrefetchOriginLimitResponse, err error) {
    if request == nil {
        request = NewDescribePrefetchOriginLimitRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribePrefetchOriginLimit")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrefetchOriginLimit require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrefetchOriginLimitResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrefetchTasksRequest() (request *DescribePrefetchTasksRequest) {
    request = &DescribePrefetchTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribePrefetchTasks")
    
    
    return
}

func NewDescribePrefetchTasksResponse() (response *DescribePrefetchTasksResponse) {
    response = &DescribePrefetchTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrefetchTasks
// This API is used to query the commit history and execution progress of preheating tasks. Tasks submitted via the CreatePrefetchTasks API can be queried through this API.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_DOMAINEMPTY = "UnauthorizedOperation.DomainEmpty"
func (c *Client) DescribePrefetchTasks(request *DescribePrefetchTasksRequest) (response *DescribePrefetchTasksResponse, err error) {
    return c.DescribePrefetchTasksWithContext(context.Background(), request)
}

// DescribePrefetchTasks
// This API is used to query the commit history and execution progress of preheating tasks. Tasks submitted via the CreatePrefetchTasks API can be queried through this API.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_DOMAINEMPTY = "UnauthorizedOperation.DomainEmpty"
func (c *Client) DescribePrefetchTasksWithContext(ctx context.Context, request *DescribePrefetchTasksRequest) (response *DescribePrefetchTasksResponse, err error) {
    if request == nil {
        request = NewDescribePrefetchTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribePrefetchTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrefetchTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrefetchTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePurgeTasksRequest() (request *DescribePurgeTasksRequest) {
    request = &DescribePurgeTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribePurgeTasks")
    
    
    return
}

func NewDescribePurgeTasksResponse() (response *DescribePurgeTasksResponse) {
    response = &DescribePurgeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePurgeTasks
// This API is used to query submitted URL refresh and directory refresh records and task progress. All tasks submitted via the CreatePurgeTasks API can be queried through this API.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribePurgeTasks(request *DescribePurgeTasksRequest) (response *DescribePurgeTasksResponse, err error) {
    return c.DescribePurgeTasksWithContext(context.Background(), request)
}

// DescribePurgeTasks
// This API is used to query submitted URL refresh and directory refresh records and task progress. All tasks submitted via the CreatePurgeTasks API can be queried through this API.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribePurgeTasksWithContext(ctx context.Context, request *DescribePurgeTasksRequest) (response *DescribePurgeTasksResponse, err error) {
    if request == nil {
        request = NewDescribePurgeTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribePurgeTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePurgeTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePurgeTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRealtimeLogDeliveryTasksRequest() (request *DescribeRealtimeLogDeliveryTasksRequest) {
    request = &DescribeRealtimeLogDeliveryTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeRealtimeLogDeliveryTasks")
    
    
    return
}

func NewDescribeRealtimeLogDeliveryTasksResponse() (response *DescribeRealtimeLogDeliveryTasksResponse) {
    response = &DescribeRealtimeLogDeliveryTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRealtimeLogDeliveryTasks
// This API is used to query real-time log delivery task list.
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeRealtimeLogDeliveryTasks(request *DescribeRealtimeLogDeliveryTasksRequest) (response *DescribeRealtimeLogDeliveryTasksResponse, err error) {
    return c.DescribeRealtimeLogDeliveryTasksWithContext(context.Background(), request)
}

// DescribeRealtimeLogDeliveryTasks
// This API is used to query real-time log delivery task list.
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeRealtimeLogDeliveryTasksWithContext(ctx context.Context, request *DescribeRealtimeLogDeliveryTasksRequest) (response *DescribeRealtimeLogDeliveryTasksResponse, err error) {
    if request == nil {
        request = NewDescribeRealtimeLogDeliveryTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeRealtimeLogDeliveryTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRealtimeLogDeliveryTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRealtimeLogDeliveryTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRulesRequest() (request *DescribeRulesRequest) {
    request = &DescribeRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeRules")
    
    
    return
}

func NewDescribeRulesResponse() (response *DescribeRulesResponse) {
    response = &DescribeRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRules
// This API is on an earlier version for querying rule engine rules. EdgeOne comprehensively upgraded the rule engine APIs on January 21, 2025. For details on the new version API for querying L7 acceleration rules, see [DescribeL7AccRules](https://www.tencentcloud.com/document/product/1552/115820?from_cn_redirect=1).
//
// <p style="color: red;">Note: Starting from January 21, 2025, the legacy version API will stop update and iteration. New features will only be provided in the new version API. The existing capacity supported by the legacy version API will not be affected. To avoid data field conflicts when using the legacy version API, migrate to the new version rule engine API as soon as possible.</p>
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeRules(request *DescribeRulesRequest) (response *DescribeRulesResponse, err error) {
    return c.DescribeRulesWithContext(context.Background(), request)
}

// DescribeRules
// This API is on an earlier version for querying rule engine rules. EdgeOne comprehensively upgraded the rule engine APIs on January 21, 2025. For details on the new version API for querying L7 acceleration rules, see [DescribeL7AccRules](https://www.tencentcloud.com/document/product/1552/115820?from_cn_redirect=1).
//
// <p style="color: red;">Note: Starting from January 21, 2025, the legacy version API will stop update and iteration. New features will only be provided in the new version API. The existing capacity supported by the legacy version API will not be affected. To avoid data field conflicts when using the legacy version API, migrate to the new version rule engine API as soon as possible.</p>
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeRulesWithContext(ctx context.Context, request *DescribeRulesRequest) (response *DescribeRulesResponse, err error) {
    if request == nil {
        request = NewDescribeRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRulesSettingRequest() (request *DescribeRulesSettingRequest) {
    request = &DescribeRulesSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeRulesSetting")
    
    
    return
}

func NewDescribeRulesSettingResponse() (response *DescribeRulesSettingResponse) {
    response = &DescribeRulesSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRulesSetting
// This API is on an earlier version. EdgeOne has completed a comprehensive upgrade of rule engine APIs. For details, see [RuleEngineAction](https://www.tencentcloud.com/document/product/1552/80721?from_cn_redirect=1#RuleEngineAction).
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeRulesSetting(request *DescribeRulesSettingRequest) (response *DescribeRulesSettingResponse, err error) {
    return c.DescribeRulesSettingWithContext(context.Background(), request)
}

// DescribeRulesSetting
// This API is on an earlier version. EdgeOne has completed a comprehensive upgrade of rule engine APIs. For details, see [RuleEngineAction](https://www.tencentcloud.com/document/product/1552/80721?from_cn_redirect=1#RuleEngineAction).
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeRulesSettingWithContext(ctx context.Context, request *DescribeRulesSettingRequest) (response *DescribeRulesSettingResponse, err error) {
    if request == nil {
        request = NewDescribeRulesSettingRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeRulesSetting")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRulesSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRulesSettingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityAPIResourceRequest() (request *DescribeSecurityAPIResourceRequest) {
    request = &DescribeSecurityAPIResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeSecurityAPIResource")
    
    
    return
}

func NewDescribeSecurityAPIResourceResponse() (response *DescribeSecurityAPIResourceResponse) {
    response = &DescribeSecurityAPIResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecurityAPIResource
// Query API resources under a site.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeSecurityAPIResource(request *DescribeSecurityAPIResourceRequest) (response *DescribeSecurityAPIResourceResponse, err error) {
    return c.DescribeSecurityAPIResourceWithContext(context.Background(), request)
}

// DescribeSecurityAPIResource
// Query API resources under a site.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeSecurityAPIResourceWithContext(ctx context.Context, request *DescribeSecurityAPIResourceRequest) (response *DescribeSecurityAPIResourceResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityAPIResourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeSecurityAPIResource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityAPIResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityAPIResourceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityAPIServiceRequest() (request *DescribeSecurityAPIServiceRequest) {
    request = &DescribeSecurityAPIServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeSecurityAPIService")
    
    
    return
}

func NewDescribeSecurityAPIServiceResponse() (response *DescribeSecurityAPIServiceResponse) {
    response = &DescribeSecurityAPIServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecurityAPIService
// Query API services under a site.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeSecurityAPIService(request *DescribeSecurityAPIServiceRequest) (response *DescribeSecurityAPIServiceResponse, err error) {
    return c.DescribeSecurityAPIServiceWithContext(context.Background(), request)
}

// DescribeSecurityAPIService
// Query API services under a site.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeSecurityAPIServiceWithContext(ctx context.Context, request *DescribeSecurityAPIServiceRequest) (response *DescribeSecurityAPIServiceResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityAPIServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeSecurityAPIService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityAPIService require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityAPIServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityClientAttesterRequest() (request *DescribeSecurityClientAttesterRequest) {
    request = &DescribeSecurityClientAttesterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeSecurityClientAttester")
    
    
    return
}

func NewDescribeSecurityClientAttesterResponse() (response *DescribeSecurityClientAttesterResponse) {
    response = &DescribeSecurityClientAttesterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecurityClientAttester
// Query client authentication option configurations.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeSecurityClientAttester(request *DescribeSecurityClientAttesterRequest) (response *DescribeSecurityClientAttesterResponse, err error) {
    return c.DescribeSecurityClientAttesterWithContext(context.Background(), request)
}

// DescribeSecurityClientAttester
// Query client authentication option configurations.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeSecurityClientAttesterWithContext(ctx context.Context, request *DescribeSecurityClientAttesterRequest) (response *DescribeSecurityClientAttesterResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityClientAttesterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeSecurityClientAttester")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityClientAttester require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityClientAttesterResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityIPGroupRequest() (request *DescribeSecurityIPGroupRequest) {
    request = &DescribeSecurityIPGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeSecurityIPGroup")
    
    
    return
}

func NewDescribeSecurityIPGroupResponse() (response *DescribeSecurityIPGroupResponse) {
    response = &DescribeSecurityIPGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecurityIPGroup
// This API is used to query the configuration information of a security IP group, including its ID, name, and content. In the query result of this API, each IP group returns up to 2,000 IPs or CIDR blocks. If an ultra-large IP group exceeds 2,000 IPs or CIDR blocks, call DescribeSecurityIPGroupContent for paging query.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeSecurityIPGroup(request *DescribeSecurityIPGroupRequest) (response *DescribeSecurityIPGroupResponse, err error) {
    return c.DescribeSecurityIPGroupWithContext(context.Background(), request)
}

// DescribeSecurityIPGroup
// This API is used to query the configuration information of a security IP group, including its ID, name, and content. In the query result of this API, each IP group returns up to 2,000 IPs or CIDR blocks. If an ultra-large IP group exceeds 2,000 IPs or CIDR blocks, call DescribeSecurityIPGroupContent for paging query.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeSecurityIPGroupWithContext(ctx context.Context, request *DescribeSecurityIPGroupRequest) (response *DescribeSecurityIPGroupResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityIPGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeSecurityIPGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityIPGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityIPGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityIPGroupContentRequest() (request *DescribeSecurityIPGroupContentRequest) {
    request = &DescribeSecurityIPGroupContentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeSecurityIPGroupContent")
    
    
    return
}

func NewDescribeSecurityIPGroupContentResponse() (response *DescribeSecurityIPGroupContentResponse) {
    response = &DescribeSecurityIPGroupContentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecurityIPGroupContent
// This API is used to query the IP address list in a specified IP group by page. When the number of IP addresses in an IP group exceeds 2000, you can use this API to perform a paging query to obtain the complete IP address list.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeSecurityIPGroupContent(request *DescribeSecurityIPGroupContentRequest) (response *DescribeSecurityIPGroupContentResponse, err error) {
    return c.DescribeSecurityIPGroupContentWithContext(context.Background(), request)
}

// DescribeSecurityIPGroupContent
// This API is used to query the IP address list in a specified IP group by page. When the number of IP addresses in an IP group exceeds 2000, you can use this API to perform a paging query to obtain the complete IP address list.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeSecurityIPGroupContentWithContext(ctx context.Context, request *DescribeSecurityIPGroupContentRequest) (response *DescribeSecurityIPGroupContentResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityIPGroupContentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeSecurityIPGroupContent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityIPGroupContent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityIPGroupContentResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityIPGroupInfoRequest() (request *DescribeSecurityIPGroupInfoRequest) {
    request = &DescribeSecurityIPGroupInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeSecurityIPGroupInfo")
    
    
    return
}

func NewDescribeSecurityIPGroupInfoResponse() (response *DescribeSecurityIPGroupInfoResponse) {
    response = &DescribeSecurityIPGroupInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecurityIPGroupInfo
// This API is deprecated and will be discontinued on June 30, 2024. Please use the API for querying security IP groups.
//
// ](https://www.tencentcloud.com/document/product/1552/105866?from_cn_redirect=1).
//
// 
//
// This API is used to query the configuration information of an IP group, including the IP group name, IP group content, and the site the IP group belongs to.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeSecurityIPGroupInfo(request *DescribeSecurityIPGroupInfoRequest) (response *DescribeSecurityIPGroupInfoResponse, err error) {
    return c.DescribeSecurityIPGroupInfoWithContext(context.Background(), request)
}

// DescribeSecurityIPGroupInfo
// This API is deprecated and will be discontinued on June 30, 2024. Please use the API for querying security IP groups.
//
// ](https://www.tencentcloud.com/document/product/1552/105866?from_cn_redirect=1).
//
// 
//
// This API is used to query the configuration information of an IP group, including the IP group name, IP group content, and the site the IP group belongs to.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeSecurityIPGroupInfoWithContext(ctx context.Context, request *DescribeSecurityIPGroupInfoRequest) (response *DescribeSecurityIPGroupInfoResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityIPGroupInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeSecurityIPGroupInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityIPGroupInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityIPGroupInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityJSInjectionRuleRequest() (request *DescribeSecurityJSInjectionRuleRequest) {
    request = &DescribeSecurityJSInjectionRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeSecurityJSInjectionRule")
    
    
    return
}

func NewDescribeSecurityJSInjectionRuleResponse() (response *DescribeSecurityJSInjectionRuleResponse) {
    response = &DescribeSecurityJSInjectionRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecurityJSInjectionRule
// This API is used to query JavaScript injection rules.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeSecurityJSInjectionRule(request *DescribeSecurityJSInjectionRuleRequest) (response *DescribeSecurityJSInjectionRuleResponse, err error) {
    return c.DescribeSecurityJSInjectionRuleWithContext(context.Background(), request)
}

// DescribeSecurityJSInjectionRule
// This API is used to query JavaScript injection rules.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeSecurityJSInjectionRuleWithContext(ctx context.Context, request *DescribeSecurityJSInjectionRuleRequest) (response *DescribeSecurityJSInjectionRuleResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityJSInjectionRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeSecurityJSInjectionRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityJSInjectionRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityJSInjectionRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityPolicyRequest() (request *DescribeSecurityPolicyRequest) {
    request = &DescribeSecurityPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeSecurityPolicy")
    
    
    return
}

func NewDescribeSecurityPolicyResponse() (response *DescribeSecurityPolicyResponse) {
    response = &DescribeSecurityPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecurityPolicy
// Query security protection configuration details.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeSecurityPolicy(request *DescribeSecurityPolicyRequest) (response *DescribeSecurityPolicyResponse, err error) {
    return c.DescribeSecurityPolicyWithContext(context.Background(), request)
}

// DescribeSecurityPolicy
// Query security protection configuration details.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeSecurityPolicyWithContext(ctx context.Context, request *DescribeSecurityPolicyRequest) (response *DescribeSecurityPolicyResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeSecurityPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityTemplateBindingsRequest() (request *DescribeSecurityTemplateBindingsRequest) {
    request = &DescribeSecurityTemplateBindingsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeSecurityTemplateBindings")
    
    
    return
}

func NewDescribeSecurityTemplateBindingsResponse() (response *DescribeSecurityTemplateBindingsResponse) {
    response = &DescribeSecurityTemplateBindingsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecurityTemplateBindings
// This API is used to query the binding relationship list of a specified policy template.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeSecurityTemplateBindings(request *DescribeSecurityTemplateBindingsRequest) (response *DescribeSecurityTemplateBindingsResponse, err error) {
    return c.DescribeSecurityTemplateBindingsWithContext(context.Background(), request)
}

// DescribeSecurityTemplateBindings
// This API is used to query the binding relationship list of a specified policy template.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeSecurityTemplateBindingsWithContext(ctx context.Context, request *DescribeSecurityTemplateBindingsRequest) (response *DescribeSecurityTemplateBindingsResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityTemplateBindingsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeSecurityTemplateBindings")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityTemplateBindings require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityTemplateBindingsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSharedCNAMERequest() (request *DescribeSharedCNAMERequest) {
    request = &DescribeSharedCNAMERequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeSharedCNAME")
    
    
    return
}

func NewDescribeSharedCNAMEResponse() (response *DescribeSharedCNAMEResponse) {
    response = &DescribeSharedCNAMEResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSharedCNAME
// This API is used to query the shared CNAME list, supporting fuzzy search, pagination, and sorting.
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeSharedCNAME(request *DescribeSharedCNAMERequest) (response *DescribeSharedCNAMEResponse, err error) {
    return c.DescribeSharedCNAMEWithContext(context.Background(), request)
}

// DescribeSharedCNAME
// This API is used to query the shared CNAME list, supporting fuzzy search, pagination, and sorting.
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeSharedCNAMEWithContext(ctx context.Context, request *DescribeSharedCNAMERequest) (response *DescribeSharedCNAMEResponse, err error) {
    if request == nil {
        request = NewDescribeSharedCNAMERequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeSharedCNAME")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSharedCNAME require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSharedCNAMEResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTimingFunctionAnalysisDataRequest() (request *DescribeTimingFunctionAnalysisDataRequest) {
    request = &DescribeTimingFunctionAnalysisDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeTimingFunctionAnalysisData")
    
    
    return
}

func NewDescribeTimingFunctionAnalysisDataResponse() (response *DescribeTimingFunctionAnalysisDataResponse) {
    response = &DescribeTimingFunctionAnalysisDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTimingFunctionAnalysisData
// This API is used to query the time series data of edge function analysis. Note that the data queried by this API has a delay of about 10 minutes. It is recommended to pull data from at least 10 minutes before the current time.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeTimingFunctionAnalysisData(request *DescribeTimingFunctionAnalysisDataRequest) (response *DescribeTimingFunctionAnalysisDataResponse, err error) {
    return c.DescribeTimingFunctionAnalysisDataWithContext(context.Background(), request)
}

// DescribeTimingFunctionAnalysisData
// This API is used to query the time series data of edge function analysis. Note that the data queried by this API has a delay of about 10 minutes. It is recommended to pull data from at least 10 minutes before the current time.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeTimingFunctionAnalysisDataWithContext(ctx context.Context, request *DescribeTimingFunctionAnalysisDataRequest) (response *DescribeTimingFunctionAnalysisDataResponse, err error) {
    if request == nil {
        request = NewDescribeTimingFunctionAnalysisDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeTimingFunctionAnalysisData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTimingFunctionAnalysisData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTimingFunctionAnalysisDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTimingL4DataRequest() (request *DescribeTimingL4DataRequest) {
    request = &DescribeTimingL4DataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeTimingL4Data")
    
    
    return
}

func NewDescribeTimingL4DataResponse() (response *DescribeTimingL4DataResponse) {
    response = &DescribeTimingL4DataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTimingL4Data
// <p>This API is used to query the L4 time series Data list.</p>
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeTimingL4Data(request *DescribeTimingL4DataRequest) (response *DescribeTimingL4DataResponse, err error) {
    return c.DescribeTimingL4DataWithContext(context.Background(), request)
}

// DescribeTimingL4Data
// <p>This API is used to query the L4 time series Data list.</p>
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeTimingL4DataWithContext(ctx context.Context, request *DescribeTimingL4DataRequest) (response *DescribeTimingL4DataResponse, err error) {
    if request == nil {
        request = NewDescribeTimingL4DataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeTimingL4Data")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTimingL4Data require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTimingL4DataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTimingL7AnalysisDataRequest() (request *DescribeTimingL7AnalysisDataRequest) {
    request = &DescribeTimingL7AnalysisDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeTimingL7AnalysisData")
    
    
    return
}

func NewDescribeTimingL7AnalysisDataResponse() (response *DescribeTimingL7AnalysisDataResponse) {
    response = &DescribeTimingL7AnalysisDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTimingL7AnalysisData
// This API is used to query time series data of the L7 domain name business.
//
// Note:
//
// 1. This API queries data with a delay of about 10 minutes. It is advisable to pull data from 10 minutes before the current time.
//
// 2. This interface returns post-protection traffic request data by default. Users can use `Filters.mitigatedByWebSecurity` to custom query defended data.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeTimingL7AnalysisData(request *DescribeTimingL7AnalysisDataRequest) (response *DescribeTimingL7AnalysisDataResponse, err error) {
    return c.DescribeTimingL7AnalysisDataWithContext(context.Background(), request)
}

// DescribeTimingL7AnalysisData
// This API is used to query time series data of the L7 domain name business.
//
// Note:
//
// 1. This API queries data with a delay of about 10 minutes. It is advisable to pull data from 10 minutes before the current time.
//
// 2. This interface returns post-protection traffic request data by default. Users can use `Filters.mitigatedByWebSecurity` to custom query defended data.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeTimingL7AnalysisDataWithContext(ctx context.Context, request *DescribeTimingL7AnalysisDataRequest) (response *DescribeTimingL7AnalysisDataResponse, err error) {
    if request == nil {
        request = NewDescribeTimingL7AnalysisDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeTimingL7AnalysisData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTimingL7AnalysisData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTimingL7AnalysisDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTimingL7CacheDataRequest() (request *DescribeTimingL7CacheDataRequest) {
    request = &DescribeTimingL7CacheDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeTimingL7CacheData")
    
    
    return
}

func NewDescribeTimingL7CacheDataResponse() (response *DescribeTimingL7CacheDataResponse) {
    response = &DescribeTimingL7CacheDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTimingL7CacheData
// This API is used to query L7 cache analysis time series traffic data. This API is to be deprecated. Please use the <a href="https://www.tencentcloud.com/document/product/1552/80648?from_cn_redirect=1">DescribeTimingL7AnalysisData</a> API.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeTimingL7CacheData(request *DescribeTimingL7CacheDataRequest) (response *DescribeTimingL7CacheDataResponse, err error) {
    return c.DescribeTimingL7CacheDataWithContext(context.Background(), request)
}

// DescribeTimingL7CacheData
// This API is used to query L7 cache analysis time series traffic data. This API is to be deprecated. Please use the <a href="https://www.tencentcloud.com/document/product/1552/80648?from_cn_redirect=1">DescribeTimingL7AnalysisData</a> API.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeTimingL7CacheDataWithContext(ctx context.Context, request *DescribeTimingL7CacheDataRequest) (response *DescribeTimingL7CacheDataResponse, err error) {
    if request == nil {
        request = NewDescribeTimingL7CacheDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeTimingL7CacheData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTimingL7CacheData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTimingL7CacheDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTimingL7OriginPullDataRequest() (request *DescribeTimingL7OriginPullDataRequest) {
    request = &DescribeTimingL7OriginPullDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeTimingL7OriginPullData")
    
    
    return
}

func NewDescribeTimingL7OriginPullDataResponse() (response *DescribeTimingL7OriginPullDataResponse) {
    response = &DescribeTimingL7OriginPullDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTimingL7OriginPullData
// This API is used to query the origin-pull time series data of an L7 domain.
//
// You can choose to group queries by specifying the query dimension <code>DimensionName</code>, which returns multiple groups of time series data. For detailed directions and usage limits, see [How to Use APIs to Implement Group Query in a Single Call](https://www.tencentcloud.com/document/product/1552/127501?from_cn_redirect=1).
//
// error code that may be returned:
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETERVALUE_FILTERSMUSTINCLUDEDIMENSIONNAME = "InvalidParameterValue.FiltersMustIncludeDimensionName"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  LIMITEXCEEDED_TIMINGDATAITEMLIMITEXCEEDED = "LimitExceeded.TimingDataItemLimitExceeded"
//  OPERATIONDENIED_ORIGINPULLDATANOTSUPPORTED = "OperationDenied.OriginPullDataNotSupported"
func (c *Client) DescribeTimingL7OriginPullData(request *DescribeTimingL7OriginPullDataRequest) (response *DescribeTimingL7OriginPullDataResponse, err error) {
    return c.DescribeTimingL7OriginPullDataWithContext(context.Background(), request)
}

// DescribeTimingL7OriginPullData
// This API is used to query the origin-pull time series data of an L7 domain.
//
// You can choose to group queries by specifying the query dimension <code>DimensionName</code>, which returns multiple groups of time series data. For detailed directions and usage limits, see [How to Use APIs to Implement Group Query in a Single Call](https://www.tencentcloud.com/document/product/1552/127501?from_cn_redirect=1).
//
// error code that may be returned:
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETERVALUE_FILTERSMUSTINCLUDEDIMENSIONNAME = "InvalidParameterValue.FiltersMustIncludeDimensionName"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  LIMITEXCEEDED_TIMINGDATAITEMLIMITEXCEEDED = "LimitExceeded.TimingDataItemLimitExceeded"
//  OPERATIONDENIED_ORIGINPULLDATANOTSUPPORTED = "OperationDenied.OriginPullDataNotSupported"
func (c *Client) DescribeTimingL7OriginPullDataWithContext(ctx context.Context, request *DescribeTimingL7OriginPullDataRequest) (response *DescribeTimingL7OriginPullDataResponse, err error) {
    if request == nil {
        request = NewDescribeTimingL7OriginPullDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeTimingL7OriginPullData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTimingL7OriginPullData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTimingL7OriginPullDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopFunctionAnalysisDataRequest() (request *DescribeTopFunctionAnalysisDataRequest) {
    request = &DescribeTopFunctionAnalysisDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeTopFunctionAnalysisData")
    
    
    return
}

func NewDescribeTopFunctionAnalysisDataResponse() (response *DescribeTopFunctionAnalysisDataResponse) {
    response = &DescribeTopFunctionAnalysisDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTopFunctionAnalysisData
// This API is used to query the topN data of Edge Functions by a specified dimension. Note that the data returned by this API has a delay of about 10 minutes. We recommend that you query the data generated at least 10 minutes before the current time.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeTopFunctionAnalysisData(request *DescribeTopFunctionAnalysisDataRequest) (response *DescribeTopFunctionAnalysisDataResponse, err error) {
    return c.DescribeTopFunctionAnalysisDataWithContext(context.Background(), request)
}

// DescribeTopFunctionAnalysisData
// This API is used to query the topN data of Edge Functions by a specified dimension. Note that the data returned by this API has a delay of about 10 minutes. We recommend that you query the data generated at least 10 minutes before the current time.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeTopFunctionAnalysisDataWithContext(ctx context.Context, request *DescribeTopFunctionAnalysisDataRequest) (response *DescribeTopFunctionAnalysisDataResponse, err error) {
    if request == nil {
        request = NewDescribeTopFunctionAnalysisDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeTopFunctionAnalysisData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopFunctionAnalysisData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopFunctionAnalysisDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopL7AnalysisDataRequest() (request *DescribeTopL7AnalysisDataRequest) {
    request = &DescribeTopL7AnalysisDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeTopL7AnalysisData")
    
    
    return
}

func NewDescribeTopL7AnalysisDataResponse() (response *DescribeTopL7AnalysisDataResponse) {
    response = &DescribeTopL7AnalysisDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTopL7AnalysisData
// This API is used to query the top N data of the L7 domain name business by specified dimension.
//
// Note:
//
// 1. This API queries data with a delay of about 10 minutes. It is recommended to pull data from 10 minutes before the current time.
//
// 2. This interface returns post-protection traffic request data by default. Users can use `Filters.mitigatedByWebSecurity` to custom query defended data.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeTopL7AnalysisData(request *DescribeTopL7AnalysisDataRequest) (response *DescribeTopL7AnalysisDataResponse, err error) {
    return c.DescribeTopL7AnalysisDataWithContext(context.Background(), request)
}

// DescribeTopL7AnalysisData
// This API is used to query the top N data of the L7 domain name business by specified dimension.
//
// Note:
//
// 1. This API queries data with a delay of about 10 minutes. It is recommended to pull data from 10 minutes before the current time.
//
// 2. This interface returns post-protection traffic request data by default. Users can use `Filters.mitigatedByWebSecurity` to custom query defended data.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeTopL7AnalysisDataWithContext(ctx context.Context, request *DescribeTopL7AnalysisDataRequest) (response *DescribeTopL7AnalysisDataResponse, err error) {
    if request == nil {
        request = NewDescribeTopL7AnalysisDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeTopL7AnalysisData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopL7AnalysisData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopL7AnalysisDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopL7CacheDataRequest() (request *DescribeTopL7CacheDataRequest) {
    request = &DescribeTopL7CacheDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeTopL7CacheData")
    
    
    return
}

func NewDescribeTopL7CacheDataResponse() (response *DescribeTopL7CacheDataResponse) {
    response = &DescribeTopL7CacheDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTopL7CacheData
// This API is used to query L7 cache analysis topN data. This API is to be deprecated. Please use the <a href="https://www.tencentcloud.com/document/product/1552/80646?from_cn_redirect=1"> DescribeTopL7AnalysisData</a> API.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeTopL7CacheData(request *DescribeTopL7CacheDataRequest) (response *DescribeTopL7CacheDataResponse, err error) {
    return c.DescribeTopL7CacheDataWithContext(context.Background(), request)
}

// DescribeTopL7CacheData
// This API is used to query L7 cache analysis topN data. This API is to be deprecated. Please use the <a href="https://www.tencentcloud.com/document/product/1552/80646?from_cn_redirect=1"> DescribeTopL7AnalysisData</a> API.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeTopL7CacheDataWithContext(ctx context.Context, request *DescribeTopL7CacheDataRequest) (response *DescribeTopL7CacheDataResponse, err error) {
    if request == nil {
        request = NewDescribeTopL7CacheDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeTopL7CacheData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopL7CacheData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopL7CacheDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWebSecurityTemplateRequest() (request *DescribeWebSecurityTemplateRequest) {
    request = &DescribeWebSecurityTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeWebSecurityTemplate")
    
    
    return
}

func NewDescribeWebSecurityTemplateResponse() (response *DescribeWebSecurityTemplateResponse) {
    response = &DescribeWebSecurityTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWebSecurityTemplate
// Queries security policy configuration template details
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeWebSecurityTemplate(request *DescribeWebSecurityTemplateRequest) (response *DescribeWebSecurityTemplateResponse, err error) {
    return c.DescribeWebSecurityTemplateWithContext(context.Background(), request)
}

// DescribeWebSecurityTemplate
// Queries security policy configuration template details
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeWebSecurityTemplateWithContext(ctx context.Context, request *DescribeWebSecurityTemplateRequest) (response *DescribeWebSecurityTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeWebSecurityTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeWebSecurityTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWebSecurityTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWebSecurityTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWebSecurityTemplatesRequest() (request *DescribeWebSecurityTemplatesRequest) {
    request = &DescribeWebSecurityTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeWebSecurityTemplates")
    
    
    return
}

func NewDescribeWebSecurityTemplatesResponse() (response *DescribeWebSecurityTemplatesResponse) {
    response = &DescribeWebSecurityTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWebSecurityTemplates
// Queries the security policy configuration template list.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeWebSecurityTemplates(request *DescribeWebSecurityTemplatesRequest) (response *DescribeWebSecurityTemplatesResponse, err error) {
    return c.DescribeWebSecurityTemplatesWithContext(context.Background(), request)
}

// DescribeWebSecurityTemplates
// Queries the security policy configuration template list.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeWebSecurityTemplatesWithContext(ctx context.Context, request *DescribeWebSecurityTemplatesRequest) (response *DescribeWebSecurityTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeWebSecurityTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeWebSecurityTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWebSecurityTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWebSecurityTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZoneConfigImportResultRequest() (request *DescribeZoneConfigImportResultRequest) {
    request = &DescribeZoneConfigImportResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeZoneConfigImportResult")
    
    
    return
}

func NewDescribeZoneConfigImportResultResponse() (response *DescribeZoneConfigImportResultResponse) {
    response = &DescribeZoneConfigImportResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeZoneConfigImportResult
// This API is used to query the results of site configuration item import via the site configuration import API.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeZoneConfigImportResult(request *DescribeZoneConfigImportResultRequest) (response *DescribeZoneConfigImportResultResponse, err error) {
    return c.DescribeZoneConfigImportResultWithContext(context.Background(), request)
}

// DescribeZoneConfigImportResult
// This API is used to query the results of site configuration item import via the site configuration import API.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeZoneConfigImportResultWithContext(ctx context.Context, request *DescribeZoneConfigImportResultRequest) (response *DescribeZoneConfigImportResultResponse, err error) {
    if request == nil {
        request = NewDescribeZoneConfigImportResultRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeZoneConfigImportResult")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZoneConfigImportResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeZoneConfigImportResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZoneCustomVariablesRequest() (request *DescribeZoneCustomVariablesRequest) {
    request = &DescribeZoneCustomVariablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeZoneCustomVariables")
    
    
    return
}

func NewDescribeZoneCustomVariablesResponse() (response *DescribeZoneCustomVariablesResponse) {
    response = &DescribeZoneCustomVariablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeZoneCustomVariables
// This API is used to query site-level custom variable configurations, including variable definitions and variable operations. If not configured, both variable definitions and variable operations return empty lists.
//
// 
//
// If needed, call the DescribeL7AccRules API to query rule-level custom variables.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeZoneCustomVariables(request *DescribeZoneCustomVariablesRequest) (response *DescribeZoneCustomVariablesResponse, err error) {
    return c.DescribeZoneCustomVariablesWithContext(context.Background(), request)
}

// DescribeZoneCustomVariables
// This API is used to query site-level custom variable configurations, including variable definitions and variable operations. If not configured, both variable definitions and variable operations return empty lists.
//
// 
//
// If needed, call the DescribeL7AccRules API to query rule-level custom variables.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeZoneCustomVariablesWithContext(ctx context.Context, request *DescribeZoneCustomVariablesRequest) (response *DescribeZoneCustomVariablesResponse, err error) {
    if request == nil {
        request = NewDescribeZoneCustomVariablesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeZoneCustomVariables")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZoneCustomVariables require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeZoneCustomVariablesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZoneSettingRequest() (request *DescribeZoneSettingRequest) {
    request = &DescribeZoneSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeZoneSetting")
    
    
    return
}

func NewDescribeZoneSettingResponse() (response *DescribeZoneSettingResponse) {
    response = &DescribeZoneSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeZoneSetting
// This API is on an earlier version. EdgeOne has completed a comprehensive upgrade of rule engine APIs. For details, see [DescribeL7AccSetting](https://www.tencentcloud.com/document/product/1552/115819?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_SETTINGINVALIDPARAM = "InvalidParameter.SettingInvalidParam"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeZoneSetting(request *DescribeZoneSettingRequest) (response *DescribeZoneSettingResponse, err error) {
    return c.DescribeZoneSettingWithContext(context.Background(), request)
}

// DescribeZoneSetting
// This API is on an earlier version. EdgeOne has completed a comprehensive upgrade of rule engine APIs. For details, see [DescribeL7AccSetting](https://www.tencentcloud.com/document/product/1552/115819?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_SETTINGINVALIDPARAM = "InvalidParameter.SettingInvalidParam"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeZoneSettingWithContext(ctx context.Context, request *DescribeZoneSettingRequest) (response *DescribeZoneSettingResponse, err error) {
    if request == nil {
        request = NewDescribeZoneSettingRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeZoneSetting")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZoneSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeZoneSettingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZonesRequest() (request *DescribeZonesRequest) {
    request = &DescribeZonesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeZones")
    
    
    return
}

func NewDescribeZonesResponse() (response *DescribeZonesResponse) {
    response = &DescribeZonesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeZones
// This API is used to query site information you have permission to access. You can filter sites based on different query conditions.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeZones(request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    return c.DescribeZonesWithContext(context.Background(), request)
}

// DescribeZones
// This API is used to query site information you have permission to access. You can filter sites based on different query conditions.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeZonesWithContext(ctx context.Context, request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    if request == nil {
        request = NewDescribeZonesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DescribeZones")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZones require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeZonesResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyPlanRequest() (request *DestroyPlanRequest) {
    request = &DestroyPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DestroyPlan")
    
    
    return
}

func NewDestroyPlanResponse() (response *DestroyPlanResponse) {
    response = &DestroyPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DestroyPlan
// This API is used to terminate a billing plan when you need to stop billing for an Edgeone plan.
//
// > The following conditions must be met to terminate a billing plan:
//
// 1. The plan has expired (except for the Enterprise Edition Plan);
//
// 2. All sites under the package have been shut down or deleted.
//
// 
//
// > The site status can be queried through the [Query Site List](https://www.tencentcloud.com/document/product/1552/80713?from_cn_redirect=1) interface.
//
// A site can be deactivated by switching the site to a closed status through the [Switch Site Status](https://www.tencentcloud.com/document/product/1552/80707?from_cn_redirect=1) interface.
//
// A site can be deleted by using the Delete Site API.
//
// error code that may be returned:
//  INVALIDPARAMETER_PLANNOTFOUND = "InvalidParameter.PlanNotFound"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
func (c *Client) DestroyPlan(request *DestroyPlanRequest) (response *DestroyPlanResponse, err error) {
    return c.DestroyPlanWithContext(context.Background(), request)
}

// DestroyPlan
// This API is used to terminate a billing plan when you need to stop billing for an Edgeone plan.
//
// > The following conditions must be met to terminate a billing plan:
//
// 1. The plan has expired (except for the Enterprise Edition Plan);
//
// 2. All sites under the package have been shut down or deleted.
//
// 
//
// > The site status can be queried through the [Query Site List](https://www.tencentcloud.com/document/product/1552/80713?from_cn_redirect=1) interface.
//
// A site can be deactivated by switching the site to a closed status through the [Switch Site Status](https://www.tencentcloud.com/document/product/1552/80707?from_cn_redirect=1) interface.
//
// A site can be deleted by using the Delete Site API.
//
// error code that may be returned:
//  INVALIDPARAMETER_PLANNOTFOUND = "InvalidParameter.PlanNotFound"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
func (c *Client) DestroyPlanWithContext(ctx context.Context, request *DestroyPlanRequest) (response *DestroyPlanResponse, err error) {
    if request == nil {
        request = NewDestroyPlanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DestroyPlan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroyPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroyPlanResponse()
    err = c.Send(request, response)
    return
}

func NewDisableOriginACLRequest() (request *DisableOriginACLRequest) {
    request = &DisableOriginACLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DisableOriginACL")
    
    
    return
}

func NewDisableOriginACLResponse() (response *DisableOriginACLResponse) {
    response = &DisableOriginACLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisableOriginACL
// This API is used to disable origin protection for a site. Once disabled, resources related to the origin will no longer use only the origin IP range provided by origin protection to request your origin, and update notifications on origin IP ranges will stop simultaneously.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DisableOriginACL(request *DisableOriginACLRequest) (response *DisableOriginACLResponse, err error) {
    return c.DisableOriginACLWithContext(context.Background(), request)
}

// DisableOriginACL
// This API is used to disable origin protection for a site. Once disabled, resources related to the origin will no longer use only the origin IP range provided by origin protection to request your origin, and update notifications on origin IP ranges will stop simultaneously.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DisableOriginACLWithContext(ctx context.Context, request *DisableOriginACLRequest) (response *DisableOriginACLResponse, err error) {
    if request == nil {
        request = NewDisableOriginACLRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DisableOriginACL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableOriginACL require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableOriginACLResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadL4LogsRequest() (request *DownloadL4LogsRequest) {
    request = &DownloadL4LogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DownloadL4Logs")
    
    
    return
}

func NewDownloadL4LogsResponse() (response *DownloadL4LogsResponse) {
    response = &DownloadL4LogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DownloadL4Logs
// This API is used to obtain the offline log download link for a Layer 4 proxy business.
//
// **Note:** EO stores offline logs and provides a download service only when the offline log feature is enabled. For specific features, see the [document](https://www.tencentcloud.com/document/product/1552/73660?from_cn_redirect=1).
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DownloadL4Logs(request *DownloadL4LogsRequest) (response *DownloadL4LogsResponse, err error) {
    return c.DownloadL4LogsWithContext(context.Background(), request)
}

// DownloadL4Logs
// This API is used to obtain the offline log download link for a Layer 4 proxy business.
//
// **Note:** EO stores offline logs and provides a download service only when the offline log feature is enabled. For specific features, see the [document](https://www.tencentcloud.com/document/product/1552/73660?from_cn_redirect=1).
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DownloadL4LogsWithContext(ctx context.Context, request *DownloadL4LogsRequest) (response *DownloadL4LogsResponse, err error) {
    if request == nil {
        request = NewDownloadL4LogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DownloadL4Logs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DownloadL4Logs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDownloadL4LogsResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadL7LogsRequest() (request *DownloadL7LogsRequest) {
    request = &DownloadL7LogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DownloadL7Logs")
    
    
    return
}

func NewDownloadL7LogsResponse() (response *DownloadL7LogsResponse) {
    response = &DownloadL7LogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DownloadL7Logs
// This API is used to obtain the offline log download link for an L7 domain.
//
// **Note:** EO stores offline logs and provides a download service only when the offline log feature is enabled. For specific features, see the [document](https://www.tencentcloud.com/document/product/1552/73660?from_cn_redirect=1).
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DownloadL7Logs(request *DownloadL7LogsRequest) (response *DownloadL7LogsResponse, err error) {
    return c.DownloadL7LogsWithContext(context.Background(), request)
}

// DownloadL7Logs
// This API is used to obtain the offline log download link for an L7 domain.
//
// **Note:** EO stores offline logs and provides a download service only when the offline log feature is enabled. For specific features, see the [document](https://www.tencentcloud.com/document/product/1552/73660?from_cn_redirect=1).
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DownloadL7LogsWithContext(ctx context.Context, request *DownloadL7LogsRequest) (response *DownloadL7LogsResponse, err error) {
    if request == nil {
        request = NewDownloadL7LogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DownloadL7Logs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DownloadL7Logs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDownloadL7LogsResponse()
    err = c.Send(request, response)
    return
}

func NewDummyParseZoneFullConfigRequest() (request *DummyParseZoneFullConfigRequest) {
    request = &DummyParseZoneFullConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DummyParseZoneFullConfig")
    
    
    return
}

func NewDummyParseZoneFullConfigResponse() (response *DummyParseZoneFullConfigResponse) {
    response = &DummyParseZoneFullConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DummyParseZoneFullConfig
// This API is used to define the structure of a complete site configuration for reference only. Note: This API does not return actual data.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DummyParseZoneFullConfig(request *DummyParseZoneFullConfigRequest) (response *DummyParseZoneFullConfigResponse, err error) {
    return c.DummyParseZoneFullConfigWithContext(context.Background(), request)
}

// DummyParseZoneFullConfig
// This API is used to define the structure of a complete site configuration for reference only. Note: This API does not return actual data.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DummyParseZoneFullConfigWithContext(ctx context.Context, request *DummyParseZoneFullConfigRequest) (response *DummyParseZoneFullConfigResponse, err error) {
    if request == nil {
        request = NewDummyParseZoneFullConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "DummyParseZoneFullConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DummyParseZoneFullConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDummyParseZoneFullConfigResponse()
    err = c.Send(request, response)
    return
}

func NewEdgeKVDeleteRequest() (request *EdgeKVDeleteRequest) {
    request = &EdgeKVDeleteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "EdgeKVDelete")
    
    
    return
}

func NewEdgeKVDeleteResponse() (response *EdgeKVDeleteResponse) {
    response = &EdgeKVDeleteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EdgeKVDelete
// This API is used to delete one or more key-value pair data in a specified namespace. Batch deletion supported. Data cannot be recovered after deletion.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_KEYSTOOMANY = "InvalidParameter.KeysTooMany"
//  RESOURCEUNAVAILABLE_NAMESPACENOTFOUND = "ResourceUnavailable.NamespaceNotFound"
func (c *Client) EdgeKVDelete(request *EdgeKVDeleteRequest) (response *EdgeKVDeleteResponse, err error) {
    return c.EdgeKVDeleteWithContext(context.Background(), request)
}

// EdgeKVDelete
// This API is used to delete one or more key-value pair data in a specified namespace. Batch deletion supported. Data cannot be recovered after deletion.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_KEYSTOOMANY = "InvalidParameter.KeysTooMany"
//  RESOURCEUNAVAILABLE_NAMESPACENOTFOUND = "ResourceUnavailable.NamespaceNotFound"
func (c *Client) EdgeKVDeleteWithContext(ctx context.Context, request *EdgeKVDeleteRequest) (response *EdgeKVDeleteResponse, err error) {
    if request == nil {
        request = NewEdgeKVDeleteRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "EdgeKVDelete")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EdgeKVDelete require credential")
    }

    request.SetContext(ctx)
    
    response = NewEdgeKVDeleteResponse()
    err = c.Send(request, response)
    return
}

func NewEdgeKVGetRequest() (request *EdgeKVGetRequest) {
    request = &EdgeKVGetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "EdgeKVGet")
    
    
    return
}

func NewEdgeKVGetResponse() (response *EdgeKVGetResponse) {
    response = &EdgeKVGetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EdgeKVGet
// This API is used to batch read key values from a specified namespace. It supports querying up to 20 keys in a single query.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_KEYSTOOMANY = "InvalidParameter.KeysTooMany"
//  RESOURCEUNAVAILABLE_NAMESPACENOTFOUND = "ResourceUnavailable.NamespaceNotFound"
func (c *Client) EdgeKVGet(request *EdgeKVGetRequest) (response *EdgeKVGetResponse, err error) {
    return c.EdgeKVGetWithContext(context.Background(), request)
}

// EdgeKVGet
// This API is used to batch read key values from a specified namespace. It supports querying up to 20 keys in a single query.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_KEYSTOOMANY = "InvalidParameter.KeysTooMany"
//  RESOURCEUNAVAILABLE_NAMESPACENOTFOUND = "ResourceUnavailable.NamespaceNotFound"
func (c *Client) EdgeKVGetWithContext(ctx context.Context, request *EdgeKVGetRequest) (response *EdgeKVGetResponse, err error) {
    if request == nil {
        request = NewEdgeKVGetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "EdgeKVGet")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EdgeKVGet require credential")
    }

    request.SetContext(ctx)
    
    response = NewEdgeKVGetResponse()
    err = c.Send(request, response)
    return
}

func NewEdgeKVListRequest() (request *EdgeKVListRequest) {
    request = &EdgeKVListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "EdgeKVList")
    
    
    return
}

func NewEdgeKVListResponse() (response *EdgeKVListResponse) {
    response = &EdgeKVListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EdgeKVList
// This API is used to list all keys in the specified namespace, with prefix filtering supported. Cursor-based traversal is implemented, and the next cursor is returned for continued querying. It is suitable for traversing all keys in a namespace.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEUNAVAILABLE_NAMESPACENOTFOUND = "ResourceUnavailable.NamespaceNotFound"
func (c *Client) EdgeKVList(request *EdgeKVListRequest) (response *EdgeKVListResponse, err error) {
    return c.EdgeKVListWithContext(context.Background(), request)
}

// EdgeKVList
// This API is used to list all keys in the specified namespace, with prefix filtering supported. Cursor-based traversal is implemented, and the next cursor is returned for continued querying. It is suitable for traversing all keys in a namespace.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEUNAVAILABLE_NAMESPACENOTFOUND = "ResourceUnavailable.NamespaceNotFound"
func (c *Client) EdgeKVListWithContext(ctx context.Context, request *EdgeKVListRequest) (response *EdgeKVListResponse, err error) {
    if request == nil {
        request = NewEdgeKVListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "EdgeKVList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EdgeKVList require credential")
    }

    request.SetContext(ctx)
    
    response = NewEdgeKVListResponse()
    err = c.Send(request, response)
    return
}

func NewEdgeKVPutRequest() (request *EdgeKVPutRequest) {
    request = &EdgeKVPutRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "EdgeKVPut")
    
    
    return
}

func NewEdgeKVPutResponse() (response *EdgeKVPutResponse) {
    response = &EdgeKVPutResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EdgeKVPut
// This API is used to write key-value pair data to a specified namespace, with support for setting expiration time. If the key already exists, the original value is overwritten. If not, a new key-value pair is created.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEUNAVAILABLE_NAMESPACENOTFOUND = "ResourceUnavailable.NamespaceNotFound"
func (c *Client) EdgeKVPut(request *EdgeKVPutRequest) (response *EdgeKVPutResponse, err error) {
    return c.EdgeKVPutWithContext(context.Background(), request)
}

// EdgeKVPut
// This API is used to write key-value pair data to a specified namespace, with support for setting expiration time. If the key already exists, the original value is overwritten. If not, a new key-value pair is created.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEUNAVAILABLE_NAMESPACENOTFOUND = "ResourceUnavailable.NamespaceNotFound"
func (c *Client) EdgeKVPutWithContext(ctx context.Context, request *EdgeKVPutRequest) (response *EdgeKVPutResponse, err error) {
    if request == nil {
        request = NewEdgeKVPutRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "EdgeKVPut")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EdgeKVPut require credential")
    }

    request.SetContext(ctx)
    
    response = NewEdgeKVPutResponse()
    err = c.Send(request, response)
    return
}

func NewEnableOriginACLRequest() (request *EnableOriginACLRequest) {
    request = &EnableOriginACLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "EnableOriginACL")
    
    
    return
}

func NewEnableOriginACLResponse() (response *EnableOriginACLResponse) {
    response = &EnableOriginACLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableOriginACL
// This API is used to enable origin protection for a site for the first time. After it is enabled, EdgeOne will use specific origin-pull IP ranges to pull from the origin for layer-7 acceleration domains/layer-4 proxy instances.
//
// This API supports two modes: **Automatic Protection Mode** and **Specified Protection Mode**. <li>When **Automatic Protection Mode** is used (L7EnableMode/L4EnableMode is set to all), the acceleration domain list and layer-4 proxy instances under the site will be automatically added to the origin protection list in real time without manual intervention;</li><li>When **Specified Protection Mode** is used (L7EnableMode/L4EnableMode is set to specific), you need to manually specify the resources to be added to origin protection. Unspecified resources will not be added to the origin protection list. If resource adjustment is required, you need to actively call the ModifyOriginACL API to make adjustments.</li>
//
// 
//
// Instructions for using the Specified Protection Mode:
//
// <li>The maximum number of layer-7 acceleration domains that can be submitted at a time is 200, and the maximum number of layer-4 proxy instances is 100. Mixed submission of layer-7 acceleration domains/layer-4 proxy instances is supported, with a maximum total of 200 instances;</li>
//
// <li>When more than 200 resources need to be enabled, you need to first configure the resources of the first shard through this API, and then configure the remaining shards through the ModifyOriginACL API. For example, if 500 layer-7 acceleration domains need to be added, first use this API to set 200 domains, and then use the ModifyOriginACL API to configure the remaining 300 domains in two batches, with no more than 200 domains per batch.</li>
//
// 
//
// Note:
//
// - When enabling, accounts on the allowlist can select other origin-pull IP range versions, such as the simplified version, to achieve the effect of pulling from the origin with fewer IP ranges.
//
// - Calling this API is deemed as your agreement to the [Special Agreement on Enabling Origin Protection](https://www.tencentcloud.com/document/product/1552/120141?from_cn_redirect=1);
//
// - Origin-pull IP ranges are subject to change from time to time. EdgeOne will send notifications 14 days, 7 days, 3 days, and 1 day before the change through one or more methods such as in-site messages, SMS, and emails. To ensure that you can receive notifications of origin-pull IP range changes, make sure that you have selected message notifications related to the EdgeOne product in the [Tencent Cloud Message Center Console](https://console.cloud.tencent.com/message) and configured the correct message recipients. For configuration methods, refer to [Message Subscription Management](https://www.tencentcloud.com/document/product/567/43476?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDOMAINS = "InvalidParameter.InvalidDomains"
//  INVALIDPARAMETER_INVALIDPROXIES = "InvalidParameter.InvalidProxies"
//  INVALIDPARAMETER_SHIELDNOTSUPPORTHOSTORIGINWHITELIST = "InvalidParameter.ShieldNotSupportHostOriginWhitelist"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_UNSUPPORTEDPLAN = "OperationDenied.UnsupportedPlan"
//  OPERATIONDENIED_VERSIONCONTROLISGRAYING = "OperationDenied.VersionControlIsGraying"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) EnableOriginACL(request *EnableOriginACLRequest) (response *EnableOriginACLResponse, err error) {
    return c.EnableOriginACLWithContext(context.Background(), request)
}

// EnableOriginACL
// This API is used to enable origin protection for a site for the first time. After it is enabled, EdgeOne will use specific origin-pull IP ranges to pull from the origin for layer-7 acceleration domains/layer-4 proxy instances.
//
// This API supports two modes: **Automatic Protection Mode** and **Specified Protection Mode**. <li>When **Automatic Protection Mode** is used (L7EnableMode/L4EnableMode is set to all), the acceleration domain list and layer-4 proxy instances under the site will be automatically added to the origin protection list in real time without manual intervention;</li><li>When **Specified Protection Mode** is used (L7EnableMode/L4EnableMode is set to specific), you need to manually specify the resources to be added to origin protection. Unspecified resources will not be added to the origin protection list. If resource adjustment is required, you need to actively call the ModifyOriginACL API to make adjustments.</li>
//
// 
//
// Instructions for using the Specified Protection Mode:
//
// <li>The maximum number of layer-7 acceleration domains that can be submitted at a time is 200, and the maximum number of layer-4 proxy instances is 100. Mixed submission of layer-7 acceleration domains/layer-4 proxy instances is supported, with a maximum total of 200 instances;</li>
//
// <li>When more than 200 resources need to be enabled, you need to first configure the resources of the first shard through this API, and then configure the remaining shards through the ModifyOriginACL API. For example, if 500 layer-7 acceleration domains need to be added, first use this API to set 200 domains, and then use the ModifyOriginACL API to configure the remaining 300 domains in two batches, with no more than 200 domains per batch.</li>
//
// 
//
// Note:
//
// - When enabling, accounts on the allowlist can select other origin-pull IP range versions, such as the simplified version, to achieve the effect of pulling from the origin with fewer IP ranges.
//
// - Calling this API is deemed as your agreement to the [Special Agreement on Enabling Origin Protection](https://www.tencentcloud.com/document/product/1552/120141?from_cn_redirect=1);
//
// - Origin-pull IP ranges are subject to change from time to time. EdgeOne will send notifications 14 days, 7 days, 3 days, and 1 day before the change through one or more methods such as in-site messages, SMS, and emails. To ensure that you can receive notifications of origin-pull IP range changes, make sure that you have selected message notifications related to the EdgeOne product in the [Tencent Cloud Message Center Console](https://console.cloud.tencent.com/message) and configured the correct message recipients. For configuration methods, refer to [Message Subscription Management](https://www.tencentcloud.com/document/product/567/43476?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDOMAINS = "InvalidParameter.InvalidDomains"
//  INVALIDPARAMETER_INVALIDPROXIES = "InvalidParameter.InvalidProxies"
//  INVALIDPARAMETER_SHIELDNOTSUPPORTHOSTORIGINWHITELIST = "InvalidParameter.ShieldNotSupportHostOriginWhitelist"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_UNSUPPORTEDPLAN = "OperationDenied.UnsupportedPlan"
//  OPERATIONDENIED_VERSIONCONTROLISGRAYING = "OperationDenied.VersionControlIsGraying"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) EnableOriginACLWithContext(ctx context.Context, request *EnableOriginACLRequest) (response *EnableOriginACLResponse, err error) {
    if request == nil {
        request = NewEnableOriginACLRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "EnableOriginACL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableOriginACL require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableOriginACLResponse()
    err = c.Send(request, response)
    return
}

func NewExportZoneConfigRequest() (request *ExportZoneConfigRequest) {
    request = &ExportZoneConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ExportZoneConfig")
    
    
    return
}

func NewExportZoneConfigResponse() (response *ExportZoneConfigResponse) {
    response = &ExportZoneConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExportZoneConfig
// This API is used to export site configuration by users based on desired configuration items. The exported configuration is used for import via the site configuration import API.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDOMAINS = "InvalidParameter.InvalidDomains"
//  INVALIDPARAMETER_INVALIDPROXIES = "InvalidParameter.InvalidProxies"
//  INVALIDPARAMETER_SHIELDNOTSUPPORTHOSTORIGINWHITELIST = "InvalidParameter.ShieldNotSupportHostOriginWhitelist"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_UNSUPPORTEDPLAN = "OperationDenied.UnsupportedPlan"
//  OPERATIONDENIED_VERSIONCONTROLISGRAYING = "OperationDenied.VersionControlIsGraying"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) ExportZoneConfig(request *ExportZoneConfigRequest) (response *ExportZoneConfigResponse, err error) {
    return c.ExportZoneConfigWithContext(context.Background(), request)
}

// ExportZoneConfig
// This API is used to export site configuration by users based on desired configuration items. The exported configuration is used for import via the site configuration import API.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDOMAINS = "InvalidParameter.InvalidDomains"
//  INVALIDPARAMETER_INVALIDPROXIES = "InvalidParameter.InvalidProxies"
//  INVALIDPARAMETER_SHIELDNOTSUPPORTHOSTORIGINWHITELIST = "InvalidParameter.ShieldNotSupportHostOriginWhitelist"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_UNSUPPORTEDPLAN = "OperationDenied.UnsupportedPlan"
//  OPERATIONDENIED_VERSIONCONTROLISGRAYING = "OperationDenied.VersionControlIsGraying"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) ExportZoneConfigWithContext(ctx context.Context, request *ExportZoneConfigRequest) (response *ExportZoneConfigResponse, err error) {
    if request == nil {
        request = NewExportZoneConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ExportZoneConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportZoneConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportZoneConfigResponse()
    err = c.Send(request, response)
    return
}

func NewHandleFunctionRuntimeEnvironmentRequest() (request *HandleFunctionRuntimeEnvironmentRequest) {
    request = &HandleFunctionRuntimeEnvironmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "HandleFunctionRuntimeEnvironment")
    
    
    return
}

func NewHandleFunctionRuntimeEnvironmentResponse() (response *HandleFunctionRuntimeEnvironmentResponse) {
    response = &HandleFunctionRuntimeEnvironmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// HandleFunctionRuntimeEnvironment
// This API is used to operate the runtime environment of an edge function. It supports related settings for environment variables.
//
// After the environment variables are set, they can be used in the function code and function replica code. For details, see [Edge Functions Referencing Environment Variables](https://www.tencentcloud.com/document/product/1552/109151?from_cn_redirect=1#0151fd9a-8b0e-407b-ae37-54553a60ded6).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FUNCTIONDEPLOYING = "FailedOperation.FunctionDeploying"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_LENGTHEXCEEDSLIMIT = "InvalidParameter.LengthExceedsLimit"
//  RESOURCEUNAVAILABLE_FUNCTIONNOTFOUND = "ResourceUnavailable.FunctionNotFound"
//  RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"
func (c *Client) HandleFunctionRuntimeEnvironment(request *HandleFunctionRuntimeEnvironmentRequest) (response *HandleFunctionRuntimeEnvironmentResponse, err error) {
    return c.HandleFunctionRuntimeEnvironmentWithContext(context.Background(), request)
}

// HandleFunctionRuntimeEnvironment
// This API is used to operate the runtime environment of an edge function. It supports related settings for environment variables.
//
// After the environment variables are set, they can be used in the function code and function replica code. For details, see [Edge Functions Referencing Environment Variables](https://www.tencentcloud.com/document/product/1552/109151?from_cn_redirect=1#0151fd9a-8b0e-407b-ae37-54553a60ded6).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FUNCTIONDEPLOYING = "FailedOperation.FunctionDeploying"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_LENGTHEXCEEDSLIMIT = "InvalidParameter.LengthExceedsLimit"
//  RESOURCEUNAVAILABLE_FUNCTIONNOTFOUND = "ResourceUnavailable.FunctionNotFound"
//  RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"
func (c *Client) HandleFunctionRuntimeEnvironmentWithContext(ctx context.Context, request *HandleFunctionRuntimeEnvironmentRequest) (response *HandleFunctionRuntimeEnvironmentResponse, err error) {
    if request == nil {
        request = NewHandleFunctionRuntimeEnvironmentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "HandleFunctionRuntimeEnvironment")
    
    if c.GetCredential() == nil {
        return nil, errors.New("HandleFunctionRuntimeEnvironment require credential")
    }

    request.SetContext(ctx)
    
    response = NewHandleFunctionRuntimeEnvironmentResponse()
    err = c.Send(request, response)
    return
}

func NewIdentifyZoneRequest() (request *IdentifyZoneRequest) {
    request = &IdentifyZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "IdentifyZone")
    
    
    return
}

func NewIdentifyZoneResponse() (response *IdentifyZoneResponse) {
    response = &IdentifyZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// IdentifyZone
// Used to verify site ownership.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) IdentifyZone(request *IdentifyZoneRequest) (response *IdentifyZoneResponse, err error) {
    return c.IdentifyZoneWithContext(context.Background(), request)
}

// IdentifyZone
// Used to verify site ownership.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) IdentifyZoneWithContext(ctx context.Context, request *IdentifyZoneRequest) (response *IdentifyZoneResponse, err error) {
    if request == nil {
        request = NewIdentifyZoneRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "IdentifyZone")
    
    if c.GetCredential() == nil {
        return nil, errors.New("IdentifyZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewIdentifyZoneResponse()
    err = c.Send(request, response)
    return
}

func NewImportZoneConfigRequest() (request *ImportZoneConfigRequest) {
    request = &ImportZoneConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ImportZoneConfig")
    
    
    return
}

func NewImportZoneConfigResponse() (response *ImportZoneConfigResponse) {
    response = &ImportZoneConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ImportZoneConfig
// This API is used to quickly import site configuration files. After the import is initiated, the API will return the corresponding task ID (TaskId). Users need to use the site configuration import result query API to obtain the results of this import task.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) ImportZoneConfig(request *ImportZoneConfigRequest) (response *ImportZoneConfigResponse, err error) {
    return c.ImportZoneConfigWithContext(context.Background(), request)
}

// ImportZoneConfig
// This API is used to quickly import site configuration files. After the import is initiated, the API will return the corresponding task ID (TaskId). Users need to use the site configuration import result query API to obtain the results of this import task.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) ImportZoneConfigWithContext(ctx context.Context, request *ImportZoneConfigRequest) (response *ImportZoneConfigResponse, err error) {
    if request == nil {
        request = NewImportZoneConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ImportZoneConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImportZoneConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewImportZoneConfigResponse()
    err = c.Send(request, response)
    return
}

func NewIncreasePlanQuotaRequest() (request *IncreasePlanQuotaRequest) {
    request = &IncreasePlanQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "IncreasePlanQuota")
    
    
    return
}

func NewIncreasePlanQuotaResponse() (response *IncreasePlanQuotaResponse) {
    response = &IncreasePlanQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// IncreasePlanQuota
// This API is used to purchase additional quotas when the number of sites bound to your plan, the number of rules under "Web Protection - Custom Rules - Precision Matching Policy", or the rules under "Web Protection - Rate Limiting - Precision Rate Limiting Module" reach the plan's quota limit.
//
// > This interface only supports the Enterprise Edition Plan.
//
// error code that may be returned:
//  FAILEDOPERATION_INSUFFICIENTACCOUNTBALANCE = "FailedOperation.InsufficientAccountBalance"
//  INVALIDPARAMETER_INVALIDQUOTANUMBER = "InvalidParameter.InvalidQuotaNumber"
//  INVALIDPARAMETER_INVALIDQUOTATYPE = "InvalidParameter.InvalidQuotaType"
//  OPERATIONDENIED_PLANHASBEENISOLATED = "OperationDenied.PlanHasBeenIsolated"
//  OPERATIONDENIED_PLANINCREASEPLANQUOTAUNSUPPORTED = "OperationDenied.PlanIncreasePlanQuotaUnsupported"
//  OPERATIONDENIED_RESOURCEHASBEENLOCKED = "OperationDenied.ResourceHasBeenLocked"
func (c *Client) IncreasePlanQuota(request *IncreasePlanQuotaRequest) (response *IncreasePlanQuotaResponse, err error) {
    return c.IncreasePlanQuotaWithContext(context.Background(), request)
}

// IncreasePlanQuota
// This API is used to purchase additional quotas when the number of sites bound to your plan, the number of rules under "Web Protection - Custom Rules - Precision Matching Policy", or the rules under "Web Protection - Rate Limiting - Precision Rate Limiting Module" reach the plan's quota limit.
//
// > This interface only supports the Enterprise Edition Plan.
//
// error code that may be returned:
//  FAILEDOPERATION_INSUFFICIENTACCOUNTBALANCE = "FailedOperation.InsufficientAccountBalance"
//  INVALIDPARAMETER_INVALIDQUOTANUMBER = "InvalidParameter.InvalidQuotaNumber"
//  INVALIDPARAMETER_INVALIDQUOTATYPE = "InvalidParameter.InvalidQuotaType"
//  OPERATIONDENIED_PLANHASBEENISOLATED = "OperationDenied.PlanHasBeenIsolated"
//  OPERATIONDENIED_PLANINCREASEPLANQUOTAUNSUPPORTED = "OperationDenied.PlanIncreasePlanQuotaUnsupported"
//  OPERATIONDENIED_RESOURCEHASBEENLOCKED = "OperationDenied.ResourceHasBeenLocked"
func (c *Client) IncreasePlanQuotaWithContext(ctx context.Context, request *IncreasePlanQuotaRequest) (response *IncreasePlanQuotaResponse, err error) {
    if request == nil {
        request = NewIncreasePlanQuotaRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "IncreasePlanQuota")
    
    if c.GetCredential() == nil {
        return nil, errors.New("IncreasePlanQuota require credential")
    }

    request.SetContext(ctx)
    
    response = NewIncreasePlanQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccelerationDomainRequest() (request *ModifyAccelerationDomainRequest) {
    request = &ModifyAccelerationDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyAccelerationDomain")
    
    
    return
}

func NewModifyAccelerationDomainResponse() (response *ModifyAccelerationDomainResponse) {
    response = &ModifyAccelerationDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAccelerationDomain
// Modify acceleration domain name information
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_CONFLICTHOSTORIGIN = "InvalidParameter.ConflictHostOrigin"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDAWSSECRETKEY = "InvalidParameter.InvalidAwsSecretKey"
//  INVALIDPARAMETER_INVALIDCLIENTIPORIGIN = "InvalidParameter.InvalidClientIpOrigin"
//  INVALIDPARAMETER_INVALIDHTTPS = "InvalidParameter.InvalidHttps"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_INVALIDPRIVATEACCESSPARAMS = "InvalidParameter.InvalidPrivateAccessParams"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGCLIENTIP = "InvalidParameter.InvalidStandardDebugClientIp"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_ORIGINISINNERIP = "InvalidParameter.OriginIsInnerIp"
//  INVALIDPARAMETER_SPACENOTBINDORIGIN = "InvalidParameter.SpaceNotBindOrigin"
//  INVALIDPARAMETER_ZONEISGRAYPUBLISHING = "InvalidParameter.ZoneIsGrayPublishing"
//  INVALIDPARAMETERVALUE_CONFLICTRECORD = "InvalidParameterValue.ConflictRecord"
//  INVALIDPARAMETERVALUE_DOMAINNOTMATCHZONE = "InvalidParameterValue.DomainNotMatchZone"
//  INVALIDPARAMETERVALUE_INVALIDDOMAINSTATUS = "InvalidParameterValue.InvalidDomainStatus"
//  INVALIDPARAMETERVALUE_SITEFAILOVERNOTSUPPORTHOSTORIGINTYPEVOD = "InvalidParameterValue.SiteFailoverNotSupportHostOriginTypeVod"
//  OPERATIONDENIED_DOMAINNOICP = "OperationDenied.DomainNoICP"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  OPERATIONDENIED_VERSIONCONTROLISGRAYING = "OperationDenied.VersionControlIsGraying"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_DNSRECORD = "ResourceInUse.DnsRecord"
//  RESOURCEUNAVAILABLE_DOMAINNOTFOUND = "ResourceUnavailable.DomainNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ModifyAccelerationDomain(request *ModifyAccelerationDomainRequest) (response *ModifyAccelerationDomainResponse, err error) {
    return c.ModifyAccelerationDomainWithContext(context.Background(), request)
}

// ModifyAccelerationDomain
// Modify acceleration domain name information
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_CONFLICTHOSTORIGIN = "InvalidParameter.ConflictHostOrigin"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDAWSSECRETKEY = "InvalidParameter.InvalidAwsSecretKey"
//  INVALIDPARAMETER_INVALIDCLIENTIPORIGIN = "InvalidParameter.InvalidClientIpOrigin"
//  INVALIDPARAMETER_INVALIDHTTPS = "InvalidParameter.InvalidHttps"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_INVALIDPRIVATEACCESSPARAMS = "InvalidParameter.InvalidPrivateAccessParams"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGCLIENTIP = "InvalidParameter.InvalidStandardDebugClientIp"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_ORIGINISINNERIP = "InvalidParameter.OriginIsInnerIp"
//  INVALIDPARAMETER_SPACENOTBINDORIGIN = "InvalidParameter.SpaceNotBindOrigin"
//  INVALIDPARAMETER_ZONEISGRAYPUBLISHING = "InvalidParameter.ZoneIsGrayPublishing"
//  INVALIDPARAMETERVALUE_CONFLICTRECORD = "InvalidParameterValue.ConflictRecord"
//  INVALIDPARAMETERVALUE_DOMAINNOTMATCHZONE = "InvalidParameterValue.DomainNotMatchZone"
//  INVALIDPARAMETERVALUE_INVALIDDOMAINSTATUS = "InvalidParameterValue.InvalidDomainStatus"
//  INVALIDPARAMETERVALUE_SITEFAILOVERNOTSUPPORTHOSTORIGINTYPEVOD = "InvalidParameterValue.SiteFailoverNotSupportHostOriginTypeVod"
//  OPERATIONDENIED_DOMAINNOICP = "OperationDenied.DomainNoICP"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  OPERATIONDENIED_VERSIONCONTROLISGRAYING = "OperationDenied.VersionControlIsGraying"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_DNSRECORD = "ResourceInUse.DnsRecord"
//  RESOURCEUNAVAILABLE_DOMAINNOTFOUND = "ResourceUnavailable.DomainNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ModifyAccelerationDomainWithContext(ctx context.Context, request *ModifyAccelerationDomainRequest) (response *ModifyAccelerationDomainResponse, err error) {
    if request == nil {
        request = NewModifyAccelerationDomainRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyAccelerationDomain")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccelerationDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccelerationDomainResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccelerationDomainStatusesRequest() (request *ModifyAccelerationDomainStatusesRequest) {
    request = &ModifyAccelerationDomainStatusesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyAccelerationDomainStatuses")
    
    
    return
}

func NewModifyAccelerationDomainStatusesResponse() (response *ModifyAccelerationDomainStatusesResponse) {
    response = &ModifyAccelerationDomainStatusesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAccelerationDomainStatuses
// Batch modify accelerated domain status
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_INVALIDERRORPAGE = "InvalidParameter.InvalidErrorPage"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_ZONEISGRAYPUBLISHING = "InvalidParameter.ZoneIsGrayPublishing"
//  INVALIDPARAMETERVALUE_DOMAINNOTMATCHZONE = "InvalidParameterValue.DomainNotMatchZone"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  OPERATIONDENIED_VERSIONCONTROLISGRAYING = "OperationDenied.VersionControlIsGraying"
//  OPERATIONDENIED_VERSIONCONTROLLOCKED = "OperationDenied.VersionControlLocked"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE_DOMAINNOTFOUND = "ResourceUnavailable.DomainNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ModifyAccelerationDomainStatuses(request *ModifyAccelerationDomainStatusesRequest) (response *ModifyAccelerationDomainStatusesResponse, err error) {
    return c.ModifyAccelerationDomainStatusesWithContext(context.Background(), request)
}

// ModifyAccelerationDomainStatuses
// Batch modify accelerated domain status
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_INVALIDERRORPAGE = "InvalidParameter.InvalidErrorPage"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_ZONEISGRAYPUBLISHING = "InvalidParameter.ZoneIsGrayPublishing"
//  INVALIDPARAMETERVALUE_DOMAINNOTMATCHZONE = "InvalidParameterValue.DomainNotMatchZone"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  OPERATIONDENIED_VERSIONCONTROLISGRAYING = "OperationDenied.VersionControlIsGraying"
//  OPERATIONDENIED_VERSIONCONTROLLOCKED = "OperationDenied.VersionControlLocked"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE_DOMAINNOTFOUND = "ResourceUnavailable.DomainNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ModifyAccelerationDomainStatusesWithContext(ctx context.Context, request *ModifyAccelerationDomainStatusesRequest) (response *ModifyAccelerationDomainStatusesResponse, err error) {
    if request == nil {
        request = NewModifyAccelerationDomainStatusesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyAccelerationDomainStatuses")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccelerationDomainStatuses require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccelerationDomainStatusesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAliasDomainRequest() (request *ModifyAliasDomainRequest) {
    request = &ModifyAliasDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyAliasDomain")
    
    
    return
}

func NewModifyAliasDomainResponse() (response *ModifyAliasDomainResponse) {
    response = &ModifyAliasDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAliasDomain
// Modify an alias domain name.
//
// This feature is only supported in the enterprise plan and is currently in beta test. To use it, please [contact us](https://www.tencentcloud.com/contact-us).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CERTNOTMATCHDOMAIN = "InvalidParameter.CertNotMatchDomain"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyAliasDomain(request *ModifyAliasDomainRequest) (response *ModifyAliasDomainResponse, err error) {
    return c.ModifyAliasDomainWithContext(context.Background(), request)
}

// ModifyAliasDomain
// Modify an alias domain name.
//
// This feature is only supported in the enterprise plan and is currently in beta test. To use it, please [contact us](https://www.tencentcloud.com/contact-us).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CERTNOTMATCHDOMAIN = "InvalidParameter.CertNotMatchDomain"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyAliasDomainWithContext(ctx context.Context, request *ModifyAliasDomainRequest) (response *ModifyAliasDomainResponse, err error) {
    if request == nil {
        request = NewModifyAliasDomainRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyAliasDomain")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAliasDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAliasDomainResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAliasDomainStatusRequest() (request *ModifyAliasDomainStatusRequest) {
    request = &ModifyAliasDomainStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyAliasDomainStatus")
    
    
    return
}

func NewModifyAliasDomainStatusResponse() (response *ModifyAliasDomainStatusResponse) {
    response = &ModifyAliasDomainStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAliasDomainStatus
// Modify the status of an alias domain name.
//
// This feature is only supported in the enterprise plan and is currently in beta test. To use it, please [contact us](https://www.tencentcloud.com/contact-us).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) ModifyAliasDomainStatus(request *ModifyAliasDomainStatusRequest) (response *ModifyAliasDomainStatusResponse, err error) {
    return c.ModifyAliasDomainStatusWithContext(context.Background(), request)
}

// ModifyAliasDomainStatus
// Modify the status of an alias domain name.
//
// This feature is only supported in the enterprise plan and is currently in beta test. To use it, please [contact us](https://www.tencentcloud.com/contact-us).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) ModifyAliasDomainStatusWithContext(ctx context.Context, request *ModifyAliasDomainStatusRequest) (response *ModifyAliasDomainStatusResponse, err error) {
    if request == nil {
        request = NewModifyAliasDomainStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyAliasDomainStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAliasDomainStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAliasDomainStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationProxyRequest() (request *ModifyApplicationProxyRequest) {
    request = &ModifyApplicationProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyApplicationProxy")
    
    
    return
}

func NewModifyApplicationProxyResponse() (response *ModifyApplicationProxyResponse) {
    response = &ModifyApplicationProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApplicationProxy
// This API is on an earlier version. If you want to call it, switch to the latest version. For details, refer to [Modify Layer 4 Proxy Instance].
//
// ](https://www.tencentcloud.com/document/product/1552/103411?from_cn_redirect=1) .
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_L4LACKOFRESOURCES = "OperationDenied.L4LackOfResources"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyApplicationProxy(request *ModifyApplicationProxyRequest) (response *ModifyApplicationProxyResponse, err error) {
    return c.ModifyApplicationProxyWithContext(context.Background(), request)
}

// ModifyApplicationProxy
// This API is on an earlier version. If you want to call it, switch to the latest version. For details, refer to [Modify Layer 4 Proxy Instance].
//
// ](https://www.tencentcloud.com/document/product/1552/103411?from_cn_redirect=1) .
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_L4LACKOFRESOURCES = "OperationDenied.L4LackOfResources"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyApplicationProxyWithContext(ctx context.Context, request *ModifyApplicationProxyRequest) (response *ModifyApplicationProxyResponse, err error) {
    if request == nil {
        request = NewModifyApplicationProxyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyApplicationProxy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplicationProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationProxyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationProxyRuleRequest() (request *ModifyApplicationProxyRuleRequest) {
    request = &ModifyApplicationProxyRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyApplicationProxyRule")
    
    
    return
}

func NewModifyApplicationProxyRuleResponse() (response *ModifyApplicationProxyRuleResponse) {
    response = &ModifyApplicationProxyRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApplicationProxyRule
// This API is on an earlier version. If you want to call it, switch to the latest version. For details, see Modifying Layer 4 Proxy Forwarding Rules.
//
// ](https://www.tencentcloud.com/document/product/1552/103410?from_cn_redirect=1) .
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyApplicationProxyRule(request *ModifyApplicationProxyRuleRequest) (response *ModifyApplicationProxyRuleResponse, err error) {
    return c.ModifyApplicationProxyRuleWithContext(context.Background(), request)
}

// ModifyApplicationProxyRule
// This API is on an earlier version. If you want to call it, switch to the latest version. For details, see Modifying Layer 4 Proxy Forwarding Rules.
//
// ](https://www.tencentcloud.com/document/product/1552/103410?from_cn_redirect=1) .
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyApplicationProxyRuleWithContext(ctx context.Context, request *ModifyApplicationProxyRuleRequest) (response *ModifyApplicationProxyRuleResponse, err error) {
    if request == nil {
        request = NewModifyApplicationProxyRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyApplicationProxyRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplicationProxyRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationProxyRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationProxyRuleStatusRequest() (request *ModifyApplicationProxyRuleStatusRequest) {
    request = &ModifyApplicationProxyRuleStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyApplicationProxyRuleStatus")
    
    
    return
}

func NewModifyApplicationProxyRuleStatusResponse() (response *ModifyApplicationProxyRuleStatusResponse) {
    response = &ModifyApplicationProxyRuleStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApplicationProxyRuleStatus
// This API is on an earlier version. If you want to call it, switch to the latest version. For details, see Modifying Layer 4 Proxy Forwarding Rule Status.
//
// ](https://www.tencentcloud.com/document/product/1552/103409?from_cn_redirect=1) .
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyApplicationProxyRuleStatus(request *ModifyApplicationProxyRuleStatusRequest) (response *ModifyApplicationProxyRuleStatusResponse, err error) {
    return c.ModifyApplicationProxyRuleStatusWithContext(context.Background(), request)
}

// ModifyApplicationProxyRuleStatus
// This API is on an earlier version. If you want to call it, switch to the latest version. For details, see Modifying Layer 4 Proxy Forwarding Rule Status.
//
// ](https://www.tencentcloud.com/document/product/1552/103409?from_cn_redirect=1) .
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyApplicationProxyRuleStatusWithContext(ctx context.Context, request *ModifyApplicationProxyRuleStatusRequest) (response *ModifyApplicationProxyRuleStatusResponse, err error) {
    if request == nil {
        request = NewModifyApplicationProxyRuleStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyApplicationProxyRuleStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplicationProxyRuleStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationProxyRuleStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationProxyStatusRequest() (request *ModifyApplicationProxyStatusRequest) {
    request = &ModifyApplicationProxyStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyApplicationProxyStatus")
    
    
    return
}

func NewModifyApplicationProxyStatusResponse() (response *ModifyApplicationProxyStatusResponse) {
    response = &ModifyApplicationProxyStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApplicationProxyStatus
// This API is on an earlier version. If you want to call it, switch to the latest version. For details, see [modify layer 4 proxy instance status](https://www.tencentcloud.com/document/product/1552/103408?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyApplicationProxyStatus(request *ModifyApplicationProxyStatusRequest) (response *ModifyApplicationProxyStatusResponse, err error) {
    return c.ModifyApplicationProxyStatusWithContext(context.Background(), request)
}

// ModifyApplicationProxyStatus
// This API is on an earlier version. If you want to call it, switch to the latest version. For details, see [modify layer 4 proxy instance status](https://www.tencentcloud.com/document/product/1552/103408?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyApplicationProxyStatusWithContext(ctx context.Context, request *ModifyApplicationProxyStatusRequest) (response *ModifyApplicationProxyStatusResponse, err error) {
    if request == nil {
        request = NewModifyApplicationProxyStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyApplicationProxyStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplicationProxyStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationProxyStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyContentIdentifierRequest() (request *ModifyContentIdentifierRequest) {
    request = &ModifyContentIdentifierRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyContentIdentifier")
    
    
    return
}

func NewModifyContentIdentifierResponse() (response *ModifyContentIdentifierResponse) {
    response = &ModifyContentIdentifierResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyContentIdentifier
// Modification content identifier. Only description modification is supported. The feature is open in the allowlist only.
//
// error code that may be returned:
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyContentIdentifier(request *ModifyContentIdentifierRequest) (response *ModifyContentIdentifierResponse, err error) {
    return c.ModifyContentIdentifierWithContext(context.Background(), request)
}

// ModifyContentIdentifier
// Modification content identifier. Only description modification is supported. The feature is open in the allowlist only.
//
// error code that may be returned:
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyContentIdentifierWithContext(ctx context.Context, request *ModifyContentIdentifierRequest) (response *ModifyContentIdentifierResponse, err error) {
    if request == nil {
        request = NewModifyContentIdentifierRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyContentIdentifier")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyContentIdentifier require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyContentIdentifierResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCustomErrorPageRequest() (request *ModifyCustomErrorPageRequest) {
    request = &ModifyCustomErrorPageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyCustomErrorPage")
    
    
    return
}

func NewModifyCustomErrorPageResponse() (response *ModifyCustomErrorPageResponse) {
    response = &ModifyCustomErrorPageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCustomErrorPage
// This API is used to modify a custom response page.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyCustomErrorPage(request *ModifyCustomErrorPageRequest) (response *ModifyCustomErrorPageResponse, err error) {
    return c.ModifyCustomErrorPageWithContext(context.Background(), request)
}

// ModifyCustomErrorPage
// This API is used to modify a custom response page.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyCustomErrorPageWithContext(ctx context.Context, request *ModifyCustomErrorPageRequest) (response *ModifyCustomErrorPageResponse, err error) {
    if request == nil {
        request = NewModifyCustomErrorPageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyCustomErrorPage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCustomErrorPage require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCustomErrorPageResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCustomGroupRequest() (request *ModifyCustomGroupRequest) {
    request = &ModifyCustomGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyCustomGroup")
    
    
    return
}

func NewModifyCustomGroupResponse() (response *ModifyCustomGroupResponse) {
    response = &ModifyCustomGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCustomGroup
// This API is used to modify the content of a custom group. GroupType cannot be modified after creation. When modifying entries in the group, you can specify replacement, append, or removal operations through GroupItemOperation.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyCustomGroup(request *ModifyCustomGroupRequest) (response *ModifyCustomGroupResponse, err error) {
    return c.ModifyCustomGroupWithContext(context.Background(), request)
}

// ModifyCustomGroup
// This API is used to modify the content of a custom group. GroupType cannot be modified after creation. When modifying entries in the group, you can specify replacement, append, or removal operations through GroupItemOperation.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyCustomGroupWithContext(ctx context.Context, request *ModifyCustomGroupRequest) (response *ModifyCustomGroupResponse, err error) {
    if request == nil {
        request = NewModifyCustomGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyCustomGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCustomGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCustomGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDDoSProtectionRequest() (request *ModifyDDoSProtectionRequest) {
    request = &ModifyDDoSProtectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyDDoSProtection")
    
    
    return
}

func NewModifyDDoSProtectionResponse() (response *ModifyDDoSProtectionResponse) {
    response = &ModifyDDoSProtectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDDoSProtection
// Modify exclusive DDoS protection for a site.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  OPERATIONDENIED_SECURITYLACKOFRESOURCES = "OperationDenied.SecurityLackOfResources"
func (c *Client) ModifyDDoSProtection(request *ModifyDDoSProtectionRequest) (response *ModifyDDoSProtectionResponse, err error) {
    return c.ModifyDDoSProtectionWithContext(context.Background(), request)
}

// ModifyDDoSProtection
// Modify exclusive DDoS protection for a site.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  OPERATIONDENIED_SECURITYLACKOFRESOURCES = "OperationDenied.SecurityLackOfResources"
func (c *Client) ModifyDDoSProtectionWithContext(ctx context.Context, request *ModifyDDoSProtectionRequest) (response *ModifyDDoSProtectionResponse, err error) {
    if request == nil {
        request = NewModifyDDoSProtectionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyDDoSProtection")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDDoSProtection require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDDoSProtectionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDnsRecordsRequest() (request *ModifyDnsRecordsRequest) {
    request = &ModifyDnsRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyDnsRecords")
    
    
    return
}

func NewModifyDnsRecordsResponse() (response *ModifyDnsRecordsResponse) {
    response = &ModifyDnsRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDnsRecords
// You can use this API to batch modify DNS records.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  OPERATIONDENIED_SECURITYLACKOFRESOURCES = "OperationDenied.SecurityLackOfResources"
func (c *Client) ModifyDnsRecords(request *ModifyDnsRecordsRequest) (response *ModifyDnsRecordsResponse, err error) {
    return c.ModifyDnsRecordsWithContext(context.Background(), request)
}

// ModifyDnsRecords
// You can use this API to batch modify DNS records.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  OPERATIONDENIED_SECURITYLACKOFRESOURCES = "OperationDenied.SecurityLackOfResources"
func (c *Client) ModifyDnsRecordsWithContext(ctx context.Context, request *ModifyDnsRecordsRequest) (response *ModifyDnsRecordsResponse, err error) {
    if request == nil {
        request = NewModifyDnsRecordsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyDnsRecords")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDnsRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDnsRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDnsRecordsStatusRequest() (request *ModifyDnsRecordsStatusRequest) {
    request = &ModifyDnsRecordsStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyDnsRecordsStatus")
    
    
    return
}

func NewModifyDnsRecordsStatusResponse() (response *ModifyDnsRecordsStatusResponse) {
    response = &ModifyDnsRecordsStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDnsRecordsStatus
// You can use this API to batch modify the status of DNS records and enable or disable records in batches.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  OPERATIONDENIED_SECURITYLACKOFRESOURCES = "OperationDenied.SecurityLackOfResources"
func (c *Client) ModifyDnsRecordsStatus(request *ModifyDnsRecordsStatusRequest) (response *ModifyDnsRecordsStatusResponse, err error) {
    return c.ModifyDnsRecordsStatusWithContext(context.Background(), request)
}

// ModifyDnsRecordsStatus
// You can use this API to batch modify the status of DNS records and enable or disable records in batches.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  OPERATIONDENIED_SECURITYLACKOFRESOURCES = "OperationDenied.SecurityLackOfResources"
func (c *Client) ModifyDnsRecordsStatusWithContext(ctx context.Context, request *ModifyDnsRecordsStatusRequest) (response *ModifyDnsRecordsStatusResponse, err error) {
    if request == nil {
        request = NewModifyDnsRecordsStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyDnsRecordsStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDnsRecordsStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDnsRecordsStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEdgeKVNamespaceRequest() (request *ModifyEdgeKVNamespaceRequest) {
    request = &ModifyEdgeKVNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyEdgeKVNamespace")
    
    
    return
}

func NewModifyEdgeKVNamespaceResponse() (response *ModifyEdgeKVNamespaceResponse) {
    response = &ModifyEdgeKVNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyEdgeKVNamespace
// This API is used to modify the attribute information of a specified KV namespace. Currently, only namespace description modification is supported.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_REMARKTOOLONG = "InvalidParameter.RemarkTooLong"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEUNAVAILABLE_NAMESPACENOTFOUND = "ResourceUnavailable.NamespaceNotFound"
func (c *Client) ModifyEdgeKVNamespace(request *ModifyEdgeKVNamespaceRequest) (response *ModifyEdgeKVNamespaceResponse, err error) {
    return c.ModifyEdgeKVNamespaceWithContext(context.Background(), request)
}

// ModifyEdgeKVNamespace
// This API is used to modify the attribute information of a specified KV namespace. Currently, only namespace description modification is supported.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_REMARKTOOLONG = "InvalidParameter.RemarkTooLong"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEUNAVAILABLE_NAMESPACENOTFOUND = "ResourceUnavailable.NamespaceNotFound"
func (c *Client) ModifyEdgeKVNamespaceWithContext(ctx context.Context, request *ModifyEdgeKVNamespaceRequest) (response *ModifyEdgeKVNamespaceResponse, err error) {
    if request == nil {
        request = NewModifyEdgeKVNamespaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyEdgeKVNamespace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEdgeKVNamespace require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEdgeKVNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyFunctionRequest() (request *ModifyFunctionRequest) {
    request = &ModifyFunctionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyFunction")
    
    
    return
}

func NewModifyFunctionResponse() (response *ModifyFunctionResponse) {
    response = &ModifyFunctionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyFunction
// This API is used to modify an edge function. It supports modifying the function content and description. The function will take effect immediately after modification and redeployment.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FUNCTIONDEPLOYING = "FailedOperation.FunctionDeploying"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_BADCONTENT = "InvalidParameter.BadContent"
//  INVALIDPARAMETER_CONTENTEXCEEDSLIMIT = "InvalidParameter.ContentExceedsLimit"
//  INVALIDPARAMETER_LENGTHEXCEEDSLIMIT = "InvalidParameter.LengthExceedsLimit"
//  INVALIDPARAMETER_MODIFYPARAMETERSMISSING = "InvalidParameter.ModifyParametersMissing"
//  RESOURCEUNAVAILABLE_FUNCTIONNOTFOUND = "ResourceUnavailable.FunctionNotFound"
//  RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"
func (c *Client) ModifyFunction(request *ModifyFunctionRequest) (response *ModifyFunctionResponse, err error) {
    return c.ModifyFunctionWithContext(context.Background(), request)
}

// ModifyFunction
// This API is used to modify an edge function. It supports modifying the function content and description. The function will take effect immediately after modification and redeployment.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FUNCTIONDEPLOYING = "FailedOperation.FunctionDeploying"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_BADCONTENT = "InvalidParameter.BadContent"
//  INVALIDPARAMETER_CONTENTEXCEEDSLIMIT = "InvalidParameter.ContentExceedsLimit"
//  INVALIDPARAMETER_LENGTHEXCEEDSLIMIT = "InvalidParameter.LengthExceedsLimit"
//  INVALIDPARAMETER_MODIFYPARAMETERSMISSING = "InvalidParameter.ModifyParametersMissing"
//  RESOURCEUNAVAILABLE_FUNCTIONNOTFOUND = "ResourceUnavailable.FunctionNotFound"
//  RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"
func (c *Client) ModifyFunctionWithContext(ctx context.Context, request *ModifyFunctionRequest) (response *ModifyFunctionResponse, err error) {
    if request == nil {
        request = NewModifyFunctionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyFunction")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyFunction require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyFunctionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyFunctionComponentBindingsRequest() (request *ModifyFunctionComponentBindingsRequest) {
    request = &ModifyFunctionComponentBindingsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyFunctionComponentBindings")
    
    
    return
}

func NewModifyFunctionComponentBindingsResponse() (response *ModifyFunctionComponentBindingsResponse) {
    response = &ModifyFunctionComponentBindingsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyFunctionComponentBindings
// Modify the binding relationship between edge functions and components. Four operation modes are supported: bind, bind-override, unbind, and rebind. Manage function component binding relationships by specifying the operation type and component list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_BINDINGNOTFOUND = "InvalidParameter.BindingNotFound"
//  INVALIDPARAMETER_DUPLICATEBINDINGNAME = "InvalidParameter.DuplicateBindingName"
//  INVALIDPARAMETER_FUNCTIONBINDVARIABLENAMECONFLICT = "InvalidParameter.FunctionBindVariableNameConflict"
//  INVALIDPARAMETER_INVALIDOPERATION = "InvalidParameter.InvalidOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEUNAVAILABLE_FUNCTIONNOTFOUND = "ResourceUnavailable.FunctionNotFound"
//  RESOURCEUNAVAILABLE_NAMESPACENOTFOUND = "ResourceUnavailable.NamespaceNotFound"
func (c *Client) ModifyFunctionComponentBindings(request *ModifyFunctionComponentBindingsRequest) (response *ModifyFunctionComponentBindingsResponse, err error) {
    return c.ModifyFunctionComponentBindingsWithContext(context.Background(), request)
}

// ModifyFunctionComponentBindings
// Modify the binding relationship between edge functions and components. Four operation modes are supported: bind, bind-override, unbind, and rebind. Manage function component binding relationships by specifying the operation type and component list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_BINDINGNOTFOUND = "InvalidParameter.BindingNotFound"
//  INVALIDPARAMETER_DUPLICATEBINDINGNAME = "InvalidParameter.DuplicateBindingName"
//  INVALIDPARAMETER_FUNCTIONBINDVARIABLENAMECONFLICT = "InvalidParameter.FunctionBindVariableNameConflict"
//  INVALIDPARAMETER_INVALIDOPERATION = "InvalidParameter.InvalidOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEUNAVAILABLE_FUNCTIONNOTFOUND = "ResourceUnavailable.FunctionNotFound"
//  RESOURCEUNAVAILABLE_NAMESPACENOTFOUND = "ResourceUnavailable.NamespaceNotFound"
func (c *Client) ModifyFunctionComponentBindingsWithContext(ctx context.Context, request *ModifyFunctionComponentBindingsRequest) (response *ModifyFunctionComponentBindingsResponse, err error) {
    if request == nil {
        request = NewModifyFunctionComponentBindingsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyFunctionComponentBindings")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyFunctionComponentBindings require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyFunctionComponentBindingsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyFunctionReplicaRequest() (request *ModifyFunctionReplicaRequest) {
    request = &ModifyFunctionReplicaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyFunctionReplica")
    
    
    return
}

func NewModifyFunctionReplicaResponse() (response *ModifyFunctionReplicaResponse) {
    response = &ModifyFunctionReplicaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyFunctionReplica
// This API is used to modify the content and description of a specified edge function replica.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_BINDINGNOTFOUND = "InvalidParameter.BindingNotFound"
//  INVALIDPARAMETER_DUPLICATEBINDINGNAME = "InvalidParameter.DuplicateBindingName"
//  INVALIDPARAMETER_FUNCTIONBINDVARIABLENAMECONFLICT = "InvalidParameter.FunctionBindVariableNameConflict"
//  INVALIDPARAMETER_INVALIDOPERATION = "InvalidParameter.InvalidOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEUNAVAILABLE_FUNCTIONNOTFOUND = "ResourceUnavailable.FunctionNotFound"
//  RESOURCEUNAVAILABLE_NAMESPACENOTFOUND = "ResourceUnavailable.NamespaceNotFound"
func (c *Client) ModifyFunctionReplica(request *ModifyFunctionReplicaRequest) (response *ModifyFunctionReplicaResponse, err error) {
    return c.ModifyFunctionReplicaWithContext(context.Background(), request)
}

// ModifyFunctionReplica
// This API is used to modify the content and description of a specified edge function replica.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_BINDINGNOTFOUND = "InvalidParameter.BindingNotFound"
//  INVALIDPARAMETER_DUPLICATEBINDINGNAME = "InvalidParameter.DuplicateBindingName"
//  INVALIDPARAMETER_FUNCTIONBINDVARIABLENAMECONFLICT = "InvalidParameter.FunctionBindVariableNameConflict"
//  INVALIDPARAMETER_INVALIDOPERATION = "InvalidParameter.InvalidOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEUNAVAILABLE_FUNCTIONNOTFOUND = "ResourceUnavailable.FunctionNotFound"
//  RESOURCEUNAVAILABLE_NAMESPACENOTFOUND = "ResourceUnavailable.NamespaceNotFound"
func (c *Client) ModifyFunctionReplicaWithContext(ctx context.Context, request *ModifyFunctionReplicaRequest) (response *ModifyFunctionReplicaResponse, err error) {
    if request == nil {
        request = NewModifyFunctionReplicaRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyFunctionReplica")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyFunctionReplica require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyFunctionReplicaResponse()
    err = c.Send(request, response)
    return
}

func NewModifyFunctionRuleRequest() (request *ModifyFunctionRuleRequest) {
    request = &ModifyFunctionRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyFunctionRule")
    
    
    return
}

func NewModifyFunctionRuleResponse() (response *ModifyFunctionRuleResponse) {
    response = &ModifyFunctionRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyFunctionRule
// This API is used to modify a trigger rule for an edge function. It supports modifying rule conditions, execution functions, and description. You can first use the DescribeFunctionRules API to get the RuleId of the rule to be modified, then input the modified rule content. The original rule content will be overwritten.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RULEOPERATIONCONFLICT = "FailedOperation.RuleOperationConflict"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_DUPLICATERULE = "InvalidParameter.DuplicateRule"
//  INVALIDPARAMETER_INVALIDCONDITIONS = "InvalidParameter.InvalidConditions"
//  INVALIDPARAMETER_MODIFYPARAMETERSMISSING = "InvalidParameter.ModifyParametersMissing"
//  RESOURCEUNAVAILABLE_FUNCTIONNOTFOUND = "ResourceUnavailable.FunctionNotFound"
//  RESOURCEUNAVAILABLE_RULENOTFOUND = "ResourceUnavailable.RuleNotFound"
func (c *Client) ModifyFunctionRule(request *ModifyFunctionRuleRequest) (response *ModifyFunctionRuleResponse, err error) {
    return c.ModifyFunctionRuleWithContext(context.Background(), request)
}

// ModifyFunctionRule
// This API is used to modify a trigger rule for an edge function. It supports modifying rule conditions, execution functions, and description. You can first use the DescribeFunctionRules API to get the RuleId of the rule to be modified, then input the modified rule content. The original rule content will be overwritten.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RULEOPERATIONCONFLICT = "FailedOperation.RuleOperationConflict"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_DUPLICATERULE = "InvalidParameter.DuplicateRule"
//  INVALIDPARAMETER_INVALIDCONDITIONS = "InvalidParameter.InvalidConditions"
//  INVALIDPARAMETER_MODIFYPARAMETERSMISSING = "InvalidParameter.ModifyParametersMissing"
//  RESOURCEUNAVAILABLE_FUNCTIONNOTFOUND = "ResourceUnavailable.FunctionNotFound"
//  RESOURCEUNAVAILABLE_RULENOTFOUND = "ResourceUnavailable.RuleNotFound"
func (c *Client) ModifyFunctionRuleWithContext(ctx context.Context, request *ModifyFunctionRuleRequest) (response *ModifyFunctionRuleResponse, err error) {
    if request == nil {
        request = NewModifyFunctionRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyFunctionRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyFunctionRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyFunctionRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyFunctionRulePriorityRequest() (request *ModifyFunctionRulePriorityRequest) {
    request = &ModifyFunctionRulePriorityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyFunctionRulePriority")
    
    
    return
}

func NewModifyFunctionRulePriorityResponse() (response *ModifyFunctionRulePriorityResponse) {
    response = &ModifyFunctionRulePriorityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyFunctionRulePriority
// This API is used to modify the priority of trigger rules for an edge function.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RULEOPERATIONCONFLICT = "FailedOperation.RuleOperationConflict"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"
func (c *Client) ModifyFunctionRulePriority(request *ModifyFunctionRulePriorityRequest) (response *ModifyFunctionRulePriorityResponse, err error) {
    return c.ModifyFunctionRulePriorityWithContext(context.Background(), request)
}

// ModifyFunctionRulePriority
// This API is used to modify the priority of trigger rules for an edge function.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RULEOPERATIONCONFLICT = "FailedOperation.RuleOperationConflict"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"
func (c *Client) ModifyFunctionRulePriorityWithContext(ctx context.Context, request *ModifyFunctionRulePriorityRequest) (response *ModifyFunctionRulePriorityResponse, err error) {
    if request == nil {
        request = NewModifyFunctionRulePriorityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyFunctionRulePriority")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyFunctionRulePriority require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyFunctionRulePriorityResponse()
    err = c.Send(request, response)
    return
}

func NewModifyHostsCertificateRequest() (request *ModifyHostsCertificateRequest) {
    request = &ModifyHostsCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyHostsCertificate")
    
    
    return
}

func NewModifyHostsCertificateResponse() (response *ModifyHostsCertificateResponse) {
    response = &ModifyHostsCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyHostsCertificate
// After domain name creation is completed, you can configure a self-owned certificate for the domain name, or alternatively use the free certificate (https://www.tencentcloud.com/document/product/1552/90437?from_cn_redirect=1) offered by EdgeOne.
//
// If you need to configure a self-owned certificate, upload the certificate to the [SSL certificate console](https://console.cloud.tencent.com/certoverview) first, then import the corresponding certificate ID in this API. For details, see [Deploy self-owned certificates to EdgeOne domain names].
//
// ](https://www.tencentcloud.com/document/product/1552/88874?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATEHASEXPIRED = "FailedOperation.CertificateHasExpired"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EDGECLIENTCERTIFICATEHASEXPIRED = "FailedOperation.EdgeClientCertificateHasExpired"
//  FAILEDOPERATION_INVALIDZONESTATUS = "FailedOperation.InvalidZoneStatus"
//  FAILEDOPERATION_MODIFYFAILED = "FailedOperation.ModifyFailed"
//  FAILEDOPERATION_UPSTREAMCLIENTCERTIFICATEHASEXPIRED = "FailedOperation.UpstreamClientCertificateHasExpired"
//  FAILEDOPERATION_UPSTREAMVERIFYCUSTOMCACERTIFICATEHASEXPIRED = "FailedOperation.UpstreamVerifyCustomCACertificateHasExpired"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWERROR = "InternalError.UnknowError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_ALIASDOMAINNOTSUPPORTKEYLESS = "InvalidParameter.AliasDomainNotSupportKeyless"
//  INVALIDPARAMETER_CERTNOTMATCHDOMAIN = "InvalidParameter.CertNotMatchDomain"
//  INVALIDPARAMETER_CERTSANTOOLONG = "InvalidParameter.CertSANTooLong"
//  INVALIDPARAMETER_CERTTOEXPIRE = "InvalidParameter.CertToExpire"
//  INVALIDPARAMETER_CERTTOOSHORTKEYSIZE = "InvalidParameter.CertTooShortKeySize"
//  INVALIDPARAMETER_CERTIFICATECONFLICTWITHKEYLESSSERVER = "InvalidParameter.CertificateConflictWithKeylessServer"
//  INVALIDPARAMETER_CNAMEWILDHOSTNOTALLOWAPPLYCERTIFICATE = "InvalidParameter.CnameWildHostNotAllowApplyCertificate"
//  INVALIDPARAMETER_EDGECLIENTCERTCHECKERROR = "InvalidParameter.EdgeClientCertCheckError"
//  INVALIDPARAMETER_HOSTSTATUSNOTALLOWAPPLYCERTIFICATE = "InvalidParameter.HostStatusNotAllowApplyCertificate"
//  INVALIDPARAMETER_INVALIDCERTINFO = "InvalidParameter.InvalidCertInfo"
//  INVALIDPARAMETER_INVALIDHTTPSCERTINFO = "InvalidParameter.InvalidHttpsCertInfo"
//  INVALIDPARAMETER_INVALIDHTTPSTLSVERSION = "InvalidParameter.InvalidHttpsTlsVersion"
//  INVALIDPARAMETER_UPSTREAMCLIENTCERTCHECKERROR = "InvalidParameter.UpstreamClientCertCheckError"
//  INVALIDPARAMETER_UPSTREAMVERIFYCERTCHECKERROR = "InvalidParameter.UpstreamVerifyCertCheckError"
//  INVALIDPARAMETER_UPSTREAMVERIFYCUSTOMCACERTNOINFO = "InvalidParameter.UpstreamVerifyCustomCACertNoInfo"
//  INVALIDPARAMETER_ZONEISGRAYPUBLISHING = "InvalidParameter.ZoneIsGrayPublishing"
//  INVALIDPARAMETERVALUE_ALIASDOMAINNOTSUPPORTEDGEMTLS = "InvalidParameterValue.AliasDomainNotSupportEdgeMTLS"
//  INVALIDPARAMETERVALUE_ALIASDOMAINNOTSUPPORTUPSTREAMMTLS = "InvalidParameterValue.AliasDomainNotSupportUpstreamMTLS"
//  INVALIDPARAMETERVALUE_ALIASDOMAINNOTSUPPORTUPSTREAMVERIFY = "InvalidParameterValue.AliasDomainNotSupportUpstreamVerify"
//  INVALIDPARAMETERVALUE_CERTIFICATEVERIFYCLIENTMUSTCA = "InvalidParameterValue.CertificateVerifyClientMustCa"
//  INVALIDPARAMETERVALUE_CERTIFICATEVERIFYCLIENTNEEDCERT = "InvalidParameterValue.CertificateVerifyClientNeedCert"
//  INVALIDPARAMETERVALUE_CERTIFICATEVERIFYUPSTREAMCLIENTMUSTRSAORECC = "InvalidParameterValue.CertificateVerifyUpstreamClientMustRSAorECC"
//  INVALIDPARAMETERVALUE_CERTIFICATEVERIFYUPSTREAMCLIENTMUSTSVR = "InvalidParameterValue.CertificateVerifyUpstreamClientMustSVR"
//  INVALIDPARAMETERVALUE_CERTIFICATEVERIFYUPSTREAMCLIENTNEEDCERT = "InvalidParameterValue.CertificateVerifyUpstreamClientNeedCert"
//  INVALIDPARAMETERVALUE_CERTIFICATEVERIFYUPSTREAMVERIFYCUSTOMCAMUSTCA = "InvalidParameterValue.CertificateVerifyUpstreamVerifyCustomCAMustCA"
//  INVALIDPARAMETERVALUE_CERTIFICATEVERIFYUPSTREAMVERIFYCUSTOMCAMUSTRSAORECC = "InvalidParameterValue.CertificateVerifyUpstreamVerifyCustomCAMustRSAorECC"
//  INVALIDPARAMETERVALUE_CERTIFICATEVERIFYUPSTREAMVERIFYCUSTOMCANEEDCERT = "InvalidParameterValue.CertificateVerifyUpstreamVerifyCustomCANeedCert"
//  INVALIDPARAMETERVALUE_CLIENTCERTINFOQUOTALIMIT = "InvalidParameterValue.ClientCertInfoQuotaLimit"
//  INVALIDPARAMETERVALUE_INVALIDKEYLESSSERVERID = "InvalidParameterValue.InvalidKeylessServerId"
//  INVALIDPARAMETERVALUE_OCDIRECTORIGINDOMAINNOTSUPPORTUPSTREAMVERIFY = "InvalidParameterValue.OCDirectOriginDomainNotSupportUpstreamVerify"
//  INVALIDPARAMETERVALUE_SERVERCERTINFONEEDCONTAINRSAORECC = "InvalidParameterValue.ServerCertInfoNeedContainRSAorECC"
//  INVALIDPARAMETERVALUE_SERVERCERTINFONEEDCONTAINSM2 = "InvalidParameterValue.ServerCertInfoNeedContainSM2"
//  INVALIDPARAMETERVALUE_SITEFAILOVERNOTSUPPORTHOSTUPSTREAMVERIFY = "InvalidParameterValue.SiteFailoverNotSupportHostUpstreamVerify"
//  INVALIDPARAMETERVALUE_UPSTREAMCLIENTCERTINFOQUOTALIMIT = "InvalidParameterValue.UpstreamClientCertInfoQuotaLimit"
//  INVALIDPARAMETERVALUE_UPSTREAMVERIFYCUSTOMCACERTINFOQUOTALIMIT = "InvalidParameterValue.UpstreamVerifyCustomCACertInfoQuotaLimit"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_CERTIFICATEPRIVATEKEYISEMPTY = "OperationDenied.CertificatePrivateKeyIsEmpty"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_HOSTSCERTIFICATEINCONSISTENCY = "OperationDenied.HostsCertificateInconsistency"
//  OPERATIONDENIED_HOSTSCLIENTCERTIFICATEINCONSISTENCY = "OperationDenied.HostsClientCertificateInconsistency"
//  OPERATIONDENIED_HOSTSKEYLESSSERVERINCONSISTENCY = "OperationDenied.HostsKeylessServerInconsistency"
//  OPERATIONDENIED_HOSTSUPSTREAMCERTIFICATEINCONSISTENCY = "OperationDenied.HostsUpstreamCertificateInconsistency"
//  OPERATIONDENIED_HOSTSUPSTREAMCERTIFICATEVERIFYINCONSISTENCY = "OperationDenied.HostsUpstreamCertificateVerifyInconsistency"
//  OPERATIONDENIED_KEYLESSCERTSWITCHTOFREECERTCONFLICT = "OperationDenied.KeylessCertSwitchToFreeCertConflict"
//  OPERATIONDENIED_KEYLESSMODECERTIFICATEPRIVATEKEYNEEDEMPTY = "OperationDenied.KeylessModeCertificatePrivateKeyNeedEmpty"
//  OPERATIONDENIED_NOTINKEYLESSWHITELIST = "OperationDenied.NotInKeylessWhiteList"
//  OPERATIONDENIED_NOTINUPSTREAMMTLSWHITELIST = "OperationDenied.NotInUpstreamMTLSWhiteList"
//  OPERATIONDENIED_UNSUPPORTTOCLOSEUPSTREAMMTLS = "OperationDenied.UnSupportToCloseUpstreamMTLS"
//  OPERATIONDENIED_USEUPSTREAMMTLSNEEDOPENHTTPS = "OperationDenied.UseUpstreamMTLSNeedOpenHttps"
//  OPERATIONDENIED_VERSIONCONTROLISGRAYING = "OperationDenied.VersionControlIsGraying"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE_CERTNOTFOUND = "ResourceUnavailable.CertNotFound"
//  RESOURCEUNAVAILABLE_HOSTNOTFOUND = "ResourceUnavailable.HostNotFound"
//  RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) ModifyHostsCertificate(request *ModifyHostsCertificateRequest) (response *ModifyHostsCertificateResponse, err error) {
    return c.ModifyHostsCertificateWithContext(context.Background(), request)
}

// ModifyHostsCertificate
// After domain name creation is completed, you can configure a self-owned certificate for the domain name, or alternatively use the free certificate (https://www.tencentcloud.com/document/product/1552/90437?from_cn_redirect=1) offered by EdgeOne.
//
// If you need to configure a self-owned certificate, upload the certificate to the [SSL certificate console](https://console.cloud.tencent.com/certoverview) first, then import the corresponding certificate ID in this API. For details, see [Deploy self-owned certificates to EdgeOne domain names].
//
// ](https://www.tencentcloud.com/document/product/1552/88874?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATEHASEXPIRED = "FailedOperation.CertificateHasExpired"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EDGECLIENTCERTIFICATEHASEXPIRED = "FailedOperation.EdgeClientCertificateHasExpired"
//  FAILEDOPERATION_INVALIDZONESTATUS = "FailedOperation.InvalidZoneStatus"
//  FAILEDOPERATION_MODIFYFAILED = "FailedOperation.ModifyFailed"
//  FAILEDOPERATION_UPSTREAMCLIENTCERTIFICATEHASEXPIRED = "FailedOperation.UpstreamClientCertificateHasExpired"
//  FAILEDOPERATION_UPSTREAMVERIFYCUSTOMCACERTIFICATEHASEXPIRED = "FailedOperation.UpstreamVerifyCustomCACertificateHasExpired"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWERROR = "InternalError.UnknowError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_ALIASDOMAINNOTSUPPORTKEYLESS = "InvalidParameter.AliasDomainNotSupportKeyless"
//  INVALIDPARAMETER_CERTNOTMATCHDOMAIN = "InvalidParameter.CertNotMatchDomain"
//  INVALIDPARAMETER_CERTSANTOOLONG = "InvalidParameter.CertSANTooLong"
//  INVALIDPARAMETER_CERTTOEXPIRE = "InvalidParameter.CertToExpire"
//  INVALIDPARAMETER_CERTTOOSHORTKEYSIZE = "InvalidParameter.CertTooShortKeySize"
//  INVALIDPARAMETER_CERTIFICATECONFLICTWITHKEYLESSSERVER = "InvalidParameter.CertificateConflictWithKeylessServer"
//  INVALIDPARAMETER_CNAMEWILDHOSTNOTALLOWAPPLYCERTIFICATE = "InvalidParameter.CnameWildHostNotAllowApplyCertificate"
//  INVALIDPARAMETER_EDGECLIENTCERTCHECKERROR = "InvalidParameter.EdgeClientCertCheckError"
//  INVALIDPARAMETER_HOSTSTATUSNOTALLOWAPPLYCERTIFICATE = "InvalidParameter.HostStatusNotAllowApplyCertificate"
//  INVALIDPARAMETER_INVALIDCERTINFO = "InvalidParameter.InvalidCertInfo"
//  INVALIDPARAMETER_INVALIDHTTPSCERTINFO = "InvalidParameter.InvalidHttpsCertInfo"
//  INVALIDPARAMETER_INVALIDHTTPSTLSVERSION = "InvalidParameter.InvalidHttpsTlsVersion"
//  INVALIDPARAMETER_UPSTREAMCLIENTCERTCHECKERROR = "InvalidParameter.UpstreamClientCertCheckError"
//  INVALIDPARAMETER_UPSTREAMVERIFYCERTCHECKERROR = "InvalidParameter.UpstreamVerifyCertCheckError"
//  INVALIDPARAMETER_UPSTREAMVERIFYCUSTOMCACERTNOINFO = "InvalidParameter.UpstreamVerifyCustomCACertNoInfo"
//  INVALIDPARAMETER_ZONEISGRAYPUBLISHING = "InvalidParameter.ZoneIsGrayPublishing"
//  INVALIDPARAMETERVALUE_ALIASDOMAINNOTSUPPORTEDGEMTLS = "InvalidParameterValue.AliasDomainNotSupportEdgeMTLS"
//  INVALIDPARAMETERVALUE_ALIASDOMAINNOTSUPPORTUPSTREAMMTLS = "InvalidParameterValue.AliasDomainNotSupportUpstreamMTLS"
//  INVALIDPARAMETERVALUE_ALIASDOMAINNOTSUPPORTUPSTREAMVERIFY = "InvalidParameterValue.AliasDomainNotSupportUpstreamVerify"
//  INVALIDPARAMETERVALUE_CERTIFICATEVERIFYCLIENTMUSTCA = "InvalidParameterValue.CertificateVerifyClientMustCa"
//  INVALIDPARAMETERVALUE_CERTIFICATEVERIFYCLIENTNEEDCERT = "InvalidParameterValue.CertificateVerifyClientNeedCert"
//  INVALIDPARAMETERVALUE_CERTIFICATEVERIFYUPSTREAMCLIENTMUSTRSAORECC = "InvalidParameterValue.CertificateVerifyUpstreamClientMustRSAorECC"
//  INVALIDPARAMETERVALUE_CERTIFICATEVERIFYUPSTREAMCLIENTMUSTSVR = "InvalidParameterValue.CertificateVerifyUpstreamClientMustSVR"
//  INVALIDPARAMETERVALUE_CERTIFICATEVERIFYUPSTREAMCLIENTNEEDCERT = "InvalidParameterValue.CertificateVerifyUpstreamClientNeedCert"
//  INVALIDPARAMETERVALUE_CERTIFICATEVERIFYUPSTREAMVERIFYCUSTOMCAMUSTCA = "InvalidParameterValue.CertificateVerifyUpstreamVerifyCustomCAMustCA"
//  INVALIDPARAMETERVALUE_CERTIFICATEVERIFYUPSTREAMVERIFYCUSTOMCAMUSTRSAORECC = "InvalidParameterValue.CertificateVerifyUpstreamVerifyCustomCAMustRSAorECC"
//  INVALIDPARAMETERVALUE_CERTIFICATEVERIFYUPSTREAMVERIFYCUSTOMCANEEDCERT = "InvalidParameterValue.CertificateVerifyUpstreamVerifyCustomCANeedCert"
//  INVALIDPARAMETERVALUE_CLIENTCERTINFOQUOTALIMIT = "InvalidParameterValue.ClientCertInfoQuotaLimit"
//  INVALIDPARAMETERVALUE_INVALIDKEYLESSSERVERID = "InvalidParameterValue.InvalidKeylessServerId"
//  INVALIDPARAMETERVALUE_OCDIRECTORIGINDOMAINNOTSUPPORTUPSTREAMVERIFY = "InvalidParameterValue.OCDirectOriginDomainNotSupportUpstreamVerify"
//  INVALIDPARAMETERVALUE_SERVERCERTINFONEEDCONTAINRSAORECC = "InvalidParameterValue.ServerCertInfoNeedContainRSAorECC"
//  INVALIDPARAMETERVALUE_SERVERCERTINFONEEDCONTAINSM2 = "InvalidParameterValue.ServerCertInfoNeedContainSM2"
//  INVALIDPARAMETERVALUE_SITEFAILOVERNOTSUPPORTHOSTUPSTREAMVERIFY = "InvalidParameterValue.SiteFailoverNotSupportHostUpstreamVerify"
//  INVALIDPARAMETERVALUE_UPSTREAMCLIENTCERTINFOQUOTALIMIT = "InvalidParameterValue.UpstreamClientCertInfoQuotaLimit"
//  INVALIDPARAMETERVALUE_UPSTREAMVERIFYCUSTOMCACERTINFOQUOTALIMIT = "InvalidParameterValue.UpstreamVerifyCustomCACertInfoQuotaLimit"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_CERTIFICATEPRIVATEKEYISEMPTY = "OperationDenied.CertificatePrivateKeyIsEmpty"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_HOSTSCERTIFICATEINCONSISTENCY = "OperationDenied.HostsCertificateInconsistency"
//  OPERATIONDENIED_HOSTSCLIENTCERTIFICATEINCONSISTENCY = "OperationDenied.HostsClientCertificateInconsistency"
//  OPERATIONDENIED_HOSTSKEYLESSSERVERINCONSISTENCY = "OperationDenied.HostsKeylessServerInconsistency"
//  OPERATIONDENIED_HOSTSUPSTREAMCERTIFICATEINCONSISTENCY = "OperationDenied.HostsUpstreamCertificateInconsistency"
//  OPERATIONDENIED_HOSTSUPSTREAMCERTIFICATEVERIFYINCONSISTENCY = "OperationDenied.HostsUpstreamCertificateVerifyInconsistency"
//  OPERATIONDENIED_KEYLESSCERTSWITCHTOFREECERTCONFLICT = "OperationDenied.KeylessCertSwitchToFreeCertConflict"
//  OPERATIONDENIED_KEYLESSMODECERTIFICATEPRIVATEKEYNEEDEMPTY = "OperationDenied.KeylessModeCertificatePrivateKeyNeedEmpty"
//  OPERATIONDENIED_NOTINKEYLESSWHITELIST = "OperationDenied.NotInKeylessWhiteList"
//  OPERATIONDENIED_NOTINUPSTREAMMTLSWHITELIST = "OperationDenied.NotInUpstreamMTLSWhiteList"
//  OPERATIONDENIED_UNSUPPORTTOCLOSEUPSTREAMMTLS = "OperationDenied.UnSupportToCloseUpstreamMTLS"
//  OPERATIONDENIED_USEUPSTREAMMTLSNEEDOPENHTTPS = "OperationDenied.UseUpstreamMTLSNeedOpenHttps"
//  OPERATIONDENIED_VERSIONCONTROLISGRAYING = "OperationDenied.VersionControlIsGraying"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE_CERTNOTFOUND = "ResourceUnavailable.CertNotFound"
//  RESOURCEUNAVAILABLE_HOSTNOTFOUND = "ResourceUnavailable.HostNotFound"
//  RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) ModifyHostsCertificateWithContext(ctx context.Context, request *ModifyHostsCertificateRequest) (response *ModifyHostsCertificateResponse, err error) {
    if request == nil {
        request = NewModifyHostsCertificateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyHostsCertificate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyHostsCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyHostsCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInferenceServiceRequest() (request *ModifyInferenceServiceRequest) {
    request = &ModifyInferenceServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyInferenceService")
    
    
    return
}

func NewModifyInferenceServiceResponse() (response *ModifyInferenceServiceResponse) {
    response = &ModifyInferenceServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInferenceService
// Modify an inference service to support updates to the listening port, request path, container image, resource configuration, and description. Only passed parameters will be modified, while unpassed parameters remain unchanged.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_ZONENOTFOUND = "InvalidParameterValue.ZoneNotFound"
//  RESOURCENOTFOUND_INFERENCESERVICE = "ResourceNotFound.InferenceService"
func (c *Client) ModifyInferenceService(request *ModifyInferenceServiceRequest) (response *ModifyInferenceServiceResponse, err error) {
    return c.ModifyInferenceServiceWithContext(context.Background(), request)
}

// ModifyInferenceService
// Modify an inference service to support updates to the listening port, request path, container image, resource configuration, and description. Only passed parameters will be modified, while unpassed parameters remain unchanged.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_ZONENOTFOUND = "InvalidParameterValue.ZoneNotFound"
//  RESOURCENOTFOUND_INFERENCESERVICE = "ResourceNotFound.InferenceService"
func (c *Client) ModifyInferenceServiceWithContext(ctx context.Context, request *ModifyInferenceServiceRequest) (response *ModifyInferenceServiceResponse, err error) {
    if request == nil {
        request = NewModifyInferenceServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyInferenceService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInferenceService require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInferenceServiceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyL4ProxyRequest() (request *ModifyL4ProxyRequest) {
    request = &ModifyL4ProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyL4Proxy")
    
    
    return
}

func NewModifyL4ProxyResponse() (response *ModifyL4ProxyResponse) {
    response = &ModifyL4ProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyL4Proxy
// This API is used to modify the configuration of a layer 4 proxy instance.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDDISABLE = "OperationDenied.AccelerateMainlandDisable"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_L4LACKOFRESOURCES = "OperationDenied.L4LackOfResources"
//  OPERATIONDENIED_L4PROXYINOFFLINESTATUS = "OperationDenied.L4ProxyInOfflineStatus"
//  OPERATIONDENIED_L4PROXYINPROCESSSTATUS = "OperationDenied.L4ProxyInProcessStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) ModifyL4Proxy(request *ModifyL4ProxyRequest) (response *ModifyL4ProxyResponse, err error) {
    return c.ModifyL4ProxyWithContext(context.Background(), request)
}

// ModifyL4Proxy
// This API is used to modify the configuration of a layer 4 proxy instance.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDDISABLE = "OperationDenied.AccelerateMainlandDisable"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_L4LACKOFRESOURCES = "OperationDenied.L4LackOfResources"
//  OPERATIONDENIED_L4PROXYINOFFLINESTATUS = "OperationDenied.L4ProxyInOfflineStatus"
//  OPERATIONDENIED_L4PROXYINPROCESSSTATUS = "OperationDenied.L4ProxyInProcessStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) ModifyL4ProxyWithContext(ctx context.Context, request *ModifyL4ProxyRequest) (response *ModifyL4ProxyResponse, err error) {
    if request == nil {
        request = NewModifyL4ProxyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyL4Proxy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyL4Proxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyL4ProxyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyL4ProxyRulesRequest() (request *ModifyL4ProxyRulesRequest) {
    request = &ModifyL4ProxyRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyL4ProxyRules")
    
    
    return
}

func NewModifyL4ProxyRulesResponse() (response *ModifyL4ProxyRulesResponse) {
    response = &ModifyL4ProxyRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyL4ProxyRules
// This API is used to modify Layer 4 proxy forwarding rules, supporting both individual and batch modification.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyL4ProxyRules(request *ModifyL4ProxyRulesRequest) (response *ModifyL4ProxyRulesResponse, err error) {
    return c.ModifyL4ProxyRulesWithContext(context.Background(), request)
}

// ModifyL4ProxyRules
// This API is used to modify Layer 4 proxy forwarding rules, supporting both individual and batch modification.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyL4ProxyRulesWithContext(ctx context.Context, request *ModifyL4ProxyRulesRequest) (response *ModifyL4ProxyRulesResponse, err error) {
    if request == nil {
        request = NewModifyL4ProxyRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyL4ProxyRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyL4ProxyRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyL4ProxyRulesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyL4ProxyRulesStatusRequest() (request *ModifyL4ProxyRulesStatusRequest) {
    request = &ModifyL4ProxyRulesStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyL4ProxyRulesStatus")
    
    
    return
}

func NewModifyL4ProxyRulesStatusResponse() (response *ModifyL4ProxyRulesStatusResponse) {
    response = &ModifyL4ProxyRulesStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyL4ProxyRulesStatus
// This API is used to enable or disable the status of layer-4 proxy rules, supporting both individual and batch operation.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyL4ProxyRulesStatus(request *ModifyL4ProxyRulesStatusRequest) (response *ModifyL4ProxyRulesStatusResponse, err error) {
    return c.ModifyL4ProxyRulesStatusWithContext(context.Background(), request)
}

// ModifyL4ProxyRulesStatus
// This API is used to enable or disable the status of layer-4 proxy rules, supporting both individual and batch operation.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyL4ProxyRulesStatusWithContext(ctx context.Context, request *ModifyL4ProxyRulesStatusRequest) (response *ModifyL4ProxyRulesStatusResponse, err error) {
    if request == nil {
        request = NewModifyL4ProxyRulesStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyL4ProxyRulesStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyL4ProxyRulesStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyL4ProxyRulesStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyL4ProxyStatusRequest() (request *ModifyL4ProxyStatusRequest) {
    request = &ModifyL4ProxyStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyL4ProxyStatus")
    
    
    return
}

func NewModifyL4ProxyStatusResponse() (response *ModifyL4ProxyStatusResponse) {
    response = &ModifyL4ProxyStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyL4ProxyStatus
// This API is used to enable or disable an L4 proxy instance.
//
// error code that may be returned:
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyL4ProxyStatus(request *ModifyL4ProxyStatusRequest) (response *ModifyL4ProxyStatusResponse, err error) {
    return c.ModifyL4ProxyStatusWithContext(context.Background(), request)
}

// ModifyL4ProxyStatus
// This API is used to enable or disable an L4 proxy instance.
//
// error code that may be returned:
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyL4ProxyStatusWithContext(ctx context.Context, request *ModifyL4ProxyStatusRequest) (response *ModifyL4ProxyStatusResponse, err error) {
    if request == nil {
        request = NewModifyL4ProxyStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyL4ProxyStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyL4ProxyStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyL4ProxyStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyL7AccRuleRequest() (request *ModifyL7AccRuleRequest) {
    request = &ModifyL7AccRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyL7AccRule")
    
    
    return
}

func NewModifyL7AccRuleResponse() (response *ModifyL7AccRuleResponse) {
    response = &ModifyL7AccRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyL7AccRule
// This API is used to modify a rule in the [rule engine](https://www.tencentcloud.com/document/product/1552/70901?from_cn_redirect=1). Only one rule can be modified at a time.
//
// error code that may be returned:
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ADVANCEDORIGINROUTINGNOTSUPPORTACCELERATEMAINLAND = "InvalidParameter.AdvancedOriginRoutingNotSupportAccelerateMainland"
//  INVALIDPARAMETER_ADVANCEDORIGINROUTINGNOTSUPPORTCURRENTPLANORIGINWHITELIST = "InvalidParameter.AdvancedOriginRoutingNotSupportCurrentPlanOriginWhitelist"
//  INVALIDPARAMETER_ADVANCEDORIGINROUTINGNOTSUPPORTMULTIPLYLAYER = "InvalidParameter.AdvancedOriginRoutingNotSupportMultiplyLayer"
//  INVALIDPARAMETER_ADVANCEDORIGINROUTINGNOTSUPPORTOCDIRECTORIGIN = "InvalidParameter.AdvancedOriginRoutingNotSupportOCDirectOrigin"
//  INVALIDPARAMETER_ADVANCEDORIGINROUTINGNOTSUPPORTRACINGPLATFORM = "InvalidParameter.AdvancedOriginRoutingNotSupportRacingPlatform"
//  INVALIDPARAMETER_CACHEKEYQUERYSTRINGTOOMANYVALUE = "InvalidParameter.CacheKeyQueryStringTooManyValue"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_COMPRESSIONINVALIDALGORITHMS = "InvalidParameter.CompressionInvalidAlgorithms"
//  INVALIDPARAMETER_ERRACTIONUNSUPPORTTARGET = "InvalidParameter.ErrActionUnsupportTarget"
//  INVALIDPARAMETER_ERRINVALIDACTION = "InvalidParameter.ErrInvalidAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONDUPLICATEACTION = "InvalidParameter.ErrInvalidActionDuplicateAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMACTION = "InvalidParameter.ErrInvalidActionParamAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMDUPLICATENAME = "InvalidParameter.ErrInvalidActionParamDuplicateName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMNAME = "InvalidParameter.ErrInvalidActionParamName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMTOOMANYVALUES = "InvalidParameter.ErrInvalidActionParamTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_ERRINVALIDACTIONTYPE = "InvalidParameter.ErrInvalidActionType"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONHOSTTOOMANYWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidConditionHostTooManyWhenModifyOriginActionConfigured"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONIGNORECASE = "InvalidParameter.ErrInvalidConditionIgnoreCase"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMEBADNAME = "InvalidParameter.ErrInvalidConditionNameBadName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMETARGETNOTSUPPORTNAME = "InvalidParameter.ErrInvalidConditionNameTargetNotSupportName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADREGULAR = "InvalidParameter.ErrInvalidConditionValueBadRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADURL = "InvalidParameter.ErrInvalidConditionValueBadUrl"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUECONTAINFILENAMEEXTENSION = "InvalidParameter.ErrInvalidConditionValueBadValueContainFileNameExtension"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOLONGVALUE = "InvalidParameter.ErrInvalidConditionValueTooLongValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYVALUES = "InvalidParameter.ErrInvalidConditionValueTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYWILDCARD = "InvalidParameter.ErrInvalidConditionValueTooManyWildcard"
//  INVALIDPARAMETER_ERRINVALIDELSEWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidElseWhenModifyOriginActionConfigured"
//  INVALIDPARAMETER_ERRNILCONDITION = "InvalidParameter.ErrNilCondition"
//  INVALIDPARAMETER_GRPCREQUIREHTTP2 = "InvalidParameter.GrpcRequireHttp2"
//  INVALIDPARAMETER_HOSTNOTFOUND = "InvalidParameter.HostNotFound"
//  INVALIDPARAMETER_INVALIDADVANCEDORIGINROUTINGUNSUPPORTED = "InvalidParameter.InvalidAdvancedOriginRoutingUnsupported"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESECRETKEY = "InvalidParameter.InvalidAuthenticationTypeSecretKey"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESIGNPARAM = "InvalidParameter.InvalidAuthenticationTypeSignParam"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPETIMEFORMAT = "InvalidParameter.InvalidAuthenticationTypeTimeFormat"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPETIMEPARAM = "InvalidParameter.InvalidAuthenticationTypeTimeParam"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDBACKUPSERVERNAME = "InvalidParameter.InvalidBackupServerName"
//  INVALIDPARAMETER_INVALIDCACHEKEY = "InvalidParameter.InvalidCacheKey"
//  INVALIDPARAMETER_INVALIDCACHEKEYCOOKIE = "InvalidParameter.InvalidCacheKeyCookie"
//  INVALIDPARAMETER_INVALIDCACHEKEYIGNORECASE = "InvalidParameter.InvalidCacheKeyIgnoreCase"
//  INVALIDPARAMETER_INVALIDCACHEKEYSCHEME = "InvalidParameter.InvalidCacheKeyScheme"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDCOSDOMAIN = "InvalidParameter.InvalidCosDomain"
//  INVALIDPARAMETER_INVALIDERRORPAGEREDIRECTURL = "InvalidParameter.InvalidErrorPageRedirectUrl"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGETIME = "InvalidParameter.InvalidMaxAgeTime"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAMEXFF = "InvalidParameter.InvalidRequestHeaderNameXff"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERVALUE = "InvalidParameter.InvalidRequestHeaderValue"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERNAME = "InvalidParameter.InvalidResponseHeaderName"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERVALUE = "InvalidParameter.InvalidResponseHeaderValue"
//  INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"
//  INVALIDPARAMETER_INVALIDRULEENGINENOTFOUND = "InvalidParameter.InvalidRuleEngineNotFound"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSEXTENSION = "InvalidParameter.InvalidRuleEngineTargetsExtension"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSURL = "InvalidParameter.InvalidRuleEngineTargetsUrl"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_INVALIDSHIELDUNSUPPORTED = "InvalidParameter.InvalidShieldUnsupported"
//  INVALIDPARAMETER_INVALIDUPSTREAMREQUESTQUERYSTRINGVALUE = "InvalidParameter.InvalidUpstreamRequestQueryStringValue"
//  INVALIDPARAMETER_INVALIDURLREDIRECTHOST = "InvalidParameter.InvalidUrlRedirectHost"
//  INVALIDPARAMETER_INVALIDURLREDIRECTURL = "InvalidParameter.InvalidUrlRedirectUrl"
//  INVALIDPARAMETER_KEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.KeyRulesInvalidQueryStringValue"
//  INVALIDPARAMETER_LOADBALANCEINSTANCEIDISREQUIRED = "InvalidParameter.LoadBalanceInstanceIdIsRequired"
//  INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"
//  INVALIDPARAMETER_ORIGINORIGINGROUPIDISREQUIRED = "InvalidParameter.OriginOriginGroupIdIsRequired"
//  INVALIDPARAMETER_ORIGINPULLPROTOCOLISREQUIRED = "InvalidParameter.OriginPullProtocolIsRequired"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_RESPONSEHEADERCACHECONTROLNOTALLOWDELETE = "InvalidParameter.ResponseHeaderCacheControlNotAllowDelete"
//  INVALIDPARAMETER_SHIELDNOTSUPPORTHOSTCACHEPREFRESH = "InvalidParameter.ShieldNotSupportHostCachePrefresh"
//  INVALIDPARAMETER_SHIELDNOTSUPPORTHOSTORIGINWHITELIST = "InvalidParameter.ShieldNotSupportHostOriginWhitelist"
//  INVALIDPARAMETER_SHIELDNOTSUPPORTHOSTSMARTROUTING = "InvalidParameter.ShieldNotSupportHostSmartRouting"
//  INVALIDPARAMETER_SHIELDNOTSUPPORTZONECACHEPREFRESH = "InvalidParameter.ShieldNotSupportZoneCachePrefresh"
//  INVALIDPARAMETER_SHIELDNOTSUPPORTZONESMARTROUTING = "InvalidParameter.ShieldNotSupportZoneSmartRouting"
//  INVALIDPARAMETER_SHIELDSPACENOTFOUND = "InvalidParameter.ShieldSpaceNotFound"
//  INVALIDPARAMETER_SHIELDSPACEREQUIREDERROR = "InvalidParameter.ShieldSpaceRequiredError"
//  INVALIDPARAMETER_SHIELDSPACESTATUSMUSTONLINE = "InvalidParameter.ShieldSpaceStatusMustOnline"
//  INVALIDPARAMETER_STATUSCODECACHEINVALIDSTATUSCODE = "InvalidParameter.StatusCodeCacheInvalidStatusCode"
//  INVALIDPARAMETER_TLSVERSIONNOTINSEQUENCE = "InvalidParameter.TlsVersionNotInSequence"
//  INVALIDPARAMETERVALUE_FORMATMISMATCH = "InvalidParameterValue.FormatMismatch"
//  INVALIDPARAMETERVALUE_GENERALMISMATCH = "InvalidParameterValue.GeneralMismatch"
//  INVALIDPARAMETERVALUE_INCLUDEINVALIDVALUE = "InvalidParameterValue.IncludeInvalidValue"
//  INVALIDPARAMETERVALUE_INVALIDSITEFAILOVERUNSUPPORTED = "InvalidParameterValue.InvalidSiteFailoverUnsupported"
//  INVALIDPARAMETERVALUE_MISSINGNECESSARYPARAM = "InvalidParameterValue.MissingNecessaryParam"
//  INVALIDPARAMETERVALUE_NOTINENUMERATION = "InvalidParameterValue.NotInEnumeration"
//  INVALIDPARAMETERVALUE_NOTWITHINRANGE = "InvalidParameterValue.NotWithinRange"
//  INVALIDPARAMETERVALUE_REGEXMISMATCH = "InvalidParameterValue.RegExMismatch"
//  INVALIDPARAMETERVALUE_SITEFAILOVERNOTSUPPORTHOSTORIGINPULLRATELIMIT = "InvalidParameterValue.SiteFailoverNotSupportHostOriginPullRateLimit"
//  INVALIDPARAMETERVALUE_SITEFAILOVERNOTSUPPORTHOSTORIGINTYPEVOD = "InvalidParameterValue.SiteFailoverNotSupportHostOriginTypeVod"
//  INVALIDPARAMETERVALUE_SITEFAILOVERNOTSUPPORTHOSTUPSTREAMVERIFY = "InvalidParameterValue.SiteFailoverNotSupportHostUpstreamVerify"
//  INVALIDPARAMETERVALUE_TRIALPLANRESPONSEPAGE = "InvalidParameterValue.TrialPlanResponsePage"
//  INVALIDPARAMETERVALUE_UNRECOGNIZABLEVALUE = "InvalidParameterValue.UnrecognizableValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ADVANCEDORIGINROUTINGCONFLICTWITHDCIOVERSEALAN = "OperationDenied.AdvancedOriginRoutingConflictWithDciOverseaLan"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_INVALIDADVANCEDDEFENSESECURITYTYPE = "OperationDenied.InvalidAdvancedDefenseSecurityType"
//  OPERATIONDENIED_INVALIDADVANCEDORIGINROUTINGBILLING = "OperationDenied.InvalidAdvancedOriginRoutingBilling"
//  OPERATIONDENIED_NOTINADVANCEDORIGINROUTINGWHITELIST = "OperationDenied.NotInAdvancedOriginRoutingWhiteList"
//  OPERATIONDENIED_NOTINSHIELDSPACEWHITELIST = "OperationDenied.NotInShieldSpaceWhiteList"
//  OPERATIONDENIED_NOTINSITEFAILOVERWHITELIST = "OperationDenied.NotInSiteFailoverWhiteList"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ModifyL7AccRule(request *ModifyL7AccRuleRequest) (response *ModifyL7AccRuleResponse, err error) {
    return c.ModifyL7AccRuleWithContext(context.Background(), request)
}

// ModifyL7AccRule
// This API is used to modify a rule in the [rule engine](https://www.tencentcloud.com/document/product/1552/70901?from_cn_redirect=1). Only one rule can be modified at a time.
//
// error code that may be returned:
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ADVANCEDORIGINROUTINGNOTSUPPORTACCELERATEMAINLAND = "InvalidParameter.AdvancedOriginRoutingNotSupportAccelerateMainland"
//  INVALIDPARAMETER_ADVANCEDORIGINROUTINGNOTSUPPORTCURRENTPLANORIGINWHITELIST = "InvalidParameter.AdvancedOriginRoutingNotSupportCurrentPlanOriginWhitelist"
//  INVALIDPARAMETER_ADVANCEDORIGINROUTINGNOTSUPPORTMULTIPLYLAYER = "InvalidParameter.AdvancedOriginRoutingNotSupportMultiplyLayer"
//  INVALIDPARAMETER_ADVANCEDORIGINROUTINGNOTSUPPORTOCDIRECTORIGIN = "InvalidParameter.AdvancedOriginRoutingNotSupportOCDirectOrigin"
//  INVALIDPARAMETER_ADVANCEDORIGINROUTINGNOTSUPPORTRACINGPLATFORM = "InvalidParameter.AdvancedOriginRoutingNotSupportRacingPlatform"
//  INVALIDPARAMETER_CACHEKEYQUERYSTRINGTOOMANYVALUE = "InvalidParameter.CacheKeyQueryStringTooManyValue"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_COMPRESSIONINVALIDALGORITHMS = "InvalidParameter.CompressionInvalidAlgorithms"
//  INVALIDPARAMETER_ERRACTIONUNSUPPORTTARGET = "InvalidParameter.ErrActionUnsupportTarget"
//  INVALIDPARAMETER_ERRINVALIDACTION = "InvalidParameter.ErrInvalidAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONDUPLICATEACTION = "InvalidParameter.ErrInvalidActionDuplicateAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMACTION = "InvalidParameter.ErrInvalidActionParamAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMDUPLICATENAME = "InvalidParameter.ErrInvalidActionParamDuplicateName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMNAME = "InvalidParameter.ErrInvalidActionParamName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMTOOMANYVALUES = "InvalidParameter.ErrInvalidActionParamTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_ERRINVALIDACTIONTYPE = "InvalidParameter.ErrInvalidActionType"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONHOSTTOOMANYWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidConditionHostTooManyWhenModifyOriginActionConfigured"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONIGNORECASE = "InvalidParameter.ErrInvalidConditionIgnoreCase"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMEBADNAME = "InvalidParameter.ErrInvalidConditionNameBadName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMETARGETNOTSUPPORTNAME = "InvalidParameter.ErrInvalidConditionNameTargetNotSupportName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADREGULAR = "InvalidParameter.ErrInvalidConditionValueBadRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADURL = "InvalidParameter.ErrInvalidConditionValueBadUrl"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUECONTAINFILENAMEEXTENSION = "InvalidParameter.ErrInvalidConditionValueBadValueContainFileNameExtension"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOLONGVALUE = "InvalidParameter.ErrInvalidConditionValueTooLongValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYVALUES = "InvalidParameter.ErrInvalidConditionValueTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYWILDCARD = "InvalidParameter.ErrInvalidConditionValueTooManyWildcard"
//  INVALIDPARAMETER_ERRINVALIDELSEWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidElseWhenModifyOriginActionConfigured"
//  INVALIDPARAMETER_ERRNILCONDITION = "InvalidParameter.ErrNilCondition"
//  INVALIDPARAMETER_GRPCREQUIREHTTP2 = "InvalidParameter.GrpcRequireHttp2"
//  INVALIDPARAMETER_HOSTNOTFOUND = "InvalidParameter.HostNotFound"
//  INVALIDPARAMETER_INVALIDADVANCEDORIGINROUTINGUNSUPPORTED = "InvalidParameter.InvalidAdvancedOriginRoutingUnsupported"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESECRETKEY = "InvalidParameter.InvalidAuthenticationTypeSecretKey"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESIGNPARAM = "InvalidParameter.InvalidAuthenticationTypeSignParam"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPETIMEFORMAT = "InvalidParameter.InvalidAuthenticationTypeTimeFormat"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPETIMEPARAM = "InvalidParameter.InvalidAuthenticationTypeTimeParam"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDBACKUPSERVERNAME = "InvalidParameter.InvalidBackupServerName"
//  INVALIDPARAMETER_INVALIDCACHEKEY = "InvalidParameter.InvalidCacheKey"
//  INVALIDPARAMETER_INVALIDCACHEKEYCOOKIE = "InvalidParameter.InvalidCacheKeyCookie"
//  INVALIDPARAMETER_INVALIDCACHEKEYIGNORECASE = "InvalidParameter.InvalidCacheKeyIgnoreCase"
//  INVALIDPARAMETER_INVALIDCACHEKEYSCHEME = "InvalidParameter.InvalidCacheKeyScheme"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDCOSDOMAIN = "InvalidParameter.InvalidCosDomain"
//  INVALIDPARAMETER_INVALIDERRORPAGEREDIRECTURL = "InvalidParameter.InvalidErrorPageRedirectUrl"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGETIME = "InvalidParameter.InvalidMaxAgeTime"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAMEXFF = "InvalidParameter.InvalidRequestHeaderNameXff"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERVALUE = "InvalidParameter.InvalidRequestHeaderValue"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERNAME = "InvalidParameter.InvalidResponseHeaderName"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERVALUE = "InvalidParameter.InvalidResponseHeaderValue"
//  INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"
//  INVALIDPARAMETER_INVALIDRULEENGINENOTFOUND = "InvalidParameter.InvalidRuleEngineNotFound"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSEXTENSION = "InvalidParameter.InvalidRuleEngineTargetsExtension"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSURL = "InvalidParameter.InvalidRuleEngineTargetsUrl"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_INVALIDSHIELDUNSUPPORTED = "InvalidParameter.InvalidShieldUnsupported"
//  INVALIDPARAMETER_INVALIDUPSTREAMREQUESTQUERYSTRINGVALUE = "InvalidParameter.InvalidUpstreamRequestQueryStringValue"
//  INVALIDPARAMETER_INVALIDURLREDIRECTHOST = "InvalidParameter.InvalidUrlRedirectHost"
//  INVALIDPARAMETER_INVALIDURLREDIRECTURL = "InvalidParameter.InvalidUrlRedirectUrl"
//  INVALIDPARAMETER_KEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.KeyRulesInvalidQueryStringValue"
//  INVALIDPARAMETER_LOADBALANCEINSTANCEIDISREQUIRED = "InvalidParameter.LoadBalanceInstanceIdIsRequired"
//  INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"
//  INVALIDPARAMETER_ORIGINORIGINGROUPIDISREQUIRED = "InvalidParameter.OriginOriginGroupIdIsRequired"
//  INVALIDPARAMETER_ORIGINPULLPROTOCOLISREQUIRED = "InvalidParameter.OriginPullProtocolIsRequired"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_RESPONSEHEADERCACHECONTROLNOTALLOWDELETE = "InvalidParameter.ResponseHeaderCacheControlNotAllowDelete"
//  INVALIDPARAMETER_SHIELDNOTSUPPORTHOSTCACHEPREFRESH = "InvalidParameter.ShieldNotSupportHostCachePrefresh"
//  INVALIDPARAMETER_SHIELDNOTSUPPORTHOSTORIGINWHITELIST = "InvalidParameter.ShieldNotSupportHostOriginWhitelist"
//  INVALIDPARAMETER_SHIELDNOTSUPPORTHOSTSMARTROUTING = "InvalidParameter.ShieldNotSupportHostSmartRouting"
//  INVALIDPARAMETER_SHIELDNOTSUPPORTZONECACHEPREFRESH = "InvalidParameter.ShieldNotSupportZoneCachePrefresh"
//  INVALIDPARAMETER_SHIELDNOTSUPPORTZONESMARTROUTING = "InvalidParameter.ShieldNotSupportZoneSmartRouting"
//  INVALIDPARAMETER_SHIELDSPACENOTFOUND = "InvalidParameter.ShieldSpaceNotFound"
//  INVALIDPARAMETER_SHIELDSPACEREQUIREDERROR = "InvalidParameter.ShieldSpaceRequiredError"
//  INVALIDPARAMETER_SHIELDSPACESTATUSMUSTONLINE = "InvalidParameter.ShieldSpaceStatusMustOnline"
//  INVALIDPARAMETER_STATUSCODECACHEINVALIDSTATUSCODE = "InvalidParameter.StatusCodeCacheInvalidStatusCode"
//  INVALIDPARAMETER_TLSVERSIONNOTINSEQUENCE = "InvalidParameter.TlsVersionNotInSequence"
//  INVALIDPARAMETERVALUE_FORMATMISMATCH = "InvalidParameterValue.FormatMismatch"
//  INVALIDPARAMETERVALUE_GENERALMISMATCH = "InvalidParameterValue.GeneralMismatch"
//  INVALIDPARAMETERVALUE_INCLUDEINVALIDVALUE = "InvalidParameterValue.IncludeInvalidValue"
//  INVALIDPARAMETERVALUE_INVALIDSITEFAILOVERUNSUPPORTED = "InvalidParameterValue.InvalidSiteFailoverUnsupported"
//  INVALIDPARAMETERVALUE_MISSINGNECESSARYPARAM = "InvalidParameterValue.MissingNecessaryParam"
//  INVALIDPARAMETERVALUE_NOTINENUMERATION = "InvalidParameterValue.NotInEnumeration"
//  INVALIDPARAMETERVALUE_NOTWITHINRANGE = "InvalidParameterValue.NotWithinRange"
//  INVALIDPARAMETERVALUE_REGEXMISMATCH = "InvalidParameterValue.RegExMismatch"
//  INVALIDPARAMETERVALUE_SITEFAILOVERNOTSUPPORTHOSTORIGINPULLRATELIMIT = "InvalidParameterValue.SiteFailoverNotSupportHostOriginPullRateLimit"
//  INVALIDPARAMETERVALUE_SITEFAILOVERNOTSUPPORTHOSTORIGINTYPEVOD = "InvalidParameterValue.SiteFailoverNotSupportHostOriginTypeVod"
//  INVALIDPARAMETERVALUE_SITEFAILOVERNOTSUPPORTHOSTUPSTREAMVERIFY = "InvalidParameterValue.SiteFailoverNotSupportHostUpstreamVerify"
//  INVALIDPARAMETERVALUE_TRIALPLANRESPONSEPAGE = "InvalidParameterValue.TrialPlanResponsePage"
//  INVALIDPARAMETERVALUE_UNRECOGNIZABLEVALUE = "InvalidParameterValue.UnrecognizableValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ADVANCEDORIGINROUTINGCONFLICTWITHDCIOVERSEALAN = "OperationDenied.AdvancedOriginRoutingConflictWithDciOverseaLan"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_INVALIDADVANCEDDEFENSESECURITYTYPE = "OperationDenied.InvalidAdvancedDefenseSecurityType"
//  OPERATIONDENIED_INVALIDADVANCEDORIGINROUTINGBILLING = "OperationDenied.InvalidAdvancedOriginRoutingBilling"
//  OPERATIONDENIED_NOTINADVANCEDORIGINROUTINGWHITELIST = "OperationDenied.NotInAdvancedOriginRoutingWhiteList"
//  OPERATIONDENIED_NOTINSHIELDSPACEWHITELIST = "OperationDenied.NotInShieldSpaceWhiteList"
//  OPERATIONDENIED_NOTINSITEFAILOVERWHITELIST = "OperationDenied.NotInSiteFailoverWhiteList"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ModifyL7AccRuleWithContext(ctx context.Context, request *ModifyL7AccRuleRequest) (response *ModifyL7AccRuleResponse, err error) {
    if request == nil {
        request = NewModifyL7AccRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyL7AccRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyL7AccRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyL7AccRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyL7AccRulePriorityRequest() (request *ModifyL7AccRulePriorityRequest) {
    request = &ModifyL7AccRulePriorityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyL7AccRulePriority")
    
    
    return
}

func NewModifyL7AccRulePriorityResponse() (response *ModifyL7AccRulePriorityResponse) {
    response = &ModifyL7AccRulePriorityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyL7AccRulePriority
// This API is used to modify the priority of the rule list in the rule engine (https://www.tencentcloud.com/document/product/1552/70901?from_cn_redirect=1). It requires the input of the complete rule ID list under the site ID. The rule ID list can be obtained through the API for querying layer-7 acceleration rules (https://www.tencentcloud.com/document/product/1552/115820?from_cn_redirect=1). The final priority order will be adjusted to the order of the rule ID list, executed from front to back.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDRULEENGINE = "InvalidParameter.InvalidRuleEngine"
//  INVALIDPARAMETER_INVALIDRULEENGINENOTFOUND = "InvalidParameter.InvalidRuleEngineNotFound"
func (c *Client) ModifyL7AccRulePriority(request *ModifyL7AccRulePriorityRequest) (response *ModifyL7AccRulePriorityResponse, err error) {
    return c.ModifyL7AccRulePriorityWithContext(context.Background(), request)
}

// ModifyL7AccRulePriority
// This API is used to modify the priority of the rule list in the rule engine (https://www.tencentcloud.com/document/product/1552/70901?from_cn_redirect=1). It requires the input of the complete rule ID list under the site ID. The rule ID list can be obtained through the API for querying layer-7 acceleration rules (https://www.tencentcloud.com/document/product/1552/115820?from_cn_redirect=1). The final priority order will be adjusted to the order of the rule ID list, executed from front to back.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDRULEENGINE = "InvalidParameter.InvalidRuleEngine"
//  INVALIDPARAMETER_INVALIDRULEENGINENOTFOUND = "InvalidParameter.InvalidRuleEngineNotFound"
func (c *Client) ModifyL7AccRulePriorityWithContext(ctx context.Context, request *ModifyL7AccRulePriorityRequest) (response *ModifyL7AccRulePriorityResponse, err error) {
    if request == nil {
        request = NewModifyL7AccRulePriorityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyL7AccRulePriority")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyL7AccRulePriority require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyL7AccRulePriorityResponse()
    err = c.Send(request, response)
    return
}

func NewModifyL7AccSettingRequest() (request *ModifyL7AccSettingRequest) {
    request = &ModifyL7AccSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyL7AccSetting")
    
    
    return
}

func NewModifyL7AccSettingResponse() (response *ModifyL7AccSettingResponse) {
    response = &ModifyL7AccSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyL7AccSetting
// This API is used to modify the global configuration of site acceleration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWERROR = "InternalError.UnknowError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_CACHEKEYQUERYSTRINGREQUIRESFULLURLCACHEOFF = "InvalidParameter.CacheKeyQueryStringRequiresFullUrlCacheOff"
//  INVALIDPARAMETER_CACHEKEYQUERYSTRINGTOOMANYVALUE = "InvalidParameter.CacheKeyQueryStringTooManyValue"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_CLIENTIPCOUNTRYCONFLICTSWITHIPV6 = "InvalidParameter.ClientIpCountryConflictsWithIpv6"
//  INVALIDPARAMETER_GRPCREQUIREHTTP2 = "InvalidParameter.GrpcRequireHttp2"
//  INVALIDPARAMETER_INVALIDAWSPRIVATEACCESS = "InvalidParameter.InvalidAwsPrivateAccess"
//  INVALIDPARAMETER_INVALIDCACHECONFIGFOLLOWORIGIN = "InvalidParameter.InvalidCacheConfigFollowOrigin"
//  INVALIDPARAMETER_INVALIDCACHEKEYQUERYSTRINGVALUE = "InvalidParameter.InvalidCacheKeyQueryStringValue"
//  INVALIDPARAMETER_INVALIDCACHEONLYONSWITCH = "InvalidParameter.InvalidCacheOnlyOnSwitch"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPCOUNTRYHEADERNAME = "InvalidParameter.InvalidClientIpCountryHeaderName"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINE = "InvalidParameter.InvalidDynamicRoutine"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINEBILLING = "InvalidParameter.InvalidDynamicRoutineBilling"
//  INVALIDPARAMETER_INVALIDFORCEREDIRECTTYPE = "InvalidParameter.InvalidForceRedirectType"
//  INVALIDPARAMETER_INVALIDHTTPS = "InvalidParameter.InvalidHttps"
//  INVALIDPARAMETER_INVALIDHTTPSCERTINFO = "InvalidParameter.InvalidHttpsCertInfo"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDHTTPSTLSVERSION = "InvalidParameter.InvalidHttpsTlsVersion"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGETIME = "InvalidParameter.InvalidMaxAgeTime"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDORIGINTYPE = "InvalidParameter.InvalidOriginType"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPOSTMAXSIZEBILLING = "InvalidParameter.InvalidPostMaxSizeBilling"
//  INVALIDPARAMETER_INVALIDPOSTSIZEVALUE = "InvalidParameter.InvalidPostSizeValue"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDRESOURCEIDBILLING = "InvalidParameter.InvalidResourceIdBilling"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGCLIENTIP = "InvalidParameter.InvalidStandardDebugClientIp"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGEMPTYLIST = "InvalidParameter.InvalidStandardDebugEmptyList"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGEXPIRETIMELIMIT = "InvalidParameter.InvalidStandardDebugExpireTimeLimit"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_MULTIPLYLAYERNOTSUPPORTSMARTROUTING = "InvalidParameter.MultiplyLayerNotSupportSmartRouting"
//  INVALIDPARAMETER_OCDIRECTORIGINREQUIRESSMARTROUTING = "InvalidParameter.OCDirectOriginRequiresSmartRouting"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_SETTINGINVALIDPARAM = "InvalidParameter.SettingInvalidParam"
//  INVALIDPARAMETER_SHIELDNOTSUPPORTZONECACHEPREFRESH = "InvalidParameter.ShieldNotSupportZoneCachePrefresh"
//  INVALIDPARAMETER_SHIELDNOTSUPPORTZONESMARTROUTING = "InvalidParameter.ShieldNotSupportZoneSmartRouting"
//  INVALIDPARAMETER_ZONEISGRAYPUBLISHING = "InvalidParameter.ZoneIsGrayPublishing"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  INVALIDPARAMETERVALUE_FORMATMISMATCH = "InvalidParameterValue.FormatMismatch"
//  INVALIDPARAMETERVALUE_GENERALMISMATCH = "InvalidParameterValue.GeneralMismatch"
//  INVALIDPARAMETERVALUE_INCLUDEINVALIDVALUE = "InvalidParameterValue.IncludeInvalidValue"
//  INVALIDPARAMETERVALUE_MISSINGNECESSARYPARAM = "InvalidParameterValue.MissingNecessaryParam"
//  INVALIDPARAMETERVALUE_NOTINENUMERATION = "InvalidParameterValue.NotInEnumeration"
//  INVALIDPARAMETERVALUE_NOTWITHINRANGE = "InvalidParameterValue.NotWithinRange"
//  INVALIDPARAMETERVALUE_REGEXMISMATCH = "InvalidParameterValue.RegExMismatch"
//  INVALIDPARAMETERVALUE_UNRECOGNIZABLEVALUE = "InvalidParameterValue.UnrecognizableValue"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDDISABLE = "OperationDenied.AccelerateMainlandDisable"
//  OPERATIONDENIED_ACCELERATEMAINLANDIPV6CONFLICT = "OperationDenied.AccelerateMainlandIpv6Conflict"
//  OPERATIONDENIED_ACCELERATEMAINLANDMULTIPLYLAYERCONFLICT = "OperationDenied.AccelerateMainlandMultiplyLayerConflict"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_NOTALLOWCLOSEACCELERATEMAINLANDWHENDOMAINBOUNDEXCLUSIVEIP = "OperationDenied.NotAllowCloseAccelerateMainlandWhenDomainBoundExclusiveIP"
//  OPERATIONDENIED_SHAREDCNAMEUNSUPPORTEDACCELERATEMAINLAND = "OperationDenied.SharedCNAMEUnsupportedAccelerateMainland"
//  OPERATIONDENIED_SHAREDCNAMEUNSUPPORTEDIPV6 = "OperationDenied.SharedCNAMEUnsupportedIPv6"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_POSTMAXSIZEQUOTANOTFOUND = "ResourceNotFound.PostMaxSizeQuotaNotFound"
//  RESOURCEUNAVAILABLE_CERTNOTFOUND = "ResourceUnavailable.CertNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ModifyL7AccSetting(request *ModifyL7AccSettingRequest) (response *ModifyL7AccSettingResponse, err error) {
    return c.ModifyL7AccSettingWithContext(context.Background(), request)
}

// ModifyL7AccSetting
// This API is used to modify the global configuration of site acceleration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWERROR = "InternalError.UnknowError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_CACHEKEYQUERYSTRINGREQUIRESFULLURLCACHEOFF = "InvalidParameter.CacheKeyQueryStringRequiresFullUrlCacheOff"
//  INVALIDPARAMETER_CACHEKEYQUERYSTRINGTOOMANYVALUE = "InvalidParameter.CacheKeyQueryStringTooManyValue"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_CLIENTIPCOUNTRYCONFLICTSWITHIPV6 = "InvalidParameter.ClientIpCountryConflictsWithIpv6"
//  INVALIDPARAMETER_GRPCREQUIREHTTP2 = "InvalidParameter.GrpcRequireHttp2"
//  INVALIDPARAMETER_INVALIDAWSPRIVATEACCESS = "InvalidParameter.InvalidAwsPrivateAccess"
//  INVALIDPARAMETER_INVALIDCACHECONFIGFOLLOWORIGIN = "InvalidParameter.InvalidCacheConfigFollowOrigin"
//  INVALIDPARAMETER_INVALIDCACHEKEYQUERYSTRINGVALUE = "InvalidParameter.InvalidCacheKeyQueryStringValue"
//  INVALIDPARAMETER_INVALIDCACHEONLYONSWITCH = "InvalidParameter.InvalidCacheOnlyOnSwitch"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPCOUNTRYHEADERNAME = "InvalidParameter.InvalidClientIpCountryHeaderName"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINE = "InvalidParameter.InvalidDynamicRoutine"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINEBILLING = "InvalidParameter.InvalidDynamicRoutineBilling"
//  INVALIDPARAMETER_INVALIDFORCEREDIRECTTYPE = "InvalidParameter.InvalidForceRedirectType"
//  INVALIDPARAMETER_INVALIDHTTPS = "InvalidParameter.InvalidHttps"
//  INVALIDPARAMETER_INVALIDHTTPSCERTINFO = "InvalidParameter.InvalidHttpsCertInfo"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDHTTPSTLSVERSION = "InvalidParameter.InvalidHttpsTlsVersion"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGETIME = "InvalidParameter.InvalidMaxAgeTime"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDORIGINTYPE = "InvalidParameter.InvalidOriginType"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPOSTMAXSIZEBILLING = "InvalidParameter.InvalidPostMaxSizeBilling"
//  INVALIDPARAMETER_INVALIDPOSTSIZEVALUE = "InvalidParameter.InvalidPostSizeValue"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDRESOURCEIDBILLING = "InvalidParameter.InvalidResourceIdBilling"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGCLIENTIP = "InvalidParameter.InvalidStandardDebugClientIp"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGEMPTYLIST = "InvalidParameter.InvalidStandardDebugEmptyList"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGEXPIRETIMELIMIT = "InvalidParameter.InvalidStandardDebugExpireTimeLimit"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_MULTIPLYLAYERNOTSUPPORTSMARTROUTING = "InvalidParameter.MultiplyLayerNotSupportSmartRouting"
//  INVALIDPARAMETER_OCDIRECTORIGINREQUIRESSMARTROUTING = "InvalidParameter.OCDirectOriginRequiresSmartRouting"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_SETTINGINVALIDPARAM = "InvalidParameter.SettingInvalidParam"
//  INVALIDPARAMETER_SHIELDNOTSUPPORTZONECACHEPREFRESH = "InvalidParameter.ShieldNotSupportZoneCachePrefresh"
//  INVALIDPARAMETER_SHIELDNOTSUPPORTZONESMARTROUTING = "InvalidParameter.ShieldNotSupportZoneSmartRouting"
//  INVALIDPARAMETER_ZONEISGRAYPUBLISHING = "InvalidParameter.ZoneIsGrayPublishing"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  INVALIDPARAMETERVALUE_FORMATMISMATCH = "InvalidParameterValue.FormatMismatch"
//  INVALIDPARAMETERVALUE_GENERALMISMATCH = "InvalidParameterValue.GeneralMismatch"
//  INVALIDPARAMETERVALUE_INCLUDEINVALIDVALUE = "InvalidParameterValue.IncludeInvalidValue"
//  INVALIDPARAMETERVALUE_MISSINGNECESSARYPARAM = "InvalidParameterValue.MissingNecessaryParam"
//  INVALIDPARAMETERVALUE_NOTINENUMERATION = "InvalidParameterValue.NotInEnumeration"
//  INVALIDPARAMETERVALUE_NOTWITHINRANGE = "InvalidParameterValue.NotWithinRange"
//  INVALIDPARAMETERVALUE_REGEXMISMATCH = "InvalidParameterValue.RegExMismatch"
//  INVALIDPARAMETERVALUE_UNRECOGNIZABLEVALUE = "InvalidParameterValue.UnrecognizableValue"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDDISABLE = "OperationDenied.AccelerateMainlandDisable"
//  OPERATIONDENIED_ACCELERATEMAINLANDIPV6CONFLICT = "OperationDenied.AccelerateMainlandIpv6Conflict"
//  OPERATIONDENIED_ACCELERATEMAINLANDMULTIPLYLAYERCONFLICT = "OperationDenied.AccelerateMainlandMultiplyLayerConflict"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_NOTALLOWCLOSEACCELERATEMAINLANDWHENDOMAINBOUNDEXCLUSIVEIP = "OperationDenied.NotAllowCloseAccelerateMainlandWhenDomainBoundExclusiveIP"
//  OPERATIONDENIED_SHAREDCNAMEUNSUPPORTEDACCELERATEMAINLAND = "OperationDenied.SharedCNAMEUnsupportedAccelerateMainland"
//  OPERATIONDENIED_SHAREDCNAMEUNSUPPORTEDIPV6 = "OperationDenied.SharedCNAMEUnsupportedIPv6"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_POSTMAXSIZEQUOTANOTFOUND = "ResourceNotFound.PostMaxSizeQuotaNotFound"
//  RESOURCEUNAVAILABLE_CERTNOTFOUND = "ResourceUnavailable.CertNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ModifyL7AccSettingWithContext(ctx context.Context, request *ModifyL7AccSettingRequest) (response *ModifyL7AccSettingResponse, err error) {
    if request == nil {
        request = NewModifyL7AccSettingRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyL7AccSetting")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyL7AccSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyL7AccSettingResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLoadBalancerRequest() (request *ModifyLoadBalancerRequest) {
    request = &ModifyLoadBalancerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyLoadBalancer")
    
    
    return
}

func NewModifyLoadBalancerResponse() (response *ModifyLoadBalancerResponse) {
    response = &ModifyLoadBalancerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLoadBalancer
// This API is used to modify LoadBalancer configuration. The load balancing feature is in beta test. If you need to use it, [contact us](https://www.tencentcloud.com/contact-us).
//
// error code that may be returned:
//  INVALIDPARAMETER_LOADBALANCERBINDL4NOTINSTABLESTATUS = "InvalidParameter.LoadBalancerBindL4NotInStableStatus"
//  INVALIDPARAMETER_LOADBALANCERBINDL7NOTINSTABLESTATUS = "InvalidParameter.LoadBalancerBindL7NotInStableStatus"
func (c *Client) ModifyLoadBalancer(request *ModifyLoadBalancerRequest) (response *ModifyLoadBalancerResponse, err error) {
    return c.ModifyLoadBalancerWithContext(context.Background(), request)
}

// ModifyLoadBalancer
// This API is used to modify LoadBalancer configuration. The load balancing feature is in beta test. If you need to use it, [contact us](https://www.tencentcloud.com/contact-us).
//
// error code that may be returned:
//  INVALIDPARAMETER_LOADBALANCERBINDL4NOTINSTABLESTATUS = "InvalidParameter.LoadBalancerBindL4NotInStableStatus"
//  INVALIDPARAMETER_LOADBALANCERBINDL7NOTINSTABLESTATUS = "InvalidParameter.LoadBalancerBindL7NotInStableStatus"
func (c *Client) ModifyLoadBalancerWithContext(ctx context.Context, request *ModifyLoadBalancerRequest) (response *ModifyLoadBalancerResponse, err error) {
    if request == nil {
        request = NewModifyLoadBalancerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyLoadBalancer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLoadBalancer require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLoadBalancerResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMultiPathGatewayRequest() (request *ModifyMultiPathGatewayRequest) {
    request = &ModifyMultiPathGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyMultiPathGateway")
    
    
    return
}

func NewModifyMultiPathGatewayResponse() (response *ModifyMultiPathGatewayResponse) {
    response = &ModifyMultiPathGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMultiPathGateway
// This API is used to modify multi-channel security acceleration gateway information, such as name, gateway ID, IP, and port.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyMultiPathGateway(request *ModifyMultiPathGatewayRequest) (response *ModifyMultiPathGatewayResponse, err error) {
    return c.ModifyMultiPathGatewayWithContext(context.Background(), request)
}

// ModifyMultiPathGateway
// This API is used to modify multi-channel security acceleration gateway information, such as name, gateway ID, IP, and port.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyMultiPathGatewayWithContext(ctx context.Context, request *ModifyMultiPathGatewayRequest) (response *ModifyMultiPathGatewayResponse, err error) {
    if request == nil {
        request = NewModifyMultiPathGatewayRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyMultiPathGateway")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMultiPathGateway require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMultiPathGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMultiPathGatewayLineRequest() (request *ModifyMultiPathGatewayLineRequest) {
    request = &ModifyMultiPathGatewayLineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyMultiPathGatewayLine")
    
    
    return
}

func NewModifyMultiPathGatewayLineResponse() (response *ModifyMultiPathGatewayLineResponse) {
    response = &ModifyMultiPathGatewayLineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMultiPathGatewayLine
// This API is used to modify the configuration for connecting lines to the multi-channel security acceleration gateway, including EdgeOne Layer 4 proxy lines and custom lines.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyMultiPathGatewayLine(request *ModifyMultiPathGatewayLineRequest) (response *ModifyMultiPathGatewayLineResponse, err error) {
    return c.ModifyMultiPathGatewayLineWithContext(context.Background(), request)
}

// ModifyMultiPathGatewayLine
// This API is used to modify the configuration for connecting lines to the multi-channel security acceleration gateway, including EdgeOne Layer 4 proxy lines and custom lines.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyMultiPathGatewayLineWithContext(ctx context.Context, request *ModifyMultiPathGatewayLineRequest) (response *ModifyMultiPathGatewayLineResponse, err error) {
    if request == nil {
        request = NewModifyMultiPathGatewayLineRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyMultiPathGatewayLine")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMultiPathGatewayLine require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMultiPathGatewayLineResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMultiPathGatewaySecretKeyRequest() (request *ModifyMultiPathGatewaySecretKeyRequest) {
    request = &ModifyMultiPathGatewaySecretKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyMultiPathGatewaySecretKey")
    
    
    return
}

func NewModifyMultiPathGatewaySecretKeyResponse() (response *ModifyMultiPathGatewaySecretKeyResponse) {
    response = &ModifyMultiPathGatewaySecretKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMultiPathGatewaySecretKey
// This API is used to modify the key for integrating a multi-channel security acceleration gateway. Customers sign with the integration key to access the multi-channel security acceleration gateway. After modification, the original key becomes invalid.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyMultiPathGatewaySecretKey(request *ModifyMultiPathGatewaySecretKeyRequest) (response *ModifyMultiPathGatewaySecretKeyResponse, err error) {
    return c.ModifyMultiPathGatewaySecretKeyWithContext(context.Background(), request)
}

// ModifyMultiPathGatewaySecretKey
// This API is used to modify the key for integrating a multi-channel security acceleration gateway. Customers sign with the integration key to access the multi-channel security acceleration gateway. After modification, the original key becomes invalid.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyMultiPathGatewaySecretKeyWithContext(ctx context.Context, request *ModifyMultiPathGatewaySecretKeyRequest) (response *ModifyMultiPathGatewaySecretKeyResponse, err error) {
    if request == nil {
        request = NewModifyMultiPathGatewaySecretKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyMultiPathGatewaySecretKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMultiPathGatewaySecretKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMultiPathGatewaySecretKeyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMultiPathGatewayStatusRequest() (request *ModifyMultiPathGatewayStatusRequest) {
    request = &ModifyMultiPathGatewayStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyMultiPathGatewayStatus")
    
    
    return
}

func NewModifyMultiPathGatewayStatusResponse() (response *ModifyMultiPathGatewayStatusResponse) {
    response = &ModifyMultiPathGatewayStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMultiPathGatewayStatus
// This API is used to update the multi-channel security gateway status.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyMultiPathGatewayStatus(request *ModifyMultiPathGatewayStatusRequest) (response *ModifyMultiPathGatewayStatusResponse, err error) {
    return c.ModifyMultiPathGatewayStatusWithContext(context.Background(), request)
}

// ModifyMultiPathGatewayStatus
// This API is used to update the multi-channel security gateway status.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyMultiPathGatewayStatusWithContext(ctx context.Context, request *ModifyMultiPathGatewayStatusRequest) (response *ModifyMultiPathGatewayStatusResponse, err error) {
    if request == nil {
        request = NewModifyMultiPathGatewayStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyMultiPathGatewayStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMultiPathGatewayStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMultiPathGatewayStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyOriginACLRequest() (request *ModifyOriginACLRequest) {
    request = &ModifyOriginACLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyOriginACL")
    
    
    return
}

func NewModifyOriginACLResponse() (response *ModifyOriginACLResponse) {
    response = &ModifyOriginACLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyOriginACL
// This API is used to modify the origin protection configuration of a site. The following operations are supported: <li>**Resource addition and deletion**: enable or disable origin-pull from specific origin IP ranges for layer-7 accelerated domains/layer-4 proxy instances. This applies to resource adjustment in **specified protection mode**, and unspecified resources will not change;</li><li>**Mode switching**: adjust the origin protection mode of layer-7 accelerated domains/layer-4 proxy instances. After switching to all, the list of accelerated domains or layer-4 proxy instances under the site will be automatically added to the origin protection list in real time, and subsequent resource changes will also be automatically adjusted without manual intervention; after switching to specific, only the mode is changed, and resources already included in protection remain unchanged;</li><li>**IP range version switching**: for whitelisted customers, it supports switching to other available origin protection IP range versions, such as the simplified version, to reduce origin-pull IP ranges.</li>
//
// 
//
// Instructions for specified protection mode:
//
// <li>The maximum number of layer-7 accelerated domains that can be submitted at a time is 200, and the maximum number of layer-4 proxy instances is 100. Mixed submission of layer-7 accelerated domains/layer-4 proxy instances is supported, with a maximum total of 200 instances;</li>
//
// <li>When more than 200 resources need to be changed, submit them in batches through this API. For example, to add 500 layer-7 accelerated domains, first submit 200 domains through this API, and then submit the remaining 300 domains in two batches, with no more than 200 domains each time.</li>
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDOMAINS = "InvalidParameter.InvalidDomains"
//  INVALIDPARAMETER_INVALIDPROXIES = "InvalidParameter.InvalidProxies"
//  INVALIDPARAMETER_SHIELDNOTSUPPORTHOSTORIGINWHITELIST = "InvalidParameter.ShieldNotSupportHostOriginWhitelist"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_UNSUPPORTEDPLAN = "OperationDenied.UnsupportedPlan"
//  OPERATIONDENIED_UPDATEIPWHITELISTFIRST = "OperationDenied.UpdateIPWhitelistFirst"
func (c *Client) ModifyOriginACL(request *ModifyOriginACLRequest) (response *ModifyOriginACLResponse, err error) {
    return c.ModifyOriginACLWithContext(context.Background(), request)
}

// ModifyOriginACL
// This API is used to modify the origin protection configuration of a site. The following operations are supported: <li>**Resource addition and deletion**: enable or disable origin-pull from specific origin IP ranges for layer-7 accelerated domains/layer-4 proxy instances. This applies to resource adjustment in **specified protection mode**, and unspecified resources will not change;</li><li>**Mode switching**: adjust the origin protection mode of layer-7 accelerated domains/layer-4 proxy instances. After switching to all, the list of accelerated domains or layer-4 proxy instances under the site will be automatically added to the origin protection list in real time, and subsequent resource changes will also be automatically adjusted without manual intervention; after switching to specific, only the mode is changed, and resources already included in protection remain unchanged;</li><li>**IP range version switching**: for whitelisted customers, it supports switching to other available origin protection IP range versions, such as the simplified version, to reduce origin-pull IP ranges.</li>
//
// 
//
// Instructions for specified protection mode:
//
// <li>The maximum number of layer-7 accelerated domains that can be submitted at a time is 200, and the maximum number of layer-4 proxy instances is 100. Mixed submission of layer-7 accelerated domains/layer-4 proxy instances is supported, with a maximum total of 200 instances;</li>
//
// <li>When more than 200 resources need to be changed, submit them in batches through this API. For example, to add 500 layer-7 accelerated domains, first submit 200 domains through this API, and then submit the remaining 300 domains in two batches, with no more than 200 domains each time.</li>
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDOMAINS = "InvalidParameter.InvalidDomains"
//  INVALIDPARAMETER_INVALIDPROXIES = "InvalidParameter.InvalidProxies"
//  INVALIDPARAMETER_SHIELDNOTSUPPORTHOSTORIGINWHITELIST = "InvalidParameter.ShieldNotSupportHostOriginWhitelist"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_UNSUPPORTEDPLAN = "OperationDenied.UnsupportedPlan"
//  OPERATIONDENIED_UPDATEIPWHITELISTFIRST = "OperationDenied.UpdateIPWhitelistFirst"
func (c *Client) ModifyOriginACLWithContext(ctx context.Context, request *ModifyOriginACLRequest) (response *ModifyOriginACLResponse, err error) {
    if request == nil {
        request = NewModifyOriginACLRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyOriginACL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyOriginACL require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyOriginACLResponse()
    err = c.Send(request, response)
    return
}

func NewModifyOriginGroupRequest() (request *ModifyOriginGroupRequest) {
    request = &ModifyOriginGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyOriginGroup")
    
    
    return
}

func NewModifyOriginGroupResponse() (response *ModifyOriginGroupResponse) {
    response = &ModifyOriginGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyOriginGroup
// Modify origin group configuration. Newly submitted origin records will overwrite the original origin records in the origin group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_HOSTHEADERINVALID = "InvalidParameter.HostHeaderInvalid"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_ORIGINL4RECORDIPV4MIXDOMAIN = "InvalidParameter.OriginL4RecordIPV4MixDomain"
//  INVALIDPARAMETER_ORIGINL4RECORDMULTIDOMAIN = "InvalidParameter.OriginL4RecordMultiDomain"
//  INVALIDPARAMETER_ORIGINNAMEEXISTS = "InvalidParameter.OriginNameExists"
//  INVALIDPARAMETER_ORIGINRECORDFORMATERROR = "InvalidParameter.OriginRecordFormatError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATIONDOMAINSTATUSNOTINONLINE = "OperationDenied.AccelerationDomainStatusNotInOnline"
//  OPERATIONDENIED_L4STATUSNOTINONLINE = "OperationDenied.L4StatusNotInOnline"
//  OPERATIONDENIED_LOADBALANCESTATUSNOTINONLINE = "OperationDenied.LoadBalanceStatusNotInOnline"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ModifyOriginGroup(request *ModifyOriginGroupRequest) (response *ModifyOriginGroupResponse, err error) {
    return c.ModifyOriginGroupWithContext(context.Background(), request)
}

// ModifyOriginGroup
// Modify origin group configuration. Newly submitted origin records will overwrite the original origin records in the origin group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_HOSTHEADERINVALID = "InvalidParameter.HostHeaderInvalid"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_ORIGINL4RECORDIPV4MIXDOMAIN = "InvalidParameter.OriginL4RecordIPV4MixDomain"
//  INVALIDPARAMETER_ORIGINL4RECORDMULTIDOMAIN = "InvalidParameter.OriginL4RecordMultiDomain"
//  INVALIDPARAMETER_ORIGINNAMEEXISTS = "InvalidParameter.OriginNameExists"
//  INVALIDPARAMETER_ORIGINRECORDFORMATERROR = "InvalidParameter.OriginRecordFormatError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATIONDOMAINSTATUSNOTINONLINE = "OperationDenied.AccelerationDomainStatusNotInOnline"
//  OPERATIONDENIED_L4STATUSNOTINONLINE = "OperationDenied.L4StatusNotInOnline"
//  OPERATIONDENIED_LOADBALANCESTATUSNOTINONLINE = "OperationDenied.LoadBalanceStatusNotInOnline"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ModifyOriginGroupWithContext(ctx context.Context, request *ModifyOriginGroupRequest) (response *ModifyOriginGroupResponse, err error) {
    if request == nil {
        request = NewModifyOriginGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyOriginGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyOriginGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyOriginGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPlanRequest() (request *ModifyPlanRequest) {
    request = &ModifyPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyPlan")
    
    
    return
}

func NewModifyPlanResponse() (response *ModifyPlanResponse) {
    response = &ModifyPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPlan
// Modify the plan settings. Currently, only the auto-renewal switch of prepaid plans can be modified.
//
// error code that may be returned:
//  INVALIDPARAMETER_PLANNOTFOUND = "InvalidParameter.PlanNotFound"
//  OPERATIONDENIED_ENTERPRISEPLANAUTORENEWUNSUPPORTED = "OperationDenied.EnterprisePlanAutoRenewUnsupported"
//  OPERATIONDENIED_PLANHASBEENISOLATED = "OperationDenied.PlanHasBeenIsolated"
func (c *Client) ModifyPlan(request *ModifyPlanRequest) (response *ModifyPlanResponse, err error) {
    return c.ModifyPlanWithContext(context.Background(), request)
}

// ModifyPlan
// Modify the plan settings. Currently, only the auto-renewal switch of prepaid plans can be modified.
//
// error code that may be returned:
//  INVALIDPARAMETER_PLANNOTFOUND = "InvalidParameter.PlanNotFound"
//  OPERATIONDENIED_ENTERPRISEPLANAUTORENEWUNSUPPORTED = "OperationDenied.EnterprisePlanAutoRenewUnsupported"
//  OPERATIONDENIED_PLANHASBEENISOLATED = "OperationDenied.PlanHasBeenIsolated"
func (c *Client) ModifyPlanWithContext(ctx context.Context, request *ModifyPlanRequest) (response *ModifyPlanResponse, err error) {
    if request == nil {
        request = NewModifyPlanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyPlan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPlanResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrefetchOriginLimitRequest() (request *ModifyPrefetchOriginLimitRequest) {
    request = &ModifyPrefetchOriginLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyPrefetchOriginLimit")
    
    
    return
}

func NewModifyPrefetchOriginLimitResponse() (response *ModifyPrefetchOriginLimitResponse) {
    response = &ModifyPrefetchOriginLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPrefetchOriginLimit
// This API is used to configure origin-pull rate limiting. This feature is in allowlist beta test.
//
// This API is used to create, modify, and delete preheated origin-pull rate limits. Each account supports up to 100 limits.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATEPREFETCHORIGINLIMITFAILED = "FailedOperation.CreatePrefetchOriginLimitFailed"
//  FAILEDOPERATION_PREFETCHORIGINLIMITCOUNTEXCEEDED = "FailedOperation.PrefetchOriginLimitCountExceeded"
//  FAILEDOPERATION_PREFETCHORIGINLIMITNOTFOUND = "FailedOperation.PrefetchOriginLimitNotFound"
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DOMAINCONFIG = "InternalError.DomainConfig"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINAREANOTSUPPORTPREFETCHORIGINLIMITAREA = "InvalidParameter.DomainAreaNotSupportPrefetchOriginLimitArea"
//  INVALIDPARAMETER_DOMAINNOTFOUND = "InvalidParameter.DomainNotFound"
//  INVALIDPARAMETER_PREFETCHORIGINLIMITAREAINVALID = "InvalidParameter.PrefetchOriginLimitAreaInvalid"
//  INVALIDPARAMETER_PREFETCHORIGINLIMITBANDWIDTHTOOLARGE = "InvalidParameter.PrefetchOriginLimitBandwidthTooLarge"
//  INVALIDPARAMETER_PREFETCHORIGINLIMITBANDWIDTHTOOSMALL = "InvalidParameter.PrefetchOriginLimitBandwidthTooSmall"
//  INVALIDPARAMETER_PREFETCHORIGINLIMITENABLEDINVALID = "InvalidParameter.PrefetchOriginLimitEnabledInvalid"
//  INVALIDPARAMETER_ZONEAREANOTSUPPORTPREFETCHORIGINLIMITAREA = "InvalidParameter.ZoneAreaNotSupportPrefetchOriginLimitArea"
//  OPERATIONDENIED_NOTINPREFETCHORIGINLIMITWHITELIST = "OperationDenied.NotInPrefetchOriginLimitWhiteList"
//  RESOURCEINUSE_PREFETCHORIGINLIMITALREADYEXISTS = "ResourceInUse.PrefetchOriginLimitAlreadyExists"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ModifyPrefetchOriginLimit(request *ModifyPrefetchOriginLimitRequest) (response *ModifyPrefetchOriginLimitResponse, err error) {
    return c.ModifyPrefetchOriginLimitWithContext(context.Background(), request)
}

// ModifyPrefetchOriginLimit
// This API is used to configure origin-pull rate limiting. This feature is in allowlist beta test.
//
// This API is used to create, modify, and delete preheated origin-pull rate limits. Each account supports up to 100 limits.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATEPREFETCHORIGINLIMITFAILED = "FailedOperation.CreatePrefetchOriginLimitFailed"
//  FAILEDOPERATION_PREFETCHORIGINLIMITCOUNTEXCEEDED = "FailedOperation.PrefetchOriginLimitCountExceeded"
//  FAILEDOPERATION_PREFETCHORIGINLIMITNOTFOUND = "FailedOperation.PrefetchOriginLimitNotFound"
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DOMAINCONFIG = "InternalError.DomainConfig"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINAREANOTSUPPORTPREFETCHORIGINLIMITAREA = "InvalidParameter.DomainAreaNotSupportPrefetchOriginLimitArea"
//  INVALIDPARAMETER_DOMAINNOTFOUND = "InvalidParameter.DomainNotFound"
//  INVALIDPARAMETER_PREFETCHORIGINLIMITAREAINVALID = "InvalidParameter.PrefetchOriginLimitAreaInvalid"
//  INVALIDPARAMETER_PREFETCHORIGINLIMITBANDWIDTHTOOLARGE = "InvalidParameter.PrefetchOriginLimitBandwidthTooLarge"
//  INVALIDPARAMETER_PREFETCHORIGINLIMITBANDWIDTHTOOSMALL = "InvalidParameter.PrefetchOriginLimitBandwidthTooSmall"
//  INVALIDPARAMETER_PREFETCHORIGINLIMITENABLEDINVALID = "InvalidParameter.PrefetchOriginLimitEnabledInvalid"
//  INVALIDPARAMETER_ZONEAREANOTSUPPORTPREFETCHORIGINLIMITAREA = "InvalidParameter.ZoneAreaNotSupportPrefetchOriginLimitArea"
//  OPERATIONDENIED_NOTINPREFETCHORIGINLIMITWHITELIST = "OperationDenied.NotInPrefetchOriginLimitWhiteList"
//  RESOURCEINUSE_PREFETCHORIGINLIMITALREADYEXISTS = "ResourceInUse.PrefetchOriginLimitAlreadyExists"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ModifyPrefetchOriginLimitWithContext(ctx context.Context, request *ModifyPrefetchOriginLimitRequest) (response *ModifyPrefetchOriginLimitResponse, err error) {
    if request == nil {
        request = NewModifyPrefetchOriginLimitRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyPrefetchOriginLimit")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPrefetchOriginLimit require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPrefetchOriginLimitResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRealtimeLogDeliveryTaskRequest() (request *ModifyRealtimeLogDeliveryTaskRequest) {
    request = &ModifyRealtimeLogDeliveryTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyRealtimeLogDeliveryTask")
    
    
    return
}

func NewModifyRealtimeLogDeliveryTaskResponse() (response *ModifyRealtimeLogDeliveryTaskResponse) {
    response = &ModifyRealtimeLogDeliveryTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRealtimeLogDeliveryTask
// This API is used to modify the configuration of a real-time log delivery task. It has the following limits:
//
// <ul>
//
// <li>The destination type (`TaskType`) of a real-time log delivery task cannot be modified;</li>
//
// <li>The data delivery type (`LogType`) cannot be modified;</li>
//
// <li>The data delivery region (`Area`) cannot be modified;</li>
//
// <li>When the destination of the original real-time log delivery task is Tencent Cloud CLS, the destination details such as the log set and log topic cannot be modified.</li>
//
// </ul>
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATELOGTOPICTASKAUTHFAILURE = "FailedOperation.CreateLogTopicTaskAuthFailure"
//  INVALIDPARAMETER_INVALIDLOGFORMATFIELDDELIMITER = "InvalidParameter.InvalidLogFormatFieldDelimiter"
//  INVALIDPARAMETER_INVALIDLOGFORMATFORMATTYPE = "InvalidParameter.InvalidLogFormatFormatType"
//  INVALIDPARAMETER_INVALIDLOGFORMATRECORDDELIMITER = "InvalidParameter.InvalidLogFormatRecordDelimiter"
//  LIMITEXCEEDED_CUSTOMLOGFIELDREGEXLIMITEXCEEDED = "LimitExceeded.CustomLogFieldRegexLimitExceeded"
func (c *Client) ModifyRealtimeLogDeliveryTask(request *ModifyRealtimeLogDeliveryTaskRequest) (response *ModifyRealtimeLogDeliveryTaskResponse, err error) {
    return c.ModifyRealtimeLogDeliveryTaskWithContext(context.Background(), request)
}

// ModifyRealtimeLogDeliveryTask
// This API is used to modify the configuration of a real-time log delivery task. It has the following limits:
//
// <ul>
//
// <li>The destination type (`TaskType`) of a real-time log delivery task cannot be modified;</li>
//
// <li>The data delivery type (`LogType`) cannot be modified;</li>
//
// <li>The data delivery region (`Area`) cannot be modified;</li>
//
// <li>When the destination of the original real-time log delivery task is Tencent Cloud CLS, the destination details such as the log set and log topic cannot be modified.</li>
//
// </ul>
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATELOGTOPICTASKAUTHFAILURE = "FailedOperation.CreateLogTopicTaskAuthFailure"
//  INVALIDPARAMETER_INVALIDLOGFORMATFIELDDELIMITER = "InvalidParameter.InvalidLogFormatFieldDelimiter"
//  INVALIDPARAMETER_INVALIDLOGFORMATFORMATTYPE = "InvalidParameter.InvalidLogFormatFormatType"
//  INVALIDPARAMETER_INVALIDLOGFORMATRECORDDELIMITER = "InvalidParameter.InvalidLogFormatRecordDelimiter"
//  LIMITEXCEEDED_CUSTOMLOGFIELDREGEXLIMITEXCEEDED = "LimitExceeded.CustomLogFieldRegexLimitExceeded"
func (c *Client) ModifyRealtimeLogDeliveryTaskWithContext(ctx context.Context, request *ModifyRealtimeLogDeliveryTaskRequest) (response *ModifyRealtimeLogDeliveryTaskResponse, err error) {
    if request == nil {
        request = NewModifyRealtimeLogDeliveryTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyRealtimeLogDeliveryTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRealtimeLogDeliveryTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRealtimeLogDeliveryTaskResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRuleRequest() (request *ModifyRuleRequest) {
    request = &ModifyRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyRule")
    
    
    return
}

func NewModifyRuleResponse() (response *ModifyRuleResponse) {
    response = &ModifyRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRule
// This API is on an earlier version for modifying rule engine APIs. EdgeOne comprehensively upgraded rule engine-related APIs on January 21, 2025. For details on the new version API for modifying layer-7 acceleration rules, refer to [ModifyL7AccRule](https://www.tencentcloud.com/document/product/1552/115818?from_cn_redirect=1).
//
// <p style="color: red;">Note: Starting from January 21, 2025, the legacy version API will stop update and iteration. New features will only be provided in the latest version API. The existing capacity supported by the legacy version API will not be affected. To avoid data field conflict when using the legacy version API, migrate to the latest version rule engine API as soon as possible.</p>
//
// error code that may be returned:
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_CACHEKEYQUERYSTRINGTOOMANYVALUE = "InvalidParameter.CacheKeyQueryStringTooManyValue"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_COMPRESSIONINVALIDALGORITHMS = "InvalidParameter.CompressionInvalidAlgorithms"
//  INVALIDPARAMETER_ERRACTIONUNSUPPORTTARGET = "InvalidParameter.ErrActionUnsupportTarget"
//  INVALIDPARAMETER_ERRINVALIDACTION = "InvalidParameter.ErrInvalidAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONDUPLICATEACTION = "InvalidParameter.ErrInvalidActionDuplicateAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMACTION = "InvalidParameter.ErrInvalidActionParamAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMDUPLICATENAME = "InvalidParameter.ErrInvalidActionParamDuplicateName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMNAME = "InvalidParameter.ErrInvalidActionParamName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMTOOMANYVALUES = "InvalidParameter.ErrInvalidActionParamTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_ERRINVALIDACTIONTYPE = "InvalidParameter.ErrInvalidActionType"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONHOSTTOOMANYWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidConditionHostTooManyWhenModifyOriginActionConfigured"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONIGNORECASE = "InvalidParameter.ErrInvalidConditionIgnoreCase"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMEBADNAME = "InvalidParameter.ErrInvalidConditionNameBadName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMETARGETNOTSUPPORTNAME = "InvalidParameter.ErrInvalidConditionNameTargetNotSupportName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADREGULAR = "InvalidParameter.ErrInvalidConditionValueBadRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADURL = "InvalidParameter.ErrInvalidConditionValueBadUrl"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUECONTAINFILENAMEEXTENSION = "InvalidParameter.ErrInvalidConditionValueBadValueContainFileNameExtension"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOLONGVALUE = "InvalidParameter.ErrInvalidConditionValueTooLongValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYVALUES = "InvalidParameter.ErrInvalidConditionValueTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYWILDCARD = "InvalidParameter.ErrInvalidConditionValueTooManyWildcard"
//  INVALIDPARAMETER_ERRINVALIDELSEWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidElseWhenModifyOriginActionConfigured"
//  INVALIDPARAMETER_ERRNILCONDITION = "InvalidParameter.ErrNilCondition"
//  INVALIDPARAMETER_GRPCREQUIREHTTP2 = "InvalidParameter.GrpcRequireHttp2"
//  INVALIDPARAMETER_HOSTNOTFOUND = "InvalidParameter.HostNotFound"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESECRETKEY = "InvalidParameter.InvalidAuthenticationTypeSecretKey"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESIGNPARAM = "InvalidParameter.InvalidAuthenticationTypeSignParam"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPETIMEFORMAT = "InvalidParameter.InvalidAuthenticationTypeTimeFormat"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPETIMEPARAM = "InvalidParameter.InvalidAuthenticationTypeTimeParam"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDBACKUPSERVERNAME = "InvalidParameter.InvalidBackupServerName"
//  INVALIDPARAMETER_INVALIDCACHEKEY = "InvalidParameter.InvalidCacheKey"
//  INVALIDPARAMETER_INVALIDCACHEKEYCOOKIE = "InvalidParameter.InvalidCacheKeyCookie"
//  INVALIDPARAMETER_INVALIDCACHEKEYIGNORECASE = "InvalidParameter.InvalidCacheKeyIgnoreCase"
//  INVALIDPARAMETER_INVALIDCACHEKEYSCHEME = "InvalidParameter.InvalidCacheKeyScheme"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDCOSDOMAIN = "InvalidParameter.InvalidCosDomain"
//  INVALIDPARAMETER_INVALIDERRORPAGEREDIRECTURL = "InvalidParameter.InvalidErrorPageRedirectUrl"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGETIME = "InvalidParameter.InvalidMaxAgeTime"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAMEXFF = "InvalidParameter.InvalidRequestHeaderNameXff"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERVALUE = "InvalidParameter.InvalidRequestHeaderValue"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERNAME = "InvalidParameter.InvalidResponseHeaderName"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERVALUE = "InvalidParameter.InvalidResponseHeaderValue"
//  INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"
//  INVALIDPARAMETER_INVALIDRULEENGINENOTFOUND = "InvalidParameter.InvalidRuleEngineNotFound"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSEXTENSION = "InvalidParameter.InvalidRuleEngineTargetsExtension"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSURL = "InvalidParameter.InvalidRuleEngineTargetsUrl"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_INVALIDUPSTREAMREQUESTQUERYSTRINGVALUE = "InvalidParameter.InvalidUpstreamRequestQueryStringValue"
//  INVALIDPARAMETER_INVALIDURLREDIRECTHOST = "InvalidParameter.InvalidUrlRedirectHost"
//  INVALIDPARAMETER_INVALIDURLREDIRECTURL = "InvalidParameter.InvalidUrlRedirectUrl"
//  INVALIDPARAMETER_KEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.KeyRulesInvalidQueryStringValue"
//  INVALIDPARAMETER_LOADBALANCEINSTANCEIDISREQUIRED = "InvalidParameter.LoadBalanceInstanceIdIsRequired"
//  INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"
//  INVALIDPARAMETER_ORIGINORIGINGROUPIDISREQUIRED = "InvalidParameter.OriginOriginGroupIdIsRequired"
//  INVALIDPARAMETER_ORIGINPULLPROTOCOLISREQUIRED = "InvalidParameter.OriginPullProtocolIsRequired"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_RESPONSEHEADERCACHECONTROLNOTALLOWDELETE = "InvalidParameter.ResponseHeaderCacheControlNotAllowDelete"
//  INVALIDPARAMETER_STATUSCODECACHEINVALIDSTATUSCODE = "InvalidParameter.StatusCodeCacheInvalidStatusCode"
//  INVALIDPARAMETER_TLSVERSIONNOTINSEQUENCE = "InvalidParameter.TlsVersionNotInSequence"
//  INVALIDPARAMETERVALUE_FORMATMISMATCH = "InvalidParameterValue.FormatMismatch"
//  INVALIDPARAMETERVALUE_GENERALMISMATCH = "InvalidParameterValue.GeneralMismatch"
//  INVALIDPARAMETERVALUE_INCLUDEINVALIDVALUE = "InvalidParameterValue.IncludeInvalidValue"
//  INVALIDPARAMETERVALUE_MISSINGNECESSARYPARAM = "InvalidParameterValue.MissingNecessaryParam"
//  INVALIDPARAMETERVALUE_NOTINENUMERATION = "InvalidParameterValue.NotInEnumeration"
//  INVALIDPARAMETERVALUE_NOTWITHINRANGE = "InvalidParameterValue.NotWithinRange"
//  INVALIDPARAMETERVALUE_REGEXMISMATCH = "InvalidParameterValue.RegExMismatch"
//  INVALIDPARAMETERVALUE_TRIALPLANRESPONSEPAGE = "InvalidParameterValue.TrialPlanResponsePage"
//  INVALIDPARAMETERVALUE_UNRECOGNIZABLEVALUE = "InvalidParameterValue.UnrecognizableValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_INVALIDADVANCEDDEFENSESECURITYTYPE = "OperationDenied.InvalidAdvancedDefenseSecurityType"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ModifyRule(request *ModifyRuleRequest) (response *ModifyRuleResponse, err error) {
    return c.ModifyRuleWithContext(context.Background(), request)
}

// ModifyRule
// This API is on an earlier version for modifying rule engine APIs. EdgeOne comprehensively upgraded rule engine-related APIs on January 21, 2025. For details on the new version API for modifying layer-7 acceleration rules, refer to [ModifyL7AccRule](https://www.tencentcloud.com/document/product/1552/115818?from_cn_redirect=1).
//
// <p style="color: red;">Note: Starting from January 21, 2025, the legacy version API will stop update and iteration. New features will only be provided in the latest version API. The existing capacity supported by the legacy version API will not be affected. To avoid data field conflict when using the legacy version API, migrate to the latest version rule engine API as soon as possible.</p>
//
// error code that may be returned:
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_CACHEKEYQUERYSTRINGTOOMANYVALUE = "InvalidParameter.CacheKeyQueryStringTooManyValue"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_COMPRESSIONINVALIDALGORITHMS = "InvalidParameter.CompressionInvalidAlgorithms"
//  INVALIDPARAMETER_ERRACTIONUNSUPPORTTARGET = "InvalidParameter.ErrActionUnsupportTarget"
//  INVALIDPARAMETER_ERRINVALIDACTION = "InvalidParameter.ErrInvalidAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONDUPLICATEACTION = "InvalidParameter.ErrInvalidActionDuplicateAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMACTION = "InvalidParameter.ErrInvalidActionParamAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMDUPLICATENAME = "InvalidParameter.ErrInvalidActionParamDuplicateName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMNAME = "InvalidParameter.ErrInvalidActionParamName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMTOOMANYVALUES = "InvalidParameter.ErrInvalidActionParamTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_ERRINVALIDACTIONTYPE = "InvalidParameter.ErrInvalidActionType"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONHOSTTOOMANYWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidConditionHostTooManyWhenModifyOriginActionConfigured"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONIGNORECASE = "InvalidParameter.ErrInvalidConditionIgnoreCase"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMEBADNAME = "InvalidParameter.ErrInvalidConditionNameBadName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMETARGETNOTSUPPORTNAME = "InvalidParameter.ErrInvalidConditionNameTargetNotSupportName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADREGULAR = "InvalidParameter.ErrInvalidConditionValueBadRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADURL = "InvalidParameter.ErrInvalidConditionValueBadUrl"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUECONTAINFILENAMEEXTENSION = "InvalidParameter.ErrInvalidConditionValueBadValueContainFileNameExtension"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOLONGVALUE = "InvalidParameter.ErrInvalidConditionValueTooLongValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYVALUES = "InvalidParameter.ErrInvalidConditionValueTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYWILDCARD = "InvalidParameter.ErrInvalidConditionValueTooManyWildcard"
//  INVALIDPARAMETER_ERRINVALIDELSEWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidElseWhenModifyOriginActionConfigured"
//  INVALIDPARAMETER_ERRNILCONDITION = "InvalidParameter.ErrNilCondition"
//  INVALIDPARAMETER_GRPCREQUIREHTTP2 = "InvalidParameter.GrpcRequireHttp2"
//  INVALIDPARAMETER_HOSTNOTFOUND = "InvalidParameter.HostNotFound"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESECRETKEY = "InvalidParameter.InvalidAuthenticationTypeSecretKey"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESIGNPARAM = "InvalidParameter.InvalidAuthenticationTypeSignParam"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPETIMEFORMAT = "InvalidParameter.InvalidAuthenticationTypeTimeFormat"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPETIMEPARAM = "InvalidParameter.InvalidAuthenticationTypeTimeParam"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDBACKUPSERVERNAME = "InvalidParameter.InvalidBackupServerName"
//  INVALIDPARAMETER_INVALIDCACHEKEY = "InvalidParameter.InvalidCacheKey"
//  INVALIDPARAMETER_INVALIDCACHEKEYCOOKIE = "InvalidParameter.InvalidCacheKeyCookie"
//  INVALIDPARAMETER_INVALIDCACHEKEYIGNORECASE = "InvalidParameter.InvalidCacheKeyIgnoreCase"
//  INVALIDPARAMETER_INVALIDCACHEKEYSCHEME = "InvalidParameter.InvalidCacheKeyScheme"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDCOSDOMAIN = "InvalidParameter.InvalidCosDomain"
//  INVALIDPARAMETER_INVALIDERRORPAGEREDIRECTURL = "InvalidParameter.InvalidErrorPageRedirectUrl"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGETIME = "InvalidParameter.InvalidMaxAgeTime"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAMEXFF = "InvalidParameter.InvalidRequestHeaderNameXff"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERVALUE = "InvalidParameter.InvalidRequestHeaderValue"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERNAME = "InvalidParameter.InvalidResponseHeaderName"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERVALUE = "InvalidParameter.InvalidResponseHeaderValue"
//  INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"
//  INVALIDPARAMETER_INVALIDRULEENGINENOTFOUND = "InvalidParameter.InvalidRuleEngineNotFound"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSEXTENSION = "InvalidParameter.InvalidRuleEngineTargetsExtension"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSURL = "InvalidParameter.InvalidRuleEngineTargetsUrl"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_INVALIDUPSTREAMREQUESTQUERYSTRINGVALUE = "InvalidParameter.InvalidUpstreamRequestQueryStringValue"
//  INVALIDPARAMETER_INVALIDURLREDIRECTHOST = "InvalidParameter.InvalidUrlRedirectHost"
//  INVALIDPARAMETER_INVALIDURLREDIRECTURL = "InvalidParameter.InvalidUrlRedirectUrl"
//  INVALIDPARAMETER_KEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.KeyRulesInvalidQueryStringValue"
//  INVALIDPARAMETER_LOADBALANCEINSTANCEIDISREQUIRED = "InvalidParameter.LoadBalanceInstanceIdIsRequired"
//  INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"
//  INVALIDPARAMETER_ORIGINORIGINGROUPIDISREQUIRED = "InvalidParameter.OriginOriginGroupIdIsRequired"
//  INVALIDPARAMETER_ORIGINPULLPROTOCOLISREQUIRED = "InvalidParameter.OriginPullProtocolIsRequired"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_RESPONSEHEADERCACHECONTROLNOTALLOWDELETE = "InvalidParameter.ResponseHeaderCacheControlNotAllowDelete"
//  INVALIDPARAMETER_STATUSCODECACHEINVALIDSTATUSCODE = "InvalidParameter.StatusCodeCacheInvalidStatusCode"
//  INVALIDPARAMETER_TLSVERSIONNOTINSEQUENCE = "InvalidParameter.TlsVersionNotInSequence"
//  INVALIDPARAMETERVALUE_FORMATMISMATCH = "InvalidParameterValue.FormatMismatch"
//  INVALIDPARAMETERVALUE_GENERALMISMATCH = "InvalidParameterValue.GeneralMismatch"
//  INVALIDPARAMETERVALUE_INCLUDEINVALIDVALUE = "InvalidParameterValue.IncludeInvalidValue"
//  INVALIDPARAMETERVALUE_MISSINGNECESSARYPARAM = "InvalidParameterValue.MissingNecessaryParam"
//  INVALIDPARAMETERVALUE_NOTINENUMERATION = "InvalidParameterValue.NotInEnumeration"
//  INVALIDPARAMETERVALUE_NOTWITHINRANGE = "InvalidParameterValue.NotWithinRange"
//  INVALIDPARAMETERVALUE_REGEXMISMATCH = "InvalidParameterValue.RegExMismatch"
//  INVALIDPARAMETERVALUE_TRIALPLANRESPONSEPAGE = "InvalidParameterValue.TrialPlanResponsePage"
//  INVALIDPARAMETERVALUE_UNRECOGNIZABLEVALUE = "InvalidParameterValue.UnrecognizableValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_INVALIDADVANCEDDEFENSESECURITYTYPE = "OperationDenied.InvalidAdvancedDefenseSecurityType"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ModifyRuleWithContext(ctx context.Context, request *ModifyRuleRequest) (response *ModifyRuleResponse, err error) {
    if request == nil {
        request = NewModifyRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifySecurityAPIResourceRequest() (request *ModifySecurityAPIResourceRequest) {
    request = &ModifySecurityAPIResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifySecurityAPIResource")
    
    
    return
}

func NewModifySecurityAPIResourceResponse() (response *ModifySecurityAPIResourceResponse) {
    response = &ModifySecurityAPIResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySecurityAPIResource
// This API is used to modify API resources.
//
// error code that may be returned:
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_CACHEKEYQUERYSTRINGTOOMANYVALUE = "InvalidParameter.CacheKeyQueryStringTooManyValue"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_COMPRESSIONINVALIDALGORITHMS = "InvalidParameter.CompressionInvalidAlgorithms"
//  INVALIDPARAMETER_ERRACTIONUNSUPPORTTARGET = "InvalidParameter.ErrActionUnsupportTarget"
//  INVALIDPARAMETER_ERRINVALIDACTION = "InvalidParameter.ErrInvalidAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONDUPLICATEACTION = "InvalidParameter.ErrInvalidActionDuplicateAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMACTION = "InvalidParameter.ErrInvalidActionParamAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMDUPLICATENAME = "InvalidParameter.ErrInvalidActionParamDuplicateName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMNAME = "InvalidParameter.ErrInvalidActionParamName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMTOOMANYVALUES = "InvalidParameter.ErrInvalidActionParamTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_ERRINVALIDACTIONTYPE = "InvalidParameter.ErrInvalidActionType"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONHOSTTOOMANYWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidConditionHostTooManyWhenModifyOriginActionConfigured"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONIGNORECASE = "InvalidParameter.ErrInvalidConditionIgnoreCase"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMEBADNAME = "InvalidParameter.ErrInvalidConditionNameBadName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMETARGETNOTSUPPORTNAME = "InvalidParameter.ErrInvalidConditionNameTargetNotSupportName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADREGULAR = "InvalidParameter.ErrInvalidConditionValueBadRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADURL = "InvalidParameter.ErrInvalidConditionValueBadUrl"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUECONTAINFILENAMEEXTENSION = "InvalidParameter.ErrInvalidConditionValueBadValueContainFileNameExtension"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOLONGVALUE = "InvalidParameter.ErrInvalidConditionValueTooLongValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYVALUES = "InvalidParameter.ErrInvalidConditionValueTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYWILDCARD = "InvalidParameter.ErrInvalidConditionValueTooManyWildcard"
//  INVALIDPARAMETER_ERRINVALIDELSEWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidElseWhenModifyOriginActionConfigured"
//  INVALIDPARAMETER_ERRNILCONDITION = "InvalidParameter.ErrNilCondition"
//  INVALIDPARAMETER_GRPCREQUIREHTTP2 = "InvalidParameter.GrpcRequireHttp2"
//  INVALIDPARAMETER_HOSTNOTFOUND = "InvalidParameter.HostNotFound"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESECRETKEY = "InvalidParameter.InvalidAuthenticationTypeSecretKey"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESIGNPARAM = "InvalidParameter.InvalidAuthenticationTypeSignParam"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPETIMEFORMAT = "InvalidParameter.InvalidAuthenticationTypeTimeFormat"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPETIMEPARAM = "InvalidParameter.InvalidAuthenticationTypeTimeParam"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDBACKUPSERVERNAME = "InvalidParameter.InvalidBackupServerName"
//  INVALIDPARAMETER_INVALIDCACHEKEY = "InvalidParameter.InvalidCacheKey"
//  INVALIDPARAMETER_INVALIDCACHEKEYCOOKIE = "InvalidParameter.InvalidCacheKeyCookie"
//  INVALIDPARAMETER_INVALIDCACHEKEYIGNORECASE = "InvalidParameter.InvalidCacheKeyIgnoreCase"
//  INVALIDPARAMETER_INVALIDCACHEKEYSCHEME = "InvalidParameter.InvalidCacheKeyScheme"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDCOSDOMAIN = "InvalidParameter.InvalidCosDomain"
//  INVALIDPARAMETER_INVALIDERRORPAGEREDIRECTURL = "InvalidParameter.InvalidErrorPageRedirectUrl"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGETIME = "InvalidParameter.InvalidMaxAgeTime"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAMEXFF = "InvalidParameter.InvalidRequestHeaderNameXff"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERVALUE = "InvalidParameter.InvalidRequestHeaderValue"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERNAME = "InvalidParameter.InvalidResponseHeaderName"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERVALUE = "InvalidParameter.InvalidResponseHeaderValue"
//  INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"
//  INVALIDPARAMETER_INVALIDRULEENGINENOTFOUND = "InvalidParameter.InvalidRuleEngineNotFound"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSEXTENSION = "InvalidParameter.InvalidRuleEngineTargetsExtension"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSURL = "InvalidParameter.InvalidRuleEngineTargetsUrl"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_INVALIDUPSTREAMREQUESTQUERYSTRINGVALUE = "InvalidParameter.InvalidUpstreamRequestQueryStringValue"
//  INVALIDPARAMETER_INVALIDURLREDIRECTHOST = "InvalidParameter.InvalidUrlRedirectHost"
//  INVALIDPARAMETER_INVALIDURLREDIRECTURL = "InvalidParameter.InvalidUrlRedirectUrl"
//  INVALIDPARAMETER_KEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.KeyRulesInvalidQueryStringValue"
//  INVALIDPARAMETER_LOADBALANCEINSTANCEIDISREQUIRED = "InvalidParameter.LoadBalanceInstanceIdIsRequired"
//  INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"
//  INVALIDPARAMETER_ORIGINORIGINGROUPIDISREQUIRED = "InvalidParameter.OriginOriginGroupIdIsRequired"
//  INVALIDPARAMETER_ORIGINPULLPROTOCOLISREQUIRED = "InvalidParameter.OriginPullProtocolIsRequired"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_RESPONSEHEADERCACHECONTROLNOTALLOWDELETE = "InvalidParameter.ResponseHeaderCacheControlNotAllowDelete"
//  INVALIDPARAMETER_STATUSCODECACHEINVALIDSTATUSCODE = "InvalidParameter.StatusCodeCacheInvalidStatusCode"
//  INVALIDPARAMETER_TLSVERSIONNOTINSEQUENCE = "InvalidParameter.TlsVersionNotInSequence"
//  INVALIDPARAMETERVALUE_FORMATMISMATCH = "InvalidParameterValue.FormatMismatch"
//  INVALIDPARAMETERVALUE_GENERALMISMATCH = "InvalidParameterValue.GeneralMismatch"
//  INVALIDPARAMETERVALUE_INCLUDEINVALIDVALUE = "InvalidParameterValue.IncludeInvalidValue"
//  INVALIDPARAMETERVALUE_MISSINGNECESSARYPARAM = "InvalidParameterValue.MissingNecessaryParam"
//  INVALIDPARAMETERVALUE_NOTINENUMERATION = "InvalidParameterValue.NotInEnumeration"
//  INVALIDPARAMETERVALUE_NOTWITHINRANGE = "InvalidParameterValue.NotWithinRange"
//  INVALIDPARAMETERVALUE_REGEXMISMATCH = "InvalidParameterValue.RegExMismatch"
//  INVALIDPARAMETERVALUE_TRIALPLANRESPONSEPAGE = "InvalidParameterValue.TrialPlanResponsePage"
//  INVALIDPARAMETERVALUE_UNRECOGNIZABLEVALUE = "InvalidParameterValue.UnrecognizableValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_INVALIDADVANCEDDEFENSESECURITYTYPE = "OperationDenied.InvalidAdvancedDefenseSecurityType"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ModifySecurityAPIResource(request *ModifySecurityAPIResourceRequest) (response *ModifySecurityAPIResourceResponse, err error) {
    return c.ModifySecurityAPIResourceWithContext(context.Background(), request)
}

// ModifySecurityAPIResource
// This API is used to modify API resources.
//
// error code that may be returned:
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_CACHEKEYQUERYSTRINGTOOMANYVALUE = "InvalidParameter.CacheKeyQueryStringTooManyValue"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_COMPRESSIONINVALIDALGORITHMS = "InvalidParameter.CompressionInvalidAlgorithms"
//  INVALIDPARAMETER_ERRACTIONUNSUPPORTTARGET = "InvalidParameter.ErrActionUnsupportTarget"
//  INVALIDPARAMETER_ERRINVALIDACTION = "InvalidParameter.ErrInvalidAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONDUPLICATEACTION = "InvalidParameter.ErrInvalidActionDuplicateAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMACTION = "InvalidParameter.ErrInvalidActionParamAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMDUPLICATENAME = "InvalidParameter.ErrInvalidActionParamDuplicateName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMNAME = "InvalidParameter.ErrInvalidActionParamName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMTOOMANYVALUES = "InvalidParameter.ErrInvalidActionParamTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_ERRINVALIDACTIONTYPE = "InvalidParameter.ErrInvalidActionType"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONHOSTTOOMANYWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidConditionHostTooManyWhenModifyOriginActionConfigured"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONIGNORECASE = "InvalidParameter.ErrInvalidConditionIgnoreCase"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMEBADNAME = "InvalidParameter.ErrInvalidConditionNameBadName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMETARGETNOTSUPPORTNAME = "InvalidParameter.ErrInvalidConditionNameTargetNotSupportName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADREGULAR = "InvalidParameter.ErrInvalidConditionValueBadRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADURL = "InvalidParameter.ErrInvalidConditionValueBadUrl"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUECONTAINFILENAMEEXTENSION = "InvalidParameter.ErrInvalidConditionValueBadValueContainFileNameExtension"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOLONGVALUE = "InvalidParameter.ErrInvalidConditionValueTooLongValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYVALUES = "InvalidParameter.ErrInvalidConditionValueTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYWILDCARD = "InvalidParameter.ErrInvalidConditionValueTooManyWildcard"
//  INVALIDPARAMETER_ERRINVALIDELSEWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidElseWhenModifyOriginActionConfigured"
//  INVALIDPARAMETER_ERRNILCONDITION = "InvalidParameter.ErrNilCondition"
//  INVALIDPARAMETER_GRPCREQUIREHTTP2 = "InvalidParameter.GrpcRequireHttp2"
//  INVALIDPARAMETER_HOSTNOTFOUND = "InvalidParameter.HostNotFound"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESECRETKEY = "InvalidParameter.InvalidAuthenticationTypeSecretKey"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESIGNPARAM = "InvalidParameter.InvalidAuthenticationTypeSignParam"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPETIMEFORMAT = "InvalidParameter.InvalidAuthenticationTypeTimeFormat"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPETIMEPARAM = "InvalidParameter.InvalidAuthenticationTypeTimeParam"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDBACKUPSERVERNAME = "InvalidParameter.InvalidBackupServerName"
//  INVALIDPARAMETER_INVALIDCACHEKEY = "InvalidParameter.InvalidCacheKey"
//  INVALIDPARAMETER_INVALIDCACHEKEYCOOKIE = "InvalidParameter.InvalidCacheKeyCookie"
//  INVALIDPARAMETER_INVALIDCACHEKEYIGNORECASE = "InvalidParameter.InvalidCacheKeyIgnoreCase"
//  INVALIDPARAMETER_INVALIDCACHEKEYSCHEME = "InvalidParameter.InvalidCacheKeyScheme"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDCOSDOMAIN = "InvalidParameter.InvalidCosDomain"
//  INVALIDPARAMETER_INVALIDERRORPAGEREDIRECTURL = "InvalidParameter.InvalidErrorPageRedirectUrl"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGETIME = "InvalidParameter.InvalidMaxAgeTime"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAMEXFF = "InvalidParameter.InvalidRequestHeaderNameXff"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERVALUE = "InvalidParameter.InvalidRequestHeaderValue"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERNAME = "InvalidParameter.InvalidResponseHeaderName"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERVALUE = "InvalidParameter.InvalidResponseHeaderValue"
//  INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"
//  INVALIDPARAMETER_INVALIDRULEENGINENOTFOUND = "InvalidParameter.InvalidRuleEngineNotFound"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSEXTENSION = "InvalidParameter.InvalidRuleEngineTargetsExtension"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSURL = "InvalidParameter.InvalidRuleEngineTargetsUrl"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_INVALIDUPSTREAMREQUESTQUERYSTRINGVALUE = "InvalidParameter.InvalidUpstreamRequestQueryStringValue"
//  INVALIDPARAMETER_INVALIDURLREDIRECTHOST = "InvalidParameter.InvalidUrlRedirectHost"
//  INVALIDPARAMETER_INVALIDURLREDIRECTURL = "InvalidParameter.InvalidUrlRedirectUrl"
//  INVALIDPARAMETER_KEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.KeyRulesInvalidQueryStringValue"
//  INVALIDPARAMETER_LOADBALANCEINSTANCEIDISREQUIRED = "InvalidParameter.LoadBalanceInstanceIdIsRequired"
//  INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"
//  INVALIDPARAMETER_ORIGINORIGINGROUPIDISREQUIRED = "InvalidParameter.OriginOriginGroupIdIsRequired"
//  INVALIDPARAMETER_ORIGINPULLPROTOCOLISREQUIRED = "InvalidParameter.OriginPullProtocolIsRequired"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_RESPONSEHEADERCACHECONTROLNOTALLOWDELETE = "InvalidParameter.ResponseHeaderCacheControlNotAllowDelete"
//  INVALIDPARAMETER_STATUSCODECACHEINVALIDSTATUSCODE = "InvalidParameter.StatusCodeCacheInvalidStatusCode"
//  INVALIDPARAMETER_TLSVERSIONNOTINSEQUENCE = "InvalidParameter.TlsVersionNotInSequence"
//  INVALIDPARAMETERVALUE_FORMATMISMATCH = "InvalidParameterValue.FormatMismatch"
//  INVALIDPARAMETERVALUE_GENERALMISMATCH = "InvalidParameterValue.GeneralMismatch"
//  INVALIDPARAMETERVALUE_INCLUDEINVALIDVALUE = "InvalidParameterValue.IncludeInvalidValue"
//  INVALIDPARAMETERVALUE_MISSINGNECESSARYPARAM = "InvalidParameterValue.MissingNecessaryParam"
//  INVALIDPARAMETERVALUE_NOTINENUMERATION = "InvalidParameterValue.NotInEnumeration"
//  INVALIDPARAMETERVALUE_NOTWITHINRANGE = "InvalidParameterValue.NotWithinRange"
//  INVALIDPARAMETERVALUE_REGEXMISMATCH = "InvalidParameterValue.RegExMismatch"
//  INVALIDPARAMETERVALUE_TRIALPLANRESPONSEPAGE = "InvalidParameterValue.TrialPlanResponsePage"
//  INVALIDPARAMETERVALUE_UNRECOGNIZABLEVALUE = "InvalidParameterValue.UnrecognizableValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_INVALIDADVANCEDDEFENSESECURITYTYPE = "OperationDenied.InvalidAdvancedDefenseSecurityType"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ModifySecurityAPIResourceWithContext(ctx context.Context, request *ModifySecurityAPIResourceRequest) (response *ModifySecurityAPIResourceResponse, err error) {
    if request == nil {
        request = NewModifySecurityAPIResourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifySecurityAPIResource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySecurityAPIResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySecurityAPIResourceResponse()
    err = c.Send(request, response)
    return
}

func NewModifySecurityAPIServiceRequest() (request *ModifySecurityAPIServiceRequest) {
    request = &ModifySecurityAPIServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifySecurityAPIService")
    
    
    return
}

func NewModifySecurityAPIServiceResponse() (response *ModifySecurityAPIServiceResponse) {
    response = &ModifySecurityAPIServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySecurityAPIService
// This API is used to modify API services.
//
// error code that may be returned:
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_CACHEKEYQUERYSTRINGTOOMANYVALUE = "InvalidParameter.CacheKeyQueryStringTooManyValue"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_COMPRESSIONINVALIDALGORITHMS = "InvalidParameter.CompressionInvalidAlgorithms"
//  INVALIDPARAMETER_ERRACTIONUNSUPPORTTARGET = "InvalidParameter.ErrActionUnsupportTarget"
//  INVALIDPARAMETER_ERRINVALIDACTION = "InvalidParameter.ErrInvalidAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONDUPLICATEACTION = "InvalidParameter.ErrInvalidActionDuplicateAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMACTION = "InvalidParameter.ErrInvalidActionParamAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMDUPLICATENAME = "InvalidParameter.ErrInvalidActionParamDuplicateName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMNAME = "InvalidParameter.ErrInvalidActionParamName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMTOOMANYVALUES = "InvalidParameter.ErrInvalidActionParamTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_ERRINVALIDACTIONTYPE = "InvalidParameter.ErrInvalidActionType"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONHOSTTOOMANYWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidConditionHostTooManyWhenModifyOriginActionConfigured"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONIGNORECASE = "InvalidParameter.ErrInvalidConditionIgnoreCase"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMEBADNAME = "InvalidParameter.ErrInvalidConditionNameBadName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMETARGETNOTSUPPORTNAME = "InvalidParameter.ErrInvalidConditionNameTargetNotSupportName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADREGULAR = "InvalidParameter.ErrInvalidConditionValueBadRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADURL = "InvalidParameter.ErrInvalidConditionValueBadUrl"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUECONTAINFILENAMEEXTENSION = "InvalidParameter.ErrInvalidConditionValueBadValueContainFileNameExtension"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOLONGVALUE = "InvalidParameter.ErrInvalidConditionValueTooLongValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYVALUES = "InvalidParameter.ErrInvalidConditionValueTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYWILDCARD = "InvalidParameter.ErrInvalidConditionValueTooManyWildcard"
//  INVALIDPARAMETER_ERRINVALIDELSEWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidElseWhenModifyOriginActionConfigured"
//  INVALIDPARAMETER_ERRNILCONDITION = "InvalidParameter.ErrNilCondition"
//  INVALIDPARAMETER_GRPCREQUIREHTTP2 = "InvalidParameter.GrpcRequireHttp2"
//  INVALIDPARAMETER_HOSTNOTFOUND = "InvalidParameter.HostNotFound"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESECRETKEY = "InvalidParameter.InvalidAuthenticationTypeSecretKey"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESIGNPARAM = "InvalidParameter.InvalidAuthenticationTypeSignParam"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPETIMEFORMAT = "InvalidParameter.InvalidAuthenticationTypeTimeFormat"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPETIMEPARAM = "InvalidParameter.InvalidAuthenticationTypeTimeParam"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDBACKUPSERVERNAME = "InvalidParameter.InvalidBackupServerName"
//  INVALIDPARAMETER_INVALIDCACHEKEY = "InvalidParameter.InvalidCacheKey"
//  INVALIDPARAMETER_INVALIDCACHEKEYCOOKIE = "InvalidParameter.InvalidCacheKeyCookie"
//  INVALIDPARAMETER_INVALIDCACHEKEYIGNORECASE = "InvalidParameter.InvalidCacheKeyIgnoreCase"
//  INVALIDPARAMETER_INVALIDCACHEKEYSCHEME = "InvalidParameter.InvalidCacheKeyScheme"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDCOSDOMAIN = "InvalidParameter.InvalidCosDomain"
//  INVALIDPARAMETER_INVALIDERRORPAGEREDIRECTURL = "InvalidParameter.InvalidErrorPageRedirectUrl"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGETIME = "InvalidParameter.InvalidMaxAgeTime"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAMEXFF = "InvalidParameter.InvalidRequestHeaderNameXff"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERVALUE = "InvalidParameter.InvalidRequestHeaderValue"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERNAME = "InvalidParameter.InvalidResponseHeaderName"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERVALUE = "InvalidParameter.InvalidResponseHeaderValue"
//  INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"
//  INVALIDPARAMETER_INVALIDRULEENGINENOTFOUND = "InvalidParameter.InvalidRuleEngineNotFound"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSEXTENSION = "InvalidParameter.InvalidRuleEngineTargetsExtension"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSURL = "InvalidParameter.InvalidRuleEngineTargetsUrl"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_INVALIDUPSTREAMREQUESTQUERYSTRINGVALUE = "InvalidParameter.InvalidUpstreamRequestQueryStringValue"
//  INVALIDPARAMETER_INVALIDURLREDIRECTHOST = "InvalidParameter.InvalidUrlRedirectHost"
//  INVALIDPARAMETER_INVALIDURLREDIRECTURL = "InvalidParameter.InvalidUrlRedirectUrl"
//  INVALIDPARAMETER_KEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.KeyRulesInvalidQueryStringValue"
//  INVALIDPARAMETER_LOADBALANCEINSTANCEIDISREQUIRED = "InvalidParameter.LoadBalanceInstanceIdIsRequired"
//  INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"
//  INVALIDPARAMETER_ORIGINORIGINGROUPIDISREQUIRED = "InvalidParameter.OriginOriginGroupIdIsRequired"
//  INVALIDPARAMETER_ORIGINPULLPROTOCOLISREQUIRED = "InvalidParameter.OriginPullProtocolIsRequired"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_RESPONSEHEADERCACHECONTROLNOTALLOWDELETE = "InvalidParameter.ResponseHeaderCacheControlNotAllowDelete"
//  INVALIDPARAMETER_STATUSCODECACHEINVALIDSTATUSCODE = "InvalidParameter.StatusCodeCacheInvalidStatusCode"
//  INVALIDPARAMETER_TLSVERSIONNOTINSEQUENCE = "InvalidParameter.TlsVersionNotInSequence"
//  INVALIDPARAMETERVALUE_FORMATMISMATCH = "InvalidParameterValue.FormatMismatch"
//  INVALIDPARAMETERVALUE_GENERALMISMATCH = "InvalidParameterValue.GeneralMismatch"
//  INVALIDPARAMETERVALUE_INCLUDEINVALIDVALUE = "InvalidParameterValue.IncludeInvalidValue"
//  INVALIDPARAMETERVALUE_MISSINGNECESSARYPARAM = "InvalidParameterValue.MissingNecessaryParam"
//  INVALIDPARAMETERVALUE_NOTINENUMERATION = "InvalidParameterValue.NotInEnumeration"
//  INVALIDPARAMETERVALUE_NOTWITHINRANGE = "InvalidParameterValue.NotWithinRange"
//  INVALIDPARAMETERVALUE_REGEXMISMATCH = "InvalidParameterValue.RegExMismatch"
//  INVALIDPARAMETERVALUE_TRIALPLANRESPONSEPAGE = "InvalidParameterValue.TrialPlanResponsePage"
//  INVALIDPARAMETERVALUE_UNRECOGNIZABLEVALUE = "InvalidParameterValue.UnrecognizableValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_INVALIDADVANCEDDEFENSESECURITYTYPE = "OperationDenied.InvalidAdvancedDefenseSecurityType"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ModifySecurityAPIService(request *ModifySecurityAPIServiceRequest) (response *ModifySecurityAPIServiceResponse, err error) {
    return c.ModifySecurityAPIServiceWithContext(context.Background(), request)
}

// ModifySecurityAPIService
// This API is used to modify API services.
//
// error code that may be returned:
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_CACHEKEYQUERYSTRINGTOOMANYVALUE = "InvalidParameter.CacheKeyQueryStringTooManyValue"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_COMPRESSIONINVALIDALGORITHMS = "InvalidParameter.CompressionInvalidAlgorithms"
//  INVALIDPARAMETER_ERRACTIONUNSUPPORTTARGET = "InvalidParameter.ErrActionUnsupportTarget"
//  INVALIDPARAMETER_ERRINVALIDACTION = "InvalidParameter.ErrInvalidAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONDUPLICATEACTION = "InvalidParameter.ErrInvalidActionDuplicateAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMACTION = "InvalidParameter.ErrInvalidActionParamAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMDUPLICATENAME = "InvalidParameter.ErrInvalidActionParamDuplicateName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMNAME = "InvalidParameter.ErrInvalidActionParamName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMTOOMANYVALUES = "InvalidParameter.ErrInvalidActionParamTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_ERRINVALIDACTIONTYPE = "InvalidParameter.ErrInvalidActionType"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONHOSTTOOMANYWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidConditionHostTooManyWhenModifyOriginActionConfigured"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONIGNORECASE = "InvalidParameter.ErrInvalidConditionIgnoreCase"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMEBADNAME = "InvalidParameter.ErrInvalidConditionNameBadName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMETARGETNOTSUPPORTNAME = "InvalidParameter.ErrInvalidConditionNameTargetNotSupportName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADREGULAR = "InvalidParameter.ErrInvalidConditionValueBadRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADURL = "InvalidParameter.ErrInvalidConditionValueBadUrl"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUECONTAINFILENAMEEXTENSION = "InvalidParameter.ErrInvalidConditionValueBadValueContainFileNameExtension"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOLONGVALUE = "InvalidParameter.ErrInvalidConditionValueTooLongValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYVALUES = "InvalidParameter.ErrInvalidConditionValueTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYWILDCARD = "InvalidParameter.ErrInvalidConditionValueTooManyWildcard"
//  INVALIDPARAMETER_ERRINVALIDELSEWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidElseWhenModifyOriginActionConfigured"
//  INVALIDPARAMETER_ERRNILCONDITION = "InvalidParameter.ErrNilCondition"
//  INVALIDPARAMETER_GRPCREQUIREHTTP2 = "InvalidParameter.GrpcRequireHttp2"
//  INVALIDPARAMETER_HOSTNOTFOUND = "InvalidParameter.HostNotFound"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESECRETKEY = "InvalidParameter.InvalidAuthenticationTypeSecretKey"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESIGNPARAM = "InvalidParameter.InvalidAuthenticationTypeSignParam"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPETIMEFORMAT = "InvalidParameter.InvalidAuthenticationTypeTimeFormat"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPETIMEPARAM = "InvalidParameter.InvalidAuthenticationTypeTimeParam"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDBACKUPSERVERNAME = "InvalidParameter.InvalidBackupServerName"
//  INVALIDPARAMETER_INVALIDCACHEKEY = "InvalidParameter.InvalidCacheKey"
//  INVALIDPARAMETER_INVALIDCACHEKEYCOOKIE = "InvalidParameter.InvalidCacheKeyCookie"
//  INVALIDPARAMETER_INVALIDCACHEKEYIGNORECASE = "InvalidParameter.InvalidCacheKeyIgnoreCase"
//  INVALIDPARAMETER_INVALIDCACHEKEYSCHEME = "InvalidParameter.InvalidCacheKeyScheme"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDCOSDOMAIN = "InvalidParameter.InvalidCosDomain"
//  INVALIDPARAMETER_INVALIDERRORPAGEREDIRECTURL = "InvalidParameter.InvalidErrorPageRedirectUrl"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGETIME = "InvalidParameter.InvalidMaxAgeTime"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAMEXFF = "InvalidParameter.InvalidRequestHeaderNameXff"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERVALUE = "InvalidParameter.InvalidRequestHeaderValue"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERNAME = "InvalidParameter.InvalidResponseHeaderName"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERVALUE = "InvalidParameter.InvalidResponseHeaderValue"
//  INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"
//  INVALIDPARAMETER_INVALIDRULEENGINENOTFOUND = "InvalidParameter.InvalidRuleEngineNotFound"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSEXTENSION = "InvalidParameter.InvalidRuleEngineTargetsExtension"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSURL = "InvalidParameter.InvalidRuleEngineTargetsUrl"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_INVALIDUPSTREAMREQUESTQUERYSTRINGVALUE = "InvalidParameter.InvalidUpstreamRequestQueryStringValue"
//  INVALIDPARAMETER_INVALIDURLREDIRECTHOST = "InvalidParameter.InvalidUrlRedirectHost"
//  INVALIDPARAMETER_INVALIDURLREDIRECTURL = "InvalidParameter.InvalidUrlRedirectUrl"
//  INVALIDPARAMETER_KEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.KeyRulesInvalidQueryStringValue"
//  INVALIDPARAMETER_LOADBALANCEINSTANCEIDISREQUIRED = "InvalidParameter.LoadBalanceInstanceIdIsRequired"
//  INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"
//  INVALIDPARAMETER_ORIGINORIGINGROUPIDISREQUIRED = "InvalidParameter.OriginOriginGroupIdIsRequired"
//  INVALIDPARAMETER_ORIGINPULLPROTOCOLISREQUIRED = "InvalidParameter.OriginPullProtocolIsRequired"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_RESPONSEHEADERCACHECONTROLNOTALLOWDELETE = "InvalidParameter.ResponseHeaderCacheControlNotAllowDelete"
//  INVALIDPARAMETER_STATUSCODECACHEINVALIDSTATUSCODE = "InvalidParameter.StatusCodeCacheInvalidStatusCode"
//  INVALIDPARAMETER_TLSVERSIONNOTINSEQUENCE = "InvalidParameter.TlsVersionNotInSequence"
//  INVALIDPARAMETERVALUE_FORMATMISMATCH = "InvalidParameterValue.FormatMismatch"
//  INVALIDPARAMETERVALUE_GENERALMISMATCH = "InvalidParameterValue.GeneralMismatch"
//  INVALIDPARAMETERVALUE_INCLUDEINVALIDVALUE = "InvalidParameterValue.IncludeInvalidValue"
//  INVALIDPARAMETERVALUE_MISSINGNECESSARYPARAM = "InvalidParameterValue.MissingNecessaryParam"
//  INVALIDPARAMETERVALUE_NOTINENUMERATION = "InvalidParameterValue.NotInEnumeration"
//  INVALIDPARAMETERVALUE_NOTWITHINRANGE = "InvalidParameterValue.NotWithinRange"
//  INVALIDPARAMETERVALUE_REGEXMISMATCH = "InvalidParameterValue.RegExMismatch"
//  INVALIDPARAMETERVALUE_TRIALPLANRESPONSEPAGE = "InvalidParameterValue.TrialPlanResponsePage"
//  INVALIDPARAMETERVALUE_UNRECOGNIZABLEVALUE = "InvalidParameterValue.UnrecognizableValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_INVALIDADVANCEDDEFENSESECURITYTYPE = "OperationDenied.InvalidAdvancedDefenseSecurityType"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ModifySecurityAPIServiceWithContext(ctx context.Context, request *ModifySecurityAPIServiceRequest) (response *ModifySecurityAPIServiceResponse, err error) {
    if request == nil {
        request = NewModifySecurityAPIServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifySecurityAPIService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySecurityAPIService require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySecurityAPIServiceResponse()
    err = c.Send(request, response)
    return
}

func NewModifySecurityClientAttesterRequest() (request *ModifySecurityClientAttesterRequest) {
    request = &ModifySecurityClientAttesterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifySecurityClientAttester")
    
    
    return
}

func NewModifySecurityClientAttesterResponse() (response *ModifySecurityClientAttesterResponse) {
    response = &ModifySecurityClientAttesterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySecurityClientAttester
// Modify client authentication options.
//
// error code that may be returned:
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_CACHEKEYQUERYSTRINGTOOMANYVALUE = "InvalidParameter.CacheKeyQueryStringTooManyValue"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_COMPRESSIONINVALIDALGORITHMS = "InvalidParameter.CompressionInvalidAlgorithms"
//  INVALIDPARAMETER_ERRACTIONUNSUPPORTTARGET = "InvalidParameter.ErrActionUnsupportTarget"
//  INVALIDPARAMETER_ERRINVALIDACTION = "InvalidParameter.ErrInvalidAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONDUPLICATEACTION = "InvalidParameter.ErrInvalidActionDuplicateAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMACTION = "InvalidParameter.ErrInvalidActionParamAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMDUPLICATENAME = "InvalidParameter.ErrInvalidActionParamDuplicateName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMNAME = "InvalidParameter.ErrInvalidActionParamName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMTOOMANYVALUES = "InvalidParameter.ErrInvalidActionParamTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_ERRINVALIDACTIONTYPE = "InvalidParameter.ErrInvalidActionType"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONHOSTTOOMANYWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidConditionHostTooManyWhenModifyOriginActionConfigured"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONIGNORECASE = "InvalidParameter.ErrInvalidConditionIgnoreCase"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMEBADNAME = "InvalidParameter.ErrInvalidConditionNameBadName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMETARGETNOTSUPPORTNAME = "InvalidParameter.ErrInvalidConditionNameTargetNotSupportName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADREGULAR = "InvalidParameter.ErrInvalidConditionValueBadRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADURL = "InvalidParameter.ErrInvalidConditionValueBadUrl"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUECONTAINFILENAMEEXTENSION = "InvalidParameter.ErrInvalidConditionValueBadValueContainFileNameExtension"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOLONGVALUE = "InvalidParameter.ErrInvalidConditionValueTooLongValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYVALUES = "InvalidParameter.ErrInvalidConditionValueTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYWILDCARD = "InvalidParameter.ErrInvalidConditionValueTooManyWildcard"
//  INVALIDPARAMETER_ERRINVALIDELSEWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidElseWhenModifyOriginActionConfigured"
//  INVALIDPARAMETER_ERRNILCONDITION = "InvalidParameter.ErrNilCondition"
//  INVALIDPARAMETER_GRPCREQUIREHTTP2 = "InvalidParameter.GrpcRequireHttp2"
//  INVALIDPARAMETER_HOSTNOTFOUND = "InvalidParameter.HostNotFound"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESECRETKEY = "InvalidParameter.InvalidAuthenticationTypeSecretKey"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESIGNPARAM = "InvalidParameter.InvalidAuthenticationTypeSignParam"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPETIMEFORMAT = "InvalidParameter.InvalidAuthenticationTypeTimeFormat"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPETIMEPARAM = "InvalidParameter.InvalidAuthenticationTypeTimeParam"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDBACKUPSERVERNAME = "InvalidParameter.InvalidBackupServerName"
//  INVALIDPARAMETER_INVALIDCACHEKEY = "InvalidParameter.InvalidCacheKey"
//  INVALIDPARAMETER_INVALIDCACHEKEYCOOKIE = "InvalidParameter.InvalidCacheKeyCookie"
//  INVALIDPARAMETER_INVALIDCACHEKEYIGNORECASE = "InvalidParameter.InvalidCacheKeyIgnoreCase"
//  INVALIDPARAMETER_INVALIDCACHEKEYSCHEME = "InvalidParameter.InvalidCacheKeyScheme"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDCOSDOMAIN = "InvalidParameter.InvalidCosDomain"
//  INVALIDPARAMETER_INVALIDERRORPAGEREDIRECTURL = "InvalidParameter.InvalidErrorPageRedirectUrl"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGETIME = "InvalidParameter.InvalidMaxAgeTime"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAMEXFF = "InvalidParameter.InvalidRequestHeaderNameXff"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERVALUE = "InvalidParameter.InvalidRequestHeaderValue"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERNAME = "InvalidParameter.InvalidResponseHeaderName"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERVALUE = "InvalidParameter.InvalidResponseHeaderValue"
//  INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"
//  INVALIDPARAMETER_INVALIDRULEENGINENOTFOUND = "InvalidParameter.InvalidRuleEngineNotFound"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSEXTENSION = "InvalidParameter.InvalidRuleEngineTargetsExtension"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSURL = "InvalidParameter.InvalidRuleEngineTargetsUrl"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_INVALIDUPSTREAMREQUESTQUERYSTRINGVALUE = "InvalidParameter.InvalidUpstreamRequestQueryStringValue"
//  INVALIDPARAMETER_INVALIDURLREDIRECTHOST = "InvalidParameter.InvalidUrlRedirectHost"
//  INVALIDPARAMETER_INVALIDURLREDIRECTURL = "InvalidParameter.InvalidUrlRedirectUrl"
//  INVALIDPARAMETER_KEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.KeyRulesInvalidQueryStringValue"
//  INVALIDPARAMETER_LOADBALANCEINSTANCEIDISREQUIRED = "InvalidParameter.LoadBalanceInstanceIdIsRequired"
//  INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"
//  INVALIDPARAMETER_ORIGINORIGINGROUPIDISREQUIRED = "InvalidParameter.OriginOriginGroupIdIsRequired"
//  INVALIDPARAMETER_ORIGINPULLPROTOCOLISREQUIRED = "InvalidParameter.OriginPullProtocolIsRequired"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_RESPONSEHEADERCACHECONTROLNOTALLOWDELETE = "InvalidParameter.ResponseHeaderCacheControlNotAllowDelete"
//  INVALIDPARAMETER_STATUSCODECACHEINVALIDSTATUSCODE = "InvalidParameter.StatusCodeCacheInvalidStatusCode"
//  INVALIDPARAMETER_TLSVERSIONNOTINSEQUENCE = "InvalidParameter.TlsVersionNotInSequence"
//  INVALIDPARAMETERVALUE_FORMATMISMATCH = "InvalidParameterValue.FormatMismatch"
//  INVALIDPARAMETERVALUE_GENERALMISMATCH = "InvalidParameterValue.GeneralMismatch"
//  INVALIDPARAMETERVALUE_INCLUDEINVALIDVALUE = "InvalidParameterValue.IncludeInvalidValue"
//  INVALIDPARAMETERVALUE_MISSINGNECESSARYPARAM = "InvalidParameterValue.MissingNecessaryParam"
//  INVALIDPARAMETERVALUE_NOTINENUMERATION = "InvalidParameterValue.NotInEnumeration"
//  INVALIDPARAMETERVALUE_NOTWITHINRANGE = "InvalidParameterValue.NotWithinRange"
//  INVALIDPARAMETERVALUE_REGEXMISMATCH = "InvalidParameterValue.RegExMismatch"
//  INVALIDPARAMETERVALUE_TRIALPLANRESPONSEPAGE = "InvalidParameterValue.TrialPlanResponsePage"
//  INVALIDPARAMETERVALUE_UNRECOGNIZABLEVALUE = "InvalidParameterValue.UnrecognizableValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_INVALIDADVANCEDDEFENSESECURITYTYPE = "OperationDenied.InvalidAdvancedDefenseSecurityType"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ModifySecurityClientAttester(request *ModifySecurityClientAttesterRequest) (response *ModifySecurityClientAttesterResponse, err error) {
    return c.ModifySecurityClientAttesterWithContext(context.Background(), request)
}

// ModifySecurityClientAttester
// Modify client authentication options.
//
// error code that may be returned:
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_CACHEKEYQUERYSTRINGTOOMANYVALUE = "InvalidParameter.CacheKeyQueryStringTooManyValue"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_COMPRESSIONINVALIDALGORITHMS = "InvalidParameter.CompressionInvalidAlgorithms"
//  INVALIDPARAMETER_ERRACTIONUNSUPPORTTARGET = "InvalidParameter.ErrActionUnsupportTarget"
//  INVALIDPARAMETER_ERRINVALIDACTION = "InvalidParameter.ErrInvalidAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONDUPLICATEACTION = "InvalidParameter.ErrInvalidActionDuplicateAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMACTION = "InvalidParameter.ErrInvalidActionParamAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMDUPLICATENAME = "InvalidParameter.ErrInvalidActionParamDuplicateName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMNAME = "InvalidParameter.ErrInvalidActionParamName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMTOOMANYVALUES = "InvalidParameter.ErrInvalidActionParamTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_ERRINVALIDACTIONTYPE = "InvalidParameter.ErrInvalidActionType"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONHOSTTOOMANYWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidConditionHostTooManyWhenModifyOriginActionConfigured"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONIGNORECASE = "InvalidParameter.ErrInvalidConditionIgnoreCase"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMEBADNAME = "InvalidParameter.ErrInvalidConditionNameBadName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMETARGETNOTSUPPORTNAME = "InvalidParameter.ErrInvalidConditionNameTargetNotSupportName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADREGULAR = "InvalidParameter.ErrInvalidConditionValueBadRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADURL = "InvalidParameter.ErrInvalidConditionValueBadUrl"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUECONTAINFILENAMEEXTENSION = "InvalidParameter.ErrInvalidConditionValueBadValueContainFileNameExtension"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOLONGVALUE = "InvalidParameter.ErrInvalidConditionValueTooLongValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYVALUES = "InvalidParameter.ErrInvalidConditionValueTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYWILDCARD = "InvalidParameter.ErrInvalidConditionValueTooManyWildcard"
//  INVALIDPARAMETER_ERRINVALIDELSEWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidElseWhenModifyOriginActionConfigured"
//  INVALIDPARAMETER_ERRNILCONDITION = "InvalidParameter.ErrNilCondition"
//  INVALIDPARAMETER_GRPCREQUIREHTTP2 = "InvalidParameter.GrpcRequireHttp2"
//  INVALIDPARAMETER_HOSTNOTFOUND = "InvalidParameter.HostNotFound"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESECRETKEY = "InvalidParameter.InvalidAuthenticationTypeSecretKey"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESIGNPARAM = "InvalidParameter.InvalidAuthenticationTypeSignParam"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPETIMEFORMAT = "InvalidParameter.InvalidAuthenticationTypeTimeFormat"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPETIMEPARAM = "InvalidParameter.InvalidAuthenticationTypeTimeParam"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDBACKUPSERVERNAME = "InvalidParameter.InvalidBackupServerName"
//  INVALIDPARAMETER_INVALIDCACHEKEY = "InvalidParameter.InvalidCacheKey"
//  INVALIDPARAMETER_INVALIDCACHEKEYCOOKIE = "InvalidParameter.InvalidCacheKeyCookie"
//  INVALIDPARAMETER_INVALIDCACHEKEYIGNORECASE = "InvalidParameter.InvalidCacheKeyIgnoreCase"
//  INVALIDPARAMETER_INVALIDCACHEKEYSCHEME = "InvalidParameter.InvalidCacheKeyScheme"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDCOSDOMAIN = "InvalidParameter.InvalidCosDomain"
//  INVALIDPARAMETER_INVALIDERRORPAGEREDIRECTURL = "InvalidParameter.InvalidErrorPageRedirectUrl"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGETIME = "InvalidParameter.InvalidMaxAgeTime"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAMEXFF = "InvalidParameter.InvalidRequestHeaderNameXff"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERVALUE = "InvalidParameter.InvalidRequestHeaderValue"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERNAME = "InvalidParameter.InvalidResponseHeaderName"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERVALUE = "InvalidParameter.InvalidResponseHeaderValue"
//  INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"
//  INVALIDPARAMETER_INVALIDRULEENGINENOTFOUND = "InvalidParameter.InvalidRuleEngineNotFound"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSEXTENSION = "InvalidParameter.InvalidRuleEngineTargetsExtension"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSURL = "InvalidParameter.InvalidRuleEngineTargetsUrl"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_INVALIDUPSTREAMREQUESTQUERYSTRINGVALUE = "InvalidParameter.InvalidUpstreamRequestQueryStringValue"
//  INVALIDPARAMETER_INVALIDURLREDIRECTHOST = "InvalidParameter.InvalidUrlRedirectHost"
//  INVALIDPARAMETER_INVALIDURLREDIRECTURL = "InvalidParameter.InvalidUrlRedirectUrl"
//  INVALIDPARAMETER_KEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.KeyRulesInvalidQueryStringValue"
//  INVALIDPARAMETER_LOADBALANCEINSTANCEIDISREQUIRED = "InvalidParameter.LoadBalanceInstanceIdIsRequired"
//  INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"
//  INVALIDPARAMETER_ORIGINORIGINGROUPIDISREQUIRED = "InvalidParameter.OriginOriginGroupIdIsRequired"
//  INVALIDPARAMETER_ORIGINPULLPROTOCOLISREQUIRED = "InvalidParameter.OriginPullProtocolIsRequired"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_RESPONSEHEADERCACHECONTROLNOTALLOWDELETE = "InvalidParameter.ResponseHeaderCacheControlNotAllowDelete"
//  INVALIDPARAMETER_STATUSCODECACHEINVALIDSTATUSCODE = "InvalidParameter.StatusCodeCacheInvalidStatusCode"
//  INVALIDPARAMETER_TLSVERSIONNOTINSEQUENCE = "InvalidParameter.TlsVersionNotInSequence"
//  INVALIDPARAMETERVALUE_FORMATMISMATCH = "InvalidParameterValue.FormatMismatch"
//  INVALIDPARAMETERVALUE_GENERALMISMATCH = "InvalidParameterValue.GeneralMismatch"
//  INVALIDPARAMETERVALUE_INCLUDEINVALIDVALUE = "InvalidParameterValue.IncludeInvalidValue"
//  INVALIDPARAMETERVALUE_MISSINGNECESSARYPARAM = "InvalidParameterValue.MissingNecessaryParam"
//  INVALIDPARAMETERVALUE_NOTINENUMERATION = "InvalidParameterValue.NotInEnumeration"
//  INVALIDPARAMETERVALUE_NOTWITHINRANGE = "InvalidParameterValue.NotWithinRange"
//  INVALIDPARAMETERVALUE_REGEXMISMATCH = "InvalidParameterValue.RegExMismatch"
//  INVALIDPARAMETERVALUE_TRIALPLANRESPONSEPAGE = "InvalidParameterValue.TrialPlanResponsePage"
//  INVALIDPARAMETERVALUE_UNRECOGNIZABLEVALUE = "InvalidParameterValue.UnrecognizableValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_INVALIDADVANCEDDEFENSESECURITYTYPE = "OperationDenied.InvalidAdvancedDefenseSecurityType"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ModifySecurityClientAttesterWithContext(ctx context.Context, request *ModifySecurityClientAttesterRequest) (response *ModifySecurityClientAttesterResponse, err error) {
    if request == nil {
        request = NewModifySecurityClientAttesterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifySecurityClientAttester")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySecurityClientAttester require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySecurityClientAttesterResponse()
    err = c.Send(request, response)
    return
}

func NewModifySecurityIPGroupRequest() (request *ModifySecurityIPGroupRequest) {
    request = &ModifySecurityIPGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifySecurityIPGroup")
    
    
    return
}

func NewModifySecurityIPGroupResponse() (response *ModifySecurityIPGroupResponse) {
    response = &ModifySecurityIPGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySecurityIPGroup
// Modifies a security IP group.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ModifySecurityIPGroup(request *ModifySecurityIPGroupRequest) (response *ModifySecurityIPGroupResponse, err error) {
    return c.ModifySecurityIPGroupWithContext(context.Background(), request)
}

// ModifySecurityIPGroup
// Modifies a security IP group.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ModifySecurityIPGroupWithContext(ctx context.Context, request *ModifySecurityIPGroupRequest) (response *ModifySecurityIPGroupResponse, err error) {
    if request == nil {
        request = NewModifySecurityIPGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifySecurityIPGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySecurityIPGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySecurityIPGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifySecurityJSInjectionRuleRequest() (request *ModifySecurityJSInjectionRuleRequest) {
    request = &ModifySecurityJSInjectionRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifySecurityJSInjectionRule")
    
    
    return
}

func NewModifySecurityJSInjectionRuleResponse() (response *ModifySecurityJSInjectionRuleResponse) {
    response = &ModifySecurityJSInjectionRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySecurityJSInjectionRule
// This API is used to modify JavaScript injection rules.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ModifySecurityJSInjectionRule(request *ModifySecurityJSInjectionRuleRequest) (response *ModifySecurityJSInjectionRuleResponse, err error) {
    return c.ModifySecurityJSInjectionRuleWithContext(context.Background(), request)
}

// ModifySecurityJSInjectionRule
// This API is used to modify JavaScript injection rules.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ModifySecurityJSInjectionRuleWithContext(ctx context.Context, request *ModifySecurityJSInjectionRuleRequest) (response *ModifySecurityJSInjectionRuleResponse, err error) {
    if request == nil {
        request = NewModifySecurityJSInjectionRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifySecurityJSInjectionRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySecurityJSInjectionRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySecurityJSInjectionRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifySecurityPolicyRequest() (request *ModifySecurityPolicyRequest) {
    request = &ModifySecurityPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifySecurityPolicy")
    
    
    return
}

func NewModifySecurityPolicyResponse() (response *ModifySecurityPolicyResponse) {
    response = &ModifySecurityPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySecurityPolicy
// Modify Web & Bot security configuration.
//
// error code that may be returned:
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  LIMITEXCEEDED_SECURITY = "LimitExceeded.Security"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifySecurityPolicy(request *ModifySecurityPolicyRequest) (response *ModifySecurityPolicyResponse, err error) {
    return c.ModifySecurityPolicyWithContext(context.Background(), request)
}

// ModifySecurityPolicy
// Modify Web & Bot security configuration.
//
// error code that may be returned:
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  LIMITEXCEEDED_SECURITY = "LimitExceeded.Security"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifySecurityPolicyWithContext(ctx context.Context, request *ModifySecurityPolicyRequest) (response *ModifySecurityPolicyResponse, err error) {
    if request == nil {
        request = NewModifySecurityPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifySecurityPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySecurityPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySecurityPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifySharedCNAMERequest() (request *ModifySharedCNAMERequest) {
    request = &ModifySharedCNAMERequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifySharedCNAME")
    
    
    return
}

func NewModifySharedCNAMEResponse() (response *ModifySharedCNAMEResponse) {
    response = &ModifySharedCNAMEResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySharedCNAME
// This API is used to modify a shared CNAME. Currently, it only supports modifying the description and settings of a shared CNAME, and associating an IP SSL domain name with an IP SSL shared CNAME. Once created, a shared CNAME itself cannot be modified. This feature is in beta test.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED_DOMAINMUSTINIPSSLSHAREDCNAMEZONEANDINSHAREDCNAME = "OperationDenied.DomainMustInIPSSLSharedCNAMEZoneAndInSharedCNAME"
//  OPERATIONDENIED_IPSSLALREADYBOUNDANOTHERDOMAIN = "OperationDenied.IPSSLAlreadyBoundAnotherDomain"
//  OPERATIONDENIED_LASTIPSSLOPERATIONNOTCOMPLETE = "OperationDenied.LastIPSSLOperationNotComplete"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifySharedCNAME(request *ModifySharedCNAMERequest) (response *ModifySharedCNAMEResponse, err error) {
    return c.ModifySharedCNAMEWithContext(context.Background(), request)
}

// ModifySharedCNAME
// This API is used to modify a shared CNAME. Currently, it only supports modifying the description and settings of a shared CNAME, and associating an IP SSL domain name with an IP SSL shared CNAME. Once created, a shared CNAME itself cannot be modified. This feature is in beta test.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED_DOMAINMUSTINIPSSLSHAREDCNAMEZONEANDINSHAREDCNAME = "OperationDenied.DomainMustInIPSSLSharedCNAMEZoneAndInSharedCNAME"
//  OPERATIONDENIED_IPSSLALREADYBOUNDANOTHERDOMAIN = "OperationDenied.IPSSLAlreadyBoundAnotherDomain"
//  OPERATIONDENIED_LASTIPSSLOPERATIONNOTCOMPLETE = "OperationDenied.LastIPSSLOperationNotComplete"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifySharedCNAMEWithContext(ctx context.Context, request *ModifySharedCNAMERequest) (response *ModifySharedCNAMEResponse, err error) {
    if request == nil {
        request = NewModifySharedCNAMERequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifySharedCNAME")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySharedCNAME require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySharedCNAMEResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWebSecurityTemplateRequest() (request *ModifyWebSecurityTemplateRequest) {
    request = &ModifyWebSecurityTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyWebSecurityTemplate")
    
    
    return
}

func NewModifyWebSecurityTemplateResponse() (response *ModifyWebSecurityTemplateResponse) {
    response = &ModifyWebSecurityTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyWebSecurityTemplate
// Modifies a security policy configuration template
//
// error code that may be returned:
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
func (c *Client) ModifyWebSecurityTemplate(request *ModifyWebSecurityTemplateRequest) (response *ModifyWebSecurityTemplateResponse, err error) {
    return c.ModifyWebSecurityTemplateWithContext(context.Background(), request)
}

// ModifyWebSecurityTemplate
// Modifies a security policy configuration template
//
// error code that may be returned:
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
func (c *Client) ModifyWebSecurityTemplateWithContext(ctx context.Context, request *ModifyWebSecurityTemplateRequest) (response *ModifyWebSecurityTemplateResponse, err error) {
    if request == nil {
        request = NewModifyWebSecurityTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyWebSecurityTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWebSecurityTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWebSecurityTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyZoneRequest() (request *ModifyZoneRequest) {
    request = &ModifyZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyZone")
    
    
    return
}

func NewModifyZoneResponse() (response *ModifyZoneResponse) {
    response = &ModifyZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyZone
// Modifies site information.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDTOCALLDNSPOD = "FailedOperation.FailedToCallDNSPod"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_LENGTHEXCEEDSLIMIT = "InvalidParameter.LengthExceedsLimit"
//  INVALIDPARAMETER_ZONENAMEISREQUIRED = "InvalidParameter.ZoneNameIsRequired"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ZONESAMEASNAME = "InvalidParameterValue.ZoneSameAsName"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DNSPODUNAUTHORIZEDROLEOPERATION = "OperationDenied.DNSPodUnauthorizedRoleOperation"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_DOMAININSHARECNAMEGROUP = "OperationDenied.DomainInShareCnameGroup"
//  OPERATIONDENIED_DOMAINNUMBERISNOTZERO = "OperationDenied.DomainNumberIsNotZero"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_MULTIPLECNAMEZONE = "OperationDenied.MultipleCnameZone"
//  OPERATIONDENIED_NSNOTALLOWTRAFFICSTRATEGY = "OperationDenied.NSNotAllowTrafficStrategy"
//  OPERATIONDENIED_NODOMAINACCESSZONEONLYALLOWMODIFIEDTOCNAME = "OperationDenied.NoDomainAccessZoneOnlyAllowModifiedToCNAME"
//  OPERATIONDENIED_NODOMAINACCESSZONEONLYSUPPORTMODIFYAREAANDALIASZONENAME = "OperationDenied.NoDomainAccessZoneOnlySupportModifyAreaAndAliasZoneName"
//  OPERATIONDENIED_NODOMAINACCESSZONEONLYSUPPORTMODIFYTYPE = "OperationDenied.NoDomainAccessZoneOnlySupportModifyType"
//  OPERATIONDENIED_PLANNOTSUPPORTMODIFYZONEAREA = "OperationDenied.PlanNotSupportModifyZoneArea"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  OPERATIONDENIED_SWITCHAREACDNPLATFORMREUSE = "OperationDenied.SwitchAreaCdnPlatformReuse"
//  OPERATIONDENIED_SWITCHAREAORIGINPROTECTIONDEPLOYING = "OperationDenied.SwitchAreaOriginProtectionDeploying"
//  OPERATIONDENIED_ZONEHASHOSTSMODIFYCONFLICT = "OperationDenied.ZoneHasHostsModifyConflict"
//  RESOURCEINUSE_CNAME = "ResourceInUse.Cname"
//  RESOURCEINUSE_DNS = "ResourceInUse.Dns"
//  RESOURCEINUSE_GENERICHOST = "ResourceInUse.GenericHost"
//  RESOURCEINUSE_NS = "ResourceInUse.NS"
//  RESOURCEINUSE_OTHERS = "ResourceInUse.Others"
//  RESOURCEINUSE_OTHERSALIASDOMAIN = "ResourceInUse.OthersAliasDomain"
//  RESOURCEINUSE_OTHERSCNAME = "ResourceInUse.OthersCname"
//  RESOURCEINUSE_OTHERSNS = "ResourceInUse.OthersNS"
//  RESOURCEINUSE_SELFANDOTHERSCNAME = "ResourceInUse.SelfAndOthersCname"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DNSPODDOMAINNOTINACCOUNT = "ResourceNotFound.DNSPodDomainNotInAccount"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ModifyZone(request *ModifyZoneRequest) (response *ModifyZoneResponse, err error) {
    return c.ModifyZoneWithContext(context.Background(), request)
}

// ModifyZone
// Modifies site information.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDTOCALLDNSPOD = "FailedOperation.FailedToCallDNSPod"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_LENGTHEXCEEDSLIMIT = "InvalidParameter.LengthExceedsLimit"
//  INVALIDPARAMETER_ZONENAMEISREQUIRED = "InvalidParameter.ZoneNameIsRequired"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ZONESAMEASNAME = "InvalidParameterValue.ZoneSameAsName"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DNSPODUNAUTHORIZEDROLEOPERATION = "OperationDenied.DNSPodUnauthorizedRoleOperation"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_DOMAININSHARECNAMEGROUP = "OperationDenied.DomainInShareCnameGroup"
//  OPERATIONDENIED_DOMAINNUMBERISNOTZERO = "OperationDenied.DomainNumberIsNotZero"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_MULTIPLECNAMEZONE = "OperationDenied.MultipleCnameZone"
//  OPERATIONDENIED_NSNOTALLOWTRAFFICSTRATEGY = "OperationDenied.NSNotAllowTrafficStrategy"
//  OPERATIONDENIED_NODOMAINACCESSZONEONLYALLOWMODIFIEDTOCNAME = "OperationDenied.NoDomainAccessZoneOnlyAllowModifiedToCNAME"
//  OPERATIONDENIED_NODOMAINACCESSZONEONLYSUPPORTMODIFYAREAANDALIASZONENAME = "OperationDenied.NoDomainAccessZoneOnlySupportModifyAreaAndAliasZoneName"
//  OPERATIONDENIED_NODOMAINACCESSZONEONLYSUPPORTMODIFYTYPE = "OperationDenied.NoDomainAccessZoneOnlySupportModifyType"
//  OPERATIONDENIED_PLANNOTSUPPORTMODIFYZONEAREA = "OperationDenied.PlanNotSupportModifyZoneArea"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  OPERATIONDENIED_SWITCHAREACDNPLATFORMREUSE = "OperationDenied.SwitchAreaCdnPlatformReuse"
//  OPERATIONDENIED_SWITCHAREAORIGINPROTECTIONDEPLOYING = "OperationDenied.SwitchAreaOriginProtectionDeploying"
//  OPERATIONDENIED_ZONEHASHOSTSMODIFYCONFLICT = "OperationDenied.ZoneHasHostsModifyConflict"
//  RESOURCEINUSE_CNAME = "ResourceInUse.Cname"
//  RESOURCEINUSE_DNS = "ResourceInUse.Dns"
//  RESOURCEINUSE_GENERICHOST = "ResourceInUse.GenericHost"
//  RESOURCEINUSE_NS = "ResourceInUse.NS"
//  RESOURCEINUSE_OTHERS = "ResourceInUse.Others"
//  RESOURCEINUSE_OTHERSALIASDOMAIN = "ResourceInUse.OthersAliasDomain"
//  RESOURCEINUSE_OTHERSCNAME = "ResourceInUse.OthersCname"
//  RESOURCEINUSE_OTHERSNS = "ResourceInUse.OthersNS"
//  RESOURCEINUSE_SELFANDOTHERSCNAME = "ResourceInUse.SelfAndOthersCname"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DNSPODDOMAINNOTINACCOUNT = "ResourceNotFound.DNSPodDomainNotInAccount"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ModifyZoneWithContext(ctx context.Context, request *ModifyZoneRequest) (response *ModifyZoneResponse, err error) {
    if request == nil {
        request = NewModifyZoneRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyZone")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyZoneResponse()
    err = c.Send(request, response)
    return
}

func NewModifyZoneCustomVariablesRequest() (request *ModifyZoneCustomVariablesRequest) {
    request = &ModifyZoneCustomVariablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyZoneCustomVariables")
    
    
    return
}

func NewModifyZoneCustomVariablesResponse() (response *ModifyZoneCustomVariablesResponse) {
    response = &ModifyZoneCustomVariablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyZoneCustomVariables
// This API is used to manage site-level custom variables, supporting the creation, modification, and deletion of variable definitions and variable operations. The submitted list will override the corresponding existing configuration. To clear the configuration, pass in an empty list. Configuration items not passed in remain unchanged. The default values for variable definitions and variable operations are both empty.
//
// 
//
// For rule-level custom variable related operations, call the CreateL7AccRules or ModifyL7AccRule API.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWERROR = "InternalError.UnknowError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_CACHEKEYQUERYSTRINGREQUIRESFULLURLCACHEOFF = "InvalidParameter.CacheKeyQueryStringRequiresFullUrlCacheOff"
//  INVALIDPARAMETER_CACHEKEYQUERYSTRINGTOOMANYVALUE = "InvalidParameter.CacheKeyQueryStringTooManyValue"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_CLIENTIPCOUNTRYCONFLICTSWITHIPV6 = "InvalidParameter.ClientIpCountryConflictsWithIpv6"
//  INVALIDPARAMETER_GRPCREQUIREHTTP2 = "InvalidParameter.GrpcRequireHttp2"
//  INVALIDPARAMETER_INVALIDAWSPRIVATEACCESS = "InvalidParameter.InvalidAwsPrivateAccess"
//  INVALIDPARAMETER_INVALIDCACHECONFIGFOLLOWORIGIN = "InvalidParameter.InvalidCacheConfigFollowOrigin"
//  INVALIDPARAMETER_INVALIDCACHEKEYQUERYSTRINGVALUE = "InvalidParameter.InvalidCacheKeyQueryStringValue"
//  INVALIDPARAMETER_INVALIDCACHEONLYONSWITCH = "InvalidParameter.InvalidCacheOnlyOnSwitch"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPCOUNTRYHEADERNAME = "InvalidParameter.InvalidClientIpCountryHeaderName"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINE = "InvalidParameter.InvalidDynamicRoutine"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINEBILLING = "InvalidParameter.InvalidDynamicRoutineBilling"
//  INVALIDPARAMETER_INVALIDFORCEREDIRECTTYPE = "InvalidParameter.InvalidForceRedirectType"
//  INVALIDPARAMETER_INVALIDHTTPS = "InvalidParameter.InvalidHttps"
//  INVALIDPARAMETER_INVALIDHTTPSCERTINFO = "InvalidParameter.InvalidHttpsCertInfo"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDHTTPSTLSVERSION = "InvalidParameter.InvalidHttpsTlsVersion"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGETIME = "InvalidParameter.InvalidMaxAgeTime"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDORIGINTYPE = "InvalidParameter.InvalidOriginType"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPOSTMAXSIZEBILLING = "InvalidParameter.InvalidPostMaxSizeBilling"
//  INVALIDPARAMETER_INVALIDPOSTSIZEVALUE = "InvalidParameter.InvalidPostSizeValue"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDRESOURCEIDBILLING = "InvalidParameter.InvalidResourceIdBilling"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGCLIENTIP = "InvalidParameter.InvalidStandardDebugClientIp"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGEMPTYLIST = "InvalidParameter.InvalidStandardDebugEmptyList"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGEXPIRETIMELIMIT = "InvalidParameter.InvalidStandardDebugExpireTimeLimit"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_MULTIPLYLAYERNOTSUPPORTSMARTROUTING = "InvalidParameter.MultiplyLayerNotSupportSmartRouting"
//  INVALIDPARAMETER_OCDIRECTORIGINREQUIRESSMARTROUTING = "InvalidParameter.OCDirectOriginRequiresSmartRouting"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_SETTINGINVALIDPARAM = "InvalidParameter.SettingInvalidParam"
//  INVALIDPARAMETER_SHIELDNOTSUPPORTZONECACHEPREFRESH = "InvalidParameter.ShieldNotSupportZoneCachePrefresh"
//  INVALIDPARAMETER_SHIELDNOTSUPPORTZONESMARTROUTING = "InvalidParameter.ShieldNotSupportZoneSmartRouting"
//  INVALIDPARAMETER_ZONEISGRAYPUBLISHING = "InvalidParameter.ZoneIsGrayPublishing"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  INVALIDPARAMETERVALUE_FORMATMISMATCH = "InvalidParameterValue.FormatMismatch"
//  INVALIDPARAMETERVALUE_GENERALMISMATCH = "InvalidParameterValue.GeneralMismatch"
//  INVALIDPARAMETERVALUE_INCLUDEINVALIDVALUE = "InvalidParameterValue.IncludeInvalidValue"
//  INVALIDPARAMETERVALUE_MISSINGNECESSARYPARAM = "InvalidParameterValue.MissingNecessaryParam"
//  INVALIDPARAMETERVALUE_NOTINENUMERATION = "InvalidParameterValue.NotInEnumeration"
//  INVALIDPARAMETERVALUE_NOTWITHINRANGE = "InvalidParameterValue.NotWithinRange"
//  INVALIDPARAMETERVALUE_REGEXMISMATCH = "InvalidParameterValue.RegExMismatch"
//  INVALIDPARAMETERVALUE_UNRECOGNIZABLEVALUE = "InvalidParameterValue.UnrecognizableValue"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDDISABLE = "OperationDenied.AccelerateMainlandDisable"
//  OPERATIONDENIED_ACCELERATEMAINLANDIPV6CONFLICT = "OperationDenied.AccelerateMainlandIpv6Conflict"
//  OPERATIONDENIED_ACCELERATEMAINLANDMULTIPLYLAYERCONFLICT = "OperationDenied.AccelerateMainlandMultiplyLayerConflict"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_NOTALLOWCLOSEACCELERATEMAINLANDWHENDOMAINBOUNDEXCLUSIVEIP = "OperationDenied.NotAllowCloseAccelerateMainlandWhenDomainBoundExclusiveIP"
//  OPERATIONDENIED_SHAREDCNAMEUNSUPPORTEDACCELERATEMAINLAND = "OperationDenied.SharedCNAMEUnsupportedAccelerateMainland"
//  OPERATIONDENIED_SHAREDCNAMEUNSUPPORTEDIPV6 = "OperationDenied.SharedCNAMEUnsupportedIPv6"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_POSTMAXSIZEQUOTANOTFOUND = "ResourceNotFound.PostMaxSizeQuotaNotFound"
//  RESOURCEUNAVAILABLE_CERTNOTFOUND = "ResourceUnavailable.CertNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ModifyZoneCustomVariables(request *ModifyZoneCustomVariablesRequest) (response *ModifyZoneCustomVariablesResponse, err error) {
    return c.ModifyZoneCustomVariablesWithContext(context.Background(), request)
}

// ModifyZoneCustomVariables
// This API is used to manage site-level custom variables, supporting the creation, modification, and deletion of variable definitions and variable operations. The submitted list will override the corresponding existing configuration. To clear the configuration, pass in an empty list. Configuration items not passed in remain unchanged. The default values for variable definitions and variable operations are both empty.
//
// 
//
// For rule-level custom variable related operations, call the CreateL7AccRules or ModifyL7AccRule API.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWERROR = "InternalError.UnknowError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_CACHEKEYQUERYSTRINGREQUIRESFULLURLCACHEOFF = "InvalidParameter.CacheKeyQueryStringRequiresFullUrlCacheOff"
//  INVALIDPARAMETER_CACHEKEYQUERYSTRINGTOOMANYVALUE = "InvalidParameter.CacheKeyQueryStringTooManyValue"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_CLIENTIPCOUNTRYCONFLICTSWITHIPV6 = "InvalidParameter.ClientIpCountryConflictsWithIpv6"
//  INVALIDPARAMETER_GRPCREQUIREHTTP2 = "InvalidParameter.GrpcRequireHttp2"
//  INVALIDPARAMETER_INVALIDAWSPRIVATEACCESS = "InvalidParameter.InvalidAwsPrivateAccess"
//  INVALIDPARAMETER_INVALIDCACHECONFIGFOLLOWORIGIN = "InvalidParameter.InvalidCacheConfigFollowOrigin"
//  INVALIDPARAMETER_INVALIDCACHEKEYQUERYSTRINGVALUE = "InvalidParameter.InvalidCacheKeyQueryStringValue"
//  INVALIDPARAMETER_INVALIDCACHEONLYONSWITCH = "InvalidParameter.InvalidCacheOnlyOnSwitch"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPCOUNTRYHEADERNAME = "InvalidParameter.InvalidClientIpCountryHeaderName"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINE = "InvalidParameter.InvalidDynamicRoutine"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINEBILLING = "InvalidParameter.InvalidDynamicRoutineBilling"
//  INVALIDPARAMETER_INVALIDFORCEREDIRECTTYPE = "InvalidParameter.InvalidForceRedirectType"
//  INVALIDPARAMETER_INVALIDHTTPS = "InvalidParameter.InvalidHttps"
//  INVALIDPARAMETER_INVALIDHTTPSCERTINFO = "InvalidParameter.InvalidHttpsCertInfo"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDHTTPSTLSVERSION = "InvalidParameter.InvalidHttpsTlsVersion"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGETIME = "InvalidParameter.InvalidMaxAgeTime"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDORIGINTYPE = "InvalidParameter.InvalidOriginType"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPOSTMAXSIZEBILLING = "InvalidParameter.InvalidPostMaxSizeBilling"
//  INVALIDPARAMETER_INVALIDPOSTSIZEVALUE = "InvalidParameter.InvalidPostSizeValue"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDRESOURCEIDBILLING = "InvalidParameter.InvalidResourceIdBilling"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGCLIENTIP = "InvalidParameter.InvalidStandardDebugClientIp"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGEMPTYLIST = "InvalidParameter.InvalidStandardDebugEmptyList"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGEXPIRETIMELIMIT = "InvalidParameter.InvalidStandardDebugExpireTimeLimit"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_MULTIPLYLAYERNOTSUPPORTSMARTROUTING = "InvalidParameter.MultiplyLayerNotSupportSmartRouting"
//  INVALIDPARAMETER_OCDIRECTORIGINREQUIRESSMARTROUTING = "InvalidParameter.OCDirectOriginRequiresSmartRouting"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_SETTINGINVALIDPARAM = "InvalidParameter.SettingInvalidParam"
//  INVALIDPARAMETER_SHIELDNOTSUPPORTZONECACHEPREFRESH = "InvalidParameter.ShieldNotSupportZoneCachePrefresh"
//  INVALIDPARAMETER_SHIELDNOTSUPPORTZONESMARTROUTING = "InvalidParameter.ShieldNotSupportZoneSmartRouting"
//  INVALIDPARAMETER_ZONEISGRAYPUBLISHING = "InvalidParameter.ZoneIsGrayPublishing"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  INVALIDPARAMETERVALUE_FORMATMISMATCH = "InvalidParameterValue.FormatMismatch"
//  INVALIDPARAMETERVALUE_GENERALMISMATCH = "InvalidParameterValue.GeneralMismatch"
//  INVALIDPARAMETERVALUE_INCLUDEINVALIDVALUE = "InvalidParameterValue.IncludeInvalidValue"
//  INVALIDPARAMETERVALUE_MISSINGNECESSARYPARAM = "InvalidParameterValue.MissingNecessaryParam"
//  INVALIDPARAMETERVALUE_NOTINENUMERATION = "InvalidParameterValue.NotInEnumeration"
//  INVALIDPARAMETERVALUE_NOTWITHINRANGE = "InvalidParameterValue.NotWithinRange"
//  INVALIDPARAMETERVALUE_REGEXMISMATCH = "InvalidParameterValue.RegExMismatch"
//  INVALIDPARAMETERVALUE_UNRECOGNIZABLEVALUE = "InvalidParameterValue.UnrecognizableValue"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDDISABLE = "OperationDenied.AccelerateMainlandDisable"
//  OPERATIONDENIED_ACCELERATEMAINLANDIPV6CONFLICT = "OperationDenied.AccelerateMainlandIpv6Conflict"
//  OPERATIONDENIED_ACCELERATEMAINLANDMULTIPLYLAYERCONFLICT = "OperationDenied.AccelerateMainlandMultiplyLayerConflict"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_NOTALLOWCLOSEACCELERATEMAINLANDWHENDOMAINBOUNDEXCLUSIVEIP = "OperationDenied.NotAllowCloseAccelerateMainlandWhenDomainBoundExclusiveIP"
//  OPERATIONDENIED_SHAREDCNAMEUNSUPPORTEDACCELERATEMAINLAND = "OperationDenied.SharedCNAMEUnsupportedAccelerateMainland"
//  OPERATIONDENIED_SHAREDCNAMEUNSUPPORTEDIPV6 = "OperationDenied.SharedCNAMEUnsupportedIPv6"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_POSTMAXSIZEQUOTANOTFOUND = "ResourceNotFound.PostMaxSizeQuotaNotFound"
//  RESOURCEUNAVAILABLE_CERTNOTFOUND = "ResourceUnavailable.CertNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ModifyZoneCustomVariablesWithContext(ctx context.Context, request *ModifyZoneCustomVariablesRequest) (response *ModifyZoneCustomVariablesResponse, err error) {
    if request == nil {
        request = NewModifyZoneCustomVariablesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyZoneCustomVariables")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyZoneCustomVariables require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyZoneCustomVariablesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyZoneSettingRequest() (request *ModifyZoneSettingRequest) {
    request = &ModifyZoneSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyZoneSetting")
    
    
    return
}

func NewModifyZoneSettingResponse() (response *ModifyZoneSettingResponse) {
    response = &ModifyZoneSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyZoneSetting
// This API is on an earlier version. EdgeOne has completed a comprehensive upgrade of rule engine APIs. For details, see [ModifyL7AccSetting](https://www.tencentcloud.com/document/product/1552/115817?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWERROR = "InternalError.UnknowError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_CACHEKEYQUERYSTRINGREQUIRESFULLURLCACHEOFF = "InvalidParameter.CacheKeyQueryStringRequiresFullUrlCacheOff"
//  INVALIDPARAMETER_CACHEKEYQUERYSTRINGTOOMANYVALUE = "InvalidParameter.CacheKeyQueryStringTooManyValue"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_CLIENTIPCOUNTRYCONFLICTSWITHIPV6 = "InvalidParameter.ClientIpCountryConflictsWithIpv6"
//  INVALIDPARAMETER_GRPCREQUIREHTTP2 = "InvalidParameter.GrpcRequireHttp2"
//  INVALIDPARAMETER_INVALIDAWSPRIVATEACCESS = "InvalidParameter.InvalidAwsPrivateAccess"
//  INVALIDPARAMETER_INVALIDCACHECONFIGFOLLOWORIGIN = "InvalidParameter.InvalidCacheConfigFollowOrigin"
//  INVALIDPARAMETER_INVALIDCACHEKEYQUERYSTRINGVALUE = "InvalidParameter.InvalidCacheKeyQueryStringValue"
//  INVALIDPARAMETER_INVALIDCACHEONLYONSWITCH = "InvalidParameter.InvalidCacheOnlyOnSwitch"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPCOUNTRYHEADERNAME = "InvalidParameter.InvalidClientIpCountryHeaderName"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINE = "InvalidParameter.InvalidDynamicRoutine"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINEBILLING = "InvalidParameter.InvalidDynamicRoutineBilling"
//  INVALIDPARAMETER_INVALIDFORCEREDIRECTTYPE = "InvalidParameter.InvalidForceRedirectType"
//  INVALIDPARAMETER_INVALIDHTTPS = "InvalidParameter.InvalidHttps"
//  INVALIDPARAMETER_INVALIDHTTPSCERTINFO = "InvalidParameter.InvalidHttpsCertInfo"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDHTTPSTLSVERSION = "InvalidParameter.InvalidHttpsTlsVersion"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGETIME = "InvalidParameter.InvalidMaxAgeTime"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDORIGINTYPE = "InvalidParameter.InvalidOriginType"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPOSTMAXSIZEBILLING = "InvalidParameter.InvalidPostMaxSizeBilling"
//  INVALIDPARAMETER_INVALIDPOSTSIZEVALUE = "InvalidParameter.InvalidPostSizeValue"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDRESOURCEIDBILLING = "InvalidParameter.InvalidResourceIdBilling"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGCLIENTIP = "InvalidParameter.InvalidStandardDebugClientIp"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGEMPTYLIST = "InvalidParameter.InvalidStandardDebugEmptyList"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGEXPIRETIMELIMIT = "InvalidParameter.InvalidStandardDebugExpireTimeLimit"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_MULTIPLYLAYERNOTSUPPORTSMARTROUTING = "InvalidParameter.MultiplyLayerNotSupportSmartRouting"
//  INVALIDPARAMETER_OCDIRECTORIGINREQUIRESSMARTROUTING = "InvalidParameter.OCDirectOriginRequiresSmartRouting"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_SETTINGINVALIDPARAM = "InvalidParameter.SettingInvalidParam"
//  INVALIDPARAMETER_ZONEISGRAYPUBLISHING = "InvalidParameter.ZoneIsGrayPublishing"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  INVALIDPARAMETERVALUE_FORMATMISMATCH = "InvalidParameterValue.FormatMismatch"
//  INVALIDPARAMETERVALUE_GENERALMISMATCH = "InvalidParameterValue.GeneralMismatch"
//  INVALIDPARAMETERVALUE_INCLUDEINVALIDVALUE = "InvalidParameterValue.IncludeInvalidValue"
//  INVALIDPARAMETERVALUE_MISSINGNECESSARYPARAM = "InvalidParameterValue.MissingNecessaryParam"
//  INVALIDPARAMETERVALUE_NOTINENUMERATION = "InvalidParameterValue.NotInEnumeration"
//  INVALIDPARAMETERVALUE_NOTWITHINRANGE = "InvalidParameterValue.NotWithinRange"
//  INVALIDPARAMETERVALUE_REGEXMISMATCH = "InvalidParameterValue.RegExMismatch"
//  INVALIDPARAMETERVALUE_UNRECOGNIZABLEVALUE = "InvalidParameterValue.UnrecognizableValue"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDDISABLE = "OperationDenied.AccelerateMainlandDisable"
//  OPERATIONDENIED_ACCELERATEMAINLANDIPV6CONFLICT = "OperationDenied.AccelerateMainlandIpv6Conflict"
//  OPERATIONDENIED_ACCELERATEMAINLANDMULTIPLYLAYERCONFLICT = "OperationDenied.AccelerateMainlandMultiplyLayerConflict"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_SHAREDCNAMEUNSUPPORTEDACCELERATEMAINLAND = "OperationDenied.SharedCNAMEUnsupportedAccelerateMainland"
//  OPERATIONDENIED_SHAREDCNAMEUNSUPPORTEDIPV6 = "OperationDenied.SharedCNAMEUnsupportedIPv6"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_POSTMAXSIZEQUOTANOTFOUND = "ResourceNotFound.PostMaxSizeQuotaNotFound"
//  RESOURCEUNAVAILABLE_CERTNOTFOUND = "ResourceUnavailable.CertNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ModifyZoneSetting(request *ModifyZoneSettingRequest) (response *ModifyZoneSettingResponse, err error) {
    return c.ModifyZoneSettingWithContext(context.Background(), request)
}

// ModifyZoneSetting
// This API is on an earlier version. EdgeOne has completed a comprehensive upgrade of rule engine APIs. For details, see [ModifyL7AccSetting](https://www.tencentcloud.com/document/product/1552/115817?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWERROR = "InternalError.UnknowError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_CACHEKEYQUERYSTRINGREQUIRESFULLURLCACHEOFF = "InvalidParameter.CacheKeyQueryStringRequiresFullUrlCacheOff"
//  INVALIDPARAMETER_CACHEKEYQUERYSTRINGTOOMANYVALUE = "InvalidParameter.CacheKeyQueryStringTooManyValue"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_CLIENTIPCOUNTRYCONFLICTSWITHIPV6 = "InvalidParameter.ClientIpCountryConflictsWithIpv6"
//  INVALIDPARAMETER_GRPCREQUIREHTTP2 = "InvalidParameter.GrpcRequireHttp2"
//  INVALIDPARAMETER_INVALIDAWSPRIVATEACCESS = "InvalidParameter.InvalidAwsPrivateAccess"
//  INVALIDPARAMETER_INVALIDCACHECONFIGFOLLOWORIGIN = "InvalidParameter.InvalidCacheConfigFollowOrigin"
//  INVALIDPARAMETER_INVALIDCACHEKEYQUERYSTRINGVALUE = "InvalidParameter.InvalidCacheKeyQueryStringValue"
//  INVALIDPARAMETER_INVALIDCACHEONLYONSWITCH = "InvalidParameter.InvalidCacheOnlyOnSwitch"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPCOUNTRYHEADERNAME = "InvalidParameter.InvalidClientIpCountryHeaderName"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINE = "InvalidParameter.InvalidDynamicRoutine"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINEBILLING = "InvalidParameter.InvalidDynamicRoutineBilling"
//  INVALIDPARAMETER_INVALIDFORCEREDIRECTTYPE = "InvalidParameter.InvalidForceRedirectType"
//  INVALIDPARAMETER_INVALIDHTTPS = "InvalidParameter.InvalidHttps"
//  INVALIDPARAMETER_INVALIDHTTPSCERTINFO = "InvalidParameter.InvalidHttpsCertInfo"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDHTTPSTLSVERSION = "InvalidParameter.InvalidHttpsTlsVersion"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGETIME = "InvalidParameter.InvalidMaxAgeTime"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDORIGINTYPE = "InvalidParameter.InvalidOriginType"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPOSTMAXSIZEBILLING = "InvalidParameter.InvalidPostMaxSizeBilling"
//  INVALIDPARAMETER_INVALIDPOSTSIZEVALUE = "InvalidParameter.InvalidPostSizeValue"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDRESOURCEIDBILLING = "InvalidParameter.InvalidResourceIdBilling"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGCLIENTIP = "InvalidParameter.InvalidStandardDebugClientIp"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGEMPTYLIST = "InvalidParameter.InvalidStandardDebugEmptyList"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGEXPIRETIMELIMIT = "InvalidParameter.InvalidStandardDebugExpireTimeLimit"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_MULTIPLYLAYERNOTSUPPORTSMARTROUTING = "InvalidParameter.MultiplyLayerNotSupportSmartRouting"
//  INVALIDPARAMETER_OCDIRECTORIGINREQUIRESSMARTROUTING = "InvalidParameter.OCDirectOriginRequiresSmartRouting"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_SETTINGINVALIDPARAM = "InvalidParameter.SettingInvalidParam"
//  INVALIDPARAMETER_ZONEISGRAYPUBLISHING = "InvalidParameter.ZoneIsGrayPublishing"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  INVALIDPARAMETERVALUE_FORMATMISMATCH = "InvalidParameterValue.FormatMismatch"
//  INVALIDPARAMETERVALUE_GENERALMISMATCH = "InvalidParameterValue.GeneralMismatch"
//  INVALIDPARAMETERVALUE_INCLUDEINVALIDVALUE = "InvalidParameterValue.IncludeInvalidValue"
//  INVALIDPARAMETERVALUE_MISSINGNECESSARYPARAM = "InvalidParameterValue.MissingNecessaryParam"
//  INVALIDPARAMETERVALUE_NOTINENUMERATION = "InvalidParameterValue.NotInEnumeration"
//  INVALIDPARAMETERVALUE_NOTWITHINRANGE = "InvalidParameterValue.NotWithinRange"
//  INVALIDPARAMETERVALUE_REGEXMISMATCH = "InvalidParameterValue.RegExMismatch"
//  INVALIDPARAMETERVALUE_UNRECOGNIZABLEVALUE = "InvalidParameterValue.UnrecognizableValue"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDDISABLE = "OperationDenied.AccelerateMainlandDisable"
//  OPERATIONDENIED_ACCELERATEMAINLANDIPV6CONFLICT = "OperationDenied.AccelerateMainlandIpv6Conflict"
//  OPERATIONDENIED_ACCELERATEMAINLANDMULTIPLYLAYERCONFLICT = "OperationDenied.AccelerateMainlandMultiplyLayerConflict"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_SHAREDCNAMEUNSUPPORTEDACCELERATEMAINLAND = "OperationDenied.SharedCNAMEUnsupportedAccelerateMainland"
//  OPERATIONDENIED_SHAREDCNAMEUNSUPPORTEDIPV6 = "OperationDenied.SharedCNAMEUnsupportedIPv6"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_POSTMAXSIZEQUOTANOTFOUND = "ResourceNotFound.PostMaxSizeQuotaNotFound"
//  RESOURCEUNAVAILABLE_CERTNOTFOUND = "ResourceUnavailable.CertNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ModifyZoneSettingWithContext(ctx context.Context, request *ModifyZoneSettingRequest) (response *ModifyZoneSettingResponse, err error) {
    if request == nil {
        request = NewModifyZoneSettingRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyZoneSetting")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyZoneSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyZoneSettingResponse()
    err = c.Send(request, response)
    return
}

func NewModifyZoneStatusRequest() (request *ModifyZoneStatusRequest) {
    request = &ModifyZoneStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyZoneStatus")
    
    
    return
}

func NewModifyZoneStatusResponse() (response *ModifyZoneStatusResponse) {
    response = &ModifyZoneStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyZoneStatus
// Used to enable or disable a site.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_L4PROXYINPROGRESSSTATUS = "OperationDenied.L4ProxyInProgressStatus"
//  OPERATIONDENIED_L4PROXYINSTOPPINGSTATUS = "OperationDenied.L4ProxyInStoppingStatus"
//  OPERATIONDENIED_L7HOSTINPROCESSSTATUS = "OperationDenied.L7HostInProcessStatus"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  OPERATIONDENIED_VERSIONCONTROLISGRAYING = "OperationDenied.VersionControlIsGraying"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) ModifyZoneStatus(request *ModifyZoneStatusRequest) (response *ModifyZoneStatusResponse, err error) {
    return c.ModifyZoneStatusWithContext(context.Background(), request)
}

// ModifyZoneStatus
// Used to enable or disable a site.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_L4PROXYINPROGRESSSTATUS = "OperationDenied.L4ProxyInProgressStatus"
//  OPERATIONDENIED_L4PROXYINSTOPPINGSTATUS = "OperationDenied.L4ProxyInStoppingStatus"
//  OPERATIONDENIED_L7HOSTINPROCESSSTATUS = "OperationDenied.L7HostInProcessStatus"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  OPERATIONDENIED_VERSIONCONTROLISGRAYING = "OperationDenied.VersionControlIsGraying"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) ModifyZoneStatusWithContext(ctx context.Context, request *ModifyZoneStatusRequest) (response *ModifyZoneStatusResponse, err error) {
    if request == nil {
        request = NewModifyZoneStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyZoneStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyZoneStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyZoneStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyZoneWorkModeRequest() (request *ModifyZoneWorkModeRequest) {
    request = &ModifyZoneWorkModeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyZoneWorkMode")
    
    
    return
}

func NewModifyZoneWorkModeResponse() (response *ModifyZoneWorkModeResponse) {
    response = &ModifyZoneWorkModeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyZoneWorkMode
// This API is used to modify the working mode of each configuration module under a site. Each configuration module of a site can enable version management mode or immediate effect mode by configuration group dimension. For details, see Version Management (https://www.tencentcloud.com/document/product/1552/113690?from_cn_redirect=1).
//
// error code that may be returned:
//  INVALIDPARAMETER_WORKMODEISSAME = "InvalidParameter.WorkModeIsSame"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  INVALIDPARAMETERVALUE_OWNERSHIPVERIFICATIONNOTPASSED = "InvalidParameterValue.OwnershipVerificationNotPassed"
//  OPERATIONDENIED_ALIASDOMAINNOTSUPPORT = "OperationDenied.AliasDomainNotSupport"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_DOMAINSTATUSUNSTABLE = "OperationDenied.DomainStatusUnstable"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_NOTINVERSIONCONTROLWHITELIST = "OperationDenied.NotInVersionControlWhiteList"
//  OPERATIONDENIED_OWNERSHIPVERIFICATIONNOTPASSED = "OperationDenied.OwnershipVerificationNotPassed"
//  OPERATIONDENIED_PLANTYPEISWRONG = "OperationDenied.PlanTypeIsWrong"
//  OPERATIONDENIED_VERSIONCONTROLLOCKED = "OperationDenied.VersionControlLocked"
//  OPERATIONDENIED_VERSIONCONTROLNEEDNSSWITCHED = "OperationDenied.VersionControlNeedNSSwitched"
//  RESOURCENOTFOUND_ZONENOTFOUND = "ResourceNotFound.ZoneNotFound"
//  RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"
func (c *Client) ModifyZoneWorkMode(request *ModifyZoneWorkModeRequest) (response *ModifyZoneWorkModeResponse, err error) {
    return c.ModifyZoneWorkModeWithContext(context.Background(), request)
}

// ModifyZoneWorkMode
// This API is used to modify the working mode of each configuration module under a site. Each configuration module of a site can enable version management mode or immediate effect mode by configuration group dimension. For details, see Version Management (https://www.tencentcloud.com/document/product/1552/113690?from_cn_redirect=1).
//
// error code that may be returned:
//  INVALIDPARAMETER_WORKMODEISSAME = "InvalidParameter.WorkModeIsSame"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  INVALIDPARAMETERVALUE_OWNERSHIPVERIFICATIONNOTPASSED = "InvalidParameterValue.OwnershipVerificationNotPassed"
//  OPERATIONDENIED_ALIASDOMAINNOTSUPPORT = "OperationDenied.AliasDomainNotSupport"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_DOMAINSTATUSUNSTABLE = "OperationDenied.DomainStatusUnstable"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_NOTINVERSIONCONTROLWHITELIST = "OperationDenied.NotInVersionControlWhiteList"
//  OPERATIONDENIED_OWNERSHIPVERIFICATIONNOTPASSED = "OperationDenied.OwnershipVerificationNotPassed"
//  OPERATIONDENIED_PLANTYPEISWRONG = "OperationDenied.PlanTypeIsWrong"
//  OPERATIONDENIED_VERSIONCONTROLLOCKED = "OperationDenied.VersionControlLocked"
//  OPERATIONDENIED_VERSIONCONTROLNEEDNSSWITCHED = "OperationDenied.VersionControlNeedNSSwitched"
//  RESOURCENOTFOUND_ZONENOTFOUND = "ResourceNotFound.ZoneNotFound"
//  RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"
func (c *Client) ModifyZoneWorkModeWithContext(ctx context.Context, request *ModifyZoneWorkModeRequest) (response *ModifyZoneWorkModeResponse, err error) {
    if request == nil {
        request = NewModifyZoneWorkModeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "ModifyZoneWorkMode")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyZoneWorkMode require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyZoneWorkModeResponse()
    err = c.Send(request, response)
    return
}

func NewOperateInferenceDomainRequest() (request *OperateInferenceDomainRequest) {
    request = &OperateInferenceDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "OperateInferenceDomain")
    
    
    return
}

func NewOperateInferenceDomainResponse() (response *OperateInferenceDomainResponse) {
    response = &OperateInferenceDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OperateInferenceDomain
// Operate an inference service domain name. You can stop, enable, and delete it. Resources cannot be recovered after deletion.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_ZONENOTFOUND = "InvalidParameterValue.ZoneNotFound"
//  RESOURCENOTFOUND_INFERENCEDOMAIN = "ResourceNotFound.InferenceDomain"
//  RESOURCENOTFOUND_INFERENCESERVICE = "ResourceNotFound.InferenceService"
func (c *Client) OperateInferenceDomain(request *OperateInferenceDomainRequest) (response *OperateInferenceDomainResponse, err error) {
    return c.OperateInferenceDomainWithContext(context.Background(), request)
}

// OperateInferenceDomain
// Operate an inference service domain name. You can stop, enable, and delete it. Resources cannot be recovered after deletion.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_ZONENOTFOUND = "InvalidParameterValue.ZoneNotFound"
//  RESOURCENOTFOUND_INFERENCEDOMAIN = "ResourceNotFound.InferenceDomain"
//  RESOURCENOTFOUND_INFERENCESERVICE = "ResourceNotFound.InferenceService"
func (c *Client) OperateInferenceDomainWithContext(ctx context.Context, request *OperateInferenceDomainRequest) (response *OperateInferenceDomainResponse, err error) {
    if request == nil {
        request = NewOperateInferenceDomainRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "OperateInferenceDomain")
    
    if c.GetCredential() == nil {
        return nil, errors.New("OperateInferenceDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewOperateInferenceDomainResponse()
    err = c.Send(request, response)
    return
}

func NewOperateInferenceServiceRequest() (request *OperateInferenceServiceRequest) {
    request = &OperateInferenceServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "OperateInferenceService")
    
    
    return
}

func NewOperateInferenceServiceResponse() (response *OperateInferenceServiceResponse) {
    response = &OperateInferenceServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OperateInferenceService
// Operate an inference service. You can stop, start, and delete it. Resources cannot be recovered after deletion.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_ZONENOTFOUND = "InvalidParameterValue.ZoneNotFound"
//  RESOURCENOTFOUND_INFERENCESERVICE = "ResourceNotFound.InferenceService"
func (c *Client) OperateInferenceService(request *OperateInferenceServiceRequest) (response *OperateInferenceServiceResponse, err error) {
    return c.OperateInferenceServiceWithContext(context.Background(), request)
}

// OperateInferenceService
// Operate an inference service. You can stop, start, and delete it. Resources cannot be recovered after deletion.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_ZONENOTFOUND = "InvalidParameterValue.ZoneNotFound"
//  RESOURCENOTFOUND_INFERENCESERVICE = "ResourceNotFound.InferenceService"
func (c *Client) OperateInferenceServiceWithContext(ctx context.Context, request *OperateInferenceServiceRequest) (response *OperateInferenceServiceResponse, err error) {
    if request == nil {
        request = NewOperateInferenceServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "OperateInferenceService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("OperateInferenceService require credential")
    }

    request.SetContext(ctx)
    
    response = NewOperateInferenceServiceResponse()
    err = c.Send(request, response)
    return
}

func NewRefreshMultiPathGatewaySecretKeyRequest() (request *RefreshMultiPathGatewaySecretKeyRequest) {
    request = &RefreshMultiPathGatewaySecretKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "RefreshMultiPathGatewaySecretKey")
    
    
    return
}

func NewRefreshMultiPathGatewaySecretKeyResponse() (response *RefreshMultiPathGatewaySecretKeyResponse) {
    response = &RefreshMultiPathGatewaySecretKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RefreshMultiPathGatewaySecretKey
// This API is used to refresh the key of a multi-channel security acceleration gateway. Customers integrate with the multi-channel security acceleration gateway based on key signatures. Each site contains only one key, which can be used to integrate all gateways under the site. After the key is refreshed, the original key becomes invalid.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) RefreshMultiPathGatewaySecretKey(request *RefreshMultiPathGatewaySecretKeyRequest) (response *RefreshMultiPathGatewaySecretKeyResponse, err error) {
    return c.RefreshMultiPathGatewaySecretKeyWithContext(context.Background(), request)
}

// RefreshMultiPathGatewaySecretKey
// This API is used to refresh the key of a multi-channel security acceleration gateway. Customers integrate with the multi-channel security acceleration gateway based on key signatures. Each site contains only one key, which can be used to integrate all gateways under the site. After the key is refreshed, the original key becomes invalid.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) RefreshMultiPathGatewaySecretKeyWithContext(ctx context.Context, request *RefreshMultiPathGatewaySecretKeyRequest) (response *RefreshMultiPathGatewaySecretKeyResponse, err error) {
    if request == nil {
        request = NewRefreshMultiPathGatewaySecretKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "RefreshMultiPathGatewaySecretKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RefreshMultiPathGatewaySecretKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewRefreshMultiPathGatewaySecretKeyResponse()
    err = c.Send(request, response)
    return
}

func NewRenewPlanRequest() (request *RenewPlanRequest) {
    request = &RenewPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "RenewPlan")
    
    
    return
}

func NewRenewPlanResponse() (response *RenewPlanResponse) {
    response = &RenewPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RenewPlan
// When your plan needs to be extended, you can use this API to renew it. Plan renewal is only supported for the Personal, Basic, and Standard Editions.
//
// > For cost details, refer to [Plan Cost](https://www.tencentcloud.com/document/product/1552/94158?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION_INSUFFICIENTACCOUNTBALANCE = "FailedOperation.InsufficientAccountBalance"
//  INVALIDPARAMETER_INVALIDAUTOUSEVOUCHER = "InvalidParameter.InvalidAutoUseVoucher"
//  INVALIDPARAMETER_INVALIDPERIOD = "InvalidParameter.InvalidPeriod"
//  INVALIDPARAMETER_PLANNOTFOUND = "InvalidParameter.PlanNotFound"
//  OPERATIONDENIED_ENTERPRISEPLANRENEWUNSUPPORTED = "OperationDenied.EnterprisePlanRenewUnsupported"
//  OPERATIONDENIED_RESOURCEHASBEENLOCKED = "OperationDenied.ResourceHasBeenLocked"
func (c *Client) RenewPlan(request *RenewPlanRequest) (response *RenewPlanResponse, err error) {
    return c.RenewPlanWithContext(context.Background(), request)
}

// RenewPlan
// When your plan needs to be extended, you can use this API to renew it. Plan renewal is only supported for the Personal, Basic, and Standard Editions.
//
// > For cost details, refer to [Plan Cost](https://www.tencentcloud.com/document/product/1552/94158?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION_INSUFFICIENTACCOUNTBALANCE = "FailedOperation.InsufficientAccountBalance"
//  INVALIDPARAMETER_INVALIDAUTOUSEVOUCHER = "InvalidParameter.InvalidAutoUseVoucher"
//  INVALIDPARAMETER_INVALIDPERIOD = "InvalidParameter.InvalidPeriod"
//  INVALIDPARAMETER_PLANNOTFOUND = "InvalidParameter.PlanNotFound"
//  OPERATIONDENIED_ENTERPRISEPLANRENEWUNSUPPORTED = "OperationDenied.EnterprisePlanRenewUnsupported"
//  OPERATIONDENIED_RESOURCEHASBEENLOCKED = "OperationDenied.ResourceHasBeenLocked"
func (c *Client) RenewPlanWithContext(ctx context.Context, request *RenewPlanRequest) (response *RenewPlanResponse, err error) {
    if request == nil {
        request = NewRenewPlanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "RenewPlan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RenewPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewRenewPlanResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradePlanRequest() (request *UpgradePlanRequest) {
    request = &UpgradePlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "UpgradePlan")
    
    
    return
}

func NewUpgradePlanResponse() (response *UpgradePlanResponse) {
    response = &UpgradePlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpgradePlan
// When you need features available only in higher-tier plans, you can upgrade your plan through this interface. Upgrades are only supported for Personal and Basic Edition Plans.
//
// > For differences between different types of Edgeone billing plans, see Edgeone billing plan comparison (https://www.tencentcloud.com/document/product/1552/94165?from_cn_redirect=1)
//
// For billing package upgrade rules and pricing details, see Edgeone billing package upgrade instructions (https://www.tencentcloud.com/document/product/1552/95291?from_cn_redirect=1)
//
// If your plan needs to upgrade to the Enterprise Edition, please [contact us](https://www.tencentcloud.com/contact-us).
//
// error code that may be returned:
//  FAILEDOPERATION_INSUFFICIENTACCOUNTBALANCE = "FailedOperation.InsufficientAccountBalance"
//  INVALIDPARAMETER_INVALIDAUTOUSEVOUCHER = "InvalidParameter.InvalidAutoUseVoucher"
//  INVALIDPARAMETER_INVALIDPLANTYPE = "InvalidParameter.InvalidPlanType"
//  INVALIDPARAMETER_PLANNOTFOUND = "InvalidParameter.PlanNotFound"
//  OPERATIONDENIED_ENTERPRISEPLANUPGRADEUNSUPPORTED = "OperationDenied.EnterprisePlanUpgradeUnsupported"
//  OPERATIONDENIED_PLANDOWNGRADENOTALLOWED = "OperationDenied.PlanDowngradeNotAllowed"
//  OPERATIONDENIED_PLANHASBEENEXPIRED = "OperationDenied.PlanHasBeenExpired"
//  OPERATIONDENIED_PLANHASBEENISOLATED = "OperationDenied.PlanHasBeenIsolated"
//  OPERATIONDENIED_RESOURCEHASBEENLOCKED = "OperationDenied.ResourceHasBeenLocked"
func (c *Client) UpgradePlan(request *UpgradePlanRequest) (response *UpgradePlanResponse, err error) {
    return c.UpgradePlanWithContext(context.Background(), request)
}

// UpgradePlan
// When you need features available only in higher-tier plans, you can upgrade your plan through this interface. Upgrades are only supported for Personal and Basic Edition Plans.
//
// > For differences between different types of Edgeone billing plans, see Edgeone billing plan comparison (https://www.tencentcloud.com/document/product/1552/94165?from_cn_redirect=1)
//
// For billing package upgrade rules and pricing details, see Edgeone billing package upgrade instructions (https://www.tencentcloud.com/document/product/1552/95291?from_cn_redirect=1)
//
// If your plan needs to upgrade to the Enterprise Edition, please [contact us](https://www.tencentcloud.com/contact-us).
//
// error code that may be returned:
//  FAILEDOPERATION_INSUFFICIENTACCOUNTBALANCE = "FailedOperation.InsufficientAccountBalance"
//  INVALIDPARAMETER_INVALIDAUTOUSEVOUCHER = "InvalidParameter.InvalidAutoUseVoucher"
//  INVALIDPARAMETER_INVALIDPLANTYPE = "InvalidParameter.InvalidPlanType"
//  INVALIDPARAMETER_PLANNOTFOUND = "InvalidParameter.PlanNotFound"
//  OPERATIONDENIED_ENTERPRISEPLANUPGRADEUNSUPPORTED = "OperationDenied.EnterprisePlanUpgradeUnsupported"
//  OPERATIONDENIED_PLANDOWNGRADENOTALLOWED = "OperationDenied.PlanDowngradeNotAllowed"
//  OPERATIONDENIED_PLANHASBEENEXPIRED = "OperationDenied.PlanHasBeenExpired"
//  OPERATIONDENIED_PLANHASBEENISOLATED = "OperationDenied.PlanHasBeenIsolated"
//  OPERATIONDENIED_RESOURCEHASBEENLOCKED = "OperationDenied.ResourceHasBeenLocked"
func (c *Client) UpgradePlanWithContext(ctx context.Context, request *UpgradePlanRequest) (response *UpgradePlanResponse, err error) {
    if request == nil {
        request = NewUpgradePlanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "UpgradePlan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradePlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradePlanResponse()
    err = c.Send(request, response)
    return
}

func NewVerifyOwnershipRequest() (request *VerifyOwnershipRequest) {
    request = &VerifyOwnershipRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "VerifyOwnership")
    
    
    return
}

func NewVerifyOwnershipResponse() (response *VerifyOwnershipResponse) {
    response = &VerifyOwnershipResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// VerifyOwnership
// In CNAME access mode, you need to verify the ownership of the site or domain name. This API is used to trigger the verification. If the site passes ownership verification, no need to verify again when adding domains subsequently. For details, see [Site/Domain Name Ownership Verification](https://www.tencentcloud.com/document/product/1552/70789?from_cn_redirect=1).
//
// 
//
// In NS access mode, you can also query through this API whether the NS server switch is successful. For details, see [Modify DNS Server](https://www.tencentcloud.com/document/product/1552/90452?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) VerifyOwnership(request *VerifyOwnershipRequest) (response *VerifyOwnershipResponse, err error) {
    return c.VerifyOwnershipWithContext(context.Background(), request)
}

// VerifyOwnership
// In CNAME access mode, you need to verify the ownership of the site or domain name. This API is used to trigger the verification. If the site passes ownership verification, no need to verify again when adding domains subsequently. For details, see [Site/Domain Name Ownership Verification](https://www.tencentcloud.com/document/product/1552/70789?from_cn_redirect=1).
//
// 
//
// In NS access mode, you can also query through this API whether the NS server switch is successful. For details, see [Modify DNS Server](https://www.tencentcloud.com/document/product/1552/90452?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) VerifyOwnershipWithContext(ctx context.Context, request *VerifyOwnershipRequest) (response *VerifyOwnershipResponse, err error) {
    if request == nil {
        request = NewVerifyOwnershipRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "teo", APIVersion, "VerifyOwnership")
    
    if c.GetCredential() == nil {
        return nil, errors.New("VerifyOwnership require credential")
    }

    request.SetContext(ctx)
    
    response = NewVerifyOwnershipResponse()
    err = c.Send(request, response)
    return
}
