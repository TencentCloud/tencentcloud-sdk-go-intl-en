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
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
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

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
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
//  FAILEDOPERATION_CLUSTERRESOURCELIMITERROR = "FailedOperation.ClusterResourceLimitError"
//  FAILEDOPERATION_DISKCOUNTPARAMERROR = "FailedOperation.DiskCountParamError"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  FAILEDOPERATION_NOTAUTHENTICATED = "FailedOperation.NotAuthenticated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
func (c *Client) CreateInstance(request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
    if request == nil {
        request = NewCreateInstanceRequest()
    }
    response = NewCreateInstanceResponse()
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
    if request == nil {
        request = NewDeleteInstanceRequest()
    }
    response = NewDeleteInstanceResponse()
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
    if request == nil {
        request = NewDescribeInstanceLogsRequest()
    }
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
    if request == nil {
        request = NewDescribeInstanceOperationsRequest()
    }
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
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    response = NewDescribeInstancesResponse()
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
    if request == nil {
        request = NewGetRequestTargetNodeTypesRequest()
    }
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) RestartInstance(request *RestartInstanceRequest) (response *RestartInstanceResponse, err error) {
    if request == nil {
        request = NewRestartInstanceRequest()
    }
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
    if request == nil {
        request = NewRestartKibanaRequest()
    }
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
    if request == nil {
        request = NewRestartNodesRequest()
    }
    response = NewRestartNodesResponse()
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
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  FAILEDOPERATION_UNSUPPORTREVERSEREGULATIONNODETYPEANDDISK = "FailedOperation.UnsupportReverseRegulationNodeTypeAndDisk"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateInstance(request *UpdateInstanceRequest) (response *UpdateInstanceResponse, err error) {
    if request == nil {
        request = NewUpdateInstanceRequest()
    }
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
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdatePlugins(request *UpdatePluginsRequest) (response *UpdatePluginsResponse, err error) {
    if request == nil {
        request = NewUpdatePluginsRequest()
    }
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
    if request == nil {
        request = NewUpdateRequestTargetNodeTypesRequest()
    }
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
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpgradeInstance(request *UpgradeInstanceRequest) (response *UpgradeInstanceResponse, err error) {
    if request == nil {
        request = NewUpgradeInstanceRequest()
    }
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
    if request == nil {
        request = NewUpgradeLicenseRequest()
    }
    response = NewUpgradeLicenseResponse()
    err = c.Send(request, response)
    return
}
