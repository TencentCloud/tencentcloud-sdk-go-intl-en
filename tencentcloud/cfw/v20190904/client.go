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

package v20190904

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2019-09-04"

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


func NewAddAcRuleRequest() (request *AddAcRuleRequest) {
    request = &AddAcRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "AddAcRule")
    
    
    return
}

func NewAddAcRuleResponse() (response *AddAcRuleResponse) {
    response = &AddAcRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddAcRule
// This API is used to add edge firewall rules.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddAcRule(request *AddAcRuleRequest) (response *AddAcRuleResponse, err error) {
    return c.AddAcRuleWithContext(context.Background(), request)
}

// AddAcRule
// This API is used to add edge firewall rules.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddAcRuleWithContext(ctx context.Context, request *AddAcRuleRequest) (response *AddAcRuleResponse, err error) {
    if request == nil {
        request = NewAddAcRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "AddAcRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddAcRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddAcRuleResponse()
    err = c.Send(request, response)
    return
}

func NewAddEnterpriseSecurityGroupRulesRequest() (request *AddEnterpriseSecurityGroupRulesRequest) {
    request = &AddEnterpriseSecurityGroupRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "AddEnterpriseSecurityGroupRules")
    
    
    return
}

func NewAddEnterpriseSecurityGroupRulesResponse() (response *AddEnterpriseSecurityGroupRulesResponse) {
    response = &AddEnterpriseSecurityGroupRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddEnterpriseSecurityGroupRules
// This API is used to create enterprise security group rules (new).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddEnterpriseSecurityGroupRules(request *AddEnterpriseSecurityGroupRulesRequest) (response *AddEnterpriseSecurityGroupRulesResponse, err error) {
    return c.AddEnterpriseSecurityGroupRulesWithContext(context.Background(), request)
}

// AddEnterpriseSecurityGroupRules
// This API is used to create enterprise security group rules (new).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddEnterpriseSecurityGroupRulesWithContext(ctx context.Context, request *AddEnterpriseSecurityGroupRulesRequest) (response *AddEnterpriseSecurityGroupRulesResponse, err error) {
    if request == nil {
        request = NewAddEnterpriseSecurityGroupRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "AddEnterpriseSecurityGroupRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddEnterpriseSecurityGroupRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddEnterpriseSecurityGroupRulesResponse()
    err = c.Send(request, response)
    return
}

func NewAddNatAcRuleRequest() (request *AddNatAcRuleRequest) {
    request = &AddNatAcRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "AddNatAcRule")
    
    
    return
}

func NewAddNatAcRuleResponse() (response *AddNatAcRuleResponse) {
    response = &AddNatAcRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddNatAcRule
// This API is used to add NAT access control rules.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddNatAcRule(request *AddNatAcRuleRequest) (response *AddNatAcRuleResponse, err error) {
    return c.AddNatAcRuleWithContext(context.Background(), request)
}

// AddNatAcRule
// This API is used to add NAT access control rules.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddNatAcRuleWithContext(ctx context.Context, request *AddNatAcRuleRequest) (response *AddNatAcRuleResponse, err error) {
    if request == nil {
        request = NewAddNatAcRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "AddNatAcRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddNatAcRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddNatAcRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAcRulesRequest() (request *CreateAcRulesRequest) {
    request = &CreateAcRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "CreateAcRules")
    
    
    return
}

func NewCreateAcRulesResponse() (response *CreateAcRulesResponse) {
    response = &CreateAcRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAcRules
// This API is used to create access control rules.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateAcRules(request *CreateAcRulesRequest) (response *CreateAcRulesResponse, err error) {
    return c.CreateAcRulesWithContext(context.Background(), request)
}

// CreateAcRules
// This API is used to create access control rules.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateAcRulesWithContext(ctx context.Context, request *CreateAcRulesRequest) (response *CreateAcRulesResponse, err error) {
    if request == nil {
        request = NewCreateAcRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "CreateAcRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAcRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAcRulesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNatFwInstanceRequest() (request *CreateNatFwInstanceRequest) {
    request = &CreateNatFwInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "CreateNatFwInstance")
    
    
    return
}

func NewCreateNatFwInstanceResponse() (response *CreateNatFwInstanceResponse) {
    response = &CreateNatFwInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateNatFwInstance
// This API is used to create a NAT firewall instance (The Region parameter is required).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateNatFwInstance(request *CreateNatFwInstanceRequest) (response *CreateNatFwInstanceResponse, err error) {
    return c.CreateNatFwInstanceWithContext(context.Background(), request)
}

// CreateNatFwInstance
// This API is used to create a NAT firewall instance (The Region parameter is required).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateNatFwInstanceWithContext(ctx context.Context, request *CreateNatFwInstanceRequest) (response *CreateNatFwInstanceResponse, err error) {
    if request == nil {
        request = NewCreateNatFwInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "CreateNatFwInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNatFwInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNatFwInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNatFwInstanceWithDomainRequest() (request *CreateNatFwInstanceWithDomainRequest) {
    request = &CreateNatFwInstanceWithDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "CreateNatFwInstanceWithDomain")
    
    
    return
}

func NewCreateNatFwInstanceWithDomainResponse() (response *CreateNatFwInstanceWithDomainResponse) {
    response = &CreateNatFwInstanceWithDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateNatFwInstanceWithDomain
// This API is used to create a firewall instance with domain name (The Region parameter is required).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateNatFwInstanceWithDomain(request *CreateNatFwInstanceWithDomainRequest) (response *CreateNatFwInstanceWithDomainResponse, err error) {
    return c.CreateNatFwInstanceWithDomainWithContext(context.Background(), request)
}

// CreateNatFwInstanceWithDomain
// This API is used to create a firewall instance with domain name (The Region parameter is required).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateNatFwInstanceWithDomainWithContext(ctx context.Context, request *CreateNatFwInstanceWithDomainRequest) (response *CreateNatFwInstanceWithDomainResponse, err error) {
    if request == nil {
        request = NewCreateNatFwInstanceWithDomainRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "CreateNatFwInstanceWithDomain")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNatFwInstanceWithDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNatFwInstanceWithDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSecurityGroupRulesRequest() (request *CreateSecurityGroupRulesRequest) {
    request = &CreateSecurityGroupRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "CreateSecurityGroupRules")
    
    
    return
}

func NewCreateSecurityGroupRulesResponse() (response *CreateSecurityGroupRulesResponse) {
    response = &CreateSecurityGroupRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSecurityGroupRules
// This API is used to create enterprise security group rules.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSecurityGroupRules(request *CreateSecurityGroupRulesRequest) (response *CreateSecurityGroupRulesResponse, err error) {
    return c.CreateSecurityGroupRulesWithContext(context.Background(), request)
}

// CreateSecurityGroupRules
// This API is used to create enterprise security group rules.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSecurityGroupRulesWithContext(ctx context.Context, request *CreateSecurityGroupRulesRequest) (response *CreateSecurityGroupRulesResponse, err error) {
    if request == nil {
        request = NewCreateSecurityGroupRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "CreateSecurityGroupRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSecurityGroupRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSecurityGroupRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAcRuleRequest() (request *DeleteAcRuleRequest) {
    request = &DeleteAcRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DeleteAcRule")
    
    
    return
}

func NewDeleteAcRuleResponse() (response *DeleteAcRuleResponse) {
    response = &DeleteAcRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAcRule
// This API is used to delete a rule.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteAcRule(request *DeleteAcRuleRequest) (response *DeleteAcRuleResponse, err error) {
    return c.DeleteAcRuleWithContext(context.Background(), request)
}

// DeleteAcRule
// This API is used to delete a rule.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteAcRuleWithContext(ctx context.Context, request *DeleteAcRuleRequest) (response *DeleteAcRuleResponse, err error) {
    if request == nil {
        request = NewDeleteAcRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "DeleteAcRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAcRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAcRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAllAccessControlRuleRequest() (request *DeleteAllAccessControlRuleRequest) {
    request = &DeleteAllAccessControlRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DeleteAllAccessControlRule")
    
    
    return
}

func NewDeleteAllAccessControlRuleResponse() (response *DeleteAllAccessControlRuleResponse) {
    response = &DeleteAllAccessControlRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAllAccessControlRule
// This API is used to delete all rules.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteAllAccessControlRule(request *DeleteAllAccessControlRuleRequest) (response *DeleteAllAccessControlRuleResponse, err error) {
    return c.DeleteAllAccessControlRuleWithContext(context.Background(), request)
}

