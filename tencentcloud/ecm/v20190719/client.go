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

package v20190719

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2019-07-19"

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


func NewAllocateAddressesRequest() (request *AllocateAddressesRequest) {
    request = &AllocateAddressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "AllocateAddresses")
    
    
    return
}

func NewAllocateAddressesResponse() (response *AllocateAddressesResponse) {
    response = &AllocateAddressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AllocateAddresses
// This API is used to apply for one or multiple EIPs.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_INVALIDSTATUS = "FailedOperation.InvalidStatus"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  FAILEDOPERATION_PRIVATEIPADDRESSBINDED = "FailedOperation.PrivateIpAddressBinded"
//  FAILEDOPERATION_PRIVATEIPADDRESSUNAVAILABLE = "FailedOperation.PrivateIpAddressUnavailable"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_INVALIDPARAMETERCONFLICT = "InvalidParameter.InvalidParameterConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ADDRESSQUOTALIMITEXCEEDED = "InvalidParameterValue.AddressQuotaLimitExceeded"
//  INVALIDPARAMETERVALUE_INVALIDADDRESSCOUNT = "InvalidParameterValue.InvalidAddressCount"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTH = "InvalidParameterValue.InvalidBandwidth"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.AddressQuotaLimitExceeded"
//  LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDEDDAILYALLOCATE = "LimitExceeded.AddressQuotaLimitExceededDailyAllocate"
//  MISSINGPARAMETER_MISSINGPRIVATEIPADDRESS = "MissingParameter.MissingPrivateIpAddress"
//  RESOURCEINSUFFICIENT_IPQUOTANOTENOUGH = "ResourceInsufficient.IPQuotaNotEnough"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_INVALIDPRIVATEIPADDRESSALREADYBINDEIP = "UnsupportedOperation.InvalidPrivateIpAddressAlreadyBindEip"
func (c *Client) AllocateAddresses(request *AllocateAddressesRequest) (response *AllocateAddressesResponse, err error) {
    return c.AllocateAddressesWithContext(context.Background(), request)
}

// AllocateAddresses
// This API is used to apply for one or multiple EIPs.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_INVALIDSTATUS = "FailedOperation.InvalidStatus"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  FAILEDOPERATION_PRIVATEIPADDRESSBINDED = "FailedOperation.PrivateIpAddressBinded"
//  FAILEDOPERATION_PRIVATEIPADDRESSUNAVAILABLE = "FailedOperation.PrivateIpAddressUnavailable"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_INVALIDPARAMETERCONFLICT = "InvalidParameter.InvalidParameterConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ADDRESSQUOTALIMITEXCEEDED = "InvalidParameterValue.AddressQuotaLimitExceeded"
//  INVALIDPARAMETERVALUE_INVALIDADDRESSCOUNT = "InvalidParameterValue.InvalidAddressCount"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTH = "InvalidParameterValue.InvalidBandwidth"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.AddressQuotaLimitExceeded"
//  LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDEDDAILYALLOCATE = "LimitExceeded.AddressQuotaLimitExceededDailyAllocate"
//  MISSINGPARAMETER_MISSINGPRIVATEIPADDRESS = "MissingParameter.MissingPrivateIpAddress"
//  RESOURCEINSUFFICIENT_IPQUOTANOTENOUGH = "ResourceInsufficient.IPQuotaNotEnough"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_INVALIDPRIVATEIPADDRESSALREADYBINDEIP = "UnsupportedOperation.InvalidPrivateIpAddressAlreadyBindEip"
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
    
    request.Init().WithApiInfo("ecm", APIVersion, "AssignIpv6Addresses")
    
    
    return
}

func NewAssignIpv6AddressesResponse() (response *AssignIpv6AddressesResponse) {
    response = &AssignIpv6AddressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AssignIpv6Addresses
// This API is used to apply for an IPv6 address for an ENI.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDECMREGION = "InvalidParameterValue.InvalidEcmRegion"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NODENOTSUPPORTIPV6 = "InvalidParameterValue.NodeNotSupportIPv6"
//  INVALIDPARAMETERVALUE_SUBNETCONFLICT = "InvalidParameterValue.SubnetConflict"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDREGION = "InvalidParameterValue.UnsupportedRegion"
//  INVALIDPARAMETERVALUE_USERNOTSUPPORTIPV6 = "InvalidParameterValue.UserNotSupportIPv6"
//  LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.AddressQuotaLimitExceeded"
//  LIMITEXCEEDED_CIDRBLOCK = "LimitExceeded.CidrBlock"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) AssignIpv6Addresses(request *AssignIpv6AddressesRequest) (response *AssignIpv6AddressesResponse, err error) {
    return c.AssignIpv6AddressesWithContext(context.Background(), request)
}

// AssignIpv6Addresses
// This API is used to apply for an IPv6 address for an ENI.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDECMREGION = "InvalidParameterValue.InvalidEcmRegion"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_NODENOTSUPPORTIPV6 = "InvalidParameterValue.NodeNotSupportIPv6"
//  INVALIDPARAMETERVALUE_SUBNETCONFLICT = "InvalidParameterValue.SubnetConflict"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDREGION = "InvalidParameterValue.UnsupportedRegion"
//  INVALIDPARAMETERVALUE_USERNOTSUPPORTIPV6 = "InvalidParameterValue.UserNotSupportIPv6"
//  LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.AddressQuotaLimitExceeded"
//  LIMITEXCEEDED_CIDRBLOCK = "LimitExceeded.CidrBlock"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
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

func NewAssignPrivateIpAddressesRequest() (request *AssignPrivateIpAddressesRequest) {
    request = &AssignPrivateIpAddressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "AssignPrivateIpAddresses")
    
    
    return
}

func NewAssignPrivateIpAddressesResponse() (response *AssignPrivateIpAddressesResponse) {
    response = &AssignPrivateIpAddressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AssignPrivateIpAddresses
// This API is used to apply for a private IP for an ENI.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  LIMITEXCEEDED_PRIVATEIPQUOTALIMITEXCEEDED = "LimitExceeded.PrivateIPQuotaLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) AssignPrivateIpAddresses(request *AssignPrivateIpAddressesRequest) (response *AssignPrivateIpAddressesResponse, err error) {
    return c.AssignPrivateIpAddressesWithContext(context.Background(), request)
}

// AssignPrivateIpAddresses
// This API is used to apply for a private IP for an ENI.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  LIMITEXCEEDED_PRIVATEIPQUOTALIMITEXCEEDED = "LimitExceeded.PrivateIPQuotaLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
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
    
    request.Init().WithApiInfo("ecm", APIVersion, "AssociateAddress")
    
    
    return
}

func NewAssociateAddressResponse() (response *AssociateAddressResponse) {
    response = &AssociateAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AssociateAddress
// This API is used to bind an EIP to an instance or the specified private IP of an ENI.
//
// Binding an EIP to an instance (ECM) is essentially to bind it to the primary private IP of the primary ENI of the instance.
//
// When you bind an EIP to the private IP of the specified ENI, if the private IP is already bound to an EIP or public IP, a failure will be reported, and you must unbind it first before you can bind it to a new EIP.
//
// Only EIPs in `UNBIND` status can be bound to a private IP.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_INVALIDSTATUS = "FailedOperation.InvalidStatus"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  FAILEDOPERATION_PRIVATEIPADDRESSUNAVAILABLE = "FailedOperation.PrivateIpAddressUnavailable"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_INVALIDPARAMETERCONFLICT = "InvalidParameter.InvalidParameterConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVAILDENIID = "InvalidParameterValue.InvaildEniID"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.AddressQuotaLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGASSOCIATEENTITY = "MissingParameter.MissingAssociateEntity"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ADDRESSIDNOTFOUND = "UnsupportedOperation.AddressIdNotFound"
//  UNSUPPORTEDOPERATION_ALREADYBINDEIP = "UnsupportedOperation.AlreadyBindEip"
//  UNSUPPORTEDOPERATION_INSTANCEIDNOTFOUND = "UnsupportedOperation.InstanceIdNotFound"
//  UNSUPPORTEDOPERATION_INSTANCEIDNOTSUPPORTED = "UnsupportedOperation.InstanceIdNotSupported"
//  UNSUPPORTEDOPERATION_INVALIDNETWORKINTERFACEIDNOTFOUND = "UnsupportedOperation.InvalidNetworkInterfaceIdNotFound"
//  UNSUPPORTEDOPERATION_INVALIDPRIVATEIPADDRESSALREADYBINDEIP = "UnsupportedOperation.InvalidPrivateIpAddressAlreadyBindEip"
//  UNSUPPORTEDOPERATION_MALFORMED = "UnsupportedOperation.Malformed"
//  UNSUPPORTEDOPERATION_STATUSNOTPERMIT = "UnsupportedOperation.StatusNotPermit"
func (c *Client) AssociateAddress(request *AssociateAddressRequest) (response *AssociateAddressResponse, err error) {
    return c.AssociateAddressWithContext(context.Background(), request)
}

// AssociateAddress
// This API is used to bind an EIP to an instance or the specified private IP of an ENI.
//
// Binding an EIP to an instance (ECM) is essentially to bind it to the primary private IP of the primary ENI of the instance.
//
// When you bind an EIP to the private IP of the specified ENI, if the private IP is already bound to an EIP or public IP, a failure will be reported, and you must unbind it first before you can bind it to a new EIP.
//
// Only EIPs in `UNBIND` status can be bound to a private IP.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_INVALIDSTATUS = "FailedOperation.InvalidStatus"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  FAILEDOPERATION_PRIVATEIPADDRESSUNAVAILABLE = "FailedOperation.PrivateIpAddressUnavailable"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_INVALIDPARAMETERCONFLICT = "InvalidParameter.InvalidParameterConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVAILDENIID = "InvalidParameterValue.InvaildEniID"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.AddressQuotaLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGASSOCIATEENTITY = "MissingParameter.MissingAssociateEntity"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ADDRESSIDNOTFOUND = "UnsupportedOperation.AddressIdNotFound"
//  UNSUPPORTEDOPERATION_ALREADYBINDEIP = "UnsupportedOperation.AlreadyBindEip"
//  UNSUPPORTEDOPERATION_INSTANCEIDNOTFOUND = "UnsupportedOperation.InstanceIdNotFound"
//  UNSUPPORTEDOPERATION_INSTANCEIDNOTSUPPORTED = "UnsupportedOperation.InstanceIdNotSupported"
//  UNSUPPORTEDOPERATION_INVALIDNETWORKINTERFACEIDNOTFOUND = "UnsupportedOperation.InvalidNetworkInterfaceIdNotFound"
//  UNSUPPORTEDOPERATION_INVALIDPRIVATEIPADDRESSALREADYBINDEIP = "UnsupportedOperation.InvalidPrivateIpAddressAlreadyBindEip"
//  UNSUPPORTEDOPERATION_MALFORMED = "UnsupportedOperation.Malformed"
//  UNSUPPORTEDOPERATION_STATUSNOTPERMIT = "UnsupportedOperation.StatusNotPermit"
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

func NewAssociateSecurityGroupsRequest() (request *AssociateSecurityGroupsRequest) {
    request = &AssociateSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "AssociateSecurityGroups")
    
    
    return
}

func NewAssociateSecurityGroupsResponse() (response *AssociateSecurityGroupsResponse) {
    response = &AssociateSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AssociateSecurityGroups
// This API is used to bind a security group.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  LIMITEXCEEDED_SECURITYGROUPINSTANCELIMITEXCEEDED = "LimitExceeded.SecurityGroupInstanceLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) AssociateSecurityGroups(request *AssociateSecurityGroupsRequest) (response *AssociateSecurityGroupsResponse, err error) {
    return c.AssociateSecurityGroupsWithContext(context.Background(), request)
}

// AssociateSecurityGroups
// This API is used to bind a security group.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  LIMITEXCEEDED_SECURITYGROUPINSTANCELIMITEXCEEDED = "LimitExceeded.SecurityGroupInstanceLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) AssociateSecurityGroupsWithContext(ctx context.Context, request *AssociateSecurityGroupsRequest) (response *AssociateSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewAssociateSecurityGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssociateSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewAssociateSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewAttachNetworkInterfaceRequest() (request *AttachNetworkInterfaceRequest) {
    request = &AttachNetworkInterfaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "AttachNetworkInterface")
    
    
    return
}

func NewAttachNetworkInterfaceResponse() (response *AttachNetworkInterfaceResponse) {
    response = &AttachNetworkInterfaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AttachNetworkInterface
// This API is used to bind an ENI to a CVM instance.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVAILDENIID = "InvalidParameterValue.InvaildEniID"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.AddressQuotaLimitExceeded"
//  LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDEDDAILYALLOCATE = "LimitExceeded.AddressQuotaLimitExceededDailyAllocate"
//  LIMITEXCEEDED_ENIQUOTALIMITEXCEEDED = "LimitExceeded.EniQuotaLimitExceeded"
//  LIMITEXCEEDED_PRIVATEIPQUOTALIMITEXCEEDED = "LimitExceeded.PrivateIPQuotaLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ATTACHMENTALREADYEXISTS = "UnsupportedOperation.AttachmentAlreadyExists"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
func (c *Client) AttachNetworkInterface(request *AttachNetworkInterfaceRequest) (response *AttachNetworkInterfaceResponse, err error) {
    return c.AttachNetworkInterfaceWithContext(context.Background(), request)
}

// AttachNetworkInterface
// This API is used to bind an ENI to a CVM instance.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVAILDENIID = "InvalidParameterValue.InvaildEniID"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.AddressQuotaLimitExceeded"
//  LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDEDDAILYALLOCATE = "LimitExceeded.AddressQuotaLimitExceededDailyAllocate"
//  LIMITEXCEEDED_ENIQUOTALIMITEXCEEDED = "LimitExceeded.EniQuotaLimitExceeded"
//  LIMITEXCEEDED_PRIVATEIPQUOTALIMITEXCEEDED = "LimitExceeded.PrivateIPQuotaLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ATTACHMENTALREADYEXISTS = "UnsupportedOperation.AttachmentAlreadyExists"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
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

func NewBatchDeregisterTargetsRequest() (request *BatchDeregisterTargetsRequest) {
    request = &BatchDeregisterTargetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "BatchDeregisterTargets")
    
    
    return
}

func NewBatchDeregisterTargetsResponse() (response *BatchDeregisterTargetsResponse) {
    response = &BatchDeregisterTargetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchDeregisterTargets
// This API is used to batch unbind real servers.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_LISTENERIDNOTFOUND = "InvalidParameter.ListenerIdNotFound"
//  INVALIDPARAMETER_LOCATIONNOTFOUND = "InvalidParameter.LocationNotFound"
//  INVALIDPARAMETER_PORTCHECKFAILED = "InvalidParameter.PortCheckFailed"
//  INVALIDPARAMETER_PROTOCOLCHECKFAILED = "InvalidParameter.ProtocolCheckFailed"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) BatchDeregisterTargets(request *BatchDeregisterTargetsRequest) (response *BatchDeregisterTargetsResponse, err error) {
    return c.BatchDeregisterTargetsWithContext(context.Background(), request)
}

// BatchDeregisterTargets
// This API is used to batch unbind real servers.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_LISTENERIDNOTFOUND = "InvalidParameter.ListenerIdNotFound"
//  INVALIDPARAMETER_LOCATIONNOTFOUND = "InvalidParameter.LocationNotFound"
//  INVALIDPARAMETER_PORTCHECKFAILED = "InvalidParameter.PortCheckFailed"
//  INVALIDPARAMETER_PROTOCOLCHECKFAILED = "InvalidParameter.ProtocolCheckFailed"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
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
    
    request.Init().WithApiInfo("ecm", APIVersion, "BatchModifyTargetWeight")
    
    
    return
}

func NewBatchModifyTargetWeightResponse() (response *BatchModifyTargetWeightResponse) {
    response = &BatchModifyTargetWeightResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchModifyTargetWeight
// This API is used to batch modify the forwarding weights of the real servers bound to a listener.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) BatchModifyTargetWeight(request *BatchModifyTargetWeightRequest) (response *BatchModifyTargetWeightResponse, err error) {
    return c.BatchModifyTargetWeightWithContext(context.Background(), request)
}

// BatchModifyTargetWeight
// This API is used to batch modify the forwarding weights of the real servers bound to a listener.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
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
    
    request.Init().WithApiInfo("ecm", APIVersion, "BatchRegisterTargets")
    
    
    return
}

func NewBatchRegisterTargetsResponse() (response *BatchRegisterTargetsResponse) {
    response = &BatchRegisterTargetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchRegisterTargets
// This API is used to batch bind backend targets.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_LISTENERIDNOTFOUND = "InvalidParameter.ListenerIdNotFound"
//  INVALIDPARAMETER_LOCATIONNOTFOUND = "InvalidParameter.LocationNotFound"
//  INVALIDPARAMETER_PORTCHECKFAILED = "InvalidParameter.PortCheckFailed"
//  INVALIDPARAMETER_PROTOCOLCHECKFAILED = "InvalidParameter.ProtocolCheckFailed"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) BatchRegisterTargets(request *BatchRegisterTargetsRequest) (response *BatchRegisterTargetsResponse, err error) {
    return c.BatchRegisterTargetsWithContext(context.Background(), request)
}

// BatchRegisterTargets
// This API is used to batch bind backend targets.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_LISTENERIDNOTFOUND = "InvalidParameter.ListenerIdNotFound"
//  INVALIDPARAMETER_LOCATIONNOTFOUND = "InvalidParameter.LocationNotFound"
//  INVALIDPARAMETER_PORTCHECKFAILED = "InvalidParameter.PortCheckFailed"
//  INVALIDPARAMETER_PROTOCOLCHECKFAILED = "InvalidParameter.ProtocolCheckFailed"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
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

func NewCreateHaVipRequest() (request *CreateHaVipRequest) {
    request = &CreateHaVipRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "CreateHaVip")
    
    
    return
}

func NewCreateHaVipResponse() (response *CreateHaVipResponse) {
    response = &CreateHaVipResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateHaVip
// This API is used to create an HAVIP.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_OBJECTNOTCURRENTSUBNET = "InvalidParameterValue.ObjectNotCurrentSubnet"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateHaVip(request *CreateHaVipRequest) (response *CreateHaVipResponse, err error) {
    return c.CreateHaVipWithContext(context.Background(), request)
}

// CreateHaVip
// This API is used to create an HAVIP.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_OBJECTNOTCURRENTSUBNET = "InvalidParameterValue.ObjectNotCurrentSubnet"
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

func NewCreateImageRequest() (request *CreateImageRequest) {
    request = &CreateImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "CreateImage")
    
    
    return
}

