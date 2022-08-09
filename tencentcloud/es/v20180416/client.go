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

package v20180416

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-04-16"

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


func NewCreateIndexRequest() (request *CreateIndexRequest) {
    request = &CreateIndexRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "CreateIndex")
    
    
    return
}

func NewCreateIndexResponse() (response *CreateIndexResponse) {
    response = &CreateIndexResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateIndex
// This API is used to create indices.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateIndex(request *CreateIndexRequest) (response *CreateIndexResponse, err error) {
    return c.CreateIndexWithContext(context.Background(), request)
}

// CreateIndex
// This API is used to create indices.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateIndexWithContext(ctx context.Context, request *CreateIndexRequest) (response *CreateIndexResponse, err error) {
    if request == nil {
        request = NewCreateIndexRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIndex require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateIndexResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstanceRequest() (request *CreateInstanceRequest) {
    request = &CreateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "CreateInstance")
    
    
    return
}

func NewCreateInstanceResponse() (response *CreateInstanceResponse) {
    response = &CreateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateInstance
// This API is used to create an ES cluster instance with the specified specification.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION_CLUSTERRESOURCELIMITERROR = "FailedOperation.ClusterResourceLimitError"
//  FAILEDOPERATION_DISKCOUNTPARAMERROR = "FailedOperation.DiskCountParamError"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  FAILEDOPERATION_NOTAUTHENTICATED = "FailedOperation.NotAuthenticated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_HIDDENZONE = "ResourceInsufficient.HiddenZone"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
func (c *Client) CreateInstance(request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
    return c.CreateInstanceWithContext(context.Background(), request)
}

// CreateInstance
// This API is used to create an ES cluster instance with the specified specification.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION_CLUSTERRESOURCELIMITERROR = "FailedOperation.ClusterResourceLimitError"
//  FAILEDOPERATION_DISKCOUNTPARAMERROR = "FailedOperation.DiskCountParamError"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  FAILEDOPERATION_NOTAUTHENTICATED = "FailedOperation.NotAuthenticated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_HIDDENZONE = "ResourceInsufficient.HiddenZone"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
func (c *Client) CreateInstanceWithContext(ctx context.Context, request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
    if request == nil {
        request = NewCreateInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteIndexRequest() (request *DeleteIndexRequest) {
    request = &DeleteIndexRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "DeleteIndex")
    
    
    return
}

func NewDeleteIndexResponse() (response *DeleteIndexResponse) {
    response = &DeleteIndexResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteIndex
// This API is used to delete indices.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteIndex(request *DeleteIndexRequest) (response *DeleteIndexResponse, err error) {
    return c.DeleteIndexWithContext(context.Background(), request)
}

// DeleteIndex
// This API is used to delete indices.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteIndexWithContext(ctx context.Context, request *DeleteIndexRequest) (response *DeleteIndexResponse, err error) {
    if request == nil {
        request = NewDeleteIndexRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteIndex require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteIndexResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteInstanceRequest() (request *DeleteInstanceRequest) {
    request = &DeleteInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "DeleteInstance")
    
    
    return
}

func NewDeleteInstanceResponse() (response *DeleteInstanceResponse) {
    response = &DeleteInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteInstance
// This API is used to terminate a cluster instance. 
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteInstance(request *DeleteInstanceRequest) (response *DeleteInstanceResponse, err error) {
    return c.DeleteInstanceWithContext(context.Background(), request)
}

// DeleteInstance
// This API is used to terminate a cluster instance. 
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteInstanceWithContext(ctx context.Context, request *DeleteInstanceRequest) (response *DeleteInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIndexListRequest() (request *DescribeIndexListRequest) {
    request = &DescribeIndexListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "DescribeIndexList")
    
    
    return
}

func NewDescribeIndexListResponse() (response *DescribeIndexListResponse) {
    response = &DescribeIndexListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIndexList
// This API is used to obtain the index list.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeIndexList(request *DescribeIndexListRequest) (response *DescribeIndexListResponse, err error) {
    return c.DescribeIndexListWithContext(context.Background(), request)
}

// DescribeIndexList
// This API is used to obtain the index list.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeIndexListWithContext(ctx context.Context, request *DescribeIndexListRequest) (response *DescribeIndexListResponse, err error) {
    if request == nil {
        request = NewDescribeIndexListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIndexList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIndexListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIndexMetaRequest() (request *DescribeIndexMetaRequest) {
    request = &DescribeIndexMetaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "DescribeIndexMeta")
    
    
    return
}

func NewDescribeIndexMetaResponse() (response *DescribeIndexMetaResponse) {
    response = &DescribeIndexMetaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIndexMeta
// This API is used to obtain index metadata.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeIndexMeta(request *DescribeIndexMetaRequest) (response *DescribeIndexMetaResponse, err error) {
    return c.DescribeIndexMetaWithContext(context.Background(), request)
}

// DescribeIndexMeta
// This API is used to obtain index metadata.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeIndexMetaWithContext(ctx context.Context, request *DescribeIndexMetaRequest) (response *DescribeIndexMetaResponse, err error) {
    if request == nil {
        request = NewDescribeIndexMetaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIndexMeta require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIndexMetaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceLogsRequest() (request *DescribeInstanceLogsRequest) {
    request = &DescribeInstanceLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "DescribeInstanceLogs")
    
    
    return
}

func NewDescribeInstanceLogsResponse() (response *DescribeInstanceLogsResponse) {
    response = &DescribeInstanceLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceLogs
// This API is used to query the eligible ES cluster logs in the current region.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeInstanceLogs(request *DescribeInstanceLogsRequest) (response *DescribeInstanceLogsResponse, err error) {
    return c.DescribeInstanceLogsWithContext(context.Background(), request)
}

// DescribeInstanceLogs
// This API is used to query the eligible ES cluster logs in the current region.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeInstanceLogsWithContext(ctx context.Context, request *DescribeInstanceLogsRequest) (response *DescribeInstanceLogsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceOperationsRequest() (request *DescribeInstanceOperationsRequest) {
    request = &DescribeInstanceOperationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "DescribeInstanceOperations")
    
    
    return
}

func NewDescribeInstanceOperationsResponse() (response *DescribeInstanceOperationsResponse) {
    response = &DescribeInstanceOperationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceOperations
// This API is used to query the operation history of an instance by specified criteria.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeInstanceOperations(request *DescribeInstanceOperationsRequest) (response *DescribeInstanceOperationsResponse, err error) {
    return c.DescribeInstanceOperationsWithContext(context.Background(), request)
}

// DescribeInstanceOperations
// This API is used to query the operation history of an instance by specified criteria.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeInstanceOperationsWithContext(ctx context.Context, request *DescribeInstanceOperationsRequest) (response *DescribeInstanceOperationsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceOperationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceOperations require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceOperationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "DescribeInstances")
    
    
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstances
// This API is used to query all eligible instances in the current region under the current account.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    return c.DescribeInstancesWithContext(context.Background(), request)
}

