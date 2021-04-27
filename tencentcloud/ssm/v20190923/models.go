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
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type CreateSecretRequest struct {
	*tchttp.BaseRequest

	// Secret name, which must be unique within a region. The name can be up to 128 bytes, contain letters, digits, hyphens (-), and underscores (_), and must begin with a letter or digit.
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// Secret version. It can be up to 64 bytes, contain letters, digits, hyphens (-), and underscores (_), and must begin with a letter or digit. `SecretName` and `VersionId` are used to query the Secret information.
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

	// Description information, such as the detailed use cases. It can be up to 2048 bytes.
	Description *string `json:"Description,omitempty" name:"Description"`

	// KMS CMK used for Secret encryption. If this parameter is left empty, SecretsManager will create a CMK by default. You can also specify a KMS CMK that is created in the same region.
	KmsKeyId *string `json:"KmsKeyId,omitempty" name:"KmsKeyId"`

	// Base64-encoded plaintext of a binary Secret. Either `SecretBinary` or `SecretString` must be set. A maximum of 4096 bytes is supported.
	SecretBinary *string `json:"SecretBinary,omitempty" name:"SecretBinary"`

	// Plaintext of a Secret, in text format. Base64 encoding is not required. Either `SecretBinary` or `SecretString` must be set. A maximum of 4096 bytes is supported.
	SecretString *string `json:"SecretString,omitempty" name:"SecretString"`

	// List of tags.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`
}

