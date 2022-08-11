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

package v20180813

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2018-08-13"

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


func NewAssumeRoleRequest() (request *AssumeRoleRequest) {
    request = &AssumeRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sts", APIVersion, "AssumeRole")
    
    
    return
}

func NewAssumeRoleResponse() (response *AssumeRoleResponse) {
    response = &AssumeRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssumeRole
// This API is used to request for the temporary security credentials of a role.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ENCRYPTERROR = "InternalError.EncryptError"
//  INTERNALERROR_GETAPPIDERROR = "InternalError.GetAppIdError"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_GETSEEDTOKENERROR = "InternalError.GetSeedTokenError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INTERNALERROR_PBSERIALIZEERROR = "InternalError.PbSerializeError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_ACCOUNTNOTAVALIABLE = "InvalidParameter.AccountNotAvaliable"
//  INVALIDPARAMETER_EXTENDSTRATEGYOVERSIZE = "InvalidParameter.ExtendStrategyOverSize"
//  INVALIDPARAMETER_GRANTOTHERRESOURCE = "InvalidParameter.GrantOtherResource"
//  INVALIDPARAMETER_OVERLIMIT = "InvalidParameter.OverLimit"
//  INVALIDPARAMETER_OVERTIMEERROR = "InvalidParameter.OverTimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYTOOLONG = "InvalidParameter.PolicyTooLong"
//  INVALIDPARAMETER_RESOUCEERROR = "InvalidParameter.ResouceError"
//  INVALIDPARAMETER_STRATEGYFORMATERROR = "InvalidParameter.StrategyFormatError"
//  INVALIDPARAMETER_STRATEGYINVALID = "InvalidParameter.StrategyInvalid"
//  INVALIDPARAMETER_TEMPCODENOTAVALIABLE = "InvalidParameter.TempCodeNotAvaliable"
//  RESOURCENOTFOUND_ROLENOTFOUND = "ResourceNotFound.RoleNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AssumeRole(request *AssumeRoleRequest) (response *AssumeRoleResponse, err error) {
    return c.AssumeRoleWithContext(context.Background(), request)
}

// AssumeRole
// This API is used to request for the temporary security credentials of a role.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ENCRYPTERROR = "InternalError.EncryptError"
//  INTERNALERROR_GETAPPIDERROR = "InternalError.GetAppIdError"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_GETSEEDTOKENERROR = "InternalError.GetSeedTokenError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INTERNALERROR_PBSERIALIZEERROR = "InternalError.PbSerializeError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_ACCOUNTNOTAVALIABLE = "InvalidParameter.AccountNotAvaliable"
//  INVALIDPARAMETER_EXTENDSTRATEGYOVERSIZE = "InvalidParameter.ExtendStrategyOverSize"
//  INVALIDPARAMETER_GRANTOTHERRESOURCE = "InvalidParameter.GrantOtherResource"
//  INVALIDPARAMETER_OVERLIMIT = "InvalidParameter.OverLimit"
//  INVALIDPARAMETER_OVERTIMEERROR = "InvalidParameter.OverTimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYTOOLONG = "InvalidParameter.PolicyTooLong"
//  INVALIDPARAMETER_RESOUCEERROR = "InvalidParameter.ResouceError"
//  INVALIDPARAMETER_STRATEGYFORMATERROR = "InvalidParameter.StrategyFormatError"
//  INVALIDPARAMETER_STRATEGYINVALID = "InvalidParameter.StrategyInvalid"
//  INVALIDPARAMETER_TEMPCODENOTAVALIABLE = "InvalidParameter.TempCodeNotAvaliable"
//  RESOURCENOTFOUND_ROLENOTFOUND = "ResourceNotFound.RoleNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AssumeRoleWithContext(ctx context.Context, request *AssumeRoleRequest) (response *AssumeRoleResponse, err error) {
    if request == nil {
        request = NewAssumeRoleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssumeRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewAssumeRoleResponse()
    err = c.Send(request, response)
    return
}

func NewAssumeRoleWithSAMLRequest() (request *AssumeRoleWithSAMLRequest) {
    request = &AssumeRoleWithSAMLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sts", APIVersion, "AssumeRoleWithSAML")
    
    
    return
}

