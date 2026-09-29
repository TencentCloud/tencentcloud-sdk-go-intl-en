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

package v20251030

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2025-10-30"

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


func NewAddTargetsToTargetGroupRequest() (request *AddTargetsToTargetGroupRequest) {
    request = &AddTargetsToTargetGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "AddTargetsToTargetGroup")
    
    
    return
}

func NewAddTargetsToTargetGroupResponse() (response *AddTargetsToTargetGroupResponse) {
    response = &AddTargetsToTargetGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddTargetsToTargetGroup
// Add a backend service in the target group.
func (c *Client) AddTargetsToTargetGroup(request *AddTargetsToTargetGroupRequest) (response *AddTargetsToTargetGroupResponse, err error) {
    return c.AddTargetsToTargetGroupWithContext(context.Background(), request)
}

// AddTargetsToTargetGroup
// Add a backend service in the target group.
func (c *Client) AddTargetsToTargetGroupWithContext(ctx context.Context, request *AddTargetsToTargetGroupRequest) (response *AddTargetsToTargetGroupResponse, err error) {
    if request == nil {
        request = NewAddTargetsToTargetGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "AddTargetsToTargetGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddTargetsToTargetGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddTargetsToTargetGroupResponse()
    err = c.Send(request, response)
    return
}

func NewAssociateBandwidthPackageWithLoadBalancerRequest() (request *AssociateBandwidthPackageWithLoadBalancerRequest) {
    request = &AssociateBandwidthPackageWithLoadBalancerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "AssociateBandwidthPackageWithLoadBalancer")
    
    
    return
}

func NewAssociateBandwidthPackageWithLoadBalancerResponse() (response *AssociateBandwidthPackageWithLoadBalancerResponse) {
    response = &AssociateBandwidthPackageWithLoadBalancerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AssociateBandwidthPackageWithLoadBalancer
// Bind a Bandwidth Package to an application CLB instance.
func (c *Client) AssociateBandwidthPackageWithLoadBalancer(request *AssociateBandwidthPackageWithLoadBalancerRequest) (response *AssociateBandwidthPackageWithLoadBalancerResponse, err error) {
    return c.AssociateBandwidthPackageWithLoadBalancerWithContext(context.Background(), request)
}

// AssociateBandwidthPackageWithLoadBalancer
// Bind a Bandwidth Package to an application CLB instance.
func (c *Client) AssociateBandwidthPackageWithLoadBalancerWithContext(ctx context.Context, request *AssociateBandwidthPackageWithLoadBalancerRequest) (response *AssociateBandwidthPackageWithLoadBalancerResponse, err error) {
    if request == nil {
        request = NewAssociateBandwidthPackageWithLoadBalancerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "AssociateBandwidthPackageWithLoadBalancer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssociateBandwidthPackageWithLoadBalancer require credential")
    }

    request.SetContext(ctx)
    
    response = NewAssociateBandwidthPackageWithLoadBalancerResponse()
    err = c.Send(request, response)
    return
}

func NewAssociateListenerAdditionalCertificatesRequest() (request *AssociateListenerAdditionalCertificatesRequest) {
    request = &AssociateListenerAdditionalCertificatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "AssociateListenerAdditionalCertificates")
    
    
    return
}

func NewAssociateListenerAdditionalCertificatesResponse() (response *AssociateListenerAdditionalCertificatesResponse) {
    response = &AssociateListenerAdditionalCertificatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AssociateListenerAdditionalCertificates
// AssociateListenerAdditionalCertificates is an async API. The system returns a request ID, but the additional cert is not yet successfully added. The add task is still in progress in the system backend. You can call the DescribeListenerCertificates API to query the add status of the additional cert.
//
// When HTTPS and QUIC listeners are in Associating status, it means certificate expansion is ongoing.
//
// When HTTPS and QUIC listeners are in the Associated status, the extension cert is successfully added.
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
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CERTIFICATE = "ResourceNotFound.Certificate"
//  RESOURCENOTFOUND_LISTENER = "ResourceNotFound.Listener"
//  RESOURCENOTFOUND_LOADBALANCER = "ResourceNotFound.LoadBalancer"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CERTIFICATE = "ResourceUnavailable.Certificate"
//  RESOURCEUNAVAILABLE_LISTENER = "ResourceUnavailable.Listener"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CERTIFICATEALREADYBOUND = "UnsupportedOperation.CertificateAlreadyBound"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDPROTOCOL = "UnsupportedOperation.UnsupportedProtocol"
func (c *Client) AssociateListenerAdditionalCertificates(request *AssociateListenerAdditionalCertificatesRequest) (response *AssociateListenerAdditionalCertificatesResponse, err error) {
    return c.AssociateListenerAdditionalCertificatesWithContext(context.Background(), request)
}

// AssociateListenerAdditionalCertificates
// AssociateListenerAdditionalCertificates is an async API. The system returns a request ID, but the additional cert is not yet successfully added. The add task is still in progress in the system backend. You can call the DescribeListenerCertificates API to query the add status of the additional cert.
//
// When HTTPS and QUIC listeners are in Associating status, it means certificate expansion is ongoing.
//
// When HTTPS and QUIC listeners are in the Associated status, the extension cert is successfully added.
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
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CERTIFICATE = "ResourceNotFound.Certificate"
//  RESOURCENOTFOUND_LISTENER = "ResourceNotFound.Listener"
//  RESOURCENOTFOUND_LOADBALANCER = "ResourceNotFound.LoadBalancer"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CERTIFICATE = "ResourceUnavailable.Certificate"
//  RESOURCEUNAVAILABLE_LISTENER = "ResourceUnavailable.Listener"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CERTIFICATEALREADYBOUND = "UnsupportedOperation.CertificateAlreadyBound"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDPROTOCOL = "UnsupportedOperation.UnsupportedProtocol"
func (c *Client) AssociateListenerAdditionalCertificatesWithContext(ctx context.Context, request *AssociateListenerAdditionalCertificatesRequest) (response *AssociateListenerAdditionalCertificatesResponse, err error) {
    if request == nil {
        request = NewAssociateListenerAdditionalCertificatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "AssociateListenerAdditionalCertificates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssociateListenerAdditionalCertificates require credential")
    }

    request.SetContext(ctx)
    
    response = NewAssociateListenerAdditionalCertificatesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateHealthCheckTemplateRequest() (request *CreateHealthCheckTemplateRequest) {
    request = &CreateHealthCheckTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "CreateHealthCheckTemplate")
    
    
    return
}

