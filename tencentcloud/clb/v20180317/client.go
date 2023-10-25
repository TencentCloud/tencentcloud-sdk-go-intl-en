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

package v20180317

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2018-03-17"

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


func NewAssociateTargetGroupsRequest() (request *AssociateTargetGroupsRequest) {
    request = &AssociateTargetGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "AssociateTargetGroups")
    
    
    return
}

func NewAssociateTargetGroupsResponse() (response *AssociateTargetGroupsResponse) {
    response = &AssociateTargetGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AssociateTargetGroups
// This API is used to bind target groups to CLB listeners (layer-4 protocol) or forwarding rules (layer-7 protocol).
//
// This is an async API. After it is returned successfully, you can call the `DescribeTaskStatus` API with the returned `RequestID` as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AssociateTargetGroups(request *AssociateTargetGroupsRequest) (response *AssociateTargetGroupsResponse, err error) {
    return c.AssociateTargetGroupsWithContext(context.Background(), request)
}

// AssociateTargetGroups
// This API is used to bind target groups to CLB listeners (layer-4 protocol) or forwarding rules (layer-7 protocol).
//
// This is an async API. After it is returned successfully, you can call the `DescribeTaskStatus` API with the returned `RequestID` as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AssociateTargetGroupsWithContext(ctx context.Context, request *AssociateTargetGroupsRequest) (response *AssociateTargetGroupsResponse, err error) {
    if request == nil {
        request = NewAssociateTargetGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssociateTargetGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewAssociateTargetGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewAutoRewriteRequest() (request *AutoRewriteRequest) {
    request = &AutoRewriteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "AutoRewrite")
    
    
    return
}

func NewAutoRewriteResponse() (response *AutoRewriteResponse) {
    response = &AutoRewriteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AutoRewrite
// An HTTPS:443 listener needs to be created first, along with a forwarding rule. When this API is called, an HTTP:80 listener will be created automatically if it did not exist and a forwarding rule corresponding to `Domains` (specified in the input parameter) under the HTTPS:443 listener will also be created. After successful creation, access requests to an HTTP:80 address will be redirected to an HTTPS:443 address automatically.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDLBSTATUS = "FailedOperation.InvalidLBStatus"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_LISTENERIDNOTFOUND = "InvalidParameter.ListenerIdNotFound"
//  INVALIDPARAMETER_LOCATIONNOTFOUND = "InvalidParameter.LocationNotFound"
//  INVALIDPARAMETER_PORTCHECKFAILED = "InvalidParameter.PortCheckFailed"
//  INVALIDPARAMETER_PROTOCOLCHECKFAILED = "InvalidParameter.ProtocolCheckFailed"
//  INVALIDPARAMETER_REWRITEALREADYEXIST = "InvalidParameter.RewriteAlreadyExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AutoRewrite(request *AutoRewriteRequest) (response *AutoRewriteResponse, err error) {
    return c.AutoRewriteWithContext(context.Background(), request)
}

// AutoRewrite
// An HTTPS:443 listener needs to be created first, along with a forwarding rule. When this API is called, an HTTP:80 listener will be created automatically if it did not exist and a forwarding rule corresponding to `Domains` (specified in the input parameter) under the HTTPS:443 listener will also be created. After successful creation, access requests to an HTTP:80 address will be redirected to an HTTPS:443 address automatically.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDLBSTATUS = "FailedOperation.InvalidLBStatus"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_LISTENERIDNOTFOUND = "InvalidParameter.ListenerIdNotFound"
//  INVALIDPARAMETER_LOCATIONNOTFOUND = "InvalidParameter.LocationNotFound"
//  INVALIDPARAMETER_PORTCHECKFAILED = "InvalidParameter.PortCheckFailed"
//  INVALIDPARAMETER_PROTOCOLCHECKFAILED = "InvalidParameter.ProtocolCheckFailed"
//  INVALIDPARAMETER_REWRITEALREADYEXIST = "InvalidParameter.RewriteAlreadyExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AutoRewriteWithContext(ctx context.Context, request *AutoRewriteRequest) (response *AutoRewriteResponse, err error) {
    if request == nil {
        request = NewAutoRewriteRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AutoRewrite require credential")
    }

    request.SetContext(ctx)
    
    response = NewAutoRewriteResponse()
    err = c.Send(request, response)
    return
}

func NewBatchDeregisterTargetsRequest() (request *BatchDeregisterTargetsRequest) {
    request = &BatchDeregisterTargetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "BatchDeregisterTargets")
    
    
    return
}

func NewBatchDeregisterTargetsResponse() (response *BatchDeregisterTargetsResponse) {
    response = &BatchDeregisterTargetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchDeregisterTargets
// This API is used to batch unbind real servers of the layer-4 and layer-7 VPC-based CLBs. Up to 500 real servers can be unbound in a batch.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_LISTENERIDNOTFOUND = "InvalidParameter.ListenerIdNotFound"
//  INVALIDPARAMETER_LOCATIONNOTFOUND = "InvalidParameter.LocationNotFound"
//  INVALIDPARAMETER_PORTCHECKFAILED = "InvalidParameter.PortCheckFailed"
//  INVALIDPARAMETER_PROTOCOLCHECKFAILED = "InvalidParameter.ProtocolCheckFailed"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) BatchDeregisterTargets(request *BatchDeregisterTargetsRequest) (response *BatchDeregisterTargetsResponse, err error) {
    return c.BatchDeregisterTargetsWithContext(context.Background(), request)
}

// BatchDeregisterTargets
// This API is used to batch unbind real servers of the layer-4 and layer-7 VPC-based CLBs. Up to 500 real servers can be unbound in a batch.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_LISTENERIDNOTFOUND = "InvalidParameter.ListenerIdNotFound"
//  INVALIDPARAMETER_LOCATIONNOTFOUND = "InvalidParameter.LocationNotFound"
//  INVALIDPARAMETER_PORTCHECKFAILED = "InvalidParameter.PortCheckFailed"
//  INVALIDPARAMETER_PROTOCOLCHECKFAILED = "InvalidParameter.ProtocolCheckFailed"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) BatchDeregisterTargetsWithContext(ctx context.Context, request *BatchDeregisterTargetsRequest) (response *BatchDeregisterTargetsResponse, err error) {
    if request == nil {
        request = NewBatchDeregisterTargetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchDeregisterTargets require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchDeregisterTargetsResponse()
    err = c.Send(request, response)
    return
}

func NewBatchModifyTargetWeightRequest() (request *BatchModifyTargetWeightRequest) {
    request = &BatchModifyTargetWeightRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "BatchModifyTargetWeight")
    
    
    return
}

func NewBatchModifyTargetWeightResponse() (response *BatchModifyTargetWeightResponse) {
    response = &BatchModifyTargetWeightResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchModifyTargetWeight
// This API is used to modify forwarding weights of real servers bound to CLB listeners in batches. Up to 500 servers can be unbound in a batch. As this API is async, you should check whether the task is successful by passing the RequestId returned to the API call `DescribeTaskStatus`.<br/> This API is supported by CLB layer-4 and layer-7 listeners, but not Classis CLB counterparts.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) BatchModifyTargetWeight(request *BatchModifyTargetWeightRequest) (response *BatchModifyTargetWeightResponse, err error) {
    return c.BatchModifyTargetWeightWithContext(context.Background(), request)
}

// BatchModifyTargetWeight
// This API is used to modify forwarding weights of real servers bound to CLB listeners in batches. Up to 500 servers can be unbound in a batch. As this API is async, you should check whether the task is successful by passing the RequestId returned to the API call `DescribeTaskStatus`.<br/> This API is supported by CLB layer-4 and layer-7 listeners, but not Classis CLB counterparts.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) BatchModifyTargetWeightWithContext(ctx context.Context, request *BatchModifyTargetWeightRequest) (response *BatchModifyTargetWeightResponse, err error) {
    if request == nil {
        request = NewBatchModifyTargetWeightRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchModifyTargetWeight require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchModifyTargetWeightResponse()
    err = c.Send(request, response)
    return
}

func NewBatchRegisterTargetsRequest() (request *BatchRegisterTargetsRequest) {
    request = &BatchRegisterTargetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "BatchRegisterTargets")
    
    
    return
}

func NewBatchRegisterTargetsResponse() (response *BatchRegisterTargetsResponse) {
    response = &BatchRegisterTargetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchRegisterTargets
// This API is used to batch bind CVM instances or ENIs. Up to 500 servers can be bound in a batch. It supports cross-region binding, layer-4 and layer-7 (TCP/UDP/HTTP/HTTPS) protocols, and VPC-based CLBs only.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_LISTENERIDNOTFOUND = "InvalidParameter.ListenerIdNotFound"
//  INVALIDPARAMETER_LOCATIONNOTFOUND = "InvalidParameter.LocationNotFound"
//  INVALIDPARAMETER_PORTCHECKFAILED = "InvalidParameter.PortCheckFailed"
//  INVALIDPARAMETER_PROTOCOLCHECKFAILED = "InvalidParameter.ProtocolCheckFailed"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) BatchRegisterTargets(request *BatchRegisterTargetsRequest) (response *BatchRegisterTargetsResponse, err error) {
    return c.BatchRegisterTargetsWithContext(context.Background(), request)
}

// BatchRegisterTargets
// This API is used to batch bind CVM instances or ENIs. Up to 500 servers can be bound in a batch. It supports cross-region binding, layer-4 and layer-7 (TCP/UDP/HTTP/HTTPS) protocols, and VPC-based CLBs only.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_LISTENERIDNOTFOUND = "InvalidParameter.ListenerIdNotFound"
//  INVALIDPARAMETER_LOCATIONNOTFOUND = "InvalidParameter.LocationNotFound"
//  INVALIDPARAMETER_PORTCHECKFAILED = "InvalidParameter.PortCheckFailed"
//  INVALIDPARAMETER_PROTOCOLCHECKFAILED = "InvalidParameter.ProtocolCheckFailed"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) BatchRegisterTargetsWithContext(ctx context.Context, request *BatchRegisterTargetsRequest) (response *BatchRegisterTargetsResponse, err error) {
    if request == nil {
        request = NewBatchRegisterTargetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchRegisterTargets require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchRegisterTargetsResponse()
    err = c.Send(request, response)
    return
}

func NewCloneLoadBalancerRequest() (request *CloneLoadBalancerRequest) {
    request = &CloneLoadBalancerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "CloneLoadBalancer")
    
    
    return
}

func NewCloneLoadBalancerResponse() (response *CloneLoadBalancerResponse) {
    response = &CloneLoadBalancerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CloneLoadBalancer
// This API is used to create a clone of the source CLB instance with the same forwarding rules and binding relations. Note that this API is asynchronous, which means that changes to the source CLB after invocation of the API are not included in the clone.
//
// 
//
// Limits:
//
// Instance attribute restrictions
//
//   Only pay-as-you-go instances can be cloned. Monthly-subscribed instances cannot be cloned. 
//
//   CLB instances without any billable items cannot be cloned.
//
//   Classic CLB instances and CLB u200dinstances created for Anti-DDoS service cannot be cloned.
//
//   Classic network-based instances cannot be cloned.
//
//   IPv6 instances, IPv6 NAT64 instances, and instances bound with both IPv4 and IPv6 cannot be cloned.
//
//   The following settings will not be cloned: **Custom configuration**, **Redirection configurations**, and **Allow Traffic by Default** in security groups.
//
//   Before cloning an instance, make sure all certificates used on the instance are valid. Cloning will fail if there are any expired certificates.
//
// Listener restrictions
//
//   Instances with QUIC listeners or port range listeners cannot be cloned.
//
//   Private network CLB instances with TCP_SSL listeners cannot be cloned.
//
//   Instances with layer-7 listeners that have no forwarding rules cannot be cloned.
//
//   u200dInstances u200dwith more than 50 listeners cannot be cloned. 
//
// Backend service restrictions
//
//   Instances with target groups and SCF cloud functions as the backend services cannot be cloned.
//
// 
//
// Notes:
//
// If you are using a BGP bandwidth package, you need to pass the package ID.
//
// To create a dedicated cluster-based CLB by cloning the source CLB, you need to pass the cluster ID. Otherwise, a normal CLB is created.
//
// This API is only available for beta users. To try it out, [submit a ticket](https://console.cloud.tencent.com/workorder/category?level1_id=6&level2_id=163&source=0&data_title=%E8%B4%9F%E8%BD%BD%E5%9D%87%E8%A1%A1%20CLB&step=1).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDLBSTATUS = "FailedOperation.InvalidLBStatus"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_LISTENERIDNOTFOUND = "InvalidParameter.ListenerIdNotFound"
//  INVALIDPARAMETER_LOCATIONNOTFOUND = "InvalidParameter.LocationNotFound"
//  INVALIDPARAMETER_PORTCHECKFAILED = "InvalidParameter.PortCheckFailed"
//  INVALIDPARAMETER_PROTOCOLCHECKFAILED = "InvalidParameter.ProtocolCheckFailed"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
func (c *Client) CloneLoadBalancer(request *CloneLoadBalancerRequest) (response *CloneLoadBalancerResponse, err error) {
    return c.CloneLoadBalancerWithContext(context.Background(), request)
}

// CloneLoadBalancer
// This API is used to create a clone of the source CLB instance with the same forwarding rules and binding relations. Note that this API is asynchronous, which means that changes to the source CLB after invocation of the API are not included in the clone.
//
// 
//
// Limits:
//
// Instance attribute restrictions
//
//   Only pay-as-you-go instances can be cloned. Monthly-subscribed instances cannot be cloned. 
//
//   CLB instances without any billable items cannot be cloned.
//
//   Classic CLB instances and CLB u200dinstances created for Anti-DDoS service cannot be cloned.
//
//   Classic network-based instances cannot be cloned.
//
//   IPv6 instances, IPv6 NAT64 instances, and instances bound with both IPv4 and IPv6 cannot be cloned.
//
//   The following settings will not be cloned: **Custom configuration**, **Redirection configurations**, and **Allow Traffic by Default** in security groups.
//
//   Before cloning an instance, make sure all certificates used on the instance are valid. Cloning will fail if there are any expired certificates.
//
// Listener restrictions
//
//   Instances with QUIC listeners or port range listeners cannot be cloned.
//
//   Private network CLB instances with TCP_SSL listeners cannot be cloned.
//
//   Instances with layer-7 listeners that have no forwarding rules cannot be cloned.
//
//   u200dInstances u200dwith more than 50 listeners cannot be cloned. 
//
// Backend service restrictions
//
//   Instances with target groups and SCF cloud functions as the backend services cannot be cloned.
//
// 
//
// Notes:
//
// If you are using a BGP bandwidth package, you need to pass the package ID.
//
// To create a dedicated cluster-based CLB by cloning the source CLB, you need to pass the cluster ID. Otherwise, a normal CLB is created.
//
// This API is only available for beta users. To try it out, [submit a ticket](https://console.cloud.tencent.com/workorder/category?level1_id=6&level2_id=163&source=0&data_title=%E8%B4%9F%E8%BD%BD%E5%9D%87%E8%A1%A1%20CLB&step=1).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDLBSTATUS = "FailedOperation.InvalidLBStatus"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_LISTENERIDNOTFOUND = "InvalidParameter.ListenerIdNotFound"
//  INVALIDPARAMETER_LOCATIONNOTFOUND = "InvalidParameter.LocationNotFound"
//  INVALIDPARAMETER_PORTCHECKFAILED = "InvalidParameter.PortCheckFailed"
//  INVALIDPARAMETER_PROTOCOLCHECKFAILED = "InvalidParameter.ProtocolCheckFailed"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
func (c *Client) CloneLoadBalancerWithContext(ctx context.Context, request *CloneLoadBalancerRequest) (response *CloneLoadBalancerResponse, err error) {
    if request == nil {
        request = NewCloneLoadBalancerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloneLoadBalancer require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloneLoadBalancerResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClsLogSetRequest() (request *CreateClsLogSetRequest) {
    request = &CreateClsLogSetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "CreateClsLogSet")
    
    
    return
}

func NewCreateClsLogSetResponse() (response *CreateClsLogSetResponse) {
    response = &CreateClsLogSetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateClsLogSet
// This API is used to create a CLB exclusive logset for storing CLB logs.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateClsLogSet(request *CreateClsLogSetRequest) (response *CreateClsLogSetResponse, err error) {
    return c.CreateClsLogSetWithContext(context.Background(), request)
}

// CreateClsLogSet
// This API is used to create a CLB exclusive logset for storing CLB logs.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateClsLogSetWithContext(ctx context.Context, request *CreateClsLogSetRequest) (response *CreateClsLogSetResponse, err error) {
    if request == nil {
        request = NewCreateClsLogSetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateClsLogSet require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateClsLogSetResponse()
    err = c.Send(request, response)
    return
}

func NewCreateListenerRequest() (request *CreateListenerRequest) {
    request = &CreateListenerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "CreateListener")
    
    
    return
}

func NewCreateListenerResponse() (response *CreateListenerResponse) {
    response = &CreateListenerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateListener
// This API is used to create a listener for a CLB instance.
//
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestId as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_PORTCHECKFAILED = "InvalidParameter.PortCheckFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateListener(request *CreateListenerRequest) (response *CreateListenerResponse, err error) {
    return c.CreateListenerWithContext(context.Background(), request)
}

// CreateListener
// This API is used to create a listener for a CLB instance.
//
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestId as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_PORTCHECKFAILED = "InvalidParameter.PortCheckFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateListenerWithContext(ctx context.Context, request *CreateListenerRequest) (response *CreateListenerResponse, err error) {
    if request == nil {
        request = NewCreateListenerRequest()
    }
    
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
    
    request.Init().WithApiInfo("clb", APIVersion, "CreateLoadBalancer")
    
    
    return
}

