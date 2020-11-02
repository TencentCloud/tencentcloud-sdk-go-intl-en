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

package v20190923

import (
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2019-09-23"

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


func NewCreateSecretRequest() (request *CreateSecretRequest) {
    request = &CreateSecretRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "CreateSecret")
    return
}

func NewCreateSecretResponse() (response *CreateSecretResponse) {
    response = &CreateSecretResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create a KMS-encrypted Secret. You can create and store up to 1,000 Secrets in each region.
func (c *Client) CreateSecret(request *CreateSecretRequest) (response *CreateSecretResponse, err error) {
    if request == nil {
        request = NewCreateSecretRequest()
    }
    response = NewCreateSecretResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSecretRequest() (request *DeleteSecretRequest) {
    request = &DeleteSecretRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "DeleteSecret")
    return
}

func NewDeleteSecretResponse() (response *DeleteSecretResponse) {
    response = &DeleteSecretResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete a Secret. You can set whether to delete the Secret immediately or on schedule using the `RecoveryWindowInDays` parameter. For a Secret to be deleted on schedule, its status will be `PendingDelete` before the scheduled deletion time. You can use `RestoreSecret` to restore a deleted Secret during this time. A deleted Secret will not be restorable after the scheduled deletion time. A Secret can only be deleted after being disabled using `DisableSecret`.
func (c *Client) DeleteSecret(request *DeleteSecretRequest) (response *DeleteSecretResponse, err error) {
    if request == nil {
        request = NewDeleteSecretRequest()
    }
    response = NewDeleteSecretResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSecretVersionRequest() (request *DeleteSecretVersionRequest) {
    request = &DeleteSecretVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "DeleteSecretVersion")
    return
}

func NewDeleteSecretVersionResponse() (response *DeleteSecretVersionResponse) {
    response = &DeleteSecretVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete a version of a Secret. The deletion takes effect immediately. Secret versions in any status can be deleted.
func (c *Client) DeleteSecretVersion(request *DeleteSecretVersionRequest) (response *DeleteSecretVersionResponse, err error) {
    if request == nil {
        request = NewDeleteSecretVersionRequest()
    }
    response = NewDeleteSecretVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecretRequest() (request *DescribeSecretRequest) {
    request = &DescribeSecretRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "DescribeSecret")
    return
}

func NewDescribeSecretResponse() (response *DescribeSecretResponse) {
    response = &DescribeSecretResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to obtain the detailed attribute information of a Secret.
func (c *Client) DescribeSecret(request *DescribeSecretRequest) (response *DescribeSecretResponse, err error) {
    if request == nil {
        request = NewDescribeSecretRequest()
    }
    response = NewDescribeSecretResponse()
    err = c.Send(request, response)
    return
}

func NewDisableSecretRequest() (request *DisableSecretRequest) {
    request = &DisableSecretRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "DisableSecret")
    return
}

func NewDisableSecretResponse() (response *DisableSecretResponse) {
    response = &DisableSecretResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to disable a Secret and will change its status to `Disabled`. The plaintext of a disabled Secret cannot be obtained through APIs.
func (c *Client) DisableSecret(request *DisableSecretRequest) (response *DisableSecretResponse, err error) {
    if request == nil {
        request = NewDisableSecretRequest()
    }
    response = NewDisableSecretResponse()
    err = c.Send(request, response)
    return
}

func NewEnableSecretRequest() (request *EnableSecretRequest) {
    request = &EnableSecretRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "EnableSecret")
    return
}

func NewEnableSecretResponse() (response *EnableSecretResponse) {
    response = &EnableSecretResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to enable a Secret and will change its status to `Enabled`. You can call the `GetSecretValue` API to obtain the plaintext of this Secret. Secrets in `PendingDelete` status can only be enabled after being restored by using `RestoreSecret`.
func (c *Client) EnableSecret(request *EnableSecretRequest) (response *EnableSecretResponse, err error) {
    if request == nil {
        request = NewEnableSecretRequest()
    }
    response = NewEnableSecretResponse()
    err = c.Send(request, response)
    return
}

func NewGetRegionsRequest() (request *GetRegionsRequest) {
    request = &GetRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "GetRegions")
    return
}

func NewGetRegionsResponse() (response *GetRegionsResponse) {
    response = &GetRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to obtain the list of regions displayed on Console.
func (c *Client) GetRegions(request *GetRegionsRequest) (response *GetRegionsResponse, err error) {
    if request == nil {
        request = NewGetRegionsRequest()
    }
    response = NewGetRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewGetSecretValueRequest() (request *GetSecretValueRequest) {
    request = &GetSecretValueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "GetSecretValue")
    return
}

func NewGetSecretValueResponse() (response *GetSecretValueResponse) {
    response = &GetSecretValueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to obtain the plaintext of a specified Secret and version. Only plaintext of an enabled Secret can be obtained.
func (c *Client) GetSecretValue(request *GetSecretValueRequest) (response *GetSecretValueResponse, err error) {
    if request == nil {
        request = NewGetSecretValueRequest()
    }
    response = NewGetSecretValueResponse()
    err = c.Send(request, response)
    return
}

func NewGetServiceStatusRequest() (request *GetServiceStatusRequest) {
    request = &GetServiceStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "GetServiceStatus")
    return
}

func NewGetServiceStatusResponse() (response *GetServiceStatusResponse) {
    response = &GetServiceStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to obtain the SecretsManager service status of a user.
func (c *Client) GetServiceStatus(request *GetServiceStatusRequest) (response *GetServiceStatusResponse, err error) {
    if request == nil {
        request = NewGetServiceStatusRequest()
    }
    response = NewGetServiceStatusResponse()
    err = c.Send(request, response)
    return
}

func NewListSecretVersionIdsRequest() (request *ListSecretVersionIdsRequest) {
    request = &ListSecretVersionIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "ListSecretVersionIds")
    return
}

func NewListSecretVersionIdsResponse() (response *ListSecretVersionIdsResponse) {
    response = &ListSecretVersionIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to obtain list of versions of a Secret.
func (c *Client) ListSecretVersionIds(request *ListSecretVersionIdsRequest) (response *ListSecretVersionIdsResponse, err error) {
    if request == nil {
        request = NewListSecretVersionIdsRequest()
    }
    response = NewListSecretVersionIdsResponse()
    err = c.Send(request, response)
    return
}

func NewListSecretsRequest() (request *ListSecretsRequest) {
    request = &ListSecretsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "ListSecrets")
    return
}

func NewListSecretsResponse() (response *ListSecretsResponse) {
    response = &ListSecretsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to obtain the detailed information list of all Secrets. You can specify the filter fields and sorting order as needed.
func (c *Client) ListSecrets(request *ListSecretsRequest) (response *ListSecretsResponse, err error) {
    if request == nil {
        request = NewListSecretsRequest()
    }
    response = NewListSecretsResponse()
    err = c.Send(request, response)
    return
}

func NewPutSecretValueRequest() (request *PutSecretValueRequest) {
    request = &PutSecretValueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "PutSecretValue")
    return
}

func NewPutSecretValueResponse() (response *PutSecretValueResponse) {
    response = &PutSecretValueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to add a new version to a specified Secret. Each Secret supports up to 10 versions. You can only add versions to Secrets in `Enabled` or `Disabled` status.
func (c *Client) PutSecretValue(request *PutSecretValueRequest) (response *PutSecretValueResponse, err error) {
    if request == nil {
        request = NewPutSecretValueRequest()
    }
    response = NewPutSecretValueResponse()
    err = c.Send(request, response)
    return
}

func NewRestoreSecretRequest() (request *RestoreSecretRequest) {
    request = &RestoreSecretRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "RestoreSecret")
    return
}

func NewRestoreSecretResponse() (response *RestoreSecretResponse) {
    response = &RestoreSecretResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to restore a `PendingDelete` Secret, canceling its scheduled deletion. The restored Secret will be in `Disabled` status. You can call the `EnableSecret` API to enable this Secret again.
func (c *Client) RestoreSecret(request *RestoreSecretRequest) (response *RestoreSecretResponse, err error) {
    if request == nil {
        request = NewRestoreSecretRequest()
    }
    response = NewRestoreSecretResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDescriptionRequest() (request *UpdateDescriptionRequest) {
    request = &UpdateDescriptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "UpdateDescription")
    return
}

func NewUpdateDescriptionResponse() (response *UpdateDescriptionResponse) {
    response = &UpdateDescriptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to update the description of a Secret. This API can only update Secrets in `Enabled` or `Disabled` status.
func (c *Client) UpdateDescription(request *UpdateDescriptionRequest) (response *UpdateDescriptionResponse, err error) {
    if request == nil {
        request = NewUpdateDescriptionRequest()
    }
    response = NewUpdateDescriptionResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateSecretRequest() (request *UpdateSecretRequest) {
    request = &UpdateSecretRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "UpdateSecret")
    return
}

func NewUpdateSecretResponse() (response *UpdateSecretResponse) {
    response = &UpdateSecretResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to update the name and version ID of a Secret. Calling this API encrypts the new Secret content and overwrites the old content. This API can only update Secrets in `Enabled` or `Disabled` status.
func (c *Client) UpdateSecret(request *UpdateSecretRequest) (response *UpdateSecretResponse, err error) {
    if request == nil {
        request = NewUpdateSecretRequest()
    }
    response = NewUpdateSecretResponse()
    err = c.Send(request, response)
    return
}