func NewCreateImageResponse() (response *CreateImageResponse) {
    response = &CreateImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateImage
// This API is used to create an image with the system disk of an instance. The image can be used to create instances.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCENOTALLSTOPPED = "FailedOperation.InstanceNotAllStopped"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_IMAGENAMEDUPLICATE = "InvalidParameterValue.ImageNameDuplicate"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGIMAGEPARAMETER = "MissingParameter.MissingImageParameter"
//  RESOURCEINSUFFICIENT_INVAILDPRIVATEIMAGENUM = "ResourceInsufficient.InvaildPrivateImageNum"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INSTANCEIDNOTFOUND = "UnsupportedOperation.InstanceIdNotFound"
//  UNSUPPORTEDOPERATION_INSTANCEIDNOTSUPPORTED = "UnsupportedOperation.InstanceIdNotSupported"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
func (c *Client) CreateImage(request *CreateImageRequest) (response *CreateImageResponse, err error) {
    return c.CreateImageWithContext(context.Background(), request)
}

// CreateImage
// This API is used to create an image with the system disk of an instance. The image can be used to create instances.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCENOTALLSTOPPED = "FailedOperation.InstanceNotAllStopped"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_IMAGENAMEDUPLICATE = "InvalidParameterValue.ImageNameDuplicate"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGIMAGEPARAMETER = "MissingParameter.MissingImageParameter"
//  RESOURCEINSUFFICIENT_INVAILDPRIVATEIMAGENUM = "ResourceInsufficient.InvaildPrivateImageNum"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INSTANCEIDNOTFOUND = "UnsupportedOperation.InstanceIdNotFound"
//  UNSUPPORTEDOPERATION_INSTANCEIDNOTSUPPORTED = "UnsupportedOperation.InstanceIdNotSupported"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
func (c *Client) CreateImageWithContext(ctx context.Context, request *CreateImageRequest) (response *CreateImageResponse, err error) {
    if request == nil {
        request = NewCreateImageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateImage require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateImageResponse()
    err = c.Send(request, response)
    return
}

func NewCreateKeyPairRequest() (request *CreateKeyPairRequest) {
    request = &CreateKeyPairRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "CreateKeyPair")
    
    
    return
}

func NewCreateKeyPairResponse() (response *CreateKeyPairResponse) {
    response = &CreateKeyPairResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateKeyPair
// This API is used to create an `OpenSSH RSA` key pair, which can be used to log in to a Linux instance.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDKEYPAIRNAME = "InvalidParameterValue.InvalidKeyPairName"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateKeyPair(request *CreateKeyPairRequest) (response *CreateKeyPairResponse, err error) {
    return c.CreateKeyPairWithContext(context.Background(), request)
}

// CreateKeyPair
// This API is used to create an `OpenSSH RSA` key pair, which can be used to log in to a Linux instance.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDKEYPAIRNAME = "InvalidParameterValue.InvalidKeyPairName"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateKeyPairWithContext(ctx context.Context, request *CreateKeyPairRequest) (response *CreateKeyPairResponse, err error) {
    if request == nil {
        request = NewCreateKeyPairRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateKeyPair require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateKeyPairResponse()
    err = c.Send(request, response)
    return
}

func NewCreateListenerRequest() (request *CreateListenerRequest) {
    request = &CreateListenerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "CreateListener")
    
    
    return
}

func NewCreateListenerResponse() (response *CreateListenerResponse) {
    response = &CreateListenerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateListener
// This API is used to create a CLB listener.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) CreateListener(request *CreateListenerRequest) (response *CreateListenerResponse, err error) {
    return c.CreateListenerWithContext(context.Background(), request)
}

// CreateListener
// This API is used to create a CLB listener.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
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
    
    request.Init().WithApiInfo("ecm", APIVersion, "CreateLoadBalancer")
    
    
    return
}

func NewCreateLoadBalancerResponse() (response *CreateLoadBalancerResponse) {
    response = &CreateLoadBalancerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLoadBalancer
// This API is used to purchase a CLB instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDECMREGION = "InvalidParameterValue.InvalidEcmRegion"
//  INVALIDPARAMETERVALUE_INVALIDISPINNODE = "InvalidParameterValue.InvalidISPInNode"
//  INVALIDPARAMETERVALUE_INVALIDLOADBALANCERNUM = "InvalidParameterValue.InvalidLoadBalancerNum"
//  INVALIDPARAMETERVALUE_INVALIDLOADBALANCERTYPE = "InvalidParameterValue.InvalidLoadBalancerType"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDSUBNETINZONE = "InvalidParameterValue.InvalidSubnetInZone"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TAGNUMOUTOFRANGE = "InvalidParameterValue.TagNumOutOfRange"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_LBLIMITEXCEEDED = "LimitExceeded.LBLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCESSOLDOUT_LOADBALANCERSOLDOUT = "ResourcesSoldOut.LoadBalancerSoldOut"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) CreateLoadBalancer(request *CreateLoadBalancerRequest) (response *CreateLoadBalancerResponse, err error) {
    return c.CreateLoadBalancerWithContext(context.Background(), request)
}

// CreateLoadBalancer
// This API is used to purchase a CLB instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDECMREGION = "InvalidParameterValue.InvalidEcmRegion"
//  INVALIDPARAMETERVALUE_INVALIDISPINNODE = "InvalidParameterValue.InvalidISPInNode"
//  INVALIDPARAMETERVALUE_INVALIDLOADBALANCERNUM = "InvalidParameterValue.InvalidLoadBalancerNum"
//  INVALIDPARAMETERVALUE_INVALIDLOADBALANCERTYPE = "InvalidParameterValue.InvalidLoadBalancerType"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDSUBNETINZONE = "InvalidParameterValue.InvalidSubnetInZone"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TAGNUMOUTOFRANGE = "InvalidParameterValue.TagNumOutOfRange"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_LBLIMITEXCEEDED = "LimitExceeded.LBLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCESSOLDOUT_LOADBALANCERSOLDOUT = "ResourcesSoldOut.LoadBalancerSoldOut"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
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

func NewCreateModuleRequest() (request *CreateModuleRequest) {
    request = &CreateModuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "CreateModule")
    
    
    return
}

func NewCreateModuleResponse() (response *CreateModuleResponse) {
    response = &CreateModuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateModule
// This API is used to create a module.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVAILDMODULENUM = "InvalidParameterValue.InvaildModuleNum"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTH = "InvalidParameterValue.InvalidBandwidth"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTHIN = "InvalidParameterValue.InvalidBandwidthIn"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTHINANDOUT = "InvalidParameterValue.InvalidBandwidthInAndOut"
//  INVALIDPARAMETERVALUE_INVALIDDATADISKSIZE = "InvalidParameterValue.InvalidDataDiskSize"
//  INVALIDPARAMETERVALUE_INVALIDEIPDIRECTSERVICE = "InvalidParameterValue.InvalidEIPDirectService"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEARCHITECTURE = "InvalidParameterValue.InvalidImageArchitecture"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageID"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCETYPECONFIGID = "InvalidParameterValue.InvalidInstanceTypeConfigID"
//  INVALIDPARAMETERVALUE_INVALIDMODULENAME = "InvalidParameterValue.InvalidModuleName"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDSYSTEMDISKSIZE = "InvalidParameterValue.InvalidSystemDiskSize"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  INVALIDPARAMETERVALUE_USERNOTSUPPORTIPV6 = "InvalidParameterValue.UserNotSupportIPv6"
//  LIMITEXCEEDED_MODULESECURITYGROUPLIMITEXCEEDED = "LimitExceeded.ModuleSecurityGroupLimitExceeded"
//  LIMITEXCEEDED_SECURITYGROUPMODULELIMITEXCEEDED = "LimitExceeded.SecurityGroupModuleLimitExceeded"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_INSTANCETYPENOTSUPPORTIMAGE = "UnsupportedOperation.InstanceTypeNotSupportImage"
func (c *Client) CreateModule(request *CreateModuleRequest) (response *CreateModuleResponse, err error) {
    return c.CreateModuleWithContext(context.Background(), request)
}

// CreateModule
// This API is used to create a module.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVAILDMODULENUM = "InvalidParameterValue.InvaildModuleNum"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTH = "InvalidParameterValue.InvalidBandwidth"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTHIN = "InvalidParameterValue.InvalidBandwidthIn"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTHINANDOUT = "InvalidParameterValue.InvalidBandwidthInAndOut"
//  INVALIDPARAMETERVALUE_INVALIDDATADISKSIZE = "InvalidParameterValue.InvalidDataDiskSize"
//  INVALIDPARAMETERVALUE_INVALIDEIPDIRECTSERVICE = "InvalidParameterValue.InvalidEIPDirectService"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEARCHITECTURE = "InvalidParameterValue.InvalidImageArchitecture"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageID"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCETYPECONFIGID = "InvalidParameterValue.InvalidInstanceTypeConfigID"
//  INVALIDPARAMETERVALUE_INVALIDMODULENAME = "InvalidParameterValue.InvalidModuleName"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDSYSTEMDISKSIZE = "InvalidParameterValue.InvalidSystemDiskSize"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  INVALIDPARAMETERVALUE_USERNOTSUPPORTIPV6 = "InvalidParameterValue.UserNotSupportIPv6"
//  LIMITEXCEEDED_MODULESECURITYGROUPLIMITEXCEEDED = "LimitExceeded.ModuleSecurityGroupLimitExceeded"
//  LIMITEXCEEDED_SECURITYGROUPMODULELIMITEXCEEDED = "LimitExceeded.SecurityGroupModuleLimitExceeded"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_INSTANCETYPENOTSUPPORTIMAGE = "UnsupportedOperation.InstanceTypeNotSupportImage"
func (c *Client) CreateModuleWithContext(ctx context.Context, request *CreateModuleRequest) (response *CreateModuleResponse, err error) {
    if request == nil {
        request = NewCreateModuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateModule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateModuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNetworkInterfaceRequest() (request *CreateNetworkInterfaceRequest) {
    request = &CreateNetworkInterfaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "CreateNetworkInterface")
    
    
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
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_RESERVED = "InvalidParameterValue.Reserved"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_ENIQUOTALIMITEXCEEDED = "LimitExceeded.EniQuotaLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGNETWORKINTERFACEPARAMETER = "MissingParameter.MissingNetworkInterfaceParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) CreateNetworkInterface(request *CreateNetworkInterfaceRequest) (response *CreateNetworkInterfaceResponse, err error) {
    return c.CreateNetworkInterfaceWithContext(context.Background(), request)
}

// CreateNetworkInterface
// This API is used to create an ENI.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_RESERVED = "InvalidParameterValue.Reserved"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_ENIQUOTALIMITEXCEEDED = "LimitExceeded.EniQuotaLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGNETWORKINTERFACEPARAMETER = "MissingParameter.MissingNetworkInterfaceParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
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
    
    request.Init().WithApiInfo("ecm", APIVersion, "CreateRouteTable")
    
    
    return
}

func NewCreateRouteTableResponse() (response *CreateRouteTableResponse) {
    response = &CreateRouteTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRouteTable
// After a VPC is created, the system will create a default route table, with which all new subnets will be associated. By default, you can use the default route table to manage your routing policies. If you have multiple routing policies, you can call the API for route table creation to create more route tables to manage your routing policies.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) CreateRouteTable(request *CreateRouteTableRequest) (response *CreateRouteTableResponse, err error) {
    return c.CreateRouteTableWithContext(context.Background(), request)
}

// CreateRouteTable
// After a VPC is created, the system will create a default route table, with which all new subnets will be associated. By default, you can use the default route table to manage your routing policies. If you have multiple routing policies, you can call the API for route table creation to create more route tables to manage your routing policies.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
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
    
    request.Init().WithApiInfo("ecm", APIVersion, "CreateRoutes")
    
    
    return
}

func NewCreateRoutesResponse() (response *CreateRoutesResponse) {
    response = &CreateRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRoutes
// This API is used to create a routing policy.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CIDRNOTINPEERVPC = "InvalidParameterValue.CidrNotInPeerVpc"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_VPCCIDRCONFLICT = "InvalidParameterValue.VpcCidrConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ECMP = "UnsupportedOperation.Ecmp"
//  UNSUPPORTEDOPERATION_ECMPWITHCCNROUTE = "UnsupportedOperation.EcmpWithCcnRoute"
//  UNSUPPORTEDOPERATION_ECMPWITHUSERROUTE = "UnsupportedOperation.EcmpWithUserRoute"
//  UNSUPPORTEDOPERATION_SYSTEMROUTE = "UnsupportedOperation.SystemRoute"
func (c *Client) CreateRoutes(request *CreateRoutesRequest) (response *CreateRoutesResponse, err error) {
    return c.CreateRoutesWithContext(context.Background(), request)
}

// CreateRoutes
// This API is used to create a routing policy.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CIDRNOTINPEERVPC = "InvalidParameterValue.CidrNotInPeerVpc"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_VPCCIDRCONFLICT = "InvalidParameterValue.VpcCidrConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ECMP = "UnsupportedOperation.Ecmp"
//  UNSUPPORTEDOPERATION_ECMPWITHCCNROUTE = "UnsupportedOperation.EcmpWithCcnRoute"
//  UNSUPPORTEDOPERATION_ECMPWITHUSERROUTE = "UnsupportedOperation.EcmpWithUserRoute"
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
    
    request.Init().WithApiInfo("ecm", APIVersion, "CreateSecurityGroup")
    
    
    return
}

func NewCreateSecurityGroupResponse() (response *CreateSecurityGroupResponse) {
    response = &CreateSecurityGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSecurityGroup
// This API is used to create a security group.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) CreateSecurityGroup(request *CreateSecurityGroupRequest) (response *CreateSecurityGroupResponse, err error) {
    return c.CreateSecurityGroupWithContext(context.Background(), request)
}

// CreateSecurityGroup
// This API is used to create a security group.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
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
    
    request.Init().WithApiInfo("ecm", APIVersion, "CreateSecurityGroupPolicies")
    
    
    return
}

func NewCreateSecurityGroupPoliciesResponse() (response *CreateSecurityGroupPoliciesResponse) {
    response = &CreateSecurityGroupPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSecurityGroupPolicies
// <p>This API is used to create a security group policy.</p>
//
// <p>In the `SecurityGroupPolicySet` parameter:</p>
//
// <ul>
//
// <li>`Version`: the version number of a security group policy, which automatically increases by one each time you update the security policy, to prevent expiration of the updated routing policies. If it is left empty, any conflicts will be ignored.</li>
//
// <li>When creating the `Egress` and `Ingress` polices,<ul>
//
// <li><code>Protocol</code>: <code>TCP</code>, <code>UDP</code>, <code>ICMP</code>, <code>GRE</code>, or <code>ALL</code>.</li>
//
// <li>`CidrBlock`: a CIDR block in the correct format. In a classic network, if a `CidrBlock` contains private IPs on Tencent Cloud for devices under your account other than CVMs, it does not mean this policy allows you to access these devices. The network isolation policies between tenants take priority over the private network policies in security groups.</li>
//
// <li>`SecurityGroupId`: ID of the security group. It can be the ID of security group to be modified, or the ID of other security group in the same project. All private IPs of all CVMs under the security group will be covered. If this field is used, the policy will automatically change according to the CVM associated with the group ID while being used to match network messages. You don’t need to change it manually.</li>
//
// <li>`Port`: a single port number such as 80, or a port range in the format of “8000-8010”. You may use this field only if the `Protocol` field takes the value `TCP` or `UDP`. Otherwise `Protocol` and `Port` are mutually exclusive.</li>
//
// <li>`Action`: only allows `ACCEPT` or `DROP`.</li>
//
// <li>`CidrBlock`, `SecurityGroupId`, and `AddressTemplate` are mutually exclusive. `Protocol` + `Port` and `ServiceTemplate` are mutually exclusive.</li>
//
// <li>You can only create policies in one direction in each request. To specify the `PolicyIndex` parameter, use the same index number in policies.</li>
//
// </ul></li></ul>
//
// <p>Default API request rate limit: 20 requests/sec.</p>
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_SECURITYGROUPPOLICYSET = "LimitExceeded.SecurityGroupPolicySet"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_DUPLICATEPOLICY = "UnsupportedOperation.DuplicatePolicy"
//  UNSUPPORTEDOPERATION_VERSIONMISMATCH = "UnsupportedOperation.VersionMismatch"
func (c *Client) CreateSecurityGroupPolicies(request *CreateSecurityGroupPoliciesRequest) (response *CreateSecurityGroupPoliciesResponse, err error) {
    return c.CreateSecurityGroupPoliciesWithContext(context.Background(), request)
}

// CreateSecurityGroupPolicies
// <p>This API is used to create a security group policy.</p>
//
// <p>In the `SecurityGroupPolicySet` parameter:</p>
//
// <ul>
//
// <li>`Version`: the version number of a security group policy, which automatically increases by one each time you update the security policy, to prevent expiration of the updated routing policies. If it is left empty, any conflicts will be ignored.</li>
//
// <li>When creating the `Egress` and `Ingress` polices,<ul>
//
// <li><code>Protocol</code>: <code>TCP</code>, <code>UDP</code>, <code>ICMP</code>, <code>GRE</code>, or <code>ALL</code>.</li>
//
// <li>`CidrBlock`: a CIDR block in the correct format. In a classic network, if a `CidrBlock` contains private IPs on Tencent Cloud for devices under your account other than CVMs, it does not mean this policy allows you to access these devices. The network isolation policies between tenants take priority over the private network policies in security groups.</li>
//
// <li>`SecurityGroupId`: ID of the security group. It can be the ID of security group to be modified, or the ID of other security group in the same project. All private IPs of all CVMs under the security group will be covered. If this field is used, the policy will automatically change according to the CVM associated with the group ID while being used to match network messages. You don’t need to change it manually.</li>
//
// <li>`Port`: a single port number such as 80, or a port range in the format of “8000-8010”. You may use this field only if the `Protocol` field takes the value `TCP` or `UDP`. Otherwise `Protocol` and `Port` are mutually exclusive.</li>
//
// <li>`Action`: only allows `ACCEPT` or `DROP`.</li>
//
// <li>`CidrBlock`, `SecurityGroupId`, and `AddressTemplate` are mutually exclusive. `Protocol` + `Port` and `ServiceTemplate` are mutually exclusive.</li>
//
// <li>You can only create policies in one direction in each request. To specify the `PolicyIndex` parameter, use the same index number in policies.</li>
//
// </ul></li></ul>
//
// <p>Default API request rate limit: 20 requests/sec.</p>
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_SECURITYGROUPPOLICYSET = "LimitExceeded.SecurityGroupPolicySet"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
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

func NewCreateSubnetRequest() (request *CreateSubnetRequest) {
    request = &CreateSubnetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "CreateSubnet")
    
    
    return
}

func NewCreateSubnetResponse() (response *CreateSubnetResponse) {
    response = &CreateSubnetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSubnet
// This API is used to create a subnet. After the subnet is created successfully, it will become the default subnet for the AZ.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDZONE = "InvalidParameterValue.InvalidZone"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_SUBNETCONFLICT = "InvalidParameterValue.SubnetConflict"
//  INVALIDPARAMETERVALUE_SUBNETRANGE = "InvalidParameterValue.SubnetRange"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
func (c *Client) CreateSubnet(request *CreateSubnetRequest) (response *CreateSubnetResponse, err error) {
    return c.CreateSubnetWithContext(context.Background(), request)
}

// CreateSubnet
// This API is used to create a subnet. After the subnet is created successfully, it will become the default subnet for the AZ.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDZONE = "InvalidParameterValue.InvalidZone"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_SUBNETCONFLICT = "InvalidParameterValue.SubnetConflict"
//  INVALIDPARAMETERVALUE_SUBNETRANGE = "InvalidParameterValue.SubnetRange"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
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

func NewCreateVpcRequest() (request *CreateVpcRequest) {
    request = &CreateVpcRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "CreateVpc")
    
    
    return
}

func NewCreateVpcResponse() (response *CreateVpcResponse) {
    response = &CreateVpcResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateVpc
// This API is used to create a VPC.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) CreateVpc(request *CreateVpcRequest) (response *CreateVpcResponse, err error) {
    return c.CreateVpcWithContext(context.Background(), request)
}

// CreateVpc
// This API is used to create a VPC.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
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

func NewDeleteHaVipRequest() (request *DeleteHaVipRequest) {
    request = &DeleteHaVipRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DeleteHaVip")
    
    
    return
}

func NewDeleteHaVipResponse() (response *DeleteHaVipResponse) {
    response = &DeleteHaVipResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteHaVip
// This API is used to delete an HAVIP.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteHaVip(request *DeleteHaVipRequest) (response *DeleteHaVipResponse, err error) {
    return c.DeleteHaVipWithContext(context.Background(), request)
}

// DeleteHaVip
// This API is used to delete an HAVIP.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
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

func NewDeleteImageRequest() (request *DeleteImageRequest) {
    request = &DeleteImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DeleteImage")
    
    
    return
}

func NewDeleteImageResponse() (response *DeleteImageResponse) {
    response = &DeleteImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteImage
// This API is used to delete an image.
//
// error code that may be returned:
//  FAILEDOPERATION_IMAGEINUSE = "FailedOperation.ImageInUse"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageID"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DeleteImage(request *DeleteImageRequest) (response *DeleteImageResponse, err error) {
    return c.DeleteImageWithContext(context.Background(), request)
}

