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

package v20201028

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2020-10-28"

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


func NewCreatePrivateDNSAccountRequest() (request *CreatePrivateDNSAccountRequest) {
    request = &CreatePrivateDNSAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("privatedns", APIVersion, "CreatePrivateDNSAccount")
    
    
    return
}

func NewCreatePrivateDNSAccountResponse() (response *CreatePrivateDNSAccountResponse) {
    response = &CreatePrivateDNSAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePrivateDNSAccount
// This API is used to create a Private DNS account.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACCOUNTEXIST = "InvalidParameter.AccountExist"
//  INVALIDPARAMETER_RECORDEXIST = "InvalidParameter.RecordExist"
//  INVALIDPARAMETER_RECORDNOTEXIST = "InvalidParameter.RecordNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreatePrivateDNSAccount(request *CreatePrivateDNSAccountRequest) (response *CreatePrivateDNSAccountResponse, err error) {
    return c.CreatePrivateDNSAccountWithContext(context.Background(), request)
}

// CreatePrivateDNSAccount
// This API is used to create a Private DNS account.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACCOUNTEXIST = "InvalidParameter.AccountExist"
//  INVALIDPARAMETER_RECORDEXIST = "InvalidParameter.RecordExist"
//  INVALIDPARAMETER_RECORDNOTEXIST = "InvalidParameter.RecordNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreatePrivateDNSAccountWithContext(ctx context.Context, request *CreatePrivateDNSAccountRequest) (response *CreatePrivateDNSAccountResponse, err error) {
    if request == nil {
        request = NewCreatePrivateDNSAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrivateDNSAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrivateDNSAccountResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrivateZoneRequest() (request *CreatePrivateZoneRequest) {
    request = &CreatePrivateZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("privatedns", APIVersion, "CreatePrivateZone")
    
    
    return
}

func NewCreatePrivateZoneResponse() (response *CreatePrivateZoneResponse) {
    response = &CreatePrivateZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePrivateZone
// This API is used to create a private domain.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATEZONEFAILED = "FailedOperation.CreateZoneFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALCIDR = "InvalidParameter.IllegalCidr"
//  INVALIDPARAMETER_ILLEGALDOMAIN = "InvalidParameter.IllegalDomain"
//  INVALIDPARAMETER_ILLEGALDOMAINTLD = "InvalidParameter.IllegalDomainTld"
//  INVALIDPARAMETER_ILLEGALRECORD = "InvalidParameter.IllegalRecord"
//  INVALIDPARAMETER_ILLEGALRECORDVALUE = "InvalidParameter.IllegalRecordValue"
//  INVALIDPARAMETER_RECORDLEVELEXCEED = "InvalidParameter.RecordLevelExceed"
//  INVALIDPARAMETER_VPCBINDED = "InvalidParameter.VpcBinded"
//  INVALIDPARAMETER_VPCBINDEDMAINDOMAIN = "InvalidParameter.VpcBindedMainDomain"
//  INVALIDPARAMETER_VPCPTRZONEBINDEXCEED = "InvalidParameter.VpcPtrZoneBindExceed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_RESERVEDDOMAIN = "InvalidParameterValue.ReservedDomain"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TLDOUTOFLIMIT = "LimitExceeded.TldOutOfLimit"
//  LIMITEXCEEDED_TLDOUTOFRANGE = "LimitExceeded.TldOutOfRange"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ROLEUNAUTHORIZED = "UnauthorizedOperation.RoleUnAuthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ACCOUNTNOTBOUND = "UnsupportedOperation.AccountNotBound"
//  UNSUPPORTEDOPERATION_NOTSUPPORTDNSFORWARD = "UnsupportedOperation.NotSupportDnsForward"
func (c *Client) CreatePrivateZone(request *CreatePrivateZoneRequest) (response *CreatePrivateZoneResponse, err error) {
    return c.CreatePrivateZoneWithContext(context.Background(), request)
}

// CreatePrivateZone
// This API is used to create a private domain.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATEZONEFAILED = "FailedOperation.CreateZoneFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALCIDR = "InvalidParameter.IllegalCidr"
//  INVALIDPARAMETER_ILLEGALDOMAIN = "InvalidParameter.IllegalDomain"
//  INVALIDPARAMETER_ILLEGALDOMAINTLD = "InvalidParameter.IllegalDomainTld"
//  INVALIDPARAMETER_ILLEGALRECORD = "InvalidParameter.IllegalRecord"
//  INVALIDPARAMETER_ILLEGALRECORDVALUE = "InvalidParameter.IllegalRecordValue"
//  INVALIDPARAMETER_RECORDLEVELEXCEED = "InvalidParameter.RecordLevelExceed"
//  INVALIDPARAMETER_VPCBINDED = "InvalidParameter.VpcBinded"
//  INVALIDPARAMETER_VPCBINDEDMAINDOMAIN = "InvalidParameter.VpcBindedMainDomain"
//  INVALIDPARAMETER_VPCPTRZONEBINDEXCEED = "InvalidParameter.VpcPtrZoneBindExceed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_RESERVEDDOMAIN = "InvalidParameterValue.ReservedDomain"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TLDOUTOFLIMIT = "LimitExceeded.TldOutOfLimit"
//  LIMITEXCEEDED_TLDOUTOFRANGE = "LimitExceeded.TldOutOfRange"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ROLEUNAUTHORIZED = "UnauthorizedOperation.RoleUnAuthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ACCOUNTNOTBOUND = "UnsupportedOperation.AccountNotBound"
//  UNSUPPORTEDOPERATION_NOTSUPPORTDNSFORWARD = "UnsupportedOperation.NotSupportDnsForward"
func (c *Client) CreatePrivateZoneWithContext(ctx context.Context, request *CreatePrivateZoneRequest) (response *CreatePrivateZoneResponse, err error) {
    if request == nil {
        request = NewCreatePrivateZoneRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrivateZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrivateZoneResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrivateZoneRecordRequest() (request *CreatePrivateZoneRecordRequest) {
    request = &CreatePrivateZoneRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("privatedns", APIVersion, "CreatePrivateZoneRecord")
    
    
    return
}

func NewCreatePrivateZoneRecordResponse() (response *CreatePrivateZoneRecordResponse) {
    response = &CreatePrivateZoneRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePrivateZoneRecord
// This API is used to add a DNS record for a private domain.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATERECORDFAILED = "FailedOperation.CreateRecordFailed"
//  FAILEDOPERATION_DATAERROR = "FailedOperation.DataError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALPTRRECORD = "InvalidParameter.IllegalPTRRecord"
//  INVALIDPARAMETER_ILLEGALRECORD = "InvalidParameter.IllegalRecord"
//  INVALIDPARAMETER_ILLEGALRECORDVALUE = "InvalidParameter.IllegalRecordValue"
//  INVALIDPARAMETER_INVALIDMX = "InvalidParameter.InvalidMX"
//  INVALIDPARAMETER_RECORDAAAACOUNTEXCEED = "InvalidParameter.RecordAAAACountExceed"
//  INVALIDPARAMETER_RECORDACOUNTEXCEED = "InvalidParameter.RecordACountExceed"
//  INVALIDPARAMETER_RECORDCNAMECOUNTEXCEED = "InvalidParameter.RecordCNAMECountExceed"
//  INVALIDPARAMETER_RECORDCONFLICT = "InvalidParameter.RecordConflict"
//  INVALIDPARAMETER_RECORDCOUNTEXCEED = "InvalidParameter.RecordCountExceed"
//  INVALIDPARAMETER_RECORDEXIST = "InvalidParameter.RecordExist"
//  INVALIDPARAMETER_RECORDMXCOUNTEXCEED = "InvalidParameter.RecordMXCountExceed"
//  INVALIDPARAMETER_RECORDROLLLIMITCOUNTEXCEED = "InvalidParameter.RecordRolllimitCountExceed"
//  INVALIDPARAMETER_RECORDTXTCOUNTEXCEED = "InvalidParameter.RecordTXTCountExceed"
//  INVALIDPARAMETER_RECORDUNSUPPORTWEIGHT = "InvalidParameter.RecordUnsupportWeight"
//  INVALIDPARAMETER_VPCBINDEDMAINDOMAIN = "InvalidParameter.VpcBindedMainDomain"
//  INVALIDPARAMETER_ZONENOTEXISTS = "InvalidParameter.ZoneNotExists"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TLDOUTOFLIMIT = "LimitExceeded.TldOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_TLDPACKAGEEXPIRED = "ResourceUnavailable.TldPackageExpired"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreatePrivateZoneRecord(request *CreatePrivateZoneRecordRequest) (response *CreatePrivateZoneRecordResponse, err error) {
    return c.CreatePrivateZoneRecordWithContext(context.Background(), request)
}

// CreatePrivateZoneRecord
// This API is used to add a DNS record for a private domain.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATERECORDFAILED = "FailedOperation.CreateRecordFailed"
//  FAILEDOPERATION_DATAERROR = "FailedOperation.DataError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALPTRRECORD = "InvalidParameter.IllegalPTRRecord"
//  INVALIDPARAMETER_ILLEGALRECORD = "InvalidParameter.IllegalRecord"
//  INVALIDPARAMETER_ILLEGALRECORDVALUE = "InvalidParameter.IllegalRecordValue"
//  INVALIDPARAMETER_INVALIDMX = "InvalidParameter.InvalidMX"
//  INVALIDPARAMETER_RECORDAAAACOUNTEXCEED = "InvalidParameter.RecordAAAACountExceed"
//  INVALIDPARAMETER_RECORDACOUNTEXCEED = "InvalidParameter.RecordACountExceed"
//  INVALIDPARAMETER_RECORDCNAMECOUNTEXCEED = "InvalidParameter.RecordCNAMECountExceed"
//  INVALIDPARAMETER_RECORDCONFLICT = "InvalidParameter.RecordConflict"
//  INVALIDPARAMETER_RECORDCOUNTEXCEED = "InvalidParameter.RecordCountExceed"
//  INVALIDPARAMETER_RECORDEXIST = "InvalidParameter.RecordExist"
//  INVALIDPARAMETER_RECORDMXCOUNTEXCEED = "InvalidParameter.RecordMXCountExceed"
//  INVALIDPARAMETER_RECORDROLLLIMITCOUNTEXCEED = "InvalidParameter.RecordRolllimitCountExceed"
//  INVALIDPARAMETER_RECORDTXTCOUNTEXCEED = "InvalidParameter.RecordTXTCountExceed"
//  INVALIDPARAMETER_RECORDUNSUPPORTWEIGHT = "InvalidParameter.RecordUnsupportWeight"
//  INVALIDPARAMETER_VPCBINDEDMAINDOMAIN = "InvalidParameter.VpcBindedMainDomain"
//  INVALIDPARAMETER_ZONENOTEXISTS = "InvalidParameter.ZoneNotExists"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TLDOUTOFLIMIT = "LimitExceeded.TldOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_TLDPACKAGEEXPIRED = "ResourceUnavailable.TldPackageExpired"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreatePrivateZoneRecordWithContext(ctx context.Context, request *CreatePrivateZoneRecordRequest) (response *CreatePrivateZoneRecordResponse, err error) {
    if request == nil {
        request = NewCreatePrivateZoneRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrivateZoneRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrivateZoneRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountVpcListRequest() (request *DescribeAccountVpcListRequest) {
    request = &DescribeAccountVpcListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("privatedns", APIVersion, "DescribeAccountVpcList")
    
    
    return
}

func NewDescribeAccountVpcListResponse() (response *DescribeAccountVpcListResponse) {
    response = &DescribeAccountVpcListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccountVpcList
// This API is used to get the VPC list of a Private DNS account.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ROLEUNAUTHORIZED = "UnauthorizedOperation.RoleUnAuthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ACCOUNTNOTBOUND = "UnsupportedOperation.AccountNotBound"
func (c *Client) DescribeAccountVpcList(request *DescribeAccountVpcListRequest) (response *DescribeAccountVpcListResponse, err error) {
    return c.DescribeAccountVpcListWithContext(context.Background(), request)
}

// DescribeAccountVpcList
// This API is used to get the VPC list of a Private DNS account.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ROLEUNAUTHORIZED = "UnauthorizedOperation.RoleUnAuthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ACCOUNTNOTBOUND = "UnsupportedOperation.AccountNotBound"
func (c *Client) DescribeAccountVpcListWithContext(ctx context.Context, request *DescribeAccountVpcListRequest) (response *DescribeAccountVpcListResponse, err error) {
    if request == nil {
        request = NewDescribeAccountVpcListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccountVpcList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccountVpcListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuditLogRequest() (request *DescribeAuditLogRequest) {
    request = &DescribeAuditLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("privatedns", APIVersion, "DescribeAuditLog")
    
    
    return
}

func NewDescribeAuditLogResponse() (response *DescribeAuditLogResponse) {
    response = &DescribeAuditLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAuditLog
// This API is used to get the list of operation logs.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAuditLog(request *DescribeAuditLogRequest) (response *DescribeAuditLogResponse, err error) {
    return c.DescribeAuditLogWithContext(context.Background(), request)
}

// DescribeAuditLog
// This API is used to get the list of operation logs.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAuditLogWithContext(ctx context.Context, request *DescribeAuditLogRequest) (response *DescribeAuditLogResponse, err error) {
    if request == nil {
        request = NewDescribeAuditLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuditLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuditLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDashboardRequest() (request *DescribeDashboardRequest) {
    request = &DescribeDashboardRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("privatedns", APIVersion, "DescribeDashboard")
    
    
    return
}

func NewDescribeDashboardResponse() (response *DescribeDashboardResponse) {
    response = &DescribeDashboardResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDashboard
// This API is used to get the overview of private DNS records.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDashboard(request *DescribeDashboardRequest) (response *DescribeDashboardResponse, err error) {
    return c.DescribeDashboardWithContext(context.Background(), request)
}

// DescribeDashboard
// This API is used to get the overview of private DNS records.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDashboardWithContext(ctx context.Context, request *DescribeDashboardRequest) (response *DescribeDashboardResponse, err error) {
    if request == nil {
        request = NewDescribeDashboardRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDashboard require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDashboardResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrivateDNSAccountListRequest() (request *DescribePrivateDNSAccountListRequest) {
    request = &DescribePrivateDNSAccountListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("privatedns", APIVersion, "DescribePrivateDNSAccountList")
    
    
    return
}

func NewDescribePrivateDNSAccountListResponse() (response *DescribePrivateDNSAccountListResponse) {
    response = &DescribePrivateDNSAccountListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePrivateDNSAccountList
// This API is used to get the list of Private DNS accounts.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePrivateDNSAccountList(request *DescribePrivateDNSAccountListRequest) (response *DescribePrivateDNSAccountListResponse, err error) {
    return c.DescribePrivateDNSAccountListWithContext(context.Background(), request)
}

// DescribePrivateDNSAccountList
// This API is used to get the list of Private DNS accounts.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePrivateDNSAccountListWithContext(ctx context.Context, request *DescribePrivateDNSAccountListRequest) (response *DescribePrivateDNSAccountListResponse, err error) {
    if request == nil {
        request = NewDescribePrivateDNSAccountListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrivateDNSAccountList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrivateDNSAccountListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrivateZoneServiceRequest() (request *DescribePrivateZoneServiceRequest) {
    request = &DescribePrivateZoneServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("privatedns", APIVersion, "DescribePrivateZoneService")
    
    
    return
}

func NewDescribePrivateZoneServiceResponse() (response *DescribePrivateZoneServiceResponse) {
    response = &DescribePrivateZoneServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePrivateZoneService
// This API is used to query the Private DNS activation status.
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePrivateZoneService(request *DescribePrivateZoneServiceRequest) (response *DescribePrivateZoneServiceResponse, err error) {
    return c.DescribePrivateZoneServiceWithContext(context.Background(), request)
}

// DescribePrivateZoneService
// This API is used to query the Private DNS activation status.
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePrivateZoneServiceWithContext(ctx context.Context, request *DescribePrivateZoneServiceRequest) (response *DescribePrivateZoneServiceResponse, err error) {
    if request == nil {
        request = NewDescribePrivateZoneServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrivateZoneService require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrivateZoneServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQuotaUsageRequest() (request *DescribeQuotaUsageRequest) {
    request = &DescribeQuotaUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("privatedns", APIVersion, "DescribeQuotaUsage")
    
    
    return
}

func NewDescribeQuotaUsageResponse() (response *DescribeQuotaUsageResponse) {
    response = &DescribeQuotaUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeQuotaUsage
// This API is used to query quota usage.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ROLEUNAUTHORIZED = "UnauthorizedOperation.RoleUnAuthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeQuotaUsage(request *DescribeQuotaUsageRequest) (response *DescribeQuotaUsageResponse, err error) {
    return c.DescribeQuotaUsageWithContext(context.Background(), request)
}

// DescribeQuotaUsage
// This API is used to query quota usage.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ROLEUNAUTHORIZED = "UnauthorizedOperation.RoleUnAuthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeQuotaUsageWithContext(ctx context.Context, request *DescribeQuotaUsageRequest) (response *DescribeQuotaUsageResponse, err error) {
    if request == nil {
        request = NewDescribeQuotaUsageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeQuotaUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeQuotaUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRequestDataRequest() (request *DescribeRequestDataRequest) {
    request = &DescribeRequestDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("privatedns", APIVersion, "DescribeRequestData")
    
    
    return
}

func NewDescribeRequestDataResponse() (response *DescribeRequestDataResponse) {
    response = &DescribeRequestDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRequestData
// This API is used to get the DNS request volume of a private domain.
//
// error code that may be returned:
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
func (c *Client) DescribeRequestData(request *DescribeRequestDataRequest) (response *DescribeRequestDataResponse, err error) {
    return c.DescribeRequestDataWithContext(context.Background(), request)
}

// DescribeRequestData
// This API is used to get the DNS request volume of a private domain.
//
// error code that may be returned:
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
func (c *Client) DescribeRequestDataWithContext(ctx context.Context, request *DescribeRequestDataRequest) (response *DescribeRequestDataResponse, err error) {
    if request == nil {
        request = NewDescribeRequestDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRequestData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRequestDataResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrivateZoneRequest() (request *ModifyPrivateZoneRequest) {
    request = &ModifyPrivateZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("privatedns", APIVersion, "ModifyPrivateZone")
    
    
    return
}

func NewModifyPrivateZoneResponse() (response *ModifyPrivateZoneResponse) {
    response = &ModifyPrivateZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPrivateZone
// This API is used to modify a private domain.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MODIFYZONEFAILED = "FailedOperation.ModifyZoneFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ZONENOTEXISTS = "InvalidParameter.ZoneNotExists"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TLDOUTOFLIMIT = "LimitExceeded.TldOutOfLimit"
//  LIMITEXCEEDED_TLDOUTOFRANGE = "LimitExceeded.TldOutOfRange"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_TLDPACKAGEEXPIRED = "ResourceUnavailable.TldPackageExpired"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_NOTSUPPORTDNSFORWARD = "UnsupportedOperation.NotSupportDnsForward"
func (c *Client) ModifyPrivateZone(request *ModifyPrivateZoneRequest) (response *ModifyPrivateZoneResponse, err error) {
    return c.ModifyPrivateZoneWithContext(context.Background(), request)
}

// ModifyPrivateZone
// This API is used to modify a private domain.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MODIFYZONEFAILED = "FailedOperation.ModifyZoneFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ZONENOTEXISTS = "InvalidParameter.ZoneNotExists"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TLDOUTOFLIMIT = "LimitExceeded.TldOutOfLimit"
//  LIMITEXCEEDED_TLDOUTOFRANGE = "LimitExceeded.TldOutOfRange"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_TLDPACKAGEEXPIRED = "ResourceUnavailable.TldPackageExpired"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_NOTSUPPORTDNSFORWARD = "UnsupportedOperation.NotSupportDnsForward"
func (c *Client) ModifyPrivateZoneWithContext(ctx context.Context, request *ModifyPrivateZoneRequest) (response *ModifyPrivateZoneResponse, err error) {
    if request == nil {
        request = NewModifyPrivateZoneRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPrivateZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPrivateZoneResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrivateZoneRecordRequest() (request *ModifyPrivateZoneRecordRequest) {
    request = &ModifyPrivateZoneRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("privatedns", APIVersion, "ModifyPrivateZoneRecord")
    
    
    return
}

func NewModifyPrivateZoneRecordResponse() (response *ModifyPrivateZoneRecordResponse) {
    response = &ModifyPrivateZoneRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPrivateZoneRecord
// This API is used to modify a DNS record for a private domain.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAERROR = "FailedOperation.DataError"
//  FAILEDOPERATION_MODIFYRECORDFAILED = "FailedOperation.ModifyRecordFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALCIDR = "InvalidParameter.IllegalCidr"
//  INVALIDPARAMETER_ILLEGALPTRRECORD = "InvalidParameter.IllegalPTRRecord"
//  INVALIDPARAMETER_ILLEGALRECORD = "InvalidParameter.IllegalRecord"
//  INVALIDPARAMETER_ILLEGALRECORDVALUE = "InvalidParameter.IllegalRecordValue"
//  INVALIDPARAMETER_INVALIDMX = "InvalidParameter.InvalidMX"
//  INVALIDPARAMETER_RECORDAAAACOUNTEXCEED = "InvalidParameter.RecordAAAACountExceed"
//  INVALIDPARAMETER_RECORDACOUNTEXCEED = "InvalidParameter.RecordACountExceed"
//  INVALIDPARAMETER_RECORDCNAMECOUNTEXCEED = "InvalidParameter.RecordCNAMECountExceed"
//  INVALIDPARAMETER_RECORDCONFLICT = "InvalidParameter.RecordConflict"
//  INVALIDPARAMETER_RECORDCOUNTEXCEED = "InvalidParameter.RecordCountExceed"
//  INVALIDPARAMETER_RECORDEXIST = "InvalidParameter.RecordExist"
//  INVALIDPARAMETER_RECORDLEVELEXCEED = "InvalidParameter.RecordLevelExceed"
//  INVALIDPARAMETER_RECORDMXCOUNTEXCEED = "InvalidParameter.RecordMXCountExceed"
//  INVALIDPARAMETER_RECORDNOTEXIST = "InvalidParameter.RecordNotExist"
//  INVALIDPARAMETER_RECORDROLLLIMITCOUNTEXCEED = "InvalidParameter.RecordRolllimitCountExceed"
//  INVALIDPARAMETER_RECORDTXTCOUNTEXCEED = "InvalidParameter.RecordTXTCountExceed"
//  INVALIDPARAMETER_RECORDUNSUPPORTWEIGHT = "InvalidParameter.RecordUnsupportWeight"
//  INVALIDPARAMETER_VPCBINDEDMAINDOMAIN = "InvalidParameter.VpcBindedMainDomain"
//  INVALIDPARAMETER_ZONENOTEXISTS = "InvalidParameter.ZoneNotExists"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TLDOUTOFLIMIT = "LimitExceeded.TldOutOfLimit"
//  LIMITEXCEEDED_TLDOUTOFRANGE = "LimitExceeded.TldOutOfRange"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_TLDPACKAGEEXPIRED = "ResourceUnavailable.TldPackageExpired"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyPrivateZoneRecord(request *ModifyPrivateZoneRecordRequest) (response *ModifyPrivateZoneRecordResponse, err error) {
    return c.ModifyPrivateZoneRecordWithContext(context.Background(), request)
}

// ModifyPrivateZoneRecord
// This API is used to modify a DNS record for a private domain.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATAERROR = "FailedOperation.DataError"
//  FAILEDOPERATION_MODIFYRECORDFAILED = "FailedOperation.ModifyRecordFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALCIDR = "InvalidParameter.IllegalCidr"
//  INVALIDPARAMETER_ILLEGALPTRRECORD = "InvalidParameter.IllegalPTRRecord"
//  INVALIDPARAMETER_ILLEGALRECORD = "InvalidParameter.IllegalRecord"
//  INVALIDPARAMETER_ILLEGALRECORDVALUE = "InvalidParameter.IllegalRecordValue"
//  INVALIDPARAMETER_INVALIDMX = "InvalidParameter.InvalidMX"
//  INVALIDPARAMETER_RECORDAAAACOUNTEXCEED = "InvalidParameter.RecordAAAACountExceed"
//  INVALIDPARAMETER_RECORDACOUNTEXCEED = "InvalidParameter.RecordACountExceed"
//  INVALIDPARAMETER_RECORDCNAMECOUNTEXCEED = "InvalidParameter.RecordCNAMECountExceed"
//  INVALIDPARAMETER_RECORDCONFLICT = "InvalidParameter.RecordConflict"
//  INVALIDPARAMETER_RECORDCOUNTEXCEED = "InvalidParameter.RecordCountExceed"
//  INVALIDPARAMETER_RECORDEXIST = "InvalidParameter.RecordExist"
//  INVALIDPARAMETER_RECORDLEVELEXCEED = "InvalidParameter.RecordLevelExceed"
//  INVALIDPARAMETER_RECORDMXCOUNTEXCEED = "InvalidParameter.RecordMXCountExceed"
//  INVALIDPARAMETER_RECORDNOTEXIST = "InvalidParameter.RecordNotExist"
//  INVALIDPARAMETER_RECORDROLLLIMITCOUNTEXCEED = "InvalidParameter.RecordRolllimitCountExceed"
//  INVALIDPARAMETER_RECORDTXTCOUNTEXCEED = "InvalidParameter.RecordTXTCountExceed"
//  INVALIDPARAMETER_RECORDUNSUPPORTWEIGHT = "InvalidParameter.RecordUnsupportWeight"
//  INVALIDPARAMETER_VPCBINDEDMAINDOMAIN = "InvalidParameter.VpcBindedMainDomain"
//  INVALIDPARAMETER_ZONENOTEXISTS = "InvalidParameter.ZoneNotExists"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TLDOUTOFLIMIT = "LimitExceeded.TldOutOfLimit"
//  LIMITEXCEEDED_TLDOUTOFRANGE = "LimitExceeded.TldOutOfRange"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_TLDPACKAGEEXPIRED = "ResourceUnavailable.TldPackageExpired"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyPrivateZoneRecordWithContext(ctx context.Context, request *ModifyPrivateZoneRecordRequest) (response *ModifyPrivateZoneRecordResponse, err error) {
    if request == nil {
        request = NewModifyPrivateZoneRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPrivateZoneRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPrivateZoneRecordResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrivateZoneVpcRequest() (request *ModifyPrivateZoneVpcRequest) {
    request = &ModifyPrivateZoneVpcRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("privatedns", APIVersion, "ModifyPrivateZoneVpc")
    
    
    return
}

func NewModifyPrivateZoneVpcResponse() (response *ModifyPrivateZoneVpcResponse) {
    response = &ModifyPrivateZoneVpcResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPrivateZoneVpc
// This API is used to modify the VPC associated with a private domain.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BINDZONEVPCFAILED = "FailedOperation.BindZoneVpcFailed"
//  FAILEDOPERATION_DATAERROR = "FailedOperation.DataError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALVPCINFO = "InvalidParameter.IllegalVpcInfo"
//  INVALIDPARAMETER_VPCBINDED = "InvalidParameter.VpcBinded"
//  INVALIDPARAMETER_VPCBINDEDMAINDOMAIN = "InvalidParameter.VpcBindedMainDomain"
//  INVALIDPARAMETER_ZONENOTEXISTS = "InvalidParameter.ZoneNotExists"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TLDOUTOFLIMIT = "LimitExceeded.TldOutOfLimit"
//  LIMITEXCEEDED_TLDOUTOFRANGE = "LimitExceeded.TldOutOfRange"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_TLDPACKAGEEXPIRED = "ResourceUnavailable.TldPackageExpired"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ROLEUNAUTHORIZED = "UnauthorizedOperation.RoleUnAuthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyPrivateZoneVpc(request *ModifyPrivateZoneVpcRequest) (response *ModifyPrivateZoneVpcResponse, err error) {
    return c.ModifyPrivateZoneVpcWithContext(context.Background(), request)
}

// ModifyPrivateZoneVpc
// This API is used to modify the VPC associated with a private domain.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BINDZONEVPCFAILED = "FailedOperation.BindZoneVpcFailed"
//  FAILEDOPERATION_DATAERROR = "FailedOperation.DataError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALVPCINFO = "InvalidParameter.IllegalVpcInfo"
//  INVALIDPARAMETER_VPCBINDED = "InvalidParameter.VpcBinded"
//  INVALIDPARAMETER_VPCBINDEDMAINDOMAIN = "InvalidParameter.VpcBindedMainDomain"
//  INVALIDPARAMETER_ZONENOTEXISTS = "InvalidParameter.ZoneNotExists"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TLDOUTOFLIMIT = "LimitExceeded.TldOutOfLimit"
//  LIMITEXCEEDED_TLDOUTOFRANGE = "LimitExceeded.TldOutOfRange"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_TLDPACKAGEEXPIRED = "ResourceUnavailable.TldPackageExpired"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ROLEUNAUTHORIZED = "UnauthorizedOperation.RoleUnAuthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyPrivateZoneVpcWithContext(ctx context.Context, request *ModifyPrivateZoneVpcRequest) (response *ModifyPrivateZoneVpcResponse, err error) {
    if request == nil {
        request = NewModifyPrivateZoneVpcRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPrivateZoneVpc require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPrivateZoneVpcResponse()
    err = c.Send(request, response)
    return
}

func NewSubscribePrivateZoneServiceRequest() (request *SubscribePrivateZoneServiceRequest) {
    request = &SubscribePrivateZoneServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("privatedns", APIVersion, "SubscribePrivateZoneService")
    
    
    return
}

func NewSubscribePrivateZoneServiceResponse() (response *SubscribePrivateZoneServiceResponse) {
    response = &SubscribePrivateZoneServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SubscribePrivateZoneService
// This API is used to activate the Private DNS service.
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SubscribePrivateZoneService(request *SubscribePrivateZoneServiceRequest) (response *SubscribePrivateZoneServiceResponse, err error) {
    return c.SubscribePrivateZoneServiceWithContext(context.Background(), request)
}

// SubscribePrivateZoneService
// This API is used to activate the Private DNS service.
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SubscribePrivateZoneServiceWithContext(ctx context.Context, request *SubscribePrivateZoneServiceRequest) (response *SubscribePrivateZoneServiceResponse, err error) {
    if request == nil {
        request = NewSubscribePrivateZoneServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubscribePrivateZoneService require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubscribePrivateZoneServiceResponse()
    err = c.Send(request, response)
    return
}
