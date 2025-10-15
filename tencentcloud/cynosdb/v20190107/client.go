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

package v20190107

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2019-01-07"

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


func NewActivateInstanceRequest() (request *ActivateInstanceRequest) {
    request = &ActivateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ActivateInstance")
    
    
    return
}

func NewActivateInstanceResponse() (response *ActivateInstanceResponse) {
    response = &ActivateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ActivateInstance
// This interface (ActivateInstance) restores access to isolated instances.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_PREPAYPAYMODEERROR = "InvalidParameterValue.PrePayPayModeError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCENOTFOUND_RESOURCEERROR = "ResourceNotFound.ResourceError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ActivateInstance(request *ActivateInstanceRequest) (response *ActivateInstanceResponse, err error) {
    return c.ActivateInstanceWithContext(context.Background(), request)
}

// ActivateInstance
// This interface (ActivateInstance) restores access to isolated instances.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_PREPAYPAYMODEERROR = "InvalidParameterValue.PrePayPayModeError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCENOTFOUND_RESOURCEERROR = "ResourceNotFound.ResourceError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ActivateInstanceWithContext(ctx context.Context, request *ActivateInstanceRequest) (response *ActivateInstanceResponse, err error) {
    if request == nil {
        request = NewActivateInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ActivateInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ActivateInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewActivateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewAddClusterSlaveZoneRequest() (request *AddClusterSlaveZoneRequest) {
    request = &AddClusterSlaveZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "AddClusterSlaveZone")
    
    
    return
}

func NewAddClusterSlaveZoneResponse() (response *AddClusterSlaveZoneResponse) {
    response = &AddClusterSlaveZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddClusterSlaveZone
// This interface (AddClusterSlaveZone) is used to enable multi-az deployment for a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) AddClusterSlaveZone(request *AddClusterSlaveZoneRequest) (response *AddClusterSlaveZoneResponse, err error) {
    return c.AddClusterSlaveZoneWithContext(context.Background(), request)
}

// AddClusterSlaveZone
// This interface (AddClusterSlaveZone) is used to enable multi-az deployment for a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) AddClusterSlaveZoneWithContext(ctx context.Context, request *AddClusterSlaveZoneRequest) (response *AddClusterSlaveZoneResponse, err error) {
    if request == nil {
        request = NewAddClusterSlaveZoneRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "AddClusterSlaveZone")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddClusterSlaveZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddClusterSlaveZoneResponse()
    err = c.Send(request, response)
    return
}

func NewAddInstancesRequest() (request *AddInstancesRequest) {
    request = &AddInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "AddInstances")
    
    
    return
}

func NewAddInstancesResponse() (response *AddInstancesResponse) {
    response = &AddInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddInstances
// This API is used to add instances to a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  OPERATIONDENIED_INSUFFICIENTBALANCEERROR = "OperationDenied.InsufficientBalanceError"
//  OPERATIONDENIED_USERNOTAUTHENTICATEDERROR = "OperationDenied.UserNotAuthenticatedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) AddInstances(request *AddInstancesRequest) (response *AddInstancesResponse, err error) {
    return c.AddInstancesWithContext(context.Background(), request)
}

// AddInstances
// This API is used to add instances to a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  OPERATIONDENIED_INSUFFICIENTBALANCEERROR = "OperationDenied.InsufficientBalanceError"
//  OPERATIONDENIED_USERNOTAUTHENTICATEDERROR = "OperationDenied.UserNotAuthenticatedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) AddInstancesWithContext(ctx context.Context, request *AddInstancesRequest) (response *AddInstancesResponse, err error) {
    if request == nil {
        request = NewAddInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "AddInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewBindClusterResourcePackagesRequest() (request *BindClusterResourcePackagesRequest) {
    request = &BindClusterResourcePackagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "BindClusterResourcePackages")
    
    
    return
}

func NewBindClusterResourcePackagesResponse() (response *BindClusterResourcePackagesResponse) {
    response = &BindClusterResourcePackagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindClusterResourcePackages
// This API is used to bind resource packages to a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_BINDSOURCEPACKAGEERROR = "FailedOperation.BindSourcePackageError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_QUERYSOURCEPACKAGEERROR = "FailedOperation.QuerySourcePackageError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) BindClusterResourcePackages(request *BindClusterResourcePackagesRequest) (response *BindClusterResourcePackagesResponse, err error) {
    return c.BindClusterResourcePackagesWithContext(context.Background(), request)
}

// BindClusterResourcePackages
// This API is used to bind resource packages to a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_BINDSOURCEPACKAGEERROR = "FailedOperation.BindSourcePackageError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_QUERYSOURCEPACKAGEERROR = "FailedOperation.QuerySourcePackageError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) BindClusterResourcePackagesWithContext(ctx context.Context, request *BindClusterResourcePackagesRequest) (response *BindClusterResourcePackagesResponse, err error) {
    if request == nil {
        request = NewBindClusterResourcePackagesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "BindClusterResourcePackages")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindClusterResourcePackages require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindClusterResourcePackagesResponse()
    err = c.Send(request, response)
    return
}

func NewCloseAuditServiceRequest() (request *CloseAuditServiceRequest) {
    request = &CloseAuditServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "CloseAuditService")
    
    
    return
}

func NewCloseAuditServiceResponse() (response *CloseAuditServiceResponse) {
    response = &CloseAuditServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CloseAuditService
// This API is used to close the database audit service for TDSQL-C MySQL instances.
//
// error code that may be returned:
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) CloseAuditService(request *CloseAuditServiceRequest) (response *CloseAuditServiceResponse, err error) {
    return c.CloseAuditServiceWithContext(context.Background(), request)
}

// CloseAuditService
// This API is used to close the database audit service for TDSQL-C MySQL instances.
//
// error code that may be returned:
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) CloseAuditServiceWithContext(ctx context.Context, request *CloseAuditServiceRequest) (response *CloseAuditServiceResponse, err error) {
    if request == nil {
        request = NewCloseAuditServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "CloseAuditService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseAuditService require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseAuditServiceResponse()
    err = c.Send(request, response)
    return
}

func NewCloseClusterPasswordComplexityRequest() (request *CloseClusterPasswordComplexityRequest) {
    request = &CloseClusterPasswordComplexityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "CloseClusterPasswordComplexity")
    
    
    return
}

func NewCloseClusterPasswordComplexityResponse() (response *CloseClusterPasswordComplexityResponse) {
    response = &CloseClusterPasswordComplexityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CloseClusterPasswordComplexity
// This API is used to close cluster password complexity.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CloseClusterPasswordComplexity(request *CloseClusterPasswordComplexityRequest) (response *CloseClusterPasswordComplexityResponse, err error) {
    return c.CloseClusterPasswordComplexityWithContext(context.Background(), request)
}

// CloseClusterPasswordComplexity
// This API is used to close cluster password complexity.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CloseClusterPasswordComplexityWithContext(ctx context.Context, request *CloseClusterPasswordComplexityRequest) (response *CloseClusterPasswordComplexityResponse, err error) {
    if request == nil {
        request = NewCloseClusterPasswordComplexityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "CloseClusterPasswordComplexity")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseClusterPasswordComplexity require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseClusterPasswordComplexityResponse()
    err = c.Send(request, response)
    return
}

func NewCloseProxyRequest() (request *CloseProxyRequest) {
    request = &CloseProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "CloseProxy")
    
    
    return
}

func NewCloseProxyResponse() (response *CloseProxyResponse) {
    response = &CloseProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CloseProxy
// This API is used to close the database proxy service of a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CloseProxy(request *CloseProxyRequest) (response *CloseProxyResponse, err error) {
    return c.CloseProxyWithContext(context.Background(), request)
}

// CloseProxy
// This API is used to close the database proxy service of a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CloseProxyWithContext(ctx context.Context, request *CloseProxyRequest) (response *CloseProxyResponse, err error) {
    if request == nil {
        request = NewCloseProxyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "CloseProxy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseProxyResponse()
    err = c.Send(request, response)
    return
}

func NewCloseProxyEndPointRequest() (request *CloseProxyEndPointRequest) {
    request = &CloseProxyEndPointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "CloseProxyEndPoint")
    
    
    return
}

func NewCloseProxyEndPointResponse() (response *CloseProxyEndPointResponse) {
    response = &CloseProxyEndPointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CloseProxyEndPoint
// This API is used to close the database proxy connection address.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CloseProxyEndPoint(request *CloseProxyEndPointRequest) (response *CloseProxyEndPointResponse, err error) {
    return c.CloseProxyEndPointWithContext(context.Background(), request)
}

// CloseProxyEndPoint
// This API is used to close the database proxy connection address.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CloseProxyEndPointWithContext(ctx context.Context, request *CloseProxyEndPointRequest) (response *CloseProxyEndPointResponse, err error) {
    if request == nil {
        request = NewCloseProxyEndPointRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "CloseProxyEndPoint")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseProxyEndPoint require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseProxyEndPointResponse()
    err = c.Send(request, response)
    return
}

func NewCloseSSLRequest() (request *CloseSSLRequest) {
    request = &CloseSSLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "CloseSSL")
    
    
    return
}

func NewCloseSSLResponse() (response *CloseSSLResponse) {
    response = &CloseSSLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CloseSSL
// This API is used to disable SSL encryption.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREGIONIDERROR = "InvalidParameterValue.InvalidRegionIdError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CloseSSL(request *CloseSSLRequest) (response *CloseSSLResponse, err error) {
    return c.CloseSSLWithContext(context.Background(), request)
}

// CloseSSL
// This API is used to disable SSL encryption.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREGIONIDERROR = "InvalidParameterValue.InvalidRegionIdError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CloseSSLWithContext(ctx context.Context, request *CloseSSLRequest) (response *CloseSSLResponse, err error) {
    if request == nil {
        request = NewCloseSSLRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "CloseSSL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseSSL require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseSSLResponse()
    err = c.Send(request, response)
    return
}

func NewCloseWanRequest() (request *CloseWanRequest) {
    request = &CloseWanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "CloseWan")
    
    
    return
}

func NewCloseWanResponse() (response *CloseWanResponse) {
    response = &CloseWanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CloseWan
// This interface (CloseWan) is used to disable public network.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_GETNETSERVICEINFOERROR = "FailedOperation.GetNetServiceInfoError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CloseWan(request *CloseWanRequest) (response *CloseWanResponse, err error) {
    return c.CloseWanWithContext(context.Background(), request)
}

// CloseWan
// This interface (CloseWan) is used to disable public network.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_GETNETSERVICEINFOERROR = "FailedOperation.GetNetServiceInfoError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CloseWanWithContext(ctx context.Context, request *CloseWanRequest) (response *CloseWanResponse, err error) {
    if request == nil {
        request = NewCloseWanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "CloseWan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseWan require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseWanResponse()
    err = c.Send(request, response)
    return
}

func NewCopyClusterPasswordComplexityRequest() (request *CopyClusterPasswordComplexityRequest) {
    request = &CopyClusterPasswordComplexityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "CopyClusterPasswordComplexity")
    
    
    return
}

func NewCopyClusterPasswordComplexityResponse() (response *CopyClusterPasswordComplexityResponse) {
    response = &CopyClusterPasswordComplexityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CopyClusterPasswordComplexity
// This API is used to copy the password complexity of a replication cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CopyClusterPasswordComplexity(request *CopyClusterPasswordComplexityRequest) (response *CopyClusterPasswordComplexityResponse, err error) {
    return c.CopyClusterPasswordComplexityWithContext(context.Background(), request)
}

// CopyClusterPasswordComplexity
// This API is used to copy the password complexity of a replication cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CopyClusterPasswordComplexityWithContext(ctx context.Context, request *CopyClusterPasswordComplexityRequest) (response *CopyClusterPasswordComplexityResponse, err error) {
    if request == nil {
        request = NewCopyClusterPasswordComplexityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "CopyClusterPasswordComplexity")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CopyClusterPasswordComplexity require credential")
    }

    request.SetContext(ctx)
    
    response = NewCopyClusterPasswordComplexityResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAccountsRequest() (request *CreateAccountsRequest) {
    request = &CreateAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "CreateAccounts")
    
    
    return
}

func NewCreateAccountsResponse() (response *CreateAccountsResponse) {
    response = &CreateAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAccounts
// This API is used to create user accounts.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_ACCOUNTALREADYEXISTERROR = "InvalidParameterValue.AccountAlreadyExistError"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateAccounts(request *CreateAccountsRequest) (response *CreateAccountsResponse, err error) {
    return c.CreateAccountsWithContext(context.Background(), request)
}

// CreateAccounts
// This API is used to create user accounts.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_ACCOUNTALREADYEXISTERROR = "InvalidParameterValue.AccountAlreadyExistError"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateAccountsWithContext(ctx context.Context, request *CreateAccountsRequest) (response *CreateAccountsResponse, err error) {
    if request == nil {
        request = NewCreateAccountsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "CreateAccounts")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAccounts require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAuditRuleTemplateRequest() (request *CreateAuditRuleTemplateRequest) {
    request = &CreateAuditRuleTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "CreateAuditRuleTemplate")
    
    
    return
}

func NewCreateAuditRuleTemplateResponse() (response *CreateAuditRuleTemplateResponse) {
    response = &CreateAuditRuleTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAuditRuleTemplate
// This API is used to create audit rule templates.
//
// error code that may be returned:
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
func (c *Client) CreateAuditRuleTemplate(request *CreateAuditRuleTemplateRequest) (response *CreateAuditRuleTemplateResponse, err error) {
    return c.CreateAuditRuleTemplateWithContext(context.Background(), request)
}

// CreateAuditRuleTemplate
// This API is used to create audit rule templates.
//
// error code that may be returned:
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
func (c *Client) CreateAuditRuleTemplateWithContext(ctx context.Context, request *CreateAuditRuleTemplateRequest) (response *CreateAuditRuleTemplateResponse, err error) {
    if request == nil {
        request = NewCreateAuditRuleTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "CreateAuditRuleTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAuditRuleTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAuditRuleTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBackupRequest() (request *CreateBackupRequest) {
    request = &CreateBackupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "CreateBackup")
    
    
    return
}

func NewCreateBackupResponse() (response *CreateBackupResponse) {
    response = &CreateBackupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateBackup
// This API is used to create a manual backup for a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateBackup(request *CreateBackupRequest) (response *CreateBackupResponse, err error) {
    return c.CreateBackupWithContext(context.Background(), request)
}

// CreateBackup
// This API is used to create a manual backup for a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateBackupWithContext(ctx context.Context, request *CreateBackupRequest) (response *CreateBackupResponse, err error) {
    if request == nil {
        request = NewCreateBackupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "CreateBackup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBackup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBackupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterDatabaseRequest() (request *CreateClusterDatabaseRequest) {
    request = &CreateClusterDatabaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "CreateClusterDatabase")
    
    
    return
}

func NewCreateClusterDatabaseResponse() (response *CreateClusterDatabaseResponse) {
    response = &CreateClusterDatabaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateClusterDatabase
// This API is used to create a database.
//
// error code that may be returned:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateClusterDatabase(request *CreateClusterDatabaseRequest) (response *CreateClusterDatabaseResponse, err error) {
    return c.CreateClusterDatabaseWithContext(context.Background(), request)
}

// CreateClusterDatabase
// This API is used to create a database.
//
// error code that may be returned:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateClusterDatabaseWithContext(ctx context.Context, request *CreateClusterDatabaseRequest) (response *CreateClusterDatabaseResponse, err error) {
    if request == nil {
        request = NewCreateClusterDatabaseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "CreateClusterDatabase")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateClusterDatabase require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateClusterDatabaseResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClustersRequest() (request *CreateClustersRequest) {
    request = &CreateClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "CreateClusters")
    
    
    return
}

func NewCreateClustersResponse() (response *CreateClustersResponse) {
    response = &CreateClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateClusters
// This API is used to purchase new clusters.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEORDER = "FailedOperation.CreateOrder"
//  FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_QUERYSPECBYSPECCODEERROR = "FailedOperation.QuerySpecBySpecCodeError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  INVALIDPARAMETERVALUE_PROJECTIDNOTFOUND = "InvalidParameterValue.ProjectIdNotFound"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSUFFICIENTBALANCEERROR = "OperationDenied.InsufficientBalanceError"
//  OPERATIONDENIED_USERNOTAUTHENTICATEDERROR = "OperationDenied.UserNotAuthenticatedError"
//  UNAUTHORIZEDOPERATION_NOTREALNAMEACCOUNT = "UnauthorizedOperation.NotRealNameAccount"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateClusters(request *CreateClustersRequest) (response *CreateClustersResponse, err error) {
    return c.CreateClustersWithContext(context.Background(), request)
}

// CreateClusters
// This API is used to purchase new clusters.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEORDER = "FailedOperation.CreateOrder"
//  FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_QUERYSPECBYSPECCODEERROR = "FailedOperation.QuerySpecBySpecCodeError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  INVALIDPARAMETERVALUE_PROJECTIDNOTFOUND = "InvalidParameterValue.ProjectIdNotFound"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSUFFICIENTBALANCEERROR = "OperationDenied.InsufficientBalanceError"
//  OPERATIONDENIED_USERNOTAUTHENTICATEDERROR = "OperationDenied.UserNotAuthenticatedError"
//  UNAUTHORIZEDOPERATION_NOTREALNAMEACCOUNT = "UnauthorizedOperation.NotRealNameAccount"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateClustersWithContext(ctx context.Context, request *CreateClustersRequest) (response *CreateClustersResponse, err error) {
    if request == nil {
        request = NewCreateClustersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "CreateClusters")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateClusters require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateClustersResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIntegrateClusterRequest() (request *CreateIntegrateClusterRequest) {
    request = &CreateIntegrateClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "CreateIntegrateCluster")
    
    
    return
}

func NewCreateIntegrateClusterResponse() (response *CreateIntegrateClusterResponse) {
    response = &CreateIntegrateClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateIntegrateCluster
// This API is used to create a newly purchased cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEORDER = "FailedOperation.CreateOrder"
//  FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_QUERYSPECBYSPECCODEERROR = "FailedOperation.QuerySpecBySpecCodeError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  INVALIDPARAMETERVALUE_PROJECTIDNOTFOUND = "InvalidParameterValue.ProjectIdNotFound"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSUFFICIENTBALANCEERROR = "OperationDenied.InsufficientBalanceError"
//  OPERATIONDENIED_USERNOTAUTHENTICATEDERROR = "OperationDenied.UserNotAuthenticatedError"
//  UNAUTHORIZEDOPERATION_NOTREALNAMEACCOUNT = "UnauthorizedOperation.NotRealNameAccount"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateIntegrateCluster(request *CreateIntegrateClusterRequest) (response *CreateIntegrateClusterResponse, err error) {
    return c.CreateIntegrateClusterWithContext(context.Background(), request)
}