func NewCreateLoadBalancerResponse() (response *CreateLoadBalancerResponse) {
    response = &CreateLoadBalancerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLoadBalancer
// This API (CreateLoadBalancer) is used to create a CLB instance. To use the CLB service, you first need to purchase one or more instances. After this API is called successfully, a unique instance ID will be returned. There are two types of instances: public network and private network. For more information, see the product types in the product documentation.
//
// Note: (1) To apply for a CLB instance in the specified AZ and cross-AZ disaster recovery, please [submit a ticket](https://console.cloud.tencent.com/workorder/category); (2) Currently, IPv6 is supported only in Beijing, Shanghai, and Guangzhou regions.
//
// This is an async API. After it is returned successfully, you can call the DescribeLoadBalancers API to query the status of the instance (such as creating and normal) to check whether it is successfully created.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateLoadBalancer(request *CreateLoadBalancerRequest) (response *CreateLoadBalancerResponse, err error) {
    return c.CreateLoadBalancerWithContext(context.Background(), request)
}

// CreateLoadBalancer
// This API (CreateLoadBalancer) is used to create a CLB instance. To use the CLB service, you first need to purchase one or more instances. After this API is called successfully, a unique instance ID will be returned. There are two types of instances: public network and private network. For more information, see the product types in the product documentation.
//
// Note: (1) To apply for a CLB instance in the specified AZ and cross-AZ disaster recovery, please [submit a ticket](https://console.cloud.tencent.com/workorder/category); (2) Currently, IPv6 is supported only in Beijing, Shanghai, and Guangzhou regions.
//
// This is an async API. After it is returned successfully, you can call the DescribeLoadBalancers API to query the status of the instance (such as creating and normal) to check whether it is successfully created.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateLoadBalancerWithContext(ctx context.Context, request *CreateLoadBalancerRequest) (response *CreateLoadBalancerResponse, err error) {
    if request == nil {
        request = NewCreateLoadBalancerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLoadBalancer require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLoadBalancerResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLoadBalancerSnatIpsRequest() (request *CreateLoadBalancerSnatIpsRequest) {
    request = &CreateLoadBalancerSnatIpsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "CreateLoadBalancerSnatIps")
    
    
    return
}

func NewCreateLoadBalancerSnatIpsResponse() (response *CreateLoadBalancerSnatIpsResponse) {
    response = &CreateLoadBalancerSnatIpsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLoadBalancerSnatIps
// This API is used to add an SNAT IP for an SnatPro CLB instance. If SnatPro is not enabled for CLB, it will be automatically enabled after the SNAT IP is added.
//
// This is an async API. After it is returned successfully, you can check the task result by calling `DescribeTaskStatus` with the returned `RequestID`.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateLoadBalancerSnatIps(request *CreateLoadBalancerSnatIpsRequest) (response *CreateLoadBalancerSnatIpsResponse, err error) {
    return c.CreateLoadBalancerSnatIpsWithContext(context.Background(), request)
}

// CreateLoadBalancerSnatIps
// This API is used to add an SNAT IP for an SnatPro CLB instance. If SnatPro is not enabled for CLB, it will be automatically enabled after the SNAT IP is added.
//
// This is an async API. After it is returned successfully, you can check the task result by calling `DescribeTaskStatus` with the returned `RequestID`.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateLoadBalancerSnatIpsWithContext(ctx context.Context, request *CreateLoadBalancerSnatIpsRequest) (response *CreateLoadBalancerSnatIpsResponse, err error) {
    if request == nil {
        request = NewCreateLoadBalancerSnatIpsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLoadBalancerSnatIps require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLoadBalancerSnatIpsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRuleRequest() (request *CreateRuleRequest) {
    request = &CreateRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "CreateRule")
    
    
    return
}

