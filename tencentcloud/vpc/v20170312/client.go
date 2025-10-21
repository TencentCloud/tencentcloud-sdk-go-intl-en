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

package v20170312

import (
    "context"
    "errors"
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
//  UNSUPPORTEDOPERATION_UNABLECROSSBORDER = "UnsupportedOperation.UnableCrossBorder"
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
//  UNSUPPORTEDOPERATION_UNABLECROSSBORDER = "UnsupportedOperation.UnableCrossBorder"
//  UNSUPPORTEDOPERATION_UNABLECROSSFINANCE = "UnsupportedOperation.UnableCrossFinance"
func (c *Client) AcceptAttachCcnInstancesWithContext(ctx context.Context, request *AcceptAttachCcnInstancesRequest) (response *AcceptAttachCcnInstancesResponse, err error) {
    if request == nil {
        request = NewAcceptAttachCcnInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "AcceptAttachCcnInstances")
    
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
//  LIMITEXCEEDED_BANDWIDTHPACKAGERESOURCEQUOTA = "LimitExceeded.BandwidthPackageResourceQuota"
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
//  LIMITEXCEEDED_BANDWIDTHPACKAGERESOURCEQUOTA = "LimitExceeded.BandwidthPackageResourceQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_BANDWIDTHPACKAGEIDNOTSUPPORTED = "UnsupportedOperation.BandwidthPackageIdNotSupported"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTSUPPORTED = "UnsupportedOperation.InstanceStateNotSupported"
//  UNSUPPORTEDOPERATION_INVALIDRESOURCEINTERNETCHARGETYPE = "UnsupportedOperation.InvalidResourceInternetChargeType"
//  UNSUPPORTEDOPERATION_INVALIDRESOURCEPROTOCOL = "UnsupportedOperation.InvalidResourceProtocol"
func (c *Client) AddBandwidthPackageResourcesWithContext(ctx context.Context, request *AddBandwidthPackageResourcesRequest) (response *AddBandwidthPackageResourcesResponse, err error) {
    if request == nil {
        request = NewAddBandwidthPackageResourcesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "AddBandwidthPackageResources")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddBandwidthPackageResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddBandwidthPackageResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewAddTemplateMemberRequest() (request *AddTemplateMemberRequest) {
    request = &AddTemplateMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "AddTemplateMember")
    
    
    return
}

func NewAddTemplateMemberResponse() (response *AddTemplateMemberResponse) {
    response = &AddTemplateMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddTemplateMember
// This API is used to add a parameter template of the IP address, protocol port, IP address group, or protocol port group type.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) AddTemplateMember(request *AddTemplateMemberRequest) (response *AddTemplateMemberResponse, err error) {
    return c.AddTemplateMemberWithContext(context.Background(), request)
}

// AddTemplateMember
// This API is used to add a parameter template of the IP address, protocol port, IP address group, or protocol port group type.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) AddTemplateMemberWithContext(ctx context.Context, request *AddTemplateMemberRequest) (response *AddTemplateMemberResponse, err error) {
    if request == nil {
        request = NewAddTemplateMemberRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "AddTemplateMember")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddTemplateMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddTemplateMemberResponse()
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
// This API is used to change the public IP of a CVM or the EIP of the associated bandwidth package.
//
// error code that may be returned:
//  FAILEDOPERATION_TASKFAILED = "FailedOperation.TaskFailed"
//  INVALIDACCOUNT_NOTSUPPORTED = "InvalidAccount.NotSupported"
//  INVALIDADDRESSID_BLOCKED = "InvalidAddressId.Blocked"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE_ADDRESSATTACKED = "InvalidParameterValue.AddressAttacked"
//  INVALIDPARAMETERVALUE_ADDRESSIPNOTAVAILABLE = "InvalidParameterValue.AddressIpNotAvailable"
//  INVALIDPARAMETERVALUE_ADDRESSNOTFOUND = "InvalidParameterValue.AddressNotFound"
//  INVALIDPARAMETERVALUE_COMBINATION = "InvalidParameterValue.Combination"
//  INVALIDPARAMETERVALUE_INSTANCEHASNOWANIP = "InvalidParameterValue.InstanceHasNoWanIP"
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
// This API is used to change the public IP of a CVM or the EIP of the associated bandwidth package.
//
// error code that may be returned:
//  FAILEDOPERATION_TASKFAILED = "FailedOperation.TaskFailed"
//  INVALIDACCOUNT_NOTSUPPORTED = "InvalidAccount.NotSupported"
//  INVALIDADDRESSID_BLOCKED = "InvalidAddressId.Blocked"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE_ADDRESSATTACKED = "InvalidParameterValue.AddressAttacked"
//  INVALIDPARAMETERVALUE_ADDRESSIPNOTAVAILABLE = "InvalidParameterValue.AddressIpNotAvailable"
//  INVALIDPARAMETERVALUE_ADDRESSNOTFOUND = "InvalidParameterValue.AddressNotFound"
//  INVALIDPARAMETERVALUE_COMBINATION = "InvalidParameterValue.Combination"
//  INVALIDPARAMETERVALUE_INSTANCEHASNOWANIP = "InvalidParameterValue.InstanceHasNoWanIP"
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "AdjustPublicAddress")
    
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
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_RESOURCEIDMALFORMED = "InvalidParameterValue.ResourceIdMalformed"
//  INVALIDPARAMETERVALUE_RESOURCENOTSUPPORT = "InvalidParameterValue.ResourceNotSupport"
//  INVALIDPARAMETERVALUE_TAGNOTEXISTED = "InvalidParameterValue.TagNotExisted"
//  INVALIDPARAMETERVALUE_UNAVAILABLEZONE = "InvalidParameterValue.UnavailableZone"
//  LIMITEXCEEDED_BANDWIDTHPACKAGEQUOTA = "LimitExceeded.BandwidthPackageQuota"
//  LIMITEXCEEDED_BANDWIDTHPACKAGERESOURCEQUOTA = "LimitExceeded.BandwidthPackageResourceQuota"
//  LIMITEXCEEDED_MONTHLYADDRESSRECOVERYQUOTA = "LimitExceeded.MonthlyAddressRecoveryQuota"
//  LIMITEXCEEDED_TAGQUOTA = "LimitExceeded.TagQuota"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION_ANYCASTEIP = "UnauthorizedOperation.AnycastEip"
//  UNAUTHORIZEDOPERATION_INVALIDACCOUNT = "UnauthorizedOperation.InvalidAccount"
//  UNSUPPORTEDOPERATION_ACCOUNTNOTSUPPORTED = "UnsupportedOperation.AccountNotSupported"
//  UNSUPPORTEDOPERATION_ACTIONNOTFOUND = "UnsupportedOperation.ActionNotFound"
//  UNSUPPORTEDOPERATION_BANDWIDTHPACKAGEIDNOTSUPPORTED = "UnsupportedOperation.BandwidthPackageIdNotSupported"
//  UNSUPPORTEDOPERATION_DELIVERYFAILED = "UnsupportedOperation.DeliveryFailed"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTSUPPORTED = "UnsupportedOperation.InstanceStateNotSupported"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_NOTSUPPORTEDPURCHASECENTEREGRESSRESOURCE = "UnsupportedOperation.NotSupportedPurchaseCenterEgressResource"
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
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_RESOURCEIDMALFORMED = "InvalidParameterValue.ResourceIdMalformed"
//  INVALIDPARAMETERVALUE_RESOURCENOTSUPPORT = "InvalidParameterValue.ResourceNotSupport"
//  INVALIDPARAMETERVALUE_TAGNOTEXISTED = "InvalidParameterValue.TagNotExisted"
//  INVALIDPARAMETERVALUE_UNAVAILABLEZONE = "InvalidParameterValue.UnavailableZone"
//  LIMITEXCEEDED_BANDWIDTHPACKAGEQUOTA = "LimitExceeded.BandwidthPackageQuota"
//  LIMITEXCEEDED_BANDWIDTHPACKAGERESOURCEQUOTA = "LimitExceeded.BandwidthPackageResourceQuota"
//  LIMITEXCEEDED_MONTHLYADDRESSRECOVERYQUOTA = "LimitExceeded.MonthlyAddressRecoveryQuota"
//  LIMITEXCEEDED_TAGQUOTA = "LimitExceeded.TagQuota"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION_ANYCASTEIP = "UnauthorizedOperation.AnycastEip"
//  UNAUTHORIZEDOPERATION_INVALIDACCOUNT = "UnauthorizedOperation.InvalidAccount"
//  UNSUPPORTEDOPERATION_ACCOUNTNOTSUPPORTED = "UnsupportedOperation.AccountNotSupported"
//  UNSUPPORTEDOPERATION_ACTIONNOTFOUND = "UnsupportedOperation.ActionNotFound"
//  UNSUPPORTEDOPERATION_BANDWIDTHPACKAGEIDNOTSUPPORTED = "UnsupportedOperation.BandwidthPackageIdNotSupported"
//  UNSUPPORTEDOPERATION_DELIVERYFAILED = "UnsupportedOperation.DeliveryFailed"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTSUPPORTED = "UnsupportedOperation.InstanceStateNotSupported"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_NOTSUPPORTEDPURCHASECENTEREGRESSRESOURCE = "UnsupportedOperation.NotSupportedPurchaseCenterEgressResource"
//  UNSUPPORTEDOPERATION_OFFLINECHARGETYPE = "UnsupportedOperation.OfflineChargeType"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDREGION = "UnsupportedOperation.UnsupportedRegion"
func (c *Client) AllocateAddressesWithContext(ctx context.Context, request *AllocateAddressesRequest) (response *AllocateAddressesResponse, err error) {
    if request == nil {
        request = NewAllocateAddressesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "AllocateAddresses")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AllocateAddresses require credential")
    }

    request.SetContext(ctx)
    
    response = NewAllocateAddressesResponse()
    err = c.Send(request, response)
    return
}

func NewAllocateIPv6AddressesRequest() (request *AllocateIPv6AddressesRequest) {
    request = &AllocateIPv6AddressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "AllocateIPv6Addresses")
    
    
    return
}

func NewAllocateIPv6AddressesResponse() (response *AllocateIPv6AddressesResponse) {
    response = &AllocateIPv6AddressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AllocateIPv6Addresses
// This API is used to apply for one or more Elastic IPv6 (EIPv6) instances.
//
// 
//
// - EIPv6 is a fixed public IPv6 address that can be independently applied for and held in a Tencent Cloud region, providing a consistent product experience with Elastic IPv4.
//
// - You can quickly bind an EIPv6 instance to the private IPv6 address of a cloud resource, so as to quickly enable IPv6 public bandwidth for the cloud resource.
//
// - You can also bind an EIPv6 instance to other cloud resources as needed, so as to shield instance failures.
//
// error code that may be returned:
//  INVALIDPARAMETERCONFLICT = "InvalidParameterConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BANDWIDTHOUTOFRANGE = "InvalidParameterValue.BandwidthOutOfRange"
//  INVALIDPARAMETERVALUE_BANDWIDTHTOOSMALL = "InvalidParameterValue.BandwidthTooSmall"
//  INVALIDPARAMETERVALUE_COMBINATION = "InvalidParameterValue.Combination"
//  INVALIDPARAMETERVALUE_UNAVAILABLEZONE = "InvalidParameterValue.UnavailableZone"
//  LIMITEXCEEDED_BANDWIDTHPACKAGERESOURCEQUOTA = "LimitExceeded.BandwidthPackageResourceQuota"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_ACCOUNTNOTSUPPORTED = "UnsupportedOperation.AccountNotSupported"
func (c *Client) AllocateIPv6Addresses(request *AllocateIPv6AddressesRequest) (response *AllocateIPv6AddressesResponse, err error) {
    return c.AllocateIPv6AddressesWithContext(context.Background(), request)
}

// AllocateIPv6Addresses
// This API is used to apply for one or more Elastic IPv6 (EIPv6) instances.
//
// 
//
// - EIPv6 is a fixed public IPv6 address that can be independently applied for and held in a Tencent Cloud region, providing a consistent product experience with Elastic IPv4.
//
// - You can quickly bind an EIPv6 instance to the private IPv6 address of a cloud resource, so as to quickly enable IPv6 public bandwidth for the cloud resource.
//
// - You can also bind an EIPv6 instance to other cloud resources as needed, so as to shield instance failures.
//
// error code that may be returned:
//  INVALIDPARAMETERCONFLICT = "InvalidParameterConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BANDWIDTHOUTOFRANGE = "InvalidParameterValue.BandwidthOutOfRange"
//  INVALIDPARAMETERVALUE_BANDWIDTHTOOSMALL = "InvalidParameterValue.BandwidthTooSmall"
//  INVALIDPARAMETERVALUE_COMBINATION = "InvalidParameterValue.Combination"
//  INVALIDPARAMETERVALUE_UNAVAILABLEZONE = "InvalidParameterValue.UnavailableZone"
//  LIMITEXCEEDED_BANDWIDTHPACKAGERESOURCEQUOTA = "LimitExceeded.BandwidthPackageResourceQuota"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_ACCOUNTNOTSUPPORTED = "UnsupportedOperation.AccountNotSupported"
func (c *Client) AllocateIPv6AddressesWithContext(ctx context.Context, request *AllocateIPv6AddressesRequest) (response *AllocateIPv6AddressesResponse, err error) {
    if request == nil {
        request = NewAllocateIPv6AddressesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "AllocateIPv6Addresses")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AllocateIPv6Addresses require credential")
    }

    request.SetContext(ctx)
    
    response = NewAllocateIPv6AddressesResponse()
    err = c.Send(request, response)
    return
}

func NewAllocateIp6AddressesBandwidthRequest() (request *AllocateIp6AddressesBandwidthRequest) {
    request = &AllocateIp6AddressesBandwidthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "AllocateIp6AddressesBandwidth")
    
    
    return
}

func NewAllocateIp6AddressesBandwidthResponse() (response *AllocateIp6AddressesBandwidthResponse) {
    response = &AllocateIp6AddressesBandwidthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AllocateIp6AddressesBandwidth
// This API is used to allocate IPv6 public network bandwidth for classic elastic public IPv6 addresses.
//
// 
//
// - Classic elastic public IPv6 addresses only have the private network communication capability by default. They can have the IPv6 public network communication capability and be displayed in the list of Classic Elastic Public IPv6 only after IPv6 public network bandwidth is allocated in the console or by calling this API. 
//
// - You can allocate public network bandwidth for one or multiple Classic elastic public IPv6 addresses each time.
//
// error code that may be returned:
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_TASKFAILED = "FailedOperation.TaskFailed"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDACCOUNT_NOTSUPPORTED = "InvalidAccount.NotSupported"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE_ADDRESSIPNOTFOUND = "InvalidParameterValue.AddressIpNotFound"
//  INVALIDPARAMETERVALUE_ADDRESSIPNOTINVPC = "InvalidParameterValue.AddressIpNotInVpc"
//  INVALIDPARAMETERVALUE_ADDRESSPUBLISHED = "InvalidParameterValue.AddressPublished"
//  INVALIDPARAMETERVALUE_BANDWIDTHOUTOFRANGE = "InvalidParameterValue.BandwidthOutOfRange"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGENOTFOUND = "InvalidParameterValue.BandwidthPackageNotFound"
//  INVALIDPARAMETERVALUE_INVALIDIPV6 = "InvalidParameterValue.InvalidIpv6"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  UNSUPPORTEDOPERATION_BANDWIDTHPACKAGEIDNOTSUPPORTED = "UnsupportedOperation.BandwidthPackageIdNotSupported"
//  UNSUPPORTEDOPERATION_INVALIDADDRESSINTERNETCHARGETYPE = "UnsupportedOperation.InvalidAddressInternetChargeType"
func (c *Client) AllocateIp6AddressesBandwidth(request *AllocateIp6AddressesBandwidthRequest) (response *AllocateIp6AddressesBandwidthResponse, err error) {
    return c.AllocateIp6AddressesBandwidthWithContext(context.Background(), request)
}

// AllocateIp6AddressesBandwidth
// This API is used to allocate IPv6 public network bandwidth for classic elastic public IPv6 addresses.
//
// 
//
// - Classic elastic public IPv6 addresses only have the private network communication capability by default. They can have the IPv6 public network communication capability and be displayed in the list of Classic Elastic Public IPv6 only after IPv6 public network bandwidth is allocated in the console or by calling this API. 
//
// - You can allocate public network bandwidth for one or multiple Classic elastic public IPv6 addresses each time.
//
// error code that may be returned:
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_TASKFAILED = "FailedOperation.TaskFailed"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDACCOUNT_NOTSUPPORTED = "InvalidAccount.NotSupported"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE_ADDRESSIPNOTFOUND = "InvalidParameterValue.AddressIpNotFound"
//  INVALIDPARAMETERVALUE_ADDRESSIPNOTINVPC = "InvalidParameterValue.AddressIpNotInVpc"
//  INVALIDPARAMETERVALUE_ADDRESSPUBLISHED = "InvalidParameterValue.AddressPublished"
//  INVALIDPARAMETERVALUE_BANDWIDTHOUTOFRANGE = "InvalidParameterValue.BandwidthOutOfRange"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGENOTFOUND = "InvalidParameterValue.BandwidthPackageNotFound"
//  INVALIDPARAMETERVALUE_INVALIDIPV6 = "InvalidParameterValue.InvalidIpv6"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  UNSUPPORTEDOPERATION_BANDWIDTHPACKAGEIDNOTSUPPORTED = "UnsupportedOperation.BandwidthPackageIdNotSupported"
//  UNSUPPORTEDOPERATION_INVALIDADDRESSINTERNETCHARGETYPE = "UnsupportedOperation.InvalidAddressInternetChargeType"
func (c *Client) AllocateIp6AddressesBandwidthWithContext(ctx context.Context, request *AllocateIp6AddressesBandwidthRequest) (response *AllocateIp6AddressesBandwidthResponse, err error) {
    if request == nil {
        request = NewAllocateIp6AddressesBandwidthRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "AllocateIp6AddressesBandwidth")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AllocateIp6AddressesBandwidth require credential")
    }

    request.SetContext(ctx)
    
    response = NewAllocateIp6AddressesBandwidthResponse()
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_RESERVED = "InvalidParameterValue.Reserved"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_ACTIONLIMITED = "LimitExceeded.ActionLimited"
//  LIMITEXCEEDED_ADDRESS = "LimitExceeded.Address"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_RESERVED = "InvalidParameterValue.Reserved"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_ACTIONLIMITED = "LimitExceeded.ActionLimited"
//  LIMITEXCEEDED_ADDRESS = "LimitExceeded.Address"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "AssignIpv6Addresses")
    
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
//  UNSUPPORTEDOPERATION_IPV6CIDRNOTDEPLOYED = "UnsupportedOperation.IPV6CidrNotDeployed"
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
//  UNSUPPORTEDOPERATION_IPV6CIDRNOTDEPLOYED = "UnsupportedOperation.IPV6CidrNotDeployed"
func (c *Client) AssignIpv6CidrBlockWithContext(ctx context.Context, request *AssignIpv6CidrBlockRequest) (response *AssignIpv6CidrBlockResponse, err error) {
    if request == nil {
        request = NewAssignIpv6CidrBlockRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "AssignIpv6CidrBlock")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "AssignIpv6SubnetCidrBlock")
    
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
//  LIMITEXCEEDED_ACTIONLIMITED = "LimitExceeded.ActionLimited"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_IDEMPOTENTPARAMETERMISMATCH = "UnsupportedOperation.IdempotentParameterMismatch"
//  UNSUPPORTEDOPERATION_IDEMPOTENTPROCESSING = "UnsupportedOperation.IdempotentProcessing"
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
//  LIMITEXCEEDED_ACTIONLIMITED = "LimitExceeded.ActionLimited"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_IDEMPOTENTPARAMETERMISMATCH = "UnsupportedOperation.IdempotentParameterMismatch"
//  UNSUPPORTEDOPERATION_IDEMPOTENTPROCESSING = "UnsupportedOperation.IdempotentProcessing"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
//  UNSUPPORTEDOPERATION_RESOURCEMISMATCH = "UnsupportedOperation.ResourceMismatch"
func (c *Client) AssignPrivateIpAddressesWithContext(ctx context.Context, request *AssignPrivateIpAddressesRequest) (response *AssignPrivateIpAddressesResponse, err error) {
    if request == nil {
        request = NewAssignPrivateIpAddressesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "AssignPrivateIpAddresses")
    
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
//  FAILEDOPERATION_TASKFAILED = "FailedOperation.TaskFailed"
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
//  INVALIDPARAMETERVALUE_INSTANCEHASNOWANIP = "InvalidParameterValue.InstanceHasNoWanIP"
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
//  OPERATIONDENIED_ADDRESSINARREARS = "OperationDenied.AddressInArrears"
//  OPERATIONDENIED_MUTEXTASKRUNNING = "OperationDenied.MutexTaskRunning"
//  UNSUPPORTEDOPERATION_ACTIONNOTFOUND = "UnsupportedOperation.ActionNotFound"
//  UNSUPPORTEDOPERATION_ADDRESSIPNOTSUPPORTINSTANCE = "UnsupportedOperation.AddressIpNotSupportInstance"
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
//  FAILEDOPERATION_TASKFAILED = "FailedOperation.TaskFailed"
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
//  INVALIDPARAMETERVALUE_INSTANCEHASNOWANIP = "InvalidParameterValue.InstanceHasNoWanIP"
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
//  OPERATIONDENIED_ADDRESSINARREARS = "OperationDenied.AddressInArrears"
//  OPERATIONDENIED_MUTEXTASKRUNNING = "OperationDenied.MutexTaskRunning"
//  UNSUPPORTEDOPERATION_ACTIONNOTFOUND = "UnsupportedOperation.ActionNotFound"
//  UNSUPPORTEDOPERATION_ADDRESSIPNOTSUPPORTINSTANCE = "UnsupportedOperation.AddressIpNotSupportInstance"
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "AssociateAddress")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "AssociateDirectConnectGatewayNatGateway")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssociateDirectConnectGatewayNatGateway require credential")
    }

    request.SetContext(ctx)
    
    response = NewAssociateDirectConnectGatewayNatGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewAssociateIPv6AddressRequest() (request *AssociateIPv6AddressRequest) {
    request = &AssociateIPv6AddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "AssociateIPv6Address")
    
    
    return
}

func NewAssociateIPv6AddressResponse() (response *AssociateIPv6AddressResponse) {
    response = &AssociateIPv6AddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AssociateIPv6Address
// This API is used to bind an EIPv6 instance to the private IPv6 address configured on the CVM or ENI.
//
// 
//
// - Binding an EIPv6 to the CVM essentially indicates binding the EIPv6 to the private IPv6 address configured on the ENI of the CVM.
//
// - Before binding an EIPv6 to the private IPv6 of a specified ENI, ensure that the private IPv6 address is unbound before the binding operation is performed.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCONFLICT = "InvalidParameterConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ADDRESSIDMALFORMED = "InvalidParameterValue.AddressIdMalformed"
//  INVALIDPARAMETERVALUE_ADDRESSIPNOTFOUND = "InvalidParameterValue.AddressIpNotFound"
//  INVALIDPARAMETERVALUE_ADDRESSNOTAPPLICABLE = "InvalidParameterValue.AddressNotApplicable"
//  INVALIDPARAMETERVALUE_MISSINGASSOCIATEENTITY = "InvalidParameterValue.MissingAssociateEntity"
//  INVALIDPARAMETERVALUE_NETWORKINTERFACEINSTANCENOTSUPPORT = "InvalidParameterValue.NetworkInterfaceInstanceNotSupport"
//  INVALIDPARAMETERVALUE_NETWORKINTERFACENOTFOUND = "InvalidParameterValue.NetworkInterfaceNotFound"
//  INVALIDPRIVATEIPADDRESS_ALREADYBINDEIP = "InvalidPrivateIpAddress.AlreadyBindEip"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ADDRESSINARREARS = "OperationDenied.AddressInArrears"
//  UNSUPPORTEDOPERATION_ADDRESSSTATUSNOTPERMIT = "UnsupportedOperation.AddressStatusNotPermit"
func (c *Client) AssociateIPv6Address(request *AssociateIPv6AddressRequest) (response *AssociateIPv6AddressResponse, err error) {
    return c.AssociateIPv6AddressWithContext(context.Background(), request)
}

// AssociateIPv6Address
// This API is used to bind an EIPv6 instance to the private IPv6 address configured on the CVM or ENI.
//
// 
//
// - Binding an EIPv6 to the CVM essentially indicates binding the EIPv6 to the private IPv6 address configured on the ENI of the CVM.
//
// - Before binding an EIPv6 to the private IPv6 of a specified ENI, ensure that the private IPv6 address is unbound before the binding operation is performed.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCONFLICT = "InvalidParameterConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ADDRESSIDMALFORMED = "InvalidParameterValue.AddressIdMalformed"
//  INVALIDPARAMETERVALUE_ADDRESSIPNOTFOUND = "InvalidParameterValue.AddressIpNotFound"
//  INVALIDPARAMETERVALUE_ADDRESSNOTAPPLICABLE = "InvalidParameterValue.AddressNotApplicable"
//  INVALIDPARAMETERVALUE_MISSINGASSOCIATEENTITY = "InvalidParameterValue.MissingAssociateEntity"
//  INVALIDPARAMETERVALUE_NETWORKINTERFACEINSTANCENOTSUPPORT = "InvalidParameterValue.NetworkInterfaceInstanceNotSupport"
//  INVALIDPARAMETERVALUE_NETWORKINTERFACENOTFOUND = "InvalidParameterValue.NetworkInterfaceNotFound"
//  INVALIDPRIVATEIPADDRESS_ALREADYBINDEIP = "InvalidPrivateIpAddress.AlreadyBindEip"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ADDRESSINARREARS = "OperationDenied.AddressInArrears"
//  UNSUPPORTEDOPERATION_ADDRESSSTATUSNOTPERMIT = "UnsupportedOperation.AddressStatusNotPermit"
func (c *Client) AssociateIPv6AddressWithContext(ctx context.Context, request *AssociateIPv6AddressRequest) (response *AssociateIPv6AddressResponse, err error) {
    if request == nil {
        request = NewAssociateIPv6AddressRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "AssociateIPv6Address")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssociateIPv6Address require credential")
    }

    request.SetContext(ctx)
    
    response = NewAssociateIPv6AddressResponse()
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETERVALUE_EIPBRANDWIDTHOUTINVALID = "InvalidParameterValue.EIPBrandWidthOutInvalid"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.AddressQuotaLimitExceeded"
//  LIMITEXCEEDED_DAILYALLOCATEADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.DailyAllocateAddressQuotaLimitExceeded"
//  LIMITEXCEEDED_PUBLICIPADDRESSPERNATGATEWAYLIMITEXCEEDED = "LimitExceeded.PublicIpAddressPerNatGatewayLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_ADDRESS = "ResourceInUse.Address"
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETERVALUE_EIPBRANDWIDTHOUTINVALID = "InvalidParameterValue.EIPBrandWidthOutInvalid"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.AddressQuotaLimitExceeded"
//  LIMITEXCEEDED_DAILYALLOCATEADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.DailyAllocateAddressQuotaLimitExceeded"
//  LIMITEXCEEDED_PUBLICIPADDRESSPERNATGATEWAYLIMITEXCEEDED = "LimitExceeded.PublicIpAddressPerNatGatewayLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_ADDRESS = "ResourceInUse.Address"
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "AssociateNatGatewayAddress")
    
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
func (c *Client) AssociateNetworkAclSubnetsWithContext(ctx context.Context, request *AssociateNetworkAclSubnetsRequest) (response *AssociateNetworkAclSubnetsResponse, err error) {
    if request == nil {
        request = NewAssociateNetworkAclSubnetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "AssociateNetworkAclSubnets")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "AssociateNetworkInterfaceSecurityGroups")
    
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
// This API is used to add a network instance to a CCN instance. Network instances include VPCs and Direct Connect gateways. <br />
//
// The number of network instances that each CCN can be associated with is limited. For more information, see the product documentation. If you need to associate more instances, please submit a ticket.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CCNATTACHBMVPCLIMITEXCEEDED = "InvalidParameterValue.CcnAttachBmvpcLimitExceeded"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CURRENTINSTANCEATTACHEDCCNINSTANCES = "LimitExceeded.CurrentInstanceAttachedCcnInstances"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_APPIDNOTFOUND = "UnsupportedOperation.AppIdNotFound"
//  UNSUPPORTEDOPERATION_CCNATTACHED = "UnsupportedOperation.CcnAttached"
//  UNSUPPORTEDOPERATION_CCNCROSSACCOUNT = "UnsupportedOperation.CcnCrossAccount"
//  UNSUPPORTEDOPERATION_CCNORDINARYACCOUNTREFUSEATTACH = "UnsupportedOperation.CcnOrdinaryAccountRefuseAttach"
//  UNSUPPORTEDOPERATION_CCNROUTETABLENOTEXIST = "UnsupportedOperation.CcnRouteTableNotExist"
//  UNSUPPORTEDOPERATION_INSTANCEANDRTBNOTMATCH = "UnsupportedOperation.InstanceAndRtbNotMatch"
//  UNSUPPORTEDOPERATION_INSTANCECDCIDNOTMATCHCCNCDCID = "UnsupportedOperation.InstanceCdcIdNotMatchCcnCdcId"
//  UNSUPPORTEDOPERATION_INSTANCEORDINARYACCOUNTREFUSEATTACH = "UnsupportedOperation.InstanceOrdinaryAccountRefuseAttach"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_ISNOTFINANCEACCOUNT = "UnsupportedOperation.IsNotFinanceAccount"
//  UNSUPPORTEDOPERATION_MULTIPLEVPCNOTSUPPORTATTACHACCOUNTHASIPV6 = "UnsupportedOperation.MultipleVpcNotSupportAttachAccountHasIpv6"
//  UNSUPPORTEDOPERATION_NOTSUPPORTATTACHEDGEANDCROSSBORDERINSTANCE = "UnsupportedOperation.NotSupportAttachEdgeAndCrossBorderInstance"
//  UNSUPPORTEDOPERATION_PURCHASELIMIT = "UnsupportedOperation.PurchaseLimit"
//  UNSUPPORTEDOPERATION_UINNOTFOUND = "UnsupportedOperation.UinNotFound"
//  UNSUPPORTEDOPERATION_UNABLECROSSBORDER = "UnsupportedOperation.UnableCrossBorder"
//  UNSUPPORTEDOPERATION_UNABLECROSSFINANCE = "UnsupportedOperation.UnableCrossFinance"
func (c *Client) AttachCcnInstances(request *AttachCcnInstancesRequest) (response *AttachCcnInstancesResponse, err error) {
    return c.AttachCcnInstancesWithContext(context.Background(), request)
}

// AttachCcnInstances
// This API is used to add a network instance to a CCN instance. Network instances include VPCs and Direct Connect gateways. <br />
//
// The number of network instances that each CCN can be associated with is limited. For more information, see the product documentation. If you need to associate more instances, please submit a ticket.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CCNATTACHBMVPCLIMITEXCEEDED = "InvalidParameterValue.CcnAttachBmvpcLimitExceeded"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CURRENTINSTANCEATTACHEDCCNINSTANCES = "LimitExceeded.CurrentInstanceAttachedCcnInstances"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_APPIDNOTFOUND = "UnsupportedOperation.AppIdNotFound"
//  UNSUPPORTEDOPERATION_CCNATTACHED = "UnsupportedOperation.CcnAttached"
//  UNSUPPORTEDOPERATION_CCNCROSSACCOUNT = "UnsupportedOperation.CcnCrossAccount"
//  UNSUPPORTEDOPERATION_CCNORDINARYACCOUNTREFUSEATTACH = "UnsupportedOperation.CcnOrdinaryAccountRefuseAttach"
//  UNSUPPORTEDOPERATION_CCNROUTETABLENOTEXIST = "UnsupportedOperation.CcnRouteTableNotExist"
//  UNSUPPORTEDOPERATION_INSTANCEANDRTBNOTMATCH = "UnsupportedOperation.InstanceAndRtbNotMatch"
//  UNSUPPORTEDOPERATION_INSTANCECDCIDNOTMATCHCCNCDCID = "UnsupportedOperation.InstanceCdcIdNotMatchCcnCdcId"
//  UNSUPPORTEDOPERATION_INSTANCEORDINARYACCOUNTREFUSEATTACH = "UnsupportedOperation.InstanceOrdinaryAccountRefuseAttach"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_ISNOTFINANCEACCOUNT = "UnsupportedOperation.IsNotFinanceAccount"
//  UNSUPPORTEDOPERATION_MULTIPLEVPCNOTSUPPORTATTACHACCOUNTHASIPV6 = "UnsupportedOperation.MultipleVpcNotSupportAttachAccountHasIpv6"
//  UNSUPPORTEDOPERATION_NOTSUPPORTATTACHEDGEANDCROSSBORDERINSTANCE = "UnsupportedOperation.NotSupportAttachEdgeAndCrossBorderInstance"
//  UNSUPPORTEDOPERATION_PURCHASELIMIT = "UnsupportedOperation.PurchaseLimit"
//  UNSUPPORTEDOPERATION_UINNOTFOUND = "UnsupportedOperation.UinNotFound"
//  UNSUPPORTEDOPERATION_UNABLECROSSBORDER = "UnsupportedOperation.UnableCrossBorder"
//  UNSUPPORTEDOPERATION_UNABLECROSSFINANCE = "UnsupportedOperation.UnableCrossFinance"
func (c *Client) AttachCcnInstancesWithContext(ctx context.Context, request *AttachCcnInstancesRequest) (response *AttachCcnInstancesResponse, err error) {
    if request == nil {
        request = NewAttachCcnInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "AttachCcnInstances")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "AttachClassicLinkVpc")
    
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_ACTIONLIMITED = "LimitExceeded.ActionLimited"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ATTACHMENTALREADYEXISTS = "UnsupportedOperation.AttachmentAlreadyExists"
//  UNSUPPORTEDOPERATION_IDEMPOTENTPARAMETERMISMATCH = "UnsupportedOperation.IdempotentParameterMismatch"
//  UNSUPPORTEDOPERATION_IDEMPOTENTPROCESSING = "UnsupportedOperation.IdempotentProcessing"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_RESOURCEISINVALIDSTATE = "UnsupportedOperation.ResourceIsInvalidState"
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_ACTIONLIMITED = "LimitExceeded.ActionLimited"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ATTACHMENTALREADYEXISTS = "UnsupportedOperation.AttachmentAlreadyExists"
//  UNSUPPORTEDOPERATION_IDEMPOTENTPARAMETERMISMATCH = "UnsupportedOperation.IdempotentParameterMismatch"
//  UNSUPPORTEDOPERATION_IDEMPOTENTPROCESSING = "UnsupportedOperation.IdempotentProcessing"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_RESOURCEISINVALIDSTATE = "UnsupportedOperation.ResourceIsInvalidState"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDINSTANCEFAMILY = "UnsupportedOperation.UnsupportedInstanceFamily"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
//  UNSUPPORTEDOPERATION_ZONEMISMATCH = "UnsupportedOperation.ZoneMismatch"
func (c *Client) AttachNetworkInterfaceWithContext(ctx context.Context, request *AttachNetworkInterfaceRequest) (response *AttachNetworkInterfaceResponse, err error) {
    if request == nil {
        request = NewAttachNetworkInterfaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "AttachNetworkInterface")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AttachNetworkInterface require credential")
    }

    request.SetContext(ctx)
    
    response = NewAttachNetworkInterfaceResponse()
    err = c.Send(request, response)
    return
}

func NewAttachSnapshotInstancesRequest() (request *AttachSnapshotInstancesRequest) {
    request = &AttachSnapshotInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "AttachSnapshotInstances")
    
    
    return
}