// DescribeInstances
// This API is used to query all eligible instances in the current region under the current account.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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

func NewDescribeViewsRequest() (request *DescribeViewsRequest) {
    request = &DescribeViewsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "DescribeViews")
    
    
    return
}

func NewDescribeViewsResponse() (response *DescribeViewsResponse) {
    response = &DescribeViewsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeViews
// This API is used to query view data from three dimensions: cluster, node, and Kibana.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeViews(request *DescribeViewsRequest) (response *DescribeViewsResponse, err error) {
    return c.DescribeViewsWithContext(context.Background(), request)
}

// DescribeViews
// This API is used to query view data from three dimensions: cluster, node, and Kibana.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeViewsWithContext(ctx context.Context, request *DescribeViewsRequest) (response *DescribeViewsResponse, err error) {
    if request == nil {
        request = NewDescribeViewsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeViews require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeViewsResponse()
    err = c.Send(request, response)
    return
}

func NewGetRequestTargetNodeTypesRequest() (request *GetRequestTargetNodeTypesRequest) {
    request = &GetRequestTargetNodeTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "GetRequestTargetNodeTypes")
    
    
    return
}

func NewGetRequestTargetNodeTypesResponse() (response *GetRequestTargetNodeTypesResponse) {
    response = &GetRequestTargetNodeTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetRequestTargetNodeTypes
// This API is used to get the node types used to receive client requests.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetRequestTargetNodeTypes(request *GetRequestTargetNodeTypesRequest) (response *GetRequestTargetNodeTypesResponse, err error) {
    return c.GetRequestTargetNodeTypesWithContext(context.Background(), request)
}

