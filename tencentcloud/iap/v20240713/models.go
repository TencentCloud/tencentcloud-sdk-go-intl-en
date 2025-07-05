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

package v20240713

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

// Predefined struct for user
type CreateIAPUserOIDCConfigRequestParams struct {
	// OpenID Connect IdP URL. It corresponds to the value of the "issuer" field in the Openid-configuration provided by the enterprise IdP.
	IdentityUrl *string `json:"IdentityUrl,omitnil,omitempty" name:"IdentityUrl"`

	// Client ID registered with the OpenID Connect IdP.
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// OpenID Connect IdP authorization endpoint. It corresponds to the value of the "authorization_endpoint" field in the Openid-configuration provided by the enterprise IdP.
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitnil,omitempty" name:"AuthorizationEndpoint"`

	// Authorization response type, which is always id_token.
	ResponseType *string `json:"ResponseType,omitnil,omitempty" name:"ResponseType"`

	// Authorization response mode. Valid values: form_post (recommended); fragment.
	ResponseMode *string `json:"ResponseMode,omitnil,omitempty" name:"ResponseMode"`

	// Mapping field name. It indicates which field in the id_token of the IdP is mapped to the username of a sub-user. It is usually the sub or name field.
	MappingFiled *string `json:"MappingFiled,omitnil,omitempty" name:"MappingFiled"`

	// Signature public key, which is used to verify the OpenID Connect IdP's ID token and must be Base64-encoded. For the security of your account, we recommend you rotate it regularly.
	IdentityKey *string `json:"IdentityKey,omitnil,omitempty" name:"IdentityKey"`

	// Authorization information scope. Valid values: openid (default); email; profile.
	Scope []*string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// Description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateIAPUserOIDCConfigRequest struct {
	*tchttp.BaseRequest
	
	// OpenID Connect IdP URL. It corresponds to the value of the "issuer" field in the Openid-configuration provided by the enterprise IdP.
	IdentityUrl *string `json:"IdentityUrl,omitnil,omitempty" name:"IdentityUrl"`

	// Client ID registered with the OpenID Connect IdP.
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// OpenID Connect IdP authorization endpoint. It corresponds to the value of the "authorization_endpoint" field in the Openid-configuration provided by the enterprise IdP.
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitnil,omitempty" name:"AuthorizationEndpoint"`

	// Authorization response type, which is always id_token.
	ResponseType *string `json:"ResponseType,omitnil,omitempty" name:"ResponseType"`

	// Authorization response mode. Valid values: form_post (recommended); fragment.
	ResponseMode *string `json:"ResponseMode,omitnil,omitempty" name:"ResponseMode"`

	// Mapping field name. It indicates which field in the id_token of the IdP is mapped to the username of a sub-user. It is usually the sub or name field.
	MappingFiled *string `json:"MappingFiled,omitnil,omitempty" name:"MappingFiled"`

	// Signature public key, which is used to verify the OpenID Connect IdP's ID token and must be Base64-encoded. For the security of your account, we recommend you rotate it regularly.
	IdentityKey *string `json:"IdentityKey,omitnil,omitempty" name:"IdentityKey"`

	// Authorization information scope. Valid values: openid (default); email; profile.
	Scope []*string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// Description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateIAPUserOIDCConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIAPUserOIDCConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdentityUrl")
	delete(f, "ClientId")
	delete(f, "AuthorizationEndpoint")
	delete(f, "ResponseType")
	delete(f, "ResponseMode")
	delete(f, "MappingFiled")
	delete(f, "IdentityKey")
	delete(f, "Scope")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIAPUserOIDCConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIAPUserOIDCConfigResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateIAPUserOIDCConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateIAPUserOIDCConfigResponseParams `json:"Response"`
}

func (r *CreateIAPUserOIDCConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIAPUserOIDCConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIAPLoginSessionDurationRequestParams struct {

}

type DescribeIAPLoginSessionDurationRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeIAPLoginSessionDurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIAPLoginSessionDurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIAPLoginSessionDurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIAPLoginSessionDurationResponseParams struct {
	// Login session duration.
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIAPLoginSessionDurationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIAPLoginSessionDurationResponseParams `json:"Response"`
}

func (r *DescribeIAPLoginSessionDurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIAPLoginSessionDurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIAPUserOIDCConfigRequestParams struct {

}

type DescribeIAPUserOIDCConfigRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeIAPUserOIDCConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIAPUserOIDCConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIAPUserOIDCConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIAPUserOIDCConfigResponseParams struct {
	// IdP type. 13: IAP user OIDC IdP.
	ProviderType *uint64 `json:"ProviderType,omitnil,omitempty" name:"ProviderType"`

	// IdP URL.
	IdentityUrl *string `json:"IdentityUrl,omitnil,omitempty" name:"IdentityUrl"`

	// Public key for signature.
	IdentityKey *string `json:"IdentityKey,omitnil,omitempty" name:"IdentityKey"`

	// Client ID.
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// Status. 0: Not set; 2: Disabled; 11: Enabled.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// The verification fingerprint of the HTTPS CA certificate. English letters and digits are allowed, and each fingerprint is 40 characters long, with a maximum of 5 fingerprints.
	Fingerprints []*string `json:"Fingerprints,omitnil,omitempty" name:"Fingerprints"`

	// Whether to enable the automatic use of the OIDC signature public key. 1: Yes, 2: No (default).
	EnableAutoPublicKey *uint64 `json:"EnableAutoPublicKey,omitnil,omitempty" name:"EnableAutoPublicKey"`

	// Authorization endpoint.
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitnil,omitempty" name:"AuthorizationEndpoint"`

	// Authorization scope.
	Scope []*string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// Authorization response type.
	ResponseType *string `json:"ResponseType,omitnil,omitempty" name:"ResponseType"`

	// Authorization response mode.
	ResponseMode *string `json:"ResponseMode,omitnil,omitempty" name:"ResponseMode"`

	// Mapping field name.
	MappingFiled *string `json:"MappingFiled,omitnil,omitempty" name:"MappingFiled"`

	// Description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIAPUserOIDCConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIAPUserOIDCConfigResponseParams `json:"Response"`
}

func (r *DescribeIAPUserOIDCConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIAPUserOIDCConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableIAPUserSSORequestParams struct {

}

type DisableIAPUserSSORequest struct {
	*tchttp.BaseRequest
	
}

func (r *DisableIAPUserSSORequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableIAPUserSSORequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableIAPUserSSORequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableIAPUserSSOResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisableIAPUserSSOResponse struct {
	*tchttp.BaseResponse
	Response *DisableIAPUserSSOResponseParams `json:"Response"`
}

func (r *DisableIAPUserSSOResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableIAPUserSSOResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIAPLoginSessionDurationRequestParams struct {
	// Login session duration.
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`
}

type ModifyIAPLoginSessionDurationRequest struct {
	*tchttp.BaseRequest
	
	// Login session duration.
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`
}

func (r *ModifyIAPLoginSessionDurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIAPLoginSessionDurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Duration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyIAPLoginSessionDurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIAPLoginSessionDurationResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyIAPLoginSessionDurationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyIAPLoginSessionDurationResponseParams `json:"Response"`
}

func (r *ModifyIAPLoginSessionDurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIAPLoginSessionDurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateIAPUserOIDCConfigRequestParams struct {
	// OpenID Connect IdP URL. It corresponds to the value of the "issuer" field in the Openid-configuration provided by the enterprise IdP.
	IdentityUrl *string `json:"IdentityUrl,omitnil,omitempty" name:"IdentityUrl"`

	// Client ID registered with the OpenID Connect IdP.
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// OpenID Connect IdP authorization endpoint. It corresponds to the value of the "authorization_endpoint" field in the Openid-configuration provided by the enterprise IdP.
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitnil,omitempty" name:"AuthorizationEndpoint"`

	// Authorization response type, which is always id_token.
	ResponseType *string `json:"ResponseType,omitnil,omitempty" name:"ResponseType"`

	// Authorization response mode. Valid values: form_post (recommended); fragment.
	ResponseMode *string `json:"ResponseMode,omitnil,omitempty" name:"ResponseMode"`

	// Mapping field name. It indicates which field in the id_token of the IdP is mapped to the username of a sub-user. It is usually the sub or name field.
	MappingFiled *string `json:"MappingFiled,omitnil,omitempty" name:"MappingFiled"`

	// RSA signature public key in the JWKS format, which is used to verify the OpenID Connect IdP's ID token and must be Base64-encoded. For the security of your account, we recommend you rotate it regularly.
	IdentityKey *string `json:"IdentityKey,omitnil,omitempty" name:"IdentityKey"`

	// Authorization information scope. Valid values: openid (default); email; profile.
	Scope []*string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// Description, with a length of 1 to 255 English or Chinese characters. It is empty by default.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type UpdateIAPUserOIDCConfigRequest struct {
	*tchttp.BaseRequest
	
	// OpenID Connect IdP URL. It corresponds to the value of the "issuer" field in the Openid-configuration provided by the enterprise IdP.
	IdentityUrl *string `json:"IdentityUrl,omitnil,omitempty" name:"IdentityUrl"`

	// Client ID registered with the OpenID Connect IdP.
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// OpenID Connect IdP authorization endpoint. It corresponds to the value of the "authorization_endpoint" field in the Openid-configuration provided by the enterprise IdP.
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitnil,omitempty" name:"AuthorizationEndpoint"`

	// Authorization response type, which is always id_token.
	ResponseType *string `json:"ResponseType,omitnil,omitempty" name:"ResponseType"`

	// Authorization response mode. Valid values: form_post (recommended); fragment.
	ResponseMode *string `json:"ResponseMode,omitnil,omitempty" name:"ResponseMode"`

	// Mapping field name. It indicates which field in the id_token of the IdP is mapped to the username of a sub-user. It is usually the sub or name field.
	MappingFiled *string `json:"MappingFiled,omitnil,omitempty" name:"MappingFiled"`

	// RSA signature public key in the JWKS format, which is used to verify the OpenID Connect IdP's ID token and must be Base64-encoded. For the security of your account, we recommend you rotate it regularly.
	IdentityKey *string `json:"IdentityKey,omitnil,omitempty" name:"IdentityKey"`

	// Authorization information scope. Valid values: openid (default); email; profile.
	Scope []*string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// Description, with a length of 1 to 255 English or Chinese characters. It is empty by default.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *UpdateIAPUserOIDCConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateIAPUserOIDCConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdentityUrl")
	delete(f, "ClientId")
	delete(f, "AuthorizationEndpoint")
	delete(f, "ResponseType")
	delete(f, "ResponseMode")
	delete(f, "MappingFiled")
	delete(f, "IdentityKey")
	delete(f, "Scope")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateIAPUserOIDCConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateIAPUserOIDCConfigResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateIAPUserOIDCConfigResponse struct {
	*tchttp.BaseResponse
	Response *UpdateIAPUserOIDCConfigResponseParams `json:"Response"`
}

func (r *UpdateIAPUserOIDCConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateIAPUserOIDCConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}