// DeleteAllAccessControlRule
// This API is used to delete all rules.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteAllAccessControlRuleWithContext(ctx context.Context, request *DeleteAllAccessControlRuleRequest) (response *DeleteAllAccessControlRuleResponse, err error) {
    if request == nil {
        request = NewDeleteAllAccessControlRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "DeleteAllAccessControlRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAllAccessControlRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAllAccessControlRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteResourceGroupRequest() (request *DeleteResourceGroupRequest) {
    request = &DeleteResourceGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DeleteResourceGroup")
    
    
    return
}

func NewDeleteResourceGroupResponse() (response *DeleteResourceGroupResponse) {
    response = &DeleteResourceGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteResourceGroup
// This API is used to delete asset groups in Asset Management.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteResourceGroup(request *DeleteResourceGroupRequest) (response *DeleteResourceGroupResponse, err error) {
    return c.DeleteResourceGroupWithContext(context.Background(), request)
}

// DeleteResourceGroup
// This API is used to delete asset groups in Asset Management.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteResourceGroupWithContext(ctx context.Context, request *DeleteResourceGroupRequest) (response *DeleteResourceGroupResponse, err error) {
    if request == nil {
        request = NewDeleteResourceGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "DeleteResourceGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteResourceGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteResourceGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSecurityGroupRuleRequest() (request *DeleteSecurityGroupRuleRequest) {
    request = &DeleteSecurityGroupRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DeleteSecurityGroupRule")
    
    
    return
}

func NewDeleteSecurityGroupRuleResponse() (response *DeleteSecurityGroupRuleResponse) {
    response = &DeleteSecurityGroupRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSecurityGroupRule
// This API is used to delete security group rules.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteSecurityGroupRule(request *DeleteSecurityGroupRuleRequest) (response *DeleteSecurityGroupRuleResponse, err error) {
    return c.DeleteSecurityGroupRuleWithContext(context.Background(), request)
}

// DeleteSecurityGroupRule
// This API is used to delete security group rules.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteSecurityGroupRuleWithContext(ctx context.Context, request *DeleteSecurityGroupRuleRequest) (response *DeleteSecurityGroupRuleResponse, err error) {
    if request == nil {
        request = NewDeleteSecurityGroupRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "DeleteSecurityGroupRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSecurityGroupRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSecurityGroupRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteVpcInstanceRequest() (request *DeleteVpcInstanceRequest) {
    request = &DeleteVpcInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DeleteVpcInstance")
    
    
    return
}

func NewDeleteVpcInstanceResponse() (response *DeleteVpcInstanceResponse) {
    response = &DeleteVpcInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteVpcInstance
// This API is used to delete firewall instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DeleteVpcInstance(request *DeleteVpcInstanceRequest) (response *DeleteVpcInstanceResponse, err error) {
    return c.DeleteVpcInstanceWithContext(context.Background(), request)
}

// DeleteVpcInstance
// This API is used to delete firewall instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DeleteVpcInstanceWithContext(ctx context.Context, request *DeleteVpcInstanceRequest) (response *DeleteVpcInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteVpcInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "DeleteVpcInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteVpcInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteVpcInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAcListsRequest() (request *DescribeAcListsRequest) {
    request = &DescribeAcListsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeAcLists")
    
    
    return
}

func NewDescribeAcListsResponse() (response *DescribeAcListsResponse) {
    response = &DescribeAcListsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAcLists
// This API is used to get the access control list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAcLists(request *DescribeAcListsRequest) (response *DescribeAcListsResponse, err error) {
    return c.DescribeAcListsWithContext(context.Background(), request)
}

// DescribeAcLists
// This API is used to get the access control list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAcListsWithContext(ctx context.Context, request *DescribeAcListsRequest) (response *DescribeAcListsResponse, err error) {
    if request == nil {
        request = NewDescribeAcListsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "DescribeAcLists")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAcLists require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAcListsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssociatedInstanceListRequest() (request *DescribeAssociatedInstanceListRequest) {
    request = &DescribeAssociatedInstanceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeAssociatedInstanceList")
    
    
    return
}

func NewDescribeAssociatedInstanceListResponse() (response *DescribeAssociatedInstanceListResponse) {
    response = &DescribeAssociatedInstanceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssociatedInstanceList
// This API is used to get the list of instances associated with a security group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAssociatedInstanceList(request *DescribeAssociatedInstanceListRequest) (response *DescribeAssociatedInstanceListResponse, err error) {
    return c.DescribeAssociatedInstanceListWithContext(context.Background(), request)
}

// DescribeAssociatedInstanceList
// This API is used to get the list of instances associated with a security group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAssociatedInstanceListWithContext(ctx context.Context, request *DescribeAssociatedInstanceListRequest) (response *DescribeAssociatedInstanceListResponse, err error) {
    if request == nil {
        request = NewDescribeAssociatedInstanceListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "DescribeAssociatedInstanceList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssociatedInstanceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssociatedInstanceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBlockByIpTimesListRequest() (request *DescribeBlockByIpTimesListRequest) {
    request = &DescribeBlockByIpTimesListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeBlockByIpTimesList")
    
    
    return
}

func NewDescribeBlockByIpTimesListResponse() (response *DescribeBlockByIpTimesListResponse) {
    response = &DescribeBlockByIpTimesListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBlockByIpTimesList
// This API is used to get blocked IP data.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBlockByIpTimesList(request *DescribeBlockByIpTimesListRequest) (response *DescribeBlockByIpTimesListResponse, err error) {
    return c.DescribeBlockByIpTimesListWithContext(context.Background(), request)
}

// DescribeBlockByIpTimesList
// This API is used to get blocked IP data.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBlockByIpTimesListWithContext(ctx context.Context, request *DescribeBlockByIpTimesListRequest) (response *DescribeBlockByIpTimesListResponse, err error) {
    if request == nil {
        request = NewDescribeBlockByIpTimesListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "DescribeBlockByIpTimesList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBlockByIpTimesList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBlockByIpTimesListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBlockIgnoreListRequest() (request *DescribeBlockIgnoreListRequest) {
    request = &DescribeBlockIgnoreListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeBlockIgnoreList")
    
    
    return
}

func NewDescribeBlockIgnoreListResponse() (response *DescribeBlockIgnoreListResponse) {
    response = &DescribeBlockIgnoreListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBlockIgnoreList
// This API is used to get allowlists or blocklists for intrusion prevention.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBlockIgnoreList(request *DescribeBlockIgnoreListRequest) (response *DescribeBlockIgnoreListResponse, err error) {
    return c.DescribeBlockIgnoreListWithContext(context.Background(), request)
}

// DescribeBlockIgnoreList
// This API is used to get allowlists or blocklists for intrusion prevention.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBlockIgnoreListWithContext(ctx context.Context, request *DescribeBlockIgnoreListRequest) (response *DescribeBlockIgnoreListResponse, err error) {
    if request == nil {
        request = NewDescribeBlockIgnoreListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "DescribeBlockIgnoreList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBlockIgnoreList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBlockIgnoreListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBlockStaticListRequest() (request *DescribeBlockStaticListRequest) {
    request = &DescribeBlockStaticListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeBlockStaticList")
    
    
    return
}

func NewDescribeBlockStaticListResponse() (response *DescribeBlockStaticListResponse) {
    response = &DescribeBlockStaticListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBlockStaticList
// This API is used to get the most frequent attacker.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBlockStaticList(request *DescribeBlockStaticListRequest) (response *DescribeBlockStaticListResponse, err error) {
    return c.DescribeBlockStaticListWithContext(context.Background(), request)
}

// DescribeBlockStaticList
// This API is used to get the most frequent attacker.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBlockStaticListWithContext(ctx context.Context, request *DescribeBlockStaticListRequest) (response *DescribeBlockStaticListResponse, err error) {
    if request == nil {
        request = NewDescribeBlockStaticListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "DescribeBlockStaticList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBlockStaticList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBlockStaticListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDefenseSwitchRequest() (request *DescribeDefenseSwitchRequest) {
    request = &DescribeDefenseSwitchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeDefenseSwitch")
    
    
    return
}