func NewAssumeRoleWithSAMLResponse() (response *AssumeRoleWithSAMLResponse) {
    response = &AssumeRoleWithSAMLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssumeRoleWithSAML
// This API is used to request for the temporary credentials for a role that has been authenticated via a SAML assertion.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ENCRYPTERROR = "InternalError.EncryptError"
//  INTERNALERROR_GETAPPIDERROR = "InternalError.GetAppIdError"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_GETSEEDTOKENERROR = "InternalError.GetSeedTokenError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INTERNALERROR_PBSERIALIZEERROR = "InternalError.PbSerializeError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_ACCOUNTNOTAVALIABLE = "InvalidParameter.AccountNotAvaliable"
//  INVALIDPARAMETER_EXTENDSTRATEGYOVERSIZE = "InvalidParameter.ExtendStrategyOverSize"
//  INVALIDPARAMETER_GRANTOTHERRESOURCE = "InvalidParameter.GrantOtherResource"
//  INVALIDPARAMETER_OVERTIMEERROR = "InvalidParameter.OverTimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYTOOLONG = "InvalidParameter.PolicyTooLong"
//  INVALIDPARAMETER_RESOUCEERROR = "InvalidParameter.ResouceError"
//  INVALIDPARAMETER_STRATEGYFORMATERROR = "InvalidParameter.StrategyFormatError"
//  INVALIDPARAMETER_STRATEGYINVALID = "InvalidParameter.StrategyInvalid"
//  INVALIDPARAMETER_TEMPCODENOTAVALIABLE = "InvalidParameter.TempCodeNotAvaliable"
//  RESOURCENOTFOUND_ROLENOTFOUND = "ResourceNotFound.RoleNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AssumeRoleWithSAML(request *AssumeRoleWithSAMLRequest) (response *AssumeRoleWithSAMLResponse, err error) {
    return c.AssumeRoleWithSAMLWithContext(context.Background(), request)
}

// AssumeRoleWithSAML
// This API is used to request for the temporary credentials for a role that has been authenticated via a SAML assertion.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ENCRYPTERROR = "InternalError.EncryptError"
//  INTERNALERROR_GETAPPIDERROR = "InternalError.GetAppIdError"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_GETSEEDTOKENERROR = "InternalError.GetSeedTokenError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INTERNALERROR_PBSERIALIZEERROR = "InternalError.PbSerializeError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_ACCOUNTNOTAVALIABLE = "InvalidParameter.AccountNotAvaliable"
//  INVALIDPARAMETER_EXTENDSTRATEGYOVERSIZE = "InvalidParameter.ExtendStrategyOverSize"
//  INVALIDPARAMETER_GRANTOTHERRESOURCE = "InvalidParameter.GrantOtherResource"
//  INVALIDPARAMETER_OVERTIMEERROR = "InvalidParameter.OverTimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYTOOLONG = "InvalidParameter.PolicyTooLong"
//  INVALIDPARAMETER_RESOUCEERROR = "InvalidParameter.ResouceError"
//  INVALIDPARAMETER_STRATEGYFORMATERROR = "InvalidParameter.StrategyFormatError"
//  INVALIDPARAMETER_STRATEGYINVALID = "InvalidParameter.StrategyInvalid"
//  INVALIDPARAMETER_TEMPCODENOTAVALIABLE = "InvalidParameter.TempCodeNotAvaliable"
//  RESOURCENOTFOUND_ROLENOTFOUND = "ResourceNotFound.RoleNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AssumeRoleWithSAMLWithContext(ctx context.Context, request *AssumeRoleWithSAMLRequest) (response *AssumeRoleWithSAMLResponse, err error) {
    if request == nil {
        request = NewAssumeRoleWithSAMLRequest()
    }
    
    request.SetContext(ctx)
    
    response = NewAssumeRoleWithSAMLResponse()
    err = c.Send(request, response)
    return
}

func NewAssumeRoleWithWebIdentityRequest() (request *AssumeRoleWithWebIdentityRequest) {
    request = &AssumeRoleWithWebIdentityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sts", APIVersion, "AssumeRoleWithWebIdentity")
    
    
    return
}

func NewAssumeRoleWithWebIdentityResponse() (response *AssumeRoleWithWebIdentityResponse) {
    response = &AssumeRoleWithWebIdentityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssumeRoleWithWebIdentity
// This API is used to apply for an OIDC role credential.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_WEBIDENTITYTOKENERROR = "InvalidParameter.WebIdentityTokenError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AssumeRoleWithWebIdentity(request *AssumeRoleWithWebIdentityRequest) (response *AssumeRoleWithWebIdentityResponse, err error) {
    return c.AssumeRoleWithWebIdentityWithContext(context.Background(), request)
}

// AssumeRoleWithWebIdentity
// This API is used to apply for an OIDC role credential.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_WEBIDENTITYTOKENERROR = "InvalidParameter.WebIdentityTokenError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AssumeRoleWithWebIdentityWithContext(ctx context.Context, request *AssumeRoleWithWebIdentityRequest) (response *AssumeRoleWithWebIdentityResponse, err error) {
    if request == nil {
        request = NewAssumeRoleWithWebIdentityRequest()
    }
    
    request.SetContext(ctx)
    
    response = NewAssumeRoleWithWebIdentityResponse()
    err = c.Send(request, response)
    return
}

func NewGetCallerIdentityRequest() (request *GetCallerIdentityRequest) {
    request = &GetCallerIdentityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sts", APIVersion, "GetCallerIdentity")
    
    
    return
}

