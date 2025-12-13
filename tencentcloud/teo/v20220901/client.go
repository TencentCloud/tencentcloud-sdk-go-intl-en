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
// This API is used to apply for a free certificate. If you need to proceed with DNS delegated verification or file verification, you can call this API to initiate the certificate application and obtain the corresponding verification content based on the application method. The order for API calls is as follows:.
//
// Step 1: Call ApplyFreeCertificate, specify the verification method for free certificate application, and obtain the verification content.
//
// Step 2: Configure the corresponding domain as verification content.
//
// Step 3: Call CheckFreeCertificateVerification to verify. After verification passes, the free certificate application is completed.
//
// Step 4: Call ModifyHostsCertificate to issue a domain certificate configured to use the EdgeOne free certificate.
//
// 
//
// The application method introduction in the document: [Free Certificate Application Description](https://www.tencentcloud.comom/document/product/1552/90437?from_cn_redirect=1). 
//
// description:.
//
// - Only CNAME access mode can call this API to specify the free certificate application method. NS/DNSPod hosting access modes use automatic validation to apply for free certificates with no need to call this API.
//
// - If you need to switch the free certificate authentication method, you can call this API again by changing the VerificationMethod field to update it.
//
// - A domain name can only apply for one free certificate. After calling this API, the backend will trigger the free certificate application task. You need to complete the domain name verification info configuration within 2 days, then finish certificate authentication.
func (c *Client) ApplyFreeCertificate(request *ApplyFreeCertificateRequest) (response *ApplyFreeCertificateResponse, err error) {
    return c.ApplyFreeCertificateWithContext(context.Background(), request)
}

// ApplyFreeCertificate
// This API is used to apply for a free certificate. If you need to proceed with DNS delegated verification or file verification, you can call this API to initiate the certificate application and obtain the corresponding verification content based on the application method. The order for API calls is as follows:.
//
// Step 1: Call ApplyFreeCertificate, specify the verification method for free certificate application, and obtain the verification content.
//
// Step 2: Configure the corresponding domain as verification content.
//
// Step 3: Call CheckFreeCertificateVerification to verify. After verification passes, the free certificate application is completed.
//
// Step 4: Call ModifyHostsCertificate to issue a domain certificate configured to use the EdgeOne free certificate.
//
// 
//
// The application method introduction in the document: [Free Certificate Application Description](https://www.tencentcloud.comom/document/product/1552/90437?from_cn_redirect=1). 
//
// description:.
//
// - Only CNAME access mode can call this API to specify the free certificate application method. NS/DNSPod hosting access modes use automatic validation to apply for free certificates with no need to call this API.
//
// - If you need to switch the free certificate authentication method, you can call this API again by changing the VerificationMethod field to update it.
//
// - A domain name can only apply for one free certificate. After calling this API, the backend will trigger the free certificate application task. You need to complete the domain name verification info configuration within 2 days, then finish certificate authentication.
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
// This API is used to bind/unbind a domain name to/from a specific policy template. 
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
// This API is used to bind/unbind a domain name to/from a specific policy template. 
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
// This API is used to bind/unbind a domain name to/from a shared CNAME. It is now only available to beta users.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) BindSharedCNAME(request *BindSharedCNAMERequest) (response *BindSharedCNAMEResponse, err error) {
    return c.BindSharedCNAMEWithContext(context.Background(), request)
}

// BindSharedCNAME
// This API is used to bind/unbind a domain name to/from a shared CNAME. It is now only available to beta users.
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
// This API is used to bind a site to a plan.
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
// This API is used to bind a site to a plan.
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
// This API is used to query the CNAME status of a domain name.
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
// This API is used to query the CNAME status of a domain name.
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
// This API is used to verify a free certificate and obtain the application result. If verified, you can query the free certificate information for the corresponding domain name application through this API. If failed to apply, this API will return the corresponding verification failure message.
//
// This API is used to check the free certificate application result after triggering the [ApplyFreeCertificate](https://www.tencentcloud.comom/document/product/1552/124807?from_cn_redirect=1) . Once the application is successful, you need to configure through the [ModifyHostsCertificate](https://www.tencentcloud.comom/document/product/1552/80764?from_cn_redirect=1) to deploy the free certificate to the acceleration domain.
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
// This API is used to verify a free certificate and obtain the application result. If verified, you can query the free certificate information for the corresponding domain name application through this API. If failed to apply, this API will return the corresponding verification failure message.
//
// This API is used to check the free certificate application result after triggering the [ApplyFreeCertificate](https://www.tencentcloud.comom/document/product/1552/124807?from_cn_redirect=1) . Once the application is successful, you need to configure through the [ModifyHostsCertificate](https://www.tencentcloud.comom/document/product/1552/80764?from_cn_redirect=1) to deploy the free certificate to the acceleration domain.
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
// This API is used to confirm the latest origin IP range is updated to the origin server firewall when the multi-channel security acceleration gateway's origin IP range changes.
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
// This API is used to confirm the latest origin IP range is updated to the origin server firewall when the multi-channel security acceleration gateway's origin IP range changes.
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
// This API is used to confirm that the latest origin ACLs have been updated to the origin server firewall when the origin ACLs change. After confirming the update to the latest version, related change notifications will stop pushing.
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
// This API is used to confirm that the latest origin ACLs have been updated to the origin server firewall when the origin ACLs change. After confirming the update to the latest version, related change notifications will stop pushing.
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
// This API is used to create an acceleration domain name. 
//
// 
//
// For sites connected via the CNAME, if you have not verified the ownership of the domain name, the ownership verification information of the domain name is returned. To verify your ownership of the domain name, see [Ownership Verification](https://intl.cloud.tencent.com/document/product/1552/70789?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
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
//  INVALIDPARAMETERVALUE_ORIGINGROUPNOTEXISTS = "InvalidParameterValue.OriginGroupNotExists"
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
// This API is used to create an acceleration domain name. 
//
// 
//
// For sites connected via the CNAME, if you have not verified the ownership of the domain name, the ownership verification information of the domain name is returned. To verify your ownership of the domain name, see [Ownership Verification](https://intl.cloud.tencent.com/document/product/1552/70789?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
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
//  INVALIDPARAMETERVALUE_ORIGINGROUPNOTEXISTS = "InvalidParameterValue.OriginGroupNotExists"
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
// The feature is only supported by the enterprise plan and is currently in closed beta testing. If you need to use it, please [contact us](https://www.tencentcloud.com/contact-us).
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
// The feature is only supported by the enterprise plan and is currently in closed beta testing. If you need to use it, please [contact us](https://www.tencentcloud.com/contact-us).
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
// This API is on an earlier version. If you want to call it, please switch to the latest version [CreateL4Proxy] (https://intl.cloud.tencent.com/document/product/1552/103417?from_cn_redirect=1).
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
// This API is on an earlier version. If you want to call it, please switch to the latest version [CreateL4Proxy] (https://intl.cloud.tencent.com/document/product/1552/103417?from_cn_redirect=1).
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
// This API is on an earlier version. If you want to call it, please switch to the latest version. For details, see [CreateL4ProxyRules] (https://intl.cloud.tencent.com/document/product/1552/103416?from_cn_redirect=1).
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
// This API is on an earlier version. If you want to call it, please switch to the latest version. For details, see [CreateL4ProxyRules] (https://intl.cloud.tencent.com/document/product/1552/103416?from_cn_redirect=1).
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
// This API is used to create key-value indexes for relevant delivered log fields in the corresponding Tencent Cloud CLS log topic for a specified real-time log delivery task (task-id). If such indexes have been created in CLS, this API will append indexes through merging.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATECLSLOGSETFAILED = "FailedOperation.CreateClsLogSetFailed"
func (c *Client) CreateCLSIndex(request *CreateCLSIndexRequest) (response *CreateCLSIndexResponse, err error) {
    return c.CreateCLSIndexWithContext(context.Background(), request)
}

// CreateCLSIndex
// This API is used to create key-value indexes for relevant delivered log fields in the corresponding Tencent Cloud CLS log topic for a specified real-time log delivery task (task-id). If such indexes have been created in CLS, this API will append indexes through merging.
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
// This API is used to create a new version for the specified configuration group in version management mode. The version management feature is currently undergoing beta testing and is accessible only to users on the whitelist.
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
// This API is used to create a new version for the specified configuration group in version management mode. The version management feature is currently undergoing beta testing and is accessible only to users on the whitelist.
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
// This API is used to create content identifiers, where you can set descriptions, tags, and other information. It is also necessary to bind an enterprise edition package for billing data statistics. A content identifier can only bind one billing package, while a billing package can bind multiple content identifiers. This feature is only available to the allowlist.
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
// This API is used to create content identifiers, where you can set descriptions, tags, and other information. It is also necessary to bind an enterprise edition package for billing data statistics. A content identifier can only bind one billing package, while a billing package can bind multiple content identifiers. This feature is only available to the allowlist.
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
// After creating a site and the site is accessed in NS mode, you can create DNS records through this API.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_CONTENTTYPENOTMATCH = "InvalidParameterValue.ContentTypeNotMatch"
//  INVALIDPARAMETERVALUE_PAGENAMEALREADYEXIST = "InvalidParameterValue.PageNameAlreadyExist"
func (c *Client) CreateDnsRecord(request *CreateDnsRecordRequest) (response *CreateDnsRecordResponse, err error) {
    return c.CreateDnsRecordWithContext(context.Background(), request)
}

// CreateDnsRecord
// After creating a site and the site is accessed in NS mode, you can create DNS records through this API.
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
// This API is used to create a trigger rule for an edge function.
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
// This API is used to create a trigger rule for an edge function.
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
// JIT transcoding already provides preset transcoding templates to meet most needs. If there are personalized transcoding requirements, you can create custom transcoding templates through this API, with up to 100 custom transcoding templates allowed.
//
// This API is used to ensure the consistency of JIT transcoding effect, avoid video output exceptions caused by EO cache or M3U8 sharding template changes during the process, and templates cannot be modified after creation.
//
// This API is used to learn about the detailed capacity of JIT transcoding. EdgeOne video instant processing function introduction (https://www.tencentcloud.comom/document/product/1552/111927?from_cn_redirect=1).
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
// JIT transcoding already provides preset transcoding templates to meet most needs. If there are personalized transcoding requirements, you can create custom transcoding templates through this API, with up to 100 custom transcoding templates allowed.
//
// This API is used to ensure the consistency of JIT transcoding effect, avoid video output exceptions caused by EO cache or M3U8 sharding template changes during the process, and templates cannot be modified after creation.
//
// This API is used to learn about the detailed capacity of JIT transcoding. EdgeOne video instant processing function introduction (https://www.tencentcloud.comom/document/product/1552/111927?from_cn_redirect=1).
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
// This API is used to create Layer 4 proxy instances.
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
// This API is used to create Layer 4 proxy instances.
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
// This API is used to create Layer 4 proxy instance rules, supporting both individual and batch creation.
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
// This API is used to create Layer 4 proxy instance rules, supporting both individual and batch creation.
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
// This API is used to create rules in the [rule engine](https://intl.cloud.tencent.com/document/product/1552/70901?from_cn_redirect=1). Batch creation is supported.
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
func (c *Client) CreateL7AccRules(request *CreateL7AccRulesRequest) (response *CreateL7AccRulesResponse, err error) {
    return c.CreateL7AccRulesWithContext(context.Background(), request)
}

