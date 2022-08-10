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

package v20180813

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type AssumeRoleRequestParams struct {
	// Resource descriptions of a role, which can be obtained by clicking the role name in the [CAM console](https://console.cloud.tencent.com/cam/role).
	// General role:
	// qcs::cam::uin/12345678:role/4611686018427397919, qcs::cam::uin/12345678:roleName/testRoleName
	// Service role:
	// qcs::cam::uin/12345678:role/tencentcloudServiceRole/4611686018427397920, qcs::cam::uin/12345678:role/tencentcloudServiceRoleName/testServiceRoleName
	RoleArn *string `json:"RoleArn,omitempty" name:"RoleArn"`

	// User-defined temporary session name.
	// It can contain 2-128 letters, digits, and symbols (=,.@_-). Regex: [\w+=,.@_-]*
	RoleSessionName *string `json:"RoleSessionName,omitempty" name:"RoleSessionName"`

	// Specifies the validity period of credentials in seconds. Default value: 7200. Maximum value: 43200
	DurationSeconds *uint64 `json:"DurationSeconds,omitempty" name:"DurationSeconds"`

	// Policy description
	// Note:
	// 1. The policy needs to be URL-encoded (if you request a TencentCloud API through the GET method, all parameters must be URL-encoded again in accordance with [Signature v3](https://intl.cloud.tencent.com/document/api/598/33159?from_cn_redirect=1#1.-.E6.8B.BC.E6.8E.A5.E8.A7.84.E8.8C.83.E8.AF.B7.E6.B1.82.E4.B8.B2) before the request is sent).
	// 2. For the policy syntax, please see CAM's [Syntax Logic](https://intl.cloud.tencent.com/document/product/598/10603?from_cn_redirect=1).
	// 3. The policy cannot contain the `principal` element.
	Policy *string `json:"Policy,omitempty" name:"Policy"`

	// External role ID, which can be obtained by clicking the role name in the [CAM console](https://console.cloud.tencent.com/cam/role).
	// It can contain 2-128 letters, digits, and symbols (=,.@:/-). Regex: [\w+=,.@:\/-]*
	ExternalId *string `json:"ExternalId,omitempty" name:"ExternalId"`

	// List of session tags. Up to 50 tags are allowed. The tag keys can not duplicate.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type AssumeRoleRequest struct {
	*tchttp.BaseRequest
	
	// Resource descriptions of a role, which can be obtained by clicking the role name in the [CAM console](https://console.cloud.tencent.com/cam/role).
	// General role:
	// qcs::cam::uin/12345678:role/4611686018427397919, qcs::cam::uin/12345678:roleName/testRoleName
	// Service role:
	// qcs::cam::uin/12345678:role/tencentcloudServiceRole/4611686018427397920, qcs::cam::uin/12345678:role/tencentcloudServiceRoleName/testServiceRoleName
	RoleArn *string `json:"RoleArn,omitempty" name:"RoleArn"`

	// User-defined temporary session name.
	// It can contain 2-128 letters, digits, and symbols (=,.@_-). Regex: [\w+=,.@_-]*
	RoleSessionName *string `json:"RoleSessionName,omitempty" name:"RoleSessionName"`

	// Specifies the validity period of credentials in seconds. Default value: 7200. Maximum value: 43200
	DurationSeconds *uint64 `json:"DurationSeconds,omitempty" name:"DurationSeconds"`

	// Policy description
	// Note:
	// 1. The policy needs to be URL-encoded (if you request a TencentCloud API through the GET method, all parameters must be URL-encoded again in accordance with [Signature v3](https://intl.cloud.tencent.com/document/api/598/33159?from_cn_redirect=1#1.-.E6.8B.BC.E6.8E.A5.E8.A7.84.E8.8C.83.E8.AF.B7.E6.B1.82.E4.B8.B2) before the request is sent).
	// 2. For the policy syntax, please see CAM's [Syntax Logic](https://intl.cloud.tencent.com/document/product/598/10603?from_cn_redirect=1).
	// 3. The policy cannot contain the `principal` element.
	Policy *string `json:"Policy,omitempty" name:"Policy"`

	// External role ID, which can be obtained by clicking the role name in the [CAM console](https://console.cloud.tencent.com/cam/role).
	// It can contain 2-128 letters, digits, and symbols (=,.@:/-). Regex: [\w+=,.@:\/-]*
	ExternalId *string `json:"ExternalId,omitempty" name:"ExternalId"`

	// List of session tags. Up to 50 tags are allowed. The tag keys can not duplicate.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *AssumeRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssumeRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoleArn")
	delete(f, "RoleSessionName")
	delete(f, "DurationSeconds")
	delete(f, "Policy")
	delete(f, "ExternalId")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssumeRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssumeRoleResponseParams struct {
	// Temporary security credentials
	Credentials *Credentials `json:"Credentials,omitempty" name:"Credentials"`

	// Credentials expiration time. A Unix timestamp will be returned which is accurate to the second
	ExpiredTime *int64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// Credentials expiration time in UTC time in ISO 8601 format.
	Expiration *string `json:"Expiration,omitempty" name:"Expiration"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AssumeRoleResponse struct {
	*tchttp.BaseResponse
	Response *AssumeRoleResponseParams `json:"Response"`
}

func (r *AssumeRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssumeRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssumeRoleWithSAMLRequestParams struct {
	// Base64-encoded SAML assertion
	SAMLAssertion *string `json:"SAMLAssertion,omitempty" name:"SAMLAssertion"`

	// Principal access description name
	PrincipalArn *string `json:"PrincipalArn,omitempty" name:"PrincipalArn"`

	// Role access description name
	RoleArn *string `json:"RoleArn,omitempty" name:"RoleArn"`

	// Session name
	RoleSessionName *string `json:"RoleSessionName,omitempty" name:"RoleSessionName"`

	// The validity period of the temporary credentials in seconds. Default value: 7,200s. Maximum value: 43,200s.
	DurationSeconds *uint64 `json:"DurationSeconds,omitempty" name:"DurationSeconds"`
}

type AssumeRoleWithSAMLRequest struct {
	*tchttp.BaseRequest
	
	// Base64-encoded SAML assertion
	SAMLAssertion *string `json:"SAMLAssertion,omitempty" name:"SAMLAssertion"`

	// Principal access description name
	PrincipalArn *string `json:"PrincipalArn,omitempty" name:"PrincipalArn"`

	// Role access description name
	RoleArn *string `json:"RoleArn,omitempty" name:"RoleArn"`

	// Session name
	RoleSessionName *string `json:"RoleSessionName,omitempty" name:"RoleSessionName"`

	// The validity period of the temporary credentials in seconds. Default value: 7,200s. Maximum value: 43,200s.
	DurationSeconds *uint64 `json:"DurationSeconds,omitempty" name:"DurationSeconds"`
}

func (r *AssumeRoleWithSAMLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssumeRoleWithSAMLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SAMLAssertion")
	delete(f, "PrincipalArn")
	delete(f, "RoleArn")
	delete(f, "RoleSessionName")
	delete(f, "DurationSeconds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssumeRoleWithSAMLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssumeRoleWithSAMLResponseParams struct {
	// An object consists of the `Token`, `TmpSecretId`, and `TmpSecretId`
	Credentials *Credentials `json:"Credentials,omitempty" name:"Credentials"`

	// Credentials expiration time. A Unix timestamp will be returned which is accurate to the second
	ExpiredTime *uint64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// Credentials expiration time in UTC time in ISO 8601 format.
	Expiration *string `json:"Expiration,omitempty" name:"Expiration"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AssumeRoleWithSAMLResponse struct {
	*tchttp.BaseResponse
	Response *AssumeRoleWithSAMLResponseParams `json:"Response"`
}

func (r *AssumeRoleWithSAMLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssumeRoleWithSAMLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssumeRoleWithWebIdentityRequestParams struct {
	// Identity provider name
	ProviderId *string `json:"ProviderId,omitempty" name:"ProviderId"`

	// OIDC token issued by the IdP
	WebIdentityToken *string `json:"WebIdentityToken,omitempty" name:"WebIdentityToken"`

	// Role access description name
	RoleArn *string `json:"RoleArn,omitempty" name:"RoleArn"`

	// Session name
	RoleSessionName *string `json:"RoleSessionName,omitempty" name:"RoleSessionName"`

	// The validity period of the temporary credential in seconds. Default value: 7,200s. Maximum value: 43,200s.
	DurationSeconds *int64 `json:"DurationSeconds,omitempty" name:"DurationSeconds"`
}

type AssumeRoleWithWebIdentityRequest struct {
	*tchttp.BaseRequest
	
	// Identity provider name
	ProviderId *string `json:"ProviderId,omitempty" name:"ProviderId"`

	// OIDC token issued by the IdP
	WebIdentityToken *string `json:"WebIdentityToken,omitempty" name:"WebIdentityToken"`

	// Role access description name
	RoleArn *string `json:"RoleArn,omitempty" name:"RoleArn"`

	// Session name
	RoleSessionName *string `json:"RoleSessionName,omitempty" name:"RoleSessionName"`

	// The validity period of the temporary credential in seconds. Default value: 7,200s. Maximum value: 43,200s.
	DurationSeconds *int64 `json:"DurationSeconds,omitempty" name:"DurationSeconds"`
}

func (r *AssumeRoleWithWebIdentityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssumeRoleWithWebIdentityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProviderId")
	delete(f, "WebIdentityToken")
	delete(f, "RoleArn")
	delete(f, "RoleSessionName")
	delete(f, "DurationSeconds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssumeRoleWithWebIdentityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssumeRoleWithWebIdentityResponseParams struct {
	// Expiration time of the temporary credential (timestamp)
	ExpiredTime *uint64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// Expiration time of the temporary credential
	Expiration *string `json:"Expiration,omitempty" name:"Expiration"`

	// Temporary credential
	Credentials *Credentials `json:"Credentials,omitempty" name:"Credentials"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AssumeRoleWithWebIdentityResponse struct {
	*tchttp.BaseResponse
	Response *AssumeRoleWithWebIdentityResponseParams `json:"Response"`
}

func (r *AssumeRoleWithWebIdentityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssumeRoleWithWebIdentityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Credentials struct {
	// Token, which contains up to 4,096 bytes depending on the associated policies.
	Token *string `json:"Token,omitempty" name:"Token"`

	// Temporary credentials key ID, which contains up to 1,024 bytes.
	TmpSecretId *string `json:"TmpSecretId,omitempty" name:"TmpSecretId"`

	// Temporary credentials key, which contains up to 1,024 bytes.
	TmpSecretKey *string `json:"TmpSecretKey,omitempty" name:"TmpSecretKey"`
}

// Predefined struct for user
type GetCallerIdentityRequestParams struct {

}

type GetCallerIdentityRequest struct {
	*tchttp.BaseRequest
	
}

func (r *GetCallerIdentityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCallerIdentityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCallerIdentityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCallerIdentityResponseParams struct {
	// ARN of the current caller.
	Arn *string `json:"Arn,omitempty" name:"Arn"`

	// Root account UIN of the current caller.
	AccountId *string `json:"AccountId,omitempty" name:"AccountId"`

	// User ID.
	// 1. If the caller is a Tencent Cloud account, the UIN of the current account is returned.
	// 2. If the caller is a role, `roleId:roleSessionName` is returned.
	// 3. If the caller is a federated user, `uin:federatedUserName` is returned.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Account UIN.
	// 1. If the caller is a Tencent Cloud account, the UIN of the current account is returned.
	// 2. If the caller is a role, the UIN of the account that applies for the role is returned.
	PrincipalId *string `json:"PrincipalId,omitempty" name:"PrincipalId"`

	// Identity type.
	Type *string `json:"Type,omitempty" name:"Type"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetCallerIdentityResponse struct {
	*tchttp.BaseResponse
	Response *GetCallerIdentityResponseParams `json:"Response"`
}

func (r *GetCallerIdentityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCallerIdentityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFederationTokenRequestParams struct {
	// The customizable name of the caller, consisting of letters
	Name *string `json:"Name,omitempty" name:"Name"`

	// Policy description
	// Note:
	// 1. The policy needs to be URL-encoded (if you request a TencentCloud API through the GET method, all parameters must be URL-encoded again in accordance with [Signature v3](https://intl.cloud.tencent.com/document/api/598/33159?from_cn_redirect=1#1.-.E6.8B.BC.E6.8E.A5.E8.A7.84.E8.8C.83.E8.AF.B7.E6.B1.82.E4.B8.B2) before the request is sent).
	// 2. For the policy syntax, please see CAM's [Syntax Logic](https://intl.cloud.tencent.com/document/product/598/10603?from_cn_redirect=1).
	// 3. The policy cannot contain the `principal` element.
	Policy *string `json:"Policy,omitempty" name:"Policy"`

	// The validity period of temporary credentials in seconds. Default value: 1,800s. Maximum value for a root account: 7,200s. Maximum value for a sub-account: 129,600s.
	DurationSeconds *uint64 `json:"DurationSeconds,omitempty" name:"DurationSeconds"`
}

type GetFederationTokenRequest struct {
	*tchttp.BaseRequest
	
	// The customizable name of the caller, consisting of letters
	Name *string `json:"Name,omitempty" name:"Name"`

	// Policy description
	// Note:
	// 1. The policy needs to be URL-encoded (if you request a TencentCloud API through the GET method, all parameters must be URL-encoded again in accordance with [Signature v3](https://intl.cloud.tencent.com/document/api/598/33159?from_cn_redirect=1#1.-.E6.8B.BC.E6.8E.A5.E8.A7.84.E8.8C.83.E8.AF.B7.E6.B1.82.E4.B8.B2) before the request is sent).
	// 2. For the policy syntax, please see CAM's [Syntax Logic](https://intl.cloud.tencent.com/document/product/598/10603?from_cn_redirect=1).
	// 3. The policy cannot contain the `principal` element.
	Policy *string `json:"Policy,omitempty" name:"Policy"`

	// The validity period of temporary credentials in seconds. Default value: 1,800s. Maximum value for a root account: 7,200s. Maximum value for a sub-account: 129,600s.
	DurationSeconds *uint64 `json:"DurationSeconds,omitempty" name:"DurationSeconds"`
}

func (r *GetFederationTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFederationTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Policy")
	delete(f, "DurationSeconds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetFederationTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFederationTokenResponseParams struct {
	// Temporary credentials
	Credentials *Credentials `json:"Credentials,omitempty" name:"Credentials"`

	// Temporary credentials expiration time. A Unix timestamp will be returned which is accurate to the second
	ExpiredTime *uint64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// Credentials expiration time in UTC time in ISO 8601 format.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Expiration *string `json:"Expiration,omitempty" name:"Expiration"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetFederationTokenResponse struct {
	*tchttp.BaseResponse
	Response *GetFederationTokenResponseParams `json:"Response"`
}

func (r *GetFederationTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFederationTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {
	// Tag key. It’s up to 128 characters and case-sensitive.
	Key *string `json:"Key,omitempty" name:"Key"`

	// Tag value. It’s up to 256 characters and case-sensitive.
	Value *string `json:"Value,omitempty" name:"Value"`
}