// GetRequestTargetNodeTypes
// This API is used to get the node types used to receive client requests.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetRequestTargetNodeTypesWithContext(ctx context.Context, request *GetRequestTargetNodeTypesRequest) (response *GetRequestTargetNodeTypesResponse, err error) {
    if request == nil {
        request = NewGetRequestTargetNodeTypesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetRequestTargetNodeTypes require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetRequestTargetNodeTypesResponse()
    err = c.Send(request, response)
    return
}

func NewRestartInstanceRequest() (request *RestartInstanceRequest) {
    request = &RestartInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "RestartInstance")
    
    
    return
}

func NewRestartInstanceResponse() (response *RestartInstanceResponse) {
    response = &RestartInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RestartInstance
// This API is used to restart an ES cluster instance (for operations such as system update). 
//
// error code that may be returned:
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) RestartInstance(request *RestartInstanceRequest) (response *RestartInstanceResponse, err error) {
    return c.RestartInstanceWithContext(context.Background(), request)
}

// RestartInstance
// This API is used to restart an ES cluster instance (for operations such as system update). 
//
// error code that may be returned:
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) RestartInstanceWithContext(ctx context.Context, request *RestartInstanceRequest) (response *RestartInstanceResponse, err error) {
    if request == nil {
        request = NewRestartInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestartInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewRestartKibanaRequest() (request *RestartKibanaRequest) {
    request = &RestartKibanaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "RestartKibana")
    
    
    return
}

func NewRestartKibanaResponse() (response *RestartKibanaResponse) {
    response = &RestartKibanaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RestartKibana
// This API is used to restart Kibana. 
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) RestartKibana(request *RestartKibanaRequest) (response *RestartKibanaResponse, err error) {
    return c.RestartKibanaWithContext(context.Background(), request)
}

// RestartKibana
// This API is used to restart Kibana. 
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) RestartKibanaWithContext(ctx context.Context, request *RestartKibanaRequest) (response *RestartKibanaResponse, err error) {
    if request == nil {
        request = NewRestartKibanaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartKibana require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestartKibanaResponse()
    err = c.Send(request, response)
    return
}

func NewRestartNodesRequest() (request *RestartNodesRequest) {
    request = &RestartNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "RestartNodes")
    
    
    return
}