func NewDescribeDefenseSwitchResponse() (response *DescribeDefenseSwitchResponse) {
    response = &DescribeDefenseSwitchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDefenseSwitch
// This API is used to query the list of firewall toggles with Intrusion Defense enabled.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeDefenseSwitch(request *DescribeDefenseSwitchRequest) (response *DescribeDefenseSwitchResponse, err error) {
    return c.DescribeDefenseSwitchWithContext(context.Background(), request)
}

// DescribeDefenseSwitch
// This API is used to query the list of firewall toggles with Intrusion Defense enabled.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeDefenseSwitchWithContext(ctx context.Context, request *DescribeDefenseSwitchRequest) (response *DescribeDefenseSwitchResponse, err error) {
    if request == nil {
        request = NewDescribeDefenseSwitchRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "DescribeDefenseSwitch")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDefenseSwitch require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDefenseSwitchResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnterpriseSecurityGroupRuleRequest() (request *DescribeEnterpriseSecurityGroupRuleRequest) {
    request = &DescribeEnterpriseSecurityGroupRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeEnterpriseSecurityGroupRule")
    
    
    return
}

func NewDescribeEnterpriseSecurityGroupRuleResponse() (response *DescribeEnterpriseSecurityGroupRuleResponse) {
    response = &DescribeEnterpriseSecurityGroupRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEnterpriseSecurityGroupRule
// This API is used to get enterprise security group rules (new).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEnterpriseSecurityGroupRule(request *DescribeEnterpriseSecurityGroupRuleRequest) (response *DescribeEnterpriseSecurityGroupRuleResponse, err error) {
    return c.DescribeEnterpriseSecurityGroupRuleWithContext(context.Background(), request)
}

// DescribeEnterpriseSecurityGroupRule
// This API is used to get enterprise security group rules (new).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeEnterpriseSecurityGroupRuleWithContext(ctx context.Context, request *DescribeEnterpriseSecurityGroupRuleRequest) (response *DescribeEnterpriseSecurityGroupRuleResponse, err error) {
    if request == nil {
        request = NewDescribeEnterpriseSecurityGroupRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "DescribeEnterpriseSecurityGroupRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEnterpriseSecurityGroupRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEnterpriseSecurityGroupRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGuideScanInfoRequest() (request *DescribeGuideScanInfoRequest) {
    request = &DescribeGuideScanInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeGuideScanInfo")
    
    
    return
}

func NewDescribeGuideScanInfoResponse() (response *DescribeGuideScanInfoResponse) {
    response = &DescribeGuideScanInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGuideScanInfo
// This API is used to get the scan interface information in Get Started.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeGuideScanInfo(request *DescribeGuideScanInfoRequest) (response *DescribeGuideScanInfoResponse, err error) {
    return c.DescribeGuideScanInfoWithContext(context.Background(), request)
}

// DescribeGuideScanInfo
// This API is used to get the scan interface information in Get Started.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeGuideScanInfoWithContext(ctx context.Context, request *DescribeGuideScanInfoRequest) (response *DescribeGuideScanInfoResponse, err error) {
    if request == nil {
        request = NewDescribeGuideScanInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "DescribeGuideScanInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGuideScanInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGuideScanInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIPStatusListRequest() (request *DescribeIPStatusListRequest) {
    request = &DescribeIPStatusListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeIPStatusList")
    
    
    return
}

func NewDescribeIPStatusListResponse() (response *DescribeIPStatusListResponse) {
    response = &DescribeIPStatusListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIPStatusList
// This API is used to get the IP protection status.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeIPStatusList(request *DescribeIPStatusListRequest) (response *DescribeIPStatusListResponse, err error) {
    return c.DescribeIPStatusListWithContext(context.Background(), request)
}

// DescribeIPStatusList
// This API is used to get the IP protection status.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeIPStatusListWithContext(ctx context.Context, request *DescribeIPStatusListRequest) (response *DescribeIPStatusListResponse, err error) {
    if request == nil {
        request = NewDescribeIPStatusListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "DescribeIPStatusList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIPStatusList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIPStatusListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNatAcRuleRequest() (request *DescribeNatAcRuleRequest) {
    request = &DescribeNatAcRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeNatAcRule")
    
    
    return
}

func NewDescribeNatAcRuleResponse() (response *DescribeNatAcRuleResponse) {
    response = &DescribeNatAcRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNatAcRule
// This API is used to get the NAT access control list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeNatAcRule(request *DescribeNatAcRuleRequest) (response *DescribeNatAcRuleResponse, err error) {
    return c.DescribeNatAcRuleWithContext(context.Background(), request)
}

// DescribeNatAcRule
// This API is used to get the NAT access control list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeNatAcRuleWithContext(ctx context.Context, request *DescribeNatAcRuleRequest) (response *DescribeNatAcRuleResponse, err error) {
    if request == nil {
        request = NewDescribeNatAcRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "DescribeNatAcRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNatAcRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNatAcRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNatFwInfoCountRequest() (request *DescribeNatFwInfoCountRequest) {
    request = &DescribeNatFwInfoCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeNatFwInfoCount")
    
    
    return
}

func NewDescribeNatFwInfoCountResponse() (response *DescribeNatFwInfoCountResponse) {
    response = &DescribeNatFwInfoCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNatFwInfoCount
// This API is used to get the number of a user's subnets connected to NAT firewall and the number of NAT firewall instances.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNatFwInfoCount(request *DescribeNatFwInfoCountRequest) (response *DescribeNatFwInfoCountResponse, err error) {
    return c.DescribeNatFwInfoCountWithContext(context.Background(), request)
}

// DescribeNatFwInfoCount
// This API is used to get the number of a user's subnets connected to NAT firewall and the number of NAT firewall instances.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNatFwInfoCountWithContext(ctx context.Context, request *DescribeNatFwInfoCountRequest) (response *DescribeNatFwInfoCountResponse, err error) {
    if request == nil {
        request = NewDescribeNatFwInfoCountRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "DescribeNatFwInfoCount")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNatFwInfoCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNatFwInfoCountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNatFwInstanceRequest() (request *DescribeNatFwInstanceRequest) {
    request = &DescribeNatFwInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeNatFwInstance")
    
    
    return
}

func NewDescribeNatFwInstanceResponse() (response *DescribeNatFwInstanceResponse) {
    response = &DescribeNatFwInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNatFwInstance
// This API is used to get all NAT instances of a tenant.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNatFwInstance(request *DescribeNatFwInstanceRequest) (response *DescribeNatFwInstanceResponse, err error) {
    return c.DescribeNatFwInstanceWithContext(context.Background(), request)
}

// DescribeNatFwInstance
// This API is used to get all NAT instances of a tenant.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNatFwInstanceWithContext(ctx context.Context, request *DescribeNatFwInstanceRequest) (response *DescribeNatFwInstanceResponse, err error) {
    if request == nil {
        request = NewDescribeNatFwInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "DescribeNatFwInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNatFwInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNatFwInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNatFwInstanceWithRegionRequest() (request *DescribeNatFwInstanceWithRegionRequest) {
    request = &DescribeNatFwInstanceWithRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeNatFwInstanceWithRegion")
    
    
    return
}

func NewDescribeNatFwInstanceWithRegionResponse() (response *DescribeNatFwInstanceWithRegionResponse) {
    response = &DescribeNatFwInstanceWithRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNatFwInstanceWithRegion
// This API is used to get the NAT instance with the region that is newly maintained by the tenant.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNatFwInstanceWithRegion(request *DescribeNatFwInstanceWithRegionRequest) (response *DescribeNatFwInstanceWithRegionResponse, err error) {
    return c.DescribeNatFwInstanceWithRegionWithContext(context.Background(), request)
}

// DescribeNatFwInstanceWithRegion
// This API is used to get the NAT instance with the region that is newly maintained by the tenant.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNatFwInstanceWithRegionWithContext(ctx context.Context, request *DescribeNatFwInstanceWithRegionRequest) (response *DescribeNatFwInstanceWithRegionResponse, err error) {
    if request == nil {
        request = NewDescribeNatFwInstanceWithRegionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "DescribeNatFwInstanceWithRegion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNatFwInstanceWithRegion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNatFwInstanceWithRegionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNatFwInstancesInfoRequest() (request *DescribeNatFwInstancesInfoRequest) {
    request = &DescribeNatFwInstancesInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeNatFwInstancesInfo")
    
    
    return
}