// CreateL7AccRules
// This API is used to create rules in the [rule engine](https://intl.cloud.tencent.com/document/product/1552/70901?from_cn_redirect=1). Batch creation is supported.
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
// This API is used to create a LoadBalancer. For details, see [Quickly Create Load Balancers](https://intl.cloud.tencent.com/document/product/1552/104223?from_cn_redirect=1). The load balancing feature is in beta test. If you need to use it, [contact us](https://www.tencentcloud.com/contact-us).
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
// This API is used to create a LoadBalancer. For details, see [Quickly Create Load Balancers](https://intl.cloud.tencent.com/document/product/1552/104223?from_cn_redirect=1). The load balancing feature is in beta test. If you need to use it, [contact us](https://www.tencentcloud.com/contact-us).
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
// Create a multi-channel security acceleration gateway via this API, including Cloud Gateway (gateway created and managed by Tencent Cloud) and private gateway (gateway deployed by users). Query the status using DescribeMultiPathGateway, and creation is successful if the status is online.
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
// Create a multi-channel security acceleration gateway via this API, including Cloud Gateway (gateway created and managed by Tencent Cloud) and private gateway (gateway deployed by users). Query the status using DescribeMultiPathGateway, and creation is successful if the status is online.
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
// This API is used to create lines integrated with the multi-channel security acceleration gateway, including EdgeOne Layer-4 proxy and custom lines.
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
// This API is used to create lines integrated with the multi-channel security acceleration gateway, including EdgeOne Layer-4 proxy and custom lines.
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
// This API creates an access key for the multi-channel security acceleration gateway. Customers use the access key to sign requests for accessing the gateway. Each site can have only one key, which is applicable to all gateways under that site. Query the key via the DescribeMultiPathGatewaySecretKey API.
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
// This API creates an access key for the multi-channel security acceleration gateway. Customers use the access key to sign requests for accessing the gateway. Each site can have only one key, which is applicable to all gateways under that site. Query the key via the DescribeMultiPathGatewaySecretKey API.
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
// This API is used to create an origin group for easy management. The created origin server group can be used for **adding acceleration domain names** and **layer-4 proxy configuration**.
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
// This API is used to create an origin group for easy management. The created origin server group can be used for **adding acceleration domain names** and **layer-4 proxy configuration**.
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
// If you need to use the EdgeOne product, you must create a billing plan through this interface.
//
// > After creating a plan, you need to complete the process of creating a site and binding the plan through [CreateZone](https://intl.cloud.tencent.com/document/product/1552/80719?from_cn_redirect=1), so that the EdgeOne can provide services properly.
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
// If you need to use the EdgeOne product, you must create a billing plan through this interface.
//
// > After creating a plan, you need to complete the process of creating a site and binding the plan through [CreateZone](https://intl.cloud.tencent.com/document/product/1552/80719?from_cn_redirect=1), so that the EdgeOne can provide services properly.
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
// This API is used to purchase a plan for a new site.
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
// This API is used to purchase a plan for a new site.
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
// This API is used to create a pre-warming task.
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
// This API is used to create a pre-warming task.
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
// When there are resources updated on the origin with the TTL remaining valid, users cannot access the latest resources. In this case, you can purge the cache using this API. There are two methods: <li>Delete: This method deletes the node cache without verification and retrieves the latest resources from the origin when receiving a request.</li><li>Invalidate: This method marks the node cache as invalid and sends a request with the If-None-Match and If-Modified-Since headers to the origin. If the origin responses with 200, the latest resources are retrieved to be cached on the node. If a 304 response is returned, the latest resources are not cached on the node.
//
// 
//
// </li>For more details, see [Cache Purge](https://intl.cloud.tencent.com/document/product/1552/70759?from_cn_redirect=1). </li>
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
// When there are resources updated on the origin with the TTL remaining valid, users cannot access the latest resources. In this case, you can purge the cache using this API. There are two methods: <li>Delete: This method deletes the node cache without verification and retrieves the latest resources from the origin when receiving a request.</li><li>Invalidate: This method marks the node cache as invalid and sends a request with the If-None-Match and If-Modified-Since headers to the origin. If the origin responses with 200, the latest resources are retrieved to be cached on the node. If a 304 response is returned, the latest resources are not cached on the node.
//
// 
//
// </li>For more details, see [Cache Purge](https://intl.cloud.tencent.com/document/product/1552/70759?from_cn_redirect=1). </li>
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
// The following restrictions apply:
//
// 
//
// - When the log type (`LogType`) is site acceleration log (L7 access log) (`domain`), L4 proxy log (`application`), or Edge Function execution log (`function`), the same entity (L7 domain, L4 proxy instance, or Edge Function instance) can be added to only one of the following `TaskType` combinations within the same `LogType`-`Area` pair:
//
//     - One task delivering to Tencent Cloud CLS plus one task delivering to a custom HTTP(S) endpoint;
//
//     - One task delivering to Tencent Cloud CLS plus one task delivering to an AWS S3-compatible bucket.
//
// - When the log type (`LogType`) is rate-limiting & CC attack protection log (`web-rateLiming`), managed rule log (`web-attack`), custom rule log (`web-rule`), or bot management log (`web-bot`), the same entity can be added to only one real-time log delivery task within the same `LogType`-`Area` pair.
//
// 
//
// Before creating a task, we recommend that you first call [DescribeRealtimeLogDeliveryTasks](https://intl.cloud.tencent.com/document/product/1552/104110?from_cn_redirect=1) to list existing tasks for the entity and verify whether it has already been added to another task.
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
// The following restrictions apply:
//
// 
//
// - When the log type (`LogType`) is site acceleration log (L7 access log) (`domain`), L4 proxy log (`application`), or Edge Function execution log (`function`), the same entity (L7 domain, L4 proxy instance, or Edge Function instance) can be added to only one of the following `TaskType` combinations within the same `LogType`-`Area` pair:
//
//     - One task delivering to Tencent Cloud CLS plus one task delivering to a custom HTTP(S) endpoint;
//
//     - One task delivering to Tencent Cloud CLS plus one task delivering to an AWS S3-compatible bucket.
//
// - When the log type (`LogType`) is rate-limiting & CC attack protection log (`web-rateLiming`), managed rule log (`web-attack`), custom rule log (`web-rule`), or bot management log (`web-bot`), the same entity can be added to only one real-time log delivery task within the same `LogType`-`Area` pair.
//
// 
//
// Before creating a task, we recommend that you first call [DescribeRealtimeLogDeliveryTasks](https://intl.cloud.tencent.com/document/product/1552/104110?from_cn_redirect=1) to list existing tasks for the entity and verify whether it has already been added to another task.
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
// This interface is the old version of the rule engine creation interface. EdgeOne has fully upgraded the rule engine related interfaces on January 21, 2025. For details on the new version of the seven-layer acceleration rule creation interface, please refer to [CreateL7AccRules](https://intl.cloud.tencent.com/document/product/1552/115822?from_cn_redirect=1).<p style="color: red;">Note: Starting from January 21, 2025, the old version of the interface will stop updating and iteration. Subsequent new features will only be provided in the new version of the interface, and the original capabilities supported by the old version of the interface will not be affected. To avoid data field conflicts when using the old version of the interface, it is recommended that you migrate to the new version of the rule engine interface as soon as possible. </p>
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
// This interface is the old version of the rule engine creation interface. EdgeOne has fully upgraded the rule engine related interfaces on January 21, 2025. For details on the new version of the seven-layer acceleration rule creation interface, please refer to [CreateL7AccRules](https://intl.cloud.tencent.com/document/product/1552/115822?from_cn_redirect=1).<p style="color: red;">Note: Starting from January 21, 2025, the old version of the interface will stop updating and iteration. Subsequent new features will only be provided in the new version of the interface, and the original capabilities supported by the old version of the interface will not be affected. To avoid data field conflicts when using the old version of the interface, it is recommended that you migrate to the new version of the rule engine interface as soon as possible. </p>
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
// This API is used to create an API resource.
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
// This API is used to create an API resource.
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
// This API is used to create an API service.
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
// This API is used to create an API service.
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
// This API is used to create client authentication options.
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
// This API is used to create client authentication options.
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
// This API is used to create a security IP group.
//
// error code that may be returned:
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateSecurityIPGroup(request *CreateSecurityIPGroupRequest) (response *CreateSecurityIPGroupResponse, err error) {
    return c.CreateSecurityIPGroupWithContext(context.Background(), request)
}

// CreateSecurityIPGroup
// This API is used to create a security IP group.
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
// This API is used to create a JavaScript injection rule.
//
// error code that may be returned:
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateSecurityJSInjectionRule(request *CreateSecurityJSInjectionRuleRequest) (response *CreateSecurityJSInjectionRuleResponse, err error) {
    return c.CreateSecurityJSInjectionRuleWithContext(context.Background(), request)
}

// CreateSecurityJSInjectionRule
// This API is used to create a JavaScript injection rule.
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
// This API is used to create a shared CNAME. It is now only available to beta users.
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
// This API is used to create a shared CNAME. It is now only available to beta users.
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
// This API is used to create a security policy configuration template.
//
// error code that may be returned:
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  LIMITEXCEEDED_SECURITY = "LimitExceeded.Security"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateWebSecurityTemplate(request *CreateWebSecurityTemplateRequest) (response *CreateWebSecurityTemplateResponse, err error) {
    return c.CreateWebSecurityTemplateWithContext(context.Background(), request)
}

