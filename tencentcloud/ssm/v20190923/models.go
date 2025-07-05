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

package v20190923

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

// Predefined struct for user
type CreateProductSecretRequestParams struct {
	// Credential name, which must be unique in the same region. It can contain 128 bytes of letters, digits, hyphens, and underscores and must begin with a letter or digit.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// Prefix of the user account name, which is specified by you and can contain up to 8 characters.
	// Supported character sets include:
	// Digits: [0, 9].
	// Lowercase letters: [a, z].
	// Uppercase letters: [A, Z].
	// Special symbols: underscore.
	// The prefix must begin with a letter.
	UserNamePrefix *string `json:"UserNamePrefix,omitnil,omitempty" name:"UserNamePrefix"`

	// Name of the Tencent Cloud service bound to the credential, such as `Mysql`. The `DescribeSupportedProducts` API can be used to get the names of the supported Tencent Cloud services.
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// Tencent Cloud service instance ID.
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// Domain name of the account in the form of IP. You can enter `%`.
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// List of permissions that need to be granted when the credential is bound to a Tencent Cloud service.
	PrivilegesList []*ProductPrivilegeUnit `json:"PrivilegesList,omitnil,omitempty" name:"PrivilegesList"`

	// Description, which is used to describe the purpose in detail and can contain up to 2,048 bytes.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Specifies the KMS CMK that encrypts the credential.
	// If this parameter is left empty, the CMK created by Secrets Manager by default will be used for encryption.
	// You can also specify a custom KMS CMK created in the same region for encryption.
	KmsKeyId *string `json:"KmsKeyId,omitnil,omitempty" name:"KmsKeyId"`

	// List of tags.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// User-Defined rotation start time in the format of 2006-01-02 15:04:05.
	// When `EnableRotation` is `True`, this parameter is required.
	RotationBeginTime *string `json:"RotationBeginTime,omitnil,omitempty" name:"RotationBeginTime"`

	// Specifies whether to enable rotation
	// True - enable
	// False - do not enable
	// If this parameter is not specified, `False` will be used by default.
	EnableRotation *bool `json:"EnableRotation,omitnil,omitempty" name:"EnableRotation"`

	// Rotation frequency in days. Default value: 1 day.
	RotationFrequency *int64 `json:"RotationFrequency,omitnil,omitempty" name:"RotationFrequency"`
}

type CreateProductSecretRequest struct {
	*tchttp.BaseRequest
	
	// Credential name, which must be unique in the same region. It can contain 128 bytes of letters, digits, hyphens, and underscores and must begin with a letter or digit.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// Prefix of the user account name, which is specified by you and can contain up to 8 characters.
	// Supported character sets include:
	// Digits: [0, 9].
	// Lowercase letters: [a, z].
	// Uppercase letters: [A, Z].
	// Special symbols: underscore.
	// The prefix must begin with a letter.
	UserNamePrefix *string `json:"UserNamePrefix,omitnil,omitempty" name:"UserNamePrefix"`

	// Name of the Tencent Cloud service bound to the credential, such as `Mysql`. The `DescribeSupportedProducts` API can be used to get the names of the supported Tencent Cloud services.
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// Tencent Cloud service instance ID.
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// Domain name of the account in the form of IP. You can enter `%`.
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// List of permissions that need to be granted when the credential is bound to a Tencent Cloud service.
	PrivilegesList []*ProductPrivilegeUnit `json:"PrivilegesList,omitnil,omitempty" name:"PrivilegesList"`

	// Description, which is used to describe the purpose in detail and can contain up to 2,048 bytes.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Specifies the KMS CMK that encrypts the credential.
	// If this parameter is left empty, the CMK created by Secrets Manager by default will be used for encryption.
	// You can also specify a custom KMS CMK created in the same region for encryption.
	KmsKeyId *string `json:"KmsKeyId,omitnil,omitempty" name:"KmsKeyId"`

	// List of tags.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// User-Defined rotation start time in the format of 2006-01-02 15:04:05.
	// When `EnableRotation` is `True`, this parameter is required.
	RotationBeginTime *string `json:"RotationBeginTime,omitnil,omitempty" name:"RotationBeginTime"`

	// Specifies whether to enable rotation
	// True - enable
	// False - do not enable
	// If this parameter is not specified, `False` will be used by default.
	EnableRotation *bool `json:"EnableRotation,omitnil,omitempty" name:"EnableRotation"`

	// Rotation frequency in days. Default value: 1 day.
	RotationFrequency *int64 `json:"RotationFrequency,omitnil,omitempty" name:"RotationFrequency"`
}