func NewAttachSnapshotInstancesResponse() (response *AttachSnapshotInstancesResponse) {
    response = &AttachSnapshotInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AttachSnapshotInstances
// This API is used to associate a snapshot policy with specified instances.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_DUPLICATEPARA = "InvalidParameterValue.DuplicatePara"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_ATTACHEDSNAPSHOTPOLICYEXCEEDED = "LimitExceeded.AttachedSnapshotPolicyExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_SNAPSHOTATTACHED = "UnsupportedOperation.SnapshotAttached"
//  UNSUPPORTEDOPERATION_SNAPSHOTINSTANCEREGIONDIFF = "UnsupportedOperation.SnapshotInstanceRegionDiff"
func (c *Client) AttachSnapshotInstances(request *AttachSnapshotInstancesRequest) (response *AttachSnapshotInstancesResponse, err error) {
    return c.AttachSnapshotInstancesWithContext(context.Background(), request)
}

// AttachSnapshotInstances
// This API is used to associate a snapshot policy with specified instances.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_DUPLICATEPARA = "InvalidParameterValue.DuplicatePara"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_ATTACHEDSNAPSHOTPOLICYEXCEEDED = "LimitExceeded.AttachedSnapshotPolicyExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_SNAPSHOTATTACHED = "UnsupportedOperation.SnapshotAttached"
//  UNSUPPORTEDOPERATION_SNAPSHOTINSTANCEREGIONDIFF = "UnsupportedOperation.SnapshotInstanceRegionDiff"
func (c *Client) AttachSnapshotInstancesWithContext(ctx context.Context, request *AttachSnapshotInstancesRequest) (response *AttachSnapshotInstancesResponse, err error) {
    if request == nil {
        request = NewAttachSnapshotInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "AttachSnapshotInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AttachSnapshotInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewAttachSnapshotInstancesResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "AuditCrossBorderCompliance")
    
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
// This API is used to check whether the secondary CIDR block conflicts with existing routes, peering connections (peer VPC CIDR blocks), and other resources. 
//
// * Check whether the secondary CIDR block overlaps with the primary/secondary CIDR block of the current VPC.
//
// * Check whether the secondary CIDR block overlaps with the routing destination of the current VPC.
//
// * If the current VPC is used in a peering connection, check whether the secondary CIDR block overlaps with the primary/secondary CIDR block of the peer VPC.
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
// This API is used to check whether the secondary CIDR block conflicts with existing routes, peering connections (peer VPC CIDR blocks), and other resources. 
//
// * Check whether the secondary CIDR block overlaps with the primary/secondary CIDR block of the current VPC.
//
// * Check whether the secondary CIDR block overlaps with the routing destination of the current VPC.
//
// * If the current VPC is used in a peering connection, check whether the secondary CIDR block overlaps with the primary/secondary CIDR block of the peer VPC.
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CheckAssistantCidr")
    
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CONFLICTWITHDOCKERROUTE = "UnsupportedOperation.ConflictWithDockerRoute"
//  UNSUPPORTEDOPERATION_ECMPWITHUSERROUTE = "UnsupportedOperation.EcmpWithUserRoute"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
func (c *Client) CheckNetDetectStateWithContext(ctx context.Context, request *CheckNetDetectStateRequest) (response *CheckNetDetectStateResponse, err error) {
    if request == nil {
        request = NewCheckNetDetectStateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CheckNetDetectState")
    
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
//  UNSUPPORTEDOPERATION_REMOTEREGIONSGHASREFERENCEDSG = "UnsupportedOperation.RemoteRegionSgHasReferencedSg"
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
//  UNSUPPORTEDOPERATION_REMOTEREGIONSGHASREFERENCEDSG = "UnsupportedOperation.RemoteRegionSgHasReferencedSg"
func (c *Client) CloneSecurityGroupWithContext(ctx context.Context, request *CloneSecurityGroupRequest) (response *CloneSecurityGroupResponse, err error) {
    if request == nil {
        request = NewCloneSecurityGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CloneSecurityGroup")
    
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
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
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
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateAddressTemplateWithContext(ctx context.Context, request *CreateAddressTemplateRequest) (response *CreateAddressTemplateResponse, err error) {
    if request == nil {
        request = NewCreateAddressTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateAddressTemplate")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateAddressTemplateGroup")
    
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
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETER_FILTERNOTDICT = "InvalidParameter.FilterNotDict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_RESERVED = "InvalidParameterValue.Reserved"
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_IDEMPOTENTPARAMETERMISMATCH = "UnsupportedOperation.IdempotentParameterMismatch"
//  UNSUPPORTEDOPERATION_IDEMPOTENTPROCESSING = "UnsupportedOperation.IdempotentProcessing"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_RESOURCEMISMATCH = "UnsupportedOperation.ResourceMismatch"
//  UNSUPPORTEDOPERATION_TAGALLOCATE = "UnsupportedOperation.TagAllocate"
//  UNSUPPORTEDOPERATION_TAGFREE = "UnsupportedOperation.TagFree"
//  UNSUPPORTEDOPERATION_TAGNOTPERMIT = "UnsupportedOperation.TagNotPermit"
//  UNSUPPORTEDOPERATION_TAGSYSTEMRESERVEDTAGKEY = "UnsupportedOperation.TagSystemReservedTagKey"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDINSTANCEFAMILY = "UnsupportedOperation.UnsupportedInstanceFamily"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
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
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETER_FILTERNOTDICT = "InvalidParameter.FilterNotDict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_RESERVED = "InvalidParameterValue.Reserved"
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_IDEMPOTENTPARAMETERMISMATCH = "UnsupportedOperation.IdempotentParameterMismatch"
//  UNSUPPORTEDOPERATION_IDEMPOTENTPROCESSING = "UnsupportedOperation.IdempotentProcessing"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_RESOURCEMISMATCH = "UnsupportedOperation.ResourceMismatch"
//  UNSUPPORTEDOPERATION_TAGALLOCATE = "UnsupportedOperation.TagAllocate"
//  UNSUPPORTEDOPERATION_TAGFREE = "UnsupportedOperation.TagFree"
//  UNSUPPORTEDOPERATION_TAGNOTPERMIT = "UnsupportedOperation.TagNotPermit"
//  UNSUPPORTEDOPERATION_TAGSYSTEMRESERVEDTAGKEY = "UnsupportedOperation.TagSystemReservedTagKey"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDINSTANCEFAMILY = "UnsupportedOperation.UnsupportedInstanceFamily"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
func (c *Client) CreateAndAttachNetworkInterfaceWithContext(ctx context.Context, request *CreateAndAttachNetworkInterfaceRequest) (response *CreateAndAttachNetworkInterfaceResponse, err error) {
    if request == nil {
        request = NewCreateAndAttachNetworkInterfaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateAndAttachNetworkInterface")
    
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
// This API is used to batch create secondary CIDR blocks.
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
// This API is used to batch create secondary CIDR blocks.
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateAssistantCidr")
    
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
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDACCOUNT_NOTSUPPORTED = "InvalidAccount.NotSupported"
//  INVALIDPARAMETERVALUE_COMBINATION = "InvalidParameterValue.Combination"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TAGNOTEXISTED = "InvalidParameterValue.TagNotExisted"
//  LIMITEXCEEDED_BANDWIDTHPACKAGEQUOTA = "LimitExceeded.BandwidthPackageQuota"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ACCOUNTNOTSUPPORTED = "UnsupportedOperation.AccountNotSupported"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTSUPPORTED = "UnsupportedOperation.InstanceStateNotSupported"
//  UNSUPPORTEDOPERATION_INVALIDRESOURCEINTERNETCHARGETYPE = "UnsupportedOperation.InvalidResourceInternetChargeType"
//  UNSUPPORTEDOPERATION_NOTSUPPORTEDPURCHASECENTEREGRESSRESOURCE = "UnsupportedOperation.NotSupportedPurchaseCenterEgressResource"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDREGION = "UnsupportedOperation.UnsupportedRegion"
func (c *Client) CreateBandwidthPackage(request *CreateBandwidthPackageRequest) (response *CreateBandwidthPackageResponse, err error) {
    return c.CreateBandwidthPackageWithContext(context.Background(), request)
}

// CreateBandwidthPackage
// This API is used to create a [device bandwidth package](https://intl.cloud.tencent.com/document/product/684/15245?from_cn_redirect=1#bwptype) or an [IP bandwidth package](https://intl.cloud.tencent.com/document/product/684/15245?from_cn_redirect=1#bwptype).
//
// error code that may be returned:
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDACCOUNT_NOTSUPPORTED = "InvalidAccount.NotSupported"
//  INVALIDPARAMETERVALUE_COMBINATION = "InvalidParameterValue.Combination"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TAGNOTEXISTED = "InvalidParameterValue.TagNotExisted"
//  LIMITEXCEEDED_BANDWIDTHPACKAGEQUOTA = "LimitExceeded.BandwidthPackageQuota"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ACCOUNTNOTSUPPORTED = "UnsupportedOperation.AccountNotSupported"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTSUPPORTED = "UnsupportedOperation.InstanceStateNotSupported"
//  UNSUPPORTEDOPERATION_INVALIDRESOURCEINTERNETCHARGETYPE = "UnsupportedOperation.InvalidResourceInternetChargeType"
//  UNSUPPORTEDOPERATION_NOTSUPPORTEDPURCHASECENTEREGRESSRESOURCE = "UnsupportedOperation.NotSupportedPurchaseCenterEgressResource"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDREGION = "UnsupportedOperation.UnsupportedRegion"
func (c *Client) CreateBandwidthPackageWithContext(ctx context.Context, request *CreateBandwidthPackageRequest) (response *CreateBandwidthPackageResponse, err error) {
    if request == nil {
        request = NewCreateBandwidthPackageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateBandwidthPackage")
    
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
// This API is used to create a CCN instance.
//
// * You can add tags to a CCN instance upon the creation. The tags are added successfully if they are listed in the response.
//
// * There is a quota of CCN instances for each account. For more information, see product documentation. To increase the quota, please submit a ticket.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_PARAMETERMISMATCH = "InvalidParameterValue.ParameterMismatch"
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
// This API is used to create a CCN instance.
//
// * You can add tags to a CCN instance upon the creation. The tags are added successfully if they are listed in the response.
//
// * There is a quota of CCN instances for each account. For more information, see product documentation. To increase the quota, please submit a ticket.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_PARAMETERMISMATCH = "InvalidParameterValue.ParameterMismatch"
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateCcn")
    
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
//  UNSUPPORTEDOPERATION_RECORDEXISTS = "UnsupportedOperation.RecordExists"
//  UNSUPPORTEDOPERATION_TAGALLOCATE = "UnsupportedOperation.TagAllocate"
//  UNSUPPORTEDOPERATION_TAGFREE = "UnsupportedOperation.TagFree"
//  UNSUPPORTEDOPERATION_TAGNOTPERMIT = "UnsupportedOperation.TagNotPermit"
//  UNSUPPORTEDOPERATION_TAGSYSTEMRESERVEDTAGKEY = "UnsupportedOperation.TagSystemReservedTagKey"
//  UNSUPPORTEDOPERATION_VPNUNSUPPORTEDBGP = "UnsupportedOperation.VpnUnsupportedBgp"
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
//  UNSUPPORTEDOPERATION_RECORDEXISTS = "UnsupportedOperation.RecordExists"
//  UNSUPPORTEDOPERATION_TAGALLOCATE = "UnsupportedOperation.TagAllocate"
//  UNSUPPORTEDOPERATION_TAGFREE = "UnsupportedOperation.TagFree"
//  UNSUPPORTEDOPERATION_TAGNOTPERMIT = "UnsupportedOperation.TagNotPermit"
//  UNSUPPORTEDOPERATION_TAGSYSTEMRESERVEDTAGKEY = "UnsupportedOperation.TagSystemReservedTagKey"
//  UNSUPPORTEDOPERATION_VPNUNSUPPORTEDBGP = "UnsupportedOperation.VpnUnsupportedBgp"
//  VPCLIMITEXCEEDED = "VpcLimitExceeded"
func (c *Client) CreateCustomerGatewayWithContext(ctx context.Context, request *CreateCustomerGatewayRequest) (response *CreateCustomerGatewayResponse, err error) {
    if request == nil {
        request = NewCreateCustomerGatewayRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateCustomerGateway")
    
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
// This API is used to create a VPC with default settings.
//
// 
//
// To create a VPC with custom settings, such as VPC name, IP range, subnet IP range, and subnet availability zone, use `CreateVpc` instead.
//
// 
//
// This API may not create a default VPC. It depends on the network attributes (`DescribeAccountAttributes`) of your account.
//
// * If both basic network and VPC are supported, the returned `VpcId` is 0.
//
// * If only VPC is supported, the default VPC information is returned.
//
// 
//
// You can also use the `Force` parameter to forcibly return a default VPC.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_SUBNETOVERLAP = "InvalidParameterValue.SubnetOverlap"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_DEFAULTVPCLIMITEXCEEDED = "LimitExceeded.DefaultVpcLimitExceeded"
//  RESOURCEINSUFFICIENT_CIDRBLOCK = "ResourceInsufficient.CidrBlock"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_RECORDEXISTS = "UnsupportedOperation.RecordExists"
func (c *Client) CreateDefaultVpc(request *CreateDefaultVpcRequest) (response *CreateDefaultVpcResponse, err error) {
    return c.CreateDefaultVpcWithContext(context.Background(), request)
}

// CreateDefaultVpc
// This API is used to create a VPC with default settings.
//
// 
//
// To create a VPC with custom settings, such as VPC name, IP range, subnet IP range, and subnet availability zone, use `CreateVpc` instead.
//
// 
//
// This API may not create a default VPC. It depends on the network attributes (`DescribeAccountAttributes`) of your account.
//
// * If both basic network and VPC are supported, the returned `VpcId` is 0.
//
// * If only VPC is supported, the default VPC information is returned.
//
// 
//
// You can also use the `Force` parameter to forcibly return a default VPC.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_SUBNETOVERLAP = "InvalidParameterValue.SubnetOverlap"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_DEFAULTVPCLIMITEXCEEDED = "LimitExceeded.DefaultVpcLimitExceeded"
//  RESOURCEINSUFFICIENT_CIDRBLOCK = "ResourceInsufficient.CidrBlock"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_RECORDEXISTS = "UnsupportedOperation.RecordExists"
func (c *Client) CreateDefaultVpcWithContext(ctx context.Context, request *CreateDefaultVpcRequest) (response *CreateDefaultVpcResponse, err error) {
    if request == nil {
        request = NewCreateDefaultVpcRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateDefaultVpc")
    
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
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_APPIDMISMATCH = "UnsupportedOperation.AppIdMismatch"
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
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_APPIDMISMATCH = "UnsupportedOperation.AppIdMismatch"
//  UNSUPPORTEDOPERATION_CCNROUTETABLENOTEXIST = "UnsupportedOperation.CcnRouteTableNotExist"
//  UNSUPPORTEDOPERATION_UNABLECROSSBORDER = "UnsupportedOperation.UnableCrossBorder"
func (c *Client) CreateDirectConnectGatewayWithContext(ctx context.Context, request *CreateDirectConnectGatewayRequest) (response *CreateDirectConnectGatewayResponse, err error) {
    if request == nil {
        request = NewCreateDirectConnectGatewayRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateDirectConnectGateway")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateDirectConnectGatewayCcnRoutes")
    
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
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
//  LIMITEXCEEDED_ACTIONLIMITED = "LimitExceeded.ActionLimited"
//  LIMITEXCEEDED_TAGKEYEXCEEDED = "LimitExceeded.TagKeyExceeded"
//  LIMITEXCEEDED_TAGKEYPERRESOURCEEXCEEDED = "LimitExceeded.TagKeyPerResourceExceeded"
//  LIMITEXCEEDED_TAGNOTENOUGHQUOTA = "LimitExceeded.TagNotEnoughQuota"
//  LIMITEXCEEDED_TAGQUOTA = "LimitExceeded.TagQuota"
//  LIMITEXCEEDED_TAGQUOTAEXCEEDED = "LimitExceeded.TagQuotaExceeded"
//  LIMITEXCEEDED_TAGTAGSEXCEEDED = "LimitExceeded.TagTagsExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_DPDKNATFLOWLOGONLYSUPPORTALLTRAFFICTYPE = "UnsupportedOperation.DpdkNatFlowLogOnlySupportAllTrafficType"
//  UNSUPPORTEDOPERATION_FLOWLOGINSTANCEEXISTED = "UnsupportedOperation.FlowLogInstanceExisted"
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
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
//  LIMITEXCEEDED_ACTIONLIMITED = "LimitExceeded.ActionLimited"
//  LIMITEXCEEDED_TAGKEYEXCEEDED = "LimitExceeded.TagKeyExceeded"
//  LIMITEXCEEDED_TAGKEYPERRESOURCEEXCEEDED = "LimitExceeded.TagKeyPerResourceExceeded"
//  LIMITEXCEEDED_TAGNOTENOUGHQUOTA = "LimitExceeded.TagNotEnoughQuota"
//  LIMITEXCEEDED_TAGQUOTA = "LimitExceeded.TagQuota"
//  LIMITEXCEEDED_TAGQUOTAEXCEEDED = "LimitExceeded.TagQuotaExceeded"
//  LIMITEXCEEDED_TAGTAGSEXCEEDED = "LimitExceeded.TagTagsExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_DPDKNATFLOWLOGONLYSUPPORTALLTRAFFICTYPE = "UnsupportedOperation.DpdkNatFlowLogOnlySupportAllTrafficType"
//  UNSUPPORTEDOPERATION_FLOWLOGINSTANCEEXISTED = "UnsupportedOperation.FlowLogInstanceExisted"
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateFlowLog")
    
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
// This API is used to create a highly available virtual IP (HAVIP).
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_INVALIDBUSINESS = "InvalidParameterValue.InvalidBusiness"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_RESERVED = "InvalidParameterValue.Reserved"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_SUBNETNOTEXISTS = "UnsupportedOperation.SubnetNotExists"
func (c *Client) CreateHaVip(request *CreateHaVipRequest) (response *CreateHaVipResponse, err error) {
    return c.CreateHaVipWithContext(context.Background(), request)
}

// CreateHaVip
// This API is used to create a highly available virtual IP (HAVIP).
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_INVALIDBUSINESS = "InvalidParameterValue.InvalidBusiness"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_RESERVED = "InvalidParameterValue.Reserved"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_SUBNETNOTEXISTS = "UnsupportedOperation.SubnetNotExists"
func (c *Client) CreateHaVipWithContext(ctx context.Context, request *CreateHaVipRequest) (response *CreateHaVipResponse, err error) {
    if request == nil {
        request = NewCreateHaVipRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateHaVip")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateLocalGateway")
    
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDACCOUNT_NOTSUPPORTED = "InvalidAccount.NotSupported"
//  INVALIDADDRESSSTATE = "InvalidAddressState"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_EIPBRANDWIDTHOUTINVALID = "InvalidParameterValue.EIPBrandWidthOutInvalid"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
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
//  RESOURCEINSUFFICIENT_INSTANCE = "ResourceInsufficient.Instance"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOREALNAMEAUTHENTICATION = "UnauthorizedOperation.NoRealNameAuthentication"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INSUFFICIENTFUNDS = "UnsupportedOperation.InsufficientFunds"
//  UNSUPPORTEDOPERATION_INSUFFICIENTINTERNETSERVICEPROVIDERS = "UnsupportedOperation.InsufficientInternetServiceProviders"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_NATGATEWAYHADEIPUNASSOCIATE = "UnsupportedOperation.NatGatewayHadEipUnassociate"
//  UNSUPPORTEDOPERATION_NOTSUPPORTZONE = "UnsupportedOperation.NotSupportZone"
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDACCOUNT_NOTSUPPORTED = "InvalidAccount.NotSupported"
//  INVALIDADDRESSSTATE = "InvalidAddressState"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_EIPBRANDWIDTHOUTINVALID = "InvalidParameterValue.EIPBrandWidthOutInvalid"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
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
//  RESOURCEINSUFFICIENT_INSTANCE = "ResourceInsufficient.Instance"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOREALNAMEAUTHENTICATION = "UnauthorizedOperation.NoRealNameAuthentication"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INSUFFICIENTFUNDS = "UnsupportedOperation.InsufficientFunds"
//  UNSUPPORTEDOPERATION_INSUFFICIENTINTERNETSERVICEPROVIDERS = "UnsupportedOperation.InsufficientInternetServiceProviders"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_NATGATEWAYHADEIPUNASSOCIATE = "UnsupportedOperation.NatGatewayHadEipUnassociate"
//  UNSUPPORTEDOPERATION_NOTSUPPORTZONE = "UnsupportedOperation.NotSupportZone"
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateNatGateway")
    
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
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NATGATEWAYDNATRULEEXISTED = "InvalidParameterValue.NatGatewayDnatRuleExisted"
//  INVALIDPARAMETERVALUE_NATGATEWAYDNATRULEPIPNEEDVM = "InvalidParameterValue.NatGatewayDnatRulePipNeedVm"
//  INVALIDPARAMETERVALUE_NATGATEWAYDNATRULEREPEATED = "InvalidParameterValue.NatGatewayDnatRuleRepeated"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_NATGATEWAYDNATLIMITEXCEEDED = "LimitExceeded.NatGatewayDnatLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_NATGATEWAYEIPNOTEXISTS = "UnsupportedOperation.NatGatewayEipNotExists"
func (c *Client) CreateNatGatewayDestinationIpPortTranslationNatRule(request *CreateNatGatewayDestinationIpPortTranslationNatRuleRequest) (response *CreateNatGatewayDestinationIpPortTranslationNatRuleResponse, err error) {
    return c.CreateNatGatewayDestinationIpPortTranslationNatRuleWithContext(context.Background(), request)
}

// CreateNatGatewayDestinationIpPortTranslationNatRule
// This API is used to create the port forwarding rules of a NAT gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NATGATEWAYDNATRULEEXISTED = "InvalidParameterValue.NatGatewayDnatRuleExisted"
//  INVALIDPARAMETERVALUE_NATGATEWAYDNATRULEPIPNEEDVM = "InvalidParameterValue.NatGatewayDnatRulePipNeedVm"
//  INVALIDPARAMETERVALUE_NATGATEWAYDNATRULEREPEATED = "InvalidParameterValue.NatGatewayDnatRuleRepeated"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_NATGATEWAYDNATLIMITEXCEEDED = "LimitExceeded.NatGatewayDnatLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_NATGATEWAYEIPNOTEXISTS = "UnsupportedOperation.NatGatewayEipNotExists"
func (c *Client) CreateNatGatewayDestinationIpPortTranslationNatRuleWithContext(ctx context.Context, request *CreateNatGatewayDestinationIpPortTranslationNatRuleRequest) (response *CreateNatGatewayDestinationIpPortTranslationNatRuleResponse, err error) {
    if request == nil {
        request = NewCreateNatGatewayDestinationIpPortTranslationNatRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateNatGatewayDestinationIpPortTranslationNatRule")
    
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
//  UNSUPPORTEDOPERATION_NATGATEWAYEIPNOTEXISTS = "UnsupportedOperation.NatGatewayEipNotExists"
//  UNSUPPORTEDOPERATION_NATGATEWAYRULEPIPEXISTS = "UnsupportedOperation.NatGatewayRulePipExists"
//  UNSUPPORTEDOPERATION_NATGATEWAYSNATPIPNEEDVM = "UnsupportedOperation.NatGatewaySnatPipNeedVm"
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
//  UNSUPPORTEDOPERATION_NATGATEWAYEIPNOTEXISTS = "UnsupportedOperation.NatGatewayEipNotExists"
//  UNSUPPORTEDOPERATION_NATGATEWAYRULEPIPEXISTS = "UnsupportedOperation.NatGatewayRulePipExists"
//  UNSUPPORTEDOPERATION_NATGATEWAYSNATPIPNEEDVM = "UnsupportedOperation.NatGatewaySnatPipNeedVm"
//  UNSUPPORTEDOPERATION_NATGATEWAYTYPENOTSUPPORTSNAT = "UnsupportedOperation.NatGatewayTypeNotSupportSNAT"
//  UNSUPPORTEDOPERATION_UNBINDEIP = "UnsupportedOperation.UnbindEIP"
func (c *Client) CreateNatGatewaySourceIpTranslationNatRuleWithContext(ctx context.Context, request *CreateNatGatewaySourceIpTranslationNatRuleRequest) (response *CreateNatGatewaySourceIpTranslationNatRuleResponse, err error) {
    if request == nil {
        request = NewCreateNatGatewaySourceIpTranslationNatRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateNatGatewaySourceIpTranslationNatRule")
    
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
// This API is used to create a network probe.
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CONFLICTWITHDOCKERROUTE = "UnsupportedOperation.ConflictWithDockerRoute"
//  UNSUPPORTEDOPERATION_ECMPWITHUSERROUTE = "UnsupportedOperation.EcmpWithUserRoute"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
func (c *Client) CreateNetDetect(request *CreateNetDetectRequest) (response *CreateNetDetectResponse, err error) {
    return c.CreateNetDetectWithContext(context.Background(), request)
}

// CreateNetDetect
// This API is used to create a network probe.
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CONFLICTWITHDOCKERROUTE = "UnsupportedOperation.ConflictWithDockerRoute"
//  UNSUPPORTEDOPERATION_ECMPWITHUSERROUTE = "UnsupportedOperation.EcmpWithUserRoute"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
func (c *Client) CreateNetDetectWithContext(ctx context.Context, request *CreateNetDetectRequest) (response *CreateNetDetectResponse, err error) {
    if request == nil {
        request = NewCreateNetDetectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateNetDetect")
    
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
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_APPIDMISMATCH = "UnsupportedOperation.AppIdMismatch"
//  UNSUPPORTEDOPERATION_TAGALLOCATE = "UnsupportedOperation.TagAllocate"
//  UNSUPPORTEDOPERATION_TAGFREE = "UnsupportedOperation.TagFree"
//  UNSUPPORTEDOPERATION_TAGNOTPERMIT = "UnsupportedOperation.TagNotPermit"
//  UNSUPPORTEDOPERATION_TAGSYSTEMRESERVEDTAGKEY = "UnsupportedOperation.TagSystemReservedTagKey"
func (c *Client) CreateNetworkAclWithContext(ctx context.Context, request *CreateNetworkAclRequest) (response *CreateNetworkAclResponse, err error) {
    if request == nil {
        request = NewCreateNetworkAclRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateNetworkAcl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNetworkAcl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNetworkAclResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNetworkAclQuintupleEntriesRequest() (request *CreateNetworkAclQuintupleEntriesRequest) {
    request = &CreateNetworkAclQuintupleEntriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "CreateNetworkAclQuintupleEntries")
    
    
    return
}

func NewCreateNetworkAclQuintupleEntriesResponse() (response *CreateNetworkAclQuintupleEntriesResponse) {
    response = &CreateNetworkAclQuintupleEntriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateNetworkAclQuintupleEntries
// This API is used to add one or more in/outbound rules of the network ACL quintuple.
//
// error code that may be returned:
//  INVALIDPARAMETER_ACLTYPEMISMATCH = "InvalidParameter.AclTypeMismatch"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_APPIDMISMATCH = "UnsupportedOperation.AppIdMismatch"
func (c *Client) CreateNetworkAclQuintupleEntries(request *CreateNetworkAclQuintupleEntriesRequest) (response *CreateNetworkAclQuintupleEntriesResponse, err error) {
    return c.CreateNetworkAclQuintupleEntriesWithContext(context.Background(), request)
}

// CreateNetworkAclQuintupleEntries
// This API is used to add one or more in/outbound rules of the network ACL quintuple.
//
// error code that may be returned:
//  INVALIDPARAMETER_ACLTYPEMISMATCH = "InvalidParameter.AclTypeMismatch"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_APPIDMISMATCH = "UnsupportedOperation.AppIdMismatch"
func (c *Client) CreateNetworkAclQuintupleEntriesWithContext(ctx context.Context, request *CreateNetworkAclQuintupleEntriesRequest) (response *CreateNetworkAclQuintupleEntriesResponse, err error) {
    if request == nil {
        request = NewCreateNetworkAclQuintupleEntriesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateNetworkAclQuintupleEntries")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNetworkAclQuintupleEntries require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNetworkAclQuintupleEntriesResponse()
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
//  FAILEDOPERATION_NETDETECTTIMEOUT = "FailedOperation.NetDetectTimeOut"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETER_FILTERNOTDICT = "InvalidParameter.FilterNotDict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_RESERVED = "InvalidParameterValue.Reserved"
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
//  LIMITEXCEEDED_ACTIONLIMITED = "LimitExceeded.ActionLimited"
//  LIMITEXCEEDED_NETWORKINTERFACELIMITEXCEEDED = "LimitExceeded.NetworkInterfaceLimitExceeded"
//  LIMITEXCEEDED_TAGKEYEXCEEDED = "LimitExceeded.TagKeyExceeded"
//  LIMITEXCEEDED_TAGKEYPERRESOURCEEXCEEDED = "LimitExceeded.TagKeyPerResourceExceeded"
//  LIMITEXCEEDED_TAGNOTENOUGHQUOTA = "LimitExceeded.TagNotEnoughQuota"
//  LIMITEXCEEDED_TAGQUOTA = "LimitExceeded.TagQuota"
//  LIMITEXCEEDED_TAGQUOTAEXCEEDED = "LimitExceeded.TagQuotaExceeded"
//  LIMITEXCEEDED_TAGTAGSEXCEEDED = "LimitExceeded.TagTagsExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_IDEMPOTENTPARAMETERMISMATCH = "UnsupportedOperation.IdempotentParameterMismatch"
//  UNSUPPORTEDOPERATION_IDEMPOTENTPROCESSING = "UnsupportedOperation.IdempotentProcessing"
//  UNSUPPORTEDOPERATION_RESOURCEMISMATCH = "UnsupportedOperation.ResourceMismatch"
//  UNSUPPORTEDOPERATION_TAGALLOCATE = "UnsupportedOperation.TagAllocate"
//  UNSUPPORTEDOPERATION_TAGFREE = "UnsupportedOperation.TagFree"
//  UNSUPPORTEDOPERATION_TAGNOTPERMIT = "UnsupportedOperation.TagNotPermit"
//  UNSUPPORTEDOPERATION_TAGSYSTEMRESERVEDTAGKEY = "UnsupportedOperation.TagSystemReservedTagKey"
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
//  FAILEDOPERATION_NETDETECTTIMEOUT = "FailedOperation.NetDetectTimeOut"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETER_FILTERNOTDICT = "InvalidParameter.FilterNotDict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_RESERVED = "InvalidParameterValue.Reserved"
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
//  LIMITEXCEEDED_ACTIONLIMITED = "LimitExceeded.ActionLimited"
//  LIMITEXCEEDED_NETWORKINTERFACELIMITEXCEEDED = "LimitExceeded.NetworkInterfaceLimitExceeded"
//  LIMITEXCEEDED_TAGKEYEXCEEDED = "LimitExceeded.TagKeyExceeded"
//  LIMITEXCEEDED_TAGKEYPERRESOURCEEXCEEDED = "LimitExceeded.TagKeyPerResourceExceeded"
//  LIMITEXCEEDED_TAGNOTENOUGHQUOTA = "LimitExceeded.TagNotEnoughQuota"
//  LIMITEXCEEDED_TAGQUOTA = "LimitExceeded.TagQuota"
//  LIMITEXCEEDED_TAGQUOTAEXCEEDED = "LimitExceeded.TagQuotaExceeded"
//  LIMITEXCEEDED_TAGTAGSEXCEEDED = "LimitExceeded.TagTagsExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_IDEMPOTENTPARAMETERMISMATCH = "UnsupportedOperation.IdempotentParameterMismatch"
//  UNSUPPORTEDOPERATION_IDEMPOTENTPROCESSING = "UnsupportedOperation.IdempotentProcessing"
//  UNSUPPORTEDOPERATION_RESOURCEMISMATCH = "UnsupportedOperation.ResourceMismatch"
//  UNSUPPORTEDOPERATION_TAGALLOCATE = "UnsupportedOperation.TagAllocate"
//  UNSUPPORTEDOPERATION_TAGFREE = "UnsupportedOperation.TagFree"
//  UNSUPPORTEDOPERATION_TAGNOTPERMIT = "UnsupportedOperation.TagNotPermit"
//  UNSUPPORTEDOPERATION_TAGSYSTEMRESERVEDTAGKEY = "UnsupportedOperation.TagSystemReservedTagKey"
func (c *Client) CreateNetworkInterfaceWithContext(ctx context.Context, request *CreateNetworkInterfaceRequest) (response *CreateNetworkInterfaceResponse, err error) {
    if request == nil {
        request = NewCreateNetworkInterfaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateNetworkInterface")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNetworkInterface require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNetworkInterfaceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateReserveIpAddressesRequest() (request *CreateReserveIpAddressesRequest) {
    request = &CreateReserveIpAddressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "CreateReserveIpAddresses")
    
    
    return
}

func NewCreateReserveIpAddressesResponse() (response *CreateReserveIpAddressesResponse) {
    response = &CreateReserveIpAddressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateReserveIpAddresses
// This API is used to create a reserved private IP address.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNKNOWNPARAMETER_WITHGUESS = "UnknownParameter.WithGuess"
func (c *Client) CreateReserveIpAddresses(request *CreateReserveIpAddressesRequest) (response *CreateReserveIpAddressesResponse, err error) {
    return c.CreateReserveIpAddressesWithContext(context.Background(), request)
}

// CreateReserveIpAddresses
// This API is used to create a reserved private IP address.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNKNOWNPARAMETER_WITHGUESS = "UnknownParameter.WithGuess"
func (c *Client) CreateReserveIpAddressesWithContext(ctx context.Context, request *CreateReserveIpAddressesRequest) (response *CreateReserveIpAddressesResponse, err error) {
    if request == nil {
        request = NewCreateReserveIpAddressesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateReserveIpAddresses")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateReserveIpAddresses require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateReserveIpAddressesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRoutePolicyRequest() (request *CreateRoutePolicyRequest) {
    request = &CreateRoutePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "CreateRoutePolicy")
    
    
    return
}

func NewCreateRoutePolicyResponse() (response *CreateRoutePolicyResponse) {
    response = &CreateRoutePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRoutePolicy
// This API is used to create a VPC route reception policy, including name, description and policy entries.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MUSTHASONE = "InvalidParameterValue.MustHasOne"
//  INVALIDPARAMETERVALUE_ROUTEPOLICYPRIORITYCONFLICT = "InvalidParameterValue.RoutePolicyPriorityConflict"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateRoutePolicy(request *CreateRoutePolicyRequest) (response *CreateRoutePolicyResponse, err error) {
    return c.CreateRoutePolicyWithContext(context.Background(), request)
}

// CreateRoutePolicy
// This API is used to create a VPC route reception policy, including name, description and policy entries.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MUSTHASONE = "InvalidParameterValue.MustHasOne"
//  INVALIDPARAMETERVALUE_ROUTEPOLICYPRIORITYCONFLICT = "InvalidParameterValue.RoutePolicyPriorityConflict"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateRoutePolicyWithContext(ctx context.Context, request *CreateRoutePolicyRequest) (response *CreateRoutePolicyResponse, err error) {
    if request == nil {
        request = NewCreateRoutePolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateRoutePolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRoutePolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRoutePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRoutePolicyAssociationsRequest() (request *CreateRoutePolicyAssociationsRequest) {
    request = &CreateRoutePolicyAssociationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "CreateRoutePolicyAssociations")
    
    
    return
}

func NewCreateRoutePolicyAssociationsResponse() (response *CreateRoutePolicyAssociationsResponse) {
    response = &CreateRoutePolicyAssociationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRoutePolicyAssociations
// This API is used to create route reception policy bindings (the binding relationship between policy instances and route table instances as well as set priorities).
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_ROUTEPOLICYASSOCIATIONEXISTS = "InvalidParameterValue.RoutePolicyAssociationExists"
//  INVALIDPARAMETERVALUE_ROUTEPOLICYPRIORITYCONFLICT = "InvalidParameterValue.RoutePolicyPriorityConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateRoutePolicyAssociations(request *CreateRoutePolicyAssociationsRequest) (response *CreateRoutePolicyAssociationsResponse, err error) {
    return c.CreateRoutePolicyAssociationsWithContext(context.Background(), request)
}

// CreateRoutePolicyAssociations
// This API is used to create route reception policy bindings (the binding relationship between policy instances and route table instances as well as set priorities).
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_ROUTEPOLICYASSOCIATIONEXISTS = "InvalidParameterValue.RoutePolicyAssociationExists"
//  INVALIDPARAMETERVALUE_ROUTEPOLICYPRIORITYCONFLICT = "InvalidParameterValue.RoutePolicyPriorityConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateRoutePolicyAssociationsWithContext(ctx context.Context, request *CreateRoutePolicyAssociationsRequest) (response *CreateRoutePolicyAssociationsResponse, err error) {
    if request == nil {
        request = NewCreateRoutePolicyAssociationsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateRoutePolicyAssociations")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRoutePolicyAssociations require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRoutePolicyAssociationsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRoutePolicyEntriesRequest() (request *CreateRoutePolicyEntriesRequest) {
    request = &CreateRoutePolicyEntriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "CreateRoutePolicyEntries")
    
    
    return
}

func NewCreateRoutePolicyEntriesResponse() (response *CreateRoutePolicyEntriesResponse) {
    response = &CreateRoutePolicyEntriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRoutePolicyEntries
// This API is used to create route reception policy entries.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_ROUTEPOLICYPRIORITYCONFLICT = "InvalidParameterValue.RoutePolicyPriorityConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateRoutePolicyEntries(request *CreateRoutePolicyEntriesRequest) (response *CreateRoutePolicyEntriesResponse, err error) {
    return c.CreateRoutePolicyEntriesWithContext(context.Background(), request)
}

// CreateRoutePolicyEntries
// This API is used to create route reception policy entries.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_ROUTEPOLICYPRIORITYCONFLICT = "InvalidParameterValue.RoutePolicyPriorityConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateRoutePolicyEntriesWithContext(ctx context.Context, request *CreateRoutePolicyEntriesRequest) (response *CreateRoutePolicyEntriesResponse, err error) {
    if request == nil {
        request = NewCreateRoutePolicyEntriesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateRoutePolicyEntries")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRoutePolicyEntries require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRoutePolicyEntriesResponse()
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
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
func (c *Client) CreateRouteTableWithContext(ctx context.Context, request *CreateRouteTableRequest) (response *CreateRouteTableResponse, err error) {
    if request == nil {
        request = NewCreateRouteTableRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateRouteTable")
    
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
// This API is used to create routes. 
//
// * You can batch add routes to a specified route table.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CIDRNOTINPEERVPC = "InvalidParameterValue.CidrNotInPeerVpc"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_ROUTEPOLICYPRIORITYCONFLICT = "InvalidParameterValue.RoutePolicyPriorityConflict"
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
// This API is used to create routes. 
//
// * You can batch add routes to a specified route table.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CIDRNOTINPEERVPC = "InvalidParameterValue.CidrNotInPeerVpc"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_ROUTEPOLICYPRIORITYCONFLICT = "InvalidParameterValue.RoutePolicyPriorityConflict"
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateRoutes")
    
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateSecurityGroup")
    
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
// This API is used to create security group policies (`SecurityGroupPolicy`).
//
// 
//
// For parameters of `SecurityGroupPolicySet`,
//
// <ul>
//
// <li>`Version`: The version number of a security group policy, which automatically increases by one each time you update the security policy, to prevent expiration of the updated routing policies. If it is left empty, any conflicts will be ignored.</li>
//
// <li>When creating the `Egress` and `Ingress` polices,<ul>
//
// <li>`Protocol`: Allows `TCP`, `UDP`, `ICMP`, `ICMPV6`, `GRE` and `ALL`.</li>
//
// <li>`CidrBlock`: For the classic network, the `CidrBlock` can contain private IPs of Tencent Cloud resources that are not under your account. It does not mean that you can access these devices. The network isolation policies between tenants take priority over the private network policies in security groups.</li>
//
// <li>`Ipv6CidrBlock`: For the classic network, `Ipv6CidrBlock` can contain private IPv6 addresses of Tencent Cloud resources that are not under your account. It does not mean that you can access these devices. The network isolation policies between tenants take priority over the private network policies in security groups.</li>
//
// <li>`SecurityGroupId`: ID of the security group to create policies. </li>
//
// <li>`Port`: A single port (“80”) or a port range ("8000-8010"). This parameter is only available when `Protocol` is `TCP` or `UDP`.</li>
//
// <li>`Action`: `ACCEPT` or `DROP`.</li>
//
// <li><code>CidrBlock</code>, <code>Ipv6CidrBlock</code>, <code>SecurityGroupId</code>, and <code>AddressTemplate</code> are mutually exclusive. <code>Protocol</code> + <code>Port</code> and <code>ServiceTemplate</code> are mutually exclusive. <code>IPv6CidrBlock</code> and <code>ICMP</code> are mutually exclusive; to use them, enter <code>ICMPV6</code>.</li>
//
// <li>You can only create policies in one direction in each request. To specify the `PolicyIndex` parameter, use the same index number in policies. If you want to insert a rule before the first rule, enter 0; if you want to add a rule after the last rule, leave it empty.</li>
//
// </ul></li></ul>
//
// error code that may be returned:
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONIDMISMATCH = "InvalidParameterValue.MemberApprovalApplicationIdMismatch"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONNOTAPPROVED = "InvalidParameterValue.MemberApprovalApplicationNotApproved"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONREJECTED = "InvalidParameterValue.MemberApprovalApplicationRejected"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONSTARTED = "InvalidParameterValue.MemberApprovalApplicationStarted"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_SECURITYGROUPPOLICYSET = "LimitExceeded.SecurityGroupPolicySet"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER_WITHGUESS = "UnknownParameter.WithGuess"
//  UNSUPPORTEDOPERATION_CLBPOLICYEXCEEDLIMIT = "UnsupportedOperation.ClbPolicyExceedLimit"
//  UNSUPPORTEDOPERATION_CLBPOLICYLIMIT = "UnsupportedOperation.ClbPolicyLimit"
//  UNSUPPORTEDOPERATION_DUPLICATEPOLICY = "UnsupportedOperation.DuplicatePolicy"
//  UNSUPPORTEDOPERATION_LOCKEDRESOURCES = "UnsupportedOperation.LockedResources"
//  UNSUPPORTEDOPERATION_VERSIONMISMATCH = "UnsupportedOperation.VersionMismatch"
func (c *Client) CreateSecurityGroupPolicies(request *CreateSecurityGroupPoliciesRequest) (response *CreateSecurityGroupPoliciesResponse, err error) {
    return c.CreateSecurityGroupPoliciesWithContext(context.Background(), request)
}

// CreateSecurityGroupPolicies
// This API is used to create security group policies (`SecurityGroupPolicy`).
//
// 
//
// For parameters of `SecurityGroupPolicySet`,
//
// <ul>
//
// <li>`Version`: The version number of a security group policy, which automatically increases by one each time you update the security policy, to prevent expiration of the updated routing policies. If it is left empty, any conflicts will be ignored.</li>
//
// <li>When creating the `Egress` and `Ingress` polices,<ul>
//
// <li>`Protocol`: Allows `TCP`, `UDP`, `ICMP`, `ICMPV6`, `GRE` and `ALL`.</li>
//
// <li>`CidrBlock`: For the classic network, the `CidrBlock` can contain private IPs of Tencent Cloud resources that are not under your account. It does not mean that you can access these devices. The network isolation policies between tenants take priority over the private network policies in security groups.</li>
//
// <li>`Ipv6CidrBlock`: For the classic network, `Ipv6CidrBlock` can contain private IPv6 addresses of Tencent Cloud resources that are not under your account. It does not mean that you can access these devices. The network isolation policies between tenants take priority over the private network policies in security groups.</li>
//
// <li>`SecurityGroupId`: ID of the security group to create policies. </li>
//
// <li>`Port`: A single port (“80”) or a port range ("8000-8010"). This parameter is only available when `Protocol` is `TCP` or `UDP`.</li>
//
// <li>`Action`: `ACCEPT` or `DROP`.</li>
//
// <li><code>CidrBlock</code>, <code>Ipv6CidrBlock</code>, <code>SecurityGroupId</code>, and <code>AddressTemplate</code> are mutually exclusive. <code>Protocol</code> + <code>Port</code> and <code>ServiceTemplate</code> are mutually exclusive. <code>IPv6CidrBlock</code> and <code>ICMP</code> are mutually exclusive; to use them, enter <code>ICMPV6</code>.</li>
//
// <li>You can only create policies in one direction in each request. To specify the `PolicyIndex` parameter, use the same index number in policies. If you want to insert a rule before the first rule, enter 0; if you want to add a rule after the last rule, leave it empty.</li>
//
// </ul></li></ul>
//
// error code that may be returned:
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONIDMISMATCH = "InvalidParameterValue.MemberApprovalApplicationIdMismatch"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONNOTAPPROVED = "InvalidParameterValue.MemberApprovalApplicationNotApproved"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONREJECTED = "InvalidParameterValue.MemberApprovalApplicationRejected"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONSTARTED = "InvalidParameterValue.MemberApprovalApplicationStarted"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_SECURITYGROUPPOLICYSET = "LimitExceeded.SecurityGroupPolicySet"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER_WITHGUESS = "UnknownParameter.WithGuess"
//  UNSUPPORTEDOPERATION_CLBPOLICYEXCEEDLIMIT = "UnsupportedOperation.ClbPolicyExceedLimit"
//  UNSUPPORTEDOPERATION_CLBPOLICYLIMIT = "UnsupportedOperation.ClbPolicyLimit"
//  UNSUPPORTEDOPERATION_DUPLICATEPOLICY = "UnsupportedOperation.DuplicatePolicy"
//  UNSUPPORTEDOPERATION_LOCKEDRESOURCES = "UnsupportedOperation.LockedResources"
//  UNSUPPORTEDOPERATION_VERSIONMISMATCH = "UnsupportedOperation.VersionMismatch"
func (c *Client) CreateSecurityGroupPoliciesWithContext(ctx context.Context, request *CreateSecurityGroupPoliciesRequest) (response *CreateSecurityGroupPoliciesResponse, err error) {
    if request == nil {
        request = NewCreateSecurityGroupPoliciesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateSecurityGroupPolicies")
    
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
// This API is used to create a security group, and add security group policies.
//
// * For the the upper limit of security groups per project in each region under each account, <a href="https://intl.cloud.tencent.com/document/product/213/12453?from_cn_redirect=1">see here</a>
//
// * For newly created security groups, the inbound and outbound policies are set to `Deny All` by default. You need to call <a href="https://intl.cloud.tencent.com/document/product/215/15807?from_cn_redirect=1">CreateSecurityGroupPolicies</a>
//
// to change it.
//
// 
//
// Description:
//
// * `Version`: The version number of a security group policy. It automatically increments by 1 every time you update the security policy, so to prevent the expiration of the updated policies. If this field is left empty, any conflicts will be ignored.
//
// * `Protocol`: Values can be `TCP`, `UDP`, `ICMP`, `ICMPV6`, `GRE`, and `ALL`.
//
// * `CidrBlock`: Enter a CIDR block in the correct format. In the classic network, even if the CIDR block specified in `CidrBlock` contains the Tencent Cloud private IPs not used for CVMs under your Tencent Cloud account, it does not mean this policy allows you to access those resources. The network isolation policies between tenants take priority over the private network policies in security groups.
//
// * `Ipv6CidrBlock`: Enter an IPv6 CIDR block in the correct format. In the classic network, even if the CIDR block specified in `Ipv6CidrBlock` contains the Tencent Cloud private IPv6 addresses not used for CVMs under your Tencent Cloud account, it does not mean this policy allows you to access those resources. The network isolation policies between tenants take priority over the private network policies in security groups.
//
// * `SecurityGroupId`: ID of the security group. It can be the ID of a security group to be modified, or the ID of another security group in the same project. All private IPs of all CVMs under the security group will be covered. If this field is used, the policy will automatically change according to the CVM associated with the group ID while being used to match network messages. You don't need to change it manually.
//
// * `Port`: Enter a single port number (such as `80`), or a port range (such as `8000-8010`). `Port` is only applicable when `Protocol` is `TCP` or `UDP`. If `Protocol` is not `TCP` or `UDP`, `Protocol` and `Port` cannot be both specified. 
//
// * `Action`: Values can be `ACCEPT` or `DROP`.
//
// * `CidrBlock`, `Ipv6CidrBlock`, `SecurityGroupId`, and `AddressTemplate` are exclusive to one another. "Protocol + Port" and `ServiceTemplate` are mutually exclusive.
//
// * Only policies in one direction can be created in each request. If you need to specify the `PolicyIndex` parameter, the indexes of policies must be consistent.
//
// error code that may be returned:
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONIDMISMATCH = "InvalidParameterValue.MemberApprovalApplicationIdMismatch"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONNOTAPPROVED = "InvalidParameterValue.MemberApprovalApplicationNotApproved"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONREJECTED = "InvalidParameterValue.MemberApprovalApplicationRejected"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONSTARTED = "InvalidParameterValue.MemberApprovalApplicationStarted"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER_WITHGUESS = "UnknownParameter.WithGuess"
//  UNSUPPORTEDOPERATION_LOCKEDRESOURCES = "UnsupportedOperation.LockedResources"
func (c *Client) CreateSecurityGroupWithPolicies(request *CreateSecurityGroupWithPoliciesRequest) (response *CreateSecurityGroupWithPoliciesResponse, err error) {
    return c.CreateSecurityGroupWithPoliciesWithContext(context.Background(), request)
}

// CreateSecurityGroupWithPolicies
// This API is used to create a security group, and add security group policies.
//
// * For the the upper limit of security groups per project in each region under each account, <a href="https://intl.cloud.tencent.com/document/product/213/12453?from_cn_redirect=1">see here</a>
//
// * For newly created security groups, the inbound and outbound policies are set to `Deny All` by default. You need to call <a href="https://intl.cloud.tencent.com/document/product/215/15807?from_cn_redirect=1">CreateSecurityGroupPolicies</a>
//
// to change it.
//
// 
//
// Description:
//
// * `Version`: The version number of a security group policy. It automatically increments by 1 every time you update the security policy, so to prevent the expiration of the updated policies. If this field is left empty, any conflicts will be ignored.
//
// * `Protocol`: Values can be `TCP`, `UDP`, `ICMP`, `ICMPV6`, `GRE`, and `ALL`.
//
// * `CidrBlock`: Enter a CIDR block in the correct format. In the classic network, even if the CIDR block specified in `CidrBlock` contains the Tencent Cloud private IPs not used for CVMs under your Tencent Cloud account, it does not mean this policy allows you to access those resources. The network isolation policies between tenants take priority over the private network policies in security groups.
//
// * `Ipv6CidrBlock`: Enter an IPv6 CIDR block in the correct format. In the classic network, even if the CIDR block specified in `Ipv6CidrBlock` contains the Tencent Cloud private IPv6 addresses not used for CVMs under your Tencent Cloud account, it does not mean this policy allows you to access those resources. The network isolation policies between tenants take priority over the private network policies in security groups.
//
// * `SecurityGroupId`: ID of the security group. It can be the ID of a security group to be modified, or the ID of another security group in the same project. All private IPs of all CVMs under the security group will be covered. If this field is used, the policy will automatically change according to the CVM associated with the group ID while being used to match network messages. You don't need to change it manually.
//
// * `Port`: Enter a single port number (such as `80`), or a port range (such as `8000-8010`). `Port` is only applicable when `Protocol` is `TCP` or `UDP`. If `Protocol` is not `TCP` or `UDP`, `Protocol` and `Port` cannot be both specified. 
//
// * `Action`: Values can be `ACCEPT` or `DROP`.
//
// * `CidrBlock`, `Ipv6CidrBlock`, `SecurityGroupId`, and `AddressTemplate` are exclusive to one another. "Protocol + Port" and `ServiceTemplate` are mutually exclusive.
//
// * Only policies in one direction can be created in each request. If you need to specify the `PolicyIndex` parameter, the indexes of policies must be consistent.
//
// error code that may be returned:
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONIDMISMATCH = "InvalidParameterValue.MemberApprovalApplicationIdMismatch"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONNOTAPPROVED = "InvalidParameterValue.MemberApprovalApplicationNotApproved"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONREJECTED = "InvalidParameterValue.MemberApprovalApplicationRejected"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONSTARTED = "InvalidParameterValue.MemberApprovalApplicationStarted"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER_WITHGUESS = "UnknownParameter.WithGuess"
//  UNSUPPORTEDOPERATION_LOCKEDRESOURCES = "UnsupportedOperation.LockedResources"
func (c *Client) CreateSecurityGroupWithPoliciesWithContext(ctx context.Context, request *CreateSecurityGroupWithPoliciesRequest) (response *CreateSecurityGroupWithPoliciesResponse, err error) {
    if request == nil {
        request = NewCreateSecurityGroupWithPoliciesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateSecurityGroupWithPolicies")
    
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
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateServiceTemplate(request *CreateServiceTemplateRequest) (response *CreateServiceTemplateResponse, err error) {
    return c.CreateServiceTemplateWithContext(context.Background(), request)
}

// CreateServiceTemplate
// This API (CreateServiceTemplate) is used to create a protocol port template.
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateServiceTemplateWithContext(ctx context.Context, request *CreateServiceTemplateRequest) (response *CreateServiceTemplateResponse, err error) {
    if request == nil {
        request = NewCreateServiceTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateServiceTemplate")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateServiceTemplateGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateServiceTemplateGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateServiceTemplateGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSnapshotPoliciesRequest() (request *CreateSnapshotPoliciesRequest) {
    request = &CreateSnapshotPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "CreateSnapshotPolicies")
    
    
    return
}

func NewCreateSnapshotPoliciesResponse() (response *CreateSnapshotPoliciesResponse) {
    response = &CreateSnapshotPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSnapshotPolicies
// This API is used to create snapshot policies.
//
// error code that may be returned:
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSnapshotPolicies(request *CreateSnapshotPoliciesRequest) (response *CreateSnapshotPoliciesResponse, err error) {
    return c.CreateSnapshotPoliciesWithContext(context.Background(), request)
}

// CreateSnapshotPolicies
// This API is used to create snapshot policies.
//
// error code that may be returned:
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSnapshotPoliciesWithContext(ctx context.Context, request *CreateSnapshotPoliciesRequest) (response *CreateSnapshotPoliciesResponse, err error) {
    if request == nil {
        request = NewCreateSnapshotPoliciesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateSnapshotPolicies")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSnapshotPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSnapshotPoliciesResponse()
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
//  FAILEDOPERATION_NETDETECTTIMEOUT = "FailedOperation.NetDetectTimeOut"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_SUBNETCONFLICT = "InvalidParameterValue.SubnetConflict"
//  INVALIDPARAMETERVALUE_SUBNETOVERLAP = "InvalidParameterValue.SubnetOverlap"
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
//  UNSUPPORTEDOPERATION_RECORDEXISTS = "UnsupportedOperation.RecordExists"
//  UNSUPPORTEDOPERATION_RECORDNOTEXISTS = "UnsupportedOperation.RecordNotExists"
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
//  FAILEDOPERATION_NETDETECTTIMEOUT = "FailedOperation.NetDetectTimeOut"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_SUBNETCONFLICT = "InvalidParameterValue.SubnetConflict"
//  INVALIDPARAMETERVALUE_SUBNETOVERLAP = "InvalidParameterValue.SubnetOverlap"
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
//  UNSUPPORTEDOPERATION_RECORDEXISTS = "UnsupportedOperation.RecordExists"
//  UNSUPPORTEDOPERATION_RECORDNOTEXISTS = "UnsupportedOperation.RecordNotExists"
//  UNSUPPORTEDOPERATION_TAGALLOCATE = "UnsupportedOperation.TagAllocate"
//  UNSUPPORTEDOPERATION_TAGFREE = "UnsupportedOperation.TagFree"
//  UNSUPPORTEDOPERATION_TAGNOTPERMIT = "UnsupportedOperation.TagNotPermit"
//  UNSUPPORTEDOPERATION_TAGSYSTEMRESERVEDTAGKEY = "UnsupportedOperation.TagSystemReservedTagKey"
func (c *Client) CreateSubnetWithContext(ctx context.Context, request *CreateSubnetRequest) (response *CreateSubnetResponse, err error) {
    if request == nil {
        request = NewCreateSubnetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateSubnet")
    
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
//  UNKNOWNPARAMETER_WITHGUESS = "UnknownParameter.WithGuess"
//  UNSUPPORTEDOPERATION_APPIDMISMATCH = "UnsupportedOperation.AppIdMismatch"
//  UNSUPPORTEDOPERATION_DCGATEWAYSNOTFOUNDINVPC = "UnsupportedOperation.DcGatewaysNotFoundInVpc"
//  UNSUPPORTEDOPERATION_RECORDEXISTS = "UnsupportedOperation.RecordExists"
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
//  UNKNOWNPARAMETER_WITHGUESS = "UnknownParameter.WithGuess"
//  UNSUPPORTEDOPERATION_APPIDMISMATCH = "UnsupportedOperation.AppIdMismatch"
//  UNSUPPORTEDOPERATION_DCGATEWAYSNOTFOUNDINVPC = "UnsupportedOperation.DcGatewaysNotFoundInVpc"
//  UNSUPPORTEDOPERATION_RECORDEXISTS = "UnsupportedOperation.RecordExists"
//  UNSUPPORTEDOPERATION_TAGALLOCATE = "UnsupportedOperation.TagAllocate"
//  UNSUPPORTEDOPERATION_TAGFREE = "UnsupportedOperation.TagFree"
//  UNSUPPORTEDOPERATION_TAGNOTPERMIT = "UnsupportedOperation.TagNotPermit"
//  UNSUPPORTEDOPERATION_TAGSYSTEMRESERVEDTAGKEY = "UnsupportedOperation.TagSystemReservedTagKey"
func (c *Client) CreateSubnetsWithContext(ctx context.Context, request *CreateSubnetsRequest) (response *CreateSubnetsResponse, err error) {
    if request == nil {
        request = NewCreateSubnetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateSubnets")
    
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
// * The subnet mask of the smallest IP address range that can be created is 28 (16 IP addresses), that of the largest IP address ranges 10.0.0.0/12 and 172.16.0.0/12 is 12 (1,048,576 IP addresses), and that of the largest IP address range 192.168.0.0/16 is 16 (65,536 IP addresses). For more information on how to plan VPC IP ranges, see [Network Planning](https://intl.cloud.tencent.com/document/product/215/30313?from_cn_redirect=1).
//
// * The number of VPC instances that can be created in a region is limited. For more information, see <a href="https://intl.cloud.tencent.com/doc/product/215/537?from_cn_redirect=1" title="VPC Use Limits">VPC Use Limits</a>. To request more resources, [submit a ticket](https://console.cloud.tencent.com/workorder/category).
//
// * You can bind tags when creating a VPC instance. The tag list in the response indicates the tags that have been successfully added.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_SUBNETOVERLAP = "InvalidParameterValue.SubnetOverlap"
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_APPIDMISMATCH = "UnsupportedOperation.AppIdMismatch"
//  UNSUPPORTEDOPERATION_ENABLEMULTICAST = "UnsupportedOperation.EnableMulticast"
//  UNSUPPORTEDOPERATION_RECORDEXISTS = "UnsupportedOperation.RecordExists"
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
// * The subnet mask of the smallest IP address range that can be created is 28 (16 IP addresses), that of the largest IP address ranges 10.0.0.0/12 and 172.16.0.0/12 is 12 (1,048,576 IP addresses), and that of the largest IP address range 192.168.0.0/16 is 16 (65,536 IP addresses). For more information on how to plan VPC IP ranges, see [Network Planning](https://intl.cloud.tencent.com/document/product/215/30313?from_cn_redirect=1).
//
// * The number of VPC instances that can be created in a region is limited. For more information, see <a href="https://intl.cloud.tencent.com/doc/product/215/537?from_cn_redirect=1" title="VPC Use Limits">VPC Use Limits</a>. To request more resources, [submit a ticket](https://console.cloud.tencent.com/workorder/category).
//
// * You can bind tags when creating a VPC instance. The tag list in the response indicates the tags that have been successfully added.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_SUBNETOVERLAP = "InvalidParameterValue.SubnetOverlap"
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_APPIDMISMATCH = "UnsupportedOperation.AppIdMismatch"
//  UNSUPPORTEDOPERATION_ENABLEMULTICAST = "UnsupportedOperation.EnableMulticast"
//  UNSUPPORTEDOPERATION_RECORDEXISTS = "UnsupportedOperation.RecordExists"
//  UNSUPPORTEDOPERATION_TAGALLOCATE = "UnsupportedOperation.TagAllocate"
//  UNSUPPORTEDOPERATION_TAGFREE = "UnsupportedOperation.TagFree"
//  UNSUPPORTEDOPERATION_TAGNOTPERMIT = "UnsupportedOperation.TagNotPermit"
//  UNSUPPORTEDOPERATION_TAGSYSTEMRESERVEDTAGKEY = "UnsupportedOperation.TagSystemReservedTagKey"
func (c *Client) CreateVpcWithContext(ctx context.Context, request *CreateVpcRequest) (response *CreateVpcResponse, err error) {
    if request == nil {
        request = NewCreateVpcRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateVpc")
    
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
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SERVICEWHITELISTNOTADDED = "ResourceUnavailable.ServiceWhiteListNotAdded"
//  UNAUTHORIZEDOPERATION_NOREALNAMEAUTHENTICATION = "UnauthorizedOperation.NoRealNameAuthentication"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ENDPOINTMISMATCHENDPOINTSERVICECDCID = "UnsupportedOperation.EndPointMismatchEndPointServiceCdcId"
//  UNSUPPORTEDOPERATION_ENDPOINTSERVICE = "UnsupportedOperation.EndPointService"
//  UNSUPPORTEDOPERATION_INSUFFICIENTFUNDS = "UnsupportedOperation.InsufficientFunds"
//  UNSUPPORTEDOPERATION_SNATSUBNET = "UnsupportedOperation.SnatSubnet"
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
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SERVICEWHITELISTNOTADDED = "ResourceUnavailable.ServiceWhiteListNotAdded"
//  UNAUTHORIZEDOPERATION_NOREALNAMEAUTHENTICATION = "UnauthorizedOperation.NoRealNameAuthentication"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ENDPOINTMISMATCHENDPOINTSERVICECDCID = "UnsupportedOperation.EndPointMismatchEndPointServiceCdcId"
//  UNSUPPORTEDOPERATION_ENDPOINTSERVICE = "UnsupportedOperation.EndPointService"
//  UNSUPPORTEDOPERATION_INSUFFICIENTFUNDS = "UnsupportedOperation.InsufficientFunds"
//  UNSUPPORTEDOPERATION_SNATSUBNET = "UnsupportedOperation.SnatSubnet"
//  UNSUPPORTEDOPERATION_SPECIALENDPOINTSERVICE = "UnsupportedOperation.SpecialEndPointService"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
func (c *Client) CreateVpcEndPointWithContext(ctx context.Context, request *CreateVpcEndPointRequest) (response *CreateVpcEndPointResponse, err error) {
    if request == nil {
        request = NewCreateVpcEndPointRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateVpcEndPoint")
    
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INSTANCEMISMATCH = "UnsupportedOperation.InstanceMismatch"
//  UNSUPPORTEDOPERATION_NOTMATCHTARGETSERVICE = "UnsupportedOperation.NotMatchTargetService"
//  UNSUPPORTEDOPERATION_RESOURCEISINVALIDSTATE = "UnsupportedOperation.ResourceIsInvalidState"
//  UNSUPPORTEDOPERATION_ROLENOTFOUND = "UnsupportedOperation.RoleNotFound"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
func (c *Client) CreateVpcEndPointService(request *CreateVpcEndPointServiceRequest) (response *CreateVpcEndPointServiceResponse, err error) {
    return c.CreateVpcEndPointServiceWithContext(context.Background(), request)
}

// CreateVpcEndPointService
// This API is used to create an endpoint service.
//
// error code that may be returned:
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INSTANCEMISMATCH = "UnsupportedOperation.InstanceMismatch"
//  UNSUPPORTEDOPERATION_NOTMATCHTARGETSERVICE = "UnsupportedOperation.NotMatchTargetService"
//  UNSUPPORTEDOPERATION_RESOURCEISINVALIDSTATE = "UnsupportedOperation.ResourceIsInvalidState"
//  UNSUPPORTEDOPERATION_ROLENOTFOUND = "UnsupportedOperation.RoleNotFound"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
func (c *Client) CreateVpcEndPointServiceWithContext(ctx context.Context, request *CreateVpcEndPointServiceRequest) (response *CreateVpcEndPointServiceResponse, err error) {
    if request == nil {
        request = NewCreateVpcEndPointServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateVpcEndPointService")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateVpcEndPointServiceWhiteList")
    
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
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
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
//  INVALIDPARAMETERVALUE_VPNCONNBGPTUNNELCIDRCONFLICT = "InvalidParameterValue.VpnConnBgpTunnelCidrConflict"
//  INVALIDPARAMETERVALUE_VPNCONNBGPTUNNELCIDRMASK = "InvalidParameterValue.VpnConnBgpTunnelCidrMask"
//  INVALIDPARAMETERVALUE_VPNCONNBGPTUNNELCIDRNOTSUPPORTED = "InvalidParameterValue.VpnConnBgpTunnelCidrNotSupported"
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
//  UNSUPPORTEDOPERATION_VPNUNSUPPORTEDBGP = "UnsupportedOperation.VpnUnsupportedBgp"
//  UNSUPPORTEDOPERATION_VPNUNSUPPORTEDBGPASNEQUAL = "UnsupportedOperation.VpnUnsupportedBgpAsnEqual"
//  UNSUPPORTEDOPERATION_VPNUNSUPPORTEDNOTEXISTBGPASN = "UnsupportedOperation.VpnUnsupportedNotExistBgpAsn"
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
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
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
//  INVALIDPARAMETERVALUE_VPNCONNBGPTUNNELCIDRCONFLICT = "InvalidParameterValue.VpnConnBgpTunnelCidrConflict"
//  INVALIDPARAMETERVALUE_VPNCONNBGPTUNNELCIDRMASK = "InvalidParameterValue.VpnConnBgpTunnelCidrMask"
//  INVALIDPARAMETERVALUE_VPNCONNBGPTUNNELCIDRNOTSUPPORTED = "InvalidParameterValue.VpnConnBgpTunnelCidrNotSupported"
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
//  UNSUPPORTEDOPERATION_VPNUNSUPPORTEDBGP = "UnsupportedOperation.VpnUnsupportedBgp"
//  UNSUPPORTEDOPERATION_VPNUNSUPPORTEDBGPASNEQUAL = "UnsupportedOperation.VpnUnsupportedBgpAsnEqual"
//  UNSUPPORTEDOPERATION_VPNUNSUPPORTEDNOTEXISTBGPASN = "UnsupportedOperation.VpnUnsupportedNotExistBgpAsn"
func (c *Client) CreateVpnConnectionWithContext(ctx context.Context, request *CreateVpnConnectionRequest) (response *CreateVpnConnectionResponse, err error) {
    if request == nil {
        request = NewCreateVpnConnectionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateVpnConnection")
    
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
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
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
//  INVALIDPARAMETERVALUE_VPNCONNCIDRCONFLICT = "InvalidParameterValue.VpnConnCidrConflict"
//  INVALIDVPCID_MALFORMED = "InvalidVpcId.Malformed"
//  INVALIDVPCID_NOTFOUND = "InvalidVpcId.NotFound"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TAGKEYEXCEEDED = "LimitExceeded.TagKeyExceeded"
//  LIMITEXCEEDED_TAGKEYPERRESOURCEEXCEEDED = "LimitExceeded.TagKeyPerResourceExceeded"
//  LIMITEXCEEDED_TAGNOTENOUGHQUOTA = "LimitExceeded.TagNotEnoughQuota"
//  LIMITEXCEEDED_TAGQUOTA = "LimitExceeded.TagQuota"
//  LIMITEXCEEDED_TAGQUOTAEXCEEDED = "LimitExceeded.TagQuotaExceeded"
//  LIMITEXCEEDED_TAGTAGSEXCEEDED = "LimitExceeded.TagTagsExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOREALNAMEAUTHENTICATION = "UnauthorizedOperation.NoRealNameAuthentication"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INSUFFICIENTFUNDS = "UnsupportedOperation.InsufficientFunds"
//  UNSUPPORTEDOPERATION_TAGALLOCATE = "UnsupportedOperation.TagAllocate"
//  UNSUPPORTEDOPERATION_TAGFREE = "UnsupportedOperation.TagFree"
//  UNSUPPORTEDOPERATION_TAGNOTPERMIT = "UnsupportedOperation.TagNotPermit"
//  UNSUPPORTEDOPERATION_TAGSYSTEMRESERVEDTAGKEY = "UnsupportedOperation.TagSystemReservedTagKey"
//  UNSUPPORTEDOPERATION_VPNGWVPCIDMUSTHAVE = "UnsupportedOperation.VpnGwVpcIdMustHave"
//  UNSUPPORTEDOPERATION_VPNUNSUPPORTEDBGP = "UnsupportedOperation.VpnUnsupportedBgp"
func (c *Client) CreateVpnGateway(request *CreateVpnGatewayRequest) (response *CreateVpnGatewayResponse, err error) {
    return c.CreateVpnGatewayWithContext(context.Background(), request)
}

// CreateVpnGateway
// This API (CreateVpnGateway) is used to create a VPN gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
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
//  INVALIDPARAMETERVALUE_VPNCONNCIDRCONFLICT = "InvalidParameterValue.VpnConnCidrConflict"
//  INVALIDVPCID_MALFORMED = "InvalidVpcId.Malformed"
//  INVALIDVPCID_NOTFOUND = "InvalidVpcId.NotFound"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TAGKEYEXCEEDED = "LimitExceeded.TagKeyExceeded"
//  LIMITEXCEEDED_TAGKEYPERRESOURCEEXCEEDED = "LimitExceeded.TagKeyPerResourceExceeded"
//  LIMITEXCEEDED_TAGNOTENOUGHQUOTA = "LimitExceeded.TagNotEnoughQuota"
//  LIMITEXCEEDED_TAGQUOTA = "LimitExceeded.TagQuota"
//  LIMITEXCEEDED_TAGQUOTAEXCEEDED = "LimitExceeded.TagQuotaExceeded"
//  LIMITEXCEEDED_TAGTAGSEXCEEDED = "LimitExceeded.TagTagsExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOREALNAMEAUTHENTICATION = "UnauthorizedOperation.NoRealNameAuthentication"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INSUFFICIENTFUNDS = "UnsupportedOperation.InsufficientFunds"
//  UNSUPPORTEDOPERATION_TAGALLOCATE = "UnsupportedOperation.TagAllocate"
//  UNSUPPORTEDOPERATION_TAGFREE = "UnsupportedOperation.TagFree"
//  UNSUPPORTEDOPERATION_TAGNOTPERMIT = "UnsupportedOperation.TagNotPermit"
//  UNSUPPORTEDOPERATION_TAGSYSTEMRESERVEDTAGKEY = "UnsupportedOperation.TagSystemReservedTagKey"
//  UNSUPPORTEDOPERATION_VPNGWVPCIDMUSTHAVE = "UnsupportedOperation.VpnGwVpcIdMustHave"
//  UNSUPPORTEDOPERATION_VPNUNSUPPORTEDBGP = "UnsupportedOperation.VpnUnsupportedBgp"
func (c *Client) CreateVpnGatewayWithContext(ctx context.Context, request *CreateVpnGatewayRequest) (response *CreateVpnGatewayResponse, err error) {
    if request == nil {
        request = NewCreateVpnGatewayRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateVpnGateway")
    
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
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
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
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "CreateVpnGatewayRoutes")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteAddressTemplate")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteAddressTemplateGroup")
    
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
// This API is used to delete a secondary CIDR block.
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
// This API is used to delete a secondary CIDR block.
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteAssistantCidr")
    
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
//  INVALIDPARAMETERVALUE_STOPCHARGINGINSTANCEINUSE = "InvalidParameterValue.StopChargingInstanceInUse"
//  LIMITEXCEEDED_ACCOUNTRETURNQUOTA = "LimitExceeded.AccountReturnQuota"
//  UNSUPPORTEDOPERATION_BANDWIDTHPACKAGEIDNOTSUPPORTED = "UnsupportedOperation.BandwidthPackageIdNotSupported"
//  UNSUPPORTEDOPERATION_INVALIDADDRESSSTATE = "UnsupportedOperation.InvalidAddressState"
//  UNSUPPORTEDOPERATION_UNPAIDORDERALREADYEXISTS = "UnsupportedOperation.UnpaidOrderAlreadyExists"
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
//  INVALIDPARAMETERVALUE_STOPCHARGINGINSTANCEINUSE = "InvalidParameterValue.StopChargingInstanceInUse"
//  LIMITEXCEEDED_ACCOUNTRETURNQUOTA = "LimitExceeded.AccountReturnQuota"
//  UNSUPPORTEDOPERATION_BANDWIDTHPACKAGEIDNOTSUPPORTED = "UnsupportedOperation.BandwidthPackageIdNotSupported"
//  UNSUPPORTEDOPERATION_INVALIDADDRESSSTATE = "UnsupportedOperation.InvalidAddressState"
//  UNSUPPORTEDOPERATION_UNPAIDORDERALREADYEXISTS = "UnsupportedOperation.UnpaidOrderAlreadyExists"
func (c *Client) DeleteBandwidthPackageWithContext(ctx context.Context, request *DeleteBandwidthPackageRequest) (response *DeleteBandwidthPackageResponse, err error) {
    if request == nil {
        request = NewDeleteBandwidthPackageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteBandwidthPackage")
    
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
//  INVALIDPARAMETERVALUE_PARAMETERMISMATCH = "InvalidParameterValue.ParameterMismatch"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_APPIDMISMATCH = "UnsupportedOperation.AppIdMismatch"
//  UNSUPPORTEDOPERATION_BANDWIDTHNOTEXPIRED = "UnsupportedOperation.BandwidthNotExpired"
//  UNSUPPORTEDOPERATION_CCNHASFLOWLOG = "UnsupportedOperation.CcnHasFlowLog"
//  UNSUPPORTEDOPERATION_ROUTETABLECANNOTDELETE = "UnsupportedOperation.RouteTableCanNotDelete"
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
//  INVALIDPARAMETERVALUE_PARAMETERMISMATCH = "InvalidParameterValue.ParameterMismatch"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_APPIDMISMATCH = "UnsupportedOperation.AppIdMismatch"
//  UNSUPPORTEDOPERATION_BANDWIDTHNOTEXPIRED = "UnsupportedOperation.BandwidthNotExpired"
//  UNSUPPORTEDOPERATION_CCNHASFLOWLOG = "UnsupportedOperation.CcnHasFlowLog"
//  UNSUPPORTEDOPERATION_ROUTETABLECANNOTDELETE = "UnsupportedOperation.RouteTableCanNotDelete"
func (c *Client) DeleteCcnWithContext(ctx context.Context, request *DeleteCcnRequest) (response *DeleteCcnResponse, err error) {
    if request == nil {
        request = NewDeleteCcnRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteCcn")
    
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
//  UNSUPPORTEDOPERATION_RECORDNOTEXISTS = "UnsupportedOperation.RecordNotExists"
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
//  UNSUPPORTEDOPERATION_RECORDNOTEXISTS = "UnsupportedOperation.RecordNotExists"
func (c *Client) DeleteCustomerGatewayWithContext(ctx context.Context, request *DeleteCustomerGatewayRequest) (response *DeleteCustomerGatewayResponse, err error) {
    if request == nil {
        request = NewDeleteCustomerGatewayRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteCustomerGateway")
    
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_APPIDMISMATCH = "UnsupportedOperation.AppIdMismatch"
//  UNSUPPORTEDOPERATION_DCGATEWAYNATRULEEXISTS = "UnsupportedOperation.DCGatewayNatRuleExists"
//  UNSUPPORTEDOPERATION_ROUTETABLEHASSUBNETRULE = "UnsupportedOperation.RouteTableHasSubnetRule"
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
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_APPIDMISMATCH = "UnsupportedOperation.AppIdMismatch"
//  UNSUPPORTEDOPERATION_DCGATEWAYNATRULEEXISTS = "UnsupportedOperation.DCGatewayNatRuleExists"
//  UNSUPPORTEDOPERATION_ROUTETABLEHASSUBNETRULE = "UnsupportedOperation.RouteTableHasSubnetRule"
func (c *Client) DeleteDirectConnectGatewayWithContext(ctx context.Context, request *DeleteDirectConnectGatewayRequest) (response *DeleteDirectConnectGatewayResponse, err error) {
    if request == nil {
        request = NewDeleteDirectConnectGatewayRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteDirectConnectGateway")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteDirectConnectGatewayCcnRoutes")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteFlowLog")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteHaVip")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteLocalGateway")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteNatGateway")
    
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
//  INVALIDPARAMETERVALUE_NATGATEWAYDNATRULENOTEXISTS = "InvalidParameterValue.NatGatewayDnatRuleNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteNatGatewayDestinationIpPortTranslationNatRule(request *DeleteNatGatewayDestinationIpPortTranslationNatRuleRequest) (response *DeleteNatGatewayDestinationIpPortTranslationNatRuleResponse, err error) {
    return c.DeleteNatGatewayDestinationIpPortTranslationNatRuleWithContext(context.Background(), request)
}

// DeleteNatGatewayDestinationIpPortTranslationNatRule
// This API is used to delete the port forwarding rule of a NAT gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_NATGATEWAYDNATRULENOTEXISTS = "InvalidParameterValue.NatGatewayDnatRuleNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteNatGatewayDestinationIpPortTranslationNatRuleWithContext(ctx context.Context, request *DeleteNatGatewayDestinationIpPortTranslationNatRuleRequest) (response *DeleteNatGatewayDestinationIpPortTranslationNatRuleResponse, err error) {
    if request == nil {
        request = NewDeleteNatGatewayDestinationIpPortTranslationNatRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteNatGatewayDestinationIpPortTranslationNatRule")
    
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
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NATGATEWAYSNATRULENOTEXISTS = "InvalidParameterValue.NatGatewaySnatRuleNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
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
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NATGATEWAYSNATRULENOTEXISTS = "InvalidParameterValue.NatGatewaySnatRuleNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteNatGatewaySourceIpTranslationNatRuleWithContext(ctx context.Context, request *DeleteNatGatewaySourceIpTranslationNatRuleRequest) (response *DeleteNatGatewaySourceIpTranslationNatRuleResponse, err error) {
    if request == nil {
        request = NewDeleteNatGatewaySourceIpTranslationNatRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteNatGatewaySourceIpTranslationNatRule")
    
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
// This API is used to delete a network probe.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_RECORDNOTEXISTS = "UnsupportedOperation.RecordNotExists"
func (c *Client) DeleteNetDetect(request *DeleteNetDetectRequest) (response *DeleteNetDetectResponse, err error) {
    return c.DeleteNetDetectWithContext(context.Background(), request)
}

// DeleteNetDetect
// This API is used to delete a network probe.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_RECORDNOTEXISTS = "UnsupportedOperation.RecordNotExists"
func (c *Client) DeleteNetDetectWithContext(ctx context.Context, request *DeleteNetDetectRequest) (response *DeleteNetDetectResponse, err error) {
    if request == nil {
        request = NewDeleteNetDetectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteNetDetect")
    
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
//  UNSUPPORTEDOPERATION_APPIDMISMATCH = "UnsupportedOperation.AppIdMismatch"
//  UNSUPPORTEDOPERATION_RECORDNOTEXISTS = "UnsupportedOperation.RecordNotExists"
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
//  UNSUPPORTEDOPERATION_APPIDMISMATCH = "UnsupportedOperation.AppIdMismatch"
//  UNSUPPORTEDOPERATION_RECORDNOTEXISTS = "UnsupportedOperation.RecordNotExists"
func (c *Client) DeleteNetworkAclWithContext(ctx context.Context, request *DeleteNetworkAclRequest) (response *DeleteNetworkAclResponse, err error) {
    if request == nil {
        request = NewDeleteNetworkAclRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteNetworkAcl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNetworkAcl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteNetworkAclResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNetworkAclQuintupleEntriesRequest() (request *DeleteNetworkAclQuintupleEntriesRequest) {
    request = &DeleteNetworkAclQuintupleEntriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteNetworkAclQuintupleEntries")
    
    
    return
}

func NewDeleteNetworkAclQuintupleEntriesResponse() (response *DeleteNetworkAclQuintupleEntriesResponse) {
    response = &DeleteNetworkAclQuintupleEntriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteNetworkAclQuintupleEntries
// This API is used to delete specified in/outbound rules of the network ACL quintuple. In the `NetworkAclQuintupleEntrySet` parameters, `NetworkAclQuintupleEntryId` is required for `NetworkAclQuintupleEntry`.
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
func (c *Client) DeleteNetworkAclQuintupleEntries(request *DeleteNetworkAclQuintupleEntriesRequest) (response *DeleteNetworkAclQuintupleEntriesResponse, err error) {
    return c.DeleteNetworkAclQuintupleEntriesWithContext(context.Background(), request)
}

// DeleteNetworkAclQuintupleEntries
// This API is used to delete specified in/outbound rules of the network ACL quintuple. In the `NetworkAclQuintupleEntrySet` parameters, `NetworkAclQuintupleEntryId` is required for `NetworkAclQuintupleEntry`.
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
func (c *Client) DeleteNetworkAclQuintupleEntriesWithContext(ctx context.Context, request *DeleteNetworkAclQuintupleEntriesRequest) (response *DeleteNetworkAclQuintupleEntriesResponse, err error) {
    if request == nil {
        request = NewDeleteNetworkAclQuintupleEntriesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteNetworkAclQuintupleEntries")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNetworkAclQuintupleEntries require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteNetworkAclQuintupleEntriesResponse()
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
//  LIMITEXCEEDED_ACTIONLIMITED = "LimitExceeded.ActionLimited"
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
//  LIMITEXCEEDED_ACTIONLIMITED = "LimitExceeded.ActionLimited"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) DeleteNetworkInterfaceWithContext(ctx context.Context, request *DeleteNetworkInterfaceRequest) (response *DeleteNetworkInterfaceResponse, err error) {
    if request == nil {
        request = NewDeleteNetworkInterfaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteNetworkInterface")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNetworkInterface require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteNetworkInterfaceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteReserveIpAddressesRequest() (request *DeleteReserveIpAddressesRequest) {
    request = &DeleteReserveIpAddressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteReserveIpAddresses")
    
    
    return
}

func NewDeleteReserveIpAddressesResponse() (response *DeleteReserveIpAddressesResponse) {
    response = &DeleteReserveIpAddressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteReserveIpAddresses
// This API is used to delete a reserved private IP address.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_PARAMETERMISMATCH = "InvalidParameterValue.ParameterMismatch"
//  INVALIDPARAMETERVALUE_RESOURCENOTFOUND = "InvalidParameterValue.ResourceNotFound"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNKNOWNPARAMETER_WITHGUESS = "UnknownParameter.WithGuess"
func (c *Client) DeleteReserveIpAddresses(request *DeleteReserveIpAddressesRequest) (response *DeleteReserveIpAddressesResponse, err error) {
    return c.DeleteReserveIpAddressesWithContext(context.Background(), request)
}

// DeleteReserveIpAddresses
// This API is used to delete a reserved private IP address.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_PARAMETERMISMATCH = "InvalidParameterValue.ParameterMismatch"
//  INVALIDPARAMETERVALUE_RESOURCENOTFOUND = "InvalidParameterValue.ResourceNotFound"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNKNOWNPARAMETER_WITHGUESS = "UnknownParameter.WithGuess"
func (c *Client) DeleteReserveIpAddressesWithContext(ctx context.Context, request *DeleteReserveIpAddressesRequest) (response *DeleteReserveIpAddressesResponse, err error) {
    if request == nil {
        request = NewDeleteReserveIpAddressesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteReserveIpAddresses")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteReserveIpAddresses require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteReserveIpAddressesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRoutePolicyRequest() (request *DeleteRoutePolicyRequest) {
    request = &DeleteRoutePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteRoutePolicy")
    
    
    return
}

func NewDeleteRoutePolicyResponse() (response *DeleteRoutePolicyResponse) {
    response = &DeleteRoutePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRoutePolicy
// This API is used to delete a route reception policy and entries.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_ROUTEPOLICYASSOCIATION = "UnsupportedOperation.RoutePolicyAssociation"
func (c *Client) DeleteRoutePolicy(request *DeleteRoutePolicyRequest) (response *DeleteRoutePolicyResponse, err error) {
    return c.DeleteRoutePolicyWithContext(context.Background(), request)
}

// DeleteRoutePolicy
// This API is used to delete a route reception policy and entries.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_ROUTEPOLICYASSOCIATION = "UnsupportedOperation.RoutePolicyAssociation"
func (c *Client) DeleteRoutePolicyWithContext(ctx context.Context, request *DeleteRoutePolicyRequest) (response *DeleteRoutePolicyResponse, err error) {
    if request == nil {
        request = NewDeleteRoutePolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteRoutePolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRoutePolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRoutePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRoutePolicyAssociationsRequest() (request *DeleteRoutePolicyAssociationsRequest) {
    request = &DeleteRoutePolicyAssociationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteRoutePolicyAssociations")
    
    
    return
}

func NewDeleteRoutePolicyAssociationsResponse() (response *DeleteRoutePolicyAssociationsResponse) {
    response = &DeleteRoutePolicyAssociationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRoutePolicyAssociations
// This API is used to delete route reception policy bindings (the binding relationship between route reception policy objects and route tables).
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_ROUTEPOLICYASSOCIATIONEXISTS = "InvalidParameterValue.RoutePolicyAssociationExists"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRoutePolicyAssociations(request *DeleteRoutePolicyAssociationsRequest) (response *DeleteRoutePolicyAssociationsResponse, err error) {
    return c.DeleteRoutePolicyAssociationsWithContext(context.Background(), request)
}

// DeleteRoutePolicyAssociations
// This API is used to delete route reception policy bindings (the binding relationship between route reception policy objects and route tables).
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_ROUTEPOLICYASSOCIATIONEXISTS = "InvalidParameterValue.RoutePolicyAssociationExists"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRoutePolicyAssociationsWithContext(ctx context.Context, request *DeleteRoutePolicyAssociationsRequest) (response *DeleteRoutePolicyAssociationsResponse, err error) {
    if request == nil {
        request = NewDeleteRoutePolicyAssociationsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteRoutePolicyAssociations")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRoutePolicyAssociations require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRoutePolicyAssociationsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRoutePolicyEntriesRequest() (request *DeleteRoutePolicyEntriesRequest) {
    request = &DeleteRoutePolicyEntriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteRoutePolicyEntries")
    
    
    return
}

func NewDeleteRoutePolicyEntriesResponse() (response *DeleteRoutePolicyEntriesResponse) {
    response = &DeleteRoutePolicyEntriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRoutePolicyEntries
// This API is used to delete route reception policy entries.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_MUSTHASONE = "InvalidParameterValue.MustHasOne"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRoutePolicyEntries(request *DeleteRoutePolicyEntriesRequest) (response *DeleteRoutePolicyEntriesResponse, err error) {
    return c.DeleteRoutePolicyEntriesWithContext(context.Background(), request)
}

// DeleteRoutePolicyEntries
// This API is used to delete route reception policy entries.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_MUSTHASONE = "InvalidParameterValue.MustHasOne"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRoutePolicyEntriesWithContext(ctx context.Context, request *DeleteRoutePolicyEntriesRequest) (response *DeleteRoutePolicyEntriesResponse, err error) {
    if request == nil {
        request = NewDeleteRoutePolicyEntriesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteRoutePolicyEntries")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRoutePolicyEntries require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRoutePolicyEntriesResponse()
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
//  UNSUPPORTEDOPERATION_ROUTETABLEHASSUBNETRULE = "UnsupportedOperation.RouteTableHasSubnetRule"
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
//  UNSUPPORTEDOPERATION_ROUTETABLEHASSUBNETRULE = "UnsupportedOperation.RouteTableHasSubnetRule"
func (c *Client) DeleteRouteTableWithContext(ctx context.Context, request *DeleteRouteTableRequest) (response *DeleteRouteTableResponse, err error) {
    if request == nil {
        request = NewDeleteRouteTableRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteRouteTable")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteRoutes")
    
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONIDMISMATCH = "InvalidParameterValue.MemberApprovalApplicationIdMismatch"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONNOTAPPROVED = "InvalidParameterValue.MemberApprovalApplicationNotApproved"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONREJECTED = "InvalidParameterValue.MemberApprovalApplicationRejected"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONSTARTED = "InvalidParameterValue.MemberApprovalApplicationStarted"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDSECURITYGROUPID_MALFORMED = "InvalidSecurityGroupID.Malformed"
//  INVALIDSECURITYGROUPID_NOTFOUND = "InvalidSecurityGroupID.NotFound"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_LOCKEDRESOURCES = "UnsupportedOperation.LockedResources"
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONIDMISMATCH = "InvalidParameterValue.MemberApprovalApplicationIdMismatch"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONNOTAPPROVED = "InvalidParameterValue.MemberApprovalApplicationNotApproved"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONREJECTED = "InvalidParameterValue.MemberApprovalApplicationRejected"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONSTARTED = "InvalidParameterValue.MemberApprovalApplicationStarted"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDSECURITYGROUPID_MALFORMED = "InvalidSecurityGroupID.Malformed"
//  INVALIDSECURITYGROUPID_NOTFOUND = "InvalidSecurityGroupID.NotFound"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_LOCKEDRESOURCES = "UnsupportedOperation.LockedResources"
func (c *Client) DeleteSecurityGroupWithContext(ctx context.Context, request *DeleteSecurityGroupRequest) (response *DeleteSecurityGroupResponse, err error) {
    if request == nil {
        request = NewDeleteSecurityGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteSecurityGroup")
    
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONIDMISMATCH = "InvalidParameterValue.MemberApprovalApplicationIdMismatch"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONNOTAPPROVED = "InvalidParameterValue.MemberApprovalApplicationNotApproved"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONREJECTED = "InvalidParameterValue.MemberApprovalApplicationRejected"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONSTARTED = "InvalidParameterValue.MemberApprovalApplicationStarted"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_LOCKEDRESOURCES = "UnsupportedOperation.LockedResources"
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONIDMISMATCH = "InvalidParameterValue.MemberApprovalApplicationIdMismatch"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONNOTAPPROVED = "InvalidParameterValue.MemberApprovalApplicationNotApproved"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONREJECTED = "InvalidParameterValue.MemberApprovalApplicationRejected"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONSTARTED = "InvalidParameterValue.MemberApprovalApplicationStarted"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_LOCKEDRESOURCES = "UnsupportedOperation.LockedResources"
//  UNSUPPORTEDOPERATION_VERSIONMISMATCH = "UnsupportedOperation.VersionMismatch"
func (c *Client) DeleteSecurityGroupPoliciesWithContext(ctx context.Context, request *DeleteSecurityGroupPoliciesRequest) (response *DeleteSecurityGroupPoliciesResponse, err error) {
    if request == nil {
        request = NewDeleteSecurityGroupPoliciesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteSecurityGroupPolicies")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteServiceTemplate")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteServiceTemplateGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteServiceTemplateGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteServiceTemplateGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSnapshotPoliciesRequest() (request *DeleteSnapshotPoliciesRequest) {
    request = &DeleteSnapshotPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteSnapshotPolicies")
    
    
    return
}

func NewDeleteSnapshotPoliciesResponse() (response *DeleteSnapshotPoliciesResponse) {
    response = &DeleteSnapshotPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSnapshotPolicies
// This API is used to delete snapshot policies.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteSnapshotPolicies(request *DeleteSnapshotPoliciesRequest) (response *DeleteSnapshotPoliciesResponse, err error) {
    return c.DeleteSnapshotPoliciesWithContext(context.Background(), request)
}

// DeleteSnapshotPolicies
// This API is used to delete snapshot policies.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteSnapshotPoliciesWithContext(ctx context.Context, request *DeleteSnapshotPoliciesRequest) (response *DeleteSnapshotPoliciesResponse, err error) {
    if request == nil {
        request = NewDeleteSnapshotPoliciesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteSnapshotPolicies")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSnapshotPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSnapshotPoliciesResponse()
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
// This API is used to delete a subnet.
//
// * Remove all resources in the subnet before deleting it
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_NATGATEWAYSNATRULENOTEXISTS = "UnsupportedOperation.NatGatewaySnatRuleNotExists"
//  UNSUPPORTEDOPERATION_RECORDNOTEXISTS = "UnsupportedOperation.RecordNotExists"
func (c *Client) DeleteSubnet(request *DeleteSubnetRequest) (response *DeleteSubnetResponse, err error) {
    return c.DeleteSubnetWithContext(context.Background(), request)
}

// DeleteSubnet
// This API is used to delete a subnet.
//
// * Remove all resources in the subnet before deleting it
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_NATGATEWAYSNATRULENOTEXISTS = "UnsupportedOperation.NatGatewaySnatRuleNotExists"
//  UNSUPPORTEDOPERATION_RECORDNOTEXISTS = "UnsupportedOperation.RecordNotExists"
func (c *Client) DeleteSubnetWithContext(ctx context.Context, request *DeleteSubnetRequest) (response *DeleteSubnetResponse, err error) {
    if request == nil {
        request = NewDeleteSubnetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteSubnet")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSubnet require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSubnetResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTemplateMemberRequest() (request *DeleteTemplateMemberRequest) {
    request = &DeleteTemplateMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteTemplateMember")
    
    
    return
}

func NewDeleteTemplateMemberResponse() (response *DeleteTemplateMemberResponse) {
    response = &DeleteTemplateMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTemplateMember
// This API is used to delete a parameter template of the IP address, protocol port, IP address group, or protocol port group type.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) DeleteTemplateMember(request *DeleteTemplateMemberRequest) (response *DeleteTemplateMemberResponse, err error) {
    return c.DeleteTemplateMemberWithContext(context.Background(), request)
}

// DeleteTemplateMember
// This API is used to delete a parameter template of the IP address, protocol port, IP address group, or protocol port group type.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) DeleteTemplateMemberWithContext(ctx context.Context, request *DeleteTemplateMemberRequest) (response *DeleteTemplateMemberResponse, err error) {
    if request == nil {
        request = NewDeleteTemplateMemberRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteTemplateMember")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTemplateMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTemplateMemberResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTrafficPackagesRequest() (request *DeleteTrafficPackagesRequest) {
    request = &DeleteTrafficPackagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteTrafficPackages")
    
    
    return
}

func NewDeleteTrafficPackagesResponse() (response *DeleteTrafficPackagesResponse) {
    response = &DeleteTrafficPackagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTrafficPackages
// This API is used to delete traffic packages. Note that only non-valid traffic packages can be deleted. 
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_TRAFFICPACKAGEIDMALFORMED = "InvalidParameterValue.TrafficPackageIdMalformed"
//  INVALIDPARAMETERVALUE_TRAFFICPACKAGENOTFOUND = "InvalidParameterValue.TrafficPackageNotFound"
//  INVALIDPARAMETERVALUE_TRAFFICPACKAGENOTSUPPORTED = "InvalidParameterValue.TrafficPackageNotSupported"
func (c *Client) DeleteTrafficPackages(request *DeleteTrafficPackagesRequest) (response *DeleteTrafficPackagesResponse, err error) {
    return c.DeleteTrafficPackagesWithContext(context.Background(), request)
}

// DeleteTrafficPackages
// This API is used to delete traffic packages. Note that only non-valid traffic packages can be deleted. 
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_TRAFFICPACKAGEIDMALFORMED = "InvalidParameterValue.TrafficPackageIdMalformed"
//  INVALIDPARAMETERVALUE_TRAFFICPACKAGENOTFOUND = "InvalidParameterValue.TrafficPackageNotFound"
//  INVALIDPARAMETERVALUE_TRAFFICPACKAGENOTSUPPORTED = "InvalidParameterValue.TrafficPackageNotSupported"
func (c *Client) DeleteTrafficPackagesWithContext(ctx context.Context, request *DeleteTrafficPackagesRequest) (response *DeleteTrafficPackagesResponse, err error) {
    if request == nil {
        request = NewDeleteTrafficPackagesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteTrafficPackages")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTrafficPackages require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTrafficPackagesResponse()
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
//  FAILEDOPERATION_NETDETECTTIMEOUT = "FailedOperation.NetDetectTimeOut"
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED_ACTIONLIMITED = "LimitExceeded.ActionLimited"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_APPIDMISMATCH = "UnsupportedOperation.AppIdMismatch"
//  UNSUPPORTEDOPERATION_RECORDNOTEXISTS = "UnsupportedOperation.RecordNotExists"
//  UNSUPPORTEDOPERATION_ROUTETABLEHASSUBNETRULE = "UnsupportedOperation.RouteTableHasSubnetRule"
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
//  FAILEDOPERATION_NETDETECTTIMEOUT = "FailedOperation.NetDetectTimeOut"
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED_ACTIONLIMITED = "LimitExceeded.ActionLimited"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_APPIDMISMATCH = "UnsupportedOperation.AppIdMismatch"
//  UNSUPPORTEDOPERATION_RECORDNOTEXISTS = "UnsupportedOperation.RecordNotExists"
//  UNSUPPORTEDOPERATION_ROUTETABLEHASSUBNETRULE = "UnsupportedOperation.RouteTableHasSubnetRule"
func (c *Client) DeleteVpcWithContext(ctx context.Context, request *DeleteVpcRequest) (response *DeleteVpcResponse, err error) {
    if request == nil {
        request = NewDeleteVpcRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteVpc")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteVpcEndPoint")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteVpcEndPointService")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteVpcEndPointServiceWhiteList")
    
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
// This API is used to delete a VPN tunnel.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_DELETEVPNCONNINVALIDSTATE = "UnsupportedOperation.DeleteVpnConnInvalidState"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
//  UNSUPPORTEDOPERATION_RECORDNOTEXISTS = "UnsupportedOperation.RecordNotExists"
func (c *Client) DeleteVpnConnection(request *DeleteVpnConnectionRequest) (response *DeleteVpnConnectionResponse, err error) {
    return c.DeleteVpnConnectionWithContext(context.Background(), request)
}

// DeleteVpnConnection
// This API is used to delete a VPN tunnel.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_DELETEVPNCONNINVALIDSTATE = "UnsupportedOperation.DeleteVpnConnInvalidState"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
//  UNSUPPORTEDOPERATION_RECORDNOTEXISTS = "UnsupportedOperation.RecordNotExists"
func (c *Client) DeleteVpnConnectionWithContext(ctx context.Context, request *DeleteVpnConnectionRequest) (response *DeleteVpnConnectionResponse, err error) {
    if request == nil {
        request = NewDeleteVpnConnectionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteVpnConnection")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteVpnGateway")
    
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
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
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
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteVpnGatewayRoutesWithContext(ctx context.Context, request *DeleteVpnGatewayRoutesRequest) (response *DeleteVpnGatewayRoutesResponse, err error) {
    if request == nil {
        request = NewDeleteVpnGatewayRoutesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DeleteVpnGatewayRoutes")
    
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
func (c *Client) DescribeAccountAttributes(request *DescribeAccountAttributesRequest) (response *DescribeAccountAttributesResponse, err error) {
    return c.DescribeAccountAttributesWithContext(context.Background(), request)
}

// DescribeAccountAttributes
// This API (DescribeAccountAttributes) is used to query your account attributes.
//
// error code that may be returned:
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
func (c *Client) DescribeAccountAttributesWithContext(ctx context.Context, request *DescribeAccountAttributesRequest) (response *DescribeAccountAttributesResponse, err error) {
    if request == nil {
        request = NewDescribeAccountAttributesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeAccountAttributes")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeAddressQuota")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeAddressTemplateGroups")
    
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAddressTemplatesWithContext(ctx context.Context, request *DescribeAddressTemplatesRequest) (response *DescribeAddressTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeAddressTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeAddressTemplates")
    
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
//  INVALIDPARAMETERVALUE_INVALIDDEDICATEDCLUSTERID = "InvalidParameterValue.InvalidDedicatedClusterId"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NETWORKINTERFACEIDMALFORMED = "InvalidParameterValue.NetworkInterfaceIdMalformed"
//  INVALIDPARAMETERVALUE_RESOURCEIDMALFORMED = "InvalidParameterValue.ResourceIdMalformed"
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
//  INVALIDPARAMETERVALUE_INVALIDDEDICATEDCLUSTERID = "InvalidParameterValue.InvalidDedicatedClusterId"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NETWORKINTERFACEIDMALFORMED = "InvalidParameterValue.NetworkInterfaceIdMalformed"
//  INVALIDPARAMETERVALUE_RESOURCEIDMALFORMED = "InvalidParameterValue.ResourceIdMalformed"
//  LIMITEXCEEDED_NUMBEROFFILTERS = "LimitExceeded.NumberOfFilters"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAddressesWithContext(ctx context.Context, request *DescribeAddressesRequest) (response *DescribeAddressesResponse, err error) {
    if request == nil {
        request = NewDescribeAddressesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeAddresses")
    
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
// This API is used to query the list of secondary CIDR blocks.
//
// error code that may be returned:
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
func (c *Client) DescribeAssistantCidr(request *DescribeAssistantCidrRequest) (response *DescribeAssistantCidrResponse, err error) {
    return c.DescribeAssistantCidrWithContext(context.Background(), request)
}

// DescribeAssistantCidr
// This API is used to query the list of secondary CIDR blocks.
//
// error code that may be returned:
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
func (c *Client) DescribeAssistantCidrWithContext(ctx context.Context, request *DescribeAssistantCidrRequest) (response *DescribeAssistantCidrResponse, err error) {
    if request == nil {
        request = NewDescribeAssistantCidrRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeAssistantCidr")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeBandwidthPackageBillUsage")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeBandwidthPackageQuota")
    
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
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGEIDMALFORMED = "InvalidParameterValue.BandwidthPackageIdMalformed"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGENOTFOUND = "InvalidParameterValue.BandwidthPackageNotFound"
//  INVALIDPARAMETERVALUE_ILLEGAL = "InvalidParameterValue.Illegal"
func (c *Client) DescribeBandwidthPackageResources(request *DescribeBandwidthPackageResourcesRequest) (response *DescribeBandwidthPackageResourcesResponse, err error) {
    return c.DescribeBandwidthPackageResourcesWithContext(context.Background(), request)
}

// DescribeBandwidthPackageResources
// This API is used to query resources in a bandwidth package based on the unique package ID. You can filter the result by specifying conditions and paginate the query results.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGEIDMALFORMED = "InvalidParameterValue.BandwidthPackageIdMalformed"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGENOTFOUND = "InvalidParameterValue.BandwidthPackageNotFound"
//  INVALIDPARAMETERVALUE_ILLEGAL = "InvalidParameterValue.Illegal"
func (c *Client) DescribeBandwidthPackageResourcesWithContext(ctx context.Context, request *DescribeBandwidthPackageResourcesRequest) (response *DescribeBandwidthPackageResourcesResponse, err error) {
    if request == nil {
        request = NewDescribeBandwidthPackageResourcesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeBandwidthPackageResources")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeBandwidthPackages")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeCcnAttachedInstances")
    
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCcnRegionBandwidthLimits(request *DescribeCcnRegionBandwidthLimitsRequest) (response *DescribeCcnRegionBandwidthLimitsResponse, err error) {
    return c.DescribeCcnRegionBandwidthLimitsWithContext(context.Background(), request)
}

// DescribeCcnRegionBandwidthLimits
// This API is used to query the outbound bandwidth caps of all regions connected with a CCN instance. The API only returns regions included in the associated network instances.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCcnRegionBandwidthLimitsWithContext(ctx context.Context, request *DescribeCcnRegionBandwidthLimitsRequest) (response *DescribeCcnRegionBandwidthLimitsResponse, err error) {
    if request == nil {
        request = NewDescribeCcnRegionBandwidthLimitsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeCcnRegionBandwidthLimits")
    
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
//  INVALIDPARAMETERVALUE_PARAMETERMISMATCH = "InvalidParameterValue.ParameterMismatch"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_APPIDNOTFOUND = "UnsupportedOperation.AppIdNotFound"
func (c *Client) DescribeCcnRoutes(request *DescribeCcnRoutesRequest) (response *DescribeCcnRoutesResponse, err error) {
    return c.DescribeCcnRoutesWithContext(context.Background(), request)
}

// DescribeCcnRoutes
// This API (DescribeCcnRoutes) is used to query routes that have been added to a CCN.
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_PARAMETERMISMATCH = "InvalidParameterValue.ParameterMismatch"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_APPIDNOTFOUND = "UnsupportedOperation.AppIdNotFound"
func (c *Client) DescribeCcnRoutesWithContext(ctx context.Context, request *DescribeCcnRoutesRequest) (response *DescribeCcnRoutesResponse, err error) {
    if request == nil {
        request = NewDescribeCcnRoutesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeCcnRoutes")
    
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
//  UNSUPPORTEDOPERATION_ACTIONNOTFOUND = "UnsupportedOperation.ActionNotFound"
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
//  UNSUPPORTEDOPERATION_ACTIONNOTFOUND = "UnsupportedOperation.ActionNotFound"
func (c *Client) DescribeCcnsWithContext(ctx context.Context, request *DescribeCcnsRequest) (response *DescribeCcnsResponse, err error) {
    if request == nil {
        request = NewDescribeCcnsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeCcns")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeClassicLinkInstances")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeCrossBorderCompliance")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeCustomerGatewayVendors")
    
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCustomerGateways(request *DescribeCustomerGatewaysRequest) (response *DescribeCustomerGatewaysResponse, err error) {
    return c.DescribeCustomerGatewaysWithContext(context.Background(), request)
}

// DescribeCustomerGateways
// This API (DescribeCustomerGateways) is used to query the customer gateway list.
//
// error code that may be returned:
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCustomerGatewaysWithContext(ctx context.Context, request *DescribeCustomerGatewaysRequest) (response *DescribeCustomerGatewaysResponse, err error) {
    if request == nil {
        request = NewDescribeCustomerGatewaysRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeCustomerGateways")
    
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
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDirectConnectGatewayCcnRoutes(request *DescribeDirectConnectGatewayCcnRoutesRequest) (response *DescribeDirectConnectGatewayCcnRoutesResponse, err error) {
    return c.DescribeDirectConnectGatewayCcnRoutesWithContext(context.Background(), request)
}

// DescribeDirectConnectGatewayCcnRoutes
// This API (DescribeDirectConnectGatewayCcnRoutes) is used to query the CCN routes (IDC IP range) of the Direct Connect gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDirectConnectGatewayCcnRoutesWithContext(ctx context.Context, request *DescribeDirectConnectGatewayCcnRoutesRequest) (response *DescribeDirectConnectGatewayCcnRoutesResponse, err error) {
    if request == nil {
        request = NewDescribeDirectConnectGatewayCcnRoutesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeDirectConnectGatewayCcnRoutes")
    
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETER_FILTERNOTDICT = "InvalidParameter.FilterNotDict"
//  INVALIDPARAMETER_FILTERVALUESNOTLIST = "InvalidParameter.FilterValuesNotList"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ACTIONNOTFOUND = "UnsupportedOperation.ActionNotFound"
func (c *Client) DescribeDirectConnectGateways(request *DescribeDirectConnectGatewaysRequest) (response *DescribeDirectConnectGatewaysResponse, err error) {
    return c.DescribeDirectConnectGatewaysWithContext(context.Background(), request)
}

// DescribeDirectConnectGateways
// This API is used to query direct connect gateways.
//
// error code that may be returned:
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETER_FILTERNOTDICT = "InvalidParameter.FilterNotDict"
//  INVALIDPARAMETER_FILTERVALUESNOTLIST = "InvalidParameter.FilterValuesNotList"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ACTIONNOTFOUND = "UnsupportedOperation.ActionNotFound"
func (c *Client) DescribeDirectConnectGatewaysWithContext(ctx context.Context, request *DescribeDirectConnectGatewaysRequest) (response *DescribeDirectConnectGatewaysResponse, err error) {
    if request == nil {
        request = NewDescribeDirectConnectGatewaysRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeDirectConnectGateways")
    
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
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeFlowLog(request *DescribeFlowLogRequest) (response *DescribeFlowLogResponse, err error) {
    return c.DescribeFlowLogWithContext(context.Background(), request)
}

// DescribeFlowLog
// This API is used to query the information of a flow log.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeFlowLogWithContext(ctx context.Context, request *DescribeFlowLogRequest) (response *DescribeFlowLogResponse, err error) {
    if request == nil {
        request = NewDescribeFlowLogRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeFlowLog")
    
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
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
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
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeFlowLogsWithContext(ctx context.Context, request *DescribeFlowLogsRequest) (response *DescribeFlowLogsResponse, err error) {
    if request == nil {
        request = NewDescribeFlowLogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeFlowLogs")
    
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeGatewayFlowMonitorDetailWithContext(ctx context.Context, request *DescribeGatewayFlowMonitorDetailRequest) (response *DescribeGatewayFlowMonitorDetailResponse, err error) {
    if request == nil {
        request = NewDescribeGatewayFlowMonitorDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeGatewayFlowMonitorDetail")
    
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
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
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
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) DescribeGatewayFlowQosWithContext(ctx context.Context, request *DescribeGatewayFlowQosRequest) (response *DescribeGatewayFlowQosResponse, err error) {
    if request == nil {
        request = NewDescribeGatewayFlowQosRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeGatewayFlowQos")
    
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
func (c *Client) DescribeHaVipsWithContext(ctx context.Context, request *DescribeHaVipsRequest) (response *DescribeHaVipsResponse, err error) {
    if request == nil {
        request = NewDescribeHaVipsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeHaVips")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHaVips require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHaVipsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIPv6AddressesRequest() (request *DescribeIPv6AddressesRequest) {
    request = &DescribeIPv6AddressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeIPv6Addresses")
    
    
    return
}

func NewDescribeIPv6AddressesResponse() (response *DescribeIPv6AddressesResponse) {
    response = &DescribeIPv6AddressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIPv6Addresses
// This API is used to query detailed information of one or more EIPv6 instances.
//
// 
//
// - You can query EIPv6 and traditional EIPv6 instance information in a specified region.
//
// - The system returns a certain number (as specified by the Limit, the default value is 20) of EIPv6 instances of the current user if the parameter is empty.
//
// error code that may be returned:
//  INVALIDADDRESSID_NOTFOUND = "InvalidAddressId.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
func (c *Client) DescribeIPv6Addresses(request *DescribeIPv6AddressesRequest) (response *DescribeIPv6AddressesResponse, err error) {
    return c.DescribeIPv6AddressesWithContext(context.Background(), request)
}

// DescribeIPv6Addresses
// This API is used to query detailed information of one or more EIPv6 instances.
//
// 
//
// - You can query EIPv6 and traditional EIPv6 instance information in a specified region.
//
// - The system returns a certain number (as specified by the Limit, the default value is 20) of EIPv6 instances of the current user if the parameter is empty.
//
// error code that may be returned:
//  INVALIDADDRESSID_NOTFOUND = "InvalidAddressId.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
func (c *Client) DescribeIPv6AddressesWithContext(ctx context.Context, request *DescribeIPv6AddressesRequest) (response *DescribeIPv6AddressesResponse, err error) {
    if request == nil {
        request = NewDescribeIPv6AddressesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeIPv6Addresses")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIPv6Addresses require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIPv6AddressesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceJumboRequest() (request *DescribeInstanceJumboRequest) {
    request = &DescribeInstanceJumboRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeInstanceJumbo")
    
    
    return
}

func NewDescribeInstanceJumboResponse() (response *DescribeInstanceJumboResponse) {
    response = &DescribeInstanceJumboResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceJumbo
// This API is used to check whether Cloud Virtual Machines support jumbo frames.
//
// Usage limits.
//
// This API is used to perform operations that require CAM policy authorization and read access to the corresponding instance. The API accesses CVM instances, so it verifies whether there are CAM permissions for the instance. For example: CAM action allows vpc:DescribeInstanceJumbo; resource allows qcs::cvm:ap-guangzhou:uin/2126195383:instance/*.
//
// This API is used to check the jumbo frame status before and after instance migration. The status returned by this API may be inconsistent before and after migration. You need to check whether the host machines of the instance before and after migration both support jumbo frames. One possible reason is that the instance has been migrated to a host machine that does not support jumbo frames.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeInstanceJumbo(request *DescribeInstanceJumboRequest) (response *DescribeInstanceJumboResponse, err error) {
    return c.DescribeInstanceJumboWithContext(context.Background(), request)
}

// DescribeInstanceJumbo
// This API is used to check whether Cloud Virtual Machines support jumbo frames.
//
// Usage limits.
//
// This API is used to perform operations that require CAM policy authorization and read access to the corresponding instance. The API accesses CVM instances, so it verifies whether there are CAM permissions for the instance. For example: CAM action allows vpc:DescribeInstanceJumbo; resource allows qcs::cvm:ap-guangzhou:uin/2126195383:instance/*.
//
// This API is used to check the jumbo frame status before and after instance migration. The status returned by this API may be inconsistent before and after migration. You need to check whether the host machines of the instance before and after migration both support jumbo frames. One possible reason is that the instance has been migrated to a host machine that does not support jumbo frames.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeInstanceJumboWithContext(ctx context.Context, request *DescribeInstanceJumboRequest) (response *DescribeInstanceJumboResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceJumboRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeInstanceJumbo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceJumbo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceJumboResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIp6AddressesRequest() (request *DescribeIp6AddressesRequest) {
    request = &DescribeIp6AddressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeIp6Addresses")
    
    
    return
}

func NewDescribeIp6AddressesResponse() (response *DescribeIp6AddressesResponse) {
    response = &DescribeIp6AddressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIp6Addresses
// This API is used to query the detailed information on one or multiple classic elastic public IPv6 instances.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDADDRESSID_NOTFOUND = "InvalidAddressId.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ADDRESSIPNOTPUBLIC = "InvalidParameterValue.AddressIpNotPublic"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NETWORKINTERFACEIDMALFORMED = "InvalidParameterValue.NetworkInterfaceIdMalformed"
func (c *Client) DescribeIp6Addresses(request *DescribeIp6AddressesRequest) (response *DescribeIp6AddressesResponse, err error) {
    return c.DescribeIp6AddressesWithContext(context.Background(), request)
}

// DescribeIp6Addresses
// This API is used to query the detailed information on one or multiple classic elastic public IPv6 instances.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDADDRESSID_NOTFOUND = "InvalidAddressId.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ADDRESSIPNOTPUBLIC = "InvalidParameterValue.AddressIpNotPublic"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NETWORKINTERFACEIDMALFORMED = "InvalidParameterValue.NetworkInterfaceIdMalformed"
func (c *Client) DescribeIp6AddressesWithContext(ctx context.Context, request *DescribeIp6AddressesRequest) (response *DescribeIp6AddressesResponse, err error) {
    if request == nil {
        request = NewDescribeIp6AddressesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeIp6Addresses")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIp6Addresses require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIp6AddressesResponse()
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
// <font color="#FF0000">This API will be discontinued soon and is only available for existing users.</font>
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
// <font color="#FF0000">This API will be discontinued soon and is only available for existing users.</font>
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeIpGeolocationDatabaseUrl")
    
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
// <font color="#FF0000">This API will be discontinued soon and is only available for existing users.</font>
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
// <font color="#FF0000">This API will be discontinued soon and is only available for existing users.</font>
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeIpGeolocationInfos")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeLocalGateway")
    
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
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
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
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeNatGatewayDestinationIpPortTranslationNatRulesWithContext(ctx context.Context, request *DescribeNatGatewayDestinationIpPortTranslationNatRulesRequest) (response *DescribeNatGatewayDestinationIpPortTranslationNatRulesResponse, err error) {
    if request == nil {
        request = NewDescribeNatGatewayDestinationIpPortTranslationNatRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeNatGatewayDestinationIpPortTranslationNatRules")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeNatGatewayDirectConnectGatewayRoute")
    
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
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
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
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNatGatewaySourceIpTranslationNatRulesWithContext(ctx context.Context, request *DescribeNatGatewaySourceIpTranslationNatRulesRequest) (response *DescribeNatGatewaySourceIpTranslationNatRulesResponse, err error) {
    if request == nil {
        request = NewDescribeNatGatewaySourceIpTranslationNatRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeNatGatewaySourceIpTranslationNatRules")
    
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETER_FILTERNOTDICT = "InvalidParameter.FilterNotDict"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ACTIONNOTFOUND = "UnsupportedOperation.ActionNotFound"
func (c *Client) DescribeNatGateways(request *DescribeNatGatewaysRequest) (response *DescribeNatGatewaysResponse, err error) {
    return c.DescribeNatGatewaysWithContext(context.Background(), request)
}

// DescribeNatGateways
// This API is used to query NAT gateways.
//
// error code that may be returned:
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETER_FILTERNOTDICT = "InvalidParameter.FilterNotDict"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ACTIONNOTFOUND = "UnsupportedOperation.ActionNotFound"
func (c *Client) DescribeNatGatewaysWithContext(ctx context.Context, request *DescribeNatGatewaysRequest) (response *DescribeNatGatewaysResponse, err error) {
    if request == nil {
        request = NewDescribeNatGatewaysRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeNatGateways")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeNetDetectStates")
    
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
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NOTUTF8ENCODINGERROR = "InvalidParameterValue.NotUtf8EncodingError"
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
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NOTUTF8ENCODINGERROR = "InvalidParameterValue.NotUtf8EncodingError"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNetDetectsWithContext(ctx context.Context, request *DescribeNetDetectsRequest) (response *DescribeNetDetectsResponse, err error) {
    if request == nil {
        request = NewDescribeNetDetectsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeNetDetects")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNetDetects require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNetDetectsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNetworkAclQuintupleEntriesRequest() (request *DescribeNetworkAclQuintupleEntriesRequest) {
    request = &DescribeNetworkAclQuintupleEntriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeNetworkAclQuintupleEntries")
    
    
    return
}

func NewDescribeNetworkAclQuintupleEntriesResponse() (response *DescribeNetworkAclQuintupleEntriesResponse) {
    response = &DescribeNetworkAclQuintupleEntriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNetworkAclQuintupleEntries
// This API is used to query the list of in/outbound network ACL quintuple entries.
//
// error code that may be returned:
//  INVALIDPARAMETER_ACLTYPEMISMATCH = "InvalidParameter.AclTypeMismatch"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NOTUTF8ENCODINGERROR = "InvalidParameterValue.NotUtf8EncodingError"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ACTIONNOTFOUND = "UnsupportedOperation.ActionNotFound"
func (c *Client) DescribeNetworkAclQuintupleEntries(request *DescribeNetworkAclQuintupleEntriesRequest) (response *DescribeNetworkAclQuintupleEntriesResponse, err error) {
    return c.DescribeNetworkAclQuintupleEntriesWithContext(context.Background(), request)
}

// DescribeNetworkAclQuintupleEntries
// This API is used to query the list of in/outbound network ACL quintuple entries.
//
// error code that may be returned:
//  INVALIDPARAMETER_ACLTYPEMISMATCH = "InvalidParameter.AclTypeMismatch"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NOTUTF8ENCODINGERROR = "InvalidParameterValue.NotUtf8EncodingError"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ACTIONNOTFOUND = "UnsupportedOperation.ActionNotFound"
func (c *Client) DescribeNetworkAclQuintupleEntriesWithContext(ctx context.Context, request *DescribeNetworkAclQuintupleEntriesRequest) (response *DescribeNetworkAclQuintupleEntriesResponse, err error) {
    if request == nil {
        request = NewDescribeNetworkAclQuintupleEntriesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeNetworkAclQuintupleEntries")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNetworkAclQuintupleEntries require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNetworkAclQuintupleEntriesResponse()
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NOTUTF8ENCODINGERROR = "InvalidParameterValue.NotUtf8EncodingError"
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NOTUTF8ENCODINGERROR = "InvalidParameterValue.NotUtf8EncodingError"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ACTIONNOTFOUND = "UnsupportedOperation.ActionNotFound"
func (c *Client) DescribeNetworkAclsWithContext(ctx context.Context, request *DescribeNetworkAclsRequest) (response *DescribeNetworkAclsResponse, err error) {
    if request == nil {
        request = NewDescribeNetworkAclsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeNetworkAcls")
    
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED_ACTIONLIMITED = "LimitExceeded.ActionLimited"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNetworkInterfaceLimit(request *DescribeNetworkInterfaceLimitRequest) (response *DescribeNetworkInterfaceLimitResponse, err error) {
    return c.DescribeNetworkInterfaceLimitWithContext(context.Background(), request)
}

// DescribeNetworkInterfaceLimit
// This API (DescribeNetworkInterfaceLimit) is used to query the ENI quota based on the ID of CVM instance or ENI. It returns the ENI quota to which the CVM instance can be bound and the IP address quota that can be allocated to the ENI.
//
// error code that may be returned:
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED_ACTIONLIMITED = "LimitExceeded.ActionLimited"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNetworkInterfaceLimitWithContext(ctx context.Context, request *DescribeNetworkInterfaceLimitRequest) (response *DescribeNetworkInterfaceLimitResponse, err error) {
    if request == nil {
        request = NewDescribeNetworkInterfaceLimitRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeNetworkInterfaceLimit")
    
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETER_FILTERNOTDICT = "InvalidParameter.FilterNotDict"
//  INVALIDPARAMETER_FILTERVALUESNOTLIST = "InvalidParameter.FilterValuesNotList"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NOTUTF8ENCODINGERROR = "InvalidParameterValue.NotUtf8EncodingError"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  LIMITEXCEEDED_ACTIONLIMITED = "LimitExceeded.ActionLimited"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNetworkInterfaces(request *DescribeNetworkInterfacesRequest) (response *DescribeNetworkInterfacesResponse, err error) {
    return c.DescribeNetworkInterfacesWithContext(context.Background(), request)
}

// DescribeNetworkInterfaces
// This API (DescribeNetworkInterfaces) is used to query the ENI list.
//
// error code that may be returned:
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETER_FILTERNOTDICT = "InvalidParameter.FilterNotDict"
//  INVALIDPARAMETER_FILTERVALUESNOTLIST = "InvalidParameter.FilterValuesNotList"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NOTUTF8ENCODINGERROR = "InvalidParameterValue.NotUtf8EncodingError"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  LIMITEXCEEDED_ACTIONLIMITED = "LimitExceeded.ActionLimited"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNetworkInterfacesWithContext(ctx context.Context, request *DescribeNetworkInterfacesRequest) (response *DescribeNetworkInterfacesResponse, err error) {
    if request == nil {
        request = NewDescribeNetworkInterfacesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeNetworkInterfaces")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNetworkInterfaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNetworkInterfacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReserveIpAddressesRequest() (request *DescribeReserveIpAddressesRequest) {
    request = &DescribeReserveIpAddressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeReserveIpAddresses")
    
    
    return
}

func NewDescribeReserveIpAddressesResponse() (response *DescribeReserveIpAddressesResponse) {
    response = &DescribeReserveIpAddressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReserveIpAddresses
// This API is used to query reserved private IP addresses.
//
// error code that may be returned:
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETER_FILTERNOTDICT = "InvalidParameter.FilterNotDict"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RESOURCENOTFOUND = "InvalidParameterValue.ResourceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeReserveIpAddresses(request *DescribeReserveIpAddressesRequest) (response *DescribeReserveIpAddressesResponse, err error) {
    return c.DescribeReserveIpAddressesWithContext(context.Background(), request)
}

// DescribeReserveIpAddresses
// This API is used to query reserved private IP addresses.
//
// error code that may be returned:
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETER_FILTERNOTDICT = "InvalidParameter.FilterNotDict"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RESOURCENOTFOUND = "InvalidParameterValue.ResourceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeReserveIpAddressesWithContext(ctx context.Context, request *DescribeReserveIpAddressesRequest) (response *DescribeReserveIpAddressesResponse, err error) {
    if request == nil {
        request = NewDescribeReserveIpAddressesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeReserveIpAddresses")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReserveIpAddresses require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReserveIpAddressesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRoutePolicyEntriesRequest() (request *DescribeRoutePolicyEntriesRequest) {
    request = &DescribeRoutePolicyEntriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeRoutePolicyEntries")
    
    
    return
}

func NewDescribeRoutePolicyEntriesResponse() (response *DescribeRoutePolicyEntriesResponse) {
    response = &DescribeRoutePolicyEntriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRoutePolicyEntries
// This API is used to query the route reception policy entry list.
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRoutePolicyEntries(request *DescribeRoutePolicyEntriesRequest) (response *DescribeRoutePolicyEntriesResponse, err error) {
    return c.DescribeRoutePolicyEntriesWithContext(context.Background(), request)
}

// DescribeRoutePolicyEntries
// This API is used to query the route reception policy entry list.
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRoutePolicyEntriesWithContext(ctx context.Context, request *DescribeRoutePolicyEntriesRequest) (response *DescribeRoutePolicyEntriesResponse, err error) {
    if request == nil {
        request = NewDescribeRoutePolicyEntriesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeRoutePolicyEntries")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRoutePolicyEntries require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRoutePolicyEntriesResponse()
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
// This API is used to query route tables.
//
// error code that may be returned:
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETER_FILTERNOTDICT = "InvalidParameter.FilterNotDict"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NOTUTF8ENCODINGERROR = "InvalidParameterValue.NotUtf8EncodingError"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRouteTables(request *DescribeRouteTablesRequest) (response *DescribeRouteTablesResponse, err error) {
    return c.DescribeRouteTablesWithContext(context.Background(), request)
}

// DescribeRouteTables
// This API is used to query route tables.
//
// error code that may be returned:
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETER_FILTERNOTDICT = "InvalidParameter.FilterNotDict"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NOTUTF8ENCODINGERROR = "InvalidParameterValue.NotUtf8EncodingError"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRouteTablesWithContext(ctx context.Context, request *DescribeRouteTablesRequest) (response *DescribeRouteTablesResponse, err error) {
    if request == nil {
        request = NewDescribeRouteTablesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeRouteTables")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeSecurityGroupAssociationStatistics")
    
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
//  INVALIDPARAMETER_FILTERNOTDICT = "InvalidParameter.FilterNotDict"
//  INVALIDPARAMETER_FILTERVALUESNOTLIST = "InvalidParameter.FilterValuesNotList"
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
//  INVALIDPARAMETER_FILTERNOTDICT = "InvalidParameter.FilterNotDict"
//  INVALIDPARAMETER_FILTERVALUESNOTLIST = "InvalidParameter.FilterValuesNotList"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSecurityGroupPoliciesWithContext(ctx context.Context, request *DescribeSecurityGroupPoliciesRequest) (response *DescribeSecurityGroupPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityGroupPoliciesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeSecurityGroupPolicies")
    
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSecurityGroupReferencesWithContext(ctx context.Context, request *DescribeSecurityGroupReferencesRequest) (response *DescribeSecurityGroupReferencesResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityGroupReferencesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeSecurityGroupReferences")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeSecurityGroups")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeServiceTemplateGroups")
    
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
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
func (c *Client) DescribeServiceTemplates(request *DescribeServiceTemplatesRequest) (response *DescribeServiceTemplatesResponse, err error) {
    return c.DescribeServiceTemplatesWithContext(context.Background(), request)
}

// DescribeServiceTemplates
// This API (DescribeServiceTemplates) is used to query protocol port templates.
//
// error code that may be returned:
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
func (c *Client) DescribeServiceTemplatesWithContext(ctx context.Context, request *DescribeServiceTemplatesRequest) (response *DescribeServiceTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeServiceTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeServiceTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServiceTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServiceTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSgSnapshotFileContentRequest() (request *DescribeSgSnapshotFileContentRequest) {
    request = &DescribeSgSnapshotFileContentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeSgSnapshotFileContent")
    
    
    return
}

func NewDescribeSgSnapshotFileContentResponse() (response *DescribeSgSnapshotFileContentResponse) {
    response = &DescribeSgSnapshotFileContentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSgSnapshotFileContent
// This API is used to query the snapshot file contents.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_SNAPSHOTFILEFAILED = "UnsupportedOperation.SnapshotFileFailed"
//  UNSUPPORTEDOPERATION_SNAPSHOTFILENOEXIST = "UnsupportedOperation.SnapshotFileNoExist"
//  UNSUPPORTEDOPERATION_SNAPSHOTFILEPROCESSING = "UnsupportedOperation.SnapshotFileProcessing"
func (c *Client) DescribeSgSnapshotFileContent(request *DescribeSgSnapshotFileContentRequest) (response *DescribeSgSnapshotFileContentResponse, err error) {
    return c.DescribeSgSnapshotFileContentWithContext(context.Background(), request)
}

// DescribeSgSnapshotFileContent
// This API is used to query the snapshot file contents.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_SNAPSHOTFILEFAILED = "UnsupportedOperation.SnapshotFileFailed"
//  UNSUPPORTEDOPERATION_SNAPSHOTFILENOEXIST = "UnsupportedOperation.SnapshotFileNoExist"
//  UNSUPPORTEDOPERATION_SNAPSHOTFILEPROCESSING = "UnsupportedOperation.SnapshotFileProcessing"
func (c *Client) DescribeSgSnapshotFileContentWithContext(ctx context.Context, request *DescribeSgSnapshotFileContentRequest) (response *DescribeSgSnapshotFileContentResponse, err error) {
    if request == nil {
        request = NewDescribeSgSnapshotFileContentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeSgSnapshotFileContent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSgSnapshotFileContent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSgSnapshotFileContentResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSnapshotAttachedInstancesRequest() (request *DescribeSnapshotAttachedInstancesRequest) {
    request = &DescribeSnapshotAttachedInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeSnapshotAttachedInstances")
    
    
    return
}

func NewDescribeSnapshotAttachedInstancesResponse() (response *DescribeSnapshotAttachedInstancesResponse) {
    response = &DescribeSnapshotAttachedInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSnapshotAttachedInstances
// This API is used to query instances associated with a snapshot policy.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSnapshotAttachedInstances(request *DescribeSnapshotAttachedInstancesRequest) (response *DescribeSnapshotAttachedInstancesResponse, err error) {
    return c.DescribeSnapshotAttachedInstancesWithContext(context.Background(), request)
}

// DescribeSnapshotAttachedInstances
// This API is used to query instances associated with a snapshot policy.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSnapshotAttachedInstancesWithContext(ctx context.Context, request *DescribeSnapshotAttachedInstancesRequest) (response *DescribeSnapshotAttachedInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeSnapshotAttachedInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeSnapshotAttachedInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSnapshotAttachedInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSnapshotAttachedInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSnapshotFilesRequest() (request *DescribeSnapshotFilesRequest) {
    request = &DescribeSnapshotFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeSnapshotFiles")
    
    
    return
}

func NewDescribeSnapshotFilesResponse() (response *DescribeSnapshotFilesResponse) {
    response = &DescribeSnapshotFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSnapshotFiles
// This API is used to query snapshot files.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_SNAPSHOTNOTATTACHED = "UnsupportedOperation.SnapshotNotAttached"
func (c *Client) DescribeSnapshotFiles(request *DescribeSnapshotFilesRequest) (response *DescribeSnapshotFilesResponse, err error) {
    return c.DescribeSnapshotFilesWithContext(context.Background(), request)
}

// DescribeSnapshotFiles
// This API is used to query snapshot files.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_SNAPSHOTNOTATTACHED = "UnsupportedOperation.SnapshotNotAttached"
func (c *Client) DescribeSnapshotFilesWithContext(ctx context.Context, request *DescribeSnapshotFilesRequest) (response *DescribeSnapshotFilesResponse, err error) {
    if request == nil {
        request = NewDescribeSnapshotFilesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeSnapshotFiles")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSnapshotFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSnapshotFilesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSnapshotPoliciesRequest() (request *DescribeSnapshotPoliciesRequest) {
    request = &DescribeSnapshotPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeSnapshotPolicies")
    
    
    return
}

func NewDescribeSnapshotPoliciesResponse() (response *DescribeSnapshotPoliciesResponse) {
    response = &DescribeSnapshotPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSnapshotPolicies
// This API is used to query snapshot policies.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NOTUTF8ENCODINGERROR = "InvalidParameterValue.NotUtf8EncodingError"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
func (c *Client) DescribeSnapshotPolicies(request *DescribeSnapshotPoliciesRequest) (response *DescribeSnapshotPoliciesResponse, err error) {
    return c.DescribeSnapshotPoliciesWithContext(context.Background(), request)
}

// DescribeSnapshotPolicies
// This API is used to query snapshot policies.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NOTUTF8ENCODINGERROR = "InvalidParameterValue.NotUtf8EncodingError"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
func (c *Client) DescribeSnapshotPoliciesWithContext(ctx context.Context, request *DescribeSnapshotPoliciesRequest) (response *DescribeSnapshotPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeSnapshotPoliciesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeSnapshotPolicies")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSnapshotPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSnapshotPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubnetResourceDashboardRequest() (request *DescribeSubnetResourceDashboardRequest) {
    request = &DescribeSubnetResourceDashboardRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeSubnetResourceDashboard")
    
    
    return
}

func NewDescribeSubnetResourceDashboardResponse() (response *DescribeSubnetResourceDashboardResponse) {
    response = &DescribeSubnetResourceDashboardResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSubnetResourceDashboard
// This API is used to query the subnet resource.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSubnetResourceDashboard(request *DescribeSubnetResourceDashboardRequest) (response *DescribeSubnetResourceDashboardResponse, err error) {
    return c.DescribeSubnetResourceDashboardWithContext(context.Background(), request)
}

// DescribeSubnetResourceDashboard
// This API is used to query the subnet resource.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSubnetResourceDashboardWithContext(ctx context.Context, request *DescribeSubnetResourceDashboardRequest) (response *DescribeSubnetResourceDashboardResponse, err error) {
    if request == nil {
        request = NewDescribeSubnetResourceDashboardRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeSubnetResourceDashboard")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSubnetResourceDashboard require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSubnetResourceDashboardResponse()
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeSubnets")
    
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
//  FAILEDOPERATION_TASKFAILED = "FailedOperation.TaskFailed"
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
//  FAILEDOPERATION_TASKFAILED = "FailedOperation.TaskFailed"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeTaskResultWithContext(ctx context.Context, request *DescribeTaskResultRequest) (response *DescribeTaskResultResponse, err error) {
    if request == nil {
        request = NewDescribeTaskResultRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeTaskResult")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrafficPackagesRequest() (request *DescribeTrafficPackagesRequest) {
    request = &DescribeTrafficPackagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeTrafficPackages")
    
    
    return
}

func NewDescribeTrafficPackagesResponse() (response *DescribeTrafficPackagesResponse) {
    response = &DescribeTrafficPackagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTrafficPackages
// This API is used to query the details of shared traffic packages.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETERVALUE_TRAFFICPACKAGEIDMALFORMED = "InvalidParameterValue.TrafficPackageIdMalformed"
func (c *Client) DescribeTrafficPackages(request *DescribeTrafficPackagesRequest) (response *DescribeTrafficPackagesResponse, err error) {
    return c.DescribeTrafficPackagesWithContext(context.Background(), request)
}

// DescribeTrafficPackages
// This API is used to query the details of shared traffic packages.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETERVALUE_TRAFFICPACKAGEIDMALFORMED = "InvalidParameterValue.TrafficPackageIdMalformed"
func (c *Client) DescribeTrafficPackagesWithContext(ctx context.Context, request *DescribeTrafficPackagesRequest) (response *DescribeTrafficPackagesResponse, err error) {
    if request == nil {
        request = NewDescribeTrafficPackagesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeTrafficPackages")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTrafficPackages require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTrafficPackagesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsedIpAddressRequest() (request *DescribeUsedIpAddressRequest) {
    request = &DescribeUsedIpAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeUsedIpAddress")
    
    
    return
}

func NewDescribeUsedIpAddressResponse() (response *DescribeUsedIpAddressResponse) {
    response = &DescribeUsedIpAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUsedIpAddress
// This API is used to query the IP usage of a subnet or VPC.
//
// If the IP is occupied, the resource type and ID associated with the are is returned. If the IP is not used, it returns null.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_PARAMETERMISMATCH = "InvalidParameterValue.ParameterMismatch"
func (c *Client) DescribeUsedIpAddress(request *DescribeUsedIpAddressRequest) (response *DescribeUsedIpAddressResponse, err error) {
    return c.DescribeUsedIpAddressWithContext(context.Background(), request)
}

// DescribeUsedIpAddress
// This API is used to query the IP usage of a subnet or VPC.
//
// If the IP is occupied, the resource type and ID associated with the are is returned. If the IP is not used, it returns null.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_PARAMETERMISMATCH = "InvalidParameterValue.ParameterMismatch"
func (c *Client) DescribeUsedIpAddressWithContext(ctx context.Context, request *DescribeUsedIpAddressRequest) (response *DescribeUsedIpAddressResponse, err error) {
    if request == nil {
        request = NewDescribeUsedIpAddressRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeUsedIpAddress")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUsedIpAddress require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUsedIpAddressResponse()
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
//  INVALIDPARAMETER_FILTERNOTDICT = "InvalidParameter.FilterNotDict"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NOTUTF8ENCODINGERROR = "InvalidParameterValue.NotUtf8EncodingError"
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
//  INVALIDPARAMETER_FILTERNOTDICT = "InvalidParameter.FilterNotDict"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NOTUTF8ENCODINGERROR = "InvalidParameterValue.NotUtf8EncodingError"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SVCNOTEXIST = "ResourceNotFound.SvcNotExist"
func (c *Client) DescribeVpcEndPointWithContext(ctx context.Context, request *DescribeVpcEndPointRequest) (response *DescribeVpcEndPointResponse, err error) {
    if request == nil {
        request = NewDescribeVpcEndPointRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeVpcEndPoint")
    
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
//  INVALIDPARAMETER_FILTERNOTDICT = "InvalidParameter.FilterNotDict"
//  INVALIDPARAMETER_FILTERVALUESNOTLIST = "InvalidParameter.FilterValuesNotList"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NOTUTF8ENCODINGERROR = "InvalidParameterValue.NotUtf8EncodingError"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_INSTANCEMISMATCH = "UnsupportedOperation.InstanceMismatch"
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
//  INVALIDPARAMETER_FILTERNOTDICT = "InvalidParameter.FilterNotDict"
//  INVALIDPARAMETER_FILTERVALUESNOTLIST = "InvalidParameter.FilterValuesNotList"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NOTUTF8ENCODINGERROR = "InvalidParameterValue.NotUtf8EncodingError"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_INSTANCEMISMATCH = "UnsupportedOperation.InstanceMismatch"
//  UNSUPPORTEDOPERATION_ROLENOTFOUND = "UnsupportedOperation.RoleNotFound"
func (c *Client) DescribeVpcEndPointServiceWithContext(ctx context.Context, request *DescribeVpcEndPointServiceRequest) (response *DescribeVpcEndPointServiceResponse, err error) {
    if request == nil {
        request = NewDescribeVpcEndPointServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeVpcEndPointService")
    
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
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
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
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  UNSUPPORTEDOPERATION_UINNOTFOUND = "UnsupportedOperation.UinNotFound"
func (c *Client) DescribeVpcEndPointServiceWhiteListWithContext(ctx context.Context, request *DescribeVpcEndPointServiceWhiteListRequest) (response *DescribeVpcEndPointServiceWhiteListResponse, err error) {
    if request == nil {
        request = NewDescribeVpcEndPointServiceWhiteListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeVpcEndPointServiceWhiteList")
    
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVpcInstancesWithContext(ctx context.Context, request *DescribeVpcInstancesRequest) (response *DescribeVpcInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeVpcInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeVpcInstances")
    
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
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
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
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVpcIpv6AddressesWithContext(ctx context.Context, request *DescribeVpcIpv6AddressesRequest) (response *DescribeVpcIpv6AddressesResponse, err error) {
    if request == nil {
        request = NewDescribeVpcIpv6AddressesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeVpcIpv6Addresses")
    
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeVpcPrivateIpAddressesWithContext(ctx context.Context, request *DescribeVpcPrivateIpAddressesRequest) (response *DescribeVpcPrivateIpAddressesResponse, err error) {
    if request == nil {
        request = NewDescribeVpcPrivateIpAddressesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeVpcPrivateIpAddresses")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeVpcResourceDashboard")
    
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeVpcTaskResultWithContext(ctx context.Context, request *DescribeVpcTaskResultRequest) (response *DescribeVpcTaskResultResponse, err error) {
    if request == nil {
        request = NewDescribeVpcTaskResultRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeVpcTaskResult")
    
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETER_FILTERNOTDICT = "InvalidParameter.FilterNotDict"
//  INVALIDPARAMETER_FILTERVALUESNOTLIST = "InvalidParameter.FilterValuesNotList"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NOTUTF8ENCODINGERROR = "InvalidParameterValue.NotUtf8EncodingError"
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
//  INVALIDPARAMETER_FILTERNOTDICT = "InvalidParameter.FilterNotDict"
//  INVALIDPARAMETER_FILTERVALUESNOTLIST = "InvalidParameter.FilterValuesNotList"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NOTUTF8ENCODINGERROR = "InvalidParameterValue.NotUtf8EncodingError"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeVpcsWithContext(ctx context.Context, request *DescribeVpcsRequest) (response *DescribeVpcsResponse, err error) {
    if request == nil {
        request = NewDescribeVpcsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeVpcs")
    
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
// This API is used to used to query the list of VPN tunnels.
//
// error code that may be returned:
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
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
// This API is used to used to query the list of VPN tunnels.
//
// error code that may be returned:
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeVpnConnections")
    
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
// This API is used to query VPN gateway-based CCN routes.
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
// This API is used to query VPN gateway-based CCN routes.
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeVpnGatewayCcnRoutes")
    
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
// This API is used to query VPN gateway routes.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
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
// This API is used to query VPN gateway routes.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FILTERINVALIDKEY = "InvalidParameter.FilterInvalidKey"
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeVpnGatewayRoutes")
    
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DescribeVpnGateways")
    
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
//  INVALIDPARAMETERVALUE_PARAMETERMISMATCH = "InvalidParameterValue.ParameterMismatch"
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
//  INVALIDPARAMETERVALUE_PARAMETERMISMATCH = "InvalidParameterValue.ParameterMismatch"
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DetachCcnInstances")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DetachClassicLinkVpc")
    
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED_ACTIONLIMITED = "LimitExceeded.ActionLimited"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_IDEMPOTENTPARAMETERMISMATCH = "UnsupportedOperation.IdempotentParameterMismatch"
//  UNSUPPORTEDOPERATION_IDEMPOTENTPROCESSING = "UnsupportedOperation.IdempotentProcessing"
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED_ACTIONLIMITED = "LimitExceeded.ActionLimited"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_IDEMPOTENTPARAMETERMISMATCH = "UnsupportedOperation.IdempotentParameterMismatch"
//  UNSUPPORTEDOPERATION_IDEMPOTENTPROCESSING = "UnsupportedOperation.IdempotentProcessing"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) DetachNetworkInterfaceWithContext(ctx context.Context, request *DetachNetworkInterfaceRequest) (response *DetachNetworkInterfaceResponse, err error) {
    if request == nil {
        request = NewDetachNetworkInterfaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DetachNetworkInterface")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetachNetworkInterface require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetachNetworkInterfaceResponse()
    err = c.Send(request, response)
    return
}

func NewDetachSnapshotInstancesRequest() (request *DetachSnapshotInstancesRequest) {
    request = &DetachSnapshotInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "DetachSnapshotInstances")
    
    
    return
}

func NewDetachSnapshotInstancesResponse() (response *DetachSnapshotInstancesResponse) {
    response = &DetachSnapshotInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DetachSnapshotInstances
// This API is used to disassociate a snapshot policy with instances.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATEPARA = "InvalidParameterValue.DuplicatePara"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_SNAPSHOTINSTANCEREGIONDIFF = "UnsupportedOperation.SnapshotInstanceRegionDiff"
//  UNSUPPORTEDOPERATION_SNAPSHOTNOTATTACHED = "UnsupportedOperation.SnapshotNotAttached"
func (c *Client) DetachSnapshotInstances(request *DetachSnapshotInstancesRequest) (response *DetachSnapshotInstancesResponse, err error) {
    return c.DetachSnapshotInstancesWithContext(context.Background(), request)
}

// DetachSnapshotInstances
// This API is used to disassociate a snapshot policy with instances.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATEPARA = "InvalidParameterValue.DuplicatePara"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_SNAPSHOTINSTANCEREGIONDIFF = "UnsupportedOperation.SnapshotInstanceRegionDiff"
//  UNSUPPORTEDOPERATION_SNAPSHOTNOTATTACHED = "UnsupportedOperation.SnapshotNotAttached"
func (c *Client) DetachSnapshotInstancesWithContext(ctx context.Context, request *DetachSnapshotInstancesRequest) (response *DetachSnapshotInstancesResponse, err error) {
    if request == nil {
        request = NewDetachSnapshotInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DetachSnapshotInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetachSnapshotInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetachSnapshotInstancesResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DisableCcnRoutes")
    
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
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DisableFlowLogs(request *DisableFlowLogsRequest) (response *DisableFlowLogsResponse, err error) {
    return c.DisableFlowLogsWithContext(context.Background(), request)
}

// DisableFlowLogs
// This API is used to disable flow log.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DisableFlowLogsWithContext(ctx context.Context, request *DisableFlowLogsRequest) (response *DisableFlowLogsResponse, err error) {
    if request == nil {
        request = NewDisableFlowLogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DisableFlowLogs")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DisableGatewayFlowMonitor")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableGatewayFlowMonitor require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableGatewayFlowMonitorResponse()
    err = c.Send(request, response)
    return
}

func NewDisableRoutesRequest() (request *DisableRoutesRequest) {
    request = &DisableRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "DisableRoutes")
    
    
    return
}

func NewDisableRoutesResponse() (response *DisableRoutesResponse) {
    response = &DisableRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisableRoutes
// This API is used to disable enabled subnet routes.
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_DISABLEDNOTIFYCCN = "UnsupportedOperation.DisabledNotifyCcn"
//  UNSUPPORTEDOPERATION_SYSTEMROUTE = "UnsupportedOperation.SystemRoute"
func (c *Client) DisableRoutes(request *DisableRoutesRequest) (response *DisableRoutesResponse, err error) {
    return c.DisableRoutesWithContext(context.Background(), request)
}

// DisableRoutes
// This API is used to disable enabled subnet routes.
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_DISABLEDNOTIFYCCN = "UnsupportedOperation.DisabledNotifyCcn"
//  UNSUPPORTEDOPERATION_SYSTEMROUTE = "UnsupportedOperation.SystemRoute"
func (c *Client) DisableRoutesWithContext(ctx context.Context, request *DisableRoutesRequest) (response *DisableRoutesResponse, err error) {
    if request == nil {
        request = NewDisableRoutesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DisableRoutes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableRoutes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewDisableSnapshotPoliciesRequest() (request *DisableSnapshotPoliciesRequest) {
    request = &DisableSnapshotPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "DisableSnapshotPolicies")
    
    
    return
}

func NewDisableSnapshotPoliciesResponse() (response *DisableSnapshotPoliciesResponse) {
    response = &DisableSnapshotPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisableSnapshotPolicies
// This API is used to disable specified snapshot policies.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DisableSnapshotPolicies(request *DisableSnapshotPoliciesRequest) (response *DisableSnapshotPoliciesResponse, err error) {
    return c.DisableSnapshotPoliciesWithContext(context.Background(), request)
}

// DisableSnapshotPolicies
// This API is used to disable specified snapshot policies.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DisableSnapshotPoliciesWithContext(ctx context.Context, request *DisableSnapshotPoliciesRequest) (response *DisableSnapshotPoliciesResponse, err error) {
    if request == nil {
        request = NewDisableSnapshotPoliciesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DisableSnapshotPolicies")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableSnapshotPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableSnapshotPoliciesResponse()
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
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTSUPPORTED = "UnsupportedOperation.InstanceStateNotSupported"
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
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTSUPPORTED = "UnsupportedOperation.InstanceStateNotSupported"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DisassociateAddressWithContext(ctx context.Context, request *DisassociateAddressRequest) (response *DisassociateAddressResponse, err error) {
    if request == nil {
        request = NewDisassociateAddressRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DisassociateAddress")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DisassociateDirectConnectGatewayNatGateway")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisassociateDirectConnectGatewayNatGateway require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisassociateDirectConnectGatewayNatGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateIPv6AddressRequest() (request *DisassociateIPv6AddressRequest) {
    request = &DisassociateIPv6AddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "DisassociateIPv6Address")
    
    
    return
}

func NewDisassociateIPv6AddressResponse() (response *DisassociateIPv6AddressResponse) {
    response = &DisassociateIPv6AddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisassociateIPv6Address
// This API is used to unbind an EIPv6 instance.
//
// 
//
// - You can unbind EIPv6 instances bound to Cloud Virtual Machine (CVM) or Elastic Network Interface (ENI).
//
// - Only EIPv6 instances in BIND or BIND_ENI status can be unbound.
//
// error code that may be returned:
//  FAILEDOPERATION_ADDRESSENIINFONOTFOUND = "FailedOperation.AddressEniInfoNotFound"
//  INVALIDPARAMETERVALUE_ADDRESSIDMALFORMED = "InvalidParameterValue.AddressIdMalformed"
//  INVALIDPARAMETERVALUE_ADDRESSIPNOTFOUND = "InvalidParameterValue.AddressIpNotFound"
//  OPERATIONDENIED_ADDRESSINARREARS = "OperationDenied.AddressInArrears"
//  UNSUPPORTEDOPERATION_ADDRESSSTATUSNOTPERMIT = "UnsupportedOperation.AddressStatusNotPermit"
func (c *Client) DisassociateIPv6Address(request *DisassociateIPv6AddressRequest) (response *DisassociateIPv6AddressResponse, err error) {
    return c.DisassociateIPv6AddressWithContext(context.Background(), request)
}

// DisassociateIPv6Address
// This API is used to unbind an EIPv6 instance.
//
// 
//
// - You can unbind EIPv6 instances bound to Cloud Virtual Machine (CVM) or Elastic Network Interface (ENI).
//
// - Only EIPv6 instances in BIND or BIND_ENI status can be unbound.
//
// error code that may be returned:
//  FAILEDOPERATION_ADDRESSENIINFONOTFOUND = "FailedOperation.AddressEniInfoNotFound"
//  INVALIDPARAMETERVALUE_ADDRESSIDMALFORMED = "InvalidParameterValue.AddressIdMalformed"
//  INVALIDPARAMETERVALUE_ADDRESSIPNOTFOUND = "InvalidParameterValue.AddressIpNotFound"
//  OPERATIONDENIED_ADDRESSINARREARS = "OperationDenied.AddressInArrears"
//  UNSUPPORTEDOPERATION_ADDRESSSTATUSNOTPERMIT = "UnsupportedOperation.AddressStatusNotPermit"
func (c *Client) DisassociateIPv6AddressWithContext(ctx context.Context, request *DisassociateIPv6AddressRequest) (response *DisassociateIPv6AddressResponse, err error) {
    if request == nil {
        request = NewDisassociateIPv6AddressRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DisassociateIPv6Address")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisassociateIPv6Address require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisassociateIPv6AddressResponse()
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
//  UNSUPPORTEDOPERATION_NATGATEWAYRESTRICTED = "UnsupportedOperation.NatGatewayRestricted"
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
//  UNSUPPORTEDOPERATION_NATGATEWAYRESTRICTED = "UnsupportedOperation.NatGatewayRestricted"
//  UNSUPPORTEDOPERATION_PUBLICIPADDRESSDISASSOCIATE = "UnsupportedOperation.PublicIpAddressDisassociate"
func (c *Client) DisassociateNatGatewayAddressWithContext(ctx context.Context, request *DisassociateNatGatewayAddressRequest) (response *DisassociateNatGatewayAddressResponse, err error) {
    if request == nil {
        request = NewDisassociateNatGatewayAddressRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DisassociateNatGatewayAddress")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DisassociateNetworkAclSubnets")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DisassociateNetworkInterfaceSecurityGroups")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DisassociateVpcEndPointSecurityGroups")
    
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
// This API is used to download VPN tunnel configurations.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DownloadCustomerGatewayConfiguration(request *DownloadCustomerGatewayConfigurationRequest) (response *DownloadCustomerGatewayConfigurationResponse, err error) {
    return c.DownloadCustomerGatewayConfigurationWithContext(context.Background(), request)
}

// DownloadCustomerGatewayConfiguration
// This API is used to download VPN tunnel configurations.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DownloadCustomerGatewayConfigurationWithContext(ctx context.Context, request *DownloadCustomerGatewayConfigurationRequest) (response *DownloadCustomerGatewayConfigurationResponse, err error) {
    if request == nil {
        request = NewDownloadCustomerGatewayConfigurationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "DownloadCustomerGatewayConfiguration")
    
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
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
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
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_ECMP = "UnsupportedOperation.Ecmp"
func (c *Client) EnableCcnRoutesWithContext(ctx context.Context, request *EnableCcnRoutesRequest) (response *EnableCcnRoutesResponse, err error) {
    if request == nil {
        request = NewEnableCcnRoutesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "EnableCcnRoutes")
    
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
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) EnableFlowLogs(request *EnableFlowLogsRequest) (response *EnableFlowLogsResponse, err error) {
    return c.EnableFlowLogsWithContext(context.Background(), request)
}

// EnableFlowLogs
// This API is used to enable flow log.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) EnableFlowLogsWithContext(ctx context.Context, request *EnableFlowLogsRequest) (response *EnableFlowLogsResponse, err error) {
    if request == nil {
        request = NewEnableFlowLogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "EnableFlowLogs")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "EnableGatewayFlowMonitor")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableGatewayFlowMonitor require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableGatewayFlowMonitorResponse()
    err = c.Send(request, response)
    return
}

func NewEnableRoutesRequest() (request *EnableRoutesRequest) {
    request = &EnableRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "EnableRoutes")
    
    
    return
}

func NewEnableRoutesResponse() (response *EnableRoutesResponse) {
    response = &EnableRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableRoutes
// This API is used to enable disabled subnet routes.<br />
//
// The API is used to verify whether the enabled route conflicts with existing routes. If they conflict, the new route cannot be enabled and will result in a failure. When a route conflict occurs, you need to first disable the conflicting route before you can enable the new one.
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CONFLICTWITHDOCKERROUTE = "UnsupportedOperation.ConflictWithDockerRoute"
//  UNSUPPORTEDOPERATION_ECMP = "UnsupportedOperation.Ecmp"
//  UNSUPPORTEDOPERATION_ECMPWITHCCNROUTE = "UnsupportedOperation.EcmpWithCcnRoute"
//  UNSUPPORTEDOPERATION_ECMPWITHUSERROUTE = "UnsupportedOperation.EcmpWithUserRoute"
//  UNSUPPORTEDOPERATION_SYSTEMROUTE = "UnsupportedOperation.SystemRoute"
func (c *Client) EnableRoutes(request *EnableRoutesRequest) (response *EnableRoutesResponse, err error) {
    return c.EnableRoutesWithContext(context.Background(), request)
}

// EnableRoutes
// This API is used to enable disabled subnet routes.<br />
//
// The API is used to verify whether the enabled route conflicts with existing routes. If they conflict, the new route cannot be enabled and will result in a failure. When a route conflict occurs, you need to first disable the conflicting route before you can enable the new one.
//
// error code that may be returned:
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CONFLICTWITHDOCKERROUTE = "UnsupportedOperation.ConflictWithDockerRoute"
//  UNSUPPORTEDOPERATION_ECMP = "UnsupportedOperation.Ecmp"
//  UNSUPPORTEDOPERATION_ECMPWITHCCNROUTE = "UnsupportedOperation.EcmpWithCcnRoute"
//  UNSUPPORTEDOPERATION_ECMPWITHUSERROUTE = "UnsupportedOperation.EcmpWithUserRoute"
//  UNSUPPORTEDOPERATION_SYSTEMROUTE = "UnsupportedOperation.SystemRoute"
func (c *Client) EnableRoutesWithContext(ctx context.Context, request *EnableRoutesRequest) (response *EnableRoutesResponse, err error) {
    if request == nil {
        request = NewEnableRoutesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "EnableRoutes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableRoutes require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewEnableSnapshotPoliciesRequest() (request *EnableSnapshotPoliciesRequest) {
    request = &EnableSnapshotPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "EnableSnapshotPolicies")
    
    
    return
}

func NewEnableSnapshotPoliciesResponse() (response *EnableSnapshotPoliciesResponse) {
    response = &EnableSnapshotPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableSnapshotPolicies
// This API is used to enable specified snapshot policies. 
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) EnableSnapshotPolicies(request *EnableSnapshotPoliciesRequest) (response *EnableSnapshotPoliciesResponse, err error) {
    return c.EnableSnapshotPoliciesWithContext(context.Background(), request)
}

// EnableSnapshotPolicies
// This API is used to enable specified snapshot policies. 
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) EnableSnapshotPoliciesWithContext(ctx context.Context, request *EnableSnapshotPoliciesRequest) (response *EnableSnapshotPoliciesResponse, err error) {
    if request == nil {
        request = NewEnableSnapshotPoliciesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "EnableSnapshotPolicies")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableSnapshotPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableSnapshotPoliciesResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "EnableVpcEndPointConnect")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableVpcEndPointConnect require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableVpcEndPointConnectResponse()
    err = c.Send(request, response)
    return
}

func NewGenerateVpnConnectionDefaultHealthCheckIpRequest() (request *GenerateVpnConnectionDefaultHealthCheckIpRequest) {
    request = &GenerateVpnConnectionDefaultHealthCheckIpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "GenerateVpnConnectionDefaultHealthCheckIp")
    
    
    return
}

func NewGenerateVpnConnectionDefaultHealthCheckIpResponse() (response *GenerateVpnConnectionDefaultHealthCheckIpResponse) {
    response = &GenerateVpnConnectionDefaultHealthCheckIpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GenerateVpnConnectionDefaultHealthCheckIp
// This API is used to get a pair of VPN tunnel health check addresses. 
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GenerateVpnConnectionDefaultHealthCheckIp(request *GenerateVpnConnectionDefaultHealthCheckIpRequest) (response *GenerateVpnConnectionDefaultHealthCheckIpResponse, err error) {
    return c.GenerateVpnConnectionDefaultHealthCheckIpWithContext(context.Background(), request)
}

// GenerateVpnConnectionDefaultHealthCheckIp
// This API is used to get a pair of VPN tunnel health check addresses. 
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GenerateVpnConnectionDefaultHealthCheckIpWithContext(ctx context.Context, request *GenerateVpnConnectionDefaultHealthCheckIpRequest) (response *GenerateVpnConnectionDefaultHealthCheckIpResponse, err error) {
    if request == nil {
        request = NewGenerateVpnConnectionDefaultHealthCheckIpRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "GenerateVpnConnectionDefaultHealthCheckIp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GenerateVpnConnectionDefaultHealthCheckIp require credential")
    }

    request.SetContext(ctx)
    
    response = NewGenerateVpnConnectionDefaultHealthCheckIpResponse()
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
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
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
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetCcnRegionBandwidthLimitsWithContext(ctx context.Context, request *GetCcnRegionBandwidthLimitsRequest) (response *GetCcnRegionBandwidthLimitsResponse, err error) {
    if request == nil {
        request = NewGetCcnRegionBandwidthLimitsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "GetCcnRegionBandwidthLimits")
    
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BINDEIP = "UnsupportedOperation.BindEIP"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDBINDLOCALZONEEIP = "UnsupportedOperation.UnsupportedBindLocalZoneEIP"
func (c *Client) HaVipAssociateAddressIp(request *HaVipAssociateAddressIpRequest) (response *HaVipAssociateAddressIpResponse, err error) {
    return c.HaVipAssociateAddressIpWithContext(context.Background(), request)
}

// HaVipAssociateAddressIp
// This API is used to bind an EIP to an HAVIP. <br />
//
// This API is completed asynchronously. If you need to query the execution result of an async task, please use the `RequestId` returned by this API to poll the `DescribeVpcTaskResult` API.
//
// error code that may be returned:
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BINDEIP = "UnsupportedOperation.BindEIP"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDBINDLOCALZONEEIP = "UnsupportedOperation.UnsupportedBindLocalZoneEIP"
func (c *Client) HaVipAssociateAddressIpWithContext(ctx context.Context, request *HaVipAssociateAddressIpRequest) (response *HaVipAssociateAddressIpResponse, err error) {
    if request == nil {
        request = NewHaVipAssociateAddressIpRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "HaVipAssociateAddressIp")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "HaVipDisassociateAddressIp")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "InquirePriceCreateDirectConnectGateway")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquirePriceCreateDirectConnectGateway require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquirePriceCreateDirectConnectGatewayResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceAllocateAddressesRequest() (request *InquiryPriceAllocateAddressesRequest) {
    request = &InquiryPriceAllocateAddressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "InquiryPriceAllocateAddresses")
    
    
    return
}

func NewInquiryPriceAllocateAddressesResponse() (response *InquiryPriceAllocateAddressesResponse) {
    response = &InquiryPriceAllocateAddressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquiryPriceAllocateAddresses
// This API (InquiryPriceAllocateAddresses) is used to query the price of purchasing EIPs.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDREGION = "FailedOperation.InvalidRegion"
//  INVALIDACCOUNT_NOTSUPPORTED = "InvalidAccount.NotSupported"
//  INVALIDPARAMETERCONFLICT = "InvalidParameterConflict"
//  INVALIDPARAMETERVALUE_BANDWIDTHOUTOFRANGE = "InvalidParameterValue.BandwidthOutOfRange"
//  INVALIDPARAMETERVALUE_COMBINATION = "InvalidParameterValue.Combination"
//  INVALIDPARAMETERVALUE_INVALIDEGRESS = "InvalidParameterValue.InvalidEgress"
//  INVALIDPARAMETERVALUE_UNAVAILABLEZONE = "InvalidParameterValue.UnavailableZone"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDEGRESS = "InvalidParameterValue.UnsupportedEgress"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_ACTIONNOTFOUND = "UnsupportedOperation.ActionNotFound"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDREGION = "UnsupportedOperation.UnsupportedRegion"
func (c *Client) InquiryPriceAllocateAddresses(request *InquiryPriceAllocateAddressesRequest) (response *InquiryPriceAllocateAddressesResponse, err error) {
    return c.InquiryPriceAllocateAddressesWithContext(context.Background(), request)
}

// InquiryPriceAllocateAddresses
// This API (InquiryPriceAllocateAddresses) is used to query the price of purchasing EIPs.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDREGION = "FailedOperation.InvalidRegion"
//  INVALIDACCOUNT_NOTSUPPORTED = "InvalidAccount.NotSupported"
//  INVALIDPARAMETERCONFLICT = "InvalidParameterConflict"
//  INVALIDPARAMETERVALUE_BANDWIDTHOUTOFRANGE = "InvalidParameterValue.BandwidthOutOfRange"
//  INVALIDPARAMETERVALUE_COMBINATION = "InvalidParameterValue.Combination"
//  INVALIDPARAMETERVALUE_INVALIDEGRESS = "InvalidParameterValue.InvalidEgress"
//  INVALIDPARAMETERVALUE_UNAVAILABLEZONE = "InvalidParameterValue.UnavailableZone"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDEGRESS = "InvalidParameterValue.UnsupportedEgress"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_ACTIONNOTFOUND = "UnsupportedOperation.ActionNotFound"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDREGION = "UnsupportedOperation.UnsupportedRegion"
func (c *Client) InquiryPriceAllocateAddressesWithContext(ctx context.Context, request *InquiryPriceAllocateAddressesRequest) (response *InquiryPriceAllocateAddressesResponse, err error) {
    if request == nil {
        request = NewInquiryPriceAllocateAddressesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "InquiryPriceAllocateAddresses")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceAllocateAddresses require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceAllocateAddressesResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceModifyAddressesBandwidthRequest() (request *InquiryPriceModifyAddressesBandwidthRequest) {
    request = &InquiryPriceModifyAddressesBandwidthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "InquiryPriceModifyAddressesBandwidth")
    
    
    return
}

func NewInquiryPriceModifyAddressesBandwidthResponse() (response *InquiryPriceModifyAddressesBandwidthResponse) {
    response = &InquiryPriceModifyAddressesBandwidthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquiryPriceModifyAddressesBandwidth
// This API is used to query the price of modifying EIP bandwidth.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_ADDRESSNOTFOUND = "InvalidParameterValue.AddressNotFound"
//  INVALIDPARAMETERVALUE_BANDWIDTHOUTOFRANGE = "InvalidParameterValue.BandwidthOutOfRange"
//  INVALIDPARAMETERVALUE_BANDWIDTHTOOSMALL = "InvalidParameterValue.BandwidthTooSmall"
//  INVALIDPARAMETERVALUE_RESOURCEEXPIRED = "InvalidParameterValue.ResourceExpired"
//  INVALIDPARAMETERVALUE_RESOURCENOTSUPPORT = "InvalidParameterValue.ResourceNotSupport"
func (c *Client) InquiryPriceModifyAddressesBandwidth(request *InquiryPriceModifyAddressesBandwidthRequest) (response *InquiryPriceModifyAddressesBandwidthResponse, err error) {
    return c.InquiryPriceModifyAddressesBandwidthWithContext(context.Background(), request)
}

// InquiryPriceModifyAddressesBandwidth
// This API is used to query the price of modifying EIP bandwidth.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_ADDRESSNOTFOUND = "InvalidParameterValue.AddressNotFound"
//  INVALIDPARAMETERVALUE_BANDWIDTHOUTOFRANGE = "InvalidParameterValue.BandwidthOutOfRange"
//  INVALIDPARAMETERVALUE_BANDWIDTHTOOSMALL = "InvalidParameterValue.BandwidthTooSmall"
//  INVALIDPARAMETERVALUE_RESOURCEEXPIRED = "InvalidParameterValue.ResourceExpired"
//  INVALIDPARAMETERVALUE_RESOURCENOTSUPPORT = "InvalidParameterValue.ResourceNotSupport"
func (c *Client) InquiryPriceModifyAddressesBandwidthWithContext(ctx context.Context, request *InquiryPriceModifyAddressesBandwidthRequest) (response *InquiryPriceModifyAddressesBandwidthResponse, err error) {
    if request == nil {
        request = NewInquiryPriceModifyAddressesBandwidthRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "InquiryPriceModifyAddressesBandwidth")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceModifyAddressesBandwidth require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceModifyAddressesBandwidthResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceRenewAddressesRequest() (request *InquiryPriceRenewAddressesRequest) {
    request = &InquiryPriceRenewAddressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "InquiryPriceRenewAddresses")
    
    
    return
}

func NewInquiryPriceRenewAddressesResponse() (response *InquiryPriceRenewAddressesResponse) {
    response = &InquiryPriceRenewAddressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquiryPriceRenewAddresses
// This API (InquiryPriceRenewAddresses) is used to query the price of renewing prepaid EIPs.
//
// error code that may be returned:
//  INVALIDADDRESSID_NOTFOUND = "InvalidAddressId.NotFound"
//  INVALIDPARAMETERVALUE_ADDRESSNOTFOUND = "InvalidParameterValue.AddressNotFound"
//  UNSUPPORTEDOPERATION_INVALIDADDRESSINTERNETCHARGETYPE = "UnsupportedOperation.InvalidAddressInternetChargeType"
func (c *Client) InquiryPriceRenewAddresses(request *InquiryPriceRenewAddressesRequest) (response *InquiryPriceRenewAddressesResponse, err error) {
    return c.InquiryPriceRenewAddressesWithContext(context.Background(), request)
}

// InquiryPriceRenewAddresses
// This API (InquiryPriceRenewAddresses) is used to query the price of renewing prepaid EIPs.
//
// error code that may be returned:
//  INVALIDADDRESSID_NOTFOUND = "InvalidAddressId.NotFound"
//  INVALIDPARAMETERVALUE_ADDRESSNOTFOUND = "InvalidParameterValue.AddressNotFound"
//  UNSUPPORTEDOPERATION_INVALIDADDRESSINTERNETCHARGETYPE = "UnsupportedOperation.InvalidAddressInternetChargeType"
func (c *Client) InquiryPriceRenewAddressesWithContext(ctx context.Context, request *InquiryPriceRenewAddressesRequest) (response *InquiryPriceRenewAddressesResponse, err error) {
    if request == nil {
        request = NewInquiryPriceRenewAddressesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "InquiryPriceRenewAddresses")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceRenewAddresses require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceRenewAddressesResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "InquiryPriceRenewVpnGateway")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "InquiryPriceResetVpnGatewayInternetMaxBandwidth")
    
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ATTACHMENTNOTFOUND = "UnsupportedOperation.AttachmentNotFound"
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ATTACHMENTNOTFOUND = "UnsupportedOperation.AttachmentNotFound"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
func (c *Client) MigrateNetworkInterfaceWithContext(ctx context.Context, request *MigrateNetworkInterfaceRequest) (response *MigrateNetworkInterfaceResponse, err error) {
    if request == nil {
        request = NewMigrateNetworkInterfaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "MigrateNetworkInterface")
    
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
// This API is used to migrate the private IPs between ENIs. 
//
// * Note that primary IPs cannot be migrated. 
//
// * The source and destination ENI must be within the same subnet.  
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
// This API is used to migrate the private IPs between ENIs. 
//
// * Note that primary IPs cannot be migrated. 
//
// * The source and destination ENI must be within the same subnet.  
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "MigratePrivateIpAddress")
    
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
//  INVALIDPARAMETERVALUE_RESOURCENOTEXISTED = "InvalidParameterValue.ResourceNotExisted"
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
//  INVALIDPARAMETERVALUE_RESOURCENOTEXISTED = "InvalidParameterValue.ResourceNotExisted"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  OPERATIONDENIED_MUTEXTASKRUNNING = "OperationDenied.MutexTaskRunning"
//  UNSUPPORTEDOPERATION_ADDRESSSTATUSNOTPERMIT = "UnsupportedOperation.AddressStatusNotPermit"
//  UNSUPPORTEDOPERATION_INCORRECTADDRESSRESOURCETYPE = "UnsupportedOperation.IncorrectAddressResourceType"
//  UNSUPPORTEDOPERATION_MODIFYADDRESSATTRIBUTE = "UnsupportedOperation.ModifyAddressAttribute"
func (c *Client) ModifyAddressAttributeWithContext(ctx context.Context, request *ModifyAddressAttributeRequest) (response *ModifyAddressAttributeResponse, err error) {
    if request == nil {
        request = NewModifyAddressAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyAddressAttribute")
    
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
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
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
//  OPERATIONDENIED_ADDRESSINARREARS = "OperationDenied.AddressInArrears"
//  OPERATIONDENIED_MUTEXTASKRUNNING = "OperationDenied.MutexTaskRunning"
//  UNSUPPORTEDOPERATION_ADDRESSSTATUSNOTPERMIT = "UnsupportedOperation.AddressStatusNotPermit"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_INVALIDADDRESSINTERNETCHARGETYPE = "UnsupportedOperation.InvalidAddressInternetChargeType"
//  UNSUPPORTEDOPERATION_NATNOTSUPPORTED = "UnsupportedOperation.NatNotSupported"
//  UNSUPPORTEDOPERATION_UNPAIDORDERALREADYEXISTS = "UnsupportedOperation.UnpaidOrderAlreadyExists"
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
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
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
//  OPERATIONDENIED_ADDRESSINARREARS = "OperationDenied.AddressInArrears"
//  OPERATIONDENIED_MUTEXTASKRUNNING = "OperationDenied.MutexTaskRunning"
//  UNSUPPORTEDOPERATION_ADDRESSSTATUSNOTPERMIT = "UnsupportedOperation.AddressStatusNotPermit"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_INVALIDADDRESSINTERNETCHARGETYPE = "UnsupportedOperation.InvalidAddressInternetChargeType"
//  UNSUPPORTEDOPERATION_NATNOTSUPPORTED = "UnsupportedOperation.NatNotSupported"
//  UNSUPPORTEDOPERATION_UNPAIDORDERALREADYEXISTS = "UnsupportedOperation.UnpaidOrderAlreadyExists"
func (c *Client) ModifyAddressInternetChargeTypeWithContext(ctx context.Context, request *ModifyAddressInternetChargeTypeRequest) (response *ModifyAddressInternetChargeTypeResponse, err error) {
    if request == nil {
        request = NewModifyAddressInternetChargeTypeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyAddressInternetChargeType")
    
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
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
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
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyAddressTemplateAttribute")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyAddressTemplateGroupAttribute")
    
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
//  FAILEDOPERATION_TASKFAILED = "FailedOperation.TaskFailed"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDADDRESSID_NOTFOUND = "InvalidAddressId.NotFound"
//  INVALIDPARAMETERVALUE_ADDRESSIDMALFORMED = "InvalidParameterValue.AddressIdMalformed"
//  INVALIDPARAMETERVALUE_ADDRESSNOTFOUND = "InvalidParameterValue.AddressNotFound"
//  INVALIDPARAMETERVALUE_ADDRESSTYPECONFLICT = "InvalidParameterValue.AddressTypeConflict"
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
//  OPERATIONDENIED_ADDRESSINARREARS = "OperationDenied.AddressInArrears"
//  OPERATIONDENIED_MUTEXTASKRUNNING = "OperationDenied.MutexTaskRunning"
//  UNSUPPORTEDOPERATION_ACTIONNOTFOUND = "UnsupportedOperation.ActionNotFound"
//  UNSUPPORTEDOPERATION_ADDRESSSTATUSNOTPERMIT = "UnsupportedOperation.AddressStatusNotPermit"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTSUPPORTED = "UnsupportedOperation.InstanceStateNotSupported"
//  UNSUPPORTEDOPERATION_UNPAIDORDERALREADYEXISTS = "UnsupportedOperation.UnpaidOrderAlreadyExists"
func (c *Client) ModifyAddressesBandwidth(request *ModifyAddressesBandwidthRequest) (response *ModifyAddressesBandwidthResponse, err error) {
    return c.ModifyAddressesBandwidthWithContext(context.Background(), request)
}

// ModifyAddressesBandwidth
// This API is used to adjust the bandwidth of [Elastic IP](https://intl.cloud.tencent.com/document/product/213/1941?from_cn_redirect=1), including EIP billed on a pay-as-you-go, monthly subscription, and bandwidth package basis.
//
// error code that may be returned:
//  FAILEDOPERATION_TASKFAILED = "FailedOperation.TaskFailed"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDADDRESSID_NOTFOUND = "InvalidAddressId.NotFound"
//  INVALIDPARAMETERVALUE_ADDRESSIDMALFORMED = "InvalidParameterValue.AddressIdMalformed"
//  INVALIDPARAMETERVALUE_ADDRESSNOTFOUND = "InvalidParameterValue.AddressNotFound"
//  INVALIDPARAMETERVALUE_ADDRESSTYPECONFLICT = "InvalidParameterValue.AddressTypeConflict"
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
//  OPERATIONDENIED_ADDRESSINARREARS = "OperationDenied.AddressInArrears"
//  OPERATIONDENIED_MUTEXTASKRUNNING = "OperationDenied.MutexTaskRunning"
//  UNSUPPORTEDOPERATION_ACTIONNOTFOUND = "UnsupportedOperation.ActionNotFound"
//  UNSUPPORTEDOPERATION_ADDRESSSTATUSNOTPERMIT = "UnsupportedOperation.AddressStatusNotPermit"
//  UNSUPPORTEDOPERATION_INSTANCESTATENOTSUPPORTED = "UnsupportedOperation.InstanceStateNotSupported"
//  UNSUPPORTEDOPERATION_UNPAIDORDERALREADYEXISTS = "UnsupportedOperation.UnpaidOrderAlreadyExists"
func (c *Client) ModifyAddressesBandwidthWithContext(ctx context.Context, request *ModifyAddressesBandwidthRequest) (response *ModifyAddressesBandwidthResponse, err error) {
    if request == nil {
        request = NewModifyAddressesBandwidthRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyAddressesBandwidth")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAddressesBandwidth require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAddressesBandwidthResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAddressesRenewFlagRequest() (request *ModifyAddressesRenewFlagRequest) {
    request = &ModifyAddressesRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyAddressesRenewFlag")
    
    
    return
}

func NewModifyAddressesRenewFlagResponse() (response *ModifyAddressesRenewFlagResponse) {
    response = &ModifyAddressesRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAddressesRenewFlag
// This API is used to adjust the renewal flag for the monthly subscription EIP.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ADDRESSIDMALFORMED = "InvalidParameterValue.AddressIdMalformed"
//  INVALIDPARAMETERVALUE_ADDRESSNOTFOUND = "InvalidParameterValue.AddressNotFound"
//  INVALIDPARAMETERVALUE_RESOURCENOTSUPPORT = "InvalidParameterValue.ResourceNotSupport"
//  UNSUPPORTEDOPERATION_ADDRESSSTATUSNOTPERMIT = "UnsupportedOperation.AddressStatusNotPermit"
func (c *Client) ModifyAddressesRenewFlag(request *ModifyAddressesRenewFlagRequest) (response *ModifyAddressesRenewFlagResponse, err error) {
    return c.ModifyAddressesRenewFlagWithContext(context.Background(), request)
}

// ModifyAddressesRenewFlag
// This API is used to adjust the renewal flag for the monthly subscription EIP.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ADDRESSIDMALFORMED = "InvalidParameterValue.AddressIdMalformed"
//  INVALIDPARAMETERVALUE_ADDRESSNOTFOUND = "InvalidParameterValue.AddressNotFound"
//  INVALIDPARAMETERVALUE_RESOURCENOTSUPPORT = "InvalidParameterValue.ResourceNotSupport"
//  UNSUPPORTEDOPERATION_ADDRESSSTATUSNOTPERMIT = "UnsupportedOperation.AddressStatusNotPermit"
func (c *Client) ModifyAddressesRenewFlagWithContext(ctx context.Context, request *ModifyAddressesRenewFlagRequest) (response *ModifyAddressesRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifyAddressesRenewFlagRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyAddressesRenewFlag")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAddressesRenewFlag require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAddressesRenewFlagResponse()
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
// This API is used to batch modify (add or delete) secondary CIDR blocks.
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
// This API is used to batch modify (add or delete) secondary CIDR blocks.
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyAssistantCidr")
    
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
//  UNSUPPORTEDOPERATION_BANDWIDTHPACKAGEIDNOTSUPPORTED = "UnsupportedOperation.BandwidthPackageIdNotSupported"
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
//  UNSUPPORTEDOPERATION_BANDWIDTHPACKAGEIDNOTSUPPORTED = "UnsupportedOperation.BandwidthPackageIdNotSupported"
func (c *Client) ModifyBandwidthPackageAttributeWithContext(ctx context.Context, request *ModifyBandwidthPackageAttributeRequest) (response *ModifyBandwidthPackageAttributeResponse, err error) {
    if request == nil {
        request = NewModifyBandwidthPackageAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyBandwidthPackageAttribute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBandwidthPackageAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBandwidthPackageAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBandwidthPackageBandwidthRequest() (request *ModifyBandwidthPackageBandwidthRequest) {
    request = &ModifyBandwidthPackageBandwidthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyBandwidthPackageBandwidth")
    
    
    return
}

func NewModifyBandwidthPackageBandwidthResponse() (response *ModifyBandwidthPackageBandwidthResponse) {
    response = &ModifyBandwidthPackageBandwidthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyBandwidthPackageBandwidth
// This API is used to adjust the bandwidth of a [bandwidth package](https://www.tencentcloud.com/document/product/684/15245).
//
// error code that may be returned:
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGEIDMALFORMED = "InvalidParameterValue.BandwidthPackageIdMalformed"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGENOTFOUND = "InvalidParameterValue.BandwidthPackageNotFound"
//  INVALIDPARAMETERVALUE_COMBINATION = "InvalidParameterValue.Combination"
//  INVALIDPARAMETERVALUE_INVALIDOLDBANDWIDTH = "InvalidParameterValue.InvalidOldBandwidth"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  UNSUPPORTEDOPERATION_BANDWIDTHPACKAGEIDNOTSUPPORTED = "UnsupportedOperation.BandwidthPackageIdNotSupported"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) ModifyBandwidthPackageBandwidth(request *ModifyBandwidthPackageBandwidthRequest) (response *ModifyBandwidthPackageBandwidthResponse, err error) {
    return c.ModifyBandwidthPackageBandwidthWithContext(context.Background(), request)
}

// ModifyBandwidthPackageBandwidth
// This API is used to adjust the bandwidth of a [bandwidth package](https://www.tencentcloud.com/document/product/684/15245).
//
// error code that may be returned:
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGEIDMALFORMED = "InvalidParameterValue.BandwidthPackageIdMalformed"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGENOTFOUND = "InvalidParameterValue.BandwidthPackageNotFound"
//  INVALIDPARAMETERVALUE_COMBINATION = "InvalidParameterValue.Combination"
//  INVALIDPARAMETERVALUE_INVALIDOLDBANDWIDTH = "InvalidParameterValue.InvalidOldBandwidth"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  UNSUPPORTEDOPERATION_BANDWIDTHPACKAGEIDNOTSUPPORTED = "UnsupportedOperation.BandwidthPackageIdNotSupported"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) ModifyBandwidthPackageBandwidthWithContext(ctx context.Context, request *ModifyBandwidthPackageBandwidthRequest) (response *ModifyBandwidthPackageBandwidthResponse, err error) {
    if request == nil {
        request = NewModifyBandwidthPackageBandwidthRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyBandwidthPackageBandwidth")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBandwidthPackageBandwidth require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBandwidthPackageBandwidthResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyCcnAttachedInstancesAttribute")
    
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
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyCcnAttribute(request *ModifyCcnAttributeRequest) (response *ModifyCcnAttributeResponse, err error) {
    return c.ModifyCcnAttributeWithContext(context.Background(), request)
}

// ModifyCcnAttribute
// This API (ModifyCcnAttribute) is used to modify CCN attributes.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyCcnAttributeWithContext(ctx context.Context, request *ModifyCcnAttributeRequest) (response *ModifyCcnAttributeResponse, err error) {
    if request == nil {
        request = NewModifyCcnAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyCcnAttribute")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyCcnRegionBandwidthLimitsType")
    
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
//  UNSUPPORTEDOPERATION_VPNUNSUPPORTEDMODIFYBGPASN = "UnsupportedOperation.VpnUnsupportedModifyBgpAsn"
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
//  UNSUPPORTEDOPERATION_VPNUNSUPPORTEDMODIFYBGPASN = "UnsupportedOperation.VpnUnsupportedModifyBgpAsn"
func (c *Client) ModifyCustomerGatewayAttributeWithContext(ctx context.Context, request *ModifyCustomerGatewayAttributeRequest) (response *ModifyCustomerGatewayAttributeResponse, err error) {
    if request == nil {
        request = NewModifyCustomerGatewayAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyCustomerGatewayAttribute")
    
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
//  UNSUPPORTEDOPERATION_APPIDMISMATCH = "UnsupportedOperation.AppIdMismatch"
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
//  UNSUPPORTEDOPERATION_APPIDMISMATCH = "UnsupportedOperation.AppIdMismatch"
//  UNSUPPORTEDOPERATION_DIRECTCONNECTGATEWAYISUPDATINGCOMMUNITY = "UnsupportedOperation.DirectConnectGatewayIsUpdatingCommunity"
func (c *Client) ModifyDirectConnectGatewayAttributeWithContext(ctx context.Context, request *ModifyDirectConnectGatewayAttributeRequest) (response *ModifyDirectConnectGatewayAttributeResponse, err error) {
    if request == nil {
        request = NewModifyDirectConnectGatewayAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyDirectConnectGatewayAttribute")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyFlowLogAttribute")
    
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
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
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
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) ModifyGatewayFlowQosWithContext(ctx context.Context, request *ModifyGatewayFlowQosRequest) (response *ModifyGatewayFlowQosResponse, err error) {
    if request == nil {
        request = NewModifyGatewayFlowQosRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyGatewayFlowQos")
    
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyHaVipAttributeWithContext(ctx context.Context, request *ModifyHaVipAttributeRequest) (response *ModifyHaVipAttributeResponse, err error) {
    if request == nil {
        request = NewModifyHaVipAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyHaVipAttribute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyHaVipAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyHaVipAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyIPv6AddressesAttributesRequest() (request *ModifyIPv6AddressesAttributesRequest) {
    request = &ModifyIPv6AddressesAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyIPv6AddressesAttributes")
    
    
    return
}

func NewModifyIPv6AddressesAttributesResponse() (response *ModifyIPv6AddressesAttributesResponse) {
    response = &ModifyIPv6AddressesAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyIPv6AddressesAttributes
// This API is used to modify the name of an EIPv6 instance.
//
// 
//
// - You can modify the name of both EIPv6 and traditional EIPv6 instances.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ADDRESSIDMALFORMED = "InvalidParameterValue.AddressIdMalformed"
//  INVALIDPARAMETERVALUE_ADDRESSIPNOTFOUND = "InvalidParameterValue.AddressIpNotFound"
//  MISSINGPARAMETER_MULTIMISSINGPARAMETER = "MissingParameter.MultiMissingParameter"
//  OPERATIONDENIED_ADDRESSINARREARS = "OperationDenied.AddressInArrears"
//  UNSUPPORTEDOPERATION_ADDRESSSTATUSNOTPERMIT = "UnsupportedOperation.AddressStatusNotPermit"
func (c *Client) ModifyIPv6AddressesAttributes(request *ModifyIPv6AddressesAttributesRequest) (response *ModifyIPv6AddressesAttributesResponse, err error) {
    return c.ModifyIPv6AddressesAttributesWithContext(context.Background(), request)
}

// ModifyIPv6AddressesAttributes
// This API is used to modify the name of an EIPv6 instance.
//
// 
//
// - You can modify the name of both EIPv6 and traditional EIPv6 instances.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ADDRESSIDMALFORMED = "InvalidParameterValue.AddressIdMalformed"
//  INVALIDPARAMETERVALUE_ADDRESSIPNOTFOUND = "InvalidParameterValue.AddressIpNotFound"
//  MISSINGPARAMETER_MULTIMISSINGPARAMETER = "MissingParameter.MultiMissingParameter"
//  OPERATIONDENIED_ADDRESSINARREARS = "OperationDenied.AddressInArrears"
//  UNSUPPORTEDOPERATION_ADDRESSSTATUSNOTPERMIT = "UnsupportedOperation.AddressStatusNotPermit"
func (c *Client) ModifyIPv6AddressesAttributesWithContext(ctx context.Context, request *ModifyIPv6AddressesAttributesRequest) (response *ModifyIPv6AddressesAttributesResponse, err error) {
    if request == nil {
        request = NewModifyIPv6AddressesAttributesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyIPv6AddressesAttributes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyIPv6AddressesAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyIPv6AddressesAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyIPv6AddressesBandwidthRequest() (request *ModifyIPv6AddressesBandwidthRequest) {
    request = &ModifyIPv6AddressesBandwidthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyIPv6AddressesBandwidth")
    
    
    return
}

func NewModifyIPv6AddressesBandwidthResponse() (response *ModifyIPv6AddressesBandwidthResponse) {
    response = &ModifyIPv6AddressesBandwidthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyIPv6AddressesBandwidth
// This API is used to modify the bandwidth cap of an EIPv6 instance.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_ADDRESSIDMALFORMED = "InvalidParameterValue.AddressIdMalformed"
//  INVALIDPARAMETERVALUE_ADDRESSIPNOTFOUND = "InvalidParameterValue.AddressIpNotFound"
//  OPERATIONDENIED_ADDRESSINARREARS = "OperationDenied.AddressInArrears"
//  UNSUPPORTEDOPERATION_ADDRESSSTATUSNOTPERMIT = "UnsupportedOperation.AddressStatusNotPermit"
func (c *Client) ModifyIPv6AddressesBandwidth(request *ModifyIPv6AddressesBandwidthRequest) (response *ModifyIPv6AddressesBandwidthResponse, err error) {
    return c.ModifyIPv6AddressesBandwidthWithContext(context.Background(), request)
}

// ModifyIPv6AddressesBandwidth
// This API is used to modify the bandwidth cap of an EIPv6 instance.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_ADDRESSIDMALFORMED = "InvalidParameterValue.AddressIdMalformed"
//  INVALIDPARAMETERVALUE_ADDRESSIPNOTFOUND = "InvalidParameterValue.AddressIpNotFound"
//  OPERATIONDENIED_ADDRESSINARREARS = "OperationDenied.AddressInArrears"
//  UNSUPPORTEDOPERATION_ADDRESSSTATUSNOTPERMIT = "UnsupportedOperation.AddressStatusNotPermit"
func (c *Client) ModifyIPv6AddressesBandwidthWithContext(ctx context.Context, request *ModifyIPv6AddressesBandwidthRequest) (response *ModifyIPv6AddressesBandwidthResponse, err error) {
    if request == nil {
        request = NewModifyIPv6AddressesBandwidthRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyIPv6AddressesBandwidth")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyIPv6AddressesBandwidth require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyIPv6AddressesBandwidthResponse()
    err = c.Send(request, response)
    return
}

func NewModifyIp6AddressesBandwidthRequest() (request *ModifyIp6AddressesBandwidthRequest) {
    request = &ModifyIp6AddressesBandwidthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyIp6AddressesBandwidth")
    
    
    return
}

func NewModifyIp6AddressesBandwidthResponse() (response *ModifyIp6AddressesBandwidthResponse) {
    response = &ModifyIp6AddressesBandwidthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyIp6AddressesBandwidth
// This API is used to adjust the bandwidth limit of a classic elastic public IPv6 instance.
//
// 
//
// - You can adjust the bandwidth limit of only classic elastic public IPv6 instances.
//
// - To adjust the bandwidth limit of an elastic public IPv6 instance, call the ModifyIPv6AddressesBandwidth API.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDACCOUNT_NOTSUPPORTED = "InvalidAccount.NotSupported"
//  INVALIDADDRESSID_NOTFOUND = "InvalidAddressId.NotFound"
//  INVALIDADDRESSIDSTATE_INARREARS = "InvalidAddressIdState.InArrears"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ADDRESSIDMALFORMED = "InvalidParameterValue.AddressIdMalformed"
//  INVALIDPARAMETERVALUE_ADDRESSIPNOTFOUND = "InvalidParameterValue.AddressIpNotFound"
//  INVALIDPARAMETERVALUE_BANDWIDTHOUTOFRANGE = "InvalidParameterValue.BandwidthOutOfRange"
//  INVALIDPARAMETERVALUE_INVALIDIPV6 = "InvalidParameterValue.InvalidIpv6"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  OPERATIONDENIED_MUTEXTASKRUNNING = "OperationDenied.MutexTaskRunning"
//  UNSUPPORTEDOPERATION_ADDRESSIPINARREAR = "UnsupportedOperation.AddressIpInArrear"
//  UNSUPPORTEDOPERATION_ADDRESSIPINTERNETCHARGETYPENOTPERMIT = "UnsupportedOperation.AddressIpInternetChargeTypeNotPermit"
//  UNSUPPORTEDOPERATION_ADDRESSIPNOTSUPPORTINSTANCE = "UnsupportedOperation.AddressIpNotSupportInstance"
//  UNSUPPORTEDOPERATION_ADDRESSIPSTATUSNOTPERMIT = "UnsupportedOperation.AddressIpStatusNotPermit"
func (c *Client) ModifyIp6AddressesBandwidth(request *ModifyIp6AddressesBandwidthRequest) (response *ModifyIp6AddressesBandwidthResponse, err error) {
    return c.ModifyIp6AddressesBandwidthWithContext(context.Background(), request)
}

// ModifyIp6AddressesBandwidth
// This API is used to adjust the bandwidth limit of a classic elastic public IPv6 instance.
//
// 
//
// - You can adjust the bandwidth limit of only classic elastic public IPv6 instances.
//
// - To adjust the bandwidth limit of an elastic public IPv6 instance, call the ModifyIPv6AddressesBandwidth API.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDACCOUNT_NOTSUPPORTED = "InvalidAccount.NotSupported"
//  INVALIDADDRESSID_NOTFOUND = "InvalidAddressId.NotFound"
//  INVALIDADDRESSIDSTATE_INARREARS = "InvalidAddressIdState.InArrears"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ADDRESSIDMALFORMED = "InvalidParameterValue.AddressIdMalformed"
//  INVALIDPARAMETERVALUE_ADDRESSIPNOTFOUND = "InvalidParameterValue.AddressIpNotFound"
//  INVALIDPARAMETERVALUE_BANDWIDTHOUTOFRANGE = "InvalidParameterValue.BandwidthOutOfRange"
//  INVALIDPARAMETERVALUE_INVALIDIPV6 = "InvalidParameterValue.InvalidIpv6"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  OPERATIONDENIED_MUTEXTASKRUNNING = "OperationDenied.MutexTaskRunning"
//  UNSUPPORTEDOPERATION_ADDRESSIPINARREAR = "UnsupportedOperation.AddressIpInArrear"
//  UNSUPPORTEDOPERATION_ADDRESSIPINTERNETCHARGETYPENOTPERMIT = "UnsupportedOperation.AddressIpInternetChargeTypeNotPermit"
//  UNSUPPORTEDOPERATION_ADDRESSIPNOTSUPPORTINSTANCE = "UnsupportedOperation.AddressIpNotSupportInstance"
//  UNSUPPORTEDOPERATION_ADDRESSIPSTATUSNOTPERMIT = "UnsupportedOperation.AddressIpStatusNotPermit"
func (c *Client) ModifyIp6AddressesBandwidthWithContext(ctx context.Context, request *ModifyIp6AddressesBandwidthRequest) (response *ModifyIp6AddressesBandwidthResponse, err error) {
    if request == nil {
        request = NewModifyIp6AddressesBandwidthRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyIp6AddressesBandwidth")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyIp6AddressesBandwidth require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyIp6AddressesBandwidthResponse()
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
//  RESOURCENOTFOUND = "ResourceNotFound"
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
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ATTACHMENTNOTFOUND = "UnsupportedOperation.AttachmentNotFound"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) ModifyIpv6AddressesAttributeWithContext(ctx context.Context, request *ModifyIpv6AddressesAttributeRequest) (response *ModifyIpv6AddressesAttributeResponse, err error) {
    if request == nil {
        request = NewModifyIpv6AddressesAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyIpv6AddressesAttribute")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyLocalGateway")
    
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
//  UNSUPPORTEDOPERATION_INSUFFICIENTINTERNETSERVICEPROVIDERS = "UnsupportedOperation.InsufficientInternetServiceProviders"
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
//  UNSUPPORTEDOPERATION_INSUFFICIENTINTERNETSERVICEPROVIDERS = "UnsupportedOperation.InsufficientInternetServiceProviders"
func (c *Client) ModifyNatGatewayAttributeWithContext(ctx context.Context, request *ModifyNatGatewayAttributeRequest) (response *ModifyNatGatewayAttributeResponse, err error) {
    if request == nil {
        request = NewModifyNatGatewayAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyNatGatewayAttribute")
    
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
//  INVALIDPARAMETERVALUE_NATGATEWAYDNATRULENOTEXISTS = "InvalidParameterValue.NatGatewayDnatRuleNotExists"
//  INVALIDPARAMETERVALUE_NATGATEWAYDNATRULEPIPNEEDVM = "InvalidParameterValue.NatGatewayDnatRulePipNeedVm"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION_NATGATEWAYRULEPIPEXISTS = "UnsupportedOperation.NatGatewayRulePipExists"
//  UNSUPPORTEDOPERATION_RECORDEXISTS = "UnsupportedOperation.RecordExists"
func (c *Client) ModifyNatGatewayDestinationIpPortTranslationNatRule(request *ModifyNatGatewayDestinationIpPortTranslationNatRuleRequest) (response *ModifyNatGatewayDestinationIpPortTranslationNatRuleResponse, err error) {
    return c.ModifyNatGatewayDestinationIpPortTranslationNatRuleWithContext(context.Background(), request)
}

// ModifyNatGatewayDestinationIpPortTranslationNatRule
// This API is used to modify the port forwarding rule of a NAT gateway.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_NATGATEWAYDNATRULENOTEXISTS = "InvalidParameterValue.NatGatewayDnatRuleNotExists"
//  INVALIDPARAMETERVALUE_NATGATEWAYDNATRULEPIPNEEDVM = "InvalidParameterValue.NatGatewayDnatRulePipNeedVm"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION_NATGATEWAYRULEPIPEXISTS = "UnsupportedOperation.NatGatewayRulePipExists"
//  UNSUPPORTEDOPERATION_RECORDEXISTS = "UnsupportedOperation.RecordExists"
func (c *Client) ModifyNatGatewayDestinationIpPortTranslationNatRuleWithContext(ctx context.Context, request *ModifyNatGatewayDestinationIpPortTranslationNatRuleRequest) (response *ModifyNatGatewayDestinationIpPortTranslationNatRuleResponse, err error) {
    if request == nil {
        request = NewModifyNatGatewayDestinationIpPortTranslationNatRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyNatGatewayDestinationIpPortTranslationNatRule")
    
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
//  RESOURCENOTFOUND = "ResourceNotFound"
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
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_UNBINDEIP = "UnsupportedOperation.UnbindEIP"
func (c *Client) ModifyNatGatewaySourceIpTranslationNatRuleWithContext(ctx context.Context, request *ModifyNatGatewaySourceIpTranslationNatRuleRequest) (response *ModifyNatGatewaySourceIpTranslationNatRuleResponse, err error) {
    if request == nil {
        request = NewModifyNatGatewaySourceIpTranslationNatRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyNatGatewaySourceIpTranslationNatRule")
    
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
//  UNSUPPORTEDOPERATION_CONFLICTWITHDOCKERROUTE = "UnsupportedOperation.ConflictWithDockerRoute"
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
//  UNSUPPORTEDOPERATION_CONFLICTWITHDOCKERROUTE = "UnsupportedOperation.ConflictWithDockerRoute"
//  UNSUPPORTEDOPERATION_ECMPWITHUSERROUTE = "UnsupportedOperation.EcmpWithUserRoute"
func (c *Client) ModifyNetDetectWithContext(ctx context.Context, request *ModifyNetDetectRequest) (response *ModifyNetDetectResponse, err error) {
    if request == nil {
        request = NewModifyNetDetectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyNetDetect")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyNetworkAclAttribute")
    
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
//  INVALIDPARAMETER_ACLTYPEMISMATCH = "InvalidParameter.AclTypeMismatch"
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
//  INVALIDPARAMETER_ACLTYPEMISMATCH = "InvalidParameter.AclTypeMismatch"
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyNetworkAclEntries")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNetworkAclEntries require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNetworkAclEntriesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNetworkAclQuintupleEntriesRequest() (request *ModifyNetworkAclQuintupleEntriesRequest) {
    request = &ModifyNetworkAclQuintupleEntriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyNetworkAclQuintupleEntries")
    
    
    return
}

func NewModifyNetworkAclQuintupleEntriesResponse() (response *ModifyNetworkAclQuintupleEntriesResponse) {
    response = &ModifyNetworkAclQuintupleEntriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyNetworkAclQuintupleEntries
// This API is used to modify the in/outbound rules of the network ACL quintuple. In the `NetworkAclQuintupleEntrySet` parameters, `NetworkAclQuintupleEntryId` is required for `NetworkAclQuintupleEntry`.
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
func (c *Client) ModifyNetworkAclQuintupleEntries(request *ModifyNetworkAclQuintupleEntriesRequest) (response *ModifyNetworkAclQuintupleEntriesResponse, err error) {
    return c.ModifyNetworkAclQuintupleEntriesWithContext(context.Background(), request)
}

// ModifyNetworkAclQuintupleEntries
// This API is used to modify the in/outbound rules of the network ACL quintuple. In the `NetworkAclQuintupleEntrySet` parameters, `NetworkAclQuintupleEntryId` is required for `NetworkAclQuintupleEntry`.
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
func (c *Client) ModifyNetworkAclQuintupleEntriesWithContext(ctx context.Context, request *ModifyNetworkAclQuintupleEntriesRequest) (response *ModifyNetworkAclQuintupleEntriesResponse, err error) {
    if request == nil {
        request = NewModifyNetworkAclQuintupleEntriesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyNetworkAclQuintupleEntries")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNetworkAclQuintupleEntries require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNetworkAclQuintupleEntriesResponse()
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
//  LIMITEXCEEDED_ACTIONLIMITED = "LimitExceeded.ActionLimited"
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
//  LIMITEXCEEDED_ACTIONLIMITED = "LimitExceeded.ActionLimited"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_SUBENINOTSUPPORTTRUNKING = "UnsupportedOperation.SubEniNotSupportTrunking"
func (c *Client) ModifyNetworkInterfaceAttributeWithContext(ctx context.Context, request *ModifyNetworkInterfaceAttributeRequest) (response *ModifyNetworkInterfaceAttributeResponse, err error) {
    if request == nil {
        request = NewModifyNetworkInterfaceAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyNetworkInterfaceAttribute")
    
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
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED_ACTIONLIMITED = "LimitExceeded.ActionLimited"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) ModifyPrivateIpAddressesAttribute(request *ModifyPrivateIpAddressesAttributeRequest) (response *ModifyPrivateIpAddressesAttributeResponse, err error) {
    return c.ModifyPrivateIpAddressesAttributeWithContext(context.Background(), request)
}

// ModifyPrivateIpAddressesAttribute
// This API (ModifyPrivateIpAddressesAttribute) is used to modify the private IP attributes of an ENI.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED_ACTIONLIMITED = "LimitExceeded.ActionLimited"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) ModifyPrivateIpAddressesAttributeWithContext(ctx context.Context, request *ModifyPrivateIpAddressesAttributeRequest) (response *ModifyPrivateIpAddressesAttributeResponse, err error) {
    if request == nil {
        request = NewModifyPrivateIpAddressesAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyPrivateIpAddressesAttribute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPrivateIpAddressesAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPrivateIpAddressesAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyReserveIpAddressRequest() (request *ModifyReserveIpAddressRequest) {
    request = &ModifyReserveIpAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyReserveIpAddress")
    
    
    return
}

func NewModifyReserveIpAddressResponse() (response *ModifyReserveIpAddressResponse) {
    response = &ModifyReserveIpAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyReserveIpAddress
// This API is used to modify a reserved private IP address.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RESOURCENOTFOUND = "InvalidParameterValue.ResourceNotFound"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyReserveIpAddress(request *ModifyReserveIpAddressRequest) (response *ModifyReserveIpAddressResponse, err error) {
    return c.ModifyReserveIpAddressWithContext(context.Background(), request)
}

// ModifyReserveIpAddress
// This API is used to modify a reserved private IP address.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RESOURCENOTFOUND = "InvalidParameterValue.ResourceNotFound"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyReserveIpAddressWithContext(ctx context.Context, request *ModifyReserveIpAddressRequest) (response *ModifyReserveIpAddressResponse, err error) {
    if request == nil {
        request = NewModifyReserveIpAddressRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyReserveIpAddress")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyReserveIpAddress require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyReserveIpAddressResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRoutePolicyAttributeRequest() (request *ModifyRoutePolicyAttributeRequest) {
    request = &ModifyRoutePolicyAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyRoutePolicyAttribute")
    
    
    return
}

func NewModifyRoutePolicyAttributeResponse() (response *ModifyRoutePolicyAttributeResponse) {
    response = &ModifyRoutePolicyAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRoutePolicyAttribute
// This API is used to modify the route reception policy attribute.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_MUSTHASONE = "InvalidParameterValue.MustHasOne"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyRoutePolicyAttribute(request *ModifyRoutePolicyAttributeRequest) (response *ModifyRoutePolicyAttributeResponse, err error) {
    return c.ModifyRoutePolicyAttributeWithContext(context.Background(), request)
}

// ModifyRoutePolicyAttribute
// This API is used to modify the route reception policy attribute.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_MUSTHASONE = "InvalidParameterValue.MustHasOne"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyRoutePolicyAttributeWithContext(ctx context.Context, request *ModifyRoutePolicyAttributeRequest) (response *ModifyRoutePolicyAttributeResponse, err error) {
    if request == nil {
        request = NewModifyRoutePolicyAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyRoutePolicyAttribute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRoutePolicyAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRoutePolicyAttributeResponse()
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
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyRouteTableAttribute(request *ModifyRouteTableAttributeRequest) (response *ModifyRouteTableAttributeResponse, err error) {
    return c.ModifyRouteTableAttributeWithContext(context.Background(), request)
}

// ModifyRouteTableAttribute
// This API (ModifyRouteTableAttribute) is used to modify the attributes of a route table.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyRouteTableAttributeWithContext(ctx context.Context, request *ModifyRouteTableAttributeRequest) (response *ModifyRouteTableAttributeResponse, err error) {
    if request == nil {
        request = NewModifyRouteTableAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyRouteTableAttribute")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifySecurityGroupAttribute")
    
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONIDMISMATCH = "InvalidParameterValue.MemberApprovalApplicationIdMismatch"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONNOTAPPROVED = "InvalidParameterValue.MemberApprovalApplicationNotApproved"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONREJECTED = "InvalidParameterValue.MemberApprovalApplicationRejected"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONSTARTED = "InvalidParameterValue.MemberApprovalApplicationStarted"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_CLBPOLICYEXCEEDLIMIT = "UnsupportedOperation.ClbPolicyExceedLimit"
//  UNSUPPORTEDOPERATION_DUPLICATEPOLICY = "UnsupportedOperation.DuplicatePolicy"
//  UNSUPPORTEDOPERATION_LOCKEDRESOURCES = "UnsupportedOperation.LockedResources"
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONIDMISMATCH = "InvalidParameterValue.MemberApprovalApplicationIdMismatch"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONNOTAPPROVED = "InvalidParameterValue.MemberApprovalApplicationNotApproved"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONREJECTED = "InvalidParameterValue.MemberApprovalApplicationRejected"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONSTARTED = "InvalidParameterValue.MemberApprovalApplicationStarted"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_CLBPOLICYEXCEEDLIMIT = "UnsupportedOperation.ClbPolicyExceedLimit"
//  UNSUPPORTEDOPERATION_DUPLICATEPOLICY = "UnsupportedOperation.DuplicatePolicy"
//  UNSUPPORTEDOPERATION_LOCKEDRESOURCES = "UnsupportedOperation.LockedResources"
func (c *Client) ModifySecurityGroupPoliciesWithContext(ctx context.Context, request *ModifySecurityGroupPoliciesRequest) (response *ModifySecurityGroupPoliciesResponse, err error) {
    if request == nil {
        request = NewModifySecurityGroupPoliciesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifySecurityGroupPolicies")
    
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
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
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
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyServiceTemplateAttribute")
    
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
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
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
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) ModifyServiceTemplateGroupAttributeWithContext(ctx context.Context, request *ModifyServiceTemplateGroupAttributeRequest) (response *ModifyServiceTemplateGroupAttributeResponse, err error) {
    if request == nil {
        request = NewModifyServiceTemplateGroupAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyServiceTemplateGroupAttribute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyServiceTemplateGroupAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyServiceTemplateGroupAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifySnapshotPoliciesRequest() (request *ModifySnapshotPoliciesRequest) {
    request = &ModifySnapshotPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "ModifySnapshotPolicies")
    
    
    return
}

func NewModifySnapshotPoliciesResponse() (response *ModifySnapshotPoliciesResponse) {
    response = &ModifySnapshotPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySnapshotPolicies
// This API is used to modify specified snapshot policies.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_SNAPSHOTBACKUPTYPEMODIFY = "UnsupportedOperation.SnapshotBackupTypeModify"
func (c *Client) ModifySnapshotPolicies(request *ModifySnapshotPoliciesRequest) (response *ModifySnapshotPoliciesResponse, err error) {
    return c.ModifySnapshotPoliciesWithContext(context.Background(), request)
}

// ModifySnapshotPolicies
// This API is used to modify specified snapshot policies.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_SNAPSHOTBACKUPTYPEMODIFY = "UnsupportedOperation.SnapshotBackupTypeModify"
func (c *Client) ModifySnapshotPoliciesWithContext(ctx context.Context, request *ModifySnapshotPoliciesRequest) (response *ModifySnapshotPoliciesResponse, err error) {
    if request == nil {
        request = NewModifySnapshotPoliciesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifySnapshotPolicies")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySnapshotPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySnapshotPoliciesResponse()
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifySubnetAttribute(request *ModifySubnetAttributeRequest) (response *ModifySubnetAttributeResponse, err error) {
    return c.ModifySubnetAttributeWithContext(context.Background(), request)
}

// ModifySubnetAttribute
// This API (ModifySubnetAttribute) is used to modify subnet attributes.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifySubnetAttributeWithContext(ctx context.Context, request *ModifySubnetAttributeRequest) (response *ModifySubnetAttributeResponse, err error) {
    if request == nil {
        request = NewModifySubnetAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifySubnetAttribute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySubnetAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySubnetAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTemplateMemberRequest() (request *ModifyTemplateMemberRequest) {
    request = &ModifyTemplateMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyTemplateMember")
    
    
    return
}

func NewModifyTemplateMemberResponse() (response *ModifyTemplateMemberResponse) {
    response = &ModifyTemplateMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTemplateMember
// This API is used to modify a parameter template of the IP address, protocol port, IP address group, or protocol port group type.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) ModifyTemplateMember(request *ModifyTemplateMemberRequest) (response *ModifyTemplateMemberResponse, err error) {
    return c.ModifyTemplateMemberWithContext(context.Background(), request)
}

// ModifyTemplateMember
// This API is used to modify a parameter template of the IP address, protocol port, IP address group, or protocol port group type.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) ModifyTemplateMemberWithContext(ctx context.Context, request *ModifyTemplateMemberRequest) (response *ModifyTemplateMemberResponse, err error) {
    if request == nil {
        request = NewModifyTemplateMemberRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyTemplateMember")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTemplateMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTemplateMemberResponse()
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_ENABLEMULTICAST = "UnsupportedOperation.EnableMulticast"
//  UNSUPPORTEDOPERATION_NOTSUPPORTEDUPDATECCNROUTEPUBLISH = "UnsupportedOperation.NotSupportedUpdateCcnRoutePublish"
func (c *Client) ModifyVpcAttribute(request *ModifyVpcAttributeRequest) (response *ModifyVpcAttributeResponse, err error) {
    return c.ModifyVpcAttributeWithContext(context.Background(), request)
}

// ModifyVpcAttribute
// This API (ModifyVpcAttribute) is used to modify VPC attributes.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_ENABLEMULTICAST = "UnsupportedOperation.EnableMulticast"
//  UNSUPPORTEDOPERATION_NOTSUPPORTEDUPDATECCNROUTEPUBLISH = "UnsupportedOperation.NotSupportedUpdateCcnRoutePublish"
func (c *Client) ModifyVpcAttributeWithContext(ctx context.Context, request *ModifyVpcAttributeRequest) (response *ModifyVpcAttributeResponse, err error) {
    if request == nil {
        request = NewModifyVpcAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyVpcAttribute")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyVpcEndPointAttribute")
    
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyVpcEndPointServiceAttribute")
    
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
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
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
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_UINNOTFOUND = "UnsupportedOperation.UinNotFound"
func (c *Client) ModifyVpcEndPointServiceWhiteListWithContext(ctx context.Context, request *ModifyVpcEndPointServiceWhiteListRequest) (response *ModifyVpcEndPointServiceWhiteListResponse, err error) {
    if request == nil {
        request = NewModifyVpcEndPointServiceWhiteListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyVpcEndPointServiceWhiteList")
    
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
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_VPNCONNCIDRCONFLICT = "InvalidParameterValue.VpnConnCidrConflict"
//  INVALIDPARAMETERVALUE_VPNCONNHEALTHCHECKIPCONFLICT = "InvalidParameterValue.VpnConnHealthCheckIpConflict"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_VPNCONNINVALIDSTATE = "UnsupportedOperation.VpnConnInvalidState"
//  UNSUPPORTEDOPERATION_VPNCONNSPDOVERLAP = "UnsupportedOperation.VpnConnSPDOverlap"
func (c *Client) ModifyVpnConnectionAttribute(request *ModifyVpnConnectionAttributeRequest) (response *ModifyVpnConnectionAttributeResponse, err error) {
    return c.ModifyVpnConnectionAttributeWithContext(context.Background(), request)
}

// ModifyVpnConnectionAttribute
// This API (ModifyVpnConnectionAttribute) is used to modify VPN tunnels.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_VPNCONNCIDRCONFLICT = "InvalidParameterValue.VpnConnCidrConflict"
//  INVALIDPARAMETERVALUE_VPNCONNHEALTHCHECKIPCONFLICT = "InvalidParameterValue.VpnConnHealthCheckIpConflict"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_VPNCONNINVALIDSTATE = "UnsupportedOperation.VpnConnInvalidState"
//  UNSUPPORTEDOPERATION_VPNCONNSPDOVERLAP = "UnsupportedOperation.VpnConnSPDOverlap"
func (c *Client) ModifyVpnConnectionAttributeWithContext(ctx context.Context, request *ModifyVpnConnectionAttributeRequest) (response *ModifyVpnConnectionAttributeResponse, err error) {
    if request == nil {
        request = NewModifyVpnConnectionAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyVpnConnectionAttribute")
    
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
//  UNSUPPORTEDOPERATION_VPNUNSUPPORTEDMODIFYBGPASN = "UnsupportedOperation.VpnUnsupportedModifyBgpAsn"
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
//  UNSUPPORTEDOPERATION_VPNUNSUPPORTEDMODIFYBGPASN = "UnsupportedOperation.VpnUnsupportedModifyBgpAsn"
func (c *Client) ModifyVpnGatewayAttributeWithContext(ctx context.Context, request *ModifyVpnGatewayAttributeRequest) (response *ModifyVpnGatewayAttributeResponse, err error) {
    if request == nil {
        request = NewModifyVpnGatewayAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyVpnGatewayAttribute")
    
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
// This API is used to modify VPN gateway-based CCN routes.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyVpnGatewayCcnRoutes(request *ModifyVpnGatewayCcnRoutesRequest) (response *ModifyVpnGatewayCcnRoutesResponse, err error) {
    return c.ModifyVpnGatewayCcnRoutesWithContext(context.Background(), request)
}

// ModifyVpnGatewayCcnRoutes
// This API is used to modify VPN gateway-based CCN routes.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyVpnGatewayCcnRoutesWithContext(ctx context.Context, request *ModifyVpnGatewayCcnRoutesRequest) (response *ModifyVpnGatewayCcnRoutesResponse, err error) {
    if request == nil {
        request = NewModifyVpnGatewayCcnRoutesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyVpnGatewayCcnRoutes")
    
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
// This API is used to modify VPN gateway routes.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_RECORDNOTEXISTS = "UnsupportedOperation.RecordNotExists"
func (c *Client) ModifyVpnGatewayRoutes(request *ModifyVpnGatewayRoutesRequest) (response *ModifyVpnGatewayRoutesResponse, err error) {
    return c.ModifyVpnGatewayRoutesWithContext(context.Background(), request)
}

// ModifyVpnGatewayRoutes
// This API is used to modify VPN gateway routes.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_RECORDNOTEXISTS = "UnsupportedOperation.RecordNotExists"
func (c *Client) ModifyVpnGatewayRoutesWithContext(ctx context.Context, request *ModifyVpnGatewayRoutesRequest) (response *ModifyVpnGatewayRoutesResponse, err error) {
    if request == nil {
        request = NewModifyVpnGatewayRoutesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ModifyVpnGatewayRoutes")
    
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
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_ASSOCIATEDVPCOFCCNHADNATROUTE = "UnsupportedOperation.AssociatedVpcOfCcnHadNatRoute"
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
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_ASSOCIATEDVPCOFCCNHADNATROUTE = "UnsupportedOperation.AssociatedVpcOfCcnHadNatRoute"
//  UNSUPPORTEDOPERATION_CCNNOTATTACHED = "UnsupportedOperation.CcnNotAttached"
//  UNSUPPORTEDOPERATION_INVALIDSTATUSNOTIFYCCN = "UnsupportedOperation.InvalidStatusNotifyCcn"
//  UNSUPPORTEDOPERATION_NOTIFYCCN = "UnsupportedOperation.NotifyCcn"
//  UNSUPPORTEDOPERATION_SYSTEMROUTE = "UnsupportedOperation.SystemRoute"
func (c *Client) NotifyRoutesWithContext(ctx context.Context, request *NotifyRoutesRequest) (response *NotifyRoutesResponse, err error) {
    if request == nil {
        request = NewNotifyRoutesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "NotifyRoutes")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "RefreshDirectConnectGatewayRouteToNatGateway")
    
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
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
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
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CCNNOTATTACHED = "UnsupportedOperation.CcnNotAttached"
//  UNSUPPORTEDOPERATION_NOTPENDINGCCNINSTANCE = "UnsupportedOperation.NotPendingCcnInstance"
func (c *Client) RejectAttachCcnInstancesWithContext(ctx context.Context, request *RejectAttachCcnInstancesRequest) (response *RejectAttachCcnInstancesResponse, err error) {
    if request == nil {
        request = NewRejectAttachCcnInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "RejectAttachCcnInstances")
    
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
//  FAILEDOPERATION_TASKFAILED = "FailedOperation.TaskFailed"
//  INVALIDADDRESSID_BLOCKED = "InvalidAddressId.Blocked"
//  INVALIDADDRESSID_NOTFOUND = "InvalidAddressId.NotFound"
//  INVALIDADDRESSSTATE = "InvalidAddressState"
//  INVALIDPARAMETERVALUE_ADDRESSIDMALFORMED = "InvalidParameterValue.AddressIdMalformed"
//  INVALIDPARAMETERVALUE_ADDRESSINTERNETCHARGETYPECONFLICT = "InvalidParameterValue.AddressInternetChargeTypeConflict"
//  INVALIDPARAMETERVALUE_ADDRESSNOTFOUND = "InvalidParameterValue.AddressNotFound"
//  INVALIDPARAMETERVALUE_ADDRESSTYPECONFLICT = "InvalidParameterValue.AddressTypeConflict"
//  LIMITEXCEEDED_ACCOUNTRETURNQUOTA = "LimitExceeded.AccountReturnQuota"
//  OPERATIONDENIED_MUTEXTASKRUNNING = "OperationDenied.MutexTaskRunning"
//  UNSUPPORTEDOPERATION_ADDRESSSTATUSNOTPERMIT = "UnsupportedOperation.AddressStatusNotPermit"
//  UNSUPPORTEDOPERATION_UNPAIDORDERALREADYEXISTS = "UnsupportedOperation.UnpaidOrderAlreadyExists"
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
//  FAILEDOPERATION_TASKFAILED = "FailedOperation.TaskFailed"
//  INVALIDADDRESSID_BLOCKED = "InvalidAddressId.Blocked"
//  INVALIDADDRESSID_NOTFOUND = "InvalidAddressId.NotFound"
//  INVALIDADDRESSSTATE = "InvalidAddressState"
//  INVALIDPARAMETERVALUE_ADDRESSIDMALFORMED = "InvalidParameterValue.AddressIdMalformed"
//  INVALIDPARAMETERVALUE_ADDRESSINTERNETCHARGETYPECONFLICT = "InvalidParameterValue.AddressInternetChargeTypeConflict"
//  INVALIDPARAMETERVALUE_ADDRESSNOTFOUND = "InvalidParameterValue.AddressNotFound"
//  INVALIDPARAMETERVALUE_ADDRESSTYPECONFLICT = "InvalidParameterValue.AddressTypeConflict"
//  LIMITEXCEEDED_ACCOUNTRETURNQUOTA = "LimitExceeded.AccountReturnQuota"
//  OPERATIONDENIED_MUTEXTASKRUNNING = "OperationDenied.MutexTaskRunning"
//  UNSUPPORTEDOPERATION_ADDRESSSTATUSNOTPERMIT = "UnsupportedOperation.AddressStatusNotPermit"
//  UNSUPPORTEDOPERATION_UNPAIDORDERALREADYEXISTS = "UnsupportedOperation.UnpaidOrderAlreadyExists"
func (c *Client) ReleaseAddressesWithContext(ctx context.Context, request *ReleaseAddressesRequest) (response *ReleaseAddressesResponse, err error) {
    if request == nil {
        request = NewReleaseAddressesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ReleaseAddresses")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReleaseAddresses require credential")
    }

    request.SetContext(ctx)
    
    response = NewReleaseAddressesResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseIPv6AddressesRequest() (request *ReleaseIPv6AddressesRequest) {
    request = &ReleaseIPv6AddressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "ReleaseIPv6Addresses")
    
    
    return
}

func NewReleaseIPv6AddressesResponse() (response *ReleaseIPv6AddressesResponse) {
    response = &ReleaseIPv6AddressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReleaseIPv6Addresses
// This API is used to release one or more EIPv6 instances.
//
// 
//
// - You can release the obtained EIPv6 instances. To use them again, please reapply.
//
// - Only EIPv6 instances in UNBIND status can be released.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_ADDRESSIDMALFORMED = "InvalidParameterValue.AddressIdMalformed"
//  INVALIDPARAMETERVALUE_ADDRESSIPNOTFOUND = "InvalidParameterValue.AddressIpNotFound"
//  OPERATIONDENIED_ADDRESSINARREARS = "OperationDenied.AddressInArrears"
//  UNSUPPORTEDOPERATION_ADDRESSSTATUSNOTPERMIT = "UnsupportedOperation.AddressStatusNotPermit"
func (c *Client) ReleaseIPv6Addresses(request *ReleaseIPv6AddressesRequest) (response *ReleaseIPv6AddressesResponse, err error) {
    return c.ReleaseIPv6AddressesWithContext(context.Background(), request)
}

// ReleaseIPv6Addresses
// This API is used to release one or more EIPv6 instances.
//
// 
//
// - You can release the obtained EIPv6 instances. To use them again, please reapply.
//
// - Only EIPv6 instances in UNBIND status can be released.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_ADDRESSIDMALFORMED = "InvalidParameterValue.AddressIdMalformed"
//  INVALIDPARAMETERVALUE_ADDRESSIPNOTFOUND = "InvalidParameterValue.AddressIpNotFound"
//  OPERATIONDENIED_ADDRESSINARREARS = "OperationDenied.AddressInArrears"
//  UNSUPPORTEDOPERATION_ADDRESSSTATUSNOTPERMIT = "UnsupportedOperation.AddressStatusNotPermit"
func (c *Client) ReleaseIPv6AddressesWithContext(ctx context.Context, request *ReleaseIPv6AddressesRequest) (response *ReleaseIPv6AddressesResponse, err error) {
    if request == nil {
        request = NewReleaseIPv6AddressesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ReleaseIPv6Addresses")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReleaseIPv6Addresses require credential")
    }

    request.SetContext(ctx)
    
    response = NewReleaseIPv6AddressesResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseIp6AddressesBandwidthRequest() (request *ReleaseIp6AddressesBandwidthRequest) {
    request = &ReleaseIp6AddressesBandwidthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "ReleaseIp6AddressesBandwidth")
    
    
    return
}

func NewReleaseIp6AddressesBandwidthResponse() (response *ReleaseIp6AddressesBandwidthResponse) {
    response = &ReleaseIp6AddressesBandwidthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReleaseIp6AddressesBandwidth
// This API is used to release the IPv6 public network bandwidth of classic elastic public IPv6 instances.
//
// 
//
// - Classic elastic public IPv6 addresses still have the IPv6 private network communication capability after the public network bandwidth is released.
//
// - To allocate IPV6 public network bandwidth, call the AllocateIp6AddressesBandwidth API.
//
// error code that may be returned:
//  FAILEDOPERATION_TASKFAILED = "FailedOperation.TaskFailed"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ADDRESSIDMALFORMED = "InvalidParameterValue.AddressIdMalformed"
//  INVALIDPARAMETERVALUE_ADDRESSIPNOTFOUND = "InvalidParameterValue.AddressIpNotFound"
//  INVALIDPARAMETERVALUE_INVALIDIPV6 = "InvalidParameterValue.InvalidIpv6"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  OPERATIONDENIED_MUTEXTASKRUNNING = "OperationDenied.MutexTaskRunning"
func (c *Client) ReleaseIp6AddressesBandwidth(request *ReleaseIp6AddressesBandwidthRequest) (response *ReleaseIp6AddressesBandwidthResponse, err error) {
    return c.ReleaseIp6AddressesBandwidthWithContext(context.Background(), request)
}

// ReleaseIp6AddressesBandwidth
// This API is used to release the IPv6 public network bandwidth of classic elastic public IPv6 instances.
//
// 
//
// - Classic elastic public IPv6 addresses still have the IPv6 private network communication capability after the public network bandwidth is released.
//
// - To allocate IPV6 public network bandwidth, call the AllocateIp6AddressesBandwidth API.
//
// error code that may be returned:
//  FAILEDOPERATION_TASKFAILED = "FailedOperation.TaskFailed"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ADDRESSIDMALFORMED = "InvalidParameterValue.AddressIdMalformed"
//  INVALIDPARAMETERVALUE_ADDRESSIPNOTFOUND = "InvalidParameterValue.AddressIpNotFound"
//  INVALIDPARAMETERVALUE_INVALIDIPV6 = "InvalidParameterValue.InvalidIpv6"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  OPERATIONDENIED_MUTEXTASKRUNNING = "OperationDenied.MutexTaskRunning"
func (c *Client) ReleaseIp6AddressesBandwidthWithContext(ctx context.Context, request *ReleaseIp6AddressesBandwidthRequest) (response *ReleaseIp6AddressesBandwidthResponse, err error) {
    if request == nil {
        request = NewReleaseIp6AddressesBandwidthRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ReleaseIp6AddressesBandwidth")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReleaseIp6AddressesBandwidth require credential")
    }

    request.SetContext(ctx)
    
    response = NewReleaseIp6AddressesBandwidthResponse()
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
//  FAILEDOPERATION_IPTYPENOTPERMIT = "FailedOperation.IpTypeNotPermit"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
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
//  FAILEDOPERATION_IPTYPENOTPERMIT = "FailedOperation.IpTypeNotPermit"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "RemoveBandwidthPackageResources")
    
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
//  UNSUPPORTEDOPERATION_INSUFFICIENTFUNDS = "UnsupportedOperation.InsufficientFunds"
func (c *Client) RenewVpnGateway(request *RenewVpnGatewayRequest) (response *RenewVpnGatewayResponse, err error) {
    return c.RenewVpnGatewayWithContext(context.Background(), request)
}

// RenewVpnGateway
// This API (RenewVpnGateway) is used to renew prepaid (monthly subscription) VPN gateways. Currently, only IPSEC gateways are supported.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_INSUFFICIENTFUNDS = "UnsupportedOperation.InsufficientFunds"
func (c *Client) RenewVpnGatewayWithContext(ctx context.Context, request *RenewVpnGatewayRequest) (response *RenewVpnGatewayResponse, err error) {
    if request == nil {
        request = NewRenewVpnGatewayRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "RenewVpnGateway")
    
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
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ReplaceDirectConnectGatewayCcnRoutes(request *ReplaceDirectConnectGatewayCcnRoutesRequest) (response *ReplaceDirectConnectGatewayCcnRoutesResponse, err error) {
    return c.ReplaceDirectConnectGatewayCcnRoutesWithContext(context.Background(), request)
}

// ReplaceDirectConnectGatewayCcnRoutes
// This API (ReplaceDirectConnectGatewayCcnRoutes) is used to modify the specified route according to the route ID. Batch modification is supported.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ReplaceDirectConnectGatewayCcnRoutesWithContext(ctx context.Context, request *ReplaceDirectConnectGatewayCcnRoutesRequest) (response *ReplaceDirectConnectGatewayCcnRoutesResponse, err error) {
    if request == nil {
        request = NewReplaceDirectConnectGatewayCcnRoutesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ReplaceDirectConnectGatewayCcnRoutes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReplaceDirectConnectGatewayCcnRoutes require credential")
    }

    request.SetContext(ctx)
    
    response = NewReplaceDirectConnectGatewayCcnRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewReplaceRoutePolicyAssociationsRequest() (request *ReplaceRoutePolicyAssociationsRequest) {
    request = &ReplaceRoutePolicyAssociationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "ReplaceRoutePolicyAssociations")
    
    
    return
}

func NewReplaceRoutePolicyAssociationsResponse() (response *ReplaceRoutePolicyAssociationsResponse) {
    response = &ReplaceRoutePolicyAssociationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReplaceRoutePolicyAssociations
// This API is used to modify the binding Priority (Priority) based on the route reception policy instance ID (RoutePolicyId) and route table instance ID (RouteTableId), supporting batch modification.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_ROUTEPOLICYASSOCIATIONEXISTS = "InvalidParameterValue.RoutePolicyAssociationExists"
//  INVALIDPARAMETERVALUE_ROUTEPOLICYPRIORITYCONFLICT = "InvalidParameterValue.RoutePolicyPriorityConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ReplaceRoutePolicyAssociations(request *ReplaceRoutePolicyAssociationsRequest) (response *ReplaceRoutePolicyAssociationsResponse, err error) {
    return c.ReplaceRoutePolicyAssociationsWithContext(context.Background(), request)
}

// ReplaceRoutePolicyAssociations
// This API is used to modify the binding Priority (Priority) based on the route reception policy instance ID (RoutePolicyId) and route table instance ID (RouteTableId), supporting batch modification.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_ROUTEPOLICYASSOCIATIONEXISTS = "InvalidParameterValue.RoutePolicyAssociationExists"
//  INVALIDPARAMETERVALUE_ROUTEPOLICYPRIORITYCONFLICT = "InvalidParameterValue.RoutePolicyPriorityConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ReplaceRoutePolicyAssociationsWithContext(ctx context.Context, request *ReplaceRoutePolicyAssociationsRequest) (response *ReplaceRoutePolicyAssociationsResponse, err error) {
    if request == nil {
        request = NewReplaceRoutePolicyAssociationsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ReplaceRoutePolicyAssociations")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReplaceRoutePolicyAssociations require credential")
    }

    request.SetContext(ctx)
    
    response = NewReplaceRoutePolicyAssociationsResponse()
    err = c.Send(request, response)
    return
}

func NewReplaceRoutePolicyEntriesRequest() (request *ReplaceRoutePolicyEntriesRequest) {
    request = &ReplaceRoutePolicyEntriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "ReplaceRoutePolicyEntries")
    
    
    return
}

func NewReplaceRoutePolicyEntriesResponse() (response *ReplaceRoutePolicyEntriesResponse) {
    response = &ReplaceRoutePolicyEntriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReplaceRoutePolicyEntries
// This API is used to modify specified routing policy entries based on route reception policy rule ID and supports batch modification.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_MUSTHASONE = "InvalidParameterValue.MustHasOne"
//  INVALIDPARAMETERVALUE_ROUTEPOLICYPRIORITYCONFLICT = "InvalidParameterValue.RoutePolicyPriorityConflict"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ReplaceRoutePolicyEntries(request *ReplaceRoutePolicyEntriesRequest) (response *ReplaceRoutePolicyEntriesResponse, err error) {
    return c.ReplaceRoutePolicyEntriesWithContext(context.Background(), request)
}

// ReplaceRoutePolicyEntries
// This API is used to modify specified routing policy entries based on route reception policy rule ID and supports batch modification.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_MUSTHASONE = "InvalidParameterValue.MustHasOne"
//  INVALIDPARAMETERVALUE_ROUTEPOLICYPRIORITYCONFLICT = "InvalidParameterValue.RoutePolicyPriorityConflict"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ReplaceRoutePolicyEntriesWithContext(ctx context.Context, request *ReplaceRoutePolicyEntriesRequest) (response *ReplaceRoutePolicyEntriesResponse, err error) {
    if request == nil {
        request = NewReplaceRoutePolicyEntriesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ReplaceRoutePolicyEntries")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReplaceRoutePolicyEntries require credential")
    }

    request.SetContext(ctx)
    
    response = NewReplaceRoutePolicyEntriesResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ReplaceRouteTableAssociation")
    
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
//  UNSUPPORTEDOPERATION_RECORDNOTEXISTS = "UnsupportedOperation.RecordNotExists"
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
//  UNSUPPORTEDOPERATION_RECORDNOTEXISTS = "UnsupportedOperation.RecordNotExists"
//  UNSUPPORTEDOPERATION_SYSTEMROUTE = "UnsupportedOperation.SystemRoute"
func (c *Client) ReplaceRoutesWithContext(ctx context.Context, request *ReplaceRoutesRequest) (response *ReplaceRoutesResponse, err error) {
    if request == nil {
        request = NewReplaceRoutesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ReplaceRoutes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReplaceRoutes require credential")
    }

    request.SetContext(ctx)
    
    response = NewReplaceRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewReplaceSecurityGroupPoliciesRequest() (request *ReplaceSecurityGroupPoliciesRequest) {
    request = &ReplaceSecurityGroupPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "ReplaceSecurityGroupPolicies")
    
    
    return
}

func NewReplaceSecurityGroupPoliciesResponse() (response *ReplaceSecurityGroupPoliciesResponse) {
    response = &ReplaceSecurityGroupPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReplaceSecurityGroupPolicies
// This API is used to batch modify security group policies.
//
// Policies to modify must be in the same direction. `PolicyIndex` must be specified.
//
// error code that may be returned:
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONIDMISMATCH = "InvalidParameterValue.MemberApprovalApplicationIdMismatch"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONNOTAPPROVED = "InvalidParameterValue.MemberApprovalApplicationNotApproved"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONREJECTED = "InvalidParameterValue.MemberApprovalApplicationRejected"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONSTARTED = "InvalidParameterValue.MemberApprovalApplicationStarted"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_CLBPOLICYEXCEEDLIMIT = "UnsupportedOperation.ClbPolicyExceedLimit"
//  UNSUPPORTEDOPERATION_CLBPOLICYLIMIT = "UnsupportedOperation.ClbPolicyLimit"
//  UNSUPPORTEDOPERATION_DUPLICATEPOLICY = "UnsupportedOperation.DuplicatePolicy"
//  UNSUPPORTEDOPERATION_LOCKEDRESOURCES = "UnsupportedOperation.LockedResources"
//  UNSUPPORTEDOPERATION_VERSIONMISMATCH = "UnsupportedOperation.VersionMismatch"
func (c *Client) ReplaceSecurityGroupPolicies(request *ReplaceSecurityGroupPoliciesRequest) (response *ReplaceSecurityGroupPoliciesResponse, err error) {
    return c.ReplaceSecurityGroupPoliciesWithContext(context.Background(), request)
}

// ReplaceSecurityGroupPolicies
// This API is used to batch modify security group policies.
//
// Policies to modify must be in the same direction. `PolicyIndex` must be specified.
//
// error code that may be returned:
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONIDMISMATCH = "InvalidParameterValue.MemberApprovalApplicationIdMismatch"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONNOTAPPROVED = "InvalidParameterValue.MemberApprovalApplicationNotApproved"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONREJECTED = "InvalidParameterValue.MemberApprovalApplicationRejected"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONSTARTED = "InvalidParameterValue.MemberApprovalApplicationStarted"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_CLBPOLICYEXCEEDLIMIT = "UnsupportedOperation.ClbPolicyExceedLimit"
//  UNSUPPORTEDOPERATION_CLBPOLICYLIMIT = "UnsupportedOperation.ClbPolicyLimit"
//  UNSUPPORTEDOPERATION_DUPLICATEPOLICY = "UnsupportedOperation.DuplicatePolicy"
//  UNSUPPORTEDOPERATION_LOCKEDRESOURCES = "UnsupportedOperation.LockedResources"
//  UNSUPPORTEDOPERATION_VERSIONMISMATCH = "UnsupportedOperation.VersionMismatch"
func (c *Client) ReplaceSecurityGroupPoliciesWithContext(ctx context.Context, request *ReplaceSecurityGroupPoliciesRequest) (response *ReplaceSecurityGroupPoliciesResponse, err error) {
    if request == nil {
        request = NewReplaceSecurityGroupPoliciesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ReplaceSecurityGroupPolicies")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReplaceSecurityGroupPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewReplaceSecurityGroupPoliciesResponse()
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONIDMISMATCH = "InvalidParameterValue.MemberApprovalApplicationIdMismatch"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONNOTAPPROVED = "InvalidParameterValue.MemberApprovalApplicationNotApproved"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONREJECTED = "InvalidParameterValue.MemberApprovalApplicationRejected"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONSTARTED = "InvalidParameterValue.MemberApprovalApplicationStarted"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER_WITHGUESS = "UnknownParameter.WithGuess"
//  UNSUPPORTEDOPERATION_CLBPOLICYEXCEEDLIMIT = "UnsupportedOperation.ClbPolicyExceedLimit"
//  UNSUPPORTEDOPERATION_CLBPOLICYLIMIT = "UnsupportedOperation.ClbPolicyLimit"
//  UNSUPPORTEDOPERATION_DUPLICATEPOLICY = "UnsupportedOperation.DuplicatePolicy"
//  UNSUPPORTEDOPERATION_LOCKEDRESOURCES = "UnsupportedOperation.LockedResources"
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
//  INTERNALERROR_MODULEERROR = "InternalError.ModuleError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EMPTY = "InvalidParameterValue.Empty"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONIDMISMATCH = "InvalidParameterValue.MemberApprovalApplicationIdMismatch"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONNOTAPPROVED = "InvalidParameterValue.MemberApprovalApplicationNotApproved"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONREJECTED = "InvalidParameterValue.MemberApprovalApplicationRejected"
//  INVALIDPARAMETERVALUE_MEMBERAPPROVALAPPLICATIONSTARTED = "InvalidParameterValue.MemberApprovalApplicationStarted"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER_WITHGUESS = "UnknownParameter.WithGuess"
//  UNSUPPORTEDOPERATION_CLBPOLICYEXCEEDLIMIT = "UnsupportedOperation.ClbPolicyExceedLimit"
//  UNSUPPORTEDOPERATION_CLBPOLICYLIMIT = "UnsupportedOperation.ClbPolicyLimit"
//  UNSUPPORTEDOPERATION_DUPLICATEPOLICY = "UnsupportedOperation.DuplicatePolicy"
//  UNSUPPORTEDOPERATION_LOCKEDRESOURCES = "UnsupportedOperation.LockedResources"
//  UNSUPPORTEDOPERATION_VERSIONMISMATCH = "UnsupportedOperation.VersionMismatch"
func (c *Client) ReplaceSecurityGroupPolicyWithContext(ctx context.Context, request *ReplaceSecurityGroupPolicyRequest) (response *ReplaceSecurityGroupPolicyResponse, err error) {
    if request == nil {
        request = NewReplaceSecurityGroupPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ReplaceSecurityGroupPolicy")
    
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
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ResetAttachCcnInstances(request *ResetAttachCcnInstancesRequest) (response *ResetAttachCcnInstancesResponse, err error) {
    return c.ResetAttachCcnInstancesWithContext(context.Background(), request)
}

// ResetAttachCcnInstances
// This API (ResetAttachCcnInstances) is used to re-apply for the association operation when the application for cross-account instance association expires.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ResetAttachCcnInstancesWithContext(ctx context.Context, request *ResetAttachCcnInstancesRequest) (response *ResetAttachCcnInstancesResponse, err error) {
    if request == nil {
        request = NewResetAttachCcnInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ResetAttachCcnInstances")
    
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INSUFFICIENTFUNDS = "UnsupportedOperation.InsufficientFunds"
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INSUFFICIENTFUNDS = "UnsupportedOperation.InsufficientFunds"
//  UNSUPPORTEDOPERATION_UNPAIDORDERALREADYEXISTS = "UnsupportedOperation.UnpaidOrderAlreadyExists"
func (c *Client) ResetNatGatewayConnectionWithContext(ctx context.Context, request *ResetNatGatewayConnectionRequest) (response *ResetNatGatewayConnectionResponse, err error) {
    if request == nil {
        request = NewResetNatGatewayConnectionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ResetNatGatewayConnection")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetNatGatewayConnection require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetNatGatewayConnectionResponse()
    err = c.Send(request, response)
    return
}

func NewResetRoutePolicyAssociationsRequest() (request *ResetRoutePolicyAssociationsRequest) {
    request = &ResetRoutePolicyAssociationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "ResetRoutePolicyAssociations")
    
    
    return
}

func NewResetRoutePolicyAssociationsResponse() (response *ResetRoutePolicyAssociationsResponse) {
    response = &ResetRoutePolicyAssociationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetRoutePolicyAssociations
// This API is used to unbind the routing policy instance already bound to a specific route table instance, set up alarms for the new binding routing policy and priority.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CIDRNOTINPEERVPC = "InvalidParameterValue.CidrNotInPeerVpc"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_VPCCIDRCONFLICT = "InvalidParameterValue.VpcCidrConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER_WITHGUESS = "UnknownParameter.WithGuess"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ECMP = "UnsupportedOperation.Ecmp"
//  UNSUPPORTEDOPERATION_SYSTEMROUTE = "UnsupportedOperation.SystemRoute"
func (c *Client) ResetRoutePolicyAssociations(request *ResetRoutePolicyAssociationsRequest) (response *ResetRoutePolicyAssociationsResponse, err error) {
    return c.ResetRoutePolicyAssociationsWithContext(context.Background(), request)
}

// ResetRoutePolicyAssociations
// This API is used to unbind the routing policy instance already bound to a specific route table instance, set up alarms for the new binding routing policy and priority.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CIDRNOTINPEERVPC = "InvalidParameterValue.CidrNotInPeerVpc"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_VPCCIDRCONFLICT = "InvalidParameterValue.VpcCidrConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER_WITHGUESS = "UnknownParameter.WithGuess"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ECMP = "UnsupportedOperation.Ecmp"
//  UNSUPPORTEDOPERATION_SYSTEMROUTE = "UnsupportedOperation.SystemRoute"
func (c *Client) ResetRoutePolicyAssociationsWithContext(ctx context.Context, request *ResetRoutePolicyAssociationsRequest) (response *ResetRoutePolicyAssociationsResponse, err error) {
    if request == nil {
        request = NewResetRoutePolicyAssociationsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ResetRoutePolicyAssociations")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetRoutePolicyAssociations require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetRoutePolicyAssociationsResponse()
    err = c.Send(request, response)
    return
}

func NewResetRoutePolicyEntriesRequest() (request *ResetRoutePolicyEntriesRequest) {
    request = &ResetRoutePolicyEntriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "ResetRoutePolicyEntries")
    
    
    return
}

func NewResetRoutePolicyEntriesResponse() (response *ResetRoutePolicyEntriesResponse) {
    response = &ResetRoutePolicyEntriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetRoutePolicyEntries
// This API is used to reset the designated route reception policy entry based on the rule ID and supports batch modification.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ResetRoutePolicyEntries(request *ResetRoutePolicyEntriesRequest) (response *ResetRoutePolicyEntriesResponse, err error) {
    return c.ResetRoutePolicyEntriesWithContext(context.Background(), request)
}

// ResetRoutePolicyEntries
// This API is used to reset the designated route reception policy entry based on the rule ID and supports batch modification.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ResetRoutePolicyEntriesWithContext(ctx context.Context, request *ResetRoutePolicyEntriesRequest) (response *ResetRoutePolicyEntriesResponse, err error) {
    if request == nil {
        request = NewResetRoutePolicyEntriesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ResetRoutePolicyEntries")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetRoutePolicyEntries require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetRoutePolicyEntriesResponse()
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
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_VPCCIDRCONFLICT = "InvalidParameterValue.VpcCidrConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER_WITHGUESS = "UnknownParameter.WithGuess"
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
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_VPCCIDRCONFLICT = "InvalidParameterValue.VpcCidrConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER_WITHGUESS = "UnknownParameter.WithGuess"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ECMP = "UnsupportedOperation.Ecmp"
//  UNSUPPORTEDOPERATION_SYSTEMROUTE = "UnsupportedOperation.SystemRoute"
func (c *Client) ResetRoutesWithContext(ctx context.Context, request *ResetRoutesRequest) (response *ResetRoutesResponse, err error) {
    if request == nil {
        request = NewResetRoutesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ResetRoutes")
    
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
// The API is used to reset a VPN tunnel.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_VPNCONNINVALIDSTATE = "UnsupportedOperation.VpnConnInvalidState"
func (c *Client) ResetVpnConnection(request *ResetVpnConnectionRequest) (response *ResetVpnConnectionResponse, err error) {
    return c.ResetVpnConnectionWithContext(context.Background(), request)
}

// ResetVpnConnection
// The API is used to reset a VPN tunnel.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_VPNCONNINVALIDSTATE = "UnsupportedOperation.VpnConnInvalidState"
func (c *Client) ResetVpnConnectionWithContext(ctx context.Context, request *ResetVpnConnectionRequest) (response *ResetVpnConnectionResponse, err error) {
    if request == nil {
        request = NewResetVpnConnectionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ResetVpnConnection")
    
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
// This API is used to adjust the bandwidth cap of a VPN gateway. The adjustment of the VPN gateway bandwidth is limited to [5,100] Mbps and [200,1000] Mbps. 
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_VPNUNSUPPORTEDMODIFYBANDWIDTH = "UnsupportedOperation.VpnUnsupportedModifyBandwidth"
func (c *Client) ResetVpnGatewayInternetMaxBandwidth(request *ResetVpnGatewayInternetMaxBandwidthRequest) (response *ResetVpnGatewayInternetMaxBandwidthResponse, err error) {
    return c.ResetVpnGatewayInternetMaxBandwidthWithContext(context.Background(), request)
}

// ResetVpnGatewayInternetMaxBandwidth
// This API is used to adjust the bandwidth cap of a VPN gateway. The adjustment of the VPN gateway bandwidth is limited to [5,100] Mbps and [200,1000] Mbps. 
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_VPNUNSUPPORTEDMODIFYBANDWIDTH = "UnsupportedOperation.VpnUnsupportedModifyBandwidth"
func (c *Client) ResetVpnGatewayInternetMaxBandwidthWithContext(ctx context.Context, request *ResetVpnGatewayInternetMaxBandwidthRequest) (response *ResetVpnGatewayInternetMaxBandwidthResponse, err error) {
    if request == nil {
        request = NewResetVpnGatewayInternetMaxBandwidthRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ResetVpnGatewayInternetMaxBandwidth")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetVpnGatewayInternetMaxBandwidth require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetVpnGatewayInternetMaxBandwidthResponse()
    err = c.Send(request, response)
    return
}

func NewResumeSnapshotInstanceRequest() (request *ResumeSnapshotInstanceRequest) {
    request = &ResumeSnapshotInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "ResumeSnapshotInstance")
    
    
    return
}

func NewResumeSnapshotInstanceResponse() (response *ResumeSnapshotInstanceResponse) {
    response = &ResumeSnapshotInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResumeSnapshotInstance
// This API is used to restore security group policies with a backup.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ResumeSnapshotInstance(request *ResumeSnapshotInstanceRequest) (response *ResumeSnapshotInstanceResponse, err error) {
    return c.ResumeSnapshotInstanceWithContext(context.Background(), request)
}

// ResumeSnapshotInstance
// This API is used to restore security group policies with a backup.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ResumeSnapshotInstanceWithContext(ctx context.Context, request *ResumeSnapshotInstanceRequest) (response *ResumeSnapshotInstanceResponse, err error) {
    if request == nil {
        request = NewResumeSnapshotInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ResumeSnapshotInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResumeSnapshotInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewResumeSnapshotInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewReturnNormalAddressesRequest() (request *ReturnNormalAddressesRequest) {
    request = &ReturnNormalAddressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "ReturnNormalAddresses")
    
    
    return
}

func NewReturnNormalAddressesResponse() (response *ReturnNormalAddressesResponse) {
    response = &ReturnNormalAddressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReturnNormalAddresses
// This API is used to unbind and release public IPs. 
//
// Note: Starting from Dec 15, 2022, CAM authorization is required for a sub-account to call this API. For more details, see [Authorization Guide](https://intl.cloud.tencent.com/document/product/598/34545?from_cn_redirect=1).
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_ADDRESSIPSNOTFOUND = "InvalidParameterValue.AddressIpsNotFound"
//  INVALIDPARAMETERVALUE_ILLEGAL = "InvalidParameterValue.Illegal"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  UNSUPPORTEDOPERATION_NOTSUPPORTEDADDRESSIPSCHARGETYPE = "UnsupportedOperation.NotSupportedAddressIpsChargeType"
func (c *Client) ReturnNormalAddresses(request *ReturnNormalAddressesRequest) (response *ReturnNormalAddressesResponse, err error) {
    return c.ReturnNormalAddressesWithContext(context.Background(), request)
}

// ReturnNormalAddresses
// This API is used to unbind and release public IPs. 
//
// Note: Starting from Dec 15, 2022, CAM authorization is required for a sub-account to call this API. For more details, see [Authorization Guide](https://intl.cloud.tencent.com/document/product/598/34545?from_cn_redirect=1).
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_ADDRESSIPSNOTFOUND = "InvalidParameterValue.AddressIpsNotFound"
//  INVALIDPARAMETERVALUE_ILLEGAL = "InvalidParameterValue.Illegal"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  UNSUPPORTEDOPERATION_NOTSUPPORTEDADDRESSIPSCHARGETYPE = "UnsupportedOperation.NotSupportedAddressIpsChargeType"
func (c *Client) ReturnNormalAddressesWithContext(ctx context.Context, request *ReturnNormalAddressesRequest) (response *ReturnNormalAddressesResponse, err error) {
    if request == nil {
        request = NewReturnNormalAddressesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "ReturnNormalAddresses")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReturnNormalAddresses require credential")
    }

    request.SetContext(ctx)
    
    response = NewReturnNormalAddressesResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "SetCcnRegionBandwidthLimits")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetCcnRegionBandwidthLimits require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetCcnRegionBandwidthLimitsResponse()
    err = c.Send(request, response)
    return
}

func NewSetVpnGatewaysRenewFlagRequest() (request *SetVpnGatewaysRenewFlagRequest) {
    request = &SetVpnGatewaysRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vpc", APIVersion, "SetVpnGatewaysRenewFlag")
    
    
    return
}

func NewSetVpnGatewaysRenewFlagResponse() (response *SetVpnGatewaysRenewFlagResponse) {
    response = &SetVpnGatewaysRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetVpnGatewaysRenewFlag
// This API is used set the auto-renewal configuration of a VPN gateway.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SetVpnGatewaysRenewFlag(request *SetVpnGatewaysRenewFlagRequest) (response *SetVpnGatewaysRenewFlagResponse, err error) {
    return c.SetVpnGatewaysRenewFlagWithContext(context.Background(), request)
}

// SetVpnGatewaysRenewFlag
// This API is used set the auto-renewal configuration of a VPN gateway.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SetVpnGatewaysRenewFlagWithContext(ctx context.Context, request *SetVpnGatewaysRenewFlagRequest) (response *SetVpnGatewaysRenewFlagResponse, err error) {
    if request == nil {
        request = NewSetVpnGatewaysRenewFlagRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "SetVpnGatewaysRenewFlag")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetVpnGatewaysRenewFlag require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetVpnGatewaysRenewFlagResponse()
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
//  FAILEDOPERATION_TASKFAILED = "FailedOperation.TaskFailed"
//  INVALIDADDRESSID_BLOCKED = "InvalidAddressId.Blocked"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_ALREADYBINDEIP = "InvalidInstanceId.AlreadyBindEip"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE_INSTANCEHASNOWANIP = "InvalidParameterValue.InstanceHasNoWanIP"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INSTANCENOWANIP = "InvalidParameterValue.InstanceNoWanIP"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCESTATE = "InvalidParameterValue.InvalidInstanceState"
//  LIMITEXCEEDED_MONTHLYADDRESSRECOVERYQUOTA = "LimitExceeded.MonthlyAddressRecoveryQuota"
//  OPERATIONDENIED_ADDRESSINARREARS = "OperationDenied.AddressInArrears"
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
//  FAILEDOPERATION_TASKFAILED = "FailedOperation.TaskFailed"
//  INVALIDADDRESSID_BLOCKED = "InvalidAddressId.Blocked"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_ALREADYBINDEIP = "InvalidInstanceId.AlreadyBindEip"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE_INSTANCEHASNOWANIP = "InvalidParameterValue.InstanceHasNoWanIP"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INSTANCENOWANIP = "InvalidParameterValue.InstanceNoWanIP"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCESTATE = "InvalidParameterValue.InvalidInstanceState"
//  LIMITEXCEEDED_MONTHLYADDRESSRECOVERYQUOTA = "LimitExceeded.MonthlyAddressRecoveryQuota"
//  OPERATIONDENIED_ADDRESSINARREARS = "OperationDenied.AddressInArrears"
//  OPERATIONDENIED_MUTEXTASKRUNNING = "OperationDenied.MutexTaskRunning"
//  UNSUPPORTEDOPERATION_ADDRESSSTATUSNOTPERMIT = "UnsupportedOperation.AddressStatusNotPermit"
//  UNSUPPORTEDOPERATION_INVALIDADDRESSINTERNETCHARGETYPE = "UnsupportedOperation.InvalidAddressInternetChargeType"
func (c *Client) TransformAddressWithContext(ctx context.Context, request *TransformAddressRequest) (response *TransformAddressResponse, err error) {
    if request == nil {
        request = NewTransformAddressRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "TransformAddress")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "UnassignIpv6Addresses")
    
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
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "UnassignIpv6CidrBlock")
    
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
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
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
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UnassignIpv6SubnetCidrBlockWithContext(ctx context.Context, request *UnassignIpv6SubnetCidrBlockRequest) (response *UnassignIpv6SubnetCidrBlockResponse, err error) {
    if request == nil {
        request = NewUnassignIpv6SubnetCidrBlockRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "UnassignIpv6SubnetCidrBlock")
    
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
//  INVALIDPARAMETERVALUE_DUPLICATEPARA = "InvalidParameterValue.DuplicatePara"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED_ACTIONLIMITED = "LimitExceeded.ActionLimited"
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
//  INVALIDPARAMETERVALUE_DUPLICATEPARA = "InvalidParameterValue.DuplicatePara"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED_ACTIONLIMITED = "LimitExceeded.ActionLimited"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ATTACHMENTNOTFOUND = "UnsupportedOperation.AttachmentNotFound"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) UnassignPrivateIpAddressesWithContext(ctx context.Context, request *UnassignPrivateIpAddressesRequest) (response *UnassignPrivateIpAddressesResponse, err error) {
    if request == nil {
        request = NewUnassignPrivateIpAddressesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "UnassignPrivateIpAddresses")
    
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
//  RESOURCENOTFOUND = "ResourceNotFound"
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
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_NOTIFYCCN = "UnsupportedOperation.NotifyCcn"
//  UNSUPPORTEDOPERATION_SYSTEMROUTE = "UnsupportedOperation.SystemRoute"
func (c *Client) WithdrawNotifyRoutesWithContext(ctx context.Context, request *WithdrawNotifyRoutesRequest) (response *WithdrawNotifyRoutesResponse, err error) {
    if request == nil {
        request = NewWithdrawNotifyRoutesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vpc", APIVersion, "WithdrawNotifyRoutes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("WithdrawNotifyRoutes require credential")
    }

    request.SetContext(ctx)
    
    response = NewWithdrawNotifyRoutesResponse()
    err = c.Send(request, response)
    return
}