// CreateWebSecurityTemplate
// This API is used to create a security policy configuration template.
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
// This API is used to create a site. After you create the site, you can connect it to EdgeOne via the CNAME or NS (see [Quick Start](https://intl.cloud.tencent.com/document/product/1552/87601?from_cn_redirect=1)), or connect it without a domain name (see [Quick Access to L4 Proxy Service](https://intl.cloud.tencent.com/document/product/1552/96051?from_cn_redirect=1)).
//
// 
//
// If there are already EdgeOne plans under the current account, it is recommended to pass in the `PlanId` to bind the site with the plan directly. If `PlanId` is not passed in, the created site is not activated. You need to call [BindZoneToPlan](https://intl.cloud.tencent.com/document/product/1552/83042?from_cn_redirect=1) to bind the site with a plan. To purchase a plan, please go to the EdgeOne console.
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
// This API is used to create a site. After you create the site, you can connect it to EdgeOne via the CNAME or NS (see [Quick Start](https://intl.cloud.tencent.com/document/product/1552/87601?from_cn_redirect=1)), or connect it without a domain name (see [Quick Access to L4 Proxy Service](https://intl.cloud.tencent.com/document/product/1552/96051?from_cn_redirect=1)).
//
// 
//
// If there are already EdgeOne plans under the current account, it is recommended to pass in the `PlanId` to bind the site with the plan directly. If `PlanId` is not passed in, the created site is not activated. You need to call [BindZoneToPlan](https://intl.cloud.tencent.com/document/product/1552/83042?from_cn_redirect=1) to bind the site with a plan. To purchase a plan, please go to the EdgeOne console.
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
// This API is used to batch remove accelerated domain names.
//
// error code that may be returned:
//  INVALIDPARAMETER_ZONEISGRAYPUBLISHING = "InvalidParameter.ZoneIsGrayPublishing"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  INVALIDPARAMETERVALUE_INVALIDDOMAINSTATUS = "InvalidParameterValue.InvalidDomainStatus"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  OPERATIONDENIED_VERSIONCONTROLISGRAYING = "OperationDenied.VersionControlIsGraying"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE_DOMAINNOTFOUND = "ResourceUnavailable.DomainNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteAccelerationDomains(request *DeleteAccelerationDomainsRequest) (response *DeleteAccelerationDomainsResponse, err error) {
    return c.DeleteAccelerationDomainsWithContext(context.Background(), request)
}

// DeleteAccelerationDomains
// This API is used to batch remove accelerated domain names.
//
// error code that may be returned:
//  INVALIDPARAMETER_ZONEISGRAYPUBLISHING = "InvalidParameter.ZoneIsGrayPublishing"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  INVALIDPARAMETERVALUE_INVALIDDOMAINSTATUS = "InvalidParameterValue.InvalidDomainStatus"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
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
// This API is used to delete an alias domain name.
//
// The feature is only supported by the enterprise plan and is currently in closed beta testing. If you need to use it, [Contact Us](https://www.tencentcloud.com/contact-us).
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
// This API is used to delete an alias domain name.
//
// The feature is only supported by the enterprise plan and is currently in closed beta testing. If you need to use it, [Contact Us](https://www.tencentcloud.com/contact-us).
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
// This API is on an earlier version. If you want to call it, please switch to the latest version. For details, see [DeleteL4Proxy] (https://intl.cloud.tencent.com/document/product/1552/103415?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteApplicationProxy(request *DeleteApplicationProxyRequest) (response *DeleteApplicationProxyResponse, err error) {
    return c.DeleteApplicationProxyWithContext(context.Background(), request)
}

// DeleteApplicationProxy
// This API is on an earlier version. If you want to call it, please switch to the latest version. For details, see [DeleteL4Proxy] (https://intl.cloud.tencent.com/document/product/1552/103415?from_cn_redirect=1).
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
// This API is on an earlier version. If you want to call it, please switch to the latest version. For details, see [DeleteL4ProxyRules] (https://intl.cloud.tencent.com/document/product/1552/103414?from_cn_redirect=1).
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteApplicationProxyRule(request *DeleteApplicationProxyRuleRequest) (response *DeleteApplicationProxyRuleResponse, err error) {
    return c.DeleteApplicationProxyRuleWithContext(context.Background(), request)
}

// DeleteApplicationProxyRule
// This API is on an earlier version. If you want to call it, please switch to the latest version. For details, see [DeleteL4ProxyRules] (https://intl.cloud.tencent.com/document/product/1552/103414?from_cn_redirect=1).
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
// Delete the specified content identifier. This feature is only available to the allowlist.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteContentIdentifier(request *DeleteContentIdentifierRequest) (response *DeleteContentIdentifierResponse, err error) {
    return c.DeleteContentIdentifierWithContext(context.Background(), request)
}

// DeleteContentIdentifier
// Delete the specified content identifier. This feature is only available to the allowlist.
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
// You can use this API to batch delete DNS records.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDnsRecords(request *DeleteDnsRecordsRequest) (response *DeleteDnsRecordsResponse, err error) {
    return c.DeleteDnsRecordsWithContext(context.Background(), request)
}

// DeleteDnsRecords
// You can use this API to batch delete DNS records.
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
// This API is used to delete an edge function. Once deleted, the function cannot be recovered, and associated trigger rules are also deleted.
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
// This API is used to delete an edge function. Once deleted, the function cannot be recovered, and associated trigger rules are also deleted.
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
// This API is used to delete the appropriate just in time transcoding template based on the unique template identifier under the site ID.
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
// This API is used to delete the appropriate just in time transcoding template based on the unique template identifier under the site ID.
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
// This API is used to delete a Layer 4 proxy instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteL4Proxy(request *DeleteL4ProxyRequest) (response *DeleteL4ProxyResponse, err error) {
    return c.DeleteL4ProxyWithContext(context.Background(), request)
}

// DeleteL4Proxy
// This API is used to delete a Layer 4 proxy instance.
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
// This API is used to delete Layer 4 proxy forwarding rules, supporting both individual and batch operation.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteL4ProxyRules(request *DeleteL4ProxyRulesRequest) (response *DeleteL4ProxyRulesResponse, err error) {
    return c.DeleteL4ProxyRulesWithContext(context.Background(), request)
}

// DeleteL4ProxyRules
// This API is used to delete Layer 4 proxy forwarding rules, supporting both individual and batch operation.
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
// This API is used to delete rules of the [rule engine](https://intl.cloud.tencent.com/document/product/1552/70901?from_cn_redirect=1), supporting batch deletion.
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
// This API is used to delete rules of the [rule engine](https://intl.cloud.tencent.com/document/product/1552/70901?from_cn_redirect=1), supporting batch deletion.
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
// This API is used to delete a LoadBalancer. If the LoadBalancer is referenced by other services (for example, Layer-4 proxy), the LoadBalancer cannot be deleted until the reference relationship is removed. The load balancing feature is in beta test. If you need to use it, [contact us](https://www.tencentcloud.com/contact-us).
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
// This API is used to delete a LoadBalancer. If the LoadBalancer is referenced by other services (for example, Layer-4 proxy), the LoadBalancer cannot be deleted until the reference relationship is removed. The load balancing feature is in beta test. If you need to use it, [contact us](https://www.tencentcloud.com/contact-us).
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
// This API is used to delete a multi-channel security acceleration gateway, including private gateways and Cloud Gateways.
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
// This API is used to delete a multi-channel security acceleration gateway, including private gateways and Cloud Gateways.
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
// This API is used to delete lines integrated with the multi-channel security acceleration gateway. Only custom lines support deletion.
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
// This API is used to delete lines integrated with the multi-channel security acceleration gateway. Only custom lines support deletion.
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
// This API is used to delete an origin group. Note that an origin group can not be deleted if it is referenced by services (e.g. L4 Proxy, domain name service, load balancing, rule engines). 
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
// This API is used to delete an origin group. Note that an origin group can not be deleted if it is referenced by services (e.g. L4 Proxy, domain name service, load balancing, rule engines). 
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
// This interface is the old version of the rule engine deletion interface. EdgeOne has fully upgraded the rule engine related interfaces on January 21, 2025. For details on the new version of the seven-layer acceleration rule deletion interface, please refer to [DeleteL7AccRules](https://intl.cloud.tencent.com/document/product/1552/115821?from_cn_redirect=1).<0>Note: Starting from January 21, 2025, the earlier version API will no longer be updated. Subsequent new features will only be provided in the latest version interface. The original capabilities supported by the earlier version API will not be affected. To avoid field conflicts when using the earlier version API, it is recommended that you migrate to the new version rule engine API as soon as possible.</0>.
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
// This interface is the old version of the rule engine deletion interface. EdgeOne has fully upgraded the rule engine related interfaces on January 21, 2025. For details on the new version of the seven-layer acceleration rule deletion interface, please refer to [DeleteL7AccRules](https://intl.cloud.tencent.com/document/product/1552/115821?from_cn_redirect=1).<0>Note: Starting from January 21, 2025, the earlier version API will no longer be updated. Subsequent new features will only be provided in the latest version interface. The original capabilities supported by the earlier version API will not be affected. To avoid field conflicts when using the earlier version API, it is recommended that you migrate to the new version rule engine API as soon as possible.</0>.
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
// This API is used to delete API resources.
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
// This API is used to delete API resources.
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
// This API is used to delete the API service.
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
// This API is used to delete the API service.
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
// This API is used to delete client authentication options.
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
// This API is used to delete client authentication options.
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
// This API is used to delete a specified security IP group. Note that the security IP group cannot be deleted if it is referenced in a rule.
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
// This API is used to delete a specified security IP group. Note that the security IP group cannot be deleted if it is referenced in a rule.
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
// This API is used to delete a shared CNAME. It is now only available to beta users.
//
// error code that may be returned:
//  RESOURCEINUSE_SHAREDCNAME = "ResourceInUse.SharedCNAME"
func (c *Client) DeleteSharedCNAME(request *DeleteSharedCNAMERequest) (response *DeleteSharedCNAMEResponse, err error) {
    return c.DeleteSharedCNAMEWithContext(context.Background(), request)
}

// DeleteSharedCNAME
// This API is used to delete a shared CNAME. It is now only available to beta users.
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
// This API is used to delete a security policy configuration template.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DeleteWebSecurityTemplate(request *DeleteWebSecurityTemplateRequest) (response *DeleteWebSecurityTemplateResponse, err error) {
    return c.DeleteWebSecurityTemplateWithContext(context.Background(), request)
}

