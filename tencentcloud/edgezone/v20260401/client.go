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

package v20260401

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2026-04-01"

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


func NewApplyPublicIpsRequest() (request *ApplyPublicIpsRequest) {
    request = &ApplyPublicIpsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("edgezone", APIVersion, "ApplyPublicIps")
    
    
    return
}

func NewApplyPublicIpsResponse() (response *ApplyPublicIpsResponse) {
    response = &ApplyPublicIpsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ApplyPublicIps
// This API is used to submit a request for multiple IP addresses from the static IP pool to specify a public network instance (random allocation). Need to check user quota before applying.
//
// This API is applicable only to public network instances with `RouteMode=static`. Calling this API for BGP/OSPF instances will return an error.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTSUPPORTEDFORDYNAMICINSTANCE = "FailedOperation.NotSupportedForDynamicInstance"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_IPV6QUOTAEXCEEDED = "LimitExceeded.Ipv6QuotaExceeded"
//  LIMITEXCEEDED_IPV6QUOTANOTCONFIGURED = "LimitExceeded.Ipv6QuotaNotConfigured"
//  LIMITEXCEEDED_QUOTAEXCEEDED = "LimitExceeded.QuotaExceeded"
//  LIMITEXCEEDED_QUOTANOTCONFIGURED = "LimitExceeded.QuotaNotConfigured"
//  RESOURCEINSUFFICIENT_PUBLICIPINSUFFICIENT = "ResourceInsufficient.PublicIpInsufficient"
//  RESOURCEINSUFFICIENT_PUBLICIPV6INSUFFICIENT = "ResourceInsufficient.PublicIpv6Insufficient"
func (c *Client) ApplyPublicIps(request *ApplyPublicIpsRequest) (response *ApplyPublicIpsResponse, err error) {
    return c.ApplyPublicIpsWithContext(context.Background(), request)
}

// ApplyPublicIps
// This API is used to submit a request for multiple IP addresses from the static IP pool to specify a public network instance (random allocation). Need to check user quota before applying.
//
// This API is applicable only to public network instances with `RouteMode=static`. Calling this API for BGP/OSPF instances will return an error.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTSUPPORTEDFORDYNAMICINSTANCE = "FailedOperation.NotSupportedForDynamicInstance"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_IPV6QUOTAEXCEEDED = "LimitExceeded.Ipv6QuotaExceeded"
//  LIMITEXCEEDED_IPV6QUOTANOTCONFIGURED = "LimitExceeded.Ipv6QuotaNotConfigured"
//  LIMITEXCEEDED_QUOTAEXCEEDED = "LimitExceeded.QuotaExceeded"
//  LIMITEXCEEDED_QUOTANOTCONFIGURED = "LimitExceeded.QuotaNotConfigured"
//  RESOURCEINSUFFICIENT_PUBLICIPINSUFFICIENT = "ResourceInsufficient.PublicIpInsufficient"
//  RESOURCEINSUFFICIENT_PUBLICIPV6INSUFFICIENT = "ResourceInsufficient.PublicIpv6Insufficient"
func (c *Client) ApplyPublicIpsWithContext(ctx context.Context, request *ApplyPublicIpsRequest) (response *ApplyPublicIpsResponse, err error) {
    if request == nil {
        request = NewApplyPublicIpsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "edgezone", APIVersion, "ApplyPublicIps")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyPublicIps require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyPublicIpsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstancesRequest() (request *CreateInstancesRequest) {
    request = &CreateInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("edgezone", APIVersion, "CreateInstances")
    
    
    return
}

func NewCreateInstancesResponse() (response *CreateInstancesResponse) {
    response = &CreateInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateInstances
// This API is used to create a physical machine instance. The system automatically allocates physical machine resources and completes installation. If the user is not in the current availability zone, the system automatically enables billing. It supports concurrent allocation of physical machine resources and async execution of network assignment and installation tasks.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_INSTANCEQUOTA = "LimitExceeded.InstanceQuota"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) CreateInstances(request *CreateInstancesRequest) (response *CreateInstancesResponse, err error) {
    return c.CreateInstancesWithContext(context.Background(), request)
}

// CreateInstances
// This API is used to create a physical machine instance. The system automatically allocates physical machine resources and completes installation. If the user is not in the current availability zone, the system automatically enables billing. It supports concurrent allocation of physical machine resources and async execution of network assignment and installation tasks.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_INSTANCEQUOTA = "LimitExceeded.InstanceQuota"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) CreateInstancesWithContext(ctx context.Context, request *CreateInstancesRequest) (response *CreateInstancesResponse, err error) {
    if request == nil {
        request = NewCreateInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "edgezone", APIVersion, "CreateInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrivateNetworkInstanceRequest() (request *CreatePrivateNetworkInstanceRequest) {
    request = &CreatePrivateNetworkInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("edgezone", APIVersion, "CreatePrivateNetworkInstance")
    
    
    return
}

func NewCreatePrivateNetworkInstanceResponse() (response *CreatePrivateNetworkInstanceResponse) {
    response = &CreatePrivateNetworkInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePrivateNetworkInstance
// Create a private network instance. A user can only create one private network instance in an availability zone. The subnet range is collectively determined by both parameters: Network (network number) and Mask (bit number of the mask). Network must be a valid network address from one of the three RFC 1918 private address ranges: 10.0.0.0/8, 172.16.0.0/12, or 192.168.0.0/16, and all host bits must be 0 (meaning the combination of Network and Mask cannot have any host bits set, such as 10.0.0.1/24 is illegal, use 10.0.0.0/24 instead). The maximum Mask is unified as 28, and the minimum is determined by the address range it belongs to: 10.x.x.x allows 8–28, 172.16.x.x allows 12–28, and 192.168.x.x allows 16–28.
//
// error code that may be returned:
//  FAILEDOPERATION_PRIVATEINSTANCEDUPLICATE = "FailedOperation.PrivateInstanceDuplicate"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDNETWORK = "InvalidParameterValue.InvalidNetwork"
func (c *Client) CreatePrivateNetworkInstance(request *CreatePrivateNetworkInstanceRequest) (response *CreatePrivateNetworkInstanceResponse, err error) {
    return c.CreatePrivateNetworkInstanceWithContext(context.Background(), request)
}

// CreatePrivateNetworkInstance
// Create a private network instance. A user can only create one private network instance in an availability zone. The subnet range is collectively determined by both parameters: Network (network number) and Mask (bit number of the mask). Network must be a valid network address from one of the three RFC 1918 private address ranges: 10.0.0.0/8, 172.16.0.0/12, or 192.168.0.0/16, and all host bits must be 0 (meaning the combination of Network and Mask cannot have any host bits set, such as 10.0.0.1/24 is illegal, use 10.0.0.0/24 instead). The maximum Mask is unified as 28, and the minimum is determined by the address range it belongs to: 10.x.x.x allows 8–28, 172.16.x.x allows 12–28, and 192.168.x.x allows 16–28.
//
// error code that may be returned:
//  FAILEDOPERATION_PRIVATEINSTANCEDUPLICATE = "FailedOperation.PrivateInstanceDuplicate"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDNETWORK = "InvalidParameterValue.InvalidNetwork"
func (c *Client) CreatePrivateNetworkInstanceWithContext(ctx context.Context, request *CreatePrivateNetworkInstanceRequest) (response *CreatePrivateNetworkInstanceResponse, err error) {
    if request == nil {
        request = NewCreatePrivateNetworkInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "edgezone", APIVersion, "CreatePrivateNetworkInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrivateNetworkInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrivateNetworkInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePublicNetworkInstanceRequest() (request *CreatePublicNetworkInstanceRequest) {
    request = &CreatePublicNetworkInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("edgezone", APIVersion, "CreatePublicNetworkInstance")
    
    
    return
}

func NewCreatePublicNetworkInstanceResponse() (response *CreatePublicNetworkInstanceResponse) {
    response = &CreatePublicNetworkInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePublicNetworkInstance
// The user inputs the availability zone ID, public network instance name, network line, and routing mode to create a public network instance. A user can only create one public network instance in an availability zone.
//
// Public network instances in **static** routing mode require users to proactively apply for and release public IP addresses.
//
// Public network instances with routing mode set to **OSPF, BGP** automatically allocate public IP ranges when creating and auto release public IP ranges upon termination.
//
// error code that may be returned:
//  FAILEDOPERATION_PUBLICINSTANCEDUPLICATE = "FailedOperation.PublicInstanceDuplicate"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDQUOTA = "InvalidParameter.InvalidQuota"
//  LIMITEXCEEDED_QUOTANOTCONFIGURED = "LimitExceeded.QuotaNotConfigured"
//  RESOURCEINSUFFICIENT_NOAVAILABLECIDR = "ResourceInsufficient.NoAvailableCidr"
func (c *Client) CreatePublicNetworkInstance(request *CreatePublicNetworkInstanceRequest) (response *CreatePublicNetworkInstanceResponse, err error) {
    return c.CreatePublicNetworkInstanceWithContext(context.Background(), request)
}

// CreatePublicNetworkInstance
// The user inputs the availability zone ID, public network instance name, network line, and routing mode to create a public network instance. A user can only create one public network instance in an availability zone.
//
// Public network instances in **static** routing mode require users to proactively apply for and release public IP addresses.
//
// Public network instances with routing mode set to **OSPF, BGP** automatically allocate public IP ranges when creating and auto release public IP ranges upon termination.
//
// error code that may be returned:
//  FAILEDOPERATION_PUBLICINSTANCEDUPLICATE = "FailedOperation.PublicInstanceDuplicate"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDQUOTA = "InvalidParameter.InvalidQuota"
//  LIMITEXCEEDED_QUOTANOTCONFIGURED = "LimitExceeded.QuotaNotConfigured"
//  RESOURCEINSUFFICIENT_NOAVAILABLECIDR = "ResourceInsufficient.NoAvailableCidr"
func (c *Client) CreatePublicNetworkInstanceWithContext(ctx context.Context, request *CreatePublicNetworkInstanceRequest) (response *CreatePublicNetworkInstanceResponse, err error) {
    if request == nil {
        request = NewCreatePublicNetworkInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "edgezone", APIVersion, "CreatePublicNetworkInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePublicNetworkInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePublicNetworkInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrivateNetworkInstanceRequest() (request *DeletePrivateNetworkInstanceRequest) {
    request = &DeletePrivateNetworkInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("edgezone", APIVersion, "DeletePrivateNetworkInstance")
    
    
    return
}

func NewDeletePrivateNetworkInstanceResponse() (response *DeletePrivateNetworkInstanceResponse) {
    response = &DeletePrivateNetworkInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePrivateNetworkInstance
// Delete a private network instance
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE_PRIVATEINSTANCEINUSE = "ResourceInUse.PrivateInstanceInUse"
//  RESOURCENOTFOUND_PRIVATEINSTANCENOTFOUND = "ResourceNotFound.PrivateInstanceNotFound"
func (c *Client) DeletePrivateNetworkInstance(request *DeletePrivateNetworkInstanceRequest) (response *DeletePrivateNetworkInstanceResponse, err error) {
    return c.DeletePrivateNetworkInstanceWithContext(context.Background(), request)
}

// DeletePrivateNetworkInstance
// Delete a private network instance
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE_PRIVATEINSTANCEINUSE = "ResourceInUse.PrivateInstanceInUse"
//  RESOURCENOTFOUND_PRIVATEINSTANCENOTFOUND = "ResourceNotFound.PrivateInstanceNotFound"
func (c *Client) DeletePrivateNetworkInstanceWithContext(ctx context.Context, request *DeletePrivateNetworkInstanceRequest) (response *DeletePrivateNetworkInstanceResponse, err error) {
    if request == nil {
        request = NewDeletePrivateNetworkInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "edgezone", APIVersion, "DeletePrivateNetworkInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePrivateNetworkInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePrivateNetworkInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePublicNetworkInstanceRequest() (request *DeletePublicNetworkInstanceRequest) {
    request = &DeletePublicNetworkInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("edgezone", APIVersion, "DeletePublicNetworkInstance")
    
    
    return
}

func NewDeletePublicNetworkInstanceResponse() (response *DeletePublicNetworkInstanceResponse) {
    response = &DeletePublicNetworkInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePublicNetworkInstance
// Modify public network instance info
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE_PUBLICINSTANCEINUSE = "ResourceInUse.PublicInstanceInUse"
//  RESOURCENOTFOUND_PUBLICINSTANCENOTFOUND = "ResourceNotFound.PublicInstanceNotFound"
func (c *Client) DeletePublicNetworkInstance(request *DeletePublicNetworkInstanceRequest) (response *DeletePublicNetworkInstanceResponse, err error) {
    return c.DeletePublicNetworkInstanceWithContext(context.Background(), request)
}

// DeletePublicNetworkInstance
// Modify public network instance info
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE_PUBLICINSTANCEINUSE = "ResourceInUse.PublicInstanceInUse"
//  RESOURCENOTFOUND_PUBLICINSTANCENOTFOUND = "ResourceNotFound.PublicInstanceNotFound"
func (c *Client) DeletePublicNetworkInstanceWithContext(ctx context.Context, request *DeletePublicNetworkInstanceRequest) (response *DeletePublicNetworkInstanceResponse, err error) {
    if request == nil {
        request = NewDeletePublicNetworkInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "edgezone", APIVersion, "DeletePublicNetworkInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePublicNetworkInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePublicNetworkInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceTypesRequest() (request *DescribeInstanceTypesRequest) {
    request = &DescribeInstanceTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("edgezone", APIVersion, "DescribeInstanceTypes")
    
    
    return
}

func NewDescribeInstanceTypesResponse() (response *DescribeInstanceTypesResponse) {
    response = &DescribeInstanceTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceTypes
// This API is used to query the model quota list under the account by availability zone dimensionality based on AppId. If a Zone is input, it returns the model quota under the designated availability zone. If not, it returns the model quota of all AZs under the account.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE_PUBLICINSTANCEINUSE = "ResourceInUse.PublicInstanceInUse"
//  RESOURCENOTFOUND_PUBLICINSTANCENOTFOUND = "ResourceNotFound.PublicInstanceNotFound"
func (c *Client) DescribeInstanceTypes(request *DescribeInstanceTypesRequest) (response *DescribeInstanceTypesResponse, err error) {
    return c.DescribeInstanceTypesWithContext(context.Background(), request)
}

// DescribeInstanceTypes
// This API is used to query the model quota list under the account by availability zone dimensionality based on AppId. If a Zone is input, it returns the model quota under the designated availability zone. If not, it returns the model quota of all AZs under the account.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE_PUBLICINSTANCEINUSE = "ResourceInUse.PublicInstanceInUse"
//  RESOURCENOTFOUND_PUBLICINSTANCENOTFOUND = "ResourceNotFound.PublicInstanceNotFound"
func (c *Client) DescribeInstanceTypesWithContext(ctx context.Context, request *DescribeInstanceTypesRequest) (response *DescribeInstanceTypesResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceTypesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "edgezone", APIVersion, "DescribeInstanceTypes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceTypes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceTypesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("edgezone", APIVersion, "DescribeInstances")
    
    
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstances
// This API is used to query physical machine instance list, support by instance ID, instance name, availability zone, instance status with conditional filtering, and support paging query.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE_PUBLICINSTANCEINUSE = "ResourceInUse.PublicInstanceInUse"
//  RESOURCENOTFOUND_PUBLICINSTANCENOTFOUND = "ResourceNotFound.PublicInstanceNotFound"
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    return c.DescribeInstancesWithContext(context.Background(), request)
}

// DescribeInstances
// This API is used to query physical machine instance list, support by instance ID, instance name, availability zone, instance status with conditional filtering, and support paging query.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE_PUBLICINSTANCEINUSE = "ResourceInUse.PublicInstanceInUse"
//  RESOURCENOTFOUND_PUBLICINSTANCENOTFOUND = "ResourceNotFound.PublicInstanceNotFound"
func (c *Client) DescribeInstancesWithContext(ctx context.Context, request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "edgezone", APIVersion, "DescribeInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrivateNetworkInstancesRequest() (request *DescribePrivateNetworkInstancesRequest) {
    request = &DescribePrivateNetworkInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("edgezone", APIVersion, "DescribePrivateNetworkInstances")
    
    
    return
}

func NewDescribePrivateNetworkInstancesResponse() (response *DescribePrivateNetworkInstancesResponse) {
    response = &DescribePrivateNetworkInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrivateNetworkInstances
// Query private network instances, support through parameters such as private network instance ID, instance name, and availability zone id.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePrivateNetworkInstances(request *DescribePrivateNetworkInstancesRequest) (response *DescribePrivateNetworkInstancesResponse, err error) {
    return c.DescribePrivateNetworkInstancesWithContext(context.Background(), request)
}

// DescribePrivateNetworkInstances
// Query private network instances, support through parameters such as private network instance ID, instance name, and availability zone id.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePrivateNetworkInstancesWithContext(ctx context.Context, request *DescribePrivateNetworkInstancesRequest) (response *DescribePrivateNetworkInstancesResponse, err error) {
    if request == nil {
        request = NewDescribePrivateNetworkInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "edgezone", APIVersion, "DescribePrivateNetworkInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrivateNetworkInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrivateNetworkInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublicIpsRequest() (request *DescribePublicIpsRequest) {
    request = &DescribePublicIpsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("edgezone", APIVersion, "DescribePublicIps")
    
    
    return
}

func NewDescribePublicIpsResponse() (response *DescribePublicIpsResponse) {
    response = &DescribePublicIpsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePublicIps
// Query the public network Ip information of the user. For public network instances with routing mode set to Static, return all applied public network Ip information. For public network instances with routing mode set to Ospf and Bgp, return Ip range information directly.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePublicIps(request *DescribePublicIpsRequest) (response *DescribePublicIpsResponse, err error) {
    return c.DescribePublicIpsWithContext(context.Background(), request)
}

// DescribePublicIps
// Query the public network Ip information of the user. For public network instances with routing mode set to Static, return all applied public network Ip information. For public network instances with routing mode set to Ospf and Bgp, return Ip range information directly.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePublicIpsWithContext(ctx context.Context, request *DescribePublicIpsRequest) (response *DescribePublicIpsResponse, err error) {
    if request == nil {
        request = NewDescribePublicIpsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "edgezone", APIVersion, "DescribePublicIps")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePublicIps require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePublicIpsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublicNetworkInstancesRequest() (request *DescribePublicNetworkInstancesRequest) {
    request = &DescribePublicNetworkInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("edgezone", APIVersion, "DescribePublicNetworkInstances")
    
    
    return
}

func NewDescribePublicNetworkInstancesResponse() (response *DescribePublicNetworkInstancesResponse) {
    response = &DescribePublicNetworkInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePublicNetworkInstances
// Query public network instance list, support conditional filtering by instance ID, instance name, availability zone, and support paging query.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_PUBLICINSTANCENOTFOUND = "ResourceNotFound.PublicInstanceNotFound"
func (c *Client) DescribePublicNetworkInstances(request *DescribePublicNetworkInstancesRequest) (response *DescribePublicNetworkInstancesResponse, err error) {
    return c.DescribePublicNetworkInstancesWithContext(context.Background(), request)
}

// DescribePublicNetworkInstances
// Query public network instance list, support conditional filtering by instance ID, instance name, availability zone, and support paging query.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_PUBLICINSTANCENOTFOUND = "ResourceNotFound.PublicInstanceNotFound"
func (c *Client) DescribePublicNetworkInstancesWithContext(ctx context.Context, request *DescribePublicNetworkInstancesRequest) (response *DescribePublicNetworkInstancesResponse, err error) {
    if request == nil {
        request = NewDescribePublicNetworkInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "edgezone", APIVersion, "DescribePublicNetworkInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePublicNetworkInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePublicNetworkInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZoneDataRequest() (request *DescribeZoneDataRequest) {
    request = &DescribeZoneDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("edgezone", APIVersion, "DescribeZoneData")
    
    
    return
}

func NewDescribeZoneDataResponse() (response *DescribeZoneDataResponse) {
    response = &DescribeZoneDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeZoneData
// Query statistics by metric name. Data is aggregated at 1-minute intervals.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) DescribeZoneData(request *DescribeZoneDataRequest) (response *DescribeZoneDataResponse, err error) {
    return c.DescribeZoneDataWithContext(context.Background(), request)
}

// DescribeZoneData
// Query statistics by metric name. Data is aggregated at 1-minute intervals.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) DescribeZoneDataWithContext(ctx context.Context, request *DescribeZoneDataRequest) (response *DescribeZoneDataResponse, err error) {
    if request == nil {
        request = NewDescribeZoneDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "edgezone", APIVersion, "DescribeZoneData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZoneData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeZoneDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZonesRequest() (request *DescribeZonesRequest) {
    request = &DescribeZonesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("edgezone", APIVersion, "DescribeZones")
    
    
    return
}

func NewDescribeZonesResponse() (response *DescribeZonesResponse) {
    response = &DescribeZonesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeZones
// Cross-regional aggregate query returns the AZ list for the specified AppId in ALL configured regions. The local region directly performs a database query, while remote regions send HTTP requests to each region's DescribeAppZones API and merge the results.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) DescribeZones(request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    return c.DescribeZonesWithContext(context.Background(), request)
}

// DescribeZones
// Cross-regional aggregate query returns the AZ list for the specified AppId in ALL configured regions. The local region directly performs a database query, while remote regions send HTTP requests to each region's DescribeAppZones API and merge the results.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) DescribeZonesWithContext(ctx context.Context, request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    if request == nil {
        request = NewDescribeZonesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "edgezone", APIVersion, "DescribeZones")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZones require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeZonesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceAttributeRequest() (request *ModifyInstanceAttributeRequest) {
    request = &ModifyInstanceAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("edgezone", APIVersion, "ModifyInstanceAttribute")
    
    
    return
}

func NewModifyInstanceAttributeResponse() (response *ModifyInstanceAttributeResponse) {
    response = &ModifyInstanceAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstanceAttribute
// This API is used to modify the attributes of a physical machine instance, supporting modification of the instance name and update of the public IP address (IPv4/IPv6). At least one of InstanceName and NewPublicIp must be input.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) ModifyInstanceAttribute(request *ModifyInstanceAttributeRequest) (response *ModifyInstanceAttributeResponse, err error) {
    return c.ModifyInstanceAttributeWithContext(context.Background(), request)
}

// ModifyInstanceAttribute
// This API is used to modify the attributes of a physical machine instance, supporting modification of the instance name and update of the public IP address (IPv4/IPv6). At least one of InstanceName and NewPublicIp must be input.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) ModifyInstanceAttributeWithContext(ctx context.Context, request *ModifyInstanceAttributeRequest) (response *ModifyInstanceAttributeResponse, err error) {
    if request == nil {
        request = NewModifyInstanceAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "edgezone", APIVersion, "ModifyInstanceAttribute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrivateNetworkInstanceRequest() (request *ModifyPrivateNetworkInstanceRequest) {
    request = &ModifyPrivateNetworkInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("edgezone", APIVersion, "ModifyPrivateNetworkInstance")
    
    
    return
}

func NewModifyPrivateNetworkInstanceResponse() (response *ModifyPrivateNetworkInstanceResponse) {
    response = &ModifyPrivateNetworkInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPrivateNetworkInstance
// Modify private network instance info
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_PRIVATEINSTANCENOTFOUND = "ResourceNotFound.PrivateInstanceNotFound"
func (c *Client) ModifyPrivateNetworkInstance(request *ModifyPrivateNetworkInstanceRequest) (response *ModifyPrivateNetworkInstanceResponse, err error) {
    return c.ModifyPrivateNetworkInstanceWithContext(context.Background(), request)
}

// ModifyPrivateNetworkInstance
// Modify private network instance info
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_PRIVATEINSTANCENOTFOUND = "ResourceNotFound.PrivateInstanceNotFound"
func (c *Client) ModifyPrivateNetworkInstanceWithContext(ctx context.Context, request *ModifyPrivateNetworkInstanceRequest) (response *ModifyPrivateNetworkInstanceResponse, err error) {
    if request == nil {
        request = NewModifyPrivateNetworkInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "edgezone", APIVersion, "ModifyPrivateNetworkInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPrivateNetworkInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPrivateNetworkInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPublicNetworkInstanceRequest() (request *ModifyPublicNetworkInstanceRequest) {
    request = &ModifyPublicNetworkInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("edgezone", APIVersion, "ModifyPublicNetworkInstance")
    
    
    return
}

func NewModifyPublicNetworkInstanceResponse() (response *ModifyPublicNetworkInstanceResponse) {
    response = &ModifyPublicNetworkInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPublicNetworkInstance
// Modify public network instance info
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_PUBLICINSTANCENOTFOUND = "ResourceNotFound.PublicInstanceNotFound"
func (c *Client) ModifyPublicNetworkInstance(request *ModifyPublicNetworkInstanceRequest) (response *ModifyPublicNetworkInstanceResponse, err error) {
    return c.ModifyPublicNetworkInstanceWithContext(context.Background(), request)
}

// ModifyPublicNetworkInstance
// Modify public network instance info
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_PUBLICINSTANCENOTFOUND = "ResourceNotFound.PublicInstanceNotFound"
func (c *Client) ModifyPublicNetworkInstanceWithContext(ctx context.Context, request *ModifyPublicNetworkInstanceRequest) (response *ModifyPublicNetworkInstanceResponse, err error) {
    if request == nil {
        request = NewModifyPublicNetworkInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "edgezone", APIVersion, "ModifyPublicNetworkInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPublicNetworkInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPublicNetworkInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewReleasePublicIpRequest() (request *ReleasePublicIpRequest) {
    request = &ReleasePublicIpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("edgezone", APIVersion, "ReleasePublicIp")
    
    
    return
}

func NewReleasePublicIpResponse() (response *ReleasePublicIpResponse) {
    response = &ReleasePublicIpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReleasePublicIp
// Batch release IPv4 addresses assigned to STATIC public network instances but not bound to physical servers
//
// This API is applicable only to STATIC mode instances. The CIDR of BGP/OSPF instances is automatically returned during deletion with no need to manually release single IP addresses.
//
// error code that may be returned:
//  FAILEDOPERATION_IPSTILLBOUNDTOSERVER = "FailedOperation.IpStillBoundToServer"
//  FAILEDOPERATION_NOTSUPPORTEDFORDYNAMICINSTANCE = "FailedOperation.NotSupportedForDynamicInstance"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_PUBLICIPNOTAVAILABLE = "InvalidParameterValue.PublicIpNotAvailable"
//  INVALIDPARAMETERVALUE_PUBLICIPV6NOTAVAILABLE = "InvalidParameterValue.PublicIpv6NotAvailable"
//  RESOURCENOTFOUND_PUBLICINSTANCENOTFOUND = "ResourceNotFound.PublicInstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ReleasePublicIp(request *ReleasePublicIpRequest) (response *ReleasePublicIpResponse, err error) {
    return c.ReleasePublicIpWithContext(context.Background(), request)
}

// ReleasePublicIp
// Batch release IPv4 addresses assigned to STATIC public network instances but not bound to physical servers
//
// This API is applicable only to STATIC mode instances. The CIDR of BGP/OSPF instances is automatically returned during deletion with no need to manually release single IP addresses.
//
// error code that may be returned:
//  FAILEDOPERATION_IPSTILLBOUNDTOSERVER = "FailedOperation.IpStillBoundToServer"
//  FAILEDOPERATION_NOTSUPPORTEDFORDYNAMICINSTANCE = "FailedOperation.NotSupportedForDynamicInstance"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_PUBLICIPNOTAVAILABLE = "InvalidParameterValue.PublicIpNotAvailable"
//  INVALIDPARAMETERVALUE_PUBLICIPV6NOTAVAILABLE = "InvalidParameterValue.PublicIpv6NotAvailable"
//  RESOURCENOTFOUND_PUBLICINSTANCENOTFOUND = "ResourceNotFound.PublicInstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ReleasePublicIpWithContext(ctx context.Context, request *ReleasePublicIpRequest) (response *ReleasePublicIpResponse, err error) {
    if request == nil {
        request = NewReleasePublicIpRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "edgezone", APIVersion, "ReleasePublicIp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReleasePublicIp require credential")
    }

    request.SetContext(ctx)
    
    response = NewReleasePublicIpResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateInstancesRequest() (request *TerminateInstancesRequest) {
    request = &TerminateInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("edgezone", APIVersion, "TerminateInstances")
    
    
    return
}

func NewTerminateInstancesResponse() (response *TerminateInstancesResponse) {
    response = &TerminateInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TerminateInstances
// This API is used to terminate a physical machine instance and free up resources. It synchronously releases network resources (IP recycling) and updates status to terminating, while performing disk cleanup asynchronously in the background. It supports partially successful operations.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
func (c *Client) TerminateInstances(request *TerminateInstancesRequest) (response *TerminateInstancesResponse, err error) {
    return c.TerminateInstancesWithContext(context.Background(), request)
}

// TerminateInstances
// This API is used to terminate a physical machine instance and free up resources. It synchronously releases network resources (IP recycling) and updates status to terminating, while performing disk cleanup asynchronously in the background. It supports partially successful operations.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
func (c *Client) TerminateInstancesWithContext(ctx context.Context, request *TerminateInstancesRequest) (response *TerminateInstancesResponse, err error) {
    if request == nil {
        request = NewTerminateInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "edgezone", APIVersion, "TerminateInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateInstancesResponse()
    err = c.Send(request, response)
    return
}