func NewDescribeNatFwInstancesInfoResponse() (response *DescribeNatFwInstancesInfoResponse) {
    response = &DescribeNatFwInstancesInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNatFwInstancesInfo
// This API is used to get all NAT instances and instance card information of a tenant.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNatFwInstancesInfo(request *DescribeNatFwInstancesInfoRequest) (response *DescribeNatFwInstancesInfoResponse, err error) {
    return c.DescribeNatFwInstancesInfoWithContext(context.Background(), request)
}

// DescribeNatFwInstancesInfo
// This API is used to get all NAT instances and instance card information of a tenant.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNatFwInstancesInfoWithContext(ctx context.Context, request *DescribeNatFwInstancesInfoRequest) (response *DescribeNatFwInstancesInfoResponse, err error) {
    if request == nil {
        request = NewDescribeNatFwInstancesInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "DescribeNatFwInstancesInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNatFwInstancesInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNatFwInstancesInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNatFwVpcDnsLstRequest() (request *DescribeNatFwVpcDnsLstRequest) {
    request = &DescribeNatFwVpcDnsLstRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeNatFwVpcDnsLst")
    
    
    return
}

func NewDescribeNatFwVpcDnsLstResponse() (response *DescribeNatFwVpcDnsLstResponse) {
    response = &DescribeNatFwVpcDnsLstResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNatFwVpcDnsLst
// This API is used to get the VPC DNS status of NAT firewall instances.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNatFwVpcDnsLst(request *DescribeNatFwVpcDnsLstRequest) (response *DescribeNatFwVpcDnsLstResponse, err error) {
    return c.DescribeNatFwVpcDnsLstWithContext(context.Background(), request)
}

// DescribeNatFwVpcDnsLst
// This API is used to get the VPC DNS status of NAT firewall instances.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNatFwVpcDnsLstWithContext(ctx context.Context, request *DescribeNatFwVpcDnsLstRequest) (response *DescribeNatFwVpcDnsLstResponse, err error) {
    if request == nil {
        request = NewDescribeNatFwVpcDnsLstRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "DescribeNatFwVpcDnsLst")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNatFwVpcDnsLst require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNatFwVpcDnsLstResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceGroupNewRequest() (request *DescribeResourceGroupNewRequest) {
    request = &DescribeResourceGroupNewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeResourceGroupNew")
    
    
    return
}

func NewDescribeResourceGroupNewResponse() (response *DescribeResourceGroupNewResponse) {
    response = &DescribeResourceGroupNewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeResourceGroupNew
// This API is used to get the asset tree information in Asset Management.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeResourceGroupNew(request *DescribeResourceGroupNewRequest) (response *DescribeResourceGroupNewResponse, err error) {
    return c.DescribeResourceGroupNewWithContext(context.Background(), request)
}

// DescribeResourceGroupNew
// This API is used to get the asset tree information in Asset Management.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeResourceGroupNewWithContext(ctx context.Context, request *DescribeResourceGroupNewRequest) (response *DescribeResourceGroupNewResponse, err error) {
    if request == nil {
        request = NewDescribeResourceGroupNewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "DescribeResourceGroupNew")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourceGroupNew require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourceGroupNewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRuleOverviewRequest() (request *DescribeRuleOverviewRequest) {
    request = &DescribeRuleOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeRuleOverview")
    
    
    return
}

func NewDescribeRuleOverviewResponse() (response *DescribeRuleOverviewResponse) {
    response = &DescribeRuleOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRuleOverview
// This API is used to get the rule list overview.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRuleOverview(request *DescribeRuleOverviewRequest) (response *DescribeRuleOverviewResponse, err error) {
    return c.DescribeRuleOverviewWithContext(context.Background(), request)
}

// DescribeRuleOverview
// This API is used to get the rule list overview.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRuleOverviewWithContext(ctx context.Context, request *DescribeRuleOverviewRequest) (response *DescribeRuleOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeRuleOverviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "DescribeRuleOverview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRuleOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRuleOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityGroupListRequest() (request *DescribeSecurityGroupListRequest) {
    request = &DescribeSecurityGroupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeSecurityGroupList")
    
    
    return
}

func NewDescribeSecurityGroupListResponse() (response *DescribeSecurityGroupListResponse) {
    response = &DescribeSecurityGroupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecurityGroupList
// This API is used to get the security group rule list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeSecurityGroupList(request *DescribeSecurityGroupListRequest) (response *DescribeSecurityGroupListResponse, err error) {
    return c.DescribeSecurityGroupListWithContext(context.Background(), request)
}

// DescribeSecurityGroupList
// This API is used to get the security group rule list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeSecurityGroupListWithContext(ctx context.Context, request *DescribeSecurityGroupListRequest) (response *DescribeSecurityGroupListResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityGroupListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "DescribeSecurityGroupList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityGroupList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityGroupListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSourceAssetRequest() (request *DescribeSourceAssetRequest) {
    request = &DescribeSourceAssetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeSourceAsset")
    
    
    return
}

func NewDescribeSourceAssetResponse() (response *DescribeSourceAssetResponse) {
    response = &DescribeSourceAssetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSourceAsset
// This API is used to get all asset information of an asset group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSourceAsset(request *DescribeSourceAssetRequest) (response *DescribeSourceAssetResponse, err error) {
    return c.DescribeSourceAssetWithContext(context.Background(), request)
}

// DescribeSourceAsset
// This API is used to get all asset information of an asset group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSourceAssetWithContext(ctx context.Context, request *DescribeSourceAssetRequest) (response *DescribeSourceAssetResponse, err error) {
    if request == nil {
        request = NewDescribeSourceAssetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "DescribeSourceAsset")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSourceAsset require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSourceAssetResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSwitchListsRequest() (request *DescribeSwitchListsRequest) {
    request = &DescribeSwitchListsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeSwitchLists")
    
    
    return
}

func NewDescribeSwitchListsResponse() (response *DescribeSwitchListsResponse) {
    response = &DescribeSwitchListsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSwitchLists
// This API is used to get the firewall status list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSwitchLists(request *DescribeSwitchListsRequest) (response *DescribeSwitchListsResponse, err error) {
    return c.DescribeSwitchListsWithContext(context.Background(), request)
}

// DescribeSwitchLists
// This API is used to get the firewall status list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSwitchListsWithContext(ctx context.Context, request *DescribeSwitchListsRequest) (response *DescribeSwitchListsResponse, err error) {
    if request == nil {
        request = NewDescribeSwitchListsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "DescribeSwitchLists")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSwitchLists require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSwitchListsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTLogInfoRequest() (request *DescribeTLogInfoRequest) {
    request = &DescribeTLogInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeTLogInfo")
    
    
    return
}

func NewDescribeTLogInfoResponse() (response *DescribeTLogInfoResponse) {
    response = &DescribeTLogInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTLogInfo
// This API is used to get the current alert monitoring data.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTLogInfo(request *DescribeTLogInfoRequest) (response *DescribeTLogInfoResponse, err error) {
    return c.DescribeTLogInfoWithContext(context.Background(), request)
}

// DescribeTLogInfo
// This API is used to get the current alert monitoring data.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTLogInfoWithContext(ctx context.Context, request *DescribeTLogInfoRequest) (response *DescribeTLogInfoResponse, err error) {
    if request == nil {
        request = NewDescribeTLogInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "DescribeTLogInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTLogInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTLogInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTLogIpListRequest() (request *DescribeTLogIpListRequest) {
    request = &DescribeTLogIpListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeTLogIpList")
    
    
    return
}

func NewDescribeTLogIpListResponse() (response *DescribeTLogIpListResponse) {
    response = &DescribeTLogIpListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTLogIpList
// This API is used to get the most frequent attacker IP.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTLogIpList(request *DescribeTLogIpListRequest) (response *DescribeTLogIpListResponse, err error) {
    return c.DescribeTLogIpListWithContext(context.Background(), request)
}

// DescribeTLogIpList
// This API is used to get the most frequent attacker IP.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTLogIpListWithContext(ctx context.Context, request *DescribeTLogIpListRequest) (response *DescribeTLogIpListResponse, err error) {
    if request == nil {
        request = NewDescribeTLogIpListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "DescribeTLogIpList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTLogIpList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTLogIpListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTableStatusRequest() (request *DescribeTableStatusRequest) {
    request = &DescribeTableStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeTableStatus")
    
    
    return
}

func NewDescribeTableStatusResponse() (response *DescribeTableStatusResponse) {
    response = &DescribeTableStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTableStatus
// This API is used to get the rule list status.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTableStatus(request *DescribeTableStatusRequest) (response *DescribeTableStatusResponse, err error) {
    return c.DescribeTableStatusWithContext(context.Background(), request)
}