// DeleteWebSecurityTemplate
// This API is used to delete a security policy configuration template.
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
// This API is used to delete a site.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
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
// This API is used to delete a site.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
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
// This API is used to release versions in version management mode. Users can deploy the version to either the testing environment or the production environment by specifying the EnvId parameter. The version management feature is currently undergoing beta testing and is accessible only to users on the whitelist.
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
// This API is used to release versions in version management mode. Users can deploy the version to either the testing environment or the production environment by specifying the EnvId parameter. The version management feature is currently undergoing beta testing and is accessible only to users on the whitelist.
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
// This API is used to query domain name information of a site, including the acceleration domain name, origin, and domain name status. You can query the information of all domain names, or specific domain names by specifying filters information.
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
// This API is used to query domain name information of a site, including the acceleration domain name, origin, and domain name status. You can query the information of all domain names, or specific domain names by specifying filters information.
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
// This API is used to query the alias domain name information list.
//
// The feature is only supported in the enterprise plan and is currently in closed beta testing. If you need to use it, [Contact Us](https://www.tencentcloud.com/contact-us).
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
// This API is used to query the alias domain name information list.
//
// The feature is only supported in the enterprise plan and is currently in closed beta testing. If you need to use it, [Contact Us](https://www.tencentcloud.com/contact-us).
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
// This API is on an earlier version. If you want to call it, please switch to the latest version. In the latest version, this API has been split into two APIs: one for querying the Layer 4 proxy instance list and the other for querying Layer 4 forwarding rules. For details, see [DescribeL4Proxy] (https://intl.cloud.tencent.com/document/product/1552/103413?from_cn_redirect=1) and [DescribeL4ProxyRules] (https://intl.cloud.tencent.com/document/product/1552/103412?from_cn_redirect=1).
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
// This API is on an earlier version. If you want to call it, please switch to the latest version. In the latest version, this API has been split into two APIs: one for querying the Layer 4 proxy instance list and the other for querying Layer 4 forwarding rules. For details, see [DescribeL4Proxy] (https://intl.cloud.tencent.com/document/product/1552/103413?from_cn_redirect=1) and [DescribeL4ProxyRules] (https://intl.cloud.tencent.com/document/product/1552/103412?from_cn_redirect=1).
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
// This API is used to query plan options available for purchase.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
func (c *Client) DescribeAvailablePlans(request *DescribeAvailablePlansRequest) (response *DescribeAvailablePlansResponse, err error) {
    return c.DescribeAvailablePlansWithContext(context.Background(), request)
}

// DescribeAvailablePlans
// This API is used to query plan options available for purchase.
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
// This API is used to query billing data.
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
// This API is used to query billing data.
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
// This API is used to obtain detailed information about a version in version management mode. The response includes the version ID, description, status, creation time, configuration group information, and the content of the version configuration file. The version management feature is currently undergoing beta testing and is accessible only to users on the whitelist.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeConfigGroupVersionDetail(request *DescribeConfigGroupVersionDetailRequest) (response *DescribeConfigGroupVersionDetailResponse, err error) {
    return c.DescribeConfigGroupVersionDetailWithContext(context.Background(), request)
}

// DescribeConfigGroupVersionDetail
// This API is used to obtain detailed information about a version in version management mode. The response includes the version ID, description, status, creation time, configuration group information, and the content of the version configuration file. The version management feature is currently undergoing beta testing and is accessible only to users on the whitelist.
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
// This API is used to query the version list for the specified configuration group in version management mode. The version management feature is currently undergoing beta testing and is accessible only to users on the whitelist.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeConfigGroupVersions(request *DescribeConfigGroupVersionsRequest) (response *DescribeConfigGroupVersionsResponse, err error) {
    return c.DescribeConfigGroupVersionsWithContext(context.Background(), request)
}

// DescribeConfigGroupVersions
// This API is used to query the version list for the specified configuration group in version management mode. The version management feature is currently undergoing beta testing and is accessible only to users on the whitelist.
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
// Batch query content identifiers, which can be filtered by ID, description, status, or Tag. Deleted content identifiers queried by status are retained for only three months. This feature is only open to the allowlist.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeContentIdentifiers(request *DescribeContentIdentifiersRequest) (response *DescribeContentIdentifiersResponse, err error) {
    return c.DescribeContentIdentifiersWithContext(context.Background(), request)
}

// DescribeContentIdentifiers
// Batch query content identifiers, which can be filtered by ID, description, status, or Tag. Deleted content identifiers queried by status are retained for only three months. This feature is only open to the allowlist.
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
// This API is used to query content management quotas.
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeContentQuota(request *DescribeContentQuotaRequest) (response *DescribeContentQuotaResponse, err error) {
    return c.DescribeContentQuotaWithContext(context.Background(), request)
}

// DescribeContentQuota
// This API is used to query content management quotas.
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
// This API is used to query the custom response page list.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
func (c *Client) DescribeCustomErrorPages(request *DescribeCustomErrorPagesRequest) (response *DescribeCustomErrorPagesResponse, err error) {
    return c.DescribeCustomErrorPagesWithContext(context.Background(), request)
}

// DescribeCustomErrorPages
// This API is used to query the custom response page list.
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
// This API is used to query the time-series data of DDoS attacks.
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
// This API is used to query the time-series data of DDoS attacks.
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
// This API is used to query DDoS attack events.
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
// This API is used to query DDoS attack events.
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
// This API is used to query the top-ranked DDoS attack data.
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
// This API is used to query the top-ranked DDoS attack data.
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
// This API is used to search for site exclusive Anti-DDoS information.
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
// This API is used to search for site exclusive Anti-DDoS information.
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
// This API is used to query a list of default certificates.
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
// This API is used to query a list of default certificates.
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
// This API is used to query the release history of versions in the production or test environment in version management mode. The version management feature is currently undergoing beta testing and is accessible only to users on the whitelist.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDeployHistory(request *DescribeDeployHistoryRequest) (response *DescribeDeployHistoryResponse, err error) {
    return c.DescribeDeployHistoryWithContext(context.Background(), request)
}

// DescribeDeployHistory
// This API is used to query the release history of versions in the production or test environment in version management mode. The version management feature is currently undergoing beta testing and is accessible only to users on the whitelist.
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
// This API is used to view DNS record information under a site, including DNS record name, record type, and record content. It supports querying specific DNS record information by specifying filter conditions.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDnsRecords(request *DescribeDnsRecordsRequest) (response *DescribeDnsRecordsResponse, err error) {
    return c.DescribeDnsRecordsWithContext(context.Background(), request)
}

// DescribeDnsRecords
// This API is used to view DNS record information under a site, including DNS record name, record type, and record content. It supports querying specific DNS record information by specifying filter conditions.
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
// This API is used to query environment information in version management mode. The response includes the environment ID, type, and current effective version. The version management feature is currently undergoing beta testing and is accessible only to users on the whitelist.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeEnvironments(request *DescribeEnvironmentsRequest) (response *DescribeEnvironmentsResponse, err error) {
    return c.DescribeEnvironmentsWithContext(context.Background(), request)
}

// DescribeEnvironments
// This API is used to query environment information in version management mode. The response includes the environment ID, type, and current effective version. The version management feature is currently undergoing beta testing and is accessible only to users on the whitelist.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
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
// This API is an old version. EdgeOne has fully upgraded the APIs related to the rule engine. You can obtain detailed configurations of domain names through [DescribeL7AccSetting](https://intl.cloud.tencent.com/document/product/1552/115819?from_cn_redirect=1) and [DescribeL7AccRules](https://intl.cloud.tencent.com/document/product/1552/115820?from_cn_redirect=1).
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
// This API is an old version. EdgeOne has fully upgraded the APIs related to the rule engine. You can obtain detailed configurations of domain names through [DescribeL7AccSetting](https://intl.cloud.tencent.com/document/product/1552/115819?from_cn_redirect=1) and [DescribeL7AccRules](https://intl.cloud.tencent.com/document/product/1552/115820?from_cn_redirect=1).
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
// This API is used to check if the IP is an EdgeOne IP.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
func (c *Client) DescribeIPRegion(request *DescribeIPRegionRequest) (response *DescribeIPRegionResponse, err error) {
    return c.DescribeIPRegionWithContext(context.Background(), request)
}

// DescribeIPRegion
// This API is used to check if the IP is an EdgeOne IP.
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
// This API is used to query the verification information of a site.
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
// This API is used to query the verification information of a site.
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
// This API is used to search the transcoding template detail list according to the name, template type, or unique identifier of the just-in-time transcoding template. The returned results include all eligible custom templates and preset templates.
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
// This API is used to search the transcoding template detail list according to the name, template type, or unique identifier of the just-in-time transcoding template. The returned results include all eligible custom templates and preset templates.
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
// This API is used to query a Layer 4 proxy instance list.
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
// This API is used to query a Layer 4 proxy instance list.
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
// This API is used to query the rule list of the rule engine (https://intl.cloud.tencent.com/document/product/1552/70901?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeL7AccRules(request *DescribeL7AccRulesRequest) (response *DescribeL7AccRulesResponse, err error) {
    return c.DescribeL7AccRulesWithContext(context.Background(), request)
}

// DescribeL7AccRules
// This API is used to query the rule list of the rule engine (https://intl.cloud.tencent.com/document/product/1552/70901?from_cn_redirect=1).
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
// This API is used to query the global configuration of [Site Acceleration](https://intl.cloud.tencent.com/document/product/1552/96193?from_cn_redirect=1).
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
// This API is used to query the global configuration of [Site Acceleration](https://intl.cloud.tencent.com/document/product/1552/96193?from_cn_redirect=1).
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
// This API is used to query the LoadBalancer list. The load balancing feature is in beta test. If you need to use it, [contact us](https://www.tencentcloud.com/contact-us).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeLoadBalancerList(request *DescribeLoadBalancerListRequest) (response *DescribeLoadBalancerListResponse, err error) {
    return c.DescribeLoadBalancerListWithContext(context.Background(), request)
}

// DescribeLoadBalancerList
// This API is used to query the LoadBalancer list. The load balancing feature is in beta test. If you need to use it, [contact us](https://www.tencentcloud.com/contact-us).
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
// This API is used to query multi-channel security acceleration gateway details such as name, Gateway ID, IP, port and type.
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
// This API is used to query multi-channel security acceleration gateway details such as name, Gateway ID, IP, port and type.
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
// Use this API to query the lines integrated with the multi-channel security acceleration gateway, including direct connection lines, EdgeOne Layer-4 proxy lines, and custom lines.
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
// Use this API to query the lines integrated with the multi-channel security acceleration gateway, including direct connection lines, EdgeOne Layer-4 proxy lines, and custom lines.
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
// This API is used to query the binding relationship between a multi-channel security acceleration gateway instance and the origin server IP range, as well as the IP range details. If the MultiPathGatewayNextOriginACL field has a return value, the latest origin server IP range must be synchronized to the origin server firewall configuration.
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
// This API is used to query the binding relationship between a multi-channel security acceleration gateway instance and the origin server IP range, as well as the IP range details. If the MultiPathGatewayNextOriginACL field has a return value, the latest origin server IP range must be synchronized to the origin server firewall configuration.
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
// This API is used to query the list of available regions for user-created multi-channel security acceleration gateways (Cloud Gateway).
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
// This API is used to query the list of available regions for user-created multi-channel security acceleration gateways (Cloud Gateway).
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
// This API is used to query keys for integrating multi-channel security acceleration gateways. Customers access multi-channel security acceleration gateways based on key signature.
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
// This API is used to query keys for integrating multi-channel security acceleration gateways. Customers access multi-channel security acceleration gateways based on key signature.
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
// Query the multi-channel security acceleration gateway list created by the user through this interface. Supports pagination.
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
// Query the multi-channel security acceleration gateway list created by the user through this interface. Supports pagination.
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
// This API is used to query the binding relationship between L7 acceleration domains/L4 proxy instances and origin ACLs under a site, as well as IP range details. If you want to periodically obtain the latest version of origin IP ranges through an automation script, you can poll this API at a low-frequency (recommended every three days). If the NextOriginACL field has a return value, synchronize the latest origin IP ranges to the origin server firewall configuration.
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
// This API is used to query the binding relationship between L7 acceleration domains/L4 proxy instances and origin ACLs under a site, as well as IP range details. If you want to periodically obtain the latest version of origin IP ranges through an automation script, you can poll this API at a low-frequency (recommended every three days). If the NextOriginACL field has a return value, synchronize the latest origin IP ranges to the origin server firewall configuration.
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
// This API is used to obtain a list of origin groups.
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeOriginGroup(request *DescribeOriginGroupRequest) (response *DescribeOriginGroupResponse, err error) {
    return c.DescribeOriginGroupWithContext(context.Background(), request)
}