func (r *CreateSecretRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
	delete(f, "SecretBinary")
	delete(f, "SecretString")
	delete(f, "Tags")
	if len(f) > 0 {
		return errors.New("CreateSecretRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateSecretResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Name of the new Secret.
		SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

		// ID of the new Secret version.
		VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

		// Return code of tag operation. `0`: success; `1`: internal error; `2`: business processing error
	// Note: This field may return `null`, indicating that no valid value was found.
		TagCode *uint64 `json:"TagCode,omitempty" name:"TagCode"`

		// Return message of tag operation.
	// Note: This field may return `null`, indicating that no valid value was found.
		TagMsg *string `json:"TagMsg,omitempty" name:"TagMsg"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSecretResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecretRequest struct {
	*tchttp.BaseRequest

	// Name of the Secret to be deleted.
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// Scheduled deletion time, in days. If set to 0, the Secret is deleted immediately. A number in the range of 1 to 30 indicates the number of retention days. The Secret will be deleted after the set value.
	RecoveryWindowInDays *uint64 `json:"RecoveryWindowInDays,omitempty" name:"RecoveryWindowInDays"`
}

func (r *DeleteSecretRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecretRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	delete(f, "RecoveryWindowInDays")
	if len(f) > 0 {
		return errors.New("DeleteSecretRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecretResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Name of deleted Secret.
		SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

		// Secret deletion time, formatted as a Unix timestamp.
		DeleteTime *int64 `json:"DeleteTime,omitempty" name:"DeleteTime"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSecretResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecretVersionRequest struct {
	*tchttp.BaseRequest

	// Name of the Secret.
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// ID of the Secret version to be deleted.
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`
}

func (r *DeleteSecretVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecretVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	delete(f, "VersionId")
	if len(f) > 0 {
		return errors.New("DeleteSecretVersionRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecretVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Name of the Secret.
		SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

		// Version ID of the Secret.
		VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSecretVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecretVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecretRequest struct {
	*tchttp.BaseRequest

	// Name of a Secret whose detailed information is to be obtained.
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`
}

func (r *DescribeSecretRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecretRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	if len(f) > 0 {
		return errors.New("DescribeSecretRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecretResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Name of the Secret.
		SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

		// Description of the Secret.
		Description *string `json:"Description,omitempty" name:"Description"`

		// ID of the KMS CMK used for encryption.
		KmsKeyId *string `json:"KmsKeyId,omitempty" name:"KmsKeyId"`

		// Creator UIN.
		CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

		// Secret status, which can be `Enabled`, `Disabled`, or `PendingDelete`.
		Status *string `json:"Status,omitempty" name:"Status"`

		// Deletion time, formatted as a Unix timestamp. For a Secret that is not in `PendingDelete` status, this value is 0.
		DeleteTime *uint64 `json:"DeleteTime,omitempty" name:"DeleteTime"`

		// Creation time.
		CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecretResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableSecretRequest struct {
	*tchttp.BaseRequest

	// Name of the Secret to be disabled.
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`
}

func (r *DisableSecretRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableSecretRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	if len(f) > 0 {
		return errors.New("DisableSecretRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DisableSecretResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Name of the disabled Secret.
		SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableSecretResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableSecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableSecretRequest struct {
	*tchttp.BaseRequest

	// Name of the Secret to be enabled.
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`
}

func (r *EnableSecretRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableSecretRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	if len(f) > 0 {
		return errors.New("EnableSecretRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type EnableSecretResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Name of the enabled Secret.
		SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableSecretResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableSecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRegionsRequest struct {
	*tchttp.BaseRequest
}

func (r *GetRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.New("GetRegionsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetRegionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of regions.
		Regions []*string `json:"Regions,omitempty" name:"Regions" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSecretValueRequest struct {
	*tchttp.BaseRequest

	// Name of a Secret.
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// ID of the Secret version.
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`
}

func (r *GetSecretValueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSecretValueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	delete(f, "VersionId")
	if len(f) > 0 {
		return errors.New("GetSecretValueRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetSecretValueResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Name of the Secret.
		SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

		// ID of the Secret version.
		VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

		// If the `SecretBinary` field in the request body is specified in the `CreateSecret` call, this field is returned and base64-encoded. The caller needs to perform base64 decoding to obtain the original data. Either `SecretBinary` or `SecretString` will be returned.
		SecretBinary *string `json:"SecretBinary,omitempty" name:"SecretBinary"`

		// If the `SecretString` field in the request body is specified in the `CreateSecret` call, this field is returned. Either `SecretBinary` or `SecretString` will be returned.
		SecretString *string `json:"SecretString,omitempty" name:"SecretString"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSecretValueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSecretValueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetServiceStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *GetServiceStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetServiceStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.New("GetServiceStatusRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetServiceStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// `true`: The service is activated; `false`: The service is not activated.
		ServiceEnabled *bool `json:"ServiceEnabled,omitempty" name:"ServiceEnabled"`

		// Invalid service type. `0`: not purchased; `1`: normal; `2`: suspended due to arrears; `3`: resource released
		InvalidType *int64 `json:"InvalidType,omitempty" name:"InvalidType"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetServiceStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetServiceStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSecretVersionIdsRequest struct {
	*tchttp.BaseRequest

	// Name of the Secret.
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`
}

func (r *ListSecretVersionIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSecretVersionIdsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	if len(f) > 0 {
		return errors.New("ListSecretVersionIdsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListSecretVersionIdsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Name of the Secret.
		SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

		// `VersionId` list.
	// Note: This field may return `null`, indicating that no valid value was found.
		Versions []*VersionInfo `json:"Versions,omitempty" name:"Versions" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListSecretVersionIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSecretVersionIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSecretsRequest struct {
	*tchttp.BaseRequest

	// Starting position of the list, starting at 0. If not specified, 0 is used by default.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of returned Secrets in a query. If not set or set to 0, 20 is used by default.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Sorting order according to the creation time. If not set or set to 0, descending order is used; if set to 1, ascending order is used.
	OrderType *uint64 `json:"OrderType,omitempty" name:"OrderType"`

	// Filter according to Secret statuses. `0` (default): all Secrets; `1`: Secrets in `Enabled` status; `2`: Secrets in `Disabled` status; `3`: Secrets in `PendingDelete` status.
	State *uint64 `json:"State,omitempty" name:"State"`

	// Filter according to Secret names. If left empty, this filter is not applied.
	SearchSecretName *string `json:"SearchSecretName,omitempty" name:"SearchSecretName"`

	// Tag filter condition.
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters" list`
}

func (r *ListSecretsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
	if len(f) > 0 {
		return errors.New("ListSecretsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListSecretsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of filtered Secrets according to `State` and `SearchSecretName`.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of Secret information.
		SecretMetadatas []*SecretMetadata `json:"SecretMetadatas,omitempty" name:"SecretMetadatas" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListSecretsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSecretsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PutSecretValueRequest struct {
	*tchttp.BaseRequest

	// Name of a Secret where the version is added to.
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// ID of the new Secret version. It can be up to 64 bytes, contain letters, digits, hyphens (-), and underscores (_), and must begin with a letter or digit.
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

	// Binary Secret information that is base64-encoded. Either `SecretBinary` or `SecretString` must be set.
	SecretBinary *string `json:"SecretBinary,omitempty" name:"SecretBinary"`

	// Secret information plaintext in text format, base64 encoding is not needed. Either `SecretBinary` or `SecretString` must be set.
	SecretString *string `json:"SecretString,omitempty" name:"SecretString"`
}

func (r *PutSecretValueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("PutSecretValueRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type PutSecretValueResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Name of the Secret.
		SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

		// Version ID that is newly added.
		VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PutSecretValueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PutSecretValueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestoreSecretRequest struct {
	*tchttp.BaseRequest

	// Name of the Secret to be restored.
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`
}

func (r *RestoreSecretRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestoreSecretRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	if len(f) > 0 {
		return errors.New("RestoreSecretRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RestoreSecretResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Name of the Secret.
		SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RestoreSecretResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestoreSecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecretMetadata struct {

	// Name of the Secret.
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// Description of the Secret.
	Description *string `json:"Description,omitempty" name:"Description"`

	// KMS Key ID used for Secret encryption.
	KmsKeyId *string `json:"KmsKeyId,omitempty" name:"KmsKeyId"`

	// Creator UIN.
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// Secret status, which can be `Enabled`, `Disabled`, or `PendingDelete`.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Secret deletion time, formatted as a Unix timestamp. This parameter is only applicable for Secrets in `PendingDelete` status.
	DeleteTime *uint64 `json:"DeleteTime,omitempty" name:"DeleteTime"`

	// Secret creation time, formatted as a Unix timestamp.
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// Type of KMS CMK used for Secret encryption. `DEFAULT`: default key created by SecretsManager; `CUSTOMER`: user-specified key.
	KmsKeyType *string `json:"KmsKeyType,omitempty" name:"KmsKeyType"`
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
	TagValue []*string `json:"TagValue,omitempty" name:"TagValue" list`
}

type UpdateDescriptionRequest struct {
	*tchttp.BaseRequest

	// Name of a Secret whose description is to be updated.
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// New description information, which can be up to 2048 bytes.
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *UpdateDescriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDescriptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	delete(f, "Description")
	if len(f) > 0 {
		return errors.New("UpdateDescriptionRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDescriptionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Name of the Secret.
		SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateDescriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDescriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSecretRequest struct {
	*tchttp.BaseRequest

	// Name of a Secret whose content is to be updated.
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// ID of the Secret version whose content is to be updated.
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

	// Use this field if the new Secret content is in binary format, and base64-encoded. Either `SecretBinary` or `SecretString` is set.
	SecretBinary *string `json:"SecretBinary,omitempty" name:"SecretBinary"`

	// Use this field if the new Secret content is in text format, and base64-encoding is not required. Either `SecretBinary` or `SecretString` is set.
	SecretString *string `json:"SecretString,omitempty" name:"SecretString"`
}

func (r *UpdateSecretRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("UpdateSecretRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateSecretResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Name of the Secret.
		SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

		// Version ID of the Secret.
		VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateSecretResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateSecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VersionInfo struct {

	// Version ID.
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

	// Creation time, formatted as a Unix timestamp.
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
}
