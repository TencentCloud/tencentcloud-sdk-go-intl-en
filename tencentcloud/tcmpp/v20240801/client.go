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

package v20240801

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2024-08-01"

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


func NewAddTeamMemberRequest() (request *AddTeamMemberRequest) {
    request = &AddTeamMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "AddTeamMember")
    
    
    return
}

func NewAddTeamMemberResponse() (response *AddTeamMemberResponse) {
    response = &AddTeamMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddTeamMember
// This API is used to add a team member.
//
// error code that may be returned:
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  FAILEDOPERATION_UNABLEOPERATEADMINACCOUNT = "FailedOperation.UnableOperateAdminAccount"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  INVALIDPARAMETERVALUE_INVALIDROLEID = "InvalidParameterValue.InvalidRoleId"
//  INVALIDPARAMETERVALUE_INVALIDTEAMID = "InvalidParameterValue.InvalidTeamId"
//  INVALIDPARAMETERVALUE_INVALIDUSERID = "InvalidParameterValue.InvalidUserId"
func (c *Client) AddTeamMember(request *AddTeamMemberRequest) (response *AddTeamMemberResponse, err error) {
    return c.AddTeamMemberWithContext(context.Background(), request)
}

// AddTeamMember
// This API is used to add a team member.
//
// error code that may be returned:
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  FAILEDOPERATION_UNABLEOPERATEADMINACCOUNT = "FailedOperation.UnableOperateAdminAccount"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  INVALIDPARAMETERVALUE_INVALIDROLEID = "InvalidParameterValue.InvalidRoleId"
//  INVALIDPARAMETERVALUE_INVALIDTEAMID = "InvalidParameterValue.InvalidTeamId"
//  INVALIDPARAMETERVALUE_INVALIDUSERID = "InvalidParameterValue.InvalidUserId"
func (c *Client) AddTeamMemberWithContext(ctx context.Context, request *AddTeamMemberRequest) (response *AddTeamMemberResponse, err error) {
    if request == nil {
        request = NewAddTeamMemberRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddTeamMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddTeamMemberResponse()
    err = c.Send(request, response)
    return
}

func NewCheckGlobalDomainRequest() (request *CheckGlobalDomainRequest) {
    request = &CheckGlobalDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "CheckGlobalDomain")
    
    
    return
}

func NewCheckGlobalDomainResponse() (response *CheckGlobalDomainResponse) {
    response = &CheckGlobalDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckGlobalDomain
// This API is used to query if the domain is in the allowlist or blocklist
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXISTBLACKDOMAIN = "FailedOperation.ExistBlackDomain"
//  FAILEDOPERATION_EXISTWHITEDOMAIN = "FailedOperation.ExistWhiteDomain"
//  FAILEDOPERATION_INVALIDREQUESTENTERPRISEINFO = "FailedOperation.InvalidRequestEnterpriseInfo"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_NOACCESSPERMISSION = "FailedOperation.NoAccessPermission"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CheckGlobalDomain(request *CheckGlobalDomainRequest) (response *CheckGlobalDomainResponse, err error) {
    return c.CheckGlobalDomainWithContext(context.Background(), request)
}

// CheckGlobalDomain
// This API is used to query if the domain is in the allowlist or blocklist
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXISTBLACKDOMAIN = "FailedOperation.ExistBlackDomain"
//  FAILEDOPERATION_EXISTWHITEDOMAIN = "FailedOperation.ExistWhiteDomain"
//  FAILEDOPERATION_INVALIDREQUESTENTERPRISEINFO = "FailedOperation.InvalidRequestEnterpriseInfo"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_NOACCESSPERMISSION = "FailedOperation.NoAccessPermission"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CheckGlobalDomainWithContext(ctx context.Context, request *CheckGlobalDomainRequest) (response *CheckGlobalDomainResponse, err error) {
    if request == nil {
        request = NewCheckGlobalDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckGlobalDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckGlobalDomainResponse()
    err = c.Send(request, response)
    return
}

func NewConfigureMNPPreviewRequest() (request *ConfigureMNPPreviewRequest) {
    request = &ConfigureMNPPreviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "ConfigureMNPPreview")
    
    
    return
}

func NewConfigureMNPPreviewResponse() (response *ConfigureMNPPreviewResponse) {
    response = &ConfigureMNPPreviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ConfigureMNPPreview
// This API is used to configure the preview version of a mini program.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_SHOWCASEVERSIONALREADYEXIST = "FailedOperation.ShowcaseVersionAlreadyExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) ConfigureMNPPreview(request *ConfigureMNPPreviewRequest) (response *ConfigureMNPPreviewResponse, err error) {
    return c.ConfigureMNPPreviewWithContext(context.Background(), request)
}

// ConfigureMNPPreview
// This API is used to configure the preview version of a mini program.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_SHOWCASEVERSIONALREADYEXIST = "FailedOperation.ShowcaseVersionAlreadyExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) ConfigureMNPPreviewWithContext(ctx context.Context, request *ConfigureMNPPreviewRequest) (response *ConfigureMNPPreviewResponse, err error) {
    if request == nil {
        request = NewConfigureMNPPreviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ConfigureMNPPreview require credential")
    }

    request.SetContext(ctx)
    
    response = NewConfigureMNPPreviewResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApplicationRequest() (request *CreateApplicationRequest) {
    request = &CreateApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "CreateApplication")
    
    
    return
}

func NewCreateApplicationResponse() (response *CreateApplicationResponse) {
    response = &CreateApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApplication
// This API is used to create an application.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPLICATIONNAMEALREADYEXISTED = "FailedOperation.ApplicationNameAlreadyExisted"
//  FAILEDOPERATION_APPLICATIONNUMBEREXCEEDLIMIT = "FailedOperation.ApplicationNumberExceedLimit"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_APPPACKAGENAMENOTCONFIG = "InvalidParameterValue.AppPackageNameNotConfig"
//  INVALIDPARAMETERVALUE_INVALIDAPPPACKAGENAMELAYOUT = "InvalidParameterValue.InvalidAppPackageNameLayout"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONINTRO = "InvalidParameterValue.InvalidApplicationIntro"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONNAME = "InvalidParameterValue.InvalidApplicationName"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONREMARK = "InvalidParameterValue.InvalidApplicationRemark"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) CreateApplication(request *CreateApplicationRequest) (response *CreateApplicationResponse, err error) {
    return c.CreateApplicationWithContext(context.Background(), request)
}

// CreateApplication
// This API is used to create an application.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPLICATIONNAMEALREADYEXISTED = "FailedOperation.ApplicationNameAlreadyExisted"
//  FAILEDOPERATION_APPLICATIONNUMBEREXCEEDLIMIT = "FailedOperation.ApplicationNumberExceedLimit"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_APPPACKAGENAMENOTCONFIG = "InvalidParameterValue.AppPackageNameNotConfig"
//  INVALIDPARAMETERVALUE_INVALIDAPPPACKAGENAMELAYOUT = "InvalidParameterValue.InvalidAppPackageNameLayout"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONINTRO = "InvalidParameterValue.InvalidApplicationIntro"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONNAME = "InvalidParameterValue.InvalidApplicationName"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONREMARK = "InvalidParameterValue.InvalidApplicationRemark"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) CreateApplicationWithContext(ctx context.Context, request *CreateApplicationRequest) (response *CreateApplicationResponse, err error) {
    if request == nil {
        request = NewCreateApplicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApplicationSensitiveAPIRequest() (request *CreateApplicationSensitiveAPIRequest) {
    request = &CreateApplicationSensitiveAPIRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "CreateApplicationSensitiveAPI")
    
    
    return
}

func NewCreateApplicationSensitiveAPIResponse() (response *CreateApplicationSensitiveAPIResponse) {
    response = &CreateApplicationSensitiveAPIResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApplicationSensitiveAPI
// This API is used to create a sensitive API of an application.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXISTSENSITIVEAPI = "FailedOperation.ExistSensitiveAPI"
//  FAILEDOPERATION_GETOPERATERESOURCEFAILED = "FailedOperation.GetOperateResourceFailed"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONID = "InvalidParameterValue.InvalidApplicationId"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) CreateApplicationSensitiveAPI(request *CreateApplicationSensitiveAPIRequest) (response *CreateApplicationSensitiveAPIResponse, err error) {
    return c.CreateApplicationSensitiveAPIWithContext(context.Background(), request)
}

// CreateApplicationSensitiveAPI
// This API is used to create a sensitive API of an application.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXISTSENSITIVEAPI = "FailedOperation.ExistSensitiveAPI"
//  FAILEDOPERATION_GETOPERATERESOURCEFAILED = "FailedOperation.GetOperateResourceFailed"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONID = "InvalidParameterValue.InvalidApplicationId"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) CreateApplicationSensitiveAPIWithContext(ctx context.Context, request *CreateApplicationSensitiveAPIRequest) (response *CreateApplicationSensitiveAPIResponse, err error) {
    if request == nil {
        request = NewCreateApplicationSensitiveAPIRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApplicationSensitiveAPI require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApplicationSensitiveAPIResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConsoleMNPVersionCompileTaskRequest() (request *CreateConsoleMNPVersionCompileTaskRequest) {
    request = &CreateConsoleMNPVersionCompileTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "CreateConsoleMNPVersionCompileTask")
    
    
    return
}

