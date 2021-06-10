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

package v20170312

import (
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2017-03-12"

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


func NewAcceptAttachCcnInstancesRequest() (request *AcceptAttachCcnInstancesRequest) {
    request = &AcceptAttachCcnInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "AcceptAttachCcnInstances")
    return
}

func NewAcceptAttachCcnInstancesResponse() (response *AcceptAttachCcnInstancesResponse) {
    response = &AcceptAttachCcnInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AcceptAttachCcnInstances
// This API (AcceptAttachCcnInstances) is used to associate instances across accounts. Cloud Connect Network (CCN) owners accept and agree to the operations.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CCNNOTATTACHED = "UnsupportedOperation.CcnNotAttached"
//  UNSUPPORTEDOPERATION_ISNOTFINANCEACCOUNT = "UnsupportedOperation.IsNotFinanceAccount"
//  UNSUPPORTEDOPERATION_NOTPENDINGCCNINSTANCE = "UnsupportedOperation.NotPendingCcnInstance"
//  UNSUPPORTEDOPERATION_UNABLECROSSFINANCE = "UnsupportedOperation.UnableCrossFinance"
func (c *Client) AcceptAttachCcnInstances(request *AcceptAttachCcnInstancesRequest) (response *AcceptAttachCcnInstancesResponse, err error) {
    if request == nil {
        request = NewAcceptAttachCcnInstancesRequest()
    }
    response = NewAcceptAttachCcnInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewAddBandwidthPackageResourcesRequest() (request *AddBandwidthPackageResourcesRequest) {
    request = &AddBandwidthPackageResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "AddBandwidthPackageResources")
    return
}

func NewAddBandwidthPackageResourcesResponse() (response *AddBandwidthPackageResourcesResponse) {
    response = &AddBandwidthPackageResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddBandwidthPackageResources
// This API is used to add resources to a bandwidth package, including [Elastic IP](https://intl.cloud.tencent.com/document/product/213/1941?from_cn_redirect=1), [Cloud Load Balancer](https://intl.cloud.tencent.com/document/product/214/517?from_cn_redirect=1), and so on.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) AddBandwidthPackageResources(request *AddBandwidthPackageResourcesRequest) (response *AddBandwidthPackageResourcesResponse, err error) {
    if request == nil {
        request = NewAddBandwidthPackageResourcesRequest()
    }
    response = NewAddBandwidthPackageResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewAllocateAddressesRequest() (request *AllocateAddressesRequest) {
    request = &AllocateAddressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "AllocateAddresses")
    return
}

func NewAllocateAddressesResponse() (response *AllocateAddressesResponse) {
    response = &AllocateAddressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AllocateAddresses
// This API is used to apply for one or more [Elastic IP Addresses](https://intl.cloud.tencent.com/document/product/213/1941?from_cn_redirect=1) (EIPs for short).
//
// * An EIP is a static IP address that is dedicated for dynamic cloud computing. You can quickly re-map an EIP to another instance under your account to protect against instance failures.
//
// * Your EIP is associated with your Tencent Cloud account rather than an instance. It remains associated with your Tencent Cloud account until you choose to explicitly release it or your account is in arrears for more than 24 hours.
//
// * The maximum number of EIPs that can be applied for a Tencent Cloud account in each region is restricted. For more information, see [EIP Product Introduction](https://intl.cloud.tencent.com/document/product/213/5733?from_cn_redirect=1). You can get the quota information through the DescribeAddressQuota API.
//
// error code that may be returned:
//  ADDRESSQUOTALIMITEXCEEDED = "AddressQuotaLimitExceeded"
//  ADDRESSQUOTALIMITEXCEEDED_DAILYALLOCATE = "AddressQuotaLimitExceeded.DailyAllocate"
//  INVALIDACCOUNT_NOTSUPPORTED = "InvalidAccount.NotSupported"
//  INVALIDADDRESSID_BLOCKED = "InvalidAddressId.Blocked"
//  INVALIDPARAMETERCONFLICT = "InvalidParameterConflict"
//  INVALIDPARAMETERVALUE_BANDWIDTHOUTOFRANGE = "InvalidParameterValue.BandwidthOutOfRange"
//  INVALIDPARAMETERVALUE_BANDWIDTHTOOSMALL = "InvalidParameterValue.BandwidthTooSmall"
//  INVALIDPARAMETERVALUE_INVALIDDEDICATEDCLUSTERID = "InvalidParameterValue.InvalidDedicatedClusterId"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) AllocateAddresses(request *AllocateAddressesRequest) (response *AllocateAddressesResponse, err error) {
    if request == nil {
        request = NewAllocateAddressesRequest()
    }
    response = NewAllocateAddressesResponse()
    err = c.Send(request, response)
    return
}

func NewAssignIpv6AddressesRequest() (request *AssignIpv6AddressesRequest) {
    request = &AssignIpv6AddressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "AssignIpv6Addresses")
    return
}

func NewAssignIpv6AddressesResponse() (response *AssignIpv6AddressesResponse) {
    response = &AssignIpv6AddressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssignIpv6Addresses
// This API (AssignIpv6Addresses) is used to apply for an IPv6 address for the ENI.<br />
//
// This API is completed asynchronously. If you need to query the async execution results, use the `RequestId` returned by this API to query the `QueryTask` API.
//
// * An ENI can only be bound with a limited number of IPs. For more information about resource limits, see<a href="/document/product/576/18527">ENI use limits</a>.
//
// * You can specify the `IPv6` address when applying. The address type cannot be the primary `IP`. Currently, `IPv6` can only be supported as the secondary `IP`.
//
// * The address must be unoccupied and is in the subnet to which the ENI belongs.
//
// * When applying for one to multiple secondary `IPv6` addresses on ENI, the API will return the specified number of secondary `IPv6` addresses in the subnet range where the ENI is located.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED_ADDRESS = "LimitExceeded.Address"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE_ADDRESS = "ResourceInUse.Address"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
//  UNSUPPORTEDOPERATION_UNASSIGNCIDRBLOCK = "UnsupportedOperation.UnassignCidrBlock"
func (c *Client) AssignIpv6Addresses(request *AssignIpv6AddressesRequest) (response *AssignIpv6AddressesResponse, err error) {
    if request == nil {
        request = NewAssignIpv6AddressesRequest()
    }
    response = NewAssignIpv6AddressesResponse()
    err = c.Send(request, response)
    return
}

func NewAssignIpv6CidrBlockRequest() (request *AssignIpv6CidrBlockRequest) {
    request = &AssignIpv6CidrBlockRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "AssignIpv6CidrBlock")
    return
}

func NewAssignIpv6CidrBlockResponse() (response *AssignIpv6CidrBlockResponse) {
    response = &AssignIpv6CidrBlockResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssignIpv6CidrBlock
// This API is used to assign IPv6 ranges.
//
// * To use this API, you must already have a VPC instance. If you do not have a VPC instance yet, use the <a href="https://intl.cloud.tencent.com/document/api/215/15774?from_cn_redirect=1" title="CreateVpc" target="_blank">CreateVpc</a> API to create one.
//
// * Each VPC can apply for only one IPv6 range.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED_CIDRBLOCK = "LimitExceeded.CidrBlock"
//  RESOURCEINSUFFICIENT_CIDRBLOCK = "ResourceInsufficient.CidrBlock"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AssignIpv6CidrBlock(request *AssignIpv6CidrBlockRequest) (response *AssignIpv6CidrBlockResponse, err error) {
    if request == nil {
        request = NewAssignIpv6CidrBlockRequest()
    }
    response = NewAssignIpv6CidrBlockResponse()
    err = c.Send(request, response)
    return
}

func NewAssignIpv6SubnetCidrBlockRequest() (request *AssignIpv6SubnetCidrBlockRequest) {
    request = &AssignIpv6SubnetCidrBlockRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "AssignIpv6SubnetCidrBlock")
    return
}

func NewAssignIpv6SubnetCidrBlockResponse() (response *AssignIpv6SubnetCidrBlockResponse) {
    response = &AssignIpv6SubnetCidrBlockResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssignIpv6SubnetCidrBlock
// This API (AssignIpv6SubnetCidrBlock) is used to assign IPv6 subnet IP ranges.
//
// * To assign an `IPv6` IP range to a subnet, the `VPC` that the subnet belongs to should have obtained the `IPv6` IP range. If this has not been assigned, use the `AssignIpv6CidrBlock` API to assign an `IPv6` IP range to the `VPC` to which the subnet belongs. Otherwise, the `IPv6` subnet IP range cannot be assigned.
//
// * Each subnet can only be assigned one IPv6 IP range.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_SUBNETCONFLICT = "InvalidParameterValue.SubnetConflict"
//  INVALIDPARAMETERVALUE_SUBNETRANGE = "InvalidParameterValue.SubnetRange"
//  LIMITEXCEEDED_SUBNETCIDRBLOCK = "LimitExceeded.SubnetCidrBlock"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AssignIpv6SubnetCidrBlock(request *AssignIpv6SubnetCidrBlockRequest) (response *AssignIpv6SubnetCidrBlockResponse, err error) {
    if request == nil {
        request = NewAssignIpv6SubnetCidrBlockRequest()
    }
    response = NewAssignIpv6SubnetCidrBlockResponse()
    err = c.Send(request, response)
    return
}

func NewAssignPrivateIpAddressesRequest() (request *AssignPrivateIpAddressesRequest) {
    request = &AssignPrivateIpAddressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "AssignPrivateIpAddresses")
    return
}

func NewAssignPrivateIpAddressesResponse() (response *AssignPrivateIpAddressesResponse) {
    response = &AssignPrivateIpAddressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssignPrivateIpAddresses
// This API (AssignPrivateIpAddresses) is used for the ENI to apply for private IPs.
//
// * An ENI can only be bound with a limited number of IP addresses. For more information about resource limits, see <a href="/document/product/576/18527">ENI Use Limits</a>.
//
// * You can specify the private IP you want to apply for. It cannot be the primary IP, which already exists and cannot be modified. The private IP must be in the same subnet as the ENI, and cannot be occupied.
//
// * You can apply for more than one secondary private IP on the ENI. The API will return the specified number of secondary private IPs in the subnet IP range of the ENI.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_RESERVED = "InvalidParameterValue.Reserved"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) AssignPrivateIpAddresses(request *AssignPrivateIpAddressesRequest) (response *AssignPrivateIpAddressesResponse, err error) {
    if request == nil {
        request = NewAssignPrivateIpAddressesRequest()
    }
    response = NewAssignPrivateIpAddressesResponse()
    err = c.Send(request, response)
    return
}

func NewAssociateAddressRequest() (request *AssociateAddressRequest) {
    request = &AssociateAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "AssociateAddress")
    return
}

func NewAssociateAddressResponse() (response *AssociateAddressResponse) {
    response = &AssociateAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssociateAddress
// This API is used to bind an [Elastic IP](https://intl.cloud.tencent.com/document/product/213/1941?from_cn_redirect=1) (EIP for short) to the specified private IP of an instance or ENI.
//
// * The EIP is essentially bound to the primary private IP of the primary ENI on a CVM instance.
//
// * The EIP binding will automatically unbind and release the public IP previously bound to the CVM instance.
//
// * To bind another EIP to the private IP of the specified ENI (not the primary private IP of the primary ENI), you must first unbind the EIP.
//
// * To bind an EIP to a NAT Gateway, use the [`AssociateNatGatewayAddress`](https://intl.cloud.tencent.com/document/product/215/36722?from_cn_redirect=1) API.
//
// * EIP that is in arrears or blocked cannot be bound.
//
// * Only EIP in the UNBIND status can be bound.
//
// error code that may be returned:
//  INVALIDACCOUNT_NOTSUPPORTED = "InvalidAccount.NotSupported"
//  INVALIDADDRESSID_BLOCKED = "InvalidAddressId.Blocked"
//  INVALIDADDRESSID_NOTFOUND = "InvalidAddressId.NotFound"
//  INVALIDADDRESSIDSTATE_INARREARS = "InvalidAddressIdState.InArrears"
//  INVALIDADDRESSIDSTATUS_NOTPERMIT = "InvalidAddressIdStatus.NotPermit"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_ALREADYBINDEIP = "InvalidInstanceId.AlreadyBindEip"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDNETWORKINTERFACEID_NOTFOUND = "InvalidNetworkInterfaceId.NotFound"
//  INVALIDPARAMETERCONFLICT = "InvalidParameterConflict"
//  INVALIDPARAMETERVALUE_COMBINATION = "InvalidParameterValue.Combination"
//  INVALIDPRIVATEIPADDRESS_ALREADYBINDEIP = "InvalidPrivateIpAddress.AlreadyBindEip"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGRESOURCEID = "MissingResourceId"
func (c *Client) AssociateAddress(request *AssociateAddressRequest) (response *AssociateAddressResponse, err error) {
    if request == nil {
        request = NewAssociateAddressRequest()
    }
    response = NewAssociateAddressResponse()
    err = c.Send(request, response)
    return
}

func NewAssociateDirectConnectGatewayNatGatewayRequest() (request *AssociateDirectConnectGatewayNatGatewayRequest) {
    request = &AssociateDirectConnectGatewayNatGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "AssociateDirectConnectGatewayNatGateway")
    return
}

func NewAssociateDirectConnectGatewayNatGatewayResponse() (response *AssociateDirectConnectGatewayNatGatewayResponse) {
    response = &AssociateDirectConnectGatewayNatGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssociateDirectConnectGatewayNatGateway
// This API is used to bind a direct connect gateway with a NAT gateway,  and direct its default route to the NAT Gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AssociateDirectConnectGatewayNatGateway(request *AssociateDirectConnectGatewayNatGatewayRequest) (response *AssociateDirectConnectGatewayNatGatewayResponse, err error) {
    if request == nil {
        request = NewAssociateDirectConnectGatewayNatGatewayRequest()
    }
    response = NewAssociateDirectConnectGatewayNatGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewAssociateNatGatewayAddressRequest() (request *AssociateNatGatewayAddressRequest) {
    request = &AssociateNatGatewayAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "AssociateNatGatewayAddress")
    return
}

func NewAssociateNatGatewayAddressResponse() (response *AssociateNatGatewayAddressResponse) {
    response = &AssociateNatGatewayAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssociateNatGatewayAddress
// This API is used to bind an EIP to NAT Gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.AddressQuotaLimitExceeded"
//  LIMITEXCEEDED_DAILYALLOCATEADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.DailyAllocateAddressQuotaLimitExceeded"
//  LIMITEXCEEDED_PUBLICIPADDRESSPERNATGATEWAYLIMITEXCEEDED = "LimitExceeded.PublicIpAddressPerNatGatewayLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
//  UNSUPPORTEDOPERATION_PUBLICIPADDRESSISNOTBGPIP = "UnsupportedOperation.PublicIpAddressIsNotBGPIp"
//  UNSUPPORTEDOPERATION_PUBLICIPADDRESSNOTBILLEDBYTRAFFIC = "UnsupportedOperation.PublicIpAddressNotBilledByTraffic"
func (c *Client) AssociateNatGatewayAddress(request *AssociateNatGatewayAddressRequest) (response *AssociateNatGatewayAddressResponse, err error) {
    if request == nil {
        request = NewAssociateNatGatewayAddressRequest()
    }
    response = NewAssociateNatGatewayAddressResponse()
    err = c.Send(request, response)
    return
}

func NewAssociateNetworkAclSubnetsRequest() (request *AssociateNetworkAclSubnetsRequest) {
    request = &AssociateNetworkAclSubnetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "AssociateNetworkAclSubnets")
    return
}

func NewAssociateNetworkAclSubnetsResponse() (response *AssociateNetworkAclSubnetsResponse) {
    response = &AssociateNetworkAclSubnetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssociateNetworkAclSubnets
// This API is used to associate a network ACL with subnets in a VPC instance.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
func (c *Client) AssociateNetworkAclSubnets(request *AssociateNetworkAclSubnetsRequest) (response *AssociateNetworkAclSubnetsResponse, err error) {
    if request == nil {
        request = NewAssociateNetworkAclSubnetsRequest()
    }
    response = NewAssociateNetworkAclSubnetsResponse()
    err = c.Send(request, response)
    return
}

func NewAssociateNetworkInterfaceSecurityGroupsRequest() (request *AssociateNetworkInterfaceSecurityGroupsRequest) {
    request = &AssociateNetworkInterfaceSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "AssociateNetworkInterfaceSecurityGroups")
    return
}

func NewAssociateNetworkInterfaceSecurityGroupsResponse() (response *AssociateNetworkInterfaceSecurityGroupsResponse) {
    response = &AssociateNetworkInterfaceSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssociateNetworkInterfaceSecurityGroups
// This API (AssociateNetworkInterfaceSecurityGroups) is used to attach a security group to an ENI.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AssociateNetworkInterfaceSecurityGroups(request *AssociateNetworkInterfaceSecurityGroupsRequest) (response *AssociateNetworkInterfaceSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewAssociateNetworkInterfaceSecurityGroupsRequest()
    }
    response = NewAssociateNetworkInterfaceSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewAttachCcnInstancesRequest() (request *AttachCcnInstancesRequest) {
    request = &AttachCcnInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "AttachCcnInstances")
    return
}