func NewCreateRuleResponse() (response *CreateRuleResponse) {
    response = &CreateRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRule
// This API (CreateRule) is used to create a forwarding rule under an existing layer-7 CLB listener, where real servers must be bound to the rule instead of the listener.
//
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestID as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateRule(request *CreateRuleRequest) (response *CreateRuleResponse, err error) {
    return c.CreateRuleWithContext(context.Background(), request)
}

// CreateRule
// This API (CreateRule) is used to create a forwarding rule under an existing layer-7 CLB listener, where real servers must be bound to the rule instead of the listener.
//
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestID as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateRuleWithContext(ctx context.Context, request *CreateRuleRequest) (response *CreateRuleResponse, err error) {
    if request == nil {
        request = NewCreateRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTargetGroupRequest() (request *CreateTargetGroupRequest) {
    request = &CreateTargetGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "CreateTargetGroup")
    
    
    return
}

func NewCreateTargetGroupResponse() (response *CreateTargetGroupResponse) {
    response = &CreateTargetGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTargetGroup
// This API is used to create a target group. This feature is in beta test, if you want to try it out, please [submit a ticket](https://console.cloud.tencent.com/workorder/category?level1_id=6&level2_id=163&source=0&data_title=%E8%B4%9F%E8%BD%BD%E5%9D%87%E8%A1%A1%20LB&step=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateTargetGroup(request *CreateTargetGroupRequest) (response *CreateTargetGroupResponse, err error) {
    return c.CreateTargetGroupWithContext(context.Background(), request)
}

// CreateTargetGroup
// This API is used to create a target group. This feature is in beta test, if you want to try it out, please [submit a ticket](https://console.cloud.tencent.com/workorder/category?level1_id=6&level2_id=163&source=0&data_title=%E8%B4%9F%E8%BD%BD%E5%9D%87%E8%A1%A1%20LB&step=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateTargetGroupWithContext(ctx context.Context, request *CreateTargetGroupRequest) (response *CreateTargetGroupResponse, err error) {
    if request == nil {
        request = NewCreateTargetGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTargetGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTargetGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTopicRequest() (request *CreateTopicRequest) {
    request = &CreateTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "CreateTopic")
    
    
    return
}

func NewCreateTopicResponse() (response *CreateTopicResponse) {
    response = &CreateTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTopic
// This API is used to create a topic with the full-text index and key-value index enabled by default. The creation will fail if there is no CLB exclusive logset.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateTopic(request *CreateTopicRequest) (response *CreateTopicResponse, err error) {
    return c.CreateTopicWithContext(context.Background(), request)
}

// CreateTopic
// This API is used to create a topic with the full-text index and key-value index enabled by default. The creation will fail if there is no CLB exclusive logset.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateTopicWithContext(ctx context.Context, request *CreateTopicRequest) (response *CreateTopicResponse, err error) {
    if request == nil {
        request = NewCreateTopicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteListenerRequest() (request *DeleteListenerRequest) {
    request = &DeleteListenerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DeleteListener")
    
    
    return
}

func NewDeleteListenerResponse() (response *DeleteListenerResponse) {
    response = &DeleteListenerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteListener
// This API is used to delete a listener from a CLB instance (layer-4 or layer-7).
//
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestID as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteListener(request *DeleteListenerRequest) (response *DeleteListenerResponse, err error) {
    return c.DeleteListenerWithContext(context.Background(), request)
}

// DeleteListener
// This API is used to delete a listener from a CLB instance (layer-4 or layer-7).
//
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestID as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteListenerWithContext(ctx context.Context, request *DeleteListenerRequest) (response *DeleteListenerResponse, err error) {
    if request == nil {
        request = NewDeleteListenerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteListener require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteListenerResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLoadBalancerRequest() (request *DeleteLoadBalancerRequest) {
    request = &DeleteLoadBalancerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DeleteLoadBalancer")
    
    
    return
}

func NewDeleteLoadBalancerResponse() (response *DeleteLoadBalancerResponse) {
    response = &DeleteLoadBalancerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLoadBalancer
// This API (DeleteLoadBalancer) is used to delete one or more specified CLB instances.
//
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestId as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteLoadBalancer(request *DeleteLoadBalancerRequest) (response *DeleteLoadBalancerResponse, err error) {
    return c.DeleteLoadBalancerWithContext(context.Background(), request)
}

// DeleteLoadBalancer
// This API (DeleteLoadBalancer) is used to delete one or more specified CLB instances.
//
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestId as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteLoadBalancerWithContext(ctx context.Context, request *DeleteLoadBalancerRequest) (response *DeleteLoadBalancerResponse, err error) {
    if request == nil {
        request = NewDeleteLoadBalancerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLoadBalancer require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLoadBalancerResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLoadBalancerListenersRequest() (request *DeleteLoadBalancerListenersRequest) {
    request = &DeleteLoadBalancerListenersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DeleteLoadBalancerListeners")
    
    
    return
}

func NewDeleteLoadBalancerListenersResponse() (response *DeleteLoadBalancerListenersResponse) {
    response = &DeleteLoadBalancerListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLoadBalancerListeners
// This API is used to delete multiple listeners of a CLB instance.
//
// This is an async API. After it is returned successfully, you can call the `DescribeTaskStatus` API with the returned `RequestID` as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteLoadBalancerListeners(request *DeleteLoadBalancerListenersRequest) (response *DeleteLoadBalancerListenersResponse, err error) {
    return c.DeleteLoadBalancerListenersWithContext(context.Background(), request)
}

// DeleteLoadBalancerListeners
// This API is used to delete multiple listeners of a CLB instance.
//
// This is an async API. After it is returned successfully, you can call the `DescribeTaskStatus` API with the returned `RequestID` as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteLoadBalancerListenersWithContext(ctx context.Context, request *DeleteLoadBalancerListenersRequest) (response *DeleteLoadBalancerListenersResponse, err error) {
    if request == nil {
        request = NewDeleteLoadBalancerListenersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLoadBalancerListeners require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLoadBalancerListenersResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLoadBalancerSnatIpsRequest() (request *DeleteLoadBalancerSnatIpsRequest) {
    request = &DeleteLoadBalancerSnatIpsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DeleteLoadBalancerSnatIps")
    
    
    return
}

func NewDeleteLoadBalancerSnatIpsResponse() (response *DeleteLoadBalancerSnatIpsResponse) {
    response = &DeleteLoadBalancerSnatIpsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLoadBalancerSnatIps
// This API is used to delete the SNAT IP for an SnatPro CLB instance.
//
// This is an async API. After it is returned successfully, you can check the task result by calling `DescribeTaskStatus` with the returned `RequestID`.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DeleteLoadBalancerSnatIps(request *DeleteLoadBalancerSnatIpsRequest) (response *DeleteLoadBalancerSnatIpsResponse, err error) {
    return c.DeleteLoadBalancerSnatIpsWithContext(context.Background(), request)
}

// DeleteLoadBalancerSnatIps
// This API is used to delete the SNAT IP for an SnatPro CLB instance.
//
// This is an async API. After it is returned successfully, you can check the task result by calling `DescribeTaskStatus` with the returned `RequestID`.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DeleteLoadBalancerSnatIpsWithContext(ctx context.Context, request *DeleteLoadBalancerSnatIpsRequest) (response *DeleteLoadBalancerSnatIpsResponse, err error) {
    if request == nil {
        request = NewDeleteLoadBalancerSnatIpsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLoadBalancerSnatIps require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLoadBalancerSnatIpsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRewriteRequest() (request *DeleteRewriteRequest) {
    request = &DeleteRewriteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DeleteRewrite")
    
    
    return
}

func NewDeleteRewriteResponse() (response *DeleteRewriteResponse) {
    response = &DeleteRewriteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRewrite
// This API (DeleteRewrite) is used to delete the redirection relationship between the specified forwarding rules.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDLBSTATUS = "FailedOperation.InvalidLBStatus"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_SOMEREWRITENOTFOUND = "InvalidParameter.SomeRewriteNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteRewrite(request *DeleteRewriteRequest) (response *DeleteRewriteResponse, err error) {
    return c.DeleteRewriteWithContext(context.Background(), request)
}

// DeleteRewrite
// This API (DeleteRewrite) is used to delete the redirection relationship between the specified forwarding rules.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDLBSTATUS = "FailedOperation.InvalidLBStatus"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_SOMEREWRITENOTFOUND = "InvalidParameter.SomeRewriteNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteRewriteWithContext(ctx context.Context, request *DeleteRewriteRequest) (response *DeleteRewriteResponse, err error) {
    if request == nil {
        request = NewDeleteRewriteRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRewrite require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRewriteResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRuleRequest() (request *DeleteRuleRequest) {
    request = &DeleteRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DeleteRule")
    
    
    return
}

func NewDeleteRuleResponse() (response *DeleteRuleResponse) {
    response = &DeleteRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRule
// This API (DeleteRule) is used to delete a forwarding rule under a layer-7 CLB instance listener
//
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestID as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteRule(request *DeleteRuleRequest) (response *DeleteRuleResponse, err error) {
    return c.DeleteRuleWithContext(context.Background(), request)
}

// DeleteRule
// This API (DeleteRule) is used to delete a forwarding rule under a layer-7 CLB instance listener
//
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestID as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteRuleWithContext(ctx context.Context, request *DeleteRuleRequest) (response *DeleteRuleResponse, err error) {
    if request == nil {
        request = NewDeleteRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTargetGroupsRequest() (request *DeleteTargetGroupsRequest) {
    request = &DeleteTargetGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DeleteTargetGroups")
    
    
    return
}

func NewDeleteTargetGroupsResponse() (response *DeleteTargetGroupsResponse) {
    response = &DeleteTargetGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTargetGroups
// This API is used to delete a target group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteTargetGroups(request *DeleteTargetGroupsRequest) (response *DeleteTargetGroupsResponse, err error) {
    return c.DeleteTargetGroupsWithContext(context.Background(), request)
}

// DeleteTargetGroups
// This API is used to delete a target group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteTargetGroupsWithContext(ctx context.Context, request *DeleteTargetGroupsRequest) (response *DeleteTargetGroupsResponse, err error) {
    if request == nil {
        request = NewDeleteTargetGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTargetGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTargetGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDeregisterFunctionTargetsRequest() (request *DeregisterFunctionTargetsRequest) {
    request = &DeregisterFunctionTargetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DeregisterFunctionTargets")
    
    
    return
}

func NewDeregisterFunctionTargetsResponse() (response *DeregisterFunctionTargetsResponse) {
    response = &DeregisterFunctionTargetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeregisterFunctionTargets
// This API is used to unbind a SCF function with a CLB forwarding rule. For L7 listeners, you need to specify the forwarding rule by using `LocationId` or `Domain+Url`. 
//
// This is an async API. After it is returned successfully, you can call the [DescribeTaskStatus](https://intl.cloud.tencent.com/document/product/214/30683?from_cn_redirect=1) API with the returned RequestID to check whether this task is successful.
//
// <br/>Limits: 
//
// 
//
// - Binding with SCF is only available in Guangzhou, Shenzhen Finance, Shanghai, Shanghai Finance, Beijing, Chengdu, Hong Kong (China), Singapore, Mumbai, Tokyo, and Silicon Valley.
//
// - SCF functions can only be bound with CLB instances of bill-by-IP accounts but not with bill-by-CVM accounts. If you are using a bill-by-CVM account, we recommend upgrading it to a bill-by-IP account. For more information, please see [Checking Account Type](https://intl.cloud.tencent.com/document/product/1199/49090?from_cn_redirect=1).
//
// - SCF functions cannot be bound with classic CLB instances.
//
// - SCF functions cannot be bound with classic network-based CLB instances.
//
// - SCF functions in the same region can be bound with CLB instances. SCF functions can only be bound across VPCs but not regions.
//
// - SCF functions can only be bound with IPv4 and IPv6 NAT64 CLB instances, but currently not with IPv6 CLB instances.
//
// - SCF functions can only be bound with layer-7 HTTP and HTTPS listeners, but not with layer-7 QUIC listeners or layer-4 (TCP, UDP, and TCP SSL) listeners.
//
// - Only SCF event-triggered functions can be bound with CLB instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeregisterFunctionTargets(request *DeregisterFunctionTargetsRequest) (response *DeregisterFunctionTargetsResponse, err error) {
    return c.DeregisterFunctionTargetsWithContext(context.Background(), request)
}

// DeregisterFunctionTargets
// This API is used to unbind a SCF function with a CLB forwarding rule. For L7 listeners, you need to specify the forwarding rule by using `LocationId` or `Domain+Url`. 
//
// This is an async API. After it is returned successfully, you can call the [DescribeTaskStatus](https://intl.cloud.tencent.com/document/product/214/30683?from_cn_redirect=1) API with the returned RequestID to check whether this task is successful.
//
// <br/>Limits: 
//
// 
//
// - Binding with SCF is only available in Guangzhou, Shenzhen Finance, Shanghai, Shanghai Finance, Beijing, Chengdu, Hong Kong (China), Singapore, Mumbai, Tokyo, and Silicon Valley.
//
// - SCF functions can only be bound with CLB instances of bill-by-IP accounts but not with bill-by-CVM accounts. If you are using a bill-by-CVM account, we recommend upgrading it to a bill-by-IP account. For more information, please see [Checking Account Type](https://intl.cloud.tencent.com/document/product/1199/49090?from_cn_redirect=1).
//
// - SCF functions cannot be bound with classic CLB instances.
//
// - SCF functions cannot be bound with classic network-based CLB instances.
//
// - SCF functions in the same region can be bound with CLB instances. SCF functions can only be bound across VPCs but not regions.
//
// - SCF functions can only be bound with IPv4 and IPv6 NAT64 CLB instances, but currently not with IPv6 CLB instances.
//
// - SCF functions can only be bound with layer-7 HTTP and HTTPS listeners, but not with layer-7 QUIC listeners or layer-4 (TCP, UDP, and TCP SSL) listeners.
//
// - Only SCF event-triggered functions can be bound with CLB instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeregisterFunctionTargetsWithContext(ctx context.Context, request *DeregisterFunctionTargetsRequest) (response *DeregisterFunctionTargetsResponse, err error) {
    if request == nil {
        request = NewDeregisterFunctionTargetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeregisterFunctionTargets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeregisterFunctionTargetsResponse()
    err = c.Send(request, response)
    return
}

func NewDeregisterTargetGroupInstancesRequest() (request *DeregisterTargetGroupInstancesRequest) {
    request = &DeregisterTargetGroupInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DeregisterTargetGroupInstances")
    
    
    return
}

func NewDeregisterTargetGroupInstancesResponse() (response *DeregisterTargetGroupInstancesResponse) {
    response = &DeregisterTargetGroupInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeregisterTargetGroupInstances
// This API is used to unbind a server from a target group.
//
// This is an async API. After it is returned successfully, you can call the API `DescribeTaskStatus` with the returned RequestId as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeregisterTargetGroupInstances(request *DeregisterTargetGroupInstancesRequest) (response *DeregisterTargetGroupInstancesResponse, err error) {
    return c.DeregisterTargetGroupInstancesWithContext(context.Background(), request)
}

// DeregisterTargetGroupInstances
// This API is used to unbind a server from a target group.
//
// This is an async API. After it is returned successfully, you can call the API `DescribeTaskStatus` with the returned RequestId as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeregisterTargetGroupInstancesWithContext(ctx context.Context, request *DeregisterTargetGroupInstancesRequest) (response *DeregisterTargetGroupInstancesResponse, err error) {
    if request == nil {
        request = NewDeregisterTargetGroupInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeregisterTargetGroupInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeregisterTargetGroupInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDeregisterTargetsRequest() (request *DeregisterTargetsRequest) {
    request = &DeregisterTargetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DeregisterTargets")
    
    
    return
}

func NewDeregisterTargetsResponse() (response *DeregisterTargetsResponse) {
    response = &DeregisterTargetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeregisterTargets
// This API (DeregisterTargets) is used to unbind one or more real servers from a CLB listener or forwarding rule. For layer-4 listeners, only the listener ID needs to be specified. For layer-7 listeners, the forwarding rule also needs to be specified through LocationId or Domain+Url.
//
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestID as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeregisterTargets(request *DeregisterTargetsRequest) (response *DeregisterTargetsResponse, err error) {
    return c.DeregisterTargetsWithContext(context.Background(), request)
}

// DeregisterTargets
// This API (DeregisterTargets) is used to unbind one or more real servers from a CLB listener or forwarding rule. For layer-4 listeners, only the listener ID needs to be specified. For layer-7 listeners, the forwarding rule also needs to be specified through LocationId or Domain+Url.
//
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestID as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeregisterTargetsWithContext(ctx context.Context, request *DeregisterTargetsRequest) (response *DeregisterTargetsResponse, err error) {
    if request == nil {
        request = NewDeregisterTargetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeregisterTargets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeregisterTargetsResponse()
    err = c.Send(request, response)
    return
}

func NewDeregisterTargetsFromClassicalLBRequest() (request *DeregisterTargetsFromClassicalLBRequest) {
    request = &DeregisterTargetsFromClassicalLBRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DeregisterTargetsFromClassicalLB")
    
    
    return
}

func NewDeregisterTargetsFromClassicalLBResponse() (response *DeregisterTargetsFromClassicalLBResponse) {
    response = &DeregisterTargetsFromClassicalLBResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeregisterTargetsFromClassicalLB
// This API is used to unbind a CLB real server. This is an async API. After it is returned successfully, you can call the API `DescribeTaskStatus` with the returned RequestId as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeregisterTargetsFromClassicalLB(request *DeregisterTargetsFromClassicalLBRequest) (response *DeregisterTargetsFromClassicalLBResponse, err error) {
    return c.DeregisterTargetsFromClassicalLBWithContext(context.Background(), request)
}

// DeregisterTargetsFromClassicalLB
// This API is used to unbind a CLB real server. This is an async API. After it is returned successfully, you can call the API `DescribeTaskStatus` with the returned RequestId as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeregisterTargetsFromClassicalLBWithContext(ctx context.Context, request *DeregisterTargetsFromClassicalLBRequest) (response *DeregisterTargetsFromClassicalLBResponse, err error) {
    if request == nil {
        request = NewDeregisterTargetsFromClassicalLBRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeregisterTargetsFromClassicalLB require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeregisterTargetsFromClassicalLBResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBlockIPListRequest() (request *DescribeBlockIPListRequest) {
    request = &DescribeBlockIPListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DescribeBlockIPList")
    
    
    return
}

func NewDescribeBlockIPListResponse() (response *DescribeBlockIPListResponse) {
    response = &DescribeBlockIPListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBlockIPList
// This API is used to query the list of blocked IPs (blocklist) of a CLB instance. (This API is in beta test. To use it, please submit a ticket.)
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeBlockIPList(request *DescribeBlockIPListRequest) (response *DescribeBlockIPListResponse, err error) {
    return c.DescribeBlockIPListWithContext(context.Background(), request)
}

// DescribeBlockIPList
// This API is used to query the list of blocked IPs (blocklist) of a CLB instance. (This API is in beta test. To use it, please submit a ticket.)
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeBlockIPListWithContext(ctx context.Context, request *DescribeBlockIPListRequest) (response *DescribeBlockIPListResponse, err error) {
    if request == nil {
        request = NewDescribeBlockIPListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBlockIPList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBlockIPListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBlockIPTaskRequest() (request *DescribeBlockIPTaskRequest) {
    request = &DescribeBlockIPTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DescribeBlockIPTask")
    
    
    return
}

func NewDescribeBlockIPTaskResponse() (response *DescribeBlockIPTaskResponse) {
    response = &DescribeBlockIPTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBlockIPTask
// This API is used to query the execution status of an async IP blocking (blocklisting) task by the async task ID returned by the `ModifyBlockIPList` API. (This API is in beta test. To use it, please submit a ticket.)
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBlockIPTask(request *DescribeBlockIPTaskRequest) (response *DescribeBlockIPTaskResponse, err error) {
    return c.DescribeBlockIPTaskWithContext(context.Background(), request)
}

// DescribeBlockIPTask
// This API is used to query the execution status of an async IP blocking (blocklisting) task by the async task ID returned by the `ModifyBlockIPList` API. (This API is in beta test. To use it, please submit a ticket.)
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBlockIPTaskWithContext(ctx context.Context, request *DescribeBlockIPTaskRequest) (response *DescribeBlockIPTaskResponse, err error) {
    if request == nil {
        request = NewDescribeBlockIPTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBlockIPTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBlockIPTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClassicalLBByInstanceIdRequest() (request *DescribeClassicalLBByInstanceIdRequest) {
    request = &DescribeClassicalLBByInstanceIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DescribeClassicalLBByInstanceId")
    
    
    return
}

func NewDescribeClassicalLBByInstanceIdResponse() (response *DescribeClassicalLBByInstanceIdResponse) {
    response = &DescribeClassicalLBByInstanceIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClassicalLBByInstanceId
// This API is used to get the list of classic CLB instance IDs through a real server ID.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeClassicalLBByInstanceId(request *DescribeClassicalLBByInstanceIdRequest) (response *DescribeClassicalLBByInstanceIdResponse, err error) {
    return c.DescribeClassicalLBByInstanceIdWithContext(context.Background(), request)
}

// DescribeClassicalLBByInstanceId
// This API is used to get the list of classic CLB instance IDs through a real server ID.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeClassicalLBByInstanceIdWithContext(ctx context.Context, request *DescribeClassicalLBByInstanceIdRequest) (response *DescribeClassicalLBByInstanceIdResponse, err error) {
    if request == nil {
        request = NewDescribeClassicalLBByInstanceIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClassicalLBByInstanceId require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClassicalLBByInstanceIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClassicalLBHealthStatusRequest() (request *DescribeClassicalLBHealthStatusRequest) {
    request = &DescribeClassicalLBHealthStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DescribeClassicalLBHealthStatus")
    
    
    return
}

func NewDescribeClassicalLBHealthStatusResponse() (response *DescribeClassicalLBHealthStatusResponse) {
    response = &DescribeClassicalLBHealthStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClassicalLBHealthStatus
// This API (DescribeClassicalLBHealthStatus) is used to get the real server health status of a classic CLB
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeClassicalLBHealthStatus(request *DescribeClassicalLBHealthStatusRequest) (response *DescribeClassicalLBHealthStatusResponse, err error) {
    return c.DescribeClassicalLBHealthStatusWithContext(context.Background(), request)
}

// DescribeClassicalLBHealthStatus
// This API (DescribeClassicalLBHealthStatus) is used to get the real server health status of a classic CLB
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeClassicalLBHealthStatusWithContext(ctx context.Context, request *DescribeClassicalLBHealthStatusRequest) (response *DescribeClassicalLBHealthStatusResponse, err error) {
    if request == nil {
        request = NewDescribeClassicalLBHealthStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClassicalLBHealthStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClassicalLBHealthStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClassicalLBListenersRequest() (request *DescribeClassicalLBListenersRequest) {
    request = &DescribeClassicalLBListenersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DescribeClassicalLBListeners")
    
    
    return
}

func NewDescribeClassicalLBListenersResponse() (response *DescribeClassicalLBListenersResponse) {
    response = &DescribeClassicalLBListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClassicalLBListeners
// This API (DescribeClassicalLBListeners) is used to get the listener information of a classic CLB.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClassicalLBListeners(request *DescribeClassicalLBListenersRequest) (response *DescribeClassicalLBListenersResponse, err error) {
    return c.DescribeClassicalLBListenersWithContext(context.Background(), request)
}

// DescribeClassicalLBListeners
// This API (DescribeClassicalLBListeners) is used to get the listener information of a classic CLB.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClassicalLBListenersWithContext(ctx context.Context, request *DescribeClassicalLBListenersRequest) (response *DescribeClassicalLBListenersResponse, err error) {
    if request == nil {
        request = NewDescribeClassicalLBListenersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClassicalLBListeners require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClassicalLBListenersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClassicalLBTargetsRequest() (request *DescribeClassicalLBTargetsRequest) {
    request = &DescribeClassicalLBTargetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DescribeClassicalLBTargets")
    
    
    return
}

func NewDescribeClassicalLBTargetsResponse() (response *DescribeClassicalLBTargetsResponse) {
    response = &DescribeClassicalLBTargetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClassicalLBTargets
// This API is used to get the real servers bound to a classic CLB instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeClassicalLBTargets(request *DescribeClassicalLBTargetsRequest) (response *DescribeClassicalLBTargetsResponse, err error) {
    return c.DescribeClassicalLBTargetsWithContext(context.Background(), request)
}

// DescribeClassicalLBTargets
// This API is used to get the real servers bound to a classic CLB instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeClassicalLBTargetsWithContext(ctx context.Context, request *DescribeClassicalLBTargetsRequest) (response *DescribeClassicalLBTargetsResponse, err error) {
    if request == nil {
        request = NewDescribeClassicalLBTargetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClassicalLBTargets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClassicalLBTargetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClsLogSetRequest() (request *DescribeClsLogSetRequest) {
    request = &DescribeClsLogSetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DescribeClsLogSet")
    
    
    return
}

func NewDescribeClsLogSetResponse() (response *DescribeClsLogSetResponse) {
    response = &DescribeClsLogSetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClsLogSet
// This API is used to get the CLB exclusive logset.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeClsLogSet(request *DescribeClsLogSetRequest) (response *DescribeClsLogSetResponse, err error) {
    return c.DescribeClsLogSetWithContext(context.Background(), request)
}

// DescribeClsLogSet
// This API is used to get the CLB exclusive logset.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeClsLogSetWithContext(ctx context.Context, request *DescribeClsLogSetRequest) (response *DescribeClsLogSetResponse, err error) {
    if request == nil {
        request = NewDescribeClsLogSetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClsLogSet require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClsLogSetResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCrossTargetsRequest() (request *DescribeCrossTargetsRequest) {
    request = &DescribeCrossTargetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DescribeCrossTargets")
    
    
    return
}

func NewDescribeCrossTargetsResponse() (response *DescribeCrossTargetsResponse) {
    response = &DescribeCrossTargetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCrossTargets
// Queries information of CVMs and ENIs that use cross-region binding 2.0
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeCrossTargets(request *DescribeCrossTargetsRequest) (response *DescribeCrossTargetsResponse, err error) {
    return c.DescribeCrossTargetsWithContext(context.Background(), request)
}

// DescribeCrossTargets
// Queries information of CVMs and ENIs that use cross-region binding 2.0
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeCrossTargetsWithContext(ctx context.Context, request *DescribeCrossTargetsRequest) (response *DescribeCrossTargetsResponse, err error) {
    if request == nil {
        request = NewDescribeCrossTargetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCrossTargets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCrossTargetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomizedConfigAssociateListRequest() (request *DescribeCustomizedConfigAssociateListRequest) {
    request = &DescribeCustomizedConfigAssociateListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DescribeCustomizedConfigAssociateList")
    
    
    return
}

func NewDescribeCustomizedConfigAssociateListResponse() (response *DescribeCustomizedConfigAssociateListResponse) {
    response = &DescribeCustomizedConfigAssociateListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCustomizedConfigAssociateList
// This API is used to query the configured location, bound server or bound CLB instance. If there are domain names, the result will be filtered by domain name.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeCustomizedConfigAssociateList(request *DescribeCustomizedConfigAssociateListRequest) (response *DescribeCustomizedConfigAssociateListResponse, err error) {
    return c.DescribeCustomizedConfigAssociateListWithContext(context.Background(), request)
}

// DescribeCustomizedConfigAssociateList
// This API is used to query the configured location, bound server or bound CLB instance. If there are domain names, the result will be filtered by domain name.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeCustomizedConfigAssociateListWithContext(ctx context.Context, request *DescribeCustomizedConfigAssociateListRequest) (response *DescribeCustomizedConfigAssociateListResponse, err error) {
    if request == nil {
        request = NewDescribeCustomizedConfigAssociateListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCustomizedConfigAssociateList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCustomizedConfigAssociateListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomizedConfigListRequest() (request *DescribeCustomizedConfigListRequest) {
    request = &DescribeCustomizedConfigListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DescribeCustomizedConfigList")
    
    
    return
}

func NewDescribeCustomizedConfigListResponse() (response *DescribeCustomizedConfigListResponse) {
    response = &DescribeCustomizedConfigListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCustomizedConfigList
// This API is used to pull custom configuration lists to return the user configuration of `AppId`.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_LISTENERIDNOTFOUND = "InvalidParameter.ListenerIdNotFound"
//  INVALIDPARAMETER_LOCATIONNOTFOUND = "InvalidParameter.LocationNotFound"
//  INVALIDPARAMETER_PORTCHECKFAILED = "InvalidParameter.PortCheckFailed"
//  INVALIDPARAMETER_PROTOCOLCHECKFAILED = "InvalidParameter.ProtocolCheckFailed"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  INVALIDPARAMETER_REWRITEALREADYEXIST = "InvalidParameter.RewriteAlreadyExist"
//  INVALIDPARAMETER_SOMEREWRITENOTFOUND = "InvalidParameter.SomeRewriteNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeCustomizedConfigList(request *DescribeCustomizedConfigListRequest) (response *DescribeCustomizedConfigListResponse, err error) {
    return c.DescribeCustomizedConfigListWithContext(context.Background(), request)
}

// DescribeCustomizedConfigList
// This API is used to pull custom configuration lists to return the user configuration of `AppId`.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_LISTENERIDNOTFOUND = "InvalidParameter.ListenerIdNotFound"
//  INVALIDPARAMETER_LOCATIONNOTFOUND = "InvalidParameter.LocationNotFound"
//  INVALIDPARAMETER_PORTCHECKFAILED = "InvalidParameter.PortCheckFailed"
//  INVALIDPARAMETER_PROTOCOLCHECKFAILED = "InvalidParameter.ProtocolCheckFailed"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  INVALIDPARAMETER_REWRITEALREADYEXIST = "InvalidParameter.RewriteAlreadyExist"
//  INVALIDPARAMETER_SOMEREWRITENOTFOUND = "InvalidParameter.SomeRewriteNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeCustomizedConfigListWithContext(ctx context.Context, request *DescribeCustomizedConfigListRequest) (response *DescribeCustomizedConfigListResponse, err error) {
    if request == nil {
        request = NewDescribeCustomizedConfigListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCustomizedConfigList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCustomizedConfigListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIdleLoadBalancersRequest() (request *DescribeIdleLoadBalancersRequest) {
    request = &DescribeIdleLoadBalancersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DescribeIdleLoadBalancers")
    
    
    return
}

func NewDescribeIdleLoadBalancersResponse() (response *DescribeIdleLoadBalancersResponse) {
    response = &DescribeIdleLoadBalancersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIdleLoadBalancers
// Idle CLB instances are pay-as-you-go load balancers that, within seven days after the creation, do not have rules configured or the configured rules are not associated with any servers. 
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDLBSTATUS = "FailedOperation.InvalidLBStatus"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_LISTENERIDNOTFOUND = "InvalidParameter.ListenerIdNotFound"
//  INVALIDPARAMETER_LOCATIONNOTFOUND = "InvalidParameter.LocationNotFound"
//  INVALIDPARAMETER_PORTCHECKFAILED = "InvalidParameter.PortCheckFailed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
func (c *Client) DescribeIdleLoadBalancers(request *DescribeIdleLoadBalancersRequest) (response *DescribeIdleLoadBalancersResponse, err error) {
    return c.DescribeIdleLoadBalancersWithContext(context.Background(), request)
}

// DescribeIdleLoadBalancers
// Idle CLB instances are pay-as-you-go load balancers that, within seven days after the creation, do not have rules configured or the configured rules are not associated with any servers. 
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDLBSTATUS = "FailedOperation.InvalidLBStatus"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_LISTENERIDNOTFOUND = "InvalidParameter.ListenerIdNotFound"
//  INVALIDPARAMETER_LOCATIONNOTFOUND = "InvalidParameter.LocationNotFound"
//  INVALIDPARAMETER_PORTCHECKFAILED = "InvalidParameter.PortCheckFailed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
func (c *Client) DescribeIdleLoadBalancersWithContext(ctx context.Context, request *DescribeIdleLoadBalancersRequest) (response *DescribeIdleLoadBalancersResponse, err error) {
    if request == nil {
        request = NewDescribeIdleLoadBalancersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIdleLoadBalancers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIdleLoadBalancersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLBListenersRequest() (request *DescribeLBListenersRequest) {
    request = &DescribeLBListenersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DescribeLBListeners")
    
    
    return
}

func NewDescribeLBListenersResponse() (response *DescribeLBListenersResponse) {
    response = &DescribeLBListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLBListeners
// This API is used to query CLB instances bound to the CVM or ENI.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_LISTENERIDNOTFOUND = "InvalidParameter.ListenerIdNotFound"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeLBListeners(request *DescribeLBListenersRequest) (response *DescribeLBListenersResponse, err error) {
    return c.DescribeLBListenersWithContext(context.Background(), request)
}

// DescribeLBListeners
// This API is used to query CLB instances bound to the CVM or ENI.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_LISTENERIDNOTFOUND = "InvalidParameter.ListenerIdNotFound"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeLBListenersWithContext(ctx context.Context, request *DescribeLBListenersRequest) (response *DescribeLBListenersResponse, err error) {
    if request == nil {
        request = NewDescribeLBListenersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLBListeners require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLBListenersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListenersRequest() (request *DescribeListenersRequest) {
    request = &DescribeListenersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DescribeListeners")
    
    
    return
}

func NewDescribeListenersResponse() (response *DescribeListenersResponse) {
    response = &DescribeListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeListeners
// This API is used to get the list of listeners by CLB ID, listener protocol, or listener port. If no filter is specified, all listeners for the CLB instance will be returned.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeListeners(request *DescribeListenersRequest) (response *DescribeListenersResponse, err error) {
    return c.DescribeListenersWithContext(context.Background(), request)
}

// DescribeListeners
// This API is used to get the list of listeners by CLB ID, listener protocol, or listener port. If no filter is specified, all listeners for the CLB instance will be returned.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeListenersWithContext(ctx context.Context, request *DescribeListenersRequest) (response *DescribeListenersResponse, err error) {
    if request == nil {
        request = NewDescribeListenersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeListeners require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeListenersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLoadBalancerListByCertIdRequest() (request *DescribeLoadBalancerListByCertIdRequest) {
    request = &DescribeLoadBalancerListByCertIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DescribeLoadBalancerListByCertId")
    
    
    return
}

func NewDescribeLoadBalancerListByCertIdResponse() (response *DescribeLoadBalancerListByCertIdResponse) {
    response = &DescribeLoadBalancerListByCertIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLoadBalancerListByCertId
// This API is used to query the list of CLB instances associated with a certificate in a region by certificate ID.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeLoadBalancerListByCertId(request *DescribeLoadBalancerListByCertIdRequest) (response *DescribeLoadBalancerListByCertIdResponse, err error) {
    return c.DescribeLoadBalancerListByCertIdWithContext(context.Background(), request)
}

// DescribeLoadBalancerListByCertId
// This API is used to query the list of CLB instances associated with a certificate in a region by certificate ID.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeLoadBalancerListByCertIdWithContext(ctx context.Context, request *DescribeLoadBalancerListByCertIdRequest) (response *DescribeLoadBalancerListByCertIdResponse, err error) {
    if request == nil {
        request = NewDescribeLoadBalancerListByCertIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLoadBalancerListByCertId require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLoadBalancerListByCertIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLoadBalancerOverviewRequest() (request *DescribeLoadBalancerOverviewRequest) {
    request = &DescribeLoadBalancerOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DescribeLoadBalancerOverview")
    
    
    return
}

func NewDescribeLoadBalancerOverviewResponse() (response *DescribeLoadBalancerOverviewResponse) {
    response = &DescribeLoadBalancerOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLoadBalancerOverview
// Queries the total number of CLB instances and the number of CLB instances in different status (running, isolated and about to expire).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
func (c *Client) DescribeLoadBalancerOverview(request *DescribeLoadBalancerOverviewRequest) (response *DescribeLoadBalancerOverviewResponse, err error) {
    return c.DescribeLoadBalancerOverviewWithContext(context.Background(), request)
}

// DescribeLoadBalancerOverview
// Queries the total number of CLB instances and the number of CLB instances in different status (running, isolated and about to expire).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
func (c *Client) DescribeLoadBalancerOverviewWithContext(ctx context.Context, request *DescribeLoadBalancerOverviewRequest) (response *DescribeLoadBalancerOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeLoadBalancerOverviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLoadBalancerOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLoadBalancerOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLoadBalancerTrafficRequest() (request *DescribeLoadBalancerTrafficRequest) {
    request = &DescribeLoadBalancerTrafficRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DescribeLoadBalancerTraffic")
    
    
    return
}

func NewDescribeLoadBalancerTrafficResponse() (response *DescribeLoadBalancerTrafficResponse) {
    response = &DescribeLoadBalancerTrafficResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLoadBalancerTraffic
// This API is used to query CLB instances with high traffic under the current account, and return the top 10 results. For queries using a sub-account, only the CLB instances authorized to the sub-account will be returned.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
func (c *Client) DescribeLoadBalancerTraffic(request *DescribeLoadBalancerTrafficRequest) (response *DescribeLoadBalancerTrafficResponse, err error) {
    return c.DescribeLoadBalancerTrafficWithContext(context.Background(), request)
}

// DescribeLoadBalancerTraffic
// This API is used to query CLB instances with high traffic under the current account, and return the top 10 results. For queries using a sub-account, only the CLB instances authorized to the sub-account will be returned.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
func (c *Client) DescribeLoadBalancerTrafficWithContext(ctx context.Context, request *DescribeLoadBalancerTrafficRequest) (response *DescribeLoadBalancerTrafficResponse, err error) {
    if request == nil {
        request = NewDescribeLoadBalancerTrafficRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLoadBalancerTraffic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLoadBalancerTrafficResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLoadBalancersRequest() (request *DescribeLoadBalancersRequest) {
    request = &DescribeLoadBalancersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DescribeLoadBalancers")
    
    
    return
}

func NewDescribeLoadBalancersResponse() (response *DescribeLoadBalancersResponse) {
    response = &DescribeLoadBalancersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLoadBalancers
// This API is used to query the list of CLB instances in a region.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeLoadBalancers(request *DescribeLoadBalancersRequest) (response *DescribeLoadBalancersResponse, err error) {
    return c.DescribeLoadBalancersWithContext(context.Background(), request)
}

// DescribeLoadBalancers
// This API is used to query the list of CLB instances in a region.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeLoadBalancersWithContext(ctx context.Context, request *DescribeLoadBalancersRequest) (response *DescribeLoadBalancersResponse, err error) {
    if request == nil {
        request = NewDescribeLoadBalancersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLoadBalancers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLoadBalancersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLoadBalancersDetailRequest() (request *DescribeLoadBalancersDetailRequest) {
    request = &DescribeLoadBalancersDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DescribeLoadBalancersDetail")
    
    
    return
}

func NewDescribeLoadBalancersDetailResponse() (response *DescribeLoadBalancersDetailResponse) {
    response = &DescribeLoadBalancersDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLoadBalancersDetail
// This API is used to query CLB instance details, including listener, rules, and target real servers.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeLoadBalancersDetail(request *DescribeLoadBalancersDetailRequest) (response *DescribeLoadBalancersDetailResponse, err error) {
    return c.DescribeLoadBalancersDetailWithContext(context.Background(), request)
}

// DescribeLoadBalancersDetail
// This API is used to query CLB instance details, including listener, rules, and target real servers.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeLoadBalancersDetailWithContext(ctx context.Context, request *DescribeLoadBalancersDetailRequest) (response *DescribeLoadBalancersDetailResponse, err error) {
    if request == nil {
        request = NewDescribeLoadBalancersDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLoadBalancersDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLoadBalancersDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQuotaRequest() (request *DescribeQuotaRequest) {
    request = &DescribeQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DescribeQuota")
    
    
    return
}

func NewDescribeQuotaResponse() (response *DescribeQuotaResponse) {
    response = &DescribeQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeQuota
// This API is used to query various quotas in the current region.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeQuota(request *DescribeQuotaRequest) (response *DescribeQuotaResponse, err error) {
    return c.DescribeQuotaWithContext(context.Background(), request)
}

// DescribeQuota
// This API is used to query various quotas in the current region.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeQuotaWithContext(ctx context.Context, request *DescribeQuotaRequest) (response *DescribeQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeQuotaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeQuota require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourcesRequest() (request *DescribeResourcesRequest) {
    request = &DescribeResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DescribeResources")
    
    
    return
}

func NewDescribeResourcesResponse() (response *DescribeResourcesResponse) {
    response = &DescribeResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeResources
// This API is used to query the list of AZs and resources supported for the user in the current region.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
func (c *Client) DescribeResources(request *DescribeResourcesRequest) (response *DescribeResourcesResponse, err error) {
    return c.DescribeResourcesWithContext(context.Background(), request)
}

// DescribeResources
// This API is used to query the list of AZs and resources supported for the user in the current region.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
func (c *Client) DescribeResourcesWithContext(ctx context.Context, request *DescribeResourcesRequest) (response *DescribeResourcesResponse, err error) {
    if request == nil {
        request = NewDescribeResourcesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRewriteRequest() (request *DescribeRewriteRequest) {
    request = &DescribeRewriteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DescribeRewrite")
    
    
    return
}

func NewDescribeRewriteResponse() (response *DescribeRewriteResponse) {
    response = &DescribeRewriteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRewrite
// This API (DescribeRewrite) is used to query the redirection relationship between the forwarding rules of a CLB instance by instance ID. If no listener ID or forwarding rule ID is specified, all redirection relationships in the instance will be returned.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRewrite(request *DescribeRewriteRequest) (response *DescribeRewriteResponse, err error) {
    return c.DescribeRewriteWithContext(context.Background(), request)
}

// DescribeRewrite
// This API (DescribeRewrite) is used to query the redirection relationship between the forwarding rules of a CLB instance by instance ID. If no listener ID or forwarding rule ID is specified, all redirection relationships in the instance will be returned.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRewriteWithContext(ctx context.Context, request *DescribeRewriteRequest) (response *DescribeRewriteResponse, err error) {
    if request == nil {
        request = NewDescribeRewriteRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRewrite require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRewriteResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTargetGroupInstancesRequest() (request *DescribeTargetGroupInstancesRequest) {
    request = &DescribeTargetGroupInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DescribeTargetGroupInstances")
    
    
    return
}

func NewDescribeTargetGroupInstancesResponse() (response *DescribeTargetGroupInstancesResponse) {
    response = &DescribeTargetGroupInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTargetGroupInstances
// This API is used to get the information of servers bound to a target group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTargetGroupInstances(request *DescribeTargetGroupInstancesRequest) (response *DescribeTargetGroupInstancesResponse, err error) {
    return c.DescribeTargetGroupInstancesWithContext(context.Background(), request)
}

// DescribeTargetGroupInstances
// This API is used to get the information of servers bound to a target group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTargetGroupInstancesWithContext(ctx context.Context, request *DescribeTargetGroupInstancesRequest) (response *DescribeTargetGroupInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeTargetGroupInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTargetGroupInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTargetGroupInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTargetGroupListRequest() (request *DescribeTargetGroupListRequest) {
    request = &DescribeTargetGroupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DescribeTargetGroupList")
    
    
    return
}

func NewDescribeTargetGroupListResponse() (response *DescribeTargetGroupListResponse) {
    response = &DescribeTargetGroupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTargetGroupList
// This API is used to get the target group list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTargetGroupList(request *DescribeTargetGroupListRequest) (response *DescribeTargetGroupListResponse, err error) {
    return c.DescribeTargetGroupListWithContext(context.Background(), request)
}

// DescribeTargetGroupList
// This API is used to get the target group list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTargetGroupListWithContext(ctx context.Context, request *DescribeTargetGroupListRequest) (response *DescribeTargetGroupListResponse, err error) {
    if request == nil {
        request = NewDescribeTargetGroupListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTargetGroupList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTargetGroupListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTargetGroupsRequest() (request *DescribeTargetGroupsRequest) {
    request = &DescribeTargetGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DescribeTargetGroups")
    
    
    return
}

func NewDescribeTargetGroupsResponse() (response *DescribeTargetGroupsResponse) {
    response = &DescribeTargetGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTargetGroups
// This API is used to query the target group information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTargetGroups(request *DescribeTargetGroupsRequest) (response *DescribeTargetGroupsResponse, err error) {
    return c.DescribeTargetGroupsWithContext(context.Background(), request)
}

// DescribeTargetGroups
// This API is used to query the target group information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTargetGroupsWithContext(ctx context.Context, request *DescribeTargetGroupsRequest) (response *DescribeTargetGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeTargetGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTargetGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTargetGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTargetHealthRequest() (request *DescribeTargetHealthRequest) {
    request = &DescribeTargetHealthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DescribeTargetHealth")
    
    
    return
}

func NewDescribeTargetHealthResponse() (response *DescribeTargetHealthResponse) {
    response = &DescribeTargetHealthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTargetHealth
// This API (DescribeTargetHealth) is used to query the health check result of a real server of a CLB instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTargetHealth(request *DescribeTargetHealthRequest) (response *DescribeTargetHealthResponse, err error) {
    return c.DescribeTargetHealthWithContext(context.Background(), request)
}

// DescribeTargetHealth
// This API (DescribeTargetHealth) is used to query the health check result of a real server of a CLB instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTargetHealthWithContext(ctx context.Context, request *DescribeTargetHealthRequest) (response *DescribeTargetHealthResponse, err error) {
    if request == nil {
        request = NewDescribeTargetHealthRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTargetHealth require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTargetHealthResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTargetsRequest() (request *DescribeTargetsRequest) {
    request = &DescribeTargetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DescribeTargets")
    
    
    return
}

func NewDescribeTargetsResponse() (response *DescribeTargetsResponse) {
    response = &DescribeTargetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTargets
// This API (DescribeTargets) is used to query the list of real servers bound to some listeners of a CLB instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTargets(request *DescribeTargetsRequest) (response *DescribeTargetsResponse, err error) {
    return c.DescribeTargetsWithContext(context.Background(), request)
}

// DescribeTargets
// This API (DescribeTargets) is used to query the list of real servers bound to some listeners of a CLB instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTargetsWithContext(ctx context.Context, request *DescribeTargetsRequest) (response *DescribeTargetsResponse, err error) {
    if request == nil {
        request = NewDescribeTargetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTargets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTargetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskStatusRequest() (request *DescribeTaskStatusRequest) {
    request = &DescribeTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DescribeTaskStatus")
    
    
    return
}

func NewDescribeTaskStatusResponse() (response *DescribeTaskStatusResponse) {
    response = &DescribeTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTaskStatus
// This API is used to query the execution status of an async task. After non-query APIs (used to create/delete CLB instances, listeners, or rules or to bind/unbind real servers) are called successfully, this API needs to be used to query whether the task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
func (c *Client) DescribeTaskStatus(request *DescribeTaskStatusRequest) (response *DescribeTaskStatusResponse, err error) {
    return c.DescribeTaskStatusWithContext(context.Background(), request)
}

// DescribeTaskStatus
// This API is used to query the execution status of an async task. After non-query APIs (used to create/delete CLB instances, listeners, or rules or to bind/unbind real servers) are called successfully, this API needs to be used to query whether the task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
func (c *Client) DescribeTaskStatusWithContext(ctx context.Context, request *DescribeTaskStatusRequest) (response *DescribeTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeTaskStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateTargetGroupsRequest() (request *DisassociateTargetGroupsRequest) {
    request = &DisassociateTargetGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "DisassociateTargetGroups")
    
    
    return
}

func NewDisassociateTargetGroupsResponse() (response *DisassociateTargetGroupsResponse) {
    response = &DisassociateTargetGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisassociateTargetGroups
// This API is used to unbind target groups from a rule.
//
// This is an async API. After it is returned successfully, you can call the `DescribeTaskStatus` API with the returned `RequestID` as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DisassociateTargetGroups(request *DisassociateTargetGroupsRequest) (response *DisassociateTargetGroupsResponse, err error) {
    return c.DisassociateTargetGroupsWithContext(context.Background(), request)
}

// DisassociateTargetGroups
// This API is used to unbind target groups from a rule.
//
// This is an async API. After it is returned successfully, you can call the `DescribeTaskStatus` API with the returned `RequestID` as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DisassociateTargetGroupsWithContext(ctx context.Context, request *DisassociateTargetGroupsRequest) (response *DisassociateTargetGroupsResponse, err error) {
    if request == nil {
        request = NewDisassociateTargetGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisassociateTargetGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisassociateTargetGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceCreateLoadBalancerRequest() (request *InquiryPriceCreateLoadBalancerRequest) {
    request = &InquiryPriceCreateLoadBalancerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "InquiryPriceCreateLoadBalancer")
    
    
    return
}

func NewInquiryPriceCreateLoadBalancerResponse() (response *InquiryPriceCreateLoadBalancerResponse) {
    response = &InquiryPriceCreateLoadBalancerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquiryPriceCreateLoadBalancer
// This API is used to query the price of creating a CLB instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) InquiryPriceCreateLoadBalancer(request *InquiryPriceCreateLoadBalancerRequest) (response *InquiryPriceCreateLoadBalancerResponse, err error) {
    return c.InquiryPriceCreateLoadBalancerWithContext(context.Background(), request)
}

// InquiryPriceCreateLoadBalancer
// This API is used to query the price of creating a CLB instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) InquiryPriceCreateLoadBalancerWithContext(ctx context.Context, request *InquiryPriceCreateLoadBalancerRequest) (response *InquiryPriceCreateLoadBalancerResponse, err error) {
    if request == nil {
        request = NewInquiryPriceCreateLoadBalancerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceCreateLoadBalancer require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceCreateLoadBalancerResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceModifyLoadBalancerRequest() (request *InquiryPriceModifyLoadBalancerRequest) {
    request = &InquiryPriceModifyLoadBalancerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "InquiryPriceModifyLoadBalancer")
    
    
    return
}

func NewInquiryPriceModifyLoadBalancerResponse() (response *InquiryPriceModifyLoadBalancerResponse) {
    response = &InquiryPriceModifyLoadBalancerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquiryPriceModifyLoadBalancer
// This API is used to query the price of adjusting the specification of a CLB instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) InquiryPriceModifyLoadBalancer(request *InquiryPriceModifyLoadBalancerRequest) (response *InquiryPriceModifyLoadBalancerResponse, err error) {
    return c.InquiryPriceModifyLoadBalancerWithContext(context.Background(), request)
}

// InquiryPriceModifyLoadBalancer
// This API is used to query the price of adjusting the specification of a CLB instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) InquiryPriceModifyLoadBalancerWithContext(ctx context.Context, request *InquiryPriceModifyLoadBalancerRequest) (response *InquiryPriceModifyLoadBalancerResponse, err error) {
    if request == nil {
        request = NewInquiryPriceModifyLoadBalancerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceModifyLoadBalancer require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceModifyLoadBalancerResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceRefundLoadBalancerRequest() (request *InquiryPriceRefundLoadBalancerRequest) {
    request = &InquiryPriceRefundLoadBalancerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "InquiryPriceRefundLoadBalancer")
    
    
    return
}

func NewInquiryPriceRefundLoadBalancerResponse() (response *InquiryPriceRefundLoadBalancerResponse) {
    response = &InquiryPriceRefundLoadBalancerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquiryPriceRefundLoadBalancer
// This API is used to query the refund amount of returning a CLB instance. 
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) InquiryPriceRefundLoadBalancer(request *InquiryPriceRefundLoadBalancerRequest) (response *InquiryPriceRefundLoadBalancerResponse, err error) {
    return c.InquiryPriceRefundLoadBalancerWithContext(context.Background(), request)
}

// InquiryPriceRefundLoadBalancer
// This API is used to query the refund amount of returning a CLB instance. 
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) InquiryPriceRefundLoadBalancerWithContext(ctx context.Context, request *InquiryPriceRefundLoadBalancerRequest) (response *InquiryPriceRefundLoadBalancerResponse, err error) {
    if request == nil {
        request = NewInquiryPriceRefundLoadBalancerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceRefundLoadBalancer require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceRefundLoadBalancerResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceRenewLoadBalancerRequest() (request *InquiryPriceRenewLoadBalancerRequest) {
    request = &InquiryPriceRenewLoadBalancerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "InquiryPriceRenewLoadBalancer")
    
    
    return
}

func NewInquiryPriceRenewLoadBalancerResponse() (response *InquiryPriceRenewLoadBalancerResponse) {
    response = &InquiryPriceRenewLoadBalancerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquiryPriceRenewLoadBalancer
// This API is used to query the price of renewing a CLB instance. It's only available to prepaid CLB instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) InquiryPriceRenewLoadBalancer(request *InquiryPriceRenewLoadBalancerRequest) (response *InquiryPriceRenewLoadBalancerResponse, err error) {
    return c.InquiryPriceRenewLoadBalancerWithContext(context.Background(), request)
}

// InquiryPriceRenewLoadBalancer
// This API is used to query the price of renewing a CLB instance. It's only available to prepaid CLB instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) InquiryPriceRenewLoadBalancerWithContext(ctx context.Context, request *InquiryPriceRenewLoadBalancerRequest) (response *InquiryPriceRenewLoadBalancerResponse, err error) {
    if request == nil {
        request = NewInquiryPriceRenewLoadBalancerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceRenewLoadBalancer require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceRenewLoadBalancerResponse()
    err = c.Send(request, response)
    return
}

func NewManualRewriteRequest() (request *ManualRewriteRequest) {
    request = &ManualRewriteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "ManualRewrite")
    
    
    return
}

func NewManualRewriteResponse() (response *ManualRewriteResponse) {
    response = &ManualRewriteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ManualRewrite
// After the original access address and the address to be redirected are configured manually, the system will automatically redirect requests made to the original access address to the target address of the corresponding path. Multiple paths can be configured as a redirection policy under one domain name to achieve automatic redirection between HTTP and HTTPS. A redirection policy should meet the following rules: if A has already been redirected to B, then it cannot be redirected to C (unless the original redirection relationship is deleted and a new one is created), and B cannot be redirected to any other addresses.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDLBSTATUS = "FailedOperation.InvalidLBStatus"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_LISTENERIDNOTFOUND = "InvalidParameter.ListenerIdNotFound"
//  INVALIDPARAMETER_LOCATIONNOTFOUND = "InvalidParameter.LocationNotFound"
//  INVALIDPARAMETER_PROTOCOLCHECKFAILED = "InvalidParameter.ProtocolCheckFailed"
//  INVALIDPARAMETER_REWRITEALREADYEXIST = "InvalidParameter.RewriteAlreadyExist"
//  INVALIDPARAMETER_SOMEREWRITENOTFOUND = "InvalidParameter.SomeRewriteNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ManualRewrite(request *ManualRewriteRequest) (response *ManualRewriteResponse, err error) {
    return c.ManualRewriteWithContext(context.Background(), request)
}

// ManualRewrite
// After the original access address and the address to be redirected are configured manually, the system will automatically redirect requests made to the original access address to the target address of the corresponding path. Multiple paths can be configured as a redirection policy under one domain name to achieve automatic redirection between HTTP and HTTPS. A redirection policy should meet the following rules: if A has already been redirected to B, then it cannot be redirected to C (unless the original redirection relationship is deleted and a new one is created), and B cannot be redirected to any other addresses.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDLBSTATUS = "FailedOperation.InvalidLBStatus"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_LISTENERIDNOTFOUND = "InvalidParameter.ListenerIdNotFound"
//  INVALIDPARAMETER_LOCATIONNOTFOUND = "InvalidParameter.LocationNotFound"
//  INVALIDPARAMETER_PROTOCOLCHECKFAILED = "InvalidParameter.ProtocolCheckFailed"
//  INVALIDPARAMETER_REWRITEALREADYEXIST = "InvalidParameter.RewriteAlreadyExist"
//  INVALIDPARAMETER_SOMEREWRITENOTFOUND = "InvalidParameter.SomeRewriteNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ManualRewriteWithContext(ctx context.Context, request *ManualRewriteRequest) (response *ManualRewriteResponse, err error) {
    if request == nil {
        request = NewManualRewriteRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ManualRewrite require credential")
    }

    request.SetContext(ctx)
    
    response = NewManualRewriteResponse()
    err = c.Send(request, response)
    return
}

func NewMigrateClassicalLoadBalancersRequest() (request *MigrateClassicalLoadBalancersRequest) {
    request = &MigrateClassicalLoadBalancersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "MigrateClassicalLoadBalancers")
    
    
    return
}

func NewMigrateClassicalLoadBalancersResponse() (response *MigrateClassicalLoadBalancersResponse) {
    response = &MigrateClassicalLoadBalancersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// MigrateClassicalLoadBalancers
// This API is used to upgrade classic CLB instances to application CLB instances.
//
// This is an async API. After it is returned successfully, you can check the action result by calling `DescribeLoadBalancers`. 
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) MigrateClassicalLoadBalancers(request *MigrateClassicalLoadBalancersRequest) (response *MigrateClassicalLoadBalancersResponse, err error) {
    return c.MigrateClassicalLoadBalancersWithContext(context.Background(), request)
}

// MigrateClassicalLoadBalancers
// This API is used to upgrade classic CLB instances to application CLB instances.
//
// This is an async API. After it is returned successfully, you can check the action result by calling `DescribeLoadBalancers`. 
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) MigrateClassicalLoadBalancersWithContext(ctx context.Context, request *MigrateClassicalLoadBalancersRequest) (response *MigrateClassicalLoadBalancersResponse, err error) {
    if request == nil {
        request = NewMigrateClassicalLoadBalancersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("MigrateClassicalLoadBalancers require credential")
    }

    request.SetContext(ctx)
    
    response = NewMigrateClassicalLoadBalancersResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBlockIPListRequest() (request *ModifyBlockIPListRequest) {
    request = &ModifyBlockIPListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "ModifyBlockIPList")
    
    
    return
}

func NewModifyBlockIPListResponse() (response *ModifyBlockIPListResponse) {
    response = &ModifyBlockIPListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyBlockIPList
// This API is used to modify the client IP blocklist of a CLB instance. One forwarding rule supports blocking up to 2,000,000 IPs. One blocklist can contain up to 2,000,000 entries.
//
// (This API is in beta test. To use it, please submit a ticket.)
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyBlockIPList(request *ModifyBlockIPListRequest) (response *ModifyBlockIPListResponse, err error) {
    return c.ModifyBlockIPListWithContext(context.Background(), request)
}

// ModifyBlockIPList
// This API is used to modify the client IP blocklist of a CLB instance. One forwarding rule supports blocking up to 2,000,000 IPs. One blocklist can contain up to 2,000,000 entries.
//
// (This API is in beta test. To use it, please submit a ticket.)
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyBlockIPListWithContext(ctx context.Context, request *ModifyBlockIPListRequest) (response *ModifyBlockIPListResponse, err error) {
    if request == nil {
        request = NewModifyBlockIPListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBlockIPList require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBlockIPListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDomainRequest() (request *ModifyDomainRequest) {
    request = &ModifyDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "ModifyDomain")
    
    
    return
}

func NewModifyDomainResponse() (response *ModifyDomainResponse) {
    response = &ModifyDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDomain
// This API (ModifyDomain) is used to modify a domain name under a layer-7 CLB listener.
//
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestID as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyDomain(request *ModifyDomainRequest) (response *ModifyDomainResponse, err error) {
    return c.ModifyDomainWithContext(context.Background(), request)
}

// ModifyDomain
// This API (ModifyDomain) is used to modify a domain name under a layer-7 CLB listener.
//
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestID as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
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

func NewModifyDomainAttributesRequest() (request *ModifyDomainAttributesRequest) {
    request = &ModifyDomainAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "ModifyDomainAttributes")
    
    
    return
}

func NewModifyDomainAttributesResponse() (response *ModifyDomainAttributesResponse) {
    response = &ModifyDomainAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDomainAttributes
// This API is used to modify the domain name-level attributes of a layer-7 listener's forwarding rule, such as modifying the domain name, changing the DefaultServer, enabling/disabling HTTP/2, and modifying certificates.
//
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestId as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyDomainAttributes(request *ModifyDomainAttributesRequest) (response *ModifyDomainAttributesResponse, err error) {
    return c.ModifyDomainAttributesWithContext(context.Background(), request)
}

// ModifyDomainAttributes
// This API is used to modify the domain name-level attributes of a layer-7 listener's forwarding rule, such as modifying the domain name, changing the DefaultServer, enabling/disabling HTTP/2, and modifying certificates.
//
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestId as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyDomainAttributesWithContext(ctx context.Context, request *ModifyDomainAttributesRequest) (response *ModifyDomainAttributesResponse, err error) {
    if request == nil {
        request = NewModifyDomainAttributesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDomainAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDomainAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyFunctionTargetsRequest() (request *ModifyFunctionTargetsRequest) {
    request = &ModifyFunctionTargetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "ModifyFunctionTargets")
    
    
    return
}

func NewModifyFunctionTargetsResponse() (response *ModifyFunctionTargetsResponse) {
    response = &ModifyFunctionTargetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyFunctionTargets
// This API is used to modify the cloud functions associated with a load balancing forwarding rule.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyFunctionTargets(request *ModifyFunctionTargetsRequest) (response *ModifyFunctionTargetsResponse, err error) {
    return c.ModifyFunctionTargetsWithContext(context.Background(), request)
}

// ModifyFunctionTargets
// This API is used to modify the cloud functions associated with a load balancing forwarding rule.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyFunctionTargetsWithContext(ctx context.Context, request *ModifyFunctionTargetsRequest) (response *ModifyFunctionTargetsResponse, err error) {
    if request == nil {
        request = NewModifyFunctionTargetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyFunctionTargets require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyFunctionTargetsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyListenerRequest() (request *ModifyListenerRequest) {
    request = &ModifyListenerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "ModifyListener")
    
    
    return
}

func NewModifyListenerResponse() (response *ModifyListenerResponse) {
    response = &ModifyListenerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyListener
// This API (ModifyListener) is used to modify the attributes of a CLB listener, such as listener name, health check parameter, certificate information, and forwarding policy.
//
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestID as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyListener(request *ModifyListenerRequest) (response *ModifyListenerResponse, err error) {
    return c.ModifyListenerWithContext(context.Background(), request)
}

// ModifyListener
// This API (ModifyListener) is used to modify the attributes of a CLB listener, such as listener name, health check parameter, certificate information, and forwarding policy.
//
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestID as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyListenerWithContext(ctx context.Context, request *ModifyListenerRequest) (response *ModifyListenerResponse, err error) {
    if request == nil {
        request = NewModifyListenerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyListener require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyListenerResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLoadBalancerAttributesRequest() (request *ModifyLoadBalancerAttributesRequest) {
    request = &ModifyLoadBalancerAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "ModifyLoadBalancerAttributes")
    
    
    return
}

func NewModifyLoadBalancerAttributesResponse() (response *ModifyLoadBalancerAttributesResponse) {
    response = &ModifyLoadBalancerAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLoadBalancerAttributes
// This API is used to modify the attributes of a CLB instance such as name and cross-region attributes.
//
// This is an async API. After it is returned successfully, you can check the task result by calling `DescribeTaskStatus` with the returned `RequestID`.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyLoadBalancerAttributes(request *ModifyLoadBalancerAttributesRequest) (response *ModifyLoadBalancerAttributesResponse, err error) {
    return c.ModifyLoadBalancerAttributesWithContext(context.Background(), request)
}

// ModifyLoadBalancerAttributes
// This API is used to modify the attributes of a CLB instance such as name and cross-region attributes.
//
// This is an async API. After it is returned successfully, you can check the task result by calling `DescribeTaskStatus` with the returned `RequestID`.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyLoadBalancerAttributesWithContext(ctx context.Context, request *ModifyLoadBalancerAttributesRequest) (response *ModifyLoadBalancerAttributesResponse, err error) {
    if request == nil {
        request = NewModifyLoadBalancerAttributesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLoadBalancerAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLoadBalancerAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLoadBalancerSlaRequest() (request *ModifyLoadBalancerSlaRequest) {
    request = &ModifyLoadBalancerSlaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "ModifyLoadBalancerSla")
    
    
    return
}

func NewModifyLoadBalancerSlaResponse() (response *ModifyLoadBalancerSlaResponse) {
    response = &ModifyLoadBalancerSlaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLoadBalancerSla
// This API is used to upgrade a pay-as-you-go shared CLB instance to an LCU-supported CLB instance. <br/>
//
// Limits
//
// - This API can only be used to upgrade pay-as-you-go shared instances. To upgrade monthly-subscribed shared instances, please go to the CLB console.
//
// - An LCU-supported instance cannot be changed back to a shared instance.
//
// - Classic CLB instances cannot be upgraded to LCU-supported instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_LISTENERIDNOTFOUND = "InvalidParameter.ListenerIdNotFound"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyLoadBalancerSla(request *ModifyLoadBalancerSlaRequest) (response *ModifyLoadBalancerSlaResponse, err error) {
    return c.ModifyLoadBalancerSlaWithContext(context.Background(), request)
}

// ModifyLoadBalancerSla
// This API is used to upgrade a pay-as-you-go shared CLB instance to an LCU-supported CLB instance. <br/>
//
// Limits
//
// - This API can only be used to upgrade pay-as-you-go shared instances. To upgrade monthly-subscribed shared instances, please go to the CLB console.
//
// - An LCU-supported instance cannot be changed back to a shared instance.
//
// - Classic CLB instances cannot be upgraded to LCU-supported instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_LISTENERIDNOTFOUND = "InvalidParameter.ListenerIdNotFound"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyLoadBalancerSlaWithContext(ctx context.Context, request *ModifyLoadBalancerSlaRequest) (response *ModifyLoadBalancerSlaResponse, err error) {
    if request == nil {
        request = NewModifyLoadBalancerSlaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLoadBalancerSla require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLoadBalancerSlaResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLoadBalancersProjectRequest() (request *ModifyLoadBalancersProjectRequest) {
    request = &ModifyLoadBalancersProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "ModifyLoadBalancersProject")
    
    
    return
}

func NewModifyLoadBalancersProjectResponse() (response *ModifyLoadBalancersProjectResponse) {
    response = &ModifyLoadBalancersProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLoadBalancersProject
// This API is used to modify the projects of CLB instances. 
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyLoadBalancersProject(request *ModifyLoadBalancersProjectRequest) (response *ModifyLoadBalancersProjectResponse, err error) {
    return c.ModifyLoadBalancersProjectWithContext(context.Background(), request)
}

// ModifyLoadBalancersProject
// This API is used to modify the projects of CLB instances. 
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyLoadBalancersProjectWithContext(ctx context.Context, request *ModifyLoadBalancersProjectRequest) (response *ModifyLoadBalancersProjectResponse, err error) {
    if request == nil {
        request = NewModifyLoadBalancersProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLoadBalancersProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLoadBalancersProjectResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRuleRequest() (request *ModifyRuleRequest) {
    request = &ModifyRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "ModifyRule")
    
    
    return
}

func NewModifyRuleResponse() (response *ModifyRuleResponse) {
    response = &ModifyRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRule
// This API (ModifyRule) is used to modify the attributes of a forwarding rule under a layer-7 CLB listener, such as forwarding path, health check attribute, and forwarding policy.
//
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestID as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyRule(request *ModifyRuleRequest) (response *ModifyRuleResponse, err error) {
    return c.ModifyRuleWithContext(context.Background(), request)
}

// ModifyRule
// This API (ModifyRule) is used to modify the attributes of a forwarding rule under a layer-7 CLB listener, such as forwarding path, health check attribute, and forwarding policy.
//
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestID as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyRuleWithContext(ctx context.Context, request *ModifyRuleRequest) (response *ModifyRuleResponse, err error) {
    if request == nil {
        request = NewModifyRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTargetGroupAttributeRequest() (request *ModifyTargetGroupAttributeRequest) {
    request = &ModifyTargetGroupAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "ModifyTargetGroupAttribute")
    
    
    return
}

func NewModifyTargetGroupAttributeResponse() (response *ModifyTargetGroupAttributeResponse) {
    response = &ModifyTargetGroupAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTargetGroupAttribute
// This API is used to rename a target group or modify its default port attribute.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyTargetGroupAttribute(request *ModifyTargetGroupAttributeRequest) (response *ModifyTargetGroupAttributeResponse, err error) {
    return c.ModifyTargetGroupAttributeWithContext(context.Background(), request)
}

// ModifyTargetGroupAttribute
// This API is used to rename a target group or modify its default port attribute.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyTargetGroupAttributeWithContext(ctx context.Context, request *ModifyTargetGroupAttributeRequest) (response *ModifyTargetGroupAttributeResponse, err error) {
    if request == nil {
        request = NewModifyTargetGroupAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTargetGroupAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTargetGroupAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTargetGroupInstancesPortRequest() (request *ModifyTargetGroupInstancesPortRequest) {
    request = &ModifyTargetGroupInstancesPortRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "ModifyTargetGroupInstancesPort")
    
    
    return
}

func NewModifyTargetGroupInstancesPortResponse() (response *ModifyTargetGroupInstancesPortResponse) {
    response = &ModifyTargetGroupInstancesPortResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTargetGroupInstancesPort
// This API is used to modify server ports of a target group in batches.
//
// This is an async API. After it is returned successfully, you can call the `DescribeTaskStatus` API with the returned `RequestID` as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyTargetGroupInstancesPort(request *ModifyTargetGroupInstancesPortRequest) (response *ModifyTargetGroupInstancesPortResponse, err error) {
    return c.ModifyTargetGroupInstancesPortWithContext(context.Background(), request)
}

// ModifyTargetGroupInstancesPort
// This API is used to modify server ports of a target group in batches.
//
// This is an async API. After it is returned successfully, you can call the `DescribeTaskStatus` API with the returned `RequestID` as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyTargetGroupInstancesPortWithContext(ctx context.Context, request *ModifyTargetGroupInstancesPortRequest) (response *ModifyTargetGroupInstancesPortResponse, err error) {
    if request == nil {
        request = NewModifyTargetGroupInstancesPortRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTargetGroupInstancesPort require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTargetGroupInstancesPortResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTargetGroupInstancesWeightRequest() (request *ModifyTargetGroupInstancesWeightRequest) {
    request = &ModifyTargetGroupInstancesWeightRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "ModifyTargetGroupInstancesWeight")
    
    
    return
}

func NewModifyTargetGroupInstancesWeightResponse() (response *ModifyTargetGroupInstancesWeightResponse) {
    response = &ModifyTargetGroupInstancesWeightResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTargetGroupInstancesWeight
// This API is used to modify server weights of a target group in batches.
//
// This is an async API. After it is returned successfully, you can call the `DescribeTaskStatus` API with the returned `RequestID` as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyTargetGroupInstancesWeight(request *ModifyTargetGroupInstancesWeightRequest) (response *ModifyTargetGroupInstancesWeightResponse, err error) {
    return c.ModifyTargetGroupInstancesWeightWithContext(context.Background(), request)
}

// ModifyTargetGroupInstancesWeight
// This API is used to modify server weights of a target group in batches.
//
// This is an async API. After it is returned successfully, you can call the `DescribeTaskStatus` API with the returned `RequestID` as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyTargetGroupInstancesWeightWithContext(ctx context.Context, request *ModifyTargetGroupInstancesWeightRequest) (response *ModifyTargetGroupInstancesWeightResponse, err error) {
    if request == nil {
        request = NewModifyTargetGroupInstancesWeightRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTargetGroupInstancesWeight require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTargetGroupInstancesWeightResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTargetPortRequest() (request *ModifyTargetPortRequest) {
    request = &ModifyTargetPortRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "ModifyTargetPort")
    
    
    return
}

func NewModifyTargetPortResponse() (response *ModifyTargetPortResponse) {
    response = &ModifyTargetPortResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTargetPort
// This API (ModifyTargetPort) is used to modify the port of a real server bound to a listener.
//
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestID as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyTargetPort(request *ModifyTargetPortRequest) (response *ModifyTargetPortResponse, err error) {
    return c.ModifyTargetPortWithContext(context.Background(), request)
}

// ModifyTargetPort
// This API (ModifyTargetPort) is used to modify the port of a real server bound to a listener.
//
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestID as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyTargetPortWithContext(ctx context.Context, request *ModifyTargetPortRequest) (response *ModifyTargetPortResponse, err error) {
    if request == nil {
        request = NewModifyTargetPortRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTargetPort require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTargetPortResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTargetWeightRequest() (request *ModifyTargetWeightRequest) {
    request = &ModifyTargetWeightRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "ModifyTargetWeight")
    
    
    return
}

func NewModifyTargetWeightResponse() (response *ModifyTargetWeightResponse) {
    response = &ModifyTargetWeightResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTargetWeight
// This API (ModifyTargetWeight) is used to modify the forwarding weight of a real server bound to a CLB instance.
//
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestID as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyTargetWeight(request *ModifyTargetWeightRequest) (response *ModifyTargetWeightResponse, err error) {
    return c.ModifyTargetWeightWithContext(context.Background(), request)
}

// ModifyTargetWeight
// This API (ModifyTargetWeight) is used to modify the forwarding weight of a real server bound to a CLB instance.
//
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestID as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyTargetWeightWithContext(ctx context.Context, request *ModifyTargetWeightRequest) (response *ModifyTargetWeightResponse, err error) {
    if request == nil {
        request = NewModifyTargetWeightRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTargetWeight require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTargetWeightResponse()
    err = c.Send(request, response)
    return
}

func NewRegisterFunctionTargetsRequest() (request *RegisterFunctionTargetsRequest) {
    request = &RegisterFunctionTargetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "RegisterFunctionTargets")
    
    
    return
}

func NewRegisterFunctionTargetsResponse() (response *RegisterFunctionTargetsResponse) {
    response = &RegisterFunctionTargetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RegisterFunctionTargets
// This API is used to bind an SCF function with the L7 forwarding rule of a CLB instance. Note that you need to create an L7 listener (HTTP, HTTPS) and forwarding rule first.
//
// This is an async API. After it is returned successfully, you can call the `DescribeTaskStatus` API with the returned `RequestID` as an input parameter to check whether this task is successful.<br/>
//
// **Limits:**
//
// - Binding with SCF is only available in Guangzhou, Shenzhen Finance, Shanghai, Shanghai Finance, Beijing, Chengdu, Hong Kong (China), Singapore, Mumbai, Tokyo, and Silicon Valley.
//
// - SCF functions can only be bound with CLB instances of bill-by-IP accounts but not with bill-by-CVM accounts. If you are using a bill-by-CVM account, we recommend upgrading it to a bill-by-IP account. For more information, please see [Checking Account Type](https://intl.cloud.tencent.com/document/product/1199/49090?from_cn_redirect=1). 
//
// - SCF functions cannot be bound with classic CLB instances.
//
// - SCF functions cannot be bound with classic network-based CLB instances.
//
// - SCF functions in the same region can be bound with CLB instances. SCF functions can only be bound across VPCs but not regions.
//
// - SCF functions can only be bound with IPv4 and IPv6 NAT64 CLB instances, but currently not with IPv6 CLB instances.
//
// - SCF functions can only be bound with layer-7 HTTP and HTTPS listeners, but not with layer-7 QUIC listeners or layer-4 (TCP, UDP, and TCP SSL) listeners.
//
// - Only SCF event-triggered functions can be bound with CLB instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_LISTENERIDNOTFOUND = "InvalidParameter.ListenerIdNotFound"
//  INVALIDPARAMETER_LOCATIONNOTFOUND = "InvalidParameter.LocationNotFound"
//  INVALIDPARAMETER_PROTOCOLCHECKFAILED = "InvalidParameter.ProtocolCheckFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) RegisterFunctionTargets(request *RegisterFunctionTargetsRequest) (response *RegisterFunctionTargetsResponse, err error) {
    return c.RegisterFunctionTargetsWithContext(context.Background(), request)
}

// RegisterFunctionTargets
// This API is used to bind an SCF function with the L7 forwarding rule of a CLB instance. Note that you need to create an L7 listener (HTTP, HTTPS) and forwarding rule first.
//
// This is an async API. After it is returned successfully, you can call the `DescribeTaskStatus` API with the returned `RequestID` as an input parameter to check whether this task is successful.<br/>
//
// **Limits:**
//
// - Binding with SCF is only available in Guangzhou, Shenzhen Finance, Shanghai, Shanghai Finance, Beijing, Chengdu, Hong Kong (China), Singapore, Mumbai, Tokyo, and Silicon Valley.
//
// - SCF functions can only be bound with CLB instances of bill-by-IP accounts but not with bill-by-CVM accounts. If you are using a bill-by-CVM account, we recommend upgrading it to a bill-by-IP account. For more information, please see [Checking Account Type](https://intl.cloud.tencent.com/document/product/1199/49090?from_cn_redirect=1). 
//
// - SCF functions cannot be bound with classic CLB instances.
//
// - SCF functions cannot be bound with classic network-based CLB instances.
//
// - SCF functions in the same region can be bound with CLB instances. SCF functions can only be bound across VPCs but not regions.
//
// - SCF functions can only be bound with IPv4 and IPv6 NAT64 CLB instances, but currently not with IPv6 CLB instances.
//
// - SCF functions can only be bound with layer-7 HTTP and HTTPS listeners, but not with layer-7 QUIC listeners or layer-4 (TCP, UDP, and TCP SSL) listeners.
//
// - Only SCF event-triggered functions can be bound with CLB instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_LISTENERIDNOTFOUND = "InvalidParameter.ListenerIdNotFound"
//  INVALIDPARAMETER_LOCATIONNOTFOUND = "InvalidParameter.LocationNotFound"
//  INVALIDPARAMETER_PROTOCOLCHECKFAILED = "InvalidParameter.ProtocolCheckFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) RegisterFunctionTargetsWithContext(ctx context.Context, request *RegisterFunctionTargetsRequest) (response *RegisterFunctionTargetsResponse, err error) {
    if request == nil {
        request = NewRegisterFunctionTargetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RegisterFunctionTargets require credential")
    }

    request.SetContext(ctx)
    
    response = NewRegisterFunctionTargetsResponse()
    err = c.Send(request, response)
    return
}

func NewRegisterTargetGroupInstancesRequest() (request *RegisterTargetGroupInstancesRequest) {
    request = &RegisterTargetGroupInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "RegisterTargetGroupInstances")
    
    
    return
}

func NewRegisterTargetGroupInstancesResponse() (response *RegisterTargetGroupInstancesResponse) {
    response = &RegisterTargetGroupInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RegisterTargetGroupInstances
// This API is used to register servers to a target group.
//
// This is an async API. After it is returned successfully, you can call the `DescribeTaskStatus` API with the returned `RequestID` as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) RegisterTargetGroupInstances(request *RegisterTargetGroupInstancesRequest) (response *RegisterTargetGroupInstancesResponse, err error) {
    return c.RegisterTargetGroupInstancesWithContext(context.Background(), request)
}

// RegisterTargetGroupInstances
// This API is used to register servers to a target group.
//
// This is an async API. After it is returned successfully, you can call the `DescribeTaskStatus` API with the returned `RequestID` as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) RegisterTargetGroupInstancesWithContext(ctx context.Context, request *RegisterTargetGroupInstancesRequest) (response *RegisterTargetGroupInstancesResponse, err error) {
    if request == nil {
        request = NewRegisterTargetGroupInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RegisterTargetGroupInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewRegisterTargetGroupInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewRegisterTargetsRequest() (request *RegisterTargetsRequest) {
    request = &RegisterTargetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "RegisterTargets")
    
    
    return
}

func NewRegisterTargetsResponse() (response *RegisterTargetsResponse) {
    response = &RegisterTargetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RegisterTargets
// This API (RegisterTargets) is used to bind one or more real servers to a CLB listener or layer-7 forwarding rule. Before using this API, you need to create relevant layer-4 listeners or layer-7 forwarding rules. For the former (TCP/UDP), only the listener ID needs to be specified, while for the latter (HTTP/HTTPS), the forwarding rule also needs to be specified through LocationId or Domain+Url.
//
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestID as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) RegisterTargets(request *RegisterTargetsRequest) (response *RegisterTargetsResponse, err error) {
    return c.RegisterTargetsWithContext(context.Background(), request)
}

// RegisterTargets
// This API (RegisterTargets) is used to bind one or more real servers to a CLB listener or layer-7 forwarding rule. Before using this API, you need to create relevant layer-4 listeners or layer-7 forwarding rules. For the former (TCP/UDP), only the listener ID needs to be specified, while for the latter (HTTP/HTTPS), the forwarding rule also needs to be specified through LocationId or Domain+Url.
//
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestID as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) RegisterTargetsWithContext(ctx context.Context, request *RegisterTargetsRequest) (response *RegisterTargetsResponse, err error) {
    if request == nil {
        request = NewRegisterTargetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RegisterTargets require credential")
    }

    request.SetContext(ctx)
    
    response = NewRegisterTargetsResponse()
    err = c.Send(request, response)
    return
}

func NewRegisterTargetsWithClassicalLBRequest() (request *RegisterTargetsWithClassicalLBRequest) {
    request = &RegisterTargetsWithClassicalLBRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "RegisterTargetsWithClassicalLB")
    
    
    return
}

func NewRegisterTargetsWithClassicalLBResponse() (response *RegisterTargetsWithClassicalLBResponse) {
    response = &RegisterTargetsWithClassicalLBResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RegisterTargetsWithClassicalLB
// This API is used to bind a real server with a classic CLB instance. This is an async API. After it is returned successfully, you can call the API `DescribeTaskStatus` with the returned RequestId as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) RegisterTargetsWithClassicalLB(request *RegisterTargetsWithClassicalLBRequest) (response *RegisterTargetsWithClassicalLBResponse, err error) {
    return c.RegisterTargetsWithClassicalLBWithContext(context.Background(), request)
}

// RegisterTargetsWithClassicalLB
// This API is used to bind a real server with a classic CLB instance. This is an async API. After it is returned successfully, you can call the API `DescribeTaskStatus` with the returned RequestId as an input parameter to check whether this task is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) RegisterTargetsWithClassicalLBWithContext(ctx context.Context, request *RegisterTargetsWithClassicalLBRequest) (response *RegisterTargetsWithClassicalLBResponse, err error) {
    if request == nil {
        request = NewRegisterTargetsWithClassicalLBRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RegisterTargetsWithClassicalLB require credential")
    }

    request.SetContext(ctx)
    
    response = NewRegisterTargetsWithClassicalLBResponse()
    err = c.Send(request, response)
    return
}

func NewReplaceCertForLoadBalancersRequest() (request *ReplaceCertForLoadBalancersRequest) {
    request = &ReplaceCertForLoadBalancersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "ReplaceCertForLoadBalancers")
    
    
    return
}

func NewReplaceCertForLoadBalancersResponse() (response *ReplaceCertForLoadBalancersResponse) {
    response = &ReplaceCertForLoadBalancersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReplaceCertForLoadBalancers
// This API (ReplaceCertForLoadBalancers) is used to replace the certificate associated with a CLB instance. A new certificates can be associated with a CLB only after the original certificate is disassociated from it.
//
// This API supports replacing server certificates and client certificates.
//
// The new certificate to be used can be specified by passing in the certificate ID. If no certificate ID is specified, relevant information such as certificate content must be passed in to create a new certificate and bind it to the CLB.
//
// Note: This API can only be called in the Guangzhou region; for other regions, an error will occur due to domain name resolution problems.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
func (c *Client) ReplaceCertForLoadBalancers(request *ReplaceCertForLoadBalancersRequest) (response *ReplaceCertForLoadBalancersResponse, err error) {
    return c.ReplaceCertForLoadBalancersWithContext(context.Background(), request)
}

// ReplaceCertForLoadBalancers
// This API (ReplaceCertForLoadBalancers) is used to replace the certificate associated with a CLB instance. A new certificates can be associated with a CLB only after the original certificate is disassociated from it.
//
// This API supports replacing server certificates and client certificates.
//
// The new certificate to be used can be specified by passing in the certificate ID. If no certificate ID is specified, relevant information such as certificate content must be passed in to create a new certificate and bind it to the CLB.
//
// Note: This API can only be called in the Guangzhou region; for other regions, an error will occur due to domain name resolution problems.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
func (c *Client) ReplaceCertForLoadBalancersWithContext(ctx context.Context, request *ReplaceCertForLoadBalancersRequest) (response *ReplaceCertForLoadBalancersResponse, err error) {
    if request == nil {
        request = NewReplaceCertForLoadBalancersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReplaceCertForLoadBalancers require credential")
    }

    request.SetContext(ctx)
    
    response = NewReplaceCertForLoadBalancersResponse()
    err = c.Send(request, response)
    return
}

func NewSetCustomizedConfigForLoadBalancerRequest() (request *SetCustomizedConfigForLoadBalancerRequest) {
    request = &SetCustomizedConfigForLoadBalancerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "SetCustomizedConfigForLoadBalancer")
    
    
    return
}

func NewSetCustomizedConfigForLoadBalancerResponse() (response *SetCustomizedConfigForLoadBalancerResponse) {
    response = &SetCustomizedConfigForLoadBalancerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetCustomizedConfigForLoadBalancer
// This API is used to create or manage a user-defined CLB configuration template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SetCustomizedConfigForLoadBalancer(request *SetCustomizedConfigForLoadBalancerRequest) (response *SetCustomizedConfigForLoadBalancerResponse, err error) {
    return c.SetCustomizedConfigForLoadBalancerWithContext(context.Background(), request)
}

// SetCustomizedConfigForLoadBalancer
// This API is used to create or manage a user-defined CLB configuration template.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SetCustomizedConfigForLoadBalancerWithContext(ctx context.Context, request *SetCustomizedConfigForLoadBalancerRequest) (response *SetCustomizedConfigForLoadBalancerResponse, err error) {
    if request == nil {
        request = NewSetCustomizedConfigForLoadBalancerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetCustomizedConfigForLoadBalancer require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetCustomizedConfigForLoadBalancerResponse()
    err = c.Send(request, response)
    return
}

func NewSetLoadBalancerClsLogRequest() (request *SetLoadBalancerClsLogRequest) {
    request = &SetLoadBalancerClsLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "SetLoadBalancerClsLog")
    
    
    return
}

func NewSetLoadBalancerClsLogResponse() (response *SetLoadBalancerClsLogResponse) {
    response = &SetLoadBalancerClsLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetLoadBalancerClsLog
// This API is used to add, delete, and update the CLS topic of a CLB instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SetLoadBalancerClsLog(request *SetLoadBalancerClsLogRequest) (response *SetLoadBalancerClsLogResponse, err error) {
    return c.SetLoadBalancerClsLogWithContext(context.Background(), request)
}

// SetLoadBalancerClsLog
// This API is used to add, delete, and update the CLS topic of a CLB instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SetLoadBalancerClsLogWithContext(ctx context.Context, request *SetLoadBalancerClsLogRequest) (response *SetLoadBalancerClsLogResponse, err error) {
    if request == nil {
        request = NewSetLoadBalancerClsLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetLoadBalancerClsLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetLoadBalancerClsLogResponse()
    err = c.Send(request, response)
    return
}

func NewSetLoadBalancerSecurityGroupsRequest() (request *SetLoadBalancerSecurityGroupsRequest) {
    request = &SetLoadBalancerSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "SetLoadBalancerSecurityGroups")
    
    
    return
}

func NewSetLoadBalancerSecurityGroupsResponse() (response *SetLoadBalancerSecurityGroupsResponse) {
    response = &SetLoadBalancerSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetLoadBalancerSecurityGroups
// This API (SetLoadBalancerSecurityGroups) is used to bind/unbind security groups for a public network CLB instance. You can use the DescribeLoadBalancers API to query the security groups bound to a CLB instance. This API uses `set` semantics.
//
// During a binding operation, the input parameters need to be all security groups to be bound to the CLB instance (including those already bound ones and new ones).
//
// During an unbinding operation, the input parameters need to be all the security groups still bound to the CLB instance after the unbinding operation. To unbind all security groups, you can leave this parameter empty or pass in an empty array. Note: Private network CLB do not support binding security groups.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SetLoadBalancerSecurityGroups(request *SetLoadBalancerSecurityGroupsRequest) (response *SetLoadBalancerSecurityGroupsResponse, err error) {
    return c.SetLoadBalancerSecurityGroupsWithContext(context.Background(), request)
}

// SetLoadBalancerSecurityGroups
// This API (SetLoadBalancerSecurityGroups) is used to bind/unbind security groups for a public network CLB instance. You can use the DescribeLoadBalancers API to query the security groups bound to a CLB instance. This API uses `set` semantics.
//
// During a binding operation, the input parameters need to be all security groups to be bound to the CLB instance (including those already bound ones and new ones).
//
// During an unbinding operation, the input parameters need to be all the security groups still bound to the CLB instance after the unbinding operation. To unbind all security groups, you can leave this parameter empty or pass in an empty array. Note: Private network CLB do not support binding security groups.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SetLoadBalancerSecurityGroupsWithContext(ctx context.Context, request *SetLoadBalancerSecurityGroupsRequest) (response *SetLoadBalancerSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewSetLoadBalancerSecurityGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetLoadBalancerSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetLoadBalancerSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewSetSecurityGroupForLoadbalancersRequest() (request *SetSecurityGroupForLoadbalancersRequest) {
    request = &SetSecurityGroupForLoadbalancersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("clb", APIVersion, "SetSecurityGroupForLoadbalancers")
    
    
    return
}

func NewSetSecurityGroupForLoadbalancersResponse() (response *SetSecurityGroupForLoadbalancersResponse) {
    response = &SetSecurityGroupForLoadbalancersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetSecurityGroupForLoadbalancers
// This API is used to bind or unbind a security group for multiple public network CLB instances. Note: Private network CLB do not support binding security groups.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SetSecurityGroupForLoadbalancers(request *SetSecurityGroupForLoadbalancersRequest) (response *SetSecurityGroupForLoadbalancersResponse, err error) {
    return c.SetSecurityGroupForLoadbalancersWithContext(context.Background(), request)
}

// SetSecurityGroupForLoadbalancers
// This API is used to bind or unbind a security group for multiple public network CLB instances. Note: Private network CLB do not support binding security groups.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SetSecurityGroupForLoadbalancersWithContext(ctx context.Context, request *SetSecurityGroupForLoadbalancersRequest) (response *SetSecurityGroupForLoadbalancersResponse, err error) {
    if request == nil {
        request = NewSetSecurityGroupForLoadbalancersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetSecurityGroupForLoadbalancers require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetSecurityGroupForLoadbalancersResponse()
    err = c.Send(request, response)
    return
}
