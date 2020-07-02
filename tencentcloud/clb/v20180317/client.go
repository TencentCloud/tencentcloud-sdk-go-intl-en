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

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
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

// This API is used to bind target groups to CLB listeners (layer-4 protocol) or forwarding rules (layer-7 protocol).
// This is an async API. After it is returned successfully, you can call the `DescribeTaskStatus` API with the returned `RequestID` as an input parameter to check whether this task is successful.
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

// An HTTPS:443 listener needs to be created first, along with a forwarding rule. When this API is called, an HTTP:80 listener will be created automatically if it did not exist and a forwarding rule corresponding to `Domains` (specified in the input parameter) under the HTTPS:443 listener will also be created. After successful creation, access requests to an HTTP:80 address will be redirected to an HTTPS:443 address automatically.
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

// This API is used to unbind layer-4/layer-7 real servers in batches.
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

// This API is used to modify the forwarding weights of real servers bound to a CLB listener in batches. It supports layer-4 and layer-7 CLB listeners but not Classic CLB.
// This is an async API. After it is returned successfully, you can call the `DescribeTaskStatus` API with the returned `RequestID` as an input parameter to check whether this task is successful.
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

// This API is used to bind CVM instances or ENIs in batches. It supports cross-region binding and layer-4 and layer-7 (TCP, UDP, HTTP, HTTPS) protocols.
func (c *Client) BatchRegisterTargets(request *BatchRegisterTargetsRequest) (response *BatchRegisterTargetsResponse, err error) {
    if request == nil {
        request = NewBatchRegisterTargetsRequest()
    }
    response = NewBatchRegisterTargetsResponse()
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

// This API is used to create a listener for a CLB instance.
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestId as an input parameter to check whether this task is successful.
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

// This API (CreateLoadBalancer) is used to create a CLB instance. To use the CLB service, you first need to purchase one or more instances. After this API is called successfully, a unique instance ID will be returned. There are two types of instances: public network and private network. For more information, see the product types in the product documentation.
// Note: (1) To apply for a CLB instance in the specified AZ and cross-AZ disaster recovery, please [submit a ticket](https://console.cloud.tencent.com/workorder/category); (2) Currently, IPv6 is supported only in Beijing, Shanghai, and Guangzhou regions.
// This is an async API. After it is returned successfully, you can call the DescribeLoadBalancers API to query the status of the instance (such as creating and normal) to check whether it is successfully created.
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

// 针对SnatPro负载均衡，这个接口用于添加SnatIp，如果负载均衡没有开启SnatPro，添加SnatIp后会自动开启
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

// This API (CreateRule) is used to create a forwarding rule under an existing layer-7 CLB listener, where real servers must be bound to the rule instead of the listener.
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestID as an input parameter to check whether this task is successful.
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

// This API is used to create a target group. This feature is in beta test, if you want to try it out, please [submit a ticket](https://console.cloud.tencent.com/workorder/category?level1_id=6&level2_id=163&source=0&data_title=%E8%B4%9F%E8%BD%BD%E5%9D%87%E8%A1%A1%20LB&step=1).
func (c *Client) CreateTargetGroup(request *CreateTargetGroupRequest) (response *CreateTargetGroupResponse, err error) {
    if request == nil {
        request = NewCreateTargetGroupRequest()
    }
    response = NewCreateTargetGroupResponse()
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

// This API is used to delete a listener from a CLB instance (layer-4 or layer-7).
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestID as an input parameter to check whether this task is successful.
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

// This API (DeleteLoadBalancer) is used to delete one or more specified CLB instances.
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestId as an input parameter to check whether this task is successful.
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

// 该接口支持删除负载均衡的多个监听器。
// 本接口为异步接口，本接口返回成功后需以返回的 RequestID 为入参，调用 DescribeTaskStatus 接口查询本次任务是否成功。
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

// 对于SnatPro的负载均衡，这个接口用于删除SnatIp
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

// This API (DeleteRewrite) is used to delete the redirection relationship between the specified forwarding rules.
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

// This API (DeleteRule) is used to delete a forwarding rule under a layer-7 CLB instance listener
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestID as an input parameter to check whether this task is successful.
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

// This API is used to delete a target group.
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

// This API is used to unbind a server from a target group.
// This is an async API. After it is returned successfully, you can call the `DescribeTaskStatus` API with the returned `RequestID` as an input parameter to check whether this task is successful.
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

// This API (DeregisterTargets) is used to unbind one or more real servers from a CLB listener or forwarding rule. For layer-4 listeners, only the listener ID needs to be specified. For layer-7 listeners, the forwarding rule also needs to be specified through LocationId or Domain+Url.
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestID as an input parameter to check whether this task is successful.
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

// This API (DeregisterTargetsFromClassicalLB) is used to unbind real servers from a classic load balancer.
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestId as an input parameter to check whether this task is successful.
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

// 查询一个负载均衡所封禁的IP列表（黑名单）。（接口灰度中，如需使用请提工单）
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

// 根据 ModifyBlockIPList 接口返回的异步任务的ID，查询封禁IP（黑名单）异步任务的执行状态。（接口灰度中，如需使用请提工单）
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

// This API (DescribeClassicalLBByInstanceId) is used to get the list of classic CLB IDs through the real server instance ID.
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

// This API (DescribeClassicalLBHealthStatus) is used to get the real server health status of a classic CLB
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

// This API (DescribeClassicalLBListeners) is used to get the listener information of a classic CLB.
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

// This API (DescribeClassicalLBTargets) is used to get the real servers bound to a classic CLB.
func (c *Client) DescribeClassicalLBTargets(request *DescribeClassicalLBTargetsRequest) (response *DescribeClassicalLBTargetsResponse, err error) {
    if request == nil {
        request = NewDescribeClassicalLBTargetsRequest()
    }
    response = NewDescribeClassicalLBTargetsResponse()
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

// This API is used to get the list of listeners by CLB instance ID, listener protocol, or port. If no filter is specified, all listeners under the CLB instance will be returned.
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

// This API is used to query the list of CLB instances associated with a certificate in a region by certificate ID.
func (c *Client) DescribeLoadBalancerListByCertId(request *DescribeLoadBalancerListByCertIdRequest) (response *DescribeLoadBalancerListByCertIdResponse, err error) {
    if request == nil {
        request = NewDescribeLoadBalancerListByCertIdRequest()
    }
    response = NewDescribeLoadBalancerListByCertIdResponse()
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

// This API is used to query the list of CLB instances in a region.
func (c *Client) DescribeLoadBalancers(request *DescribeLoadBalancersRequest) (response *DescribeLoadBalancersResponse, err error) {
    if request == nil {
        request = NewDescribeLoadBalancersRequest()
    }
    response = NewDescribeLoadBalancersResponse()
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

// This API (DescribeRewrite) is used to query the redirection relationship between the forwarding rules of a CLB instance by instance ID. If no listener ID or forwarding rule ID is specified, all redirection relationships in the instance will be returned.
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

// This API is used to get the information of servers bound to a target group.
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

// This API is used to get the target group list.
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

// This API is used to query the target group information.
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

// This API (DescribeTargetHealth) is used to query the health check result of a real server of a CLB instance.
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

// This API (DescribeTargets) is used to query the list of real servers bound to some listeners of a CLB instance.
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

// This API is used to query the execution status of an async task. After non-query APIs (used to create/delete CLB instances, listeners, or rules or to bind/unbind real servers) are called successfully, this API needs to be used to query whether the task is successful.
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

// This API is used to unbind target groups from a rule.
// This is an async API. After it is returned successfully, you can call the `DescribeTaskStatus` API with the returned `RequestID` as an input parameter to check whether this task is successful.
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

// After the original access address and the address to be redirected are configured manually, the system will automatically redirect requests made to the original access address to the target address of the corresponding path. Multiple paths can be configured as a redirection policy under one domain name to achieve automatic redirection between HTTP and HTTPS. A redirection policy should meet the following rules: if A has already been redirected to B, then it cannot be redirected to C (unless the original redirection relationship is deleted and a new one is created), and B cannot be redirected to any other addresses.
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

// 修改负载均衡的IP（client IP）封禁黑名单列表，一个转发规则最多支持封禁 2000000 个IP，及黑名单容量为 2000000。
// （接口灰度中，如需使用请提工单）
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

// This API (ModifyDomain) is used to modify a domain name under a layer-7 CLB listener.
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestID as an input parameter to check whether this task is successful.
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

// This API is used to modify the domain name-level attributes of a layer-7 listener's forwarding rule, such as modifying the domain name, changing the DefaultServer, enabling/disabling HTTP/2, and modifying certificates.
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestId as an input parameter to check whether this task is successful.
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

// This API (ModifyListener) is used to modify the attributes of a CLB listener, such as listener name, health check parameter, certificate information, and forwarding policy.
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestID as an input parameter to check whether this task is successful.
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

// This API is used to modify the attributes of a CLB instance such as name and cross-region attributes.
func (c *Client) ModifyLoadBalancerAttributes(request *ModifyLoadBalancerAttributesRequest) (response *ModifyLoadBalancerAttributesResponse, err error) {
    if request == nil {
        request = NewModifyLoadBalancerAttributesRequest()
    }
    response = NewModifyLoadBalancerAttributesResponse()
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

// This API (ModifyRule) is used to modify the attributes of a forwarding rule under a layer-7 CLB listener, such as forwarding path, health check attribute, and forwarding policy.
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestID as an input parameter to check whether this task is successful.
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

// This API is used to rename a target group or modify its default port attribute.
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

// This API is used to modify server ports of a target group in batches.
// This is an async API. After it is returned successfully, you can call the `DescribeTaskStatus` API with the returned `RequestID` as an input parameter to check whether this task is successful.
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

// This API is used to modify server weights of a target group in batches.
// This is an async API. After it is returned successfully, you can call the `DescribeTaskStatus` API with the returned `RequestID` as an input parameter to check whether this task is successful.
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

// This API (ModifyTargetPort) is used to modify the port of a real server bound to a listener.
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestID as an input parameter to check whether this task is successful.
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

// This API (ModifyTargetWeight) is used to modify the forwarding weight of a real server bound to a CLB instance.
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestID as an input parameter to check whether this task is successful.
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

// This API is used to register servers to a target group.
// This is an async API. After it is returned successfully, you can call the `DescribeTaskStatus` API with the returned `RequestID` as an input parameter to check whether this task is successful.
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

// This API (RegisterTargets) is used to bind one or more real servers to a CLB listener or layer-7 forwarding rule. Before using this API, you need to create relevant layer-4 listeners or layer-7 forwarding rules. For the former (TCP/UDP), only the listener ID needs to be specified, while for the latter (HTTP/HTTPS), the forwarding rule also needs to be specified through LocationId or Domain+Url.
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestID as an input parameter to check whether this task is successful.
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

// This API (RegisterTargetsWithClassicalLB) is used to bind real servers to a classic CLB.
// This is an async API. After it is returned successfully, you can call the DescribeTaskStatus API with the returned RequestId as an input parameter to check whether this task is successful.
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

// This API (ReplaceCertForLoadBalancers) is used to replace the certificate associated with a CLB instance. A new certificates can be associated with a CLB only after the original certificate is disassociated from it.
// This API supports replacing server certificates and client certificates.
// The new certificate to be used can be specified by passing in the certificate ID. If no certificate ID is specified, relevant information such as certificate content must be passed in to create a new certificate and bind it to the CLB.
// Note: This API can only be called in the Guangzhou region; for other regions, an error will occur due to domain name resolution problems.
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

// This API is used to add, delete, and update the CLS topic of a CLB instance.
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

// This API (SetLoadBalancerSecurityGroups) is used to bind/unbind security groups for a public network CLB instance. You can use the DescribeLoadBalancers API to query the security groups bound to a CLB instance. This API uses `set` semantics.
// During a binding operation, the input parameters need to be all security groups to be bound to the CLB instance (including those already bound ones and new ones).
// During an unbinding operation, the input parameters need to be all the security groups still bound to the CLB instance after the unbinding operation. To unbind all security groups, you can leave this parameter empty or pass in an empty array. Note: Private network CLB do not support binding security groups.
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

// This API is used to bind or unbind a security group for multiple public network CLB instances. Note: Private network CLB do not support binding security groups.
func (c *Client) SetSecurityGroupForLoadbalancers(request *SetSecurityGroupForLoadbalancersRequest) (response *SetSecurityGroupForLoadbalancersResponse, err error) {
    if request == nil {
        request = NewSetSecurityGroupForLoadbalancersRequest()
    }
    response = NewSetSecurityGroupForLoadbalancersResponse()
    err = c.Send(request, response)
    return
}
