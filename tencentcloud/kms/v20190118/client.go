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

package v20190118

import (
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2019-01-18"

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


func NewAsymmetricRsaDecryptRequest() (request *AsymmetricRsaDecryptRequest) {
    request = &AsymmetricRsaDecryptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "AsymmetricRsaDecrypt")
    return
}

func NewAsymmetricRsaDecryptResponse() (response *AsymmetricRsaDecryptResponse) {
    response = &AsymmetricRsaDecryptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to decrypt data with the specified private key that is encrypted with RSA asymmetric cryptographic algorithm. The ciphertext must be encrypted with the corresponding public key. The asymmetric key must be in `Enabled` state for decryption.
func (c *Client) AsymmetricRsaDecrypt(request *AsymmetricRsaDecryptRequest) (response *AsymmetricRsaDecryptResponse, err error) {
    if request == nil {
        request = NewAsymmetricRsaDecryptRequest()
    }
    response = NewAsymmetricRsaDecryptResponse()
    err = c.Send(request, response)
    return
}

func NewAsymmetricSm2DecryptRequest() (request *AsymmetricSm2DecryptRequest) {
    request = &AsymmetricSm2DecryptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "AsymmetricSm2Decrypt")
    return
}

func NewAsymmetricSm2DecryptResponse() (response *AsymmetricSm2DecryptResponse) {
    response = &AsymmetricSm2DecryptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to decrypt data with the specified private key that is encrypted with SM2 asymmetric cryptographic algorithm. The ciphertext must be encrypted with the corresponding public key. The asymmetric key must be in `Enabled` state for decryption. The length of the ciphertext passed in cannot exceed 256 bytes.
func (c *Client) AsymmetricSm2Decrypt(request *AsymmetricSm2DecryptRequest) (response *AsymmetricSm2DecryptResponse, err error) {
    if request == nil {
        request = NewAsymmetricSm2DecryptRequest()
    }
    response = NewAsymmetricSm2DecryptResponse()
    err = c.Send(request, response)
    return
}

func NewCancelKeyDeletionRequest() (request *CancelKeyDeletionRequest) {
    request = &CancelKeyDeletionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "CancelKeyDeletion")
    return
}