func NewCreateHealthCheckTemplateResponse() (response *CreateHealthCheckTemplateResponse) {
    response = &CreateHealthCheckTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateHealthCheckTemplate
// This API is used to create a health check Template.
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
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CERTIFICATE = "ResourceNotFound.Certificate"
//  RESOURCENOTFOUND_LISTENER = "ResourceNotFound.Listener"
//  RESOURCENOTFOUND_LOADBALANCER = "ResourceNotFound.LoadBalancer"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CERTIFICATE = "ResourceUnavailable.Certificate"
//  RESOURCEUNAVAILABLE_LISTENER = "ResourceUnavailable.Listener"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CERTIFICATEALREADYBOUND = "UnsupportedOperation.CertificateAlreadyBound"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDPROTOCOL = "UnsupportedOperation.UnsupportedProtocol"
func (c *Client) CreateHealthCheckTemplate(request *CreateHealthCheckTemplateRequest) (response *CreateHealthCheckTemplateResponse, err error) {
    return c.CreateHealthCheckTemplateWithContext(context.Background(), request)
}

// CreateHealthCheckTemplate
// This API is used to create a health check Template.
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
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CERTIFICATE = "ResourceNotFound.Certificate"
//  RESOURCENOTFOUND_LISTENER = "ResourceNotFound.Listener"
//  RESOURCENOTFOUND_LOADBALANCER = "ResourceNotFound.LoadBalancer"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CERTIFICATE = "ResourceUnavailable.Certificate"
//  RESOURCEUNAVAILABLE_LISTENER = "ResourceUnavailable.Listener"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CERTIFICATEALREADYBOUND = "UnsupportedOperation.CertificateAlreadyBound"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDPROTOCOL = "UnsupportedOperation.UnsupportedProtocol"
func (c *Client) CreateHealthCheckTemplateWithContext(ctx context.Context, request *CreateHealthCheckTemplateRequest) (response *CreateHealthCheckTemplateResponse, err error) {
    if request == nil {
        request = NewCreateHealthCheckTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "CreateHealthCheckTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateHealthCheckTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateHealthCheckTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateListenerRequest() (request *CreateListenerRequest) {
    request = &CreateListenerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "CreateListener")
    
    
    return
}

func NewCreateListenerResponse() (response *CreateListenerResponse) {
    response = &CreateListenerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateListener
// This API is used to create a listener.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION_TARGETGROUPPROTOCOLMISMATCH = "UnsupportedOperation.TargetGroupProtocolMismatch"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDPROTOCOL = "UnsupportedOperation.UnsupportedProtocol"
func (c *Client) CreateListener(request *CreateListenerRequest) (response *CreateListenerResponse, err error) {
    return c.CreateListenerWithContext(context.Background(), request)
}

// CreateListener
// This API is used to create a listener.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION_TARGETGROUPPROTOCOLMISMATCH = "UnsupportedOperation.TargetGroupProtocolMismatch"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDPROTOCOL = "UnsupportedOperation.UnsupportedProtocol"
func (c *Client) CreateListenerWithContext(ctx context.Context, request *CreateListenerRequest) (response *CreateListenerResponse, err error) {
    if request == nil {
        request = NewCreateListenerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "CreateListener")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateListener require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateListenerResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLoadBalancerRequest() (request *CreateLoadBalancerRequest) {
    request = &CreateLoadBalancerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "CreateLoadBalancer")
    
    
    return
}

func NewCreateLoadBalancerResponse() (response *CreateLoadBalancerResponse) {
    response = &CreateLoadBalancerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLoadBalancer
// **CreateLoadBalancer** is an async API. The system returns an instance ID, but the application CLB instance is not created successfully yet, and the creation task is still in progress in the system backend. You can call [DescribeLoadBalancerDetail](https://www.tencentcloud.com/document/product/1311/84267) to query the creation status of the application CLB instance.
//
// - When an application CLB instance is in the **Provisioning** status, it means the application CLB instance is being created.
//
// -When an application CLB instance is in the **Active** status, the application CLB instance is successfully created.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_RESOURCESSOLDOUT = "UnsupportedOperation.ResourcesSoldOut"
func (c *Client) CreateLoadBalancer(request *CreateLoadBalancerRequest) (response *CreateLoadBalancerResponse, err error) {
    return c.CreateLoadBalancerWithContext(context.Background(), request)
}

// CreateLoadBalancer
// **CreateLoadBalancer** is an async API. The system returns an instance ID, but the application CLB instance is not created successfully yet, and the creation task is still in progress in the system backend. You can call [DescribeLoadBalancerDetail](https://www.tencentcloud.com/document/product/1311/84267) to query the creation status of the application CLB instance.
//
// - When an application CLB instance is in the **Provisioning** status, it means the application CLB instance is being created.
//
// -When an application CLB instance is in the **Active** status, the application CLB instance is successfully created.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_RESOURCESSOLDOUT = "UnsupportedOperation.ResourcesSoldOut"
func (c *Client) CreateLoadBalancerWithContext(ctx context.Context, request *CreateLoadBalancerRequest) (response *CreateLoadBalancerResponse, err error) {
    if request == nil {
        request = NewCreateLoadBalancerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "CreateLoadBalancer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLoadBalancer require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLoadBalancerResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRulesRequest() (request *CreateRulesRequest) {
    request = &CreateRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "CreateRules")
    
    
    return
}

func NewCreateRulesResponse() (response *CreateRulesResponse) {
    response = &CreateRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRules
// This API is used to create forwarding rules. It is an async API. After returning successfully, call the DescribeAsyncJobs API with the returned RequestID as an input parameter to check whether this task is successful.
//
// A rule supports up to 10 forward Conditions and 5 forward Actions.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCEALREADYEXISTS = "FailedOperation.ResourceAlreadyExists"
//  FAILEDOPERATION_RULEALREADYEXISTS = "FailedOperation.RuleAlreadyExists"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCLIENTTOKEN = "InvalidParameter.InvalidClientToken"
//  INVALIDPARAMETER_INVALIDFIELDFORMAT = "InvalidParameter.InvalidFieldFormat"
//  INVALIDPARAMETER_INVALIDFIELDTYPE = "InvalidParameter.InvalidFieldType"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETER_MISSINGACTION = "InvalidParameter.MissingAction"
//  INVALIDPARAMETER_MISSINGFIELD = "InvalidParameter.MissingField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BUSINESSPARAMETERMISMATCH = "InvalidParameterValue.BusinessParameterMismatch"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_LISTENER = "ResourceNotFound.Listener"
//  RESOURCENOTFOUND_LOADBALANCER = "ResourceNotFound.LoadBalancer"
//  RESOURCENOTFOUND_TARGETGROUP = "ResourceNotFound.TargetGroup"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_LISTENER = "ResourceUnavailable.Listener"
//  RESOURCEUNAVAILABLE_LOADBALANCER = "ResourceUnavailable.LoadBalancer"
//  RESOURCEUNAVAILABLE_PREVIOUSTASKNOTCOMPLETED = "ResourceUnavailable.PreviousTaskNotCompleted"
//  RESOURCEUNAVAILABLE_TARGETGROUP = "ResourceUnavailable.TargetGroup"
//  RESOURCEUNAVAILABLE_UPSTREAMSTATUSBLOCKED = "ResourceUnavailable.UpstreamStatusBlocked"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICEROLEUNAUTHORIZED = "UnauthorizedOperation.ServiceRoleUnauthorized"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDSTATETRANSITION = "UnsupportedOperation.InvalidStateTransition"
func (c *Client) CreateRules(request *CreateRulesRequest) (response *CreateRulesResponse, err error) {
    return c.CreateRulesWithContext(context.Background(), request)
}

// CreateRules
// This API is used to create forwarding rules. It is an async API. After returning successfully, call the DescribeAsyncJobs API with the returned RequestID as an input parameter to check whether this task is successful.
//
// A rule supports up to 10 forward Conditions and 5 forward Actions.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCEALREADYEXISTS = "FailedOperation.ResourceAlreadyExists"
//  FAILEDOPERATION_RULEALREADYEXISTS = "FailedOperation.RuleAlreadyExists"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCLIENTTOKEN = "InvalidParameter.InvalidClientToken"
//  INVALIDPARAMETER_INVALIDFIELDFORMAT = "InvalidParameter.InvalidFieldFormat"
//  INVALIDPARAMETER_INVALIDFIELDTYPE = "InvalidParameter.InvalidFieldType"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETER_MISSINGACTION = "InvalidParameter.MissingAction"
//  INVALIDPARAMETER_MISSINGFIELD = "InvalidParameter.MissingField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BUSINESSPARAMETERMISMATCH = "InvalidParameterValue.BusinessParameterMismatch"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_LISTENER = "ResourceNotFound.Listener"
//  RESOURCENOTFOUND_LOADBALANCER = "ResourceNotFound.LoadBalancer"
//  RESOURCENOTFOUND_TARGETGROUP = "ResourceNotFound.TargetGroup"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_LISTENER = "ResourceUnavailable.Listener"
//  RESOURCEUNAVAILABLE_LOADBALANCER = "ResourceUnavailable.LoadBalancer"
//  RESOURCEUNAVAILABLE_PREVIOUSTASKNOTCOMPLETED = "ResourceUnavailable.PreviousTaskNotCompleted"
//  RESOURCEUNAVAILABLE_TARGETGROUP = "ResourceUnavailable.TargetGroup"
//  RESOURCEUNAVAILABLE_UPSTREAMSTATUSBLOCKED = "ResourceUnavailable.UpstreamStatusBlocked"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICEROLEUNAUTHORIZED = "UnauthorizedOperation.ServiceRoleUnauthorized"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDSTATETRANSITION = "UnsupportedOperation.InvalidStateTransition"
func (c *Client) CreateRulesWithContext(ctx context.Context, request *CreateRulesRequest) (response *CreateRulesResponse, err error) {
    if request == nil {
        request = NewCreateRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "CreateRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRulesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSecurityPolicyRequest() (request *CreateSecurityPolicyRequest) {
    request = &CreateSecurityPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "CreateSecurityPolicy")
    
    
    return
}

func NewCreateSecurityPolicyResponse() (response *CreateSecurityPolicyResponse) {
    response = &CreateSecurityPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSecurityPolicy
// Create a custom security policy for configuring the TLS protocol version and encryption suite of an HTTPS listener. With a security policy, you can flexibly control the security level of HTTPS communication between clients and load balancing.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSecurityPolicy(request *CreateSecurityPolicyRequest) (response *CreateSecurityPolicyResponse, err error) {
    return c.CreateSecurityPolicyWithContext(context.Background(), request)
}

// CreateSecurityPolicy
// Create a custom security policy for configuring the TLS protocol version and encryption suite of an HTTPS listener. With a security policy, you can flexibly control the security level of HTTPS communication between clients and load balancing.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSecurityPolicyWithContext(ctx context.Context, request *CreateSecurityPolicyRequest) (response *CreateSecurityPolicyResponse, err error) {
    if request == nil {
        request = NewCreateSecurityPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "CreateSecurityPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSecurityPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSecurityPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTargetGroupRequest() (request *CreateTargetGroupRequest) {
    request = &CreateTargetGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "CreateTargetGroup")
    
    
    return
}

func NewCreateTargetGroupResponse() (response *CreateTargetGroupResponse) {
    response = &CreateTargetGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTargetGroup
// Target Group APIs
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateTargetGroup(request *CreateTargetGroupRequest) (response *CreateTargetGroupResponse, err error) {
    return c.CreateTargetGroupWithContext(context.Background(), request)
}

// CreateTargetGroup
// Target Group APIs
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateTargetGroupWithContext(ctx context.Context, request *CreateTargetGroupRequest) (response *CreateTargetGroupResponse, err error) {
    if request == nil {
        request = NewCreateTargetGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "CreateTargetGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTargetGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTargetGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteHealthCheckTemplatesRequest() (request *DeleteHealthCheckTemplatesRequest) {
    request = &DeleteHealthCheckTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DeleteHealthCheckTemplates")
    
    
    return
}

func NewDeleteHealthCheckTemplatesResponse() (response *DeleteHealthCheckTemplatesResponse) {
    response = &DeleteHealthCheckTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteHealthCheckTemplates
// Deletes a health check Template
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteHealthCheckTemplates(request *DeleteHealthCheckTemplatesRequest) (response *DeleteHealthCheckTemplatesResponse, err error) {
    return c.DeleteHealthCheckTemplatesWithContext(context.Background(), request)
}

// DeleteHealthCheckTemplates
// Deletes a health check Template
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteHealthCheckTemplatesWithContext(ctx context.Context, request *DeleteHealthCheckTemplatesRequest) (response *DeleteHealthCheckTemplatesResponse, err error) {
    if request == nil {
        request = NewDeleteHealthCheckTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DeleteHealthCheckTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteHealthCheckTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteHealthCheckTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteListenerRequest() (request *DeleteListenerRequest) {
    request = &DeleteListenerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DeleteListener")
    
    
    return
}

func NewDeleteListenerResponse() (response *DeleteListenerResponse) {
    response = &DeleteListenerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteListener
// Delete a listener
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteListener(request *DeleteListenerRequest) (response *DeleteListenerResponse, err error) {
    return c.DeleteListenerWithContext(context.Background(), request)
}

// DeleteListener
// Delete a listener
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteListenerWithContext(ctx context.Context, request *DeleteListenerRequest) (response *DeleteListenerResponse, err error) {
    if request == nil {
        request = NewDeleteListenerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DeleteListener")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteListener require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteListenerResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLoadBalancersRequest() (request *DeleteLoadBalancersRequest) {
    request = &DeleteLoadBalancersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DeleteLoadBalancers")
    
    
    return
}

func NewDeleteLoadBalancersResponse() (response *DeleteLoadBalancersResponse) {
    response = &DeleteLoadBalancersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLoadBalancers
// The **DeleteLoadBalancers** API is an async API. The system returns a request ID, but the application CLB instance is not yet deleted successfully. The deletion task is still in progress in the system backend. You can call [DescribeLoadBalancerDetail](https://www.tencentcloud.com/document/product/1311/84267) to query the deletion status of the application CLB instance.
//
// - When an application CLB instance is in the **Deleting** status, it means the application CLB instance is being deleted.
//
// -If the specified application CLB instance cannot be queried, the application CLB instance has been deleted successfully.
//
// error code that may be returned:
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE_LOADBALANCERHASLISTENERS = "ResourceInUse.LoadBalancerHasListeners"
func (c *Client) DeleteLoadBalancers(request *DeleteLoadBalancersRequest) (response *DeleteLoadBalancersResponse, err error) {
    return c.DeleteLoadBalancersWithContext(context.Background(), request)
}

// DeleteLoadBalancers
// The **DeleteLoadBalancers** API is an async API. The system returns a request ID, but the application CLB instance is not yet deleted successfully. The deletion task is still in progress in the system backend. You can call [DescribeLoadBalancerDetail](https://www.tencentcloud.com/document/product/1311/84267) to query the deletion status of the application CLB instance.
//
// - When an application CLB instance is in the **Deleting** status, it means the application CLB instance is being deleted.
//
// -If the specified application CLB instance cannot be queried, the application CLB instance has been deleted successfully.
//
// error code that may be returned:
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE_LOADBALANCERHASLISTENERS = "ResourceInUse.LoadBalancerHasListeners"
func (c *Client) DeleteLoadBalancersWithContext(ctx context.Context, request *DeleteLoadBalancersRequest) (response *DeleteLoadBalancersResponse, err error) {
    if request == nil {
        request = NewDeleteLoadBalancersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DeleteLoadBalancers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLoadBalancers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLoadBalancersResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRulesRequest() (request *DeleteRulesRequest) {
    request = &DeleteRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DeleteRules")
    
    
    return
}

func NewDeleteRulesResponse() (response *DeleteRulesResponse) {
    response = &DeleteRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRules
// DeleteRules deletes forwarding rules. This is an async API. After returning successfully, call the DescribeAsyncJobs API with the returned RequestID as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCLIENTTOKEN = "InvalidParameter.InvalidClientToken"
//  INVALIDPARAMETER_INVALIDFIELDFORMAT = "InvalidParameter.InvalidFieldFormat"
//  INVALIDPARAMETER_INVALIDFIELDTYPE = "InvalidParameter.InvalidFieldType"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETER_MISSINGACTION = "InvalidParameter.MissingAction"
//  INVALIDPARAMETER_MISSINGFIELD = "InvalidParameter.MissingField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_LISTENER = "ResourceNotFound.Listener"
//  RESOURCENOTFOUND_LOADBALANCER = "ResourceNotFound.LoadBalancer"
//  RESOURCENOTFOUND_RULE = "ResourceNotFound.Rule"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_LISTENER = "ResourceUnavailable.Listener"
//  RESOURCEUNAVAILABLE_LOADBALANCER = "ResourceUnavailable.LoadBalancer"
//  RESOURCEUNAVAILABLE_PREVIOUSTASKNOTCOMPLETED = "ResourceUnavailable.PreviousTaskNotCompleted"
//  RESOURCEUNAVAILABLE_RULE = "ResourceUnavailable.Rule"
//  RESOURCEUNAVAILABLE_UPSTREAMSTATUSBLOCKED = "ResourceUnavailable.UpstreamStatusBlocked"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICEROLEUNAUTHORIZED = "UnauthorizedOperation.ServiceRoleUnauthorized"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRules(request *DeleteRulesRequest) (response *DeleteRulesResponse, err error) {
    return c.DeleteRulesWithContext(context.Background(), request)
}

// DeleteRules
// DeleteRules deletes forwarding rules. This is an async API. After returning successfully, call the DescribeAsyncJobs API with the returned RequestID as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCLIENTTOKEN = "InvalidParameter.InvalidClientToken"
//  INVALIDPARAMETER_INVALIDFIELDFORMAT = "InvalidParameter.InvalidFieldFormat"
//  INVALIDPARAMETER_INVALIDFIELDTYPE = "InvalidParameter.InvalidFieldType"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETER_MISSINGACTION = "InvalidParameter.MissingAction"
//  INVALIDPARAMETER_MISSINGFIELD = "InvalidParameter.MissingField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_LISTENER = "ResourceNotFound.Listener"
//  RESOURCENOTFOUND_LOADBALANCER = "ResourceNotFound.LoadBalancer"
//  RESOURCENOTFOUND_RULE = "ResourceNotFound.Rule"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_LISTENER = "ResourceUnavailable.Listener"
//  RESOURCEUNAVAILABLE_LOADBALANCER = "ResourceUnavailable.LoadBalancer"
//  RESOURCEUNAVAILABLE_PREVIOUSTASKNOTCOMPLETED = "ResourceUnavailable.PreviousTaskNotCompleted"
//  RESOURCEUNAVAILABLE_RULE = "ResourceUnavailable.Rule"
//  RESOURCEUNAVAILABLE_UPSTREAMSTATUSBLOCKED = "ResourceUnavailable.UpstreamStatusBlocked"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICEROLEUNAUTHORIZED = "UnauthorizedOperation.ServiceRoleUnauthorized"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRulesWithContext(ctx context.Context, request *DeleteRulesRequest) (response *DeleteRulesResponse, err error) {
    if request == nil {
        request = NewDeleteRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DeleteRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSecurityPolicyRequest() (request *DeleteSecurityPolicyRequest) {
    request = &DeleteSecurityPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DeleteSecurityPolicy")
    
    
    return
}

func NewDeleteSecurityPolicyResponse() (response *DeleteSecurityPolicyResponse) {
    response = &DeleteSecurityPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSecurityPolicy
// Delete one or more custom security policies. Before deletion, please ensure the policy hasn't been referenced by any HTTPS listener, otherwise the deletion will fail.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteSecurityPolicy(request *DeleteSecurityPolicyRequest) (response *DeleteSecurityPolicyResponse, err error) {
    return c.DeleteSecurityPolicyWithContext(context.Background(), request)
}

// DeleteSecurityPolicy
// Delete one or more custom security policies. Before deletion, please ensure the policy hasn't been referenced by any HTTPS listener, otherwise the deletion will fail.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteSecurityPolicyWithContext(ctx context.Context, request *DeleteSecurityPolicyRequest) (response *DeleteSecurityPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteSecurityPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DeleteSecurityPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSecurityPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSecurityPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTargetGroupsRequest() (request *DeleteTargetGroupsRequest) {
    request = &DeleteTargetGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DeleteTargetGroups")
    
    
    return
}

func NewDeleteTargetGroupsResponse() (response *DeleteTargetGroupsResponse) {
    response = &DeleteTargetGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTargetGroups
// Delete a target group.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteTargetGroups(request *DeleteTargetGroupsRequest) (response *DeleteTargetGroupsResponse, err error) {
    return c.DeleteTargetGroupsWithContext(context.Background(), request)
}

// DeleteTargetGroups
// Delete a target group.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteTargetGroupsWithContext(ctx context.Context, request *DeleteTargetGroupsRequest) (response *DeleteTargetGroupsResponse, err error) {
    if request == nil {
        request = NewDeleteTargetGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DeleteTargetGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTargetGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTargetGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAsyncJobsRequest() (request *DescribeAsyncJobsRequest) {
    request = &DescribeAsyncJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DescribeAsyncJobs")
    
    
    return
}

func NewDescribeAsyncJobsResponse() (response *DescribeAsyncJobsResponse) {
    response = &DescribeAsyncJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAsyncJobs
// Query API for async tasks
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAsyncJobs(request *DescribeAsyncJobsRequest) (response *DescribeAsyncJobsResponse, err error) {
    return c.DescribeAsyncJobsWithContext(context.Background(), request)
}

// DescribeAsyncJobs
// Query API for async tasks
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAsyncJobsWithContext(ctx context.Context, request *DescribeAsyncJobsRequest) (response *DescribeAsyncJobsResponse, err error) {
    if request == nil {
        request = NewDescribeAsyncJobsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DescribeAsyncJobs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAsyncJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAsyncJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHealthCheckTemplatesRequest() (request *DescribeHealthCheckTemplatesRequest) {
    request = &DescribeHealthCheckTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DescribeHealthCheckTemplates")
    
    
    return
}

func NewDescribeHealthCheckTemplatesResponse() (response *DescribeHealthCheckTemplatesResponse) {
    response = &DescribeHealthCheckTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHealthCheckTemplates
// This API is used to query the health check template list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeHealthCheckTemplates(request *DescribeHealthCheckTemplatesRequest) (response *DescribeHealthCheckTemplatesResponse, err error) {
    return c.DescribeHealthCheckTemplatesWithContext(context.Background(), request)
}

// DescribeHealthCheckTemplates
// This API is used to query the health check template list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeHealthCheckTemplatesWithContext(ctx context.Context, request *DescribeHealthCheckTemplatesRequest) (response *DescribeHealthCheckTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeHealthCheckTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DescribeHealthCheckTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHealthCheckTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHealthCheckTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListenerCertificatesRequest() (request *DescribeListenerCertificatesRequest) {
    request = &DescribeListenerCertificatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DescribeListenerCertificates")
    
    
    return
}

func NewDescribeListenerCertificatesResponse() (response *DescribeListenerCertificatesResponse) {
    response = &DescribeListenerCertificatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeListenerCertificates
// This API is used to query the list of certificates bound to a specified listener by instance id and listener id.
//
// If `CertificateType` is set to `SVR`, the information of the extended server certificate and the default server certificate is returned.
//
// If CertificateType is set to CA, the default CA certificate info is returned.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListenerCertificates(request *DescribeListenerCertificatesRequest) (response *DescribeListenerCertificatesResponse, err error) {
    return c.DescribeListenerCertificatesWithContext(context.Background(), request)
}

// DescribeListenerCertificates
// This API is used to query the list of certificates bound to a specified listener by instance id and listener id.
//
// If `CertificateType` is set to `SVR`, the information of the extended server certificate and the default server certificate is returned.
//
// If CertificateType is set to CA, the default CA certificate info is returned.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListenerCertificatesWithContext(ctx context.Context, request *DescribeListenerCertificatesRequest) (response *DescribeListenerCertificatesResponse, err error) {
    if request == nil {
        request = NewDescribeListenerCertificatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DescribeListenerCertificates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeListenerCertificates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeListenerCertificatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListenerDetailRequest() (request *DescribeListenerDetailRequest) {
    request = &DescribeListenerDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DescribeListenerDetail")
    
    
    return
}

func NewDescribeListenerDetailResponse() (response *DescribeListenerDetailResponse) {
    response = &DescribeListenerDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeListenerDetail
// Queries details of one listener.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListenerDetail(request *DescribeListenerDetailRequest) (response *DescribeListenerDetailResponse, err error) {
    return c.DescribeListenerDetailWithContext(context.Background(), request)
}

// DescribeListenerDetail
// Queries details of one listener.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListenerDetailWithContext(ctx context.Context, request *DescribeListenerDetailRequest) (response *DescribeListenerDetailResponse, err error) {
    if request == nil {
        request = NewDescribeListenerDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DescribeListenerDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeListenerDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeListenerDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListenerHealthStatusRequest() (request *DescribeListenerHealthStatusRequest) {
    request = &DescribeListenerHealthStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DescribeListenerHealthStatus")
    
    
    return
}

func NewDescribeListenerHealthStatusResponse() (response *DescribeListenerHealthStatusResponse) {
    response = &DescribeListenerHealthStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeListenerHealthStatus
// Queries the health status of a listener.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListenerHealthStatus(request *DescribeListenerHealthStatusRequest) (response *DescribeListenerHealthStatusResponse, err error) {
    return c.DescribeListenerHealthStatusWithContext(context.Background(), request)
}

// DescribeListenerHealthStatus
// Queries the health status of a listener.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListenerHealthStatusWithContext(ctx context.Context, request *DescribeListenerHealthStatusRequest) (response *DescribeListenerHealthStatusResponse, err error) {
    if request == nil {
        request = NewDescribeListenerHealthStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DescribeListenerHealthStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeListenerHealthStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeListenerHealthStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListenersRequest() (request *DescribeListenersRequest) {
    request = &DescribeListenersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DescribeListeners")
    
    
    return
}

func NewDescribeListenersResponse() (response *DescribeListenersResponse) {
    response = &DescribeListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeListeners
// Queries the listener list
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListeners(request *DescribeListenersRequest) (response *DescribeListenersResponse, err error) {
    return c.DescribeListenersWithContext(context.Background(), request)
}

// DescribeListeners
// Queries the listener list
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListenersWithContext(ctx context.Context, request *DescribeListenersRequest) (response *DescribeListenersResponse, err error) {
    if request == nil {
        request = NewDescribeListenersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DescribeListeners")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeListeners require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeListenersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLoadBalancerDetailRequest() (request *DescribeLoadBalancerDetailRequest) {
    request = &DescribeLoadBalancerDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DescribeLoadBalancerDetail")
    
    
    return
}

func NewDescribeLoadBalancerDetailResponse() (response *DescribeLoadBalancerDetailResponse) {
    response = &DescribeLoadBalancerDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLoadBalancerDetail
// Queries detailed information of a specified load balancing instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeLoadBalancerDetail(request *DescribeLoadBalancerDetailRequest) (response *DescribeLoadBalancerDetailResponse, err error) {
    return c.DescribeLoadBalancerDetailWithContext(context.Background(), request)
}

// DescribeLoadBalancerDetail
// Queries detailed information of a specified load balancing instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeLoadBalancerDetailWithContext(ctx context.Context, request *DescribeLoadBalancerDetailRequest) (response *DescribeLoadBalancerDetailResponse, err error) {
    if request == nil {
        request = NewDescribeLoadBalancerDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DescribeLoadBalancerDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLoadBalancerDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLoadBalancerDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLoadBalancersRequest() (request *DescribeLoadBalancersRequest) {
    request = &DescribeLoadBalancersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DescribeLoadBalancers")
    
    
    return
}

func NewDescribeLoadBalancersResponse() (response *DescribeLoadBalancersResponse) {
    response = &DescribeLoadBalancersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLoadBalancers
// Query instance configuration.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeLoadBalancers(request *DescribeLoadBalancersRequest) (response *DescribeLoadBalancersResponse, err error) {
    return c.DescribeLoadBalancersWithContext(context.Background(), request)
}

// DescribeLoadBalancers
// Query instance configuration.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeLoadBalancersWithContext(ctx context.Context, request *DescribeLoadBalancersRequest) (response *DescribeLoadBalancersResponse, err error) {
    if request == nil {
        request = NewDescribeLoadBalancersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DescribeLoadBalancers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLoadBalancers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLoadBalancersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQuotaRequest() (request *DescribeQuotaRequest) {
    request = &DescribeQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DescribeQuota")
    
    
    return
}

func NewDescribeQuotaResponse() (response *DescribeQuotaResponse) {
    response = &DescribeQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeQuota
// Queries the ALB quota configuration of the current account. It supports querying by quota type and allows you to pass a resource ID to query resource-level quotas. You can use DisplayFields to return the used amount and remaining available quantity as needed.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeQuota(request *DescribeQuotaRequest) (response *DescribeQuotaResponse, err error) {
    return c.DescribeQuotaWithContext(context.Background(), request)
}

// DescribeQuota
// Queries the ALB quota configuration of the current account. It supports querying by quota type and allows you to pass a resource ID to query resource-level quotas. You can use DisplayFields to return the used amount and remaining available quantity as needed.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeQuotaWithContext(ctx context.Context, request *DescribeQuotaRequest) (response *DescribeQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeQuotaRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DescribeQuota")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeQuota require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRulesRequest() (request *DescribeRulesRequest) {
    request = &DescribeRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DescribeRules")
    
    
    return
}

func NewDescribeRulesResponse() (response *DescribeRulesResponse) {
    response = &DescribeRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRules
// This API is used to query forwarding rules.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFIELDTYPE = "InvalidParameter.InvalidFieldType"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETER_MISSINGACTION = "InvalidParameter.MissingAction"
//  INVALIDPARAMETER_MISSINGFIELD = "InvalidParameter.MissingField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_LISTENER = "ResourceNotFound.Listener"
//  RESOURCENOTFOUND_LOADBALANCER = "ResourceNotFound.LoadBalancer"
//  RESOURCENOTFOUND_RULE = "ResourceNotFound.Rule"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICEROLEUNAUTHORIZED = "UnauthorizedOperation.ServiceRoleUnauthorized"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRules(request *DescribeRulesRequest) (response *DescribeRulesResponse, err error) {
    return c.DescribeRulesWithContext(context.Background(), request)
}

// DescribeRules
// This API is used to query forwarding rules.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFIELDTYPE = "InvalidParameter.InvalidFieldType"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETER_MISSINGACTION = "InvalidParameter.MissingAction"
//  INVALIDPARAMETER_MISSINGFIELD = "InvalidParameter.MissingField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_LISTENER = "ResourceNotFound.Listener"
//  RESOURCENOTFOUND_LOADBALANCER = "ResourceNotFound.LoadBalancer"
//  RESOURCENOTFOUND_RULE = "ResourceNotFound.Rule"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICEROLEUNAUTHORIZED = "UnauthorizedOperation.ServiceRoleUnauthorized"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRulesWithContext(ctx context.Context, request *DescribeRulesRequest) (response *DescribeRulesResponse, err error) {
    if request == nil {
        request = NewDescribeRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DescribeRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityPoliciesRequest() (request *DescribeSecurityPoliciesRequest) {
    request = &DescribeSecurityPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DescribeSecurityPolicies")
    
    
    return
}

func NewDescribeSecurityPoliciesResponse() (response *DescribeSecurityPoliciesResponse) {
    response = &DescribeSecurityPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecurityPolicies
// Queries the custom security policy list, supports filtering by security policy ID, name, or tag, and supports paging query.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSecurityPolicies(request *DescribeSecurityPoliciesRequest) (response *DescribeSecurityPoliciesResponse, err error) {
    return c.DescribeSecurityPoliciesWithContext(context.Background(), request)
}

// DescribeSecurityPolicies
// Queries the custom security policy list, supports filtering by security policy ID, name, or tag, and supports paging query.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSecurityPoliciesWithContext(ctx context.Context, request *DescribeSecurityPoliciesRequest) (response *DescribeSecurityPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityPoliciesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DescribeSecurityPolicies")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityPolicyCapabilitiesRequest() (request *DescribeSecurityPolicyCapabilitiesRequest) {
    request = &DescribeSecurityPolicyCapabilitiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DescribeSecurityPolicyCapabilities")
    
    
    return
}

func NewDescribeSecurityPolicyCapabilitiesResponse() (response *DescribeSecurityPolicyCapabilitiesResponse) {
    response = &DescribeSecurityPolicyCapabilitiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecurityPolicyCapabilities
// Query the security policy configuration capacity supported in the current region, including optional TLS protocol versions and the encryption suite list for each version. Before creating or modifying a custom security policy, call this API to get available configuration options.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSecurityPolicyCapabilities(request *DescribeSecurityPolicyCapabilitiesRequest) (response *DescribeSecurityPolicyCapabilitiesResponse, err error) {
    return c.DescribeSecurityPolicyCapabilitiesWithContext(context.Background(), request)
}

// DescribeSecurityPolicyCapabilities
// Query the security policy configuration capacity supported in the current region, including optional TLS protocol versions and the encryption suite list for each version. Before creating or modifying a custom security policy, call this API to get available configuration options.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSecurityPolicyCapabilitiesWithContext(ctx context.Context, request *DescribeSecurityPolicyCapabilitiesRequest) (response *DescribeSecurityPolicyCapabilitiesResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityPolicyCapabilitiesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DescribeSecurityPolicyCapabilities")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityPolicyCapabilities require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityPolicyCapabilitiesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityPolicyRelationsRequest() (request *DescribeSecurityPolicyRelationsRequest) {
    request = &DescribeSecurityPolicyRelationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DescribeSecurityPolicyRelations")
    
    
    return
}

func NewDescribeSecurityPolicyRelationsResponse() (response *DescribeSecurityPolicyRelationsResponse) {
    response = &DescribeSecurityPolicyRelationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecurityPolicyRelations
// Query the relationship between a security policy and the HTTPS listeners that refer to it. Before deleting or modifying a security policy, it is advisable to call this API to confirm the impact.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSecurityPolicyRelations(request *DescribeSecurityPolicyRelationsRequest) (response *DescribeSecurityPolicyRelationsResponse, err error) {
    return c.DescribeSecurityPolicyRelationsWithContext(context.Background(), request)
}

// DescribeSecurityPolicyRelations
// Query the relationship between a security policy and the HTTPS listeners that refer to it. Before deleting or modifying a security policy, it is advisable to call this API to confirm the impact.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSecurityPolicyRelationsWithContext(ctx context.Context, request *DescribeSecurityPolicyRelationsRequest) (response *DescribeSecurityPolicyRelationsResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityPolicyRelationsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DescribeSecurityPolicyRelations")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityPolicyRelations require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityPolicyRelationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSystemSecurityPoliciesRequest() (request *DescribeSystemSecurityPoliciesRequest) {
    request = &DescribeSystemSecurityPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DescribeSystemSecurityPolicies")
    
    
    return
}

func NewDescribeSystemSecurityPoliciesResponse() (response *DescribeSystemSecurityPoliciesResponse) {
    response = &DescribeSystemSecurityPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSystemSecurityPolicies
// Queries system security policies.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSystemSecurityPolicies(request *DescribeSystemSecurityPoliciesRequest) (response *DescribeSystemSecurityPoliciesResponse, err error) {
    return c.DescribeSystemSecurityPoliciesWithContext(context.Background(), request)
}

// DescribeSystemSecurityPolicies
// Queries system security policies.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSystemSecurityPoliciesWithContext(ctx context.Context, request *DescribeSystemSecurityPoliciesRequest) (response *DescribeSystemSecurityPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeSystemSecurityPoliciesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DescribeSystemSecurityPolicies")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSystemSecurityPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSystemSecurityPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTargetGroupTargetsRequest() (request *DescribeTargetGroupTargetsRequest) {
    request = &DescribeTargetGroupTargetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DescribeTargetGroupTargets")
    
    
    return
}

func NewDescribeTargetGroupTargetsResponse() (response *DescribeTargetGroupTargetsResponse) {
    response = &DescribeTargetGroupTargetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTargetGroupTargets
// Queries backend services in the target group.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTargetGroupTargets(request *DescribeTargetGroupTargetsRequest) (response *DescribeTargetGroupTargetsResponse, err error) {
    return c.DescribeTargetGroupTargetsWithContext(context.Background(), request)
}

// DescribeTargetGroupTargets
// Queries backend services in the target group.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTargetGroupTargetsWithContext(ctx context.Context, request *DescribeTargetGroupTargetsRequest) (response *DescribeTargetGroupTargetsResponse, err error) {
    if request == nil {
        request = NewDescribeTargetGroupTargetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DescribeTargetGroupTargets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTargetGroupTargets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTargetGroupTargetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTargetGroupsRequest() (request *DescribeTargetGroupsRequest) {
    request = &DescribeTargetGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DescribeTargetGroups")
    
    
    return
}

func NewDescribeTargetGroupsResponse() (response *DescribeTargetGroupsResponse) {
    response = &DescribeTargetGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTargetGroups
// Query the target group list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTargetGroups(request *DescribeTargetGroupsRequest) (response *DescribeTargetGroupsResponse, err error) {
    return c.DescribeTargetGroupsWithContext(context.Background(), request)
}

// DescribeTargetGroups
// Query the target group list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTargetGroupsWithContext(ctx context.Context, request *DescribeTargetGroupsRequest) (response *DescribeTargetGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeTargetGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DescribeTargetGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTargetGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTargetGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTargetGroupsByTargetRequest() (request *DescribeTargetGroupsByTargetRequest) {
    request = &DescribeTargetGroupsByTargetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DescribeTargetGroupsByTarget")
    
    
    return
}

func NewDescribeTargetGroupsByTargetResponse() (response *DescribeTargetGroupsByTargetResponse) {
    response = &DescribeTargetGroupsByTargetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTargetGroupsByTarget
// Query bound target groups based on the slave machine.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTargetGroupsByTarget(request *DescribeTargetGroupsByTargetRequest) (response *DescribeTargetGroupsByTargetResponse, err error) {
    return c.DescribeTargetGroupsByTargetWithContext(context.Background(), request)
}

// DescribeTargetGroupsByTarget
// Query bound target groups based on the slave machine.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTargetGroupsByTargetWithContext(ctx context.Context, request *DescribeTargetGroupsByTargetRequest) (response *DescribeTargetGroupsByTargetResponse, err error) {
    if request == nil {
        request = NewDescribeTargetGroupsByTargetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DescribeTargetGroupsByTarget")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTargetGroupsByTarget require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTargetGroupsByTargetResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZonesRequest() (request *DescribeZonesRequest) {
    request = &DescribeZonesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DescribeZones")
    
    
    return
}

func NewDescribeZonesResponse() (response *DescribeZonesResponse) {
    response = &DescribeZonesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeZones
// Querying Availability Zones
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeZones(request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    return c.DescribeZonesWithContext(context.Background(), request)
}

// DescribeZones
// Querying Availability Zones
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeZonesWithContext(ctx context.Context, request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    if request == nil {
        request = NewDescribeZonesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DescribeZones")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZones require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeZonesResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateBandwidthPackageFromLoadBalancerRequest() (request *DisassociateBandwidthPackageFromLoadBalancerRequest) {
    request = &DisassociateBandwidthPackageFromLoadBalancerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DisassociateBandwidthPackageFromLoadBalancer")
    
    
    return
}

func NewDisassociateBandwidthPackageFromLoadBalancerResponse() (response *DisassociateBandwidthPackageFromLoadBalancerResponse) {
    response = &DisassociateBandwidthPackageFromLoadBalancerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisassociateBandwidthPackageFromLoadBalancer
// Unbind a Bandwidth Package from an application CLB instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DisassociateBandwidthPackageFromLoadBalancer(request *DisassociateBandwidthPackageFromLoadBalancerRequest) (response *DisassociateBandwidthPackageFromLoadBalancerResponse, err error) {
    return c.DisassociateBandwidthPackageFromLoadBalancerWithContext(context.Background(), request)
}

// DisassociateBandwidthPackageFromLoadBalancer
// Unbind a Bandwidth Package from an application CLB instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DisassociateBandwidthPackageFromLoadBalancerWithContext(ctx context.Context, request *DisassociateBandwidthPackageFromLoadBalancerRequest) (response *DisassociateBandwidthPackageFromLoadBalancerResponse, err error) {
    if request == nil {
        request = NewDisassociateBandwidthPackageFromLoadBalancerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DisassociateBandwidthPackageFromLoadBalancer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisassociateBandwidthPackageFromLoadBalancer require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisassociateBandwidthPackageFromLoadBalancerResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateListenerAdditionalCertificatesRequest() (request *DisassociateListenerAdditionalCertificatesRequest) {
    request = &DisassociateListenerAdditionalCertificatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "DisassociateListenerAdditionalCertificates")
    
    
    return
}

func NewDisassociateListenerAdditionalCertificatesResponse() (response *DisassociateListenerAdditionalCertificatesResponse) {
    response = &DisassociateListenerAdditionalCertificatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisassociateListenerAdditionalCertificates
// DisassociateListenerAdditionalCertificates is an async API. The system returns a request ID, but the additional cert is not yet unbound. The unbinding task is still in progress in the system backend. You can call the DescribeListenerCertificates API to query the cert unbinding status. If the cert is in Disassociating status, it is being unbound.
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
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_LISTENER = "ResourceNotFound.Listener"
//  RESOURCENOTFOUND_LOADBALANCER = "ResourceNotFound.LoadBalancer"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CERTIFICATE = "ResourceUnavailable.Certificate"
//  RESOURCEUNAVAILABLE_LISTENER = "ResourceUnavailable.Listener"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CERTIFICATENOTBOUND = "UnsupportedOperation.CertificateNotBound"
//  UNSUPPORTEDOPERATION_DISASSOCIATEDEFAULTCERTIFICATE = "UnsupportedOperation.DisassociateDefaultCertificate"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDPROTOCOL = "UnsupportedOperation.UnsupportedProtocol"
func (c *Client) DisassociateListenerAdditionalCertificates(request *DisassociateListenerAdditionalCertificatesRequest) (response *DisassociateListenerAdditionalCertificatesResponse, err error) {
    return c.DisassociateListenerAdditionalCertificatesWithContext(context.Background(), request)
}

// DisassociateListenerAdditionalCertificates
// DisassociateListenerAdditionalCertificates is an async API. The system returns a request ID, but the additional cert is not yet unbound. The unbinding task is still in progress in the system backend. You can call the DescribeListenerCertificates API to query the cert unbinding status. If the cert is in Disassociating status, it is being unbound.
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
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_LISTENER = "ResourceNotFound.Listener"
//  RESOURCENOTFOUND_LOADBALANCER = "ResourceNotFound.LoadBalancer"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CERTIFICATE = "ResourceUnavailable.Certificate"
//  RESOURCEUNAVAILABLE_LISTENER = "ResourceUnavailable.Listener"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CERTIFICATENOTBOUND = "UnsupportedOperation.CertificateNotBound"
//  UNSUPPORTEDOPERATION_DISASSOCIATEDEFAULTCERTIFICATE = "UnsupportedOperation.DisassociateDefaultCertificate"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDPROTOCOL = "UnsupportedOperation.UnsupportedProtocol"
func (c *Client) DisassociateListenerAdditionalCertificatesWithContext(ctx context.Context, request *DisassociateListenerAdditionalCertificatesRequest) (response *DisassociateListenerAdditionalCertificatesResponse, err error) {
    if request == nil {
        request = NewDisassociateListenerAdditionalCertificatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "DisassociateListenerAdditionalCertificates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisassociateListenerAdditionalCertificates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisassociateListenerAdditionalCertificatesResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceCreateLoadBalancerRequest() (request *InquirePriceCreateLoadBalancerRequest) {
    request = &InquirePriceCreateLoadBalancerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "InquirePriceCreateLoadBalancer")
    
    
    return
}

func NewInquirePriceCreateLoadBalancerResponse() (response *InquirePriceCreateLoadBalancerResponse) {
    response = &InquirePriceCreateLoadBalancerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquirePriceCreateLoadBalancer
// This API is used to query the price for creating a load balancer.
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
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_LISTENER = "ResourceNotFound.Listener"
//  RESOURCENOTFOUND_LOADBALANCER = "ResourceNotFound.LoadBalancer"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CERTIFICATE = "ResourceUnavailable.Certificate"
//  RESOURCEUNAVAILABLE_LISTENER = "ResourceUnavailable.Listener"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CERTIFICATENOTBOUND = "UnsupportedOperation.CertificateNotBound"
//  UNSUPPORTEDOPERATION_DISASSOCIATEDEFAULTCERTIFICATE = "UnsupportedOperation.DisassociateDefaultCertificate"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDPROTOCOL = "UnsupportedOperation.UnsupportedProtocol"
func (c *Client) InquirePriceCreateLoadBalancer(request *InquirePriceCreateLoadBalancerRequest) (response *InquirePriceCreateLoadBalancerResponse, err error) {
    return c.InquirePriceCreateLoadBalancerWithContext(context.Background(), request)
}

// InquirePriceCreateLoadBalancer
// This API is used to query the price for creating a load balancer.
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
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_LISTENER = "ResourceNotFound.Listener"
//  RESOURCENOTFOUND_LOADBALANCER = "ResourceNotFound.LoadBalancer"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CERTIFICATE = "ResourceUnavailable.Certificate"
//  RESOURCEUNAVAILABLE_LISTENER = "ResourceUnavailable.Listener"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CERTIFICATENOTBOUND = "UnsupportedOperation.CertificateNotBound"
//  UNSUPPORTEDOPERATION_DISASSOCIATEDEFAULTCERTIFICATE = "UnsupportedOperation.DisassociateDefaultCertificate"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDPROTOCOL = "UnsupportedOperation.UnsupportedProtocol"
func (c *Client) InquirePriceCreateLoadBalancerWithContext(ctx context.Context, request *InquirePriceCreateLoadBalancerRequest) (response *InquirePriceCreateLoadBalancerResponse, err error) {
    if request == nil {
        request = NewInquirePriceCreateLoadBalancerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "InquirePriceCreateLoadBalancer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquirePriceCreateLoadBalancer require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquirePriceCreateLoadBalancerResponse()
    err = c.Send(request, response)
    return
}

func NewModifyHealthCheckTemplateRequest() (request *ModifyHealthCheckTemplateRequest) {
    request = &ModifyHealthCheckTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "ModifyHealthCheckTemplate")
    
    
    return
}

func NewModifyHealthCheckTemplateResponse() (response *ModifyHealthCheckTemplateResponse) {
    response = &ModifyHealthCheckTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyHealthCheckTemplate
// Modify a health check template
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyHealthCheckTemplate(request *ModifyHealthCheckTemplateRequest) (response *ModifyHealthCheckTemplateResponse, err error) {
    return c.ModifyHealthCheckTemplateWithContext(context.Background(), request)
}

// ModifyHealthCheckTemplate
// Modify a health check template
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyHealthCheckTemplateWithContext(ctx context.Context, request *ModifyHealthCheckTemplateRequest) (response *ModifyHealthCheckTemplateResponse, err error) {
    if request == nil {
        request = NewModifyHealthCheckTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "ModifyHealthCheckTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyHealthCheckTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyHealthCheckTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyListenerAttributesRequest() (request *ModifyListenerAttributesRequest) {
    request = &ModifyListenerAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "ModifyListenerAttributes")
    
    
    return
}

func NewModifyListenerAttributesResponse() (response *ModifyListenerAttributesResponse) {
    response = &ModifyListenerAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyListenerAttributes
// Modifies listener properties.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyListenerAttributes(request *ModifyListenerAttributesRequest) (response *ModifyListenerAttributesResponse, err error) {
    return c.ModifyListenerAttributesWithContext(context.Background(), request)
}

// ModifyListenerAttributes
// Modifies listener properties.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyListenerAttributesWithContext(ctx context.Context, request *ModifyListenerAttributesRequest) (response *ModifyListenerAttributesResponse, err error) {
    if request == nil {
        request = NewModifyListenerAttributesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "ModifyListenerAttributes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyListenerAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyListenerAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLoadBalancerAddressTypeRequest() (request *ModifyLoadBalancerAddressTypeRequest) {
    request = &ModifyLoadBalancerAddressTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "ModifyLoadBalancerAddressType")
    
    
    return
}

func NewModifyLoadBalancerAddressTypeResponse() (response *ModifyLoadBalancerAddressTypeResponse) {
    response = &ModifyLoadBalancerAddressTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLoadBalancerAddressType
// **Prerequisite:**
//
// You have created an application CLB instance. For detailed operations, please see CreateLoadBalancer.
//
// When you need to change the network type of an application CLB instance from private network to public network through this API, you need to create an Elastic IP first.
//
// **Instructions:**
//
// The ModifyLoadBalancerAddressType API is an async API. The system returns a request ID, but the network type of the application CLB instance has not been changed yet. The change task is still in progress in the system backend. You can call DescribeLoadBalancerDetail to query the change status of the network type of the application CLB instance.
//
// When an application CLB instance is in the Configuring status, it means the network type of the instance is changing.
//
// When an application CLB instance is in the Active status, the network type change of the instance is successful.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyLoadBalancerAddressType(request *ModifyLoadBalancerAddressTypeRequest) (response *ModifyLoadBalancerAddressTypeResponse, err error) {
    return c.ModifyLoadBalancerAddressTypeWithContext(context.Background(), request)
}

// ModifyLoadBalancerAddressType
// **Prerequisite:**
//
// You have created an application CLB instance. For detailed operations, please see CreateLoadBalancer.
//
// When you need to change the network type of an application CLB instance from private network to public network through this API, you need to create an Elastic IP first.
//
// **Instructions:**
//
// The ModifyLoadBalancerAddressType API is an async API. The system returns a request ID, but the network type of the application CLB instance has not been changed yet. The change task is still in progress in the system backend. You can call DescribeLoadBalancerDetail to query the change status of the network type of the application CLB instance.
//
// When an application CLB instance is in the Configuring status, it means the network type of the instance is changing.
//
// When an application CLB instance is in the Active status, the network type change of the instance is successful.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyLoadBalancerAddressTypeWithContext(ctx context.Context, request *ModifyLoadBalancerAddressTypeRequest) (response *ModifyLoadBalancerAddressTypeResponse, err error) {
    if request == nil {
        request = NewModifyLoadBalancerAddressTypeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "ModifyLoadBalancerAddressType")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLoadBalancerAddressType require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLoadBalancerAddressTypeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLoadBalancerAttributesRequest() (request *ModifyLoadBalancerAttributesRequest) {
    request = &ModifyLoadBalancerAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "ModifyLoadBalancerAttributes")
    
    
    return
}

func NewModifyLoadBalancerAttributesResponse() (response *ModifyLoadBalancerAttributesResponse) {
    response = &ModifyLoadBalancerAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLoadBalancerAttributes
// The **ModifyLoadBalancerAttributes** API is an async API. It returns a request ID, but the application CLB instance attribute has not been modified yet. The modifying task is still in progress in the system backend. You can call [DescribeLoadBalancerDetail](https://www.tencentcloud.com/document/product/1311/84267) to query the modification status of the application CLB instance attribute.
//
// -When the application CLB instance attribute is in the **Configuring** status, it means the application CLB instance attribute is being modified.
//
// - When the application CLB instance attribute is in the **Active** status, it means the application CLB instance attribute was modified successfully.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyLoadBalancerAttributes(request *ModifyLoadBalancerAttributesRequest) (response *ModifyLoadBalancerAttributesResponse, err error) {
    return c.ModifyLoadBalancerAttributesWithContext(context.Background(), request)
}

// ModifyLoadBalancerAttributes
// The **ModifyLoadBalancerAttributes** API is an async API. It returns a request ID, but the application CLB instance attribute has not been modified yet. The modifying task is still in progress in the system backend. You can call [DescribeLoadBalancerDetail](https://www.tencentcloud.com/document/product/1311/84267) to query the modification status of the application CLB instance attribute.
//
// -When the application CLB instance attribute is in the **Configuring** status, it means the application CLB instance attribute is being modified.
//
// - When the application CLB instance attribute is in the **Active** status, it means the application CLB instance attribute was modified successfully.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyLoadBalancerAttributesWithContext(ctx context.Context, request *ModifyLoadBalancerAttributesRequest) (response *ModifyLoadBalancerAttributesResponse, err error) {
    if request == nil {
        request = NewModifyLoadBalancerAttributesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "ModifyLoadBalancerAttributes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLoadBalancerAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLoadBalancerAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLoadBalancerModificationProtectionRequest() (request *ModifyLoadBalancerModificationProtectionRequest) {
    request = &ModifyLoadBalancerModificationProtectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "ModifyLoadBalancerModificationProtection")
    
    
    return
}

func NewModifyLoadBalancerModificationProtectionResponse() (response *ModifyLoadBalancerModificationProtectionResponse) {
    response = &ModifyLoadBalancerModificationProtectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLoadBalancerModificationProtection
// Set load balancing instance modification protection.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyLoadBalancerModificationProtection(request *ModifyLoadBalancerModificationProtectionRequest) (response *ModifyLoadBalancerModificationProtectionResponse, err error) {
    return c.ModifyLoadBalancerModificationProtectionWithContext(context.Background(), request)
}

// ModifyLoadBalancerModificationProtection
// Set load balancing instance modification protection.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyLoadBalancerModificationProtectionWithContext(ctx context.Context, request *ModifyLoadBalancerModificationProtectionRequest) (response *ModifyLoadBalancerModificationProtectionResponse, err error) {
    if request == nil {
        request = NewModifyLoadBalancerModificationProtectionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "ModifyLoadBalancerModificationProtection")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLoadBalancerModificationProtection require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLoadBalancerModificationProtectionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRulesAttributesRequest() (request *ModifyRulesAttributesRequest) {
    request = &ModifyRulesAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "ModifyRulesAttributes")
    
    
    return
}

func NewModifyRulesAttributesResponse() (response *ModifyRulesAttributesResponse) {
    response = &ModifyRulesAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRulesAttributes
// This API is used to modify forwarding rule attributes. This is an async API. After the API return succeeds, you can call the DescribeAsyncJobs API with the returned RequestID as an input parameter to check whether this task is successful.
//
// A rule supports up to 10 forward Conditions and 5 forward Actions.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCLIENTTOKEN = "InvalidParameter.InvalidClientToken"
//  INVALIDPARAMETER_INVALIDFIELDTYPE = "InvalidParameter.InvalidFieldType"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETER_INVALIDQUOTACHECKREQUEST = "InvalidParameter.InvalidQuotaCheckRequest"
//  INVALIDPARAMETER_MISSINGACTION = "InvalidParameter.MissingAction"
//  INVALIDPARAMETER_MISSINGFIELD = "InvalidParameter.MissingField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_LISTENER = "ResourceNotFound.Listener"
//  RESOURCENOTFOUND_LOADBALANCER = "ResourceNotFound.LoadBalancer"
//  RESOURCENOTFOUND_RULE = "ResourceNotFound.Rule"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_LISTENER = "ResourceUnavailable.Listener"
//  RESOURCEUNAVAILABLE_LOADBALANCER = "ResourceUnavailable.LoadBalancer"
//  RESOURCEUNAVAILABLE_RULE = "ResourceUnavailable.Rule"
//  RESOURCEUNAVAILABLE_TARGETGROUP = "ResourceUnavailable.TargetGroup"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyRulesAttributes(request *ModifyRulesAttributesRequest) (response *ModifyRulesAttributesResponse, err error) {
    return c.ModifyRulesAttributesWithContext(context.Background(), request)
}

// ModifyRulesAttributes
// This API is used to modify forwarding rule attributes. This is an async API. After the API return succeeds, you can call the DescribeAsyncJobs API with the returned RequestID as an input parameter to check whether this task is successful.
//
// A rule supports up to 10 forward Conditions and 5 forward Actions.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCLIENTTOKEN = "InvalidParameter.InvalidClientToken"
//  INVALIDPARAMETER_INVALIDFIELDTYPE = "InvalidParameter.InvalidFieldType"
//  INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"
//  INVALIDPARAMETER_INVALIDQUOTACHECKREQUEST = "InvalidParameter.InvalidQuotaCheckRequest"
//  INVALIDPARAMETER_MISSINGACTION = "InvalidParameter.MissingAction"
//  INVALIDPARAMETER_MISSINGFIELD = "InvalidParameter.MissingField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_LISTENER = "ResourceNotFound.Listener"
//  RESOURCENOTFOUND_LOADBALANCER = "ResourceNotFound.LoadBalancer"
//  RESOURCENOTFOUND_RULE = "ResourceNotFound.Rule"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_LISTENER = "ResourceUnavailable.Listener"
//  RESOURCEUNAVAILABLE_LOADBALANCER = "ResourceUnavailable.LoadBalancer"
//  RESOURCEUNAVAILABLE_RULE = "ResourceUnavailable.Rule"
//  RESOURCEUNAVAILABLE_TARGETGROUP = "ResourceUnavailable.TargetGroup"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyRulesAttributesWithContext(ctx context.Context, request *ModifyRulesAttributesRequest) (response *ModifyRulesAttributesResponse, err error) {
    if request == nil {
        request = NewModifyRulesAttributesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "ModifyRulesAttributes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRulesAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRulesAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewModifySecurityPolicyAttributesRequest() (request *ModifySecurityPolicyAttributesRequest) {
    request = &ModifySecurityPolicyAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "ModifySecurityPolicyAttributes")
    
    
    return
}

func NewModifySecurityPolicyAttributesResponse() (response *ModifySecurityPolicyAttributesResponse) {
    response = &ModifySecurityPolicyAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySecurityPolicyAttributes
// Modify the properties of a custom security policy, including the policy name, TLS protocol version, and encryption suite. The modified configuration will be applied to all HTTPS listeners associated with this policy immediately.
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
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifySecurityPolicyAttributes(request *ModifySecurityPolicyAttributesRequest) (response *ModifySecurityPolicyAttributesResponse, err error) {
    return c.ModifySecurityPolicyAttributesWithContext(context.Background(), request)
}

// ModifySecurityPolicyAttributes
// Modify the properties of a custom security policy, including the policy name, TLS protocol version, and encryption suite. The modified configuration will be applied to all HTTPS listeners associated with this policy immediately.
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
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifySecurityPolicyAttributesWithContext(ctx context.Context, request *ModifySecurityPolicyAttributesRequest) (response *ModifySecurityPolicyAttributesResponse, err error) {
    if request == nil {
        request = NewModifySecurityPolicyAttributesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "ModifySecurityPolicyAttributes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySecurityPolicyAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySecurityPolicyAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTargetGroupAttributesRequest() (request *ModifyTargetGroupAttributesRequest) {
    request = &ModifyTargetGroupAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "ModifyTargetGroupAttributes")
    
    
    return
}

func NewModifyTargetGroupAttributesResponse() (response *ModifyTargetGroupAttributesResponse) {
    response = &ModifyTargetGroupAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTargetGroupAttributes
// Modify the target group.
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
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyTargetGroupAttributes(request *ModifyTargetGroupAttributesRequest) (response *ModifyTargetGroupAttributesResponse, err error) {
    return c.ModifyTargetGroupAttributesWithContext(context.Background(), request)
}

// ModifyTargetGroupAttributes
// Modify the target group.
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
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyTargetGroupAttributesWithContext(ctx context.Context, request *ModifyTargetGroupAttributesRequest) (response *ModifyTargetGroupAttributesResponse, err error) {
    if request == nil {
        request = NewModifyTargetGroupAttributesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "ModifyTargetGroupAttributes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTargetGroupAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTargetGroupAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTargetsInTargetGroupRequest() (request *ModifyTargetsInTargetGroupRequest) {
    request = &ModifyTargetsInTargetGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "ModifyTargetsInTargetGroup")
    
    
    return
}

func NewModifyTargetsInTargetGroupResponse() (response *ModifyTargetsInTargetGroupResponse) {
    response = &ModifyTargetsInTargetGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTargetsInTargetGroup
// Modifies backend service information in the target group.
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
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyTargetsInTargetGroup(request *ModifyTargetsInTargetGroupRequest) (response *ModifyTargetsInTargetGroupResponse, err error) {
    return c.ModifyTargetsInTargetGroupWithContext(context.Background(), request)
}

// ModifyTargetsInTargetGroup
// Modifies backend service information in the target group.
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
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyTargetsInTargetGroupWithContext(ctx context.Context, request *ModifyTargetsInTargetGroupRequest) (response *ModifyTargetsInTargetGroupResponse, err error) {
    if request == nil {
        request = NewModifyTargetsInTargetGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "ModifyTargetsInTargetGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTargetsInTargetGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTargetsInTargetGroupResponse()
    err = c.Send(request, response)
    return
}

func NewNotifyUnbindTargetRequest() (request *NotifyUnbindTargetRequest) {
    request = &NotifyUnbindTargetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "NotifyUnbindTarget")
    
    
    return
}

func NewNotifyUnbindTargetResponse() (response *NotifyUnbindTargetResponse) {
    response = &NotifyUnbindTargetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// NotifyUnbindTarget
// Notify load balancing to unbind real servers
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
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) NotifyUnbindTarget(request *NotifyUnbindTargetRequest) (response *NotifyUnbindTargetResponse, err error) {
    return c.NotifyUnbindTargetWithContext(context.Background(), request)
}

// NotifyUnbindTarget
// Notify load balancing to unbind real servers
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
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) NotifyUnbindTargetWithContext(ctx context.Context, request *NotifyUnbindTargetRequest) (response *NotifyUnbindTargetResponse, err error) {
    if request == nil {
        request = NewNotifyUnbindTargetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "NotifyUnbindTarget")
    
    if c.GetCredential() == nil {
        return nil, errors.New("NotifyUnbindTarget require credential")
    }

    request.SetContext(ctx)
    
    response = NewNotifyUnbindTargetResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveTargetsFromTargetGroupRequest() (request *RemoveTargetsFromTargetGroupRequest) {
    request = &RemoveTargetsFromTargetGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "RemoveTargetsFromTargetGroup")
    
    
    return
}

func NewRemoveTargetsFromTargetGroupResponse() (response *RemoveTargetsFromTargetGroupResponse) {
    response = &RemoveTargetsFromTargetGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveTargetsFromTargetGroup
// Removes a backend service from the target group
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
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RemoveTargetsFromTargetGroup(request *RemoveTargetsFromTargetGroupRequest) (response *RemoveTargetsFromTargetGroupResponse, err error) {
    return c.RemoveTargetsFromTargetGroupWithContext(context.Background(), request)
}

// RemoveTargetsFromTargetGroup
// Removes a backend service from the target group
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
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RemoveTargetsFromTargetGroupWithContext(ctx context.Context, request *RemoveTargetsFromTargetGroupRequest) (response *RemoveTargetsFromTargetGroupResponse, err error) {
    if request == nil {
        request = NewRemoveTargetsFromTargetGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "RemoveTargetsFromTargetGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveTargetsFromTargetGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveTargetsFromTargetGroupResponse()
    err = c.Send(request, response)
    return
}

func NewSetLoadBalancerSecurityGroupsRequest() (request *SetLoadBalancerSecurityGroupsRequest) {
    request = &SetLoadBalancerSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("alb", APIVersion, "SetLoadBalancerSecurityGroups")
    
    
    return
}

func NewSetLoadBalancerSecurityGroupsResponse() (response *SetLoadBalancerSecurityGroupsResponse) {
    response = &SetLoadBalancerSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetLoadBalancerSecurityGroups
// The SetLoadBalancerSecurityGroups API supports setting (binding and unbinding) security groups for a public network load balancing instance. To query the security groups currently bound to a load balancing instance, use the DescribeLoadBalancerDetail API (https://www.tencentcloud.com/document/api/1822/133711?from_cn_redirect=1). This API uses SET semantics.
//
// For the binding operation, input parameters need to be passed in for all security groups that should be bound to the load balancing instance (bound + new binding).
//
// During unbinding, input parameters need to pass in all security groups bound to a CLB instance after unbinding. To unbind all security groups, omit this parameter or specify an empty array.
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
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SetLoadBalancerSecurityGroups(request *SetLoadBalancerSecurityGroupsRequest) (response *SetLoadBalancerSecurityGroupsResponse, err error) {
    return c.SetLoadBalancerSecurityGroupsWithContext(context.Background(), request)
}

// SetLoadBalancerSecurityGroups
// The SetLoadBalancerSecurityGroups API supports setting (binding and unbinding) security groups for a public network load balancing instance. To query the security groups currently bound to a load balancing instance, use the DescribeLoadBalancerDetail API (https://www.tencentcloud.com/document/api/1822/133711?from_cn_redirect=1). This API uses SET semantics.
//
// For the binding operation, input parameters need to be passed in for all security groups that should be bound to the load balancing instance (bound + new binding).
//
// During unbinding, input parameters need to pass in all security groups bound to a CLB instance after unbinding. To unbind all security groups, omit this parameter or specify an empty array.
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
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SetLoadBalancerSecurityGroupsWithContext(ctx context.Context, request *SetLoadBalancerSecurityGroupsRequest) (response *SetLoadBalancerSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewSetLoadBalancerSecurityGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "alb", APIVersion, "SetLoadBalancerSecurityGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetLoadBalancerSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetLoadBalancerSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}