func NewRestartNodesResponse() (response *RestartNodesResponse) {
    response = &RestartNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RestartNodes
// This API is used to restart cluster nodes.
//
// error code that may be returned:
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RestartNodes(request *RestartNodesRequest) (response *RestartNodesResponse, err error) {
    return c.RestartNodesWithContext(context.Background(), request)
}

// RestartNodes
// This API is used to restart cluster nodes.
//
// error code that may be returned:
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RestartNodesWithContext(ctx context.Context, request *RestartNodesRequest) (response *RestartNodesResponse, err error) {
    if request == nil {
        request = NewRestartNodesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestartNodesResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDictionariesRequest() (request *UpdateDictionariesRequest) {
    request = &UpdateDictionariesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "UpdateDictionaries")
    
    
    return
}

func NewUpdateDictionariesResponse() (response *UpdateDictionariesResponse) {
    response = &UpdateDictionariesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateDictionaries
// This API is used to update ES cluster dictionaries.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateDictionaries(request *UpdateDictionariesRequest) (response *UpdateDictionariesResponse, err error) {
    return c.UpdateDictionariesWithContext(context.Background(), request)
}

// UpdateDictionaries
// This API is used to update ES cluster dictionaries.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateDictionariesWithContext(ctx context.Context, request *UpdateDictionariesRequest) (response *UpdateDictionariesResponse, err error) {
    if request == nil {
        request = NewUpdateDictionariesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateDictionaries require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateDictionariesResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateIndexRequest() (request *UpdateIndexRequest) {
    request = &UpdateIndexRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "UpdateIndex")
    
    
    return
}

func NewUpdateIndexResponse() (response *UpdateIndexResponse) {
    response = &UpdateIndexResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateIndex
// This API is used to update indices.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateIndex(request *UpdateIndexRequest) (response *UpdateIndexResponse, err error) {
    return c.UpdateIndexWithContext(context.Background(), request)
}

// UpdateIndex
// This API is used to update indices.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateIndexWithContext(ctx context.Context, request *UpdateIndexRequest) (response *UpdateIndexResponse, err error) {
    if request == nil {
        request = NewUpdateIndexRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateIndex require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateIndexResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateInstanceRequest() (request *UpdateInstanceRequest) {
    request = &UpdateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "UpdateInstance")
    
    
    return
}

func NewUpdateInstanceResponse() (response *UpdateInstanceResponse) {
    response = &UpdateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateInstance
// This API is used for operations such as modifying node specification, renaming an instance, modifying configuration, resetting password, and setting Kibana blocklist/allowlist. `InstanceId` is required, while `ForceRestart` is optional. Other parameters or parameter combinations and their meanings are as follows:
//
// - InstanceName: renames an instance (only for instance identification)
//
// - NodeInfoList: modifies node configuration (horizontally scaling nodes, vertically scaling nodes, adding primary nodes, adding cold nodes, etc.)
//
// - EsConfig: modifies cluster configuration
//
// - Password: changes the password of the default user "elastic"
//
// - EsAcl: modifies the ACL
//
// - CosBackUp: sets auto-backup to COS for a cluster
//
// Only one of the parameters or parameter combinations above can be passed in at a time, while passing fewer or more ones will cause the request to fail.
//
// error code that may be returned:
//  FAILEDOPERATION_CLUSTERRESOURCELIMITERROR = "FailedOperation.ClusterResourceLimitError"
//  FAILEDOPERATION_DISKCOUNTPARAMERROR = "FailedOperation.DiskCountParamError"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  FAILEDOPERATION_UNSUPPORTRESETNODETYPEANDSCALEOUTDISK = "FailedOperation.UnsupportResetNodeTypeAndScaleoutDisk"
//  FAILEDOPERATION_UNSUPPORTRESETSCALEDOWNANDMODIFYDISK = "FailedOperation.UnsupportResetScaledownAndModifyDisk"
//  FAILEDOPERATION_UNSUPPORTREVERSEREGULATIONNODETYPEANDDISK = "FailedOperation.UnsupportReverseRegulationNodeTypeAndDisk"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateInstance(request *UpdateInstanceRequest) (response *UpdateInstanceResponse, err error) {
    return c.UpdateInstanceWithContext(context.Background(), request)
}

// UpdateInstance
// This API is used for operations such as modifying node specification, renaming an instance, modifying configuration, resetting password, and setting Kibana blocklist/allowlist. `InstanceId` is required, while `ForceRestart` is optional. Other parameters or parameter combinations and their meanings are as follows:
//
// - InstanceName: renames an instance (only for instance identification)
//
// - NodeInfoList: modifies node configuration (horizontally scaling nodes, vertically scaling nodes, adding primary nodes, adding cold nodes, etc.)
//
// - EsConfig: modifies cluster configuration
//
// - Password: changes the password of the default user "elastic"
//
// - EsAcl: modifies the ACL
//
// - CosBackUp: sets auto-backup to COS for a cluster
//
// Only one of the parameters or parameter combinations above can be passed in at a time, while passing fewer or more ones will cause the request to fail.
//
// error code that may be returned:
//  FAILEDOPERATION_CLUSTERRESOURCELIMITERROR = "FailedOperation.ClusterResourceLimitError"
//  FAILEDOPERATION_DISKCOUNTPARAMERROR = "FailedOperation.DiskCountParamError"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  FAILEDOPERATION_UNSUPPORTRESETNODETYPEANDSCALEOUTDISK = "FailedOperation.UnsupportResetNodeTypeAndScaleoutDisk"
//  FAILEDOPERATION_UNSUPPORTRESETSCALEDOWNANDMODIFYDISK = "FailedOperation.UnsupportResetScaledownAndModifyDisk"
//  FAILEDOPERATION_UNSUPPORTREVERSEREGULATIONNODETYPEANDDISK = "FailedOperation.UnsupportReverseRegulationNodeTypeAndDisk"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateInstanceWithContext(ctx context.Context, request *UpdateInstanceRequest) (response *UpdateInstanceResponse, err error) {
    if request == nil {
        request = NewUpdateInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpdatePluginsRequest() (request *UpdatePluginsRequest) {
    request = &UpdatePluginsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "UpdatePlugins")
    
    
    return
}

func NewUpdatePluginsResponse() (response *UpdatePluginsResponse) {
    response = &UpdatePluginsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdatePlugins
// This API is used to change the list of plugins.
//
// error code that may be returned:
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATENOREPLICATION = "FailedOperation.ErrorClusterStateNoReplication"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdatePlugins(request *UpdatePluginsRequest) (response *UpdatePluginsResponse, err error) {
    return c.UpdatePluginsWithContext(context.Background(), request)
}

// UpdatePlugins
// This API is used to change the list of plugins.
//
// error code that may be returned:
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATENOREPLICATION = "FailedOperation.ErrorClusterStateNoReplication"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdatePluginsWithContext(ctx context.Context, request *UpdatePluginsRequest) (response *UpdatePluginsResponse, err error) {
    if request == nil {
        request = NewUpdatePluginsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdatePlugins require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdatePluginsResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateRequestTargetNodeTypesRequest() (request *UpdateRequestTargetNodeTypesRequest) {
    request = &UpdateRequestTargetNodeTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "UpdateRequestTargetNodeTypes")
    
    
    return
}

func NewUpdateRequestTargetNodeTypesResponse() (response *UpdateRequestTargetNodeTypesResponse) {
    response = &UpdateRequestTargetNodeTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateRequestTargetNodeTypes
// This API is used to update the node types used to receive client requests.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateRequestTargetNodeTypes(request *UpdateRequestTargetNodeTypesRequest) (response *UpdateRequestTargetNodeTypesResponse, err error) {
    return c.UpdateRequestTargetNodeTypesWithContext(context.Background(), request)
}

// UpdateRequestTargetNodeTypes
// This API is used to update the node types used to receive client requests.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateRequestTargetNodeTypesWithContext(ctx context.Context, request *UpdateRequestTargetNodeTypesRequest) (response *UpdateRequestTargetNodeTypesResponse, err error) {
    if request == nil {
        request = NewUpdateRequestTargetNodeTypesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateRequestTargetNodeTypes require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateRequestTargetNodeTypesResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeInstanceRequest() (request *UpgradeInstanceRequest) {
    request = &UpgradeInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "UpgradeInstance")
    
    
    return
}

func NewUpgradeInstanceResponse() (response *UpgradeInstanceResponse) {
    response = &UpgradeInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpgradeInstance
// This API is used to upgrade ES cluster version
//
// error code that may be returned:
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpgradeInstance(request *UpgradeInstanceRequest) (response *UpgradeInstanceResponse, err error) {
    return c.UpgradeInstanceWithContext(context.Background(), request)
}

// UpgradeInstance
// This API is used to upgrade ES cluster version
//
// error code that may be returned:
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpgradeInstanceWithContext(ctx context.Context, request *UpgradeInstanceRequest) (response *UpgradeInstanceResponse, err error) {
    if request == nil {
        request = NewUpgradeInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeLicenseRequest() (request *UpgradeLicenseRequest) {
    request = &UpgradeLicenseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("es", APIVersion, "UpgradeLicense")
    
    
    return
}

func NewUpgradeLicenseResponse() (response *UpgradeLicenseResponse) {
    response = &UpgradeLicenseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpgradeLicense
// This API is used to upgrade ES X-Pack.
//
// error code that may be returned:
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpgradeLicense(request *UpgradeLicenseRequest) (response *UpgradeLicenseResponse, err error) {
    return c.UpgradeLicenseWithContext(context.Background(), request)
}

// UpgradeLicense
// This API is used to upgrade ES X-Pack.
//
// error code that may be returned:
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpgradeLicenseWithContext(ctx context.Context, request *UpgradeLicenseRequest) (response *UpgradeLicenseResponse, err error) {
    if request == nil {
        request = NewUpgradeLicenseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeLicense require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeLicenseResponse()
    err = c.Send(request, response)
    return
}