func (r *CreateProductSecretRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProductSecretRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	delete(f, "UserNamePrefix")
	delete(f, "ProductName")
	delete(f, "InstanceID")
	delete(f, "Domains")
	delete(f, "PrivilegesList")
	delete(f, "Description")
	delete(f, "KmsKeyId")
	delete(f, "Tags")
	delete(f, "RotationBeginTime")
	delete(f, "EnableRotation")
	delete(f, "RotationFrequency")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProductSecretRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProductSecretResponseParams struct {
	// Name of the created credential.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// Tag operation return code. 0: success; 1: internal error; 2: business processing error.
	// Note: this field may return null, indicating that no valid values can be obtained.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TagCode *uint64 `json:"TagCode,omitnil,omitempty" name:"TagCode"`

	// Tag operation return message.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TagMsg *string `json:"TagMsg,omitnil,omitempty" name:"TagMsg"`

	// ID of the created Tencent Cloud service credential async task.
	FlowID *int64 `json:"FlowID,omitnil,omitempty" name:"FlowID"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateProductSecretResponse struct {
	*tchttp.BaseResponse
	Response *CreateProductSecretResponseParams `json:"Response"`
}

func (r *CreateProductSecretResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProductSecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSSHKeyPairSecretRequestParams struct {
	// Secret name, which must be unique in the same region. It can contain 128 bytes of letters, digits, hyphens and underscores and must begin with a letter or digit.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// ID of the project to which the created SSH key belongs.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Description, such as what it is used for. It contains up to 2,048 bytes.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Specifies a KMS CMK to encrypt the secret.
	// If this parameter is left empty, the CMK created by Secrets Manager by default will be used for encryption.
	// You can also specify a custom KMS CMK created in the same region for encryption.
	KmsKeyId *string `json:"KmsKeyId,omitnil,omitempty" name:"KmsKeyId"`

	// List of tags.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Name of the SSH key pair, which only contains digits, letters and underscores and must start with a digit or letter. The maximum length is 25 characters.
	SSHKeyName *string `json:"SSHKeyName,omitnil,omitempty" name:"SSHKeyName"`
}

type CreateSSHKeyPairSecretRequest struct {
	*tchttp.BaseRequest
	
	// Secret name, which must be unique in the same region. It can contain 128 bytes of letters, digits, hyphens and underscores and must begin with a letter or digit.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// ID of the project to which the created SSH key belongs.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Description, such as what it is used for. It contains up to 2,048 bytes.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Specifies a KMS CMK to encrypt the secret.
	// If this parameter is left empty, the CMK created by Secrets Manager by default will be used for encryption.
	// You can also specify a custom KMS CMK created in the same region for encryption.
	KmsKeyId *string `json:"KmsKeyId,omitnil,omitempty" name:"KmsKeyId"`

	// List of tags.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Name of the SSH key pair, which only contains digits, letters and underscores and must start with a digit or letter. The maximum length is 25 characters.
	SSHKeyName *string `json:"SSHKeyName,omitnil,omitempty" name:"SSHKeyName"`
}

func (r *CreateSSHKeyPairSecretRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSSHKeyPairSecretRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	delete(f, "ProjectId")
	delete(f, "Description")
	delete(f, "KmsKeyId")
	delete(f, "Tags")
	delete(f, "SSHKeyName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSSHKeyPairSecretRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSSHKeyPairSecretResponseParams struct {
	// Name of the created secret.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// ID of the created SSH key.
	SSHKeyID *string `json:"SSHKeyID,omitnil,omitempty" name:"SSHKeyID"`

	// Name of the created SSH key.
	SSHKeyName *string `json:"SSHKeyName,omitnil,omitempty" name:"SSHKeyName"`

	// Tag return code. `0`: success; `1`: internal error; `2`: business processing error.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TagCode *uint64 `json:"TagCode,omitnil,omitempty" name:"TagCode"`

	// Tag return message.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TagMsg *string `json:"TagMsg,omitnil,omitempty" name:"TagMsg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSSHKeyPairSecretResponse struct {
	*tchttp.BaseResponse
	Response *CreateSSHKeyPairSecretResponseParams `json:"Response"`
}

func (r *CreateSSHKeyPairSecretResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSSHKeyPairSecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecretRequestParams struct {
	// Secret name, which must be unique in the same region. It can contain 128 bytes ([a-z], [A-Z], [0-9], [-_]). It must begin with a letter or digit. Note that it cannot be modified once created. 
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// Secret version. It can contain up to 64 bytes ([a-z], [A-Z], [0-9], [-_.]). It must begin with a letter or digit. `SecretName` and `VersionId` are used to query the Secret information. If it is left empty, the initial Secret version number is used by default.
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// Description information, such as the detailed use cases. It can be up to 2048 bytes.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// KMS CMK used for Secret encryption. If this parameter is left empty, SecretsManager will create a CMK by default. You can also specify a KMS CMK that is created in the same region.
	KmsKeyId *string `json:"KmsKeyId,omitnil,omitempty" name:"KmsKeyId"`

	// Secret type. It defaults to `custom`.
	SecretType *uint64 `json:"SecretType,omitnil,omitempty" name:"SecretType"`

	// Base64-encoded plaintext of a binary Secret. Either `SecretBinary` or `SecretString` must be set. A maximum of 4096 bytes is supported.
	SecretBinary *string `json:"SecretBinary,omitnil,omitempty" name:"SecretBinary"`

	// Plaintext of a Secret, in text format. Base64 encoding is not required. Either `SecretBinary` or `SecretString` must be set. A maximum of 4096 bytes is supported.
	SecretString *string `json:"SecretString,omitnil,omitempty" name:"SecretString"`

	// Additional configuration of the Secret in JSON format
	AdditionalConfig *string `json:"AdditionalConfig,omitnil,omitempty" name:"AdditionalConfig"`

	// List of tags.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateSecretRequest struct {
	*tchttp.BaseRequest
	
	// Secret name, which must be unique in the same region. It can contain 128 bytes ([a-z], [A-Z], [0-9], [-_]). It must begin with a letter or digit. Note that it cannot be modified once created. 
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// Secret version. It can contain up to 64 bytes ([a-z], [A-Z], [0-9], [-_.]). It must begin with a letter or digit. `SecretName` and `VersionId` are used to query the Secret information. If it is left empty, the initial Secret version number is used by default.
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// Description information, such as the detailed use cases. It can be up to 2048 bytes.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// KMS CMK used for Secret encryption. If this parameter is left empty, SecretsManager will create a CMK by default. You can also specify a KMS CMK that is created in the same region.
	KmsKeyId *string `json:"KmsKeyId,omitnil,omitempty" name:"KmsKeyId"`

	// Secret type. It defaults to `custom`.
	SecretType *uint64 `json:"SecretType,omitnil,omitempty" name:"SecretType"`

	// Base64-encoded plaintext of a binary Secret. Either `SecretBinary` or `SecretString` must be set. A maximum of 4096 bytes is supported.
	SecretBinary *string `json:"SecretBinary,omitnil,omitempty" name:"SecretBinary"`

	// Plaintext of a Secret, in text format. Base64 encoding is not required. Either `SecretBinary` or `SecretString` must be set. A maximum of 4096 bytes is supported.
	SecretString *string `json:"SecretString,omitnil,omitempty" name:"SecretString"`

	// Additional configuration of the Secret in JSON format
	AdditionalConfig *string `json:"AdditionalConfig,omitnil,omitempty" name:"AdditionalConfig"`

	// List of tags.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateSecretRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecretRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	delete(f, "VersionId")
	delete(f, "Description")
	delete(f, "KmsKeyId")
	delete(f, "SecretType")
	delete(f, "SecretBinary")
	delete(f, "SecretString")
	delete(f, "AdditionalConfig")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSecretRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecretResponseParams struct {
	// Name of the new Secret.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// ID of the new Secret version.
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// Return code of tag operation. `0`: success; `1`: internal error; `2`: business processing error
	// Note: This field may return `null`, indicating that no valid value was found.
	TagCode *uint64 `json:"TagCode,omitnil,omitempty" name:"TagCode"`

	// Return message of tag operation.
	// Note: This field may return `null`, indicating that no valid value was found.
	TagMsg *string `json:"TagMsg,omitnil,omitempty" name:"TagMsg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSecretResponse struct {
	*tchttp.BaseResponse
	Response *CreateSecretResponseParams `json:"Response"`
}

func (r *CreateSecretResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecretRequestParams struct {
	// Name of the Secret to be deleted.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// Scheduled deletion time (in days), indicating the number of retention days for the secret. Value range: 0-30. If it is `0`, the secret is deleted immediately.
	// For an SSH key secret, this field can only be `0`.
	RecoveryWindowInDays *uint64 `json:"RecoveryWindowInDays,omitnil,omitempty" name:"RecoveryWindowInDays"`

	// Specifies whether to delete the SSH key from both the secret and the SSH key list in the CVM console. This field is only valid for SSH key secrets. Valid values:
	// `True`: deletes SSH key from both the secret and SSH key list in the CVM console. Note that the deletion will fail if the SSH key is already bound to a CVM instance.
	// `False`: only deletes the SSH key information in the secret.
	CleanSSHKey *bool `json:"CleanSSHKey,omitnil,omitempty" name:"CleanSSHKey"`
}

type DeleteSecretRequest struct {
	*tchttp.BaseRequest
	
	// Name of the Secret to be deleted.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// Scheduled deletion time (in days), indicating the number of retention days for the secret. Value range: 0-30. If it is `0`, the secret is deleted immediately.
	// For an SSH key secret, this field can only be `0`.
	RecoveryWindowInDays *uint64 `json:"RecoveryWindowInDays,omitnil,omitempty" name:"RecoveryWindowInDays"`

	// Specifies whether to delete the SSH key from both the secret and the SSH key list in the CVM console. This field is only valid for SSH key secrets. Valid values:
	// `True`: deletes SSH key from both the secret and SSH key list in the CVM console. Note that the deletion will fail if the SSH key is already bound to a CVM instance.
	// `False`: only deletes the SSH key information in the secret.
	CleanSSHKey *bool `json:"CleanSSHKey,omitnil,omitempty" name:"CleanSSHKey"`
}

func (r *DeleteSecretRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecretRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	delete(f, "RecoveryWindowInDays")
	delete(f, "CleanSSHKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSecretRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecretResponseParams struct {
	// Name of deleted Secret.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// Secret deletion time, formatted as a Unix timestamp.
	DeleteTime *int64 `json:"DeleteTime,omitnil,omitempty" name:"DeleteTime"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSecretResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSecretResponseParams `json:"Response"`
}

func (r *DeleteSecretResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecretVersionRequestParams struct {
	// Name of the Secret.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// ID of the Secret version to be deleted.
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

type DeleteSecretVersionRequest struct {
	*tchttp.BaseRequest
	
	// Name of the Secret.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// ID of the Secret version to be deleted.
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

func (r *DeleteSecretVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecretVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	delete(f, "VersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSecretVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecretVersionResponseParams struct {
	// Name of the Secret.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// Version ID of the Secret.
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSecretVersionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSecretVersionResponseParams `json:"Response"`
}

func (r *DeleteSecretVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecretVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAsyncRequestInfoRequestParams struct {
	// Async task ID.
	FlowID *int64 `json:"FlowID,omitnil,omitempty" name:"FlowID"`
}

type DescribeAsyncRequestInfoRequest struct {
	*tchttp.BaseRequest
	
	// Async task ID.
	FlowID *int64 `json:"FlowID,omitnil,omitempty" name:"FlowID"`
}

func (r *DescribeAsyncRequestInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAsyncRequestInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAsyncRequestInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAsyncRequestInfoResponseParams struct {
	// 0: processing, 1: processing succeeded, 2: processing failed
	TaskStatus *int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// Task description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAsyncRequestInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAsyncRequestInfoResponseParams `json:"Response"`
}

func (r *DescribeAsyncRequestInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAsyncRequestInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRotationDetailRequestParams struct {
	// Specifies the name of the credential for which to get the credential rotation details.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`
}

type DescribeRotationDetailRequest struct {
	*tchttp.BaseRequest
	
	// Specifies the name of the credential for which to get the credential rotation details.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`
}

func (r *DescribeRotationDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRotationDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRotationDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRotationDetailResponseParams struct {
	// Whether to enable rotation. `true`: enabled; `false`: disabled.
	EnableRotation *bool `json:"EnableRotation,omitnil,omitempty" name:"EnableRotation"`

	// Rotation frequency in days. Default value: 1 day.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Frequency *int64 `json:"Frequency,omitnil,omitempty" name:"Frequency"`

	// Last rotation time, which is an explicitly visible time string in the format of 2006-01-02 15:04:05.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LatestRotateTime *string `json:"LatestRotateTime,omitnil,omitempty" name:"LatestRotateTime"`

	// Next rotation start time, which is an explicitly visible time string in the format of 2006-01-02 15:04:05.
	// Note: this field may return null, indicating that no valid values can be obtained.
	NextRotateBeginTime *string `json:"NextRotateBeginTime,omitnil,omitempty" name:"NextRotateBeginTime"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRotationDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRotationDetailResponseParams `json:"Response"`
}

func (r *DescribeRotationDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRotationDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRotationHistoryRequestParams struct {
	// Specifies the name of the credential for which to get the credential rotation records.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`
}

type DescribeRotationHistoryRequest struct {
	*tchttp.BaseRequest
	
	// Specifies the name of the credential for which to get the credential rotation records.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`
}

func (r *DescribeRotationHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRotationHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRotationHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRotationHistoryResponseParams struct {
	// List of version numbers.
	VersionIDs []*string `json:"VersionIDs,omitnil,omitempty" name:"VersionIDs"`

	// Number of version numbers. The maximum number of version numbers that can be shown to users is 10.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRotationHistoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRotationHistoryResponseParams `json:"Response"`
}

func (r *DescribeRotationHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRotationHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecretRequestParams struct {
	// Name of a Secret whose detailed information is to be obtained.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`
}

type DescribeSecretRequest struct {
	*tchttp.BaseRequest
	
	// Name of a Secret whose detailed information is to be obtained.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`
}

func (r *DescribeSecretRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecretRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecretRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecretResponseParams struct {
	// Name of the Secret.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// Description of the Secret.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// ID of the KMS CMK used for encryption.
	KmsKeyId *string `json:"KmsKeyId,omitnil,omitempty" name:"KmsKeyId"`

	// Creator UIN.
	CreateUin *uint64 `json:"CreateUin,omitnil,omitempty" name:"CreateUin"`

	// Credential status: Enabled, Disabled, PendingDelete, Creating, Failed.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Deletion time, formatted as a Unix timestamp. For a Secret that is not in `PendingDelete` status, this value is 0.
	DeleteTime *uint64 `json:"DeleteTime,omitnil,omitempty" name:"DeleteTime"`

	// Creation time.
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// `0`: user-defined secret; `1`: database credential; `2`: SSH key secret.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SecretType *int64 `json:"SecretType,omitnil,omitempty" name:"SecretType"`

	// Tencent Cloud service name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// Tencent Cloud service instance ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ResourceID *string `json:"ResourceID,omitnil,omitempty" name:"ResourceID"`

	// Whether to enable rotation. `True`: enable rotation; `False`: disable rotation.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	RotationStatus *bool `json:"RotationStatus,omitnil,omitempty" name:"RotationStatus"`

	// Rotation frequency in days by default.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RotationFrequency *int64 `json:"RotationFrequency,omitnil,omitempty" name:"RotationFrequency"`

	// Secret name. This field is only valid when the `SecretType` is set to `2` (SSH key secret).
	// Note: this field may return null, indicating that no valid values can be obtained.
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// Project ID. This field is only valid when the `SecretType` is set to `2` (SSH key secret).
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// ID of the CVM instance associated with the SSH key. ID. This field is only valid when the `SecretType` is set to `2` (SSH key secret).
	// Note: this field may return null, indicating that no valid values can be obtained.
	AssociatedInstanceIDs []*string `json:"AssociatedInstanceIDs,omitnil,omitempty" name:"AssociatedInstanceIDs"`

	// UIN of the Tencent Cloud API key. This field is valid when the secret type is Tencent Cloud API key secret.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TargetUin *uint64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// Additional configuration of the Secret
	// Note: This field may return null, indicating that no valid values can be obtained.
	AdditionalConfig *string `json:"AdditionalConfig,omitnil,omitempty" name:"AdditionalConfig"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSecretResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecretResponseParams `json:"Response"`
}

func (r *DescribeSecretResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSupportedProductsRequestParams struct {

}

type DescribeSupportedProductsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeSupportedProductsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSupportedProductsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSupportedProductsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSupportedProductsResponseParams struct {
	// List of supported services.
	Products []*string `json:"Products,omitnil,omitempty" name:"Products"`

	// Number of supported services
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSupportedProductsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSupportedProductsResponseParams `json:"Response"`
}

func (r *DescribeSupportedProductsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSupportedProductsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableSecretRequestParams struct {
	// Name of the Secret to be disabled.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`
}

type DisableSecretRequest struct {
	*tchttp.BaseRequest
	
	// Name of the Secret to be disabled.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`
}

func (r *DisableSecretRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableSecretRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableSecretRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableSecretResponseParams struct {
	// Name of the disabled Secret.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisableSecretResponse struct {
	*tchttp.BaseResponse
	Response *DisableSecretResponseParams `json:"Response"`
}

func (r *DisableSecretResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableSecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableSecretRequestParams struct {
	// Name of the Secret to be enabled.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`
}

type EnableSecretRequest struct {
	*tchttp.BaseRequest
	
	// Name of the Secret to be enabled.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`
}

func (r *EnableSecretRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableSecretRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableSecretRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableSecretResponseParams struct {
	// Name of the enabled Secret.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnableSecretResponse struct {
	*tchttp.BaseResponse
	Response *EnableSecretResponseParams `json:"Response"`
}

func (r *EnableSecretResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableSecretResponse) FromJsonString(s string) error {
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
	// List of regions.
	Regions []*string `json:"Regions,omitnil,omitempty" name:"Regions"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type GetSSHKeyPairValueRequestParams struct {
	// Secret name. This field is only valid for SSH key secrets.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// ID of the key pair, which is the unique identifier of the key pair in the CVM.
	SSHKeyId *string `json:"SSHKeyId,omitnil,omitempty" name:"SSHKeyId"`
}

type GetSSHKeyPairValueRequest struct {
	*tchttp.BaseRequest
	
	// Secret name. This field is only valid for SSH key secrets.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// ID of the key pair, which is the unique identifier of the key pair in the CVM.
	SSHKeyId *string `json:"SSHKeyId,omitnil,omitempty" name:"SSHKeyId"`
}

func (r *GetSSHKeyPairValueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSSHKeyPairValueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	delete(f, "SSHKeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetSSHKeyPairValueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSSHKeyPairValueResponseParams struct {
	// ID of the SSH key.
	SSHKeyID *string `json:"SSHKeyID,omitnil,omitempty" name:"SSHKeyID"`

	// Plaintext value of the Base64-encoded public key.
	PublicKey *string `json:"PublicKey,omitnil,omitempty" name:"PublicKey"`

	// Plaintext value of the Base64-encoded private key.
	PrivateKey *string `json:"PrivateKey,omitnil,omitempty" name:"PrivateKey"`

	// ID of the project to which the SSH key belongs.
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// Description of the SSH key.
	// The description can be modified in the CVM console.
	SSHKeyDescription *string `json:"SSHKeyDescription,omitnil,omitempty" name:"SSHKeyDescription"`

	// Name of the SSH key.
	// The name can be modified in the CVM console.
	SSHKeyName *string `json:"SSHKeyName,omitnil,omitempty" name:"SSHKeyName"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetSSHKeyPairValueResponse struct {
	*tchttp.BaseResponse
	Response *GetSSHKeyPairValueResponseParams `json:"Response"`
}

func (r *GetSSHKeyPairValueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSSHKeyPairValueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSecretValueRequestParams struct {
	// Name of a Secret.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// Specifies the version number of the corresponding credential.
	// For Tencent Cloud service credentials such as MySQL credentials, this API is used to get the plaintext information of a previously rotated credential by specifying the credential name and historical version number. If you want to get the plaintext of the credential version currently in use, you need to specify the version number as `SSM_Current`.
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

type GetSecretValueRequest struct {
	*tchttp.BaseRequest
	
	// Name of a Secret.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// Specifies the version number of the corresponding credential.
	// For Tencent Cloud service credentials such as MySQL credentials, this API is used to get the plaintext information of a previously rotated credential by specifying the credential name and historical version number. If you want to get the plaintext of the credential version currently in use, you need to specify the version number as `SSM_Current`.
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

func (r *GetSecretValueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSecretValueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	delete(f, "VersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetSecretValueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSecretValueResponseParams struct {
	// Name of the Secret.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// ID of the Secret version.
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// When creating a credential (CreateSecret), if you specify binary data, this field will be the Base64-encoded returned result. The application needs to Base64-decode the result to get the original data.
	// Either `SecretBinary` or `SecretString` cannot be empty.
	SecretBinary *string `json:"SecretBinary,omitnil,omitempty" name:"SecretBinary"`

	// When creating a credential (CreateSecret), if you specify general text data, this field will be the returned result.
	// Either `SecretBinary` or `SecretString` cannot be empty.
	SecretString *string `json:"SecretString,omitnil,omitempty" name:"SecretString"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetSecretValueResponse struct {
	*tchttp.BaseResponse
	Response *GetSecretValueResponseParams `json:"Response"`
}

func (r *GetSecretValueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSecretValueResponse) FromJsonString(s string) error {
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
	// `true`: The service is activated; `false`: The service is not activated.
	ServiceEnabled *bool `json:"ServiceEnabled,omitnil,omitempty" name:"ServiceEnabled"`

	// Invalid service type. `0`: not purchased; `1`: normal; `2`: suspended due to arrears; `3`: resource released
	InvalidType *int64 `json:"InvalidType,omitnil,omitempty" name:"InvalidType"`

	// `true`: Allow SSM to manage Tencent Cloud API key secrets.
	// `false`: Forbid SSM to manage Tencent Cloud API key secrets.
	AccessKeyEscrowEnabled *bool `json:"AccessKeyEscrowEnabled,omitnil,omitempty" name:"AccessKeyEscrowEnabled"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ListSecretVersionIdsRequestParams struct {
	// Name of the Secret.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`
}

type ListSecretVersionIdsRequest struct {
	*tchttp.BaseRequest
	
	// Name of the Secret.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`
}

func (r *ListSecretVersionIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSecretVersionIdsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListSecretVersionIdsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSecretVersionIdsResponseParams struct {
	// Name of the Secret.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// `VersionId` list.
	// Note: This field may return `null`, indicating that no valid value was found.
	Versions []*VersionInfo `json:"Versions,omitnil,omitempty" name:"Versions"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListSecretVersionIdsResponse struct {
	*tchttp.BaseResponse
	Response *ListSecretVersionIdsResponseParams `json:"Response"`
}

func (r *ListSecretVersionIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSecretVersionIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSecretsRequestParams struct {
	// Starting position of the list, starting at 0. If not specified, 0 is used by default.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of returned Secrets in a query. If not set or set to 0, 20 is used by default.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Sorting order according to the creation time. If not set or set to 0, descending order is used; if set to 1, ascending order is used.
	OrderType *uint64 `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// Filter based on credential status.
	// The default value is 0, indicating to query all.
	// 1: query the list of credentials in `Enabled` status.
	// 2: query the list of credentials in `Disabled` status.
	// 3: query the list of credentials in `PendingDelete` status.
	// 4: query the list of credentials in `PendingCreate` status.
	// 5: query the list of credentials in `CreateFailed` status.
	// The `PendingCreate` and `CreateFailed` status only take effect when `SecretType` is Tencent Cloud service credential
	State *uint64 `json:"State,omitnil,omitempty" name:"State"`

	// Filter according to Secret names. If left empty, this filter is not applied.
	SearchSecretName *string `json:"SearchSecretName,omitnil,omitempty" name:"SearchSecretName"`

	// Tag filter.
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`

	// `0` (default): user-defined secret.
	// `1`: Tencent Cloud services secret.
	// `2`: SSH key secret.
	// `3`: Tencent Cloud API key secret.
	SecretType *uint64 `json:"SecretType,omitnil,omitempty" name:"SecretType"`

	// This parameter only takes effect when the value of the SecretType parameter is 1.\nWhen the value of SecretType is `1`:
	// If the ProductName value is empty, it means querying all types of Tencent Cloud product secrets;If the ProductName value is a specific cloud product value such as MySQL, it means querying MySQL database credential;If the ProductName value is multiple cloud product values, such as: Mysql, Tdsql-mysql, Tdsql_C_Mysql (multiple values are separated by commas in English), it means querying the secrets of three cloud product types;To query the list of supported cloud products, use the interface: `DescribeSupportedProducts`.
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`
}

type ListSecretsRequest struct {
	*tchttp.BaseRequest
	
	// Starting position of the list, starting at 0. If not specified, 0 is used by default.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of returned Secrets in a query. If not set or set to 0, 20 is used by default.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Sorting order according to the creation time. If not set or set to 0, descending order is used; if set to 1, ascending order is used.
	OrderType *uint64 `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// Filter based on credential status.
	// The default value is 0, indicating to query all.
	// 1: query the list of credentials in `Enabled` status.
	// 2: query the list of credentials in `Disabled` status.
	// 3: query the list of credentials in `PendingDelete` status.
	// 4: query the list of credentials in `PendingCreate` status.
	// 5: query the list of credentials in `CreateFailed` status.
	// The `PendingCreate` and `CreateFailed` status only take effect when `SecretType` is Tencent Cloud service credential
	State *uint64 `json:"State,omitnil,omitempty" name:"State"`

	// Filter according to Secret names. If left empty, this filter is not applied.
	SearchSecretName *string `json:"SearchSecretName,omitnil,omitempty" name:"SearchSecretName"`

	// Tag filter.
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`

	// `0` (default): user-defined secret.
	// `1`: Tencent Cloud services secret.
	// `2`: SSH key secret.
	// `3`: Tencent Cloud API key secret.
	SecretType *uint64 `json:"SecretType,omitnil,omitempty" name:"SecretType"`

	// This parameter only takes effect when the value of the SecretType parameter is 1.\nWhen the value of SecretType is `1`:
	// If the ProductName value is empty, it means querying all types of Tencent Cloud product secrets;If the ProductName value is a specific cloud product value such as MySQL, it means querying MySQL database credential;If the ProductName value is multiple cloud product values, such as: Mysql, Tdsql-mysql, Tdsql_C_Mysql (multiple values are separated by commas in English), it means querying the secrets of three cloud product types;To query the list of supported cloud products, use the interface: `DescribeSupportedProducts`.
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`
}

func (r *ListSecretsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSecretsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderType")
	delete(f, "State")
	delete(f, "SearchSecretName")
	delete(f, "TagFilters")
	delete(f, "SecretType")
	delete(f, "ProductName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListSecretsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSecretsResponseParams struct {
	// Number of filtered Secrets according to `State` and `SearchSecretName`.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of Secret information.
	SecretMetadatas []*SecretMetadata `json:"SecretMetadatas,omitnil,omitempty" name:"SecretMetadatas"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListSecretsResponse struct {
	*tchttp.BaseResponse
	Response *ListSecretsResponseParams `json:"Response"`
}

func (r *ListSecretsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSecretsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductPrivilegeUnit struct {
	// Permission name. Valid values:
	// GlobalPrivileges
	// DatabasePrivileges
	// TablePrivileges
	// ColumnPrivileges
	// 
	// When the permission is `DatabasePrivileges`, the database name must be specified by the `Database` parameter;
	// 
	// When the permission is `TablePrivileges`, the database name and the table name in the database must be specified by the `Database` and `TableName` parameters;
	// 
	// When the permission is `ColumnPrivileges`, the database name, table name in the database, and column name in the table must be specified by the `Database`, `TableName`, and `ColumnName` parameters.
	PrivilegeName *string `json:"PrivilegeName,omitnil,omitempty" name:"PrivilegeName"`

	// Permission list.
	// For the `Mysql` service, optional permission values are:
	// 
	// 1. Valid values of `GlobalPrivileges`: "SELECT","INSERT","UPDATE","DELETE","CREATE", "PROCESS", "DROP","REFERENCES","INDEX","ALTER","SHOW DATABASES","CREATE TEMPORARY TABLES","LOCK TABLES","EXECUTE","CREATE VIEW","SHOW VIEW","CREATE ROUTINE","ALTER ROUTINE","EVENT","TRIGGER".
	// Note: if this parameter is not passed in, it means to clear the permission.
	// 
	// 2. Valid values of `DatabasePrivileges`: "SELECT","INSERT","UPDATE","DELETE","CREATE", "DROP","REFERENCES","INDEX","ALTER","CREATE TEMPORARY TABLES","LOCK TABLES","EXECUTE","CREATE VIEW","SHOW VIEW","CREATE ROUTINE","ALTER ROUTINE","EVENT","TRIGGER".
	// Note: if this parameter is not passed in, it means to clear the permission.
	// 
	// 3. Valid values of `TablePrivileges`: "SELECT","INSERT","UPDATE","DELETE","CREATE", "DROP","REFERENCES","INDEX","ALTER","CREATE VIEW","SHOW VIEW", "TRIGGER".
	// Note: if this parameter is not passed in, it means to clear the permission.
	// 
	// 4. Valid values of `ColumnPrivileges`: "SELECT","INSERT","UPDATE","REFERENCES".
	// Note: if this parameter is not passed in, it means to clear the permission.
	Privileges []*string `json:"Privileges,omitnil,omitempty" name:"Privileges"`

	// This value takes effect only when `PrivilegeName` is `DatabasePrivileges`.
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// This value takes effect only when `PrivilegeName` is `TablePrivileges`, and the `Database` parameter is required in this case to explicitly indicate the database instance.
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// This value takes effect only when `PrivilegeName` is `ColumnPrivileges`, and the following parameters are required in this case:
	// Database: explicitly indicate the database instance.
	// TableName: explicitly indicate the table
	ColumnName *string `json:"ColumnName,omitnil,omitempty" name:"ColumnName"`
}

// Predefined struct for user
type PutSecretValueRequestParams struct {
	// Name of a Secret where the version is added to.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// ID of the new Secret version. It can be up to 64 bytes, contain letters, digits, hyphens (-), and underscores (_), and must begin with a letter or digit.
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// Base64-encoded binary credential information.
	// Either `SecretBinary` or `SecretString` must be set.
	SecretBinary *string `json:"SecretBinary,omitnil,omitempty" name:"SecretBinary"`

	// Secret information plaintext in text format, base64 encoding is not needed. Either `SecretBinary` or `SecretString` must be set.
	SecretString *string `json:"SecretString,omitnil,omitempty" name:"SecretString"`
}

type PutSecretValueRequest struct {
	*tchttp.BaseRequest
	
	// Name of a Secret where the version is added to.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// ID of the new Secret version. It can be up to 64 bytes, contain letters, digits, hyphens (-), and underscores (_), and must begin with a letter or digit.
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// Base64-encoded binary credential information.
	// Either `SecretBinary` or `SecretString` must be set.
	SecretBinary *string `json:"SecretBinary,omitnil,omitempty" name:"SecretBinary"`

	// Secret information plaintext in text format, base64 encoding is not needed. Either `SecretBinary` or `SecretString` must be set.
	SecretString *string `json:"SecretString,omitnil,omitempty" name:"SecretString"`
}

func (r *PutSecretValueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PutSecretValueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	delete(f, "VersionId")
	delete(f, "SecretBinary")
	delete(f, "SecretString")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PutSecretValueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PutSecretValueResponseParams struct {
	// Name of the Secret.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// Version ID that is newly added.
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PutSecretValueResponse struct {
	*tchttp.BaseResponse
	Response *PutSecretValueResponseParams `json:"Response"`
}

func (r *PutSecretValueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PutSecretValueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestoreSecretRequestParams struct {
	// Name of the Secret to be restored.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`
}

type RestoreSecretRequest struct {
	*tchttp.BaseRequest
	
	// Name of the Secret to be restored.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`
}

func (r *RestoreSecretRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestoreSecretRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestoreSecretRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestoreSecretResponseParams struct {
	// Name of the Secret.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestoreSecretResponse struct {
	*tchttp.BaseResponse
	Response *RestoreSecretResponseParams `json:"Response"`
}

func (r *RestoreSecretResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestoreSecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RotateProductSecretRequestParams struct {
	// Name of the credential to be rotated.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`
}

type RotateProductSecretRequest struct {
	*tchttp.BaseRequest
	
	// Name of the credential to be rotated.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`
}

func (r *RotateProductSecretRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RotateProductSecretRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RotateProductSecretRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RotateProductSecretResponseParams struct {
	// Asynchronous rotation task ID. This field is valid when `SecretType` is `1` (i.e., the secret type is Tencent Cloud services secret, such as MySQL/TDSQL credentials).
	FlowID *int64 `json:"FlowID,omitnil,omitempty" name:"FlowID"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RotateProductSecretResponse struct {
	*tchttp.BaseResponse
	Response *RotateProductSecretResponseParams `json:"Response"`
}

func (r *RotateProductSecretResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RotateProductSecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecretMetadata struct {
	// Credential name
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// Credential description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// KMS `KeyId` used to encrypt the credential
	KmsKeyId *string `json:"KmsKeyId,omitnil,omitempty" name:"KmsKeyId"`

	// Creator UIN
	CreateUin *uint64 `json:"CreateUin,omitnil,omitempty" name:"CreateUin"`

	// Credential status: Enabled, Disabled, PendingDelete, Creating, Failed.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Credential deletion date, which takes effect for credentials in `PendingDelete` status and is in UNIX timestamp format
	DeleteTime *uint64 `json:"DeleteTime,omitnil,omitempty" name:"DeleteTime"`

	// Credential creation time in UNIX timestamp format
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Type of the KMS CMK used to encrypt the credential. `DEFAULT` represents the default key created by Secrets Manager, and `CUSTOMER` represents the user-specified key
	KmsKeyType *string `json:"KmsKeyType,omitnil,omitempty" name:"KmsKeyType"`

	// 1: enable rotation; 0: disable rotation
	// Note: this field may return null, indicating that no valid values can be obtained.
	RotationStatus *int64 `json:"RotationStatus,omitnil,omitempty" name:"RotationStatus"`

	// Start time of the next rotation in UNIX timestamp format
	// Note: this field may return null, indicating that no valid values can be obtained.
	NextRotationTime *uint64 `json:"NextRotationTime,omitnil,omitempty" name:"NextRotationTime"`

	// 0: custom secret;1: database credential;2: SSH key secret;3: cloud API key secret;4: Redis secret;Note: This field may return `null`, indicating no valid value can be obtained.
	SecretType *int64 `json:"SecretType,omitnil,omitempty" name:"SecretType"`

	// Tencent Cloud service name, which takes effect only when `SecretType` is 1 (Tencent Cloud service credential)
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// Secret name. This field is only valid when the `SecretType` is set to `2` (SSH key secret).
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// Project ID. This field is only valid when the `SecretType` is set to `2` (SSH key secret).
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ProjectID *int64 `json:"ProjectID,omitnil,omitempty" name:"ProjectID"`

	// ID of the CVM instance associated with the SSH key. ID. This field is only valid when the `SecretType` is set to `2` (SSH key secret).
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AssociatedInstanceIDs []*string `json:"AssociatedInstanceIDs,omitnil,omitempty" name:"AssociatedInstanceIDs"`

	// UIN of the Tencent Cloud API key. This field is valid when the secret type is Tencent Cloud API key secret.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TargetUin *uint64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// Rotation frequency in days. It takes effect when the rotation feature is enabled. 
	// Note: This field may return null, indicating that no valid values can be obtained.
	RotationFrequency *int64 `json:"RotationFrequency,omitnil,omitempty" name:"RotationFrequency"`

	// ID of Tencent Cloud resource corresponding with the Secret. 
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResourceID *string `json:"ResourceID,omitnil,omitempty" name:"ResourceID"`

	// The rotation start time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RotationBeginTime *string `json:"RotationBeginTime,omitnil,omitempty" name:"RotationBeginTime"`
}

type Tag struct {
	// Tag key
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TagFilter struct {
	// Tag key
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag value
	TagValue []*string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

// Predefined struct for user
type UpdateDescriptionRequestParams struct {
	// Name of a Secret whose description is to be updated.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// New description information, which can be up to 2048 bytes.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type UpdateDescriptionRequest struct {
	*tchttp.BaseRequest
	
	// Name of a Secret whose description is to be updated.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// New description information, which can be up to 2048 bytes.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *UpdateDescriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDescriptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDescriptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDescriptionResponseParams struct {
	// Name of the Secret.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateDescriptionResponse struct {
	*tchttp.BaseResponse
	Response *UpdateDescriptionResponseParams `json:"Response"`
}

func (r *UpdateDescriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDescriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRotationStatusRequestParams struct {
	// Tencent Cloud service credential name.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// Specifies whether to enable rotation.
	// `true`: enables rotation.
	// `false`: disables rotation.
	EnableRotation *bool `json:"EnableRotation,omitnil,omitempty" name:"EnableRotation"`

	// Rotation frequency in days. Value range: 30–365.
	Frequency *int64 `json:"Frequency,omitnil,omitempty" name:"Frequency"`

	// User-defined rotation start time in the format of 2006-01-02 15:04:05.
	// When `EnableRotation` is `true` and `RotationBeginTime` is left empty, the current time will be entered by default.
	RotationBeginTime *string `json:"RotationBeginTime,omitnil,omitempty" name:"RotationBeginTime"`
}

type UpdateRotationStatusRequest struct {
	*tchttp.BaseRequest
	
	// Tencent Cloud service credential name.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// Specifies whether to enable rotation.
	// `true`: enables rotation.
	// `false`: disables rotation.
	EnableRotation *bool `json:"EnableRotation,omitnil,omitempty" name:"EnableRotation"`

	// Rotation frequency in days. Value range: 30–365.
	Frequency *int64 `json:"Frequency,omitnil,omitempty" name:"Frequency"`

	// User-defined rotation start time in the format of 2006-01-02 15:04:05.
	// When `EnableRotation` is `true` and `RotationBeginTime` is left empty, the current time will be entered by default.
	RotationBeginTime *string `json:"RotationBeginTime,omitnil,omitempty" name:"RotationBeginTime"`
}

func (r *UpdateRotationStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRotationStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	delete(f, "EnableRotation")
	delete(f, "Frequency")
	delete(f, "RotationBeginTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateRotationStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRotationStatusResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateRotationStatusResponse struct {
	*tchttp.BaseResponse
	Response *UpdateRotationStatusResponseParams `json:"Response"`
}

func (r *UpdateRotationStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRotationStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateSecretRequestParams struct {
	// Name of a Secret whose content is to be updated.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// ID of the Secret version whose content is to be updated.
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// This field should be used and Base64-encoded if the content of the new credential is binary.
	// Either `SecretBinary` or `SecretString` cannot be empty.
	SecretBinary *string `json:"SecretBinary,omitnil,omitempty" name:"SecretBinary"`

	// This field should be used without being Base64-encoded if the content of the new credential is text. Either `SecretBinary` or `SecretString` cannot be empty.
	SecretString *string `json:"SecretString,omitnil,omitempty" name:"SecretString"`
}

type UpdateSecretRequest struct {
	*tchttp.BaseRequest
	
	// Name of a Secret whose content is to be updated.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// ID of the Secret version whose content is to be updated.
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// This field should be used and Base64-encoded if the content of the new credential is binary.
	// Either `SecretBinary` or `SecretString` cannot be empty.
	SecretBinary *string `json:"SecretBinary,omitnil,omitempty" name:"SecretBinary"`

	// This field should be used without being Base64-encoded if the content of the new credential is text. Either `SecretBinary` or `SecretString` cannot be empty.
	SecretString *string `json:"SecretString,omitnil,omitempty" name:"SecretString"`
}

func (r *UpdateSecretRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateSecretRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	delete(f, "VersionId")
	delete(f, "SecretBinary")
	delete(f, "SecretString")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateSecretRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateSecretResponseParams struct {
	// Name of the Secret.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// Version ID of the Secret.
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateSecretResponse struct {
	*tchttp.BaseResponse
	Response *UpdateSecretResponseParams `json:"Response"`
}

func (r *UpdateSecretResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateSecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VersionInfo struct {
	// Version ID.
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// Creation time, formatted as a Unix timestamp.
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}