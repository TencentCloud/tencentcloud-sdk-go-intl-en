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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AlgorithmInfo struct {
	// Algorithm ID
	KeyUsage *string `json:"KeyUsage,omitempty" name:"KeyUsage"`

	// Algorithm name
	Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`
}

// Predefined struct for user
type ArchiveKeyRequestParams struct {
	// Unique CMK ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

type ArchiveKeyRequest struct {
	*tchttp.BaseRequest
	
	// Unique CMK ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *ArchiveKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ArchiveKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ArchiveKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ArchiveKeyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ArchiveKeyResponse struct {
	*tchttp.BaseResponse
	Response *ArchiveKeyResponseParams `json:"Response"`
}

func (r *ArchiveKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ArchiveKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AsymmetricRsaDecryptRequestParams struct {
	// Unique CMK ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Base64-encoded ciphertext encrypted with `PublicKey`
	Ciphertext *string `json:"Ciphertext,omitempty" name:"Ciphertext"`

	// Corresponding algorithm when a public key is used for encryption. Valid values: RSAES_PKCS1_V1_5, RSAES_OAEP_SHA_1, RSAES_OAEP_SHA_256
	Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`
}

type AsymmetricRsaDecryptRequest struct {
	*tchttp.BaseRequest
	
	// Unique CMK ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Base64-encoded ciphertext encrypted with `PublicKey`
	Ciphertext *string `json:"Ciphertext,omitempty" name:"Ciphertext"`

	// Corresponding algorithm when a public key is used for encryption. Valid values: RSAES_PKCS1_V1_5, RSAES_OAEP_SHA_1, RSAES_OAEP_SHA_256
	Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`
}

func (r *AsymmetricRsaDecryptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AsymmetricRsaDecryptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	delete(f, "Ciphertext")
	delete(f, "Algorithm")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AsymmetricRsaDecryptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AsymmetricRsaDecryptResponseParams struct {
	// Unique CMK ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Base64-encoded plaintext after decryption
	Plaintext *string `json:"Plaintext,omitempty" name:"Plaintext"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AsymmetricRsaDecryptResponse struct {
	*tchttp.BaseResponse
	Response *AsymmetricRsaDecryptResponseParams `json:"Response"`
}

func (r *AsymmetricRsaDecryptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AsymmetricRsaDecryptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AsymmetricSm2DecryptRequestParams struct {
	// Unique CMK ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Base64-encoded ciphertext encrypted with `PublicKey`, whose length cannot exceed 256 bytes.
	Ciphertext *string `json:"Ciphertext,omitempty" name:"Ciphertext"`
}

type AsymmetricSm2DecryptRequest struct {
	*tchttp.BaseRequest
	
	// Unique CMK ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Base64-encoded ciphertext encrypted with `PublicKey`, whose length cannot exceed 256 bytes.
	Ciphertext *string `json:"Ciphertext,omitempty" name:"Ciphertext"`
}

func (r *AsymmetricSm2DecryptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AsymmetricSm2DecryptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	delete(f, "Ciphertext")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AsymmetricSm2DecryptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AsymmetricSm2DecryptResponseParams struct {
	// Unique CMK ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Base64-encoded plaintext after decryption
	Plaintext *string `json:"Plaintext,omitempty" name:"Plaintext"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AsymmetricSm2DecryptResponse struct {
	*tchttp.BaseResponse
	Response *AsymmetricSm2DecryptResponseParams `json:"Response"`
}

func (r *AsymmetricSm2DecryptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AsymmetricSm2DecryptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindCloudResourceRequestParams struct {
	// CMK ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Unique ID of a Tencent Cloud service
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// Resource/instance ID, which is stored as a string and defined by the caller based on the Tencent Cloud service’s features.
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

type BindCloudResourceRequest struct {
	*tchttp.BaseRequest
	
	// CMK ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Unique ID of a Tencent Cloud service
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// Resource/instance ID, which is stored as a string and defined by the caller based on the Tencent Cloud service’s features.
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

func (r *BindCloudResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindCloudResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	delete(f, "ProductId")
	delete(f, "ResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindCloudResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindCloudResourceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindCloudResourceResponse struct {
	*tchttp.BaseResponse
	Response *BindCloudResourceResponseParams `json:"Response"`
}

func (r *BindCloudResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindCloudResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelKeyArchiveRequestParams struct {
	// Unique CMK ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

type CancelKeyArchiveRequest struct {
	*tchttp.BaseRequest
	
	// Unique CMK ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *CancelKeyArchiveRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelKeyArchiveRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelKeyArchiveRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelKeyArchiveResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CancelKeyArchiveResponse struct {
	*tchttp.BaseResponse
	Response *CancelKeyArchiveResponseParams `json:"Response"`
}

func (r *CancelKeyArchiveResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelKeyArchiveResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelKeyDeletionRequestParams struct {
	// Unique ID of the CMK for which to cancel schedule deletion
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

type CancelKeyDeletionRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of the CMK for which to cancel schedule deletion
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *CancelKeyDeletionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelKeyDeletionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelKeyDeletionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelKeyDeletionResponseParams struct {
	// Unique ID of the CMK for which the schedule deletion is canceled
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CancelKeyDeletionResponse struct {
	*tchttp.BaseResponse
	Response *CancelKeyDeletionResponseParams `json:"Response"`
}

func (r *CancelKeyDeletionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelKeyDeletionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateKeyRequestParams struct {
	// Unique alias that makes a key more recognizable and understandable. This parameter cannot be empty, can contain 1-60 letters, digits, `-`, and `_`, and must begin with a letter or digit. The `kms-` prefix is used for Tencent Cloud products.
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// CMK description of up to 1,024 bytes in length
	Description *string `json:"Description,omitempty" name:"Description"`

	// Defines the purpose of the key. The valid values are as follows: `ENCRYPT_DECRYPT` (default): creates a symmetric encryption/decryption key; `ASYMMETRIC_DECRYPT_RSA_2048`: creates an asymmetric encryption/decryption 2048-bit RSA key; `ASYMMETRIC_DECRYPT_SM2`: creates an asymmetric encryption/decryption SM2 key; `ASYMMETRIC_SIGN_VERIFY_SM2`: creates an asymmetric SM2 key for signature verification; `ASYMMETRIC_SIGN_VERIFY_ECC`: creates an asymmetric 2048-bit RSA key for signature verification; `ASYMMETRIC_SIGN_VERIFY_ECDSA384`: creates an asymmetric ECDSA384 key for signature verification. You can get a full list of supported key purposes and algorithms using the ListAlgorithms API.
	KeyUsage *string `json:"KeyUsage,omitempty" name:"KeyUsage"`

	// Specifies the key type. Default value: 1. Valid value: 1 - default type, indicating that the CMK is created by KMS; 2 - EXTERNAL type, indicating that you need to import key material. For more information, please see the `GetParametersForImport` and `ImportKeyMaterial` API documents.
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// Tag list
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type CreateKeyRequest struct {
	*tchttp.BaseRequest
	
	// Unique alias that makes a key more recognizable and understandable. This parameter cannot be empty, can contain 1-60 letters, digits, `-`, and `_`, and must begin with a letter or digit. The `kms-` prefix is used for Tencent Cloud products.
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// CMK description of up to 1,024 bytes in length
	Description *string `json:"Description,omitempty" name:"Description"`

	// Defines the purpose of the key. The valid values are as follows: `ENCRYPT_DECRYPT` (default): creates a symmetric encryption/decryption key; `ASYMMETRIC_DECRYPT_RSA_2048`: creates an asymmetric encryption/decryption 2048-bit RSA key; `ASYMMETRIC_DECRYPT_SM2`: creates an asymmetric encryption/decryption SM2 key; `ASYMMETRIC_SIGN_VERIFY_SM2`: creates an asymmetric SM2 key for signature verification; `ASYMMETRIC_SIGN_VERIFY_ECC`: creates an asymmetric 2048-bit RSA key for signature verification; `ASYMMETRIC_SIGN_VERIFY_ECDSA384`: creates an asymmetric ECDSA384 key for signature verification. You can get a full list of supported key purposes and algorithms using the ListAlgorithms API.
	KeyUsage *string `json:"KeyUsage,omitempty" name:"KeyUsage"`

	// Specifies the key type. Default value: 1. Valid value: 1 - default type, indicating that the CMK is created by KMS; 2 - EXTERNAL type, indicating that you need to import key material. For more information, please see the `GetParametersForImport` and `ImportKeyMaterial` API documents.
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// Tag list
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreateKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Alias")
	delete(f, "Description")
	delete(f, "KeyUsage")
	delete(f, "Type")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateKeyResponseParams struct {
	// Globally unique CMK ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Alias that makes a key more recognizable and understandable
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// Key creation time in UNIX timestamp format
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// CMK description
	Description *string `json:"Description,omitempty" name:"Description"`

	// CMK status
	KeyState *string `json:"KeyState,omitempty" name:"KeyState"`

	// CMK usage
	KeyUsage *string `json:"KeyUsage,omitempty" name:"KeyUsage"`

	// Tag operation return code. 0: success; 1: internal error; 2: business processing error
	TagCode *uint64 `json:"TagCode,omitempty" name:"TagCode"`

	// Tag operation return information
	TagMsg *string `json:"TagMsg,omitempty" name:"TagMsg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateKeyResponse struct {
	*tchttp.BaseResponse
	Response *CreateKeyResponseParams `json:"Response"`
}

func (r *CreateKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWhiteBoxKeyRequestParams struct {
	// Unique alias that makes a key more recognizable and understandable. This parameter should be 1 to 60 letters, digits, `-`, and `_`; it must begin with a letter or digit and cannot be left empty.
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// All algorithm types for creating keys. Valid values: AES_256, SM4
	Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`

	// Key description of up to 1024 bytes
	Description *string `json:"Description,omitempty" name:"Description"`

	// Tag list
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type CreateWhiteBoxKeyRequest struct {
	*tchttp.BaseRequest
	
	// Unique alias that makes a key more recognizable and understandable. This parameter should be 1 to 60 letters, digits, `-`, and `_`; it must begin with a letter or digit and cannot be left empty.
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// All algorithm types for creating keys. Valid values: AES_256, SM4
	Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`

	// Key description of up to 1024 bytes
	Description *string `json:"Description,omitempty" name:"Description"`

	// Tag list
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreateWhiteBoxKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWhiteBoxKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Alias")
	delete(f, "Algorithm")
	delete(f, "Description")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWhiteBoxKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWhiteBoxKeyResponseParams struct {
	// Base64-encoded encryption key
	EncryptKey *string `json:"EncryptKey,omitempty" name:"EncryptKey"`

	// Base64-encoded decryption key
	DecryptKey *string `json:"DecryptKey,omitempty" name:"DecryptKey"`

	// Globally unique white-box key ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Tag operation return code. 0: success; 1: internal error; 2: business processing error
	TagCode *uint64 `json:"TagCode,omitempty" name:"TagCode"`

	// Tag operation return message
	TagMsg *string `json:"TagMsg,omitempty" name:"TagMsg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateWhiteBoxKeyResponse struct {
	*tchttp.BaseResponse
	Response *CreateWhiteBoxKeyResponseParams `json:"Response"`
}

func (r *CreateWhiteBoxKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWhiteBoxKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DecryptRequestParams struct {
	// The ciphertext data to be decrypted.
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" name:"CiphertextBlob"`

	// JSON string of key-value pair. If this parameter is specified for `Encrypt`, the same parameter needs to be provided when the `Decrypt` API is called. The maximum length is 1,024 bytes.
	EncryptionContext *string `json:"EncryptionContext,omitempty" name:"EncryptionContext"`

	// PEM-encoded public key (2048-bit RSA/SM2 key), which can be used to encrypt the `Plaintext` returned. If this field is left empty, the `Plaintext` will not be encrypted.
	EncryptionPublicKey *string `json:"EncryptionPublicKey,omitempty" name:"EncryptionPublicKey"`

	// Asymmetric encryption algorithm. Valid values: `SM2(C1C3C2)`, `RSAES_PKCS1_V1_5`, `RSAES_OAEP_SHA_1`, and `RSAES_OAEP_SHA_256`. This field is used in combination with `EncryptionPublicKey` for encryption. If it is left empty, a SM2 public key will be used by default.
	EncryptionAlgorithm *string `json:"EncryptionAlgorithm,omitempty" name:"EncryptionAlgorithm"`
}

type DecryptRequest struct {
	*tchttp.BaseRequest
	
	// The ciphertext data to be decrypted.
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" name:"CiphertextBlob"`

	// JSON string of key-value pair. If this parameter is specified for `Encrypt`, the same parameter needs to be provided when the `Decrypt` API is called. The maximum length is 1,024 bytes.
	EncryptionContext *string `json:"EncryptionContext,omitempty" name:"EncryptionContext"`

	// PEM-encoded public key (2048-bit RSA/SM2 key), which can be used to encrypt the `Plaintext` returned. If this field is left empty, the `Plaintext` will not be encrypted.
	EncryptionPublicKey *string `json:"EncryptionPublicKey,omitempty" name:"EncryptionPublicKey"`

	// Asymmetric encryption algorithm. Valid values: `SM2(C1C3C2)`, `RSAES_PKCS1_V1_5`, `RSAES_OAEP_SHA_1`, and `RSAES_OAEP_SHA_256`. This field is used in combination with `EncryptionPublicKey` for encryption. If it is left empty, a SM2 public key will be used by default.
	EncryptionAlgorithm *string `json:"EncryptionAlgorithm,omitempty" name:"EncryptionAlgorithm"`
}

func (r *DecryptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DecryptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CiphertextBlob")
	delete(f, "EncryptionContext")
	delete(f, "EncryptionPublicKey")
	delete(f, "EncryptionAlgorithm")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DecryptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DecryptResponseParams struct {
	// Globally unique CMK ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// If `EncryptionPublicKey` is left empty, a Base64-encoded ciphertext will be returned. To get the plaintext, you need to decode the ciphertext first.
	// If `EncryptionPublicKey` is specified, this field will return the Base64-encoded ciphertext encrypted with the specified public key. To get the plaintext, you need to decode the ciphertext and upload the corresponding private key.
	Plaintext *string `json:"Plaintext,omitempty" name:"Plaintext"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DecryptResponse struct {
	*tchttp.BaseResponse
	Response *DecryptResponseParams `json:"Response"`
}

func (r *DecryptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DecryptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImportedKeyMaterialRequestParams struct {
	// Specifies the EXTERNAL CMK for which to delete the key material.
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImportedKeyMaterialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteImportedKeyMaterialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImportedKeyMaterialResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteImportedKeyMaterialResponse struct {
	*tchttp.BaseResponse
	Response *DeleteImportedKeyMaterialResponseParams `json:"Response"`
}

func (r *DeleteImportedKeyMaterialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImportedKeyMaterialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWhiteBoxKeyRequestParams struct {
	// Globally unique white-box key ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

type DeleteWhiteBoxKeyRequest struct {
	*tchttp.BaseRequest
	
	// Globally unique white-box key ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *DeleteWhiteBoxKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWhiteBoxKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWhiteBoxKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWhiteBoxKeyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteWhiteBoxKeyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWhiteBoxKeyResponseParams `json:"Response"`
}

func (r *DeleteWhiteBoxKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWhiteBoxKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKeyRequestParams struct {
	// Globally unique CMK ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

type DescribeKeyRequest struct {
	*tchttp.BaseRequest
	
	// Globally unique CMK ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *DescribeKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKeyResponseParams struct {
	// Key attribute information
	// Note: this field may return null, indicating that no valid values can be obtained.
	KeyMetadata *KeyMetadata `json:"KeyMetadata,omitempty" name:"KeyMetadata"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeKeyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKeyResponseParams `json:"Response"`
}

func (r *DescribeKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKeysRequestParams struct {
	// List of IDs of the CMKs to be queried in batches. Up to 100 `KeyId` values are supported in one query.
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
}

type DescribeKeysRequest struct {
	*tchttp.BaseRequest
	
	// List of IDs of the CMKs to be queried in batches. Up to 100 `KeyId` values are supported in one query.
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
}

func (r *DescribeKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKeysResponseParams struct {
	// List of returned attribute information
	// Note: this field may return null, indicating that no valid values can be obtained.
	KeyMetadatas []*KeyMetadata `json:"KeyMetadatas,omitempty" name:"KeyMetadatas"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeKeysResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKeysResponseParams `json:"Response"`
}

func (r *DescribeKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhiteBoxDecryptKeyRequestParams struct {
	// Globally unique white-box key ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

type DescribeWhiteBoxDecryptKeyRequest struct {
	*tchttp.BaseRequest
	
	// Globally unique white-box key ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *DescribeWhiteBoxDecryptKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhiteBoxDecryptKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWhiteBoxDecryptKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhiteBoxDecryptKeyResponseParams struct {
	// Base64-encoded white-box decryption key
	DecryptKey *string `json:"DecryptKey,omitempty" name:"DecryptKey"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWhiteBoxDecryptKeyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWhiteBoxDecryptKeyResponseParams `json:"Response"`
}

func (r *DescribeWhiteBoxDecryptKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhiteBoxDecryptKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhiteBoxDeviceFingerprintsRequestParams struct {
	// White-box key ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

type DescribeWhiteBoxDeviceFingerprintsRequest struct {
	*tchttp.BaseRequest
	
	// White-box key ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *DescribeWhiteBoxDeviceFingerprintsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhiteBoxDeviceFingerprintsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWhiteBoxDeviceFingerprintsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhiteBoxDeviceFingerprintsResponseParams struct {
	// Device fingerprint list
	DeviceFingerprints []*DeviceFingerprint `json:"DeviceFingerprints,omitempty" name:"DeviceFingerprints"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWhiteBoxDeviceFingerprintsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWhiteBoxDeviceFingerprintsResponseParams `json:"Response"`
}

func (r *DescribeWhiteBoxDeviceFingerprintsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhiteBoxDeviceFingerprintsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhiteBoxKeyDetailsRequestParams struct {
	// Filter: key status. 0: disabled, 1: enabled
	KeyStatus *int64 `json:"KeyStatus,omitempty" name:"KeyStatus"`

	// This parameter has the same meaning of the `Offset` in an SQL query, indicating that this acquisition starts from the "No. Offset value" element of the array arranged in a certain order. The default value is 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// This parameter has the same meaning of the `Limit` in an SQL query, indicating that up to `Limit` value elements can be obtained in this request. The default value is 0, indicating not to paginate.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Tag filter condition
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`
}

type DescribeWhiteBoxKeyDetailsRequest struct {
	*tchttp.BaseRequest
	
	// Filter: key status. 0: disabled, 1: enabled
	KeyStatus *int64 `json:"KeyStatus,omitempty" name:"KeyStatus"`

	// This parameter has the same meaning of the `Offset` in an SQL query, indicating that this acquisition starts from the "No. Offset value" element of the array arranged in a certain order. The default value is 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// This parameter has the same meaning of the `Limit` in an SQL query, indicating that up to `Limit` value elements can be obtained in this request. The default value is 0, indicating not to paginate.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Tag filter condition
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`
}

func (r *DescribeWhiteBoxKeyDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhiteBoxKeyDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyStatus")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "TagFilters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWhiteBoxKeyDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhiteBoxKeyDetailsResponseParams struct {
	// White-box key information list
	KeyInfos []*WhiteboxKeyInfo `json:"KeyInfos,omitempty" name:"KeyInfos"`

	// Total number of keys
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWhiteBoxKeyDetailsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWhiteBoxKeyDetailsResponseParams `json:"Response"`
}

func (r *DescribeWhiteBoxKeyDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhiteBoxKeyDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhiteBoxKeyRequestParams struct {
	// Globally unique white-box key ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

type DescribeWhiteBoxKeyRequest struct {
	*tchttp.BaseRequest
	
	// Globally unique white-box key ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *DescribeWhiteBoxKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhiteBoxKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWhiteBoxKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhiteBoxKeyResponseParams struct {
	// White-box key information
	KeyInfo *WhiteboxKeyInfo `json:"KeyInfo,omitempty" name:"KeyInfo"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWhiteBoxKeyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWhiteBoxKeyResponseParams `json:"Response"`
}

func (r *DescribeWhiteBoxKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhiteBoxKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhiteBoxServiceStatusRequestParams struct {

}

type DescribeWhiteBoxServiceStatusRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeWhiteBoxServiceStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhiteBoxServiceStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWhiteBoxServiceStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhiteBoxServiceStatusResponseParams struct {
	// Whether the user's white-box key service is available
	ServiceEnabled *bool `json:"ServiceEnabled,omitempty" name:"ServiceEnabled"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWhiteBoxServiceStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWhiteBoxServiceStatusResponseParams `json:"Response"`
}

func (r *DescribeWhiteBoxServiceStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhiteBoxServiceStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeviceFingerprint struct {
	// Fingerprint information collected by device fingerprint collector. Its format must be in the following regular expression: ^[0-9a-f]{8}[\-][0-9a-f]{14}[\-][0-9a-f]{14}[\-][0-9a-f]{14}[\-][0-9a-f]{16}$
	Identity *string `json:"Identity,omitempty" name:"Identity"`

	// Description, such as IP and device name. Length limit: 1,024 bytes
	// Note: this field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitempty" name:"Description"`
}

// Predefined struct for user
type DisableKeyRequestParams struct {
	// Unique CMK ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

type DisableKeyRequest struct {
	*tchttp.BaseRequest
	
	// Unique CMK ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *DisableKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableKeyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisableKeyResponse struct {
	*tchttp.BaseResponse
	Response *DisableKeyResponseParams `json:"Response"`
}

func (r *DisableKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableKeyRotationRequestParams struct {
	// Unique CMK ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

type DisableKeyRotationRequest struct {
	*tchttp.BaseRequest
	
	// Unique CMK ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *DisableKeyRotationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableKeyRotationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableKeyRotationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableKeyRotationResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisableKeyRotationResponse struct {
	*tchttp.BaseResponse
	Response *DisableKeyRotationResponseParams `json:"Response"`
}

func (r *DisableKeyRotationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableKeyRotationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableKeysRequestParams struct {
	// List of IDs of the CMKs to be disabled in batches. Up to 100 CMKs are supported at a time
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
}

type DisableKeysRequest struct {
	*tchttp.BaseRequest
	
	// List of IDs of the CMKs to be disabled in batches. Up to 100 CMKs are supported at a time
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
}

func (r *DisableKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableKeysResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisableKeysResponse struct {
	*tchttp.BaseResponse
	Response *DisableKeysResponseParams `json:"Response"`
}

func (r *DisableKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableWhiteBoxKeyRequestParams struct {
	// Globally unique white-box key ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

type DisableWhiteBoxKeyRequest struct {
	*tchttp.BaseRequest
	
	// Globally unique white-box key ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *DisableWhiteBoxKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableWhiteBoxKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableWhiteBoxKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableWhiteBoxKeyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisableWhiteBoxKeyResponse struct {
	*tchttp.BaseResponse
	Response *DisableWhiteBoxKeyResponseParams `json:"Response"`
}

func (r *DisableWhiteBoxKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableWhiteBoxKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableWhiteBoxKeysRequestParams struct {
	// List of globally unique white-box key IDs. Note: you should make sure that all provided `KeyId` values are in valid format, unique, and actually exist. Up to 50 ones are allowed.
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
}

type DisableWhiteBoxKeysRequest struct {
	*tchttp.BaseRequest
	
	// List of globally unique white-box key IDs. Note: you should make sure that all provided `KeyId` values are in valid format, unique, and actually exist. Up to 50 ones are allowed.
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
}

func (r *DisableWhiteBoxKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableWhiteBoxKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableWhiteBoxKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableWhiteBoxKeysResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisableWhiteBoxKeysResponse struct {
	*tchttp.BaseResponse
	Response *DisableWhiteBoxKeysResponseParams `json:"Response"`
}

func (r *DisableWhiteBoxKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableWhiteBoxKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableKeyRequestParams struct {
	// Unique CMK ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

type EnableKeyRequest struct {
	*tchttp.BaseRequest
	
	// Unique CMK ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *EnableKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableKeyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EnableKeyResponse struct {
	*tchttp.BaseResponse
	Response *EnableKeyResponseParams `json:"Response"`
}

func (r *EnableKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableKeyRotationRequestParams struct {
	// Unique CMK ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

type EnableKeyRotationRequest struct {
	*tchttp.BaseRequest
	
	// Unique CMK ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *EnableKeyRotationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableKeyRotationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableKeyRotationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableKeyRotationResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EnableKeyRotationResponse struct {
	*tchttp.BaseResponse
	Response *EnableKeyRotationResponseParams `json:"Response"`
}

func (r *EnableKeyRotationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableKeyRotationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableKeysRequestParams struct {
	// List of IDs of the CMKs to be enabled in batches. Up to 100 CMKs are supported at a time
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
}

type EnableKeysRequest struct {
	*tchttp.BaseRequest
	
	// List of IDs of the CMKs to be enabled in batches. Up to 100 CMKs are supported at a time
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
}

func (r *EnableKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableKeysResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EnableKeysResponse struct {
	*tchttp.BaseResponse
	Response *EnableKeysResponseParams `json:"Response"`
}

func (r *EnableKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableWhiteBoxKeyRequestParams struct {
	// Globally unique white-box key ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

type EnableWhiteBoxKeyRequest struct {
	*tchttp.BaseRequest
	
	// Globally unique white-box key ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *EnableWhiteBoxKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableWhiteBoxKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableWhiteBoxKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableWhiteBoxKeyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EnableWhiteBoxKeyResponse struct {
	*tchttp.BaseResponse
	Response *EnableWhiteBoxKeyResponseParams `json:"Response"`
}

func (r *EnableWhiteBoxKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableWhiteBoxKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableWhiteBoxKeysRequestParams struct {
	// List of globally unique white-box key IDs. Note: you should make sure that all provided `KeyId` values are in valid format, unique, and actually exist. Up to 50 ones are allowed.
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
}

type EnableWhiteBoxKeysRequest struct {
	*tchttp.BaseRequest
	
	// List of globally unique white-box key IDs. Note: you should make sure that all provided `KeyId` values are in valid format, unique, and actually exist. Up to 50 ones are allowed.
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
}

func (r *EnableWhiteBoxKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableWhiteBoxKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableWhiteBoxKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableWhiteBoxKeysResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EnableWhiteBoxKeysResponse struct {
	*tchttp.BaseResponse
	Response *EnableWhiteBoxKeysResponseParams `json:"Response"`
}

func (r *EnableWhiteBoxKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableWhiteBoxKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EncryptByWhiteBoxRequestParams struct {
	// Globally unique white-box key ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Base64-encoded text to be encrypted. The size of the original text cannot exceed 4 KB.
	PlainText *string `json:"PlainText,omitempty" name:"PlainText"`

	// Base64-encoded initialization vector of 16 bytes, which will be used by the encryption algorithm. If this parameter is not passed in, the backend service will generate a random one. You should save this value as a parameter for decryption.
	InitializationVector *string `json:"InitializationVector,omitempty" name:"InitializationVector"`
}

type EncryptByWhiteBoxRequest struct {
	*tchttp.BaseRequest
	
	// Globally unique white-box key ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Base64-encoded text to be encrypted. The size of the original text cannot exceed 4 KB.
	PlainText *string `json:"PlainText,omitempty" name:"PlainText"`

	// Base64-encoded initialization vector of 16 bytes, which will be used by the encryption algorithm. If this parameter is not passed in, the backend service will generate a random one. You should save this value as a parameter for decryption.
	InitializationVector *string `json:"InitializationVector,omitempty" name:"InitializationVector"`
}

func (r *EncryptByWhiteBoxRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EncryptByWhiteBoxRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	delete(f, "PlainText")
	delete(f, "InitializationVector")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EncryptByWhiteBoxRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EncryptByWhiteBoxResponseParams struct {
	// Base64-encoded initialization vector, which will be used by the encryption algorithm. If this parameter is passed in by the caller, it will be returned as-is; otherwise, the backend service will generate a random one and return it.
	InitializationVector *string `json:"InitializationVector,omitempty" name:"InitializationVector"`

	// Base64-encoded ciphertext after encryption
	CipherText *string `json:"CipherText,omitempty" name:"CipherText"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EncryptByWhiteBoxResponse struct {
	*tchttp.BaseResponse
	Response *EncryptByWhiteBoxResponseParams `json:"Response"`
}

func (r *EncryptByWhiteBoxResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EncryptByWhiteBoxResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EncryptRequestParams struct {
	// Globally unique ID of the CMK generated by calling the `CreateKey` API
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Encrypted plaintext data. This field must be Base64-encoded. The maximum size of the original data is 4 KB
	Plaintext *string `json:"Plaintext,omitempty" name:"Plaintext"`

	// JSON string of key-value pair. If this parameter is specified, the same parameter needs to be provided when the `Decrypt` API is called. It is up to 1,024 characters
	EncryptionContext *string `json:"EncryptionContext,omitempty" name:"EncryptionContext"`
}

type EncryptRequest struct {
	*tchttp.BaseRequest
	
	// Globally unique ID of the CMK generated by calling the `CreateKey` API
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Encrypted plaintext data. This field must be Base64-encoded. The maximum size of the original data is 4 KB
	Plaintext *string `json:"Plaintext,omitempty" name:"Plaintext"`

	// JSON string of key-value pair. If this parameter is specified, the same parameter needs to be provided when the `Decrypt` API is called. It is up to 1,024 characters
	EncryptionContext *string `json:"EncryptionContext,omitempty" name:"EncryptionContext"`
}

func (r *EncryptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EncryptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	delete(f, "Plaintext")
	delete(f, "EncryptionContext")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EncryptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EncryptResponseParams struct {
	// Base64-encoded ciphertext, which is the encrypted information of the ciphertext and key. To get the plaintext, you need to pass in this field to the Decrypt API.
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" name:"CiphertextBlob"`

	// Globally unique ID of the CMK used for encryption
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EncryptResponse struct {
	*tchttp.BaseResponse
	Response *EncryptResponseParams `json:"Response"`
}

func (r *EncryptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EncryptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateDataKeyRequestParams struct {
	// Globally unique CMK ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Specifies the encryption algorithm and size of the `DataKey`. Valid values: AES_128, AES_256. Either `KeySpec` or `NumberOfBytes` must be specified.
	KeySpec *string `json:"KeySpec,omitempty" name:"KeySpec"`

	// Length of the `DataKey`. If both `NumberOfBytes` and `KeySpec` are specified, `NumberOfBytes` will prevail. Minimum value: 1; maximum value: 1024. Either `KeySpec` or `NumberOfBytes` must be specified.
	NumberOfBytes *uint64 `json:"NumberOfBytes,omitempty" name:"NumberOfBytes"`

	// JSON string of key-value pair. If this field is used, the same string should be entered when the returned `DataKey` is decrypted.
	EncryptionContext *string `json:"EncryptionContext,omitempty" name:"EncryptionContext"`

	// PEM-encoded public key (2048-bit RSA/SM2 key), which can be used to encrypt the `Plaintext` returned. If this field is left empty, the `Plaintext` will not be encrypted.
	EncryptionPublicKey *string `json:"EncryptionPublicKey,omitempty" name:"EncryptionPublicKey"`

	// Asymmetric encryption algorithm. Valid values: `SM2(C1C3C2)`, `RSAES_PKCS1_V1_5`, `RSAES_OAEP_SHA_1`, and `RSAES_OAEP_SHA_256`. This field is used with `EncryptionPublicKey` for encryption. If it is left empty, a SM2 public key will be used by default.
	EncryptionAlgorithm *string `json:"EncryptionAlgorithm,omitempty" name:"EncryptionAlgorithm"`
}

type GenerateDataKeyRequest struct {
	*tchttp.BaseRequest
	
	// Globally unique CMK ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Specifies the encryption algorithm and size of the `DataKey`. Valid values: AES_128, AES_256. Either `KeySpec` or `NumberOfBytes` must be specified.
	KeySpec *string `json:"KeySpec,omitempty" name:"KeySpec"`

	// Length of the `DataKey`. If both `NumberOfBytes` and `KeySpec` are specified, `NumberOfBytes` will prevail. Minimum value: 1; maximum value: 1024. Either `KeySpec` or `NumberOfBytes` must be specified.
	NumberOfBytes *uint64 `json:"NumberOfBytes,omitempty" name:"NumberOfBytes"`

	// JSON string of key-value pair. If this field is used, the same string should be entered when the returned `DataKey` is decrypted.
	EncryptionContext *string `json:"EncryptionContext,omitempty" name:"EncryptionContext"`

	// PEM-encoded public key (2048-bit RSA/SM2 key), which can be used to encrypt the `Plaintext` returned. If this field is left empty, the `Plaintext` will not be encrypted.
	EncryptionPublicKey *string `json:"EncryptionPublicKey,omitempty" name:"EncryptionPublicKey"`

	// Asymmetric encryption algorithm. Valid values: `SM2(C1C3C2)`, `RSAES_PKCS1_V1_5`, `RSAES_OAEP_SHA_1`, and `RSAES_OAEP_SHA_256`. This field is used with `EncryptionPublicKey` for encryption. If it is left empty, a SM2 public key will be used by default.
	EncryptionAlgorithm *string `json:"EncryptionAlgorithm,omitempty" name:"EncryptionAlgorithm"`
}

func (r *GenerateDataKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateDataKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	delete(f, "KeySpec")
	delete(f, "NumberOfBytes")
	delete(f, "EncryptionContext")
	delete(f, "EncryptionPublicKey")
	delete(f, "EncryptionAlgorithm")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GenerateDataKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateDataKeyResponseParams struct {
	// Globally unique CMK ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// If `EncryptionPublicKey` is left empty, a Base64-encoded ciphertext will be returned. To get the plaintext, you need to decode the ciphertext first.
	// If `EncryptionPublicKey` is specified, this field will return the Base64-encoded ciphertext encrypted with the specified public key. To get the plaintext, you need to decode the ciphertext and upload the corresponding private key.
	Plaintext *string `json:"Plaintext,omitempty" name:"Plaintext"`

	// Ciphertext of the data key, which should be kept by yourself. KMS does not host user data keys. You can call the `Decrypt` API to get the plaintext of the data key from `CiphertextBlob`.
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" name:"CiphertextBlob"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GenerateDataKeyResponse struct {
	*tchttp.BaseResponse
	Response *GenerateDataKeyResponseParams `json:"Response"`
}

func (r *GenerateDataKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateDataKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateRandomRequestParams struct {
	// Length of the random number. Minimum value: 1. Maximum value: 1024
	NumberOfBytes *uint64 `json:"NumberOfBytes,omitempty" name:"NumberOfBytes"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateRandomRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NumberOfBytes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GenerateRandomRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateRandomResponseParams struct {
	// Base64-encoded plaintext of the randomly generated number. You need to Base64-decode it to get the plaintext.
	Plaintext *string `json:"Plaintext,omitempty" name:"Plaintext"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GenerateRandomResponse struct {
	*tchttp.BaseResponse
	Response *GenerateRandomResponseParams `json:"Response"`
}

func (r *GenerateRandomResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateRandomResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetKeyRotationStatusRequestParams struct {
	// Unique CMK ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

type GetKeyRotationStatusRequest struct {
	*tchttp.BaseRequest
	
	// Unique CMK ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *GetKeyRotationStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetKeyRotationStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetKeyRotationStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetKeyRotationStatusResponseParams struct {
	// Whether key rotation is enabled
	KeyRotationEnabled *bool `json:"KeyRotationEnabled,omitempty" name:"KeyRotationEnabled"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetKeyRotationStatusResponse struct {
	*tchttp.BaseResponse
	Response *GetKeyRotationStatusResponseParams `json:"Response"`
}

func (r *GetKeyRotationStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetKeyRotationStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetParametersForImportRequestParams struct {
	// Unique ID of a CMK. The CMK for which to get the key parameters must be of the `EXTERNAL` type, i.e., Type = 2 when the CMK is created by the `CreateKey` API.
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Specifies the algorithm for key material encryption. Currently, `RSAES_PKCS1_V1_5`, `RSAES_OAEP_SHA_1`, and `RSAES_OAEP_SHA_256` are supported.
	WrappingAlgorithm *string `json:"WrappingAlgorithm,omitempty" name:"WrappingAlgorithm"`

	// Specifies the type of wrapping key. Currently, only `RSA_2048` is supported.
	WrappingKeySpec *string `json:"WrappingKeySpec,omitempty" name:"WrappingKeySpec"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetParametersForImportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	delete(f, "WrappingAlgorithm")
	delete(f, "WrappingKeySpec")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetParametersForImportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetParametersForImportResponseParams struct {
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
}

type GetParametersForImportResponse struct {
	*tchttp.BaseResponse
	Response *GetParametersForImportResponseParams `json:"Response"`
}

func (r *GetParametersForImportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetParametersForImportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPublicKeyRequestParams struct {
	// Unique CMK ID.
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

type GetPublicKeyRequest struct {
	*tchttp.BaseRequest
	
	// Unique CMK ID.
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *GetPublicKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPublicKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetPublicKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPublicKeyResponseParams struct {
	// Unique CMK ID.
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Base64-encoded public key content.
	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`

	// Public key content in PEM format.
	PublicKeyPem *string `json:"PublicKeyPem,omitempty" name:"PublicKeyPem"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetPublicKeyResponse struct {
	*tchttp.BaseResponse
	Response *GetPublicKeyResponseParams `json:"Response"`
}

func (r *GetPublicKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPublicKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRegionsRequestParams struct {

}

type GetRegionsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *GetRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRegionsResponseParams struct {
	// The list of supported regions
	// Note: this field may return null, indicating that no valid values can be obtained.
	Regions []*string `json:"Regions,omitempty" name:"Regions"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetRegionsResponse struct {
	*tchttp.BaseResponse
	Response *GetRegionsResponseParams `json:"Response"`
}

func (r *GetRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetServiceStatusRequestParams struct {

}

type GetServiceStatusRequest struct {
	*tchttp.BaseRequest
	
}

func (r *GetServiceStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetServiceStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetServiceStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetServiceStatusResponseParams struct {
	// Whether the KMS service has been activated. true: activated
	ServiceEnabled *bool `json:"ServiceEnabled,omitempty" name:"ServiceEnabled"`

	// Service unavailability type. 0: not purchased; 1: normal; 2: suspended due to arrears; 3: resource released
	InvalidType *int64 `json:"InvalidType,omitempty" name:"InvalidType"`

	// 0: Basic Edition, 1: Ultimate Edition
	UserLevel *uint64 `json:"UserLevel,omitempty" name:"UserLevel"`

	// Ultimate Edition expiration time
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProExpireTime *uint64 `json:"ProExpireTime,omitempty" name:"ProExpireTime"`

	// Whether to automatically renew Ultimate Edition. 0: no, 1: yes
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProRenewFlag *uint64 `json:"ProRenewFlag,omitempty" name:"ProRenewFlag"`

	// Unique ID of the Ultimate Edition purchase record. If the Ultimate Edition is not activated, the returned value will be null.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProResourceId *string `json:"ProResourceId,omitempty" name:"ProResourceId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetServiceStatusResponse struct {
	*tchttp.BaseResponse
	Response *GetServiceStatusResponseParams `json:"Response"`
}

func (r *GetServiceStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetServiceStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportKeyMaterialRequestParams struct {
	// Base64-encoded key material that encrypted with the `PublicKey` returned by `GetParametersForImport`. For the KMS of SM-CRYPTO version, the length of the key material should be 128 bits, while for KMS of FIPS-compliant version, the length should be 256 bits.
	EncryptedKeyMaterial *string `json:"EncryptedKeyMaterial,omitempty" name:"EncryptedKeyMaterial"`

	// Import token obtained by calling `GetParametersForImport`.
	ImportToken *string `json:"ImportToken,omitempty" name:"ImportToken"`

	// Specifies the CMK into which to import key material, which must be the same as the one specified by `GetParametersForImport`.
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Unix timestamp of the key material's expiration time. If this value is empty or 0, the key material will never expire. To specify the expiration time, it should be later than the current time. Maximum value: 2147443200.
	ValidTo *uint64 `json:"ValidTo,omitempty" name:"ValidTo"`
}

type ImportKeyMaterialRequest struct {
	*tchttp.BaseRequest
	
	// Base64-encoded key material that encrypted with the `PublicKey` returned by `GetParametersForImport`. For the KMS of SM-CRYPTO version, the length of the key material should be 128 bits, while for KMS of FIPS-compliant version, the length should be 256 bits.
	EncryptedKeyMaterial *string `json:"EncryptedKeyMaterial,omitempty" name:"EncryptedKeyMaterial"`

	// Import token obtained by calling `GetParametersForImport`.
	ImportToken *string `json:"ImportToken,omitempty" name:"ImportToken"`

	// Specifies the CMK into which to import key material, which must be the same as the one specified by `GetParametersForImport`.
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Unix timestamp of the key material's expiration time. If this value is empty or 0, the key material will never expire. To specify the expiration time, it should be later than the current time. Maximum value: 2147443200.
	ValidTo *uint64 `json:"ValidTo,omitempty" name:"ValidTo"`
}

func (r *ImportKeyMaterialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportKeyMaterialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EncryptedKeyMaterial")
	delete(f, "ImportToken")
	delete(f, "KeyId")
	delete(f, "ValidTo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImportKeyMaterialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportKeyMaterialResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ImportKeyMaterialResponse struct {
	*tchttp.BaseResponse
	Response *ImportKeyMaterialResponseParams `json:"Response"`
}

func (r *ImportKeyMaterialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportKeyMaterialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Key struct {
	// Globally unique CMK ID.
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

type KeyMetadata struct {
	// Globally unique CMK ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Alias that makes a key more recognizable and understandable
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// Key creation time
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// CMK description
	Description *string `json:"Description,omitempty" name:"Description"`

	// CMK status. Valid values: Enabled, Disabled, PendingDelete, PendingImport, Archived.
	KeyState *string `json:"KeyState,omitempty" name:"KeyState"`

	// CMK purpose. Valid values: `ENCRYPT_DECRYPT`, `ASYMMETRIC_DECRYPT_RSA_2048`, `ASYMMETRIC_DECRYPT_SM2`, `ASYMMETRIC_SIGN_VERIFY_SM2`, `ASYMMETRIC_SIGN_VERIFY_RSA_2048`, and `ASYMMETRIC_SIGN_VERIFY_ECC`.
	KeyUsage *string `json:"KeyUsage,omitempty" name:"KeyUsage"`

	// CMK type. 2: FIPS-compliant; 4: SM-CRYPTO
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// Creator
	CreatorUin *uint64 `json:"CreatorUin,omitempty" name:"CreatorUin"`

	// Whether key rotation is enabled
	KeyRotationEnabled *bool `json:"KeyRotationEnabled,omitempty" name:"KeyRotationEnabled"`

	// CMK creator. The value of this parameter is `user` if the CMK is created by the user, or the corresponding service name if it is created automatically by an authorized Tencent Cloud service.
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// Time of next rotation if key rotation is enabled
	NextRotateTime *uint64 `json:"NextRotateTime,omitempty" name:"NextRotateTime"`

	// Scheduled deletion time
	DeletionDate *uint64 `json:"DeletionDate,omitempty" name:"DeletionDate"`

	// CMK key material type. TENCENT_KMS: created by KMS; EXTERNAL: imported by user.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Origin *string `json:"Origin,omitempty" name:"Origin"`

	// It's valid when `Origin` is `EXTERNAL`, indicating the expiration date of key material. 0 means valid forever.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ValidTo *uint64 `json:"ValidTo,omitempty" name:"ValidTo"`

	// Resource ID in the format of `creatorUin/$creatorUin/$keyId`.
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

// Predefined struct for user
type ListAlgorithmsRequestParams struct {

}

type ListAlgorithmsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ListAlgorithmsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAlgorithmsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAlgorithmsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAlgorithmsResponseParams struct {
	// Symmetric encryption algorithms supported in this region
	SymmetricAlgorithms []*AlgorithmInfo `json:"SymmetricAlgorithms,omitempty" name:"SymmetricAlgorithms"`

	// Asymmetric encryption algorithms supported in this region
	AsymmetricAlgorithms []*AlgorithmInfo `json:"AsymmetricAlgorithms,omitempty" name:"AsymmetricAlgorithms"`

	// Asymmetric signature verification algorithms supported in the current region
	AsymmetricSignVerifyAlgorithms []*AlgorithmInfo `json:"AsymmetricSignVerifyAlgorithms,omitempty" name:"AsymmetricSignVerifyAlgorithms"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListAlgorithmsResponse struct {
	*tchttp.BaseResponse
	Response *ListAlgorithmsResponseParams `json:"Response"`
}

func (r *ListAlgorithmsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAlgorithmsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListKeyDetailRequestParams struct {
	// This parameter has the same meaning of the `Offset` in an SQL query, indicating that this acquisition starts from the "No. Offset value" element of the array arranged in a certain order. The default value is 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// This parameter has the same meaning of the `Limit` in an SQL query, indicating that up to `Limit` value elements can be obtained in this request. The default value is 10 and the maximum value is 200.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Filters by creator role. 0 (default value): the CMK is created by the user; 1: the CMK is created automatically by an authorized Tencent Cloud service.
	Role *uint64 `json:"Role,omitempty" name:"Role"`

	// Sorts by CMK creation time. 0: descending; 1: ascending
	OrderType *uint64 `json:"OrderType,omitempty" name:"OrderType"`

	// Filters by CMK status. 0: all CMKs; 1: CMKs in `Enabled` status only; 2: CMKs in `Disabled` status only; 3: CMKs in `PendingDelete` status only (i.e., keys with schedule deletion enabled); 4: CMKs in `PendingImport` status only; 5: CMKs in `Archived` status only.
	KeyState *uint64 `json:"KeyState,omitempty" name:"KeyState"`

	// Performs a fuzzy query by `KeyId` or `Alias`
	SearchKeyAlias *string `json:"SearchKeyAlias,omitempty" name:"SearchKeyAlias"`

	// Filters by CMK type. "TENCENT_KMS" indicates to filter CMKs whose key materials are created by KMS; "EXTERNAL" indicates to filter CMKs of `EXTERNAL` type whose key materials are imported by users; "ALL" or empty indicates to filter CMKs of both types. This value is case-sensitive.
	Origin *string `json:"Origin,omitempty" name:"Origin"`

	// Filters by the `KeyUsage` field value. Valid values: `ALL` (all CMKs), `ENCRYPT_DECRYPT` (used when this field is left empty), `ASYMMETRIC_DECRYPT_RSA_2048`, `ASYMMETRIC_DECRYPT_SM2`, `ASYMMETRIC_SIGN_VERIFY_SM2`, `ASYMMETRIC_SIGN_VERIFY_RSA_2048`, and `ASYMMETRIC_SIGN_VERIFY_ECC`.
	KeyUsage *string `json:"KeyUsage,omitempty" name:"KeyUsage"`

	// Tag filter condition
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`
}

type ListKeyDetailRequest struct {
	*tchttp.BaseRequest
	
	// This parameter has the same meaning of the `Offset` in an SQL query, indicating that this acquisition starts from the "No. Offset value" element of the array arranged in a certain order. The default value is 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// This parameter has the same meaning of the `Limit` in an SQL query, indicating that up to `Limit` value elements can be obtained in this request. The default value is 10 and the maximum value is 200.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Filters by creator role. 0 (default value): the CMK is created by the user; 1: the CMK is created automatically by an authorized Tencent Cloud service.
	Role *uint64 `json:"Role,omitempty" name:"Role"`

	// Sorts by CMK creation time. 0: descending; 1: ascending
	OrderType *uint64 `json:"OrderType,omitempty" name:"OrderType"`

	// Filters by CMK status. 0: all CMKs; 1: CMKs in `Enabled` status only; 2: CMKs in `Disabled` status only; 3: CMKs in `PendingDelete` status only (i.e., keys with schedule deletion enabled); 4: CMKs in `PendingImport` status only; 5: CMKs in `Archived` status only.
	KeyState *uint64 `json:"KeyState,omitempty" name:"KeyState"`

	// Performs a fuzzy query by `KeyId` or `Alias`
	SearchKeyAlias *string `json:"SearchKeyAlias,omitempty" name:"SearchKeyAlias"`

	// Filters by CMK type. "TENCENT_KMS" indicates to filter CMKs whose key materials are created by KMS; "EXTERNAL" indicates to filter CMKs of `EXTERNAL` type whose key materials are imported by users; "ALL" or empty indicates to filter CMKs of both types. This value is case-sensitive.
	Origin *string `json:"Origin,omitempty" name:"Origin"`

	// Filters by the `KeyUsage` field value. Valid values: `ALL` (all CMKs), `ENCRYPT_DECRYPT` (used when this field is left empty), `ASYMMETRIC_DECRYPT_RSA_2048`, `ASYMMETRIC_DECRYPT_SM2`, `ASYMMETRIC_SIGN_VERIFY_SM2`, `ASYMMETRIC_SIGN_VERIFY_RSA_2048`, and `ASYMMETRIC_SIGN_VERIFY_ECC`.
	KeyUsage *string `json:"KeyUsage,omitempty" name:"KeyUsage"`

	// Tag filter condition
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`
}

func (r *ListKeyDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListKeyDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Role")
	delete(f, "OrderType")
	delete(f, "KeyState")
	delete(f, "SearchKeyAlias")
	delete(f, "Origin")
	delete(f, "KeyUsage")
	delete(f, "TagFilters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListKeyDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListKeyDetailResponseParams struct {
	// Total number of CMKs
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of returned attribute information.
	// Note: this field may return null, indicating that no valid values can be obtained.
	KeyMetadatas []*KeyMetadata `json:"KeyMetadatas,omitempty" name:"KeyMetadatas"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListKeyDetailResponse struct {
	*tchttp.BaseResponse
	Response *ListKeyDetailResponseParams `json:"Response"`
}

func (r *ListKeyDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListKeyDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListKeysRequestParams struct {
	// This parameter has the same meaning of the `Offset` in an SQL query, indicating that this acquisition starts from the "No. Offset value" element of the array arranged in a certain order. The default value is 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// This parameter has the same meaning of the `Limit` in an SQL query, indicating that up to `Limit` value elements can be obtained in this request. The default value is 10 and the maximum value is 200
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Filter by creator role. 0 (default value): the CMK is created by the user; 1: the CMK is created automatically by an authorized Tencent Cloud service
	Role *uint64 `json:"Role,omitempty" name:"Role"`
}

type ListKeysRequest struct {
	*tchttp.BaseRequest
	
	// This parameter has the same meaning of the `Offset` in an SQL query, indicating that this acquisition starts from the "No. Offset value" element of the array arranged in a certain order. The default value is 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// This parameter has the same meaning of the `Limit` in an SQL query, indicating that up to `Limit` value elements can be obtained in this request. The default value is 10 and the maximum value is 200
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Filter by creator role. 0 (default value): the CMK is created by the user; 1: the CMK is created automatically by an authorized Tencent Cloud service
	Role *uint64 `json:"Role,omitempty" name:"Role"`
}

func (r *ListKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Role")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListKeysResponseParams struct {
	// CMK list array
	Keys []*Key `json:"Keys,omitempty" name:"Keys"`

	// Total number of CMKs
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListKeysResponse struct {
	*tchttp.BaseResponse
	Response *ListKeysResponseParams `json:"Response"`
}

func (r *ListKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OverwriteWhiteBoxDeviceFingerprintsRequestParams struct {
	// White-box key ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Device fingerprint list. If the list is empty, it means to delete all fingerprint information corresponding to the key. There can be up to 200 entries in the list.
	DeviceFingerprints []*DeviceFingerprint `json:"DeviceFingerprints,omitempty" name:"DeviceFingerprints"`
}

type OverwriteWhiteBoxDeviceFingerprintsRequest struct {
	*tchttp.BaseRequest
	
	// White-box key ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Device fingerprint list. If the list is empty, it means to delete all fingerprint information corresponding to the key. There can be up to 200 entries in the list.
	DeviceFingerprints []*DeviceFingerprint `json:"DeviceFingerprints,omitempty" name:"DeviceFingerprints"`
}

func (r *OverwriteWhiteBoxDeviceFingerprintsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OverwriteWhiteBoxDeviceFingerprintsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	delete(f, "DeviceFingerprints")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OverwriteWhiteBoxDeviceFingerprintsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OverwriteWhiteBoxDeviceFingerprintsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type OverwriteWhiteBoxDeviceFingerprintsResponse struct {
	*tchttp.BaseResponse
	Response *OverwriteWhiteBoxDeviceFingerprintsResponseParams `json:"Response"`
}

func (r *OverwriteWhiteBoxDeviceFingerprintsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OverwriteWhiteBoxDeviceFingerprintsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReEncryptRequestParams struct {
	// Ciphertext to be re-encrypted
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" name:"CiphertextBlob"`

	// CMK used for re-encryption. If this parameter is empty, the ciphertext will be re-encrypted by using the original CMK (as long as the key is not rotated, the ciphertext will not be refreshed)
	DestinationKeyId *string `json:"DestinationKeyId,omitempty" name:"DestinationKeyId"`

	// JSON string of the key-value pair used during ciphertext encryption by `CiphertextBlob`. If not used during encryption, this parameter will be empty
	SourceEncryptionContext *string `json:"SourceEncryptionContext,omitempty" name:"SourceEncryptionContext"`

	// JSON string of the key-value pair used during re-encryption. If this field is used, the same string should be entered when the returned new ciphertext is decrypted
	DestinationEncryptionContext *string `json:"DestinationEncryptionContext,omitempty" name:"DestinationEncryptionContext"`
}

type ReEncryptRequest struct {
	*tchttp.BaseRequest
	
	// Ciphertext to be re-encrypted
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" name:"CiphertextBlob"`

	// CMK used for re-encryption. If this parameter is empty, the ciphertext will be re-encrypted by using the original CMK (as long as the key is not rotated, the ciphertext will not be refreshed)
	DestinationKeyId *string `json:"DestinationKeyId,omitempty" name:"DestinationKeyId"`

	// JSON string of the key-value pair used during ciphertext encryption by `CiphertextBlob`. If not used during encryption, this parameter will be empty
	SourceEncryptionContext *string `json:"SourceEncryptionContext,omitempty" name:"SourceEncryptionContext"`

	// JSON string of the key-value pair used during re-encryption. If this field is used, the same string should be entered when the returned new ciphertext is decrypted
	DestinationEncryptionContext *string `json:"DestinationEncryptionContext,omitempty" name:"DestinationEncryptionContext"`
}

func (r *ReEncryptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReEncryptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CiphertextBlob")
	delete(f, "DestinationKeyId")
	delete(f, "SourceEncryptionContext")
	delete(f, "DestinationEncryptionContext")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReEncryptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReEncryptResponseParams struct {
	// Re-encrypted ciphertext
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" name:"CiphertextBlob"`

	// CMK used for re-encryption
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// CMK used by ciphertext before re-encryption
	SourceKeyId *string `json:"SourceKeyId,omitempty" name:"SourceKeyId"`

	// `true` indicates that the ciphertext has been re-encrypted. When re-encryption is initiated by using the same CMK, as long as the CMK is not rotated, no actual re-encryption will be performed, and the original ciphertext will be returned
	ReEncrypted *bool `json:"ReEncrypted,omitempty" name:"ReEncrypted"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReEncryptResponse struct {
	*tchttp.BaseResponse
	Response *ReEncryptResponseParams `json:"Response"`
}

func (r *ReEncryptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReEncryptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScheduleKeyDeletionRequestParams struct {
	// Unique CMK ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Schedule deletion time range. Value range: [7,30]
	PendingWindowInDays *uint64 `json:"PendingWindowInDays,omitempty" name:"PendingWindowInDays"`
}

type ScheduleKeyDeletionRequest struct {
	*tchttp.BaseRequest
	
	// Unique CMK ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Schedule deletion time range. Value range: [7,30]
	PendingWindowInDays *uint64 `json:"PendingWindowInDays,omitempty" name:"PendingWindowInDays"`
}

func (r *ScheduleKeyDeletionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScheduleKeyDeletionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	delete(f, "PendingWindowInDays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScheduleKeyDeletionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScheduleKeyDeletionResponseParams struct {
	// Schedule deletion execution time
	DeletionDate *uint64 `json:"DeletionDate,omitempty" name:"DeletionDate"`

	// Unique ID of the CMK scheduled for deletion
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ScheduleKeyDeletionResponse struct {
	*tchttp.BaseResponse
	Response *ScheduleKeyDeletionResponseParams `json:"Response"`
}

func (r *ScheduleKeyDeletionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScheduleKeyDeletionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SignByAsymmetricKeyRequestParams struct {
	// Signature algorithm. The valid values include `SM2DSA`, `ECC_P256_R1`, `RSA_PSS_SHA_256`, and `RSA_PKCS1_SHA_256`, etc. You can get a full list of supported algorithms using the ListAlgorithms API.
	Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`

	// Full message or message abstract. Before Base64 encoding, an original message can contain up to 4,096 bytes while a message abstract must be 32 bytes.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Unique ID of a key
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Message type. Valid values: `RAW` (indicating an original message; used by default if the parameter is not passed in) and `DIGEST`.
	MessageType *string `json:"MessageType,omitempty" name:"MessageType"`
}

type SignByAsymmetricKeyRequest struct {
	*tchttp.BaseRequest
	
	// Signature algorithm. The valid values include `SM2DSA`, `ECC_P256_R1`, `RSA_PSS_SHA_256`, and `RSA_PKCS1_SHA_256`, etc. You can get a full list of supported algorithms using the ListAlgorithms API.
	Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`

	// Full message or message abstract. Before Base64 encoding, an original message can contain up to 4,096 bytes while a message abstract must be 32 bytes.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Unique ID of a key
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Message type. Valid values: `RAW` (indicating an original message; used by default if the parameter is not passed in) and `DIGEST`.
	MessageType *string `json:"MessageType,omitempty" name:"MessageType"`
}

func (r *SignByAsymmetricKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SignByAsymmetricKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Algorithm")
	delete(f, "Message")
	delete(f, "KeyId")
	delete(f, "MessageType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SignByAsymmetricKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SignByAsymmetricKeyResponseParams struct {
	// Base64-encoded signature
	Signature *string `json:"Signature,omitempty" name:"Signature"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SignByAsymmetricKeyResponse struct {
	*tchttp.BaseResponse
	Response *SignByAsymmetricKeyResponseParams `json:"Response"`
}

func (r *SignByAsymmetricKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SignByAsymmetricKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {
	// Tag key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TagFilter struct {
	// Tag key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	TagValue []*string `json:"TagValue,omitempty" name:"TagValue"`
}

// Predefined struct for user
type UnbindCloudResourceRequestParams struct {
	// CMK ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Unique ID of a Tencent Cloud service
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// Resource/instance ID, which is stored as a string and defined by the caller based on the Tencent Cloud service’s features.
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

type UnbindCloudResourceRequest struct {
	*tchttp.BaseRequest
	
	// CMK ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Unique ID of a Tencent Cloud service
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// Resource/instance ID, which is stored as a string and defined by the caller based on the Tencent Cloud service’s features.
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

func (r *UnbindCloudResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindCloudResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	delete(f, "ProductId")
	delete(f, "ResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindCloudResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindCloudResourceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnbindCloudResourceResponse struct {
	*tchttp.BaseResponse
	Response *UnbindCloudResourceResponseParams `json:"Response"`
}

func (r *UnbindCloudResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindCloudResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAliasRequestParams struct {
	// New alias containing 1-60 characters or digits
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// Globally unique CMK ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

type UpdateAliasRequest struct {
	*tchttp.BaseRequest
	
	// New alias containing 1-60 characters or digits
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// Globally unique CMK ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *UpdateAliasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAliasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Alias")
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAliasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAliasResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateAliasResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAliasResponseParams `json:"Response"`
}

func (r *UpdateAliasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAliasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateKeyDescriptionRequestParams struct {
	// New description of up to 1,024 bytes in length
	Description *string `json:"Description,omitempty" name:"Description"`

	// ID of the CMK for which to modify the description
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

type UpdateKeyDescriptionRequest struct {
	*tchttp.BaseRequest
	
	// New description of up to 1,024 bytes in length
	Description *string `json:"Description,omitempty" name:"Description"`

	// ID of the CMK for which to modify the description
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *UpdateKeyDescriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateKeyDescriptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Description")
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateKeyDescriptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateKeyDescriptionResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateKeyDescriptionResponse struct {
	*tchttp.BaseResponse
	Response *UpdateKeyDescriptionResponseParams `json:"Response"`
}

func (r *UpdateKeyDescriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateKeyDescriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyByAsymmetricKeyRequestParams struct {
	// Unique ID of a key
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Signature value, which is generated by calling the KMS signature API.
	SignatureValue *string `json:"SignatureValue,omitempty" name:"SignatureValue"`

	// Full message or message abstract. Before Base64 encoding, an original message can contain up to 4,096 bytes while a message abstract must be 32 bytes.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Signature algorithm. The valid values include `SM2DSA`, `ECC_P256_R1`, `RSA_PSS_SHA_256`, and `RSA_PKCS1_SHA_256`, etc. You can get a full list of supported algorithms using the ListAlgorithms API.
	Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`

	// Message type. Valid values: `RAW` (indicating an original message; used by default if the parameter is not passed in) and `DIGEST`.
	MessageType *string `json:"MessageType,omitempty" name:"MessageType"`
}

type VerifyByAsymmetricKeyRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of a key
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Signature value, which is generated by calling the KMS signature API.
	SignatureValue *string `json:"SignatureValue,omitempty" name:"SignatureValue"`

	// Full message or message abstract. Before Base64 encoding, an original message can contain up to 4,096 bytes while a message abstract must be 32 bytes.
	Message *string `json:"Message,omitempty" name:"Message"`

	// Signature algorithm. The valid values include `SM2DSA`, `ECC_P256_R1`, `RSA_PSS_SHA_256`, and `RSA_PKCS1_SHA_256`, etc. You can get a full list of supported algorithms using the ListAlgorithms API.
	Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`

	// Message type. Valid values: `RAW` (indicating an original message; used by default if the parameter is not passed in) and `DIGEST`.
	MessageType *string `json:"MessageType,omitempty" name:"MessageType"`
}

func (r *VerifyByAsymmetricKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyByAsymmetricKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyId")
	delete(f, "SignatureValue")
	delete(f, "Message")
	delete(f, "Algorithm")
	delete(f, "MessageType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyByAsymmetricKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyByAsymmetricKeyResponseParams struct {
	// Whether the signature is valid. `true`: the signature is valid; `false`: the signature is invalid.
	SignatureValid *bool `json:"SignatureValid,omitempty" name:"SignatureValid"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type VerifyByAsymmetricKeyResponse struct {
	*tchttp.BaseResponse
	Response *VerifyByAsymmetricKeyResponseParams `json:"Response"`
}

func (r *VerifyByAsymmetricKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyByAsymmetricKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WhiteboxKeyInfo struct {
	// Globally unique white-box key ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// Unique alias that makes a key more recognizable and understandable. This parameter cannot be empty, can contain 1 to 60 letters, digits, hyphens (-), and underscores (_), and must begin with a letter or digit.
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// Creator
	CreatorUin *uint64 `json:"CreatorUin,omitempty" name:"CreatorUin"`

	// Key description information
	Description *string `json:"Description,omitempty" name:"Description"`

	// Key creation time in Unix timestamp
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// White-box key status. Valid values: Enabled, Disabled
	Status *string `json:"Status,omitempty" name:"Status"`

	// Creator
	OwnerUin *uint64 `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// Key algorithm type
	Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`

	// Base64-encoded white-box encryption key
	EncryptKey *string `json:"EncryptKey,omitempty" name:"EncryptKey"`

	// Base64-encoded white-box decryption key
	DecryptKey *string `json:"DecryptKey,omitempty" name:"DecryptKey"`

	// Resource ID in the format of `creatorUin/$creatorUin/$keyId`
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Whether there is a device fingerprint bound to the current key
	// Note: this field may return null, indicating that no valid values can be obtained.
	DeviceFingerprintBind *bool `json:"DeviceFingerprintBind,omitempty" name:"DeviceFingerprintBind"`
}