// DescribeOriginGroup
// This API is used to obtain a list of origin groups.
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
// This API is used to query the health status of origin server groups under a LoadBalancer. The load balancing feature is in beta test. If you need to use it, [contact us](https://www.tencentcloud.com/contact-us).
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeOriginGroupHealthStatus(request *DescribeOriginGroupHealthStatusRequest) (response *DescribeOriginGroupHealthStatusResponse, err error) {
    return c.DescribeOriginGroupHealthStatusWithContext(context.Background(), request)
}

// DescribeOriginGroupHealthStatus
// This API is used to query the health status of origin server groups under a LoadBalancer. The load balancing feature is in beta test. If you need to use it, [contact us](https://www.tencentcloud.com/contact-us).
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
// This API is used to query origin protection on an earlier version. EdgeOne comprehensively upgraded relevant APIs for origin protection on June 27, 2025. For details on the new version, see [DescribeOriginACL](https://intl.cloud.tencent.com/document/product/1552/120408?from_cn_redirect=1).
//
// 
//
// Note: Starting from June 27, 2025, the legacy version APIs will stop updating. New features will only be provided in the latest version APIs. To avoid data field conflicts when using legacy version APIs, it is recommended to migrate to the new version origin protection APIs as soon as possible.
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
// This API is used to query origin protection on an earlier version. EdgeOne comprehensively upgraded relevant APIs for origin protection on June 27, 2025. For details on the new version, see [DescribeOriginACL](https://intl.cloud.tencent.com/document/product/1552/120408?from_cn_redirect=1).
//
// 
//
// Note: Starting from June 27, 2025, the legacy version APIs will stop updating. New features will only be provided in the latest version APIs. To avoid data field conflicts when using legacy version APIs, it is recommended to migrate to the new version origin protection APIs as soon as possible.
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
// This API is used to query the time sequence traffic data of the monitoring category in L7. This API is to be discarded. Please use the API <a href="https://intl.cloud.tencent.com/document/product/1552/80648?from_cn_redirect=1">DescribeTimingL7AnalysisData</a>.
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
// This API is used to query the time sequence traffic data of the monitoring category in L7. This API is to be discarded. Please use the API <a href="https://intl.cloud.tencent.com/document/product/1552/80648?from_cn_redirect=1">DescribeTimingL7AnalysisData</a>.
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
// This API is used to query package information list with pagination support.
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribePlans(request *DescribePlansRequest) (response *DescribePlansResponse, err error) {
    return c.DescribePlansWithContext(context.Background(), request)
}

// DescribePlans
// This API is used to query package information list with pagination support.
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
// DescribePrefetchTasks is used to query the submission history and execution progress of preheating tasks. This interface can be used to query the tasks submitted by the CreatePrefetchTasks interface.
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
// DescribePrefetchTasks is used to query the submission history and execution progress of preheating tasks. This interface can be used to query the tasks submitted by the CreatePrefetchTasks interface.
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
// DescribePurgeTasks is used to query the submitted URL refreshing and directory refreshing records and execution progress. This interface can be used to query the tasks submitted by the CreatePurgeTasks API.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribePurgeTasks(request *DescribePurgeTasksRequest) (response *DescribePurgeTasksResponse, err error) {
    return c.DescribePurgeTasksWithContext(context.Background(), request)
}

// DescribePurgeTasks
// DescribePurgeTasks is used to query the submitted URL refreshing and directory refreshing records and execution progress. This interface can be used to query the tasks submitted by the CreatePurgeTasks API.
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
// This API is used to query the real-time log delivery task list.
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeRealtimeLogDeliveryTasks(request *DescribeRealtimeLogDeliveryTasksRequest) (response *DescribeRealtimeLogDeliveryTasksResponse, err error) {
    return c.DescribeRealtimeLogDeliveryTasksWithContext(context.Background(), request)
}

// DescribeRealtimeLogDeliveryTasks
// This API is used to query the real-time log delivery task list.
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
// This API is on an earlier version to query engine rules. EdgeOne has comprehensively upgraded relevant APIs of the rule engine on January 21, 2025. For details about the new version API to query layer-7 acceleration rules, see DescribeL7AccRules(https://intl.cloud.tencent.com/document/product/1552/115820?from_cn_redirect=1).
//
// <p style="color: red;">Note: Starting from January 21, 2025, the old version of the interface will stop updating and iteration. Subsequent new features will only be provided in the new version of the interface, and the original capabilities supported by the old version of the interface will not be affected. To avoid data field conflicts when using the old version of the interface, it is recommended that you migrate to the new version of the rule engine interface as soon as possible. </p>
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeRules(request *DescribeRulesRequest) (response *DescribeRulesResponse, err error) {
    return c.DescribeRulesWithContext(context.Background(), request)
}

// DescribeRules
// This API is on an earlier version to query engine rules. EdgeOne has comprehensively upgraded relevant APIs of the rule engine on January 21, 2025. For details about the new version API to query layer-7 acceleration rules, see DescribeL7AccRules(https://intl.cloud.tencent.com/document/product/1552/115820?from_cn_redirect=1).
//
// <p style="color: red;">Note: Starting from January 21, 2025, the old version of the interface will stop updating and iteration. Subsequent new features will only be provided in the new version of the interface, and the original capabilities supported by the old version of the interface will not be affected. To avoid data field conflicts when using the old version of the interface, it is recommended that you migrate to the new version of the rule engine interface as soon as possible. </p>
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
// This API is an older version. EdgeOne has fully upgraded the APIs related to the rule engine. For details, please refer to [RuleEngineAction](https://intl.cloud.tencent.com/document/product/1552/80721?from_cn_redirect=1#RuleEngineAction).
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeRulesSetting(request *DescribeRulesSettingRequest) (response *DescribeRulesSettingResponse, err error) {
    return c.DescribeRulesSettingWithContext(context.Background(), request)
}

// DescribeRulesSetting
// This API is an older version. EdgeOne has fully upgraded the APIs related to the rule engine. For details, please refer to [RuleEngineAction](https://intl.cloud.tencent.com/document/product/1552/80721?from_cn_redirect=1#RuleEngineAction).
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
// This API is used to query API resources under a site.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeSecurityAPIResource(request *DescribeSecurityAPIResourceRequest) (response *DescribeSecurityAPIResourceResponse, err error) {
    return c.DescribeSecurityAPIResourceWithContext(context.Background(), request)
}

// DescribeSecurityAPIResource
// This API is used to query API resources under a site.
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
// This API is used to query API services under a site.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeSecurityAPIService(request *DescribeSecurityAPIServiceRequest) (response *DescribeSecurityAPIServiceResponse, err error) {
    return c.DescribeSecurityAPIServiceWithContext(context.Background(), request)
}

// DescribeSecurityAPIService
// This API is used to query API services under a site.
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
// This API is used to query client authentication option configuration.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeSecurityClientAttester(request *DescribeSecurityClientAttesterRequest) (response *DescribeSecurityClientAttesterResponse, err error) {
    return c.DescribeSecurityClientAttesterWithContext(context.Background(), request)
}

// DescribeSecurityClientAttester
// This API is used to query client authentication option configuration.
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
// This API is used to query the configuration information of a security IP group, including the ID, name and content of the security IP group. The query result of this API only returns up to 2000 IPs or CIDR blocks for each IP group. If there is a very large IP group exceeding 2000 IPs or CIDR blocks, call DescribeSecurityIPGroupContent to perform a paging query.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeSecurityIPGroup(request *DescribeSecurityIPGroupRequest) (response *DescribeSecurityIPGroupResponse, err error) {
    return c.DescribeSecurityIPGroupWithContext(context.Background(), request)
}

// DescribeSecurityIPGroup
// This API is used to query the configuration information of a security IP group, including the ID, name and content of the security IP group. The query result of this API only returns up to 2000 IPs or CIDR blocks for each IP group. If there is a very large IP group exceeding 2000 IPs or CIDR blocks, call DescribeSecurityIPGroupContent to perform a paging query.
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
// This API is used to perform a paging query for the IP address list in a designated IP group. When the number of IP addresses in the group exceeds 2000, you can use this API to perform a paging query to obtain the complete IP address list.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeSecurityIPGroupContent(request *DescribeSecurityIPGroupContentRequest) (response *DescribeSecurityIPGroupContentResponse, err error) {
    return c.DescribeSecurityIPGroupContentWithContext(context.Background(), request)
}

