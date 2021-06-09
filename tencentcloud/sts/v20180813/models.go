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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AssumeRoleRequest struct {
	*tchttp.BaseRequest

	// Role resource description, such as qcs::cam::uin/12345678:role/4611686018427397919, qcs::cam::uin/12345678:roleName/testRoleName
	RoleArn *string `json:"RoleArn,omitempty" name:"RoleArn"`

	// User-defined temporary session name
	RoleSessionName *string `json:"RoleSessionName,omitempty" name:"RoleSessionName"`

	// Specifies the validity period of credentials in seconds. Default value: 7200. Maximum value: 43200
	DurationSeconds *uint64 `json:"DurationSeconds,omitempty" name:"DurationSeconds"`

	// Policy description
	// Note:
	// 1. The policy needs to be URL-encoded (if you request a TencentCloud API through the GET method, all parameters must be URL-encoded again in accordance with [Signature v3](https://cloud.tencent.com/document/api/598/33159#1.-.E6.8B.BC.E6.8E.A5.E8.A7.84.E8.8C.83.E8.AF.B7.E6.B1.82.E4.B8.B2) before the request is sent).
	// 2. For the policy syntax, please see CAM’s [Syntax Logic](https://cloud.tencent.com/document/product/598/10603).
	// 3. The policy cannot contain the `principal` element.
	Policy *string `json:"Policy,omitempty" name:"Policy"`
}

func (r *AssumeRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssumeRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AssumeRoleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Temporary security credentials
		Credentials *Credentials `json:"Credentials,omitempty" name:"Credentials"`

		// Credentials expiration time. A Unix timestamp will be returned which is accurate to the second
		ExpiredTime *int64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

		// Credentials expiration time in UTC time in ISO 8601 format.
		Expiration *string `json:"Expiration,omitempty" name:"Expiration"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssumeRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssumeRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

	// Specifies the validity period of credentials in seconds. Default value: 7200. Maximum value: 7200
	DurationSeconds *uint64 `json:"DurationSeconds,omitempty" name:"DurationSeconds"`
}

func (r *AssumeRoleWithSAMLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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

type AssumeRoleWithSAMLResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// An object consists of the `Token`, `TmpSecretId`, and `TmpSecretId`
		Credentials *Credentials `json:"Credentials,omitempty" name:"Credentials"`

		// Credentials expiration time. A Unix timestamp will be returned which is accurate to the second
		ExpiredTime *uint64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

		// Credentials expiration time in UTC time in ISO 8601 format.
		Expiration *string `json:"Expiration,omitempty" name:"Expiration"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssumeRoleWithSAMLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssumeRoleWithSAMLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Credentials struct {

	// token
	Token *string `json:"Token,omitempty" name:"Token"`

	// Temporary credentials secret ID
	TmpSecretId *string `json:"TmpSecretId,omitempty" name:"TmpSecretId"`

	// Temporary credentials secret key
	TmpSecretKey *string `json:"TmpSecretKey,omitempty" name:"TmpSecretKey"`
}

type GetFederationTokenRequest struct {
	*tchttp.BaseRequest

	// The customizable name of the caller, consisting of letters
	Name *string `json:"Name,omitempty" name:"Name"`

	// Policy description
	// Note:
	// 1. The policy needs to be URL-encoded (if you request a TencentCloud API through the GET method, all parameters must be URL-encoded again in accordance with [Signature v3](https://cloud.tencent.com/document/api/598/33159#1.-.E6.8B.BC.E6.8E.A5.E8.A7.84.E8.8C.83.E8.AF.B7.E6.B1.82.E4.B8.B2) before the request is sent).
	// 2. For the policy syntax, please see CAM’s [Syntax Logic](https://cloud.tencent.com/document/product/598/10603).
	// 3. The policy cannot contain the `principal` element.
	Policy *string `json:"Policy,omitempty" name:"Policy"`

	// Specifies the validity period of credentials in seconds. Default value: 1800. Maximum value: 7200
	DurationSeconds *uint64 `json:"DurationSeconds,omitempty" name:"DurationSeconds"`
}

func (r *GetFederationTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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

type GetFederationTokenResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Temporary credentials
		Credentials *Credentials `json:"Credentials,omitempty" name:"Credentials"`

		// Temporary credentials expiration time. A Unix timestamp will be returned which is accurate to the second
		ExpiredTime *uint64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

		// Credentials expiration time in UTC time in ISO 8601 format.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Expiration *string `json:"Expiration,omitempty" name:"Expiration"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetFederationTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFederationTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
