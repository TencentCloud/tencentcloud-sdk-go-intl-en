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

package v20230417

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2023-04-17"

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
// The BatchModifyTargetWeight API is used to modify the forwarding weight of backend machines bound to a Cloud Load Balancer listener in batch. The resource limit is 500. This is an async API. After it returns a successful result, call DescribeTaskStatus API with the returned RequestID as input parameter to check whether this task is successful.<br/>This API is supported by layer-4 and layer-7 CLB listeners but not by classic CLB.
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
// The BatchModifyTargetWeight API is used to modify the forwarding weight of backend machines bound to a Cloud Load Balancer listener in batch. The resource limit is 500. This is an async API. After it returns a successful result, call DescribeTaskStatus API with the returned RequestID as input parameter to check whether this task is successful.<br/>This API is supported by layer-4 and layer-7 CLB listeners but not by classic CLB.
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
    c.InitBaseRequest(&request.BaseRequest, "clb", APIVersion, "BatchModifyTargetWeight")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchModifyTargetWeight require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchModifyTargetWeightResponse()
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
// This API is used to clone a Cloud Load Balancer instance. Based on the designated CLB instance, it creates a new instance with identical rules and binding relationships. The cloning operation is asynchronous. The cloned data is based on the call to CloneLoadBalancer. If the cloned CLB changes after calling CloneLoadBalancer, the change rules will not be cloned.
//
// 
//
// Limitation notes:
//
// Does not support basic networks, classic CLB, IPv6, or NAT64.
//
// Unsupported Monthly Subscription CLB
//
// The listener does not support QUIC or port ranges.
//
// No support for backend types: target group, Serverless Cloud Function (SCF).
//
// Personalized configuration, redirection configuration, and security group default allow switch will not be cloned and must be manually configured.
//
// 
//
// API calling
//
// BGP bandwidth package must include bandwidth package ID
//
// Exclusive cluster cloning must pass corresponding parameters, otherwise shared instance creation is used.
//
// The feature is in beta test. You can submit a [beta test application](https://www.tencentcloud.com/apply/p/1akuvsmyn0g?from_cn_redirect=1).
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
// This API is used to clone a Cloud Load Balancer instance. Based on the designated CLB instance, it creates a new instance with identical rules and binding relationships. The cloning operation is asynchronous. The cloned data is based on the call to CloneLoadBalancer. If the cloned CLB changes after calling CloneLoadBalancer, the change rules will not be cloned.
//
// 
//
// Limitation notes:
//
// Does not support basic networks, classic CLB, IPv6, or NAT64.
//
// Unsupported Monthly Subscription CLB
//
// The listener does not support QUIC or port ranges.
//
// No support for backend types: target group, Serverless Cloud Function (SCF).
//
// Personalized configuration, redirection configuration, and security group default allow switch will not be cloned and must be manually configured.
//
// 
//
// API calling
//
// BGP bandwidth package must include bandwidth package ID
//
// Exclusive cluster cloning must pass corresponding parameters, otherwise shared instance creation is used.
//
// The feature is in beta test. You can submit a [beta test application](https://www.tencentcloud.com/apply/p/1akuvsmyn0g?from_cn_redirect=1).
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
    c.InitBaseRequest(&request.BaseRequest, "clb", APIVersion, "CloneLoadBalancer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloneLoadBalancer require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloneLoadBalancerResponse()
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
// This API is used to create a Cloud Load Balancer instance (this interface supports only pay-as-you-go CLB instances. For annual/monthly subscription, proceed to purchase through the console). To use the CLB service, you must purchase one or more CLB instances. After successfully calling the API, the unique ID of the CLB instance will be returned. CLB instances are divided into public network and private network types. For details, refer to the product type in the product description.
//
// Note: (1) To apply for Cloud Load Balancer (CLB) in a specified availability zone or cross-zone disaster recovery (only supported in Hong Kong (China)), [submit a request](https://console.cloud.tencent.com/workorder/category) if you need to experience the feature. (2) Currently only Beijing, Shanghai, and Guangzhou support IPv6. (3) The default purchase quota for an account in each region is 100 for public network and 100 for private network.
//
// This is an async API. After the API returns successfully, you can use the DescribeLoadBalancers API to query the status of the Cloud Load Balancer instance (such as creating and normal) to confirm whether the creation is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLIENTTOKENLIMITEXCEEDED = "InvalidParameter.ClientTokenLimitExceeded"
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
// This API is used to create a Cloud Load Balancer instance (this interface supports only pay-as-you-go CLB instances. For annual/monthly subscription, proceed to purchase through the console). To use the CLB service, you must purchase one or more CLB instances. After successfully calling the API, the unique ID of the CLB instance will be returned. CLB instances are divided into public network and private network types. For details, refer to the product type in the product description.
//
// Note: (1) To apply for Cloud Load Balancer (CLB) in a specified availability zone or cross-zone disaster recovery (only supported in Hong Kong (China)), [submit a request](https://console.cloud.tencent.com/workorder/category) if you need to experience the feature. (2) Currently only Beijing, Shanghai, and Guangzhou support IPv6. (3) The default purchase quota for an account in each region is 100 for public network and 100 for private network.
//
// This is an async API. After the API returns successfully, you can use the DescribeLoadBalancers API to query the status of the Cloud Load Balancer instance (such as creating and normal) to confirm whether the creation is successful.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLIENTTOKENLIMITEXCEEDED = "InvalidParameter.ClientTokenLimitExceeded"
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
    c.InitBaseRequest(&request.BaseRequest, "clb", APIVersion, "CreateLoadBalancer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLoadBalancer require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLoadBalancerResponse()
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
// This API is used to obtain listener information of classic CLB.
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
// This API is used to obtain listener information of classic CLB.
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
    c.InitBaseRequest(&request.BaseRequest, "clb", APIVersion, "DescribeClassicalLBListeners")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClassicalLBListeners require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClassicalLBListenersResponse()
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
// This API is used to query the locations, servers, or CLB instances bound to a configuration. If there are domains, the result will be filtered by domain.
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
// This API is used to query the locations, servers, or CLB instances bound to a configuration. If there are domains, the result will be filtered by domain.
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
    c.InitBaseRequest(&request.BaseRequest, "clb", APIVersion, "DescribeCustomizedConfigAssociateList")
    
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
// Pull the list of custom configurations. The configurations of the specified type under the user's AppId will be returned.
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
// Pull the list of custom configurations. The configurations of the specified type under the user's AppId will be returned.
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
    c.InitBaseRequest(&request.BaseRequest, "clb", APIVersion, "DescribeCustomizedConfigList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCustomizedConfigList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCustomizedConfigListResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "clb", APIVersion, "DescribeLoadBalancerListByCertId")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLoadBalancerListByCertId require credential")
    }

    request.SetContext(ctx)
    
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
    c.InitBaseRequest(&request.BaseRequest, "clb", APIVersion, "DescribeLoadBalancers")
    
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
// Query the detailed information of Cloud Load Balancer, including listeners, rules, and backend targets.
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
// Query the detailed information of Cloud Load Balancer, including listeners, rules, and backend targets.
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
    c.InitBaseRequest(&request.BaseRequest, "clb", APIVersion, "DescribeLoadBalancersDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLoadBalancersDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLoadBalancersDetailResponse()
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
// This API (DescribeTargetHealth) is used to query the health check result of a real server of a CLB instance. It is not supported by Classic CLB.
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
// This API (DescribeTargetHealth) is used to query the health check result of a real server of a CLB instance. It is not supported by Classic CLB.
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
    c.InitBaseRequest(&request.BaseRequest, "clb", APIVersion, "DescribeTargetHealth")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTargetHealth require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTargetHealthResponse()
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
// This API is used to create, delete, modify, bind, and unbind custom CLB configurations.
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
// This API is used to create, delete, modify, bind, and unbind custom CLB configurations.
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
    c.InitBaseRequest(&request.BaseRequest, "clb", APIVersion, "SetCustomizedConfigForLoadBalancer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetCustomizedConfigForLoadBalancer require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetCustomizedConfigForLoadBalancerResponse()
    err = c.Send(request, response)
    return
}