// DescribeSecurityIPGroupContent
// This API is used to perform a paging query for the IP address list in a designated IP group. When the number of IP addresses in the group exceeds 2000, you can use this API to perform a paging query to obtain the complete IP address list.
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
// The API is deprecated and will be discontinued on June 30, 2024. Please use the API [DescribeSecurityIPGroup
//
// ](https://intl.cloud.tencent.com/document/product/1552/105866?from_cn_redirect=1).
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
// The API is deprecated and will be discontinued on June 30, 2024. Please use the API [DescribeSecurityIPGroup
//
// ](https://intl.cloud.tencent.com/document/product/1552/105866?from_cn_redirect=1).
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
// This API is used to query the web and security protection configurations.
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
// This API is used to query the web and security protection configurations.
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
// This API is used to query bindings of a policy template.
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
// This API is used to query bindings of a policy template.
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
// This API is used to query the list of L4 traffic data recorded over time.
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
// This API is used to query the list of L4 traffic data recorded over time.
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
// This API is used to query time series data for L7 domain name business.
//
// Create and bind policy Query instance Reset instance access password.
//
// This API is used to query data with a delay of about 10 minutes. It is recommended to pull data from at least 10 minutes before the current time.
//
// This API is used to return post-protection traffic request data by default. Users can query defended data in `Filters.mitigatedByWebSecurity`.
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
// This API is used to query time series data for L7 domain name business.
//
// Create and bind policy Query instance Reset instance access password.
//
// This API is used to query data with a delay of about 10 minutes. It is recommended to pull data from at least 10 minutes before the current time.
//
// This API is used to return post-protection traffic request data by default. Users can query defended data in `Filters.mitigatedByWebSecurity`.
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
// This API is used to query the time series traffic data of the L7 cache analysis. It will be deprecated. Use the <a href="https://intl.cloud.tencent.com/document/product/1552/80648?from_cn_redirect=1">DescribeTimingL7AnalysisData</a> API instead.
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
// This API is used to query the time series traffic data of the L7 cache analysis. It will be deprecated. Use the <a href="https://intl.cloud.tencent.com/document/product/1552/80648?from_cn_redirect=1">DescribeTimingL7AnalysisData</a> API instead.
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
// Create and bind policy Query instance Reset instance access password.
//
// This API is used to query data with a delay of about 10 minutes. It is recommended to pull data from at least 10 minutes before the current time.
//
// This API is used to return post-protection traffic request data by default. Users can query defended data in `Filters.mitigatedByWebSecurity`.
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
// Create and bind policy Query instance Reset instance access password.
//
// This API is used to query data with a delay of about 10 minutes. It is recommended to pull data from at least 10 minutes before the current time.
//
// This API is used to return post-protection traffic request data by default. Users can query defended data in `Filters.mitigatedByWebSecurity`.
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
// This API is used to query the top N data of the L7 cache analysis. It will be deprecated. Use the <a href="https://intl.cloud.tencent.com/document/product/1552/80646?from_cn_redirect=1"> DescribeTopL7AnalysisData</a> API instead.
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
// This API is used to query the top N data of the L7 cache analysis. It will be deprecated. Use the <a href="https://intl.cloud.tencent.com/document/product/1552/80646?from_cn_redirect=1"> DescribeTopL7AnalysisData</a> API instead.
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
// This API is used to query security policy configuration template details.
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
// This API is used to query security policy configuration template details.
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
// This API is used to query the security policy configuration template list.
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
// This API is used to query the security policy configuration template list.
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
// This API is used to query the results of site configuration import via API (ImportZoneConfig). This feature only supports the sites in the plans of the Standard Edition and the Enterprise Edition.
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
// This API is used to query the results of site configuration import via API (ImportZoneConfig). This feature only supports the sites in the plans of the Standard Edition and the Enterprise Edition.
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
// This API is an old version. EdgeOne has fully upgraded the APIs related to the rule engine. For details, please refer to [DescribeL7AccSetting](https://intl.cloud.tencent.com/document/product/1552/115819?from_cn_redirect=1).
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
// This API is an old version. EdgeOne has fully upgraded the APIs related to the rule engine. For details, please refer to [DescribeL7AccSetting](https://intl.cloud.tencent.com/document/product/1552/115819?from_cn_redirect=1).
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
// This API is used to query the information of sites that you have access to. You can filter sites based on different query criteria.
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
// This API is used to query the information of sites that you have access to. You can filter sites based on different query criteria.
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
// To stop billing for your EdgeOne plan, you can use this interface to terminate the billing plan.
//
// > Terminating a billing plan requires the following conditions:
//
//     1. The plan has expired (except for the Enterprise Edition Plan);
//
//     2. All sites under the plan have been either shut down or deleted.
//
// 
//
// > The site status can be queried through the [Query Site List](https://www.tencentcloud.com/zh/document/product/1145/50481) interface.
//
// A site can be deactivated by switching the site to a closed status through the [Switch Site Status](https://intl.cloud.tencent.com/document/product/1552/80707?from_cn_redirect=1) interface.
//
// A site can be deleted by using the [Delete Site](https://intl.cloud.tencent.com/document/product/1552/80717?from_cn_redirect=1) interface.
//
// error code that may be returned:
//  INVALIDPARAMETER_PLANNOTFOUND = "InvalidParameter.PlanNotFound"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
func (c *Client) DestroyPlan(request *DestroyPlanRequest) (response *DestroyPlanResponse, err error) {
    return c.DestroyPlanWithContext(context.Background(), request)
}

// DestroyPlan
// To stop billing for your EdgeOne plan, you can use this interface to terminate the billing plan.
//
// > Terminating a billing plan requires the following conditions:
//
//     1. The plan has expired (except for the Enterprise Edition Plan);
//
//     2. All sites under the plan have been either shut down or deleted.
//
// 
//
// > The site status can be queried through the [Query Site List](https://www.tencentcloud.com/zh/document/product/1145/50481) interface.
//
// A site can be deactivated by switching the site to a closed status through the [Switch Site Status](https://intl.cloud.tencent.com/document/product/1552/80707?from_cn_redirect=1) interface.
//
// A site can be deleted by using the [Delete Site](https://intl.cloud.tencent.com/document/product/1552/80717?from_cn_redirect=1) interface.
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
// This API is used to disable 'Origin Protection' of a site. Once disabled, resources related to it will no longer use only the origin ACLs provided by "origin protection" to request your origin, and stops sending update notifications on the origin ACLs.
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
// This API is used to disable 'Origin Protection' of a site. Once disabled, resources related to it will no longer use only the origin ACLs provided by "origin protection" to request your origin, and stops sending update notifications on the origin ACLs.
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
// This API is used to download L4 logs.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DownloadL4Logs(request *DownloadL4LogsRequest) (response *DownloadL4LogsResponse, err error) {
    return c.DownloadL4LogsWithContext(context.Background(), request)
}

// DownloadL4Logs
// This API is used to download L4 logs.
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
// This API is used to download L7 logs.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DownloadL7Logs(request *DownloadL7LogsRequest) (response *DownloadL7LogsResponse, err error) {
    return c.DownloadL7LogsWithContext(context.Background(), request)
}

// DownloadL7Logs
// This API is used to download L7 logs.
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
// This API is used to enable origin protection for a site for the first time. Enabled, EdgeOne will use specific origin IP ranges to backhaul traffic for L7 acceleration domains/L4 proxy instances. The maximum allowed number of L7 acceleration domains per submission is 200, and the maximum allowed number of L4 proxy instances is 100. Mixing L7 acceleration domains and L4 proxy instances in a single submission is supported, with a total maximum of 200 instances. To enable more than 200 resources, first enable the maximum quantity via specified resources, then enable the remaining resources via the ModifyOriginACL API. Subsequent addition of L7 acceleration domains/L4 proxy instances should be configured via the ModifyOriginACL API.
//
// 
//
// Create and bind policy Query instance Reset instance access password.
//
// -Call this API to deem as consent to the origin protection enablement special agreement (https://intl.cloud.tencent.com/document/product/1552/120141?from_cn_redirect=1);.
//
// -The origin IP range may change irregularly. tencent cloud EdgeOne (EdgeOne) will trigger notifications via message Center, SMS, or email 14 days, 7 days, 3 days, and 1 day before the change. To ensure you receive the change notification for the origin IP range, please ensure you have selected EdgeOne product services in the [tencent cloud message Center console](https://console.cloud.tencent.com/message) and configured the correct message recipient. For the setting method, refer to [message Subscription Management](https://intl.cloud.tencent.com/document/product/567/43476?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDOMAINS = "InvalidParameter.InvalidDomains"
//  INVALIDPARAMETER_INVALIDPROXIES = "InvalidParameter.InvalidProxies"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_UNSUPPORTEDPLAN = "OperationDenied.UnsupportedPlan"
//  OPERATIONDENIED_VERSIONCONTROLISGRAYING = "OperationDenied.VersionControlIsGraying"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) EnableOriginACL(request *EnableOriginACLRequest) (response *EnableOriginACLResponse, err error) {
    return c.EnableOriginACLWithContext(context.Background(), request)
}

// EnableOriginACL
// This API is used to enable origin protection for a site for the first time. Enabled, EdgeOne will use specific origin IP ranges to backhaul traffic for L7 acceleration domains/L4 proxy instances. The maximum allowed number of L7 acceleration domains per submission is 200, and the maximum allowed number of L4 proxy instances is 100. Mixing L7 acceleration domains and L4 proxy instances in a single submission is supported, with a total maximum of 200 instances. To enable more than 200 resources, first enable the maximum quantity via specified resources, then enable the remaining resources via the ModifyOriginACL API. Subsequent addition of L7 acceleration domains/L4 proxy instances should be configured via the ModifyOriginACL API.
//
// 
//
// Create and bind policy Query instance Reset instance access password.
//
// -Call this API to deem as consent to the origin protection enablement special agreement (https://intl.cloud.tencent.com/document/product/1552/120141?from_cn_redirect=1);.
//
// -The origin IP range may change irregularly. tencent cloud EdgeOne (EdgeOne) will trigger notifications via message Center, SMS, or email 14 days, 7 days, 3 days, and 1 day before the change. To ensure you receive the change notification for the origin IP range, please ensure you have selected EdgeOne product services in the [tencent cloud message Center console](https://console.cloud.tencent.com/message) and configured the correct message recipient. For the setting method, refer to [message Subscription Management](https://intl.cloud.tencent.com/document/product/567/43476?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDOMAINS = "InvalidParameter.InvalidDomains"
//  INVALIDPARAMETER_INVALIDPROXIES = "InvalidParameter.InvalidProxies"
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
// This API is used to export site configuration . The exported configuration is used for import via the API (ImportZoneConfig). This feature only supports the sites in the plans of the Standard Edition and the Enterprise Edition.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDOMAINS = "InvalidParameter.InvalidDomains"
//  INVALIDPARAMETER_INVALIDPROXIES = "InvalidParameter.InvalidProxies"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_UNSUPPORTEDPLAN = "OperationDenied.UnsupportedPlan"
//  OPERATIONDENIED_VERSIONCONTROLISGRAYING = "OperationDenied.VersionControlIsGraying"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) ExportZoneConfig(request *ExportZoneConfigRequest) (response *ExportZoneConfigResponse, err error) {
    return c.ExportZoneConfigWithContext(context.Background(), request)
}

// ExportZoneConfig
// This API is used to export site configuration . The exported configuration is used for import via the API (ImportZoneConfig). This feature only supports the sites in the plans of the Standard Edition and the Enterprise Edition.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDOMAINS = "InvalidParameter.InvalidDomains"
//  INVALIDPARAMETER_INVALIDPROXIES = "InvalidParameter.InvalidProxies"
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
// After the environment variables are set, they can be used in the function code. For details, see [Edge Functions Referencing Environment Variables](https://intl.cloud.tencent.com/document/product/1552/109151?from_cn_redirect=1#0151fd9a-8b0e-407b-ae37-54553a60ded6).
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
// After the environment variables are set, they can be used in the function code. For details, see [Edge Functions Referencing Environment Variables](https://intl.cloud.tencent.com/document/product/1552/109151?from_cn_redirect=1#0151fd9a-8b0e-407b-ae37-54553a60ded6).
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
// This API is used to verify ownership of the site.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) IdentifyZone(request *IdentifyZoneRequest) (response *IdentifyZoneResponse, err error) {
    return c.IdentifyZoneWithContext(context.Background(), request)
}

