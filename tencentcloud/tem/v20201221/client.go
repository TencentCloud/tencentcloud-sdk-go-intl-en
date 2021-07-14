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

package v20201221

import (
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2020-12-21"

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


func NewCreateCosTokenV2Request() (request *CreateCosTokenV2Request) {
    request = &CreateCosTokenV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "CreateCosTokenV2")
    return
}

func NewCreateCosTokenV2Response() (response *CreateCosTokenV2Response) {
    response = &CreateCosTokenV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCosTokenV2
// This API is used to generate a COS temporary key.
func (c *Client) CreateCosTokenV2(request *CreateCosTokenV2Request) (response *CreateCosTokenV2Response, err error) {
    if request == nil {
        request = NewCreateCosTokenV2Request()
    }
    response = NewCreateCosTokenV2Response()
    err = c.Send(request, response)
    return
}

func NewCreateNamespaceRequest() (request *CreateNamespaceRequest) {
    request = &CreateNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "CreateNamespace")
    return
}

func NewCreateNamespaceResponse() (response *CreateNamespaceResponse) {
    response = &CreateNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateNamespace
// This API is used to create an environment.
func (c *Client) CreateNamespace(request *CreateNamespaceRequest) (response *CreateNamespaceResponse, err error) {
    if request == nil {
        request = NewCreateNamespaceRequest()
    }
    response = NewCreateNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateResourceRequest() (request *CreateResourceRequest) {
    request = &CreateResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "CreateResource")
    return
}

func NewCreateResourceResponse() (response *CreateResourceResponse) {
    response = &CreateResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateResource
// This API is used to bind a cloud resource.
func (c *Client) CreateResource(request *CreateResourceRequest) (response *CreateResourceResponse, err error) {
    if request == nil {
        request = NewCreateResourceRequest()
    }
    response = NewCreateResourceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateServiceV2Request() (request *CreateServiceV2Request) {
    request = &CreateServiceV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "CreateServiceV2")
    return
}

func NewCreateServiceV2Response() (response *CreateServiceV2Response) {
    response = &CreateServiceV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateServiceV2
// This API is used to create a service.
//
// error code that may be returned:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) CreateServiceV2(request *CreateServiceV2Request) (response *CreateServiceV2Response, err error) {
    if request == nil {
        request = NewCreateServiceV2Request()
    }
    response = NewCreateServiceV2Response()
    err = c.Send(request, response)
    return
}

func NewDeleteIngressRequest() (request *DeleteIngressRequest) {
    request = &DeleteIngressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "DeleteIngress")
    return
}

func NewDeleteIngressResponse() (response *DeleteIngressResponse) {
    response = &DeleteIngressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteIngress
// This API is used to delete an ingress rule.
//
// error code that may be returned:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) DeleteIngress(request *DeleteIngressRequest) (response *DeleteIngressResponse, err error) {
    if request == nil {
        request = NewDeleteIngressRequest()
    }
    response = NewDeleteIngressResponse()
    err = c.Send(request, response)
    return
}

func NewDeployServiceV2Request() (request *DeployServiceV2Request) {
    request = &DeployServiceV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "DeployServiceV2")
    return
}

func NewDeployServiceV2Response() (response *DeployServiceV2Response) {
    response = &DeployServiceV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeployServiceV2
// This API is used to deploy a service.
//
// error code that may be returned:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) DeployServiceV2(request *DeployServiceV2Request) (response *DeployServiceV2Response, err error) {
    if request == nil {
        request = NewDeployServiceV2Request()
    }
    response = NewDeployServiceV2Response()
    err = c.Send(request, response)
    return
}

func NewDescribeIngressRequest() (request *DescribeIngressRequest) {
    request = &DescribeIngressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "DescribeIngress")
    return
}

func NewDescribeIngressResponse() (response *DescribeIngressResponse) {
    response = &DescribeIngressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIngress
// This API is used to query an ingress rule.
//
// error code that may be returned:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) DescribeIngress(request *DescribeIngressRequest) (response *DescribeIngressResponse, err error) {
    if request == nil {
        request = NewDescribeIngressRequest()
    }
    response = NewDescribeIngressResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIngressesRequest() (request *DescribeIngressesRequest) {
    request = &DescribeIngressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "DescribeIngresses")
    return
}

func NewDescribeIngressesResponse() (response *DescribeIngressesResponse) {
    response = &DescribeIngressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIngresses
// This API is used to query the list of ingress rules.
//
// error code that may be returned:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) DescribeIngresses(request *DescribeIngressesRequest) (response *DescribeIngressesResponse, err error) {
    if request == nil {
        request = NewDescribeIngressesRequest()
    }
    response = NewDescribeIngressesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNamespacesRequest() (request *DescribeNamespacesRequest) {
    request = &DescribeNamespacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "DescribeNamespaces")
    return
}