// DeleteImage
// This API is used to delete an image.
//
// error code that may be returned:
//  FAILEDOPERATION_IMAGEINUSE = "FailedOperation.ImageInUse"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageID"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DeleteImageWithContext(ctx context.Context, request *DeleteImageRequest) (response *DeleteImageResponse, err error) {
    if request == nil {
        request = NewDeleteImageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteImage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteImageResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteListenerRequest() (request *DeleteListenerRequest) {
    request = &DeleteListenerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DeleteListener")
    
    
    return
}

func NewDeleteListenerResponse() (response *DeleteListenerResponse) {
    response = &DeleteListenerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteListener
// This API is used to delete a CLB listener.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DeleteListener(request *DeleteListenerRequest) (response *DeleteListenerResponse, err error) {
    return c.DeleteListenerWithContext(context.Background(), request)
}

// DeleteListener
// This API is used to delete a CLB listener.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
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
    
    request.Init().WithApiInfo("ecm", APIVersion, "DeleteLoadBalancer")
    
    
    return
}

func NewDeleteLoadBalancerResponse() (response *DeleteLoadBalancerResponse) {
    response = &DeleteLoadBalancerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLoadBalancer
// This API is used to delete a CLB instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DeleteLoadBalancer(request *DeleteLoadBalancerRequest) (response *DeleteLoadBalancerResponse, err error) {
    return c.DeleteLoadBalancerWithContext(context.Background(), request)
}

// DeleteLoadBalancer
// This API is used to delete a CLB instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
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
    
    request.Init().WithApiInfo("ecm", APIVersion, "DeleteLoadBalancerListeners")
    
    
    return
}

func NewDeleteLoadBalancerListenersResponse() (response *DeleteLoadBalancerListenersResponse) {
    response = &DeleteLoadBalancerListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLoadBalancerListeners
// This API is used to delete multiple CLB listeners.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DeleteLoadBalancerListeners(request *DeleteLoadBalancerListenersRequest) (response *DeleteLoadBalancerListenersResponse, err error) {
    return c.DeleteLoadBalancerListenersWithContext(context.Background(), request)
}

// DeleteLoadBalancerListeners
// This API is used to delete multiple CLB listeners.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
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

func NewDeleteModuleRequest() (request *DeleteModuleRequest) {
    request = &DeleteModuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DeleteModule")
    
    
    return
}

func NewDeleteModuleResponse() (response *DeleteModuleResponse) {
    response = &DeleteModuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteModule
// This API is used to delete a business module.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCEINMODULE = "FailedOperation.InstanceInModule"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONCONFLICT = "FailedOperation.OperationConflict"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DeleteModule(request *DeleteModuleRequest) (response *DeleteModuleResponse, err error) {
    return c.DeleteModuleWithContext(context.Background(), request)
}

// DeleteModule
// This API is used to delete a business module.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCEINMODULE = "FailedOperation.InstanceInModule"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONCONFLICT = "FailedOperation.OperationConflict"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DeleteModuleWithContext(ctx context.Context, request *DeleteModuleRequest) (response *DeleteModuleResponse, err error) {
    if request == nil {
        request = NewDeleteModuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteModule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteModuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNetworkInterfaceRequest() (request *DeleteNetworkInterfaceRequest) {
    request = &DeleteNetworkInterfaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DeleteNetworkInterface")
    
    
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
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER_MISSINGNETWORKINTERFACEPARAMETER = "MissingParameter.MissingNetworkInterfaceParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteNetworkInterface(request *DeleteNetworkInterfaceRequest) (response *DeleteNetworkInterfaceResponse, err error) {
    return c.DeleteNetworkInterfaceWithContext(context.Background(), request)
}

// DeleteNetworkInterface
// This API is used to delete an ENI.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER_MISSINGNETWORKINTERFACEPARAMETER = "MissingParameter.MissingNetworkInterfaceParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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
    
    request.Init().WithApiInfo("ecm", APIVersion, "DeleteRouteTable")
    
    
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
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_DELDEFAULTROUTE = "UnsupportedOperation.DelDefaultRoute"
//  UNSUPPORTEDOPERATION_DELROUTEWITHSUBNET = "UnsupportedOperation.DelRouteWithSubnet"
func (c *Client) DeleteRouteTable(request *DeleteRouteTableRequest) (response *DeleteRouteTableResponse, err error) {
    return c.DeleteRouteTableWithContext(context.Background(), request)
}

// DeleteRouteTable
// This API is used to delete a route table.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_DELDEFAULTROUTE = "UnsupportedOperation.DelDefaultRoute"
//  UNSUPPORTEDOPERATION_DELROUTEWITHSUBNET = "UnsupportedOperation.DelRouteWithSubnet"
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
    
    request.Init().WithApiInfo("ecm", APIVersion, "DeleteRoutes")
    
    
    return
}

func NewDeleteRoutesResponse() (response *DeleteRoutesResponse) {
    response = &DeleteRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRoutes
// This API is used to batch delete routing policies from a route table.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_SYSTEMROUTE = "UnsupportedOperation.SystemRoute"
func (c *Client) DeleteRoutes(request *DeleteRoutesRequest) (response *DeleteRoutesResponse, err error) {
    return c.DeleteRoutesWithContext(context.Background(), request)
}

// DeleteRoutes
// This API is used to batch delete routing policies from a route table.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  RESOURCENOTFOUND = "ResourceNotFound"
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
    
    request.Init().WithApiInfo("ecm", APIVersion, "DeleteSecurityGroup")
    
    
    return
}

func NewDeleteSecurityGroupResponse() (response *DeleteSecurityGroupResponse) {
    response = &DeleteSecurityGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSecurityGroup
// Only security groups under the current account can be deleted.
//
// A security group cannot be deleted directly if its instance ID is used in the policy of another security group. In this case, you need to modify the policy first before deleting the security group.
//
// Deleted security groups cannot be recovered. Therefore, call this API with caution.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteSecurityGroup(request *DeleteSecurityGroupRequest) (response *DeleteSecurityGroupResponse, err error) {
    return c.DeleteSecurityGroupWithContext(context.Background(), request)
}

// DeleteSecurityGroup
// Only security groups under the current account can be deleted.
//
// A security group cannot be deleted directly if its instance ID is used in the policy of another security group. In this case, you need to modify the policy first before deleting the security group.
//
// Deleted security groups cannot be recovered. Therefore, call this API with caution.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
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
    
    request.Init().WithApiInfo("ecm", APIVersion, "DeleteSecurityGroupPolicies")
    
    
    return
}