// IdentifyZone
// This API is used to verify ownership of the site.
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
// This API is used to quickly import site configuration files. After the import is initiated, the API will return the corresponding task ID (TaskId). Users need to use the site configuration import result query API (DescribeZoneConfigImportResult) to obtain the results of this import task. This feature only supports the sites in the plans of the Standard Edition and the Enterprise Edition.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) ImportZoneConfig(request *ImportZoneConfigRequest) (response *ImportZoneConfigResponse, err error) {
    return c.ImportZoneConfigWithContext(context.Background(), request)
}

// ImportZoneConfig
// This API is used to quickly import site configuration files. After the import is initiated, the API will return the corresponding task ID (TaskId). Users need to use the site configuration import result query API (DescribeZoneConfigImportResult) to obtain the results of this import task. This feature only supports the sites in the plans of the Standard Edition and the Enterprise Edition.
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
// When the number of sites bound to your plan, the number of rules under "Web Protection - Custom Rules - Precision Matching Policy", or the number of rules under "Web Protection - Rate Limiting - Precision Rate Limiting Module" reaches the plan's quota, you can use this interface to purchase additional quotas.
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
// When the number of sites bound to your plan, the number of rules under "Web Protection - Custom Rules - Precision Matching Policy", or the number of rules under "Web Protection - Rate Limiting - Precision Rate Limiting Module" reaches the plan's quota, you can use this interface to purchase additional quotas.
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
// This API is used to modify an accelerated domain name.
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
// This API is used to modify an accelerated domain name.
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
// This API is used to batch modify the status of accelerated domains.
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
// This API is used to batch modify the status of accelerated domains.
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
// This API is used to modify an alias domain name.
//
// The feature is only supported in the enterprise plan and is currently in closed beta testing. If you need to use it, [contact us](https://www.tencentcloud.com/contact-us).
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
// This API is used to modify an alias domain name.
//
// The feature is only supported in the enterprise plan and is currently in closed beta testing. If you need to use it, [contact us](https://www.tencentcloud.com/contact-us).
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
// This API is used to modify the status of an alias domain name.
//
// The feature is only supported in the enterprise plan and is currently in closed beta testing. If you need to use it, [Contact Us](https://www.tencentcloud.com/contact-us).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) ModifyAliasDomainStatus(request *ModifyAliasDomainStatusRequest) (response *ModifyAliasDomainStatusResponse, err error) {
    return c.ModifyAliasDomainStatusWithContext(context.Background(), request)
}

// ModifyAliasDomainStatus
// This API is used to modify the status of an alias domain name.
//
// The feature is only supported in the enterprise plan and is currently in closed beta testing. If you need to use it, [Contact Us](https://www.tencentcloud.com/contact-us).
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
// This API is on an earlier version. If you want to call it, please switch to the latest version. For details, see [ModifyL4Proxy
//
// ] (https://intl.cloud.tencent.com/document/product/1552/103411?from_cn_redirect=1).
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
// This API is on an earlier version. If you want to call it, please switch to the latest version. For details, see [ModifyL4Proxy
//
// ] (https://intl.cloud.tencent.com/document/product/1552/103411?from_cn_redirect=1).
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
// This API is on an earlier version. If you want to call it, please switch to the latest version. For details, see [ModifyL4ProxyRules] (https://intl.cloud.tencent.com/document/product/1552/103410?from_cn_redirect=1).
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
// This API is on an earlier version. If you want to call it, please switch to the latest version. For details, see [ModifyL4ProxyRules] (https://intl.cloud.tencent.com/document/product/1552/103410?from_cn_redirect=1).
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
// This API is on an earlier version. If you want to call it, please switch to the latest version. For details, see [ModifyL4ProxyRulesStatus
//
// ] (https://intl.cloud.tencent.com/document/product/1552/103409?from_cn_redirect=1).
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
// This API is on an earlier version. If you want to call it, please switch to the latest version. For details, see [ModifyL4ProxyRulesStatus
//
// ] (https://intl.cloud.tencent.com/document/product/1552/103409?from_cn_redirect=1).
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
// This API is on an earlier version. If you want to call it, please switch to the latest version. For details, see [ModifyL4ProxyStatus] (https://intl.cloud.tencent.com/document/product/1552/103408?from_cn_redirect=1).
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
// This API is on an earlier version. If you want to call it, please switch to the latest version. For details, see [ModifyL4ProxyStatus] (https://intl.cloud.tencent.com/document/product/1552/103408?from_cn_redirect=1).
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
// Modify content identifier, only description modification is supported. This feature is only open to the allowlist.
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
// Modify content identifier, only description modification is supported. This feature is only open to the allowlist.
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
// This API is used to modify site exclusive Anti-DDoS protection.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  OPERATIONDENIED_SECURITYLACKOFRESOURCES = "OperationDenied.SecurityLackOfResources"
func (c *Client) ModifyDDoSProtection(request *ModifyDDoSProtectionRequest) (response *ModifyDDoSProtectionResponse, err error) {
    return c.ModifyDDoSProtectionWithContext(context.Background(), request)
}

// ModifyDDoSProtection
// This API is used to modify site exclusive Anti-DDoS protection.
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
// This API is used to bulk modify DNS records.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  OPERATIONDENIED_SECURITYLACKOFRESOURCES = "OperationDenied.SecurityLackOfResources"
func (c *Client) ModifyDnsRecords(request *ModifyDnsRecordsRequest) (response *ModifyDnsRecordsResponse, err error) {
    return c.ModifyDnsRecordsWithContext(context.Background(), request)
}

// ModifyDnsRecords
// This API is used to bulk modify DNS records.
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
// You can batch modify the status of DNS records through this API, enabling and disabling records in bulk.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  OPERATIONDENIED_SECURITYLACKOFRESOURCES = "OperationDenied.SecurityLackOfResources"
func (c *Client) ModifyDnsRecordsStatus(request *ModifyDnsRecordsStatusRequest) (response *ModifyDnsRecordsStatusResponse, err error) {
    return c.ModifyDnsRecordsStatusWithContext(context.Background(), request)
}

// ModifyDnsRecordsStatus
// You can batch modify the status of DNS records through this API, enabling and disabling records in bulk.
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
// This API is used to modify a trigger rule for an edge function. It supports modifying rule conditions, execution functions, and description.
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
// This API is used to modify a trigger rule for an edge function. It supports modifying rule conditions, execution functions, and description.
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
// This API is used to configure the certificate of a site. You can use your own certificate or [apply for a free certificate](https://intl.cloud.tencent.com/document/product/1552/90437?from_cn_redirect=1).
//
// To use an external certificate, upload the certificate to [SSL Certificates Console](https://console.cloud.tencent.com/certoview) first, and then input the certificate ID in this API. For details, see [Deploying Own Certificates to EdgeOne Domains](https://intl.cloud.tencent.com/document/product/1552/88874?from_cn_redirect=1).
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
//  INVALIDPARAMETERVALUE_UPSTREAMCLIENTCERTINFOQUOTALIMIT = "InvalidParameterValue.UpstreamClientCertInfoQuotaLimit"
//  INVALIDPARAMETERVALUE_UPSTREAMVERIFYCUSTOMCACERTINFOQUOTALIMIT = "InvalidParameterValue.UpstreamVerifyCustomCACertInfoQuotaLimit"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_CERTIFICATEPRIVATEKEYISEMPTY = "OperationDenied.CertificatePrivateKeyIsEmpty"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
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
// This API is used to configure the certificate of a site. You can use your own certificate or [apply for a free certificate](https://intl.cloud.tencent.com/document/product/1552/90437?from_cn_redirect=1).
//
// To use an external certificate, upload the certificate to [SSL Certificates Console](https://console.cloud.tencent.com/certoview) first, and then input the certificate ID in this API. For details, see [Deploying Own Certificates to EdgeOne Domains](https://intl.cloud.tencent.com/document/product/1552/88874?from_cn_redirect=1).
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
//  INVALIDPARAMETERVALUE_UPSTREAMCLIENTCERTINFOQUOTALIMIT = "InvalidParameterValue.UpstreamClientCertInfoQuotaLimit"
//  INVALIDPARAMETERVALUE_UPSTREAMVERIFYCUSTOMCACERTINFOQUOTALIMIT = "InvalidParameterValue.UpstreamVerifyCustomCACertInfoQuotaLimit"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_CERTIFICATEPRIVATEKEYISEMPTY = "OperationDenied.CertificatePrivateKeyIsEmpty"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
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
// This API is used to modify the configuration of a Layer 4 proxy instance.
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
// This API is used to modify the configuration of a Layer 4 proxy instance.
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
// This API is used to start or stop Layer 4 proxy forwarding rules, supporting both individual and batch operation.
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
// This API is used to start or stop Layer 4 proxy forwarding rules, supporting both individual and batch operation.
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
// This API is used to enable or disable a Layer 4 proxy instance.
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
// This API is used to enable or disable a Layer 4 proxy instance.
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
// This API is used to modify rules in the [rule engine](https://intl.cloud.tencent.com/document/product/1552/70901?from_cn_redirect=1), supporting only one rule modification per request.
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
func (c *Client) ModifyL7AccRule(request *ModifyL7AccRuleRequest) (response *ModifyL7AccRuleResponse, err error) {
    return c.ModifyL7AccRuleWithContext(context.Background(), request)
}

// ModifyL7AccRule
// This API is used to modify rules in the [rule engine](https://intl.cloud.tencent.com/document/product/1552/70901?from_cn_redirect=1), supporting only one rule modification per request.
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
// This interface is used to modify the priority of the rule list in the [Rule Engine](https://intl.cloud.tencent.com/document/product/1552/70901?from_cn_redirect=1). This interface requires the complete rule ID list under the site ID to be passed in. The rule ID list can be obtained through the [Query Seven-Layer Acceleration Rules](https://intl.cloud.tencent.com/document/product/1552/115820?from_cn_redirect=1) interface. The final priority order will be adjusted to the order of the rule ID list, and will be executed from front to back.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDRULEENGINE = "InvalidParameter.InvalidRuleEngine"
//  INVALIDPARAMETER_INVALIDRULEENGINENOTFOUND = "InvalidParameter.InvalidRuleEngineNotFound"
func (c *Client) ModifyL7AccRulePriority(request *ModifyL7AccRulePriorityRequest) (response *ModifyL7AccRulePriorityResponse, err error) {
    return c.ModifyL7AccRulePriorityWithContext(context.Background(), request)
}

