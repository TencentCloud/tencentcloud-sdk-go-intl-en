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

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewAddDBInstanceToReadOnlyGroupRequest() (request *AddDBInstanceToReadOnlyGroupRequest) {
    request = &AddDBInstanceToReadOnlyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "AddDBInstanceToReadOnlyGroup")
    
    return
}

func NewAddDBInstanceToReadOnlyGroupResponse() (response *AddDBInstanceToReadOnlyGroupResponse) {
    response = &AddDBInstanceToReadOnlyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddDBInstanceToReadOnlyGroup
// This API is used to add a read-only replica to an RO group.
//
// error code that may be returned:
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_ROGROUPMASTERINSTANCENOTRIGHT = "FailedOperation.ROGroupMasterInstanceNotRight"
//  FAILEDOPERATION_ROGROUPNOTFOUNDERROR = "FailedOperation.ROGroupNotFoundError"
//  FAILEDOPERATION_ROINSTANCEHASINROGROUPERROR = "FailedOperation.ROInstanceHasInROGroupError"
//  INSTANCENOTEXIST = "InstanceNotExist"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREADONLYGROUPSTATUS = "InvalidParameterValue.InvalidReadOnlyGroupStatus"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_READONLYGROUPNOTEXIST = "InvalidParameterValue.ReadOnlyGroupNotExist"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  OPERATIONDENIED_ROGROUPSTATUSERROR = "OperationDenied.ROGroupStatusError"
//  OPERATIONDENIED_ROINSTANCECOUNTEXEEDERROR = "OperationDenied.RoInstanceCountExeedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
func (c *Client) AddDBInstanceToReadOnlyGroup(request *AddDBInstanceToReadOnlyGroupRequest) (response *AddDBInstanceToReadOnlyGroupResponse, err error) {
    if request == nil {
        request = NewAddDBInstanceToReadOnlyGroupRequest()
    }
    response = NewAddDBInstanceToReadOnlyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCloseDBExtranetAccessRequest() (request *CloseDBExtranetAccessRequest) {
    request = &CloseDBExtranetAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "CloseDBExtranetAccess")
    
    return
}

func NewCloseDBExtranetAccessResponse() (response *CloseDBExtranetAccessResponse) {
    response = &CloseDBExtranetAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CloseDBExtranetAccess
// This API is used to disable the public network link to an instance.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_FLOWERROR = "InternalError.FlowError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
func (c *Client) CloseDBExtranetAccess(request *CloseDBExtranetAccessRequest) (response *CloseDBExtranetAccessResponse, err error) {
    if request == nil {
        request = NewCloseDBExtranetAccessRequest()
    }
    response = NewCloseDBExtranetAccessResponse()
    err = c.Send(request, response)
    return
}

func NewCloseServerlessDBExtranetAccessRequest() (request *CloseServerlessDBExtranetAccessRequest) {
    request = &CloseServerlessDBExtranetAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "CloseServerlessDBExtranetAccess")
    
    return
}

func NewCloseServerlessDBExtranetAccessResponse() (response *CloseServerlessDBExtranetAccessResponse) {
    response = &CloseServerlessDBExtranetAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CloseServerlessDBExtranetAccess
// This API is used to disable public network access for a PostgreSQL for Serverless instance.
//
// error code that may be returned:
//  ACCOUNTNOTEXIST = "AccountNotExist"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DBERROR = "DBError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_INVALIDTRADEOPERATE = "FailedOperation.InvalidTradeOperate"
//  FAILEDOPERATION_LIMITOPERATION = "FailedOperation.LimitOperation"
//  FAILEDOPERATION_QUERYVPCFALIED = "FailedOperation.QueryVpcFalied"
//  FLOWERROR = "FlowError"
//  INSTANCENOTEXIST = "InstanceNotExist"
//  INTERFACENAMENOTFOUND = "InterfaceNameNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_FLOWERROR = "InternalError.FlowError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDACCOUNTSTATUS = "InvalidAccountStatus"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INTERFACENAMENOTFOUND = "InvalidParameterValue.InterfaceNameNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_STRUCTPARSEFAILED = "InvalidParameterValue.StructParseFailed"
//  LIMITOPERATION = "LimitOperation"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
//  RESOURCEUNAVAILABLE_RESOURCENOPERMISSION = "ResourceUnavailable.ResourceNoPermission"
//  RESOURCEUNAVAILABLE_VPCRESOURCENOTFOUND = "ResourceUnavailable.VpcResourceNotFound"
//  STRUCTPARSEFAILED = "StructParseFailed"
//  SYSTEMERROR = "SystemError"
//  UNKNOWNERROR = "UnknownError"
//  VPCERROR = "VpcError"
func (c *Client) CloseServerlessDBExtranetAccess(request *CloseServerlessDBExtranetAccessRequest) (response *CloseServerlessDBExtranetAccessResponse, err error) {
    if request == nil {
        request = NewCloseServerlessDBExtranetAccessRequest()
    }
    response = NewCloseServerlessDBExtranetAccessResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDBInstancesRequest() (request *CreateDBInstancesRequest) {
    request = &CreateDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "CreateDBInstances")
    
    return
}

func NewCreateDBInstancesResponse() (response *CreateDBInstancesResponse) {
    response = &CreateDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDBInstances
// This API is used to create (but not initialize) one or more TencentDB for PostgreSQL instances.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_CDBCGWCONNECTERROR = "FailedOperation.CdbCgwConnectError"
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_GETSUBNETERROR = "FailedOperation.GetSubnetError"
//  FAILEDOPERATION_GETVPCINFOERROR = "FailedOperation.GetVpcInfoError"
//  FAILEDOPERATION_INVALIDTRADEOPERATE = "FailedOperation.InvalidTradeOperate"
//  FAILEDOPERATION_PAYORDERFAILED = "FailedOperation.PayOrderFailed"
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  FAILEDOPERATION_QUERYSPECERROR = "FailedOperation.QuerySpecError"
//  FAILEDOPERATION_QUERYVPCFAILED = "FailedOperation.QueryVpcFailed"
//  FAILEDOPERATION_QUERYVPCFALIED = "FailedOperation.QueryVpcFalied"
//  FAILEDOPERATION_SERVICEACCESSERROR = "FailedOperation.ServiceAccessError"
//  FAILEDOPERATION_STORAGEMEMORYCHECKERROR = "FailedOperation.StorageMemoryCheckError"
//  FAILEDOPERATION_TRADECREATEERROR = "FailedOperation.TradeCreateError"
//  FAILEDOPERATION_TRADEPAYORDERSERROR = "FailedOperation.TradePayOrdersError"
//  INTERNALERROR_CGWERROR = "InternalError.CgwError"
//  INTERNALERROR_CNSERROR = "InternalError.CnsError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_FLOWERROR = "InternalError.FlowError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETER_TRADEACCESSDENIEDERROR = "InvalidParameter.TradeAccessDeniedError"
//  INVALIDPARAMETER_VPCNOTFOUNDERROR = "InvalidParameter.VpcNotFoundError"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_BADSPEC = "InvalidParameterValue.BadSpec"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCECHARGETYPE = "InvalidParameterValue.IllegalInstanceChargeType"
//  INVALIDPARAMETERVALUE_ILLEGALPROJECTID = "InvalidParameterValue.IllegalProjectId"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_INTERFACENAMENOTFOUND = "InvalidParameterValue.InterfaceNameNotFound"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTFORMAT = "InvalidParameterValue.InvalidAccountFormat"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAME = "InvalidParameterValue.InvalidAccountName"
//  INVALIDPARAMETERVALUE_INVALIDCHARSET = "InvalidParameterValue.InvalidCharset"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCENUM = "InvalidParameterValue.InvalidInstanceNum"
//  INVALIDPARAMETERVALUE_INVALIDORDERNUM = "InvalidParameterValue.InvalidOrderNum"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDFORMAT = "InvalidParameterValue.InvalidPasswordFormat"
//  INVALIDPARAMETERVALUE_INVALIDPID = "InvalidParameterValue.InvalidPid"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  INVALIDPARAMETERVALUE_PARAMETERCHARACTERERROR = "InvalidParameterValue.ParameterCharacterError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_PARAMETERLENGTHLIMITERROR = "InvalidParameterValue.ParameterLengthLimitError"
//  INVALIDPARAMETERVALUE_REGIONNOTSUPPORTED = "InvalidParameterValue.RegionNotSupported"
//  INVALIDPARAMETERVALUE_SPECNOTRECOGNIZEDERROR = "InvalidParameterValue.SpecNotRecognizedError"
//  INVALIDPARAMETERVALUE_STRUCTPARSEFAILED = "InvalidParameterValue.StructParseFailed"
//  INVALIDPID = "InvalidPid"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  OPERATIONDENIED_USERNOTAUTHENTICATEDERROR = "OperationDenied.UserNotAuthenticatedError"
//  OPERATIONDENIED_VPCDENIEDERROR = "OperationDenied.VpcDeniedError"
//  RESOURCEINSUFFICIENT_RESOURCENOTENOUGH = "ResourceInsufficient.ResourceNotEnough"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
//  RESOURCEUNAVAILABLE_RESOURCENOPERMISSION = "ResourceUnavailable.ResourceNoPermission"
//  RESOURCEUNAVAILABLE_VPCRESOURCENOTFOUND = "ResourceUnavailable.VpcResourceNotFound"
//  UNAUTHORIZEDOPERATION_USERHASNOREALNAMEAUTHENTICATION = "UnauthorizedOperation.UserHasNoRealnameAuthentication"
//  UNKNOWNERROR = "UnknownError"
func (c *Client) CreateDBInstances(request *CreateDBInstancesRequest) (response *CreateDBInstancesResponse, err error) {
    if request == nil {
        request = NewCreateDBInstancesRequest()
    }
    response = NewCreateDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstancesRequest() (request *CreateInstancesRequest) {
    request = &CreateInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "CreateInstances")
    
    return
}

func NewCreateInstancesResponse() (response *CreateInstancesResponse) {
    response = &CreateInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateInstances
// This API is used to create and initialize one or more TencentDB for PostgreSQL instances.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_CDBCGWCONNECTERROR = "FailedOperation.CdbCgwConnectError"
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_GETSUBNETERROR = "FailedOperation.GetSubnetError"
//  FAILEDOPERATION_GETVPCINFOERROR = "FailedOperation.GetVpcInfoError"
//  FAILEDOPERATION_INVALIDTRADEOPERATE = "FailedOperation.InvalidTradeOperate"
//  FAILEDOPERATION_PAYORDERFAILED = "FailedOperation.PayOrderFailed"
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  FAILEDOPERATION_QUERYSPECERROR = "FailedOperation.QuerySpecError"
//  FAILEDOPERATION_QUERYVPCFAILED = "FailedOperation.QueryVpcFailed"
//  FAILEDOPERATION_QUERYVPCFALIED = "FailedOperation.QueryVpcFalied"
//  FAILEDOPERATION_SERVICEACCESSERROR = "FailedOperation.ServiceAccessError"
//  FAILEDOPERATION_STORAGEMEMORYCHECKERROR = "FailedOperation.StorageMemoryCheckError"
//  FAILEDOPERATION_TRADECREATEERROR = "FailedOperation.TradeCreateError"
//  FAILEDOPERATION_TRADEPAYORDERSERROR = "FailedOperation.TradePayOrdersError"
//  INTERNALERROR_CGWERROR = "InternalError.CgwError"
//  INTERNALERROR_CNSERROR = "InternalError.CnsError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_FLOWERROR = "InternalError.FlowError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETER_TRADEACCESSDENIEDERROR = "InvalidParameter.TradeAccessDeniedError"
//  INVALIDPARAMETER_VPCNOTFOUNDERROR = "InvalidParameter.VpcNotFoundError"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_BADSPEC = "InvalidParameterValue.BadSpec"
//  INVALIDPARAMETERVALUE_CHARSETNOTFOUNDERROR = "InvalidParameterValue.CharsetNotFoundError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCECHARGETYPE = "InvalidParameterValue.IllegalInstanceChargeType"
//  INVALIDPARAMETERVALUE_ILLEGALPROJECTID = "InvalidParameterValue.IllegalProjectId"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_INTERFACENAMENOTFOUND = "InvalidParameterValue.InterfaceNameNotFound"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTERROR = "InvalidParameterValue.InvalidAccountError"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTFORMAT = "InvalidParameterValue.InvalidAccountFormat"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAME = "InvalidParameterValue.InvalidAccountName"
//  INVALIDPARAMETERVALUE_INVALIDCHARSET = "InvalidParameterValue.InvalidCharset"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCENUM = "InvalidParameterValue.InvalidInstanceNum"
//  INVALIDPARAMETERVALUE_INVALIDORDERNUM = "InvalidParameterValue.InvalidOrderNum"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDFORMAT = "InvalidParameterValue.InvalidPasswordFormat"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDLENGTHERROR = "InvalidParameterValue.InvalidPasswordLengthError"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDVALUEERROR = "InvalidParameterValue.InvalidPasswordValueError"
//  INVALIDPARAMETERVALUE_INVALIDPID = "InvalidParameterValue.InvalidPid"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  INVALIDPARAMETERVALUE_PARAMETERCHARACTERERROR = "InvalidParameterValue.ParameterCharacterError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_PARAMETERLENGTHLIMITERROR = "InvalidParameterValue.ParameterLengthLimitError"
//  INVALIDPARAMETERVALUE_REGIONNOTSUPPORTED = "InvalidParameterValue.RegionNotSupported"
//  INVALIDPARAMETERVALUE_SPECNOTRECOGNIZEDERROR = "InvalidParameterValue.SpecNotRecognizedError"
//  INVALIDPARAMETERVALUE_STRUCTPARSEFAILED = "InvalidParameterValue.StructParseFailed"
//  INVALIDPID = "InvalidPid"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  OPERATIONDENIED_USERNOTAUTHENTICATEDERROR = "OperationDenied.UserNotAuthenticatedError"
//  OPERATIONDENIED_VPCDENIEDERROR = "OperationDenied.VpcDeniedError"
//  RESOURCEINSUFFICIENT_RESOURCENOTENOUGH = "ResourceInsufficient.ResourceNotEnough"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
//  RESOURCEUNAVAILABLE_RESOURCENOPERMISSION = "ResourceUnavailable.ResourceNoPermission"
//  RESOURCEUNAVAILABLE_VPCRESOURCENOTFOUND = "ResourceUnavailable.VpcResourceNotFound"
//  UNAUTHORIZEDOPERATION_USERHASNOREALNAMEAUTHENTICATION = "UnauthorizedOperation.UserHasNoRealnameAuthentication"
//  UNKNOWNERROR = "UnknownError"
func (c *Client) CreateInstances(request *CreateInstancesRequest) (response *CreateInstancesResponse, err error) {
    if request == nil {
        request = NewCreateInstancesRequest()
    }
    response = NewCreateInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateReadOnlyDBInstanceRequest() (request *CreateReadOnlyDBInstanceRequest) {
    request = &CreateReadOnlyDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "CreateReadOnlyDBInstance")
    
    return
}

func NewCreateReadOnlyDBInstanceResponse() (response *CreateReadOnlyDBInstanceResponse) {
    response = &CreateReadOnlyDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateReadOnlyDBInstance
// This API is used to create read-only replicas.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_CDBCGWCONNECTERROR = "FailedOperation.CdbCgwConnectError"
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_GETSUBNETERROR = "FailedOperation.GetSubnetError"
//  FAILEDOPERATION_GETVPCINFOERROR = "FailedOperation.GetVpcInfoError"
//  FAILEDOPERATION_INVALIDTRADEOPERATE = "FailedOperation.InvalidTradeOperate"
//  FAILEDOPERATION_PAYORDERFAILED = "FailedOperation.PayOrderFailed"
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  FAILEDOPERATION_QUERYSPECERROR = "FailedOperation.QuerySpecError"
//  FAILEDOPERATION_QUERYVPCFAILED = "FailedOperation.QueryVpcFailed"
//  FAILEDOPERATION_QUERYVPCFALIED = "FailedOperation.QueryVpcFalied"
//  FAILEDOPERATION_ROGROUPNOTFOUNDERROR = "FailedOperation.ROGroupNotFoundError"
//  FAILEDOPERATION_ROGROUPNUMEXCEEDERROR = "FailedOperation.ROGroupNumExceedError"
//  FAILEDOPERATION_SERVICEACCESSERROR = "FailedOperation.ServiceAccessError"
//  FAILEDOPERATION_STORAGEMEMORYCHECKERROR = "FailedOperation.StorageMemoryCheckError"
//  FAILEDOPERATION_TRADECREATEERROR = "FailedOperation.TradeCreateError"
//  FAILEDOPERATION_TRADEPAYORDERSERROR = "FailedOperation.TradePayOrdersError"
//  INTERNALERROR_CGWERROR = "InternalError.CgwError"
//  INTERNALERROR_CNSERROR = "InternalError.CnsError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_FLOWERROR = "InternalError.FlowError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETER_TRADEACCESSDENIEDERROR = "InvalidParameter.TradeAccessDeniedError"
//  INVALIDPARAMETER_VPCNOTFOUNDERROR = "InvalidParameter.VpcNotFoundError"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_BADSPEC = "InvalidParameterValue.BadSpec"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCECHARGETYPE = "InvalidParameterValue.IllegalInstanceChargeType"
//  INVALIDPARAMETERVALUE_ILLEGALPROJECTID = "InvalidParameterValue.IllegalProjectId"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_INTERFACENAMENOTFOUND = "InvalidParameterValue.InterfaceNameNotFound"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTFORMAT = "InvalidParameterValue.InvalidAccountFormat"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTNAME = "InvalidParameterValue.InvalidAccountName"
//  INVALIDPARAMETERVALUE_INVALIDCHARSET = "InvalidParameterValue.InvalidCharset"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCENUM = "InvalidParameterValue.InvalidInstanceNum"
//  INVALIDPARAMETERVALUE_INVALIDORDERNUM = "InvalidParameterValue.InvalidOrderNum"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDFORMAT = "InvalidParameterValue.InvalidPasswordFormat"
//  INVALIDPARAMETERVALUE_INVALIDPID = "InvalidParameterValue.InvalidPid"
//  INVALIDPARAMETERVALUE_PARAMETERCHARACTERERROR = "InvalidParameterValue.ParameterCharacterError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_PARAMETERLENGTHLIMITERROR = "InvalidParameterValue.ParameterLengthLimitError"
//  INVALIDPARAMETERVALUE_REGIONNOTSUPPORTED = "InvalidParameterValue.RegionNotSupported"
//  INVALIDPARAMETERVALUE_STRUCTPARSEFAILED = "InvalidParameterValue.StructParseFailed"
//  INVALIDPID = "InvalidPid"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  OPERATIONDENIED_PAYMODEERROR = "OperationDenied.PayModeError"
//  OPERATIONDENIED_ROGROUPSTATUSERROR = "OperationDenied.ROGroupStatusError"
//  OPERATIONDENIED_ROINSTANCECOUNTEXEEDERROR = "OperationDenied.RoInstanceCountExeedError"
//  OPERATIONDENIED_USERNOTAUTHENTICATEDERROR = "OperationDenied.UserNotAuthenticatedError"
//  OPERATIONDENIED_VPCDENIEDERROR = "OperationDenied.VpcDeniedError"
//  RESOURCEINSUFFICIENT_RESOURCENOTENOUGH = "ResourceInsufficient.ResourceNotEnough"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
//  RESOURCEUNAVAILABLE_RESOURCENOPERMISSION = "ResourceUnavailable.ResourceNoPermission"
//  RESOURCEUNAVAILABLE_VPCRESOURCENOTFOUND = "ResourceUnavailable.VpcResourceNotFound"
//  UNAUTHORIZEDOPERATION_USERHASNOREALNAMEAUTHENTICATION = "UnauthorizedOperation.UserHasNoRealnameAuthentication"
//  UNKNOWNERROR = "UnknownError"
func (c *Client) CreateReadOnlyDBInstance(request *CreateReadOnlyDBInstanceRequest) (response *CreateReadOnlyDBInstanceResponse, err error) {
    if request == nil {
        request = NewCreateReadOnlyDBInstanceRequest()
    }
    response = NewCreateReadOnlyDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateReadOnlyGroupRequest() (request *CreateReadOnlyGroupRequest) {
    request = &CreateReadOnlyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "CreateReadOnlyGroup")
    
    return
}

func NewCreateReadOnlyGroupResponse() (response *CreateReadOnlyGroupResponse) {
    response = &CreateReadOnlyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateReadOnlyGroup
// This API is used to create an RO group.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_GETVPCINFOERROR = "FailedOperation.GetVpcInfoError"
//  FAILEDOPERATION_ILLEGALROINSTANCENUM = "FailedOperation.IllegalROInstanceNum"
//  FAILEDOPERATION_MASTERINSTANCEQUERYERROR = "FailedOperation.MasterInstanceQueryError"
//  FAILEDOPERATION_ROGROUPNUMEXCEEDERROR = "FailedOperation.ROGroupNumExceedError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETER_VPCNOTFOUNDERROR = "InvalidParameter.VpcNotFoundError"
//  INVALIDPARAMETERVALUE_INTERFACENAMENOTFOUND = "InvalidParameterValue.InterfaceNameNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_PARAMETERLENGTHLIMITERROR = "InvalidParameterValue.ParameterLengthLimitError"
//  INVALIDPARAMETERVALUE_REGIONNOTSUPPORTED = "InvalidParameterValue.RegionNotSupported"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  STRUCTPARSEFAILED = "StructParseFailed"
//  UNKNOWNERROR = "UnknownError"
func (c *Client) CreateReadOnlyGroup(request *CreateReadOnlyGroupRequest) (response *CreateReadOnlyGroupResponse, err error) {
    if request == nil {
        request = NewCreateReadOnlyGroupRequest()
    }
    response = NewCreateReadOnlyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateServerlessDBInstanceRequest() (request *CreateServerlessDBInstanceRequest) {
    request = &CreateServerlessDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "CreateServerlessDBInstance")
    
    return
}

func NewCreateServerlessDBInstanceResponse() (response *CreateServerlessDBInstanceResponse) {
    response = &CreateServerlessDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateServerlessDBInstance
// This API is used to create a PostgreSQL for Serverless instance. If the creation succeeds, the instance ID will be returned.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_GETVPCINFOERROR = "FailedOperation.GetVpcInfoError"
//  FAILEDOPERATION_QUERYVPCFAILED = "FailedOperation.QueryVpcFailed"
//  FLOWERROR = "FlowError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_FLOWERROR = "InternalError.FlowError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENAMEEXIST = "InvalidParameter.InstanceNameExist"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_INSTANCENAMEEXIST = "InvalidParameterValue.InstanceNameExist"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTFORMAT = "InvalidParameterValue.InvalidAccountFormat"
//  INVALIDPARAMETERVALUE_INVALIDCHARSET = "InvalidParameterValue.InvalidCharset"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDFORMAT = "InvalidParameterValue.InvalidPasswordFormat"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_REGIONNOTSUPPORTED = "InvalidParameterValue.RegionNotSupported"
//  INVALIDPARAMETERVALUE_STRUCTPARSEFAILED = "InvalidParameterValue.StructParseFailed"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_NOTSUPPORTZONEERROR = "OperationDenied.NotSupportZoneError"
//  OPERATIONDENIED_VERSIONNOTSUPPORTERROR = "OperationDenied.VersionNotSupportError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
//  UNAUTHORIZEDOPERATION_USERHASNOREALNAMEAUTHENTICATION = "UnauthorizedOperation.UserHasNoRealnameAuthentication"
func (c *Client) CreateServerlessDBInstance(request *CreateServerlessDBInstanceRequest) (response *CreateServerlessDBInstanceResponse, err error) {
    if request == nil {
        request = NewCreateServerlessDBInstanceRequest()
    }
    response = NewCreateServerlessDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteReadOnlyGroupRequest() (request *DeleteReadOnlyGroupRequest) {
    request = &DeleteReadOnlyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "DeleteReadOnlyGroup")
    
    return
}

func NewDeleteReadOnlyGroupResponse() (response *DeleteReadOnlyGroupResponse) {
    response = &DeleteReadOnlyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteReadOnlyGroup
// This API is used to delete an RO group.
//
// error code that may be returned:
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_ROGROUPCANNOTBEDELETEDERROR = "FailedOperation.ROGroupCannotBeDeletedError"
//  INTERNALERROR_FLOWERROR = "InternalError.FlowError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INTERFACENAMENOTFOUND = "InvalidParameterValue.InterfaceNameNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  SYSTEMERROR = "SystemError"
func (c *Client) DeleteReadOnlyGroup(request *DeleteReadOnlyGroupRequest) (response *DeleteReadOnlyGroupResponse, err error) {
    if request == nil {
        request = NewDeleteReadOnlyGroupRequest()
    }
    response = NewDeleteReadOnlyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteServerlessDBInstanceRequest() (request *DeleteServerlessDBInstanceRequest) {
    request = &DeleteServerlessDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "DeleteServerlessDBInstance")
    
    return
}

func NewDeleteServerlessDBInstanceResponse() (response *DeleteServerlessDBInstanceResponse) {
    response = &DeleteServerlessDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteServerlessDBInstance
// This API is used to delete a PostgreSQL for Serverless instance.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_FLOWERROR = "InternalError.FlowError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDINSTANCESTATUS = "InvalidInstanceStatus"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_STRUCTPARSEFAILED = "InvalidParameterValue.StructParseFailed"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  REGIONNOTSUPPORTED = "RegionNotSupported"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
func (c *Client) DeleteServerlessDBInstance(request *DeleteServerlessDBInstanceRequest) (response *DeleteServerlessDBInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteServerlessDBInstanceRequest()
    }
    response = NewDeleteServerlessDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountsRequest() (request *DescribeAccountsRequest) {
    request = &DescribeAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeAccounts")
    
    return
}

func NewDescribeAccountsResponse() (response *DescribeAccountsResponse) {
    response = &DescribeAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccounts
// This API is used to get the instance user list.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXIST = "InvalidParameterValue.AccountNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeAccounts(request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeAccountsRequest()
    }
    response = NewDescribeAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBBackupsRequest() (request *DescribeDBBackupsRequest) {
    request = &DescribeDBBackupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeDBBackups")
    
    return
}

func NewDescribeDBBackupsResponse() (response *DescribeDBBackupsResponse) {
    response = &DescribeDBBackupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDBBackups
// This API is used to query the instance backup list.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_OSSACCESSERROR = "FailedOperation.OssAccessError"
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  FAILEDOPERATION_PRESIGNEDURLGETERROR = "FailedOperation.PresignedURLGetError"
//  FAILEDOPERATION_STOREPATHSPLITERROR = "FailedOperation.StorePathSplitError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INSTANCEDATAERROR = "InternalError.InstanceDataError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNKNOWNERROR = "UnknownError"
func (c *Client) DescribeDBBackups(request *DescribeDBBackupsRequest) (response *DescribeDBBackupsResponse, err error) {
    if request == nil {
        request = NewDescribeDBBackupsRequest()
    }
    response = NewDescribeDBBackupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBErrlogsRequest() (request *DescribeDBErrlogsRequest) {
    request = &DescribeDBErrlogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeDBErrlogs")
    
    return
}

func NewDescribeDBErrlogsResponse() (response *DescribeDBErrlogsResponse) {
    response = &DescribeDBErrlogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDBErrlogs
// This API is used to get error logs.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_ESCONNECTERROR = "FailedOperation.ESConnectError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INTERFACENAMENOTFOUND = "InvalidParameterValue.InterfaceNameNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeDBErrlogs(request *DescribeDBErrlogsRequest) (response *DescribeDBErrlogsResponse, err error) {
    if request == nil {
        request = NewDescribeDBErrlogsRequest()
    }
    response = NewDescribeDBErrlogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceAttributeRequest() (request *DescribeDBInstanceAttributeRequest) {
    request = &DescribeDBInstanceAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeDBInstanceAttribute")
    
    return
}

func NewDescribeDBInstanceAttributeResponse() (response *DescribeDBInstanceAttributeResponse) {
    response = &DescribeDBInstanceAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDBInstanceAttribute
// This API is used to query the details of one instance.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_QUERYINSTANCEBATCHERROR = "FailedOperation.QueryInstanceBatchError"
//  FAILEDOPERATION_QUERYSPECBYSPECCODEERROR = "FailedOperation.QuerySpecBySpecCodeError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeDBInstanceAttribute(request *DescribeDBInstanceAttributeRequest) (response *DescribeDBInstanceAttributeResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceAttributeRequest()
    }
    response = NewDescribeDBInstanceAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceParametersRequest() (request *DescribeDBInstanceParametersRequest) {
    request = &DescribeDBInstanceParametersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeDBInstanceParameters")
    
    return
}

func NewDescribeDBInstanceParametersResponse() (response *DescribeDBInstanceParametersResponse) {
    response = &DescribeDBInstanceParametersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDBInstanceParameters
// This API is used to get the list of modifiable parameters of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeDBInstanceParameters(request *DescribeDBInstanceParametersRequest) (response *DescribeDBInstanceParametersResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceParametersRequest()
    }
    response = NewDescribeDBInstanceParametersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstancesRequest() (request *DescribeDBInstancesRequest) {
    request = &DescribeDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeDBInstances")
    
    return
}

func NewDescribeDBInstancesResponse() (response *DescribeDBInstancesResponse) {
    response = &DescribeDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDBInstances
// This API is used to query the details of one or more instances.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_CAMCHECKRESOURCEFAILED = "FailedOperation.CamCheckResourceFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_LIMITOPERATION = "FailedOperation.LimitOperation"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_REGIONNOTSUPPORTED = "InvalidParameterValue.RegionNotSupported"
//  INVALIDPARAMETERVALUE_STRUCTPARSEFAILED = "InvalidParameterValue.StructParseFailed"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
func (c *Client) DescribeDBInstances(request *DescribeDBInstancesRequest) (response *DescribeDBInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstancesRequest()
    }
    response = NewDescribeDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBSlowlogsRequest() (request *DescribeDBSlowlogsRequest) {
    request = &DescribeDBSlowlogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeDBSlowlogs")
    
    return
}

func NewDescribeDBSlowlogsResponse() (response *DescribeDBSlowlogsResponse) {
    response = &DescribeDBSlowlogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDBSlowlogs
// This API was used to get slow query logs. Since it was deprecated on September 1, 2021, it has no longer returned data. Please use the [DescribeSlowQueryList](https://intl.cloud.tencent.com/document/product/409/60540?from_cn_redirect=1) API instead to get slow query logs.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_ESCONNECTERROR = "FailedOperation.ESConnectError"
//  FAILEDOPERATION_ESQUERYERROR = "FailedOperation.ESQueryError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeDBSlowlogs(request *DescribeDBSlowlogsRequest) (response *DescribeDBSlowlogsResponse, err error) {
    if request == nil {
        request = NewDescribeDBSlowlogsRequest()
    }
    response = NewDescribeDBSlowlogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBXlogsRequest() (request *DescribeDBXlogsRequest) {
    request = &DescribeDBXlogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeDBXlogs")
    
    return
}

func NewDescribeDBXlogsResponse() (response *DescribeDBXlogsResponse) {
    response = &DescribeDBXlogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDBXlogs
// This API is used to get the instance Xlog list.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_OSSACCESSERROR = "FailedOperation.OssAccessError"
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  FAILEDOPERATION_PRESIGNEDURLGETERROR = "FailedOperation.PresignedURLGetError"
//  FAILEDOPERATION_STOREPATHSPLITERROR = "FailedOperation.StorePathSplitError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INSTANCEDATAERROR = "InternalError.InstanceDataError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNKNOWNERROR = "UnknownError"
func (c *Client) DescribeDBXlogs(request *DescribeDBXlogsRequest) (response *DescribeDBXlogsResponse, err error) {
    if request == nil {
        request = NewDescribeDBXlogsRequest()
    }
    response = NewDescribeDBXlogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatabasesRequest() (request *DescribeDatabasesRequest) {
    request = &DescribeDatabasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeDatabases")
    
    return
}

func NewDescribeDatabasesResponse() (response *DescribeDatabasesResponse) {
    response = &DescribeDatabasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDatabases
// This API is used to pull the list of databases.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_LIMITOPERATION = "FailedOperation.LimitOperation"
//  FAILEDOPERATION_OSSACCESSERROR = "FailedOperation.OssAccessError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INSTANCEDATAERROR = "InternalError.InstanceDataError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDINSTANCESTATUS = "InvalidInstanceStatus"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
func (c *Client) DescribeDatabases(request *DescribeDatabasesRequest) (response *DescribeDatabasesResponse, err error) {
    if request == nil {
        request = NewDescribeDatabasesRequest()
    }
    response = NewDescribeDatabasesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrdersRequest() (request *DescribeOrdersRequest) {
    request = &DescribeOrdersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeOrders")
    
    return
}

func NewDescribeOrdersResponse() (response *DescribeOrdersResponse) {
    response = &DescribeOrdersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOrders
// This API is used to get order information.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_QUERYDEALFAILED = "FailedOperation.QueryDealFailed"
//  FAILEDOPERATION_TRADEQUERYORDERERROR = "FailedOperation.TradeQueryOrderError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_NULLDEALNAMES = "InvalidParameterValue.NullDealNames"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeOrders(request *DescribeOrdersRequest) (response *DescribeOrdersResponse, err error) {
    if request == nil {
        request = NewDescribeOrdersRequest()
    }
    response = NewDescribeOrdersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeParamsEventRequest() (request *DescribeParamsEventRequest) {
    request = &DescribeParamsEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeParamsEvent")
    
    return
}

func NewDescribeParamsEventResponse() (response *DescribeParamsEventResponse) {
    response = &DescribeParamsEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeParamsEvent
// This API is used to get the details of parameter modification events.
//
// error code that may be returned:
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeParamsEvent(request *DescribeParamsEventRequest) (response *DescribeParamsEventResponse, err error) {
    if request == nil {
        request = NewDescribeParamsEventRequest()
    }
    response = NewDescribeParamsEventResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProductConfigRequest() (request *DescribeProductConfigRequest) {
    request = &DescribeProductConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeProductConfig")
    
    return
}

func NewDescribeProductConfigResponse() (response *DescribeProductConfigResponse) {
    response = &DescribeProductConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProductConfig
// This API is used to query the purchasable specification configuration.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BADSPEC = "InvalidParameterValue.BadSpec"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeProductConfig(request *DescribeProductConfigRequest) (response *DescribeProductConfigResponse, err error) {
    if request == nil {
        request = NewDescribeProductConfigRequest()
    }
    response = NewDescribeProductConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReadOnlyGroupsRequest() (request *DescribeReadOnlyGroupsRequest) {
    request = &DescribeReadOnlyGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeReadOnlyGroups")
    
    return
}

func NewDescribeReadOnlyGroupsResponse() (response *DescribeReadOnlyGroupsResponse) {
    response = &DescribeReadOnlyGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeReadOnlyGroups
// This API is used to query RO group information by specifying the primary instance IDs.
//
// error code that may be returned:
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INTERFACENAMENOTFOUND = "InvalidParameterValue.InterfaceNameNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeReadOnlyGroups(request *DescribeReadOnlyGroupsRequest) (response *DescribeReadOnlyGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeReadOnlyGroupsRequest()
    }
    response = NewDescribeReadOnlyGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegionsRequest() (request *DescribeRegionsRequest) {
    request = &DescribeRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeRegions")
    
    return
}

func NewDescribeRegionsResponse() (response *DescribeRegionsResponse) {
    response = &DescribeRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRegions
// This API is used to query the purchasable regions.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INTERFACENAMENOTFOUND = "InvalidParameterValue.InterfaceNameNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_REGIONNOTSUPPORTED = "InvalidParameterValue.RegionNotSupported"
//  INVALIDPARAMETERVALUE_STRUCTPARSEFAILED = "InvalidParameterValue.StructParseFailed"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeRegions(request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeRegionsRequest()
    }
    response = NewDescribeRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServerlessDBInstancesRequest() (request *DescribeServerlessDBInstancesRequest) {
    request = &DescribeServerlessDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeServerlessDBInstances")
    
    return
}

func NewDescribeServerlessDBInstancesResponse() (response *DescribeServerlessDBInstancesResponse) {
    response = &DescribeServerlessDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeServerlessDBInstances
// This API is used to query the details of one or more PostgreSQL for Serverless instances.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DBERROR = "DBError"
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERFACENAMENOTFOUND = "InterfaceNameNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_VPCNOTFOUNDERROR = "InvalidParameter.VpcNotFoundError"
//  INVALIDPARAMETERVALUE_ILLEGALPROJECTID = "InvalidParameterValue.IllegalProjectId"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_VPCRESOURCENOTFOUND = "ResourceUnavailable.VpcResourceNotFound"
//  STRUCTPARSEFAILED = "StructParseFailed"
//  SYSTEMERROR = "SystemError"
//  UNKNOWNERROR = "UnknownError"
func (c *Client) DescribeServerlessDBInstances(request *DescribeServerlessDBInstancesRequest) (response *DescribeServerlessDBInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeServerlessDBInstancesRequest()
    }
    response = NewDescribeServerlessDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlowQueryAnalysisRequest() (request *DescribeSlowQueryAnalysisRequest) {
    request = &DescribeSlowQueryAnalysisRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeSlowQueryAnalysis")
    
    return
}

func NewDescribeSlowQueryAnalysisResponse() (response *DescribeSlowQueryAnalysisResponse) {
    response = &DescribeSlowQueryAnalysisResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSlowQueryAnalysis
// This API is used to count and analyze slow query statements during the specified period of time and return aggregated statistical analysis results which are classified by statement with abstract parameter values.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEEXCEEDERROR = "InvalidParameterValue.ParameterValueExceedError"
func (c *Client) DescribeSlowQueryAnalysis(request *DescribeSlowQueryAnalysisRequest) (response *DescribeSlowQueryAnalysisResponse, err error) {
    if request == nil {
        request = NewDescribeSlowQueryAnalysisRequest()
    }
    response = NewDescribeSlowQueryAnalysisResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlowQueryListRequest() (request *DescribeSlowQueryListRequest) {
    request = &DescribeSlowQueryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeSlowQueryList")
    
    return
}

func NewDescribeSlowQueryListResponse() (response *DescribeSlowQueryListResponse) {
    response = &DescribeSlowQueryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSlowQueryList
// This API is used to get the slow queries during the specified period of time.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEEXCEEDERROR = "InvalidParameterValue.ParameterValueExceedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeSlowQueryList(request *DescribeSlowQueryListRequest) (response *DescribeSlowQueryListResponse, err error) {
    if request == nil {
        request = NewDescribeSlowQueryListRequest()
    }
    response = NewDescribeSlowQueryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZonesRequest() (request *DescribeZonesRequest) {
    request = &DescribeZonesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "DescribeZones")
    
    return
}

func NewDescribeZonesResponse() (response *DescribeZonesResponse) {
    response = &DescribeZonesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeZones
// This API is used to query the supported AZs.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_REGIONNOTSUPPORTED = "InvalidParameterValue.RegionNotSupported"
//  INVALIDPARAMETERVALUE_STRUCTPARSEFAILED = "InvalidParameterValue.StructParseFailed"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeZones(request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    if request == nil {
        request = NewDescribeZonesRequest()
    }
    response = NewDescribeZonesResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyDBInstanceRequest() (request *DestroyDBInstanceRequest) {
    request = &DestroyDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "DestroyDBInstance")
    
    return
}

func NewDestroyDBInstanceResponse() (response *DestroyDBInstanceResponse) {
    response = &DestroyDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DestroyDBInstance
// This API is used to eliminate an isolated instance by specifying the `DBInstanceId` parameter. The data of an eliminated instance will be deleted and cannot be recovered.
//
// error code that may be returned:
//  ACCOUNTNOTEXIST = "AccountNotExist"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CMQRESPONSEERROR = "FailedOperation.CMQResponseError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DELETERESOURCEPROJECTTAGERROR = "FailedOperation.DeleteResourceProjectTagError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_LIMITOPERATION = "FailedOperation.LimitOperation"
//  FAILEDOPERATION_POSTPAIDUNBLOCKERROR = "FailedOperation.PostPaidUnblockError"
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  FAILEDOPERATION_TRADEACCESSERROR = "FailedOperation.TradeAccessError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_FLOWERROR = "InternalError.FlowError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDACCOUNTSTATUS = "InvalidAccountStatus"
//  INVALIDINSTANCESTATUS = "InvalidInstanceStatus"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCECHARGETYPE = "InvalidParameterValue.IllegalInstanceChargeType"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_STRUCTPARSEFAILED = "InvalidParameterValue.StructParseFailed"
//  LIMITOPERATION = "LimitOperation"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
//  STRUCTPARSEFAILED = "StructParseFailed"
//  SYSTEMERROR = "SystemError"
//  UNKNOWNERROR = "UnknownError"
func (c *Client) DestroyDBInstance(request *DestroyDBInstanceRequest) (response *DestroyDBInstanceResponse, err error) {
    if request == nil {
        request = NewDestroyDBInstanceRequest()
    }
    response = NewDestroyDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDisIsolateDBInstancesRequest() (request *DisIsolateDBInstancesRequest) {
    request = &DisIsolateDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "DisIsolateDBInstances")
    
    return
}

func NewDisIsolateDBInstancesResponse() (response *DisIsolateDBInstancesResponse) {
    response = &DisIsolateDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisIsolateDBInstances
// This API is used to remove one or more instances from isolation.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_CAMCHECKRESOURCEFAILED = "FailedOperation.CamCheckResourceFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_INVALIDTRADEOPERATE = "FailedOperation.InvalidTradeOperate"
//  FAILEDOPERATION_QUERYTRADESTATUSERROR = "FailedOperation.QueryTradeStatusError"
//  FAILEDOPERATION_TRADEACCESSERROR = "FailedOperation.TradeAccessError"
//  FAILEDOPERATION_TRADECREATEERROR = "FailedOperation.TradeCreateError"
//  FAILEDOPERATION_TRADEPAYORDERSERROR = "FailedOperation.TradePayOrdersError"
//  FAILEDOPERATION_TRADEQUERYORDERERROR = "FailedOperation.TradeQueryOrderError"
//  FAILEDOPERATION_TRADEQUERYPRICEERROR = "FailedOperation.TradeQueryPriceError"
//  INVALIDINSTANCESTATUS = "InvalidInstanceStatus"
//  INVALIDORDERNUM = "InvalidOrderNum"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETER_TRADEACCESSDENIEDERROR = "InvalidParameter.TradeAccessDeniedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERCHARACTERLIMITERROR = "InvalidParameterValue.ParameterCharacterLimitError"
//  INVALIDPARAMETERVALUE_PARAMETERLENGTHLIMITERROR = "InvalidParameterValue.ParameterLengthLimitError"
//  INVALIDPARAMETERVALUE_STRUCTPARSEFAILED = "InvalidParameterValue.StructParseFailed"
//  INVALIDTRADEOPERATE = "InvalidTradeOperate"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITOPERATION = "LimitOperation"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DisIsolateDBInstances(request *DisIsolateDBInstancesRequest) (response *DisIsolateDBInstancesResponse, err error) {
    if request == nil {
        request = NewDisIsolateDBInstancesRequest()
    }
    response = NewDisIsolateDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewInitDBInstancesRequest() (request *InitDBInstancesRequest) {
    request = &InitDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "InitDBInstances")
    
    return
}

func NewInitDBInstancesResponse() (response *InitDBInstancesResponse) {
    response = &InitDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InitDBInstances
// This API is used to initialize a TencentDB for PostgreSQL instance.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_CHARSETNOTFOUNDERROR = "InvalidParameterValue.CharsetNotFoundError"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTERROR = "InvalidParameterValue.InvalidAccountError"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTFORMAT = "InvalidParameterValue.InvalidAccountFormat"
//  INVALIDPARAMETERVALUE_INVALIDCHARSET = "InvalidParameterValue.InvalidCharset"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDFORMAT = "InvalidParameterValue.InvalidPasswordFormat"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDLENGTHERROR = "InvalidParameterValue.InvalidPasswordLengthError"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDVALUEERROR = "InvalidParameterValue.InvalidPasswordValueError"
//  INVALIDPARAMETERVALUE_PARAMETERCHARACTERLIMITERROR = "InvalidParameterValue.ParameterCharacterLimitError"
//  INVALIDPARAMETERVALUE_PARAMETERCHARACTERPRELIMITERROR = "InvalidParameterValue.ParameterCharacterPreLimitError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
func (c *Client) InitDBInstances(request *InitDBInstancesRequest) (response *InitDBInstancesResponse, err error) {
    if request == nil {
        request = NewInitDBInstancesRequest()
    }
    response = NewInitDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceCreateDBInstancesRequest() (request *InquiryPriceCreateDBInstancesRequest) {
    request = &InquiryPriceCreateDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "InquiryPriceCreateDBInstances")
    
    return
}

func NewInquiryPriceCreateDBInstancesResponse() (response *InquiryPriceCreateDBInstancesResponse) {
    response = &InquiryPriceCreateDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquiryPriceCreateDBInstances
// This API is used to query the purchase price of one or multiple instances.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  FAILEDOPERATION_QUERYSPECBYSPECCODEERROR = "FailedOperation.QuerySpecBySpecCodeError"
//  FAILEDOPERATION_TRADEQUERYPRICEERROR = "FailedOperation.TradeQueryPriceError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BADSPEC = "InvalidParameterValue.BadSpec"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCECHARGETYPE = "InvalidParameterValue.IllegalInstanceChargeType"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_SPECNOTRECOGNIZEDERROR = "InvalidParameterValue.SpecNotRecognizedError"
//  INVALIDPID = "InvalidPid"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_PAYMODEERROR = "OperationDenied.PayModeError"
func (c *Client) InquiryPriceCreateDBInstances(request *InquiryPriceCreateDBInstancesRequest) (response *InquiryPriceCreateDBInstancesResponse, err error) {
    if request == nil {
        request = NewInquiryPriceCreateDBInstancesRequest()
    }
    response = NewInquiryPriceCreateDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceRenewDBInstanceRequest() (request *InquiryPriceRenewDBInstanceRequest) {
    request = &InquiryPriceRenewDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "InquiryPriceRenewDBInstance")
    
    return
}

func NewInquiryPriceRenewDBInstanceResponse() (response *InquiryPriceRenewDBInstanceResponse) {
    response = &InquiryPriceRenewDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquiryPriceRenewDBInstance
// This API is used to query the renewal price of an instance.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  FAILEDOPERATION_TRADEQUERYPRICEERROR = "FailedOperation.TradeQueryPriceError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BADSPEC = "InvalidParameterValue.BadSpec"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCECHARGETYPE = "InvalidParameterValue.IllegalInstanceChargeType"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_POSTPAIDPAYMODEERROR = "OperationDenied.PostPaidPayModeError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) InquiryPriceRenewDBInstance(request *InquiryPriceRenewDBInstanceRequest) (response *InquiryPriceRenewDBInstanceResponse, err error) {
    if request == nil {
        request = NewInquiryPriceRenewDBInstanceRequest()
    }
    response = NewInquiryPriceRenewDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceUpgradeDBInstanceRequest() (request *InquiryPriceUpgradeDBInstanceRequest) {
    request = &InquiryPriceUpgradeDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "InquiryPriceUpgradeDBInstance")
    
    return
}

func NewInquiryPriceUpgradeDBInstanceResponse() (response *InquiryPriceUpgradeDBInstanceResponse) {
    response = &InquiryPriceUpgradeDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquiryPriceUpgradeDBInstance
// This API is used to query the upgrade price of an instance.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_CDBCGWCONNECTERROR = "FailedOperation.CdbCgwConnectError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_GETPURCHASERECORDFAILED = "FailedOperation.GetPurchaseRecordFailed"
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  FAILEDOPERATION_TRADEQUERYPRICEERROR = "FailedOperation.TradeQueryPriceError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BADSPEC = "InvalidParameterValue.BadSpec"
//  INVALIDPARAMETERVALUE_ILLEGALCHARGETYPE = "InvalidParameterValue.IllegalChargeType"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCECHARGETYPE = "InvalidParameterValue.IllegalInstanceChargeType"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEVOLUME = "InvalidParameterValue.InvalidInstanceVolume"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
func (c *Client) InquiryPriceUpgradeDBInstance(request *InquiryPriceUpgradeDBInstanceRequest) (response *InquiryPriceUpgradeDBInstanceResponse, err error) {
    if request == nil {
        request = NewInquiryPriceUpgradeDBInstanceRequest()
    }
    response = NewInquiryPriceUpgradeDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewIsolateDBInstancesRequest() (request *IsolateDBInstancesRequest) {
    request = &IsolateDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "IsolateDBInstances")
    
    return
}

func NewIsolateDBInstancesResponse() (response *IsolateDBInstancesResponse) {
    response = &IsolateDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// IsolateDBInstances
// This API is used to isolate one or more instances.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DBERROR = "DBError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_GETINSTANCEBYRESOURCEIDERROR = "FailedOperation.GetInstanceByResourceIdError"
//  FAILEDOPERATION_INVALIDACCOUNTSTATUS = "FailedOperation.InvalidAccountStatus"
//  FAILEDOPERATION_QUERYTRADESTATUSERROR = "FailedOperation.QueryTradeStatusError"
//  FAILEDOPERATION_TRADEACCESSERROR = "FailedOperation.TradeAccessError"
//  FAILEDOPERATION_TRADECREATEERROR = "FailedOperation.TradeCreateError"
//  FAILEDOPERATION_TRADEPAYORDERSERROR = "FailedOperation.TradePayOrdersError"
//  FAILEDOPERATION_TRADEQUERYORDERERROR = "FailedOperation.TradeQueryOrderError"
//  FAILEDOPERATION_TRADEQUERYPRICEERROR = "FailedOperation.TradeQueryPriceError"
//  INSTANCENOTEXIST = "InstanceNotExist"
//  INTERNALERROR_INSTANCEDATAERROR = "InternalError.InstanceDataError"
//  INVALIDINSTANCESTATUS = "InvalidInstanceStatus"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TRADEACCESSDENIEDERROR = "InvalidParameter.TradeAccessDeniedError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDVOUCHERID = "InvalidVoucherId"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITOPERATION = "LimitOperation"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_DTSINSTANCESTATUSERROR = "OperationDenied.DTSInstanceStatusError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  REGIONNOTSUPPORTED = "RegionNotSupported"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) IsolateDBInstances(request *IsolateDBInstancesRequest) (response *IsolateDBInstancesResponse, err error) {
    if request == nil {
        request = NewIsolateDBInstancesRequest()
    }
    response = NewIsolateDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccountRemarkRequest() (request *ModifyAccountRemarkRequest) {
    request = &ModifyAccountRemarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "ModifyAccountRemark")
    
    return
}

func NewModifyAccountRemarkResponse() (response *ModifyAccountRemarkResponse) {
    response = &ModifyAccountRemarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAccountRemark
// This API is used to modify account remarks.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTFORMAT = "InvalidParameterValue.InvalidAccountFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERCHARACTERERROR = "InvalidParameterValue.ParameterCharacterError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_PARAMETERLENGTHLIMITERROR = "InvalidParameterValue.ParameterLengthLimitError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ModifyAccountRemark(request *ModifyAccountRemarkRequest) (response *ModifyAccountRemarkResponse, err error) {
    if request == nil {
        request = NewModifyAccountRemarkRequest()
    }
    response = NewModifyAccountRemarkResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceNameRequest() (request *ModifyDBInstanceNameRequest) {
    request = &ModifyDBInstanceNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "ModifyDBInstanceName")
    
    return
}

func NewModifyDBInstanceNameResponse() (response *ModifyDBInstanceNameResponse) {
    response = &ModifyDBInstanceNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDBInstanceName
// This API is used to rename a TencentDB for PostgreSQL instance.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_OSSACCESSERROR = "FailedOperation.OssAccessError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDACCOUNTFORMAT = "InvalidParameterValue.InvalidAccountFormat"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERCHARACTERERROR = "InvalidParameterValue.ParameterCharacterError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_PARAMETERLENGTHLIMITERROR = "InvalidParameterValue.ParameterLengthLimitError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ModifyDBInstanceName(request *ModifyDBInstanceNameRequest) (response *ModifyDBInstanceNameResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceNameRequest()
    }
    response = NewModifyDBInstanceNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceParametersRequest() (request *ModifyDBInstanceParametersRequest) {
    request = &ModifyDBInstanceParametersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "ModifyDBInstanceParameters")
    
    return
}

func NewModifyDBInstanceParametersResponse() (response *ModifyDBInstanceParametersResponse) {
    response = &ModifyDBInstanceParametersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDBInstanceParameters
// This API is used to modify parameters in batches.
//
// error code that may be returned:
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ModifyDBInstanceParameters(request *ModifyDBInstanceParametersRequest) (response *ModifyDBInstanceParametersResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceParametersRequest()
    }
    response = NewModifyDBInstanceParametersResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceReadOnlyGroupRequest() (request *ModifyDBInstanceReadOnlyGroupRequest) {
    request = &ModifyDBInstanceReadOnlyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "ModifyDBInstanceReadOnlyGroup")
    
    return
}

func NewModifyDBInstanceReadOnlyGroupResponse() (response *ModifyDBInstanceReadOnlyGroupResponse) {
    response = &ModifyDBInstanceReadOnlyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDBInstanceReadOnlyGroup
// This API is used to modify the RO group of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_ILLEGALROINSTANCENUM = "FailedOperation.IllegalROInstanceNum"
//  FAILEDOPERATION_ROGROUPNOTFOUNDERROR = "FailedOperation.ROGroupNotFoundError"
//  INSTANCENOTEXIST = "InstanceNotExist"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDINSTANCESTATUS = "InvalidInstanceStatus"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREADONLYGROUPSTATUS = "InvalidParameterValue.InvalidReadOnlyGroupStatus"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_READONLYGROUPNOTEXIST = "InvalidParameterValue.ReadOnlyGroupNotExist"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_ROGROUPSTATUSERROR = "OperationDenied.ROGroupStatusError"
//  OPERATIONDENIED_ROINSTANCECOUNTEXEEDERROR = "OperationDenied.RoInstanceCountExeedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
func (c *Client) ModifyDBInstanceReadOnlyGroup(request *ModifyDBInstanceReadOnlyGroupRequest) (response *ModifyDBInstanceReadOnlyGroupResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceReadOnlyGroupRequest()
    }
    response = NewModifyDBInstanceReadOnlyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceSpecRequest() (request *ModifyDBInstanceSpecRequest) {
    request = &ModifyDBInstanceSpecRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "ModifyDBInstanceSpec")
    
    return
}

func NewModifyDBInstanceSpecResponse() (response *ModifyDBInstanceSpecResponse) {
    response = &ModifyDBInstanceSpecResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDBInstanceSpec
// This API is used to modify instance specifications including memory and disk size.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_PAYORDERFAILED = "FailedOperation.PayOrderFailed"
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  FAILEDOPERATION_QUERYSPECERROR = "FailedOperation.QuerySpecError"
func (c *Client) ModifyDBInstanceSpec(request *ModifyDBInstanceSpecRequest) (response *ModifyDBInstanceSpecResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceSpecRequest()
    }
    response = NewModifyDBInstanceSpecResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstancesProjectRequest() (request *ModifyDBInstancesProjectRequest) {
    request = &ModifyDBInstancesProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "ModifyDBInstancesProject")
    
    return
}

func NewModifyDBInstancesProjectResponse() (response *ModifyDBInstancesProjectResponse) {
    response = &ModifyDBInstancesProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDBInstancesProject
// This API is used to transfer an instance to another project.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_CDBCGWCONNECTERROR = "FailedOperation.CdbCgwConnectError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_UPDATERESOURCEPROJECTTAGVALUEERROR = "FailedOperation.UpdateResourceProjectTagValueError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ILLEGALPROJECTID = "InvalidParameterValue.IllegalProjectId"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ModifyDBInstancesProject(request *ModifyDBInstancesProjectRequest) (response *ModifyDBInstancesProjectResponse, err error) {
    if request == nil {
        request = NewModifyDBInstancesProjectRequest()
    }
    response = NewModifyDBInstancesProjectResponse()
    err = c.Send(request, response)
    return
}

func NewModifyReadOnlyGroupConfigRequest() (request *ModifyReadOnlyGroupConfigRequest) {
    request = &ModifyReadOnlyGroupConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "ModifyReadOnlyGroupConfig")
    
    return
}

func NewModifyReadOnlyGroupConfigResponse() (response *ModifyReadOnlyGroupConfigResponse) {
    response = &ModifyReadOnlyGroupConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyReadOnlyGroupConfig
// This API is used to modify RO group configuration.
//
// error code that may be returned:
//  FAILEDOPERATION_BASENETWORKACCESSERROR = "FailedOperation.BaseNetworkAccessError"
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DELETEALLROUTE = "FailedOperation.DeleteAllRoute"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_MODIFYROGROUPERROR = "FailedOperation.ModifyROGroupError"
//  FAILEDOPERATION_ROGROUPNOTFOUNDERROR = "FailedOperation.ROGroupNotFoundError"
//  FAILEDOPERATION_SERVICEACCESSERROR = "FailedOperation.ServiceAccessError"
//  FAILEDOPERATION_VPCRESETSERVICEERROR = "FailedOperation.VPCResetServiceError"
//  FAILEDOPERATION_VPCUPDATEROUTEERROR = "FailedOperation.VPCUpdateRouteError"
//  INTERFACENAMENOTFOUND = "InterfaceNameNotFound"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_READONLYGROUPNOTEXIST = "InvalidParameterValue.ReadOnlyGroupNotExist"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ModifyReadOnlyGroupConfig(request *ModifyReadOnlyGroupConfigRequest) (response *ModifyReadOnlyGroupConfigResponse, err error) {
    if request == nil {
        request = NewModifyReadOnlyGroupConfigRequest()
    }
    response = NewModifyReadOnlyGroupConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifySwitchTimePeriodRequest() (request *ModifySwitchTimePeriodRequest) {
    request = &ModifySwitchTimePeriodRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "ModifySwitchTimePeriod")
    
    return
}

func NewModifySwitchTimePeriodResponse() (response *ModifySwitchTimePeriodResponse) {
    response = &ModifySwitchTimePeriodResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySwitchTimePeriod
// This API is used to perform a primary-standby switch for an instance waiting for the switch after it is upgraded.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_CAMCHECKRESOURCEFAILED = "FailedOperation.CamCheckResourceFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_QUERYSPECERROR = "FailedOperation.QuerySpecError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNKNOWNERROR = "UnknownError"
func (c *Client) ModifySwitchTimePeriod(request *ModifySwitchTimePeriodRequest) (response *ModifySwitchTimePeriodResponse, err error) {
    if request == nil {
        request = NewModifySwitchTimePeriodRequest()
    }
    response = NewModifySwitchTimePeriodResponse()
    err = c.Send(request, response)
    return
}

func NewOpenDBExtranetAccessRequest() (request *OpenDBExtranetAccessRequest) {
    request = &OpenDBExtranetAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "OpenDBExtranetAccess")
    
    return
}

func NewOpenDBExtranetAccessResponse() (response *OpenDBExtranetAccessResponse) {
    response = &OpenDBExtranetAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// OpenDBExtranetAccess
// This API is used to enable public network access.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_FLOWERROR = "InternalError.FlowError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
func (c *Client) OpenDBExtranetAccess(request *OpenDBExtranetAccessRequest) (response *OpenDBExtranetAccessResponse, err error) {
    if request == nil {
        request = NewOpenDBExtranetAccessRequest()
    }
    response = NewOpenDBExtranetAccessResponse()
    err = c.Send(request, response)
    return
}

func NewOpenServerlessDBExtranetAccessRequest() (request *OpenServerlessDBExtranetAccessRequest) {
    request = &OpenServerlessDBExtranetAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "OpenServerlessDBExtranetAccess")
    
    return
}

func NewOpenServerlessDBExtranetAccessResponse() (response *OpenServerlessDBExtranetAccessResponse) {
    response = &OpenServerlessDBExtranetAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// OpenServerlessDBExtranetAccess
// This API is used to enable public network access for a PostgreSQL for Serverless instance.
//
// error code that may be returned:
//  ACCOUNTNOTEXIST = "AccountNotExist"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DBERROR = "DBError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_QUERYVPCFALIED = "FailedOperation.QueryVpcFalied"
//  FLOWERROR = "FlowError"
//  INSTANCENOTEXIST = "InstanceNotExist"
//  INTERFACENAMENOTFOUND = "InterfaceNameNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_FLOWERROR = "InternalError.FlowError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDACCOUNTSTATUS = "InvalidAccountStatus"
//  INVALIDINSTANCESTATUS = "InvalidInstanceStatus"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INTERFACENAMENOTFOUND = "InvalidParameterValue.InterfaceNameNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_STRUCTPARSEFAILED = "InvalidParameterValue.StructParseFailed"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
//  RESOURCEUNAVAILABLE_RESOURCENOPERMISSION = "ResourceUnavailable.ResourceNoPermission"
//  RESOURCEUNAVAILABLE_VPCRESOURCENOTFOUND = "ResourceUnavailable.VpcResourceNotFound"
//  STRUCTPARSEFAILED = "StructParseFailed"
//  SYSTEMERROR = "SystemError"
//  UNKNOWNERROR = "UnknownError"
//  VPCERROR = "VpcError"
func (c *Client) OpenServerlessDBExtranetAccess(request *OpenServerlessDBExtranetAccessRequest) (response *OpenServerlessDBExtranetAccessResponse, err error) {
    if request == nil {
        request = NewOpenServerlessDBExtranetAccessRequest()
    }
    response = NewOpenServerlessDBExtranetAccessResponse()
    err = c.Send(request, response)
    return
}

func NewRebalanceReadOnlyGroupRequest() (request *RebalanceReadOnlyGroupRequest) {
    request = &RebalanceReadOnlyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "RebalanceReadOnlyGroup")
    
    return
}

func NewRebalanceReadOnlyGroupResponse() (response *RebalanceReadOnlyGroupResponse) {
    response = &RebalanceReadOnlyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RebalanceReadOnlyGroup
// This API is used to rebalance the loads of read-only replicas in an RO group. Please note that connections to those read-only replicas will be interrupted transiently; therefore, you should ensure that your application can reconnect to the databases. This operation should be performed with caution.
//
// error code that may be returned:
//  FAILEDOPERATION_BASENETWORKACCESSERROR = "FailedOperation.BaseNetworkAccessError"
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DELETEALLROUTE = "FailedOperation.DeleteAllRoute"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_ROGROUPNOTFOUNDERROR = "FailedOperation.ROGroupNotFoundError"
//  FAILEDOPERATION_VPCRESETSERVICEERROR = "FailedOperation.VPCResetServiceError"
//  FAILEDOPERATION_VPCUPDATEROUTEERROR = "FailedOperation.VPCUpdateRouteError"
//  INTERFACENAMENOTFOUND = "InterfaceNameNotFound"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  VPCERROR = "VpcError"
func (c *Client) RebalanceReadOnlyGroup(request *RebalanceReadOnlyGroupRequest) (response *RebalanceReadOnlyGroupResponse, err error) {
    if request == nil {
        request = NewRebalanceReadOnlyGroupRequest()
    }
    response = NewRebalanceReadOnlyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveDBInstanceFromReadOnlyGroupRequest() (request *RemoveDBInstanceFromReadOnlyGroupRequest) {
    request = &RemoveDBInstanceFromReadOnlyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "RemoveDBInstanceFromReadOnlyGroup")
    
    return
}

func NewRemoveDBInstanceFromReadOnlyGroupResponse() (response *RemoveDBInstanceFromReadOnlyGroupResponse) {
    response = &RemoveDBInstanceFromReadOnlyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RemoveDBInstanceFromReadOnlyGroup
// This API is used to remove a read-only replica from an RO group.
//
// error code that may be returned:
//  FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_ILLEGALROINSTANCENUM = "FailedOperation.IllegalROInstanceNum"
//  FAILEDOPERATION_ROGROUPMASTERINSTANCENOTRIGHT = "FailedOperation.ROGroupMasterInstanceNotRight"
//  FAILEDOPERATION_ROGROUPNOTFOUNDERROR = "FailedOperation.ROGroupNotFoundError"
//  INSTANCENOTEXIST = "InstanceNotExist"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDINSTANCESTATUS = "InvalidInstanceStatus"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREADONLYGROUPSTATUS = "InvalidParameterValue.InvalidReadOnlyGroupStatus"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPARAMETERVALUE_READONLYGROUPNOTEXIST = "InvalidParameterValue.ReadOnlyGroupNotExist"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_ROGROUPSTATUSERROR = "OperationDenied.ROGroupStatusError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
func (c *Client) RemoveDBInstanceFromReadOnlyGroup(request *RemoveDBInstanceFromReadOnlyGroupRequest) (response *RemoveDBInstanceFromReadOnlyGroupResponse, err error) {
    if request == nil {
        request = NewRemoveDBInstanceFromReadOnlyGroupRequest()
    }
    response = NewRemoveDBInstanceFromReadOnlyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewRenewInstanceRequest() (request *RenewInstanceRequest) {
    request = &RenewInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "RenewInstance")
    
    return
}

func NewRenewInstanceResponse() (response *RenewInstanceResponse) {
    response = &RenewInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RenewInstance
// This API is used to renew an instance.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_GETVPCINFOERROR = "FailedOperation.GetVpcInfoError"
//  FAILEDOPERATION_QUERYVPCFAILED = "FailedOperation.QueryVpcFailed"
//  FAILEDOPERATION_QUERYVPCFALIED = "FailedOperation.QueryVpcFalied"
//  FAILEDOPERATION_TRADECREATEERROR = "FailedOperation.TradeCreateError"
//  FAILEDOPERATION_TRADEPAYORDERSERROR = "FailedOperation.TradePayOrdersError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_VPCNOTFOUNDERROR = "InvalidParameter.VpcNotFoundError"
//  INVALIDPARAMETERVALUE_BADSPEC = "InvalidParameterValue.BadSpec"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCECHARGETYPE = "InvalidParameterValue.IllegalInstanceChargeType"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  OPERATIONDENIED_POSTPAIDPAYMODEERROR = "OperationDenied.PostPaidPayModeError"
//  OPERATIONDENIED_VPCDENIEDERROR = "OperationDenied.VpcDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCENOPERMISSION = "ResourceUnavailable.ResourceNoPermission"
//  RESOURCEUNAVAILABLE_VPCRESOURCENOTFOUND = "ResourceUnavailable.VpcResourceNotFound"
func (c *Client) RenewInstance(request *RenewInstanceRequest) (response *RenewInstanceResponse, err error) {
    if request == nil {
        request = NewRenewInstanceRequest()
    }
    response = NewRenewInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewResetAccountPasswordRequest() (request *ResetAccountPasswordRequest) {
    request = &ResetAccountPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "ResetAccountPassword")
    
    return
}

func NewResetAccountPasswordResponse() (response *ResetAccountPasswordResponse) {
    response = &ResetAccountPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetAccountPassword
// This API is used to reset the account password of an instance.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_OSSACCESSERROR = "FailedOperation.OssAccessError"
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDACCOUNTPASSWORD = "InvalidAccountPassword"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDFORMAT = "InvalidParameterValue.InvalidPasswordFormat"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDLENGTHERROR = "InvalidParameterValue.InvalidPasswordLengthError"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDVALUEERROR = "InvalidParameterValue.InvalidPasswordValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ResetAccountPassword(request *ResetAccountPasswordRequest) (response *ResetAccountPasswordResponse, err error) {
    if request == nil {
        request = NewResetAccountPasswordRequest()
    }
    response = NewResetAccountPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewRestartDBInstanceRequest() (request *RestartDBInstanceRequest) {
    request = &RestartDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "RestartDBInstance")
    
    return
}

func NewRestartDBInstanceResponse() (response *RestartDBInstanceResponse) {
    response = &RestartDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RestartDBInstance
// This API is used to restart an instance.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_FLOWERROR = "InternalError.FlowError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
func (c *Client) RestartDBInstance(request *RestartDBInstanceRequest) (response *RestartDBInstanceResponse, err error) {
    if request == nil {
        request = NewRestartDBInstanceRequest()
    }
    response = NewRestartDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewSetAutoRenewFlagRequest() (request *SetAutoRenewFlagRequest) {
    request = &SetAutoRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "SetAutoRenewFlag")
    
    return
}

func NewSetAutoRenewFlagResponse() (response *SetAutoRenewFlagResponse) {
    response = &SetAutoRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetAutoRenewFlag
// This API is used to set auto-renewal.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_SETAUTORENEWFLAGERROR = "FailedOperation.SetAutoRenewFlagError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCECHARGETYPE = "InvalidParameterValue.IllegalInstanceChargeType"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  OPERATIONDENIED_POSTPAIDPAYMODEERROR = "OperationDenied.PostPaidPayModeError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
func (c *Client) SetAutoRenewFlag(request *SetAutoRenewFlagRequest) (response *SetAutoRenewFlagResponse, err error) {
    if request == nil {
        request = NewSetAutoRenewFlagRequest()
    }
    response = NewSetAutoRenewFlagResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeDBInstanceRequest() (request *UpgradeDBInstanceRequest) {
    request = &UpgradeDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("postgres", APIVersion, "UpgradeDBInstance")
    
    return
}

func NewUpgradeDBInstanceResponse() (response *UpgradeDBInstanceResponse) {
    response = &UpgradeDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpgradeDBInstance
// This API is used to upgrade instance configurations.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_GETPURCHASERECORDFAILED = "FailedOperation.GetPurchaseRecordFailed"
//  FAILEDOPERATION_PAYORDERFAILED = "FailedOperation.PayOrderFailed"
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  FAILEDOPERATION_QUERYSPECERROR = "FailedOperation.QuerySpecError"
//  FAILEDOPERATION_STORAGEMEMORYCHECKERROR = "FailedOperation.StorageMemoryCheckError"
//  FAILEDOPERATION_TRADEACCESSERROR = "FailedOperation.TradeAccessError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BADSPEC = "InvalidParameterValue.BadSpec"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEVOLUME = "InvalidParameterValue.InvalidInstanceVolume"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"
//  INVALIDPID = "InvalidPid"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"
func (c *Client) UpgradeDBInstance(request *UpgradeDBInstanceRequest) (response *UpgradeDBInstanceResponse, err error) {
    if request == nil {
        request = NewUpgradeDBInstanceRequest()
    }
    response = NewUpgradeDBInstanceResponse()
    err = c.Send(request, response)
    return
}