func NewAttachCcnInstancesResponse() (response *AttachCcnInstancesResponse) {
    response = &AttachCcnInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AttachCcnInstances
// This API (AttachCcnInstances) is used to load a network instance to a CCN instance. Network instances include VPCs and Direct Connect gateways.<br />
//
// The number of network instances that each CCN can be associated with is limited. For more information, see the product documentation. If you need to associate more instances, please contact online customer service.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CCNATTACHBMVPCLIMITEXCEEDED = "InvalidParameterValue.CcnAttachBmvpcLimitExceeded"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CCNATTACHED = "UnsupportedOperation.CcnAttached"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_ISNOTFINANCEACCOUNT = "UnsupportedOperation.IsNotFinanceAccount"
//  UNSUPPORTEDOPERATION_UINNOTFOUND = "UnsupportedOperation.UinNotFound"
//  UNSUPPORTEDOPERATION_UNABLECROSSBORDER = "UnsupportedOperation.UnableCrossBorder"
//  UNSUPPORTEDOPERATION_UNABLECROSSFINANCE = "UnsupportedOperation.UnableCrossFinance"
func (c *Client) AttachCcnInstances(request *AttachCcnInstancesRequest) (response *AttachCcnInstancesResponse, err error) {
    if request == nil {
        request = NewAttachCcnInstancesRequest()
    }
    response = NewAttachCcnInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewAttachClassicLinkVpcRequest() (request *AttachClassicLinkVpcRequest) {
    request = &AttachClassicLinkVpcRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "AttachClassicLinkVpc")
    return
}

func NewAttachClassicLinkVpcResponse() (response *AttachClassicLinkVpcResponse) {
    response = &AttachClassicLinkVpcResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AttachClassicLinkVpc
// This API is used to create a Classiclink between a VPC instance and a basic network device.
//
// * The VPC instance and the basic network device must be in the same region.
//
// * For differences between VPC and basic networks, see <a href="https://intl.cloud.tencent.com/document/product/215/30720?from_cn_redirect=1">VPC and Basic Networks</a>.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CIDRUNSUPPORTEDCLASSICLINK = "UnsupportedOperation.CIDRUnSupportedClassicLink"
//  UNSUPPORTEDOPERATION_CLASSICINSTANCEIDALREADYEXISTS = "UnsupportedOperation.ClassicInstanceIdAlreadyExists"
func (c *Client) AttachClassicLinkVpc(request *AttachClassicLinkVpcRequest) (response *AttachClassicLinkVpcResponse, err error) {
    if request == nil {
        request = NewAttachClassicLinkVpcRequest()
    }
    response = NewAttachClassicLinkVpcResponse()
    err = c.Send(request, response)
    return
}

func NewAttachNetworkInterfaceRequest() (request *AttachNetworkInterfaceRequest) {
    request = &AttachNetworkInterfaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "AttachNetworkInterface")
    return
}

func NewAttachNetworkInterfaceResponse() (response *AttachNetworkInterfaceResponse) {
    response = &AttachNetworkInterfaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AttachNetworkInterface
// This API is used to bind an ENI to a CVM.
//
// * One CVM can be bound with multiple ENIs, but only one primary ENI. For more information about the limits, please see <a href="https://intl.cloud.tencent.com/document/product/576/18527?from_cn_redirect=1">ENI Use Limits</a>.
//
// * An ENI can only be bound to one CVM at a time.
//
// * Only the running or shutdown CVMs can be bound with ENIs. For more information about the CVM status, see <a href="https://intl.cloud.tencent.com/document/api/213/9452?from_cn_redirect=1#InstanceStatus">InstanceStatus</a> in the Data Types.
//
// * An ENI can only be bound to a VPC-based CVM under the same availability zone as the ENI subnet.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ATTACHMENTALREADYEXISTS = "UnsupportedOperation.AttachmentAlreadyExists"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDINSTANCEFAMILY = "UnsupportedOperation.UnsupportedInstanceFamily"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
//  UNSUPPORTEDOPERATION_ZONEMISMATCH = "UnsupportedOperation.ZoneMismatch"
func (c *Client) AttachNetworkInterface(request *AttachNetworkInterfaceRequest) (response *AttachNetworkInterfaceResponse, err error) {
    if request == nil {
        request = NewAttachNetworkInterfaceRequest()
    }
    response = NewAttachNetworkInterfaceResponse()
    err = c.Send(request, response)
    return
}

func NewAuditCrossBorderComplianceRequest() (request *AuditCrossBorderComplianceRequest) {
    request = &AuditCrossBorderComplianceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "AuditCrossBorderCompliance")
    return
}

func NewAuditCrossBorderComplianceResponse() (response *AuditCrossBorderComplianceResponse) {
    response = &AuditCrossBorderComplianceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AuditCrossBorderCompliance
// This API is used by the service provider to perform a compliance audit.
//
// * This API is only provided for service providers to audit compliance review requests received. Tencent Cloud will verify the identity of the service provider by the `APPID`. 
//
// * The status of the review request can be changed between `APPROVED` and `DENY`.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AuditCrossBorderCompliance(request *AuditCrossBorderComplianceRequest) (response *AuditCrossBorderComplianceResponse, err error) {
    if request == nil {
        request = NewAuditCrossBorderComplianceRequest()
    }
    response = NewAuditCrossBorderComplianceResponse()
    err = c.Send(request, response)
    return
}

func NewCheckAssistantCidrRequest() (request *CheckAssistantCidrRequest) {
    request = &CheckAssistantCidrRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CheckAssistantCidr")
    return
}

func NewCheckAssistantCidrResponse() (response *CheckAssistantCidrResponse) {
    response = &CheckAssistantCidrResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckAssistantCidr
// This API (CheckAssistantCidr) is used to check overlapping of a secondary CIDR block with inventory routing, peering connection (opposite VPC CIDR block), and any other resources. If an overlap is present, the overlapped resources are returned. (To use this API that is in Beta, please submit a ticket.)
//
// * Check whether the secondary CIDR block overlaps with a primary or secondary CIDR block of the current VPC.
//
// * Check whether the secondary CIDR block overlaps with the routing destination of the current VPC.
//
// * Check whether the secondary CIDR block is peer-connected to the current VPC, and whether it overlaps with a main or secondary CIDR block of the opposite VPC.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_SUBNETCONFLICT = "InvalidParameterValue.SubnetConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CheckAssistantCidr(request *CheckAssistantCidrRequest) (response *CheckAssistantCidrResponse, err error) {
    if request == nil {
        request = NewCheckAssistantCidrRequest()
    }
    response = NewCheckAssistantCidrResponse()
    err = c.Send(request, response)
    return
}

func NewCheckNetDetectStateRequest() (request *CheckNetDetectStateRequest) {
    request = &CheckNetDetectStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CheckNetDetectState")
    return
}

func NewCheckNetDetectStateResponse() (response *CheckNetDetectStateResponse) {
    response = &CheckNetDetectStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckNetDetectState
// This API is used to verify the network detection status.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_NEXTHOPMISMATCH = "InvalidParameter.NextHopMismatch"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NETDETECTSAMEIP = "InvalidParameterValue.NetDetectSameIp"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_CONFLICTWITHDOCKERROUTE = "UnsupportedOperation.ConflictWithDockerRoute"
//  UNSUPPORTEDOPERATION_ECMPWITHUSERROUTE = "UnsupportedOperation.EcmpWithUserRoute"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
func (c *Client) CheckNetDetectState(request *CheckNetDetectStateRequest) (response *CheckNetDetectStateResponse, err error) {
    if request == nil {
        request = NewCheckNetDetectStateRequest()
    }
    response = NewCheckNetDetectStateResponse()
    err = c.Send(request, response)
    return
}

func NewCloneSecurityGroupRequest() (request *CloneSecurityGroupRequest) {
    request = &CloneSecurityGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CloneSecurityGroup")
    return
}

func NewCloneSecurityGroupResponse() (response *CloneSecurityGroupResponse) {
    response = &CloneSecurityGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CloneSecurityGroup
// This API is used to create a security group with the same rule configurations as an existing security group. The cloning only copies the security group and its rules, but not the security group tags.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CloneSecurityGroup(request *CloneSecurityGroupRequest) (response *CloneSecurityGroupResponse, err error) {
    if request == nil {
        request = NewCloneSecurityGroupRequest()
    }
    response = NewCloneSecurityGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAddressTemplateRequest() (request *CreateAddressTemplateRequest) {
    request = &CreateAddressTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateAddressTemplate")
    return
}

func NewCreateAddressTemplateResponse() (response *CreateAddressTemplateResponse) {
    response = &CreateAddressTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAddressTemplate
// This API (CreateAddressTemplate) is used to create an IP address template.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateAddressTemplate(request *CreateAddressTemplateRequest) (response *CreateAddressTemplateResponse, err error) {
    if request == nil {
        request = NewCreateAddressTemplateRequest()
    }
    response = NewCreateAddressTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAddressTemplateGroupRequest() (request *CreateAddressTemplateGroupRequest) {
    request = &CreateAddressTemplateGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateAddressTemplateGroup")
    return
}

func NewCreateAddressTemplateGroupResponse() (response *CreateAddressTemplateGroupResponse) {
    response = &CreateAddressTemplateGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAddressTemplateGroup
// This API (CreateAddressTemplateGroup) is used to create an IP address template group.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) CreateAddressTemplateGroup(request *CreateAddressTemplateGroupRequest) (response *CreateAddressTemplateGroupResponse, err error) {
    if request == nil {
        request = NewCreateAddressTemplateGroupRequest()
    }
    response = NewCreateAddressTemplateGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAndAttachNetworkInterfaceRequest() (request *CreateAndAttachNetworkInterfaceRequest) {
    request = &CreateAndAttachNetworkInterfaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateAndAttachNetworkInterface")
    return
}

func NewCreateAndAttachNetworkInterfaceResponse() (response *CreateAndAttachNetworkInterfaceResponse) {
    response = &CreateAndAttachNetworkInterfaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAndAttachNetworkInterface
// This API is used to create an ENI and bind it to a CVM.
//
// * You can specify private IP addresses and a primary IP when creating an ENI. The specified private IP must be idle and in the same subnet as the ENI.
//
// * When creating an ENI, you can specify the number of private IPs that you want to apply for. The system will randomly generate private IP addresses.
//
// * The number of IPs bound with an ENI is limited. For more information, see <a href="/document/product/576/18527">ENI Use Limits</a>.
//
// * You can bind an existing security group when creating an ENI.
//
// * You can bind a tag when creating an ENI. The tag list in the response indicates the tags that have been successfully added.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_RESERVED = "InvalidParameterValue.Reserved"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDINSTANCEFAMILY = "UnsupportedOperation.UnsupportedInstanceFamily"
func (c *Client) CreateAndAttachNetworkInterface(request *CreateAndAttachNetworkInterfaceRequest) (response *CreateAndAttachNetworkInterfaceResponse, err error) {
    if request == nil {
        request = NewCreateAndAttachNetworkInterfaceRequest()
    }
    response = NewCreateAndAttachNetworkInterfaceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAssistantCidrRequest() (request *CreateAssistantCidrRequest) {
    request = &CreateAssistantCidrRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateAssistantCidr")
    return
}

func NewCreateAssistantCidrResponse() (response *CreateAssistantCidrResponse) {
    response = &CreateAssistantCidrResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAssistantCidr
// This API (CreateAssistantCidr) is used to batch create secondary CIDR blocks. (To use this API that is in Beta, please submit a ticket.)
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_SUBNETCONFLICT = "InvalidParameterValue.SubnetConflict"
//  INVALIDPARAMETERVALUE_SUBNETRANGE = "InvalidParameterValue.SubnetRange"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateAssistantCidr(request *CreateAssistantCidrRequest) (response *CreateAssistantCidrResponse, err error) {
    if request == nil {
        request = NewCreateAssistantCidrRequest()
    }
    response = NewCreateAssistantCidrResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBandwidthPackageRequest() (request *CreateBandwidthPackageRequest) {
    request = &CreateBandwidthPackageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateBandwidthPackage")
    return
}

func NewCreateBandwidthPackageResponse() (response *CreateBandwidthPackageResponse) {
    response = &CreateBandwidthPackageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBandwidthPackage
// This API is used to create [device bandwidth packages](https://intl.cloud.tencent.com/document/product/684/15246?from_cn_redirect=1#.E8.AE.BE.E5.A4.87.E5.B8.A6.E5.AE.BD.E5.8C.85) and [IP bandwidth packages](https://intl.cloud.tencent.com/document/product/684/15246?from_cn_redirect=1#ip-.E5.B8.A6.E5.AE.BD.E5.8C.85)
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDACCOUNT_NOTSUPPORTED = "InvalidAccount.NotSupported"
func (c *Client) CreateBandwidthPackage(request *CreateBandwidthPackageRequest) (response *CreateBandwidthPackageResponse, err error) {
    if request == nil {
        request = NewCreateBandwidthPackageRequest()
    }
    response = NewCreateBandwidthPackageResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCcnRequest() (request *CreateCcnRequest) {
    request = &CreateCcnRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateCcn")
    return
}

func NewCreateCcnResponse() (response *CreateCcnResponse) {
    response = &CreateCcnResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCcn
// This API is used to create a Cloud Connect Network (CCN).<br />
//
// * You can bind a tag when creating a CCN instance. The tag list in the response indicates the tags that have been successfully added.
//
// Each account can only create a limited number of CCN instances. For more information, see product documentation. To create more instances, contact the online customer service.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_NOREALNAMEAUTHENTICATION = "UnauthorizedOperation.NoRealNameAuthentication"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INSUFFICIENTFUNDS = "UnsupportedOperation.InsufficientFunds"
//  UNSUPPORTEDOPERATION_PREPAIDCCNONLYSUPPORTINTERREGIONLIMIT = "UnsupportedOperation.PrepaidCcnOnlySupportInterRegionLimit"
//  UNSUPPORTEDOPERATION_USERANDCCNCHARGETYPENOTMATCH = "UnsupportedOperation.UserAndCcnChargeTypeNotMatch"
func (c *Client) CreateCcn(request *CreateCcnRequest) (response *CreateCcnResponse, err error) {
    if request == nil {
        request = NewCreateCcnRequest()
    }
    response = NewCreateCcnResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCustomerGatewayRequest() (request *CreateCustomerGatewayRequest) {
    request = &CreateCustomerGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateCustomerGateway")
    return
}

func NewCreateCustomerGatewayResponse() (response *CreateCustomerGatewayResponse) {
    response = &CreateCustomerGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCustomerGateway
// This API (CreateCustomerGateway) is used to create customer gateways.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  LIMITEXCEEDED = "LimitExceeded"
//  VPCLIMITEXCEEDED = "VpcLimitExceeded"
func (c *Client) CreateCustomerGateway(request *CreateCustomerGatewayRequest) (response *CreateCustomerGatewayResponse, err error) {
    if request == nil {
        request = NewCreateCustomerGatewayRequest()
    }
    response = NewCreateCustomerGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDefaultVpcRequest() (request *CreateDefaultVpcRequest) {
    request = &CreateDefaultVpcRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateDefaultVpc")
    return
}

func NewCreateDefaultVpcResponse() (response *CreateDefaultVpcResponse) {
    response = &CreateDefaultVpcResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDefaultVpc
// This API is used to create a default VPC.
//
// 
//
// The default VPC is suitable for getting started with and launching public instances, and it can be used like any other VPCs. To create a standard VPC, for which you need to specify a VPC name, VPC IP range, subnet IP range, and subnet availability zone, use the regular CreateVpc API.
//
// 
//
// Under normal circumstances, this API may not create a default VPC. It depends on the network attributes (DescribeAccountAttributes) of your account.
//
// * If both basic network and VPC are supported, the returned VpcId is 0.
//
// * If only VPC is supported, the default VPC information is returned.
//
// 
//
// You can also use the Force parameter to forcibly return a default VPC.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT_CIDRBLOCK = "ResourceInsufficient.CidrBlock"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateDefaultVpc(request *CreateDefaultVpcRequest) (response *CreateDefaultVpcResponse, err error) {
    if request == nil {
        request = NewCreateDefaultVpcRequest()
    }
    response = NewCreateDefaultVpcResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDirectConnectGatewayRequest() (request *CreateDirectConnectGatewayRequest) {
    request = &CreateDirectConnectGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateDirectConnectGateway")
    return
}

func NewCreateDirectConnectGatewayResponse() (response *CreateDirectConnectGatewayResponse) {
    response = &CreateDirectConnectGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDirectConnectGateway
// This API is used to create a direct connect gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_UNABLECROSSBORDER = "UnsupportedOperation.UnableCrossBorder"
func (c *Client) CreateDirectConnectGateway(request *CreateDirectConnectGatewayRequest) (response *CreateDirectConnectGatewayResponse, err error) {
    if request == nil {
        request = NewCreateDirectConnectGatewayRequest()
    }
    response = NewCreateDirectConnectGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDirectConnectGatewayCcnRoutesRequest() (request *CreateDirectConnectGatewayCcnRoutesRequest) {
    request = &CreateDirectConnectGatewayCcnRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateDirectConnectGatewayCcnRoutes")
    return
}

func NewCreateDirectConnectGatewayCcnRoutesResponse() (response *CreateDirectConnectGatewayCcnRoutesResponse) {
    response = &CreateDirectConnectGatewayCcnRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDirectConnectGatewayCcnRoutes
// This API (CreateDirectConnectGatewayCcnRoutes) is used to create the CCN route (IDC IP range) of a Direct Connect gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
func (c *Client) CreateDirectConnectGatewayCcnRoutes(request *CreateDirectConnectGatewayCcnRoutesRequest) (response *CreateDirectConnectGatewayCcnRoutesResponse, err error) {
    if request == nil {
        request = NewCreateDirectConnectGatewayCcnRoutesRequest()
    }
    response = NewCreateDirectConnectGatewayCcnRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFlowLogRequest() (request *CreateFlowLogRequest) {
    request = &CreateFlowLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateFlowLog")
    return
}

func NewCreateFlowLogResponse() (response *CreateFlowLogResponse) {
    response = &CreateFlowLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateFlowLog
// This API is used to create a flow log.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateFlowLog(request *CreateFlowLogRequest) (response *CreateFlowLogResponse, err error) {
    if request == nil {
        request = NewCreateFlowLogRequest()
    }
    response = NewCreateFlowLogResponse()
    err = c.Send(request, response)
    return
}

func NewCreateHaVipRequest() (request *CreateHaVipRequest) {
    request = &CreateHaVipRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateHaVip")
    return
}

func NewCreateHaVipResponse() (response *CreateHaVipResponse) {
    response = &CreateHaVipResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateHaVip
// This API (CreateHaVip) is used to create a highly available virtual IP (HAVIP)
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateHaVip(request *CreateHaVipRequest) (response *CreateHaVipResponse, err error) {
    if request == nil {
        request = NewCreateHaVipRequest()
    }
    response = NewCreateHaVipResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLocalGatewayRequest() (request *CreateLocalGatewayRequest) {
    request = &CreateLocalGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateLocalGateway")
    return
}

func NewCreateLocalGatewayResponse() (response *CreateLocalGatewayResponse) {
    response = &CreateLocalGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateLocalGateway
// This API is used to create a local gateway for a CDC instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_LOCALGATEWAYALREADYEXISTS = "UnsupportedOperation.LocalGateWayAlreadyExists"
func (c *Client) CreateLocalGateway(request *CreateLocalGatewayRequest) (response *CreateLocalGatewayResponse, err error) {
    if request == nil {
        request = NewCreateLocalGatewayRequest()
    }
    response = NewCreateLocalGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNatGatewayRequest() (request *CreateNatGatewayRequest) {
    request = &CreateNatGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateNatGateway")
    return
}

func NewCreateNatGatewayResponse() (response *CreateNatGatewayResponse) {
    response = &CreateNatGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateNatGateway
// This API (CreateNatGateway) is used to create a NAT gateway.
//
// error code that may be returned:
//  ADDRESSQUOTALIMITEXCEEDED = "AddressQuotaLimitExceeded"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDACCOUNT_NOTSUPPORTED = "InvalidAccount.NotSupported"
//  INVALIDADDRESSSTATE = "InvalidAddressState"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDVPCID_MALFORMED = "InvalidVpcId.Malformed"
//  INVALIDVPCID_NOTFOUND = "InvalidVpcId.NotFound"
//  LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.AddressQuotaLimitExceeded"
//  LIMITEXCEEDED_DAILYALLOCATEADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.DailyAllocateAddressQuotaLimitExceeded"
//  LIMITEXCEEDED_NATGATEWAYPERVPCLIMITEXCEEDED = "LimitExceeded.NatGatewayPerVpcLimitExceeded"
//  LIMITEXCEEDED_PUBLICIPADDRESSPERNATGATEWAYLIMITEXCEEDED = "LimitExceeded.PublicIpAddressPerNatGatewayLimitExceeded"
//  RESOURCEINUSE_ADDRESS = "ResourceInUse.Address"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOREALNAMEAUTHENTICATION = "UnauthorizedOperation.NoRealNameAuthentication"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INSUFFICIENTFUNDS = "UnsupportedOperation.InsufficientFunds"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_PUBLICIPADDRESSISNOTBGPIP = "UnsupportedOperation.PublicIpAddressIsNotBGPIp"
//  UNSUPPORTEDOPERATION_PUBLICIPADDRESSISNOTEXISTED = "UnsupportedOperation.PublicIpAddressIsNotExisted"
//  UNSUPPORTEDOPERATION_PUBLICIPADDRESSNOTBILLEDBYTRAFFIC = "UnsupportedOperation.PublicIpAddressNotBilledByTraffic"
func (c *Client) CreateNatGateway(request *CreateNatGatewayRequest) (response *CreateNatGatewayResponse, err error) {
    if request == nil {
        request = NewCreateNatGatewayRequest()
    }
    response = NewCreateNatGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNatGatewayDestinationIpPortTranslationNatRuleRequest() (request *CreateNatGatewayDestinationIpPortTranslationNatRuleRequest) {
    request = &CreateNatGatewayDestinationIpPortTranslationNatRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateNatGatewayDestinationIpPortTranslationNatRule")
    return
}

func NewCreateNatGatewayDestinationIpPortTranslationNatRuleResponse() (response *CreateNatGatewayDestinationIpPortTranslationNatRuleResponse) {
    response = &CreateNatGatewayDestinationIpPortTranslationNatRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateNatGatewayDestinationIpPortTranslationNatRule
// This API (CreateNatGatewayDestinationIpPortTranslationNatRule) is used to create a port forwarding rule for a NAT gateway.
//
// error code that may be returned:
//  ADDRESSQUOTALIMITEXCEEDED = "AddressQuotaLimitExceeded"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDACCOUNT_NOTSUPPORTED = "InvalidAccount.NotSupported"
//  INVALIDADDRESSSTATE = "InvalidAddressState"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDVPCID_MALFORMED = "InvalidVpcId.Malformed"
//  INVALIDVPCID_NOTFOUND = "InvalidVpcId.NotFound"
//  LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.AddressQuotaLimitExceeded"
//  LIMITEXCEEDED_DAILYALLOCATEADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.DailyAllocateAddressQuotaLimitExceeded"
//  LIMITEXCEEDED_NATGATEWAYPERVPCLIMITEXCEEDED = "LimitExceeded.NatGatewayPerVpcLimitExceeded"
//  LIMITEXCEEDED_PUBLICIPADDRESSPERNATGATEWAYLIMITEXCEEDED = "LimitExceeded.PublicIpAddressPerNatGatewayLimitExceeded"
//  RESOURCEINUSE_ADDRESS = "ResourceInUse.Address"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOREALNAMEAUTHENTICATION = "UnauthorizedOperation.NoRealNameAuthentication"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INSUFFICIENTFUNDS = "UnsupportedOperation.InsufficientFunds"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_PUBLICIPADDRESSISNOTBGPIP = "UnsupportedOperation.PublicIpAddressIsNotBGPIp"
//  UNSUPPORTEDOPERATION_PUBLICIPADDRESSISNOTEXISTED = "UnsupportedOperation.PublicIpAddressIsNotExisted"
//  UNSUPPORTEDOPERATION_PUBLICIPADDRESSNOTBILLEDBYTRAFFIC = "UnsupportedOperation.PublicIpAddressNotBilledByTraffic"
func (c *Client) CreateNatGatewayDestinationIpPortTranslationNatRule(request *CreateNatGatewayDestinationIpPortTranslationNatRuleRequest) (response *CreateNatGatewayDestinationIpPortTranslationNatRuleResponse, err error) {
    if request == nil {
        request = NewCreateNatGatewayDestinationIpPortTranslationNatRuleRequest()
    }
    response = NewCreateNatGatewayDestinationIpPortTranslationNatRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNatGatewaySourceIpTranslationNatRuleRequest() (request *CreateNatGatewaySourceIpTranslationNatRuleRequest) {
    request = &CreateNatGatewaySourceIpTranslationNatRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateNatGatewaySourceIpTranslationNatRule")
    return
}

func NewCreateNatGatewaySourceIpTranslationNatRuleResponse() (response *CreateNatGatewaySourceIpTranslationNatRuleResponse) {
    response = &CreateNatGatewaySourceIpTranslationNatRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateNatGatewaySourceIpTranslationNatRule
// This API is used to create a SNAT rule for the NAT Gateway.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NATSNATRULEEXISTS = "InvalidParameterValue.NatSnatRuleExists"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION_NATGATEWAYTYPENOTSUPPORTSNAT = "UnsupportedOperation.NatGatewayTypeNotSupportSNAT"
func (c *Client) CreateNatGatewaySourceIpTranslationNatRule(request *CreateNatGatewaySourceIpTranslationNatRuleRequest) (response *CreateNatGatewaySourceIpTranslationNatRuleResponse, err error) {
    if request == nil {
        request = NewCreateNatGatewaySourceIpTranslationNatRuleRequest()
    }
    response = NewCreateNatGatewaySourceIpTranslationNatRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNetDetectRequest() (request *CreateNetDetectRequest) {
    request = &CreateNetDetectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateNetDetect")
    return
}

func NewCreateNetDetectResponse() (response *CreateNetDetectResponse) {
    response = &CreateNetDetectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateNetDetect
// This API is used to create a network detection instance.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_NEXTHOPMISMATCH = "InvalidParameter.NextHopMismatch"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NETDETECTSAMEIP = "InvalidParameterValue.NetDetectSameIp"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_CONFLICTWITHDOCKERROUTE = "UnsupportedOperation.ConflictWithDockerRoute"
//  UNSUPPORTEDOPERATION_ECMPWITHUSERROUTE = "UnsupportedOperation.EcmpWithUserRoute"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
func (c *Client) CreateNetDetect(request *CreateNetDetectRequest) (response *CreateNetDetectResponse, err error) {
    if request == nil {
        request = NewCreateNetDetectRequest()
    }
    response = NewCreateNetDetectResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNetworkAclRequest() (request *CreateNetworkAclRequest) {
    request = &CreateNetworkAclRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateNetworkAcl")
    return
}

func NewCreateNetworkAclResponse() (response *CreateNetworkAclResponse) {
    response = &CreateNetworkAclResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateNetworkAcl
// This API is used to create a <a href="https://intl.cloud.tencent.com/document/product/215/20088?from_cn_redirect=1">network ACL</a>.
//
// * The inbound and outbound rules for a new network ACL are "Deny All" by default. You need to call `ModifyNetworkAclEntries` after creation to set rules for the network ACL as needed.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateNetworkAcl(request *CreateNetworkAclRequest) (response *CreateNetworkAclResponse, err error) {
    if request == nil {
        request = NewCreateNetworkAclRequest()
    }
    response = NewCreateNetworkAclResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNetworkInterfaceRequest() (request *CreateNetworkInterfaceRequest) {
    request = &CreateNetworkInterfaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateNetworkInterface")
    return
}

func NewCreateNetworkInterfaceResponse() (response *CreateNetworkInterfaceResponse) {
    response = &CreateNetworkInterfaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateNetworkInterface
// This API is used to create one or more ENIs.
//
// * You can specify private IP addresses and a primary IP when creating an ENI. The specified private IP must be in the same subnet as the ENI and is not occupied.
//
// * When creating an ENI, you can specify the number of private IP addresses that you want to apply for. The system will randomly generate private IP addresses.
//
// * An ENI can only be bound with a limited number of IP addresses. For more information about resource limits, see <a href="/document/product/576/18527">ENI Use Limits</a>.
//
// * You can bind an existing security group when creating an ENI.
//
// * You can bind a tag when creating an ENI. The tag list in the response indicates the tags that have been successfully added.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_RESERVED = "InvalidParameterValue.Reserved"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateNetworkInterface(request *CreateNetworkInterfaceRequest) (response *CreateNetworkInterfaceResponse, err error) {
    if request == nil {
        request = NewCreateNetworkInterfaceRequest()
    }
    response = NewCreateNetworkInterfaceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRouteTableRequest() (request *CreateRouteTableRequest) {
    request = &CreateRouteTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateRouteTable")
    return
}

func NewCreateRouteTableResponse() (response *CreateRouteTableResponse) {
    response = &CreateRouteTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRouteTable
// This API is used to create a route table.
//
// * After the VPC instance has been created, the system creates a default route table with which all newly created subnets will be associated. By default, you can use this route table to manage your routing policies. If you have multiple routing policies, you can call the API for creating route tables to create more route tables to manage these routing policies.
//
// * You can bind a tag when creating a route table. The tag list in the response indicates the tags that have been successfully added.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateRouteTable(request *CreateRouteTableRequest) (response *CreateRouteTableResponse, err error) {
    if request == nil {
        request = NewCreateRouteTableRequest()
    }
    response = NewCreateRouteTableResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRoutesRequest() (request *CreateRoutesRequest) {
    request = &CreateRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateRoutes")
    return
}

func NewCreateRoutesResponse() (response *CreateRoutesResponse) {
    response = &CreateRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRoutes
// This API (CreateRoutes) is used to create a routing policy.
//
// * You can create routing policies in batch for a specified route table.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CIDRNOTINPEERVPC = "InvalidParameterValue.CidrNotInPeerVpc"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_VPCCIDRCONFLICT = "InvalidParameterValue.VpcCidrConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER_WITHGUESS = "UnknownParameter.WithGuess"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CDCSUBNETNOTSUPPORTUNLOCALGATEWAY = "UnsupportedOperation.CdcSubnetNotSupportUnLocalGateway"
//  UNSUPPORTEDOPERATION_CONFLICTWITHDOCKERROUTE = "UnsupportedOperation.ConflictWithDockerRoute"
//  UNSUPPORTEDOPERATION_ECMP = "UnsupportedOperation.Ecmp"
//  UNSUPPORTEDOPERATION_ECMPWITHCCNROUTE = "UnsupportedOperation.EcmpWithCcnRoute"
//  UNSUPPORTEDOPERATION_ECMPWITHUSERROUTE = "UnsupportedOperation.EcmpWithUserRoute"
//  UNSUPPORTEDOPERATION_NORMALSUBNETNOTSUPPORTLOCALGATEWAY = "UnsupportedOperation.NormalSubnetNotSupportLocalGateway"
//  UNSUPPORTEDOPERATION_SYSTEMROUTE = "UnsupportedOperation.SystemRoute"
func (c *Client) CreateRoutes(request *CreateRoutesRequest) (response *CreateRoutesResponse, err error) {
    if request == nil {
        request = NewCreateRoutesRequest()
    }
    response = NewCreateRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSecurityGroupRequest() (request *CreateSecurityGroupRequest) {
    request = &CreateSecurityGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateSecurityGroup")
    return
}

func NewCreateSecurityGroupResponse() (response *CreateSecurityGroupResponse) {
    response = &CreateSecurityGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSecurityGroup
// This API is used to create a security group (SecurityGroup).
//
// * Note the <a href="https://intl.cloud.tencent.com/document/product/213/12453?from_cn_redirect=1">maximum number of security groups</a> per project in each region under each account.
//
// * Both the inbound and outbound rules for a newly created security group are "Deny All" by default. You need to call CreateSecurityGroupPolicies to set security group rules based on your needs.
//
// * You can bind a tag when creating a security group. The tag list in the response indicates the tags that have been successfully added.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateSecurityGroup(request *CreateSecurityGroupRequest) (response *CreateSecurityGroupResponse, err error) {
    if request == nil {
        request = NewCreateSecurityGroupRequest()
    }
    response = NewCreateSecurityGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSecurityGroupPoliciesRequest() (request *CreateSecurityGroupPoliciesRequest) {
    request = &CreateSecurityGroupPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateSecurityGroupPolicies")
    return
}

func NewCreateSecurityGroupPoliciesResponse() (response *CreateSecurityGroupPoliciesResponse) {
    response = &CreateSecurityGroupPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSecurityGroupPolicies
// This API is used to create a security group policy (SecurityGroupPolicy).
//
// 
//
// For parameters of SecurityGroupPolicySet,
//
// <ul>
//
// <li>`Version`: the version number of a security group policy, which automatically increases by one each time you update the security policy, to prevent expiration of the updated routing policies. If it is left empty, any conflicts will be ignored.</li>
//
// <li>When creating the `Egress` and `Ingress` policies,<ul>
//
// <li>`Protocol`: allows `TCP`, `UDP`, `ICMP`, `ICMPV6`, `GRE`, or `ALL`.</li>
//
// <li>`CidrBlock`: a CIDR block in the correct format. In the classic network, if a `CidrBlock` contains private IPs of devices under your account other than CVMs, it does not mean this policy allows you to access these devices. The network isolation policies between tenants take priority over the private network policies in security groups.</li>
//
// <li>`Ipv6CidrBlock`: an IPv6 CIDR block in the correct format. In a classic network, if an `Ipv6CidrBlock` contains private IPv6 addresses on Tencent Cloud for devices under your account other than CVMs, it does not mean this policy allows you to access these devices. The network isolation policies between tenants take priority over the private network policies in security groups.</li>
//
// <li>`SecurityGroupId`: ID of the security group. It can be the ID of security group to be modified, or the ID of another security group in the same project. All private IPs of all CVMs under the security group will be covered. If this field is used, the policy will automatically change according to the CVM associated with the group ID while being used to match network messages. You don’t need to change it manually.</li>
//
// <li>`Port`: a single port number such as 80, or a port range in the format of “8000-8010”. You may use this field only if the `Protocol` field takes the value `TCP` or `UDP`. Otherwise `Protocol` and `Port` are mutually exclusive.</li>
//
// <li>`Action`: only allows `ACCEPT` or `DROP`.</li>
//
// <li>`CidrBlock`, `Ipv6CidrBlock`, `SecurityGroupId`, and `AddressTemplate` are mutually exclusive. `Protocol` + `Port` and `ServiceTemplate` are mutually exclusive.</li>
//
// <li>You can only create policies in one direction in each request. To specify the `PolicyIndex` parameter, use the same index number in policies.</li>
//
// </ul></li></ul>
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_SECURITYGROUPPOLICYSET = "LimitExceeded.SecurityGroupPolicySet"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER_WITHGUESS = "UnknownParameter.WithGuess"
//  UNSUPPORTEDOPERATION_CLBPOLICYLIMIT = "UnsupportedOperation.ClbPolicyLimit"
//  UNSUPPORTEDOPERATION_DUPLICATEPOLICY = "UnsupportedOperation.DuplicatePolicy"
//  UNSUPPORTEDOPERATION_VERSIONMISMATCH = "UnsupportedOperation.VersionMismatch"
func (c *Client) CreateSecurityGroupPolicies(request *CreateSecurityGroupPoliciesRequest) (response *CreateSecurityGroupPoliciesResponse, err error) {
    if request == nil {
        request = NewCreateSecurityGroupPoliciesRequest()
    }
    response = NewCreateSecurityGroupPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSecurityGroupWithPoliciesRequest() (request *CreateSecurityGroupWithPoliciesRequest) {
    request = &CreateSecurityGroupWithPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateSecurityGroupWithPolicies")
    return
}

func NewCreateSecurityGroupWithPoliciesResponse() (response *CreateSecurityGroupWithPoliciesResponse) {
    response = &CreateSecurityGroupWithPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSecurityGroupWithPolicies
// This API (CreateSecurityGroupWithPolicies) is used to create security groups, and add security group policies.
//
// * Note the<a href="https://intl.cloud.tencent.com/document/product/213/12453?from_cn_redirect=1">maximum number of security groups</a>per project in each region under each account.
//
// * Both the inbound and outbound policies for a newly created security group are Deny All by default. You need to call CreateSecurityGroupPolicies to set security group policies according to your needs.
//
// 
//
// Description:
//
// * `Version`: Indicates the version number of a security group policy, which will automatically increment by 1 every time you update the security policy, to prevent the expiration of the updated policies. If this field is left empty, any conflicts will be ignored.
//
// * `Protocol`: Values can be TCP, UDP, ICMP, ICMPV6, GRE, or ALL.
//
// * `CidrBlock`:  A CIDR block in the correct format. In a basic network, if a CidrBlock contains private IPs on Tencent Cloud for devices under your account other than CVMs, it does not mean this policy allows you to access these devices. The network isolation policies between tenants take priority over the private network policies in security groups.
//
// * `Ipv6CidrBlock`: An IPv6 CIDR block in the correct format. In a basic network, if an Ipv6CidrBlock contains private IPv6 addresses on Tencent Cloud for devices under your account other than CVMs, it does not mean this policy allows you to access these devices. The network isolation policies between tenants take priority over the private network policies in security groups.
//
// * `SecurityGroupId`: ID of the security group. It can be in the same project as the security group to be modified, including the ID of the security group itself, to represent private IP addresses of all CVMs under the security group. If this field is used, the policy will change without manual modification according to the CVM associated with the policy ID while being used to match network messages.
//
// * `Port`: A single port number, or a port range in the format of “8000-8010”. The Port field is accepted only if the value of the `Protocol` field is `TCP` or `UDP`. Otherwise Protocol and Port are mutually exclusive. 
//
// * `Action`: Values can be `ACCEPT` or `DROP`.
//
// * CidrBlock, Ipv6CidrBlock, SecurityGroupId, and AddressTemplate are exclusive and cannot be entered at the same time. “Protocol + Port” and ServiceTemplate are mutually exclusive and cannot be entered at the same time.
//
// * Only policies in one direction can be created in each request. If you need to specify the `PolicyIndex` parameter, the indexes of policies must be consistent.
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateSecurityGroupWithPolicies(request *CreateSecurityGroupWithPoliciesRequest) (response *CreateSecurityGroupWithPoliciesResponse, err error) {
    if request == nil {
        request = NewCreateSecurityGroupWithPoliciesRequest()
    }
    response = NewCreateSecurityGroupWithPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateServiceTemplateRequest() (request *CreateServiceTemplateRequest) {
    request = &CreateServiceTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateServiceTemplate")
    return
}

func NewCreateServiceTemplateResponse() (response *CreateServiceTemplateResponse) {
    response = &CreateServiceTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateServiceTemplate
// This API (CreateServiceTemplate) is used to create a protocol port template.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateServiceTemplate(request *CreateServiceTemplateRequest) (response *CreateServiceTemplateResponse, err error) {
    if request == nil {
        request = NewCreateServiceTemplateRequest()
    }
    response = NewCreateServiceTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateServiceTemplateGroupRequest() (request *CreateServiceTemplateGroupRequest) {
    request = &CreateServiceTemplateGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateServiceTemplateGroup")
    return
}

func NewCreateServiceTemplateGroupResponse() (response *CreateServiceTemplateGroupResponse) {
    response = &CreateServiceTemplateGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateServiceTemplateGroup
// This API (CreateServiceTemplateGroup) is used to create a protocol port template group.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) CreateServiceTemplateGroup(request *CreateServiceTemplateGroupRequest) (response *CreateServiceTemplateGroupResponse, err error) {
    if request == nil {
        request = NewCreateServiceTemplateGroupRequest()
    }
    response = NewCreateServiceTemplateGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSubnetRequest() (request *CreateSubnetRequest) {
    request = &CreateSubnetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateSubnet")
    return
}

func NewCreateSubnetResponse() (response *CreateSubnetResponse) {
    response = &CreateSubnetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSubnet
// This API is used to create a subnet.
//
// * You must create a VPC instance before creating a subnet.
//
// * After the subnet is successfully created, its IP address range cannot be modified. The subnet IP address range must fall within the VPC IP address range. They can be the same if the VPC instance has only one subnet. We recommend that you keep the subnet IP address range within the VPC IP address range to reserve IP address ranges for other subnets.
//
// * The subnet mask of the smallest IP address range that can be created is 28 (16 IP addresses), and that of the largest IP address range is 16 (65,536 IP addresses).
//
// * IP address ranges of different subnets cannot overlap with each other within the same VPC instance.
//
// * A subnet is automatically associated with the default route table once created.
//
// * You can bind a tag when creating a subnet. The tag list in the response indicates the tags that have been successfully added.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_SUBNETCONFLICT = "InvalidParameterValue.SubnetConflict"
//  INVALIDPARAMETERVALUE_SUBNETRANGE = "InvalidParameterValue.SubnetRange"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_ZONECONFLICT = "InvalidParameterValue.ZoneConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_DCGATEWAYSNOTFOUNDINVPC = "UnsupportedOperation.DcGatewaysNotFoundInVpc"
func (c *Client) CreateSubnet(request *CreateSubnetRequest) (response *CreateSubnetResponse, err error) {
    if request == nil {
        request = NewCreateSubnetRequest()
    }
    response = NewCreateSubnetResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSubnetsRequest() (request *CreateSubnetsRequest) {
    request = &CreateSubnetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateSubnets")
    return
}

func NewCreateSubnetsResponse() (response *CreateSubnetsResponse) {
    response = &CreateSubnetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSubnets
// This API is used to create subnets in batches.
//
// * You must create a VPC instance before creating a subnet.
//
// * After the subnet is successfully created, its IP address range cannot be modified. The subnet IP address range must fall within the VPC IP address range. They can be the same if the VPC has only one subnet. We recommend that you keep the subnet IP address range within the VPC IP address range to reserve IP address ranges for other subnets.
//
// * The subnet mask of the smallest IP address range that can be created is 28 (16 IP addresses), and that of the largest IP address range is 16 (65,536 IP addresses).
//
// * IP address ranges of different subnets cannot overlap with each other within the same VPC instance.
//
// * A subnet is automatically associated with the default route table once created.
//
// * You can bind a tag when creating a subnet. The tag list in the response indicates the tags that have been successfully added.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_SUBNETCONFLICT = "InvalidParameterValue.SubnetConflict"
//  INVALIDPARAMETERVALUE_SUBNETRANGE = "InvalidParameterValue.SubnetRange"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_ZONECONFLICT = "InvalidParameterValue.ZoneConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_DCGATEWAYSNOTFOUNDINVPC = "UnsupportedOperation.DcGatewaysNotFoundInVpc"
func (c *Client) CreateSubnets(request *CreateSubnetsRequest) (response *CreateSubnetsResponse, err error) {
    if request == nil {
        request = NewCreateSubnetsRequest()
    }
    response = NewCreateSubnetsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVpcRequest() (request *CreateVpcRequest) {
    request = &CreateVpcRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateVpc")
    return
}

func NewCreateVpcResponse() (response *CreateVpcResponse) {
    response = &CreateVpcResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateVpc
// This API is used to create a VPC instance.
//
// * The subnet mask of the smallest IP address range that can be created is 28 (16 IP addresses), and that of the largest IP address range is 16 (65,536 IP addresses). For more information on how to plan VPC IP ranges, see [Network Planning](https://intl.cloud.tencent.com/document/product/215/30313?from_cn_redirect=1).
//
// * The number of VPC instances that can be created in a region is limited. For more information, see <a href="https://intl.cloud.tencent.com/doc/product/215/537?from_cn_redirect=1" title="VPC Use Limits">VPC Use Limits</a>. To request more resources, please [submit a ticket](https://console.cloud.tencent.com/workorder/category).
//
// * You can bind tags when creating a VPC instance. The tag list in the response indicates the tags that have been successfully added.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateVpc(request *CreateVpcRequest) (response *CreateVpcResponse, err error) {
    if request == nil {
        request = NewCreateVpcRequest()
    }
    response = NewCreateVpcResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVpcEndPointRequest() (request *CreateVpcEndPointRequest) {
    request = &CreateVpcEndPointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateVpcEndPoint")
    return
}

func NewCreateVpcEndPointResponse() (response *CreateVpcEndPointResponse) {
    response = &CreateVpcEndPointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateVpcEndPoint
// This API is used to create an endpoint.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_RESERVED = "InvalidParameterValue.Reserved"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION_SPECIALENDPOINTSERVICE = "UnsupportedOperation.SpecialEndPointService"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
func (c *Client) CreateVpcEndPoint(request *CreateVpcEndPointRequest) (response *CreateVpcEndPointResponse, err error) {
    if request == nil {
        request = NewCreateVpcEndPointRequest()
    }
    response = NewCreateVpcEndPointResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVpcEndPointServiceRequest() (request *CreateVpcEndPointServiceRequest) {
    request = &CreateVpcEndPointServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateVpcEndPointService")
    return
}

func NewCreateVpcEndPointServiceResponse() (response *CreateVpcEndPointServiceResponse) {
    response = &CreateVpcEndPointServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateVpcEndPointService
// This API is used to create an endpoint service.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
func (c *Client) CreateVpcEndPointService(request *CreateVpcEndPointServiceRequest) (response *CreateVpcEndPointServiceResponse, err error) {
    if request == nil {
        request = NewCreateVpcEndPointServiceRequest()
    }
    response = NewCreateVpcEndPointServiceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVpcEndPointServiceWhiteListRequest() (request *CreateVpcEndPointServiceWhiteListRequest) {
    request = &CreateVpcEndPointServiceWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateVpcEndPointServiceWhiteList")
    return
}

func NewCreateVpcEndPointServiceWhiteListResponse() (response *CreateVpcEndPointServiceWhiteListResponse) {
    response = &CreateVpcEndPointServiceWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateVpcEndPointServiceWhiteList
// This API is used to create the endpoint service allowlist.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_UINNOTFOUND = "UnsupportedOperation.UinNotFound"
func (c *Client) CreateVpcEndPointServiceWhiteList(request *CreateVpcEndPointServiceWhiteListRequest) (response *CreateVpcEndPointServiceWhiteListResponse, err error) {
    if request == nil {
        request = NewCreateVpcEndPointServiceWhiteListRequest()
    }
    response = NewCreateVpcEndPointServiceWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVpnConnectionRequest() (request *CreateVpnConnectionRequest) {
    request = &CreateVpnConnectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateVpnConnection")
    return
}

func NewCreateVpnConnectionResponse() (response *CreateVpnConnectionResponse) {
    response = &CreateVpnConnectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateVpnConnection
// This API (CreateVpnConnection) is used to create VPN tunnel.
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_VPCCIDRCONFLICT = "InvalidParameterValue.VpcCidrConflict"
//  INVALIDPARAMETERVALUE_VPNCONNCIDRCONFLICT = "InvalidParameterValue.VpnConnCidrConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) CreateVpnConnection(request *CreateVpnConnectionRequest) (response *CreateVpnConnectionResponse, err error) {
    if request == nil {
        request = NewCreateVpnConnectionRequest()
    }
    response = NewCreateVpnConnectionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVpnGatewayRequest() (request *CreateVpnGatewayRequest) {
    request = &CreateVpnGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateVpnGateway")
    return
}

func NewCreateVpnGatewayResponse() (response *CreateVpnGatewayResponse) {
    response = &CreateVpnGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateVpnGateway
// This API (CreateVpnGateway) is used to create a VPN gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_VPNCONNCIDRCONFLICT = "InvalidParameterValue.VpnConnCidrConflict"
//  INVALIDVPCID_MALFORMED = "InvalidVpcId.Malformed"
//  INVALIDVPCID_NOTFOUND = "InvalidVpcId.NotFound"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOREALNAMEAUTHENTICATION = "UnauthorizedOperation.NoRealNameAuthentication"
//  UNSUPPORTEDOPERATION_INSUFFICIENTFUNDS = "UnsupportedOperation.InsufficientFunds"
func (c *Client) CreateVpnGateway(request *CreateVpnGatewayRequest) (response *CreateVpnGatewayResponse, err error) {
    if request == nil {
        request = NewCreateVpnGatewayRequest()
    }
    response = NewCreateVpnGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAddressTemplateRequest() (request *DeleteAddressTemplateRequest) {
    request = &DeleteAddressTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteAddressTemplate")
    return
}

func NewDeleteAddressTemplateResponse() (response *DeleteAddressTemplateResponse) {
    response = &DeleteAddressTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAddressTemplate
// This API (DeleteAddressTemplate) is used to delete an IP address template.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) DeleteAddressTemplate(request *DeleteAddressTemplateRequest) (response *DeleteAddressTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteAddressTemplateRequest()
    }
    response = NewDeleteAddressTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAddressTemplateGroupRequest() (request *DeleteAddressTemplateGroupRequest) {
    request = &DeleteAddressTemplateGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteAddressTemplateGroup")
    return
}

func NewDeleteAddressTemplateGroupResponse() (response *DeleteAddressTemplateGroupResponse) {
    response = &DeleteAddressTemplateGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAddressTemplateGroup
// This API (DeleteAddressTemplateGroup) is used to delete an IP address template group.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) DeleteAddressTemplateGroup(request *DeleteAddressTemplateGroupRequest) (response *DeleteAddressTemplateGroupResponse, err error) {
    if request == nil {
        request = NewDeleteAddressTemplateGroupRequest()
    }
    response = NewDeleteAddressTemplateGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAssistantCidrRequest() (request *DeleteAssistantCidrRequest) {
    request = &DeleteAssistantCidrRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteAssistantCidr")
    return
}

func NewDeleteAssistantCidrResponse() (response *DeleteAssistantCidrResponse) {
    response = &DeleteAssistantCidrResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAssistantCidr
// This API (DeleteAssistantCidr) is used to delete secondary CIDR blocks. (To use this API that is in Beta, please submit a ticket.)
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteAssistantCidr(request *DeleteAssistantCidrRequest) (response *DeleteAssistantCidrResponse, err error) {
    if request == nil {
        request = NewDeleteAssistantCidrRequest()
    }
    response = NewDeleteAssistantCidrResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBandwidthPackageRequest() (request *DeleteBandwidthPackageRequest) {
    request = &DeleteBandwidthPackageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteBandwidthPackage")
    return
}

func NewDeleteBandwidthPackageResponse() (response *DeleteBandwidthPackageResponse) {
    response = &DeleteBandwidthPackageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteBandwidthPackage
// This API is used to delete bandwidth packages, including [device bandwidth packages](https://intl.cloud.tencent.com/document/product/684/15246?from_cn_redirect=1#.E8.AE.BE.E5.A4.87.E5.B8.A6.E5.AE.BD.E5.8C.85) and [IP bandwidth packages](https://intl.cloud.tencent.com/document/product/684/15246?from_cn_redirect=1#ip-.E5.B8.A6.E5.AE.BD.E5.8C.85).
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
func (c *Client) DeleteBandwidthPackage(request *DeleteBandwidthPackageRequest) (response *DeleteBandwidthPackageResponse, err error) {
    if request == nil {
        request = NewDeleteBandwidthPackageRequest()
    }
    response = NewDeleteBandwidthPackageResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCcnRequest() (request *DeleteCcnRequest) {
    request = &DeleteCcnRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteCcn")
    return
}

func NewDeleteCcnResponse() (response *DeleteCcnResponse) {
    response = &DeleteCcnResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCcn
// This API (DeleteCcn) is used to delete CCNs.
//
// * After deletion, the routes between all instances associated with the CCN will be deleted, and the network will be interrupted. Please confirm this operation in advance.
//
// * CCN deletion is an irreversible operation. Please proceed with caution.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BANDWIDTHNOTEXPIRED = "UnsupportedOperation.BandwidthNotExpired"
func (c *Client) DeleteCcn(request *DeleteCcnRequest) (response *DeleteCcnResponse, err error) {
    if request == nil {
        request = NewDeleteCcnRequest()
    }
    response = NewDeleteCcnResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCustomerGatewayRequest() (request *DeleteCustomerGatewayRequest) {
    request = &DeleteCustomerGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteCustomerGateway")
    return
}

func NewDeleteCustomerGatewayResponse() (response *DeleteCustomerGatewayResponse) {
    response = &DeleteCustomerGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCustomerGateway
// This API (DeleteCustomerGateway) is used to delete customer gateways.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteCustomerGateway(request *DeleteCustomerGatewayRequest) (response *DeleteCustomerGatewayResponse, err error) {
    if request == nil {
        request = NewDeleteCustomerGatewayRequest()
    }
    response = NewDeleteCustomerGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDirectConnectGatewayRequest() (request *DeleteDirectConnectGatewayRequest) {
    request = &DeleteDirectConnectGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteDirectConnectGateway")
    return
}

func NewDeleteDirectConnectGatewayResponse() (response *DeleteDirectConnectGatewayResponse) {
    response = &DeleteDirectConnectGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDirectConnectGateway
// This API is used to delete a direct connect gateway.
//
// <li>For a NAT gateway, NAT and ACL rules will be cleared upon the deletion of a direct connect gateway.
//
// <li>After the deletion of a direct connect gateway, the routing policy associated with the gateway in the route table will also be deleted.
//
// This API is completed asynchronously. If you need to query the async job execution results, please use the `RequestId` returned by this API to poll the `QueryTask` API.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteDirectConnectGateway(request *DeleteDirectConnectGatewayRequest) (response *DeleteDirectConnectGatewayResponse, err error) {
    if request == nil {
        request = NewDeleteDirectConnectGatewayRequest()
    }
    response = NewDeleteDirectConnectGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDirectConnectGatewayCcnRoutesRequest() (request *DeleteDirectConnectGatewayCcnRoutesRequest) {
    request = &DeleteDirectConnectGatewayCcnRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteDirectConnectGatewayCcnRoutes")
    return
}

func NewDeleteDirectConnectGatewayCcnRoutesResponse() (response *DeleteDirectConnectGatewayCcnRoutesResponse) {
    response = &DeleteDirectConnectGatewayCcnRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDirectConnectGatewayCcnRoutes
// This API (DeleteDirectConnectGatewayCcnRoutes) is used to delete the CCN routes (IDC IP range) of a Direct Connect gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDirectConnectGatewayCcnRoutes(request *DeleteDirectConnectGatewayCcnRoutesRequest) (response *DeleteDirectConnectGatewayCcnRoutesResponse, err error) {
    if request == nil {
        request = NewDeleteDirectConnectGatewayCcnRoutesRequest()
    }
    response = NewDeleteDirectConnectGatewayCcnRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFlowLogRequest() (request *DeleteFlowLogRequest) {
    request = &DeleteFlowLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteFlowLog")
    return
}

func NewDeleteFlowLogResponse() (response *DeleteFlowLogResponse) {
    response = &DeleteFlowLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteFlowLog
// This API is used to delete a flow log.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteFlowLog(request *DeleteFlowLogRequest) (response *DeleteFlowLogResponse, err error) {
    if request == nil {
        request = NewDeleteFlowLogRequest()
    }
    response = NewDeleteFlowLogResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteHaVipRequest() (request *DeleteHaVipRequest) {
    request = &DeleteHaVipRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteHaVip")
    return
}

func NewDeleteHaVipResponse() (response *DeleteHaVipResponse) {
    response = &DeleteHaVipResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteHaVip
// This API (DeleteHaVip) is used to delete Highly Available Virtual IP (HAVIP)<br />
//
// This API is completed asynchronously. If you need to query the async job execution results, please use the `RequestId` returned by this API to query the `QueryTask` API.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteHaVip(request *DeleteHaVipRequest) (response *DeleteHaVipResponse, err error) {
    if request == nil {
        request = NewDeleteHaVipRequest()
    }
    response = NewDeleteHaVipResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLocalGatewayRequest() (request *DeleteLocalGatewayRequest) {
    request = &DeleteLocalGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteLocalGateway")
    return
}

func NewDeleteLocalGatewayResponse() (response *DeleteLocalGatewayResponse) {
    response = &DeleteLocalGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLocalGateway
// This API is used to delete the local gateway of a CDC instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteLocalGateway(request *DeleteLocalGatewayRequest) (response *DeleteLocalGatewayResponse, err error) {
    if request == nil {
        request = NewDeleteLocalGatewayRequest()
    }
    response = NewDeleteLocalGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNatGatewayRequest() (request *DeleteNatGatewayRequest) {
    request = &DeleteNatGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteNatGateway")
    return
}

func NewDeleteNatGatewayResponse() (response *DeleteNatGatewayResponse) {
    response = &DeleteNatGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteNatGateway
// This API (DeleteNatGateway) is used to delete a NAT gateway.
//
// After the deletion of a NAT gateway, the system will automatically delete the routing entry that contains the NAT gateway from the route table. It will also unbind the Elastic IP.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) DeleteNatGateway(request *DeleteNatGatewayRequest) (response *DeleteNatGatewayResponse, err error) {
    if request == nil {
        request = NewDeleteNatGatewayRequest()
    }
    response = NewDeleteNatGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNatGatewayDestinationIpPortTranslationNatRuleRequest() (request *DeleteNatGatewayDestinationIpPortTranslationNatRuleRequest) {
    request = &DeleteNatGatewayDestinationIpPortTranslationNatRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteNatGatewayDestinationIpPortTranslationNatRule")
    return
}

func NewDeleteNatGatewayDestinationIpPortTranslationNatRuleResponse() (response *DeleteNatGatewayDestinationIpPortTranslationNatRuleResponse) {
    response = &DeleteNatGatewayDestinationIpPortTranslationNatRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteNatGatewayDestinationIpPortTranslationNatRule
// This API (DeleteNatGatewayDestinationIpPortTranslationNatRule) is used to delete a port forwarding rule for a NAT gateway.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) DeleteNatGatewayDestinationIpPortTranslationNatRule(request *DeleteNatGatewayDestinationIpPortTranslationNatRuleRequest) (response *DeleteNatGatewayDestinationIpPortTranslationNatRuleResponse, err error) {
    if request == nil {
        request = NewDeleteNatGatewayDestinationIpPortTranslationNatRuleRequest()
    }
    response = NewDeleteNatGatewayDestinationIpPortTranslationNatRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNatGatewaySourceIpTranslationNatRuleRequest() (request *DeleteNatGatewaySourceIpTranslationNatRuleRequest) {
    request = &DeleteNatGatewaySourceIpTranslationNatRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteNatGatewaySourceIpTranslationNatRule")
    return
}

func NewDeleteNatGatewaySourceIpTranslationNatRuleResponse() (response *DeleteNatGatewaySourceIpTranslationNatRuleResponse) {
    response = &DeleteNatGatewaySourceIpTranslationNatRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteNatGatewaySourceIpTranslationNatRule
// This API is used to delete a SNAT forwarding rule of the NAT Gateway.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteNatGatewaySourceIpTranslationNatRule(request *DeleteNatGatewaySourceIpTranslationNatRuleRequest) (response *DeleteNatGatewaySourceIpTranslationNatRuleResponse, err error) {
    if request == nil {
        request = NewDeleteNatGatewaySourceIpTranslationNatRuleRequest()
    }
    response = NewDeleteNatGatewaySourceIpTranslationNatRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNetDetectRequest() (request *DeleteNetDetectRequest) {
    request = &DeleteNetDetectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteNetDetect")
    return
}

func NewDeleteNetDetectResponse() (response *DeleteNetDetectResponse) {
    response = &DeleteNetDetectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteNetDetect
// This API (DeleteNetDetect) is used to delete a network detection instance.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteNetDetect(request *DeleteNetDetectRequest) (response *DeleteNetDetectResponse, err error) {
    if request == nil {
        request = NewDeleteNetDetectRequest()
    }
    response = NewDeleteNetDetectResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNetworkAclRequest() (request *DeleteNetworkAclRequest) {
    request = &DeleteNetworkAclRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteNetworkAcl")
    return
}

func NewDeleteNetworkAclResponse() (response *DeleteNetworkAclResponse) {
    response = &DeleteNetworkAclResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteNetworkAcl
// This API is used to delete a network ACL.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteNetworkAcl(request *DeleteNetworkAclRequest) (response *DeleteNetworkAclResponse, err error) {
    if request == nil {
        request = NewDeleteNetworkAclRequest()
    }
    response = NewDeleteNetworkAclResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNetworkInterfaceRequest() (request *DeleteNetworkInterfaceRequest) {
    request = &DeleteNetworkInterfaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteNetworkInterface")
    return
}

func NewDeleteNetworkInterfaceResponse() (response *DeleteNetworkInterfaceResponse) {
    response = &DeleteNetworkInterfaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteNetworkInterface
// This API (DeleteNetworkInterface) is used to delete ENIs.
//
// * An ENI that has been bound to a CVM cannot be deleted.
//
// * An ENI can be deleted only after being unbound from the server. After the deletion, all private IP addresses associated with the ENI will be released.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) DeleteNetworkInterface(request *DeleteNetworkInterfaceRequest) (response *DeleteNetworkInterfaceResponse, err error) {
    if request == nil {
        request = NewDeleteNetworkInterfaceRequest()
    }
    response = NewDeleteNetworkInterfaceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRouteTableRequest() (request *DeleteRouteTableRequest) {
    request = &DeleteRouteTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteRouteTable")
    return
}

func NewDeleteRouteTableResponse() (response *DeleteRouteTableResponse) {
    response = &DeleteRouteTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRouteTable
// This API is used to delete a route table.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRouteTable(request *DeleteRouteTableRequest) (response *DeleteRouteTableResponse, err error) {
    if request == nil {
        request = NewDeleteRouteTableRequest()
    }
    response = NewDeleteRouteTableResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRoutesRequest() (request *DeleteRoutesRequest) {
    request = &DeleteRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteRoutes")
    return
}

func NewDeleteRoutesResponse() (response *DeleteRoutesResponse) {
    response = &DeleteRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRoutes
// This API (DeleteRoutes) is used to delete routing policies in batches from a route table.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER_WITHGUESS = "UnknownParameter.WithGuess"
//  UNSUPPORTEDOPERATION_DISABLEDNOTIFYCCN = "UnsupportedOperation.DisabledNotifyCcn"
//  UNSUPPORTEDOPERATION_SYSTEMROUTE = "UnsupportedOperation.SystemRoute"
func (c *Client) DeleteRoutes(request *DeleteRoutesRequest) (response *DeleteRoutesResponse, err error) {
    if request == nil {
        request = NewDeleteRoutesRequest()
    }
    response = NewDeleteRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSecurityGroupRequest() (request *DeleteSecurityGroupRequest) {
    request = &DeleteSecurityGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteSecurityGroup")
    return
}

func NewDeleteSecurityGroupResponse() (response *DeleteSecurityGroupResponse) {
    response = &DeleteSecurityGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteSecurityGroup
// This API (DeleteSecurityGroup) is used to delete security groups (SecurityGroup).
//
// * Only security groups under the current account can be deleted.
//
// * A security group cannot be deleted directly if its instance ID is used in the policy of another security group. You need to modify the policy first and then delete the security group.
//
// * A security group cannot be recovered after deletion, please proceed with caution.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDSECURITYGROUPID_MALFORMED = "InvalidSecurityGroupID.Malformed"
//  INVALIDSECURITYGROUPID_NOTFOUND = "InvalidSecurityGroupID.NotFound"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteSecurityGroup(request *DeleteSecurityGroupRequest) (response *DeleteSecurityGroupResponse, err error) {
    if request == nil {
        request = NewDeleteSecurityGroupRequest()
    }
    response = NewDeleteSecurityGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSecurityGroupPoliciesRequest() (request *DeleteSecurityGroupPoliciesRequest) {
    request = &DeleteSecurityGroupPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteSecurityGroupPolicies")
    return
}

func NewDeleteSecurityGroupPoliciesResponse() (response *DeleteSecurityGroupPoliciesResponse) {
    response = &DeleteSecurityGroupPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteSecurityGroupPolicies
// This API (DeleteSecurityGroupPolicies) is used to delete security group policies (SecurityGroupPolicy).
//
// * SecurityGroupPolicySet.Version is used to specify the version of the security group you are operating. If the specified Version number differs from the latest version of the current security group, a failure will be returned. If Version is not specified, the policy of the specified PolicyIndex will be deleted directly.
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_VERSIONMISMATCH = "UnsupportedOperation.VersionMismatch"
func (c *Client) DeleteSecurityGroupPolicies(request *DeleteSecurityGroupPoliciesRequest) (response *DeleteSecurityGroupPoliciesResponse, err error) {
    if request == nil {
        request = NewDeleteSecurityGroupPoliciesRequest()
    }
    response = NewDeleteSecurityGroupPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteServiceTemplateRequest() (request *DeleteServiceTemplateRequest) {
    request = &DeleteServiceTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteServiceTemplate")
    return
}

func NewDeleteServiceTemplateResponse() (response *DeleteServiceTemplateResponse) {
    response = &DeleteServiceTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteServiceTemplate
// This API (DeleteServiceTemplate) is used to delete a protocol port template.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) DeleteServiceTemplate(request *DeleteServiceTemplateRequest) (response *DeleteServiceTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteServiceTemplateRequest()
    }
    response = NewDeleteServiceTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteServiceTemplateGroupRequest() (request *DeleteServiceTemplateGroupRequest) {
    request = &DeleteServiceTemplateGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteServiceTemplateGroup")
    return
}

func NewDeleteServiceTemplateGroupResponse() (response *DeleteServiceTemplateGroupResponse) {
    response = &DeleteServiceTemplateGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteServiceTemplateGroup
// This API (DeleteServiceTemplateGroup) is used to delete a protocol port template group.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) DeleteServiceTemplateGroup(request *DeleteServiceTemplateGroupRequest) (response *DeleteServiceTemplateGroupResponse, err error) {
    if request == nil {
        request = NewDeleteServiceTemplateGroupRequest()
    }
    response = NewDeleteServiceTemplateGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSubnetRequest() (request *DeleteSubnetRequest) {
    request = &DeleteSubnetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteSubnet")
    return
}

func NewDeleteSubnetResponse() (response *DeleteSubnetResponse) {
    response = &DeleteSubnetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteSubnet
// This API (DeleteSubnet) is used to delete subnets.
//
// Before deleting a subnet, you need to remove all resources in the subnet, including CVMs, load balancers, cloud data, NoSQL databases, and ENIs.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteSubnet(request *DeleteSubnetRequest) (response *DeleteSubnetResponse, err error) {
    if request == nil {
        request = NewDeleteSubnetRequest()
    }
    response = NewDeleteSubnetResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteVpcRequest() (request *DeleteVpcRequest) {
    request = &DeleteVpcRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteVpc")
    return
}

func NewDeleteVpcResponse() (response *DeleteVpcResponse) {
    response = &DeleteVpcResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteVpc
// This API (DeleteVpc) is used to delete VPCs.
//
// * Before deleting a VPC, ensure that the VPC contains no resources, including CVMs, cloud databases, NoSQL databases, VPN gateways, direct connect gateways, load balancers, peering connections, and basic network devices that are linked to the VPC.
//
// * The deletion of VPCs is irreversible. Proceed with caution.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteVpc(request *DeleteVpcRequest) (response *DeleteVpcResponse, err error) {
    if request == nil {
        request = NewDeleteVpcRequest()
    }
    response = NewDeleteVpcResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteVpcEndPointRequest() (request *DeleteVpcEndPointRequest) {
    request = &DeleteVpcEndPointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteVpcEndPoint")
    return
}

func NewDeleteVpcEndPointResponse() (response *DeleteVpcEndPointResponse) {
    response = &DeleteVpcEndPointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteVpcEndPoint
// This API is used to delete an endpoint.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteVpcEndPoint(request *DeleteVpcEndPointRequest) (response *DeleteVpcEndPointResponse, err error) {
    if request == nil {
        request = NewDeleteVpcEndPointRequest()
    }
    response = NewDeleteVpcEndPointResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteVpcEndPointServiceRequest() (request *DeleteVpcEndPointServiceRequest) {
    request = &DeleteVpcEndPointServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteVpcEndPointService")
    return
}

func NewDeleteVpcEndPointServiceResponse() (response *DeleteVpcEndPointServiceResponse) {
    response = &DeleteVpcEndPointServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteVpcEndPointService
// This API is used to delete an endpoint service.
//
// 
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteVpcEndPointService(request *DeleteVpcEndPointServiceRequest) (response *DeleteVpcEndPointServiceResponse, err error) {
    if request == nil {
        request = NewDeleteVpcEndPointServiceRequest()
    }
    response = NewDeleteVpcEndPointServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteVpcEndPointServiceWhiteListRequest() (request *DeleteVpcEndPointServiceWhiteListRequest) {
    request = &DeleteVpcEndPointServiceWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteVpcEndPointServiceWhiteList")
    return
}

func NewDeleteVpcEndPointServiceWhiteListResponse() (response *DeleteVpcEndPointServiceWhiteListResponse) {
    response = &DeleteVpcEndPointServiceWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteVpcEndPointServiceWhiteList
// This API is used to delete the endpoint service allowlist.
//
// error code that may be returned:
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_UINNOTFOUND = "UnsupportedOperation.UinNotFound"
func (c *Client) DeleteVpcEndPointServiceWhiteList(request *DeleteVpcEndPointServiceWhiteListRequest) (response *DeleteVpcEndPointServiceWhiteListResponse, err error) {
    if request == nil {
        request = NewDeleteVpcEndPointServiceWhiteListRequest()
    }
    response = NewDeleteVpcEndPointServiceWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteVpnConnectionRequest() (request *DeleteVpnConnectionRequest) {
    request = &DeleteVpnConnectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteVpnConnection")
    return
}

func NewDeleteVpnConnectionResponse() (response *DeleteVpnConnectionResponse) {
    response = &DeleteVpnConnectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteVpnConnection
// This API (DeleteVpnConnection) is used to delete VPN tunnels.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) DeleteVpnConnection(request *DeleteVpnConnectionRequest) (response *DeleteVpnConnectionResponse, err error) {
    if request == nil {
        request = NewDeleteVpnConnectionRequest()
    }
    response = NewDeleteVpnConnectionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteVpnGatewayRequest() (request *DeleteVpnGatewayRequest) {
    request = &DeleteVpnGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteVpnGateway")
    return
}

func NewDeleteVpnGatewayResponse() (response *DeleteVpnGatewayResponse) {
    response = &DeleteVpnGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteVpnGateway
// This API (DeleteVpnGateway) is used to delete a VPN gateway. Currently, only deletion of pay-as-you-go IPSEC gateway instances in running status is supported.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDVPNGATEWAYID_MALFORMED = "InvalidVpnGatewayId.Malformed"
//  INVALIDVPNGATEWAYID_NOTFOUND = "InvalidVpnGatewayId.NotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteVpnGateway(request *DeleteVpnGatewayRequest) (response *DeleteVpnGatewayResponse, err error) {
    if request == nil {
        request = NewDeleteVpnGatewayRequest()
    }
    response = NewDeleteVpnGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountAttributesRequest() (request *DescribeAccountAttributesRequest) {
    request = &DescribeAccountAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeAccountAttributes")
    return
}

func NewDescribeAccountAttributesResponse() (response *DescribeAccountAttributesResponse) {
    response = &DescribeAccountAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccountAttributes
// This API (DescribeAccountAttributes) is used to query your account attributes.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDVPNGATEWAYID_MALFORMED = "InvalidVpnGatewayId.Malformed"
//  INVALIDVPNGATEWAYID_NOTFOUND = "InvalidVpnGatewayId.NotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAccountAttributes(request *DescribeAccountAttributesRequest) (response *DescribeAccountAttributesResponse, err error) {
    if request == nil {
        request = NewDescribeAccountAttributesRequest()
    }
    response = NewDescribeAccountAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAddressQuotaRequest() (request *DescribeAddressQuotaRequest) {
    request = &DescribeAddressQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeAddressQuota")
    return
}

func NewDescribeAddressQuotaResponse() (response *DescribeAddressQuotaResponse) {
    response = &DescribeAddressQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAddressQuota
// This API (DescribeAddressQuota) is used to query the quota information of your [Elastic IP](https://intl.cloud.tencent.com/document/product/213/1941?from_cn_redirect=1) (EIP) in the current region. For more information, see [EIP product introduction](https://intl.cloud.tencent.com/document/product/213/5733?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
func (c *Client) DescribeAddressQuota(request *DescribeAddressQuotaRequest) (response *DescribeAddressQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeAddressQuotaRequest()
    }
    response = NewDescribeAddressQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAddressTemplateGroupsRequest() (request *DescribeAddressTemplateGroupsRequest) {
    request = &DescribeAddressTemplateGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeAddressTemplateGroups")
    return
}

func NewDescribeAddressTemplateGroupsResponse() (response *DescribeAddressTemplateGroupsResponse) {
    response = &DescribeAddressTemplateGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAddressTemplateGroups
// This API (DescribeAddressTemplateGroups) is used to query an IP address template group.
//
// error code that may be returned:
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
func (c *Client) DescribeAddressTemplateGroups(request *DescribeAddressTemplateGroupsRequest) (response *DescribeAddressTemplateGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeAddressTemplateGroupsRequest()
    }
    response = NewDescribeAddressTemplateGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAddressTemplatesRequest() (request *DescribeAddressTemplatesRequest) {
    request = &DescribeAddressTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeAddressTemplates")
    return
}

func NewDescribeAddressTemplatesResponse() (response *DescribeAddressTemplatesResponse) {
    response = &DescribeAddressTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAddressTemplates
// This API (DescribeAddressTemplates) is used to query an IP address template.
//
// error code that may be returned:
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
func (c *Client) DescribeAddressTemplates(request *DescribeAddressTemplatesRequest) (response *DescribeAddressTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeAddressTemplatesRequest()
    }
    response = NewDescribeAddressTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAddressesRequest() (request *DescribeAddressesRequest) {
    request = &DescribeAddressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeAddresses")
    return
}

func NewDescribeAddressesResponse() (response *DescribeAddressesResponse) {
    response = &DescribeAddressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAddresses
// This API (DescribeAddresses) is used to query the information of one or multiple [Elastic IPs](https://intl.cloud.tencent.com/document/product/213/1941?from_cn_redirect=1).
//
// * If the parameter is empty, a number (as specified by the `Limit`, the default value is 20) of EIPs will be returned.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
func (c *Client) DescribeAddresses(request *DescribeAddressesRequest) (response *DescribeAddressesResponse, err error) {
    if request == nil {
        request = NewDescribeAddressesRequest()
    }
    response = NewDescribeAddressesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssistantCidrRequest() (request *DescribeAssistantCidrRequest) {
    request = &DescribeAssistantCidrRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeAssistantCidr")
    return
}

func NewDescribeAssistantCidrResponse() (response *DescribeAssistantCidrResponse) {
    response = &DescribeAssistantCidrResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssistantCidr
// This API (DescribeAssistantCidr) is used to query a list of secondary CIDR blocks. (To use this API that is in Beta, please submit a ticket.)
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
func (c *Client) DescribeAssistantCidr(request *DescribeAssistantCidrRequest) (response *DescribeAssistantCidrResponse, err error) {
    if request == nil {
        request = NewDescribeAssistantCidrRequest()
    }
    response = NewDescribeAssistantCidrResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBandwidthPackageBillUsageRequest() (request *DescribeBandwidthPackageBillUsageRequest) {
    request = &DescribeBandwidthPackageBillUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeBandwidthPackageBillUsage")
    return
}

func NewDescribeBandwidthPackageBillUsageResponse() (response *DescribeBandwidthPackageBillUsageResponse) {
    response = &DescribeBandwidthPackageBillUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBandwidthPackageBillUsage
// This API is used to query the current billable usage of a pay-as-you-go bandwidth package.
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
func (c *Client) DescribeBandwidthPackageBillUsage(request *DescribeBandwidthPackageBillUsageRequest) (response *DescribeBandwidthPackageBillUsageResponse, err error) {
    if request == nil {
        request = NewDescribeBandwidthPackageBillUsageRequest()
    }
    response = NewDescribeBandwidthPackageBillUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBandwidthPackageQuotaRequest() (request *DescribeBandwidthPackageQuotaRequest) {
    request = &DescribeBandwidthPackageQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeBandwidthPackageQuota")
    return
}

func NewDescribeBandwidthPackageQuotaResponse() (response *DescribeBandwidthPackageQuotaResponse) {
    response = &DescribeBandwidthPackageQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBandwidthPackageQuota
// This API is used to query the maximum and used number of bandwidth packages under the account in the current region.
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
func (c *Client) DescribeBandwidthPackageQuota(request *DescribeBandwidthPackageQuotaRequest) (response *DescribeBandwidthPackageQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeBandwidthPackageQuotaRequest()
    }
    response = NewDescribeBandwidthPackageQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBandwidthPackageResourcesRequest() (request *DescribeBandwidthPackageResourcesRequest) {
    request = &DescribeBandwidthPackageResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeBandwidthPackageResources")
    return
}

func NewDescribeBandwidthPackageResourcesResponse() (response *DescribeBandwidthPackageResourcesResponse) {
    response = &DescribeBandwidthPackageResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBandwidthPackageResources
// This API is used to query resources in a bandwidth package based on the unique package ID. You can filter the result by specifying conditions and paginate the query results.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBandwidthPackageResources(request *DescribeBandwidthPackageResourcesRequest) (response *DescribeBandwidthPackageResourcesResponse, err error) {
    if request == nil {
        request = NewDescribeBandwidthPackageResourcesRequest()
    }
    response = NewDescribeBandwidthPackageResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBandwidthPackagesRequest() (request *DescribeBandwidthPackagesRequest) {
    request = &DescribeBandwidthPackagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeBandwidthPackages")
    return
}

func NewDescribeBandwidthPackagesResponse() (response *DescribeBandwidthPackagesResponse) {
    response = &DescribeBandwidthPackagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBandwidthPackages
// This API is used to query bandwidth package information, including the unique ID of the bandwidth package, the type, the billing mode, the name, and the resource information.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBandwidthPackages(request *DescribeBandwidthPackagesRequest) (response *DescribeBandwidthPackagesResponse, err error) {
    if request == nil {
        request = NewDescribeBandwidthPackagesRequest()
    }
    response = NewDescribeBandwidthPackagesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCcnAttachedInstancesRequest() (request *DescribeCcnAttachedInstancesRequest) {
    request = &DescribeCcnAttachedInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeCcnAttachedInstances")
    return
}

func NewDescribeCcnAttachedInstancesResponse() (response *DescribeCcnAttachedInstancesResponse) {
    response = &DescribeCcnAttachedInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCcnAttachedInstances
// This API (DescribeCcnAttachedInstances) is used to query the network instances associated with the CCN instance.
//
// error code that may be returned:
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETER_FILTERNOTDICT = "InvalidParameter.FilterNotDict"
//  INVALIDPARAMETER_FILTERVALUESNOTLIST = "InvalidParameter.FilterValuesNotList"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
func (c *Client) DescribeCcnAttachedInstances(request *DescribeCcnAttachedInstancesRequest) (response *DescribeCcnAttachedInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeCcnAttachedInstancesRequest()
    }
    response = NewDescribeCcnAttachedInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCcnRegionBandwidthLimitsRequest() (request *DescribeCcnRegionBandwidthLimitsRequest) {
    request = &DescribeCcnRegionBandwidthLimitsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeCcnRegionBandwidthLimits")
    return
}

func NewDescribeCcnRegionBandwidthLimitsResponse() (response *DescribeCcnRegionBandwidthLimitsResponse) {
    response = &DescribeCcnRegionBandwidthLimitsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCcnRegionBandwidthLimits
// This API is used to query the outbound bandwidth caps of all regions connected with a CCN instance. The API only returns regions included in the associated network instances.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
func (c *Client) DescribeCcnRegionBandwidthLimits(request *DescribeCcnRegionBandwidthLimitsRequest) (response *DescribeCcnRegionBandwidthLimitsResponse, err error) {
    if request == nil {
        request = NewDescribeCcnRegionBandwidthLimitsRequest()
    }
    response = NewDescribeCcnRegionBandwidthLimitsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCcnRoutesRequest() (request *DescribeCcnRoutesRequest) {
    request = &DescribeCcnRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeCcnRoutes")
    return
}

func NewDescribeCcnRoutesResponse() (response *DescribeCcnRoutesResponse) {
    response = &DescribeCcnRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCcnRoutes
// This API (DescribeCcnRoutes) is used to query routes that have been added to a CCN.
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCcnRoutes(request *DescribeCcnRoutesRequest) (response *DescribeCcnRoutesResponse, err error) {
    if request == nil {
        request = NewDescribeCcnRoutesRequest()
    }
    response = NewDescribeCcnRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCcnsRequest() (request *DescribeCcnsRequest) {
    request = &DescribeCcnsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeCcns")
    return
}

func NewDescribeCcnsResponse() (response *DescribeCcnsResponse) {
    response = &DescribeCcnsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCcns
// This API (DescribeCcns) is used to query the CCN list.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_FILTERNOTDICT = "InvalidParameter.FilterNotDict"
//  INVALIDPARAMETER_FILTERVALUESNOTLIST = "InvalidParameter.FilterValuesNotList"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCcns(request *DescribeCcnsRequest) (response *DescribeCcnsResponse, err error) {
    if request == nil {
        request = NewDescribeCcnsRequest()
    }
    response = NewDescribeCcnsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClassicLinkInstancesRequest() (request *DescribeClassicLinkInstancesRequest) {
    request = &DescribeClassicLinkInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeClassicLinkInstances")
    return
}

func NewDescribeClassicLinkInstancesResponse() (response *DescribeClassicLinkInstancesResponse) {
    response = &DescribeClassicLinkInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClassicLinkInstances
// This API (DescribeClassicLinkInstances) is used to query the Classiclink instances list.
//
// error code that may be returned:
//  INVALIDPARAMETER_FILTERVALUESNOTLIST = "InvalidParameter.FilterValuesNotList"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
func (c *Client) DescribeClassicLinkInstances(request *DescribeClassicLinkInstancesRequest) (response *DescribeClassicLinkInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeClassicLinkInstancesRequest()
    }
    response = NewDescribeClassicLinkInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCrossBorderComplianceRequest() (request *DescribeCrossBorderComplianceRequest) {
    request = &DescribeCrossBorderComplianceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeCrossBorderCompliance")
    return
}

func NewDescribeCrossBorderComplianceResponse() (response *DescribeCrossBorderComplianceResponse) {
    response = &DescribeCrossBorderComplianceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCrossBorderCompliance
// This API is used to query the compliance review requests created by the user. 
//
// A service provider can query all review requests created by any `APPID` under its account. Other users can only query their own review requests.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCrossBorderCompliance(request *DescribeCrossBorderComplianceRequest) (response *DescribeCrossBorderComplianceResponse, err error) {
    if request == nil {
        request = NewDescribeCrossBorderComplianceRequest()
    }
    response = NewDescribeCrossBorderComplianceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomerGatewayVendorsRequest() (request *DescribeCustomerGatewayVendorsRequest) {
    request = &DescribeCustomerGatewayVendorsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeCustomerGatewayVendors")
    return
}

func NewDescribeCustomerGatewayVendorsResponse() (response *DescribeCustomerGatewayVendorsResponse) {
    response = &DescribeCustomerGatewayVendorsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCustomerGatewayVendors
// This API (DescribeCustomerGatewayVendors) is used to query the information of supported customer gateway vendors.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCustomerGatewayVendors(request *DescribeCustomerGatewayVendorsRequest) (response *DescribeCustomerGatewayVendorsResponse, err error) {
    if request == nil {
        request = NewDescribeCustomerGatewayVendorsRequest()
    }
    response = NewDescribeCustomerGatewayVendorsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomerGatewaysRequest() (request *DescribeCustomerGatewaysRequest) {
    request = &DescribeCustomerGatewaysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeCustomerGateways")
    return
}

func NewDescribeCustomerGatewaysResponse() (response *DescribeCustomerGatewaysResponse) {
    response = &DescribeCustomerGatewaysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCustomerGateways
// This API (DescribeCustomerGateways) is used to query the customer gateway list.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCustomerGateways(request *DescribeCustomerGatewaysRequest) (response *DescribeCustomerGatewaysResponse, err error) {
    if request == nil {
        request = NewDescribeCustomerGatewaysRequest()
    }
    response = NewDescribeCustomerGatewaysResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDirectConnectGatewayCcnRoutesRequest() (request *DescribeDirectConnectGatewayCcnRoutesRequest) {
    request = &DescribeDirectConnectGatewayCcnRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeDirectConnectGatewayCcnRoutes")
    return
}

func NewDescribeDirectConnectGatewayCcnRoutesResponse() (response *DescribeDirectConnectGatewayCcnRoutesResponse) {
    response = &DescribeDirectConnectGatewayCcnRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDirectConnectGatewayCcnRoutes
// This API (DescribeDirectConnectGatewayCcnRoutes) is used to query the CCN routes (IDC IP range) of the Direct Connect gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDirectConnectGatewayCcnRoutes(request *DescribeDirectConnectGatewayCcnRoutesRequest) (response *DescribeDirectConnectGatewayCcnRoutesResponse, err error) {
    if request == nil {
        request = NewDescribeDirectConnectGatewayCcnRoutesRequest()
    }
    response = NewDescribeDirectConnectGatewayCcnRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDirectConnectGatewaysRequest() (request *DescribeDirectConnectGatewaysRequest) {
    request = &DescribeDirectConnectGatewaysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeDirectConnectGateways")
    return
}

func NewDescribeDirectConnectGatewaysResponse() (response *DescribeDirectConnectGatewaysResponse) {
    response = &DescribeDirectConnectGatewaysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDirectConnectGateways
// This API is used to query direct connect gateways.
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_FILTERNOTDICT = "InvalidParameter.FilterNotDict"
//  INVALIDPARAMETER_FILTERVALUESNOTLIST = "InvalidParameter.FilterValuesNotList"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
func (c *Client) DescribeDirectConnectGateways(request *DescribeDirectConnectGatewaysRequest) (response *DescribeDirectConnectGatewaysResponse, err error) {
    if request == nil {
        request = NewDescribeDirectConnectGatewaysRequest()
    }
    response = NewDescribeDirectConnectGatewaysResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlowLogRequest() (request *DescribeFlowLogRequest) {
    request = &DescribeFlowLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeFlowLog")
    return
}

func NewDescribeFlowLogResponse() (response *DescribeFlowLogResponse) {
    response = &DescribeFlowLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFlowLog
// This API is used to query the information of a flow log.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
func (c *Client) DescribeFlowLog(request *DescribeFlowLogRequest) (response *DescribeFlowLogResponse, err error) {
    if request == nil {
        request = NewDescribeFlowLogRequest()
    }
    response = NewDescribeFlowLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlowLogsRequest() (request *DescribeFlowLogsRequest) {
    request = &DescribeFlowLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeFlowLogs")
    return
}

func NewDescribeFlowLogsResponse() (response *DescribeFlowLogsResponse) {
    response = &DescribeFlowLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFlowLogs
// This API is used to query all the flow logs of the current account.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeFlowLogs(request *DescribeFlowLogsRequest) (response *DescribeFlowLogsResponse, err error) {
    if request == nil {
        request = NewDescribeFlowLogsRequest()
    }
    response = NewDescribeFlowLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGatewayFlowMonitorDetailRequest() (request *DescribeGatewayFlowMonitorDetailRequest) {
    request = &DescribeGatewayFlowMonitorDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeGatewayFlowMonitorDetail")
    return
}

func NewDescribeGatewayFlowMonitorDetailResponse() (response *DescribeGatewayFlowMonitorDetailResponse) {
    response = &DescribeGatewayFlowMonitorDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGatewayFlowMonitorDetail
// This API is used to query the traffic monitoring details of the gateway.
//
// * You can only use this API to query a single gateway instance, which means you must pass in only one of `VpnId`, `DirectConnectGatewayId`, `PeeringConnectionId`, or `NatId`.
//
// * If the gateway has traffic, but no data is returned when this API is called, please check whether gateway traffic monitoring has been enabled in the corresponding gateway details page in the console.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeGatewayFlowMonitorDetail(request *DescribeGatewayFlowMonitorDetailRequest) (response *DescribeGatewayFlowMonitorDetailResponse, err error) {
    if request == nil {
        request = NewDescribeGatewayFlowMonitorDetailRequest()
    }
    response = NewDescribeGatewayFlowMonitorDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGatewayFlowQosRequest() (request *DescribeGatewayFlowQosRequest) {
    request = &DescribeGatewayFlowQosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeGatewayFlowQos")
    return
}

func NewDescribeGatewayFlowQosResponse() (response *DescribeGatewayFlowQosResponse) {
    response = &DescribeGatewayFlowQosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGatewayFlowQos
// This API (DescribeGatewayFlowQos) is used to query the QoS bandwidth limit of inbound IP flow in a gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) DescribeGatewayFlowQos(request *DescribeGatewayFlowQosRequest) (response *DescribeGatewayFlowQosResponse, err error) {
    if request == nil {
        request = NewDescribeGatewayFlowQosRequest()
    }
    response = NewDescribeGatewayFlowQosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHaVipsRequest() (request *DescribeHaVipsRequest) {
    request = &DescribeHaVipsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeHaVips")
    return
}

func NewDescribeHaVipsResponse() (response *DescribeHaVipsResponse) {
    response = &DescribeHaVipsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeHaVips
// This API (DescribeHaVips) is used to query the list of highly available virtual IPs (HAVIP).
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
func (c *Client) DescribeHaVips(request *DescribeHaVipsRequest) (response *DescribeHaVipsResponse, err error) {
    if request == nil {
        request = NewDescribeHaVipsRequest()
    }
    response = NewDescribeHaVipsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIpGeolocationDatabaseUrlRequest() (request *DescribeIpGeolocationDatabaseUrlRequest) {
    request = &DescribeIpGeolocationDatabaseUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeIpGeolocationDatabaseUrl")
    return
}

func NewDescribeIpGeolocationDatabaseUrlResponse() (response *DescribeIpGeolocationDatabaseUrlResponse) {
    response = &DescribeIpGeolocationDatabaseUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIpGeolocationDatabaseUrl
// This API is used to obtain the download link of an IP location database.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERCONFLICT = "InvalidParameterConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeIpGeolocationDatabaseUrl(request *DescribeIpGeolocationDatabaseUrlRequest) (response *DescribeIpGeolocationDatabaseUrlResponse, err error) {
    if request == nil {
        request = NewDescribeIpGeolocationDatabaseUrlRequest()
    }
    response = NewDescribeIpGeolocationDatabaseUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIpGeolocationInfosRequest() (request *DescribeIpGeolocationInfosRequest) {
    request = &DescribeIpGeolocationInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeIpGeolocationInfos")
    return
}

func NewDescribeIpGeolocationInfosResponse() (response *DescribeIpGeolocationInfosResponse) {
    response = &DescribeIpGeolocationInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIpGeolocationInfos
// This API is used to query the information of IP addresses, including their geographical locations and networks.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDACCOUNT_NOTSUPPORTED = "InvalidAccount.NotSupported"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeIpGeolocationInfos(request *DescribeIpGeolocationInfosRequest) (response *DescribeIpGeolocationInfosResponse, err error) {
    if request == nil {
        request = NewDescribeIpGeolocationInfosRequest()
    }
    response = NewDescribeIpGeolocationInfosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLocalGatewayRequest() (request *DescribeLocalGatewayRequest) {
    request = &DescribeLocalGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeLocalGateway")
    return
}

func NewDescribeLocalGatewayResponse() (response *DescribeLocalGatewayResponse) {
    response = &DescribeLocalGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLocalGateway
// This API is used to query local gateways of a CDC instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLocalGateway(request *DescribeLocalGatewayRequest) (response *DescribeLocalGatewayResponse, err error) {
    if request == nil {
        request = NewDescribeLocalGatewayRequest()
    }
    response = NewDescribeLocalGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNatGatewayDestinationIpPortTranslationNatRulesRequest() (request *DescribeNatGatewayDestinationIpPortTranslationNatRulesRequest) {
    request = &DescribeNatGatewayDestinationIpPortTranslationNatRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeNatGatewayDestinationIpPortTranslationNatRules")
    return
}

func NewDescribeNatGatewayDestinationIpPortTranslationNatRulesResponse() (response *DescribeNatGatewayDestinationIpPortTranslationNatRulesResponse) {
    response = &DescribeNatGatewayDestinationIpPortTranslationNatRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNatGatewayDestinationIpPortTranslationNatRules
// This API (DescribeNatGatewayDestinationIpPortTranslationNatRules) is used to query the array of objects of the port forwarding rules for a NAT gateway.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDADDRESSID_NOTFOUND = "InvalidAddressId.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeNatGatewayDestinationIpPortTranslationNatRules(request *DescribeNatGatewayDestinationIpPortTranslationNatRulesRequest) (response *DescribeNatGatewayDestinationIpPortTranslationNatRulesResponse, err error) {
    if request == nil {
        request = NewDescribeNatGatewayDestinationIpPortTranslationNatRulesRequest()
    }
    response = NewDescribeNatGatewayDestinationIpPortTranslationNatRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNatGatewaySourceIpTranslationNatRulesRequest() (request *DescribeNatGatewaySourceIpTranslationNatRulesRequest) {
    request = &DescribeNatGatewaySourceIpTranslationNatRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeNatGatewaySourceIpTranslationNatRules")
    return
}

func NewDescribeNatGatewaySourceIpTranslationNatRulesResponse() (response *DescribeNatGatewaySourceIpTranslationNatRulesResponse) {
    response = &DescribeNatGatewaySourceIpTranslationNatRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNatGatewaySourceIpTranslationNatRules
// This API is used to query the object arrays of SNAT forwarding rules of the NAT Gateway.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDADDRESSID_NOTFOUND = "InvalidAddressId.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNatGatewaySourceIpTranslationNatRules(request *DescribeNatGatewaySourceIpTranslationNatRulesRequest) (response *DescribeNatGatewaySourceIpTranslationNatRulesResponse, err error) {
    if request == nil {
        request = NewDescribeNatGatewaySourceIpTranslationNatRulesRequest()
    }
    response = NewDescribeNatGatewaySourceIpTranslationNatRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNatGatewaysRequest() (request *DescribeNatGatewaysRequest) {
    request = &DescribeNatGatewaysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeNatGateways")
    return
}

func NewDescribeNatGatewaysResponse() (response *DescribeNatGatewaysResponse) {
    response = &DescribeNatGatewaysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNatGateways
// This API (DescribeNatGateways) is used to query NAT gateways.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
func (c *Client) DescribeNatGateways(request *DescribeNatGatewaysRequest) (response *DescribeNatGatewaysResponse, err error) {
    if request == nil {
        request = NewDescribeNatGatewaysRequest()
    }
    response = NewDescribeNatGatewaysResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNetDetectStatesRequest() (request *DescribeNetDetectStatesRequest) {
    request = &DescribeNetDetectStatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeNetDetectStates")
    return
}

func NewDescribeNetDetectStatesResponse() (response *DescribeNetDetectStatesResponse) {
    response = &DescribeNetDetectStatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNetDetectStates
// This API (DescribeNetDetectStates) is used to query the list of network detection verification results.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNetDetectStates(request *DescribeNetDetectStatesRequest) (response *DescribeNetDetectStatesResponse, err error) {
    if request == nil {
        request = NewDescribeNetDetectStatesRequest()
    }
    response = NewDescribeNetDetectStatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNetDetectsRequest() (request *DescribeNetDetectsRequest) {
    request = &DescribeNetDetectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeNetDetects")
    return
}

func NewDescribeNetDetectsResponse() (response *DescribeNetDetectsResponse) {
    response = &DescribeNetDetectsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNetDetects
// This API (DescribeNetDetects) is used to query the list of network detection instances.
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNetDetects(request *DescribeNetDetectsRequest) (response *DescribeNetDetectsResponse, err error) {
    if request == nil {
        request = NewDescribeNetDetectsRequest()
    }
    response = NewDescribeNetDetectsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNetworkAclsRequest() (request *DescribeNetworkAclsRequest) {
    request = &DescribeNetworkAclsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeNetworkAcls")
    return
}

func NewDescribeNetworkAclsResponse() (response *DescribeNetworkAclsResponse) {
    response = &DescribeNetworkAclsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNetworkAcls
// This API is used to query a list of network ACLs.
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_ACTIONNOTFOUND = "UnsupportedOperation.ActionNotFound"
func (c *Client) DescribeNetworkAcls(request *DescribeNetworkAclsRequest) (response *DescribeNetworkAclsResponse, err error) {
    if request == nil {
        request = NewDescribeNetworkAclsRequest()
    }
    response = NewDescribeNetworkAclsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNetworkInterfaceLimitRequest() (request *DescribeNetworkInterfaceLimitRequest) {
    request = &DescribeNetworkInterfaceLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeNetworkInterfaceLimit")
    return
}

func NewDescribeNetworkInterfaceLimitResponse() (response *DescribeNetworkInterfaceLimitResponse) {
    response = &DescribeNetworkInterfaceLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNetworkInterfaceLimit
// This API (DescribeNetworkInterfaceLimit) is used to query the ENI quota based on the ID of CVM instance or ENI. It returns the ENI quota to which the CVM instance can be bound and the IP address quota that can be allocated to the ENI.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNetworkInterfaceLimit(request *DescribeNetworkInterfaceLimitRequest) (response *DescribeNetworkInterfaceLimitResponse, err error) {
    if request == nil {
        request = NewDescribeNetworkInterfaceLimitRequest()
    }
    response = NewDescribeNetworkInterfaceLimitResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNetworkInterfacesRequest() (request *DescribeNetworkInterfacesRequest) {
    request = &DescribeNetworkInterfacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeNetworkInterfaces")
    return
}

func NewDescribeNetworkInterfacesResponse() (response *DescribeNetworkInterfacesResponse) {
    response = &DescribeNetworkInterfacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNetworkInterfaces
// This API (DescribeNetworkInterfaces) is used to query the ENI list.
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETER_FILTERNOTDICT = "InvalidParameter.FilterNotDict"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNetworkInterfaces(request *DescribeNetworkInterfacesRequest) (response *DescribeNetworkInterfacesResponse, err error) {
    if request == nil {
        request = NewDescribeNetworkInterfacesRequest()
    }
    response = NewDescribeNetworkInterfacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRouteTablesRequest() (request *DescribeRouteTablesRequest) {
    request = &DescribeRouteTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeRouteTables")
    return
}

func NewDescribeRouteTablesResponse() (response *DescribeRouteTablesResponse) {
    response = &DescribeRouteTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRouteTables
//  This API (DescribeRouteTables) is used to query route tables.
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRouteTables(request *DescribeRouteTablesRequest) (response *DescribeRouteTablesResponse, err error) {
    if request == nil {
        request = NewDescribeRouteTablesRequest()
    }
    response = NewDescribeRouteTablesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityGroupAssociationStatisticsRequest() (request *DescribeSecurityGroupAssociationStatisticsRequest) {
    request = &DescribeSecurityGroupAssociationStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeSecurityGroupAssociationStatistics")
    return
}

func NewDescribeSecurityGroupAssociationStatisticsResponse() (response *DescribeSecurityGroupAssociationStatisticsResponse) {
    response = &DescribeSecurityGroupAssociationStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSecurityGroupAssociationStatistics
// This API (DescribeSecurityGroupAssociationStatistics) is used to query statistics on the instances associated with a security group.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSecurityGroupAssociationStatistics(request *DescribeSecurityGroupAssociationStatisticsRequest) (response *DescribeSecurityGroupAssociationStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityGroupAssociationStatisticsRequest()
    }
    response = NewDescribeSecurityGroupAssociationStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityGroupPoliciesRequest() (request *DescribeSecurityGroupPoliciesRequest) {
    request = &DescribeSecurityGroupPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeSecurityGroupPolicies")
    return
}

func NewDescribeSecurityGroupPoliciesResponse() (response *DescribeSecurityGroupPoliciesResponse) {
    response = &DescribeSecurityGroupPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSecurityGroupPolicies
// This API (DescribeSecurityGroupPolicies) is used to query security group policies.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSecurityGroupPolicies(request *DescribeSecurityGroupPoliciesRequest) (response *DescribeSecurityGroupPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityGroupPoliciesRequest()
    }
    response = NewDescribeSecurityGroupPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityGroupReferencesRequest() (request *DescribeSecurityGroupReferencesRequest) {
    request = &DescribeSecurityGroupReferencesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeSecurityGroupReferences")
    return
}

func NewDescribeSecurityGroupReferencesResponse() (response *DescribeSecurityGroupReferencesResponse) {
    response = &DescribeSecurityGroupReferencesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSecurityGroupReferences
// This API (DescribeSecurityGroupReferences) is used to query referred security groups.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSecurityGroupReferences(request *DescribeSecurityGroupReferencesRequest) (response *DescribeSecurityGroupReferencesResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityGroupReferencesRequest()
    }
    response = NewDescribeSecurityGroupReferencesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityGroupsRequest() (request *DescribeSecurityGroupsRequest) {
    request = &DescribeSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeSecurityGroups")
    return
}

func NewDescribeSecurityGroupsResponse() (response *DescribeSecurityGroupsResponse) {
    response = &DescribeSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSecurityGroups
// This API (DescribeSecurityGroups) is used to query security groups.
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETER_FILTERNOTDICT = "InvalidParameter.FilterNotDict"
//  INVALIDPARAMETER_FILTERVALUESNOTLIST = "InvalidParameter.FilterValuesNotList"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSecurityGroups(request *DescribeSecurityGroupsRequest) (response *DescribeSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityGroupsRequest()
    }
    response = NewDescribeSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceTemplateGroupsRequest() (request *DescribeServiceTemplateGroupsRequest) {
    request = &DescribeServiceTemplateGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeServiceTemplateGroups")
    return
}

func NewDescribeServiceTemplateGroupsResponse() (response *DescribeServiceTemplateGroupsResponse) {
    response = &DescribeServiceTemplateGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeServiceTemplateGroups
// This API (DescribeServiceTemplateGroups) is used to query a protocol port template group.
//
// error code that may be returned:
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
func (c *Client) DescribeServiceTemplateGroups(request *DescribeServiceTemplateGroupsRequest) (response *DescribeServiceTemplateGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeServiceTemplateGroupsRequest()
    }
    response = NewDescribeServiceTemplateGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceTemplatesRequest() (request *DescribeServiceTemplatesRequest) {
    request = &DescribeServiceTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeServiceTemplates")
    return
}

func NewDescribeServiceTemplatesResponse() (response *DescribeServiceTemplatesResponse) {
    response = &DescribeServiceTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeServiceTemplates
// This API (DescribeServiceTemplates) is used to query protocol port templates.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
func (c *Client) DescribeServiceTemplates(request *DescribeServiceTemplatesRequest) (response *DescribeServiceTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeServiceTemplatesRequest()
    }
    response = NewDescribeServiceTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubnetsRequest() (request *DescribeSubnetsRequest) {
    request = &DescribeSubnetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeSubnets")
    return
}

func NewDescribeSubnetsResponse() (response *DescribeSubnetsResponse) {
    response = &DescribeSubnetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSubnets
// This API (DescribeSubnets) is used to query the list of subnets.
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETER_FILTERNOTDICT = "InvalidParameter.FilterNotDict"
//  INVALIDPARAMETER_FILTERVALUESNOTLIST = "InvalidParameter.FilterValuesNotList"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSubnets(request *DescribeSubnetsRequest) (response *DescribeSubnetsResponse, err error) {
    if request == nil {
        request = NewDescribeSubnetsRequest()
    }
    response = NewDescribeSubnetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskResultRequest() (request *DescribeTaskResultRequest) {
    request = &DescribeTaskResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeTaskResult")
    return
}

func NewDescribeTaskResultResponse() (response *DescribeTaskResultResponse) {
    response = &DescribeTaskResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTaskResult
// This API is used to query the EIP async job execution results.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeTaskResult(request *DescribeTaskResultRequest) (response *DescribeTaskResultResponse, err error) {
    if request == nil {
        request = NewDescribeTaskResultRequest()
    }
    response = NewDescribeTaskResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVpcEndPointRequest() (request *DescribeVpcEndPointRequest) {
    request = &DescribeVpcEndPointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpcEndPoint")
    return
}

func NewDescribeVpcEndPointResponse() (response *DescribeVpcEndPointResponse) {
    response = &DescribeVpcEndPointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVpcEndPoint
// This API is used to query the endpoint list.
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVpcEndPoint(request *DescribeVpcEndPointRequest) (response *DescribeVpcEndPointResponse, err error) {
    if request == nil {
        request = NewDescribeVpcEndPointRequest()
    }
    response = NewDescribeVpcEndPointResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVpcEndPointServiceRequest() (request *DescribeVpcEndPointServiceRequest) {
    request = &DescribeVpcEndPointServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpcEndPointService")
    return
}

func NewDescribeVpcEndPointServiceResponse() (response *DescribeVpcEndPointServiceResponse) {
    response = &DescribeVpcEndPointServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVpcEndPointService
// This API is used to query the endpoint service list.
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVpcEndPointService(request *DescribeVpcEndPointServiceRequest) (response *DescribeVpcEndPointServiceResponse, err error) {
    if request == nil {
        request = NewDescribeVpcEndPointServiceRequest()
    }
    response = NewDescribeVpcEndPointServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVpcEndPointServiceWhiteListRequest() (request *DescribeVpcEndPointServiceWhiteListRequest) {
    request = &DescribeVpcEndPointServiceWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpcEndPointServiceWhiteList")
    return
}

func NewDescribeVpcEndPointServiceWhiteListResponse() (response *DescribeVpcEndPointServiceWhiteListResponse) {
    response = &DescribeVpcEndPointServiceWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVpcEndPointServiceWhiteList
// This API is used to query the endpoint service allowlist.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  UNSUPPORTEDOPERATION_UINNOTFOUND = "UnsupportedOperation.UinNotFound"
func (c *Client) DescribeVpcEndPointServiceWhiteList(request *DescribeVpcEndPointServiceWhiteListRequest) (response *DescribeVpcEndPointServiceWhiteListResponse, err error) {
    if request == nil {
        request = NewDescribeVpcEndPointServiceWhiteListRequest()
    }
    response = NewDescribeVpcEndPointServiceWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVpcInstancesRequest() (request *DescribeVpcInstancesRequest) {
    request = &DescribeVpcInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpcInstances")
    return
}

func NewDescribeVpcInstancesResponse() (response *DescribeVpcInstancesResponse) {
    response = &DescribeVpcInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVpcInstances
//  This API (DescribeVpcInstances) is used to query a list of VCM instances on VPC.
//
// error code that may be returned:
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVpcInstances(request *DescribeVpcInstancesRequest) (response *DescribeVpcInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeVpcInstancesRequest()
    }
    response = NewDescribeVpcInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVpcIpv6AddressesRequest() (request *DescribeVpcIpv6AddressesRequest) {
    request = &DescribeVpcIpv6AddressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpcIpv6Addresses")
    return
}

func NewDescribeVpcIpv6AddressesResponse() (response *DescribeVpcIpv6AddressesResponse) {
    response = &DescribeVpcIpv6AddressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVpcIpv6Addresses
// This API (DescribeVpcIpv6Addresses) is used to query `VPC` `IPv6` information.
//
// This API is used to query only the information of `IPv6` addresses that are already in use. When querying IPs that have not yet been used, this API will not report an error, but the IPs will not appear in the returned results.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVpcIpv6Addresses(request *DescribeVpcIpv6AddressesRequest) (response *DescribeVpcIpv6AddressesResponse, err error) {
    if request == nil {
        request = NewDescribeVpcIpv6AddressesRequest()
    }
    response = NewDescribeVpcIpv6AddressesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVpcPrivateIpAddressesRequest() (request *DescribeVpcPrivateIpAddressesRequest) {
    request = &DescribeVpcPrivateIpAddressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpcPrivateIpAddresses")
    return
}

func NewDescribeVpcPrivateIpAddressesResponse() (response *DescribeVpcPrivateIpAddressesResponse) {
    response = &DescribeVpcPrivateIpAddressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVpcPrivateIpAddresses
// This API (DescribeVpcPrivateIpAddresses) is used to query the private IP information of a VPC.<br />
//
// This API is used to query only the information of IP addresses that are already in use. When querying IPs that have not yet been used, this API will not report an error, but the IPs will not appear in the returned results.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVpcPrivateIpAddresses(request *DescribeVpcPrivateIpAddressesRequest) (response *DescribeVpcPrivateIpAddressesResponse, err error) {
    if request == nil {
        request = NewDescribeVpcPrivateIpAddressesRequest()
    }
    response = NewDescribeVpcPrivateIpAddressesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVpcResourceDashboardRequest() (request *DescribeVpcResourceDashboardRequest) {
    request = &DescribeVpcResourceDashboardRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpcResourceDashboard")
    return
}

func NewDescribeVpcResourceDashboardResponse() (response *DescribeVpcResourceDashboardResponse) {
    response = &DescribeVpcResourceDashboardResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVpcResourceDashboard
// View VPC resources.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVpcResourceDashboard(request *DescribeVpcResourceDashboardRequest) (response *DescribeVpcResourceDashboardResponse, err error) {
    if request == nil {
        request = NewDescribeVpcResourceDashboardRequest()
    }
    response = NewDescribeVpcResourceDashboardResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVpcsRequest() (request *DescribeVpcsRequest) {
    request = &DescribeVpcsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpcs")
    return
}

func NewDescribeVpcsResponse() (response *DescribeVpcsResponse) {
    response = &DescribeVpcsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVpcs
// This API (DescribeVpcs) is used to query the VPC list.
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETER_FILTERVALUESNOTLIST = "InvalidParameter.FilterValuesNotList"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVpcs(request *DescribeVpcsRequest) (response *DescribeVpcsResponse, err error) {
    if request == nil {
        request = NewDescribeVpcsRequest()
    }
    response = NewDescribeVpcsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVpnConnectionsRequest() (request *DescribeVpnConnectionsRequest) {
    request = &DescribeVpnConnectionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpnConnections")
    return
}

func NewDescribeVpnConnectionsResponse() (response *DescribeVpnConnectionsResponse) {
    response = &DescribeVpnConnectionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVpnConnections
//  This API (DescribeVpnConnections) is used to query the VPN tunnel list.
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVpnConnections(request *DescribeVpnConnectionsRequest) (response *DescribeVpnConnectionsResponse, err error) {
    if request == nil {
        request = NewDescribeVpnConnectionsRequest()
    }
    response = NewDescribeVpnConnectionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVpnGatewayCcnRoutesRequest() (request *DescribeVpnGatewayCcnRoutesRequest) {
    request = &DescribeVpnGatewayCcnRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpnGatewayCcnRoutes")
    return
}

func NewDescribeVpnGatewayCcnRoutesResponse() (response *DescribeVpnGatewayCcnRoutesResponse) {
    response = &DescribeVpnGatewayCcnRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVpnGatewayCcnRoutes
// This API (DescribeVpnGatewayCcnRoutes) is used to query VPN gateway-based CCN routes.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVpnGatewayCcnRoutes(request *DescribeVpnGatewayCcnRoutesRequest) (response *DescribeVpnGatewayCcnRoutesResponse, err error) {
    if request == nil {
        request = NewDescribeVpnGatewayCcnRoutesRequest()
    }
    response = NewDescribeVpnGatewayCcnRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVpnGatewaysRequest() (request *DescribeVpnGatewaysRequest) {
    request = &DescribeVpnGatewaysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpnGateways")
    return
}

func NewDescribeVpnGatewaysResponse() (response *DescribeVpnGatewaysResponse) {
    response = &DescribeVpnGatewaysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVpnGateways
// This API (DescribeVpnGateways) is used to query the VPN gateway list.
//
// error code that may be returned:
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETER_FILTERNOTDICT = "InvalidParameter.FilterNotDict"
//  INVALIDPARAMETER_FILTERVALUESNOTLIST = "InvalidParameter.FilterValuesNotList"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDVPNGATEWAYID_MALFORMED = "InvalidVpnGatewayId.Malformed"
//  INVALIDVPNGATEWAYID_NOTFOUND = "InvalidVpnGatewayId.NotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVpnGateways(request *DescribeVpnGatewaysRequest) (response *DescribeVpnGatewaysResponse, err error) {
    if request == nil {
        request = NewDescribeVpnGatewaysRequest()
    }
    response = NewDescribeVpnGatewaysResponse()
    err = c.Send(request, response)
    return
}

func NewDetachCcnInstancesRequest() (request *DetachCcnInstancesRequest) {
    request = &DetachCcnInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DetachCcnInstances")
    return
}

func NewDetachCcnInstancesResponse() (response *DetachCcnInstancesResponse) {
    response = &DetachCcnInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetachCcnInstances
// This API (DetachCcnInstances) is used to unbind a specified network instance from a CCN instance.<br />
//
// After unbinding the network instance, the corresponding routing policy will also be deleted.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_CCNNOTATTACHED = "UnsupportedOperation.CcnNotAttached"
func (c *Client) DetachCcnInstances(request *DetachCcnInstancesRequest) (response *DetachCcnInstancesResponse, err error) {
    if request == nil {
        request = NewDetachCcnInstancesRequest()
    }
    response = NewDetachCcnInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDetachClassicLinkVpcRequest() (request *DetachClassicLinkVpcRequest) {
    request = &DetachClassicLinkVpcRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DetachClassicLinkVpc")
    return
}

func NewDetachClassicLinkVpcResponse() (response *DetachClassicLinkVpcResponse) {
    response = &DetachClassicLinkVpcResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetachClassicLinkVpc
// This API (DetachClassicLinkVpc) is used to delete a Classiclink.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DetachClassicLinkVpc(request *DetachClassicLinkVpcRequest) (response *DetachClassicLinkVpcResponse, err error) {
    if request == nil {
        request = NewDetachClassicLinkVpcRequest()
    }
    response = NewDetachClassicLinkVpcResponse()
    err = c.Send(request, response)
    return
}

func NewDetachNetworkInterfaceRequest() (request *DetachNetworkInterfaceRequest) {
    request = &DetachNetworkInterfaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DetachNetworkInterface")
    return
}

func NewDetachNetworkInterfaceResponse() (response *DetachNetworkInterfaceResponse) {
    response = &DetachNetworkInterfaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetachNetworkInterface
// This API is used to unbind an ENI from a CVM.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) DetachNetworkInterface(request *DetachNetworkInterfaceRequest) (response *DetachNetworkInterfaceResponse, err error) {
    if request == nil {
        request = NewDetachNetworkInterfaceRequest()
    }
    response = NewDetachNetworkInterfaceResponse()
    err = c.Send(request, response)
    return
}

func NewDisableCcnRoutesRequest() (request *DisableCcnRoutesRequest) {
    request = &DisableCcnRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DisableCcnRoutes")
    return
}

func NewDisableCcnRoutesResponse() (response *DisableCcnRoutesResponse) {
    response = &DisableCcnRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisableCcnRoutes
// This API (DisableCcnRoutes) is used to disable CCN routes that are already enabled.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DisableCcnRoutes(request *DisableCcnRoutesRequest) (response *DisableCcnRoutesResponse, err error) {
    if request == nil {
        request = NewDisableCcnRoutesRequest()
    }
    response = NewDisableCcnRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewDisableGatewayFlowMonitorRequest() (request *DisableGatewayFlowMonitorRequest) {
    request = &DisableGatewayFlowMonitorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DisableGatewayFlowMonitor")
    return
}

func NewDisableGatewayFlowMonitorResponse() (response *DisableGatewayFlowMonitorResponse) {
    response = &DisableGatewayFlowMonitorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisableGatewayFlowMonitor
// This API (DisableGatewayFlowMonitor) is used to disable gateway flow monitor.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) DisableGatewayFlowMonitor(request *DisableGatewayFlowMonitorRequest) (response *DisableGatewayFlowMonitorResponse, err error) {
    if request == nil {
        request = NewDisableGatewayFlowMonitorRequest()
    }
    response = NewDisableGatewayFlowMonitorResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateAddressRequest() (request *DisassociateAddressRequest) {
    request = &DisassociateAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DisassociateAddress")
    return
}

func NewDisassociateAddressResponse() (response *DisassociateAddressResponse) {
    response = &DisassociateAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisassociateAddress
// This API is used to unbind an [Elastic IP](https://intl.cloud.tencent.com/document/product/213/1941?from_cn_redirect=1) (EIP for short).
//
// * This API supports unbinding an EIP from CVM instances and ENIs.
//
// * This API does not support unbinding an EIP from a NAT Gateway. To unbind an EIP from a NAT Gateway, use the [`DisassociateNatGatewayAddress`](https://intl.cloud.tencent.com/document/api/215/36716?from_cn_redirect=1) API.
//
// * Only EIPs in BIND or BIND_ENI status can be unbound.
//
// * Blocked EIPs cannot be unbound.
//
// error code that may be returned:
//  ADDRESSQUOTALIMITEXCEEDED_DAILYALLOCATE = "AddressQuotaLimitExceeded.DailyAllocate"
//  INVALIDADDRESSID_BLOCKED = "InvalidAddressId.Blocked"
//  INVALIDADDRESSID_NOTFOUND = "InvalidAddressId.NotFound"
//  INVALIDADDRESSIDSTATUS_NOTPERMIT = "InvalidAddressIdStatus.NotPermit"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ONLYSUPPORTEDFORMASTERNETWORKCARD = "InvalidParameterValue.OnlySupportedForMasterNetworkCard"
func (c *Client) DisassociateAddress(request *DisassociateAddressRequest) (response *DisassociateAddressResponse, err error) {
    if request == nil {
        request = NewDisassociateAddressRequest()
    }
    response = NewDisassociateAddressResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateDirectConnectGatewayNatGatewayRequest() (request *DisassociateDirectConnectGatewayNatGatewayRequest) {
    request = &DisassociateDirectConnectGatewayNatGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DisassociateDirectConnectGatewayNatGateway")
    return
}

func NewDisassociateDirectConnectGatewayNatGatewayResponse() (response *DisassociateDirectConnectGatewayNatGatewayResponse) {
    response = &DisassociateDirectConnectGatewayNatGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisassociateDirectConnectGatewayNatGateway
// This API is used to unbind a direct connect gateway from a NAT Gateway. After unbinding, the direct connect gateway cannot access internet through the NAT Gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DisassociateDirectConnectGatewayNatGateway(request *DisassociateDirectConnectGatewayNatGatewayRequest) (response *DisassociateDirectConnectGatewayNatGatewayResponse, err error) {
    if request == nil {
        request = NewDisassociateDirectConnectGatewayNatGatewayRequest()
    }
    response = NewDisassociateDirectConnectGatewayNatGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateNatGatewayAddressRequest() (request *DisassociateNatGatewayAddressRequest) {
    request = &DisassociateNatGatewayAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DisassociateNatGatewayAddress")
    return
}

func NewDisassociateNatGatewayAddressResponse() (response *DisassociateNatGatewayAddressResponse) {
    response = &DisassociateNatGatewayAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisassociateNatGatewayAddress
// This API (DisassociateNatGatewayAddress) is used to unbind an EIP from a NAT gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_PUBLICIPADDRESSDISASSOCIATE = "UnsupportedOperation.PublicIpAddressDisassociate"
func (c *Client) DisassociateNatGatewayAddress(request *DisassociateNatGatewayAddressRequest) (response *DisassociateNatGatewayAddressResponse, err error) {
    if request == nil {
        request = NewDisassociateNatGatewayAddressRequest()
    }
    response = NewDisassociateNatGatewayAddressResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateNetworkAclSubnetsRequest() (request *DisassociateNetworkAclSubnetsRequest) {
    request = &DisassociateNetworkAclSubnetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DisassociateNetworkAclSubnets")
    return
}

func NewDisassociateNetworkAclSubnetsResponse() (response *DisassociateNetworkAclSubnetsResponse) {
    response = &DisassociateNetworkAclSubnetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisassociateNetworkAclSubnets
// This API is used to disassociate a network ACL from subnets in a VPC instance.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
func (c *Client) DisassociateNetworkAclSubnets(request *DisassociateNetworkAclSubnetsRequest) (response *DisassociateNetworkAclSubnetsResponse, err error) {
    if request == nil {
        request = NewDisassociateNetworkAclSubnetsRequest()
    }
    response = NewDisassociateNetworkAclSubnetsResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateNetworkInterfaceSecurityGroupsRequest() (request *DisassociateNetworkInterfaceSecurityGroupsRequest) {
    request = &DisassociateNetworkInterfaceSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DisassociateNetworkInterfaceSecurityGroups")
    return
}

func NewDisassociateNetworkInterfaceSecurityGroupsResponse() (response *DisassociateNetworkInterfaceSecurityGroupsResponse) {
    response = &DisassociateNetworkInterfaceSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisassociateNetworkInterfaceSecurityGroups
// This API (DisassociateNetworkInterfaceSecurityGroups) is used to detach (or fully detach if possible) a security group from an ENI.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DisassociateNetworkInterfaceSecurityGroups(request *DisassociateNetworkInterfaceSecurityGroupsRequest) (response *DisassociateNetworkInterfaceSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDisassociateNetworkInterfaceSecurityGroupsRequest()
    }
    response = NewDisassociateNetworkInterfaceSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateVpcEndPointSecurityGroupsRequest() (request *DisassociateVpcEndPointSecurityGroupsRequest) {
    request = &DisassociateVpcEndPointSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DisassociateVpcEndPointSecurityGroups")
    return
}

func NewDisassociateVpcEndPointSecurityGroupsResponse() (response *DisassociateVpcEndPointSecurityGroupsResponse) {
    response = &DisassociateVpcEndPointSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisassociateVpcEndPointSecurityGroups
// This API is used to unbind an endpoint from a security group.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DisassociateVpcEndPointSecurityGroups(request *DisassociateVpcEndPointSecurityGroupsRequest) (response *DisassociateVpcEndPointSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDisassociateVpcEndPointSecurityGroupsRequest()
    }
    response = NewDisassociateVpcEndPointSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadCustomerGatewayConfigurationRequest() (request *DownloadCustomerGatewayConfigurationRequest) {
    request = &DownloadCustomerGatewayConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DownloadCustomerGatewayConfiguration")
    return
}

func NewDownloadCustomerGatewayConfigurationResponse() (response *DownloadCustomerGatewayConfigurationResponse) {
    response = &DownloadCustomerGatewayConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DownloadCustomerGatewayConfiguration
// This API (DownloadCustomerGatewayConfiguration) is used to download a VPN tunnel configuration.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DownloadCustomerGatewayConfiguration(request *DownloadCustomerGatewayConfigurationRequest) (response *DownloadCustomerGatewayConfigurationResponse, err error) {
    if request == nil {
        request = NewDownloadCustomerGatewayConfigurationRequest()
    }
    response = NewDownloadCustomerGatewayConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewEnableCcnRoutesRequest() (request *EnableCcnRoutesRequest) {
    request = &EnableCcnRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "EnableCcnRoutes")
    return
}

func NewEnableCcnRoutesResponse() (response *EnableCcnRoutesResponse) {
    response = &EnableCcnRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EnableCcnRoutes
// This API (EnableCcnRoutes) is used to enable CCN routes that are already added.<br />
//
// This API is used to verify whether there will be conflict with an existing route after a CCN route is enabled. If there is a conflict, the route will not be enabled, and the process will fail. When a conflict occurs, you must disable the conflicting route before you can enable the desired route.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_ECMP = "UnsupportedOperation.Ecmp"
func (c *Client) EnableCcnRoutes(request *EnableCcnRoutesRequest) (response *EnableCcnRoutesResponse, err error) {
    if request == nil {
        request = NewEnableCcnRoutesRequest()
    }
    response = NewEnableCcnRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewEnableGatewayFlowMonitorRequest() (request *EnableGatewayFlowMonitorRequest) {
    request = &EnableGatewayFlowMonitorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "EnableGatewayFlowMonitor")
    return
}

func NewEnableGatewayFlowMonitorResponse() (response *EnableGatewayFlowMonitorResponse) {
    response = &EnableGatewayFlowMonitorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EnableGatewayFlowMonitor
// This API (EnableGatewayFlowMonitor) is used to enable gateway flow monitor.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) EnableGatewayFlowMonitor(request *EnableGatewayFlowMonitorRequest) (response *EnableGatewayFlowMonitorResponse, err error) {
    if request == nil {
        request = NewEnableGatewayFlowMonitorRequest()
    }
    response = NewEnableGatewayFlowMonitorResponse()
    err = c.Send(request, response)
    return
}

func NewEnableVpcEndPointConnectRequest() (request *EnableVpcEndPointConnectRequest) {
    request = &EnableVpcEndPointConnectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "EnableVpcEndPointConnect")
    return
}

func NewEnableVpcEndPointConnectResponse() (response *EnableVpcEndPointConnectResponse) {
    response = &EnableVpcEndPointConnectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EnableVpcEndPointConnect
// This API is used to determine whether to accept the request of connecting with an endpoint.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) EnableVpcEndPointConnect(request *EnableVpcEndPointConnectRequest) (response *EnableVpcEndPointConnectResponse, err error) {
    if request == nil {
        request = NewEnableVpcEndPointConnectRequest()
    }
    response = NewEnableVpcEndPointConnectResponse()
    err = c.Send(request, response)
    return
}

func NewGetCcnRegionBandwidthLimitsRequest() (request *GetCcnRegionBandwidthLimitsRequest) {
    request = &GetCcnRegionBandwidthLimitsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "GetCcnRegionBandwidthLimits")
    return
}

func NewGetCcnRegionBandwidthLimitsResponse() (response *GetCcnRegionBandwidthLimitsResponse) {
    response = &GetCcnRegionBandwidthLimitsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetCcnRegionBandwidthLimits
// This API is used to query the bandwidth limits of a CCN instance. Monthly-subscribed CCNs only support Inter-region Bandwidth Limits, while the pay-as-you-go CCNs support both the Inter-region Bandwidth Limits and Region Outbound Bandwidth Limits. 
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetCcnRegionBandwidthLimits(request *GetCcnRegionBandwidthLimitsRequest) (response *GetCcnRegionBandwidthLimitsResponse, err error) {
    if request == nil {
        request = NewGetCcnRegionBandwidthLimitsRequest()
    }
    response = NewGetCcnRegionBandwidthLimitsResponse()
    err = c.Send(request, response)
    return
}

func NewHaVipAssociateAddressIpRequest() (request *HaVipAssociateAddressIpRequest) {
    request = &HaVipAssociateAddressIpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "HaVipAssociateAddressIp")
    return
}

func NewHaVipAssociateAddressIpResponse() (response *HaVipAssociateAddressIpResponse) {
    response = &HaVipAssociateAddressIpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// HaVipAssociateAddressIp
// This API (HaVipAssociateAddressIp) is used to bind an EIP to an HAVIP.<br />
//
// This API is completed asynchronously. If you need to query the async job execution results, please use the `RequestId` returned by this API to query the `QueryTask` API.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_BINDEIP = "UnsupportedOperation.BindEIP"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) HaVipAssociateAddressIp(request *HaVipAssociateAddressIpRequest) (response *HaVipAssociateAddressIpResponse, err error) {
    if request == nil {
        request = NewHaVipAssociateAddressIpRequest()
    }
    response = NewHaVipAssociateAddressIpResponse()
    err = c.Send(request, response)
    return
}

func NewHaVipDisassociateAddressIpRequest() (request *HaVipDisassociateAddressIpRequest) {
    request = &HaVipDisassociateAddressIpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "HaVipDisassociateAddressIp")
    return
}

func NewHaVipDisassociateAddressIpResponse() (response *HaVipDisassociateAddressIpResponse) {
    response = &HaVipDisassociateAddressIpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// HaVipDisassociateAddressIp
// This API (HaVipDisassociateAddressIp) is used to unbind an EIP which has been bound to an HAVIP.<br />
//
// This API is completed asynchronously. If you need to query the async job execution results, please use the `RequestId` returned by this API to query the `QueryTask` API.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
//  UNSUPPORTEDOPERATION_UNBINDEIP = "UnsupportedOperation.UnbindEIP"
func (c *Client) HaVipDisassociateAddressIp(request *HaVipDisassociateAddressIpRequest) (response *HaVipDisassociateAddressIpResponse, err error) {
    if request == nil {
        request = NewHaVipDisassociateAddressIpRequest()
    }
    response = NewHaVipDisassociateAddressIpResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceCreateDirectConnectGatewayRequest() (request *InquirePriceCreateDirectConnectGatewayRequest) {
    request = &InquirePriceCreateDirectConnectGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "InquirePriceCreateDirectConnectGateway")
    return
}

func NewInquirePriceCreateDirectConnectGatewayResponse() (response *InquirePriceCreateDirectConnectGatewayResponse) {
    response = &InquirePriceCreateDirectConnectGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquirePriceCreateDirectConnectGateway
// This API is used to query the price of creating a direct connect gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
//  UNSUPPORTEDOPERATION_UNBINDEIP = "UnsupportedOperation.UnbindEIP"
func (c *Client) InquirePriceCreateDirectConnectGateway(request *InquirePriceCreateDirectConnectGatewayRequest) (response *InquirePriceCreateDirectConnectGatewayResponse, err error) {
    if request == nil {
        request = NewInquirePriceCreateDirectConnectGatewayRequest()
    }
    response = NewInquirePriceCreateDirectConnectGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceCreateVpnGatewayRequest() (request *InquiryPriceCreateVpnGatewayRequest) {
    request = &InquiryPriceCreateVpnGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "InquiryPriceCreateVpnGateway")
    return
}

func NewInquiryPriceCreateVpnGatewayResponse() (response *InquiryPriceCreateVpnGatewayResponse) {
    response = &InquiryPriceCreateVpnGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquiryPriceCreateVpnGateway
// This API (InquiryPriceCreateVpnGateway) is used to query the price for VPN gateway creation.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) InquiryPriceCreateVpnGateway(request *InquiryPriceCreateVpnGatewayRequest) (response *InquiryPriceCreateVpnGatewayResponse, err error) {
    if request == nil {
        request = NewInquiryPriceCreateVpnGatewayRequest()
    }
    response = NewInquiryPriceCreateVpnGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceRenewVpnGatewayRequest() (request *InquiryPriceRenewVpnGatewayRequest) {
    request = &InquiryPriceRenewVpnGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "InquiryPriceRenewVpnGateway")
    return
}

func NewInquiryPriceRenewVpnGatewayResponse() (response *InquiryPriceRenewVpnGatewayResponse) {
    response = &InquiryPriceRenewVpnGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquiryPriceRenewVpnGateway
// This API (InquiryPriceRenewVpnGateway) is used to query the price for VPN gateway renewal. Currently, only querying prices for IPSEC-type gateways is supported.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) InquiryPriceRenewVpnGateway(request *InquiryPriceRenewVpnGatewayRequest) (response *InquiryPriceRenewVpnGatewayResponse, err error) {
    if request == nil {
        request = NewInquiryPriceRenewVpnGatewayRequest()
    }
    response = NewInquiryPriceRenewVpnGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceResetVpnGatewayInternetMaxBandwidthRequest() (request *InquiryPriceResetVpnGatewayInternetMaxBandwidthRequest) {
    request = &InquiryPriceResetVpnGatewayInternetMaxBandwidthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "InquiryPriceResetVpnGatewayInternetMaxBandwidth")
    return
}

func NewInquiryPriceResetVpnGatewayInternetMaxBandwidthResponse() (response *InquiryPriceResetVpnGatewayInternetMaxBandwidthResponse) {
    response = &InquiryPriceResetVpnGatewayInternetMaxBandwidthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquiryPriceResetVpnGatewayInternetMaxBandwidth
// This API (InquiryPriceResetVpnGatewayInternetMaxBandwidth) is used to query the price for adjusting the bandwidth cap of a VPN gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) InquiryPriceResetVpnGatewayInternetMaxBandwidth(request *InquiryPriceResetVpnGatewayInternetMaxBandwidthRequest) (response *InquiryPriceResetVpnGatewayInternetMaxBandwidthResponse, err error) {
    if request == nil {
        request = NewInquiryPriceResetVpnGatewayInternetMaxBandwidthRequest()
    }
    response = NewInquiryPriceResetVpnGatewayInternetMaxBandwidthResponse()
    err = c.Send(request, response)
    return
}

func NewMigrateNetworkInterfaceRequest() (request *MigrateNetworkInterfaceRequest) {
    request = &MigrateNetworkInterfaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "MigrateNetworkInterface")
    return
}

func NewMigrateNetworkInterfaceResponse() (response *MigrateNetworkInterfaceResponse) {
    response = &MigrateNetworkInterfaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// MigrateNetworkInterface
// This API (MigrateNetworkInterface) is used to migrate ENIs.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
func (c *Client) MigrateNetworkInterface(request *MigrateNetworkInterfaceRequest) (response *MigrateNetworkInterfaceResponse, err error) {
    if request == nil {
        request = NewMigrateNetworkInterfaceRequest()
    }
    response = NewMigrateNetworkInterfaceResponse()
    err = c.Send(request, response)
    return
}

func NewMigratePrivateIpAddressRequest() (request *MigratePrivateIpAddressRequest) {
    request = &MigratePrivateIpAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "MigratePrivateIpAddress")
    return
}

func NewMigratePrivateIpAddressResponse() (response *MigratePrivateIpAddressResponse) {
    response = &MigratePrivateIpAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// MigratePrivateIpAddress
//  This API (MigratePrivateIpAddress) is used to migrate the private IPs of ENIs.
//
// 
//
// * This API is used to migrate a private IP from one ENI to another. Primary IPs cannot be migrated.
//
// * The ENIs before and after migration must belong to the same subnet.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_ATTACHMENTNOTFOUND = "UnauthorizedOperation.AttachmentNotFound"
//  UNAUTHORIZEDOPERATION_PRIMARYIP = "UnauthorizedOperation.PrimaryIp"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
//  UNSUPPORTEDOPERATION_PRIMARYIP = "UnsupportedOperation.PrimaryIp"
func (c *Client) MigratePrivateIpAddress(request *MigratePrivateIpAddressRequest) (response *MigratePrivateIpAddressResponse, err error) {
    if request == nil {
        request = NewMigratePrivateIpAddressRequest()
    }
    response = NewMigratePrivateIpAddressResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAddressAttributeRequest() (request *ModifyAddressAttributeRequest) {
    request = &ModifyAddressAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyAddressAttribute")
    return
}

func NewModifyAddressAttributeResponse() (response *ModifyAddressAttributeResponse) {
    response = &ModifyAddressAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAddressAttribute
// This API (ModifyAddressAttribute) is used to modify the name of an [Elastic IP](https://intl.cloud.tencent.com/document/product/213/1941?from_cn_redirect=1).
//
// error code that may be returned:
//  INVALIDADDRESSID_BLOCKED = "InvalidAddressId.Blocked"
//  INVALIDADDRESSID_NOTFOUND = "InvalidAddressId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
func (c *Client) ModifyAddressAttribute(request *ModifyAddressAttributeRequest) (response *ModifyAddressAttributeResponse, err error) {
    if request == nil {
        request = NewModifyAddressAttributeRequest()
    }
    response = NewModifyAddressAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAddressInternetChargeTypeRequest() (request *ModifyAddressInternetChargeTypeRequest) {
    request = &ModifyAddressInternetChargeTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyAddressInternetChargeType")
    return
}

func NewModifyAddressInternetChargeTypeResponse() (response *ModifyAddressInternetChargeTypeResponse) {
    response = &ModifyAddressInternetChargeTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAddressInternetChargeType
// This API is used to adjust the network billing mode of an EIP. Please note that it's available to users whose network fees are billed on IPs but not CVMs.
//
// * The network billing mode can be switched between `BANDWIDTH_PREPAID_BY_MONTH` and `TRAFFIC_POSTPAID_BY_HOUR`.
//
// * The network billing mode for each EIP be changed for up to twice.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDACCOUNT_NOTSUPPORTED = "InvalidAccount.NotSupported"
//  INVALIDADDRESSID_NOTFOUND = "InvalidAddressId.NotFound"
//  INVALIDADDRESSIDSTATE_INARREARS = "InvalidAddressIdState.InArrears"
//  INVALIDADDRESSSTATE = "InvalidAddressState"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) ModifyAddressInternetChargeType(request *ModifyAddressInternetChargeTypeRequest) (response *ModifyAddressInternetChargeTypeResponse, err error) {
    if request == nil {
        request = NewModifyAddressInternetChargeTypeRequest()
    }
    response = NewModifyAddressInternetChargeTypeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAddressTemplateAttributeRequest() (request *ModifyAddressTemplateAttributeRequest) {
    request = &ModifyAddressTemplateAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyAddressTemplateAttribute")
    return
}

func NewModifyAddressTemplateAttributeResponse() (response *ModifyAddressTemplateAttributeResponse) {
    response = &ModifyAddressTemplateAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAddressTemplateAttribute
// This API (ModifyAddressTemplateAttribute) is used to modify an IP address template.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) ModifyAddressTemplateAttribute(request *ModifyAddressTemplateAttributeRequest) (response *ModifyAddressTemplateAttributeResponse, err error) {
    if request == nil {
        request = NewModifyAddressTemplateAttributeRequest()
    }
    response = NewModifyAddressTemplateAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAddressTemplateGroupAttributeRequest() (request *ModifyAddressTemplateGroupAttributeRequest) {
    request = &ModifyAddressTemplateGroupAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyAddressTemplateGroupAttribute")
    return
}

func NewModifyAddressTemplateGroupAttributeResponse() (response *ModifyAddressTemplateGroupAttributeResponse) {
    response = &ModifyAddressTemplateGroupAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAddressTemplateGroupAttribute
// This API (ModifyAddressTemplateGroupAttribute) is used to modify an IP address template group.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) ModifyAddressTemplateGroupAttribute(request *ModifyAddressTemplateGroupAttributeRequest) (response *ModifyAddressTemplateGroupAttributeResponse, err error) {
    if request == nil {
        request = NewModifyAddressTemplateGroupAttributeRequest()
    }
    response = NewModifyAddressTemplateGroupAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAddressesBandwidthRequest() (request *ModifyAddressesBandwidthRequest) {
    request = &ModifyAddressesBandwidthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyAddressesBandwidth")
    return
}

func NewModifyAddressesBandwidthResponse() (response *ModifyAddressesBandwidthResponse) {
    response = &ModifyAddressesBandwidthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAddressesBandwidth
// This API is used to adjust the bandwidth of [Elastic IP](https://intl.cloud.tencent.com/document/product/213/1941?from_cn_redirect=1), including EIP billed on a pay-as-you-go, monthly subscription, and bandwidth package basis.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDADDRESSID_NOTFOUND = "InvalidAddressId.NotFound"
//  INVALIDPARAMETERVALUE_BANDWIDTHOUTOFRANGE = "InvalidParameterValue.BandwidthOutOfRange"
func (c *Client) ModifyAddressesBandwidth(request *ModifyAddressesBandwidthRequest) (response *ModifyAddressesBandwidthResponse, err error) {
    if request == nil {
        request = NewModifyAddressesBandwidthRequest()
    }
    response = NewModifyAddressesBandwidthResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAssistantCidrRequest() (request *ModifyAssistantCidrRequest) {
    request = &ModifyAssistantCidrRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyAssistantCidr")
    return
}

func NewModifyAssistantCidrResponse() (response *ModifyAssistantCidrResponse) {
    response = &ModifyAssistantCidrResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAssistantCidr
// This API (ModifyAssistantCidr) is used to batch modify (e.g. add and delete) secondary CIDR blocks. (To use this API that is in Beta, please submit a ticket.)
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_SUBNETCONFLICT = "InvalidParameterValue.SubnetConflict"
//  INVALIDPARAMETERVALUE_VPCCIDRCONFLICT = "InvalidParameterValue.VpcCidrConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAssistantCidr(request *ModifyAssistantCidrRequest) (response *ModifyAssistantCidrResponse, err error) {
    if request == nil {
        request = NewModifyAssistantCidrRequest()
    }
    response = NewModifyAssistantCidrResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBandwidthPackageAttributeRequest() (request *ModifyBandwidthPackageAttributeRequest) {
    request = &ModifyBandwidthPackageAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyBandwidthPackageAttribute")
    return
}

func NewModifyBandwidthPackageAttributeResponse() (response *ModifyBandwidthPackageAttributeResponse) {
    response = &ModifyBandwidthPackageAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyBandwidthPackageAttribute
// This API is used to modify the attributes of a bandwidth package, including the bandwidth package name, and so on.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
func (c *Client) ModifyBandwidthPackageAttribute(request *ModifyBandwidthPackageAttributeRequest) (response *ModifyBandwidthPackageAttributeResponse, err error) {
    if request == nil {
        request = NewModifyBandwidthPackageAttributeRequest()
    }
    response = NewModifyBandwidthPackageAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCcnAttachedInstancesAttributeRequest() (request *ModifyCcnAttachedInstancesAttributeRequest) {
    request = &ModifyCcnAttachedInstancesAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyCcnAttachedInstancesAttribute")
    return
}

func NewModifyCcnAttachedInstancesAttributeResponse() (response *ModifyCcnAttachedInstancesAttributeResponse) {
    response = &ModifyCcnAttachedInstancesAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCcnAttachedInstancesAttribute
// This API is used to modify CCN-associated instance attributes. Currently, only the `description` can be modified.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCcnAttachedInstancesAttribute(request *ModifyCcnAttachedInstancesAttributeRequest) (response *ModifyCcnAttachedInstancesAttributeResponse, err error) {
    if request == nil {
        request = NewModifyCcnAttachedInstancesAttributeRequest()
    }
    response = NewModifyCcnAttachedInstancesAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCcnAttributeRequest() (request *ModifyCcnAttributeRequest) {
    request = &ModifyCcnAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyCcnAttribute")
    return
}

func NewModifyCcnAttributeResponse() (response *ModifyCcnAttributeResponse) {
    response = &ModifyCcnAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCcnAttribute
// This API (ModifyCcnAttribute) is used to modify CCN attributes.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyCcnAttribute(request *ModifyCcnAttributeRequest) (response *ModifyCcnAttributeResponse, err error) {
    if request == nil {
        request = NewModifyCcnAttributeRequest()
    }
    response = NewModifyCcnAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCcnRegionBandwidthLimitsTypeRequest() (request *ModifyCcnRegionBandwidthLimitsTypeRequest) {
    request = &ModifyCcnRegionBandwidthLimitsTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyCcnRegionBandwidthLimitsType")
    return
}

func NewModifyCcnRegionBandwidthLimitsTypeResponse() (response *ModifyCcnRegionBandwidthLimitsTypeResponse) {
    response = &ModifyCcnRegionBandwidthLimitsTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCcnRegionBandwidthLimitsType
// This API is used to modify the bandwidth limit policy of a postpaid CCN instance.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_NOTPOSTPAIDCCNOPERATION = "UnsupportedOperation.NotPostpaidCcnOperation"
func (c *Client) ModifyCcnRegionBandwidthLimitsType(request *ModifyCcnRegionBandwidthLimitsTypeRequest) (response *ModifyCcnRegionBandwidthLimitsTypeResponse, err error) {
    if request == nil {
        request = NewModifyCcnRegionBandwidthLimitsTypeRequest()
    }
    response = NewModifyCcnRegionBandwidthLimitsTypeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCustomerGatewayAttributeRequest() (request *ModifyCustomerGatewayAttributeRequest) {
    request = &ModifyCustomerGatewayAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyCustomerGatewayAttribute")
    return
}

func NewModifyCustomerGatewayAttributeResponse() (response *ModifyCustomerGatewayAttributeResponse) {
    response = &ModifyCustomerGatewayAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCustomerGatewayAttribute
// This API (ModifyCustomerGatewayAttribute) is used to modify the customer gateway information.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyCustomerGatewayAttribute(request *ModifyCustomerGatewayAttributeRequest) (response *ModifyCustomerGatewayAttributeResponse, err error) {
    if request == nil {
        request = NewModifyCustomerGatewayAttributeRequest()
    }
    response = NewModifyCustomerGatewayAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDirectConnectGatewayAttributeRequest() (request *ModifyDirectConnectGatewayAttributeRequest) {
    request = &ModifyDirectConnectGatewayAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyDirectConnectGatewayAttribute")
    return
}

func NewModifyDirectConnectGatewayAttributeResponse() (response *ModifyDirectConnectGatewayAttributeResponse) {
    response = &ModifyDirectConnectGatewayAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDirectConnectGatewayAttribute
// This API is used to modify the attributes of a direct connect gateway.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_DIRECTCONNECTGATEWAYISUPDATINGCOMMUNITY = "UnsupportedOperation.DirectConnectGatewayIsUpdatingCommunity"
func (c *Client) ModifyDirectConnectGatewayAttribute(request *ModifyDirectConnectGatewayAttributeRequest) (response *ModifyDirectConnectGatewayAttributeResponse, err error) {
    if request == nil {
        request = NewModifyDirectConnectGatewayAttributeRequest()
    }
    response = NewModifyDirectConnectGatewayAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyFlowLogAttributeRequest() (request *ModifyFlowLogAttributeRequest) {
    request = &ModifyFlowLogAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyFlowLogAttribute")
    return
}

func NewModifyFlowLogAttributeResponse() (response *ModifyFlowLogAttributeResponse) {
    response = &ModifyFlowLogAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyFlowLogAttribute
// This API is used to modify the attributes of a flow log.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
func (c *Client) ModifyFlowLogAttribute(request *ModifyFlowLogAttributeRequest) (response *ModifyFlowLogAttributeResponse, err error) {
    if request == nil {
        request = NewModifyFlowLogAttributeRequest()
    }
    response = NewModifyFlowLogAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyGatewayFlowQosRequest() (request *ModifyGatewayFlowQosRequest) {
    request = &ModifyGatewayFlowQosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyGatewayFlowQos")
    return
}

func NewModifyGatewayFlowQosResponse() (response *ModifyGatewayFlowQosResponse) {
    response = &ModifyGatewayFlowQosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyGatewayFlowQos
// This API (ModifyGatewayFlowQos) is used to adjust the QoS bandwidth limit in a gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) ModifyGatewayFlowQos(request *ModifyGatewayFlowQosRequest) (response *ModifyGatewayFlowQosResponse, err error) {
    if request == nil {
        request = NewModifyGatewayFlowQosRequest()
    }
    response = NewModifyGatewayFlowQosResponse()
    err = c.Send(request, response)
    return
}

func NewModifyHaVipAttributeRequest() (request *ModifyHaVipAttributeRequest) {
    request = &ModifyHaVipAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyHaVipAttribute")
    return
}

func NewModifyHaVipAttributeResponse() (response *ModifyHaVipAttributeResponse) {
    response = &ModifyHaVipAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyHaVipAttribute
// This API (ModifyHaVipAttribute) is used to modify HAVIP attributes.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyHaVipAttribute(request *ModifyHaVipAttributeRequest) (response *ModifyHaVipAttributeResponse, err error) {
    if request == nil {
        request = NewModifyHaVipAttributeRequest()
    }
    response = NewModifyHaVipAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyIpv6AddressesAttributeRequest() (request *ModifyIpv6AddressesAttributeRequest) {
    request = &ModifyIpv6AddressesAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyIpv6AddressesAttribute")
    return
}

func NewModifyIpv6AddressesAttributeResponse() (response *ModifyIpv6AddressesAttributeResponse) {
    response = &ModifyIpv6AddressesAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyIpv6AddressesAttribute
// This API (ModifyIpv6AddressesAttribute) is used to modify the private IPv6 address attributes of an ENI.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ATTACHMENTNOTFOUND = "UnsupportedOperation.AttachmentNotFound"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) ModifyIpv6AddressesAttribute(request *ModifyIpv6AddressesAttributeRequest) (response *ModifyIpv6AddressesAttributeResponse, err error) {
    if request == nil {
        request = NewModifyIpv6AddressesAttributeRequest()
    }
    response = NewModifyIpv6AddressesAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLocalGatewayRequest() (request *ModifyLocalGatewayRequest) {
    request = &ModifyLocalGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyLocalGateway")
    return
}

func NewModifyLocalGatewayResponse() (response *ModifyLocalGatewayResponse) {
    response = &ModifyLocalGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLocalGateway
// This API is used to modify the local gateway of a CDC instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyLocalGateway(request *ModifyLocalGatewayRequest) (response *ModifyLocalGatewayResponse, err error) {
    if request == nil {
        request = NewModifyLocalGatewayRequest()
    }
    response = NewModifyLocalGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNatGatewayAttributeRequest() (request *ModifyNatGatewayAttributeRequest) {
    request = &ModifyNatGatewayAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyNatGatewayAttribute")
    return
}

func NewModifyNatGatewayAttributeResponse() (response *ModifyNatGatewayAttributeResponse) {
    response = &ModifyNatGatewayAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyNatGatewayAttribute
// This API (ModifyNatGatewayAttribute) is used to modify the attributes of a NAT gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyNatGatewayAttribute(request *ModifyNatGatewayAttributeRequest) (response *ModifyNatGatewayAttributeResponse, err error) {
    if request == nil {
        request = NewModifyNatGatewayAttributeRequest()
    }
    response = NewModifyNatGatewayAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNatGatewayDestinationIpPortTranslationNatRuleRequest() (request *ModifyNatGatewayDestinationIpPortTranslationNatRuleRequest) {
    request = &ModifyNatGatewayDestinationIpPortTranslationNatRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyNatGatewayDestinationIpPortTranslationNatRule")
    return
}

func NewModifyNatGatewayDestinationIpPortTranslationNatRuleResponse() (response *ModifyNatGatewayDestinationIpPortTranslationNatRuleResponse) {
    response = &ModifyNatGatewayDestinationIpPortTranslationNatRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyNatGatewayDestinationIpPortTranslationNatRule
// This API (ModifyNatGatewayDestinationIpPortTranslationNatRule) is used to modify a port forwarding rule for a NAT gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyNatGatewayDestinationIpPortTranslationNatRule(request *ModifyNatGatewayDestinationIpPortTranslationNatRuleRequest) (response *ModifyNatGatewayDestinationIpPortTranslationNatRuleResponse, err error) {
    if request == nil {
        request = NewModifyNatGatewayDestinationIpPortTranslationNatRuleRequest()
    }
    response = NewModifyNatGatewayDestinationIpPortTranslationNatRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNatGatewaySourceIpTranslationNatRuleRequest() (request *ModifyNatGatewaySourceIpTranslationNatRuleRequest) {
    request = &ModifyNatGatewaySourceIpTranslationNatRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyNatGatewaySourceIpTranslationNatRule")
    return
}

func NewModifyNatGatewaySourceIpTranslationNatRuleResponse() (response *ModifyNatGatewaySourceIpTranslationNatRuleResponse) {
    response = &ModifyNatGatewaySourceIpTranslationNatRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyNatGatewaySourceIpTranslationNatRule
// This API is used to modify a SNAT forwarding rule of the NAT Gateway.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyNatGatewaySourceIpTranslationNatRule(request *ModifyNatGatewaySourceIpTranslationNatRuleRequest) (response *ModifyNatGatewaySourceIpTranslationNatRuleResponse, err error) {
    if request == nil {
        request = NewModifyNatGatewaySourceIpTranslationNatRuleRequest()
    }
    response = NewModifyNatGatewaySourceIpTranslationNatRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNetDetectRequest() (request *ModifyNetDetectRequest) {
    request = &ModifyNetDetectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyNetDetect")
    return
}

func NewModifyNetDetectResponse() (response *ModifyNetDetectResponse) {
    response = &ModifyNetDetectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyNetDetect
// This API (ModifyNetDetect) is used to modify network detection parameters.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_NEXTHOPMISMATCH = "InvalidParameter.NextHopMismatch"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NETDETECTSAMEIP = "InvalidParameterValue.NetDetectSameIp"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_ECMPWITHUSERROUTE = "UnsupportedOperation.EcmpWithUserRoute"
func (c *Client) ModifyNetDetect(request *ModifyNetDetectRequest) (response *ModifyNetDetectResponse, err error) {
    if request == nil {
        request = NewModifyNetDetectRequest()
    }
    response = NewModifyNetDetectResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNetworkAclAttributeRequest() (request *ModifyNetworkAclAttributeRequest) {
    request = &ModifyNetworkAclAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyNetworkAclAttribute")
    return
}

func NewModifyNetworkAclAttributeResponse() (response *ModifyNetworkAclAttributeResponse) {
    response = &ModifyNetworkAclAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyNetworkAclAttribute
// This API is used to modify the attributes of a network ACL.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyNetworkAclAttribute(request *ModifyNetworkAclAttributeRequest) (response *ModifyNetworkAclAttributeResponse, err error) {
    if request == nil {
        request = NewModifyNetworkAclAttributeRequest()
    }
    response = NewModifyNetworkAclAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNetworkAclEntriesRequest() (request *ModifyNetworkAclEntriesRequest) {
    request = &ModifyNetworkAclEntriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyNetworkAclEntries")
    return
}

func NewModifyNetworkAclEntriesResponse() (response *ModifyNetworkAclEntriesResponse) {
    response = &ModifyNetworkAclEntriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyNetworkAclEntries
// This API is used to modify (add or delete) the inbound and outbound rules of a network ACL.
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyNetworkAclEntries(request *ModifyNetworkAclEntriesRequest) (response *ModifyNetworkAclEntriesResponse, err error) {
    if request == nil {
        request = NewModifyNetworkAclEntriesRequest()
    }
    response = NewModifyNetworkAclEntriesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNetworkInterfaceAttributeRequest() (request *ModifyNetworkInterfaceAttributeRequest) {
    request = &ModifyNetworkInterfaceAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyNetworkInterfaceAttribute")
    return
}

func NewModifyNetworkInterfaceAttributeResponse() (response *ModifyNetworkInterfaceAttributeResponse) {
    response = &ModifyNetworkInterfaceAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyNetworkInterfaceAttribute
// This API (ModifyNetworkInterfaceAttribute) is used to modify ENI attributes.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyNetworkInterfaceAttribute(request *ModifyNetworkInterfaceAttributeRequest) (response *ModifyNetworkInterfaceAttributeResponse, err error) {
    if request == nil {
        request = NewModifyNetworkInterfaceAttributeRequest()
    }
    response = NewModifyNetworkInterfaceAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrivateIpAddressesAttributeRequest() (request *ModifyPrivateIpAddressesAttributeRequest) {
    request = &ModifyPrivateIpAddressesAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyPrivateIpAddressesAttribute")
    return
}

func NewModifyPrivateIpAddressesAttributeResponse() (response *ModifyPrivateIpAddressesAttributeResponse) {
    response = &ModifyPrivateIpAddressesAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPrivateIpAddressesAttribute
// This API (ModifyPrivateIpAddressesAttribute) is used to modify the private IP attributes of an ENI.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) ModifyPrivateIpAddressesAttribute(request *ModifyPrivateIpAddressesAttributeRequest) (response *ModifyPrivateIpAddressesAttributeResponse, err error) {
    if request == nil {
        request = NewModifyPrivateIpAddressesAttributeRequest()
    }
    response = NewModifyPrivateIpAddressesAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRouteTableAttributeRequest() (request *ModifyRouteTableAttributeRequest) {
    request = &ModifyRouteTableAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyRouteTableAttribute")
    return
}

func NewModifyRouteTableAttributeResponse() (response *ModifyRouteTableAttributeResponse) {
    response = &ModifyRouteTableAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRouteTableAttribute
// This API (ModifyRouteTableAttribute) is used to modify the attributes of a route table.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyRouteTableAttribute(request *ModifyRouteTableAttributeRequest) (response *ModifyRouteTableAttributeResponse, err error) {
    if request == nil {
        request = NewModifyRouteTableAttributeRequest()
    }
    response = NewModifyRouteTableAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifySecurityGroupAttributeRequest() (request *ModifySecurityGroupAttributeRequest) {
    request = &ModifySecurityGroupAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifySecurityGroupAttribute")
    return
}

func NewModifySecurityGroupAttributeResponse() (response *ModifySecurityGroupAttributeResponse) {
    response = &ModifySecurityGroupAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySecurityGroupAttribute
// This API (ModifySecurityGroupAttribute) is used to modify the attributes of a security group (SecurityGroupPolicy).
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifySecurityGroupAttribute(request *ModifySecurityGroupAttributeRequest) (response *ModifySecurityGroupAttributeResponse, err error) {
    if request == nil {
        request = NewModifySecurityGroupAttributeRequest()
    }
    response = NewModifySecurityGroupAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifySecurityGroupPoliciesRequest() (request *ModifySecurityGroupPoliciesRequest) {
    request = &ModifySecurityGroupPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifySecurityGroupPolicies")
    return
}

func NewModifySecurityGroupPoliciesResponse() (response *ModifySecurityGroupPoliciesResponse) {
    response = &ModifySecurityGroupPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySecurityGroupPolicies
// This API is used to reset the egress and ingress policies (SecurityGroupPolicy) of a security group.
//
// 
//
// <ul>
//
// <li>This API deletes all the existing egress and ingress policies, and then adds Egress and Ingress policies. It does not support custom indexes `PolicyIndex`.</li>
//
// <li>For parameters of SecurityGroupPolicySet,<ul>
//
// 	<li>If `SecurityGroupPolicySet.Version` is set to 0, all policies will be cleared, and `Egress` and `Ingress` will be ignored.</li>
//
// 	<li>If `SecurityGroupPolicySet.Version` is not set to 0, add `Egress` and `Ingress` policies:<ul>
//
// 		<li>`Protocol`: allows TCP, UDP, ICMP, ICMPV6, GRE, or ALL.</li>
//
// 		<li>`CidrBlock`: a CIDR block in the correct format. In a classic network, if a `CidrBlock` contains private IPs on Tencent Cloud for devices under your account other than CVMs, it does not mean this policy allows you to access these devices. The network isolation policies between tenants take priority over the private network policies in security groups.</li>
//
// 		<li>`Ipv6CidrBlock`: an IPv6 CIDR block in the correct format. In a classic network, if an `Ipv6CidrBlock` contains private IPv6 addresses on Tencent Cloud for devices under your account other than CVMs, it does not mean this policy allows you to access these devices. The network isolation policies between tenants take priority over the private network policies in security groups.</li>
//
// 		<li>`SecurityGroupId`: ID of the security group. It can be the ID of security group to be modified, or the ID of other security group in the same project. All private IPs of all CVMs under the security group will be covered. If this field is used, the policy will automatically change according to the CVM associated with the group ID while being used to match network messages. You don't need to change it manually.</li>
//
// 		<li>`Port`: a single port number such as 80, or a port range in the format of '8000-8010'. You may use this field only if the `Protocol` field takes the value `TCP` or `UDP`.</li>
//
// 		<li>`Action`: only allows ACCEPT or DROP.</li>
//
// 		<li>`CidrBlock`, `Ipv6CidrBlock`, `SecurityGroupId`, and `AddressTemplate` are mutually exclusive. `Protocol` + `Port` and `ServiceTemplate` are mutually exclusive.</li>
//
// </ul></li></ul></li>
//
// </ul>
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_DUPLICATEPOLICY = "UnsupportedOperation.DuplicatePolicy"
func (c *Client) ModifySecurityGroupPolicies(request *ModifySecurityGroupPoliciesRequest) (response *ModifySecurityGroupPoliciesResponse, err error) {
    if request == nil {
        request = NewModifySecurityGroupPoliciesRequest()
    }
    response = NewModifySecurityGroupPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyServiceTemplateAttributeRequest() (request *ModifyServiceTemplateAttributeRequest) {
    request = &ModifyServiceTemplateAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyServiceTemplateAttribute")
    return
}

func NewModifyServiceTemplateAttributeResponse() (response *ModifyServiceTemplateAttributeResponse) {
    response = &ModifyServiceTemplateAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyServiceTemplateAttribute
// This API (ModifyServiceTemplateAttribute) is used to modify a protocol port template.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) ModifyServiceTemplateAttribute(request *ModifyServiceTemplateAttributeRequest) (response *ModifyServiceTemplateAttributeResponse, err error) {
    if request == nil {
        request = NewModifyServiceTemplateAttributeRequest()
    }
    response = NewModifyServiceTemplateAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyServiceTemplateGroupAttributeRequest() (request *ModifyServiceTemplateGroupAttributeRequest) {
    request = &ModifyServiceTemplateGroupAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyServiceTemplateGroupAttribute")
    return
}

func NewModifyServiceTemplateGroupAttributeResponse() (response *ModifyServiceTemplateGroupAttributeResponse) {
    response = &ModifyServiceTemplateGroupAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyServiceTemplateGroupAttribute
// This API (ModifyServiceTemplateGroupAttribute) is used to modify a protocol port template group.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyServiceTemplateGroupAttribute(request *ModifyServiceTemplateGroupAttributeRequest) (response *ModifyServiceTemplateGroupAttributeResponse, err error) {
    if request == nil {
        request = NewModifyServiceTemplateGroupAttributeRequest()
    }
    response = NewModifyServiceTemplateGroupAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifySubnetAttributeRequest() (request *ModifySubnetAttributeRequest) {
    request = &ModifySubnetAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifySubnetAttribute")
    return
}

func NewModifySubnetAttributeResponse() (response *ModifySubnetAttributeResponse) {
    response = &ModifySubnetAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySubnetAttribute
// This API (ModifySubnetAttribute) is used to modify subnet attributes.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifySubnetAttribute(request *ModifySubnetAttributeRequest) (response *ModifySubnetAttributeResponse, err error) {
    if request == nil {
        request = NewModifySubnetAttributeRequest()
    }
    response = NewModifySubnetAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVpcAttributeRequest() (request *ModifyVpcAttributeRequest) {
    request = &ModifyVpcAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyVpcAttribute")
    return
}

func NewModifyVpcAttributeResponse() (response *ModifyVpcAttributeResponse) {
    response = &ModifyVpcAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyVpcAttribute
// This API (ModifyVpcAttribute) is used to modify VPC attributes.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyVpcAttribute(request *ModifyVpcAttributeRequest) (response *ModifyVpcAttributeResponse, err error) {
    if request == nil {
        request = NewModifyVpcAttributeRequest()
    }
    response = NewModifyVpcAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVpcEndPointAttributeRequest() (request *ModifyVpcEndPointAttributeRequest) {
    request = &ModifyVpcEndPointAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyVpcEndPointAttribute")
    return
}

func NewModifyVpcEndPointAttributeResponse() (response *ModifyVpcEndPointAttributeResponse) {
    response = &ModifyVpcEndPointAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyVpcEndPointAttribute
// This API is used to modify endpoint attributes.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_SPECIALENDPOINTSERVICE = "UnsupportedOperation.SpecialEndPointService"
func (c *Client) ModifyVpcEndPointAttribute(request *ModifyVpcEndPointAttributeRequest) (response *ModifyVpcEndPointAttributeResponse, err error) {
    if request == nil {
        request = NewModifyVpcEndPointAttributeRequest()
    }
    response = NewModifyVpcEndPointAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVpcEndPointServiceAttributeRequest() (request *ModifyVpcEndPointServiceAttributeRequest) {
    request = &ModifyVpcEndPointServiceAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyVpcEndPointServiceAttribute")
    return
}

func NewModifyVpcEndPointServiceAttributeResponse() (response *ModifyVpcEndPointServiceAttributeResponse) {
    response = &ModifyVpcEndPointServiceAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyVpcEndPointServiceAttribute
// This API is used to modify endpoint service attributes.
//
// 
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
func (c *Client) ModifyVpcEndPointServiceAttribute(request *ModifyVpcEndPointServiceAttributeRequest) (response *ModifyVpcEndPointServiceAttributeResponse, err error) {
    if request == nil {
        request = NewModifyVpcEndPointServiceAttributeRequest()
    }
    response = NewModifyVpcEndPointServiceAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVpcEndPointServiceWhiteListRequest() (request *ModifyVpcEndPointServiceWhiteListRequest) {
    request = &ModifyVpcEndPointServiceWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyVpcEndPointServiceWhiteList")
    return
}

func NewModifyVpcEndPointServiceWhiteListResponse() (response *ModifyVpcEndPointServiceWhiteListResponse) {
    response = &ModifyVpcEndPointServiceWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyVpcEndPointServiceWhiteList
// This API is used to modify the attributes of the endpoint service allowlist.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_UINNOTFOUND = "UnsupportedOperation.UinNotFound"
func (c *Client) ModifyVpcEndPointServiceWhiteList(request *ModifyVpcEndPointServiceWhiteListRequest) (response *ModifyVpcEndPointServiceWhiteListResponse, err error) {
    if request == nil {
        request = NewModifyVpcEndPointServiceWhiteListRequest()
    }
    response = NewModifyVpcEndPointServiceWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVpnConnectionAttributeRequest() (request *ModifyVpnConnectionAttributeRequest) {
    request = &ModifyVpnConnectionAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyVpnConnectionAttribute")
    return
}

func NewModifyVpnConnectionAttributeResponse() (response *ModifyVpnConnectionAttributeResponse) {
    response = &ModifyVpnConnectionAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyVpnConnectionAttribute
// This API (ModifyVpnConnectionAttribute) is used to modify VPN tunnels.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_VPNCONNCIDRCONFLICT = "InvalidParameterValue.VpnConnCidrConflict"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyVpnConnectionAttribute(request *ModifyVpnConnectionAttributeRequest) (response *ModifyVpnConnectionAttributeResponse, err error) {
    if request == nil {
        request = NewModifyVpnConnectionAttributeRequest()
    }
    response = NewModifyVpnConnectionAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVpnGatewayAttributeRequest() (request *ModifyVpnGatewayAttributeRequest) {
    request = &ModifyVpnGatewayAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyVpnGatewayAttribute")
    return
}

func NewModifyVpnGatewayAttributeResponse() (response *ModifyVpnGatewayAttributeResponse) {
    response = &ModifyVpnGatewayAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyVpnGatewayAttribute
// This API (ModifyVpnGatewayAttribute) is used to modify the attributes of VPN gateways.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyVpnGatewayAttribute(request *ModifyVpnGatewayAttributeRequest) (response *ModifyVpnGatewayAttributeResponse, err error) {
    if request == nil {
        request = NewModifyVpnGatewayAttributeRequest()
    }
    response = NewModifyVpnGatewayAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVpnGatewayCcnRoutesRequest() (request *ModifyVpnGatewayCcnRoutesRequest) {
    request = &ModifyVpnGatewayCcnRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyVpnGatewayCcnRoutes")
    return
}

func NewModifyVpnGatewayCcnRoutesResponse() (response *ModifyVpnGatewayCcnRoutesResponse) {
    response = &ModifyVpnGatewayCcnRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyVpnGatewayCcnRoutes
// This API (ModifyVpnGatewayCcnRoutes) is used to modify VPN gateway-based CCN routes.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyVpnGatewayCcnRoutes(request *ModifyVpnGatewayCcnRoutesRequest) (response *ModifyVpnGatewayCcnRoutesResponse, err error) {
    if request == nil {
        request = NewModifyVpnGatewayCcnRoutesRequest()
    }
    response = NewModifyVpnGatewayCcnRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewNotifyRoutesRequest() (request *NotifyRoutesRequest) {
    request = &NotifyRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "NotifyRoutes")
    return
}

func NewNotifyRoutesResponse() (response *NotifyRoutesResponse) {
    response = &NotifyRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// NotifyRoutes
// This API is used to publish a route to CCN. This can also be done by clicking the **Publish to CCN** button on the route table page.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDROUTEID_NOTFOUND = "InvalidRouteId.NotFound"
//  INVALIDROUTETABLEID_MALFORMED = "InvalidRouteTableId.Malformed"
//  INVALIDROUTETABLEID_NOTFOUND = "InvalidRouteTableId.NotFound"
//  UNSUPPORTEDOPERATION_CCNNOTATTACHED = "UnsupportedOperation.CcnNotAttached"
//  UNSUPPORTEDOPERATION_NOTIFYCCN = "UnsupportedOperation.NotifyCcn"
//  UNSUPPORTEDOPERATION_SYSTEMROUTE = "UnsupportedOperation.SystemRoute"
func (c *Client) NotifyRoutes(request *NotifyRoutesRequest) (response *NotifyRoutesResponse, err error) {
    if request == nil {
        request = NewNotifyRoutesRequest()
    }
    response = NewNotifyRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewRejectAttachCcnInstancesRequest() (request *RejectAttachCcnInstancesRequest) {
    request = &RejectAttachCcnInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "RejectAttachCcnInstances")
    return
}

func NewRejectAttachCcnInstancesResponse() (response *RejectAttachCcnInstancesResponse) {
    response = &RejectAttachCcnInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RejectAttachCcnInstances
// This API (RejectAttachCcnInstances) is used to reject association operations when instances are associated across accounts for the CCN owner.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CCNNOTATTACHED = "UnsupportedOperation.CcnNotAttached"
//  UNSUPPORTEDOPERATION_NOTPENDINGCCNINSTANCE = "UnsupportedOperation.NotPendingCcnInstance"
func (c *Client) RejectAttachCcnInstances(request *RejectAttachCcnInstancesRequest) (response *RejectAttachCcnInstancesResponse, err error) {
    if request == nil {
        request = NewRejectAttachCcnInstancesRequest()
    }
    response = NewRejectAttachCcnInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseAddressesRequest() (request *ReleaseAddressesRequest) {
    request = &ReleaseAddressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ReleaseAddresses")
    return
}

func NewReleaseAddressesResponse() (response *ReleaseAddressesResponse) {
    response = &ReleaseAddressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReleaseAddresses
// This API (ReleaseAddresses) is used to release one or multiple [Elastic IPs](https://intl.cloud.tencent.com/document/product/213/1941?from_cn_redirect=1).
//
// * This operation is irreversible. Once you release an EIP, the IP address associated with the EIP no longer belongs to you.
//
// * Only EIPs in UNBIND status can be released.
//
// error code that may be returned:
//  INVALIDADDRESSID_BLOCKED = "InvalidAddressId.Blocked"
//  INVALIDADDRESSID_NOTFOUND = "InvalidAddressId.NotFound"
//  INVALIDADDRESSSTATE = "InvalidAddressState"
func (c *Client) ReleaseAddresses(request *ReleaseAddressesRequest) (response *ReleaseAddressesResponse, err error) {
    if request == nil {
        request = NewReleaseAddressesRequest()
    }
    response = NewReleaseAddressesResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveBandwidthPackageResourcesRequest() (request *RemoveBandwidthPackageResourcesRequest) {
    request = &RemoveBandwidthPackageResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "RemoveBandwidthPackageResources")
    return
}

func NewRemoveBandwidthPackageResourcesResponse() (response *RemoveBandwidthPackageResourcesResponse) {
    response = &RemoveBandwidthPackageResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RemoveBandwidthPackageResources
// This API is used to delete a bandwidth package resource, including [Elastic IP](https://intl.cloud.tencent.com/document/product/213/1941?from_cn_redirect=1), [Cloud Load Balancer](https://intl.cloud.tencent.com/document/product/214/517?from_cn_redirect=1), and so on.
//
// error code that may be returned:
//  INVALIDADDRESSID_BLOCKED = "InvalidAddressId.Blocked"
//  INVALIDADDRESSID_NOTFOUND = "InvalidAddressId.NotFound"
//  INVALIDADDRESSSTATE = "InvalidAddressState"
func (c *Client) RemoveBandwidthPackageResources(request *RemoveBandwidthPackageResourcesRequest) (response *RemoveBandwidthPackageResourcesResponse, err error) {
    if request == nil {
        request = NewRemoveBandwidthPackageResourcesRequest()
    }
    response = NewRemoveBandwidthPackageResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewRenewVpnGatewayRequest() (request *RenewVpnGatewayRequest) {
    request = &RenewVpnGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "RenewVpnGateway")
    return
}

func NewRenewVpnGatewayResponse() (response *RenewVpnGatewayResponse) {
    response = &RenewVpnGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RenewVpnGateway
// This API (RenewVpnGateway) is used to renew prepaid (monthly subscription) VPN gateways. Currently, only IPSEC gateways are supported.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RenewVpnGateway(request *RenewVpnGatewayRequest) (response *RenewVpnGatewayResponse, err error) {
    if request == nil {
        request = NewRenewVpnGatewayRequest()
    }
    response = NewRenewVpnGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewReplaceDirectConnectGatewayCcnRoutesRequest() (request *ReplaceDirectConnectGatewayCcnRoutesRequest) {
    request = &ReplaceDirectConnectGatewayCcnRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ReplaceDirectConnectGatewayCcnRoutes")
    return
}

func NewReplaceDirectConnectGatewayCcnRoutesResponse() (response *ReplaceDirectConnectGatewayCcnRoutesResponse) {
    response = &ReplaceDirectConnectGatewayCcnRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReplaceDirectConnectGatewayCcnRoutes
// This API (ReplaceDirectConnectGatewayCcnRoutes) is used to modify the specified route according to the route ID. Batch modification is supported.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ReplaceDirectConnectGatewayCcnRoutes(request *ReplaceDirectConnectGatewayCcnRoutesRequest) (response *ReplaceDirectConnectGatewayCcnRoutesResponse, err error) {
    if request == nil {
        request = NewReplaceDirectConnectGatewayCcnRoutesRequest()
    }
    response = NewReplaceDirectConnectGatewayCcnRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewReplaceRouteTableAssociationRequest() (request *ReplaceRouteTableAssociationRequest) {
    request = &ReplaceRouteTableAssociationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ReplaceRouteTableAssociation")
    return
}

func NewReplaceRouteTableAssociationResponse() (response *ReplaceRouteTableAssociationResponse) {
    response = &ReplaceRouteTableAssociationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReplaceRouteTableAssociation
// This API (ReplaceRouteTableAssociation) is used to modify the route table associated with a subnet.
//
// * A subnet can only be associated with one route table.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
func (c *Client) ReplaceRouteTableAssociation(request *ReplaceRouteTableAssociationRequest) (response *ReplaceRouteTableAssociationResponse, err error) {
    if request == nil {
        request = NewReplaceRouteTableAssociationRequest()
    }
    response = NewReplaceRouteTableAssociationResponse()
    err = c.Send(request, response)
    return
}

func NewReplaceRoutesRequest() (request *ReplaceRoutesRequest) {
    request = &ReplaceRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ReplaceRoutes")
    return
}

func NewReplaceRoutesResponse() (response *ReplaceRoutesResponse) {
    response = &ReplaceRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReplaceRoutes
// This API (ReplaceRoutes) is used to modify a specified routing policy by its ID (RouteId). Batch modification is supported.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CIDRNOTINPEERVPC = "InvalidParameterValue.CidrNotInPeerVpc"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_VPCCIDRCONFLICT = "InvalidParameterValue.VpcCidrConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER_WITHGUESS = "UnknownParameter.WithGuess"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CDCSUBNETNOTSUPPORTUNLOCALGATEWAY = "UnsupportedOperation.CdcSubnetNotSupportUnLocalGateway"
//  UNSUPPORTEDOPERATION_CONFLICTWITHDOCKERROUTE = "UnsupportedOperation.ConflictWithDockerRoute"
//  UNSUPPORTEDOPERATION_ECMP = "UnsupportedOperation.Ecmp"
//  UNSUPPORTEDOPERATION_NORMALSUBNETNOTSUPPORTLOCALGATEWAY = "UnsupportedOperation.NormalSubnetNotSupportLocalGateway"
//  UNSUPPORTEDOPERATION_SYSTEMROUTE = "UnsupportedOperation.SystemRoute"
func (c *Client) ReplaceRoutes(request *ReplaceRoutesRequest) (response *ReplaceRoutesResponse, err error) {
    if request == nil {
        request = NewReplaceRoutesRequest()
    }
    response = NewReplaceRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewReplaceSecurityGroupPolicyRequest() (request *ReplaceSecurityGroupPolicyRequest) {
    request = &ReplaceSecurityGroupPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ReplaceSecurityGroupPolicy")
    return
}

func NewReplaceSecurityGroupPolicyResponse() (response *ReplaceSecurityGroupPolicyResponse) {
    response = &ReplaceSecurityGroupPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReplaceSecurityGroupPolicy
// This API (ReplaceSecurityGroupPolicy) is used to replace a single security group policy (SecurityGroupPolicy).
//
// Only one policy in a single direction can be replaced in each request, and the PolicyIndex parameter must be specified.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_DUPLICATEPOLICY = "UnsupportedOperation.DuplicatePolicy"
//  UNSUPPORTEDOPERATION_VERSIONMISMATCH = "UnsupportedOperation.VersionMismatch"
func (c *Client) ReplaceSecurityGroupPolicy(request *ReplaceSecurityGroupPolicyRequest) (response *ReplaceSecurityGroupPolicyResponse, err error) {
    if request == nil {
        request = NewReplaceSecurityGroupPolicyRequest()
    }
    response = NewReplaceSecurityGroupPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewResetAttachCcnInstancesRequest() (request *ResetAttachCcnInstancesRequest) {
    request = &ResetAttachCcnInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ResetAttachCcnInstances")
    return
}

func NewResetAttachCcnInstancesResponse() (response *ResetAttachCcnInstancesResponse) {
    response = &ResetAttachCcnInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetAttachCcnInstances
// This API (ResetAttachCcnInstances) is used to re-apply for the association operation when the application for cross-account instance association expires.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ResetAttachCcnInstances(request *ResetAttachCcnInstancesRequest) (response *ResetAttachCcnInstancesResponse, err error) {
    if request == nil {
        request = NewResetAttachCcnInstancesRequest()
    }
    response = NewResetAttachCcnInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewResetNatGatewayConnectionRequest() (request *ResetNatGatewayConnectionRequest) {
    request = &ResetNatGatewayConnectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ResetNatGatewayConnection")
    return
}

func NewResetNatGatewayConnectionResponse() (response *ResetNatGatewayConnectionResponse) {
    response = &ResetNatGatewayConnectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetNatGatewayConnection
// This API (ResetNatGatewayConnection) is used to adjust concurrent connection cap for the NAT gateway.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) ResetNatGatewayConnection(request *ResetNatGatewayConnectionRequest) (response *ResetNatGatewayConnectionResponse, err error) {
    if request == nil {
        request = NewResetNatGatewayConnectionRequest()
    }
    response = NewResetNatGatewayConnectionResponse()
    err = c.Send(request, response)
    return
}

func NewResetRoutesRequest() (request *ResetRoutesRequest) {
    request = &ResetRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ResetRoutes")
    return
}

func NewResetRoutesResponse() (response *ResetRoutesResponse) {
    response = &ResetRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetRoutes
// This API (ResetRoutes) is used to reset the name of a route table and all its routing policies.<br />
//
// Note: When this API is called, all routing policies in the current route table are deleted before new routing policies are saved, which may incur network interruption.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CIDRNOTINPEERVPC = "InvalidParameterValue.CidrNotInPeerVpc"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_VPCCIDRCONFLICT = "InvalidParameterValue.VpcCidrConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ECMP = "UnsupportedOperation.Ecmp"
//  UNSUPPORTEDOPERATION_SYSTEMROUTE = "UnsupportedOperation.SystemRoute"
func (c *Client) ResetRoutes(request *ResetRoutesRequest) (response *ResetRoutesResponse, err error) {
    if request == nil {
        request = NewResetRoutesRequest()
    }
    response = NewResetRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewResetVpnConnectionRequest() (request *ResetVpnConnectionRequest) {
    request = &ResetVpnConnectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ResetVpnConnection")
    return
}

func NewResetVpnConnectionResponse() (response *ResetVpnConnectionResponse) {
    response = &ResetVpnConnectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetVpnConnection
// The API (ResetVpnConnection) is used to reset VPN tunnels.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ResetVpnConnection(request *ResetVpnConnectionRequest) (response *ResetVpnConnectionResponse, err error) {
    if request == nil {
        request = NewResetVpnConnectionRequest()
    }
    response = NewResetVpnConnectionResponse()
    err = c.Send(request, response)
    return
}

func NewResetVpnGatewayInternetMaxBandwidthRequest() (request *ResetVpnGatewayInternetMaxBandwidthRequest) {
    request = &ResetVpnGatewayInternetMaxBandwidthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ResetVpnGatewayInternetMaxBandwidth")
    return
}

func NewResetVpnGatewayInternetMaxBandwidthResponse() (response *ResetVpnGatewayInternetMaxBandwidthResponse) {
    response = &ResetVpnGatewayInternetMaxBandwidthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetVpnGatewayInternetMaxBandwidth
// This API (ResetVpnGatewayInternetMaxBandwidth) is used to adjust the bandwidth cap of VPN gateways. Currently, only configuration upgrade is supported. VPN gateways with monthly subscription must be within the validity period.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) ResetVpnGatewayInternetMaxBandwidth(request *ResetVpnGatewayInternetMaxBandwidthRequest) (response *ResetVpnGatewayInternetMaxBandwidthResponse, err error) {
    if request == nil {
        request = NewResetVpnGatewayInternetMaxBandwidthRequest()
    }
    response = NewResetVpnGatewayInternetMaxBandwidthResponse()
    err = c.Send(request, response)
    return
}

func NewSetCcnRegionBandwidthLimitsRequest() (request *SetCcnRegionBandwidthLimitsRequest) {
    request = &SetCcnRegionBandwidthLimitsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "SetCcnRegionBandwidthLimits")
    return
}

func NewSetCcnRegionBandwidthLimitsResponse() (response *SetCcnRegionBandwidthLimitsResponse) {
    response = &SetCcnRegionBandwidthLimitsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetCcnRegionBandwidthLimits
// This API (SetCcnRegionBandwidthLimits) is used to set the outbound bandwidth cap for CCNs in each region. This API can only set the outbound bandwidth cap for regions in the network instances that have already been associated.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_NOTPOSTPAIDCCNOPERATION = "UnsupportedOperation.NotPostpaidCcnOperation"
func (c *Client) SetCcnRegionBandwidthLimits(request *SetCcnRegionBandwidthLimitsRequest) (response *SetCcnRegionBandwidthLimitsResponse, err error) {
    if request == nil {
        request = NewSetCcnRegionBandwidthLimitsRequest()
    }
    response = NewSetCcnRegionBandwidthLimitsResponse()
    err = c.Send(request, response)
    return
}

func NewTransformAddressRequest() (request *TransformAddressRequest) {
    request = &TransformAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "TransformAddress")
    return
}

func NewTransformAddressResponse() (response *TransformAddressResponse) {
    response = &TransformAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TransformAddress
// This API (TransformAddress) is used to switch common public IPs into [Elastic IPs](https://intl.cloud.tencent.comhttps://intl.cloud.tencent.com/document/product/213/1941?from_cn_redirect=1?from_cn_redirect=1).
//
// * The platform limits the number of times that a user can unbind an EIP and reassign public IPs in each region per day. For more information, see [EIP product introduction](https://intl.cloud.tencent.com/document/product/213/1941?from_cn_redirect=1)). The preceding quota can be obtained through the [DescribeAddressQuota](https://intl.cloud.tencent.com/document/api/213/1378?from_cn_redirect=1) API.
//
// error code that may be returned:
//  ADDRESSQUOTALIMITEXCEEDED = "AddressQuotaLimitExceeded"
//  ADDRESSQUOTALIMITEXCEEDED_DAILYALLOCATE = "AddressQuotaLimitExceeded.DailyAllocate"
//  INVALIDADDRESSID_BLOCKED = "InvalidAddressId.Blocked"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_ALREADYBINDEIP = "InvalidInstanceId.AlreadyBindEip"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
func (c *Client) TransformAddress(request *TransformAddressRequest) (response *TransformAddressResponse, err error) {
    if request == nil {
        request = NewTransformAddressRequest()
    }
    response = NewTransformAddressResponse()
    err = c.Send(request, response)
    return
}

func NewUnassignIpv6AddressesRequest() (request *UnassignIpv6AddressesRequest) {
    request = &UnassignIpv6AddressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "UnassignIpv6Addresses")
    return
}

func NewUnassignIpv6AddressesResponse() (response *UnassignIpv6AddressesResponse) {
    response = &UnassignIpv6AddressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnassignIpv6Addresses
// This API (UnassignIpv6Addresses) is used to release ENI `IPv6` addresses.<br />
//
// This API is completed asynchronously. If you need to query the async execution results, use the `RequestId` returned by this API to query the `QueryTask` API.
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_ATTACHMENTNOTFOUND = "UnauthorizedOperation.AttachmentNotFound"
//  UNAUTHORIZEDOPERATION_PRIMARYIP = "UnauthorizedOperation.PrimaryIp"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ATTACHMENTNOTFOUND = "UnsupportedOperation.AttachmentNotFound"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) UnassignIpv6Addresses(request *UnassignIpv6AddressesRequest) (response *UnassignIpv6AddressesResponse, err error) {
    if request == nil {
        request = NewUnassignIpv6AddressesRequest()
    }
    response = NewUnassignIpv6AddressesResponse()
    err = c.Send(request, response)
    return
}

func NewUnassignIpv6CidrBlockRequest() (request *UnassignIpv6CidrBlockRequest) {
    request = &UnassignIpv6CidrBlockRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "UnassignIpv6CidrBlock")
    return
}

func NewUnassignIpv6CidrBlockResponse() (response *UnassignIpv6CidrBlockResponse) {
    response = &UnassignIpv6CidrBlockResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnassignIpv6CidrBlock
// This API (UnassignIpv6CidrBlock) is used to release IPv6 IP ranges.
//
// If the IP range still has occupied IPs that are not yet repossessed, the IP range cannot be released.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UnassignIpv6CidrBlock(request *UnassignIpv6CidrBlockRequest) (response *UnassignIpv6CidrBlockResponse, err error) {
    if request == nil {
        request = NewUnassignIpv6CidrBlockRequest()
    }
    response = NewUnassignIpv6CidrBlockResponse()
    err = c.Send(request, response)
    return
}

func NewUnassignIpv6SubnetCidrBlockRequest() (request *UnassignIpv6SubnetCidrBlockRequest) {
    request = &UnassignIpv6SubnetCidrBlockRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "UnassignIpv6SubnetCidrBlock")
    return
}

func NewUnassignIpv6SubnetCidrBlockResponse() (response *UnassignIpv6SubnetCidrBlockResponse) {
    response = &UnassignIpv6SubnetCidrBlockResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnassignIpv6SubnetCidrBlock
// This API (UnassignIpv6SubnetCidrBlock) is used to release IPv6 subnet IP ranges.
//
// If the subnet IP range still has occupied IPs that are not yet repossessed, the subnet IP range cannot be released.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UnassignIpv6SubnetCidrBlock(request *UnassignIpv6SubnetCidrBlockRequest) (response *UnassignIpv6SubnetCidrBlockResponse, err error) {
    if request == nil {
        request = NewUnassignIpv6SubnetCidrBlockRequest()
    }
    response = NewUnassignIpv6SubnetCidrBlockResponse()
    err = c.Send(request, response)
    return
}

func NewUnassignPrivateIpAddressesRequest() (request *UnassignPrivateIpAddressesRequest) {
    request = &UnassignPrivateIpAddressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "UnassignPrivateIpAddresses")
    return
}

func NewUnassignPrivateIpAddressesResponse() (response *UnassignPrivateIpAddressesResponse) {
    response = &UnassignPrivateIpAddressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnassignPrivateIpAddresses
// This API (UnassignPrivateIpAddresses) is used to return the private IPs of ENI.
//
// * To return the secondary private IPs of an ENI, the API will automatically unbind the IPs of an ENI. The primary private IP of the ENI cannot be returned.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ATTACHMENTNOTFOUND = "UnsupportedOperation.AttachmentNotFound"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) UnassignPrivateIpAddresses(request *UnassignPrivateIpAddressesRequest) (response *UnassignPrivateIpAddressesResponse, err error) {
    if request == nil {
        request = NewUnassignPrivateIpAddressesRequest()
    }
    response = NewUnassignPrivateIpAddressesResponse()
    err = c.Send(request, response)
    return
}

func NewWithdrawNotifyRoutesRequest() (request *WithdrawNotifyRoutesRequest) {
    request = &WithdrawNotifyRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "WithdrawNotifyRoutes")
    return
}

func NewWithdrawNotifyRoutesResponse() (response *WithdrawNotifyRoutesResponse) {
    response = &WithdrawNotifyRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// WithdrawNotifyRoutes
// This API is used to withdraw a route from CCN. This can also be done by clicking the **Withdraw from CCN** button on the route table page.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDROUTEID_NOTFOUND = "InvalidRouteId.NotFound"
//  INVALIDROUTETABLEID_MALFORMED = "InvalidRouteTableId.Malformed"
//  INVALIDROUTETABLEID_NOTFOUND = "InvalidRouteTableId.NotFound"
//  UNSUPPORTEDOPERATION_NOTIFYCCN = "UnsupportedOperation.NotifyCcn"
//  UNSUPPORTEDOPERATION_SYSTEMROUTE = "UnsupportedOperation.SystemRoute"
func (c *Client) WithdrawNotifyRoutes(request *WithdrawNotifyRoutesRequest) (response *WithdrawNotifyRoutesResponse, err error) {
    if request == nil {
        request = NewWithdrawNotifyRoutesRequest()
    }
    response = NewWithdrawNotifyRoutesResponse()
    err = c.Send(request, response)
    return
}
