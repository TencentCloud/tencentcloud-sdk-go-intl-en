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
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AlgorithmInfo struct {

	// Algorithm identification
	KeyUsage *string `json:"KeyUsage,omitempty" name:"KeyUsage"`

	// The name of the algorithm
	Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`
}

type AsymmetricRsaDecryptRequest struct {
	*tchttp.BaseRequest

	// CMK unique identifier
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Encrypted ciphertext using PublicKey, Base64 encoded
	Ciphertext *string `json:"Ciphertext,omitempty" name:"Ciphertext"`

	// Corresponding algorithm when using public key encryption: currently supports RSAES_PKCS1_V1_5, RSAES_OAEP_SHA_1, RSAES_OAEP_SHA_256
	Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`
}

func (r *AsymmetricRsaDecryptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AsymmetricRsaDecryptRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AsymmetricRsaDecryptResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// CMK unique identifier
		KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

		// Decrypted plaintext, base64 encoded
		Plaintext *string `json:"Plaintext,omitempty" name:"Plaintext"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AsymmetricRsaDecryptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AsymmetricRsaDecryptResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AsymmetricSm2DecryptRequest struct {
	*tchttp.BaseRequest

	// CMK unique identifier
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Encrypted ciphertext using PublicKey, Base64 encoded. The cipher text cannot exceed 256 bytes.
	Ciphertext *string `json:"Ciphertext,omitempty" name:"Ciphertext"`
}

func (r *AsymmetricSm2DecryptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AsymmetricSm2DecryptRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AsymmetricSm2DecryptResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// CMK unique identifier
		KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

		// Decrypted plaintext, base64 encoded
		Plaintext *string `json:"Plaintext,omitempty" name:"Plaintext"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AsymmetricSm2DecryptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AsymmetricSm2DecryptResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CancelKeyDeletionRequest struct {
	*tchttp.BaseRequest
}

func (r *CancelKeyDeletionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CancelKeyDeletionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CancelKeyDeletionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelKeyDeletionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CancelKeyDeletionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateKeyRequest struct {
	*tchttp.BaseRequest

	// As an alias that is easier to identify and easier to understand, the key must not be empty, a combination of 1-60 alphanumeric characters-_, and the first character must be a letter or number. The kms- prefix is used for cloud product usage. Alias is not repeatable.
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 
	Description *string `json:"Description,omitempty" name:"Description"`

	// Purpose of the specified key, Default "ENCRYPT_DECRYPT" Indicates the creation of a symmetric encryption and decryption key, other supported uses "ASYMMETRIC_DECRYPT_RSA_2048" means creating an RSA2048 asymmetric key for encryption and decryption,"ASYMMETRIC_DECRYPT_SM2" means creating a SM2 asymmetric key for encryption and decryption
	KeyUsage *string `json:"KeyUsage,omitempty" name:"KeyUsage"`

	// Specifies the key type. Default value: 1. Valid value: 1 - default type, indicating that the CMK is created by KMS; 2 - EXTERNAL type, indicating that you need to import key material. For more information, please see the `GetParametersForImport` and `ImportKeyMaterial` API documents.
	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *CreateKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DecryptRequest struct {
	*tchttp.BaseRequest

	// The ciphertext data to be decrypted.
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" name:"CiphertextBlob"`

	// 
	EncryptionContext *string `json:"EncryptionContext,omitempty" name:"EncryptionContext"`
}

func (r *DecryptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DecryptRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DecryptResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DecryptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DecryptResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteImportedKeyMaterialRequest struct {
	*tchttp.BaseRequest

	// Specifies the EXTERNAL CMK for which to delete the key material.
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *DeleteImportedKeyMaterialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteImportedKeyMaterialRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteImportedKeyMaterialResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteImportedKeyMaterialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteImportedKeyMaterialResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeKeyRequest struct {
	*tchttp.BaseRequest

	// CMK globally unique identifier
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *DescribeKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Key attribute information 
	//  Note: This field may return null, indicating that a valid value cannot be taken.
		KeyMetadata *KeyMetadata `json:"KeyMetadata,omitempty" name:"KeyMetadata"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeKeysRequest struct {
	*tchttp.BaseRequest

	// Query the CMK ID list. Batch query supports up to 100 KeyIds at a time.
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds" list`
}

func (r *DescribeKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeKeysRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeKeysResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of attribute information returned 
	//  Note: This field may return null, indicating that a valid value could not be taken.
		KeyMetadatas []*KeyMetadata `json:"KeyMetadatas,omitempty" name:"KeyMetadatas" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeKeysResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableKeyRequest struct {
	*tchttp.BaseRequest
}

func (r *DisableKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableKeyRotationRequest struct {
	*tchttp.BaseRequest
}

func (r *DisableKeyRotationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableKeyRotationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableKeyRotationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableKeyRotationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableKeyRotationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableKeysRequest struct {
	*tchttp.BaseRequest
}

func (r *DisableKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableKeysRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableKeysResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableKeysResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableKeyRequest struct {
	*tchttp.BaseRequest
}

func (r *EnableKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableKeyRotationRequest struct {
	*tchttp.BaseRequest
}

func (r *EnableKeyRotationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableKeyRotationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableKeyRotationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableKeyRotationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableKeyRotationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableKeysRequest struct {
	*tchttp.BaseRequest
}

func (r *EnableKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableKeysRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableKeysResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableKeysResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EncryptRequest struct {
	*tchttp.BaseRequest
}

func (r *EncryptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EncryptRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EncryptResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Base64-encoded encrypted ciphertext
		CiphertextBlob *string `json:"CiphertextBlob,omitempty" name:"CiphertextBlob"`

		// 
		KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EncryptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EncryptResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GenerateDataKeyRequest struct {
	*tchttp.BaseRequest

	// 
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Specifies the encryption algorithm and size of the `DataKey`. Valid values: AES_128, AES_256. Either `KeySpec` or `NumberOfBytes` must be specified.
	KeySpec *string `json:"KeySpec,omitempty" name:"KeySpec"`

	// Length of the `DataKey`. If both `NumberOfBytes` and `KeySpec` are specified, `NumberOfBytes` will prevail. Minimum value: 1; maximum value: 1024. Either `KeySpec` or `NumberOfBytes` must be specified.
	NumberOfBytes *uint64 `json:"NumberOfBytes,omitempty" name:"NumberOfBytes"`

	// 
	EncryptionContext *string `json:"EncryptionContext,omitempty" name:"EncryptionContext"`
}

func (r *GenerateDataKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GenerateDataKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GenerateDataKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 
		KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

		// 
		Plaintext *string `json:"Plaintext,omitempty" name:"Plaintext"`

		// Base64-encoded ciphertext that is encrypted by `DataKey`. You should keep the ciphertext private.
		CiphertextBlob *string `json:"CiphertextBlob,omitempty" name:"CiphertextBlob"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GenerateDataKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GenerateDataKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GenerateRandomRequest struct {
	*tchttp.BaseRequest

	// Length of the random number. Minimum value: 1. Maximum value: 1024
	NumberOfBytes *uint64 `json:"NumberOfBytes,omitempty" name:"NumberOfBytes"`
}

func (r *GenerateRandomRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GenerateRandomRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GenerateRandomResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Base64-encoded plaintext of the randomly generated number. You need to Base64-decode it to get the plaintext.
		Plaintext *string `json:"Plaintext,omitempty" name:"Plaintext"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GenerateRandomResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GenerateRandomResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetKeyRotationStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *GetKeyRotationStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetKeyRotationStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetKeyRotationStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetKeyRotationStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetKeyRotationStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetParametersForImportRequest struct {
	*tchttp.BaseRequest

	// Unique ID of a CMK. The CMK for which to get the key parameters must be of the `EXTERNAL` type, i.e., Type = 2 when the CMK is created by the `CreateKey` API.
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Specifies the algorithm for key material encryption. Currently, `RSAES_PKCS1_V1_5`, `RSAES_OAEP_SHA_1`, and `RSAES_OAEP_SHA_256` are supported.
	WrappingAlgorithm *string `json:"WrappingAlgorithm,omitempty" name:"WrappingAlgorithm"`

	// Specifies the type of wrapping key. Currently, only `RSA_2048` is supported.
	WrappingKeySpec *string `json:"WrappingKeySpec,omitempty" name:"WrappingKeySpec"`
}

func (r *GetParametersForImportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetParametersForImportRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetParametersForImportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Unique ID of a CMK, which is used to specify the CMK into which to import key material.
		KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

		// The token required for importing key material, which is used as a parameter for `ImportKeyMaterial`.
		ImportToken *string `json:"ImportToken,omitempty" name:"ImportToken"`

		// The Base64-encoded RSA public key used to encrypt key material before importing it with `ImportKeyMaterial`.
		PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`

		// Validity period of the token and public key. A token and public key can only be imported when they are valid. If they are expired, you will need to call the `GetParametersForImport` API again to get a new token and public key.
		ParametersValidTo *uint64 `json:"ParametersValidTo,omitempty" name:"ParametersValidTo"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetParametersForImportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetParametersForImportResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPublicKeyRequest struct {
	*tchttp.BaseRequest

	// The unique identifier of the CMK.
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *GetPublicKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPublicKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPublicKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// CThe unique identifier of the CMK.
		KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

		// Base64-encoded public key content.
		PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`

		// Public key content in PEM format.
		PublicKeyPem *string `json:"PublicKeyPem,omitempty" name:"PublicKeyPem"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPublicKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPublicKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetServiceStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *GetServiceStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetServiceStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetServiceStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetServiceStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetServiceStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImportKeyMaterialRequest struct {
	*tchttp.BaseRequest

	// Base64-encoded key material that encrypted with the `PublicKey` returned by `GetParametersForImport`. For the KMS of SM-CRYPTO version, the length of the key material should be 128 bits, while for KMS of FIPS-compliant version, the length should be 256 bits.
	EncryptedKeyMaterial *string `json:"EncryptedKeyMaterial,omitempty" name:"EncryptedKeyMaterial"`

	// Import token obtained by calling `GetParametersForImport`.
	ImportToken *string `json:"ImportToken,omitempty" name:"ImportToken"`

	// Specifies the CMK into which to import key material, which must be the same as the one specified by `GetParametersForImport`.
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Unix timestamp of the key material’s expiration time. If this value is empty or 0, the key material will never expire. To specify the expiration time, it should be later than the current time. Maximum value: 2147443200.
	ValidTo *uint64 `json:"ValidTo,omitempty" name:"ValidTo"`
}

func (r *ImportKeyMaterialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImportKeyMaterialRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImportKeyMaterialResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportKeyMaterialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImportKeyMaterialResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type KeyMetadata struct {

	// 
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// 
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 
	Description *string `json:"Description,omitempty" name:"Description"`

	// CMK status. Valid values: Enabled, Disabled, PendingDelete, PendingImport.
	KeyState *string `json:"KeyState,omitempty" name:"KeyState"`

	// CMK usage, the value is: ENCRYPT_DECRYPT | ASYMMETRIC_DECRYPT_RSA_2048 | ASYMMETRIC_DECRYPT_SM2
	KeyUsage *string `json:"KeyUsage,omitempty" name:"KeyUsage"`

	// CMK type. 2: FIPS-compliant; 4: SM-CRYPTO
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 
	CreatorUin *uint64 `json:"CreatorUin,omitempty" name:"CreatorUin"`

	// 
	KeyRotationEnabled *bool `json:"KeyRotationEnabled,omitempty" name:"KeyRotationEnabled"`

	// 
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// 
	NextRotateTime *uint64 `json:"NextRotateTime,omitempty" name:"NextRotateTime"`

	// 
	DeletionDate *uint64 `json:"DeletionDate,omitempty" name:"DeletionDate"`

	// CMK key material type. TENCENT_KMS: created by KMS; EXTERNAL: imported by user.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Origin *string `json:"Origin,omitempty" name:"Origin"`

	// It’s valid when `Origin` is `EXTERNAL`, indicating the expiration date of key material. 0 means valid forever.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ValidTo *uint64 `json:"ValidTo,omitempty" name:"ValidTo"`
}

type ListAlgorithmsRequest struct {
	*tchttp.BaseRequest
}

func (r *ListAlgorithmsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListAlgorithmsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListAlgorithmsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Symmetric encryption algorithms supported in this region
		SymmetricAlgorithms []*AlgorithmInfo `json:"SymmetricAlgorithms,omitempty" name:"SymmetricAlgorithms" list`

		// Asymmetric encryption algorithms supported in this region
		AsymmetricAlgorithms []*AlgorithmInfo `json:"AsymmetricAlgorithms,omitempty" name:"AsymmetricAlgorithms" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAlgorithmsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListAlgorithmsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListKeyDetailRequest struct {
	*tchttp.BaseRequest

	// 
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// The meaning is the same as the Limit of the SQL query, which means that a maximum of Limit elements are obtained this time. The default value is 10 and the maximum value is 200
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 
	Role *uint64 `json:"Role,omitempty" name:"Role"`

	// 
	OrderType *uint64 `json:"OrderType,omitempty" name:"OrderType"`

	// Filters by CMK status. 0: all CMKs; 1: CMKs in `Enabled` status only; 2: CMKs in `Disabled` status only; 3: CMKs in `PendingDelete` status only (i.e., keys with schedule deletion enabled); 4: CMKs in `PendingImport` status only.
	KeyState *uint64 `json:"KeyState,omitempty" name:"KeyState"`

	// 
	SearchKeyAlias *string `json:"SearchKeyAlias,omitempty" name:"SearchKeyAlias"`

	// Filters by CMK type. "TENCENT_KMS" indicates to filter CMKs whose key materials are created by KMS; "EXTERNAL" indicates to filter CMKs of `EXTERNAL` type whose key materials are imported by users; "ALL" or empty indicates to filter CMKs of both types. This value is case-sensitive.
	Origin *string `json:"Origin,omitempty" name:"Origin"`

	// Filter according to CMK's KeyUsage. If it is empty, it means to filter all. The available parameters are: ENCRYPT_DECRYPT or ASYMMETRIC_DECRYPT_RSA_2048 or ASYMMETRIC_DECRYPT_SM2.
	KeyUsage *string `json:"KeyUsage,omitempty" name:"KeyUsage"`
}

func (r *ListKeyDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListKeyDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListKeyDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of attribute information returned. 
	//  Note: This field may return null, which means that a valid value cannot be taken.
		KeyMetadatas []*KeyMetadata `json:"KeyMetadatas,omitempty" name:"KeyMetadatas" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListKeyDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListKeyDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListKeysRequest struct {
	*tchttp.BaseRequest
}

func (r *ListKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListKeysRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListKeysResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListKeysResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReEncryptRequest struct {
	*tchttp.BaseRequest
}

func (r *ReEncryptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReEncryptRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReEncryptResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReEncryptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReEncryptResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ScheduleKeyDeletionRequest struct {
	*tchttp.BaseRequest
}

func (r *ScheduleKeyDeletionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ScheduleKeyDeletionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ScheduleKeyDeletionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ScheduleKeyDeletionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ScheduleKeyDeletionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateAliasRequest struct {
	*tchttp.BaseRequest
}

func (r *UpdateAliasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateAliasRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateAliasResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateAliasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateAliasResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateKeyDescriptionRequest struct {
	*tchttp.BaseRequest

	// 
	Description *string `json:"Description,omitempty" name:"Description"`

	// CMK ID that needs to be modified
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *UpdateKeyDescriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateKeyDescriptionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateKeyDescriptionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateKeyDescriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateKeyDescriptionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