// DescribeTableStatus
// This API is used to get the rule list status.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTableStatusWithContext(ctx context.Context, request *DescribeTableStatusRequest) (response *DescribeTableStatusResponse, err error) {
    if request == nil {
        request = NewDescribeTableStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "DescribeTableStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTableStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTableStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUnHandleEventTabListRequest() (request *DescribeUnHandleEventTabListRequest) {
    request = &DescribeUnHandleEventTabListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeUnHandleEventTabList")
    
    
    return
}

func NewDescribeUnHandleEventTabListResponse() (response *DescribeUnHandleEventTabListResponse) {
    response = &DescribeUnHandleEventTabListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUnHandleEventTabList
// This API is used to get unhandled security events.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUnHandleEventTabList(request *DescribeUnHandleEventTabListRequest) (response *DescribeUnHandleEventTabListResponse, err error) {
    return c.DescribeUnHandleEventTabListWithContext(context.Background(), request)
}

// DescribeUnHandleEventTabList
// This API is used to get unhandled security events.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUnHandleEventTabListWithContext(ctx context.Context, request *DescribeUnHandleEventTabListRequest) (response *DescribeUnHandleEventTabListResponse, err error) {
    if request == nil {
        request = NewDescribeUnHandleEventTabListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "DescribeUnHandleEventTabList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUnHandleEventTabList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUnHandleEventTabListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVpcAcRuleRequest() (request *DescribeVpcAcRuleRequest) {
    request = &DescribeVpcAcRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "DescribeVpcAcRule")
    
    
    return
}