func NewGetCallerIdentityResponse() (response *GetCallerIdentityResponse) {
    response = &GetCallerIdentityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetCallerIdentity
// This API is used to get the identity information of the current caller.
//
// The persistent keys of the root account and sub-account as well as the temporary credentials generated by `AssumeRole` and `GetFederationToken` can be used to get the identity information.
//
// error code that may be returned:
//  AUTHFAILURE_ACCESSKEYILLEGAL = "AuthFailure.AccessKeyIllegal"
//  INTERNALERROR_GETSEEDTOKENERROR = "InternalError.GetSeedTokenError"
//  INVALIDPARAMETER_ACCESSKEYNOTSUPPORT = "InvalidParameter.AccessKeyNotSupport"
func (c *Client) GetCallerIdentity(request *GetCallerIdentityRequest) (response *GetCallerIdentityResponse, err error) {
    return c.GetCallerIdentityWithContext(context.Background(), request)
}

// GetCallerIdentity
// This API is used to get the identity information of the current caller.
//
// The persistent keys of the root account and sub-account as well as the temporary credentials generated by `AssumeRole` and `GetFederationToken` can be used to get the identity information.
//
// error code that may be returned:
//  AUTHFAILURE_ACCESSKEYILLEGAL = "AuthFailure.AccessKeyIllegal"
//  INTERNALERROR_GETSEEDTOKENERROR = "InternalError.GetSeedTokenError"
//  INVALIDPARAMETER_ACCESSKEYNOTSUPPORT = "InvalidParameter.AccessKeyNotSupport"
func (c *Client) GetCallerIdentityWithContext(ctx context.Context, request *GetCallerIdentityRequest) (response *GetCallerIdentityResponse, err error) {
    if request == nil {
        request = NewGetCallerIdentityRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetCallerIdentity require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetCallerIdentityResponse()
    err = c.Send(request, response)
    return
}

func NewGetFederationTokenRequest() (request *GetFederationTokenRequest) {
    request = &GetFederationTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sts", APIVersion, "GetFederationToken")
    
    
    return
}

func NewGetFederationTokenResponse() (response *GetFederationTokenResponse) {
    response = &GetFederationTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetFederationToken
// This API is used to get temporary credentials for a federated user.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ENCRYPTERROR = "InternalError.EncryptError"
//  INTERNALERROR_GETAPPIDERROR = "InternalError.GetAppIdError"
//  INTERNALERROR_GETSEEDTOKENERROR = "InternalError.GetSeedTokenError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACCOUNTNOTAVALIABLE = "InvalidParameter.AccountNotAvaliable"
//  INVALIDPARAMETER_EXTENDSTRATEGYOVERSIZE = "InvalidParameter.ExtendStrategyOverSize"
//  INVALIDPARAMETER_GRANTOTHERRESOURCE = "InvalidParameter.GrantOtherResource"
//  INVALIDPARAMETER_OVERTIMEERROR = "InvalidParameter.OverTimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYTOOLONG = "InvalidParameter.PolicyTooLong"
//  INVALIDPARAMETER_RESOUCEERROR = "InvalidParameter.ResouceError"
//  INVALIDPARAMETER_STRATEGYFORMATERROR = "InvalidParameter.StrategyFormatError"
//  INVALIDPARAMETER_STRATEGYINVALID = "InvalidParameter.StrategyInvalid"
//  INVALIDPARAMETER_TEMPCODENOTAVALIABLE = "InvalidParameter.TempCodeNotAvaliable"
//  RESOURCENOTFOUND_ROLENOTFOUND = "ResourceNotFound.RoleNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) GetFederationToken(request *GetFederationTokenRequest) (response *GetFederationTokenResponse, err error) {
    return c.GetFederationTokenWithContext(context.Background(), request)
}

// GetFederationToken
// This API is used to get temporary credentials for a federated user.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ENCRYPTERROR = "InternalError.EncryptError"
//  INTERNALERROR_GETAPPIDERROR = "InternalError.GetAppIdError"
//  INTERNALERROR_GETSEEDTOKENERROR = "InternalError.GetSeedTokenError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACCOUNTNOTAVALIABLE = "InvalidParameter.AccountNotAvaliable"
//  INVALIDPARAMETER_EXTENDSTRATEGYOVERSIZE = "InvalidParameter.ExtendStrategyOverSize"
//  INVALIDPARAMETER_GRANTOTHERRESOURCE = "InvalidParameter.GrantOtherResource"
//  INVALIDPARAMETER_OVERTIMEERROR = "InvalidParameter.OverTimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYTOOLONG = "InvalidParameter.PolicyTooLong"
//  INVALIDPARAMETER_RESOUCEERROR = "InvalidParameter.ResouceError"
//  INVALIDPARAMETER_STRATEGYFORMATERROR = "InvalidParameter.StrategyFormatError"
//  INVALIDPARAMETER_STRATEGYINVALID = "InvalidParameter.StrategyInvalid"
//  INVALIDPARAMETER_TEMPCODENOTAVALIABLE = "InvalidParameter.TempCodeNotAvaliable"
//  RESOURCENOTFOUND_ROLENOTFOUND = "ResourceNotFound.RoleNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) GetFederationTokenWithContext(ctx context.Context, request *GetFederationTokenRequest) (response *GetFederationTokenResponse, err error) {
    if request == nil {
        request = NewGetFederationTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetFederationToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetFederationTokenResponse()
    err = c.Send(request, response)
    return
}
