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
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
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

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
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
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_ISNOTFINANCEACCOUNT = "UnsupportedOperation.IsNotFinanceAccount"
//  UNSUPPORTEDOPERATION_NOTPENDINGCCNINSTANCE = "UnsupportedOperation.NotPendingCcnInstance"
//  UNSUPPORTEDOPERATION_UNABLECROSSFINANCE = "UnsupportedOperation.UnableCrossFinance"
func (c *Client) AcceptAttachCcnInstances(request *AcceptAttachCcnInstancesRequest) (response *AcceptAttachCcnInstancesResponse, err error) {
    return c.AcceptAttachCcnInstancesWithContext(context.Background(), request)
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
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_ISNOTFINANCEACCOUNT = "UnsupportedOperation.IsNotFinanceAccount"
//  UNSUPPORTEDOPERATION_NOTPENDINGCCNINSTANCE = "UnsupportedOperation.NotPendingCcnInstance"
//  UNSUPPORTEDOPERATION_UNABLECROSSFINANCE = "UnsupportedOperation.UnableCrossFinance"
func (c *Client) AcceptAttachCcnInstancesWithContext(ctx context.Context, request *AcceptAttachCcnInstancesRequest) (response *AcceptAttachCcnInstancesResponse, err error) {
    if request == nil {
        request = NewAcceptAttachCcnInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AcceptAttachCcnInstances require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGEIDMALFORMED = "InvalidParameterValue.BandwidthPackageIdMalformed"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGENOTFOUND = "InvalidParameterValue.BandwidthPackageNotFound"
//  INVALIDPARAMETERVALUE_COMBINATION = "InvalidParameterValue.Combination"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXISTED = "InvalidParameterValue.ResourceAlreadyExisted"
//  INVALIDPARAMETERVALUE_RESOURCEIDMALFORMED = "InvalidParameterValue.ResourceIdMalformed"
//  INVALIDPARAMETERVALUE_RESOURCENOTFOUND = "InvalidParameterValue.ResourceNotFound"
//  LIMITEXCEEDED_BANDWIDTHPACKAGEQUOTA = "LimitExceeded.BandwidthPackageQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_BANDWIDTHPACKAGEIDNOTSUPPORTED = "UnsupportedOperation.BandwidthPackageIdNotSupported"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTSUPPORTED = "UnsupportedOperation.InstanceStateNotSupported"
//  UNSUPPORTEDOPERATION_INVALIDRESOURCEINTERNETCHARGETYPE = "UnsupportedOperation.InvalidResourceInternetChargeType"
//  UNSUPPORTEDOPERATION_INVALIDRESOURCEPROTOCOL = "UnsupportedOperation.InvalidResourceProtocol"
func (c *Client) AddBandwidthPackageResources(request *AddBandwidthPackageResourcesRequest) (response *AddBandwidthPackageResourcesResponse, err error) {
    return c.AddBandwidthPackageResourcesWithContext(context.Background(), request)
}

// AddBandwidthPackageResources
// This API is used to add resources to a bandwidth package, including [Elastic IP](https://intl.cloud.tencent.com/document/product/213/1941?from_cn_redirect=1), [Cloud Load Balancer](https://intl.cloud.tencent.com/document/product/214/517?from_cn_redirect=1), and so on.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGEIDMALFORMED = "InvalidParameterValue.BandwidthPackageIdMalformed"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGENOTFOUND = "InvalidParameterValue.BandwidthPackageNotFound"
//  INVALIDPARAMETERVALUE_COMBINATION = "InvalidParameterValue.Combination"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_RESOURCEALREADYEXISTED = "InvalidParameterValue.ResourceAlreadyExisted"
//  INVALIDPARAMETERVALUE_RESOURCEIDMALFORMED = "InvalidParameterValue.ResourceIdMalformed"
//  INVALIDPARAMETERVALUE_RESOURCENOTFOUND = "InvalidParameterValue.ResourceNotFound"
//  LIMITEXCEEDED_BANDWIDTHPACKAGEQUOTA = "LimitExceeded.BandwidthPackageQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_BANDWIDTHPACKAGEIDNOTSUPPORTED = "UnsupportedOperation.BandwidthPackageIdNotSupported"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTSUPPORTED = "UnsupportedOperation.InstanceStateNotSupported"
//  UNSUPPORTEDOPERATION_INVALIDRESOURCEINTERNETCHARGETYPE = "UnsupportedOperation.InvalidResourceInternetChargeType"
//  UNSUPPORTEDOPERATION_INVALIDRESOURCEPROTOCOL = "UnsupportedOperation.InvalidResourceProtocol"
func (c *Client) AddBandwidthPackageResourcesWithContext(ctx context.Context, request *AddBandwidthPackageResourcesRequest) (response *AddBandwidthPackageResourcesResponse, err error) {
    if request == nil {
        request = NewAddBandwidthPackageResourcesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddBandwidthPackageResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddBandwidthPackageResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewAdjustPublicAddressRequest() (request *AdjustPublicAddressRequest) {
    request = &AdjustPublicAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "AdjustPublicAddress")
    
    
    return
}

func NewAdjustPublicAddressResponse() (response *AdjustPublicAddressResponse) {
    response = &AdjustPublicAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AdjustPublicAddress
// This API is used to change the IP address. It supports changing the common public IPs and EIPs billed by monthly subscribed bandwidth of a CVM instance.
//
// error code that may be returned:
//  INVALIDACCOUNT_NOTSUPPORTED = "InvalidAccount.NotSupported"
//  INVALIDADDRESSID_BLOCKED = "InvalidAddressId.Blocked"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE_ADDRESSATTACKED = "InvalidParameterValue.AddressAttacked"
//  INVALIDPARAMETERVALUE_ADDRESSIPNOTAVAILABLE = "InvalidParameterValue.AddressIpNotAvailable"
//  INVALIDPARAMETERVALUE_ADDRESSNOTFOUND = "InvalidParameterValue.AddressNotFound"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INSTANCENOWANIP = "InvalidParameterValue.InstanceNoWanIP"
//  INVALIDPARAMETERVALUE_RESOURCENOTSUPPORT = "InvalidParameterValue.ResourceNotSupport"
//  INVALIDPARAMETERVALUE_UNAVAILABLEZONE = "InvalidParameterValue.UnavailableZone"
//  LIMITEXCEEDED_CHANGEADDRESSQUOTA = "LimitExceeded.ChangeAddressQuota"
//  LIMITEXCEEDED_DAILYCHANGEADDRESSQUOTA = "LimitExceeded.DailyChangeAddressQuota"
//  OPERATIONDENIED_ADDRESSINARREARS = "OperationDenied.AddressInArrears"
//  OPERATIONDENIED_MUTEXTASKRUNNING = "OperationDenied.MutexTaskRunning"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNSUPPORTEDOPERATION_ADDRESSSTATUSNOTPERMIT = "UnsupportedOperation.AddressStatusNotPermit"
//  UNSUPPORTEDOPERATION_INVALIDADDRESSINTERNETCHARGETYPE = "UnsupportedOperation.InvalidAddressInternetChargeType"
//  UNSUPPORTEDOPERATION_ISPNOTSUPPORTED = "UnsupportedOperation.IspNotSupported"
func (c *Client) AdjustPublicAddress(request *AdjustPublicAddressRequest) (response *AdjustPublicAddressResponse, err error) {
    return c.AdjustPublicAddressWithContext(context.Background(), request)
}

// AdjustPublicAddress
// This API is used to change the IP address. It supports changing the common public IPs and EIPs billed by monthly subscribed bandwidth of a CVM instance.
//
// error code that may be returned:
//  INVALIDACCOUNT_NOTSUPPORTED = "InvalidAccount.NotSupported"
//  INVALIDADDRESSID_BLOCKED = "InvalidAddressId.Blocked"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE_ADDRESSATTACKED = "InvalidParameterValue.AddressAttacked"
//  INVALIDPARAMETERVALUE_ADDRESSIPNOTAVAILABLE = "InvalidParameterValue.AddressIpNotAvailable"
//  INVALIDPARAMETERVALUE_ADDRESSNOTFOUND = "InvalidParameterValue.AddressNotFound"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INSTANCENOWANIP = "InvalidParameterValue.InstanceNoWanIP"
//  INVALIDPARAMETERVALUE_RESOURCENOTSUPPORT = "InvalidParameterValue.ResourceNotSupport"
//  INVALIDPARAMETERVALUE_UNAVAILABLEZONE = "InvalidParameterValue.UnavailableZone"
//  LIMITEXCEEDED_CHANGEADDRESSQUOTA = "LimitExceeded.ChangeAddressQuota"
//  LIMITEXCEEDED_DAILYCHANGEADDRESSQUOTA = "LimitExceeded.DailyChangeAddressQuota"
//  OPERATIONDENIED_ADDRESSINARREARS = "OperationDenied.AddressInArrears"
//  OPERATIONDENIED_MUTEXTASKRUNNING = "OperationDenied.MutexTaskRunning"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNSUPPORTEDOPERATION_ADDRESSSTATUSNOTPERMIT = "UnsupportedOperation.AddressStatusNotPermit"
//  UNSUPPORTEDOPERATION_INVALIDADDRESSINTERNETCHARGETYPE = "UnsupportedOperation.InvalidAddressInternetChargeType"
//  UNSUPPORTEDOPERATION_ISPNOTSUPPORTED = "UnsupportedOperation.IspNotSupported"
func (c *Client) AdjustPublicAddressWithContext(ctx context.Context, request *AdjustPublicAddressRequest) (response *AdjustPublicAddressResponse, err error) {
    if request == nil {
        request = NewAdjustPublicAddressRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AdjustPublicAddress require credential")
    }

    request.SetContext(ctx)
    
    response = NewAdjustPublicAddressResponse()
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
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_INVALIDREGION = "FailedOperation.InvalidRegion"
//  INVALIDACCOUNT_NOTSUPPORTED = "InvalidAccount.NotSupported"
//  INVALIDADDRESSID_BLOCKED = "InvalidAddressId.Blocked"
//  INVALIDPARAMETERCONFLICT = "InvalidParameterConflict"
//  INVALIDPARAMETERVALUE_ADDRESSATTACKED = "InvalidParameterValue.AddressAttacked"
//  INVALIDPARAMETERVALUE_ADDRESSIPNOTAVAILABLE = "InvalidParameterValue.AddressIpNotAvailable"
//  INVALIDPARAMETERVALUE_BANDWIDTHOUTOFRANGE = "InvalidParameterValue.BandwidthOutOfRange"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGEIDMALFORMED = "InvalidParameterValue.BandwidthPackageIdMalformed"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGENOTFOUND = "InvalidParameterValue.BandwidthPackageNotFound"
//  INVALIDPARAMETERVALUE_BANDWIDTHTOOSMALL = "InvalidParameterValue.BandwidthTooSmall"
//  INVALIDPARAMETERVALUE_COMBINATION = "InvalidParameterValue.Combination"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDDEDICATEDCLUSTERID = "InvalidParameterValue.InvalidDedicatedClusterId"
//  INVALIDPARAMETERVALUE_INVALIDTAG = "InvalidParameterValue.InvalidTag"
//  INVALIDPARAMETERVALUE_MIXEDADDRESSIPSETTYPE = "InvalidParameterValue.MixedAddressIpSetType"
//  INVALIDPARAMETERVALUE_RESOURCENOTSUPPORT = "InvalidParameterValue.ResourceNotSupport"
//  INVALIDPARAMETERVALUE_TAGNOTEXISTED = "InvalidParameterValue.TagNotExisted"
//  INVALIDPARAMETERVALUE_UNAVAILABLEZONE = "InvalidParameterValue.UnavailableZone"
//  LIMITEXCEEDED_BANDWIDTHPACKAGEQUOTA = "LimitExceeded.BandwidthPackageQuota"
//  LIMITEXCEEDED_MONTHLYADDRESSRECOVERYQUOTA = "LimitExceeded.MonthlyAddressRecoveryQuota"
//  LIMITEXCEEDED_TAGQUOTA = "LimitExceeded.TagQuota"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION_ANYCASTEIP = "UnauthorizedOperation.AnycastEip"
//  UNAUTHORIZEDOPERATION_INVALIDACCOUNT = "UnauthorizedOperation.InvalidAccount"
//  UNSUPPORTEDOPERATION_ACTIONNOTFOUND = "UnsupportedOperation.ActionNotFound"
//  UNSUPPORTEDOPERATION_BANDWIDTHPACKAGEIDNOTSUPPORTED = "UnsupportedOperation.BandwidthPackageIdNotSupported"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTSUPPORTED = "UnsupportedOperation.InstanceStateNotSupported"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_OFFLINECHARGETYPE = "UnsupportedOperation.OfflineChargeType"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDREGION = "UnsupportedOperation.UnsupportedRegion"
func (c *Client) AllocateAddresses(request *AllocateAddressesRequest) (response *AllocateAddressesResponse, err error) {
    return c.AllocateAddressesWithContext(context.Background(), request)
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
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_INVALIDREGION = "FailedOperation.InvalidRegion"
//  INVALIDACCOUNT_NOTSUPPORTED = "InvalidAccount.NotSupported"
//  INVALIDADDRESSID_BLOCKED = "InvalidAddressId.Blocked"
//  INVALIDPARAMETERCONFLICT = "InvalidParameterConflict"
//  INVALIDPARAMETERVALUE_ADDRESSATTACKED = "InvalidParameterValue.AddressAttacked"
//  INVALIDPARAMETERVALUE_ADDRESSIPNOTAVAILABLE = "InvalidParameterValue.AddressIpNotAvailable"
//  INVALIDPARAMETERVALUE_BANDWIDTHOUTOFRANGE = "InvalidParameterValue.BandwidthOutOfRange"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGEIDMALFORMED = "InvalidParameterValue.BandwidthPackageIdMalformed"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGENOTFOUND = "InvalidParameterValue.BandwidthPackageNotFound"
//  INVALIDPARAMETERVALUE_BANDWIDTHTOOSMALL = "InvalidParameterValue.BandwidthTooSmall"
//  INVALIDPARAMETERVALUE_COMBINATION = "InvalidParameterValue.Combination"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDDEDICATEDCLUSTERID = "InvalidParameterValue.InvalidDedicatedClusterId"
//  INVALIDPARAMETERVALUE_INVALIDTAG = "InvalidParameterValue.InvalidTag"
//  INVALIDPARAMETERVALUE_MIXEDADDRESSIPSETTYPE = "InvalidParameterValue.MixedAddressIpSetType"
//  INVALIDPARAMETERVALUE_RESOURCENOTSUPPORT = "InvalidParameterValue.ResourceNotSupport"
//  INVALIDPARAMETERVALUE_TAGNOTEXISTED = "InvalidParameterValue.TagNotExisted"
//  INVALIDPARAMETERVALUE_UNAVAILABLEZONE = "InvalidParameterValue.UnavailableZone"
//  LIMITEXCEEDED_BANDWIDTHPACKAGEQUOTA = "LimitExceeded.BandwidthPackageQuota"
//  LIMITEXCEEDED_MONTHLYADDRESSRECOVERYQUOTA = "LimitExceeded.MonthlyAddressRecoveryQuota"
//  LIMITEXCEEDED_TAGQUOTA = "LimitExceeded.TagQuota"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION_ANYCASTEIP = "UnauthorizedOperation.AnycastEip"
//  UNAUTHORIZEDOPERATION_INVALIDACCOUNT = "UnauthorizedOperation.InvalidAccount"
//  UNSUPPORTEDOPERATION_ACTIONNOTFOUND = "UnsupportedOperation.ActionNotFound"
//  UNSUPPORTEDOPERATION_BANDWIDTHPACKAGEIDNOTSUPPORTED = "UnsupportedOperation.BandwidthPackageIdNotSupported"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTSUPPORTED = "UnsupportedOperation.InstanceStateNotSupported"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_OFFLINECHARGETYPE = "UnsupportedOperation.OfflineChargeType"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDREGION = "UnsupportedOperation.UnsupportedRegion"
func (c *Client) AllocateAddressesWithContext(ctx context.Context, request *AllocateAddressesRequest) (response *AllocateAddressesResponse, err error) {
    if request == nil {
        request = NewAllocateAddressesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AllocateAddresses require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to apply for an IPv6 address for the ENI. <br />
//
// This API is completed asynchronously. If you need to query the execution result of an async task, please use the `RequestId` returned by this API to poll the `DescribeVpcTaskResult` API.
//
// * The number of IPs bound with an ENI is limited. For more information, see <a href="/document/product/576/18527">ENI Use Limits</a>.
//
// * You can apply for a specified IPv6 address. Currently, the IPv6 address can only be used as a secondary IP, instead of the primary IP.
//
// * The address must be an idle IP in the subnet to which the ENI belongs.
//
// * When applying for one or more secondary IPv6 addresses for an ENI, the API will return the specified number of secondary IPv6 addresses in the subnet range where the ENI is located.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED_ADDRESS = "LimitExceeded.Address"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE_ADDRESS = "ResourceInUse.Address"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
//  UNSUPPORTEDOPERATION_UNASSIGNCIDRBLOCK = "UnsupportedOperation.UnassignCidrBlock"
func (c *Client) AssignIpv6Addresses(request *AssignIpv6AddressesRequest) (response *AssignIpv6AddressesResponse, err error) {
    return c.AssignIpv6AddressesWithContext(context.Background(), request)
}

// AssignIpv6Addresses
// This API is used to apply for an IPv6 address for the ENI. <br />
//
// This API is completed asynchronously. If you need to query the execution result of an async task, please use the `RequestId` returned by this API to poll the `DescribeVpcTaskResult` API.
//
// * The number of IPs bound with an ENI is limited. For more information, see <a href="/document/product/576/18527">ENI Use Limits</a>.
//
// * You can apply for a specified IPv6 address. Currently, the IPv6 address can only be used as a secondary IP, instead of the primary IP.
//
// * The address must be an idle IP in the subnet to which the ENI belongs.
//
// * When applying for one or more secondary IPv6 addresses for an ENI, the API will return the specified number of secondary IPv6 addresses in the subnet range where the ENI is located.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED_ADDRESS = "LimitExceeded.Address"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE_ADDRESS = "ResourceInUse.Address"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
//  UNSUPPORTEDOPERATION_UNASSIGNCIDRBLOCK = "UnsupportedOperation.UnassignCidrBlock"
func (c *Client) AssignIpv6AddressesWithContext(ctx context.Context, request *AssignIpv6AddressesRequest) (response *AssignIpv6AddressesResponse, err error) {
    if request == nil {
        request = NewAssignIpv6AddressesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssignIpv6Addresses require credential")
    }

    request.SetContext(ctx)
    
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
    return c.AssignIpv6CidrBlockWithContext(context.Background(), request)
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
func (c *Client) AssignIpv6CidrBlockWithContext(ctx context.Context, request *AssignIpv6CidrBlockRequest) (response *AssignIpv6CidrBlockResponse, err error) {
    if request == nil {
        request = NewAssignIpv6CidrBlockRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssignIpv6CidrBlock require credential")
    }

    request.SetContext(ctx)
    
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
    return c.AssignIpv6SubnetCidrBlockWithContext(context.Background(), request)
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
func (c *Client) AssignIpv6SubnetCidrBlockWithContext(ctx context.Context, request *AssignIpv6SubnetCidrBlockRequest) (response *AssignIpv6SubnetCidrBlockResponse, err error) {
    if request == nil {
        request = NewAssignIpv6SubnetCidrBlockRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssignIpv6SubnetCidrBlock require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to apply for private IPs for an ENI.
//
// * An ENI can only be bound with a limited number of IP addresses. For more information about resource limits, see <a href="/document/product/576/18527">ENI Use Limits</a>.
//
// * You can apply for a specified private IP. It cannot be a primary IP because the primary IP already exists and cannot be modified. The private IP address must be an idle IP in the subnet to which the ENI belongs.
//
// * You can apply for more than one secondary private IP on the ENI. The API will return the specified number of secondary private IPs in the subnet IP range.
//
// >?This API is async. You can call the [`DescribeVpcTaskResult`](https://intl.cloud.tencent.com/document/api/215/59037?from_cn_redirect=1) API to query the task result. When the task is completed, you can continue other tasks.
//
// >
//
// error code that may be returned:
//  FAILEDOPERATION_NETDETECTTIMEOUT = "FailedOperation.NetDetectTimeOut"
//  INVALIDPARAMETERVALUE_DUPLICATEPARA = "InvalidParameterValue.DuplicatePara"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
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
//  UNSUPPORTEDOPERATION_RESOURCEMISMATCH = "UnsupportedOperation.ResourceMismatch"
func (c *Client) AssignPrivateIpAddresses(request *AssignPrivateIpAddressesRequest) (response *AssignPrivateIpAddressesResponse, err error) {
    return c.AssignPrivateIpAddressesWithContext(context.Background(), request)
}

// AssignPrivateIpAddresses
// This API is used to apply for private IPs for an ENI.
//
// * An ENI can only be bound with a limited number of IP addresses. For more information about resource limits, see <a href="/document/product/576/18527">ENI Use Limits</a>.
//
// * You can apply for a specified private IP. It cannot be a primary IP because the primary IP already exists and cannot be modified. The private IP address must be an idle IP in the subnet to which the ENI belongs.
//
// * You can apply for more than one secondary private IP on the ENI. The API will return the specified number of secondary private IPs in the subnet IP range.
//
// >?This API is async. You can call the [`DescribeVpcTaskResult`](https://intl.cloud.tencent.com/document/api/215/59037?from_cn_redirect=1) API to query the task result. When the task is completed, you can continue other tasks.
//
// >
//
// error code that may be returned:
//  FAILEDOPERATION_NETDETECTTIMEOUT = "FailedOperation.NetDetectTimeOut"
//  INVALIDPARAMETERVALUE_DUPLICATEPARA = "InvalidParameterValue.DuplicatePara"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
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
//  UNSUPPORTEDOPERATION_RESOURCEMISMATCH = "UnsupportedOperation.ResourceMismatch"
func (c *Client) AssignPrivateIpAddressesWithContext(ctx context.Context, request *AssignPrivateIpAddressesRequest) (response *AssignPrivateIpAddressesResponse, err error) {
    if request == nil {
        request = NewAssignPrivateIpAddressesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssignPrivateIpAddresses require credential")
    }

    request.SetContext(ctx)
    
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
// * Binding an EIP to a CVM instance is actually binding the EIP to the primary private IP of the primary ENI on the CVM instance.
//
// * When an EIP is bound, the public IP previously bound to the CVM instance will be unbound and released automatically.
//
// * To bind another EIP to the private IP of the specified ENI, you must first unbind the EIP.
//
// * To bind an EIP to a NAT Gateway, use the API [AssociateNatGatewayAddress](https://intl.cloud.tencent.com/document/product/215/36722?from_cn_redirect=1).
//
// * An EIP cannot be bound if it’s overdue or blocked
//
// * Only EIP in the `UNBIND` status can be bound.
//
// error code that may be returned:
//  FAILEDOPERATION_ADDRESSENIINFONOTFOUND = "FailedOperation.AddressEniInfoNotFound"
//  INVALIDACCOUNT_NOTSUPPORTED = "InvalidAccount.NotSupported"
//  INVALIDADDRESSID_BLOCKED = "InvalidAddressId.Blocked"
//  INVALIDADDRESSID_NOTFOUND = "InvalidAddressId.NotFound"
//  INVALIDINSTANCEID_ALREADYBINDEIP = "InvalidInstanceId.AlreadyBindEip"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDNETWORKINTERFACEID_NOTFOUND = "InvalidNetworkInterfaceId.NotFound"
//  INVALIDPARAMETERCONFLICT = "InvalidParameterConflict"
//  INVALIDPARAMETERVALUE_ADDRESSIDMALFORMED = "InvalidParameterValue.AddressIdMalformed"
//  INVALIDPARAMETERVALUE_ADDRESSNOTAPPLICABLE = "InvalidParameterValue.AddressNotApplicable"
//  INVALIDPARAMETERVALUE_ADDRESSNOTFOUND = "InvalidParameterValue.AddressNotFound"
//  INVALIDPARAMETERVALUE_COMBINATION = "InvalidParameterValue.Combination"
//  INVALIDPARAMETERVALUE_INSTANCEDOESNOTSUPPORTANYCAST = "InvalidParameterValue.InstanceDoesNotSupportAnycast"
//  INVALIDPARAMETERVALUE_INSTANCEHASWANIP = "InvalidParameterValue.InstanceHasWanIP"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INSTANCENOWANIP = "InvalidParameterValue.InstanceNoWanIP"
//  INVALIDPARAMETERVALUE_INSTANCENORMALPUBLICIPBLOCKED = "InvalidParameterValue.InstanceNormalPublicIpBlocked"
//  INVALIDPARAMETERVALUE_INSTANCENOTMATCHASSOCIATEENI = "InvalidParameterValue.InstanceNotMatchAssociateEni"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEINTERNETCHARGETYPE = "InvalidParameterValue.InvalidInstanceInternetChargeType"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCESTATE = "InvalidParameterValue.InvalidInstanceState"
//  INVALIDPARAMETERVALUE_LBALREADYBINDEIP = "InvalidParameterValue.LBAlreadyBindEip"
//  INVALIDPARAMETERVALUE_MISSINGASSOCIATEENTITY = "InvalidParameterValue.MissingAssociateEntity"
//  INVALIDPARAMETERVALUE_NETWORKINTERFACENOTFOUND = "InvalidParameterValue.NetworkInterfaceNotFound"
//  INVALIDPARAMETERVALUE_RESOURCENOTSUPPORT = "InvalidParameterValue.ResourceNotSupport"
//  INVALIDPRIVATEIPADDRESS_ALREADYBINDEIP = "InvalidPrivateIpAddress.AlreadyBindEip"
//  LIMITEXCEEDED_INSTANCEADDRESSQUOTA = "LimitExceeded.InstanceAddressQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_MUTEXTASKRUNNING = "OperationDenied.MutexTaskRunning"
//  UNSUPPORTEDOPERATION_ACTIONNOTFOUND = "UnsupportedOperation.ActionNotFound"
//  UNSUPPORTEDOPERATION_ADDRESSSTATUSNOTPERMIT = "UnsupportedOperation.AddressStatusNotPermit"
//  UNSUPPORTEDOPERATION_INCORRECTADDRESSRESOURCETYPE = "UnsupportedOperation.IncorrectAddressResourceType"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTSUPPORTED = "UnsupportedOperation.InstanceStateNotSupported"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_INVALIDADDRESSINTERNETCHARGETYPE = "UnsupportedOperation.InvalidAddressInternetChargeType"
//  UNSUPPORTEDOPERATION_ISPNOTSUPPORTED = "UnsupportedOperation.IspNotSupported"
func (c *Client) AssociateAddress(request *AssociateAddressRequest) (response *AssociateAddressResponse, err error) {
    return c.AssociateAddressWithContext(context.Background(), request)
}

// AssociateAddress
// This API is used to bind an [Elastic IP](https://intl.cloud.tencent.com/document/product/213/1941?from_cn_redirect=1) (EIP for short) to the specified private IP of an instance or ENI.
//
// * Binding an EIP to a CVM instance is actually binding the EIP to the primary private IP of the primary ENI on the CVM instance.
//
// * When an EIP is bound, the public IP previously bound to the CVM instance will be unbound and released automatically.
//
// * To bind another EIP to the private IP of the specified ENI, you must first unbind the EIP.
//
// * To bind an EIP to a NAT Gateway, use the API [AssociateNatGatewayAddress](https://intl.cloud.tencent.com/document/product/215/36722?from_cn_redirect=1).
//
// * An EIP cannot be bound if it’s overdue or blocked
//
// * Only EIP in the `UNBIND` status can be bound.
//
// error code that may be returned:
//  FAILEDOPERATION_ADDRESSENIINFONOTFOUND = "FailedOperation.AddressEniInfoNotFound"
//  INVALIDACCOUNT_NOTSUPPORTED = "InvalidAccount.NotSupported"
//  INVALIDADDRESSID_BLOCKED = "InvalidAddressId.Blocked"
//  INVALIDADDRESSID_NOTFOUND = "InvalidAddressId.NotFound"
//  INVALIDINSTANCEID_ALREADYBINDEIP = "InvalidInstanceId.AlreadyBindEip"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDNETWORKINTERFACEID_NOTFOUND = "InvalidNetworkInterfaceId.NotFound"
//  INVALIDPARAMETERCONFLICT = "InvalidParameterConflict"
//  INVALIDPARAMETERVALUE_ADDRESSIDMALFORMED = "InvalidParameterValue.AddressIdMalformed"
//  INVALIDPARAMETERVALUE_ADDRESSNOTAPPLICABLE = "InvalidParameterValue.AddressNotApplicable"
//  INVALIDPARAMETERVALUE_ADDRESSNOTFOUND = "InvalidParameterValue.AddressNotFound"
//  INVALIDPARAMETERVALUE_COMBINATION = "InvalidParameterValue.Combination"
//  INVALIDPARAMETERVALUE_INSTANCEDOESNOTSUPPORTANYCAST = "InvalidParameterValue.InstanceDoesNotSupportAnycast"
//  INVALIDPARAMETERVALUE_INSTANCEHASWANIP = "InvalidParameterValue.InstanceHasWanIP"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INSTANCENOWANIP = "InvalidParameterValue.InstanceNoWanIP"
//  INVALIDPARAMETERVALUE_INSTANCENORMALPUBLICIPBLOCKED = "InvalidParameterValue.InstanceNormalPublicIpBlocked"
//  INVALIDPARAMETERVALUE_INSTANCENOTMATCHASSOCIATEENI = "InvalidParameterValue.InstanceNotMatchAssociateEni"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEINTERNETCHARGETYPE = "InvalidParameterValue.InvalidInstanceInternetChargeType"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCESTATE = "InvalidParameterValue.InvalidInstanceState"
//  INVALIDPARAMETERVALUE_LBALREADYBINDEIP = "InvalidParameterValue.LBAlreadyBindEip"
//  INVALIDPARAMETERVALUE_MISSINGASSOCIATEENTITY = "InvalidParameterValue.MissingAssociateEntity"
//  INVALIDPARAMETERVALUE_NETWORKINTERFACENOTFOUND = "InvalidParameterValue.NetworkInterfaceNotFound"
//  INVALIDPARAMETERVALUE_RESOURCENOTSUPPORT = "InvalidParameterValue.ResourceNotSupport"
//  INVALIDPRIVATEIPADDRESS_ALREADYBINDEIP = "InvalidPrivateIpAddress.AlreadyBindEip"
//  LIMITEXCEEDED_INSTANCEADDRESSQUOTA = "LimitExceeded.InstanceAddressQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_MUTEXTASKRUNNING = "OperationDenied.MutexTaskRunning"
//  UNSUPPORTEDOPERATION_ACTIONNOTFOUND = "UnsupportedOperation.ActionNotFound"
//  UNSUPPORTEDOPERATION_ADDRESSSTATUSNOTPERMIT = "UnsupportedOperation.AddressStatusNotPermit"
//  UNSUPPORTEDOPERATION_INCORRECTADDRESSRESOURCETYPE = "UnsupportedOperation.IncorrectAddressResourceType"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTSUPPORTED = "UnsupportedOperation.InstanceStateNotSupported"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_INVALIDADDRESSINTERNETCHARGETYPE = "UnsupportedOperation.InvalidAddressInternetChargeType"
//  UNSUPPORTEDOPERATION_ISPNOTSUPPORTED = "UnsupportedOperation.IspNotSupported"
func (c *Client) AssociateAddressWithContext(ctx context.Context, request *AssociateAddressRequest) (response *AssociateAddressResponse, err error) {
    if request == nil {
        request = NewAssociateAddressRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssociateAddress require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_VPGTYPENOTMATCH = "InvalidParameterValue.VpgTypeNotMatch"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AssociateDirectConnectGatewayNatGateway(request *AssociateDirectConnectGatewayNatGatewayRequest) (response *AssociateDirectConnectGatewayNatGatewayResponse, err error) {
    return c.AssociateDirectConnectGatewayNatGatewayWithContext(context.Background(), request)
}

// AssociateDirectConnectGatewayNatGateway
// This API is used to bind a direct connect gateway with a NAT gateway,  and direct its default route to the NAT Gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_VPGTYPENOTMATCH = "InvalidParameterValue.VpgTypeNotMatch"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AssociateDirectConnectGatewayNatGatewayWithContext(ctx context.Context, request *AssociateDirectConnectGatewayNatGatewayRequest) (response *AssociateDirectConnectGatewayNatGatewayResponse, err error) {
    if request == nil {
        request = NewAssociateDirectConnectGatewayNatGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssociateDirectConnectGatewayNatGateway require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to bind an EIP to a NAT gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_EIPBRANDWIDTHOUTINVALID = "InvalidParameterValue.EIPBrandWidthOutInvalid"
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
//  UNSUPPORTEDOPERATION_PUBLICIPADDRESSISNOTEXISTED = "UnsupportedOperation.PublicIpAddressIsNotExisted"
//  UNSUPPORTEDOPERATION_PUBLICIPADDRESSNOTBILLEDBYTRAFFIC = "UnsupportedOperation.PublicIpAddressNotBilledByTraffic"
func (c *Client) AssociateNatGatewayAddress(request *AssociateNatGatewayAddressRequest) (response *AssociateNatGatewayAddressResponse, err error) {
    return c.AssociateNatGatewayAddressWithContext(context.Background(), request)
}

// AssociateNatGatewayAddress
// This API is used to bind an EIP to a NAT gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_EIPBRANDWIDTHOUTINVALID = "InvalidParameterValue.EIPBrandWidthOutInvalid"
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
//  UNSUPPORTEDOPERATION_PUBLICIPADDRESSISNOTEXISTED = "UnsupportedOperation.PublicIpAddressIsNotExisted"
//  UNSUPPORTEDOPERATION_PUBLICIPADDRESSNOTBILLEDBYTRAFFIC = "UnsupportedOperation.PublicIpAddressNotBilledByTraffic"
func (c *Client) AssociateNatGatewayAddressWithContext(ctx context.Context, request *AssociateNatGatewayAddressRequest) (response *AssociateNatGatewayAddressResponse, err error) {
    if request == nil {
        request = NewAssociateNatGatewayAddressRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssociateNatGatewayAddress require credential")
    }

    request.SetContext(ctx)
    
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
    return c.AssociateNetworkAclSubnetsWithContext(context.Background(), request)
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
func (c *Client) AssociateNetworkAclSubnetsWithContext(ctx context.Context, request *AssociateNetworkAclSubnetsRequest) (response *AssociateNetworkAclSubnetsResponse, err error) {
    if request == nil {
        request = NewAssociateNetworkAclSubnetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssociateNetworkAclSubnets require credential")
    }

    request.SetContext(ctx)
    
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
    return c.AssociateNetworkInterfaceSecurityGroupsWithContext(context.Background(), request)
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
func (c *Client) AssociateNetworkInterfaceSecurityGroupsWithContext(ctx context.Context, request *AssociateNetworkInterfaceSecurityGroupsRequest) (response *AssociateNetworkInterfaceSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewAssociateNetworkInterfaceSecurityGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssociateNetworkInterfaceSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
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
//  UNSUPPORTEDOPERATION_APPIDNOTFOUND = "UnsupportedOperation.AppIdNotFound"
//  UNSUPPORTEDOPERATION_CCNATTACHED = "UnsupportedOperation.CcnAttached"
//  UNSUPPORTEDOPERATION_CCNORDINARYACCOUNTREFUSEATTACH = "UnsupportedOperation.CcnOrdinaryAccountRefuseAttach"
//  UNSUPPORTEDOPERATION_CCNROUTETABLENOTEXIST = "UnsupportedOperation.CcnRouteTableNotExist"
//  UNSUPPORTEDOPERATION_INSTANCEANDRTBNOTMATCH = "UnsupportedOperation.InstanceAndRtbNotMatch"
//  UNSUPPORTEDOPERATION_INSTANCEORDINARYACCOUNTREFUSEATTACH = "UnsupportedOperation.InstanceOrdinaryAccountRefuseAttach"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_ISNOTFINANCEACCOUNT = "UnsupportedOperation.IsNotFinanceAccount"
//  UNSUPPORTEDOPERATION_PURCHASELIMIT = "UnsupportedOperation.PurchaseLimit"
//  UNSUPPORTEDOPERATION_UINNOTFOUND = "UnsupportedOperation.UinNotFound"
//  UNSUPPORTEDOPERATION_UNABLECROSSBORDER = "UnsupportedOperation.UnableCrossBorder"
//  UNSUPPORTEDOPERATION_UNABLECROSSFINANCE = "UnsupportedOperation.UnableCrossFinance"
func (c *Client) AttachCcnInstances(request *AttachCcnInstancesRequest) (response *AttachCcnInstancesResponse, err error) {
    return c.AttachCcnInstancesWithContext(context.Background(), request)
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
//  UNSUPPORTEDOPERATION_APPIDNOTFOUND = "UnsupportedOperation.AppIdNotFound"
//  UNSUPPORTEDOPERATION_CCNATTACHED = "UnsupportedOperation.CcnAttached"
//  UNSUPPORTEDOPERATION_CCNORDINARYACCOUNTREFUSEATTACH = "UnsupportedOperation.CcnOrdinaryAccountRefuseAttach"
//  UNSUPPORTEDOPERATION_CCNROUTETABLENOTEXIST = "UnsupportedOperation.CcnRouteTableNotExist"
//  UNSUPPORTEDOPERATION_INSTANCEANDRTBNOTMATCH = "UnsupportedOperation.InstanceAndRtbNotMatch"
//  UNSUPPORTEDOPERATION_INSTANCEORDINARYACCOUNTREFUSEATTACH = "UnsupportedOperation.InstanceOrdinaryAccountRefuseAttach"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_ISNOTFINANCEACCOUNT = "UnsupportedOperation.IsNotFinanceAccount"
//  UNSUPPORTEDOPERATION_PURCHASELIMIT = "UnsupportedOperation.PurchaseLimit"
//  UNSUPPORTEDOPERATION_UINNOTFOUND = "UnsupportedOperation.UinNotFound"
//  UNSUPPORTEDOPERATION_UNABLECROSSBORDER = "UnsupportedOperation.UnableCrossBorder"
//  UNSUPPORTEDOPERATION_UNABLECROSSFINANCE = "UnsupportedOperation.UnableCrossFinance"
func (c *Client) AttachCcnInstancesWithContext(ctx context.Context, request *AttachCcnInstancesRequest) (response *AttachCcnInstancesResponse, err error) {
    if request == nil {
        request = NewAttachCcnInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AttachCcnInstances require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to create a Classiclink between a VPC instance and a classic network device.
//
// * The VPC instance and the classic network device must be in the same region.
//
// * For differences between VPC and the classic network, see <a href="https://intl.cloud.tencent.com/document/product/215/30720?from_cn_redirect=1">VPC and Classic Network</a>.
//
// >?This API is async. You can call the [`DescribeVpcTaskResult`](https://intl.cloud.tencent.com/document/api/215/59037?from_cn_redirect=1) API to query the task result. When the task is completed, you can continue other tasks.
//
// >
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
    return c.AttachClassicLinkVpcWithContext(context.Background(), request)
}

// AttachClassicLinkVpc
// This API is used to create a Classiclink between a VPC instance and a classic network device.
//
// * The VPC instance and the classic network device must be in the same region.
//
// * For differences between VPC and the classic network, see <a href="https://intl.cloud.tencent.com/document/product/215/30720?from_cn_redirect=1">VPC and Classic Network</a>.
//
// >?This API is async. You can call the [`DescribeVpcTaskResult`](https://intl.cloud.tencent.com/document/api/215/59037?from_cn_redirect=1) API to query the task result. When the task is completed, you can continue other tasks.
//
// >
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CIDRUNSUPPORTEDCLASSICLINK = "UnsupportedOperation.CIDRUnSupportedClassicLink"
//  UNSUPPORTEDOPERATION_CLASSICINSTANCEIDALREADYEXISTS = "UnsupportedOperation.ClassicInstanceIdAlreadyExists"
func (c *Client) AttachClassicLinkVpcWithContext(ctx context.Context, request *AttachClassicLinkVpcRequest) (response *AttachClassicLinkVpcResponse, err error) {
    if request == nil {
        request = NewAttachClassicLinkVpcRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AttachClassicLinkVpc require credential")
    }

    request.SetContext(ctx)
    
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
// * An ENI must be bound with one security group at least. To bind it, see <a href="https://intl.cloud.tencent.com/document/product/215/43132?from_cn_redirect=1">AssociateNetworkInterfaceSecurityGroups</a>.
//
// * One CVM can be bound with multiple ENIs, but only one can be the primary ENI. For more information about the limits, see <a href="https://intl.cloud.tencent.com/document/product/576/18527?from_cn_redirect=1">ENI Use Limits</a>.
//
// * An ENI can only be bound to one CVM.
//
// * Only the running or shutdown CVMs can be bound with ENIs. For more information about the CVM status, see <a href="https://intl.cloud.tencent.com/document/api/213/9452?from_cn_redirect=1#InstanceStatus">InstanceStatus</a> in the Data Types.
//
// * An ENI can only be bound to a VPC-based CVM under the same availability zone as the ENI subnet.
//
// 
//
// This API is completed asynchronously. If you need to query the execution result of an async task, please use the `RequestId` returned by this API to poll the `DescribeVpcTaskResult` API.
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
    return c.AttachNetworkInterfaceWithContext(context.Background(), request)
}

// AttachNetworkInterface
// This API is used to bind an ENI to a CVM.
//
// * An ENI must be bound with one security group at least. To bind it, see <a href="https://intl.cloud.tencent.com/document/product/215/43132?from_cn_redirect=1">AssociateNetworkInterfaceSecurityGroups</a>.
//
// * One CVM can be bound with multiple ENIs, but only one can be the primary ENI. For more information about the limits, see <a href="https://intl.cloud.tencent.com/document/product/576/18527?from_cn_redirect=1">ENI Use Limits</a>.
//
// * An ENI can only be bound to one CVM.
//
// * Only the running or shutdown CVMs can be bound with ENIs. For more information about the CVM status, see <a href="https://intl.cloud.tencent.com/document/api/213/9452?from_cn_redirect=1#InstanceStatus">InstanceStatus</a> in the Data Types.
//
// * An ENI can only be bound to a VPC-based CVM under the same availability zone as the ENI subnet.
//
// 
//
// This API is completed asynchronously. If you need to query the execution result of an async task, please use the `RequestId` returned by this API to poll the `DescribeVpcTaskResult` API.
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
func (c *Client) AttachNetworkInterfaceWithContext(ctx context.Context, request *AttachNetworkInterfaceRequest) (response *AttachNetworkInterfaceResponse, err error) {
    if request == nil {
        request = NewAttachNetworkInterfaceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AttachNetworkInterface require credential")
    }

    request.SetContext(ctx)
    
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
    return c.AuditCrossBorderComplianceWithContext(context.Background(), request)
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
func (c *Client) AuditCrossBorderComplianceWithContext(ctx context.Context, request *AuditCrossBorderComplianceRequest) (response *AuditCrossBorderComplianceResponse, err error) {
    if request == nil {
        request = NewAuditCrossBorderComplianceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AuditCrossBorderCompliance require credential")
    }

    request.SetContext(ctx)
    
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
    return c.CheckAssistantCidrWithContext(context.Background(), request)
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
func (c *Client) CheckAssistantCidrWithContext(ctx context.Context, request *CheckAssistantCidrRequest) (response *CheckAssistantCidrResponse, err error) {
    if request == nil {
        request = NewCheckAssistantCidrRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckAssistantCidr require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION_NETDETECTTIMEOUT = "FailedOperation.NetDetectTimeOut"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_NEXTHOPMISMATCH = "InvalidParameter.NextHopMismatch"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NETDETECTINVPC = "InvalidParameterValue.NetDetectInVpc"
//  INVALIDPARAMETERVALUE_NETDETECTNOTFOUNDIP = "InvalidParameterValue.NetDetectNotFoundIp"
//  INVALIDPARAMETERVALUE_NETDETECTSAMEIP = "InvalidParameterValue.NetDetectSameIp"
//  INVALIDPARAMETERVALUE_RESOURCENOTFOUND = "InvalidParameterValue.ResourceNotFound"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_CONFLICTWITHDOCKERROUTE = "UnsupportedOperation.ConflictWithDockerRoute"
//  UNSUPPORTEDOPERATION_ECMPWITHUSERROUTE = "UnsupportedOperation.EcmpWithUserRoute"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
func (c *Client) CheckNetDetectState(request *CheckNetDetectStateRequest) (response *CheckNetDetectStateResponse, err error) {
    return c.CheckNetDetectStateWithContext(context.Background(), request)
}

// CheckNetDetectState
// This API is used to verify the network detection status.
//
// error code that may be returned:
//  FAILEDOPERATION_NETDETECTTIMEOUT = "FailedOperation.NetDetectTimeOut"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_NEXTHOPMISMATCH = "InvalidParameter.NextHopMismatch"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NETDETECTINVPC = "InvalidParameterValue.NetDetectInVpc"
//  INVALIDPARAMETERVALUE_NETDETECTNOTFOUNDIP = "InvalidParameterValue.NetDetectNotFoundIp"
//  INVALIDPARAMETERVALUE_NETDETECTSAMEIP = "InvalidParameterValue.NetDetectSameIp"
//  INVALIDPARAMETERVALUE_RESOURCENOTFOUND = "InvalidParameterValue.ResourceNotFound"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_CONFLICTWITHDOCKERROUTE = "UnsupportedOperation.ConflictWithDockerRoute"
//  UNSUPPORTEDOPERATION_ECMPWITHUSERROUTE = "UnsupportedOperation.EcmpWithUserRoute"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
func (c *Client) CheckNetDetectStateWithContext(ctx context.Context, request *CheckNetDetectStateRequest) (response *CheckNetDetectStateResponse, err error) {
    if request == nil {
        request = NewCheckNetDetectStateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckNetDetectState require credential")
    }

    request.SetContext(ctx)
    
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
    return c.CloneSecurityGroupWithContext(context.Background(), request)
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
func (c *Client) CloneSecurityGroupWithContext(ctx context.Context, request *CloneSecurityGroupRequest) (response *CloneSecurityGroupResponse, err error) {
    if request == nil {
        request = NewCloneSecurityGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloneSecurityGroup require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to create an IP address template.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateAddressTemplate(request *CreateAddressTemplateRequest) (response *CreateAddressTemplateResponse, err error) {
    return c.CreateAddressTemplateWithContext(context.Background(), request)
}

// CreateAddressTemplate
// This API is used to create an IP address template.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateAddressTemplateWithContext(ctx context.Context, request *CreateAddressTemplateRequest) (response *CreateAddressTemplateResponse, err error) {
    if request == nil {
        request = NewCreateAddressTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAddressTemplate require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to create an IP address template group.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) CreateAddressTemplateGroup(request *CreateAddressTemplateGroupRequest) (response *CreateAddressTemplateGroupResponse, err error) {
    return c.CreateAddressTemplateGroupWithContext(context.Background(), request)
}

// CreateAddressTemplateGroup
// This API is used to create an IP address template group.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) CreateAddressTemplateGroupWithContext(ctx context.Context, request *CreateAddressTemplateGroupRequest) (response *CreateAddressTemplateGroupResponse, err error) {
    if request == nil {
        request = NewCreateAddressTemplateGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAddressTemplateGroup require credential")
    }

    request.SetContext(ctx)
    
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
// >?This API is async. You can call the [`DescribeVpcTaskResult`](https://intl.cloud.tencent.com/document/api/215/59037?from_cn_redirect=1) API to query the task result. When the task is completed, you can continue other tasks.
//
// >
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
    return c.CreateAndAttachNetworkInterfaceWithContext(context.Background(), request)
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
// >?This API is async. You can call the [`DescribeVpcTaskResult`](https://intl.cloud.tencent.com/document/api/215/59037?from_cn_redirect=1) API to query the task result. When the task is completed, you can continue other tasks.
//
// >
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
func (c *Client) CreateAndAttachNetworkInterfaceWithContext(ctx context.Context, request *CreateAndAttachNetworkInterfaceRequest) (response *CreateAndAttachNetworkInterfaceResponse, err error) {
    if request == nil {
        request = NewCreateAndAttachNetworkInterfaceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAndAttachNetworkInterface require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to batch create secondary CIDR blocks. This API is in beta test. To use it, please submit a ticket.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_SUBNETCONFLICT = "InvalidParameterValue.SubnetConflict"
//  INVALIDPARAMETERVALUE_SUBNETOVERLAPASSISTCIDR = "InvalidParameterValue.SubnetOverlapAssistCidr"
//  INVALIDPARAMETERVALUE_SUBNETRANGE = "InvalidParameterValue.SubnetRange"
//  INVALIDPARAMETERVALUE_VPCCIDRCONFLICT = "InvalidParameterValue.VpcCidrConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateAssistantCidr(request *CreateAssistantCidrRequest) (response *CreateAssistantCidrResponse, err error) {
    return c.CreateAssistantCidrWithContext(context.Background(), request)
}

// CreateAssistantCidr
// This API is used to batch create secondary CIDR blocks. This API is in beta test. To use it, please submit a ticket.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_SUBNETCONFLICT = "InvalidParameterValue.SubnetConflict"
//  INVALIDPARAMETERVALUE_SUBNETOVERLAPASSISTCIDR = "InvalidParameterValue.SubnetOverlapAssistCidr"
//  INVALIDPARAMETERVALUE_SUBNETRANGE = "InvalidParameterValue.SubnetRange"
//  INVALIDPARAMETERVALUE_VPCCIDRCONFLICT = "InvalidParameterValue.VpcCidrConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateAssistantCidrWithContext(ctx context.Context, request *CreateAssistantCidrRequest) (response *CreateAssistantCidrResponse, err error) {
    if request == nil {
        request = NewCreateAssistantCidrRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAssistantCidr require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to create a [device bandwidth package](https://intl.cloud.tencent.com/document/product/684/15245?from_cn_redirect=1#bwptype) or an [IP bandwidth package](https://intl.cloud.tencent.com/document/product/684/15245?from_cn_redirect=1#bwptype).
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDACCOUNT_NOTSUPPORTED = "InvalidAccount.NotSupported"
//  INVALIDPARAMETERVALUE_COMBINATION = "InvalidParameterValue.Combination"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TAGNOTEXISTED = "InvalidParameterValue.TagNotExisted"
//  LIMITEXCEEDED_BANDWIDTHPACKAGEQUOTA = "LimitExceeded.BandwidthPackageQuota"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTSUPPORTED = "UnsupportedOperation.InstanceStateNotSupported"
//  UNSUPPORTEDOPERATION_INVALIDRESOURCEINTERNETCHARGETYPE = "UnsupportedOperation.InvalidResourceInternetChargeType"
func (c *Client) CreateBandwidthPackage(request *CreateBandwidthPackageRequest) (response *CreateBandwidthPackageResponse, err error) {
    return c.CreateBandwidthPackageWithContext(context.Background(), request)
}

// CreateBandwidthPackage
// This API is used to create a [device bandwidth package](https://intl.cloud.tencent.com/document/product/684/15245?from_cn_redirect=1#bwptype) or an [IP bandwidth package](https://intl.cloud.tencent.com/document/product/684/15245?from_cn_redirect=1#bwptype).
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDACCOUNT_NOTSUPPORTED = "InvalidAccount.NotSupported"
//  INVALIDPARAMETERVALUE_COMBINATION = "InvalidParameterValue.Combination"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TAGNOTEXISTED = "InvalidParameterValue.TagNotExisted"
//  LIMITEXCEEDED_BANDWIDTHPACKAGEQUOTA = "LimitExceeded.BandwidthPackageQuota"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTSUPPORTED = "UnsupportedOperation.InstanceStateNotSupported"
//  UNSUPPORTEDOPERATION_INVALIDRESOURCEINTERNETCHARGETYPE = "UnsupportedOperation.InvalidResourceInternetChargeType"
func (c *Client) CreateBandwidthPackageWithContext(ctx context.Context, request *CreateBandwidthPackageRequest) (response *CreateBandwidthPackageResponse, err error) {
    if request == nil {
        request = NewCreateBandwidthPackageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBandwidthPackage require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_TAGDUPLICATEKEY = "InvalidParameterValue.TagDuplicateKey"
//  INVALIDPARAMETERVALUE_TAGDUPLICATERESOURCETYPE = "InvalidParameterValue.TagDuplicateResourceType"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEY = "InvalidParameterValue.TagInvalidKey"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEYLEN = "InvalidParameterValue.TagInvalidKeyLen"
//  INVALIDPARAMETERVALUE_TAGINVALIDVAL = "InvalidParameterValue.TagInvalidVal"
//  INVALIDPARAMETERVALUE_TAGKEYNOTEXISTS = "InvalidParameterValue.TagKeyNotExists"
//  INVALIDPARAMETERVALUE_TAGNOTALLOCATEDQUOTA = "InvalidParameterValue.TagNotAllocatedQuota"
//  INVALIDPARAMETERVALUE_TAGNOTEXISTED = "InvalidParameterValue.TagNotExisted"
//  INVALIDPARAMETERVALUE_TAGNOTSUPPORTTAG = "InvalidParameterValue.TagNotSupportTag"
//  INVALIDPARAMETERVALUE_TAGRESOURCEFORMATERROR = "InvalidParameterValue.TagResourceFormatError"
//  INVALIDPARAMETERVALUE_TAGTIMESTAMPEXCEEDED = "InvalidParameterValue.TagTimestampExceeded"
//  INVALIDPARAMETERVALUE_TAGVALNOTEXISTS = "InvalidParameterValue.TagValNotExists"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TAGKEYEXCEEDED = "LimitExceeded.TagKeyExceeded"
//  LIMITEXCEEDED_TAGKEYPERRESOURCEEXCEEDED = "LimitExceeded.TagKeyPerResourceExceeded"
//  LIMITEXCEEDED_TAGNOTENOUGHQUOTA = "LimitExceeded.TagNotEnoughQuota"
//  LIMITEXCEEDED_TAGQUOTA = "LimitExceeded.TagQuota"
//  LIMITEXCEEDED_TAGQUOTAEXCEEDED = "LimitExceeded.TagQuotaExceeded"
//  LIMITEXCEEDED_TAGTAGSEXCEEDED = "LimitExceeded.TagTagsExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_NOREALNAMEAUTHENTICATION = "UnauthorizedOperation.NoRealNameAuthentication"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INSUFFICIENTFUNDS = "UnsupportedOperation.InsufficientFunds"
//  UNSUPPORTEDOPERATION_PREPAIDCCNONLYSUPPORTINTERREGIONLIMIT = "UnsupportedOperation.PrepaidCcnOnlySupportInterRegionLimit"
//  UNSUPPORTEDOPERATION_TAGALLOCATE = "UnsupportedOperation.TagAllocate"
//  UNSUPPORTEDOPERATION_TAGFREE = "UnsupportedOperation.TagFree"
//  UNSUPPORTEDOPERATION_TAGNOTPERMIT = "UnsupportedOperation.TagNotPermit"
//  UNSUPPORTEDOPERATION_TAGSYSTEMRESERVEDTAGKEY = "UnsupportedOperation.TagSystemReservedTagKey"
//  UNSUPPORTEDOPERATION_USERANDCCNCHARGETYPENOTMATCH = "UnsupportedOperation.UserAndCcnChargeTypeNotMatch"
func (c *Client) CreateCcn(request *CreateCcnRequest) (response *CreateCcnResponse, err error) {
    return c.CreateCcnWithContext(context.Background(), request)
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
//  INVALIDPARAMETERVALUE_TAGDUPLICATEKEY = "InvalidParameterValue.TagDuplicateKey"
//  INVALIDPARAMETERVALUE_TAGDUPLICATERESOURCETYPE = "InvalidParameterValue.TagDuplicateResourceType"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEY = "InvalidParameterValue.TagInvalidKey"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEYLEN = "InvalidParameterValue.TagInvalidKeyLen"
//  INVALIDPARAMETERVALUE_TAGINVALIDVAL = "InvalidParameterValue.TagInvalidVal"
//  INVALIDPARAMETERVALUE_TAGKEYNOTEXISTS = "InvalidParameterValue.TagKeyNotExists"
//  INVALIDPARAMETERVALUE_TAGNOTALLOCATEDQUOTA = "InvalidParameterValue.TagNotAllocatedQuota"
//  INVALIDPARAMETERVALUE_TAGNOTEXISTED = "InvalidParameterValue.TagNotExisted"
//  INVALIDPARAMETERVALUE_TAGNOTSUPPORTTAG = "InvalidParameterValue.TagNotSupportTag"
//  INVALIDPARAMETERVALUE_TAGRESOURCEFORMATERROR = "InvalidParameterValue.TagResourceFormatError"
//  INVALIDPARAMETERVALUE_TAGTIMESTAMPEXCEEDED = "InvalidParameterValue.TagTimestampExceeded"
//  INVALIDPARAMETERVALUE_TAGVALNOTEXISTS = "InvalidParameterValue.TagValNotExists"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TAGKEYEXCEEDED = "LimitExceeded.TagKeyExceeded"
//  LIMITEXCEEDED_TAGKEYPERRESOURCEEXCEEDED = "LimitExceeded.TagKeyPerResourceExceeded"
//  LIMITEXCEEDED_TAGNOTENOUGHQUOTA = "LimitExceeded.TagNotEnoughQuota"
//  LIMITEXCEEDED_TAGQUOTA = "LimitExceeded.TagQuota"
//  LIMITEXCEEDED_TAGQUOTAEXCEEDED = "LimitExceeded.TagQuotaExceeded"
//  LIMITEXCEEDED_TAGTAGSEXCEEDED = "LimitExceeded.TagTagsExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_NOREALNAMEAUTHENTICATION = "UnauthorizedOperation.NoRealNameAuthentication"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INSUFFICIENTFUNDS = "UnsupportedOperation.InsufficientFunds"
//  UNSUPPORTEDOPERATION_PREPAIDCCNONLYSUPPORTINTERREGIONLIMIT = "UnsupportedOperation.PrepaidCcnOnlySupportInterRegionLimit"
//  UNSUPPORTEDOPERATION_TAGALLOCATE = "UnsupportedOperation.TagAllocate"
//  UNSUPPORTEDOPERATION_TAGFREE = "UnsupportedOperation.TagFree"
//  UNSUPPORTEDOPERATION_TAGNOTPERMIT = "UnsupportedOperation.TagNotPermit"
//  UNSUPPORTEDOPERATION_TAGSYSTEMRESERVEDTAGKEY = "UnsupportedOperation.TagSystemReservedTagKey"
//  UNSUPPORTEDOPERATION_USERANDCCNCHARGETYPENOTMATCH = "UnsupportedOperation.UserAndCcnChargeTypeNotMatch"
func (c *Client) CreateCcnWithContext(ctx context.Context, request *CreateCcnRequest) (response *CreateCcnResponse, err error) {
    if request == nil {
        request = NewCreateCcnRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCcn require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_TAGDUPLICATEKEY = "InvalidParameterValue.TagDuplicateKey"
//  INVALIDPARAMETERVALUE_TAGDUPLICATERESOURCETYPE = "InvalidParameterValue.TagDuplicateResourceType"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEY = "InvalidParameterValue.TagInvalidKey"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEYLEN = "InvalidParameterValue.TagInvalidKeyLen"
//  INVALIDPARAMETERVALUE_TAGINVALIDVAL = "InvalidParameterValue.TagInvalidVal"
//  INVALIDPARAMETERVALUE_TAGKEYNOTEXISTS = "InvalidParameterValue.TagKeyNotExists"
//  INVALIDPARAMETERVALUE_TAGNOTALLOCATEDQUOTA = "InvalidParameterValue.TagNotAllocatedQuota"
//  INVALIDPARAMETERVALUE_TAGNOTEXISTED = "InvalidParameterValue.TagNotExisted"
//  INVALIDPARAMETERVALUE_TAGNOTSUPPORTTAG = "InvalidParameterValue.TagNotSupportTag"
//  INVALIDPARAMETERVALUE_TAGRESOURCEFORMATERROR = "InvalidParameterValue.TagResourceFormatError"
//  INVALIDPARAMETERVALUE_TAGTIMESTAMPEXCEEDED = "InvalidParameterValue.TagTimestampExceeded"
//  INVALIDPARAMETERVALUE_TAGVALNOTEXISTS = "InvalidParameterValue.TagValNotExists"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TAGKEYEXCEEDED = "LimitExceeded.TagKeyExceeded"
//  LIMITEXCEEDED_TAGKEYPERRESOURCEEXCEEDED = "LimitExceeded.TagKeyPerResourceExceeded"
//  LIMITEXCEEDED_TAGNOTENOUGHQUOTA = "LimitExceeded.TagNotEnoughQuota"
//  LIMITEXCEEDED_TAGQUOTA = "LimitExceeded.TagQuota"
//  LIMITEXCEEDED_TAGQUOTAEXCEEDED = "LimitExceeded.TagQuotaExceeded"
//  LIMITEXCEEDED_TAGTAGSEXCEEDED = "LimitExceeded.TagTagsExceeded"
//  UNSUPPORTEDOPERATION_TAGALLOCATE = "UnsupportedOperation.TagAllocate"
//  UNSUPPORTEDOPERATION_TAGFREE = "UnsupportedOperation.TagFree"
//  UNSUPPORTEDOPERATION_TAGNOTPERMIT = "UnsupportedOperation.TagNotPermit"
//  UNSUPPORTEDOPERATION_TAGSYSTEMRESERVEDTAGKEY = "UnsupportedOperation.TagSystemReservedTagKey"
//  VPCLIMITEXCEEDED = "VpcLimitExceeded"
func (c *Client) CreateCustomerGateway(request *CreateCustomerGatewayRequest) (response *CreateCustomerGatewayResponse, err error) {
    return c.CreateCustomerGatewayWithContext(context.Background(), request)
}

// CreateCustomerGateway
// This API (CreateCustomerGateway) is used to create customer gateways.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_TAGDUPLICATEKEY = "InvalidParameterValue.TagDuplicateKey"
//  INVALIDPARAMETERVALUE_TAGDUPLICATERESOURCETYPE = "InvalidParameterValue.TagDuplicateResourceType"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEY = "InvalidParameterValue.TagInvalidKey"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEYLEN = "InvalidParameterValue.TagInvalidKeyLen"
//  INVALIDPARAMETERVALUE_TAGINVALIDVAL = "InvalidParameterValue.TagInvalidVal"
//  INVALIDPARAMETERVALUE_TAGKEYNOTEXISTS = "InvalidParameterValue.TagKeyNotExists"
//  INVALIDPARAMETERVALUE_TAGNOTALLOCATEDQUOTA = "InvalidParameterValue.TagNotAllocatedQuota"
//  INVALIDPARAMETERVALUE_TAGNOTEXISTED = "InvalidParameterValue.TagNotExisted"
//  INVALIDPARAMETERVALUE_TAGNOTSUPPORTTAG = "InvalidParameterValue.TagNotSupportTag"
//  INVALIDPARAMETERVALUE_TAGRESOURCEFORMATERROR = "InvalidParameterValue.TagResourceFormatError"
//  INVALIDPARAMETERVALUE_TAGTIMESTAMPEXCEEDED = "InvalidParameterValue.TagTimestampExceeded"
//  INVALIDPARAMETERVALUE_TAGVALNOTEXISTS = "InvalidParameterValue.TagValNotExists"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TAGKEYEXCEEDED = "LimitExceeded.TagKeyExceeded"
//  LIMITEXCEEDED_TAGKEYPERRESOURCEEXCEEDED = "LimitExceeded.TagKeyPerResourceExceeded"
//  LIMITEXCEEDED_TAGNOTENOUGHQUOTA = "LimitExceeded.TagNotEnoughQuota"
//  LIMITEXCEEDED_TAGQUOTA = "LimitExceeded.TagQuota"
//  LIMITEXCEEDED_TAGQUOTAEXCEEDED = "LimitExceeded.TagQuotaExceeded"
//  LIMITEXCEEDED_TAGTAGSEXCEEDED = "LimitExceeded.TagTagsExceeded"
//  UNSUPPORTEDOPERATION_TAGALLOCATE = "UnsupportedOperation.TagAllocate"
//  UNSUPPORTEDOPERATION_TAGFREE = "UnsupportedOperation.TagFree"
//  UNSUPPORTEDOPERATION_TAGNOTPERMIT = "UnsupportedOperation.TagNotPermit"
//  UNSUPPORTEDOPERATION_TAGSYSTEMRESERVEDTAGKEY = "UnsupportedOperation.TagSystemReservedTagKey"
//  VPCLIMITEXCEEDED = "VpcLimitExceeded"
func (c *Client) CreateCustomerGatewayWithContext(ctx context.Context, request *CreateCustomerGatewayRequest) (response *CreateCustomerGatewayResponse, err error) {
    if request == nil {
        request = NewCreateCustomerGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCustomerGateway require credential")
    }

    request.SetContext(ctx)
    
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
    return c.CreateDefaultVpcWithContext(context.Background(), request)
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
func (c *Client) CreateDefaultVpcWithContext(ctx context.Context, request *CreateDefaultVpcRequest) (response *CreateDefaultVpcResponse, err error) {
    if request == nil {
        request = NewCreateDefaultVpcRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDefaultVpc require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETER_VPGHAGROUPNOTFOUND = "InvalidParameter.VpgHaGroupNotFound"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CCNROUTETABLENOTEXIST = "UnsupportedOperation.CcnRouteTableNotExist"
//  UNSUPPORTEDOPERATION_UNABLECROSSBORDER = "UnsupportedOperation.UnableCrossBorder"
func (c *Client) CreateDirectConnectGateway(request *CreateDirectConnectGatewayRequest) (response *CreateDirectConnectGatewayResponse, err error) {
    return c.CreateDirectConnectGatewayWithContext(context.Background(), request)
}

// CreateDirectConnectGateway
// This API is used to create a direct connect gateway.
//
// error code that may be returned:
//  INVALIDPARAMETER_VPGHAGROUPNOTFOUND = "InvalidParameter.VpgHaGroupNotFound"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CCNROUTETABLENOTEXIST = "UnsupportedOperation.CcnRouteTableNotExist"
//  UNSUPPORTEDOPERATION_UNABLECROSSBORDER = "UnsupportedOperation.UnableCrossBorder"
func (c *Client) CreateDirectConnectGatewayWithContext(ctx context.Context, request *CreateDirectConnectGatewayRequest) (response *CreateDirectConnectGatewayResponse, err error) {
    if request == nil {
        request = NewCreateDirectConnectGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDirectConnectGateway require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateDirectConnectGatewayCcnRoutes(request *CreateDirectConnectGatewayCcnRoutesRequest) (response *CreateDirectConnectGatewayCcnRoutesResponse, err error) {
    return c.CreateDirectConnectGatewayCcnRoutesWithContext(context.Background(), request)
}

// CreateDirectConnectGatewayCcnRoutes
// This API (CreateDirectConnectGatewayCcnRoutes) is used to create the CCN route (IDC IP range) of a Direct Connect gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateDirectConnectGatewayCcnRoutesWithContext(ctx context.Context, request *CreateDirectConnectGatewayCcnRoutesRequest) (response *CreateDirectConnectGatewayCcnRoutesResponse, err error) {
    if request == nil {
        request = NewCreateDirectConnectGatewayCcnRoutesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDirectConnectGatewayCcnRoutes require credential")
    }

    request.SetContext(ctx)
    
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
//  INTERNALERROR_CREATECKAFKAROUTEERROR = "InternalError.CreateCkafkaRouteError"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TAGDUPLICATEKEY = "InvalidParameterValue.TagDuplicateKey"
//  INVALIDPARAMETERVALUE_TAGDUPLICATERESOURCETYPE = "InvalidParameterValue.TagDuplicateResourceType"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEY = "InvalidParameterValue.TagInvalidKey"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEYLEN = "InvalidParameterValue.TagInvalidKeyLen"
//  INVALIDPARAMETERVALUE_TAGINVALIDVAL = "InvalidParameterValue.TagInvalidVal"
//  INVALIDPARAMETERVALUE_TAGKEYNOTEXISTS = "InvalidParameterValue.TagKeyNotExists"
//  INVALIDPARAMETERVALUE_TAGNOTALLOCATEDQUOTA = "InvalidParameterValue.TagNotAllocatedQuota"
//  INVALIDPARAMETERVALUE_TAGNOTEXISTED = "InvalidParameterValue.TagNotExisted"
//  INVALIDPARAMETERVALUE_TAGNOTSUPPORTTAG = "InvalidParameterValue.TagNotSupportTag"
//  INVALIDPARAMETERVALUE_TAGRESOURCEFORMATERROR = "InvalidParameterValue.TagResourceFormatError"
//  INVALIDPARAMETERVALUE_TAGTIMESTAMPEXCEEDED = "InvalidParameterValue.TagTimestampExceeded"
//  INVALIDPARAMETERVALUE_TAGVALNOTEXISTS = "InvalidParameterValue.TagValNotExists"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TAGKEYEXCEEDED = "LimitExceeded.TagKeyExceeded"
//  LIMITEXCEEDED_TAGKEYPERRESOURCEEXCEEDED = "LimitExceeded.TagKeyPerResourceExceeded"
//  LIMITEXCEEDED_TAGNOTENOUGHQUOTA = "LimitExceeded.TagNotEnoughQuota"
//  LIMITEXCEEDED_TAGQUOTA = "LimitExceeded.TagQuota"
//  LIMITEXCEEDED_TAGQUOTAEXCEEDED = "LimitExceeded.TagQuotaExceeded"
//  LIMITEXCEEDED_TAGTAGSEXCEEDED = "LimitExceeded.TagTagsExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_FLOWLOGSNOTSUPPORTKOINSTANCEENI = "UnsupportedOperation.FlowLogsNotSupportKoInstanceEni"
//  UNSUPPORTEDOPERATION_FLOWLOGSNOTSUPPORTNULLINSTANCEENI = "UnsupportedOperation.FlowLogsNotSupportNullInstanceEni"
//  UNSUPPORTEDOPERATION_ONLYSUPPORTPROFESSIONKAFKA = "UnsupportedOperation.OnlySupportProfessionKafka"
//  UNSUPPORTEDOPERATION_TAGALLOCATE = "UnsupportedOperation.TagAllocate"
//  UNSUPPORTEDOPERATION_TAGFREE = "UnsupportedOperation.TagFree"
//  UNSUPPORTEDOPERATION_TAGNOTPERMIT = "UnsupportedOperation.TagNotPermit"
//  UNSUPPORTEDOPERATION_TAGSYSTEMRESERVEDTAGKEY = "UnsupportedOperation.TagSystemReservedTagKey"
func (c *Client) CreateFlowLog(request *CreateFlowLogRequest) (response *CreateFlowLogResponse, err error) {
    return c.CreateFlowLogWithContext(context.Background(), request)
}

// CreateFlowLog
// This API is used to create a flow log.
//
// error code that may be returned:
//  INTERNALERROR_CREATECKAFKAROUTEERROR = "InternalError.CreateCkafkaRouteError"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TAGDUPLICATEKEY = "InvalidParameterValue.TagDuplicateKey"
//  INVALIDPARAMETERVALUE_TAGDUPLICATERESOURCETYPE = "InvalidParameterValue.TagDuplicateResourceType"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEY = "InvalidParameterValue.TagInvalidKey"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEYLEN = "InvalidParameterValue.TagInvalidKeyLen"
//  INVALIDPARAMETERVALUE_TAGINVALIDVAL = "InvalidParameterValue.TagInvalidVal"
//  INVALIDPARAMETERVALUE_TAGKEYNOTEXISTS = "InvalidParameterValue.TagKeyNotExists"
//  INVALIDPARAMETERVALUE_TAGNOTALLOCATEDQUOTA = "InvalidParameterValue.TagNotAllocatedQuota"
//  INVALIDPARAMETERVALUE_TAGNOTEXISTED = "InvalidParameterValue.TagNotExisted"
//  INVALIDPARAMETERVALUE_TAGNOTSUPPORTTAG = "InvalidParameterValue.TagNotSupportTag"
//  INVALIDPARAMETERVALUE_TAGRESOURCEFORMATERROR = "InvalidParameterValue.TagResourceFormatError"
//  INVALIDPARAMETERVALUE_TAGTIMESTAMPEXCEEDED = "InvalidParameterValue.TagTimestampExceeded"
//  INVALIDPARAMETERVALUE_TAGVALNOTEXISTS = "InvalidParameterValue.TagValNotExists"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TAGKEYEXCEEDED = "LimitExceeded.TagKeyExceeded"
//  LIMITEXCEEDED_TAGKEYPERRESOURCEEXCEEDED = "LimitExceeded.TagKeyPerResourceExceeded"
//  LIMITEXCEEDED_TAGNOTENOUGHQUOTA = "LimitExceeded.TagNotEnoughQuota"
//  LIMITEXCEEDED_TAGQUOTA = "LimitExceeded.TagQuota"
//  LIMITEXCEEDED_TAGQUOTAEXCEEDED = "LimitExceeded.TagQuotaExceeded"
//  LIMITEXCEEDED_TAGTAGSEXCEEDED = "LimitExceeded.TagTagsExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_FLOWLOGSNOTSUPPORTKOINSTANCEENI = "UnsupportedOperation.FlowLogsNotSupportKoInstanceEni"
//  UNSUPPORTEDOPERATION_FLOWLOGSNOTSUPPORTNULLINSTANCEENI = "UnsupportedOperation.FlowLogsNotSupportNullInstanceEni"
//  UNSUPPORTEDOPERATION_ONLYSUPPORTPROFESSIONKAFKA = "UnsupportedOperation.OnlySupportProfessionKafka"
//  UNSUPPORTEDOPERATION_TAGALLOCATE = "UnsupportedOperation.TagAllocate"
//  UNSUPPORTEDOPERATION_TAGFREE = "UnsupportedOperation.TagFree"
//  UNSUPPORTEDOPERATION_TAGNOTPERMIT = "UnsupportedOperation.TagNotPermit"
//  UNSUPPORTEDOPERATION_TAGSYSTEMRESERVEDTAGKEY = "UnsupportedOperation.TagSystemReservedTagKey"
func (c *Client) CreateFlowLogWithContext(ctx context.Context, request *CreateFlowLogRequest) (response *CreateFlowLogResponse, err error) {
    if request == nil {
        request = NewCreateFlowLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFlowLog require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_INVALIDBUSINESS = "InvalidParameterValue.InvalidBusiness"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateHaVip(request *CreateHaVipRequest) (response *CreateHaVipResponse, err error) {
    return c.CreateHaVipWithContext(context.Background(), request)
}

// CreateHaVip
// This API (CreateHaVip) is used to create a highly available virtual IP (HAVIP)
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_INVALIDBUSINESS = "InvalidParameterValue.InvalidBusiness"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateHaVipWithContext(ctx context.Context, request *CreateHaVipRequest) (response *CreateHaVipResponse, err error) {
    if request == nil {
        request = NewCreateHaVipRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateHaVip require credential")
    }

    request.SetContext(ctx)
    
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
//  UNSUPPORTEDOPERATION_LOCALGATEWAYALREADYEXISTS = "UnsupportedOperation.LocalGatewayAlreadyExists"
func (c *Client) CreateLocalGateway(request *CreateLocalGatewayRequest) (response *CreateLocalGatewayResponse, err error) {
    return c.CreateLocalGatewayWithContext(context.Background(), request)
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
//  UNSUPPORTEDOPERATION_LOCALGATEWAYALREADYEXISTS = "UnsupportedOperation.LocalGatewayAlreadyExists"
func (c *Client) CreateLocalGatewayWithContext(ctx context.Context, request *CreateLocalGatewayRequest) (response *CreateLocalGatewayResponse, err error) {
    if request == nil {
        request = NewCreateLocalGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLocalGateway require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to create a NAT Gateway.
//
// Before taking actions on a NAT gateway, ensure that it has been successfully created, namely, the `State` field in the response of the `DescribeNatGateway` API is `AVAILABLE`.
//
// error code that may be returned:
//  ADDRESSQUOTALIMITEXCEEDED = "AddressQuotaLimitExceeded"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDACCOUNT_NOTSUPPORTED = "InvalidAccount.NotSupported"
//  INVALIDADDRESSSTATE = "InvalidAddressState"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_EIPBRANDWIDTHOUTINVALID = "InvalidParameterValue.EIPBrandWidthOutInvalid"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TAGDUPLICATEKEY = "InvalidParameterValue.TagDuplicateKey"
//  INVALIDPARAMETERVALUE_TAGDUPLICATERESOURCETYPE = "InvalidParameterValue.TagDuplicateResourceType"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEY = "InvalidParameterValue.TagInvalidKey"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEYLEN = "InvalidParameterValue.TagInvalidKeyLen"
//  INVALIDPARAMETERVALUE_TAGINVALIDVAL = "InvalidParameterValue.TagInvalidVal"
//  INVALIDPARAMETERVALUE_TAGKEYNOTEXISTS = "InvalidParameterValue.TagKeyNotExists"
//  INVALIDPARAMETERVALUE_TAGNOTALLOCATEDQUOTA = "InvalidParameterValue.TagNotAllocatedQuota"
//  INVALIDPARAMETERVALUE_TAGNOTEXISTED = "InvalidParameterValue.TagNotExisted"
//  INVALIDPARAMETERVALUE_TAGNOTSUPPORTTAG = "InvalidParameterValue.TagNotSupportTag"
//  INVALIDPARAMETERVALUE_TAGRESOURCEFORMATERROR = "InvalidParameterValue.TagResourceFormatError"
//  INVALIDPARAMETERVALUE_TAGTIMESTAMPEXCEEDED = "InvalidParameterValue.TagTimestampExceeded"
//  INVALIDPARAMETERVALUE_TAGVALNOTEXISTS = "InvalidParameterValue.TagValNotExists"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDVPCID_MALFORMED = "InvalidVpcId.Malformed"
//  INVALIDVPCID_NOTFOUND = "InvalidVpcId.NotFound"
//  LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.AddressQuotaLimitExceeded"
//  LIMITEXCEEDED_DAILYALLOCATEADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.DailyAllocateAddressQuotaLimitExceeded"
//  LIMITEXCEEDED_NATGATEWAYLIMITEXCEEDED = "LimitExceeded.NatGatewayLimitExceeded"
//  LIMITEXCEEDED_NATGATEWAYPERVPCLIMITEXCEEDED = "LimitExceeded.NatGatewayPerVpcLimitExceeded"
//  LIMITEXCEEDED_PUBLICIPADDRESSPERNATGATEWAYLIMITEXCEEDED = "LimitExceeded.PublicIpAddressPerNatGatewayLimitExceeded"
//  LIMITEXCEEDED_TAGKEYEXCEEDED = "LimitExceeded.TagKeyExceeded"
//  LIMITEXCEEDED_TAGKEYPERRESOURCEEXCEEDED = "LimitExceeded.TagKeyPerResourceExceeded"
//  LIMITEXCEEDED_TAGNOTENOUGHQUOTA = "LimitExceeded.TagNotEnoughQuota"
//  LIMITEXCEEDED_TAGQUOTA = "LimitExceeded.TagQuota"
//  LIMITEXCEEDED_TAGQUOTAEXCEEDED = "LimitExceeded.TagQuotaExceeded"
//  LIMITEXCEEDED_TAGTAGSEXCEEDED = "LimitExceeded.TagTagsExceeded"
//  RESOURCEINUSE_ADDRESS = "ResourceInUse.Address"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOREALNAMEAUTHENTICATION = "UnauthorizedOperation.NoRealNameAuthentication"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INSUFFICIENTFUNDS = "UnsupportedOperation.InsufficientFunds"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_PUBLICIPADDRESSISNOTBGPIP = "UnsupportedOperation.PublicIpAddressIsNotBGPIp"
//  UNSUPPORTEDOPERATION_PUBLICIPADDRESSISNOTEXISTED = "UnsupportedOperation.PublicIpAddressIsNotExisted"
//  UNSUPPORTEDOPERATION_PUBLICIPADDRESSNOTBILLEDBYTRAFFIC = "UnsupportedOperation.PublicIpAddressNotBilledByTraffic"
//  UNSUPPORTEDOPERATION_TAGALLOCATE = "UnsupportedOperation.TagAllocate"
//  UNSUPPORTEDOPERATION_TAGFREE = "UnsupportedOperation.TagFree"
//  UNSUPPORTEDOPERATION_TAGNOTPERMIT = "UnsupportedOperation.TagNotPermit"
//  UNSUPPORTEDOPERATION_TAGSYSTEMRESERVEDTAGKEY = "UnsupportedOperation.TagSystemReservedTagKey"
func (c *Client) CreateNatGateway(request *CreateNatGatewayRequest) (response *CreateNatGatewayResponse, err error) {
    return c.CreateNatGatewayWithContext(context.Background(), request)
}

// CreateNatGateway
// This API is used to create a NAT Gateway.
//
// Before taking actions on a NAT gateway, ensure that it has been successfully created, namely, the `State` field in the response of the `DescribeNatGateway` API is `AVAILABLE`.
//
// error code that may be returned:
//  ADDRESSQUOTALIMITEXCEEDED = "AddressQuotaLimitExceeded"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDACCOUNT_NOTSUPPORTED = "InvalidAccount.NotSupported"
//  INVALIDADDRESSSTATE = "InvalidAddressState"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_EIPBRANDWIDTHOUTINVALID = "InvalidParameterValue.EIPBrandWidthOutInvalid"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TAGDUPLICATEKEY = "InvalidParameterValue.TagDuplicateKey"
//  INVALIDPARAMETERVALUE_TAGDUPLICATERESOURCETYPE = "InvalidParameterValue.TagDuplicateResourceType"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEY = "InvalidParameterValue.TagInvalidKey"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEYLEN = "InvalidParameterValue.TagInvalidKeyLen"
//  INVALIDPARAMETERVALUE_TAGINVALIDVAL = "InvalidParameterValue.TagInvalidVal"
//  INVALIDPARAMETERVALUE_TAGKEYNOTEXISTS = "InvalidParameterValue.TagKeyNotExists"
//  INVALIDPARAMETERVALUE_TAGNOTALLOCATEDQUOTA = "InvalidParameterValue.TagNotAllocatedQuota"
//  INVALIDPARAMETERVALUE_TAGNOTEXISTED = "InvalidParameterValue.TagNotExisted"
//  INVALIDPARAMETERVALUE_TAGNOTSUPPORTTAG = "InvalidParameterValue.TagNotSupportTag"
//  INVALIDPARAMETERVALUE_TAGRESOURCEFORMATERROR = "InvalidParameterValue.TagResourceFormatError"
//  INVALIDPARAMETERVALUE_TAGTIMESTAMPEXCEEDED = "InvalidParameterValue.TagTimestampExceeded"
//  INVALIDPARAMETERVALUE_TAGVALNOTEXISTS = "InvalidParameterValue.TagValNotExists"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDVPCID_MALFORMED = "InvalidVpcId.Malformed"
//  INVALIDVPCID_NOTFOUND = "InvalidVpcId.NotFound"
//  LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.AddressQuotaLimitExceeded"
//  LIMITEXCEEDED_DAILYALLOCATEADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.DailyAllocateAddressQuotaLimitExceeded"
//  LIMITEXCEEDED_NATGATEWAYLIMITEXCEEDED = "LimitExceeded.NatGatewayLimitExceeded"
//  LIMITEXCEEDED_NATGATEWAYPERVPCLIMITEXCEEDED = "LimitExceeded.NatGatewayPerVpcLimitExceeded"
//  LIMITEXCEEDED_PUBLICIPADDRESSPERNATGATEWAYLIMITEXCEEDED = "LimitExceeded.PublicIpAddressPerNatGatewayLimitExceeded"
//  LIMITEXCEEDED_TAGKEYEXCEEDED = "LimitExceeded.TagKeyExceeded"
//  LIMITEXCEEDED_TAGKEYPERRESOURCEEXCEEDED = "LimitExceeded.TagKeyPerResourceExceeded"
//  LIMITEXCEEDED_TAGNOTENOUGHQUOTA = "LimitExceeded.TagNotEnoughQuota"
//  LIMITEXCEEDED_TAGQUOTA = "LimitExceeded.TagQuota"
//  LIMITEXCEEDED_TAGQUOTAEXCEEDED = "LimitExceeded.TagQuotaExceeded"
//  LIMITEXCEEDED_TAGTAGSEXCEEDED = "LimitExceeded.TagTagsExceeded"
//  RESOURCEINUSE_ADDRESS = "ResourceInUse.Address"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOREALNAMEAUTHENTICATION = "UnauthorizedOperation.NoRealNameAuthentication"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INSUFFICIENTFUNDS = "UnsupportedOperation.InsufficientFunds"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_PUBLICIPADDRESSISNOTBGPIP = "UnsupportedOperation.PublicIpAddressIsNotBGPIp"
//  UNSUPPORTEDOPERATION_PUBLICIPADDRESSISNOTEXISTED = "UnsupportedOperation.PublicIpAddressIsNotExisted"
//  UNSUPPORTEDOPERATION_PUBLICIPADDRESSNOTBILLEDBYTRAFFIC = "UnsupportedOperation.PublicIpAddressNotBilledByTraffic"
//  UNSUPPORTEDOPERATION_TAGALLOCATE = "UnsupportedOperation.TagAllocate"
//  UNSUPPORTEDOPERATION_TAGFREE = "UnsupportedOperation.TagFree"
//  UNSUPPORTEDOPERATION_TAGNOTPERMIT = "UnsupportedOperation.TagNotPermit"
//  UNSUPPORTEDOPERATION_TAGSYSTEMRESERVEDTAGKEY = "UnsupportedOperation.TagSystemReservedTagKey"
func (c *Client) CreateNatGatewayWithContext(ctx context.Context, request *CreateNatGatewayRequest) (response *CreateNatGatewayResponse, err error) {
    if request == nil {
        request = NewCreateNatGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNatGateway require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to create the port forwarding rules of a NAT gateway.
//
// error code that may be returned:
//  ADDRESSQUOTALIMITEXCEEDED = "AddressQuotaLimitExceeded"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDACCOUNT_NOTSUPPORTED = "InvalidAccount.NotSupported"
//  INVALIDADDRESSSTATE = "InvalidAddressState"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_EIPBRANDWIDTHOUTINVALID = "InvalidParameterValue.EIPBrandWidthOutInvalid"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TAGDUPLICATEKEY = "InvalidParameterValue.TagDuplicateKey"
//  INVALIDPARAMETERVALUE_TAGDUPLICATERESOURCETYPE = "InvalidParameterValue.TagDuplicateResourceType"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEY = "InvalidParameterValue.TagInvalidKey"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEYLEN = "InvalidParameterValue.TagInvalidKeyLen"
//  INVALIDPARAMETERVALUE_TAGINVALIDVAL = "InvalidParameterValue.TagInvalidVal"
//  INVALIDPARAMETERVALUE_TAGKEYNOTEXISTS = "InvalidParameterValue.TagKeyNotExists"
//  INVALIDPARAMETERVALUE_TAGNOTALLOCATEDQUOTA = "InvalidParameterValue.TagNotAllocatedQuota"
//  INVALIDPARAMETERVALUE_TAGNOTEXISTED = "InvalidParameterValue.TagNotExisted"
//  INVALIDPARAMETERVALUE_TAGNOTSUPPORTTAG = "InvalidParameterValue.TagNotSupportTag"
//  INVALIDPARAMETERVALUE_TAGRESOURCEFORMATERROR = "InvalidParameterValue.TagResourceFormatError"
//  INVALIDPARAMETERVALUE_TAGTIMESTAMPEXCEEDED = "InvalidParameterValue.TagTimestampExceeded"
//  INVALIDPARAMETERVALUE_TAGVALNOTEXISTS = "InvalidParameterValue.TagValNotExists"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDVPCID_MALFORMED = "InvalidVpcId.Malformed"
//  INVALIDVPCID_NOTFOUND = "InvalidVpcId.NotFound"
//  LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.AddressQuotaLimitExceeded"
//  LIMITEXCEEDED_DAILYALLOCATEADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.DailyAllocateAddressQuotaLimitExceeded"
//  LIMITEXCEEDED_NATGATEWAYLIMITEXCEEDED = "LimitExceeded.NatGatewayLimitExceeded"
//  LIMITEXCEEDED_NATGATEWAYPERVPCLIMITEXCEEDED = "LimitExceeded.NatGatewayPerVpcLimitExceeded"
//  LIMITEXCEEDED_PUBLICIPADDRESSPERNATGATEWAYLIMITEXCEEDED = "LimitExceeded.PublicIpAddressPerNatGatewayLimitExceeded"
//  LIMITEXCEEDED_TAGKEYEXCEEDED = "LimitExceeded.TagKeyExceeded"
//  LIMITEXCEEDED_TAGKEYPERRESOURCEEXCEEDED = "LimitExceeded.TagKeyPerResourceExceeded"
//  LIMITEXCEEDED_TAGNOTENOUGHQUOTA = "LimitExceeded.TagNotEnoughQuota"
//  LIMITEXCEEDED_TAGQUOTA = "LimitExceeded.TagQuota"
//  LIMITEXCEEDED_TAGQUOTAEXCEEDED = "LimitExceeded.TagQuotaExceeded"
//  LIMITEXCEEDED_TAGTAGSEXCEEDED = "LimitExceeded.TagTagsExceeded"
//  RESOURCEINUSE_ADDRESS = "ResourceInUse.Address"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOREALNAMEAUTHENTICATION = "UnauthorizedOperation.NoRealNameAuthentication"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INSUFFICIENTFUNDS = "UnsupportedOperation.InsufficientFunds"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_PUBLICIPADDRESSISNOTBGPIP = "UnsupportedOperation.PublicIpAddressIsNotBGPIp"
//  UNSUPPORTEDOPERATION_PUBLICIPADDRESSISNOTEXISTED = "UnsupportedOperation.PublicIpAddressIsNotExisted"
//  UNSUPPORTEDOPERATION_PUBLICIPADDRESSNOTBILLEDBYTRAFFIC = "UnsupportedOperation.PublicIpAddressNotBilledByTraffic"
//  UNSUPPORTEDOPERATION_TAGALLOCATE = "UnsupportedOperation.TagAllocate"
//  UNSUPPORTEDOPERATION_TAGFREE = "UnsupportedOperation.TagFree"
//  UNSUPPORTEDOPERATION_TAGNOTPERMIT = "UnsupportedOperation.TagNotPermit"
//  UNSUPPORTEDOPERATION_TAGSYSTEMRESERVEDTAGKEY = "UnsupportedOperation.TagSystemReservedTagKey"
func (c *Client) CreateNatGatewayDestinationIpPortTranslationNatRule(request *CreateNatGatewayDestinationIpPortTranslationNatRuleRequest) (response *CreateNatGatewayDestinationIpPortTranslationNatRuleResponse, err error) {
    return c.CreateNatGatewayDestinationIpPortTranslationNatRuleWithContext(context.Background(), request)
}

// CreateNatGatewayDestinationIpPortTranslationNatRule
// This API is used to create the port forwarding rules of a NAT gateway.
//
// error code that may be returned:
//  ADDRESSQUOTALIMITEXCEEDED = "AddressQuotaLimitExceeded"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDACCOUNT_NOTSUPPORTED = "InvalidAccount.NotSupported"
//  INVALIDADDRESSSTATE = "InvalidAddressState"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_EIPBRANDWIDTHOUTINVALID = "InvalidParameterValue.EIPBrandWidthOutInvalid"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TAGDUPLICATEKEY = "InvalidParameterValue.TagDuplicateKey"
//  INVALIDPARAMETERVALUE_TAGDUPLICATERESOURCETYPE = "InvalidParameterValue.TagDuplicateResourceType"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEY = "InvalidParameterValue.TagInvalidKey"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEYLEN = "InvalidParameterValue.TagInvalidKeyLen"
//  INVALIDPARAMETERVALUE_TAGINVALIDVAL = "InvalidParameterValue.TagInvalidVal"
//  INVALIDPARAMETERVALUE_TAGKEYNOTEXISTS = "InvalidParameterValue.TagKeyNotExists"
//  INVALIDPARAMETERVALUE_TAGNOTALLOCATEDQUOTA = "InvalidParameterValue.TagNotAllocatedQuota"
//  INVALIDPARAMETERVALUE_TAGNOTEXISTED = "InvalidParameterValue.TagNotExisted"
//  INVALIDPARAMETERVALUE_TAGNOTSUPPORTTAG = "InvalidParameterValue.TagNotSupportTag"
//  INVALIDPARAMETERVALUE_TAGRESOURCEFORMATERROR = "InvalidParameterValue.TagResourceFormatError"
//  INVALIDPARAMETERVALUE_TAGTIMESTAMPEXCEEDED = "InvalidParameterValue.TagTimestampExceeded"
//  INVALIDPARAMETERVALUE_TAGVALNOTEXISTS = "InvalidParameterValue.TagValNotExists"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDVPCID_MALFORMED = "InvalidVpcId.Malformed"
//  INVALIDVPCID_NOTFOUND = "InvalidVpcId.NotFound"
//  LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.AddressQuotaLimitExceeded"
//  LIMITEXCEEDED_DAILYALLOCATEADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.DailyAllocateAddressQuotaLimitExceeded"
//  LIMITEXCEEDED_NATGATEWAYLIMITEXCEEDED = "LimitExceeded.NatGatewayLimitExceeded"
//  LIMITEXCEEDED_NATGATEWAYPERVPCLIMITEXCEEDED = "LimitExceeded.NatGatewayPerVpcLimitExceeded"
//  LIMITEXCEEDED_PUBLICIPADDRESSPERNATGATEWAYLIMITEXCEEDED = "LimitExceeded.PublicIpAddressPerNatGatewayLimitExceeded"
//  LIMITEXCEEDED_TAGKEYEXCEEDED = "LimitExceeded.TagKeyExceeded"
//  LIMITEXCEEDED_TAGKEYPERRESOURCEEXCEEDED = "LimitExceeded.TagKeyPerResourceExceeded"
//  LIMITEXCEEDED_TAGNOTENOUGHQUOTA = "LimitExceeded.TagNotEnoughQuota"
//  LIMITEXCEEDED_TAGQUOTA = "LimitExceeded.TagQuota"
//  LIMITEXCEEDED_TAGQUOTAEXCEEDED = "LimitExceeded.TagQuotaExceeded"
//  LIMITEXCEEDED_TAGTAGSEXCEEDED = "LimitExceeded.TagTagsExceeded"
//  RESOURCEINUSE_ADDRESS = "ResourceInUse.Address"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOREALNAMEAUTHENTICATION = "UnauthorizedOperation.NoRealNameAuthentication"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INSUFFICIENTFUNDS = "UnsupportedOperation.InsufficientFunds"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_PUBLICIPADDRESSISNOTBGPIP = "UnsupportedOperation.PublicIpAddressIsNotBGPIp"
//  UNSUPPORTEDOPERATION_PUBLICIPADDRESSISNOTEXISTED = "UnsupportedOperation.PublicIpAddressIsNotExisted"
//  UNSUPPORTEDOPERATION_PUBLICIPADDRESSNOTBILLEDBYTRAFFIC = "UnsupportedOperation.PublicIpAddressNotBilledByTraffic"
//  UNSUPPORTEDOPERATION_TAGALLOCATE = "UnsupportedOperation.TagAllocate"
//  UNSUPPORTEDOPERATION_TAGFREE = "UnsupportedOperation.TagFree"
//  UNSUPPORTEDOPERATION_TAGNOTPERMIT = "UnsupportedOperation.TagNotPermit"
//  UNSUPPORTEDOPERATION_TAGSYSTEMRESERVEDTAGKEY = "UnsupportedOperation.TagSystemReservedTagKey"
func (c *Client) CreateNatGatewayDestinationIpPortTranslationNatRuleWithContext(ctx context.Context, request *CreateNatGatewayDestinationIpPortTranslationNatRuleRequest) (response *CreateNatGatewayDestinationIpPortTranslationNatRuleResponse, err error) {
    if request == nil {
        request = NewCreateNatGatewayDestinationIpPortTranslationNatRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNatGatewayDestinationIpPortTranslationNatRule require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to create SNAT rules for a NAT gateway.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NATSNATRULEEXISTS = "InvalidParameterValue.NatSnatRuleExists"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION_NATGATEWAYRULEPIPEXISTS = "UnsupportedOperation.NatGatewayRulePipExists"
//  UNSUPPORTEDOPERATION_NATGATEWAYTYPENOTSUPPORTSNAT = "UnsupportedOperation.NatGatewayTypeNotSupportSNAT"
//  UNSUPPORTEDOPERATION_UNBINDEIP = "UnsupportedOperation.UnbindEIP"
func (c *Client) CreateNatGatewaySourceIpTranslationNatRule(request *CreateNatGatewaySourceIpTranslationNatRuleRequest) (response *CreateNatGatewaySourceIpTranslationNatRuleResponse, err error) {
    return c.CreateNatGatewaySourceIpTranslationNatRuleWithContext(context.Background(), request)
}

// CreateNatGatewaySourceIpTranslationNatRule
// This API is used to create SNAT rules for a NAT gateway.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NATSNATRULEEXISTS = "InvalidParameterValue.NatSnatRuleExists"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION_NATGATEWAYRULEPIPEXISTS = "UnsupportedOperation.NatGatewayRulePipExists"
//  UNSUPPORTEDOPERATION_NATGATEWAYTYPENOTSUPPORTSNAT = "UnsupportedOperation.NatGatewayTypeNotSupportSNAT"
//  UNSUPPORTEDOPERATION_UNBINDEIP = "UnsupportedOperation.UnbindEIP"
func (c *Client) CreateNatGatewaySourceIpTranslationNatRuleWithContext(ctx context.Context, request *CreateNatGatewaySourceIpTranslationNatRuleRequest) (response *CreateNatGatewaySourceIpTranslationNatRuleResponse, err error) {
    if request == nil {
        request = NewCreateNatGatewaySourceIpTranslationNatRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNatGatewaySourceIpTranslationNatRule require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION_NETDETECTTIMEOUT = "FailedOperation.NetDetectTimeOut"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_NEXTHOPMISMATCH = "InvalidParameter.NextHopMismatch"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NETDETECTINVPC = "InvalidParameterValue.NetDetectInVpc"
//  INVALIDPARAMETERVALUE_NETDETECTNOTFOUNDIP = "InvalidParameterValue.NetDetectNotFoundIp"
//  INVALIDPARAMETERVALUE_NETDETECTSAMEIP = "InvalidParameterValue.NetDetectSameIp"
//  INVALIDPARAMETERVALUE_RESOURCENOTFOUND = "InvalidParameterValue.ResourceNotFound"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_CONFLICTWITHDOCKERROUTE = "UnsupportedOperation.ConflictWithDockerRoute"
//  UNSUPPORTEDOPERATION_ECMPWITHUSERROUTE = "UnsupportedOperation.EcmpWithUserRoute"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
func (c *Client) CreateNetDetect(request *CreateNetDetectRequest) (response *CreateNetDetectResponse, err error) {
    return c.CreateNetDetectWithContext(context.Background(), request)
}

// CreateNetDetect
// This API is used to create a network detection instance.
//
// error code that may be returned:
//  FAILEDOPERATION_NETDETECTTIMEOUT = "FailedOperation.NetDetectTimeOut"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_NEXTHOPMISMATCH = "InvalidParameter.NextHopMismatch"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NETDETECTINVPC = "InvalidParameterValue.NetDetectInVpc"
//  INVALIDPARAMETERVALUE_NETDETECTNOTFOUNDIP = "InvalidParameterValue.NetDetectNotFoundIp"
//  INVALIDPARAMETERVALUE_NETDETECTSAMEIP = "InvalidParameterValue.NetDetectSameIp"
//  INVALIDPARAMETERVALUE_RESOURCENOTFOUND = "InvalidParameterValue.ResourceNotFound"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_CONFLICTWITHDOCKERROUTE = "UnsupportedOperation.ConflictWithDockerRoute"
//  UNSUPPORTEDOPERATION_ECMPWITHUSERROUTE = "UnsupportedOperation.EcmpWithUserRoute"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
func (c *Client) CreateNetDetectWithContext(ctx context.Context, request *CreateNetDetectRequest) (response *CreateNetDetectResponse, err error) {
    if request == nil {
        request = NewCreateNetDetectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNetDetect require credential")
    }

    request.SetContext(ctx)
    
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
// * The inbound and outbound rules for a new network ACL are "Deny All" by default. You need to call `ModifyNetworkAclEntries` to set rules for the new network ACL as needed.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TAGDUPLICATEKEY = "InvalidParameterValue.TagDuplicateKey"
//  INVALIDPARAMETERVALUE_TAGDUPLICATERESOURCETYPE = "InvalidParameterValue.TagDuplicateResourceType"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEY = "InvalidParameterValue.TagInvalidKey"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEYLEN = "InvalidParameterValue.TagInvalidKeyLen"
//  INVALIDPARAMETERVALUE_TAGINVALIDVAL = "InvalidParameterValue.TagInvalidVal"
//  INVALIDPARAMETERVALUE_TAGKEYNOTEXISTS = "InvalidParameterValue.TagKeyNotExists"
//  INVALIDPARAMETERVALUE_TAGNOTALLOCATEDQUOTA = "InvalidParameterValue.TagNotAllocatedQuota"
//  INVALIDPARAMETERVALUE_TAGNOTEXISTED = "InvalidParameterValue.TagNotExisted"
//  INVALIDPARAMETERVALUE_TAGNOTSUPPORTTAG = "InvalidParameterValue.TagNotSupportTag"
//  INVALIDPARAMETERVALUE_TAGRESOURCEFORMATERROR = "InvalidParameterValue.TagResourceFormatError"
//  INVALIDPARAMETERVALUE_TAGTIMESTAMPEXCEEDED = "InvalidParameterValue.TagTimestampExceeded"
//  INVALIDPARAMETERVALUE_TAGVALNOTEXISTS = "InvalidParameterValue.TagValNotExists"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TAGKEYEXCEEDED = "LimitExceeded.TagKeyExceeded"
//  LIMITEXCEEDED_TAGKEYPERRESOURCEEXCEEDED = "LimitExceeded.TagKeyPerResourceExceeded"
//  LIMITEXCEEDED_TAGNOTENOUGHQUOTA = "LimitExceeded.TagNotEnoughQuota"
//  LIMITEXCEEDED_TAGQUOTA = "LimitExceeded.TagQuota"
//  LIMITEXCEEDED_TAGQUOTAEXCEEDED = "LimitExceeded.TagQuotaExceeded"
//  LIMITEXCEEDED_TAGTAGSEXCEEDED = "LimitExceeded.TagTagsExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER_WITHGUESS = "UnknownParameter.WithGuess"
//  UNSUPPORTEDOPERATION_APPIDMISMATCH = "UnsupportedOperation.AppIdMismatch"
//  UNSUPPORTEDOPERATION_TAGALLOCATE = "UnsupportedOperation.TagAllocate"
//  UNSUPPORTEDOPERATION_TAGFREE = "UnsupportedOperation.TagFree"
//  UNSUPPORTEDOPERATION_TAGNOTPERMIT = "UnsupportedOperation.TagNotPermit"
//  UNSUPPORTEDOPERATION_TAGSYSTEMRESERVEDTAGKEY = "UnsupportedOperation.TagSystemReservedTagKey"
func (c *Client) CreateNetworkAcl(request *CreateNetworkAclRequest) (response *CreateNetworkAclResponse, err error) {
    return c.CreateNetworkAclWithContext(context.Background(), request)
}

// CreateNetworkAcl
// This API is used to create a <a href="https://intl.cloud.tencent.com/document/product/215/20088?from_cn_redirect=1">network ACL</a>.
//
// * The inbound and outbound rules for a new network ACL are "Deny All" by default. You need to call `ModifyNetworkAclEntries` to set rules for the new network ACL as needed.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TAGDUPLICATEKEY = "InvalidParameterValue.TagDuplicateKey"
//  INVALIDPARAMETERVALUE_TAGDUPLICATERESOURCETYPE = "InvalidParameterValue.TagDuplicateResourceType"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEY = "InvalidParameterValue.TagInvalidKey"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEYLEN = "InvalidParameterValue.TagInvalidKeyLen"
//  INVALIDPARAMETERVALUE_TAGINVALIDVAL = "InvalidParameterValue.TagInvalidVal"
//  INVALIDPARAMETERVALUE_TAGKEYNOTEXISTS = "InvalidParameterValue.TagKeyNotExists"
//  INVALIDPARAMETERVALUE_TAGNOTALLOCATEDQUOTA = "InvalidParameterValue.TagNotAllocatedQuota"
//  INVALIDPARAMETERVALUE_TAGNOTEXISTED = "InvalidParameterValue.TagNotExisted"
//  INVALIDPARAMETERVALUE_TAGNOTSUPPORTTAG = "InvalidParameterValue.TagNotSupportTag"
//  INVALIDPARAMETERVALUE_TAGRESOURCEFORMATERROR = "InvalidParameterValue.TagResourceFormatError"
//  INVALIDPARAMETERVALUE_TAGTIMESTAMPEXCEEDED = "InvalidParameterValue.TagTimestampExceeded"
//  INVALIDPARAMETERVALUE_TAGVALNOTEXISTS = "InvalidParameterValue.TagValNotExists"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TAGKEYEXCEEDED = "LimitExceeded.TagKeyExceeded"
//  LIMITEXCEEDED_TAGKEYPERRESOURCEEXCEEDED = "LimitExceeded.TagKeyPerResourceExceeded"
//  LIMITEXCEEDED_TAGNOTENOUGHQUOTA = "LimitExceeded.TagNotEnoughQuota"
//  LIMITEXCEEDED_TAGQUOTA = "LimitExceeded.TagQuota"
//  LIMITEXCEEDED_TAGQUOTAEXCEEDED = "LimitExceeded.TagQuotaExceeded"
//  LIMITEXCEEDED_TAGTAGSEXCEEDED = "LimitExceeded.TagTagsExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER_WITHGUESS = "UnknownParameter.WithGuess"
//  UNSUPPORTEDOPERATION_APPIDMISMATCH = "UnsupportedOperation.AppIdMismatch"
//  UNSUPPORTEDOPERATION_TAGALLOCATE = "UnsupportedOperation.TagAllocate"
//  UNSUPPORTEDOPERATION_TAGFREE = "UnsupportedOperation.TagFree"
//  UNSUPPORTEDOPERATION_TAGNOTPERMIT = "UnsupportedOperation.TagNotPermit"
//  UNSUPPORTEDOPERATION_TAGSYSTEMRESERVEDTAGKEY = "UnsupportedOperation.TagSystemReservedTagKey"
func (c *Client) CreateNetworkAclWithContext(ctx context.Context, request *CreateNetworkAclRequest) (response *CreateNetworkAclResponse, err error) {
    if request == nil {
        request = NewCreateNetworkAclRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNetworkAcl require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to create an ENI.
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
// >?This API is async. You can call the [`DescribeVpcTaskResult`](https://intl.cloud.tencent.com/document/api/215/59037?from_cn_redirect=1) API to query the task result. When the task is completed, you can continue other tasks.
//
// >
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_RESERVED = "InvalidParameterValue.Reserved"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TAGNOTENOUGHQUOTA = "LimitExceeded.TagNotEnoughQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_RESOURCEMISMATCH = "UnsupportedOperation.ResourceMismatch"
func (c *Client) CreateNetworkInterface(request *CreateNetworkInterfaceRequest) (response *CreateNetworkInterfaceResponse, err error) {
    return c.CreateNetworkInterfaceWithContext(context.Background(), request)
}

// CreateNetworkInterface
// This API is used to create an ENI.
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
// >?This API is async. You can call the [`DescribeVpcTaskResult`](https://intl.cloud.tencent.com/document/api/215/59037?from_cn_redirect=1) API to query the task result. When the task is completed, you can continue other tasks.
//
// >
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_RESERVED = "InvalidParameterValue.Reserved"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TAGNOTENOUGHQUOTA = "LimitExceeded.TagNotEnoughQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_RESOURCEMISMATCH = "UnsupportedOperation.ResourceMismatch"
func (c *Client) CreateNetworkInterfaceWithContext(ctx context.Context, request *CreateNetworkInterfaceRequest) (response *CreateNetworkInterfaceResponse, err error) {
    if request == nil {
        request = NewCreateNetworkInterfaceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNetworkInterface require credential")
    }

    request.SetContext(ctx)
    
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
//  LIMITEXCEEDED_TAGNOTENOUGHQUOTA = "LimitExceeded.TagNotEnoughQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateRouteTable(request *CreateRouteTableRequest) (response *CreateRouteTableResponse, err error) {
    return c.CreateRouteTableWithContext(context.Background(), request)
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
//  LIMITEXCEEDED_TAGNOTENOUGHQUOTA = "LimitExceeded.TagNotEnoughQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateRouteTableWithContext(ctx context.Context, request *CreateRouteTableRequest) (response *CreateRouteTableResponse, err error) {
    if request == nil {
        request = NewCreateRouteTableRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRouteTable require credential")
    }

    request.SetContext(ctx)
    
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
    return c.CreateRoutesWithContext(context.Background(), request)
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
func (c *Client) CreateRoutesWithContext(ctx context.Context, request *CreateRoutesRequest) (response *CreateRoutesResponse, err error) {
    if request == nil {
        request = NewCreateRoutesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRoutes require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TAGDUPLICATEKEY = "InvalidParameterValue.TagDuplicateKey"
//  INVALIDPARAMETERVALUE_TAGDUPLICATERESOURCETYPE = "InvalidParameterValue.TagDuplicateResourceType"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEY = "InvalidParameterValue.TagInvalidKey"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEYLEN = "InvalidParameterValue.TagInvalidKeyLen"
//  INVALIDPARAMETERVALUE_TAGINVALIDVAL = "InvalidParameterValue.TagInvalidVal"
//  INVALIDPARAMETERVALUE_TAGKEYNOTEXISTS = "InvalidParameterValue.TagKeyNotExists"
//  INVALIDPARAMETERVALUE_TAGNOTALLOCATEDQUOTA = "InvalidParameterValue.TagNotAllocatedQuota"
//  INVALIDPARAMETERVALUE_TAGNOTEXISTED = "InvalidParameterValue.TagNotExisted"
//  INVALIDPARAMETERVALUE_TAGNOTSUPPORTTAG = "InvalidParameterValue.TagNotSupportTag"
//  INVALIDPARAMETERVALUE_TAGRESOURCEFORMATERROR = "InvalidParameterValue.TagResourceFormatError"
//  INVALIDPARAMETERVALUE_TAGTIMESTAMPEXCEEDED = "InvalidParameterValue.TagTimestampExceeded"
//  INVALIDPARAMETERVALUE_TAGVALNOTEXISTS = "InvalidParameterValue.TagValNotExists"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TAGKEYEXCEEDED = "LimitExceeded.TagKeyExceeded"
//  LIMITEXCEEDED_TAGKEYPERRESOURCEEXCEEDED = "LimitExceeded.TagKeyPerResourceExceeded"
//  LIMITEXCEEDED_TAGNOTENOUGHQUOTA = "LimitExceeded.TagNotEnoughQuota"
//  LIMITEXCEEDED_TAGQUOTA = "LimitExceeded.TagQuota"
//  LIMITEXCEEDED_TAGQUOTAEXCEEDED = "LimitExceeded.TagQuotaExceeded"
//  LIMITEXCEEDED_TAGTAGSEXCEEDED = "LimitExceeded.TagTagsExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_TAGALLOCATE = "UnsupportedOperation.TagAllocate"
//  UNSUPPORTEDOPERATION_TAGFREE = "UnsupportedOperation.TagFree"
//  UNSUPPORTEDOPERATION_TAGNOTPERMIT = "UnsupportedOperation.TagNotPermit"
//  UNSUPPORTEDOPERATION_TAGSYSTEMRESERVEDTAGKEY = "UnsupportedOperation.TagSystemReservedTagKey"
func (c *Client) CreateSecurityGroup(request *CreateSecurityGroupRequest) (response *CreateSecurityGroupResponse, err error) {
    return c.CreateSecurityGroupWithContext(context.Background(), request)
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
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TAGDUPLICATEKEY = "InvalidParameterValue.TagDuplicateKey"
//  INVALIDPARAMETERVALUE_TAGDUPLICATERESOURCETYPE = "InvalidParameterValue.TagDuplicateResourceType"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEY = "InvalidParameterValue.TagInvalidKey"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEYLEN = "InvalidParameterValue.TagInvalidKeyLen"
//  INVALIDPARAMETERVALUE_TAGINVALIDVAL = "InvalidParameterValue.TagInvalidVal"
//  INVALIDPARAMETERVALUE_TAGKEYNOTEXISTS = "InvalidParameterValue.TagKeyNotExists"
//  INVALIDPARAMETERVALUE_TAGNOTALLOCATEDQUOTA = "InvalidParameterValue.TagNotAllocatedQuota"
//  INVALIDPARAMETERVALUE_TAGNOTEXISTED = "InvalidParameterValue.TagNotExisted"
//  INVALIDPARAMETERVALUE_TAGNOTSUPPORTTAG = "InvalidParameterValue.TagNotSupportTag"
//  INVALIDPARAMETERVALUE_TAGRESOURCEFORMATERROR = "InvalidParameterValue.TagResourceFormatError"
//  INVALIDPARAMETERVALUE_TAGTIMESTAMPEXCEEDED = "InvalidParameterValue.TagTimestampExceeded"
//  INVALIDPARAMETERVALUE_TAGVALNOTEXISTS = "InvalidParameterValue.TagValNotExists"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TAGKEYEXCEEDED = "LimitExceeded.TagKeyExceeded"
//  LIMITEXCEEDED_TAGKEYPERRESOURCEEXCEEDED = "LimitExceeded.TagKeyPerResourceExceeded"
//  LIMITEXCEEDED_TAGNOTENOUGHQUOTA = "LimitExceeded.TagNotEnoughQuota"
//  LIMITEXCEEDED_TAGQUOTA = "LimitExceeded.TagQuota"
//  LIMITEXCEEDED_TAGQUOTAEXCEEDED = "LimitExceeded.TagQuotaExceeded"
//  LIMITEXCEEDED_TAGTAGSEXCEEDED = "LimitExceeded.TagTagsExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_TAGALLOCATE = "UnsupportedOperation.TagAllocate"
//  UNSUPPORTEDOPERATION_TAGFREE = "UnsupportedOperation.TagFree"
//  UNSUPPORTEDOPERATION_TAGNOTPERMIT = "UnsupportedOperation.TagNotPermit"
//  UNSUPPORTEDOPERATION_TAGSYSTEMRESERVEDTAGKEY = "UnsupportedOperation.TagSystemReservedTagKey"
func (c *Client) CreateSecurityGroupWithContext(ctx context.Context, request *CreateSecurityGroupRequest) (response *CreateSecurityGroupResponse, err error) {
    if request == nil {
        request = NewCreateSecurityGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSecurityGroup require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to create security group policies.
//
// 
//
// For parameters of SecurityGroupPolicySet,
//
// <ul>
//
// <li>`Version`: The version number of a security group policy, which automatically increases by one each time you update the security policy, to prevent expiration of the updated routing policies. If it is left empty, any conflicts will be ignored.</li>
//
// <li>When creating the `Egress` and `Ingress` polices,<ul>
//
// <li>`Protocol`: `TCP`, `UDP`, `ICMP`, `ICMPV6`, `GRE`, or `ALL`.</li>
//
// <li>`CidrBlock`: A CIDR block in the correct format. </li>For 
//
// <li>`Ipv6CidrBlock`: An IPv6 CIDR block in the correct format. In a classic network, if an `Ipv6CidrBlock` contains private IPv6 addresses on Tencent Cloud for devices under your account other than CVMs, it does not mean this policy allows you to access these devices. The network isolation policies between tenants take priority over the private network policies in security groups.</li>
//
// <li>`SecurityGroupId`: ID of the security group. It can be the ID of security group to be modified, or the ID of other security group in the same project. All private IPs of all CVMs under the security group will be covered. If this field is used, the policy will automatically change according to the CVM associated with the group ID while being used to match network messages. You don’t need to change it manually.</li>
//
// <li>`Port`: A single port number such as 80, or a port range in the format of "8000-8010". This parameter is only available when the `Protocol` is `TCP` or `UDP`. Otherwise, `Protocol` and `Port` are mutually exclusive.</li>
//
// <li>`Action`: `ACCEPT` or `DROP`.</li>
//
// <li>`CidrBlock`, `Ipv6CidrBlock`, `SecurityGroupId`, and `AddressTemplate` are mutually exclusive. `Protocol` + `Port` and `ServiceTemplate` are mutually exclusive.</li>
//
// <li>You can only create policies in one direction in each request. To specify the `PolicyIndex` parameter, use the same index number in policies. If you want to insert a rule before the first rule, enter 0; if you want to add a rule after the last rule, leave it empty.</li>
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
    return c.CreateSecurityGroupPoliciesWithContext(context.Background(), request)
}

// CreateSecurityGroupPolicies
// This API is used to create security group policies.
//
// 
//
// For parameters of SecurityGroupPolicySet,
//
// <ul>
//
// <li>`Version`: The version number of a security group policy, which automatically increases by one each time you update the security policy, to prevent expiration of the updated routing policies. If it is left empty, any conflicts will be ignored.</li>
//
// <li>When creating the `Egress` and `Ingress` polices,<ul>
//
// <li>`Protocol`: `TCP`, `UDP`, `ICMP`, `ICMPV6`, `GRE`, or `ALL`.</li>
//
// <li>`CidrBlock`: A CIDR block in the correct format. </li>For 
//
// <li>`Ipv6CidrBlock`: An IPv6 CIDR block in the correct format. In a classic network, if an `Ipv6CidrBlock` contains private IPv6 addresses on Tencent Cloud for devices under your account other than CVMs, it does not mean this policy allows you to access these devices. The network isolation policies between tenants take priority over the private network policies in security groups.</li>
//
// <li>`SecurityGroupId`: ID of the security group. It can be the ID of security group to be modified, or the ID of other security group in the same project. All private IPs of all CVMs under the security group will be covered. If this field is used, the policy will automatically change according to the CVM associated with the group ID while being used to match network messages. You don’t need to change it manually.</li>
//
// <li>`Port`: A single port number such as 80, or a port range in the format of "8000-8010". This parameter is only available when the `Protocol` is `TCP` or `UDP`. Otherwise, `Protocol` and `Port` are mutually exclusive.</li>
//
// <li>`Action`: `ACCEPT` or `DROP`.</li>
//
// <li>`CidrBlock`, `Ipv6CidrBlock`, `SecurityGroupId`, and `AddressTemplate` are mutually exclusive. `Protocol` + `Port` and `ServiceTemplate` are mutually exclusive.</li>
//
// <li>You can only create policies in one direction in each request. To specify the `PolicyIndex` parameter, use the same index number in policies. If you want to insert a rule before the first rule, enter 0; if you want to add a rule after the last rule, leave it empty.</li>
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
func (c *Client) CreateSecurityGroupPoliciesWithContext(ctx context.Context, request *CreateSecurityGroupPoliciesRequest) (response *CreateSecurityGroupPoliciesResponse, err error) {
    if request == nil {
        request = NewCreateSecurityGroupPoliciesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSecurityGroupPolicies require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateSecurityGroupWithPolicies(request *CreateSecurityGroupWithPoliciesRequest) (response *CreateSecurityGroupWithPoliciesResponse, err error) {
    return c.CreateSecurityGroupWithPoliciesWithContext(context.Background(), request)
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
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateSecurityGroupWithPoliciesWithContext(ctx context.Context, request *CreateSecurityGroupWithPoliciesRequest) (response *CreateSecurityGroupWithPoliciesResponse, err error) {
    if request == nil {
        request = NewCreateSecurityGroupWithPoliciesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSecurityGroupWithPolicies require credential")
    }

    request.SetContext(ctx)
    
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
    return c.CreateServiceTemplateWithContext(context.Background(), request)
}

// CreateServiceTemplate
// This API (CreateServiceTemplate) is used to create a protocol port template.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateServiceTemplateWithContext(ctx context.Context, request *CreateServiceTemplateRequest) (response *CreateServiceTemplateResponse, err error) {
    if request == nil {
        request = NewCreateServiceTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateServiceTemplate require credential")
    }

    request.SetContext(ctx)
    
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
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) CreateServiceTemplateGroup(request *CreateServiceTemplateGroupRequest) (response *CreateServiceTemplateGroupResponse, err error) {
    return c.CreateServiceTemplateGroupWithContext(context.Background(), request)
}

// CreateServiceTemplateGroup
// This API (CreateServiceTemplateGroup) is used to create a protocol port template group.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) CreateServiceTemplateGroupWithContext(ctx context.Context, request *CreateServiceTemplateGroupRequest) (response *CreateServiceTemplateGroupResponse, err error) {
    if request == nil {
        request = NewCreateServiceTemplateGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateServiceTemplateGroup require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_TAGDUPLICATEKEY = "InvalidParameterValue.TagDuplicateKey"
//  INVALIDPARAMETERVALUE_TAGDUPLICATERESOURCETYPE = "InvalidParameterValue.TagDuplicateResourceType"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEY = "InvalidParameterValue.TagInvalidKey"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEYLEN = "InvalidParameterValue.TagInvalidKeyLen"
//  INVALIDPARAMETERVALUE_TAGINVALIDVAL = "InvalidParameterValue.TagInvalidVal"
//  INVALIDPARAMETERVALUE_TAGKEYNOTEXISTS = "InvalidParameterValue.TagKeyNotExists"
//  INVALIDPARAMETERVALUE_TAGNOTALLOCATEDQUOTA = "InvalidParameterValue.TagNotAllocatedQuota"
//  INVALIDPARAMETERVALUE_TAGNOTEXISTED = "InvalidParameterValue.TagNotExisted"
//  INVALIDPARAMETERVALUE_TAGNOTSUPPORTTAG = "InvalidParameterValue.TagNotSupportTag"
//  INVALIDPARAMETERVALUE_TAGRESOURCEFORMATERROR = "InvalidParameterValue.TagResourceFormatError"
//  INVALIDPARAMETERVALUE_TAGTIMESTAMPEXCEEDED = "InvalidParameterValue.TagTimestampExceeded"
//  INVALIDPARAMETERVALUE_TAGVALNOTEXISTS = "InvalidParameterValue.TagValNotExists"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_ZONECONFLICT = "InvalidParameterValue.ZoneConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TAGKEYEXCEEDED = "LimitExceeded.TagKeyExceeded"
//  LIMITEXCEEDED_TAGKEYPERRESOURCEEXCEEDED = "LimitExceeded.TagKeyPerResourceExceeded"
//  LIMITEXCEEDED_TAGNOTENOUGHQUOTA = "LimitExceeded.TagNotEnoughQuota"
//  LIMITEXCEEDED_TAGQUOTA = "LimitExceeded.TagQuota"
//  LIMITEXCEEDED_TAGQUOTAEXCEEDED = "LimitExceeded.TagQuotaExceeded"
//  LIMITEXCEEDED_TAGTAGSEXCEEDED = "LimitExceeded.TagTagsExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_APPIDMISMATCH = "UnsupportedOperation.AppIdMismatch"
//  UNSUPPORTEDOPERATION_DCGATEWAYSNOTFOUNDINVPC = "UnsupportedOperation.DcGatewaysNotFoundInVpc"
//  UNSUPPORTEDOPERATION_TAGALLOCATE = "UnsupportedOperation.TagAllocate"
//  UNSUPPORTEDOPERATION_TAGFREE = "UnsupportedOperation.TagFree"
//  UNSUPPORTEDOPERATION_TAGNOTPERMIT = "UnsupportedOperation.TagNotPermit"
//  UNSUPPORTEDOPERATION_TAGSYSTEMRESERVEDTAGKEY = "UnsupportedOperation.TagSystemReservedTagKey"
func (c *Client) CreateSubnet(request *CreateSubnetRequest) (response *CreateSubnetResponse, err error) {
    return c.CreateSubnetWithContext(context.Background(), request)
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
//  INVALIDPARAMETERVALUE_TAGDUPLICATEKEY = "InvalidParameterValue.TagDuplicateKey"
//  INVALIDPARAMETERVALUE_TAGDUPLICATERESOURCETYPE = "InvalidParameterValue.TagDuplicateResourceType"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEY = "InvalidParameterValue.TagInvalidKey"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEYLEN = "InvalidParameterValue.TagInvalidKeyLen"
//  INVALIDPARAMETERVALUE_TAGINVALIDVAL = "InvalidParameterValue.TagInvalidVal"
//  INVALIDPARAMETERVALUE_TAGKEYNOTEXISTS = "InvalidParameterValue.TagKeyNotExists"
//  INVALIDPARAMETERVALUE_TAGNOTALLOCATEDQUOTA = "InvalidParameterValue.TagNotAllocatedQuota"
//  INVALIDPARAMETERVALUE_TAGNOTEXISTED = "InvalidParameterValue.TagNotExisted"
//  INVALIDPARAMETERVALUE_TAGNOTSUPPORTTAG = "InvalidParameterValue.TagNotSupportTag"
//  INVALIDPARAMETERVALUE_TAGRESOURCEFORMATERROR = "InvalidParameterValue.TagResourceFormatError"
//  INVALIDPARAMETERVALUE_TAGTIMESTAMPEXCEEDED = "InvalidParameterValue.TagTimestampExceeded"
//  INVALIDPARAMETERVALUE_TAGVALNOTEXISTS = "InvalidParameterValue.TagValNotExists"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_ZONECONFLICT = "InvalidParameterValue.ZoneConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TAGKEYEXCEEDED = "LimitExceeded.TagKeyExceeded"
//  LIMITEXCEEDED_TAGKEYPERRESOURCEEXCEEDED = "LimitExceeded.TagKeyPerResourceExceeded"
//  LIMITEXCEEDED_TAGNOTENOUGHQUOTA = "LimitExceeded.TagNotEnoughQuota"
//  LIMITEXCEEDED_TAGQUOTA = "LimitExceeded.TagQuota"
//  LIMITEXCEEDED_TAGQUOTAEXCEEDED = "LimitExceeded.TagQuotaExceeded"
//  LIMITEXCEEDED_TAGTAGSEXCEEDED = "LimitExceeded.TagTagsExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_APPIDMISMATCH = "UnsupportedOperation.AppIdMismatch"
//  UNSUPPORTEDOPERATION_DCGATEWAYSNOTFOUNDINVPC = "UnsupportedOperation.DcGatewaysNotFoundInVpc"
//  UNSUPPORTEDOPERATION_TAGALLOCATE = "UnsupportedOperation.TagAllocate"
//  UNSUPPORTEDOPERATION_TAGFREE = "UnsupportedOperation.TagFree"
//  UNSUPPORTEDOPERATION_TAGNOTPERMIT = "UnsupportedOperation.TagNotPermit"
//  UNSUPPORTEDOPERATION_TAGSYSTEMRESERVEDTAGKEY = "UnsupportedOperation.TagSystemReservedTagKey"
func (c *Client) CreateSubnetWithContext(ctx context.Context, request *CreateSubnetRequest) (response *CreateSubnetResponse, err error) {
    if request == nil {
        request = NewCreateSubnetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSubnet require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_TAGDUPLICATEKEY = "InvalidParameterValue.TagDuplicateKey"
//  INVALIDPARAMETERVALUE_TAGDUPLICATERESOURCETYPE = "InvalidParameterValue.TagDuplicateResourceType"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEY = "InvalidParameterValue.TagInvalidKey"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEYLEN = "InvalidParameterValue.TagInvalidKeyLen"
//  INVALIDPARAMETERVALUE_TAGINVALIDVAL = "InvalidParameterValue.TagInvalidVal"
//  INVALIDPARAMETERVALUE_TAGKEYNOTEXISTS = "InvalidParameterValue.TagKeyNotExists"
//  INVALIDPARAMETERVALUE_TAGNOTALLOCATEDQUOTA = "InvalidParameterValue.TagNotAllocatedQuota"
//  INVALIDPARAMETERVALUE_TAGNOTEXISTED = "InvalidParameterValue.TagNotExisted"
//  INVALIDPARAMETERVALUE_TAGNOTSUPPORTTAG = "InvalidParameterValue.TagNotSupportTag"
//  INVALIDPARAMETERVALUE_TAGRESOURCEFORMATERROR = "InvalidParameterValue.TagResourceFormatError"
//  INVALIDPARAMETERVALUE_TAGTIMESTAMPEXCEEDED = "InvalidParameterValue.TagTimestampExceeded"
//  INVALIDPARAMETERVALUE_TAGVALNOTEXISTS = "InvalidParameterValue.TagValNotExists"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_ZONECONFLICT = "InvalidParameterValue.ZoneConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TAGKEYEXCEEDED = "LimitExceeded.TagKeyExceeded"
//  LIMITEXCEEDED_TAGKEYPERRESOURCEEXCEEDED = "LimitExceeded.TagKeyPerResourceExceeded"
//  LIMITEXCEEDED_TAGNOTENOUGHQUOTA = "LimitExceeded.TagNotEnoughQuota"
//  LIMITEXCEEDED_TAGQUOTA = "LimitExceeded.TagQuota"
//  LIMITEXCEEDED_TAGQUOTAEXCEEDED = "LimitExceeded.TagQuotaExceeded"
//  LIMITEXCEEDED_TAGTAGSEXCEEDED = "LimitExceeded.TagTagsExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_APPIDMISMATCH = "UnsupportedOperation.AppIdMismatch"
//  UNSUPPORTEDOPERATION_DCGATEWAYSNOTFOUNDINVPC = "UnsupportedOperation.DcGatewaysNotFoundInVpc"
//  UNSUPPORTEDOPERATION_TAGALLOCATE = "UnsupportedOperation.TagAllocate"
//  UNSUPPORTEDOPERATION_TAGFREE = "UnsupportedOperation.TagFree"
//  UNSUPPORTEDOPERATION_TAGNOTPERMIT = "UnsupportedOperation.TagNotPermit"
//  UNSUPPORTEDOPERATION_TAGSYSTEMRESERVEDTAGKEY = "UnsupportedOperation.TagSystemReservedTagKey"
func (c *Client) CreateSubnets(request *CreateSubnetsRequest) (response *CreateSubnetsResponse, err error) {
    return c.CreateSubnetsWithContext(context.Background(), request)
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
//  INVALIDPARAMETERVALUE_TAGDUPLICATEKEY = "InvalidParameterValue.TagDuplicateKey"
//  INVALIDPARAMETERVALUE_TAGDUPLICATERESOURCETYPE = "InvalidParameterValue.TagDuplicateResourceType"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEY = "InvalidParameterValue.TagInvalidKey"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEYLEN = "InvalidParameterValue.TagInvalidKeyLen"
//  INVALIDPARAMETERVALUE_TAGINVALIDVAL = "InvalidParameterValue.TagInvalidVal"
//  INVALIDPARAMETERVALUE_TAGKEYNOTEXISTS = "InvalidParameterValue.TagKeyNotExists"
//  INVALIDPARAMETERVALUE_TAGNOTALLOCATEDQUOTA = "InvalidParameterValue.TagNotAllocatedQuota"
//  INVALIDPARAMETERVALUE_TAGNOTEXISTED = "InvalidParameterValue.TagNotExisted"
//  INVALIDPARAMETERVALUE_TAGNOTSUPPORTTAG = "InvalidParameterValue.TagNotSupportTag"
//  INVALIDPARAMETERVALUE_TAGRESOURCEFORMATERROR = "InvalidParameterValue.TagResourceFormatError"
//  INVALIDPARAMETERVALUE_TAGTIMESTAMPEXCEEDED = "InvalidParameterValue.TagTimestampExceeded"
//  INVALIDPARAMETERVALUE_TAGVALNOTEXISTS = "InvalidParameterValue.TagValNotExists"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_ZONECONFLICT = "InvalidParameterValue.ZoneConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TAGKEYEXCEEDED = "LimitExceeded.TagKeyExceeded"
//  LIMITEXCEEDED_TAGKEYPERRESOURCEEXCEEDED = "LimitExceeded.TagKeyPerResourceExceeded"
//  LIMITEXCEEDED_TAGNOTENOUGHQUOTA = "LimitExceeded.TagNotEnoughQuota"
//  LIMITEXCEEDED_TAGQUOTA = "LimitExceeded.TagQuota"
//  LIMITEXCEEDED_TAGQUOTAEXCEEDED = "LimitExceeded.TagQuotaExceeded"
//  LIMITEXCEEDED_TAGTAGSEXCEEDED = "LimitExceeded.TagTagsExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_APPIDMISMATCH = "UnsupportedOperation.AppIdMismatch"
//  UNSUPPORTEDOPERATION_DCGATEWAYSNOTFOUNDINVPC = "UnsupportedOperation.DcGatewaysNotFoundInVpc"
//  UNSUPPORTEDOPERATION_TAGALLOCATE = "UnsupportedOperation.TagAllocate"
//  UNSUPPORTEDOPERATION_TAGFREE = "UnsupportedOperation.TagFree"
//  UNSUPPORTEDOPERATION_TAGNOTPERMIT = "UnsupportedOperation.TagNotPermit"
//  UNSUPPORTEDOPERATION_TAGSYSTEMRESERVEDTAGKEY = "UnsupportedOperation.TagSystemReservedTagKey"
func (c *Client) CreateSubnetsWithContext(ctx context.Context, request *CreateSubnetsRequest) (response *CreateSubnetsResponse, err error) {
    if request == nil {
        request = NewCreateSubnetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSubnets require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_TAGDUPLICATEKEY = "InvalidParameterValue.TagDuplicateKey"
//  INVALIDPARAMETERVALUE_TAGDUPLICATERESOURCETYPE = "InvalidParameterValue.TagDuplicateResourceType"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEY = "InvalidParameterValue.TagInvalidKey"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEYLEN = "InvalidParameterValue.TagInvalidKeyLen"
//  INVALIDPARAMETERVALUE_TAGINVALIDVAL = "InvalidParameterValue.TagInvalidVal"
//  INVALIDPARAMETERVALUE_TAGKEYNOTEXISTS = "InvalidParameterValue.TagKeyNotExists"
//  INVALIDPARAMETERVALUE_TAGNOTALLOCATEDQUOTA = "InvalidParameterValue.TagNotAllocatedQuota"
//  INVALIDPARAMETERVALUE_TAGNOTEXISTED = "InvalidParameterValue.TagNotExisted"
//  INVALIDPARAMETERVALUE_TAGNOTSUPPORTTAG = "InvalidParameterValue.TagNotSupportTag"
//  INVALIDPARAMETERVALUE_TAGRESOURCEFORMATERROR = "InvalidParameterValue.TagResourceFormatError"
//  INVALIDPARAMETERVALUE_TAGTIMESTAMPEXCEEDED = "InvalidParameterValue.TagTimestampExceeded"
//  INVALIDPARAMETERVALUE_TAGVALNOTEXISTS = "InvalidParameterValue.TagValNotExists"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TAGKEYEXCEEDED = "LimitExceeded.TagKeyExceeded"
//  LIMITEXCEEDED_TAGKEYPERRESOURCEEXCEEDED = "LimitExceeded.TagKeyPerResourceExceeded"
//  LIMITEXCEEDED_TAGNOTENOUGHQUOTA = "LimitExceeded.TagNotEnoughQuota"
//  LIMITEXCEEDED_TAGQUOTA = "LimitExceeded.TagQuota"
//  LIMITEXCEEDED_TAGQUOTAEXCEEDED = "LimitExceeded.TagQuotaExceeded"
//  LIMITEXCEEDED_TAGTAGSEXCEEDED = "LimitExceeded.TagTagsExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_TAGALLOCATE = "UnsupportedOperation.TagAllocate"
//  UNSUPPORTEDOPERATION_TAGFREE = "UnsupportedOperation.TagFree"
//  UNSUPPORTEDOPERATION_TAGNOTPERMIT = "UnsupportedOperation.TagNotPermit"
//  UNSUPPORTEDOPERATION_TAGSYSTEMRESERVEDTAGKEY = "UnsupportedOperation.TagSystemReservedTagKey"
func (c *Client) CreateVpc(request *CreateVpcRequest) (response *CreateVpcResponse, err error) {
    return c.CreateVpcWithContext(context.Background(), request)
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
//  INVALIDPARAMETERVALUE_TAGDUPLICATEKEY = "InvalidParameterValue.TagDuplicateKey"
//  INVALIDPARAMETERVALUE_TAGDUPLICATERESOURCETYPE = "InvalidParameterValue.TagDuplicateResourceType"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEY = "InvalidParameterValue.TagInvalidKey"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEYLEN = "InvalidParameterValue.TagInvalidKeyLen"
//  INVALIDPARAMETERVALUE_TAGINVALIDVAL = "InvalidParameterValue.TagInvalidVal"
//  INVALIDPARAMETERVALUE_TAGKEYNOTEXISTS = "InvalidParameterValue.TagKeyNotExists"
//  INVALIDPARAMETERVALUE_TAGNOTALLOCATEDQUOTA = "InvalidParameterValue.TagNotAllocatedQuota"
//  INVALIDPARAMETERVALUE_TAGNOTEXISTED = "InvalidParameterValue.TagNotExisted"
//  INVALIDPARAMETERVALUE_TAGNOTSUPPORTTAG = "InvalidParameterValue.TagNotSupportTag"
//  INVALIDPARAMETERVALUE_TAGRESOURCEFORMATERROR = "InvalidParameterValue.TagResourceFormatError"
//  INVALIDPARAMETERVALUE_TAGTIMESTAMPEXCEEDED = "InvalidParameterValue.TagTimestampExceeded"
//  INVALIDPARAMETERVALUE_TAGVALNOTEXISTS = "InvalidParameterValue.TagValNotExists"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TAGKEYEXCEEDED = "LimitExceeded.TagKeyExceeded"
//  LIMITEXCEEDED_TAGKEYPERRESOURCEEXCEEDED = "LimitExceeded.TagKeyPerResourceExceeded"
//  LIMITEXCEEDED_TAGNOTENOUGHQUOTA = "LimitExceeded.TagNotEnoughQuota"
//  LIMITEXCEEDED_TAGQUOTA = "LimitExceeded.TagQuota"
//  LIMITEXCEEDED_TAGQUOTAEXCEEDED = "LimitExceeded.TagQuotaExceeded"
//  LIMITEXCEEDED_TAGTAGSEXCEEDED = "LimitExceeded.TagTagsExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_TAGALLOCATE = "UnsupportedOperation.TagAllocate"
//  UNSUPPORTEDOPERATION_TAGFREE = "UnsupportedOperation.TagFree"
//  UNSUPPORTEDOPERATION_TAGNOTPERMIT = "UnsupportedOperation.TagNotPermit"
//  UNSUPPORTEDOPERATION_TAGSYSTEMRESERVEDTAGKEY = "UnsupportedOperation.TagSystemReservedTagKey"
func (c *Client) CreateVpcWithContext(ctx context.Context, request *CreateVpcRequest) (response *CreateVpcResponse, err error) {
    if request == nil {
        request = NewCreateVpcRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVpc require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SERVICEWHITELISTNOTADDED = "ResourceUnavailable.ServiceWhiteListNotAdded"
//  UNSUPPORTEDOPERATION_ENDPOINTSERVICE = "UnsupportedOperation.EndPointService"
//  UNSUPPORTEDOPERATION_INSUFFICIENTFUNDS = "UnsupportedOperation.InsufficientFunds"
//  UNSUPPORTEDOPERATION_SPECIALENDPOINTSERVICE = "UnsupportedOperation.SpecialEndPointService"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
func (c *Client) CreateVpcEndPoint(request *CreateVpcEndPointRequest) (response *CreateVpcEndPointResponse, err error) {
    return c.CreateVpcEndPointWithContext(context.Background(), request)
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
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SERVICEWHITELISTNOTADDED = "ResourceUnavailable.ServiceWhiteListNotAdded"
//  UNSUPPORTEDOPERATION_ENDPOINTSERVICE = "UnsupportedOperation.EndPointService"
//  UNSUPPORTEDOPERATION_INSUFFICIENTFUNDS = "UnsupportedOperation.InsufficientFunds"
//  UNSUPPORTEDOPERATION_SPECIALENDPOINTSERVICE = "UnsupportedOperation.SpecialEndPointService"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
func (c *Client) CreateVpcEndPointWithContext(ctx context.Context, request *CreateVpcEndPointRequest) (response *CreateVpcEndPointResponse, err error) {
    if request == nil {
        request = NewCreateVpcEndPointRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVpcEndPoint require credential")
    }

    request.SetContext(ctx)
    
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
func (c *Client) CreateVpcEndPointService(request *CreateVpcEndPointServiceRequest) (response *CreateVpcEndPointServiceResponse, err error) {
    return c.CreateVpcEndPointServiceWithContext(context.Background(), request)
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
func (c *Client) CreateVpcEndPointServiceWithContext(ctx context.Context, request *CreateVpcEndPointServiceRequest) (response *CreateVpcEndPointServiceResponse, err error) {
    if request == nil {
        request = NewCreateVpcEndPointServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVpcEndPointService require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_UINNOTFOUND = "UnsupportedOperation.UinNotFound"
func (c *Client) CreateVpcEndPointServiceWhiteList(request *CreateVpcEndPointServiceWhiteListRequest) (response *CreateVpcEndPointServiceWhiteListResponse, err error) {
    return c.CreateVpcEndPointServiceWhiteListWithContext(context.Background(), request)
}

// CreateVpcEndPointServiceWhiteList
// This API is used to create the endpoint service allowlist.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_UINNOTFOUND = "UnsupportedOperation.UinNotFound"
func (c *Client) CreateVpcEndPointServiceWhiteListWithContext(ctx context.Context, request *CreateVpcEndPointServiceWhiteListRequest) (response *CreateVpcEndPointServiceWhiteListResponse, err error) {
    if request == nil {
        request = NewCreateVpcEndPointServiceWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVpcEndPointServiceWhiteList require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to create a VPN tunnel.
//
// >?This API is async. You can call the [`DescribeVpcTaskResult`](https://intl.cloud.tencent.com/document/api/215/59037?from_cn_redirect=1) API to query the task result. When the task is completed, you can continue other tasks.
//
// >
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TAGDUPLICATEKEY = "InvalidParameterValue.TagDuplicateKey"
//  INVALIDPARAMETERVALUE_TAGDUPLICATERESOURCETYPE = "InvalidParameterValue.TagDuplicateResourceType"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEY = "InvalidParameterValue.TagInvalidKey"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEYLEN = "InvalidParameterValue.TagInvalidKeyLen"
//  INVALIDPARAMETERVALUE_TAGINVALIDVAL = "InvalidParameterValue.TagInvalidVal"
//  INVALIDPARAMETERVALUE_TAGKEYNOTEXISTS = "InvalidParameterValue.TagKeyNotExists"
//  INVALIDPARAMETERVALUE_TAGNOTALLOCATEDQUOTA = "InvalidParameterValue.TagNotAllocatedQuota"
//  INVALIDPARAMETERVALUE_TAGNOTEXISTED = "InvalidParameterValue.TagNotExisted"
//  INVALIDPARAMETERVALUE_TAGNOTSUPPORTTAG = "InvalidParameterValue.TagNotSupportTag"
//  INVALIDPARAMETERVALUE_TAGRESOURCEFORMATERROR = "InvalidParameterValue.TagResourceFormatError"
//  INVALIDPARAMETERVALUE_TAGTIMESTAMPEXCEEDED = "InvalidParameterValue.TagTimestampExceeded"
//  INVALIDPARAMETERVALUE_TAGVALNOTEXISTS = "InvalidParameterValue.TagValNotExists"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_VPCCIDRCONFLICT = "InvalidParameterValue.VpcCidrConflict"
//  INVALIDPARAMETERVALUE_VPNCONNCIDRCONFLICT = "InvalidParameterValue.VpnConnCidrConflict"
//  INVALIDPARAMETERVALUE_VPNCONNHEALTHCHECKIPCONFLICT = "InvalidParameterValue.VpnConnHealthCheckIpConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TAGKEYEXCEEDED = "LimitExceeded.TagKeyExceeded"
//  LIMITEXCEEDED_TAGKEYPERRESOURCEEXCEEDED = "LimitExceeded.TagKeyPerResourceExceeded"
//  LIMITEXCEEDED_TAGNOTENOUGHQUOTA = "LimitExceeded.TagNotEnoughQuota"
//  LIMITEXCEEDED_TAGQUOTA = "LimitExceeded.TagQuota"
//  LIMITEXCEEDED_TAGQUOTAEXCEEDED = "LimitExceeded.TagQuotaExceeded"
//  LIMITEXCEEDED_TAGTAGSEXCEEDED = "LimitExceeded.TagTagsExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_TAGALLOCATE = "UnsupportedOperation.TagAllocate"
//  UNSUPPORTEDOPERATION_TAGFREE = "UnsupportedOperation.TagFree"
//  UNSUPPORTEDOPERATION_TAGNOTPERMIT = "UnsupportedOperation.TagNotPermit"
//  UNSUPPORTEDOPERATION_TAGSYSTEMRESERVEDTAGKEY = "UnsupportedOperation.TagSystemReservedTagKey"
func (c *Client) CreateVpnConnection(request *CreateVpnConnectionRequest) (response *CreateVpnConnectionResponse, err error) {
    return c.CreateVpnConnectionWithContext(context.Background(), request)
}

// CreateVpnConnection
// This API is used to create a VPN tunnel.
//
// >?This API is async. You can call the [`DescribeVpcTaskResult`](https://intl.cloud.tencent.com/document/api/215/59037?from_cn_redirect=1) API to query the task result. When the task is completed, you can continue other tasks.
//
// >
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TAGDUPLICATEKEY = "InvalidParameterValue.TagDuplicateKey"
//  INVALIDPARAMETERVALUE_TAGDUPLICATERESOURCETYPE = "InvalidParameterValue.TagDuplicateResourceType"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEY = "InvalidParameterValue.TagInvalidKey"
//  INVALIDPARAMETERVALUE_TAGINVALIDKEYLEN = "InvalidParameterValue.TagInvalidKeyLen"
//  INVALIDPARAMETERVALUE_TAGINVALIDVAL = "InvalidParameterValue.TagInvalidVal"
//  INVALIDPARAMETERVALUE_TAGKEYNOTEXISTS = "InvalidParameterValue.TagKeyNotExists"
//  INVALIDPARAMETERVALUE_TAGNOTALLOCATEDQUOTA = "InvalidParameterValue.TagNotAllocatedQuota"
//  INVALIDPARAMETERVALUE_TAGNOTEXISTED = "InvalidParameterValue.TagNotExisted"
//  INVALIDPARAMETERVALUE_TAGNOTSUPPORTTAG = "InvalidParameterValue.TagNotSupportTag"
//  INVALIDPARAMETERVALUE_TAGRESOURCEFORMATERROR = "InvalidParameterValue.TagResourceFormatError"
//  INVALIDPARAMETERVALUE_TAGTIMESTAMPEXCEEDED = "InvalidParameterValue.TagTimestampExceeded"
//  INVALIDPARAMETERVALUE_TAGVALNOTEXISTS = "InvalidParameterValue.TagValNotExists"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_VPCCIDRCONFLICT = "InvalidParameterValue.VpcCidrConflict"
//  INVALIDPARAMETERVALUE_VPNCONNCIDRCONFLICT = "InvalidParameterValue.VpnConnCidrConflict"
//  INVALIDPARAMETERVALUE_VPNCONNHEALTHCHECKIPCONFLICT = "InvalidParameterValue.VpnConnHealthCheckIpConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TAGKEYEXCEEDED = "LimitExceeded.TagKeyExceeded"
//  LIMITEXCEEDED_TAGKEYPERRESOURCEEXCEEDED = "LimitExceeded.TagKeyPerResourceExceeded"
//  LIMITEXCEEDED_TAGNOTENOUGHQUOTA = "LimitExceeded.TagNotEnoughQuota"
//  LIMITEXCEEDED_TAGQUOTA = "LimitExceeded.TagQuota"
//  LIMITEXCEEDED_TAGQUOTAEXCEEDED = "LimitExceeded.TagQuotaExceeded"
//  LIMITEXCEEDED_TAGTAGSEXCEEDED = "LimitExceeded.TagTagsExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_TAGALLOCATE = "UnsupportedOperation.TagAllocate"
//  UNSUPPORTEDOPERATION_TAGFREE = "UnsupportedOperation.TagFree"
//  UNSUPPORTEDOPERATION_TAGNOTPERMIT = "UnsupportedOperation.TagNotPermit"
//  UNSUPPORTEDOPERATION_TAGSYSTEMRESERVEDTAGKEY = "UnsupportedOperation.TagSystemReservedTagKey"
func (c *Client) CreateVpnConnectionWithContext(ctx context.Context, request *CreateVpnConnectionRequest) (response *CreateVpnConnectionResponse, err error) {
    if request == nil {
        request = NewCreateVpnConnectionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVpnConnection require credential")
    }

    request.SetContext(ctx)
    
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
    return c.CreateVpnGatewayWithContext(context.Background(), request)
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
func (c *Client) CreateVpnGatewayWithContext(ctx context.Context, request *CreateVpnGatewayRequest) (response *CreateVpnGatewayResponse, err error) {
    if request == nil {
        request = NewCreateVpnGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVpnGateway require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVpnGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVpnGatewayRoutesRequest() (request *CreateVpnGatewayRoutesRequest) {
    request = &CreateVpnGatewayRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateVpnGatewayRoutes")
    
    
    return
}

func NewCreateVpnGatewayRoutesResponse() (response *CreateVpnGatewayRoutesResponse) {
    response = &CreateVpnGatewayRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateVpnGatewayRoutes
// This API is used to create destination routes of a route-based VPN gateway.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateVpnGatewayRoutes(request *CreateVpnGatewayRoutesRequest) (response *CreateVpnGatewayRoutesResponse, err error) {
    return c.CreateVpnGatewayRoutesWithContext(context.Background(), request)
}

// CreateVpnGatewayRoutes
// This API is used to create destination routes of a route-based VPN gateway.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateVpnGatewayRoutesWithContext(ctx context.Context, request *CreateVpnGatewayRoutesRequest) (response *CreateVpnGatewayRoutesResponse, err error) {
    if request == nil {
        request = NewCreateVpnGatewayRoutesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVpnGatewayRoutes require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVpnGatewayRoutesResponse()
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
    return c.DeleteAddressTemplateWithContext(context.Background(), request)
}

// DeleteAddressTemplate
// This API (DeleteAddressTemplate) is used to delete an IP address template.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) DeleteAddressTemplateWithContext(ctx context.Context, request *DeleteAddressTemplateRequest) (response *DeleteAddressTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteAddressTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAddressTemplate require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteAddressTemplateGroupWithContext(context.Background(), request)
}

// DeleteAddressTemplateGroup
// This API (DeleteAddressTemplateGroup) is used to delete an IP address template group.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) DeleteAddressTemplateGroupWithContext(ctx context.Context, request *DeleteAddressTemplateGroupRequest) (response *DeleteAddressTemplateGroupResponse, err error) {
    if request == nil {
        request = NewDeleteAddressTemplateGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAddressTemplateGroup require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to delete secondary CIDR blocks. This API is in beta test. To use it, please submit a ticket.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteAssistantCidr(request *DeleteAssistantCidrRequest) (response *DeleteAssistantCidrResponse, err error) {
    return c.DeleteAssistantCidrWithContext(context.Background(), request)
}

// DeleteAssistantCidr
// This API is used to delete secondary CIDR blocks. This API is in beta test. To use it, please submit a ticket.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteAssistantCidrWithContext(ctx context.Context, request *DeleteAssistantCidrRequest) (response *DeleteAssistantCidrResponse, err error) {
    if request == nil {
        request = NewDeleteAssistantCidrRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAssistantCidr require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION_INVALIDREGION = "FailedOperation.InvalidRegion"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGEIDMALFORMED = "InvalidParameterValue.BandwidthPackageIdMalformed"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGEINUSE = "InvalidParameterValue.BandwidthPackageInUse"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGENOTFOUND = "InvalidParameterValue.BandwidthPackageNotFound"
//  UNSUPPORTEDOPERATION_BANDWIDTHPACKAGEIDNOTSUPPORTED = "UnsupportedOperation.BandwidthPackageIdNotSupported"
//  UNSUPPORTEDOPERATION_INVALIDADDRESSSTATE = "UnsupportedOperation.InvalidAddressState"
func (c *Client) DeleteBandwidthPackage(request *DeleteBandwidthPackageRequest) (response *DeleteBandwidthPackageResponse, err error) {
    return c.DeleteBandwidthPackageWithContext(context.Background(), request)
}

// DeleteBandwidthPackage
// This API is used to delete bandwidth packages, including [device bandwidth packages](https://intl.cloud.tencent.com/document/product/684/15246?from_cn_redirect=1#.E8.AE.BE.E5.A4.87.E5.B8.A6.E5.AE.BD.E5.8C.85) and [IP bandwidth packages](https://intl.cloud.tencent.com/document/product/684/15246?from_cn_redirect=1#ip-.E5.B8.A6.E5.AE.BD.E5.8C.85).
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDREGION = "FailedOperation.InvalidRegion"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGEIDMALFORMED = "InvalidParameterValue.BandwidthPackageIdMalformed"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGEINUSE = "InvalidParameterValue.BandwidthPackageInUse"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGENOTFOUND = "InvalidParameterValue.BandwidthPackageNotFound"
//  UNSUPPORTEDOPERATION_BANDWIDTHPACKAGEIDNOTSUPPORTED = "UnsupportedOperation.BandwidthPackageIdNotSupported"
//  UNSUPPORTEDOPERATION_INVALIDADDRESSSTATE = "UnsupportedOperation.InvalidAddressState"
func (c *Client) DeleteBandwidthPackageWithContext(ctx context.Context, request *DeleteBandwidthPackageRequest) (response *DeleteBandwidthPackageResponse, err error) {
    if request == nil {
        request = NewDeleteBandwidthPackageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBandwidthPackage require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BANDWIDTHNOTEXPIRED = "UnsupportedOperation.BandwidthNotExpired"
//  UNSUPPORTEDOPERATION_CCNHASFLOWLOG = "UnsupportedOperation.CcnHasFlowLog"
func (c *Client) DeleteCcn(request *DeleteCcnRequest) (response *DeleteCcnResponse, err error) {
    return c.DeleteCcnWithContext(context.Background(), request)
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BANDWIDTHNOTEXPIRED = "UnsupportedOperation.BandwidthNotExpired"
//  UNSUPPORTEDOPERATION_CCNHASFLOWLOG = "UnsupportedOperation.CcnHasFlowLog"
func (c *Client) DeleteCcnWithContext(ctx context.Context, request *DeleteCcnRequest) (response *DeleteCcnResponse, err error) {
    if request == nil {
        request = NewDeleteCcnRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCcn require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteCustomerGatewayWithContext(context.Background(), request)
}

// DeleteCustomerGateway
// This API (DeleteCustomerGateway) is used to delete customer gateways.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteCustomerGatewayWithContext(ctx context.Context, request *DeleteCustomerGatewayRequest) (response *DeleteCustomerGatewayResponse, err error) {
    if request == nil {
        request = NewDeleteCustomerGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCustomerGateway require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteDirectConnectGatewayWithContext(context.Background(), request)
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
func (c *Client) DeleteDirectConnectGatewayWithContext(ctx context.Context, request *DeleteDirectConnectGatewayRequest) (response *DeleteDirectConnectGatewayResponse, err error) {
    if request == nil {
        request = NewDeleteDirectConnectGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDirectConnectGateway require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteDirectConnectGatewayCcnRoutesWithContext(context.Background(), request)
}

// DeleteDirectConnectGatewayCcnRoutes
// This API (DeleteDirectConnectGatewayCcnRoutes) is used to delete the CCN routes (IDC IP range) of a Direct Connect gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDirectConnectGatewayCcnRoutesWithContext(ctx context.Context, request *DeleteDirectConnectGatewayCcnRoutesRequest) (response *DeleteDirectConnectGatewayCcnRoutesResponse, err error) {
    if request == nil {
        request = NewDeleteDirectConnectGatewayCcnRoutesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDirectConnectGatewayCcnRoutes require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteFlowLogWithContext(context.Background(), request)
}

// DeleteFlowLog
// This API is used to delete a flow log.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteFlowLogWithContext(ctx context.Context, request *DeleteFlowLogRequest) (response *DeleteFlowLogResponse, err error) {
    if request == nil {
        request = NewDeleteFlowLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteFlowLog require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to delete an HAVIP. <br />
//
// This API is completed asynchronously. If you need to query the execution result of an async task, please use the `RequestId` returned by this API to poll the `DescribeVpcTaskResult` API.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteHaVip(request *DeleteHaVipRequest) (response *DeleteHaVipResponse, err error) {
    return c.DeleteHaVipWithContext(context.Background(), request)
}

// DeleteHaVip
// This API is used to delete an HAVIP. <br />
//
// This API is completed asynchronously. If you need to query the execution result of an async task, please use the `RequestId` returned by this API to poll the `DescribeVpcTaskResult` API.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteHaVipWithContext(ctx context.Context, request *DeleteHaVipRequest) (response *DeleteHaVipResponse, err error) {
    if request == nil {
        request = NewDeleteHaVipRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteHaVip require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteLocalGatewayWithContext(context.Background(), request)
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
func (c *Client) DeleteLocalGatewayWithContext(ctx context.Context, request *DeleteLocalGatewayRequest) (response *DeleteLocalGatewayResponse, err error) {
    if request == nil {
        request = NewDeleteLocalGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLocalGateway require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to delete a NAT gateway.
//
// When a NAT gateway is deleted, all routes containing this gateway are deleted automatically, and the elastic IP is unbound.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) DeleteNatGateway(request *DeleteNatGatewayRequest) (response *DeleteNatGatewayResponse, err error) {
    return c.DeleteNatGatewayWithContext(context.Background(), request)
}

// DeleteNatGateway
// This API is used to delete a NAT gateway.
//
// When a NAT gateway is deleted, all routes containing this gateway are deleted automatically, and the elastic IP is unbound.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) DeleteNatGatewayWithContext(ctx context.Context, request *DeleteNatGatewayRequest) (response *DeleteNatGatewayResponse, err error) {
    if request == nil {
        request = NewDeleteNatGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNatGateway require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to delete the port forwarding rule of a NAT gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) DeleteNatGatewayDestinationIpPortTranslationNatRule(request *DeleteNatGatewayDestinationIpPortTranslationNatRuleRequest) (response *DeleteNatGatewayDestinationIpPortTranslationNatRuleResponse, err error) {
    return c.DeleteNatGatewayDestinationIpPortTranslationNatRuleWithContext(context.Background(), request)
}

// DeleteNatGatewayDestinationIpPortTranslationNatRule
// This API is used to delete the port forwarding rule of a NAT gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) DeleteNatGatewayDestinationIpPortTranslationNatRuleWithContext(ctx context.Context, request *DeleteNatGatewayDestinationIpPortTranslationNatRuleRequest) (response *DeleteNatGatewayDestinationIpPortTranslationNatRuleResponse, err error) {
    if request == nil {
        request = NewDeleteNatGatewayDestinationIpPortTranslationNatRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNatGatewayDestinationIpPortTranslationNatRule require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to delete a SNAT forwarding rule of a NAT gateway.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NATGATEWAYSNATRULENOTEXISTS = "InvalidParameterValue.NatGatewaySnatRuleNotExists"
func (c *Client) DeleteNatGatewaySourceIpTranslationNatRule(request *DeleteNatGatewaySourceIpTranslationNatRuleRequest) (response *DeleteNatGatewaySourceIpTranslationNatRuleResponse, err error) {
    return c.DeleteNatGatewaySourceIpTranslationNatRuleWithContext(context.Background(), request)
}

// DeleteNatGatewaySourceIpTranslationNatRule
// This API is used to delete a SNAT forwarding rule of a NAT gateway.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NATGATEWAYSNATRULENOTEXISTS = "InvalidParameterValue.NatGatewaySnatRuleNotExists"
func (c *Client) DeleteNatGatewaySourceIpTranslationNatRuleWithContext(ctx context.Context, request *DeleteNatGatewaySourceIpTranslationNatRuleRequest) (response *DeleteNatGatewaySourceIpTranslationNatRuleResponse, err error) {
    if request == nil {
        request = NewDeleteNatGatewaySourceIpTranslationNatRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNatGatewaySourceIpTranslationNatRule require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteNetDetectWithContext(context.Background(), request)
}

// DeleteNetDetect
// This API (DeleteNetDetect) is used to delete a network detection instance.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteNetDetectWithContext(ctx context.Context, request *DeleteNetDetectRequest) (response *DeleteNetDetectResponse, err error) {
    if request == nil {
        request = NewDeleteNetDetectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNetDetect require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteNetworkAclWithContext(context.Background(), request)
}

// DeleteNetworkAcl
// This API is used to delete a network ACL.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteNetworkAclWithContext(ctx context.Context, request *DeleteNetworkAclRequest) (response *DeleteNetworkAclResponse, err error) {
    if request == nil {
        request = NewDeleteNetworkAclRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNetworkAcl require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to delete an ENI.
//
// * An ENI cannot be deleted when it’s bound to a CVM.
//
//  * After the deletion, all of its private IP addresses will be released.
//
// 
//
// This API is completed asynchronously. If you need to query the execution result of an async task, please use the `RequestId` returned by this API to poll the `DescribeVpcTaskResult` API.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) DeleteNetworkInterface(request *DeleteNetworkInterfaceRequest) (response *DeleteNetworkInterfaceResponse, err error) {
    return c.DeleteNetworkInterfaceWithContext(context.Background(), request)
}

// DeleteNetworkInterface
// This API is used to delete an ENI.
//
// * An ENI cannot be deleted when it’s bound to a CVM.
//
//  * After the deletion, all of its private IP addresses will be released.
//
// 
//
// This API is completed asynchronously. If you need to query the execution result of an async task, please use the `RequestId` returned by this API to poll the `DescribeVpcTaskResult` API.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) DeleteNetworkInterfaceWithContext(ctx context.Context, request *DeleteNetworkInterfaceRequest) (response *DeleteNetworkInterfaceResponse, err error) {
    if request == nil {
        request = NewDeleteNetworkInterfaceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNetworkInterface require credential")
    }

    request.SetContext(ctx)
    
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
//  UNSUPPORTEDOPERATION_DELDEFAULTROUTE = "UnsupportedOperation.DelDefaultRoute"
//  UNSUPPORTEDOPERATION_DELROUTEWITHSUBNET = "UnsupportedOperation.DelRouteWithSubnet"
//  UNSUPPORTEDOPERATION_NOTSUPPORTDELETEDEFAULTROUTETABLE = "UnsupportedOperation.NotSupportDeleteDefaultRouteTable"
func (c *Client) DeleteRouteTable(request *DeleteRouteTableRequest) (response *DeleteRouteTableResponse, err error) {
    return c.DeleteRouteTableWithContext(context.Background(), request)
}

// DeleteRouteTable
// This API is used to delete a route table.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_DELDEFAULTROUTE = "UnsupportedOperation.DelDefaultRoute"
//  UNSUPPORTEDOPERATION_DELROUTEWITHSUBNET = "UnsupportedOperation.DelRouteWithSubnet"
//  UNSUPPORTEDOPERATION_NOTSUPPORTDELETEDEFAULTROUTETABLE = "UnsupportedOperation.NotSupportDeleteDefaultRouteTable"
func (c *Client) DeleteRouteTableWithContext(ctx context.Context, request *DeleteRouteTableRequest) (response *DeleteRouteTableResponse, err error) {
    if request == nil {
        request = NewDeleteRouteTableRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRouteTable require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteRoutesWithContext(context.Background(), request)
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
func (c *Client) DeleteRoutesWithContext(ctx context.Context, request *DeleteRoutesRequest) (response *DeleteRoutesResponse, err error) {
    if request == nil {
        request = NewDeleteRoutesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRoutes require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteSecurityGroupWithContext(context.Background(), request)
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
func (c *Client) DeleteSecurityGroupWithContext(ctx context.Context, request *DeleteSecurityGroupRequest) (response *DeleteSecurityGroupResponse, err error) {
    if request == nil {
        request = NewDeleteSecurityGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSecurityGroup require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteSecurityGroupPoliciesWithContext(context.Background(), request)
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
func (c *Client) DeleteSecurityGroupPoliciesWithContext(ctx context.Context, request *DeleteSecurityGroupPoliciesRequest) (response *DeleteSecurityGroupPoliciesResponse, err error) {
    if request == nil {
        request = NewDeleteSecurityGroupPoliciesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSecurityGroupPolicies require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteServiceTemplateWithContext(context.Background(), request)
}

// DeleteServiceTemplate
// This API (DeleteServiceTemplate) is used to delete a protocol port template.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) DeleteServiceTemplateWithContext(ctx context.Context, request *DeleteServiceTemplateRequest) (response *DeleteServiceTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteServiceTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteServiceTemplate require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteServiceTemplateGroupWithContext(context.Background(), request)
}

// DeleteServiceTemplateGroup
// This API (DeleteServiceTemplateGroup) is used to delete a protocol port template group.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) DeleteServiceTemplateGroupWithContext(ctx context.Context, request *DeleteServiceTemplateGroupRequest) (response *DeleteServiceTemplateGroupResponse, err error) {
    if request == nil {
        request = NewDeleteServiceTemplateGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteServiceTemplateGroup require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteSubnetWithContext(context.Background(), request)
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
func (c *Client) DeleteSubnetWithContext(ctx context.Context, request *DeleteSubnetRequest) (response *DeleteSubnetResponse, err error) {
    if request == nil {
        request = NewDeleteSubnetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSubnet require credential")
    }

    request.SetContext(ctx)
    
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
//  UNSUPPORTEDOPERATION_APPIDMISMATCH = "UnsupportedOperation.AppIdMismatch"
func (c *Client) DeleteVpc(request *DeleteVpcRequest) (response *DeleteVpcResponse, err error) {
    return c.DeleteVpcWithContext(context.Background(), request)
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
//  UNSUPPORTEDOPERATION_APPIDMISMATCH = "UnsupportedOperation.AppIdMismatch"
func (c *Client) DeleteVpcWithContext(ctx context.Context, request *DeleteVpcRequest) (response *DeleteVpcResponse, err error) {
    if request == nil {
        request = NewDeleteVpcRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteVpc require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteVpcEndPointWithContext(context.Background(), request)
}

// DeleteVpcEndPoint
// This API is used to delete an endpoint.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteVpcEndPointWithContext(ctx context.Context, request *DeleteVpcEndPointRequest) (response *DeleteVpcEndPointResponse, err error) {
    if request == nil {
        request = NewDeleteVpcEndPointRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteVpcEndPoint require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteVpcEndPointServiceWithContext(context.Background(), request)
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
func (c *Client) DeleteVpcEndPointServiceWithContext(ctx context.Context, request *DeleteVpcEndPointServiceRequest) (response *DeleteVpcEndPointServiceResponse, err error) {
    if request == nil {
        request = NewDeleteVpcEndPointServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteVpcEndPointService require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_UINNOTFOUND = "UnsupportedOperation.UinNotFound"
func (c *Client) DeleteVpcEndPointServiceWhiteList(request *DeleteVpcEndPointServiceWhiteListRequest) (response *DeleteVpcEndPointServiceWhiteListResponse, err error) {
    return c.DeleteVpcEndPointServiceWhiteListWithContext(context.Background(), request)
}

// DeleteVpcEndPointServiceWhiteList
// This API is used to delete the endpoint service allowlist.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_UINNOTFOUND = "UnsupportedOperation.UinNotFound"
func (c *Client) DeleteVpcEndPointServiceWhiteListWithContext(ctx context.Context, request *DeleteVpcEndPointServiceWhiteListRequest) (response *DeleteVpcEndPointServiceWhiteListResponse, err error) {
    if request == nil {
        request = NewDeleteVpcEndPointServiceWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteVpcEndPointServiceWhiteList require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteVpnConnectionWithContext(context.Background(), request)
}

// DeleteVpnConnection
// This API (DeleteVpnConnection) is used to delete VPN tunnels.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) DeleteVpnConnectionWithContext(ctx context.Context, request *DeleteVpnConnectionRequest) (response *DeleteVpnConnectionResponse, err error) {
    if request == nil {
        request = NewDeleteVpnConnectionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteVpnConnection require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteVpnGatewayWithContext(context.Background(), request)
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
func (c *Client) DeleteVpnGatewayWithContext(ctx context.Context, request *DeleteVpnGatewayRequest) (response *DeleteVpnGatewayResponse, err error) {
    if request == nil {
        request = NewDeleteVpnGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteVpnGateway require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteVpnGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteVpnGatewayRoutesRequest() (request *DeleteVpnGatewayRoutesRequest) {
    request = &DeleteVpnGatewayRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteVpnGatewayRoutes")
    
    
    return
}

func NewDeleteVpnGatewayRoutesResponse() (response *DeleteVpnGatewayRoutesResponse) {
    response = &DeleteVpnGatewayRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteVpnGatewayRoutes
// This API is used to delete routes of a VPN gateway.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteVpnGatewayRoutes(request *DeleteVpnGatewayRoutesRequest) (response *DeleteVpnGatewayRoutesResponse, err error) {
    return c.DeleteVpnGatewayRoutesWithContext(context.Background(), request)
}

// DeleteVpnGatewayRoutes
// This API is used to delete routes of a VPN gateway.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteVpnGatewayRoutesWithContext(ctx context.Context, request *DeleteVpnGatewayRoutesRequest) (response *DeleteVpnGatewayRoutesResponse, err error) {
    if request == nil {
        request = NewDeleteVpnGatewayRoutesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteVpnGatewayRoutes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteVpnGatewayRoutesResponse()
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
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAccountAttributes(request *DescribeAccountAttributesRequest) (response *DescribeAccountAttributesResponse, err error) {
    return c.DescribeAccountAttributesWithContext(context.Background(), request)
}

// DescribeAccountAttributes
// This API (DescribeAccountAttributes) is used to query your account attributes.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAccountAttributesWithContext(ctx context.Context, request *DescribeAccountAttributesRequest) (response *DescribeAccountAttributesResponse, err error) {
    if request == nil {
        request = NewDescribeAccountAttributesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccountAttributes require credential")
    }

    request.SetContext(ctx)
    
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAddressQuota(request *DescribeAddressQuotaRequest) (response *DescribeAddressQuotaResponse, err error) {
    return c.DescribeAddressQuotaWithContext(context.Background(), request)
}

// DescribeAddressQuota
// This API (DescribeAddressQuota) is used to query the quota information of your [Elastic IP](https://intl.cloud.tencent.com/document/product/213/1941?from_cn_redirect=1) (EIP) in the current region. For more information, see [EIP product introduction](https://intl.cloud.tencent.com/document/product/213/5733?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAddressQuotaWithContext(ctx context.Context, request *DescribeAddressQuotaRequest) (response *DescribeAddressQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeAddressQuotaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAddressQuota require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
func (c *Client) DescribeAddressTemplateGroups(request *DescribeAddressTemplateGroupsRequest) (response *DescribeAddressTemplateGroupsResponse, err error) {
    return c.DescribeAddressTemplateGroupsWithContext(context.Background(), request)
}

// DescribeAddressTemplateGroups
// This API (DescribeAddressTemplateGroups) is used to query an IP address template group.
//
// error code that may be returned:
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
func (c *Client) DescribeAddressTemplateGroupsWithContext(ctx context.Context, request *DescribeAddressTemplateGroupsRequest) (response *DescribeAddressTemplateGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeAddressTemplateGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAddressTemplateGroups require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
func (c *Client) DescribeAddressTemplates(request *DescribeAddressTemplatesRequest) (response *DescribeAddressTemplatesResponse, err error) {
    return c.DescribeAddressTemplatesWithContext(context.Background(), request)
}

// DescribeAddressTemplates
// This API (DescribeAddressTemplates) is used to query an IP address template.
//
// error code that may be returned:
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
func (c *Client) DescribeAddressTemplatesWithContext(ctx context.Context, request *DescribeAddressTemplatesRequest) (response *DescribeAddressTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeAddressTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAddressTemplates require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETERVALUE_ADDRESSIDMALFORMED = "InvalidParameterValue.AddressIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  LIMITEXCEEDED_NUMBEROFFILTERS = "LimitExceeded.NumberOfFilters"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAddresses(request *DescribeAddressesRequest) (response *DescribeAddressesResponse, err error) {
    return c.DescribeAddressesWithContext(context.Background(), request)
}

// DescribeAddresses
// This API (DescribeAddresses) is used to query the information of one or multiple [Elastic IPs](https://intl.cloud.tencent.com/document/product/213/1941?from_cn_redirect=1).
//
// * If the parameter is empty, a number (as specified by the `Limit`, the default value is 20) of EIPs will be returned.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETERVALUE_ADDRESSIDMALFORMED = "InvalidParameterValue.AddressIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  LIMITEXCEEDED_NUMBEROFFILTERS = "LimitExceeded.NumberOfFilters"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAddressesWithContext(ctx context.Context, request *DescribeAddressesRequest) (response *DescribeAddressesResponse, err error) {
    if request == nil {
        request = NewDescribeAddressesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAddresses require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
func (c *Client) DescribeAssistantCidr(request *DescribeAssistantCidrRequest) (response *DescribeAssistantCidrResponse, err error) {
    return c.DescribeAssistantCidrWithContext(context.Background(), request)
}

// DescribeAssistantCidr
// This API (DescribeAssistantCidr) is used to query a list of secondary CIDR blocks. (To use this API that is in Beta, please submit a ticket.)
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
func (c *Client) DescribeAssistantCidrWithContext(ctx context.Context, request *DescribeAssistantCidrRequest) (response *DescribeAssistantCidrResponse, err error) {
    if request == nil {
        request = NewDescribeAssistantCidrRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssistantCidr require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGEIDMALFORMED = "InvalidParameterValue.BandwidthPackageIdMalformed"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGENOTFOUND = "InvalidParameterValue.BandwidthPackageNotFound"
//  UNSUPPORTEDOPERATION_BANDWIDTHPACKAGEIDNOTSUPPORTED = "UnsupportedOperation.BandwidthPackageIdNotSupported"
func (c *Client) DescribeBandwidthPackageBillUsage(request *DescribeBandwidthPackageBillUsageRequest) (response *DescribeBandwidthPackageBillUsageResponse, err error) {
    return c.DescribeBandwidthPackageBillUsageWithContext(context.Background(), request)
}

// DescribeBandwidthPackageBillUsage
// This API is used to query the current billable usage of a pay-as-you-go bandwidth package.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGEIDMALFORMED = "InvalidParameterValue.BandwidthPackageIdMalformed"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGENOTFOUND = "InvalidParameterValue.BandwidthPackageNotFound"
//  UNSUPPORTEDOPERATION_BANDWIDTHPACKAGEIDNOTSUPPORTED = "UnsupportedOperation.BandwidthPackageIdNotSupported"
func (c *Client) DescribeBandwidthPackageBillUsageWithContext(ctx context.Context, request *DescribeBandwidthPackageBillUsageRequest) (response *DescribeBandwidthPackageBillUsageResponse, err error) {
    if request == nil {
        request = NewDescribeBandwidthPackageBillUsageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBandwidthPackageBillUsage require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_RESOURCENOTSUPPORT = "InvalidParameterValue.ResourceNotSupport"
func (c *Client) DescribeBandwidthPackageQuota(request *DescribeBandwidthPackageQuotaRequest) (response *DescribeBandwidthPackageQuotaResponse, err error) {
    return c.DescribeBandwidthPackageQuotaWithContext(context.Background(), request)
}

// DescribeBandwidthPackageQuota
// This API is used to query the maximum and used number of bandwidth packages under the account in the current region.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_RESOURCENOTSUPPORT = "InvalidParameterValue.ResourceNotSupport"
func (c *Client) DescribeBandwidthPackageQuotaWithContext(ctx context.Context, request *DescribeBandwidthPackageQuotaRequest) (response *DescribeBandwidthPackageQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeBandwidthPackageQuotaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBandwidthPackageQuota require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGEIDMALFORMED = "InvalidParameterValue.BandwidthPackageIdMalformed"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGENOTFOUND = "InvalidParameterValue.BandwidthPackageNotFound"
//  INVALIDPARAMETERVALUE_RESOURCEIDMALFORMED = "InvalidParameterValue.ResourceIdMalformed"
func (c *Client) DescribeBandwidthPackageResources(request *DescribeBandwidthPackageResourcesRequest) (response *DescribeBandwidthPackageResourcesResponse, err error) {
    return c.DescribeBandwidthPackageResourcesWithContext(context.Background(), request)
}

// DescribeBandwidthPackageResources
// This API is used to query resources in a bandwidth package based on the unique package ID. You can filter the result by specifying conditions and paginate the query results.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGEIDMALFORMED = "InvalidParameterValue.BandwidthPackageIdMalformed"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGENOTFOUND = "InvalidParameterValue.BandwidthPackageNotFound"
//  INVALIDPARAMETERVALUE_RESOURCEIDMALFORMED = "InvalidParameterValue.ResourceIdMalformed"
func (c *Client) DescribeBandwidthPackageResourcesWithContext(ctx context.Context, request *DescribeBandwidthPackageResourcesRequest) (response *DescribeBandwidthPackageResourcesResponse, err error) {
    if request == nil {
        request = NewDescribeBandwidthPackageResourcesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBandwidthPackageResources require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGEIDMALFORMED = "InvalidParameterValue.BandwidthPackageIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTHPACKAGECHARGETYPE = "InvalidParameterValue.InvalidBandwidthPackageChargeType"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBandwidthPackages(request *DescribeBandwidthPackagesRequest) (response *DescribeBandwidthPackagesResponse, err error) {
    return c.DescribeBandwidthPackagesWithContext(context.Background(), request)
}

// DescribeBandwidthPackages
// This API is used to query bandwidth package information, including the unique ID of the bandwidth package, the type, the billing mode, the name, and the resource information.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGEIDMALFORMED = "InvalidParameterValue.BandwidthPackageIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTHPACKAGECHARGETYPE = "InvalidParameterValue.InvalidBandwidthPackageChargeType"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBandwidthPackagesWithContext(ctx context.Context, request *DescribeBandwidthPackagesRequest) (response *DescribeBandwidthPackagesResponse, err error) {
    if request == nil {
        request = NewDescribeBandwidthPackagesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBandwidthPackages require credential")
    }

    request.SetContext(ctx)
    
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_APPIDNOTFOUND = "UnsupportedOperation.AppIdNotFound"
func (c *Client) DescribeCcnAttachedInstances(request *DescribeCcnAttachedInstancesRequest) (response *DescribeCcnAttachedInstancesResponse, err error) {
    return c.DescribeCcnAttachedInstancesWithContext(context.Background(), request)
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_APPIDNOTFOUND = "UnsupportedOperation.AppIdNotFound"
func (c *Client) DescribeCcnAttachedInstancesWithContext(ctx context.Context, request *DescribeCcnAttachedInstancesRequest) (response *DescribeCcnAttachedInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeCcnAttachedInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCcnAttachedInstances require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeCcnRegionBandwidthLimitsWithContext(context.Background(), request)
}

// DescribeCcnRegionBandwidthLimits
// This API is used to query the outbound bandwidth caps of all regions connected with a CCN instance. The API only returns regions included in the associated network instances.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
func (c *Client) DescribeCcnRegionBandwidthLimitsWithContext(ctx context.Context, request *DescribeCcnRegionBandwidthLimitsRequest) (response *DescribeCcnRegionBandwidthLimitsResponse, err error) {
    if request == nil {
        request = NewDescribeCcnRegionBandwidthLimitsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCcnRegionBandwidthLimits require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeCcnRoutesWithContext(context.Background(), request)
}

// DescribeCcnRoutes
// This API (DescribeCcnRoutes) is used to query routes that have been added to a CCN.
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCcnRoutesWithContext(ctx context.Context, request *DescribeCcnRoutesRequest) (response *DescribeCcnRoutesResponse, err error) {
    if request == nil {
        request = NewDescribeCcnRoutesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCcnRoutes require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETER_FILTERNOTDICT = "InvalidParameter.FilterNotDict"
//  INVALIDPARAMETER_FILTERVALUESNOTLIST = "InvalidParameter.FilterValuesNotList"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCcns(request *DescribeCcnsRequest) (response *DescribeCcnsResponse, err error) {
    return c.DescribeCcnsWithContext(context.Background(), request)
}

// DescribeCcns
// This API (DescribeCcns) is used to query the CCN list.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETER_FILTERNOTDICT = "InvalidParameter.FilterNotDict"
//  INVALIDPARAMETER_FILTERVALUESNOTLIST = "InvalidParameter.FilterValuesNotList"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCcnsWithContext(ctx context.Context, request *DescribeCcnsRequest) (response *DescribeCcnsResponse, err error) {
    if request == nil {
        request = NewDescribeCcnsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCcns require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeClassicLinkInstancesWithContext(context.Background(), request)
}

// DescribeClassicLinkInstances
// This API (DescribeClassicLinkInstances) is used to query the Classiclink instances list.
//
// error code that may be returned:
//  INVALIDPARAMETER_FILTERVALUESNOTLIST = "InvalidParameter.FilterValuesNotList"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
func (c *Client) DescribeClassicLinkInstancesWithContext(ctx context.Context, request *DescribeClassicLinkInstancesRequest) (response *DescribeClassicLinkInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeClassicLinkInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClassicLinkInstances require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
func (c *Client) DescribeCrossBorderCompliance(request *DescribeCrossBorderComplianceRequest) (response *DescribeCrossBorderComplianceResponse, err error) {
    return c.DescribeCrossBorderComplianceWithContext(context.Background(), request)
}

// DescribeCrossBorderCompliance
// This API is used to query the compliance review requests created by the user. 
//
// A service provider can query all review requests created by any `APPID` under its account. Other users can only query their own review requests.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
func (c *Client) DescribeCrossBorderComplianceWithContext(ctx context.Context, request *DescribeCrossBorderComplianceRequest) (response *DescribeCrossBorderComplianceResponse, err error) {
    if request == nil {
        request = NewDescribeCrossBorderComplianceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCrossBorderCompliance require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
func (c *Client) DescribeCustomerGatewayVendors(request *DescribeCustomerGatewayVendorsRequest) (response *DescribeCustomerGatewayVendorsResponse, err error) {
    return c.DescribeCustomerGatewayVendorsWithContext(context.Background(), request)
}

// DescribeCustomerGatewayVendors
// This API (DescribeCustomerGatewayVendors) is used to query the information of supported customer gateway vendors.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
func (c *Client) DescribeCustomerGatewayVendorsWithContext(ctx context.Context, request *DescribeCustomerGatewayVendorsRequest) (response *DescribeCustomerGatewayVendorsResponse, err error) {
    if request == nil {
        request = NewDescribeCustomerGatewayVendorsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCustomerGatewayVendors require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeCustomerGatewaysWithContext(context.Background(), request)
}

// DescribeCustomerGateways
// This API (DescribeCustomerGateways) is used to query the customer gateway list.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCustomerGatewaysWithContext(ctx context.Context, request *DescribeCustomerGatewaysRequest) (response *DescribeCustomerGatewaysResponse, err error) {
    if request == nil {
        request = NewDescribeCustomerGatewaysRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCustomerGateways require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeDirectConnectGatewayCcnRoutesWithContext(context.Background(), request)
}

// DescribeDirectConnectGatewayCcnRoutes
// This API (DescribeDirectConnectGatewayCcnRoutes) is used to query the CCN routes (IDC IP range) of the Direct Connect gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDirectConnectGatewayCcnRoutesWithContext(ctx context.Context, request *DescribeDirectConnectGatewayCcnRoutesRequest) (response *DescribeDirectConnectGatewayCcnRoutesResponse, err error) {
    if request == nil {
        request = NewDescribeDirectConnectGatewayCcnRoutesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDirectConnectGatewayCcnRoutes require credential")
    }

    request.SetContext(ctx)
    
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDirectConnectGateways(request *DescribeDirectConnectGatewaysRequest) (response *DescribeDirectConnectGatewaysResponse, err error) {
    return c.DescribeDirectConnectGatewaysWithContext(context.Background(), request)
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDirectConnectGatewaysWithContext(ctx context.Context, request *DescribeDirectConnectGatewaysRequest) (response *DescribeDirectConnectGatewaysResponse, err error) {
    if request == nil {
        request = NewDescribeDirectConnectGatewaysRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDirectConnectGateways require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeFlowLogWithContext(context.Background(), request)
}

// DescribeFlowLog
// This API is used to query the information of a flow log.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
func (c *Client) DescribeFlowLogWithContext(ctx context.Context, request *DescribeFlowLogRequest) (response *DescribeFlowLogResponse, err error) {
    if request == nil {
        request = NewDescribeFlowLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFlowLog require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeFlowLogsWithContext(context.Background(), request)
}

// DescribeFlowLogs
// This API is used to query all the flow logs of the current account.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeFlowLogsWithContext(ctx context.Context, request *DescribeFlowLogsRequest) (response *DescribeFlowLogsResponse, err error) {
    if request == nil {
        request = NewDescribeFlowLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFlowLogs require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeGatewayFlowMonitorDetailWithContext(context.Background(), request)
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
func (c *Client) DescribeGatewayFlowMonitorDetailWithContext(ctx context.Context, request *DescribeGatewayFlowMonitorDetailRequest) (response *DescribeGatewayFlowMonitorDetailResponse, err error) {
    if request == nil {
        request = NewDescribeGatewayFlowMonitorDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGatewayFlowMonitorDetail require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to query the inbound IP bandwidth limit of a gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) DescribeGatewayFlowQos(request *DescribeGatewayFlowQosRequest) (response *DescribeGatewayFlowQosResponse, err error) {
    return c.DescribeGatewayFlowQosWithContext(context.Background(), request)
}

// DescribeGatewayFlowQos
// This API is used to query the inbound IP bandwidth limit of a gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) DescribeGatewayFlowQosWithContext(ctx context.Context, request *DescribeGatewayFlowQosRequest) (response *DescribeGatewayFlowQosResponse, err error) {
    if request == nil {
        request = NewDescribeGatewayFlowQosRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGatewayFlowQos require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
func (c *Client) DescribeHaVips(request *DescribeHaVipsRequest) (response *DescribeHaVipsResponse, err error) {
    return c.DescribeHaVipsWithContext(context.Background(), request)
}

// DescribeHaVips
// This API (DescribeHaVips) is used to query the list of highly available virtual IPs (HAVIP).
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
func (c *Client) DescribeHaVipsWithContext(ctx context.Context, request *DescribeHaVipsRequest) (response *DescribeHaVipsResponse, err error) {
    if request == nil {
        request = NewDescribeHaVipsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHaVips require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDACCOUNT_NOTSUPPORTED = "InvalidAccount.NotSupported"
//  INVALIDPARAMETERCONFLICT = "InvalidParameterConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeIpGeolocationDatabaseUrl(request *DescribeIpGeolocationDatabaseUrlRequest) (response *DescribeIpGeolocationDatabaseUrlResponse, err error) {
    return c.DescribeIpGeolocationDatabaseUrlWithContext(context.Background(), request)
}

// DescribeIpGeolocationDatabaseUrl
// This API is used to obtain the download link of an IP location database.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDACCOUNT_NOTSUPPORTED = "InvalidAccount.NotSupported"
//  INVALIDPARAMETERCONFLICT = "InvalidParameterConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeIpGeolocationDatabaseUrlWithContext(ctx context.Context, request *DescribeIpGeolocationDatabaseUrlRequest) (response *DescribeIpGeolocationDatabaseUrlResponse, err error) {
    if request == nil {
        request = NewDescribeIpGeolocationDatabaseUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIpGeolocationDatabaseUrl require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to query the location and network information of one or more IP addresses.
//
// This API is currently in beta test. To use it, please [submit a ticket](https://console.cloud.tencent.com/workorder/category?level1_id=6&level2_id=660&source=0&data_title=%E5%BC%B9%E6%80%A7%E5%85%AC%E7%BD%91%20EIP&level3_id=662&queue=96&scene_code=16400&step=2).
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDACCOUNT_NOTSUPPORTED = "InvalidAccount.NotSupported"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_COMBINATION = "InvalidParameterValue.Combination"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeIpGeolocationInfos(request *DescribeIpGeolocationInfosRequest) (response *DescribeIpGeolocationInfosResponse, err error) {
    return c.DescribeIpGeolocationInfosWithContext(context.Background(), request)
}

// DescribeIpGeolocationInfos
// This API is used to query the location and network information of one or more IP addresses.
//
// This API is currently in beta test. To use it, please [submit a ticket](https://console.cloud.tencent.com/workorder/category?level1_id=6&level2_id=660&source=0&data_title=%E5%BC%B9%E6%80%A7%E5%85%AC%E7%BD%91%20EIP&level3_id=662&queue=96&scene_code=16400&step=2).
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDACCOUNT_NOTSUPPORTED = "InvalidAccount.NotSupported"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_COMBINATION = "InvalidParameterValue.Combination"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeIpGeolocationInfosWithContext(ctx context.Context, request *DescribeIpGeolocationInfosRequest) (response *DescribeIpGeolocationInfosResponse, err error) {
    if request == nil {
        request = NewDescribeIpGeolocationInfosRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIpGeolocationInfos require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeLocalGatewayWithContext(context.Background(), request)
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
func (c *Client) DescribeLocalGatewayWithContext(ctx context.Context, request *DescribeLocalGatewayRequest) (response *DescribeLocalGatewayResponse, err error) {
    if request == nil {
        request = NewDescribeLocalGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLocalGateway require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to query the array of objects of a NAT gateway's port forwarding rules.
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
    return c.DescribeNatGatewayDestinationIpPortTranslationNatRulesWithContext(context.Background(), request)
}

// DescribeNatGatewayDestinationIpPortTranslationNatRules
// This API is used to query the array of objects of a NAT gateway's port forwarding rules.
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
func (c *Client) DescribeNatGatewayDestinationIpPortTranslationNatRulesWithContext(ctx context.Context, request *DescribeNatGatewayDestinationIpPortTranslationNatRulesRequest) (response *DescribeNatGatewayDestinationIpPortTranslationNatRulesResponse, err error) {
    if request == nil {
        request = NewDescribeNatGatewayDestinationIpPortTranslationNatRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNatGatewayDestinationIpPortTranslationNatRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNatGatewayDestinationIpPortTranslationNatRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNatGatewayDirectConnectGatewayRouteRequest() (request *DescribeNatGatewayDirectConnectGatewayRouteRequest) {
    request = &DescribeNatGatewayDirectConnectGatewayRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeNatGatewayDirectConnectGatewayRoute")
    
    
    return
}

func NewDescribeNatGatewayDirectConnectGatewayRouteResponse() (response *DescribeNatGatewayDirectConnectGatewayRouteResponse) {
    response = &DescribeNatGatewayDirectConnectGatewayRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNatGatewayDirectConnectGatewayRoute
// This API is used to query the routes between a NAT gateway and Direct Connect.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeNatGatewayDirectConnectGatewayRoute(request *DescribeNatGatewayDirectConnectGatewayRouteRequest) (response *DescribeNatGatewayDirectConnectGatewayRouteResponse, err error) {
    return c.DescribeNatGatewayDirectConnectGatewayRouteWithContext(context.Background(), request)
}

// DescribeNatGatewayDirectConnectGatewayRoute
// This API is used to query the routes between a NAT gateway and Direct Connect.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeNatGatewayDirectConnectGatewayRouteWithContext(ctx context.Context, request *DescribeNatGatewayDirectConnectGatewayRouteRequest) (response *DescribeNatGatewayDirectConnectGatewayRouteResponse, err error) {
    if request == nil {
        request = NewDescribeNatGatewayDirectConnectGatewayRouteRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNatGatewayDirectConnectGatewayRoute require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNatGatewayDirectConnectGatewayRouteResponse()
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
// This API is used to query the NAT gateway's SNAT forwarding rules.
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
    return c.DescribeNatGatewaySourceIpTranslationNatRulesWithContext(context.Background(), request)
}

// DescribeNatGatewaySourceIpTranslationNatRules
// This API is used to query the NAT gateway's SNAT forwarding rules.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDADDRESSID_NOTFOUND = "InvalidAddressId.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNatGatewaySourceIpTranslationNatRulesWithContext(ctx context.Context, request *DescribeNatGatewaySourceIpTranslationNatRulesRequest) (response *DescribeNatGatewaySourceIpTranslationNatRulesResponse, err error) {
    if request == nil {
        request = NewDescribeNatGatewaySourceIpTranslationNatRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNatGatewaySourceIpTranslationNatRules require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to query NAT gateways.
//
// error code that may be returned:
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNatGateways(request *DescribeNatGatewaysRequest) (response *DescribeNatGatewaysResponse, err error) {
    return c.DescribeNatGatewaysWithContext(context.Background(), request)
}

// DescribeNatGateways
// This API is used to query NAT gateways.
//
// error code that may be returned:
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNatGatewaysWithContext(ctx context.Context, request *DescribeNatGatewaysRequest) (response *DescribeNatGatewaysResponse, err error) {
    if request == nil {
        request = NewDescribeNatGatewaysRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNatGateways require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION_NETDETECTTIMEOUT = "FailedOperation.NetDetectTimeOut"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_RESOURCENOTFOUND = "InvalidParameterValue.ResourceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNetDetectStates(request *DescribeNetDetectStatesRequest) (response *DescribeNetDetectStatesResponse, err error) {
    return c.DescribeNetDetectStatesWithContext(context.Background(), request)
}

// DescribeNetDetectStates
// This API (DescribeNetDetectStates) is used to query the list of network detection verification results.
//
// error code that may be returned:
//  FAILEDOPERATION_NETDETECTTIMEOUT = "FailedOperation.NetDetectTimeOut"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_RESOURCENOTFOUND = "InvalidParameterValue.ResourceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNetDetectStatesWithContext(ctx context.Context, request *DescribeNetDetectStatesRequest) (response *DescribeNetDetectStatesResponse, err error) {
    if request == nil {
        request = NewDescribeNetDetectStatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNetDetectStates require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeNetDetectsWithContext(context.Background(), request)
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
func (c *Client) DescribeNetDetectsWithContext(ctx context.Context, request *DescribeNetDetectsRequest) (response *DescribeNetDetectsResponse, err error) {
    if request == nil {
        request = NewDescribeNetDetectsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNetDetects require credential")
    }

    request.SetContext(ctx)
    
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ACTIONNOTFOUND = "UnsupportedOperation.ActionNotFound"
func (c *Client) DescribeNetworkAcls(request *DescribeNetworkAclsRequest) (response *DescribeNetworkAclsResponse, err error) {
    return c.DescribeNetworkAclsWithContext(context.Background(), request)
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ACTIONNOTFOUND = "UnsupportedOperation.ActionNotFound"
func (c *Client) DescribeNetworkAclsWithContext(ctx context.Context, request *DescribeNetworkAclsRequest) (response *DescribeNetworkAclsResponse, err error) {
    if request == nil {
        request = NewDescribeNetworkAclsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNetworkAcls require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNetworkInterfaceLimit(request *DescribeNetworkInterfaceLimitRequest) (response *DescribeNetworkInterfaceLimitResponse, err error) {
    return c.DescribeNetworkInterfaceLimitWithContext(context.Background(), request)
}

// DescribeNetworkInterfaceLimit
// This API (DescribeNetworkInterfaceLimit) is used to query the ENI quota based on the ID of CVM instance or ENI. It returns the ENI quota to which the CVM instance can be bound and the IP address quota that can be allocated to the ENI.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNetworkInterfaceLimitWithContext(ctx context.Context, request *DescribeNetworkInterfaceLimitRequest) (response *DescribeNetworkInterfaceLimitResponse, err error) {
    if request == nil {
        request = NewDescribeNetworkInterfaceLimitRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNetworkInterfaceLimit require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETER_FILTERVALUESNOTLIST = "InvalidParameter.FilterValuesNotList"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNetworkInterfaces(request *DescribeNetworkInterfacesRequest) (response *DescribeNetworkInterfacesResponse, err error) {
    return c.DescribeNetworkInterfacesWithContext(context.Background(), request)
}

// DescribeNetworkInterfaces
// This API (DescribeNetworkInterfaces) is used to query the ENI list.
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNetworkInterfacesWithContext(ctx context.Context, request *DescribeNetworkInterfacesRequest) (response *DescribeNetworkInterfacesResponse, err error) {
    if request == nil {
        request = NewDescribeNetworkInterfacesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNetworkInterfaces require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETER_FILTERNOTDICT = "InvalidParameter.FilterNotDict"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRouteTables(request *DescribeRouteTablesRequest) (response *DescribeRouteTablesResponse, err error) {
    return c.DescribeRouteTablesWithContext(context.Background(), request)
}

// DescribeRouteTables
//  This API (DescribeRouteTables) is used to query route tables.
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETER_FILTERNOTDICT = "InvalidParameter.FilterNotDict"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRouteTablesWithContext(ctx context.Context, request *DescribeRouteTablesRequest) (response *DescribeRouteTablesResponse, err error) {
    if request == nil {
        request = NewDescribeRouteTablesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRouteTables require credential")
    }

    request.SetContext(ctx)
    
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSecurityGroupAssociationStatistics(request *DescribeSecurityGroupAssociationStatisticsRequest) (response *DescribeSecurityGroupAssociationStatisticsResponse, err error) {
    return c.DescribeSecurityGroupAssociationStatisticsWithContext(context.Background(), request)
}

// DescribeSecurityGroupAssociationStatistics
// This API (DescribeSecurityGroupAssociationStatistics) is used to query statistics on the instances associated with a security group.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSecurityGroupAssociationStatisticsWithContext(ctx context.Context, request *DescribeSecurityGroupAssociationStatisticsRequest) (response *DescribeSecurityGroupAssociationStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityGroupAssociationStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityGroupAssociationStatistics require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSecurityGroupPolicies(request *DescribeSecurityGroupPoliciesRequest) (response *DescribeSecurityGroupPoliciesResponse, err error) {
    return c.DescribeSecurityGroupPoliciesWithContext(context.Background(), request)
}

// DescribeSecurityGroupPolicies
// This API (DescribeSecurityGroupPolicies) is used to query security group policies.
//
// error code that may be returned:
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSecurityGroupPoliciesWithContext(ctx context.Context, request *DescribeSecurityGroupPoliciesRequest) (response *DescribeSecurityGroupPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityGroupPoliciesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityGroupPolicies require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeSecurityGroupReferencesWithContext(context.Background(), request)
}

// DescribeSecurityGroupReferences
// This API (DescribeSecurityGroupReferences) is used to query referred security groups.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSecurityGroupReferencesWithContext(ctx context.Context, request *DescribeSecurityGroupReferencesRequest) (response *DescribeSecurityGroupReferencesResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityGroupReferencesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityGroupReferences require credential")
    }

    request.SetContext(ctx)
    
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSecurityGroups(request *DescribeSecurityGroupsRequest) (response *DescribeSecurityGroupsResponse, err error) {
    return c.DescribeSecurityGroupsWithContext(context.Background(), request)
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSecurityGroupsWithContext(ctx context.Context, request *DescribeSecurityGroupsRequest) (response *DescribeSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
func (c *Client) DescribeServiceTemplateGroups(request *DescribeServiceTemplateGroupsRequest) (response *DescribeServiceTemplateGroupsResponse, err error) {
    return c.DescribeServiceTemplateGroupsWithContext(context.Background(), request)
}

// DescribeServiceTemplateGroups
// This API (DescribeServiceTemplateGroups) is used to query a protocol port template group.
//
// error code that may be returned:
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
func (c *Client) DescribeServiceTemplateGroupsWithContext(ctx context.Context, request *DescribeServiceTemplateGroupsRequest) (response *DescribeServiceTemplateGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeServiceTemplateGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServiceTemplateGroups require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeServiceTemplatesWithContext(context.Background(), request)
}

// DescribeServiceTemplates
// This API (DescribeServiceTemplates) is used to query protocol port templates.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
func (c *Client) DescribeServiceTemplatesWithContext(ctx context.Context, request *DescribeServiceTemplatesRequest) (response *DescribeServiceTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeServiceTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServiceTemplates require credential")
    }

    request.SetContext(ctx)
    
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_APPIDMISMATCH = "UnsupportedOperation.AppIdMismatch"
func (c *Client) DescribeSubnets(request *DescribeSubnetsRequest) (response *DescribeSubnetsResponse, err error) {
    return c.DescribeSubnetsWithContext(context.Background(), request)
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_APPIDMISMATCH = "UnsupportedOperation.AppIdMismatch"
func (c *Client) DescribeSubnetsWithContext(ctx context.Context, request *DescribeSubnetsRequest) (response *DescribeSubnetsResponse, err error) {
    if request == nil {
        request = NewDescribeSubnetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSubnets require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeTaskResultWithContext(context.Background(), request)
}

// DescribeTaskResult
// This API is used to query the EIP async job execution results.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeTaskResultWithContext(ctx context.Context, request *DescribeTaskResultRequest) (response *DescribeTaskResultResponse, err error) {
    if request == nil {
        request = NewDescribeTaskResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskResult require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SVCNOTEXIST = "ResourceNotFound.SvcNotExist"
func (c *Client) DescribeVpcEndPoint(request *DescribeVpcEndPointRequest) (response *DescribeVpcEndPointResponse, err error) {
    return c.DescribeVpcEndPointWithContext(context.Background(), request)
}

// DescribeVpcEndPoint
// This API is used to query the endpoint list.
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SVCNOTEXIST = "ResourceNotFound.SvcNotExist"
func (c *Client) DescribeVpcEndPointWithContext(ctx context.Context, request *DescribeVpcEndPointRequest) (response *DescribeVpcEndPointResponse, err error) {
    if request == nil {
        request = NewDescribeVpcEndPointRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVpcEndPoint require credential")
    }

    request.SetContext(ctx)
    
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
//  UNSUPPORTEDOPERATION_ROLENOTFOUND = "UnsupportedOperation.RoleNotFound"
func (c *Client) DescribeVpcEndPointService(request *DescribeVpcEndPointServiceRequest) (response *DescribeVpcEndPointServiceResponse, err error) {
    return c.DescribeVpcEndPointServiceWithContext(context.Background(), request)
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
//  UNSUPPORTEDOPERATION_ROLENOTFOUND = "UnsupportedOperation.RoleNotFound"
func (c *Client) DescribeVpcEndPointServiceWithContext(ctx context.Context, request *DescribeVpcEndPointServiceRequest) (response *DescribeVpcEndPointServiceResponse, err error) {
    if request == nil {
        request = NewDescribeVpcEndPointServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVpcEndPointService require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  UNSUPPORTEDOPERATION_UINNOTFOUND = "UnsupportedOperation.UinNotFound"
func (c *Client) DescribeVpcEndPointServiceWhiteList(request *DescribeVpcEndPointServiceWhiteListRequest) (response *DescribeVpcEndPointServiceWhiteListResponse, err error) {
    return c.DescribeVpcEndPointServiceWhiteListWithContext(context.Background(), request)
}

// DescribeVpcEndPointServiceWhiteList
// This API is used to query the endpoint service allowlist.
//
// error code that may be returned:
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  UNSUPPORTEDOPERATION_UINNOTFOUND = "UnsupportedOperation.UinNotFound"
func (c *Client) DescribeVpcEndPointServiceWhiteListWithContext(ctx context.Context, request *DescribeVpcEndPointServiceWhiteListRequest) (response *DescribeVpcEndPointServiceWhiteListResponse, err error) {
    if request == nil {
        request = NewDescribeVpcEndPointServiceWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVpcEndPointServiceWhiteList require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeVpcInstancesWithContext(context.Background(), request)
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
func (c *Client) DescribeVpcInstancesWithContext(ctx context.Context, request *DescribeVpcInstancesRequest) (response *DescribeVpcInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeVpcInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVpcInstances require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeVpcIpv6AddressesWithContext(context.Background(), request)
}

// DescribeVpcIpv6Addresses
// This API (DescribeVpcIpv6Addresses) is used to query `VPC` `IPv6` information.
//
// This API is used to query only the information of `IPv6` addresses that are already in use. When querying IPs that have not yet been used, this API will not report an error, but the IPs will not appear in the returned results.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVpcIpv6AddressesWithContext(ctx context.Context, request *DescribeVpcIpv6AddressesRequest) (response *DescribeVpcIpv6AddressesResponse, err error) {
    if request == nil {
        request = NewDescribeVpcIpv6AddressesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVpcIpv6Addresses require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeVpcPrivateIpAddressesWithContext(context.Background(), request)
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
func (c *Client) DescribeVpcPrivateIpAddressesWithContext(ctx context.Context, request *DescribeVpcPrivateIpAddressesRequest) (response *DescribeVpcPrivateIpAddressesResponse, err error) {
    if request == nil {
        request = NewDescribeVpcPrivateIpAddressesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVpcPrivateIpAddresses require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVpcResourceDashboard(request *DescribeVpcResourceDashboardRequest) (response *DescribeVpcResourceDashboardResponse, err error) {
    return c.DescribeVpcResourceDashboardWithContext(context.Background(), request)
}

// DescribeVpcResourceDashboard
// View VPC resources.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVpcResourceDashboardWithContext(ctx context.Context, request *DescribeVpcResourceDashboardRequest) (response *DescribeVpcResourceDashboardResponse, err error) {
    if request == nil {
        request = NewDescribeVpcResourceDashboardRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVpcResourceDashboard require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVpcResourceDashboardResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVpcTaskResultRequest() (request *DescribeVpcTaskResultRequest) {
    request = &DescribeVpcTaskResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpcTaskResult")
    
    
    return
}

func NewDescribeVpcTaskResultResponse() (response *DescribeVpcTaskResultResponse) {
    response = &DescribeVpcTaskResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVpcTaskResult
// This API is used to query the execution result of a VPC task.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeVpcTaskResult(request *DescribeVpcTaskResultRequest) (response *DescribeVpcTaskResultResponse, err error) {
    return c.DescribeVpcTaskResultWithContext(context.Background(), request)
}

// DescribeVpcTaskResult
// This API is used to query the execution result of a VPC task.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeVpcTaskResultWithContext(ctx context.Context, request *DescribeVpcTaskResultRequest) (response *DescribeVpcTaskResultResponse, err error) {
    if request == nil {
        request = NewDescribeVpcTaskResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVpcTaskResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVpcTaskResultResponse()
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
//  INVALIDPARAMETER_FILTERNOTDICT = "InvalidParameter.FilterNotDict"
//  INVALIDPARAMETER_FILTERVALUESNOTLIST = "InvalidParameter.FilterValuesNotList"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeVpcs(request *DescribeVpcsRequest) (response *DescribeVpcsResponse, err error) {
    return c.DescribeVpcsWithContext(context.Background(), request)
}

// DescribeVpcs
// This API (DescribeVpcs) is used to query the VPC list.
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeVpcsWithContext(ctx context.Context, request *DescribeVpcsRequest) (response *DescribeVpcsResponse, err error) {
    if request == nil {
        request = NewDescribeVpcsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVpcs require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeVpnConnections(request *DescribeVpnConnectionsRequest) (response *DescribeVpnConnectionsResponse, err error) {
    return c.DescribeVpnConnectionsWithContext(context.Background(), request)
}

// DescribeVpnConnections
//  This API (DescribeVpnConnections) is used to query the VPN tunnel list.
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeVpnConnectionsWithContext(ctx context.Context, request *DescribeVpnConnectionsRequest) (response *DescribeVpnConnectionsResponse, err error) {
    if request == nil {
        request = NewDescribeVpnConnectionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVpnConnections require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVpnGatewayCcnRoutes(request *DescribeVpnGatewayCcnRoutesRequest) (response *DescribeVpnGatewayCcnRoutesResponse, err error) {
    return c.DescribeVpnGatewayCcnRoutesWithContext(context.Background(), request)
}

// DescribeVpnGatewayCcnRoutes
// This API (DescribeVpnGatewayCcnRoutes) is used to query VPN gateway-based CCN routes.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVpnGatewayCcnRoutesWithContext(ctx context.Context, request *DescribeVpnGatewayCcnRoutesRequest) (response *DescribeVpnGatewayCcnRoutesResponse, err error) {
    if request == nil {
        request = NewDescribeVpnGatewayCcnRoutesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVpnGatewayCcnRoutes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVpnGatewayCcnRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVpnGatewayRoutesRequest() (request *DescribeVpnGatewayRoutesRequest) {
    request = &DescribeVpnGatewayRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpnGatewayRoutes")
    
    
    return
}

func NewDescribeVpnGatewayRoutesResponse() (response *DescribeVpnGatewayRoutesResponse) {
    response = &DescribeVpnGatewayRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVpnGatewayRoutes
// This API is used to query destination routes of a route-based VPN gateway.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeVpnGatewayRoutes(request *DescribeVpnGatewayRoutesRequest) (response *DescribeVpnGatewayRoutesResponse, err error) {
    return c.DescribeVpnGatewayRoutesWithContext(context.Background(), request)
}

// DescribeVpnGatewayRoutes
// This API is used to query destination routes of a route-based VPN gateway.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeVpnGatewayRoutesWithContext(ctx context.Context, request *DescribeVpnGatewayRoutesRequest) (response *DescribeVpnGatewayRoutesResponse, err error) {
    if request == nil {
        request = NewDescribeVpnGatewayRoutesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVpnGatewayRoutes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVpnGatewayRoutesResponse()
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
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDVPNGATEWAYID_MALFORMED = "InvalidVpnGatewayId.Malformed"
//  INVALIDVPNGATEWAYID_NOTFOUND = "InvalidVpnGatewayId.NotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeVpnGateways(request *DescribeVpnGatewaysRequest) (response *DescribeVpnGatewaysResponse, err error) {
    return c.DescribeVpnGatewaysWithContext(context.Background(), request)
}

// DescribeVpnGateways
// This API (DescribeVpnGateways) is used to query the VPN gateway list.
//
// error code that may be returned:
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETER_FILTERNOTDICT = "InvalidParameter.FilterNotDict"
//  INVALIDPARAMETER_FILTERVALUESNOTLIST = "InvalidParameter.FilterValuesNotList"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDVPNGATEWAYID_MALFORMED = "InvalidVpnGatewayId.Malformed"
//  INVALIDVPNGATEWAYID_NOTFOUND = "InvalidVpnGatewayId.NotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeVpnGatewaysWithContext(ctx context.Context, request *DescribeVpnGatewaysRequest) (response *DescribeVpnGatewaysResponse, err error) {
    if request == nil {
        request = NewDescribeVpnGatewaysRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVpnGateways require credential")
    }

    request.SetContext(ctx)
    
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_APPIDMISMATCH = "UnsupportedOperation.AppIdMismatch"
//  UNSUPPORTEDOPERATION_APPIDNOTFOUND = "UnsupportedOperation.AppIdNotFound"
//  UNSUPPORTEDOPERATION_CCNNOTATTACHED = "UnsupportedOperation.CcnNotAttached"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) DetachCcnInstances(request *DetachCcnInstancesRequest) (response *DetachCcnInstancesResponse, err error) {
    return c.DetachCcnInstancesWithContext(context.Background(), request)
}

// DetachCcnInstances
// This API (DetachCcnInstances) is used to unbind a specified network instance from a CCN instance.<br />
//
// After unbinding the network instance, the corresponding routing policy will also be deleted.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_APPIDMISMATCH = "UnsupportedOperation.AppIdMismatch"
//  UNSUPPORTEDOPERATION_APPIDNOTFOUND = "UnsupportedOperation.AppIdNotFound"
//  UNSUPPORTEDOPERATION_CCNNOTATTACHED = "UnsupportedOperation.CcnNotAttached"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) DetachCcnInstancesWithContext(ctx context.Context, request *DetachCcnInstancesRequest) (response *DetachCcnInstancesResponse, err error) {
    if request == nil {
        request = NewDetachCcnInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetachCcnInstances require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to delete a Classiclink.
//
// >?This API is async. You can call the [`DescribeVpcTaskResult`](https://intl.cloud.tencent.com/document/api/215/59037?from_cn_redirect=1) API to query the task result. When the task is completed, you can continue other tasks.
//
// >
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DetachClassicLinkVpc(request *DetachClassicLinkVpcRequest) (response *DetachClassicLinkVpcResponse, err error) {
    return c.DetachClassicLinkVpcWithContext(context.Background(), request)
}

// DetachClassicLinkVpc
// This API is used to delete a Classiclink.
//
// >?This API is async. You can call the [`DescribeVpcTaskResult`](https://intl.cloud.tencent.com/document/api/215/59037?from_cn_redirect=1) API to query the task result. When the task is completed, you can continue other tasks.
//
// >
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DetachClassicLinkVpcWithContext(ctx context.Context, request *DetachClassicLinkVpcRequest) (response *DetachClassicLinkVpcResponse, err error) {
    if request == nil {
        request = NewDetachClassicLinkVpcRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetachClassicLinkVpc require credential")
    }

    request.SetContext(ctx)
    
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
// This API is completed asynchronously. If you need to query the execution result of an async task, please use the `RequestId` returned by this API to poll the `DescribeVpcTaskResult` API.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) DetachNetworkInterface(request *DetachNetworkInterfaceRequest) (response *DetachNetworkInterfaceResponse, err error) {
    return c.DetachNetworkInterfaceWithContext(context.Background(), request)
}

// DetachNetworkInterface
// This API is used to unbind an ENI from a CVM.
//
// This API is completed asynchronously. If you need to query the execution result of an async task, please use the `RequestId` returned by this API to poll the `DescribeVpcTaskResult` API.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) DetachNetworkInterfaceWithContext(ctx context.Context, request *DetachNetworkInterfaceRequest) (response *DetachNetworkInterfaceResponse, err error) {
    if request == nil {
        request = NewDetachNetworkInterfaceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetachNetworkInterface require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DisableCcnRoutes(request *DisableCcnRoutesRequest) (response *DisableCcnRoutesResponse, err error) {
    return c.DisableCcnRoutesWithContext(context.Background(), request)
}

// DisableCcnRoutes
// This API (DisableCcnRoutes) is used to disable CCN routes that are already enabled.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DisableCcnRoutesWithContext(ctx context.Context, request *DisableCcnRoutesRequest) (response *DisableCcnRoutesResponse, err error) {
    if request == nil {
        request = NewDisableCcnRoutesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableCcnRoutes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableCcnRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewDisableFlowLogsRequest() (request *DisableFlowLogsRequest) {
    request = &DisableFlowLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DisableFlowLogs")
    
    
    return
}

func NewDisableFlowLogsResponse() (response *DisableFlowLogsResponse) {
    response = &DisableFlowLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisableFlowLogs
// This API is used to disable flow log.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DisableFlowLogs(request *DisableFlowLogsRequest) (response *DisableFlowLogsResponse, err error) {
    return c.DisableFlowLogsWithContext(context.Background(), request)
}

// DisableFlowLogs
// This API is used to disable flow log.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DisableFlowLogsWithContext(ctx context.Context, request *DisableFlowLogsRequest) (response *DisableFlowLogsResponse, err error) {
    if request == nil {
        request = NewDisableFlowLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableFlowLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableFlowLogsResponse()
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
// This API is used to disable gateway traffic monitor.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) DisableGatewayFlowMonitor(request *DisableGatewayFlowMonitorRequest) (response *DisableGatewayFlowMonitorResponse, err error) {
    return c.DisableGatewayFlowMonitorWithContext(context.Background(), request)
}

// DisableGatewayFlowMonitor
// This API is used to disable gateway traffic monitor.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) DisableGatewayFlowMonitorWithContext(ctx context.Context, request *DisableGatewayFlowMonitorRequest) (response *DisableGatewayFlowMonitorResponse, err error) {
    if request == nil {
        request = NewDisableGatewayFlowMonitorRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableGatewayFlowMonitor require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION_ADDRESSENIINFONOTFOUND = "FailedOperation.AddressEniInfoNotFound"
//  FAILEDOPERATION_MASTERENINOTFOUND = "FailedOperation.MasterEniNotFound"
//  FAILEDOPERATION_TASKFAILED = "FailedOperation.TaskFailed"
//  INVALIDADDRESSID_BLOCKED = "InvalidAddressId.Blocked"
//  INVALIDADDRESSID_NOTFOUND = "InvalidAddressId.NotFound"
//  INVALIDADDRESSIDSTATUS_NOTPERMIT = "InvalidAddressIdStatus.NotPermit"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ADDRESSIDMALFORMED = "InvalidParameterValue.AddressIdMalformed"
//  INVALIDPARAMETERVALUE_ADDRESSNOTFOUND = "InvalidParameterValue.AddressNotFound"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_ONLYSUPPORTEDFORMASTERNETWORKCARD = "InvalidParameterValue.OnlySupportedForMasterNetworkCard"
//  INVALIDPARAMETERVALUE_RESOURCENOTSUPPORT = "InvalidParameterValue.ResourceNotSupport"
//  OPERATIONDENIED_MUTEXTASKRUNNING = "OperationDenied.MutexTaskRunning"
//  UNSUPPORTEDOPERATION_ACTIONNOTFOUND = "UnsupportedOperation.ActionNotFound"
//  UNSUPPORTEDOPERATION_ADDRESSSTATUSNOTPERMIT = "UnsupportedOperation.AddressStatusNotPermit"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DisassociateAddress(request *DisassociateAddressRequest) (response *DisassociateAddressResponse, err error) {
    return c.DisassociateAddressWithContext(context.Background(), request)
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
//  FAILEDOPERATION_ADDRESSENIINFONOTFOUND = "FailedOperation.AddressEniInfoNotFound"
//  FAILEDOPERATION_MASTERENINOTFOUND = "FailedOperation.MasterEniNotFound"
//  FAILEDOPERATION_TASKFAILED = "FailedOperation.TaskFailed"
//  INVALIDADDRESSID_BLOCKED = "InvalidAddressId.Blocked"
//  INVALIDADDRESSID_NOTFOUND = "InvalidAddressId.NotFound"
//  INVALIDADDRESSIDSTATUS_NOTPERMIT = "InvalidAddressIdStatus.NotPermit"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ADDRESSIDMALFORMED = "InvalidParameterValue.AddressIdMalformed"
//  INVALIDPARAMETERVALUE_ADDRESSNOTFOUND = "InvalidParameterValue.AddressNotFound"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_ONLYSUPPORTEDFORMASTERNETWORKCARD = "InvalidParameterValue.OnlySupportedForMasterNetworkCard"
//  INVALIDPARAMETERVALUE_RESOURCENOTSUPPORT = "InvalidParameterValue.ResourceNotSupport"
//  OPERATIONDENIED_MUTEXTASKRUNNING = "OperationDenied.MutexTaskRunning"
//  UNSUPPORTEDOPERATION_ACTIONNOTFOUND = "UnsupportedOperation.ActionNotFound"
//  UNSUPPORTEDOPERATION_ADDRESSSTATUSNOTPERMIT = "UnsupportedOperation.AddressStatusNotPermit"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DisassociateAddressWithContext(ctx context.Context, request *DisassociateAddressRequest) (response *DisassociateAddressResponse, err error) {
    if request == nil {
        request = NewDisassociateAddressRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisassociateAddress require credential")
    }

    request.SetContext(ctx)
    
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
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DisassociateDirectConnectGatewayNatGateway(request *DisassociateDirectConnectGatewayNatGatewayRequest) (response *DisassociateDirectConnectGatewayNatGatewayResponse, err error) {
    return c.DisassociateDirectConnectGatewayNatGatewayWithContext(context.Background(), request)
}

// DisassociateDirectConnectGatewayNatGateway
// This API is used to unbind a direct connect gateway from a NAT Gateway. After unbinding, the direct connect gateway cannot access internet through the NAT Gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DisassociateDirectConnectGatewayNatGatewayWithContext(ctx context.Context, request *DisassociateDirectConnectGatewayNatGatewayRequest) (response *DisassociateDirectConnectGatewayNatGatewayResponse, err error) {
    if request == nil {
        request = NewDisassociateDirectConnectGatewayNatGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisassociateDirectConnectGatewayNatGateway require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to unbind an EIP from a NAT gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_PUBLICIPADDRESSDISASSOCIATE = "UnsupportedOperation.PublicIpAddressDisassociate"
func (c *Client) DisassociateNatGatewayAddress(request *DisassociateNatGatewayAddressRequest) (response *DisassociateNatGatewayAddressResponse, err error) {
    return c.DisassociateNatGatewayAddressWithContext(context.Background(), request)
}

// DisassociateNatGatewayAddress
// This API is used to unbind an EIP from a NAT gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_PUBLICIPADDRESSDISASSOCIATE = "UnsupportedOperation.PublicIpAddressDisassociate"
func (c *Client) DisassociateNatGatewayAddressWithContext(ctx context.Context, request *DisassociateNatGatewayAddressRequest) (response *DisassociateNatGatewayAddressResponse, err error) {
    if request == nil {
        request = NewDisassociateNatGatewayAddressRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisassociateNatGatewayAddress require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DisassociateNetworkAclSubnetsWithContext(context.Background(), request)
}

// DisassociateNetworkAclSubnets
// This API is used to disassociate a network ACL from subnets in a VPC instance.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
func (c *Client) DisassociateNetworkAclSubnetsWithContext(ctx context.Context, request *DisassociateNetworkAclSubnetsRequest) (response *DisassociateNetworkAclSubnetsResponse, err error) {
    if request == nil {
        request = NewDisassociateNetworkAclSubnetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisassociateNetworkAclSubnets require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DisassociateNetworkInterfaceSecurityGroups(request *DisassociateNetworkInterfaceSecurityGroupsRequest) (response *DisassociateNetworkInterfaceSecurityGroupsResponse, err error) {
    return c.DisassociateNetworkInterfaceSecurityGroupsWithContext(context.Background(), request)
}

// DisassociateNetworkInterfaceSecurityGroups
// This API (DisassociateNetworkInterfaceSecurityGroups) is used to detach (or fully detach if possible) a security group from an ENI.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DisassociateNetworkInterfaceSecurityGroupsWithContext(ctx context.Context, request *DisassociateNetworkInterfaceSecurityGroupsRequest) (response *DisassociateNetworkInterfaceSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDisassociateNetworkInterfaceSecurityGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisassociateNetworkInterfaceSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DisassociateVpcEndPointSecurityGroupsWithContext(context.Background(), request)
}

// DisassociateVpcEndPointSecurityGroups
// This API is used to unbind an endpoint from a security group.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DisassociateVpcEndPointSecurityGroupsWithContext(ctx context.Context, request *DisassociateVpcEndPointSecurityGroupsRequest) (response *DisassociateVpcEndPointSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDisassociateVpcEndPointSecurityGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisassociateVpcEndPointSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DownloadCustomerGatewayConfiguration(request *DownloadCustomerGatewayConfigurationRequest) (response *DownloadCustomerGatewayConfigurationResponse, err error) {
    return c.DownloadCustomerGatewayConfigurationWithContext(context.Background(), request)
}

// DownloadCustomerGatewayConfiguration
// This API (DownloadCustomerGatewayConfiguration) is used to download a VPN tunnel configuration.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DownloadCustomerGatewayConfigurationWithContext(ctx context.Context, request *DownloadCustomerGatewayConfigurationRequest) (response *DownloadCustomerGatewayConfigurationResponse, err error) {
    if request == nil {
        request = NewDownloadCustomerGatewayConfigurationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DownloadCustomerGatewayConfiguration require credential")
    }

    request.SetContext(ctx)
    
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
    return c.EnableCcnRoutesWithContext(context.Background(), request)
}

// EnableCcnRoutes
// This API (EnableCcnRoutes) is used to enable CCN routes that are already added.<br />
//
// This API is used to verify whether there will be conflict with an existing route after a CCN route is enabled. If there is a conflict, the route will not be enabled, and the process will fail. When a conflict occurs, you must disable the conflicting route before you can enable the desired route.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_ECMP = "UnsupportedOperation.Ecmp"
func (c *Client) EnableCcnRoutesWithContext(ctx context.Context, request *EnableCcnRoutesRequest) (response *EnableCcnRoutesResponse, err error) {
    if request == nil {
        request = NewEnableCcnRoutesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableCcnRoutes require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableCcnRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewEnableFlowLogsRequest() (request *EnableFlowLogsRequest) {
    request = &EnableFlowLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "EnableFlowLogs")
    
    
    return
}

func NewEnableFlowLogsResponse() (response *EnableFlowLogsResponse) {
    response = &EnableFlowLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EnableFlowLogs
// This API is used to enable flow log.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) EnableFlowLogs(request *EnableFlowLogsRequest) (response *EnableFlowLogsResponse, err error) {
    return c.EnableFlowLogsWithContext(context.Background(), request)
}

// EnableFlowLogs
// This API is used to enable flow log.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) EnableFlowLogsWithContext(ctx context.Context, request *EnableFlowLogsRequest) (response *EnableFlowLogsResponse, err error) {
    if request == nil {
        request = NewEnableFlowLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableFlowLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableFlowLogsResponse()
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
// This API is used to enable gateway traffic monitor.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) EnableGatewayFlowMonitor(request *EnableGatewayFlowMonitorRequest) (response *EnableGatewayFlowMonitorResponse, err error) {
    return c.EnableGatewayFlowMonitorWithContext(context.Background(), request)
}

// EnableGatewayFlowMonitor
// This API is used to enable gateway traffic monitor.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) EnableGatewayFlowMonitorWithContext(ctx context.Context, request *EnableGatewayFlowMonitorRequest) (response *EnableGatewayFlowMonitorResponse, err error) {
    if request == nil {
        request = NewEnableGatewayFlowMonitorRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableGatewayFlowMonitor require credential")
    }

    request.SetContext(ctx)
    
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
    return c.EnableVpcEndPointConnectWithContext(context.Background(), request)
}

// EnableVpcEndPointConnect
// This API is used to determine whether to accept the request of connecting with an endpoint.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) EnableVpcEndPointConnectWithContext(ctx context.Context, request *EnableVpcEndPointConnectRequest) (response *EnableVpcEndPointConnectResponse, err error) {
    if request == nil {
        request = NewEnableVpcEndPointConnectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableVpcEndPointConnect require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetCcnRegionBandwidthLimits(request *GetCcnRegionBandwidthLimitsRequest) (response *GetCcnRegionBandwidthLimitsResponse, err error) {
    return c.GetCcnRegionBandwidthLimitsWithContext(context.Background(), request)
}

// GetCcnRegionBandwidthLimits
// This API is used to query the bandwidth limits of a CCN instance. Monthly-subscribed CCNs only support Inter-region Bandwidth Limits, while the pay-as-you-go CCNs support both the Inter-region Bandwidth Limits and Region Outbound Bandwidth Limits. 
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetCcnRegionBandwidthLimitsWithContext(ctx context.Context, request *GetCcnRegionBandwidthLimitsRequest) (response *GetCcnRegionBandwidthLimitsResponse, err error) {
    if request == nil {
        request = NewGetCcnRegionBandwidthLimitsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetCcnRegionBandwidthLimits require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to bind an EIP to an HAVIP. <br />
//
// This API is completed asynchronously. If you need to query the execution result of an async task, please use the `RequestId` returned by this API to poll the `DescribeVpcTaskResult` API.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_BINDEIP = "UnsupportedOperation.BindEIP"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) HaVipAssociateAddressIp(request *HaVipAssociateAddressIpRequest) (response *HaVipAssociateAddressIpResponse, err error) {
    return c.HaVipAssociateAddressIpWithContext(context.Background(), request)
}

// HaVipAssociateAddressIp
// This API is used to bind an EIP to an HAVIP. <br />
//
// This API is completed asynchronously. If you need to query the execution result of an async task, please use the `RequestId` returned by this API to poll the `DescribeVpcTaskResult` API.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_BINDEIP = "UnsupportedOperation.BindEIP"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) HaVipAssociateAddressIpWithContext(ctx context.Context, request *HaVipAssociateAddressIpRequest) (response *HaVipAssociateAddressIpResponse, err error) {
    if request == nil {
        request = NewHaVipAssociateAddressIpRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("HaVipAssociateAddressIp require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to unbind an EIP from an HAVIP. <br />
//
// This API is completed asynchronously. If you need to query the execution result of an async task, please use the `RequestId` returned by this API to poll the `DescribeVpcTaskResult` API.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
//  UNSUPPORTEDOPERATION_UNBINDEIP = "UnsupportedOperation.UnbindEIP"
func (c *Client) HaVipDisassociateAddressIp(request *HaVipDisassociateAddressIpRequest) (response *HaVipDisassociateAddressIpResponse, err error) {
    return c.HaVipDisassociateAddressIpWithContext(context.Background(), request)
}

// HaVipDisassociateAddressIp
// This API is used to unbind an EIP from an HAVIP. <br />
//
// This API is completed asynchronously. If you need to query the execution result of an async task, please use the `RequestId` returned by this API to poll the `DescribeVpcTaskResult` API.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
//  UNSUPPORTEDOPERATION_UNBINDEIP = "UnsupportedOperation.UnbindEIP"
func (c *Client) HaVipDisassociateAddressIpWithContext(ctx context.Context, request *HaVipDisassociateAddressIpRequest) (response *HaVipDisassociateAddressIpResponse, err error) {
    if request == nil {
        request = NewHaVipDisassociateAddressIpRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("HaVipDisassociateAddressIp require credential")
    }

    request.SetContext(ctx)
    
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
    return c.InquirePriceCreateDirectConnectGatewayWithContext(context.Background(), request)
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
func (c *Client) InquirePriceCreateDirectConnectGatewayWithContext(ctx context.Context, request *InquirePriceCreateDirectConnectGatewayRequest) (response *InquirePriceCreateDirectConnectGatewayResponse, err error) {
    if request == nil {
        request = NewInquirePriceCreateDirectConnectGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquirePriceCreateDirectConnectGateway require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquirePriceCreateDirectConnectGatewayResponse()
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
    return c.InquiryPriceRenewVpnGatewayWithContext(context.Background(), request)
}

// InquiryPriceRenewVpnGateway
// This API (InquiryPriceRenewVpnGateway) is used to query the price for VPN gateway renewal. Currently, only querying prices for IPSEC-type gateways is supported.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) InquiryPriceRenewVpnGatewayWithContext(ctx context.Context, request *InquiryPriceRenewVpnGatewayRequest) (response *InquiryPriceRenewVpnGatewayResponse, err error) {
    if request == nil {
        request = NewInquiryPriceRenewVpnGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceRenewVpnGateway require credential")
    }

    request.SetContext(ctx)
    
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
    return c.InquiryPriceResetVpnGatewayInternetMaxBandwidthWithContext(context.Background(), request)
}

// InquiryPriceResetVpnGatewayInternetMaxBandwidth
// This API (InquiryPriceResetVpnGatewayInternetMaxBandwidth) is used to query the price for adjusting the bandwidth cap of a VPN gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) InquiryPriceResetVpnGatewayInternetMaxBandwidthWithContext(ctx context.Context, request *InquiryPriceResetVpnGatewayInternetMaxBandwidthRequest) (response *InquiryPriceResetVpnGatewayInternetMaxBandwidthResponse, err error) {
    if request == nil {
        request = NewInquiryPriceResetVpnGatewayInternetMaxBandwidthRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceResetVpnGatewayInternetMaxBandwidth require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to migrate ENIs.
//
// This API is completed asynchronously. If you need to query the execution result of an async task, please use the `RequestId` returned by this API to poll the `DescribeVpcTaskResult` API.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
func (c *Client) MigrateNetworkInterface(request *MigrateNetworkInterfaceRequest) (response *MigrateNetworkInterfaceResponse, err error) {
    return c.MigrateNetworkInterfaceWithContext(context.Background(), request)
}

// MigrateNetworkInterface
// This API is used to migrate ENIs.
//
// This API is completed asynchronously. If you need to query the execution result of an async task, please use the `RequestId` returned by this API to poll the `DescribeVpcTaskResult` API.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
func (c *Client) MigrateNetworkInterfaceWithContext(ctx context.Context, request *MigrateNetworkInterfaceRequest) (response *MigrateNetworkInterfaceResponse, err error) {
    if request == nil {
        request = NewMigrateNetworkInterfaceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("MigrateNetworkInterface require credential")
    }

    request.SetContext(ctx)
    
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
//  This API is used to migrate the private IPs between ENIs.
//
// * This API is used to migrate a private IP from one ENI to another. Primary IPs cannot be migrated.
//
// * The source and destination ENIs must be in the same subnet.  
//
// 
//
// This API is completed asynchronously. If you need to query the execution result of an async task, please use the `RequestId` returned by this API to poll the `DescribeVpcTaskResult` API.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_ATTACHMENTNOTFOUND = "UnauthorizedOperation.AttachmentNotFound"
//  UNAUTHORIZEDOPERATION_PRIMARYIP = "UnauthorizedOperation.PrimaryIp"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
//  UNSUPPORTEDOPERATION_PRIMARYIP = "UnsupportedOperation.PrimaryIp"
func (c *Client) MigratePrivateIpAddress(request *MigratePrivateIpAddressRequest) (response *MigratePrivateIpAddressResponse, err error) {
    return c.MigratePrivateIpAddressWithContext(context.Background(), request)
}

// MigratePrivateIpAddress
//  This API is used to migrate the private IPs between ENIs.
//
// * This API is used to migrate a private IP from one ENI to another. Primary IPs cannot be migrated.
//
// * The source and destination ENIs must be in the same subnet.  
//
// 
//
// This API is completed asynchronously. If you need to query the execution result of an async task, please use the `RequestId` returned by this API to poll the `DescribeVpcTaskResult` API.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_ATTACHMENTNOTFOUND = "UnauthorizedOperation.AttachmentNotFound"
//  UNAUTHORIZEDOPERATION_PRIMARYIP = "UnauthorizedOperation.PrimaryIp"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
//  UNSUPPORTEDOPERATION_PRIMARYIP = "UnsupportedOperation.PrimaryIp"
func (c *Client) MigratePrivateIpAddressWithContext(ctx context.Context, request *MigratePrivateIpAddressRequest) (response *MigratePrivateIpAddressResponse, err error) {
    if request == nil {
        request = NewMigratePrivateIpAddressRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("MigratePrivateIpAddress require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_ADDRESSIDMALFORMED = "InvalidParameterValue.AddressIdMalformed"
//  INVALIDPARAMETERVALUE_ADDRESSNOTFOUND = "InvalidParameterValue.AddressNotFound"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  OPERATIONDENIED_MUTEXTASKRUNNING = "OperationDenied.MutexTaskRunning"
//  UNSUPPORTEDOPERATION_ADDRESSSTATUSNOTPERMIT = "UnsupportedOperation.AddressStatusNotPermit"
//  UNSUPPORTEDOPERATION_INCORRECTADDRESSRESOURCETYPE = "UnsupportedOperation.IncorrectAddressResourceType"
//  UNSUPPORTEDOPERATION_MODIFYADDRESSATTRIBUTE = "UnsupportedOperation.ModifyAddressAttribute"
func (c *Client) ModifyAddressAttribute(request *ModifyAddressAttributeRequest) (response *ModifyAddressAttributeResponse, err error) {
    return c.ModifyAddressAttributeWithContext(context.Background(), request)
}

// ModifyAddressAttribute
// This API (ModifyAddressAttribute) is used to modify the name of an [Elastic IP](https://intl.cloud.tencent.com/document/product/213/1941?from_cn_redirect=1).
//
// error code that may be returned:
//  INVALIDADDRESSID_BLOCKED = "InvalidAddressId.Blocked"
//  INVALIDADDRESSID_NOTFOUND = "InvalidAddressId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ADDRESSIDMALFORMED = "InvalidParameterValue.AddressIdMalformed"
//  INVALIDPARAMETERVALUE_ADDRESSNOTFOUND = "InvalidParameterValue.AddressNotFound"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  OPERATIONDENIED_MUTEXTASKRUNNING = "OperationDenied.MutexTaskRunning"
//  UNSUPPORTEDOPERATION_ADDRESSSTATUSNOTPERMIT = "UnsupportedOperation.AddressStatusNotPermit"
//  UNSUPPORTEDOPERATION_INCORRECTADDRESSRESOURCETYPE = "UnsupportedOperation.IncorrectAddressResourceType"
//  UNSUPPORTEDOPERATION_MODIFYADDRESSATTRIBUTE = "UnsupportedOperation.ModifyAddressAttribute"
func (c *Client) ModifyAddressAttributeWithContext(ctx context.Context, request *ModifyAddressAttributeRequest) (response *ModifyAddressAttributeResponse, err error) {
    if request == nil {
        request = NewModifyAddressAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAddressAttribute require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_ADDRESSIDMALFORMED = "InvalidParameterValue.AddressIdMalformed"
//  INVALIDPARAMETERVALUE_ADDRESSNOTCALCIP = "InvalidParameterValue.AddressNotCalcIP"
//  INVALIDPARAMETERVALUE_ADDRESSNOTFOUND = "InvalidParameterValue.AddressNotFound"
//  INVALIDPARAMETERVALUE_INTERNETCHARGETYPENOTCHANGED = "InvalidParameterValue.InternetChargeTypeNotChanged"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_MODIFYADDRESSINTERNETCHARGETYPEQUOTA = "LimitExceeded.ModifyAddressInternetChargeTypeQuota"
//  UNSUPPORTEDOPERATION_ADDRESSSTATUSNOTPERMIT = "UnsupportedOperation.AddressStatusNotPermit"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_INVALIDADDRESSINTERNETCHARGETYPE = "UnsupportedOperation.InvalidAddressInternetChargeType"
//  UNSUPPORTEDOPERATION_NATNOTSUPPORTED = "UnsupportedOperation.NatNotSupported"
func (c *Client) ModifyAddressInternetChargeType(request *ModifyAddressInternetChargeTypeRequest) (response *ModifyAddressInternetChargeTypeResponse, err error) {
    return c.ModifyAddressInternetChargeTypeWithContext(context.Background(), request)
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
//  INVALIDPARAMETERVALUE_ADDRESSIDMALFORMED = "InvalidParameterValue.AddressIdMalformed"
//  INVALIDPARAMETERVALUE_ADDRESSNOTCALCIP = "InvalidParameterValue.AddressNotCalcIP"
//  INVALIDPARAMETERVALUE_ADDRESSNOTFOUND = "InvalidParameterValue.AddressNotFound"
//  INVALIDPARAMETERVALUE_INTERNETCHARGETYPENOTCHANGED = "InvalidParameterValue.InternetChargeTypeNotChanged"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_MODIFYADDRESSINTERNETCHARGETYPEQUOTA = "LimitExceeded.ModifyAddressInternetChargeTypeQuota"
//  UNSUPPORTEDOPERATION_ADDRESSSTATUSNOTPERMIT = "UnsupportedOperation.AddressStatusNotPermit"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_INVALIDADDRESSINTERNETCHARGETYPE = "UnsupportedOperation.InvalidAddressInternetChargeType"
//  UNSUPPORTEDOPERATION_NATNOTSUPPORTED = "UnsupportedOperation.NatNotSupported"
func (c *Client) ModifyAddressInternetChargeTypeWithContext(ctx context.Context, request *ModifyAddressInternetChargeTypeRequest) (response *ModifyAddressInternetChargeTypeResponse, err error) {
    if request == nil {
        request = NewModifyAddressInternetChargeTypeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAddressInternetChargeType require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) ModifyAddressTemplateAttribute(request *ModifyAddressTemplateAttributeRequest) (response *ModifyAddressTemplateAttributeResponse, err error) {
    return c.ModifyAddressTemplateAttributeWithContext(context.Background(), request)
}

// ModifyAddressTemplateAttribute
// This API (ModifyAddressTemplateAttribute) is used to modify an IP address template.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) ModifyAddressTemplateAttributeWithContext(ctx context.Context, request *ModifyAddressTemplateAttributeRequest) (response *ModifyAddressTemplateAttributeResponse, err error) {
    if request == nil {
        request = NewModifyAddressTemplateAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAddressTemplateAttribute require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) ModifyAddressTemplateGroupAttribute(request *ModifyAddressTemplateGroupAttributeRequest) (response *ModifyAddressTemplateGroupAttributeResponse, err error) {
    return c.ModifyAddressTemplateGroupAttributeWithContext(context.Background(), request)
}

// ModifyAddressTemplateGroupAttribute
// This API (ModifyAddressTemplateGroupAttribute) is used to modify an IP address template group.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) ModifyAddressTemplateGroupAttributeWithContext(ctx context.Context, request *ModifyAddressTemplateGroupAttributeRequest) (response *ModifyAddressTemplateGroupAttributeResponse, err error) {
    if request == nil {
        request = NewModifyAddressTemplateGroupAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAddressTemplateGroupAttribute require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_ADDRESSIDMALFORMED = "InvalidParameterValue.AddressIdMalformed"
//  INVALIDPARAMETERVALUE_ADDRESSNOTFOUND = "InvalidParameterValue.AddressNotFound"
//  INVALIDPARAMETERVALUE_BANDWIDTHOUTOFRANGE = "InvalidParameterValue.BandwidthOutOfRange"
//  INVALIDPARAMETERVALUE_BANDWIDTHTOOSMALL = "InvalidParameterValue.BandwidthTooSmall"
//  INVALIDPARAMETERVALUE_INCONSISTENTINSTANCEINTERNETCHARGETYPE = "InvalidParameterValue.InconsistentInstanceInternetChargeType"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INSTANCENOCALCIP = "InvalidParameterValue.InstanceNoCalcIP"
//  INVALIDPARAMETERVALUE_INSTANCENOWANIP = "InvalidParameterValue.InstanceNoWanIP"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_RESOURCEEXPIRED = "InvalidParameterValue.ResourceExpired"
//  INVALIDPARAMETERVALUE_RESOURCENOTEXISTED = "InvalidParameterValue.ResourceNotExisted"
//  INVALIDPARAMETERVALUE_RESOURCENOTSUPPORT = "InvalidParameterValue.ResourceNotSupport"
//  UNSUPPORTEDOPERATION_ACTIONNOTFOUND = "UnsupportedOperation.ActionNotFound"
//  UNSUPPORTEDOPERATION_ADDRESSSTATUSNOTPERMIT = "UnsupportedOperation.AddressStatusNotPermit"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTSUPPORTED = "UnsupportedOperation.InstanceStateNotSupported"
func (c *Client) ModifyAddressesBandwidth(request *ModifyAddressesBandwidthRequest) (response *ModifyAddressesBandwidthResponse, err error) {
    return c.ModifyAddressesBandwidthWithContext(context.Background(), request)
}

// ModifyAddressesBandwidth
// This API is used to adjust the bandwidth of [Elastic IP](https://intl.cloud.tencent.com/document/product/213/1941?from_cn_redirect=1), including EIP billed on a pay-as-you-go, monthly subscription, and bandwidth package basis.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDADDRESSID_NOTFOUND = "InvalidAddressId.NotFound"
//  INVALIDPARAMETERVALUE_ADDRESSIDMALFORMED = "InvalidParameterValue.AddressIdMalformed"
//  INVALIDPARAMETERVALUE_ADDRESSNOTFOUND = "InvalidParameterValue.AddressNotFound"
//  INVALIDPARAMETERVALUE_BANDWIDTHOUTOFRANGE = "InvalidParameterValue.BandwidthOutOfRange"
//  INVALIDPARAMETERVALUE_BANDWIDTHTOOSMALL = "InvalidParameterValue.BandwidthTooSmall"
//  INVALIDPARAMETERVALUE_INCONSISTENTINSTANCEINTERNETCHARGETYPE = "InvalidParameterValue.InconsistentInstanceInternetChargeType"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INSTANCENOCALCIP = "InvalidParameterValue.InstanceNoCalcIP"
//  INVALIDPARAMETERVALUE_INSTANCENOWANIP = "InvalidParameterValue.InstanceNoWanIP"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_RESOURCEEXPIRED = "InvalidParameterValue.ResourceExpired"
//  INVALIDPARAMETERVALUE_RESOURCENOTEXISTED = "InvalidParameterValue.ResourceNotExisted"
//  INVALIDPARAMETERVALUE_RESOURCENOTSUPPORT = "InvalidParameterValue.ResourceNotSupport"
//  UNSUPPORTEDOPERATION_ACTIONNOTFOUND = "UnsupportedOperation.ActionNotFound"
//  UNSUPPORTEDOPERATION_ADDRESSSTATUSNOTPERMIT = "UnsupportedOperation.AddressStatusNotPermit"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTSUPPORTED = "UnsupportedOperation.InstanceStateNotSupported"
func (c *Client) ModifyAddressesBandwidthWithContext(ctx context.Context, request *ModifyAddressesBandwidthRequest) (response *ModifyAddressesBandwidthResponse, err error) {
    if request == nil {
        request = NewModifyAddressesBandwidthRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAddressesBandwidth require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to modify (add or delete) secondary CIDR blocks in batch. This API is in beta test. To use it, please submit a ticket.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_SUBNETCONFLICT = "InvalidParameterValue.SubnetConflict"
//  INVALIDPARAMETERVALUE_SUBNETOVERLAPASSISTCIDR = "InvalidParameterValue.SubnetOverlapAssistCidr"
//  INVALIDPARAMETERVALUE_VPCCIDRCONFLICT = "InvalidParameterValue.VpcCidrConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAssistantCidr(request *ModifyAssistantCidrRequest) (response *ModifyAssistantCidrResponse, err error) {
    return c.ModifyAssistantCidrWithContext(context.Background(), request)
}

// ModifyAssistantCidr
// This API is used to modify (add or delete) secondary CIDR blocks in batch. This API is in beta test. To use it, please submit a ticket.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_SUBNETCONFLICT = "InvalidParameterValue.SubnetConflict"
//  INVALIDPARAMETERVALUE_SUBNETOVERLAPASSISTCIDR = "InvalidParameterValue.SubnetOverlapAssistCidr"
//  INVALIDPARAMETERVALUE_VPCCIDRCONFLICT = "InvalidParameterValue.VpcCidrConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAssistantCidrWithContext(ctx context.Context, request *ModifyAssistantCidrRequest) (response *ModifyAssistantCidrResponse, err error) {
    if request == nil {
        request = NewModifyAssistantCidrRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAssistantCidr require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGEIDMALFORMED = "InvalidParameterValue.BandwidthPackageIdMalformed"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGENOTFOUND = "InvalidParameterValue.BandwidthPackageNotFound"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTHPACKAGECHARGETYPE = "InvalidParameterValue.InvalidBandwidthPackageChargeType"
func (c *Client) ModifyBandwidthPackageAttribute(request *ModifyBandwidthPackageAttributeRequest) (response *ModifyBandwidthPackageAttributeResponse, err error) {
    return c.ModifyBandwidthPackageAttributeWithContext(context.Background(), request)
}

// ModifyBandwidthPackageAttribute
// This API is used to modify the attributes of a bandwidth package, including the bandwidth package name, and so on.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGEIDMALFORMED = "InvalidParameterValue.BandwidthPackageIdMalformed"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGENOTFOUND = "InvalidParameterValue.BandwidthPackageNotFound"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTHPACKAGECHARGETYPE = "InvalidParameterValue.InvalidBandwidthPackageChargeType"
func (c *Client) ModifyBandwidthPackageAttributeWithContext(ctx context.Context, request *ModifyBandwidthPackageAttributeRequest) (response *ModifyBandwidthPackageAttributeResponse, err error) {
    if request == nil {
        request = NewModifyBandwidthPackageAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBandwidthPackageAttribute require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ModifyCcnAttachedInstancesAttributeWithContext(context.Background(), request)
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
func (c *Client) ModifyCcnAttachedInstancesAttributeWithContext(ctx context.Context, request *ModifyCcnAttachedInstancesAttributeRequest) (response *ModifyCcnAttachedInstancesAttributeResponse, err error) {
    if request == nil {
        request = NewModifyCcnAttachedInstancesAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCcnAttachedInstancesAttribute require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyCcnAttribute(request *ModifyCcnAttributeRequest) (response *ModifyCcnAttributeResponse, err error) {
    return c.ModifyCcnAttributeWithContext(context.Background(), request)
}

// ModifyCcnAttribute
// This API (ModifyCcnAttribute) is used to modify CCN attributes.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyCcnAttributeWithContext(ctx context.Context, request *ModifyCcnAttributeRequest) (response *ModifyCcnAttributeResponse, err error) {
    if request == nil {
        request = NewModifyCcnAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCcnAttribute require credential")
    }

    request.SetContext(ctx)
    
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
//  UNSUPPORTEDOPERATION_NOTLOCKEDINSTANCEOPERATION = "UnsupportedOperation.NotLockedInstanceOperation"
//  UNSUPPORTEDOPERATION_NOTPOSTPAIDCCNOPERATION = "UnsupportedOperation.NotPostpaidCcnOperation"
func (c *Client) ModifyCcnRegionBandwidthLimitsType(request *ModifyCcnRegionBandwidthLimitsTypeRequest) (response *ModifyCcnRegionBandwidthLimitsTypeResponse, err error) {
    return c.ModifyCcnRegionBandwidthLimitsTypeWithContext(context.Background(), request)
}

// ModifyCcnRegionBandwidthLimitsType
// This API is used to modify the bandwidth limit policy of a postpaid CCN instance.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_NOTLOCKEDINSTANCEOPERATION = "UnsupportedOperation.NotLockedInstanceOperation"
//  UNSUPPORTEDOPERATION_NOTPOSTPAIDCCNOPERATION = "UnsupportedOperation.NotPostpaidCcnOperation"
func (c *Client) ModifyCcnRegionBandwidthLimitsTypeWithContext(ctx context.Context, request *ModifyCcnRegionBandwidthLimitsTypeRequest) (response *ModifyCcnRegionBandwidthLimitsTypeResponse, err error) {
    if request == nil {
        request = NewModifyCcnRegionBandwidthLimitsTypeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCcnRegionBandwidthLimitsType require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyCustomerGatewayAttribute(request *ModifyCustomerGatewayAttributeRequest) (response *ModifyCustomerGatewayAttributeResponse, err error) {
    return c.ModifyCustomerGatewayAttributeWithContext(context.Background(), request)
}

// ModifyCustomerGatewayAttribute
// This API (ModifyCustomerGatewayAttribute) is used to modify the customer gateway information.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyCustomerGatewayAttributeWithContext(ctx context.Context, request *ModifyCustomerGatewayAttributeRequest) (response *ModifyCustomerGatewayAttributeResponse, err error) {
    if request == nil {
        request = NewModifyCustomerGatewayAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCustomerGatewayAttribute require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ModifyDirectConnectGatewayAttributeWithContext(context.Background(), request)
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
func (c *Client) ModifyDirectConnectGatewayAttributeWithContext(ctx context.Context, request *ModifyDirectConnectGatewayAttributeRequest) (response *ModifyDirectConnectGatewayAttributeResponse, err error) {
    if request == nil {
        request = NewModifyDirectConnectGatewayAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDirectConnectGatewayAttribute require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
func (c *Client) ModifyFlowLogAttribute(request *ModifyFlowLogAttributeRequest) (response *ModifyFlowLogAttributeResponse, err error) {
    return c.ModifyFlowLogAttributeWithContext(context.Background(), request)
}

// ModifyFlowLogAttribute
// This API is used to modify the attributes of a flow log.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
func (c *Client) ModifyFlowLogAttributeWithContext(ctx context.Context, request *ModifyFlowLogAttributeRequest) (response *ModifyFlowLogAttributeResponse, err error) {
    if request == nil {
        request = NewModifyFlowLogAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyFlowLogAttribute require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to adjust the bandwidth limit of a gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) ModifyGatewayFlowQos(request *ModifyGatewayFlowQosRequest) (response *ModifyGatewayFlowQosResponse, err error) {
    return c.ModifyGatewayFlowQosWithContext(context.Background(), request)
}

// ModifyGatewayFlowQos
// This API is used to adjust the bandwidth limit of a gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) ModifyGatewayFlowQosWithContext(ctx context.Context, request *ModifyGatewayFlowQosRequest) (response *ModifyGatewayFlowQosResponse, err error) {
    if request == nil {
        request = NewModifyGatewayFlowQosRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyGatewayFlowQos require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ModifyHaVipAttributeWithContext(context.Background(), request)
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
func (c *Client) ModifyHaVipAttributeWithContext(ctx context.Context, request *ModifyHaVipAttributeRequest) (response *ModifyHaVipAttributeResponse, err error) {
    if request == nil {
        request = NewModifyHaVipAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyHaVipAttribute require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ModifyIpv6AddressesAttributeWithContext(context.Background(), request)
}

// ModifyIpv6AddressesAttribute
// This API (ModifyIpv6AddressesAttribute) is used to modify the private IPv6 address attributes of an ENI.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ATTACHMENTNOTFOUND = "UnsupportedOperation.AttachmentNotFound"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) ModifyIpv6AddressesAttributeWithContext(ctx context.Context, request *ModifyIpv6AddressesAttributeRequest) (response *ModifyIpv6AddressesAttributeResponse, err error) {
    if request == nil {
        request = NewModifyIpv6AddressesAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyIpv6AddressesAttribute require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ModifyLocalGatewayWithContext(context.Background(), request)
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
func (c *Client) ModifyLocalGatewayWithContext(ctx context.Context, request *ModifyLocalGatewayRequest) (response *ModifyLocalGatewayResponse, err error) {
    if request == nil {
        request = NewModifyLocalGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLocalGateway require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to modify the attributes of a NAT gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyNatGatewayAttribute(request *ModifyNatGatewayAttributeRequest) (response *ModifyNatGatewayAttributeResponse, err error) {
    return c.ModifyNatGatewayAttributeWithContext(context.Background(), request)
}

// ModifyNatGatewayAttribute
// This API is used to modify the attributes of a NAT gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyNatGatewayAttributeWithContext(ctx context.Context, request *ModifyNatGatewayAttributeRequest) (response *ModifyNatGatewayAttributeResponse, err error) {
    if request == nil {
        request = NewModifyNatGatewayAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNatGatewayAttribute require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to modify the port forwarding rule of a NAT gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyNatGatewayDestinationIpPortTranslationNatRule(request *ModifyNatGatewayDestinationIpPortTranslationNatRuleRequest) (response *ModifyNatGatewayDestinationIpPortTranslationNatRuleResponse, err error) {
    return c.ModifyNatGatewayDestinationIpPortTranslationNatRuleWithContext(context.Background(), request)
}

// ModifyNatGatewayDestinationIpPortTranslationNatRule
// This API is used to modify the port forwarding rule of a NAT gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyNatGatewayDestinationIpPortTranslationNatRuleWithContext(ctx context.Context, request *ModifyNatGatewayDestinationIpPortTranslationNatRuleRequest) (response *ModifyNatGatewayDestinationIpPortTranslationNatRuleResponse, err error) {
    if request == nil {
        request = NewModifyNatGatewayDestinationIpPortTranslationNatRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNatGatewayDestinationIpPortTranslationNatRule require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to modify a NAT gateway's SNAT forwarding rules.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NATGATEWAYSNATRULENOTEXISTS = "InvalidParameterValue.NatGatewaySnatRuleNotExists"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_UNBINDEIP = "UnsupportedOperation.UnbindEIP"
func (c *Client) ModifyNatGatewaySourceIpTranslationNatRule(request *ModifyNatGatewaySourceIpTranslationNatRuleRequest) (response *ModifyNatGatewaySourceIpTranslationNatRuleResponse, err error) {
    return c.ModifyNatGatewaySourceIpTranslationNatRuleWithContext(context.Background(), request)
}

// ModifyNatGatewaySourceIpTranslationNatRule
// This API is used to modify a NAT gateway's SNAT forwarding rules.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NATGATEWAYSNATRULENOTEXISTS = "InvalidParameterValue.NatGatewaySnatRuleNotExists"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_UNBINDEIP = "UnsupportedOperation.UnbindEIP"
func (c *Client) ModifyNatGatewaySourceIpTranslationNatRuleWithContext(ctx context.Context, request *ModifyNatGatewaySourceIpTranslationNatRuleRequest) (response *ModifyNatGatewaySourceIpTranslationNatRuleResponse, err error) {
    if request == nil {
        request = NewModifyNatGatewaySourceIpTranslationNatRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNatGatewaySourceIpTranslationNatRule require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION_NETDETECTTIMEOUT = "FailedOperation.NetDetectTimeOut"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_NEXTHOPMISMATCH = "InvalidParameter.NextHopMismatch"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NETDETECTNOTFOUNDIP = "InvalidParameterValue.NetDetectNotFoundIp"
//  INVALIDPARAMETERVALUE_NETDETECTSAMEIP = "InvalidParameterValue.NetDetectSameIp"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_ECMPWITHUSERROUTE = "UnsupportedOperation.EcmpWithUserRoute"
func (c *Client) ModifyNetDetect(request *ModifyNetDetectRequest) (response *ModifyNetDetectResponse, err error) {
    return c.ModifyNetDetectWithContext(context.Background(), request)
}

// ModifyNetDetect
// This API (ModifyNetDetect) is used to modify network detection parameters.
//
// error code that may be returned:
//  FAILEDOPERATION_NETDETECTTIMEOUT = "FailedOperation.NetDetectTimeOut"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_NEXTHOPMISMATCH = "InvalidParameter.NextHopMismatch"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NETDETECTNOTFOUNDIP = "InvalidParameterValue.NetDetectNotFoundIp"
//  INVALIDPARAMETERVALUE_NETDETECTSAMEIP = "InvalidParameterValue.NetDetectSameIp"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_ECMPWITHUSERROUTE = "UnsupportedOperation.EcmpWithUserRoute"
func (c *Client) ModifyNetDetectWithContext(ctx context.Context, request *ModifyNetDetectRequest) (response *ModifyNetDetectResponse, err error) {
    if request == nil {
        request = NewModifyNetDetectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNetDetect require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ModifyNetworkAclAttributeWithContext(context.Background(), request)
}

// ModifyNetworkAclAttribute
// This API is used to modify the attributes of a network ACL.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyNetworkAclAttributeWithContext(ctx context.Context, request *ModifyNetworkAclAttributeRequest) (response *ModifyNetworkAclAttributeResponse, err error) {
    if request == nil {
        request = NewModifyNetworkAclAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNetworkAclAttribute require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to modify (add or delete) the inbound and outbound rules of a network ACL. In `NetworkAclEntrySet` parameters,
//
// * Passing in the new inbound/outbound rules will reset the original rules.
//
// * Passing in the inbound rules will only reset the original inbound rules and not affect the original outbound rules, and vice versa.
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_APPIDMISMATCH = "UnsupportedOperation.AppIdMismatch"
func (c *Client) ModifyNetworkAclEntries(request *ModifyNetworkAclEntriesRequest) (response *ModifyNetworkAclEntriesResponse, err error) {
    return c.ModifyNetworkAclEntriesWithContext(context.Background(), request)
}

// ModifyNetworkAclEntries
// This API is used to modify (add or delete) the inbound and outbound rules of a network ACL. In `NetworkAclEntrySet` parameters,
//
// * Passing in the new inbound/outbound rules will reset the original rules.
//
// * Passing in the inbound rules will only reset the original inbound rules and not affect the original outbound rules, and vice versa.
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_APPIDMISMATCH = "UnsupportedOperation.AppIdMismatch"
func (c *Client) ModifyNetworkAclEntriesWithContext(ctx context.Context, request *ModifyNetworkAclEntriesRequest) (response *ModifyNetworkAclEntriesResponse, err error) {
    if request == nil {
        request = NewModifyNetworkAclEntriesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNetworkAclEntries require credential")
    }

    request.SetContext(ctx)
    
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
//  UNSUPPORTEDOPERATION_SUBENINOTSUPPORTTRUNKING = "UnsupportedOperation.SubEniNotSupportTrunking"
func (c *Client) ModifyNetworkInterfaceAttribute(request *ModifyNetworkInterfaceAttributeRequest) (response *ModifyNetworkInterfaceAttributeResponse, err error) {
    return c.ModifyNetworkInterfaceAttributeWithContext(context.Background(), request)
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
//  UNSUPPORTEDOPERATION_SUBENINOTSUPPORTTRUNKING = "UnsupportedOperation.SubEniNotSupportTrunking"
func (c *Client) ModifyNetworkInterfaceAttributeWithContext(ctx context.Context, request *ModifyNetworkInterfaceAttributeRequest) (response *ModifyNetworkInterfaceAttributeResponse, err error) {
    if request == nil {
        request = NewModifyNetworkInterfaceAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNetworkInterfaceAttribute require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ModifyPrivateIpAddressesAttributeWithContext(context.Background(), request)
}

// ModifyPrivateIpAddressesAttribute
// This API (ModifyPrivateIpAddressesAttribute) is used to modify the private IP attributes of an ENI.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) ModifyPrivateIpAddressesAttributeWithContext(ctx context.Context, request *ModifyPrivateIpAddressesAttributeRequest) (response *ModifyPrivateIpAddressesAttributeResponse, err error) {
    if request == nil {
        request = NewModifyPrivateIpAddressesAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPrivateIpAddressesAttribute require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ModifyRouteTableAttributeWithContext(context.Background(), request)
}

// ModifyRouteTableAttribute
// This API (ModifyRouteTableAttribute) is used to modify the attributes of a route table.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyRouteTableAttributeWithContext(ctx context.Context, request *ModifyRouteTableAttributeRequest) (response *ModifyRouteTableAttributeResponse, err error) {
    if request == nil {
        request = NewModifyRouteTableAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRouteTableAttribute require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ModifySecurityGroupAttributeWithContext(context.Background(), request)
}

// ModifySecurityGroupAttribute
// This API (ModifySecurityGroupAttribute) is used to modify the attributes of a security group (SecurityGroupPolicy).
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifySecurityGroupAttributeWithContext(ctx context.Context, request *ModifySecurityGroupAttributeRequest) (response *ModifySecurityGroupAttributeResponse, err error) {
    if request == nil {
        request = NewModifySecurityGroupAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySecurityGroupAttribute require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to reset the `Egress` and `Ingress` rules (SecurityGroupPolicy) of a security group.
//
// 
//
// <ul>
//
// <li>This API does not support custom indexes <code>PolicyIndex</code>. </li>
//
// <li>For <code>SecurityGroupPolicySet</code> parameter,<ul> <ul>
//
// 	<li>If <code>SecurityGroupPolicySet.Version</code> is set to `0`, all policies will be cleared, and <code>Egress</code> and <code>Ingress</code> will be ignored. </li>
//
// 	<li>If <code>SecurityGroupPolicySet.Version</code> is not set to `0`, add <code>Egress</code> and <code>Ingress</code> policies: <ul>
//
// 		<li><code>Protocol</code>: <code>TCP</code>, <code>UDP</code>, <code>ICMP</code>, <code>ICMPV6</code>, <code>GRE</code>, or <code>ALL</code>. </li>
//
// 		<li><code>CidrBlock</code>: a CIDR block in the correct format. In the classic network, even if the CIDR block specified in <code>CidrBlock</code> contains the Tencent Cloud private IPs that are not using for CVMs under your Tencent Cloud account, it does not mean this policy allows you to access those resources. The network isolation policies between tenants take priority over the private network policies in security groups. </li>
//
// 		<li><code>Ipv6CidrBlock</code>: an IPv6 CIDR block in the correct format. In the classic network, even if the CIDR block specified in <code>Ipv6CidrBlock</code> contains the Tencent Cloud private IPv6 addresses that are not using for CVMs under your Tencent Cloud account, it does not mean this policy allows you to access those resources. The network isolation policies between tenants take priority over the private network policies in security groups. </li>
//
// 		<li><code>SecurityGroupId</code>: ID of the security group. It can be the ID of a security group to be modified, or the ID of another security group in the same project. All private IPs of all CVMs under the security group will be covered. If this field is used, the policy will automatically change according to the CVM associated with the group ID while being used to match network messages. You don't need to change it manually. </li>
//
// 		<li><code>Port</code>: a single port number such as 80, or a port range in the format of '8000-8010'.  You may use this field only if the <code>Protocol</code> field takes the value <code>TCP</code> or <code>UDP</code>. </li>
//
// 		<li><code>Action</code>: only allows <code>ACCEPT</code> or <code>DROP</code>. </li>
//
// 		<li><code>CidrBlock</code>, <code>Ipv6CidrBlock</code>, <code>SecurityGroupId</code>, and <code>AddressTemplate</code> are mutually exclusive. <code>Protocol</code> + <code>Port</code> and <code>ServiceTemplate</code> are mutually exclusive.</li> </li>
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
    return c.ModifySecurityGroupPoliciesWithContext(context.Background(), request)
}

// ModifySecurityGroupPolicies
// This API is used to reset the `Egress` and `Ingress` rules (SecurityGroupPolicy) of a security group.
//
// 
//
// <ul>
//
// <li>This API does not support custom indexes <code>PolicyIndex</code>. </li>
//
// <li>For <code>SecurityGroupPolicySet</code> parameter,<ul> <ul>
//
// 	<li>If <code>SecurityGroupPolicySet.Version</code> is set to `0`, all policies will be cleared, and <code>Egress</code> and <code>Ingress</code> will be ignored. </li>
//
// 	<li>If <code>SecurityGroupPolicySet.Version</code> is not set to `0`, add <code>Egress</code> and <code>Ingress</code> policies: <ul>
//
// 		<li><code>Protocol</code>: <code>TCP</code>, <code>UDP</code>, <code>ICMP</code>, <code>ICMPV6</code>, <code>GRE</code>, or <code>ALL</code>. </li>
//
// 		<li><code>CidrBlock</code>: a CIDR block in the correct format. In the classic network, even if the CIDR block specified in <code>CidrBlock</code> contains the Tencent Cloud private IPs that are not using for CVMs under your Tencent Cloud account, it does not mean this policy allows you to access those resources. The network isolation policies between tenants take priority over the private network policies in security groups. </li>
//
// 		<li><code>Ipv6CidrBlock</code>: an IPv6 CIDR block in the correct format. In the classic network, even if the CIDR block specified in <code>Ipv6CidrBlock</code> contains the Tencent Cloud private IPv6 addresses that are not using for CVMs under your Tencent Cloud account, it does not mean this policy allows you to access those resources. The network isolation policies between tenants take priority over the private network policies in security groups. </li>
//
// 		<li><code>SecurityGroupId</code>: ID of the security group. It can be the ID of a security group to be modified, or the ID of another security group in the same project. All private IPs of all CVMs under the security group will be covered. If this field is used, the policy will automatically change according to the CVM associated with the group ID while being used to match network messages. You don't need to change it manually. </li>
//
// 		<li><code>Port</code>: a single port number such as 80, or a port range in the format of '8000-8010'.  You may use this field only if the <code>Protocol</code> field takes the value <code>TCP</code> or <code>UDP</code>. </li>
//
// 		<li><code>Action</code>: only allows <code>ACCEPT</code> or <code>DROP</code>. </li>
//
// 		<li><code>CidrBlock</code>, <code>Ipv6CidrBlock</code>, <code>SecurityGroupId</code>, and <code>AddressTemplate</code> are mutually exclusive. <code>Protocol</code> + <code>Port</code> and <code>ServiceTemplate</code> are mutually exclusive.</li> </li>
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
func (c *Client) ModifySecurityGroupPoliciesWithContext(ctx context.Context, request *ModifySecurityGroupPoliciesRequest) (response *ModifySecurityGroupPoliciesResponse, err error) {
    if request == nil {
        request = NewModifySecurityGroupPoliciesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySecurityGroupPolicies require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) ModifyServiceTemplateAttribute(request *ModifyServiceTemplateAttributeRequest) (response *ModifyServiceTemplateAttributeResponse, err error) {
    return c.ModifyServiceTemplateAttributeWithContext(context.Background(), request)
}

// ModifyServiceTemplateAttribute
// This API (ModifyServiceTemplateAttribute) is used to modify a protocol port template.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) ModifyServiceTemplateAttributeWithContext(ctx context.Context, request *ModifyServiceTemplateAttributeRequest) (response *ModifyServiceTemplateAttributeResponse, err error) {
    if request == nil {
        request = NewModifyServiceTemplateAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyServiceTemplateAttribute require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ModifyServiceTemplateGroupAttributeWithContext(context.Background(), request)
}

// ModifyServiceTemplateGroupAttribute
// This API (ModifyServiceTemplateGroupAttribute) is used to modify a protocol port template group.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyServiceTemplateGroupAttributeWithContext(ctx context.Context, request *ModifyServiceTemplateGroupAttributeRequest) (response *ModifyServiceTemplateGroupAttributeResponse, err error) {
    if request == nil {
        request = NewModifyServiceTemplateGroupAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyServiceTemplateGroupAttribute require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ModifySubnetAttributeWithContext(context.Background(), request)
}

// ModifySubnetAttribute
// This API (ModifySubnetAttribute) is used to modify subnet attributes.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifySubnetAttributeWithContext(ctx context.Context, request *ModifySubnetAttributeRequest) (response *ModifySubnetAttributeResponse, err error) {
    if request == nil {
        request = NewModifySubnetAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySubnetAttribute require credential")
    }

    request.SetContext(ctx)
    
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
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_NOTSUPPORTEDUPDATECCNROUTEPUBLISH = "UnsupportedOperation.NotSupportedUpdateCcnRoutePublish"
func (c *Client) ModifyVpcAttribute(request *ModifyVpcAttributeRequest) (response *ModifyVpcAttributeResponse, err error) {
    return c.ModifyVpcAttributeWithContext(context.Background(), request)
}

// ModifyVpcAttribute
// This API (ModifyVpcAttribute) is used to modify VPC attributes.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_NOTSUPPORTEDUPDATECCNROUTEPUBLISH = "UnsupportedOperation.NotSupportedUpdateCcnRoutePublish"
func (c *Client) ModifyVpcAttributeWithContext(ctx context.Context, request *ModifyVpcAttributeRequest) (response *ModifyVpcAttributeResponse, err error) {
    if request == nil {
        request = NewModifyVpcAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVpcAttribute require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SVCNOTEXIST = "ResourceNotFound.SvcNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_SPECIALENDPOINTSERVICE = "UnsupportedOperation.SpecialEndPointService"
func (c *Client) ModifyVpcEndPointAttribute(request *ModifyVpcEndPointAttributeRequest) (response *ModifyVpcEndPointAttributeResponse, err error) {
    return c.ModifyVpcEndPointAttributeWithContext(context.Background(), request)
}

// ModifyVpcEndPointAttribute
// This API is used to modify endpoint attributes.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SVCNOTEXIST = "ResourceNotFound.SvcNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_SPECIALENDPOINTSERVICE = "UnsupportedOperation.SpecialEndPointService"
func (c *Client) ModifyVpcEndPointAttributeWithContext(ctx context.Context, request *ModifyVpcEndPointAttributeRequest) (response *ModifyVpcEndPointAttributeResponse, err error) {
    if request == nil {
        request = NewModifyVpcEndPointAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVpcEndPointAttribute require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to modify the VPC endpoint service attributes.
//
// 
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
func (c *Client) ModifyVpcEndPointServiceAttribute(request *ModifyVpcEndPointServiceAttributeRequest) (response *ModifyVpcEndPointServiceAttributeResponse, err error) {
    return c.ModifyVpcEndPointServiceAttributeWithContext(context.Background(), request)
}

// ModifyVpcEndPointServiceAttribute
// This API is used to modify the VPC endpoint service attributes.
//
// 
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
func (c *Client) ModifyVpcEndPointServiceAttributeWithContext(ctx context.Context, request *ModifyVpcEndPointServiceAttributeRequest) (response *ModifyVpcEndPointServiceAttributeResponse, err error) {
    if request == nil {
        request = NewModifyVpcEndPointServiceAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVpcEndPointServiceAttribute require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ModifyVpcEndPointServiceWhiteListWithContext(context.Background(), request)
}

// ModifyVpcEndPointServiceWhiteList
// This API is used to modify the attributes of the endpoint service allowlist.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_UINNOTFOUND = "UnsupportedOperation.UinNotFound"
func (c *Client) ModifyVpcEndPointServiceWhiteListWithContext(ctx context.Context, request *ModifyVpcEndPointServiceWhiteListRequest) (response *ModifyVpcEndPointServiceWhiteListResponse, err error) {
    if request == nil {
        request = NewModifyVpcEndPointServiceWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVpcEndPointServiceWhiteList require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_VPNCONNHEALTHCHECKIPCONFLICT = "InvalidParameterValue.VpnConnHealthCheckIpConflict"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyVpnConnectionAttribute(request *ModifyVpnConnectionAttributeRequest) (response *ModifyVpnConnectionAttributeResponse, err error) {
    return c.ModifyVpnConnectionAttributeWithContext(context.Background(), request)
}

// ModifyVpnConnectionAttribute
// This API (ModifyVpnConnectionAttribute) is used to modify VPN tunnels.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_VPNCONNCIDRCONFLICT = "InvalidParameterValue.VpnConnCidrConflict"
//  INVALIDPARAMETERVALUE_VPNCONNHEALTHCHECKIPCONFLICT = "InvalidParameterValue.VpnConnHealthCheckIpConflict"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyVpnConnectionAttributeWithContext(ctx context.Context, request *ModifyVpnConnectionAttributeRequest) (response *ModifyVpnConnectionAttributeResponse, err error) {
    if request == nil {
        request = NewModifyVpnConnectionAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVpnConnectionAttribute require credential")
    }

    request.SetContext(ctx)
    
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
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) ModifyVpnGatewayAttribute(request *ModifyVpnGatewayAttributeRequest) (response *ModifyVpnGatewayAttributeResponse, err error) {
    return c.ModifyVpnGatewayAttributeWithContext(context.Background(), request)
}

// ModifyVpnGatewayAttribute
// This API (ModifyVpnGatewayAttribute) is used to modify the attributes of VPN gateways.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) ModifyVpnGatewayAttributeWithContext(ctx context.Context, request *ModifyVpnGatewayAttributeRequest) (response *ModifyVpnGatewayAttributeResponse, err error) {
    if request == nil {
        request = NewModifyVpnGatewayAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVpnGatewayAttribute require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ModifyVpnGatewayCcnRoutesWithContext(context.Background(), request)
}

// ModifyVpnGatewayCcnRoutes
// This API (ModifyVpnGatewayCcnRoutes) is used to modify VPN gateway-based CCN routes.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyVpnGatewayCcnRoutesWithContext(ctx context.Context, request *ModifyVpnGatewayCcnRoutesRequest) (response *ModifyVpnGatewayCcnRoutesResponse, err error) {
    if request == nil {
        request = NewModifyVpnGatewayCcnRoutesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVpnGatewayCcnRoutes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVpnGatewayCcnRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVpnGatewayRoutesRequest() (request *ModifyVpnGatewayRoutesRequest) {
    request = &ModifyVpnGatewayRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyVpnGatewayRoutes")
    
    
    return
}

func NewModifyVpnGatewayRoutesResponse() (response *ModifyVpnGatewayRoutesResponse) {
    response = &ModifyVpnGatewayRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyVpnGatewayRoutes
// This API is used to modify the route status of a VPN gateway.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyVpnGatewayRoutes(request *ModifyVpnGatewayRoutesRequest) (response *ModifyVpnGatewayRoutesResponse, err error) {
    return c.ModifyVpnGatewayRoutesWithContext(context.Background(), request)
}

// ModifyVpnGatewayRoutes
// This API is used to modify the route status of a VPN gateway.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyVpnGatewayRoutesWithContext(ctx context.Context, request *ModifyVpnGatewayRoutesRequest) (response *ModifyVpnGatewayRoutesResponse, err error) {
    if request == nil {
        request = NewModifyVpnGatewayRoutesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVpnGatewayRoutes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVpnGatewayRoutesResponse()
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
// This API is used to publish a route to CCN. This can also be done by clicking "Publish to CCN" in the operation column on the page of route table list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDROUTEID_NOTFOUND = "InvalidRouteId.NotFound"
//  INVALIDROUTETABLEID_MALFORMED = "InvalidRouteTableId.Malformed"
//  INVALIDROUTETABLEID_NOTFOUND = "InvalidRouteTableId.NotFound"
//  UNSUPPORTEDOPERATION_CCNNOTATTACHED = "UnsupportedOperation.CcnNotAttached"
//  UNSUPPORTEDOPERATION_INVALIDSTATUSNOTIFYCCN = "UnsupportedOperation.InvalidStatusNotifyCcn"
//  UNSUPPORTEDOPERATION_NOTIFYCCN = "UnsupportedOperation.NotifyCcn"
//  UNSUPPORTEDOPERATION_SYSTEMROUTE = "UnsupportedOperation.SystemRoute"
func (c *Client) NotifyRoutes(request *NotifyRoutesRequest) (response *NotifyRoutesResponse, err error) {
    return c.NotifyRoutesWithContext(context.Background(), request)
}

// NotifyRoutes
// This API is used to publish a route to CCN. This can also be done by clicking "Publish to CCN" in the operation column on the page of route table list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDROUTEID_NOTFOUND = "InvalidRouteId.NotFound"
//  INVALIDROUTETABLEID_MALFORMED = "InvalidRouteTableId.Malformed"
//  INVALIDROUTETABLEID_NOTFOUND = "InvalidRouteTableId.NotFound"
//  UNSUPPORTEDOPERATION_CCNNOTATTACHED = "UnsupportedOperation.CcnNotAttached"
//  UNSUPPORTEDOPERATION_INVALIDSTATUSNOTIFYCCN = "UnsupportedOperation.InvalidStatusNotifyCcn"
//  UNSUPPORTEDOPERATION_NOTIFYCCN = "UnsupportedOperation.NotifyCcn"
//  UNSUPPORTEDOPERATION_SYSTEMROUTE = "UnsupportedOperation.SystemRoute"
func (c *Client) NotifyRoutesWithContext(ctx context.Context, request *NotifyRoutesRequest) (response *NotifyRoutesResponse, err error) {
    if request == nil {
        request = NewNotifyRoutesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("NotifyRoutes require credential")
    }

    request.SetContext(ctx)
    
    response = NewNotifyRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewRefreshDirectConnectGatewayRouteToNatGatewayRequest() (request *RefreshDirectConnectGatewayRouteToNatGatewayRequest) {
    request = &RefreshDirectConnectGatewayRouteToNatGatewayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "RefreshDirectConnectGatewayRouteToNatGateway")
    
    
    return
}

func NewRefreshDirectConnectGatewayRouteToNatGatewayResponse() (response *RefreshDirectConnectGatewayRouteToNatGatewayResponse) {
    response = &RefreshDirectConnectGatewayRouteToNatGatewayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RefreshDirectConnectGatewayRouteToNatGateway
// This API is used to refresh the route between a NAT gateway and  Direct Connect and update the associated route table.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) RefreshDirectConnectGatewayRouteToNatGateway(request *RefreshDirectConnectGatewayRouteToNatGatewayRequest) (response *RefreshDirectConnectGatewayRouteToNatGatewayResponse, err error) {
    return c.RefreshDirectConnectGatewayRouteToNatGatewayWithContext(context.Background(), request)
}

// RefreshDirectConnectGatewayRouteToNatGateway
// This API is used to refresh the route between a NAT gateway and  Direct Connect and update the associated route table.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) RefreshDirectConnectGatewayRouteToNatGatewayWithContext(ctx context.Context, request *RefreshDirectConnectGatewayRouteToNatGatewayRequest) (response *RefreshDirectConnectGatewayRouteToNatGatewayResponse, err error) {
    if request == nil {
        request = NewRefreshDirectConnectGatewayRouteToNatGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RefreshDirectConnectGatewayRouteToNatGateway require credential")
    }

    request.SetContext(ctx)
    
    response = NewRefreshDirectConnectGatewayRouteToNatGatewayResponse()
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
    return c.RejectAttachCcnInstancesWithContext(context.Background(), request)
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
func (c *Client) RejectAttachCcnInstancesWithContext(ctx context.Context, request *RejectAttachCcnInstancesRequest) (response *RejectAttachCcnInstancesResponse, err error) {
    if request == nil {
        request = NewRejectAttachCcnInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RejectAttachCcnInstances require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_ADDRESSIDMALFORMED = "InvalidParameterValue.AddressIdMalformed"
//  INVALIDPARAMETERVALUE_ADDRESSINTERNETCHARGETYPECONFLICT = "InvalidParameterValue.AddressInternetChargeTypeConflict"
//  INVALIDPARAMETERVALUE_ADDRESSNOTEIP = "InvalidParameterValue.AddressNotEIP"
//  INVALIDPARAMETERVALUE_ADDRESSNOTFOUND = "InvalidParameterValue.AddressNotFound"
//  LIMITEXCEEDED_ACCOUNTRETURNQUOTA = "LimitExceeded.AccountReturnQuota"
//  OPERATIONDENIED_MUTEXTASKRUNNING = "OperationDenied.MutexTaskRunning"
//  UNSUPPORTEDOPERATION_ADDRESSSTATUSNOTPERMIT = "UnsupportedOperation.AddressStatusNotPermit"
func (c *Client) ReleaseAddresses(request *ReleaseAddressesRequest) (response *ReleaseAddressesResponse, err error) {
    return c.ReleaseAddressesWithContext(context.Background(), request)
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
//  INVALIDPARAMETERVALUE_ADDRESSIDMALFORMED = "InvalidParameterValue.AddressIdMalformed"
//  INVALIDPARAMETERVALUE_ADDRESSINTERNETCHARGETYPECONFLICT = "InvalidParameterValue.AddressInternetChargeTypeConflict"
//  INVALIDPARAMETERVALUE_ADDRESSNOTEIP = "InvalidParameterValue.AddressNotEIP"
//  INVALIDPARAMETERVALUE_ADDRESSNOTFOUND = "InvalidParameterValue.AddressNotFound"
//  LIMITEXCEEDED_ACCOUNTRETURNQUOTA = "LimitExceeded.AccountReturnQuota"
//  OPERATIONDENIED_MUTEXTASKRUNNING = "OperationDenied.MutexTaskRunning"
//  UNSUPPORTEDOPERATION_ADDRESSSTATUSNOTPERMIT = "UnsupportedOperation.AddressStatusNotPermit"
func (c *Client) ReleaseAddressesWithContext(ctx context.Context, request *ReleaseAddressesRequest) (response *ReleaseAddressesResponse, err error) {
    if request == nil {
        request = NewReleaseAddressesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReleaseAddresses require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGEIDMALFORMED = "InvalidParameterValue.BandwidthPackageIdMalformed"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGENOTFOUND = "InvalidParameterValue.BandwidthPackageNotFound"
//  INVALIDPARAMETERVALUE_RESOURCEIDMALFORMED = "InvalidParameterValue.ResourceIdMalformed"
//  INVALIDPARAMETERVALUE_RESOURCENOTEXISTED = "InvalidParameterValue.ResourceNotExisted"
//  INVALIDPARAMETERVALUE_RESOURCENOTFOUND = "InvalidParameterValue.ResourceNotFound"
//  UNSUPPORTEDOPERATION_BANDWIDTHPACKAGEIDNOTSUPPORTED = "UnsupportedOperation.BandwidthPackageIdNotSupported"
//  UNSUPPORTEDOPERATION_INVALIDRESOURCEPROTOCOL = "UnsupportedOperation.InvalidResourceProtocol"
func (c *Client) RemoveBandwidthPackageResources(request *RemoveBandwidthPackageResourcesRequest) (response *RemoveBandwidthPackageResourcesResponse, err error) {
    return c.RemoveBandwidthPackageResourcesWithContext(context.Background(), request)
}

// RemoveBandwidthPackageResources
// This API is used to delete a bandwidth package resource, including [Elastic IP](https://intl.cloud.tencent.com/document/product/213/1941?from_cn_redirect=1), [Cloud Load Balancer](https://intl.cloud.tencent.com/document/product/214/517?from_cn_redirect=1), and so on.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGEIDMALFORMED = "InvalidParameterValue.BandwidthPackageIdMalformed"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGENOTFOUND = "InvalidParameterValue.BandwidthPackageNotFound"
//  INVALIDPARAMETERVALUE_RESOURCEIDMALFORMED = "InvalidParameterValue.ResourceIdMalformed"
//  INVALIDPARAMETERVALUE_RESOURCENOTEXISTED = "InvalidParameterValue.ResourceNotExisted"
//  INVALIDPARAMETERVALUE_RESOURCENOTFOUND = "InvalidParameterValue.ResourceNotFound"
//  UNSUPPORTEDOPERATION_BANDWIDTHPACKAGEIDNOTSUPPORTED = "UnsupportedOperation.BandwidthPackageIdNotSupported"
//  UNSUPPORTEDOPERATION_INVALIDRESOURCEPROTOCOL = "UnsupportedOperation.InvalidResourceProtocol"
func (c *Client) RemoveBandwidthPackageResourcesWithContext(ctx context.Context, request *RemoveBandwidthPackageResourcesRequest) (response *RemoveBandwidthPackageResourcesResponse, err error) {
    if request == nil {
        request = NewRemoveBandwidthPackageResourcesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveBandwidthPackageResources require credential")
    }

    request.SetContext(ctx)
    
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
    return c.RenewVpnGatewayWithContext(context.Background(), request)
}

// RenewVpnGateway
// This API (RenewVpnGateway) is used to renew prepaid (monthly subscription) VPN gateways. Currently, only IPSEC gateways are supported.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RenewVpnGatewayWithContext(ctx context.Context, request *RenewVpnGatewayRequest) (response *RenewVpnGatewayResponse, err error) {
    if request == nil {
        request = NewRenewVpnGatewayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RenewVpnGateway require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ReplaceDirectConnectGatewayCcnRoutes(request *ReplaceDirectConnectGatewayCcnRoutesRequest) (response *ReplaceDirectConnectGatewayCcnRoutesResponse, err error) {
    return c.ReplaceDirectConnectGatewayCcnRoutesWithContext(context.Background(), request)
}

// ReplaceDirectConnectGatewayCcnRoutes
// This API (ReplaceDirectConnectGatewayCcnRoutes) is used to modify the specified route according to the route ID. Batch modification is supported.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ReplaceDirectConnectGatewayCcnRoutesWithContext(ctx context.Context, request *ReplaceDirectConnectGatewayCcnRoutesRequest) (response *ReplaceDirectConnectGatewayCcnRoutesResponse, err error) {
    if request == nil {
        request = NewReplaceDirectConnectGatewayCcnRoutesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReplaceDirectConnectGatewayCcnRoutes require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ReplaceRouteTableAssociationWithContext(context.Background(), request)
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
func (c *Client) ReplaceRouteTableAssociationWithContext(ctx context.Context, request *ReplaceRouteTableAssociationRequest) (response *ReplaceRouteTableAssociationResponse, err error) {
    if request == nil {
        request = NewReplaceRouteTableAssociationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReplaceRouteTableAssociation require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ReplaceRoutesWithContext(context.Background(), request)
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
func (c *Client) ReplaceRoutesWithContext(ctx context.Context, request *ReplaceRoutesRequest) (response *ReplaceRoutesResponse, err error) {
    if request == nil {
        request = NewReplaceRoutesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReplaceRoutes require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_CLBPOLICYLIMIT = "UnsupportedOperation.ClbPolicyLimit"
//  UNSUPPORTEDOPERATION_DUPLICATEPOLICY = "UnsupportedOperation.DuplicatePolicy"
//  UNSUPPORTEDOPERATION_VERSIONMISMATCH = "UnsupportedOperation.VersionMismatch"
func (c *Client) ReplaceSecurityGroupPolicy(request *ReplaceSecurityGroupPolicyRequest) (response *ReplaceSecurityGroupPolicyResponse, err error) {
    return c.ReplaceSecurityGroupPolicyWithContext(context.Background(), request)
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
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_CLBPOLICYLIMIT = "UnsupportedOperation.ClbPolicyLimit"
//  UNSUPPORTEDOPERATION_DUPLICATEPOLICY = "UnsupportedOperation.DuplicatePolicy"
//  UNSUPPORTEDOPERATION_VERSIONMISMATCH = "UnsupportedOperation.VersionMismatch"
func (c *Client) ReplaceSecurityGroupPolicyWithContext(ctx context.Context, request *ReplaceSecurityGroupPolicyRequest) (response *ReplaceSecurityGroupPolicyResponse, err error) {
    if request == nil {
        request = NewReplaceSecurityGroupPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReplaceSecurityGroupPolicy require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ResetAttachCcnInstancesWithContext(context.Background(), request)
}

// ResetAttachCcnInstances
// This API (ResetAttachCcnInstances) is used to re-apply for the association operation when the application for cross-account instance association expires.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ResetAttachCcnInstancesWithContext(ctx context.Context, request *ResetAttachCcnInstancesRequest) (response *ResetAttachCcnInstancesResponse, err error) {
    if request == nil {
        request = NewResetAttachCcnInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetAttachCcnInstances require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to adjust concurrent connection cap for the NAT gateway.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_UNPAIDORDERALREADYEXISTS = "UnsupportedOperation.UnpaidOrderAlreadyExists"
func (c *Client) ResetNatGatewayConnection(request *ResetNatGatewayConnectionRequest) (response *ResetNatGatewayConnectionResponse, err error) {
    return c.ResetNatGatewayConnectionWithContext(context.Background(), request)
}

// ResetNatGatewayConnection
// This API is used to adjust concurrent connection cap for the NAT gateway.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_UNPAIDORDERALREADYEXISTS = "UnsupportedOperation.UnpaidOrderAlreadyExists"
func (c *Client) ResetNatGatewayConnectionWithContext(ctx context.Context, request *ResetNatGatewayConnectionRequest) (response *ResetNatGatewayConnectionResponse, err error) {
    if request == nil {
        request = NewResetNatGatewayConnectionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetNatGatewayConnection require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ResetRoutesWithContext(context.Background(), request)
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
func (c *Client) ResetRoutesWithContext(ctx context.Context, request *ResetRoutesRequest) (response *ResetRoutesResponse, err error) {
    if request == nil {
        request = NewResetRoutesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetRoutes require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ResetVpnConnectionWithContext(context.Background(), request)
}

// ResetVpnConnection
// The API (ResetVpnConnection) is used to reset VPN tunnels.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ResetVpnConnectionWithContext(ctx context.Context, request *ResetVpnConnectionRequest) (response *ResetVpnConnectionResponse, err error) {
    if request == nil {
        request = NewResetVpnConnectionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetVpnConnection require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ResetVpnGatewayInternetMaxBandwidthWithContext(context.Background(), request)
}

// ResetVpnGatewayInternetMaxBandwidth
// This API (ResetVpnGatewayInternetMaxBandwidth) is used to adjust the bandwidth cap of VPN gateways. Currently, only configuration upgrade is supported. VPN gateways with monthly subscription must be within the validity period.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) ResetVpnGatewayInternetMaxBandwidthWithContext(ctx context.Context, request *ResetVpnGatewayInternetMaxBandwidthRequest) (response *ResetVpnGatewayInternetMaxBandwidthResponse, err error) {
    if request == nil {
        request = NewResetVpnGatewayInternetMaxBandwidthRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetVpnGatewayInternetMaxBandwidth require credential")
    }

    request.SetContext(ctx)
    
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
    return c.SetCcnRegionBandwidthLimitsWithContext(context.Background(), request)
}

// SetCcnRegionBandwidthLimits
// This API (SetCcnRegionBandwidthLimits) is used to set the outbound bandwidth cap for CCNs in each region. This API can only set the outbound bandwidth cap for regions in the network instances that have already been associated.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_NOTPOSTPAIDCCNOPERATION = "UnsupportedOperation.NotPostpaidCcnOperation"
func (c *Client) SetCcnRegionBandwidthLimitsWithContext(ctx context.Context, request *SetCcnRegionBandwidthLimitsRequest) (response *SetCcnRegionBandwidthLimitsResponse, err error) {
    if request == nil {
        request = NewSetCcnRegionBandwidthLimitsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetCcnRegionBandwidthLimits require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to convert a common public IP into an [Elastic IP](https://intl.cloud.tencent.com/document/product/213/1941?from_cn_redirect=1) (EIP for short).
//
// * Tencent Cloud limits the number of times that a user can unbind EIPs and reassign public IPs in each region per day. For more information, see product introduction of [Elastic IP](https://intl.cloud.tencent.com/document/product/213/5733?from_cn_redirect=1). The preceding quota can be obtained through the API [DescribeAddressQuota](https://intl.cloud.tencent.com/document/product/215/16701).
//
// error code that may be returned:
//  ADDRESSQUOTALIMITEXCEEDED = "AddressQuotaLimitExceeded"
//  ADDRESSQUOTALIMITEXCEEDED_DAILYALLOCATE = "AddressQuotaLimitExceeded.DailyAllocate"
//  INVALIDADDRESSID_BLOCKED = "InvalidAddressId.Blocked"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_ALREADYBINDEIP = "InvalidInstanceId.AlreadyBindEip"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INSTANCENOWANIP = "InvalidParameterValue.InstanceNoWanIP"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCESTATE = "InvalidParameterValue.InvalidInstanceState"
//  LIMITEXCEEDED_MONTHLYADDRESSRECOVERYQUOTA = "LimitExceeded.MonthlyAddressRecoveryQuota"
//  OPERATIONDENIED_MUTEXTASKRUNNING = "OperationDenied.MutexTaskRunning"
//  UNSUPPORTEDOPERATION_ADDRESSSTATUSNOTPERMIT = "UnsupportedOperation.AddressStatusNotPermit"
//  UNSUPPORTEDOPERATION_INVALIDADDRESSINTERNETCHARGETYPE = "UnsupportedOperation.InvalidAddressInternetChargeType"
func (c *Client) TransformAddress(request *TransformAddressRequest) (response *TransformAddressResponse, err error) {
    return c.TransformAddressWithContext(context.Background(), request)
}

// TransformAddress
// This API is used to convert a common public IP into an [Elastic IP](https://intl.cloud.tencent.com/document/product/213/1941?from_cn_redirect=1) (EIP for short).
//
// * Tencent Cloud limits the number of times that a user can unbind EIPs and reassign public IPs in each region per day. For more information, see product introduction of [Elastic IP](https://intl.cloud.tencent.com/document/product/213/5733?from_cn_redirect=1). The preceding quota can be obtained through the API [DescribeAddressQuota](https://intl.cloud.tencent.com/document/product/215/16701).
//
// error code that may be returned:
//  ADDRESSQUOTALIMITEXCEEDED = "AddressQuotaLimitExceeded"
//  ADDRESSQUOTALIMITEXCEEDED_DAILYALLOCATE = "AddressQuotaLimitExceeded.DailyAllocate"
//  INVALIDADDRESSID_BLOCKED = "InvalidAddressId.Blocked"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_ALREADYBINDEIP = "InvalidInstanceId.AlreadyBindEip"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INSTANCENOWANIP = "InvalidParameterValue.InstanceNoWanIP"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCESTATE = "InvalidParameterValue.InvalidInstanceState"
//  LIMITEXCEEDED_MONTHLYADDRESSRECOVERYQUOTA = "LimitExceeded.MonthlyAddressRecoveryQuota"
//  OPERATIONDENIED_MUTEXTASKRUNNING = "OperationDenied.MutexTaskRunning"
//  UNSUPPORTEDOPERATION_ADDRESSSTATUSNOTPERMIT = "UnsupportedOperation.AddressStatusNotPermit"
//  UNSUPPORTEDOPERATION_INVALIDADDRESSINTERNETCHARGETYPE = "UnsupportedOperation.InvalidAddressInternetChargeType"
func (c *Client) TransformAddressWithContext(ctx context.Context, request *TransformAddressRequest) (response *TransformAddressResponse, err error) {
    if request == nil {
        request = NewTransformAddressRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TransformAddress require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to release the IPv6 addresses of an ENI. <br />
//
// This API is completed asynchronously. If you need to query the execution result of an async task, please use the `RequestId` returned by this API to poll the `DescribeVpcTaskResult` API.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_ATTACHMENTNOTFOUND = "UnauthorizedOperation.AttachmentNotFound"
//  UNAUTHORIZEDOPERATION_PRIMARYIP = "UnauthorizedOperation.PrimaryIp"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ATTACHMENTNOTFOUND = "UnsupportedOperation.AttachmentNotFound"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) UnassignIpv6Addresses(request *UnassignIpv6AddressesRequest) (response *UnassignIpv6AddressesResponse, err error) {
    return c.UnassignIpv6AddressesWithContext(context.Background(), request)
}

// UnassignIpv6Addresses
// This API is used to release the IPv6 addresses of an ENI. <br />
//
// This API is completed asynchronously. If you need to query the execution result of an async task, please use the `RequestId` returned by this API to poll the `DescribeVpcTaskResult` API.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_ATTACHMENTNOTFOUND = "UnauthorizedOperation.AttachmentNotFound"
//  UNAUTHORIZEDOPERATION_PRIMARYIP = "UnauthorizedOperation.PrimaryIp"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ATTACHMENTNOTFOUND = "UnsupportedOperation.AttachmentNotFound"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) UnassignIpv6AddressesWithContext(ctx context.Context, request *UnassignIpv6AddressesRequest) (response *UnassignIpv6AddressesResponse, err error) {
    if request == nil {
        request = NewUnassignIpv6AddressesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnassignIpv6Addresses require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UnassignIpv6CidrBlock(request *UnassignIpv6CidrBlockRequest) (response *UnassignIpv6CidrBlockResponse, err error) {
    return c.UnassignIpv6CidrBlockWithContext(context.Background(), request)
}

// UnassignIpv6CidrBlock
// This API (UnassignIpv6CidrBlock) is used to release IPv6 IP ranges.
//
// If the IP range still has occupied IPs that are not yet repossessed, the IP range cannot be released.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UnassignIpv6CidrBlockWithContext(ctx context.Context, request *UnassignIpv6CidrBlockRequest) (response *UnassignIpv6CidrBlockResponse, err error) {
    if request == nil {
        request = NewUnassignIpv6CidrBlockRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnassignIpv6CidrBlock require credential")
    }

    request.SetContext(ctx)
    
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
    return c.UnassignIpv6SubnetCidrBlockWithContext(context.Background(), request)
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
func (c *Client) UnassignIpv6SubnetCidrBlockWithContext(ctx context.Context, request *UnassignIpv6SubnetCidrBlockRequest) (response *UnassignIpv6SubnetCidrBlockResponse, err error) {
    if request == nil {
        request = NewUnassignIpv6SubnetCidrBlockRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnassignIpv6SubnetCidrBlock require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to return the private IP addresses of an ENI.
//
// * If a secondary private IP of an ENI is returned, the EIP will be automatically unassociated as well. The primary private IP of the ENI cannot be returned.
//
// 
//
// This API is completed asynchronously. If you need to query the execution result of an async task, please use the `RequestId` returned by this API to poll the `DescribeVpcTaskResult` API.
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
    return c.UnassignPrivateIpAddressesWithContext(context.Background(), request)
}

// UnassignPrivateIpAddresses
// This API is used to return the private IP addresses of an ENI.
//
// * If a secondary private IP of an ENI is returned, the EIP will be automatically unassociated as well. The primary private IP of the ENI cannot be returned.
//
// 
//
// This API is completed asynchronously. If you need to query the execution result of an async task, please use the `RequestId` returned by this API to poll the `DescribeVpcTaskResult` API.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ATTACHMENTNOTFOUND = "UnsupportedOperation.AttachmentNotFound"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) UnassignPrivateIpAddressesWithContext(ctx context.Context, request *UnassignPrivateIpAddressesRequest) (response *UnassignPrivateIpAddressesResponse, err error) {
    if request == nil {
        request = NewUnassignPrivateIpAddressesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnassignPrivateIpAddresses require credential")
    }

    request.SetContext(ctx)
    
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
// This API is used to withdraw a route from CCN. 
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
    return c.WithdrawNotifyRoutesWithContext(context.Background(), request)
}

// WithdrawNotifyRoutes
// This API is used to withdraw a route from CCN. 
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
func (c *Client) WithdrawNotifyRoutesWithContext(ctx context.Context, request *WithdrawNotifyRoutesRequest) (response *WithdrawNotifyRoutesResponse, err error) {
    if request == nil {
        request = NewWithdrawNotifyRoutesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("WithdrawNotifyRoutes require credential")
    }

    request.SetContext(ctx)
    
    response = NewWithdrawNotifyRoutesResponse()
    err = c.Send(request, response)
    return
}