func NewDescribeVpcAcRuleResponse() (response *DescribeVpcAcRuleResponse) {
    response = &DescribeVpcAcRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVpcAcRule
// Query Inter-VPC rules
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeVpcAcRule(request *DescribeVpcAcRuleRequest) (response *DescribeVpcAcRuleResponse, err error) {
    return c.DescribeVpcAcRuleWithContext(context.Background(), request)
}

// DescribeVpcAcRule
// Query Inter-VPC rules
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeVpcAcRuleWithContext(ctx context.Context, request *DescribeVpcAcRuleRequest) (response *DescribeVpcAcRuleResponse, err error) {
    if request == nil {
        request = NewDescribeVpcAcRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "DescribeVpcAcRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVpcAcRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVpcAcRuleResponse()
    err = c.Send(request, response)
    return
}

func NewExpandCfwVerticalRequest() (request *ExpandCfwVerticalRequest) {
    request = &ExpandCfwVerticalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ExpandCfwVertical")
    
    
    return
}

func NewExpandCfwVerticalResponse() (response *ExpandCfwVerticalResponse) {
    response = &ExpandCfwVerticalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExpandCfwVertical
// This API is used to increase the firewall bandwidth.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ExpandCfwVertical(request *ExpandCfwVerticalRequest) (response *ExpandCfwVerticalResponse, err error) {
    return c.ExpandCfwVerticalWithContext(context.Background(), request)
}

// ExpandCfwVertical
// This API is used to increase the firewall bandwidth.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ExpandCfwVerticalWithContext(ctx context.Context, request *ExpandCfwVerticalRequest) (response *ExpandCfwVerticalResponse, err error) {
    if request == nil {
        request = NewExpandCfwVerticalRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "ExpandCfwVertical")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExpandCfwVertical require credential")
    }

    request.SetContext(ctx)
    
    response = NewExpandCfwVerticalResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAcRuleRequest() (request *ModifyAcRuleRequest) {
    request = &ModifyAcRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyAcRule")
    
    
    return
}

func NewModifyAcRuleResponse() (response *ModifyAcRuleResponse) {
    response = &ModifyAcRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAcRule
// This API is used to modify rules.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyAcRule(request *ModifyAcRuleRequest) (response *ModifyAcRuleResponse, err error) {
    return c.ModifyAcRuleWithContext(context.Background(), request)
}

// ModifyAcRule
// This API is used to modify rules.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyAcRuleWithContext(ctx context.Context, request *ModifyAcRuleRequest) (response *ModifyAcRuleResponse, err error) {
    if request == nil {
        request = NewModifyAcRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "ModifyAcRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAcRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAcRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAllPublicIPSwitchStatusRequest() (request *ModifyAllPublicIPSwitchStatusRequest) {
    request = &ModifyAllPublicIPSwitchStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyAllPublicIPSwitchStatus")
    
    
    return
}

func NewModifyAllPublicIPSwitchStatusResponse() (response *ModifyAllPublicIPSwitchStatusResponse) {
    response = &ModifyAllPublicIPSwitchStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAllPublicIPSwitchStatus
// This API is used to enable or disable one or multiple edge firewalls.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) ModifyAllPublicIPSwitchStatus(request *ModifyAllPublicIPSwitchStatusRequest) (response *ModifyAllPublicIPSwitchStatusResponse, err error) {
    return c.ModifyAllPublicIPSwitchStatusWithContext(context.Background(), request)
}

// ModifyAllPublicIPSwitchStatus
// This API is used to enable or disable one or multiple edge firewalls.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) ModifyAllPublicIPSwitchStatusWithContext(ctx context.Context, request *ModifyAllPublicIPSwitchStatusRequest) (response *ModifyAllPublicIPSwitchStatusResponse, err error) {
    if request == nil {
        request = NewModifyAllPublicIPSwitchStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "ModifyAllPublicIPSwitchStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAllPublicIPSwitchStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAllPublicIPSwitchStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAllRuleStatusRequest() (request *ModifyAllRuleStatusRequest) {
    request = &ModifyAllRuleStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyAllRuleStatus")
    
    
    return
}

func NewModifyAllRuleStatusResponse() (response *ModifyAllRuleStatusResponse) {
    response = &ModifyAllRuleStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAllRuleStatus
// This API is used to enable or disable all rules.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyAllRuleStatus(request *ModifyAllRuleStatusRequest) (response *ModifyAllRuleStatusResponse, err error) {
    return c.ModifyAllRuleStatusWithContext(context.Background(), request)
}

// ModifyAllRuleStatus
// This API is used to enable or disable all rules.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyAllRuleStatusWithContext(ctx context.Context, request *ModifyAllRuleStatusRequest) (response *ModifyAllRuleStatusResponse, err error) {
    if request == nil {
        request = NewModifyAllRuleStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "ModifyAllRuleStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAllRuleStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAllRuleStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAllVPCSwitchStatusRequest() (request *ModifyAllVPCSwitchStatusRequest) {
    request = &ModifyAllVPCSwitchStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyAllVPCSwitchStatus")
    
    
    return
}

func NewModifyAllVPCSwitchStatusResponse() (response *ModifyAllVPCSwitchStatusResponse) {
    response = &ModifyAllVPCSwitchStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAllVPCSwitchStatus
// This API is used to enable or disable one or multiple VPC firewalls.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) ModifyAllVPCSwitchStatus(request *ModifyAllVPCSwitchStatusRequest) (response *ModifyAllVPCSwitchStatusResponse, err error) {
    return c.ModifyAllVPCSwitchStatusWithContext(context.Background(), request)
}

// ModifyAllVPCSwitchStatus
// This API is used to enable or disable one or multiple VPC firewalls.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) ModifyAllVPCSwitchStatusWithContext(ctx context.Context, request *ModifyAllVPCSwitchStatusRequest) (response *ModifyAllVPCSwitchStatusResponse, err error) {
    if request == nil {
        request = NewModifyAllVPCSwitchStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "ModifyAllVPCSwitchStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAllVPCSwitchStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAllVPCSwitchStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAssetScanRequest() (request *ModifyAssetScanRequest) {
    request = &ModifyAssetScanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyAssetScan")
    
    
    return
}

func NewModifyAssetScanResponse() (response *ModifyAssetScanResponse) {
    response = &ModifyAssetScanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAssetScan
// This API is used to modify asset scan settings.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) ModifyAssetScan(request *ModifyAssetScanRequest) (response *ModifyAssetScanResponse, err error) {
    return c.ModifyAssetScanWithContext(context.Background(), request)
}

// ModifyAssetScan
// This API is used to modify asset scan settings.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) ModifyAssetScanWithContext(ctx context.Context, request *ModifyAssetScanRequest) (response *ModifyAssetScanResponse, err error) {
    if request == nil {
        request = NewModifyAssetScanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "ModifyAssetScan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAssetScan require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAssetScanResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBlockIgnoreListRequest() (request *ModifyBlockIgnoreListRequest) {
    request = &ModifyBlockIgnoreListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyBlockIgnoreList")
    
    
    return
}

func NewModifyBlockIgnoreListResponse() (response *ModifyBlockIgnoreListResponse) {
    response = &ModifyBlockIgnoreListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyBlockIgnoreList
// This API is used to manage blocked/allowed IPs and domains.
//
// Add IPs/domains to the blocked/allowed list
//
// Remove IPs/domains from the blocked/allowed list
//
// Modify events related with the IPs/domains in the blocked/allowed list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyBlockIgnoreList(request *ModifyBlockIgnoreListRequest) (response *ModifyBlockIgnoreListResponse, err error) {
    return c.ModifyBlockIgnoreListWithContext(context.Background(), request)
}

// ModifyBlockIgnoreList
// This API is used to manage blocked/allowed IPs and domains.
//
// Add IPs/domains to the blocked/allowed list
//
// Remove IPs/domains from the blocked/allowed list
//
// Modify events related with the IPs/domains in the blocked/allowed list
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyBlockIgnoreListWithContext(ctx context.Context, request *ModifyBlockIgnoreListRequest) (response *ModifyBlockIgnoreListResponse, err error) {
    if request == nil {
        request = NewModifyBlockIgnoreListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "ModifyBlockIgnoreList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBlockIgnoreList require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBlockIgnoreListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBlockTopRequest() (request *ModifyBlockTopRequest) {
    request = &ModifyBlockTopRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyBlockTop")
    
    
    return
}

func NewModifyBlockTopResponse() (response *ModifyBlockTopResponse) {
    response = &ModifyBlockTopResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyBlockTop
// This API is used to pin or unpin a blocking log.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyBlockTop(request *ModifyBlockTopRequest) (response *ModifyBlockTopResponse, err error) {
    return c.ModifyBlockTopWithContext(context.Background(), request)
}

// ModifyBlockTop
// This API is used to pin or unpin a blocking log.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyBlockTopWithContext(ctx context.Context, request *ModifyBlockTopRequest) (response *ModifyBlockTopResponse, err error) {
    if request == nil {
        request = NewModifyBlockTopRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "ModifyBlockTop")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBlockTop require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBlockTopResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEnterpriseSecurityDispatchStatusRequest() (request *ModifyEnterpriseSecurityDispatchStatusRequest) {
    request = &ModifyEnterpriseSecurityDispatchStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyEnterpriseSecurityDispatchStatus")
    
    
    return
}

func NewModifyEnterpriseSecurityDispatchStatusResponse() (response *ModifyEnterpriseSecurityDispatchStatusResponse) {
    response = &ModifyEnterpriseSecurityDispatchStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyEnterpriseSecurityDispatchStatus
// This API is used to modify the publishing status of an enterprise security group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyEnterpriseSecurityDispatchStatus(request *ModifyEnterpriseSecurityDispatchStatusRequest) (response *ModifyEnterpriseSecurityDispatchStatusResponse, err error) {
    return c.ModifyEnterpriseSecurityDispatchStatusWithContext(context.Background(), request)
}

// ModifyEnterpriseSecurityDispatchStatus
// This API is used to modify the publishing status of an enterprise security group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyEnterpriseSecurityDispatchStatusWithContext(ctx context.Context, request *ModifyEnterpriseSecurityDispatchStatusRequest) (response *ModifyEnterpriseSecurityDispatchStatusResponse, err error) {
    if request == nil {
        request = NewModifyEnterpriseSecurityDispatchStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "ModifyEnterpriseSecurityDispatchStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEnterpriseSecurityDispatchStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEnterpriseSecurityDispatchStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEnterpriseSecurityGroupRuleRequest() (request *ModifyEnterpriseSecurityGroupRuleRequest) {
    request = &ModifyEnterpriseSecurityGroupRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyEnterpriseSecurityGroupRule")
    
    
    return
}

func NewModifyEnterpriseSecurityGroupRuleResponse() (response *ModifyEnterpriseSecurityGroupRuleResponse) {
    response = &ModifyEnterpriseSecurityGroupRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyEnterpriseSecurityGroupRule
// This API is used to modify a new enterprise security group rule.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyEnterpriseSecurityGroupRule(request *ModifyEnterpriseSecurityGroupRuleRequest) (response *ModifyEnterpriseSecurityGroupRuleResponse, err error) {
    return c.ModifyEnterpriseSecurityGroupRuleWithContext(context.Background(), request)
}

// ModifyEnterpriseSecurityGroupRule
// This API is used to modify a new enterprise security group rule.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyEnterpriseSecurityGroupRuleWithContext(ctx context.Context, request *ModifyEnterpriseSecurityGroupRuleRequest) (response *ModifyEnterpriseSecurityGroupRuleResponse, err error) {
    if request == nil {
        request = NewModifyEnterpriseSecurityGroupRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "ModifyEnterpriseSecurityGroupRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEnterpriseSecurityGroupRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEnterpriseSecurityGroupRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNatAcRuleRequest() (request *ModifyNatAcRuleRequest) {
    request = &ModifyNatAcRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyNatAcRule")
    
    
    return
}

func NewModifyNatAcRuleResponse() (response *ModifyNatAcRuleResponse) {
    response = &ModifyNatAcRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyNatAcRule
// This API is used to modify NAT access control rules.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyNatAcRule(request *ModifyNatAcRuleRequest) (response *ModifyNatAcRuleResponse, err error) {
    return c.ModifyNatAcRuleWithContext(context.Background(), request)
}

// ModifyNatAcRule
// This API is used to modify NAT access control rules.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyNatAcRuleWithContext(ctx context.Context, request *ModifyNatAcRuleRequest) (response *ModifyNatAcRuleResponse, err error) {
    if request == nil {
        request = NewModifyNatAcRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "ModifyNatAcRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNatAcRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNatAcRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNatFwReSelectRequest() (request *ModifyNatFwReSelectRequest) {
    request = &ModifyNatFwReSelectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyNatFwReSelect")
    
    
    return
}

func NewModifyNatFwReSelectResponse() (response *ModifyNatFwReSelectResponse) {
    response = &ModifyNatFwReSelectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyNatFwReSelect
// This API is used to get the VPC or NAT list for changing associated firewall instances.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyNatFwReSelect(request *ModifyNatFwReSelectRequest) (response *ModifyNatFwReSelectResponse, err error) {
    return c.ModifyNatFwReSelectWithContext(context.Background(), request)
}

// ModifyNatFwReSelect
// This API is used to get the VPC or NAT list for changing associated firewall instances.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyNatFwReSelectWithContext(ctx context.Context, request *ModifyNatFwReSelectRequest) (response *ModifyNatFwReSelectResponse, err error) {
    if request == nil {
        request = NewModifyNatFwReSelectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "ModifyNatFwReSelect")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNatFwReSelect require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNatFwReSelectResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNatFwSwitchRequest() (request *ModifyNatFwSwitchRequest) {
    request = &ModifyNatFwSwitchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyNatFwSwitch")
    
    
    return
}

func NewModifyNatFwSwitchResponse() (response *ModifyNatFwSwitchResponse) {
    response = &ModifyNatFwSwitchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyNatFwSwitch
// This API is used to enable or disable a NAT firewall.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyNatFwSwitch(request *ModifyNatFwSwitchRequest) (response *ModifyNatFwSwitchResponse, err error) {
    return c.ModifyNatFwSwitchWithContext(context.Background(), request)
}

// ModifyNatFwSwitch
// This API is used to enable or disable a NAT firewall.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyNatFwSwitchWithContext(ctx context.Context, request *ModifyNatFwSwitchRequest) (response *ModifyNatFwSwitchResponse, err error) {
    if request == nil {
        request = NewModifyNatFwSwitchRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "ModifyNatFwSwitch")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNatFwSwitch require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNatFwSwitchResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNatFwVpcDnsSwitchRequest() (request *ModifyNatFwVpcDnsSwitchRequest) {
    request = &ModifyNatFwVpcDnsSwitchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyNatFwVpcDnsSwitch")
    
    
    return
}

func NewModifyNatFwVpcDnsSwitchResponse() (response *ModifyNatFwVpcDnsSwitchResponse) {
    response = &ModifyNatFwVpcDnsSwitchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyNatFwVpcDnsSwitch
// This API is used to modify the VPC DNS status of NAT firewall instances.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyNatFwVpcDnsSwitch(request *ModifyNatFwVpcDnsSwitchRequest) (response *ModifyNatFwVpcDnsSwitchResponse, err error) {
    return c.ModifyNatFwVpcDnsSwitchWithContext(context.Background(), request)
}

// ModifyNatFwVpcDnsSwitch
// This API is used to modify the VPC DNS status of NAT firewall instances.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyNatFwVpcDnsSwitchWithContext(ctx context.Context, request *ModifyNatFwVpcDnsSwitchRequest) (response *ModifyNatFwVpcDnsSwitchResponse, err error) {
    if request == nil {
        request = NewModifyNatFwVpcDnsSwitchRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "ModifyNatFwVpcDnsSwitch")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNatFwVpcDnsSwitch require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNatFwVpcDnsSwitchResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNatSequenceRulesRequest() (request *ModifyNatSequenceRulesRequest) {
    request = &ModifyNatSequenceRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyNatSequenceRules")
    
    
    return
}

func NewModifyNatSequenceRulesResponse() (response *ModifyNatSequenceRulesResponse) {
    response = &ModifyNatSequenceRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyNatSequenceRules
// This API is used to change the sequence number of NAT firewall rules.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyNatSequenceRules(request *ModifyNatSequenceRulesRequest) (response *ModifyNatSequenceRulesResponse, err error) {
    return c.ModifyNatSequenceRulesWithContext(context.Background(), request)
}

// ModifyNatSequenceRules
// This API is used to change the sequence number of NAT firewall rules.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyNatSequenceRulesWithContext(ctx context.Context, request *ModifyNatSequenceRulesRequest) (response *ModifyNatSequenceRulesResponse, err error) {
    if request == nil {
        request = NewModifyNatSequenceRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "ModifyNatSequenceRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNatSequenceRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNatSequenceRulesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPublicIPSwitchStatusRequest() (request *ModifyPublicIPSwitchStatusRequest) {
    request = &ModifyPublicIPSwitchStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyPublicIPSwitchStatus")
    
    
    return
}

func NewModifyPublicIPSwitchStatusResponse() (response *ModifyPublicIPSwitchStatusResponse) {
    response = &ModifyPublicIPSwitchStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPublicIPSwitchStatus
// This API is used to enable or disable an edge firewall.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyPublicIPSwitchStatus(request *ModifyPublicIPSwitchStatusRequest) (response *ModifyPublicIPSwitchStatusResponse, err error) {
    return c.ModifyPublicIPSwitchStatusWithContext(context.Background(), request)
}

// ModifyPublicIPSwitchStatus
// This API is used to enable or disable an edge firewall.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyPublicIPSwitchStatusWithContext(ctx context.Context, request *ModifyPublicIPSwitchStatusRequest) (response *ModifyPublicIPSwitchStatusResponse, err error) {
    if request == nil {
        request = NewModifyPublicIPSwitchStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "ModifyPublicIPSwitchStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPublicIPSwitchStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPublicIPSwitchStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyResourceGroupRequest() (request *ModifyResourceGroupRequest) {
    request = &ModifyResourceGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyResourceGroup")
    
    
    return
}

func NewModifyResourceGroupResponse() (response *ModifyResourceGroupResponse) {
    response = &ModifyResourceGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyResourceGroup
// This API is used to modify the asset group information in Asset Management.
//
// 
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyResourceGroup(request *ModifyResourceGroupRequest) (response *ModifyResourceGroupResponse, err error) {
    return c.ModifyResourceGroupWithContext(context.Background(), request)
}

// ModifyResourceGroup
// This API is used to modify the asset group information in Asset Management.
//
// 
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyResourceGroupWithContext(ctx context.Context, request *ModifyResourceGroupRequest) (response *ModifyResourceGroupResponse, err error) {
    if request == nil {
        request = NewModifyResourceGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "ModifyResourceGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyResourceGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyResourceGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRunSyncAssetRequest() (request *ModifyRunSyncAssetRequest) {
    request = &ModifyRunSyncAssetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyRunSyncAsset")
    
    
    return
}

func NewModifyRunSyncAssetResponse() (response *ModifyRunSyncAssetResponse) {
    response = &ModifyRunSyncAssetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRunSyncAsset
// This API is used to sync assets - Internet & VPC (new).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyRunSyncAsset(request *ModifyRunSyncAssetRequest) (response *ModifyRunSyncAssetResponse, err error) {
    return c.ModifyRunSyncAssetWithContext(context.Background(), request)
}

// ModifyRunSyncAsset
// This API is used to sync assets - Internet & VPC (new).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyRunSyncAssetWithContext(ctx context.Context, request *ModifyRunSyncAssetRequest) (response *ModifyRunSyncAssetResponse, err error) {
    if request == nil {
        request = NewModifyRunSyncAssetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "ModifyRunSyncAsset")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRunSyncAsset require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRunSyncAssetResponse()
    err = c.Send(request, response)
    return
}

func NewModifySecurityGroupItemRuleStatusRequest() (request *ModifySecurityGroupItemRuleStatusRequest) {
    request = &ModifySecurityGroupItemRuleStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifySecurityGroupItemRuleStatus")
    
    
    return
}

func NewModifySecurityGroupItemRuleStatusResponse() (response *ModifySecurityGroupItemRuleStatusResponse) {
    response = &ModifySecurityGroupItemRuleStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySecurityGroupItemRuleStatus
// This API is used to enable or disable an enterprise security group rule.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifySecurityGroupItemRuleStatus(request *ModifySecurityGroupItemRuleStatusRequest) (response *ModifySecurityGroupItemRuleStatusResponse, err error) {
    return c.ModifySecurityGroupItemRuleStatusWithContext(context.Background(), request)
}

// ModifySecurityGroupItemRuleStatus
// This API is used to enable or disable an enterprise security group rule.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifySecurityGroupItemRuleStatusWithContext(ctx context.Context, request *ModifySecurityGroupItemRuleStatusRequest) (response *ModifySecurityGroupItemRuleStatusResponse, err error) {
    if request == nil {
        request = NewModifySecurityGroupItemRuleStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "ModifySecurityGroupItemRuleStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySecurityGroupItemRuleStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySecurityGroupItemRuleStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifySecurityGroupSequenceRulesRequest() (request *ModifySecurityGroupSequenceRulesRequest) {
    request = &ModifySecurityGroupSequenceRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifySecurityGroupSequenceRules")
    
    
    return
}

func NewModifySecurityGroupSequenceRulesResponse() (response *ModifySecurityGroupSequenceRulesResponse) {
    response = &ModifySecurityGroupSequenceRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySecurityGroupSequenceRules
// This API is used to sort enterprise security group rules.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) ModifySecurityGroupSequenceRules(request *ModifySecurityGroupSequenceRulesRequest) (response *ModifySecurityGroupSequenceRulesResponse, err error) {
    return c.ModifySecurityGroupSequenceRulesWithContext(context.Background(), request)
}

// ModifySecurityGroupSequenceRules
// This API is used to sort enterprise security group rules.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) ModifySecurityGroupSequenceRulesWithContext(ctx context.Context, request *ModifySecurityGroupSequenceRulesRequest) (response *ModifySecurityGroupSequenceRulesResponse, err error) {
    if request == nil {
        request = NewModifySecurityGroupSequenceRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "ModifySecurityGroupSequenceRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySecurityGroupSequenceRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySecurityGroupSequenceRulesResponse()
    err = c.Send(request, response)
    return
}

func NewModifySequenceRulesRequest() (request *ModifySequenceRulesRequest) {
    request = &ModifySequenceRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifySequenceRules")
    
    
    return
}

func NewModifySequenceRulesResponse() (response *ModifySequenceRulesResponse) {
    response = &ModifySequenceRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySequenceRules
// This API is used to modify rule priority.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifySequenceRules(request *ModifySequenceRulesRequest) (response *ModifySequenceRulesResponse, err error) {
    return c.ModifySequenceRulesWithContext(context.Background(), request)
}

// ModifySequenceRules
// This API is used to modify rule priority.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifySequenceRulesWithContext(ctx context.Context, request *ModifySequenceRulesRequest) (response *ModifySequenceRulesResponse, err error) {
    if request == nil {
        request = NewModifySequenceRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "ModifySequenceRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySequenceRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySequenceRulesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyStorageSettingRequest() (request *ModifyStorageSettingRequest) {
    request = &ModifyStorageSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyStorageSetting")
    
    
    return
}

func NewModifyStorageSettingResponse() (response *ModifyStorageSettingResponse) {
    response = &ModifyStorageSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyStorageSetting
// This API is used to modify the log retention period or to clear logs.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyStorageSetting(request *ModifyStorageSettingRequest) (response *ModifyStorageSettingResponse, err error) {
    return c.ModifyStorageSettingWithContext(context.Background(), request)
}

// ModifyStorageSetting
// This API is used to modify the log retention period or to clear logs.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyStorageSettingWithContext(ctx context.Context, request *ModifyStorageSettingRequest) (response *ModifyStorageSettingResponse, err error) {
    if request == nil {
        request = NewModifyStorageSettingRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "ModifyStorageSetting")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyStorageSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyStorageSettingResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTableStatusRequest() (request *ModifyTableStatusRequest) {
    request = &ModifyTableStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "ModifyTableStatus")
    
    
    return
}

func NewModifyTableStatusResponse() (response *ModifyTableStatusResponse) {
    response = &ModifyTableStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTableStatus
// This API is used to modify rule list status.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyTableStatus(request *ModifyTableStatusRequest) (response *ModifyTableStatusResponse, err error) {
    return c.ModifyTableStatusWithContext(context.Background(), request)
}

// ModifyTableStatus
// This API is used to modify rule list status.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyTableStatusWithContext(ctx context.Context, request *ModifyTableStatusRequest) (response *ModifyTableStatusResponse, err error) {
    if request == nil {
        request = NewModifyTableStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "ModifyTableStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTableStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTableStatusResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveAcRuleRequest() (request *RemoveAcRuleRequest) {
    request = &RemoveAcRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "RemoveAcRule")
    
    
    return
}

func NewRemoveAcRuleResponse() (response *RemoveAcRuleResponse) {
    response = &RemoveAcRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveAcRule
// This API is used to delete edge firewall rules.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RemoveAcRule(request *RemoveAcRuleRequest) (response *RemoveAcRuleResponse, err error) {
    return c.RemoveAcRuleWithContext(context.Background(), request)
}

// RemoveAcRule
// This API is used to delete edge firewall rules.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RemoveAcRuleWithContext(ctx context.Context, request *RemoveAcRuleRequest) (response *RemoveAcRuleResponse, err error) {
    if request == nil {
        request = NewRemoveAcRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "RemoveAcRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveAcRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveAcRuleResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveEnterpriseSecurityGroupRuleRequest() (request *RemoveEnterpriseSecurityGroupRuleRequest) {
    request = &RemoveEnterpriseSecurityGroupRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "RemoveEnterpriseSecurityGroupRule")
    
    
    return
}

