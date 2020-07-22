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


func NewCopyFunctionRequest() (request *CopyFunctionRequest) {
    request = &CopyFunctionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "CopyFunction")
    return
}

func NewCopyFunctionResponse() (response *CopyFunctionResponse) {
    response = &CopyFunctionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to replicate a function. You can store the replicated function in a specified Region and Namespace.
// Note: This API **does not** replicate the following objects or attributes of the function:
// 1. Function trigger
// 2. Versions other than $LATEST
// 3. CLS target of the logs configured in the function
// 
// You can manually configure the function after replication as required.
func (c *Client) CopyFunction(request *CopyFunctionRequest) (response *CopyFunctionResponse, err error) {
    if request == nil {
        request = NewCopyFunctionRequest()
    }
    response = NewCopyFunctionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAliasRequest() (request *CreateAliasRequest) {
    request = &CreateAliasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "CreateAlias")
    return
}

func NewCreateAliasResponse() (response *CreateAliasResponse) {
    response = &CreateAliasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create an alias for a function version. You can use the alias to mark a specific function version such as DEV/RELEASE. You can also modify the version pointed to by the alias at any time.
// An alias must point to a main version and can point to an additional version at the same time. If you specify an alias when invoking a function, the request will be sent to the versions pointed to by the alias. You can configure the ratio between the main version and additional version during request sending.
func (c *Client) CreateAlias(request *CreateAliasRequest) (response *CreateAliasResponse, err error) {
    if request == nil {
        request = NewCreateAliasRequest()
    }
    response = NewCreateAliasResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFunctionRequest() (request *CreateFunctionRequest) {
    request = &CreateFunctionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "CreateFunction")
    return
}

func NewCreateFunctionResponse() (response *CreateFunctionResponse) {
    response = &CreateFunctionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create a function based on the input parameters.
func (c *Client) CreateFunction(request *CreateFunctionRequest) (response *CreateFunctionResponse, err error) {
    if request == nil {
        request = NewCreateFunctionRequest()
    }
    response = NewCreateFunctionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNamespaceRequest() (request *CreateNamespaceRequest) {
    request = &CreateNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "CreateNamespace")
    return
}

func NewCreateNamespaceResponse() (response *CreateNamespaceResponse) {
    response = &CreateNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create a namespace based on the input parameters.
func (c *Client) CreateNamespace(request *CreateNamespaceRequest) (response *CreateNamespaceResponse, err error) {
    if request == nil {
        request = NewCreateNamespaceRequest()
    }
    response = NewCreateNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTriggerRequest() (request *CreateTriggerRequest) {
    request = &CreateTriggerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "CreateTrigger")
    return
}

func NewCreateTriggerResponse() (response *CreateTriggerResponse) {
    response = &CreateTriggerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create a trigger based on the input parameters.
func (c *Client) CreateTrigger(request *CreateTriggerRequest) (response *CreateTriggerResponse, err error) {
    if request == nil {
        request = NewCreateTriggerRequest()
    }
    response = NewCreateTriggerResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAliasRequest() (request *DeleteAliasRequest) {
    request = &DeleteAliasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "DeleteAlias")
    return
}

func NewDeleteAliasResponse() (response *DeleteAliasResponse) {
    response = &DeleteAliasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete an alias of a function version.
func (c *Client) DeleteAlias(request *DeleteAliasRequest) (response *DeleteAliasResponse, err error) {
    if request == nil {
        request = NewDeleteAliasRequest()
    }
    response = NewDeleteAliasResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFunctionRequest() (request *DeleteFunctionRequest) {
    request = &DeleteFunctionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "DeleteFunction")
    return
}

func NewDeleteFunctionResponse() (response *DeleteFunctionResponse) {
    response = &DeleteFunctionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete a function based on the input parameters.
func (c *Client) DeleteFunction(request *DeleteFunctionRequest) (response *DeleteFunctionResponse, err error) {
    if request == nil {
        request = NewDeleteFunctionRequest()
    }
    response = NewDeleteFunctionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLayerVersionRequest() (request *DeleteLayerVersionRequest) {
    request = &DeleteLayerVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "DeleteLayerVersion")
    return
}

func NewDeleteLayerVersionResponse() (response *DeleteLayerVersionResponse) {
    response = &DeleteLayerVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete a specified version of a specified layer. The deleted version cannot be associated with a function, but the deletion does not affect functions that are referencing this layer.
func (c *Client) DeleteLayerVersion(request *DeleteLayerVersionRequest) (response *DeleteLayerVersionResponse, err error) {
    if request == nil {
        request = NewDeleteLayerVersionRequest()
    }
    response = NewDeleteLayerVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNamespaceRequest() (request *DeleteNamespaceRequest) {
    request = &DeleteNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "DeleteNamespace")
    return
}

func NewDeleteNamespaceResponse() (response *DeleteNamespaceResponse) {
    response = &DeleteNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create a namespace based on the input parameters.
func (c *Client) DeleteNamespace(request *DeleteNamespaceRequest) (response *DeleteNamespaceResponse, err error) {
    if request == nil {
        request = NewDeleteNamespaceRequest()
    }
    response = NewDeleteNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTriggerRequest() (request *DeleteTriggerRequest) {
    request = &DeleteTriggerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "DeleteTrigger")
    return
}

func NewDeleteTriggerResponse() (response *DeleteTriggerResponse) {
    response = &DeleteTriggerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete an existing trigger based on the input parameters.
func (c *Client) DeleteTrigger(request *DeleteTriggerRequest) (response *DeleteTriggerResponse, err error) {
    if request == nil {
        request = NewDeleteTriggerRequest()
    }
    response = NewDeleteTriggerResponse()
    err = c.Send(request, response)
    return
}

func NewGetAliasRequest() (request *GetAliasRequest) {
    request = &GetAliasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "GetAlias")
    return
}

func NewGetAliasResponse() (response *GetAliasResponse) {
    response = &GetAliasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the alias details such as the name, description, version, and routing information.
func (c *Client) GetAlias(request *GetAliasRequest) (response *GetAliasResponse, err error) {
    if request == nil {
        request = NewGetAliasRequest()
    }
    response = NewGetAliasResponse()
    err = c.Send(request, response)
    return
}

func NewGetFunctionRequest() (request *GetFunctionRequest) {
    request = &GetFunctionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "GetFunction")
    return
}

func NewGetFunctionResponse() (response *GetFunctionResponse) {
    response = &GetFunctionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to obtain function details, such as name, code, handler, associated trigger, and timeout.
func (c *Client) GetFunction(request *GetFunctionRequest) (response *GetFunctionResponse, err error) {
    if request == nil {
        request = NewGetFunctionRequest()
    }
    response = NewGetFunctionResponse()
    err = c.Send(request, response)
    return
}

func NewGetFunctionAddressRequest() (request *GetFunctionAddressRequest) {
    request = &GetFunctionAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "GetFunctionAddress")
    return
}

func NewGetFunctionAddressResponse() (response *GetFunctionAddressResponse) {
    response = &GetFunctionAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to obtain the download address of the function code package.
func (c *Client) GetFunctionAddress(request *GetFunctionAddressRequest) (response *GetFunctionAddressResponse, err error) {
    if request == nil {
        request = NewGetFunctionAddressRequest()
    }
    response = NewGetFunctionAddressResponse()
    err = c.Send(request, response)
    return
}

func NewGetFunctionLogsRequest() (request *GetFunctionLogsRequest) {
    request = &GetFunctionLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "GetFunctionLogs")
    return
}

func NewGetFunctionLogsResponse() (response *GetFunctionLogsResponse) {
    response = &GetFunctionLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to return function running logs according to the specified log query criteria.
func (c *Client) GetFunctionLogs(request *GetFunctionLogsRequest) (response *GetFunctionLogsResponse, err error) {
    if request == nil {
        request = NewGetFunctionLogsRequest()
    }
    response = NewGetFunctionLogsResponse()
    err = c.Send(request, response)
    return
}

func NewGetLayerVersionRequest() (request *GetLayerVersionRequest) {
    request = &GetLayerVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "GetLayerVersion")
    return
}

func NewGetLayerVersionResponse() (response *GetLayerVersionResponse) {
    response = &GetLayerVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the layer version details, including links used to download files in the layer.
func (c *Client) GetLayerVersion(request *GetLayerVersionRequest) (response *GetLayerVersionResponse, err error) {
    if request == nil {
        request = NewGetLayerVersionRequest()
    }
    response = NewGetLayerVersionResponse()
    err = c.Send(request, response)
    return
}

func NewInvokeRequest() (request *InvokeRequest) {
    request = &InvokeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "Invoke")
    return
}

func NewInvokeResponse() (response *InvokeResponse) {
    response = &InvokeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to run a function.
func (c *Client) Invoke(request *InvokeRequest) (response *InvokeResponse, err error) {
    if request == nil {
        request = NewInvokeRequest()
    }
    response = NewInvokeResponse()
    err = c.Send(request, response)
    return
}

func NewListAliasesRequest() (request *ListAliasesRequest) {
    request = &ListAliasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "ListAliases")
    return
}

func NewListAliasesResponse() (response *ListAliasesResponse) {
    response = &ListAliasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to return the list of all aliases under a function. You can filter them by the specific function version.
func (c *Client) ListAliases(request *ListAliasesRequest) (response *ListAliasesResponse, err error) {
    if request == nil {
        request = NewListAliasesRequest()
    }
    response = NewListAliasesResponse()
    err = c.Send(request, response)
    return
}

func NewListFunctionsRequest() (request *ListFunctionsRequest) {
    request = &ListFunctionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "ListFunctions")
    return
}

func NewListFunctionsResponse() (response *ListFunctionsResponse) {
    response = &ListFunctionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to return relevant function information based on the input query parameters.
func (c *Client) ListFunctions(request *ListFunctionsRequest) (response *ListFunctionsResponse, err error) {
    if request == nil {
        request = NewListFunctionsRequest()
    }
    response = NewListFunctionsResponse()
    err = c.Send(request, response)
    return
}

func NewListLayerVersionsRequest() (request *ListLayerVersionsRequest) {
    request = &ListLayerVersionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "ListLayerVersions")
    return
}

func NewListLayerVersionsResponse() (response *ListLayerVersionsResponse) {
    response = &ListLayerVersionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the information of all versions of a specified layer.
func (c *Client) ListLayerVersions(request *ListLayerVersionsRequest) (response *ListLayerVersionsResponse, err error) {
    if request == nil {
        request = NewListLayerVersionsRequest()
    }
    response = NewListLayerVersionsResponse()
    err = c.Send(request, response)
    return
}

func NewListLayersRequest() (request *ListLayersRequest) {
    request = &ListLayersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "ListLayers")
    return
}

func NewListLayersResponse() (response *ListLayersResponse) {
    response = &ListLayersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to return the list of all layers, including the information of the latest version of each layer. You can filter them by the compatible runtime.
func (c *Client) ListLayers(request *ListLayersRequest) (response *ListLayersResponse, err error) {
    if request == nil {
        request = NewListLayersRequest()
    }
    response = NewListLayersResponse()
    err = c.Send(request, response)
    return
}

func NewListNamespacesRequest() (request *ListNamespacesRequest) {
    request = &ListNamespacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "ListNamespaces")
    return
}

func NewListNamespacesResponse() (response *ListNamespacesResponse) {
    response = &ListNamespacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to display a namespace list.
func (c *Client) ListNamespaces(request *ListNamespacesRequest) (response *ListNamespacesResponse, err error) {
    if request == nil {
        request = NewListNamespacesRequest()
    }
    response = NewListNamespacesResponse()
    err = c.Send(request, response)
    return
}

func NewListTriggersRequest() (request *ListTriggersRequest) {
    request = &ListTriggersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "ListTriggers")
    return
}

func NewListTriggersResponse() (response *ListTriggersResponse) {
    response = &ListTriggersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the function trigger list.
func (c *Client) ListTriggers(request *ListTriggersRequest) (response *ListTriggersResponse, err error) {
    if request == nil {
        request = NewListTriggersRequest()
    }
    response = NewListTriggersResponse()
    err = c.Send(request, response)
    return
}

func NewListVersionByFunctionRequest() (request *ListVersionByFunctionRequest) {
    request = &ListVersionByFunctionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "ListVersionByFunction")
    return
}

func NewListVersionByFunctionResponse() (response *ListVersionByFunctionResponse) {
    response = &ListVersionByFunctionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the function version based on the input parameters.
func (c *Client) ListVersionByFunction(request *ListVersionByFunctionRequest) (response *ListVersionByFunctionResponse, err error) {
    if request == nil {
        request = NewListVersionByFunctionRequest()
    }
    response = NewListVersionByFunctionResponse()
    err = c.Send(request, response)
    return
}

func NewPublishLayerVersionRequest() (request *PublishLayerVersionRequest) {
    request = &PublishLayerVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "PublishLayerVersion")
    return
}

func NewPublishLayerVersionResponse() (response *PublishLayerVersionResponse) {
    response = &PublishLayerVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create a version for a layer by using the given .zip file or COS object. Each time this API is called with the same layer name, a new version will be generated.
func (c *Client) PublishLayerVersion(request *PublishLayerVersionRequest) (response *PublishLayerVersionResponse, err error) {
    if request == nil {
        request = NewPublishLayerVersionRequest()
    }
    response = NewPublishLayerVersionResponse()
    err = c.Send(request, response)
    return
}

func NewPublishVersionRequest() (request *PublishVersionRequest) {
    request = &PublishVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "PublishVersion")
    return
}

func NewPublishVersionResponse() (response *PublishVersionResponse) {
    response = &PublishVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used for users to release a new version of the function.
func (c *Client) PublishVersion(request *PublishVersionRequest) (response *PublishVersionResponse, err error) {
    if request == nil {
        request = NewPublishVersionRequest()
    }
    response = NewPublishVersionResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAliasRequest() (request *UpdateAliasRequest) {
    request = &UpdateAliasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "UpdateAlias")
    return
}

func NewUpdateAliasResponse() (response *UpdateAliasResponse) {
    response = &UpdateAliasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to update the configuration of an alias.
func (c *Client) UpdateAlias(request *UpdateAliasRequest) (response *UpdateAliasResponse, err error) {
    if request == nil {
        request = NewUpdateAliasRequest()
    }
    response = NewUpdateAliasResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateFunctionCodeRequest() (request *UpdateFunctionCodeRequest) {
    request = &UpdateFunctionCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "UpdateFunctionCode")
    return
}

func NewUpdateFunctionCodeResponse() (response *UpdateFunctionCodeResponse) {
    response = &UpdateFunctionCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to update the function code based on the input parameters.
func (c *Client) UpdateFunctionCode(request *UpdateFunctionCodeRequest) (response *UpdateFunctionCodeResponse, err error) {
    if request == nil {
        request = NewUpdateFunctionCodeRequest()
    }
    response = NewUpdateFunctionCodeResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateFunctionConfigurationRequest() (request *UpdateFunctionConfigurationRequest) {
    request = &UpdateFunctionConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "UpdateFunctionConfiguration")
    return
}

func NewUpdateFunctionConfigurationResponse() (response *UpdateFunctionConfigurationResponse) {
    response = &UpdateFunctionConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to update the function configuration based on the input parameters.
func (c *Client) UpdateFunctionConfiguration(request *UpdateFunctionConfigurationRequest) (response *UpdateFunctionConfigurationResponse, err error) {
    if request == nil {
        request = NewUpdateFunctionConfigurationRequest()
    }
    response = NewUpdateFunctionConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateNamespaceRequest() (request *UpdateNamespaceRequest) {
    request = &UpdateNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "UpdateNamespace")
    return
}

func NewUpdateNamespaceResponse() (response *UpdateNamespaceResponse) {
    response = &UpdateNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to update a namespace.
func (c *Client) UpdateNamespace(request *UpdateNamespaceRequest) (response *UpdateNamespaceResponse, err error) {
    if request == nil {
        request = NewUpdateNamespaceRequest()
    }
    response = NewUpdateNamespaceResponse()
    err = c.Send(request, response)
    return
}