// CreateIntegrateCluster
// This API is used to create a newly purchased cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEORDER = "FailedOperation.CreateOrder"
//  FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_QUERYSPECBYSPECCODEERROR = "FailedOperation.QuerySpecBySpecCodeError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  INVALIDPARAMETERVALUE_PROJECTIDNOTFOUND = "InvalidParameterValue.ProjectIdNotFound"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSUFFICIENTBALANCEERROR = "OperationDenied.InsufficientBalanceError"
//  OPERATIONDENIED_USERNOTAUTHENTICATEDERROR = "OperationDenied.UserNotAuthenticatedError"
//  UNAUTHORIZEDOPERATION_NOTREALNAMEACCOUNT = "UnauthorizedOperation.NotRealNameAccount"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateIntegrateClusterWithContext(ctx context.Context, request *CreateIntegrateClusterRequest) (response *CreateIntegrateClusterResponse, err error) {
    if request == nil {
        request = NewCreateIntegrateClusterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "CreateIntegrateCluster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIntegrateCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateIntegrateClusterResponse()
    err = c.Send(request, response)
    return
}

func NewCreateParamTemplateRequest() (request *CreateParamTemplateRequest) {
    request = &CreateParamTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "CreateParamTemplate")
    
    
    return
}

func NewCreateParamTemplateResponse() (response *CreateParamTemplateResponse) {
    response = &CreateParamTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateParamTemplate
// This API is used to create parameter templates.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateParamTemplate(request *CreateParamTemplateRequest) (response *CreateParamTemplateResponse, err error) {
    return c.CreateParamTemplateWithContext(context.Background(), request)
}

// CreateParamTemplate
// This API is used to create parameter templates.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateParamTemplateWithContext(ctx context.Context, request *CreateParamTemplateRequest) (response *CreateParamTemplateResponse, err error) {
    if request == nil {
        request = NewCreateParamTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "CreateParamTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateParamTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateParamTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProxyRequest() (request *CreateProxyRequest) {
    request = &CreateProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "CreateProxy")
    
    
    return
}

func NewCreateProxyResponse() (response *CreateProxyResponse) {
    response = &CreateProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateProxy
// This API is used to enable the database proxy of a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateProxy(request *CreateProxyRequest) (response *CreateProxyResponse, err error) {
    return c.CreateProxyWithContext(context.Background(), request)
}

// CreateProxy
// This API is used to enable the database proxy of a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateProxyWithContext(ctx context.Context, request *CreateProxyRequest) (response *CreateProxyResponse, err error) {
    if request == nil {
        request = NewCreateProxyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "CreateProxy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateProxyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProxyEndPointRequest() (request *CreateProxyEndPointRequest) {
    request = &CreateProxyEndPointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "CreateProxyEndPoint")
    
    
    return
}

func NewCreateProxyEndPointResponse() (response *CreateProxyEndPointResponse) {
    response = &CreateProxyEndPointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateProxyEndPoint
// This API is used to create a database proxy connection point.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateProxyEndPoint(request *CreateProxyEndPointRequest) (response *CreateProxyEndPointResponse, err error) {
    return c.CreateProxyEndPointWithContext(context.Background(), request)
}

// CreateProxyEndPoint
// This API is used to create a database proxy connection point.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateProxyEndPointWithContext(ctx context.Context, request *CreateProxyEndPointRequest) (response *CreateProxyEndPointResponse, err error) {
    if request == nil {
        request = NewCreateProxyEndPointRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "CreateProxyEndPoint")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateProxyEndPoint require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateProxyEndPointResponse()
    err = c.Send(request, response)
    return
}

func NewCreateResourcePackageRequest() (request *CreateResourcePackageRequest) {
    request = &CreateResourcePackageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "CreateResourcePackage")
    
    
    return
}

func NewCreateResourcePackageResponse() (response *CreateResourcePackageResponse) {
    response = &CreateResourcePackageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateResourcePackage
// This API is used to purchase new resource packets.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATESOURCEPACKAGEERROR = "FailedOperation.CreateSourcePackageError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateResourcePackage(request *CreateResourcePackageRequest) (response *CreateResourcePackageResponse, err error) {
    return c.CreateResourcePackageWithContext(context.Background(), request)
}

// CreateResourcePackage
// This API is used to purchase new resource packets.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATESOURCEPACKAGEERROR = "FailedOperation.CreateSourcePackageError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateResourcePackageWithContext(ctx context.Context, request *CreateResourcePackageRequest) (response *CreateResourcePackageResponse, err error) {
    if request == nil {
        request = NewCreateResourcePackageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "CreateResourcePackage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateResourcePackage require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateResourcePackageResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAccountsRequest() (request *DeleteAccountsRequest) {
    request = &DeleteAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DeleteAccounts")
    
    
    return
}

func NewDeleteAccountsResponse() (response *DeleteAccountsResponse) {
    response = &DeleteAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAccounts
// This API is used to delete user accounts.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteAccounts(request *DeleteAccountsRequest) (response *DeleteAccountsResponse, err error) {
    return c.DeleteAccountsWithContext(context.Background(), request)
}

// DeleteAccounts
// This API is used to delete user accounts.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteAccountsWithContext(ctx context.Context, request *DeleteAccountsRequest) (response *DeleteAccountsResponse, err error) {
    if request == nil {
        request = NewDeleteAccountsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DeleteAccounts")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAccounts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAuditRuleTemplatesRequest() (request *DeleteAuditRuleTemplatesRequest) {
    request = &DeleteAuditRuleTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DeleteAuditRuleTemplates")
    
    
    return
}