func NewRemoveEnterpriseSecurityGroupRuleResponse() (response *RemoveEnterpriseSecurityGroupRuleResponse) {
    response = &RemoveEnterpriseSecurityGroupRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveEnterpriseSecurityGroupRule
// This API is used to delete enterprise security group rules (new).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RemoveEnterpriseSecurityGroupRule(request *RemoveEnterpriseSecurityGroupRuleRequest) (response *RemoveEnterpriseSecurityGroupRuleResponse, err error) {
    return c.RemoveEnterpriseSecurityGroupRuleWithContext(context.Background(), request)
}

// RemoveEnterpriseSecurityGroupRule
// This API is used to delete enterprise security group rules (new).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RemoveEnterpriseSecurityGroupRuleWithContext(ctx context.Context, request *RemoveEnterpriseSecurityGroupRuleRequest) (response *RemoveEnterpriseSecurityGroupRuleResponse, err error) {
    if request == nil {
        request = NewRemoveEnterpriseSecurityGroupRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "RemoveEnterpriseSecurityGroupRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveEnterpriseSecurityGroupRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveEnterpriseSecurityGroupRuleResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveNatAcRuleRequest() (request *RemoveNatAcRuleRequest) {
    request = &RemoveNatAcRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "RemoveNatAcRule")
    
    
    return
}