func NewDeleteSecurityGroupPoliciesResponse() (response *DeleteSecurityGroupPoliciesResponse) {
    response = &DeleteSecurityGroupPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSecurityGroupPolicies
// `SecurityGroupPolicySet.Version` is used to specify the version of the security group to be manipulated. If the `Version` value passed in differs from the current latest version of the security group, a failure will be returned. If `Version` is not passed in, the policy of the specified `PolicyIndex` will be deleted directly.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_VERSIONMISMATCH = "UnsupportedOperation.VersionMismatch"
func (c *Client) DeleteSecurityGroupPolicies(request *DeleteSecurityGroupPoliciesRequest) (response *DeleteSecurityGroupPoliciesResponse, err error) {
    return c.DeleteSecurityGroupPoliciesWithContext(context.Background(), request)
}

// DeleteSecurityGroupPolicies
// `SecurityGroupPolicySet.Version` is used to specify the version of the security group to be manipulated. If the `Version` value passed in differs from the current latest version of the security group, a failure will be returned. If `Version` is not passed in, the policy of the specified `PolicyIndex` will be deleted directly.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  MISSINGPARAMETER = "MissingParameter"
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

func NewDeleteSnapshotsRequest() (request *DeleteSnapshotsRequest) {
    request = &DeleteSnapshotsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DeleteSnapshots")
    
    
    return
}

func NewDeleteSnapshotsResponse() (response *DeleteSnapshotsResponse) {
    response = &DeleteSnapshotsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSnapshots
// This API is used to delete a snapshot.
//
// 
//
// * Only snapshots in the `NORMAL` state can be deleted. To query the state of a snapshot, you can call the [DescribeSnapshots](https://intl.cloud.tencent.com/document/product/362/15647?from_cn_redirect=1) API and check the `SnapshotState` field in the response.
//
// * Batch operations are supported. If there is any snapshot that cannot be deleted, the operation will not be performed and a specific error code will be returned.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDSNAPSHOT = "InvalidParameterValue.InvalidSnapshot"
//  INVALIDPARAMETERVALUE_INVALIDSNAPSHOTID = "InvalidParameterValue.InvalidSnapshotId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNSUPPORTEDOPERATION_SNAPHASSHARED = "UnsupportedOperation.SnapHasShared"
//  UNSUPPORTEDOPERATION_SNAPSHOTHASBINDEDIMAGE = "UnsupportedOperation.SnapshotHasBindedImage"
func (c *Client) DeleteSnapshots(request *DeleteSnapshotsRequest) (response *DeleteSnapshotsResponse, err error) {
    return c.DeleteSnapshotsWithContext(context.Background(), request)
}

// DeleteSnapshots
// This API is used to delete a snapshot.
//
// 
//
// * Only snapshots in the `NORMAL` state can be deleted. To query the state of a snapshot, you can call the [DescribeSnapshots](https://intl.cloud.tencent.com/document/product/362/15647?from_cn_redirect=1) API and check the `SnapshotState` field in the response.
//
// * Batch operations are supported. If there is any snapshot that cannot be deleted, the operation will not be performed and a specific error code will be returned.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDSNAPSHOT = "InvalidParameterValue.InvalidSnapshot"
//  INVALIDPARAMETERVALUE_INVALIDSNAPSHOTID = "InvalidParameterValue.InvalidSnapshotId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNSUPPORTEDOPERATION_SNAPHASSHARED = "UnsupportedOperation.SnapHasShared"
//  UNSUPPORTEDOPERATION_SNAPSHOTHASBINDEDIMAGE = "UnsupportedOperation.SnapshotHasBindedImage"
func (c *Client) DeleteSnapshotsWithContext(ctx context.Context, request *DeleteSnapshotsRequest) (response *DeleteSnapshotsResponse, err error) {
    if request == nil {
        request = NewDeleteSnapshotsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSnapshots require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSnapshotsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSubnetRequest() (request *DeleteSubnetRequest) {
    request = &DeleteSubnetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DeleteSubnet")
    
    
    return
}

func NewDeleteSubnetResponse() (response *DeleteSubnetResponse) {
    response = &DeleteSubnetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSubnet
// This API is used to delete a subnet. If the subnet is the current default subnet in the AZ, after it is deleted, the default subnet automatically created by the system rather than the last subnet created by you will be used as the new default subnet. If the new default subnet does not meet your needs, you can call the API for setting the default subnet to configure it.
//
// error code that may be returned:
//  FAILEDOPERATION_DATAOPERATIONFAILED = "FailedOperation.DataOperationFailed"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DeleteSubnet(request *DeleteSubnetRequest) (response *DeleteSubnetResponse, err error) {
    return c.DeleteSubnetWithContext(context.Background(), request)
}

// DeleteSubnet
// This API is used to delete a subnet. If the subnet is the current default subnet in the AZ, after it is deleted, the default subnet automatically created by the system rather than the last subnet created by you will be used as the new default subnet. If the new default subnet does not meet your needs, you can call the API for setting the default subnet to configure it.
//
// error code that may be returned:
//  FAILEDOPERATION_DATAOPERATIONFAILED = "FailedOperation.DataOperationFailed"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
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
    
    request.Init().WithApiInfo("ecm", APIVersion, "DeleteVpc")
    
    
    return
}

func NewDeleteVpcResponse() (response *DeleteVpcResponse) {
    response = &DeleteVpcResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteVpc
// This API is used to delete a VPC.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DeleteVpc(request *DeleteVpcRequest) (response *DeleteVpcResponse, err error) {
    return c.DeleteVpcWithContext(context.Background(), request)
}

// DeleteVpc
// This API is used to delete a VPC.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
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

func NewDescribeAddressQuotaRequest() (request *DescribeAddressQuotaRequest) {
    request = &DescribeAddressQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeAddressQuota")
    
    
    return
}

func NewDescribeAddressQuotaResponse() (response *DescribeAddressQuotaResponse) {
    response = &DescribeAddressQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAddressQuota
// This API is used to query the quota information of the EIP under your account in the current region.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeAddressQuota(request *DescribeAddressQuotaRequest) (response *DescribeAddressQuotaResponse, err error) {
    return c.DescribeAddressQuotaWithContext(context.Background(), request)
}

// DescribeAddressQuota
// This API is used to query the quota information of the EIP under your account in the current region.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
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

func NewDescribeAddressesRequest() (request *DescribeAddressesRequest) {
    request = &DescribeAddressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeAddresses")
    
    
    return
}

func NewDescribeAddressesResponse() (response *DescribeAddressesResponse) {
    response = &DescribeAddressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAddresses
// This API is used to query the list of EIPs.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_MALFORMED = "UnsupportedOperation.Malformed"
func (c *Client) DescribeAddresses(request *DescribeAddressesRequest) (response *DescribeAddressesResponse, err error) {
    return c.DescribeAddressesWithContext(context.Background(), request)
}

// DescribeAddresses
// This API is used to query the list of EIPs.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_MALFORMED = "UnsupportedOperation.Malformed"
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

func NewDescribeBaseOverviewRequest() (request *DescribeBaseOverviewRequest) {
    request = &DescribeBaseOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeBaseOverview")
    
    
    return
}

func NewDescribeBaseOverviewResponse() (response *DescribeBaseOverviewResponse) {
    response = &DescribeBaseOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBaseOverview
// This API is used to get the basic data displayed on the overview page.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  MISSINGPARAMETER_MISSINGOVERVIEWPARAMETER = "MissingParameter.MissingOverViewParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeBaseOverview(request *DescribeBaseOverviewRequest) (response *DescribeBaseOverviewResponse, err error) {
    return c.DescribeBaseOverviewWithContext(context.Background(), request)
}

// DescribeBaseOverview
// This API is used to get the basic data displayed on the overview page.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  MISSINGPARAMETER_MISSINGOVERVIEWPARAMETER = "MissingParameter.MissingOverViewParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeBaseOverviewWithContext(ctx context.Context, request *DescribeBaseOverviewRequest) (response *DescribeBaseOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeBaseOverviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaseOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaseOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigRequest() (request *DescribeConfigRequest) {
    request = &DescribeConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeConfig")
    
    
    return
}

func NewDescribeConfigResponse() (response *DescribeConfigResponse) {
    response = &DescribeConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConfig
// This API is used to get the limits of data such as bandwidth and disk.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  MISSINGPARAMETER_MISSINGBASECONFIGPARAMETER = "MissingParameter.MissingBaseConfigParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeConfig(request *DescribeConfigRequest) (response *DescribeConfigResponse, err error) {
    return c.DescribeConfigWithContext(context.Background(), request)
}

// DescribeConfig
// This API is used to get the limits of data such as bandwidth and disk.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  MISSINGPARAMETER_MISSINGBASECONFIGPARAMETER = "MissingParameter.MissingBaseConfigParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeConfigWithContext(ctx context.Context, request *DescribeConfigRequest) (response *DescribeConfigResponse, err error) {
    if request == nil {
        request = NewDescribeConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomImageTaskRequest() (request *DescribeCustomImageTaskRequest) {
    request = &DescribeCustomImageTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeCustomImageTask")
    
    
    return
}

func NewDescribeCustomImageTaskResponse() (response *DescribeCustomImageTaskResponse) {
    response = &DescribeCustomImageTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCustomImageTask
// This API is used to query an image import task.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCustomImageTask(request *DescribeCustomImageTaskRequest) (response *DescribeCustomImageTaskResponse, err error) {
    return c.DescribeCustomImageTaskWithContext(context.Background(), request)
}

// DescribeCustomImageTask
// This API is used to query an image import task.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCustomImageTaskWithContext(ctx context.Context, request *DescribeCustomImageTaskRequest) (response *DescribeCustomImageTaskResponse, err error) {
    if request == nil {
        request = NewDescribeCustomImageTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCustomImageTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCustomImageTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDefaultSubnetRequest() (request *DescribeDefaultSubnetRequest) {
    request = &DescribeDefaultSubnetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeDefaultSubnet")
    
    
    return
}

func NewDescribeDefaultSubnetResponse() (response *DescribeDefaultSubnetResponse) {
    response = &DescribeDefaultSubnetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDefaultSubnet
// This API is used to query the default subnet in an AZ.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDZONE = "InvalidParameterValue.InvalidZone"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeDefaultSubnet(request *DescribeDefaultSubnetRequest) (response *DescribeDefaultSubnetResponse, err error) {
    return c.DescribeDefaultSubnetWithContext(context.Background(), request)
}

// DescribeDefaultSubnet
// This API is used to query the default subnet in an AZ.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDZONE = "InvalidParameterValue.InvalidZone"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeDefaultSubnetWithContext(ctx context.Context, request *DescribeDefaultSubnetRequest) (response *DescribeDefaultSubnetResponse, err error) {
    if request == nil {
        request = NewDescribeDefaultSubnetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDefaultSubnet require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDefaultSubnetResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHaVipsRequest() (request *DescribeHaVipsRequest) {
    request = &DescribeHaVipsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeHaVips")
    
    
    return
}

func NewDescribeHaVipsResponse() (response *DescribeHaVipsResponse) {
    response = &DescribeHaVipsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHaVips
// This API is used to query the list of HAVIPs.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE_INVAILDPAGEPARAM = "InvalidParameterValue.InvaildPageParam"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeHaVips(request *DescribeHaVipsRequest) (response *DescribeHaVipsResponse, err error) {
    return c.DescribeHaVipsWithContext(context.Background(), request)
}

// DescribeHaVips
// This API is used to query the list of HAVIPs.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE_INVAILDPAGEPARAM = "InvalidParameterValue.InvaildPageParam"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
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

func NewDescribeImageRequest() (request *DescribeImageRequest) {
    request = &DescribeImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeImage")
    
    
    return
}

func NewDescribeImageResponse() (response *DescribeImageResponse) {
    response = &DescribeImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeImage
// This API is used to display the list of images.
//
// error code that may be returned:
//  FAILEDOPERATION_IMAGEINUSE = "FailedOperation.ImageInUse"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVAILDPAGEPARAM = "InvalidParameterValue.InvaildPageParam"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageID"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  MISSINGPARAMETER_MISSINGIMAGEPARAMETER = "MissingParameter.MissingImageParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeImage(request *DescribeImageRequest) (response *DescribeImageResponse, err error) {
    return c.DescribeImageWithContext(context.Background(), request)
}

// DescribeImage
// This API is used to display the list of images.
//
// error code that may be returned:
//  FAILEDOPERATION_IMAGEINUSE = "FailedOperation.ImageInUse"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVAILDPAGEPARAM = "InvalidParameterValue.InvaildPageParam"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageID"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  MISSINGPARAMETER_MISSINGIMAGEPARAMETER = "MissingParameter.MissingImageParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeImageWithContext(ctx context.Context, request *DescribeImageRequest) (response *DescribeImageResponse, err error) {
    if request == nil {
        request = NewDescribeImageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImportImageOsRequest() (request *DescribeImportImageOsRequest) {
    request = &DescribeImportImageOsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeImportImageOs")
    
    
    return
}

func NewDescribeImportImageOsResponse() (response *DescribeImportImageOsResponse) {
    response = &DescribeImportImageOsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeImportImageOs
// This API is used to query the list of operating systems supported by an image imported from an external resource.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
func (c *Client) DescribeImportImageOs(request *DescribeImportImageOsRequest) (response *DescribeImportImageOsResponse, err error) {
    return c.DescribeImportImageOsWithContext(context.Background(), request)
}

// DescribeImportImageOs
// This API is used to query the list of operating systems supported by an image imported from an external resource.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
func (c *Client) DescribeImportImageOsWithContext(ctx context.Context, request *DescribeImportImageOsRequest) (response *DescribeImportImageOsResponse, err error) {
    if request == nil {
        request = NewDescribeImportImageOsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImportImageOs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImportImageOsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceTypeConfigRequest() (request *DescribeInstanceTypeConfigRequest) {
    request = &DescribeInstanceTypeConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeInstanceTypeConfig")
    
    
    return
}

func NewDescribeInstanceTypeConfigResponse() (response *DescribeInstanceTypeConfigResponse) {
    response = &DescribeInstanceTypeConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceTypeConfig
// This API is used to get the list of model configurations.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  MISSINGPARAMETER_MISSINGINSTANCETYPECONFIGPARAMETER = "MissingParameter.MissingInstanceTypeConfigParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeInstanceTypeConfig(request *DescribeInstanceTypeConfigRequest) (response *DescribeInstanceTypeConfigResponse, err error) {
    return c.DescribeInstanceTypeConfigWithContext(context.Background(), request)
}

// DescribeInstanceTypeConfig
// This API is used to get the list of model configurations.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  MISSINGPARAMETER_MISSINGINSTANCETYPECONFIGPARAMETER = "MissingParameter.MissingInstanceTypeConfigParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeInstanceTypeConfigWithContext(ctx context.Context, request *DescribeInstanceTypeConfigRequest) (response *DescribeInstanceTypeConfigResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceTypeConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceTypeConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceTypeConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceVncUrlRequest() (request *DescribeInstanceVncUrlRequest) {
    request = &DescribeInstanceVncUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeInstanceVncUrl")
    
    
    return
}

func NewDescribeInstanceVncUrlResponse() (response *DescribeInstanceVncUrlResponse) {
    response = &DescribeInstanceVncUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceVncUrl
// This API is used to query the VNC URL of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  MISSINGPARAMETER_MISSINGINSTANCESPARAMETER = "MissingParameter.MissingInstancesParameter"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
func (c *Client) DescribeInstanceVncUrl(request *DescribeInstanceVncUrlRequest) (response *DescribeInstanceVncUrlResponse, err error) {
    return c.DescribeInstanceVncUrlWithContext(context.Background(), request)
}

// DescribeInstanceVncUrl
// This API is used to query the VNC URL of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  MISSINGPARAMETER_MISSINGINSTANCESPARAMETER = "MissingParameter.MissingInstancesParameter"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
func (c *Client) DescribeInstanceVncUrlWithContext(ctx context.Context, request *DescribeInstanceVncUrlRequest) (response *DescribeInstanceVncUrlResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceVncUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceVncUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceVncUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeInstances")
    
    
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstances
// This API is used to get the information of an instance.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVAILDDESCRIBEINSTANCE = "InvalidParameterValue.InvaildDescribeInstance"
//  INVALIDPARAMETERVALUE_INVAILDPAGEPARAM = "InvalidParameterValue.InvaildPageParam"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_INVALIDORDERBYFIELD = "InvalidParameterValue.InvalidOrderByField"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    return c.DescribeInstancesWithContext(context.Background(), request)
}

// DescribeInstances
// This API is used to get the information of an instance.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVAILDDESCRIBEINSTANCE = "InvalidParameterValue.InvaildDescribeInstance"
//  INVALIDPARAMETERVALUE_INVAILDPAGEPARAM = "InvalidParameterValue.InvaildPageParam"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_INVALIDORDERBYFIELD = "InvalidParameterValue.InvalidOrderByField"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeInstancesWithContext(ctx context.Context, request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesDeniedActionsRequest() (request *DescribeInstancesDeniedActionsRequest) {
    request = &DescribeInstancesDeniedActionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeInstancesDeniedActions")
    
    
    return
}

func NewDescribeInstancesDeniedActionsResponse() (response *DescribeInstancesDeniedActionsResponse) {
    response = &DescribeInstancesDeniedActionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstancesDeniedActions
// This API is used to get the information of prohibited operations by instance ID.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  MISSINGPARAMETER_MISSINGINSTANCESPARAMETER = "MissingParameter.MissingInstancesParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeInstancesDeniedActions(request *DescribeInstancesDeniedActionsRequest) (response *DescribeInstancesDeniedActionsResponse, err error) {
    return c.DescribeInstancesDeniedActionsWithContext(context.Background(), request)
}

// DescribeInstancesDeniedActions
// This API is used to get the information of prohibited operations by instance ID.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  MISSINGPARAMETER_MISSINGINSTANCESPARAMETER = "MissingParameter.MissingInstancesParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeInstancesDeniedActionsWithContext(ctx context.Context, request *DescribeInstancesDeniedActionsRequest) (response *DescribeInstancesDeniedActionsResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesDeniedActionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstancesDeniedActions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesDeniedActionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListenersRequest() (request *DescribeListenersRequest) {
    request = &DescribeListenersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeListeners")
    
    
    return
}

func NewDescribeListenersResponse() (response *DescribeListenersResponse) {
    response = &DescribeListenersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeListeners
// This API is used to query the list of CLB listeners.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeListeners(request *DescribeListenersRequest) (response *DescribeListenersResponse, err error) {
    return c.DescribeListenersWithContext(context.Background(), request)
}

// DescribeListeners
// This API is used to query the list of CLB listeners.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
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

func NewDescribeLoadBalanceTaskStatusRequest() (request *DescribeLoadBalanceTaskStatusRequest) {
    request = &DescribeLoadBalanceTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeLoadBalanceTaskStatus")
    
    
    return
}

func NewDescribeLoadBalanceTaskStatusResponse() (response *DescribeLoadBalanceTaskStatusResponse) {
    response = &DescribeLoadBalanceTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLoadBalanceTaskStatus
// This API is used to query the task status of a CLB instance.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeLoadBalanceTaskStatus(request *DescribeLoadBalanceTaskStatusRequest) (response *DescribeLoadBalanceTaskStatusResponse, err error) {
    return c.DescribeLoadBalanceTaskStatusWithContext(context.Background(), request)
}

// DescribeLoadBalanceTaskStatus
// This API is used to query the task status of a CLB instance.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeLoadBalanceTaskStatusWithContext(ctx context.Context, request *DescribeLoadBalanceTaskStatusRequest) (response *DescribeLoadBalanceTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeLoadBalanceTaskStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLoadBalanceTaskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLoadBalanceTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLoadBalancersRequest() (request *DescribeLoadBalancersRequest) {
    request = &DescribeLoadBalancersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeLoadBalancers")
    
    
    return
}

func NewDescribeLoadBalancersResponse() (response *DescribeLoadBalancersResponse) {
    response = &DescribeLoadBalancersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLoadBalancers
// This API is used to query the list of CLB instances.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeLoadBalancers(request *DescribeLoadBalancersRequest) (response *DescribeLoadBalancersResponse, err error) {
    return c.DescribeLoadBalancersWithContext(context.Background(), request)
}

// DescribeLoadBalancers
// This API is used to query the list of CLB instances.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
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

func NewDescribeModuleRequest() (request *DescribeModuleRequest) {
    request = &DescribeModuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeModule")
    
    
    return
}

func NewDescribeModuleResponse() (response *DescribeModuleResponse) {
    response = &DescribeModuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeModule
// This API is used to get the list of modules.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVAILDPAGEPARAM = "InvalidParameterValue.InvaildPageParam"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeModule(request *DescribeModuleRequest) (response *DescribeModuleResponse, err error) {
    return c.DescribeModuleWithContext(context.Background(), request)
}

// DescribeModule
// This API is used to get the list of modules.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVAILDPAGEPARAM = "InvalidParameterValue.InvaildPageParam"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeModuleWithContext(ctx context.Context, request *DescribeModuleRequest) (response *DescribeModuleResponse, err error) {
    if request == nil {
        request = NewDescribeModuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeModuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModuleDetailRequest() (request *DescribeModuleDetailRequest) {
    request = &DescribeModuleDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeModuleDetail")
    
    
    return
}

func NewDescribeModuleDetailResponse() (response *DescribeModuleDetailResponse) {
    response = &DescribeModuleDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeModuleDetail
// This API is used to display the details of a module.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeModuleDetail(request *DescribeModuleDetailRequest) (response *DescribeModuleDetailResponse, err error) {
    return c.DescribeModuleDetailWithContext(context.Background(), request)
}

// DescribeModuleDetail
// This API is used to display the details of a module.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeModuleDetailWithContext(ctx context.Context, request *DescribeModuleDetailRequest) (response *DescribeModuleDetailResponse, err error) {
    if request == nil {
        request = NewDescribeModuleDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModuleDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeModuleDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMonthPeakNetworkRequest() (request *DescribeMonthPeakNetworkRequest) {
    request = &DescribeMonthPeakNetworkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeMonthPeakNetwork")
    
    
    return
}

func NewDescribeMonthPeakNetworkResponse() (response *DescribeMonthPeakNetworkResponse) {
    response = &DescribeMonthPeakNetworkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMonthPeakNetwork
// This API is used to get the monthly peak and billable inbound/outbound bandwidth values of your node.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDISPINNODE = "InvalidParameterValue.InvalidISPInNode"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeMonthPeakNetwork(request *DescribeMonthPeakNetworkRequest) (response *DescribeMonthPeakNetworkResponse, err error) {
    return c.DescribeMonthPeakNetworkWithContext(context.Background(), request)
}

// DescribeMonthPeakNetwork
// This API is used to get the monthly peak and billable inbound/outbound bandwidth values of your node.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDISPINNODE = "InvalidParameterValue.InvalidISPInNode"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeMonthPeakNetworkWithContext(ctx context.Context, request *DescribeMonthPeakNetworkRequest) (response *DescribeMonthPeakNetworkResponse, err error) {
    if request == nil {
        request = NewDescribeMonthPeakNetworkRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMonthPeakNetwork require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMonthPeakNetworkResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNetworkInterfacesRequest() (request *DescribeNetworkInterfacesRequest) {
    request = &DescribeNetworkInterfacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeNetworkInterfaces")
    
    
    return
}

func NewDescribeNetworkInterfacesResponse() (response *DescribeNetworkInterfacesResponse) {
    response = &DescribeNetworkInterfacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNetworkInterfaces
// This API is used to query the list of ENIs.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNetworkInterfaces(request *DescribeNetworkInterfacesRequest) (response *DescribeNetworkInterfacesResponse, err error) {
    return c.DescribeNetworkInterfacesWithContext(context.Background(), request)
}

// DescribeNetworkInterfaces
// This API is used to query the list of ENIs.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
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

func NewDescribeNodeRequest() (request *DescribeNodeRequest) {
    request = &DescribeNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeNode")
    
    
    return
}

func NewDescribeNodeResponse() (response *DescribeNodeResponse) {
    response = &DescribeNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNode
// This API is used to get the list of nodes.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  MISSINGPARAMETER_MISSINGNODEPARAMETER = "MissingParameter.MissingNodeParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeNode(request *DescribeNodeRequest) (response *DescribeNodeResponse, err error) {
    return c.DescribeNodeWithContext(context.Background(), request)
}

// DescribeNode
// This API is used to get the list of nodes.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  MISSINGPARAMETER_MISSINGNODEPARAMETER = "MissingParameter.MissingNodeParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeNodeWithContext(ctx context.Context, request *DescribeNodeRequest) (response *DescribeNodeResponse, err error) {
    if request == nil {
        request = NewDescribeNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNodeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePackingQuotaGroupRequest() (request *DescribePackingQuotaGroupRequest) {
    request = &DescribePackingQuotaGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DescribePackingQuotaGroup")
    
    
    return
}

func NewDescribePackingQuotaGroupResponse() (response *DescribePackingQuotaGroupResponse) {
    response = &DescribePackingQuotaGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePackingQuotaGroup
// This API is used to get the packing quota of a model in a region (when a virtual model is used, a set of correlated packing quotas will be returned).
//
// error code that may be returned:
//  FAILEDOPERATION_DATAOPERATIONFAILED = "FailedOperation.DataOperationFailed"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePackingQuotaGroup(request *DescribePackingQuotaGroupRequest) (response *DescribePackingQuotaGroupResponse, err error) {
    return c.DescribePackingQuotaGroupWithContext(context.Background(), request)
}

// DescribePackingQuotaGroup
// This API is used to get the packing quota of a model in a region (when a virtual model is used, a set of correlated packing quotas will be returned).
//
// error code that may be returned:
//  FAILEDOPERATION_DATAOPERATIONFAILED = "FailedOperation.DataOperationFailed"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePackingQuotaGroupWithContext(ctx context.Context, request *DescribePackingQuotaGroupRequest) (response *DescribePackingQuotaGroupResponse, err error) {
    if request == nil {
        request = NewDescribePackingQuotaGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePackingQuotaGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePackingQuotaGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePeakBaseOverviewRequest() (request *DescribePeakBaseOverviewRequest) {
    request = &DescribePeakBaseOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DescribePeakBaseOverview")
    
    
    return
}

func NewDescribePeakBaseOverviewResponse() (response *DescribePeakBaseOverviewResponse) {
    response = &DescribePeakBaseOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePeakBaseOverview
// This API is used to get the peak data of basic information such as CPU, memory, and disk.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDTIME = "InvalidParameterValue.InvalidTime"
//  MISSINGPARAMETER_MISSINGOVERVIEWPARAMETER = "MissingParameter.MissingOverViewParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribePeakBaseOverview(request *DescribePeakBaseOverviewRequest) (response *DescribePeakBaseOverviewResponse, err error) {
    return c.DescribePeakBaseOverviewWithContext(context.Background(), request)
}

// DescribePeakBaseOverview
// This API is used to get the peak data of basic information such as CPU, memory, and disk.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDTIME = "InvalidParameterValue.InvalidTime"
//  MISSINGPARAMETER_MISSINGOVERVIEWPARAMETER = "MissingParameter.MissingOverViewParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribePeakBaseOverviewWithContext(ctx context.Context, request *DescribePeakBaseOverviewRequest) (response *DescribePeakBaseOverviewResponse, err error) {
    if request == nil {
        request = NewDescribePeakBaseOverviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePeakBaseOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePeakBaseOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePeakNetworkOverviewRequest() (request *DescribePeakNetworkOverviewRequest) {
    request = &DescribePeakNetworkOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DescribePeakNetworkOverview")
    
    
    return
}

func NewDescribePeakNetworkOverviewResponse() (response *DescribePeakNetworkOverviewResponse) {
    response = &DescribePeakNetworkOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePeakNetworkOverview
// This API is used to get the peak network data.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDTIME = "InvalidParameterValue.InvalidTime"
//  MISSINGPARAMETER_MISSINGOVERVIEWPARAMETER = "MissingParameter.MissingOverViewParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribePeakNetworkOverview(request *DescribePeakNetworkOverviewRequest) (response *DescribePeakNetworkOverviewResponse, err error) {
    return c.DescribePeakNetworkOverviewWithContext(context.Background(), request)
}

// DescribePeakNetworkOverview
// This API is used to get the peak network data.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDTIME = "InvalidParameterValue.InvalidTime"
//  MISSINGPARAMETER_MISSINGOVERVIEWPARAMETER = "MissingParameter.MissingOverViewParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribePeakNetworkOverviewWithContext(ctx context.Context, request *DescribePeakNetworkOverviewRequest) (response *DescribePeakNetworkOverviewResponse, err error) {
    if request == nil {
        request = NewDescribePeakNetworkOverviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePeakNetworkOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePeakNetworkOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePriceRunInstanceRequest() (request *DescribePriceRunInstanceRequest) {
    request = &DescribePriceRunInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DescribePriceRunInstance")
    
    
    return
}

func NewDescribePriceRunInstanceResponse() (response *DescribePriceRunInstanceResponse) {
    response = &DescribePriceRunInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePriceRunInstance
// This API is used to query the price of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCETYPECONFIGID = "InvalidParameterValue.InvalidInstanceTypeConfigID"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDSYSTEMDISKSIZE = "InvalidParameterValue.InvalidSystemDiskSize"
//  INVALIDPARAMETERVALUE_INVALIDSYSTEMDISKTYPE = "InvalidParameterValue.InvalidSystemDiskType"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribePriceRunInstance(request *DescribePriceRunInstanceRequest) (response *DescribePriceRunInstanceResponse, err error) {
    return c.DescribePriceRunInstanceWithContext(context.Background(), request)
}

// DescribePriceRunInstance
// This API is used to query the price of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCETYPECONFIGID = "InvalidParameterValue.InvalidInstanceTypeConfigID"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDSYSTEMDISKSIZE = "InvalidParameterValue.InvalidSystemDiskSize"
//  INVALIDPARAMETERVALUE_INVALIDSYSTEMDISKTYPE = "InvalidParameterValue.InvalidSystemDiskType"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribePriceRunInstanceWithContext(ctx context.Context, request *DescribePriceRunInstanceRequest) (response *DescribePriceRunInstanceResponse, err error) {
    if request == nil {
        request = NewDescribePriceRunInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePriceRunInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePriceRunInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRouteConflictsRequest() (request *DescribeRouteConflictsRequest) {
    request = &DescribeRouteConflictsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeRouteConflicts")
    
    
    return
}

func NewDescribeRouteConflictsResponse() (response *DescribeRouteConflictsResponse) {
    response = &DescribeRouteConflictsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRouteConflicts
// This API is used to query the list of conflicts between a custom routing policy and a CCN routing policy.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeRouteConflicts(request *DescribeRouteConflictsRequest) (response *DescribeRouteConflictsResponse, err error) {
    return c.DescribeRouteConflictsWithContext(context.Background(), request)
}

// DescribeRouteConflicts
// This API is used to query the list of conflicts between a custom routing policy and a CCN routing policy.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeRouteConflictsWithContext(ctx context.Context, request *DescribeRouteConflictsRequest) (response *DescribeRouteConflictsResponse, err error) {
    if request == nil {
        request = NewDescribeRouteConflictsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRouteConflicts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRouteConflictsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRouteTablesRequest() (request *DescribeRouteTablesRequest) {
    request = &DescribeRouteTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeRouteTables")
    
    
    return
}

func NewDescribeRouteTablesResponse() (response *DescribeRouteTablesResponse) {
    response = &DescribeRouteTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRouteTables
// This API is used to query the list of the objects in a route table.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeRouteTables(request *DescribeRouteTablesRequest) (response *DescribeRouteTablesResponse, err error) {
    return c.DescribeRouteTablesWithContext(context.Background(), request)
}

// DescribeRouteTables
// This API is used to query the list of the objects in a route table.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
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
    
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeSecurityGroupAssociationStatistics")
    
    
    return
}

func NewDescribeSecurityGroupAssociationStatisticsResponse() (response *DescribeSecurityGroupAssociationStatisticsResponse) {
    response = &DescribeSecurityGroupAssociationStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecurityGroupAssociationStatistics
// This API is used to query statistics on the instances associated with a security group.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSecurityGroupAssociationStatistics(request *DescribeSecurityGroupAssociationStatisticsRequest) (response *DescribeSecurityGroupAssociationStatisticsResponse, err error) {
    return c.DescribeSecurityGroupAssociationStatisticsWithContext(context.Background(), request)
}

// DescribeSecurityGroupAssociationStatistics
// This API is used to query statistics on the instances associated with a security group.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
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

func NewDescribeSecurityGroupLimitsRequest() (request *DescribeSecurityGroupLimitsRequest) {
    request = &DescribeSecurityGroupLimitsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeSecurityGroupLimits")
    
    
    return
}

func NewDescribeSecurityGroupLimitsResponse() (response *DescribeSecurityGroupLimitsResponse) {
    response = &DescribeSecurityGroupLimitsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecurityGroupLimits
// This API is used to query the security group quota.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeSecurityGroupLimits(request *DescribeSecurityGroupLimitsRequest) (response *DescribeSecurityGroupLimitsResponse, err error) {
    return c.DescribeSecurityGroupLimitsWithContext(context.Background(), request)
}

// DescribeSecurityGroupLimits
// This API is used to query the security group quota.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeSecurityGroupLimitsWithContext(ctx context.Context, request *DescribeSecurityGroupLimitsRequest) (response *DescribeSecurityGroupLimitsResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityGroupLimitsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityGroupLimits require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityGroupLimitsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityGroupPoliciesRequest() (request *DescribeSecurityGroupPoliciesRequest) {
    request = &DescribeSecurityGroupPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeSecurityGroupPolicies")
    
    
    return
}

func NewDescribeSecurityGroupPoliciesResponse() (response *DescribeSecurityGroupPoliciesResponse) {
    response = &DescribeSecurityGroupPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecurityGroupPolicies
// This API is used to query a security group rule.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSecurityGroupPolicies(request *DescribeSecurityGroupPoliciesRequest) (response *DescribeSecurityGroupPoliciesResponse, err error) {
    return c.DescribeSecurityGroupPoliciesWithContext(context.Background(), request)
}

// DescribeSecurityGroupPolicies
// This API is used to query a security group rule.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
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

func NewDescribeSecurityGroupsRequest() (request *DescribeSecurityGroupsRequest) {
    request = &DescribeSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeSecurityGroups")
    
    
    return
}

func NewDescribeSecurityGroupsResponse() (response *DescribeSecurityGroupsResponse) {
    response = &DescribeSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecurityGroups
// This API is used to view a security group.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeSecurityGroups(request *DescribeSecurityGroupsRequest) (response *DescribeSecurityGroupsResponse, err error) {
    return c.DescribeSecurityGroupsWithContext(context.Background(), request)
}

// DescribeSecurityGroups
// This API is used to view a security group.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
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

func NewDescribeSnapshotsRequest() (request *DescribeSnapshotsRequest) {
    request = &DescribeSnapshotsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeSnapshots")
    
    
    return
}

func NewDescribeSnapshotsResponse() (response *DescribeSnapshotsResponse) {
    response = &DescribeSnapshotsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSnapshots
// This API is used to query the list of snapshots.
//
// 
//
// * You can filter results by snapshot ID and the ID and type of the cloud disk for which the snapshot is created. The relationship between different filters is `AND`. For more information on filters, see `Filter`.
//
// * If no parameter is defined, the status of a certain number of snapshots under the current account will be returned. The number is specified by `Limit` and is 20 by default.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeSnapshots(request *DescribeSnapshotsRequest) (response *DescribeSnapshotsResponse, err error) {
    return c.DescribeSnapshotsWithContext(context.Background(), request)
}

// DescribeSnapshots
// This API is used to query the list of snapshots.
//
// 
//
// * You can filter results by snapshot ID and the ID and type of the cloud disk for which the snapshot is created. The relationship between different filters is `AND`. For more information on filters, see `Filter`.
//
// * If no parameter is defined, the status of a certain number of snapshots under the current account will be returned. The number is specified by `Limit` and is 20 by default.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeSnapshotsWithContext(ctx context.Context, request *DescribeSnapshotsRequest) (response *DescribeSnapshotsResponse, err error) {
    if request == nil {
        request = NewDescribeSnapshotsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSnapshots require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSnapshotsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubnetsRequest() (request *DescribeSubnetsRequest) {
    request = &DescribeSubnetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeSubnets")
    
    
    return
}

func NewDescribeSubnetsResponse() (response *DescribeSubnetsResponse) {
    response = &DescribeSubnetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSubnets
// This API is used to query the list of subnets.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FILTERLIMITEXCEEDED = "InvalidParameterValue.FilterLimitExceeded"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeSubnets(request *DescribeSubnetsRequest) (response *DescribeSubnetsResponse, err error) {
    return c.DescribeSubnetsWithContext(context.Background(), request)
}

// DescribeSubnets
// This API is used to query the list of subnets.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FILTERLIMITEXCEEDED = "InvalidParameterValue.FilterLimitExceeded"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
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

func NewDescribeTargetHealthRequest() (request *DescribeTargetHealthRequest) {
    request = &DescribeTargetHealthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeTargetHealth")
    
    
    return
}

func NewDescribeTargetHealthResponse() (response *DescribeTargetHealthResponse) {
    response = &DescribeTargetHealthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTargetHealth
// This API is used to get the health check status of a CLB real server.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeTargetHealth(request *DescribeTargetHealthRequest) (response *DescribeTargetHealthResponse, err error) {
    return c.DescribeTargetHealthWithContext(context.Background(), request)
}

// DescribeTargetHealth
// This API is used to get the health check status of a CLB real server.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
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
    
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeTargets")
    
    
    return
}

func NewDescribeTargetsResponse() (response *DescribeTargetsResponse) {
    response = &DescribeTargetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTargets
// This API is used to query the list of the real servers bound to a CLB instance.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeTargets(request *DescribeTargetsRequest) (response *DescribeTargetsResponse, err error) {
    return c.DescribeTargetsWithContext(context.Background(), request)
}

// DescribeTargets
// This API is used to query the list of the real servers bound to a CLB instance.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
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

func NewDescribeTaskResultRequest() (request *DescribeTaskResultRequest) {
    request = &DescribeTaskResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeTaskResult")
    
    
    return
}

func NewDescribeTaskResultResponse() (response *DescribeTaskResultResponse) {
    response = &DescribeTaskResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTaskResult
// This API is used to query the execution result of an EIP async task.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeTaskResult(request *DescribeTaskResultRequest) (response *DescribeTaskResultResponse, err error) {
    return c.DescribeTaskResultWithContext(context.Background(), request)
}

// DescribeTaskResult
// This API is used to query the execution result of an EIP async task.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
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

func NewDescribeTaskStatusRequest() (request *DescribeTaskStatusRequest) {
    request = &DescribeTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeTaskStatus")
    
    
    return
}

func NewDescribeTaskStatusResponse() (response *DescribeTaskStatusResponse) {
    response = &DescribeTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTaskStatus
// This API is used to get the status of an async task.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_TASKNOTFOUND = "InvalidParameterValue.TaskNotFound"
func (c *Client) DescribeTaskStatus(request *DescribeTaskStatusRequest) (response *DescribeTaskStatusResponse, err error) {
    return c.DescribeTaskStatusWithContext(context.Background(), request)
}

// DescribeTaskStatus
// This API is used to get the status of an async task.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_TASKNOTFOUND = "InvalidParameterValue.TaskNotFound"
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

func NewDescribeVpcsRequest() (request *DescribeVpcsRequest) {
    request = &DescribeVpcsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeVpcs")
    
    
    return
}

func NewDescribeVpcsResponse() (response *DescribeVpcsResponse) {
    response = &DescribeVpcsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVpcs
// This API is used to query the list of VPCs.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FILTERLIMITEXCEEDED = "InvalidParameterValue.FilterLimitExceeded"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDZONEINSTANCECOUNT = "InvalidParameterValue.InvalidZoneInstanceCount"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) DescribeVpcs(request *DescribeVpcsRequest) (response *DescribeVpcsResponse, err error) {
    return c.DescribeVpcsWithContext(context.Background(), request)
}

// DescribeVpcs
// This API is used to query the list of VPCs.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FILTERLIMITEXCEEDED = "InvalidParameterValue.FilterLimitExceeded"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDZONEINSTANCECOUNT = "InvalidParameterValue.InvalidZoneInstanceCount"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
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

func NewDetachNetworkInterfaceRequest() (request *DetachNetworkInterfaceRequest) {
    request = &DetachNetworkInterfaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DetachNetworkInterface")
    
    
    return
}

func NewDetachNetworkInterfaceResponse() (response *DetachNetworkInterfaceResponse) {
    response = &DetachNetworkInterfaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DetachNetworkInterface
// This API is used to unbind an ENI from a CVM instance.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) DetachNetworkInterface(request *DetachNetworkInterfaceRequest) (response *DetachNetworkInterfaceResponse, err error) {
    return c.DetachNetworkInterfaceWithContext(context.Background(), request)
}

// DetachNetworkInterface
// This API is used to unbind an ENI from a CVM instance.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
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

func NewDisableRoutesRequest() (request *DisableRoutesRequest) {
    request = &DisableRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DisableRoutes")
    
    
    return
}

func NewDisableRoutesResponse() (response *DisableRoutesResponse) {
    response = &DisableRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisableRoutes
// This API is used to disable a subnet route.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_SYSTEMROUTE = "UnsupportedOperation.SystemRoute"
func (c *Client) DisableRoutes(request *DisableRoutesRequest) (response *DisableRoutesResponse, err error) {
    return c.DisableRoutesWithContext(context.Background(), request)
}

// DisableRoutes
// This API is used to disable a subnet route.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_SYSTEMROUTE = "UnsupportedOperation.SystemRoute"
func (c *Client) DisableRoutesWithContext(ctx context.Context, request *DisableRoutesRequest) (response *DisableRoutesResponse, err error) {
    if request == nil {
        request = NewDisableRoutesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableRoutes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateAddressRequest() (request *DisassociateAddressRequest) {
    request = &DisassociateAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DisassociateAddress")
    
    
    return
}

func NewDisassociateAddressResponse() (response *DisassociateAddressResponse) {
    response = &DisassociateAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisassociateAddress
// This API is used to unbind an EIP.
//
// Only EIPs in `BIND` or `BIND_ENI` status can be unbound.
//
// Blocked EIPs cannot be unbound.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVAILDADDRESSID = "InvalidParameterValue.InvaildAddressId"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_ADDRESSIDNOTFOUND = "UnsupportedOperation.AddressIdNotFound"
//  UNSUPPORTEDOPERATION_INSTANCEIDNOTFOUND = "UnsupportedOperation.InstanceIdNotFound"
//  UNSUPPORTEDOPERATION_INSTANCEIDNOTSUPPORTED = "UnsupportedOperation.InstanceIdNotSupported"
//  UNSUPPORTEDOPERATION_QUOTALIMITEXCEEDED = "UnsupportedOperation.QuotaLimitExceeded"
//  UNSUPPORTEDOPERATION_STATUSNOTPERMIT = "UnsupportedOperation.StatusNotPermit"
func (c *Client) DisassociateAddress(request *DisassociateAddressRequest) (response *DisassociateAddressResponse, err error) {
    return c.DisassociateAddressWithContext(context.Background(), request)
}

// DisassociateAddress
// This API is used to unbind an EIP.
//
// Only EIPs in `BIND` or `BIND_ENI` status can be unbound.
//
// Blocked EIPs cannot be unbound.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVAILDADDRESSID = "InvalidParameterValue.InvaildAddressId"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_ADDRESSIDNOTFOUND = "UnsupportedOperation.AddressIdNotFound"
//  UNSUPPORTEDOPERATION_INSTANCEIDNOTFOUND = "UnsupportedOperation.InstanceIdNotFound"
//  UNSUPPORTEDOPERATION_INSTANCEIDNOTSUPPORTED = "UnsupportedOperation.InstanceIdNotSupported"
//  UNSUPPORTEDOPERATION_QUOTALIMITEXCEEDED = "UnsupportedOperation.QuotaLimitExceeded"
//  UNSUPPORTEDOPERATION_STATUSNOTPERMIT = "UnsupportedOperation.StatusNotPermit"
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

func NewDisassociateInstancesKeyPairsRequest() (request *DisassociateInstancesKeyPairsRequest) {
    request = &DisassociateInstancesKeyPairsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DisassociateInstancesKeyPairs")
    
    
    return
}

func NewDisassociateInstancesKeyPairsResponse() (response *DisassociateInstancesKeyPairsResponse) {
    response = &DisassociateInstancesKeyPairsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisassociateInstancesKeyPairs
// This API is used to unbind a key pair from an instance.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCEIDNOTSUPPORTED = "InvalidParameterValue.InstanceIdNotSupported"
//  INVALIDPARAMETERVALUE_INVALIDKEYPAIRID = "InvalidParameterValue.InvalidKeyPairId"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DisassociateInstancesKeyPairs(request *DisassociateInstancesKeyPairsRequest) (response *DisassociateInstancesKeyPairsResponse, err error) {
    return c.DisassociateInstancesKeyPairsWithContext(context.Background(), request)
}

// DisassociateInstancesKeyPairs
// This API is used to unbind a key pair from an instance.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCEIDNOTSUPPORTED = "InvalidParameterValue.InstanceIdNotSupported"
//  INVALIDPARAMETERVALUE_INVALIDKEYPAIRID = "InvalidParameterValue.InvalidKeyPairId"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DisassociateInstancesKeyPairsWithContext(ctx context.Context, request *DisassociateInstancesKeyPairsRequest) (response *DisassociateInstancesKeyPairsResponse, err error) {
    if request == nil {
        request = NewDisassociateInstancesKeyPairsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisassociateInstancesKeyPairs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisassociateInstancesKeyPairsResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateSecurityGroupsRequest() (request *DisassociateSecurityGroupsRequest) {
    request = &DisassociateSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "DisassociateSecurityGroups")
    
    
    return
}

func NewDisassociateSecurityGroupsResponse() (response *DisassociateSecurityGroupsResponse) {
    response = &DisassociateSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisassociateSecurityGroups
// This API is used to unbind a security group.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DisassociateSecurityGroups(request *DisassociateSecurityGroupsRequest) (response *DisassociateSecurityGroupsResponse, err error) {
    return c.DisassociateSecurityGroupsWithContext(context.Background(), request)
}

// DisassociateSecurityGroups
// This API is used to unbind a security group.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DisassociateSecurityGroupsWithContext(ctx context.Context, request *DisassociateSecurityGroupsRequest) (response *DisassociateSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDisassociateSecurityGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisassociateSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisassociateSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewEnableRoutesRequest() (request *EnableRoutesRequest) {
    request = &EnableRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "EnableRoutes")
    
    
    return
}

func NewEnableRoutesResponse() (response *EnableRoutesResponse) {
    response = &EnableRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableRoutes
// This API is used to enable a disabled subnet route.
//
// This API verifies whether a CCN route will conflict with an existing route after it is enabled, and if so, you cannot enable it before you disable the conflicting route first.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ECMP = "UnsupportedOperation.Ecmp"
//  UNSUPPORTEDOPERATION_ECMPWITHCCNROUTE = "UnsupportedOperation.EcmpWithCcnRoute"
//  UNSUPPORTEDOPERATION_ECMPWITHUSERROUTE = "UnsupportedOperation.EcmpWithUserRoute"
//  UNSUPPORTEDOPERATION_SYSTEMROUTE = "UnsupportedOperation.SystemRoute"
func (c *Client) EnableRoutes(request *EnableRoutesRequest) (response *EnableRoutesResponse, err error) {
    return c.EnableRoutesWithContext(context.Background(), request)
}

// EnableRoutes
// This API is used to enable a disabled subnet route.
//
// This API verifies whether a CCN route will conflict with an existing route after it is enabled, and if so, you cannot enable it before you disable the conflicting route first.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_COEXIST = "InvalidParameter.Coexist"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ECMP = "UnsupportedOperation.Ecmp"
//  UNSUPPORTEDOPERATION_ECMPWITHCCNROUTE = "UnsupportedOperation.EcmpWithCcnRoute"
//  UNSUPPORTEDOPERATION_ECMPWITHUSERROUTE = "UnsupportedOperation.EcmpWithUserRoute"
//  UNSUPPORTEDOPERATION_SYSTEMROUTE = "UnsupportedOperation.SystemRoute"
func (c *Client) EnableRoutesWithContext(ctx context.Context, request *EnableRoutesRequest) (response *EnableRoutesResponse, err error) {
    if request == nil {
        request = NewEnableRoutesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableRoutes require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewImportImageRequest() (request *ImportImageRequest) {
    request = &ImportImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "ImportImage")
    
    
    return
}

func NewImportImageResponse() (response *ImportImageResponse) {
    response = &ImportImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ImportImage
// This API is used to import an image from a CVM instance to an ECM instance.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_IMAGEDUPLICATE = "InvalidParameterValue.ImageDuplicate"
//  INVALIDPARAMETERVALUE_IMAGENAMEDUPLICATE = "InvalidParameterValue.ImageNameDuplicate"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageID"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDREGION = "InvalidParameterValue.UnsupportedRegion"
//  MISSINGPARAMETER_MISSINGIMAGEPARAMETER = "MissingParameter.MissingImageParameter"
//  RESOURCEINSUFFICIENT_INVAILDPRIVATEIMAGENUM = "ResourceInsufficient.InvaildPrivateImageNum"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ImportImage(request *ImportImageRequest) (response *ImportImageResponse, err error) {
    return c.ImportImageWithContext(context.Background(), request)
}

// ImportImage
// This API is used to import an image from a CVM instance to an ECM instance.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_IMAGEDUPLICATE = "InvalidParameterValue.ImageDuplicate"
//  INVALIDPARAMETERVALUE_IMAGENAMEDUPLICATE = "InvalidParameterValue.ImageNameDuplicate"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageID"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDREGION = "InvalidParameterValue.UnsupportedRegion"
//  MISSINGPARAMETER_MISSINGIMAGEPARAMETER = "MissingParameter.MissingImageParameter"
//  RESOURCEINSUFFICIENT_INVAILDPRIVATEIMAGENUM = "ResourceInsufficient.InvaildPrivateImageNum"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ImportImageWithContext(ctx context.Context, request *ImportImageRequest) (response *ImportImageResponse, err error) {
    if request == nil {
        request = NewImportImageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImportImage require credential")
    }

    request.SetContext(ctx)
    
    response = NewImportImageResponse()
    err = c.Send(request, response)
    return
}

func NewMigrateNetworkInterfaceRequest() (request *MigrateNetworkInterfaceRequest) {
    request = &MigrateNetworkInterfaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "MigrateNetworkInterface")
    
    
    return
}

func NewMigrateNetworkInterfaceResponse() (response *MigrateNetworkInterfaceResponse) {
    response = &MigrateNetworkInterfaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// MigrateNetworkInterface
// This API is used to migrate an ENI.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_ATTACHMENTNOTFOUND = "UnsupportedOperation.AttachmentNotFound"
func (c *Client) MigrateNetworkInterface(request *MigrateNetworkInterfaceRequest) (response *MigrateNetworkInterfaceResponse, err error) {
    return c.MigrateNetworkInterfaceWithContext(context.Background(), request)
}

// MigrateNetworkInterface
// This API is used to migrate an ENI.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_ATTACHMENTNOTFOUND = "UnsupportedOperation.AttachmentNotFound"
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
    
    request.Init().WithApiInfo("ecm", APIVersion, "MigratePrivateIpAddress")
    
    
    return
}

func NewMigratePrivateIpAddressResponse() (response *MigratePrivateIpAddressResponse) {
    response = &MigratePrivateIpAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// MigratePrivateIpAddress
// This API is used to migrate a private IP from an ENI.
//
// It migrates a private IP from one ENI to another. Primary IPs cannot be migrated.
//
// The source and destination ENIs must be in the same subnet.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) MigratePrivateIpAddress(request *MigratePrivateIpAddressRequest) (response *MigratePrivateIpAddressResponse, err error) {
    return c.MigratePrivateIpAddressWithContext(context.Background(), request)
}

// MigratePrivateIpAddress
// This API is used to migrate a private IP from an ENI.
//
// It migrates a private IP from one ENI to another. Primary IPs cannot be migrated.
//
// The source and destination ENIs must be in the same subnet.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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
    
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyAddressAttribute")
    
    
    return
}

func NewModifyAddressAttributeResponse() (response *ModifyAddressAttributeResponse) {
    response = &ModifyAddressAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAddressAttribute
// This API is used to modify EIP attributes.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_ADDRESSIDNOTFOUND = "UnsupportedOperation.AddressIdNotFound"
//  UNSUPPORTEDOPERATION_MALFORMED = "UnsupportedOperation.Malformed"
//  UNSUPPORTEDOPERATION_STATUSNOTPERMIT = "UnsupportedOperation.StatusNotPermit"
func (c *Client) ModifyAddressAttribute(request *ModifyAddressAttributeRequest) (response *ModifyAddressAttributeResponse, err error) {
    return c.ModifyAddressAttributeWithContext(context.Background(), request)
}

// ModifyAddressAttribute
// This API is used to modify EIP attributes.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_ADDRESSIDNOTFOUND = "UnsupportedOperation.AddressIdNotFound"
//  UNSUPPORTEDOPERATION_MALFORMED = "UnsupportedOperation.Malformed"
//  UNSUPPORTEDOPERATION_STATUSNOTPERMIT = "UnsupportedOperation.StatusNotPermit"
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

func NewModifyAddressesBandwidthRequest() (request *ModifyAddressesBandwidthRequest) {
    request = &ModifyAddressesBandwidthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyAddressesBandwidth")
    
    
    return
}

func NewModifyAddressesBandwidthResponse() (response *ModifyAddressesBandwidthResponse) {
    response = &ModifyAddressesBandwidthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAddressesBandwidth
// This API is used to modify the EIP bandwidth.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BANDWIDTHOUTOFRANGE = "InvalidParameterValue.BandwidthOutOfRange"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTH = "InvalidParameterValue.InvalidBandwidth"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_ADDRESSIDNOTFOUND = "UnsupportedOperation.AddressIdNotFound"
//  UNSUPPORTEDOPERATION_STATUSNOTPERMIT = "UnsupportedOperation.StatusNotPermit"
func (c *Client) ModifyAddressesBandwidth(request *ModifyAddressesBandwidthRequest) (response *ModifyAddressesBandwidthResponse, err error) {
    return c.ModifyAddressesBandwidthWithContext(context.Background(), request)
}

// ModifyAddressesBandwidth
// This API is used to modify the EIP bandwidth.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BANDWIDTHOUTOFRANGE = "InvalidParameterValue.BandwidthOutOfRange"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTH = "InvalidParameterValue.InvalidBandwidth"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_ADDRESSIDNOTFOUND = "UnsupportedOperation.AddressIdNotFound"
//  UNSUPPORTEDOPERATION_STATUSNOTPERMIT = "UnsupportedOperation.StatusNotPermit"
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

func NewModifyDefaultSubnetRequest() (request *ModifyDefaultSubnetRequest) {
    request = &ModifyDefaultSubnetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyDefaultSubnet")
    
    
    return
}

func NewModifyDefaultSubnetResponse() (response *ModifyDefaultSubnetResponse) {
    response = &ModifyDefaultSubnetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDefaultSubnet
// This API is used to modify the default subnet used when you create an instance in an AZ (if you don't specify the VPC parameter when creating the instance, `sunbetId` will be used).
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ModifyDefaultSubnet(request *ModifyDefaultSubnetRequest) (response *ModifyDefaultSubnetResponse, err error) {
    return c.ModifyDefaultSubnetWithContext(context.Background(), request)
}

// ModifyDefaultSubnet
// This API is used to modify the default subnet used when you create an instance in an AZ (if you don't specify the VPC parameter when creating the instance, `sunbetId` will be used).
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ModifyDefaultSubnetWithContext(ctx context.Context, request *ModifyDefaultSubnetRequest) (response *ModifyDefaultSubnetResponse, err error) {
    if request == nil {
        request = NewModifyDefaultSubnetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDefaultSubnet require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDefaultSubnetResponse()
    err = c.Send(request, response)
    return
}

func NewModifyHaVipAttributeRequest() (request *ModifyHaVipAttributeRequest) {
    request = &ModifyHaVipAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyHaVipAttribute")
    
    
    return
}

func NewModifyHaVipAttributeResponse() (response *ModifyHaVipAttributeResponse) {
    response = &ModifyHaVipAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyHaVipAttribute
// This API is used to modify the attributes of an HAVIP.
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
// This API is used to modify the attributes of an HAVIP.
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

func NewModifyImageAttributeRequest() (request *ModifyImageAttributeRequest) {
    request = &ModifyImageAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyImageAttribute")
    
    
    return
}

func NewModifyImageAttributeResponse() (response *ModifyImageAttributeResponse) {
    response = &ModifyImageAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyImageAttribute
// This API is used to modify the attributes of an image.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_IMAGENAMEDUPLICATE = "InvalidParameterValue.ImageNameDuplicate"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGIMAGEPARAMETER = "MissingParameter.MissingImageParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyImageAttribute(request *ModifyImageAttributeRequest) (response *ModifyImageAttributeResponse, err error) {
    return c.ModifyImageAttributeWithContext(context.Background(), request)
}

// ModifyImageAttribute
// This API is used to modify the attributes of an image.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_IMAGENAMEDUPLICATE = "InvalidParameterValue.ImageNameDuplicate"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGIMAGEPARAMETER = "MissingParameter.MissingImageParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyImageAttributeWithContext(ctx context.Context, request *ModifyImageAttributeRequest) (response *ModifyImageAttributeResponse, err error) {
    if request == nil {
        request = NewModifyImageAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyImageAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyImageAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstancesAttributeRequest() (request *ModifyInstancesAttributeRequest) {
    request = &ModifyInstancesAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyInstancesAttribute")
    
    
    return
}

func NewModifyInstancesAttributeResponse() (response *ModifyInstancesAttributeResponse) {
    response = &ModifyInstancesAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstancesAttribute
// This API is used to modify the attributes of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCEOWNERCHECKFAILED = "FailedOperation.InstanceOwnerCheckFailed"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  LIMITEXCEEDED_SECURITYGROUPINSTANCELIMITEXCEEDED = "LimitExceeded.SecurityGroupInstanceLimitExceeded"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ModifyInstancesAttribute(request *ModifyInstancesAttributeRequest) (response *ModifyInstancesAttributeResponse, err error) {
    return c.ModifyInstancesAttributeWithContext(context.Background(), request)
}

// ModifyInstancesAttribute
// This API is used to modify the attributes of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCEOWNERCHECKFAILED = "FailedOperation.InstanceOwnerCheckFailed"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  LIMITEXCEEDED_SECURITYGROUPINSTANCELIMITEXCEEDED = "LimitExceeded.SecurityGroupInstanceLimitExceeded"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ModifyInstancesAttributeWithContext(ctx context.Context, request *ModifyInstancesAttributeRequest) (response *ModifyInstancesAttributeResponse, err error) {
    if request == nil {
        request = NewModifyInstancesAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstancesAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstancesAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyIpv6AddressesAttributeRequest() (request *ModifyIpv6AddressesAttributeRequest) {
    request = &ModifyIpv6AddressesAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyIpv6AddressesAttribute")
    
    
    return
}

func NewModifyIpv6AddressesAttributeResponse() (response *ModifyIpv6AddressesAttributeResponse) {
    response = &ModifyIpv6AddressesAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyIpv6AddressesAttribute
// This API is used to modify the IPv6 address attributes of an ENI.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_ATTACHMENTNOTFOUND = "UnsupportedOperation.AttachmentNotFound"
func (c *Client) ModifyIpv6AddressesAttribute(request *ModifyIpv6AddressesAttributeRequest) (response *ModifyIpv6AddressesAttributeResponse, err error) {
    return c.ModifyIpv6AddressesAttributeWithContext(context.Background(), request)
}

// ModifyIpv6AddressesAttribute
// This API is used to modify the IPv6 address attributes of an ENI.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_ATTACHMENTNOTFOUND = "UnsupportedOperation.AttachmentNotFound"
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

func NewModifyListenerRequest() (request *ModifyListenerRequest) {
    request = &ModifyListenerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyListener")
    
    
    return
}

func NewModifyListenerResponse() (response *ModifyListenerResponse) {
    response = &ModifyListenerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyListener
// This API is used to modify the attributes of a CLB listener.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ModifyListener(request *ModifyListenerRequest) (response *ModifyListenerResponse, err error) {
    return c.ModifyListenerWithContext(context.Background(), request)
}

// ModifyListener
// This API is used to modify the attributes of a CLB listener.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
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
    
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyLoadBalancerAttributes")
    
    
    return
}

func NewModifyLoadBalancerAttributesResponse() (response *ModifyLoadBalancerAttributesResponse) {
    response = &ModifyLoadBalancerAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLoadBalancerAttributes
// This API is used to modify the attributes of a CLB instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ModifyLoadBalancerAttributes(request *ModifyLoadBalancerAttributesRequest) (response *ModifyLoadBalancerAttributesResponse, err error) {
    return c.ModifyLoadBalancerAttributesWithContext(context.Background(), request)
}

// ModifyLoadBalancerAttributes
// This API is used to modify the attributes of a CLB instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
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

func NewModifyModuleConfigRequest() (request *ModifyModuleConfigRequest) {
    request = &ModifyModuleConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyModuleConfig")
    
    
    return
}

func NewModifyModuleConfigResponse() (response *ModifyModuleConfigResponse) {
    response = &ModifyModuleConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyModuleConfig
// This API is used to modify the configuration of a module. You cannot modify the configuration of the module if it is associated with an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVAILDMODIFYPARAM = "InvalidParameterValue.InvaildModifyParam"
//  INVALIDPARAMETERVALUE_INVALIDDATADISKSIZE = "InvalidParameterValue.InvalidDataDiskSize"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEARCHITECTURE = "InvalidParameterValue.InvalidImageArchitecture"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCETYPE = "InvalidParameterValue.InvalidInstanceType"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCETYPECONFIGID = "InvalidParameterValue.InvalidInstanceTypeConfigID"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_MODULENOTALLOWCHANGE = "InvalidParameterValue.ModuleNotAllowChange"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ModifyModuleConfig(request *ModifyModuleConfigRequest) (response *ModifyModuleConfigResponse, err error) {
    return c.ModifyModuleConfigWithContext(context.Background(), request)
}

// ModifyModuleConfig
// This API is used to modify the configuration of a module. You cannot modify the configuration of the module if it is associated with an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVAILDMODIFYPARAM = "InvalidParameterValue.InvaildModifyParam"
//  INVALIDPARAMETERVALUE_INVALIDDATADISKSIZE = "InvalidParameterValue.InvalidDataDiskSize"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEARCHITECTURE = "InvalidParameterValue.InvalidImageArchitecture"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCETYPE = "InvalidParameterValue.InvalidInstanceType"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCETYPECONFIGID = "InvalidParameterValue.InvalidInstanceTypeConfigID"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_MODULENOTALLOWCHANGE = "InvalidParameterValue.ModuleNotAllowChange"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ModifyModuleConfigWithContext(ctx context.Context, request *ModifyModuleConfigRequest) (response *ModifyModuleConfigResponse, err error) {
    if request == nil {
        request = NewModifyModuleConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyModuleConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyModuleConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyModuleDisableWanIpRequest() (request *ModifyModuleDisableWanIpRequest) {
    request = &ModifyModuleDisableWanIpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyModuleDisableWanIp")
    
    
    return
}

func NewModifyModuleDisableWanIpResponse() (response *ModifyModuleDisableWanIpResponse) {
    response = &ModifyModuleDisableWanIpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyModuleDisableWanIp
// This API is used to specify whether to prohibit public IP assignment for a module.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
func (c *Client) ModifyModuleDisableWanIp(request *ModifyModuleDisableWanIpRequest) (response *ModifyModuleDisableWanIpResponse, err error) {
    return c.ModifyModuleDisableWanIpWithContext(context.Background(), request)
}

// ModifyModuleDisableWanIp
// This API is used to specify whether to prohibit public IP assignment for a module.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
func (c *Client) ModifyModuleDisableWanIpWithContext(ctx context.Context, request *ModifyModuleDisableWanIpRequest) (response *ModifyModuleDisableWanIpResponse, err error) {
    if request == nil {
        request = NewModifyModuleDisableWanIpRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyModuleDisableWanIp require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyModuleDisableWanIpResponse()
    err = c.Send(request, response)
    return
}

func NewModifyModuleImageRequest() (request *ModifyModuleImageRequest) {
    request = &ModifyModuleImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyModuleImage")
    
    
    return
}

func NewModifyModuleImageResponse() (response *ModifyModuleImageResponse) {
    response = &ModifyModuleImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyModuleImage
// This API is used to modify the default image of a module.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVAILDMODIFYPARAM = "InvalidParameterValue.InvaildModifyParam"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEARCHITECTURE = "InvalidParameterValue.InvalidImageArchitecture"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageID"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_INSTANCETYPENOTSUPPORTIMAGE = "UnsupportedOperation.InstanceTypeNotSupportImage"
func (c *Client) ModifyModuleImage(request *ModifyModuleImageRequest) (response *ModifyModuleImageResponse, err error) {
    return c.ModifyModuleImageWithContext(context.Background(), request)
}

// ModifyModuleImage
// This API is used to modify the default image of a module.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVAILDMODIFYPARAM = "InvalidParameterValue.InvaildModifyParam"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEARCHITECTURE = "InvalidParameterValue.InvalidImageArchitecture"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageID"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_INSTANCETYPENOTSUPPORTIMAGE = "UnsupportedOperation.InstanceTypeNotSupportImage"
func (c *Client) ModifyModuleImageWithContext(ctx context.Context, request *ModifyModuleImageRequest) (response *ModifyModuleImageResponse, err error) {
    if request == nil {
        request = NewModifyModuleImageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyModuleImage require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyModuleImageResponse()
    err = c.Send(request, response)
    return
}

func NewModifyModuleIpDirectRequest() (request *ModifyModuleIpDirectRequest) {
    request = &ModifyModuleIpDirectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyModuleIpDirect")
    
    
    return
}

func NewModifyModuleIpDirectResponse() (response *ModifyModuleIpDirectResponse) {
    response = &ModifyModuleIpDirectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyModuleIpDirect
// This API is used to modify the IP direct access of a module.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDEIPDIRECTSERVICE = "InvalidParameterValue.InvalidEIPDirectService"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
func (c *Client) ModifyModuleIpDirect(request *ModifyModuleIpDirectRequest) (response *ModifyModuleIpDirectResponse, err error) {
    return c.ModifyModuleIpDirectWithContext(context.Background(), request)
}

// ModifyModuleIpDirect
// This API is used to modify the IP direct access of a module.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDEIPDIRECTSERVICE = "InvalidParameterValue.InvalidEIPDirectService"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
func (c *Client) ModifyModuleIpDirectWithContext(ctx context.Context, request *ModifyModuleIpDirectRequest) (response *ModifyModuleIpDirectResponse, err error) {
    if request == nil {
        request = NewModifyModuleIpDirectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyModuleIpDirect require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyModuleIpDirectResponse()
    err = c.Send(request, response)
    return
}

func NewModifyModuleNameRequest() (request *ModifyModuleNameRequest) {
    request = &ModifyModuleNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyModuleName")
    
    
    return
}

func NewModifyModuleNameResponse() (response *ModifyModuleNameResponse) {
    response = &ModifyModuleNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyModuleName
// This API is used to rename a module.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  INVALIDPARAMETERVALUE_INVALIDMODULENAME = "InvalidParameterValue.InvalidModuleName"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
func (c *Client) ModifyModuleName(request *ModifyModuleNameRequest) (response *ModifyModuleNameResponse, err error) {
    return c.ModifyModuleNameWithContext(context.Background(), request)
}

// ModifyModuleName
// This API is used to rename a module.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  INVALIDPARAMETERVALUE_INVALIDMODULENAME = "InvalidParameterValue.InvalidModuleName"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
func (c *Client) ModifyModuleNameWithContext(ctx context.Context, request *ModifyModuleNameRequest) (response *ModifyModuleNameResponse, err error) {
    if request == nil {
        request = NewModifyModuleNameRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyModuleName require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyModuleNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyModuleNetworkRequest() (request *ModifyModuleNetworkRequest) {
    request = &ModifyModuleNetworkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyModuleNetwork")
    
    
    return
}

func NewModifyModuleNetworkResponse() (response *ModifyModuleNetworkResponse) {
    response = &ModifyModuleNetworkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyModuleNetwork
// This API is used to modify the default bandwidth cap of a module.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVAILDMODIFYPARAM = "InvalidParameterValue.InvaildModifyParam"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTH = "InvalidParameterValue.InvalidBandwidth"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTHIN = "InvalidParameterValue.InvalidBandwidthIn"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTHINANDOUT = "InvalidParameterValue.InvalidBandwidthInAndOut"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ModifyModuleNetwork(request *ModifyModuleNetworkRequest) (response *ModifyModuleNetworkResponse, err error) {
    return c.ModifyModuleNetworkWithContext(context.Background(), request)
}

// ModifyModuleNetwork
// This API is used to modify the default bandwidth cap of a module.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVAILDMODIFYPARAM = "InvalidParameterValue.InvaildModifyParam"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTH = "InvalidParameterValue.InvalidBandwidth"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTHIN = "InvalidParameterValue.InvalidBandwidthIn"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTHINANDOUT = "InvalidParameterValue.InvalidBandwidthInAndOut"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ModifyModuleNetworkWithContext(ctx context.Context, request *ModifyModuleNetworkRequest) (response *ModifyModuleNetworkResponse, err error) {
    if request == nil {
        request = NewModifyModuleNetworkRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyModuleNetwork require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyModuleNetworkResponse()
    err = c.Send(request, response)
    return
}

func NewModifyModuleSecurityGroupsRequest() (request *ModifyModuleSecurityGroupsRequest) {
    request = &ModifyModuleSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyModuleSecurityGroups")
    
    
    return
}

func NewModifyModuleSecurityGroupsResponse() (response *ModifyModuleSecurityGroupsResponse) {
    response = &ModifyModuleSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyModuleSecurityGroups
// This API is used to modify the default security group of a module.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  LIMITEXCEEDED_MODULESECURITYGROUPLIMITEXCEEDED = "LimitExceeded.ModuleSecurityGroupLimitExceeded"
//  LIMITEXCEEDED_SECURITYGROUPMODULELIMITEXCEEDED = "LimitExceeded.SecurityGroupModuleLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyModuleSecurityGroups(request *ModifyModuleSecurityGroupsRequest) (response *ModifyModuleSecurityGroupsResponse, err error) {
    return c.ModifyModuleSecurityGroupsWithContext(context.Background(), request)
}

// ModifyModuleSecurityGroups
// This API is used to modify the default security group of a module.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  LIMITEXCEEDED_MODULESECURITYGROUPLIMITEXCEEDED = "LimitExceeded.ModuleSecurityGroupLimitExceeded"
//  LIMITEXCEEDED_SECURITYGROUPMODULELIMITEXCEEDED = "LimitExceeded.SecurityGroupModuleLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyModuleSecurityGroupsWithContext(ctx context.Context, request *ModifyModuleSecurityGroupsRequest) (response *ModifyModuleSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewModifyModuleSecurityGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyModuleSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyModuleSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrivateIpAddressesAttributeRequest() (request *ModifyPrivateIpAddressesAttributeRequest) {
    request = &ModifyPrivateIpAddressesAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyPrivateIpAddressesAttribute")
    
    
    return
}

func NewModifyPrivateIpAddressesAttributeResponse() (response *ModifyPrivateIpAddressesAttributeResponse) {
    response = &ModifyPrivateIpAddressesAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPrivateIpAddressesAttribute
// This API is used to modify the private IP attributes of an ENI.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_INVALIDSTATE = "UnsupportedOperation.InvalidState"
func (c *Client) ModifyPrivateIpAddressesAttribute(request *ModifyPrivateIpAddressesAttributeRequest) (response *ModifyPrivateIpAddressesAttributeResponse, err error) {
    return c.ModifyPrivateIpAddressesAttributeWithContext(context.Background(), request)
}

// ModifyPrivateIpAddressesAttribute
// This API is used to modify the private IP attributes of an ENI.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
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
    
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyRouteTableAttribute")
    
    
    return
}

func NewModifyRouteTableAttributeResponse() (response *ModifyRouteTableAttributeResponse) {
    response = &ModifyRouteTableAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRouteTableAttribute
// This API is used to modify the attributes of a route table.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyRouteTableAttribute(request *ModifyRouteTableAttributeRequest) (response *ModifyRouteTableAttributeResponse, err error) {
    return c.ModifyRouteTableAttributeWithContext(context.Background(), request)
}

// ModifyRouteTableAttribute
// This API is used to modify the attributes of a route table.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
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
    
    request.Init().WithApiInfo("ecm", APIVersion, "ModifySecurityGroupAttribute")
    
    
    return
}

func NewModifySecurityGroupAttributeResponse() (response *ModifySecurityGroupAttributeResponse) {
    response = &ModifySecurityGroupAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySecurityGroupAttribute
// This API is used to modify the attributes of a security group.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifySecurityGroupAttribute(request *ModifySecurityGroupAttributeRequest) (response *ModifySecurityGroupAttributeResponse, err error) {
    return c.ModifySecurityGroupAttributeWithContext(context.Background(), request)
}

// ModifySecurityGroupAttribute
// This API is used to modify the attributes of a security group.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  MISSINGPARAMETER = "MissingParameter"
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
    
    request.Init().WithApiInfo("ecm", APIVersion, "ModifySecurityGroupPolicies")
    
    
    return
}

func NewModifySecurityGroupPoliciesResponse() (response *ModifySecurityGroupPoliciesResponse) {
    response = &ModifySecurityGroupPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySecurityGroupPolicies
// This API is used to modify the outbound and inbound rules of a security group.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifySecurityGroupPolicies(request *ModifySecurityGroupPoliciesRequest) (response *ModifySecurityGroupPoliciesResponse, err error) {
    return c.ModifySecurityGroupPoliciesWithContext(context.Background(), request)
}

// ModifySecurityGroupPolicies
// This API is used to modify the outbound and inbound rules of a security group.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
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

func NewModifySubnetAttributeRequest() (request *ModifySubnetAttributeRequest) {
    request = &ModifySubnetAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "ModifySubnetAttribute")
    
    
    return
}

func NewModifySubnetAttributeResponse() (response *ModifySubnetAttributeResponse) {
    response = &ModifySubnetAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySubnetAttribute
// This API is used to modify the attributes of a subnet.
//
// error code that may be returned:
//  FAILEDOPERATION_DATAOPERATIONFAILED = "FailedOperation.DataOperationFailed"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ModifySubnetAttribute(request *ModifySubnetAttributeRequest) (response *ModifySubnetAttributeResponse, err error) {
    return c.ModifySubnetAttributeWithContext(context.Background(), request)
}

// ModifySubnetAttribute
// This API is used to modify the attributes of a subnet.
//
// error code that may be returned:
//  FAILEDOPERATION_DATAOPERATIONFAILED = "FailedOperation.DataOperationFailed"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
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

func NewModifyTargetPortRequest() (request *ModifyTargetPortRequest) {
    request = &ModifyTargetPortRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyTargetPort")
    
    
    return
}

func NewModifyTargetPortResponse() (response *ModifyTargetPortResponse) {
    response = &ModifyTargetPortResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTargetPort
// This API is used to modify the port of a real server bound to a listener.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ModifyTargetPort(request *ModifyTargetPortRequest) (response *ModifyTargetPortResponse, err error) {
    return c.ModifyTargetPortWithContext(context.Background(), request)
}

// ModifyTargetPort
// This API is used to modify the port of a real server bound to a listener.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
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
    
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyTargetWeight")
    
    
    return
}

func NewModifyTargetWeightResponse() (response *ModifyTargetWeightResponse) {
    response = &ModifyTargetWeightResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTargetWeight
// This API is used to modify the forwarding weight of a real server bound to a listener.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ModifyTargetWeight(request *ModifyTargetWeightRequest) (response *ModifyTargetWeightResponse, err error) {
    return c.ModifyTargetWeightWithContext(context.Background(), request)
}

// ModifyTargetWeight
// This API is used to modify the forwarding weight of a real server bound to a listener.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
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

func NewModifyVpcAttributeRequest() (request *ModifyVpcAttributeRequest) {
    request = &ModifyVpcAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyVpcAttribute")
    
    
    return
}

func NewModifyVpcAttributeResponse() (response *ModifyVpcAttributeResponse) {
    response = &ModifyVpcAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyVpcAttribute
// This API is used to modify the attributes of a VPC.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ModifyVpcAttribute(request *ModifyVpcAttributeRequest) (response *ModifyVpcAttributeResponse, err error) {
    return c.ModifyVpcAttributeWithContext(context.Background(), request)
}

// ModifyVpcAttribute
// This API is used to modify the attributes of a VPC.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
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

func NewRebootInstancesRequest() (request *RebootInstancesRequest) {
    request = &RebootInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "RebootInstances")
    
    
    return
}

func NewRebootInstancesResponse() (response *RebootInstancesResponse) {
    response = &RebootInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RebootInstances
// This API is used to restart an instance. Only instances in `RUNNING` status can be restarted. The instance status will become `REBOOTING` when the API is called successfully and will become `RUNNING` when the instance is restarted successfully. Forced restart is supported. Just like powering off a physical PC and restarting it, a forced restart may cause data loss or file system corruption. Make sure that you use this API only when the server cannot be restarted normally.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDDATADISKTYPE = "InvalidParameterValue.InvalidDataDiskType"
//  INVALIDPARAMETERVALUE_INVALIDISPINNODE = "InvalidParameterValue.InvalidISPInNode"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  RESOURCEUNAVAILABLE_INSTANCENOTRUNNING = "ResourceUnavailable.InstanceNotRunning"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) RebootInstances(request *RebootInstancesRequest) (response *RebootInstancesResponse, err error) {
    return c.RebootInstancesWithContext(context.Background(), request)
}

// RebootInstances
// This API is used to restart an instance. Only instances in `RUNNING` status can be restarted. The instance status will become `REBOOTING` when the API is called successfully and will become `RUNNING` when the instance is restarted successfully. Forced restart is supported. Just like powering off a physical PC and restarting it, a forced restart may cause data loss or file system corruption. Make sure that you use this API only when the server cannot be restarted normally.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDDATADISKTYPE = "InvalidParameterValue.InvalidDataDiskType"
//  INVALIDPARAMETERVALUE_INVALIDISPINNODE = "InvalidParameterValue.InvalidISPInNode"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  RESOURCEUNAVAILABLE_INSTANCENOTRUNNING = "ResourceUnavailable.InstanceNotRunning"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) RebootInstancesWithContext(ctx context.Context, request *RebootInstancesRequest) (response *RebootInstancesResponse, err error) {
    if request == nil {
        request = NewRebootInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RebootInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewRebootInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseAddressesRequest() (request *ReleaseAddressesRequest) {
    request = &ReleaseAddressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "ReleaseAddresses")
    
    
    return
}

func NewReleaseAddressesResponse() (response *ReleaseAddressesResponse) {
    response = &ReleaseAddressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReleaseAddresses
// This API is used to release one or multiple EIPs.
//
// This operation is irreversible. Once you release an EIP, the IP address associated with it will no longer belong to you.
//
// Only EIPs in `UNBIND` status can be released.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_ADDRESSIDNOTFOUND = "UnsupportedOperation.AddressIdNotFound"
//  UNSUPPORTEDOPERATION_MALFORMED = "UnsupportedOperation.Malformed"
//  UNSUPPORTEDOPERATION_STATUSNOTPERMIT = "UnsupportedOperation.StatusNotPermit"
func (c *Client) ReleaseAddresses(request *ReleaseAddressesRequest) (response *ReleaseAddressesResponse, err error) {
    return c.ReleaseAddressesWithContext(context.Background(), request)
}

// ReleaseAddresses
// This API is used to release one or multiple EIPs.
//
// This operation is irreversible. Once you release an EIP, the IP address associated with it will no longer belong to you.
//
// Only EIPs in `UNBIND` status can be released.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_ADDRESSIDNOTFOUND = "UnsupportedOperation.AddressIdNotFound"
//  UNSUPPORTEDOPERATION_MALFORMED = "UnsupportedOperation.Malformed"
//  UNSUPPORTEDOPERATION_STATUSNOTPERMIT = "UnsupportedOperation.StatusNotPermit"
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

func NewReleaseIpv6AddressesRequest() (request *ReleaseIpv6AddressesRequest) {
    request = &ReleaseIpv6AddressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "ReleaseIpv6Addresses")
    
    
    return
}

func NewReleaseIpv6AddressesResponse() (response *ReleaseIpv6AddressesResponse) {
    response = &ReleaseIpv6AddressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReleaseIpv6Addresses
// This API is used to release the IPv6 addresses of an ENI.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_ATTACHMENTNOTFOUND = "UnsupportedOperation.AttachmentNotFound"
func (c *Client) ReleaseIpv6Addresses(request *ReleaseIpv6AddressesRequest) (response *ReleaseIpv6AddressesResponse, err error) {
    return c.ReleaseIpv6AddressesWithContext(context.Background(), request)
}

// ReleaseIpv6Addresses
// This API is used to release the IPv6 addresses of an ENI.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER_MISSINGMODULEPARAMETER = "MissingParameter.MissingModuleParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_ATTACHMENTNOTFOUND = "UnsupportedOperation.AttachmentNotFound"
func (c *Client) ReleaseIpv6AddressesWithContext(ctx context.Context, request *ReleaseIpv6AddressesRequest) (response *ReleaseIpv6AddressesResponse, err error) {
    if request == nil {
        request = NewReleaseIpv6AddressesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReleaseIpv6Addresses require credential")
    }

    request.SetContext(ctx)
    
    response = NewReleaseIpv6AddressesResponse()
    err = c.Send(request, response)
    return
}

func NewRemovePrivateIpAddressesRequest() (request *RemovePrivateIpAddressesRequest) {
    request = &RemovePrivateIpAddressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "RemovePrivateIpAddresses")
    
    
    return
}

func NewRemovePrivateIpAddressesResponse() (response *RemovePrivateIpAddressesResponse) {
    response = &RemovePrivateIpAddressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemovePrivateIpAddresses
// This API is used to return the private IPs of an ENI.
//
// To return the secondary private IPs of an ENI, the API will automatically unbind them from the ENI. The primary private IP of the ENI cannot be returned.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ATTACHMENTNOTFOUND = "UnsupportedOperation.AttachmentNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) RemovePrivateIpAddresses(request *RemovePrivateIpAddressesRequest) (response *RemovePrivateIpAddressesResponse, err error) {
    return c.RemovePrivateIpAddressesWithContext(context.Background(), request)
}

// RemovePrivateIpAddresses
// This API is used to return the private IPs of an ENI.
//
// To return the secondary private IPs of an ENI, the API will automatically unbind them from the ENI. The primary private IP of the ENI cannot be returned.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ATTACHMENTNOTFOUND = "UnsupportedOperation.AttachmentNotFound"
//  UNSUPPORTEDOPERATION_MUTEXOPERATIONTASKRUNNING = "UnsupportedOperation.MutexOperationTaskRunning"
func (c *Client) RemovePrivateIpAddressesWithContext(ctx context.Context, request *RemovePrivateIpAddressesRequest) (response *RemovePrivateIpAddressesResponse, err error) {
    if request == nil {
        request = NewRemovePrivateIpAddressesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemovePrivateIpAddresses require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemovePrivateIpAddressesResponse()
    err = c.Send(request, response)
    return
}

func NewReplaceRouteTableAssociationRequest() (request *ReplaceRouteTableAssociationRequest) {
    request = &ReplaceRouteTableAssociationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "ReplaceRouteTableAssociation")
    
    
    return
}

func NewReplaceRouteTableAssociationResponse() (response *ReplaceRouteTableAssociationResponse) {
    response = &ReplaceRouteTableAssociationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReplaceRouteTableAssociation
// This API is used to modify the route table associated with a subnet. A subnet can be associated with only one route table.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_OBJECTVPCNOTCURRENTVPC = "InvalidParameterValue.ObjectVpcNotCurrentVpc"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_VPCMISMATCH = "UnsupportedOperation.VpcMismatch"
func (c *Client) ReplaceRouteTableAssociation(request *ReplaceRouteTableAssociationRequest) (response *ReplaceRouteTableAssociationResponse, err error) {
    return c.ReplaceRouteTableAssociationWithContext(context.Background(), request)
}

// ReplaceRouteTableAssociation
// This API is used to modify the route table associated with a subnet. A subnet can be associated with only one route table.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_OBJECTVPCNOTCURRENTVPC = "InvalidParameterValue.ObjectVpcNotCurrentVpc"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
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
    
    request.Init().WithApiInfo("ecm", APIVersion, "ReplaceRoutes")
    
    
    return
}

func NewReplaceRoutesResponse() (response *ReplaceRoutesResponse) {
    response = &ReplaceRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReplaceRoutes
// This API is used to replace a routing policy.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CIDRNOTINPEERVPC = "InvalidParameterValue.CidrNotInPeerVpc"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_VPCCIDRCONFLICT = "InvalidParameterValue.VpcCidrConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ECMP = "UnsupportedOperation.Ecmp"
//  UNSUPPORTEDOPERATION_SYSTEMROUTE = "UnsupportedOperation.SystemRoute"
func (c *Client) ReplaceRoutes(request *ReplaceRoutesRequest) (response *ReplaceRoutesResponse, err error) {
    return c.ReplaceRoutesWithContext(context.Background(), request)
}

// ReplaceRoutes
// This API is used to replace a routing policy.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CIDRNOTINPEERVPC = "InvalidParameterValue.CidrNotInPeerVpc"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  INVALIDPARAMETERVALUE_VPCCIDRCONFLICT = "InvalidParameterValue.VpcCidrConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ECMP = "UnsupportedOperation.Ecmp"
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
    
    request.Init().WithApiInfo("ecm", APIVersion, "ReplaceSecurityGroupPolicy")
    
    
    return
}

func NewReplaceSecurityGroupPolicyResponse() (response *ReplaceSecurityGroupPolicyResponse) {
    response = &ReplaceSecurityGroupPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReplaceSecurityGroupPolicy
// This API is used to replace a single routing rule of a security group. You can replace only one rule in a single direction in one request and must specify the index (PolicyIndex).
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_VERSIONMISMATCH = "UnsupportedOperation.VersionMismatch"
func (c *Client) ReplaceSecurityGroupPolicy(request *ReplaceSecurityGroupPolicyRequest) (response *ReplaceSecurityGroupPolicyResponse, err error) {
    return c.ReplaceSecurityGroupPolicyWithContext(context.Background(), request)
}

// ReplaceSecurityGroupPolicy
// This API is used to replace a single routing rule of a security group. You can replace only one rule in a single direction in one request and must specify the index (PolicyIndex).
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
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

func NewResetInstancesRequest() (request *ResetInstancesRequest) {
    request = &ResetInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "ResetInstances")
    
    
    return
}

func NewResetInstancesResponse() (response *ResetInstancesResponse) {
    response = &ResetInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetInstances
// This API is used to reinstall an instance. If you specify the `ImageId` parameter, the specified image will be used; otherwise, the image used by the current instance will be used. If you don't specify the password, a password will be sent later in Message Center.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE_INVAILDPASSWORD = "InvalidParameterValue.InvaildPassword"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEARCHITECTURE = "InvalidParameterValue.InvalidImageArchitecture"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageID"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ResetInstances(request *ResetInstancesRequest) (response *ResetInstancesResponse, err error) {
    return c.ResetInstancesWithContext(context.Background(), request)
}

// ResetInstances
// This API is used to reinstall an instance. If you specify the `ImageId` parameter, the specified image will be used; otherwise, the image used by the current instance will be used. If you don't specify the password, a password will be sent later in Message Center.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE_INVAILDPASSWORD = "InvalidParameterValue.InvaildPassword"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEARCHITECTURE = "InvalidParameterValue.InvalidImageArchitecture"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageID"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ResetInstancesWithContext(ctx context.Context, request *ResetInstancesRequest) (response *ResetInstancesResponse, err error) {
    if request == nil {
        request = NewResetInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewResetInstancesMaxBandwidthRequest() (request *ResetInstancesMaxBandwidthRequest) {
    request = &ResetInstancesMaxBandwidthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "ResetInstancesMaxBandwidth")
    
    
    return
}

func NewResetInstancesMaxBandwidthResponse() (response *ResetInstancesMaxBandwidthResponse) {
    response = &ResetInstancesMaxBandwidthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetInstancesMaxBandwidth
// This API is used to reset the bandwidth cap of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTH = "InvalidParameterValue.InvalidBandwidth"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTHIN = "InvalidParameterValue.InvalidBandwidthIn"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTHINANDOUT = "InvalidParameterValue.InvalidBandwidthInAndOut"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ResetInstancesMaxBandwidth(request *ResetInstancesMaxBandwidthRequest) (response *ResetInstancesMaxBandwidthResponse, err error) {
    return c.ResetInstancesMaxBandwidthWithContext(context.Background(), request)
}

// ResetInstancesMaxBandwidth
// This API is used to reset the bandwidth cap of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTH = "InvalidParameterValue.InvalidBandwidth"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTHIN = "InvalidParameterValue.InvalidBandwidthIn"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTHINANDOUT = "InvalidParameterValue.InvalidBandwidthInAndOut"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ResetInstancesMaxBandwidthWithContext(ctx context.Context, request *ResetInstancesMaxBandwidthRequest) (response *ResetInstancesMaxBandwidthResponse, err error) {
    if request == nil {
        request = NewResetInstancesMaxBandwidthRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetInstancesMaxBandwidth require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetInstancesMaxBandwidthResponse()
    err = c.Send(request, response)
    return
}

func NewResetInstancesPasswordRequest() (request *ResetInstancesPasswordRequest) {
    request = &ResetInstancesPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "ResetInstancesPassword")
    
    
    return
}

func NewResetInstancesPasswordResponse() (response *ResetInstancesPasswordResponse) {
    response = &ResetInstancesPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetInstancesPassword
// This API is used to reset the password for a running status. You need to explicitly specify the `ForceStop` parameter; otherwise, you can reset the password only for instances that have been shut down.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCENOTALLSTOPPED = "FailedOperation.InstanceNotAllStopped"
//  FAILEDOPERATION_INSTANCEOWNERCHECKFAILED = "FailedOperation.InstanceOwnerCheckFailed"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVAILDMODIFYPARAM = "InvalidParameterValue.InvaildModifyParam"
//  INVALIDPARAMETERVALUE_INVAILDPASSWORD = "InvalidParameterValue.InvaildPassword"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ResetInstancesPassword(request *ResetInstancesPasswordRequest) (response *ResetInstancesPasswordResponse, err error) {
    return c.ResetInstancesPasswordWithContext(context.Background(), request)
}

// ResetInstancesPassword
// This API is used to reset the password for a running status. You need to explicitly specify the `ForceStop` parameter; otherwise, you can reset the password only for instances that have been shut down.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCENOTALLSTOPPED = "FailedOperation.InstanceNotAllStopped"
//  FAILEDOPERATION_INSTANCEOWNERCHECKFAILED = "FailedOperation.InstanceOwnerCheckFailed"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVAILDMODIFYPARAM = "InvalidParameterValue.InvaildModifyParam"
//  INVALIDPARAMETERVALUE_INVAILDPASSWORD = "InvalidParameterValue.InvaildPassword"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) ResetInstancesPasswordWithContext(ctx context.Context, request *ResetInstancesPasswordRequest) (response *ResetInstancesPasswordResponse, err error) {
    if request == nil {
        request = NewResetInstancesPasswordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetInstancesPassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetInstancesPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewResetRoutesRequest() (request *ResetRoutesRequest) {
    request = &ResetRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "ResetRoutes")
    
    
    return
}

func NewResetRoutesResponse() (response *ResetRoutesResponse) {
    response = &ResetRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetRoutes
// This API is used to reset a route table name and all routing policies.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CIDRNOTINPEERVPC = "InvalidParameterValue.CidrNotInPeerVpc"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_VPCCIDRCONFLICT = "InvalidParameterValue.VpcCidrConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_SYSTEMROUTE = "UnsupportedOperation.SystemRoute"
func (c *Client) ResetRoutes(request *ResetRoutesRequest) (response *ResetRoutesResponse, err error) {
    return c.ResetRoutesWithContext(context.Background(), request)
}

// ResetRoutes
// This API is used to reset a route table name and all routing policies.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CIDRNOTINPEERVPC = "InvalidParameterValue.CidrNotInPeerVpc"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_VPCCIDRCONFLICT = "InvalidParameterValue.VpcCidrConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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

func NewRunInstancesRequest() (request *RunInstancesRequest) {
    request = &RunInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "RunInstances")
    
    
    return
}

func NewRunInstancesResponse() (response *RunInstancesResponse) {
    response = &RunInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RunInstances
// This API is used to create an ECM instance.
//
// error code that may be returned:
//  FAILEDOPERATION_BLOCKBALANCE = "FailedOperation.BlockBalance"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONCONFLICT = "FailedOperation.OperationConflict"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ADDRESSQUOTALIMITEXCEEDED = "InvalidParameterValue.AddressQuotaLimitExceeded"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_FILTERLIMITEXCEEDED = "InvalidParameterValue.FilterLimitExceeded"
//  INVALIDPARAMETERVALUE_IMAGESIZELARGETHANSYSDISKSIZE = "InvalidParameterValue.ImageSizeLargeThanSysDiskSize"
//  INVALIDPARAMETERVALUE_INSTANCECONFIGNOTMATCH = "InvalidParameterValue.InstanceConfigNotMatch"
//  INVALIDPARAMETERVALUE_INSTANCENAMETOOLONG = "InvalidParameterValue.InstanceNameTooLong"
//  INVALIDPARAMETERVALUE_INSTANCETYPENOTMATCHPID = "InvalidParameterValue.InstanceTypeNotMatchPid"
//  INVALIDPARAMETERVALUE_INVAILDHOSTNAME = "InvalidParameterValue.InvaildHostName"
//  INVALIDPARAMETERVALUE_INVAILDMODIFYPARAM = "InvalidParameterValue.InvaildModifyParam"
//  INVALIDPARAMETERVALUE_INVAILDPASSWORD = "InvalidParameterValue.InvaildPassword"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTH = "InvalidParameterValue.InvalidBandwidth"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTHIN = "InvalidParameterValue.InvalidBandwidthIn"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTHINANDOUT = "InvalidParameterValue.InvalidBandwidthInAndOut"
//  INVALIDPARAMETERVALUE_INVALIDBILLINGTYPE = "InvalidParameterValue.InvalidBillingType"
//  INVALIDPARAMETERVALUE_INVALIDDATADISKNUM = "InvalidParameterValue.InvalidDataDiskNum"
//  INVALIDPARAMETERVALUE_INVALIDDATADISKSIZE = "InvalidParameterValue.InvalidDataDiskSize"
//  INVALIDPARAMETERVALUE_INVALIDDATADISKTYPE = "InvalidParameterValue.InvalidDataDiskType"
//  INVALIDPARAMETERVALUE_INVALIDEIPDIRECTSERVICE = "InvalidParameterValue.InvalidEIPDirectService"
//  INVALIDPARAMETERVALUE_INVALIDISPINNODE = "InvalidParameterValue.InvalidISPInNode"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEARCHITECTURE = "InvalidParameterValue.InvalidImageArchitecture"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageID"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCECHARGETYPE = "InvalidParameterValue.InvalidInstanceChargeType"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCETYPECONFIGID = "InvalidParameterValue.InvalidInstanceTypeConfigID"
//  INVALIDPARAMETERVALUE_INVALIDINTERNETCHARGETYPE = "InvalidParameterValue.InvalidInternetChargeType"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  INVALIDPARAMETERVALUE_INVALIDMODULEIDANDINSTANCETYPECONFIGID = "InvalidParameterValue.InvalidModuleIDAndInstanceTypeConfigID"
//  INVALIDPARAMETERVALUE_INVALIDMODULEIDANDINSTANCETYPEID = "InvalidParameterValue.InvalidModuleIDAndInstanceTypeID"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDSECURITYGROUPID = "InvalidParameterValue.InvalidSecurityGroupID"
//  INVALIDPARAMETERVALUE_INVALIDSYSTEMDISKSIZE = "InvalidParameterValue.InvalidSystemDiskSize"
//  INVALIDPARAMETERVALUE_INVALIDSYSTEMDISKTYPE = "InvalidParameterValue.InvalidSystemDiskType"
//  INVALIDPARAMETERVALUE_INVALIDZONE = "InvalidParameterValue.InvalidZone"
//  INVALIDPARAMETERVALUE_INVALIDZONEINSTANCECOUNT = "InvalidParameterValue.InvalidZoneInstanceCount"
//  INVALIDPARAMETERVALUE_INVALIDZONEINSTANCETYPE = "InvalidParameterValue.InvalidZoneInstanceType"
//  INVALIDPARAMETERVALUE_NODENOTSUPPORTIPV6 = "InvalidParameterValue.NodeNotSupportIPv6"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  INVALIDPARAMETERVALUE_SUBNETNOTEXIST = "InvalidParameterValue.SubnetNotExist"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_UNMATCHEDBILLINGTYPE = "InvalidParameterValue.UnmatchedBillingType"
//  INVALIDPARAMETERVALUE_USERNOTSUPPORTIPV6 = "InvalidParameterValue.UserNotSupportIPv6"
//  LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.AddressQuotaLimitExceeded"
//  LIMITEXCEEDED_ENIQUOTALIMITEXCEEDED = "LimitExceeded.EniQuotaLimitExceeded"
//  LIMITEXCEEDED_INSTANCESECURITYGROUPLIMITEXCEEDED = "LimitExceeded.InstanceSecurityGroupLimitExceeded"
//  LIMITEXCEEDED_NICORIPLIMITEXCEEDED = "LimitExceeded.NicOrIPLimitExceeded"
//  LIMITEXCEEDED_SECURITYGROUPINSTANCELIMITEXCEEDED = "LimitExceeded.SecurityGroupInstanceLimitExceeded"
//  LIMITEXCEEDED_VCPULIMITEXCEEDED = "LimitExceeded.VcpuLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT_IPQUOTANOTENOUGH = "ResourceInsufficient.IPQuotaNotEnough"
//  RESOURCEINSUFFICIENT_INSTANCEQUOTANOTENOUGH = "ResourceInsufficient.InstanceQuotaNotEnough"
//  RESOURCEINSUFFICIENT_PRIVATEIPQUOTANOTENOUGH = "ResourceInsufficient.PrivateIPQuotaNotEnough"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCESSOLDOUT_SPECIFIEDINSTANCETYPE = "ResourcesSoldOut.SpecifiedInstanceType"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_INSTANCETYPENOTSUPPORTIMAGE = "UnsupportedOperation.InstanceTypeNotSupportImage"
func (c *Client) RunInstances(request *RunInstancesRequest) (response *RunInstancesResponse, err error) {
    return c.RunInstancesWithContext(context.Background(), request)
}

// RunInstances
// This API is used to create an ECM instance.
//
// error code that may be returned:
//  FAILEDOPERATION_BLOCKBALANCE = "FailedOperation.BlockBalance"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONCONFLICT = "FailedOperation.OperationConflict"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ADDRESSQUOTALIMITEXCEEDED = "InvalidParameterValue.AddressQuotaLimitExceeded"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_FILTERLIMITEXCEEDED = "InvalidParameterValue.FilterLimitExceeded"
//  INVALIDPARAMETERVALUE_IMAGESIZELARGETHANSYSDISKSIZE = "InvalidParameterValue.ImageSizeLargeThanSysDiskSize"
//  INVALIDPARAMETERVALUE_INSTANCECONFIGNOTMATCH = "InvalidParameterValue.InstanceConfigNotMatch"
//  INVALIDPARAMETERVALUE_INSTANCENAMETOOLONG = "InvalidParameterValue.InstanceNameTooLong"
//  INVALIDPARAMETERVALUE_INSTANCETYPENOTMATCHPID = "InvalidParameterValue.InstanceTypeNotMatchPid"
//  INVALIDPARAMETERVALUE_INVAILDHOSTNAME = "InvalidParameterValue.InvaildHostName"
//  INVALIDPARAMETERVALUE_INVAILDMODIFYPARAM = "InvalidParameterValue.InvaildModifyParam"
//  INVALIDPARAMETERVALUE_INVAILDPASSWORD = "InvalidParameterValue.InvaildPassword"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTH = "InvalidParameterValue.InvalidBandwidth"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTHIN = "InvalidParameterValue.InvalidBandwidthIn"
//  INVALIDPARAMETERVALUE_INVALIDBANDWIDTHINANDOUT = "InvalidParameterValue.InvalidBandwidthInAndOut"
//  INVALIDPARAMETERVALUE_INVALIDBILLINGTYPE = "InvalidParameterValue.InvalidBillingType"
//  INVALIDPARAMETERVALUE_INVALIDDATADISKNUM = "InvalidParameterValue.InvalidDataDiskNum"
//  INVALIDPARAMETERVALUE_INVALIDDATADISKSIZE = "InvalidParameterValue.InvalidDataDiskSize"
//  INVALIDPARAMETERVALUE_INVALIDDATADISKTYPE = "InvalidParameterValue.InvalidDataDiskType"
//  INVALIDPARAMETERVALUE_INVALIDEIPDIRECTSERVICE = "InvalidParameterValue.InvalidEIPDirectService"
//  INVALIDPARAMETERVALUE_INVALIDISPINNODE = "InvalidParameterValue.InvalidISPInNode"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEARCHITECTURE = "InvalidParameterValue.InvalidImageArchitecture"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageID"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCECHARGETYPE = "InvalidParameterValue.InvalidInstanceChargeType"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCETYPECONFIGID = "InvalidParameterValue.InvalidInstanceTypeConfigID"
//  INVALIDPARAMETERVALUE_INVALIDINTERNETCHARGETYPE = "InvalidParameterValue.InvalidInternetChargeType"
//  INVALIDPARAMETERVALUE_INVALIDMODULEID = "InvalidParameterValue.InvalidModuleID"
//  INVALIDPARAMETERVALUE_INVALIDMODULEIDANDINSTANCETYPECONFIGID = "InvalidParameterValue.InvalidModuleIDAndInstanceTypeConfigID"
//  INVALIDPARAMETERVALUE_INVALIDMODULEIDANDINSTANCETYPEID = "InvalidParameterValue.InvalidModuleIDAndInstanceTypeID"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_INVALIDSECURITYGROUPID = "InvalidParameterValue.InvalidSecurityGroupID"
//  INVALIDPARAMETERVALUE_INVALIDSYSTEMDISKSIZE = "InvalidParameterValue.InvalidSystemDiskSize"
//  INVALIDPARAMETERVALUE_INVALIDSYSTEMDISKTYPE = "InvalidParameterValue.InvalidSystemDiskType"
//  INVALIDPARAMETERVALUE_INVALIDZONE = "InvalidParameterValue.InvalidZone"
//  INVALIDPARAMETERVALUE_INVALIDZONEINSTANCECOUNT = "InvalidParameterValue.InvalidZoneInstanceCount"
//  INVALIDPARAMETERVALUE_INVALIDZONEINSTANCETYPE = "InvalidParameterValue.InvalidZoneInstanceType"
//  INVALIDPARAMETERVALUE_NODENOTSUPPORTIPV6 = "InvalidParameterValue.NodeNotSupportIPv6"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  INVALIDPARAMETERVALUE_SUBNETNOTEXIST = "InvalidParameterValue.SubnetNotExist"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_UNMATCHEDBILLINGTYPE = "InvalidParameterValue.UnmatchedBillingType"
//  INVALIDPARAMETERVALUE_USERNOTSUPPORTIPV6 = "InvalidParameterValue.UserNotSupportIPv6"
//  LIMITEXCEEDED_ADDRESSQUOTALIMITEXCEEDED = "LimitExceeded.AddressQuotaLimitExceeded"
//  LIMITEXCEEDED_ENIQUOTALIMITEXCEEDED = "LimitExceeded.EniQuotaLimitExceeded"
//  LIMITEXCEEDED_INSTANCESECURITYGROUPLIMITEXCEEDED = "LimitExceeded.InstanceSecurityGroupLimitExceeded"
//  LIMITEXCEEDED_NICORIPLIMITEXCEEDED = "LimitExceeded.NicOrIPLimitExceeded"
//  LIMITEXCEEDED_SECURITYGROUPINSTANCELIMITEXCEEDED = "LimitExceeded.SecurityGroupInstanceLimitExceeded"
//  LIMITEXCEEDED_VCPULIMITEXCEEDED = "LimitExceeded.VcpuLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT_IPQUOTANOTENOUGH = "ResourceInsufficient.IPQuotaNotEnough"
//  RESOURCEINSUFFICIENT_INSTANCEQUOTANOTENOUGH = "ResourceInsufficient.InstanceQuotaNotEnough"
//  RESOURCEINSUFFICIENT_PRIVATEIPQUOTANOTENOUGH = "ResourceInsufficient.PrivateIPQuotaNotEnough"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCESSOLDOUT_SPECIFIEDINSTANCETYPE = "ResourcesSoldOut.SpecifiedInstanceType"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
//  UNSUPPORTEDOPERATION_INSTANCETYPENOTSUPPORTIMAGE = "UnsupportedOperation.InstanceTypeNotSupportImage"
func (c *Client) RunInstancesWithContext(ctx context.Context, request *RunInstancesRequest) (response *RunInstancesResponse, err error) {
    if request == nil {
        request = NewRunInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewSetLoadBalancerSecurityGroupsRequest() (request *SetLoadBalancerSecurityGroupsRequest) {
    request = &SetLoadBalancerSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "SetLoadBalancerSecurityGroups")
    
    
    return
}

func NewSetLoadBalancerSecurityGroupsResponse() (response *SetLoadBalancerSecurityGroupsResponse) {
    response = &SetLoadBalancerSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetLoadBalancerSecurityGroups
// This API is used to configure security groups for a CLB instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) SetLoadBalancerSecurityGroups(request *SetLoadBalancerSecurityGroupsRequest) (response *SetLoadBalancerSecurityGroupsResponse, err error) {
    return c.SetLoadBalancerSecurityGroupsWithContext(context.Background(), request)
}

// SetLoadBalancerSecurityGroups
// This API is used to configure security groups for a CLB instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
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
    
    request.Init().WithApiInfo("ecm", APIVersion, "SetSecurityGroupForLoadbalancers")
    
    
    return
}

func NewSetSecurityGroupForLoadbalancersResponse() (response *SetSecurityGroupForLoadbalancersResponse) {
    response = &SetSecurityGroupForLoadbalancersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetSecurityGroupForLoadbalancers
// This API is used to bind or unbind a security group to or from multiple CLB instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) SetSecurityGroupForLoadbalancers(request *SetSecurityGroupForLoadbalancersRequest) (response *SetSecurityGroupForLoadbalancersResponse, err error) {
    return c.SetSecurityGroupForLoadbalancersWithContext(context.Background(), request)
}

// SetSecurityGroupForLoadbalancers
// This API is used to bind or unbind a security group to or from multiple CLB instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICPARAM = "InvalidParameterValue.InvalidPublicParam"
//  INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
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

func NewStartInstancesRequest() (request *StartInstancesRequest) {
    request = &StartInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "StartInstances")
    
    
    return
}

func NewStartInstancesResponse() (response *StartInstancesResponse) {
    response = &StartInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartInstances
// This API is used to start an instance. Only instances in `STOPPED` status can be started. The instance status will become `STARTING` when this API is called successfully and will become `RUNNING` when the instance is started successfully.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCENOTALLSTOPPED = "FailedOperation.InstanceNotAllStopped"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) StartInstances(request *StartInstancesRequest) (response *StartInstancesResponse, err error) {
    return c.StartInstancesWithContext(context.Background(), request)
}

// StartInstances
// This API is used to start an instance. Only instances in `STOPPED` status can be started. The instance status will become `STARTING` when this API is called successfully and will become `RUNNING` when the instance is started successfully.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCENOTALLSTOPPED = "FailedOperation.InstanceNotAllStopped"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) StartInstancesWithContext(ctx context.Context, request *StartInstancesRequest) (response *StartInstancesResponse, err error) {
    if request == nil {
        request = NewStartInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewStopInstancesRequest() (request *StopInstancesRequest) {
    request = &StopInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "StopInstances")
    
    
    return
}

func NewStopInstancesResponse() (response *StopInstancesResponse) {
    response = &StopInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopInstances
// Only instances in `RUNNING` status can be shut down.
//
// The instance status will become `STOPPING` when the API is called successfully and will become `STOPPED` when the instance is shut down successfully.
//
// Forced shutdown is supported. Just like powering off a physical PC, a forced shutdown may cause data loss or file system corruption. Make sure that you use this API only when the server cannot be shut down normally.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCEOWNERCHECKFAILED = "FailedOperation.InstanceOwnerCheckFailed"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVAILDMODIFYPARAM = "InvalidParameterValue.InvaildModifyParam"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  RESOURCEUNAVAILABLE_INSTANCENOTRUNNING = "ResourceUnavailable.InstanceNotRunning"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) StopInstances(request *StopInstancesRequest) (response *StopInstancesResponse, err error) {
    return c.StopInstancesWithContext(context.Background(), request)
}

// StopInstances
// Only instances in `RUNNING` status can be shut down.
//
// The instance status will become `STOPPING` when the API is called successfully and will become `STOPPED` when the instance is shut down successfully.
//
// Forced shutdown is supported. Just like powering off a physical PC, a forced shutdown may cause data loss or file system corruption. Make sure that you use this API only when the server cannot be shut down normally.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCEOWNERCHECKFAILED = "FailedOperation.InstanceOwnerCheckFailed"
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVAILDMODIFYPARAM = "InvalidParameterValue.InvaildModifyParam"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  RESOURCEUNAVAILABLE_INSTANCENOTRUNNING = "ResourceUnavailable.InstanceNotRunning"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) StopInstancesWithContext(ctx context.Context, request *StopInstancesRequest) (response *StopInstancesResponse, err error) {
    if request == nil {
        request = NewStopInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateInstancesRequest() (request *TerminateInstancesRequest) {
    request = &TerminateInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecm", APIVersion, "TerminateInstances")
    
    
    return
}

func NewTerminateInstancesResponse() (response *TerminateInstancesResponse) {
    response = &TerminateInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TerminateInstances
// This API is used to terminate an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_INVALIDTIME = "InvalidParameterValue.InvalidTime"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  INVALIDPARAMETERVALUE_TERMINATETIMESMALLER = "InvalidParameterValue.TerminateTimeSmaller"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) TerminateInstances(request *TerminateInstancesRequest) (response *TerminateInstancesResponse, err error) {
    return c.TerminateInstancesWithContext(context.Background(), request)
}

// TerminateInstances
// This API is used to terminate an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_INTERNALOPERATIONFAILURE = "FailedOperation.InternalOperationFailure"
//  FAILEDOPERATION_OPERATIONNOTALLOW = "FailedOperation.OperationNotAllow"
//  INVALIDPARAMETER_INVALIDDATAFORMAT = "InvalidParameter.InvalidDataFormat"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceID"
//  INVALIDPARAMETERVALUE_INVALIDTIME = "InvalidParameterValue.InvalidTime"
//  INVALIDPARAMETERVALUE_PARAMETERVALUETOOLARGE = "InvalidParameterValue.ParameterValueTooLarge"
//  INVALIDPARAMETERVALUE_TERMINATETIMESMALLER = "InvalidParameterValue.TerminateTimeSmaller"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  UNAUTHORIZEDOPERATION_FORBIDDENOPERATION = "UnauthorizedOperation.ForbiddenOperation"
func (c *Client) TerminateInstancesWithContext(ctx context.Context, request *TerminateInstancesRequest) (response *TerminateInstancesResponse, err error) {
    if request == nil {
        request = NewTerminateInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateInstancesResponse()
    err = c.Send(request, response)
    return
}