func NewDeleteAuditRuleTemplatesResponse() (response *DeleteAuditRuleTemplatesResponse) {
    response = &DeleteAuditRuleTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAuditRuleTemplates
// This API is used to delete audit rule templates.
//
// error code that may be returned:
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
func (c *Client) DeleteAuditRuleTemplates(request *DeleteAuditRuleTemplatesRequest) (response *DeleteAuditRuleTemplatesResponse, err error) {
    return c.DeleteAuditRuleTemplatesWithContext(context.Background(), request)
}

// DeleteAuditRuleTemplates
// This API is used to delete audit rule templates.
//
// error code that may be returned:
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
func (c *Client) DeleteAuditRuleTemplatesWithContext(ctx context.Context, request *DeleteAuditRuleTemplatesRequest) (response *DeleteAuditRuleTemplatesResponse, err error) {
    if request == nil {
        request = NewDeleteAuditRuleTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DeleteAuditRuleTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAuditRuleTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAuditRuleTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBackupRequest() (request *DeleteBackupRequest) {
    request = &DeleteBackupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DeleteBackup")
    
    
    return
}

func NewDeleteBackupResponse() (response *DeleteBackupResponse) {
    response = &DeleteBackupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteBackup
// This API is used to delete manual backups for a cluster. Automatic backups cannot be deleted.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteBackup(request *DeleteBackupRequest) (response *DeleteBackupResponse, err error) {
    return c.DeleteBackupWithContext(context.Background(), request)
}

// DeleteBackup
// This API is used to delete manual backups for a cluster. Automatic backups cannot be deleted.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteBackupWithContext(ctx context.Context, request *DeleteBackupRequest) (response *DeleteBackupResponse, err error) {
    if request == nil {
        request = NewDeleteBackupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DeleteBackup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBackup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBackupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterDatabaseRequest() (request *DeleteClusterDatabaseRequest) {
    request = &DeleteClusterDatabaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DeleteClusterDatabase")
    
    
    return
}

func NewDeleteClusterDatabaseResponse() (response *DeleteClusterDatabaseResponse) {
    response = &DeleteClusterDatabaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteClusterDatabase
// This interface is used to delete a database.
//
// error code that may be returned:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteClusterDatabase(request *DeleteClusterDatabaseRequest) (response *DeleteClusterDatabaseResponse, err error) {
    return c.DeleteClusterDatabaseWithContext(context.Background(), request)
}

// DeleteClusterDatabase
// This interface is used to delete a database.
//
// error code that may be returned:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteClusterDatabaseWithContext(ctx context.Context, request *DeleteClusterDatabaseRequest) (response *DeleteClusterDatabaseResponse, err error) {
    if request == nil {
        request = NewDeleteClusterDatabaseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DeleteClusterDatabase")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteClusterDatabase require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteClusterDatabaseResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteParamTemplateRequest() (request *DeleteParamTemplateRequest) {
    request = &DeleteParamTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DeleteParamTemplate")
    
    
    return
}

func NewDeleteParamTemplateResponse() (response *DeleteParamTemplateResponse) {
    response = &DeleteParamTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteParamTemplate
// This API is used to delete a parameter template.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
func (c *Client) DeleteParamTemplate(request *DeleteParamTemplateRequest) (response *DeleteParamTemplateResponse, err error) {
    return c.DeleteParamTemplateWithContext(context.Background(), request)
}

// DeleteParamTemplate
// This API is used to delete a parameter template.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
func (c *Client) DeleteParamTemplateWithContext(ctx context.Context, request *DeleteParamTemplateRequest) (response *DeleteParamTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteParamTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DeleteParamTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteParamTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteParamTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountPrivilegesRequest() (request *DescribeAccountPrivilegesRequest) {
    request = &DescribeAccountPrivilegesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeAccountPrivileges")
    
    
    return
}

func NewDescribeAccountPrivilegesResponse() (response *DescribeAccountPrivilegesResponse) {
    response = &DescribeAccountPrivilegesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccountPrivileges
// This API is used to query account privileges.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeAccountPrivileges(request *DescribeAccountPrivilegesRequest) (response *DescribeAccountPrivilegesResponse, err error) {
    return c.DescribeAccountPrivilegesWithContext(context.Background(), request)
}

// DescribeAccountPrivileges
// This API is used to query account privileges.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeAccountPrivilegesWithContext(ctx context.Context, request *DescribeAccountPrivilegesRequest) (response *DescribeAccountPrivilegesResponse, err error) {
    if request == nil {
        request = NewDescribeAccountPrivilegesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeAccountPrivileges")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccountPrivileges require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccountPrivilegesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountsRequest() (request *DescribeAccountsRequest) {
    request = &DescribeAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeAccounts")
    
    
    return
}

func NewDescribeAccountsResponse() (response *DescribeAccountsResponse) {
    response = &DescribeAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccounts
// This API is used to query the database account list.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeAccounts(request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
    return c.DescribeAccountsWithContext(context.Background(), request)
}

// DescribeAccounts
// This API is used to query the database account list.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeAccountsWithContext(ctx context.Context, request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeAccountsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeAccounts")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccounts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuditRuleTemplatesRequest() (request *DescribeAuditRuleTemplatesRequest) {
    request = &DescribeAuditRuleTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeAuditRuleTemplates")
    
    
    return
}

func NewDescribeAuditRuleTemplatesResponse() (response *DescribeAuditRuleTemplatesResponse) {
    response = &DescribeAuditRuleTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAuditRuleTemplates
// This API is used to query audit rule template information.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeAuditRuleTemplates(request *DescribeAuditRuleTemplatesRequest) (response *DescribeAuditRuleTemplatesResponse, err error) {
    return c.DescribeAuditRuleTemplatesWithContext(context.Background(), request)
}

// DescribeAuditRuleTemplates
// This API is used to query audit rule template information.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeAuditRuleTemplatesWithContext(ctx context.Context, request *DescribeAuditRuleTemplatesRequest) (response *DescribeAuditRuleTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeAuditRuleTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeAuditRuleTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuditRuleTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuditRuleTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuditRuleWithInstanceIdsRequest() (request *DescribeAuditRuleWithInstanceIdsRequest) {
    request = &DescribeAuditRuleWithInstanceIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeAuditRuleWithInstanceIds")
    
    
    return
}

func NewDescribeAuditRuleWithInstanceIdsResponse() (response *DescribeAuditRuleWithInstanceIdsResponse) {
    response = &DescribeAuditRuleWithInstanceIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAuditRuleWithInstanceIds
// This API is used to obtain the audit rules of the instance.
//
// error code that may be returned:
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
func (c *Client) DescribeAuditRuleWithInstanceIds(request *DescribeAuditRuleWithInstanceIdsRequest) (response *DescribeAuditRuleWithInstanceIdsResponse, err error) {
    return c.DescribeAuditRuleWithInstanceIdsWithContext(context.Background(), request)
}

// DescribeAuditRuleWithInstanceIds
// This API is used to obtain the audit rules of the instance.
//
// error code that may be returned:
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
func (c *Client) DescribeAuditRuleWithInstanceIdsWithContext(ctx context.Context, request *DescribeAuditRuleWithInstanceIdsRequest) (response *DescribeAuditRuleWithInstanceIdsResponse, err error) {
    if request == nil {
        request = NewDescribeAuditRuleWithInstanceIdsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeAuditRuleWithInstanceIds")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuditRuleWithInstanceIds require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuditRuleWithInstanceIdsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupConfigRequest() (request *DescribeBackupConfigRequest) {
    request = &DescribeBackupConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeBackupConfig")
    
    
    return
}

func NewDescribeBackupConfigResponse() (response *DescribeBackupConfigResponse) {
    response = &DescribeBackupConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupConfig
// This API is used to obtain the backup configuration information of a specified cluster, including the full backup time period and the backup file retention time.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_GETBACKUPSTRATEGYERROR = "FailedOperation.GetBackupStrategyError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupConfig(request *DescribeBackupConfigRequest) (response *DescribeBackupConfigResponse, err error) {
    return c.DescribeBackupConfigWithContext(context.Background(), request)
}

// DescribeBackupConfig
// This API is used to obtain the backup configuration information of a specified cluster, including the full backup time period and the backup file retention time.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_GETBACKUPSTRATEGYERROR = "FailedOperation.GetBackupStrategyError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupConfigWithContext(ctx context.Context, request *DescribeBackupConfigRequest) (response *DescribeBackupConfigResponse, err error) {
    if request == nil {
        request = NewDescribeBackupConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeBackupConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupDownloadRestrictionRequest() (request *DescribeBackupDownloadRestrictionRequest) {
    request = &DescribeBackupDownloadRestrictionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeBackupDownloadRestriction")
    
    
    return
}

func NewDescribeBackupDownloadRestrictionResponse() (response *DescribeBackupDownloadRestrictionResponse) {
    response = &DescribeBackupDownloadRestrictionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupDownloadRestriction
// This API is used to query the download source limit of the default backup configured by the user in the current region.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREGIONIDERROR = "InvalidParameterValue.InvalidRegionIdError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupDownloadRestriction(request *DescribeBackupDownloadRestrictionRequest) (response *DescribeBackupDownloadRestrictionResponse, err error) {
    return c.DescribeBackupDownloadRestrictionWithContext(context.Background(), request)
}

// DescribeBackupDownloadRestriction
// This API is used to query the download source limit of the default backup configured by the user in the current region.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREGIONIDERROR = "InvalidParameterValue.InvalidRegionIdError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupDownloadRestrictionWithContext(ctx context.Context, request *DescribeBackupDownloadRestrictionRequest) (response *DescribeBackupDownloadRestrictionResponse, err error) {
    if request == nil {
        request = NewDescribeBackupDownloadRestrictionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeBackupDownloadRestriction")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupDownloadRestriction require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupDownloadRestrictionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupDownloadUrlRequest() (request *DescribeBackupDownloadUrlRequest) {
    request = &DescribeBackupDownloadUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeBackupDownloadUrl")
    
    
    return
}

func NewDescribeBackupDownloadUrlResponse() (response *DescribeBackupDownloadUrlResponse) {
    response = &DescribeBackupDownloadUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupDownloadUrl
// This API is used to query the download link of cluster backup files.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupDownloadUrl(request *DescribeBackupDownloadUrlRequest) (response *DescribeBackupDownloadUrlResponse, err error) {
    return c.DescribeBackupDownloadUrlWithContext(context.Background(), request)
}

// DescribeBackupDownloadUrl
// This API is used to query the download link of cluster backup files.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupDownloadUrlWithContext(ctx context.Context, request *DescribeBackupDownloadUrlRequest) (response *DescribeBackupDownloadUrlResponse, err error) {
    if request == nil {
        request = NewDescribeBackupDownloadUrlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeBackupDownloadUrl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupDownloadUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupDownloadUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupDownloadUserRestrictionRequest() (request *DescribeBackupDownloadUserRestrictionRequest) {
    request = &DescribeBackupDownloadUserRestrictionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeBackupDownloadUserRestriction")
    
    
    return
}

func NewDescribeBackupDownloadUserRestrictionResponse() (response *DescribeBackupDownloadUserRestrictionResponse) {
    response = &DescribeBackupDownloadUserRestrictionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupDownloadUserRestriction
// This API is used to query the default backup download access restrictions of user-level settings in the current region.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREGIONIDERROR = "InvalidParameterValue.InvalidRegionIdError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupDownloadUserRestriction(request *DescribeBackupDownloadUserRestrictionRequest) (response *DescribeBackupDownloadUserRestrictionResponse, err error) {
    return c.DescribeBackupDownloadUserRestrictionWithContext(context.Background(), request)
}

// DescribeBackupDownloadUserRestriction
// This API is used to query the default backup download access restrictions of user-level settings in the current region.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREGIONIDERROR = "InvalidParameterValue.InvalidRegionIdError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupDownloadUserRestrictionWithContext(ctx context.Context, request *DescribeBackupDownloadUserRestrictionRequest) (response *DescribeBackupDownloadUserRestrictionResponse, err error) {
    if request == nil {
        request = NewDescribeBackupDownloadUserRestrictionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeBackupDownloadUserRestriction")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupDownloadUserRestriction require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupDownloadUserRestrictionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupListRequest() (request *DescribeBackupListRequest) {
    request = &DescribeBackupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeBackupList")
    
    
    return
}

func NewDescribeBackupListResponse() (response *DescribeBackupListResponse) {
    response = &DescribeBackupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupList
// This API is used to query the backup file list of a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_GETBACKUPSTRATEGYERROR = "FailedOperation.GetBackupStrategyError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupList(request *DescribeBackupListRequest) (response *DescribeBackupListResponse, err error) {
    return c.DescribeBackupListWithContext(context.Background(), request)
}

// DescribeBackupList
// This API is used to query the backup file list of a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_GETBACKUPSTRATEGYERROR = "FailedOperation.GetBackupStrategyError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupListWithContext(ctx context.Context, request *DescribeBackupListRequest) (response *DescribeBackupListResponse, err error) {
    if request == nil {
        request = NewDescribeBackupListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeBackupList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBinlogConfigRequest() (request *DescribeBinlogConfigRequest) {
    request = &DescribeBinlogConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeBinlogConfig")
    
    
    return
}

func NewDescribeBinlogConfigResponse() (response *DescribeBinlogConfigResponse) {
    response = &DescribeBinlogConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBinlogConfig
// This API is used to query binlog configurations.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
func (c *Client) DescribeBinlogConfig(request *DescribeBinlogConfigRequest) (response *DescribeBinlogConfigResponse, err error) {
    return c.DescribeBinlogConfigWithContext(context.Background(), request)
}

// DescribeBinlogConfig
// This API is used to query binlog configurations.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
func (c *Client) DescribeBinlogConfigWithContext(ctx context.Context, request *DescribeBinlogConfigRequest) (response *DescribeBinlogConfigResponse, err error) {
    if request == nil {
        request = NewDescribeBinlogConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeBinlogConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBinlogConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBinlogConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBinlogDownloadUrlRequest() (request *DescribeBinlogDownloadUrlRequest) {
    request = &DescribeBinlogDownloadUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeBinlogDownloadUrl")
    
    
    return
}

func NewDescribeBinlogDownloadUrlResponse() (response *DescribeBinlogDownloadUrlResponse) {
    response = &DescribeBinlogDownloadUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBinlogDownloadUrl
// This API is used to query the download address of Binlog.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBinlogDownloadUrl(request *DescribeBinlogDownloadUrlRequest) (response *DescribeBinlogDownloadUrlResponse, err error) {
    return c.DescribeBinlogDownloadUrlWithContext(context.Background(), request)
}

// DescribeBinlogDownloadUrl
// This API is used to query the download address of Binlog.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBinlogDownloadUrlWithContext(ctx context.Context, request *DescribeBinlogDownloadUrlRequest) (response *DescribeBinlogDownloadUrlResponse, err error) {
    if request == nil {
        request = NewDescribeBinlogDownloadUrlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeBinlogDownloadUrl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBinlogDownloadUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBinlogDownloadUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBinlogSaveDaysRequest() (request *DescribeBinlogSaveDaysRequest) {
    request = &DescribeBinlogSaveDaysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeBinlogSaveDays")
    
    
    return
}

func NewDescribeBinlogSaveDaysResponse() (response *DescribeBinlogSaveDaysResponse) {
    response = &DescribeBinlogSaveDaysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBinlogSaveDays
// This API is used to query the binlog retention period of a cluster in days.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_GETBACKUPSTRATEGYERROR = "FailedOperation.GetBackupStrategyError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBinlogSaveDays(request *DescribeBinlogSaveDaysRequest) (response *DescribeBinlogSaveDaysResponse, err error) {
    return c.DescribeBinlogSaveDaysWithContext(context.Background(), request)
}

// DescribeBinlogSaveDays
// This API is used to query the binlog retention period of a cluster in days.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_GETBACKUPSTRATEGYERROR = "FailedOperation.GetBackupStrategyError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBinlogSaveDaysWithContext(ctx context.Context, request *DescribeBinlogSaveDaysRequest) (response *DescribeBinlogSaveDaysResponse, err error) {
    if request == nil {
        request = NewDescribeBinlogSaveDaysRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeBinlogSaveDays")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBinlogSaveDays require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBinlogSaveDaysResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBinlogsRequest() (request *DescribeBinlogsRequest) {
    request = &DescribeBinlogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeBinlogs")
    
    
    return
}

func NewDescribeBinlogsResponse() (response *DescribeBinlogsResponse) {
    response = &DescribeBinlogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBinlogs
// This interface (DescribeBinlogs) queries the cluster binlog list.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBinlogs(request *DescribeBinlogsRequest) (response *DescribeBinlogsResponse, err error) {
    return c.DescribeBinlogsWithContext(context.Background(), request)
}

// DescribeBinlogs
// This interface (DescribeBinlogs) queries the cluster binlog list.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBinlogsWithContext(ctx context.Context, request *DescribeBinlogsRequest) (response *DescribeBinlogsResponse, err error) {
    if request == nil {
        request = NewDescribeBinlogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeBinlogs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBinlogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBinlogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterDatabaseTablesRequest() (request *DescribeClusterDatabaseTablesRequest) {
    request = &DescribeClusterDatabaseTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeClusterDatabaseTables")
    
    
    return
}

func NewDescribeClusterDatabaseTablesResponse() (response *DescribeClusterDatabaseTablesResponse) {
    response = &DescribeClusterDatabaseTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterDatabaseTables
// This API is used to access the table list.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterDatabaseTables(request *DescribeClusterDatabaseTablesRequest) (response *DescribeClusterDatabaseTablesResponse, err error) {
    return c.DescribeClusterDatabaseTablesWithContext(context.Background(), request)
}

// DescribeClusterDatabaseTables
// This API is used to access the table list.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterDatabaseTablesWithContext(ctx context.Context, request *DescribeClusterDatabaseTablesRequest) (response *DescribeClusterDatabaseTablesResponse, err error) {
    if request == nil {
        request = NewDescribeClusterDatabaseTablesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeClusterDatabaseTables")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterDatabaseTables require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterDatabaseTablesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterDetailRequest() (request *DescribeClusterDetailRequest) {
    request = &DescribeClusterDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeClusterDetail")
    
    
    return
}

func NewDescribeClusterDetailResponse() (response *DescribeClusterDetailResponse) {
    response = &DescribeClusterDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterDetail
// This API is used to display cluster details.
//
// error code that may be returned:
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  OPERATIONDENIED_VERSIONNOTSUPPORTERROR = "OperationDenied.VersionNotSupportError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterDetail(request *DescribeClusterDetailRequest) (response *DescribeClusterDetailResponse, err error) {
    return c.DescribeClusterDetailWithContext(context.Background(), request)
}

// DescribeClusterDetail
// This API is used to display cluster details.
//
// error code that may be returned:
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  OPERATIONDENIED_VERSIONNOTSUPPORTERROR = "OperationDenied.VersionNotSupportError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterDetailWithContext(ctx context.Context, request *DescribeClusterDetailRequest) (response *DescribeClusterDetailResponse, err error) {
    if request == nil {
        request = NewDescribeClusterDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeClusterDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterDetailDatabasesRequest() (request *DescribeClusterDetailDatabasesRequest) {
    request = &DescribeClusterDetailDatabasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeClusterDetailDatabases")
    
    
    return
}

func NewDescribeClusterDetailDatabasesResponse() (response *DescribeClusterDetailDatabasesResponse) {
    response = &DescribeClusterDetailDatabasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterDetailDatabases
// This API is used to query database list.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterDetailDatabases(request *DescribeClusterDetailDatabasesRequest) (response *DescribeClusterDetailDatabasesResponse, err error) {
    return c.DescribeClusterDetailDatabasesWithContext(context.Background(), request)
}

// DescribeClusterDetailDatabases
// This API is used to query database list.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterDetailDatabasesWithContext(ctx context.Context, request *DescribeClusterDetailDatabasesRequest) (response *DescribeClusterDetailDatabasesResponse, err error) {
    if request == nil {
        request = NewDescribeClusterDetailDatabasesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeClusterDetailDatabases")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterDetailDatabases require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterDetailDatabasesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterInstanceGrpsRequest() (request *DescribeClusterInstanceGrpsRequest) {
    request = &DescribeClusterInstanceGrpsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeClusterInstanceGrps")
    
    
    return
}

func NewDescribeClusterInstanceGrpsResponse() (response *DescribeClusterInstanceGrpsResponse) {
    response = &DescribeClusterInstanceGrpsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterInstanceGrps
// This API is used to query instance groups.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterInstanceGrps(request *DescribeClusterInstanceGrpsRequest) (response *DescribeClusterInstanceGrpsResponse, err error) {
    return c.DescribeClusterInstanceGrpsWithContext(context.Background(), request)
}

// DescribeClusterInstanceGrps
// This API is used to query instance groups.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterInstanceGrpsWithContext(ctx context.Context, request *DescribeClusterInstanceGrpsRequest) (response *DescribeClusterInstanceGrpsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterInstanceGrpsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeClusterInstanceGrps")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterInstanceGrps require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterInstanceGrpsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterParamsRequest() (request *DescribeClusterParamsRequest) {
    request = &DescribeClusterParamsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeClusterParams")
    
    
    return
}

func NewDescribeClusterParamsResponse() (response *DescribeClusterParamsResponse) {
    response = &DescribeClusterParamsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterParams
// This API is used to query cluster parameters.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterParams(request *DescribeClusterParamsRequest) (response *DescribeClusterParamsResponse, err error) {
    return c.DescribeClusterParamsWithContext(context.Background(), request)
}

// DescribeClusterParams
// This API is used to query cluster parameters.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterParamsWithContext(ctx context.Context, request *DescribeClusterParamsRequest) (response *DescribeClusterParamsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterParamsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeClusterParams")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterParams require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterParamsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterPasswordComplexityRequest() (request *DescribeClusterPasswordComplexityRequest) {
    request = &DescribeClusterPasswordComplexityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeClusterPasswordComplexity")
    
    
    return
}

func NewDescribeClusterPasswordComplexityResponse() (response *DescribeClusterPasswordComplexityResponse) {
    response = &DescribeClusterPasswordComplexityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterPasswordComplexity
// This API is used to view the cluster password complexity details.
//
// error code that may be returned:
//  FAILEDOPERATION_GETOSSINFOERROR = "FailedOperation.GetOssInfoError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterPasswordComplexity(request *DescribeClusterPasswordComplexityRequest) (response *DescribeClusterPasswordComplexityResponse, err error) {
    return c.DescribeClusterPasswordComplexityWithContext(context.Background(), request)
}

// DescribeClusterPasswordComplexity
// This API is used to view the cluster password complexity details.
//
// error code that may be returned:
//  FAILEDOPERATION_GETOSSINFOERROR = "FailedOperation.GetOssInfoError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterPasswordComplexityWithContext(ctx context.Context, request *DescribeClusterPasswordComplexityRequest) (response *DescribeClusterPasswordComplexityResponse, err error) {
    if request == nil {
        request = NewDescribeClusterPasswordComplexityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeClusterPasswordComplexity")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterPasswordComplexity require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterPasswordComplexityResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterReadOnlyRequest() (request *DescribeClusterReadOnlyRequest) {
    request = &DescribeClusterReadOnlyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeClusterReadOnly")
    
    
    return
}

func NewDescribeClusterReadOnlyResponse() (response *DescribeClusterReadOnlyResponse) {
    response = &DescribeClusterReadOnlyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterReadOnly
// This API is used to query the cluster read-only switch.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterReadOnly(request *DescribeClusterReadOnlyRequest) (response *DescribeClusterReadOnlyResponse, err error) {
    return c.DescribeClusterReadOnlyWithContext(context.Background(), request)
}

// DescribeClusterReadOnly
// This API is used to query the cluster read-only switch.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterReadOnlyWithContext(ctx context.Context, request *DescribeClusterReadOnlyRequest) (response *DescribeClusterReadOnlyResponse, err error) {
    if request == nil {
        request = NewDescribeClusterReadOnlyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeClusterReadOnly")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterReadOnly require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterReadOnlyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterTransparentEncryptInfoRequest() (request *DescribeClusterTransparentEncryptInfoRequest) {
    request = &DescribeClusterTransparentEncryptInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeClusterTransparentEncryptInfo")
    
    
    return
}

func NewDescribeClusterTransparentEncryptInfoResponse() (response *DescribeClusterTransparentEncryptInfoResponse) {
    response = &DescribeClusterTransparentEncryptInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterTransparentEncryptInfo
// This API is used to query cluster transparent encryption information.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterTransparentEncryptInfo(request *DescribeClusterTransparentEncryptInfoRequest) (response *DescribeClusterTransparentEncryptInfoResponse, err error) {
    return c.DescribeClusterTransparentEncryptInfoWithContext(context.Background(), request)
}

// DescribeClusterTransparentEncryptInfo
// This API is used to query cluster transparent encryption information.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterTransparentEncryptInfoWithContext(ctx context.Context, request *DescribeClusterTransparentEncryptInfoRequest) (response *DescribeClusterTransparentEncryptInfoResponse, err error) {
    if request == nil {
        request = NewDescribeClusterTransparentEncryptInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeClusterTransparentEncryptInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterTransparentEncryptInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterTransparentEncryptInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClustersRequest() (request *DescribeClustersRequest) {
    request = &DescribeClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeClusters")
    
    
    return
}

func NewDescribeClustersResponse() (response *DescribeClustersResponse) {
    response = &DescribeClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusters
// This API is used to describe clusters.
//
// error code that may be returned:
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMBOTHSETERROR = "InvalidParameterValue.ParamBothSetError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusters(request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
    return c.DescribeClustersWithContext(context.Background(), request)
}

// DescribeClusters
// This API is used to describe clusters.
//
// error code that may be returned:
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMBOTHSETERROR = "InvalidParameterValue.ParamBothSetError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClustersWithContext(ctx context.Context, request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
    if request == nil {
        request = NewDescribeClustersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeClusters")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusters require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClustersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBSecurityGroupsRequest() (request *DescribeDBSecurityGroupsRequest) {
    request = &DescribeDBSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeDBSecurityGroups")
    
    
    return
}

func NewDescribeDBSecurityGroupsResponse() (response *DescribeDBSecurityGroupsResponse) {
    response = &DescribeDBSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBSecurityGroups
// This API is used to query instance security group information.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBSecurityGroups(request *DescribeDBSecurityGroupsRequest) (response *DescribeDBSecurityGroupsResponse, err error) {
    return c.DescribeDBSecurityGroupsWithContext(context.Background(), request)
}

// DescribeDBSecurityGroups
// This API is used to query instance security group information.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBSecurityGroupsWithContext(ctx context.Context, request *DescribeDBSecurityGroupsRequest) (response *DescribeDBSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeDBSecurityGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeDBSecurityGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlowRequest() (request *DescribeFlowRequest) {
    request = &DescribeFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeFlow")
    
    
    return
}

func NewDescribeFlowResponse() (response *DescribeFlowResponse) {
    response = &DescribeFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFlow
// This API is used to query task flow information.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWNOTFOUNDERROR = "FailedOperation.FlowNotFoundError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeFlow(request *DescribeFlowRequest) (response *DescribeFlowResponse, err error) {
    return c.DescribeFlowWithContext(context.Background(), request)
}

// DescribeFlow
// This API is used to query task flow information.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWNOTFOUNDERROR = "FailedOperation.FlowNotFoundError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeFlowWithContext(ctx context.Context, request *DescribeFlowRequest) (response *DescribeFlowResponse, err error) {
    if request == nil {
        request = NewDescribeFlowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeFlow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFlowResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceDetailRequest() (request *DescribeInstanceDetailRequest) {
    request = &DescribeInstanceDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeInstanceDetail")
    
    
    return
}

func NewDescribeInstanceDetailResponse() (response *DescribeInstanceDetailResponse) {
    response = &DescribeInstanceDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceDetail
// This API is used to query instance details.
//
// error code that may be returned:
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstanceDetail(request *DescribeInstanceDetailRequest) (response *DescribeInstanceDetailResponse, err error) {
    return c.DescribeInstanceDetailWithContext(context.Background(), request)
}

// DescribeInstanceDetail
// This API is used to query instance details.
//
// error code that may be returned:
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstanceDetailWithContext(ctx context.Context, request *DescribeInstanceDetailRequest) (response *DescribeInstanceDetailResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeInstanceDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceErrorLogsRequest() (request *DescribeInstanceErrorLogsRequest) {
    request = &DescribeInstanceErrorLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeInstanceErrorLogs")
    
    
    return
}

func NewDescribeInstanceErrorLogsResponse() (response *DescribeInstanceErrorLogsResponse) {
    response = &DescribeInstanceErrorLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceErrorLogs
// This API is used to query the list of instance error logs.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_VALUENOTFOUND = "InvalidParameterValue.ValueNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstanceErrorLogs(request *DescribeInstanceErrorLogsRequest) (response *DescribeInstanceErrorLogsResponse, err error) {
    return c.DescribeInstanceErrorLogsWithContext(context.Background(), request)
}

// DescribeInstanceErrorLogs
// This API is used to query the list of instance error logs.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_VALUENOTFOUND = "InvalidParameterValue.ValueNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstanceErrorLogsWithContext(ctx context.Context, request *DescribeInstanceErrorLogsRequest) (response *DescribeInstanceErrorLogsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceErrorLogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeInstanceErrorLogs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceErrorLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceErrorLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceParamsRequest() (request *DescribeInstanceParamsRequest) {
    request = &DescribeInstanceParamsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeInstanceParams")
    
    
    return
}

func NewDescribeInstanceParamsResponse() (response *DescribeInstanceParamsResponse) {
    response = &DescribeInstanceParamsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceParams
// This API is used to query the instance parameter list.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstanceParams(request *DescribeInstanceParamsRequest) (response *DescribeInstanceParamsResponse, err error) {
    return c.DescribeInstanceParamsWithContext(context.Background(), request)
}

// DescribeInstanceParams
// This API is used to query the instance parameter list.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstanceParamsWithContext(ctx context.Context, request *DescribeInstanceParamsRequest) (response *DescribeInstanceParamsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceParamsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeInstanceParams")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceParams require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceParamsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceSlowQueriesRequest() (request *DescribeInstanceSlowQueriesRequest) {
    request = &DescribeInstanceSlowQueriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeInstanceSlowQueries")
    
    
    return
}

func NewDescribeInstanceSlowQueriesResponse() (response *DescribeInstanceSlowQueriesResponse) {
    response = &DescribeInstanceSlowQueriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceSlowQueries
// This API is used to query the slow query logs of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_LENGTHOVERLIMIT = "OperationDenied.LengthOverLimit"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstanceSlowQueries(request *DescribeInstanceSlowQueriesRequest) (response *DescribeInstanceSlowQueriesResponse, err error) {
    return c.DescribeInstanceSlowQueriesWithContext(context.Background(), request)
}

// DescribeInstanceSlowQueries
// This API is used to query the slow query logs of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_LENGTHOVERLIMIT = "OperationDenied.LengthOverLimit"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstanceSlowQueriesWithContext(ctx context.Context, request *DescribeInstanceSlowQueriesRequest) (response *DescribeInstanceSlowQueriesResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceSlowQueriesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeInstanceSlowQueries")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceSlowQueries require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceSlowQueriesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceSpecsRequest() (request *DescribeInstanceSpecsRequest) {
    request = &DescribeInstanceSpecsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeInstanceSpecs")
    
    
    return
}

func NewDescribeInstanceSpecsResponse() (response *DescribeInstanceSpecsResponse) {
    response = &DescribeInstanceSpecsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceSpecs
// This interface (DescribeInstanceSpecs) is used to query the instance specifications available for purchase on the query purchase page.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstanceSpecs(request *DescribeInstanceSpecsRequest) (response *DescribeInstanceSpecsResponse, err error) {
    return c.DescribeInstanceSpecsWithContext(context.Background(), request)
}

// DescribeInstanceSpecs
// This interface (DescribeInstanceSpecs) is used to query the instance specifications available for purchase on the query purchase page.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstanceSpecsWithContext(ctx context.Context, request *DescribeInstanceSpecsRequest) (response *DescribeInstanceSpecsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceSpecsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeInstanceSpecs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceSpecs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceSpecsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeInstances")
    
    
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstances
// This API is used to query the list of instances.
//
// error code that may be returned:
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMBOTHSETERROR = "InvalidParameterValue.ParamBothSetError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    return c.DescribeInstancesWithContext(context.Background(), request)
}

// DescribeInstances
// This API is used to query the list of instances.
//
// error code that may be returned:
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMBOTHSETERROR = "InvalidParameterValue.ParamBothSetError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstancesWithContext(ctx context.Context, request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesWithinSameClusterRequest() (request *DescribeInstancesWithinSameClusterRequest) {
    request = &DescribeInstancesWithinSameClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeInstancesWithinSameCluster")
    
    
    return
}

func NewDescribeInstancesWithinSameClusterResponse() (response *DescribeInstancesWithinSameClusterResponse) {
    response = &DescribeInstancesWithinSameClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstancesWithinSameCluster
// This API is used to query the instance list under the same cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstancesWithinSameCluster(request *DescribeInstancesWithinSameClusterRequest) (response *DescribeInstancesWithinSameClusterResponse, err error) {
    return c.DescribeInstancesWithinSameClusterWithContext(context.Background(), request)
}

// DescribeInstancesWithinSameCluster
// This API is used to query the instance list under the same cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstancesWithinSameClusterWithContext(ctx context.Context, request *DescribeInstancesWithinSameClusterRequest) (response *DescribeInstancesWithinSameClusterResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesWithinSameClusterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeInstancesWithinSameCluster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstancesWithinSameCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesWithinSameClusterResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIntegrateTaskRequest() (request *DescribeIntegrateTaskRequest) {
    request = &DescribeIntegrateTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeIntegrateTask")
    
    
    return
}

func NewDescribeIntegrateTaskResponse() (response *DescribeIntegrateTaskResponse) {
    response = &DescribeIntegrateTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIntegrateTask
// This API is used to query cluster tasks.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCENOTFOUND_RESOURCEERROR = "ResourceNotFound.ResourceError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeIntegrateTask(request *DescribeIntegrateTaskRequest) (response *DescribeIntegrateTaskResponse, err error) {
    return c.DescribeIntegrateTaskWithContext(context.Background(), request)
}

// DescribeIntegrateTask
// This API is used to query cluster tasks.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCENOTFOUND_RESOURCEERROR = "ResourceNotFound.ResourceError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeIntegrateTaskWithContext(ctx context.Context, request *DescribeIntegrateTaskRequest) (response *DescribeIntegrateTaskResponse, err error) {
    if request == nil {
        request = NewDescribeIntegrateTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeIntegrateTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIntegrateTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIntegrateTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIsolatedInstancesRequest() (request *DescribeIsolatedInstancesRequest) {
    request = &DescribeIsolatedInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeIsolatedInstances")
    
    
    return
}

func NewDescribeIsolatedInstancesResponse() (response *DescribeIsolatedInstancesResponse) {
    response = &DescribeIsolatedInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIsolatedInstances
// This interface is used for querying the recycle bin instance list.
//
// error code that may be returned:
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMBOTHSETERROR = "InvalidParameterValue.ParamBothSetError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeIsolatedInstances(request *DescribeIsolatedInstancesRequest) (response *DescribeIsolatedInstancesResponse, err error) {
    return c.DescribeIsolatedInstancesWithContext(context.Background(), request)
}

// DescribeIsolatedInstances
// This interface is used for querying the recycle bin instance list.
//
// error code that may be returned:
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMBOTHSETERROR = "InvalidParameterValue.ParamBothSetError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeIsolatedInstancesWithContext(ctx context.Context, request *DescribeIsolatedInstancesRequest) (response *DescribeIsolatedInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeIsolatedInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeIsolatedInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIsolatedInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIsolatedInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMaintainPeriodRequest() (request *DescribeMaintainPeriodRequest) {
    request = &DescribeMaintainPeriodRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeMaintainPeriod")
    
    
    return
}

func NewDescribeMaintainPeriodResponse() (response *DescribeMaintainPeriodResponse) {
    response = &DescribeMaintainPeriodResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMaintainPeriod
// This interface (DescribeMaintainPeriod) is used to query the instance maintenance window.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeMaintainPeriod(request *DescribeMaintainPeriodRequest) (response *DescribeMaintainPeriodResponse, err error) {
    return c.DescribeMaintainPeriodWithContext(context.Background(), request)
}

// DescribeMaintainPeriod
// This interface (DescribeMaintainPeriod) is used to query the instance maintenance window.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeMaintainPeriodWithContext(ctx context.Context, request *DescribeMaintainPeriodRequest) (response *DescribeMaintainPeriodResponse, err error) {
    if request == nil {
        request = NewDescribeMaintainPeriodRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeMaintainPeriod")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMaintainPeriod require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMaintainPeriodResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeParamTemplateDetailRequest() (request *DescribeParamTemplateDetailRequest) {
    request = &DescribeParamTemplateDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeParamTemplateDetail")
    
    
    return
}

func NewDescribeParamTemplateDetailResponse() (response *DescribeParamTemplateDetailResponse) {
    response = &DescribeParamTemplateDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeParamTemplateDetail
// This API is used to query user parameter template details.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeParamTemplateDetail(request *DescribeParamTemplateDetailRequest) (response *DescribeParamTemplateDetailResponse, err error) {
    return c.DescribeParamTemplateDetailWithContext(context.Background(), request)
}

// DescribeParamTemplateDetail
// This API is used to query user parameter template details.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeParamTemplateDetailWithContext(ctx context.Context, request *DescribeParamTemplateDetailRequest) (response *DescribeParamTemplateDetailResponse, err error) {
    if request == nil {
        request = NewDescribeParamTemplateDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeParamTemplateDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeParamTemplateDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeParamTemplateDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeParamTemplatesRequest() (request *DescribeParamTemplatesRequest) {
    request = &DescribeParamTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeParamTemplates")
    
    
    return
}

func NewDescribeParamTemplatesResponse() (response *DescribeParamTemplatesResponse) {
    response = &DescribeParamTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeParamTemplates
// This API is used to query all parameter template information under the user-specified product.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeParamTemplates(request *DescribeParamTemplatesRequest) (response *DescribeParamTemplatesResponse, err error) {
    return c.DescribeParamTemplatesWithContext(context.Background(), request)
}

// DescribeParamTemplates
// This API is used to query all parameter template information under the user-specified product.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeParamTemplatesWithContext(ctx context.Context, request *DescribeParamTemplatesRequest) (response *DescribeParamTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeParamTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeParamTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeParamTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeParamTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectSecurityGroupsRequest() (request *DescribeProjectSecurityGroupsRequest) {
    request = &DescribeProjectSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeProjectSecurityGroups")
    
    
    return
}

func NewDescribeProjectSecurityGroupsResponse() (response *DescribeProjectSecurityGroupsResponse) {
    response = &DescribeProjectSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProjectSecurityGroups
// This API is used to query project security group information.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeProjectSecurityGroups(request *DescribeProjectSecurityGroupsRequest) (response *DescribeProjectSecurityGroupsResponse, err error) {
    return c.DescribeProjectSecurityGroupsWithContext(context.Background(), request)
}

// DescribeProjectSecurityGroups
// This API is used to query project security group information.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeProjectSecurityGroupsWithContext(ctx context.Context, request *DescribeProjectSecurityGroupsRequest) (response *DescribeProjectSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeProjectSecurityGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeProjectSecurityGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProjectSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProjectSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProxiesRequest() (request *DescribeProxiesRequest) {
    request = &DescribeProxiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeProxies")
    
    
    return
}

func NewDescribeProxiesResponse() (response *DescribeProxiesResponse) {
    response = &DescribeProxiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProxies
// This API is used to query agent list.
//
// error code that may be returned:
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMBOTHSETERROR = "InvalidParameterValue.ParamBothSetError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeProxies(request *DescribeProxiesRequest) (response *DescribeProxiesResponse, err error) {
    return c.DescribeProxiesWithContext(context.Background(), request)
}

// DescribeProxies
// This API is used to query agent list.
//
// error code that may be returned:
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMBOTHSETERROR = "InvalidParameterValue.ParamBothSetError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeProxiesWithContext(ctx context.Context, request *DescribeProxiesRequest) (response *DescribeProxiesResponse, err error) {
    if request == nil {
        request = NewDescribeProxiesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeProxies")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProxies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProxiesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProxyNodesRequest() (request *DescribeProxyNodesRequest) {
    request = &DescribeProxyNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeProxyNodes")
    
    
    return
}

func NewDescribeProxyNodesResponse() (response *DescribeProxyNodesResponse) {
    response = &DescribeProxyNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProxyNodes
// This API is used to query the list of proxy nodes.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeProxyNodes(request *DescribeProxyNodesRequest) (response *DescribeProxyNodesResponse, err error) {
    return c.DescribeProxyNodesWithContext(context.Background(), request)
}

// DescribeProxyNodes
// This API is used to query the list of proxy nodes.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeProxyNodesWithContext(ctx context.Context, request *DescribeProxyNodesRequest) (response *DescribeProxyNodesResponse, err error) {
    if request == nil {
        request = NewDescribeProxyNodesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeProxyNodes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProxyNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProxyNodesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProxySpecsRequest() (request *DescribeProxySpecsRequest) {
    request = &DescribeProxySpecsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeProxySpecs")
    
    
    return
}

func NewDescribeProxySpecsResponse() (response *DescribeProxySpecsResponse) {
    response = &DescribeProxySpecsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProxySpecs
// This API is used to query database proxy specifications.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
func (c *Client) DescribeProxySpecs(request *DescribeProxySpecsRequest) (response *DescribeProxySpecsResponse, err error) {
    return c.DescribeProxySpecsWithContext(context.Background(), request)
}

// DescribeProxySpecs
// This API is used to query database proxy specifications.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
func (c *Client) DescribeProxySpecsWithContext(ctx context.Context, request *DescribeProxySpecsRequest) (response *DescribeProxySpecsResponse, err error) {
    if request == nil {
        request = NewDescribeProxySpecsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeProxySpecs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProxySpecs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProxySpecsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourcePackageDetailRequest() (request *DescribeResourcePackageDetailRequest) {
    request = &DescribeResourcePackageDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeResourcePackageDetail")
    
    
    return
}

func NewDescribeResourcePackageDetailResponse() (response *DescribeResourcePackageDetailResponse) {
    response = &DescribeResourcePackageDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeResourcePackageDetail
// This API is used to query resource package usage details.
//
// error code that may be returned:
//  FAILEDOPERATION_QUERYSOURCEPACKAGEDETAILERROR = "FailedOperation.QuerySourcePackageDetailError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeResourcePackageDetail(request *DescribeResourcePackageDetailRequest) (response *DescribeResourcePackageDetailResponse, err error) {
    return c.DescribeResourcePackageDetailWithContext(context.Background(), request)
}

// DescribeResourcePackageDetail
// This API is used to query resource package usage details.
//
// error code that may be returned:
//  FAILEDOPERATION_QUERYSOURCEPACKAGEDETAILERROR = "FailedOperation.QuerySourcePackageDetailError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeResourcePackageDetailWithContext(ctx context.Context, request *DescribeResourcePackageDetailRequest) (response *DescribeResourcePackageDetailResponse, err error) {
    if request == nil {
        request = NewDescribeResourcePackageDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeResourcePackageDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourcePackageDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourcePackageDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourcePackageListRequest() (request *DescribeResourcePackageListRequest) {
    request = &DescribeResourcePackageListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeResourcePackageList")
    
    
    return
}

func NewDescribeResourcePackageListResponse() (response *DescribeResourcePackageListResponse) {
    response = &DescribeResourcePackageListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeResourcePackageList
// This API is used to query resource package list.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_QUERYSOURCEPACKAGEERROR = "FailedOperation.QuerySourcePackageError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeResourcePackageList(request *DescribeResourcePackageListRequest) (response *DescribeResourcePackageListResponse, err error) {
    return c.DescribeResourcePackageListWithContext(context.Background(), request)
}

// DescribeResourcePackageList
// This API is used to query resource package list.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_QUERYSOURCEPACKAGEERROR = "FailedOperation.QuerySourcePackageError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeResourcePackageListWithContext(ctx context.Context, request *DescribeResourcePackageListRequest) (response *DescribeResourcePackageListResponse, err error) {
    if request == nil {
        request = NewDescribeResourcePackageListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeResourcePackageList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourcePackageList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourcePackageListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourcePackageSaleSpecRequest() (request *DescribeResourcePackageSaleSpecRequest) {
    request = &DescribeResourcePackageSaleSpecRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeResourcePackageSaleSpec")
    
    
    return
}

func NewDescribeResourcePackageSaleSpecResponse() (response *DescribeResourcePackageSaleSpecResponse) {
    response = &DescribeResourcePackageSaleSpecResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeResourcePackageSaleSpec
// This API is used to query resource package specifications.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_UNSUPPORTSALESPECERROR = "OperationDenied.UnSupportSaleSpecError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeResourcePackageSaleSpec(request *DescribeResourcePackageSaleSpecRequest) (response *DescribeResourcePackageSaleSpecResponse, err error) {
    return c.DescribeResourcePackageSaleSpecWithContext(context.Background(), request)
}

// DescribeResourcePackageSaleSpec
// This API is used to query resource package specifications.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_UNSUPPORTSALESPECERROR = "OperationDenied.UnSupportSaleSpecError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeResourcePackageSaleSpecWithContext(ctx context.Context, request *DescribeResourcePackageSaleSpecRequest) (response *DescribeResourcePackageSaleSpecResponse, err error) {
    if request == nil {
        request = NewDescribeResourcePackageSaleSpecRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeResourcePackageSaleSpec")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourcePackageSaleSpec require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourcePackageSaleSpecResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourcesByDealNameRequest() (request *DescribeResourcesByDealNameRequest) {
    request = &DescribeResourcesByDealNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeResourcesByDealName")
    
    
    return
}

func NewDescribeResourcesByDealNameResponse() (response *DescribeResourcesByDealNameResponse) {
    response = &DescribeResourcesByDealNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeResourcesByDealName
// This interface (DescribeResourcesByDealName) is used to query order-associated instances.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_DEALNAMENOTFOUND = "InvalidParameterValue.DealNameNotFound"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_PARAMETEROUTRANGEERROR = "InvalidParameterValue.ParameterOutRangeError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeResourcesByDealName(request *DescribeResourcesByDealNameRequest) (response *DescribeResourcesByDealNameResponse, err error) {
    return c.DescribeResourcesByDealNameWithContext(context.Background(), request)
}

// DescribeResourcesByDealName
// This interface (DescribeResourcesByDealName) is used to query order-associated instances.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_DEALNAMENOTFOUND = "InvalidParameterValue.DealNameNotFound"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_PARAMETEROUTRANGEERROR = "InvalidParameterValue.ParameterOutRangeError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeResourcesByDealNameWithContext(ctx context.Context, request *DescribeResourcesByDealNameRequest) (response *DescribeResourcesByDealNameResponse, err error) {
    if request == nil {
        request = NewDescribeResourcesByDealNameRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeResourcesByDealName")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourcesByDealName require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourcesByDealNameResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRollbackTimeRangeRequest() (request *DescribeRollbackTimeRangeRequest) {
    request = &DescribeRollbackTimeRangeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeRollbackTimeRange")
    
    
    return
}

func NewDescribeRollbackTimeRangeResponse() (response *DescribeRollbackTimeRangeResponse) {
    response = &DescribeRollbackTimeRangeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRollbackTimeRange
// This API is used to query the rollback time range.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDREGIONIDERROR = "InvalidParameterValue.InvalidRegionIdError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeRollbackTimeRange(request *DescribeRollbackTimeRangeRequest) (response *DescribeRollbackTimeRangeResponse, err error) {
    return c.DescribeRollbackTimeRangeWithContext(context.Background(), request)
}

// DescribeRollbackTimeRange
// This API is used to query the rollback time range.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDREGIONIDERROR = "InvalidParameterValue.InvalidRegionIdError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeRollbackTimeRangeWithContext(ctx context.Context, request *DescribeRollbackTimeRangeRequest) (response *DescribeRollbackTimeRangeResponse, err error) {
    if request == nil {
        request = NewDescribeRollbackTimeRangeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeRollbackTimeRange")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRollbackTimeRange require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRollbackTimeRangeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServerlessInstanceSpecsRequest() (request *DescribeServerlessInstanceSpecsRequest) {
    request = &DescribeServerlessInstanceSpecsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeServerlessInstanceSpecs")
    
    
    return
}

func NewDescribeServerlessInstanceSpecsResponse() (response *DescribeServerlessInstanceSpecsResponse) {
    response = &DescribeServerlessInstanceSpecsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeServerlessInstanceSpecs
// This API is used to query available specifications of Serverless instances.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeServerlessInstanceSpecs(request *DescribeServerlessInstanceSpecsRequest) (response *DescribeServerlessInstanceSpecsResponse, err error) {
    return c.DescribeServerlessInstanceSpecsWithContext(context.Background(), request)
}

// DescribeServerlessInstanceSpecs
// This API is used to query available specifications of Serverless instances.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeServerlessInstanceSpecsWithContext(ctx context.Context, request *DescribeServerlessInstanceSpecsRequest) (response *DescribeServerlessInstanceSpecsResponse, err error) {
    if request == nil {
        request = NewDescribeServerlessInstanceSpecsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeServerlessInstanceSpecs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServerlessInstanceSpecs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServerlessInstanceSpecsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServerlessStrategyRequest() (request *DescribeServerlessStrategyRequest) {
    request = &DescribeServerlessStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeServerlessStrategy")
    
    
    return
}

func NewDescribeServerlessStrategyResponse() (response *DescribeServerlessStrategyResponse) {
    response = &DescribeServerlessStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeServerlessStrategy
// This API is used to query serverless policies.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_DBMODENOTSERVERLESSERROR = "InvalidParameterValue.DbModeNotServerlessError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
func (c *Client) DescribeServerlessStrategy(request *DescribeServerlessStrategyRequest) (response *DescribeServerlessStrategyResponse, err error) {
    return c.DescribeServerlessStrategyWithContext(context.Background(), request)
}

// DescribeServerlessStrategy
// This API is used to query serverless policies.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_DBMODENOTSERVERLESSERROR = "InvalidParameterValue.DbModeNotServerlessError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
func (c *Client) DescribeServerlessStrategyWithContext(ctx context.Context, request *DescribeServerlessStrategyRequest) (response *DescribeServerlessStrategyResponse, err error) {
    if request == nil {
        request = NewDescribeServerlessStrategyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeServerlessStrategy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServerlessStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServerlessStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlaveZonesRequest() (request *DescribeSlaveZonesRequest) {
    request = &DescribeSlaveZonesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeSlaveZones")
    
    
    return
}

func NewDescribeSlaveZonesResponse() (response *DescribeSlaveZonesResponse) {
    response = &DescribeSlaveZonesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSlaveZones
// This API is used to query from availability zones.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
func (c *Client) DescribeSlaveZones(request *DescribeSlaveZonesRequest) (response *DescribeSlaveZonesResponse, err error) {
    return c.DescribeSlaveZonesWithContext(context.Background(), request)
}

// DescribeSlaveZones
// This API is used to query from availability zones.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
func (c *Client) DescribeSlaveZonesWithContext(ctx context.Context, request *DescribeSlaveZonesRequest) (response *DescribeSlaveZonesResponse, err error) {
    if request == nil {
        request = NewDescribeSlaveZonesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeSlaveZones")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSlaveZones require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSlaveZonesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSupportProxyVersionRequest() (request *DescribeSupportProxyVersionRequest) {
    request = &DescribeSupportProxyVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeSupportProxyVersion")
    
    
    return
}

func NewDescribeSupportProxyVersionResponse() (response *DescribeSupportProxyVersionResponse) {
    response = &DescribeSupportProxyVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSupportProxyVersion
// This API is used to query supported database proxy versions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeSupportProxyVersion(request *DescribeSupportProxyVersionRequest) (response *DescribeSupportProxyVersionResponse, err error) {
    return c.DescribeSupportProxyVersionWithContext(context.Background(), request)
}

// DescribeSupportProxyVersion
// This API is used to query supported database proxy versions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeSupportProxyVersionWithContext(ctx context.Context, request *DescribeSupportProxyVersionRequest) (response *DescribeSupportProxyVersionResponse, err error) {
    if request == nil {
        request = NewDescribeSupportProxyVersionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeSupportProxyVersion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSupportProxyVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSupportProxyVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZonesRequest() (request *DescribeZonesRequest) {
    request = &DescribeZonesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeZones")
    
    
    return
}

func NewDescribeZonesResponse() (response *DescribeZonesResponse) {
    response = &DescribeZonesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeZones
// This API is used to query marketable regional availability zone information.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeZones(request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    return c.DescribeZonesWithContext(context.Background(), request)
}

// DescribeZones
// This API is used to query marketable regional availability zone information.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeZonesWithContext(ctx context.Context, request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    if request == nil {
        request = NewDescribeZonesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeZones")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZones require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeZonesResponse()
    err = c.Send(request, response)
    return
}

func NewExportInstanceErrorLogsRequest() (request *ExportInstanceErrorLogsRequest) {
    request = &ExportInstanceErrorLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ExportInstanceErrorLogs")
    
    
    return
}

func NewExportInstanceErrorLogsResponse() (response *ExportInstanceErrorLogsResponse) {
    response = &ExportInstanceErrorLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExportInstanceErrorLogs
// This API is used to export the error logs of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ExportInstanceErrorLogs(request *ExportInstanceErrorLogsRequest) (response *ExportInstanceErrorLogsResponse, err error) {
    return c.ExportInstanceErrorLogsWithContext(context.Background(), request)
}

// ExportInstanceErrorLogs
// This API is used to export the error logs of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ExportInstanceErrorLogsWithContext(ctx context.Context, request *ExportInstanceErrorLogsRequest) (response *ExportInstanceErrorLogsResponse, err error) {
    if request == nil {
        request = NewExportInstanceErrorLogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ExportInstanceErrorLogs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportInstanceErrorLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportInstanceErrorLogsResponse()
    err = c.Send(request, response)
    return
}

func NewExportInstanceSlowQueriesRequest() (request *ExportInstanceSlowQueriesRequest) {
    request = &ExportInstanceSlowQueriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ExportInstanceSlowQueries")
    
    
    return
}

func NewExportInstanceSlowQueriesResponse() (response *ExportInstanceSlowQueriesResponse) {
    response = &ExportInstanceSlowQueriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExportInstanceSlowQueries
// This API is used to export instance slow logs.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ExportInstanceSlowQueries(request *ExportInstanceSlowQueriesRequest) (response *ExportInstanceSlowQueriesResponse, err error) {
    return c.ExportInstanceSlowQueriesWithContext(context.Background(), request)
}

// ExportInstanceSlowQueries
// This API is used to export instance slow logs.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ExportInstanceSlowQueriesWithContext(ctx context.Context, request *ExportInstanceSlowQueriesRequest) (response *ExportInstanceSlowQueriesResponse, err error) {
    if request == nil {
        request = NewExportInstanceSlowQueriesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ExportInstanceSlowQueries")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportInstanceSlowQueries require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportInstanceSlowQueriesResponse()
    err = c.Send(request, response)
    return
}

func NewExportResourcePackageDeductDetailsRequest() (request *ExportResourcePackageDeductDetailsRequest) {
    request = &ExportResourcePackageDeductDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ExportResourcePackageDeductDetails")
    
    
    return
}

func NewExportResourcePackageDeductDetailsResponse() (response *ExportResourcePackageDeductDetailsResponse) {
    response = &ExportResourcePackageDeductDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExportResourcePackageDeductDetails
// This API is used to export the usage details of a resource package.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_QUERYSOURCEPACKAGEERROR = "FailedOperation.QuerySourcePackageError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ExportResourcePackageDeductDetails(request *ExportResourcePackageDeductDetailsRequest) (response *ExportResourcePackageDeductDetailsResponse, err error) {
    return c.ExportResourcePackageDeductDetailsWithContext(context.Background(), request)
}

// ExportResourcePackageDeductDetails
// This API is used to export the usage details of a resource package.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_QUERYSOURCEPACKAGEERROR = "FailedOperation.QuerySourcePackageError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ExportResourcePackageDeductDetailsWithContext(ctx context.Context, request *ExportResourcePackageDeductDetailsRequest) (response *ExportResourcePackageDeductDetailsResponse, err error) {
    if request == nil {
        request = NewExportResourcePackageDeductDetailsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ExportResourcePackageDeductDetails")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportResourcePackageDeductDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportResourcePackageDeductDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceCreateRequest() (request *InquirePriceCreateRequest) {
    request = &InquirePriceCreateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "InquirePriceCreate")
    
    
    return
}

func NewInquirePriceCreateResponse() (response *InquirePriceCreateResponse) {
    response = &InquirePriceCreateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquirePriceCreate
// This interface (InquirePriceCreate) is used for price inquiry of newly purchased clusters.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) InquirePriceCreate(request *InquirePriceCreateRequest) (response *InquirePriceCreateResponse, err error) {
    return c.InquirePriceCreateWithContext(context.Background(), request)
}

// InquirePriceCreate
// This interface (InquirePriceCreate) is used for price inquiry of newly purchased clusters.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) InquirePriceCreateWithContext(ctx context.Context, request *InquirePriceCreateRequest) (response *InquirePriceCreateResponse, err error) {
    if request == nil {
        request = NewInquirePriceCreateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "InquirePriceCreate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquirePriceCreate require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquirePriceCreateResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceModifyRequest() (request *InquirePriceModifyRequest) {
    request = &InquirePriceModifyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "InquirePriceModify")
    
    
    return
}

func NewInquirePriceModifyResponse() (response *InquirePriceModifyResponse) {
    response = &InquirePriceModifyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquirePriceModify
// This API is used to query the price for modifying the specifications of a prepaid cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) InquirePriceModify(request *InquirePriceModifyRequest) (response *InquirePriceModifyResponse, err error) {
    return c.InquirePriceModifyWithContext(context.Background(), request)
}

// InquirePriceModify
// This API is used to query the price for modifying the specifications of a prepaid cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) InquirePriceModifyWithContext(ctx context.Context, request *InquirePriceModifyRequest) (response *InquirePriceModifyResponse, err error) {
    if request == nil {
        request = NewInquirePriceModifyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "InquirePriceModify")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquirePriceModify require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquirePriceModifyResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceRenewRequest() (request *InquirePriceRenewRequest) {
    request = &InquirePriceRenewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "InquirePriceRenew")
    
    
    return
}

func NewInquirePriceRenewResponse() (response *InquirePriceRenewResponse) {
    response = &InquirePriceRenewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquirePriceRenew
// This API is used to query the renewal price of a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_BATCHGETINSTANCEERROR = "FailedOperation.BatchGetInstanceError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) InquirePriceRenew(request *InquirePriceRenewRequest) (response *InquirePriceRenewResponse, err error) {
    return c.InquirePriceRenewWithContext(context.Background(), request)
}

// InquirePriceRenew
// This API is used to query the renewal price of a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_BATCHGETINSTANCEERROR = "FailedOperation.BatchGetInstanceError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) InquirePriceRenewWithContext(ctx context.Context, request *InquirePriceRenewRequest) (response *InquirePriceRenewResponse, err error) {
    if request == nil {
        request = NewInquirePriceRenewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "InquirePriceRenew")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquirePriceRenew require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquirePriceRenewResponse()
    err = c.Send(request, response)
    return
}

func NewIsolateClusterRequest() (request *IsolateClusterRequest) {
    request = &IsolateClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "IsolateCluster")
    
    
    return
}

func NewIsolateClusterResponse() (response *IsolateClusterResponse) {
    response = &IsolateClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// IsolateCluster
// This interface (IsolateCluster) is used to isolate a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREGIONIDERROR = "InvalidParameterValue.InvalidRegionIdError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) IsolateCluster(request *IsolateClusterRequest) (response *IsolateClusterResponse, err error) {
    return c.IsolateClusterWithContext(context.Background(), request)
}

// IsolateCluster
// This interface (IsolateCluster) is used to isolate a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREGIONIDERROR = "InvalidParameterValue.InvalidRegionIdError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) IsolateClusterWithContext(ctx context.Context, request *IsolateClusterRequest) (response *IsolateClusterResponse, err error) {
    if request == nil {
        request = NewIsolateClusterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "IsolateCluster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("IsolateCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewIsolateClusterResponse()
    err = c.Send(request, response)
    return
}

func NewIsolateInstanceRequest() (request *IsolateInstanceRequest) {
    request = &IsolateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "IsolateInstance")
    
    
    return
}

func NewIsolateInstanceResponse() (response *IsolateInstanceResponse) {
    response = &IsolateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// IsolateInstance
// This API is used to isolate an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCENOTFOUND_RESOURCEERROR = "ResourceNotFound.ResourceError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) IsolateInstance(request *IsolateInstanceRequest) (response *IsolateInstanceResponse, err error) {
    return c.IsolateInstanceWithContext(context.Background(), request)
}

// IsolateInstance
// This API is used to isolate an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCENOTFOUND_RESOURCEERROR = "ResourceNotFound.ResourceError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) IsolateInstanceWithContext(ctx context.Context, request *IsolateInstanceRequest) (response *IsolateInstanceResponse, err error) {
    if request == nil {
        request = NewIsolateInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "IsolateInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("IsolateInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewIsolateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccountDescriptionRequest() (request *ModifyAccountDescriptionRequest) {
    request = &ModifyAccountDescriptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyAccountDescription")
    
    
    return
}

func NewModifyAccountDescriptionResponse() (response *ModifyAccountDescriptionResponse) {
    response = &ModifyAccountDescriptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAccountDescription
// This API is used to modify the descriptions of a database account.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyAccountDescription(request *ModifyAccountDescriptionRequest) (response *ModifyAccountDescriptionResponse, err error) {
    return c.ModifyAccountDescriptionWithContext(context.Background(), request)
}

// ModifyAccountDescription
// This API is used to modify the descriptions of a database account.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyAccountDescriptionWithContext(ctx context.Context, request *ModifyAccountDescriptionRequest) (response *ModifyAccountDescriptionResponse, err error) {
    if request == nil {
        request = NewModifyAccountDescriptionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyAccountDescription")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccountDescription require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccountDescriptionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccountHostRequest() (request *ModifyAccountHostRequest) {
    request = &ModifyAccountHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyAccountHost")
    
    
    return
}

func NewModifyAccountHostResponse() (response *ModifyAccountHostResponse) {
    response = &ModifyAccountHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAccountHost
// This API is used to modify account hosts.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyAccountHost(request *ModifyAccountHostRequest) (response *ModifyAccountHostResponse, err error) {
    return c.ModifyAccountHostWithContext(context.Background(), request)
}

// ModifyAccountHost
// This API is used to modify account hosts.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyAccountHostWithContext(ctx context.Context, request *ModifyAccountHostRequest) (response *ModifyAccountHostResponse, err error) {
    if request == nil {
        request = NewModifyAccountHostRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyAccountHost")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccountHost require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccountHostResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccountPrivilegesRequest() (request *ModifyAccountPrivilegesRequest) {
    request = &ModifyAccountPrivilegesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyAccountPrivileges")
    
    
    return
}

func NewModifyAccountPrivilegesResponse() (response *ModifyAccountPrivilegesResponse) {
    response = &ModifyAccountPrivilegesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAccountPrivileges
// This API is used to modify account database and table permissions.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyAccountPrivileges(request *ModifyAccountPrivilegesRequest) (response *ModifyAccountPrivilegesResponse, err error) {
    return c.ModifyAccountPrivilegesWithContext(context.Background(), request)
}

// ModifyAccountPrivileges
// This API is used to modify account database and table permissions.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyAccountPrivilegesWithContext(ctx context.Context, request *ModifyAccountPrivilegesRequest) (response *ModifyAccountPrivilegesResponse, err error) {
    if request == nil {
        request = NewModifyAccountPrivilegesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyAccountPrivileges")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccountPrivileges require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccountPrivilegesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAuditRuleTemplatesRequest() (request *ModifyAuditRuleTemplatesRequest) {
    request = &ModifyAuditRuleTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyAuditRuleTemplates")
    
    
    return
}

func NewModifyAuditRuleTemplatesResponse() (response *ModifyAuditRuleTemplatesResponse) {
    response = &ModifyAuditRuleTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAuditRuleTemplates
// This API is used to modify audit rule templates.
//
// error code that may be returned:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
func (c *Client) ModifyAuditRuleTemplates(request *ModifyAuditRuleTemplatesRequest) (response *ModifyAuditRuleTemplatesResponse, err error) {
    return c.ModifyAuditRuleTemplatesWithContext(context.Background(), request)
}

// ModifyAuditRuleTemplates
// This API is used to modify audit rule templates.
//
// error code that may be returned:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
func (c *Client) ModifyAuditRuleTemplatesWithContext(ctx context.Context, request *ModifyAuditRuleTemplatesRequest) (response *ModifyAuditRuleTemplatesResponse, err error) {
    if request == nil {
        request = NewModifyAuditRuleTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyAuditRuleTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAuditRuleTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAuditRuleTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAuditServiceRequest() (request *ModifyAuditServiceRequest) {
    request = &ModifyAuditServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyAuditService")
    
    
    return
}

func NewModifyAuditServiceResponse() (response *ModifyAuditServiceResponse) {
    response = &ModifyAuditServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAuditService
// This API is used to modify the audit configurations of an instance, such as audit log retention period and audit rule.
//
// error code that may be returned:
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  OPERATIONDENIED_AUDITSTATUSERROR = "OperationDenied.AuditStatusError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
func (c *Client) ModifyAuditService(request *ModifyAuditServiceRequest) (response *ModifyAuditServiceResponse, err error) {
    return c.ModifyAuditServiceWithContext(context.Background(), request)
}

// ModifyAuditService
// This API is used to modify the audit configurations of an instance, such as audit log retention period and audit rule.
//
// error code that may be returned:
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  OPERATIONDENIED_AUDITSTATUSERROR = "OperationDenied.AuditStatusError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
func (c *Client) ModifyAuditServiceWithContext(ctx context.Context, request *ModifyAuditServiceRequest) (response *ModifyAuditServiceResponse, err error) {
    if request == nil {
        request = NewModifyAuditServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyAuditService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAuditService require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAuditServiceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBackupConfigRequest() (request *ModifyBackupConfigRequest) {
    request = &ModifyBackupConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyBackupConfig")
    
    
    return
}

func NewModifyBackupConfigResponse() (response *ModifyBackupConfigResponse) {
    response = &ModifyBackupConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyBackupConfig
// This API is used to modify the backup configuration of a specified cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_CYNOSDBMYSQLSETBACKUPSTRATEGY = "FailedOperation.CynosdbMysqlSetBackupStrategy"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyBackupConfig(request *ModifyBackupConfigRequest) (response *ModifyBackupConfigResponse, err error) {
    return c.ModifyBackupConfigWithContext(context.Background(), request)
}

// ModifyBackupConfig
// This API is used to modify the backup configuration of a specified cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_CYNOSDBMYSQLSETBACKUPSTRATEGY = "FailedOperation.CynosdbMysqlSetBackupStrategy"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyBackupConfigWithContext(ctx context.Context, request *ModifyBackupConfigRequest) (response *ModifyBackupConfigResponse, err error) {
    if request == nil {
        request = NewModifyBackupConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyBackupConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBackupConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBackupConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBackupDownloadRestrictionRequest() (request *ModifyBackupDownloadRestrictionRequest) {
    request = &ModifyBackupDownloadRestrictionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyBackupDownloadRestriction")
    
    
    return
}

func NewModifyBackupDownloadRestrictionResponse() (response *ModifyBackupDownloadRestrictionResponse) {
    response = &ModifyBackupDownloadRestrictionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyBackupDownloadRestriction
// This API is used to modify the download source limit of the backup file for the user in the current region. It can be configured to allow downloads from both private and public networks, only the private network, or a designated vpc or ip within the private network.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREGIONIDERROR = "InvalidParameterValue.InvalidRegionIdError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyBackupDownloadRestriction(request *ModifyBackupDownloadRestrictionRequest) (response *ModifyBackupDownloadRestrictionResponse, err error) {
    return c.ModifyBackupDownloadRestrictionWithContext(context.Background(), request)
}

// ModifyBackupDownloadRestriction
// This API is used to modify the download source limit of the backup file for the user in the current region. It can be configured to allow downloads from both private and public networks, only the private network, or a designated vpc or ip within the private network.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREGIONIDERROR = "InvalidParameterValue.InvalidRegionIdError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyBackupDownloadRestrictionWithContext(ctx context.Context, request *ModifyBackupDownloadRestrictionRequest) (response *ModifyBackupDownloadRestrictionResponse, err error) {
    if request == nil {
        request = NewModifyBackupDownloadRestrictionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyBackupDownloadRestriction")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBackupDownloadRestriction require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBackupDownloadRestrictionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBackupDownloadUserRestrictionRequest() (request *ModifyBackupDownloadUserRestrictionRequest) {
    request = &ModifyBackupDownloadUserRestrictionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyBackupDownloadUserRestriction")
    
    
    return
}

func NewModifyBackupDownloadUserRestrictionResponse() (response *ModifyBackupDownloadUserRestrictionResponse) {
    response = &ModifyBackupDownloadUserRestrictionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyBackupDownloadUserRestriction
// This API is used to modify the download source restrictions for backup files in the user's current region. It can be configured to allow downloads from both private and public networks, only from a private network, or from a designated vpc or ip within the private network.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREGIONIDERROR = "InvalidParameterValue.InvalidRegionIdError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyBackupDownloadUserRestriction(request *ModifyBackupDownloadUserRestrictionRequest) (response *ModifyBackupDownloadUserRestrictionResponse, err error) {
    return c.ModifyBackupDownloadUserRestrictionWithContext(context.Background(), request)
}

// ModifyBackupDownloadUserRestriction
// This API is used to modify the download source restrictions for backup files in the user's current region. It can be configured to allow downloads from both private and public networks, only from a private network, or from a designated vpc or ip within the private network.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREGIONIDERROR = "InvalidParameterValue.InvalidRegionIdError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyBackupDownloadUserRestrictionWithContext(ctx context.Context, request *ModifyBackupDownloadUserRestrictionRequest) (response *ModifyBackupDownloadUserRestrictionResponse, err error) {
    if request == nil {
        request = NewModifyBackupDownloadUserRestrictionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyBackupDownloadUserRestriction")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBackupDownloadUserRestriction require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBackupDownloadUserRestrictionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBackupNameRequest() (request *ModifyBackupNameRequest) {
    request = &ModifyBackupNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyBackupName")
    
    
    return
}

func NewModifyBackupNameResponse() (response *ModifyBackupNameResponse) {
    response = &ModifyBackupNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyBackupName
// This API is used to rename a backup file.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyBackupName(request *ModifyBackupNameRequest) (response *ModifyBackupNameResponse, err error) {
    return c.ModifyBackupNameWithContext(context.Background(), request)
}

// ModifyBackupName
// This API is used to rename a backup file.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyBackupNameWithContext(ctx context.Context, request *ModifyBackupNameRequest) (response *ModifyBackupNameResponse, err error) {
    if request == nil {
        request = NewModifyBackupNameRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyBackupName")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBackupName require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBackupNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBinlogConfigRequest() (request *ModifyBinlogConfigRequest) {
    request = &ModifyBinlogConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyBinlogConfig")
    
    
    return
}

func NewModifyBinlogConfigResponse() (response *ModifyBinlogConfigResponse) {
    response = &ModifyBinlogConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyBinlogConfig
// This API is used to modify Binlog configuration.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
func (c *Client) ModifyBinlogConfig(request *ModifyBinlogConfigRequest) (response *ModifyBinlogConfigResponse, err error) {
    return c.ModifyBinlogConfigWithContext(context.Background(), request)
}

// ModifyBinlogConfig
// This API is used to modify Binlog configuration.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
func (c *Client) ModifyBinlogConfigWithContext(ctx context.Context, request *ModifyBinlogConfigRequest) (response *ModifyBinlogConfigResponse, err error) {
    if request == nil {
        request = NewModifyBinlogConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyBinlogConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBinlogConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBinlogConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBinlogSaveDaysRequest() (request *ModifyBinlogSaveDaysRequest) {
    request = &ModifyBinlogSaveDaysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyBinlogSaveDays")
    
    
    return
}

func NewModifyBinlogSaveDaysResponse() (response *ModifyBinlogSaveDaysResponse) {
    response = &ModifyBinlogSaveDaysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyBinlogSaveDays
// This API is used to modify the binlog retention period in days.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyBinlogSaveDays(request *ModifyBinlogSaveDaysRequest) (response *ModifyBinlogSaveDaysResponse, err error) {
    return c.ModifyBinlogSaveDaysWithContext(context.Background(), request)
}

// ModifyBinlogSaveDays
// This API is used to modify the binlog retention period in days.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyBinlogSaveDaysWithContext(ctx context.Context, request *ModifyBinlogSaveDaysRequest) (response *ModifyBinlogSaveDaysResponse, err error) {
    if request == nil {
        request = NewModifyBinlogSaveDaysRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyBinlogSaveDays")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBinlogSaveDays require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBinlogSaveDaysResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterDatabaseRequest() (request *ModifyClusterDatabaseRequest) {
    request = &ModifyClusterDatabaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyClusterDatabase")
    
    
    return
}

func NewModifyClusterDatabaseResponse() (response *ModifyClusterDatabaseResponse) {
    response = &ModifyClusterDatabaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyClusterDatabase
// This API is used to modify account authorization of a database.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyClusterDatabase(request *ModifyClusterDatabaseRequest) (response *ModifyClusterDatabaseResponse, err error) {
    return c.ModifyClusterDatabaseWithContext(context.Background(), request)
}

// ModifyClusterDatabase
// This API is used to modify account authorization of a database.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyClusterDatabaseWithContext(ctx context.Context, request *ModifyClusterDatabaseRequest) (response *ModifyClusterDatabaseResponse, err error) {
    if request == nil {
        request = NewModifyClusterDatabaseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyClusterDatabase")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterDatabase require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterDatabaseResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterNameRequest() (request *ModifyClusterNameRequest) {
    request = &ModifyClusterNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyClusterName")
    
    
    return
}

func NewModifyClusterNameResponse() (response *ModifyClusterNameResponse) {
    response = &ModifyClusterNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyClusterName
// This API is used to modify cluster names.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyClusterName(request *ModifyClusterNameRequest) (response *ModifyClusterNameResponse, err error) {
    return c.ModifyClusterNameWithContext(context.Background(), request)
}

// ModifyClusterName
// This API is used to modify cluster names.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyClusterNameWithContext(ctx context.Context, request *ModifyClusterNameRequest) (response *ModifyClusterNameResponse, err error) {
    if request == nil {
        request = NewModifyClusterNameRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyClusterName")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterName require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterParamRequest() (request *ModifyClusterParamRequest) {
    request = &ModifyClusterParamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyClusterParam")
    
    
    return
}

func NewModifyClusterParamResponse() (response *ModifyClusterParamResponse) {
    response = &ModifyClusterParamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyClusterParam
// This API is used to modify cluster parameters.
//
// error code that may be returned:
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  OPERATIONDENIED_TASKCONFLICTERROR = "OperationDenied.TaskConflictError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyClusterParam(request *ModifyClusterParamRequest) (response *ModifyClusterParamResponse, err error) {
    return c.ModifyClusterParamWithContext(context.Background(), request)
}

// ModifyClusterParam
// This API is used to modify cluster parameters.
//
// error code that may be returned:
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  OPERATIONDENIED_TASKCONFLICTERROR = "OperationDenied.TaskConflictError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyClusterParamWithContext(ctx context.Context, request *ModifyClusterParamRequest) (response *ModifyClusterParamResponse, err error) {
    if request == nil {
        request = NewModifyClusterParamRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyClusterParam")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterParam require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterParamResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterPasswordComplexityRequest() (request *ModifyClusterPasswordComplexityRequest) {
    request = &ModifyClusterPasswordComplexityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyClusterPasswordComplexity")
    
    
    return
}

func NewModifyClusterPasswordComplexityResponse() (response *ModifyClusterPasswordComplexityResponse) {
    response = &ModifyClusterPasswordComplexityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyClusterPasswordComplexity
// This API is used to modify or enable cluster password complexity.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyClusterPasswordComplexity(request *ModifyClusterPasswordComplexityRequest) (response *ModifyClusterPasswordComplexityResponse, err error) {
    return c.ModifyClusterPasswordComplexityWithContext(context.Background(), request)
}

// ModifyClusterPasswordComplexity
// This API is used to modify or enable cluster password complexity.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyClusterPasswordComplexityWithContext(ctx context.Context, request *ModifyClusterPasswordComplexityRequest) (response *ModifyClusterPasswordComplexityResponse, err error) {
    if request == nil {
        request = NewModifyClusterPasswordComplexityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyClusterPasswordComplexity")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterPasswordComplexity require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterPasswordComplexityResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterReadOnlyRequest() (request *ModifyClusterReadOnlyRequest) {
    request = &ModifyClusterReadOnlyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyClusterReadOnly")
    
    
    return
}

func NewModifyClusterReadOnlyResponse() (response *ModifyClusterReadOnlyResponse) {
    response = &ModifyClusterReadOnlyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyClusterReadOnly
// This API is used to modify the read-only switch of a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  OPERATIONDENIED_TASKCONFLICTERROR = "OperationDenied.TaskConflictError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyClusterReadOnly(request *ModifyClusterReadOnlyRequest) (response *ModifyClusterReadOnlyResponse, err error) {
    return c.ModifyClusterReadOnlyWithContext(context.Background(), request)
}

// ModifyClusterReadOnly
// This API is used to modify the read-only switch of a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  OPERATIONDENIED_TASKCONFLICTERROR = "OperationDenied.TaskConflictError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyClusterReadOnlyWithContext(ctx context.Context, request *ModifyClusterReadOnlyRequest) (response *ModifyClusterReadOnlyResponse, err error) {
    if request == nil {
        request = NewModifyClusterReadOnlyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyClusterReadOnly")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterReadOnly require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterReadOnlyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterSlaveZoneRequest() (request *ModifyClusterSlaveZoneRequest) {
    request = &ModifyClusterSlaveZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyClusterSlaveZone")
    
    
    return
}

func NewModifyClusterSlaveZoneResponse() (response *ModifyClusterSlaveZoneResponse) {
    response = &ModifyClusterSlaveZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyClusterSlaveZone
// This API is used to modify the slave availability zone of a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_PAUSEDSLSNOTALLOWMODIFYSLAVE = "OperationDenied.PausedSlsNotAllowModifySlave"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyClusterSlaveZone(request *ModifyClusterSlaveZoneRequest) (response *ModifyClusterSlaveZoneResponse, err error) {
    return c.ModifyClusterSlaveZoneWithContext(context.Background(), request)
}

// ModifyClusterSlaveZone
// This API is used to modify the slave availability zone of a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_PAUSEDSLSNOTALLOWMODIFYSLAVE = "OperationDenied.PausedSlsNotAllowModifySlave"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyClusterSlaveZoneWithContext(ctx context.Context, request *ModifyClusterSlaveZoneRequest) (response *ModifyClusterSlaveZoneResponse, err error) {
    if request == nil {
        request = NewModifyClusterSlaveZoneRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyClusterSlaveZone")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterSlaveZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterSlaveZoneResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceSecurityGroupsRequest() (request *ModifyDBInstanceSecurityGroupsRequest) {
    request = &ModifyDBInstanceSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyDBInstanceSecurityGroups")
    
    
    return
}

func NewModifyDBInstanceSecurityGroupsResponse() (response *ModifyDBInstanceSecurityGroupsResponse) {
    response = &ModifyDBInstanceSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBInstanceSecurityGroups
// This API is used to modify the security group bound to the instance.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDBInstanceSecurityGroups(request *ModifyDBInstanceSecurityGroupsRequest) (response *ModifyDBInstanceSecurityGroupsResponse, err error) {
    return c.ModifyDBInstanceSecurityGroupsWithContext(context.Background(), request)
}

// ModifyDBInstanceSecurityGroups
// This API is used to modify the security group bound to the instance.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDBInstanceSecurityGroupsWithContext(ctx context.Context, request *ModifyDBInstanceSecurityGroupsRequest) (response *ModifyDBInstanceSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceSecurityGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyDBInstanceSecurityGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBInstanceSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceNameRequest() (request *ModifyInstanceNameRequest) {
    request = &ModifyInstanceNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyInstanceName")
    
    
    return
}

func NewModifyInstanceNameResponse() (response *ModifyInstanceNameResponse) {
    response = &ModifyInstanceNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstanceName
// This API is used to modify instance name.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyInstanceName(request *ModifyInstanceNameRequest) (response *ModifyInstanceNameResponse, err error) {
    return c.ModifyInstanceNameWithContext(context.Background(), request)
}

// ModifyInstanceName
// This API is used to modify instance name.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyInstanceNameWithContext(ctx context.Context, request *ModifyInstanceNameRequest) (response *ModifyInstanceNameResponse, err error) {
    if request == nil {
        request = NewModifyInstanceNameRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyInstanceName")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceName require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceParamRequest() (request *ModifyInstanceParamRequest) {
    request = &ModifyInstanceParamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyInstanceParam")
    
    
    return
}

func NewModifyInstanceParamResponse() (response *ModifyInstanceParamResponse) {
    response = &ModifyInstanceParamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstanceParam
// This API is used to modify the instance parameters.
//
// error code that may be returned:
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyInstanceParam(request *ModifyInstanceParamRequest) (response *ModifyInstanceParamResponse, err error) {
    return c.ModifyInstanceParamWithContext(context.Background(), request)
}

// ModifyInstanceParam
// This API is used to modify the instance parameters.
//
// error code that may be returned:
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyInstanceParamWithContext(ctx context.Context, request *ModifyInstanceParamRequest) (response *ModifyInstanceParamResponse, err error) {
    if request == nil {
        request = NewModifyInstanceParamRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyInstanceParam")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceParam require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceParamResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMaintainPeriodConfigRequest() (request *ModifyMaintainPeriodConfigRequest) {
    request = &ModifyMaintainPeriodConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyMaintainPeriodConfig")
    
    
    return
}

func NewModifyMaintainPeriodConfigResponse() (response *ModifyMaintainPeriodConfigResponse) {
    response = &ModifyMaintainPeriodConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMaintainPeriodConfig
// This API is used to modify maintenance time configuration.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyMaintainPeriodConfig(request *ModifyMaintainPeriodConfigRequest) (response *ModifyMaintainPeriodConfigResponse, err error) {
    return c.ModifyMaintainPeriodConfigWithContext(context.Background(), request)
}

// ModifyMaintainPeriodConfig
// This API is used to modify maintenance time configuration.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyMaintainPeriodConfigWithContext(ctx context.Context, request *ModifyMaintainPeriodConfigRequest) (response *ModifyMaintainPeriodConfigResponse, err error) {
    if request == nil {
        request = NewModifyMaintainPeriodConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyMaintainPeriodConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMaintainPeriodConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMaintainPeriodConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyParamTemplateRequest() (request *ModifyParamTemplateRequest) {
    request = &ModifyParamTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyParamTemplate")
    
    
    return
}

func NewModifyParamTemplateResponse() (response *ModifyParamTemplateResponse) {
    response = &ModifyParamTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyParamTemplate
// This API is used to modify a parameter template.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
func (c *Client) ModifyParamTemplate(request *ModifyParamTemplateRequest) (response *ModifyParamTemplateResponse, err error) {
    return c.ModifyParamTemplateWithContext(context.Background(), request)
}

// ModifyParamTemplate
// This API is used to modify a parameter template.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
func (c *Client) ModifyParamTemplateWithContext(ctx context.Context, request *ModifyParamTemplateRequest) (response *ModifyParamTemplateResponse, err error) {
    if request == nil {
        request = NewModifyParamTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyParamTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyParamTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyParamTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProxyDescRequest() (request *ModifyProxyDescRequest) {
    request = &ModifyProxyDescRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyProxyDesc")
    
    
    return
}

func NewModifyProxyDescResponse() (response *ModifyProxyDescResponse) {
    response = &ModifyProxyDescResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyProxyDesc
// This API is used to modify the description of a database proxy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ModifyProxyDesc(request *ModifyProxyDescRequest) (response *ModifyProxyDescResponse, err error) {
    return c.ModifyProxyDescWithContext(context.Background(), request)
}

// ModifyProxyDesc
// This API is used to modify the description of a database proxy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ModifyProxyDescWithContext(ctx context.Context, request *ModifyProxyDescRequest) (response *ModifyProxyDescResponse, err error) {
    if request == nil {
        request = NewModifyProxyDescRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyProxyDesc")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyProxyDesc require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyProxyDescResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProxyRwSplitRequest() (request *ModifyProxyRwSplitRequest) {
    request = &ModifyProxyRwSplitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyProxyRwSplit")
    
    
    return
}

func NewModifyProxyRwSplitResponse() (response *ModifyProxyRwSplitResponse) {
    response = &ModifyProxyRwSplitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyProxyRwSplit
// This API is used to configure read-write separation for database proxy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyProxyRwSplit(request *ModifyProxyRwSplitRequest) (response *ModifyProxyRwSplitResponse, err error) {
    return c.ModifyProxyRwSplitWithContext(context.Background(), request)
}

// ModifyProxyRwSplit
// This API is used to configure read-write separation for database proxy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyProxyRwSplitWithContext(ctx context.Context, request *ModifyProxyRwSplitRequest) (response *ModifyProxyRwSplitResponse, err error) {
    if request == nil {
        request = NewModifyProxyRwSplitRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyProxyRwSplit")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyProxyRwSplit require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyProxyRwSplitResponse()
    err = c.Send(request, response)
    return
}

func NewModifyResourcePackageClustersRequest() (request *ModifyResourcePackageClustersRequest) {
    request = &ModifyResourcePackageClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyResourcePackageClusters")
    
    
    return
}

func NewModifyResourcePackageClustersResponse() (response *ModifyResourcePackageClustersResponse) {
    response = &ModifyResourcePackageClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyResourcePackageClusters
// This API is used to modify the binding relationship between resource packages and clusters.
//
// error code that may be returned:
//  FAILEDOPERATION_BINDSOURCEPACKAGEERROR = "FailedOperation.BindSourcePackageError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_QUERYSOURCEPACKAGEERROR = "FailedOperation.QuerySourcePackageError"
//  FAILEDOPERATION_UNBINDSOURCEPACKAGEERROR = "FailedOperation.UnBindSourcePackageError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyResourcePackageClusters(request *ModifyResourcePackageClustersRequest) (response *ModifyResourcePackageClustersResponse, err error) {
    return c.ModifyResourcePackageClustersWithContext(context.Background(), request)
}

// ModifyResourcePackageClusters
// This API is used to modify the binding relationship between resource packages and clusters.
//
// error code that may be returned:
//  FAILEDOPERATION_BINDSOURCEPACKAGEERROR = "FailedOperation.BindSourcePackageError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_QUERYSOURCEPACKAGEERROR = "FailedOperation.QuerySourcePackageError"
//  FAILEDOPERATION_UNBINDSOURCEPACKAGEERROR = "FailedOperation.UnBindSourcePackageError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyResourcePackageClustersWithContext(ctx context.Context, request *ModifyResourcePackageClustersRequest) (response *ModifyResourcePackageClustersResponse, err error) {
    if request == nil {
        request = NewModifyResourcePackageClustersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyResourcePackageClusters")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyResourcePackageClusters require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyResourcePackageClustersResponse()
    err = c.Send(request, response)
    return
}

func NewModifyResourcePackageNameRequest() (request *ModifyResourcePackageNameRequest) {
    request = &ModifyResourcePackageNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyResourcePackageName")
    
    
    return
}

func NewModifyResourcePackageNameResponse() (response *ModifyResourcePackageNameResponse) {
    response = &ModifyResourcePackageNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyResourcePackageName
// This API is used to modify resource package name.
//
// error code that may be returned:
//  FAILEDOPERATION_BINDSOURCEPACKAGEERROR = "FailedOperation.BindSourcePackageError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_QUERYSOURCEPACKAGEERROR = "FailedOperation.QuerySourcePackageError"
//  FAILEDOPERATION_UNBINDSOURCEPACKAGEERROR = "FailedOperation.UnBindSourcePackageError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyResourcePackageName(request *ModifyResourcePackageNameRequest) (response *ModifyResourcePackageNameResponse, err error) {
    return c.ModifyResourcePackageNameWithContext(context.Background(), request)
}

// ModifyResourcePackageName
// This API is used to modify resource package name.
//
// error code that may be returned:
//  FAILEDOPERATION_BINDSOURCEPACKAGEERROR = "FailedOperation.BindSourcePackageError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_QUERYSOURCEPACKAGEERROR = "FailedOperation.QuerySourcePackageError"
//  FAILEDOPERATION_UNBINDSOURCEPACKAGEERROR = "FailedOperation.UnBindSourcePackageError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyResourcePackageNameWithContext(ctx context.Context, request *ModifyResourcePackageNameRequest) (response *ModifyResourcePackageNameResponse, err error) {
    if request == nil {
        request = NewModifyResourcePackageNameRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyResourcePackageName")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyResourcePackageName require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyResourcePackageNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyResourcePackagesDeductionPriorityRequest() (request *ModifyResourcePackagesDeductionPriorityRequest) {
    request = &ModifyResourcePackagesDeductionPriorityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyResourcePackagesDeductionPriority")
    
    
    return
}

func NewModifyResourcePackagesDeductionPriorityResponse() (response *ModifyResourcePackagesDeductionPriorityResponse) {
    response = &ModifyResourcePackagesDeductionPriorityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyResourcePackagesDeductionPriority
// This API is used to modify the deduction priority of the bound resource package.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_MODIFYDEDUCTIONPRIORITYERROR = "FailedOperation.ModifyDeductionPriorityError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyResourcePackagesDeductionPriority(request *ModifyResourcePackagesDeductionPriorityRequest) (response *ModifyResourcePackagesDeductionPriorityResponse, err error) {
    return c.ModifyResourcePackagesDeductionPriorityWithContext(context.Background(), request)
}

// ModifyResourcePackagesDeductionPriority
// This API is used to modify the deduction priority of the bound resource package.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_MODIFYDEDUCTIONPRIORITYERROR = "FailedOperation.ModifyDeductionPriorityError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyResourcePackagesDeductionPriorityWithContext(ctx context.Context, request *ModifyResourcePackagesDeductionPriorityRequest) (response *ModifyResourcePackagesDeductionPriorityResponse, err error) {
    if request == nil {
        request = NewModifyResourcePackagesDeductionPriorityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyResourcePackagesDeductionPriority")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyResourcePackagesDeductionPriority require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyResourcePackagesDeductionPriorityResponse()
    err = c.Send(request, response)
    return
}

func NewModifyServerlessStrategyRequest() (request *ModifyServerlessStrategyRequest) {
    request = &ModifyServerlessStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyServerlessStrategy")
    
    
    return
}

func NewModifyServerlessStrategyResponse() (response *ModifyServerlessStrategyResponse) {
    response = &ModifyServerlessStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyServerlessStrategy
// This API is used to modify the serverless policy.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_SERVERLESSSETSTRATEGYERROR = "FailedOperation.ServerlessSetStrategyError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_DBMODENOTSERVERLESSERROR = "InvalidParameterValue.DbModeNotServerlessError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyServerlessStrategy(request *ModifyServerlessStrategyRequest) (response *ModifyServerlessStrategyResponse, err error) {
    return c.ModifyServerlessStrategyWithContext(context.Background(), request)
}

// ModifyServerlessStrategy
// This API is used to modify the serverless policy.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_SERVERLESSSETSTRATEGYERROR = "FailedOperation.ServerlessSetStrategyError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_DBMODENOTSERVERLESSERROR = "InvalidParameterValue.DbModeNotServerlessError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyServerlessStrategyWithContext(ctx context.Context, request *ModifyServerlessStrategyRequest) (response *ModifyServerlessStrategyResponse, err error) {
    if request == nil {
        request = NewModifyServerlessStrategyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyServerlessStrategy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyServerlessStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyServerlessStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVipVportRequest() (request *ModifyVipVportRequest) {
    request = &ModifyVipVportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyVipVport")
    
    
    return
}

func NewModifyVipVportResponse() (response *ModifyVipVportResponse) {
    response = &ModifyVipVportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyVipVport
// This API is used to modify the ip and port of an instance group.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyVipVport(request *ModifyVipVportRequest) (response *ModifyVipVportResponse, err error) {
    return c.ModifyVipVportWithContext(context.Background(), request)
}

// ModifyVipVport
// This API is used to modify the ip and port of an instance group.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyVipVportWithContext(ctx context.Context, request *ModifyVipVportRequest) (response *ModifyVipVportResponse, err error) {
    if request == nil {
        request = NewModifyVipVportRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyVipVport")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVipVport require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVipVportResponse()
    err = c.Send(request, response)
    return
}

func NewOfflineClusterRequest() (request *OfflineClusterRequest) {
    request = &OfflineClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "OfflineCluster")
    
    
    return
}

func NewOfflineClusterResponse() (response *OfflineClusterResponse) {
    response = &OfflineClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OfflineCluster
// This interface (OfflineCluster) is used to terminate clusters.
//
// error code that may be returned:
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) OfflineCluster(request *OfflineClusterRequest) (response *OfflineClusterResponse, err error) {
    return c.OfflineClusterWithContext(context.Background(), request)
}

// OfflineCluster
// This interface (OfflineCluster) is used to terminate clusters.
//
// error code that may be returned:
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) OfflineClusterWithContext(ctx context.Context, request *OfflineClusterRequest) (response *OfflineClusterResponse, err error) {
    if request == nil {
        request = NewOfflineClusterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "OfflineCluster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("OfflineCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewOfflineClusterResponse()
    err = c.Send(request, response)
    return
}

func NewOfflineInstanceRequest() (request *OfflineInstanceRequest) {
    request = &OfflineInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "OfflineInstance")
    
    
    return
}

func NewOfflineInstanceResponse() (response *OfflineInstanceResponse) {
    response = &OfflineInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OfflineInstance
// This interface (OfflineInstance) is used to terminate an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCENOTFOUND_RESOURCEERROR = "ResourceNotFound.ResourceError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) OfflineInstance(request *OfflineInstanceRequest) (response *OfflineInstanceResponse, err error) {
    return c.OfflineInstanceWithContext(context.Background(), request)
}

// OfflineInstance
// This interface (OfflineInstance) is used to terminate an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCENOTFOUND_RESOURCEERROR = "ResourceNotFound.ResourceError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) OfflineInstanceWithContext(ctx context.Context, request *OfflineInstanceRequest) (response *OfflineInstanceResponse, err error) {
    if request == nil {
        request = NewOfflineInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "OfflineInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("OfflineInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewOfflineInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewOpenAuditServiceRequest() (request *OpenAuditServiceRequest) {
    request = &OpenAuditServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "OpenAuditService")
    
    
    return
}

func NewOpenAuditServiceResponse() (response *OpenAuditServiceResponse) {
    response = &OpenAuditServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OpenAuditService
// This API is used to enable database audit service for an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCEQUERYERROR = "FailedOperation.InstanceQueryError"
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
func (c *Client) OpenAuditService(request *OpenAuditServiceRequest) (response *OpenAuditServiceResponse, err error) {
    return c.OpenAuditServiceWithContext(context.Background(), request)
}

// OpenAuditService
// This API is used to enable database audit service for an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCEQUERYERROR = "FailedOperation.InstanceQueryError"
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
func (c *Client) OpenAuditServiceWithContext(ctx context.Context, request *OpenAuditServiceRequest) (response *OpenAuditServiceResponse, err error) {
    if request == nil {
        request = NewOpenAuditServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "OpenAuditService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenAuditService require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenAuditServiceResponse()
    err = c.Send(request, response)
    return
}

func NewOpenClusterPasswordComplexityRequest() (request *OpenClusterPasswordComplexityRequest) {
    request = &OpenClusterPasswordComplexityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "OpenClusterPasswordComplexity")
    
    
    return
}

func NewOpenClusterPasswordComplexityResponse() (response *OpenClusterPasswordComplexityResponse) {
    response = &OpenClusterPasswordComplexityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OpenClusterPasswordComplexity
// This API is used to enable the custom password complexity feature.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) OpenClusterPasswordComplexity(request *OpenClusterPasswordComplexityRequest) (response *OpenClusterPasswordComplexityResponse, err error) {
    return c.OpenClusterPasswordComplexityWithContext(context.Background(), request)
}

// OpenClusterPasswordComplexity
// This API is used to enable the custom password complexity feature.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) OpenClusterPasswordComplexityWithContext(ctx context.Context, request *OpenClusterPasswordComplexityRequest) (response *OpenClusterPasswordComplexityResponse, err error) {
    if request == nil {
        request = NewOpenClusterPasswordComplexityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "OpenClusterPasswordComplexity")
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenClusterPasswordComplexity require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenClusterPasswordComplexityResponse()
    err = c.Send(request, response)
    return
}

func NewOpenClusterReadOnlyInstanceGroupAccessRequest() (request *OpenClusterReadOnlyInstanceGroupAccessRequest) {
    request = &OpenClusterReadOnlyInstanceGroupAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "OpenClusterReadOnlyInstanceGroupAccess")
    
    
    return
}

func NewOpenClusterReadOnlyInstanceGroupAccessResponse() (response *OpenClusterReadOnlyInstanceGroupAccessResponse) {
    response = &OpenClusterReadOnlyInstanceGroupAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OpenClusterReadOnlyInstanceGroupAccess
// This API is used to enable read-only instance group access.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) OpenClusterReadOnlyInstanceGroupAccess(request *OpenClusterReadOnlyInstanceGroupAccessRequest) (response *OpenClusterReadOnlyInstanceGroupAccessResponse, err error) {
    return c.OpenClusterReadOnlyInstanceGroupAccessWithContext(context.Background(), request)
}

// OpenClusterReadOnlyInstanceGroupAccess
// This API is used to enable read-only instance group access.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) OpenClusterReadOnlyInstanceGroupAccessWithContext(ctx context.Context, request *OpenClusterReadOnlyInstanceGroupAccessRequest) (response *OpenClusterReadOnlyInstanceGroupAccessResponse, err error) {
    if request == nil {
        request = NewOpenClusterReadOnlyInstanceGroupAccessRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "OpenClusterReadOnlyInstanceGroupAccess")
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenClusterReadOnlyInstanceGroupAccess require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenClusterReadOnlyInstanceGroupAccessResponse()
    err = c.Send(request, response)
    return
}

func NewOpenClusterTransparentEncryptRequest() (request *OpenClusterTransparentEncryptRequest) {
    request = &OpenClusterTransparentEncryptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "OpenClusterTransparentEncrypt")
    
    
    return
}

func NewOpenClusterTransparentEncryptResponse() (response *OpenClusterTransparentEncryptResponse) {
    response = &OpenClusterTransparentEncryptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OpenClusterTransparentEncrypt
// Enable transparent data encryption for the cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) OpenClusterTransparentEncrypt(request *OpenClusterTransparentEncryptRequest) (response *OpenClusterTransparentEncryptResponse, err error) {
    return c.OpenClusterTransparentEncryptWithContext(context.Background(), request)
}

// OpenClusterTransparentEncrypt
// Enable transparent data encryption for the cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) OpenClusterTransparentEncryptWithContext(ctx context.Context, request *OpenClusterTransparentEncryptRequest) (response *OpenClusterTransparentEncryptResponse, err error) {
    if request == nil {
        request = NewOpenClusterTransparentEncryptRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "OpenClusterTransparentEncrypt")
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenClusterTransparentEncrypt require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenClusterTransparentEncryptResponse()
    err = c.Send(request, response)
    return
}

func NewOpenReadOnlyInstanceExclusiveAccessRequest() (request *OpenReadOnlyInstanceExclusiveAccessRequest) {
    request = &OpenReadOnlyInstanceExclusiveAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "OpenReadOnlyInstanceExclusiveAccess")
    
    
    return
}

func NewOpenReadOnlyInstanceExclusiveAccessResponse() (response *OpenReadOnlyInstanceExclusiveAccessResponse) {
    response = &OpenReadOnlyInstanceExclusiveAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OpenReadOnlyInstanceExclusiveAccess
// This interface (OpenReadOnlyInstanceExclusiveAccess) is used to enable the dedicated access access group for a read-only instance.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) OpenReadOnlyInstanceExclusiveAccess(request *OpenReadOnlyInstanceExclusiveAccessRequest) (response *OpenReadOnlyInstanceExclusiveAccessResponse, err error) {
    return c.OpenReadOnlyInstanceExclusiveAccessWithContext(context.Background(), request)
}

// OpenReadOnlyInstanceExclusiveAccess
// This interface (OpenReadOnlyInstanceExclusiveAccess) is used to enable the dedicated access access group for a read-only instance.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) OpenReadOnlyInstanceExclusiveAccessWithContext(ctx context.Context, request *OpenReadOnlyInstanceExclusiveAccessRequest) (response *OpenReadOnlyInstanceExclusiveAccessResponse, err error) {
    if request == nil {
        request = NewOpenReadOnlyInstanceExclusiveAccessRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "OpenReadOnlyInstanceExclusiveAccess")
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenReadOnlyInstanceExclusiveAccess require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenReadOnlyInstanceExclusiveAccessResponse()
    err = c.Send(request, response)
    return
}

func NewOpenWanRequest() (request *OpenWanRequest) {
    request = &OpenWanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "OpenWan")
    
    
    return
}

func NewOpenWanResponse() (response *OpenWanResponse) {
    response = &OpenWanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OpenWan
// This interface (OpenWan) is used to enable external network.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_GETNETSERVICEINFOERROR = "FailedOperation.GetNetServiceInfoError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) OpenWan(request *OpenWanRequest) (response *OpenWanResponse, err error) {
    return c.OpenWanWithContext(context.Background(), request)
}

// OpenWan
// This interface (OpenWan) is used to enable external network.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_GETNETSERVICEINFOERROR = "FailedOperation.GetNetServiceInfoError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) OpenWanWithContext(ctx context.Context, request *OpenWanRequest) (response *OpenWanResponse, err error) {
    if request == nil {
        request = NewOpenWanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "OpenWan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenWan require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenWanResponse()
    err = c.Send(request, response)
    return
}

func NewPauseServerlessRequest() (request *PauseServerlessRequest) {
    request = &PauseServerlessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "PauseServerless")
    
    
    return
}

func NewPauseServerlessResponse() (response *PauseServerlessResponse) {
    response = &PauseServerlessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PauseServerless
// This API is used to suspend a serverless cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_DBMODENOTSERVERLESSERROR = "InvalidParameterValue.DbModeNotServerlessError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) PauseServerless(request *PauseServerlessRequest) (response *PauseServerlessResponse, err error) {
    return c.PauseServerlessWithContext(context.Background(), request)
}

// PauseServerless
// This API is used to suspend a serverless cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_DBMODENOTSERVERLESSERROR = "InvalidParameterValue.DbModeNotServerlessError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) PauseServerlessWithContext(ctx context.Context, request *PauseServerlessRequest) (response *PauseServerlessResponse, err error) {
    if request == nil {
        request = NewPauseServerlessRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "PauseServerless")
    
    if c.GetCredential() == nil {
        return nil, errors.New("PauseServerless require credential")
    }

    request.SetContext(ctx)
    
    response = NewPauseServerlessResponse()
    err = c.Send(request, response)
    return
}

func NewRefundResourcePackageRequest() (request *RefundResourcePackageRequest) {
    request = &RefundResourcePackageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "RefundResourcePackage")
    
    
    return
}

func NewRefundResourcePackageResponse() (response *RefundResourcePackageResponse) {
    response = &RefundResourcePackageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RefundResourcePackage
// This API is used to refund a resource package.
//
// error code that may be returned:
//  FAILEDOPERATION_REFUNDSOURCEPACKAGEERROR = "FailedOperation.RefundSourcePackageError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RefundResourcePackage(request *RefundResourcePackageRequest) (response *RefundResourcePackageResponse, err error) {
    return c.RefundResourcePackageWithContext(context.Background(), request)
}

// RefundResourcePackage
// This API is used to refund a resource package.
//
// error code that may be returned:
//  FAILEDOPERATION_REFUNDSOURCEPACKAGEERROR = "FailedOperation.RefundSourcePackageError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RefundResourcePackageWithContext(ctx context.Context, request *RefundResourcePackageRequest) (response *RefundResourcePackageResponse, err error) {
    if request == nil {
        request = NewRefundResourcePackageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "RefundResourcePackage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RefundResourcePackage require credential")
    }

    request.SetContext(ctx)
    
    response = NewRefundResourcePackageResponse()
    err = c.Send(request, response)
    return
}

func NewReloadBalanceProxyNodeRequest() (request *ReloadBalanceProxyNodeRequest) {
    request = &ReloadBalanceProxyNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ReloadBalanceProxyNode")
    
    
    return
}

func NewReloadBalanceProxyNodeResponse() (response *ReloadBalanceProxyNodeResponse) {
    response = &ReloadBalanceProxyNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReloadBalanceProxyNode
// This API is used to reload the database proxy of Cloud Load Balancer.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) ReloadBalanceProxyNode(request *ReloadBalanceProxyNodeRequest) (response *ReloadBalanceProxyNodeResponse, err error) {
    return c.ReloadBalanceProxyNodeWithContext(context.Background(), request)
}

// ReloadBalanceProxyNode
// This API is used to reload the database proxy of Cloud Load Balancer.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) ReloadBalanceProxyNodeWithContext(ctx context.Context, request *ReloadBalanceProxyNodeRequest) (response *ReloadBalanceProxyNodeResponse, err error) {
    if request == nil {
        request = NewReloadBalanceProxyNodeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ReloadBalanceProxyNode")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReloadBalanceProxyNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewReloadBalanceProxyNodeResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveClusterSlaveZoneRequest() (request *RemoveClusterSlaveZoneRequest) {
    request = &RemoveClusterSlaveZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "RemoveClusterSlaveZone")
    
    
    return
}

func NewRemoveClusterSlaveZoneResponse() (response *RemoveClusterSlaveZoneResponse) {
    response = &RemoveClusterSlaveZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveClusterSlaveZone
// This API is used to disable multi-AZ deployment for a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RemoveClusterSlaveZone(request *RemoveClusterSlaveZoneRequest) (response *RemoveClusterSlaveZoneResponse, err error) {
    return c.RemoveClusterSlaveZoneWithContext(context.Background(), request)
}

// RemoveClusterSlaveZone
// This API is used to disable multi-AZ deployment for a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RemoveClusterSlaveZoneWithContext(ctx context.Context, request *RemoveClusterSlaveZoneRequest) (response *RemoveClusterSlaveZoneResponse, err error) {
    if request == nil {
        request = NewRemoveClusterSlaveZoneRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "RemoveClusterSlaveZone")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveClusterSlaveZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveClusterSlaveZoneResponse()
    err = c.Send(request, response)
    return
}

func NewReplayInstanceAuditLogRequest() (request *ReplayInstanceAuditLogRequest) {
    request = &ReplayInstanceAuditLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ReplayInstanceAuditLog")
    
    
    return
}

func NewReplayInstanceAuditLogResponse() (response *ReplayInstanceAuditLogResponse) {
    response = &ReplayInstanceAuditLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReplayInstanceAuditLog
// This API is used to replay instance audit logs.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_AUDITLOGCLOSEDERROR = "OperationDenied.AuditLogClosedError"
//  OPERATIONDENIED_FEATURENOTSUPPORTERROR = "OperationDenied.FeatureNotSupportError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_ISNOTROLLBACKCLUSTERERROR = "OperationDenied.IsNotRollbackClusterError"
//  OPERATIONDENIED_LIMITDAYFORAUDITREPLAYERROR = "OperationDenied.LimitDayForAuditReplayError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ReplayInstanceAuditLog(request *ReplayInstanceAuditLogRequest) (response *ReplayInstanceAuditLogResponse, err error) {
    return c.ReplayInstanceAuditLogWithContext(context.Background(), request)
}

// ReplayInstanceAuditLog
// This API is used to replay instance audit logs.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_AUDITLOGCLOSEDERROR = "OperationDenied.AuditLogClosedError"
//  OPERATIONDENIED_FEATURENOTSUPPORTERROR = "OperationDenied.FeatureNotSupportError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_ISNOTROLLBACKCLUSTERERROR = "OperationDenied.IsNotRollbackClusterError"
//  OPERATIONDENIED_LIMITDAYFORAUDITREPLAYERROR = "OperationDenied.LimitDayForAuditReplayError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ReplayInstanceAuditLogWithContext(ctx context.Context, request *ReplayInstanceAuditLogRequest) (response *ReplayInstanceAuditLogResponse, err error) {
    if request == nil {
        request = NewReplayInstanceAuditLogRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ReplayInstanceAuditLog")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReplayInstanceAuditLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewReplayInstanceAuditLogResponse()
    err = c.Send(request, response)
    return
}

func NewResetAccountPasswordRequest() (request *ResetAccountPasswordRequest) {
    request = &ResetAccountPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ResetAccountPassword")
    
    
    return
}

func NewResetAccountPasswordResponse() (response *ResetAccountPasswordResponse) {
    response = &ResetAccountPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetAccountPassword
// This API is used to modify the database account password.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDVALUEERROR = "InvalidParameterValue.InvalidPasswordValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ResetAccountPassword(request *ResetAccountPasswordRequest) (response *ResetAccountPasswordResponse, err error) {
    return c.ResetAccountPasswordWithContext(context.Background(), request)
}

// ResetAccountPassword
// This API is used to modify the database account password.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDVALUEERROR = "InvalidParameterValue.InvalidPasswordValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ResetAccountPasswordWithContext(ctx context.Context, request *ResetAccountPasswordRequest) (response *ResetAccountPasswordResponse, err error) {
    if request == nil {
        request = NewResetAccountPasswordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ResetAccountPassword")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetAccountPassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetAccountPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewRestartInstanceRequest() (request *RestartInstanceRequest) {
    request = &RestartInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "RestartInstance")
    
    
    return
}

func NewRestartInstanceResponse() (response *RestartInstanceResponse) {
    response = &RestartInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RestartInstance
// This API is used to reboot an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_CLUSTERINSTANCELIMIT = "LimitExceeded.ClusterInstanceLimit"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_NOTREALNAMEACCOUNT = "UnauthorizedOperation.NotRealNameAccount"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RestartInstance(request *RestartInstanceRequest) (response *RestartInstanceResponse, err error) {
    return c.RestartInstanceWithContext(context.Background(), request)
}

// RestartInstance
// This API is used to reboot an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_CLUSTERINSTANCELIMIT = "LimitExceeded.ClusterInstanceLimit"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_NOTREALNAMEACCOUNT = "UnauthorizedOperation.NotRealNameAccount"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RestartInstanceWithContext(ctx context.Context, request *RestartInstanceRequest) (response *RestartInstanceResponse, err error) {
    if request == nil {
        request = NewRestartInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "RestartInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestartInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewResumeServerlessRequest() (request *ResumeServerlessRequest) {
    request = &ResumeServerlessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ResumeServerless")
    
    
    return
}

func NewResumeServerlessResponse() (response *ResumeServerlessResponse) {
    response = &ResumeServerlessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResumeServerless
// This API is used to restore a serverless cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_DBMODENOTSERVERLESSERROR = "InvalidParameterValue.DbModeNotServerlessError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ResumeServerless(request *ResumeServerlessRequest) (response *ResumeServerlessResponse, err error) {
    return c.ResumeServerlessWithContext(context.Background(), request)
}

// ResumeServerless
// This API is used to restore a serverless cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_DBMODENOTSERVERLESSERROR = "InvalidParameterValue.DbModeNotServerlessError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ResumeServerlessWithContext(ctx context.Context, request *ResumeServerlessRequest) (response *ResumeServerlessResponse, err error) {
    if request == nil {
        request = NewResumeServerlessRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ResumeServerless")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResumeServerless require credential")
    }

    request.SetContext(ctx)
    
    response = NewResumeServerlessResponse()
    err = c.Send(request, response)
    return
}

func NewSearchClusterDatabasesRequest() (request *SearchClusterDatabasesRequest) {
    request = &SearchClusterDatabasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "SearchClusterDatabases")
    
    
    return
}

func NewSearchClusterDatabasesResponse() (response *SearchClusterDatabasesResponse) {
    response = &SearchClusterDatabasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SearchClusterDatabases
// This API is used to search cluster database lists.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
func (c *Client) SearchClusterDatabases(request *SearchClusterDatabasesRequest) (response *SearchClusterDatabasesResponse, err error) {
    return c.SearchClusterDatabasesWithContext(context.Background(), request)
}

// SearchClusterDatabases
// This API is used to search cluster database lists.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
func (c *Client) SearchClusterDatabasesWithContext(ctx context.Context, request *SearchClusterDatabasesRequest) (response *SearchClusterDatabasesResponse, err error) {
    if request == nil {
        request = NewSearchClusterDatabasesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "SearchClusterDatabases")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchClusterDatabases require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchClusterDatabasesResponse()
    err = c.Send(request, response)
    return
}

func NewSearchClusterTablesRequest() (request *SearchClusterTablesRequest) {
    request = &SearchClusterTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "SearchClusterTables")
    
    
    return
}

func NewSearchClusterTablesResponse() (response *SearchClusterTablesResponse) {
    response = &SearchClusterTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SearchClusterTables
// This API is used to search cluster data table lists.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
func (c *Client) SearchClusterTables(request *SearchClusterTablesRequest) (response *SearchClusterTablesResponse, err error) {
    return c.SearchClusterTablesWithContext(context.Background(), request)
}

// SearchClusterTables
// This API is used to search cluster data table lists.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
func (c *Client) SearchClusterTablesWithContext(ctx context.Context, request *SearchClusterTablesRequest) (response *SearchClusterTablesResponse, err error) {
    if request == nil {
        request = NewSearchClusterTablesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "SearchClusterTables")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchClusterTables require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchClusterTablesResponse()
    err = c.Send(request, response)
    return
}

func NewSetRenewFlagRequest() (request *SetRenewFlagRequest) {
    request = &SetRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "SetRenewFlag")
    
    
    return
}

func NewSetRenewFlagResponse() (response *SetRenewFlagResponse) {
    response = &SetRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetRenewFlag
// This API is used to set the auto-renewal feature of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) SetRenewFlag(request *SetRenewFlagRequest) (response *SetRenewFlagResponse, err error) {
    return c.SetRenewFlagWithContext(context.Background(), request)
}

// SetRenewFlag
// This API is used to set the auto-renewal feature of an instance.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) SetRenewFlagWithContext(ctx context.Context, request *SetRenewFlagRequest) (response *SetRenewFlagResponse, err error) {
    if request == nil {
        request = NewSetRenewFlagRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "SetRenewFlag")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetRenewFlag require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetRenewFlagResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchClusterVpcRequest() (request *SwitchClusterVpcRequest) {
    request = &SwitchClusterVpcRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "SwitchClusterVpc")
    
    
    return
}

func NewSwitchClusterVpcResponse() (response *SwitchClusterVpcResponse) {
    response = &SwitchClusterVpcResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SwitchClusterVpc
// This API is used to replace the cluster vpc.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) SwitchClusterVpc(request *SwitchClusterVpcRequest) (response *SwitchClusterVpcResponse, err error) {
    return c.SwitchClusterVpcWithContext(context.Background(), request)
}

// SwitchClusterVpc
// This API is used to replace the cluster vpc.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) SwitchClusterVpcWithContext(ctx context.Context, request *SwitchClusterVpcRequest) (response *SwitchClusterVpcResponse, err error) {
    if request == nil {
        request = NewSwitchClusterVpcRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "SwitchClusterVpc")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SwitchClusterVpc require credential")
    }

    request.SetContext(ctx)
    
    response = NewSwitchClusterVpcResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchClusterZoneRequest() (request *SwitchClusterZoneRequest) {
    request = &SwitchClusterZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "SwitchClusterZone")
    
    
    return
}

func NewSwitchClusterZoneResponse() (response *SwitchClusterZoneResponse) {
    response = &SwitchClusterZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SwitchClusterZone
// This API is used to switch the primary and secondary AZs of a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) SwitchClusterZone(request *SwitchClusterZoneRequest) (response *SwitchClusterZoneResponse, err error) {
    return c.SwitchClusterZoneWithContext(context.Background(), request)
}

// SwitchClusterZone
// This API is used to switch the primary and secondary AZs of a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) SwitchClusterZoneWithContext(ctx context.Context, request *SwitchClusterZoneRequest) (response *SwitchClusterZoneResponse, err error) {
    if request == nil {
        request = NewSwitchClusterZoneRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "SwitchClusterZone")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SwitchClusterZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewSwitchClusterZoneResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchProxyVpcRequest() (request *SwitchProxyVpcRequest) {
    request = &SwitchProxyVpcRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "SwitchProxyVpc")
    
    
    return
}

func NewSwitchProxyVpcResponse() (response *SwitchProxyVpcResponse) {
    response = &SwitchProxyVpcResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SwitchProxyVpc
// This API is used to replace the vpc of the database proxy.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) SwitchProxyVpc(request *SwitchProxyVpcRequest) (response *SwitchProxyVpcResponse, err error) {
    return c.SwitchProxyVpcWithContext(context.Background(), request)
}

// SwitchProxyVpc
// This API is used to replace the vpc of the database proxy.
//
// error code that may be returned:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) SwitchProxyVpcWithContext(ctx context.Context, request *SwitchProxyVpcRequest) (response *SwitchProxyVpcResponse, err error) {
    if request == nil {
        request = NewSwitchProxyVpcRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "SwitchProxyVpc")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SwitchProxyVpc require credential")
    }

    request.SetContext(ctx)
    
    response = NewSwitchProxyVpcResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindClusterResourcePackagesRequest() (request *UnbindClusterResourcePackagesRequest) {
    request = &UnbindClusterResourcePackagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "UnbindClusterResourcePackages")
    
    
    return
}

func NewUnbindClusterResourcePackagesResponse() (response *UnbindClusterResourcePackagesResponse) {
    response = &UnbindClusterResourcePackagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnbindClusterResourcePackages
// This API is used to unbind resource packages from clusters.
//
// error code that may be returned:
//  FAILEDOPERATION_QUERYSOURCEPACKAGEERROR = "FailedOperation.QuerySourcePackageError"
//  FAILEDOPERATION_UNBINDSOURCEPACKAGEERROR = "FailedOperation.UnBindSourcePackageError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) UnbindClusterResourcePackages(request *UnbindClusterResourcePackagesRequest) (response *UnbindClusterResourcePackagesResponse, err error) {
    return c.UnbindClusterResourcePackagesWithContext(context.Background(), request)
}

// UnbindClusterResourcePackages
// This API is used to unbind resource packages from clusters.
//
// error code that may be returned:
//  FAILEDOPERATION_QUERYSOURCEPACKAGEERROR = "FailedOperation.QuerySourcePackageError"
//  FAILEDOPERATION_UNBINDSOURCEPACKAGEERROR = "FailedOperation.UnBindSourcePackageError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) UnbindClusterResourcePackagesWithContext(ctx context.Context, request *UnbindClusterResourcePackagesRequest) (response *UnbindClusterResourcePackagesResponse, err error) {
    if request == nil {
        request = NewUnbindClusterResourcePackagesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "UnbindClusterResourcePackages")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindClusterResourcePackages require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindClusterResourcePackagesResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeClusterVersionRequest() (request *UpgradeClusterVersionRequest) {
    request = &UpgradeClusterVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "UpgradeClusterVersion")
    
    
    return
}

func NewUpgradeClusterVersionResponse() (response *UpgradeClusterVersionResponse) {
    response = &UpgradeClusterVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpgradeClusterVersion
// This interface (UpgradeClusterVersion) is used to update the kernel minor version.
//
// error code that may be returned:
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_CLUSTERINSTANCELIMIT = "LimitExceeded.ClusterInstanceLimit"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_NOTREALNAMEACCOUNT = "UnauthorizedOperation.NotRealNameAccount"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) UpgradeClusterVersion(request *UpgradeClusterVersionRequest) (response *UpgradeClusterVersionResponse, err error) {
    return c.UpgradeClusterVersionWithContext(context.Background(), request)
}

// UpgradeClusterVersion
// This interface (UpgradeClusterVersion) is used to update the kernel minor version.
//
// error code that may be returned:
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_CLUSTERINSTANCELIMIT = "LimitExceeded.ClusterInstanceLimit"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_NOTREALNAMEACCOUNT = "UnauthorizedOperation.NotRealNameAccount"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) UpgradeClusterVersionWithContext(ctx context.Context, request *UpgradeClusterVersionRequest) (response *UpgradeClusterVersionResponse, err error) {
    if request == nil {
        request = NewUpgradeClusterVersionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "UpgradeClusterVersion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeClusterVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeClusterVersionResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeInstanceRequest() (request *UpgradeInstanceRequest) {
    request = &UpgradeInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "UpgradeInstance")
    
    
    return
}

func NewUpgradeInstanceResponse() (response *UpgradeInstanceResponse) {
    response = &UpgradeInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpgradeInstance
// This interface (UpgradeInstance) is used to upgrade instances.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEORDER = "FailedOperation.CreateOrder"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_INSUFFICIENTBALANCEERROR = "OperationDenied.InsufficientBalanceError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) UpgradeInstance(request *UpgradeInstanceRequest) (response *UpgradeInstanceResponse, err error) {
    return c.UpgradeInstanceWithContext(context.Background(), request)
}

// UpgradeInstance
// This interface (UpgradeInstance) is used to upgrade instances.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEORDER = "FailedOperation.CreateOrder"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_INSUFFICIENTBALANCEERROR = "OperationDenied.InsufficientBalanceError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) UpgradeInstanceWithContext(ctx context.Context, request *UpgradeInstanceRequest) (response *UpgradeInstanceResponse, err error) {
    if request == nil {
        request = NewUpgradeInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "UpgradeInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeProxyRequest() (request *UpgradeProxyRequest) {
    request = &UpgradeProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "UpgradeProxy")
    
    
    return
}

func NewUpgradeProxyResponse() (response *UpgradeProxyResponse) {
    response = &UpgradeProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpgradeProxy
// This API is used to upgrade database proxy configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_SPECNOTFOUNDERROR = "FailedOperation.SpecNotFoundError"
//  OPERATIONDENIED_GETPROXYGROUPFAILEDERROR = "OperationDenied.GetProxyGroupFailedError"
//  OPERATIONDENIED_PROXYCONNECTCOUNTCHECKERROR = "OperationDenied.ProxyConnectCountCheckError"
//  OPERATIONDENIED_PROXYNODECOUNTCHECKERROR = "OperationDenied.ProxyNodeCountCheckError"
//  OPERATIONDENIED_PROXYNOTRUNNINGERROR = "OperationDenied.ProxyNotRunningError"
//  OPERATIONDENIED_PROXYSALEZONECHECKERROR = "OperationDenied.ProxySaleZoneCheckError"
//  OPERATIONDENIED_PROXYVERSIONCHECKERROR = "OperationDenied.ProxyVersionCheckError"
//  OPERATIONDENIED_PROXYZONECHECKERROR = "OperationDenied.ProxyZoneCheckError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) UpgradeProxy(request *UpgradeProxyRequest) (response *UpgradeProxyResponse, err error) {
    return c.UpgradeProxyWithContext(context.Background(), request)
}

// UpgradeProxy
// This API is used to upgrade database proxy configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_SPECNOTFOUNDERROR = "FailedOperation.SpecNotFoundError"
//  OPERATIONDENIED_GETPROXYGROUPFAILEDERROR = "OperationDenied.GetProxyGroupFailedError"
//  OPERATIONDENIED_PROXYCONNECTCOUNTCHECKERROR = "OperationDenied.ProxyConnectCountCheckError"
//  OPERATIONDENIED_PROXYNODECOUNTCHECKERROR = "OperationDenied.ProxyNodeCountCheckError"
//  OPERATIONDENIED_PROXYNOTRUNNINGERROR = "OperationDenied.ProxyNotRunningError"
//  OPERATIONDENIED_PROXYSALEZONECHECKERROR = "OperationDenied.ProxySaleZoneCheckError"
//  OPERATIONDENIED_PROXYVERSIONCHECKERROR = "OperationDenied.ProxyVersionCheckError"
//  OPERATIONDENIED_PROXYZONECHECKERROR = "OperationDenied.ProxyZoneCheckError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) UpgradeProxyWithContext(ctx context.Context, request *UpgradeProxyRequest) (response *UpgradeProxyResponse, err error) {
    if request == nil {
        request = NewUpgradeProxyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "UpgradeProxy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeProxyResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeProxyVersionRequest() (request *UpgradeProxyVersionRequest) {
    request = &UpgradeProxyVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "UpgradeProxyVersion")
    
    
    return
}

func NewUpgradeProxyVersionResponse() (response *UpgradeProxyVersionResponse) {
    response = &UpgradeProxyVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpgradeProxyVersion
// This API is used to upgrade the database proxy version.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  OPERATIONDENIED_PROXYSTOCKCHECKERROR = "OperationDenied.ProxyStockCheckError"
func (c *Client) UpgradeProxyVersion(request *UpgradeProxyVersionRequest) (response *UpgradeProxyVersionResponse, err error) {
    return c.UpgradeProxyVersionWithContext(context.Background(), request)
}

// UpgradeProxyVersion
// This API is used to upgrade the database proxy version.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  OPERATIONDENIED_PROXYSTOCKCHECKERROR = "OperationDenied.ProxyStockCheckError"
func (c *Client) UpgradeProxyVersionWithContext(ctx context.Context, request *UpgradeProxyVersionRequest) (response *UpgradeProxyVersionResponse, err error) {
    if request == nil {
        request = NewUpgradeProxyVersionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "UpgradeProxyVersion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeProxyVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeProxyVersionResponse()
    err = c.Send(request, response)
    return
}