func NewDescribeNamespacesResponse() (response *DescribeNamespacesResponse) {
    response = &DescribeNamespacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNamespaces
// This API is used to get the list of tenant environments.
//
// error code that may be returned:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) DescribeNamespaces(request *DescribeNamespacesRequest) (response *DescribeNamespacesResponse, err error) {
    if request == nil {
        request = NewDescribeNamespacesRequest()
    }
    response = NewDescribeNamespacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRelatedIngressesRequest() (request *DescribeRelatedIngressesRequest) {
    request = &DescribeRelatedIngressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "DescribeRelatedIngresses")
    return
}

func NewDescribeRelatedIngressesResponse() (response *DescribeRelatedIngressesResponse) {
    response = &DescribeRelatedIngressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRelatedIngresses
// This API is used to query the list of ingress rules associated with the service.
//
// error code that may be returned:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) DescribeRelatedIngresses(request *DescribeRelatedIngressesRequest) (response *DescribeRelatedIngressesResponse, err error) {
    if request == nil {
        request = NewDescribeRelatedIngressesRequest()
    }
    response = NewDescribeRelatedIngressesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceRunPodListV2Request() (request *DescribeServiceRunPodListV2Request) {
    request = &DescribeServiceRunPodListV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "DescribeServiceRunPodListV2")
    return
}

func NewDescribeServiceRunPodListV2Response() (response *DescribeServiceRunPodListV2Response) {
    response = &DescribeServiceRunPodListV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeServiceRunPodListV2
// This API is used to get the list of running pods under a service.
//
// error code that may be returned:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) DescribeServiceRunPodListV2(request *DescribeServiceRunPodListV2Request) (response *DescribeServiceRunPodListV2Response, err error) {
    if request == nil {
        request = NewDescribeServiceRunPodListV2Request()
    }
    response = NewDescribeServiceRunPodListV2Response()
    err = c.Send(request, response)
    return
}

func NewGenerateDownloadUrlRequest() (request *GenerateDownloadUrlRequest) {
    request = &GenerateDownloadUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "GenerateDownloadUrl")
    return
}

func NewGenerateDownloadUrlResponse() (response *GenerateDownloadUrlResponse) {
    response = &GenerateDownloadUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GenerateDownloadUrl
// Generate the pre-signed download URL for the specified package
//
// error code that may be returned:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) GenerateDownloadUrl(request *GenerateDownloadUrlRequest) (response *GenerateDownloadUrlResponse, err error) {
    if request == nil {
        request = NewGenerateDownloadUrlRequest()
    }
    response = NewGenerateDownloadUrlResponse()
    err = c.Send(request, response)
    return
}

func NewModifyIngressRequest() (request *ModifyIngressRequest) {
    request = &ModifyIngressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "ModifyIngress")
    return
}

func NewModifyIngressResponse() (response *ModifyIngressResponse) {
    response = &ModifyIngressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyIngress
// This API is used to create or update an ingress rule.
//
// error code that may be returned:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) ModifyIngress(request *ModifyIngressRequest) (response *ModifyIngressResponse, err error) {
    if request == nil {
        request = NewModifyIngressRequest()
    }
    response = NewModifyIngressResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNamespaceRequest() (request *ModifyNamespaceRequest) {
    request = &ModifyNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "ModifyNamespace")
    return
}

func NewModifyNamespaceResponse() (response *ModifyNamespaceResponse) {
    response = &ModifyNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyNamespace
// This API is used to edit an environment.
//
// error code that may be returned:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) ModifyNamespace(request *ModifyNamespaceRequest) (response *ModifyNamespaceResponse, err error) {
    if request == nil {
        request = NewModifyNamespaceRequest()
    }
    response = NewModifyNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyServiceInfoRequest() (request *ModifyServiceInfoRequest) {
    request = &ModifyServiceInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "ModifyServiceInfo")
    return
}

func NewModifyServiceInfoResponse() (response *ModifyServiceInfoResponse) {
    response = &ModifyServiceInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyServiceInfo
// This API is used to modify a service’s basic information.
//
// error code that may be returned:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) ModifyServiceInfo(request *ModifyServiceInfoRequest) (response *ModifyServiceInfoResponse, err error) {
    if request == nil {
        request = NewModifyServiceInfoRequest()
    }
    response = NewModifyServiceInfoResponse()
    err = c.Send(request, response)
    return
}

func NewRestartServiceRunPodRequest() (request *RestartServiceRunPodRequest) {
    request = &RestartServiceRunPodRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "RestartServiceRunPod")
    return
}

func NewRestartServiceRunPodResponse() (response *RestartServiceRunPodResponse) {
    response = &RestartServiceRunPodResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RestartServiceRunPod
// This API is used to restart an instance.
//
// error code that may be returned:
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
func (c *Client) RestartServiceRunPod(request *RestartServiceRunPodRequest) (response *RestartServiceRunPodResponse, err error) {
    if request == nil {
        request = NewRestartServiceRunPodRequest()
    }
    response = NewRestartServiceRunPodResponse()
    err = c.Send(request, response)
    return
}