func NewCreateConsoleMNPVersionCompileTaskResponse() (response *CreateConsoleMNPVersionCompileTaskResponse) {
    response = &CreateConsoleMNPVersionCompileTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateConsoleMNPVersionCompileTask
// This API is used to add a new mini program version
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDMINIAPPID = "InvalidParameterValue.InvalidMiniAppId"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
func (c *Client) CreateConsoleMNPVersionCompileTask(request *CreateConsoleMNPVersionCompileTaskRequest) (response *CreateConsoleMNPVersionCompileTaskResponse, err error) {
    return c.CreateConsoleMNPVersionCompileTaskWithContext(context.Background(), request)
}

// CreateConsoleMNPVersionCompileTask
// This API is used to add a new mini program version
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDMINIAPPID = "InvalidParameterValue.InvalidMiniAppId"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
func (c *Client) CreateConsoleMNPVersionCompileTaskWithContext(ctx context.Context, request *CreateConsoleMNPVersionCompileTaskRequest) (response *CreateConsoleMNPVersionCompileTaskResponse, err error) {
    if request == nil {
        request = NewCreateConsoleMNPVersionCompileTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConsoleMNPVersionCompileTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConsoleMNPVersionCompileTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDomainRequest() (request *CreateDomainRequest) {
    request = &CreateDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "CreateDomain")
    
    
    return
}

func NewCreateDomainResponse() (response *CreateDomainResponse) {
    response = &CreateDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDomain
// This API is used to create a mini program service domain
//
// error code that may be returned:
//  FAILEDOPERATION_EXISTREPEATDOMAINURL = "FailedOperation.ExistRepeatDomainUrl"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
func (c *Client) CreateDomain(request *CreateDomainRequest) (response *CreateDomainResponse, err error) {
    return c.CreateDomainWithContext(context.Background(), request)
}

// CreateDomain
// This API is used to create a mini program service domain
//
// error code that may be returned:
//  FAILEDOPERATION_EXISTREPEATDOMAINURL = "FailedOperation.ExistRepeatDomainUrl"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
func (c *Client) CreateDomainWithContext(ctx context.Context, request *CreateDomainRequest) (response *CreateDomainResponse, err error) {
    if request == nil {
        request = NewCreateDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGlobalDomainRequest() (request *CreateGlobalDomainRequest) {
    request = &CreateGlobalDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "CreateGlobalDomain")
    
    
    return
}

func NewCreateGlobalDomainResponse() (response *CreateGlobalDomainResponse) {
    response = &CreateGlobalDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateGlobalDomain
// This API is used to add domains to allowlist or blocklist
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXISTBLACKDOMAIN = "FailedOperation.ExistBlackDomain"
//  FAILEDOPERATION_EXISTWHITEDOMAIN = "FailedOperation.ExistWhiteDomain"
//  FAILEDOPERATION_INVALIDREQUESTENTERPRISEINFO = "FailedOperation.InvalidRequestEnterpriseInfo"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_NOACCESSPERMISSION = "FailedOperation.NoAccessPermission"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_USERTEAMRELATIONNOTEXIST = "InvalidParameterValue.UserTeamRelationNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateGlobalDomain(request *CreateGlobalDomainRequest) (response *CreateGlobalDomainResponse, err error) {
    return c.CreateGlobalDomainWithContext(context.Background(), request)
}

// CreateGlobalDomain
// This API is used to add domains to allowlist or blocklist
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXISTBLACKDOMAIN = "FailedOperation.ExistBlackDomain"
//  FAILEDOPERATION_EXISTWHITEDOMAIN = "FailedOperation.ExistWhiteDomain"
//  FAILEDOPERATION_INVALIDREQUESTENTERPRISEINFO = "FailedOperation.InvalidRequestEnterpriseInfo"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_NOACCESSPERMISSION = "FailedOperation.NoAccessPermission"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_USERTEAMRELATIONNOTEXIST = "InvalidParameterValue.UserTeamRelationNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateGlobalDomainWithContext(ctx context.Context, request *CreateGlobalDomainRequest) (response *CreateGlobalDomainResponse, err error) {
    if request == nil {
        request = NewCreateGlobalDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGlobalDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateGlobalDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGlobalDomainACLRequest() (request *CreateGlobalDomainACLRequest) {
    request = &CreateGlobalDomainACLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "CreateGlobalDomainACL")
    
    
    return
}

func NewCreateGlobalDomainACLResponse() (response *CreateGlobalDomainACLResponse) {
    response = &CreateGlobalDomainACLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateGlobalDomainACL
// This API is used to create a global domain allowlist or blocklist.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETOPERATERESOURCEFAILED = "FailedOperation.GetOperateResourceFailed"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) CreateGlobalDomainACL(request *CreateGlobalDomainACLRequest) (response *CreateGlobalDomainACLResponse, err error) {
    return c.CreateGlobalDomainACLWithContext(context.Background(), request)
}

// CreateGlobalDomainACL
// This API is used to create a global domain allowlist or blocklist.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETOPERATERESOURCEFAILED = "FailedOperation.GetOperateResourceFailed"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) CreateGlobalDomainACLWithContext(ctx context.Context, request *CreateGlobalDomainACLRequest) (response *CreateGlobalDomainACLResponse, err error) {
    if request == nil {
        request = NewCreateGlobalDomainACLRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGlobalDomainACL require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateGlobalDomainACLResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMNPRequest() (request *CreateMNPRequest) {
    request = &CreateMNPRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "CreateMNP")
    
    
    return
}

func NewCreateMNPResponse() (response *CreateMNPResponse) {
    response = &CreateMNPResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateMNP
// This API is used to create a mini program
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MNPNUMBEREXCEEDLIMIT = "FailedOperation.MNPNumberExceedLimit"
//  FAILEDOPERATION_OPERATIONSTEAMNOMINIPROGRAMPERMISSION = "FailedOperation.OperationsTeamNoMiniProgramPermission"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMNPICON = "InvalidParameterValue.InvalidMNPIcon"
//  INVALIDPARAMETERVALUE_INVALIDMNPINTRO = "InvalidParameterValue.InvalidMNPIntro"
//  INVALIDPARAMETERVALUE_INVALIDMNPNAME = "InvalidParameterValue.InvalidMNPName"
//  INVALIDPARAMETERVALUE_INVALIDMNPTYPE = "InvalidParameterValue.InvalidMNPType"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  INVALIDPARAMETERVALUE_MNPTYPENUMBEREXCEEDLIMIT = "InvalidParameterValue.MNPTypeNumberExceedLimit"
//  INVALIDPARAMETERVALUE_MINIPROGRAMNAMEALREADYEXIST = "InvalidParameterValue.MiniProgramNameAlreadyExist"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateMNP(request *CreateMNPRequest) (response *CreateMNPResponse, err error) {
    return c.CreateMNPWithContext(context.Background(), request)
}

// CreateMNP
// This API is used to create a mini program
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MNPNUMBEREXCEEDLIMIT = "FailedOperation.MNPNumberExceedLimit"
//  FAILEDOPERATION_OPERATIONSTEAMNOMINIPROGRAMPERMISSION = "FailedOperation.OperationsTeamNoMiniProgramPermission"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMNPICON = "InvalidParameterValue.InvalidMNPIcon"
//  INVALIDPARAMETERVALUE_INVALIDMNPINTRO = "InvalidParameterValue.InvalidMNPIntro"
//  INVALIDPARAMETERVALUE_INVALIDMNPNAME = "InvalidParameterValue.InvalidMNPName"
//  INVALIDPARAMETERVALUE_INVALIDMNPTYPE = "InvalidParameterValue.InvalidMNPType"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  INVALIDPARAMETERVALUE_MNPTYPENUMBEREXCEEDLIMIT = "InvalidParameterValue.MNPTypeNumberExceedLimit"
//  INVALIDPARAMETERVALUE_MINIPROGRAMNAMEALREADYEXIST = "InvalidParameterValue.MiniProgramNameAlreadyExist"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateMNPWithContext(ctx context.Context, request *CreateMNPRequest) (response *CreateMNPResponse, err error) {
    if request == nil {
        request = NewCreateMNPRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMNP require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMNPResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMNPApprovalRequest() (request *CreateMNPApprovalRequest) {
    request = &CreateMNPApprovalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "CreateMNPApproval")
    
    
    return
}

func NewCreateMNPApprovalResponse() (response *CreateMNPApprovalResponse) {
    response = &CreateMNPApprovalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateMNPApproval
// This API is used to create a mini program approval request.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXISTMNPAPPROVALRECORD = "FailedOperation.ExistMNPApprovalRecord"
//  FAILEDOPERATION_MNPTEAMASSOCIATEDAPPLICATIONTEAMNOTCREATEAPPLICATION = "FailedOperation.MNPTeamAssociatedApplicationTeamNotCreateApplication"
//  FAILEDOPERATION_MNPTEAMNOTASSOCIATEDAPPLICATIONTEAM = "FailedOperation.MNPTeamNotAssociatedApplicationTeam"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDAPPLYACTION = "InvalidParameterValue.InvalidApplyAction"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  INVALIDPARAMETERVALUE_MNPAPPROVALSTATUSERROR = "InvalidParameterValue.MNPApprovalStatusError"
//  INVALIDPARAMETERVALUE_MNPVERSIONISNOTPLATFORMPHASE = "InvalidParameterValue.MNPVersionIsNotPlatformPhase"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateMNPApproval(request *CreateMNPApprovalRequest) (response *CreateMNPApprovalResponse, err error) {
    return c.CreateMNPApprovalWithContext(context.Background(), request)
}

// CreateMNPApproval
// This API is used to create a mini program approval request.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXISTMNPAPPROVALRECORD = "FailedOperation.ExistMNPApprovalRecord"
//  FAILEDOPERATION_MNPTEAMASSOCIATEDAPPLICATIONTEAMNOTCREATEAPPLICATION = "FailedOperation.MNPTeamAssociatedApplicationTeamNotCreateApplication"
//  FAILEDOPERATION_MNPTEAMNOTASSOCIATEDAPPLICATIONTEAM = "FailedOperation.MNPTeamNotAssociatedApplicationTeam"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDAPPLYACTION = "InvalidParameterValue.InvalidApplyAction"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  INVALIDPARAMETERVALUE_MNPAPPROVALSTATUSERROR = "InvalidParameterValue.MNPApprovalStatusError"
//  INVALIDPARAMETERVALUE_MNPVERSIONISNOTPLATFORMPHASE = "InvalidParameterValue.MNPVersionIsNotPlatformPhase"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateMNPApprovalWithContext(ctx context.Context, request *CreateMNPApprovalRequest) (response *CreateMNPApprovalResponse, err error) {
    if request == nil {
        request = NewCreateMNPApprovalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMNPApproval require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMNPApprovalResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMNPDomainACLRequest() (request *CreateMNPDomainACLRequest) {
    request = &CreateMNPDomainACLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "CreateMNPDomainACL")
    
    
    return
}

func NewCreateMNPDomainACLResponse() (response *CreateMNPDomainACLResponse) {
    response = &CreateMNPDomainACLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateMNPDomainACL
// This API is used to add a domain name to the allowlist/blocklist of a mini program.
//
// error code that may be returned:
//  FAILEDOPERATION_EXISTREPEATDOMAINURL = "FailedOperation.ExistRepeatDomainUrl"
//  FAILEDOPERATION_GETOPERATERESOURCEFAILED = "FailedOperation.GetOperateResourceFailed"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDTEAMID = "InvalidParameterValue.InvalidTeamId"
func (c *Client) CreateMNPDomainACL(request *CreateMNPDomainACLRequest) (response *CreateMNPDomainACLResponse, err error) {
    return c.CreateMNPDomainACLWithContext(context.Background(), request)
}

// CreateMNPDomainACL
// This API is used to add a domain name to the allowlist/blocklist of a mini program.
//
// error code that may be returned:
//  FAILEDOPERATION_EXISTREPEATDOMAINURL = "FailedOperation.ExistRepeatDomainUrl"
//  FAILEDOPERATION_GETOPERATERESOURCEFAILED = "FailedOperation.GetOperateResourceFailed"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDTEAMID = "InvalidParameterValue.InvalidTeamId"
func (c *Client) CreateMNPDomainACLWithContext(ctx context.Context, request *CreateMNPDomainACLRequest) (response *CreateMNPDomainACLResponse, err error) {
    if request == nil {
        request = NewCreateMNPDomainACLRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMNPDomainACL require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMNPDomainACLResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMNPSensitiveAPIPermissionApprovalRequest() (request *CreateMNPSensitiveAPIPermissionApprovalRequest) {
    request = &CreateMNPSensitiveAPIPermissionApprovalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "CreateMNPSensitiveAPIPermissionApproval")
    
    
    return
}

func NewCreateMNPSensitiveAPIPermissionApprovalResponse() (response *CreateMNPSensitiveAPIPermissionApprovalResponse) {
    response = &CreateMNPSensitiveAPIPermissionApprovalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateMNPSensitiveAPIPermissionApproval
// This API is used to create a permission request to allow a mini program calling sensitive APIs.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETOPERATERESOURCEFAILED = "FailedOperation.GetOperateResourceFailed"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) CreateMNPSensitiveAPIPermissionApproval(request *CreateMNPSensitiveAPIPermissionApprovalRequest) (response *CreateMNPSensitiveAPIPermissionApprovalResponse, err error) {
    return c.CreateMNPSensitiveAPIPermissionApprovalWithContext(context.Background(), request)
}

// CreateMNPSensitiveAPIPermissionApproval
// This API is used to create a permission request to allow a mini program calling sensitive APIs.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETOPERATERESOURCEFAILED = "FailedOperation.GetOperateResourceFailed"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) CreateMNPSensitiveAPIPermissionApprovalWithContext(ctx context.Context, request *CreateMNPSensitiveAPIPermissionApprovalRequest) (response *CreateMNPSensitiveAPIPermissionApprovalResponse, err error) {
    if request == nil {
        request = NewCreateMNPSensitiveAPIPermissionApprovalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMNPSensitiveAPIPermissionApproval require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMNPSensitiveAPIPermissionApprovalResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMNPVersionRequest() (request *CreateMNPVersionRequest) {
    request = &CreateMNPVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "CreateMNPVersion")
    
    
    return
}

func NewCreateMNPVersionResponse() (response *CreateMNPVersionResponse) {
    response = &CreateMNPVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateMNPVersion
// This API is used to create a mini program version.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FILEDOWNLOADFAILED = "FailedOperation.FileDownloadFailed"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_EXISTMISMATCHPARAMETERTYPE = "InvalidParameterValue.ExistMismatchParameterType"
//  INVALIDPARAMETERVALUE_INVALIDMNPID = "InvalidParameterValue.InvalidMNPId"
//  INVALIDPARAMETERVALUE_INVALIDMNPVERSION = "InvalidParameterValue.InvalidMNPVersion"
//  INVALIDPARAMETERVALUE_INVALIDMNPVERSIONINTRO = "InvalidParameterValue.InvalidMNPVersionIntro"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  INVALIDPARAMETERVALUE_MNPFILESIZEXCEEDLIMIT = "InvalidParameterValue.MNPFileSizExceedLimit"
func (c *Client) CreateMNPVersion(request *CreateMNPVersionRequest) (response *CreateMNPVersionResponse, err error) {
    return c.CreateMNPVersionWithContext(context.Background(), request)
}

// CreateMNPVersion
// This API is used to create a mini program version.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FILEDOWNLOADFAILED = "FailedOperation.FileDownloadFailed"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_EXISTMISMATCHPARAMETERTYPE = "InvalidParameterValue.ExistMismatchParameterType"
//  INVALIDPARAMETERVALUE_INVALIDMNPID = "InvalidParameterValue.InvalidMNPId"
//  INVALIDPARAMETERVALUE_INVALIDMNPVERSION = "InvalidParameterValue.InvalidMNPVersion"
//  INVALIDPARAMETERVALUE_INVALIDMNPVERSIONINTRO = "InvalidParameterValue.InvalidMNPVersionIntro"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  INVALIDPARAMETERVALUE_MNPFILESIZEXCEEDLIMIT = "InvalidParameterValue.MNPFileSizExceedLimit"
func (c *Client) CreateMNPVersionWithContext(ctx context.Context, request *CreateMNPVersionRequest) (response *CreateMNPVersionResponse, err error) {
    if request == nil {
        request = NewCreateMNPVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMNPVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMNPVersionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOnlineApplyRequest() (request *CreateOnlineApplyRequest) {
    request = &CreateOnlineApplyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "CreateOnlineApply")
    
    
    return
}

func NewCreateOnlineApplyResponse() (response *CreateOnlineApplyResponse) {
    response = &CreateOnlineApplyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateOnlineApply
// This API is used to release the mini program
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateOnlineApply(request *CreateOnlineApplyRequest) (response *CreateOnlineApplyResponse, err error) {
    return c.CreateOnlineApplyWithContext(context.Background(), request)
}

// CreateOnlineApply
// This API is used to release the mini program
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateOnlineApplyWithContext(ctx context.Context, request *CreateOnlineApplyRequest) (response *CreateOnlineApplyResponse, err error) {
    if request == nil {
        request = NewCreateOnlineApplyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOnlineApply require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOnlineApplyResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePlatformAuditRequest() (request *CreatePlatformAuditRequest) {
    request = &CreatePlatformAuditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "CreatePlatformAudit")
    
    
    return
}

func NewCreatePlatformAuditResponse() (response *CreatePlatformAuditResponse) {
    response = &CreatePlatformAuditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePlatformAudit
// This API is used to submit mini program version for approval
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreatePlatformAudit(request *CreatePlatformAuditRequest) (response *CreatePlatformAuditResponse, err error) {
    return c.CreatePlatformAuditWithContext(context.Background(), request)
}

// CreatePlatformAudit
// This API is used to submit mini program version for approval
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreatePlatformAuditWithContext(ctx context.Context, request *CreatePlatformAuditRequest) (response *CreatePlatformAuditResponse, err error) {
    if request == nil {
        request = NewCreatePlatformAuditRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePlatformAudit require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePlatformAuditResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePresetKeyRequest() (request *CreatePresetKeyRequest) {
    request = &CreatePresetKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "CreatePresetKey")
    
    
    return
}

func NewCreatePresetKeyResponse() (response *CreatePresetKeyResponse) {
    response = &CreatePresetKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePresetKey
// This API is used to obtain the encryption key
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
func (c *Client) CreatePresetKey(request *CreatePresetKeyRequest) (response *CreatePresetKeyResponse, err error) {
    return c.CreatePresetKeyWithContext(context.Background(), request)
}

// CreatePresetKey
// This API is used to obtain the encryption key
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
func (c *Client) CreatePresetKeyWithContext(ctx context.Context, request *CreatePresetKeyRequest) (response *CreatePresetKeyResponse, err error) {
    if request == nil {
        request = NewCreatePresetKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePresetKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePresetKeyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSensitiveAPIRequest() (request *CreateSensitiveAPIRequest) {
    request = &CreateSensitiveAPIRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "CreateSensitiveAPI")
    
    
    return
}

func NewCreateSensitiveAPIResponse() (response *CreateSensitiveAPIResponse) {
    response = &CreateSensitiveAPIResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSensitiveAPI
// This API is used to add a sensitive API
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXISTSENSITIVEAPI = "FailedOperation.ExistSensitiveAPI"
//  FAILEDOPERATION_GETOPERATERESOURCEFAILED = "FailedOperation.GetOperateResourceFailed"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateSensitiveAPI(request *CreateSensitiveAPIRequest) (response *CreateSensitiveAPIResponse, err error) {
    return c.CreateSensitiveAPIWithContext(context.Background(), request)
}

// CreateSensitiveAPI
// This API is used to add a sensitive API
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXISTSENSITIVEAPI = "FailedOperation.ExistSensitiveAPI"
//  FAILEDOPERATION_GETOPERATERESOURCEFAILED = "FailedOperation.GetOperateResourceFailed"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateSensitiveAPIWithContext(ctx context.Context, request *CreateSensitiveAPIRequest) (response *CreateSensitiveAPIResponse, err error) {
    if request == nil {
        request = NewCreateSensitiveAPIRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSensitiveAPI require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSensitiveAPIResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSensitiveApiApplyRequest() (request *CreateSensitiveApiApplyRequest) {
    request = &CreateSensitiveApiApplyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "CreateSensitiveApiApply")
    
    
    return
}

func NewCreateSensitiveApiApplyResponse() (response *CreateSensitiveApiApplyResponse) {
    response = &CreateSensitiveApiApplyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSensitiveApiApply
// This API is used to apply for sensitive API permissions
//
// error code that may be returned:
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
func (c *Client) CreateSensitiveApiApply(request *CreateSensitiveApiApplyRequest) (response *CreateSensitiveApiApplyResponse, err error) {
    return c.CreateSensitiveApiApplyWithContext(context.Background(), request)
}

// CreateSensitiveApiApply
// This API is used to apply for sensitive API permissions
//
// error code that may be returned:
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
func (c *Client) CreateSensitiveApiApplyWithContext(ctx context.Context, request *CreateSensitiveApiApplyRequest) (response *CreateSensitiveApiApplyResponse, err error) {
    if request == nil {
        request = NewCreateSensitiveApiApplyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSensitiveApiApply require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSensitiveApiApplyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTeamRequest() (request *CreateTeamRequest) {
    request = &CreateTeamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "CreateTeam")
    
    
    return
}

func NewCreateTeamResponse() (response *CreateTeamResponse) {
    response = &CreateTeamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTeam
// This API is used to create a team.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETOPERATERESOURCEFAILED = "FailedOperation.GetOperateResourceFailed"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EXISTTEAMNAME = "InvalidParameterValue.ExistTeamName"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) CreateTeam(request *CreateTeamRequest) (response *CreateTeamResponse, err error) {
    return c.CreateTeamWithContext(context.Background(), request)
}

// CreateTeam
// This API is used to create a team.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETOPERATERESOURCEFAILED = "FailedOperation.GetOperateResourceFailed"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EXISTTEAMNAME = "InvalidParameterValue.ExistTeamName"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) CreateTeamWithContext(ctx context.Context, request *CreateTeamRequest) (response *CreateTeamResponse, err error) {
    if request == nil {
        request = NewCreateTeamRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTeam require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTeamResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTeamMemberRequest() (request *CreateTeamMemberRequest) {
    request = &CreateTeamMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "CreateTeamMember")
    
    
    return
}

func NewCreateTeamMemberResponse() (response *CreateTeamMemberResponse) {
    response = &CreateTeamMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTeamMember
// This API is used to add team members
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDREQUESTENTERPRISEINFO = "FailedOperation.InvalidRequestEnterpriseInfo"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_NOACCESSPERMISSION = "FailedOperation.NoAccessPermission"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EXISTUSERACCOUNT = "InvalidParameterValue.ExistUserAccount"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateTeamMember(request *CreateTeamMemberRequest) (response *CreateTeamMemberResponse, err error) {
    return c.CreateTeamMemberWithContext(context.Background(), request)
}

// CreateTeamMember
// This API is used to add team members
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDREQUESTENTERPRISEINFO = "FailedOperation.InvalidRequestEnterpriseInfo"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_NOACCESSPERMISSION = "FailedOperation.NoAccessPermission"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EXISTUSERACCOUNT = "InvalidParameterValue.ExistUserAccount"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateTeamMemberWithContext(ctx context.Context, request *CreateTeamMemberRequest) (response *CreateTeamMemberResponse, err error) {
    if request == nil {
        request = NewCreateTeamMemberRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTeamMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTeamMemberResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserRequest() (request *CreateUserRequest) {
    request = &CreateUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "CreateUser")
    
    
    return
}

func NewCreateUserResponse() (response *CreateUserResponse) {
    response = &CreateUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateUser
// This API is used to create a user
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETOPERATERESOURCEFAILED = "FailedOperation.GetOperateResourceFailed"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EXISTUSERACCOUNT = "InvalidParameterValue.ExistUserAccount"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) CreateUser(request *CreateUserRequest) (response *CreateUserResponse, err error) {
    return c.CreateUserWithContext(context.Background(), request)
}

// CreateUser
// This API is used to create a user
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETOPERATERESOURCEFAILED = "FailedOperation.GetOperateResourceFailed"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EXISTUSERACCOUNT = "InvalidParameterValue.ExistUserAccount"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) CreateUserWithContext(ctx context.Context, request *CreateUserRequest) (response *CreateUserResponse, err error) {
    if request == nil {
        request = NewCreateUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApplicationRequest() (request *DeleteApplicationRequest) {
    request = &DeleteApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DeleteApplication")
    
    
    return
}

func NewDeleteApplicationResponse() (response *DeleteApplicationResponse) {
    response = &DeleteApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteApplication
// This API is used to delete applications
//
// error code that may be returned:
//  FAILEDOPERATION_APPALREADYBINDAUDIT = "FailedOperation.AppAlreadyBindAudit"
//  FAILEDOPERATION_APPALREADYBINDMINIPROGRAM = "FailedOperation.AppAlreadyBindMiniProgram"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONID = "InvalidParameterValue.InvalidApplicationId"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  RESOURCENOTFOUND_NOTFOUNDPURCHASEDPACKAGE = "ResourceNotFound.NotFoundPurchasedPackage"
func (c *Client) DeleteApplication(request *DeleteApplicationRequest) (response *DeleteApplicationResponse, err error) {
    return c.DeleteApplicationWithContext(context.Background(), request)
}

// DeleteApplication
// This API is used to delete applications
//
// error code that may be returned:
//  FAILEDOPERATION_APPALREADYBINDAUDIT = "FailedOperation.AppAlreadyBindAudit"
//  FAILEDOPERATION_APPALREADYBINDMINIPROGRAM = "FailedOperation.AppAlreadyBindMiniProgram"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONID = "InvalidParameterValue.InvalidApplicationId"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  RESOURCENOTFOUND_NOTFOUNDPURCHASEDPACKAGE = "ResourceNotFound.NotFoundPurchasedPackage"
func (c *Client) DeleteApplicationWithContext(ctx context.Context, request *DeleteApplicationRequest) (response *DeleteApplicationResponse, err error) {
    if request == nil {
        request = NewDeleteApplicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApplicationSensitiveAPIRequest() (request *DeleteApplicationSensitiveAPIRequest) {
    request = &DeleteApplicationSensitiveAPIRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DeleteApplicationSensitiveAPI")
    
    
    return
}

func NewDeleteApplicationSensitiveAPIResponse() (response *DeleteApplicationSensitiveAPIResponse) {
    response = &DeleteApplicationSensitiveAPIResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteApplicationSensitiveAPI
// This API is used to delete a sensitive API.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOACCESSPERMISSION = "FailedOperation.NoAccessPermission"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DeleteApplicationSensitiveAPI(request *DeleteApplicationSensitiveAPIRequest) (response *DeleteApplicationSensitiveAPIResponse, err error) {
    return c.DeleteApplicationSensitiveAPIWithContext(context.Background(), request)
}

// DeleteApplicationSensitiveAPI
// This API is used to delete a sensitive API.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOACCESSPERMISSION = "FailedOperation.NoAccessPermission"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DeleteApplicationSensitiveAPIWithContext(ctx context.Context, request *DeleteApplicationSensitiveAPIRequest) (response *DeleteApplicationSensitiveAPIResponse, err error) {
    if request == nil {
        request = NewDeleteApplicationSensitiveAPIRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApplicationSensitiveAPI require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteApplicationSensitiveAPIResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGlobalDomainRequest() (request *DeleteGlobalDomainRequest) {
    request = &DeleteGlobalDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DeleteGlobalDomain")
    
    
    return
}

func NewDeleteGlobalDomainResponse() (response *DeleteGlobalDomainResponse) {
    response = &DeleteGlobalDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteGlobalDomain
// This API is used to delete domains from allowlist or blocklist
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DeleteGlobalDomain(request *DeleteGlobalDomainRequest) (response *DeleteGlobalDomainResponse, err error) {
    return c.DeleteGlobalDomainWithContext(context.Background(), request)
}

// DeleteGlobalDomain
// This API is used to delete domains from allowlist or blocklist
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DeleteGlobalDomainWithContext(ctx context.Context, request *DeleteGlobalDomainRequest) (response *DeleteGlobalDomainResponse, err error) {
    if request == nil {
        request = NewDeleteGlobalDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGlobalDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGlobalDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMNPRequest() (request *DeleteMNPRequest) {
    request = &DeleteMNPRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DeleteMNP")
    
    
    return
}

func NewDeleteMNPResponse() (response *DeleteMNPResponse) {
    response = &DeleteMNPResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteMNP
// This API is used to delete the mini program
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMNPID = "InvalidParameterValue.InvalidMNPId"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DeleteMNP(request *DeleteMNPRequest) (response *DeleteMNPResponse, err error) {
    return c.DeleteMNPWithContext(context.Background(), request)
}

// DeleteMNP
// This API is used to delete the mini program
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMNPID = "InvalidParameterValue.InvalidMNPId"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DeleteMNPWithContext(ctx context.Context, request *DeleteMNPRequest) (response *DeleteMNPResponse, err error) {
    if request == nil {
        request = NewDeleteMNPRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMNP require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMNPResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSensitiveAPIRequest() (request *DeleteSensitiveAPIRequest) {
    request = &DeleteSensitiveAPIRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DeleteSensitiveAPI")
    
    
    return
}

func NewDeleteSensitiveAPIResponse() (response *DeleteSensitiveAPIResponse) {
    response = &DeleteSensitiveAPIResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSensitiveAPI
// This API is used to delete sensitive API
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteSensitiveAPI(request *DeleteSensitiveAPIRequest) (response *DeleteSensitiveAPIResponse, err error) {
    return c.DeleteSensitiveAPIWithContext(context.Background(), request)
}

// DeleteSensitiveAPI
// This API is used to delete sensitive API
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteSensitiveAPIWithContext(ctx context.Context, request *DeleteSensitiveAPIRequest) (response *DeleteSensitiveAPIResponse, err error) {
    if request == nil {
        request = NewDeleteSensitiveAPIRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSensitiveAPI require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSensitiveAPIResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTeamRequest() (request *DeleteTeamRequest) {
    request = &DeleteTeamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DeleteTeam")
    
    
    return
}

func NewDeleteTeamResponse() (response *DeleteTeamResponse) {
    response = &DeleteTeamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTeam
// This API is used to delete a team
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  INVALIDPARAMETERVALUE_INVALIDTEAMID = "InvalidParameterValue.InvalidTeamId"
func (c *Client) DeleteTeam(request *DeleteTeamRequest) (response *DeleteTeamResponse, err error) {
    return c.DeleteTeamWithContext(context.Background(), request)
}

// DeleteTeam
// This API is used to delete a team
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  INVALIDPARAMETERVALUE_INVALIDTEAMID = "InvalidParameterValue.InvalidTeamId"
func (c *Client) DeleteTeamWithContext(ctx context.Context, request *DeleteTeamRequest) (response *DeleteTeamResponse, err error) {
    if request == nil {
        request = NewDeleteTeamRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTeam require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTeamResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTeamMemberRequest() (request *DeleteTeamMemberRequest) {
    request = &DeleteTeamMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DeleteTeamMember")
    
    
    return
}

func NewDeleteTeamMemberResponse() (response *DeleteTeamMemberResponse) {
    response = &DeleteTeamMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTeamMember
// This API is used to delete a team member
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  FAILEDOPERATION_UNABLEOPERATEADMINACCOUNT = "FailedOperation.UnableOperateAdminAccount"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  INVALIDPARAMETERVALUE_INVALIDTEAMID = "InvalidParameterValue.InvalidTeamId"
func (c *Client) DeleteTeamMember(request *DeleteTeamMemberRequest) (response *DeleteTeamMemberResponse, err error) {
    return c.DeleteTeamMemberWithContext(context.Background(), request)
}

// DeleteTeamMember
// This API is used to delete a team member
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  FAILEDOPERATION_UNABLEOPERATEADMINACCOUNT = "FailedOperation.UnableOperateAdminAccount"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  INVALIDPARAMETERVALUE_INVALIDTEAMID = "InvalidParameterValue.InvalidTeamId"
func (c *Client) DeleteTeamMemberWithContext(ctx context.Context, request *DeleteTeamMemberRequest) (response *DeleteTeamMemberResponse, err error) {
    if request == nil {
        request = NewDeleteTeamMemberRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTeamMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTeamMemberResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUserRequest() (request *DeleteUserRequest) {
    request = &DeleteUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DeleteUser")
    
    
    return
}

func NewDeleteUserResponse() (response *DeleteUserResponse) {
    response = &DeleteUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteUser
// This API is used to delete a user
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  FAILEDOPERATION_UNABLEOPERATEADMINACCOUNT = "FailedOperation.UnableOperateAdminAccount"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  INVALIDPARAMETERVALUE_INVALIDUSERID = "InvalidParameterValue.InvalidUserId"
func (c *Client) DeleteUser(request *DeleteUserRequest) (response *DeleteUserResponse, err error) {
    return c.DeleteUserWithContext(context.Background(), request)
}

// DeleteUser
// This API is used to delete a user
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  FAILEDOPERATION_UNABLEOPERATEADMINACCOUNT = "FailedOperation.UnableOperateAdminAccount"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  INVALIDPARAMETERVALUE_INVALIDUSERID = "InvalidParameterValue.InvalidUserId"
func (c *Client) DeleteUserWithContext(ctx context.Context, request *DeleteUserRequest) (response *DeleteUserResponse, err error) {
    if request == nil {
        request = NewDeleteUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUserResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationRequest() (request *DescribeApplicationRequest) {
    request = &DescribeApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeApplication")
    
    
    return
}

func NewDescribeApplicationResponse() (response *DescribeApplicationResponse) {
    response = &DescribeApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApplication
// This API is used to query details of an application. 
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONID = "InvalidParameterValue.InvalidApplicationId"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DescribeApplication(request *DescribeApplicationRequest) (response *DescribeApplicationResponse, err error) {
    return c.DescribeApplicationWithContext(context.Background(), request)
}

// DescribeApplication
// This API is used to query details of an application. 
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONID = "InvalidParameterValue.InvalidApplicationId"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DescribeApplicationWithContext(ctx context.Context, request *DescribeApplicationRequest) (response *DescribeApplicationResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationConfigRequest() (request *DescribeApplicationConfigRequest) {
    request = &DescribeApplicationConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeApplicationConfig")
    
    
    return
}

func NewDescribeApplicationConfigResponse() (response *DescribeApplicationConfigResponse) {
    response = &DescribeApplicationConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApplicationConfig
// This API is used to query application configuration information
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeApplicationConfig(request *DescribeApplicationConfigRequest) (response *DescribeApplicationConfigResponse, err error) {
    return c.DescribeApplicationConfigWithContext(context.Background(), request)
}

// DescribeApplicationConfig
// This API is used to query application configuration information
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeApplicationConfigWithContext(ctx context.Context, request *DescribeApplicationConfigRequest) (response *DescribeApplicationConfigResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationConfigFileRequest() (request *DescribeApplicationConfigFileRequest) {
    request = &DescribeApplicationConfigFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeApplicationConfigFile")
    
    
    return
}

func NewDescribeApplicationConfigFileResponse() (response *DescribeApplicationConfigFileResponse) {
    response = &DescribeApplicationConfigFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApplicationConfigFile
// This API is used to query the configuration files of an application.
//
// error code that may be returned:
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONID = "InvalidParameterValue.InvalidApplicationId"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeApplicationConfigFile(request *DescribeApplicationConfigFileRequest) (response *DescribeApplicationConfigFileResponse, err error) {
    return c.DescribeApplicationConfigFileWithContext(context.Background(), request)
}

// DescribeApplicationConfigFile
// This API is used to query the configuration files of an application.
//
// error code that may be returned:
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONID = "InvalidParameterValue.InvalidApplicationId"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeApplicationConfigFileWithContext(ctx context.Context, request *DescribeApplicationConfigFileRequest) (response *DescribeApplicationConfigFileResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationConfigFileRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationConfigFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationConfigFileResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationListRequest() (request *DescribeApplicationListRequest) {
    request = &DescribeApplicationListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeApplicationList")
    
    
    return
}

func NewDescribeApplicationListResponse() (response *DescribeApplicationListResponse) {
    response = &DescribeApplicationListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApplicationList
// This API is used to query the list of application. 
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  INVALIDPARAMETERVALUE_QUERYPARAMETERLENGTHEXCEEDLIMIT = "InvalidParameterValue.QueryParameterLengthExceedLimit"
func (c *Client) DescribeApplicationList(request *DescribeApplicationListRequest) (response *DescribeApplicationListResponse, err error) {
    return c.DescribeApplicationListWithContext(context.Background(), request)
}

// DescribeApplicationList
// This API is used to query the list of application. 
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  INVALIDPARAMETERVALUE_QUERYPARAMETERLENGTHEXCEEDLIMIT = "InvalidParameterValue.QueryParameterLengthExceedLimit"
func (c *Client) DescribeApplicationListWithContext(ctx context.Context, request *DescribeApplicationListRequest) (response *DescribeApplicationListResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationMNPVersionAuditListRequest() (request *DescribeApplicationMNPVersionAuditListRequest) {
    request = &DescribeApplicationMNPVersionAuditListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeApplicationMNPVersionAuditList")
    
    
    return
}

func NewDescribeApplicationMNPVersionAuditListResponse() (response *DescribeApplicationMNPVersionAuditListResponse) {
    response = &DescribeApplicationMNPVersionAuditListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApplicationMNPVersionAuditList
// This API is used to query the approval list of the mini program versions
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
func (c *Client) DescribeApplicationMNPVersionAuditList(request *DescribeApplicationMNPVersionAuditListRequest) (response *DescribeApplicationMNPVersionAuditListResponse, err error) {
    return c.DescribeApplicationMNPVersionAuditListWithContext(context.Background(), request)
}

// DescribeApplicationMNPVersionAuditList
// This API is used to query the approval list of the mini program versions
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
func (c *Client) DescribeApplicationMNPVersionAuditListWithContext(ctx context.Context, request *DescribeApplicationMNPVersionAuditListRequest) (response *DescribeApplicationMNPVersionAuditListResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationMNPVersionAuditListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationMNPVersionAuditList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationMNPVersionAuditListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationSensitiveAPIListRequest() (request *DescribeApplicationSensitiveAPIListRequest) {
    request = &DescribeApplicationSensitiveAPIListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeApplicationSensitiveAPIList")
    
    
    return
}

func NewDescribeApplicationSensitiveAPIListResponse() (response *DescribeApplicationSensitiveAPIListResponse) {
    response = &DescribeApplicationSensitiveAPIListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApplicationSensitiveAPIList
// This API is used to list sensitive APIs of an application. 
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONID = "InvalidParameterValue.InvalidApplicationId"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeApplicationSensitiveAPIList(request *DescribeApplicationSensitiveAPIListRequest) (response *DescribeApplicationSensitiveAPIListResponse, err error) {
    return c.DescribeApplicationSensitiveAPIListWithContext(context.Background(), request)
}

// DescribeApplicationSensitiveAPIList
// This API is used to list sensitive APIs of an application. 
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONID = "InvalidParameterValue.InvalidApplicationId"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeApplicationSensitiveAPIListWithContext(ctx context.Context, request *DescribeApplicationSensitiveAPIListRequest) (response *DescribeApplicationSensitiveAPIListResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationSensitiveAPIListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationSensitiveAPIList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationSensitiveAPIListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConsoleMNPVersionCompileTaskRequest() (request *DescribeConsoleMNPVersionCompileTaskRequest) {
    request = &DescribeConsoleMNPVersionCompileTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeConsoleMNPVersionCompileTask")
    
    
    return
}

func NewDescribeConsoleMNPVersionCompileTaskResponse() (response *DescribeConsoleMNPVersionCompileTaskResponse) {
    response = &DescribeConsoleMNPVersionCompileTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConsoleMNPVersionCompileTask
// This API is used to query if the mini program version is created successfully
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
func (c *Client) DescribeConsoleMNPVersionCompileTask(request *DescribeConsoleMNPVersionCompileTaskRequest) (response *DescribeConsoleMNPVersionCompileTaskResponse, err error) {
    return c.DescribeConsoleMNPVersionCompileTaskWithContext(context.Background(), request)
}

// DescribeConsoleMNPVersionCompileTask
// This API is used to query if the mini program version is created successfully
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
func (c *Client) DescribeConsoleMNPVersionCompileTaskWithContext(ctx context.Context, request *DescribeConsoleMNPVersionCompileTaskRequest) (response *DescribeConsoleMNPVersionCompileTaskResponse, err error) {
    if request == nil {
        request = NewDescribeConsoleMNPVersionCompileTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConsoleMNPVersionCompileTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConsoleMNPVersionCompileTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainInfoRequest() (request *DescribeDomainInfoRequest) {
    request = &DescribeDomainInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeDomainInfo")
    
    
    return
}

func NewDescribeDomainInfoResponse() (response *DescribeDomainInfoResponse) {
    response = &DescribeDomainInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDomainInfo
// This API is used to query the domain name list configured for the mini program
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
func (c *Client) DescribeDomainInfo(request *DescribeDomainInfoRequest) (response *DescribeDomainInfoResponse, err error) {
    return c.DescribeDomainInfoWithContext(context.Background(), request)
}

// DescribeDomainInfo
// This API is used to query the domain name list configured for the mini program
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
func (c *Client) DescribeDomainInfoWithContext(ctx context.Context, request *DescribeDomainInfoRequest) (response *DescribeDomainInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDomainInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomainInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainTeamListRequest() (request *DescribeDomainTeamListRequest) {
    request = &DescribeDomainTeamListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeDomainTeamList")
    
    
    return
}

func NewDescribeDomainTeamListResponse() (response *DescribeDomainTeamListResponse) {
    response = &DescribeDomainTeamListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDomainTeamList
// This API is used to query the list of teams with domains that obtained the ICP filing
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETOPERATERESOURCEFAILED = "FailedOperation.GetOperateResourceFailed"
//  FAILEDOPERATION_INVALIDREQUESTENTERPRISEINFO = "FailedOperation.InvalidRequestEnterpriseInfo"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_NOACCESSPERMISSION = "FailedOperation.NoAccessPermission"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_USERTEAMRELATIONNOTEXIST = "InvalidParameterValue.UserTeamRelationNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDomainTeamList(request *DescribeDomainTeamListRequest) (response *DescribeDomainTeamListResponse, err error) {
    return c.DescribeDomainTeamListWithContext(context.Background(), request)
}

// DescribeDomainTeamList
// This API is used to query the list of teams with domains that obtained the ICP filing
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETOPERATERESOURCEFAILED = "FailedOperation.GetOperateResourceFailed"
//  FAILEDOPERATION_INVALIDREQUESTENTERPRISEINFO = "FailedOperation.InvalidRequestEnterpriseInfo"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_NOACCESSPERMISSION = "FailedOperation.NoAccessPermission"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_USERTEAMRELATIONNOTEXIST = "InvalidParameterValue.UserTeamRelationNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDomainTeamListWithContext(ctx context.Context, request *DescribeDomainTeamListRequest) (response *DescribeDomainTeamListResponse, err error) {
    if request == nil {
        request = NewDescribeDomainTeamListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomainTeamList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainTeamListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGlobalDomainACLRequest() (request *DescribeGlobalDomainACLRequest) {
    request = &DescribeGlobalDomainACLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeGlobalDomainACL")
    
    
    return
}

func NewDescribeGlobalDomainACLResponse() (response *DescribeGlobalDomainACLResponse) {
    response = &DescribeGlobalDomainACLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGlobalDomainACL
// This API is used to query the global domain allowlist and blocklist.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  INVALIDPARAMETERVALUE_QUERYPARAMETERLENGTHEXCEEDLIMIT = "InvalidParameterValue.QueryParameterLengthExceedLimit"
func (c *Client) DescribeGlobalDomainACL(request *DescribeGlobalDomainACLRequest) (response *DescribeGlobalDomainACLResponse, err error) {
    return c.DescribeGlobalDomainACLWithContext(context.Background(), request)
}

// DescribeGlobalDomainACL
// This API is used to query the global domain allowlist and blocklist.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  INVALIDPARAMETERVALUE_QUERYPARAMETERLENGTHEXCEEDLIMIT = "InvalidParameterValue.QueryParameterLengthExceedLimit"
func (c *Client) DescribeGlobalDomainACLWithContext(ctx context.Context, request *DescribeGlobalDomainACLRequest) (response *DescribeGlobalDomainACLResponse, err error) {
    if request == nil {
        request = NewDescribeGlobalDomainACLRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGlobalDomainACL require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGlobalDomainACLResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGlobalDomainListRequest() (request *DescribeGlobalDomainListRequest) {
    request = &DescribeGlobalDomainListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeGlobalDomainList")
    
    
    return
}

func NewDescribeGlobalDomainListResponse() (response *DescribeGlobalDomainListResponse) {
    response = &DescribeGlobalDomainListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGlobalDomainList
// This API is used to query domain allowlist and blocklist
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDREQUESTENTERPRISEINFO = "FailedOperation.InvalidRequestEnterpriseInfo"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_USERTEAMRELATIONNOTEXIST = "InvalidParameterValue.UserTeamRelationNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeGlobalDomainList(request *DescribeGlobalDomainListRequest) (response *DescribeGlobalDomainListResponse, err error) {
    return c.DescribeGlobalDomainListWithContext(context.Background(), request)
}

// DescribeGlobalDomainList
// This API is used to query domain allowlist and blocklist
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDREQUESTENTERPRISEINFO = "FailedOperation.InvalidRequestEnterpriseInfo"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_USERTEAMRELATIONNOTEXIST = "InvalidParameterValue.UserTeamRelationNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeGlobalDomainListWithContext(ctx context.Context, request *DescribeGlobalDomainListRequest) (response *DescribeGlobalDomainListResponse, err error) {
    if request == nil {
        request = NewDescribeGlobalDomainListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGlobalDomainList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGlobalDomainListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMNPRequest() (request *DescribeMNPRequest) {
    request = &DescribeMNPRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeMNP")
    
    
    return
}

func NewDescribeMNPResponse() (response *DescribeMNPResponse) {
    response = &DescribeMNPResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMNP
// This API is used to query details of a mini program. 
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMNPID = "InvalidParameterValue.InvalidMNPId"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DescribeMNP(request *DescribeMNPRequest) (response *DescribeMNPResponse, err error) {
    return c.DescribeMNPWithContext(context.Background(), request)
}

// DescribeMNP
// This API is used to query details of a mini program. 
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMNPID = "InvalidParameterValue.InvalidMNPId"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DescribeMNPWithContext(ctx context.Context, request *DescribeMNPRequest) (response *DescribeMNPResponse, err error) {
    if request == nil {
        request = NewDescribeMNPRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMNP require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMNPResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMNPAllStageVersionsRequest() (request *DescribeMNPAllStageVersionsRequest) {
    request = &DescribeMNPAllStageVersionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeMNPAllStageVersions")
    
    
    return
}

func NewDescribeMNPAllStageVersionsResponse() (response *DescribeMNPAllStageVersionsResponse) {
    response = &DescribeMNPAllStageVersionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMNPAllStageVersions
// This API is used to query the mini program version management information
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMNPID = "InvalidParameterValue.InvalidMNPId"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DescribeMNPAllStageVersions(request *DescribeMNPAllStageVersionsRequest) (response *DescribeMNPAllStageVersionsResponse, err error) {
    return c.DescribeMNPAllStageVersionsWithContext(context.Background(), request)
}

// DescribeMNPAllStageVersions
// This API is used to query the mini program version management information
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMNPID = "InvalidParameterValue.InvalidMNPId"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DescribeMNPAllStageVersionsWithContext(ctx context.Context, request *DescribeMNPAllStageVersionsRequest) (response *DescribeMNPAllStageVersionsResponse, err error) {
    if request == nil {
        request = NewDescribeMNPAllStageVersionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMNPAllStageVersions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMNPAllStageVersionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMNPApprovalListRequest() (request *DescribeMNPApprovalListRequest) {
    request = &DescribeMNPApprovalListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeMNPApprovalList")
    
    
    return
}

func NewDescribeMNPApprovalListResponse() (response *DescribeMNPApprovalListResponse) {
    response = &DescribeMNPApprovalListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMNPApprovalList
// This API is used to list the approval requests related with a mini program version.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  INVALIDPARAMETERVALUE_QUERYPARAMETERLENGTHEXCEEDLIMIT = "InvalidParameterValue.QueryParameterLengthExceedLimit"
func (c *Client) DescribeMNPApprovalList(request *DescribeMNPApprovalListRequest) (response *DescribeMNPApprovalListResponse, err error) {
    return c.DescribeMNPApprovalListWithContext(context.Background(), request)
}

// DescribeMNPApprovalList
// This API is used to list the approval requests related with a mini program version.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  INVALIDPARAMETERVALUE_QUERYPARAMETERLENGTHEXCEEDLIMIT = "InvalidParameterValue.QueryParameterLengthExceedLimit"
func (c *Client) DescribeMNPApprovalListWithContext(ctx context.Context, request *DescribeMNPApprovalListRequest) (response *DescribeMNPApprovalListResponse, err error) {
    if request == nil {
        request = NewDescribeMNPApprovalListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMNPApprovalList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMNPApprovalListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMNPBoardRequest() (request *DescribeMNPBoardRequest) {
    request = &DescribeMNPBoardRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeMNPBoard")
    
    
    return
}

func NewDescribeMNPBoardResponse() (response *DescribeMNPBoardResponse) {
    response = &DescribeMNPBoardResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMNPBoard
// This API is used to query the mini program version management information
//
// error code that may be returned:
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeMNPBoard(request *DescribeMNPBoardRequest) (response *DescribeMNPBoardResponse, err error) {
    return c.DescribeMNPBoardWithContext(context.Background(), request)
}

// DescribeMNPBoard
// This API is used to query the mini program version management information
//
// error code that may be returned:
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeMNPBoardWithContext(ctx context.Context, request *DescribeMNPBoardRequest) (response *DescribeMNPBoardResponse, err error) {
    if request == nil {
        request = NewDescribeMNPBoardRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMNPBoard require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMNPBoardResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMNPCategoryRequest() (request *DescribeMNPCategoryRequest) {
    request = &DescribeMNPCategoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeMNPCategory")
    
    
    return
}

func NewDescribeMNPCategoryResponse() (response *DescribeMNPCategoryResponse) {
    response = &DescribeMNPCategoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMNPCategory
// This API is used to query the list of mini program types.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DescribeMNPCategory(request *DescribeMNPCategoryRequest) (response *DescribeMNPCategoryResponse, err error) {
    return c.DescribeMNPCategoryWithContext(context.Background(), request)
}

// DescribeMNPCategory
// This API is used to query the list of mini program types.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DescribeMNPCategoryWithContext(ctx context.Context, request *DescribeMNPCategoryRequest) (response *DescribeMNPCategoryResponse, err error) {
    if request == nil {
        request = NewDescribeMNPCategoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMNPCategory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMNPCategoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMNPDetailRequest() (request *DescribeMNPDetailRequest) {
    request = &DescribeMNPDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeMNPDetail")
    
    
    return
}

func NewDescribeMNPDetailResponse() (response *DescribeMNPDetailResponse) {
    response = &DescribeMNPDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMNPDetail
// This API is used to query mini program details
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeMNPDetail(request *DescribeMNPDetailRequest) (response *DescribeMNPDetailResponse, err error) {
    return c.DescribeMNPDetailWithContext(context.Background(), request)
}

// DescribeMNPDetail
// This API is used to query mini program details
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeMNPDetailWithContext(ctx context.Context, request *DescribeMNPDetailRequest) (response *DescribeMNPDetailResponse, err error) {
    if request == nil {
        request = NewDescribeMNPDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMNPDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMNPDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMNPDomainACLRequest() (request *DescribeMNPDomainACLRequest) {
    request = &DescribeMNPDomainACLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeMNPDomainACL")
    
    
    return
}

func NewDescribeMNPDomainACLResponse() (response *DescribeMNPDomainACLResponse) {
    response = &DescribeMNPDomainACLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMNPDomainACL
// This API is used to query the domain name allowlist / blocklist of a mini program. 
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDTEAMID = "InvalidParameterValue.InvalidTeamId"
func (c *Client) DescribeMNPDomainACL(request *DescribeMNPDomainACLRequest) (response *DescribeMNPDomainACLResponse, err error) {
    return c.DescribeMNPDomainACLWithContext(context.Background(), request)
}

// DescribeMNPDomainACL
// This API is used to query the domain name allowlist / blocklist of a mini program. 
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDTEAMID = "InvalidParameterValue.InvalidTeamId"
func (c *Client) DescribeMNPDomainACLWithContext(ctx context.Context, request *DescribeMNPDomainACLRequest) (response *DescribeMNPDomainACLResponse, err error) {
    if request == nil {
        request = NewDescribeMNPDomainACLRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMNPDomainACL require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMNPDomainACLResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMNPListRequest() (request *DescribeMNPListRequest) {
    request = &DescribeMNPListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeMNPList")
    
    
    return
}

func NewDescribeMNPListResponse() (response *DescribeMNPListResponse) {
    response = &DescribeMNPListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMNPList
// This API is used to query the list of mini programs.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DescribeMNPList(request *DescribeMNPListRequest) (response *DescribeMNPListResponse, err error) {
    return c.DescribeMNPListWithContext(context.Background(), request)
}

// DescribeMNPList
// This API is used to query the list of mini programs.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DescribeMNPListWithContext(ctx context.Context, request *DescribeMNPListRequest) (response *DescribeMNPListResponse, err error) {
    if request == nil {
        request = NewDescribeMNPListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMNPList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMNPListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMNPManagerDetailRequest() (request *DescribeMNPManagerDetailRequest) {
    request = &DescribeMNPManagerDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeMNPManagerDetail")
    
    
    return
}

func NewDescribeMNPManagerDetailResponse() (response *DescribeMNPManagerDetailResponse) {
    response = &DescribeMNPManagerDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMNPManagerDetail
// This API is used to query mini program details in the mini program list
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeMNPManagerDetail(request *DescribeMNPManagerDetailRequest) (response *DescribeMNPManagerDetailResponse, err error) {
    return c.DescribeMNPManagerDetailWithContext(context.Background(), request)
}

// DescribeMNPManagerDetail
// This API is used to query mini program details in the mini program list
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeMNPManagerDetailWithContext(ctx context.Context, request *DescribeMNPManagerDetailRequest) (response *DescribeMNPManagerDetailResponse, err error) {
    if request == nil {
        request = NewDescribeMNPManagerDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMNPManagerDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMNPManagerDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMNPManagerListRequest() (request *DescribeMNPManagerListRequest) {
    request = &DescribeMNPManagerListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeMNPManagerList")
    
    
    return
}

func NewDescribeMNPManagerListResponse() (response *DescribeMNPManagerListResponse) {
    response = &DescribeMNPManagerListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMNPManagerList
// This API is used to query the mini program list
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeMNPManagerList(request *DescribeMNPManagerListRequest) (response *DescribeMNPManagerListResponse, err error) {
    return c.DescribeMNPManagerListWithContext(context.Background(), request)
}

// DescribeMNPManagerList
// This API is used to query the mini program list
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeMNPManagerListWithContext(ctx context.Context, request *DescribeMNPManagerListRequest) (response *DescribeMNPManagerListResponse, err error) {
    if request == nil {
        request = NewDescribeMNPManagerListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMNPManagerList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMNPManagerListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMNPOfflinePackageURLRequest() (request *DescribeMNPOfflinePackageURLRequest) {
    request = &DescribeMNPOfflinePackageURLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeMNPOfflinePackageURL")
    
    
    return
}

func NewDescribeMNPOfflinePackageURLResponse() (response *DescribeMNPOfflinePackageURLResponse) {
    response = &DescribeMNPOfflinePackageURLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMNPOfflinePackageURL
// DescribeMNPOfflinePackageURL
//
// error code that may be returned:
//  FAILEDOPERATION_MNPONLINEVERSIONNOTEXIST = "FailedOperation.MNPOnlineVersionNotExist"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DescribeMNPOfflinePackageURL(request *DescribeMNPOfflinePackageURLRequest) (response *DescribeMNPOfflinePackageURLResponse, err error) {
    return c.DescribeMNPOfflinePackageURLWithContext(context.Background(), request)
}

// DescribeMNPOfflinePackageURL
// DescribeMNPOfflinePackageURL
//
// error code that may be returned:
//  FAILEDOPERATION_MNPONLINEVERSIONNOTEXIST = "FailedOperation.MNPOnlineVersionNotExist"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DescribeMNPOfflinePackageURLWithContext(ctx context.Context, request *DescribeMNPOfflinePackageURLRequest) (response *DescribeMNPOfflinePackageURLResponse, err error) {
    if request == nil {
        request = NewDescribeMNPOfflinePackageURLRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMNPOfflinePackageURL require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMNPOfflinePackageURLResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMNPPreviewRequest() (request *DescribeMNPPreviewRequest) {
    request = &DescribeMNPPreviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeMNPPreview")
    
    
    return
}

func NewDescribeMNPPreviewResponse() (response *DescribeMNPPreviewResponse) {
    response = &DescribeMNPPreviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMNPPreview
// This API is used to query the details of a mini program preview version.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMNPID = "InvalidParameterValue.InvalidMNPId"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DescribeMNPPreview(request *DescribeMNPPreviewRequest) (response *DescribeMNPPreviewResponse, err error) {
    return c.DescribeMNPPreviewWithContext(context.Background(), request)
}

// DescribeMNPPreview
// This API is used to query the details of a mini program preview version.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMNPID = "InvalidParameterValue.InvalidMNPId"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DescribeMNPPreviewWithContext(ctx context.Context, request *DescribeMNPPreviewRequest) (response *DescribeMNPPreviewResponse, err error) {
    if request == nil {
        request = NewDescribeMNPPreviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMNPPreview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMNPPreviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMNPPrivacyRequest() (request *DescribeMNPPrivacyRequest) {
    request = &DescribeMNPPrivacyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeMNPPrivacy")
    
    
    return
}

func NewDescribeMNPPrivacyResponse() (response *DescribeMNPPrivacyResponse) {
    response = &DescribeMNPPrivacyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMNPPrivacy
// This API is used to query the details filled in the service description
//
// error code that may be returned:
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeMNPPrivacy(request *DescribeMNPPrivacyRequest) (response *DescribeMNPPrivacyResponse, err error) {
    return c.DescribeMNPPrivacyWithContext(context.Background(), request)
}

// DescribeMNPPrivacy
// This API is used to query the details filled in the service description
//
// error code that may be returned:
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeMNPPrivacyWithContext(ctx context.Context, request *DescribeMNPPrivacyRequest) (response *DescribeMNPPrivacyResponse, err error) {
    if request == nil {
        request = NewDescribeMNPPrivacyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMNPPrivacy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMNPPrivacyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMNPReleasedVersionHistoryRequest() (request *DescribeMNPReleasedVersionHistoryRequest) {
    request = &DescribeMNPReleasedVersionHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeMNPReleasedVersionHistory")
    
    
    return
}

func NewDescribeMNPReleasedVersionHistoryResponse() (response *DescribeMNPReleasedVersionHistoryResponse) {
    response = &DescribeMNPReleasedVersionHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMNPReleasedVersionHistory
// This API is used to list all released versions of a mini program.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMNPID = "InvalidParameterValue.InvalidMNPId"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DescribeMNPReleasedVersionHistory(request *DescribeMNPReleasedVersionHistoryRequest) (response *DescribeMNPReleasedVersionHistoryResponse, err error) {
    return c.DescribeMNPReleasedVersionHistoryWithContext(context.Background(), request)
}

// DescribeMNPReleasedVersionHistory
// This API is used to list all released versions of a mini program.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMNPID = "InvalidParameterValue.InvalidMNPId"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DescribeMNPReleasedVersionHistoryWithContext(ctx context.Context, request *DescribeMNPReleasedVersionHistoryRequest) (response *DescribeMNPReleasedVersionHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeMNPReleasedVersionHistoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMNPReleasedVersionHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMNPReleasedVersionHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMNPSensitiveAPIPermissionApprovalRequest() (request *DescribeMNPSensitiveAPIPermissionApprovalRequest) {
    request = &DescribeMNPSensitiveAPIPermissionApprovalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeMNPSensitiveAPIPermissionApproval")
    
    
    return
}

func NewDescribeMNPSensitiveAPIPermissionApprovalResponse() (response *DescribeMNPSensitiveAPIPermissionApprovalResponse) {
    response = &DescribeMNPSensitiveAPIPermissionApprovalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMNPSensitiveAPIPermissionApproval
// This API is used to query details of a specific permission request to call sensitive APIs.
//
// error code that may be returned:
//  FAILEDOPERATION_NOACCESSPERMISSION = "FailedOperation.NoAccessPermission"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DescribeMNPSensitiveAPIPermissionApproval(request *DescribeMNPSensitiveAPIPermissionApprovalRequest) (response *DescribeMNPSensitiveAPIPermissionApprovalResponse, err error) {
    return c.DescribeMNPSensitiveAPIPermissionApprovalWithContext(context.Background(), request)
}

// DescribeMNPSensitiveAPIPermissionApproval
// This API is used to query details of a specific permission request to call sensitive APIs.
//
// error code that may be returned:
//  FAILEDOPERATION_NOACCESSPERMISSION = "FailedOperation.NoAccessPermission"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DescribeMNPSensitiveAPIPermissionApprovalWithContext(ctx context.Context, request *DescribeMNPSensitiveAPIPermissionApprovalRequest) (response *DescribeMNPSensitiveAPIPermissionApprovalResponse, err error) {
    if request == nil {
        request = NewDescribeMNPSensitiveAPIPermissionApprovalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMNPSensitiveAPIPermissionApproval require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMNPSensitiveAPIPermissionApprovalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMNPSensitiveAPIPermissionApprovalListRequest() (request *DescribeMNPSensitiveAPIPermissionApprovalListRequest) {
    request = &DescribeMNPSensitiveAPIPermissionApprovalListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeMNPSensitiveAPIPermissionApprovalList")
    
    
    return
}

func NewDescribeMNPSensitiveAPIPermissionApprovalListResponse() (response *DescribeMNPSensitiveAPIPermissionApprovalListResponse) {
    response = &DescribeMNPSensitiveAPIPermissionApprovalListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMNPSensitiveAPIPermissionApprovalList
// This API is used to query permission requests to allow a mini program calling sensitive APIs.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DescribeMNPSensitiveAPIPermissionApprovalList(request *DescribeMNPSensitiveAPIPermissionApprovalListRequest) (response *DescribeMNPSensitiveAPIPermissionApprovalListResponse, err error) {
    return c.DescribeMNPSensitiveAPIPermissionApprovalListWithContext(context.Background(), request)
}

// DescribeMNPSensitiveAPIPermissionApprovalList
// This API is used to query permission requests to allow a mini program calling sensitive APIs.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DescribeMNPSensitiveAPIPermissionApprovalListWithContext(ctx context.Context, request *DescribeMNPSensitiveAPIPermissionApprovalListRequest) (response *DescribeMNPSensitiveAPIPermissionApprovalListResponse, err error) {
    if request == nil {
        request = NewDescribeMNPSensitiveAPIPermissionApprovalListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMNPSensitiveAPIPermissionApprovalList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMNPSensitiveAPIPermissionApprovalListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMNPSensitiveAPIPermissionListRequest() (request *DescribeMNPSensitiveAPIPermissionListRequest) {
    request = &DescribeMNPSensitiveAPIPermissionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeMNPSensitiveAPIPermissionList")
    
    
    return
}

func NewDescribeMNPSensitiveAPIPermissionListResponse() (response *DescribeMNPSensitiveAPIPermissionListResponse) {
    response = &DescribeMNPSensitiveAPIPermissionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMNPSensitiveAPIPermissionList
// This API is used to query the list of sensitive APIs that available to a mini program.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMNPID = "InvalidParameterValue.InvalidMNPId"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DescribeMNPSensitiveAPIPermissionList(request *DescribeMNPSensitiveAPIPermissionListRequest) (response *DescribeMNPSensitiveAPIPermissionListResponse, err error) {
    return c.DescribeMNPSensitiveAPIPermissionListWithContext(context.Background(), request)
}

// DescribeMNPSensitiveAPIPermissionList
// This API is used to query the list of sensitive APIs that available to a mini program.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMNPID = "InvalidParameterValue.InvalidMNPId"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DescribeMNPSensitiveAPIPermissionListWithContext(ctx context.Context, request *DescribeMNPSensitiveAPIPermissionListRequest) (response *DescribeMNPSensitiveAPIPermissionListResponse, err error) {
    if request == nil {
        request = NewDescribeMNPSensitiveAPIPermissionListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMNPSensitiveAPIPermissionList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMNPSensitiveAPIPermissionListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMNPTypeRequest() (request *DescribeMNPTypeRequest) {
    request = &DescribeMNPTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeMNPType")
    
    
    return
}

func NewDescribeMNPTypeResponse() (response *DescribeMNPTypeResponse) {
    response = &DescribeMNPTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMNPType
// This API is used to query the mini program type list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
func (c *Client) DescribeMNPType(request *DescribeMNPTypeRequest) (response *DescribeMNPTypeResponse, err error) {
    return c.DescribeMNPTypeWithContext(context.Background(), request)
}

// DescribeMNPType
// This API is used to query the mini program type list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
func (c *Client) DescribeMNPTypeWithContext(ctx context.Context, request *DescribeMNPTypeRequest) (response *DescribeMNPTypeResponse, err error) {
    if request == nil {
        request = NewDescribeMNPTypeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMNPType require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMNPTypeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMNPVersionRequest() (request *DescribeMNPVersionRequest) {
    request = &DescribeMNPVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeMNPVersion")
    
    
    return
}

func NewDescribeMNPVersionResponse() (response *DescribeMNPVersionResponse) {
    response = &DescribeMNPVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMNPVersion
// This API is used to query the result of the task to create a mini program version.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMNPVERSIONTASKID = "InvalidParameterValue.InvalidMNPVersionTaskId"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DescribeMNPVersion(request *DescribeMNPVersionRequest) (response *DescribeMNPVersionResponse, err error) {
    return c.DescribeMNPVersionWithContext(context.Background(), request)
}

// DescribeMNPVersion
// This API is used to query the result of the task to create a mini program version.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMNPVERSIONTASKID = "InvalidParameterValue.InvalidMNPVersionTaskId"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DescribeMNPVersionWithContext(ctx context.Context, request *DescribeMNPVersionRequest) (response *DescribeMNPVersionResponse, err error) {
    if request == nil {
        request = NewDescribeMNPVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMNPVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMNPVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMNPVersionPreviewRequest() (request *DescribeMNPVersionPreviewRequest) {
    request = &DescribeMNPVersionPreviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeMNPVersionPreview")
    
    
    return
}

func NewDescribeMNPVersionPreviewResponse() (response *DescribeMNPVersionPreviewResponse) {
    response = &DescribeMNPVersionPreviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMNPVersionPreview
// This API is used to query the details of the mini program preview version
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
func (c *Client) DescribeMNPVersionPreview(request *DescribeMNPVersionPreviewRequest) (response *DescribeMNPVersionPreviewResponse, err error) {
    return c.DescribeMNPVersionPreviewWithContext(context.Background(), request)
}

// DescribeMNPVersionPreview
// This API is used to query the details of the mini program preview version
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
func (c *Client) DescribeMNPVersionPreviewWithContext(ctx context.Context, request *DescribeMNPVersionPreviewRequest) (response *DescribeMNPVersionPreviewResponse, err error) {
    if request == nil {
        request = NewDescribeMNPVersionPreviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMNPVersionPreview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMNPVersionPreviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOfflineMNPPackageRequest() (request *DescribeOfflineMNPPackageRequest) {
    request = &DescribeOfflineMNPPackageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeOfflineMNPPackage")
    
    
    return
}

func NewDescribeOfflineMNPPackageResponse() (response *DescribeOfflineMNPPackageResponse) {
    response = &DescribeOfflineMNPPackageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOfflineMNPPackage
// This API is used to query offline mini program packages
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETOPERATERESOURCEFAILED = "FailedOperation.GetOperateResourceFailed"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_OPERATIONSTEAMNOAPPLICATIONPERMISSION = "FailedOperation.OperationsTeamNoApplicationPermission"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TEAMTYPEMISMATCH = "InvalidParameterValue.TeamTypeMismatch"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeOfflineMNPPackage(request *DescribeOfflineMNPPackageRequest) (response *DescribeOfflineMNPPackageResponse, err error) {
    return c.DescribeOfflineMNPPackageWithContext(context.Background(), request)
}

// DescribeOfflineMNPPackage
// This API is used to query offline mini program packages
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETOPERATERESOURCEFAILED = "FailedOperation.GetOperateResourceFailed"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_OPERATIONSTEAMNOAPPLICATIONPERMISSION = "FailedOperation.OperationsTeamNoApplicationPermission"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TEAMTYPEMISMATCH = "InvalidParameterValue.TeamTypeMismatch"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeOfflineMNPPackageWithContext(ctx context.Context, request *DescribeOfflineMNPPackageRequest) (response *DescribeOfflineMNPPackageResponse, err error) {
    if request == nil {
        request = NewDescribeOfflineMNPPackageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOfflineMNPPackage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOfflineMNPPackageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOnlineVersionRequest() (request *DescribeOnlineVersionRequest) {
    request = &DescribeOnlineVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeOnlineVersion")
    
    
    return
}

func NewDescribeOnlineVersionResponse() (response *DescribeOnlineVersionResponse) {
    response = &DescribeOnlineVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOnlineVersion
// This API is used to query the release version history
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
func (c *Client) DescribeOnlineVersion(request *DescribeOnlineVersionRequest) (response *DescribeOnlineVersionResponse, err error) {
    return c.DescribeOnlineVersionWithContext(context.Background(), request)
}

// DescribeOnlineVersion
// This API is used to query the release version history
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
func (c *Client) DescribeOnlineVersionWithContext(ctx context.Context, request *DescribeOnlineVersionRequest) (response *DescribeOnlineVersionResponse, err error) {
    if request == nil {
        request = NewDescribeOnlineVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOnlineVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOnlineVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRoleDetailRequest() (request *DescribeRoleDetailRequest) {
    request = &DescribeRoleDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeRoleDetail")
    
    
    return
}

func NewDescribeRoleDetailResponse() (response *DescribeRoleDetailResponse) {
    response = &DescribeRoleDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRoleDetail
// This API is used to query role permission information
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
func (c *Client) DescribeRoleDetail(request *DescribeRoleDetailRequest) (response *DescribeRoleDetailResponse, err error) {
    return c.DescribeRoleDetailWithContext(context.Background(), request)
}

// DescribeRoleDetail
// This API is used to query role permission information
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
func (c *Client) DescribeRoleDetailWithContext(ctx context.Context, request *DescribeRoleDetailRequest) (response *DescribeRoleDetailResponse, err error) {
    if request == nil {
        request = NewDescribeRoleDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRoleDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRoleDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRoleListRequest() (request *DescribeRoleListRequest) {
    request = &DescribeRoleListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeRoleList")
    
    
    return
}

func NewDescribeRoleListResponse() (response *DescribeRoleListResponse) {
    response = &DescribeRoleListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRoleList
// This API is used to query the list of roles. 
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DescribeRoleList(request *DescribeRoleListRequest) (response *DescribeRoleListResponse, err error) {
    return c.DescribeRoleListWithContext(context.Background(), request)
}

// DescribeRoleList
// This API is used to query the list of roles. 
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DescribeRoleListWithContext(ctx context.Context, request *DescribeRoleListRequest) (response *DescribeRoleListResponse, err error) {
    if request == nil {
        request = NewDescribeRoleListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRoleList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRoleListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSensitiveAPIAuditListRequest() (request *DescribeSensitiveAPIAuditListRequest) {
    request = &DescribeSensitiveAPIAuditListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeSensitiveAPIAuditList")
    
    
    return
}

func NewDescribeSensitiveAPIAuditListResponse() (response *DescribeSensitiveAPIAuditListResponse) {
    response = &DescribeSensitiveAPIAuditListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSensitiveAPIAuditList
// This API is used to query sensitive API approval list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeSensitiveAPIAuditList(request *DescribeSensitiveAPIAuditListRequest) (response *DescribeSensitiveAPIAuditListResponse, err error) {
    return c.DescribeSensitiveAPIAuditListWithContext(context.Background(), request)
}

// DescribeSensitiveAPIAuditList
// This API is used to query sensitive API approval list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeSensitiveAPIAuditListWithContext(ctx context.Context, request *DescribeSensitiveAPIAuditListRequest) (response *DescribeSensitiveAPIAuditListResponse, err error) {
    if request == nil {
        request = NewDescribeSensitiveAPIAuditListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSensitiveAPIAuditList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSensitiveAPIAuditListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSensitiveAPIListRequest() (request *DescribeSensitiveAPIListRequest) {
    request = &DescribeSensitiveAPIListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeSensitiveAPIList")
    
    
    return
}

func NewDescribeSensitiveAPIListResponse() (response *DescribeSensitiveAPIListResponse) {
    response = &DescribeSensitiveAPIListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSensitiveAPIList
// This API is used to query all sensitive APIs list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeSensitiveAPIList(request *DescribeSensitiveAPIListRequest) (response *DescribeSensitiveAPIListResponse, err error) {
    return c.DescribeSensitiveAPIListWithContext(context.Background(), request)
}

// DescribeSensitiveAPIList
// This API is used to query all sensitive APIs list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeSensitiveAPIListWithContext(ctx context.Context, request *DescribeSensitiveAPIListRequest) (response *DescribeSensitiveAPIListResponse, err error) {
    if request == nil {
        request = NewDescribeSensitiveAPIListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSensitiveAPIList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSensitiveAPIListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSensitiveApiApplyDetailRequest() (request *DescribeSensitiveApiApplyDetailRequest) {
    request = &DescribeSensitiveApiApplyDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeSensitiveApiApplyDetail")
    
    
    return
}

func NewDescribeSensitiveApiApplyDetailResponse() (response *DescribeSensitiveApiApplyDetailResponse) {
    response = &DescribeSensitiveApiApplyDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSensitiveApiApplyDetail
// This API is used to query sensitive API permission approval details
//
// error code that may be returned:
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
func (c *Client) DescribeSensitiveApiApplyDetail(request *DescribeSensitiveApiApplyDetailRequest) (response *DescribeSensitiveApiApplyDetailResponse, err error) {
    return c.DescribeSensitiveApiApplyDetailWithContext(context.Background(), request)
}

// DescribeSensitiveApiApplyDetail
// This API is used to query sensitive API permission approval details
//
// error code that may be returned:
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
func (c *Client) DescribeSensitiveApiApplyDetailWithContext(ctx context.Context, request *DescribeSensitiveApiApplyDetailRequest) (response *DescribeSensitiveApiApplyDetailResponse, err error) {
    if request == nil {
        request = NewDescribeSensitiveApiApplyDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSensitiveApiApplyDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSensitiveApiApplyDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSensitiveApiAuthListRequest() (request *DescribeSensitiveApiAuthListRequest) {
    request = &DescribeSensitiveApiAuthListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeSensitiveApiAuthList")
    
    
    return
}

func NewDescribeSensitiveApiAuthListResponse() (response *DescribeSensitiveApiAuthListResponse) {
    response = &DescribeSensitiveApiAuthListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSensitiveApiAuthList
// This API is used to query the sensitive API permission list
//
// error code that may be returned:
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
func (c *Client) DescribeSensitiveApiAuthList(request *DescribeSensitiveApiAuthListRequest) (response *DescribeSensitiveApiAuthListResponse, err error) {
    return c.DescribeSensitiveApiAuthListWithContext(context.Background(), request)
}

// DescribeSensitiveApiAuthList
// This API is used to query the sensitive API permission list
//
// error code that may be returned:
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
func (c *Client) DescribeSensitiveApiAuthListWithContext(ctx context.Context, request *DescribeSensitiveApiAuthListRequest) (response *DescribeSensitiveApiAuthListResponse, err error) {
    if request == nil {
        request = NewDescribeSensitiveApiAuthListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSensitiveApiAuthList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSensitiveApiAuthListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSimpleApplicationInfoListRequest() (request *DescribeSimpleApplicationInfoListRequest) {
    request = &DescribeSimpleApplicationInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeSimpleApplicationInfoList")
    
    
    return
}

func NewDescribeSimpleApplicationInfoListResponse() (response *DescribeSimpleApplicationInfoListResponse) {
    response = &DescribeSimpleApplicationInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSimpleApplicationInfoList
// This API is used to query application list information
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDTEAMID = "InvalidParameterValue.InvalidTeamId"
//  INVALIDPARAMETERVALUE_USERTEAMRELATIONNOTEXIST = "InvalidParameterValue.UserTeamRelationNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeSimpleApplicationInfoList(request *DescribeSimpleApplicationInfoListRequest) (response *DescribeSimpleApplicationInfoListResponse, err error) {
    return c.DescribeSimpleApplicationInfoListWithContext(context.Background(), request)
}

// DescribeSimpleApplicationInfoList
// This API is used to query application list information
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDTEAMID = "InvalidParameterValue.InvalidTeamId"
//  INVALIDPARAMETERVALUE_USERTEAMRELATIONNOTEXIST = "InvalidParameterValue.UserTeamRelationNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeSimpleApplicationInfoListWithContext(ctx context.Context, request *DescribeSimpleApplicationInfoListRequest) (response *DescribeSimpleApplicationInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeSimpleApplicationInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSimpleApplicationInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSimpleApplicationInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSimpleTeamListRequest() (request *DescribeSimpleTeamListRequest) {
    request = &DescribeSimpleTeamListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeSimpleTeamList")
    
    
    return
}

func NewDescribeSimpleTeamListResponse() (response *DescribeSimpleTeamListResponse) {
    response = &DescribeSimpleTeamListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSimpleTeamList
// This API is used to query the team information list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDREQUESTENTERPRISEINFO = "FailedOperation.InvalidRequestEnterpriseInfo"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeSimpleTeamList(request *DescribeSimpleTeamListRequest) (response *DescribeSimpleTeamListResponse, err error) {
    return c.DescribeSimpleTeamListWithContext(context.Background(), request)
}

// DescribeSimpleTeamList
// This API is used to query the team information list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDREQUESTENTERPRISEINFO = "FailedOperation.InvalidRequestEnterpriseInfo"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeSimpleTeamListWithContext(ctx context.Context, request *DescribeSimpleTeamListRequest) (response *DescribeSimpleTeamListResponse, err error) {
    if request == nil {
        request = NewDescribeSimpleTeamListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSimpleTeamList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSimpleTeamListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTeamRequest() (request *DescribeTeamRequest) {
    request = &DescribeTeamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeTeam")
    
    
    return
}

func NewDescribeTeamResponse() (response *DescribeTeamResponse) {
    response = &DescribeTeamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTeam
// This API is used to query details of a team. 
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DescribeTeam(request *DescribeTeamRequest) (response *DescribeTeamResponse, err error) {
    return c.DescribeTeamWithContext(context.Background(), request)
}

// DescribeTeam
// This API is used to query details of a team. 
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DescribeTeamWithContext(ctx context.Context, request *DescribeTeamRequest) (response *DescribeTeamResponse, err error) {
    if request == nil {
        request = NewDescribeTeamRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTeam require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTeamResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTeamDomainListRequest() (request *DescribeTeamDomainListRequest) {
    request = &DescribeTeamDomainListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeTeamDomainList")
    
    
    return
}

func NewDescribeTeamDomainListResponse() (response *DescribeTeamDomainListResponse) {
    response = &DescribeTeamDomainListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTeamDomainList
// This API is used to query a specified team’s domains that obtained ICP filing
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDREQUESTENTERPRISEINFO = "FailedOperation.InvalidRequestEnterpriseInfo"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_NOACCESSPERMISSION = "FailedOperation.NoAccessPermission"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_USERTEAMRELATIONNOTEXIST = "InvalidParameterValue.UserTeamRelationNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTeamDomainList(request *DescribeTeamDomainListRequest) (response *DescribeTeamDomainListResponse, err error) {
    return c.DescribeTeamDomainListWithContext(context.Background(), request)
}

// DescribeTeamDomainList
// This API is used to query a specified team’s domains that obtained ICP filing
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDREQUESTENTERPRISEINFO = "FailedOperation.InvalidRequestEnterpriseInfo"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_NOACCESSPERMISSION = "FailedOperation.NoAccessPermission"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_USERTEAMRELATIONNOTEXIST = "InvalidParameterValue.UserTeamRelationNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTeamDomainListWithContext(ctx context.Context, request *DescribeTeamDomainListRequest) (response *DescribeTeamDomainListResponse, err error) {
    if request == nil {
        request = NewDescribeTeamDomainListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTeamDomainList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTeamDomainListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTeamListRequest() (request *DescribeTeamListRequest) {
    request = &DescribeTeamListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeTeamList")
    
    
    return
}

func NewDescribeTeamListResponse() (response *DescribeTeamListResponse) {
    response = &DescribeTeamListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTeamList
// This API is used to query the list of teams. 
//
// error code that may be returned:
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  INVALIDPARAMETERVALUE_QUERYPARAMETERLENGTHEXCEEDLIMIT = "InvalidParameterValue.QueryParameterLengthExceedLimit"
func (c *Client) DescribeTeamList(request *DescribeTeamListRequest) (response *DescribeTeamListResponse, err error) {
    return c.DescribeTeamListWithContext(context.Background(), request)
}

// DescribeTeamList
// This API is used to query the list of teams. 
//
// error code that may be returned:
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  INVALIDPARAMETERVALUE_QUERYPARAMETERLENGTHEXCEEDLIMIT = "InvalidParameterValue.QueryParameterLengthExceedLimit"
func (c *Client) DescribeTeamListWithContext(ctx context.Context, request *DescribeTeamListRequest) (response *DescribeTeamListResponse, err error) {
    if request == nil {
        request = NewDescribeTeamListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTeamList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTeamListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTeamMemberListRequest() (request *DescribeTeamMemberListRequest) {
    request = &DescribeTeamMemberListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeTeamMemberList")
    
    
    return
}

func NewDescribeTeamMemberListResponse() (response *DescribeTeamMemberListResponse) {
    response = &DescribeTeamMemberListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTeamMemberList
// This API is used to query team member list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DescribeTeamMemberList(request *DescribeTeamMemberListRequest) (response *DescribeTeamMemberListResponse, err error) {
    return c.DescribeTeamMemberListWithContext(context.Background(), request)
}

// DescribeTeamMemberList
// This API is used to query team member list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DescribeTeamMemberListWithContext(ctx context.Context, request *DescribeTeamMemberListRequest) (response *DescribeTeamMemberListResponse, err error) {
    if request == nil {
        request = NewDescribeTeamMemberListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTeamMemberList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTeamMemberListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTeamMembersRequest() (request *DescribeTeamMembersRequest) {
    request = &DescribeTeamMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeTeamMembers")
    
    
    return
}

func NewDescribeTeamMembersResponse() (response *DescribeTeamMembersResponse) {
    response = &DescribeTeamMembersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTeamMembers
// This API is used to query the member list of a specified team
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDREQUESTENTERPRISEINFO = "FailedOperation.InvalidRequestEnterpriseInfo"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTeamMembers(request *DescribeTeamMembersRequest) (response *DescribeTeamMembersResponse, err error) {
    return c.DescribeTeamMembersWithContext(context.Background(), request)
}

// DescribeTeamMembers
// This API is used to query the member list of a specified team
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDREQUESTENTERPRISEINFO = "FailedOperation.InvalidRequestEnterpriseInfo"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTeamMembersWithContext(ctx context.Context, request *DescribeTeamMembersRequest) (response *DescribeTeamMembersResponse, err error) {
    if request == nil {
        request = NewDescribeTeamMembersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTeamMembers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTeamMembersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTempSecret4UploadFile2CosRequest() (request *DescribeTempSecret4UploadFile2CosRequest) {
    request = &DescribeTempSecret4UploadFile2CosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeTempSecret4UploadFile2Cos")
    
    
    return
}

func NewDescribeTempSecret4UploadFile2CosResponse() (response *DescribeTempSecret4UploadFile2CosResponse) {
    response = &DescribeTempSecret4UploadFile2CosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTempSecret4UploadFile2Cos
// This API is used to obtain a temporary key for file uploads
//
// error code that may be returned:
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDTEAMID = "InvalidParameterValue.InvalidTeamId"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
func (c *Client) DescribeTempSecret4UploadFile2Cos(request *DescribeTempSecret4UploadFile2CosRequest) (response *DescribeTempSecret4UploadFile2CosResponse, err error) {
    return c.DescribeTempSecret4UploadFile2CosWithContext(context.Background(), request)
}

// DescribeTempSecret4UploadFile2Cos
// This API is used to obtain a temporary key for file uploads
//
// error code that may be returned:
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDTEAMID = "InvalidParameterValue.InvalidTeamId"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
func (c *Client) DescribeTempSecret4UploadFile2CosWithContext(ctx context.Context, request *DescribeTempSecret4UploadFile2CosRequest) (response *DescribeTempSecret4UploadFile2CosResponse, err error) {
    if request == nil {
        request = NewDescribeTempSecret4UploadFile2CosRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTempSecret4UploadFile2Cos require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTempSecret4UploadFile2CosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserRequest() (request *DescribeUserRequest) {
    request = &DescribeUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeUser")
    
    
    return
}

func NewDescribeUserResponse() (response *DescribeUserResponse) {
    response = &DescribeUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUser
// This API is used to query details of a user. 
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  INVALIDPARAMETERVALUE_INVALIDUSERID = "InvalidParameterValue.InvalidUserId"
func (c *Client) DescribeUser(request *DescribeUserRequest) (response *DescribeUserResponse, err error) {
    return c.DescribeUserWithContext(context.Background(), request)
}

// DescribeUser
// This API is used to query details of a user. 
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  INVALIDPARAMETERVALUE_INVALIDUSERID = "InvalidParameterValue.InvalidUserId"
func (c *Client) DescribeUserWithContext(ctx context.Context, request *DescribeUserRequest) (response *DescribeUserResponse, err error) {
    if request == nil {
        request = NewDescribeUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserDetailRequest() (request *DescribeUserDetailRequest) {
    request = &DescribeUserDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeUserDetail")
    
    
    return
}

func NewDescribeUserDetailResponse() (response *DescribeUserDetailResponse) {
    response = &DescribeUserDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserDetail
// This API is used to query user details
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
func (c *Client) DescribeUserDetail(request *DescribeUserDetailRequest) (response *DescribeUserDetailResponse, err error) {
    return c.DescribeUserDetailWithContext(context.Background(), request)
}

// DescribeUserDetail
// This API is used to query user details
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
func (c *Client) DescribeUserDetailWithContext(ctx context.Context, request *DescribeUserDetailRequest) (response *DescribeUserDetailResponse, err error) {
    if request == nil {
        request = NewDescribeUserDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserListRequest() (request *DescribeUserListRequest) {
    request = &DescribeUserListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DescribeUserList")
    
    
    return
}

func NewDescribeUserListResponse() (response *DescribeUserListResponse) {
    response = &DescribeUserListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserList
// This API is used to query the user list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTTYPE = "InvalidParameterValue.InvalidAccountType"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  INVALIDPARAMETERVALUE_INVALIDTEAMID = "InvalidParameterValue.InvalidTeamId"
func (c *Client) DescribeUserList(request *DescribeUserListRequest) (response *DescribeUserListResponse, err error) {
    return c.DescribeUserListWithContext(context.Background(), request)
}

// DescribeUserList
// This API is used to query the user list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTTYPE = "InvalidParameterValue.InvalidAccountType"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  INVALIDPARAMETERVALUE_INVALIDTEAMID = "InvalidParameterValue.InvalidTeamId"
func (c *Client) DescribeUserListWithContext(ctx context.Context, request *DescribeUserListRequest) (response *DescribeUserListResponse, err error) {
    if request == nil {
        request = NewDescribeUserListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserListResponse()
    err = c.Send(request, response)
    return
}

func NewDisableApplicationSensitiveAPIRequest() (request *DisableApplicationSensitiveAPIRequest) {
    request = &DisableApplicationSensitiveAPIRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DisableApplicationSensitiveAPI")
    
    
    return
}

func NewDisableApplicationSensitiveAPIResponse() (response *DisableApplicationSensitiveAPIResponse) {
    response = &DisableApplicationSensitiveAPIResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisableApplicationSensitiveAPI
// This API is used to set a sensitive API to Restricted.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOACCESSPERMISSION = "FailedOperation.NoAccessPermission"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DisableApplicationSensitiveAPI(request *DisableApplicationSensitiveAPIRequest) (response *DisableApplicationSensitiveAPIResponse, err error) {
    return c.DisableApplicationSensitiveAPIWithContext(context.Background(), request)
}

// DisableApplicationSensitiveAPI
// This API is used to set a sensitive API to Restricted.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOACCESSPERMISSION = "FailedOperation.NoAccessPermission"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DisableApplicationSensitiveAPIWithContext(ctx context.Context, request *DisableApplicationSensitiveAPIRequest) (response *DisableApplicationSensitiveAPIResponse, err error) {
    if request == nil {
        request = NewDisableApplicationSensitiveAPIRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableApplicationSensitiveAPI require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableApplicationSensitiveAPIResponse()
    err = c.Send(request, response)
    return
}

func NewDisableTeamDomainRequest() (request *DisableTeamDomainRequest) {
    request = &DisableTeamDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "DisableTeamDomain")
    
    
    return
}

func NewDisableTeamDomainResponse() (response *DisableTeamDomainResponse) {
    response = &DisableTeamDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisableTeamDomain
// This API is used to disable the company’s domain name that obtained the ICP filing
//
// error code that may be returned:
//  FAILEDOPERATION_EXISTBLACKDOMAIN = "FailedOperation.ExistBlackDomain"
//  FAILEDOPERATION_EXISTWHITEDOMAIN = "FailedOperation.ExistWhiteDomain"
//  FAILEDOPERATION_INVALIDREQUESTENTERPRISEINFO = "FailedOperation.InvalidRequestEnterpriseInfo"
//  FAILEDOPERATION_NOACCESSPERMISSION = "FailedOperation.NoAccessPermission"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DisableTeamDomain(request *DisableTeamDomainRequest) (response *DisableTeamDomainResponse, err error) {
    return c.DisableTeamDomainWithContext(context.Background(), request)
}

// DisableTeamDomain
// This API is used to disable the company’s domain name that obtained the ICP filing
//
// error code that may be returned:
//  FAILEDOPERATION_EXISTBLACKDOMAIN = "FailedOperation.ExistBlackDomain"
//  FAILEDOPERATION_EXISTWHITEDOMAIN = "FailedOperation.ExistWhiteDomain"
//  FAILEDOPERATION_INVALIDREQUESTENTERPRISEINFO = "FailedOperation.InvalidRequestEnterpriseInfo"
//  FAILEDOPERATION_NOACCESSPERMISSION = "FailedOperation.NoAccessPermission"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DisableTeamDomainWithContext(ctx context.Context, request *DisableTeamDomainRequest) (response *DisableTeamDomainResponse, err error) {
    if request == nil {
        request = NewDisableTeamDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableTeamDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableTeamDomainResponse()
    err = c.Send(request, response)
    return
}

func NewEnableApplicationSensitiveAPIRequest() (request *EnableApplicationSensitiveAPIRequest) {
    request = &EnableApplicationSensitiveAPIRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "EnableApplicationSensitiveAPI")
    
    
    return
}

func NewEnableApplicationSensitiveAPIResponse() (response *EnableApplicationSensitiveAPIResponse) {
    response = &EnableApplicationSensitiveAPIResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableApplicationSensitiveAPI
// This API is used to set an application sensitive API to public.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOACCESSPERMISSION = "FailedOperation.NoAccessPermission"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) EnableApplicationSensitiveAPI(request *EnableApplicationSensitiveAPIRequest) (response *EnableApplicationSensitiveAPIResponse, err error) {
    return c.EnableApplicationSensitiveAPIWithContext(context.Background(), request)
}

// EnableApplicationSensitiveAPI
// This API is used to set an application sensitive API to public.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOACCESSPERMISSION = "FailedOperation.NoAccessPermission"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) EnableApplicationSensitiveAPIWithContext(ctx context.Context, request *EnableApplicationSensitiveAPIRequest) (response *EnableApplicationSensitiveAPIResponse, err error) {
    if request == nil {
        request = NewEnableApplicationSensitiveAPIRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableApplicationSensitiveAPI require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableApplicationSensitiveAPIResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationRequest() (request *ModifyApplicationRequest) {
    request = &ModifyApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "ModifyApplication")
    
    
    return
}

func NewModifyApplicationResponse() (response *ModifyApplicationResponse) {
    response = &ModifyApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApplication
// This API is used to change application information
//
// error code that may be returned:
//  FAILEDOPERATION_APPLICATIONNAMEALREADYEXISTED = "FailedOperation.ApplicationNameAlreadyExisted"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDAPPPACKAGENAMELAYOUT = "InvalidParameterValue.InvalidAppPackageNameLayout"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONID = "InvalidParameterValue.InvalidApplicationId"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONINTRO = "InvalidParameterValue.InvalidApplicationIntro"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONNAME = "InvalidParameterValue.InvalidApplicationName"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONREMARK = "InvalidParameterValue.InvalidApplicationRemark"
//  INVALIDPARAMETERVALUE_INVALIDMNPNAME = "InvalidParameterValue.InvalidMNPName"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) ModifyApplication(request *ModifyApplicationRequest) (response *ModifyApplicationResponse, err error) {
    return c.ModifyApplicationWithContext(context.Background(), request)
}

// ModifyApplication
// This API is used to change application information
//
// error code that may be returned:
//  FAILEDOPERATION_APPLICATIONNAMEALREADYEXISTED = "FailedOperation.ApplicationNameAlreadyExisted"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDAPPPACKAGENAMELAYOUT = "InvalidParameterValue.InvalidAppPackageNameLayout"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONID = "InvalidParameterValue.InvalidApplicationId"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONINTRO = "InvalidParameterValue.InvalidApplicationIntro"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONNAME = "InvalidParameterValue.InvalidApplicationName"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONREMARK = "InvalidParameterValue.InvalidApplicationRemark"
//  INVALIDPARAMETERVALUE_INVALIDMNPNAME = "InvalidParameterValue.InvalidMNPName"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) ModifyApplicationWithContext(ctx context.Context, request *ModifyApplicationRequest) (response *ModifyApplicationResponse, err error) {
    if request == nil {
        request = NewModifyApplicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationAppKeyRequest() (request *ModifyApplicationAppKeyRequest) {
    request = &ModifyApplicationAppKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "ModifyApplicationAppKey")
    
    
    return
}

func NewModifyApplicationAppKeyResponse() (response *ModifyApplicationAppKeyResponse) {
    response = &ModifyApplicationAppKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApplicationAppKey
// This API is used to modify the application package name
//
// error code that may be returned:
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyApplicationAppKey(request *ModifyApplicationAppKeyRequest) (response *ModifyApplicationAppKeyResponse, err error) {
    return c.ModifyApplicationAppKeyWithContext(context.Background(), request)
}

// ModifyApplicationAppKey
// This API is used to modify the application package name
//
// error code that may be returned:
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyApplicationAppKeyWithContext(ctx context.Context, request *ModifyApplicationAppKeyRequest) (response *ModifyApplicationAppKeyResponse, err error) {
    if request == nil {
        request = NewModifyApplicationAppKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplicationAppKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationAppKeyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationAppUrlRequest() (request *ModifyApplicationAppUrlRequest) {
    request = &ModifyApplicationAppUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "ModifyApplicationAppUrl")
    
    
    return
}

func NewModifyApplicationAppUrlResponse() (response *ModifyApplicationAppUrlResponse) {
    response = &ModifyApplicationAppUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApplicationAppUrl
// This API is used to change the app download address
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyApplicationAppUrl(request *ModifyApplicationAppUrlRequest) (response *ModifyApplicationAppUrlResponse, err error) {
    return c.ModifyApplicationAppUrlWithContext(context.Background(), request)
}

// ModifyApplicationAppUrl
// This API is used to change the app download address
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyApplicationAppUrlWithContext(ctx context.Context, request *ModifyApplicationAppUrlRequest) (response *ModifyApplicationAppUrlResponse, err error) {
    if request == nil {
        request = NewModifyApplicationAppUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplicationAppUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationAppUrlResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDomainRequest() (request *ModifyDomainRequest) {
    request = &ModifyDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "ModifyDomain")
    
    
    return
}

func NewModifyDomainResponse() (response *ModifyDomainResponse) {
    response = &ModifyDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDomain
// This API is used to edit the mini program domain information
//
// error code that may be returned:
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
func (c *Client) ModifyDomain(request *ModifyDomainRequest) (response *ModifyDomainResponse, err error) {
    return c.ModifyDomainWithContext(context.Background(), request)
}

// ModifyDomain
// This API is used to edit the mini program domain information
//
// error code that may be returned:
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
func (c *Client) ModifyDomainWithContext(ctx context.Context, request *ModifyDomainRequest) (response *ModifyDomainResponse, err error) {
    if request == nil {
        request = NewModifyDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDomainResponse()
    err = c.Send(request, response)
    return
}

func NewModifyGlobalDomainRequest() (request *ModifyGlobalDomainRequest) {
    request = &ModifyGlobalDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "ModifyGlobalDomain")
    
    
    return
}

func NewModifyGlobalDomainResponse() (response *ModifyGlobalDomainResponse) {
    response = &ModifyGlobalDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyGlobalDomain
// This API is used to modify domain allowlist or blocklist
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETOPERATERESOURCEFAILED = "FailedOperation.GetOperateResourceFailed"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) ModifyGlobalDomain(request *ModifyGlobalDomainRequest) (response *ModifyGlobalDomainResponse, err error) {
    return c.ModifyGlobalDomainWithContext(context.Background(), request)
}

// ModifyGlobalDomain
// This API is used to modify domain allowlist or blocklist
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETOPERATERESOURCEFAILED = "FailedOperation.GetOperateResourceFailed"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) ModifyGlobalDomainWithContext(ctx context.Context, request *ModifyGlobalDomainRequest) (response *ModifyGlobalDomainResponse, err error) {
    if request == nil {
        request = NewModifyGlobalDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyGlobalDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyGlobalDomainResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMNPRequest() (request *ModifyMNPRequest) {
    request = &ModifyMNPRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "ModifyMNP")
    
    
    return
}

func NewModifyMNPResponse() (response *ModifyMNPResponse) {
    response = &ModifyMNPResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMNP
// This API is used to modify mini program information
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MINIPROGRAMICONANALYSISFAILED = "FailedOperation.MiniProgramIconAnalysisFailed"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMNPID = "InvalidParameterValue.InvalidMNPId"
//  INVALIDPARAMETERVALUE_INVALIDMNPNAME = "InvalidParameterValue.InvalidMNPName"
//  INVALIDPARAMETERVALUE_INVALIDMNPTYPE = "InvalidParameterValue.InvalidMNPType"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  INVALIDPARAMETERVALUE_MNPTYPENUMBEREXCEEDLIMIT = "InvalidParameterValue.MNPTypeNumberExceedLimit"
//  INVALIDPARAMETERVALUE_MINIPROGRAMNAMEALREADYEXIST = "InvalidParameterValue.MiniProgramNameAlreadyExist"
func (c *Client) ModifyMNP(request *ModifyMNPRequest) (response *ModifyMNPResponse, err error) {
    return c.ModifyMNPWithContext(context.Background(), request)
}

// ModifyMNP
// This API is used to modify mini program information
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MINIPROGRAMICONANALYSISFAILED = "FailedOperation.MiniProgramIconAnalysisFailed"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMNPID = "InvalidParameterValue.InvalidMNPId"
//  INVALIDPARAMETERVALUE_INVALIDMNPNAME = "InvalidParameterValue.InvalidMNPName"
//  INVALIDPARAMETERVALUE_INVALIDMNPTYPE = "InvalidParameterValue.InvalidMNPType"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  INVALIDPARAMETERVALUE_MNPTYPENUMBEREXCEEDLIMIT = "InvalidParameterValue.MNPTypeNumberExceedLimit"
//  INVALIDPARAMETERVALUE_MINIPROGRAMNAMEALREADYEXIST = "InvalidParameterValue.MiniProgramNameAlreadyExist"
func (c *Client) ModifyMNPWithContext(ctx context.Context, request *ModifyMNPRequest) (response *ModifyMNPResponse, err error) {
    if request == nil {
        request = NewModifyMNPRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMNP require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMNPResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMNPDomainRequest() (request *ModifyMNPDomainRequest) {
    request = &ModifyMNPDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "ModifyMNPDomain")
    
    
    return
}

func NewModifyMNPDomainResponse() (response *ModifyMNPDomainResponse) {
    response = &ModifyMNPDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMNPDomain
// This API is used to edit the mini program domain information.
//
// error code that may be returned:
//  FAILEDOPERATION_GETOPERATERESOURCEFAILED = "FailedOperation.GetOperateResourceFailed"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMNPID = "InvalidParameterValue.InvalidMNPId"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) ModifyMNPDomain(request *ModifyMNPDomainRequest) (response *ModifyMNPDomainResponse, err error) {
    return c.ModifyMNPDomainWithContext(context.Background(), request)
}

// ModifyMNPDomain
// This API is used to edit the mini program domain information.
//
// error code that may be returned:
//  FAILEDOPERATION_GETOPERATERESOURCEFAILED = "FailedOperation.GetOperateResourceFailed"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMNPID = "InvalidParameterValue.InvalidMNPId"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) ModifyMNPDomainWithContext(ctx context.Context, request *ModifyMNPDomainRequest) (response *ModifyMNPDomainResponse, err error) {
    if request == nil {
        request = NewModifyMNPDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMNPDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMNPDomainResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMNPStatusOfflineRequest() (request *ModifyMNPStatusOfflineRequest) {
    request = &ModifyMNPStatusOfflineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "ModifyMNPStatusOffline")
    
    
    return
}

func NewModifyMNPStatusOfflineResponse() (response *ModifyMNPStatusOfflineResponse) {
    response = &ModifyMNPStatusOfflineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMNPStatusOffline
// This API is used to remove the mini program
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDREQUESTENTERPRISEINFO = "FailedOperation.InvalidRequestEnterpriseInfo"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_NOACCESSPERMISSION = "FailedOperation.NoAccessPermission"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_USERTEAMRELATIONNOTEXIST = "InvalidParameterValue.UserTeamRelationNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyMNPStatusOffline(request *ModifyMNPStatusOfflineRequest) (response *ModifyMNPStatusOfflineResponse, err error) {
    return c.ModifyMNPStatusOfflineWithContext(context.Background(), request)
}

// ModifyMNPStatusOffline
// This API is used to remove the mini program
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDREQUESTENTERPRISEINFO = "FailedOperation.InvalidRequestEnterpriseInfo"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_NOACCESSPERMISSION = "FailedOperation.NoAccessPermission"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_USERTEAMRELATIONNOTEXIST = "InvalidParameterValue.UserTeamRelationNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyMNPStatusOfflineWithContext(ctx context.Context, request *ModifyMNPStatusOfflineRequest) (response *ModifyMNPStatusOfflineResponse, err error) {
    if request == nil {
        request = NewModifyMNPStatusOfflineRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMNPStatusOffline require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMNPStatusOfflineResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMNPVersionPreviewRequest() (request *ModifyMNPVersionPreviewRequest) {
    request = &ModifyMNPVersionPreviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "ModifyMNPVersionPreview")
    
    
    return
}

func NewModifyMNPVersionPreviewResponse() (response *ModifyMNPVersionPreviewResponse) {
    response = &ModifyMNPVersionPreviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMNPVersionPreview
// This API is used to configure the mini program preview version
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  FAILEDOPERATION_SHOWCASEVERSIONALREADYEXIST = "FailedOperation.ShowcaseVersionAlreadyExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyMNPVersionPreview(request *ModifyMNPVersionPreviewRequest) (response *ModifyMNPVersionPreviewResponse, err error) {
    return c.ModifyMNPVersionPreviewWithContext(context.Background(), request)
}

// ModifyMNPVersionPreview
// This API is used to configure the mini program preview version
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  FAILEDOPERATION_SHOWCASEVERSIONALREADYEXIST = "FailedOperation.ShowcaseVersionAlreadyExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyMNPVersionPreviewWithContext(ctx context.Context, request *ModifyMNPVersionPreviewRequest) (response *ModifyMNPVersionPreviewResponse, err error) {
    if request == nil {
        request = NewModifyMNPVersionPreviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMNPVersionPreview require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMNPVersionPreviewResponse()
    err = c.Send(request, response)
    return
}

func NewModifyOnlineVersionRequest() (request *ModifyOnlineVersionRequest) {
    request = &ModifyOnlineVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "ModifyOnlineVersion")
    
    
    return
}

func NewModifyOnlineVersionResponse() (response *ModifyOnlineVersionResponse) {
    response = &ModifyOnlineVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyOnlineVersion
// This API is used to rollback a mini program release version
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_NOACCESSPERMISSION = "FailedOperation.NoAccessPermission"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMINIAPPID = "InvalidParameterValue.InvalidMiniAppId"
//  INVALIDPARAMETERVALUE_INVALIDUSERID = "InvalidParameterValue.InvalidUserId"
//  INVALIDPARAMETERVALUE_USERTEAMRELATIONNOTEXIST = "InvalidParameterValue.UserTeamRelationNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyOnlineVersion(request *ModifyOnlineVersionRequest) (response *ModifyOnlineVersionResponse, err error) {
    return c.ModifyOnlineVersionWithContext(context.Background(), request)
}

// ModifyOnlineVersion
// This API is used to rollback a mini program release version
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_NOACCESSPERMISSION = "FailedOperation.NoAccessPermission"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMINIAPPID = "InvalidParameterValue.InvalidMiniAppId"
//  INVALIDPARAMETERVALUE_INVALIDUSERID = "InvalidParameterValue.InvalidUserId"
//  INVALIDPARAMETERVALUE_USERTEAMRELATIONNOTEXIST = "InvalidParameterValue.UserTeamRelationNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyOnlineVersionWithContext(ctx context.Context, request *ModifyOnlineVersionRequest) (response *ModifyOnlineVersionResponse, err error) {
    if request == nil {
        request = NewModifyOnlineVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyOnlineVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyOnlineVersionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPlatformAuditStatusRequest() (request *ModifyPlatformAuditStatusRequest) {
    request = &ModifyPlatformAuditStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "ModifyPlatformAuditStatus")
    
    
    return
}

func NewModifyPlatformAuditStatusResponse() (response *ModifyPlatformAuditStatusResponse) {
    response = &ModifyPlatformAuditStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPlatformAuditStatus
// This API is used to approve or reject the release of the mini program version
//
// error code that may be returned:
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyPlatformAuditStatus(request *ModifyPlatformAuditStatusRequest) (response *ModifyPlatformAuditStatusResponse, err error) {
    return c.ModifyPlatformAuditStatusWithContext(context.Background(), request)
}

// ModifyPlatformAuditStatus
// This API is used to approve or reject the release of the mini program version
//
// error code that may be returned:
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyPlatformAuditStatusWithContext(ctx context.Context, request *ModifyPlatformAuditStatusRequest) (response *ModifyPlatformAuditStatusResponse, err error) {
    if request == nil {
        request = NewModifyPlatformAuditStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPlatformAuditStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPlatformAuditStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifySensitiveAPIAuditStatusRequest() (request *ModifySensitiveAPIAuditStatusRequest) {
    request = &ModifySensitiveAPIAuditStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "ModifySensitiveAPIAuditStatus")
    
    
    return
}

func NewModifySensitiveAPIAuditStatusResponse() (response *ModifySensitiveAPIAuditStatusResponse) {
    response = &ModifySensitiveAPIAuditStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySensitiveAPIAuditStatus
// This API is used to approve or reject the sensitive API permission application
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifySensitiveAPIAuditStatus(request *ModifySensitiveAPIAuditStatusRequest) (response *ModifySensitiveAPIAuditStatusResponse, err error) {
    return c.ModifySensitiveAPIAuditStatusWithContext(context.Background(), request)
}

// ModifySensitiveAPIAuditStatus
// This API is used to approve or reject the sensitive API permission application
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifySensitiveAPIAuditStatusWithContext(ctx context.Context, request *ModifySensitiveAPIAuditStatusRequest) (response *ModifySensitiveAPIAuditStatusResponse, err error) {
    if request == nil {
        request = NewModifySensitiveAPIAuditStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySensitiveAPIAuditStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySensitiveAPIAuditStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTeamRequest() (request *ModifyTeamRequest) {
    request = &ModifyTeamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "ModifyTeam")
    
    
    return
}

func NewModifyTeamResponse() (response *ModifyTeamResponse) {
    response = &ModifyTeamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTeam
// This API is used to change team information
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EXISTTEAMNAME = "InvalidParameterValue.ExistTeamName"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  INVALIDPARAMETERVALUE_INVALIDTEAMID = "InvalidParameterValue.InvalidTeamId"
func (c *Client) ModifyTeam(request *ModifyTeamRequest) (response *ModifyTeamResponse, err error) {
    return c.ModifyTeamWithContext(context.Background(), request)
}

// ModifyTeam
// This API is used to change team information
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EXISTTEAMNAME = "InvalidParameterValue.ExistTeamName"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  INVALIDPARAMETERVALUE_INVALIDTEAMID = "InvalidParameterValue.InvalidTeamId"
func (c *Client) ModifyTeamWithContext(ctx context.Context, request *ModifyTeamRequest) (response *ModifyTeamResponse, err error) {
    if request == nil {
        request = NewModifyTeamRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTeam require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTeamResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTeamMemberRequest() (request *ModifyTeamMemberRequest) {
    request = &ModifyTeamMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "ModifyTeamMember")
    
    
    return
}

func NewModifyTeamMemberResponse() (response *ModifyTeamMemberResponse) {
    response = &ModifyTeamMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTeamMember
// This API is used to modify team member information.
//
// error code that may be returned:
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDROLEID = "InvalidParameterValue.InvalidRoleId"
//  INVALIDPARAMETERVALUE_INVALIDTEAMID = "InvalidParameterValue.InvalidTeamId"
//  INVALIDPARAMETERVALUE_INVALIDUSERID = "InvalidParameterValue.InvalidUserId"
func (c *Client) ModifyTeamMember(request *ModifyTeamMemberRequest) (response *ModifyTeamMemberResponse, err error) {
    return c.ModifyTeamMemberWithContext(context.Background(), request)
}

// ModifyTeamMember
// This API is used to modify team member information.
//
// error code that may be returned:
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDROLEID = "InvalidParameterValue.InvalidRoleId"
//  INVALIDPARAMETERVALUE_INVALIDTEAMID = "InvalidParameterValue.InvalidTeamId"
//  INVALIDPARAMETERVALUE_INVALIDUSERID = "InvalidParameterValue.InvalidUserId"
func (c *Client) ModifyTeamMemberWithContext(ctx context.Context, request *ModifyTeamMemberRequest) (response *ModifyTeamMemberResponse, err error) {
    if request == nil {
        request = NewModifyTeamMemberRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTeamMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTeamMemberResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserRequest() (request *ModifyUserRequest) {
    request = &ModifyUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "ModifyUser")
    
    
    return
}

func NewModifyUserResponse() (response *ModifyUserResponse) {
    response = &ModifyUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUser
// This API is used to modify user information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  FAILEDOPERATION_UNABLEOPERATEADMINACCOUNT = "FailedOperation.UnableOperateAdminAccount"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTTYPE = "InvalidParameterValue.InvalidAccountType"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  INVALIDPARAMETERVALUE_INVALIDUSERNAME = "InvalidParameterValue.InvalidUserName"
func (c *Client) ModifyUser(request *ModifyUserRequest) (response *ModifyUserResponse, err error) {
    return c.ModifyUserWithContext(context.Background(), request)
}

// ModifyUser
// This API is used to modify user information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  FAILEDOPERATION_UNABLEOPERATEADMINACCOUNT = "FailedOperation.UnableOperateAdminAccount"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTTYPE = "InvalidParameterValue.InvalidAccountType"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  INVALIDPARAMETERVALUE_INVALIDUSERNAME = "InvalidParameterValue.InvalidUserName"
func (c *Client) ModifyUserWithContext(ctx context.Context, request *ModifyUserRequest) (response *ModifyUserResponse, err error) {
    if request == nil {
        request = NewModifyUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserPasswordRequest() (request *ModifyUserPasswordRequest) {
    request = &ModifyUserPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "ModifyUserPassword")
    
    
    return
}

func NewModifyUserPasswordResponse() (response *ModifyUserPasswordResponse) {
    response = &ModifyUserPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUserPassword
// This API is used to reset user password
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
func (c *Client) ModifyUserPassword(request *ModifyUserPasswordRequest) (response *ModifyUserPasswordResponse, err error) {
    return c.ModifyUserPasswordWithContext(context.Background(), request)
}

// ModifyUserPassword
// This API is used to reset user password
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAUTHFAILED = "FailedOperation.LoginAuthFailed"
//  FAILEDOPERATION_REQUESTPARAMANALYSISFAILED = "FailedOperation.RequestParamAnalysisFailed"
//  MISSINGPARAMETER_LOGININFONOTFOUND = "MissingParameter.LoginInfoNotFound"
func (c *Client) ModifyUserPasswordWithContext(ctx context.Context, request *ModifyUserPasswordRequest) (response *ModifyUserPasswordResponse, err error) {
    if request == nil {
        request = NewModifyUserPasswordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUserPassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewProcessMNPApprovalRequest() (request *ProcessMNPApprovalRequest) {
    request = &ProcessMNPApprovalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "ProcessMNPApproval")
    
    
    return
}

func NewProcessMNPApprovalResponse() (response *ProcessMNPApprovalResponse) {
    response = &ProcessMNPApprovalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ProcessMNPApproval
// This API is used to approve or reject the release of a mini program version.
//
// error code that may be returned:
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDAPPROVALNO = "InvalidParameterValue.InvalidApprovalNo"
//  INVALIDPARAMETERVALUE_INVALIDAPPROVALNOTE = "InvalidParameterValue.InvalidApprovalNote"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  INVALIDPARAMETERVALUE_PROCESSAPPLICATIONANDAPPROVALAPPLICATIONMISMATCH = "InvalidParameterValue.ProcessApplicationAndApprovalApplicationMismatch"
//  INVALIDPARAMETERVALUE_PROCESSAPPROVALITEMISEMPTY = "InvalidParameterValue.ProcessApprovalItemIsEmpty"
//  INVALIDPARAMETERVALUE_PROCESSMNPAPPROVALSTATUSERROR = "InvalidParameterValue.ProcessMNPApprovalStatusError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ProcessMNPApproval(request *ProcessMNPApprovalRequest) (response *ProcessMNPApprovalResponse, err error) {
    return c.ProcessMNPApprovalWithContext(context.Background(), request)
}

// ProcessMNPApproval
// This API is used to approve or reject the release of a mini program version.
//
// error code that may be returned:
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDAPPROVALNO = "InvalidParameterValue.InvalidApprovalNo"
//  INVALIDPARAMETERVALUE_INVALIDAPPROVALNOTE = "InvalidParameterValue.InvalidApprovalNote"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  INVALIDPARAMETERVALUE_PROCESSAPPLICATIONANDAPPROVALAPPLICATIONMISMATCH = "InvalidParameterValue.ProcessApplicationAndApprovalApplicationMismatch"
//  INVALIDPARAMETERVALUE_PROCESSAPPROVALITEMISEMPTY = "InvalidParameterValue.ProcessApprovalItemIsEmpty"
//  INVALIDPARAMETERVALUE_PROCESSMNPAPPROVALSTATUSERROR = "InvalidParameterValue.ProcessMNPApprovalStatusError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ProcessMNPApprovalWithContext(ctx context.Context, request *ProcessMNPApprovalRequest) (response *ProcessMNPApprovalResponse, err error) {
    if request == nil {
        request = NewProcessMNPApprovalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ProcessMNPApproval require credential")
    }

    request.SetContext(ctx)
    
    response = NewProcessMNPApprovalResponse()
    err = c.Send(request, response)
    return
}

func NewProcessMNPSensitiveAPIPermissionApprovalRequest() (request *ProcessMNPSensitiveAPIPermissionApprovalRequest) {
    request = &ProcessMNPSensitiveAPIPermissionApprovalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "ProcessMNPSensitiveAPIPermissionApproval")
    
    
    return
}

func NewProcessMNPSensitiveAPIPermissionApprovalResponse() (response *ProcessMNPSensitiveAPIPermissionApprovalResponse) {
    response = &ProcessMNPSensitiveAPIPermissionApprovalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ProcessMNPSensitiveAPIPermissionApproval
// This API is used to approve or reject the sensitive API permission application. 
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) ProcessMNPSensitiveAPIPermissionApproval(request *ProcessMNPSensitiveAPIPermissionApprovalRequest) (response *ProcessMNPSensitiveAPIPermissionApprovalResponse, err error) {
    return c.ProcessMNPSensitiveAPIPermissionApprovalWithContext(context.Background(), request)
}

// ProcessMNPSensitiveAPIPermissionApproval
// This API is used to approve or reject the sensitive API permission application. 
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) ProcessMNPSensitiveAPIPermissionApprovalWithContext(ctx context.Context, request *ProcessMNPSensitiveAPIPermissionApprovalRequest) (response *ProcessMNPSensitiveAPIPermissionApprovalResponse, err error) {
    if request == nil {
        request = NewProcessMNPSensitiveAPIPermissionApprovalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ProcessMNPSensitiveAPIPermissionApproval require credential")
    }

    request.SetContext(ctx)
    
    response = NewProcessMNPSensitiveAPIPermissionApprovalResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseMNPVersionRequest() (request *ReleaseMNPVersionRequest) {
    request = &ReleaseMNPVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "ReleaseMNPVersion")
    
    
    return
}

func NewReleaseMNPVersionResponse() (response *ReleaseMNPVersionResponse) {
    response = &ReleaseMNPVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReleaseMNPVersion
// This API is used to release a mini program version.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MNPRELEASENUMBEREXCEEDLIMIT = "FailedOperation.MNPReleaseNumberExceedLimit"
//  FAILEDOPERATION_MNPVERSIONRELEASED = "FailedOperation.MNPVersionReleased"
//  FAILEDOPERATION_ONLYRELEASEPLATFORMAPPROVEDMNPVERSION = "FailedOperation.OnlyReleasePlatformApprovedMNPVersion"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  FAILEDOPERATION_RELEASEMNPVERSIONNEEDGREATERTHANONLINEVERSION = "FailedOperation.ReleaseMNPVersionNeedGreaterThanOnlineVersion"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMNPVERSIONID = "InvalidParameterValue.InvalidMNPVersionId"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ReleaseMNPVersion(request *ReleaseMNPVersionRequest) (response *ReleaseMNPVersionResponse, err error) {
    return c.ReleaseMNPVersionWithContext(context.Background(), request)
}

// ReleaseMNPVersion
// This API is used to release a mini program version.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MNPRELEASENUMBEREXCEEDLIMIT = "FailedOperation.MNPReleaseNumberExceedLimit"
//  FAILEDOPERATION_MNPVERSIONRELEASED = "FailedOperation.MNPVersionReleased"
//  FAILEDOPERATION_ONLYRELEASEPLATFORMAPPROVEDMNPVERSION = "FailedOperation.OnlyReleasePlatformApprovedMNPVersion"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  FAILEDOPERATION_RELEASEMNPVERSIONNEEDGREATERTHANONLINEVERSION = "FailedOperation.ReleaseMNPVersionNeedGreaterThanOnlineVersion"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMNPVERSIONID = "InvalidParameterValue.InvalidMNPVersionId"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ReleaseMNPVersionWithContext(ctx context.Context, request *ReleaseMNPVersionRequest) (response *ReleaseMNPVersionResponse, err error) {
    if request == nil {
        request = NewReleaseMNPVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReleaseMNPVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewReleaseMNPVersionResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveMNPRequest() (request *RemoveMNPRequest) {
    request = &RemoveMNPRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "RemoveMNP")
    
    
    return
}

func NewRemoveMNPResponse() (response *RemoveMNPResponse) {
    response = &RemoveMNPResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveMNP
// This API is used to remove a mini program.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMNPID = "InvalidParameterValue.InvalidMNPId"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) RemoveMNP(request *RemoveMNPRequest) (response *RemoveMNPResponse, err error) {
    return c.RemoveMNPWithContext(context.Background(), request)
}

// RemoveMNP
// This API is used to remove a mini program.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMNPID = "InvalidParameterValue.InvalidMNPId"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) RemoveMNPWithContext(ctx context.Context, request *RemoveMNPRequest) (response *RemoveMNPResponse, err error) {
    if request == nil {
        request = NewRemoveMNPRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveMNP require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveMNPResponse()
    err = c.Send(request, response)
    return
}

func NewRollbackMNPVersionRequest() (request *RollbackMNPVersionRequest) {
    request = &RollbackMNPVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcmpp", APIVersion, "RollbackMNPVersion")
    
    
    return
}

func NewRollbackMNPVersionResponse() (response *RollbackMNPVersionResponse) {
    response = &RollbackMNPVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RollbackMNPVersion
// This API is used to rollback a mini program online version.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMINIAPPID = "InvalidParameterValue.InvalidMiniAppId"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) RollbackMNPVersion(request *RollbackMNPVersionRequest) (response *RollbackMNPVersionResponse, err error) {
    return c.RollbackMNPVersionWithContext(context.Background(), request)
}

// RollbackMNPVersion
// This API is used to rollback a mini program online version.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMINIAPPID = "InvalidParameterValue.InvalidMiniAppId"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) RollbackMNPVersionWithContext(ctx context.Context, request *RollbackMNPVersionRequest) (response *RollbackMNPVersionResponse, err error) {
    if request == nil {
        request = NewRollbackMNPVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RollbackMNPVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewRollbackMNPVersionResponse()
    err = c.Send(request, response)
    return
}
