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

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewAllocateHostsRequest() (request *AllocateHostsRequest) {
    request = &AllocateHostsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "AllocateHosts")
    return
}

func NewAllocateHostsResponse() (response *AllocateHostsResponse) {
    response = &AllocateHostsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create CDH instances with specified configuration.
// * When HostChargeType is PREPAID, the HostChargePrepaid parameter must be specified.
func (c *Client) AllocateHosts(request *AllocateHostsRequest) (response *AllocateHostsResponse, err error) {
    if request == nil {
        request = NewAllocateHostsRequest()
    }
    response = NewAllocateHostsResponse()
    err = c.Send(request, response)
    return
}

func NewAssociateInstancesKeyPairsRequest() (request *AssociateInstancesKeyPairsRequest) {
    request = &AssociateInstancesKeyPairsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "AssociateInstancesKeyPairs")
    return
}

func NewAssociateInstancesKeyPairsResponse() (response *AssociateInstancesKeyPairsResponse) {
    response = &AssociateInstancesKeyPairsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to associate key pairs with instances.
// 
// * If the public key of a key pair is written to the `SSH` configuration of the instance, users will be able to log in to the instance with the private key of the key pair.
// * If the instance is already associated with a key, the old key will become invalid.
// If you currently use a password to log in, you will no longer be able to do so after you associate the instance with a key.
// * Batch operations are supported. The maximum number of instances in each request is 100. If any instance in the request cannot be associated with a key, you will get an error code.
func (c *Client) AssociateInstancesKeyPairs(request *AssociateInstancesKeyPairsRequest) (response *AssociateInstancesKeyPairsResponse, err error) {
    if request == nil {
        request = NewAssociateInstancesKeyPairsRequest()
    }
    response = NewAssociateInstancesKeyPairsResponse()
    err = c.Send(request, response)
    return
}

func NewAssociateSecurityGroupsRequest() (request *AssociateSecurityGroupsRequest) {
    request = &AssociateSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "AssociateSecurityGroups")
    return
}

func NewAssociateSecurityGroupsResponse() (response *AssociateSecurityGroupsResponse) {
    response = &AssociateSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to associate security groups with specified instances.
func (c *Client) AssociateSecurityGroups(request *AssociateSecurityGroupsRequest) (response *AssociateSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewAssociateSecurityGroupsRequest()
    }
    response = NewAssociateSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDisasterRecoverGroupRequest() (request *CreateDisasterRecoverGroupRequest) {
    request = &CreateDisasterRecoverGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "CreateDisasterRecoverGroup")
    return
}