func NewCancelKeyDeletionResponse() (response *CancelKeyDeletionResponse) {
    response = &CancelKeyDeletionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Cancel the scheduled deletion of CMK
func (c *Client) CancelKeyDeletion(request *CancelKeyDeletionRequest) (response *CancelKeyDeletionResponse, err error) {
    if request == nil {
        request = NewCancelKeyDeletionRequest()
    }
    response = NewCancelKeyDeletionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateKeyRequest() (request *CreateKeyRequest) {
    request = &CreateKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "CreateKey")
    return
}

func NewCreateKeyResponse() (response *CreateKeyResponse) {
    response = &CreateKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Create a master key CMK (Custom Master Key) for user management data keys
func (c *Client) CreateKey(request *CreateKeyRequest) (response *CreateKeyResponse, err error) {
    if request == nil {
        request = NewCreateKeyRequest()
    }
    response = NewCreateKeyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWhiteBoxKeyRequest() (request *CreateWhiteBoxKeyRequest) {
    request = &CreateWhiteBoxKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "CreateWhiteBoxKey")
    return
}

func NewCreateWhiteBoxKeyResponse() (response *CreateWhiteBoxKeyResponse) {
    response = &CreateWhiteBoxKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create a white-box key. Up to 50 ones can be created.
func (c *Client) CreateWhiteBoxKey(request *CreateWhiteBoxKeyRequest) (response *CreateWhiteBoxKeyResponse, err error) {
    if request == nil {
        request = NewCreateWhiteBoxKeyRequest()
    }
    response = NewCreateWhiteBoxKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDecryptRequest() (request *DecryptRequest) {
    request = &DecryptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "Decrypt")
    return
}

func NewDecryptResponse() (response *DecryptResponse) {
    response = &DecryptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to decrypt the ciphertext and obtain the plaintext data.
func (c *Client) Decrypt(request *DecryptRequest) (response *DecryptResponse, err error) {
    if request == nil {
        request = NewDecryptRequest()
    }
    response = NewDecryptResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteImportedKeyMaterialRequest() (request *DeleteImportedKeyMaterialRequest) {
    request = &DeleteImportedKeyMaterialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "DeleteImportedKeyMaterial")
    return
}

func NewDeleteImportedKeyMaterialResponse() (response *DeleteImportedKeyMaterialResponse) {
    response = &DeleteImportedKeyMaterialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete the imported key material. It is only valid for EXTERNAL CMKs. Specifically, it puts a CMK into `PendingImport` status instead of deleting the CMK, so that the CMK can be used again after key material is reimported. To delete the CMK completely, please call the `ScheduleKeyDeletion` API.
func (c *Client) DeleteImportedKeyMaterial(request *DeleteImportedKeyMaterialRequest) (response *DeleteImportedKeyMaterialResponse, err error) {
    if request == nil {
        request = NewDeleteImportedKeyMaterialRequest()
    }
    response = NewDeleteImportedKeyMaterialResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWhiteBoxKeyRequest() (request *DeleteWhiteBoxKeyRequest) {
    request = &DeleteWhiteBoxKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "DeleteWhiteBoxKey")
    return
}

func NewDeleteWhiteBoxKeyResponse() (response *DeleteWhiteBoxKeyResponse) {
    response = &DeleteWhiteBoxKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete a white-box key. Note: only disabled white-box keys can be deleted.
func (c *Client) DeleteWhiteBoxKey(request *DeleteWhiteBoxKeyRequest) (response *DeleteWhiteBoxKeyResponse, err error) {
    if request == nil {
        request = NewDeleteWhiteBoxKeyRequest()
    }
    response = NewDeleteWhiteBoxKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKeyRequest() (request *DescribeKeyRequest) {
    request = &DescribeKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "DescribeKey")
    return
}

func NewDescribeKeyResponse() (response *DescribeKeyResponse) {
    response = &DescribeKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the attribute details of the CMK with a specified `KeyId`.
func (c *Client) DescribeKey(request *DescribeKeyRequest) (response *DescribeKeyResponse, err error) {
    if request == nil {
        request = NewDescribeKeyRequest()
    }
    response = NewDescribeKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKeysRequest() (request *DescribeKeysRequest) {
    request = &DescribeKeysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "DescribeKeys")
    return
}

func NewDescribeKeysResponse() (response *DescribeKeysResponse) {
    response = &DescribeKeysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the attribute information of CMKs in batches.
func (c *Client) DescribeKeys(request *DescribeKeysRequest) (response *DescribeKeysResponse, err error) {
    if request == nil {
        request = NewDescribeKeysRequest()
    }
    response = NewDescribeKeysResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWhiteBoxDecryptKeyRequest() (request *DescribeWhiteBoxDecryptKeyRequest) {
    request = &DescribeWhiteBoxDecryptKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "DescribeWhiteBoxDecryptKey")
    return
}

func NewDescribeWhiteBoxDecryptKeyResponse() (response *DescribeWhiteBoxDecryptKeyResponse) {
    response = &DescribeWhiteBoxDecryptKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get a white-box decryption key.
func (c *Client) DescribeWhiteBoxDecryptKey(request *DescribeWhiteBoxDecryptKeyRequest) (response *DescribeWhiteBoxDecryptKeyResponse, err error) {
    if request == nil {
        request = NewDescribeWhiteBoxDecryptKeyRequest()
    }
    response = NewDescribeWhiteBoxDecryptKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWhiteBoxDeviceFingerprintsRequest() (request *DescribeWhiteBoxDeviceFingerprintsRequest) {
    request = &DescribeWhiteBoxDeviceFingerprintsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "DescribeWhiteBoxDeviceFingerprints")
    return
}

func NewDescribeWhiteBoxDeviceFingerprintsResponse() (response *DescribeWhiteBoxDeviceFingerprintsResponse) {
    response = &DescribeWhiteBoxDeviceFingerprintsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the device fingerprint list of a specified key.
func (c *Client) DescribeWhiteBoxDeviceFingerprints(request *DescribeWhiteBoxDeviceFingerprintsRequest) (response *DescribeWhiteBoxDeviceFingerprintsResponse, err error) {
    if request == nil {
        request = NewDescribeWhiteBoxDeviceFingerprintsRequest()
    }
    response = NewDescribeWhiteBoxDeviceFingerprintsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWhiteBoxKeyRequest() (request *DescribeWhiteBoxKeyRequest) {
    request = &DescribeWhiteBoxKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "DescribeWhiteBoxKey")
    return
}

func NewDescribeWhiteBoxKeyResponse() (response *DescribeWhiteBoxKeyResponse) {
    response = &DescribeWhiteBoxKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to display white-box key information.
func (c *Client) DescribeWhiteBoxKey(request *DescribeWhiteBoxKeyRequest) (response *DescribeWhiteBoxKeyResponse, err error) {
    if request == nil {
        request = NewDescribeWhiteBoxKeyRequest()
    }
    response = NewDescribeWhiteBoxKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWhiteBoxKeyDetailsRequest() (request *DescribeWhiteBoxKeyDetailsRequest) {
    request = &DescribeWhiteBoxKeyDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "DescribeWhiteBoxKeyDetails")
    return
}

func NewDescribeWhiteBoxKeyDetailsResponse() (response *DescribeWhiteBoxKeyDetailsResponse) {
    response = &DescribeWhiteBoxKeyDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the white-box key list.
func (c *Client) DescribeWhiteBoxKeyDetails(request *DescribeWhiteBoxKeyDetailsRequest) (response *DescribeWhiteBoxKeyDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeWhiteBoxKeyDetailsRequest()
    }
    response = NewDescribeWhiteBoxKeyDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWhiteBoxServiceStatusRequest() (request *DescribeWhiteBoxServiceStatusRequest) {
    request = &DescribeWhiteBoxServiceStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "DescribeWhiteBoxServiceStatus")
    return
}

func NewDescribeWhiteBoxServiceStatusResponse() (response *DescribeWhiteBoxServiceStatusResponse) {
    response = &DescribeWhiteBoxServiceStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the white-box key service status.
func (c *Client) DescribeWhiteBoxServiceStatus(request *DescribeWhiteBoxServiceStatusRequest) (response *DescribeWhiteBoxServiceStatusResponse, err error) {
    if request == nil {
        request = NewDescribeWhiteBoxServiceStatusRequest()
    }
    response = NewDescribeWhiteBoxServiceStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDisableKeyRequest() (request *DisableKeyRequest) {
    request = &DisableKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "DisableKey")
    return
}

func NewDisableKeyResponse() (response *DisableKeyResponse) {
    response = &DisableKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to disable a master key. The disabled key cannot be used for encryption and decryption operations.
func (c *Client) DisableKey(request *DisableKeyRequest) (response *DisableKeyResponse, err error) {
    if request == nil {
        request = NewDisableKeyRequest()
    }
    response = NewDisableKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDisableKeyRotationRequest() (request *DisableKeyRotationRequest) {
    request = &DisableKeyRotationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "DisableKeyRotation")
    return
}

func NewDisableKeyRotationResponse() (response *DisableKeyRotationResponse) {
    response = &DisableKeyRotationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Disabled key rotation for the specified CMK.
func (c *Client) DisableKeyRotation(request *DisableKeyRotationRequest) (response *DisableKeyRotationResponse, err error) {
    if request == nil {
        request = NewDisableKeyRotationRequest()
    }
    response = NewDisableKeyRotationResponse()
    err = c.Send(request, response)
    return
}

func NewDisableKeysRequest() (request *DisableKeysRequest) {
    request = &DisableKeysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "DisableKeys")
    return
}

func NewDisableKeysResponse() (response *DisableKeysResponse) {
    response = &DisableKeysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to batch prohibit the use of CMK.
func (c *Client) DisableKeys(request *DisableKeysRequest) (response *DisableKeysResponse, err error) {
    if request == nil {
        request = NewDisableKeysRequest()
    }
    response = NewDisableKeysResponse()
    err = c.Send(request, response)
    return
}

func NewDisableWhiteBoxKeyRequest() (request *DisableWhiteBoxKeyRequest) {
    request = &DisableWhiteBoxKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "DisableWhiteBoxKey")
    return
}

func NewDisableWhiteBoxKeyResponse() (response *DisableWhiteBoxKeyResponse) {
    response = &DisableWhiteBoxKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to disable a white-box key.
func (c *Client) DisableWhiteBoxKey(request *DisableWhiteBoxKeyRequest) (response *DisableWhiteBoxKeyResponse, err error) {
    if request == nil {
        request = NewDisableWhiteBoxKeyRequest()
    }
    response = NewDisableWhiteBoxKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDisableWhiteBoxKeysRequest() (request *DisableWhiteBoxKeysRequest) {
    request = &DisableWhiteBoxKeysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "DisableWhiteBoxKeys")
    return
}

func NewDisableWhiteBoxKeysResponse() (response *DisableWhiteBoxKeysResponse) {
    response = &DisableWhiteBoxKeysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to disable white-box keys in batches.
func (c *Client) DisableWhiteBoxKeys(request *DisableWhiteBoxKeysRequest) (response *DisableWhiteBoxKeysResponse, err error) {
    if request == nil {
        request = NewDisableWhiteBoxKeysRequest()
    }
    response = NewDisableWhiteBoxKeysResponse()
    err = c.Send(request, response)
    return
}

func NewEnableKeyRequest() (request *EnableKeyRequest) {
    request = &EnableKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "EnableKey")
    return
}

func NewEnableKeyResponse() (response *EnableKeyResponse) {
    response = &EnableKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Enable a specified CMK.
func (c *Client) EnableKey(request *EnableKeyRequest) (response *EnableKeyResponse, err error) {
    if request == nil {
        request = NewEnableKeyRequest()
    }
    response = NewEnableKeyResponse()
    err = c.Send(request, response)
    return
}

func NewEnableKeyRotationRequest() (request *EnableKeyRotationRequest) {
    request = &EnableKeyRotationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "EnableKeyRotation")
    return
}

func NewEnableKeyRotationResponse() (response *EnableKeyRotationResponse) {
    response = &EnableKeyRotationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Turn on the key rotation function for the specified CMK.
func (c *Client) EnableKeyRotation(request *EnableKeyRotationRequest) (response *EnableKeyRotationResponse, err error) {
    if request == nil {
        request = NewEnableKeyRotationRequest()
    }
    response = NewEnableKeyRotationResponse()
    err = c.Send(request, response)
    return
}

func NewEnableKeysRequest() (request *EnableKeysRequest) {
    request = &EnableKeysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "EnableKeys")
    return
}

func NewEnableKeysResponse() (response *EnableKeysResponse) {
    response = &EnableKeysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to enable CMK in batches.
func (c *Client) EnableKeys(request *EnableKeysRequest) (response *EnableKeysResponse, err error) {
    if request == nil {
        request = NewEnableKeysRequest()
    }
    response = NewEnableKeysResponse()
    err = c.Send(request, response)
    return
}

func NewEnableWhiteBoxKeyRequest() (request *EnableWhiteBoxKeyRequest) {
    request = &EnableWhiteBoxKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "EnableWhiteBoxKey")
    return
}

func NewEnableWhiteBoxKeyResponse() (response *EnableWhiteBoxKeyResponse) {
    response = &EnableWhiteBoxKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to enable a white-box key.
func (c *Client) EnableWhiteBoxKey(request *EnableWhiteBoxKeyRequest) (response *EnableWhiteBoxKeyResponse, err error) {
    if request == nil {
        request = NewEnableWhiteBoxKeyRequest()
    }
    response = NewEnableWhiteBoxKeyResponse()
    err = c.Send(request, response)
    return
}

func NewEnableWhiteBoxKeysRequest() (request *EnableWhiteBoxKeysRequest) {
    request = &EnableWhiteBoxKeysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "EnableWhiteBoxKeys")
    return
}

func NewEnableWhiteBoxKeysResponse() (response *EnableWhiteBoxKeysResponse) {
    response = &EnableWhiteBoxKeysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to enable white-box keys in batches.
func (c *Client) EnableWhiteBoxKeys(request *EnableWhiteBoxKeysRequest) (response *EnableWhiteBoxKeysResponse, err error) {
    if request == nil {
        request = NewEnableWhiteBoxKeysRequest()
    }
    response = NewEnableWhiteBoxKeysResponse()
    err = c.Send(request, response)
    return
}

func NewEncryptRequest() (request *EncryptRequest) {
    request = &EncryptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "Encrypt")
    return
}

func NewEncryptResponse() (response *EncryptResponse) {
    response = &EncryptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to encrypt any data up to 4KB. It can be used to encrypt database passwords, RSA Key, or other small sensitive information. For application data encryption, use the DataKey generated by GenerateDataKey to perform local data encryption and decryption operations
func (c *Client) Encrypt(request *EncryptRequest) (response *EncryptResponse, err error) {
    if request == nil {
        request = NewEncryptRequest()
    }
    response = NewEncryptResponse()
    err = c.Send(request, response)
    return
}

func NewEncryptByWhiteBoxRequest() (request *EncryptByWhiteBoxRequest) {
    request = &EncryptByWhiteBoxRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "EncryptByWhiteBox")
    return
}

func NewEncryptByWhiteBoxResponse() (response *EncryptByWhiteBoxResponse) {
    response = &EncryptByWhiteBoxResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to encrypt data with a white-box key.
func (c *Client) EncryptByWhiteBox(request *EncryptByWhiteBoxRequest) (response *EncryptByWhiteBoxResponse, err error) {
    if request == nil {
        request = NewEncryptByWhiteBoxRequest()
    }
    response = NewEncryptByWhiteBoxResponse()
    err = c.Send(request, response)
    return
}

func NewGenerateDataKeyRequest() (request *GenerateDataKeyRequest) {
    request = &GenerateDataKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "GenerateDataKey")
    return
}

func NewGenerateDataKeyResponse() (response *GenerateDataKeyResponse) {
    response = &GenerateDataKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API generates a data key, which you can use to encrypt local data.
func (c *Client) GenerateDataKey(request *GenerateDataKeyRequest) (response *GenerateDataKeyResponse, err error) {
    if request == nil {
        request = NewGenerateDataKeyRequest()
    }
    response = NewGenerateDataKeyResponse()
    err = c.Send(request, response)
    return
}

func NewGenerateRandomRequest() (request *GenerateRandomRequest) {
    request = &GenerateRandomRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "GenerateRandom")
    return
}

func NewGenerateRandomResponse() (response *GenerateRandomResponse) {
    response = &GenerateRandomResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to generate a random number.
func (c *Client) GenerateRandom(request *GenerateRandomRequest) (response *GenerateRandomResponse, err error) {
    if request == nil {
        request = NewGenerateRandomRequest()
    }
    response = NewGenerateRandomResponse()
    err = c.Send(request, response)
    return
}

func NewGetKeyRotationStatusRequest() (request *GetKeyRotationStatusRequest) {
    request = &GetKeyRotationStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "GetKeyRotationStatus")
    return
}

func NewGetKeyRotationStatusResponse() (response *GetKeyRotationStatusResponse) {
    response = &GetKeyRotationStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Query whether the specified CMK has the key rotation function.
func (c *Client) GetKeyRotationStatus(request *GetKeyRotationStatusRequest) (response *GetKeyRotationStatusResponse, err error) {
    if request == nil {
        request = NewGetKeyRotationStatusRequest()
    }
    response = NewGetKeyRotationStatusResponse()
    err = c.Send(request, response)
    return
}

func NewGetParametersForImportRequest() (request *GetParametersForImportRequest) {
    request = &GetParametersForImportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "GetParametersForImport")
    return
}

func NewGetParametersForImportResponse() (response *GetParametersForImportResponse) {
    response = &GetParametersForImportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to obtain the parameters of the material to be imported into a CMK. The returned `Token` is used as one of the parameters to execute `ImportKeyMaterial`, and the returned `PublicKey` is used to encrypt the key material. The `Token` and `PublicKey` are valid for 24 hours. If they are expired, you will need to call the API again to get a new `Token` and `PublicKey`.
func (c *Client) GetParametersForImport(request *GetParametersForImportRequest) (response *GetParametersForImportResponse, err error) {
    if request == nil {
        request = NewGetParametersForImportRequest()
    }
    response = NewGetParametersForImportResponse()
    err = c.Send(request, response)
    return
}

func NewGetPublicKeyRequest() (request *GetPublicKeyRequest) {
    request = &GetPublicKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "GetPublicKey")
    return
}

func NewGetPublicKeyResponse() (response *GetPublicKeyResponse) {
    response = &GetPublicKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the information of the public key that is encrypted with the asymmetric cryptographic algorithm and of which the `KeyUsage` is `ASYMMETRIC_DECRYPT_RSA_2048` or `ASYMMETRIC_DECRYPT_SM2`. This public key can be used to encrypt data locally, and the data encrypted with it can only be decrypted with the corresponding private key through KMS. The public key can only be obtained from the asymmetric key in `Enabled` state.
func (c *Client) GetPublicKey(request *GetPublicKeyRequest) (response *GetPublicKeyResponse, err error) {
    if request == nil {
        request = NewGetPublicKeyRequest()
    }
    response = NewGetPublicKeyResponse()
    err = c.Send(request, response)
    return
}

func NewGetServiceStatusRequest() (request *GetServiceStatusRequest) {
    request = &GetServiceStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "GetServiceStatus")
    return
}

func NewGetServiceStatusResponse() (response *GetServiceStatusResponse) {
    response = &GetServiceStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Used to query whether the user has activated the KMS service.
func (c *Client) GetServiceStatus(request *GetServiceStatusRequest) (response *GetServiceStatusResponse, err error) {
    if request == nil {
        request = NewGetServiceStatusRequest()
    }
    response = NewGetServiceStatusResponse()
    err = c.Send(request, response)
    return
}

func NewImportKeyMaterialRequest() (request *ImportKeyMaterialRequest) {
    request = &ImportKeyMaterialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "ImportKeyMaterial")
    return
}

func NewImportKeyMaterialResponse() (response *ImportKeyMaterialResponse) {
    response = &ImportKeyMaterialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to import key material into an EXTERNAL CMK. The key obtained through the `GetParametersForImport` API is used to encrypt the key material. You can only reimport the same key material into the specified CMK and set a new expiration time. After the CMK key material is imported, it cannot be replaced. After the key material is expired or deleted, the CMK will remain unavailable until the same key material is reimported. CMKs are independent, which means that the same key material can be imported into different CMKs, but data encrypted by one CMK cannot be decrypted by another one.
// Key material can only be imported into CMKs in `Enabled` and `PendingImport` status.
func (c *Client) ImportKeyMaterial(request *ImportKeyMaterialRequest) (response *ImportKeyMaterialResponse, err error) {
    if request == nil {
        request = NewImportKeyMaterialRequest()
    }
    response = NewImportKeyMaterialResponse()
    err = c.Send(request, response)
    return
}

func NewListAlgorithmsRequest() (request *ListAlgorithmsRequest) {
    request = &ListAlgorithmsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "ListAlgorithms")
    return
}

func NewListAlgorithmsResponse() (response *ListAlgorithmsResponse) {
    response = &ListAlgorithmsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to list the encryption methods supported in the current region.
func (c *Client) ListAlgorithms(request *ListAlgorithmsRequest) (response *ListAlgorithmsResponse, err error) {
    if request == nil {
        request = NewListAlgorithmsRequest()
    }
    response = NewListAlgorithmsResponse()
    err = c.Send(request, response)
    return
}

func NewListKeyDetailRequest() (request *ListKeyDetailRequest) {
    request = &ListKeyDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "ListKeyDetail")
    return
}

func NewListKeyDetailResponse() (response *ListKeyDetailResponse) {
    response = &ListKeyDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Get the master key list details according to the specified Offset and Limit.
func (c *Client) ListKeyDetail(request *ListKeyDetailRequest) (response *ListKeyDetailResponse, err error) {
    if request == nil {
        request = NewListKeyDetailRequest()
    }
    response = NewListKeyDetailResponse()
    err = c.Send(request, response)
    return
}

func NewListKeysRequest() (request *ListKeysRequest) {
    request = &ListKeysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "ListKeys")
    return
}

func NewListKeysResponse() (response *ListKeysResponse) {
    response = &ListKeysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to list the KeyIds of CMKs in `Enabled`, `Disabled`, and `PendingImport` status under the account.
func (c *Client) ListKeys(request *ListKeysRequest) (response *ListKeysResponse, err error) {
    if request == nil {
        request = NewListKeysRequest()
    }
    response = NewListKeysResponse()
    err = c.Send(request, response)
    return
}

func NewOverwriteWhiteBoxDeviceFingerprintsRequest() (request *OverwriteWhiteBoxDeviceFingerprintsRequest) {
    request = &OverwriteWhiteBoxDeviceFingerprintsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "OverwriteWhiteBoxDeviceFingerprints")
    return
}

func NewOverwriteWhiteBoxDeviceFingerprintsResponse() (response *OverwriteWhiteBoxDeviceFingerprintsResponse) {
    response = &OverwriteWhiteBoxDeviceFingerprintsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to overwrite the device fingerprint information of a specified key.
func (c *Client) OverwriteWhiteBoxDeviceFingerprints(request *OverwriteWhiteBoxDeviceFingerprintsRequest) (response *OverwriteWhiteBoxDeviceFingerprintsResponse, err error) {
    if request == nil {
        request = NewOverwriteWhiteBoxDeviceFingerprintsRequest()
    }
    response = NewOverwriteWhiteBoxDeviceFingerprintsResponse()
    err = c.Send(request, response)
    return
}

func NewReEncryptRequest() (request *ReEncryptRequest) {
    request = &ReEncryptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "ReEncrypt")
    return
}

func NewReEncryptResponse() (response *ReEncryptResponse) {
    response = &ReEncryptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Re-encrypt the ciphertext using the specified CMK.
func (c *Client) ReEncrypt(request *ReEncryptRequest) (response *ReEncryptResponse, err error) {
    if request == nil {
        request = NewReEncryptRequest()
    }
    response = NewReEncryptResponse()
    err = c.Send(request, response)
    return
}

func NewScheduleKeyDeletionRequest() (request *ScheduleKeyDeletionRequest) {
    request = &ScheduleKeyDeletionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "ScheduleKeyDeletion")
    return
}

func NewScheduleKeyDeletionResponse() (response *ScheduleKeyDeletionResponse) {
    response = &ScheduleKeyDeletionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CMK planned deletion API, used to specify the time of CMK deletion, the optional time interval is [7,30] days
func (c *Client) ScheduleKeyDeletion(request *ScheduleKeyDeletionRequest) (response *ScheduleKeyDeletionResponse, err error) {
    if request == nil {
        request = NewScheduleKeyDeletionRequest()
    }
    response = NewScheduleKeyDeletionResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAliasRequest() (request *UpdateAliasRequest) {
    request = &UpdateAliasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "UpdateAlias")
    return
}

func NewUpdateAliasResponse() (response *UpdateAliasResponse) {
    response = &UpdateAliasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify the alias of a CMK. CMKs in `PendingDelete` status cannot be modified.
func (c *Client) UpdateAlias(request *UpdateAliasRequest) (response *UpdateAliasResponse, err error) {
    if request == nil {
        request = NewUpdateAliasRequest()
    }
    response = NewUpdateAliasResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateKeyDescriptionRequest() (request *UpdateKeyDescriptionRequest) {
    request = &UpdateKeyDescriptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "UpdateKeyDescription")
    return
}

func NewUpdateKeyDescriptionResponse() (response *UpdateKeyDescriptionResponse) {
    response = &UpdateKeyDescriptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify the description of the specified CMK. CMKs in `PendingDelete` status cannot be modified.
func (c *Client) UpdateKeyDescription(request *UpdateKeyDescriptionRequest) (response *UpdateKeyDescriptionResponse, err error) {
    if request == nil {
        request = NewUpdateKeyDescriptionRequest()
    }
    response = NewUpdateKeyDescriptionResponse()
    err = c.Send(request, response)
    return
}