// ModifyL7AccRulePriority
// This interface is used to modify the priority of the rule list in the [Rule Engine](https://intl.cloud.tencent.com/document/product/1552/70901?from_cn_redirect=1). This interface requires the complete rule ID list under the site ID to be passed in. The rule ID list can be obtained through the [Query Seven-Layer Acceleration Rules](https://intl.cloud.tencent.com/document/product/1552/115820?from_cn_redirect=1) interface. The final priority order will be adjusted to the order of the rule ID list, and will be executed from front to back.
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
// This API is used to modify the global configuration of [Site Acceleration](https://intl.cloud.tencent.com/document/product/1552/96193?from_cn_redirect=1).
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
func (c *Client) ModifyL7AccSetting(request *ModifyL7AccSettingRequest) (response *ModifyL7AccSettingResponse, err error) {
    return c.ModifyL7AccSettingWithContext(context.Background(), request)
}

// ModifyL7AccSetting
// This API is used to modify the global configuration of [Site Acceleration](https://intl.cloud.tencent.com/document/product/1552/96193?from_cn_redirect=1).
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
// This API is used to modify multi-channel security acceleration gateway information, such as name, gateway ID, IP and port.
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
// This API is used to modify multi-channel security acceleration gateway information, such as name, gateway ID, IP and port.
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
// This API is used to modify the access lines of the multi-channel security acceleration gateway, including EdgeOne Layer-4 proxy lines and custom lines.
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
// This API is used to modify the access lines of the multi-channel security acceleration gateway, including EdgeOne Layer-4 proxy lines and custom lines.
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
// This API is used to modify the access key for the multi-channel security acceleration gateway.The access key is used by customers to sign requests for gateway access. The original key becomes invalid after modification.
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
// This API is used to modify the access key for the multi-channel security acceleration gateway.The access key is used by customers to sign requests for gateway access. The original key becomes invalid after modification.
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
// This API is used to update the status of a multi-channel security gateway.
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
// This API is used to update the status of a multi-channel security gateway.
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
// This API is used to enable or disable specific origin ACLs for L7 acceleration domain names or L4 proxy instances. A single submission supports up to 200 L7 acceleration domain names or 100 L4 proxy instances. Hybrid submissions of L7 acceleration domain names and L4 proxy instances are supported, with a maximum total number of instances of 200. If changes are needed for exceeding 200 instances, submit them in batches via this API.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDOMAINS = "InvalidParameter.InvalidDomains"
//  INVALIDPARAMETER_INVALIDPROXIES = "InvalidParameter.InvalidProxies"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_UNSUPPORTEDPLAN = "OperationDenied.UnsupportedPlan"
//  OPERATIONDENIED_UPDATEIPWHITELISTFIRST = "OperationDenied.UpdateIPWhitelistFirst"
func (c *Client) ModifyOriginACL(request *ModifyOriginACLRequest) (response *ModifyOriginACLResponse, err error) {
    return c.ModifyOriginACLWithContext(context.Background(), request)
}

// ModifyOriginACL
// This API is used to enable or disable specific origin ACLs for L7 acceleration domain names or L4 proxy instances. A single submission supports up to 200 L7 acceleration domain names or 100 L4 proxy instances. Hybrid submissions of L7 acceleration domain names and L4 proxy instances are supported, with a maximum total number of instances of 200. If changes are needed for exceeding 200 instances, submit them in batches via this API.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDOMAINS = "InvalidParameter.InvalidDomains"
//  INVALIDPARAMETER_INVALIDPROXIES = "InvalidParameter.InvalidProxies"
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
// This API is used to modify the configuration of an origin group. The original configuration will be overwritten. 
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
// This API is used to modify the configuration of an origin group. The original configuration will be overwritten. 
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
// This API is used to modify the real-time log delivery task configuration. This API has the following restrictions:<li>Does not support modifying the destination type of the real-time log delivery task (TaskType);</li><li>Does not support modifying the data delivery type (LogType);</li><li>Does not support modifying the data delivery area (Area);</li><li>Does not support modifying the detailed destination configuration, such as log set and log topic, when the destination of the original real-time log delivery task is Tencent Cloud CLS.</li>
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
// This API is used to modify the real-time log delivery task configuration. This API has the following restrictions:<li>Does not support modifying the destination type of the real-time log delivery task (TaskType);</li><li>Does not support modifying the data delivery type (LogType);</li><li>Does not support modifying the data delivery area (Area);</li><li>Does not support modifying the detailed destination configuration, such as log set and log topic, when the destination of the original real-time log delivery task is Tencent Cloud CLS.</li>
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
// This API is on an earlier version. EdgeOne has comprehensively upgraded the relevant APIs of the rule engine on January 21, 2025. For details about the new version of the API for modifying layer-7 acceleration rules, see ModifyL7AccRule(https://intl.cloud.tencent.com/document/product/1552/115818?from_cn_redirect=1).
//
// <p style="color: red;">Note: Starting from January 21, 2025, the old version of the interface will stop updating and iteration. Subsequent new features will only be provided in the new version of the interface, and the original capabilities supported by the old version of the interface will not be affected. To avoid data field conflicts when using the old version of the interface, it is recommended that you migrate to the new version of the rule engine interface as soon as possible. </p>
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
// This API is on an earlier version. EdgeOne has comprehensively upgraded the relevant APIs of the rule engine on January 21, 2025. For details about the new version of the API for modifying layer-7 acceleration rules, see ModifyL7AccRule(https://intl.cloud.tencent.com/document/product/1552/115818?from_cn_redirect=1).
//
// <p style="color: red;">Note: Starting from January 21, 2025, the old version of the interface will stop updating and iteration. Subsequent new features will only be provided in the new version of the interface, and the original capabilities supported by the old version of the interface will not be affected. To avoid data field conflicts when using the old version of the interface, it is recommended that you migrate to the new version of the rule engine interface as soon as possible. </p>
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
// This API is used to modify an API resource.
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
// This API is used to modify an API resource.
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
// This API is used to modify an API service.
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
// This API is used to modify an API service.
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
// This API is used to modify client authentication options.
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
// This API is used to modify client authentication options.
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
// This API is used to modify a security IP group.
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
// This API is used to modify a security IP group.
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
// This API is used to modify the web and bot security configurations.
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
// This API is used to modify the web and bot security configurations.
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
// This API is used to modify the security policy configuration template.
//
// error code that may be returned:
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
func (c *Client) ModifyWebSecurityTemplate(request *ModifyWebSecurityTemplateRequest) (response *ModifyWebSecurityTemplateResponse, err error) {
    return c.ModifyWebSecurityTemplateWithContext(context.Background(), request)
}

// ModifyWebSecurityTemplate
// This API is used to modify the security policy configuration template.
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
// This API is used to modify a site.
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
//  OPERATIONDENIED_NODOMAINACCESSZONEONLYSUPPORTMODIFYTYPE = "OperationDenied.NoDomainAccessZoneOnlySupportModifyType"
//  OPERATIONDENIED_PLANNOTSUPPORTMODIFYZONEAREA = "OperationDenied.PlanNotSupportModifyZoneArea"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
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
// This API is used to modify a site.
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
//  OPERATIONDENIED_NODOMAINACCESSZONEONLYSUPPORTMODIFYTYPE = "OperationDenied.NoDomainAccessZoneOnlySupportModifyType"
//  OPERATIONDENIED_PLANNOTSUPPORTMODIFYZONEAREA = "OperationDenied.PlanNotSupportModifyZoneArea"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
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
// This API is an older version. EdgeOne has fully upgraded the APIs related to the rule engine. For details, please refer to [ModifyL7AccSetting](https://intl.cloud.tencent.com/document/product/1552/115817?from_cn_redirect=1).
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
// This API is an older version. EdgeOne has fully upgraded the APIs related to the rule engine. For details, please refer to [ModifyL7AccSetting](https://intl.cloud.tencent.com/document/product/1552/115817?from_cn_redirect=1).
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
// This API is used to change the site status.
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
// This API is used to change the site status.
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
// This API is used to refresh keys for multi-channel security acceleration gateways. Customers access multi-channel security acceleration gateways based on integration key signatures. Each site has only one access key, which applies to all gateways under that site. After refreshing the key, the original key becomes invalid.
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
// This API is used to refresh keys for multi-channel security acceleration gateways. Customers access multi-channel security acceleration gateways based on integration key signatures. Each site has only one access key, which applies to all gateways under that site. After refreshing the key, the original key becomes invalid.
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
// When your plan needs to be extended, you can use this interface to renew it. Plan renewal is only supported for the Personal, Basic, and Standard Editions.
//
// > For cost details, refer to [Plan Fees](https://intl.cloud.tencent.com/document/product/1552/94158?from_cn_redirect=1).
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
// When your plan needs to be extended, you can use this interface to renew it. Plan renewal is only supported for the Personal, Basic, and Standard Editions.
//
// > For cost details, refer to [Plan Fees](https://intl.cloud.tencent.com/document/product/1552/94158?from_cn_redirect=1).
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
// > For differences between EdgeOne billing plans, refer to [Comparison of EdgeOne Plans](https://intl.cloud.tencent.com/document/product/1552/94165?from_cn_redirect=1).
//
// For EdgeOne plan upgrade rules and pricing details, refer to [EdgeOne Plan Upgrade Guide](https://intl.cloud.tencent.com/document/product/1552/95291?from_cn_redirect=1).
//
// If your plan needs to upgrade to the Enterprise Edition, [Contact Us](https://www.tencentcloud.com/contact-us).
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
// > For differences between EdgeOne billing plans, refer to [Comparison of EdgeOne Plans](https://intl.cloud.tencent.com/document/product/1552/94165?from_cn_redirect=1).
//
// For EdgeOne plan upgrade rules and pricing details, refer to [EdgeOne Plan Upgrade Guide](https://intl.cloud.tencent.com/document/product/1552/95291?from_cn_redirect=1).
//
// If your plan needs to upgrade to the Enterprise Edition, [Contact Us](https://www.tencentcloud.com/contact-us).
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
// This API is used to verify your ownership of a site or domain name. It's required in the CNAME access mode. After a site is verified, you don't need to verify the ownership again for domain names added to it in the future. For details, see [Ownership Verification](https://intl.cloud.tencent.com/document/product/1552/70789?from_cn_redirect=1).
//
// 
//
// For sites connected via the NS, you can query whether the NS is successfully switched through this API. For details, see [Modifying DNS Servers](https://intl.cloud.tencent.com/document/product/1552/90452?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) VerifyOwnership(request *VerifyOwnershipRequest) (response *VerifyOwnershipResponse, err error) {
    return c.VerifyOwnershipWithContext(context.Background(), request)
}

// VerifyOwnership
// This API is used to verify your ownership of a site or domain name. It's required in the CNAME access mode. After a site is verified, you don't need to verify the ownership again for domain names added to it in the future. For details, see [Ownership Verification](https://intl.cloud.tencent.com/document/product/1552/70789?from_cn_redirect=1).
//
// 
//
// For sites connected via the NS, you can query whether the NS is successfully switched through this API. For details, see [Modifying DNS Servers](https://intl.cloud.tencent.com/document/product/1552/90452?from_cn_redirect=1).
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