func NewCreateDisasterRecoverGroupResponse() (response *CreateDisasterRecoverGroupResponse) {
    response = &CreateDisasterRecoverGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create a [spread placement group](https://intl.cloud.tencent.com/document/product/213/15486?from_cn_redirect=1). After you create one, you can specify it for an instance when you [create the instance](https://intl.cloud.tencent.com/document/api/213/15730?from_cn_redirect=1), 
func (c *Client) CreateDisasterRecoverGroup(request *CreateDisasterRecoverGroupRequest) (response *CreateDisasterRecoverGroupResponse, err error) {
    if request == nil {
        request = NewCreateDisasterRecoverGroupRequest()
    }
    response = NewCreateDisasterRecoverGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateImageRequest() (request *CreateImageRequest) {
    request = &CreateImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "CreateImage")
    return
}

func NewCreateImageResponse() (response *CreateImageResponse) {
    response = &CreateImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create a new image with the system disk of an instance. The image can be used to create new instances.
func (c *Client) CreateImage(request *CreateImageRequest) (response *CreateImageResponse, err error) {
    if request == nil {
        request = NewCreateImageRequest()
    }
    response = NewCreateImageResponse()
    err = c.Send(request, response)
    return
}

func NewCreateKeyPairRequest() (request *CreateKeyPairRequest) {
    request = &CreateKeyPairRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "CreateKeyPair")
    return
}

func NewCreateKeyPairResponse() (response *CreateKeyPairResponse) {
    response = &CreateKeyPairResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create an `OpenSSH RSA` key pair, which you can use to log in to a `Linux` instance.
// 
// * You only need to specify a name, and the system will automatically create a key pair and return its `ID` and the public and private keys.
// * The name of the key pair must be unique.
// * You can save the private key to a file and use it as an authentication method for `SSH`.
// * Tencent Cloud does not save users' private keys. Be sure to save it yourself.
func (c *Client) CreateKeyPair(request *CreateKeyPairRequest) (response *CreateKeyPairResponse, err error) {
    if request == nil {
        request = NewCreateKeyPairRequest()
    }
    response = NewCreateKeyPairResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDisasterRecoverGroupsRequest() (request *DeleteDisasterRecoverGroupsRequest) {
    request = &DeleteDisasterRecoverGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DeleteDisasterRecoverGroups")
    return
}

func NewDeleteDisasterRecoverGroupsResponse() (response *DeleteDisasterRecoverGroupsResponse) {
    response = &DeleteDisasterRecoverGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete a [spread placement group](https://intl.cloud.tencent.com/document/product/213/15486?from_cn_redirect=1). Only empty placement groups can be deleted. To delete a non-empty group, you need to terminate all the CVM instances in it first. Otherwise, the deletion will fail.
func (c *Client) DeleteDisasterRecoverGroups(request *DeleteDisasterRecoverGroupsRequest) (response *DeleteDisasterRecoverGroupsResponse, err error) {
    if request == nil {
        request = NewDeleteDisasterRecoverGroupsRequest()
    }
    response = NewDeleteDisasterRecoverGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteImagesRequest() (request *DeleteImagesRequest) {
    request = &DeleteImagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DeleteImages")
    return
}

func NewDeleteImagesResponse() (response *DeleteImagesResponse) {
    response = &DeleteImagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete images.
// 
// * If the [ImageState](https://intl.cloud.tencent.com/document/api/213/9452?from_cn_redirect=1#image_state) of an image is `Creating` or `In Use`, it cannot be deleted. Use [DescribeImages](https://intl.cloud.tencent.com/document/api/213/9418?from_cn_redirect=1) to query the image state.
// * You can only create up to 10 custom images in each region. If you have used up the quota, you can delete images to create new ones.
// * A shared image cannot be deleted.
func (c *Client) DeleteImages(request *DeleteImagesRequest) (response *DeleteImagesResponse, err error) {
    if request == nil {
        request = NewDeleteImagesRequest()
    }
    response = NewDeleteImagesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteKeyPairsRequest() (request *DeleteKeyPairsRequest) {
    request = &DeleteKeyPairsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DeleteKeyPairs")
    return
}

func NewDeleteKeyPairsResponse() (response *DeleteKeyPairsResponse) {
    response = &DeleteKeyPairsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete the key pairs hosted in Tencent Cloud.
// 
// * You can delete multiple key pairs at the same time.
// * A key pair used by an instance or image cannot be deleted. Therefore, you need to verify whether all the key pairs have been deleted successfully.
func (c *Client) DeleteKeyPairs(request *DeleteKeyPairsRequest) (response *DeleteKeyPairsResponse, err error) {
    if request == nil {
        request = NewDeleteKeyPairsRequest()
    }
    response = NewDeleteKeyPairsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDisasterRecoverGroupQuotaRequest() (request *DescribeDisasterRecoverGroupQuotaRequest) {
    request = &DescribeDisasterRecoverGroupQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeDisasterRecoverGroupQuota")
    return
}

func NewDescribeDisasterRecoverGroupQuotaResponse() (response *DescribeDisasterRecoverGroupQuotaResponse) {
    response = &DescribeDisasterRecoverGroupQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the quota of [spread placement groups](https://intl.cloud.tencent.com/document/product/213/15486?from_cn_redirect=1).
func (c *Client) DescribeDisasterRecoverGroupQuota(request *DescribeDisasterRecoverGroupQuotaRequest) (response *DescribeDisasterRecoverGroupQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeDisasterRecoverGroupQuotaRequest()
    }
    response = NewDescribeDisasterRecoverGroupQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDisasterRecoverGroupsRequest() (request *DescribeDisasterRecoverGroupsRequest) {
    request = &DescribeDisasterRecoverGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeDisasterRecoverGroups")
    return
}

func NewDescribeDisasterRecoverGroupsResponse() (response *DescribeDisasterRecoverGroupsResponse) {
    response = &DescribeDisasterRecoverGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the information on [spread placement groups](https://intl.cloud.tencent.com/document/product/213/15486?from_cn_redirect=1).
func (c *Client) DescribeDisasterRecoverGroups(request *DescribeDisasterRecoverGroupsRequest) (response *DescribeDisasterRecoverGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeDisasterRecoverGroupsRequest()
    }
    response = NewDescribeDisasterRecoverGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostsRequest() (request *DescribeHostsRequest) {
    request = &DescribeHostsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeHosts")
    return
}

func NewDescribeHostsResponse() (response *DescribeHostsResponse) {
    response = &DescribeHostsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the details of CDH instances.
func (c *Client) DescribeHosts(request *DescribeHostsRequest) (response *DescribeHostsResponse, err error) {
    if request == nil {
        request = NewDescribeHostsRequest()
    }
    response = NewDescribeHostsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageQuotaRequest() (request *DescribeImageQuotaRequest) {
    request = &DescribeImageQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeImageQuota")
    return
}

func NewDescribeImageQuotaResponse() (response *DescribeImageQuotaResponse) {
    response = &DescribeImageQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the image quota of an user account.
func (c *Client) DescribeImageQuota(request *DescribeImageQuotaRequest) (response *DescribeImageQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeImageQuotaRequest()
    }
    response = NewDescribeImageQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageSharePermissionRequest() (request *DescribeImageSharePermissionRequest) {
    request = &DescribeImageSharePermissionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeImageSharePermission")
    return
}

func NewDescribeImageSharePermissionResponse() (response *DescribeImageSharePermissionResponse) {
    response = &DescribeImageSharePermissionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query image sharing information.
func (c *Client) DescribeImageSharePermission(request *DescribeImageSharePermissionRequest) (response *DescribeImageSharePermissionResponse, err error) {
    if request == nil {
        request = NewDescribeImageSharePermissionRequest()
    }
    response = NewDescribeImageSharePermissionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImagesRequest() (request *DescribeImagesRequest) {
    request = &DescribeImagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeImages")
    return
}

func NewDescribeImagesResponse() (response *DescribeImagesResponse) {
    response = &DescribeImagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to view the list of images.
// 
// * You specify the image ID or set filters to query the details of certain images.
// * You can specify `Offset` and `Limit` to select a certain part of the results. By default, the information on the first 20 matching results is returned.
func (c *Client) DescribeImages(request *DescribeImagesRequest) (response *DescribeImagesResponse, err error) {
    if request == nil {
        request = NewDescribeImagesRequest()
    }
    response = NewDescribeImagesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImportImageOsRequest() (request *DescribeImportImageOsRequest) {
    request = &DescribeImportImageOsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeImportImageOs")
    return
}

func NewDescribeImportImageOsResponse() (response *DescribeImportImageOsResponse) {
    response = &DescribeImportImageOsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the list of supported operating systems of imported images.
func (c *Client) DescribeImportImageOs(request *DescribeImportImageOsRequest) (response *DescribeImportImageOsResponse, err error) {
    if request == nil {
        request = NewDescribeImportImageOsRequest()
    }
    response = NewDescribeImportImageOsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceFamilyConfigsRequest() (request *DescribeInstanceFamilyConfigsRequest) {
    request = &DescribeInstanceFamilyConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstanceFamilyConfigs")
    return
}

func NewDescribeInstanceFamilyConfigsResponse() (response *DescribeInstanceFamilyConfigsResponse) {
    response = &DescribeInstanceFamilyConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query a list of model families available to the current user in the current region.
func (c *Client) DescribeInstanceFamilyConfigs(request *DescribeInstanceFamilyConfigsRequest) (response *DescribeInstanceFamilyConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceFamilyConfigsRequest()
    }
    response = NewDescribeInstanceFamilyConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceTypeConfigsRequest() (request *DescribeInstanceTypeConfigsRequest) {
    request = &DescribeInstanceTypeConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstanceTypeConfigs")
    return
}

func NewDescribeInstanceTypeConfigsResponse() (response *DescribeInstanceTypeConfigsResponse) {
    response = &DescribeInstanceTypeConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the model configuration of an instance.
// 
// * You can filter the query results with `zone` or `instance-family`. For more information on filtering conditions, see [`Filter`](https://intl.cloud.tencent.com/document/api/213/15753?from_cn_redirect=1#Filter).
// * If no parameter is defined, the model configuration of all the instances in the specified region will be returned.
func (c *Client) DescribeInstanceTypeConfigs(request *DescribeInstanceTypeConfigsRequest) (response *DescribeInstanceTypeConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceTypeConfigsRequest()
    }
    response = NewDescribeInstanceTypeConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceVncUrlRequest() (request *DescribeInstanceVncUrlRequest) {
    request = &DescribeInstanceVncUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstanceVncUrl")
    return
}

func NewDescribeInstanceVncUrlResponse() (response *DescribeInstanceVncUrlResponse) {
    response = &DescribeInstanceVncUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the Virtual Network Console (VNC) URL of an instance for its login to the VNC.
// 
// * It does not support `STOPPED` CVMs.
// * A VNC URL is only valid for 15 seconds. If you do not access the URL within 15 seconds, it will become invalid and you have to query a URL again.
// * Once the VNC URL is accessed, it will become invalid and you have to query a URL again if needed.
// * If the connection is interrupted, you can make up to 30 reconnection attempts per minute.
// * After getting the value `InstanceVncUrl`, you need to append `InstanceVncUrl=xxxx` to the end of the link <https://img.qcloud.com/qcloud/app/active_vnc/index.html?>.
// 
//   - `InstanceVncUrl`: its value will be returned after the API is successfully called.
// 
//     The final URL is in the following format:
// 
// ```
// https://img.qcloud.com/qcloud/app/active_vnc/index.html?InstanceVncUrl=wss%3A%2F%2Fbjvnc.qcloud.com%3A26789%2Fvnc%3Fs%3DaHpjWnRVMFNhYmxKdDM5MjRHNlVTSVQwajNUSW0wb2tBbmFtREFCTmFrcy8vUUNPMG0wSHZNOUUxRm5PMmUzWmFDcWlOdDJIbUJxSTZDL0RXcHZxYnZZMmRkWWZWcEZia2lyb09XMzdKNmM9
// ```
func (c *Client) DescribeInstanceVncUrl(request *DescribeInstanceVncUrlRequest) (response *DescribeInstanceVncUrlResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceVncUrlRequest()
    }
    response = NewDescribeInstanceVncUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstances")
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the details of instances.
// 
// * You can filter the query results with the instance `ID`, name, or billing method. See `Filter` for more information.
// * If no parameter is defined, a certain number of instances under the current account will be returned. The number is specified by `Limit` and is 20 by default.
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesOperationLimitRequest() (request *DescribeInstancesOperationLimitRequest) {
    request = &DescribeInstancesOperationLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstancesOperationLimit")
    return
}

func NewDescribeInstancesOperationLimitResponse() (response *DescribeInstancesOperationLimitResponse) {
    response = &DescribeInstancesOperationLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query limitations on operations on an instance.
// 
// * Currently you can use this API to query the maximum number of times you can modify the configuration of an instance.
func (c *Client) DescribeInstancesOperationLimit(request *DescribeInstancesOperationLimitRequest) (response *DescribeInstancesOperationLimitResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesOperationLimitRequest()
    }
    response = NewDescribeInstancesOperationLimitResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesStatusRequest() (request *DescribeInstancesStatusRequest) {
    request = &DescribeInstancesStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstancesStatus")
    return
}

func NewDescribeInstancesStatusResponse() (response *DescribeInstancesStatusResponse) {
    response = &DescribeInstancesStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the status of instances.
// 
// * You can query the status of an instance with its `ID`.
// * If no parameter is defined, the status of a certain number of instances under the current account will be returned. The number is specified by `Limit` and is 20 by default.
func (c *Client) DescribeInstancesStatus(request *DescribeInstancesStatusRequest) (response *DescribeInstancesStatusResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesStatusRequest()
    }
    response = NewDescribeInstancesStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInternetChargeTypeConfigsRequest() (request *DescribeInternetChargeTypeConfigsRequest) {
    request = &DescribeInternetChargeTypeConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeInternetChargeTypeConfigs")
    return
}

func NewDescribeInternetChargeTypeConfigsResponse() (response *DescribeInternetChargeTypeConfigsResponse) {
    response = &DescribeInternetChargeTypeConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the network billing methods.
func (c *Client) DescribeInternetChargeTypeConfigs(request *DescribeInternetChargeTypeConfigsRequest) (response *DescribeInternetChargeTypeConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeInternetChargeTypeConfigsRequest()
    }
    response = NewDescribeInternetChargeTypeConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKeyPairsRequest() (request *DescribeKeyPairsRequest) {
    request = &DescribeKeyPairsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeKeyPairs")
    return
}

func NewDescribeKeyPairsResponse() (response *DescribeKeyPairsResponse) {
    response = &DescribeKeyPairsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query key pairs.
// 
// * A key pair is a pair of keys generated by an algorithm in which the public key is available to the public and the private key is available only to the user. You can use this API to query the public key but not the private key.
func (c *Client) DescribeKeyPairs(request *DescribeKeyPairsRequest) (response *DescribeKeyPairsResponse, err error) {
    if request == nil {
        request = NewDescribeKeyPairsRequest()
    }
    response = NewDescribeKeyPairsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegionsRequest() (request *DescribeRegionsRequest) {
    request = &DescribeRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeRegions")
    return
}

func NewDescribeRegionsResponse() (response *DescribeRegionsResponse) {
    response = &DescribeRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query regions.
func (c *Client) DescribeRegions(request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeRegionsRequest()
    }
    response = NewDescribeRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReservedInstancesRequest() (request *DescribeReservedInstancesRequest) {
    request = &DescribeReservedInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeReservedInstances")
    return
}

func NewDescribeReservedInstancesResponse() (response *DescribeReservedInstancesResponse) {
    response = &DescribeReservedInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to list reserved instances the user has purchased.
func (c *Client) DescribeReservedInstances(request *DescribeReservedInstancesRequest) (response *DescribeReservedInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeReservedInstancesRequest()
    }
    response = NewDescribeReservedInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReservedInstancesOfferingsRequest() (request *DescribeReservedInstancesOfferingsRequest) {
    request = &DescribeReservedInstancesOfferingsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeReservedInstancesOfferings")
    return
}

func NewDescribeReservedInstancesOfferingsResponse() (response *DescribeReservedInstancesOfferingsResponse) {
    response = &DescribeReservedInstancesOfferingsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to describe Reserved Instance offerings that are available for purchase.
func (c *Client) DescribeReservedInstancesOfferings(request *DescribeReservedInstancesOfferingsRequest) (response *DescribeReservedInstancesOfferingsResponse, err error) {
    if request == nil {
        request = NewDescribeReservedInstancesOfferingsRequest()
    }
    response = NewDescribeReservedInstancesOfferingsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZoneInstanceConfigInfosRequest() (request *DescribeZoneInstanceConfigInfosRequest) {
    request = &DescribeZoneInstanceConfigInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeZoneInstanceConfigInfos")
    return
}

func NewDescribeZoneInstanceConfigInfosResponse() (response *DescribeZoneInstanceConfigInfosResponse) {
    response = &DescribeZoneInstanceConfigInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the configurations of models in an availability zone.
func (c *Client) DescribeZoneInstanceConfigInfos(request *DescribeZoneInstanceConfigInfosRequest) (response *DescribeZoneInstanceConfigInfosResponse, err error) {
    if request == nil {
        request = NewDescribeZoneInstanceConfigInfosRequest()
    }
    response = NewDescribeZoneInstanceConfigInfosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZonesRequest() (request *DescribeZonesRequest) {
    request = &DescribeZonesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeZones")
    return
}

func NewDescribeZonesResponse() (response *DescribeZonesResponse) {
    response = &DescribeZonesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query availability zones.
func (c *Client) DescribeZones(request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    if request == nil {
        request = NewDescribeZonesRequest()
    }
    response = NewDescribeZonesResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateInstancesKeyPairsRequest() (request *DisassociateInstancesKeyPairsRequest) {
    request = &DisassociateInstancesKeyPairsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DisassociateInstancesKeyPairs")
    return
}

func NewDisassociateInstancesKeyPairsResponse() (response *DisassociateInstancesKeyPairsResponse) {
    response = &DisassociateInstancesKeyPairsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to unbind one or more key pairs from one or more instances.
// 
// * It only supports [`STOPPED`](https://intl.cloud.tencent.com/document/product/213/15753?from_cn_redirect=1#InstanceStatus) Linux instances.
// * After a key pair is disassociated from an instance, you can log in to the instance with password.
// * If you did not set a password for the instance, you will not be able to log in via SSH after the unbinding. In this case, you can call [ResetInstancesPassword](https://intl.cloud.tencent.com/document/api/213/15736?from_cn_redirect=1) to set a login password.
// * Batch operations are supported. The maximum number of instances in each request is 100. If instances not available for the operation are selected, you will get an error code.
func (c *Client) DisassociateInstancesKeyPairs(request *DisassociateInstancesKeyPairsRequest) (response *DisassociateInstancesKeyPairsResponse, err error) {
    if request == nil {
        request = NewDisassociateInstancesKeyPairsRequest()
    }
    response = NewDisassociateInstancesKeyPairsResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateSecurityGroupsRequest() (request *DisassociateSecurityGroupsRequest) {
    request = &DisassociateSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DisassociateSecurityGroups")
    return
}

func NewDisassociateSecurityGroupsResponse() (response *DisassociateSecurityGroupsResponse) {
    response = &DisassociateSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to disassociate security groups from instances.
func (c *Client) DisassociateSecurityGroups(request *DisassociateSecurityGroupsRequest) (response *DisassociateSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDisassociateSecurityGroupsRequest()
    }
    response = NewDisassociateSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewImportImageRequest() (request *ImportImageRequest) {
    request = &ImportImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "ImportImage")
    return
}

func NewImportImageResponse() (response *ImportImageResponse) {
    response = &ImportImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to import images. Imported images can be used to create instances. 
func (c *Client) ImportImage(request *ImportImageRequest) (response *ImportImageResponse, err error) {
    if request == nil {
        request = NewImportImageRequest()
    }
    response = NewImportImageResponse()
    err = c.Send(request, response)
    return
}

func NewImportKeyPairRequest() (request *ImportKeyPairRequest) {
    request = &ImportKeyPairRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "ImportKeyPair")
    return
}

func NewImportKeyPairResponse() (response *ImportKeyPairResponse) {
    response = &ImportKeyPairResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to import key pairs.
// 
// * You can use this API to import key pairs to a user account, but the key pairs will not be automatically associated with any instance. You may use [AssociasteInstancesKeyPair](https://intl.cloud.tencent.com/document/api/213/9404?from_cn_redirect=1) to associate key pairs with instances.
// * You need to specify the names of the key pairs and the content of the public keys.
// * If you only have private keys, you can convert them to public keys with the `SSL` tool before importing them.
func (c *Client) ImportKeyPair(request *ImportKeyPairRequest) (response *ImportKeyPairResponse, err error) {
    if request == nil {
        request = NewImportKeyPairRequest()
    }
    response = NewImportKeyPairResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceResetInstanceRequest() (request *InquiryPriceResetInstanceRequest) {
    request = &InquiryPriceResetInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "InquiryPriceResetInstance")
    return
}

func NewInquiryPriceResetInstanceResponse() (response *InquiryPriceResetInstanceResponse) {
    response = &InquiryPriceResetInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the price for reinstalling an instance.
// 
// * If you have specified the `ImageId` parameter, the price query is performed with the specified image. Otherwise, the image used by the current instance is used.
// * You can only query the price for reinstallation caused by switching between Linux and Windows OS. And the [system disk type](https://intl.cloud.tencent.com/document/api/213/15753?from_cn_redirect=1#SystemDisk) of the instance must be `CLOUD_BASIC`, `CLOUD_PREMIUM`, or `CLOUD_SSD`.
// * Currently, this API only supports instances in Mainland China regions.
func (c *Client) InquiryPriceResetInstance(request *InquiryPriceResetInstanceRequest) (response *InquiryPriceResetInstanceResponse, err error) {
    if request == nil {
        request = NewInquiryPriceResetInstanceRequest()
    }
    response = NewInquiryPriceResetInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceResetInstancesInternetMaxBandwidthRequest() (request *InquiryPriceResetInstancesInternetMaxBandwidthRequest) {
    request = &InquiryPriceResetInstancesInternetMaxBandwidthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "InquiryPriceResetInstancesInternetMaxBandwidth")
    return
}

func NewInquiryPriceResetInstancesInternetMaxBandwidthResponse() (response *InquiryPriceResetInstancesInternetMaxBandwidthResponse) {
    response = &InquiryPriceResetInstancesInternetMaxBandwidthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the price for upgrading the public bandwidth cap of an instance.
// 
// * The allowed bandwidth cap varies for different models. For details, see [Purchasing Network Bandwidth](https://intl.cloud.tencent.com/document/product/213/509?from_cn_redirect=1).
// * For bandwidth billed by the `TRAFFIC_POSTPAID_BY_HOUR` method, changing the bandwidth cap through this API takes effect in real time. You can increase or reduce bandwidth within applicable limits.
func (c *Client) InquiryPriceResetInstancesInternetMaxBandwidth(request *InquiryPriceResetInstancesInternetMaxBandwidthRequest) (response *InquiryPriceResetInstancesInternetMaxBandwidthResponse, err error) {
    if request == nil {
        request = NewInquiryPriceResetInstancesInternetMaxBandwidthRequest()
    }
    response = NewInquiryPriceResetInstancesInternetMaxBandwidthResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceResetInstancesTypeRequest() (request *InquiryPriceResetInstancesTypeRequest) {
    request = &InquiryPriceResetInstancesTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "InquiryPriceResetInstancesType")
    return
}

func NewInquiryPriceResetInstancesTypeResponse() (response *InquiryPriceResetInstancesTypeResponse) {
    response = &InquiryPriceResetInstancesTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the price for adjusting the instance model.
// 
// * Currently, you can only use this API to query the prices of instances whose [system disk type](https://intl.cloud.tencent.com/document/api/213/9452?from_cn_redirect=1#block_device) is `CLOUD_BASIC`, `CLOUD_PREMIUM`, or `CLOUD_SSD`.
// * Currently, you cannot use this API to query the prices of [CDH](https://intl.cloud.tencent.com/document/product/416?from_cn_redirect=1) instances.
func (c *Client) InquiryPriceResetInstancesType(request *InquiryPriceResetInstancesTypeRequest) (response *InquiryPriceResetInstancesTypeResponse, err error) {
    if request == nil {
        request = NewInquiryPriceResetInstancesTypeRequest()
    }
    response = NewInquiryPriceResetInstancesTypeResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceResizeInstanceDisksRequest() (request *InquiryPriceResizeInstanceDisksRequest) {
    request = &InquiryPriceResizeInstanceDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "InquiryPriceResizeInstanceDisks")
    return
}

func NewInquiryPriceResizeInstanceDisksResponse() (response *InquiryPriceResizeInstanceDisksResponse) {
    response = &InquiryPriceResizeInstanceDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the price for expanding data disks of an instance.
// 
// * Currently, you can only use this API to query the price of non-elastic data disks whose [disk type](https://intl.cloud.tencent.com/document/api/213/9452?from_cn_redirect=1#block_device) is `CLOUD_BASIC`, `CLOUD_PREMIUM`, or `CLOUD_SSD`. You can use [`DescribeDisks`](https://intl.cloud.tencent.com/document/api/362/16315?from_cn_redirect=1) to check whether a disk is elastic. If the `Portable` field in the response is `false`, it means that the disk is non-elastic.
// * Currently, you cannot use this API to query the price for [CDH](https://intl.cloud.tencent.com/document/product/416?from_cn_redirect=1) instances. *Also, you can only query the price of expanding one data disk at a time.
func (c *Client) InquiryPriceResizeInstanceDisks(request *InquiryPriceResizeInstanceDisksRequest) (response *InquiryPriceResizeInstanceDisksResponse, err error) {
    if request == nil {
        request = NewInquiryPriceResizeInstanceDisksRequest()
    }
    response = NewInquiryPriceResizeInstanceDisksResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceRunInstancesRequest() (request *InquiryPriceRunInstancesRequest) {
    request = &InquiryPriceRunInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "InquiryPriceRunInstances")
    return
}

func NewInquiryPriceRunInstancesResponse() (response *InquiryPriceRunInstancesResponse) {
    response = &InquiryPriceRunInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the price of creating instances. You can only use this API for instances whose configuration is within the purchase limit. For more information, see [RunInstances](https://intl.cloud.tencent.com/document/api/213/15730?from_cn_redirect=1).
func (c *Client) InquiryPriceRunInstances(request *InquiryPriceRunInstancesRequest) (response *InquiryPriceRunInstancesResponse, err error) {
    if request == nil {
        request = NewInquiryPriceRunInstancesRequest()
    }
    response = NewInquiryPriceRunInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDisasterRecoverGroupAttributeRequest() (request *ModifyDisasterRecoverGroupAttributeRequest) {
    request = &ModifyDisasterRecoverGroupAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "ModifyDisasterRecoverGroupAttribute")
    return
}

func NewModifyDisasterRecoverGroupAttributeResponse() (response *ModifyDisasterRecoverGroupAttributeResponse) {
    response = &ModifyDisasterRecoverGroupAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify the attributes of [spread placement groups](https://intl.cloud.tencent.com/document/product/213/15486?from_cn_redirect=1).
func (c *Client) ModifyDisasterRecoverGroupAttribute(request *ModifyDisasterRecoverGroupAttributeRequest) (response *ModifyDisasterRecoverGroupAttributeResponse, err error) {
    if request == nil {
        request = NewModifyDisasterRecoverGroupAttributeRequest()
    }
    response = NewModifyDisasterRecoverGroupAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyHostsAttributeRequest() (request *ModifyHostsAttributeRequest) {
    request = &ModifyHostsAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "ModifyHostsAttribute")
    return
}

func NewModifyHostsAttributeResponse() (response *ModifyHostsAttributeResponse) {
    response = &ModifyHostsAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify the attributes of a CDH instance, such as instance name and renewal flag. One of the two parameters, HostName and RenewFlag, must be set, but you cannot set both of them at the same time.
func (c *Client) ModifyHostsAttribute(request *ModifyHostsAttributeRequest) (response *ModifyHostsAttributeResponse, err error) {
    if request == nil {
        request = NewModifyHostsAttributeRequest()
    }
    response = NewModifyHostsAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyImageAttributeRequest() (request *ModifyImageAttributeRequest) {
    request = &ModifyImageAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "ModifyImageAttribute")
    return
}

func NewModifyImageAttributeResponse() (response *ModifyImageAttributeResponse) {
    response = &ModifyImageAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify image attributes.
// 
// * Attributes of shared images cannot be modified.
func (c *Client) ModifyImageAttribute(request *ModifyImageAttributeRequest) (response *ModifyImageAttributeResponse, err error) {
    if request == nil {
        request = NewModifyImageAttributeRequest()
    }
    response = NewModifyImageAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyImageSharePermissionRequest() (request *ModifyImageSharePermissionRequest) {
    request = &ModifyImageSharePermissionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "ModifyImageSharePermission")
    return
}

func NewModifyImageSharePermissionResponse() (response *ModifyImageSharePermissionResponse) {
    response = &ModifyImageSharePermissionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify image sharing information.
// 
// * The accounts with which an image is shared can use the shared image to create instances.
// * Each custom image can be shared with up to 50 accounts.
// * You can use a shared image to create instances, but you cannot change its name and description.
// * If an image is shared with another account, the shared image will be in the same region as the original image.
func (c *Client) ModifyImageSharePermission(request *ModifyImageSharePermissionRequest) (response *ModifyImageSharePermissionResponse, err error) {
    if request == nil {
        request = NewModifyImageSharePermissionRequest()
    }
    response = NewModifyImageSharePermissionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstancesAttributeRequest() (request *ModifyInstancesAttributeRequest) {
    request = &ModifyInstancesAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "ModifyInstancesAttribute")
    return
}

func NewModifyInstancesAttributeResponse() (response *ModifyInstancesAttributeResponse) {
    response = &ModifyInstancesAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify the attributes of an instance. Currently you can only use the API to modify the name and the associated security groups of the instance.
// 
// * Instance names are used only for users' convenience. Tencent Cloud does not use the name for ticket submission or instance management.
// * Batch operations are supported. The maximum number of instances in each request is 100.
// * When you change the security groups associated with an instance, the original security groups will be disassociated.
func (c *Client) ModifyInstancesAttribute(request *ModifyInstancesAttributeRequest) (response *ModifyInstancesAttributeResponse, err error) {
    if request == nil {
        request = NewModifyInstancesAttributeRequest()
    }
    response = NewModifyInstancesAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstancesProjectRequest() (request *ModifyInstancesProjectRequest) {
    request = &ModifyInstancesProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "ModifyInstancesProject")
    return
}

func NewModifyInstancesProjectResponse() (response *ModifyInstancesProjectResponse) {
    response = &ModifyInstancesProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to change the project to which an instance belongs.
// 
// * Project is a virtual concept. You can create multiple projects under one account, manage different resources in each project, and assign different instances to different projects. You may use the [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) API to query instances and use the project ID to filter results.
// * You cannot modify the project of an instance that is bound to a load balancer. You need to firstly unbind the load balancer from the instance by using the [`DeregisterInstancesFromLoadBalancer`](https://intl.cloud.tencent.com/document/api/214/1258?from_cn_redirect=1) API.
// [^_^]: # (If you modify the project of an instance, security groups associated with the instance will be automatically disassociated. You can use the [`ModifyInstancesAttribute`](https://intl.cloud.tencent.com/document/api/213/15739?from_cn_redirect=1) API to associate the instance with the security groups again.
// * Batch operations are supported. You can operate up to 100 instances in each request.
// * You can call the [DescribeInstances](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1#.E7.A4.BA.E4.BE.8B3-.E6.9F.A5.E8.AF.A2.E5.AE.9E.E4.BE.8B.E7.9A.84.E6.9C.80.E6.96.B0.E6.93.8D.E4.BD.9C.E6.83.85.E5.86.B5) API and find the result of the operation in the response parameter `LatestOperationState`. If the value is `SUCCESS`, the operation is successful.
func (c *Client) ModifyInstancesProject(request *ModifyInstancesProjectRequest) (response *ModifyInstancesProjectResponse, err error) {
    if request == nil {
        request = NewModifyInstancesProjectRequest()
    }
    response = NewModifyInstancesProjectResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstancesVpcAttributeRequest() (request *ModifyInstancesVpcAttributeRequest) {
    request = &ModifyInstancesVpcAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "ModifyInstancesVpcAttribute")
    return
}

func NewModifyInstancesVpcAttributeResponse() (response *ModifyInstancesVpcAttributeResponse) {
    response = &ModifyInstancesVpcAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify the VPC attributes of an instance, such as the VPC IP address.
// * By default, the instances will shut down when you perform this operation and restart upon completion.
// * If the specified VPC ID and subnet ID (the subnet must be in the same availability zone as the instance) are different from the VPC where the specified instance resides, the instance will be migrated to a subnet of the specified VPC. Before performing this operation, make sure that the specified instance is not bound with an [ENI](https://intl.cloud.tencent.com/document/product/576?from_cn_redirect=1) or [CLB](https://intl.cloud.tencent.com/document/product/214?from_cn_redirect=1).
func (c *Client) ModifyInstancesVpcAttribute(request *ModifyInstancesVpcAttributeRequest) (response *ModifyInstancesVpcAttributeResponse, err error) {
    if request == nil {
        request = NewModifyInstancesVpcAttributeRequest()
    }
    response = NewModifyInstancesVpcAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyKeyPairAttributeRequest() (request *ModifyKeyPairAttributeRequest) {
    request = &ModifyKeyPairAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "ModifyKeyPairAttribute")
    return
}

func NewModifyKeyPairAttributeResponse() (response *ModifyKeyPairAttributeResponse) {
    response = &ModifyKeyPairAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify the attributes of key pairs.
// 
// * This API modifies the name and description of the key pair identified by the key pair ID.
// * The name of the key pair must be unique.
// * Key pair ID is the unique identifier of a key pair and cannot be modified.
func (c *Client) ModifyKeyPairAttribute(request *ModifyKeyPairAttributeRequest) (response *ModifyKeyPairAttributeResponse, err error) {
    if request == nil {
        request = NewModifyKeyPairAttributeRequest()
    }
    response = NewModifyKeyPairAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewPurchaseReservedInstancesOfferingRequest() (request *PurchaseReservedInstancesOfferingRequest) {
    request = &PurchaseReservedInstancesOfferingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "PurchaseReservedInstancesOffering")
    return
}

func NewPurchaseReservedInstancesOfferingResponse() (response *PurchaseReservedInstancesOfferingResponse) {
    response = &PurchaseReservedInstancesOfferingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to purchase one or more specific Reserved Instances.
func (c *Client) PurchaseReservedInstancesOffering(request *PurchaseReservedInstancesOfferingRequest) (response *PurchaseReservedInstancesOfferingResponse, err error) {
    if request == nil {
        request = NewPurchaseReservedInstancesOfferingRequest()
    }
    response = NewPurchaseReservedInstancesOfferingResponse()
    err = c.Send(request, response)
    return
}

func NewRebootInstancesRequest() (request *RebootInstancesRequest) {
    request = &RebootInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "RebootInstances")
    return
}

func NewRebootInstancesResponse() (response *RebootInstancesResponse) {
    response = &RebootInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to restart instances.
// 
// * You can only perform this operation on instances whose status is `RUNNING`.
// * If the API is called successfully, the instance status will become `REBOOTING`. After the instance is restarted, its status will become `RUNNING` again.
// * Forced restart is supported. A forced restart is similar to switching off the power of a physical computer and starting it again. It may cause data loss or file system corruption. Be sure to only force start a CVM when it cannot be restarted normally.
// * Batch operations are supported. The maximum number of instances in each request is 100.
func (c *Client) RebootInstances(request *RebootInstancesRequest) (response *RebootInstancesResponse, err error) {
    if request == nil {
        request = NewRebootInstancesRequest()
    }
    response = NewRebootInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewResetInstanceRequest() (request *ResetInstanceRequest) {
    request = &ResetInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "ResetInstance")
    return
}

func NewResetInstanceResponse() (response *ResetInstanceResponse) {
    response = &ResetInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to reinstall the operating system of the specified instance.
// 
// * If you specify an `ImageId`, the specified image is used. Otherwise, the image used by the current instance is used.
// * The system disk will be formatted and reset. Therefore, make sure that no important files are stored on the system disk.
// * If the operating system switches between `Linux` and `Windows`, the system disk `ID` of the instance will change, and the snapshots that are associated with the system disk can no longer be used to roll back and restore data.
// * If no password is specified, you will get a random password via internal message.
// * You can only use this API to switch the operating system between `Linux` and `Windows` for instances whose [system disk type](https://intl.cloud.tencent.com/document/api/213/9452?from_cn_redirect=1#SystemDisk) is `CLOUD_BASIC`, `CLOUD_PREMIUM`, or `CLOUD_SSD`.
// * Currently, this API only supports instances in Mainland China regions.
func (c *Client) ResetInstance(request *ResetInstanceRequest) (response *ResetInstanceResponse, err error) {
    if request == nil {
        request = NewResetInstanceRequest()
    }
    response = NewResetInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewResetInstancesInternetMaxBandwidthRequest() (request *ResetInstancesInternetMaxBandwidthRequest) {
    request = &ResetInstancesInternetMaxBandwidthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "ResetInstancesInternetMaxBandwidth")
    return
}

func NewResetInstancesInternetMaxBandwidthResponse() (response *ResetInstancesInternetMaxBandwidthResponse) {
    response = &ResetInstancesInternetMaxBandwidthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to change the public bandwidth cap of an instance.
// 
// * The allowed bandwidth cap varies for different models. For details, see [Purchasing Network Bandwidth](https://intl.cloud.tencent.com/document/product/213/509?from_cn_redirect=1).
// * For bandwidth billed by the `TRAFFIC_POSTPAID_BY_HOUR` method, changing the bandwidth cap through this API takes effect in real time. Users can increase or reduce bandwidth within applicable limits.
func (c *Client) ResetInstancesInternetMaxBandwidth(request *ResetInstancesInternetMaxBandwidthRequest) (response *ResetInstancesInternetMaxBandwidthResponse, err error) {
    if request == nil {
        request = NewResetInstancesInternetMaxBandwidthRequest()
    }
    response = NewResetInstancesInternetMaxBandwidthResponse()
    err = c.Send(request, response)
    return
}

func NewResetInstancesPasswordRequest() (request *ResetInstancesPasswordRequest) {
    request = &ResetInstancesPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "ResetInstancesPassword")
    return
}

func NewResetInstancesPasswordResponse() (response *ResetInstancesPasswordResponse) {
    response = &ResetInstancesPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to reset the password of the operating system instances to a user-specified password.
// 
// * To modify the password of the administrator account: the name of the administrator account varies with the operating system. In Windows, it is `Administrator`; in Ubuntu, it is `ubuntu`; in Linux, it is `root`.
// * To reset the password of a running instance, you need to set the parameter `ForceStop` to `True` for a forced shutdown. If not, only passwords of stopped instances can be reset.
// * Batch operations are supported. You can reset the passwords of up to 100 instances to the same value once.
// * You can call the [DescribeInstances](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1#.E7.A4.BA.E4.BE.8B3-.E6.9F.A5.E8.AF.A2.E5.AE.9E.E4.BE.8B.E7.9A.84.E6.9C.80.E6.96.B0.E6.93.8D.E4.BD.9C.E6.83.85.E5.86.B5) API and find the result of the operation in the response parameter `LatestOperationState`. If the value is `SUCCESS`, the operation is successful.
func (c *Client) ResetInstancesPassword(request *ResetInstancesPasswordRequest) (response *ResetInstancesPasswordResponse, err error) {
    if request == nil {
        request = NewResetInstancesPasswordRequest()
    }
    response = NewResetInstancesPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewResetInstancesTypeRequest() (request *ResetInstancesTypeRequest) {
    request = &ResetInstancesTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "ResetInstancesType")
    return
}

func NewResetInstancesTypeResponse() (response *ResetInstancesTypeResponse) {
    response = &ResetInstancesTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to change the model of an instance.
// * You can only use this API to change the models of instances whose [system disk type](https://intl.cloud.tencent.com/document/api/213/9452?from_cn_redirect=1#block_device) is `CLOUD_BASIC`, `CLOUD_PREMIUM`, or `CLOUD_SSD`.
// * Currently, you cannot use this API to change the models of [CDH](https://intl.cloud.tencent.com/document/product/416?from_cn_redirect=1) instances.
func (c *Client) ResetInstancesType(request *ResetInstancesTypeRequest) (response *ResetInstancesTypeResponse, err error) {
    if request == nil {
        request = NewResetInstancesTypeRequest()
    }
    response = NewResetInstancesTypeResponse()
    err = c.Send(request, response)
    return
}

func NewResizeInstanceDisksRequest() (request *ResizeInstanceDisksRequest) {
    request = &ResizeInstanceDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "ResizeInstanceDisks")
    return
}

func NewResizeInstanceDisksResponse() (response *ResizeInstanceDisksResponse) {
    response = &ResizeInstanceDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (ResizeInstanceDisks) is used to expand the data disks of an instance.
// 
// * Currently, you can only use the API to expand non-elastic data disks whose [disk type](https://intl.cloud.tencent.com/document/api/213/9452?from_cn_redirect=1#block_device) is `CLOUD_BASIC`, `CLOUD_PREMIUM`, or `CLOUD_SSD`. You can use [`DescribeDisks`](https://intl.cloud.tencent.com/document/api/362/16315?from_cn_redirect=1) to check whether a disk is elastic. If the `Portable` field in the response is `false`, it means that the disk is non-elastic.
// * Currently, this API does not support [CDH](https://intl.cloud.tencent.com/document/product/416?from_cn_redirect=1) instances.
// * Currently, only one data disk can be expanded at a time.
func (c *Client) ResizeInstanceDisks(request *ResizeInstanceDisksRequest) (response *ResizeInstanceDisksResponse, err error) {
    if request == nil {
        request = NewResizeInstanceDisksRequest()
    }
    response = NewResizeInstanceDisksResponse()
    err = c.Send(request, response)
    return
}

func NewRunInstancesRequest() (request *RunInstancesRequest) {
    request = &RunInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "RunInstances")
    return
}

func NewRunInstancesResponse() (response *RunInstancesResponse) {
    response = &RunInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create one or more instances with a specified configuration.
// 
// * After an instance is created successfully, it will start up automatically, and the [instance state](https://intl.cloud.tencent.com/document/api/213/9452?from_cn_redirect=1#instance_state) will become "Running".
// * If you create a pay-as-you-go instance billed on an hourly basis, an amount equivalent to the hourly rate will be frozen before the creation. Make sure your account balance is sufficient before calling this API.
// * The number of instances you can purchase through this API is subject to the [CVM instance purchase limit](https://intl.cloud.tencent.com/document/product/213/2664?from_cn_redirect=1). Both the instances created through this API and the console will be counted toward the quota.
// * This API is an async API. An instance `ID` list will be returned after you successfully make a creation request. However, it does not mean the creation has been completed. The state of the instance will be `Creating` during the creation. You can use [DescribeInstances](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) to query the status of the instance. If the status changes from `Creating` to `Running`, it means that the instance has been created successfully.
func (c *Client) RunInstances(request *RunInstancesRequest) (response *RunInstancesResponse, err error) {
    if request == nil {
        request = NewRunInstancesRequest()
    }
    response = NewRunInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewStartInstancesRequest() (request *StartInstancesRequest) {
    request = &StartInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "StartInstances")
    return
}

func NewStartInstancesResponse() (response *StartInstancesResponse) {
    response = &StartInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to start instances.
// 
// * You can only perform this operation on instances whose status is `STOPPED`.
// * The instance status will become `STARTING` when the API is called successfully and `RUNNING` when the instance is successfully started.
// * Batch operations are supported. The maximum number of instances in each request is 100.
func (c *Client) StartInstances(request *StartInstancesRequest) (response *StartInstancesResponse, err error) {
    if request == nil {
        request = NewStartInstancesRequest()
    }
    response = NewStartInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewStopInstancesRequest() (request *StopInstancesRequest) {
    request = &StopInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "StopInstances")
    return
}

func NewStopInstancesResponse() (response *StopInstancesResponse) {
    response = &StopInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to shut down instances.
// 
// * You can only perform this operation on instances whose status is `RUNNING`.
// * The instance status will become `STOPPING` when the API is called successfully and `STOPPED` when the instance is successfully shut down.
// * Forced shutdown is supported. A forced shutdown is similar to switching off the power of a physical computer. It may cause data loss or file system corruption. Be sure to only force shut down a CVM when it cannot be sht down normally.
// * Batch operations are supported. The maximum number of instances in each request is 100.
func (c *Client) StopInstances(request *StopInstancesRequest) (response *StopInstancesResponse, err error) {
    if request == nil {
        request = NewStopInstancesRequest()
    }
    response = NewStopInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewSyncImagesRequest() (request *SyncImagesRequest) {
    request = &SyncImagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "SyncImages")
    return
}

func NewSyncImagesResponse() (response *SyncImagesResponse) {
    response = &SyncImagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to sync a custom image to other regions.
// 
// * Each API call syncs a single image.
// * This API supports syncing an image to multiple regions.
// * Each account can have up to 10 custom images in each region. 
func (c *Client) SyncImages(request *SyncImagesRequest) (response *SyncImagesResponse, err error) {
    if request == nil {
        request = NewSyncImagesRequest()
    }
    response = NewSyncImagesResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateInstancesRequest() (request *TerminateInstancesRequest) {
    request = &TerminateInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "TerminateInstances")
    return
}

func NewTerminateInstancesResponse() (response *TerminateInstancesResponse) {
    response = &TerminateInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to return instances.
// 
// * Use this API to return instances that are no longer required.
// * Pay-as-you-go instances can be returned directly through this API.
// * When this API is called for the first time, the instance will be moved to the recycle bin. When this API is called for the second time, the instance will be terminated and cannot be recovered.
// * Batch operations are supported. The allowed maximum number of instances in each request is 100.
func (c *Client) TerminateInstances(request *TerminateInstancesRequest) (response *TerminateInstancesResponse, err error) {
    if request == nil {
        request = NewTerminateInstancesRequest()
    }
    response = NewTerminateInstancesResponse()
    err = c.Send(request, response)
    return
}