func NewRemoveNatAcRuleResponse() (response *RemoveNatAcRuleResponse) {
    response = &RemoveNatAcRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveNatAcRule
// This API is used to delete NAT access control rules.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RemoveNatAcRule(request *RemoveNatAcRuleRequest) (response *RemoveNatAcRuleResponse, err error) {
    return c.RemoveNatAcRuleWithContext(context.Background(), request)
}

// RemoveNatAcRule
// This API is used to delete NAT access control rules.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RemoveNatAcRuleWithContext(ctx context.Context, request *RemoveNatAcRuleRequest) (response *RemoveNatAcRuleResponse, err error) {
    if request == nil {
        request = NewRemoveNatAcRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "RemoveNatAcRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveNatAcRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveNatAcRuleResponse()
    err = c.Send(request, response)
    return
}

func NewSetNatFwDnatRuleRequest() (request *SetNatFwDnatRuleRequest) {
    request = &SetNatFwDnatRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "SetNatFwDnatRule")
    
    
    return
}

func NewSetNatFwDnatRuleResponse() (response *SetNatFwDnatRuleResponse) {
    response = &SetNatFwDnatRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetNatFwDnatRule
// This API is used to configure firewall DNAT rules.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SetNatFwDnatRule(request *SetNatFwDnatRuleRequest) (response *SetNatFwDnatRuleResponse, err error) {
    return c.SetNatFwDnatRuleWithContext(context.Background(), request)
}

// SetNatFwDnatRule
// This API is used to configure firewall DNAT rules.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SetNatFwDnatRuleWithContext(ctx context.Context, request *SetNatFwDnatRuleRequest) (response *SetNatFwDnatRuleResponse, err error) {
    if request == nil {
        request = NewSetNatFwDnatRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "SetNatFwDnatRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetNatFwDnatRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetNatFwDnatRuleResponse()
    err = c.Send(request, response)
    return
}

func NewSetNatFwEipRequest() (request *SetNatFwEipRequest) {
    request = &SetNatFwEipRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "SetNatFwEip")
    
    
    return
}

func NewSetNatFwEipResponse() (response *SetNatFwEipResponse) {
    response = &SetNatFwEipResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetNatFwEip
// This API is used to set the firewall instance EIP. (Available for firewall instances in the "Create new" mode. only)
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SetNatFwEip(request *SetNatFwEipRequest) (response *SetNatFwEipResponse, err error) {
    return c.SetNatFwEipWithContext(context.Background(), request)
}

// SetNatFwEip
// This API is used to set the firewall instance EIP. (Available for firewall instances in the "Create new" mode. only)
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SetNatFwEipWithContext(ctx context.Context, request *SetNatFwEipRequest) (response *SetNatFwEipResponse, err error) {
    if request == nil {
        request = NewSetNatFwEipRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "SetNatFwEip")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetNatFwEip require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetNatFwEipResponse()
    err = c.Send(request, response)
    return
}

func NewStopSecurityGroupRuleDispatchRequest() (request *StopSecurityGroupRuleDispatchRequest) {
    request = &StopSecurityGroupRuleDispatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfw", APIVersion, "StopSecurityGroupRuleDispatch")
    
    
    return
}

func NewStopSecurityGroupRuleDispatchResponse() (response *StopSecurityGroupRuleDispatchResponse) {
    response = &StopSecurityGroupRuleDispatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopSecurityGroupRuleDispatch
// This API is used to stop publishing security group rules.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StopSecurityGroupRuleDispatch(request *StopSecurityGroupRuleDispatchRequest) (response *StopSecurityGroupRuleDispatchResponse, err error) {
    return c.StopSecurityGroupRuleDispatchWithContext(context.Background(), request)
}

// StopSecurityGroupRuleDispatch
// This API is used to stop publishing security group rules.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StopSecurityGroupRuleDispatchWithContext(ctx context.Context, request *StopSecurityGroupRuleDispatchRequest) (response *StopSecurityGroupRuleDispatchResponse, err error) {
    if request == nil {
        request = NewStopSecurityGroupRuleDispatchRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfw", APIVersion, "StopSecurityGroupRuleDispatch")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopSecurityGroupRuleDispatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopSecurityGroupRuleDispatchResponse()
    err = c.Send(request, response)
    return
}
