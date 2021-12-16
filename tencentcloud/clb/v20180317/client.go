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
    if request == nil {
        request = NewAssociateTargetGroupsRequest()
    }
    
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
    if request == nil {
        request = NewAutoRewriteRequest()
    }
    
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
// This API is used to unbind layer-4 and layer-7 real servers in batches. Up to 500 servers can be unbound in a batch.
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
    if request == nil {
        request = NewBatchDeregisterTargetsRequest()
    }
    
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
    if request == nil {
        request = NewBatchModifyTargetWeightRequest()
    }
    
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
// This API is used to bind CVM instances or ENIs in batches. Up to 500 servers can be bound in a batch. It supports cross-region binding, and layer-4 and layer-7 (TCP/UDP/HTTP/HTTPS) protocols.
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
    if request == nil {
        request = NewBatchRegisterTargetsRequest()
    }
    
    response = NewBatchRegisterTargetsResponse()
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
    if request == nil {
        request = NewCreateClsLogSetRequest()
    }
    
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateListener(request *CreateListenerRequest) (response *CreateListenerResponse, err error) {
    if request == nil {
        request = NewCreateListenerRequest()
    }
    
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
func (c *Client) CreateLoadBalancer(request *CreateLoadBalancerRequest) (response *CreateLoadBalancerResponse, err error) {
    if request == nil {
        request = NewCreateLoadBalancerRequest()
    }
    
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
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateLoadBalancerSnatIps(request *CreateLoadBalancerSnatIpsRequest) (response *CreateLoadBalancerSnatIpsResponse, err error) {
    if request == nil {
        request = NewCreateLoadBalancerSnatIpsRequest()
    }
    
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
    if request == nil {
        request = NewCreateRuleRequest()
    }
    
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
    if request == nil {
        request = NewCreateTargetGroupRequest()
    }
    
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
    if request == nil {
        request = NewCreateTopicRequest()
    }
    
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
    if request == nil {
        request = NewDeleteListenerRequest()
    }
    
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
    if request == nil {
        request = NewDeleteLoadBalancerRequest()
    }
    
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
    if request == nil {
        request = NewDeleteLoadBalancerListenersRequest()
    }
    
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
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DeleteLoadBalancerSnatIps(request *DeleteLoadBalancerSnatIpsRequest) (response *DeleteLoadBalancerSnatIpsResponse, err error) {
    if request == nil {
        request = NewDeleteLoadBalancerSnatIpsRequest()
    }
    
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
    if request == nil {
        request = NewDeleteRewriteRequest()
    }
    
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
    if request == nil {
        request = NewDeleteRuleRequest()
    }
    
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
    if request == nil {
        request = NewDeleteTargetGroupsRequest()
    }
    
    response = NewDeleteTargetGroupsResponse()
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
    if request == nil {
        request = NewDeregisterTargetGroupInstancesRequest()
    }
    
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
    if request == nil {
        request = NewDeregisterTargetsRequest()
    }
    
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
    if request == nil {
        request = NewDeregisterTargetsFromClassicalLBRequest()
    }
    
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
    if request == nil {
        request = NewDescribeBlockIPListRequest()
    }
    
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBlockIPTask(request *DescribeBlockIPTaskRequest) (response *DescribeBlockIPTaskResponse, err error) {
    if request == nil {
        request = NewDescribeBlockIPTaskRequest()
    }
    
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
    if request == nil {
        request = NewDescribeClassicalLBByInstanceIdRequest()
    }
    
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
    if request == nil {
        request = NewDescribeClassicalLBHealthStatusRequest()
    }
    
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
func (c *Client) DescribeClassicalLBListeners(request *DescribeClassicalLBListenersRequest) (response *DescribeClassicalLBListenersResponse, err error) {
    if request == nil {
        request = NewDescribeClassicalLBListenersRequest()
    }
    
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
    if request == nil {
        request = NewDescribeClassicalLBTargetsRequest()
    }
    
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
    if request == nil {
        request = NewDescribeClsLogSetRequest()
    }
    
    response = NewDescribeClsLogSetResponse()
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
    if request == nil {
        request = NewDescribeCustomizedConfigAssociateListRequest()
    }
    
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
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeCustomizedConfigList(request *DescribeCustomizedConfigListRequest) (response *DescribeCustomizedConfigListResponse, err error) {
    if request == nil {
        request = NewDescribeCustomizedConfigListRequest()
    }
    
    response = NewDescribeCustomizedConfigListResponse()
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
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeLBListeners(request *DescribeLBListenersRequest) (response *DescribeLBListenersResponse, err error) {
    if request == nil {
        request = NewDescribeLBListenersRequest()
    }
    
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeListeners(request *DescribeListenersRequest) (response *DescribeListenersResponse, err error) {
    if request == nil {
        request = NewDescribeListenersRequest()
    }
    
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
    if request == nil {
        request = NewDescribeLoadBalancerListByCertIdRequest()
    }
    
    response = NewDescribeLoadBalancerListByCertIdResponse()
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
    if request == nil {
        request = NewDescribeLoadBalancerTrafficRequest()
    }
    
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeLoadBalancers(request *DescribeLoadBalancersRequest) (response *DescribeLoadBalancersResponse, err error) {
    if request == nil {
        request = NewDescribeLoadBalancersRequest()
    }
    
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
    if request == nil {
        request = NewDescribeLoadBalancersDetailRequest()
    }
    
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
func (c *Client) DescribeQuota(request *DescribeQuotaRequest) (response *DescribeQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeQuotaRequest()
    }
    
    response = NewDescribeQuotaResponse()
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
    if request == nil {
        request = NewDescribeRewriteRequest()
    }
    
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
    if request == nil {
        request = NewDescribeTargetGroupInstancesRequest()
    }
    
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
    if request == nil {
        request = NewDescribeTargetGroupListRequest()
    }
    
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
    if request == nil {
        request = NewDescribeTargetGroupsRequest()
    }
    
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
func (c *Client) DescribeTargetHealth(request *DescribeTargetHealthRequest) (response *DescribeTargetHealthResponse, err error) {
    if request == nil {
        request = NewDescribeTargetHealthRequest()
    }
    
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
func (c *Client) DescribeTargets(request *DescribeTargetsRequest) (response *DescribeTargetsResponse, err error) {
    if request == nil {
        request = NewDescribeTargetsRequest()
    }
    
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
func (c *Client) DescribeTaskStatus(request *DescribeTaskStatusRequest) (response *DescribeTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeTaskStatusRequest()
    }
    
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
    if request == nil {
        request = NewDisassociateTargetGroupsRequest()
    }
    
    response = NewDisassociateTargetGroupsResponse()
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
    if request == nil {
        request = NewManualRewriteRequest()
    }
    
    response = NewManualRewriteResponse()
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
    if request == nil {
        request = NewModifyBlockIPListRequest()
    }
    
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
    if request == nil {
        request = NewModifyDomainRequest()
    }
    
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
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyDomainAttributes(request *ModifyDomainAttributesRequest) (response *ModifyDomainAttributesResponse, err error) {
    if request == nil {
        request = NewModifyDomainAttributesRequest()
    }
    
    response = NewModifyDomainAttributesResponse()
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
    if request == nil {
        request = NewModifyListenerRequest()
    }
    
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
    if request == nil {
        request = NewModifyLoadBalancerAttributesRequest()
    }
    
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
// This API is used to upgrade shared CLB instances to LCU-supported CLB instances (downgrade is not allowed) and upgrade/downgrade the specification of LCU-supported instances.
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
    if request == nil {
        request = NewModifyLoadBalancerSlaRequest()
    }
    
    response = NewModifyLoadBalancerSlaResponse()
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
    if request == nil {
        request = NewModifyRuleRequest()
    }
    
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
    if request == nil {
        request = NewModifyTargetGroupAttributeRequest()
    }
    
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
    if request == nil {
        request = NewModifyTargetGroupInstancesPortRequest()
    }
    
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
    if request == nil {
        request = NewModifyTargetGroupInstancesWeightRequest()
    }
    
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
    if request == nil {
        request = NewModifyTargetPortRequest()
    }
    
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
    if request == nil {
        request = NewModifyTargetWeightRequest()
    }
    
    response = NewModifyTargetWeightResponse()
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
    if request == nil {
        request = NewRegisterTargetGroupInstancesRequest()
    }
    
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
    if request == nil {
        request = NewRegisterTargetsRequest()
    }
    
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
    if request == nil {
        request = NewRegisterTargetsWithClassicalLBRequest()
    }
    
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
    if request == nil {
        request = NewReplaceCertForLoadBalancersRequest()
    }
    
    response = NewReplaceCertForLoadBalancersResponse()
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
    if request == nil {
        request = NewSetLoadBalancerClsLogRequest()
    }
    
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
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SetLoadBalancerSecurityGroups(request *SetLoadBalancerSecurityGroupsRequest) (response *SetLoadBalancerSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewSetLoadBalancerSecurityGroupsRequest()
    }
    
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
    if request == nil {
        request = NewSetSecurityGroupForLoadbalancersRequest()
    }
    
    response = NewSetSecurityGroupForLoadbalancersResponse()
    err = c.Send(request, response)
    return
}
